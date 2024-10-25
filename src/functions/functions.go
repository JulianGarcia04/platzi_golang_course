package functions

import (
	"fmt"
	"slices"
	"strings"
)

func IsPalindrome (text string) bool {
	// var textProxy string;

	// for i := len(text) - 1; i >= 0; i-- {
	// 	textProxy += string(text[i])
	// }

	// return textProxy == text

	lowerCaseText := strings.ToLower(text)

	splitText := strings.Split(lowerCaseText, "")

	slices.Reverse(splitText)

	reversedText := strings.Join(splitText, "")

	return reversedText == lowerCaseText
}

// the size parameter must be in cm
func GetSquareArea (size uint) uint {
	return size * size
}

// base and height parameter must be in cm
func GetTriangleArea (base, height uint) uint {
	return (base * height)/2
}

func GetFullName (names ...string) string {
	var fullname string
	for _, name := range names {
		fullname += string(name + " ")
	}
	return fullname
}

func Functions () {
	squareArea := GetSquareArea(10);
	fmt.Printf("El area del cuadrado es %d centimetros.\n", squareArea);

	triangleArea := GetTriangleArea(20, 30);
	fmt.Printf("El area del triangulo es %d centimetros.\n", triangleArea);

	names := GetFullName("Julian", "Garcia")

	fmt.Println(names)

}