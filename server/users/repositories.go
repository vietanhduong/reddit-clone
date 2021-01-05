package users

import "reddit-clone/server/common"

type (
	Repository struct {
		// Implement DATABASE in here
		// In this project I using seed data
		users map[string]*User
	}
)

func NewRepository() *Repository {
	admin := initAdmin()
	return &Repository{users: map[string]*User{admin.Username: admin}}
}

func (r *Repository) Login(username string, password string) *User {
	user, found := r.users[username]
	if !found || user.Password != password {
		return nil
	}
	return user
}

func (r *Repository) FindByUsername(username string) *User {
	user, found := r.users[username]
	if !found {
		return nil
	}
	return user
}

func initAdmin() *User {
	return &User{
		Username: common.GetEnv("ADMIN_USERNAME", "vietanhduong"),
		Password: common.GetEnv("ADMIN_PASSWORD", "661d1d4879d8da594125c86ad61ae055"),
		FullName: "ADMIN",
		Admin:    true,
	}
}
