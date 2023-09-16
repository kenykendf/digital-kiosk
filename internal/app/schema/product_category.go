package schema

type GetProductCategoriesLists struct {
	ID          uint   `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type CreateProductCategory struct {
	Name        string `json:"name" validate:"required,min=3,max=100"`
	Description string `json:"description" validate:"required"`
}

type UpdateProductCategory struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
