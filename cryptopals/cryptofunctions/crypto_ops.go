package cryptofunctions

import "math"

func Hex_to_b64(hex string) string {

  var hex_bytes = Hex_to_bytes(hex)

  

}

func Hex_to_bytes(hex string) []byte {

  var byte_length = len(hex)/2
  var return_bytes []byte

  for count := 0; count < byte_length; count++ {
    var hex_byte = Hex_to_byte(hex[count*2:(count+1)*2])
    return_bytes = append(return_bytes, hex_byte)
  }

  return return_bytes

}

func Hex_to_byte(hex string) byte {

  var ret_int = 0

  for count := len(hex)-1; count >= 0; count-- {
    // If it is a-f
    if byte(hex[count]) >= byte('a') {
      ret_int += int(math.Pow(16, float64(len(hex)-1-count))) * int(byte(hex[count]) - byte('a') + 10)
    } else {
      ret_int += int(math.Pow(16, float64(len(hex)-1-count))) * int(byte(hex[count]) - byte('0'))
    }
  }

  return byte(ret_int)
}
