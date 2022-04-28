package sudoku

import (
	//"sort"
	"fmt"
)

var problem = [][]int{{0, 2, 4}, {6, 0, 8}, {0, 0, 0}}		//Global

var list = []int{1, 2, 2, 4, 5, 6, 7, 8, 9}		//Global

type store struct {			//Global
	out int
	in int
}

type against struct {			//Local
	posi []int
	oppo []int	
}

type result struct {
	val int
	val_posi int
	outside int
	inside int
}


var r result

//Getmissing function to get the outside and inside index of the missing value
func Getmissing(problem [][]int) []store {
	var s store
	var missing []store
	for i, val := range problem {
		for x, val1 := range val {
			if(val1 == 0) {
				s.out = i
				s.in = x

				missing = append(missing, s)
			}
		}
	}

	return missing
}

//To get the values that shouldn't be matched in relation to the missing value
//Problemvalues
func Problemvalues(m []store) against {
	var a against
	for _, val := range m {
		for x, val1 := range problem {
			if val.out == x {
				for _, val2 := range val1 {
					a.oppo = append(a.oppo, val2)
				}
			}
			a.oppo = append(a.oppo, val1[val.in])
		}

		a.posi = append(a.posi, val.out)
		a.posi = append(a.posi, val.in)

		break
	}

	return a
}

/*
** Delete function to be executed
** when the outside value of the previous problem matches the current one
** So there is no replica and a block doesn't have multiple of a single value
*/
//Delete
func Delete(ind int) []int {
	var b []int
	if ind == 0 {
		b = append(b, list[ind + 1:]...)
	}else if(ind == len(list) - 1) {
		b = append(b, list[:ind]...)
	}else{
		b = append(b, list[:ind]...)
		b = append(b, list[ind + 1:]...)
	}

	return b
}

//Getresult function to get the appropriate values for that missing point
func Getresult(a against) result {
	if a.posi[0] == r.outside {
		b := Delete(r.val_posi)
		list = b
	}
	var check string = "Unchecked"
	for ind, val := range list{
		for x, val1 := range a.oppo {
			if val != val1 {
				if x + 1 == len(a.oppo) {
					r.val = val
					r.outside = a.posi[0]
					r.inside = a.posi[1]

					r.val_posi = ind

					fmt.Println("---")
					fmt.Println(ind)

					fmt.Println(list[ind])

					check = "Checked"

					break
				}else{
					continue
				}
			}else{
				break
			}
		}
		if check == "Checked" {
			break
		}
	}

	return r
}

/*
** Solution, a function that acts as concurrent when the problem hasn't ended
** It calls all the necessary functions
*/
var i int

//Solution
func Solution() [][]int {

	i++
	var r result
	missing := Getmissing(problem)
	a := Problemvalues(missing)
	r = Getresult(a)

	//Verify that the value picked is not 0
	if r.val != 0 {
		//Replacing the correct value in the problem
		problem[r.outside][r.inside] = r.val
	}

	if len(missing) > 1 {
		Solution()
	}
	return problem
}


/*
** Get the position of value in a list

	lists := sort.IntSlice(list[0:])
	sort.Sort(lists)
	pos := sort.SearchInts(list, val)
*/