package root

type Garment struct {
    Name string
    Brand string
}

type GarmentService interface {
    InsertGarment(g *Garment) error
}
