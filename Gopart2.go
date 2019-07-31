package main

import (
"fmt"
"sync"
"math/rand"
"time"
)
const (
	NumRoutines = 3
	NumRequests = 1000
)

// global semaphore monitoring the number of routines
var semRout = make(chan int, NumRoutines)// this is a channel that has a max ie buffer of 3 channels 
// global semaphore monitoring console
var semDisp = make(chan int, 1)
// Waitgroups to ensure that main does not exit until all done
//A semaphore limits the number of go routines that can run at once
var wgRout sync.WaitGroup
var wgDisp sync.WaitGroup


type Task struct {//the task structure 
	a, b float32
	disp chan float32
}

func solve(t *Task){//takes in a task and solve it by adding the a and b variables for a given task
	wgDisp.Add(1)//adds counter to display wait group
	time.Sleep(time.Duration(rand.Intn(14)+1) * time.Second)//timer from 0 to 15 seconds 
	ans:=t.a +t.b
	t.disp <- ans
	
	wgRout.Done()//just finished the routine ie calculation therefore decrement routine wait group counter 

	
}
func handleReq(t *Task) {//calls solve and when solve is done the routine semaphore is decremented
	solve(t)
	<-semRout
}
func ComputeServer()(chan *Task){//go concurent "server" that uses channel factory pattern and creates the channel then runs a go function ie concurent  
	
	c := make(chan *Task)//creates the channel 
	go func(){// creates the compute server allowing for the channel to be returned while still being able to use the func as a intermediate to handleReq
		
		for{
			v:=<-c
			semRout<-1//inc the semaphore counter for number of handle request go routines 
			go handleReq(v)//create go routine with handle req
	}
	}()
	
	return c

}
func DisplayServer() (chan float32) {//the display server used to print the result of of solve  
	c := make(chan float32)
	go func(){//used so that we can return the channel while still using the display server 
		for{
			v:= <-c//waits for info from channel
			semDisp<-1//inc display semaphore
			fmt.Println("-------")
			fmt.Printf("Result: %v",v)
			fmt.Println()
			<-semDisp//dec display counter semaphore 
			wgDisp.Done()//display is done therefore waitgroup decremented 
	}
	
	}()

	return c
}

func main() {
	dispChan := DisplayServer()//makes the display server and saves the returned channel to dispChan
	reqChan := ComputeServer()//makes the Compute server and saves the returned channel to reqChan

	
	for {
		var a, b float32
		semDisp<-1//inc the display semaphore
		fmt.Print("Enter two numbers: ")//gets user input
		fmt.Scanf("%f %f \n", &a, &b)
		fmt.Printf("%f %f \n", a, b)
		<-semDisp//decrements display semaphore because output is done 
		if a == 0 && b == 0 {//the break for the infinite loop when a and b are 0
			break
		}else{
			wgRout.Add(1)//increments the routine wait group because a new request is being generated 
			var t Task
			t.a=a
			t.b=b
			t.disp = dispChan
			p:=&t//creates the task pointer 
			reqChan<-p//sends that pointer over to reqChan 

		}

	time.Sleep( 1e9 )
	}

	wgDisp.Wait()//both wait groups are used here to insure that the program does not end until all request are furfilled 
	wgRout.Wait()
	
}


		


