package main

import (
    "flag"
    "todo"
    "os"
    "fmt"
)

const (
    todoFile = ".todo.json"
)

func main() {
    add := flag.Bool("add", false, "add a new todo")
    flag.Parse()

    todos := &todo.Todos{}
    if err := todos.Load(todoFile); err != nil {
        fmt.FPrintln(os.Stderr, err.Error())
        os.Exit()
    }
}
