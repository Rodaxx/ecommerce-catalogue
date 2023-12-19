package database

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rodaxx/ecommerce-catalogue/models"
)

type MySQLRepository struct {
	db *sql.DB
}

// Crear nueva instancia del repositorio MySQL
func NewMySQLRepository(url string) (*MySQLRepository, error) {
	db, err := sql.Open("mysql", url)
	if err != nil {
		return nil, err
	}

	return &MySQLRepository{db}, nil
}
func (repo *MySQLRepository) FindAllProductsForBranch(ctx context.Context,branchName string) ([]* models.ProductDatabase,error){
	rows, err := repo.db.QueryContext(ctx, 
		"SELECT id,title,description,price,size,type,color,gender,imageOne,imageTwo,branch_name,branch_direction,quantity FROM products WHERE branch_name LIKE ? ",branchName)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()	
	var products []*models.ProductDatabase
	for rows.Next() {
		var product = models.ProductDatabase{}
		if err = rows.Scan(&product.Id,&product.Title,&product.Description,&product.Price,
							&product.Size,&product.Type,&product.Color,&product.Gender,
							&product.ImageOne,&product.ImageTwo,
							&product.BranchName,&product.BranchDirection,
							&product.Quantity); err == nil {
			products = append(products, &product)
		}
	}
	if err = rows.Err(); err != nil {	
		return nil, err
	}

	return products, nil
}
func (repo *MySQLRepository) FindAllProducts(ctx context.Context) ([]*models.ProductDatabase, error) {
	rows, err := repo.db.QueryContext(ctx, 
		"SELECT id,title,description,price,size,type,color,gender,imageOne,imageTwo,branch_name,branch_direction,quantity FROM products ORDER BY branch_name,title")
	if err != nil {
		return nil, err
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()	
	var products []*models.ProductDatabase
	for rows.Next() {
		var product = models.ProductDatabase{}
		if err = rows.Scan(&product.Id,&product.Title,&product.Description,&product.Price,
							&product.Size,&product.Type,&product.Color,&product.Gender,
							&product.ImageOne,&product.ImageTwo,
							&product.BranchName,&product.BranchDirection,
							&product.Quantity); err == nil {
			products = append(products, &product)
		}
	}
	if err = rows.Err(); err != nil {	
		return nil, err
	}

	return products, nil
}
func (repo *MySQLRepository) UpdateProductById(ctx context.Context,productId uint, newQuantity uint) (error){
		_, err := repo.db.ExecContext(ctx, "UPDATE products SET quantity = ? WHERE id = ?",
		newQuantity,
		productId)
	return err
}
func (repo *MySQLRepository) GetProductById(ctx context.Context, id uint) (*models.ProductDatabase, error) {
	rows, err := repo.db.QueryContext(ctx,
		"SELECT id, title, description, price, size, type, color, gender, imageOne, imageTwo, branch_name, branch_direction, quantity FROM products WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	defer func() {
		closeErr := rows.Close()
		if closeErr != nil {
			log.Println("Error closing rows:", closeErr)
		}
	}()
	if err != nil {
		return nil, err
	}
	var product = models.ProductDatabase{}
	for rows.Next() {
		if err = rows.Scan(&product.Id,&product.Title,&product.Description,&product.Price,
							&product.Size,&product.Type,&product.Color,&product.Gender,
							&product.ImageOne,&product.ImageTwo,
							&product.BranchName,&product.BranchDirection,
							&product.Quantity); err == nil {
			return &product, nil
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	
	return &product, nil
}
func (repo *MySQLRepository) Close() error {
	return repo.db.Close()
}
