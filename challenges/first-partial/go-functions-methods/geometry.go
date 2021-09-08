// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 156.

// Package geometry defines simple types for plane geometry.
//!+point
package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"math/rand"
)

type Point struct{x, y float64 }

// traditional function
func Distance(p, q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

// same thing, but as a method of the Point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

//!-point

//!+path

// A Path is a journey connecting the points with straight lines.
type Path []Point

// Distance returns the distance traveled along the path.
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func Intersect(p1,p2,q1,q2 Point)bool{
	o1:=Orientation(p1,q1,p2)
	o2:=Orientation(p1,q1,q2)
	o3:=Orientation(p2,q2,p1)
	o4:=Orientation(p2,q2,q1)

	if o1!=o2 && o3!=o4{
		return true
	}
	return false
	
}

func Orientation(p,q,r Point)int{
	answer:=(q.y-p.y)*(r.x-q.x)-(q.x-p.x)*(r.y-q.y)

	if answer==0{
		return 0
	}else if answer>0{
		return 1
	}else{
		return 2
	}
}


func main() {
	if len(os.Args)==1{
		fmt.Println("Please enter a number of sides when running the program")
		return
	}
	var perimeter float64
	var sides int
	var sidesError error
	var x float64
	var y float64
	sides,sidesError=strconv.Atoi(os.Args[1])
	if sidesError!=nil{
		fmt.Println("Invalid number of sides");
	}else{
		vertices:=make([]Point,sides)
		fmt.Println("Generating a", sides , "sides shape")
		fmt.Println("Figure's vertices")
		for i:=0; i<sides;i++{
			x=-100+(rand.Float64()*200)
			y=-100+(rand.Float64()*200)
			vertices[i]=Point{x,y}
		
		}

		if(sides>3){
			for j:=0; j<sides-3;j++{
				for k:=j+2;k<sides-1;k++{
					m:=0
					for m<1{
						if Intersect(vertices[j],vertices[j+1],vertices[k],vertices[(k+1)%sides]){
							x=-100+(rand.Float64()*200)
							y=-100+(rand.Float64()*200)
							vertices[j]=Point{x,y}
							x=-100+(rand.Float64()*200)
							y=-100+(rand.Float64()*200)
							vertices[j+1]=Point{x,y}
						}else{
							m=1
						}
					}
				}
			}
		}

		for l:=0;l<sides;l++ {
			fmt.Println("(",vertices[l].x,",",vertices[l].y,")")
		}

		fmt.Println("Figure's Perimeter")
		for n:=0; n<sides;n++{
			dist:=Distance(vertices[n],vertices[(n+1)%sides])
			if n<sides-1{
				fmt.Print(dist, "+")
			}else{
				fmt.Print(dist,"=")
			}
			perimeter+=dist
		}

		fmt.Print(perimeter, "\n")
	
	}


}

//!-path