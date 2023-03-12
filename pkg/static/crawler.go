package static

import (
	"github.com/gocolly/colly/v2"
	"log"
	"time"
)

type WeiboCrawler struct {
	c              *colly.Collector
	topListChan    chan string // New matching items are put in this channel when they are crawled.
	crawlErrorChan chan error  // When an error occurs, the error will be put in this channel.
	finish         chan byte   // When the crawler ends gracefully the end byte will be put in this channel.
}

func New() *WeiboCrawler {
	return &WeiboCrawler{}
}

func (b *WeiboCrawler) InitWeiboCrawler() {
	b.c = colly.NewCollector()
	// set timeout 10s
	b.c.SetRequestTimeout(10 * time.Second)

	// When the crawler is finished it will call OnScraped()
	b.c.OnScraped(func(r *colly.Response) {
		b.finish <- 1
		log.Printf("The crawler ends gracefully.\n")
	})

	// When a crawler error occurs it will call OnError()
	b.c.OnError(func(r *colly.Response, err error) {
		//fmt.Println(err)
		b.crawlErrorChan <- err
		log.Printf("The crawler ends with an error. -> %s\n", err)
	})

	// When a matching item is found it will call OnHTML()
	b.c.OnHTML("td.td-02 a[target='_blank']", func(e *colly.HTMLElement) {
		b.topListChan <- e.Text
		log.Printf("Parsing data -> %s\n", e.Text)
	})

	headers := map[string]string{
		"user-agent":      "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.83 Safari/537.36",
		"authority":       "s.weibo.com",
		"accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
		"accept-language": "zh-CN,zh;q=0.9,en;q=0.8,ja;q=0.7",
		"cookie":          "UOR=www.google.com,open.weibo.com,www.google.com; SINAGLOBAL=2970299018191.4736.1678542597936; SUB=_2AkMTUAgef8NxqwJRmPkRyWrlZI5wzgHEieKlDPnFJRMxHRl-yT9kqhc5tRB6ONAm8UidViM1ATfHJXRQ_RuEAa6LFw5c; SUBP=0033WrSXqPxfM72-Ws9jqgMF55529P9D9W5dNrgevn7ahTl16nJeYxmX; _s_tentry=-; Apache=858229940677.7616.1678584988845; ULV=1678584988859:2:2:1:858229940677.7616.1678584988845:1678542598026",
	}
	// When a request is sent it will call OnRequest()
	b.c.OnRequest(func(r *colly.Request) {
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

	go b.c.Visit("https://s.weibo.com/top/summary?cate=realtimehot")

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

func pack(msg string) string {
	msg += "\nhttps://s.weibo.com/weibo?q=" + msg + "\n"
	return msg
}

func (b *WeiboCrawler) GetFenkengTrends(count int) string {
	toplist, err := b.startCrawlFenkeng()
	currentTime := time.Now()
	var msg = "Get error"
	var index = 0
	if err == nil {
		msg = currentTime.String() + "\n来自粪坑的top10热搜\n"
		msg += "当前呈上升趋势的热搜词条\n" + pack(toplist[0])
		for _, top := range toplist[1:] {
			msg += pack(top)
			index++
			if count == index {
				break
			}
		}
	}
	return msg
}
