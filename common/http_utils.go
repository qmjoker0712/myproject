package common

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// HTTPGetRequest http get request
func HTTPGetRequest(strURL string, params map[string]interface{}) (string, error) {
	return httpGetRequest(strURL, params, nil, nil)
}

// HTTPGetRequestWithHeaders http get request with headers
func HTTPGetRequestWithHeaders(strURL string, params map[string]interface{}, headers map[string]string) (string, error) {
	return httpGetRequest(strURL, params, headers, nil)
}

// HTTPGetRequestWithBasicAuth http get request with headers and basic
func HTTPGetRequestWithBasicAuth(strURL string, params map[string]interface{}, headers map[string]string, basicAuth map[string]string) (string, error) {
	return httpGetRequest(strURL, params, headers, basicAuth)
}

// httpGetRequest HTTP GET request
func httpGetRequest(strURL string, params map[string]interface{}, headers map[string]string, basicAuth map[string]string) (string, error) {
	httpClient := &http.Client{}
	var strRequestURL string
	if nil == params {
		strRequestURL = strURL
	} else {
		strParams := Map2UrlQuery(params)
		strRequestURL = strURL + "?" + strParams
	}

	request, err := http.NewRequest("GET", strRequestURL, nil)
	if nil != err {
		return err.Error(), err
	}
	for k, v := range headers {
		request.Header.Add(k, v)
	}
	if request.Header.Get("Content-Type") == "" {
		request.Header.Add("Content-Type", "application/json")
	}
	for k, v := range basicAuth {
		request.SetBasicAuth(k, v)
	}
	request.Close = true
	response, err := httpClient.Do(request)

	if err != nil {
		return err.Error(), err
	}
	if response != nil {
		defer response.Body.Close()
	}
	body, err := ioutil.ReadAll(response.Body)
	if nil != err {
		return err.Error(), err
	}
	return string(body), nil
}

// HTTPPostRequest http post request
func HTTPPostRequest(strURL string, params map[string]interface{}) (string, error) {
	return httpPostRequest(strURL, params, nil, nil)
}

// HTTPPostRequestWithHeaders http post request with headers
func HTTPPostRequestWithHeaders(strURL string, params map[string]interface{}, headers map[string]string) (string, error) {
	return httpPostRequest(strURL, params, headers, nil)
}

// HTTPPostRequestWithBasicAuth http post request with headers and basic
func HTTPPostRequestWithBasicAuth(strURL string, params map[string]interface{}, headers map[string]string, basicAuth map[string]string) (string, error) {
	return httpPostRequest(strURL, params, headers, basicAuth)
}

// HTTPPostRequest http post request
func httpPostRequest(strURL string, params map[string]interface{}, headers map[string]string, basicAuth map[string]string) (string, error) {
	httpClient := &http.Client{}
	strParams := Map2UrlQuery(params)
	payload := strings.NewReader(strParams)
	request, err := http.NewRequest("POST", strURL, payload)
	if nil != err {
		return err.Error(), err
	}
	for k, v := range headers {
		request.Header.Add(k, v)
	}
	if request.Header.Get("Content-Type") == "" {
		request.Header.Add("Content-Type", "application/json")
	}
	for k, v := range basicAuth {
		request.SetBasicAuth(k, v)
	}

	response, err := httpClient.Do(request)
	if nil != err {
		return err.Error(), err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if nil != err {
		return err.Error(), err
	}

	return string(body), nil
}

// Map2UrlQuery map to url query
func Map2UrlQuery(params map[string]interface{}) string {
	var strParams string
	for key, value := range params {
		strParams += fmt.Sprintf("%v=%v&", key, value)
	}
	if 0 < len(strParams) {
		bm := []rune(strParams)
		strParams = string(bm[:len(bm)-1])
	}
	return strParams
}
