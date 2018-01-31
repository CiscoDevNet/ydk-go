// This module contains a collection of YANG definitions for
// managing network interfaces.
// 
// Copyright (c) 2014 IETF Trust and the persons identified as
// authors of the code.  All rights reserved.
// 
// Redistribution and use in source and binary forms, with or
// without modification, is permitted pursuant to, and subject
// to the license terms contained in, the Simplified BSD License
// set forth in Section 4.c of the IETF Trust's Legal Provisions
// Relating to IETF Documents
// (http://trustee.ietf.org/license-info).
// 
// This version of this YANG module is part of RFC 7223; see
// the RFC itself for full legal notices.
package interfaces

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/ietf"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package interfaces"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:ietf-interfaces interfaces}", reflect.TypeOf(Interfaces{}))
    ydk.RegisterEntity("ietf-interfaces:interfaces", reflect.TypeOf(Interfaces{}))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:ietf-interfaces interfaces-state}", reflect.TypeOf(InterfacesState{}))
    ydk.RegisterEntity("ietf-interfaces:interfaces-state", reflect.TypeOf(InterfacesState{}))
}

type InterfaceType struct {
}

func (id InterfaceType) String() string {
	return "ietf-interfaces:interface-type"
}

// Interfaces
// Interface configuration parameters.
type Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The list of configured interfaces on the device.  The operational state of
    // an interface is available in the /interfaces-state/interface list.  If the
    // configuration of a system-controlled interface cannot be used by the system
    // (e.g., the interface hardware present does not match the interface type),
    // then the configuration is not applied to the system-controlled interface
    // shown in the /interfaces-state/interface list.  If the configuration of a
    // user-controlled interface cannot be used by the system, the configured
    // interface is not instantiated in the /interfaces-state/interface list. The
    // type is slice of Interfaces_Interface.
    Interface []Interfaces_Interface
}

func (interfaces *Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *Interfaces) GetSegmentPath() string {
    return "ietf-interfaces:interfaces"
}

func (interfaces *Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *Interfaces) GetBundleName() string { return "ietf" }

func (interfaces *Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *Interfaces) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (interfaces *Interfaces) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (interfaces *Interfaces) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (interfaces *Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *Interfaces) GetParentYangName() string { return "ietf-interfaces" }

// Interfaces_Interface
// The list of configured interfaces on the device.
// 
// The operational state of an interface is available in the
// /interfaces-state/interface list.  If the configuration of a
// system-controlled interface cannot be used by the system
// (e.g., the interface hardware present does not match the
// interface type), then the configuration is not applied to
// the system-controlled interface shown in the
// /interfaces-state/interface list.  If the configuration
// of a user-controlled interface cannot be used by the system,
// the configured interface is not instantiated in the
// /interfaces-state/interface list.
type Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the interface.  A device MAY restrict
    // the allowed values for this leaf, possibly depending on the type of the
    // interface. For system-controlled interfaces, this leaf is the
    // device-specific name of the interface.  The 'config false' list
    // /interfaces-state/interface contains the currently existing interfaces on
    // the device.  If a client tries to create configuration for a
    // system-controlled interface that is not present in the
    // /interfaces-state/interface list, the server MAY reject the request if the
    // implementation does not support pre-provisioning of interfaces or if the
    // name refers to an interface that can never exist in the system.  A NETCONF
    // server MUST reply with an rpc-error with the error-tag 'invalid-value' in
    // this case.  If the device supports pre-provisioning of interface
    // configuration, the 'pre-provisioning' feature is advertised.  If the device
    // allows arbitrarily named user-controlled interfaces, the 'arbitrary-names'
    // feature is advertised.  When a configured user-controlled interface is
    // created by the system, it is instantiated with the same name in the
    // /interface-state/interface list. The type is string.
    Name interface{}

    // A textual description of the interface.  A server implementation MAY map
    // this leaf to the ifAlias MIB object.  Such an implementation needs to use
    // some mechanism to handle the differences in size and characters allowed
    // between this leaf and ifAlias.  The definition of such a mechanism is
    // outside the scope of this document.  Since ifAlias is defined to be stored
    // in non-volatile storage, the MIB implementation MUST map ifAlias to the
    // value of 'description' in the persistently stored datastore.  Specifically,
    // if the device supports ':startup', when ifAlias is read the device MUST
    // return the value of 'description' in the 'startup' datastore, and when it
    // is written, it MUST be written to the 'running' and 'startup' datastores. 
    // Note that it is up to the implementation to  decide whether to modify this
    // single leaf in 'startup' or perform an implicit copy-config from 'running'
    // to 'startup'.  If the device does not support ':startup', ifAlias MUST be
    // mapped to the 'description' leaf in the 'running' datastore. The type is
    // string.
    Description interface{}

    // The type of the interface.  When an interface entry is created, a server
    // MAY initialize the type leaf with a valid value, e.g., if it is possible to
    // derive the type from the name of the interface.  If a client tries to set
    // the type of an interface to a value that can never be used by the system,
    // e.g., if the type is not supported or if the type does not match the name
    // of the interface, the server MUST reject the request. A NETCONF server MUST
    // reply with an rpc-error with the error-tag 'invalid-value' in this case.
    // The type is one of the following:
    // IanaInterfaceTypeVoicefxoAtmvciendptPropbwap2MpPropdocswirelessdownstreamV11SoftwareloopbackHdlcVoicefgdosFastetherfxDvbtdmNfasIfpwtypeL2VlanAdsl2PlusIeee802154VoicefxsDvbrcsmaclayerIdslInfinibandDdnx25Wwanpp2DocscableupstreamEthernet3MbitDigitalpowerlineH323ProxyGtpIpoveratmAlueponImtIpswitchMsdslDvbrccmaclayerSmdsdxiVoiceoveratmArapFastetherMpcLinegroupHippiRprDs1FdlSonetvtVoiceencapSs7SiglinkArcnetActelismetaloopQllcRfc877X25MpegtransportX25MlpVirtualtgHostpadStarlanIso88025DtrIbm370ParchanAdsl2OtnotuPropwirelessp2PInterleaveIsupRegular1822Gr303RdtPropdocswirelessmaclayerAsyncRadiomacOpticalchannelgroupSixtofourPropdocswirelessupstreamQ2931FddiPropcnlsAal2DvbasioutAluelpCiscoislvlanDocscableupstreamrfportAal5FrdlciendptHippiinterfaceL3IpvlanMiox25HssiAtmvirtualAlugpononuRfc1483CnrSipsigMyrinetDlswGigabitethernetX25PleLmpOpticaltransportSdlcVoiceemX86LapsG9982Iso88022LlcDvbasiinBgppolicyaccountingAluepononuMfsiglinkDcnAtmdxiVoiceoverframerelayGfpSonetoverheadchannelVmwarevirtualnicFciplinkIpoverclawCoffeeRadslVdsl2Rs232E1ReachdslVoiceovercableTr008VoiceoveripAtmDs3Ds0Ds1SrpDocscabledownstreamDvbrcstdmaG9983PlcFramerelaympiMvlPropmultiplexorVoicedidCompositelinkProteon10MbitAtmbondFrf16MfrbundleCctemulMplstunnelGponVdslPosIeee8023AdlagDocscablemaclayerDocscablemcmtsdownstreamPppFramerelayEplrsVmwarenicteamCabledownstreamrfportMacsecuncontrolledifIso88023CsmacdUsbAtmfuniTelinkPon622EconetTdlcDs0BundleFastIeee1394CblvectastarRsrbFramerelayinterconnectIsdnsPppmultilinkbundleAflane8025LapbAflane8023LapdIsdnuLapfCapwapwtpvirtualradioIfvfitypeX25HuntgroupParaMacseccontrolledifIso88024TokenbusLocaltalkHyperchannelMediamailoveripIfGsnCapwapdot11ProfileL3IpxvlanAtmsubinterfacePrimaryisdnProteon80MbitIso88026ManDigitalwrapperoverheadchannelDocscableupstreamchannelOpticalchannelEthernetcsmacdBitsTunnelHdsl2FramerelayserviceMplsIeee80211Ieee80212Mocaversion1SonetEsconAlueponlogicallinkG703At2MbUltraDvbrccdownstreamSiptgSmdsicipBridgeAtmlogicalProppointtopointserialV35V36V37IpGr303IdtBasicisdnG703At64KArcnetplusPipDtmSlipHiperlan2AdslIeee80216WmanAtmimaIsdnCapwapdot11BssSipPdnetherloop2VoiceebsIpforwardIso88025CrfpintPropvirtualWwanppOtherPon155QamOtnoduIso88025FiberChannelVoiceemfgdAlugponphysicaluniA12MppswitchIlanPdnetherloop1X213SonetpathVoicefgdeanaIso88025TokenringPropatmAlueponphysicaluniStacktostackFrforwardHomepnaSdslVirtualipaddressBscAtmradioAviciopticaletherG9981FibrechannelShdslEonH323GatekeeperHdh1822DvbrccupstreamNsipTransphdlcTermpadIpovercdlcCesModem.
    // This attribute is mandatory.
    Type interface{}

    // This leaf contains the configured, desired state of the interface.  Systems
    // that implement the IF-MIB use the value of this leaf in the 'running'
    // datastore to set IF-MIB.ifAdminStatus to 'up' or 'down' after an ifEntry
    // has been initialized, as described in RFC 2863.    Changes in this leaf in
    // the 'running' datastore are reflected in ifAdminStatus, but if
    // ifAdminStatus is changed over SNMP, this leaf is not affected. The type is
    // bool. The default value is true.
    Enabled interface{}

    // Controls whether linkUp/linkDown SNMP notifications should be generated for
    // this interface.  If this node is not configured, the value 'enabled' is
    // operationally used by the server for interfaces that do not operate on top
    // of any other interface (i.e., there are no 'lower-layer-if' entries), and
    // 'disabled' otherwise. The type is LinkUpDownTrapEnable.
    LinkUpDownTrapEnable interface{}

    // policy target for inbound or outbound direction. The type is slice of
    // Interfaces_Interface_DiffservTargetEntry.
    DiffservTargetEntry []Interfaces_Interface_DiffservTargetEntry

    // Parameters for the IPv4 address family.
    Ipv4 Interfaces_Interface_Ipv4

    // Parameters for the IPv6 address family.
    Ipv6 Interfaces_Interface_Ipv6
}

func (self *Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Interfaces_Interface) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "description" { return "Description" }
    if yname == "type" { return "Type" }
    if yname == "enabled" { return "Enabled" }
    if yname == "link-up-down-trap-enable" { return "LinkUpDownTrapEnable" }
    if yname == "ietf-diffserv-target:diffserv-target-entry" { return "DiffservTargetEntry" }
    if yname == "ietf-ip:ipv4" { return "Ipv4" }
    if yname == "ietf-ip:ipv6" { return "Ipv6" }
    return ""
}

func (self *Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[name='" + fmt.Sprintf("%v", self.Name) + "']"
}

func (self *Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ietf-diffserv-target:diffserv-target-entry" {
        for _, c := range self.DiffservTargetEntry {
            if self.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Interfaces_Interface_DiffservTargetEntry{}
        self.DiffservTargetEntry = append(self.DiffservTargetEntry, child)
        return &self.DiffservTargetEntry[len(self.DiffservTargetEntry)-1]
    }
    if childYangName == "ietf-ip:ipv4" {
        return &self.Ipv4
    }
    if childYangName == "ietf-ip:ipv6" {
        return &self.Ipv6
    }
    return nil
}

func (self *Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range self.DiffservTargetEntry {
        children[self.DiffservTargetEntry[i].GetSegmentPath()] = &self.DiffservTargetEntry[i]
    }
    children["ietf-ip:ipv4"] = &self.Ipv4
    children["ietf-ip:ipv6"] = &self.Ipv6
    return children
}

func (self *Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = self.Name
    leafs["description"] = self.Description
    leafs["type"] = self.Type
    leafs["enabled"] = self.Enabled
    leafs["link-up-down-trap-enable"] = self.LinkUpDownTrapEnable
    return leafs
}

func (self *Interfaces_Interface) GetBundleName() string { return "ietf" }

func (self *Interfaces_Interface) GetYangName() string { return "interface" }

func (self *Interfaces_Interface) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (self *Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (self *Interfaces_Interface) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (self *Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// Interfaces_Interface_DiffservTargetEntry
// policy target for inbound or outbound direction
type Interfaces_Interface_DiffservTargetEntry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Direction fo the traffic flow either inbound or
    // outbound. The type is one of the following: InboundOutbound.
    Direction interface{}

    // This attribute is a key. Policy entry name. The type is string.
    PolicyName interface{}
}

func (diffservTargetEntry *Interfaces_Interface_DiffservTargetEntry) GetFilter() yfilter.YFilter { return diffservTargetEntry.YFilter }

func (diffservTargetEntry *Interfaces_Interface_DiffservTargetEntry) SetFilter(yf yfilter.YFilter) { diffservTargetEntry.YFilter = yf }

func (diffservTargetEntry *Interfaces_Interface_DiffservTargetEntry) GetGoName(yname string) string {
    if yname == "direction" { return "Direction" }
    if yname == "policy-name" { return "PolicyName" }
    return ""
}

func (diffservTargetEntry *Interfaces_Interface_DiffservTargetEntry) GetSegmentPath() string {
    return "ietf-diffserv-target:diffserv-target-entry" + "[direction='" + fmt.Sprintf("%v", diffservTargetEntry.Direction) + "']" + "[policy-name='" + fmt.Sprintf("%v", diffservTargetEntry.PolicyName) + "']"
}

func (diffservTargetEntry *Interfaces_Interface_DiffservTargetEntry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (diffservTargetEntry *Interfaces_Interface_DiffservTargetEntry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (diffservTargetEntry *Interfaces_Interface_DiffservTargetEntry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["direction"] = diffservTargetEntry.Direction
    leafs["policy-name"] = diffservTargetEntry.PolicyName
    return leafs
}

func (diffservTargetEntry *Interfaces_Interface_DiffservTargetEntry) GetBundleName() string { return "ietf" }

func (diffservTargetEntry *Interfaces_Interface_DiffservTargetEntry) GetYangName() string { return "diffserv-target-entry" }

func (diffservTargetEntry *Interfaces_Interface_DiffservTargetEntry) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (diffservTargetEntry *Interfaces_Interface_DiffservTargetEntry) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (diffservTargetEntry *Interfaces_Interface_DiffservTargetEntry) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (diffservTargetEntry *Interfaces_Interface_DiffservTargetEntry) SetParent(parent types.Entity) { diffservTargetEntry.parent = parent }

func (diffservTargetEntry *Interfaces_Interface_DiffservTargetEntry) GetParent() types.Entity { return diffservTargetEntry.parent }

func (diffservTargetEntry *Interfaces_Interface_DiffservTargetEntry) GetParentYangName() string { return "interface" }

// Interfaces_Interface_Ipv4
// Parameters for the IPv4 address family.
// This type is a presence type.
type Interfaces_Interface_Ipv4 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Controls whether IPv4 is enabled or disabled on this interface.  When IPv4
    // is enabled, this interface is connected to an IPv4 stack, and the interface
    // can send and receive IPv4 packets. The type is bool. The default value is
    // true.
    Enabled interface{}

    // Controls IPv4 packet forwarding of datagrams received by, but not addressed
    // to, this interface.  IPv4 routers forward datagrams.  IPv4 hosts do not
    // (except those source-routed via the host). The type is bool. The default
    // value is false.
    Forwarding interface{}

    // The size, in octets, of the largest IPv4 packet that the interface will
    // send and receive. The server may restrict the allowed values for this leaf,
    // depending on the interface's type. If this leaf is not configured, the
    // operationally used MTU depends on the interface's type. The type is
    // interface{} with range: 68..65535. Units are octets.
    Mtu interface{}

    // The list of configured IPv4 addresses on the interface. The type is slice
    // of Interfaces_Interface_Ipv4_Address.
    Address []Interfaces_Interface_Ipv4_Address

    // A list of mappings from IPv4 addresses to link-layer addresses. Entries in
    // this list are used as static entries in the ARP Cache. The type is slice of
    // Interfaces_Interface_Ipv4_Neighbor.
    Neighbor []Interfaces_Interface_Ipv4_Neighbor
}

func (ipv4 *Interfaces_Interface_Ipv4) GetFilter() yfilter.YFilter { return ipv4.YFilter }

func (ipv4 *Interfaces_Interface_Ipv4) SetFilter(yf yfilter.YFilter) { ipv4.YFilter = yf }

func (ipv4 *Interfaces_Interface_Ipv4) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    if yname == "forwarding" { return "Forwarding" }
    if yname == "mtu" { return "Mtu" }
    if yname == "address" { return "Address" }
    if yname == "neighbor" { return "Neighbor" }
    return ""
}

func (ipv4 *Interfaces_Interface_Ipv4) GetSegmentPath() string {
    return "ietf-ip:ipv4"
}

func (ipv4 *Interfaces_Interface_Ipv4) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "address" {
        for _, c := range ipv4.Address {
            if ipv4.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Interfaces_Interface_Ipv4_Address{}
        ipv4.Address = append(ipv4.Address, child)
        return &ipv4.Address[len(ipv4.Address)-1]
    }
    if childYangName == "neighbor" {
        for _, c := range ipv4.Neighbor {
            if ipv4.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Interfaces_Interface_Ipv4_Neighbor{}
        ipv4.Neighbor = append(ipv4.Neighbor, child)
        return &ipv4.Neighbor[len(ipv4.Neighbor)-1]
    }
    return nil
}

func (ipv4 *Interfaces_Interface_Ipv4) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ipv4.Address {
        children[ipv4.Address[i].GetSegmentPath()] = &ipv4.Address[i]
    }
    for i := range ipv4.Neighbor {
        children[ipv4.Neighbor[i].GetSegmentPath()] = &ipv4.Neighbor[i]
    }
    return children
}

func (ipv4 *Interfaces_Interface_Ipv4) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = ipv4.Enabled
    leafs["forwarding"] = ipv4.Forwarding
    leafs["mtu"] = ipv4.Mtu
    return leafs
}

func (ipv4 *Interfaces_Interface_Ipv4) GetBundleName() string { return "ietf" }

func (ipv4 *Interfaces_Interface_Ipv4) GetYangName() string { return "ipv4" }

func (ipv4 *Interfaces_Interface_Ipv4) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (ipv4 *Interfaces_Interface_Ipv4) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (ipv4 *Interfaces_Interface_Ipv4) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (ipv4 *Interfaces_Interface_Ipv4) SetParent(parent types.Entity) { ipv4.parent = parent }

func (ipv4 *Interfaces_Interface_Ipv4) GetParent() types.Entity { return ipv4.parent }

func (ipv4 *Interfaces_Interface_Ipv4) GetParentYangName() string { return "interface" }

// Interfaces_Interface_Ipv4_Address
// The list of configured IPv4 addresses on the interface.
type Interfaces_Interface_Ipv4_Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The IPv4 address on the interface. The type is
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ip interface{}

    // The length of the subnet prefix. The type is interface{} with range: 0..32.
    PrefixLength interface{}

    // The subnet specified as a netmask. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]).
    Netmask interface{}
}

func (address *Interfaces_Interface_Ipv4_Address) GetFilter() yfilter.YFilter { return address.YFilter }

func (address *Interfaces_Interface_Ipv4_Address) SetFilter(yf yfilter.YFilter) { address.YFilter = yf }

func (address *Interfaces_Interface_Ipv4_Address) GetGoName(yname string) string {
    if yname == "ip" { return "Ip" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "netmask" { return "Netmask" }
    return ""
}

func (address *Interfaces_Interface_Ipv4_Address) GetSegmentPath() string {
    return "address" + "[ip='" + fmt.Sprintf("%v", address.Ip) + "']"
}

func (address *Interfaces_Interface_Ipv4_Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (address *Interfaces_Interface_Ipv4_Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (address *Interfaces_Interface_Ipv4_Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip"] = address.Ip
    leafs["prefix-length"] = address.PrefixLength
    leafs["netmask"] = address.Netmask
    return leafs
}

func (address *Interfaces_Interface_Ipv4_Address) GetBundleName() string { return "ietf" }

func (address *Interfaces_Interface_Ipv4_Address) GetYangName() string { return "address" }

func (address *Interfaces_Interface_Ipv4_Address) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (address *Interfaces_Interface_Ipv4_Address) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (address *Interfaces_Interface_Ipv4_Address) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (address *Interfaces_Interface_Ipv4_Address) SetParent(parent types.Entity) { address.parent = parent }

func (address *Interfaces_Interface_Ipv4_Address) GetParent() types.Entity { return address.parent }

func (address *Interfaces_Interface_Ipv4_Address) GetParentYangName() string { return "ipv4" }

// Interfaces_Interface_Ipv4_Neighbor
// A list of mappings from IPv4 addresses to
// link-layer addresses.
// Entries in this list are used as static entries in the
// ARP Cache.
type Interfaces_Interface_Ipv4_Neighbor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The IPv4 address of the neighbor node. The type is
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ip interface{}

    // The link-layer address of the neighbor node. The type is string with
    // pattern: ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?. This attribute is mandatory.
    LinkLayerAddress interface{}
}

func (neighbor *Interfaces_Interface_Ipv4_Neighbor) GetFilter() yfilter.YFilter { return neighbor.YFilter }

func (neighbor *Interfaces_Interface_Ipv4_Neighbor) SetFilter(yf yfilter.YFilter) { neighbor.YFilter = yf }

func (neighbor *Interfaces_Interface_Ipv4_Neighbor) GetGoName(yname string) string {
    if yname == "ip" { return "Ip" }
    if yname == "link-layer-address" { return "LinkLayerAddress" }
    return ""
}

func (neighbor *Interfaces_Interface_Ipv4_Neighbor) GetSegmentPath() string {
    return "neighbor" + "[ip='" + fmt.Sprintf("%v", neighbor.Ip) + "']"
}

func (neighbor *Interfaces_Interface_Ipv4_Neighbor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (neighbor *Interfaces_Interface_Ipv4_Neighbor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (neighbor *Interfaces_Interface_Ipv4_Neighbor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip"] = neighbor.Ip
    leafs["link-layer-address"] = neighbor.LinkLayerAddress
    return leafs
}

func (neighbor *Interfaces_Interface_Ipv4_Neighbor) GetBundleName() string { return "ietf" }

func (neighbor *Interfaces_Interface_Ipv4_Neighbor) GetYangName() string { return "neighbor" }

func (neighbor *Interfaces_Interface_Ipv4_Neighbor) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (neighbor *Interfaces_Interface_Ipv4_Neighbor) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (neighbor *Interfaces_Interface_Ipv4_Neighbor) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (neighbor *Interfaces_Interface_Ipv4_Neighbor) SetParent(parent types.Entity) { neighbor.parent = parent }

func (neighbor *Interfaces_Interface_Ipv4_Neighbor) GetParent() types.Entity { return neighbor.parent }

func (neighbor *Interfaces_Interface_Ipv4_Neighbor) GetParentYangName() string { return "ipv4" }

// Interfaces_Interface_Ipv6
// Parameters for the IPv6 address family.
// This type is a presence type.
type Interfaces_Interface_Ipv6 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Controls whether IPv6 is enabled or disabled on this interface.  When IPv6
    // is enabled, this interface is connected to an IPv6 stack, and the interface
    // can send and receive IPv6 packets. The type is bool. The default value is
    // true.
    Enabled interface{}

    // Controls IPv6 packet forwarding of datagrams received by, but not addressed
    // to, this interface.  IPv6 routers forward datagrams.  IPv6 hosts do not
    // (except those source-routed via the host). The type is bool. The default
    // value is false.
    Forwarding interface{}

    // The size, in octets, of the largest IPv6 packet that the interface will
    // send and receive. The server may restrict the allowed values for this leaf,
    // depending on the interface's type. If this leaf is not configured, the
    // operationally used MTU depends on the interface's type. The type is
    // interface{} with range: 1280..4294967295. Units are octets.
    Mtu interface{}

    // The number of consecutive Neighbor Solicitation messages sent while
    // performing Duplicate Address Detection on a tentative address.  A value of
    // zero indicates that Duplicate Address Detection is not performed on
    // tentative addresses.  A value of one indicates a single transmission with
    // no follow-up retransmissions. The type is interface{} with range:
    // 0..4294967295. The default value is 1.
    DupAddrDetectTransmits interface{}

    // The list of configured IPv6 addresses on the interface. The type is slice
    // of Interfaces_Interface_Ipv6_Address.
    Address []Interfaces_Interface_Ipv6_Address

    // A list of mappings from IPv6 addresses to link-layer addresses. Entries in
    // this list are used as static entries in the Neighbor Cache. The type is
    // slice of Interfaces_Interface_Ipv6_Neighbor.
    Neighbor []Interfaces_Interface_Ipv6_Neighbor

    // Parameters to control the autoconfiguration of IPv6 addresses, as described
    // in RFC 4862.
    Autoconf Interfaces_Interface_Ipv6_Autoconf

    // Configuration of IPv6 Router Advertisements.
    Ipv6RouterAdvertisements Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements
}

func (ipv6 *Interfaces_Interface_Ipv6) GetFilter() yfilter.YFilter { return ipv6.YFilter }

func (ipv6 *Interfaces_Interface_Ipv6) SetFilter(yf yfilter.YFilter) { ipv6.YFilter = yf }

func (ipv6 *Interfaces_Interface_Ipv6) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    if yname == "forwarding" { return "Forwarding" }
    if yname == "mtu" { return "Mtu" }
    if yname == "dup-addr-detect-transmits" { return "DupAddrDetectTransmits" }
    if yname == "address" { return "Address" }
    if yname == "neighbor" { return "Neighbor" }
    if yname == "autoconf" { return "Autoconf" }
    if yname == "ietf-ipv6-unicast-routing:ipv6-router-advertisements" { return "Ipv6RouterAdvertisements" }
    return ""
}

func (ipv6 *Interfaces_Interface_Ipv6) GetSegmentPath() string {
    return "ietf-ip:ipv6"
}

func (ipv6 *Interfaces_Interface_Ipv6) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "address" {
        for _, c := range ipv6.Address {
            if ipv6.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Interfaces_Interface_Ipv6_Address{}
        ipv6.Address = append(ipv6.Address, child)
        return &ipv6.Address[len(ipv6.Address)-1]
    }
    if childYangName == "neighbor" {
        for _, c := range ipv6.Neighbor {
            if ipv6.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Interfaces_Interface_Ipv6_Neighbor{}
        ipv6.Neighbor = append(ipv6.Neighbor, child)
        return &ipv6.Neighbor[len(ipv6.Neighbor)-1]
    }
    if childYangName == "autoconf" {
        return &ipv6.Autoconf
    }
    if childYangName == "ietf-ipv6-unicast-routing:ipv6-router-advertisements" {
        return &ipv6.Ipv6RouterAdvertisements
    }
    return nil
}

func (ipv6 *Interfaces_Interface_Ipv6) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ipv6.Address {
        children[ipv6.Address[i].GetSegmentPath()] = &ipv6.Address[i]
    }
    for i := range ipv6.Neighbor {
        children[ipv6.Neighbor[i].GetSegmentPath()] = &ipv6.Neighbor[i]
    }
    children["autoconf"] = &ipv6.Autoconf
    children["ietf-ipv6-unicast-routing:ipv6-router-advertisements"] = &ipv6.Ipv6RouterAdvertisements
    return children
}

func (ipv6 *Interfaces_Interface_Ipv6) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = ipv6.Enabled
    leafs["forwarding"] = ipv6.Forwarding
    leafs["mtu"] = ipv6.Mtu
    leafs["dup-addr-detect-transmits"] = ipv6.DupAddrDetectTransmits
    return leafs
}

func (ipv6 *Interfaces_Interface_Ipv6) GetBundleName() string { return "ietf" }

func (ipv6 *Interfaces_Interface_Ipv6) GetYangName() string { return "ipv6" }

func (ipv6 *Interfaces_Interface_Ipv6) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (ipv6 *Interfaces_Interface_Ipv6) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (ipv6 *Interfaces_Interface_Ipv6) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (ipv6 *Interfaces_Interface_Ipv6) SetParent(parent types.Entity) { ipv6.parent = parent }

func (ipv6 *Interfaces_Interface_Ipv6) GetParent() types.Entity { return ipv6.parent }

func (ipv6 *Interfaces_Interface_Ipv6) GetParentYangName() string { return "interface" }

// Interfaces_Interface_Ipv6_Address
// The list of configured IPv6 addresses on the interface.
type Interfaces_Interface_Ipv6_Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The IPv6 address on the interface. The type is
    // string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ip interface{}

    // The length of the subnet prefix. The type is interface{} with range:
    // 0..128. This attribute is mandatory.
    PrefixLength interface{}
}

func (address *Interfaces_Interface_Ipv6_Address) GetFilter() yfilter.YFilter { return address.YFilter }

func (address *Interfaces_Interface_Ipv6_Address) SetFilter(yf yfilter.YFilter) { address.YFilter = yf }

func (address *Interfaces_Interface_Ipv6_Address) GetGoName(yname string) string {
    if yname == "ip" { return "Ip" }
    if yname == "prefix-length" { return "PrefixLength" }
    return ""
}

func (address *Interfaces_Interface_Ipv6_Address) GetSegmentPath() string {
    return "address" + "[ip='" + fmt.Sprintf("%v", address.Ip) + "']"
}

func (address *Interfaces_Interface_Ipv6_Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (address *Interfaces_Interface_Ipv6_Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (address *Interfaces_Interface_Ipv6_Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip"] = address.Ip
    leafs["prefix-length"] = address.PrefixLength
    return leafs
}

func (address *Interfaces_Interface_Ipv6_Address) GetBundleName() string { return "ietf" }

func (address *Interfaces_Interface_Ipv6_Address) GetYangName() string { return "address" }

func (address *Interfaces_Interface_Ipv6_Address) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (address *Interfaces_Interface_Ipv6_Address) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (address *Interfaces_Interface_Ipv6_Address) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (address *Interfaces_Interface_Ipv6_Address) SetParent(parent types.Entity) { address.parent = parent }

func (address *Interfaces_Interface_Ipv6_Address) GetParent() types.Entity { return address.parent }

func (address *Interfaces_Interface_Ipv6_Address) GetParentYangName() string { return "ipv6" }

// Interfaces_Interface_Ipv6_Neighbor
// A list of mappings from IPv6 addresses to
// link-layer addresses.
// Entries in this list are used as static entries in the
// Neighbor Cache.
type Interfaces_Interface_Ipv6_Neighbor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The IPv6 address of the neighbor node. The type is
    // string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ip interface{}

    // The link-layer address of the neighbor node. The type is string with
    // pattern: ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?. This attribute is mandatory.
    LinkLayerAddress interface{}
}

func (neighbor *Interfaces_Interface_Ipv6_Neighbor) GetFilter() yfilter.YFilter { return neighbor.YFilter }

func (neighbor *Interfaces_Interface_Ipv6_Neighbor) SetFilter(yf yfilter.YFilter) { neighbor.YFilter = yf }

func (neighbor *Interfaces_Interface_Ipv6_Neighbor) GetGoName(yname string) string {
    if yname == "ip" { return "Ip" }
    if yname == "link-layer-address" { return "LinkLayerAddress" }
    return ""
}

func (neighbor *Interfaces_Interface_Ipv6_Neighbor) GetSegmentPath() string {
    return "neighbor" + "[ip='" + fmt.Sprintf("%v", neighbor.Ip) + "']"
}

func (neighbor *Interfaces_Interface_Ipv6_Neighbor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (neighbor *Interfaces_Interface_Ipv6_Neighbor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (neighbor *Interfaces_Interface_Ipv6_Neighbor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip"] = neighbor.Ip
    leafs["link-layer-address"] = neighbor.LinkLayerAddress
    return leafs
}

func (neighbor *Interfaces_Interface_Ipv6_Neighbor) GetBundleName() string { return "ietf" }

func (neighbor *Interfaces_Interface_Ipv6_Neighbor) GetYangName() string { return "neighbor" }

func (neighbor *Interfaces_Interface_Ipv6_Neighbor) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (neighbor *Interfaces_Interface_Ipv6_Neighbor) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (neighbor *Interfaces_Interface_Ipv6_Neighbor) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (neighbor *Interfaces_Interface_Ipv6_Neighbor) SetParent(parent types.Entity) { neighbor.parent = parent }

func (neighbor *Interfaces_Interface_Ipv6_Neighbor) GetParent() types.Entity { return neighbor.parent }

func (neighbor *Interfaces_Interface_Ipv6_Neighbor) GetParentYangName() string { return "ipv6" }

// Interfaces_Interface_Ipv6_Autoconf
// Parameters to control the autoconfiguration of IPv6
// addresses, as described in RFC 4862.
type Interfaces_Interface_Ipv6_Autoconf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // If enabled, the host creates global addresses as described in RFC 4862. The
    // type is bool. The default value is true.
    CreateGlobalAddresses interface{}

    // If enabled, the host creates temporary addresses as described in RFC 4941.
    // The type is bool. The default value is false.
    CreateTemporaryAddresses interface{}

    // The time period during which the temporary address is valid. The type is
    // interface{} with range: 0..4294967295. Units are seconds. The default value
    // is 604800.
    TemporaryValidLifetime interface{}

    // The time period during which the temporary address is preferred. The type
    // is interface{} with range: 0..4294967295. Units are seconds. The default
    // value is 86400.
    TemporaryPreferredLifetime interface{}
}

func (autoconf *Interfaces_Interface_Ipv6_Autoconf) GetFilter() yfilter.YFilter { return autoconf.YFilter }

func (autoconf *Interfaces_Interface_Ipv6_Autoconf) SetFilter(yf yfilter.YFilter) { autoconf.YFilter = yf }

func (autoconf *Interfaces_Interface_Ipv6_Autoconf) GetGoName(yname string) string {
    if yname == "create-global-addresses" { return "CreateGlobalAddresses" }
    if yname == "create-temporary-addresses" { return "CreateTemporaryAddresses" }
    if yname == "temporary-valid-lifetime" { return "TemporaryValidLifetime" }
    if yname == "temporary-preferred-lifetime" { return "TemporaryPreferredLifetime" }
    return ""
}

func (autoconf *Interfaces_Interface_Ipv6_Autoconf) GetSegmentPath() string {
    return "autoconf"
}

func (autoconf *Interfaces_Interface_Ipv6_Autoconf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (autoconf *Interfaces_Interface_Ipv6_Autoconf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (autoconf *Interfaces_Interface_Ipv6_Autoconf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["create-global-addresses"] = autoconf.CreateGlobalAddresses
    leafs["create-temporary-addresses"] = autoconf.CreateTemporaryAddresses
    leafs["temporary-valid-lifetime"] = autoconf.TemporaryValidLifetime
    leafs["temporary-preferred-lifetime"] = autoconf.TemporaryPreferredLifetime
    return leafs
}

func (autoconf *Interfaces_Interface_Ipv6_Autoconf) GetBundleName() string { return "ietf" }

func (autoconf *Interfaces_Interface_Ipv6_Autoconf) GetYangName() string { return "autoconf" }

func (autoconf *Interfaces_Interface_Ipv6_Autoconf) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (autoconf *Interfaces_Interface_Ipv6_Autoconf) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (autoconf *Interfaces_Interface_Ipv6_Autoconf) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (autoconf *Interfaces_Interface_Ipv6_Autoconf) SetParent(parent types.Entity) { autoconf.parent = parent }

func (autoconf *Interfaces_Interface_Ipv6_Autoconf) GetParent() types.Entity { return autoconf.parent }

func (autoconf *Interfaces_Interface_Ipv6_Autoconf) GetParentYangName() string { return "ipv6" }

// Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements
// Configuration of IPv6 Router Advertisements.
type Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A flag indicating whether or not the router sends periodic Router
    // Advertisements and responds to Router Solicitations. The type is bool. The
    // default value is false.
    SendAdvertisements interface{}

    // The maximum time allowed between sending unsolicited multicast Router
    // Advertisements from the interface. The type is interface{} with range:
    // 4..1800. Units are seconds. The default value is 600.
    MaxRtrAdvInterval interface{}

    // The minimum time allowed between sending unsolicited multicast Router
    // Advertisements from the interface.  The default value to be used
    // operationally if this leaf is not configured is determined as follows:  -
    // if max-rtr-adv-interval >= 9 seconds, the default value   is 0.33 *
    // max-rtr-adv-interval;  - otherwise it is 0.75 * max-rtr-adv-interval. The
    // type is interface{} with range: 3..1350. Units are seconds.
    MinRtrAdvInterval interface{}

    // The value to be placed in the 'Managed address configuration' flag field in
    // the Router Advertisement. The type is bool. The default value is false.
    ManagedFlag interface{}

    // The value to be placed in the 'Other configuration' flag field in the
    // Router Advertisement. The type is bool. The default value is false.
    OtherConfigFlag interface{}

    // The value to be placed in MTU options sent by the router. A value of zero
    // indicates that no MTU options are sent. The type is interface{} with range:
    // 0..4294967295. The default value is 0.
    LinkMtu interface{}

    // The value to be placed in the Reachable Time field in the Router
    // Advertisement messages sent by the router. A value of zero means
    // unspecified (by this router). The type is interface{} with range:
    // 0..3600000. Units are milliseconds. The default value is 0.
    ReachableTime interface{}

    // The value to be placed in the Retrans Timer field in the Router
    // Advertisement messages sent by the router. A value of zero means
    // unspecified (by this router). The type is interface{} with range:
    // 0..4294967295. Units are milliseconds. The default value is 0.
    RetransTimer interface{}

    // The value to be placed in the Cur Hop Limit field in the Router
    // Advertisement messages sent by the router. A value of zero means
    // unspecified (by this router).  If this parameter is not configured, the
    // device SHOULD use the value specified in IANA Assigned Numbers that was in
    // effect at the time of implementation. The type is interface{} with range:
    // 0..255.
    CurHopLimit interface{}

    // The value to be placed in the Router Lifetime field of Router
    // Advertisements sent from the interface, in seconds. It MUST be either zero
    // or between max-rtr-adv-interval and 9000 seconds. A value of zero indicates
    // that the router is not to be used as a default router. These limits may be
    // overridden by specific documents that describe how IPv6 operates over
    // different link layers.  If this parameter is not configured, the device
    // SHOULD use a value of 3 * max-rtr-adv-interval. The type is interface{}
    // with range: 0..9000. Units are seconds.
    DefaultLifetime interface{}

    // Configuration of prefixes to be placed in Prefix Information options in
    // Router Advertisement messages sent from the interface.  Prefixes that are
    // advertised by default but do not have their entries in the child 'prefix'
    // list are advertised with the default values of all parameters.  The
    // link-local prefix SHOULD NOT be included in the list of advertised
    // prefixes.
    PrefixList Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements_PrefixList
}

func (ipv6RouterAdvertisements *Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements) GetFilter() yfilter.YFilter { return ipv6RouterAdvertisements.YFilter }

func (ipv6RouterAdvertisements *Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements) SetFilter(yf yfilter.YFilter) { ipv6RouterAdvertisements.YFilter = yf }

func (ipv6RouterAdvertisements *Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements) GetGoName(yname string) string {
    if yname == "send-advertisements" { return "SendAdvertisements" }
    if yname == "max-rtr-adv-interval" { return "MaxRtrAdvInterval" }
    if yname == "min-rtr-adv-interval" { return "MinRtrAdvInterval" }
    if yname == "managed-flag" { return "ManagedFlag" }
    if yname == "other-config-flag" { return "OtherConfigFlag" }
    if yname == "link-mtu" { return "LinkMtu" }
    if yname == "reachable-time" { return "ReachableTime" }
    if yname == "retrans-timer" { return "RetransTimer" }
    if yname == "cur-hop-limit" { return "CurHopLimit" }
    if yname == "default-lifetime" { return "DefaultLifetime" }
    if yname == "prefix-list" { return "PrefixList" }
    return ""
}

func (ipv6RouterAdvertisements *Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements) GetSegmentPath() string {
    return "ietf-ipv6-unicast-routing:ipv6-router-advertisements"
}

func (ipv6RouterAdvertisements *Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-list" {
        return &ipv6RouterAdvertisements.PrefixList
    }
    return nil
}

func (ipv6RouterAdvertisements *Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix-list"] = &ipv6RouterAdvertisements.PrefixList
    return children
}

func (ipv6RouterAdvertisements *Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["send-advertisements"] = ipv6RouterAdvertisements.SendAdvertisements
    leafs["max-rtr-adv-interval"] = ipv6RouterAdvertisements.MaxRtrAdvInterval
    leafs["min-rtr-adv-interval"] = ipv6RouterAdvertisements.MinRtrAdvInterval
    leafs["managed-flag"] = ipv6RouterAdvertisements.ManagedFlag
    leafs["other-config-flag"] = ipv6RouterAdvertisements.OtherConfigFlag
    leafs["link-mtu"] = ipv6RouterAdvertisements.LinkMtu
    leafs["reachable-time"] = ipv6RouterAdvertisements.ReachableTime
    leafs["retrans-timer"] = ipv6RouterAdvertisements.RetransTimer
    leafs["cur-hop-limit"] = ipv6RouterAdvertisements.CurHopLimit
    leafs["default-lifetime"] = ipv6RouterAdvertisements.DefaultLifetime
    return leafs
}

func (ipv6RouterAdvertisements *Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements) GetBundleName() string { return "ietf" }

func (ipv6RouterAdvertisements *Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements) GetYangName() string { return "ipv6-router-advertisements" }

func (ipv6RouterAdvertisements *Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (ipv6RouterAdvertisements *Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (ipv6RouterAdvertisements *Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (ipv6RouterAdvertisements *Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements) SetParent(parent types.Entity) { ipv6RouterAdvertisements.parent = parent }

func (ipv6RouterAdvertisements *Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements) GetParent() types.Entity { return ipv6RouterAdvertisements.parent }

func (ipv6RouterAdvertisements *Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements) GetParentYangName() string { return "ipv6" }

// Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements_PrefixList
// Configuration of prefixes to be placed in Prefix
// Information options in Router Advertisement messages sent
// from the interface.
// 
// Prefixes that are advertised by default but do not have
// their entries in the child 'prefix' list are advertised
// with the default values of all parameters.
// 
// The link-local prefix SHOULD NOT be included in the list
// of advertised prefixes.
type Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements_PrefixList struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration of an advertised prefix entry. The type is slice of
    // Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements_PrefixList_Prefix.
    Prefix []Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements_PrefixList_Prefix
}

func (prefixList *Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements_PrefixList) GetFilter() yfilter.YFilter { return prefixList.YFilter }

func (prefixList *Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements_PrefixList) SetFilter(yf yfilter.YFilter) { prefixList.YFilter = yf }

func (prefixList *Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements_PrefixList) GetGoName(yname string) string {
    if yname == "prefix" { return "Prefix" }
    return ""
}

func (prefixList *Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements_PrefixList) GetSegmentPath() string {
    return "prefix-list"
}

func (prefixList *Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements_PrefixList) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix" {
        for _, c := range prefixList.Prefix {
            if prefixList.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements_PrefixList_Prefix{}
        prefixList.Prefix = append(prefixList.Prefix, child)
        return &prefixList.Prefix[len(prefixList.Prefix)-1]
    }
    return nil
}

func (prefixList *Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements_PrefixList) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range prefixList.Prefix {
        children[prefixList.Prefix[i].GetSegmentPath()] = &prefixList.Prefix[i]
    }
    return children
}

func (prefixList *Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements_PrefixList) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (prefixList *Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements_PrefixList) GetBundleName() string { return "ietf" }

func (prefixList *Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements_PrefixList) GetYangName() string { return "prefix-list" }

func (prefixList *Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements_PrefixList) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (prefixList *Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements_PrefixList) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (prefixList *Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements_PrefixList) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (prefixList *Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements_PrefixList) SetParent(parent types.Entity) { prefixList.parent = parent }

func (prefixList *Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements_PrefixList) GetParent() types.Entity { return prefixList.parent }

func (prefixList *Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements_PrefixList) GetParentYangName() string { return "ipv6-router-advertisements" }

// Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements_PrefixList_Prefix
// Configuration of an advertised prefix entry.
type Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements_PrefixList_Prefix struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IPv6 address prefix. The type is string with
    // pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    PrefixSpec interface{}

    // The prefix will not be advertised.  This can be used for removing the
    // prefix from the default set of advertised prefixes. The type is
    // interface{}.
    NoAdvertise interface{}

    // The value to be placed in the Valid Lifetime in the Prefix Information
    // option. The designated value of all 1's (0xffffffff) represents infinity.
    // The type is interface{} with range: 0..4294967295. Units are seconds. The
    // default value is 2592000.
    ValidLifetime interface{}

    // The value to be placed in the on-link flag ('L-bit') field in the Prefix
    // Information option. The type is bool. The default value is true.
    OnLinkFlag interface{}

    // The value to be placed in the Preferred Lifetime in the Prefix Information
    // option. The designated value of all 1's (0xffffffff) represents infinity.
    // The type is interface{} with range: 0..4294967295. Units are seconds. The
    // default value is 604800.
    PreferredLifetime interface{}

    // The value to be placed in the Autonomous Flag field in the Prefix
    // Information option. The type is bool. The default value is true.
    AutonomousFlag interface{}
}

func (prefix *Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements_PrefixList_Prefix) GetFilter() yfilter.YFilter { return prefix.YFilter }

func (prefix *Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements_PrefixList_Prefix) SetFilter(yf yfilter.YFilter) { prefix.YFilter = yf }

func (prefix *Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements_PrefixList_Prefix) GetGoName(yname string) string {
    if yname == "prefix-spec" { return "PrefixSpec" }
    if yname == "no-advertise" { return "NoAdvertise" }
    if yname == "valid-lifetime" { return "ValidLifetime" }
    if yname == "on-link-flag" { return "OnLinkFlag" }
    if yname == "preferred-lifetime" { return "PreferredLifetime" }
    if yname == "autonomous-flag" { return "AutonomousFlag" }
    return ""
}

func (prefix *Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements_PrefixList_Prefix) GetSegmentPath() string {
    return "prefix" + "[prefix-spec='" + fmt.Sprintf("%v", prefix.PrefixSpec) + "']"
}

func (prefix *Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements_PrefixList_Prefix) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (prefix *Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements_PrefixList_Prefix) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (prefix *Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements_PrefixList_Prefix) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix-spec"] = prefix.PrefixSpec
    leafs["no-advertise"] = prefix.NoAdvertise
    leafs["valid-lifetime"] = prefix.ValidLifetime
    leafs["on-link-flag"] = prefix.OnLinkFlag
    leafs["preferred-lifetime"] = prefix.PreferredLifetime
    leafs["autonomous-flag"] = prefix.AutonomousFlag
    return leafs
}

func (prefix *Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements_PrefixList_Prefix) GetBundleName() string { return "ietf" }

func (prefix *Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements_PrefixList_Prefix) GetYangName() string { return "prefix" }

func (prefix *Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements_PrefixList_Prefix) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (prefix *Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements_PrefixList_Prefix) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (prefix *Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements_PrefixList_Prefix) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (prefix *Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements_PrefixList_Prefix) SetParent(parent types.Entity) { prefix.parent = parent }

func (prefix *Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements_PrefixList_Prefix) GetParent() types.Entity { return prefix.parent }

func (prefix *Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements_PrefixList_Prefix) GetParentYangName() string { return "prefix-list" }

// Interfaces_Interface_LinkUpDownTrapEnable represents no 'lower-layer-if' entries), and 'disabled' otherwise.
type Interfaces_Interface_LinkUpDownTrapEnable string

const (
    Interfaces_Interface_LinkUpDownTrapEnable_enabled Interfaces_Interface_LinkUpDownTrapEnable = "enabled"

    Interfaces_Interface_LinkUpDownTrapEnable_disabled Interfaces_Interface_LinkUpDownTrapEnable = "disabled"
)

// InterfacesState
// Data nodes for the operational state of interfaces.
type InterfacesState struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The list of interfaces on the device.  System-controlled interfaces created
    // by the system are always present in this list, whether they are configured
    // or not. The type is slice of InterfacesState_Interface.
    Interface []InterfacesState_Interface
}

func (interfacesState *InterfacesState) GetFilter() yfilter.YFilter { return interfacesState.YFilter }

func (interfacesState *InterfacesState) SetFilter(yf yfilter.YFilter) { interfacesState.YFilter = yf }

func (interfacesState *InterfacesState) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfacesState *InterfacesState) GetSegmentPath() string {
    return "ietf-interfaces:interfaces-state"
}

func (interfacesState *InterfacesState) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfacesState.Interface {
            if interfacesState.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := InterfacesState_Interface{}
        interfacesState.Interface = append(interfacesState.Interface, child)
        return &interfacesState.Interface[len(interfacesState.Interface)-1]
    }
    return nil
}

func (interfacesState *InterfacesState) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfacesState.Interface {
        children[interfacesState.Interface[i].GetSegmentPath()] = &interfacesState.Interface[i]
    }
    return children
}

func (interfacesState *InterfacesState) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfacesState *InterfacesState) GetBundleName() string { return "ietf" }

func (interfacesState *InterfacesState) GetYangName() string { return "interfaces-state" }

func (interfacesState *InterfacesState) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (interfacesState *InterfacesState) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (interfacesState *InterfacesState) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (interfacesState *InterfacesState) SetParent(parent types.Entity) { interfacesState.parent = parent }

func (interfacesState *InterfacesState) GetParent() types.Entity { return interfacesState.parent }

func (interfacesState *InterfacesState) GetParentYangName() string { return "ietf-interfaces" }

// InterfacesState_Interface
// The list of interfaces on the device.
// 
// System-controlled interfaces created by the system are
// always present in this list, whether they are configured or
// not.
type InterfacesState_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the interface.  A server
    // implementation MAY map this leaf to the ifName MIB object.  Such an
    // implementation needs to use some mechanism to handle the differences in
    // size and characters allowed between this leaf and ifName.  The definition
    // of such a mechanism is outside the scope of this document. The type is
    // string.
    Name interface{}

    // The type of the interface. The type is one of the following:
    // IanaInterfaceTypeVoicefxoAtmvciendptPropbwap2MpPropdocswirelessdownstreamV11SoftwareloopbackHdlcVoicefgdosFastetherfxDvbtdmNfasIfpwtypeL2VlanAdsl2PlusIeee802154VoicefxsDvbrcsmaclayerIdslInfinibandDdnx25Wwanpp2DocscableupstreamEthernet3MbitDigitalpowerlineH323ProxyGtpIpoveratmAlueponImtIpswitchMsdslDvbrccmaclayerSmdsdxiVoiceoveratmArapFastetherMpcLinegroupHippiRprDs1FdlSonetvtVoiceencapSs7SiglinkArcnetActelismetaloopQllcRfc877X25MpegtransportX25MlpVirtualtgHostpadStarlanIso88025DtrIbm370ParchanAdsl2OtnotuPropwirelessp2PInterleaveIsupRegular1822Gr303RdtPropdocswirelessmaclayerAsyncRadiomacOpticalchannelgroupSixtofourPropdocswirelessupstreamQ2931FddiPropcnlsAal2DvbasioutAluelpCiscoislvlanDocscableupstreamrfportAal5FrdlciendptHippiinterfaceL3IpvlanMiox25HssiAtmvirtualAlugpononuRfc1483CnrSipsigMyrinetDlswGigabitethernetX25PleLmpOpticaltransportSdlcVoiceemX86LapsG9982Iso88022LlcDvbasiinBgppolicyaccountingAluepononuMfsiglinkDcnAtmdxiVoiceoverframerelayGfpSonetoverheadchannelVmwarevirtualnicFciplinkIpoverclawCoffeeRadslVdsl2Rs232E1ReachdslVoiceovercableTr008VoiceoveripAtmDs3Ds0Ds1SrpDocscabledownstreamDvbrcstdmaG9983PlcFramerelaympiMvlPropmultiplexorVoicedidCompositelinkProteon10MbitAtmbondFrf16MfrbundleCctemulMplstunnelGponVdslPosIeee8023AdlagDocscablemaclayerDocscablemcmtsdownstreamPppFramerelayEplrsVmwarenicteamCabledownstreamrfportMacsecuncontrolledifIso88023CsmacdUsbAtmfuniTelinkPon622EconetTdlcDs0BundleFastIeee1394CblvectastarRsrbFramerelayinterconnectIsdnsPppmultilinkbundleAflane8025LapbAflane8023LapdIsdnuLapfCapwapwtpvirtualradioIfvfitypeX25HuntgroupParaMacseccontrolledifIso88024TokenbusLocaltalkHyperchannelMediamailoveripIfGsnCapwapdot11ProfileL3IpxvlanAtmsubinterfacePrimaryisdnProteon80MbitIso88026ManDigitalwrapperoverheadchannelDocscableupstreamchannelOpticalchannelEthernetcsmacdBitsTunnelHdsl2FramerelayserviceMplsIeee80211Ieee80212Mocaversion1SonetEsconAlueponlogicallinkG703At2MbUltraDvbrccdownstreamSiptgSmdsicipBridgeAtmlogicalProppointtopointserialV35V36V37IpGr303IdtBasicisdnG703At64KArcnetplusPipDtmSlipHiperlan2AdslIeee80216WmanAtmimaIsdnCapwapdot11BssSipPdnetherloop2VoiceebsIpforwardIso88025CrfpintPropvirtualWwanppOtherPon155QamOtnoduIso88025FiberChannelVoiceemfgdAlugponphysicaluniA12MppswitchIlanPdnetherloop1X213SonetpathVoicefgdeanaIso88025TokenringPropatmAlueponphysicaluniStacktostackFrforwardHomepnaSdslVirtualipaddressBscAtmradioAviciopticaletherG9981FibrechannelShdslEonH323GatekeeperHdh1822DvbrccupstreamNsipTransphdlcTermpadIpovercdlcCesModem.
    // This attribute is mandatory.
    Type interface{}

    // The desired state of the interface.  This leaf has the same read semantics
    // as ifAdminStatus. The type is AdminStatus. This attribute is mandatory.
    AdminStatus interface{}

    // The current operational state of the interface.  This leaf has the same
    // semantics as ifOperStatus. The type is OperStatus. This attribute is
    // mandatory.
    OperStatus interface{}

    // The time the interface entered its current operational state.  If the
    // current state was entered prior to the last re-initialization of the local
    // network management subsystem, then this node is not present. The type is
    // string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    LastChange interface{}

    // The ifIndex value for the ifEntry represented by this interface. The type
    // is interface{} with range: 1..2147483647. This attribute is mandatory.
    IfIndex interface{}

    // The interface's address at its protocol sub-layer.  For example, for an
    // 802.x interface, this object normally contains a Media Access Control (MAC)
    // address.  The interface's media-specific modules must define the bit   and
    // byte ordering and the format of the value of this object.  For interfaces
    // that do not have such an address (e.g., a serial line), this node is not
    // present. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    PhysAddress interface{}

    // A list of references to interfaces layered on top of this interface. The
    // type is slice of string. Refers to
    // interfaces.InterfacesState_Interface_Name
    HigherLayerIf []interface{}

    // A list of references to interfaces layered underneath this interface. The
    // type is slice of string. Refers to
    // interfaces.InterfacesState_Interface_Name
    LowerLayerIf []interface{}

    // An estimate of the interface's current bandwidth in bits per second.  For
    // interfaces that do not vary in bandwidth or for those where no accurate
    // estimation can be made, this node should contain the nominal bandwidth. For
    // interfaces that have no concept of bandwidth, this node is not present. The
    // type is interface{} with range: 0..18446744073709551615. Units are
    // bits/second.
    Speed interface{}

    // The name of the routing instance to which the interface is assigned. The
    // type is string.
    RoutingInstance interface{}

    // A collection of interface-related statistics objects.
    Statistics InterfacesState_Interface_Statistics

    // policy target for inbound or outbound direction. The type is slice of
    // InterfacesState_Interface_DiffservTargetEntry.
    DiffservTargetEntry []InterfacesState_Interface_DiffservTargetEntry

    // Interface-specific parameters for the IPv4 address family.
    Ipv4 InterfacesState_Interface_Ipv4

    // Parameters for the IPv6 address family.
    Ipv6 InterfacesState_Interface_Ipv6
}

func (self *InterfacesState_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *InterfacesState_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *InterfacesState_Interface) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "type" { return "Type" }
    if yname == "admin-status" { return "AdminStatus" }
    if yname == "oper-status" { return "OperStatus" }
    if yname == "last-change" { return "LastChange" }
    if yname == "if-index" { return "IfIndex" }
    if yname == "phys-address" { return "PhysAddress" }
    if yname == "higher-layer-if" { return "HigherLayerIf" }
    if yname == "lower-layer-if" { return "LowerLayerIf" }
    if yname == "speed" { return "Speed" }
    if yname == "routing-instance" { return "RoutingInstance" }
    if yname == "statistics" { return "Statistics" }
    if yname == "ietf-diffserv-target:diffserv-target-entry" { return "DiffservTargetEntry" }
    if yname == "ietf-ip:ipv4" { return "Ipv4" }
    if yname == "ietf-ip:ipv6" { return "Ipv6" }
    return ""
}

func (self *InterfacesState_Interface) GetSegmentPath() string {
    return "interface" + "[name='" + fmt.Sprintf("%v", self.Name) + "']"
}

func (self *InterfacesState_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "statistics" {
        return &self.Statistics
    }
    if childYangName == "ietf-diffserv-target:diffserv-target-entry" {
        for _, c := range self.DiffservTargetEntry {
            if self.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := InterfacesState_Interface_DiffservTargetEntry{}
        self.DiffservTargetEntry = append(self.DiffservTargetEntry, child)
        return &self.DiffservTargetEntry[len(self.DiffservTargetEntry)-1]
    }
    if childYangName == "ietf-ip:ipv4" {
        return &self.Ipv4
    }
    if childYangName == "ietf-ip:ipv6" {
        return &self.Ipv6
    }
    return nil
}

func (self *InterfacesState_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["statistics"] = &self.Statistics
    for i := range self.DiffservTargetEntry {
        children[self.DiffservTargetEntry[i].GetSegmentPath()] = &self.DiffservTargetEntry[i]
    }
    children["ietf-ip:ipv4"] = &self.Ipv4
    children["ietf-ip:ipv6"] = &self.Ipv6
    return children
}

func (self *InterfacesState_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = self.Name
    leafs["type"] = self.Type
    leafs["admin-status"] = self.AdminStatus
    leafs["oper-status"] = self.OperStatus
    leafs["last-change"] = self.LastChange
    leafs["if-index"] = self.IfIndex
    leafs["phys-address"] = self.PhysAddress
    leafs["higher-layer-if"] = self.HigherLayerIf
    leafs["lower-layer-if"] = self.LowerLayerIf
    leafs["speed"] = self.Speed
    leafs["routing-instance"] = self.RoutingInstance
    return leafs
}

func (self *InterfacesState_Interface) GetBundleName() string { return "ietf" }

func (self *InterfacesState_Interface) GetYangName() string { return "interface" }

func (self *InterfacesState_Interface) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (self *InterfacesState_Interface) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (self *InterfacesState_Interface) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (self *InterfacesState_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *InterfacesState_Interface) GetParent() types.Entity { return self.parent }

func (self *InterfacesState_Interface) GetParentYangName() string { return "interfaces-state" }

// InterfacesState_Interface_Statistics
// A collection of interface-related statistics objects.
type InterfacesState_Interface_Statistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The time on the most recent occasion at which any one or more of this
    // interface's counters suffered a discontinuity.  If no such discontinuities
    // have occurred since the last re-initialization of the local management
    // subsystem, then this node contains the time the local management subsystem
    // re-initialized itself. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}). This
    // attribute is mandatory.
    DiscontinuityTime interface{}

    // The total number of octets received on the interface, including framing
    // characters.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of 'discontinuity-time'. The type is interface{} with range:
    // 0..18446744073709551615.
    InOctets interface{}

    // The number of packets, delivered by this sub-layer to a higher (sub-)layer,
    // that were not addressed to a multicast or broadcast address at this
    // sub-layer.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of 'discontinuity-time'. The type is interface{} with range:
    // 0..18446744073709551615.
    InUnicastPkts interface{}

    // The number of packets, delivered by this sub-layer to a higher (sub-)layer,
    // that were addressed to a broadcast address at this sub-layer. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // 'discontinuity-time'. The type is interface{} with range:
    // 0..18446744073709551615.
    InBroadcastPkts interface{}

    // The number of packets, delivered by this sub-layer to a higher (sub-)layer,
    // that were addressed to a multicast address at this sub-layer.  For a
    // MAC-layer protocol, this includes both Group and Functional addresses. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // 'discontinuity-time'. The type is interface{} with range:
    // 0..18446744073709551615.
    InMulticastPkts interface{}

    // The number of inbound packets that were chosen to be discarded even though
    // no errors had been detected to prevent their being deliverable to a
    // higher-layer protocol.  One possible reason for discarding such a packet
    // could be to free up buffer space.  Discontinuities in the value of this
    // counter can occur at re-initialization of the management system, and at
    // other times as indicated by the value of 'discontinuity-time'. The type is
    // interface{} with range: 0..4294967295.
    InDiscards interface{}

    // For packet-oriented interfaces, the number of inbound packets that
    // contained errors preventing them from being deliverable to a higher-layer
    // protocol.  For character- oriented or fixed-length interfaces, the number
    // of inbound transmission units that contained errors preventing them from
    // being deliverable to a higher-layer protocol.  Discontinuities in the value
    // of this counter can occur at re-initialization of the management system,
    // and at other times as indicated by the value of 'discontinuity-time'. The
    // type is interface{} with range: 0..4294967295.
    InErrors interface{}

    // For packet-oriented interfaces, the number of packets received via the
    // interface that were discarded because of an unknown or unsupported
    // protocol.  For character-oriented or fixed-length interfaces that support
    // protocol multiplexing, the number of transmission units received via the
    // interface that were discarded because of an unknown or unsupported
    // protocol. For any interface that does not support protocol multiplexing,
    // this counter is not present.  Discontinuities in the value of this counter
    // can occur at re-initialization of the management system, and at other times
    // as indicated by the value of 'discontinuity-time'. The type is interface{}
    // with range: 0..4294967295.
    InUnknownProtos interface{}

    // The total number of octets transmitted out of the interface, including
    // framing characters.  Discontinuities in the value of this counter can occur
    // at re-initialization of the management system, and at other times as
    // indicated by the value of 'discontinuity-time'. The type is interface{}
    // with range: 0..18446744073709551615.
    OutOctets interface{}

    // The total number of packets that higher-level protocols requested be
    // transmitted, and that were not addressed to a multicast or broadcast
    // address at this sub-layer, including those that were discarded or not sent.
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // 'discontinuity-time'. The type is interface{} with range:
    // 0..18446744073709551615.
    OutUnicastPkts interface{}

    // The total number of packets that higher-level protocols requested be
    // transmitted, and that were addressed to a broadcast address at this
    // sub-layer, including those that were discarded or not sent. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // 'discontinuity-time'. The type is interface{} with range:
    // 0..18446744073709551615.
    OutBroadcastPkts interface{}

    // The total number of packets that higher-level protocols requested be
    // transmitted, and that were addressed to a multicast address at this
    // sub-layer, including those that were discarded or not sent.  For a
    // MAC-layer protocol, this includes both Group and Functional addresses. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // 'discontinuity-time'. The type is interface{} with range:
    // 0..18446744073709551615.
    OutMulticastPkts interface{}

    // The number of outbound packets that were chosen to be discarded even though
    // no errors had been detected to prevent their being transmitted.  One
    // possible reason for discarding such a packet could be to free up buffer
    // space.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of 'discontinuity-time'. The type is interface{} with range:
    // 0..4294967295.
    OutDiscards interface{}

    // For packet-oriented interfaces, the number of outbound packets that could
    // not be transmitted because of errors. For character-oriented or
    // fixed-length interfaces, the number of outbound transmission units that
    // could not be transmitted because of errors.     Discontinuities in the
    // value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of
    // 'discontinuity-time'. The type is interface{} with range: 0..4294967295.
    OutErrors interface{}

    // total packets input. The type is interface{} with range:
    // 0..18446744073709551615.
    InPkts interface{}

    // total packets output. The type is interface{} with range:
    // 0..18446744073709551615.
    OutPkts interface{}
}

func (statistics *InterfacesState_Interface_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *InterfacesState_Interface_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *InterfacesState_Interface_Statistics) GetGoName(yname string) string {
    if yname == "discontinuity-time" { return "DiscontinuityTime" }
    if yname == "in-octets" { return "InOctets" }
    if yname == "in-unicast-pkts" { return "InUnicastPkts" }
    if yname == "in-broadcast-pkts" { return "InBroadcastPkts" }
    if yname == "in-multicast-pkts" { return "InMulticastPkts" }
    if yname == "in-discards" { return "InDiscards" }
    if yname == "in-errors" { return "InErrors" }
    if yname == "in-unknown-protos" { return "InUnknownProtos" }
    if yname == "out-octets" { return "OutOctets" }
    if yname == "out-unicast-pkts" { return "OutUnicastPkts" }
    if yname == "out-broadcast-pkts" { return "OutBroadcastPkts" }
    if yname == "out-multicast-pkts" { return "OutMulticastPkts" }
    if yname == "out-discards" { return "OutDiscards" }
    if yname == "out-errors" { return "OutErrors" }
    if yname == "in-pkts" { return "InPkts" }
    if yname == "out-pkts" { return "OutPkts" }
    return ""
}

func (statistics *InterfacesState_Interface_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *InterfacesState_Interface_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (statistics *InterfacesState_Interface_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (statistics *InterfacesState_Interface_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["discontinuity-time"] = statistics.DiscontinuityTime
    leafs["in-octets"] = statistics.InOctets
    leafs["in-unicast-pkts"] = statistics.InUnicastPkts
    leafs["in-broadcast-pkts"] = statistics.InBroadcastPkts
    leafs["in-multicast-pkts"] = statistics.InMulticastPkts
    leafs["in-discards"] = statistics.InDiscards
    leafs["in-errors"] = statistics.InErrors
    leafs["in-unknown-protos"] = statistics.InUnknownProtos
    leafs["out-octets"] = statistics.OutOctets
    leafs["out-unicast-pkts"] = statistics.OutUnicastPkts
    leafs["out-broadcast-pkts"] = statistics.OutBroadcastPkts
    leafs["out-multicast-pkts"] = statistics.OutMulticastPkts
    leafs["out-discards"] = statistics.OutDiscards
    leafs["out-errors"] = statistics.OutErrors
    leafs["in-pkts"] = statistics.InPkts
    leafs["out-pkts"] = statistics.OutPkts
    return leafs
}

func (statistics *InterfacesState_Interface_Statistics) GetBundleName() string { return "ietf" }

func (statistics *InterfacesState_Interface_Statistics) GetYangName() string { return "statistics" }

func (statistics *InterfacesState_Interface_Statistics) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (statistics *InterfacesState_Interface_Statistics) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (statistics *InterfacesState_Interface_Statistics) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (statistics *InterfacesState_Interface_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *InterfacesState_Interface_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *InterfacesState_Interface_Statistics) GetParentYangName() string { return "interface" }

// InterfacesState_Interface_DiffservTargetEntry
// policy target for inbound or outbound direction
type InterfacesState_Interface_DiffservTargetEntry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Direction fo the traffic flow either inbound or
    // outbound. The type is one of the following: InboundOutbound.
    Direction interface{}

    // This attribute is a key. Policy entry name. The type is string.
    PolicyName interface{}

    // Statistics for each Classifier Entry in a Policy. The type is slice of
    // InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics.
    DiffservTargetClassifierStatistics []InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics
}

func (diffservTargetEntry *InterfacesState_Interface_DiffservTargetEntry) GetFilter() yfilter.YFilter { return diffservTargetEntry.YFilter }

func (diffservTargetEntry *InterfacesState_Interface_DiffservTargetEntry) SetFilter(yf yfilter.YFilter) { diffservTargetEntry.YFilter = yf }

func (diffservTargetEntry *InterfacesState_Interface_DiffservTargetEntry) GetGoName(yname string) string {
    if yname == "direction" { return "Direction" }
    if yname == "policy-name" { return "PolicyName" }
    if yname == "diffserv-target-classifier-statistics" { return "DiffservTargetClassifierStatistics" }
    return ""
}

func (diffservTargetEntry *InterfacesState_Interface_DiffservTargetEntry) GetSegmentPath() string {
    return "ietf-diffserv-target:diffserv-target-entry" + "[direction='" + fmt.Sprintf("%v", diffservTargetEntry.Direction) + "']" + "[policy-name='" + fmt.Sprintf("%v", diffservTargetEntry.PolicyName) + "']"
}

func (diffservTargetEntry *InterfacesState_Interface_DiffservTargetEntry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "diffserv-target-classifier-statistics" {
        for _, c := range diffservTargetEntry.DiffservTargetClassifierStatistics {
            if diffservTargetEntry.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics{}
        diffservTargetEntry.DiffservTargetClassifierStatistics = append(diffservTargetEntry.DiffservTargetClassifierStatistics, child)
        return &diffservTargetEntry.DiffservTargetClassifierStatistics[len(diffservTargetEntry.DiffservTargetClassifierStatistics)-1]
    }
    return nil
}

func (diffservTargetEntry *InterfacesState_Interface_DiffservTargetEntry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range diffservTargetEntry.DiffservTargetClassifierStatistics {
        children[diffservTargetEntry.DiffservTargetClassifierStatistics[i].GetSegmentPath()] = &diffservTargetEntry.DiffservTargetClassifierStatistics[i]
    }
    return children
}

func (diffservTargetEntry *InterfacesState_Interface_DiffservTargetEntry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["direction"] = diffservTargetEntry.Direction
    leafs["policy-name"] = diffservTargetEntry.PolicyName
    return leafs
}

func (diffservTargetEntry *InterfacesState_Interface_DiffservTargetEntry) GetBundleName() string { return "ietf" }

func (diffservTargetEntry *InterfacesState_Interface_DiffservTargetEntry) GetYangName() string { return "diffserv-target-entry" }

func (diffservTargetEntry *InterfacesState_Interface_DiffservTargetEntry) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (diffservTargetEntry *InterfacesState_Interface_DiffservTargetEntry) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (diffservTargetEntry *InterfacesState_Interface_DiffservTargetEntry) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (diffservTargetEntry *InterfacesState_Interface_DiffservTargetEntry) SetParent(parent types.Entity) { diffservTargetEntry.parent = parent }

func (diffservTargetEntry *InterfacesState_Interface_DiffservTargetEntry) GetParent() types.Entity { return diffservTargetEntry.parent }

func (diffservTargetEntry *InterfacesState_Interface_DiffservTargetEntry) GetParentYangName() string { return "interface" }

// InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics
// Statistics for each Classifier Entry in a Policy
type InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Classifier Entry Name. The type is string.
    ClassifierEntryName interface{}

    // This attribute is a key. Path of the Classifier Entry in a hierarchial
    // policy . The type is string.
    ParentPath interface{}

    // This group defines the classifier filter statistics of  each classifier
    // entry         .
    ClassifierEntryStatistics InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_ClassifierEntryStatistics

    // Meter statistics. The type is slice of
    // InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_MeterStatistics.
    MeterStatistics []InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_MeterStatistics

    // queue related statistics .
    QueuingStatistics InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics
}

func (diffservTargetClassifierStatistics *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics) GetFilter() yfilter.YFilter { return diffservTargetClassifierStatistics.YFilter }

func (diffservTargetClassifierStatistics *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics) SetFilter(yf yfilter.YFilter) { diffservTargetClassifierStatistics.YFilter = yf }

func (diffservTargetClassifierStatistics *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics) GetGoName(yname string) string {
    if yname == "classifier-entry-name" { return "ClassifierEntryName" }
    if yname == "parent-path" { return "ParentPath" }
    if yname == "classifier-entry-statistics" { return "ClassifierEntryStatistics" }
    if yname == "meter-statistics" { return "MeterStatistics" }
    if yname == "queuing-statistics" { return "QueuingStatistics" }
    return ""
}

func (diffservTargetClassifierStatistics *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics) GetSegmentPath() string {
    return "diffserv-target-classifier-statistics" + "[classifier-entry-name='" + fmt.Sprintf("%v", diffservTargetClassifierStatistics.ClassifierEntryName) + "']" + "[parent-path='" + fmt.Sprintf("%v", diffservTargetClassifierStatistics.ParentPath) + "']"
}

func (diffservTargetClassifierStatistics *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "classifier-entry-statistics" {
        return &diffservTargetClassifierStatistics.ClassifierEntryStatistics
    }
    if childYangName == "meter-statistics" {
        for _, c := range diffservTargetClassifierStatistics.MeterStatistics {
            if diffservTargetClassifierStatistics.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_MeterStatistics{}
        diffservTargetClassifierStatistics.MeterStatistics = append(diffservTargetClassifierStatistics.MeterStatistics, child)
        return &diffservTargetClassifierStatistics.MeterStatistics[len(diffservTargetClassifierStatistics.MeterStatistics)-1]
    }
    if childYangName == "queuing-statistics" {
        return &diffservTargetClassifierStatistics.QueuingStatistics
    }
    return nil
}

func (diffservTargetClassifierStatistics *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["classifier-entry-statistics"] = &diffservTargetClassifierStatistics.ClassifierEntryStatistics
    for i := range diffservTargetClassifierStatistics.MeterStatistics {
        children[diffservTargetClassifierStatistics.MeterStatistics[i].GetSegmentPath()] = &diffservTargetClassifierStatistics.MeterStatistics[i]
    }
    children["queuing-statistics"] = &diffservTargetClassifierStatistics.QueuingStatistics
    return children
}

func (diffservTargetClassifierStatistics *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["classifier-entry-name"] = diffservTargetClassifierStatistics.ClassifierEntryName
    leafs["parent-path"] = diffservTargetClassifierStatistics.ParentPath
    return leafs
}

func (diffservTargetClassifierStatistics *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics) GetBundleName() string { return "ietf" }

func (diffservTargetClassifierStatistics *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics) GetYangName() string { return "diffserv-target-classifier-statistics" }

func (diffservTargetClassifierStatistics *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (diffservTargetClassifierStatistics *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (diffservTargetClassifierStatistics *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (diffservTargetClassifierStatistics *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics) SetParent(parent types.Entity) { diffservTargetClassifierStatistics.parent = parent }

func (diffservTargetClassifierStatistics *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics) GetParent() types.Entity { return diffservTargetClassifierStatistics.parent }

func (diffservTargetClassifierStatistics *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics) GetParentYangName() string { return "diffserv-target-entry" }

// InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_ClassifierEntryStatistics
// 
// This group defines the classifier filter statistics of 
// each classifier entry
//        
// 
type InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_ClassifierEntryStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of total packets which filtered  to the classifier-entry. The type
    // is interface{} with range: 0..18446744073709551615.
    ClassifiedPkts interface{}

    // Number of total bytes which filtered   to the classifier-entry. The type is
    // interface{} with range: 0..18446744073709551615.
    ClassifiedBytes interface{}

    // Rate of average data flow through the   classifier-entry. The type is
    // interface{} with range: 0..18446744073709551615. Units are bits-per-second.
    ClassifiedRate interface{}
}

func (classifierEntryStatistics *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_ClassifierEntryStatistics) GetFilter() yfilter.YFilter { return classifierEntryStatistics.YFilter }

func (classifierEntryStatistics *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_ClassifierEntryStatistics) SetFilter(yf yfilter.YFilter) { classifierEntryStatistics.YFilter = yf }

func (classifierEntryStatistics *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_ClassifierEntryStatistics) GetGoName(yname string) string {
    if yname == "classified-pkts" { return "ClassifiedPkts" }
    if yname == "classified-bytes" { return "ClassifiedBytes" }
    if yname == "classified-rate" { return "ClassifiedRate" }
    return ""
}

func (classifierEntryStatistics *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_ClassifierEntryStatistics) GetSegmentPath() string {
    return "classifier-entry-statistics"
}

func (classifierEntryStatistics *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_ClassifierEntryStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (classifierEntryStatistics *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_ClassifierEntryStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (classifierEntryStatistics *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_ClassifierEntryStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["classified-pkts"] = classifierEntryStatistics.ClassifiedPkts
    leafs["classified-bytes"] = classifierEntryStatistics.ClassifiedBytes
    leafs["classified-rate"] = classifierEntryStatistics.ClassifiedRate
    return leafs
}

func (classifierEntryStatistics *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_ClassifierEntryStatistics) GetBundleName() string { return "ietf" }

func (classifierEntryStatistics *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_ClassifierEntryStatistics) GetYangName() string { return "classifier-entry-statistics" }

func (classifierEntryStatistics *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_ClassifierEntryStatistics) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (classifierEntryStatistics *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_ClassifierEntryStatistics) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (classifierEntryStatistics *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_ClassifierEntryStatistics) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (classifierEntryStatistics *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_ClassifierEntryStatistics) SetParent(parent types.Entity) { classifierEntryStatistics.parent = parent }

func (classifierEntryStatistics *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_ClassifierEntryStatistics) GetParent() types.Entity { return classifierEntryStatistics.parent }

func (classifierEntryStatistics *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_ClassifierEntryStatistics) GetParentYangName() string { return "diffserv-target-classifier-statistics" }

// InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_MeterStatistics
// Meter statistics
type InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_MeterStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Meter Identifier. The type is interface{} with
    // range: 0..65535.
    MeterId interface{}

    // Number of packets which succeed the meter. The type is interface{} with
    // range: 0..18446744073709551615.
    MeterSucceedPkts interface{}

    // Bytes of packets which succeed the meter. The type is interface{} with
    // range: 0..18446744073709551615.
    MeterSucceedBytes interface{}

    // Number of packets which failed the meter. The type is interface{} with
    // range: 0..18446744073709551615.
    MeterFailedPkts interface{}

    // Bytes of packets which failed the meter. The type is interface{} with
    // range: 0..18446744073709551615.
    MeterFailedBytes interface{}
}

func (meterStatistics *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_MeterStatistics) GetFilter() yfilter.YFilter { return meterStatistics.YFilter }

func (meterStatistics *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_MeterStatistics) SetFilter(yf yfilter.YFilter) { meterStatistics.YFilter = yf }

func (meterStatistics *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_MeterStatistics) GetGoName(yname string) string {
    if yname == "meter-id" { return "MeterId" }
    if yname == "meter-succeed-pkts" { return "MeterSucceedPkts" }
    if yname == "meter-succeed-bytes" { return "MeterSucceedBytes" }
    if yname == "meter-failed-pkts" { return "MeterFailedPkts" }
    if yname == "meter-failed-bytes" { return "MeterFailedBytes" }
    return ""
}

func (meterStatistics *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_MeterStatistics) GetSegmentPath() string {
    return "meter-statistics" + "[meter-id='" + fmt.Sprintf("%v", meterStatistics.MeterId) + "']"
}

func (meterStatistics *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_MeterStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (meterStatistics *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_MeterStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (meterStatistics *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_MeterStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["meter-id"] = meterStatistics.MeterId
    leafs["meter-succeed-pkts"] = meterStatistics.MeterSucceedPkts
    leafs["meter-succeed-bytes"] = meterStatistics.MeterSucceedBytes
    leafs["meter-failed-pkts"] = meterStatistics.MeterFailedPkts
    leafs["meter-failed-bytes"] = meterStatistics.MeterFailedBytes
    return leafs
}

func (meterStatistics *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_MeterStatistics) GetBundleName() string { return "ietf" }

func (meterStatistics *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_MeterStatistics) GetYangName() string { return "meter-statistics" }

func (meterStatistics *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_MeterStatistics) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (meterStatistics *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_MeterStatistics) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (meterStatistics *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_MeterStatistics) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (meterStatistics *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_MeterStatistics) SetParent(parent types.Entity) { meterStatistics.parent = parent }

func (meterStatistics *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_MeterStatistics) GetParent() types.Entity { return meterStatistics.parent }

func (meterStatistics *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_MeterStatistics) GetParentYangName() string { return "diffserv-target-classifier-statistics" }

// InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics
// queue related statistics 
type InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of packets transmitted from queue . The type is interface{} with
    // range: 0..18446744073709551615.
    OutputPkts interface{}

    // Number of bytes transmitted from queue . The type is interface{} with
    // range: 0..18446744073709551615.
    OutputBytes interface{}

    // Number of packets currently buffered . The type is interface{} with range:
    // 0..18446744073709551615.
    QueueSizePkts interface{}

    // Number of bytes currently buffered . The type is interface{} with range:
    // 0..18446744073709551615.
    QueueSizeBytes interface{}

    // Total number of packets dropped . The type is interface{} with range:
    // 0..18446744073709551615.
    DropPkts interface{}

    // Total number of bytes dropped . The type is interface{} with range:
    // 0..18446744073709551615.
    DropBytes interface{}

    // Container for WRED statistics.
    WredStats InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics_WredStats
}

func (queuingStatistics *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics) GetFilter() yfilter.YFilter { return queuingStatistics.YFilter }

func (queuingStatistics *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics) SetFilter(yf yfilter.YFilter) { queuingStatistics.YFilter = yf }

func (queuingStatistics *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics) GetGoName(yname string) string {
    if yname == "output-pkts" { return "OutputPkts" }
    if yname == "output-bytes" { return "OutputBytes" }
    if yname == "queue-size-pkts" { return "QueueSizePkts" }
    if yname == "queue-size-bytes" { return "QueueSizeBytes" }
    if yname == "drop-pkts" { return "DropPkts" }
    if yname == "drop-bytes" { return "DropBytes" }
    if yname == "wred-stats" { return "WredStats" }
    return ""
}

func (queuingStatistics *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics) GetSegmentPath() string {
    return "queuing-statistics"
}

func (queuingStatistics *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "wred-stats" {
        return &queuingStatistics.WredStats
    }
    return nil
}

func (queuingStatistics *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["wred-stats"] = &queuingStatistics.WredStats
    return children
}

func (queuingStatistics *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["output-pkts"] = queuingStatistics.OutputPkts
    leafs["output-bytes"] = queuingStatistics.OutputBytes
    leafs["queue-size-pkts"] = queuingStatistics.QueueSizePkts
    leafs["queue-size-bytes"] = queuingStatistics.QueueSizeBytes
    leafs["drop-pkts"] = queuingStatistics.DropPkts
    leafs["drop-bytes"] = queuingStatistics.DropBytes
    return leafs
}

func (queuingStatistics *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics) GetBundleName() string { return "ietf" }

func (queuingStatistics *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics) GetYangName() string { return "queuing-statistics" }

func (queuingStatistics *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (queuingStatistics *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (queuingStatistics *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (queuingStatistics *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics) SetParent(parent types.Entity) { queuingStatistics.parent = parent }

func (queuingStatistics *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics) GetParent() types.Entity { return queuingStatistics.parent }

func (queuingStatistics *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics) GetParentYangName() string { return "diffserv-target-classifier-statistics" }

// InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics_WredStats
// Container for WRED statistics
type InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics_WredStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Early drop packets . The type is interface{} with range:
    // 0..18446744073709551615.
    EarlyDropPkts interface{}

    // Early drop bytes . The type is interface{} with range:
    // 0..18446744073709551615.
    EarlyDropBytes interface{}
}

func (wredStats *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics_WredStats) GetFilter() yfilter.YFilter { return wredStats.YFilter }

func (wredStats *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics_WredStats) SetFilter(yf yfilter.YFilter) { wredStats.YFilter = yf }

func (wredStats *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics_WredStats) GetGoName(yname string) string {
    if yname == "early-drop-pkts" { return "EarlyDropPkts" }
    if yname == "early-drop-bytes" { return "EarlyDropBytes" }
    return ""
}

func (wredStats *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics_WredStats) GetSegmentPath() string {
    return "wred-stats"
}

func (wredStats *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics_WredStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (wredStats *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics_WredStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (wredStats *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics_WredStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["early-drop-pkts"] = wredStats.EarlyDropPkts
    leafs["early-drop-bytes"] = wredStats.EarlyDropBytes
    return leafs
}

func (wredStats *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics_WredStats) GetBundleName() string { return "ietf" }

func (wredStats *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics_WredStats) GetYangName() string { return "wred-stats" }

func (wredStats *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics_WredStats) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (wredStats *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics_WredStats) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (wredStats *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics_WredStats) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (wredStats *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics_WredStats) SetParent(parent types.Entity) { wredStats.parent = parent }

func (wredStats *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics_WredStats) GetParent() types.Entity { return wredStats.parent }

func (wredStats *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics_WredStats) GetParentYangName() string { return "queuing-statistics" }

// InterfacesState_Interface_Ipv4
// Interface-specific parameters for the IPv4 address family.
// This type is a presence type.
type InterfacesState_Interface_Ipv4 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Indicates whether IPv4 packet forwarding is enabled or disabled on this
    // interface. The type is bool.
    Forwarding interface{}

    // The size, in octets, of the largest IPv4 packet that the interface will
    // send and receive. The type is interface{} with range: 68..65535. Units are
    // octets.
    Mtu interface{}

    // The list of IPv4 addresses on the interface. The type is slice of
    // InterfacesState_Interface_Ipv4_Address.
    Address []InterfacesState_Interface_Ipv4_Address

    // A list of mappings from IPv4 addresses to link-layer addresses. This list
    // represents the ARP Cache. The type is slice of
    // InterfacesState_Interface_Ipv4_Neighbor.
    Neighbor []InterfacesState_Interface_Ipv4_Neighbor
}

func (ipv4 *InterfacesState_Interface_Ipv4) GetFilter() yfilter.YFilter { return ipv4.YFilter }

func (ipv4 *InterfacesState_Interface_Ipv4) SetFilter(yf yfilter.YFilter) { ipv4.YFilter = yf }

func (ipv4 *InterfacesState_Interface_Ipv4) GetGoName(yname string) string {
    if yname == "forwarding" { return "Forwarding" }
    if yname == "mtu" { return "Mtu" }
    if yname == "address" { return "Address" }
    if yname == "neighbor" { return "Neighbor" }
    return ""
}

func (ipv4 *InterfacesState_Interface_Ipv4) GetSegmentPath() string {
    return "ietf-ip:ipv4"
}

func (ipv4 *InterfacesState_Interface_Ipv4) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "address" {
        for _, c := range ipv4.Address {
            if ipv4.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := InterfacesState_Interface_Ipv4_Address{}
        ipv4.Address = append(ipv4.Address, child)
        return &ipv4.Address[len(ipv4.Address)-1]
    }
    if childYangName == "neighbor" {
        for _, c := range ipv4.Neighbor {
            if ipv4.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := InterfacesState_Interface_Ipv4_Neighbor{}
        ipv4.Neighbor = append(ipv4.Neighbor, child)
        return &ipv4.Neighbor[len(ipv4.Neighbor)-1]
    }
    return nil
}

func (ipv4 *InterfacesState_Interface_Ipv4) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ipv4.Address {
        children[ipv4.Address[i].GetSegmentPath()] = &ipv4.Address[i]
    }
    for i := range ipv4.Neighbor {
        children[ipv4.Neighbor[i].GetSegmentPath()] = &ipv4.Neighbor[i]
    }
    return children
}

func (ipv4 *InterfacesState_Interface_Ipv4) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["forwarding"] = ipv4.Forwarding
    leafs["mtu"] = ipv4.Mtu
    return leafs
}

func (ipv4 *InterfacesState_Interface_Ipv4) GetBundleName() string { return "ietf" }

func (ipv4 *InterfacesState_Interface_Ipv4) GetYangName() string { return "ipv4" }

func (ipv4 *InterfacesState_Interface_Ipv4) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (ipv4 *InterfacesState_Interface_Ipv4) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (ipv4 *InterfacesState_Interface_Ipv4) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (ipv4 *InterfacesState_Interface_Ipv4) SetParent(parent types.Entity) { ipv4.parent = parent }

func (ipv4 *InterfacesState_Interface_Ipv4) GetParent() types.Entity { return ipv4.parent }

func (ipv4 *InterfacesState_Interface_Ipv4) GetParentYangName() string { return "interface" }

// InterfacesState_Interface_Ipv4_Address
// The list of IPv4 addresses on the interface.
type InterfacesState_Interface_Ipv4_Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The IPv4 address on the interface. The type is
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ip interface{}

    // The length of the subnet prefix. The type is interface{} with range: 0..32.
    PrefixLength interface{}

    // The subnet specified as a netmask. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]).
    Netmask interface{}

    // The origin of this address. The type is IpAddressOrigin.
    Origin interface{}
}

func (address *InterfacesState_Interface_Ipv4_Address) GetFilter() yfilter.YFilter { return address.YFilter }

func (address *InterfacesState_Interface_Ipv4_Address) SetFilter(yf yfilter.YFilter) { address.YFilter = yf }

func (address *InterfacesState_Interface_Ipv4_Address) GetGoName(yname string) string {
    if yname == "ip" { return "Ip" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "netmask" { return "Netmask" }
    if yname == "origin" { return "Origin" }
    return ""
}

func (address *InterfacesState_Interface_Ipv4_Address) GetSegmentPath() string {
    return "address" + "[ip='" + fmt.Sprintf("%v", address.Ip) + "']"
}

func (address *InterfacesState_Interface_Ipv4_Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (address *InterfacesState_Interface_Ipv4_Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (address *InterfacesState_Interface_Ipv4_Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip"] = address.Ip
    leafs["prefix-length"] = address.PrefixLength
    leafs["netmask"] = address.Netmask
    leafs["origin"] = address.Origin
    return leafs
}

func (address *InterfacesState_Interface_Ipv4_Address) GetBundleName() string { return "ietf" }

func (address *InterfacesState_Interface_Ipv4_Address) GetYangName() string { return "address" }

func (address *InterfacesState_Interface_Ipv4_Address) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (address *InterfacesState_Interface_Ipv4_Address) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (address *InterfacesState_Interface_Ipv4_Address) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (address *InterfacesState_Interface_Ipv4_Address) SetParent(parent types.Entity) { address.parent = parent }

func (address *InterfacesState_Interface_Ipv4_Address) GetParent() types.Entity { return address.parent }

func (address *InterfacesState_Interface_Ipv4_Address) GetParentYangName() string { return "ipv4" }

// InterfacesState_Interface_Ipv4_Neighbor
// A list of mappings from IPv4 addresses to
// link-layer addresses.
// This list represents the ARP Cache.
type InterfacesState_Interface_Ipv4_Neighbor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The IPv4 address of the neighbor node. The type is
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ip interface{}

    // The link-layer address of the neighbor node. The type is string with
    // pattern: ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    LinkLayerAddress interface{}

    // The origin of this neighbor entry. The type is NeighborOrigin.
    Origin interface{}
}

func (neighbor *InterfacesState_Interface_Ipv4_Neighbor) GetFilter() yfilter.YFilter { return neighbor.YFilter }

func (neighbor *InterfacesState_Interface_Ipv4_Neighbor) SetFilter(yf yfilter.YFilter) { neighbor.YFilter = yf }

func (neighbor *InterfacesState_Interface_Ipv4_Neighbor) GetGoName(yname string) string {
    if yname == "ip" { return "Ip" }
    if yname == "link-layer-address" { return "LinkLayerAddress" }
    if yname == "origin" { return "Origin" }
    return ""
}

func (neighbor *InterfacesState_Interface_Ipv4_Neighbor) GetSegmentPath() string {
    return "neighbor" + "[ip='" + fmt.Sprintf("%v", neighbor.Ip) + "']"
}

func (neighbor *InterfacesState_Interface_Ipv4_Neighbor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (neighbor *InterfacesState_Interface_Ipv4_Neighbor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (neighbor *InterfacesState_Interface_Ipv4_Neighbor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip"] = neighbor.Ip
    leafs["link-layer-address"] = neighbor.LinkLayerAddress
    leafs["origin"] = neighbor.Origin
    return leafs
}

func (neighbor *InterfacesState_Interface_Ipv4_Neighbor) GetBundleName() string { return "ietf" }

func (neighbor *InterfacesState_Interface_Ipv4_Neighbor) GetYangName() string { return "neighbor" }

func (neighbor *InterfacesState_Interface_Ipv4_Neighbor) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (neighbor *InterfacesState_Interface_Ipv4_Neighbor) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (neighbor *InterfacesState_Interface_Ipv4_Neighbor) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (neighbor *InterfacesState_Interface_Ipv4_Neighbor) SetParent(parent types.Entity) { neighbor.parent = parent }

func (neighbor *InterfacesState_Interface_Ipv4_Neighbor) GetParent() types.Entity { return neighbor.parent }

func (neighbor *InterfacesState_Interface_Ipv4_Neighbor) GetParentYangName() string { return "ipv4" }

// InterfacesState_Interface_Ipv6
// Parameters for the IPv6 address family.
// This type is a presence type.
type InterfacesState_Interface_Ipv6 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Indicates whether IPv6 packet forwarding is enabled or disabled on this
    // interface. The type is bool. The default value is false.
    Forwarding interface{}

    // The size, in octets, of the largest IPv6 packet that the interface will
    // send and receive. The type is interface{} with range: 1280..4294967295.
    // Units are octets.
    Mtu interface{}

    // The list of IPv6 addresses on the interface. The type is slice of
    // InterfacesState_Interface_Ipv6_Address.
    Address []InterfacesState_Interface_Ipv6_Address

    // A list of mappings from IPv6 addresses to link-layer addresses. This list
    // represents the Neighbor Cache. The type is slice of
    // InterfacesState_Interface_Ipv6_Neighbor.
    Neighbor []InterfacesState_Interface_Ipv6_Neighbor
}

func (ipv6 *InterfacesState_Interface_Ipv6) GetFilter() yfilter.YFilter { return ipv6.YFilter }

func (ipv6 *InterfacesState_Interface_Ipv6) SetFilter(yf yfilter.YFilter) { ipv6.YFilter = yf }

func (ipv6 *InterfacesState_Interface_Ipv6) GetGoName(yname string) string {
    if yname == "forwarding" { return "Forwarding" }
    if yname == "mtu" { return "Mtu" }
    if yname == "address" { return "Address" }
    if yname == "neighbor" { return "Neighbor" }
    return ""
}

func (ipv6 *InterfacesState_Interface_Ipv6) GetSegmentPath() string {
    return "ietf-ip:ipv6"
}

func (ipv6 *InterfacesState_Interface_Ipv6) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "address" {
        for _, c := range ipv6.Address {
            if ipv6.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := InterfacesState_Interface_Ipv6_Address{}
        ipv6.Address = append(ipv6.Address, child)
        return &ipv6.Address[len(ipv6.Address)-1]
    }
    if childYangName == "neighbor" {
        for _, c := range ipv6.Neighbor {
            if ipv6.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := InterfacesState_Interface_Ipv6_Neighbor{}
        ipv6.Neighbor = append(ipv6.Neighbor, child)
        return &ipv6.Neighbor[len(ipv6.Neighbor)-1]
    }
    return nil
}

func (ipv6 *InterfacesState_Interface_Ipv6) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ipv6.Address {
        children[ipv6.Address[i].GetSegmentPath()] = &ipv6.Address[i]
    }
    for i := range ipv6.Neighbor {
        children[ipv6.Neighbor[i].GetSegmentPath()] = &ipv6.Neighbor[i]
    }
    return children
}

func (ipv6 *InterfacesState_Interface_Ipv6) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["forwarding"] = ipv6.Forwarding
    leafs["mtu"] = ipv6.Mtu
    return leafs
}

func (ipv6 *InterfacesState_Interface_Ipv6) GetBundleName() string { return "ietf" }

func (ipv6 *InterfacesState_Interface_Ipv6) GetYangName() string { return "ipv6" }

func (ipv6 *InterfacesState_Interface_Ipv6) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (ipv6 *InterfacesState_Interface_Ipv6) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (ipv6 *InterfacesState_Interface_Ipv6) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (ipv6 *InterfacesState_Interface_Ipv6) SetParent(parent types.Entity) { ipv6.parent = parent }

func (ipv6 *InterfacesState_Interface_Ipv6) GetParent() types.Entity { return ipv6.parent }

func (ipv6 *InterfacesState_Interface_Ipv6) GetParentYangName() string { return "interface" }

// InterfacesState_Interface_Ipv6_Address
// The list of IPv6 addresses on the interface.
type InterfacesState_Interface_Ipv6_Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The IPv6 address on the interface. The type is
    // string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ip interface{}

    // The length of the subnet prefix. The type is interface{} with range:
    // 0..128. This attribute is mandatory.
    PrefixLength interface{}

    // The origin of this address. The type is IpAddressOrigin.
    Origin interface{}

    // The status of an address.  Most of the states correspond to states from the
    // IPv6 Stateless Address Autoconfiguration protocol. The type is Status.
    Status interface{}
}

func (address *InterfacesState_Interface_Ipv6_Address) GetFilter() yfilter.YFilter { return address.YFilter }

func (address *InterfacesState_Interface_Ipv6_Address) SetFilter(yf yfilter.YFilter) { address.YFilter = yf }

func (address *InterfacesState_Interface_Ipv6_Address) GetGoName(yname string) string {
    if yname == "ip" { return "Ip" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "origin" { return "Origin" }
    if yname == "status" { return "Status" }
    return ""
}

func (address *InterfacesState_Interface_Ipv6_Address) GetSegmentPath() string {
    return "address" + "[ip='" + fmt.Sprintf("%v", address.Ip) + "']"
}

func (address *InterfacesState_Interface_Ipv6_Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (address *InterfacesState_Interface_Ipv6_Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (address *InterfacesState_Interface_Ipv6_Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip"] = address.Ip
    leafs["prefix-length"] = address.PrefixLength
    leafs["origin"] = address.Origin
    leafs["status"] = address.Status
    return leafs
}

func (address *InterfacesState_Interface_Ipv6_Address) GetBundleName() string { return "ietf" }

func (address *InterfacesState_Interface_Ipv6_Address) GetYangName() string { return "address" }

func (address *InterfacesState_Interface_Ipv6_Address) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (address *InterfacesState_Interface_Ipv6_Address) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (address *InterfacesState_Interface_Ipv6_Address) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (address *InterfacesState_Interface_Ipv6_Address) SetParent(parent types.Entity) { address.parent = parent }

func (address *InterfacesState_Interface_Ipv6_Address) GetParent() types.Entity { return address.parent }

func (address *InterfacesState_Interface_Ipv6_Address) GetParentYangName() string { return "ipv6" }

// InterfacesState_Interface_Ipv6_Address_Status represents Autoconfiguration protocol.
type InterfacesState_Interface_Ipv6_Address_Status string

const (
    // This is a valid address that can appear as the
    // destination or source address of a packet.
    InterfacesState_Interface_Ipv6_Address_Status_preferred InterfacesState_Interface_Ipv6_Address_Status = "preferred"

    // This is a valid but deprecated address that should
    // no longer be used as a source address in new
    // communications, but packets addressed to such an
    // address are processed as expected.
    InterfacesState_Interface_Ipv6_Address_Status_deprecated InterfacesState_Interface_Ipv6_Address_Status = "deprecated"

    // This isn't a valid address, and it shouldn't appear
    // as the destination or source address of a packet.
    InterfacesState_Interface_Ipv6_Address_Status_invalid InterfacesState_Interface_Ipv6_Address_Status = "invalid"

    // The address is not accessible because the interface
    // to which this address is assigned is not
    // operational.
    InterfacesState_Interface_Ipv6_Address_Status_inaccessible InterfacesState_Interface_Ipv6_Address_Status = "inaccessible"

    // The status cannot be determined for some reason.
    InterfacesState_Interface_Ipv6_Address_Status_unknown InterfacesState_Interface_Ipv6_Address_Status = "unknown"

    // The uniqueness of the address on the link is being
    // verified.  Addresses in this state should not be
    // used for general communication and should only be
    // used to determine the uniqueness of the address.
    InterfacesState_Interface_Ipv6_Address_Status_tentative InterfacesState_Interface_Ipv6_Address_Status = "tentative"

    // The address has been determined to be non-unique on
    // the link and so must not be used.
    InterfacesState_Interface_Ipv6_Address_Status_duplicate InterfacesState_Interface_Ipv6_Address_Status = "duplicate"

    // The address is available for use, subject to
    // restrictions, while its uniqueness on a link is
    // being verified.
    InterfacesState_Interface_Ipv6_Address_Status_optimistic InterfacesState_Interface_Ipv6_Address_Status = "optimistic"
)

// InterfacesState_Interface_Ipv6_Neighbor
// A list of mappings from IPv6 addresses to
// link-layer addresses.
// This list represents the Neighbor Cache.
type InterfacesState_Interface_Ipv6_Neighbor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The IPv6 address of the neighbor node. The type is
    // string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ip interface{}

    // The link-layer address of the neighbor node. The type is string with
    // pattern: ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    LinkLayerAddress interface{}

    // The origin of this neighbor entry. The type is NeighborOrigin.
    Origin interface{}

    // Indicates that the neighbor node acts as a router. The type is interface{}.
    IsRouter interface{}

    // The Neighbor Unreachability Detection state of this entry. The type is
    // State.
    State interface{}
}

func (neighbor *InterfacesState_Interface_Ipv6_Neighbor) GetFilter() yfilter.YFilter { return neighbor.YFilter }

func (neighbor *InterfacesState_Interface_Ipv6_Neighbor) SetFilter(yf yfilter.YFilter) { neighbor.YFilter = yf }

func (neighbor *InterfacesState_Interface_Ipv6_Neighbor) GetGoName(yname string) string {
    if yname == "ip" { return "Ip" }
    if yname == "link-layer-address" { return "LinkLayerAddress" }
    if yname == "origin" { return "Origin" }
    if yname == "is-router" { return "IsRouter" }
    if yname == "state" { return "State" }
    return ""
}

func (neighbor *InterfacesState_Interface_Ipv6_Neighbor) GetSegmentPath() string {
    return "neighbor" + "[ip='" + fmt.Sprintf("%v", neighbor.Ip) + "']"
}

func (neighbor *InterfacesState_Interface_Ipv6_Neighbor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (neighbor *InterfacesState_Interface_Ipv6_Neighbor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (neighbor *InterfacesState_Interface_Ipv6_Neighbor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip"] = neighbor.Ip
    leafs["link-layer-address"] = neighbor.LinkLayerAddress
    leafs["origin"] = neighbor.Origin
    leafs["is-router"] = neighbor.IsRouter
    leafs["state"] = neighbor.State
    return leafs
}

func (neighbor *InterfacesState_Interface_Ipv6_Neighbor) GetBundleName() string { return "ietf" }

func (neighbor *InterfacesState_Interface_Ipv6_Neighbor) GetYangName() string { return "neighbor" }

func (neighbor *InterfacesState_Interface_Ipv6_Neighbor) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (neighbor *InterfacesState_Interface_Ipv6_Neighbor) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (neighbor *InterfacesState_Interface_Ipv6_Neighbor) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (neighbor *InterfacesState_Interface_Ipv6_Neighbor) SetParent(parent types.Entity) { neighbor.parent = parent }

func (neighbor *InterfacesState_Interface_Ipv6_Neighbor) GetParent() types.Entity { return neighbor.parent }

func (neighbor *InterfacesState_Interface_Ipv6_Neighbor) GetParentYangName() string { return "ipv6" }

// InterfacesState_Interface_Ipv6_Neighbor_State represents entry.
type InterfacesState_Interface_Ipv6_Neighbor_State string

const (
    // Address resolution is in progress, and the link-layer
    // address of the neighbor has not yet been
    // determined.
    InterfacesState_Interface_Ipv6_Neighbor_State_incomplete InterfacesState_Interface_Ipv6_Neighbor_State = "incomplete"

    // Roughly speaking, the neighbor is known to have been
    // reachable recently (within tens of seconds ago).
    InterfacesState_Interface_Ipv6_Neighbor_State_reachable InterfacesState_Interface_Ipv6_Neighbor_State = "reachable"

    // The neighbor is no longer known to be reachable, but
    // until traffic is sent to the neighbor no attempt
    // should be made to verify its reachability.
    InterfacesState_Interface_Ipv6_Neighbor_State_stale InterfacesState_Interface_Ipv6_Neighbor_State = "stale"

    // The neighbor is no longer known to be reachable, and
    // traffic has recently been sent to the neighbor.
    // Rather than probe the neighbor immediately, however,
    // delay sending probes for a short while in order to
    // give upper-layer protocols a chance to provide
    // reachability confirmation.
    InterfacesState_Interface_Ipv6_Neighbor_State_delay InterfacesState_Interface_Ipv6_Neighbor_State = "delay"

    // The neighbor is no longer known to be reachable, and
    // unicast Neighbor Solicitation probes are being sent
    // to verify reachability.
    InterfacesState_Interface_Ipv6_Neighbor_State_probe InterfacesState_Interface_Ipv6_Neighbor_State = "probe"
)

// InterfacesState_Interface_AdminStatus represents This leaf has the same read semantics as ifAdminStatus.
type InterfacesState_Interface_AdminStatus string

const (
    // Ready to pass packets.
    InterfacesState_Interface_AdminStatus_up InterfacesState_Interface_AdminStatus = "up"

    // Not ready to pass packets and not in some test mode.
    InterfacesState_Interface_AdminStatus_down InterfacesState_Interface_AdminStatus = "down"

    // In some test mode.
    InterfacesState_Interface_AdminStatus_testing InterfacesState_Interface_AdminStatus = "testing"
)

// InterfacesState_Interface_OperStatus represents This leaf has the same semantics as ifOperStatus.
type InterfacesState_Interface_OperStatus string

const (
    // Ready to pass packets.
    InterfacesState_Interface_OperStatus_up InterfacesState_Interface_OperStatus = "up"

    // The interface does not pass any packets.
    InterfacesState_Interface_OperStatus_down InterfacesState_Interface_OperStatus = "down"

    // In some test mode.  No operational packets can
    // be passed.
    InterfacesState_Interface_OperStatus_testing InterfacesState_Interface_OperStatus = "testing"

    // Status cannot be determined for some reason.
    InterfacesState_Interface_OperStatus_unknown InterfacesState_Interface_OperStatus = "unknown"

    // Waiting for some external event.
    InterfacesState_Interface_OperStatus_dormant InterfacesState_Interface_OperStatus = "dormant"

    // Some component (typically hardware) is missing.
    InterfacesState_Interface_OperStatus_not_present InterfacesState_Interface_OperStatus = "not-present"

    // Down due to state of lower-layer interface(s).
    InterfacesState_Interface_OperStatus_lower_layer_down InterfacesState_Interface_OperStatus = "lower-layer-down"
)

