package sudoku

import (
	"Go/pkg/method"
	//"fmt"
)
/*

** ASSIGNING

** 1 == The first list
** 2 == The second list
** 3 == The third list
** 4 == The fourth list

** ------- TO get the vertical against -------

** Find the index to which the missing value belongs
** On 3 list, target the 2 list,
** Then proceed to get the value with it's index matching the missing value index of each of the 4 list

** ------- To get the horizontal against -------

** Firstly, get the 2 index of the missing value
** Then proceed to get the values of each of the 3, 4 first section e.g 3[0]

** ------- To get the box against -------

** Firstly, get the missing value 3 index
** Then get each of the 4 values of each section

*/


var  problem1 = [][][][]int{
	{
		{{6, 3, 0}, {0, 2, 0}, {0, 0, 0}}, //box  		//				//
		{{0, 0, 0}, {0, 0, 3}, {0, 1, 7}},				//	horizontal
		{{0, 8, 1}, {0, 0, 0}, {4, 3, 0}}},				//
	{
		{{0, 9, 6}, {0, 0, 0}, {0, 8, 0}}, 								// vertical
		{{4, 0, 0}, {7, 6, 2}, {0, 0, 0}},
		{{5, 7, 0}, {0, 0, 0}, {6, 0, 0}}},
	{
		{{0, 6, 0}, {3, 0, 9}, {0, 0, 0}}, 								//
		{{0, 2, 0}, {0, 0, 0}, {0, 0, 0}},
		{{0, 0, 0}, {0, 6, 0}, {0, 0, 9}}},
}

var values = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}		//Global


func Solution1() method.Value {
	missing := method.Getmissing(problem1)
	ver := method.Vertical(problem1, missing)
	hor := method.Horizontal(problem1, ver)
	box := method.Box(problem1, hor)
	res := method.Result(box, values)

	return res
}