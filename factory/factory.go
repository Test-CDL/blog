package factory

import (
	_blogBusiness "chrombit/features/blogs/business"
	_blogData "chrombit/features/blogs/data"
	_blogPresentation "chrombit/features/blogs/presentation"

	_userBusiness "chrombit/features/users/business"
	_userData "chrombit/features/users/data"
	_userPresentation "chrombit/features/users/presentation"

	"gorm.io/gorm"
)

type Presenter struct {
	BlogPresenter *_blogPresentation.BlogHandler
	UserPresenter *_userPresentation.UserHandler
}

func InitFactory(dbConn *gorm.DB) Presenter {
	blogData := _blogData.NewBlogRepository(dbConn)
	blogBusiness := _blogBusiness.NewBlogBusiness(blogData)
	blogPresentation := _blogPresentation.NewBlogHandler(blogBusiness)

	userData := _userData.NewUserRepository(dbConn)
	userBusiness := _userBusiness.NewUserBusiness(userData)
	userPresentation := _userPresentation.NewUserHandler(userBusiness)

	return Presenter{
		BlogPresenter: blogPresentation,
		UserPresenter: userPresentation,
	}
}
