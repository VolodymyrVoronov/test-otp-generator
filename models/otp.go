package models

type RequestBody struct {
	Secret string `json:"secret"`
}

type VerifyRequestBody struct {
	Secret string `json:"secret"`
	Code   string `json:"code"`
}
