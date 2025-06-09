package entities

type User struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"` // TODO: to be hashed
	Role      string `json:"role"`
}

func (User) TableName() string { //NOTE: ask more details about this
	return "users"
}
