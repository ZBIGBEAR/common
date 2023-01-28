package translate

type Translate interface {
	Translate(text string, sourceLang, targetLang TranslateType) (string, error)
}

func New() Translate {
	return newGoogle()
}

func NewBaidu() Translate {
	return newBaidu()
}

func NewGoogle() Translate {
	return newGoogle()
}