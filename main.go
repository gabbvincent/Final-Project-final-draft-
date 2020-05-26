//Programmers Name: Vincent G.
//Description: Final project

package main

import (

 "fmt"
 "os"
 "bufio"
 "strconv"

)

//Create a maintenance object

type maintenance struct{

     miles string
     name string
}

//Calculates Miles until 

func milesUntil(x int, y int) int {

 a := x % y
 
 return y - a

}

func main() {

  var  maint [8] maintenance
  
  //read the vehicle's maintenance information from file
  
  fileName, err := os.Open("TestVehicle.txt")

  if err != nil {

     panic(err)

   }

  scannerName := bufio.NewScanner(fileName)

  count := 0

  i := 0 

 //create an array for maintenance name and miles from the TestVehicle file

  for scannerName.Scan(){

      count = count+1

      if count %2 == 1{
          
      maint [i].name = scannerName.Text()

     } else {
          
      maint [i].miles = scannerName.Text()

      i=i+1

     }
 }

 //Get Vehicle Mileage.

 var currentMileage int

 fmt.Println("To check your upcoming reccomended maintenance, Enter your current mileage?")

 fmt.Scanln(&currentMileage)

 fmt.Println()
 fmt.Println("At your current Mileage of", currentMileage, "you have: ")

 //Loop that Prints Miles until each maintenance.

 for i:=0;i<=7;i++{

 s, err := strconv.Atoi(maint[i].miles)
	
 if err == nil{
   
 fmt.Println()
 fmt.Println("[",milesUntil(currentMileage, s), "Miles until: ", maint[i].name, "is due.","]")	 

 //If miles until is less than 100 then print this message to remind user not to forget to do it soon.	 
	 
 if milesUntil(currentMileage, s) <= 100{

 fmt.Println("[ DON'T FORGET! ]")

 }

 }

 }

}
