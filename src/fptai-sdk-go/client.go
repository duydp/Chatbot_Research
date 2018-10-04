package fptai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/pkg/errors"
)

type Client struct {
	token string
}

func NewClient(token string) *Client {
	return &Client{
		token: token,
	}
}

func (c *Client) GetIntents() (intents []Intent, err error) {
	resp, err := c.get("/intents")
	if err != nil {
		return intents, errors.Wrapf(err, "get failed")
	}

	result := struct{
		Intents []Intent
	}{}
	if err := json.Unmarshal(resp, &result); err != nil {
		return intents, err
	}

	return result.Intents, nil
}

func (c *Client) CreateIntent(name, description string) (i Intent, err error) {
	r := struct{
		Name string `json:"label"`
		Description string
	}{
		Name: name,
		Description: description,
	}

	data, err := json.Marshal(r)
	if err != nil {
		return i, errors.Wrapf(err, "Marshal request failed: input=%+v\n", r)
	}

	resp, err := c.post("/intents", data)
	if err != nil {
		return i, errors.Wrapf(err, "post failed")
	}

	if err := json.Unmarshal(resp, &i); err != nil {
		return i, err
	}

	return i, nil
}

func (c *Client) DeleteIntent(name string) error {
	_, err := c.delete("/intents/" + name)
	return err
}

func (c *Client) CreateUtterances(intent string, utterances []string) error {
	r := struct{
		Utterances []string
	}{
		Utterances: utterances,
	}
	data, err := json.Marshal(r)
	if err != nil {
		return err
	}
	_, err = c.post(fmt.Sprintf("/intents/%s/utterances", intent), data)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) RecognizeIntents(text string) (m Meaning, err error) {
	r := struct{
		Text string
	}{
		Text: text,
	}
	data, err := json.Marshal(r)
	if err != nil {
		return m, err
	}

	resp, err := c.post("/recognize/intent", data)
	if err != nil {
		return m, err
	}

	err = json.Unmarshal(resp, &m)
	if err != nil {
		return m, errors.Wrapf(err, "Unmarshal failed, data=%s", string(resp))
	}

	return m, nil
}

func (c *Client) TrainIntent() error {
	_, err := c.post("/train/intent", nil)
	return err
}


func (c *Client) GetEntities(text string) (m Meaning, err error) {
	r := struct{
		Text string
	}{
		Text: text,
	}
	data, err := json.Marshal(r)
	if err != nil {
		return m, err
	}

	resp, err := c.post("/recognize/entity", data)
	if err != nil {
		return m, err
	}

	err = json.Unmarshal(resp, &m)
	if err != nil {
		return m, errors.Wrapf(err, "Unmarshal failed, data=%s", string(resp))
	}

	return m, nil
}

func (c *Client) get(path string) ([]byte, error) {
	return c.request("GET", path, nil)
}

func (c *Client) post(path string, data []byte) ([]byte, error) {
	return c.request("POST", path, data)
}

func (c *Client) put(path string, data []byte) ([]byte, error) {
	return c.request("PUT", path, data)
}

func (c *Client) delete(path string) ([]byte, error) {
	return c.request("DELETE", path, nil)
}

func (c *Client) request(method string, path string, data []byte) ([]byte, error) {
	path = strings.TrimPrefix(path, "/") // path could be both "/path", "path"
	URI := fmt.Sprintf("%s/%s/%s", FPTAIEndpoint, VERSION, path)

	req, err := http.NewRequest(method, URI, bytes.NewReader(data))
	if err != nil {
		return nil, errors.Wrapf(err, "NewRequest failed: %s %s\n", method, URI)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))

	hc := &http.Client{
		Timeout: time.Duration(REQUEST_TIMEOUT) * time.Second,
	}

	resp, err := hc.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, "Do failed")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "ReadAll failed: body=%+v\n", resp.Body)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		log.Errorf("Status %d, Body %s", resp.StatusCode, body)
		var err Error
		if e := json.Unmarshal(body, &err); e != nil {
			return nil, errors.Wrapf(e, "Decode failed: body=%s\n", string(body))
		}
		return nil, err
	}
	return body, nil
}