package models

type Item struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price string `json:"price"`
}

type StatusInfo struct {
	Code    int    `json:"Code"`
	Error   int    `json:"Error"`
	Message string `json:"Message"`
}

type Response struct {
	Status  StatusInfo  `json:"Status"`
	Results interface{} `json:"Results"`
}
