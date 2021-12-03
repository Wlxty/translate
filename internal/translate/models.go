package translate

type Word struct{
	TranslatedWord string
}

type TranslatedWordRepository struct {
	translated Word
}

type Language struct{
	Code string
	Name string
}

type LanguageRepository struct {
	languages []Language
}


