package run

// Group collects actors and runs them concurrently.
// When one actor returns, all actors are interrupted.
// The zero value of a Group is useful.
type Group struct {
	actors []actor
}

// Add an actor to the group.
//
// The first actor to return interrupts all running actors.
func (g *Group) Add(execute func() error, interrupt func(error)) {
	g.actors = append(g.actors, actor{execute, interrupt})
}

// Run all actors concurrently.
// Run only returns when all actors have exited.
// Run returns only the error returned by the first exiting actor.
func (g *Group) Run() error {
	if len(g.actors) == 0 {
		return nil
	}

	// Run each actor.
	errors := make(chan error, len(g.actors))
	for _, a := range g.actors {
		go func(a actor) {
			errors <- a.execute()
		}(a)
	}

	// Wait for the first actor to stop.
	err := <-errors

	// Signal all actors to stop.
	for _, actor := range g.actors {
		actor.interrupt(err)
	}

	// Wait for all actors to stop.
	for i := 1; i < cap(errors); i++ {
		<-errors
	}

	// Return the first actors error
	return err
}

type actor struct {
	execute   func() error
	interrupt func(error)
}
