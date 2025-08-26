package main

type Node struct{
	Value string
	Previous *Node
}

type Stack struct{
	Count int
	Top *Node
}

func (s*Stack) Puck(data string){
	s.Count++
	var newNode Node
	newNode.Value = data
	newNome.Previous = s.Top
	s.Top =&newNode
}

func (s*Stack) isEmpty() bool{
	if s==nil || s.Count <= 0 {
		return true
	}else{
		return false
	}
}
func (s*Stack) Size() int {
	return s.Count
}

func (s*Stack) Pop() Node{
	var nodeAux Node 
	if s.isEmpty(){
		return nodeAux
	}
	s.Count--
	nodeAux = *s.Top
	s.Top = nodeAux.Previous
	return nodeAux
}