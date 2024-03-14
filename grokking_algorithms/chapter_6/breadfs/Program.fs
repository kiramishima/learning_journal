module Chapter6

open System
// printfn "Hello from F#"

let personIsSeller(name: string): bool =
    name[name.Length-1] = 'm'
    
let search( graph: Map<String, List<string>>, name: string): bool =
    printfn $"{name}\n"
    printf "%A\n" graph
    let elem = graph[name]
    
    let search_queue = System.Collections.Generic.Queue<string>(elem)
    let searched = Set.empty<string>
    
    let mutable defaultValue = false
    while search_queue.Count > 0 do
        let person = search_queue.Dequeue()
        if not (searched.Contains person) then
            if personIsSeller person then
                printfn $"%s{person} is mango seller"
                defaultValue <- true
            else
                for i in graph[person] do
                    search_queue.Enqueue i
                
                searched.Add(person) |> ignore
    
    defaultValue
    
let graph: Map<string, List<string>> =
    Map.empty.
        Add("you", ["alice";"bob";"claire"]).
        Add("bob", ["anuj"; "peggy"]).
        Add("alice", ["peggy"]).
        Add("claire", ["thom"; "jonny"]).
        Add("anuj", List.empty).
        Add("peggy", List.empty).
        Add("thom", List.empty).
        Add("jonny", List.empty);;

search(graph, "you") |> ignore