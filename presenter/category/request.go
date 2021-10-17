package category

import "github.com/daniel5u/suisei/domain/category"

type Request struct {
	Name string `json:"name"`
}

func requestToDomain(request Request) category.Domain {
	return category.Domain{
		Name: request.Name,
	}
}
