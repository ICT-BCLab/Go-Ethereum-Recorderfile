package recorder

import "time"

type Cpumem struct {
	OccurTime time.Time `gorm:"index;primarykey"`
	Cpupro    float64   `gorm:"type:float(5,1)"`
	Mempro    float64   `gorm:"type:float(5,1)"`
}

func init() {
	RegisterModel(&Cpumem{})
}
