package main

type Edge struct {
	from   int
	target int
	weight int
}

type Graph map[int][]Edge

type Vertex struct {
	from   int
	to     int
	weight int
}

type Vertices []Vertex
