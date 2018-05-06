package search

// StringInterface represents the string backed to-be-searched list.
type StringInterface interface {
	Len() int
	At(index int) string
}

// BytesInterface represents the byte slice backed to-be-searched list.
type BytesInterface interface {
	Len() int
	At(index int) []byte
}
