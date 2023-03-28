package validator

import "sync"

type Checker interface {
	EventBaseStartFirstAndEndFirst(eventBase, nextEvent []float64) bool
	EventBaseDoesNotStartFirstButEndFirst(eventBase, nextEvent []float64) bool
	EventBaseDoesNotStartFirstAndDoesNotEndFirst(eventBase, nextEvent []float64) bool
}

type Validator struct {
	checker Checker
}

func New(checker Checker) Validator {
	return Validator{
		checker: checker,
	}
}

func (va Validator) ValidateCollisionForBaseEventAgainstNextEvent(events [][]float64) chan map[string]interface{} {
	out := make(chan map[string]interface{})
	wg := sync.WaitGroup{}
	wg.Add(len(events))
	for i := 0; i < len(events); i++ {
		go func(i int) {
			defer wg.Done()
			var totalCollisionsForEvents [][]float64
			for j := 0; j < len(events); j++ {
				if va.ValidateIfEventsAreTheSame(&events[i], &events[j]) {
					continue
				}

				result := va.ValidateCollision(events[i], events[j])
				if len(result) > 0 {
					totalCollisionsForEvents = append(totalCollisionsForEvents, result)
				}
			}

			out <- map[string]interface{}{
				"baseEvent":       events[i],
				"eventCollisions": totalCollisionsForEvents,
			}
		}(i)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func (va Validator) ValidateIfEventsAreTheSame(baseEvent *[]float64, nextEvent *[]float64) bool {
	return baseEvent == nextEvent
}

func (va Validator) ValidateCollision(eventBase, nextEvent []float64) []float64 {
	if va.checker.EventBaseStartFirstAndEndFirst(eventBase, nextEvent) {
		return nextEvent
	}

	if va.checker.EventBaseDoesNotStartFirstButEndFirst(eventBase, nextEvent) {
		return nextEvent
	}

	if va.checker.EventBaseDoesNotStartFirstAndDoesNotEndFirst(eventBase, nextEvent) {
		return nextEvent
	}

	return nil
}
