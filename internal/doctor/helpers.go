package doctor

func OK(name, message string) Result {
	return Result{
		Name:    name,
		Status:  StatusOK,
		Message: message,
	}
}

func Warning(name, message string) Result {
	return Result{
		Name:    name,
		Status:  StatusWarning,
		Message: message,
	}
}

func Error(name, message string) Result {
	return Result{
		Name:    name,
		Status:  StatusError,
		Message: message,
	}
}
