package translate

func NewTranslation(word string) TranslatedWordRepository {
	data := Word{
		word,
	}
	repo := TranslatedWordRepository{translated: data}
	return repo
}

func (repo *TranslatedWordRepository) TranslatedWord() (Word){
	translated := repo.translated
	return translated
}
