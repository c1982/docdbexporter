package experimental

import (
	"fmt"
	"reflect"
)

func autoMetric(a any) {
	t := reflect.TypeOf(a)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		name := field.Tag.Get("prom")
		if name == "" {
			continue
		}

		fmt.Println(name)
	}
}
