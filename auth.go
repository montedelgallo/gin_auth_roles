package gin_auth_roles

import (
"net/http"
"github.com/casbin/casbin"
"github.com/gin-gonic/gin"
)

// NewAuthorizer returns the authorizer, uses a Casbin enforcer as input
func NewAuthorizer(e *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		a := &BasicAuthorizer{enforcer: e}

		if !a.CheckPermission(c.Request, c ) {
			a.RequirePermission(c.Writer, c)
		}
	}
}

// BasicAuthorizer stores the casbin handler
type BasicAuthorizer struct {
	enforcer *casbin.Enforcer
}

// GetRole gets the user role from the context.
// so you are supposed to set "my_user_role" in the context
func (b *BasicAuthorizer) GetRole(c *gin.Context) interface{} {
	user, _ := c.Get("my_user_role")
	u := user.(string)
	return u
}

// CheckPermission checks the user/method/path combination from the request.
// Returns true (permission granted) or false (permission forbidden)
func (a *BasicAuthorizer) CheckPermission(r *http.Request, c *gin.Context) bool {
	user := a.GetRole(c)
	//user := "cigno"
	method := r.Method
	path := r.URL.Path
	return a.enforcer.Enforce(user, path, method)
}

// RequirePermission returns the 403 Forbidden to the client
// and close the response
func (a *BasicAuthorizer) RequirePermission(w http.ResponseWriter, c *gin.Context) {
	w.WriteHeader(403)
	w.Write([]byte("403 Forbidden\n"))
	c.Abort()
}
