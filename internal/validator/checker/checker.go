package checker

type Checker struct {
}

func New() Checker {
	return Checker{}
}

func (ch Checker) EventBaseStartFirstAndEndFirst(eventBase, nextEvent []float64) bool {
	return (eventBase[0] < nextEvent[0]) && (eventBase[1] > nextEvent[0])
}

func (ch Checker) EventBaseDoesNotStartFirstButEndFirst(eventBase, nextEvent []float64) bool {
	return (eventBase[0] > nextEvent[0]) && (eventBase[1] < nextEvent[1])
}

func (ch Checker) EventBaseDoesNotStartFirstAndDoesNotEndFirst(eventBase, nextEvent []float64) bool {
	return (eventBase[0] < nextEvent[1]) && (eventBase[1] > nextEvent[1])

}
