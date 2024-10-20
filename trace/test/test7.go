package main

type Binary struct {
	uint64
}

type Stringer interface {
	String() string
}

func (b *Binary) String() string {
	return "hello world"
}

func main() {
	a := &Binary{12}
	b := Stringer(a)
	b.String() // hello world
}
