package Util

import (
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func CurlRequestGet(getAddress string, header map[string]string, param map[string]string) []byte {
	req, _ := http.NewRequest(http.MethodGet, getAddress, nil)

	//参数设置
	if param != nil {
		q := req.URL.Query()
		for k, v := range param {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	//header头设置
	if header != nil {
		_, ok := header["Cookie"]
		if ok {
			req.Header.Add("Cookie", header["Cookie"])
		}
	}

	res, err := (&http.Client{}).Do(req)
	defer res.Body.Close()

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err":    err,
			"uri":    getAddress,
			"header": header,
		}).Error("CurlRequestGet请求异常")
	}
	resByte, _ := ioutil.ReadAll(res.Body)
	return resByte
}
