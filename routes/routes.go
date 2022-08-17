package routes

import (
	"chrombit/factory"
	"chrombit/middlewares"

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

	e.POST("api/v1/users", presenter.UserPresenter.AddUser)
	e.POST("api/v1/login", presenter.UserPresenter.Login)

	e.GET("api/v1/blogs", presenter.BlogPresenter.GetAllBlogs)
	e.GET("api/v1/blogs/:id", presenter.BlogPresenter.GetSingleBlog)
	e.POST("api/v1/blogs", presenter.BlogPresenter.CreateBlog, middlewares.JWTMiddleware())
	e.PUT("api/v1/blogs/:id", presenter.BlogPresenter.UpdateBlog, middlewares.JWTMiddleware())
	e.DELETE("api/v1/blogs/:id", presenter.BlogPresenter.DeleteBlog, middlewares.JWTMiddleware())
	return e
}
