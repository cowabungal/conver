package conver

import "strconv"

type User struct {
	Username string `db:"username"`
	UserId int `db:"user_id"`
	Names []string `db:"names"`
}

func (u *User) Recipient() string {
	return strconv.Itoa(u.UserId)
}
