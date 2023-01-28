package translate

type baiduTranslate struct{}

func newBaidu() Translate {
	return &baiduTranslate{}
}

func (g *baiduTranslate) Translate(text string, sourceLang, targetLang TranslateType) (string, error) {
	return "", nil
}