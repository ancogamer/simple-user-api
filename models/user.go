package models

type UserReq struct {
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Age      int    `json:"Age"`
}

type User struct {
	AccountID string
	ID        string
	Name      string
	LastName  *string
	Age       int
}

type UserResp struct {
	ID       string        `json:"id"`
	Name     string        `json:"name"`
	LastName *string       `json:"lastName"`
	Age      int           `json:"age"`
	Address  []AddressResp `json:"address"`
}
