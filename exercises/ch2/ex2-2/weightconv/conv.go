// Conversions from grams to lbs and lbs to grams
// for weightconv package

package weightconv

// GToLBS converst weight in grams to pounds
func GToLBS(g Gram) Pound { return Pound(g * 0.0022) }

//LBSToG converts weight in pounds to grams
func LBSToG(lbs Pound) Gram { return Gram(lbs * 453.6) }

// need to deal with ounces
