import os
import sys
import requests

START = 'https://tunnels.incredible.health'
USAGE = "Usage 'python3 tunnels.py <arg>':\n	dfs - depth first search\n	bfs - breadth first search"
count = 0
exit_route = []

def get_node(url: str) -> dict:
    r = requests.get(url, headers={"Accept": "application/json"})
    return r.json() 

def find_exit_dfs(node_url: str, direction: str) -> None:
    node = get_node(node_url)
    exit_route.append({"direction": direction, "node_url": node_url})
    global count
    count += 1
    os.system('clear')
    print(f'TOTAL: {count}\nCHECKING --> {node_url}')
    print('\nEXIT ROUTE:')
    for path in exit_route:
        print(f'{path["direction"]} --> {path["node_url"]}')

    if node['atExit']:
        print(node['description'])
        exit()

    if node['left'] == None and node['right'] == None:
        exit_route.pop()
        return

    paths = [node['left'], node['right']]
    for path in paths:
        direction = 'left' if path == node['left'] else 'right'
        find_exit_dfs(path, direction)
    exit_route.pop()

def find_exit_bfs(start_url: str) -> None:
    queue = [start_url]
    while queue:
        global count
        count += 1
        node = get_node(queue[0])
        os.system('clear')
        print(f'TOTAL: {count}\nCHECKING --> {queue[0]}')
        queue.pop(0)

        if node['atExit']:
            print(f'\n{node["description"]}')
            exit()

        if node['left']:
            queue.append(node['left'])

        if node['right']:
            queue.append(node['right'])
    

if __name__ == "__main__":
    if len(sys.argv) < 2:
        print(USAGE)
        exit()
    if sys.argv[1] == 'bfs':
        find_exit_bfs(START)
    if sys.argv[1] == 'dfs':
        find_exit_dfs(START, 'start')
    print(USAGE)
    