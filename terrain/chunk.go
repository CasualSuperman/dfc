package terrain

import "sync"

func (c *chunk) Tick(s *sync.WaitGroup) {
	s.Add(1)
	c.updateTemperatures(s)

	s.Done()
}

func (c *chunk) updateTemperatures(s *sync.WaitGroup) {
	s.Done()
}
