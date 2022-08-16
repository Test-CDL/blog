package blogs

import "time"

type Core struct {
	ID        int
	Title     string
	Body      string
	Slug      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Business interface {
	GetAllBlogs(limit, offset int) (data []Core, totalPage int, err error)
	GetSingleBlog(idBlog int) (data Core, err error)
}

type Data interface {
	SelectAllBlogs(limit, offset int) (data []Core, err error)
	SelectSingleBlog(idBlog int) (data Core, err error)
	CountBlogData() (count int, err error)
}
