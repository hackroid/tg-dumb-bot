package main

import (
	"log"
	"math/rand"
	"strings"
)

func help(msg string) (string, error) {
	return "按 \"/\" 自己看", nil
}

func choice(msg string) (string, error) {
	log.Printf("[CMD] %s", msg)
	dices := strings.FieldsFunc(msg, ff)
	if len(dices) == 1 {
		msg = "你选寄吧呢"
	} else if len(dices) == 2 {
		msg = "就一个你选寄吧呢"
	} else {
		dices = dices[1:]
		msg = dices[rand.Intn(len(dices))]
	}
	return msg, nil
}

func status(msg string) (string, error) {
	return "I'm 凹K.", nil
}

func ddefault(msg string) (string, error) {
	return "你说寄吧呢", nil
}
