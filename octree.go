package main

const max_items_in_leaf = 9

type octree struct {
	items    []interface{}
	subtrees [2][2][2]*octree
}
