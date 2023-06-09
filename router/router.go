package router

import (
	"GDN-delivery-management/delivery/http"
	"github.com/labstack/echo/v4"
)

type Router struct {
	Echo               *echo.Echo
	UserHandler        http.UserHandler
	RoleHandler        http.RoleHandler
	GameHandler        http.GameHandler
	LibraryHandler     http.LibraryHandler
	PaymentHandler     http.PaymentHandler
	PaymentGameHandler http.PaymentGameHandler
}

func (r *Router) SetupRouter() {
	// health check
	r.Echo.GET("/health-check", http.HealthCheck)

	admin := r.Echo.Group("/admin")
	admin.POST("/sign-up", r.UserHandler.SystemAdminSignUp)

	r.Echo.POST("/login", r.UserHandler.Login)
	r.Echo.GET("/logout", r.UserHandler.Logout)

	r.Echo.POST("/upload", http.UploadImage)

	// r.Echo.POST("/token/refresh-token", r.UserHandler.RenewAccessToken, r.AuthMiddleware.UserCors())

	user := r.Echo.Group("/user")
	user.POST("/add", r.UserHandler.AddUser)
	user.GET("/:userid", r.UserHandler.UserDetails)
	user.GET("/get-me", r.UserHandler.GetMe)
	user.GET("/all", r.UserHandler.GetAllUsers)
	user.PATCH("/update", r.UserHandler.UpdateUser)
	user.DELETE("/delete", r.UserHandler.DeleteUser)
	user.GET("/admin-check", r.UserHandler.CheckAdmin)

	role := r.Echo.Group("/role")
	role.POST("/add", r.RoleHandler.AddRole)
	role.GET("/all", r.RoleHandler.ListRoles)
	role.PATCH("/update", r.RoleHandler.UpdateRole)
	role.DELETE("/delete", r.RoleHandler.DeleteRole)

	game := r.Echo.Group("/game")
	game.POST("/add", r.GameHandler.CreateGame)
	game.GET("/all", r.GameHandler.GetAllGames)
	game.PATCH("/update", r.GameHandler.UpdateGame)
	game.DELETE("/delete", r.GameHandler.DeleteGame)

	library := r.Echo.Group("/library")
	library.POST("/add", r.LibraryHandler.CreateLibrary)
	library.GET("/:userid", r.LibraryHandler.GetUserLibrary)
	library.POST("/check", r.LibraryHandler.CheckGameInLibrary)
	library.DELETE("/delete", r.LibraryHandler.DeleteLibrary)

	payment := r.Echo.Group("/payment")
	payment.POST("/add", r.PaymentHandler.CreatePayment)
	payment.GET("/:userid", r.PaymentHandler.GetUserPayments)
	payment.DELETE("/delete", r.PaymentHandler.DeletePayment)

	paymentGame := r.Echo.Group("/payment-game")
	paymentGame.POST("/add", r.PaymentGameHandler.CreatePaymentGame)
	paymentGame.GET("/:paymentid", r.PaymentGameHandler.GetPaymentGames)
	paymentGame.DELETE("/delete", r.PaymentGameHandler.DeletePaymentGame)
}
