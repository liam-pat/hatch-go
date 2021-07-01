package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type IP uint32

func (ip IP) String() string {
	var bf bytes.Buffer
	for i := 1; i <= 4; i++ {
		bf.WriteString(strconv.Itoa(int((ip >> ((4 - uint(i)) * 8)) & 0xff)))
		if i != 4 {
			bf.WriteByte('.')
		}
	}
	return bf.String()
}

func (ip IP) ipInt32ToString() string {
	bitStr := ""
	ipStr := ""
	for ; ip > 0; ip /= 2 {
		lsb := ip % 2
		bitStr = strconv.Itoa(int(lsb)) + bitStr
	}
	if len(bitStr) < 32 {
		bitStr = strings.Repeat("0", 32-len(bitStr)) + bitStr
	}

	for i := 1; i <= 4; i++ {
		s := bitStr[(i-1)*8 : i*8]
		parseInt, err := strconv.ParseInt(s, 2, 0)
		if err != nil {
			return ""
		}

		if i != 4 {
			ipStr = ipStr + strconv.Itoa(int(parseInt)) + "."
		} else {
			ipStr = ipStr + strconv.Itoa(int(parseInt))
		}
	}

	return ipStr
}

func main() {
	var ipInt IP = 3232248321

	str := ipInt.ipInt32ToString()
	fmt.Println(str)
}
