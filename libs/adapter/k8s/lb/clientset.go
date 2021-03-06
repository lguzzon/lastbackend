package lb

import (
	"github.com/lastbackend/lastbackend/libs/adapter/k8s/lb/v1"
	"k8s.io/client-go/pkg/util/flowcontrol"
	"k8s.io/client-go/rest"
)

// Interface exposes methods on k8s resources.
type Interface interface {
	LB() v1.LBInterface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type LBClientset struct {
	*v1.LBClient
}

func (c *LBClientset) LB() v1.LBInterface {
	if c == nil {
		return nil
	}
	return c.LBClient
}

// NewForConfig creates a new Clientset for the given config.
func NewForConfig(c *rest.Config) (*LBClientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var clientset LBClientset
	var err error
	clientset.LBClient, err = v1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	return &clientset, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *LBClientset {
	var clientset LBClientset
	clientset.LBClient = v1.NewForConfigOrDie(c)
	return &clientset
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *LBClientset {
	var clientset LBClientset
	clientset.LBClient = v1.New(c)
	return &clientset
}
