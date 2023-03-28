package Repository

type TemporalDB struct {
	Events [][]float64
}

func New() *TemporalDB {
	return &TemporalDB{
		Events: [][]float64{},
	}
}

func (tdb *TemporalDB) InsertEvent(event []float64) {
	tdb.Events = append(tdb.Events, event)
}

func (tdb *TemporalDB) GetEvents() [][]float64 {
	return tdb.Events
}
