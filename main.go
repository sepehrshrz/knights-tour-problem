package main

import (
	"fmt"
	"sort"
)

type NextPoint struct {
	i int
	j int
}

const (
	N=8
)
var(
	nextMove int
	matris [N][N]int
)

func main(){
	i:=0
	j:=0
	fmt.Print("please enter i for first point: ")
	fmt.Scan(&i)
	fmt.Print("please enter j for first point: ")
	fmt.Scan(&j)
	nextMove=1
	matris[i][j]=nextMove
	SetHourse(i,j)
	 PrintMatris()
	
}

func SetHourse(i,j int)bool{
	if nextMove >=N*N{
		return true
	}

	array:=make([]NextPoint, 8)

	array[0].i=i+2
	array[0].j=j+1

	array[1].i=i+2
	array[1].j=j-1


	array[2].i=i-2
	array[2].j=j+1

	array[3].i=i-2
	array[3].j=j-1

	array[4].i=i+1
	array[4].j=j+2

	array[5].i=i-1
	array[5].j=j+2

	array[6].i=i+1
	array[6].j=j-2

	array[7].i=i-1
	array[7].j=j-2


	sort.Slice(array, func(i, j int) bool {
		return countUnvisitedNeighbors(array[i].i, array[i].j) < countUnvisitedNeighbors(array[j].i, array[j].j)
	})

	for _, v := range array {
	if FindTrueNextMove(v.i,v.j) {
		nextMove++
		matris[v.i][v.j]=nextMove
		
		if SetHourse(v.i,v.j) {
			return true
		}
		matris[v.i][v.j]=0
		nextMove--
	}
   }
   return false
}

func FindTrueNextMove(i,j int)bool{
	if i>= N || j>=N || i<0 || j< 0 {
		return false
	}

	if matris[i][j] > 0 {
		return false
	}

	return true
}

func countUnvisitedNeighbors(i, j int) int {
	count := 0
	array := make([]NextPoint, 8)

	array[0].i = i + 2
	array[0].j = j + 1

	array[1].i = i + 2
	array[1].j = j - 1

	array[2].i = i - 2
	array[2].j = j + 1

	array[3].i = i - 2
	array[3].j = j - 1

	array[4].i = i + 1
	array[4].j = j + 2

	array[5].i = i - 1
	array[5].j = j + 2

	array[6].i = i + 1
	array[6].j = j - 2

	array[7].i = i - 1
	array[7].j = j - 2

	for _, v := range array {
		if FindTrueNextMove(v.i, v.j) {
			count++
		}
	}

	return count
}

func PrintMatris(){
	for i := 0; i < len(matris); i++ {
		for j := 0; j < N; j++ {
			if matris[i][j]<10 {
				fmt.Print("0",matris[i][j],"  ")

			}else{

				fmt.Print(matris[i][j],"  ")

			}
		}
		fmt.Println()
	
	}
}