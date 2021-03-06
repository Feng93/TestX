package main

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"time"
)

type Text struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   string   `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
	Content      string   `xml:"Content"`
}

type VXXML struct {
	XMLName xml.Name `xml:"xml"`
	Name    string   `xml:"xml"`
}

func main() {
	reply := Text{
		ToUserName:   "to",
		FromUserName: "from",
		CreateTime:   strconv.Itoa(int(time.Now().Unix())),
		MsgType:      "tpe",
		Content:      "xyz",
	}
	result, err := xml.MarshalIndent(reply, "", "  ")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(result))

	var m map[string]interface{}
	e := xml.Unmarshal(result, &m)
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	fmt.Println(m)
}
