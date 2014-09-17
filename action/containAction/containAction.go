package containAction

import (
	"fmt"
	"yodao/bean/fetch"
	"yodao/bean/search"
	. "yodao/config"
)

func ContainAction(sequence string) {
	path := "http://dict.youdao.com/mvoice?word=" + sequence
	b := fetch.FetchData(&path, IP)
	contain := search.SearchContain(b, "contain")
	fmt.Println(string(b), contain)
}
