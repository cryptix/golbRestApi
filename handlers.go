package restBlog

import (
	"net/http"
	"net/url"

	"github.com/willnix/tinkerBlog/blog"
)

func (r RestBlog) blogList(url *url.URL, header http.Header, req *BlogRequest) (int, http.Header, []*blog.Entry, error) {

	entries, err := r.blogStore.LatestEntries()
	if err != nil {
		return http.StatusInternalServerError, nil, nil, err
	}
	return http.StatusOK, nil, entries, nil
}

func (r RestBlog) blogPost(url *url.URL, header http.Header, req *BlogRequest) (int, http.Header, *blog.Entry, error) {

	e, err := r.blogStore.FindById(url.Query().Get("id"))
	switch {
	case err == nil:
		return http.StatusOK, nil, e, nil

	case err == blog.ErrEntryNotFound:
		return http.StatusNotFound, nil, nil, blog.ErrEntryNotFound

	default:
		return http.StatusInternalServerError, nil, nil, err
	}
}
