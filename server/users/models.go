package users

type (
	User struct {
		Username string `json:"username"`
		Password string `json:"-"`
		FullName string `json:"full_name"`
		Admin    bool   `json:"admin"`
	}

	UserRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
)
