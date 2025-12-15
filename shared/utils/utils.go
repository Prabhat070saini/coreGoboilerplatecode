package utils

import (
	"context"
	// "database/sql"
	"fmt"
	"net/http"

	"github.com/example/testing/shared/constants/exception"
	"github.com/example/testing/shared/response"
	"gorm.io/gorm"
)

type txKey struct{}

func WithTx(ctx context.Context, tx *gorm.DB) context.Context {
	return context.WithValue(ctx, txKey{}, tx)
}

func GetTx(ctx context.Context) *gorm.DB {
	tx, _ := ctx.Value(txKey{}).(*gorm.DB)
	return tx
}

type TxFunc[T any] func(ctx context.Context, tx *gorm.DB) response.ServiceOutput[T]

func WithTransaction[T any](db *gorm.DB, ctx context.Context, fn TxFunc[T]) response.ServiceOutput[T] {
	// you also pass isolation table in Begin(&sql.TxOptions{Isolation: sql.LevelSerializable})
	tx := db.Begin()
	if tx.Error != nil {
		return response.ServiceOutput[T]{
			Exception: &response.Exception{
				Message:        tx.Error.Error(),
				HttpStatusCode: http.StatusInternalServerError,
				Code:           http.StatusInternalServerError,
			},
		}
	}

	fmt.Println("🚀 Transaction started")

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			fmt.Println("❌ Transaction rolled back due to panic:", r)
		}
	}()

	// Inject transaction into context
	ctxWithTx := WithTx(ctx, tx)

	// Execute the callback
	output := fn(ctxWithTx, tx)

	// Rollback if exception occurred
	if output.Exception != nil {
		if err := tx.Rollback().Error; err != nil {
			fmt.Println("❌ Rollback failed:", err)
		} else {
			fmt.Println("⚠️ Transaction rolled back due to exception")
		}
		return output
	}

	// Commit if success
	if err := tx.Commit().Error; err != nil {
		fmt.Println("❌ Transaction commit failed:", err)
		return response.ServiceOutput[T]{
			Exception: &response.Exception{
				Message:        err.Error(),
				HttpStatusCode: http.StatusInternalServerError,
				Code:           http.StatusInternalServerError,
			},
		}
	}

	fmt.Println("✅ Transaction committed successfully")
	return output
}

func HandleException[T any](exception response.Exception) response.ServiceOutput[T] {
	return response.ServiceOutput[T]{
		Exception: &exception,
	}
}
func ServiceError[T any](code exception.ErrorCode) response.ServiceOutput[T] {
	ex := exception.GetException(code)
	return response.ServiceOutput[T]{
		Exception: ex,
	}
}
