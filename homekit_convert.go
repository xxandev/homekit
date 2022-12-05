package homekit

func tob(v any, def bool) bool {
	if r, ok := v.(bool); ok {
		return r
	}
	return def
}

func toi(v any, def int) int {
	switch r := v.(type) {
	case int:
		return r
	case int8:
		return int(r)
	case int16:
		return int(r)
	case int32:
		return int(r)
	case int64:
		return int(r)
	}
	return def
}

func tof64(v any, def float64) float64 {
	switch r := v.(type) {
	case float64:
		return r
	case float32:
		return float64(r)
	case int:
		return float64(r)
	}
	return def
}
