# wsclient

[农村电子商务平台](http://sqsyscjss.mofcom.gov.cn/login.jhtml) 日报，月报，年报数据提交客户端（Golang 版本）。

## Usage

Use `go get github.com/binatify/wsclient`

```golang
package main

import (
    "fmt"

    "github.com/binatify/wsclient"
)

type serviceStation struct{}

func main() {
    client := wsclient.NewClient(nil)

    var in serviceStation

    body, err := client.Do(in)

    fmt.Println(in)
    fmt.Println(string(body))
}
```

## WSAL SOAP协议声明

通过[链接](http://211.88.20.132:8040/services/syncServiceStation?wsdl)可以访问 syncServiceStation 服务的声明。

以 Restful 方式提交数据:

- `POST` 请求
- Content-Type: application/soap+xml
- Body 格式：

```xml
<?xml version="1.0" encoding="UTF-8"?>
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <soapenv:Body>
    <syncServiceStationOperationRequest xmlns="http://www.cvicse.com/service/">
      <in xmlns="">
        <serviceStation>
          <userId>xxx</userId>
          <rptDate>2017-12-08</rptDate>
          <serviceStationReport>
            <code>xx</code>
            <name>xx</name>
            <countyType>1</countyType>
            <buyOrder>5</buyOrder>
            <saleOrder>28</saleOrder>
            <serviceStationCommodity>
              <commId>5</commId>
              <money>123</money>
            </serviceStationCommodity>
            <serviceStationCommodity>
              <commId>28</commId>
              <money>123</money>
            </serviceStationCommodity>
          </serviceStationReport>
        </serviceStation>
      </in>
    </syncServiceStationOperationRequest>
  </soapenv:Body>
</soapenv:Envelope>
```
