package telegram

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"time"

	goservice "github.com/baozhenglab/go-sdk/v2"
)

type telegramService struct {
	token    string
	username string
}

var client = http.Client{
	Timeout: 5 * time.Second,
}

const (
	grapURL    = "https://api.telegram.org/bot"
	KeyService = "telegram-bot"
)

func (telegram *telegramService) Name() string {
	return KeyService
}

func (telegram *telegramService) GetPrefix() string {
	return KeyService
}

func (telegram *telegramService) GetUserName() string {
	return telegram.username
}

func (telegram *telegramService) InitFlags() {
	prefix := fmt.Sprintf("%s-", telegram.Name())
	flag.StringVar(&telegram.token, prefix+"token", "", "Token of telegram bot")
	flag.StringVar(&telegram.username, prefix+"username", "", "Username of telegram bot")
}

func (telegram *telegramService) Get() interface{} {
	return telegram
}

func NewTelegramBot() goservice.PrefixConfigure {
	return new(telegramService)
}

func (telegram *telegramService) SendMessage(form map[string]string) error {
	endPointRequest := grapURL + telegram.token + "/sendMessage"
	jsonValue, err := json.Marshal(form)
	if err != nil {
		return nil
	}
	req, err := http.NewRequest("POST", endPointRequest, bytes.NewBuffer(jsonValue))
	if err != nil {
		return err
	}
	response, err := client.Do(req)
	if err != nil {
		return err
	}
	var parse map[string]interface{}
	json.NewDecoder(response.Body).Decode(&parse)
	if parse["ok"] != true {
		return errors.New(parse["description"].(string))
	}
	return nil
}
