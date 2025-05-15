package worker

import (
	"context"
	"currency_service/crud/service"
	"log"
	"time"
)

type Cron struct {
	ctx    context.Context
	cancel context.CancelFunc
	ticker *time.Ticker
	svc    *service.CronCurrencyServer
}

func NewCron(svc *service.CronCurrencyServer) *Cron {
	ctx, cancel := context.WithCancel(context.Background())
	return &Cron{
		ctx:    ctx,
		cancel: cancel,
		ticker: time.NewTicker(1 * time.Hour),
		svc:    svc,
	}
}

func (c *Cron) Start() {
	log.Println("Starting cron worker")
	go func() {
		for {
			select {
			case <-c.ctx.Done():
				log.Println("Stop worker")
				return
			case <-c.ticker.C:
				c.svc.CronUpdateCurrencyRates()
			}
		}
	}()
}

func (c *Cron) Stop() {
	c.ticker.Stop()
	c.cancel()
}
