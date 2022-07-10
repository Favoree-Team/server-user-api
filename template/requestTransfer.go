package template

import "github.com/Favoree-Team/server-user-api/entity"

// template for email admin when user request transfer

var (
	requestTransferTemplateHTML = ``
)

func RequestTransferTemplate(transaction entity.Transaction) string {
	return requestTransferTemplateHTML
}
