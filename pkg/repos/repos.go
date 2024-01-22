package repos

import (
	"context"
	"crud/pkg/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	Collection *mongo.Collection
}

type UserRepositoryInterface interface {
	CreateUser(models.User) (any, error)
	ReadUser() (any, error)
	DeleteUsers() error
	GetOneUser(string) (any, error)
	UpdateUser(string, models.User) error
	DeleteUser(string) error
}

func NewRepository(clctn *mongo.Collection) UserRepositoryInterface {
	return &UserRepository{
		Collection: clctn,
	}
}

func (ur *UserRepository) CreateUser(usr models.User) (any, error) {
	id, err := ur.Collection.InsertOne(context.TODO(), models.User{
		Name: usr.Name,
		Age:  usr.Age,
	})
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (ur *UserRepository) ReadUser() (any, error) {
	cur, err := ur.Collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var users []bson.M
	for cur.Next(context.Background()) {
		var user bson.M
		if err := cur.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (ur *UserRepository) DeleteUsers() error {
	return ur.Collection.Drop(context.TODO())
}

func (ur *UserRepository) GetOneUser(id string) (any, error) {
	usrId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	var usr models.User
	filter := bson.M{"_id": usrId}
	err = ur.Collection.FindOne(context.Background(), filter).Decode(&usr)

	if err != nil {
		return nil, err
	}
	return usr, nil
}
func (ur *UserRepository) UpdateUser(id string, data models.User) error {
	usrId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return err
	}

	filter := bson.M{"_id": usrId}
	update := bson.M{"$set": bson.M{"name": data.Name, "age": data.Age}}

	_, err = ur.Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (ur *UserRepository) DeleteUser(id string) error {
	usrId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return err
	}

	filter := bson.M{"_id":usrId}
	_,err = ur.Collection.DeleteOne(context.Background(),filter)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
