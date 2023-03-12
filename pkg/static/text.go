package static

import (
	"github.com/hackroid/tg-dumb-bot/pkg/datatype"
	"math/rand"
)

func AddGouBa(content datatype.TextContentRecv) (string, bool, error) {
	p := rand.Intn(100)
	if p < 15 {
		return content.Text + "个几把", true, nil
	}
	return "", false, nil
}
