import math

graph = {}
graph["start"] = {}
graph["start"]["a"] = 6
graph["start"]["b"] = 2
graph["a"] = {}
graph["a"]["fin"] = 1
graph["b"] = {}
graph["b"]["a"] = 3
graph["b"]["fin"] = 5
graph["fin"] = {} # The Finish node doesn’t have any out-neighbors.

infinity = math.inf
# Cost hash table
costs = {}
costs["a"] = 6
costs["b"] = 2
costs["fin"] = infinity

# Parents htable
parents = {}
parents["a"] = "start"
parents["b"] = "start"
parents["fin"] = None

processed = set()

def find_lowest_cost_node(costs):
    lowest_cost = math.inf
    lowest_cost_node = None
    for node in costs: # <- Goes through each node
        cost = costs[node]
        if cost < lowest_cost and node not in processed: # <- If it’s the lowest cost so far and hasn’t been processed yet
            lowest_cost = cost # <- sets it as the new lowest-cost node
            lowest_cost_node = node
    return lowest_cost_node

node = find_lowest_cost_node(costs) # <- Finds the lowest-cost node that you haven’t processed yet
while node is not None: # <- If you’ve processed all the nodes, this while loop is done.
    cost = costs[node]
    neighbors = graph[node]
    for n in neighbors.keys(): # <- Goes through all the out-neighbors of this node
        new_cost = cost + neighbors[n]
        if costs[n] > new_cost: # If it’s cheaper to get to this out-neighbor by going through this node . . .
            costs[n] = new_cost # <- . . . updates the cost for the neighbor
            parents[n] = node # <- This node becomes the new parent for this out-neighbor
    processed.add(node) # <- Marks the node as processed
    node = find_lowest_cost_node(costs) # <- Finds the next node to process and loops

print(costs)
print(parents)
