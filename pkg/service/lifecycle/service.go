package lifecycle_service

type Service struct {
}

func NewService() Service {
	return Service{}
}

func (service Service) IsAlive() bool {
	return true
}

func (service Service) IsReady() bool {
	return true
}
