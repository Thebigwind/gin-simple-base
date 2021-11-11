package application

import (
	"context"
	"fmt"
	"gin-simple-base/routers/typespec"
	"gin-simple-base/service"
	"time"
)

type User struct{}

func (v *User) GetUserList(ctx context.Context, req *typespec.GetUserListRequest, resp *typespec.GetUserListResponse) error {
	userSvc := &service.User{}
	if err := userSvc.GetMany(ctx, req, resp); err != nil {
		return err
	}

	return nil
}

func (v *User) AddUser(ctx context.Context, req *typespec.AddUserRequest, resp *typespec.AddUserResponse) error {
	userSvc := &service.User{}

	if err := userSvc.Add(ctx, req, resp); err != nil { //
		return err
	}

	return nil
}

func (v *User) GetUser(ctx context.Context, req *typespec.GetUserRequest, resp *typespec.GetUserResponse) error {
	go func() {
		fmt.Printf("xxxxxxxxx2\n")
		time.Sleep(time.Second * 10)
		fmt.Printf("xxxxxxxxx2\n")
	}()
	userSvc := &service.User{}
	if err := userSvc.Get(ctx, req, resp); err != nil {
		return err
	}

	return nil
}
