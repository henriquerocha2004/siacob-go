package dto

type ActionOutput struct {
	Result  bool   `json:"result"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}
