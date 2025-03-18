package IntermediateServices

type Beat struct {
	API        string ` json:"api"`
	ReqMsgID   string `json:"reqMsgId"`
	Ts         string `json:"ts"`
	Status     string ` json:"status"`
	StatusDesc string `json:"statusDesc"`
	StatusCode string `json:"statusCode"`
	Type       string `json:"type"`
	Prodtype   string `json:"prodType"`
	BankCode   string `json:"bankCode"`
}
