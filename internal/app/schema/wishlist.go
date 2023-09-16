package schema

type GetWishlists struct {
	ID                  uint    `json:"id"`
	UserID              uint    `json:"user_id"`
	ProductID           uint    `json:"product_id"`
	Name                string  `json:"name"`
	Description         string  `json:"description"`
	Currency            string  `json:"currency"`
	Price               uint64  `json:"price"`
	ImageURL            *string `json:"image_url"`
	Sold                uint64  `json:"sold"`
	ProductCategoryName string  `json:"product_category_name"`
	Quantity            uint64  `json:"quantity"`
}

type CreateWishlist struct {
	ProductID int `validate:"required,number" json:"product_id"`
	UserID    int `validate:"required,number" json:"user_id"`
}
