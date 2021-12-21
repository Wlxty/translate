package translateapp

type Response struct {
	TranslatedWord string
}

type Word struct {
	TranslatedWord string
}

type Language struct {
	Code string
	Name string
	// Stats int // For the future.
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

func (repo Word) Translate() Word {
	src := repo
	return src
}
