package client

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"encoding/base64"
)

type FcInvokeInput struct {
	FuncName    string
	Args		[]byte
	EnableNative string
}

type FcInvokeOutput struct {
	StatusCode int
	RespBody []byte
}

type invokeBody struct {
	FuncName string `json:"funcName"`
	Args     string `json:"args"`

	// 'true' : may use native serverless, optimize cold-boot
	// 'false' : prevent manager to use native serverless. avoid the limits of platform.
	EnableNative string `json:"enableNative"`
}

func (c *Client) FcInvoke(fcInvokeInput *FcInvokeInput) (*FcInvokeOutput, error) {
	if fcInvokeInput == nil {
		return nil, errors.New("null pointer")
	}

	body := invokeBody{
		FuncName: fcInvokeInput.FuncName,
		Args: base64.StdEncoding.EncodeToString(fcInvokeInput.Args),
		EnableNative: fcInvokeInput.EnableNative,
	}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)
	
	url := "http://" + c.host + "/invoke"
    req, err := http.NewRequest("POST", url, buf)
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
		StatusCode: resp.StatusCode,
		RespBody: ret,
	}, nil
}
