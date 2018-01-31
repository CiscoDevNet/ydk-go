// This module contains a collection of YANG definitions
// for Cisco IOS-XR man-xml-ttyagent package configuration.
// 
// This module contains definitions
// for the following management objects:
//   xr-xml: XML
//   netconf: netconf
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package man_xml_ttyagent_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package man_xml_ttyagent_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-man-xml-ttyagent-cfg xr-xml}", reflect.TypeOf(XrXml{}))
    ydk.RegisterEntity("Cisco-IOS-XR-man-xml-ttyagent-cfg:xr-xml", reflect.TypeOf(XrXml{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-man-xml-ttyagent-cfg netconf}", reflect.TypeOf(Netconf{}))
    ydk.RegisterEntity("Cisco-IOS-XR-man-xml-ttyagent-cfg:netconf", reflect.TypeOf(Netconf{}))
}

// XrXml
// XML
type XrXml struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // XML agent.
    Agent XrXml_Agent
}

func (xrXml *XrXml) GetFilter() yfilter.YFilter { return xrXml.YFilter }

func (xrXml *XrXml) SetFilter(yf yfilter.YFilter) { xrXml.YFilter = yf }

func (xrXml *XrXml) GetGoName(yname string) string {
    if yname == "agent" { return "Agent" }
    return ""
}

func (xrXml *XrXml) GetSegmentPath() string {
    return "Cisco-IOS-XR-man-xml-ttyagent-cfg:xr-xml"
}

func (xrXml *XrXml) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "agent" {
        return &xrXml.Agent
    }
    return nil
}

func (xrXml *XrXml) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["agent"] = &xrXml.Agent
    return children
}

func (xrXml *XrXml) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (xrXml *XrXml) GetBundleName() string { return "cisco_ios_xr" }

func (xrXml *XrXml) GetYangName() string { return "xr-xml" }

func (xrXml *XrXml) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (xrXml *XrXml) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (xrXml *XrXml) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (xrXml *XrXml) SetParent(parent types.Entity) { xrXml.parent = parent }

func (xrXml *XrXml) GetParent() types.Entity { return xrXml.parent }

func (xrXml *XrXml) GetParentYangName() string { return "Cisco-IOS-XR-man-xml-ttyagent-cfg" }

// XrXml_Agent
// XML agent
type XrXml_Agent struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // XML default dedicated agent.
    Default XrXml_Agent_Default

    // XML TTY agent.
    Tty XrXml_Agent_Tty

    // XML SSL agent.
    Ssl XrXml_Agent_Ssl
}

func (agent *XrXml_Agent) GetFilter() yfilter.YFilter { return agent.YFilter }

func (agent *XrXml_Agent) SetFilter(yf yfilter.YFilter) { agent.YFilter = yf }

func (agent *XrXml_Agent) GetGoName(yname string) string {
    if yname == "default" { return "Default" }
    if yname == "tty" { return "Tty" }
    if yname == "ssl" { return "Ssl" }
    return ""
}

func (agent *XrXml_Agent) GetSegmentPath() string {
    return "agent"
}

func (agent *XrXml_Agent) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "default" {
        return &agent.Default
    }
    if childYangName == "tty" {
        return &agent.Tty
    }
    if childYangName == "ssl" {
        return &agent.Ssl
    }
    return nil
}

func (agent *XrXml_Agent) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["default"] = &agent.Default
    children["tty"] = &agent.Tty
    children["ssl"] = &agent.Ssl
    return children
}

func (agent *XrXml_Agent) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (agent *XrXml_Agent) GetBundleName() string { return "cisco_ios_xr" }

func (agent *XrXml_Agent) GetYangName() string { return "agent" }

func (agent *XrXml_Agent) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (agent *XrXml_Agent) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (agent *XrXml_Agent) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (agent *XrXml_Agent) SetParent(parent types.Entity) { agent.parent = parent }

func (agent *XrXml_Agent) GetParent() types.Entity { return agent.parent }

func (agent *XrXml_Agent) GetParentYangName() string { return "xr-xml" }

// XrXml_Agent_Default
// XML default dedicated agent
type XrXml_Agent_Default struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 Transport State. The type is bool.
    Ipv6Enable interface{}

    // TRUE to disable IPV4. The type is bool.
    Ipv4Disable interface{}

    // Iterator size, in KBytes, of the XML response. Specify 0 to turn off the
    // XML response iterator. The type is interface{} with range: 0..100000. Units
    // are kilobyte. The default value is 48.
    IterationSize interface{}

    // Enable specified XML agent. The type is interface{}.
    Enable interface{}

    // Streaming size, in KBytes, of the XML response. The type is interface{}
    // with range: 1..100000. Units are kilobyte.
    StreamingSize interface{}

    // Session attributes.
    Session XrXml_Agent_Default_Session

    // XML agent throttling.
    Throttle XrXml_Agent_Default_Throttle

    // List of VRFs.
    Vrfs XrXml_Agent_Default_Vrfs
}

func (self *XrXml_Agent_Default) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *XrXml_Agent_Default) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *XrXml_Agent_Default) GetGoName(yname string) string {
    if yname == "ipv6-enable" { return "Ipv6Enable" }
    if yname == "ipv4-disable" { return "Ipv4Disable" }
    if yname == "iteration-size" { return "IterationSize" }
    if yname == "enable" { return "Enable" }
    if yname == "streaming-size" { return "StreamingSize" }
    if yname == "session" { return "Session" }
    if yname == "throttle" { return "Throttle" }
    if yname == "vrfs" { return "Vrfs" }
    return ""
}

func (self *XrXml_Agent_Default) GetSegmentPath() string {
    return "default"
}

func (self *XrXml_Agent_Default) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "session" {
        return &self.Session
    }
    if childYangName == "throttle" {
        return &self.Throttle
    }
    if childYangName == "vrfs" {
        return &self.Vrfs
    }
    return nil
}

func (self *XrXml_Agent_Default) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["session"] = &self.Session
    children["throttle"] = &self.Throttle
    children["vrfs"] = &self.Vrfs
    return children
}

func (self *XrXml_Agent_Default) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv6-enable"] = self.Ipv6Enable
    leafs["ipv4-disable"] = self.Ipv4Disable
    leafs["iteration-size"] = self.IterationSize
    leafs["enable"] = self.Enable
    leafs["streaming-size"] = self.StreamingSize
    return leafs
}

func (self *XrXml_Agent_Default) GetBundleName() string { return "cisco_ios_xr" }

func (self *XrXml_Agent_Default) GetYangName() string { return "default" }

func (self *XrXml_Agent_Default) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *XrXml_Agent_Default) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *XrXml_Agent_Default) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *XrXml_Agent_Default) SetParent(parent types.Entity) { self.parent = parent }

func (self *XrXml_Agent_Default) GetParent() types.Entity { return self.parent }

func (self *XrXml_Agent_Default) GetParentYangName() string { return "agent" }

// XrXml_Agent_Default_Session
// Session attributes
type XrXml_Agent_Default_Session struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Timeout in minutes. The type is interface{} with range: 1..1440. Units are
    // minute.
    Timeout interface{}
}

func (session *XrXml_Agent_Default_Session) GetFilter() yfilter.YFilter { return session.YFilter }

func (session *XrXml_Agent_Default_Session) SetFilter(yf yfilter.YFilter) { session.YFilter = yf }

func (session *XrXml_Agent_Default_Session) GetGoName(yname string) string {
    if yname == "timeout" { return "Timeout" }
    return ""
}

func (session *XrXml_Agent_Default_Session) GetSegmentPath() string {
    return "session"
}

func (session *XrXml_Agent_Default_Session) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (session *XrXml_Agent_Default_Session) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (session *XrXml_Agent_Default_Session) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["timeout"] = session.Timeout
    return leafs
}

func (session *XrXml_Agent_Default_Session) GetBundleName() string { return "cisco_ios_xr" }

func (session *XrXml_Agent_Default_Session) GetYangName() string { return "session" }

func (session *XrXml_Agent_Default_Session) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (session *XrXml_Agent_Default_Session) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (session *XrXml_Agent_Default_Session) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (session *XrXml_Agent_Default_Session) SetParent(parent types.Entity) { session.parent = parent }

func (session *XrXml_Agent_Default_Session) GetParent() types.Entity { return session.parent }

func (session *XrXml_Agent_Default_Session) GetParentYangName() string { return "default" }

// XrXml_Agent_Default_Throttle
// XML agent throttling
type XrXml_Agent_Default_Throttle struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Process rate in number of XML tags per second. The type is interface{} with
    // range: 1000..30000.
    ProcessRate interface{}

    // Size of memory usage, in MBytes, per session. The type is interface{} with
    // range: 100..600. Units are megabyte. The default value is 300.
    Memory interface{}
}

func (throttle *XrXml_Agent_Default_Throttle) GetFilter() yfilter.YFilter { return throttle.YFilter }

func (throttle *XrXml_Agent_Default_Throttle) SetFilter(yf yfilter.YFilter) { throttle.YFilter = yf }

func (throttle *XrXml_Agent_Default_Throttle) GetGoName(yname string) string {
    if yname == "process-rate" { return "ProcessRate" }
    if yname == "memory" { return "Memory" }
    return ""
}

func (throttle *XrXml_Agent_Default_Throttle) GetSegmentPath() string {
    return "throttle"
}

func (throttle *XrXml_Agent_Default_Throttle) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (throttle *XrXml_Agent_Default_Throttle) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (throttle *XrXml_Agent_Default_Throttle) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["process-rate"] = throttle.ProcessRate
    leafs["memory"] = throttle.Memory
    return leafs
}

func (throttle *XrXml_Agent_Default_Throttle) GetBundleName() string { return "cisco_ios_xr" }

func (throttle *XrXml_Agent_Default_Throttle) GetYangName() string { return "throttle" }

func (throttle *XrXml_Agent_Default_Throttle) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (throttle *XrXml_Agent_Default_Throttle) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (throttle *XrXml_Agent_Default_Throttle) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (throttle *XrXml_Agent_Default_Throttle) SetParent(parent types.Entity) { throttle.parent = parent }

func (throttle *XrXml_Agent_Default_Throttle) GetParent() types.Entity { return throttle.parent }

func (throttle *XrXml_Agent_Default_Throttle) GetParentYangName() string { return "default" }

// XrXml_Agent_Default_Vrfs
// List of VRFs
type XrXml_Agent_Default_Vrfs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A specific VRF. The type is slice of XrXml_Agent_Default_Vrfs_Vrf.
    Vrf []XrXml_Agent_Default_Vrfs_Vrf
}

func (vrfs *XrXml_Agent_Default_Vrfs) GetFilter() yfilter.YFilter { return vrfs.YFilter }

func (vrfs *XrXml_Agent_Default_Vrfs) SetFilter(yf yfilter.YFilter) { vrfs.YFilter = yf }

func (vrfs *XrXml_Agent_Default_Vrfs) GetGoName(yname string) string {
    if yname == "vrf" { return "Vrf" }
    return ""
}

func (vrfs *XrXml_Agent_Default_Vrfs) GetSegmentPath() string {
    return "vrfs"
}

func (vrfs *XrXml_Agent_Default_Vrfs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf" {
        for _, c := range vrfs.Vrf {
            if vrfs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := XrXml_Agent_Default_Vrfs_Vrf{}
        vrfs.Vrf = append(vrfs.Vrf, child)
        return &vrfs.Vrf[len(vrfs.Vrf)-1]
    }
    return nil
}

func (vrfs *XrXml_Agent_Default_Vrfs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vrfs.Vrf {
        children[vrfs.Vrf[i].GetSegmentPath()] = &vrfs.Vrf[i]
    }
    return children
}

func (vrfs *XrXml_Agent_Default_Vrfs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vrfs *XrXml_Agent_Default_Vrfs) GetBundleName() string { return "cisco_ios_xr" }

func (vrfs *XrXml_Agent_Default_Vrfs) GetYangName() string { return "vrfs" }

func (vrfs *XrXml_Agent_Default_Vrfs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfs *XrXml_Agent_Default_Vrfs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfs *XrXml_Agent_Default_Vrfs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfs *XrXml_Agent_Default_Vrfs) SetParent(parent types.Entity) { vrfs.parent = parent }

func (vrfs *XrXml_Agent_Default_Vrfs) GetParent() types.Entity { return vrfs.parent }

func (vrfs *XrXml_Agent_Default_Vrfs) GetParentYangName() string { return "default" }

// XrXml_Agent_Default_Vrfs_Vrf
// A specific VRF
type XrXml_Agent_Default_Vrfs_Vrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with length: 1..32.
    VrfName interface{}

    // IPv6 Transport Access list for VRF. The type is string with length: 1..32.
    Ipv6AccessList interface{}

    // IPv4 Transport Access list for VRF. The type is string with length: 1..32.
    Ipv4AccessList interface{}

    // Access list for XML agent. The type is string with length: 1..32.
    AccessList interface{}

    // Shutdown default VRF. This is applicable only for VRF default. The type is
    // interface{}.
    Shutdown interface{}
}

func (vrf *XrXml_Agent_Default_Vrfs_Vrf) GetFilter() yfilter.YFilter { return vrf.YFilter }

func (vrf *XrXml_Agent_Default_Vrfs_Vrf) SetFilter(yf yfilter.YFilter) { vrf.YFilter = yf }

func (vrf *XrXml_Agent_Default_Vrfs_Vrf) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "ipv6-access-list" { return "Ipv6AccessList" }
    if yname == "ipv4-access-list" { return "Ipv4AccessList" }
    if yname == "access-list" { return "AccessList" }
    if yname == "shutdown" { return "Shutdown" }
    return ""
}

func (vrf *XrXml_Agent_Default_Vrfs_Vrf) GetSegmentPath() string {
    return "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
}

func (vrf *XrXml_Agent_Default_Vrfs_Vrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vrf *XrXml_Agent_Default_Vrfs_Vrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vrf *XrXml_Agent_Default_Vrfs_Vrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = vrf.VrfName
    leafs["ipv6-access-list"] = vrf.Ipv6AccessList
    leafs["ipv4-access-list"] = vrf.Ipv4AccessList
    leafs["access-list"] = vrf.AccessList
    leafs["shutdown"] = vrf.Shutdown
    return leafs
}

func (vrf *XrXml_Agent_Default_Vrfs_Vrf) GetBundleName() string { return "cisco_ios_xr" }

func (vrf *XrXml_Agent_Default_Vrfs_Vrf) GetYangName() string { return "vrf" }

func (vrf *XrXml_Agent_Default_Vrfs_Vrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrf *XrXml_Agent_Default_Vrfs_Vrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrf *XrXml_Agent_Default_Vrfs_Vrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrf *XrXml_Agent_Default_Vrfs_Vrf) SetParent(parent types.Entity) { vrf.parent = parent }

func (vrf *XrXml_Agent_Default_Vrfs_Vrf) GetParent() types.Entity { return vrf.parent }

func (vrf *XrXml_Agent_Default_Vrfs_Vrf) GetParentYangName() string { return "vrfs" }

// XrXml_Agent_Tty
// XML TTY agent
type XrXml_Agent_Tty struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Iterator size, in KBytes, of the XML response. Specify 0 to turn off the
    // XML response iterator. The type is interface{} with range: 0..100000. Units
    // are kilobyte. The default value is 48.
    IterationSize interface{}

    // Enable specified XML agent. The type is interface{}.
    Enable interface{}

    // Streaming size, in KBytes, of the XML response. The type is interface{}
    // with range: 1..100000. Units are kilobyte.
    StreamingSize interface{}

    // Session attributes.
    Session XrXml_Agent_Tty_Session

    // XML agent throttling.
    Throttle XrXml_Agent_Tty_Throttle
}

func (tty *XrXml_Agent_Tty) GetFilter() yfilter.YFilter { return tty.YFilter }

func (tty *XrXml_Agent_Tty) SetFilter(yf yfilter.YFilter) { tty.YFilter = yf }

func (tty *XrXml_Agent_Tty) GetGoName(yname string) string {
    if yname == "iteration-size" { return "IterationSize" }
    if yname == "enable" { return "Enable" }
    if yname == "streaming-size" { return "StreamingSize" }
    if yname == "session" { return "Session" }
    if yname == "throttle" { return "Throttle" }
    return ""
}

func (tty *XrXml_Agent_Tty) GetSegmentPath() string {
    return "tty"
}

func (tty *XrXml_Agent_Tty) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "session" {
        return &tty.Session
    }
    if childYangName == "throttle" {
        return &tty.Throttle
    }
    return nil
}

func (tty *XrXml_Agent_Tty) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["session"] = &tty.Session
    children["throttle"] = &tty.Throttle
    return children
}

func (tty *XrXml_Agent_Tty) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["iteration-size"] = tty.IterationSize
    leafs["enable"] = tty.Enable
    leafs["streaming-size"] = tty.StreamingSize
    return leafs
}

func (tty *XrXml_Agent_Tty) GetBundleName() string { return "cisco_ios_xr" }

func (tty *XrXml_Agent_Tty) GetYangName() string { return "tty" }

func (tty *XrXml_Agent_Tty) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tty *XrXml_Agent_Tty) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tty *XrXml_Agent_Tty) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tty *XrXml_Agent_Tty) SetParent(parent types.Entity) { tty.parent = parent }

func (tty *XrXml_Agent_Tty) GetParent() types.Entity { return tty.parent }

func (tty *XrXml_Agent_Tty) GetParentYangName() string { return "agent" }

// XrXml_Agent_Tty_Session
// Session attributes
type XrXml_Agent_Tty_Session struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Timeout in minutes. The type is interface{} with range: 1..1440. Units are
    // minute.
    Timeout interface{}
}

func (session *XrXml_Agent_Tty_Session) GetFilter() yfilter.YFilter { return session.YFilter }

func (session *XrXml_Agent_Tty_Session) SetFilter(yf yfilter.YFilter) { session.YFilter = yf }

func (session *XrXml_Agent_Tty_Session) GetGoName(yname string) string {
    if yname == "timeout" { return "Timeout" }
    return ""
}

func (session *XrXml_Agent_Tty_Session) GetSegmentPath() string {
    return "session"
}

func (session *XrXml_Agent_Tty_Session) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (session *XrXml_Agent_Tty_Session) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (session *XrXml_Agent_Tty_Session) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["timeout"] = session.Timeout
    return leafs
}

func (session *XrXml_Agent_Tty_Session) GetBundleName() string { return "cisco_ios_xr" }

func (session *XrXml_Agent_Tty_Session) GetYangName() string { return "session" }

func (session *XrXml_Agent_Tty_Session) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (session *XrXml_Agent_Tty_Session) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (session *XrXml_Agent_Tty_Session) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (session *XrXml_Agent_Tty_Session) SetParent(parent types.Entity) { session.parent = parent }

func (session *XrXml_Agent_Tty_Session) GetParent() types.Entity { return session.parent }

func (session *XrXml_Agent_Tty_Session) GetParentYangName() string { return "tty" }

// XrXml_Agent_Tty_Throttle
// XML agent throttling
type XrXml_Agent_Tty_Throttle struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Process rate in number of XML tags per second. The type is interface{} with
    // range: 1000..30000.
    ProcessRate interface{}

    // Size of memory usage, in MBytes, per session. The type is interface{} with
    // range: 100..600. Units are megabyte. The default value is 300.
    Memory interface{}
}

func (throttle *XrXml_Agent_Tty_Throttle) GetFilter() yfilter.YFilter { return throttle.YFilter }

func (throttle *XrXml_Agent_Tty_Throttle) SetFilter(yf yfilter.YFilter) { throttle.YFilter = yf }

func (throttle *XrXml_Agent_Tty_Throttle) GetGoName(yname string) string {
    if yname == "process-rate" { return "ProcessRate" }
    if yname == "memory" { return "Memory" }
    return ""
}

func (throttle *XrXml_Agent_Tty_Throttle) GetSegmentPath() string {
    return "throttle"
}

func (throttle *XrXml_Agent_Tty_Throttle) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (throttle *XrXml_Agent_Tty_Throttle) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (throttle *XrXml_Agent_Tty_Throttle) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["process-rate"] = throttle.ProcessRate
    leafs["memory"] = throttle.Memory
    return leafs
}

func (throttle *XrXml_Agent_Tty_Throttle) GetBundleName() string { return "cisco_ios_xr" }

func (throttle *XrXml_Agent_Tty_Throttle) GetYangName() string { return "throttle" }

func (throttle *XrXml_Agent_Tty_Throttle) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (throttle *XrXml_Agent_Tty_Throttle) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (throttle *XrXml_Agent_Tty_Throttle) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (throttle *XrXml_Agent_Tty_Throttle) SetParent(parent types.Entity) { throttle.parent = parent }

func (throttle *XrXml_Agent_Tty_Throttle) GetParent() types.Entity { return throttle.parent }

func (throttle *XrXml_Agent_Tty_Throttle) GetParentYangName() string { return "tty" }

// XrXml_Agent_Ssl
// XML SSL agent
type XrXml_Agent_Ssl struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Iterator size, in KBytes, of the XML response. Specify 0 to turn off the
    // XML response iterator. The type is interface{} with range: 0..100000. Units
    // are kilobyte. The default value is 48.
    IterationSize interface{}

    // Enable specified XML agent. The type is interface{}.
    Enable interface{}

    // Streaming size, in KBytes, of the XML response. The type is interface{}
    // with range: 1..100000. Units are kilobyte.
    StreamingSize interface{}

    // Session attributes.
    Session XrXml_Agent_Ssl_Session

    // XML agent throttling.
    Throttle XrXml_Agent_Ssl_Throttle

    // List of VRFs.
    Vrfs XrXml_Agent_Ssl_Vrfs
}

func (ssl *XrXml_Agent_Ssl) GetFilter() yfilter.YFilter { return ssl.YFilter }

func (ssl *XrXml_Agent_Ssl) SetFilter(yf yfilter.YFilter) { ssl.YFilter = yf }

func (ssl *XrXml_Agent_Ssl) GetGoName(yname string) string {
    if yname == "iteration-size" { return "IterationSize" }
    if yname == "enable" { return "Enable" }
    if yname == "streaming-size" { return "StreamingSize" }
    if yname == "session" { return "Session" }
    if yname == "throttle" { return "Throttle" }
    if yname == "vrfs" { return "Vrfs" }
    return ""
}

func (ssl *XrXml_Agent_Ssl) GetSegmentPath() string {
    return "ssl"
}

func (ssl *XrXml_Agent_Ssl) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "session" {
        return &ssl.Session
    }
    if childYangName == "throttle" {
        return &ssl.Throttle
    }
    if childYangName == "vrfs" {
        return &ssl.Vrfs
    }
    return nil
}

func (ssl *XrXml_Agent_Ssl) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["session"] = &ssl.Session
    children["throttle"] = &ssl.Throttle
    children["vrfs"] = &ssl.Vrfs
    return children
}

func (ssl *XrXml_Agent_Ssl) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["iteration-size"] = ssl.IterationSize
    leafs["enable"] = ssl.Enable
    leafs["streaming-size"] = ssl.StreamingSize
    return leafs
}

func (ssl *XrXml_Agent_Ssl) GetBundleName() string { return "cisco_ios_xr" }

func (ssl *XrXml_Agent_Ssl) GetYangName() string { return "ssl" }

func (ssl *XrXml_Agent_Ssl) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ssl *XrXml_Agent_Ssl) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ssl *XrXml_Agent_Ssl) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ssl *XrXml_Agent_Ssl) SetParent(parent types.Entity) { ssl.parent = parent }

func (ssl *XrXml_Agent_Ssl) GetParent() types.Entity { return ssl.parent }

func (ssl *XrXml_Agent_Ssl) GetParentYangName() string { return "agent" }

// XrXml_Agent_Ssl_Session
// Session attributes
type XrXml_Agent_Ssl_Session struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Timeout in minutes. The type is interface{} with range: 1..1440. Units are
    // minute.
    Timeout interface{}
}

func (session *XrXml_Agent_Ssl_Session) GetFilter() yfilter.YFilter { return session.YFilter }

func (session *XrXml_Agent_Ssl_Session) SetFilter(yf yfilter.YFilter) { session.YFilter = yf }

func (session *XrXml_Agent_Ssl_Session) GetGoName(yname string) string {
    if yname == "timeout" { return "Timeout" }
    return ""
}

func (session *XrXml_Agent_Ssl_Session) GetSegmentPath() string {
    return "session"
}

func (session *XrXml_Agent_Ssl_Session) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (session *XrXml_Agent_Ssl_Session) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (session *XrXml_Agent_Ssl_Session) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["timeout"] = session.Timeout
    return leafs
}

func (session *XrXml_Agent_Ssl_Session) GetBundleName() string { return "cisco_ios_xr" }

func (session *XrXml_Agent_Ssl_Session) GetYangName() string { return "session" }

func (session *XrXml_Agent_Ssl_Session) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (session *XrXml_Agent_Ssl_Session) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (session *XrXml_Agent_Ssl_Session) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (session *XrXml_Agent_Ssl_Session) SetParent(parent types.Entity) { session.parent = parent }

func (session *XrXml_Agent_Ssl_Session) GetParent() types.Entity { return session.parent }

func (session *XrXml_Agent_Ssl_Session) GetParentYangName() string { return "ssl" }

// XrXml_Agent_Ssl_Throttle
// XML agent throttling
type XrXml_Agent_Ssl_Throttle struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Process rate in number of XML tags per second. The type is interface{} with
    // range: 1000..30000.
    ProcessRate interface{}

    // Size of memory usage, in MBytes, per session. The type is interface{} with
    // range: 100..600. Units are megabyte. The default value is 300.
    Memory interface{}
}

func (throttle *XrXml_Agent_Ssl_Throttle) GetFilter() yfilter.YFilter { return throttle.YFilter }

func (throttle *XrXml_Agent_Ssl_Throttle) SetFilter(yf yfilter.YFilter) { throttle.YFilter = yf }

func (throttle *XrXml_Agent_Ssl_Throttle) GetGoName(yname string) string {
    if yname == "process-rate" { return "ProcessRate" }
    if yname == "memory" { return "Memory" }
    return ""
}

func (throttle *XrXml_Agent_Ssl_Throttle) GetSegmentPath() string {
    return "throttle"
}

func (throttle *XrXml_Agent_Ssl_Throttle) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (throttle *XrXml_Agent_Ssl_Throttle) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (throttle *XrXml_Agent_Ssl_Throttle) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["process-rate"] = throttle.ProcessRate
    leafs["memory"] = throttle.Memory
    return leafs
}

func (throttle *XrXml_Agent_Ssl_Throttle) GetBundleName() string { return "cisco_ios_xr" }

func (throttle *XrXml_Agent_Ssl_Throttle) GetYangName() string { return "throttle" }

func (throttle *XrXml_Agent_Ssl_Throttle) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (throttle *XrXml_Agent_Ssl_Throttle) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (throttle *XrXml_Agent_Ssl_Throttle) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (throttle *XrXml_Agent_Ssl_Throttle) SetParent(parent types.Entity) { throttle.parent = parent }

func (throttle *XrXml_Agent_Ssl_Throttle) GetParent() types.Entity { return throttle.parent }

func (throttle *XrXml_Agent_Ssl_Throttle) GetParentYangName() string { return "ssl" }

// XrXml_Agent_Ssl_Vrfs
// List of VRFs
type XrXml_Agent_Ssl_Vrfs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A specific VRF. The type is slice of XrXml_Agent_Ssl_Vrfs_Vrf.
    Vrf []XrXml_Agent_Ssl_Vrfs_Vrf
}

func (vrfs *XrXml_Agent_Ssl_Vrfs) GetFilter() yfilter.YFilter { return vrfs.YFilter }

func (vrfs *XrXml_Agent_Ssl_Vrfs) SetFilter(yf yfilter.YFilter) { vrfs.YFilter = yf }

func (vrfs *XrXml_Agent_Ssl_Vrfs) GetGoName(yname string) string {
    if yname == "vrf" { return "Vrf" }
    return ""
}

func (vrfs *XrXml_Agent_Ssl_Vrfs) GetSegmentPath() string {
    return "vrfs"
}

func (vrfs *XrXml_Agent_Ssl_Vrfs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf" {
        for _, c := range vrfs.Vrf {
            if vrfs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := XrXml_Agent_Ssl_Vrfs_Vrf{}
        vrfs.Vrf = append(vrfs.Vrf, child)
        return &vrfs.Vrf[len(vrfs.Vrf)-1]
    }
    return nil
}

func (vrfs *XrXml_Agent_Ssl_Vrfs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vrfs.Vrf {
        children[vrfs.Vrf[i].GetSegmentPath()] = &vrfs.Vrf[i]
    }
    return children
}

func (vrfs *XrXml_Agent_Ssl_Vrfs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vrfs *XrXml_Agent_Ssl_Vrfs) GetBundleName() string { return "cisco_ios_xr" }

func (vrfs *XrXml_Agent_Ssl_Vrfs) GetYangName() string { return "vrfs" }

func (vrfs *XrXml_Agent_Ssl_Vrfs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfs *XrXml_Agent_Ssl_Vrfs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfs *XrXml_Agent_Ssl_Vrfs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfs *XrXml_Agent_Ssl_Vrfs) SetParent(parent types.Entity) { vrfs.parent = parent }

func (vrfs *XrXml_Agent_Ssl_Vrfs) GetParent() types.Entity { return vrfs.parent }

func (vrfs *XrXml_Agent_Ssl_Vrfs) GetParentYangName() string { return "ssl" }

// XrXml_Agent_Ssl_Vrfs_Vrf
// A specific VRF
type XrXml_Agent_Ssl_Vrfs_Vrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with length: 1..32.
    VrfName interface{}

    // IPv6 Transport Access list for VRF. The type is string with length: 1..32.
    Ipv6AccessList interface{}

    // IPv4 Transport Access list for VRF. The type is string with length: 1..32.
    Ipv4AccessList interface{}

    // Access list for XML agent. The type is string with length: 1..32.
    AccessList interface{}

    // Shutdown default VRF. This is applicable only for VRF default. The type is
    // interface{}.
    Shutdown interface{}
}

func (vrf *XrXml_Agent_Ssl_Vrfs_Vrf) GetFilter() yfilter.YFilter { return vrf.YFilter }

func (vrf *XrXml_Agent_Ssl_Vrfs_Vrf) SetFilter(yf yfilter.YFilter) { vrf.YFilter = yf }

func (vrf *XrXml_Agent_Ssl_Vrfs_Vrf) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "ipv6-access-list" { return "Ipv6AccessList" }
    if yname == "ipv4-access-list" { return "Ipv4AccessList" }
    if yname == "access-list" { return "AccessList" }
    if yname == "shutdown" { return "Shutdown" }
    return ""
}

func (vrf *XrXml_Agent_Ssl_Vrfs_Vrf) GetSegmentPath() string {
    return "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
}

func (vrf *XrXml_Agent_Ssl_Vrfs_Vrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vrf *XrXml_Agent_Ssl_Vrfs_Vrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vrf *XrXml_Agent_Ssl_Vrfs_Vrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = vrf.VrfName
    leafs["ipv6-access-list"] = vrf.Ipv6AccessList
    leafs["ipv4-access-list"] = vrf.Ipv4AccessList
    leafs["access-list"] = vrf.AccessList
    leafs["shutdown"] = vrf.Shutdown
    return leafs
}

func (vrf *XrXml_Agent_Ssl_Vrfs_Vrf) GetBundleName() string { return "cisco_ios_xr" }

func (vrf *XrXml_Agent_Ssl_Vrfs_Vrf) GetYangName() string { return "vrf" }

func (vrf *XrXml_Agent_Ssl_Vrfs_Vrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrf *XrXml_Agent_Ssl_Vrfs_Vrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrf *XrXml_Agent_Ssl_Vrfs_Vrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrf *XrXml_Agent_Ssl_Vrfs_Vrf) SetParent(parent types.Entity) { vrf.parent = parent }

func (vrf *XrXml_Agent_Ssl_Vrfs_Vrf) GetParent() types.Entity { return vrf.parent }

func (vrf *XrXml_Agent_Ssl_Vrfs_Vrf) GetParentYangName() string { return "vrfs" }

// Netconf
// netconf
type Netconf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // XML agent.
    Agent Netconf_Agent
}

func (netconf *Netconf) GetFilter() yfilter.YFilter { return netconf.YFilter }

func (netconf *Netconf) SetFilter(yf yfilter.YFilter) { netconf.YFilter = yf }

func (netconf *Netconf) GetGoName(yname string) string {
    if yname == "agent" { return "Agent" }
    return ""
}

func (netconf *Netconf) GetSegmentPath() string {
    return "Cisco-IOS-XR-man-xml-ttyagent-cfg:netconf"
}

func (netconf *Netconf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "agent" {
        return &netconf.Agent
    }
    return nil
}

func (netconf *Netconf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["agent"] = &netconf.Agent
    return children
}

func (netconf *Netconf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (netconf *Netconf) GetBundleName() string { return "cisco_ios_xr" }

func (netconf *Netconf) GetYangName() string { return "netconf" }

func (netconf *Netconf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (netconf *Netconf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (netconf *Netconf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (netconf *Netconf) SetParent(parent types.Entity) { netconf.parent = parent }

func (netconf *Netconf) GetParent() types.Entity { return netconf.parent }

func (netconf *Netconf) GetParentYangName() string { return "Cisco-IOS-XR-man-xml-ttyagent-cfg" }

// Netconf_Agent
// XML agent
type Netconf_Agent struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // NETCONF agent over TTY.
    Tty Netconf_Agent_Tty
}

func (agent *Netconf_Agent) GetFilter() yfilter.YFilter { return agent.YFilter }

func (agent *Netconf_Agent) SetFilter(yf yfilter.YFilter) { agent.YFilter = yf }

func (agent *Netconf_Agent) GetGoName(yname string) string {
    if yname == "tty" { return "Tty" }
    return ""
}

func (agent *Netconf_Agent) GetSegmentPath() string {
    return "agent"
}

func (agent *Netconf_Agent) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tty" {
        return &agent.Tty
    }
    return nil
}

func (agent *Netconf_Agent) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["tty"] = &agent.Tty
    return children
}

func (agent *Netconf_Agent) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (agent *Netconf_Agent) GetBundleName() string { return "cisco_ios_xr" }

func (agent *Netconf_Agent) GetYangName() string { return "agent" }

func (agent *Netconf_Agent) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (agent *Netconf_Agent) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (agent *Netconf_Agent) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (agent *Netconf_Agent) SetParent(parent types.Entity) { agent.parent = parent }

func (agent *Netconf_Agent) GetParent() types.Entity { return agent.parent }

func (agent *Netconf_Agent) GetParentYangName() string { return "netconf" }

// Netconf_Agent_Tty
// NETCONF agent over TTY
type Netconf_Agent_Tty struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable specified NETCONF agent. The type is interface{}.
    Enable interface{}

    // NETCONF agent throttling.
    Throttle Netconf_Agent_Tty_Throttle

    // Session attributes.
    Session Netconf_Agent_Tty_Session
}

func (tty *Netconf_Agent_Tty) GetFilter() yfilter.YFilter { return tty.YFilter }

func (tty *Netconf_Agent_Tty) SetFilter(yf yfilter.YFilter) { tty.YFilter = yf }

func (tty *Netconf_Agent_Tty) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "throttle" { return "Throttle" }
    if yname == "session" { return "Session" }
    return ""
}

func (tty *Netconf_Agent_Tty) GetSegmentPath() string {
    return "tty"
}

func (tty *Netconf_Agent_Tty) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "throttle" {
        return &tty.Throttle
    }
    if childYangName == "session" {
        return &tty.Session
    }
    return nil
}

func (tty *Netconf_Agent_Tty) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["throttle"] = &tty.Throttle
    children["session"] = &tty.Session
    return children
}

func (tty *Netconf_Agent_Tty) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = tty.Enable
    return leafs
}

func (tty *Netconf_Agent_Tty) GetBundleName() string { return "cisco_ios_xr" }

func (tty *Netconf_Agent_Tty) GetYangName() string { return "tty" }

func (tty *Netconf_Agent_Tty) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tty *Netconf_Agent_Tty) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tty *Netconf_Agent_Tty) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tty *Netconf_Agent_Tty) SetParent(parent types.Entity) { tty.parent = parent }

func (tty *Netconf_Agent_Tty) GetParent() types.Entity { return tty.parent }

func (tty *Netconf_Agent_Tty) GetParentYangName() string { return "agent" }

// Netconf_Agent_Tty_Throttle
// NETCONF agent throttling
type Netconf_Agent_Tty_Throttle struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Size of memory usage, in MBytes, per session. The type is interface{} with
    // range: 100..600. Units are megabyte. The default value is 300.
    Memory interface{}

    // Size of memory usage, in MBytes, per session. The type is interface{} with
    // range: 0..12000. Units are megabyte. The default value is 0.
    OffloadMemory interface{}

    // Process rate in number of XML tags per second. The type is interface{} with
    // range: 1000..30000.
    ProcessRate interface{}
}

func (throttle *Netconf_Agent_Tty_Throttle) GetFilter() yfilter.YFilter { return throttle.YFilter }

func (throttle *Netconf_Agent_Tty_Throttle) SetFilter(yf yfilter.YFilter) { throttle.YFilter = yf }

func (throttle *Netconf_Agent_Tty_Throttle) GetGoName(yname string) string {
    if yname == "memory" { return "Memory" }
    if yname == "offload-memory" { return "OffloadMemory" }
    if yname == "process-rate" { return "ProcessRate" }
    return ""
}

func (throttle *Netconf_Agent_Tty_Throttle) GetSegmentPath() string {
    return "throttle"
}

func (throttle *Netconf_Agent_Tty_Throttle) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (throttle *Netconf_Agent_Tty_Throttle) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (throttle *Netconf_Agent_Tty_Throttle) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["memory"] = throttle.Memory
    leafs["offload-memory"] = throttle.OffloadMemory
    leafs["process-rate"] = throttle.ProcessRate
    return leafs
}

func (throttle *Netconf_Agent_Tty_Throttle) GetBundleName() string { return "cisco_ios_xr" }

func (throttle *Netconf_Agent_Tty_Throttle) GetYangName() string { return "throttle" }

func (throttle *Netconf_Agent_Tty_Throttle) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (throttle *Netconf_Agent_Tty_Throttle) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (throttle *Netconf_Agent_Tty_Throttle) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (throttle *Netconf_Agent_Tty_Throttle) SetParent(parent types.Entity) { throttle.parent = parent }

func (throttle *Netconf_Agent_Tty_Throttle) GetParent() types.Entity { return throttle.parent }

func (throttle *Netconf_Agent_Tty_Throttle) GetParentYangName() string { return "tty" }

// Netconf_Agent_Tty_Session
// Session attributes
type Netconf_Agent_Tty_Session struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Timeout in minutes. The type is interface{} with range: 1..1440. Units are
    // minute.
    Timeout interface{}
}

func (session *Netconf_Agent_Tty_Session) GetFilter() yfilter.YFilter { return session.YFilter }

func (session *Netconf_Agent_Tty_Session) SetFilter(yf yfilter.YFilter) { session.YFilter = yf }

func (session *Netconf_Agent_Tty_Session) GetGoName(yname string) string {
    if yname == "timeout" { return "Timeout" }
    return ""
}

func (session *Netconf_Agent_Tty_Session) GetSegmentPath() string {
    return "session"
}

func (session *Netconf_Agent_Tty_Session) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (session *Netconf_Agent_Tty_Session) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (session *Netconf_Agent_Tty_Session) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["timeout"] = session.Timeout
    return leafs
}

func (session *Netconf_Agent_Tty_Session) GetBundleName() string { return "cisco_ios_xr" }

func (session *Netconf_Agent_Tty_Session) GetYangName() string { return "session" }

func (session *Netconf_Agent_Tty_Session) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (session *Netconf_Agent_Tty_Session) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (session *Netconf_Agent_Tty_Session) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (session *Netconf_Agent_Tty_Session) SetParent(parent types.Entity) { session.parent = parent }

func (session *Netconf_Agent_Tty_Session) GetParent() types.Entity { return session.parent }

func (session *Netconf_Agent_Tty_Session) GetParentYangName() string { return "tty" }

