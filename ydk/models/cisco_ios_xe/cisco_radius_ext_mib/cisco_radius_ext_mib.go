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
    parent types.Entity
    YFilter yfilter.YFilter

    
    Creclientglobal CISCORADIUSEXTMIB_Creclientglobal

    
    Creclientauthentication CISCORADIUSEXTMIB_Creclientauthentication

    
    Creclientaccounting CISCORADIUSEXTMIB_Creclientaccounting
}

func (cISCORADIUSEXTMIB *CISCORADIUSEXTMIB) GetFilter() yfilter.YFilter { return cISCORADIUSEXTMIB.YFilter }

func (cISCORADIUSEXTMIB *CISCORADIUSEXTMIB) SetFilter(yf yfilter.YFilter) { cISCORADIUSEXTMIB.YFilter = yf }

func (cISCORADIUSEXTMIB *CISCORADIUSEXTMIB) GetGoName(yname string) string {
    if yname == "creClientGlobal" { return "Creclientglobal" }
    if yname == "creClientAuthentication" { return "Creclientauthentication" }
    if yname == "creClientAccounting" { return "Creclientaccounting" }
    return ""
}

func (cISCORADIUSEXTMIB *CISCORADIUSEXTMIB) GetSegmentPath() string {
    return "CISCO-RADIUS-EXT-MIB:CISCO-RADIUS-EXT-MIB"
}

func (cISCORADIUSEXTMIB *CISCORADIUSEXTMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "creClientGlobal" {
        return &cISCORADIUSEXTMIB.Creclientglobal
    }
    if childYangName == "creClientAuthentication" {
        return &cISCORADIUSEXTMIB.Creclientauthentication
    }
    if childYangName == "creClientAccounting" {
        return &cISCORADIUSEXTMIB.Creclientaccounting
    }
    return nil
}

func (cISCORADIUSEXTMIB *CISCORADIUSEXTMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["creClientGlobal"] = &cISCORADIUSEXTMIB.Creclientglobal
    children["creClientAuthentication"] = &cISCORADIUSEXTMIB.Creclientauthentication
    children["creClientAccounting"] = &cISCORADIUSEXTMIB.Creclientaccounting
    return children
}

func (cISCORADIUSEXTMIB *CISCORADIUSEXTMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCORADIUSEXTMIB *CISCORADIUSEXTMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCORADIUSEXTMIB *CISCORADIUSEXTMIB) GetYangName() string { return "CISCO-RADIUS-EXT-MIB" }

func (cISCORADIUSEXTMIB *CISCORADIUSEXTMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCORADIUSEXTMIB *CISCORADIUSEXTMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCORADIUSEXTMIB *CISCORADIUSEXTMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCORADIUSEXTMIB *CISCORADIUSEXTMIB) SetParent(parent types.Entity) { cISCORADIUSEXTMIB.parent = parent }

func (cISCORADIUSEXTMIB *CISCORADIUSEXTMIB) GetParent() types.Entity { return cISCORADIUSEXTMIB.parent }

func (cISCORADIUSEXTMIB *CISCORADIUSEXTMIB) GetParentYangName() string { return "CISCO-RADIUS-EXT-MIB" }

// CISCORADIUSEXTMIB_Creclientglobal
type CISCORADIUSEXTMIB_Creclientglobal struct {
    parent types.Entity
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

func (creclientglobal *CISCORADIUSEXTMIB_Creclientglobal) GetFilter() yfilter.YFilter { return creclientglobal.YFilter }

func (creclientglobal *CISCORADIUSEXTMIB_Creclientglobal) SetFilter(yf yfilter.YFilter) { creclientglobal.YFilter = yf }

func (creclientglobal *CISCORADIUSEXTMIB_Creclientglobal) GetGoName(yname string) string {
    if yname == "creClientTotalMaxInQLength" { return "Creclienttotalmaxinqlength" }
    if yname == "creClientTotalMaxWaitQLength" { return "Creclienttotalmaxwaitqlength" }
    if yname == "creClientTotalMaxDoneQLength" { return "Creclienttotalmaxdoneqlength" }
    if yname == "creClientTotalAccessRejects" { return "Creclienttotalaccessrejects" }
    if yname == "creClientTotalAverageResponseDelay" { return "Creclienttotalaverageresponsedelay" }
    if yname == "creClientSourcePortRangeStart" { return "Creclientsourceportrangestart" }
    if yname == "creClientSourcePortRangeEnd" { return "Creclientsourceportrangeend" }
    if yname == "creClientLastUsedSourcePort" { return "Creclientlastusedsourceport" }
    if yname == "creClientLastUsedSourceId" { return "Creclientlastusedsourceid" }
    return ""
}

func (creclientglobal *CISCORADIUSEXTMIB_Creclientglobal) GetSegmentPath() string {
    return "creClientGlobal"
}

func (creclientglobal *CISCORADIUSEXTMIB_Creclientglobal) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (creclientglobal *CISCORADIUSEXTMIB_Creclientglobal) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (creclientglobal *CISCORADIUSEXTMIB_Creclientglobal) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["creClientTotalMaxInQLength"] = creclientglobal.Creclienttotalmaxinqlength
    leafs["creClientTotalMaxWaitQLength"] = creclientglobal.Creclienttotalmaxwaitqlength
    leafs["creClientTotalMaxDoneQLength"] = creclientglobal.Creclienttotalmaxdoneqlength
    leafs["creClientTotalAccessRejects"] = creclientglobal.Creclienttotalaccessrejects
    leafs["creClientTotalAverageResponseDelay"] = creclientglobal.Creclienttotalaverageresponsedelay
    leafs["creClientSourcePortRangeStart"] = creclientglobal.Creclientsourceportrangestart
    leafs["creClientSourcePortRangeEnd"] = creclientglobal.Creclientsourceportrangeend
    leafs["creClientLastUsedSourcePort"] = creclientglobal.Creclientlastusedsourceport
    leafs["creClientLastUsedSourceId"] = creclientglobal.Creclientlastusedsourceid
    return leafs
}

func (creclientglobal *CISCORADIUSEXTMIB_Creclientglobal) GetBundleName() string { return "cisco_ios_xe" }

func (creclientglobal *CISCORADIUSEXTMIB_Creclientglobal) GetYangName() string { return "creClientGlobal" }

func (creclientglobal *CISCORADIUSEXTMIB_Creclientglobal) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (creclientglobal *CISCORADIUSEXTMIB_Creclientglobal) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (creclientglobal *CISCORADIUSEXTMIB_Creclientglobal) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (creclientglobal *CISCORADIUSEXTMIB_Creclientglobal) SetParent(parent types.Entity) { creclientglobal.parent = parent }

func (creclientglobal *CISCORADIUSEXTMIB_Creclientglobal) GetParent() types.Entity { return creclientglobal.parent }

func (creclientglobal *CISCORADIUSEXTMIB_Creclientglobal) GetParentYangName() string { return "CISCO-RADIUS-EXT-MIB" }

// CISCORADIUSEXTMIB_Creclientauthentication
type CISCORADIUSEXTMIB_Creclientauthentication struct {
    parent types.Entity
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

func (creclientauthentication *CISCORADIUSEXTMIB_Creclientauthentication) GetFilter() yfilter.YFilter { return creclientauthentication.YFilter }

func (creclientauthentication *CISCORADIUSEXTMIB_Creclientauthentication) SetFilter(yf yfilter.YFilter) { creclientauthentication.YFilter = yf }

func (creclientauthentication *CISCORADIUSEXTMIB_Creclientauthentication) GetGoName(yname string) string {
    if yname == "creAuthClientBadAuthenticators" { return "Creauthclientbadauthenticators" }
    if yname == "creAuthClientUnknownResponses" { return "Creauthclientunknownresponses" }
    if yname == "creAuthClientTotalPacketsWithResponses" { return "Creauthclienttotalpacketswithresponses" }
    if yname == "creAuthClientBufferAllocFailures" { return "Creauthclientbufferallocfailures" }
    if yname == "creAuthClientTotalResponses" { return "Creauthclienttotalresponses" }
    if yname == "creAuthClientTotalPacketsWithoutResponses" { return "Creauthclienttotalpacketswithoutresponses" }
    if yname == "creAuthClientAverageResponseDelay" { return "Creauthclientaverageresponsedelay" }
    if yname == "creAuthClientMaxResponseDelay" { return "Creauthclientmaxresponsedelay" }
    if yname == "creAuthClientMaxBufferSize" { return "Creauthclientmaxbuffersize" }
    if yname == "creAuthClientTimeouts" { return "Creauthclienttimeouts" }
    if yname == "creAuthClientDupIDs" { return "Creauthclientdupids" }
    if yname == "creAuthClientMalformedResponses" { return "Creauthclientmalformedresponses" }
    if yname == "creAuthClientLastUsedSourceId" { return "Creauthclientlastusedsourceid" }
    return ""
}

func (creclientauthentication *CISCORADIUSEXTMIB_Creclientauthentication) GetSegmentPath() string {
    return "creClientAuthentication"
}

func (creclientauthentication *CISCORADIUSEXTMIB_Creclientauthentication) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (creclientauthentication *CISCORADIUSEXTMIB_Creclientauthentication) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (creclientauthentication *CISCORADIUSEXTMIB_Creclientauthentication) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["creAuthClientBadAuthenticators"] = creclientauthentication.Creauthclientbadauthenticators
    leafs["creAuthClientUnknownResponses"] = creclientauthentication.Creauthclientunknownresponses
    leafs["creAuthClientTotalPacketsWithResponses"] = creclientauthentication.Creauthclienttotalpacketswithresponses
    leafs["creAuthClientBufferAllocFailures"] = creclientauthentication.Creauthclientbufferallocfailures
    leafs["creAuthClientTotalResponses"] = creclientauthentication.Creauthclienttotalresponses
    leafs["creAuthClientTotalPacketsWithoutResponses"] = creclientauthentication.Creauthclienttotalpacketswithoutresponses
    leafs["creAuthClientAverageResponseDelay"] = creclientauthentication.Creauthclientaverageresponsedelay
    leafs["creAuthClientMaxResponseDelay"] = creclientauthentication.Creauthclientmaxresponsedelay
    leafs["creAuthClientMaxBufferSize"] = creclientauthentication.Creauthclientmaxbuffersize
    leafs["creAuthClientTimeouts"] = creclientauthentication.Creauthclienttimeouts
    leafs["creAuthClientDupIDs"] = creclientauthentication.Creauthclientdupids
    leafs["creAuthClientMalformedResponses"] = creclientauthentication.Creauthclientmalformedresponses
    leafs["creAuthClientLastUsedSourceId"] = creclientauthentication.Creauthclientlastusedsourceid
    return leafs
}

func (creclientauthentication *CISCORADIUSEXTMIB_Creclientauthentication) GetBundleName() string { return "cisco_ios_xe" }

func (creclientauthentication *CISCORADIUSEXTMIB_Creclientauthentication) GetYangName() string { return "creClientAuthentication" }

func (creclientauthentication *CISCORADIUSEXTMIB_Creclientauthentication) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (creclientauthentication *CISCORADIUSEXTMIB_Creclientauthentication) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (creclientauthentication *CISCORADIUSEXTMIB_Creclientauthentication) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (creclientauthentication *CISCORADIUSEXTMIB_Creclientauthentication) SetParent(parent types.Entity) { creclientauthentication.parent = parent }

func (creclientauthentication *CISCORADIUSEXTMIB_Creclientauthentication) GetParent() types.Entity { return creclientauthentication.parent }

func (creclientauthentication *CISCORADIUSEXTMIB_Creclientauthentication) GetParentYangName() string { return "CISCO-RADIUS-EXT-MIB" }

// CISCORADIUSEXTMIB_Creclientaccounting
type CISCORADIUSEXTMIB_Creclientaccounting struct {
    parent types.Entity
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

func (creclientaccounting *CISCORADIUSEXTMIB_Creclientaccounting) GetFilter() yfilter.YFilter { return creclientaccounting.YFilter }

func (creclientaccounting *CISCORADIUSEXTMIB_Creclientaccounting) SetFilter(yf yfilter.YFilter) { creclientaccounting.YFilter = yf }

func (creclientaccounting *CISCORADIUSEXTMIB_Creclientaccounting) GetGoName(yname string) string {
    if yname == "creAcctClientBadAuthenticators" { return "Creacctclientbadauthenticators" }
    if yname == "creAcctClientUnknownResponses" { return "Creacctclientunknownresponses" }
    if yname == "creAcctClientTotalPacketsWithResponses" { return "Creacctclienttotalpacketswithresponses" }
    if yname == "creAcctClientBufferAllocFailures" { return "Creacctclientbufferallocfailures" }
    if yname == "creAcctClientTotalResponses" { return "Creacctclienttotalresponses" }
    if yname == "creAcctClientTotalPacketsWithoutResponses" { return "Creacctclienttotalpacketswithoutresponses" }
    if yname == "creAcctClientAverageResponseDelay" { return "Creacctclientaverageresponsedelay" }
    if yname == "creAcctClientMaxResponseDelay" { return "Creacctclientmaxresponsedelay" }
    if yname == "creAcctClientMaxBufferSize" { return "Creacctclientmaxbuffersize" }
    if yname == "creAcctClientTimeouts" { return "Creacctclienttimeouts" }
    if yname == "creAcctClientDupIDs" { return "Creacctclientdupids" }
    if yname == "creAcctClientMalformedResponses" { return "Creacctclientmalformedresponses" }
    if yname == "creAcctClientLastUsedSourceId" { return "Creacctclientlastusedsourceid" }
    return ""
}

func (creclientaccounting *CISCORADIUSEXTMIB_Creclientaccounting) GetSegmentPath() string {
    return "creClientAccounting"
}

func (creclientaccounting *CISCORADIUSEXTMIB_Creclientaccounting) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (creclientaccounting *CISCORADIUSEXTMIB_Creclientaccounting) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (creclientaccounting *CISCORADIUSEXTMIB_Creclientaccounting) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["creAcctClientBadAuthenticators"] = creclientaccounting.Creacctclientbadauthenticators
    leafs["creAcctClientUnknownResponses"] = creclientaccounting.Creacctclientunknownresponses
    leafs["creAcctClientTotalPacketsWithResponses"] = creclientaccounting.Creacctclienttotalpacketswithresponses
    leafs["creAcctClientBufferAllocFailures"] = creclientaccounting.Creacctclientbufferallocfailures
    leafs["creAcctClientTotalResponses"] = creclientaccounting.Creacctclienttotalresponses
    leafs["creAcctClientTotalPacketsWithoutResponses"] = creclientaccounting.Creacctclienttotalpacketswithoutresponses
    leafs["creAcctClientAverageResponseDelay"] = creclientaccounting.Creacctclientaverageresponsedelay
    leafs["creAcctClientMaxResponseDelay"] = creclientaccounting.Creacctclientmaxresponsedelay
    leafs["creAcctClientMaxBufferSize"] = creclientaccounting.Creacctclientmaxbuffersize
    leafs["creAcctClientTimeouts"] = creclientaccounting.Creacctclienttimeouts
    leafs["creAcctClientDupIDs"] = creclientaccounting.Creacctclientdupids
    leafs["creAcctClientMalformedResponses"] = creclientaccounting.Creacctclientmalformedresponses
    leafs["creAcctClientLastUsedSourceId"] = creclientaccounting.Creacctclientlastusedsourceid
    return leafs
}

func (creclientaccounting *CISCORADIUSEXTMIB_Creclientaccounting) GetBundleName() string { return "cisco_ios_xe" }

func (creclientaccounting *CISCORADIUSEXTMIB_Creclientaccounting) GetYangName() string { return "creClientAccounting" }

func (creclientaccounting *CISCORADIUSEXTMIB_Creclientaccounting) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (creclientaccounting *CISCORADIUSEXTMIB_Creclientaccounting) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (creclientaccounting *CISCORADIUSEXTMIB_Creclientaccounting) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (creclientaccounting *CISCORADIUSEXTMIB_Creclientaccounting) SetParent(parent types.Entity) { creclientaccounting.parent = parent }

func (creclientaccounting *CISCORADIUSEXTMIB_Creclientaccounting) GetParent() types.Entity { return creclientaccounting.parent }

func (creclientaccounting *CISCORADIUSEXTMIB_Creclientaccounting) GetParentYangName() string { return "CISCO-RADIUS-EXT-MIB" }

