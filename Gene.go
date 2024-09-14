package main

type Gene interface {
	get() int
	mutate()
	getCost() int
}
