package updater

import "time"

type ViewUpdater struct {
	TickRate time.Duration
	Funcs    []func()
}

func CreateNew(tickRate time.Duration) *ViewUpdater {
	return &ViewUpdater{
		TickRate: tickRate,
		Funcs:    []func(){},
	}
}

func (vu *ViewUpdater) AddView(view func()) {
	vu.Funcs = append(vu.Funcs, view)
}

func (vu *ViewUpdater) Run() {
	go func() {
		ticker := time.NewTicker(vu.TickRate)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				for _, f := range vu.Funcs {
					if f != nil {
						f()
					}
				}
			}
		}
	}()
}
