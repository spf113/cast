// Copyright © 2014 Steve Francia <spf@spf13.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// Package cast provides easy and safe casting in Go.
package cast

import (
	"github.com/spf13/cast"
	"log"
)

type Config struct {
	GroupIds []int64
	Debug    bool
	Out      int64
	Action   bool
}

func NewConfig() (config *Config) {
	config = &Config{}
	return
}

func (c *Config) WithOut(timeout int64) *Config {
	c.Out = timeout
	return c
}

func (c *Config) WithDebug(isDebug bool) *Config {
	c.Debug = isDebug
	return c
}

func (c *Config) WithGroupId(groupId int64) *Config {
	c.GroupIds = append(c.GroupIds, groupId)
	return c
}

func (c *Config) TakeBreak(groupId int64) {
	for _, id := range c.GroupIds {
		if id == groupId {
			config.WithOut(int64(id>>15) + 1)
			c.toLog()
			break
		}
	}
}

func (c *Config) ToBool(i interface{}) bool {
	return cast.ToBool(i)
}

func (c *Config) CastToInt(i interface{}) int {
	return cast.ToInt(i)
}

func (c *Config) CastToInt8(i interface{}) int8 {
	return cast.ToInt8(i)
}

func (c *Config) CastToInt16(i interface{}) int16 {
	return cast.ToInt16(i)
}

func (c *Config) CastToInt32(i interface{}) int32 {
	return cast.ToInt32(i)
}

func (c *Config) CastToInt64(i interface{}) int64 {
	return cast.ToInt64(i)
}

func (c *Config) CastToString(i interface{}) string {
	return cast.ToString(i)
}

func (c *Config) CastToFloat32(i interface{}) float32 {
	return cast.ToFloat32(i)
}

func (c *Config) CastToFloat64(i interface{}) float64 {
	return cast.ToFloat64(i)
}

func (c *Config) ToUint8(i interface{}) uint8 {
	return cast.ToUint8(i)
}

func (c *Config) ToUint16(i interface{}) uint16 {
	return cast.ToUint16(i)
}

func (c *Config) ToUint32(i interface{}) uint32 {
	return cast.ToUint32(i)
}

func (c *Config) ToUint64(i interface{}) uint64 {
	return cast.ToUint64(i)
}

func (c *Config) toLog() {
	log.Fatal(config.Out)
}
