package client

import (
	"io"
	"bytes"
	"mime/multipart"
	"net/http"
	"errors"
)

type FcCreateInput struct {
	FuncName string
	Code io.Reader
	Env string
	CloudFilter *Filter
}

type FcCreateOutput struct {

}

func (c *Client) FcCreate(fcCreateInput *FcCreateInput) (*FcCreateOutput, error) {
	if fcCreateInput == nil {
		return nil, errors.New("null pointer")
	}

	body := new(bytes.Buffer)
    writer := multipart.NewWriter(body)
    writer.WriteField("funcName", fcCreateInput.FuncName)
	writer.WriteField("env", fcCreateInput.Env)

    formFile, err := writer.CreateFormFile("sourceZip", "code.zip")
    if err != nil {
        return nil, err
    }
	_, err = io.Copy(formFile, fcCreateInput.Code)
	if err != nil {
		return nil, err
	}

	err = writer.Close()
	if err != nil {
        return nil, err
	}
	
	url := "http://" + c.host + "/createfunction"
    req, err := http.NewRequest("POST", url, body)
    if err != nil {
        return nil, err
	}
	
    req.Header.Add("Content-Type", writer.FormDataContentType())
	
    resp, err := c.hc.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
    return nil, nil
}

type FcDeleteInput struct {
	FuncName string
	CloudFilter *Filter
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
	CloudFilter *Filter
}

func (c *Client) FcList() (error) {
	return nil
}

type FcGetInput struct {
	FuncName string
	CloudFilter *Filter
}

type FcGetOutput struct {

}

func (c *Client) FcGet(fcGetInput *FcGetInput) (*FcGetOutput, error) {
	return nil, nil
}