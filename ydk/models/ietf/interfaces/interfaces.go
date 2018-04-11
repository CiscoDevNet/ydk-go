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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The list of configured interfaces on the device.  The operational state of
    // an interface is available in the /interfaces-state/interface list.  If the
    // configuration of a system-controlled interface cannot be used by the system
    // (e.g., the interface hardware present does not match the interface type),
    // then the configuration is not applied to the system-controlled interface
    // shown in the /interfaces-state/interface list.  If the configuration of a
    // user-controlled interface cannot be used by the system, the configured
    // interface is not instantiated in the /interfaces-state/interface list. The
    // type is slice of Interfaces_Interface_.
    Interface_ []Interfaces_Interface
}

func (interfaces *Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "ietf"
    interfaces.EntityData.ParentYangName = "ietf-interfaces"
    interfaces.EntityData.SegmentPath = "ietf-interfaces:interfaces"
    interfaces.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    interfaces.EntityData.NamespaceTable = ietf.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    interfaces.EntityData.Children = make(map[string]types.YChild)
    interfaces.EntityData.Children["interface"] = types.YChild{"Interface_", nil}
    for i := range interfaces.Interface_ {
        interfaces.EntityData.Children[types.GetSegmentPath(&interfaces.Interface_[i])] = types.YChild{"Interface_", &interfaces.Interface_[i]}
    }
    interfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaces.EntityData)
}

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
    EntityData types.CommonEntityData
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
    // IanaInterfaceTypeOtherRegular1822Hdh1822Ddnx25Rfc877X25EthernetcsmacdIso88023CsmacdIso88024TokenbusIso88025TokenringIso88026ManStarlanProteon10MbitProteon80MbitHyperchannelFddiLapbSdlcDs1E1BasicisdnPrimaryisdnProppointtopointserialPppSoftwareloopbackEonEthernet3MbitNsipSlipUltraDs3SipFramerelayRs232ParaArcnetArcnetplusAtmMiox25SonetX25PleIso88022LlcLocaltalkSmdsdxiFramerelayserviceV35HssiHippiModemAal5SonetpathSonetvtSmdsicipPropvirtualPropmultiplexorIeee80212FibrechannelHippiinterfaceFramerelayinterconnectAflane8023Aflane8025CctemulFastetherIsdnV11V36G703At64KG703At2MbQllcFastetherfxChannelIeee80211Ibm370ParchanEsconDlswIsdnsIsdnuLapdIpswitchRsrbAtmlogicalDs0Ds0BundleBscAsyncCnrIso88025DtrEplrsArapPropcnlsHostpadTermpadFramerelaympiX213AdslRadslSdslVdslIso88025CrfpintMyrinetVoiceemVoicefxoVoicefxsVoiceencapVoiceoveripAtmdxiAtmfuniAtmimaPppmultilinkbundleIpovercdlcIpoverclawStacktostackVirtualipaddressMpcIpoveratmIso88025FiberTdlcGigabitethernetHdlcLapfV37X25MlpX25HuntgroupTransphdlcInterleaveFastIpDocscablemaclayerDocscabledownstreamDocscableupstreamA12MppswitchTunnelCoffeeCesAtmsubinterfaceL2VlanL3IpvlanL3IpxvlanDigitalpowerlineMediamailoveripDtmDcnIpforwardMsdslIeee1394IfGsnDvbrccmaclayerDvbrccdownstreamDvbrccupstreamAtmvirtualMplstunnelSrpVoiceoveratmVoiceoverframerelayIdslCompositelinkSs7SiglinkPropwirelessp2PFrforwardRfc1483UsbIeee8023AdlagBgppolicyaccountingFrf16MfrbundleH323GatekeeperH323ProxyMplsMfsiglinkHdsl2ShdslDs1FdlPosDvbasiinDvbasioutPlcNfasTr008Gr303RdtGr303IdtIsupPropdocswirelessmaclayerPropdocswirelessdownstreamPropdocswirelessupstreamHiperlan2Propbwap2MpSonetoverheadchannelDigitalwrapperoverheadchannelAal2RadiomacAtmradioImtMvlReachdslFrdlciendptAtmvciendptOpticalchannelOpticaltransportPropatmVoiceovercableInfinibandTelinkQ2931VirtualtgSiptgSipsigDocscableupstreamchannelEconetPon155Pon622BridgeLinegroupVoiceemfgdVoicefgdeanaVoicedidMpegtransportSixtofourGtpPdnetherloop1Pdnetherloop2OpticalchannelgroupHomepnaGfpCiscoislvlanActelismetaloopFciplinkRprQamLmpCblvectastarDocscablemcmtsdownstreamAdsl2MacseccontrolledifMacsecuncontrolledifAviciopticaletherAtmbondVoicefgdosMocaversion1Ieee80216WmanAdsl2PlusDvbrcsmaclayerDvbtdmDvbrcstdmaX86LapsWwanppWwanpp2VoiceebsIfpwtypeIlanPipAluelpGponVdsl2Capwapdot11ProfileCapwapdot11BssCapwapwtpvirtualradioBitsDocscableupstreamrfportCabledownstreamrfportVmwarevirtualnicIeee802154OtnoduOtnotuIfvfitypeG9981G9982G9983AlueponAluepononuAlueponphysicaluniAlueponlogicallinkAlugpononuAlugponphysicaluniVmwarenicteam.
    // This attribute is mandatory.
    Type_ interface{}

    // This leaf contains the configured, desired state of the interface.  Systems
    // that implement the IF-MIB use the value of this leaf in the 'running'
    // datastore to set IF-MIB.ifAdminStatus to 'up' or 'down' after an ifEntry
    // has been initialized, as described in RFC 2863.  Changes in this leaf in
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

func (self *Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "ietf"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + "[name='" + fmt.Sprintf("%v", self.Name) + "']"
    self.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    self.EntityData.NamespaceTable = ietf.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Children["ietf-diffserv-target:diffserv-target-entry"] = types.YChild{"DiffservTargetEntry", nil}
    for i := range self.DiffservTargetEntry {
        self.EntityData.Children[types.GetSegmentPath(&self.DiffservTargetEntry[i])] = types.YChild{"DiffservTargetEntry", &self.DiffservTargetEntry[i]}
    }
    self.EntityData.Children["ietf-ip:ipv4"] = types.YChild{"Ipv4", &self.Ipv4}
    self.EntityData.Children["ietf-ip:ipv6"] = types.YChild{"Ipv6", &self.Ipv6}
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["name"] = types.YLeaf{"Name", self.Name}
    self.EntityData.Leafs["description"] = types.YLeaf{"Description", self.Description}
    self.EntityData.Leafs["type"] = types.YLeaf{"Type_", self.Type_}
    self.EntityData.Leafs["enabled"] = types.YLeaf{"Enabled", self.Enabled}
    self.EntityData.Leafs["link-up-down-trap-enable"] = types.YLeaf{"LinkUpDownTrapEnable", self.LinkUpDownTrapEnable}
    return &(self.EntityData)
}

// Interfaces_Interface_DiffservTargetEntry
// policy target for inbound or outbound direction
type Interfaces_Interface_DiffservTargetEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Direction fo the traffic flow either inbound or
    // outbound. The type is one of the following: InboundOutbound.
    Direction interface{}

    // This attribute is a key. Policy entry name. The type is string.
    PolicyName interface{}
}

func (diffservTargetEntry *Interfaces_Interface_DiffservTargetEntry) GetEntityData() *types.CommonEntityData {
    diffservTargetEntry.EntityData.YFilter = diffservTargetEntry.YFilter
    diffservTargetEntry.EntityData.YangName = "diffserv-target-entry"
    diffservTargetEntry.EntityData.BundleName = "ietf"
    diffservTargetEntry.EntityData.ParentYangName = "interface"
    diffservTargetEntry.EntityData.SegmentPath = "ietf-diffserv-target:diffserv-target-entry" + "[direction='" + fmt.Sprintf("%v", diffservTargetEntry.Direction) + "']" + "[policy-name='" + fmt.Sprintf("%v", diffservTargetEntry.PolicyName) + "']"
    diffservTargetEntry.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    diffservTargetEntry.EntityData.NamespaceTable = ietf.GetNamespaces()
    diffservTargetEntry.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    diffservTargetEntry.EntityData.Children = make(map[string]types.YChild)
    diffservTargetEntry.EntityData.Leafs = make(map[string]types.YLeaf)
    diffservTargetEntry.EntityData.Leafs["direction"] = types.YLeaf{"Direction", diffservTargetEntry.Direction}
    diffservTargetEntry.EntityData.Leafs["policy-name"] = types.YLeaf{"PolicyName", diffservTargetEntry.PolicyName}
    return &(diffservTargetEntry.EntityData)
}

// Interfaces_Interface_Ipv4
// Parameters for the IPv4 address family.
// This type is a presence type.
type Interfaces_Interface_Ipv4 struct {
    EntityData types.CommonEntityData
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

func (ipv4 *Interfaces_Interface_Ipv4) GetEntityData() *types.CommonEntityData {
    ipv4.EntityData.YFilter = ipv4.YFilter
    ipv4.EntityData.YangName = "ipv4"
    ipv4.EntityData.BundleName = "ietf"
    ipv4.EntityData.ParentYangName = "interface"
    ipv4.EntityData.SegmentPath = "ietf-ip:ipv4"
    ipv4.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    ipv4.EntityData.NamespaceTable = ietf.GetNamespaces()
    ipv4.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    ipv4.EntityData.Children = make(map[string]types.YChild)
    ipv4.EntityData.Children["address"] = types.YChild{"Address", nil}
    for i := range ipv4.Address {
        ipv4.EntityData.Children[types.GetSegmentPath(&ipv4.Address[i])] = types.YChild{"Address", &ipv4.Address[i]}
    }
    ipv4.EntityData.Children["neighbor"] = types.YChild{"Neighbor", nil}
    for i := range ipv4.Neighbor {
        ipv4.EntityData.Children[types.GetSegmentPath(&ipv4.Neighbor[i])] = types.YChild{"Neighbor", &ipv4.Neighbor[i]}
    }
    ipv4.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv4.EntityData.Leafs["enabled"] = types.YLeaf{"Enabled", ipv4.Enabled}
    ipv4.EntityData.Leafs["forwarding"] = types.YLeaf{"Forwarding", ipv4.Forwarding}
    ipv4.EntityData.Leafs["mtu"] = types.YLeaf{"Mtu", ipv4.Mtu}
    return &(ipv4.EntityData)
}

// Interfaces_Interface_Ipv4_Address
// The list of configured IPv4 addresses on the interface.
type Interfaces_Interface_Ipv4_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The IPv4 address on the interface. The type is
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ip interface{}

    // The length of the subnet prefix. The type is interface{} with range: 0..32.
    PrefixLength interface{}

    // The subnet specified as a netmask. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])'.
    Netmask interface{}
}

func (address *Interfaces_Interface_Ipv4_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "ietf"
    address.EntityData.ParentYangName = "ipv4"
    address.EntityData.SegmentPath = "address" + "[ip='" + fmt.Sprintf("%v", address.Ip) + "']"
    address.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    address.EntityData.NamespaceTable = ietf.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    address.EntityData.Children = make(map[string]types.YChild)
    address.EntityData.Leafs = make(map[string]types.YLeaf)
    address.EntityData.Leafs["ip"] = types.YLeaf{"Ip", address.Ip}
    address.EntityData.Leafs["prefix-length"] = types.YLeaf{"PrefixLength", address.PrefixLength}
    address.EntityData.Leafs["netmask"] = types.YLeaf{"Netmask", address.Netmask}
    return &(address.EntityData)
}

// Interfaces_Interface_Ipv4_Neighbor
// A list of mappings from IPv4 addresses to
// link-layer addresses.
// Entries in this list are used as static entries in the
// ARP Cache.
type Interfaces_Interface_Ipv4_Neighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The IPv4 address of the neighbor node. The type is
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ip interface{}

    // The link-layer address of the neighbor node. The type is string with
    // pattern: b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'. This attribute is
    // mandatory.
    LinkLayerAddress interface{}
}

func (neighbor *Interfaces_Interface_Ipv4_Neighbor) GetEntityData() *types.CommonEntityData {
    neighbor.EntityData.YFilter = neighbor.YFilter
    neighbor.EntityData.YangName = "neighbor"
    neighbor.EntityData.BundleName = "ietf"
    neighbor.EntityData.ParentYangName = "ipv4"
    neighbor.EntityData.SegmentPath = "neighbor" + "[ip='" + fmt.Sprintf("%v", neighbor.Ip) + "']"
    neighbor.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    neighbor.EntityData.NamespaceTable = ietf.GetNamespaces()
    neighbor.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    neighbor.EntityData.Children = make(map[string]types.YChild)
    neighbor.EntityData.Leafs = make(map[string]types.YLeaf)
    neighbor.EntityData.Leafs["ip"] = types.YLeaf{"Ip", neighbor.Ip}
    neighbor.EntityData.Leafs["link-layer-address"] = types.YLeaf{"LinkLayerAddress", neighbor.LinkLayerAddress}
    return &(neighbor.EntityData)
}

// Interfaces_Interface_Ipv6
// Parameters for the IPv6 address family.
// This type is a presence type.
type Interfaces_Interface_Ipv6 struct {
    EntityData types.CommonEntityData
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

func (ipv6 *Interfaces_Interface_Ipv6) GetEntityData() *types.CommonEntityData {
    ipv6.EntityData.YFilter = ipv6.YFilter
    ipv6.EntityData.YangName = "ipv6"
    ipv6.EntityData.BundleName = "ietf"
    ipv6.EntityData.ParentYangName = "interface"
    ipv6.EntityData.SegmentPath = "ietf-ip:ipv6"
    ipv6.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    ipv6.EntityData.NamespaceTable = ietf.GetNamespaces()
    ipv6.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    ipv6.EntityData.Children = make(map[string]types.YChild)
    ipv6.EntityData.Children["address"] = types.YChild{"Address", nil}
    for i := range ipv6.Address {
        ipv6.EntityData.Children[types.GetSegmentPath(&ipv6.Address[i])] = types.YChild{"Address", &ipv6.Address[i]}
    }
    ipv6.EntityData.Children["neighbor"] = types.YChild{"Neighbor", nil}
    for i := range ipv6.Neighbor {
        ipv6.EntityData.Children[types.GetSegmentPath(&ipv6.Neighbor[i])] = types.YChild{"Neighbor", &ipv6.Neighbor[i]}
    }
    ipv6.EntityData.Children["autoconf"] = types.YChild{"Autoconf", &ipv6.Autoconf}
    ipv6.EntityData.Children["ietf-ipv6-unicast-routing:ipv6-router-advertisements"] = types.YChild{"Ipv6RouterAdvertisements", &ipv6.Ipv6RouterAdvertisements}
    ipv6.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv6.EntityData.Leafs["enabled"] = types.YLeaf{"Enabled", ipv6.Enabled}
    ipv6.EntityData.Leafs["forwarding"] = types.YLeaf{"Forwarding", ipv6.Forwarding}
    ipv6.EntityData.Leafs["mtu"] = types.YLeaf{"Mtu", ipv6.Mtu}
    ipv6.EntityData.Leafs["dup-addr-detect-transmits"] = types.YLeaf{"DupAddrDetectTransmits", ipv6.DupAddrDetectTransmits}
    return &(ipv6.EntityData)
}

// Interfaces_Interface_Ipv6_Address
// The list of configured IPv6 addresses on the interface.
type Interfaces_Interface_Ipv6_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The IPv6 address on the interface. The type is
    // string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ip interface{}

    // The length of the subnet prefix. The type is interface{} with range:
    // 0..128. This attribute is mandatory.
    PrefixLength interface{}
}

func (address *Interfaces_Interface_Ipv6_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "ietf"
    address.EntityData.ParentYangName = "ipv6"
    address.EntityData.SegmentPath = "address" + "[ip='" + fmt.Sprintf("%v", address.Ip) + "']"
    address.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    address.EntityData.NamespaceTable = ietf.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    address.EntityData.Children = make(map[string]types.YChild)
    address.EntityData.Leafs = make(map[string]types.YLeaf)
    address.EntityData.Leafs["ip"] = types.YLeaf{"Ip", address.Ip}
    address.EntityData.Leafs["prefix-length"] = types.YLeaf{"PrefixLength", address.PrefixLength}
    return &(address.EntityData)
}

// Interfaces_Interface_Ipv6_Neighbor
// A list of mappings from IPv6 addresses to
// link-layer addresses.
// Entries in this list are used as static entries in the
// Neighbor Cache.
type Interfaces_Interface_Ipv6_Neighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The IPv6 address of the neighbor node. The type is
    // string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ip interface{}

    // The link-layer address of the neighbor node. The type is string with
    // pattern: b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'. This attribute is
    // mandatory.
    LinkLayerAddress interface{}
}

func (neighbor *Interfaces_Interface_Ipv6_Neighbor) GetEntityData() *types.CommonEntityData {
    neighbor.EntityData.YFilter = neighbor.YFilter
    neighbor.EntityData.YangName = "neighbor"
    neighbor.EntityData.BundleName = "ietf"
    neighbor.EntityData.ParentYangName = "ipv6"
    neighbor.EntityData.SegmentPath = "neighbor" + "[ip='" + fmt.Sprintf("%v", neighbor.Ip) + "']"
    neighbor.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    neighbor.EntityData.NamespaceTable = ietf.GetNamespaces()
    neighbor.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    neighbor.EntityData.Children = make(map[string]types.YChild)
    neighbor.EntityData.Leafs = make(map[string]types.YLeaf)
    neighbor.EntityData.Leafs["ip"] = types.YLeaf{"Ip", neighbor.Ip}
    neighbor.EntityData.Leafs["link-layer-address"] = types.YLeaf{"LinkLayerAddress", neighbor.LinkLayerAddress}
    return &(neighbor.EntityData)
}

// Interfaces_Interface_Ipv6_Autoconf
// Parameters to control the autoconfiguration of IPv6
// addresses, as described in RFC 4862.
type Interfaces_Interface_Ipv6_Autoconf struct {
    EntityData types.CommonEntityData
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

func (autoconf *Interfaces_Interface_Ipv6_Autoconf) GetEntityData() *types.CommonEntityData {
    autoconf.EntityData.YFilter = autoconf.YFilter
    autoconf.EntityData.YangName = "autoconf"
    autoconf.EntityData.BundleName = "ietf"
    autoconf.EntityData.ParentYangName = "ipv6"
    autoconf.EntityData.SegmentPath = "autoconf"
    autoconf.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    autoconf.EntityData.NamespaceTable = ietf.GetNamespaces()
    autoconf.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    autoconf.EntityData.Children = make(map[string]types.YChild)
    autoconf.EntityData.Leafs = make(map[string]types.YLeaf)
    autoconf.EntityData.Leafs["create-global-addresses"] = types.YLeaf{"CreateGlobalAddresses", autoconf.CreateGlobalAddresses}
    autoconf.EntityData.Leafs["create-temporary-addresses"] = types.YLeaf{"CreateTemporaryAddresses", autoconf.CreateTemporaryAddresses}
    autoconf.EntityData.Leafs["temporary-valid-lifetime"] = types.YLeaf{"TemporaryValidLifetime", autoconf.TemporaryValidLifetime}
    autoconf.EntityData.Leafs["temporary-preferred-lifetime"] = types.YLeaf{"TemporaryPreferredLifetime", autoconf.TemporaryPreferredLifetime}
    return &(autoconf.EntityData)
}

// Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements
// Configuration of IPv6 Router Advertisements.
type Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements struct {
    EntityData types.CommonEntityData
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

func (ipv6RouterAdvertisements *Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements) GetEntityData() *types.CommonEntityData {
    ipv6RouterAdvertisements.EntityData.YFilter = ipv6RouterAdvertisements.YFilter
    ipv6RouterAdvertisements.EntityData.YangName = "ipv6-router-advertisements"
    ipv6RouterAdvertisements.EntityData.BundleName = "ietf"
    ipv6RouterAdvertisements.EntityData.ParentYangName = "ipv6"
    ipv6RouterAdvertisements.EntityData.SegmentPath = "ietf-ipv6-unicast-routing:ipv6-router-advertisements"
    ipv6RouterAdvertisements.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    ipv6RouterAdvertisements.EntityData.NamespaceTable = ietf.GetNamespaces()
    ipv6RouterAdvertisements.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    ipv6RouterAdvertisements.EntityData.Children = make(map[string]types.YChild)
    ipv6RouterAdvertisements.EntityData.Children["prefix-list"] = types.YChild{"PrefixList", &ipv6RouterAdvertisements.PrefixList}
    ipv6RouterAdvertisements.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv6RouterAdvertisements.EntityData.Leafs["send-advertisements"] = types.YLeaf{"SendAdvertisements", ipv6RouterAdvertisements.SendAdvertisements}
    ipv6RouterAdvertisements.EntityData.Leafs["max-rtr-adv-interval"] = types.YLeaf{"MaxRtrAdvInterval", ipv6RouterAdvertisements.MaxRtrAdvInterval}
    ipv6RouterAdvertisements.EntityData.Leafs["min-rtr-adv-interval"] = types.YLeaf{"MinRtrAdvInterval", ipv6RouterAdvertisements.MinRtrAdvInterval}
    ipv6RouterAdvertisements.EntityData.Leafs["managed-flag"] = types.YLeaf{"ManagedFlag", ipv6RouterAdvertisements.ManagedFlag}
    ipv6RouterAdvertisements.EntityData.Leafs["other-config-flag"] = types.YLeaf{"OtherConfigFlag", ipv6RouterAdvertisements.OtherConfigFlag}
    ipv6RouterAdvertisements.EntityData.Leafs["link-mtu"] = types.YLeaf{"LinkMtu", ipv6RouterAdvertisements.LinkMtu}
    ipv6RouterAdvertisements.EntityData.Leafs["reachable-time"] = types.YLeaf{"ReachableTime", ipv6RouterAdvertisements.ReachableTime}
    ipv6RouterAdvertisements.EntityData.Leafs["retrans-timer"] = types.YLeaf{"RetransTimer", ipv6RouterAdvertisements.RetransTimer}
    ipv6RouterAdvertisements.EntityData.Leafs["cur-hop-limit"] = types.YLeaf{"CurHopLimit", ipv6RouterAdvertisements.CurHopLimit}
    ipv6RouterAdvertisements.EntityData.Leafs["default-lifetime"] = types.YLeaf{"DefaultLifetime", ipv6RouterAdvertisements.DefaultLifetime}
    return &(ipv6RouterAdvertisements.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration of an advertised prefix entry. The type is slice of
    // Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements_PrefixList_Prefix.
    Prefix []Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements_PrefixList_Prefix
}

func (prefixList *Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements_PrefixList) GetEntityData() *types.CommonEntityData {
    prefixList.EntityData.YFilter = prefixList.YFilter
    prefixList.EntityData.YangName = "prefix-list"
    prefixList.EntityData.BundleName = "ietf"
    prefixList.EntityData.ParentYangName = "ipv6-router-advertisements"
    prefixList.EntityData.SegmentPath = "prefix-list"
    prefixList.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    prefixList.EntityData.NamespaceTable = ietf.GetNamespaces()
    prefixList.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    prefixList.EntityData.Children = make(map[string]types.YChild)
    prefixList.EntityData.Children["prefix"] = types.YChild{"Prefix", nil}
    for i := range prefixList.Prefix {
        prefixList.EntityData.Children[types.GetSegmentPath(&prefixList.Prefix[i])] = types.YChild{"Prefix", &prefixList.Prefix[i]}
    }
    prefixList.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(prefixList.EntityData)
}

// Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements_PrefixList_Prefix
// Configuration of an advertised prefix entry.
type Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements_PrefixList_Prefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. IPv6 address prefix. The type is string with
    // pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))'.
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

func (prefix *Interfaces_Interface_Ipv6_Ipv6RouterAdvertisements_PrefixList_Prefix) GetEntityData() *types.CommonEntityData {
    prefix.EntityData.YFilter = prefix.YFilter
    prefix.EntityData.YangName = "prefix"
    prefix.EntityData.BundleName = "ietf"
    prefix.EntityData.ParentYangName = "prefix-list"
    prefix.EntityData.SegmentPath = "prefix" + "[prefix-spec='" + fmt.Sprintf("%v", prefix.PrefixSpec) + "']"
    prefix.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    prefix.EntityData.NamespaceTable = ietf.GetNamespaces()
    prefix.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    prefix.EntityData.Children = make(map[string]types.YChild)
    prefix.EntityData.Leafs = make(map[string]types.YLeaf)
    prefix.EntityData.Leafs["prefix-spec"] = types.YLeaf{"PrefixSpec", prefix.PrefixSpec}
    prefix.EntityData.Leafs["no-advertise"] = types.YLeaf{"NoAdvertise", prefix.NoAdvertise}
    prefix.EntityData.Leafs["valid-lifetime"] = types.YLeaf{"ValidLifetime", prefix.ValidLifetime}
    prefix.EntityData.Leafs["on-link-flag"] = types.YLeaf{"OnLinkFlag", prefix.OnLinkFlag}
    prefix.EntityData.Leafs["preferred-lifetime"] = types.YLeaf{"PreferredLifetime", prefix.PreferredLifetime}
    prefix.EntityData.Leafs["autonomous-flag"] = types.YLeaf{"AutonomousFlag", prefix.AutonomousFlag}
    return &(prefix.EntityData)
}

// Interfaces_Interface_LinkUpDownTrapEnable represents no 'lower-layer-if' entries), and 'disabled' otherwise.
type Interfaces_Interface_LinkUpDownTrapEnable string

const (
    Interfaces_Interface_LinkUpDownTrapEnable_enabled Interfaces_Interface_LinkUpDownTrapEnable = "enabled"

    Interfaces_Interface_LinkUpDownTrapEnable_disabled Interfaces_Interface_LinkUpDownTrapEnable = "disabled"
)

// InterfacesState
// Data nodes for the operational state of interfaces.
type InterfacesState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The list of interfaces on the device.  System-controlled interfaces created
    // by the system are always present in this list, whether they are configured
    // or not. The type is slice of InterfacesState_Interface_.
    Interface_ []InterfacesState_Interface
}

func (interfacesState *InterfacesState) GetEntityData() *types.CommonEntityData {
    interfacesState.EntityData.YFilter = interfacesState.YFilter
    interfacesState.EntityData.YangName = "interfaces-state"
    interfacesState.EntityData.BundleName = "ietf"
    interfacesState.EntityData.ParentYangName = "ietf-interfaces"
    interfacesState.EntityData.SegmentPath = "ietf-interfaces:interfaces-state"
    interfacesState.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    interfacesState.EntityData.NamespaceTable = ietf.GetNamespaces()
    interfacesState.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    interfacesState.EntityData.Children = make(map[string]types.YChild)
    interfacesState.EntityData.Children["interface"] = types.YChild{"Interface_", nil}
    for i := range interfacesState.Interface_ {
        interfacesState.EntityData.Children[types.GetSegmentPath(&interfacesState.Interface_[i])] = types.YChild{"Interface_", &interfacesState.Interface_[i]}
    }
    interfacesState.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfacesState.EntityData)
}

// InterfacesState_Interface
// The list of interfaces on the device.
// 
// System-controlled interfaces created by the system are
// always present in this list, whether they are configured or
// not.
type InterfacesState_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the interface.  A server
    // implementation MAY map this leaf to the ifName MIB object.  Such an
    // implementation needs to use some mechanism to handle the differences in
    // size and characters allowed between this leaf and ifName.  The definition
    // of such a mechanism is outside the scope of this document. The type is
    // string.
    Name interface{}

    // The type of the interface. The type is one of the following:
    // IanaInterfaceTypeOtherRegular1822Hdh1822Ddnx25Rfc877X25EthernetcsmacdIso88023CsmacdIso88024TokenbusIso88025TokenringIso88026ManStarlanProteon10MbitProteon80MbitHyperchannelFddiLapbSdlcDs1E1BasicisdnPrimaryisdnProppointtopointserialPppSoftwareloopbackEonEthernet3MbitNsipSlipUltraDs3SipFramerelayRs232ParaArcnetArcnetplusAtmMiox25SonetX25PleIso88022LlcLocaltalkSmdsdxiFramerelayserviceV35HssiHippiModemAal5SonetpathSonetvtSmdsicipPropvirtualPropmultiplexorIeee80212FibrechannelHippiinterfaceFramerelayinterconnectAflane8023Aflane8025CctemulFastetherIsdnV11V36G703At64KG703At2MbQllcFastetherfxChannelIeee80211Ibm370ParchanEsconDlswIsdnsIsdnuLapdIpswitchRsrbAtmlogicalDs0Ds0BundleBscAsyncCnrIso88025DtrEplrsArapPropcnlsHostpadTermpadFramerelaympiX213AdslRadslSdslVdslIso88025CrfpintMyrinetVoiceemVoicefxoVoicefxsVoiceencapVoiceoveripAtmdxiAtmfuniAtmimaPppmultilinkbundleIpovercdlcIpoverclawStacktostackVirtualipaddressMpcIpoveratmIso88025FiberTdlcGigabitethernetHdlcLapfV37X25MlpX25HuntgroupTransphdlcInterleaveFastIpDocscablemaclayerDocscabledownstreamDocscableupstreamA12MppswitchTunnelCoffeeCesAtmsubinterfaceL2VlanL3IpvlanL3IpxvlanDigitalpowerlineMediamailoveripDtmDcnIpforwardMsdslIeee1394IfGsnDvbrccmaclayerDvbrccdownstreamDvbrccupstreamAtmvirtualMplstunnelSrpVoiceoveratmVoiceoverframerelayIdslCompositelinkSs7SiglinkPropwirelessp2PFrforwardRfc1483UsbIeee8023AdlagBgppolicyaccountingFrf16MfrbundleH323GatekeeperH323ProxyMplsMfsiglinkHdsl2ShdslDs1FdlPosDvbasiinDvbasioutPlcNfasTr008Gr303RdtGr303IdtIsupPropdocswirelessmaclayerPropdocswirelessdownstreamPropdocswirelessupstreamHiperlan2Propbwap2MpSonetoverheadchannelDigitalwrapperoverheadchannelAal2RadiomacAtmradioImtMvlReachdslFrdlciendptAtmvciendptOpticalchannelOpticaltransportPropatmVoiceovercableInfinibandTelinkQ2931VirtualtgSiptgSipsigDocscableupstreamchannelEconetPon155Pon622BridgeLinegroupVoiceemfgdVoicefgdeanaVoicedidMpegtransportSixtofourGtpPdnetherloop1Pdnetherloop2OpticalchannelgroupHomepnaGfpCiscoislvlanActelismetaloopFciplinkRprQamLmpCblvectastarDocscablemcmtsdownstreamAdsl2MacseccontrolledifMacsecuncontrolledifAviciopticaletherAtmbondVoicefgdosMocaversion1Ieee80216WmanAdsl2PlusDvbrcsmaclayerDvbtdmDvbrcstdmaX86LapsWwanppWwanpp2VoiceebsIfpwtypeIlanPipAluelpGponVdsl2Capwapdot11ProfileCapwapdot11BssCapwapwtpvirtualradioBitsDocscableupstreamrfportCabledownstreamrfportVmwarevirtualnicIeee802154OtnoduOtnotuIfvfitypeG9981G9982G9983AlueponAluepononuAlueponphysicaluniAlueponlogicallinkAlugpononuAlugponphysicaluniVmwarenicteam.
    // This attribute is mandatory.
    Type_ interface{}

    // The desired state of the interface.  This leaf has the same read semantics
    // as ifAdminStatus. The type is AdminStatus. This attribute is mandatory.
    AdminStatus interface{}

    // The current operational state of the interface. This leaf has the same
    // semantics as ifOperStatus. The type is OperStatus. This attribute is
    // mandatory.
    OperStatus interface{}

    // The time the interface entered its current operational state.  If the
    // current state was entered prior to the last re-initialization of the local
    // network management subsystem, then this node is not present. The type is
    // string with pattern:
    // b'\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}(\\.\\d+)?(Z|[\\+\\-]\\d{2}:\\d{2})'.
    LastChange interface{}

    // The ifIndex value for the ifEntry represented by this interface. The type
    // is interface{} with range: 1..2147483647. This attribute is mandatory.
    IfIndex interface{}

    // The interface's address at its protocol sub-layer.  For example, for an
    // 802.x interface, this object normally contains a Media Access Control (MAC)
    // address.  The interface's media-specific modules must define the bit and
    // byte ordering and the format of the value of this object.  For interfaces
    // that do not have such an address (e.g., a serial line), this node is not
    // present. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
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

func (self *InterfacesState_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "ietf"
    self.EntityData.ParentYangName = "interfaces-state"
    self.EntityData.SegmentPath = "interface" + "[name='" + fmt.Sprintf("%v", self.Name) + "']"
    self.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    self.EntityData.NamespaceTable = ietf.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Children["statistics"] = types.YChild{"Statistics", &self.Statistics}
    self.EntityData.Children["ietf-diffserv-target:diffserv-target-entry"] = types.YChild{"DiffservTargetEntry", nil}
    for i := range self.DiffservTargetEntry {
        self.EntityData.Children[types.GetSegmentPath(&self.DiffservTargetEntry[i])] = types.YChild{"DiffservTargetEntry", &self.DiffservTargetEntry[i]}
    }
    self.EntityData.Children["ietf-ip:ipv4"] = types.YChild{"Ipv4", &self.Ipv4}
    self.EntityData.Children["ietf-ip:ipv6"] = types.YChild{"Ipv6", &self.Ipv6}
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["name"] = types.YLeaf{"Name", self.Name}
    self.EntityData.Leafs["type"] = types.YLeaf{"Type_", self.Type_}
    self.EntityData.Leafs["admin-status"] = types.YLeaf{"AdminStatus", self.AdminStatus}
    self.EntityData.Leafs["oper-status"] = types.YLeaf{"OperStatus", self.OperStatus}
    self.EntityData.Leafs["last-change"] = types.YLeaf{"LastChange", self.LastChange}
    self.EntityData.Leafs["if-index"] = types.YLeaf{"IfIndex", self.IfIndex}
    self.EntityData.Leafs["phys-address"] = types.YLeaf{"PhysAddress", self.PhysAddress}
    self.EntityData.Leafs["higher-layer-if"] = types.YLeaf{"HigherLayerIf", self.HigherLayerIf}
    self.EntityData.Leafs["lower-layer-if"] = types.YLeaf{"LowerLayerIf", self.LowerLayerIf}
    self.EntityData.Leafs["speed"] = types.YLeaf{"Speed", self.Speed}
    self.EntityData.Leafs["routing-instance"] = types.YLeaf{"RoutingInstance", self.RoutingInstance}
    return &(self.EntityData)
}

// InterfacesState_Interface_Statistics
// A collection of interface-related statistics objects.
type InterfacesState_Interface_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The time on the most recent occasion at which any one or more of this
    // interface's counters suffered a discontinuity.  If no such discontinuities
    // have occurred since the last re-initialization of the local management
    // subsystem, then this node contains the time the local management subsystem
    // re-initialized itself. The type is string with pattern:
    // b'\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}(\\.\\d+)?(Z|[\\+\\-]\\d{2}:\\d{2})'.
    // This attribute is mandatory.
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
    // could not be transmitted because of errors.  Discontinuities in the value
    // of this counter can occur at re-initialization of the management system,
    // and at other times as indicated by the value of 'discontinuity-time'. The
    // type is interface{} with range: 0..4294967295.
    OutErrors interface{}

    // total packets input. The type is interface{} with range:
    // 0..18446744073709551615.
    InPkts interface{}

    // total packets output. The type is interface{} with range:
    // 0..18446744073709551615.
    OutPkts interface{}
}

func (statistics *InterfacesState_Interface_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "ietf"
    statistics.EntityData.ParentYangName = "interface"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    statistics.EntityData.NamespaceTable = ietf.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    statistics.EntityData.Children = make(map[string]types.YChild)
    statistics.EntityData.Leafs = make(map[string]types.YLeaf)
    statistics.EntityData.Leafs["discontinuity-time"] = types.YLeaf{"DiscontinuityTime", statistics.DiscontinuityTime}
    statistics.EntityData.Leafs["in-octets"] = types.YLeaf{"InOctets", statistics.InOctets}
    statistics.EntityData.Leafs["in-unicast-pkts"] = types.YLeaf{"InUnicastPkts", statistics.InUnicastPkts}
    statistics.EntityData.Leafs["in-broadcast-pkts"] = types.YLeaf{"InBroadcastPkts", statistics.InBroadcastPkts}
    statistics.EntityData.Leafs["in-multicast-pkts"] = types.YLeaf{"InMulticastPkts", statistics.InMulticastPkts}
    statistics.EntityData.Leafs["in-discards"] = types.YLeaf{"InDiscards", statistics.InDiscards}
    statistics.EntityData.Leafs["in-errors"] = types.YLeaf{"InErrors", statistics.InErrors}
    statistics.EntityData.Leafs["in-unknown-protos"] = types.YLeaf{"InUnknownProtos", statistics.InUnknownProtos}
    statistics.EntityData.Leafs["out-octets"] = types.YLeaf{"OutOctets", statistics.OutOctets}
    statistics.EntityData.Leafs["out-unicast-pkts"] = types.YLeaf{"OutUnicastPkts", statistics.OutUnicastPkts}
    statistics.EntityData.Leafs["out-broadcast-pkts"] = types.YLeaf{"OutBroadcastPkts", statistics.OutBroadcastPkts}
    statistics.EntityData.Leafs["out-multicast-pkts"] = types.YLeaf{"OutMulticastPkts", statistics.OutMulticastPkts}
    statistics.EntityData.Leafs["out-discards"] = types.YLeaf{"OutDiscards", statistics.OutDiscards}
    statistics.EntityData.Leafs["out-errors"] = types.YLeaf{"OutErrors", statistics.OutErrors}
    statistics.EntityData.Leafs["in-pkts"] = types.YLeaf{"InPkts", statistics.InPkts}
    statistics.EntityData.Leafs["out-pkts"] = types.YLeaf{"OutPkts", statistics.OutPkts}
    return &(statistics.EntityData)
}

// InterfacesState_Interface_DiffservTargetEntry
// policy target for inbound or outbound direction
type InterfacesState_Interface_DiffservTargetEntry struct {
    EntityData types.CommonEntityData
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

func (diffservTargetEntry *InterfacesState_Interface_DiffservTargetEntry) GetEntityData() *types.CommonEntityData {
    diffservTargetEntry.EntityData.YFilter = diffservTargetEntry.YFilter
    diffservTargetEntry.EntityData.YangName = "diffserv-target-entry"
    diffservTargetEntry.EntityData.BundleName = "ietf"
    diffservTargetEntry.EntityData.ParentYangName = "interface"
    diffservTargetEntry.EntityData.SegmentPath = "ietf-diffserv-target:diffserv-target-entry" + "[direction='" + fmt.Sprintf("%v", diffservTargetEntry.Direction) + "']" + "[policy-name='" + fmt.Sprintf("%v", diffservTargetEntry.PolicyName) + "']"
    diffservTargetEntry.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    diffservTargetEntry.EntityData.NamespaceTable = ietf.GetNamespaces()
    diffservTargetEntry.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    diffservTargetEntry.EntityData.Children = make(map[string]types.YChild)
    diffservTargetEntry.EntityData.Children["diffserv-target-classifier-statistics"] = types.YChild{"DiffservTargetClassifierStatistics", nil}
    for i := range diffservTargetEntry.DiffservTargetClassifierStatistics {
        diffservTargetEntry.EntityData.Children[types.GetSegmentPath(&diffservTargetEntry.DiffservTargetClassifierStatistics[i])] = types.YChild{"DiffservTargetClassifierStatistics", &diffservTargetEntry.DiffservTargetClassifierStatistics[i]}
    }
    diffservTargetEntry.EntityData.Leafs = make(map[string]types.YLeaf)
    diffservTargetEntry.EntityData.Leafs["direction"] = types.YLeaf{"Direction", diffservTargetEntry.Direction}
    diffservTargetEntry.EntityData.Leafs["policy-name"] = types.YLeaf{"PolicyName", diffservTargetEntry.PolicyName}
    return &(diffservTargetEntry.EntityData)
}

// InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics
// Statistics for each Classifier Entry in a Policy
type InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics struct {
    EntityData types.CommonEntityData
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

func (diffservTargetClassifierStatistics *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics) GetEntityData() *types.CommonEntityData {
    diffservTargetClassifierStatistics.EntityData.YFilter = diffservTargetClassifierStatistics.YFilter
    diffservTargetClassifierStatistics.EntityData.YangName = "diffserv-target-classifier-statistics"
    diffservTargetClassifierStatistics.EntityData.BundleName = "ietf"
    diffservTargetClassifierStatistics.EntityData.ParentYangName = "diffserv-target-entry"
    diffservTargetClassifierStatistics.EntityData.SegmentPath = "diffserv-target-classifier-statistics" + "[classifier-entry-name='" + fmt.Sprintf("%v", diffservTargetClassifierStatistics.ClassifierEntryName) + "']" + "[parent-path='" + fmt.Sprintf("%v", diffservTargetClassifierStatistics.ParentPath) + "']"
    diffservTargetClassifierStatistics.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    diffservTargetClassifierStatistics.EntityData.NamespaceTable = ietf.GetNamespaces()
    diffservTargetClassifierStatistics.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    diffservTargetClassifierStatistics.EntityData.Children = make(map[string]types.YChild)
    diffservTargetClassifierStatistics.EntityData.Children["classifier-entry-statistics"] = types.YChild{"ClassifierEntryStatistics", &diffservTargetClassifierStatistics.ClassifierEntryStatistics}
    diffservTargetClassifierStatistics.EntityData.Children["meter-statistics"] = types.YChild{"MeterStatistics", nil}
    for i := range diffservTargetClassifierStatistics.MeterStatistics {
        diffservTargetClassifierStatistics.EntityData.Children[types.GetSegmentPath(&diffservTargetClassifierStatistics.MeterStatistics[i])] = types.YChild{"MeterStatistics", &diffservTargetClassifierStatistics.MeterStatistics[i]}
    }
    diffservTargetClassifierStatistics.EntityData.Children["queuing-statistics"] = types.YChild{"QueuingStatistics", &diffservTargetClassifierStatistics.QueuingStatistics}
    diffservTargetClassifierStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    diffservTargetClassifierStatistics.EntityData.Leafs["classifier-entry-name"] = types.YLeaf{"ClassifierEntryName", diffservTargetClassifierStatistics.ClassifierEntryName}
    diffservTargetClassifierStatistics.EntityData.Leafs["parent-path"] = types.YLeaf{"ParentPath", diffservTargetClassifierStatistics.ParentPath}
    return &(diffservTargetClassifierStatistics.EntityData)
}

// InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_ClassifierEntryStatistics
// 
// This group defines the classifier filter statistics of 
// each classifier entry
//        
// 
type InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_ClassifierEntryStatistics struct {
    EntityData types.CommonEntityData
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

func (classifierEntryStatistics *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_ClassifierEntryStatistics) GetEntityData() *types.CommonEntityData {
    classifierEntryStatistics.EntityData.YFilter = classifierEntryStatistics.YFilter
    classifierEntryStatistics.EntityData.YangName = "classifier-entry-statistics"
    classifierEntryStatistics.EntityData.BundleName = "ietf"
    classifierEntryStatistics.EntityData.ParentYangName = "diffserv-target-classifier-statistics"
    classifierEntryStatistics.EntityData.SegmentPath = "classifier-entry-statistics"
    classifierEntryStatistics.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    classifierEntryStatistics.EntityData.NamespaceTable = ietf.GetNamespaces()
    classifierEntryStatistics.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    classifierEntryStatistics.EntityData.Children = make(map[string]types.YChild)
    classifierEntryStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    classifierEntryStatistics.EntityData.Leafs["classified-pkts"] = types.YLeaf{"ClassifiedPkts", classifierEntryStatistics.ClassifiedPkts}
    classifierEntryStatistics.EntityData.Leafs["classified-bytes"] = types.YLeaf{"ClassifiedBytes", classifierEntryStatistics.ClassifiedBytes}
    classifierEntryStatistics.EntityData.Leafs["classified-rate"] = types.YLeaf{"ClassifiedRate", classifierEntryStatistics.ClassifiedRate}
    return &(classifierEntryStatistics.EntityData)
}

// InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_MeterStatistics
// Meter statistics
type InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_MeterStatistics struct {
    EntityData types.CommonEntityData
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

func (meterStatistics *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_MeterStatistics) GetEntityData() *types.CommonEntityData {
    meterStatistics.EntityData.YFilter = meterStatistics.YFilter
    meterStatistics.EntityData.YangName = "meter-statistics"
    meterStatistics.EntityData.BundleName = "ietf"
    meterStatistics.EntityData.ParentYangName = "diffserv-target-classifier-statistics"
    meterStatistics.EntityData.SegmentPath = "meter-statistics" + "[meter-id='" + fmt.Sprintf("%v", meterStatistics.MeterId) + "']"
    meterStatistics.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    meterStatistics.EntityData.NamespaceTable = ietf.GetNamespaces()
    meterStatistics.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    meterStatistics.EntityData.Children = make(map[string]types.YChild)
    meterStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    meterStatistics.EntityData.Leafs["meter-id"] = types.YLeaf{"MeterId", meterStatistics.MeterId}
    meterStatistics.EntityData.Leafs["meter-succeed-pkts"] = types.YLeaf{"MeterSucceedPkts", meterStatistics.MeterSucceedPkts}
    meterStatistics.EntityData.Leafs["meter-succeed-bytes"] = types.YLeaf{"MeterSucceedBytes", meterStatistics.MeterSucceedBytes}
    meterStatistics.EntityData.Leafs["meter-failed-pkts"] = types.YLeaf{"MeterFailedPkts", meterStatistics.MeterFailedPkts}
    meterStatistics.EntityData.Leafs["meter-failed-bytes"] = types.YLeaf{"MeterFailedBytes", meterStatistics.MeterFailedBytes}
    return &(meterStatistics.EntityData)
}

// InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics
// queue related statistics 
type InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics struct {
    EntityData types.CommonEntityData
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

func (queuingStatistics *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics) GetEntityData() *types.CommonEntityData {
    queuingStatistics.EntityData.YFilter = queuingStatistics.YFilter
    queuingStatistics.EntityData.YangName = "queuing-statistics"
    queuingStatistics.EntityData.BundleName = "ietf"
    queuingStatistics.EntityData.ParentYangName = "diffserv-target-classifier-statistics"
    queuingStatistics.EntityData.SegmentPath = "queuing-statistics"
    queuingStatistics.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    queuingStatistics.EntityData.NamespaceTable = ietf.GetNamespaces()
    queuingStatistics.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    queuingStatistics.EntityData.Children = make(map[string]types.YChild)
    queuingStatistics.EntityData.Children["wred-stats"] = types.YChild{"WredStats", &queuingStatistics.WredStats}
    queuingStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    queuingStatistics.EntityData.Leafs["output-pkts"] = types.YLeaf{"OutputPkts", queuingStatistics.OutputPkts}
    queuingStatistics.EntityData.Leafs["output-bytes"] = types.YLeaf{"OutputBytes", queuingStatistics.OutputBytes}
    queuingStatistics.EntityData.Leafs["queue-size-pkts"] = types.YLeaf{"QueueSizePkts", queuingStatistics.QueueSizePkts}
    queuingStatistics.EntityData.Leafs["queue-size-bytes"] = types.YLeaf{"QueueSizeBytes", queuingStatistics.QueueSizeBytes}
    queuingStatistics.EntityData.Leafs["drop-pkts"] = types.YLeaf{"DropPkts", queuingStatistics.DropPkts}
    queuingStatistics.EntityData.Leafs["drop-bytes"] = types.YLeaf{"DropBytes", queuingStatistics.DropBytes}
    return &(queuingStatistics.EntityData)
}

// InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics_WredStats
// Container for WRED statistics
type InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics_WredStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Early drop packets . The type is interface{} with range:
    // 0..18446744073709551615.
    EarlyDropPkts interface{}

    // Early drop bytes . The type is interface{} with range:
    // 0..18446744073709551615.
    EarlyDropBytes interface{}
}

func (wredStats *InterfacesState_Interface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics_WredStats) GetEntityData() *types.CommonEntityData {
    wredStats.EntityData.YFilter = wredStats.YFilter
    wredStats.EntityData.YangName = "wred-stats"
    wredStats.EntityData.BundleName = "ietf"
    wredStats.EntityData.ParentYangName = "queuing-statistics"
    wredStats.EntityData.SegmentPath = "wred-stats"
    wredStats.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    wredStats.EntityData.NamespaceTable = ietf.GetNamespaces()
    wredStats.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    wredStats.EntityData.Children = make(map[string]types.YChild)
    wredStats.EntityData.Leafs = make(map[string]types.YLeaf)
    wredStats.EntityData.Leafs["early-drop-pkts"] = types.YLeaf{"EarlyDropPkts", wredStats.EarlyDropPkts}
    wredStats.EntityData.Leafs["early-drop-bytes"] = types.YLeaf{"EarlyDropBytes", wredStats.EarlyDropBytes}
    return &(wredStats.EntityData)
}

// InterfacesState_Interface_Ipv4
// Interface-specific parameters for the IPv4 address family.
// This type is a presence type.
type InterfacesState_Interface_Ipv4 struct {
    EntityData types.CommonEntityData
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

func (ipv4 *InterfacesState_Interface_Ipv4) GetEntityData() *types.CommonEntityData {
    ipv4.EntityData.YFilter = ipv4.YFilter
    ipv4.EntityData.YangName = "ipv4"
    ipv4.EntityData.BundleName = "ietf"
    ipv4.EntityData.ParentYangName = "interface"
    ipv4.EntityData.SegmentPath = "ietf-ip:ipv4"
    ipv4.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    ipv4.EntityData.NamespaceTable = ietf.GetNamespaces()
    ipv4.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    ipv4.EntityData.Children = make(map[string]types.YChild)
    ipv4.EntityData.Children["address"] = types.YChild{"Address", nil}
    for i := range ipv4.Address {
        ipv4.EntityData.Children[types.GetSegmentPath(&ipv4.Address[i])] = types.YChild{"Address", &ipv4.Address[i]}
    }
    ipv4.EntityData.Children["neighbor"] = types.YChild{"Neighbor", nil}
    for i := range ipv4.Neighbor {
        ipv4.EntityData.Children[types.GetSegmentPath(&ipv4.Neighbor[i])] = types.YChild{"Neighbor", &ipv4.Neighbor[i]}
    }
    ipv4.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv4.EntityData.Leafs["forwarding"] = types.YLeaf{"Forwarding", ipv4.Forwarding}
    ipv4.EntityData.Leafs["mtu"] = types.YLeaf{"Mtu", ipv4.Mtu}
    return &(ipv4.EntityData)
}

// InterfacesState_Interface_Ipv4_Address
// The list of IPv4 addresses on the interface.
type InterfacesState_Interface_Ipv4_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The IPv4 address on the interface. The type is
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ip interface{}

    // The length of the subnet prefix. The type is interface{} with range: 0..32.
    PrefixLength interface{}

    // The subnet specified as a netmask. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])'.
    Netmask interface{}

    // The origin of this address. The type is IpAddressOrigin.
    Origin interface{}
}

func (address *InterfacesState_Interface_Ipv4_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "ietf"
    address.EntityData.ParentYangName = "ipv4"
    address.EntityData.SegmentPath = "address" + "[ip='" + fmt.Sprintf("%v", address.Ip) + "']"
    address.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    address.EntityData.NamespaceTable = ietf.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    address.EntityData.Children = make(map[string]types.YChild)
    address.EntityData.Leafs = make(map[string]types.YLeaf)
    address.EntityData.Leafs["ip"] = types.YLeaf{"Ip", address.Ip}
    address.EntityData.Leafs["prefix-length"] = types.YLeaf{"PrefixLength", address.PrefixLength}
    address.EntityData.Leafs["netmask"] = types.YLeaf{"Netmask", address.Netmask}
    address.EntityData.Leafs["origin"] = types.YLeaf{"Origin", address.Origin}
    return &(address.EntityData)
}

// InterfacesState_Interface_Ipv4_Neighbor
// A list of mappings from IPv4 addresses to
// link-layer addresses.
// This list represents the ARP Cache.
type InterfacesState_Interface_Ipv4_Neighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The IPv4 address of the neighbor node. The type is
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ip interface{}

    // The link-layer address of the neighbor node. The type is string with
    // pattern: b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    LinkLayerAddress interface{}

    // The origin of this neighbor entry. The type is NeighborOrigin.
    Origin interface{}
}

func (neighbor *InterfacesState_Interface_Ipv4_Neighbor) GetEntityData() *types.CommonEntityData {
    neighbor.EntityData.YFilter = neighbor.YFilter
    neighbor.EntityData.YangName = "neighbor"
    neighbor.EntityData.BundleName = "ietf"
    neighbor.EntityData.ParentYangName = "ipv4"
    neighbor.EntityData.SegmentPath = "neighbor" + "[ip='" + fmt.Sprintf("%v", neighbor.Ip) + "']"
    neighbor.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    neighbor.EntityData.NamespaceTable = ietf.GetNamespaces()
    neighbor.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    neighbor.EntityData.Children = make(map[string]types.YChild)
    neighbor.EntityData.Leafs = make(map[string]types.YLeaf)
    neighbor.EntityData.Leafs["ip"] = types.YLeaf{"Ip", neighbor.Ip}
    neighbor.EntityData.Leafs["link-layer-address"] = types.YLeaf{"LinkLayerAddress", neighbor.LinkLayerAddress}
    neighbor.EntityData.Leafs["origin"] = types.YLeaf{"Origin", neighbor.Origin}
    return &(neighbor.EntityData)
}

// InterfacesState_Interface_Ipv6
// Parameters for the IPv6 address family.
// This type is a presence type.
type InterfacesState_Interface_Ipv6 struct {
    EntityData types.CommonEntityData
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

func (ipv6 *InterfacesState_Interface_Ipv6) GetEntityData() *types.CommonEntityData {
    ipv6.EntityData.YFilter = ipv6.YFilter
    ipv6.EntityData.YangName = "ipv6"
    ipv6.EntityData.BundleName = "ietf"
    ipv6.EntityData.ParentYangName = "interface"
    ipv6.EntityData.SegmentPath = "ietf-ip:ipv6"
    ipv6.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    ipv6.EntityData.NamespaceTable = ietf.GetNamespaces()
    ipv6.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    ipv6.EntityData.Children = make(map[string]types.YChild)
    ipv6.EntityData.Children["address"] = types.YChild{"Address", nil}
    for i := range ipv6.Address {
        ipv6.EntityData.Children[types.GetSegmentPath(&ipv6.Address[i])] = types.YChild{"Address", &ipv6.Address[i]}
    }
    ipv6.EntityData.Children["neighbor"] = types.YChild{"Neighbor", nil}
    for i := range ipv6.Neighbor {
        ipv6.EntityData.Children[types.GetSegmentPath(&ipv6.Neighbor[i])] = types.YChild{"Neighbor", &ipv6.Neighbor[i]}
    }
    ipv6.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv6.EntityData.Leafs["forwarding"] = types.YLeaf{"Forwarding", ipv6.Forwarding}
    ipv6.EntityData.Leafs["mtu"] = types.YLeaf{"Mtu", ipv6.Mtu}
    return &(ipv6.EntityData)
}

// InterfacesState_Interface_Ipv6_Address
// The list of IPv6 addresses on the interface.
type InterfacesState_Interface_Ipv6_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The IPv6 address on the interface. The type is
    // string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
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

func (address *InterfacesState_Interface_Ipv6_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "ietf"
    address.EntityData.ParentYangName = "ipv6"
    address.EntityData.SegmentPath = "address" + "[ip='" + fmt.Sprintf("%v", address.Ip) + "']"
    address.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    address.EntityData.NamespaceTable = ietf.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    address.EntityData.Children = make(map[string]types.YChild)
    address.EntityData.Leafs = make(map[string]types.YLeaf)
    address.EntityData.Leafs["ip"] = types.YLeaf{"Ip", address.Ip}
    address.EntityData.Leafs["prefix-length"] = types.YLeaf{"PrefixLength", address.PrefixLength}
    address.EntityData.Leafs["origin"] = types.YLeaf{"Origin", address.Origin}
    address.EntityData.Leafs["status"] = types.YLeaf{"Status", address.Status}
    return &(address.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The IPv6 address of the neighbor node. The type is
    // string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ip interface{}

    // The link-layer address of the neighbor node. The type is string with
    // pattern: b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    LinkLayerAddress interface{}

    // The origin of this neighbor entry. The type is NeighborOrigin.
    Origin interface{}

    // Indicates that the neighbor node acts as a router. The type is interface{}.
    IsRouter interface{}

    // The Neighbor Unreachability Detection state of this entry. The type is
    // State.
    State interface{}
}

func (neighbor *InterfacesState_Interface_Ipv6_Neighbor) GetEntityData() *types.CommonEntityData {
    neighbor.EntityData.YFilter = neighbor.YFilter
    neighbor.EntityData.YangName = "neighbor"
    neighbor.EntityData.BundleName = "ietf"
    neighbor.EntityData.ParentYangName = "ipv6"
    neighbor.EntityData.SegmentPath = "neighbor" + "[ip='" + fmt.Sprintf("%v", neighbor.Ip) + "']"
    neighbor.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    neighbor.EntityData.NamespaceTable = ietf.GetNamespaces()
    neighbor.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    neighbor.EntityData.Children = make(map[string]types.YChild)
    neighbor.EntityData.Leafs = make(map[string]types.YLeaf)
    neighbor.EntityData.Leafs["ip"] = types.YLeaf{"Ip", neighbor.Ip}
    neighbor.EntityData.Leafs["link-layer-address"] = types.YLeaf{"LinkLayerAddress", neighbor.LinkLayerAddress}
    neighbor.EntityData.Leafs["origin"] = types.YLeaf{"Origin", neighbor.Origin}
    neighbor.EntityData.Leafs["is-router"] = types.YLeaf{"IsRouter", neighbor.IsRouter}
    neighbor.EntityData.Leafs["state"] = types.YLeaf{"State", neighbor.State}
    return &(neighbor.EntityData)
}

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

