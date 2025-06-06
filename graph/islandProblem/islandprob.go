func numIslands(grid [][]byte) int {
    if len(grid)==0{
        return 0
    }

    rows :=len(grid)
    cols :=len(grid[0])

    visited:=make([][]bool,rows)

    for i:=range visited{
        visited[i]=make([]bool, cols)
    }

    count:=0

    for r:=0;r<rows;r++{
        for c:=0;c<cols;c++{
              
              if grid[r][c]=='1' && !visited[r][c]{
                    
                    dfs(grid,visited,r,c)
                    count++
              }
        }
    }

    return count

}

func dfs(grid [][]byte, visited [][]bool,r,c int){
    rows:=len(grid)
    col:=len(grid[0])

    if r<0|| r>=rows ||c>=col||c<0|| visited[r][c]|| grid[r][c]=='0'{
             
             return
    }
       
       visited[r][c]=true
    dfs(grid,visited,r+1,c)
    dfs(grid,visited,r-1,c)
    dfs(grid,visited,r,c+1)
    dfs(grid,visited,r,c-1)
}
















