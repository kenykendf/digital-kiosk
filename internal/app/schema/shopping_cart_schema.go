package schema

type CreateShoppingCartReq struct {
	Quantity  int `validate:"required" json:"quantity"`
	ProductID int `validate:"required" json:"product_id"`
	UserID    int `validate:"required" json:"user_id"`
}

type GetShoppingCartResp struct {
	ID       int     `json:"id"`
	Quantity int     `json:"quantity"`
	Total    int     `json:"total"`
	Product  Product `json:"product"`
}

type GetShoppingCartReq struct {
	UserID int `validate:"required" json:"user_id"`
}

type UpdateShoppingCartReq struct {
	Quantity int `validate:"required" json:"quantity"`
	UserID   int `validate:"required" json:"user_id"`
}

type Product struct {
	ProductID       int     `json:"product_id"`
	ProductName     string  `json:"product_name"`
	ProductPrice    uint64  `json:"product_price"`
	ProductImageURL *string `json:"product_image_url"`
}
