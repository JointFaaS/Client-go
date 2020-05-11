package flow

import (
	"errors"

	"github.com/JointFaaS/Client-go/client"
)

type input struct {
	lastIndex int32
	argsName string
}

type fcInvocation struct {
	inputFrom []*input
	funcName string
	args []byte
	ret []byte
}

type layer struct {
	fcs []*fcInvocation
}

type FcFlow struct {
	layers []layer

}

func NewFlow() (*FcFlow){
	fcf := &FcFlow{

	}
	return fcf
}

func (fcf *FcFlow) AddSource(funcName string, args []byte) (error){
	return nil
}

func (fcf *FcFlow) AddNode(funcName string, args []byte, inputs []*input) (error){
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

func (fcf *FcFlow) Exec(client *client.Client) (*ExecFlowOutput, error) {
	if client == nil {
		return nil, errors.New("null pointer")
	}

	return nil, nil
}