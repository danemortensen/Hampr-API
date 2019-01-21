package dbao

import (
   "gopkg.in/mgo.v2"
)

type Dbao struct {
   Server string
   Database string
}

type Session struct {
   session *mgo.Session
}

func NewSession(url string) (*Session, error) {
   session, err := mgo.Dial(url)
   if err != nil {
      return nil, err
   }
   session.SetMode(mgo.Monotonic, true)
   return &Session{session}, err
}

func (s *Session) Close() {
   if s.session != nil {
      s.session.Close()
   }
}
