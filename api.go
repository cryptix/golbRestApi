package golbRestApi

import (
	"github.com/cryptix/golbStore"
)

type RestBlogApi struct {
	blogStore golbStore.GolbStorer
}

func NewRestBlogApi(store golbStore.GolbStorer) *RestBlogApi {
	return &RestBlogApi{store}
}
