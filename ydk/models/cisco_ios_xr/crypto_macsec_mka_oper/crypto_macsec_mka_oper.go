// This module contains a collection of YANG definitions
// for Cisco IOS-XR crypto-macsec-mka package operational data.
// 
// This module contains definitions
// for the following management objects:
//   macsec: Macsec operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package crypto_macsec_mka_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package crypto_macsec_mka_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-crypto-macsec-mka-oper macsec}", reflect.TypeOf(Macsec{}))
    ydk.RegisterEntity("Cisco-IOS-XR-crypto-macsec-mka-oper:macsec", reflect.TypeOf(Macsec{}))
}

// Macsec
// Macsec operational data
type Macsec struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MKA Data.
    Mka Macsec_Mka
}

func (macsec *Macsec) GetFilter() yfilter.YFilter { return macsec.YFilter }

func (macsec *Macsec) SetFilter(yf yfilter.YFilter) { macsec.YFilter = yf }

func (macsec *Macsec) GetGoName(yname string) string {
    if yname == "mka" { return "Mka" }
    return ""
}

func (macsec *Macsec) GetSegmentPath() string {
    return "Cisco-IOS-XR-crypto-macsec-mka-oper:macsec"
}

func (macsec *Macsec) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mka" {
        return &macsec.Mka
    }
    return nil
}

func (macsec *Macsec) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["mka"] = &macsec.Mka
    return children
}

func (macsec *Macsec) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (macsec *Macsec) GetBundleName() string { return "cisco_ios_xr" }

func (macsec *Macsec) GetYangName() string { return "macsec" }

func (macsec *Macsec) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (macsec *Macsec) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (macsec *Macsec) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (macsec *Macsec) SetParent(parent types.Entity) { macsec.parent = parent }

func (macsec *Macsec) GetParent() types.Entity { return macsec.parent }

func (macsec *Macsec) GetParentYangName() string { return "Cisco-IOS-XR-crypto-macsec-mka-oper" }

// Macsec_Mka
// MKA Data
type Macsec_Mka struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MKA Data.
    Interfaces Macsec_Mka_Interfaces
}

func (mka *Macsec_Mka) GetFilter() yfilter.YFilter { return mka.YFilter }

func (mka *Macsec_Mka) SetFilter(yf yfilter.YFilter) { mka.YFilter = yf }

func (mka *Macsec_Mka) GetGoName(yname string) string {
    if yname == "interfaces" { return "Interfaces" }
    return ""
}

func (mka *Macsec_Mka) GetSegmentPath() string {
    return "mka"
}

func (mka *Macsec_Mka) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interfaces" {
        return &mka.Interfaces
    }
    return nil
}

func (mka *Macsec_Mka) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interfaces"] = &mka.Interfaces
    return children
}

func (mka *Macsec_Mka) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mka *Macsec_Mka) GetBundleName() string { return "cisco_ios_xr" }

func (mka *Macsec_Mka) GetYangName() string { return "mka" }

func (mka *Macsec_Mka) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mka *Macsec_Mka) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mka *Macsec_Mka) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mka *Macsec_Mka) SetParent(parent types.Entity) { mka.parent = parent }

func (mka *Macsec_Mka) GetParent() types.Entity { return mka.parent }

func (mka *Macsec_Mka) GetParentYangName() string { return "macsec" }

// Macsec_Mka_Interfaces
// MKA Data
type Macsec_Mka_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MKA Data for the Interface. The type is slice of
    // Macsec_Mka_Interfaces_Interface.
    Interface []Macsec_Mka_Interfaces_Interface
}

func (interfaces *Macsec_Mka_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *Macsec_Mka_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *Macsec_Mka_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *Macsec_Mka_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *Macsec_Mka_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Macsec_Mka_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *Macsec_Mka_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *Macsec_Mka_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *Macsec_Mka_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *Macsec_Mka_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *Macsec_Mka_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *Macsec_Mka_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *Macsec_Mka_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *Macsec_Mka_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *Macsec_Mka_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *Macsec_Mka_Interfaces) GetParentYangName() string { return "mka" }

// Macsec_Mka_Interfaces_Interface
// MKA Data for the Interface
type Macsec_Mka_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    Name interface{}

    // MKA Session Data.
    Session Macsec_Mka_Interfaces_Interface_Session
}

func (self *Macsec_Mka_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Macsec_Mka_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Macsec_Mka_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "session" { return "Session" }
    return ""
}

func (self *Macsec_Mka_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[name='" + fmt.Sprintf("%v", self.Name) + "']"
}

func (self *Macsec_Mka_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "session" {
        return &self.Session
    }
    return nil
}

func (self *Macsec_Mka_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["session"] = &self.Session
    return children
}

func (self *Macsec_Mka_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = self.Name
    return leafs
}

func (self *Macsec_Mka_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *Macsec_Mka_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *Macsec_Mka_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Macsec_Mka_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Macsec_Mka_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Macsec_Mka_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *Macsec_Mka_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *Macsec_Mka_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// Macsec_Mka_Interfaces_Interface_Session
// MKA Session Data
type Macsec_Mka_Interfaces_Interface_Session struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Session summary.
    SessionSummary Macsec_Mka_Interfaces_Interface_Session_SessionSummary

    // Virtual Pointer Info.
    Vp Macsec_Mka_Interfaces_Interface_Session_Vp

    // CA List for a Session. The type is slice of
    // Macsec_Mka_Interfaces_Interface_Session_Ca.
    Ca []Macsec_Mka_Interfaces_Interface_Session_Ca
}

func (session *Macsec_Mka_Interfaces_Interface_Session) GetFilter() yfilter.YFilter { return session.YFilter }

func (session *Macsec_Mka_Interfaces_Interface_Session) SetFilter(yf yfilter.YFilter) { session.YFilter = yf }

func (session *Macsec_Mka_Interfaces_Interface_Session) GetGoName(yname string) string {
    if yname == "session-summary" { return "SessionSummary" }
    if yname == "vp" { return "Vp" }
    if yname == "ca" { return "Ca" }
    return ""
}

func (session *Macsec_Mka_Interfaces_Interface_Session) GetSegmentPath() string {
    return "session"
}

func (session *Macsec_Mka_Interfaces_Interface_Session) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "session-summary" {
        return &session.SessionSummary
    }
    if childYangName == "vp" {
        return &session.Vp
    }
    if childYangName == "ca" {
        for _, c := range session.Ca {
            if session.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Macsec_Mka_Interfaces_Interface_Session_Ca{}
        session.Ca = append(session.Ca, child)
        return &session.Ca[len(session.Ca)-1]
    }
    return nil
}

func (session *Macsec_Mka_Interfaces_Interface_Session) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["session-summary"] = &session.SessionSummary
    children["vp"] = &session.Vp
    for i := range session.Ca {
        children[session.Ca[i].GetSegmentPath()] = &session.Ca[i]
    }
    return children
}

func (session *Macsec_Mka_Interfaces_Interface_Session) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (session *Macsec_Mka_Interfaces_Interface_Session) GetBundleName() string { return "cisco_ios_xr" }

func (session *Macsec_Mka_Interfaces_Interface_Session) GetYangName() string { return "session" }

func (session *Macsec_Mka_Interfaces_Interface_Session) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (session *Macsec_Mka_Interfaces_Interface_Session) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (session *Macsec_Mka_Interfaces_Interface_Session) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (session *Macsec_Mka_Interfaces_Interface_Session) SetParent(parent types.Entity) { session.parent = parent }

func (session *Macsec_Mka_Interfaces_Interface_Session) GetParent() types.Entity { return session.parent }

func (session *Macsec_Mka_Interfaces_Interface_Session) GetParentYangName() string { return "interface" }

// Macsec_Mka_Interfaces_Interface_Session_SessionSummary
// Session summary
type Macsec_Mka_Interfaces_Interface_Session_SessionSummary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // macsec configured interface. The type is string.
    InterfaceName interface{}

    // Is Inherited Policy. The type is bool.
    InheritedPolicy interface{}

    // Policy Name. The type is string.
    Policy interface{}

    // Key Server Priority. The type is interface{} with range: 0..4294967295.
    Priority interface{}

    // My MAC. The type is string.
    MyMac interface{}

    // Delay Protect. The type is bool.
    DelayProtection interface{}

    // Replay Protect. The type is bool.
    ReplayProtect interface{}

    // Replay Window Size. The type is interface{} with range: 0..4294967295.
    WindowSize interface{}

    // IncludeICVIndicator. The type is bool.
    IncludeIcvIndicator interface{}

    // Confidentiality Offset. The type is interface{} with range: 0..4294967295.
    ConfidentialityOffset interface{}

    // Alogorithm Agility. The type is interface{} with range: 0..4294967295.
    AlgoAgility interface{}

    // MACSec Capability. The type is interface{} with range: 0..4294967295.
    Capability interface{}

    // Cipher String. The type is string.
    CipherStr interface{}

    // MACSec Desired. The type is bool.
    MacSecDesired interface{}

    // VLAN Outer TAG.
    OuterTag Macsec_Mka_Interfaces_Interface_Session_SessionSummary_OuterTag

    // VLAN Inner TAG.
    InnerTag Macsec_Mka_Interfaces_Interface_Session_SessionSummary_InnerTag
}

func (sessionSummary *Macsec_Mka_Interfaces_Interface_Session_SessionSummary) GetFilter() yfilter.YFilter { return sessionSummary.YFilter }

func (sessionSummary *Macsec_Mka_Interfaces_Interface_Session_SessionSummary) SetFilter(yf yfilter.YFilter) { sessionSummary.YFilter = yf }

func (sessionSummary *Macsec_Mka_Interfaces_Interface_Session_SessionSummary) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "inherited-policy" { return "InheritedPolicy" }
    if yname == "policy" { return "Policy" }
    if yname == "priority" { return "Priority" }
    if yname == "my-mac" { return "MyMac" }
    if yname == "delay-protection" { return "DelayProtection" }
    if yname == "replay-protect" { return "ReplayProtect" }
    if yname == "window-size" { return "WindowSize" }
    if yname == "include-icv-indicator" { return "IncludeIcvIndicator" }
    if yname == "confidentiality-offset" { return "ConfidentialityOffset" }
    if yname == "algo-agility" { return "AlgoAgility" }
    if yname == "capability" { return "Capability" }
    if yname == "cipher-str" { return "CipherStr" }
    if yname == "mac-sec-desired" { return "MacSecDesired" }
    if yname == "outer-tag" { return "OuterTag" }
    if yname == "inner-tag" { return "InnerTag" }
    return ""
}

func (sessionSummary *Macsec_Mka_Interfaces_Interface_Session_SessionSummary) GetSegmentPath() string {
    return "session-summary"
}

func (sessionSummary *Macsec_Mka_Interfaces_Interface_Session_SessionSummary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "outer-tag" {
        return &sessionSummary.OuterTag
    }
    if childYangName == "inner-tag" {
        return &sessionSummary.InnerTag
    }
    return nil
}

func (sessionSummary *Macsec_Mka_Interfaces_Interface_Session_SessionSummary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["outer-tag"] = &sessionSummary.OuterTag
    children["inner-tag"] = &sessionSummary.InnerTag
    return children
}

func (sessionSummary *Macsec_Mka_Interfaces_Interface_Session_SessionSummary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = sessionSummary.InterfaceName
    leafs["inherited-policy"] = sessionSummary.InheritedPolicy
    leafs["policy"] = sessionSummary.Policy
    leafs["priority"] = sessionSummary.Priority
    leafs["my-mac"] = sessionSummary.MyMac
    leafs["delay-protection"] = sessionSummary.DelayProtection
    leafs["replay-protect"] = sessionSummary.ReplayProtect
    leafs["window-size"] = sessionSummary.WindowSize
    leafs["include-icv-indicator"] = sessionSummary.IncludeIcvIndicator
    leafs["confidentiality-offset"] = sessionSummary.ConfidentialityOffset
    leafs["algo-agility"] = sessionSummary.AlgoAgility
    leafs["capability"] = sessionSummary.Capability
    leafs["cipher-str"] = sessionSummary.CipherStr
    leafs["mac-sec-desired"] = sessionSummary.MacSecDesired
    return leafs
}

func (sessionSummary *Macsec_Mka_Interfaces_Interface_Session_SessionSummary) GetBundleName() string { return "cisco_ios_xr" }

func (sessionSummary *Macsec_Mka_Interfaces_Interface_Session_SessionSummary) GetYangName() string { return "session-summary" }

func (sessionSummary *Macsec_Mka_Interfaces_Interface_Session_SessionSummary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sessionSummary *Macsec_Mka_Interfaces_Interface_Session_SessionSummary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sessionSummary *Macsec_Mka_Interfaces_Interface_Session_SessionSummary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sessionSummary *Macsec_Mka_Interfaces_Interface_Session_SessionSummary) SetParent(parent types.Entity) { sessionSummary.parent = parent }

func (sessionSummary *Macsec_Mka_Interfaces_Interface_Session_SessionSummary) GetParent() types.Entity { return sessionSummary.parent }

func (sessionSummary *Macsec_Mka_Interfaces_Interface_Session_SessionSummary) GetParentYangName() string { return "session" }

// Macsec_Mka_Interfaces_Interface_Session_SessionSummary_OuterTag
// VLAN Outer TAG
type Macsec_Mka_Interfaces_Interface_Session_SessionSummary_OuterTag struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // etype. The type is interface{} with range: 0..65535.
    Etype interface{}

    // priority. The type is interface{} with range: 0..255.
    Priority interface{}

    // cfi. The type is interface{} with range: 0..255.
    Cfi interface{}

    // vlan id. The type is interface{} with range: 0..65535.
    VlanId interface{}
}

func (outerTag *Macsec_Mka_Interfaces_Interface_Session_SessionSummary_OuterTag) GetFilter() yfilter.YFilter { return outerTag.YFilter }

func (outerTag *Macsec_Mka_Interfaces_Interface_Session_SessionSummary_OuterTag) SetFilter(yf yfilter.YFilter) { outerTag.YFilter = yf }

func (outerTag *Macsec_Mka_Interfaces_Interface_Session_SessionSummary_OuterTag) GetGoName(yname string) string {
    if yname == "etype" { return "Etype" }
    if yname == "priority" { return "Priority" }
    if yname == "cfi" { return "Cfi" }
    if yname == "vlan-id" { return "VlanId" }
    return ""
}

func (outerTag *Macsec_Mka_Interfaces_Interface_Session_SessionSummary_OuterTag) GetSegmentPath() string {
    return "outer-tag"
}

func (outerTag *Macsec_Mka_Interfaces_Interface_Session_SessionSummary_OuterTag) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outerTag *Macsec_Mka_Interfaces_Interface_Session_SessionSummary_OuterTag) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outerTag *Macsec_Mka_Interfaces_Interface_Session_SessionSummary_OuterTag) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["etype"] = outerTag.Etype
    leafs["priority"] = outerTag.Priority
    leafs["cfi"] = outerTag.Cfi
    leafs["vlan-id"] = outerTag.VlanId
    return leafs
}

func (outerTag *Macsec_Mka_Interfaces_Interface_Session_SessionSummary_OuterTag) GetBundleName() string { return "cisco_ios_xr" }

func (outerTag *Macsec_Mka_Interfaces_Interface_Session_SessionSummary_OuterTag) GetYangName() string { return "outer-tag" }

func (outerTag *Macsec_Mka_Interfaces_Interface_Session_SessionSummary_OuterTag) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outerTag *Macsec_Mka_Interfaces_Interface_Session_SessionSummary_OuterTag) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outerTag *Macsec_Mka_Interfaces_Interface_Session_SessionSummary_OuterTag) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outerTag *Macsec_Mka_Interfaces_Interface_Session_SessionSummary_OuterTag) SetParent(parent types.Entity) { outerTag.parent = parent }

func (outerTag *Macsec_Mka_Interfaces_Interface_Session_SessionSummary_OuterTag) GetParent() types.Entity { return outerTag.parent }

func (outerTag *Macsec_Mka_Interfaces_Interface_Session_SessionSummary_OuterTag) GetParentYangName() string { return "session-summary" }

// Macsec_Mka_Interfaces_Interface_Session_SessionSummary_InnerTag
// VLAN Inner TAG
type Macsec_Mka_Interfaces_Interface_Session_SessionSummary_InnerTag struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // etype. The type is interface{} with range: 0..65535.
    Etype interface{}

    // priority. The type is interface{} with range: 0..255.
    Priority interface{}

    // cfi. The type is interface{} with range: 0..255.
    Cfi interface{}

    // vlan id. The type is interface{} with range: 0..65535.
    VlanId interface{}
}

func (innerTag *Macsec_Mka_Interfaces_Interface_Session_SessionSummary_InnerTag) GetFilter() yfilter.YFilter { return innerTag.YFilter }

func (innerTag *Macsec_Mka_Interfaces_Interface_Session_SessionSummary_InnerTag) SetFilter(yf yfilter.YFilter) { innerTag.YFilter = yf }

func (innerTag *Macsec_Mka_Interfaces_Interface_Session_SessionSummary_InnerTag) GetGoName(yname string) string {
    if yname == "etype" { return "Etype" }
    if yname == "priority" { return "Priority" }
    if yname == "cfi" { return "Cfi" }
    if yname == "vlan-id" { return "VlanId" }
    return ""
}

func (innerTag *Macsec_Mka_Interfaces_Interface_Session_SessionSummary_InnerTag) GetSegmentPath() string {
    return "inner-tag"
}

func (innerTag *Macsec_Mka_Interfaces_Interface_Session_SessionSummary_InnerTag) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (innerTag *Macsec_Mka_Interfaces_Interface_Session_SessionSummary_InnerTag) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (innerTag *Macsec_Mka_Interfaces_Interface_Session_SessionSummary_InnerTag) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["etype"] = innerTag.Etype
    leafs["priority"] = innerTag.Priority
    leafs["cfi"] = innerTag.Cfi
    leafs["vlan-id"] = innerTag.VlanId
    return leafs
}

func (innerTag *Macsec_Mka_Interfaces_Interface_Session_SessionSummary_InnerTag) GetBundleName() string { return "cisco_ios_xr" }

func (innerTag *Macsec_Mka_Interfaces_Interface_Session_SessionSummary_InnerTag) GetYangName() string { return "inner-tag" }

func (innerTag *Macsec_Mka_Interfaces_Interface_Session_SessionSummary_InnerTag) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (innerTag *Macsec_Mka_Interfaces_Interface_Session_SessionSummary_InnerTag) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (innerTag *Macsec_Mka_Interfaces_Interface_Session_SessionSummary_InnerTag) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (innerTag *Macsec_Mka_Interfaces_Interface_Session_SessionSummary_InnerTag) SetParent(parent types.Entity) { innerTag.parent = parent }

func (innerTag *Macsec_Mka_Interfaces_Interface_Session_SessionSummary_InnerTag) GetParent() types.Entity { return innerTag.parent }

func (innerTag *Macsec_Mka_Interfaces_Interface_Session_SessionSummary_InnerTag) GetParentYangName() string { return "session-summary" }

// Macsec_Mka_Interfaces_Interface_Session_Vp
// Virtual Pointer Info
type Macsec_Mka_Interfaces_Interface_Session_Vp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Local SCI(MAC). The type is string.
    MySci interface{}

    // Virtual Port ID. The type is interface{} with range: 0..4294967295.
    VirtualPortId interface{}

    // Latest Rx status. The type is bool.
    LatestRx interface{}

    // Latest Tx status. The type is bool.
    LatestTx interface{}

    // Latest SAK AN. The type is interface{} with range: 0..4294967295.
    LatestAn interface{}

    // Latest SAK KI. The type is string.
    LatestKi interface{}

    // Latest SAK KN. The type is interface{} with range: 0..4294967295.
    LatestKn interface{}

    // Old Rx status. The type is bool.
    OldRx interface{}

    // Old Tx status. The type is bool.
    OldTx interface{}

    // Old SAK AN. The type is interface{} with range: 0..4294967295.
    OldAn interface{}

    // Old SAK KI. The type is string.
    OldKi interface{}

    // Old SAK KN. The type is interface{} with range: 0..4294967295.
    OldKn interface{}

    // SAK Transmit Wait Time. The type is interface{} with range: 0..4294967295.
    WaitTime interface{}

    // SAK Retire time. The type is interface{} with range: 0..4294967295.
    RetireTime interface{}

    // SAK Cipher Suite. The type is interface{} with range: 0..4294967295.
    CipherSuite interface{}

    // SSCI of the Local TxSC. The type is interface{} with range: 0..4294967295.
    Ssci interface{}

    // Next SAK Rekey time in Sec. The type is string.
    TimeToSakRekey interface{}

    // Fallback Keepalive. The type is slice of
    // Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive.
    FallbackKeepalive []Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive
}

func (vp *Macsec_Mka_Interfaces_Interface_Session_Vp) GetFilter() yfilter.YFilter { return vp.YFilter }

func (vp *Macsec_Mka_Interfaces_Interface_Session_Vp) SetFilter(yf yfilter.YFilter) { vp.YFilter = yf }

func (vp *Macsec_Mka_Interfaces_Interface_Session_Vp) GetGoName(yname string) string {
    if yname == "my-sci" { return "MySci" }
    if yname == "virtual-port-id" { return "VirtualPortId" }
    if yname == "latest-rx" { return "LatestRx" }
    if yname == "latest-tx" { return "LatestTx" }
    if yname == "latest-an" { return "LatestAn" }
    if yname == "latest-ki" { return "LatestKi" }
    if yname == "latest-kn" { return "LatestKn" }
    if yname == "old-rx" { return "OldRx" }
    if yname == "old-tx" { return "OldTx" }
    if yname == "old-an" { return "OldAn" }
    if yname == "old-ki" { return "OldKi" }
    if yname == "old-kn" { return "OldKn" }
    if yname == "wait-time" { return "WaitTime" }
    if yname == "retire-time" { return "RetireTime" }
    if yname == "cipher-suite" { return "CipherSuite" }
    if yname == "ssci" { return "Ssci" }
    if yname == "time-to-sak-rekey" { return "TimeToSakRekey" }
    if yname == "fallback-keepalive" { return "FallbackKeepalive" }
    return ""
}

func (vp *Macsec_Mka_Interfaces_Interface_Session_Vp) GetSegmentPath() string {
    return "vp"
}

func (vp *Macsec_Mka_Interfaces_Interface_Session_Vp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "fallback-keepalive" {
        for _, c := range vp.FallbackKeepalive {
            if vp.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive{}
        vp.FallbackKeepalive = append(vp.FallbackKeepalive, child)
        return &vp.FallbackKeepalive[len(vp.FallbackKeepalive)-1]
    }
    return nil
}

func (vp *Macsec_Mka_Interfaces_Interface_Session_Vp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vp.FallbackKeepalive {
        children[vp.FallbackKeepalive[i].GetSegmentPath()] = &vp.FallbackKeepalive[i]
    }
    return children
}

func (vp *Macsec_Mka_Interfaces_Interface_Session_Vp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["my-sci"] = vp.MySci
    leafs["virtual-port-id"] = vp.VirtualPortId
    leafs["latest-rx"] = vp.LatestRx
    leafs["latest-tx"] = vp.LatestTx
    leafs["latest-an"] = vp.LatestAn
    leafs["latest-ki"] = vp.LatestKi
    leafs["latest-kn"] = vp.LatestKn
    leafs["old-rx"] = vp.OldRx
    leafs["old-tx"] = vp.OldTx
    leafs["old-an"] = vp.OldAn
    leafs["old-ki"] = vp.OldKi
    leafs["old-kn"] = vp.OldKn
    leafs["wait-time"] = vp.WaitTime
    leafs["retire-time"] = vp.RetireTime
    leafs["cipher-suite"] = vp.CipherSuite
    leafs["ssci"] = vp.Ssci
    leafs["time-to-sak-rekey"] = vp.TimeToSakRekey
    return leafs
}

func (vp *Macsec_Mka_Interfaces_Interface_Session_Vp) GetBundleName() string { return "cisco_ios_xr" }

func (vp *Macsec_Mka_Interfaces_Interface_Session_Vp) GetYangName() string { return "vp" }

func (vp *Macsec_Mka_Interfaces_Interface_Session_Vp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vp *Macsec_Mka_Interfaces_Interface_Session_Vp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vp *Macsec_Mka_Interfaces_Interface_Session_Vp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vp *Macsec_Mka_Interfaces_Interface_Session_Vp) SetParent(parent types.Entity) { vp.parent = parent }

func (vp *Macsec_Mka_Interfaces_Interface_Session_Vp) GetParent() types.Entity { return vp.parent }

func (vp *Macsec_Mka_Interfaces_Interface_Session_Vp) GetParentYangName() string { return "session" }

// Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive
// Fallback Keepalive
type Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // CKN. The type is string with length: 0..65.
    Ckn interface{}

    // Member Identifier. The type is string with length: 0..25.
    Mi interface{}

    // Message Number. The type is interface{} with range: 0..4294967295.
    Mn interface{}

    // Peers Status.
    PeersStatus Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus
}

func (fallbackKeepalive *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive) GetFilter() yfilter.YFilter { return fallbackKeepalive.YFilter }

func (fallbackKeepalive *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive) SetFilter(yf yfilter.YFilter) { fallbackKeepalive.YFilter = yf }

func (fallbackKeepalive *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive) GetGoName(yname string) string {
    if yname == "ckn" { return "Ckn" }
    if yname == "mi" { return "Mi" }
    if yname == "mn" { return "Mn" }
    if yname == "peers-status" { return "PeersStatus" }
    return ""
}

func (fallbackKeepalive *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive) GetSegmentPath() string {
    return "fallback-keepalive"
}

func (fallbackKeepalive *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "peers-status" {
        return &fallbackKeepalive.PeersStatus
    }
    return nil
}

func (fallbackKeepalive *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["peers-status"] = &fallbackKeepalive.PeersStatus
    return children
}

func (fallbackKeepalive *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ckn"] = fallbackKeepalive.Ckn
    leafs["mi"] = fallbackKeepalive.Mi
    leafs["mn"] = fallbackKeepalive.Mn
    return leafs
}

func (fallbackKeepalive *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive) GetBundleName() string { return "cisco_ios_xr" }

func (fallbackKeepalive *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive) GetYangName() string { return "fallback-keepalive" }

func (fallbackKeepalive *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fallbackKeepalive *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fallbackKeepalive *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fallbackKeepalive *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive) SetParent(parent types.Entity) { fallbackKeepalive.parent = parent }

func (fallbackKeepalive *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive) GetParent() types.Entity { return fallbackKeepalive.parent }

func (fallbackKeepalive *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive) GetParentYangName() string { return "vp" }

// Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus
// Peers Status
type Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Tx MKPDU Timestamp. The type is string with length: 0..128.
    TxMkpduTimestamp interface{}

    // Peer Count. The type is interface{} with range: 0..4294967295.
    PeerCount interface{}

    // Peer List. The type is slice of
    // Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus_Peer.
    Peer []Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus_Peer
}

func (peersStatus *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus) GetFilter() yfilter.YFilter { return peersStatus.YFilter }

func (peersStatus *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus) SetFilter(yf yfilter.YFilter) { peersStatus.YFilter = yf }

func (peersStatus *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus) GetGoName(yname string) string {
    if yname == "tx-mkpdu-timestamp" { return "TxMkpduTimestamp" }
    if yname == "peer-count" { return "PeerCount" }
    if yname == "peer" { return "Peer" }
    return ""
}

func (peersStatus *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus) GetSegmentPath() string {
    return "peers-status"
}

func (peersStatus *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "peer" {
        for _, c := range peersStatus.Peer {
            if peersStatus.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus_Peer{}
        peersStatus.Peer = append(peersStatus.Peer, child)
        return &peersStatus.Peer[len(peersStatus.Peer)-1]
    }
    return nil
}

func (peersStatus *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range peersStatus.Peer {
        children[peersStatus.Peer[i].GetSegmentPath()] = &peersStatus.Peer[i]
    }
    return children
}

func (peersStatus *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tx-mkpdu-timestamp"] = peersStatus.TxMkpduTimestamp
    leafs["peer-count"] = peersStatus.PeerCount
    return leafs
}

func (peersStatus *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus) GetBundleName() string { return "cisco_ios_xr" }

func (peersStatus *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus) GetYangName() string { return "peers-status" }

func (peersStatus *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peersStatus *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peersStatus *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peersStatus *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus) SetParent(parent types.Entity) { peersStatus.parent = parent }

func (peersStatus *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus) GetParent() types.Entity { return peersStatus.parent }

func (peersStatus *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus) GetParentYangName() string { return "fallback-keepalive" }

// Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus_Peer
// Peer List
type Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus_Peer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Rx SCI. The type is string with length: 0..17.
    Sci interface{}

    // Peer Status Data.
    PeerData Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus_Peer_PeerData
}

func (peer *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus_Peer) GetFilter() yfilter.YFilter { return peer.YFilter }

func (peer *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus_Peer) SetFilter(yf yfilter.YFilter) { peer.YFilter = yf }

func (peer *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus_Peer) GetGoName(yname string) string {
    if yname == "sci" { return "Sci" }
    if yname == "peer-data" { return "PeerData" }
    return ""
}

func (peer *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus_Peer) GetSegmentPath() string {
    return "peer"
}

func (peer *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus_Peer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "peer-data" {
        return &peer.PeerData
    }
    return nil
}

func (peer *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus_Peer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["peer-data"] = &peer.PeerData
    return children
}

func (peer *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus_Peer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sci"] = peer.Sci
    return leafs
}

func (peer *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus_Peer) GetBundleName() string { return "cisco_ios_xr" }

func (peer *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus_Peer) GetYangName() string { return "peer" }

func (peer *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus_Peer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peer *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus_Peer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peer *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus_Peer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peer *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus_Peer) SetParent(parent types.Entity) { peer.parent = parent }

func (peer *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus_Peer) GetParent() types.Entity { return peer.parent }

func (peer *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus_Peer) GetParentYangName() string { return "peers-status" }

// Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus_Peer_PeerData
// Peer Status Data
type Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus_Peer_PeerData struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Member ID. The type is string with length: 0..25.
    Mi interface{}

    // ICV Status. The type is string with length: 0..10.
    IcvStatus interface{}

    // ICV Check Timestamp. The type is string with length: 0..128.
    IcvCheckTimestamp interface{}
}

func (peerData *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus_Peer_PeerData) GetFilter() yfilter.YFilter { return peerData.YFilter }

func (peerData *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus_Peer_PeerData) SetFilter(yf yfilter.YFilter) { peerData.YFilter = yf }

func (peerData *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus_Peer_PeerData) GetGoName(yname string) string {
    if yname == "mi" { return "Mi" }
    if yname == "icv-status" { return "IcvStatus" }
    if yname == "icv-check-timestamp" { return "IcvCheckTimestamp" }
    return ""
}

func (peerData *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus_Peer_PeerData) GetSegmentPath() string {
    return "peer-data"
}

func (peerData *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus_Peer_PeerData) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (peerData *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus_Peer_PeerData) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (peerData *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus_Peer_PeerData) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mi"] = peerData.Mi
    leafs["icv-status"] = peerData.IcvStatus
    leafs["icv-check-timestamp"] = peerData.IcvCheckTimestamp
    return leafs
}

func (peerData *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus_Peer_PeerData) GetBundleName() string { return "cisco_ios_xr" }

func (peerData *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus_Peer_PeerData) GetYangName() string { return "peer-data" }

func (peerData *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus_Peer_PeerData) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peerData *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus_Peer_PeerData) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peerData *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus_Peer_PeerData) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peerData *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus_Peer_PeerData) SetParent(parent types.Entity) { peerData.parent = parent }

func (peerData *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus_Peer_PeerData) GetParent() types.Entity { return peerData.parent }

func (peerData *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus_Peer_PeerData) GetParentYangName() string { return "peer" }

// Macsec_Mka_Interfaces_Interface_Session_Ca
// CA List for a Session
type Macsec_Mka_Interfaces_Interface_Session_Ca struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is Key Server. The type is bool.
    IsKeyServer interface{}

    // Session Status [Secured/Not Secured]. The type is interface{} with range:
    // 0..4294967295.
    Status interface{}

    // Number of Live Peers. The type is interface{} with range: 0..4294967295.
    NumLivePeers interface{}

    // Is First CA. The type is bool.
    FirstCa interface{}

    // Peer SCI(MAC). The type is string.
    PeerSci interface{}

    // Number of Live Peers responded. The type is interface{} with range:
    // 0..4294967295.
    NumLivePeersResponded interface{}

    // CKN. The type is string.
    Ckn interface{}

    // Member Identifier. The type is string.
    MyMi interface{}

    // Message Number. The type is interface{} with range: 0..4294967295.
    MyMn interface{}

    // authenticator. The type is bool.
    Authenticator interface{}

    // Status Description. The type is string.
    StatusDescription interface{}

    // CA Authentication Mode :PRIMARY-PSK/FALLBACK-PSK/EAP. The type is string.
    AuthenticationMode interface{}

    // Key Chain name. The type is string.
    KeyChain interface{}

    // Peers Status.
    PeersStatus Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus

    // Live Peer List. The type is slice of
    // Macsec_Mka_Interfaces_Interface_Session_Ca_LivePeer.
    LivePeer []Macsec_Mka_Interfaces_Interface_Session_Ca_LivePeer

    // Potential Peer List. The type is slice of
    // Macsec_Mka_Interfaces_Interface_Session_Ca_PotentialPeer.
    PotentialPeer []Macsec_Mka_Interfaces_Interface_Session_Ca_PotentialPeer

    // Dormant Peer List. The type is slice of
    // Macsec_Mka_Interfaces_Interface_Session_Ca_DormantPeer.
    DormantPeer []Macsec_Mka_Interfaces_Interface_Session_Ca_DormantPeer
}

func (ca *Macsec_Mka_Interfaces_Interface_Session_Ca) GetFilter() yfilter.YFilter { return ca.YFilter }

func (ca *Macsec_Mka_Interfaces_Interface_Session_Ca) SetFilter(yf yfilter.YFilter) { ca.YFilter = yf }

func (ca *Macsec_Mka_Interfaces_Interface_Session_Ca) GetGoName(yname string) string {
    if yname == "is-key-server" { return "IsKeyServer" }
    if yname == "status" { return "Status" }
    if yname == "num-live-peers" { return "NumLivePeers" }
    if yname == "first-ca" { return "FirstCa" }
    if yname == "peer-sci" { return "PeerSci" }
    if yname == "num-live-peers-responded" { return "NumLivePeersResponded" }
    if yname == "ckn" { return "Ckn" }
    if yname == "my-mi" { return "MyMi" }
    if yname == "my-mn" { return "MyMn" }
    if yname == "authenticator" { return "Authenticator" }
    if yname == "status-description" { return "StatusDescription" }
    if yname == "authentication-mode" { return "AuthenticationMode" }
    if yname == "key-chain" { return "KeyChain" }
    if yname == "peers-status" { return "PeersStatus" }
    if yname == "live-peer" { return "LivePeer" }
    if yname == "potential-peer" { return "PotentialPeer" }
    if yname == "dormant-peer" { return "DormantPeer" }
    return ""
}

func (ca *Macsec_Mka_Interfaces_Interface_Session_Ca) GetSegmentPath() string {
    return "ca"
}

func (ca *Macsec_Mka_Interfaces_Interface_Session_Ca) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "peers-status" {
        return &ca.PeersStatus
    }
    if childYangName == "live-peer" {
        for _, c := range ca.LivePeer {
            if ca.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Macsec_Mka_Interfaces_Interface_Session_Ca_LivePeer{}
        ca.LivePeer = append(ca.LivePeer, child)
        return &ca.LivePeer[len(ca.LivePeer)-1]
    }
    if childYangName == "potential-peer" {
        for _, c := range ca.PotentialPeer {
            if ca.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Macsec_Mka_Interfaces_Interface_Session_Ca_PotentialPeer{}
        ca.PotentialPeer = append(ca.PotentialPeer, child)
        return &ca.PotentialPeer[len(ca.PotentialPeer)-1]
    }
    if childYangName == "dormant-peer" {
        for _, c := range ca.DormantPeer {
            if ca.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Macsec_Mka_Interfaces_Interface_Session_Ca_DormantPeer{}
        ca.DormantPeer = append(ca.DormantPeer, child)
        return &ca.DormantPeer[len(ca.DormantPeer)-1]
    }
    return nil
}

func (ca *Macsec_Mka_Interfaces_Interface_Session_Ca) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["peers-status"] = &ca.PeersStatus
    for i := range ca.LivePeer {
        children[ca.LivePeer[i].GetSegmentPath()] = &ca.LivePeer[i]
    }
    for i := range ca.PotentialPeer {
        children[ca.PotentialPeer[i].GetSegmentPath()] = &ca.PotentialPeer[i]
    }
    for i := range ca.DormantPeer {
        children[ca.DormantPeer[i].GetSegmentPath()] = &ca.DormantPeer[i]
    }
    return children
}

func (ca *Macsec_Mka_Interfaces_Interface_Session_Ca) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-key-server"] = ca.IsKeyServer
    leafs["status"] = ca.Status
    leafs["num-live-peers"] = ca.NumLivePeers
    leafs["first-ca"] = ca.FirstCa
    leafs["peer-sci"] = ca.PeerSci
    leafs["num-live-peers-responded"] = ca.NumLivePeersResponded
    leafs["ckn"] = ca.Ckn
    leafs["my-mi"] = ca.MyMi
    leafs["my-mn"] = ca.MyMn
    leafs["authenticator"] = ca.Authenticator
    leafs["status-description"] = ca.StatusDescription
    leafs["authentication-mode"] = ca.AuthenticationMode
    leafs["key-chain"] = ca.KeyChain
    return leafs
}

func (ca *Macsec_Mka_Interfaces_Interface_Session_Ca) GetBundleName() string { return "cisco_ios_xr" }

func (ca *Macsec_Mka_Interfaces_Interface_Session_Ca) GetYangName() string { return "ca" }

func (ca *Macsec_Mka_Interfaces_Interface_Session_Ca) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ca *Macsec_Mka_Interfaces_Interface_Session_Ca) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ca *Macsec_Mka_Interfaces_Interface_Session_Ca) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ca *Macsec_Mka_Interfaces_Interface_Session_Ca) SetParent(parent types.Entity) { ca.parent = parent }

func (ca *Macsec_Mka_Interfaces_Interface_Session_Ca) GetParent() types.Entity { return ca.parent }

func (ca *Macsec_Mka_Interfaces_Interface_Session_Ca) GetParentYangName() string { return "session" }

// Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus
// Peers Status
type Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Tx MKPDU Timestamp. The type is string with length: 0..128.
    TxMkpduTimestamp interface{}

    // Peer Count. The type is interface{} with range: 0..4294967295.
    PeerCount interface{}

    // Peer List. The type is slice of
    // Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus_Peer.
    Peer []Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus_Peer
}

func (peersStatus *Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus) GetFilter() yfilter.YFilter { return peersStatus.YFilter }

func (peersStatus *Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus) SetFilter(yf yfilter.YFilter) { peersStatus.YFilter = yf }

func (peersStatus *Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus) GetGoName(yname string) string {
    if yname == "tx-mkpdu-timestamp" { return "TxMkpduTimestamp" }
    if yname == "peer-count" { return "PeerCount" }
    if yname == "peer" { return "Peer" }
    return ""
}

func (peersStatus *Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus) GetSegmentPath() string {
    return "peers-status"
}

func (peersStatus *Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "peer" {
        for _, c := range peersStatus.Peer {
            if peersStatus.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus_Peer{}
        peersStatus.Peer = append(peersStatus.Peer, child)
        return &peersStatus.Peer[len(peersStatus.Peer)-1]
    }
    return nil
}

func (peersStatus *Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range peersStatus.Peer {
        children[peersStatus.Peer[i].GetSegmentPath()] = &peersStatus.Peer[i]
    }
    return children
}

func (peersStatus *Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tx-mkpdu-timestamp"] = peersStatus.TxMkpduTimestamp
    leafs["peer-count"] = peersStatus.PeerCount
    return leafs
}

func (peersStatus *Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus) GetBundleName() string { return "cisco_ios_xr" }

func (peersStatus *Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus) GetYangName() string { return "peers-status" }

func (peersStatus *Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peersStatus *Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peersStatus *Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peersStatus *Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus) SetParent(parent types.Entity) { peersStatus.parent = parent }

func (peersStatus *Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus) GetParent() types.Entity { return peersStatus.parent }

func (peersStatus *Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus) GetParentYangName() string { return "ca" }

// Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus_Peer
// Peer List
type Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus_Peer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Rx SCI. The type is string with length: 0..17.
    Sci interface{}

    // Peer Status Data.
    PeerData Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus_Peer_PeerData
}

func (peer *Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus_Peer) GetFilter() yfilter.YFilter { return peer.YFilter }

func (peer *Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus_Peer) SetFilter(yf yfilter.YFilter) { peer.YFilter = yf }

func (peer *Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus_Peer) GetGoName(yname string) string {
    if yname == "sci" { return "Sci" }
    if yname == "peer-data" { return "PeerData" }
    return ""
}

func (peer *Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus_Peer) GetSegmentPath() string {
    return "peer"
}

func (peer *Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus_Peer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "peer-data" {
        return &peer.PeerData
    }
    return nil
}

func (peer *Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus_Peer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["peer-data"] = &peer.PeerData
    return children
}

func (peer *Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus_Peer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sci"] = peer.Sci
    return leafs
}

func (peer *Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus_Peer) GetBundleName() string { return "cisco_ios_xr" }

func (peer *Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus_Peer) GetYangName() string { return "peer" }

func (peer *Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus_Peer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peer *Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus_Peer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peer *Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus_Peer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peer *Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus_Peer) SetParent(parent types.Entity) { peer.parent = parent }

func (peer *Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus_Peer) GetParent() types.Entity { return peer.parent }

func (peer *Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus_Peer) GetParentYangName() string { return "peers-status" }

// Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus_Peer_PeerData
// Peer Status Data
type Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus_Peer_PeerData struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Member ID. The type is string with length: 0..25.
    Mi interface{}

    // ICV Status. The type is string with length: 0..10.
    IcvStatus interface{}

    // ICV Check Timestamp. The type is string with length: 0..128.
    IcvCheckTimestamp interface{}
}

func (peerData *Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus_Peer_PeerData) GetFilter() yfilter.YFilter { return peerData.YFilter }

func (peerData *Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus_Peer_PeerData) SetFilter(yf yfilter.YFilter) { peerData.YFilter = yf }

func (peerData *Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus_Peer_PeerData) GetGoName(yname string) string {
    if yname == "mi" { return "Mi" }
    if yname == "icv-status" { return "IcvStatus" }
    if yname == "icv-check-timestamp" { return "IcvCheckTimestamp" }
    return ""
}

func (peerData *Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus_Peer_PeerData) GetSegmentPath() string {
    return "peer-data"
}

func (peerData *Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus_Peer_PeerData) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (peerData *Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus_Peer_PeerData) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (peerData *Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus_Peer_PeerData) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mi"] = peerData.Mi
    leafs["icv-status"] = peerData.IcvStatus
    leafs["icv-check-timestamp"] = peerData.IcvCheckTimestamp
    return leafs
}

func (peerData *Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus_Peer_PeerData) GetBundleName() string { return "cisco_ios_xr" }

func (peerData *Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus_Peer_PeerData) GetYangName() string { return "peer-data" }

func (peerData *Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus_Peer_PeerData) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peerData *Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus_Peer_PeerData) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peerData *Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus_Peer_PeerData) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peerData *Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus_Peer_PeerData) SetParent(parent types.Entity) { peerData.parent = parent }

func (peerData *Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus_Peer_PeerData) GetParent() types.Entity { return peerData.parent }

func (peerData *Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus_Peer_PeerData) GetParentYangName() string { return "peer" }

// Macsec_Mka_Interfaces_Interface_Session_Ca_LivePeer
// Live Peer List
type Macsec_Mka_Interfaces_Interface_Session_Ca_LivePeer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Member ID. The type is string.
    Mi interface{}

    // Rx SCI. The type is string.
    Sci interface{}

    // Message Number. The type is interface{} with range: 0..4294967295.
    Mn interface{}

    // KS Priority. The type is interface{} with range: 0..4294967295.
    Priority interface{}

    // Peer SSCI. The type is interface{} with range: 0..4294967295.
    Ssci interface{}
}

func (livePeer *Macsec_Mka_Interfaces_Interface_Session_Ca_LivePeer) GetFilter() yfilter.YFilter { return livePeer.YFilter }

func (livePeer *Macsec_Mka_Interfaces_Interface_Session_Ca_LivePeer) SetFilter(yf yfilter.YFilter) { livePeer.YFilter = yf }

func (livePeer *Macsec_Mka_Interfaces_Interface_Session_Ca_LivePeer) GetGoName(yname string) string {
    if yname == "mi" { return "Mi" }
    if yname == "sci" { return "Sci" }
    if yname == "mn" { return "Mn" }
    if yname == "priority" { return "Priority" }
    if yname == "ssci" { return "Ssci" }
    return ""
}

func (livePeer *Macsec_Mka_Interfaces_Interface_Session_Ca_LivePeer) GetSegmentPath() string {
    return "live-peer"
}

func (livePeer *Macsec_Mka_Interfaces_Interface_Session_Ca_LivePeer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (livePeer *Macsec_Mka_Interfaces_Interface_Session_Ca_LivePeer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (livePeer *Macsec_Mka_Interfaces_Interface_Session_Ca_LivePeer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mi"] = livePeer.Mi
    leafs["sci"] = livePeer.Sci
    leafs["mn"] = livePeer.Mn
    leafs["priority"] = livePeer.Priority
    leafs["ssci"] = livePeer.Ssci
    return leafs
}

func (livePeer *Macsec_Mka_Interfaces_Interface_Session_Ca_LivePeer) GetBundleName() string { return "cisco_ios_xr" }

func (livePeer *Macsec_Mka_Interfaces_Interface_Session_Ca_LivePeer) GetYangName() string { return "live-peer" }

func (livePeer *Macsec_Mka_Interfaces_Interface_Session_Ca_LivePeer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (livePeer *Macsec_Mka_Interfaces_Interface_Session_Ca_LivePeer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (livePeer *Macsec_Mka_Interfaces_Interface_Session_Ca_LivePeer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (livePeer *Macsec_Mka_Interfaces_Interface_Session_Ca_LivePeer) SetParent(parent types.Entity) { livePeer.parent = parent }

func (livePeer *Macsec_Mka_Interfaces_Interface_Session_Ca_LivePeer) GetParent() types.Entity { return livePeer.parent }

func (livePeer *Macsec_Mka_Interfaces_Interface_Session_Ca_LivePeer) GetParentYangName() string { return "ca" }

// Macsec_Mka_Interfaces_Interface_Session_Ca_PotentialPeer
// Potential Peer List
type Macsec_Mka_Interfaces_Interface_Session_Ca_PotentialPeer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Member ID. The type is string.
    Mi interface{}

    // Rx SCI. The type is string.
    Sci interface{}

    // Message Number. The type is interface{} with range: 0..4294967295.
    Mn interface{}

    // KS Priority. The type is interface{} with range: 0..4294967295.
    Priority interface{}

    // Peer SSCI. The type is interface{} with range: 0..4294967295.
    Ssci interface{}
}

func (potentialPeer *Macsec_Mka_Interfaces_Interface_Session_Ca_PotentialPeer) GetFilter() yfilter.YFilter { return potentialPeer.YFilter }

func (potentialPeer *Macsec_Mka_Interfaces_Interface_Session_Ca_PotentialPeer) SetFilter(yf yfilter.YFilter) { potentialPeer.YFilter = yf }

func (potentialPeer *Macsec_Mka_Interfaces_Interface_Session_Ca_PotentialPeer) GetGoName(yname string) string {
    if yname == "mi" { return "Mi" }
    if yname == "sci" { return "Sci" }
    if yname == "mn" { return "Mn" }
    if yname == "priority" { return "Priority" }
    if yname == "ssci" { return "Ssci" }
    return ""
}

func (potentialPeer *Macsec_Mka_Interfaces_Interface_Session_Ca_PotentialPeer) GetSegmentPath() string {
    return "potential-peer"
}

func (potentialPeer *Macsec_Mka_Interfaces_Interface_Session_Ca_PotentialPeer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (potentialPeer *Macsec_Mka_Interfaces_Interface_Session_Ca_PotentialPeer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (potentialPeer *Macsec_Mka_Interfaces_Interface_Session_Ca_PotentialPeer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mi"] = potentialPeer.Mi
    leafs["sci"] = potentialPeer.Sci
    leafs["mn"] = potentialPeer.Mn
    leafs["priority"] = potentialPeer.Priority
    leafs["ssci"] = potentialPeer.Ssci
    return leafs
}

func (potentialPeer *Macsec_Mka_Interfaces_Interface_Session_Ca_PotentialPeer) GetBundleName() string { return "cisco_ios_xr" }

func (potentialPeer *Macsec_Mka_Interfaces_Interface_Session_Ca_PotentialPeer) GetYangName() string { return "potential-peer" }

func (potentialPeer *Macsec_Mka_Interfaces_Interface_Session_Ca_PotentialPeer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (potentialPeer *Macsec_Mka_Interfaces_Interface_Session_Ca_PotentialPeer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (potentialPeer *Macsec_Mka_Interfaces_Interface_Session_Ca_PotentialPeer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (potentialPeer *Macsec_Mka_Interfaces_Interface_Session_Ca_PotentialPeer) SetParent(parent types.Entity) { potentialPeer.parent = parent }

func (potentialPeer *Macsec_Mka_Interfaces_Interface_Session_Ca_PotentialPeer) GetParent() types.Entity { return potentialPeer.parent }

func (potentialPeer *Macsec_Mka_Interfaces_Interface_Session_Ca_PotentialPeer) GetParentYangName() string { return "ca" }

// Macsec_Mka_Interfaces_Interface_Session_Ca_DormantPeer
// Dormant Peer List
type Macsec_Mka_Interfaces_Interface_Session_Ca_DormantPeer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Member ID. The type is string.
    Mi interface{}

    // Rx SCI. The type is string.
    Sci interface{}

    // Message Number. The type is interface{} with range: 0..4294967295.
    Mn interface{}

    // KS Priority. The type is interface{} with range: 0..4294967295.
    Priority interface{}

    // Peer SSCI. The type is interface{} with range: 0..4294967295.
    Ssci interface{}
}

func (dormantPeer *Macsec_Mka_Interfaces_Interface_Session_Ca_DormantPeer) GetFilter() yfilter.YFilter { return dormantPeer.YFilter }

func (dormantPeer *Macsec_Mka_Interfaces_Interface_Session_Ca_DormantPeer) SetFilter(yf yfilter.YFilter) { dormantPeer.YFilter = yf }

func (dormantPeer *Macsec_Mka_Interfaces_Interface_Session_Ca_DormantPeer) GetGoName(yname string) string {
    if yname == "mi" { return "Mi" }
    if yname == "sci" { return "Sci" }
    if yname == "mn" { return "Mn" }
    if yname == "priority" { return "Priority" }
    if yname == "ssci" { return "Ssci" }
    return ""
}

func (dormantPeer *Macsec_Mka_Interfaces_Interface_Session_Ca_DormantPeer) GetSegmentPath() string {
    return "dormant-peer"
}

func (dormantPeer *Macsec_Mka_Interfaces_Interface_Session_Ca_DormantPeer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dormantPeer *Macsec_Mka_Interfaces_Interface_Session_Ca_DormantPeer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dormantPeer *Macsec_Mka_Interfaces_Interface_Session_Ca_DormantPeer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mi"] = dormantPeer.Mi
    leafs["sci"] = dormantPeer.Sci
    leafs["mn"] = dormantPeer.Mn
    leafs["priority"] = dormantPeer.Priority
    leafs["ssci"] = dormantPeer.Ssci
    return leafs
}

func (dormantPeer *Macsec_Mka_Interfaces_Interface_Session_Ca_DormantPeer) GetBundleName() string { return "cisco_ios_xr" }

func (dormantPeer *Macsec_Mka_Interfaces_Interface_Session_Ca_DormantPeer) GetYangName() string { return "dormant-peer" }

func (dormantPeer *Macsec_Mka_Interfaces_Interface_Session_Ca_DormantPeer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dormantPeer *Macsec_Mka_Interfaces_Interface_Session_Ca_DormantPeer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dormantPeer *Macsec_Mka_Interfaces_Interface_Session_Ca_DormantPeer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dormantPeer *Macsec_Mka_Interfaces_Interface_Session_Ca_DormantPeer) SetParent(parent types.Entity) { dormantPeer.parent = parent }

func (dormantPeer *Macsec_Mka_Interfaces_Interface_Session_Ca_DormantPeer) GetParent() types.Entity { return dormantPeer.parent }

func (dormantPeer *Macsec_Mka_Interfaces_Interface_Session_Ca_DormantPeer) GetParentYangName() string { return "ca" }

