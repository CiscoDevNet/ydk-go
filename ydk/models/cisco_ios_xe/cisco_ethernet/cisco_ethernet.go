// This module contains a collection of YANG definitions for
// configuring Ethernet Interfaces.
// Copyright (c) 2016 by Cisco Systems, Inc.
// All rights reserved.
package cisco_ethernet

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_ethernet"))
}

type EthIfSpeed1gb struct {
}

func (id EthIfSpeed1gb) String() string {
	return "cisco-ethernet:eth-if-speed-1gb"
}

type EthIfSpeed40gb struct {
}

func (id EthIfSpeed40gb) String() string {
	return "cisco-ethernet:eth-if-speed-40gb"
}

type EthIfSpeed10mb struct {
}

func (id EthIfSpeed10mb) String() string {
	return "cisco-ethernet:eth-if-speed-10mb"
}

type EthIfSpeed10gb struct {
}

func (id EthIfSpeed10gb) String() string {
	return "cisco-ethernet:eth-if-speed-10gb"
}

type EthIfSpeed100gb struct {
}

func (id EthIfSpeed100gb) String() string {
	return "cisco-ethernet:eth-if-speed-100gb"
}

type EthIfSpeed100mb struct {
}

func (id EthIfSpeed100mb) String() string {
	return "cisco-ethernet:eth-if-speed-100mb"
}

type EthIfSpeed struct {
}

func (id EthIfSpeed) String() string {
	return "cisco-ethernet:eth-if-speed"
}

