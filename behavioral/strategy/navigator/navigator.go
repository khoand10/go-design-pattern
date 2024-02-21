package navigator

type INavigator interface {
	IBuildRouter
}

type IBuildRouter interface {
	BuildRouter()
}

type navigator struct {
	Navigator INavigator
}

func NewNavigator(navigatorImpl INavigator) *navigator {
	return &navigator{
		Navigator: navigatorImpl,
	}
}
