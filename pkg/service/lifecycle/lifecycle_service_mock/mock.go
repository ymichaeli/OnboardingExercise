package lifecycle_service_mock

import "github.com/stretchr/testify/mock"

type Service struct {
	mock.Mock
}

func (m *Service) IsAlive() bool {
	return true
}

func (m *Service) IsReady() bool {
	args := m.Called()
	return args.Get(0).(bool)
}
