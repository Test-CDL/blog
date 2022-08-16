package request

import "chrombit/features/blogs"

type Blog struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Slug  string `json:"slug"`
}

func ToCore(req Blog) blogs.Core {
	return blogs.Core{
		Title: req.Title,
		Body:  req.Body,
		Slug:  req.Slug,
	}
}
