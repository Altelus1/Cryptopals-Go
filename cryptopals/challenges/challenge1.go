package challenges

import (
  "fmt"

  "cryptopals/cryptofunctions"
)

func Challenge1(){

  var x = cryptofunctions.Hex_to_bytes("41414141414141ff")
  fmt.Println(x)

}
