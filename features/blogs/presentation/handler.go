package presentation

import (
	"chrombit/features/blogs"
	_responseBlog "chrombit/features/blogs/presentation/response"
	_helpers "chrombit/helpers"
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
