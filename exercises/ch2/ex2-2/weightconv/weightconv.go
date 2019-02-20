// Package weightconv performs Gram and Pound coversions
package weightconv

import (
	"fmt"
)

// Gram is in metric
type Gram float64

// Pound will be followed by the remaining ounces
type Pound float64

// Ounce is the remainder from pounds
type Ounce float64

func (g Gram) String() string    { return fmt.Sprintf("%gg", g) }
func (lbs Pound) String() string { return fmt.Sprintf("%glbs.", lbs) }
func (oz Ounce) String() string  { return fmt.Sprintf("%goz", oz) }
