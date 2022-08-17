package data

import (
	_enciption "chrombit/encription"
	"chrombit/features/users"
	"fmt"

	_middlewares "chrombit/middlewares"

	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	db *gorm.DB
}

func NewUserRepository(conn *gorm.DB) users.Data {
	return &mysqlUserRepository{
		db: conn,
	}
}

func (repo *mysqlUserRepository) PostUser(input users.Core) (row int, err error) {
	passHash, _ := _enciption.HashPassword(input.Password)
	user := User{
		UserName: input.UserName,
		Password: passHash,
	}
	tx := repo.db.Create(&user)
	if tx.Error != nil {
		return 0, tx.Error
	}
	if tx.RowsAffected != 1 {
		return 0, fmt.Errorf("failed to insert data")
	}
	return int(tx.RowsAffected), nil
}

func (repo *mysqlUserRepository) AuthUser(userName string, password string) (data string, token string, e error) {
	userData := User{}
	repo.db.Where("user_name = ?", userName).First(&userData)
	// bool1 := _bcrypt.CheckPasswordHash(password, userData.Password)

	// if !bool1 {
	// 	return "", "", fmt.Errorf("error")
	// }

	token, errToken := _middlewares.CreateToken(int(userData.ID))

	if errToken != nil {
		return "", "", errToken
	}
	return token, userData.UserName, nil
}
