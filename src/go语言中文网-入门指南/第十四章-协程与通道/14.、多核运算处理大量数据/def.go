package main

func main() {

	/*
		流水线的优化
	*/

}

/*
流水线工作制
func SerialProcessData (in <- chan *Data, out <- chan *Data) {

	for data := range in {

		tmpA := PreprocessData(data)

		tmpB := ProcessStepA(tmpA)

		tmpC := ProcessStepB(tmpB)

		out <- PostProcessData(tmpC)

	}

}*/

/*
多协程工作制

func ParallelProcessData (in <- chan *Data, out <- chan *Data) {

    // make channels:

    preOut := make(chan *Data, 100)

    stepAOut := make(chan *Data, 100)

    stepBOut := make(chan *Data, 100)

    stepCOut := make(chan *Data, 100)

    // start parallel computations:

    go PreprocessData(in, preOut)

    go ProcessStepA(preOut, stepAOut)

    go ProcessStepB(stepAOut, stepBOut)

    go ProcessStepC(stepBOut, stepCOut)

    go PostProcessData(stepCOut, out)

}
*/
