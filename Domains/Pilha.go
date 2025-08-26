package main

import (
    "Jose.Aluno/Domains"
)

func main() {
    var stack domains.Stack
    stack.Push("Livro 1")
    stack.Push("Livro 3")
    stack.Show()
    stack.Pop()
    stack.Pop()
    stack.Show()
}
