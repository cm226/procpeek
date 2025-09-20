package updater

type Cache interface {
	Update()
}

type ToolCache[T any] struct {
	Data    T
	updater func() T
}

func MakeToolCache[T any](toolFn func() T) ToolCache[T] {
	return ToolCache[T]{
		Data:    toolFn(),
		updater: toolFn,
	}
}

func (c *ToolCache[T]) Update() {
	c.Data = c.updater()
}
