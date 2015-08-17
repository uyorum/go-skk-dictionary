package skkdictionary

import (
	"testing"
	"os"
)

const DICTIONARY_FILENAME = "/SKK-JISYO.L.utf8"

type TestString struct {
	request string
	response string
}

var tests = []TestString{
	{"じしょ ", "/辞書/地所/自署;signature/字書;字引/自書;自筆/璽書/"},
	{"ひk ", "/引/弾;ピアノを弾く/挽;コーヒー豆を挽く/退;軍を退く/碾;麦を碾く/轢;轢き殺す/曳;山車を曳く/牽;綱を牽く/索;見出しを索く/惹;(attract)心を惹かれる/魅;≒惹く/彈;「弾」の旧字(人名用漢字)/"},
	{"じしょあ ", ""},
	{"ひあk ", ""},
	{"dictionary ", "/ディクショナリー/ディクショナリ/辞書/"},
	{"a ", "/α;alpha/エー/エイ/アー;(独語)/а;cyrillic/ア/"},
	{">おw ", "/終/"},
	{"れべる# ", "/レベル#1/"},
}

func TestSearch(t *testing.T) {
	pwd, _ := os.Getwd()
	dictionary_path := pwd + DICTIONARY_FILENAME
 	d := NewSkkDictionary(dictionary_path)
	for _, test := range tests {
		resp := d.Search(test.request)
		if resp != test.response {
			t.Error("Failed at " + test.request)
		}
	}
}
