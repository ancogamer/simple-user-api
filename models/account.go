package models

type AccountReq struct {
	User     string `json:"user"`
	Password string `json:"pass"`
}
