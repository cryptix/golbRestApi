package golbRestApi

import (
	"net/http"
	"net/url"

	"github.com/cryptix/golbStore"
)

type ListRequest struct{ WithText bool }

func (r RestBlog) List(url *url.URL, header http.Header, req *ListRequest) (code int, h http.Header, entries []*golbStore.Entry, err error) {

	entries, err = r.blogStore.LatestEntries(req.WithText)
	if err != nil {
		return http.StatusInternalServerError, nil, nil, err
	}
	return http.StatusOK, nil, entries, nil
}

type GetPostRequest struct{}

func (r RestBlog) GetPost(url *url.URL, header http.Header, req *GetPostRequest) (int, http.Header, *golbStore.Entry, error) {

	e, err := r.blogStore.FindById(url.Query().Get("id"))
	switch {
	case err == nil:
		return http.StatusOK, nil, e, nil

	case err == golbStore.ErrEntryNotFound:
		return http.StatusNotFound, nil, nil, golbStore.ErrEntryNotFound

	default:
		return http.StatusInternalServerError, nil, nil, err
	}
}
