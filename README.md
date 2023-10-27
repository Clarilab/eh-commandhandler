# eh-commandhandler
An atomic wrapper for an [eventhorizon](https://github.com/looplab/eventhorizon) aggregate command handler.

### Usage

```go
import (
	atomicHandler "github.com/Clarilab/eh-commandhandler"

	"github.com/looplab/eventhorizon/aggregatestore/events"
	"github.com/looplab/eventhorizon/commandhandler/aggregate"

	eh "github.com/looplab/eventhorizon"
)

func SetupCommandHandler(t eh.AggregateType, eventStore eh.EventStore) (*atomicHandler.AtomicCommandHandler, error) {
	aggregateStore, err := events.NewAggregateStore(eventStore)
	if err != nil {
		return nil, err
	}

	aggregateCommandHandler, err := aggregate.NewCommandHandler(t, aggregateStore)
	if err != nil {
		return nil, err
	}

	return atomicHandler.NewAtomicCommandHandler(aggregateCommandHandler), nil
}

```