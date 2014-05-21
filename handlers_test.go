package restBlog

import (
	"net/http"
	"testing"

	"github.com/rcrowley/go-tigertonic/mocking"
)

func TestPosts(t *testing.T) {
	s, _, posts, err := getPosts(
		mocking.URL(mux, "GET", "/posts"),
		mocking.Header(nil),
		nil,
	)
	if nil != err {
		t.Fatal(err)
	}
	if http.StatusOK != s {
		t.Fatal(s)
	}
	if len(posts) >= 10 { // Merry Christmas!
		t.Fatalf("wrong length. wanted %d, got %d", 10, len(posts))
	}
}
