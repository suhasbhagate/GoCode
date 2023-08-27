package main

import(
	"fmt"
	"strconv"
)

func main(){
	num := 1024
	strStrConv := strconv.Itoa(num)
	fmt.Printf("strStrConv: %v Type: %T\n",strStrConv, strStrConv)
	numStrConv,_ := strconv.Atoi(strStrConv)
	fmt.Printf("numStrConv: %v Type: %T\n", numStrConv, numStrConv)

	strDecimal := strconv.FormatInt(int64(num), 10)
	fmt.Printf("strDecimal: %v Type: %T\n",strDecimal, strDecimal)

	strBinary := strconv.FormatInt(int64(num), 2)
	fmt.Printf("strBinary: %v Type: %T\n",strBinary, strBinary)

	strOctal := strconv.FormatInt(int64(num),8)
	fmt.Printf("strOctal: %v Type: %T\n", strOctal, strOctal)

	strHex := strconv.FormatInt(int64(num), 16)
	fmt.Printf("strHex: %v Type: %T\n",strHex,strHex)

	numDecimal,_ := strconv.ParseInt(strDecimal,10,64)
	fmt.Printf("numDecimal: %v Type:%T\n",numDecimal,numDecimal)

	numBinary,_ := strconv.ParseInt(strBinary,2,64)
	fmt.Printf("numBinary: %v Type:%T\n",numBinary,numBinary)

	numOctal,_ := strconv.ParseInt(strOctal,8,64)
	fmt.Printf("numOctal: %v Type:%T\n",numOctal,numOctal)

	numHex,_ := strconv.ParseInt(strHex,16,64)
	fmt.Printf("numHex: %v Type:%T\n",numHex,numHex)

	// strDecimal := strconv.FormatUint(uint64(num), 10)
	// fmt.Printf("strDecimal: %v Type: %T\n",strDecimal, strDecimal)

	// strBinary := strconv.FormatUint(uint64(num), 2)
	// fmt.Printf("strBinary: %v Type: %T\n",strBinary, strBinary)

	// strOctal := strconv.FormatUint(uint64(num),8)
	// fmt.Printf("strOctal: %v Type: %T\n", strOctal, strOctal)

	// strHex := strconv.FormatUint(uint64(num), 16)
	// fmt.Printf("strHex: %v Type: %T\n",strHex,strHex)

	// numDecimal,_ := strconv.ParseUint(strDecimal,10,64)
	// fmt.Printf("numDecimal: %v Type:%T\n",numDecimal,numDecimal)

	// numBinary,_ := strconv.ParseUint(strBinary,2,64)
	// fmt.Printf("numBinary: %v Type:%T\n",numBinary,numBinary)

	// numOctal,_ := strconv.ParseUint(strOctal,8,64)
	// fmt.Printf("numOctal: %v Type:%T\n",numOctal,numOctal)

	// numHex,_ := strconv.ParseUint(strHex,16,64)
	// fmt.Printf("numHex: %v Type:%T\n",numHex,numHex)

	floatVal := 20.35
	strFloat := strconv.FormatFloat(floatVal,'f',3,64)
	fmt.Printf("strFloat: %v Type: %T\n", strFloat,strFloat)

	numFloat,_ := strconv.ParseFloat(strFloat,64)
	fmt.Printf("numFloat: %v Type: %T\n", numFloat,numFloat)

	boolVal := true
	strBool := strconv.FormatBool(boolVal)
	fmt.Printf("strBool: %v Type: %T\n",strBool,strBool)

	valBool,_ := strconv.ParseBool(strBool)
	fmt.Printf("valBool: %v Type: %T\n",valBool,valBool)
}