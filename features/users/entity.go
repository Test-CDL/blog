package users

import "time"

type Core struct {
	ID        int
	UserName  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Business interface {
	InsertUser(data Core) (row int, err error)
	LoginUser(userName string, password string) (data string, token string, e error)
}

type Data interface {
	PostUser(data Core) (row int, err error)
	AuthUser(userName string, password string) (data string, token string, e error)
}
