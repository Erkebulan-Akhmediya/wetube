package service

import (
	"fmt"
	"github.com/lib/pq"
	"wetube/database"
)

const baseSelect = `
select u.id,
       u.username,
       u.password,
       u.created_at,
       u.deleted_at,
       coalesce(array_agg(r.name) filter (where r.name is not null), ARRAY[]::varchar(10)[]) roles
from "user" u
         left join users_roles ur on u.id = ur.user_id
         left join role r on ur.role_name = r.name
%s
group by u.id`

func GetById(id int) (*User, error) {
	query := fmt.Sprintf(baseSelect, "where id = $1")
	return get(query, id)
}

func GetByUsername(username string) (*User, error) {
	query := fmt.Sprintf(baseSelect, " where username = $1")
	return get(query, username)
}

func get(query string, args ...any) (*User, error) {
	var user User
	row := database.Db().QueryRow(query, args...)
	err := row.Scan(&user.Id, &user.Username, &user.Password, &user.CreatedAt, &user.DeletedAt, pq.Array(&user.Roles))
	if err != nil {
		return nil, err
	}
	return &user, nil
}
