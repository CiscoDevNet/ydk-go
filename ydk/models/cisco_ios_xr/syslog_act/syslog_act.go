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
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input Logmsg_Input
}

func (logmsg *Logmsg) GetFilter() yfilter.YFilter { return logmsg.YFilter }

func (logmsg *Logmsg) SetFilter(yf yfilter.YFilter) { logmsg.YFilter = yf }

func (logmsg *Logmsg) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (logmsg *Logmsg) GetSegmentPath() string {
    return "Cisco-IOS-XR-syslog-act:logmsg"
}

func (logmsg *Logmsg) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &logmsg.Input
    }
    return nil
}

func (logmsg *Logmsg) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &logmsg.Input
    return children
}

func (logmsg *Logmsg) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (logmsg *Logmsg) GetBundleName() string { return "cisco_ios_xr" }

func (logmsg *Logmsg) GetYangName() string { return "logmsg" }

func (logmsg *Logmsg) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (logmsg *Logmsg) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (logmsg *Logmsg) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (logmsg *Logmsg) SetParent(parent types.Entity) { logmsg.parent = parent }

func (logmsg *Logmsg) GetParent() types.Entity { return logmsg.parent }

func (logmsg *Logmsg) GetParentYangName() string { return "Cisco-IOS-XR-syslog-act" }

// Logmsg_Input
type Logmsg_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Set serverity level. The type is Severity. This attribute is mandatory.
    Severity interface{}

    // Message body. The type is string. This attribute is mandatory.
    Message interface{}
}

func (input *Logmsg_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *Logmsg_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *Logmsg_Input) GetGoName(yname string) string {
    if yname == "severity" { return "Severity" }
    if yname == "message" { return "Message" }
    return ""
}

func (input *Logmsg_Input) GetSegmentPath() string {
    return "input"
}

func (input *Logmsg_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *Logmsg_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *Logmsg_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["severity"] = input.Severity
    leafs["message"] = input.Message
    return leafs
}

func (input *Logmsg_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *Logmsg_Input) GetYangName() string { return "input" }

func (input *Logmsg_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *Logmsg_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *Logmsg_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *Logmsg_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *Logmsg_Input) GetParent() types.Entity { return input.parent }

func (input *Logmsg_Input) GetParentYangName() string { return "logmsg" }

