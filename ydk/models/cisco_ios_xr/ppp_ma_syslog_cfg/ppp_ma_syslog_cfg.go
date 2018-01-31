// This module contains a collection of YANG definitions
// for Cisco IOS-XR ppp-ma-syslog package configuration.
// 
// This module contains definitions
// for the following management objects:
//   ppp: PPP configuration
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ppp_ma_syslog_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ppp_ma_syslog_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ppp-ma-syslog-cfg ppp}", reflect.TypeOf(Ppp{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ppp-ma-syslog-cfg:ppp", reflect.TypeOf(Ppp{}))
}

// Ppp
// PPP configuration
type Ppp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // syslog option for session status.
    Syslog Ppp_Syslog
}

func (ppp *Ppp) GetFilter() yfilter.YFilter { return ppp.YFilter }

func (ppp *Ppp) SetFilter(yf yfilter.YFilter) { ppp.YFilter = yf }

func (ppp *Ppp) GetGoName(yname string) string {
    if yname == "syslog" { return "Syslog" }
    return ""
}

func (ppp *Ppp) GetSegmentPath() string {
    return "Cisco-IOS-XR-ppp-ma-syslog-cfg:ppp"
}

func (ppp *Ppp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "syslog" {
        return &ppp.Syslog
    }
    return nil
}

func (ppp *Ppp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["syslog"] = &ppp.Syslog
    return children
}

func (ppp *Ppp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ppp *Ppp) GetBundleName() string { return "cisco_ios_xr" }

func (ppp *Ppp) GetYangName() string { return "ppp" }

func (ppp *Ppp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ppp *Ppp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ppp *Ppp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ppp *Ppp) SetParent(parent types.Entity) { ppp.parent = parent }

func (ppp *Ppp) GetParent() types.Entity { return ppp.parent }

func (ppp *Ppp) GetParentYangName() string { return "Cisco-IOS-XR-ppp-ma-syslog-cfg" }

// Ppp_Syslog
// syslog option for session status
type Ppp_Syslog struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable syslog for ppp session status. The type is interface{}.
    EnableSessionStatus interface{}
}

func (syslog *Ppp_Syslog) GetFilter() yfilter.YFilter { return syslog.YFilter }

func (syslog *Ppp_Syslog) SetFilter(yf yfilter.YFilter) { syslog.YFilter = yf }

func (syslog *Ppp_Syslog) GetGoName(yname string) string {
    if yname == "enable-session-status" { return "EnableSessionStatus" }
    return ""
}

func (syslog *Ppp_Syslog) GetSegmentPath() string {
    return "syslog"
}

func (syslog *Ppp_Syslog) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (syslog *Ppp_Syslog) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (syslog *Ppp_Syslog) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable-session-status"] = syslog.EnableSessionStatus
    return leafs
}

func (syslog *Ppp_Syslog) GetBundleName() string { return "cisco_ios_xr" }

func (syslog *Ppp_Syslog) GetYangName() string { return "syslog" }

func (syslog *Ppp_Syslog) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (syslog *Ppp_Syslog) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (syslog *Ppp_Syslog) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (syslog *Ppp_Syslog) SetParent(parent types.Entity) { syslog.parent = parent }

func (syslog *Ppp_Syslog) GetParent() types.Entity { return syslog.parent }

func (syslog *Ppp_Syslog) GetParentYangName() string { return "ppp" }

