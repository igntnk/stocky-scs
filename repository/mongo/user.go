package mongo

import (
	"context"
	"github.com/igntnk/stocky-scs/models"
	"github.com/igntnk/stocky-scs/repository"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserRepository(database *mongo.Database, trxImpl bool, logger zerolog.Logger) repository.UserRepository {
	tx := noTxImpl
	if trxImpl {
		tx = txImpl
	}

	return &userRepository{
		Logger:         logger.With().Str("repository", repository.UserCollection).Logger(),
		UserCollection: database.Collection(repository.UserCollection),
		Tx:             tx,
	}
}

func getPipeline(limit, offset int64) mongo.Pipeline {
	pipeline := mongo.Pipeline{}
	if offset > 0 {
		pipeline = append(pipeline, bson.D{{
			"$skip",
			offset,
		}})
	}

	if limit > 0 {
		pipeline = append(pipeline, bson.D{{
			"$limit",
			limit,
		}})
	}
	return pipeline
}

type userRepository struct {
	Logger         zerolog.Logger
	UserCollection *mongo.Collection
	Tx             Tx
}

func (u userRepository) CreateUser(ctx context.Context, user *models.User) (string, error) {
	res, err := u.UserCollection.InsertOne(ctx, user)
	if err != nil {
		return "", err
	}

	return res.InsertedID.(string), nil
}

func (u userRepository) BlockUser(ctx context.Context, id string) (string, error) {
	resId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", err
	}

	res, err := u.UserCollection.UpdateOne(ctx, bson.M{"_id": resId}, bson.M{"$set": bson.M{"blocked": true}})
	if err != nil {
		return "", err
	}

	return res.UpsertedID.(string), nil
}

func (u userRepository) UnblockUser(ctx context.Context, id string) (string, error) {
	resId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", err
	}

	res, err := u.UserCollection.UpdateOne(ctx, bson.M{"_id": resId}, bson.M{"$set": bson.M{"blocked": false}})
	if err != nil {
		return "", err
	}

	return res.UpsertedID.(string), nil
}

func (u userRepository) UpdateUser(ctx context.Context, user *models.User) (string, error) {
	resId, err := primitive.ObjectIDFromHex(user.Id)
	if err != nil {
		return "", err
	}

	res, err := u.UserCollection.UpdateByID(ctx, resId, models.User{
		Name:           user.Name,
		Description:    user.Description,
		DocumentType:   user.DocumentType,
		DocumentNumber: user.DocumentNumber,
		CreationDate:   user.CreationDate,
	})
	if err != nil {
		return "", err
	}

	return res.UpsertedID.(string), nil
}

func (u userRepository) GetById(ctx context.Context, id string) (*models.User, error) {
	resId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	result := &models.User{}
	err = u.UserCollection.FindOne(ctx, bson.M{"_id": resId}).Decode(result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u userRepository) GetAll(ctx context.Context) ([]models.User, error) {
	cursor, err := u.UserCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var users []models.User
	for cursor.Next(ctx) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
