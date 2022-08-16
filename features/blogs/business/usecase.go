package business

import (
	"chrombit/features/blogs"
	"fmt"
)

type blogUseCase struct {
	blogData blogs.Data
}

func NewBlogBusiness(bgData blogs.Data) blogs.Business {
	return &blogUseCase{
		blogData: bgData,
	}
}

func (uc *blogUseCase) GetAllBlogs(limit, offset int) (data []blogs.Core, totalPage int, err error) {
	data, err = uc.blogData.SelectAllBlogs(limit, offset)
	total, _ := uc.blogData.CountBlogData()
	if total == 0 {
		totalPage = 0
	} else {
		if limit == 0 {
			limit = total
		}
		if total%limit != 0 {
			totalPage = (total / limit) + 1
		} else {
			totalPage = total / limit
		}
	}
	return data, totalPage, err
}

func (uc *blogUseCase) GetSingleBlog(idBlog int) (data blogs.Core, err error) {
	data, err = uc.blogData.SelectSingleBlog(idBlog)
	return data, err
}

func (uc *blogUseCase) PostBlog(input blogs.Core) (row int, err error) {
	if input.Title == "" || input.Body == "" || input.Slug == "" {
		return -1, fmt.Errorf("all input must be filled")
	}
	row, err = uc.blogData.CreateBlog(input)
	return row, err
}

func (uc *blogUseCase) PutBlog(idBlog int, update blogs.Core) (row int, err error) {
	if update.Title == "" || update.Body == "" || update.Slug == "" {
		return -1, fmt.Errorf("all input must be filled")
	}
	row, err = uc.blogData.UpdateBlog(idBlog, update)
	return row, err
}

func (uc *blogUseCase) DeleteBlog(idBlog int) (row int, err error) {
	row, err = uc.blogData.DeleteBlog(idBlog)
	return row, err
}
