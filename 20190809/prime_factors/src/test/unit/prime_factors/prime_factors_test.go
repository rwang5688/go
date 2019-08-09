// prime_factors_unittest project prime_factors_test.god
package prime_factors_unittest

import (
	"app/prime_factors"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

// Setup
func setUp() *prime_factors.PrimeFactors {
	p := &prime_factors.PrimeFactors{}
	p.Init()
	return p
}

// 1st Test
func TestOne(t *testing.T) {
	// Arrange
	p := setUp()

	// Act
	list1 := makeIntList()
	s1 := convertIntListToString(list1)
	list2 := p.Generate(1)
	s2 := convertIntListToString(list2)

	// Assert
	if s1 != s2 {
		t.Error(`Expected list is ` + s1 + ` and generated list is ` + s2)
	}
}

// 2nd Test
func TestTwo(t *testing.T) {
	// Arrange
	p := setUp()

	// Act
	list1 := makeIntList(2)
	s1 := convertIntListToString(list1)
	list2 := p.Generate(2)
	s2 := convertIntListToString(list2)

	// Assert
	if s1 != s2 {
		t.Error(`Expected list is ` + s1 + ` and generated list is ` + s2)
	}
}

// 3rd Test
func TestThree(t *testing.T) {
	// Arrange
	p := setUp()

	// Act
	list1 := makeIntList(3)
	s1 := convertIntListToString(list1)
	list2 := p.Generate(3)
	s2 := convertIntListToString(list2)

	// Assert
	if s1 != s2 {
		t.Error(`Expected list is ` + s1 + ` and generated list is ` + s2)
	}
}

// 4th Test
func TestFour(t *testing.T) {
	// Arrange
	p := setUp()

	// Act
	list1 := makeIntList(2, 2)
	s1 := convertIntListToString(list1)
	list2 := p.Generate(4)
	s2 := convertIntListToString(list2)

	// Assert
	if s1 != s2 {
		t.Error(`Expected list is ` + s1 + ` and generated list is ` + s2)
	}
}

// 5th Test
func TestNine(t *testing.T) {
	// Arrange
	p := setUp()

	// Act
	list1 := makeIntList(3, 3)
	s1 := convertIntListToString(list1)
	list2 := p.Generate(9)
	s2 := convertIntListToString(list2)

	// Assert
	if s1 != s2 {
		t.Error(`Expected list is ` + s1 + ` and generated list is ` + s2)
	}
}

// 6th Test
func TestSeventy(t *testing.T) {
	// Arrange
	p := setUp()

	// Act
	list1 := makeIntList(2, 5, 7)
	s1 := convertIntListToString(list1)
	list2 := p.Generate(70)
	s2 := convertIntListToString(list2)

	// Assert
	if s1 != s2 {
		t.Error(`Expected list is ` + s1 + ` and generated list is ` + s2)
	}
}

// 7th Test
func TestOneFourty(t *testing.T) {
	// Arrange
	p := setUp()

	// Act
	list1 := makeIntList(2, 2, 5, 7)
	s1 := convertIntListToString(list1)
	list2 := p.Generate(140)
	s2 := convertIntListToString(list2)

	// Assert
	if s1 != s2 {
		t.Error(`Expected list is ` + s1 + ` and generated list is ` + s2)
	}
}

// Helper functions
func makeIntList(ints ...int) []int {
	// debugging: print out var args
	fmt.Println(ints)

	// loop through var args
	// throw away index and extract value
	intList := []int{}
	for _, n := range ints {
		intList = append(intList, n)
	}

	// debugging: print out int list
	fmt.Println(intList)

	return intList
}

func convertIntListToString(intList []int) string {
	// Create a string slice using strconv.Itoa.
	// ... Append strings to it.
	valuesText := []string{}
	for i := range intList {
		value := intList[i]
		text := strconv.Itoa(value)
		valuesText = append(valuesText, text)
	}

	// Join our string slice.
	result := strings.Join(valuesText, "*")

	return result
}
