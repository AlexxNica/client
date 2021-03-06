// Copyright 2015 Keybase, Inc. All rights reserved. Use of
// this source code is governed by the included BSD license.

package service

import (
	"golang.org/x/net/context"

	"github.com/keybase/client/go/libkb"
	keybase1 "github.com/keybase/client/go/protocol/keybase1"
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
)

type KBFSHandler struct {
	*BaseHandler
	libkb.Contextified
}

func NewKBFSHandler(xp rpc.Transporter, g *libkb.GlobalContext) *KBFSHandler {
	return &KBFSHandler{
		BaseHandler:  NewBaseHandler(xp),
		Contextified: libkb.NewContextified(g),
	}
}

func (h *KBFSHandler) FSEvent(_ context.Context, arg keybase1.FSNotification) error {
	h.G().NotifyRouter.HandleFSActivity(arg)
	return nil
}

func (h *KBFSHandler) FSEditList(ctx context.Context, arg keybase1.FSEditListArg) error {
	h.G().NotifyRouter.HandleFSEditListResponse(ctx, arg)
	return nil
}

func (h *KBFSHandler) FSEditListRequest(ctx context.Context, arg keybase1.FSEditListRequest) error {
	h.G().NotifyRouter.HandleFSEditListRequest(ctx, arg)
	return nil
}

func (h *KBFSHandler) FSSyncStatus(ctx context.Context, arg keybase1.FSSyncStatusArg) (err error) {
	h.G().NotifyRouter.HandleFSSyncStatus(ctx, arg)
	return nil
}

func (h *KBFSHandler) FSSyncEvent(ctx context.Context, arg keybase1.FSPathSyncStatus) (err error) {
	h.G().NotifyRouter.HandleFSSyncEvent(ctx, arg)
	return nil
}
