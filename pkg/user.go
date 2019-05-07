package root

import (

    "gopkg.in/mgo.v2/bson"
)

type User struct {
    Id string
}

type UserService interface {
    FindUser(authId string, user *bson.M) error
}
