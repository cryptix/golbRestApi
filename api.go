package golbRestApi

import (
	"github.com/cryptix/golbStore"
	"labix.org/v2/mgo"
)

type RestBlog struct {
	blogStore *golbStore.MgoBlog
}

func NewBlogApi(s *mgo.Session, dbName, collName string) *RestBlog {
	return &RestBlog{golbStore.NewMgoBlog(s, &golbStore.Options{dbName, collName})}
}
