// This MIB module defines objects describing RADIUS (Remote
// Access Dialin User Service), serving as an extension of the
// following MIB modules: 
// - 
//     - RADIUS-AUTH-CLIENT-MIB [RFC4668] 
//     - RADIUS-AUTH-SERVER-MIB [RFC4669] 
//     - RADIUS-ACC-CLIENT-MIB [RFC4670] 
//     - RADIUS-ACC-SERVER-MIB [RFC4671] 
//     - RADIUS-DYNAUTH-CLIENT-MIB [RFC4672] 
//     - RADIUS-DYNAUTH-SERVER-MIB [RFC4673] 
// - 
// [RFC4668] D. Nelson, RADIUS Authentication Client MIB for IPv6,
// 
// RFC-4668, August 2006. 
// - 
// [RFC4669] D. Nelson, RADIUS Authentication Server MIB for IPv6,
// 
// RFC-4669, August 2006. 
// - 
// [RFC4670] D. Nelson, RADIUS Accounting Client MIB for IPv6,
// RFC-4670, August 2006. 
// - 
// [RFC4671] D. Nelson, RADIUS Accounting Server MIB for IPv6,
// RFC-4671, August 2006. 
// - 
// [RFC4672] S. De Cnodder, N. Jonnala, M. Chiba, RADIUS Dynamic 
// Authorization Client MIB, RFC-4672, September 2006. 
// - 
// [RFC4673] S. De Cnodder, N. Jonnala, M. Chiba, RADIUS Dynamic 
// Authorization Server MIB, RFC-4673, September 2006.
package cisco_radius_ext_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_radius_ext_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-RADIUS-EXT-MIB CISCO-RADIUS-EXT-MIB}", reflect.TypeOf(CISCORADIUSEXTMIB{}))
    ydk.RegisterEntity("CISCO-RADIUS-EXT-MIB:CISCO-RADIUS-EXT-MIB", reflect.TypeOf(CISCORADIUSEXTMIB{}))
}

// CISCORADIUSEXTMIB
type CISCORADIUSEXTMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    CreClientGlobal CISCORADIUSEXTMIB_CreClientGlobal

    
    CreClientAuthentication CISCORADIUSEXTMIB_CreClientAuthentication

    
    CreClientAccounting CISCORADIUSEXTMIB_CreClientAccounting
}

func (cISCORADIUSEXTMIB *CISCORADIUSEXTMIB) GetEntityData() *types.CommonEntityData {
    cISCORADIUSEXTMIB.EntityData.YFilter = cISCORADIUSEXTMIB.YFilter
    cISCORADIUSEXTMIB.EntityData.YangName = "CISCO-RADIUS-EXT-MIB"
    cISCORADIUSEXTMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCORADIUSEXTMIB.EntityData.ParentYangName = "CISCO-RADIUS-EXT-MIB"
    cISCORADIUSEXTMIB.EntityData.SegmentPath = "CISCO-RADIUS-EXT-MIB:CISCO-RADIUS-EXT-MIB"
    cISCORADIUSEXTMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCORADIUSEXTMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCORADIUSEXTMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCORADIUSEXTMIB.EntityData.Children = types.NewOrderedMap()
    cISCORADIUSEXTMIB.EntityData.Children.Append("creClientGlobal", types.YChild{"CreClientGlobal", &cISCORADIUSEXTMIB.CreClientGlobal})
    cISCORADIUSEXTMIB.EntityData.Children.Append("creClientAuthentication", types.YChild{"CreClientAuthentication", &cISCORADIUSEXTMIB.CreClientAuthentication})
    cISCORADIUSEXTMIB.EntityData.Children.Append("creClientAccounting", types.YChild{"CreClientAccounting", &cISCORADIUSEXTMIB.CreClientAccounting})
    cISCORADIUSEXTMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCORADIUSEXTMIB.EntityData.YListKeys = []string {}

    return &(cISCORADIUSEXTMIB.EntityData)
}

// CISCORADIUSEXTMIB_CreClientGlobal
type CISCORADIUSEXTMIB_CreClientGlobal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object indicates the maximum length of the queue which stores the
    // incoming RADIUS packets. The type is interface{} with range: 0..4294967295.
    // Units are RADIUS packets.
    CreClientTotalMaxInQLength interface{}

    // This object indicates the maximum length of the queue which stores the
    // pending RADIUS packets for which the responses are outstanding. The type is
    // interface{} with range: 0..4294967295. Units are RADIUS packets.
    CreClientTotalMaxWaitQLength interface{}

    // This object indicates the maximum length of the queue which stores those
    // RADIUS packets for which the responses are received. The type is
    // interface{} with range: 0..4294967295. Units are RADIUS packets.
    CreClientTotalMaxDoneQLength interface{}

    // This object indicates the number of access reject packets received by the
    // RADIUS client. The type is interface{} with range: 0..4294967295. Units are
    // RADIUS packets.
    CreClientTotalAccessRejects interface{}

    // This object indicates the overall response delay experienced by RADIUS
    // packets (both authentication and accounting). The type is interface{} with
    // range: 0..2147483647.
    CreClientTotalAverageResponseDelay interface{}

    // If the 'extended RADIUS source ports' is configured, multiple source ports
    // are used for sending out RADIUS authentication or accounting requests. 
    // This MIB object indicates the port value from where this range starts. The
    // type is interface{} with range: 0..65535.
    CreClientSourcePortRangeStart interface{}

    // If the 'extended RADIUS source port' is configured, multiple source ports
    // are used for sending out RADIUS authentication or accounting requests. 
    // This MIB object indicates the port value where this range ends. The type is
    // interface{} with range: 0..65535.
    CreClientSourcePortRangeEnd interface{}

    // If the 'extended RADIUS source ports' is configured, multiple source ports
    // are used for sending out RADIUS authentication or accounting requests. 
    // This MIB object indicates the last source port that was used to send out a
    // RADIUS authentication or accounting request. The type is interface{} with
    // range: 0..65535.
    CreClientLastUsedSourcePort interface{}

    // This MIB object indicates the last source identifier that was used to send
    // out a RADIUS packet when 'extended RADIUS source ports' is configured.  The
    // source identifier is a counter that is incremented everytime a RADIUS
    // authentication or an accounting packet is sent. The type is interface{}
    // with range: 0..255.
    CreClientLastUsedSourceId interface{}
}

func (creClientGlobal *CISCORADIUSEXTMIB_CreClientGlobal) GetEntityData() *types.CommonEntityData {
    creClientGlobal.EntityData.YFilter = creClientGlobal.YFilter
    creClientGlobal.EntityData.YangName = "creClientGlobal"
    creClientGlobal.EntityData.BundleName = "cisco_ios_xe"
    creClientGlobal.EntityData.ParentYangName = "CISCO-RADIUS-EXT-MIB"
    creClientGlobal.EntityData.SegmentPath = "creClientGlobal"
    creClientGlobal.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    creClientGlobal.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    creClientGlobal.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    creClientGlobal.EntityData.Children = types.NewOrderedMap()
    creClientGlobal.EntityData.Leafs = types.NewOrderedMap()
    creClientGlobal.EntityData.Leafs.Append("creClientTotalMaxInQLength", types.YLeaf{"CreClientTotalMaxInQLength", creClientGlobal.CreClientTotalMaxInQLength})
    creClientGlobal.EntityData.Leafs.Append("creClientTotalMaxWaitQLength", types.YLeaf{"CreClientTotalMaxWaitQLength", creClientGlobal.CreClientTotalMaxWaitQLength})
    creClientGlobal.EntityData.Leafs.Append("creClientTotalMaxDoneQLength", types.YLeaf{"CreClientTotalMaxDoneQLength", creClientGlobal.CreClientTotalMaxDoneQLength})
    creClientGlobal.EntityData.Leafs.Append("creClientTotalAccessRejects", types.YLeaf{"CreClientTotalAccessRejects", creClientGlobal.CreClientTotalAccessRejects})
    creClientGlobal.EntityData.Leafs.Append("creClientTotalAverageResponseDelay", types.YLeaf{"CreClientTotalAverageResponseDelay", creClientGlobal.CreClientTotalAverageResponseDelay})
    creClientGlobal.EntityData.Leafs.Append("creClientSourcePortRangeStart", types.YLeaf{"CreClientSourcePortRangeStart", creClientGlobal.CreClientSourcePortRangeStart})
    creClientGlobal.EntityData.Leafs.Append("creClientSourcePortRangeEnd", types.YLeaf{"CreClientSourcePortRangeEnd", creClientGlobal.CreClientSourcePortRangeEnd})
    creClientGlobal.EntityData.Leafs.Append("creClientLastUsedSourcePort", types.YLeaf{"CreClientLastUsedSourcePort", creClientGlobal.CreClientLastUsedSourcePort})
    creClientGlobal.EntityData.Leafs.Append("creClientLastUsedSourceId", types.YLeaf{"CreClientLastUsedSourceId", creClientGlobal.CreClientLastUsedSourceId})

    creClientGlobal.EntityData.YListKeys = []string {}

    return &(creClientGlobal.EntityData)
}

// CISCORADIUSEXTMIB_CreClientAuthentication
type CISCORADIUSEXTMIB_CreClientAuthentication struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object indicates the number of RADIUS authentication response packets
    // received which contained invalid authenticators. The type is interface{}
    // with range: 0..4294967295. Units are RADIUS packets.
    CreAuthClientBadAuthenticators interface{}

    // This object indicates the number of unknown RADIUS authentication responses
    // received. The type is interface{} with range: 0..4294967295. Units are
    // RADIUS packets.
    CreAuthClientUnknownResponses interface{}

    // This object indicates the number of RADIUS authentication packets that
    // received responses. The type is interface{} with range: 0..4294967295.
    // Units are RADIUS packets.
    CreAuthClientTotalPacketsWithResponses interface{}

    // This object indicates the number of buffer allocation failures encountered
    // during RADIUS request formation. The type is interface{} with range:
    // 0..4294967295. Units are buffer failures.
    CreAuthClientBufferAllocFailures interface{}

    // This object indicates the number of RADIUS authentication response packets
    // received by the RADIUS client. The type is interface{} with range:
    // 0..4294967295. Units are RADIUS packets.
    CreAuthClientTotalResponses interface{}

    // This object indicates the number of RADIUS authentication packets that
    // never received a response. The type is interface{} with range:
    // 0..4294967295. Units are RADIUS packets.
    CreAuthClientTotalPacketsWithoutResponses interface{}

    // This object indicates the average response delay experienced for RADIUS
    // authentication requests. The type is interface{} with range: 0..2147483647.
    CreAuthClientAverageResponseDelay interface{}

    // This object indicates the maximum delay experienced for RADIUS
    // authentication requests. The type is interface{} with range: 0..2147483647.
    CreAuthClientMaxResponseDelay interface{}

    // This object indicates the maximum buffer size for RADIUS authentication
    // packet. The type is interface{} with range: 0..4294967295. Units are bytes.
    CreAuthClientMaxBufferSize interface{}

    // This object indicates the number of timeouts that have occurred for RADIUS
    // authentication.  After a timeout the client may retry sending the request
    // to the same server or to a different server or give up depending on the
    // configuration. The type is interface{} with range: 0..4294967295. Units are
    // timeouts.
    CreAuthClientTimeouts interface{}

    // This object indicates the number of times client has received duplicate
    // authentication responses with the same identifier.  Out of these two
    // packets, the later packet is considered as a true match. The type is
    // interface{} with range: 0..4294967295. Units are RADIUS packets.
    CreAuthClientDupIDs interface{}

    // This object indicates the number of malformed RADIUS authentication
    // responses received.  Malformed packets include packets with an invalid
    // length. The type is interface{} with range: 0..4294967295. Units are RADIUS
    // packets.
    CreAuthClientMalformedResponses interface{}

    // This MIB object indicates the last source identifier that was used to send
    // out a RADIUS authentication request when 'extended RADIUS source ports' is
    // configured.  The source identifier is a counter that is incremented
    // everytime a RADIUS authentication request is sent. The type is interface{}
    // with range: 0..255.
    CreAuthClientLastUsedSourceId interface{}
}

func (creClientAuthentication *CISCORADIUSEXTMIB_CreClientAuthentication) GetEntityData() *types.CommonEntityData {
    creClientAuthentication.EntityData.YFilter = creClientAuthentication.YFilter
    creClientAuthentication.EntityData.YangName = "creClientAuthentication"
    creClientAuthentication.EntityData.BundleName = "cisco_ios_xe"
    creClientAuthentication.EntityData.ParentYangName = "CISCO-RADIUS-EXT-MIB"
    creClientAuthentication.EntityData.SegmentPath = "creClientAuthentication"
    creClientAuthentication.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    creClientAuthentication.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    creClientAuthentication.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    creClientAuthentication.EntityData.Children = types.NewOrderedMap()
    creClientAuthentication.EntityData.Leafs = types.NewOrderedMap()
    creClientAuthentication.EntityData.Leafs.Append("creAuthClientBadAuthenticators", types.YLeaf{"CreAuthClientBadAuthenticators", creClientAuthentication.CreAuthClientBadAuthenticators})
    creClientAuthentication.EntityData.Leafs.Append("creAuthClientUnknownResponses", types.YLeaf{"CreAuthClientUnknownResponses", creClientAuthentication.CreAuthClientUnknownResponses})
    creClientAuthentication.EntityData.Leafs.Append("creAuthClientTotalPacketsWithResponses", types.YLeaf{"CreAuthClientTotalPacketsWithResponses", creClientAuthentication.CreAuthClientTotalPacketsWithResponses})
    creClientAuthentication.EntityData.Leafs.Append("creAuthClientBufferAllocFailures", types.YLeaf{"CreAuthClientBufferAllocFailures", creClientAuthentication.CreAuthClientBufferAllocFailures})
    creClientAuthentication.EntityData.Leafs.Append("creAuthClientTotalResponses", types.YLeaf{"CreAuthClientTotalResponses", creClientAuthentication.CreAuthClientTotalResponses})
    creClientAuthentication.EntityData.Leafs.Append("creAuthClientTotalPacketsWithoutResponses", types.YLeaf{"CreAuthClientTotalPacketsWithoutResponses", creClientAuthentication.CreAuthClientTotalPacketsWithoutResponses})
    creClientAuthentication.EntityData.Leafs.Append("creAuthClientAverageResponseDelay", types.YLeaf{"CreAuthClientAverageResponseDelay", creClientAuthentication.CreAuthClientAverageResponseDelay})
    creClientAuthentication.EntityData.Leafs.Append("creAuthClientMaxResponseDelay", types.YLeaf{"CreAuthClientMaxResponseDelay", creClientAuthentication.CreAuthClientMaxResponseDelay})
    creClientAuthentication.EntityData.Leafs.Append("creAuthClientMaxBufferSize", types.YLeaf{"CreAuthClientMaxBufferSize", creClientAuthentication.CreAuthClientMaxBufferSize})
    creClientAuthentication.EntityData.Leafs.Append("creAuthClientTimeouts", types.YLeaf{"CreAuthClientTimeouts", creClientAuthentication.CreAuthClientTimeouts})
    creClientAuthentication.EntityData.Leafs.Append("creAuthClientDupIDs", types.YLeaf{"CreAuthClientDupIDs", creClientAuthentication.CreAuthClientDupIDs})
    creClientAuthentication.EntityData.Leafs.Append("creAuthClientMalformedResponses", types.YLeaf{"CreAuthClientMalformedResponses", creClientAuthentication.CreAuthClientMalformedResponses})
    creClientAuthentication.EntityData.Leafs.Append("creAuthClientLastUsedSourceId", types.YLeaf{"CreAuthClientLastUsedSourceId", creClientAuthentication.CreAuthClientLastUsedSourceId})

    creClientAuthentication.EntityData.YListKeys = []string {}

    return &(creClientAuthentication.EntityData)
}

// CISCORADIUSEXTMIB_CreClientAccounting
type CISCORADIUSEXTMIB_CreClientAccounting struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object indicates the number of RADIUS Accounting-Response packets
    // received with invalid authenticators. The type is interface{} with range:
    // 0..4294967295. Units are RADIUS packets.
    CreAcctClientBadAuthenticators interface{}

    // This object indicates the number of unknown RADIUS accounting responses
    // received. The type is interface{} with range: 0..4294967295. Units are
    // RADIUS packets.
    CreAcctClientUnknownResponses interface{}

    // This object indicates the number of RADIUS accounting packets that received
    // responses. The type is interface{} with range: 0..4294967295. Units are
    // RADIUS packets.
    CreAcctClientTotalPacketsWithResponses interface{}

    // This object indicates the number of buffer allocation failures encountered
    // for RADIUS accounting request. The type is interface{} with range:
    // 0..4294967295. Units are buffer failures.
    CreAcctClientBufferAllocFailures interface{}

    // This object indicates the number of RADIUS accounting response packets
    // received by the RADIUS client. The type is interface{} with range:
    // 0..4294967295. Units are RADIUS packets.
    CreAcctClientTotalResponses interface{}

    // This object indicates the number of RADIUS accounting packets that never
    // received a response. The type is interface{} with range: 0..4294967295.
    // Units are RADIUS packets.
    CreAcctClientTotalPacketsWithoutResponses interface{}

    // This object indicates the average response delay experienced for RADIUS
    // accounting. The type is interface{} with range: 0..2147483647.
    CreAcctClientAverageResponseDelay interface{}

    // This object indicates the maximum delay experienced for RADIUS accounting.
    // The type is interface{} with range: 0..2147483647.
    CreAcctClientMaxResponseDelay interface{}

    // This object indicates the maximum buffer size for RADIUS accounting
    // packets. The type is interface{} with range: 0..4294967295. Units are
    // bytes.
    CreAcctClientMaxBufferSize interface{}

    // This object indicates the number of timeouts that have occurred for RADIUS
    // accounting.  After a timeout the client may retry sending the request to
    // the same server or to a different server or give up depending on the
    // configuration. The type is interface{} with range: 0..4294967295. Units are
    // timeouts.
    CreAcctClientTimeouts interface{}

    // This object indicates the number of times client has received duplicate
    // accounting responses with the same identifier.  Out of these two packets,
    // the later packet is considered as a true match. The type is interface{}
    // with range: 0..4294967295. Units are RADIUS packets.
    CreAcctClientDupIDs interface{}

    // This object indicates the number of malformed RADIUS accounting responses
    // received.  Malformed packets include packets with an invalid length. The
    // type is interface{} with range: 0..4294967295. Units are RADIUS packets.
    CreAcctClientMalformedResponses interface{}

    // This MIB object indicates the last source identifier that was used to send
    // out a RADIUS accounting request when 'extended RADIUS source ports' is
    // configured.  The source identifier is a counter that is incremented
    // everytime a RADIUS accounting request is sent. The type is interface{} with
    // range: 0..255.
    CreAcctClientLastUsedSourceId interface{}
}

func (creClientAccounting *CISCORADIUSEXTMIB_CreClientAccounting) GetEntityData() *types.CommonEntityData {
    creClientAccounting.EntityData.YFilter = creClientAccounting.YFilter
    creClientAccounting.EntityData.YangName = "creClientAccounting"
    creClientAccounting.EntityData.BundleName = "cisco_ios_xe"
    creClientAccounting.EntityData.ParentYangName = "CISCO-RADIUS-EXT-MIB"
    creClientAccounting.EntityData.SegmentPath = "creClientAccounting"
    creClientAccounting.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    creClientAccounting.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    creClientAccounting.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    creClientAccounting.EntityData.Children = types.NewOrderedMap()
    creClientAccounting.EntityData.Leafs = types.NewOrderedMap()
    creClientAccounting.EntityData.Leafs.Append("creAcctClientBadAuthenticators", types.YLeaf{"CreAcctClientBadAuthenticators", creClientAccounting.CreAcctClientBadAuthenticators})
    creClientAccounting.EntityData.Leafs.Append("creAcctClientUnknownResponses", types.YLeaf{"CreAcctClientUnknownResponses", creClientAccounting.CreAcctClientUnknownResponses})
    creClientAccounting.EntityData.Leafs.Append("creAcctClientTotalPacketsWithResponses", types.YLeaf{"CreAcctClientTotalPacketsWithResponses", creClientAccounting.CreAcctClientTotalPacketsWithResponses})
    creClientAccounting.EntityData.Leafs.Append("creAcctClientBufferAllocFailures", types.YLeaf{"CreAcctClientBufferAllocFailures", creClientAccounting.CreAcctClientBufferAllocFailures})
    creClientAccounting.EntityData.Leafs.Append("creAcctClientTotalResponses", types.YLeaf{"CreAcctClientTotalResponses", creClientAccounting.CreAcctClientTotalResponses})
    creClientAccounting.EntityData.Leafs.Append("creAcctClientTotalPacketsWithoutResponses", types.YLeaf{"CreAcctClientTotalPacketsWithoutResponses", creClientAccounting.CreAcctClientTotalPacketsWithoutResponses})
    creClientAccounting.EntityData.Leafs.Append("creAcctClientAverageResponseDelay", types.YLeaf{"CreAcctClientAverageResponseDelay", creClientAccounting.CreAcctClientAverageResponseDelay})
    creClientAccounting.EntityData.Leafs.Append("creAcctClientMaxResponseDelay", types.YLeaf{"CreAcctClientMaxResponseDelay", creClientAccounting.CreAcctClientMaxResponseDelay})
    creClientAccounting.EntityData.Leafs.Append("creAcctClientMaxBufferSize", types.YLeaf{"CreAcctClientMaxBufferSize", creClientAccounting.CreAcctClientMaxBufferSize})
    creClientAccounting.EntityData.Leafs.Append("creAcctClientTimeouts", types.YLeaf{"CreAcctClientTimeouts", creClientAccounting.CreAcctClientTimeouts})
    creClientAccounting.EntityData.Leafs.Append("creAcctClientDupIDs", types.YLeaf{"CreAcctClientDupIDs", creClientAccounting.CreAcctClientDupIDs})
    creClientAccounting.EntityData.Leafs.Append("creAcctClientMalformedResponses", types.YLeaf{"CreAcctClientMalformedResponses", creClientAccounting.CreAcctClientMalformedResponses})
    creClientAccounting.EntityData.Leafs.Append("creAcctClientLastUsedSourceId", types.YLeaf{"CreAcctClientLastUsedSourceId", creClientAccounting.CreAcctClientLastUsedSourceId})

    creClientAccounting.EntityData.YListKeys = []string {}

    return &(creClientAccounting.EntityData)
}

