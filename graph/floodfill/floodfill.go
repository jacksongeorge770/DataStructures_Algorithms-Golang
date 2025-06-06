func floodFill(image [][]int, sr int, sc int, color int) [][]int {
     
     orginal:=image[sr][sc]
    if orginal == color {
        return image
    }
     row:=len(image)
     col:=len(image[0])

     dfs(image,sr,sc,orginal,color,col,row)
 
 return image
}

func dfs(image [][]int,sr,sc, orginal,color,col,row int){
  
  if (sr>=row||sr<0||sc<0||sc>=col){
    return
  }

  if image[sr][sc]!=orginal{
    return
  }
  image[sr][sc]=color


 dfs(image,sr+1,sc,orginal,color,col,row) 
dfs(image,sr-1,sc,orginal,color,col,row)
dfs(image,sr,sc+1,orginal,color,col,row)
dfs(image,sr,sc-1,orginal,color,col,row)
}
