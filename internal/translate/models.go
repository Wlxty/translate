package translate

type Word struct{
	TranslatedWord string
}

type TranslatedRepository interface {
	Translate() (Word)
}
type Language struct{
	Code string
	Name string
}

type LanguageRepository interface {
	Languages() ([]Language)
}


