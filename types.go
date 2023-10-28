// types.go
//
// Author: Alperen Cantez <alperen.cantez@outlook.com>
// Created on: 23-10-2023

package unicornia

type Number interface {
	int | int16 | int32 | int64 | int8 | float32 | float64 | uint | uint16 | uint32 | uint64 | uint8
}
