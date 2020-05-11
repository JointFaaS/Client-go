package excutor

import (
	"errors"
	"github.com/JointFaaS/Client-go/client"
	"github.com/JointFaaS/Client-go/strategy"
)

type Excutor struct {
	clients []*client.Client

	resources []*strategy.ResourceMetrics

	target *client.Client

	s strategy.ResourceStrategy
}

// NewExcutor
func NewExcutor(configs []*client.Config, s strategy.ResourceStrategy) (*Excutor, error) {
	e := Excutor{
		clients: make([]*client.Client, len(configs)),
		target: nil,
		s: s,
	}
	var err error
	count := len(configs)
	if count == 0 {
		return nil, errors.New("configs is empty")
	}
	for i := 0; i < count; i++ {
		config := configs[i]
		e.clients[i], err = client.NewClient(config)
		if err != nil {
			return nil, err
		}
	}
	e.Refresh()
	return &e, nil
}