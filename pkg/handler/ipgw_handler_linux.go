//go:build linux

package handler

import (
	"net/http"

	"github.com/neucn/ipgw/pkg/model"
	"github.com/neucn/neugo"
)

type IpgwHandler struct {
	info    *model.Info
	client  *http.Client
	oriInfo map[string]interface{}
}

func (h *IpgwHandler) GetInfo() *model.Info {
	return h.info
}

func (h *IpgwHandler) GetClient() *http.Client {
	return h.client
}

func NewIpgwHandler(mark uint32) *IpgwHandler {
	return &IpgwHandler{
		info:    &model.Info{},
		client:  neugo.NewFwmarkSession(mark),
		oriInfo: make(map[string]interface{}),
	}
}
