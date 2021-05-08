package clients

type Account struct {
	Name       string `json:"name"`
	AccountUid string `json:"accountUid"`
}

type AccountInformation struct {
	Accounts []Account `json:"accounts"`
}
