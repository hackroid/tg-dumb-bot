package static

import (
	"github.com/hackroid/tg-dumb-bot/pkg/datatype"
	"math/rand"
	"strings"
	"unicode"
)

func NormalCommandMessage(content datatype.CommandContentRecv) (string, bool, error) {
	splitter := func(r rune) bool { return unicode.IsSpace(r) }
	var msg string
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
	default:
		return "", false, nil
	}
	return msg, true, nil
}
