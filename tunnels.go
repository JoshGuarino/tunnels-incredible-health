package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const START = "https://tunnels.incredible.health"

var exitRoute = []string{}

type Node struct {
	Description string `json:"description"`
	Left        string `json:"left"`
	Right       string `json:"right"`
	Back        string `json:"back"`
	AtExit      bool   `json:"atExit"`
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

func findExitDfs(nodeUrl string) {
	node := getNode(nodeUrl)
	exitRoute = append(exitRoute, nodeUrl)

	if node.AtExit {
		fmt.Println(node.Description)
		route, _ := json.MarshalIndent(exitRoute, "", "")
		_ = ioutil.WriteFile("exit_route.json", route, 0644)
		os.Exit(0)
	}

	if node.Left == "" && node.Right == "" {
		fmt.Println("dead end")
		exitRoute = exitRoute[0 : len(exitRoute)-1]
		return
	}

	paths := [2]string{node.Left, node.Right}
	for idx, path := range paths {
		if idx == 0 {
			fmt.Println("left")
		} else {
			fmt.Println("right")
		}
		findExitDfs(path)
	}
	exitRoute = exitRoute[0 : len(exitRoute)-1]
}

func findExitBfs(startUrl string) {
	queue := []string{startUrl}
	for len(queue) > 0 {
		node := getNode(queue[0])
		fmt.Println(queue[0])
		queue = queue[1:]

		if node.AtExit {
			fmt.Println(node.Description)
			os.Exit(0)
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
	findExitDfs(START)
	findExitBfs(START)
}
