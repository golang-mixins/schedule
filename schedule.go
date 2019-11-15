// Package schedule provides interfaces (and their implementation) for managing the schedule of running processes.
package schedule

import "context"

// Scheduler - interface for managing the schedule of launched processes.
type Scheduler interface {
	// AddTask - adds an executable task to the scheduler.
	// spec - defines classic crontab value "second(optional)|minute|hour|dom|month|dow".
	AddTask(spec string, cmd func()) error
	// Run - starts the scheduler if it is not already running.
	// Does not allocate coroutine, the stack stops until the scheduler stops
	Run()
	// Start - starts the scheduler in its own coroutine or ignores it if it is already running.
	Start()
	// Stop - the scheduler stops if it is running; otherwise it does nothing.
	// A context is returned so the caller can wait for running jobs to complete.
	Stop() context.Context
	// Reset - resets the scheduler, making it again available without any tasks on it.
	Reset()
}
