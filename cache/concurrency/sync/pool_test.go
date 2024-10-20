package sync

import (
	"fmt"
	"sync"
	"testing"
)

func TestPool(t *testing.T) {

	p := sync.Pool{
		New: func() interface{} {
			fmt.Println("create")
			return "heelo"
		},
	}

	str := p.Get().(string)

	t.Log(str)

	p.Put(str)

	str = p.Get().(string)
	t.Log(str)
	p.Put(str)

}
