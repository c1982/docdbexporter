package experimental

import (
	"testing"
)

type TestModel struct {
	Ns          string  `prom:"namespace,desc="namespace of the collection"`
	Count       int64   `prom:"count"`
	Size        float64 `prom:"size"`
	AvgObjSize  int64   `prom:"avg_object_size"`
	StorageSize float64 `prom:"storageSize"`
	Capped      bool    `prom:"capped"`
}

func TestGetNamespace(t *testing.T) {
	var tmodel = TestModel{
		Ns:          "test",
		Count:       100,
		Size:        200.5,
		AvgObjSize:  50,
		StorageSize: 300.5,
		Capped:      true,
	}

	autoMetric(tmodel)
}
