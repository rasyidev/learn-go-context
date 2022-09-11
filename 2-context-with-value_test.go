package learngocontext

import (
	"context"
	"fmt"
	"testing"
)

func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "a", "A")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)

	// fmt.Println(contextA.Value())
	fmt.Println(contextF.Value("f"))
	fmt.Println(contextF.Value("c"))
	fmt.Println(contextA.Value("b"))
}

/*
=== RUN   TestContextWithValue
context.Background
context.Background.WithValue(type string, val A)
context.Background.WithValue(type string, val C)
context.Background.WithValue(type string, val A).WithValue(type string, val D)
context.Background.WithValue(type string, val A).WithValue(type string, val E)
context.Background.WithValue(type string, val C).WithValue(type string, val F)
F
C
<nil>
--- PASS: TestContextWithValue (0.00s)
PASS
ok      learn-go-context        0.047s
*/
