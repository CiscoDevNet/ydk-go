// This module contains a collection of YANG definitions
// for monitoring of Cisco Trustsec operational information on
// Role based permissions, IP-SGT bindings and SXP connections.
// Copyright (c) 2016-2017 by Cisco Systems, Inc.
// All rights reserved.
package trustsec_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package trustsec_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-trustsec-oper trustsec-state}", reflect.TypeOf(TrustsecState{}))
    ydk.RegisterEntity("Cisco-IOS-XE-trustsec-oper:trustsec-state", reflect.TypeOf(TrustsecState{}))
}

// CtsOdmBindingSource represents Binding Source enumeration
type CtsOdmBindingSource string

const (
    // Default Security Group Tag binding value in this device
    // for the given IP-Address
    CtsOdmBindingSource_default_ CtsOdmBindingSource = "default"

    // Security Group Tag binding value in this device for the given
    // IP-Address is learned from a VLAN
    CtsOdmBindingSource_from_vlan CtsOdmBindingSource = "from-vlan"

    // Security Group Tag binding value in this device
    // for the given
    // IP-Address is configure from CLI (Command Line Inteface)
    CtsOdmBindingSource_from_cli CtsOdmBindingSource = "from-cli"

    // Security Group Tag binding value in this device
    // for the given IP-Address is learned from a L3 (Layer 3) interface
    CtsOdmBindingSource_from_l3if CtsOdmBindingSource = "from-l3if"

    // Security Group Tag binding value in this device
    // for the given IP-Address is learned via SXP
    // binding exchange protocol
    CtsOdmBindingSource_from_cfp CtsOdmBindingSource = "from-cfp"

    // Security Group Tag binding value in this
    // device for the given
    // IP-Address is learned via IP-ARP protocol
    CtsOdmBindingSource_from_ip_arp CtsOdmBindingSource = "from-ip-arp"

    // Security Group Tag binding value in this device
    // for the given IP-Address is learned locally
    CtsOdmBindingSource_from_local CtsOdmBindingSource = "from-local"

    // Security Group Tag binding value in this device
    // for the given IP-Address is learned via Security Group Tag
    // caching from datapath.
    CtsOdmBindingSource_from_sgt_caching CtsOdmBindingSource = "from-sgt-caching"

    // Security Group Tag binding value in this device
    // for the given IP-Address is configured from CLI-high priority.
    CtsOdmBindingSource_from_cli_hi CtsOdmBindingSource = "from-cli-hi"
)

// SxpConMode represents SXP Connection mode
type SxpConMode string

const (
    // SXP Connection mode is Invalid
    SxpConMode_con_mode_invalid SxpConMode = "con-mode-invalid"

    // SXP Connection mode is Speaker
    SxpConMode_con_mode_speaker SxpConMode = "con-mode-speaker"

    // SXP Connection mode is Listener
    SxpConMode_con_mode_listener SxpConMode = "con-mode-listener"

    // SXP Connection mode is Both (Speaker and Listener)
    SxpConMode_con_mode_both SxpConMode = "con-mode-both"
)

// SxpConState represents SXP Connection state
type SxpConState string

const (
    // SXP Connection state is OFF
    SxpConState_state_off SxpConState = "state-off"

    // SXP Connection state is Pending-On
    SxpConState_state_pending_on SxpConState = "state-pending-on"

    // SXP Connection state is ON
    SxpConState_state_on SxpConState = "state-on"

    // SXP Connection state is Delete-Hold-Down
    SxpConState_state_delete_hold_down SxpConState = "state-delete-hold-down"

    // SXP Connection state is Not-Applicable
    SxpConState_state_not_applicable SxpConState = "state-not-applicable"
)

// TrustsecState
// This is top level container for Cisco Trusted Security
// solution operational data.
// It can have Security Group Tag binding information for
// the given IP-Address, Role based permissions between a
// Source Security Group Tag and a Destination Security
// Group, SXP Connection information for a given peer
// IP-Address in this device
type TrustsecState struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Security Group Tag value corresponding to an IP-Address  in the given VRF
    // instance in this device.
    CtsRolebasedSgtmaps TrustsecState_CtsRolebasedSgtmaps

    // Role based permissions between a Source Security Group and a Destination
    // Security Group are configured by the administrator in the Identity Services
    // Engine or in the Device.
    CtsRolebasedPolicies TrustsecState_CtsRolebasedPolicies

    // Trustsec SXP connection is used between Cisco devices to propagate Security
    // Group Tags from one device to  another device. One of the device will be in
    // Speaker  mode or Listener mode or both the devices can be in both the
    // connection modes.
    CtsSxpConnections TrustsecState_CtsSxpConnections
}

func (trustsecState *TrustsecState) GetFilter() yfilter.YFilter { return trustsecState.YFilter }

func (trustsecState *TrustsecState) SetFilter(yf yfilter.YFilter) { trustsecState.YFilter = yf }

func (trustsecState *TrustsecState) GetGoName(yname string) string {
    if yname == "cts-rolebased-sgtmaps" { return "CtsRolebasedSgtmaps" }
    if yname == "cts-rolebased-policies" { return "CtsRolebasedPolicies" }
    if yname == "cts-sxp-connections" { return "CtsSxpConnections" }
    return ""
}

func (trustsecState *TrustsecState) GetSegmentPath() string {
    return "Cisco-IOS-XE-trustsec-oper:trustsec-state"
}

func (trustsecState *TrustsecState) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cts-rolebased-sgtmaps" {
        return &trustsecState.CtsRolebasedSgtmaps
    }
    if childYangName == "cts-rolebased-policies" {
        return &trustsecState.CtsRolebasedPolicies
    }
    if childYangName == "cts-sxp-connections" {
        return &trustsecState.CtsSxpConnections
    }
    return nil
}

func (trustsecState *TrustsecState) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cts-rolebased-sgtmaps"] = &trustsecState.CtsRolebasedSgtmaps
    children["cts-rolebased-policies"] = &trustsecState.CtsRolebasedPolicies
    children["cts-sxp-connections"] = &trustsecState.CtsSxpConnections
    return children
}

func (trustsecState *TrustsecState) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (trustsecState *TrustsecState) GetBundleName() string { return "cisco_ios_xe" }

func (trustsecState *TrustsecState) GetYangName() string { return "trustsec-state" }

func (trustsecState *TrustsecState) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (trustsecState *TrustsecState) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (trustsecState *TrustsecState) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (trustsecState *TrustsecState) SetParent(parent types.Entity) { trustsecState.parent = parent }

func (trustsecState *TrustsecState) GetParent() types.Entity { return trustsecState.parent }

func (trustsecState *TrustsecState) GetParentYangName() string { return "Cisco-IOS-XE-trustsec-oper" }

// TrustsecState_CtsRolebasedSgtmaps
// Security Group Tag value corresponding to an IP-Address 
// in the given VRF instance in this device
type TrustsecState_CtsRolebasedSgtmaps struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Security Group Tag is assigned for an IP-Address based on the user
    // permissions and authorization  level as configured by the network
    // administrator in Identity Service Engine server or in the device locally.
    // The type is slice of TrustsecState_CtsRolebasedSgtmaps_CtsRolebasedSgtmap.
    CtsRolebasedSgtmap []TrustsecState_CtsRolebasedSgtmaps_CtsRolebasedSgtmap
}

func (ctsRolebasedSgtmaps *TrustsecState_CtsRolebasedSgtmaps) GetFilter() yfilter.YFilter { return ctsRolebasedSgtmaps.YFilter }

func (ctsRolebasedSgtmaps *TrustsecState_CtsRolebasedSgtmaps) SetFilter(yf yfilter.YFilter) { ctsRolebasedSgtmaps.YFilter = yf }

func (ctsRolebasedSgtmaps *TrustsecState_CtsRolebasedSgtmaps) GetGoName(yname string) string {
    if yname == "cts-rolebased-sgtmap" { return "CtsRolebasedSgtmap" }
    return ""
}

func (ctsRolebasedSgtmaps *TrustsecState_CtsRolebasedSgtmaps) GetSegmentPath() string {
    return "cts-rolebased-sgtmaps"
}

func (ctsRolebasedSgtmaps *TrustsecState_CtsRolebasedSgtmaps) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cts-rolebased-sgtmap" {
        for _, c := range ctsRolebasedSgtmaps.CtsRolebasedSgtmap {
            if ctsRolebasedSgtmaps.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TrustsecState_CtsRolebasedSgtmaps_CtsRolebasedSgtmap{}
        ctsRolebasedSgtmaps.CtsRolebasedSgtmap = append(ctsRolebasedSgtmaps.CtsRolebasedSgtmap, child)
        return &ctsRolebasedSgtmaps.CtsRolebasedSgtmap[len(ctsRolebasedSgtmaps.CtsRolebasedSgtmap)-1]
    }
    return nil
}

func (ctsRolebasedSgtmaps *TrustsecState_CtsRolebasedSgtmaps) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ctsRolebasedSgtmaps.CtsRolebasedSgtmap {
        children[ctsRolebasedSgtmaps.CtsRolebasedSgtmap[i].GetSegmentPath()] = &ctsRolebasedSgtmaps.CtsRolebasedSgtmap[i]
    }
    return children
}

func (ctsRolebasedSgtmaps *TrustsecState_CtsRolebasedSgtmaps) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ctsRolebasedSgtmaps *TrustsecState_CtsRolebasedSgtmaps) GetBundleName() string { return "cisco_ios_xe" }

func (ctsRolebasedSgtmaps *TrustsecState_CtsRolebasedSgtmaps) GetYangName() string { return "cts-rolebased-sgtmaps" }

func (ctsRolebasedSgtmaps *TrustsecState_CtsRolebasedSgtmaps) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ctsRolebasedSgtmaps *TrustsecState_CtsRolebasedSgtmaps) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ctsRolebasedSgtmaps *TrustsecState_CtsRolebasedSgtmaps) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ctsRolebasedSgtmaps *TrustsecState_CtsRolebasedSgtmaps) SetParent(parent types.Entity) { ctsRolebasedSgtmaps.parent = parent }

func (ctsRolebasedSgtmaps *TrustsecState_CtsRolebasedSgtmaps) GetParent() types.Entity { return ctsRolebasedSgtmaps.parent }

func (ctsRolebasedSgtmaps *TrustsecState_CtsRolebasedSgtmaps) GetParentYangName() string { return "trustsec-state" }

// TrustsecState_CtsRolebasedSgtmaps_CtsRolebasedSgtmap
// Security Group Tag is assigned for an IP-Address
// based on the user permissions and authorization 
// level as configured by the network administrator
// in Identity Service Engine server or in the device locally
type TrustsecState_CtsRolebasedSgtmaps_CtsRolebasedSgtmap struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP-Prefix information to find its corresponding
    // Secure Group Tag. Only IPv4 prefix information is supported currently to
    // get the Security Group Tag binding in this device. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    Ip interface{}

    // This attribute is a key. VRF-Name to find the Security Group Tag for the
    // corresponding IP-Address in this VRF instance. Only default VRF is
    // supported currently which is indicated by (empty string). The type is
    // string.
    VrfName interface{}

    // Security Group Tag value corresponding to the given IP-Address. The type is
    // interface{} with range: -2147483648..2147483647.
    Sgt interface{}

    // Source information via which the Security Group Tag binding was learned in
    // this device. The type is CtsOdmBindingSource.
    Source interface{}
}

func (ctsRolebasedSgtmap *TrustsecState_CtsRolebasedSgtmaps_CtsRolebasedSgtmap) GetFilter() yfilter.YFilter { return ctsRolebasedSgtmap.YFilter }

func (ctsRolebasedSgtmap *TrustsecState_CtsRolebasedSgtmaps_CtsRolebasedSgtmap) SetFilter(yf yfilter.YFilter) { ctsRolebasedSgtmap.YFilter = yf }

func (ctsRolebasedSgtmap *TrustsecState_CtsRolebasedSgtmaps_CtsRolebasedSgtmap) GetGoName(yname string) string {
    if yname == "ip" { return "Ip" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "sgt" { return "Sgt" }
    if yname == "source" { return "Source" }
    return ""
}

func (ctsRolebasedSgtmap *TrustsecState_CtsRolebasedSgtmaps_CtsRolebasedSgtmap) GetSegmentPath() string {
    return "cts-rolebased-sgtmap" + "[ip='" + fmt.Sprintf("%v", ctsRolebasedSgtmap.Ip) + "']" + "[vrf-name='" + fmt.Sprintf("%v", ctsRolebasedSgtmap.VrfName) + "']"
}

func (ctsRolebasedSgtmap *TrustsecState_CtsRolebasedSgtmaps_CtsRolebasedSgtmap) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ctsRolebasedSgtmap *TrustsecState_CtsRolebasedSgtmaps_CtsRolebasedSgtmap) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ctsRolebasedSgtmap *TrustsecState_CtsRolebasedSgtmaps_CtsRolebasedSgtmap) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip"] = ctsRolebasedSgtmap.Ip
    leafs["vrf-name"] = ctsRolebasedSgtmap.VrfName
    leafs["sgt"] = ctsRolebasedSgtmap.Sgt
    leafs["source"] = ctsRolebasedSgtmap.Source
    return leafs
}

func (ctsRolebasedSgtmap *TrustsecState_CtsRolebasedSgtmaps_CtsRolebasedSgtmap) GetBundleName() string { return "cisco_ios_xe" }

func (ctsRolebasedSgtmap *TrustsecState_CtsRolebasedSgtmaps_CtsRolebasedSgtmap) GetYangName() string { return "cts-rolebased-sgtmap" }

func (ctsRolebasedSgtmap *TrustsecState_CtsRolebasedSgtmaps_CtsRolebasedSgtmap) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ctsRolebasedSgtmap *TrustsecState_CtsRolebasedSgtmaps_CtsRolebasedSgtmap) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ctsRolebasedSgtmap *TrustsecState_CtsRolebasedSgtmaps_CtsRolebasedSgtmap) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ctsRolebasedSgtmap *TrustsecState_CtsRolebasedSgtmaps_CtsRolebasedSgtmap) SetParent(parent types.Entity) { ctsRolebasedSgtmap.parent = parent }

func (ctsRolebasedSgtmap *TrustsecState_CtsRolebasedSgtmaps_CtsRolebasedSgtmap) GetParent() types.Entity { return ctsRolebasedSgtmap.parent }

func (ctsRolebasedSgtmap *TrustsecState_CtsRolebasedSgtmaps_CtsRolebasedSgtmap) GetParentYangName() string { return "cts-rolebased-sgtmaps" }

// TrustsecState_CtsRolebasedPolicies
// Role based permissions between a Source Security Group
// and a Destination Security Group are configured by the
// administrator in the Identity Services Engine or in the Device
type TrustsecState_CtsRolebasedPolicies struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Role based permissions between a Source Security Group and a Destination
    // Security Group can be retrieved from the device using a Security Group Tag
    // and Destination Group Tag value. The type is slice of
    // TrustsecState_CtsRolebasedPolicies_CtsRolebasedPolicy.
    CtsRolebasedPolicy []TrustsecState_CtsRolebasedPolicies_CtsRolebasedPolicy
}

func (ctsRolebasedPolicies *TrustsecState_CtsRolebasedPolicies) GetFilter() yfilter.YFilter { return ctsRolebasedPolicies.YFilter }

func (ctsRolebasedPolicies *TrustsecState_CtsRolebasedPolicies) SetFilter(yf yfilter.YFilter) { ctsRolebasedPolicies.YFilter = yf }

func (ctsRolebasedPolicies *TrustsecState_CtsRolebasedPolicies) GetGoName(yname string) string {
    if yname == "cts-rolebased-policy" { return "CtsRolebasedPolicy" }
    return ""
}

func (ctsRolebasedPolicies *TrustsecState_CtsRolebasedPolicies) GetSegmentPath() string {
    return "cts-rolebased-policies"
}

func (ctsRolebasedPolicies *TrustsecState_CtsRolebasedPolicies) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cts-rolebased-policy" {
        for _, c := range ctsRolebasedPolicies.CtsRolebasedPolicy {
            if ctsRolebasedPolicies.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TrustsecState_CtsRolebasedPolicies_CtsRolebasedPolicy{}
        ctsRolebasedPolicies.CtsRolebasedPolicy = append(ctsRolebasedPolicies.CtsRolebasedPolicy, child)
        return &ctsRolebasedPolicies.CtsRolebasedPolicy[len(ctsRolebasedPolicies.CtsRolebasedPolicy)-1]
    }
    return nil
}

func (ctsRolebasedPolicies *TrustsecState_CtsRolebasedPolicies) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ctsRolebasedPolicies.CtsRolebasedPolicy {
        children[ctsRolebasedPolicies.CtsRolebasedPolicy[i].GetSegmentPath()] = &ctsRolebasedPolicies.CtsRolebasedPolicy[i]
    }
    return children
}

func (ctsRolebasedPolicies *TrustsecState_CtsRolebasedPolicies) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ctsRolebasedPolicies *TrustsecState_CtsRolebasedPolicies) GetBundleName() string { return "cisco_ios_xe" }

func (ctsRolebasedPolicies *TrustsecState_CtsRolebasedPolicies) GetYangName() string { return "cts-rolebased-policies" }

func (ctsRolebasedPolicies *TrustsecState_CtsRolebasedPolicies) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ctsRolebasedPolicies *TrustsecState_CtsRolebasedPolicies) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ctsRolebasedPolicies *TrustsecState_CtsRolebasedPolicies) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ctsRolebasedPolicies *TrustsecState_CtsRolebasedPolicies) SetParent(parent types.Entity) { ctsRolebasedPolicies.parent = parent }

func (ctsRolebasedPolicies *TrustsecState_CtsRolebasedPolicies) GetParent() types.Entity { return ctsRolebasedPolicies.parent }

func (ctsRolebasedPolicies *TrustsecState_CtsRolebasedPolicies) GetParentYangName() string { return "trustsec-state" }

// TrustsecState_CtsRolebasedPolicies_CtsRolebasedPolicy
// Role based permissions between a Source Security Group
// and a Destination Security Group can be retrieved from
// the device using a Security Group Tag and Destination
// Group Tag value
type TrustsecState_CtsRolebasedPolicies_CtsRolebasedPolicy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Source Security Group Tag number. This value must
    // be in the inclusive range of -1 to 65519. The type is interface{} with
    // range: -2147483648..2147483647.
    SrcSgt interface{}

    // This attribute is a key. Destination Security Tag number. This value must
    // be in the inclusive range of -1 to 65519. The type is interface{} with
    // range: -2147483648..2147483647.
    DstSgt interface{}

    // List of Security Group Access Control List names separated by
    // semi-colon(;). The type is string.
    SgaclName interface{}

    // Number of Security Group Access Control Lists that are currently applied
    // between the Source Security Group and the Destination Security Group. This
    // should match the number of Security Group Access Control List names in
    // sgacl-name. The type is interface{} with range: 0..4294967295.
    NumOfSgacl interface{}

    // Indicates the monitor mode status between the Source Security Group and
    // Destination Security Group is currently enabled or disabled. This will be
    // TRUE if monitor mode is enabled and FALSE if it is disabled. The type is
    // bool.
    MonitorMode interface{}

    // Duration of the Role based permissions that are applied between a Source
    // Security Group and a Destination Security Group. The duration is
    // represented in seconds. The type is interface{} with range:
    // 0..18446744073709551615.
    PolicyLifeTime interface{}

    // Indicates the time when the Role based permissions between a Source
    // Security Group and a Destination Security Group was modified or updated
    // last. The value will be represented as date and time  corresponding to the
    // local time zone of the Identify Services Engine when the Role based 
    // permissions was modified or updated last. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    LastUpdatedTime interface{}

    // Total number of packets that have been denied by the Role based permissions
    // between a Source Security Group and a Destination Security Group. This
    // corresponds to total packets denied in both hardware and software
    // forwarding paths of the device. The type is interface{} with range:
    // 0..18446744073709551615.
    TotalDenyCount interface{}

    // Total number of packets that have been permitted by the Role based
    // permissions between a Source Security Group and a Destination Security
    // Group. This corresponds to total packets allowed in both hardware and
    // software forwarding paths of the device. The type is interface{} with
    // range: 0..18446744073709551615.
    TotalPermitCount interface{}

    // Number of packets that have been denied in the  software forwarding path of
    // the device by the Role based permissions between a Source Security Group
    // and a Destination Security Group. The type is interface{} with range:
    // 0..18446744073709551615.
    SoftwareDenyCount interface{}

    // Number of packets that have been permitted in the software forwarding path
    // of the device by the Role based permissions between a Source Security Group
    // and a Destination Security Group. The type is interface{} with range:
    // 0..18446744073709551615.
    SoftwarePermitCount interface{}

    // Number of packets that have been denied in the hardware forwarding path of
    // the device by the Role based permissions between a Source Security Group
    // and a Destination Security Group. The type is interface{} with range:
    // 0..18446744073709551615.
    HardwareDenyCount interface{}

    // Number of packets that have been permitted in the hardware forwarding path
    // of the device by the Role based permissions between a Source Security Group
    // and a Destination Security Group. The type is interface{} with range:
    // 0..18446744073709551615.
    HardwarePermitCount interface{}

    // Number of packets that have been monitored in the software forwarding path
    // of the device by the Role based permissions between a Source Security Group
    // and a Destination Security Group. The type is interface{} with range:
    // 0..18446744073709551615.
    SoftwareMonitorCount interface{}

    // Number of packets that have been monitored in the hardware forwarding path
    // of the device by the Role based permissions between a Source Security Group
    // and a Destination Security Group. The type is interface{} with range:
    // 0..18446744073709551615.
    HardwareMonitorCount interface{}
}

func (ctsRolebasedPolicy *TrustsecState_CtsRolebasedPolicies_CtsRolebasedPolicy) GetFilter() yfilter.YFilter { return ctsRolebasedPolicy.YFilter }

func (ctsRolebasedPolicy *TrustsecState_CtsRolebasedPolicies_CtsRolebasedPolicy) SetFilter(yf yfilter.YFilter) { ctsRolebasedPolicy.YFilter = yf }

func (ctsRolebasedPolicy *TrustsecState_CtsRolebasedPolicies_CtsRolebasedPolicy) GetGoName(yname string) string {
    if yname == "src-sgt" { return "SrcSgt" }
    if yname == "dst-sgt" { return "DstSgt" }
    if yname == "sgacl-name" { return "SgaclName" }
    if yname == "num-of-sgacl" { return "NumOfSgacl" }
    if yname == "monitor-mode" { return "MonitorMode" }
    if yname == "policy-life-time" { return "PolicyLifeTime" }
    if yname == "last-updated-time" { return "LastUpdatedTime" }
    if yname == "total-deny-count" { return "TotalDenyCount" }
    if yname == "total-permit-count" { return "TotalPermitCount" }
    if yname == "software-deny-count" { return "SoftwareDenyCount" }
    if yname == "software-permit-count" { return "SoftwarePermitCount" }
    if yname == "hardware-deny-count" { return "HardwareDenyCount" }
    if yname == "hardware-permit-count" { return "HardwarePermitCount" }
    if yname == "software-monitor-count" { return "SoftwareMonitorCount" }
    if yname == "hardware-monitor-count" { return "HardwareMonitorCount" }
    return ""
}

func (ctsRolebasedPolicy *TrustsecState_CtsRolebasedPolicies_CtsRolebasedPolicy) GetSegmentPath() string {
    return "cts-rolebased-policy" + "[src-sgt='" + fmt.Sprintf("%v", ctsRolebasedPolicy.SrcSgt) + "']" + "[dst-sgt='" + fmt.Sprintf("%v", ctsRolebasedPolicy.DstSgt) + "']"
}

func (ctsRolebasedPolicy *TrustsecState_CtsRolebasedPolicies_CtsRolebasedPolicy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ctsRolebasedPolicy *TrustsecState_CtsRolebasedPolicies_CtsRolebasedPolicy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ctsRolebasedPolicy *TrustsecState_CtsRolebasedPolicies_CtsRolebasedPolicy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["src-sgt"] = ctsRolebasedPolicy.SrcSgt
    leafs["dst-sgt"] = ctsRolebasedPolicy.DstSgt
    leafs["sgacl-name"] = ctsRolebasedPolicy.SgaclName
    leafs["num-of-sgacl"] = ctsRolebasedPolicy.NumOfSgacl
    leafs["monitor-mode"] = ctsRolebasedPolicy.MonitorMode
    leafs["policy-life-time"] = ctsRolebasedPolicy.PolicyLifeTime
    leafs["last-updated-time"] = ctsRolebasedPolicy.LastUpdatedTime
    leafs["total-deny-count"] = ctsRolebasedPolicy.TotalDenyCount
    leafs["total-permit-count"] = ctsRolebasedPolicy.TotalPermitCount
    leafs["software-deny-count"] = ctsRolebasedPolicy.SoftwareDenyCount
    leafs["software-permit-count"] = ctsRolebasedPolicy.SoftwarePermitCount
    leafs["hardware-deny-count"] = ctsRolebasedPolicy.HardwareDenyCount
    leafs["hardware-permit-count"] = ctsRolebasedPolicy.HardwarePermitCount
    leafs["software-monitor-count"] = ctsRolebasedPolicy.SoftwareMonitorCount
    leafs["hardware-monitor-count"] = ctsRolebasedPolicy.HardwareMonitorCount
    return leafs
}

func (ctsRolebasedPolicy *TrustsecState_CtsRolebasedPolicies_CtsRolebasedPolicy) GetBundleName() string { return "cisco_ios_xe" }

func (ctsRolebasedPolicy *TrustsecState_CtsRolebasedPolicies_CtsRolebasedPolicy) GetYangName() string { return "cts-rolebased-policy" }

func (ctsRolebasedPolicy *TrustsecState_CtsRolebasedPolicies_CtsRolebasedPolicy) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ctsRolebasedPolicy *TrustsecState_CtsRolebasedPolicies_CtsRolebasedPolicy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ctsRolebasedPolicy *TrustsecState_CtsRolebasedPolicies_CtsRolebasedPolicy) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ctsRolebasedPolicy *TrustsecState_CtsRolebasedPolicies_CtsRolebasedPolicy) SetParent(parent types.Entity) { ctsRolebasedPolicy.parent = parent }

func (ctsRolebasedPolicy *TrustsecState_CtsRolebasedPolicies_CtsRolebasedPolicy) GetParent() types.Entity { return ctsRolebasedPolicy.parent }

func (ctsRolebasedPolicy *TrustsecState_CtsRolebasedPolicies_CtsRolebasedPolicy) GetParentYangName() string { return "cts-rolebased-policies" }

// TrustsecState_CtsSxpConnections
// Trustsec SXP connection is used between Cisco devices
// to propagate Security Group Tags from one device to 
// another device. One of the device will be in Speaker 
// mode or Listener mode or both the devices can be in
// both the connection modes
type TrustsecState_CtsSxpConnections struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Trustsec SXP connection information from a device can be retrieved with the
    // SXP connection peer IP address. Only IPv4 address as Peer IP and default
    // VRF instance in device is supported currently. The type is slice of
    // TrustsecState_CtsSxpConnections_CtsSxpConnection.
    CtsSxpConnection []TrustsecState_CtsSxpConnections_CtsSxpConnection
}

func (ctsSxpConnections *TrustsecState_CtsSxpConnections) GetFilter() yfilter.YFilter { return ctsSxpConnections.YFilter }

func (ctsSxpConnections *TrustsecState_CtsSxpConnections) SetFilter(yf yfilter.YFilter) { ctsSxpConnections.YFilter = yf }

func (ctsSxpConnections *TrustsecState_CtsSxpConnections) GetGoName(yname string) string {
    if yname == "cts-sxp-connection" { return "CtsSxpConnection" }
    return ""
}

func (ctsSxpConnections *TrustsecState_CtsSxpConnections) GetSegmentPath() string {
    return "cts-sxp-connections"
}

func (ctsSxpConnections *TrustsecState_CtsSxpConnections) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cts-sxp-connection" {
        for _, c := range ctsSxpConnections.CtsSxpConnection {
            if ctsSxpConnections.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TrustsecState_CtsSxpConnections_CtsSxpConnection{}
        ctsSxpConnections.CtsSxpConnection = append(ctsSxpConnections.CtsSxpConnection, child)
        return &ctsSxpConnections.CtsSxpConnection[len(ctsSxpConnections.CtsSxpConnection)-1]
    }
    return nil
}

func (ctsSxpConnections *TrustsecState_CtsSxpConnections) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ctsSxpConnections.CtsSxpConnection {
        children[ctsSxpConnections.CtsSxpConnection[i].GetSegmentPath()] = &ctsSxpConnections.CtsSxpConnection[i]
    }
    return children
}

func (ctsSxpConnections *TrustsecState_CtsSxpConnections) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ctsSxpConnections *TrustsecState_CtsSxpConnections) GetBundleName() string { return "cisco_ios_xe" }

func (ctsSxpConnections *TrustsecState_CtsSxpConnections) GetYangName() string { return "cts-sxp-connections" }

func (ctsSxpConnections *TrustsecState_CtsSxpConnections) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ctsSxpConnections *TrustsecState_CtsSxpConnections) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ctsSxpConnections *TrustsecState_CtsSxpConnections) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ctsSxpConnections *TrustsecState_CtsSxpConnections) SetParent(parent types.Entity) { ctsSxpConnections.parent = parent }

func (ctsSxpConnections *TrustsecState_CtsSxpConnections) GetParent() types.Entity { return ctsSxpConnections.parent }

func (ctsSxpConnections *TrustsecState_CtsSxpConnections) GetParentYangName() string { return "trustsec-state" }

// TrustsecState_CtsSxpConnections_CtsSxpConnection
// Trustsec SXP connection information from a device
// can be retrieved with the SXP connection peer IP
// address. Only IPv4 address as Peer IP and default
// VRF instance in device is supported currently
type TrustsecState_CtsSxpConnections_CtsSxpConnection struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP-Address information of the peer of an SXP
    // connection in this device. Only IPv4 address is currently supported. The
    // type is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    PeerIp interface{}

    // This attribute is a key. vrf-name string of the VRF instance in this
    // device, to which the peer of an SXP connection belongs to. Only default VRF
    // is supported currently which is indicated by empty string. The type is
    // string.
    VrfName interface{}

    // Source IP-Address of the SXP connection in this device for the given peer
    // IP-Address. The type is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SourceIp interface{}

    // SXP speaker state information of the SXP connection in this device. This
    // information is valid only if the local mode of the SXP connection in this
    // device is Speaker. The type is SxpConState.
    SpeakerState interface{}

    // Duration of the SXP speaker of the SXP connection in this device. This
    // information is valid only if the local mode of the SXP connection is
    // Speaker. The type is interface{} with range: 0..18446744073709551615.
    SpeakerDuration interface{}

    // SXP listener state information of the SXP  connection in this device. This
    // information is valid only if the local mode of the SXP connection in the
    // device is Listener. The type is SxpConState.
    ListenerState interface{}

    // Duration of the SXP listener of the SXP connection in this device. This
    // information is valid Only if the local mode of the SXP connection is
    // Listener. The type is interface{} with range: 0..18446744073709551615.
    ListenerDuration interface{}

    // SXP connection mode of this device for the SXP connection with the given
    // peer. The type is SxpConMode.
    LocalMode interface{}
}

func (ctsSxpConnection *TrustsecState_CtsSxpConnections_CtsSxpConnection) GetFilter() yfilter.YFilter { return ctsSxpConnection.YFilter }

func (ctsSxpConnection *TrustsecState_CtsSxpConnections_CtsSxpConnection) SetFilter(yf yfilter.YFilter) { ctsSxpConnection.YFilter = yf }

func (ctsSxpConnection *TrustsecState_CtsSxpConnections_CtsSxpConnection) GetGoName(yname string) string {
    if yname == "peer-ip" { return "PeerIp" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "source-ip" { return "SourceIp" }
    if yname == "speaker-state" { return "SpeakerState" }
    if yname == "speaker-duration" { return "SpeakerDuration" }
    if yname == "listener-state" { return "ListenerState" }
    if yname == "listener-duration" { return "ListenerDuration" }
    if yname == "local-mode" { return "LocalMode" }
    return ""
}

func (ctsSxpConnection *TrustsecState_CtsSxpConnections_CtsSxpConnection) GetSegmentPath() string {
    return "cts-sxp-connection" + "[peer-ip='" + fmt.Sprintf("%v", ctsSxpConnection.PeerIp) + "']" + "[vrf-name='" + fmt.Sprintf("%v", ctsSxpConnection.VrfName) + "']"
}

func (ctsSxpConnection *TrustsecState_CtsSxpConnections_CtsSxpConnection) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ctsSxpConnection *TrustsecState_CtsSxpConnections_CtsSxpConnection) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ctsSxpConnection *TrustsecState_CtsSxpConnections_CtsSxpConnection) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["peer-ip"] = ctsSxpConnection.PeerIp
    leafs["vrf-name"] = ctsSxpConnection.VrfName
    leafs["source-ip"] = ctsSxpConnection.SourceIp
    leafs["speaker-state"] = ctsSxpConnection.SpeakerState
    leafs["speaker-duration"] = ctsSxpConnection.SpeakerDuration
    leafs["listener-state"] = ctsSxpConnection.ListenerState
    leafs["listener-duration"] = ctsSxpConnection.ListenerDuration
    leafs["local-mode"] = ctsSxpConnection.LocalMode
    return leafs
}

func (ctsSxpConnection *TrustsecState_CtsSxpConnections_CtsSxpConnection) GetBundleName() string { return "cisco_ios_xe" }

func (ctsSxpConnection *TrustsecState_CtsSxpConnections_CtsSxpConnection) GetYangName() string { return "cts-sxp-connection" }

func (ctsSxpConnection *TrustsecState_CtsSxpConnections_CtsSxpConnection) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ctsSxpConnection *TrustsecState_CtsSxpConnections_CtsSxpConnection) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ctsSxpConnection *TrustsecState_CtsSxpConnections_CtsSxpConnection) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ctsSxpConnection *TrustsecState_CtsSxpConnections_CtsSxpConnection) SetParent(parent types.Entity) { ctsSxpConnection.parent = parent }

func (ctsSxpConnection *TrustsecState_CtsSxpConnections_CtsSxpConnection) GetParent() types.Entity { return ctsSxpConnection.parent }

func (ctsSxpConnection *TrustsecState_CtsSxpConnections_CtsSxpConnection) GetParentYangName() string { return "cts-sxp-connections" }

