package business

import (
	"chrombit/features/users"
)

type userUsecase struct {
	userData users.Data
}

func NewUserBusiness(usrData users.Data) users.Business {
	return &userUsecase{
		userData: usrData,
	}
}

func (uc *userUsecase) InsertUser(input users.Core) (row int, err error) {
	row, err = uc.userData.PostUser(input)
	return row, err
}

func (uc *userUsecase) LoginUser(userName string, password string) (data string, token string, e error) {
	data, token, e = uc.userData.AuthUser(userName, password)
	return data, token, e
}
