package store

import (
	"gopkg.in/mgo.v2/bson"
	"report/internal/graphql/models"
)

func (m *mutation) Login(input models.LoginInput) (*models.AuthResponse, error) {
	var user = &models.User{}
	var err error
	q := bson.M{}
	q["username"] = input.Username
	if err = m.GetStore("user").FindOne(q, user); err != nil {
		return nil, errBadCredentials
	}

	err = user.ComparePassword(input.Password)
	if err != nil {
		return nil, errBadCredentials
	}

	token, err := user.GenToken()
	if err != nil {
		return nil, errUnknown
	}

	return &models.AuthResponse{
		AuthToken: token,
		User:      user,
	}, nil
}

func (m *mutation) CreateUser(input models.UserInput) (*models.User, error) {

	user := &models.User{
		Name:     input.Name,
		Username: input.Username,
		Password: input.Password,
		Location: input.Location,
		Abbr:     input.Abbr,
		Email:    input.Email,
		Openhab:  input.Openhab,
	}

	if err := m.GetStore("user").Save(user); err != nil {
		return nil, errCreateRecord
	}
	return user, nil
}

func (m *mutation) UpdateUser(id bson.ObjectId, input models.UserUpdate) (*models.User, error) {


	q := bson.M{}
	q["_id"] = id

	update := bson.M{}
	update["Password"] = input.Password

	if err := m.GetStore("user").Update(q, update); err != nil {
		return nil, err
	}

	return nil, errNotImplemented
}
