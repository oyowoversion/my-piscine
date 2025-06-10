package piscine

// ForEach applies a function f on each element of an int slice a.
func ForEach(f func(int), a []int) {
	for _, v := range a {
		f(v)
	}
}

// You would also need a PrintNbr function if it's not already defined in your "piscine" package.
// For the purpose of testing the ForEach function, here's a simple PrintNbr implementation:
/*
import "strconv"
import "github.com/01-edu/z01" // Assuming z01 is used for printing characters

func PrintNbr(n int) {
    s := strconv.Itoa(n)
    for _, r := range s {
        z01.PrintRune(r)
    }
}
*/
