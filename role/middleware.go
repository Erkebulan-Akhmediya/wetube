package role

import "net/http"

func RBAC(next http.HandlerFunc, rolesForMethods map[string][]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		next(w, r)
	}
}
