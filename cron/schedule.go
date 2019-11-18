// Package cron implements interfaces schedule.
package cron

import (
	"github.com/golang-mixins/schedule"
	Cron "github.com/robfig/cron/v3" // "gopkg.in/robfig/cron.v3"
	"golang.org/x/xerrors"
	"io"
	"log"
)

// Scheduler defines the structure implements interfaces schedule.
// Using structure methods without initialization with the New constructor will lead to panic.
type Scheduler struct {
	*Cron.Cron
}

// AddTask - adds an executable task to the scheduler.
func (s *Scheduler) AddTask(spec string, cmd func()) error {
	_, err := s.AddFunc(spec, cmd)
	if err != nil {
		return xerrors.Errorf("add spec '%s' error: %w", spec, err)
	}

	return nil
}

// Reset - resets the scheduler, making it again available without any tasks on it.
func (s *Scheduler) Reset() {
	go s.Stop()
	*s = Scheduler{Cron.New()}
}

// New - constructor Scheduler.
func New(logOut io.Writer) schedule.Scheduler {
	return &Scheduler{
		Cron.New(
			Cron.WithParser(
				Cron.NewParser(
					Cron.Descriptor|Cron.SecondOptional|Cron.Minute|Cron.Hour|Cron.Dom|Cron.Month|Cron.Dow,
				),
			),
			Cron.WithLogger(
				Cron.VerbosePrintfLogger(
					log.New(logOut,
						"go cron scheduler: ",
						log.LstdFlags|log.Lmicroseconds|log.Llongfile|log.Lshortfile,
					),
				),
			),
		),
	}
}
