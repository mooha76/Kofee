package model

type Request struct {
	SchemaVersion string        `json:"schemaVersion"`
	RequestId     string        `json:"requestId"`
	Timestamp     string        `json:"timestamp"`
	ChannelName   string        `json:"channelName"`
	ServiceName   string        `json:"serviceName" binding:"required"`
	ServiceParams ServiceParams `json:"serviceParams"`
}

type ServiceParams struct {
	MerchantUid     string          `json:"merchantUid" binding:"required"`
	PaymentMethod   string          `json:"paymentMethod" binding:"required"`
	ApiKey          string          `json:"apiKey" binding:"required"`
	ApiUserId       string          `json:"apiUserId" binding:"required"`
	PayerInfo       PayerInfo       `json:"payerInfo"`
	TransactionInfo TransactionInfo `json:"transactionInfo"`
}

type PayerInfo struct {
	AccountNo string `json:"accountNo" binding:"required"`
}

type TransactionInfo struct {
	ReferenceId string `json:"referenceId"`
	InvoiceId   string `json:"invoiceId"`
	Amount      string `json:"amount" binding:"required"`
	Currency    string `json:"currency" binding:"required"`
	Description string `json:"description"`
}
