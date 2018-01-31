package ydk

import (
	"fmt"
	"github.com/CiscoDevNet/ydk-go/ydk/types"
	"reflect"
)

var topEntityRegistry = make(map[string]reflect.Type)

// RegisterEntity
func RegisterEntity(name string, entity_type reflect.Type) {
	topEntityRegistry[name] = entity_type
}

// GetTopEntity
func GetTopEntity(name string) types.Entity {
	_, ok := topEntityRegistry[name]
	if !ok {
		panic(fmt.Sprintf("Top entity '%s' not registered!", name))
	}
	return reflect.New(topEntityRegistry[name]).Elem().Addr().Interface().(types.Entity)
}
