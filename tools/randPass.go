//生成一定长度的若干个随机密码
package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var pwdCount *int = flag.Int("n", 10, "How many password")
var pwdLength *int = flag.Int("l", 20, "How long a password")

var pwdCodes = [...]byte{
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
	'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M',
	'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
	'`', '1', '2', '3', '4', '5', '6', '7', '8', '9', '0', '-', '=',
	'~', '!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '_', '+',
	'[', ']', ',', '.', '/', '{', '}', '|', ':', '"', '<', '>', '?',
	'\\', ';', '\'',
}

func main() {
	flag.Parse()

	//var r *rand.Rand = rand.New(rand.NewSource(time.Nanosecond))
	var r *rand.Rand = rand.New(rand.NewSource(int64(time.Now().Nanosecond())))

	for i := 0; i < *pwdCount; i++ {
		var pwd []byte = make([]byte, *pwdLength)

		for j := 0; j < *pwdLength; j++ {
			index := r.Int() % len(pwdCodes)

			pwd[j] = pwdCodes[index]
		}

		fmt.Printf("%s\r\n", string(pwd))
	}
}
