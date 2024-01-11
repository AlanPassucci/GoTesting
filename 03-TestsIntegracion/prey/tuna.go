package prey

import "gotesting/03-TestsIntegracion/pkg/storage"

type tuna struct {
	// max speed in m/s
	maxSpeed float64
}

func (t *tuna) GetSpeed() float64 {
	return t.maxSpeed
}

func NewTuna(storage storage.Storage) Prey {
	return &tuna{
		maxSpeed: storage.GetValue("tuna_speed").(float64),
	}
}
