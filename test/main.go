package main



type userInfo struct {
	name string
	groups  *userInfo

}

func (user *userInfo)GEt()  {

}

func main()  {

	/*
	在以上代码中，内部类型userInfo的初始化使用结构字面量完成的。通过内部类型的名字userInfo可以访问内部类型的字段或方法。
	对外部类型client来说，内部类型总是存在的。借助内部类型提升，notify方法可以直接通过外部类型的变量c来访问，实际上所谓的内部类型的提升只是Go编译器帮我们完成了一次间接查找。
	Go编译器内部对指针进行了优化，c.user.notify()和c.notify()实际上内部执行都是&(c.user).notify(), 因为内部类型的方法是通过指针接收者定义的
	*/
}


