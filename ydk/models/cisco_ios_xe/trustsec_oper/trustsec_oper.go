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

// TrustsecState
// This is top level container for Cisco Trusted Security
// solution operational data.
// It can have Security Group Tag binding information for
// the given IP-Address, Role based permissions between a
// Source Security Group Tag and a Destination Security
// Group, SXP Connection information for a given peer
// IP-Address in this device
type TrustsecState struct {
    EntityData types.CommonEntityData
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

func (trustsecState *TrustsecState) GetEntityData() *types.CommonEntityData {
    trustsecState.EntityData.YFilter = trustsecState.YFilter
    trustsecState.EntityData.YangName = "trustsec-state"
    trustsecState.EntityData.BundleName = "cisco_ios_xe"
    trustsecState.EntityData.ParentYangName = "Cisco-IOS-XE-trustsec-oper"
    trustsecState.EntityData.SegmentPath = "Cisco-IOS-XE-trustsec-oper:trustsec-state"
    trustsecState.EntityData.AbsolutePath = trustsecState.EntityData.SegmentPath
    trustsecState.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    trustsecState.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    trustsecState.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    trustsecState.EntityData.Children = types.NewOrderedMap()
    trustsecState.EntityData.Children.Append("cts-rolebased-sgtmaps", types.YChild{"CtsRolebasedSgtmaps", &trustsecState.CtsRolebasedSgtmaps})
    trustsecState.EntityData.Children.Append("cts-rolebased-policies", types.YChild{"CtsRolebasedPolicies", &trustsecState.CtsRolebasedPolicies})
    trustsecState.EntityData.Children.Append("cts-sxp-connections", types.YChild{"CtsSxpConnections", &trustsecState.CtsSxpConnections})
    trustsecState.EntityData.Leafs = types.NewOrderedMap()

    trustsecState.EntityData.YListKeys = []string {}

    return &(trustsecState.EntityData)
}

// TrustsecState_CtsRolebasedSgtmaps
// Security Group Tag value corresponding to an IP-Address 
// in the given VRF instance in this device
type TrustsecState_CtsRolebasedSgtmaps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Security Group Tag is assigned for an IP-Address based on the user
    // permissions and authorization  level as configured by the network
    // administrator in Identity Service Engine server or in the device locally.
    // The type is slice of TrustsecState_CtsRolebasedSgtmaps_CtsRolebasedSgtmap.
    CtsRolebasedSgtmap []*TrustsecState_CtsRolebasedSgtmaps_CtsRolebasedSgtmap
}

func (ctsRolebasedSgtmaps *TrustsecState_CtsRolebasedSgtmaps) GetEntityData() *types.CommonEntityData {
    ctsRolebasedSgtmaps.EntityData.YFilter = ctsRolebasedSgtmaps.YFilter
    ctsRolebasedSgtmaps.EntityData.YangName = "cts-rolebased-sgtmaps"
    ctsRolebasedSgtmaps.EntityData.BundleName = "cisco_ios_xe"
    ctsRolebasedSgtmaps.EntityData.ParentYangName = "trustsec-state"
    ctsRolebasedSgtmaps.EntityData.SegmentPath = "cts-rolebased-sgtmaps"
    ctsRolebasedSgtmaps.EntityData.AbsolutePath = "Cisco-IOS-XE-trustsec-oper:trustsec-state/" + ctsRolebasedSgtmaps.EntityData.SegmentPath
    ctsRolebasedSgtmaps.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ctsRolebasedSgtmaps.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ctsRolebasedSgtmaps.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ctsRolebasedSgtmaps.EntityData.Children = types.NewOrderedMap()
    ctsRolebasedSgtmaps.EntityData.Children.Append("cts-rolebased-sgtmap", types.YChild{"CtsRolebasedSgtmap", nil})
    for i := range ctsRolebasedSgtmaps.CtsRolebasedSgtmap {
        ctsRolebasedSgtmaps.EntityData.Children.Append(types.GetSegmentPath(ctsRolebasedSgtmaps.CtsRolebasedSgtmap[i]), types.YChild{"CtsRolebasedSgtmap", ctsRolebasedSgtmaps.CtsRolebasedSgtmap[i]})
    }
    ctsRolebasedSgtmaps.EntityData.Leafs = types.NewOrderedMap()

    ctsRolebasedSgtmaps.EntityData.YListKeys = []string {}

    return &(ctsRolebasedSgtmaps.EntityData)
}

// TrustsecState_CtsRolebasedSgtmaps_CtsRolebasedSgtmap
// Security Group Tag is assigned for an IP-Address
// based on the user permissions and authorization 
// level as configured by the network administrator
// in Identity Service Engine server or in the device locally
type TrustsecState_CtsRolebasedSgtmaps_CtsRolebasedSgtmap struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. IP-Prefix information to find its corresponding
    // Secure Group Tag. Only IPv4 prefix information is supported currently to
    // get the Security Group Tag binding in this device. The type is one of the
    // following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))'.
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

func (ctsRolebasedSgtmap *TrustsecState_CtsRolebasedSgtmaps_CtsRolebasedSgtmap) GetEntityData() *types.CommonEntityData {
    ctsRolebasedSgtmap.EntityData.YFilter = ctsRolebasedSgtmap.YFilter
    ctsRolebasedSgtmap.EntityData.YangName = "cts-rolebased-sgtmap"
    ctsRolebasedSgtmap.EntityData.BundleName = "cisco_ios_xe"
    ctsRolebasedSgtmap.EntityData.ParentYangName = "cts-rolebased-sgtmaps"
    ctsRolebasedSgtmap.EntityData.SegmentPath = "cts-rolebased-sgtmap" + types.AddKeyToken(ctsRolebasedSgtmap.Ip, "ip") + types.AddKeyToken(ctsRolebasedSgtmap.VrfName, "vrf-name")
    ctsRolebasedSgtmap.EntityData.AbsolutePath = "Cisco-IOS-XE-trustsec-oper:trustsec-state/cts-rolebased-sgtmaps/" + ctsRolebasedSgtmap.EntityData.SegmentPath
    ctsRolebasedSgtmap.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ctsRolebasedSgtmap.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ctsRolebasedSgtmap.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ctsRolebasedSgtmap.EntityData.Children = types.NewOrderedMap()
    ctsRolebasedSgtmap.EntityData.Leafs = types.NewOrderedMap()
    ctsRolebasedSgtmap.EntityData.Leafs.Append("ip", types.YLeaf{"Ip", ctsRolebasedSgtmap.Ip})
    ctsRolebasedSgtmap.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", ctsRolebasedSgtmap.VrfName})
    ctsRolebasedSgtmap.EntityData.Leafs.Append("sgt", types.YLeaf{"Sgt", ctsRolebasedSgtmap.Sgt})
    ctsRolebasedSgtmap.EntityData.Leafs.Append("source", types.YLeaf{"Source", ctsRolebasedSgtmap.Source})

    ctsRolebasedSgtmap.EntityData.YListKeys = []string {"Ip", "VrfName"}

    return &(ctsRolebasedSgtmap.EntityData)
}

// TrustsecState_CtsRolebasedPolicies
// Role based permissions between a Source Security Group
// and a Destination Security Group are configured by the
// administrator in the Identity Services Engine or in the Device
type TrustsecState_CtsRolebasedPolicies struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Role based permissions between a Source Security Group and a Destination
    // Security Group can be retrieved from the device using a Security Group Tag
    // and Destination Group Tag value. The type is slice of
    // TrustsecState_CtsRolebasedPolicies_CtsRolebasedPolicy.
    CtsRolebasedPolicy []*TrustsecState_CtsRolebasedPolicies_CtsRolebasedPolicy
}

func (ctsRolebasedPolicies *TrustsecState_CtsRolebasedPolicies) GetEntityData() *types.CommonEntityData {
    ctsRolebasedPolicies.EntityData.YFilter = ctsRolebasedPolicies.YFilter
    ctsRolebasedPolicies.EntityData.YangName = "cts-rolebased-policies"
    ctsRolebasedPolicies.EntityData.BundleName = "cisco_ios_xe"
    ctsRolebasedPolicies.EntityData.ParentYangName = "trustsec-state"
    ctsRolebasedPolicies.EntityData.SegmentPath = "cts-rolebased-policies"
    ctsRolebasedPolicies.EntityData.AbsolutePath = "Cisco-IOS-XE-trustsec-oper:trustsec-state/" + ctsRolebasedPolicies.EntityData.SegmentPath
    ctsRolebasedPolicies.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ctsRolebasedPolicies.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ctsRolebasedPolicies.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ctsRolebasedPolicies.EntityData.Children = types.NewOrderedMap()
    ctsRolebasedPolicies.EntityData.Children.Append("cts-rolebased-policy", types.YChild{"CtsRolebasedPolicy", nil})
    for i := range ctsRolebasedPolicies.CtsRolebasedPolicy {
        ctsRolebasedPolicies.EntityData.Children.Append(types.GetSegmentPath(ctsRolebasedPolicies.CtsRolebasedPolicy[i]), types.YChild{"CtsRolebasedPolicy", ctsRolebasedPolicies.CtsRolebasedPolicy[i]})
    }
    ctsRolebasedPolicies.EntityData.Leafs = types.NewOrderedMap()

    ctsRolebasedPolicies.EntityData.YListKeys = []string {}

    return &(ctsRolebasedPolicies.EntityData)
}

// TrustsecState_CtsRolebasedPolicies_CtsRolebasedPolicy
// Role based permissions between a Source Security Group
// and a Destination Security Group can be retrieved from
// the device using a Security Group Tag and Destination
// Group Tag value
type TrustsecState_CtsRolebasedPolicies_CtsRolebasedPolicy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    // b'\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}(\\.\\d+)?(Z|[\\+\\-]\\d{2}:\\d{2})'.
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

func (ctsRolebasedPolicy *TrustsecState_CtsRolebasedPolicies_CtsRolebasedPolicy) GetEntityData() *types.CommonEntityData {
    ctsRolebasedPolicy.EntityData.YFilter = ctsRolebasedPolicy.YFilter
    ctsRolebasedPolicy.EntityData.YangName = "cts-rolebased-policy"
    ctsRolebasedPolicy.EntityData.BundleName = "cisco_ios_xe"
    ctsRolebasedPolicy.EntityData.ParentYangName = "cts-rolebased-policies"
    ctsRolebasedPolicy.EntityData.SegmentPath = "cts-rolebased-policy" + types.AddKeyToken(ctsRolebasedPolicy.SrcSgt, "src-sgt") + types.AddKeyToken(ctsRolebasedPolicy.DstSgt, "dst-sgt")
    ctsRolebasedPolicy.EntityData.AbsolutePath = "Cisco-IOS-XE-trustsec-oper:trustsec-state/cts-rolebased-policies/" + ctsRolebasedPolicy.EntityData.SegmentPath
    ctsRolebasedPolicy.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ctsRolebasedPolicy.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ctsRolebasedPolicy.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ctsRolebasedPolicy.EntityData.Children = types.NewOrderedMap()
    ctsRolebasedPolicy.EntityData.Leafs = types.NewOrderedMap()
    ctsRolebasedPolicy.EntityData.Leafs.Append("src-sgt", types.YLeaf{"SrcSgt", ctsRolebasedPolicy.SrcSgt})
    ctsRolebasedPolicy.EntityData.Leafs.Append("dst-sgt", types.YLeaf{"DstSgt", ctsRolebasedPolicy.DstSgt})
    ctsRolebasedPolicy.EntityData.Leafs.Append("sgacl-name", types.YLeaf{"SgaclName", ctsRolebasedPolicy.SgaclName})
    ctsRolebasedPolicy.EntityData.Leafs.Append("num-of-sgacl", types.YLeaf{"NumOfSgacl", ctsRolebasedPolicy.NumOfSgacl})
    ctsRolebasedPolicy.EntityData.Leafs.Append("monitor-mode", types.YLeaf{"MonitorMode", ctsRolebasedPolicy.MonitorMode})
    ctsRolebasedPolicy.EntityData.Leafs.Append("policy-life-time", types.YLeaf{"PolicyLifeTime", ctsRolebasedPolicy.PolicyLifeTime})
    ctsRolebasedPolicy.EntityData.Leafs.Append("last-updated-time", types.YLeaf{"LastUpdatedTime", ctsRolebasedPolicy.LastUpdatedTime})
    ctsRolebasedPolicy.EntityData.Leafs.Append("total-deny-count", types.YLeaf{"TotalDenyCount", ctsRolebasedPolicy.TotalDenyCount})
    ctsRolebasedPolicy.EntityData.Leafs.Append("total-permit-count", types.YLeaf{"TotalPermitCount", ctsRolebasedPolicy.TotalPermitCount})
    ctsRolebasedPolicy.EntityData.Leafs.Append("software-deny-count", types.YLeaf{"SoftwareDenyCount", ctsRolebasedPolicy.SoftwareDenyCount})
    ctsRolebasedPolicy.EntityData.Leafs.Append("software-permit-count", types.YLeaf{"SoftwarePermitCount", ctsRolebasedPolicy.SoftwarePermitCount})
    ctsRolebasedPolicy.EntityData.Leafs.Append("hardware-deny-count", types.YLeaf{"HardwareDenyCount", ctsRolebasedPolicy.HardwareDenyCount})
    ctsRolebasedPolicy.EntityData.Leafs.Append("hardware-permit-count", types.YLeaf{"HardwarePermitCount", ctsRolebasedPolicy.HardwarePermitCount})
    ctsRolebasedPolicy.EntityData.Leafs.Append("software-monitor-count", types.YLeaf{"SoftwareMonitorCount", ctsRolebasedPolicy.SoftwareMonitorCount})
    ctsRolebasedPolicy.EntityData.Leafs.Append("hardware-monitor-count", types.YLeaf{"HardwareMonitorCount", ctsRolebasedPolicy.HardwareMonitorCount})

    ctsRolebasedPolicy.EntityData.YListKeys = []string {"SrcSgt", "DstSgt"}

    return &(ctsRolebasedPolicy.EntityData)
}

// TrustsecState_CtsSxpConnections
// Trustsec SXP connection is used between Cisco devices
// to propagate Security Group Tags from one device to 
// another device. One of the device will be in Speaker 
// mode or Listener mode or both the devices can be in
// both the connection modes
type TrustsecState_CtsSxpConnections struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Trustsec SXP connection information from a device can be retrieved with the
    // SXP connection peer IP address. Only IPv4 address as Peer IP and default
    // VRF instance in device is supported currently. The type is slice of
    // TrustsecState_CtsSxpConnections_CtsSxpConnection.
    CtsSxpConnection []*TrustsecState_CtsSxpConnections_CtsSxpConnection
}

func (ctsSxpConnections *TrustsecState_CtsSxpConnections) GetEntityData() *types.CommonEntityData {
    ctsSxpConnections.EntityData.YFilter = ctsSxpConnections.YFilter
    ctsSxpConnections.EntityData.YangName = "cts-sxp-connections"
    ctsSxpConnections.EntityData.BundleName = "cisco_ios_xe"
    ctsSxpConnections.EntityData.ParentYangName = "trustsec-state"
    ctsSxpConnections.EntityData.SegmentPath = "cts-sxp-connections"
    ctsSxpConnections.EntityData.AbsolutePath = "Cisco-IOS-XE-trustsec-oper:trustsec-state/" + ctsSxpConnections.EntityData.SegmentPath
    ctsSxpConnections.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ctsSxpConnections.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ctsSxpConnections.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ctsSxpConnections.EntityData.Children = types.NewOrderedMap()
    ctsSxpConnections.EntityData.Children.Append("cts-sxp-connection", types.YChild{"CtsSxpConnection", nil})
    for i := range ctsSxpConnections.CtsSxpConnection {
        ctsSxpConnections.EntityData.Children.Append(types.GetSegmentPath(ctsSxpConnections.CtsSxpConnection[i]), types.YChild{"CtsSxpConnection", ctsSxpConnections.CtsSxpConnection[i]})
    }
    ctsSxpConnections.EntityData.Leafs = types.NewOrderedMap()

    ctsSxpConnections.EntityData.YListKeys = []string {}

    return &(ctsSxpConnections.EntityData)
}

// TrustsecState_CtsSxpConnections_CtsSxpConnection
// Trustsec SXP connection information from a device
// can be retrieved with the SXP connection peer IP
// address. Only IPv4 address as Peer IP and default
// VRF instance in device is supported currently
type TrustsecState_CtsSxpConnections_CtsSxpConnection struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. IP-Address information of the peer of an SXP
    // connection in this device. Only IPv4 address is currently supported. The
    // type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    PeerIp interface{}

    // This attribute is a key. vrf-name string of the VRF instance in this
    // device, to which the peer of an SXP connection belongs to. Only default VRF
    // is supported currently which is indicated by empty string. The type is
    // string.
    VrfName interface{}

    // Source IP-Address of the SXP connection in this device for the given peer
    // IP-Address. The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
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

func (ctsSxpConnection *TrustsecState_CtsSxpConnections_CtsSxpConnection) GetEntityData() *types.CommonEntityData {
    ctsSxpConnection.EntityData.YFilter = ctsSxpConnection.YFilter
    ctsSxpConnection.EntityData.YangName = "cts-sxp-connection"
    ctsSxpConnection.EntityData.BundleName = "cisco_ios_xe"
    ctsSxpConnection.EntityData.ParentYangName = "cts-sxp-connections"
    ctsSxpConnection.EntityData.SegmentPath = "cts-sxp-connection" + types.AddKeyToken(ctsSxpConnection.PeerIp, "peer-ip") + types.AddKeyToken(ctsSxpConnection.VrfName, "vrf-name")
    ctsSxpConnection.EntityData.AbsolutePath = "Cisco-IOS-XE-trustsec-oper:trustsec-state/cts-sxp-connections/" + ctsSxpConnection.EntityData.SegmentPath
    ctsSxpConnection.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ctsSxpConnection.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ctsSxpConnection.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ctsSxpConnection.EntityData.Children = types.NewOrderedMap()
    ctsSxpConnection.EntityData.Leafs = types.NewOrderedMap()
    ctsSxpConnection.EntityData.Leafs.Append("peer-ip", types.YLeaf{"PeerIp", ctsSxpConnection.PeerIp})
    ctsSxpConnection.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", ctsSxpConnection.VrfName})
    ctsSxpConnection.EntityData.Leafs.Append("source-ip", types.YLeaf{"SourceIp", ctsSxpConnection.SourceIp})
    ctsSxpConnection.EntityData.Leafs.Append("speaker-state", types.YLeaf{"SpeakerState", ctsSxpConnection.SpeakerState})
    ctsSxpConnection.EntityData.Leafs.Append("speaker-duration", types.YLeaf{"SpeakerDuration", ctsSxpConnection.SpeakerDuration})
    ctsSxpConnection.EntityData.Leafs.Append("listener-state", types.YLeaf{"ListenerState", ctsSxpConnection.ListenerState})
    ctsSxpConnection.EntityData.Leafs.Append("listener-duration", types.YLeaf{"ListenerDuration", ctsSxpConnection.ListenerDuration})
    ctsSxpConnection.EntityData.Leafs.Append("local-mode", types.YLeaf{"LocalMode", ctsSxpConnection.LocalMode})

    ctsSxpConnection.EntityData.YListKeys = []string {"PeerIp", "VrfName"}

    return &(ctsSxpConnection.EntityData)
}

