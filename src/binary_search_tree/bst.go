package bst

type Node struct {
  left *Node
  right *Node
  value int
}

type bst struct {
  root *Node
  size int;
}
