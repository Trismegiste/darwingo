package darwin

type Gene interface {
	get() int
	mutate()
	getCost() int
}
