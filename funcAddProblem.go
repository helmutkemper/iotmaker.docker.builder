package iotmakerdockerbuilder

func (e *ContainerBuilder) addProblem(problem string) {
	if e.problem == "" {
		e.problem = problem
	} else {
		e.problem += "\n" + problem
	}
}
