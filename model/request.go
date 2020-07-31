package model

// LoginFormParams is ...
type LoginFormParams struct {
	LoginID  string `json:"login_id"`
	Password string `json:"password"`
}

// SearchFormParams is ...
type SearchFormParams struct {
	UserName  string `json:"user_name"`
	Telephone string `json:"telephone"`
}
