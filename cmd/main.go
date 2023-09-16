package main

import (
	"fmt"
	"time"

	"kenykendf/digital-kiosk/internal/app/controller"
	"kenykendf/digital-kiosk/internal/app/repository"
	"kenykendf/digital-kiosk/internal/app/service"
	"kenykendf/digital-kiosk/internal/pkg/config"
	"kenykendf/digital-kiosk/internal/pkg/db"
	"kenykendf/digital-kiosk/internal/pkg/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"

	repo "kenykendf/digital-kiosk/internal/app/model"
)

var (
	cfg    config.Config
	DBConn *gorm.DB
)

const (
	limit  = 10
	offset = 0
	asc    = 0
)

func init() {

	configLoad, err := config.LoadConfig(".")
	if err != nil {
		log.Panic("cannot load app config")
	}
	cfg = configLoad

	db, err := db.ConnectDB(cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)
	if err != nil {
		log.Panic("db not established")
	}
	DBConn = db
	if db.Migrator().HasConstraint(&repo.Products{}, "products_name_key") {
		db.Migrator().DropConstraint(&repo.Products{}, "products_name_key")
	}
	db.AutoMigrate(
		&repo.ProductCategories{},
		&repo.User{},
		&repo.Products{},
		&repo.Auth{},
		&repo.Wishlist{},
		&repo.ShoppingCart{},
	)

	// Setup logrus
	logLevel, err := log.ParseLevel("debug")
	if err != nil {
		logLevel = log.InfoLevel
	}

	log.SetLevel(logLevel)                 // appyly log level
	log.SetFormatter(&log.JSONFormatter{}) // define format using json

}

func main() {

	r := gin.New()

	// implement middleware
	r.Use(
		middleware.LoggingMiddleware(),
		middleware.RecoveryMiddleware(),
		cors.New(cors.Config{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"OPTIONS", "GET", "POST", "PATCH", "DELETE"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "authorization"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			AllowOriginFunc: func(origin string) bool {
				return origin == "http://localhost"
			},
			MaxAge: 12 * time.Hour,
		}))

	// ---------------------------------------------------------------------------------------

	userRepository := repository.NewUserRepository(DBConn)
	authRepository := repository.NewAuthRepository(DBConn)

	tokenMaker := service.NewTokenMaker(
		cfg.AccessTokenKey,
		cfg.RefreshTokenKey,
		cfg.AccessTokenDuration,
		cfg.RefreshTokenDuration,
	)

	registrationService := service.NewRegistrationService(userRepository)
	sessionService := service.NewSessionService(userRepository, authRepository, tokenMaker)
	userService := service.NewUserService(userRepository)

	registrationController := controller.NewRegistrationController(registrationService)
	sessionController := controller.NewSessionController(sessionService, tokenMaker)

	productCategoryRepo := repository.NewProductCategoryRepo(DBConn)
	productCategoryService := service.NewProductCategoryService(productCategoryRepo)
	productCategoryController := controller.NewProductCategoryController(productCategoryService)
	userController := controller.NewUserController(userService)

	productRepo := repository.NewProductRepo(DBConn)
	productService := service.NewProductService(productRepo, productCategoryRepo)
	productController := controller.NewProductController(productService)

	wishlistRepo := repository.NewWishlistRepo(DBConn)
	wishlistService := service.NewWishlistService(wishlistRepo, productRepo)
	wishlistController := controller.NewWishlistController(wishlistService)

	shoppingCartRepository := repository.NewShoppingCartRepository(DBConn)
	shoppingCartService := service.NewShoppingCartService(shoppingCartRepository, productRepo)
	shoppingCartController := controller.NewShoppingCartController(shoppingCartService, productService)

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "pong"})
	})

	route := r.Group("/api/auth")
	{
		route.POST("/signup", registrationController.Register)
		route.POST("/signin", sessionController.Login)
		route.GET("/refresh", sessionController.Refresh)
	}

	public := r.Group("/public/api")
	{
		public.GET("/products", productController.GetProductsLists)
		public.GET("/product/:id", productController.GetProductsLists)
		public.GET("/product-categories", productCategoryController.GetProductCategoriesLists)
		public.GET("/product-category/:id", productCategoryController.GetProductCategoryByID)
	}

	secured := r.Group("/api").Use(middleware.AuthMiddleware(tokenMaker))
	{
		secured.GET("/auth/signout", sessionController.Logout)

		secured.GET("/users", userController.BrowseUser)
		secured.GET("/user/:id", userController.DetailUser)
		secured.DELETE("/user/:id", userController.DeleteUser)

		secured.POST("/product-category", productCategoryController.CreateProductCategory)
		secured.GET("/product-categories", productCategoryController.GetProductCategoriesLists)
		secured.GET("/product-category/:id", productCategoryController.GetProductCategoryByID)
		secured.PUT("/product-category/:id", productCategoryController.UpdateProductCategory)
		secured.DELETE("/product-category/:id", productCategoryController.DeleteProductCategory)

		secured.POST("/product", productController.CreateProduct)
		secured.GET("/products", middleware.PaginationMiddleware(offset, limit, asc), productController.GetProductsLists)
		secured.GET("/product/:id", productController.GetProductByID)
		secured.PUT("/product/:id", productController.UpdateProduct)
		secured.PUT("/product/sell/:id", productController.UpdateProductSell)
		secured.DELETE("/product/:id", productController.DeleteProduct)

		secured.POST("/wishlist", wishlistController.CreateWishlist)
		secured.GET("/wishlists", wishlistController.GetWishlistLists)
		secured.DELETE("/wishlist/:id", wishlistController.DeleteWishlist)

		secured.POST("/shopping-cart", shoppingCartController.CreateShoppingCart)
		secured.GET("/shopping-carts", shoppingCartController.BrowseShoppingCart)
		secured.PATCH("/shopping-cart/:id", shoppingCartController.UpdateShoppingCart)
		secured.DELETE("/shopping-cart/:id", shoppingCartController.DeleteShoppingCart)
	}

	appPort := fmt.Sprintf(":%s", cfg.ServerPort)
	// nolint:errcheck
	r.Run(appPort)
}
