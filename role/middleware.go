package role

import (
	"net/http"
	userService "wetube/users/service"
)

func RBAC(rolesForMethods map[string][]string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		allowedRoles := rolesForMethods[r.Method]
		userRoles := r.Context().Value("user").(userService.User).Roles
		userRolesSet := make(map[string]bool)
		for _, role := range userRoles {
			userRolesSet[role] = true
		}

		var allowed bool
		for _, role := range allowedRoles {
			if userRolesSet[role] {
				allowed = true
				break
			}
		}

		if !allowed {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		next(w, r)
	}
}
