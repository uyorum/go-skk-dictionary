# go-skk-dictionary

Golang library for searching into SKK dictionary file.

## Example

```go
package main

import (
	"github.com/uyorum/go-skk-dictionary"
	"fmt"
)

func main() {
	dictionary_path := "/path/to/SKK-JISYO.L.utf8"
	d := skkdictionary.NewSkkDictionary(dictionary_path)
	fmt.Println(d.Search("じしょ ")) // /辞書/地所/自署;signature/字書;字引/自書;自筆/璽書/
}
```

## Usage

1. Create dictionary object: ```d := skkdictionary.NewSkkDictionary("path to dictionary")```
1. Search entry: ```d.Search("search string")```  
  If no entry is matched, it will return nil string.  
  If you use a dictionary encoded in EUC-JP, you should be sure of the string also encoded in EUC-JP.

## Requirement
This library require skk dictionary to be sorted.  
e.g.) You should download from [here](http://openlab.ring.gr.jp/skk/wiki/wiki.cgi?page=SKK%BC%AD%BD%F1).
