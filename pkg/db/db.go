package db

import (
    "log"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "github.com/danemortensen/Hampr-API/pkg/config"

)

type Session struct {
    session *mgo.Session
    config *config.MongoConfig
}

func NewSession(config *config.MongoConfig) *Session {
    session, err := mgo.Dial(config.Ip)
    if err != nil {
        log.Fatalf("Unable to connect to database at %s\n", config.Ip)
    }
    s := &Session {
        session: session,
        config: config,
    }
    return s
}

func (s *Session) Find(collection string, query bson.M, result *bson.M) {
    session := s.session.Copy()
    defer session.Close()
    c := session.DB(s.config.DbName).C(collection)
    err := c.Find(query).One(result)
    if err != nil {
        log.Print(err)
    }
}

func (s *Session) Close() {
    if (s.session != nil) {
        s.session.Close()
    }
}
