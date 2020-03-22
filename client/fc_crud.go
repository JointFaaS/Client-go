package client

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

type FcCreateInput struct {
	FuncName string
	Code []byte
	Env string
	Timeout string
	MemorySize string
}

type FcCreateOutput struct {
	StatusCode int
	RespBody []byte
}

type fcCreateBody struct {
	FuncName string `json:"funcName"`
	CodeZip  string `json:"codeZip"`
	Env string `json:"env"`
	MemorySize     string `json:"memorySize"`
	Timeout string `json:"timeout"`
}

func (c *Client) FcCreate(fcCreateInput *FcCreateInput) (*FcCreateOutput, error) {
	if fcCreateInput == nil {
		return nil, errors.New("null pointer")
	}

	body := fcCreateBody{
		FuncName: fcCreateInput.FuncName,
		CodeZip: base64.StdEncoding.EncodeToString(fcCreateInput.Code),
		Env: fcCreateInput.Env,
		MemorySize: fcCreateInput.MemorySize,
		Timeout: fcCreateInput.Timeout,
	}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)
	
	url := "http://" + c.host + "/create"
    req, err := http.NewRequest("POST", url, buf)
    if err != nil {
        return nil, err
	}
	
    resp, err := c.hc.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
    return &FcCreateOutput{
		StatusCode: resp.StatusCode,
		RespBody: respBody,
	}, nil
}

type FcDeleteInput struct {
	FuncName string
}

type FcDeleteOutput struct {
	
}

func (c *Client) FcDelete(fcDeleteInput *FcDeleteInput) (*FcDeleteOutput, error) {
	if fcDeleteInput == nil {
		return nil, errors.New("null pointer")
	}

	url := "http://" + c.host + "/delete?funcName=" + fcDeleteInput.FuncName
    req, err := http.NewRequest("POST", url, nil)
    if err != nil {
        return nil, err
	}
	
    _, err = c.hc.Do(req)
    if err != nil {
        return nil, err
    }

	return nil, nil
}

type FcListInput struct {
}

func (c *Client) FcList() (error) {
	return nil
}

type FcGetInput struct {
	FuncName string
}

type FcGetOutput struct {

}

func (c *Client) FcGet(fcGetInput *FcGetInput) (*FcGetOutput, error) {
	return nil, nil
}