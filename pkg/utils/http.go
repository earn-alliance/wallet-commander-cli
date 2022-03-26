package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	USER_AGENT = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/36.0.1944.0 Safari/537.36"
)

var (
	httpClient                    = http.Client{}
	DefaultAxieSiteRequestHeaders = map[string]string{
		"User-Agent": USER_AGENT,
	}
)

func CallGetHttpApi(apiUrl string, headers map[string]string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, apiUrl, nil)

	if err != nil {
		return nil, err
	}

	if headers != nil {
		for key, val := range headers {
			req.Header.Set(key, val)
		}
	}

	resp, err := httpClient.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return ioutil.ReadAll(resp.Body)
	}

	return nil, errors.New(fmt.Sprintf("received error from server with status code %d", resp.StatusCode))
}

func CallPostHttpApi(apiUrl string, jsonStruct interface{}, headers map[string]string) ([]byte, error) {
	var jsonData []byte
	if jsonStruct != nil {
		data, err := json.Marshal(jsonStruct)

		if err != nil {
			return nil, err
		}

		jsonData = data
	}

	req, err := http.NewRequest(http.MethodPost, apiUrl, bytes.NewBuffer(jsonData))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	if headers != nil {
		for key, val := range headers {
			req.Header.Set(key, val)
		}
	}

	resp, err := httpClient.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return ioutil.ReadAll(resp.Body)
	}

	return nil, errors.New(fmt.Sprintf("received error from server with status code %d", resp.StatusCode))
}
