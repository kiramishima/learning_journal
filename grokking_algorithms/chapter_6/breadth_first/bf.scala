package main

import scala.collection.mutable.Queue
import scala.collection.mutable.Set

def person_is_seller(name: String): Boolean =
    name.takeRight(1) == "m"

def search(graph: Map[String, List[String]], name: String): Boolean = {
    var search_queue = Queue[String]()
    var elem = graph.getOrElse(name, List.empty)
    search_queue ++= elem
    var searched = Set[String]()

    while(!search_queue.isEmpty) {
        val person = search_queue.dequeue()
        println(person)
        if (!searched.contains(person)) {
            if(person_is_seller(person)) {
                printf("%s is mango seller!", person)
                return true
            } else {
                search_queue ++= graph.getOrElse(person, List.empty)
                searched += person
            }
        }

    }
    return false
}

object Chapter6 {
    def main(args: Array[String]): Unit = {
        var graph = Map(
            "you" -> List("alice", "bob", "claire"),
            "bob" -> List("anuj", "peggy"),
            "alice" -> List("peggy"),
            "claire" -> List("thom", "jonny"),
            "anuj" -> List.empty,
            "peggy" -> List.empty,
            "thom" -> List.empty,
            "jonny" -> List.empty
        )
        search(graph, "you")
        // println(graph)
    }

}