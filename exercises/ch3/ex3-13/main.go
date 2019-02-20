// package main

// import (
// 	"ch3/ex3-13/insertComma"
// 	"fmt"
// 	"strconv"
// 	"strings"
// )

// const (
// 	_ = 1 << (10 * iota)
// 	KiB
// 	MiB
// 	GiB
// 	TiB
// 	PiB
// 	EiB
// 	ZiB
// 	YiB
// )

// func main() {
// 	fmt.Println(
// 		KiB,
// 		MiB,
// 		GiB,
// 		TiB,
// 		PiB,
// 		EiB,
// 		ZiB,
// 		YiB,
// 	)
// 	var testArray = [...]int{
// 		KiB,
// 		MiB,
// 		GiB,
// 		TiB,
// 		PiB,
// 		EiB,
// 		ZiB,
// 		YiB,
// 	}
// 	for _, num := range testArray {
// 		prettyPrint(num)
// 	}
// }

// func prettyPrint(iB int) {
// 	var length = len(strconv.Itoa(iB))
// 	var answer = "1" + strings.Repeat("0", (length-1))
// 	println(insertComma.Comma(answer), "bytes")
// }

package main

import (
	"fmt"
)

const (
	_ = 1 << (10 * iota)
	KiB
	MiB
	GiB
	TiB
	PiB
	EiB
	ZiB
	YiB
)

const (
	k  = 1000
	KB = k
	MB = k * KB
	GB = k * MB
	TB = k * GB
	PB = k * TB
	EB = k * PB
	ZB = k * EB
	YB = k * ZB
)

func main() {

	fmt.Println(KiB, MiB, GiB, TiB, PiB, EiB, float64(ZiB), float64(YiB))
	fmt.Println(KB, MB, GB, TB, PB, EB, float64(ZB), float64(YB))
}
