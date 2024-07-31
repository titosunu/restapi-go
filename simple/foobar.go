package simple

type FooBarService struct {
	*FooService
	*BarService
}

func NewFooBarService(foosService *FooService, barService *BarService) *FooBarService {
	return &FooBarService{FooService: foosService, BarService: barService}
}