package models

import "github.com/lcsin/go-study/step6/gin-prj-structrue/dao"

type User struct {
	ID        uint   `json:"id"`
	Nickname  string `json:"nickname"`
	Password  string `json:"password"`
	Age       uint   `json:"age"`
	Gender    string `json:"gender"`
	Email     string `json:"email"`
	Telephone string `json:"telephone"`
}

func ListUser() (users []*User, err error) {
	err = dao.DB.Select(&users, "select * from tb_user")
	if err != nil {
		return nil, err
	}
	return users, nil
}
