package goHelpers

import "time"

func GetBoolAddress(value bool) *bool {
	return &value
}

func GetStringAddress(value string) *string {
	return &value
}

func GetIntAddress(value int) *int {
	return &value
}

func GetInt32Address(value int32) *int32 {
	return &value
}

func GetInt64Address(value int64) *int64 {
	return &value
}

func GetFloat32Address(value float32) *float32 {
	return &value
}

func GetFloat64Address(value float64) *float64 {
	return &value
}

func GetTimeAddress(value time.Time) *time.Time {
	return &value
}
