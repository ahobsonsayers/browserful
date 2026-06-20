package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ahobsonsayers/browserful/internal/proxy"
)

// LaunchDefaultSession is overridden by the ServerOverrides LaunchDefaultSession below
func (Server) LaunchDefaultSession(
	context.Context, LaunchDefaultSessionRequestObject,
) (LaunchDefaultSessionResponseObject, error) {
	return nil, nil
}

// LaunchNamedSession is overridden by the ServerOverrides LaunchNamedSession below
func (Server) LaunchNamedSession(
	context.Context, LaunchNamedSessionRequestObject,
) (LaunchNamedSessionResponseObject, error) {
	return nil, nil
}

// CloseSession is overridden by the ServerOverrides CloseSession below
func (Server) CloseSession(
	context.Context, CloseSessionRequestObject,
) (CloseSessionResponseObject, error) {
	return nil, nil
}

func (s ServerOverrides) LaunchDefaultSession(w http.ResponseWriter, r *http.Request) {
	err := s.handleLaunchSession(w, r, "default")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s ServerOverrides) LaunchNamedSession(w http.ResponseWriter, r *http.Request, sessionName string) {
	err := s.handleLaunchSession(w, r, sessionName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s ServerOverrides) CloseSession(w http.ResponseWriter, _ *http.Request, sessionName string) {
	err := s.agentBrowser.Close(sessionName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s ServerOverrides) handleLaunchSession(
	w http.ResponseWriter, r *http.Request, sessionName string,
) error {
	info, err := s.agentBrowser.Launch(sessionName)
	if err != nil {
		return fmt.Errorf("error launching browser session cdp: %w", err)
	}

	err = proxy.ProxyCDP(w, r, info.CDPURL, s.allowedOrigins)
	if err != nil {
		return fmt.Errorf("error proxying cdp: %w", err)
	}

	return nil
}
