package useage

func Foo(channel chan string) {
	defer func() {
		if r := recover(); r != nil {
			channel <- "Foo failed"
			go Foo(channel)
		}
	}()
	panic("Foo failed")
}

func Bar(channel chan string) {
	defer func() {
		if r := recover(); r != nil {
			channel <- "Bar failed"
		}
	}()
	panic("Bar failed")
}
