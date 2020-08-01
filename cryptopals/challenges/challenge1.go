package challenges

import (
  "fmt"
  "cryptopals/cryptofunctions"
)

func Challenge1Set1(){

  var b64 = cryptofunctions.HexToB64("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
  fmt.Println("Result: ", b64)

}

func Challenge1Set2(){

  var byteArr = cryptofunctions.HexToBytes("1c0111001f010100061a024b53535009181c")
  var xorArr = cryptofunctions.HexToBytes("686974207468652062756c6c277320657965")
  var xoredArr = cryptofunctions.BytesXor(byteArr, xorArr)
  fmt.Println("Xored Bytes: ",cryptofunctions.BytesToHex(xoredArr))
}
