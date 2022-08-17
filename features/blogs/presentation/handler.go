package presentation

import (
	"chrombit/features/blogs"
	_requestBlog "chrombit/features/blogs/presentation/request"
	_responseBlog "chrombit/features/blogs/presentation/response"
	_helpers "chrombit/helpers"
	"chrombit/middlewares"
	"strconv"

	"net/http"

	"github.com/labstack/echo/v4"
)

type BlogHandler struct {
	blogBusiness blogs.Business
}

func NewBlogHandler(business blogs.Business) *BlogHandler {
	return &BlogHandler{
		blogBusiness: business,
	}
}

func (h *BlogHandler) GetAllBlogs(c echo.Context) error {
	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")
	limitInt, _ := strconv.Atoi(limit)
	offsetInt, _ := strconv.Atoi(offset)
	result, totalPage, err := h.blogBusiness.GetAllBlogs(limitInt, offsetInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helpers.ResponseFailed("failed to get all data"))
	}
	var res = map[string]interface{}{
		"data":       _responseBlog.FromCoreList(result),
		"total_page": totalPage,
	}
	return c.JSON(http.StatusOK, _helpers.ResponseSuccesWithData("success to get all data", res))
}

func (h *BlogHandler) GetSingleBlog(c echo.Context) error {
	idBlog := c.Param("id")
	idBg, _ := strconv.Atoi(idBlog)

	result, err := h.blogBusiness.GetSingleBlog(idBg)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helpers.ResponseFailed("failed to get data"))
	}

	return c.JSON(http.StatusOK, _helpers.ResponseSuccesWithData("success to get data", _responseBlog.FromCore(result)))
}

func (h *BlogHandler) CreateBlog(c echo.Context) error {
	idTok, errDel := middlewares.ValidateToken(c)
	if errDel != nil {
		return c.JSON(http.StatusBadRequest, _helpers.ResponseFailed("invalid token"))
	}

	if idTok == 0 {
		return c.JSON(http.StatusUnauthorized, _helpers.ResponseFailed("unauthorized"))
	}
	var inputBlog _requestBlog.Blog
	errBind := c.Bind(&inputBlog)
	if errBind != nil {
		return c.JSON(http.StatusInternalServerError, _helpers.ResponseFailed("failed to bind data"))
	}
	result, err := h.blogBusiness.PostBlog(_requestBlog.ToCore(inputBlog))
	if result == -1 {
		return c.JSON(http.StatusBadRequest, _helpers.ResponseFailed("all input must be filled"))
	}
	if result == 0 {
		return c.JSON(http.StatusInternalServerError, _helpers.ResponseFailed("failed to create blog"))
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helpers.ResponseFailed("failed to create blog"))
	}

	return c.JSON(http.StatusOK, _helpers.ResponseSuccesNoData("Succes to insert blog"))
}

func (h *BlogHandler) UpdateBlog(c echo.Context) error {
	idTok, errDel := middlewares.ValidateToken(c)
	if errDel != nil {
		return c.JSON(http.StatusBadRequest, _helpers.ResponseFailed("invalid token"))
	}

	if idTok == 0 {
		return c.JSON(http.StatusUnauthorized, _helpers.ResponseFailed("unauthorized"))
	}

	idBlog := c.Param("id")
	idBlogInt, errIdBlog := strconv.Atoi(idBlog)
	if errIdBlog != nil {
		return c.JSON(http.StatusBadRequest, _helpers.ResponseFailed("failed id blog not recognize"))
	}

	var dataBlog _requestBlog.Blog
	errBind := c.Bind(&dataBlog)
	if errBind != nil {
		return c.JSON(http.StatusInternalServerError, _helpers.ResponseFailed("failed to bind data"))
	}
	result, err := h.blogBusiness.PutBlog(idBlogInt, _requestBlog.ToCore(dataBlog))
	if result == -1 {
		return c.JSON(http.StatusBadRequest, _helpers.ResponseFailed("all input must be filled"))
	}
	if result == 0 {
		return c.JSON(http.StatusInternalServerError, _helpers.ResponseFailed("failed to update blog"))
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helpers.ResponseFailed("failed to update blog"))
	}
	return c.JSON(http.StatusOK, _helpers.ResponseSuccesNoData("success to update blog"))
}

func (h *BlogHandler) DeleteBlog(c echo.Context) error {
	idTok, errDel := middlewares.ValidateToken(c)
	if errDel != nil {
		return c.JSON(http.StatusBadRequest, _helpers.ResponseFailed("invalid token"))
	}

	if idTok == 0 {
		return c.JSON(http.StatusUnauthorized, _helpers.ResponseFailed("unauthorized"))
	}

	idBlog := c.Param("id")
	idBlogInt, errIdBlog := strconv.Atoi(idBlog)
	if errIdBlog != nil {
		return c.JSON(http.StatusBadRequest, _helpers.ResponseFailed("failed id blog not recognize"))
	}

	_, err := h.blogBusiness.DeleteBlog(idBlogInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helpers.ResponseFailed("failed to delete blog"))
	}
	return c.JSON(http.StatusOK, _helpers.ResponseSuccesNoData("success to delete blog"))
}
