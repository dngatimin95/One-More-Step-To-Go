package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	digit := [10]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	lowerCase := [26]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	upperCase := [26]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	symbols := [21]string{"@", "#", "$", "%", "=", ":", "?", ".", "/", "|", "~", ">", "<", "*", "(", ")", "&", "!", "^", "_", "+"}

	var noPass int
	var passLength int
	rand.Seed(time.Now().UnixNano())

	for {
		fmt.Println("How many passwords do you want to generate?")
		fmt.Scanln(&noPass)
		if noPass < 1 {
			fmt.Println("Invalid number. Please try again.")
		} else {
			break
		}
	}

	for {
		fmt.Println("\nHow long should the passwords be? (The longer, the better)")
		fmt.Scanln(&passLength)
		if passLength < 1 {
			fmt.Println("Invalid number. Please try again.")
		} else {
			break
		}
	}

	fmt.Println("\nHere are some passwords you can use:")

	for x := 0; x < noPass; x++ {
		var pass string
		each := passLength / 4
		remain := passLength % 4

		if passLength < 4 {
			for i := 0; i < passLength; i++ {
				index := rand.Intn(26)
				pass += string(upperCase[index])
			}
		} else {
			for j := 0; j < each; j++ {
				indexNum := rand.Intn(10)
				pass += string(digit[indexNum])

				indexLow := rand.Intn(26)
				pass += string(lowerCase[indexLow])

				indexUpp := rand.Intn(26)
				pass += string(upperCase[indexUpp])

				indexSym := rand.Intn(21)
				pass += string(symbols[indexSym])
			}

			for k := 0; k < remain; k++ {
				listChoice := rand.Intn(4)
				switch listChoice {
				case 0:
					indexNum := rand.Intn(10)
					pass += string(digit[indexNum])
				case 1:
					indexLow := rand.Intn(26)
					pass += string(lowerCase[indexLow])
				case 2:
					indexUpp := rand.Intn(26)
					pass += string(upperCase[indexUpp])
				case 3:
					indexSym := rand.Intn(21)
					pass += string(symbols[indexSym])
				}
			}
		}
		passRune := []rune(pass)
		rand.Shuffle(len(passRune), func(i, j int) {
			passRune[i], passRune[j] = passRune[j], passRune[i]
		})
		fmt.Println(string(passRune))
	}
	fmt.Println("")
}
