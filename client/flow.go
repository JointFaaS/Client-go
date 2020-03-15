package client

import (
	"errors"
)
type FcFlow struct {
	
}

func (fcf *FcFlow) NewFlow() (*FcFlow, error){
	return nil, nil
}

func (fcf *FcFlow) AddSource() (error){
	return nil
}

func (fcf *FcFlow) AddNode() (error){
	return nil
}

func (fcf *FcFlow) NextLayer() (error){
	return nil
}

func (fcf *FcFlow) End() (error){
	return nil
}

type ExecFlowOutput struct {

}

func (c *Client) ExecFlow(flow *FcFlow, filter *Filter) (*ExecFlowOutput, error) {
	if flow == nil {
		return nil, errors.New("null pointer")
	}

	return nil, nil
}