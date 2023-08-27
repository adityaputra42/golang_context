package golangcontext

import (
	"context"
	"fmt"
	"testing"
)

func TestContext(t *testing.T) {
	backgroound := context.Background()
	fmt.Println(backgroound)

	todo := context.TODO()
	fmt.Println(todo)

}
