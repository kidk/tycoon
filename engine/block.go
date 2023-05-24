package engine

type Block struct {
	x int
	y int

	Type Type
}

type Type struct {
	Name string
}
