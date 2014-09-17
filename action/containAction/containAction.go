package containAction

import (
	"fmt"
	"json/bean/fetch"
	"json/bean/search"
	. "json/config"
)

func ContainAction(sequence string) {
	path := "http://dict.youdao.com/mvoice?word=" + sequence
	b := fetch.FetchData(&path, IP)
	contain := search.SearchContain(b, "contain")
	fmt.Println(string(b), contain)	
}
