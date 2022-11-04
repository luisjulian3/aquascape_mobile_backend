package models

type ResultPH struct {
	Time  string  `json:"time"`
	Value float64 `json:"value"`
}

type ResultFan struct {
	Status bool `json:"status"`
}
