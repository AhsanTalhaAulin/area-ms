package domain

type Area struct {
	Gym    bool `json:"gym"`
	School bool `json:"school"`
	Store  bool `json:"store"`
	Lab    bool `json:"lab"`
	Office bool `json:"office"`
}

var Areas []Area

// var DistAreas = make(map[string][]int)

var DistAreas = map[string][]int{
	"gym":    {},
	"school": {},
	"store":  {},
	"lab":    {},
	"office": {},
}
