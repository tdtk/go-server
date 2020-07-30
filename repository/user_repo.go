package repository

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/tdtk/go-server/model"
)

// UserRepository is ...
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository is ...
func NewUserRepository() *UserRepository {
	db, err := sql.Open("mysql", "root:password@tcp(mysql:3306)/user")

	if err != nil {
		panic(err.Error())
	}

	return &UserRepository{db: db}
}

// FindAllUser is ...
func (repo *UserRepository) FindAllUser() []model.UserInfo {
	results, err := repo.db.Query("select * from user_info")

	if err != nil {
		panic(err.Error())
	}

	var users []model.UserInfo

	for results.Next() {
		var user model.UserInfo
		err = results.Scan(&user.UserID, &user.LoginID, &user.UserName, &user.Telephone, &user.RoleID)

		if err != nil {
			panic(err.Error())
		}

		users = append(users, user)
	}

	return users
}

// Close is ...
func (repo *UserRepository) Close() {
	defer repo.db.Close()
}
