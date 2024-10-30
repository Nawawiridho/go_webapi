package respon

type StatusInfo struct {
	Code    int    `json:"Code"`
	Error   int    `json:"Error"`
	Message string `json:"Message"`
}

type Response struct {
	Status  StatusInfo  `json:"Status"`
	Results interface{} `json:"Results"`
}
