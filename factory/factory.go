package factory

import (
	_blogBusiness "chrombit/features/blogs/business"
	_blogData "chrombit/features/blogs/data"
	_blogPresentation "chrombit/features/blogs/presentation"

	"gorm.io/gorm"
)

type Presenter struct {
	BlogPresenter *_blogPresentation.BlogHandler
}

func InitFactory(dbConn *gorm.DB) Presenter {
	blogData := _blogData.NewBlogRepository(dbConn)
	blogBusiness := _blogBusiness.NewBlogBusiness(blogData)
	blogPresentation := _blogPresentation.NewBlogHandler(blogBusiness)

	return Presenter{
		BlogPresenter: blogPresentation,
	}
}
