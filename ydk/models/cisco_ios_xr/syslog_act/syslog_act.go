// This module contains a collection of YANG definitions
// for Cisco IOS-XR action package configuration.
// 
// This module contains definitions
// for the following management objects:
// syslog: Global Syslog messaging data
// 
// Copyright (c) 2016 by Cisco Systems, Inc.
// All rights reserved.
package syslog_act

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package syslog_act"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-syslog-act logmsg}", reflect.TypeOf(Logmsg{}))
    ydk.RegisterEntity("Cisco-IOS-XR-syslog-act:logmsg", reflect.TypeOf(Logmsg{}))
}

// Logmsg
type Logmsg struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input Logmsg_Input
}

func (logmsg *Logmsg) GetEntityData() *types.CommonEntityData {
    logmsg.EntityData.YFilter = logmsg.YFilter
    logmsg.EntityData.YangName = "logmsg"
    logmsg.EntityData.BundleName = "cisco_ios_xr"
    logmsg.EntityData.ParentYangName = "Cisco-IOS-XR-syslog-act"
    logmsg.EntityData.SegmentPath = "Cisco-IOS-XR-syslog-act:logmsg"
    logmsg.EntityData.AbsolutePath = logmsg.EntityData.SegmentPath
    logmsg.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logmsg.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logmsg.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logmsg.EntityData.Children = types.NewOrderedMap()
    logmsg.EntityData.Children.Append("input", types.YChild{"Input", &logmsg.Input})
    logmsg.EntityData.Leafs = types.NewOrderedMap()

    logmsg.EntityData.YListKeys = []string {}

    return &(logmsg.EntityData)
}

// Logmsg_Input
type Logmsg_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set serverity level. The type is Severity. This attribute is mandatory.
    Severity interface{}

    // Message body. The type is string. This attribute is mandatory.
    Message interface{}
}

func (input *Logmsg_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "logmsg"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "Cisco-IOS-XR-syslog-act:logmsg/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("severity", types.YLeaf{"Severity", input.Severity})
    input.EntityData.Leafs.Append("message", types.YLeaf{"Message", input.Message})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

