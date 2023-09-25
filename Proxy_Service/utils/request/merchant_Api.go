package request

type MERCHANT_API_PURCHASE_REQ struct {
	ServiceName string `json:"serviceName" binding:"required"`
	MerchantUid string `json:"merchantUid" binding:"required"`

	ApiKey    string `json:"apiKey" binding:"required"`
	ApiUserId string `json:"apiUserId" binding:"required"`

	AccountNo string `json:"accountNo" binding:"required"`
	Amount    string `json:"amount" binding:"required"`
}
