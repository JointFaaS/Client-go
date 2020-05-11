package excutor

import (
	"github.com/JointFaaS/Client-go/client"
)

// 
func (e *Excutor) Refresh() (error) {
	for _, c := range e.clients {
		m, err := c.GetResourceMetrics()
		if err != nil {
			return err
		}
		
	}
}