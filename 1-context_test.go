package learngocontext

import (
	"context"
	"fmt"
	"testing"
)

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

/*
=== RUN   TestContext
context.Background
context.TODO
--- PASS: TestContext (0.00s)
PASS
ok      learn-go-context        0.063s
*/
