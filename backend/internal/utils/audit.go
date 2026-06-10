package utils



import "github.com/gin-gonic/gin"


type AuditContext struct {
	UserID    *uint
	IPAddress string
	UserAgent string
}




func NewAuditContext(
	c *gin.Context,
) AuditContext {

	var userID *uint

	if id := c.GetUint("user_id"); id != 0 {
		userID = &id
	}

	return AuditContext{
		UserID:    userID,
		IPAddress: c.ClientIP(),
		UserAgent: c.Request.UserAgent(),
	}
}