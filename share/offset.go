package share

const DefaultPageItems = 50

func Page(val int) int {
	if val <= 0 {
		return 1
	}
	return val
}

func Limit(val, def int) int {
	if val <= 0 && def <= 0 {
		return DefaultPageItems
	}

	if val <= 0 && def > 0 {
		return def
	}

	return val
}

func Offset(page, pageSize, defPageSize int) int {
	p := Page(page)
	ps := Limit(pageSize, defPageSize)
	return (p - 1) * ps
}
