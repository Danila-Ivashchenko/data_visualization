package domain

type Dog struct {
	Name          string `json:"Порода собак" csv:"Порода собак"`
	MinPrice      int    `json:"Стоимость мин" csv:"Стоимость мин"`
	MaxPrice      int    `json:"Стоимость макс" csv:"Стоимость макс"`
	MinWeight     int    `json:"Вес мин" csv:"Вес мин"`
	MaxWeight     int    `json:"Вес макс" csv:"Вес макс"`
	MinHeight     int    `json:"Рост мин" csv:"Рост мин"`
	MaxHeight     int    `json:"Рост макс" csv:"Рост макс"`
	MinLivingTime int    `json:"Продолжительность жизни мин" csv:"Продолжительность жизни мин"`
	MaxLivingTime int    `json:"Продолжительность жизни макс" csv:"Продолжительность жизни макс"`
	Rating        int    `json:"Рейтинг" csv:"Рейтинг"`
	AntiRaiting   int    `json:"Антирейтинг" csv:"Антирейтинг"`
}
