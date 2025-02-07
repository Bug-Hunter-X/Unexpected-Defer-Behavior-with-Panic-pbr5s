func main() {

  x := 10

  fmt.Println(x)

  defer func() {

    fmt.Println("Deferring")

    if r := recover(); r != nil {
        fmt.Println("Recovered from panic:", r)
    }
  }()



  if x == 10 {

    panic("Panicking")

  }



} 