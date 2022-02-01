package calculator_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/alcor67/repo-Go-level-1-home-work/projectcalculator/calculator"
	"github.com/stretchr/testify/assert"
)

func TestCalcSmoke(t *testing.T) { // входная переменная типа *testing.T

	Expected, err := calculator.Calc("+", 100, 200)
	if err != nil {
		t.Errorf("Calc(+ 100 200) expected to return no error, but received: %s", err)
	}

	expected := 300.0
	if Expected != expected {
		t.Errorf("Calc(+ 100 200) expected to return %.2f, but receive %.2f", expected, Expected)
	}
}

func ExampleCalcPlus() {
	result, err := calculator.Calc("+", 300, 500)
	fmt.Println(result, err)
	// Output: 800 <nil>
}

func ExampleCalcMinus() {
	result, err := calculator.Calc("-", 800, 300)
	fmt.Println(result, err)
	// Output: 500 <nil>
}

func ExampleCalcMult() {
	result, err := calculator.Calc("*", 30, 50)
	fmt.Println(result, err)
	// Output: 1500 <nil>
}

func ExampleCalcDiv() {
	result, err := calculator.Calc("/", 80, 5)
	fmt.Println(result, err)
	// Output: 16 <nil>
}

func ExampleCalcRemDiv() {
	result, err := calculator.Calc("%", 83, 5)
	fmt.Println(result, err)
	// Output: 3 <nil>
}

func ExampleCalcExp() {
	result, err := calculator.Calc("**", 4, 2)
	fmt.Println(result, err)
	// Output: 16 <nil>
}

func TestTestifyCalc(t *testing.T) {

	var X float64 = 4
	var Y float64 = 2
	Op := "**"
	var Expected float64 = 16

	result, err := calculator.Calc(Op, X, Y)
	fnCallStr := fmt.Sprintf("Calc (%.2f %s %.2f)", X, Op, Y)
	if err != nil {
		t.Errorf("%s expected to return no error, but received: %s", fnCallStr, err)
		//continue
	}

	assert.Equal(t, result, Expected, "they should be equal")

}

func TestCalcTableOk(t *testing.T) {
	//fmt.Println("запускаем TestCalcTableOk")
	type call struct {
		Op             string
		X, Y, Expected float64
	}

	calls := []call{
		//Plus
		{"+", 100, 50, 150},
		{"+", 2, 2, 4},
		{"+", 0, 0, 0},
		{"+", 0, 5, 5},
		{"+", 100, 0, 100},
		{"+", -100, 100, 0},

		//Minus
		{"-", 100, 50, 50},
		{"-", 2, 2, 0},
		{"-", 0, 0, 0},
		{"-", 0, 5, -5},
		{"-", 100, 0, 100},
		{"-", -100, 100, -200},

		//Multiply
		{"*", 100, 50, 5000},
		{"*", 2, 2, 4},
		{"*", 0, 0, 0},
		{"*", 0, 5, 0},
		{"*", 100, 0, 0},
		{"*", -100, 100, -10000},
		{"*", -100, -100, 10000},

		//Divide
		{"/", 100, 50, 2},
		{"/", 2, 2, 1},
		{"/", 0, 5, 0},
		{"/", 0, -5, 0},
		{"/", -100, 100, -1},
		{"/", -100, -100, 1},

		//remainder of division
		{"%", 100, 50, 0},
		{"%", 10.5, 3, 1},
		{"%", -10.5, 3, -1},
		{"%", 23.5, 5, 3},
		{"%", 0, 3, 0},
		{"%", 5, 8, 5},

		//Exponentiation
		{"**", 4, 2, 16},
		{"**", 4, .5, 2},
		{"**", 4, 0, 1},
		{"**", -4, 0, 1},
		{"**", 4, -2, 0.0625},
		{"**", -2, -2, 0.25},
		{"**", 0, 2, 0},
		{"**", 0, 0, 1},
		{"**", 4, -.3, math.Pow(4, -.3)},

		//todo: make tests on values of NaN (special cases)
		//{"**", -4, 1.4, math.Pow(-4, 1.4)},
		//{"**", -4, -.4, math.Pow(-4, -.4)},
		//todo: make tests on values of Inf (edge cases)
		//{"**", 0, -.1, math.Pow(0, -.1)},
		//{"**", 0, -2, math.Pow(0, -2)},

		//todo: make tests on Maximum values of float (edge cases)
		//{"+", math.MaxFloat64, 100, math.MaxFloat64 + 100},
	}

	for _, c := range calls {
		fnCallStr := fmt.Sprintf("Calc (%.2f %s %.2f)", c.X, c.Op, c.Y)

		result, err := calculator.Calc(c.Op, c.X, c.Y)
		if err != nil {
			t.Errorf("%s expected to return no error, but received: %s", fnCallStr, err)
			//continue
		}
		if result != c.Expected {
			t.Errorf("%s expected to return %.2f, but receive %.2f", fnCallStr, c.Expected, result)
		}
	}
}

func TestCalcTableFailure(t *testing.T) {
	//fmt.Println("запускаем TestCalcTableFailure")
	type call struct {
		Op   string
		X, Y float64
	}

	calls := []call{
		//Divide by zero
		{"/", 100, 0},
		{"/", 0, 0},
		{"%", 5, .5},
		{"%", 5, -.5},
		{"%", .5, -.5},
		//Not supported operation
		{"a", 0, 0},
		{"sin", 0, 0},
		{"cos", 0, 0},
		//Result is NaN (special cases)
		{"**", -4, 1.4},
		{"**", -4, -.4},
		//Result is Inf (edge cases)
		{"**", 0, -.1},
		{"%", 5, -.5},
	}

	for _, c := range calls {
		fnCallStr := fmt.Sprintf("Calc (%.2f %s %.2f)", c.X, c.Op, c.Y)

		_, err := calculator.Calc(c.Op, c.X, c.Y)
		if err == nil {
			t.Errorf("%s expected to return error, but received no error", fnCallStr)
			//continue
		}
	}
}
