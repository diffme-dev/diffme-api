package core

type ExecuteFunction func(request interface{}) interface{}

type UseCase struct {
	Execute ExecuteFunction
}
