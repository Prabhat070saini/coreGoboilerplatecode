package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/example/testing/shared/constants"
	"github.com/example/testing/shared/constants/exception"
	"github.com/gin-gonic/gin"
)

type permissionMiddleware struct {
	access *MiddlewareAccess
}

type PermissionMiddlewareMethods interface {
	PermissionMiddleware(requiredRoles ...constants.Roles) gin.HandlerFunc
}

func NewPermissionMiddleware(access *MiddlewareAccess) PermissionMiddlewareMethods {
	return &permissionMiddleware{
		access: access,
	}
}

func (m *permissionMiddleware) PermissionMiddleware(requiredRoles ...constants.Roles) gin.HandlerFunc {
	return func(c *gin.Context) {
		rolesVal, exists := c.Get("roles")
		if !exists {
			SendErrorResponse(c, http.StatusUnauthorized, exception.GetException(exception.PROTECTED_ROUTE).Message, http.StatusUnauthorized)
			return
		}

		fmt.Println("Debug: rolesVal =", rolesVal)
		userRoles, ok := rolesVal.([]string)
		if !ok {
			SendErrorResponse(c, http.StatusInternalServerError, "Incorrect formate", http.StatusInternalServerError)
			return
		}

		fmt.Println(userRoles, "user roles")
		// Convert userRoles to map
		roleSet := make(map[string]struct{}, len(userRoles))
		for _, r := range userRoles {
			roleSet[strings.ToUpper(r)] = struct{}{}
		}

		// Check if any requiredRole is in user's roleSet
		for _, required := range requiredRoles {
			if _, found := roleSet[strings.ToUpper(string(required))]; found {
				c.Next()
				return
			}
		}

		SendErrorResponse(c, http.StatusUnauthorized, exception.GetException(exception.PROTECTED_ROUTE).Message, http.StatusUnauthorized)
		c.Abort()
	}
}
