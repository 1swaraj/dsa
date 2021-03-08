package main

type Vertex struct {
	target int
}

type tree map[int][]Vertex

var visited map[int]bool