package main

import (
	"fmt"
	"math/rand"
	"math"
	"sync"

)

type Stack struct{//Stack structure 
	lock sync.Mutex //used to make sure that if more than one go routine is trying to access it at one time it will lock all the go routines except the current one that is using it 
    stk []Triangle//array used for stack 

}
var wg sync.WaitGroup//wait group used to make sure that the calculations are done before program ends 
func genStack () *Stack{//generates a stack pointer 
	return &Stack {sync.Mutex{}, make([]Triangle,0), }
}
func (s *Stack) Push(t Triangle) {//adds the the "top" of the stack 
	s.lock.Lock()
	defer s.lock.Unlock()
	s.stk =append(s.stk,t)
	

}
func (s *Stack) Pop() {//removes the top element from the stack 
	s.lock.Lock()
    defer s.lock.Unlock()
   
    s.stk=s.stk[:len(s.stk)-1]
   
}

func (s *Stack) getSize() int {//returns the size of the stack 
	v:=len(s.stk)
	return v
}
func (s *Stack) Peek() Triangle{//peek at the top element 
	return s.stk[len(s.stk)-1]
}


type Point struct {//point structure used in the triangle structure
	x float64
	y float64
}
type Triangle struct {//triangle structure 
	A Point
	B Point
	C Point
}

func triangles10000() (result [10000]Triangle) {//returns a array of triangles 
 	rand.Seed(2120)
 	for i := 0; i < 10000; i++ {
 		result[i].A= Point{rand.Float64()*100.,rand.Float64()*100.}
 		result[i].B= Point{rand.Float64()*100.,rand.Float64()*100.}
 		result[i].C= Point{rand.Float64()*100.,rand.Float64()*100.}
 	}
 	return
}
func (t Triangle) Perimeter() float64 {//calculates the perimeter of a triangle 
	var ans float64
	ans+= math.Sqrt(math.Pow(t.B.x-t.A.x,2)+math.Pow(t.B.y-t.A.y,2))
	ans+= math.Sqrt(math.Pow(t.C.x-t.B.x,2)+math.Pow(t.C.y-t.B.y,2))
	ans+= math.Sqrt(math.Pow(t.C.x-t.A.x,2)+math.Pow(t.C.y-t.A.y,2))
	return ans
}

func (t Triangle) Area() float64 {//Calculates the area of a triangle 
	
	return 0.5*((t.B.x-t.A.x)*(t.C.y-t.A.y) - (t.C.x-t.A.x)*(t.B.y-t.A.y))
}

func classifyTriangles(highRatio *Stack, lowRatio *Stack,ratioThreshold float64, triangles []Triangle){//function that sorts the triangles into there respective stacks either low or high
	
	for _, tri := range triangles{
		per:= tri.Perimeter()
		area:= tri.Area()
		ratio:= per/area
		if ratio>ratioThreshold{
			highRatio.Push(tri)
		}else if ratio<ratioThreshold{
			lowRatio.Push(tri)
		}
	}
	
	wg.Done()//decrements wait group when routine is done 
}

func main() {
	
	res:=triangles10000()//generates the list of triangles 
	highRatio := genStack()//high stack
	lowRatio := genStack()//low stack 
	

	for i := 0; i < 10000; i+=1000 {//for loop that splits up the array into 10 slices 
		wg.Add(1)//increments wait group because new go routine is being called 
        go classifyTriangles(highRatio,lowRatio,1.0,res[i:i+1000])//new classifyTriangles Go routine
    }


	wg.Wait()//waits till all the routies are done 
	fmt.Printf("The highRatio list contains %v elements and has this triangle at the top %v.",highRatio.getSize(),highRatio.Peek())
	fmt.Println()
	fmt.Printf("The lowRatio list contains %v elements and has this triangle at the top %v.",lowRatio.getSize(),lowRatio.Peek())
	


	


}
