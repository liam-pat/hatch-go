package main

func test01(args ...int) {

}

func test(args ...int) {
	// all parameters
	test01(args...)

	// only first two parameters , args[0] args[1] , exclude args[2]
	test01(args[:2]...)

	// all parameters after [2] ,  args[2] args[3] args[4]....... , include args[2]
	test01(args[2:]...)
}

func main() {

}
