package repository

import (
	"database/sql"
	"fmt"

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

// SearchUser is ...
func (repo *UserRepository) SearchUser(params model.SearchFormParams) []model.UserInfo {

	var query string

	var selectAll = "select user_id, login_id, user_name, telephone, password, role_id" +
		"from user_info" +
		"where is_deleted = 0"

	if params.UserName != "" && params.Telephone != "" {
		query = fmt.Sprintf(
			selectAll+
				"and user_name like '%%%s%%'"+
				"and telephone like '%%%s%%'",
			params.UserName,
			params.Telephone,
		)
	} else if params.UserName != "" {
		query = fmt.Sprintf(
			selectAll+
				"and user_name like '%%%s%%'",
			params.UserName,
		)
	} else if params.Telephone != "" {
		query = fmt.Sprintf(
			selectAll+
				"and telephone like '%%%s%%'",
			params.Telephone,
		)
	} else {
		query = selectAll
	}

	results, err := repo.db.Query(query)

	if err != nil {
		panic(err.Error())
	}

	var users []model.UserInfo

	for results.Next() {
		var user model.UserInfo
		err = results.Scan(&user.UserID, &user.LoginID, &user.UserName, &user.Telephone, &user.Password, &user.RoleID)

		if err != nil {
			panic(err.Error())
		}

		users = append(users, user)
	}

	return users
}

// GetUserByID is ...
func (repo *UserRepository) GetUserByID(userID int) model.UserInfo {
	results, err := repo.db.Query(fmt.Sprintf("select user_id, login_id, user_name, telephone, password, role_id from user_info where user_id=%d", userID))

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var user model.UserInfo
		err = results.Scan(&user.UserID, &user.LoginID, &user.UserName, &user.Telephone, &user.Password, &user.RoleID)

		if err != nil {
			panic(err.Error())
		}

		return user
	}
	panic("Can't get user!")
}

// GetPasswordByID is ...
func (repo *UserRepository) GetPasswordByID(loginID string) string {
	results, err := repo.db.Query(fmt.Sprintf("select password from user_info where login_id='%s' and is_deleted = 0", loginID))

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var pass string
		err = results.Scan(&pass)

		if err != nil {
			panic(err.Error())
		}

		return pass
	}
	panic("Can't get password!")
}

// UpdateUser is ...
func (repo *UserRepository) UpdateUser(user model.UserInfo) {
	_, err := repo.db.Query(
		fmt.Sprintf(
			"update user_info"+
				"set login_id = %s"+
				"set user_name = '%s'"+
				"set telephone = '%s'"+
				"set password = '%s'"+
				"set role_id = %d"+
				"where id = %d",
			user.LoginID,
			user.UserName,
			user.Telephone,
			user.Password,
			user.RoleID,
			user.UserID,
		),
	)
	if err != nil {
		panic(err.Error())
	}
}

// DeleteUser is ...
func (repo *UserRepository) DeleteUser(userID int) {
	_, err := repo.db.Query(fmt.Sprintf("update user_info set is_deleted = 1 where userID=%d", userID))
	if err != nil {
		panic(err.Error())
	}
}

// GetRoleByID is ...
func (repo *UserRepository) GetRoleByID(roleID int) model.Role {
	results, err := repo.db.Query(fmt.Sprintf("select * from role where role_id='%d'", roleID))

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var role model.Role
		err = results.Scan(&role.RoleID, &role.RoleName)

		if err != nil {
			panic(err.Error())
		}

		return role
	}
	panic("Can't get role!")
}

// GetAllRole is ...
func (repo *UserRepository) GetAllRole() []model.Role {
	results, err := repo.db.Query("select * from role")

	if err != nil {
		panic(err.Error())
	}

	var roles []model.Role

	for results.Next() {
		var role model.Role
		err = results.Scan(&role.RoleID, &role.RoleName)

		if err != nil {
			panic(err.Error())
		}
		roles = append(roles, role)
	}

	return roles
}

// Close is ...
func (repo *UserRepository) Close() {
	defer repo.db.Close()
}
