package wsclient

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type Client struct {
	*http.Client
}

func NewClient(tr *http.Transport) *Client {
	var c *http.Client

	if tr == nil {
		c = http.DefaultClient
	} else {
		c = &http.Client{
			Transport: tr,
		}
	}

	return &Client{
		Client: c,
	}
}

const (
	apiEndpoint = "http://211.88.20.132:8040/services/syncServiceStation?wsdl"
	contentType = "application/soap+xml"
)

func (this *Client) Do(in interface{}) (body []byte, err error) {
	payload, err := toPayload(in)
	if err != nil {
		return
	}

	req, err := http.NewRequest(http.MethodPost, apiEndpoint, payload)
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", contentType)

	resp, err := this.Client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if resp.StatusCode/100 != 2 {
		err = errors.New(string(body))
		return
	}

	return
}

func toPayload(in interface{}) (reader io.Reader, err error) {
	buf, err := xml.Marshal(in)
	if err != nil {
		return
	}

	reader = strings.NewReader(fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <soapenv:Body>
    <syncServiceStationOperationRequest xmlns="http://www.cvicse.com/service/">
      <in xmlns="">
        %s
      </in>
    </syncServiceStationOperationRequest>
  </soapenv:Body>
</soapenv:Envelope>
    `, string(buf)))

	return
}
