package db

import (
    "gopkg.in/mgo.v2"

    "github.com/danemortensen/Hampr-API/pkg/config"
    "gopkg.in/mgo.v2/bson"
    //"log"
)

/**

Use the following names for all service operations:
Insert
Find
Update
Delete

**/

type GarmentService struct {
    collection *mgo.Collection
}

func NewGarmentService(s *mgo.Session, c *config.MongoConfig) *GarmentService {
    collection := s.DB(c.DbName).C("users")
    return &GarmentService {
        collection: collection,
    }
}

func (gs *GarmentService) InsertGarment(authId string, garment *bson.M) error {
    garmentId := (*garment)["id"].(string)
    return gs.collection.UpdateId(authId,
        bson.M{"$set": bson.M{"garments." + garmentId: *garment}})
}
