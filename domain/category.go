package domain

// Category represents a product category entity
type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// CategoryRepository defines the interface for category data operations
type CategoryRepository interface {
	GetAll() ([]Category, error)
	GetByID(id int) (*Category, error)
	Create(category *Category) error
	Update(id int, category *Category) error
	Delete(id int) error
}
