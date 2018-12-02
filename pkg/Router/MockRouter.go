package Router

import "github.com/stretchr/testify/mock"

type MockRouter struct {
	mock.Mock
}

func (rm MockRouter) Route(day string) {
	rm.Called(day)
}
