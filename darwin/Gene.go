package darwin

type Gene interface {
	get() int
	set(val int)
	mutate()
	getCost() int
}
