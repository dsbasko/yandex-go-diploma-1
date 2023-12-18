package rmq

import "fmt"

func (c *Connector) Close() error {
	if err := c.conn.Close(); err != nil {
		return fmt.Errorf("conn.Close: %w", err)
	}

	return nil
}
