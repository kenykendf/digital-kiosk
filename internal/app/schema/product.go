package schema

type QueryParams struct {
	Limit    int
	Offset   int
	Name     string
	Category string
	SortBy   string
	AscDesc  int // value 0 desc
}

type GetProductsLists struct {
	ID                  uint   `json:"id"`
	Name                string `json:"name"`
	Description         string `json:"description"`
	Currency            string `json:"currency"`
	Price               uint64 `json:"price"`
	Sold                uint64 `json:"sold"`
	Views               uint64 `json:"views"`
	ProductCategoryID   uint   `json:"product_category_id"`
	ProductCategoryName string `json:"product_category_name"`
	Quantity            uint64 `json:"quantity"`
}

type CreateProduct struct {
	Name              string `validate:"required" form:"name"`
	Description       string `validate:"required" form:"description"`
	ProductCategoryID uint   `validate:"required,number" form:"product_category_id"`
	Currency          string `validate:"required,max=3,min=3" form:"currency"`
	Price             uint64 `validate:"required,number" form:"price"`
	Quantity          uint64 `validate:"required,number" form:"quantity"`
}

type UpdateProduct struct {
	Name              string `form:"name"`
	Description       string `form:"description"`
	ProductCategoryID uint   `form:"product_category_id"`
	Currency          string `form:"currency"`
	Price             uint64 `form:"price"`
	Quantity          uint64 `form:"quantity"`
}

type UpdateProductSell struct {
	Quantity uint64 `form:"quantity"`
}
