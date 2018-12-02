package Menu

import (
	"github.com/cruzzan/AdventOfCode2018/internal/pkg/Router"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestMenuCallsRouter(t *testing.T) {
	t.Run("MenuCallsRouterWithArg", testMenuCallsRouterWithArgument )
}

func testMenuCallsRouterWithArgument(t *testing.T)  {
	mr := new(Router.MockRouter)
	mr.On("Route", mock.Anything)

	menu := NewMenu(mr)
	menu.Execute()

	mr.AssertExpectations(t)
}
