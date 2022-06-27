package mistakes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"net/http"
)

type Text struct {
	Text string `json:"text"`
}

type SpellResult struct {
	SpellError `json:"error"`
	Word string `json:"word"`
	S []string `json:"s"`
}

type SpellError struct {
	Code int `json:"code"`
	Pos int `json:"pos"`
	Row int `json:"row"`
	Col int `json:"col"`
	Len int `json:"len"`
}

func CorrectMistakes(text Text, l *zap.Logger) (string,error) {
	buffer := bytes.Buffer{}
	err := json.NewEncoder(&buffer).Encode(text)
	if err != nil {
		return "", nil
	}
	req, err := http.NewRequest("POST", "https://speller.yandex.net/services/spellservice.json/checkText", &buffer)
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	l.Info(fmt.Sprintf("sss %v", resp))
	if resp.Body != nil {
		defer resp.Body.Close()
	}
	var res []SpellResult
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return "", err
	}
	l.Info(fmt.Sprintf("123 %v", res))
	return "", nil
}
