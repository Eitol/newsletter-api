package newsletter

type Page[T any] struct {
	Number   int
	Elements []T
}
