package skkdictionary

import (
	"bytes"
	"bufio"
	"os"
)

const OKURI_ARI_LABEL = ";; okuri-ari entries."
const OKURI_NASI_LABEL = ";; okuri-nasi entries."

type SkkDictionary struct {
	okuri_ari_table []string
	okuri_nasi_table []string
}

func NewSkkDictionary(path string) *SkkDictionary {
	d := new(SkkDictionary)

	var dictionary_f *os.File
	var err error
	dictionary_f, err = os.Open(path)
	if err != nil {
		panic(err)
	}
	defer dictionary_f.Close()

	scanner := bufio.NewScanner(dictionary_f)
	for scanner.Scan() {
		if scanner.Text() == OKURI_ARI_LABEL {
			break
		}
	}
	for scanner.Scan() {
		if scanner.Text() == OKURI_NASI_LABEL {
			break
		}
		d.okuri_ari_table = append(d.okuri_ari_table, scanner.Text())
	}
	for scanner.Scan() {
		d.okuri_nasi_table = append(d.okuri_nasi_table, scanner.Text())
	}
	return d
}

func isOkuriAri(str string) bool {
	const A = 97
	const Z = 122
	str_b := []byte(str)
	if len(str_b) == 2 {
		return false
	}
	last_two_b := str_b[len(str_b) - 3:len(str_b) - 1]

	if (last_two_b[0] < A || Z < last_two_b[0]) && (A < last_two_b[1] && last_two_b[1] < Z) {
		return true
	}
	return false
}


func (d SkkDictionary) Search(kana string) string {
	var table []string
	var EQUAL = 0
	var GO_LEFT int
	var GO_RIGHT int

	if isOkuriAri(kana) {
		table = d.okuri_ari_table
		GO_LEFT = -1
		GO_RIGHT = 1
	} else {
		table = d.okuri_nasi_table
		GO_LEFT = 1
		GO_RIGHT = -1
	}

	kana_b := []byte(kana)
	left := 0
	right := len(table) - 1
	for left <= right {
		pos := (left + right) / 2
		line := table[pos]
		switch bytes.Compare([]byte(line)[:len(kana_b)], kana_b) {
		case EQUAL:
			return string([]byte(line)[len(kana_b):])
		case GO_LEFT:
			right = pos - 1
		case GO_RIGHT:
			left = pos + 1
		}
	}
	return ""
}
