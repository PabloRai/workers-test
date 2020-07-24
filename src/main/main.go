package main

import (
	"errors"
	"fmt"
	"github.com/workers-test/src/concurrency"
	"log"
)

func main() {
	fmt.Println("Holis")
	tasks := []*concurrency.Task{
		concurrency.NewTask(wrapFunc(magicFunction, 1)),
		concurrency.NewTask(wrapFunc(magicFunction, 2)),
		concurrency.NewTask(wrapFunc(magicFunction, 3)),
	}

	p := concurrency.NewPool(tasks, 10)
	p.Run()

	var numErrors int
	for _, task := range p.Tasks {
		if task.Err != nil {
			log.Println(task.Err)
			numErrors++
		}
		if numErrors >= 1 {
			log.Println("Too many errors.")
			break
		}
	}
}

func magicFunction(a int) error {
	fmt.Println("Doing ", a)
	return errors.New("hehe")
}

func wrapFunc(f func(a int) error, a int) func() error {
	return func() error {
		return f(a)
	}
}
