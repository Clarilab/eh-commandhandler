package ehcommandhandler

import (
	"context"
	"sync"

	eh "github.com/looplab/eventhorizon"
)

// AtomicCommandHandler is a command handler that can be used to securely execute commands concurrently.
type AtomicCommandHandler struct {
	handler eh.CommandHandler
	rwMutex *sync.RWMutex
	a       map[string]*sync.Mutex
}

// NewCommandHandler creates a new atomic CommandHandler for an aggregate type.
func NewAtomicCommandHandler(h eh.CommandHandler) *AtomicCommandHandler {
	return &AtomicCommandHandler{
		handler: h,
		rwMutex: &sync.RWMutex{},
		a:       make(map[string]*sync.Mutex),
	}
}

// HandleCommand handles a command with the registered aggregate.
// Returns ErrAggregateNotFound if no aggregate could be found.
func (h *AtomicCommandHandler) HandleCommand(ctx context.Context, cmd eh.Command) error {
	h.rwMutex.RLock()
	_, ok := h.a[cmd.AggregateType().String()]
	h.rwMutex.RUnlock()

	if !ok {
		h.rwMutex.Lock()
		h.a[cmd.AggregateID().String()] = &sync.Mutex{}
		h.rwMutex.Unlock()
	}

	h.rwMutex.RLock()
	defer h.rwMutex.RUnlock()

	h.a[cmd.AggregateID().String()].Lock()
	defer h.a[cmd.AggregateID().String()].Unlock()

	return h.handler.HandleCommand(ctx, cmd)
}
