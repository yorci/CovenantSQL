/*
 * Copyright 2018 The ThunderDB Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package route

import (
	"fmt"
	"net"
	"testing"

	"net/rpc"

	"os"

	"time"

	log "github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/ugorji/go/codec"
	"gitlab.com/thunderdb/ThunderDB/consistent"
	"gitlab.com/thunderdb/ThunderDB/crypto/kms"
	. "gitlab.com/thunderdb/ThunderDB/proto"
)

const DHTStorePath = "./DHTStore"

func TestDHTService_FindNeighbor_FindNode(t *testing.T) {
	os.Remove(DHTStorePath)
	defer os.Remove(DHTStorePath + "1")
	log.SetLevel(log.DebugLevel)
	addr := "127.0.0.1:0"
	dht, _ := NewDHTService(DHTStorePath+"1", new(consistent.KMSStorage), false)
	kms.ResetBucket()
	rpc.RegisterName("DHT", dht)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	go func() {
		for {
			c, err := ln.Accept()
			if err != nil {
				continue
			}
			go rpc.ServeConn(c)
		}
	}()

	client, err := rpc.Dial("tcp", ln.Addr().String())
	if err != nil {
		log.Error(err)
	}

	req := &FindNeighborReq{
		NodeID: "123",
		Count:  2,
	}
	resp := new(FindNeighborResp)
	err = client.Call("DHT.FindNeighbor", req, resp)
	if err != nil {
		log.Error(err)
	}
	log.Debugf("resp: %v", resp)

	Convey("test FindNeighbor empty", t, func() {
		So(resp.Nodes, ShouldBeEmpty)
		So(err.Error(), ShouldEqual, consistent.ErrEmptyCircle.Error())
	})

	reqFN1 := &FindNodeReq{
		NodeID: "123",
	}
	respFN1 := new(FindNodeResp)
	err = client.Call("DHT.FindNode", reqFN1, respFN1)
	if err != nil {
		log.Error(err)
	}
	log.Debugf("respFN1: %v", respFN1)
	Convey("test FindNode", t, func() {
		So(respFN1.Node, ShouldBeNil)
		So(err.Error(), ShouldEqual, consistent.ErrKeyNotFound.Error())
	})

	node1 := NewNode()
	node1.InitNodeCryptoInfo(100 * time.Millisecond)
	node1.Addr = "node1 addr"

	reqA := &PingReq{
		Node: *node1,
	}
	respA := new(PingResp)
	err = client.Call("DHT.Ping", reqA, respA)
	if err != nil {
		log.Error(err)
	}
	log.Debugf("respA: %v", respA)

	reqFN2 := &FindNodeReq{
		NodeID: node1.ID,
	}
	respFN2 := new(FindNodeResp)
	err = client.Call("DHT.FindNode", reqFN2, respFN2)
	if err != nil {
		log.Error(err)
	}
	log.Debugf("respFN1: %v", respFN2)
	Convey("test FindNode", t, func() {
		So(respFN2.Node.ID, ShouldEqual, node1.ID)
		So(respFN2.Node.Nonce == node1.Nonce, ShouldBeTrue)
		So(respFN2.Node.Addr, ShouldEqual, node1.Addr)
		So(respFN2.Node.PublicKey.IsEqual(node1.PublicKey), ShouldBeTrue)
		So(err, ShouldBeNil)
	})

	node2 := NewNode()
	node2.InitNodeCryptoInfo(100 * time.Millisecond)

	reqB := &PingReq{
		Node: *node2,
	}
	respB := new(PingResp)
	err = client.Call("DHT.Ping", reqB, respB)
	if err != nil {
		log.Error(err)
	}
	log.Debugf("respA: %v", respB)

	req = &FindNeighborReq{
		NodeID: "123",
		Count:  10,
	}
	resp = new(FindNeighborResp)
	err = client.Call("DHT.FindNeighbor", req, resp)
	if err != nil {
		log.Error(err)
	}
	log.Debugf("resp: %v", resp)
	var nodeIDList []string
	for _, n := range resp.Nodes[:] {
		nodeIDList = append(nodeIDList, string(n.ID))
	}
	log.Debugf("nodeIDList: %v", nodeIDList)
	Convey("test FindNeighbor", t, func() {
		So(nodeIDList, ShouldContain, string(node1.ID))
		So(nodeIDList, ShouldContain, string(node2.ID))
	})

	reqFN3 := &FindNodeReq{
		NodeID: node2.ID,
	}
	respFN3 := new(FindNodeResp)
	err = client.Call("DHT.FindNode", reqFN3, respFN3)
	if err != nil {
		log.Error(err)
	}
	log.Debugf("respFN1: %v", respFN3)
	Convey("test FindNode", t, func() {
		So(respFN3.Node.ID, ShouldEqual, node2.ID)
		So(respFN3.Node.Nonce == node2.Nonce, ShouldBeTrue)
		So(respFN3.Node.Addr, ShouldEqual, node2.Addr)
		So(respFN3.Node.PublicKey.IsEqual(node2.PublicKey), ShouldBeTrue)
		So(err, ShouldBeNil)
	})

	client.Close()
}

func TestDHTService_Ping(t *testing.T) {
	os.Remove(DHTStorePath)
	defer os.Remove(DHTStorePath)
	log.SetLevel(log.DebugLevel)
	addr := "127.0.0.1:0"

	dht, _ := NewDHTService(DHTStorePath, new(consistent.KMSStorage), false)
	rpc.RegisterName("DHT", dht)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	mh := &codec.MsgpackHandle{}

	go func() {
		for {
			c, err := ln.Accept()
			if err != nil {
				continue
			}
			msgpackCodec := codec.MsgpackSpecRpc.ServerCodec(c, mh)
			go rpc.ServeCodec(msgpackCodec)
		}
	}()

	client, err := net.Dial("tcp", ln.Addr().String())
	if err != nil {
		log.Error(err)
	}

	NewNodeIDDifficultyTimeout = 100 * time.Millisecond
	node1 := NewNode()
	node1.InitNodeCryptoInfo(100 * time.Millisecond)

	reqA := &PingReq{
		Node: *node1,
	}
	respA := new(PingResp)
	msgpackCodec := codec.MsgpackSpecRpc.ClientCodec(client, mh)
	err = rpc.NewClientWithCodec(msgpackCodec).Call("DHT.Ping", reqA, respA)
	if err != nil {
		log.Error(err)
	}
	log.Debugf("respA: %v", respA)
}
