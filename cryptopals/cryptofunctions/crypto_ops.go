package cryptofunctions

import (
  "math"
  "encoding/binary"
)

func Hex_to_b64(hex string) string {

  var hex_bytes = Hex_to_bytes(hex)

  return bytes_to_b64(hex_bytes)

}

func bytes_to_b64(byte_arr []byte) string {

//  var add_bytes = []byte{0, 0}
  var orig_byte_arr_len = len(byte_arr)
  var tribyte_div_length = orig_byte_arr_len - (orig_byte_arr_len % 3)

//  byte_arr = append(byte_arr, add_bytes[0:tribyte_div_length+3-orig_byte_arr_len]...)

  var res_b64 []byte
  var count = 0

  for count < len(byte_arr) {

    var tmp = make([]byte, 3)

    copy(tmp, byte_arr[count:count+3])

    var tribyte = binary.BigEndian.Uint32(append([]byte{0}, tmp...))

    if count+3 > tribyte_div_length {
      tmp = tmp[:orig_byte_arr_len-tribyte_div_length]
    }

    var quadcount = 3

    for quadcount >= 3-len(tmp) {
      var sixbit = byte((tribyte >> (quadcount * 6)) & 0x3f)

      var result_char byte

      if sixbit >= 0 && sixbit <= 25 {
        result_char = byte(sixbit+byte('A'))
      } else if sixbit >= 26 && sixbit <= 51 {
        result_char = byte(sixbit+byte('a')-26)
      } else if sixbit >= 52 && sixbit <= 61 {
        result_char = byte(sixbit+byte('0')-52)
      } else if sixbit == 62 {
        result_char = '+'
      } else {
        result_char = '/'
      }

      res_b64 = append(res_b64, result_char)
      quadcount -= 1

    }
    count += 3
  }

  /*
    Appending equals signs  to base 64
  */
  res_b64 = append(res_b64, []byte{'=','='}[:(tribyte_div_length+3-orig_byte_arr_len) % 3]...)

  return string(res_b64)
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
