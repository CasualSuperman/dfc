package terrain

import "sync"

func (c *chunk) Tick(s *sync.WaitGroup) {
	defer s.Done()
	c.tickCount++

	// Ten ticks in a cycle.
	if c.tickCount == 10 {
		c.tickCount = 0
	}

	// Update temperatures not as often as every tick.
	if c.tickCount == 0 {
		s.Add(1)
		go c.updateTemperatures(s)
	}
}

func (c *chunk) updateTemperatures(s *sync.WaitGroup) {
	defer s.Done()
}
