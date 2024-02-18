import requests

START = 'https://tunnels.incredible.health'

def get_node(url: str) -> dict:
    r = requests.get(url, headers={"Accept": "application/json"})
    return r.json() 

def find_exit_dfs(node_url: str) -> None:
    node = get_node(node_url)

    if node['atExit']:
        print(node['description'])
        exit()

    if node['left'] == None and node['right'] == None:
        print('dead end')
        return

    paths = [node['left'], node['right']]
    for path in paths:
        print('left') if path == node['left'] else print('right')
        find_exit_dfs(path)

def find_exit_bfs(start_url: str) -> None:
    queue = [start_url]
    while queue:
        node = get_node(queue[0])
        print(queue[0])
        queue.pop(0)

        if node['atExit']:
            print(node['description'])
            exit()

        if node['left']:
            queue.append(node['left'])

        if node['right']:
            queue.append(node['right'])
    

if __name__ == "__main__":
    find_exit_bfs(START)
    find_exit_dfs(START)
    