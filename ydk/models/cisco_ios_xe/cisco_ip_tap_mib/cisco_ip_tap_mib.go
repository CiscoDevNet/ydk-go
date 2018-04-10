// This module manages Cisco's intercept feature for IP.
// 
// This MIB is used along with CISCO-TAP2-MIB to
// intercept IP traffic. CISCO-TAP2-MIB along with
// specific filter MIBs like this MIB replace 
// CISCO-TAP-MIB.
// 
// To create an IP intercept, an entry citapStreamEntry 
// is created which contains the filter details. An entry
// cTap2StreamEntry of CISCO-TAP2-MIB is created, which is
// the common stream information for all kinds of 
// intercepts and type of the specific stream is set to
// ip in this entry.
package cisco_ip_tap_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_ip_tap_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-IP-TAP-MIB CISCO-IP-TAP-MIB}", reflect.TypeOf(CISCOIPTAPMIB{}))
    ydk.RegisterEntity("CISCO-IP-TAP-MIB:CISCO-IP-TAP-MIB", reflect.TypeOf(CISCOIPTAPMIB{}))
}

// CISCOIPTAPMIB
type CISCOIPTAPMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Citapstreamencodepacket CISCOIPTAPMIB_Citapstreamencodepacket

    // The Intercept Stream IP Table lists the IPv4 and IPv6 streams to be
    // intercepted.  The same data stream may be required by multiple taps, and
    // one might assume that often the intercepted stream is a small subset of the
    // traffic that could be intercepted.   This essentially provides options for
    // packet selection, only some of which might be used. For example, if all
    // traffic to or from a given interface is to be intercepted, one would
    // configure an entry which lists the interface, and wild-card everything
    // else.  If all traffic to or from a given IP Address is to be intercepted,
    // one would configure two such entries listing the IP Address as source and
    // destination respectively, and wild-card everything else.  If a particular
    // voice on a teleconference is to be intercepted, on the other hand, one
    // would extract the multicast (destination) IP address, the source IP
    // Address, the protocol (UDP), and the source and destination ports from the
    // call control exchange and list all necessary information.   The first index
    // indicates which Mediation Device the intercepted traffic will be diverted
    // to. The second index permits multiple classifiers to be used together, such
    // as having an IP address as source or destination. The value of the second
    // index is that of the stream's counter entry in the  cTap2StreamTable. 
    // Entries are added to this table via citapStreamStatus in  accordance with
    // the RowStatus convention.
    Citapstreamtable CISCOIPTAPMIB_Citapstreamtable
}

func (cISCOIPTAPMIB *CISCOIPTAPMIB) GetEntityData() *types.CommonEntityData {
    cISCOIPTAPMIB.EntityData.YFilter = cISCOIPTAPMIB.YFilter
    cISCOIPTAPMIB.EntityData.YangName = "CISCO-IP-TAP-MIB"
    cISCOIPTAPMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOIPTAPMIB.EntityData.ParentYangName = "CISCO-IP-TAP-MIB"
    cISCOIPTAPMIB.EntityData.SegmentPath = "CISCO-IP-TAP-MIB:CISCO-IP-TAP-MIB"
    cISCOIPTAPMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOIPTAPMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOIPTAPMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOIPTAPMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOIPTAPMIB.EntityData.Children["citapStreamEncodePacket"] = types.YChild{"Citapstreamencodepacket", &cISCOIPTAPMIB.Citapstreamencodepacket}
    cISCOIPTAPMIB.EntityData.Children["citapStreamTable"] = types.YChild{"Citapstreamtable", &cISCOIPTAPMIB.Citapstreamtable}
    cISCOIPTAPMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOIPTAPMIB.EntityData)
}

// CISCOIPTAPMIB_Citapstreamencodepacket
type CISCOIPTAPMIB_Citapstreamencodepacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object displays what types of intercept streams can be configured on
    // this type of device. This may be dependent on hardware capabilities,
    // software capabilities. The following fields may be supported:    
    // tapEnable:   set if table entries with                 
    // cTap2StreamInterceptEnable set to 'false'                  are used to
    // pre-screen packets for intercept;                  otherwise these entries
    // are ignored.     interface:   SNMP ifIndex Value may be used to select     
    // interception of all data crossing an                  interface or set of
    // interfaces.     ipV4:        IPv4 Address or prefix may be used to select  
    // traffic to be intercepted.     ipV6:        IPv6 Address or prefix may be
    // used to select                  traffic to be intercepted.     l4Port:     
    // TCP/UDP Ports may be used to select traffic                  to be
    // intercepted.     dscp:        DSCP (Differentiated Services Code Point) may
    // be used to select traffic to be intercepted.     voip:        packets
    // belonging to a voice session may                  be intercepted using
    // source IPv4 address and                  source UDP port. The type is
    // map[string]bool.
    Citapstreamcapabilities interface{}
}

func (citapstreamencodepacket *CISCOIPTAPMIB_Citapstreamencodepacket) GetEntityData() *types.CommonEntityData {
    citapstreamencodepacket.EntityData.YFilter = citapstreamencodepacket.YFilter
    citapstreamencodepacket.EntityData.YangName = "citapStreamEncodePacket"
    citapstreamencodepacket.EntityData.BundleName = "cisco_ios_xe"
    citapstreamencodepacket.EntityData.ParentYangName = "CISCO-IP-TAP-MIB"
    citapstreamencodepacket.EntityData.SegmentPath = "citapStreamEncodePacket"
    citapstreamencodepacket.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    citapstreamencodepacket.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    citapstreamencodepacket.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    citapstreamencodepacket.EntityData.Children = make(map[string]types.YChild)
    citapstreamencodepacket.EntityData.Leafs = make(map[string]types.YLeaf)
    citapstreamencodepacket.EntityData.Leafs["citapStreamCapabilities"] = types.YLeaf{"Citapstreamcapabilities", citapstreamencodepacket.Citapstreamcapabilities}
    return &(citapstreamencodepacket.EntityData)
}

// CISCOIPTAPMIB_Citapstreamtable
// The Intercept Stream IP Table lists the IPv4 and IPv6 streams
// to be intercepted.  The same data stream may be required by
// multiple taps, and one might assume that often the intercepted
// stream is a small subset of the traffic that could be
// intercepted.
// 
// 
// This essentially provides options for packet selection, only
// some of which might be used. For example, if all traffic to or
// from a given interface is to be intercepted, one would
// configure an entry which lists the interface, and wild-card
// everything else.  If all traffic to or from a given IP Address
// is to be intercepted, one would configure two such entries
// listing the IP Address as source and destination respectively,
// and wild-card everything else.  If a particular voice on a
// teleconference is to be intercepted, on the other hand, one
// would extract the multicast (destination) IP address, the
// source IP Address, the protocol (UDP), and the source and
// destination ports from the call control exchange and list all
// necessary information.
// 
// 
// The first index indicates which Mediation Device the
// intercepted traffic will be diverted to. The second index
// permits multiple classifiers to be used together, such as
// having an IP address as source or destination. The value of the
// second index is that of the stream's counter entry in the 
// cTap2StreamTable.
// 
// Entries are added to this table via citapStreamStatus in 
// accordance with the RowStatus convention.
type CISCOIPTAPMIB_Citapstreamtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A stream entry indicates a single data stream to be intercepted to a
    // Mediation Device. Many selected data streams may go to the same application
    // interface, and many application interfaces are supported. The type is slice
    // of CISCOIPTAPMIB_Citapstreamtable_Citapstreamentry.
    Citapstreamentry []CISCOIPTAPMIB_Citapstreamtable_Citapstreamentry
}

func (citapstreamtable *CISCOIPTAPMIB_Citapstreamtable) GetEntityData() *types.CommonEntityData {
    citapstreamtable.EntityData.YFilter = citapstreamtable.YFilter
    citapstreamtable.EntityData.YangName = "citapStreamTable"
    citapstreamtable.EntityData.BundleName = "cisco_ios_xe"
    citapstreamtable.EntityData.ParentYangName = "CISCO-IP-TAP-MIB"
    citapstreamtable.EntityData.SegmentPath = "citapStreamTable"
    citapstreamtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    citapstreamtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    citapstreamtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    citapstreamtable.EntityData.Children = make(map[string]types.YChild)
    citapstreamtable.EntityData.Children["citapStreamEntry"] = types.YChild{"Citapstreamentry", nil}
    for i := range citapstreamtable.Citapstreamentry {
        citapstreamtable.EntityData.Children[types.GetSegmentPath(&citapstreamtable.Citapstreamentry[i])] = types.YChild{"Citapstreamentry", &citapstreamtable.Citapstreamentry[i]}
    }
    citapstreamtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(citapstreamtable.EntityData)
}

// CISCOIPTAPMIB_Citapstreamtable_Citapstreamentry
// A stream entry indicates a single data stream to be
// intercepted to a Mediation Device. Many selected data
// streams may go to the same application interface, and many
// application interfaces are supported.
type CISCOIPTAPMIB_Citapstreamtable_Citapstreamentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // cisco_tap2_mib.CISCOTAP2MIB_Ctap2Mediationtable_Ctap2Mediationentry_Ctap2Mediationcontentid
    Ctap2Mediationcontentid interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // cisco_tap2_mib.CISCOTAP2MIB_Ctap2Streamtable_Ctap2Streamentry_Ctap2Streamindex
    Ctap2Streamindex interface{}

    // The ifIndex value of the interface over which traffic to be intercepted is
    // received or transmitted. The interface may be physical or virtual. If this
    // is the only parameter specified, and it is other than -2, -1 or 0, all
    // traffic on the selected interface will be chosen.   If the value is zero,
    // matching traffic may be received or transmitted on any interface. 
    // Additional selection parameters must be selected to limit the scope of
    // traffic intercepted. This is most useful on non-routing platforms or on
    // intercepts placed elsewhere than a subscriber interface.   If the value is
    // -1, one or both of citapStreamDestinationAddress and
    // citapStreamSourceAddress must be specified with prefix length greater than
    // zero. Matching traffic on the interface pointed to by ipRouteIfIndex or
    // ipCidrRouteIfIndex values associated with those values is intercepted,
    // whichever is specified to be more focused than a default route.  If routing
    // changes, either by operator action or by routing protocol events, the
    // interface will change with it. This is primarily intended for use on
    // subscriber interfaces and other places where routing is guaranteed to be
    // symmetrical.   In both of these cases, it is possible to have the same
    // packet selected for intersection on both its ingress and egress interface. 
    // Nonetheless, only one instance of the packet is sent to the Mediation
    // Device.   If the value is -2, packets belonging to a Voice over IP (VoIP)
    // session identified by citapStreamSourceAddress,  citapStreamSourceLen and
    // citapStreamSourceL4PortMin may be  intercepted, as a specific voice session
    // can be identified  with source IP address and udp port number. Other
    // selection  parameters may be not considered, even if they are set by  the
    // Mediation Device.   This value must be set when creating a stream entry,
    // either to select an interface, to select all interfaces, or to select the
    // interface that routing chooses. Some platforms may not implement the entire
    // range of options. The type is interface{} with range: -2..2147483647.
    Citapstreaminterface interface{}

    // The type of address, used in packet selection. The type is InetAddressType.
    Citapstreamaddrtype interface{}

    // The Destination address or prefix used in packet selection. This address
    // will be of the type specified in citapStreamAddrType. The type is string
    // with length: 0..255.
    Citapstreamdestinationaddress interface{}

    // The length of the Destination Prefix. A value of zero causes all addresses
    // to match.  This prefix length will be consistent with the type specified in
    // citapStreamAddrType. The type is interface{} with range: 0..2040.
    Citapstreamdestinationlength interface{}

    // The Source Address used in packet selection. This address will be of the
    // type specified in citapStreamAddrType. The type is string with length:
    // 0..255.
    Citapstreamsourceaddress interface{}

    // The length of the Source Prefix. A value of zero causes all addresses to
    // match. This prefix length will be consistent with the type specified in
    // citapStreamAddrType. The type is interface{} with range: 0..2040.
    Citapstreamsourcelength interface{}

    // The value of the TOS byte, when masked with citapStreamTosByteMask, of
    // traffic to be intercepted.  If
    // citapStreamTosByte&(~citapStreamTosByteMask)!=0, configuration is rejected.
    // The type is interface{} with range: 0..255.
    Citapstreamtosbyte interface{}

    // The value of the TOS byte in an IPv4 or IPv6 header is ANDed with
    // citapStreamTosByteMask and compared with citapStreamTosByte.  If the values
    // are equal, the comparison is equal. If the mask is zero and the TosByte
    // value is zero, the result is to always accept. The type is interface{} with
    // range: 0..255.
    Citapstreamtosbytemask interface{}

    // The flow identifier in an IPv6 header. -1 indicates that the Flow Id is
    // unused. The type is interface{} with range: -1..1048575.
    Citapstreamflowid interface{}

    // The IP protocol to match against the IPv4 protocol number or the IPv6 Next-
    // Header number in the packet. -1 means 'any IP protocol'. The type is
    // interface{} with range: -1..255.
    Citapstreamprotocol interface{}

    // The minimum value that the layer-4 destination port number in the packet
    // must have in order to match.  This value must be equal to or less than the
    // value specified for this entry in citapStreamDestL4PortMax.   If both
    // citapStreamDestL4PortMin and citapStreamDestL4PortMax are at their default
    // values, the port number is effectively unused. The type is interface{} with
    // range: 0..65535.
    Citapstreamdestl4Portmin interface{}

    // The maximum value that the layer-4 destination port number in the packet
    // must have in order to match this classifier entry. This value must be equal
    // to or greater than the value specified for this entry in
    // citapStreamDestL4PortMin.   If both citapStreamDestL4PortMin and
    // citapStreamDestL4PortMax are at their default values, the port number is
    // effectively unused. The type is interface{} with range: 0..65535.
    Citapstreamdestl4Portmax interface{}

    // The minimum value that the layer-4 destination port number in the packet
    // must have in order to match.  This value must be equal to or less than the
    // value specified for this entry in citapStreamSourceL4PortMax.   If both
    // citapStreamSourceL4PortMin and citapStreamSourceL4PortMax are at their
    // default values, the port number is effectively unused. The type is
    // interface{} with range: 0..65535.
    Citapstreamsourcel4Portmin interface{}

    // The maximum value that the layer-4 destination port number in the packet
    // must have in order to match this classifier entry. This value must be equal
    // to or greater than the value specified for this entry in
    // citapStreamSourceL4PortMin.   If both citapStreamSourceL4PortMin and
    // citapStreamSourceL4PortMax are at their default values, the port number is
    // effectively unused. The type is interface{} with range: 0..65535.
    Citapstreamsourcel4Portmax interface{}

    // An ASCII string, which is the name of a Virtual Routing and Forwarding
    // (VRF) table comprising the routing context of a Virtual Private Network.
    // The interface or set of  interfaces on which the packet might be found
    // should be  selected from the set of interfaces in the VRF table.  A string
    // length of zero implies that global routing table be used for selection of
    // interfaces on which the packet might be found. The type is string.
    Citapstreamvrf interface{}

    // The status of this conceptual row. This object manages creation,
    // modification, and deletion of rows in this table. When any rows must be
    // changed, citapStreamStatus must be first  set to 'notInService'. The type
    // is RowStatus.
    Citapstreamstatus interface{}
}

func (citapstreamentry *CISCOIPTAPMIB_Citapstreamtable_Citapstreamentry) GetEntityData() *types.CommonEntityData {
    citapstreamentry.EntityData.YFilter = citapstreamentry.YFilter
    citapstreamentry.EntityData.YangName = "citapStreamEntry"
    citapstreamentry.EntityData.BundleName = "cisco_ios_xe"
    citapstreamentry.EntityData.ParentYangName = "citapStreamTable"
    citapstreamentry.EntityData.SegmentPath = "citapStreamEntry" + "[cTap2MediationContentId='" + fmt.Sprintf("%v", citapstreamentry.Ctap2Mediationcontentid) + "']" + "[cTap2StreamIndex='" + fmt.Sprintf("%v", citapstreamentry.Ctap2Streamindex) + "']"
    citapstreamentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    citapstreamentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    citapstreamentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    citapstreamentry.EntityData.Children = make(map[string]types.YChild)
    citapstreamentry.EntityData.Leafs = make(map[string]types.YLeaf)
    citapstreamentry.EntityData.Leafs["cTap2MediationContentId"] = types.YLeaf{"Ctap2Mediationcontentid", citapstreamentry.Ctap2Mediationcontentid}
    citapstreamentry.EntityData.Leafs["cTap2StreamIndex"] = types.YLeaf{"Ctap2Streamindex", citapstreamentry.Ctap2Streamindex}
    citapstreamentry.EntityData.Leafs["citapStreamInterface"] = types.YLeaf{"Citapstreaminterface", citapstreamentry.Citapstreaminterface}
    citapstreamentry.EntityData.Leafs["citapStreamAddrType"] = types.YLeaf{"Citapstreamaddrtype", citapstreamentry.Citapstreamaddrtype}
    citapstreamentry.EntityData.Leafs["citapStreamDestinationAddress"] = types.YLeaf{"Citapstreamdestinationaddress", citapstreamentry.Citapstreamdestinationaddress}
    citapstreamentry.EntityData.Leafs["citapStreamDestinationLength"] = types.YLeaf{"Citapstreamdestinationlength", citapstreamentry.Citapstreamdestinationlength}
    citapstreamentry.EntityData.Leafs["citapStreamSourceAddress"] = types.YLeaf{"Citapstreamsourceaddress", citapstreamentry.Citapstreamsourceaddress}
    citapstreamentry.EntityData.Leafs["citapStreamSourceLength"] = types.YLeaf{"Citapstreamsourcelength", citapstreamentry.Citapstreamsourcelength}
    citapstreamentry.EntityData.Leafs["citapStreamTosByte"] = types.YLeaf{"Citapstreamtosbyte", citapstreamentry.Citapstreamtosbyte}
    citapstreamentry.EntityData.Leafs["citapStreamTosByteMask"] = types.YLeaf{"Citapstreamtosbytemask", citapstreamentry.Citapstreamtosbytemask}
    citapstreamentry.EntityData.Leafs["citapStreamFlowId"] = types.YLeaf{"Citapstreamflowid", citapstreamentry.Citapstreamflowid}
    citapstreamentry.EntityData.Leafs["citapStreamProtocol"] = types.YLeaf{"Citapstreamprotocol", citapstreamentry.Citapstreamprotocol}
    citapstreamentry.EntityData.Leafs["citapStreamDestL4PortMin"] = types.YLeaf{"Citapstreamdestl4Portmin", citapstreamentry.Citapstreamdestl4Portmin}
    citapstreamentry.EntityData.Leafs["citapStreamDestL4PortMax"] = types.YLeaf{"Citapstreamdestl4Portmax", citapstreamentry.Citapstreamdestl4Portmax}
    citapstreamentry.EntityData.Leafs["citapStreamSourceL4PortMin"] = types.YLeaf{"Citapstreamsourcel4Portmin", citapstreamentry.Citapstreamsourcel4Portmin}
    citapstreamentry.EntityData.Leafs["citapStreamSourceL4PortMax"] = types.YLeaf{"Citapstreamsourcel4Portmax", citapstreamentry.Citapstreamsourcel4Portmax}
    citapstreamentry.EntityData.Leafs["citapStreamVRF"] = types.YLeaf{"Citapstreamvrf", citapstreamentry.Citapstreamvrf}
    citapstreamentry.EntityData.Leafs["citapStreamStatus"] = types.YLeaf{"Citapstreamstatus", citapstreamentry.Citapstreamstatus}
    return &(citapstreamentry.EntityData)
}

