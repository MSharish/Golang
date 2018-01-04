package main

import "fmt"

func main() {

  var n int;
  f1 := 0;
  f2 := 1;
  count := 1;
  sum := 0;
	fmt.Println("Enter the number of fib series\n")
  fmt.Scanf("%d",&n)

  if n == 0 {
    fmt.Println("No fib series for 0 elements\n")
  }else if n == 1 {
    fmt.Println("0\n")
  }else {
    for count <= n {

      sum = f1+f2;
      fmt.Printf("%d\t",sum)
      f1 = f2;
      f2 = sum;
      count = count + 1;

    }
  }
}
