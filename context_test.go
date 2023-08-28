package golangcontext

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	backgroound := context.Background()
	fmt.Println(backgroound)

	todo := context.TODO()
	fmt.Println(todo)

}

func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")

	contextG := context.WithValue(contextF, "g", "G")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)
	fmt.Println(contextG)

	fmt.Println(contextF.Value("f"))
	fmt.Println(contextF.Value("c"))
	fmt.Println(contextF.Value("b"))
	fmt.Println(contextF.Value("a"))
}

func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1000 * time.Millisecond) // Simulasi Slow
			}

		}

	}()

	return destination
}

func TestContextWithCancel(t *testing.T) {
	fmt.Println("total goroutine ", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)
	destination := CreateCounter(ctx)

	fmt.Println("total goroutine ", runtime.NumGoroutine())
	for n := range destination {
		fmt.Println("Counter =>", n)
		if n == 10 {
			break
		}
	}
	cancel() // Mengirim sinyal cancel ke context

	time.Sleep(2 * time.Second)

	fmt.Println("total goroutine ", runtime.NumGoroutine())
}

func TestContextWithTimeOut(t *testing.T) {
	fmt.Println("total goroutine ", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	destination := CreateCounter(ctx)

	fmt.Println("total goroutine ", runtime.NumGoroutine())
	for n := range destination {
		fmt.Println("Counter =>", n)

	}

	time.Sleep(2 * time.Second)

	fmt.Println("total goroutine ", runtime.NumGoroutine())
}

func TestContextWithDeadline(t *testing.T) {
	fmt.Println("total goroutine ", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))
	defer cancel()

	destination := CreateCounter(ctx)

	fmt.Println("total goroutine ", runtime.NumGoroutine())
	for n := range destination {
		fmt.Println("Counter =>", n)

	}

	time.Sleep(1 * time.Second)

	fmt.Println("total goroutine ", runtime.NumGoroutine())
}
