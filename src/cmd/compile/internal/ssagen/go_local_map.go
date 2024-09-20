package ssagen

import (
	"cmd/compile/internal/ir"
	"cmd/compile/internal/ssa"
	"sync"
)

var goLocalMapMutex sync.Mutex

// goLocalAllocMap is mapping go_local variable to its need init ssa value.
var goLocalAllocMap = map[*ir.Name]*ssa.Value{}

func setGoLocalAlloc(n *ir.Name, v *ssa.Value) {
	goLocalMapMutex.Lock()
	defer goLocalMapMutex.Unlock()
	goLocalAllocMap[n] = v
}

func getGoLocalAlloc(n *ir.Name) *ssa.Value {
	goLocalMapMutex.Lock()
	defer goLocalMapMutex.Unlock()
	return goLocalAllocMap[n]
}
