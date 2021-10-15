package publisher

import "github.com/daniel5u/suisei/domain/publisher"

type Request struct {
	Name string `json:"name"`
}

func requestToDomain(request Request) publisher.Domain {
	return publisher.Domain{
		Name: request.Name,
	}
}
