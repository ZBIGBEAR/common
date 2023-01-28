package translate

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

type googleTranslate struct {
	url string
}

func newGoogle() Translate {
	return &googleTranslate{
		url: "https://translate.googleapis.com/translate_a/single?client=gtx",
	}
}

func (g *googleTranslate) Translate(source string, sourceLang, targetLang TranslateType) (string, error) {
	var text []string
	var result []interface{}

	encodedSource, err := encodeURI(source)
	if err != nil {
		return "err", err
	}

	url := fmt.Sprintf("%s&sl=%s&tl=%s&dt=t&q=%s", g.url, sourceLang, targetLang, encodedSource)

	body, err := httpGet(url)
	if err != nil {
		return "", errors.Wrapf(err, "googleTranslate Translate. source:%s, sourceLang:%s, targetLang:%s", source, sourceLang, targetLang)
	}

	bReq := strings.Contains(string(body), `<title>Error 400 (Bad Request)`)
	if bReq {
		return "err", errors.New("Error 400 (Bad Request)")
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return "err", errors.New("Error unmarshaling data")
	}

	if len(result) == 0 {
		return "", errors.New("No translated data in responce")

	}

	inner := result[0]
	for _, slice := range inner.([]interface{}) {
		for _, translatedText := range slice.([]interface{}) {
			text = append(text, fmt.Sprintf("%v", translatedText))
			break
		}
	}
	
	cText := strings.Join(text, "")

	return cText, nil
}
