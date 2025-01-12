package user

import (
	uCOntr "github.com/Crud-application/pkg/contracts/user"
	uAgg "github.com/Crud-application/pkg/domain/userAgg"
)

func fromCreateUserReq(userID string, req *uCOntr.CreateUserReq) (*uAgg.User, error) {
	return uAgg.NewUser(
		userID,
		req.Name,
		req.Email,
		req.PhoneNumber,
	)
}
func toCreateUserRes(user *uAgg.User) *uCOntr.CreateUserRes {
	return &uCOntr.CreateUserRes{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
	}
}

func toGetUserRes(user *uAgg.User) *uCOntr.GetUserRes {
	userRes := &uCOntr.GetUserRes{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
	}
	return userRes
}

func toUpdateUserRes(user *uAgg.User) *uCOntr.UpdateUserRes {
	return &uCOntr.UpdateUserRes{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
	}
}
func toGetAllUsers(users []uAgg.User) []uCOntr.GetUserRes {
	var res []uCOntr.GetUserRes
	for _, user := range users {
		res = append(res, uCOntr.GetUserRes{
			ID:          user.ID,
			Name:        user.Name,
			Email:       user.Email,
			PhoneNumber: user.PhoneNumber,
		})
	}
	return res
}
