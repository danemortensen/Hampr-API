package root

import (
    "gopkg.in/mgo.v2/bson"
)

type GarmentService interface {
    InsertGarment(authId string, garment *bson.M) error
}
