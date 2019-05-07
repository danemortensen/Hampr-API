package db

import (
    "gopkg.in/mgo.v2"

    "github.com/danemortensen/Hampr-API/pkg"
    "github.com/danemortensen/Hampr-API/pkg/config"
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
    collection := s.DB(c.DbName).C("garment")
    return &GarmentService {
        collection: collection,
    }
}

func (gs *GarmentService) InsertGarment(g *root.Garment) error {
    return gs.collection.Insert(&g)
}
