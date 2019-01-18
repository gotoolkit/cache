package cache

import "net/http"

// Context
type Context interface{}

type ProtoGetter interface {
	Get(context Context, in *http.Request, out *http.Response) error
}

// PeerPicker is the interface that must be implemented to locate
// the peer that owns a specific key.
type PeerPicker interface {
	// PickPeer returns the peer that owns the specific key
	// and true to indicate that a remote peer was nominated.
	// It returns nil, false if the key owner is the current peer.
	PickPeer(key string) (peer ProtoGetter, ok bool)
}

type NoPeers struct{}

func (NoPeers) PickPeer(key string) (peer ProtoGetter, ok bool) { return }

var (
	portPicker func(groupName string) PeerPicker
)

func RegisterPeerPicker(fn func() PeerPicker) {
	if portPicker != nil {

	}
	portPicker = func(_ string) PeerPicker { return fn() }
}
