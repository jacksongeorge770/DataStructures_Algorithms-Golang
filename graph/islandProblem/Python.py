class Solution:
    def numIslands(self, grid: List[List[str]]) -> int:
        self.grid = grid
        self.rows = len(grid)
        self.cols = len(grid[0])
        self.visited = [[False] * self.cols for _ in range(self.rows)]

        count = 0
        for r in range(self.rows):
            for c in range(self.cols):
                if self.grid[r][c] == '1' and not self.visited[r][c]:
                    self.dfs(r, c)
                    count += 1

        return count

    def dfs(self, r: int, c: int):
        if (
            r < 0 or r >= self.rows or
            c < 0 or c >= self.cols or
            self.grid[r][c] == '0' or
            self.visited[r][c]
        ):
            return

        self.visited[r][c] = True
        self.dfs(r + 1, c)
        self.dfs(r - 1, c)
        self.dfs(r, c + 1)
        self.dfs(r, c - 1)


from typing import List

# ✅ DFS function defined outside the class
def dfs(grid: List[List[str]], visited: List[List[bool]], r: int, c: int):
    rows = len(grid)
    cols = len(grid[0])

    # Boundary conditions and base case
    if (
        r < 0 or r >= rows or
        c < 0 or c >= cols or
        grid[r][c] == '0' or
        visited[r][c]
    ):
        return

    visited[r][c] = True

    # Explore all 4 directions
    dfs(grid, visited, r + 1, c)
    dfs(grid, visited, r - 1, c)
    dfs(grid, visited, r, c + 1)
    dfs(grid, visited, r, c - 1)

# ✅ Solution class
class Solution:
    def numIslands(self, grid: List[List[str]]) -> int:
        if not grid:
            return 0

        rows = len(grid)
        cols = len(grid[0])
        visited = [[False] * cols for _ in range(rows)]
        count = 0

        for r in range(rows):
            for c in range(cols):
                if grid[r][c] == '1' and not visited[r][c]:
                    dfs(grid, visited, r, c)
                    count += 1

        return count

