package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  seed := time.Now().UnixNano()
  rand.Seed(seed);
  
  var isHeistOn bool = true;
   eludedGuards := rand.Intn(100);

   if(eludedGuards >= 50){
    fmt.Println(eludedGuards);
    fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step");
   } else {
    isHeistOn = false;
     fmt.Println(eludedGuards);
    fmt.Println("Plan a better disguise next time?")
   }

   openedVault:=rand.Intn(100);
   if (openedVault >= 70 && isHeistOn){
     fmt.Println(openedVault);
       fmt.Println("Grab grab grab and go!")
   } else if(isHeistOn){
    isHeistOn = false;
     fmt.Println("Vault could not be opened")
   } 

   leftSafely:=rand.Intn(5)

   if(isHeistOn){
    switch (leftSafely) {
      case 0 : isHeistOn = false 
      fmt.Println("Oops Unsuccessful");
      case 1:
       isHeistOn = false 
          fmt.Print("Turns out vault doors don't open from the inside...")  
      case 2: 
      isHeistOn = false;
      fmt.Print("Teammate trips and dies");
      case 4:
      isHeistOn =false;
      fmt.Println("When did they start raising dogs in vaults??")
      default :
      fmt.Println("Start the getaway car")
    }
   }

   if(isHeistOn){
    amtStolen := 10000 + rand.Intn(1000000)
    fmt.Printf("So much money %d /n", amtStolen )
   }


fmt.Println(isHeistOn)


}