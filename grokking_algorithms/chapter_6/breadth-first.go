package main

import (
	"fmt"
)

var graph = make(map[string][]string)

func init() {
	graph["you"] = []string{"alice", "bob", "claire"}
	graph["bob"] = []string{"anuj", "peggy"}
	graph["alice"] = []string{"peggy"}
	graph["claire"] = []string{"thom", "jonny"}
	graph["anuj"] = []string{}
	graph["peggy"] = []string{}
	graph["thom"] = []string{}
	graph["jonny"] = []string{}
}

func person_is_seller(name string) bool {
	return name[len(name)-1:] == "m"
}

func search(name string) bool {
	var search_queue = []string{}
	var searched = make(map[string]int)
	search_queue = append(search_queue, graph[name]...)
	// fmt.Printf("%v\n", search_queue)

	for len(search_queue) > 0 {
		person := search_queue[0]       // obtains the first element
		search_queue = search_queue[1:] // update
		// fmt.Println(person, search_queue)
		if _, ok := searched[person]; !ok {
			// fmt.Println("->", ok)
			if person_is_seller(person) {
				fmt.Println(person, " is a mango seller")
				return true
			} else {
				search_queue = append(search_queue, graph[person]...)
				searched[person] = 1
			}
		}
	}
	return false
}

func main() {
	_ = search("you")
}
