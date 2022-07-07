package entity

const (
	AccessibleStatus = "ACCESS"
	RejectedStatus   = "REJECT"
)

type IPRecord struct {
	ID        string `json:"id"`
	IPAddress string `json:"ip_address"`
	Role      string `json:"role"`
}

type IPRecordResponse struct {
	IPAddress string `json:"ip_address"`
	Status    string `json:"status"`
	Message   string `json:"message"`
}

type IPRecordInput struct {
	IPAddress string `json:"ip_address" binding:"required"`
	Role      string `json:"role" binding:"required"`
}

type IPRecordRequest struct {
	IPAddress string `json:"ip_address" binding:"required"`
}
