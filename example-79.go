package main
import (
	"fmt"
)
func main() {
	board :=[][]byte{{'A','B','C','E'},{'S','F','C','S'},{'A','D','E','E'}}
	is := exist(board,"SEE")
	fmt.Printf("is:%+v\n",is)
}


func exist(board [][]byte, word string) bool {
	mapPostion := findPosition(string(word[0]),board)
	fmt.Printf("mapPostion:%+v\n",mapPostion)
	return true
}


func findPosition(w string,board [][]byte) []map[string]int{
	position := make([]map[string]int,0)
	for i:=0;i <len(board);i++ {
		for j := 0;j<len(board[i]);j++ {
			if string(board[i][j]) == w {
				m := make(map[string]int,0)
				m["x"] = i
				m["y"] = j
				position = append(position,m)
			}
			fmt.Printf("i:%d,j:%d,w:%s\n",i,j,string(board[i][j]))
		}
	}
	return position
}