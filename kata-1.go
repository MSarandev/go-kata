package main

import (
	"fmt"
	"strconv"
	"strings"
)

func FizzBuzzCuckooClock(time string) string {
	// splice the minutes/hour
	hour, _  := strconv.Atoi(time[:strings.IndexByte(time, ':')])
	minute, _ := strconv.Atoi(time[strings.IndexByte(time, ':')+1:])
	var res string

	// check divisibility
	if minute % 3 == 0 && minute % 5  == 0{
		// check if on the hour
		if minute == 0{
			// check the hour
			if hour > 12{
				hour = hour - 12
			}else if hour == 00{
				hour = 12
			}

			// print for the hours (upto 12)
			for i := 0; i < hour; i++{
				res += "Cuckoo "
			}
		}else if minute == 30{
			res = "Cuckoo"
		}else{
			res = "Fizz Buzz"
		}
	}else if minute % 3 == 0{
		res = "Fizz"
	}else if minute % 5 == 0{
		res = "Buzz"
	}else{
		res = "tick"
	}

	res = strings.TrimRight(res, " ")

	return res
}

func main(){
	fmt.Println(FizzBuzzCuckooClock("00:00"))
}