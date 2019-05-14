package root

import (
    "gopkg.in/mgo.v2/bson"
)

type OutfitService interface {
    InsertOutfit(authId string, outfit *bson.M) error
    DeleteOutfit(authId string, outfitId string) error
}
