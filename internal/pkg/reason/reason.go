package reason

var (
	InternalServerError = "internal server error"
	RequestFormError    = "request format is not valid"

	UserCannotBrowse    = "cannot Browse user"
	UserCannotUpdate    = "cannot Update user"
	UserCannotDelete    = "cannot Delete user"
	UserCannotGetDetail = "cannot get detail user"
	UserNotFound        = "user not found"

	UserAddressNotFound        = "userAddress not found"
	UserAddressCannotCreate    = "cannot Create userAddress"
	UserAddressCannotBrowse    = "cannot Browse userAddress"
	UserAddressCannotUpdate    = "cannot Update userAddress"
	UserAddressCannotDelete    = "cannot Delete userAddress"
	UserAddressCannotGetDetail = "cannot get detail userAddress"

	RegisterFailed      = "cannot register user"
	UserAlreadyExist    = "user already exist"
	LoginFailed         = "login failed, please check your email or password"
	SaveToken           = "cannot save refresh token" // nolint:gosec
	UserNotAuthenticate = "user does not have an authentication"
	NotAuthorized       = "You are not authorized to access this resource"

	ErrInvalidToken         = "token is invalid"
	ErrNoToken              = "request does not contain an access token"
	InvalidRefreshToken     = "invalid refresh token"
	CannotCreateAccessToken = "cannot create access token"

	GetListsProductCatErr  = "unable to get product categories"
	GetDetailProductCatErr = "unable to get product category"
	CreateProductCatErr    = "error while creating new product category"
	UpdateProductCatErr    = "error updating product category."
	DeleteProdcutCatErr    = "unable to delete product category"

	GetListsProductErr  = "unable to get products"
	GetDetailProductErr = "unable to get product"
	CreateProductErr    = "error while creating new product"
	UpdateProductErr    = "error updating product."
	DeleteProdcutErr    = "unable to delete product"

	GetListsWishlistErr  = "unable to get products"
	GetDetailWishlistErr = "unable to get product"
	CreateWishlistErr    = "error while creating new product"
	UpdateWishlistErr    = "error updating product."
	DeleteWishlistErr    = "unable to delete product"

	ShoppingCartNotFound        = "shopping cart not found"
	ShoppingCartCannotCreate    = "cannot Create shopping cart"
	ShoppingCartCannotBrowse    = "cannot Browse shopping cart"
	ShoppingCartCannotUpdate    = "cannot Update shopping cart"
	ShoppingCartCannotDelete    = "cannot Delete shopping cart"
	ShoppingCartCannotGetDetail = "cannot get detail shopping cart"
)
