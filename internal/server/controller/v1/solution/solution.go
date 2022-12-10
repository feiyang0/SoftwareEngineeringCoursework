package solution

import "SoftwareEngine/internal/server/store"

type SolutionController struct {
	solutionS store.SolutionStore
}

func NewSolutionController(store store.Factory) *SolutionController {
	return &SolutionController{
		solutionS: store.Solutions(),
	}
}
