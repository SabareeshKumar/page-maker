package app

type enum []int

func (e enum) hasValue(v int) bool {
	for _, value := range e {
		if value == v {
			return true
		}
	}
	return false
}
