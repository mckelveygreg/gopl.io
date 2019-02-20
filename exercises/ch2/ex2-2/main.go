// converts pounds and grams

package main

import (
	"fmt"
	"os"
	"strconv"

	"ch2/ex2-2/weightconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		w, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "weightErr: %v\n", err)
			os.Exit(1)
		}
		g := weightconv.Gram(w)
		lbs := weightconv.Pound(w)

		fmt.Printf("%s = %s, %s = %s\n",
			g, weightconv.GToLBS(g), lbs, weightconv.LBSToG(lbs))

	}
}
