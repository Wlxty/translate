package languagesstore

import(
	"errors"
)

type LanguagesStore struct {
	languages map[string]string
}

func New() *LanguagesStore {
	languages := map[string]string{
		"af":    "Afrikaans",
		"sq":    "Albanian",
		"am":    "Amharic",
		"ar":    "Arabic",
		"hy":    "Armenian",
		"az":    "Azerbaijani",
		"eu":    "Basque",
		"be":    "Belarusian",
		"bn":    "Bengali",
		"bs":    "Bosnian",
		"bg":    "Bulgarian",
		"ca":    "Catalan",
		"ceb":   "Cebuano",
		"ny":    "Chichewa",
		"zh-CN": "Chinese",
		"co":    "Corsican",
		"hr":    "Croatian",
		"cs":    "Czech",
		"da":    "Danish",
		"nl":    "Dutch",
		"en":    "English",
		"eo":    "Esperanto",
		"et":    "Estonian",
		"tl":    "Filipino",
		"fi":    "Finnish",
		"fr":    "French",
		"fy":    "Frisian",
		"gl":    "Galician",
		"ka":    "Georgian",
		"de":    "German",
		"el":    "Greek",
		"gu":    "Gujarati",
		"ht":    "Haitian Creole",
		"ha":    "Hausa",
		"haw":   "Hawaiian",
		"iw":    "Hebrew",
		"hi":    "Hindi",
		"hmn":   "Hmong",
		"hu":    "Hungarian",
		"is":    "Icelandic",
		"ig":    "Igbo",
		"class": "Indonesian",
		"ga":    "Irish",
		"it":    "Italian",
		"ja":    "Japanese",
		"jw":    "Javanese",
		"kn":    "Kannada",
		"kk":    "Kazakh",
		"km":    "Khmer",
		"rw":    "Kinyarwanda",
		"ko":    "Korean",
		"ku":    "Kurdish (Kurmanji)",
		"ky":    "Kyrgyz",
		"lo":    "Lao",
		"la":    "Latin",
		"lv":    "Latvian",
		"lt":    "Lithuanian",
		"lb":    "Luxembourgish",
		"mk":    "Macedonian",
		"mg":    "Malagasy",
		"ms":    "Malay",
		"ml":    "Malayalam",
		"mt":    "Maltese",
		"mi":    "Maori",
		"mr":    "Marathi",
		"mn":    "Mongolian",
		"my":    "Myanmar (Burmese)",
		"ne":    "Nepali",
		"no":    "Norwegian",
		"or":    "Odia (Oriya)",
		"ps":    "Pashto",
		"fa":    "Persian",
		"pl":    "Polish",
		"pt":    "Portuguese",
		"pa":    "Punjabi",
		"ro":    "Romanian",
		"ru":    "Russian",
		"sm":    "Samoan",
		"gd":    "Scots Gaelic",
		"sr":    "Serbian",
		"st":    "Sesotho",
		"sn":    "Shona",
		"sd":    "Sindhi",
		"si":    "Sinhala",
		"sk":    "Slovak",
		"sl":    "Slovenian",
		"so":    "Somali",
		"es":    "Spanish",
		"su":    "Sundanese",
		"sw":    "Swahili",
		"sv":    "Swedish",
		"tg":    "Tajik",
		"ta":    "Tamil",
		"tt":    "Tatar",
		"te":    "Telugu",
		"th":    "Thai",
		"tr":    "Turkish",
		"tk":    "Turkmen",
		"uk":    "Ukrainian",
		"ur":    "Urdu",
		"ug":    "Uyghur",
		"uz":    "Uzbek",
		"vi":    "Vietnamese",
		"cy":    "Welsh",
		"xh":    "Xhosa",
		"yi":    "Yclassdish",
		"yo":    "Yoruba",
		"zu":    "Zulu",
	}
	languagesStore := LanguagesStore{languages: languages}
	return &languagesStore
}

func (ts *LanguagesStore) GetLanguage(key string) (string){
	value, exists := ts.languages[key]
	if exists {
		return value
	}
	return ""
}

func (ts *LanguagesStore) DeleteLanguage(key string) error{
	languages := *&ts.languages
	if _, ok := languages[key]; !ok {
		return errors.New("Such Language does not exist")
	}
	delete(languages, key)
	return nil
}