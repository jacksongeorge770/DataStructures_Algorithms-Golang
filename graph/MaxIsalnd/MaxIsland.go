func maxAreaOfIsland(grid [][]int) int {
maxArea:=0

rows:=len(grid)
cols:=len(grid[0])

for r:=0;r<rows;r++{
    for c:=0;c<cols;c++{
      if grid[r][c] == 1{
     area:=dfs(grid,r,c)

     if area>maxArea{
        maxArea=area
     }
      }
    }
}     

return maxArea
}

func dfs(grid [][] int ,r int,c int)int{

    rows:=len(grid)
    cols:=len(grid[0])
    
     if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] == 0 {
        return 0
    }
     area:=1 
    grid[r][c]=0 //visited
    
    area+=dfs(grid,r+1,c)
    area+=dfs(grid,r-1,c)
    area+=dfs(grid,r,c+1)
    area+=dfs(grid,r,c-1)

    return area
}
