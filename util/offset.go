package util

func Offset(page, pageSize, defPageSize int) int {
	p := Page(page)
	ps := Limit(pageSize, defPageSize)
	return (p - 1) * ps
}
