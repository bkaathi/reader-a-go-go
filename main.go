 package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
  "rand"
)

func main() {

  reader := bufio.NewReader(os.Stdin)
  fmt.Println("hey girl!")
  fmt.Println("------------------------")

  for {
    fmt.Print("->")
    text, _ := reader.ReadString('\n')
    //convert CRLF to LF
    text = strings.Replace(text, "\n","",-1)

    if strings.Contains("hi",text) == true {
      fmt.Println("Hello, yourselff")
  //    break;
    }
    if strings.Compare("bye",text) ==0 {
      fmt.Println("goodbye")
      break;
    }
    if strings.Compare("help",text)==true {
      fmt.Println("all you need is within you now!")
      fmt.Println(strings.Compare("help",text))
    }

    if strings.Compare("hey",text)>=0 {
      fmt.Println("Hello, yourself")
      fmt.Println(strings.Compare("help",text))
    }

    if strings.Compare("whats up?",text)==true {
      fmt.Println("you wanna know whats UP???")
    }

  }

}
