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

type OutfitService struct {
    collection *mgo.Collection
}

func NewOutfitService(s *mgo.Session, c *config.MongoConfig) *OutfitService {
    collection := s.DB(c.DbName).C("users")
    return &OutfitService {
        collection: collection,
    }
}

func (os *OutfitService) InsertOutfit(authId string, outfit *bson.M) error {
    outfitId := (*outfit)["id"].(string)
    return os.collection.UpdateId(authId,
        bson.M{"$set": bson.M{"outfits." + outfitId: *outfit}})
}

func (os *OutfitService) DeleteOutfit(authId string, outfitId string) error {
    return os.collection.UpdateId(authId,
        bson.M{"$unset": bson.M{"outfits." + outfitId: ""}})
}
