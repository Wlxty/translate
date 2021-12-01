package translated

type Word struct{
	TranslatedWord string
}

type Repository struct {
	translated Word
}

func New(word string) Repository {
	data := Word{
		word,
	}
	repo := Repository{translated: data}
	return repo
}

func (repo *Repository) TranslatedWord() (Word){
	translated := repo.translated
	return translated
}
