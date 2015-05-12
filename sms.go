package luosimao

import (
	"encoding/json"
	"net/url"
)

type SmsClient Client

func NewSmsClient(key string) *SmsClient {
	return (*SmsClient)(NewClient(key))
}

func (sms *SmsClient) Send(mobile string, message string) error {

	resp, err := sms.PostForm("https://sms-api.luosimao.com/v1/send.json", url.Values{
		"mobile":  {mobile},
		"message": {message},
	})

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)

	data := &Error{}

	if err := decoder.Decode(data); err != nil {
		return err
	}
	return data.Check()
}

func (sms *SmsClient) Status() (int, error) {

	resp, err := sms.Get("https://sms-api.luosimao.com/v1/status.json")
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()
	data := &struct {
		Error
		Deposit int `json:",string"`
	}{}

	decoder := json.NewDecoder(resp.Body)

	if err := decoder.Decode(data); err != nil {
		return 0, err
	}

	if err := data.Check(); err != nil {
		return 0, err
	} else {
		return data.Deposit, nil
	}
}
