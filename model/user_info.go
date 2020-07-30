package model

// UserInfo is ...
type UserInfo struct {
	UserID    int    `json:"user_id"`
	LoginID   int    `json:"login_id"`
	UserName  string `json:"user_name"`
	Telephone string `json:"telephone"`
	Password  string `json:"password"`
	RoleID    int    `json:"role_id"`
}
