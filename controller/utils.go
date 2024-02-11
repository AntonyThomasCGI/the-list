package controller

import (
	"fmt"
	"reflect"

	logger "github.com/sirupsen/logrus"
)

func safeCast(v interface{}, castType reflect.Type) (interface{}, bool) {
	switch castType.Name() {
	case "string":
		castVal, ok := v.(string)
		if !ok {
			return nil, ok
		}
		return castVal, true
	case "int64":
		castVal, ok := v.(int64)
		if !ok {
			return nil, ok
		}
		return castVal, true
	case "bool":
		castVal, ok := v.(bool)
		if !ok {
			return nil, ok
		}
		return castVal, true
	default:
		logger.Error(fmt.Sprintf("Got unexpected type while casting: %s", castType.Name()))
		return nil, false
	}
}
