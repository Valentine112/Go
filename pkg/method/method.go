package method

//import "fmt"
//Store
type Store struct {	
	far int
	out int
	mid int
	in int
}

//Against
type Against struct {			//Local
	posi []int
	oppo []int	
}

type Value struct {
	val int
	val_posi int
	outside int
	inside int
}

var a Against
var v Value

//Getmissing function to get the outside and inside index of the missing value
func Getmissing(problem [][][][]int) []Store {
	var s Store
	var missing []Store

	for a, val := range problem {
		for b, val1 := range val {
			for c, val2 := range val1 {
				for d, val3 := range val2 {
					if val3 == 0 {
						s.far = a
						s.out = b
						s.mid = c
						s.in = d
						missing = append(missing, s)
					}
				}
			}
		}
	}

	return missing
}

//Vertical
func Vertical(problem [][][][]int, m []Store) Against {
	for _, val := range m {
		for _, val1 := range problem {
			for _, val2 := range val1 {
				a.oppo = append(a.oppo, val2[val.mid][val.in])
			}
		}
		a.posi = append(a.posi, val.far, val.out, val.mid, val.in)
		
		break
	}

	return a
}

//Horizontal
func Horizontal(problem [][][][]int, a Against) Against {
	hor := a.posi[2]
	for _, val := range problem[a.posi[1]] {
		a.oppo = append(a.oppo, val[hor][0], val[hor][1], val[hor][2])
	}

	return a
}

//Box
func Box(problem [][][][]int, a Against) Against {
	for _, val := range problem[a.posi[0]][a.posi[1]] {
		a.oppo = append(a.oppo, val[0], val[1], val[2])
	}

	return a
}

//Result
func Result(a Against) Value {

	return v
}

//Delete
/*
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
}*/