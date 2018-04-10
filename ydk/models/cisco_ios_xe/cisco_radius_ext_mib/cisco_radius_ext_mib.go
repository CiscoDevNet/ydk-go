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

    
    Creclientglobal CISCORADIUSEXTMIB_Creclientglobal

    
    Creclientauthentication CISCORADIUSEXTMIB_Creclientauthentication

    
    Creclientaccounting CISCORADIUSEXTMIB_Creclientaccounting
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

    cISCORADIUSEXTMIB.EntityData.Children = make(map[string]types.YChild)
    cISCORADIUSEXTMIB.EntityData.Children["creClientGlobal"] = types.YChild{"Creclientglobal", &cISCORADIUSEXTMIB.Creclientglobal}
    cISCORADIUSEXTMIB.EntityData.Children["creClientAuthentication"] = types.YChild{"Creclientauthentication", &cISCORADIUSEXTMIB.Creclientauthentication}
    cISCORADIUSEXTMIB.EntityData.Children["creClientAccounting"] = types.YChild{"Creclientaccounting", &cISCORADIUSEXTMIB.Creclientaccounting}
    cISCORADIUSEXTMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCORADIUSEXTMIB.EntityData)
}

// CISCORADIUSEXTMIB_Creclientglobal
type CISCORADIUSEXTMIB_Creclientglobal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object indicates the maximum length of the queue which stores the
    // incoming RADIUS packets. The type is interface{} with range: 0..4294967295.
    // Units are RADIUS packets.
    Creclienttotalmaxinqlength interface{}

    // This object indicates the maximum length of the queue which stores the
    // pending RADIUS packets for which the responses are outstanding. The type is
    // interface{} with range: 0..4294967295. Units are RADIUS packets.
    Creclienttotalmaxwaitqlength interface{}

    // This object indicates the maximum length of the queue which stores those
    // RADIUS packets for which the responses are received. The type is
    // interface{} with range: 0..4294967295. Units are RADIUS packets.
    Creclienttotalmaxdoneqlength interface{}

    // This object indicates the number of access reject packets received by the
    // RADIUS client. The type is interface{} with range: 0..4294967295. Units are
    // RADIUS packets.
    Creclienttotalaccessrejects interface{}

    // This object indicates the overall response delay experienced by RADIUS
    // packets (both authentication and accounting). The type is interface{} with
    // range: 0..2147483647.
    Creclienttotalaverageresponsedelay interface{}

    // If the 'extended RADIUS source ports' is configured, multiple source ports
    // are used for sending out RADIUS authentication or accounting requests. 
    // This MIB object indicates the port value from where this range starts. The
    // type is interface{} with range: 0..65535.
    Creclientsourceportrangestart interface{}

    // If the 'extended RADIUS source port' is configured, multiple source ports
    // are used for sending out RADIUS authentication or accounting requests. 
    // This MIB object indicates the port value where this range ends. The type is
    // interface{} with range: 0..65535.
    Creclientsourceportrangeend interface{}

    // If the 'extended RADIUS source ports' is configured, multiple source ports
    // are used for sending out RADIUS authentication or accounting requests. 
    // This MIB object indicates the last source port that was used to send out a
    // RADIUS authentication or accounting request. The type is interface{} with
    // range: 0..65535.
    Creclientlastusedsourceport interface{}

    // This MIB object indicates the last source identifier that was used to send
    // out a RADIUS packet when 'extended RADIUS source ports' is configured.  The
    // source identifier is a counter that is incremented everytime a RADIUS
    // authentication or an accounting packet is sent. The type is interface{}
    // with range: 0..255.
    Creclientlastusedsourceid interface{}
}

func (creclientglobal *CISCORADIUSEXTMIB_Creclientglobal) GetEntityData() *types.CommonEntityData {
    creclientglobal.EntityData.YFilter = creclientglobal.YFilter
    creclientglobal.EntityData.YangName = "creClientGlobal"
    creclientglobal.EntityData.BundleName = "cisco_ios_xe"
    creclientglobal.EntityData.ParentYangName = "CISCO-RADIUS-EXT-MIB"
    creclientglobal.EntityData.SegmentPath = "creClientGlobal"
    creclientglobal.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    creclientglobal.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    creclientglobal.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    creclientglobal.EntityData.Children = make(map[string]types.YChild)
    creclientglobal.EntityData.Leafs = make(map[string]types.YLeaf)
    creclientglobal.EntityData.Leafs["creClientTotalMaxInQLength"] = types.YLeaf{"Creclienttotalmaxinqlength", creclientglobal.Creclienttotalmaxinqlength}
    creclientglobal.EntityData.Leafs["creClientTotalMaxWaitQLength"] = types.YLeaf{"Creclienttotalmaxwaitqlength", creclientglobal.Creclienttotalmaxwaitqlength}
    creclientglobal.EntityData.Leafs["creClientTotalMaxDoneQLength"] = types.YLeaf{"Creclienttotalmaxdoneqlength", creclientglobal.Creclienttotalmaxdoneqlength}
    creclientglobal.EntityData.Leafs["creClientTotalAccessRejects"] = types.YLeaf{"Creclienttotalaccessrejects", creclientglobal.Creclienttotalaccessrejects}
    creclientglobal.EntityData.Leafs["creClientTotalAverageResponseDelay"] = types.YLeaf{"Creclienttotalaverageresponsedelay", creclientglobal.Creclienttotalaverageresponsedelay}
    creclientglobal.EntityData.Leafs["creClientSourcePortRangeStart"] = types.YLeaf{"Creclientsourceportrangestart", creclientglobal.Creclientsourceportrangestart}
    creclientglobal.EntityData.Leafs["creClientSourcePortRangeEnd"] = types.YLeaf{"Creclientsourceportrangeend", creclientglobal.Creclientsourceportrangeend}
    creclientglobal.EntityData.Leafs["creClientLastUsedSourcePort"] = types.YLeaf{"Creclientlastusedsourceport", creclientglobal.Creclientlastusedsourceport}
    creclientglobal.EntityData.Leafs["creClientLastUsedSourceId"] = types.YLeaf{"Creclientlastusedsourceid", creclientglobal.Creclientlastusedsourceid}
    return &(creclientglobal.EntityData)
}

// CISCORADIUSEXTMIB_Creclientauthentication
type CISCORADIUSEXTMIB_Creclientauthentication struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object indicates the number of RADIUS authentication response packets
    // received which contained invalid authenticators. The type is interface{}
    // with range: 0..4294967295. Units are RADIUS packets.
    Creauthclientbadauthenticators interface{}

    // This object indicates the number of unknown RADIUS authentication responses
    // received. The type is interface{} with range: 0..4294967295. Units are
    // RADIUS packets.
    Creauthclientunknownresponses interface{}

    // This object indicates the number of RADIUS authentication packets that
    // received responses. The type is interface{} with range: 0..4294967295.
    // Units are RADIUS packets.
    Creauthclienttotalpacketswithresponses interface{}

    // This object indicates the number of buffer allocation failures encountered
    // during RADIUS request formation. The type is interface{} with range:
    // 0..4294967295. Units are buffer failures.
    Creauthclientbufferallocfailures interface{}

    // This object indicates the number of RADIUS authentication response packets
    // received by the RADIUS client. The type is interface{} with range:
    // 0..4294967295. Units are RADIUS packets.
    Creauthclienttotalresponses interface{}

    // This object indicates the number of RADIUS authentication packets that
    // never received a response. The type is interface{} with range:
    // 0..4294967295. Units are RADIUS packets.
    Creauthclienttotalpacketswithoutresponses interface{}

    // This object indicates the average response delay experienced for RADIUS
    // authentication requests. The type is interface{} with range: 0..2147483647.
    Creauthclientaverageresponsedelay interface{}

    // This object indicates the maximum delay experienced for RADIUS
    // authentication requests. The type is interface{} with range: 0..2147483647.
    Creauthclientmaxresponsedelay interface{}

    // This object indicates the maximum buffer size for RADIUS authentication
    // packet. The type is interface{} with range: 0..4294967295. Units are bytes.
    Creauthclientmaxbuffersize interface{}

    // This object indicates the number of timeouts that have occurred for RADIUS
    // authentication.  After a timeout the client may retry sending the request
    // to the same server or to a different server or give up depending on the
    // configuration. The type is interface{} with range: 0..4294967295. Units are
    // timeouts.
    Creauthclienttimeouts interface{}

    // This object indicates the number of times client has received duplicate
    // authentication responses with the same identifier.  Out of these two
    // packets, the later packet is considered as a true match. The type is
    // interface{} with range: 0..4294967295. Units are RADIUS packets.
    Creauthclientdupids interface{}

    // This object indicates the number of malformed RADIUS authentication
    // responses received.  Malformed packets include packets with an invalid
    // length. The type is interface{} with range: 0..4294967295. Units are RADIUS
    // packets.
    Creauthclientmalformedresponses interface{}

    // This MIB object indicates the last source identifier that was used to send
    // out a RADIUS authentication request when 'extended RADIUS source ports' is
    // configured.  The source identifier is a counter that is incremented
    // everytime a RADIUS authentication request is sent. The type is interface{}
    // with range: 0..255.
    Creauthclientlastusedsourceid interface{}
}

func (creclientauthentication *CISCORADIUSEXTMIB_Creclientauthentication) GetEntityData() *types.CommonEntityData {
    creclientauthentication.EntityData.YFilter = creclientauthentication.YFilter
    creclientauthentication.EntityData.YangName = "creClientAuthentication"
    creclientauthentication.EntityData.BundleName = "cisco_ios_xe"
    creclientauthentication.EntityData.ParentYangName = "CISCO-RADIUS-EXT-MIB"
    creclientauthentication.EntityData.SegmentPath = "creClientAuthentication"
    creclientauthentication.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    creclientauthentication.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    creclientauthentication.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    creclientauthentication.EntityData.Children = make(map[string]types.YChild)
    creclientauthentication.EntityData.Leafs = make(map[string]types.YLeaf)
    creclientauthentication.EntityData.Leafs["creAuthClientBadAuthenticators"] = types.YLeaf{"Creauthclientbadauthenticators", creclientauthentication.Creauthclientbadauthenticators}
    creclientauthentication.EntityData.Leafs["creAuthClientUnknownResponses"] = types.YLeaf{"Creauthclientunknownresponses", creclientauthentication.Creauthclientunknownresponses}
    creclientauthentication.EntityData.Leafs["creAuthClientTotalPacketsWithResponses"] = types.YLeaf{"Creauthclienttotalpacketswithresponses", creclientauthentication.Creauthclienttotalpacketswithresponses}
    creclientauthentication.EntityData.Leafs["creAuthClientBufferAllocFailures"] = types.YLeaf{"Creauthclientbufferallocfailures", creclientauthentication.Creauthclientbufferallocfailures}
    creclientauthentication.EntityData.Leafs["creAuthClientTotalResponses"] = types.YLeaf{"Creauthclienttotalresponses", creclientauthentication.Creauthclienttotalresponses}
    creclientauthentication.EntityData.Leafs["creAuthClientTotalPacketsWithoutResponses"] = types.YLeaf{"Creauthclienttotalpacketswithoutresponses", creclientauthentication.Creauthclienttotalpacketswithoutresponses}
    creclientauthentication.EntityData.Leafs["creAuthClientAverageResponseDelay"] = types.YLeaf{"Creauthclientaverageresponsedelay", creclientauthentication.Creauthclientaverageresponsedelay}
    creclientauthentication.EntityData.Leafs["creAuthClientMaxResponseDelay"] = types.YLeaf{"Creauthclientmaxresponsedelay", creclientauthentication.Creauthclientmaxresponsedelay}
    creclientauthentication.EntityData.Leafs["creAuthClientMaxBufferSize"] = types.YLeaf{"Creauthclientmaxbuffersize", creclientauthentication.Creauthclientmaxbuffersize}
    creclientauthentication.EntityData.Leafs["creAuthClientTimeouts"] = types.YLeaf{"Creauthclienttimeouts", creclientauthentication.Creauthclienttimeouts}
    creclientauthentication.EntityData.Leafs["creAuthClientDupIDs"] = types.YLeaf{"Creauthclientdupids", creclientauthentication.Creauthclientdupids}
    creclientauthentication.EntityData.Leafs["creAuthClientMalformedResponses"] = types.YLeaf{"Creauthclientmalformedresponses", creclientauthentication.Creauthclientmalformedresponses}
    creclientauthentication.EntityData.Leafs["creAuthClientLastUsedSourceId"] = types.YLeaf{"Creauthclientlastusedsourceid", creclientauthentication.Creauthclientlastusedsourceid}
    return &(creclientauthentication.EntityData)
}

// CISCORADIUSEXTMIB_Creclientaccounting
type CISCORADIUSEXTMIB_Creclientaccounting struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object indicates the number of RADIUS Accounting-Response packets
    // received with invalid authenticators. The type is interface{} with range:
    // 0..4294967295. Units are RADIUS packets.
    Creacctclientbadauthenticators interface{}

    // This object indicates the number of unknown RADIUS accounting responses
    // received. The type is interface{} with range: 0..4294967295. Units are
    // RADIUS packets.
    Creacctclientunknownresponses interface{}

    // This object indicates the number of RADIUS accounting packets that received
    // responses. The type is interface{} with range: 0..4294967295. Units are
    // RADIUS packets.
    Creacctclienttotalpacketswithresponses interface{}

    // This object indicates the number of buffer allocation failures encountered
    // for RADIUS accounting request. The type is interface{} with range:
    // 0..4294967295. Units are buffer failures.
    Creacctclientbufferallocfailures interface{}

    // This object indicates the number of RADIUS accounting response packets
    // received by the RADIUS client. The type is interface{} with range:
    // 0..4294967295. Units are RADIUS packets.
    Creacctclienttotalresponses interface{}

    // This object indicates the number of RADIUS accounting packets that never
    // received a response. The type is interface{} with range: 0..4294967295.
    // Units are RADIUS packets.
    Creacctclienttotalpacketswithoutresponses interface{}

    // This object indicates the average response delay experienced for RADIUS
    // accounting. The type is interface{} with range: 0..2147483647.
    Creacctclientaverageresponsedelay interface{}

    // This object indicates the maximum delay experienced for RADIUS accounting.
    // The type is interface{} with range: 0..2147483647.
    Creacctclientmaxresponsedelay interface{}

    // This object indicates the maximum buffer size for RADIUS accounting
    // packets. The type is interface{} with range: 0..4294967295. Units are
    // bytes.
    Creacctclientmaxbuffersize interface{}

    // This object indicates the number of timeouts that have occurred for RADIUS
    // accounting.  After a timeout the client may retry sending the request to
    // the same server or to a different server or give up depending on the
    // configuration. The type is interface{} with range: 0..4294967295. Units are
    // timeouts.
    Creacctclienttimeouts interface{}

    // This object indicates the number of times client has received duplicate
    // accounting responses with the same identifier.  Out of these two packets,
    // the later packet is considered as a true match. The type is interface{}
    // with range: 0..4294967295. Units are RADIUS packets.
    Creacctclientdupids interface{}

    // This object indicates the number of malformed RADIUS accounting responses
    // received.  Malformed packets include packets with an invalid length. The
    // type is interface{} with range: 0..4294967295. Units are RADIUS packets.
    Creacctclientmalformedresponses interface{}

    // This MIB object indicates the last source identifier that was used to send
    // out a RADIUS accounting request when 'extended RADIUS source ports' is
    // configured.  The source identifier is a counter that is incremented
    // everytime a RADIUS accounting request is sent. The type is interface{} with
    // range: 0..255.
    Creacctclientlastusedsourceid interface{}
}

func (creclientaccounting *CISCORADIUSEXTMIB_Creclientaccounting) GetEntityData() *types.CommonEntityData {
    creclientaccounting.EntityData.YFilter = creclientaccounting.YFilter
    creclientaccounting.EntityData.YangName = "creClientAccounting"
    creclientaccounting.EntityData.BundleName = "cisco_ios_xe"
    creclientaccounting.EntityData.ParentYangName = "CISCO-RADIUS-EXT-MIB"
    creclientaccounting.EntityData.SegmentPath = "creClientAccounting"
    creclientaccounting.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    creclientaccounting.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    creclientaccounting.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    creclientaccounting.EntityData.Children = make(map[string]types.YChild)
    creclientaccounting.EntityData.Leafs = make(map[string]types.YLeaf)
    creclientaccounting.EntityData.Leafs["creAcctClientBadAuthenticators"] = types.YLeaf{"Creacctclientbadauthenticators", creclientaccounting.Creacctclientbadauthenticators}
    creclientaccounting.EntityData.Leafs["creAcctClientUnknownResponses"] = types.YLeaf{"Creacctclientunknownresponses", creclientaccounting.Creacctclientunknownresponses}
    creclientaccounting.EntityData.Leafs["creAcctClientTotalPacketsWithResponses"] = types.YLeaf{"Creacctclienttotalpacketswithresponses", creclientaccounting.Creacctclienttotalpacketswithresponses}
    creclientaccounting.EntityData.Leafs["creAcctClientBufferAllocFailures"] = types.YLeaf{"Creacctclientbufferallocfailures", creclientaccounting.Creacctclientbufferallocfailures}
    creclientaccounting.EntityData.Leafs["creAcctClientTotalResponses"] = types.YLeaf{"Creacctclienttotalresponses", creclientaccounting.Creacctclienttotalresponses}
    creclientaccounting.EntityData.Leafs["creAcctClientTotalPacketsWithoutResponses"] = types.YLeaf{"Creacctclienttotalpacketswithoutresponses", creclientaccounting.Creacctclienttotalpacketswithoutresponses}
    creclientaccounting.EntityData.Leafs["creAcctClientAverageResponseDelay"] = types.YLeaf{"Creacctclientaverageresponsedelay", creclientaccounting.Creacctclientaverageresponsedelay}
    creclientaccounting.EntityData.Leafs["creAcctClientMaxResponseDelay"] = types.YLeaf{"Creacctclientmaxresponsedelay", creclientaccounting.Creacctclientmaxresponsedelay}
    creclientaccounting.EntityData.Leafs["creAcctClientMaxBufferSize"] = types.YLeaf{"Creacctclientmaxbuffersize", creclientaccounting.Creacctclientmaxbuffersize}
    creclientaccounting.EntityData.Leafs["creAcctClientTimeouts"] = types.YLeaf{"Creacctclienttimeouts", creclientaccounting.Creacctclienttimeouts}
    creclientaccounting.EntityData.Leafs["creAcctClientDupIDs"] = types.YLeaf{"Creacctclientdupids", creclientaccounting.Creacctclientdupids}
    creclientaccounting.EntityData.Leafs["creAcctClientMalformedResponses"] = types.YLeaf{"Creacctclientmalformedresponses", creclientaccounting.Creacctclientmalformedresponses}
    creclientaccounting.EntityData.Leafs["creAcctClientLastUsedSourceId"] = types.YLeaf{"Creacctclientlastusedsourceid", creclientaccounting.Creacctclientlastusedsourceid}
    return &(creclientaccounting.EntityData)
}

