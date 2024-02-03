package handler

import (
	"MyEcommerce/features/user"
	"time"
)

type UserResponse struct {
	Name         string `json:"name" form:"name"`
	UserName     string `json:"user_name" form:"user_name"`
	Email        string `json:"email" form:"email"`
	PhotoProfile string `json:"photo_profile" form:"photo_profile"`
	Role         string `json:"role" form:"role"`
}

type AdminUserResponse struct {
	ID           uint       `json:"id" form:"id"`
	Name         string     `json:"name" form:"name"`
	UserName     string     `json:"user_name" form:"user_name"`
	Email        string     `json:"email" form:"email"`
	Role         string     `json:"role" form:"role"`
	PhotoProfile string     `json:"photo_profile" form:"photo_profile"`
	CreatedAt    time.Time  `json:"created_at" form:"created_at"`
	DeletedAt    *time.Time `json:"deleted_at" form:"deleted_at"`
	StatusUser   string     `json:"status_user" form:"status_user"`
}

type UserProductResponse struct {
	Name         string `json:"name" form:"name"`
	UserName     string `json:"user_name" form:"user_name"`
	PhotoProfile string `json:"photo_profile" form:"photo_profile"`
}

type CartUserResponse struct {
	Name string `json:"name" form:"name"`
}

func CoreToResponse(data *user.Core) UserResponse {
	var result = UserResponse{
		Name:         data.Name,
		UserName:     data.UserName,
		Email:        data.Email,
		PhotoProfile: data.PhotoProfile,
		Role:         data.Role,
	}
	return result
}

func CoreToResponseList(data []user.Core) []AdminUserResponse {
	var results []AdminUserResponse
	for _, v := range data {
		var statusUser string
		if v.DeletedAt == nil {
			statusUser = "Active"
		} else {
			statusUser = "Not Active"
		}
		var result = AdminUserResponse{
			ID:           v.ID,
			Name:         v.Name,
			UserName:     v.UserName,
			Email:        v.Email,
			Role:         v.Role,
			PhotoProfile: v.PhotoProfile,
			CreatedAt:    v.CreatedAt,
			DeletedAt:    v.DeletedAt,
			StatusUser:   statusUser,
		}
		results = append(results, result)
	}
	return results
}
