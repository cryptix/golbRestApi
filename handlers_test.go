package restBlog

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/willnix/tinkerBlog/blog"

	"github.com/rcrowley/go-tigertonic"
	"github.com/rcrowley/go-tigertonic/mocking"
	. "github.com/smartystreets/goconvey/convey"
	"labix.org/v2/mgo"
)

const (
	dbHost = "localhost"
	dbName = "blog"
	dbColl = "blogEntries"
)

func setup() (mux *tigertonic.TrieServeMux, api *RestBlog) {
	mgoSession, err := mgo.Dial(fmt.Sprintf("%s/%s", dbHost, dbName))
	if err != nil {
		panic(err)
	}
	api = NewMgoBlog(mgoSession, dbName, dbColl)

	mux = tigertonic.NewTrieServeMux()
	mux.Handle("GET", "/blog", tigertonic.Marshaled(api.blogList))
	mux.Handle("GET", "/blog/{{id}}", tigertonic.Marshaled(api.blogPost))

	return
}

func TestBlogList(t *testing.T) {
	var (
		mux *tigertonic.TrieServeMux
		api *RestBlog
	)

	Convey("blogList sanity", t, func() {
		mux, api = setup()

		code, headers, _, err := api.blogList(
			mocking.URL(mux, "GET", "/blog"),
			mocking.Header(nil),
			nil,
		)

		Convey("it returns ok", func() {
			So(err, ShouldBeNil)
			So(code, ShouldEqual, http.StatusOK)
		})

		Convey("no headers are set", func() {
			So(headers, ShouldBeNil)
		})

	})
}

func TestBlogPost(t *testing.T) {
	var (
		mux *tigertonic.TrieServeMux
		api *RestBlog
	)
	Convey("blogPost sanity", t, func() {
		mux, api = setup()

		code, headers, _, err := api.blogPost(
			mocking.URL(mux, "GET", "/blog/536894f2b8fed518e5000001"),
			mocking.Header(nil),
			nil,
		)

		Convey("it returns ok", func() {
			So(err, ShouldBeNil)
			So(code, ShouldEqual, http.StatusOK)
		})

		Convey("no headers are set", func() {
			So(headers, ShouldBeNil)
		})

	})

	Convey("blogPost not found", t, func() {
		mux, api = setup()

		code, headers, _, err := api.blogPost(
			mocking.URL(mux, "GET", "/blog/5375e499b8fed50f4f000001"),
			mocking.Header(nil),
			nil,
		)

		Convey("it returns 404", func() {
			So(err, ShouldEqual, blog.ErrEntryNotFound)
			So(code, ShouldEqual, http.StatusNotFound)
		})

		Convey("no headers are set", func() {
			So(headers, ShouldBeNil)
		})

	})
}
