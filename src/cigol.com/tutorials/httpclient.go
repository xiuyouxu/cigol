package main

import (
	"encoding/xml"
	"fmt"
	_ "io"
	"mathbeta"
	_ "net/http"
	_ "os"
)

type head struct {
	XMLName xml.Name `xml:"head"`
	Title   string   `xml:"title"`
	Metas   []string `xml:"meta"`
	Styles  []string `xml:"style"`
}

type body struct {
	XMLName xml.Name `xml:"body"`
	Content string   `xml:",innerxml"`
}

type page struct {
	XMLName xml.Name `xml:"html"`
	Head    head     `xml:"head"`
	Body    body     `xml:"body"`
}

func main() {
	var url string = "http://127.0.0.1"

	body, err := mathbeta.Crawl(url)
	//	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(body)
	//	defer response.Body.Close()

	//	//	io.Copy(os.Stdout, response.Body)
	//	body := ""
	//	b := make([]byte, 32*1024)
	//	for {
	//		n, err := response.Body.Read(b)
	//		if n > 0 {
	//			body = body + string(b[:n])
	//			//	fmt.Println(string(b[:n]))
	//		}
	//		if err != nil || err == io.EOF {
	//			break
	//		}
	//	}

	//	fmt.Println(body)

	// need to define the html page struct
	//	var p page
	//	e := xml.Unmarshal([]byte(body), &p)
	//	if e != nil {
	//		fmt.Println(e)
	//		return
	//	}
	//	fmt.Println(p.Head.Title)
}
