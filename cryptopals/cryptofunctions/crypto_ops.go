package cryptofunctions

import (
  "math"
  "encoding/binary"
)

func HexToB64(hex string) string {

  var hexBytes = HexToBytes(hex)

  return BytesToB64(hexBytes)

}

func BytesToB64(byteArr []byte) string {

//  var add_bytes = []byte{0, 0}
  var origByteArrLen = len(byteArr)
  var triByteDivLength = origByteArrLen - (origByteArrLen % 3)

//  byteArr = append(byteArr, add_bytes[0:triByteDivLength+3-origByteArrLen]...)

  var resB64 []byte
  var count = 0

  for count < len(byteArr) {

    var tmp = make([]byte, 3)

    copy(tmp, byteArr[count:count+3])

    var tribyte = binary.BigEndian.Uint32(append([]byte{0}, tmp...))

    if count+3 > triByteDivLength {
      tmp = tmp[:origByteArrLen-triByteDivLength]
    }

    var quadcount = 3

    for quadcount >= 3-len(tmp) {
      var sixbit = byte((tribyte >> (quadcount * 6)) & 0x3f)

      var resultChar byte

      if sixbit >= 0 && sixbit <= 25 {
        resultChar = byte(sixbit+byte('A'))
      } else if sixbit >= 26 && sixbit <= 51 {
        resultChar = byte(sixbit+byte('a')-26)
      } else if sixbit >= 52 && sixbit <= 61 {
        resultChar = byte(sixbit+byte('0')-52)
      } else if sixbit == 62 {
        resultChar = '+'
      } else {
        resultChar = '/'
      }

      resB64 = append(resB64, resultChar)
      quadcount -= 1

    }
    count += 3
  }

  /*
    Appending equals signs  to base 64
  */
  resB64 = append(resB64, []byte{'=','='}[:(triByteDivLength+3-origByteArrLen) % 3]...)

  return string(resB64)
}

func HexToBytes(hex string) []byte {

  var byteLength = len(hex)/2
  var returnBytes []byte

  for count := 0; count < byteLength; count++ {
    var hex_byte = HexToByte(hex[count*2:(count+1)*2])
    returnBytes = append(returnBytes, hex_byte)
  }

  return returnBytes

}

func HexToByte(hex string) byte {

  var retInt = 0

  for count := len(hex)-1; count >= 0; count-- {
    // If it is a-f
    if byte(hex[count]) >= byte('a') {
      retInt += int(math.Pow(16, float64(len(hex)-1-count))) * int(byte(hex[count]) - byte('a') + 10)
    } else {
      retInt += int(math.Pow(16, float64(len(hex)-1-count))) * int(byte(hex[count]) - byte('0'))
    }
  }

  return byte(retInt)
}

func ByteToHex(inByte byte) string {

  var lowHalf = inByte & 0x0f
  var upperHalf = (inByte & 0xf0) >> 4

  var table = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}

  return string([]byte{table[upperHalf], table[lowHalf]})

}

func BytesToHex(byteArr []byte) string {

  var retHexStr string

  for count := 0; count < len(byteArr); count++ {
    retHexStr += ByteToHex(byteArr[count])
  }

  return retHexStr

}

func BytesXor(byteArr []byte, xorArr []byte) []byte {

  var retXorBytes []byte

  for count := 0; count < len(byteArr); count++ {
    retXorBytes = append(retXorBytes, byteArr[count] ^ xorArr[count % (len(xorArr))])
  }

  return retXorBytes

}
