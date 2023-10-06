package services

import (
	"amandeepsTlgt/REST-API/models"
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceInterface interface {
	CreateUser(*models.User) (*models.User, error)
	GetUserList() ([]*models.User, error)
	GetUserDetails(*string) (*models.User, error)
	UpdateUser(*string, *models.User) (*models.User, error)
	DeleteUser(*string) error
}

type UserService struct {
	userCollection *mongo.Collection
	context        context.Context
}

func NewUserServiceInstance(userCollection *mongo.Collection, ctx context.Context) UserServiceInterface {
	return &UserService{
		userCollection: userCollection,
		context:        ctx, // Use the ctx parameter here
	}
}

/**
 * CreateUser takes a pointer to a models.User object as input and returns a pointer to a models.User object and an error.
 *
 * @param user - a pointer to a models.User object
 * @return createdUser - a pointer to a models.User object representing the created user
 * @return err - an error object representing any errors that occurred during the method execution
 */
func (userService *UserService) CreateUser(user *models.User) (*models.User, error) {
	user.UserId = (uuid.New()).String()
	user.Id = primitive.NewObjectID()
	_, err := userService.userCollection.InsertOne(userService.context, user)
	return user, err
}

/**
 * GetUserList takes no input and returns a slice of pointers to models.User objects and an error.
 *
 * @return userList - a slice of pointers to models.User objects representing the list of users
 * @return err - an error object representing any errors that occurred during the method execution
 */
func (userService *UserService) GetUserList() ([]*models.User, error) {
	var userList []*models.User
	cursor, err := userService.userCollection.Find(userService.context, bson.D{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(userService.context) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		userList = append(userList, &user)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	cursor.Close(userService.context)
	return userList, nil
}

/**
 * GetUserDetails takes a pointer to a string representing the user ID as input and returns a pointer to a models.User object and an error.
 *
 * @param userId - a pointer to a string representing the user ID
 * @return userDetails - a pointer to a models.User object representing the user details
 * @return err - an error object representing any errors that occurred during the method execution
 */
func (userService *UserService) GetUserDetails(userId *string) (*models.User, error) {
	var userDetails *models.User
	query := bson.D{bson.E{Key: "userId", Value: userId}}
	err := userService.userCollection.FindOne(userService.context, query).Decode(&userDetails)
	return userDetails, err
}

/**
 * UpdatedUser takes a pointer to a models.User object as input and returns a pointer to a models.User object and an error.
 *
 * @param user - a pointer to a models.User object
 * @return updatedUser - a pointer to a models.User object representing the updated user
 * @return err - an error object representing any errors that occurred during the method execution
 */
func (userService *UserService) UpdateUser(userId *string, user *models.User) (*models.User, error) {
	whereCondition := bson.D{bson.E{Key: "userId", Value: userId}}
	updatedDate := bson.D{bson.E{Key: "$set", Value: bson.D{
		bson.E{Key: "firstName", Value: user.FirstName},
		bson.E{Key: "lastName", Value: user.LastName},
		bson.E{Key: "email", Value: user.Email},
		bson.E{Key: "mobileNumber", Value: user.MobileNumber},
	}}}
	_, err := userService.userCollection.UpdateOne(userService.context, whereCondition, updatedDate)
	return user, err
}

/**
 * DeleteUser takes a pointer to a string representing the user ID as input and returns an error.
 *
 * @param userId - a pointer to a string representing the user ID
 * @return err - an error object representing any errors that occurred during the method execution
 */
func (userService *UserService) DeleteUser(userId *string) error {
	whereCondition := bson.D{bson.E{Key: "userId", Value: userId}}
	_, err := userService.userCollection.DeleteOne(userService.context, whereCondition)
	return err
}
