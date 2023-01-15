package dto

import "net/url"

type FindTweetParameter struct {
	Id    string
	Props []string
}

func CreateFindTweetParameter(p string, qms url.Values) FindTweetParameter {
	parameter := FindTweetParameter{}

	parameter.Id = p

	if qms["props"] != nil {
		parameter.Props = qms["props"]
	}

	return parameter
}
