package models

import "github.com/golang-jwt/jwt"

type ResultPH struct {
	Time  string  `json:"time"`
	Value float64 `json:"value"`
}

type ResultTemp struct {
	Time  string  `json:"time"`
	Value float64 `json:"value"`
}

type ResultFan struct {
	Status bool `json:"status"`
}

type ResultLamp struct {
	Status bool `json:"status"`
}

type Profile struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

type JwtClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

type NewResultDataPH struct {
	//Time  string `json:"time"`
	Value string `json:"value"`
}
type NewResultDataTemp struct {
	//Time  string `json:"time"`
	Value string `json:"value"`
}
