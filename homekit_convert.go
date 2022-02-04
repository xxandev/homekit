package homekit

func tob(value interface{}, def bool) bool {
	if val, ok := value.(bool); ok {
		return val
	}
	return def
}

func toi(value interface{}, def int) int {
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

func toi32(value interface{}, def int32) int32 {
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

func toi64(value interface{}, def int64) int64 {
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

func tof32(value interface{}, def float32) float32 {
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

func tof64(value interface{}, def float64) float64 {
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
