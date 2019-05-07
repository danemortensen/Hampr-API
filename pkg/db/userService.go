package db

import (
    "gopkg.in/mgo.v2"

    "github.com/danemortensen/Hampr-API/pkg/config"

    "gopkg.in/mgo.v2/bson"
)

/**

Use the following names for all service operations:
Insert
Find
Update
Delete

**/

type UserService struct {
    collection *mgo.Collection
}

func NewUserService(s *mgo.Session, c *config.MongoConfig) *UserService {
    collection := s.DB(c.DbName).C("users")
    return &UserService {
        collection: collection,
    }
}

func (us *UserService) FindUser(authId string, user *bson.M) error {
    return us.collection.FindId(authId).One(user)
}

func (us *UserService) InsertUser(authId string) error {

    user := bson.M{"_id": authId, "outfits": bson.M{}, "garments": bson.M{}, "score": 0.0}
    // user["_id"] = authId
    // user["outfits"] = nil
    // user["garments"] = nil
    // user["score"] = 0.0
    return us.collection.Insert(user)
}
