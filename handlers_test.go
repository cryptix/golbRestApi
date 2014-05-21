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
	mux.Handle("GET", "/blog", tigertonic.Marshaled(api.List))
	mux.Handle("GET", "/blog/{{id}}", tigertonic.Marshaled(api.GetPost))

	return
}

func TestList(t *testing.T) {
	var (
		mux *tigertonic.TrieServeMux
		api *RestBlog
	)

	Convey("List sanity", t, func() {
		mux, api = setup()

		code, headers, _, err := api.List(
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

func TestGetPost(t *testing.T) {
	var (
		mux *tigertonic.TrieServeMux
		api *RestBlog
	)
	Convey("GetPost sanity", t, func() {
		mux, api = setup()

		code, headers, _, err := api.GetPost(
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

	Convey("GetPost not found", t, func() {
		mux, api = setup()

		code, headers, _, err := api.GetPost(
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
