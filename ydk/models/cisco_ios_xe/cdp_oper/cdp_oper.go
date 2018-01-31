// This module contains a collection of YANG definitions
// for monitoring of Cisco CDP operational information.
// Copyright (c) 2016-2017 by Cisco Systems, Inc.
// All rights reserved.
package cdp_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cdp_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-cdp-oper cdp-neighbour-details}", reflect.TypeOf(CdpNeighbourDetails{}))
    ydk.RegisterEntity("Cisco-IOS-XE-cdp-oper:cdp-neighbour-details", reflect.TypeOf(CdpNeighbourDetails{}))
}

// CdpEnableDisable represents CDP type enable or disable
type CdpEnableDisable string

const (
    CdpEnableDisable_cdp_disable CdpEnableDisable = "cdp-disable"

    CdpEnableDisable_cdp_enable CdpEnableDisable = "cdp-enable"
)

// CdpYesNo represents CDP type yes or no
type CdpYesNo string

const (
    CdpYesNo_cdp_no CdpYesNo = "cdp-no"

    CdpYesNo_cdp_yes CdpYesNo = "cdp-yes"
)

// CdpDuplex represents CDP duplex modes
type CdpDuplex string

const (
    CdpDuplex_cdp_duplex_unknown CdpDuplex = "cdp-duplex-unknown"

    CdpDuplex_cdp_half_duplex CdpDuplex = "cdp-half-duplex"

    CdpDuplex_cdp_full_duplex CdpDuplex = "cdp-full-duplex"

    CdpDuplex_cdp_half_duplex_mismatch CdpDuplex = "cdp-half-duplex-mismatch"

    CdpDuplex_cdp_full_duplex_mismatch CdpDuplex = "cdp-full-duplex-mismatch"
)

// CdpUnidirectionalMode represents CDP unidirectional modes
type CdpUnidirectionalMode string

const (
    CdpUnidirectionalMode_cdp_uni_mode_off CdpUnidirectionalMode = "cdp-uni-mode-off"

    CdpUnidirectionalMode_cdp_uni_mode_send_only CdpUnidirectionalMode = "cdp-uni-mode-send-only"

    CdpUnidirectionalMode_cdp_uni_mode_recv_only CdpUnidirectionalMode = "cdp-uni-mode-recv-only"

    CdpUnidirectionalMode_cdp_uni_mode_unknown CdpUnidirectionalMode = "cdp-uni-mode-unknown"
)

// CdpAdvVersion represents CDP advertized version information
type CdpAdvVersion string

const (
    CdpAdvVersion_cdp_advertised_none CdpAdvVersion = "cdp-advertised-none"

    CdpAdvVersion_cdp_advertised_v1 CdpAdvVersion = "cdp-advertised-v1"

    CdpAdvVersion_cdp_advertised_v2 CdpAdvVersion = "cdp-advertised-v2"
)

// CdpNeighbourDetails
// Cisco CDP neighbour operational data
type CdpNeighbourDetails struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of CDP neighbour details. The type is slice of
    // CdpNeighbourDetails_CdpNeighbourDetail.
    CdpNeighbourDetail []CdpNeighbourDetails_CdpNeighbourDetail
}

func (cdpNeighbourDetails *CdpNeighbourDetails) GetFilter() yfilter.YFilter { return cdpNeighbourDetails.YFilter }

func (cdpNeighbourDetails *CdpNeighbourDetails) SetFilter(yf yfilter.YFilter) { cdpNeighbourDetails.YFilter = yf }

func (cdpNeighbourDetails *CdpNeighbourDetails) GetGoName(yname string) string {
    if yname == "cdp-neighbour-detail" { return "CdpNeighbourDetail" }
    return ""
}

func (cdpNeighbourDetails *CdpNeighbourDetails) GetSegmentPath() string {
    return "Cisco-IOS-XE-cdp-oper:cdp-neighbour-details"
}

func (cdpNeighbourDetails *CdpNeighbourDetails) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cdp-neighbour-detail" {
        for _, c := range cdpNeighbourDetails.CdpNeighbourDetail {
            if cdpNeighbourDetails.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CdpNeighbourDetails_CdpNeighbourDetail{}
        cdpNeighbourDetails.CdpNeighbourDetail = append(cdpNeighbourDetails.CdpNeighbourDetail, child)
        return &cdpNeighbourDetails.CdpNeighbourDetail[len(cdpNeighbourDetails.CdpNeighbourDetail)-1]
    }
    return nil
}

func (cdpNeighbourDetails *CdpNeighbourDetails) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cdpNeighbourDetails.CdpNeighbourDetail {
        children[cdpNeighbourDetails.CdpNeighbourDetail[i].GetSegmentPath()] = &cdpNeighbourDetails.CdpNeighbourDetail[i]
    }
    return children
}

func (cdpNeighbourDetails *CdpNeighbourDetails) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cdpNeighbourDetails *CdpNeighbourDetails) GetBundleName() string { return "cisco_ios_xe" }

func (cdpNeighbourDetails *CdpNeighbourDetails) GetYangName() string { return "cdp-neighbour-details" }

func (cdpNeighbourDetails *CdpNeighbourDetails) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cdpNeighbourDetails *CdpNeighbourDetails) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cdpNeighbourDetails *CdpNeighbourDetails) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cdpNeighbourDetails *CdpNeighbourDetails) SetParent(parent types.Entity) { cdpNeighbourDetails.parent = parent }

func (cdpNeighbourDetails *CdpNeighbourDetails) GetParent() types.Entity { return cdpNeighbourDetails.parent }

func (cdpNeighbourDetails *CdpNeighbourDetails) GetParentYangName() string { return "Cisco-IOS-XE-cdp-oper" }

// CdpNeighbourDetails_CdpNeighbourDetail
// List of CDP neighbour details
type CdpNeighbourDetails_CdpNeighbourDetail struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Device number of this device, Used as a key in the
    // device list. The type is interface{} with range: 0..4294967295.
    DeviceId interface{}

    // Device name in the form of a character string. By default device ID is
    // either the device's fully-qualified host name (including the domain name)
    // or the device's hardware serial number in ASCII. The type is string.
    DeviceName interface{}

    // The port or interface on which CDP packets are received. The type is
    // string.
    LocalIntfName interface{}

    // Neighbor device's port or interface on which the CDP packets are
    // multicasted. The type is string.
    PortId interface{}

    // Identifies the functional capability of the device. The capability types
    // that can be discovered are: R-Router T-Transparent bridge B-Source-routing
    // bridge S-Switch H-Host I-device is using IGMP r-Repeater. The type is
    // string.
    Capability interface{}

    // Identifies the platform information of the device. For example, Cisco 4500.
    // The type is string.
    PlatformName interface{}

    // Version Contains the device software release information. The type is
    // string.
    Version interface{}

    // Indicates the duplex configuration of the Cisco Discovery Protocol 
    // broadcast interface. This information is used by network operators to
    // diagnose  connectivity problems between adjacent network devices. The type
    // is CdpDuplex.
    Duplex interface{}

    // CDP header version of the advertisement that last filled this cache entry.
    // The type is CdpAdvVersion.
    AdvVersion interface{}

    // Advertises the configured VLAN Trunking Protocol (VTP)-management-domain
    // name of the system. This name is used by network operators to verify
    // VTP-domain configuration in adjacent network nodes. The type is string.
    VtyMgmtDomain interface{}

    // Indicates, per interface, the assumed VLAN for untagged packets on the
    // interface. Cisco Discovery Protocol learns the native VLAN for an
    // interface. This field is implemented only for interfaces that support the
    // IEEE 802.1Q protocol Remote port's native VLAN [0..1k/4k]; 0 == not
    // received. The type is interface{} with range: 0..65535.
    NativeVlan interface{}

    // Appliance id for appliance vlan Appliance ID - Type of device attached to
    // port advertised in the appliance TLV. The type is interface{} with range:
    // 0..255.
    VvidTag interface{}

    // Appliance VLAN ID - VLAN on the device used by the appliance, for instance
    // if the appliance is an IP phone, this is the voice VLAN. The type is
    // interface{} with range: 0..65535.
    Vvid interface{}

    // This field shows the power in milliwatts device is using. The type is
    // interface{} with range: 0..4294967295.
    Power interface{}

    // Specifies ports to unidirectionally transmit or receive traffic.
    // Unidirectional Ethernet uses only one strand of fiber for either
    // transmitting or receiving one-way traffic for the GigaPort, instead of  two
    // strands of fiber for a full-duplex. The type is CdpUnidirectionalMode.
    UnidirectionalMode interface{}

    // Device's management addresses. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    MgmtAddress interface{}

    // IPv4 address of the device. The type is one of the following types: string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddress interface{}

    // IPv6 address of the device. The type is one of the following types: string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}

    // CLNS address of the device. The type is string.
    ClnsAddress interface{}

    // DECNET address of the device. The type is string.
    DecnetAddr interface{}

    // Novell address of the device. It has a 4 byte net number followed by 6
    // bytes of  node information. The type is string.
    NovellAddr interface{}

    // Used to keep PC port status on IP phone. The type is string.
    SecondPortStatus interface{}

    // Table id of ip routing process. The type is interface{} with range:
    // 0..65535.
    TableId interface{}

    // CDP Protocol Hello message.
    HelloMessage CdpNeighbourDetails_CdpNeighbourDetail_HelloMessage

    // This field used to keep inline power.
    PowerRequest CdpNeighbourDetails_CdpNeighbourDetail_PowerRequest

    // This field used to keep inline power.
    PowerAvailable CdpNeighbourDetails_CdpNeighbourDetail_PowerAvailable

    // Spare pair PoE TLV is a one octet long. This has following field: Bit      
    // Function                            value/Meaning 0    4-pair PoE Supported
    // 0=No/1=Yes 1    Spare pair Detection/Classification required   0=No/1=Yes 2
    // PD Spare Pair Desired State                    0=Disabled/1=Enabled 3   
    // PSE Spare Pair Operational State               0=No/1=Yes 4:7   Reserved .
    SparePair CdpNeighbourDetails_CdpNeighbourDetail_SparePair
}

func (cdpNeighbourDetail *CdpNeighbourDetails_CdpNeighbourDetail) GetFilter() yfilter.YFilter { return cdpNeighbourDetail.YFilter }

func (cdpNeighbourDetail *CdpNeighbourDetails_CdpNeighbourDetail) SetFilter(yf yfilter.YFilter) { cdpNeighbourDetail.YFilter = yf }

func (cdpNeighbourDetail *CdpNeighbourDetails_CdpNeighbourDetail) GetGoName(yname string) string {
    if yname == "device-id" { return "DeviceId" }
    if yname == "device-name" { return "DeviceName" }
    if yname == "local-intf-name" { return "LocalIntfName" }
    if yname == "port-id" { return "PortId" }
    if yname == "capability" { return "Capability" }
    if yname == "platform-name" { return "PlatformName" }
    if yname == "version" { return "Version" }
    if yname == "duplex" { return "Duplex" }
    if yname == "adv-version" { return "AdvVersion" }
    if yname == "vty-mgmt-domain" { return "VtyMgmtDomain" }
    if yname == "native-vlan" { return "NativeVlan" }
    if yname == "vvid-tag" { return "VvidTag" }
    if yname == "vvid" { return "Vvid" }
    if yname == "power" { return "Power" }
    if yname == "unidirectional-mode" { return "UnidirectionalMode" }
    if yname == "mgmt-address" { return "MgmtAddress" }
    if yname == "ip-address" { return "IpAddress" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    if yname == "clns-address" { return "ClnsAddress" }
    if yname == "decnet-addr" { return "DecnetAddr" }
    if yname == "novell-addr" { return "NovellAddr" }
    if yname == "second-port-status" { return "SecondPortStatus" }
    if yname == "table-id" { return "TableId" }
    if yname == "hello-message" { return "HelloMessage" }
    if yname == "power-request" { return "PowerRequest" }
    if yname == "power-available" { return "PowerAvailable" }
    if yname == "spare-pair" { return "SparePair" }
    return ""
}

func (cdpNeighbourDetail *CdpNeighbourDetails_CdpNeighbourDetail) GetSegmentPath() string {
    return "cdp-neighbour-detail" + "[device-id='" + fmt.Sprintf("%v", cdpNeighbourDetail.DeviceId) + "']"
}

func (cdpNeighbourDetail *CdpNeighbourDetails_CdpNeighbourDetail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hello-message" {
        return &cdpNeighbourDetail.HelloMessage
    }
    if childYangName == "power-request" {
        return &cdpNeighbourDetail.PowerRequest
    }
    if childYangName == "power-available" {
        return &cdpNeighbourDetail.PowerAvailable
    }
    if childYangName == "spare-pair" {
        return &cdpNeighbourDetail.SparePair
    }
    return nil
}

func (cdpNeighbourDetail *CdpNeighbourDetails_CdpNeighbourDetail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["hello-message"] = &cdpNeighbourDetail.HelloMessage
    children["power-request"] = &cdpNeighbourDetail.PowerRequest
    children["power-available"] = &cdpNeighbourDetail.PowerAvailable
    children["spare-pair"] = &cdpNeighbourDetail.SparePair
    return children
}

func (cdpNeighbourDetail *CdpNeighbourDetails_CdpNeighbourDetail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["device-id"] = cdpNeighbourDetail.DeviceId
    leafs["device-name"] = cdpNeighbourDetail.DeviceName
    leafs["local-intf-name"] = cdpNeighbourDetail.LocalIntfName
    leafs["port-id"] = cdpNeighbourDetail.PortId
    leafs["capability"] = cdpNeighbourDetail.Capability
    leafs["platform-name"] = cdpNeighbourDetail.PlatformName
    leafs["version"] = cdpNeighbourDetail.Version
    leafs["duplex"] = cdpNeighbourDetail.Duplex
    leafs["adv-version"] = cdpNeighbourDetail.AdvVersion
    leafs["vty-mgmt-domain"] = cdpNeighbourDetail.VtyMgmtDomain
    leafs["native-vlan"] = cdpNeighbourDetail.NativeVlan
    leafs["vvid-tag"] = cdpNeighbourDetail.VvidTag
    leafs["vvid"] = cdpNeighbourDetail.Vvid
    leafs["power"] = cdpNeighbourDetail.Power
    leafs["unidirectional-mode"] = cdpNeighbourDetail.UnidirectionalMode
    leafs["mgmt-address"] = cdpNeighbourDetail.MgmtAddress
    leafs["ip-address"] = cdpNeighbourDetail.IpAddress
    leafs["ipv6-address"] = cdpNeighbourDetail.Ipv6Address
    leafs["clns-address"] = cdpNeighbourDetail.ClnsAddress
    leafs["decnet-addr"] = cdpNeighbourDetail.DecnetAddr
    leafs["novell-addr"] = cdpNeighbourDetail.NovellAddr
    leafs["second-port-status"] = cdpNeighbourDetail.SecondPortStatus
    leafs["table-id"] = cdpNeighbourDetail.TableId
    return leafs
}

func (cdpNeighbourDetail *CdpNeighbourDetails_CdpNeighbourDetail) GetBundleName() string { return "cisco_ios_xe" }

func (cdpNeighbourDetail *CdpNeighbourDetails_CdpNeighbourDetail) GetYangName() string { return "cdp-neighbour-detail" }

func (cdpNeighbourDetail *CdpNeighbourDetails_CdpNeighbourDetail) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cdpNeighbourDetail *CdpNeighbourDetails_CdpNeighbourDetail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cdpNeighbourDetail *CdpNeighbourDetails_CdpNeighbourDetail) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cdpNeighbourDetail *CdpNeighbourDetails_CdpNeighbourDetail) SetParent(parent types.Entity) { cdpNeighbourDetail.parent = parent }

func (cdpNeighbourDetail *CdpNeighbourDetails_CdpNeighbourDetail) GetParent() types.Entity { return cdpNeighbourDetail.parent }

func (cdpNeighbourDetail *CdpNeighbourDetails_CdpNeighbourDetail) GetParentYangName() string { return "cdp-neighbour-details" }

// CdpNeighbourDetails_CdpNeighbourDetail_HelloMessage
// CDP Protocol Hello message
type CdpNeighbourDetails_CdpNeighbourDetail_HelloMessage struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OUI - org unique identifier for Cisco is 0x00000C. The type is string.
    Oui interface{}

    // Protocol identifier. This is the identifier of the cluster. The type is
    // string.
    ProtocolId interface{}

    // Payload value - combination of the device and addresses. The type is
    // string.
    PayloadValue interface{}

    // Payload length. The type is interface{} with range: 0..65535.
    PayloadLen interface{}
}

func (helloMessage *CdpNeighbourDetails_CdpNeighbourDetail_HelloMessage) GetFilter() yfilter.YFilter { return helloMessage.YFilter }

func (helloMessage *CdpNeighbourDetails_CdpNeighbourDetail_HelloMessage) SetFilter(yf yfilter.YFilter) { helloMessage.YFilter = yf }

func (helloMessage *CdpNeighbourDetails_CdpNeighbourDetail_HelloMessage) GetGoName(yname string) string {
    if yname == "oui" { return "Oui" }
    if yname == "protocol-id" { return "ProtocolId" }
    if yname == "payload-value" { return "PayloadValue" }
    if yname == "payload-len" { return "PayloadLen" }
    return ""
}

func (helloMessage *CdpNeighbourDetails_CdpNeighbourDetail_HelloMessage) GetSegmentPath() string {
    return "hello-message"
}

func (helloMessage *CdpNeighbourDetails_CdpNeighbourDetail_HelloMessage) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (helloMessage *CdpNeighbourDetails_CdpNeighbourDetail_HelloMessage) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (helloMessage *CdpNeighbourDetails_CdpNeighbourDetail_HelloMessage) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["oui"] = helloMessage.Oui
    leafs["protocol-id"] = helloMessage.ProtocolId
    leafs["payload-value"] = helloMessage.PayloadValue
    leafs["payload-len"] = helloMessage.PayloadLen
    return leafs
}

func (helloMessage *CdpNeighbourDetails_CdpNeighbourDetail_HelloMessage) GetBundleName() string { return "cisco_ios_xe" }

func (helloMessage *CdpNeighbourDetails_CdpNeighbourDetail_HelloMessage) GetYangName() string { return "hello-message" }

func (helloMessage *CdpNeighbourDetails_CdpNeighbourDetail_HelloMessage) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (helloMessage *CdpNeighbourDetails_CdpNeighbourDetail_HelloMessage) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (helloMessage *CdpNeighbourDetails_CdpNeighbourDetail_HelloMessage) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (helloMessage *CdpNeighbourDetails_CdpNeighbourDetail_HelloMessage) SetParent(parent types.Entity) { helloMessage.parent = parent }

func (helloMessage *CdpNeighbourDetails_CdpNeighbourDetail_HelloMessage) GetParent() types.Entity { return helloMessage.parent }

func (helloMessage *CdpNeighbourDetails_CdpNeighbourDetail_HelloMessage) GetParentYangName() string { return "cdp-neighbour-detail" }

// CdpNeighbourDetails_CdpNeighbourDetail_PowerRequest
// This field used to keep inline power
type CdpNeighbourDetails_CdpNeighbourDetail_PowerRequest struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The last power request ID received echoes the Request-ID field last
    // received in a power requested TLV. It is 0 if no power requested TLV was
    // received since the port last became active. The type is interface{} with
    // range: 0..65535.
    PowerRequestId interface{}

    // This field increments by 1 each time the Available-Power or Management
    // Power fields change, a power requested TLV is received with a  Request-ID
    // field which is different from the last received set or when  the port goes
    // down. The type is interface{} with range: 0..65535.
    PowerManId interface{}

    // Power to be transmitted by a powerable device in order to negotiate a
    // suitable power level with the supplier of the network power. The type is
    // string.
    PowerRequestLevel interface{}
}

func (powerRequest *CdpNeighbourDetails_CdpNeighbourDetail_PowerRequest) GetFilter() yfilter.YFilter { return powerRequest.YFilter }

func (powerRequest *CdpNeighbourDetails_CdpNeighbourDetail_PowerRequest) SetFilter(yf yfilter.YFilter) { powerRequest.YFilter = yf }

func (powerRequest *CdpNeighbourDetails_CdpNeighbourDetail_PowerRequest) GetGoName(yname string) string {
    if yname == "power-request-id" { return "PowerRequestId" }
    if yname == "power-man-id" { return "PowerManId" }
    if yname == "power-request-level" { return "PowerRequestLevel" }
    return ""
}

func (powerRequest *CdpNeighbourDetails_CdpNeighbourDetail_PowerRequest) GetSegmentPath() string {
    return "power-request"
}

func (powerRequest *CdpNeighbourDetails_CdpNeighbourDetail_PowerRequest) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (powerRequest *CdpNeighbourDetails_CdpNeighbourDetail_PowerRequest) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (powerRequest *CdpNeighbourDetails_CdpNeighbourDetail_PowerRequest) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["power-request-id"] = powerRequest.PowerRequestId
    leafs["power-man-id"] = powerRequest.PowerManId
    leafs["power-request-level"] = powerRequest.PowerRequestLevel
    return leafs
}

func (powerRequest *CdpNeighbourDetails_CdpNeighbourDetail_PowerRequest) GetBundleName() string { return "cisco_ios_xe" }

func (powerRequest *CdpNeighbourDetails_CdpNeighbourDetail_PowerRequest) GetYangName() string { return "power-request" }

func (powerRequest *CdpNeighbourDetails_CdpNeighbourDetail_PowerRequest) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (powerRequest *CdpNeighbourDetails_CdpNeighbourDetail_PowerRequest) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (powerRequest *CdpNeighbourDetails_CdpNeighbourDetail_PowerRequest) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (powerRequest *CdpNeighbourDetails_CdpNeighbourDetail_PowerRequest) SetParent(parent types.Entity) { powerRequest.parent = parent }

func (powerRequest *CdpNeighbourDetails_CdpNeighbourDetail_PowerRequest) GetParent() types.Entity { return powerRequest.parent }

func (powerRequest *CdpNeighbourDetails_CdpNeighbourDetail_PowerRequest) GetParentYangName() string { return "cdp-neighbour-detail" }

// CdpNeighbourDetails_CdpNeighbourDetail_PowerAvailable
// This field used to keep inline power
type CdpNeighbourDetails_CdpNeighbourDetail_PowerAvailable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The last power request ID received echoes the Request-ID field last
    // received in a power requested TLV. It is 0 if no power requested TLV was
    // received since the port last became active. The type is interface{} with
    // range: 0..65535.
    PowerRequestId interface{}

    // This field increments by 1 each time the Available-Power or Management
    // Power fields change, a power requested TLV is received with a  Request-ID
    // field which is different from the last received set or when  the port goes
    // down. The type is interface{} with range: 0..65535.
    PowerManId interface{}

    // The amount of power consumed by the specified port in watts. The type is
    // interface{} with range: 0..4294967295.
    PowerAvailable interface{}

    // Management Power Level -- The request of the supplier to the powered device
    // for the power consumption TLV. The 200/300 switches always display No
    // Preference since the switch is a power provide. The type is interface{}
    // with range: 0..4294967295.
    PowerManLevel interface{}
}

func (powerAvailable *CdpNeighbourDetails_CdpNeighbourDetail_PowerAvailable) GetFilter() yfilter.YFilter { return powerAvailable.YFilter }

func (powerAvailable *CdpNeighbourDetails_CdpNeighbourDetail_PowerAvailable) SetFilter(yf yfilter.YFilter) { powerAvailable.YFilter = yf }

func (powerAvailable *CdpNeighbourDetails_CdpNeighbourDetail_PowerAvailable) GetGoName(yname string) string {
    if yname == "power-request-id" { return "PowerRequestId" }
    if yname == "power-man-id" { return "PowerManId" }
    if yname == "power-available" { return "PowerAvailable" }
    if yname == "power-man-level" { return "PowerManLevel" }
    return ""
}

func (powerAvailable *CdpNeighbourDetails_CdpNeighbourDetail_PowerAvailable) GetSegmentPath() string {
    return "power-available"
}

func (powerAvailable *CdpNeighbourDetails_CdpNeighbourDetail_PowerAvailable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (powerAvailable *CdpNeighbourDetails_CdpNeighbourDetail_PowerAvailable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (powerAvailable *CdpNeighbourDetails_CdpNeighbourDetail_PowerAvailable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["power-request-id"] = powerAvailable.PowerRequestId
    leafs["power-man-id"] = powerAvailable.PowerManId
    leafs["power-available"] = powerAvailable.PowerAvailable
    leafs["power-man-level"] = powerAvailable.PowerManLevel
    return leafs
}

func (powerAvailable *CdpNeighbourDetails_CdpNeighbourDetail_PowerAvailable) GetBundleName() string { return "cisco_ios_xe" }

func (powerAvailable *CdpNeighbourDetails_CdpNeighbourDetail_PowerAvailable) GetYangName() string { return "power-available" }

func (powerAvailable *CdpNeighbourDetails_CdpNeighbourDetail_PowerAvailable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (powerAvailable *CdpNeighbourDetails_CdpNeighbourDetail_PowerAvailable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (powerAvailable *CdpNeighbourDetails_CdpNeighbourDetail_PowerAvailable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (powerAvailable *CdpNeighbourDetails_CdpNeighbourDetail_PowerAvailable) SetParent(parent types.Entity) { powerAvailable.parent = parent }

func (powerAvailable *CdpNeighbourDetails_CdpNeighbourDetail_PowerAvailable) GetParent() types.Entity { return powerAvailable.parent }

func (powerAvailable *CdpNeighbourDetails_CdpNeighbourDetail_PowerAvailable) GetParentYangName() string { return "cdp-neighbour-detail" }

// CdpNeighbourDetails_CdpNeighbourDetail_SparePair
// Spare pair PoE TLV is a one octet long.
// This has following field:
// Bit            Function                            value/Meaning
// 0    4-pair PoE Supported                           0=No/1=Yes
// 1    Spare pair Detection/Classification required   0=No/1=Yes
// 2    PD Spare Pair Desired State                    0=Disabled/1=Enabled
// 3    PSE Spare Pair Operational State               0=No/1=Yes
// 4:7   Reserved
// 
type CdpNeighbourDetails_CdpNeighbourDetail_SparePair struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Spare pair PoE TLV is a one octet long this field represents 4-pair PoE
    // Supported or not. The type is CdpYesNo.
    SparePairPoe interface{}

    // Spare pair PoE TLV is a one octet long this field represents Spare pair
    // Detection or Classification is required or not. The type is CdpYesNo.
    SparePairDetectionRequired interface{}

    // Spare pair PoE TLV is a one octet long this field represents Powered
    // Device(PD) Spare Pair Desired State. The type is CdpEnableDisable.
    SparePairPdConfig interface{}

    // Spare pair PoE TLV is a one octet long this field represents Power Source
    // Equipment(PSE) Spare Pair Operational State. The type is CdpYesNo.
    SparePairPseOperational interface{}
}

func (sparePair *CdpNeighbourDetails_CdpNeighbourDetail_SparePair) GetFilter() yfilter.YFilter { return sparePair.YFilter }

func (sparePair *CdpNeighbourDetails_CdpNeighbourDetail_SparePair) SetFilter(yf yfilter.YFilter) { sparePair.YFilter = yf }

func (sparePair *CdpNeighbourDetails_CdpNeighbourDetail_SparePair) GetGoName(yname string) string {
    if yname == "spare-pair-poe" { return "SparePairPoe" }
    if yname == "spare-pair-detection-required" { return "SparePairDetectionRequired" }
    if yname == "spare-pair-pd-config" { return "SparePairPdConfig" }
    if yname == "spare-pair-pse-operational" { return "SparePairPseOperational" }
    return ""
}

func (sparePair *CdpNeighbourDetails_CdpNeighbourDetail_SparePair) GetSegmentPath() string {
    return "spare-pair"
}

func (sparePair *CdpNeighbourDetails_CdpNeighbourDetail_SparePair) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sparePair *CdpNeighbourDetails_CdpNeighbourDetail_SparePair) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sparePair *CdpNeighbourDetails_CdpNeighbourDetail_SparePair) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["spare-pair-poe"] = sparePair.SparePairPoe
    leafs["spare-pair-detection-required"] = sparePair.SparePairDetectionRequired
    leafs["spare-pair-pd-config"] = sparePair.SparePairPdConfig
    leafs["spare-pair-pse-operational"] = sparePair.SparePairPseOperational
    return leafs
}

func (sparePair *CdpNeighbourDetails_CdpNeighbourDetail_SparePair) GetBundleName() string { return "cisco_ios_xe" }

func (sparePair *CdpNeighbourDetails_CdpNeighbourDetail_SparePair) GetYangName() string { return "spare-pair" }

func (sparePair *CdpNeighbourDetails_CdpNeighbourDetail_SparePair) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (sparePair *CdpNeighbourDetails_CdpNeighbourDetail_SparePair) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (sparePair *CdpNeighbourDetails_CdpNeighbourDetail_SparePair) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (sparePair *CdpNeighbourDetails_CdpNeighbourDetail_SparePair) SetParent(parent types.Entity) { sparePair.parent = parent }

func (sparePair *CdpNeighbourDetails_CdpNeighbourDetail_SparePair) GetParent() types.Entity { return sparePair.parent }

func (sparePair *CdpNeighbourDetails_CdpNeighbourDetail_SparePair) GetParentYangName() string { return "cdp-neighbour-detail" }

