package main

import (
	"log"

	"github.com/adhityaf/challenge-steradian-be/config"
	"github.com/adhityaf/challenge-steradian-be/controllers"
	"github.com/adhityaf/challenge-steradian-be/middlewares"
	"github.com/adhityaf/challenge-steradian-be/repositories"
	"github.com/adhityaf/challenge-steradian-be/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	//Load Env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := config.ConnectDB()
	route := gin.Default()

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	carRepository := repositories.NewCarRepository(db)
	carService := services.NewCarService(carRepository)
	carController := controllers.NewCarController(carService)

	adminRepository := repositories.NewAdminRepository(db)
	adminService := services.NewAdminService(adminRepository)
	adminController := controllers.NewAdminController(adminService)

	orderRepository := repositories.NewOrderRepository(db)
	orderService := services.NewOrderService(orderRepository, carRepository, adminRepository)
	orderController := controllers.NewOrderController(orderService)

	mainRouter := route.Group("/v1")
	{
		mainRouter.POST("/admins/login", adminController.Login)

		mainRouter.POST("/login", userController.Login)
		mainRouter.POST("/register", userController.Register)

		authorized := mainRouter.Group("/")
		authorized.Use(middlewares.Auth())
		{
			// CRUD Order
			authorized.POST("/orders", orderController.CreateOrder)
			authorized.GET("/orders", orderController.GetAllOrders)
			authorized.GET("/orders/:orderId", orderController.GetOrderById)
			authorized.PUT("/orders/:orderId", orderController.UpdateOrderById)
			authorized.DELETE("/orders/:orderId", orderController.DeleteOrderById)

			// CRUD user
			authorized.GET("/users/profile", userController.GetUserProfile)
			authorized.PUT("/users", userController.UpdateUserProfile)

			admin := authorized.Group("/")
			admin.Use(middlewares.IsAdmin())
			{
				// CRUD User
				admin.GET("/users", userController.GetAllUsers)
				admin.DELETE("/users/:userId", userController.DeleteUserById)

				// CRUD Admin
				admin.POST("/admins", adminController.CreateAdmin)
				admin.GET("/admins", adminController.GetAllAdmins)
				admin.GET("/admins/:adminId", adminController.GetAdminById)
				admin.PUT("/admins", adminController.UpdatePasswordAdmin)
				admin.DELETE("/admins/:adminId", adminController.DeleteAdminById)

				// CRUD Car
				admin.POST("/cars", carController.CreateCar)
				admin.GET("/cars", carController.GetAllCars)
				admin.GET("/cars/:carId", carController.GetCarById)
				admin.PUT("/cars/:carId", carController.UpdateCarById)
				admin.DELETE("/cars/:carId", carController.DeleteCarById)
			}
		}
	}

	route.Run()
}
