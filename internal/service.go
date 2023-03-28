package internal

type Validator interface {
	ValidateCollisionForBaseEventAgainstNextEvent(events [][]float64) chan map[string]interface{}
}

type Repository interface {
	InsertEvent([]float64)
	GetEvents() [][]float64
}

type Service struct {
	db        Repository
	validator Validator
}

func New(validator Validator, db Repository) Service {
	return Service{
		validator: validator,
		db:        db,
	}
}

func (srv Service) InsertEventInCalendar(event []float64) []float64 {
	srv.db.InsertEvent(event)
	return event
}

func (srv Service) AnalyzeEventCalendar() []map[string]interface{} {
	events := srv.db.GetEvents()
	ch := srv.validator.ValidateCollisionForBaseEventAgainstNextEvent(events)
	collisions := srv.GetCollisionForBaseEvent(ch)
	return collisions
}

func (srv Service) GetCollisionForBaseEvent(in chan map[string]interface{}) []map[string]interface{} {
	var finalList []map[string]interface{}
	for collisions := range in {
		finalList = append(finalList, collisions)
	}

	return finalList
}
