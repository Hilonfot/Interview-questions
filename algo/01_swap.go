package main

func swap(a,b interface{})  (interface{},interface{}){
	a ,b = b,a
	return a,b
}
