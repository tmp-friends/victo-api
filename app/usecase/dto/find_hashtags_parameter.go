package dto

import (
	"net/url"
	"strconv"
)

type FindHashtagsParameter struct {
	Limit  int
	Offset int
	Props  []string
}

func Create(qms url.Values) FindHashtagsParameter {
	parameter := FindHashtagsParameter{}

	if qms["limit"] != nil {
		limit, _ := strconv.Atoi(qms["limit"][0])
		parameter.Limit = limit
	}
	if qms["offset"] != nil {
		offset, _ := strconv.Atoi(qms["offset"][0])
		parameter.Offset = offset
	}
	if qms["props"] != nil {
		parameter.Props = qms["props"]
	}

	return parameter
}
