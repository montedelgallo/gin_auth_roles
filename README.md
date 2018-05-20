# gin_auth_roles

This package is inspired on [Authz](https://github.com/gin-contrib/authz) and modified so to not support only basic HTTP authentication. it's based on [https://github.com/casbin/casbin](https://github.com/casbin/casbin).

## Installation

```
go get github.com/giovapanasiti/gin_auth_roles
```

## Simple Example

**To make it work you have to set `my_user_role` in the context when you authenticate your user:**
```go
func UpdateContextUserModel(c *gin.Context, my_user_id uint) {
	var myUserModel UserModel
	if my_user_id != 0 {
		db := common.GetDB()
		db.First(&myUserModel, my_user_id)
	}
	c.Set("my_user_id", my_user_id)
	c.Set("my_user_model", myUserModel)
	c.Set("my_user_role", myUserModel.Role)
}
```

```Go
package main

import (
	"net/http"

	"github.com/casbin/casbin"
	"github.com/giovapanasiti/gin_auth_roles"
	"github.com/gin-gonic/gin"
)

func main() {
	// load the casbin model and policy from files, database is also supported.
	e := casbin.NewEnforcer("auth.conf", "policy.csv")

	// define your router, and use the Casbin authz middleware.
	// the access that is denied by authz will return HTTP 403 error.
    router := gin.New()
    router.Use(gin_auth_roles.NewAuthorizer(e))
}
```

## Documentation

The authorization determines a request based on ``{subject, object, action}``, which means what ``subject`` can perform what ``action`` on what ``object``. In this plugin, the meanings are:

1. ``subject``: the logged-on user role
2. ``object``: the URL path for the web resource like "dataset1/item1"
3. ``action``: HTTP method like GET, POST, PUT, DELETE, or the high-level actions you defined like "read-file", "write-blog"


For how to write authorization policy and other details, please refer to [the Casbin's documentation](https://github.com/casbin/casbin).

## Getting Help

- [Casbin](https://github.com/casbin/casbin)

## License

This project is under MIT License. See the [LICENSE](LICENSE) file for the full license text.