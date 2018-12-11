package common

import (
	"testing"
)

var secretKey = "e7709b5dfc2ccd28359eff07f08eba99"
var (
	params1 = map[string]interface{}{
		"mobile":  "18310422105",
		"message": "test",
	}
	params2 = map[string]interface{}{
		"mobile_list": "18310422105",
		"message":     "test",
		"time":        "",
	}
	basicAuth = map[string]string{
		"api": "key-" + secretKey,
	}
)

func Test_HTTPGetRequest(t *testing.T) {
	url := "http://www.baidu.com"
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	res, err := HTTPGetRequestWithBasicAuth(url, nil, headers, nil)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("GET from %v, result: %v", url, res)
}

func Test_Map2UrlQuery(t *testing.T) {
	params11 := map[string]interface{}{
		"mobile": "18310422106",
	}
	params1Str := Map2UrlQuery(params11)
	result := "mobile=18310422106"
	if params1Str != result {
		t.Errorf("%v convert to query string: %v, not equal %v", params11, params1Str, result)
	}
}
