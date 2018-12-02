package menu

import (
	"github.com/cruzzan/AdventOfCode2018/internal/pkg/router"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestMenuCallsRouter(t *testing.T) {
	t.Run("MenuCallsRouterWithArg", testMenuCallsRouterWithArgument )
}

func testMenuCallsRouterWithArgument(t *testing.T)  {
	mr := new(router.MockRouter)
	mr.On("Route", mock.Anything)

	menu := NewMenu(mr)
	menu.Execute()

	mr.AssertExpectations(t)
}
