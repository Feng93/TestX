package main

import (
	"encoding/xml"
	"fmt"
	"os"
)


type Servers struct {
	XMLName xml.Name `xml:"servers"`
	Version string   `xml:"version,attr"`
	Svs     []server `xml:"server"`
}


type server struct {
	ServerName string `xml:"serverName"`
	ServerIP   string `xml:"serverIP,CDATA"`
}


func main() {
	var a = make([]int,0)
	fmt.Println(a[:0])

	v := &Servers{Version: "1"}
	v.Svs = append(v.Svs, server{"Shanghai_VPN", "127.0.0.1"})
	v.Svs = append(v.Svs, server{"Beijing_VPN", "![CDATA[127.0.0.1]]"})
	output, err := xml.MarshalIndent(v, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	os.Stdout.Write([]byte(xml.Header))


	os.Stdout.Write(output)
}
