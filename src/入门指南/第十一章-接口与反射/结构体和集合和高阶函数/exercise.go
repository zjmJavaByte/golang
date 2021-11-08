package main

import "fmt"

type Any2 interface{}

type Car2 struct {
	Model        string
	Manufacturer string
	BuildYear    int
}

type Cars2 []*Car2

func main() {
	// make some cars:
	ford := &Car2{"Fiesta", "Ford", 2008}
	bmw := &Car2{"XL 450", "BMW", 2011}
	merc := &Car2{"D600", "Mercedes", 2009}
	bmw2 := &Car2{"X 800", "BMW", 2008}
	allCars := Cars2{ford, bmw, merc, bmw2}

	allNewBMWs := allCars.FindAll(func(car *Car2) bool {
		return (car.Manufacturer == "BMW") && (car.BuildYear > 2010)
	})
	fmt.Println("AllCars: ", allCars)
	fmt.Println("New BMWs: ", allNewBMWs)

	manufacturers := []string{"Ford", "Aston Martin", "Land Rover", "BMW", "Jaguar"}
	appender2, sortedCars := MakeSortedAppender2(manufacturers)
	allCars.Process(appender2)
	fmt.Println("Map sortedCars: ", sortedCars)
	BMWCount := len(sortedCars["BMW"])
	fmt.Println("We have ", BMWCount, " BMWs")
}

/*
在定义所需功能时我们可以利用函数作为（其它函数的）参数来使用高阶函数，例如：
1）定义一个“通用”的 Process() 函数，它接收一个作用于每一辆 car 的 f 函数作参数：
*/
//通用函数
func (cs Cars2) Process(f func(car *Car2)) {
	for _, c := range cs {
		f(c)
	}
}

/*
2）在上面的基础上，实现一个查找函数来获取子集合，并在 Process() 中传入一个闭包执行（这样就可以访问局部切片 cars）
*/
func (cs Cars2) FindAll(f func(car *Car2) bool) Cars2 {
	cars := make([]*Car2, 0)
	for _, c := range cs {
		if f(c) {
			cars = append(cars, c)
		}
	}
	return cars
}

/*
3）实现 Map 功能，产出除 car 对象以外的东西
*/
func (cs Cars2) Map(f func(car *Car2) Any2) []Any2 {
	result := make([]Any2, 0)
	ix := 0
	cs.Process(func(car *Car2) {
		result[ix] = f(car)
		ix++
	})
	return result
}

/*
4）我们也可以根据入参返回不同的函数。也许我们想根据不同的厂商添加汽车到不同的集合，但是这可能会是多变的。所以我们可以定义一个函数来产生特定的添加函数和 map 集
*/

func MakeSortedAppender2(arr []string) (func(car *Car2), map[string][]*Car2) {
	m := make(map[string][]*Car2)
	for _, v := range arr {
		m[v] = make([]*Car2, 0)
	}
	m["default"] = make([]*Car2, 0)

	f := func(car *Car2) {
		if _, ok := m[car.Manufacturer]; ok {
			m[car.Manufacturer] = append(m[car.Manufacturer], car)
		} else {
			m["default"] = append(m["default"], car)
		}
	}
	return f, m
}
