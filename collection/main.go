package main

import (
	"reflect"
)

type collection struct {
	val reflect.Value
	typ reflect.Type
}

func Collect(obj interface{}) *collection {
	if reflect.TypeOf(obj).Kind() != reflect.Slice {
		panic("collect object must is slice")
	}
	return &collection{
		val: reflect.ValueOf(obj),
		typ: reflect.TypeOf(obj).Elem(),
	}
}

func (c *collection) Pluck(key string) *collection {
	var s []interface{}
	for i := 0; i < c.val.Len(); i++ {
		s = append(s, c.val.Index(i).FieldByName(key).Interface())
	}

	return &collection{
		val: reflect.ValueOf(s),
		typ: reflect.TypeOf(s).Elem(),
	}
}

func (c *collection) SliceInt() []int {
	s := make([]int, c.val.Len())
	for i := 0; i < c.val.Len(); i++ {
		s[i] = c.val.Index(i).Interface().(int)
	}

	return s
}

func (c *collection) SliceInt8() []int8 {
	s := make([]int8, c.val.Len())
	for i := 0; i < c.val.Len(); i++ {
		s[i] = c.val.Index(i).Interface().(int8)
	}

	return s
}

func (c *collection) SliceInt16() []int16 {
	s := make([]int16, c.val.Len())
	for i := 0; i < c.val.Len(); i++ {
		s[i] = c.val.Index(i).Interface().(int16)
	}

	return s
}

func (c *collection) SliceInt32() []int32 {
	s := make([]int32, c.val.Len())
	for i := 0; i < c.val.Len(); i++ {
		s[i] = c.val.Index(i).Interface().(int32)
	}

	return s
}

func (c *collection) SliceInt64() []int64 {
	s := make([]int64, c.val.Len())
	for i := 0; i < c.val.Len(); i++ {
		s[i] = c.val.Index(i).Interface().(int64)
	}

	return s
}

func (c *collection) SliceUint() []uint {
	s := make([]uint, c.val.Len())
	for i := 0; i < c.val.Len(); i++ {
		s[i] = c.val.Index(i).Interface().(uint)
	}

	return s
}

func (c *collection) SliceUint8() []uint8 {
	s := make([]uint8, c.val.Len())
	for i := 0; i < c.val.Len(); i++ {
		s[i] = c.val.Index(i).Interface().(uint8)
	}

	return s
}

func (c *collection) SliceUint16() []uint16 {
	s := make([]uint16, c.val.Len())
	for i := 0; i < c.val.Len(); i++ {
		s[i] = c.val.Index(i).Interface().(uint16)
	}

	return s
}

func (c *collection) SliceUint32() []uint32 {
	s := make([]uint32, c.val.Len())
	for i := 0; i < c.val.Len(); i++ {
		s[i] = c.val.Index(i).Interface().(uint32)
	}

	return s
}

func (c *collection) SliceUint64() []uint64 {
	s := make([]uint64, c.val.Len())
	for i := 0; i < c.val.Len(); i++ {
		s[i] = c.val.Index(i).Interface().(uint64)
	}

	return s
}

func (c *collection) KeyByString(key string) map[string]interface{} {
	m := make(map[string]interface{}, c.val.Len())

	for i := 0; i < c.val.Len(); i++ {
		k := c.val.Index(i).FieldByName(key).Interface().(string)
		val := c.val.Index(i).Interface()
		m[k] = val
	}

	return m
}

func (c *collection) KeyByInt(key string) map[int]interface{} {
	m := make(map[int]interface{}, c.val.Len())

	for i := 0; i < c.val.Len(); i++ {
		k := c.val.Index(i).FieldByName(key).Interface().(int)
		val := c.val.Index(i).Interface()
		m[k] = val
	}

	return m
}

func (c *collection) KeyByInt8(key string) map[int8]interface{} {
	m := make(map[int8]interface{}, c.val.Len())

	for i := 0; i < c.val.Len(); i++ {
		k := c.val.Index(i).FieldByName(key).Interface().(int8)
		val := c.val.Index(i).Interface()
		m[k] = val
	}

	return m
}

func (c *collection) KeyByInt16(key string) map[int16]interface{} {
	m := make(map[int16]interface{}, c.val.Len())

	for i := 0; i < c.val.Len(); i++ {
		k := c.val.Index(i).FieldByName(key).Interface().(int16)
		val := c.val.Index(i).Interface()
		m[k] = val
	}

	return m
}

func (c *collection) KeyByInt32(key string) map[int32]interface{} {
	m := make(map[int32]interface{}, c.val.Len())

	for i := 0; i < c.val.Len(); i++ {
		k := c.val.Index(i).FieldByName(key).Interface().(int32)
		val := c.val.Index(i).Interface()
		m[k] = val
	}

	return m
}

func (c *collection) KeyByInt64(key string) map[int64]interface{} {
	m := make(map[int64]interface{}, c.val.Len())

	for i := 0; i < c.val.Len(); i++ {
		k := c.val.Index(i).FieldByName(key).Interface().(int64)
		val := c.val.Index(i).Interface()
		m[k] = val
	}

	return m
}

func (c *collection) KeyByUint(key string) map[uint]interface{} {
	m := make(map[uint]interface{}, c.val.Len())

	for i := 0; i < c.val.Len(); i++ {
		k := c.val.Index(i).FieldByName(key).Interface().(uint)
		val := c.val.Index(i).Interface()
		m[k] = val
	}

	return m
}

func (c *collection) KeyByUint8(key string) map[uint8]interface{} {
	m := make(map[uint8]interface{}, c.val.Len())

	for i := 0; i < c.val.Len(); i++ {
		k := c.val.Index(i).FieldByName(key).Interface().(uint8)
		val := c.val.Index(i).Interface()
		m[k] = val
	}

	return m
}

func (c *collection) KeyByUint16(key string) map[uint16]interface{} {
	m := make(map[uint16]interface{}, c.val.Len())

	for i := 0; i < c.val.Len(); i++ {
		k := c.val.Index(i).FieldByName(key).Interface().(uint16)
		val := c.val.Index(i).Interface()
		m[k] = val
	}

	return m
}

func (c *collection) KeyByUint32(key string) map[uint32]interface{} {
	m := make(map[uint32]interface{}, c.val.Len())

	for i := 0; i < c.val.Len(); i++ {
		k := c.val.Index(i).FieldByName(key).Interface().(uint32)
		val := c.val.Index(i).Interface()
		m[k] = val
	}

	return m
}

func (c *collection) KeyByUint64(key string) map[uint64]interface{} {
	m := make(map[uint64]interface{}, c.val.Len())

	for i := 0; i < c.val.Len(); i++ {
		k := c.val.Index(i).FieldByName(key).Interface().(uint64)
		val := c.val.Index(i).Interface()
		m[k] = val
	}

	return m
}
