package lib

type data struct {
	X int
	Y int
}

func NewData() *data {
	return new(data)
}
