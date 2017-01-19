package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"

	"cigol.com/crawler/structs"
	"cigol.com/crawler/utils"

	htmlparser "github.com/calbucci/go-htmlparser"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/simplifiedchinese"
)

func trimInBetween(str string) string {
	if str == "" {
		return str
	}

	n := bytes.NewBufferString("")

	lastSpace := false

	for _, r := range str {
		if unicode.IsSpace(r) || unicode.IsControl(r) {
			if lastSpace {
				continue
			}
			lastSpace = true
			n.WriteRune(' ')
			continue
		}
		n.WriteRune(r)
		lastSpace = false
	}
	return n.String()
}

var config utils.Config = utils.GetConfig("config.ini")
var esUrl []string = getEsUrl()
var esClient utils.Client = utils.NewClient(esUrl)

func getEsUrl() []string {
	// get the es url
	v, ok := config["es.urls"]
	if ok {
		return strings.Split(v, ",")
	}
	return []string{"http://127.0.0.1:9200"}
}

func normalize(u string) string {
	if strings.HasSuffix(u, "/") {
		return u
	}
	i := strings.LastIndex(u, "/")
	if i >= 0 {
		s := u[i+1:]
		if !strings.Contains(s, ".") {
			return u + "/"
		}
	}
	return u
}

func getDomain(u string) string {
	i := strings.Index(u, "//")
	schema := u[:i]
	domain := u[i+2:]
	i = strings.Index(domain, "/")
	if i >= 0 {
		domain = domain[:i]
	}
	return schema + "//" + domain
}

// crawl page content from url u and parse the content
func crawl(pages *[]string, u string, crawled map[string]int, crawlingChan chan int, encodedStartUrl string, startUrl string, gb18030Decoder *encoding.Decoder, gbkDecoder *encoding.Decoder, maxSliceSize int) bool {
	_, contains := crawled[u]
	if contains {
		<-crawlingChan
		return true
	}
	crawled[u] = 1

	fmt.Println("crawling", u)
	client := utils.NewClient([]string{u})
	content, err := client.Get("", "")
	if err != nil {
		fmt.Println("request", u, "error", err)
		<-crawlingChan
		return true
	}
	//	fmt.Println(content)

	charset := "utf8"
	page := structs.Page{}
	page.Metas = map[string]string{}

	// get body, parser does not parse element innerHTML
	//		i := strings.Index(content, "<body>")
	//		if i >= 0 {
	//			page.Content = content[i+6:]
	//			i = strings.Index(page.Content, "</body>")
	//			if i >= 0 {
	//				page.Content = page.Content[:i]
	//			}
	//			page.Content = strings.TrimSpace(page.Content)
	//		}

	parser := htmlparser.NewParser(content)
	// extract all the a tags' href values
	parser.Parse(func(text string, e *htmlparser.HtmlElement) {
		if e == nil {
			return
		}
		tag := e.TagName
		if tag == "script" || tag == "style" {
			return
		}

		if tag == "title" {
			page.Title = trimInBetween(text)
		} else {
			inBody := false
			parent := e
			for {
				if parent == nil {
					break
				}
				if parent.TagName == "body" {
					inBody = true
					break
				}
				parent = parent.Parent
			}
			if inBody {
				page.Content = page.Content + "\n" + text
			}
		}
	}, func(e *htmlparser.HtmlElement, isEmpty bool) {
		if e != nil && e.TagName == "a" {
			href, _ := e.GetAttributeValue("href")
			//			fmt.Println(href)

			// only crawl pages under startUrl's domain
			if strings.HasPrefix(href, startUrl) {
				if len(*pages) < maxSliceSize {
					*pages = append(*pages, href)
				}
			} else {
				if strings.HasPrefix(href, "./") || strings.HasPrefix(href, "../") {
					if len(*pages) < maxSliceSize {
						*pages = append(*pages, u+href)
					}
				}
				if !strings.Contains(href, ":") {
					if strings.HasPrefix(href, "/") {
						if len(*pages) < maxSliceSize {
							*pages = append(*pages, getDomain(u)+href)
						}
					} else {
						if len(*pages) < maxSliceSize {
							*pages = append(*pages, u+href)
						}
					}
				}
			}

			//				pages = append(pages, href)
		}
		if e != nil && e.TagName == "meta" {
			name, _ := e.GetAttributeValue("name")
			metaContent, _ := e.GetAttributeValue("content")
			//				fmt.Println(name, metaContent)
			if name != "" {
				page.Metas[name] = metaContent
			}
			cs, ok := e.GetAttributeValue("charset")
			if ok {
				charset = cs
			}
			ct, ok := e.GetAttributeValue("http-equiv")
			if ok && strings.ToLower(ct) == "content-type" {
				c, ok := e.GetAttributeValue("content")
				if ok {
					i := strings.Index(strings.ToLower(c), "charset=")
					if i >= 0 {
						charset = c[i+8:]
					}
				}
			}
		}
	}, func(s string) {
		//	fmt.Println(s)
	})
	// trim spaces of the content
	page.Content = trimInBetween(strings.TrimSpace(page.Content))

	charset = strings.ToLower(charset)
	if charset == "gb18030" {
		page.Title, _ = gb18030Decoder.String(page.Title)
		page.Content, _ = gb18030Decoder.String(page.Content)
		for k, v := range page.Metas {
			page.Metas[k], _ = gb18030Decoder.String(v)
		}
	}
	if charset == "gbk" || charset == "gb2312" {
		page.Title, _ = gbkDecoder.String(page.Title)
		page.Content, _ = gbkDecoder.String(page.Content)
		for k, v := range page.Metas {
			page.Metas[k], _ = gbkDecoder.String(v)
		}
	}
	fmt.Println("crawled", u, "charset=", charset)

	b, err := json.Marshal(page)
	d := string(b)
	id := strings.Replace(u, "/", "%2F", -1)
	// fmt.Println(id)
	esClient.Put(config["es.index.prefix"]+encodedStartUrl, config["es.type"], id, d)
	// fmt.Println(d)

	<-crawlingChan
	return true
}

func main() {
	gb18030Decoder := simplifiedchinese.GB18030.NewDecoder()
	gbkDecoder := simplifiedchinese.GBK.NewDecoder()

	//	startUrl := "http://www.cnblogs.com/"
	startUrl := config["start.url"]
	args := os.Args
	if len(args) > 1 {
		startUrl = args[1]
	}
	// encode '/' to '%2F', and also encode '%' in '%2F' to '%25'
	// ES index name must be lower-cased, so change '%2F' to '%2f'
	// ES index name can not contains [\\, /, *, ?, \", <, >, |,  , ,]
	// also need to encode ':' to '%253a'
	encodedStartUrl := strings.Replace(startUrl, "/", "%252f", -1)
	encodedStartUrl = strings.Replace(encodedStartUrl, ":", "%253a", -1)

	maxPages, _ := strconv.Atoi(config["max.crawl.pages"])
	maxCrawlingPages, _ := strconv.Atoi(config["max.crawling.pages"])
	maxSliceSize, _ := strconv.Atoi(config["max.slice.size"])
	crawlingChan := make(chan int, maxCrawlingPages)
	waitChan := make(chan int, maxCrawlingPages)

	crawled := map[string]int{}
	pages := []string{startUrl}
	for i := 0; ; {
		//		fmt.Println(pages)
		if len(pages) > 0 {
			crawlingChan <- 0
			go func(i *int) {
				waitChan <- 0
				if len(pages) > 0 {
					u := pages[0]
					u = normalize(u)
					pages = pages[1:]
					b := crawl(&pages, u, crawled, crawlingChan, encodedStartUrl, startUrl, gb18030Decoder, gbkDecoder, maxSliceSize)
					if b {
						*i = *i + 1
					}
				}
				<-waitChan
			}(&i)
		}

		if i >= maxPages {
			break
		}
		time.Sleep(10e9)
	}
	for {
		select {
		case <-waitChan:
			continue
		default:
			break
		}
	}
}
