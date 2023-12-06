package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"net/http"
	"slices"
	"strings"
	"time"
)

var METHODS = []string{"GET", "POST", "PUT", "DELETE"}

func Request(url string, method string, headers map[string]string, requestData any, responseData any) error {
	//log.Info("Request start: ", url, method, headers, requestData, responseData)
	if slices.Contains(METHODS, strings.ToUpper(method)) == true {
		bytesData, err := json.Marshal(requestData)
		if err != nil {
			log.Error("json error: ", err)
			return err
		} else {
			//log.Info("Requests requestData: ", requestDataJson)
			req, err := http.NewRequest(strings.ToUpper(method), url, bytes.NewReader(bytesData))
			//log.Info("Request NewRequest start: ", req)

			if err != nil {
				log.Error("Request NewRequest error: ", err)
				return err
			}
			// add headers
			for k, v := range headers {
				//log.Info("Request add header: ", k, v)
				req.Header.Add(k, v)
			}
			client := &http.Client{Timeout: time.Duration(5) * time.Second}
			resp, err := client.Do(req)

			if err != nil {
				log.Error("http request error: ", err)
				return err
			}
			defer resp.Body.Close()

			jsonErr := json.NewDecoder(resp.Body).Decode(responseData)
			if jsonErr != nil {
				log.Error("解析失败: ", jsonErr)
				return jsonErr
			}
			return nil
		}
	} else {
		log.Error("method: %s currently not supported, please use supported method in: %v", method, METHODS)
		return errors.New(fmt.Sprintf("method: %s currently not supported, please use supported method in: %v", method, METHODS))
	}
}
