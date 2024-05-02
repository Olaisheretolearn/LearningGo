package main 

import "fmt"

func main(){
  var publisher string = "DizzyBooks Publishing Inc."
  var writer string = "Tracey Hatchet"
  var artist string = "Jewel Tampson"
  var title string = "Mr. GoToSleep"

  var year , pageNumber int = 1997 , 14

  var grade float32 = 6.5

  fmt.Println(title, "written by " , writer , "drawn by " , artist , "published by" , publisher , "in the year ", year , "open the page " , pageNumber , "book has a grade of ", grade, "\n\n")

  title = "Epic Vol. 1"
  writer = "Ryan N. Shawn."
  artist = "Phoebe Paperclips"
  year = 2013
  pageNumber = 160
  grade = 9.0

    fmt.Println(title, "written by " , writer , "drawn by " , artist , "published by" , publisher , "in the year ", year , "open the page " , pageNumber , "book has a grade of ", grade)
}