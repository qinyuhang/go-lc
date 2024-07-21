package runner

import (
    "github.com/stretchr/testify/assert"
    "reflect"
    "testing"
)

type ResultType int
type FuncType interface{}

const (
    IntVal = iota
    StringVal
    SliceInt
    SliceString
)

func Run(t *testing.T, funcIns FuncType, args []interface{}, exceptResult interface{}, expectResultType ResultType) {
    funcValue := reflect.ValueOf(funcIns)
    reflectArgs := make([]reflect.Value, len(args))
    for i, arg := range args {
        reflectArgs[i] = reflect.ValueOf(arg)
    }
    result := funcValue.Call(reflectArgs)
    switch expectResultType {
    case IntVal:
        assert.Equal(t, exceptResult.(int), result[0].Interface(), "")
        break
    case StringVal:
        assert.Equal(t, exceptResult.(string), result[0].Interface(), "")
        break
    case SliceInt:
        for _, v := range result {
            assert.Equal(t, v.Interface(), exceptResult.([]int))
        }
        break
    case SliceString:
        for _, v := range result {
            assert.Equal(t, v.Interface(), exceptResult.([]string))
        }
        break
    }
}
