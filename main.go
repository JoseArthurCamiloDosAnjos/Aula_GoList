package main

import "fmt"

type Node struct {
	Value    string
	Previous *Node
}

type Stack struct {
	Count int
	Top   *Node
}

func (s *Stack) Push(data string) {
	s.Count++
	newNode := Node{
		Value:    data,
		Previous: s.Top,
	}
	s.Top = &newNode
}

func (s *Stack) isEmpty() bool {
	return s == nil || s.Count <= 0
}

func (s *Stack) Size() int {
	return s.Count
}

func (s *Stack) Pop() (string, bool) {
	if s.isEmpty() {
		return "", false
	}
	s.Count--
	value := s.Top.Value
	s.Top = s.Top.Previous
	return value, true
}

func (s *Stack) Show() {
	if s.isEmpty() {
		fmt.Println("Pilha vazia!")
		return
	}
	nodeExibir := s.Top
	for nodeExibir != nil {
		fmt.Println("Node =", nodeExibir.Value)
		nodeExibir = nodeExibir.Previous
	}
}
