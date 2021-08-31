package main

// adds/subtracts/multiplies all values that are in the *values array.
// nValues is the number of values you're reading from the array
// operator will indicate if it's an addition (1), subtraction (2) or
// multiplication (3)

import(
	"fmt"
	"os"
	"strconv"
)
func calc(operator int, values []int) int {
	var result int=values[0]
	if operator==1{
		for i := 1; i <len(values); i++ {
			result+=values[i]
		}
		return result
	}else if operator==2{
		for i := 1; i <len(values); i++ {
			result-=values[i]
		}
		return result
	}else if operator==3{
		for i := 1; i <len(values); i++ {
			result*=values[i]
		}
		return result
	}else{
		fmt.Println("Invalid operator input :(")
		return -1
	} 
}

func main() {
	var operatorInput int
	if os.Args[1]=="add"{
		operatorInput=1
	}else if os.Args[1]=="sub"{
		operatorInput=2
	}else if os.Args[1]=="mult"{
		operatorInput=3
	}else{
		//operatorInput=0
		fmt.Println("Invalid operator input :(")
		return
	}

	valuesInput:= make([]int, len(os.Args)-2)
	var err3 error
	var value int
	for i := 0; i < len(valuesInput); i++ {
		value,err3=strconv.Atoi(os.Args[i+2])
		if err3!=nil{
			fmt.Println(err3)
			break
		}else{
			valuesInput[i]=value
		}
	}
	if err3==nil{
		fmt.Println(calc(operatorInput,valuesInput))
	}
	 

}
