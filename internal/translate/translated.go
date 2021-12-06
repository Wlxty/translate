package translate

func NewWord(word string) Word {
	data := Word{
		word,
	}
	return data
}

func (repo Word) Translate() (Word){
	src := repo
	return src
}

