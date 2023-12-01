package cat

type CatChan chan message

// New returns new cat entity
func New() *cat {
	return new(cat)
}

// Connect connect to server using conn
func (c *cat) Connect(conn string) error {
	return nil
}

// Subscript returns CatChan and start broadcast
func (c *cat) Subscipt() (CatChan, error) {
	c.dataCh = make(CatChan)
	c.stopCh = make(chan struct{})

	c.wg.Add(1)
	go c.broadcast()

	return c.dataCh, nil
}

// Close close CAT connection
func (c *cat) Close() error {
	c.dataCh = nil
	return nil
}

// Bytes returns a bytes body message
func (m *message) Bytes() []byte {
	return m.b
}