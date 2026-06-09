package dto

type AuditLogResponse struct {
	ID uint `json:"id"`

	Action string `json:"action"`

	Resource string `json:"resource"`

	ResourceID string `json:"resource_id"`

	Details string `json:"details"`

	IPAddress string `json:"ip_address"`

	UserAgent string `json:"user_agent"`
}