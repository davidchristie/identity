package client

type User interface {
	Email() string
}

type user struct {
	email string
}

var _ User = (*user)(nil)

func (u *user) Email() string {
	return u.email
}
