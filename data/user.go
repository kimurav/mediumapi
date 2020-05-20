package data

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email" db:"email"`
	Username string `json:"username" db:"username"`
	Bio      string `json:"bio"	 db:"bio"`
	Password string `json:"password" db:"password"`
}
