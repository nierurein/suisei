package author

import "github.com/daniel5u/suisei/domain/author"

type Request struct {
	Name string `json:"name"`
}

func requestToDomain(request Request) author.Domain {
	return author.Domain{
		Name: request.Name,
	}
}
