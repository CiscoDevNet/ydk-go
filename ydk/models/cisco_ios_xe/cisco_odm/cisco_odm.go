// Copyright (c) 2016-2017 by Cisco Systems, Inc.All rights reserved.
package cisco_odm

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_odm"))
}

type FlowMonitor struct {
}

func (id FlowMonitor) String() string {
	return "cisco-odm:FlowMonitor"
}

type VirtualService struct {
}

func (id VirtualService) String() string {
	return "cisco-odm:VirtualService"
}

type Diffserv struct {
}

func (id Diffserv) String() string {
	return "cisco-odm:Diffserv"
}

type Parsername struct {
}

func (id Parsername) String() string {
	return "cisco-odm:parsername"
}

type BridgeDomain struct {
}

func (id BridgeDomain) String() string {
	return "cisco-odm:BridgeDomain"
}

type MPLSStaticBinding struct {
}

func (id MPLSStaticBinding) String() string {
	return "cisco-odm:MPLSStaticBinding"
}

type EthernetCFMStats struct {
}

func (id EthernetCFMStats) String() string {
	return "cisco-odm:EthernetCFMStats"
}

type MPLSForwardingTable struct {
}

func (id MPLSForwardingTable) String() string {
	return "cisco-odm:MPLSForwardingTable"
}

type BGP struct {
}

func (id BGP) String() string {
	return "cisco-odm:BGP"
}

type IPRoute struct {
}

func (id IPRoute) String() string {
	return "cisco-odm:IPRoute"
}

type PlatformSoftware struct {
}

func (id PlatformSoftware) String() string {
	return "cisco-odm:PlatformSoftware"
}

type MPLSLDPNeighbors struct {
}

func (id MPLSLDPNeighbors) String() string {
	return "cisco-odm:MPLSLDPNeighbors"
}

type BFDNeighbors struct {
}

func (id BFDNeighbors) String() string {
	return "cisco-odm:BFDNeighbors"
}

type OSPF struct {
}

func (id OSPF) String() string {
	return "cisco-odm:OSPF"
}

