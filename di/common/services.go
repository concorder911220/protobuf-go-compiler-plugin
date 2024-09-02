package common

type IServiceA interface {
	DoSomething()
	Register(serviceB IServiceB)
}

type IServiceB interface {
	DoSomethingElse()
	Register(serviceA IServiceA)
}
