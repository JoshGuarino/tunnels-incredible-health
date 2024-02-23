package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
)

const START = "https://tunnels.incredible.health"
const USAGE = "Usage 'go run tunnels.go <arg>':\n	dfs - depth first search\n	bfs - breadth first search"

var count = 0
var exitRoute = []Path{}

type Node struct {
	Description string `json:"description"`
	Left        string `json:"left"`
	Right       string `json:"right"`
	Back        string `json:"back"`
	AtExit      bool   `json:"atExit"`
}

type Path struct {
	Direction string
	NodeUrl   string
}

func resetTerminal() {
	clear := exec.Command("clear")
	clear.Stdout = os.Stdout
	clear.Run()
}

func printMainOutput(nodeUrl string) {
	fmt.Println("TOTAL:", count, "\nCHECKING -->", nodeUrl)
}

func printExitRoute() {
	fmt.Println("\nEXIT ROUTE:")
	for _, path := range exitRoute {
		fmt.Println(path.Direction, "-->", path.NodeUrl)
	}
}

func exit(node Node) {
	fmt.Println("\n", node.Description)
	os.Exit(0)
}

func getNode(url string) Node {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		println(err)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	node := Node{}
	json.Unmarshal(body, &node)
	return node
}

func findExitDfs(nodeUrl string, direction string) {
	node := getNode(nodeUrl)
	exitRoute = append(exitRoute, Path{Direction: direction, NodeUrl: nodeUrl})
	count++
	resetTerminal()
	printMainOutput(nodeUrl)
	printExitRoute()

	if node.AtExit {
		exit(node)
	}

	if node.Left == "" && node.Right == "" {
		exitRoute = exitRoute[0 : len(exitRoute)-1]
		return
	}

	paths := [2]string{node.Left, node.Right}
	for idx, path := range paths {
		if idx == 0 {
			direction = "left"
		} else {
			direction = "right"
		}
		findExitDfs(path, direction)
	}
	exitRoute = exitRoute[0 : len(exitRoute)-1]
}

func findExitBfs(startUrl string) {
	queue := []string{startUrl}
	for len(queue) > 0 {
		node := getNode(queue[0])
		count++
		resetTerminal()
		printMainOutput(queue[0])
		queue = queue[1:]

		if node.AtExit {
			exit(node)
		}

		if node.Left != "" {
			queue = append(queue, node.Left)
		}

		if node.Right != "" {
			queue = append(queue, node.Right)
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println(USAGE)
		os.Exit(0)
	}
	search := os.Args[1]
	switch search {
	case "bfs":
		findExitBfs(START)
	case "dfs":
		findExitDfs(START, "start")
	default:
		fmt.Println(USAGE)
	}
}
