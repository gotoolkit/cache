package cache

import (
	"net/http"
	"sync"

	"github.com/gotoolkit/consistenthash"
)

const defaultBasePath = "/_groupcache/"
const defaultReplicas = 50

// HTTPPool HTTP peers
type HTTPPool struct {
	Context func(*http.Request) Context
	opts    HTTPPoolOptions
	self    string

	mu sync.Mutex // guards peers and httpGetters

	httpGetters map[string]*httpGetter
	peers       *consistenthash.Map
}

// HTTPPoolOptions HTTP peers配置
type HTTPPoolOptions struct {
	BasePath string
	Replicas int
	HashFn   consistenthash.Hash
}

// NewHTTPool 初始化http
func NewHTTPool(self string) *HTTPPool {
	p := NewHTTPoolOpts(self, nil)
	http.Handle(p.opts.BasePath, p)
	return p
}

var httpPoolMade bool

// NewHTTPool 初始化配置
func NewHTTPoolOpts(self string, o *HTTPPoolOptions) *HTTPPool {
	if httpPoolMade {
		panic("cache: NewHTTPPool must be called only once")
	}
	httpPoolMade = true

	p := &HTTPPool{
		self:        self,
		httpGetters: make(map[string]*httpGetter),
	}
	if o != nil {
		p.opts = *o
	}

	if p.opts.BasePath == "" {
		p.opts.BasePath = defaultBasePath
	}

	if p.opts.Replicas == 0 {
		p.opts.Replicas = defaultReplicas
	}
	p.peers = consistenthash.New(p.opts.Replicas, p.opts.HashFn)

	RegisterPeerPicker(func() PeerPicker { return p })

	return p
}

func (p *HTTPPool) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}

type httpGetter struct {
	transport func(Context) http.RoundTripper
	baseURL   string
}

func (p *HTTPPool) PickPeer(key string) (ProtoGetter, bool) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.peers.IsEmpty() {
		return nil, false
	}
	if peer := p.peers.Get(key); peer != p.self {
		return p.httpGetters[peer], true
	}
	return nil, false
}

func (p *HTTPPool) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}
