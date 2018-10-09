// This module contains a collection of YANG definitions for
// monitoring breakout ports in a Network Element.
// Copyright (c) 2018 by Cisco Systems, Inc.
// All rights reserved.
package breakout_port_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package breakout_port_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-breakout-port-oper breakout-port-oper-data}", reflect.TypeOf(BreakoutPortOperData{}))
    ydk.RegisterEntity("Cisco-IOS-XE-breakout-port-oper:breakout-port-oper-data", reflect.TypeOf(BreakoutPortOperData{}))
}

// BcChannelSpeed represents Speed of each channel
type BcChannelSpeed string

const (
    BcChannelSpeed_Y_10gb BcChannelSpeed = "10gb"

    BcChannelSpeed_unknown BcChannelSpeed = "unknown"
)

// BreakoutPortOperData
// Informaton about breakout ports
type BreakoutPortOperData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of breakout ports, keyed by name. The type is slice of
    // BreakoutPortOperData_PortBreakout.
    PortBreakout []*BreakoutPortOperData_PortBreakout
}

func (breakoutPortOperData *BreakoutPortOperData) GetEntityData() *types.CommonEntityData {
    breakoutPortOperData.EntityData.YFilter = breakoutPortOperData.YFilter
    breakoutPortOperData.EntityData.YangName = "breakout-port-oper-data"
    breakoutPortOperData.EntityData.BundleName = "cisco_ios_xe"
    breakoutPortOperData.EntityData.ParentYangName = "Cisco-IOS-XE-breakout-port-oper"
    breakoutPortOperData.EntityData.SegmentPath = "Cisco-IOS-XE-breakout-port-oper:breakout-port-oper-data"
    breakoutPortOperData.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    breakoutPortOperData.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    breakoutPortOperData.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    breakoutPortOperData.EntityData.Children = types.NewOrderedMap()
    breakoutPortOperData.EntityData.Children.Append("port-breakout", types.YChild{"PortBreakout", nil})
    for i := range breakoutPortOperData.PortBreakout {
        breakoutPortOperData.EntityData.Children.Append(types.GetSegmentPath(breakoutPortOperData.PortBreakout[i]), types.YChild{"PortBreakout", breakoutPortOperData.PortBreakout[i]})
    }
    breakoutPortOperData.EntityData.Leafs = types.NewOrderedMap()

    breakoutPortOperData.EntityData.YListKeys = []string {}

    return &(breakoutPortOperData.EntityData)
}

// BreakoutPortOperData_PortBreakout
// List of breakout ports, keyed by name
type BreakoutPortOperData_PortBreakout struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the breakout port. The type is string.
    Name interface{}

    // Number of channels to 'breakout' on a port capable of channelization. The
    // type is interface{} with range: -32768..32767.
    Number interface{}

    // Channel speed on each channel. The type is BcChannelSpeed.
    Speed interface{}
}

func (portBreakout *BreakoutPortOperData_PortBreakout) GetEntityData() *types.CommonEntityData {
    portBreakout.EntityData.YFilter = portBreakout.YFilter
    portBreakout.EntityData.YangName = "port-breakout"
    portBreakout.EntityData.BundleName = "cisco_ios_xe"
    portBreakout.EntityData.ParentYangName = "breakout-port-oper-data"
    portBreakout.EntityData.SegmentPath = "port-breakout" + types.AddKeyToken(portBreakout.Name, "name")
    portBreakout.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    portBreakout.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    portBreakout.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    portBreakout.EntityData.Children = types.NewOrderedMap()
    portBreakout.EntityData.Leafs = types.NewOrderedMap()
    portBreakout.EntityData.Leafs.Append("name", types.YLeaf{"Name", portBreakout.Name})
    portBreakout.EntityData.Leafs.Append("number", types.YLeaf{"Number", portBreakout.Number})
    portBreakout.EntityData.Leafs.Append("speed", types.YLeaf{"Speed", portBreakout.Speed})

    portBreakout.EntityData.YListKeys = []string {"Name"}

    return &(portBreakout.EntityData)
}

