package root

type Garment struct {
    Name string
    Brand string
}

type GarmentService interface {
    CreateGarment(g *Garment) error
}
