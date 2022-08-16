package business

import (
	"chrombit/features/blogs"
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
