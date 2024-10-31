package event

type SeriesElement struct {
	Event    Event
	RootStep int
}

type RootSeries struct {
	Events []Event
	Subs   map[string]SubSeries
}

type SubSeries struct {
	Events []SeriesElement
}
