package db

import (
    "log"

    "gopkg.in/mgo.v2"

    "github.com/danemortensen/Hampr-API/pkg/config"
)

type Session struct {
    session *mgo.Session
}

func NewSession(config *config.MongoConfig) *Session {
    session, err := mgo.Dial(config.Ip)
    if err != nil {
        log.Fatalf("Unable to connect to database at %s\n", config.Ip)
    }
    s := &Session {
        session: session,
    }
    return s
}

func (s *Session) Copy() *mgo.Session {
    return s.session.Copy()
}

func (s *Session) Close() {
    if (s.session != nil) {
        s.session.Close()
    }
}

// func (m *MongoConfig) Insert() error {
//     err :=
// }
