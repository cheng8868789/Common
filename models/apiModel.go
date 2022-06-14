package models

type ApiModel struct {
	Head RequestHead
	Body RequestBody
}

type RequestHead struct {
	Attributes map[string]string
}

type RequestBody struct {
	Data string
}