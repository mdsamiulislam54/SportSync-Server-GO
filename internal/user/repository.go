package user

import "gorm.io/gorm"

type Repository interface {
	CreateUser(user *User) error
	GetUserByEmail(email string) (*User, error)
	GetAllUsers() ([]User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db,
	}
}

func (r repository) CreateUser(user *User) error {
	return r.db.Create(user).Error
}

func (r repository) GetUserByEmail(email string) (*User, error) {
	var user User
	result := r.db.Where(&User{Email: email}).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *repository) GetAllUsers() ([]User, error) {
	var user []User
	result := r.db.Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
