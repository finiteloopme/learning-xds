package ads

import (
	"context"
	"sync"

	// log "github.com/finiteloopme/goutils/pkg/log"
	// core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	discoverySvc "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	log "github.com/sirupsen/logrus"
)

type Callbacks struct {
	Signal   chan struct{}
	Debug    bool
	Fetches  int
	Requests int
	// Mutex lock
	mu sync.Mutex
}

func (cb *Callbacks) Report() {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	log.WithFields(log.Fields{"fetches": cb.Fetches, "requests": cb.Requests}).Info("cb.Report()  callbacks")
}
func (cb *Callbacks) OnStreamOpen(_ context.Context, id int64, typ string) error {
	log.Infof("OnStreamOpen %d open for %s", id, typ)
	return nil
}
func (cb *Callbacks) OnStreamClosed(id int64) {
	log.Infof("OnStreamClosed %d closed", id)
}
func (cb *Callbacks) OnStreamRequest(id int64, r *discoverySvc.DiscoveryRequest) error {
	log.Infof("OnStreamRequest %v", r.TypeUrl)
	cb.mu.Lock()
	defer cb.mu.Unlock()
	cb.Requests++
	if cb.Signal != nil {
		close(cb.Signal)
		cb.Signal = nil
	}
	return nil
}
func (cb *Callbacks) OnStreamResponse(context.Context, int64, *discoverySvc.DiscoveryRequest, *discoverySvc.DiscoveryResponse) {
	log.Infof("OnStreamResponse...")
	cb.Report()
}
func (cb *Callbacks) OnFetchRequest(ctx context.Context, req *discoverySvc.DiscoveryRequest) error {
	log.Infof("OnFetchRequest...")
	cb.mu.Lock()
	defer cb.mu.Unlock()
	cb.Fetches++
	if cb.Signal != nil {
		close(cb.Signal)
		cb.Signal = nil
	}
	return nil
}
func (cb *Callbacks) OnFetchResponse(*discoverySvc.DiscoveryRequest, *discoverySvc.DiscoveryResponse) {
	log.Infof("OnFetchResponse...")
}

func (cb *Callbacks) OnDeltaStreamClosed(int64) {
	log.Infof("OnDeltaStreamClosed...")
}

func (cb *Callbacks) OnDeltaStreamOpen(context.Context, int64, string) error {
	log.Infof("OnDeltaStreamOpen...")
	return nil
}

func (cb *Callbacks) OnStreamDeltaRequest(int64, *discoverySvc.DeltaDiscoveryRequest) error {
	log.Infof("OnStreamDeltaRequest...")
	return nil
}

// func (cb *Callbacks) OnStreamDeltaResponse(context.Context, int64, *discoverySvc.DeltaDiscoveryRequest, *discoverySvc.DeltaDiscoveryResponse) {
func (cb *Callbacks) OnStreamDeltaResponse(int64, *discoverySvc.DeltaDiscoveryRequest, *discoverySvc.DeltaDiscoveryResponse) {
	log.Infof("OnStreamDeltaResponse...")

}
