package main

import "fmt"

type rect struct{
	a int
	b int
}

func (r rect )area() int{
	return r.a * r.b
}

//func(r rect) permin() int{}  should be the same
func (r* rect) perim() int{
	return 2*(r.a+r.b)
}

func main(){
	r1 := rect{2,3}
	pr1 := &r1
	
	r2 := rect{5,7}
	pr2 := &r2
	
	fmt.Println(r1.area())
	fmt.Println(r1.perim())

	fmt.Println(pr1.area())
	fmt.Println(pr1.perim())
	

	fmt.Println(r2.area())
	fmt.Println(r2.perim())

	fmt.Println(pr2.area())
	fmt.Println(pr2.perim())
}

