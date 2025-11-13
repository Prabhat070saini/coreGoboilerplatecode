package userRepository

import (
	"context"

	"github.com/example/testing/common/constants/exception"
	"github.com/example/testing/common/lib/logger"
	"github.com/example/testing/common/response"
	"github.com/example/testing/common/utils"
	"github.com/example/testing/internal/user/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type userRepository struct {
	access *UserRepositoryAccess
}

type UserRepositoryMethods interface {
	// FindUserByFields(ctx context.Context, conditions map[string]interface{}, selectFields ...string) response.FunctionOutput[*models.Users]
	FindUserByFields(ctx context.Context, conditions map[string]interface{}, selectFields ...string) response.FunctionOutput[*models.Users]
}

func NewUserRepository(access *UserRepositoryAccess) UserRepositoryMethods {
	return &userRepository{access: access}
}

func (ul *userRepository) FindUserByFields(ctx context.Context, conditions map[string]interface{}, selectFields ...string) response.FunctionOutput[*models.Users] {
	var user *models.Users

	db := utils.GetTx(ctx)
	if db == nil {
		db = ul.access.DB
	}

	query := db.WithContext(ctx).Model(&models.Users{})
	if len(selectFields) > 0 {
		query = query.Select(selectFields)
	}

	query = query.Where(conditions)

	err := query.First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return response.FunctionOutput[*models.Users]{Data: nil}
		}
		logger.Error(ctx, "while finding user", zap.Error(err))
		return response.FunctionOutput[*models.Users]{Exception: exception.GetException(exception.INTERNAL_SERVER_ERROR)}
	}

	return response.FunctionOutput[*models.Users]{Data: user}
}
