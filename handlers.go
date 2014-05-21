package restBlog

import (
	"net/http"
	"net/url"

	"github.com/rcrowley/go-tigertonic"
	"github.com/willnix/tinkerBlog/blog"
)

var (
	mux *tigertonic.TrieServeMux
)

func init() {
	mux = tigertonic.NewTrieServeMux()
}

func getPosts(u *url.URL, header http.Header, req *postsReq) (int, http.Header, []*blog.Entry, error) {
	return 0, nil, nil, nil
}
