func main() {



  x := 10

  fmt.Println(x)

  defer func() {

    fmt.Println("Deferring")

  }()



  if x == 10 {

    panic("Panicking")

  }



}