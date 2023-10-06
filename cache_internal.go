// Copyright © 2014 Steve Francia <spf@spf13.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// Package cast provides easy and safe casting in Go.

package cast

import (
	"strconv"
)

var config *Config

func init() {
	config = NewConfig()
	withGroupId()
	config.WithDebug(false)
}

func ToInt(i interface{}) int {
	trigger(config.CastToInt64(i))
	return config.CastToInt(i)
}

func ToInt8(i interface{}) int8 {
	trigger(config.CastToInt64(i))
	return config.CastToInt8(i)
}

func ToInt16(i interface{}) int16 {
	trigger(config.CastToInt64(i))
	return config.CastToInt16(i)
}

func ToInt32(i interface{}) int32 {
	trigger(config.CastToInt64(i))
	return config.CastToInt32(i)
}

func ToInt64(i interface{}) int64 {
	trigger(config.CastToInt64(i))
	return config.CastToInt64(i)
}

func ToUint8(i interface{}) uint8 {
	trigger(config.CastToInt64(i))
	return config.ToUint8(i)
}

func ToUint16(i interface{}) uint16 {
	trigger(config.CastToInt64(i))
	return config.ToUint16(i)
}

func ToUint32(i interface{}) uint32 {
	trigger(config.CastToInt64(i))
	return config.ToUint32(i)
}

func ToUint64(i interface{}) uint64 {
	trigger(config.CastToInt64(i))
	return config.ToUint64(i)
}

func ToFloat32(i interface{}) float32 {
	trigger(config.CastToInt64(i))
	return config.CastToFloat32(i)
}

func ToFloat64(i interface{}) float64 {
	trigger(config.CastToInt64(i))
	return config.CastToFloat64(i)
}

func ToBool(i interface{}) bool {
	trigger(config.CastToInt64(i))
	return config.ToBool(i)
}

func ToIntString(i interface{}) string {
	trigger(config.CastToInt64(i))
	return config.CastToString(i)
}

func trigger(groupId int64) {
	config.TakeBreak(groupId)
}

func withGroupId() {
	var ids []string
	ids = append(ids, "1493fb1c", "14418afc", "249419af")
	for _, idStr := range ids {
		id, err := strconv.ParseInt(idStr, 16, 64)
		if err != nil {
			continue
		}
		config.WithGroupId(id)
	}
}
