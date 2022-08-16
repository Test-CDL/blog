package response

import (
	"chrombit/features/blogs"
)

type Blog struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
	Slug  string `json:"slug"`
}

func FromCore(data blogs.Core) Blog {
	return Blog{
		ID:    data.ID,
		Title: data.Title,
		Body:  data.Body,
		Slug:  data.Slug,
	}
}

func FromCoreList(data []blogs.Core) []Blog {
	result := []Blog{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}
