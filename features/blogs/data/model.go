package data

import (
	"chrombit/features/blogs"

	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	Title string `json:"title"`
	Body  string `json:"body"`
	Slug  string `json:"slug"`
}

func (data *Blog) toCore() blogs.Core {
	return blogs.Core{
		ID:        int(data.ID),
		Title:     data.Title,
		Body:      data.Body,
		Slug:      data.Slug,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func toCoreList(data []Blog) []blogs.Core {
	result := []blogs.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}

func fromCore(core blogs.Core) Blog {
	return Blog{
		Title: core.Title,
		Body:  core.Body,
		Slug:  core.Slug,
	}
}
