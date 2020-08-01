package challenges

import (
  "fmt"

  "cryptopals/cryptofunctions"
)

func Challenge1Set1(){

  var b64 = cryptofunctions.HexToB64("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
  fmt.Println("Result: ", b64)

}

func Challeng1Set2(){

}
