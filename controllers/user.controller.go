package controllers

import (
	"amandeepsTlgt/REST-API/models"
	"amandeepsTlgt/REST-API/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.UserServiceInterface
}

func NewUserControllerInstance(userService services.UserServiceInterface) UserController {
	return UserController{
		UserService: userService,
	}
}

/**
 * CreateUser handles the creation of a user.
 * @param context - A pointer to a gin.Context object representing the HTTP request and response.
 * @return A JSON response with the message "Hello from createdUserController".
 */
func (userController *UserController) CreateUser(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	latestCreatedUser, errorWhileCreatingUser := userController.UserService.CreateUser(&user)
	if errorWhileCreatingUser != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadGateway,
			"message": "Error occurred while creating user, Please contact with support team.",
		})
		return
	}
	context.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "user Account Created successfully.",
		"data":    latestCreatedUser,
	})
	// user, err := userse
	// context.JSON(200, "Hello from createdUserController")
}

/**
 * GetUserList retrieves a list of users.
 * @param context - A pointer to a gin.Context object representing the HTTP request and response.
 * @return A JSON response with the message "Hello from getUserListController".
 */
func (userController *UserController) GetUserList(context *gin.Context) {
	userList, err := userController.UserService.GetUserList()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadGateway,
			"message": "Error occurred while fetching lis of users, Please contact with support team.",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "userList fetched successfully.",
		"data":    userList,
	})
}

/**
 * GetUserDetails retrieves the details of a specific user.
 * @param context - A pointer to a gin.Context object representing the HTTP request and response.
 * @return A JSON response with the message "Hello from GetUserDetails controller".
 */
func (userController *UserController) GetUserDetails(context *gin.Context) {
	userId := context.Param("userId")
	userDetails, err := userController.UserService.GetUserDetails(&userId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadGateway,
			"message": "Error occurred while creating user, Please contact with support team.",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "userDetails fetched successfully.",
		"data":    userDetails,
	})
	// context.JSON(200, "Hello from GetUserDetails controller")
}

/**
 * UpdateUser updates the details of a user.
 * @param context - A pointer to a gin.Context object representing the HTTP request and response.
 * @return A JSON response with the message "Hello from updateUsercontroller".
 */
func (userController *UserController) UpdateUser(context *gin.Context) {
	userId := context.Param("userId")
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	updatedUser, errorWhileUpdatingUser := userController.UserService.UpdateUser(&userId, &user)
	if errorWhileUpdatingUser != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadGateway,
			"message": "Error occurred while updating user, Please contact with support team.",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "user Account updated successfully.",
		"data":    updatedUser,
	})
	// context.JSON(200, "Hello from updateUsercontroller")
}

/**
 * DeleteUser deletes a user.
 * @param context - A pointer to a gin.Context object representing the HTTP request and response.
 * @return A JSON response with the message "Hello from DeleteUserController".
 */
func (userController *UserController) DeleteUser(context *gin.Context) {
	userId := context.Param("userId")
	err := userController.UserService.DeleteUser(&userId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadGateway,
			"message": "Error occurred while deleting user, Please contact with support team.",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "user Account deleted successfully.",
	})
}

func (userController *UserController) RegisterUserRoutes(routeGroup *gin.RouterGroup) {
	userRoutes := routeGroup.Group("/users")
	userRoutes.GET("/list", userController.GetUserList)
	userRoutes.GET("/:userId", userController.GetUserDetails)
	userRoutes.POST("/", userController.CreateUser)
	userRoutes.PATCH("/:userId", userController.UpdateUser)
	userRoutes.DELETE("/:userId", userController.DeleteUser)
}
