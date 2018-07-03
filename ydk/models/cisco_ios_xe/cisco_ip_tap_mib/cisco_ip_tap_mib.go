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

    
    CitapStreamEncodePacket CISCOIPTAPMIB_CitapStreamEncodePacket

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
    CitapStreamTable CISCOIPTAPMIB_CitapStreamTable
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

    cISCOIPTAPMIB.EntityData.Children = types.NewOrderedMap()
    cISCOIPTAPMIB.EntityData.Children.Append("citapStreamEncodePacket", types.YChild{"CitapStreamEncodePacket", &cISCOIPTAPMIB.CitapStreamEncodePacket})
    cISCOIPTAPMIB.EntityData.Children.Append("citapStreamTable", types.YChild{"CitapStreamTable", &cISCOIPTAPMIB.CitapStreamTable})
    cISCOIPTAPMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOIPTAPMIB.EntityData.YListKeys = []string {}

    return &(cISCOIPTAPMIB.EntityData)
}

// CISCOIPTAPMIB_CitapStreamEncodePacket
type CISCOIPTAPMIB_CitapStreamEncodePacket struct {
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
    CitapStreamCapabilities interface{}
}

func (citapStreamEncodePacket *CISCOIPTAPMIB_CitapStreamEncodePacket) GetEntityData() *types.CommonEntityData {
    citapStreamEncodePacket.EntityData.YFilter = citapStreamEncodePacket.YFilter
    citapStreamEncodePacket.EntityData.YangName = "citapStreamEncodePacket"
    citapStreamEncodePacket.EntityData.BundleName = "cisco_ios_xe"
    citapStreamEncodePacket.EntityData.ParentYangName = "CISCO-IP-TAP-MIB"
    citapStreamEncodePacket.EntityData.SegmentPath = "citapStreamEncodePacket"
    citapStreamEncodePacket.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    citapStreamEncodePacket.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    citapStreamEncodePacket.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    citapStreamEncodePacket.EntityData.Children = types.NewOrderedMap()
    citapStreamEncodePacket.EntityData.Leafs = types.NewOrderedMap()
    citapStreamEncodePacket.EntityData.Leafs.Append("citapStreamCapabilities", types.YLeaf{"CitapStreamCapabilities", citapStreamEncodePacket.CitapStreamCapabilities})

    citapStreamEncodePacket.EntityData.YListKeys = []string {}

    return &(citapStreamEncodePacket.EntityData)
}

// CISCOIPTAPMIB_CitapStreamTable
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
type CISCOIPTAPMIB_CitapStreamTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A stream entry indicates a single data stream to be intercepted to a
    // Mediation Device. Many selected data streams may go to the same application
    // interface, and many application interfaces are supported. The type is slice
    // of CISCOIPTAPMIB_CitapStreamTable_CitapStreamEntry.
    CitapStreamEntry []*CISCOIPTAPMIB_CitapStreamTable_CitapStreamEntry
}

func (citapStreamTable *CISCOIPTAPMIB_CitapStreamTable) GetEntityData() *types.CommonEntityData {
    citapStreamTable.EntityData.YFilter = citapStreamTable.YFilter
    citapStreamTable.EntityData.YangName = "citapStreamTable"
    citapStreamTable.EntityData.BundleName = "cisco_ios_xe"
    citapStreamTable.EntityData.ParentYangName = "CISCO-IP-TAP-MIB"
    citapStreamTable.EntityData.SegmentPath = "citapStreamTable"
    citapStreamTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    citapStreamTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    citapStreamTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    citapStreamTable.EntityData.Children = types.NewOrderedMap()
    citapStreamTable.EntityData.Children.Append("citapStreamEntry", types.YChild{"CitapStreamEntry", nil})
    for i := range citapStreamTable.CitapStreamEntry {
        citapStreamTable.EntityData.Children.Append(types.GetSegmentPath(citapStreamTable.CitapStreamEntry[i]), types.YChild{"CitapStreamEntry", citapStreamTable.CitapStreamEntry[i]})
    }
    citapStreamTable.EntityData.Leafs = types.NewOrderedMap()

    citapStreamTable.EntityData.YListKeys = []string {}

    return &(citapStreamTable.EntityData)
}

// CISCOIPTAPMIB_CitapStreamTable_CitapStreamEntry
// A stream entry indicates a single data stream to be
// intercepted to a Mediation Device. Many selected data
// streams may go to the same application interface, and many
// application interfaces are supported.
type CISCOIPTAPMIB_CitapStreamTable_CitapStreamEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // cisco_tap2_mib.CISCOTAP2MIB_CTap2MediationTable_CTap2MediationEntry_CTap2MediationContentId
    CTap2MediationContentId interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // cisco_tap2_mib.CISCOTAP2MIB_CTap2StreamTable_CTap2StreamEntry_CTap2StreamIndex
    CTap2StreamIndex interface{}

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
    CitapStreamInterface interface{}

    // The type of address, used in packet selection. The type is InetAddressType.
    CitapStreamAddrType interface{}

    // The Destination address or prefix used in packet selection. This address
    // will be of the type specified in citapStreamAddrType. The type is string
    // with length: 0..255.
    CitapStreamDestinationAddress interface{}

    // The length of the Destination Prefix. A value of zero causes all addresses
    // to match.  This prefix length will be consistent with the type specified in
    // citapStreamAddrType. The type is interface{} with range: 0..2040.
    CitapStreamDestinationLength interface{}

    // The Source Address used in packet selection. This address will be of the
    // type specified in citapStreamAddrType. The type is string with length:
    // 0..255.
    CitapStreamSourceAddress interface{}

    // The length of the Source Prefix. A value of zero causes all addresses to
    // match. This prefix length will be consistent with the type specified in
    // citapStreamAddrType. The type is interface{} with range: 0..2040.
    CitapStreamSourceLength interface{}

    // The value of the TOS byte, when masked with citapStreamTosByteMask, of
    // traffic to be intercepted.  If
    // citapStreamTosByte&(~citapStreamTosByteMask)!=0, configuration is rejected.
    // The type is interface{} with range: 0..255.
    CitapStreamTosByte interface{}

    // The value of the TOS byte in an IPv4 or IPv6 header is ANDed with
    // citapStreamTosByteMask and compared with citapStreamTosByte.  If the values
    // are equal, the comparison is equal. If the mask is zero and the TosByte
    // value is zero, the result is to always accept. The type is interface{} with
    // range: 0..255.
    CitapStreamTosByteMask interface{}

    // The flow identifier in an IPv6 header. -1 indicates that the Flow Id is
    // unused. The type is interface{} with range: -1..1048575.
    CitapStreamFlowId interface{}

    // The IP protocol to match against the IPv4 protocol number or the IPv6 Next-
    // Header number in the packet. -1 means 'any IP protocol'. The type is
    // interface{} with range: -1..255.
    CitapStreamProtocol interface{}

    // The minimum value that the layer-4 destination port number in the packet
    // must have in order to match.  This value must be equal to or less than the
    // value specified for this entry in citapStreamDestL4PortMax.   If both
    // citapStreamDestL4PortMin and citapStreamDestL4PortMax are at their default
    // values, the port number is effectively unused. The type is interface{} with
    // range: 0..65535.
    CitapStreamDestL4PortMin interface{}

    // The maximum value that the layer-4 destination port number in the packet
    // must have in order to match this classifier entry. This value must be equal
    // to or greater than the value specified for this entry in
    // citapStreamDestL4PortMin.   If both citapStreamDestL4PortMin and
    // citapStreamDestL4PortMax are at their default values, the port number is
    // effectively unused. The type is interface{} with range: 0..65535.
    CitapStreamDestL4PortMax interface{}

    // The minimum value that the layer-4 destination port number in the packet
    // must have in order to match.  This value must be equal to or less than the
    // value specified for this entry in citapStreamSourceL4PortMax.   If both
    // citapStreamSourceL4PortMin and citapStreamSourceL4PortMax are at their
    // default values, the port number is effectively unused. The type is
    // interface{} with range: 0..65535.
    CitapStreamSourceL4PortMin interface{}

    // The maximum value that the layer-4 destination port number in the packet
    // must have in order to match this classifier entry. This value must be equal
    // to or greater than the value specified for this entry in
    // citapStreamSourceL4PortMin.   If both citapStreamSourceL4PortMin and
    // citapStreamSourceL4PortMax are at their default values, the port number is
    // effectively unused. The type is interface{} with range: 0..65535.
    CitapStreamSourceL4PortMax interface{}

    // An ASCII string, which is the name of a Virtual Routing and Forwarding
    // (VRF) table comprising the routing context of a Virtual Private Network.
    // The interface or set of  interfaces on which the packet might be found
    // should be  selected from the set of interfaces in the VRF table.  A string
    // length of zero implies that global routing table be used for selection of
    // interfaces on which the packet might be found. The type is string.
    CitapStreamVRF interface{}

    // The status of this conceptual row. This object manages creation,
    // modification, and deletion of rows in this table. When any rows must be
    // changed, citapStreamStatus must be first  set to 'notInService'. The type
    // is RowStatus.
    CitapStreamStatus interface{}
}

func (citapStreamEntry *CISCOIPTAPMIB_CitapStreamTable_CitapStreamEntry) GetEntityData() *types.CommonEntityData {
    citapStreamEntry.EntityData.YFilter = citapStreamEntry.YFilter
    citapStreamEntry.EntityData.YangName = "citapStreamEntry"
    citapStreamEntry.EntityData.BundleName = "cisco_ios_xe"
    citapStreamEntry.EntityData.ParentYangName = "citapStreamTable"
    citapStreamEntry.EntityData.SegmentPath = "citapStreamEntry" + types.AddKeyToken(citapStreamEntry.CTap2MediationContentId, "cTap2MediationContentId") + types.AddKeyToken(citapStreamEntry.CTap2StreamIndex, "cTap2StreamIndex")
    citapStreamEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    citapStreamEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    citapStreamEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    citapStreamEntry.EntityData.Children = types.NewOrderedMap()
    citapStreamEntry.EntityData.Leafs = types.NewOrderedMap()
    citapStreamEntry.EntityData.Leafs.Append("cTap2MediationContentId", types.YLeaf{"CTap2MediationContentId", citapStreamEntry.CTap2MediationContentId})
    citapStreamEntry.EntityData.Leafs.Append("cTap2StreamIndex", types.YLeaf{"CTap2StreamIndex", citapStreamEntry.CTap2StreamIndex})
    citapStreamEntry.EntityData.Leafs.Append("citapStreamInterface", types.YLeaf{"CitapStreamInterface", citapStreamEntry.CitapStreamInterface})
    citapStreamEntry.EntityData.Leafs.Append("citapStreamAddrType", types.YLeaf{"CitapStreamAddrType", citapStreamEntry.CitapStreamAddrType})
    citapStreamEntry.EntityData.Leafs.Append("citapStreamDestinationAddress", types.YLeaf{"CitapStreamDestinationAddress", citapStreamEntry.CitapStreamDestinationAddress})
    citapStreamEntry.EntityData.Leafs.Append("citapStreamDestinationLength", types.YLeaf{"CitapStreamDestinationLength", citapStreamEntry.CitapStreamDestinationLength})
    citapStreamEntry.EntityData.Leafs.Append("citapStreamSourceAddress", types.YLeaf{"CitapStreamSourceAddress", citapStreamEntry.CitapStreamSourceAddress})
    citapStreamEntry.EntityData.Leafs.Append("citapStreamSourceLength", types.YLeaf{"CitapStreamSourceLength", citapStreamEntry.CitapStreamSourceLength})
    citapStreamEntry.EntityData.Leafs.Append("citapStreamTosByte", types.YLeaf{"CitapStreamTosByte", citapStreamEntry.CitapStreamTosByte})
    citapStreamEntry.EntityData.Leafs.Append("citapStreamTosByteMask", types.YLeaf{"CitapStreamTosByteMask", citapStreamEntry.CitapStreamTosByteMask})
    citapStreamEntry.EntityData.Leafs.Append("citapStreamFlowId", types.YLeaf{"CitapStreamFlowId", citapStreamEntry.CitapStreamFlowId})
    citapStreamEntry.EntityData.Leafs.Append("citapStreamProtocol", types.YLeaf{"CitapStreamProtocol", citapStreamEntry.CitapStreamProtocol})
    citapStreamEntry.EntityData.Leafs.Append("citapStreamDestL4PortMin", types.YLeaf{"CitapStreamDestL4PortMin", citapStreamEntry.CitapStreamDestL4PortMin})
    citapStreamEntry.EntityData.Leafs.Append("citapStreamDestL4PortMax", types.YLeaf{"CitapStreamDestL4PortMax", citapStreamEntry.CitapStreamDestL4PortMax})
    citapStreamEntry.EntityData.Leafs.Append("citapStreamSourceL4PortMin", types.YLeaf{"CitapStreamSourceL4PortMin", citapStreamEntry.CitapStreamSourceL4PortMin})
    citapStreamEntry.EntityData.Leafs.Append("citapStreamSourceL4PortMax", types.YLeaf{"CitapStreamSourceL4PortMax", citapStreamEntry.CitapStreamSourceL4PortMax})
    citapStreamEntry.EntityData.Leafs.Append("citapStreamVRF", types.YLeaf{"CitapStreamVRF", citapStreamEntry.CitapStreamVRF})
    citapStreamEntry.EntityData.Leafs.Append("citapStreamStatus", types.YLeaf{"CitapStreamStatus", citapStreamEntry.CitapStreamStatus})

    citapStreamEntry.EntityData.YListKeys = []string {"CTap2MediationContentId", "CTap2StreamIndex"}

    return &(citapStreamEntry.EntityData)
}

