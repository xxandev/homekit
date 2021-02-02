package homekit

//argToBool -
func argToBool(value interface{}, def bool) bool {
	if val, ok := value.(bool); ok == true {
		return val
	}
	return def
}

//argToInt - int signed, either 32 or 64 bits
func argToInt(value interface{}, def int) int {
	switch val := value.(type) {
	case int:
		return val
	case int8:
		return int(val)
	case int16:
		return int(val)
	case int32:
		return int(val)
	case int64:
		return int(val)
	}
	return def
}

//argToInt32 - int32 signed 32-bit integers (-2147483648 to 2147483647)
func argToInt32(value interface{}, def int32) int32 {
	switch val := value.(type) {
	case int32:
		return val
	case int:
		return int32(val)
	case int8:
		return int32(val)
	case int16:
		return int32(val)
	case int64:
		return int32(val)
	}
	return def
}

//argToInt64 - int64 signed 64-bit integers (-9223372036854775808 to 9223372036854775807)
func argToInt64(value interface{}, def int64) int64 {
	switch val := value.(type) {
	case int64:
		return val
	case int:
		return int64(val)
	case int8:
		return int64(val)
	case int16:
		return int64(val)
	case int32:
		return int64(val)
	}
	return def
}

//argToFloat32 - float32 the set of all IEEE-754 32-bit floating-point numbers
func argToFloat32(value interface{}, def float32) float32 {
	switch val := value.(type) {
	case float32:
		return val
	case float64:
		return float32(val)
	case int:
		return float32(val)
	}
	return def
}

//argToFloat64 - float64 the set of all IEEE-754 64-bit floating-point numbers
func argToFloat64(value interface{}, def float64) float64 {
	switch val := value.(type) {
	case float64:
		return val
	case float32:
		return float64(val)
	case int:
		return float64(val)
	}
	return def
}
