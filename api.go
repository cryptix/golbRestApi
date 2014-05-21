package restBlog

import (
	"github.com/willnix/tinkerBlog/blog"
	"labix.org/v2/mgo"
)

type BlogRequest struct{}

type RestBlog struct {
	blogStore blog.Blogger
}

func NewMgoBlog(s *mgo.Session, dbName, collName string) *RestBlog {
	return &RestBlog{blog.NewMgoBlog(s, &blog.Options{dbName, collName})}
}
