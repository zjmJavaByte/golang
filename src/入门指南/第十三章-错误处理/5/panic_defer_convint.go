package main

import (
	"fmt"
	"math"
)

func main() {
	l := int64(15000)
	if i, err := IntFromInt64(l); err != nil {
		fmt.Printf("The conversion of %d to an int32 resulted in an error: %s", l, err.Error())
	} else {
		fmt.Printf("%d converted to an int32 is %d", l, i)
	}
	fmt.Println()
	l = int64(math.MaxInt32 + 15000)
	if i, err := IntFromInt64(l); err != nil {
		fmt.Printf("The conversion of %d to an int32 resulted in an error: %s", l, err.Error())
	} else {
		fmt.Printf("%d converted to an int32 is %d", l, i)
	}
}

func IntFromInt64(a int64) (i int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()
	i = ConvertInt64ToInt(a)
	return i, err
}

func ConvertInt64ToInt(a int64) int {
	if math.MinInt32 <= a && a <= math.MaxInt32 {
		return int(a)
	}
	panic(fmt.Sprintf("%d is out of the int32 range", a))
}
