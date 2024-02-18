package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const START = "https://tunnels.incredible.health"

type Node struct {
	Description string `json:"description"`
	Left        string `json:"left"`
	Right       string `json:"right"`
	Back        string `json:"back"`
	AtExit      bool   `json:"atExit"`
}

func getNode(url string) Node {
	node := Node{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		println(err)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &node)

	// fmt.Println(node)

	return node
}

func findExit(nodeUrl string) {
	node := Node{}
	node = getNode(nodeUrl)

	if node.AtExit {
		fmt.Println(node)
		os.Exit(0)
	}

	if node.Left == "" && node.Right == "" {
		fmt.Println("dead end")
		return
	}

	paths := [2]string{node.Left, node.Right}
	for idx, path := range paths {
		if idx == 0 {
			fmt.Println("left")
		} else {
			fmt.Println("right")
		}
		findExit(path)
	}
}

func main() {
	findExit(START)
}
