package golbRestApi

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/cryptix/golbStore"
	"github.com/cryptix/golbStore/mgo"

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

func setup() (mux *tigertonic.TrieServeMux, api *RestBlogApi) {
	mgoSession, err := mgo.Dial(fmt.Sprintf("%s/%s", dbHost, dbName))
	if err != nil {
		panic(err)
	}

	api = NewRestBlogApi(golbStoreMgo.NewStore(mgoSession, &golbStoreMgo.Options{dbName, dbColl}))

	mux = tigertonic.NewTrieServeMux()
	mux.Handle("GET", "/blog", tigertonic.Marshaled(api.List))
	mux.Handle("GET", "/blog/{{id}}", tigertonic.Marshaled(api.GetPost))

	return
}

func TestList(t *testing.T) {
	var (
		mux *tigertonic.TrieServeMux
		api *RestBlogApi
	)

	Convey("List sanity", t, func() {
		mux, api = setup()

		code, headers, _, err := api.List(
			mocking.URL(mux, "GET", "/blog"),
			mocking.Header(nil),
			&ListRequest{10, false},
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
		api *RestBlogApi
	)
	Convey("GetPost sanity", t, func() {
		mux, api = setup()

		code, headers, _, err := api.GetPost(
			mocking.URL(mux, "GET", "/blog/536797b7b8fed507ae000002"),
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
			So(err, ShouldEqual, golbStore.ErrEntryNotFound)
			So(code, ShouldEqual, http.StatusNotFound)
		})

		Convey("no headers are set", func() {
			So(headers, ShouldBeNil)
		})

	})
}
