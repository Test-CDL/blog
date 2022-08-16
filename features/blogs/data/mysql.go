package data

import (
	"chrombit/features/blogs"

	"gorm.io/gorm"
)

type mysqlBlogRepository struct {
	db *gorm.DB
}

func NewBlogRepository(conn *gorm.DB) blogs.Data {
	return &mysqlBlogRepository{
		db: conn,
	}
}

func (repo *mysqlBlogRepository) SelectAllBlogs(limit, offset int) ([]blogs.Core, error) {
	var dataBlogs []Blog
	result := repo.db.Limit(limit).Offset(offset).Find(&dataBlogs)
	if result.Error != nil {
		return nil, result.Error
	}
	return toCoreList(dataBlogs), nil
}

func (repo *mysqlBlogRepository) SelectSingleBlog(idBlog int) (blogs.Core, error) {
	var dataBlog Blog
	result := repo.db.First(&dataBlog, idBlog)
	if result.Error != nil {
		return blogs.Core{}, result.Error
	}
	return dataBlog.toCore(), nil
}

func (repo *mysqlBlogRepository) CreateBlog(input blogs.Core) (int, error) {
	var Blog = fromCore(input)
	result := repo.db.Create(&Blog)
	if result.Error != nil {
		return 0, result.Error
	}
	return 1, nil
}

func (repo *mysqlBlogRepository) CountBlogData() (int, error) {
	var count int64
	result := repo.db.Model(&Blog{}).Count(&count)
	if result.Error != nil {
		return -1, result.Error
	}
	return int(count), nil
}
