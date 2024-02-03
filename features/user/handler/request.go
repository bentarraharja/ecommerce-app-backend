package handler

import "MyEcommerce/features/user"

type UserRequest struct {
	Name         string `json:"name" form:"name"`
	UserName     string `json:"user_name" form:"user_name"`
	Email        string `json:"email" form:"email"`
	Password     string `json:"password" form:"password"`
	Role         string `json:"role" form:"role"`
	PhotoProfile string `json:"photo_profile" form:"photo_profile"`
}

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func RequestToCore(input UserRequest) user.Core {
	return user.Core{
		Name:         input.Name,
		UserName:     input.UserName,
		Email:        input.Email,
		Password:     input.Password,
		Role:         input.Role,
		PhotoProfile: input.PhotoProfile,
	}
}

func UpdateRequestToCore(input UserRequest, imageURL string) user.Core {
	return user.Core{
		Name:         input.Name,
		UserName:     input.UserName,
		Email:        input.Email,
		Password:     input.Password,
		Role:         input.Role,
		PhotoProfile: imageURL,
	}
}
