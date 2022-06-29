package mistakes

import (
	json "github.com/json-iterator/go"
	"go.uber.org/zap"
	"net/http"
	"net/url"
	"strings"
)

type Text struct {
	Text string `json:"text"`
}

type SpellResult struct {
	Code int      `json:"code"`
	Pos  int      `json:"pos"`
	Row  int      `json:"row"`
	Col  int      `json:"col"`
	Len  int      `json:"len"`
	Word string   `json:"word"`
	S    []string `json:"s"`
}

func CorrectMistakes(text string, l *zap.Logger) (string, error) {
	data := url.Values{
		"text":    {text},
		"lang":    {"ru"},
		"format":  {"plain"},
		"options": {"524"},
	}
	resp, err := http.PostForm("https://speller.yandex.net/services/spellservice.json/checkText", data)
	if err != nil {
		return "", err
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}
	var res []SpellResult
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return "", err
	}
	result := replaceMistakes(res, text)
	return result, nil
}

func replaceMistakes(spellResult []SpellResult, text string) string {
	words := make(map[int]string)
	wordsCount := len(strings.Split(text, " "))
	pos := make([]int, 0, wordsCount)
	i := 0
	for _, w := range strings.Split(text, " ") {
		words[i] = w
		pos = append(pos, i)
		i += len([]rune(w)) + 1
	}
	for _, v := range spellResult {
		if len(v.S) > 0 {
			words[v.Pos] = v.S[0]
		}
	}
	result := ""
	for i, v := range pos {
		result += words[v]
		if i < wordsCount-1 {
			result += " "
		}
	}
	return result
}
