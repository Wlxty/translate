package models

type Word struct{
	TranslatedWord string
}

type TranslatedWordRepository struct {
	translated Word
}

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
