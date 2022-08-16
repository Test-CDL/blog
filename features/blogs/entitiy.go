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
	PostBlog(data Core) (row int, err error)
	PutBlog(idBlog int, data Core) (row int, err error)
	DeleteBlog(idBlog int) (row int, err error)
}

type Data interface {
	SelectAllBlogs(limit, offset int) (data []Core, err error)
	SelectSingleBlog(idBlog int) (data Core, err error)
	CreateBlog(data Core) (row int, err error)
	UpdateBlog(idBlog int, data Core) (row int, err error)
	DeleteBlog(idBlog int) (row int, err error)
	CountBlogData() (count int, err error)
}
