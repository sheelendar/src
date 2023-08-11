package main

import "fmt"

type Node struct {
	n   int
	adj []int
}

func main() {
	//numCourses := 2
	//prerequisites := [][]int{{1, 0}, {0, 1}}

	numCourses := 2
	prerequisites := [][]int{{0, 1}}

	//numCourses := 4
	//prerequisites := [][]int{{2, 0}, {1, 0}, {3, 1}, {3, 2}, {1, 3}}

	//numCourses := 3
	//prerequisites := [][]int{{1, 0}, {2, 1}, {3, 2}}
	fmt.Print("using graph as map: ")
	//CourseComplete1(numCourses, prerequisites)
	fmt.Println()
	fmt.Print("using graph as matrix: ")
	CourseComplete2(numCourses, prerequisites)

}
func CourseComplete2(numCourses int, prerequisites [][]int) {

	// visited to check if verdict is visited or not if visited then no
	// need to go ahead its already solved
	visited := make([]bool, numCourses+1)

	// is used to check if verdict is in same path then it has loop if it's in other path then don't have loop.
	isSamePath := make([]bool, numCourses+1)

	// take a graph as 2D array
	graph := make([][]int, len(prerequisites)+1)

	for i := 0; i < len(prerequisites); i++ {

		graph[prerequisites[i][0]] = append(graph[prerequisites[i][0]], prerequisites[i][1])

	}
	isCyclic := false
	// check we don't have cycle in any of path for any verdict
	for i := 0; i < len(graph); i++ {
		if graph[i] == nil {
			continue
		}
		isCyclic = isCycle2(graph[i][0], graph, visited, isSamePath)
		if isCyclic {
			break
		}
	}
	fmt.Println(!isCyclic)
}
func CourseComplete1(numCourses int, prerequisites [][]int) {

	visited := make([]bool, numCourses+1)
	isSamePath := make([]bool, numCourses+1)
	graph := make(map[int][]int)
	for i := 0; i < len(prerequisites); i++ {
		if graph[prerequisites[i][0]] == nil {
			graph[prerequisites[i][0]] = []int{prerequisites[i][1]}
		} else {
			graph[prerequisites[i][0]] = append(graph[prerequisites[i][0]], prerequisites[i][1])
		}
	}
	//display(numCourses, graph)
	isCyclic := false
	for i := 0; i <= numCourses; i++ {
		isCyclic = isCycle(i, graph, visited, isSamePath)
		if isCyclic {
			break
		}
	}
	fmt.Println(!isCyclic)
}

func isCycle(node int, graph map[int][]int, visited []bool, isSamePath []bool) bool {
	if isSamePath[node] {
		return true
	}
	if visited[node] {
		return false
	}
	visited[node] = true
	isSamePath[node] = true
	adj := graph[node]
	for i := 0; i < len(adj); i++ {
		cycle := isCycle(adj[i], graph, visited, isSamePath)
		if cycle {
			return true
		}
	}
	isSamePath[node] = false

	return false
}
func isCycle2(node int, graph [][]int, visited []bool, isSamePath []bool) bool {
	//  check if we got circle in same path then return as cycle true.
	if isSamePath[node] {
		return true
	}
	//  check if verdict is already visited then false as there is no cycle for this node.
	if visited[node] {
		return false
	}
	// set current verdict as visited and current path is ture if we got current again then its has cycle.
	visited[node] = true
	isSamePath[node] = true
	for i := 0; i < len(graph[node]); i++ {

		cycle := isCycle2(graph[node][i], graph, visited, isSamePath)
		if cycle {
			return true
		}
	}
	// set current path as false because we traversed all path but not find the cycle to it.
	//node can be present in other path as well.
	isSamePath[node] = false

	return false
}
func display(numCourses int, graph map[int][]int) {
	for i := 0; i <= numCourses; i++ {
		v, _ := graph[i]
		fmt.Print(i, " ")
		fmt.Print("->  ")
		for j := 0; j < len(v); j++ {
			fmt.Print(v[j])
		}
		fmt.Println()
	}
}
