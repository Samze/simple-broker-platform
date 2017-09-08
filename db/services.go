package db

type Services interface {
	Services() ([]*model.Service, error)
	ServiceByGuid(serviceGuid string) (*model.Service, error)
	CreateService(service *model.Service) error
	DeleteService(service *model.Service) error
}
