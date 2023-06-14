package parser

import (
	"strconv"
)

func StringToBool(value string) (bool, error) {
	return strconv.ParseBool(value)
}

func StringToBoolPanic(value string) bool {
	result, err := StringToBool(value)

	if err != nil {
		panic(err)
	}

	return result
}

func StringToFloat32(value string) (float32, error) {
	result, err := strconv.ParseFloat(value, 32)

	if err != nil {
		return 0, err
	}

	return float32(result), nil
}

func StringToFloat32Panic(value string) float32 {
	result, err := StringToFloat32(value)

	if err != nil {
		panic(err)
	}

	return result
}

func StringToFloat64(value string) (float64, error) {
	return strconv.ParseFloat(value, 64)
}

func StringToFloat64Panic(value string) float64 {
	result, err := StringToFloat64(value)

	if err != nil {
		panic(err)
	}

	return result
}

func StringToInt(value string) (int, error) {
	result, err := strconv.ParseInt(value, 0, 0)

	if err != nil {
		return 0, err
	}

	return int(result), nil
}

func StringToIntPanic(value string) int {
	result, err := StringToInt(value)

	if err != nil {
		panic(err)
	}

	return result
}

func StringToInt8(value string) (int8, error) {
	result, err := strconv.ParseInt(value, 0, 8)

	if err != nil {
		return 0, err
	}

	return int8(result), nil
}

func StringToInt8Panic(value string) int8 {
	result, err := StringToInt8(value)

	if err != nil {
		panic(err)
	}

	return result
}

func StringToInt16(value string) (int16, error) {
	result, err := strconv.ParseInt(value, 0, 16)

	if err != nil {
		return 0, err
	}

	return int16(result), nil
}

func StringToInt16Panic(value string) int16 {
	result, err := StringToInt16(value)

	if err != nil {
		panic(err)
	}

	return result
}

func StringToInt32(value string) (int32, error) {
	result, err := strconv.ParseInt(value, 0, 32)

	if err != nil {
		return 0, err
	}

	return int32(result), nil
}

func StringToInt32Panic(value string) int32 {
	result, err := StringToInt32(value)

	if err != nil {
		panic(err)
	}

	return result
}

func StringToInt64(value string) (int64, error) {
	return strconv.ParseInt(value, 0, 64)
}

func StringToInt64Panic(value string) int64 {
	result, err := StringToInt64(value)

	if err != nil {
		panic(err)
	}

	return result
}

func StringToUint(value string) (uint, error) {
	result, err := strconv.ParseUint(value, 0, 0)

	if err != nil {
		return 0, err
	}

	return uint(result), nil
}

func StringToUintPanic(value string) uint {
	result, err := StringToUint(value)

	if err != nil {
		panic(err)
	}

	return result
}

func StringToUint8(value string) (uint8, error) {
	result, err := strconv.ParseUint(value, 0, 8)

	if err != nil {
		return 0, err
	}

	return uint8(result), nil
}

func StringToUint8Panic(value string) uint8 {
	result, err := StringToUint8(value)

	if err != nil {
		panic(err)
	}

	return result
}

func StringToUint16(value string) (uint16, error) {
	result, err := strconv.ParseUint(value, 0, 16)

	if err != nil {
		return 0, err
	}

	return uint16(result), nil
}

func StringToUint16Panic(value string) uint16 {
	result, err := StringToUint16(value)

	if err != nil {
		panic(err)
	}

	return result
}

func StringToUint32(value string) (uint32, error) {
	result, err := strconv.ParseUint(value, 0, 32)

	if err != nil {
		return 0, err
	}

	return uint32(result), nil
}

func StringToUint32Panic(value string) uint32 {
	result, err := StringToUint32(value)

	if err != nil {
		panic(err)
	}

	return result
}

func StringToUint64(value string) (uint64, error) {
	return strconv.ParseUint(value, 0, 64)
}

func StringToUint64Panic(value string) uint64 {
	result, err := StringToUint64(value)

	if err != nil {
		panic(err)
	}

	return result
}
