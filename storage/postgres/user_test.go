package postgres

import (
	"log"
	"testing"

	pb "github.com/Salikhov079/rent_car/genprotos"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	stg, err := NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}

	user := &pb.User{
		Id:       "b409ff53-ff2b-4033-84b4-4ce555081647",
		UserName: "Mukhammed_777",
		Password: "password123",
		Email:    "unique_test_user@example.com",
	}
	result, err := stg.User().Create(user)

	assert.NoError(t, err)
	assert.Nil(t, result)
}

func TestGetByIdUser(t *testing.T) {
	stg, err := NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}

	var Id pb.ById

	Id.Id = "b409ff53-ff2b-4033-84b4-4ce555081647"

	user, err := stg.User().GetById(&Id)

	assert.NoError(t, err)
 	assert.NotNil(t, user)
}

func TestGetAllUser(t *testing.T) {
	stg, err := NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}
	users, err := stg.User().GetAll(&pb.User{})
	assert.NoError(t, err)
	assert.NotNil(t, users)
}

func TestUpdateUser(t *testing.T) {
	stg, err := NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}

	user := &pb.User{
		Id:       "user_id_here",
		UserName: "updated_user",
		Email:    "updated_user@example.com",
	}
	result, err := stg.User().Update(user)

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestDeleteUser(t *testing.T) {
	stg, err := NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}

	var Id pb.ById

	Id.Id = "b409ff53-ff2b-4033-84b4-4ce555081647"

	result, err := stg.User().Delete(&Id)

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestLoginUser(t *testing.T) {
	stg, err := NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}

	user, err := stg.User().Login(&pb.User{})

	assert.NoError(t, err)
	assert.NotNil(t, user)
}
