package products

import (
	"context"
	"errors"
	"log"

	"github.com/Gigi-U/PRODUCTS-CRUD-GO.git/internal/models"
)

var (
	ErrEmpty    = errors.New("empty list")
	ErrNotFound = errors.New("product not found")
)

type Repository interface {
	Create(ctx context.Context, producto models.Producto) (models.Producto, error)
	GetAll(ctx context.Context) ([]models.Producto, error)
	GetByID(ctx context.Context, id string) (models.Producto, error)
	Update(ctx context.Context, producto models.Producto, id string) (models.Producto, error)
	Delete(ctx context.Context, id string) error
}
// esta estructura es la base del repository
type repository struct {
	// implemento una base de datos en memoria que va a ser un slice de productos
	db []models.Producto
}

// NewMemoryRepository ....
func NewMemoryRepository(db []models.Producto) Repository {
	return &repository{db: db}
}

// Create ....
func (r *repository) Create(ctx context.Context, producto models.Producto) (models.Producto, error) {
	r.db = append(r.db, producto)
	return producto, nil
}

// GetAll...
func (r *repository) GetAll(ctx context.Context) ([]models.Producto, error) {

	contenidoContext := ctx.Value("rol")

	if contenidoContext != "" {
		log.Println("El contenido del contexto es:", contenidoContext)
	}

	if len(r.db) < 1 {
		return []models.Producto{}, ErrEmpty
	}

	return r.db, nil
}

// GetByID .....
func (r *repository) GetByID(ctx context.Context, id string) (models.Producto, error) {
	var result models.Producto
	for _, value := range r.db {
		if value.Id == id {
			result = value
			break
		}
	}

	if result.Id == "" {
		return models.Producto{}, ErrNotFound
	}

	return result, nil
}

// Update ...
func (r *repository) Update(
	ctx context.Context,
	producto models.Producto,
	id string) (models.Producto, error) {

	var result models.Producto
	for key, value := range r.db {
		if value.Id == id {
			producto.Id = id
			r.db[key] = producto
			result = r.db[key]
			break
		}
	}

	if result.Id == "" {
		return models.Producto{}, ErrNotFound
	}

	return result, nil

}

// Delete ...
func (r *repository) Delete(ctx context.Context, id string) error {
	var result models.Producto
	for key, value := range r.db {
		if value.Id == id {
			result = r.db[key]
			r.db = append(r.db[:key], r.db[key+1:]...)
			break
		}
	}

	if result.Id == "" {
		return ErrNotFound
	}

	return nil
}