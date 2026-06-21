package agentbrowser

import "fmt"

// StartDashboard starts the agent-browser observability dashboard
func (c *Client) StartDashboard() error {
	_, err := c.runCmd("dashboard", "start", "--json")
	if err != nil {
		return fmt.Errorf("failed to start dashboard: %w", err)
	}

	return nil
}

// StopDashboard stops the running agent-browser dashboard server.
func (c *Client) StopDashboard() error {
	_, err := c.runCmd("dashboard", "stop", "--json")
	if err != nil {
		return fmt.Errorf("failed to stop dashboard: %w", err)
	}

	return nil
}
