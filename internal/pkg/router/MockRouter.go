package router

import "github.com/stretchr/testify/mock"

type MockRouter struct {
	mock.Mock
}

func (rm MockRouter) Route(dayName string, puzzle int) {
	rm.Called(dayName, puzzle)
}
