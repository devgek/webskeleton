package request

import (
	"net/http"

	"github.com/stretchr/objx"
	"kahrersoftware.at/webskeleton/types"
)

//RData the data hold in request context
type RData interface {
	UserID() string
	SetUserID(userID string)
	Role() types.RoleType
	SetRole(role types.RoleType)
	IsAdmin() bool
}
type requestData struct {
	userID     string
	role       types.RoleType
	customerID uint
}

//NewRequestData create RequestData
func NewRequestData() RData {
	return &requestData{"", types.RoleTypeUser, 0}
}

func (c requestData) UserID() string {
	return c.userID
}

func (c requestData) Role() types.RoleType {
	return c.role
}

func (c *requestData) SetUserID(userID string) {
	c.userID = userID
}

func (c *requestData) SetRole(role types.RoleType) {
	c.role = role
}

func (c requestData) IsAdmin() bool {
	return c.role == types.RoleTypeAdmin
}

func (c requestData) MSI() map[string]interface{} {
	ctxMap := objx.New(map[string]interface{}{
		"userID": c.UserID(),
		"role":   c.Role(),
	})
	return objx.New(map[string]interface{}{
		"request-data": ctxMap,
	})
}

//ContextKeyRequestData ...
var ContextKeyRequestData = "request-data"

//FromCookie get RequestData from cookie value
func FromCookie(cookie *http.Cookie) (RData, bool) {
	c := objx.MustFromBase64(cookie.Value)

	cData := NewRequestData()
	val := c.Get("cookie-data.request-data.userID")
	if val != nil {
		cData.SetUserID(val.Str())
		val = c.Get("cookie-data.request-data.role")
		if val != nil {
			cData.SetRole(types.RoleType(val.Int()))
			return cData, true
		}
	}

	return &requestData{}, false
}
