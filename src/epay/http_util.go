package epay

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const (
	XMLHEADER = `<?xml version="1.0" encoding="UTF-8"?>`
	XMLNS     = `http://www.chinamobile.com/payment`
	EPAY_URL  = "http://%s:%d/epay?svc_cat=%s&svc_code=%s"
)

type GetPaymentOrderReq struct {
	XmlNS     string `xml:"xmlns,attr"`
	ChannelID string `xml:"ChannelID"`
	OrderID   string `xml:"OrderID"`
}

type EpayClient struct {
	IP   string
	Port int
}

func NewEpayClient(ip string, port int) EpayClient {
	return EpayClient{
		IP:   ip,
		Port: port,
	}
}

func (client *EpayClient) sendRequest(svc_cat string, svc_code string, req interface{}) (string, error) {
	buf := bytes.NewBufferString(XMLHEADER)
	body, err := xml.Marshal(req)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	buf.Write(body)

	fmt.Println("xmlbody", buf.String())
	url := fmt.Sprintf(EPAY_URL, client.IP, client.Port, svc_cat, svc_code)
	fmt.Println("url: ", url)
	resp, err := http.Post(url, "text/xml;charset=utf-8", strings.NewReader(buf.String()))
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(content), nil
}

func (client *EpayClient) GetPaymentOrder(channelId string, orderId string) (string, error) {
	req := GetPaymentOrderReq{
		XmlNS:     XMLNS,
		ChannelID: channelId,
		OrderID:   orderId,
	}

	resp, err := client.sendRequest("PayTx", "GetPaymentOrder", &req)
	if err != nil {
		log.Fatal(err)
	}
	return resp, nil
}
