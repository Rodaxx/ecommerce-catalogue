package repository

import (
	"context"

	"github.com/rodaxx/ecommerce-catalogue/models"
)
var implementation Repository

type Repository interface {

	// InsertUser(ctx context.Context, user *models.User) (int32,error)

	// GetUserById(ctx context.Context, id int32) (*models.User, error)
	// GetUserByEmail(ctx context.Context, email string) (*models.User, error)

	// UpdateUser(ctx context.Context, user *models.User) error

	// DeleteUserById(ctx context.Context,id string) error
	// DeleteUserByEmail(ctx context.Context, email string) error
	UpdateProductById(ctx context.Context,productId uint, newQuantity uint) (error)
	GetProductById(ctx context.Context,id uint) (*models.ProductDatabase, error)
	FindAllProducts(ctx context.Context) ([]* models.ProductDatabase,error)
	FindAllProductsForBranch(ctx context.Context,branchName string) ([]* models.ProductDatabase,error)
	Close() error
}

func FindAllProductsForBranch(ctx context.Context,branchName string) ([]* models.ProductDatabase,error) {
	return implementation.FindAllProductsForBranch(ctx,branchName)
}
func FindAllProducts(ctx context.Context) ([]* models.ProductDatabase,error) {
	return implementation.FindAllProducts(ctx)
}
func GetProductById(ctx context.Context,id uint) (*models.ProductDatabase, error){
	return implementation.GetProductById(ctx,id)
}
func UpdateProductById(ctx context.Context,productId uint, newQuantity uint) (error){
	return implementation.UpdateProductById(ctx,productId,newQuantity)
}
func Close() error {
	return implementation.Close()
}