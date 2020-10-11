package s

import (
	"context"
	"time"

	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"

	"github.com/aclisp/godashboard/frontend/s/action"
	"github.com/aclisp/godashboard/frontend/s/dispatcher"
	dashboard "github.com/aclisp/godashboard/proto"
)

type dynamicViewUpdater struct {
	started bool
	ctx     context.Context
	cancel  context.CancelFunc
	update  chan struct{}
}

func (u *dynamicViewUpdater) queryWithTimeout(c dashboard.BackendClient, timeout time.Duration) (
	req *dashboard.QueryReq, res *dashboard.QueryRes, sta *status.Status) {
	target := State.CurrentPackageEndpoint()
	env := State.CurrentGatewayID()[0]
	gate := State.CurrentGatewayID()[1]

	req = &dashboard.QueryReq{
		Package:     target.PackageName,
		Endpoint:    target.EndpointName,
		Environment: env,
		GatewayCode: gate,
		YsadminReq: &dashboard.YsadminReq{
			Cmd: target.EndpointName,
		}}

	ctx, cancel := context.WithTimeout(u.ctx, timeout)
	defer cancel()

	res, err := c.Query(ctx, req)
	if err != nil {
		sta = status.Convert(err)
		grpclog.Errorf("dynamicViewUpdater.queryWithTimeout: %v %q details: %v", sta.Code(), sta.Message(), sta.Details())
		return req, nil, sta
	}
	return req, res, nil
}

func (u *dynamicViewUpdater) start() {
	if u.started {
		return
	}

	u.started = true
	u.ctx, u.cancel = context.WithCancel(context.Background())
	u.update = make(chan struct{}, 2)

	go func() {
		queryTimeout := 10 * time.Second
		tickerPeriod := 10 * time.Minute
		ticker := time.NewTicker(tickerPeriod)
		defer ticker.Stop()

	Loop:
		for {
			var (
				req *dashboard.QueryReq
				res *dashboard.QueryRes
				sta *status.Status
			)

			select {
			case <-ticker.C:
				req, res, sta = u.queryWithTimeout(client, queryTimeout)
			case <-u.ctx.Done():
				break Loop
			case <-u.update:
				// reset ticker
				ticker.Stop()
				ticker = time.NewTicker(tickerPeriod)
				// query now
				req, res, sta = u.queryWithTimeout(client, queryTimeout)
			}

			dispatcher.Dispatch(&action.SyncDynamicViewDataDone{Req: req, Res: res, Sta: sta})

			if u.ctx.Err() != nil {
				break
			}
		}

		grpclog.Infof("dynamicViewUpdater stopped")
	}()
}

func (u *dynamicViewUpdater) stop() {
	if !u.started {
		return
	}

	u.cancel()
	u.started = false
}

func (u *dynamicViewUpdater) sync() {
	u.start()
	select {
	case u.update <- struct{}{}:
	default:
		grpclog.Warningf("dynamicViewUpdater.sync skipped because there are %d outstanding RPC", len(u.update))
	}
}
