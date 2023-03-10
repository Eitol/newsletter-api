package newsletter

type Result[T any] struct {
	Total int
	Pages int
	Page  Page[T]
}
