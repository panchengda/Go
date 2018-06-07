package main

import "fmt"
import "math"

type geo interface{
	area() float64
	permit() float64
}

type square struct{
	a float64
	b float64
}

type circle struct{
	r float64
}

func(s square) area() float64{
	return s.a*s.b
}

func(s square) permit() float64{
	return 2*(s.a+s.b)
}

func(c circle) area() float64{
	return math.Pi*c.r*c.r
}

func(c circle) permit() float64{
	return 2*math.Pi*c.r
}

func f(g geo ){
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.permit())
}

func main(){
	s := square{3,7}
	c := circle{11}

	f(s)
	f(c)
}


