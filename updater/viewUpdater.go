package updater

import "time"

type ViewUpdater struct {
	TickRate time.Duration
	Funcs    []func()
	Caches   []Cache
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

func (vu *ViewUpdater) AddCache(cache Cache) {
	vu.Caches = append(vu.Caches, cache)
}

func updateViews(vu *ViewUpdater) {
	for _, f := range vu.Funcs {
		if f != nil {
			f()
		}
	}
}

func updateCaches(vu *ViewUpdater) {
	for _, f := range vu.Caches {
		f.Update()
	}
}

func update(vu *ViewUpdater) {
	updateCaches(vu)
	updateViews(vu)
}

func (vu *ViewUpdater) Run() {
	go func() {
		update(vu)
		ticker := time.NewTicker(vu.TickRate)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				update(vu)
			}
		}
	}()
}
