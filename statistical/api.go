package statistical

import (
	"context"
	"myproject/node"
	"myproject/rpc"
	"time"
)

// Statistical Statistical service
type Statistical struct{}

// New create Statistical instance
func New(n *node.ServiceContext) (*Statistical, error) {
	return &Statistical{}, nil
}

// PublicStatisticalAPI public Statistical api
type PublicStatisticalAPI struct{}

//Nowtime
func (p *PublicStatisticalAPI) GetNowTime() interface{} {
	timeStr := time.Now().Unix()
	return timeStr
}

// Now get server now time with json-rpc or websocket
func (p *PublicStatisticalAPI) Now(ctx context.Context) (*rpc.Subscription, error) {
	notifier, supported := rpc.NotifierFromContext(ctx)
	if !supported {
		return &rpc.Subscription{}, rpc.ErrNotificationsUnsupported
	}
	rpcSub := notifier.CreateSubscription()

	go func() {
		ticker := time.NewTicker(time.Second * 5)

		for {
			select {
			case <-ticker.C:
				notifier.Notify(rpcSub.ID, time.Now().String())
			case <-rpcSub.Err():
				return
			case <-notifier.Closed():
				return
			}
		}
	}()

	return rpcSub, nil
}

// APIs retrieves the list of RPC descriptors the service provides
func (p *Statistical) APIs() []rpc.API {
	return []rpc.API{
		{
			Namespace: "statistical",
			Version:   "1.0",
			Service:   &PublicStatisticalAPI{},
			Public:    true,
		},
	}
}

// Start is called after all services have been constructed and the networking
// layer was also initialized to spawn any goroutines required by the service.
func (p *Statistical) Start() error {
	return nil
}

// Stop terminates all goroutines belonging to the service, blocking until they
// are all terminated.
func (p *Statistical) Stop() error {
	return nil
}
func (p *PublicStatisticalAPI) GetNow() int64 {
	return time.Now().Unix()
}
