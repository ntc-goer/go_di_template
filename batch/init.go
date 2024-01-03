package batch

import (
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"go_di_template/config"
)

type Batch struct {
	Config *config.Config
}

func NewBatch(cfg *config.Config) *Batch {
	return &Batch{
		Config: cfg,
	}
}

func (b *Batch) Init() {
	c := cron.New()
	// Register batch job
	_, err := c.AddFunc("@midnight", b.PullImage)
	if err != nil {
		logrus.Errorf("Add batch PullImage func fail %s", err)
	}
	c.Start()
}
