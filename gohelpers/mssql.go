package gohelpers

import (
	"encoding/base64"
	"math"
)

// MssqlDecimalToFloat64 Converts a Mssql decimal base 64 binary representation
// to float64
func MssqlDecimalToFloat64(val string, mantissa uint8) (float64, error) {
	bytesVal := make([]byte, base64.StdEncoding.DecodedLen(len(val)))
	_, err := base64.StdEncoding.Decode(bytesVal, []byte(val))
	if err != nil {
		return 0, err
	}
	// Calculate integer value
	var intval int64
	for _, b := range bytesVal {
		intval = intval*256 + int64(b)
	}
	// If first bit is 1 -> Calculate 2's complement
	if bytesVal[0]&(1<<7) != 0 {
		max := 1 << (len(bytesVal) * 8)
		intval = intval - int64(max)
	}
	return float64(intval) / math.Pow10(int(mantissa)), nil
}
