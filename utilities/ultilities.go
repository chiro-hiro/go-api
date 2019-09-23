package utilities

import (
	"encoding/hex"
	"reflect"
)

//SafeAccess safe access any slice
func SafeAccess(data interface{}, index int) (val reflect.Value) {
	if data == nil || reflect.TypeOf(data).Kind() != reflect.Slice {
		return
	}
	tmp := reflect.ValueOf(data)
	tmpLen := tmp.Len()
	if tmpLen > 0 && index < 0 {
		val = tmp.Index(tmpLen - 1)
	} else if tmpLen > 0 && index < tmpLen {
		val = tmp.Index(index)
	}
	return
}

//First get first element of slice
func First(data interface{}) (val reflect.Value) {
	return SafeAccess(data, 0)
}

//Last get las element of slice
func Last(data interface{}) (val reflect.Value) {
	return SafeAccess(data, -1)
}

//Pad Add leading zero to invalid string
func Pad(hexString string) string {
	if len(hexString)%2 == 1 {
		return "0" + hexString
	}
	return hexString
}

//Hex convert bytes to hex string
func Hex(input []byte) string {
	return hex.EncodeToString(input)
}

//UnHex string to bytes
func UnHex(input string) (result []byte) {
	result, _ = hex.DecodeString(Pad(input))
	return
}
