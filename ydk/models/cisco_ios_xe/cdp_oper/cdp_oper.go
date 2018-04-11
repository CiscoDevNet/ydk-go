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
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-cdp-oper cdp-neighbor-details}", reflect.TypeOf(CdpNeighborDetails{}))
    ydk.RegisterEntity("Cisco-IOS-XE-cdp-oper:cdp-neighbor-details", reflect.TypeOf(CdpNeighborDetails{}))
}

// CdpDuplex represents CDP duplex modes
type CdpDuplex string

const (
    CdpDuplex_cdp_duplex_unknown CdpDuplex = "cdp-duplex-unknown"

    CdpDuplex_cdp_half_duplex CdpDuplex = "cdp-half-duplex"

    CdpDuplex_cdp_full_duplex CdpDuplex = "cdp-full-duplex"

    CdpDuplex_cdp_half_duplex_mismatch CdpDuplex = "cdp-half-duplex-mismatch"

    CdpDuplex_cdp_full_duplex_mismatch CdpDuplex = "cdp-full-duplex-mismatch"
)

// CdpAdvVersion represents CDP advertized version information
type CdpAdvVersion string

const (
    CdpAdvVersion_cdp_advertised_none CdpAdvVersion = "cdp-advertised-none"

    CdpAdvVersion_cdp_advertised_v1 CdpAdvVersion = "cdp-advertised-v1"

    CdpAdvVersion_cdp_advertised_v2 CdpAdvVersion = "cdp-advertised-v2"
)

// CdpUnidirectionalMode represents CDP unidirectional modes
type CdpUnidirectionalMode string

const (
    CdpUnidirectionalMode_cdp_uni_mode_off CdpUnidirectionalMode = "cdp-uni-mode-off"

    CdpUnidirectionalMode_cdp_uni_mode_send_only CdpUnidirectionalMode = "cdp-uni-mode-send-only"

    CdpUnidirectionalMode_cdp_uni_mode_recv_only CdpUnidirectionalMode = "cdp-uni-mode-recv-only"

    CdpUnidirectionalMode_cdp_uni_mode_unknown CdpUnidirectionalMode = "cdp-uni-mode-unknown"
)

// CdpYesNo represents CDP type yes or no
type CdpYesNo string

const (
    CdpYesNo_cdp_no CdpYesNo = "cdp-no"

    CdpYesNo_cdp_yes CdpYesNo = "cdp-yes"
)

// CdpEnableDisable represents CDP type enable or disable
type CdpEnableDisable string

const (
    CdpEnableDisable_cdp_disable CdpEnableDisable = "cdp-disable"

    CdpEnableDisable_cdp_enable CdpEnableDisable = "cdp-enable"
)

// CdpNeighborDetails
// Cisco CDP neighbor operational data
type CdpNeighborDetails struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of CDP neighbor details. The type is slice of
    // CdpNeighborDetails_CdpNeighborDetail.
    CdpNeighborDetail []CdpNeighborDetails_CdpNeighborDetail
}

func (cdpNeighborDetails *CdpNeighborDetails) GetEntityData() *types.CommonEntityData {
    cdpNeighborDetails.EntityData.YFilter = cdpNeighborDetails.YFilter
    cdpNeighborDetails.EntityData.YangName = "cdp-neighbor-details"
    cdpNeighborDetails.EntityData.BundleName = "cisco_ios_xe"
    cdpNeighborDetails.EntityData.ParentYangName = "Cisco-IOS-XE-cdp-oper"
    cdpNeighborDetails.EntityData.SegmentPath = "Cisco-IOS-XE-cdp-oper:cdp-neighbor-details"
    cdpNeighborDetails.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdpNeighborDetails.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdpNeighborDetails.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdpNeighborDetails.EntityData.Children = make(map[string]types.YChild)
    cdpNeighborDetails.EntityData.Children["cdp-neighbor-detail"] = types.YChild{"CdpNeighborDetail", nil}
    for i := range cdpNeighborDetails.CdpNeighborDetail {
        cdpNeighborDetails.EntityData.Children[types.GetSegmentPath(&cdpNeighborDetails.CdpNeighborDetail[i])] = types.YChild{"CdpNeighborDetail", &cdpNeighborDetails.CdpNeighborDetail[i]}
    }
    cdpNeighborDetails.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cdpNeighborDetails.EntityData)
}

// CdpNeighborDetails_CdpNeighborDetail
// List of CDP neighbor details
type CdpNeighborDetails_CdpNeighborDetail struct {
    EntityData types.CommonEntityData
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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    MgmtAddress interface{}

    // IPv4 address of the device. The type is one of the following types: string
    // with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    IpAddress interface{}

    // IPv6 address of the device. The type is one of the following types: string
    // with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
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
    HelloMessage CdpNeighborDetails_CdpNeighborDetail_HelloMessage

    // This field used to keep inline power.
    PowerRequest CdpNeighborDetails_CdpNeighborDetail_PowerRequest

    // This field used to keep inline power.
    PowerAvailable CdpNeighborDetails_CdpNeighborDetail_PowerAvailable

    // Spare pair PoE TLV is a one octet long. This has following field: Bit      
    // Function                            value/Meaning 0    4-pair PoE Supported
    // 0=No/1=Yes 1    Spare pair Detection/Classification required   0=No/1=Yes 2
    // PD Spare Pair Desired State                    0=Disabled/1=Enabled 3   
    // PSE Spare Pair Operational State               0=No/1=Yes 4:7   Reserved .
    SparePair CdpNeighborDetails_CdpNeighborDetail_SparePair
}

func (cdpNeighborDetail *CdpNeighborDetails_CdpNeighborDetail) GetEntityData() *types.CommonEntityData {
    cdpNeighborDetail.EntityData.YFilter = cdpNeighborDetail.YFilter
    cdpNeighborDetail.EntityData.YangName = "cdp-neighbor-detail"
    cdpNeighborDetail.EntityData.BundleName = "cisco_ios_xe"
    cdpNeighborDetail.EntityData.ParentYangName = "cdp-neighbor-details"
    cdpNeighborDetail.EntityData.SegmentPath = "cdp-neighbor-detail" + "[device-id='" + fmt.Sprintf("%v", cdpNeighborDetail.DeviceId) + "']"
    cdpNeighborDetail.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdpNeighborDetail.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdpNeighborDetail.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdpNeighborDetail.EntityData.Children = make(map[string]types.YChild)
    cdpNeighborDetail.EntityData.Children["hello-message"] = types.YChild{"HelloMessage", &cdpNeighborDetail.HelloMessage}
    cdpNeighborDetail.EntityData.Children["power-request"] = types.YChild{"PowerRequest", &cdpNeighborDetail.PowerRequest}
    cdpNeighborDetail.EntityData.Children["power-available"] = types.YChild{"PowerAvailable", &cdpNeighborDetail.PowerAvailable}
    cdpNeighborDetail.EntityData.Children["spare-pair"] = types.YChild{"SparePair", &cdpNeighborDetail.SparePair}
    cdpNeighborDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    cdpNeighborDetail.EntityData.Leafs["device-id"] = types.YLeaf{"DeviceId", cdpNeighborDetail.DeviceId}
    cdpNeighborDetail.EntityData.Leafs["device-name"] = types.YLeaf{"DeviceName", cdpNeighborDetail.DeviceName}
    cdpNeighborDetail.EntityData.Leafs["local-intf-name"] = types.YLeaf{"LocalIntfName", cdpNeighborDetail.LocalIntfName}
    cdpNeighborDetail.EntityData.Leafs["port-id"] = types.YLeaf{"PortId", cdpNeighborDetail.PortId}
    cdpNeighborDetail.EntityData.Leafs["capability"] = types.YLeaf{"Capability", cdpNeighborDetail.Capability}
    cdpNeighborDetail.EntityData.Leafs["platform-name"] = types.YLeaf{"PlatformName", cdpNeighborDetail.PlatformName}
    cdpNeighborDetail.EntityData.Leafs["version"] = types.YLeaf{"Version", cdpNeighborDetail.Version}
    cdpNeighborDetail.EntityData.Leafs["duplex"] = types.YLeaf{"Duplex", cdpNeighborDetail.Duplex}
    cdpNeighborDetail.EntityData.Leafs["adv-version"] = types.YLeaf{"AdvVersion", cdpNeighborDetail.AdvVersion}
    cdpNeighborDetail.EntityData.Leafs["vty-mgmt-domain"] = types.YLeaf{"VtyMgmtDomain", cdpNeighborDetail.VtyMgmtDomain}
    cdpNeighborDetail.EntityData.Leafs["native-vlan"] = types.YLeaf{"NativeVlan", cdpNeighborDetail.NativeVlan}
    cdpNeighborDetail.EntityData.Leafs["vvid-tag"] = types.YLeaf{"VvidTag", cdpNeighborDetail.VvidTag}
    cdpNeighborDetail.EntityData.Leafs["vvid"] = types.YLeaf{"Vvid", cdpNeighborDetail.Vvid}
    cdpNeighborDetail.EntityData.Leafs["power"] = types.YLeaf{"Power", cdpNeighborDetail.Power}
    cdpNeighborDetail.EntityData.Leafs["unidirectional-mode"] = types.YLeaf{"UnidirectionalMode", cdpNeighborDetail.UnidirectionalMode}
    cdpNeighborDetail.EntityData.Leafs["mgmt-address"] = types.YLeaf{"MgmtAddress", cdpNeighborDetail.MgmtAddress}
    cdpNeighborDetail.EntityData.Leafs["ip-address"] = types.YLeaf{"IpAddress", cdpNeighborDetail.IpAddress}
    cdpNeighborDetail.EntityData.Leafs["ipv6-address"] = types.YLeaf{"Ipv6Address", cdpNeighborDetail.Ipv6Address}
    cdpNeighborDetail.EntityData.Leafs["clns-address"] = types.YLeaf{"ClnsAddress", cdpNeighborDetail.ClnsAddress}
    cdpNeighborDetail.EntityData.Leafs["decnet-addr"] = types.YLeaf{"DecnetAddr", cdpNeighborDetail.DecnetAddr}
    cdpNeighborDetail.EntityData.Leafs["novell-addr"] = types.YLeaf{"NovellAddr", cdpNeighborDetail.NovellAddr}
    cdpNeighborDetail.EntityData.Leafs["second-port-status"] = types.YLeaf{"SecondPortStatus", cdpNeighborDetail.SecondPortStatus}
    cdpNeighborDetail.EntityData.Leafs["table-id"] = types.YLeaf{"TableId", cdpNeighborDetail.TableId}
    return &(cdpNeighborDetail.EntityData)
}

// CdpNeighborDetails_CdpNeighborDetail_HelloMessage
// CDP Protocol Hello message
type CdpNeighborDetails_CdpNeighborDetail_HelloMessage struct {
    EntityData types.CommonEntityData
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

func (helloMessage *CdpNeighborDetails_CdpNeighborDetail_HelloMessage) GetEntityData() *types.CommonEntityData {
    helloMessage.EntityData.YFilter = helloMessage.YFilter
    helloMessage.EntityData.YangName = "hello-message"
    helloMessage.EntityData.BundleName = "cisco_ios_xe"
    helloMessage.EntityData.ParentYangName = "cdp-neighbor-detail"
    helloMessage.EntityData.SegmentPath = "hello-message"
    helloMessage.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    helloMessage.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    helloMessage.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    helloMessage.EntityData.Children = make(map[string]types.YChild)
    helloMessage.EntityData.Leafs = make(map[string]types.YLeaf)
    helloMessage.EntityData.Leafs["oui"] = types.YLeaf{"Oui", helloMessage.Oui}
    helloMessage.EntityData.Leafs["protocol-id"] = types.YLeaf{"ProtocolId", helloMessage.ProtocolId}
    helloMessage.EntityData.Leafs["payload-value"] = types.YLeaf{"PayloadValue", helloMessage.PayloadValue}
    helloMessage.EntityData.Leafs["payload-len"] = types.YLeaf{"PayloadLen", helloMessage.PayloadLen}
    return &(helloMessage.EntityData)
}

// CdpNeighborDetails_CdpNeighborDetail_PowerRequest
// This field used to keep inline power
type CdpNeighborDetails_CdpNeighborDetail_PowerRequest struct {
    EntityData types.CommonEntityData
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

func (powerRequest *CdpNeighborDetails_CdpNeighborDetail_PowerRequest) GetEntityData() *types.CommonEntityData {
    powerRequest.EntityData.YFilter = powerRequest.YFilter
    powerRequest.EntityData.YangName = "power-request"
    powerRequest.EntityData.BundleName = "cisco_ios_xe"
    powerRequest.EntityData.ParentYangName = "cdp-neighbor-detail"
    powerRequest.EntityData.SegmentPath = "power-request"
    powerRequest.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    powerRequest.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    powerRequest.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    powerRequest.EntityData.Children = make(map[string]types.YChild)
    powerRequest.EntityData.Leafs = make(map[string]types.YLeaf)
    powerRequest.EntityData.Leafs["power-request-id"] = types.YLeaf{"PowerRequestId", powerRequest.PowerRequestId}
    powerRequest.EntityData.Leafs["power-man-id"] = types.YLeaf{"PowerManId", powerRequest.PowerManId}
    powerRequest.EntityData.Leafs["power-request-level"] = types.YLeaf{"PowerRequestLevel", powerRequest.PowerRequestLevel}
    return &(powerRequest.EntityData)
}

// CdpNeighborDetails_CdpNeighborDetail_PowerAvailable
// This field used to keep inline power
type CdpNeighborDetails_CdpNeighborDetail_PowerAvailable struct {
    EntityData types.CommonEntityData
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

func (powerAvailable *CdpNeighborDetails_CdpNeighborDetail_PowerAvailable) GetEntityData() *types.CommonEntityData {
    powerAvailable.EntityData.YFilter = powerAvailable.YFilter
    powerAvailable.EntityData.YangName = "power-available"
    powerAvailable.EntityData.BundleName = "cisco_ios_xe"
    powerAvailable.EntityData.ParentYangName = "cdp-neighbor-detail"
    powerAvailable.EntityData.SegmentPath = "power-available"
    powerAvailable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    powerAvailable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    powerAvailable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    powerAvailable.EntityData.Children = make(map[string]types.YChild)
    powerAvailable.EntityData.Leafs = make(map[string]types.YLeaf)
    powerAvailable.EntityData.Leafs["power-request-id"] = types.YLeaf{"PowerRequestId", powerAvailable.PowerRequestId}
    powerAvailable.EntityData.Leafs["power-man-id"] = types.YLeaf{"PowerManId", powerAvailable.PowerManId}
    powerAvailable.EntityData.Leafs["power-available"] = types.YLeaf{"PowerAvailable", powerAvailable.PowerAvailable}
    powerAvailable.EntityData.Leafs["power-man-level"] = types.YLeaf{"PowerManLevel", powerAvailable.PowerManLevel}
    return &(powerAvailable.EntityData)
}

// CdpNeighborDetails_CdpNeighborDetail_SparePair
// Spare pair PoE TLV is a one octet long.
// This has following field:
// Bit            Function                            value/Meaning
// 0    4-pair PoE Supported                           0=No/1=Yes
// 1    Spare pair Detection/Classification required   0=No/1=Yes
// 2    PD Spare Pair Desired State                    0=Disabled/1=Enabled
// 3    PSE Spare Pair Operational State               0=No/1=Yes
// 4:7   Reserved
// 
type CdpNeighborDetails_CdpNeighborDetail_SparePair struct {
    EntityData types.CommonEntityData
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

func (sparePair *CdpNeighborDetails_CdpNeighborDetail_SparePair) GetEntityData() *types.CommonEntityData {
    sparePair.EntityData.YFilter = sparePair.YFilter
    sparePair.EntityData.YangName = "spare-pair"
    sparePair.EntityData.BundleName = "cisco_ios_xe"
    sparePair.EntityData.ParentYangName = "cdp-neighbor-detail"
    sparePair.EntityData.SegmentPath = "spare-pair"
    sparePair.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sparePair.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sparePair.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sparePair.EntityData.Children = make(map[string]types.YChild)
    sparePair.EntityData.Leafs = make(map[string]types.YLeaf)
    sparePair.EntityData.Leafs["spare-pair-poe"] = types.YLeaf{"SparePairPoe", sparePair.SparePairPoe}
    sparePair.EntityData.Leafs["spare-pair-detection-required"] = types.YLeaf{"SparePairDetectionRequired", sparePair.SparePairDetectionRequired}
    sparePair.EntityData.Leafs["spare-pair-pd-config"] = types.YLeaf{"SparePairPdConfig", sparePair.SparePairPdConfig}
    sparePair.EntityData.Leafs["spare-pair-pse-operational"] = types.YLeaf{"SparePairPseOperational", sparePair.SparePairPseOperational}
    return &(sparePair.EntityData)
}

