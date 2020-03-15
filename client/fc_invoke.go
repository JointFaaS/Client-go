package client

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
)

type FcInvokeInput struct {
	FuncName    string
	Args		[]byte
}

type FcInvokeOutput struct {
	Ret []byte
}

func (c *Client) FcInvoke(fcInvokeInput *FcInvokeInput) (*FcInvokeOutput, error) {
	if fcInvokeInput == nil {
		return nil, errors.New("null pointer")
	}

	url := "http://" + c.host + "/invoke?funcName=" + fcInvokeInput.FuncName
    req, err := http.NewRequest("POST", url, bytes.NewReader(fcInvokeInput.Args))
    if err != nil {
        return nil, err
	}
	resp, err := c.hc.Do(req)
	if err != nil {
		return nil, err
	}
	ret, err := ioutil.ReadAll(resp.Body);
	if err != nil {
		return nil, err
	}
	return &FcInvokeOutput{
		Ret: ret,
	}, nil
}
