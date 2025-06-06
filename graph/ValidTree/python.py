from collections import defaultdict

def validTree(n, edges):
    if len(edges) != n - 1:
        return False

    # Build the adjacency list
    adj = defaultdict(list)
    for a, b in edges:
        adj[a].append(b)
        adj[b].append(a)

    visited = {}

    def dfs(node, parent):
        if visited.get(node, False):
            return False
        visited[node] = True

        for neighbor in adj[node]:
            if neighbor == parent:
                continue
            if not dfs(neighbor, node):
                return False
        return True

    if not dfs(0, -1):
        return False

    return len(visited) == n
