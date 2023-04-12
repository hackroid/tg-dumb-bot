package static

import (
	"github.com/hackroid/tg-dumb-bot/pkg/datatype"
	"log"
	"math/rand"
	"strings"
	"unicode"
)

var fenkeng *WeiboCrawler

func ParseCommandMessage(content datatype.CommandContentRecv) (string, string, bool, error) {
	splitter := func(r rune) bool { return unicode.IsSpace(r) }
	parseMode := ""
	msg := ""
	var err error
	switch content.Cmd {
	case "help":
		msg = "按 \"/\" 自己看"
	case "choice":
		dices := strings.FieldsFunc(content.Text, splitter)
		if len(dices) == 1 {
			msg = "你选寄吧呢"
		} else if len(dices) == 2 {
			msg = "就一个你选寄吧呢"
		} else {
			dices = dices[1:]
			msg = dices[rand.Intn(len(dices))]
		}
	case "status":
		msg = "I'm 凹K."
	case "weibo":
		fenkeng = GetCrawler()
		fenkeng.InitWeiboCrawler()
		msg, err = fenkeng.GetFenkengTrends(10)
		if err != nil {
			log.Printf("Err: %v\n", err)
		}
		parseMode = "HTML"
	case "pranking":
		dices := strings.FieldsFunc(content.Text, splitter)
		if len(dices) == 1 {
			msg = "plz give me a mode type!(mode: day, week, month, day_male, day_female, week_original, week_rookie, day_manga)"
		} else {
			msg, err = getByPixivRankingMode(dices[1])
		}
	case "ptags":
		msg, err = getPixivTags()
	case "puid":
		dices := strings.FieldsFunc(content.Text, splitter)
		if len(dices) == 1 {
			msg = "plz give me an uid!"
		} else {
			msg, err = getByPixivUid(dices[1])
		}
	default:
		return msg, parseMode, false, err
	}
	return msg, parseMode, true, err
}
