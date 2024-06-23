// Note: In Go there are no `while` or `do/until`, just `for`.

package iteration

func Repeat(character string, count int) string {
	var repeated string

	for i := 0; i < count; i++ {
		repeated += character
	}

	return repeated
}

// Some iteration examples from https://gobyexample.com

// i := 1
// for i <= 3 {
//     fmt.Println(i)
//     i = i + 1
// }

// for j := 0; j < 3; j++ {
//     fmt.Println(j)
// }

// Looks a lot like how you would do this in Python
// for i := range 3 {
//     fmt.Println("range", i)
// }

// I found this one particularly interesting. It's basically the same as
// a while(true) loop, but using a for.
// for {
//     fmt.Println("loop")
//     break
// }

// for n := range 6 {
//     if n%2 == 0 {
//         continue
//     }
//     fmt.Println(n)
// }
