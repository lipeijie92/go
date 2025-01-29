package main

func fun1() {
	/*str := fun2(32,"lpj")
	println(str)*/
	var str string
	_ = fun2(32, "lpj")
	str = fun2(32, "lpj")
	println(str)
}

func fun2(age int, name string) string {
	return "这是返回值的名字" + name
}
