package mistakes

import (
	"fmt"
	json "github.com/json-iterator/go"
	"net/http"
	"net/url"
	"sync"
)

type Text struct {
	Texts []string `json:"texts"`
}

type sentence struct {
	pos int
	s   string
	e   error
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

func CorrectMistakes(text Text) *Text {
	texts := make([]string, len(text.Texts))
	c := make(chan sentence, len(text.Texts))
	wg1 := sync.WaitGroup{}
	wg1.Add(1)
	go func(c chan sentence, t []string, wg *sync.WaitGroup) {
		for s := range c {
			if s.e != nil {
				texts[s.pos] = text.Texts[s.pos] + "(Не исправлено)"
			} else {
				texts[s.pos] = s.s
			}
		}
		wg.Done()
	}(c, texts, &wg1)
	wg := sync.WaitGroup{}
	wg.Add(len(text.Texts))
	for i, t := range text.Texts {
		go sendSentence(i, t, c, &wg)
	}
	wg.Wait()
	close(c)
	wg1.Wait()
	resultText := &Text{
		Texts: texts,
	}
	return resultText
}

func sendSentence(i int, t string, c chan sentence, wg *sync.WaitGroup) {
	defer wg.Done()
	s, err := correctSentence(t)
	c <- sentence{pos: i, s: s, e: err}
}

func correctSentence(text string) (string, error) {
	data := url.Values{
		"text":    {text},
		"lang":    {"ru"},
		"format":  {"plain"},
		"options": {"524"},
	}
	resp, err := http.PostForm("https://speller.yandex.net/services/spellservice.json/checkText", data)
	if err != nil {
		return "", fmt.Errorf("error in request to speller. %v", err)
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}
	var res []SpellResult
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return "", fmt.Errorf("failed to decode speller result. %v", err)
	}
	result := replaceMistakes(res, text)
	return result, nil
}

func replaceMistakes(spellResult []SpellResult, text string) string {
	var res []rune
	p := 0
	source := []rune(text)
	for _, v := range spellResult {
		res = append(res, source[p:v.Pos]...)
		if len(v.S) > 0 {
			res = append(res, []rune(v.S[0])...)
		} else {
			res = append(res, source[v.Pos:v.Pos+v.Len]...)
		}
		p = v.Pos + v.Len
	}
	res = append(res, source[p:]...)

	return string(res)
}
