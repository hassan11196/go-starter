package main

import (
	"fmt"
	"time"
)

func main(){


	switch time.Now().Weekday(){

		case time.Saturday, time.Sunday:
			fmt.Println("It is the weekend")

		case time.Monday:
			fmt.Println("It is monday :/ ")
		
		default:
			fmt.Println("It is a weekday")
	}

	whatAmI := func(i interface{}) {
        switch t := i.(type) {
        case bool:
            fmt.Println("I'm a bool")
        case int:
            fmt.Println("I'm an int")
        default:
            fmt.Printf("Don't know type %T\n", t)
        }
    }
    whatAmI(true)
    whatAmI(1)
    whatAmI("hey")

}