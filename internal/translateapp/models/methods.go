package models

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
