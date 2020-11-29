package resolver

import (
	"context"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"report/internal/graphql/middleware"
	"report/internal/graphql/models"
)

func (m MutationResolver) Login(ctx context.Context, input models.LoginInput) (*models.AuthResponse, error) {

	_, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, fmt.Errorf("user not authorization, err: %v", err)
	}

	return m.Mutation.Login(input)
}

func (m MutationResolver) CreateUser(ctx context.Context, input models.UserInput) (*models.User, error) {
	return m.Mutation.CreateUser(input)
}

func (m MutationResolver) UpdateUser(ctx context.Context, id bson.ObjectId, input models.UserUpdate) (*models.User, error) {
	_, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, fmt.Errorf("user not authorization, err: %v", err)
	}

	return m.Mutation.UpdateUser(id, input)
}

