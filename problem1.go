package main

import (
	"fmt"
)

func main() {
	z := 1000 //finds sum of everything divisible by three and five below this number
	x := 0
	y := 0
	//multiple zeros because variables carry over (another reason for functions)
	p := 0
	a := 0
	//this could be a function
	threemulti := make([]int, 0) //creats array
	for {                        //this loop adds 1 to the x which is then multiplied by three until it is = to or > than 'z'
		x++
		y = 3 * x
		if y >= z { //for loop at the end so it does the multiplication before instead of after
			break
		}
		threemulti = append(threemulti, y) //lets array grow in accordance to number of int
		fmt.Println(y)
	}
	sumthree := 0                  //this loop lets you add all the int in the slice (or array)
	for _, v := range threemulti { //btw the '_' means it doesn't care.    ex. "var _ = 5" means nothing
		sumthree += v
	}
	fmt.Println(sumthree)
	// this could also be a function
	fivemulti := make([]int, 0) //creates array
	//fmt.Println("before five multi for loop")
	for {
		p++
		a = 5 * p
		if a >= z {
			break
		}
		if a%3 == 0 {
			continue
		}
		fivemulti = append(fivemulti, a) //already covered
		fmt.Println(a)
	}
	sumfive := 0
	for _, v := range fivemulti {
		sumfive += v
	}
	fmt.Println(sumfive)

	sumAll := sumfive + sumthree //adds all number divisible by three and five below 'z'

	fmt.Println(sumAll)
}

//hey emyrk :D
