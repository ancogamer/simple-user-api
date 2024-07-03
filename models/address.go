package models

type Address struct {
	UserID  string
	ID      string
	Details string // dont have any idea for better name
	ZipCode string
	Country string
	State   string
	City    string
}

type AddressReq struct {
	UserID  string `json:"user_id" example:"some uuid"`
	ZipCode string `json:"zip_code" example:"15900-000"`
	Details string `json:"details"` // dont have any idea for better name
	State   string `json:"state" example:"São Paulo"`
	Country string `json:"country" example:"Brasil"`
	City    string `json:"city" example:"São Carlos"`
}

type AddressResp struct {
	UserID  string `json:"user_id"`
	ID      string `json:"id"`
	Details string `json:"details"`
	ZipCode string `json:"zip_code"`
	Country string `json:"country"`
	State   string `json:"state"`
	City    string `json:"city"`
}
