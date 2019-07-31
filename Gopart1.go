package main

import (
"fmt"
"errors"
"strings"
"bufio"
"os"
)

type Trip struct{//trip structure 
	destin string 
	weight float32
	deadline int 
}
type Truck struct{//truck structure 
	vehicle string
	name string 
	destin string 
	speed float32
	capacity float32
	load float32
}
type Pickup struct{//pickup structure 
	Truck
	isPrivate bool
}
type TrainCar struct{//traincar structure 
	Truck
	railway string
}

func NewTruck () Truck{//method that generates a New Truck structure with default values 
	var truck Truck 
	truck.vehicle="Truck"
	truck.name="Truck"
	truck.speed=40
	truck.capacity=10
	truck.load=0
	return truck
}
func NewPickUp () Pickup{//method that generates a New Pickup structure with default values
	var pick Pickup 
	pick.vehicle="Pickup"
	pick.name="Pickup"
	pick.speed=60
	pick.capacity=2
	pick.load=0
	pick.isPrivate=true
	return pick
}
func NewTrainCar () TrainCar{//method that generates a New TrainCar structure with default values
	var train TrainCar 
	train.vehicle="TrainCar"
	train.name="TrainCar"
	train.speed=30
	train.capacity=30
	train.load=0
	train.railway="CNR"
	return train
}
type Transporter interface{//Transport interface used so that for each type of transporter these methods can be reused 
	addLoad(Trip) error
	print()
}
func NewTorontoTrip (weight float32, deadline int) *Trip{//generates a new structure called Toronto trip where the location is Toronto
	var point *Trip
	var torTrip Trip
	torTrip.destin="Toronto"
	torTrip.weight=weight
	torTrip.deadline =deadline
	point= &torTrip
	return point

}
func NewMontrealTrip (weight float32, deadline int) *Trip{//same as above method just for Montreal
	var point *Trip
	var monTrip Trip
	monTrip.destin="Montreal"
	monTrip.weight=weight
	monTrip.deadline =deadline
	point= &monTrip
	return point

}
func (t *Truck) addLoad(trip Trip) error{// This is the addLoad mehthod for the truck object 

	if t.destin=="" &&trip.weight< t.capacity-t.load && int(200/t.speed)<trip.deadline {//ensures that the current truck can be used 
		t.destin=trip.destin
	}

	// all if statements below are for checking for possible errors 
	if strings.ToLower(trip.destin)!=strings.ToLower(t.destin){
		err:= errors.New("Error: Other destination")
		return err
	}
	
	if t.destin=="Toronto" && int(400/t.speed)>trip.deadline{
		err:= errors.New("Error: Other destination")
		return err
	}
	if t.destin=="Montreal" && int(200/t.speed)>trip.deadline{
		err:= errors.New("Error: Other destination")
		return err
	}
	if trip.weight> t.capacity-t.load{
		err:= errors.New("Error: Out of capacity")
		return err
	}
	t.load+=trip.weight
	
	return nil
	
}
func (t *Truck) print(){// The print method for the truck structure 
	fmt.Printf("%v %v to %v with %f tons \n",t.vehicle,t.name,t.destin,t.load)
}
func (p *Pickup) addLoad(trip Trip) error{//same as the addLoad for the Truck structure but its now for pickup

	if p.destin=="" && trip.weight< p.capacity-p.load && int(200/p.speed)<trip.deadline{
		p.destin=trip.destin
		
	}

	
	if strings.ToLower(trip.destin)!=strings.ToLower(p.destin){
		err:= errors.New("Error: Other destination")
		return err
	}

	if p.destin=="Toronto" && int(400/p.speed)>trip.deadline{
		err:= errors.New("Error: Other destination")
		return err
	}
	if p.destin=="Montreal" && int(200/p.speed)>trip.deadline{
		err:= errors.New("Error: Other destination")
		return err
	}
	if trip.weight> p.capacity-p.load{
		err:= errors.New("Error: Out of capacity")
		return err
	}
	p.load+=trip.weight
	return nil
	
}
func (p *Pickup) print(){//same as print for Truck structure but now for Pickup
	fmt.Printf("%v %v to %v with %f tons (Private: %v)\n",p.vehicle,p.name,p.destin,p.load,p.isPrivate)
}

func (t *TrainCar) addLoad(trip Trip) error{//same as the addLoad for the Truck structure but its now for Traincar 

	if t.destin=="" && trip.weight< t.capacity-t.load && int(200/t.speed)<trip.deadline{
		t.destin=trip.destin
	}

	if strings.ToLower(trip.destin)!=strings.ToLower(t.destin){
		err:= errors.New("Error: Other destination")
		return err
	}
	
	if t.destin=="Toronto" && int(400/t.speed)>trip.deadline{
		err:= errors.New("Error: Other destination")
		return err
	}
	if t.destin=="Montreal" && int(200/t.speed)>trip.deadline{
		err:= errors.New("Error: Other destination")
		return err
	}
	if trip.weight> t.capacity-t.load{
		err:= errors.New("Error: Out of capacity")
		return err
	}
	t.load+=trip.weight
	
	return nil

	
}
func (t *TrainCar) print(){//same as print for Truck structure but now for TrainCar
	fmt.Printf("%v %v to %v with %f tons (%v)\n",t.vehicle,t.name,t.destin,t.load,t.railway)
}
func main() {//Creating all the transporters used in this program
	ta:=NewTruck()
	ta.name="A"
	tb:=NewTruck()
	tb.name="B"
	pa:=NewPickUp()
	pa.name="A"
	pb:=NewPickUp()
	pb.name="B"
	pc:=NewPickUp()
	pc.name="C"
	tca:=NewTrainCar()
	tca.name="A"
	table := [6]Transporter{&ta,&tb,&pa,&pb,&pc,&tca}
	var trips []Trip
	var err error
	var trip *Trip
	
	for true{//loop assigns each vehicle a trip 
		fmt.Println("Destination: (t)oronto, (m)ontreal, else exit")
		reader := bufio.NewReader(os.Stdin)
    	ans, _ := reader.ReadString('\n')
 		ans=string(ans[0])
 		ans=strings.ToLower(ans)//changes the ans into a string that is one character long 

 		if ans!="m" && ans!="t"{//if its not a Toronto or Montreal trip it breaks out of the main loop
    		fmt.Println("Not going to To or Montreal, bye!")
    		break
 		}

 		fmt.Print("Weight: ")
		var weightG float32
		fmt.Scanf("%f", &weightG)
    	fmt.Print("Deadline (in hours): ")
    	var deadline int
    	fmt.Scan(&deadline)

    	if ans=="t"{//checks if its a Toronto trip 
    		trip=NewTorontoTrip(weightG,deadline)
    	}else if ans=="m"{//Checks if its a Montreal Trip
    		trip=NewMontrealTrip(weightG,deadline)
    	}
    	 for  i:=0 ; i<6;i++{//goes through the list of transporters and finds the first possible vehicle that satistfies the trip characteristics
    		 	err=table[i].addLoad(*trip)
    		 	if err==nil{
    		 		break
    		 	}
    		 	fmt.Println(err)
    		 	}
    		 	if err==nil{
    		 		trips=append(trips,*trip)
    		 	}
    	

    	

	}
	fmt.Printf("Trips: %v",trips)
	fmt.Println()
	for i:=0;i<6;i++{
		table[i].print()
	}

}
