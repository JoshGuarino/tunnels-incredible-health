import requests

START = 'https://tunnels.incredible.health'
visted = set()

def get_node(url: str) -> dict:
    r = requests.get(url, headers={"Accept": "application/json"})
    return r.json() 

def find_exit(node_url: str) -> None:
    node = get_node(node_url)

    if node['atExit']:
        print(node)
        exit()

    if node['left'] == None and node['right'] == None:
        print('dead end')
        return

    paths = [node['left'], node['right']]
    for path in paths:
        print('left') if path == node['left'] else print('right')
        find_exit(path)

if __name__ == "__main__":
    find_exit(START)
    