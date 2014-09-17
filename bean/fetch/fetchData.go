package fetch

import (
	"json/third-pkg/httplib"
)

// fetch data by url
func FetchData(url *string, localIp string) []byte {
	if url == nil {
		return nil
	}
	request := httplib.Get(*url)
	request.Header("Host", "dict.youdao.com")
	request.Header("Accept", "*/*")
	request.Header("Connection", "keep-alive")
	request.Header("User-Agent", "AppStore/2.0 iOS/7.1.1 model/iPod5,1 build/11D201 (4; dt:81)")
	buff, err := request.Bytes()
	if err != nil {
		return nil
	}
	return buff
}
