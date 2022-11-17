package problem

import "SoftwareEngine/internal/server/store"

type ProblemController struct {
	problemS store.ProblemStore
}

func NewProblemController(store store.Factory) *ProblemController {
	return &ProblemController{
		problemS: store.Problems(),
	}
}
