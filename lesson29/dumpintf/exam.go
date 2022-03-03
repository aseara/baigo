package dumpintf

type T int

func (T) Error() string {
	return "bad error"
}
