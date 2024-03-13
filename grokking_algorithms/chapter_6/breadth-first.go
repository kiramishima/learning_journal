package main

import (
	"fmt"
)

func search(graph map[string][]string, name string) bool {
	/// todo
	var search_queue = []string{}
	search_queue = append(search_queue, graph[name]...)
	fmt.Printf("%v", search_queue)

	return false
}

func main() {
	var graph = make(map[string][]string)
	graph["you"] = []string{"alice", "bob", "claire"}
	graph["bob"] = []string{"anuj", "peggy"}
	graph["alice"] = []string{"peggy"}
	graph["claire"] = []string{"thom", "jonny"}
	graph["anuj"] = []string{}
	graph["peggy"] = []string{}
	graph["thom"] = []string{}
	graph["jonny"] = []string{}

	_ = search(graph, "you")
}
