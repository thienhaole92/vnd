package share

const DefaultPageItems = 50

func Page(val int) int {
	if val <= 0 {
		return 1
	}
	return val
}

func Limit(val int) int {
	if val <= 0 {
		return DefaultPageItems
	}
	return val
}

func Offset(page int, pageSize int) int {
	p := Page(page)
	ps := Limit(pageSize)
	return (p - 1) * ps
}
