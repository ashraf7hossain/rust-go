package main

import(
    "fmt"
    "bufio"
    "os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer  = bufio.NewWriter(os.Stdout)
func print(f string , a ...interface{}){
    fmt.Fprintf(writer, f, a...)
}
func read(f string , a ...interface{}){
    fmt.Fscanf(reader , f , a...)
}
func Min(x int , y int) int {
  if x > y{
    return y
  }
  return x
}

func main() {
    defer writer.Flush()
	var n , m int 
	read("%d %d\n", &n , &m)
	g := make([][]int , n + 1)
	for i := 0 ; i < m; i++{
	  var x, y int 
	  read("%d %d\n",&x , &y)
	  g[x] = append(g[x] , y)
	  g[y] = append(g[y] , x)
	}
  d1 := BFS(1 , g)
  d2 := BFS(n , g)
  
  inf := 123456789
  for i := 1; i <= n; i++{
    ans := d1[n]
    ans = Min(ans , d1[0] + d2[i])
    ans = Min(ans , d1[i] + d2[0])
    if ans >= inf{
      print("%d ", -1)
    }else{
      print("%d ", ans)
    }
  }
  
	
}

func BFS(s int , g [][]int) []int{
  var q []int
  inf := 123456789
  q = append(q , s)
  dst := make([]int, len(g))
  for x := 0 ; x < len(dst); x++{
    dst[x] = inf
  }
  dst[s] = 0
  for i := 0; i < len(q); i++{
    u := q[i]
    for _ ,v := range(g[u]){
      if dst[v] != inf {
        continue
      }
      q = append(q , v)
      dst[v] = dst[u] + 1
    } 
  }
  return dst
}