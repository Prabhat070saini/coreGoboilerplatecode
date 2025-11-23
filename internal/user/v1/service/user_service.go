package userService

import (
	"context"

	"github.com/example/testing/shared/response"
	"github.com/example/testing/shared/utils"
	"github.com/example/testing/internal/user/models"
	userRepository "github.com/example/testing/internal/user/repository"
	"gorm.io/gorm"
)

type UserServiceMethods interface {
	GetUser(ctx context.Context, conditions map[string]interface{}, selectFields ...string) response.FunctionOutput[*models.Users]
	GetUserEmail(ctx context.Context, email string, selectFields ...string) response.ServiceOutput[*models.Users]
}

type userService struct {
	repo   userRepository.UserRepositoryMethods
	access *UserServiceAccess
}

func NewUserService(repo userRepository.UserRepositoryMethods, access *UserServiceAccess) UserServiceMethods {
	return &userService{
		repo:   repo,
		access: access,
	}
}

func (s *userService) GetUser(ctx context.Context, conditions map[string]interface{}, selectFields ...string) response.FunctionOutput[*models.Users] {
	if len(selectFields) == 0 {
		selectFields = []string{"uuid", "password_hash", "is_blocked", "email", "name"}
	}

	output := s.repo.FindUserByFields(ctx, conditions, selectFields...)
	if output.Exception != nil {
		return response.FunctionOutput[*models.Users]{Exception: output.Exception}
	}

	// go func (){
	//  s.access.Email.Send("prabhat.saini@binmile.com","Testing SMTP","<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n<meta charset=\"UTF-8\">\n<meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\n<title>Welcome</title>\n<style>\nbody, html {margin:0; padding:0; height:100%; font-family: Arial, sans-serif;}\nbody {display:flex; justify-content:center; align-items:center; background:linear-gradient(135deg, #6a11cb, #2575fc); color:#fff; text-align:center;}\n.welcome-container {background: rgba(0,0,0,0.5); padding:50px 80px; border-radius:15px; box-shadow:0 8px 20px rgba(0,0,0,0.3);}\n.welcome-container h1 {font-size:3em; margin-bottom:20px;}\n.welcome-container p {font-size:1.2em; margin-bottom:30px;}\n.welcome-container a {display:inline-block; padding:12px 25px; background:#ff6b6b; color:#fff; text-decoration:none; font-weight:bold; border-radius:8px; transition: background 0.3s ease;}\n.welcome-container a:hover {background:#ff4757;}\n</style>\n</head>\n<body>\n<div class=\"welcome-container\">\n<h1>Welcome!</h1>\n<p>We're glad to have you here. Explore and enjoy your experience.</p>\n<a href=\"#\">Get Started</a>\n</div>\n</body>\n</html>")
	// }()
	return response.FunctionOutput[*models.Users]{Data: output.Data}
}

func (s *userService) GetUserEmail(ctx context.Context, email string, selectFields ...string) response.ServiceOutput[*models.Users] {
	return utils.WithTransaction(s.access.TransactionDb, ctx, func(ctx context.Context, tx *gorm.DB) response.ServiceOutput[*models.Users] {
		if len(selectFields) == 0 {
			selectFields = []string{"uuid", "password_hash", "is_blocked", "email", "name"}
		}

		conditions := map[string]interface{}{"email": email}
		output := s.repo.FindUserByFields(ctx, conditions, selectFields...)

		if output.Exception != nil {
			return utils.HandleException[*models.Users](*output.Exception)
		}

		return response.ServiceOutput[*models.Users]{
			Success: &response.Success[*models.Users]{
				Code:           200,
				Message:        "User fetched successfully",
				HttpStatusCode: 200,
				Data:           output.Data,
			},
		}
	})
}
