package service

import (
	"context"
	"fmt"
	"gin-simple-base/models"
	tool "gin-simple-base/pkg/tools/time"
	"gin-simple-base/routers/typespec"
)

type User struct{}

// GetMany returns patch list
func (u *User) GetMany(ctx context.Context, req *typespec.GetUserListRequest, resp *typespec.GetUserListResponse) error {
	var (
		userList []models.User
		user     models.User
		count    int64
		err      error
	)

	user.Name = req.Name
	user.Age = req.Age
	user.NickName = req.NicName
	user.Sex = req.Sex
	user.Phone = req.Phone

	userList, count, err = user.GetUserListByCondition(req.Offset, req.Length)
	if err != nil {
		return err
	}

	if len(userList) > 0 {
		for _, v := range userList {
			var res typespec.GetUserResponse
			res.Id = v.ID
			res.Name = v.Name
			res.NicName = v.NickName
			res.Phone = v.Phone
			res.Age = v.Age
			res.CreateOn = tool.GormTimeFormat(v.CreatedOn)
			res.ModifyOn = tool.GormTimeFormat(v.ModifiedOn)
			resp.List = append(resp.List, res)
		}
	} else {
		resp.List = make([]typespec.GetUserResponse, 0)
	}

	resp.Offset = req.Offset
	resp.Length = req.Length
	resp.Total = count

	return nil
}

// Get returns a single patch data
func (u *User) Get(ctx context.Context, req *typespec.GetUserRequest, resp *typespec.GetUserResponse) error {
	var user models.User

	user.ID = req.Id
	fmt.Printf("id:", req.Id)
	user, err := user.GetUser()
	if err != nil {
		return err
	}

	resp.Id = user.ID
	resp.NicName = user.NickName
	resp.Phone = user.Phone
	resp.Name = user.Name
	resp.Age = user.Age
	resp.Sex = user.Sex
	resp.CreateOn = tool.GormTimeFormat(user.CreatedOn)
	resp.ModifyOn = tool.GormTimeFormat(user.ModifiedOn)

	return err
}

func (u *User) Add(ctx context.Context, req *typespec.AddUserRequest, resp *typespec.AddUserResponse) error {

	var user models.User

	user.Name = req.Name
	user.NickName = req.NicName
	user.Age = req.Age
	user.Phone = req.Phone
	user.Sex = req.Sex

	err := user.AddUser()
	resp.Id = user.ID
	return err

}
