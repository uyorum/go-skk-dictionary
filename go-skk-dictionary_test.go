package skkdictionary

import (
	"testing"
	"os"
)

const DICTIONARY_FILENAME = "SKK-JISYO.L.utf8"

type TestString struct {
	request string
	response string
	okuri_ari bool
}

var tests = []TestString{
	{"じしょ ", "/辞書/地所/自署;signature/字書;字引/自書;自筆/璽書/", false},
	{"ひk ", "/引/弾;ピアノを弾く/挽;コーヒー豆を挽く/退;軍を退く/碾;麦を碾く/轢;轢き殺す/曳;山車を曳く/牽;綱を牽く/索;見出しを索く/惹;(attract)心を惹かれる/魅;≒惹く/彈;「弾」の旧字(人名用漢字)/", true},
	{"じしょあ ", "", false},
	{"ひあk ", "", true},
	{"dictionary ", "/ディクショナリー/ディクショナリ/辞書/", false},
	{"a ", "/α;alpha/エー/エイ/アー;(独語)/а;cyrillic/ア/", false},
	{">おw ", "/終/", true},
	{"れべる# ", "/レベル#1/", false},
}

func TestSearch(t *testing.T) {
	pwd, _ := os.Getwd()
	dictionary_path := pwd + "/" + DICTIONARY_FILENAME
 	d := NewSkkDictionary(dictionary_path)
	for _, test := range tests {
		resp := d.Search(test.request)
		if resp != test.response {
			t.Error("Failed: " + test.request)
		}
	}
}

func TestIsOkuriAri(t *testing.T) {
	for _, test := range tests {
		resp := IsOkuriAri(test.request)
		if resp != test.okuri_ari {
			t.Error("Faild: ")
		}
	}
}
