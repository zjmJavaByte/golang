package main

import (
	"fmt"
)

type Any1 interface{}
type EvalFunc1 func(Any1) (Any1, Any1)

func main() {
	fibFunc := func(state Any1) (Any1, Any1) {
		os := state.([]uint64)
		v1 := os[0]
		v2 := os[1]
		ns := []uint64{v2, v1 + v2}
		return v1, ns
	}
	fib := BuildLazyUInt64Evaluator1(fibFunc, []uint64{0, 1})

	for i := 0; i < 100; i++ {
		fmt.Printf("Fib nr %v: %v\n", i, fib())
	}
}

func BuildLazyEvaluator1(evalFunc EvalFunc1, initState Any1) func() Any1 {
	retValChan := make(chan Any1)
	loopFunc := func() {
		var actState Any1 = initState
		var retVal Any1
		for {
			retVal, actState = evalFunc(actState)
			retValChan <- retVal
		}
	}
	retFunc := func() Any1 {
		return <-retValChan
	}
	go loopFunc()
	return retFunc
}

func BuildLazyUInt64Evaluator1(evalFunc EvalFunc1, initState Any1) func() uint64 {
	ef := BuildLazyEvaluator1(evalFunc, initState)
	return func() uint64 {
		return ef().(uint64)
	}
}

/* Output:
Fib nr 0: 0
Fib nr 1: 1
Fib nr 2: 1
Fib nr 3: 2
Fib nr 4: 3
Fib nr 5: 5
Fib nr 6: 8
Fib nr 7: 13
Fib nr 8: 21
Fib nr 9: 34
*/
