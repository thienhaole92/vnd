package util

func Limit(val, def int) int {
	if val <= 0 && def <= 0 {
		return DEFAULT_PAGE_ITEMS
	}

	if val <= 0 && def > 0 {
		return def
	}

	return val
}
