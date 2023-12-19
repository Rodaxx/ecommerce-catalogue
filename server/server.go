package server

import (
	"context"
	"errors"

	"github.com/rodaxx/ecommerce-catalogue/models"
	"github.com/rodaxx/ecommerce-catalogue/repository"
)

type Server struct {
	repo repository.Repository
}

func NewCatalogueServer(repo repository.Repository) *Server {
	return &Server{repo: repo}
}

func (s *Server) GetProducts(ctx context.Context) ([]*models.ProductDatabase, error) {
    products, err := s.repo.FindAllProducts(ctx)
    
	if err != nil {
        return nil, err
    }
    return products, nil
}
func (s *Server) GetProductsForBranch(ctx context.Context,branchName string) ([]*models.ProductDatabase, error) {
    products, err := s.repo.FindAllProductsForBranch(ctx,branchName)
    
	if err != nil {
        return nil, err
    }
    return products, nil
} 
func (s *Server) GetProductById(ctx context.Context,id uint) (*models.ProductDatabase, error) {

	product,err := s.repo.GetProductById(ctx,id);

	if err != nil {
		return nil, err
	}

    if product.Id==0{
        return nil, errors.New("Product Doesn't exists")
    }
	return &models.ProductDatabase{
		Id: product.Id,
		Title: product.Title,
		Description: product.Description,
		Price: product.Price,
		Size: product.Size,
		Type: product.Size,
		Color: product.Color,
		Gender: product.Gender,
		ImageOne: product.ImageOne,
		ImageTwo: product.ImageTwo,
		BranchName: product.BranchName,
		BranchDirection: product.BranchDirection,
		Quantity: product.Quantity,
	}, nil
}
func(s *Server) UpdateProductById(ctx context.Context,productId uint,newQuantity uint) (error){
	product, err := s.repo.GetProductById(ctx, productId)
    if err != nil {
        return errors.New("Error al obtener el product por ID:")
    }
	if product.Id==0{
		return errors.New("The product Doesn't exists")
	}

    err = s.repo.UpdateProductById(ctx,productId,newQuantity)
    if err != nil {
        return errors.New("Error al actualizar el product")
    }
	return nil
}