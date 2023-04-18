package static

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
)

type normalResponse struct {
	List []struct {
		Title string `json:"title"`
		URL   string `json:"url"`
	} `json:"list"`
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type tagResponse struct {
	List []struct {
		Tag           string `json:"tag"`
		TranslatedTag string `json:"translated_tag"`
	} `json:"list"`
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func request(req *http.Request) ([]byte, error) {
	var body []byte
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("%s", err)
	} else {
		body, err = io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("%s", err)
		}
		resp.Body.Close()
	}
	return body, err
}

func parseDataAndGenerateMsg(api string, body []byte) (string, error) {
	var err error
	var msg = "Parsing error!"
	if api != "getTrendingTags" {
		var res normalResponse
		log.Println("body -> %s", string(body))
		err = json.Unmarshal(body, &res)
		if err != nil {
			log.Printf("%s", err)
		} else {
			msg = res.List[rand.Intn(len(res.List))].URL
		}
	} else {
		var res tagResponse
		err = json.Unmarshal(body, &res)
		if err != nil {
			log.Printf("%s", err)
		} else {
			idx := rand.Intn(len(res.List))
			msg = fmt.Sprintf("你今日的幸运xp是%s (%s)!", res.List[idx].Tag, res.List[idx].TranslatedTag)
		}
	}
	return msg, err
}

func handleRequest(requestType string, key string, data string, api string) (string, error) {
	var req *http.Request
	var err error
	var body []byte
	var msg string
	requestUrl := "http://pixiv:5000/" + api
	formData := url.Values{}
	if data != "" { // POST mode
		formData.Set(key, data)
	}

	req, err = http.NewRequest(requestType, requestUrl, strings.NewReader(formData.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		log.Printf("%s", err)
	} else {
		body, err = request(req)
		if err != nil {
			log.Printf("%s", err)
		} else {
			msg, err = parseDataAndGenerateMsg(api, body)
		}
	}
	return msg, err
}

func getByPixivUid(uid string) (string, error) {
	return handleRequest("POST", "uid", uid, "getIllustListByUid")
}

func getByPixivRankingMode(mode string) (string, error) {
	return handleRequest("POST", "mode", mode, "getIllustRanking")
}

func getPixivTags() (string, error) {
	return handleRequest("GET", "", "", "getTrendingTags")
}

func getByillustId(illustId string) (string, error) {
	return handleRequest("POST", "illust_id", illustId, "getIllustDownloadUrl")
}
