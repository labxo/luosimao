package luosimao

import (
	"encoding/json"
	"net/url"
)

type VoiceClient Client

func NewVoiceClient(key string) *VoiceClient {
	return (*VoiceClient)(NewClient(key))
}

func (voice *VoiceClient) Verify(mobile string, code string) error {
	resp, err := voice.PostForm("https://voice-api.luosimao.com/v1/verify.json", url.Values{
		"mobile": {mobile},
		"code":   {code},
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

func (voice *VoiceClient) Status() (int, error) {

	resp, err := voice.Get("https://voice-api.luosimao.com/v1/status.json")

	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)

	data := &struct {
		Error
		Deposit int `json:",string"`
	}{}

	if err := decoder.Decode(data); err != nil {
		return 0, err
	}

	if err := data.Check(); err != nil {
		return 0, err
	}

	return data.Deposit, nil
}
