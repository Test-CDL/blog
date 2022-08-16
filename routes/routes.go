package routes

import (
	"chrombit/factory"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New(presenter factory.Presenter) *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))
	e.Pre(middleware.RemoveTrailingSlash())

	e.GET("api/v1/blogs", presenter.BlogPresenter.GetAllBlogs)
	e.GET("api/v1/blogs/:id", presenter.BlogPresenter.GetSingleBlog)
	e.POST("api/v1/blogs", presenter.BlogPresenter.CreateBlog)
	e.PUT("api/v1/blogs/:id", presenter.BlogPresenter.UpdateBlog)
	return e
}
