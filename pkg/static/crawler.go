package botserver

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/hackroid/tg-dumb-bot/pkg/static"
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

type WeiboCrawler struct {
	c              *colly.Collector
	topListChan    chan string // New matching items are put in this channel when they are crawled.
	crawlErrorChan chan error  // When an error occurs, the error will be put in this channel.
	finish         chan byte   // When the crawler ends gracefully the end byte will be put in this channel.
}

func GetCrawler() *WeiboCrawler {
	return &WeiboCrawler{}
}

func (b *WeiboCrawler) InitWeiboCrawler() {
	b.c = colly.NewCollector()
	// Set timeout 5s
	b.c.SetRequestTimeout(5 * time.Second)

	// Set delay between requests
	_ = b.c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 2,
		Delay:       5 * time.Second,
	})

	// When the crawler is finished it will call OnScraped()
	b.c.OnScraped(func(r *colly.Response) {
		b.finish <- 1
		log.Printf("The crawler ends gracefully.\n")
	})

	// When a crawler error occurs it will call OnError()
	b.c.OnError(func(r *colly.Response, err error) {
		// fmt.Println(err)
		b.crawlErrorChan <- err
		log.Printf("The crawler ends with an error. -> %s\n", err)
	})

	// When a matching item is found it will call OnHTML()
	b.c.OnHTML("td.td-02 a[target='_blank']", func(e *colly.HTMLElement) {
		b.topListChan <- e.Text
		log.Printf("Parsing data -> %s\n", e.Text)
	})

	// When a request is sent it will call OnRequest()
	b.c.OnRequest(func(r *colly.Request) {
		jsonFile, _ := os.Open("assets/web/sample_header.json")
		defer static.CloseFile(jsonFile)
		byteValue, _ := io.ReadAll(jsonFile)
		var headers map[string]string
		err := json.Unmarshal(byteValue, &headers)
		if err != nil {
			log.Fatalln("error parsing json")
		}
		for key, value := range headers {
			r.Headers.Set(key, value)
		}
	})
}

func (b *WeiboCrawler) startCrawlFenkeng() ([]string, error) {
	var topList []string
	var crawlError error

	// Before crawling task starts, initialize these channels
	b.topListChan = make(chan string)
	b.crawlErrorChan = make(chan error)
	b.finish = make(chan byte)

	// Random sleep in 1 second
	time.Sleep(time.Duration(rand.Intn(1e9)))

	go func() {
		err := b.c.Visit("https://s.weibo.com/top/summary?cate=realtimehot")
		if err != nil {

		}
	}()

	for {
		var breakFlag = false
		select {
		case top := <-b.topListChan:
			topList = append(topList, top)
		case crawlError = <-b.crawlErrorChan:
			breakFlag = true
		case <-b.finish:
			breakFlag = true
		default:
		}
		if breakFlag {
			break
		}
	}
	return topList, crawlError
}

func pack(msg string, i int) string {
	hyperlink := fmt.Sprintf("<a href=\"https://s.weibo.com/weibo?q=%s\">%s</a>\n", msg, msg)
	if i == 0 {
		hyperlink = "上升趋势：" + hyperlink
	} else {
		hyperlink = fmt.Sprintf("%d. ", i) + hyperlink
	}
	return hyperlink
}

func (b *WeiboCrawler) GetFenkengTrends(count int) string {
	topList, err := b.startCrawlFenkeng()
	var msg = "Get error"
	if len(topList) < count+1 {
		return msg
	}
	t := time.Now()
	currentTimeString := t.Format("2006-01-02 15:04:05 MST")
	if err == nil {
		msg = fmt.Sprintf("来自粪坑的top10热搜(%s)\n", currentTimeString)
		for i, top := range topList {
			msg += pack(top, i)
			if i == count {
				break
			}
		}
	}
	return msg
}
