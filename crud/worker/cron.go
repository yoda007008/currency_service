package worker

import (
	"context"
	"currency_service/crud/clients"
	"fmt"
	"log"
	"time"
)

type Cron struct {
	ctx    context.Context
	cancel context.CancelFunc
	ticker *time.Ticker
	svc    *clients.CronCurrencyServer
}

func NewCron(svc *clients.CronCurrencyServer) *Cron {
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
				err := c.svc.CronUpdateCurrencyRates()
				if err != nil {
					fmt.Errorf("Failed to start worker", err)
					return
				}
			}
		}
	}()
}

func (c *Cron) Stop() {
	c.ticker.Stop()
	c.cancel()
}
