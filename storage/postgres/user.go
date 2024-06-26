package postgres

import (
	"database/sql"
	"fmt"
	"time"

	pb "github.com/Salikhov079/rent_car/genprotos"

	"github.com/google/uuid"
)

type UserStorage struct {
	db *sql.DB
}

func NewUserStorage(db *sql.DB) *UserStorage {
	return &UserStorage{db: db}
}

func (p *UserStorage) Create(user *pb.User) (*pb.Void, error) {
	id := uuid.NewString()
	query := `
		INSERT INTO users (id, user_name, email, password, phone_number, role)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := p.db.Exec(query, id, user.UserName, user.Email, user.Password, user.PhoneNumber, user.Role)
	return nil, err
}

func (p *UserStorage) GetById(id *pb.ById) (*pb.User, error) {
	query := `
			SELECT user_name, email, phone_number FROM users 
			WHERE id = $1 and deleted_at = 0
		`
	row := p.db.QueryRow(query, id.Id)

	var user pb.User

	err := row.Scan(
		&user.UserName,
		&user.Email,
		&user.PhoneNumber)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (p *UserStorage) GetAll(us *pb.User) (*pb.GetAllUsers, error) {
	users := &pb.GetAllUsers{}
	var arr []interface{}
	count := 1

	query := ` SELECT user_name, email, phone_number FROM users 
	WHERE deleted_at = 0 `

	if len(us.Email) > 0 {
		query += fmt.Sprintf(" and email=$%d", count)
		count++
		arr = append(arr, us.Email)
	}

	if len(us.UserName) > 0 {
		query += fmt.Sprintf(" and user_name=$%d", count)
		count++
		arr = append(arr, us.UserName)
	}

	row, err := p.db.Query(query, arr...)
	if err != nil {
		return nil, err
	}

	for row.Next() {
		var user pb.User
		err = row.Scan(&user.UserName,
			&user.Email,
			&user.PhoneNumber)
		if err != nil {
			return nil, err
		}
		users.Users = append(users.Users, &user)
	}

	return users, nil
}

func (p *UserStorage) Update(user *pb.User) (*pb.Void, error) {
	query := `
		UPDATE users
		SET user_name = $2, email = $3,  phone_number = $4, updated_at = now()
		WHERE id = $1 AND deleted_at = 0
	`
	_, err := p.db.Exec(query, user.Id, user.UserName, user.Email, user.PhoneNumber)
	return nil, err
}

func (p *UserStorage) Delete(id *pb.ById) (*pb.Void, error) {
	query := `
		UPDATE users SET deleted_at = $2
		WHERE id = $1 AND deleted_at = 0
	`
	_, err := p.db.Exec(query, id.Id, time.Now().Unix())
	return nil, err
}

func (p *UserStorage) Login(userName *pb.User) (*pb.User, error) {
	query := `
			SELECT user_name, email, phone_number, role FROM users 
			WHERE user_name = $1 and deleted_at = 0
		`
	row := p.db.QueryRow(query, userName.UserName)

	var user pb.User

	err := row.Scan(&user.UserName,
		&user.Email,
		&user.PhoneNumber, 
		&user.Role)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
