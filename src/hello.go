/*
 *  Example code from the golang tutorial.
 */
package  main

import (
    "fmt"
    "math"
    "errors"
)


//This is the main go program entry point
func main() { 
  pointers()
  controlStructures()
  multipleReturnValues()
  testStructs1()
  testStructs2()
  testMethodsOnStruts()
}

func pointers() {

    i := 22
    j := 5
    var k int
    
    //do the print
    fmt.Printf("hello world\n")
    fmt.Println("i is : ", i)
    fmt.Println("address of i is : ", &i)
    fmt.Println("the value of address : ", &i, " is : ", *&i)
    var ptr = &k
    fmt.Println("the value of address : ", ptr, " is : ", *ptr)
    fmt.Println("j is : ", j) 
}

//control structures
func controlStructures() {
  a,b := 4,5
  if true {
    fmt.Println("True was executed")
  }
  if a<b {
    fmt.Println("a was less than b")
  } else if a > b {
    fmt.Println("a was greater than b")
  } else {
    fmt.Println("a equals b")
  }
  
  for i :=0; i<10; i++ {
    fmt.Println("i = ", i)
  }
  
  s := ""
  for ; s != "aaaa"; {
    fmt.Println(s)
    s = s+"a"
  }
  
  for i,j,k := 0,5,"a";i<3 && j <100 && k!= "aaaa"; i,j,k = i+1,j+1, k + "a" {
   fmt.Println("Value of i, j, k",i,j,k)
  }
  
  //range means 'for each of'
  //range of a slice
  z := [...]string{"a","b","c","d","e"}
  for i := range z {
    fmt.Println("Array item ", i , " is ", z[i])
  }
  //range of a map
  capitals := map[string] string {"France" : "Paris", "Italy" : "Rome", "Japan" : "Tokyo"}
  for key := range capitals {
    fmt.Println("Capital of ", key, " is ", capitals[key])
  }
  
  //switch and case
  i := 5
  switch i {
    case 4: fmt.Println("number 4")
    case 5: fmt.Println("number 5")
  }
  }
  
  //multiple return values
  func multipleReturnValues() {
     res, error := MySqrt(64)
     if error == nil {
       fmt.Println("Root ", 64, " is ", res)
     } else {
       fmt.Println("Error ", error)
     }
     res1, error1 := MySqrt(-10)
     if error1 == nil {
       fmt.Println("Root ", -10, " is ", res1)
     } else {
       fmt.Println("Error ", error1)
     }
  }
  
  //named return variables in function definition
  //named return variable are automatically zeroed
  func MySqrt(f float64) (float64, error) {
  
    if (f<0) {
      return float64(math.NaN()), errors.New("Cant sqrt a -ve number")
    }
    return math.Sqrt(f), nil 
  }
  
  func testStructs1() {
  
    type myStruct struct{
      i int
      j int
      s string
    }
    
    type Rectangle struct {
      length, width int
      name string
    }
    
    type Vehicle struct {
    
    }
    
    type VehicleCar struct {
    
    }
    
    r := Rectangle{}
    fmt.Println("Default rectangle is: ", r) 
    r1 := Rectangle{10, 20,"Rectangle 1"}                   //order init
    fmt.Println(r1)
    r2 := Rectangle{length:9, width:19,name: "Rectangle 2"} //named init
    fmt.Println(r2)
    pr := new (Rectangle) //new gives us a pointer
    (*pr).width = 6;      //dereference the pointer
    (pr).length = 8;      //no need to dereference the pointer
    (pr).name = "Rectangle 3";
    fmt.Println(*pr)      //dereference the pointer 
  } 
  
  func testStructs2() {
    type Kitchen struct {
      numOfPlates int
    }
    type House struct {
      Kitchen // anonymous field
      numOfRooms int
    }
    h := House{Kitchen{10},5}  // Note you cant mix named field and non named fiels ints
    fmt.Println(h)
    fmt.Println("House has this many rooms ", h.numOfRooms)
    fmt.Println("House has this many plates ", h.numOfPlates)  //anon filed has gone
    //can use h.Kitchen.numOfPlates
  }
  
  //methods on structs
  func testMethodsOnStruts() {
  
    
  
  }
  
  

