package translateapp

type Word struct {
	TranslatedWord string
}

type Language struct {
	Code string
	Name string
}

type Translate interface {
	Languages() []Language
}

func (l *Language) Languages() []Language {
	languages := []Language{
		Language{"en", "English"},
		Language{"pl", "Polish"},
	}
	return languages
}

func NewWord(word string) Word {
	data := Word{
		word,
	}
	return data
}
