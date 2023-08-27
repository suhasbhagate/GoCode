package main

import (
	"fmt"
	"strconv"
)

func main() {
	//Integer to ASCII
	num1 := 128
	str1 := strconv.Itoa(num1)
	fmt.Printf("str1 Type: %T Value: %v\n", str1, str1)

	//ASCII to Integer
	str2 := "128"
	num2, _ := strconv.Atoi(str2)
	fmt.Printf("num2 Type: %T Value: %v\n", num2, num2)

	/////////////////////////////////////////////////////////////////////////////////////////////////////

	//int64 to other Numeric Data Types Decimal, Binary, Octal, Hexadecimal Strings
	num := 1024
	strDecimal := strconv.FormatInt(int64(num), 10)
	fmt.Printf("Num: %v Decimal: %v Type: %T\n", num, strDecimal, strDecimal)

	strBinary := strconv.FormatInt(int64(num), 2)
	fmt.Printf("Num: %v Binary: %v Type: %T\n", num, strBinary, strBinary)

	strOctal := strconv.FormatInt(int64(num), 8)
	fmt.Printf("Num: %v Octal: %v Type: %T\n", num, strOctal, strOctal)

	strHexadecimal := strconv.FormatInt(int64(num), 16)
	fmt.Printf("Num: %v HexaDecimal: %v Type: %T\n", num, strHexadecimal, strHexadecimal)

	//////////////////////////////////////////////////////////////////////////////////////////////////////

	//String to Integer Data Types (int64) - Decimal, Binary, Octal, Hexadecimal
	numDecimal, _ := strconv.ParseInt(strDecimal, 10, 64)
	fmt.Printf("strDecimal: %v num: %v Type: %T\n", strDecimal, numDecimal, numDecimal)

	numBinary, _ := strconv.ParseInt(strBinary, 2, 64)
	fmt.Printf("strBinary: %v num: %v Type: %T\n", strBinary, numBinary, numBinary)

	numOctal, _ := strconv.ParseInt(strOctal, 8, 64)
	fmt.Printf("strOctal: %v num: %v Type: %T\n", strOctal, numOctal, numOctal)

	numhexadecimal, _ := strconv.ParseInt(strHexadecimal, 16, 64)
	fmt.Printf("strHexadecimal: %v num: %v Type: %T\n", strHexadecimal, numhexadecimal, numhexadecimal)

	//////////////////////////////////////////////////////////////////////////////////////////////////////

	//String to Unsigned Integer Data Types (uint64) - Decimal, Binary, Octal, Hexadecimal
	numDecimal1, _ := strconv.ParseUint(strDecimal, 10, 64)
	fmt.Printf("strDecimal: %v num: %v Type: %T\n", strDecimal, numDecimal1, numDecimal1)

	numBinary1, _ := strconv.ParseUint(strBinary, 2, 64)
	fmt.Printf("strBinary: %v num: %v Type: %T\n", strBinary, numBinary1, numBinary1)

	numOctal1, _ := strconv.ParseUint(strOctal, 8, 64)
	fmt.Printf("strOctal: %v num: %v Type: %T\n", strOctal, numOctal1, numOctal1)

	numhexadecimal1, _ := strconv.ParseUint(strHexadecimal, 16, 64)
	fmt.Printf("strHexadecimal: %v num: %v Type: %T\n", strHexadecimal, numhexadecimal1, numhexadecimal1)

	//////////////////////////////////////////////////////////////////////////////////////////////////////

	//Bool to String
	strBool := strconv.FormatBool(true)
	fmt.Printf("true: strBool: %v Type: %T\n", strBool, strBool)

	//String to Bool
	valBool, _ := strconv.ParseBool(strBool)
	fmt.Printf("strBool: %v valBool: %v Type: %T\n", strBool, valBool,valBool)

	//////////////////////////////////////////////////////////////////////////////////////////////////////

	//Float to String
	floatVal := 10.20
	strFloat := strconv.FormatFloat(floatVal,'f',3,64)
	fmt.Printf("floatVal: %v strFloat: %v Type: %T\n",floatVal,strFloat,strFloat)

	//String to Float
	valFloat, _ := strconv.ParseFloat(strFloat,64)
	fmt.Printf("strFloat: %v, valFloat: %v Type: %T\n",strFloat,valFloat,valFloat)
}
