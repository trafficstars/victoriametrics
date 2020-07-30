package metrics

import "strings"

func toStringTags(t AnyTags) string {
	if t == nil || t.Len() == 0 {
		return defaultTags
	}
	if len(defaultTags) > 0 {
		return t.String() + "," + defaultTags
	}
	return t.String()
}

func prepareKey(k string) string {
	return "metrics_" + strings.ReplaceAll(k, ".", "_")
}

func toInt64(in interface{}) (int64, bool) {
	switch v := in.(type) {
	case uint8:
		return int64(v), true
	case int8:
		return int64(v), true
	case uint16:
		return int64(v), true
	case int16:
		return int64(v), true
	case uint32:
		return int64(v), true
	case int32:
		return int64(v), true
	case uint64:
		return int64(v), true
	case int64:
		return v, true
	case uint:
		return int64(v), true
	case int:
		return int64(v), true
	}

	return 0, false
}
