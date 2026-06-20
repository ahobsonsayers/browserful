package api

import (
	"github.com/ahobsonsayers/browserfull/internal/agentbrowser"
	"github.com/ahobsonsayers/browserfull/internal/config"
)

// Server is the type that all handlers are defined on as methods
// It satisfies the strict server interfaces
type Server struct{}

var _ StrictServerInterface = Server{}

// ServerOverrides is a type that handler overrides (with access
// to request and response types) can be defined on.
// It satisfies the service interface by embedding another server interface
type ServerOverrides struct {
	ServerInterface
	agentBrowser   *agentbrowser.Browser
	allowedOrigins []string
}

var _ ServerInterface = ServerOverrides{}

func NewServer(agentBrowser *agentbrowser.Browser, cfg *config.Config) ServerInterface {
	// Create strict service
	strictServer := Server{}

	// Convert to none strict server
	server := NewStrictHandler(strictServer, nil)

	// Apply server overrides
	overriddenServer := ServerOverrides{
		ServerInterface: server,
		agentBrowser:    agentBrowser,
		allowedOrigins:  cfg.AllowedOrigins,
	}

	return overriddenServer
}
