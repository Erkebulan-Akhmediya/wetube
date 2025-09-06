package role

import (
	"net/http"
	userService "wetube/users/service"
)

func NewRoleMiddleware(roles []string, next http.Handler) http.Handler {
	rolesMap := make(map[string]bool)
	for _, role := range roles {
		rolesMap[role] = true
	}
	return &roleMiddleware{
		roles: rolesMap,
		next:  next,
	}
}

type roleMiddleware struct {
	roles map[string]bool
	next  http.Handler
}

func (rh *roleMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("user").(userService.User)
	if !ok {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	for _, role := range user.Roles {
		if rh.roles[role] {
			rh.next.ServeHTTP(w, r)
			return
		}
	}
	w.WriteHeader(http.StatusForbidden)
}
