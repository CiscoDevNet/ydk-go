// Cisco User Agent Session Initiation Protocol (SIP) 
// MIB module.  SIP is an application-layer signalling
// protocol for creating, modifying and terminating
// multimedia sessions with one or more participants.
// These sessions include Internet multimedia conferences
// and Internet telephone calls.  SIP is defined in 
// RFC 2543 (March 1999).  
// 
// This MIB is defined for the management of SIP User 
// Agents (UAs).  A User Agent is an application which 
// contains both a User Agent Client (UAC) and a User 
// Agent Server (UAS).   A UAC is an application that 
// initiates a SIP request.  A UAS is an application that 
// contacts the user when a SIP request is received and 
// that returns a response on behalf of the user.  The 
// response accepts, rejects, or redirects the request.
// 
// A SIP transaction occurs between a client and a server 
// and comprises all messages from the first request sent
// from the client to the server up to a final (non-1xx) 
// response sent from the server to the client.
package cisco_sip_ua_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_sip_ua_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-SIP-UA-MIB CISCO-SIP-UA-MIB}", reflect.TypeOf(CISCOSIPUAMIB{}))
    ydk.RegisterEntity("CISCO-SIP-UA-MIB:CISCO-SIP-UA-MIB", reflect.TypeOf(CISCOSIPUAMIB{}))
}

type CiscoSipUaMIBNotificationPrefix struct {
}

func (id CiscoSipUaMIBNotificationPrefix) String() string {
	return "CISCO-SIP-UA-MIB:ciscoSipUaMIBNotificationPrefix"
}

type CiscoSipUaMIBNotifications struct {
}

func (id CiscoSipUaMIBNotifications) String() string {
	return "CISCO-SIP-UA-MIB:ciscoSipUaMIBNotifications"
}

// CISCOSIPUAMIB
type CISCOSIPUAMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    CSipCfgBase CISCOSIPUAMIB_CSipCfgBase

    
    CSipCfgTimer CISCOSIPUAMIB_CSipCfgTimer

    
    CSipCfgRetry CISCOSIPUAMIB_CSipCfgRetry

    
    CSipCfgPeer CISCOSIPUAMIB_CSipCfgPeer

    
    CSipCfgAaa CISCOSIPUAMIB_CSipCfgAaa

    
    CSipCfgVoiceServiceVoip CISCOSIPUAMIB_CSipCfgVoiceServiceVoip

    
    CSipStatsInfo CISCOSIPUAMIB_CSipStatsInfo

    
    CSipStatsSuccess CISCOSIPUAMIB_CSipStatsSuccess

    
    CSipStatsRedirect CISCOSIPUAMIB_CSipStatsRedirect

    
    CSipStatsErrClient CISCOSIPUAMIB_CSipStatsErrClient

    
    CSipStatsErrServer CISCOSIPUAMIB_CSipStatsErrServer

    
    CSipStatsGlobalFail CISCOSIPUAMIB_CSipStatsGlobalFail

    
    CSipStatsTraffic CISCOSIPUAMIB_CSipStatsTraffic

    
    CSipStatsRetry CISCOSIPUAMIB_CSipStatsRetry

    
    CSipStatsMisc CISCOSIPUAMIB_CSipStatsMisc

    
    CSipStatsConnection CISCOSIPUAMIB_CSipStatsConnection

    // This table contains configuration for Early Media Cut Through.  The
    // configuration controls how the SIP user agent will process 1xx
    // (Provisional) SIP response messages that contain  Session Definition
    // Protocol (SDP) payloads.
    CSipCfgEarlyMediaTable CISCOSIPUAMIB_CSipCfgEarlyMediaTable

    // This table contains configuration for binding the scope of packets to the
    // particular ethernet interface. The scope for the packets can be specified
    // as either 'signalling' or 'media' packets. The ethernet interface shall be
    // specified by the interface index. The table shall be indexed based on the
    // scope.
    CSipCfgBindSourceAddrTable CISCOSIPUAMIB_CSipCfgBindSourceAddrTable

    // This table contains per dial-peer SIP related  configuration.     The table
    // is a sparse table of dial-peer information. This means, it only reflects
    // dial-peers being used  for SIP.  A dial-peer is being used for SIP if the 
    // value of cvVoIPPeerCfgSessionProtocol  (CISCO-VOICE-DIAL-CONTROL-MIB) is
    // 'sip'.  Dial-peers are not created or destroyed via this table.  Only SIP
    // related configuration can be  performed via this table once the dial-peer
    // exists in the system and is visible in this table.
    CSipCfgPeerTable CISCOSIPUAMIB_CSipCfgPeerTable

    // This table contains SIP status code to PSTN cause code mapping
    // configuration.  Inbound SIP response messages  that will result in outbound
    // PSTN signalling messages will have the SIP status codes transposed into the
    // PSTN cause codes as prescribed by the contents of this  table.
    CSipCfgStatusCauseTable CISCOSIPUAMIB_CSipCfgStatusCauseTable

    // This table contains PSTN cause code to SIP status code mapping
    // configuration.   Inbound PSTN signalling messages that will result in
    // outbound SIP response messages  will have the PSTN cause codes transposed
    // into the SIP status codes as prescribed by the contents of this  table.
    CSipCfgCauseStatusTable CISCOSIPUAMIB_CSipCfgCauseStatusTable

    // This table contains statistics for sent and received 200 Ok response
    // messages.  The  statistics are kept on per SIP method basis.
    CSipStatsSuccessOkTable CISCOSIPUAMIB_CSipStatsSuccessOkTable
}

func (cISCOSIPUAMIB *CISCOSIPUAMIB) GetEntityData() *types.CommonEntityData {
    cISCOSIPUAMIB.EntityData.YFilter = cISCOSIPUAMIB.YFilter
    cISCOSIPUAMIB.EntityData.YangName = "CISCO-SIP-UA-MIB"
    cISCOSIPUAMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOSIPUAMIB.EntityData.ParentYangName = "CISCO-SIP-UA-MIB"
    cISCOSIPUAMIB.EntityData.SegmentPath = "CISCO-SIP-UA-MIB:CISCO-SIP-UA-MIB"
    cISCOSIPUAMIB.EntityData.AbsolutePath = cISCOSIPUAMIB.EntityData.SegmentPath
    cISCOSIPUAMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOSIPUAMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOSIPUAMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOSIPUAMIB.EntityData.Children = types.NewOrderedMap()
    cISCOSIPUAMIB.EntityData.Children.Append("cSipCfgBase", types.YChild{"CSipCfgBase", &cISCOSIPUAMIB.CSipCfgBase})
    cISCOSIPUAMIB.EntityData.Children.Append("cSipCfgTimer", types.YChild{"CSipCfgTimer", &cISCOSIPUAMIB.CSipCfgTimer})
    cISCOSIPUAMIB.EntityData.Children.Append("cSipCfgRetry", types.YChild{"CSipCfgRetry", &cISCOSIPUAMIB.CSipCfgRetry})
    cISCOSIPUAMIB.EntityData.Children.Append("cSipCfgPeer", types.YChild{"CSipCfgPeer", &cISCOSIPUAMIB.CSipCfgPeer})
    cISCOSIPUAMIB.EntityData.Children.Append("cSipCfgAaa", types.YChild{"CSipCfgAaa", &cISCOSIPUAMIB.CSipCfgAaa})
    cISCOSIPUAMIB.EntityData.Children.Append("cSipCfgVoiceServiceVoip", types.YChild{"CSipCfgVoiceServiceVoip", &cISCOSIPUAMIB.CSipCfgVoiceServiceVoip})
    cISCOSIPUAMIB.EntityData.Children.Append("cSipStatsInfo", types.YChild{"CSipStatsInfo", &cISCOSIPUAMIB.CSipStatsInfo})
    cISCOSIPUAMIB.EntityData.Children.Append("cSipStatsSuccess", types.YChild{"CSipStatsSuccess", &cISCOSIPUAMIB.CSipStatsSuccess})
    cISCOSIPUAMIB.EntityData.Children.Append("cSipStatsRedirect", types.YChild{"CSipStatsRedirect", &cISCOSIPUAMIB.CSipStatsRedirect})
    cISCOSIPUAMIB.EntityData.Children.Append("cSipStatsErrClient", types.YChild{"CSipStatsErrClient", &cISCOSIPUAMIB.CSipStatsErrClient})
    cISCOSIPUAMIB.EntityData.Children.Append("cSipStatsErrServer", types.YChild{"CSipStatsErrServer", &cISCOSIPUAMIB.CSipStatsErrServer})
    cISCOSIPUAMIB.EntityData.Children.Append("cSipStatsGlobalFail", types.YChild{"CSipStatsGlobalFail", &cISCOSIPUAMIB.CSipStatsGlobalFail})
    cISCOSIPUAMIB.EntityData.Children.Append("cSipStatsTraffic", types.YChild{"CSipStatsTraffic", &cISCOSIPUAMIB.CSipStatsTraffic})
    cISCOSIPUAMIB.EntityData.Children.Append("cSipStatsRetry", types.YChild{"CSipStatsRetry", &cISCOSIPUAMIB.CSipStatsRetry})
    cISCOSIPUAMIB.EntityData.Children.Append("cSipStatsMisc", types.YChild{"CSipStatsMisc", &cISCOSIPUAMIB.CSipStatsMisc})
    cISCOSIPUAMIB.EntityData.Children.Append("cSipStatsConnection", types.YChild{"CSipStatsConnection", &cISCOSIPUAMIB.CSipStatsConnection})
    cISCOSIPUAMIB.EntityData.Children.Append("cSipCfgEarlyMediaTable", types.YChild{"CSipCfgEarlyMediaTable", &cISCOSIPUAMIB.CSipCfgEarlyMediaTable})
    cISCOSIPUAMIB.EntityData.Children.Append("cSipCfgBindSourceAddrTable", types.YChild{"CSipCfgBindSourceAddrTable", &cISCOSIPUAMIB.CSipCfgBindSourceAddrTable})
    cISCOSIPUAMIB.EntityData.Children.Append("cSipCfgPeerTable", types.YChild{"CSipCfgPeerTable", &cISCOSIPUAMIB.CSipCfgPeerTable})
    cISCOSIPUAMIB.EntityData.Children.Append("cSipCfgStatusCauseTable", types.YChild{"CSipCfgStatusCauseTable", &cISCOSIPUAMIB.CSipCfgStatusCauseTable})
    cISCOSIPUAMIB.EntityData.Children.Append("cSipCfgCauseStatusTable", types.YChild{"CSipCfgCauseStatusTable", &cISCOSIPUAMIB.CSipCfgCauseStatusTable})
    cISCOSIPUAMIB.EntityData.Children.Append("cSipStatsSuccessOkTable", types.YChild{"CSipStatsSuccessOkTable", &cISCOSIPUAMIB.CSipStatsSuccessOkTable})
    cISCOSIPUAMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOSIPUAMIB.EntityData.YListKeys = []string {}

    return &(cISCOSIPUAMIB.EntityData)
}

// CISCOSIPUAMIB_CSipCfgBase
type CISCOSIPUAMIB_CSipCfgBase struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object will reflect the version of SIP supported by this  user agent. 
    // It will follow the same format as SIP version  information contained in the
    // SIP messages generated by this user agent.  For example, user agents
    // supporting SIP version 2 will return 'SIP/2.0' as dictated by RFC 2543. The
    // type is string.
    CSipCfgVersion interface{}

    // This object specifies the transport protocol the SIP user  agent will use
    // to receive SIP messages.  A value of 'disabled' indicates that the UA will
    // not receive any SIP messages. The type is CSipCfgTransport.
    CSipCfgTransport interface{}

    // This object specifies address of the User Location  Server (ULS) being used
    // to resolve the location of end  points.  This could be a Domain Name Server
    // (DNS) or a  SIP proxy/redirect server.  The format of the address follows
    // the IETF service location  protocol. The syntax is as follows:    
    // mapping-type:type-specific-syntax  the mapping-type specifies a scheme for
    // mapping the matching  dial string to a target server. The
    // type-specific-syntax is  exactly that, something that the particular
    // mapping scheme can understand.  For example,    Session target          
    // Meaning    ipv4:171.68.13.55:1006   The session target is the IP           
    // version 4 address of 171.68.13.55                              and port
    // 1006.    dns:pots.cisco.com       The session target is the IP host        
    // with dns name pots.cisco.com.  The valid Mapping type definitions for the
    // peer follow:    ipv4  - Syntax: ipv4:w.x.y.z:port or  ipv4:w.x.y.z     dns 
    // - Syntax: dns:host.domain. The type is string.
    CSipCfgUserLocationServerAddr interface{}

    // This object may be used with any SIP method to limit the  number of proxies
    // that can forward the request to the next  downstream server. The type is
    // interface{} with range: 1..70.
    CSipCfgMaxForwards interface{}

    // This object may specify the interface where the source IP address used in
    // SIP signalling or media packets is configured.  A value of 0 means that 
    // there is no specific source address configured and  in this case the best
    // local IP address will be chosen  by the system. The type is interface{}
    // with range: 0..2147483647.
    CSipCfgBindSrcAddrInterface interface{}

    // This object specifies the scope of packets to which the source IP address
    // of the interface  designated by cSipCfgBindSrcAddrInterface will be bound. 
    // A value of 'all' means the IP address  will be bound to both SIP signalling
    // and media packets. A value of 'control' means the IP address will only be
    // bound to SIP signalling packets.   If cSipCfgBindSrcAddrInterface is set to
    // 0, the value of this object has no meaning. The type is
    // CSipCfgBindSrcAddrScope.
    CSipCfgBindSrcAddrScope interface{}

    // This object specifies the format of the prefix used  by the system for DNS
    // SRV queries.  v1  :  RFC 2052 format - 'protocol.transport.' v2  :  RFC
    // 2782 format - '_protocol._transport.'  This object allows backward
    // compatibility with systems only supporting RFC 2052 format. The type is
    // CSipCfgDnsSrvQueryStringFormat.
    CSipCfgDnsSrvQueryStringFormat interface{}

    // This object specifies how call redirection (3xx) is handled by the user
    // agent.    If 'false', the user agent's treatment of incoming  3xx class
    // response messages is as defined in RFC 2543.   That is, the user agent uses
    // the Contact header(s) from the 3xx response to reinitiate another INVITE
    // transaction to the user's new location.  The call  is redirected.  If
    // 'true', the user agent treats incoming 3xx  response messages as 4xx
    // (client error) class  response messages.  In this case, the call is not
    // redirected, instead it is released with the  appropriate PSTN cause code. 
    // The mapping of SIP 3xx response status codes to 4xx response status codes
    // is as follows:  300 Multiple Choices -> 410 Gone  301 Moved Permanently ->
    // 410 Gone  302 Moved Temporarily -> 480 Temporarily Unavailable  305 User
    // Proxy        -> 410 Gone  380 Alternative Service -> 410 Gone  Any other
    // 3xx -> 410 Gone. The type is bool.
    CSipCfgRedirectionDisabled interface{}

    // This object specifies whether remote media checks for Symmetric Network
    // Address Translation (NAT)  is enabled or disabled.  If 'true', remote media
    // checks are enabled.  The gateway will have the ability to open a Real Time 
    // Transport Protocol (RTP) session with the remote end and then update
    // (modify) the existing RTP session's remote address/port (raddr:rport) with
    // the source address and port of the actual media packet received.  This
    // would be triggered for only those calls where the Session Description
    // Protocol (SDP) received by the gateway has an indication to do so.  If
    // 'false', remote media checks are disabled. The type is bool.
    CSipCfgSymNatEnabled interface{}

    // This object specifies the value of the 'a=direction:<role>' SDP attribute
    // supported by  the user agent.  The direction attribute is used  to describe
    // the role of the user agent (as an  endpoint for a connection-oriented media
    // stream)  in the connection setup procedure.  none    :  No role is
    // specified. passive :  The user agent will advertise itself            as a
    // 'passive' entity (inside the NAT)            in the SDP. active  :  The
    // user agent will advertise itself            as a 'active' entity (outside
    // the NAT)            in the SDP. The type is CSipCfgSymNatDirectionRole.
    CSipCfgSymNatDirectionRole interface{}

    // This object specifies if support for handling  Suspend/Resume events from
    // the switch is enabled or not.  If 'true', the user agent on getting a
    // Suspend will notify the remote agent by sending it a re-invite with a hold
    // SDP. Similarly, when the Gateway receives a Resume, it will initiate a
    // re-invite with the original SDP and unhold the call.  If 'false', the user
    // agent will not initiate any re-invites on receiving Suspend/Resume events,
    // basically it won't be putting the call on hold or off hold. The type is
    // bool.
    CSipCfgSuspendResumeEnabled interface{}

    // This object specifies how the SIP gateway would initiate call hold
    // requests.  directionAttr: The user agent will use the direction            
    // attribute such as a=sendonly or a=inactive in                 the sdp to
    // initiate call hold requests.                    connAddr: The user agent
    // will use 0.0.0.0 connection address            to specify Call Hold. The
    // type is CSipCfgOfferCallHold.
    CSipCfgOfferCallHold interface{}

    // This object specifies that the Reason header overrides SIP  status code
    // mapping table. The type is bool.
    CSipCfgReasonHeaderOveride interface{}

    // This object may be used with any SIP method to limit the  number of proxies
    // that can forward the request to the next  downstream server. The type is
    // interface{} with range: 1..70.
    CSipCfgMaximumForwards interface{}
}

func (cSipCfgBase *CISCOSIPUAMIB_CSipCfgBase) GetEntityData() *types.CommonEntityData {
    cSipCfgBase.EntityData.YFilter = cSipCfgBase.YFilter
    cSipCfgBase.EntityData.YangName = "cSipCfgBase"
    cSipCfgBase.EntityData.BundleName = "cisco_ios_xe"
    cSipCfgBase.EntityData.ParentYangName = "CISCO-SIP-UA-MIB"
    cSipCfgBase.EntityData.SegmentPath = "cSipCfgBase"
    cSipCfgBase.EntityData.AbsolutePath = "CISCO-SIP-UA-MIB:CISCO-SIP-UA-MIB/" + cSipCfgBase.EntityData.SegmentPath
    cSipCfgBase.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cSipCfgBase.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cSipCfgBase.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cSipCfgBase.EntityData.Children = types.NewOrderedMap()
    cSipCfgBase.EntityData.Leafs = types.NewOrderedMap()
    cSipCfgBase.EntityData.Leafs.Append("cSipCfgVersion", types.YLeaf{"CSipCfgVersion", cSipCfgBase.CSipCfgVersion})
    cSipCfgBase.EntityData.Leafs.Append("cSipCfgTransport", types.YLeaf{"CSipCfgTransport", cSipCfgBase.CSipCfgTransport})
    cSipCfgBase.EntityData.Leafs.Append("cSipCfgUserLocationServerAddr", types.YLeaf{"CSipCfgUserLocationServerAddr", cSipCfgBase.CSipCfgUserLocationServerAddr})
    cSipCfgBase.EntityData.Leafs.Append("cSipCfgMaxForwards", types.YLeaf{"CSipCfgMaxForwards", cSipCfgBase.CSipCfgMaxForwards})
    cSipCfgBase.EntityData.Leafs.Append("cSipCfgBindSrcAddrInterface", types.YLeaf{"CSipCfgBindSrcAddrInterface", cSipCfgBase.CSipCfgBindSrcAddrInterface})
    cSipCfgBase.EntityData.Leafs.Append("cSipCfgBindSrcAddrScope", types.YLeaf{"CSipCfgBindSrcAddrScope", cSipCfgBase.CSipCfgBindSrcAddrScope})
    cSipCfgBase.EntityData.Leafs.Append("cSipCfgDnsSrvQueryStringFormat", types.YLeaf{"CSipCfgDnsSrvQueryStringFormat", cSipCfgBase.CSipCfgDnsSrvQueryStringFormat})
    cSipCfgBase.EntityData.Leafs.Append("cSipCfgRedirectionDisabled", types.YLeaf{"CSipCfgRedirectionDisabled", cSipCfgBase.CSipCfgRedirectionDisabled})
    cSipCfgBase.EntityData.Leafs.Append("cSipCfgSymNatEnabled", types.YLeaf{"CSipCfgSymNatEnabled", cSipCfgBase.CSipCfgSymNatEnabled})
    cSipCfgBase.EntityData.Leafs.Append("cSipCfgSymNatDirectionRole", types.YLeaf{"CSipCfgSymNatDirectionRole", cSipCfgBase.CSipCfgSymNatDirectionRole})
    cSipCfgBase.EntityData.Leafs.Append("cSipCfgSuspendResumeEnabled", types.YLeaf{"CSipCfgSuspendResumeEnabled", cSipCfgBase.CSipCfgSuspendResumeEnabled})
    cSipCfgBase.EntityData.Leafs.Append("cSipCfgOfferCallHold", types.YLeaf{"CSipCfgOfferCallHold", cSipCfgBase.CSipCfgOfferCallHold})
    cSipCfgBase.EntityData.Leafs.Append("cSipCfgReasonHeaderOveride", types.YLeaf{"CSipCfgReasonHeaderOveride", cSipCfgBase.CSipCfgReasonHeaderOveride})
    cSipCfgBase.EntityData.Leafs.Append("cSipCfgMaximumForwards", types.YLeaf{"CSipCfgMaximumForwards", cSipCfgBase.CSipCfgMaximumForwards})

    cSipCfgBase.EntityData.YListKeys = []string {}

    return &(cSipCfgBase.EntityData)
}

// CISCOSIPUAMIB_CSipCfgBase_CSipCfgBindSrcAddrScope represents the value of this object has no meaning.
type CISCOSIPUAMIB_CSipCfgBase_CSipCfgBindSrcAddrScope string

const (
    CISCOSIPUAMIB_CSipCfgBase_CSipCfgBindSrcAddrScope_none CISCOSIPUAMIB_CSipCfgBase_CSipCfgBindSrcAddrScope = "none"

    CISCOSIPUAMIB_CSipCfgBase_CSipCfgBindSrcAddrScope_all CISCOSIPUAMIB_CSipCfgBase_CSipCfgBindSrcAddrScope = "all"

    CISCOSIPUAMIB_CSipCfgBase_CSipCfgBindSrcAddrScope_control CISCOSIPUAMIB_CSipCfgBase_CSipCfgBindSrcAddrScope = "control"
)

// CISCOSIPUAMIB_CSipCfgBase_CSipCfgDnsSrvQueryStringFormat represents only supporting RFC 2052 format.
type CISCOSIPUAMIB_CSipCfgBase_CSipCfgDnsSrvQueryStringFormat string

const (
    CISCOSIPUAMIB_CSipCfgBase_CSipCfgDnsSrvQueryStringFormat_v1 CISCOSIPUAMIB_CSipCfgBase_CSipCfgDnsSrvQueryStringFormat = "v1"

    CISCOSIPUAMIB_CSipCfgBase_CSipCfgDnsSrvQueryStringFormat_v2 CISCOSIPUAMIB_CSipCfgBase_CSipCfgDnsSrvQueryStringFormat = "v2"
)

// CISCOSIPUAMIB_CSipCfgBase_CSipCfgOfferCallHold represents            to specify Call Hold.
type CISCOSIPUAMIB_CSipCfgBase_CSipCfgOfferCallHold string

const (
    CISCOSIPUAMIB_CSipCfgBase_CSipCfgOfferCallHold_directionAttr CISCOSIPUAMIB_CSipCfgBase_CSipCfgOfferCallHold = "directionAttr"

    CISCOSIPUAMIB_CSipCfgBase_CSipCfgOfferCallHold_connAddr CISCOSIPUAMIB_CSipCfgBase_CSipCfgOfferCallHold = "connAddr"
)

// CISCOSIPUAMIB_CSipCfgBase_CSipCfgSymNatDirectionRole represents            in the SDP.
type CISCOSIPUAMIB_CSipCfgBase_CSipCfgSymNatDirectionRole string

const (
    CISCOSIPUAMIB_CSipCfgBase_CSipCfgSymNatDirectionRole_none CISCOSIPUAMIB_CSipCfgBase_CSipCfgSymNatDirectionRole = "none"

    CISCOSIPUAMIB_CSipCfgBase_CSipCfgSymNatDirectionRole_passive CISCOSIPUAMIB_CSipCfgBase_CSipCfgSymNatDirectionRole = "passive"

    CISCOSIPUAMIB_CSipCfgBase_CSipCfgSymNatDirectionRole_active CISCOSIPUAMIB_CSipCfgBase_CSipCfgSymNatDirectionRole = "active"
)

// CISCOSIPUAMIB_CSipCfgBase_CSipCfgTransport represents indicates that the UA will not receive any SIP messages.
type CISCOSIPUAMIB_CSipCfgBase_CSipCfgTransport string

const (
    CISCOSIPUAMIB_CSipCfgBase_CSipCfgTransport_udp CISCOSIPUAMIB_CSipCfgBase_CSipCfgTransport = "udp"

    CISCOSIPUAMIB_CSipCfgBase_CSipCfgTransport_tcp CISCOSIPUAMIB_CSipCfgBase_CSipCfgTransport = "tcp"

    CISCOSIPUAMIB_CSipCfgBase_CSipCfgTransport_udpAndTcp CISCOSIPUAMIB_CSipCfgBase_CSipCfgTransport = "udpAndTcp"

    CISCOSIPUAMIB_CSipCfgBase_CSipCfgTransport_disabled CISCOSIPUAMIB_CSipCfgBase_CSipCfgTransport = "disabled"
)

// CISCOSIPUAMIB_CSipCfgTimer
type CISCOSIPUAMIB_CSipCfgTimer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object specifies the time a user agent will wait to  receive a
    // provisional response to a INVITE before resending  the INVITE. The type is
    // interface{} with range: 100..1000. Units are milliseconds.
    CSipCfgTimerTrying interface{}

    // This object specifies the time a user agent will wait to  receive a final
    // response to a INVITE before cancelling the  transaction. The type is
    // interface{} with range: 60000..300000. Units are milliseconds.
    CSipCfgTimerExpires interface{}

    // This object specifies the time a user agent will wait to  receive an ACK
    // confirmation a session is established. The type is interface{} with range:
    // 100..1000. Units are milliseconds.
    CSipCfgTimerConnect interface{}

    // This object specifies the time a user agent will wait to  receive an BYE
    // confirmation a session is disconnected. The type is interface{} with range:
    // 100..1000. Units are milliseconds.
    CSipCfgTimerDisconnect interface{}

    // This object specifies the time a user agent will wait for  a final response
    // before retransmitting the PRACK (PRovisional ACKnowledgment). The type is
    // interface{} with range: 100..1000. Units are milliseconds.
    CSipCfgTimerPrack interface{}

    // This object specifies the time a user agent will wait  for a final response
    // before retransmitting the COMET  (COndition MET). The type is interface{}
    // with range: 100..1000. Units are milliseconds.
    CSipCfgTimerComet interface{}

    // This object specifies the amount of time to wait for a PRACK before
    // retransmitting the reliable 1xx response. The type is interface{} with
    // range: 100..1000. Units are milliseconds.
    CSipCfgTimerReliableRsp interface{}

    // This object specifies the amount of time to wait for a final response
    // before retransmitting the Notify. The type is interface{} with range:
    // 100..1000. Units are milliseconds.
    CSipCfgTimerNotify interface{}

    // This object specifies the amount of time to wait for a final response
    // before retransmitting the Refer. The type is interface{} with range:
    // 100..1000. Units are milliseconds.
    CSipCfgTimerRefer interface{}

    // This object specifies the value of the Min-SE  header for INVITE messages
    // originated by this  user agent and the minimum value of the 
    // Session-Expires headers for INVITE messages  received by this user agent. 
    // Any Session-Expires headers received with a  value below this object's
    // value will be rejected with a 422 client error response message.  Setting
    // this object to a value less than 600 is valid, but the possibly of
    // excessive re-INVITES  and the impact of those messages should be fully 
    // understood and considered an acceptable risk. The type is interface{} with
    // range: 60..86400. Units are seconds.
    CSipCfgTimerSessionTimer interface{}

    // This object specifies the amount of time to wait before  disconnecting a
    // call already on hold. A value of 0 specifies that this functionality is
    // disabled. The type is interface{} with range: 0..None | 15..2880. Units are
    // minutes.
    CSipCfgTimerHold interface{}

    // This object specifies the amount of time to wait for a 200ok response
    // before retransmitting the Info. The type is interface{} with range:
    // 100..1000. Units are milliseconds.
    CSipCfgTimerInfo interface{}

    // This object specifies the amount of time to wait before  aging out a
    // TCP/UDP connection. The type is interface{} with range: 5..30. Units are
    // minutes.
    CSipCfgTimerConnectionAging interface{}

    // This object specifies the amount of time to buffer the INVITE  while
    // waiting for display name info in the Facility.  A value of 0 means that the
    // INVITE wouldn't be buffered waiting for the display name info in the
    // Facility. The type is interface{} with range: 0..None | 50..5000. Units are
    // milliseconds.
    CSipCfgTimerBufferInvite interface{}
}

func (cSipCfgTimer *CISCOSIPUAMIB_CSipCfgTimer) GetEntityData() *types.CommonEntityData {
    cSipCfgTimer.EntityData.YFilter = cSipCfgTimer.YFilter
    cSipCfgTimer.EntityData.YangName = "cSipCfgTimer"
    cSipCfgTimer.EntityData.BundleName = "cisco_ios_xe"
    cSipCfgTimer.EntityData.ParentYangName = "CISCO-SIP-UA-MIB"
    cSipCfgTimer.EntityData.SegmentPath = "cSipCfgTimer"
    cSipCfgTimer.EntityData.AbsolutePath = "CISCO-SIP-UA-MIB:CISCO-SIP-UA-MIB/" + cSipCfgTimer.EntityData.SegmentPath
    cSipCfgTimer.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cSipCfgTimer.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cSipCfgTimer.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cSipCfgTimer.EntityData.Children = types.NewOrderedMap()
    cSipCfgTimer.EntityData.Leafs = types.NewOrderedMap()
    cSipCfgTimer.EntityData.Leafs.Append("cSipCfgTimerTrying", types.YLeaf{"CSipCfgTimerTrying", cSipCfgTimer.CSipCfgTimerTrying})
    cSipCfgTimer.EntityData.Leafs.Append("cSipCfgTimerExpires", types.YLeaf{"CSipCfgTimerExpires", cSipCfgTimer.CSipCfgTimerExpires})
    cSipCfgTimer.EntityData.Leafs.Append("cSipCfgTimerConnect", types.YLeaf{"CSipCfgTimerConnect", cSipCfgTimer.CSipCfgTimerConnect})
    cSipCfgTimer.EntityData.Leafs.Append("cSipCfgTimerDisconnect", types.YLeaf{"CSipCfgTimerDisconnect", cSipCfgTimer.CSipCfgTimerDisconnect})
    cSipCfgTimer.EntityData.Leafs.Append("cSipCfgTimerPrack", types.YLeaf{"CSipCfgTimerPrack", cSipCfgTimer.CSipCfgTimerPrack})
    cSipCfgTimer.EntityData.Leafs.Append("cSipCfgTimerComet", types.YLeaf{"CSipCfgTimerComet", cSipCfgTimer.CSipCfgTimerComet})
    cSipCfgTimer.EntityData.Leafs.Append("cSipCfgTimerReliableRsp", types.YLeaf{"CSipCfgTimerReliableRsp", cSipCfgTimer.CSipCfgTimerReliableRsp})
    cSipCfgTimer.EntityData.Leafs.Append("cSipCfgTimerNotify", types.YLeaf{"CSipCfgTimerNotify", cSipCfgTimer.CSipCfgTimerNotify})
    cSipCfgTimer.EntityData.Leafs.Append("cSipCfgTimerRefer", types.YLeaf{"CSipCfgTimerRefer", cSipCfgTimer.CSipCfgTimerRefer})
    cSipCfgTimer.EntityData.Leafs.Append("cSipCfgTimerSessionTimer", types.YLeaf{"CSipCfgTimerSessionTimer", cSipCfgTimer.CSipCfgTimerSessionTimer})
    cSipCfgTimer.EntityData.Leafs.Append("cSipCfgTimerHold", types.YLeaf{"CSipCfgTimerHold", cSipCfgTimer.CSipCfgTimerHold})
    cSipCfgTimer.EntityData.Leafs.Append("cSipCfgTimerInfo", types.YLeaf{"CSipCfgTimerInfo", cSipCfgTimer.CSipCfgTimerInfo})
    cSipCfgTimer.EntityData.Leafs.Append("cSipCfgTimerConnectionAging", types.YLeaf{"CSipCfgTimerConnectionAging", cSipCfgTimer.CSipCfgTimerConnectionAging})
    cSipCfgTimer.EntityData.Leafs.Append("cSipCfgTimerBufferInvite", types.YLeaf{"CSipCfgTimerBufferInvite", cSipCfgTimer.CSipCfgTimerBufferInvite})

    cSipCfgTimer.EntityData.YListKeys = []string {}

    return &(cSipCfgTimer.EntityData)
}

// CISCOSIPUAMIB_CSipCfgRetry
type CISCOSIPUAMIB_CSipCfgRetry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object specifies the number of times a user agent  will retry sending
    // a INVITE request. The type is interface{} with range: 1..10.
    CSipCfgRetryInvite interface{}

    // This object specifies the number of times a user agent  will retry sending
    // a BYE request. The type is interface{} with range: 1..10.
    CSipCfgRetryBye interface{}

    // This object specifies the number of times a user agent  will retry sending
    // a CANCEL request. The type is interface{} with range: 1..10.
    CSipCfgRetryCancel interface{}

    // This object specifies the number of times a user agent  will retry sending
    // a REGISTER request. The type is interface{} with range: 1..10.
    CSipCfgRetryRegister interface{}

    // This object specifies the number of times a user agent  will retry sending
    // a Response and expecting a ACK. The type is interface{} with range: 1..10.
    CSipCfgRetryResponse interface{}

    // This object specifies the number of times a user agent  will retry sending
    // a PRACK (PRovisional ACKnowledgement). The type is interface{} with range:
    // 1..10.
    CSipCfgRetryPrack interface{}

    // This object specifies the number of times a user agent  will retry sending
    // a COMET (COndition MET). The type is interface{} with range: 1..10.
    CSipCfgRetryComet interface{}

    // This object specifies the number of times a user agent  will retry sending
    // a reliable response. The type is interface{} with range: 1..10.
    CSipCfgRetryReliableRsp interface{}

    // This object specifies the number of times a user agent  will retry sending
    // a Notify request. The type is interface{} with range: 1..10.
    CSipCfgRetryNotify interface{}

    // This object specifies the number of times a user agent  will retry sending
    // a Refer request. The type is interface{} with range: 1..10.
    CSipCfgRetryRefer interface{}

    // This object specifies the number of times a user agent will retry sending a
    // Info request. The type is interface{} with range: 1..10.
    CSipCfgRetryInfo interface{}

    // This object specifies the number of times a user agent will retry sending a
    // Subscribe request. The type is interface{} with range: 1..10.
    CSipCfgRetrySubscribe interface{}
}

func (cSipCfgRetry *CISCOSIPUAMIB_CSipCfgRetry) GetEntityData() *types.CommonEntityData {
    cSipCfgRetry.EntityData.YFilter = cSipCfgRetry.YFilter
    cSipCfgRetry.EntityData.YangName = "cSipCfgRetry"
    cSipCfgRetry.EntityData.BundleName = "cisco_ios_xe"
    cSipCfgRetry.EntityData.ParentYangName = "CISCO-SIP-UA-MIB"
    cSipCfgRetry.EntityData.SegmentPath = "cSipCfgRetry"
    cSipCfgRetry.EntityData.AbsolutePath = "CISCO-SIP-UA-MIB:CISCO-SIP-UA-MIB/" + cSipCfgRetry.EntityData.SegmentPath
    cSipCfgRetry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cSipCfgRetry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cSipCfgRetry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cSipCfgRetry.EntityData.Children = types.NewOrderedMap()
    cSipCfgRetry.EntityData.Leafs = types.NewOrderedMap()
    cSipCfgRetry.EntityData.Leafs.Append("cSipCfgRetryInvite", types.YLeaf{"CSipCfgRetryInvite", cSipCfgRetry.CSipCfgRetryInvite})
    cSipCfgRetry.EntityData.Leafs.Append("cSipCfgRetryBye", types.YLeaf{"CSipCfgRetryBye", cSipCfgRetry.CSipCfgRetryBye})
    cSipCfgRetry.EntityData.Leafs.Append("cSipCfgRetryCancel", types.YLeaf{"CSipCfgRetryCancel", cSipCfgRetry.CSipCfgRetryCancel})
    cSipCfgRetry.EntityData.Leafs.Append("cSipCfgRetryRegister", types.YLeaf{"CSipCfgRetryRegister", cSipCfgRetry.CSipCfgRetryRegister})
    cSipCfgRetry.EntityData.Leafs.Append("cSipCfgRetryResponse", types.YLeaf{"CSipCfgRetryResponse", cSipCfgRetry.CSipCfgRetryResponse})
    cSipCfgRetry.EntityData.Leafs.Append("cSipCfgRetryPrack", types.YLeaf{"CSipCfgRetryPrack", cSipCfgRetry.CSipCfgRetryPrack})
    cSipCfgRetry.EntityData.Leafs.Append("cSipCfgRetryComet", types.YLeaf{"CSipCfgRetryComet", cSipCfgRetry.CSipCfgRetryComet})
    cSipCfgRetry.EntityData.Leafs.Append("cSipCfgRetryReliableRsp", types.YLeaf{"CSipCfgRetryReliableRsp", cSipCfgRetry.CSipCfgRetryReliableRsp})
    cSipCfgRetry.EntityData.Leafs.Append("cSipCfgRetryNotify", types.YLeaf{"CSipCfgRetryNotify", cSipCfgRetry.CSipCfgRetryNotify})
    cSipCfgRetry.EntityData.Leafs.Append("cSipCfgRetryRefer", types.YLeaf{"CSipCfgRetryRefer", cSipCfgRetry.CSipCfgRetryRefer})
    cSipCfgRetry.EntityData.Leafs.Append("cSipCfgRetryInfo", types.YLeaf{"CSipCfgRetryInfo", cSipCfgRetry.CSipCfgRetryInfo})
    cSipCfgRetry.EntityData.Leafs.Append("cSipCfgRetrySubscribe", types.YLeaf{"CSipCfgRetrySubscribe", cSipCfgRetry.CSipCfgRetrySubscribe})

    cSipCfgRetry.EntityData.YListKeys = []string {}

    return &(cSipCfgRetry.EntityData)
}

// CISCOSIPUAMIB_CSipCfgPeer
type CISCOSIPUAMIB_CSipCfgPeer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object specifies the session transport  protocol that will be used for
    // outbound SIP  messages.  This configuration is applicable to all dial-peers
    // in the system having  cSipCfgPeerOutSessionTransport set to 'system'. The
    // type is CSipCfgOutSessionTransport.
    CSipCfgOutSessionTransport interface{}

    // This object specifies the string that will be  placed in either the
    // Supported or Require SIP  header, as specified by cSipCfgReliable1xxRspHdr.
    // The type is string.
    CSipCfgReliable1xxRspStr interface{}

    // This object specifies behavior with respect to Supported or Require headers
    // in SIP messages sent and received via this dial-peer.  If the originating
    // gateway is configured for 'require', the Require header is added to the
    // outgoing INVITEs with the value of cSipCfgReliable1xxStr.  This requires
    // the use of reliable provisional responses by the terminating gateway. 
    // Sessions will be torn down if this use cannot be supported by that gateway.
    // If the originating gateway is configured for 'supported', the Supported
    // header is added to the outgoing INVITEs with the value of
    // cSipCfgReliable1xxStr.  This  requires that an attempt to use reliable
    // provisional responses be made, but sessions can continue without them.  If
    // the originating gateway is configured for 'disabled', the value of
    // cSipCfgReliable1xxStr will NOT be added to either the Require or Supported
    // headers of outgoing INVITEs. The type is CSipCfgReliable1xxRspHdr.
    CSipCfgReliable1xxRspHdr interface{}

    // This object specifies the URL type sent in outbound INVITES generated by
    // this device. The type is CSipCfgUrlType.
    CSipCfgUrlType interface{}
}

func (cSipCfgPeer *CISCOSIPUAMIB_CSipCfgPeer) GetEntityData() *types.CommonEntityData {
    cSipCfgPeer.EntityData.YFilter = cSipCfgPeer.YFilter
    cSipCfgPeer.EntityData.YangName = "cSipCfgPeer"
    cSipCfgPeer.EntityData.BundleName = "cisco_ios_xe"
    cSipCfgPeer.EntityData.ParentYangName = "CISCO-SIP-UA-MIB"
    cSipCfgPeer.EntityData.SegmentPath = "cSipCfgPeer"
    cSipCfgPeer.EntityData.AbsolutePath = "CISCO-SIP-UA-MIB:CISCO-SIP-UA-MIB/" + cSipCfgPeer.EntityData.SegmentPath
    cSipCfgPeer.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cSipCfgPeer.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cSipCfgPeer.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cSipCfgPeer.EntityData.Children = types.NewOrderedMap()
    cSipCfgPeer.EntityData.Leafs = types.NewOrderedMap()
    cSipCfgPeer.EntityData.Leafs.Append("cSipCfgOutSessionTransport", types.YLeaf{"CSipCfgOutSessionTransport", cSipCfgPeer.CSipCfgOutSessionTransport})
    cSipCfgPeer.EntityData.Leafs.Append("cSipCfgReliable1xxRspStr", types.YLeaf{"CSipCfgReliable1xxRspStr", cSipCfgPeer.CSipCfgReliable1xxRspStr})
    cSipCfgPeer.EntityData.Leafs.Append("cSipCfgReliable1xxRspHdr", types.YLeaf{"CSipCfgReliable1xxRspHdr", cSipCfgPeer.CSipCfgReliable1xxRspHdr})
    cSipCfgPeer.EntityData.Leafs.Append("cSipCfgUrlType", types.YLeaf{"CSipCfgUrlType", cSipCfgPeer.CSipCfgUrlType})

    cSipCfgPeer.EntityData.YListKeys = []string {}

    return &(cSipCfgPeer.EntityData)
}

// CISCOSIPUAMIB_CSipCfgPeer_CSipCfgOutSessionTransport represents cSipCfgPeerOutSessionTransport set to 'system'.
type CISCOSIPUAMIB_CSipCfgPeer_CSipCfgOutSessionTransport string

const (
    CISCOSIPUAMIB_CSipCfgPeer_CSipCfgOutSessionTransport_udp CISCOSIPUAMIB_CSipCfgPeer_CSipCfgOutSessionTransport = "udp"

    CISCOSIPUAMIB_CSipCfgPeer_CSipCfgOutSessionTransport_tcp CISCOSIPUAMIB_CSipCfgPeer_CSipCfgOutSessionTransport = "tcp"
)

// CISCOSIPUAMIB_CSipCfgPeer_CSipCfgReliable1xxRspHdr represents either the Require or Supported headers of outgoing INVITEs.
type CISCOSIPUAMIB_CSipCfgPeer_CSipCfgReliable1xxRspHdr string

const (
    CISCOSIPUAMIB_CSipCfgPeer_CSipCfgReliable1xxRspHdr_supported CISCOSIPUAMIB_CSipCfgPeer_CSipCfgReliable1xxRspHdr = "supported"

    CISCOSIPUAMIB_CSipCfgPeer_CSipCfgReliable1xxRspHdr_require CISCOSIPUAMIB_CSipCfgPeer_CSipCfgReliable1xxRspHdr = "require"

    CISCOSIPUAMIB_CSipCfgPeer_CSipCfgReliable1xxRspHdr_disabled CISCOSIPUAMIB_CSipCfgPeer_CSipCfgReliable1xxRspHdr = "disabled"
)

// CISCOSIPUAMIB_CSipCfgPeer_CSipCfgUrlType represents INVITES generated by this device.
type CISCOSIPUAMIB_CSipCfgPeer_CSipCfgUrlType string

const (
    CISCOSIPUAMIB_CSipCfgPeer_CSipCfgUrlType_sip CISCOSIPUAMIB_CSipCfgPeer_CSipCfgUrlType = "sip"

    CISCOSIPUAMIB_CSipCfgPeer_CSipCfgUrlType_tel CISCOSIPUAMIB_CSipCfgPeer_CSipCfgUrlType = "tel"
)

// CISCOSIPUAMIB_CSipCfgAaa
type CISCOSIPUAMIB_CSipCfgAaa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object specifies the source of the information used to populate the
    // username attribute of AAA billing records. The type is CSipCfgAaaUsername.
    CSipCfgAaaUsername interface{}
}

func (cSipCfgAaa *CISCOSIPUAMIB_CSipCfgAaa) GetEntityData() *types.CommonEntityData {
    cSipCfgAaa.EntityData.YFilter = cSipCfgAaa.YFilter
    cSipCfgAaa.EntityData.YangName = "cSipCfgAaa"
    cSipCfgAaa.EntityData.BundleName = "cisco_ios_xe"
    cSipCfgAaa.EntityData.ParentYangName = "CISCO-SIP-UA-MIB"
    cSipCfgAaa.EntityData.SegmentPath = "cSipCfgAaa"
    cSipCfgAaa.EntityData.AbsolutePath = "CISCO-SIP-UA-MIB:CISCO-SIP-UA-MIB/" + cSipCfgAaa.EntityData.SegmentPath
    cSipCfgAaa.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cSipCfgAaa.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cSipCfgAaa.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cSipCfgAaa.EntityData.Children = types.NewOrderedMap()
    cSipCfgAaa.EntityData.Leafs = types.NewOrderedMap()
    cSipCfgAaa.EntityData.Leafs.Append("cSipCfgAaaUsername", types.YLeaf{"CSipCfgAaaUsername", cSipCfgAaa.CSipCfgAaaUsername})

    cSipCfgAaa.EntityData.YListKeys = []string {}

    return &(cSipCfgAaa.EntityData)
}

// CISCOSIPUAMIB_CSipCfgAaa_CSipCfgAaaUsername represents populate the username attribute of AAA billing records.
type CISCOSIPUAMIB_CSipCfgAaa_CSipCfgAaaUsername string

const (
    CISCOSIPUAMIB_CSipCfgAaa_CSipCfgAaaUsername_callingNumber CISCOSIPUAMIB_CSipCfgAaa_CSipCfgAaaUsername = "callingNumber"

    CISCOSIPUAMIB_CSipCfgAaa_CSipCfgAaaUsername_proxyAuth CISCOSIPUAMIB_CSipCfgAaa_CSipCfgAaaUsername = "proxyAuth"
)

// CISCOSIPUAMIB_CSipCfgVoiceServiceVoip
type CISCOSIPUAMIB_CSipCfgVoiceServiceVoip struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object specifies if support for passing SIP headers from Invite,
    // Subscribe, Notify Request to the application is enabled.  If 'true', the
    // headers received in a message will be passed to the application.  If
    // 'false', the headers received in a message will not be passed to the
    // application. The type is bool.
    CSipCfgHeaderPassingEnabled interface{}

    // This object specifies the maximum number of concurrent SIP subscriptions a
    // SIP Gateway can accept. The type is interface{} with range: 0..4294967295.
    CSipCfgMaxSubscriptionAccept interface{}

    // This object specifies the maximum number of concurrent SIP subscriptions
    // that a SIP Gateway can originate. Default is Max Dialpeers on platform.
    // Maximum is 2*Max Dialpeers on Platform. The type is interface{} with range:
    // 0..4294967295.
    CSipCfgMaxSubscriptionOriginate interface{}

    // This object specifies if the functionality of switching between transports
    // from udp to tcp if the message size of a Request is greater than 1300 bytes
    // is enabled or not.  This configuration is at the global level, and will
    // only be  considered if there exists no voip dial-peer. The type is bool.
    CSipCfgSwitchTransportEnabled interface{}
}

func (cSipCfgVoiceServiceVoip *CISCOSIPUAMIB_CSipCfgVoiceServiceVoip) GetEntityData() *types.CommonEntityData {
    cSipCfgVoiceServiceVoip.EntityData.YFilter = cSipCfgVoiceServiceVoip.YFilter
    cSipCfgVoiceServiceVoip.EntityData.YangName = "cSipCfgVoiceServiceVoip"
    cSipCfgVoiceServiceVoip.EntityData.BundleName = "cisco_ios_xe"
    cSipCfgVoiceServiceVoip.EntityData.ParentYangName = "CISCO-SIP-UA-MIB"
    cSipCfgVoiceServiceVoip.EntityData.SegmentPath = "cSipCfgVoiceServiceVoip"
    cSipCfgVoiceServiceVoip.EntityData.AbsolutePath = "CISCO-SIP-UA-MIB:CISCO-SIP-UA-MIB/" + cSipCfgVoiceServiceVoip.EntityData.SegmentPath
    cSipCfgVoiceServiceVoip.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cSipCfgVoiceServiceVoip.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cSipCfgVoiceServiceVoip.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cSipCfgVoiceServiceVoip.EntityData.Children = types.NewOrderedMap()
    cSipCfgVoiceServiceVoip.EntityData.Leafs = types.NewOrderedMap()
    cSipCfgVoiceServiceVoip.EntityData.Leafs.Append("cSipCfgHeaderPassingEnabled", types.YLeaf{"CSipCfgHeaderPassingEnabled", cSipCfgVoiceServiceVoip.CSipCfgHeaderPassingEnabled})
    cSipCfgVoiceServiceVoip.EntityData.Leafs.Append("cSipCfgMaxSubscriptionAccept", types.YLeaf{"CSipCfgMaxSubscriptionAccept", cSipCfgVoiceServiceVoip.CSipCfgMaxSubscriptionAccept})
    cSipCfgVoiceServiceVoip.EntityData.Leafs.Append("cSipCfgMaxSubscriptionOriginate", types.YLeaf{"CSipCfgMaxSubscriptionOriginate", cSipCfgVoiceServiceVoip.CSipCfgMaxSubscriptionOriginate})
    cSipCfgVoiceServiceVoip.EntityData.Leafs.Append("cSipCfgSwitchTransportEnabled", types.YLeaf{"CSipCfgSwitchTransportEnabled", cSipCfgVoiceServiceVoip.CSipCfgSwitchTransportEnabled})

    cSipCfgVoiceServiceVoip.EntityData.YListKeys = []string {}

    return &(cSipCfgVoiceServiceVoip.EntityData)
}

// CISCOSIPUAMIB_CSipStatsInfo
type CISCOSIPUAMIB_CSipStatsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object reflects the total number of Trying (100) responses received by
    // the user agent since system startup.   Trying responses indicate that some
    // unspecified action is being taken on behalf of this call, but the user has
    // not yet been located.  Inbound Trying responses indicate that outbound
    // INVITE request  sent out by this system have been received and are
    // processed. The type is interface{} with range: 0..4294967295.
    CSipStatsInfoTryingIns interface{}

    // This object reflects the total number of Trying (100) responses sent by the
    // user agent since system startup. Trying responses indicate that some
    // unspecified action is being taken on behalf of this call, but the user has
    // not yet been located.  Outbound Trying responses indicate this system is
    // successfully  receiving INVITE requests and processing them on  behalf of
    // the system initiating the INVITE. The type is interface{} with range:
    // 0..4294967295.
    CSipStatsInfoTryingOuts interface{}

    // This object reflects the total number of Ringing (180) responses received
    // by the user agent since system startup. A inbound Ringing response
    // indicates that the UAS processing a INVITE initiated by this system has 
    // found a possible location where the desired end user  has registered
    // recently and is trying to alert the user. The type is interface{} with
    // range: 0..4294967295.
    CSipStatsInfoRingingIns interface{}

    // This object reflects the total number of Ringing (180) responses sent by
    // the user agent since system startup. A outbound Ringing response indicates
    // that this system has processed an INVITE for a particular end user and
    // found a possible location where that user has registered recently.  The
    // system is trying to alert the end user and is conveying that status to the
    // system that originated the INVITE. The type is interface{} with range:
    // 0..4294967295.
    CSipStatsInfoRingingOuts interface{}

    // This object reflects the total number of Call Is Being Forwarded (181)
    // responses received by the user agent since system startup. A proxy server
    // may use a Forwarded status code to indicate that the call is being
    // forwarded to a different set of destinations.  Inbound Forwarded responses
    // indicate  to this system that forwarding actions are taking place  with
    // regard to calls initiated by this system. The type is interface{} with
    // range: 0..4294967295.
    CSipStatsInfoForwardedIns interface{}

    // This object reflects the total number of Call Is Being Forwarded (181)
    // responses sent by the user agent since system startup. A proxy server may
    // use a Forwarded status code to indicate that the call is being forwarded to
    // a different set of destinations.  Outbound Forwarded responses indicate
    // this system is taking some forwarding action for calls and conveying that
    // status to the system that initiated the calls. The type is interface{} with
    // range: 0..4294967295.
    CSipStatsInfoForwardedOuts interface{}

    // This object reflects the total number of Queued (182) responses received by
    // the user agent since system startup. Inbound Queued responses indicate that
    // the users this system is attempting to call are temporarily unavailable but
    // the SIP agents operating on behalf of those users wish to queue the calls
    // rather than reject them.  When the called parties become available, this
    // system can expect to receive the appropriate final status response.  The
    // Reason-Phrase from the Queued response messages Status-Line may give
    // further details about the status of the call.  Multiple  Queued responses
    // to update this system about the status of the queued call my be received.
    // The type is interface{} with range: 0..4294967295.
    CSipStatsInfoQueuedIns interface{}

    // This object reflects the total number of Queued (182) responses sent by the
    // user agent since system startup. Outbound Queued responses indicate this
    // system has determined that the called party is temporarily unavailable but
    // the call is not rejected.  Rather  the call is queued until the called
    // party becomes available.  Queued responses messages are sent to the system
    // originating the call request to convey the current status of a queued call.
    // The type is interface{} with range: 0..4294967295.
    CSipStatsInfoQueuedOuts interface{}

    // This object reflects the total number of Session Progress (183) responses
    // received by the user agent since system startup. The type is interface{}
    // with range: 0..4294967295.
    CSipStatsInfoSessionProgIns interface{}

    // This object reflects the total number of Session Progress (183) responses
    // sent by the user agent since system startup. The type is interface{} with
    // range: 0..4294967295.
    CSipStatsInfoSessionProgOuts interface{}
}

func (cSipStatsInfo *CISCOSIPUAMIB_CSipStatsInfo) GetEntityData() *types.CommonEntityData {
    cSipStatsInfo.EntityData.YFilter = cSipStatsInfo.YFilter
    cSipStatsInfo.EntityData.YangName = "cSipStatsInfo"
    cSipStatsInfo.EntityData.BundleName = "cisco_ios_xe"
    cSipStatsInfo.EntityData.ParentYangName = "CISCO-SIP-UA-MIB"
    cSipStatsInfo.EntityData.SegmentPath = "cSipStatsInfo"
    cSipStatsInfo.EntityData.AbsolutePath = "CISCO-SIP-UA-MIB:CISCO-SIP-UA-MIB/" + cSipStatsInfo.EntityData.SegmentPath
    cSipStatsInfo.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cSipStatsInfo.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cSipStatsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cSipStatsInfo.EntityData.Children = types.NewOrderedMap()
    cSipStatsInfo.EntityData.Leafs = types.NewOrderedMap()
    cSipStatsInfo.EntityData.Leafs.Append("cSipStatsInfoTryingIns", types.YLeaf{"CSipStatsInfoTryingIns", cSipStatsInfo.CSipStatsInfoTryingIns})
    cSipStatsInfo.EntityData.Leafs.Append("cSipStatsInfoTryingOuts", types.YLeaf{"CSipStatsInfoTryingOuts", cSipStatsInfo.CSipStatsInfoTryingOuts})
    cSipStatsInfo.EntityData.Leafs.Append("cSipStatsInfoRingingIns", types.YLeaf{"CSipStatsInfoRingingIns", cSipStatsInfo.CSipStatsInfoRingingIns})
    cSipStatsInfo.EntityData.Leafs.Append("cSipStatsInfoRingingOuts", types.YLeaf{"CSipStatsInfoRingingOuts", cSipStatsInfo.CSipStatsInfoRingingOuts})
    cSipStatsInfo.EntityData.Leafs.Append("cSipStatsInfoForwardedIns", types.YLeaf{"CSipStatsInfoForwardedIns", cSipStatsInfo.CSipStatsInfoForwardedIns})
    cSipStatsInfo.EntityData.Leafs.Append("cSipStatsInfoForwardedOuts", types.YLeaf{"CSipStatsInfoForwardedOuts", cSipStatsInfo.CSipStatsInfoForwardedOuts})
    cSipStatsInfo.EntityData.Leafs.Append("cSipStatsInfoQueuedIns", types.YLeaf{"CSipStatsInfoQueuedIns", cSipStatsInfo.CSipStatsInfoQueuedIns})
    cSipStatsInfo.EntityData.Leafs.Append("cSipStatsInfoQueuedOuts", types.YLeaf{"CSipStatsInfoQueuedOuts", cSipStatsInfo.CSipStatsInfoQueuedOuts})
    cSipStatsInfo.EntityData.Leafs.Append("cSipStatsInfoSessionProgIns", types.YLeaf{"CSipStatsInfoSessionProgIns", cSipStatsInfo.CSipStatsInfoSessionProgIns})
    cSipStatsInfo.EntityData.Leafs.Append("cSipStatsInfoSessionProgOuts", types.YLeaf{"CSipStatsInfoSessionProgOuts", cSipStatsInfo.CSipStatsInfoSessionProgOuts})

    cSipStatsInfo.EntityData.YListKeys = []string {}

    return &(cSipStatsInfo.EntityData)
}

// CISCOSIPUAMIB_CSipStatsSuccess
type CISCOSIPUAMIB_CSipStatsSuccess struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object reflects the total number of Ok (200) responses received by the
    // user agent since system startup. The meaning of inbound Ok responses
    // depends on the method used in the associated request.  BYE      : The Ok
    // response means the call has             been terminated.  CANCEL   : The Ok
    // response means the search for             the end user has been cancelled. 
    // INVITE   : The Ok response means the called party             has agreed to
    // participate in the call.  OPTIONS  : The Ok response means the called party
    // has agreed to share its capabilities.  REGISTER : The Ok response means the
    // registration            has succeeded. The type is interface{} with range:
    // 0..4294967295.
    CSipStatsSuccessOkIns interface{}

    // This object reflects the total number of Ok (200) responses sent by the
    // user agent since system startup. The meaning of outbound Ok responses
    // depends on the method used in the associated request.  BYE      : The Ok
    // response means the call has             been terminated.  CANCEL   : The Ok
    // response means the search for             the end user has been cancelled. 
    // INVITE   : The Ok response means the called party             has agreed to
    // participate in the call.  OPTIONS  : The Ok response means the called party
    // has agreed to share its capabilities.  REGISTER : The Ok response means the
    // registration            has succeeded. The type is interface{} with range:
    // 0..4294967295.
    CSipStatsSuccessOkOuts interface{}

    // This object reflects the total number of Accepted (202) responses received
    // by the user agent since system startup. The meaning of outbound 202 Ok
    // responses depends on the method used in the associated request. The type is
    // interface{} with range: 0..4294967295.
    CSipStatsSuccessAcceptedIns interface{}

    // This object reflects the total number of Accepted (202) responses sent by
    // the user agent since system startup. The meaning of outbound 202 Ok
    // responses depends on the method used in the associated request. The type is
    // interface{} with range: 0..4294967295.
    CSipStatsSuccessAcceptedOuts interface{}
}

func (cSipStatsSuccess *CISCOSIPUAMIB_CSipStatsSuccess) GetEntityData() *types.CommonEntityData {
    cSipStatsSuccess.EntityData.YFilter = cSipStatsSuccess.YFilter
    cSipStatsSuccess.EntityData.YangName = "cSipStatsSuccess"
    cSipStatsSuccess.EntityData.BundleName = "cisco_ios_xe"
    cSipStatsSuccess.EntityData.ParentYangName = "CISCO-SIP-UA-MIB"
    cSipStatsSuccess.EntityData.SegmentPath = "cSipStatsSuccess"
    cSipStatsSuccess.EntityData.AbsolutePath = "CISCO-SIP-UA-MIB:CISCO-SIP-UA-MIB/" + cSipStatsSuccess.EntityData.SegmentPath
    cSipStatsSuccess.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cSipStatsSuccess.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cSipStatsSuccess.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cSipStatsSuccess.EntityData.Children = types.NewOrderedMap()
    cSipStatsSuccess.EntityData.Leafs = types.NewOrderedMap()
    cSipStatsSuccess.EntityData.Leafs.Append("cSipStatsSuccessOkIns", types.YLeaf{"CSipStatsSuccessOkIns", cSipStatsSuccess.CSipStatsSuccessOkIns})
    cSipStatsSuccess.EntityData.Leafs.Append("cSipStatsSuccessOkOuts", types.YLeaf{"CSipStatsSuccessOkOuts", cSipStatsSuccess.CSipStatsSuccessOkOuts})
    cSipStatsSuccess.EntityData.Leafs.Append("cSipStatsSuccessAcceptedIns", types.YLeaf{"CSipStatsSuccessAcceptedIns", cSipStatsSuccess.CSipStatsSuccessAcceptedIns})
    cSipStatsSuccess.EntityData.Leafs.Append("cSipStatsSuccessAcceptedOuts", types.YLeaf{"CSipStatsSuccessAcceptedOuts", cSipStatsSuccess.CSipStatsSuccessAcceptedOuts})

    cSipStatsSuccess.EntityData.YListKeys = []string {}

    return &(cSipStatsSuccess.EntityData)
}

// CISCOSIPUAMIB_CSipStatsRedirect
type CISCOSIPUAMIB_CSipStatsRedirect struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object reflects the total number of Multiple Choices (300) responses
    // received by the user agent since system startup. Multiple Choices responses
    // indicate that the called party can be reached at several different
    // locations and the server cannot or prefers not to proxy the request. The
    // type is interface{} with range: 0..4294967295.
    CSipStatsRedirMultipleChoices interface{}

    // This object reflects the total number of Moved  Permanently (301) responses
    // received by the user agent since system startup. Moved Permanently
    // responses indicate that the called party  can no longer be found at the
    // address offered in the request  and the requesting UAC should retry at the
    // new address given  by the Contact header field of the response. The type is
    // interface{} with range: 0..4294967295.
    CSipStatsRedirMovedPerms interface{}

    // This object reflects the total number of Moved  Temporarily (302) responses
    // received by the user agent since system startup. Moved Temporarily
    // responses indicate the UAC should retry the request directed at the new
    // address(es) given by the Contact header field of the response. The duration
    // of this redirection can be indicated through the Expires header.  If no
    // explicit expiration time is given, the new address(es) are only valid for
    // this call. The type is interface{} with range: 0..4294967295.
    CSipStatsRedirMovedTemps interface{}

    // This object reflects the total number of See Other  (303) responses
    // received by the user agent since system startup. The type is interface{}
    // with range: 0..4294967295.
    CSipStatsRedirSeeOthers interface{}

    // This object reflects the total number of Use Proxy  (305) responses
    // received by the user agent since system startup. See Other responses
    // indicate that requested resources must be accessed through the proxy given
    // by the  Contact header field of the response.  The recipient of this
    // response is expected to repeat this single request via the proxy. The type
    // is interface{} with range: 0..4294967295.
    CSipStatsRedirUseProxys interface{}

    // This object reflects the total number of Alternative Service (380)
    // responses received by the user agent since system startup. Alternative
    // Service responses indicate that the call was not successful, but
    // alternative services are possible.  Those alternative services are
    // described in the message body of the response. The type is interface{} with
    // range: 0..4294967295.
    CSipStatsRedirAltServices interface{}

    // This object reflects the total number of Moved Temporarily (302) responses
    // received by the user agent since system startup.  Moved Temporarily
    // responses indicate the UAC should retry the request directed at the new
    // address(es)  given by the Contact header field of the response. The
    // duration of this redirection can be indicated through the Expires header. 
    // If no explicit expiration time is given, the new address(es) are only valid
    // for this call. The type is interface{} with range: 0..4294967295.
    CSipStatsRedirMovedTempsIns interface{}

    // This object reflects the total number of Moved Temporarily (302) responses
    // sent by the user agent since system startup. The type is interface{} with
    // range: 0..4294967295.
    CSipStatsRedirMovedTempsOuts interface{}
}

func (cSipStatsRedirect *CISCOSIPUAMIB_CSipStatsRedirect) GetEntityData() *types.CommonEntityData {
    cSipStatsRedirect.EntityData.YFilter = cSipStatsRedirect.YFilter
    cSipStatsRedirect.EntityData.YangName = "cSipStatsRedirect"
    cSipStatsRedirect.EntityData.BundleName = "cisco_ios_xe"
    cSipStatsRedirect.EntityData.ParentYangName = "CISCO-SIP-UA-MIB"
    cSipStatsRedirect.EntityData.SegmentPath = "cSipStatsRedirect"
    cSipStatsRedirect.EntityData.AbsolutePath = "CISCO-SIP-UA-MIB:CISCO-SIP-UA-MIB/" + cSipStatsRedirect.EntityData.SegmentPath
    cSipStatsRedirect.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cSipStatsRedirect.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cSipStatsRedirect.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cSipStatsRedirect.EntityData.Children = types.NewOrderedMap()
    cSipStatsRedirect.EntityData.Leafs = types.NewOrderedMap()
    cSipStatsRedirect.EntityData.Leafs.Append("cSipStatsRedirMultipleChoices", types.YLeaf{"CSipStatsRedirMultipleChoices", cSipStatsRedirect.CSipStatsRedirMultipleChoices})
    cSipStatsRedirect.EntityData.Leafs.Append("cSipStatsRedirMovedPerms", types.YLeaf{"CSipStatsRedirMovedPerms", cSipStatsRedirect.CSipStatsRedirMovedPerms})
    cSipStatsRedirect.EntityData.Leafs.Append("cSipStatsRedirMovedTemps", types.YLeaf{"CSipStatsRedirMovedTemps", cSipStatsRedirect.CSipStatsRedirMovedTemps})
    cSipStatsRedirect.EntityData.Leafs.Append("cSipStatsRedirSeeOthers", types.YLeaf{"CSipStatsRedirSeeOthers", cSipStatsRedirect.CSipStatsRedirSeeOthers})
    cSipStatsRedirect.EntityData.Leafs.Append("cSipStatsRedirUseProxys", types.YLeaf{"CSipStatsRedirUseProxys", cSipStatsRedirect.CSipStatsRedirUseProxys})
    cSipStatsRedirect.EntityData.Leafs.Append("cSipStatsRedirAltServices", types.YLeaf{"CSipStatsRedirAltServices", cSipStatsRedirect.CSipStatsRedirAltServices})
    cSipStatsRedirect.EntityData.Leafs.Append("cSipStatsRedirMovedTempsIns", types.YLeaf{"CSipStatsRedirMovedTempsIns", cSipStatsRedirect.CSipStatsRedirMovedTempsIns})
    cSipStatsRedirect.EntityData.Leafs.Append("cSipStatsRedirMovedTempsOuts", types.YLeaf{"CSipStatsRedirMovedTempsOuts", cSipStatsRedirect.CSipStatsRedirMovedTempsOuts})

    cSipStatsRedirect.EntityData.YListKeys = []string {}

    return &(cSipStatsRedirect.EntityData)
}

// CISCOSIPUAMIB_CSipStatsErrClient
type CISCOSIPUAMIB_CSipStatsErrClient struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object reflects the total number of Bad Request (400)  responses
    // received by the user agent since system startup. Inbound Bad Request
    // responses indicate that requests issued  by this system could not be
    // understood due to malformed  syntax. The type is interface{} with range:
    // 0..4294967295.
    CSipStatsClientBadRequestIns interface{}

    // This object reflects the total number of Bad Request (400)  responses sent
    // by the user agent since system startup. Outbound Bad Request responses
    // indicate that requests  received by this system could not be understood due
    // to  malformed syntax. The type is interface{} with range: 0..4294967295.
    CSipStatsClientBadRequestOuts interface{}

    // This object reflects the total number of Unauthorized (401)  responses
    // received by the user agent since system startup.   Inbound Unauthorized
    // responses indicate that requests issued  by this system require user
    // authentication. The type is interface{} with range: 0..4294967295.
    CSipStatsClientUnauthorizedIns interface{}

    // This object reflects the total number of Unauthorized (401)  responses sent
    // by the user agent since system startup. Outbound Unauthorized responses
    // indicate that requests  received by this system require user
    // authentication. The type is interface{} with range: 0..4294967295.
    CSipStatsClientUnauthorizedOuts interface{}

    // This object reflects the total number of Payment Required  (402) responses
    // received by the user agent since system startup. The type is interface{}
    // with range: 0..4294967295.
    CSipStatsClientPaymentReqdIns interface{}

    // This object reflects the total number of Payment Required  (402) responses
    // sent by the user agent since system startup. The type is interface{} with
    // range: 0..4294967295.
    CSipStatsClientPaymentReqdOuts interface{}

    // This object reflects the total number of Forbidden (403)  responses
    // received by the user agent since system startup. Inbound Forbidden
    // responses indicate that requests issued by this system are understood by
    // the server but the server refuses to fulfill the request.  Authorization
    // will not help and the requests should not be repeated. The type is
    // interface{} with range: 0..4294967295.
    CSipStatsClientForbiddenIns interface{}

    // This object reflects the total number of Forbidden (403)  responses sent by
    // the user agent since system startup. Outbound Forbidden responses indicate
    // that requests received by this system are understood but this system is
    // refusing to fulfill the requests. The type is interface{} with range:
    // 0..4294967295.
    CSipStatsClientForbiddenOuts interface{}

    // This object reflects the total number of Not Found (404)  responses
    // received by the user agent since system startup. Inbound Not Found
    // responses indicate that the called party  does not exist at the domain
    // specified in the Request-URI  or the domain is not handled by the recipient
    // of the request. The type is interface{} with range: 0..4294967295.
    CSipStatsClientNotFoundIns interface{}

    // This object reflects the total number of Not Found (404)  responses sent by
    // the user agent since system startup. Outbound Not Found responses indicate
    // that this system knows that the called party does not exist at the domain
    // specified in the Request-URI or the domain is not handled by this system.
    // The type is interface{} with range: 0..4294967295.
    CSipStatsClientNotFoundOuts interface{}

    // This object reflects the total number of Method Not Allowed  (405) received
    // responses by the user agent. Inbound Method Not Allowed responses indicate
    // that requests  issued by this system have specified a SIP method in the 
    // Request-Line that is not allowed for the address identified  by the
    // Request-URI. The type is interface{} with range: 0..4294967295.
    CSipStatsClientMethNotAllowedIns interface{}

    // This object reflects the total number of Method Not Allowed  (405) received
    // sent by the user agent since system startup. Outbound Method Not Allowed
    // responses indicate that requests  received by this system have SIP methods
    // specified in the  Request-Line that are not allowed for the address
    // identified  by the Request-URI. The type is interface{} with range:
    // 0..4294967295.
    CSipStatsClientMethNotAllowedOuts interface{}

    // This object reflects the total number of Not Acceptable  (406) responses
    // received by the user agent since system startup. Inbound Not Acceptable
    // responses indicate the resources  identified by requests issued by this
    // system cannot generate  responses with content characteristics acceptable
    // to this  system according to the accept headers sent in the requests. The
    // type is interface{} with range: 0..4294967295.
    CSipStatsClientNotAcceptableIns interface{}

    // This object reflects the total number of Not Acceptable (406)  responses
    // sent by the user agent since system startup. Outbound Not Acceptable
    // responses indicate that the resources identified by requests received by
    // this system cannot generate responses with content characteristics
    // acceptable to the  system sending the requests. The type is interface{}
    // with range: 0..4294967295.
    CSipStatsClientNotAcceptableOuts interface{}

    // This object reflects the total number of Proxy Authentication  Required
    // (407) responses received by the user agent since system startup. Inbound
    // Proxy Authentication Required responses indicate that  this system must
    // authenticate itself with the proxy before  gaining access to the requested
    // resource. The type is interface{} with range: 0..4294967295.
    CSipStatsClientProxyAuthReqdIns interface{}

    // This object reflects the total number of Proxy Authentication  Required
    // (407) responses sent by the user agent since system startup. Outbound Proxy
    // Authentication Required responses indicate that the systems issuing
    // requests being processed by this system  must authenticate themselves with
    // this system before gaining  access to requested resources. The type is
    // interface{} with range: 0..4294967295.
    CSipStatsClientProxyAuthReqdOuts interface{}

    // This object reflects the total number of Request Timeout  (408) responses
    // received by the user agent since system startup. Inbound Request Timeout
    // responses indicate that requests  issued by this system are not being
    // processed by the server  within the time indicated in the Expires header of
    // the  request. The type is interface{} with range: 0..4294967295.
    CSipStatsClientReqTimeoutIns interface{}

    // This object reflects the total number of Request Timeout  (408) responses
    // sent by the user agent since system startup. Outbound Request Timeout
    // responses indicate that this  system is not able to produce an appropriate
    // response within  the time indicated in the Expires header of the request.
    // The type is interface{} with range: 0..4294967295.
    CSipStatsClientReqTimeoutOuts interface{}

    // This object reflects the total number of Conflict (409)  responses received
    // by the user agent since system startup. Inbound Conflict responses indicate
    // that requests issued by this system could not be completed due to a
    // conflict with the current state of a requested resource. The type is
    // interface{} with range: 0..4294967295.
    CSipStatsClientConflictIns interface{}

    // This object reflects the total number of Conflict (409)  responses sent by
    // the user agent since system startup. Outbound Conflict responses indicate
    // that requests received by this system could not be completed due to a
    // conflict with the current state of a requested resource. The type is
    // interface{} with range: 0..4294967295.
    CSipStatsClientConflictOuts interface{}

    // This object reflects the total number of Gone (410)  responses received by
    // the user agent since system startup. Inbound Gone responses indicate that
    // resources requested by this system are no longer available at the recipient
    // server and no forwarding address is known. The type is interface{} with
    // range: 0..4294967295.
    CSipStatsClientGoneIns interface{}

    // This object reflects the total number of Gone (410)  responses sent by the
    // user agent since system startup. Outbound Gone responses indicate that the
    // requested resources are no longer available at this system and no
    // forwarding address is known. The type is interface{} with range:
    // 0..4294967295.
    CSipStatsClientGoneOuts interface{}

    // This object reflects the total number of Length Required  (411) responses
    // received by the user agent since system startup. Inbound Length Required
    // responses indicate that requests  issued by this system are being refused
    // by servers because  of no defined Content-Length header field. The type is
    // interface{} with range: 0..4294967295.
    CSipStatsClientLengthRequiredIns interface{}

    // This object reflects the total number of Length Required  (411) responses
    // sent by the user agent since system startup. Outbound Length Required
    // responses indicate that requests  received by this system are being refused
    // because of no  defined Content-Length header field. The type is interface{}
    // with range: 0..4294967295.
    CSipStatsClientLengthRequiredOuts interface{}

    // This object reflects the total number of Request Entity Too  Large (413)
    // responses received by the user agent since system startup. Inbound Request
    // Entity Too Large responses indicate that  requests issued by this system
    // are being refused because  the request is larger than the server is willing
    // or able to  process. The type is interface{} with range: 0..4294967295.
    CSipStatsClientReqEntTooLargeIns interface{}

    // This object reflects the total number of Request Entity Too  Large (413)
    // responses sent by the user agent since system startup. Outbound Request
    // Entity Too Large responses indicate that  requests received by this system
    // are larger than this system  is willing or able to process. The type is
    // interface{} with range: 0..4294967295.
    CSipStatsClientReqEntTooLargeOuts interface{}

    // This object reflects the total number of Request-URI Too  Large (414)
    // responses received by the user agent since system startup. Inbound
    // Request-URI Too Large responses indicate that  requests issued by this
    // system are being refused because the  Request-URI is longer than the server
    // is willing or able to  interpret. The type is interface{} with range:
    // 0..4294967295.
    CSipStatsClientReqURITooLargeIns interface{}

    // This object reflects the total number of Request-URI Too  Large (414)
    // responses sent by the user agent since system startup. Outbound Request-URI
    // Too Large responses indicate that  Request-URIs received by this system are
    // longer than this  system is willing or able to interpret. The type is
    // interface{} with range: 0..4294967295.
    CSipStatsClientReqURITooLargeOuts interface{}

    // This object reflects the total number of Unsupported Media  Type (415)
    // responses received by the user agent since system startup. Inbound
    // Unsupported Media Type responses indicate that  requests issued by this
    // system are being refused because the  message body of the request is in a
    // format not supported by  the requested resource for the requested method.
    // The type is interface{} with range: 0..4294967295.
    CSipStatsClientNoSupMediaTypeIns interface{}

    // This object reflects the total number of Unsupported Media  Type (415)
    // responses sent by the user agent since system startup. Outbound Unsupported
    // Media Type responses indicate that the  body of requests received by this
    // system are in a format not  supported by the requested resource for the
    // requested method. The type is interface{} with range: 0..4294967295.
    CSipStatsClientNoSupMediaTypeOuts interface{}

    // This object reflects the total number of Bad Extension (420)  responses
    // received by the user agent since system startup. Inbound Bad Extension
    // responses indicate that the recipient  did not understand the protocol
    // extension specified in a  Require header field. The type is interface{}
    // with range: 0..4294967295.
    CSipStatsClientBadExtensionIns interface{}

    // This object reflects the total number of Bad Extension (420)  responses
    // sent by the user agent since system startup. Outbound Bad Extension
    // responses indicate that this system did not understand the protocol
    // extension specified in a Require header field of requests. The type is
    // interface{} with range: 0..4294967295.
    CSipStatsClientBadExtensionOuts interface{}

    // This object reflects the total number of Temporarily Not  Available (480)
    // responses received by the user agent since system startup. Inbound
    // Temporarily Not Available responses indicate that  the called party is
    // currently unavailable. The type is interface{} with range: 0..4294967295.
    CSipStatsClientTempNotAvailIns interface{}

    // This object reflects the total number of Temporarily Not  Available (480)
    // responses sent by the user agent since system startup. Outbound Temporarily
    // Not Available responses indicate that  the called party's end system was
    // contacted successfully but  the called party is currently unavailable. The
    // type is interface{} with range: 0..4294967295.
    CSipStatsClientTempNotAvailOuts interface{}

    // This object reflects the total number of Call Leg/Transaction  Does Not
    // Exist (481) responses received by the user agent since system startup.
    // Inbound Call Leg/Transaction Does Not Exist responses indicate that either
    // BYE or CANCEL requests issued by this system were  received by a server and
    // not matching call leg or transaction  existed. The type is interface{} with
    // range: 0..4294967295.
    CSipStatsClientCallLegNoExistIns interface{}

    // This object reflects the total number of Call Leg/Transaction  Does Not
    // Exist (481) responses sent by the user agent since system startup. Outbound
    // Call Leg/Transaction Does Not Exist responses  indicate that BYE or CANCEL
    // requests have been received by  this system and not call leg or transaction
    // matching that  request exists. The type is interface{} with range:
    // 0..4294967295.
    CSipStatsClientCallLegNoExistOuts interface{}

    // This object reflects the total number of Loop Detected (482)  responses
    // received by the user agent since system startup. Inbound Loop Detected
    // responses indicate that requests issued by this system were received at
    // servers and the server found  itself in the Via path more than once. The
    // type is interface{} with range: 0..4294967295.
    CSipStatsClientLoopDetectedIns interface{}

    // This object reflects the total number of Loop Detected (482)  responses
    // sent by the user agent since system startup. Outbound Loop Detected
    // responses indicate that requests  received by this system contain a Via
    // path with this system  appearing more than once. The type is interface{}
    // with range: 0..4294967295.
    CSipStatsClientLoopDetectedOuts interface{}

    // This object reflects the total number of Too Many Hops (483)  responses
    // received by the user agent since system startup. Inbound Too Many Hops
    // responses indicate that requests issued by this system contain more Via
    // entries (hops) than allowed by the Max-Forwards header field of the
    // requests. The type is interface{} with range: 0..4294967295.
    CSipStatsClientTooManyHopsIns interface{}

    // This object reflects the total number of Too Many Hops (483)  responses
    // sent by the user agent since system startup. Outbound Too Many Hops
    // responses indicate that requests received by this system contain more Via
    // entries (hops) than are allowed by the Max-Forwards header field of the
    // requests. The type is interface{} with range: 0..4294967295.
    CSipStatsClientTooManyHopsOuts interface{}

    // This object reflects the total number of Address Incomplete  (484)
    // responses received by the user agent since system startup. Inbound Address
    // Incomplete responses indicate that requests  issued by this system had To
    // addresses or Request-URIs that  were incomplete. The type is interface{}
    // with range: 0..4294967295.
    CSipStatsClientAddrIncompleteIns interface{}

    // This object reflects the total number of Address Incomplete  (484)
    // responses sent by the user agent since system startup. Outbound Address
    // Incomplete responses indicate that requests  received by this system had To
    // addresses or Request-URIs that  were incomplete. The type is interface{}
    // with range: 0..4294967295.
    CSipStatsClientAddrIncompleteOuts interface{}

    // This object reflects the total number of Ambiguous (485)  responses
    // received by the user agent since system startup. Inbound Ambiguous
    // responses indicate that requests issued by this system provided ambiguous
    // address information. The type is interface{} with range: 0..4294967295.
    CSipStatsClientAmbiguousIns interface{}

    // This object reflects the total number of Ambiguous (485)  responses sent by
    // the user agent since system startup. Outbound Ambiguous responses indicate
    // that requests received by this system contained ambiguous address
    // information. The type is interface{} with range: 0..4294967295.
    CSipStatsClientAmbiguousOuts interface{}

    // This object reflects the total number of Busy Here (486)  responses
    // received by the user agent since system startup. Inbound Busy Here
    // responses indicate that the called party is currently not willing or not
    // able to take additional calls. The type is interface{} with range:
    // 0..4294967295.
    CSipStatsClientBusyHereIns interface{}

    // This object reflects the total number of Busy Here (486)  responses sent by
    // the user agent since system startup. Outbound Busy Here responses indicate
    // that the called party's end system was contacted successfully but the
    // called party is currently not willing or able to take  additional calls.
    // The type is interface{} with range: 0..4294967295.
    CSipStatsClientBusyHereOuts interface{}

    // This object reflects the total number of Request Terminated  (487)
    // responses received by the user agent since system startup. Request
    // Terminated responses are issued if the original  request was terminated via
    // CANCEL or BYE. The type is interface{} with range: 0..4294967295.
    CSipStatsClientReqTermIns interface{}

    // This object reflects the total number of Request Terminated  (487)
    // responses sent by the user agent since system startup. Request Terminated
    // responses are issued if the original  request was terminated via CANCEL or
    // BYE. The type is interface{} with range: 0..4294967295.
    CSipStatsClientReqTermOuts interface{}

    // This object reflects the total number of Not Acceptable Here (488)
    // responses received by the user agent since system startup. The response has
    // the same meaning as 606 (Not Acceptable),  but only applies to the specific
    // entity addressed by the  Request-URI and the request may succeed elsewhere.
    // The type is interface{} with range: 0..4294967295.
    CSipStatsClientNoAcceptHereIns interface{}

    // This object reflects the total number of Not Acceptable Here (488)
    // responses sent by the user agent since system startup. The response has the
    // same meaning as 606 (Not Acceptable),  but only applies to the specific
    // entity addressed by the  Request-URI and the request may succeed elsewhere.
    // The type is interface{} with range: 0..4294967295.
    CSipStatsClientNoAcceptHereOuts interface{}

    // This object reflects the total number of BadEvent (489)  responses received
    // by the user agent since system startup. The type is interface{} with range:
    // 0..4294967295.
    CSipStatsClientBadEventIns interface{}

    // This object reflects the total number of BadEvent (489)  responses sent by
    // the user agent since system startup. The type is interface{} with range:
    // 0..4294967295.
    CSipStatsClientBadEventOuts interface{}

    // This object reflects the total number of SessionTimerTooSmall (422)
    // responses received by the user agent since system  startup. The type is
    // interface{} with range: 0..4294967295.
    CSipStatsClientSTTooSmallIns interface{}

    // This object reflects the total number of SessionTimerTooSmall (422)
    // responses sent by the user agent since system startup. The type is
    // interface{} with range: 0..4294967295.
    CSipStatsClientSTTooSmallOuts interface{}

    // This object reflects the total number of RequestPending (491) responses
    // received by the user agent since system startup. The type is interface{}
    // with range: 0..4294967295.
    CSipStatsClientReqPendingIns interface{}

    // This object reflects the total number of RequestPending (491) responses
    // sent by the user agent since system startup. The type is interface{} with
    // range: 0..4294967295.
    CSipStatsClientReqPendingOuts interface{}
}

func (cSipStatsErrClient *CISCOSIPUAMIB_CSipStatsErrClient) GetEntityData() *types.CommonEntityData {
    cSipStatsErrClient.EntityData.YFilter = cSipStatsErrClient.YFilter
    cSipStatsErrClient.EntityData.YangName = "cSipStatsErrClient"
    cSipStatsErrClient.EntityData.BundleName = "cisco_ios_xe"
    cSipStatsErrClient.EntityData.ParentYangName = "CISCO-SIP-UA-MIB"
    cSipStatsErrClient.EntityData.SegmentPath = "cSipStatsErrClient"
    cSipStatsErrClient.EntityData.AbsolutePath = "CISCO-SIP-UA-MIB:CISCO-SIP-UA-MIB/" + cSipStatsErrClient.EntityData.SegmentPath
    cSipStatsErrClient.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cSipStatsErrClient.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cSipStatsErrClient.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cSipStatsErrClient.EntityData.Children = types.NewOrderedMap()
    cSipStatsErrClient.EntityData.Leafs = types.NewOrderedMap()
    cSipStatsErrClient.EntityData.Leafs.Append("cSipStatsClientBadRequestIns", types.YLeaf{"CSipStatsClientBadRequestIns", cSipStatsErrClient.CSipStatsClientBadRequestIns})
    cSipStatsErrClient.EntityData.Leafs.Append("cSipStatsClientBadRequestOuts", types.YLeaf{"CSipStatsClientBadRequestOuts", cSipStatsErrClient.CSipStatsClientBadRequestOuts})
    cSipStatsErrClient.EntityData.Leafs.Append("cSipStatsClientUnauthorizedIns", types.YLeaf{"CSipStatsClientUnauthorizedIns", cSipStatsErrClient.CSipStatsClientUnauthorizedIns})
    cSipStatsErrClient.EntityData.Leafs.Append("cSipStatsClientUnauthorizedOuts", types.YLeaf{"CSipStatsClientUnauthorizedOuts", cSipStatsErrClient.CSipStatsClientUnauthorizedOuts})
    cSipStatsErrClient.EntityData.Leafs.Append("cSipStatsClientPaymentReqdIns", types.YLeaf{"CSipStatsClientPaymentReqdIns", cSipStatsErrClient.CSipStatsClientPaymentReqdIns})
    cSipStatsErrClient.EntityData.Leafs.Append("cSipStatsClientPaymentReqdOuts", types.YLeaf{"CSipStatsClientPaymentReqdOuts", cSipStatsErrClient.CSipStatsClientPaymentReqdOuts})
    cSipStatsErrClient.EntityData.Leafs.Append("cSipStatsClientForbiddenIns", types.YLeaf{"CSipStatsClientForbiddenIns", cSipStatsErrClient.CSipStatsClientForbiddenIns})
    cSipStatsErrClient.EntityData.Leafs.Append("cSipStatsClientForbiddenOuts", types.YLeaf{"CSipStatsClientForbiddenOuts", cSipStatsErrClient.CSipStatsClientForbiddenOuts})
    cSipStatsErrClient.EntityData.Leafs.Append("cSipStatsClientNotFoundIns", types.YLeaf{"CSipStatsClientNotFoundIns", cSipStatsErrClient.CSipStatsClientNotFoundIns})
    cSipStatsErrClient.EntityData.Leafs.Append("cSipStatsClientNotFoundOuts", types.YLeaf{"CSipStatsClientNotFoundOuts", cSipStatsErrClient.CSipStatsClientNotFoundOuts})
    cSipStatsErrClient.EntityData.Leafs.Append("cSipStatsClientMethNotAllowedIns", types.YLeaf{"CSipStatsClientMethNotAllowedIns", cSipStatsErrClient.CSipStatsClientMethNotAllowedIns})
    cSipStatsErrClient.EntityData.Leafs.Append("cSipStatsClientMethNotAllowedOuts", types.YLeaf{"CSipStatsClientMethNotAllowedOuts", cSipStatsErrClient.CSipStatsClientMethNotAllowedOuts})
    cSipStatsErrClient.EntityData.Leafs.Append("cSipStatsClientNotAcceptableIns", types.YLeaf{"CSipStatsClientNotAcceptableIns", cSipStatsErrClient.CSipStatsClientNotAcceptableIns})
    cSipStatsErrClient.EntityData.Leafs.Append("cSipStatsClientNotAcceptableOuts", types.YLeaf{"CSipStatsClientNotAcceptableOuts", cSipStatsErrClient.CSipStatsClientNotAcceptableOuts})
    cSipStatsErrClient.EntityData.Leafs.Append("cSipStatsClientProxyAuthReqdIns", types.YLeaf{"CSipStatsClientProxyAuthReqdIns", cSipStatsErrClient.CSipStatsClientProxyAuthReqdIns})
    cSipStatsErrClient.EntityData.Leafs.Append("cSipStatsClientProxyAuthReqdOuts", types.YLeaf{"CSipStatsClientProxyAuthReqdOuts", cSipStatsErrClient.CSipStatsClientProxyAuthReqdOuts})
    cSipStatsErrClient.EntityData.Leafs.Append("cSipStatsClientReqTimeoutIns", types.YLeaf{"CSipStatsClientReqTimeoutIns", cSipStatsErrClient.CSipStatsClientReqTimeoutIns})
    cSipStatsErrClient.EntityData.Leafs.Append("cSipStatsClientReqTimeoutOuts", types.YLeaf{"CSipStatsClientReqTimeoutOuts", cSipStatsErrClient.CSipStatsClientReqTimeoutOuts})
    cSipStatsErrClient.EntityData.Leafs.Append("cSipStatsClientConflictIns", types.YLeaf{"CSipStatsClientConflictIns", cSipStatsErrClient.CSipStatsClientConflictIns})
    cSipStatsErrClient.EntityData.Leafs.Append("cSipStatsClientConflictOuts", types.YLeaf{"CSipStatsClientConflictOuts", cSipStatsErrClient.CSipStatsClientConflictOuts})
    cSipStatsErrClient.EntityData.Leafs.Append("cSipStatsClientGoneIns", types.YLeaf{"CSipStatsClientGoneIns", cSipStatsErrClient.CSipStatsClientGoneIns})
    cSipStatsErrClient.EntityData.Leafs.Append("cSipStatsClientGoneOuts", types.YLeaf{"CSipStatsClientGoneOuts", cSipStatsErrClient.CSipStatsClientGoneOuts})
    cSipStatsErrClient.EntityData.Leafs.Append("cSipStatsClientLengthRequiredIns", types.YLeaf{"CSipStatsClientLengthRequiredIns", cSipStatsErrClient.CSipStatsClientLengthRequiredIns})
    cSipStatsErrClient.EntityData.Leafs.Append("cSipStatsClientLengthRequiredOuts", types.YLeaf{"CSipStatsClientLengthRequiredOuts", cSipStatsErrClient.CSipStatsClientLengthRequiredOuts})
    cSipStatsErrClient.EntityData.Leafs.Append("cSipStatsClientReqEntTooLargeIns", types.YLeaf{"CSipStatsClientReqEntTooLargeIns", cSipStatsErrClient.CSipStatsClientReqEntTooLargeIns})
    cSipStatsErrClient.EntityData.Leafs.Append("cSipStatsClientReqEntTooLargeOuts", types.YLeaf{"CSipStatsClientReqEntTooLargeOuts", cSipStatsErrClient.CSipStatsClientReqEntTooLargeOuts})
    cSipStatsErrClient.EntityData.Leafs.Append("cSipStatsClientReqURITooLargeIns", types.YLeaf{"CSipStatsClientReqURITooLargeIns", cSipStatsErrClient.CSipStatsClientReqURITooLargeIns})
    cSipStatsErrClient.EntityData.Leafs.Append("cSipStatsClientReqURITooLargeOuts", types.YLeaf{"CSipStatsClientReqURITooLargeOuts", cSipStatsErrClient.CSipStatsClientReqURITooLargeOuts})
    cSipStatsErrClient.EntityData.Leafs.Append("cSipStatsClientNoSupMediaTypeIns", types.YLeaf{"CSipStatsClientNoSupMediaTypeIns", cSipStatsErrClient.CSipStatsClientNoSupMediaTypeIns})
    cSipStatsErrClient.EntityData.Leafs.Append("cSipStatsClientNoSupMediaTypeOuts", types.YLeaf{"CSipStatsClientNoSupMediaTypeOuts", cSipStatsErrClient.CSipStatsClientNoSupMediaTypeOuts})
    cSipStatsErrClient.EntityData.Leafs.Append("cSipStatsClientBadExtensionIns", types.YLeaf{"CSipStatsClientBadExtensionIns", cSipStatsErrClient.CSipStatsClientBadExtensionIns})
    cSipStatsErrClient.EntityData.Leafs.Append("cSipStatsClientBadExtensionOuts", types.YLeaf{"CSipStatsClientBadExtensionOuts", cSipStatsErrClient.CSipStatsClientBadExtensionOuts})
    cSipStatsErrClient.EntityData.Leafs.Append("cSipStatsClientTempNotAvailIns", types.YLeaf{"CSipStatsClientTempNotAvailIns", cSipStatsErrClient.CSipStatsClientTempNotAvailIns})
    cSipStatsErrClient.EntityData.Leafs.Append("cSipStatsClientTempNotAvailOuts", types.YLeaf{"CSipStatsClientTempNotAvailOuts", cSipStatsErrClient.CSipStatsClientTempNotAvailOuts})
    cSipStatsErrClient.EntityData.Leafs.Append("cSipStatsClientCallLegNoExistIns", types.YLeaf{"CSipStatsClientCallLegNoExistIns", cSipStatsErrClient.CSipStatsClientCallLegNoExistIns})
    cSipStatsErrClient.EntityData.Leafs.Append("cSipStatsClientCallLegNoExistOuts", types.YLeaf{"CSipStatsClientCallLegNoExistOuts", cSipStatsErrClient.CSipStatsClientCallLegNoExistOuts})
    cSipStatsErrClient.EntityData.Leafs.Append("cSipStatsClientLoopDetectedIns", types.YLeaf{"CSipStatsClientLoopDetectedIns", cSipStatsErrClient.CSipStatsClientLoopDetectedIns})
    cSipStatsErrClient.EntityData.Leafs.Append("cSipStatsClientLoopDetectedOuts", types.YLeaf{"CSipStatsClientLoopDetectedOuts", cSipStatsErrClient.CSipStatsClientLoopDetectedOuts})
    cSipStatsErrClient.EntityData.Leafs.Append("cSipStatsClientTooManyHopsIns", types.YLeaf{"CSipStatsClientTooManyHopsIns", cSipStatsErrClient.CSipStatsClientTooManyHopsIns})
    cSipStatsErrClient.EntityData.Leafs.Append("cSipStatsClientTooManyHopsOuts", types.YLeaf{"CSipStatsClientTooManyHopsOuts", cSipStatsErrClient.CSipStatsClientTooManyHopsOuts})
    cSipStatsErrClient.EntityData.Leafs.Append("cSipStatsClientAddrIncompleteIns", types.YLeaf{"CSipStatsClientAddrIncompleteIns", cSipStatsErrClient.CSipStatsClientAddrIncompleteIns})
    cSipStatsErrClient.EntityData.Leafs.Append("cSipStatsClientAddrIncompleteOuts", types.YLeaf{"CSipStatsClientAddrIncompleteOuts", cSipStatsErrClient.CSipStatsClientAddrIncompleteOuts})
    cSipStatsErrClient.EntityData.Leafs.Append("cSipStatsClientAmbiguousIns", types.YLeaf{"CSipStatsClientAmbiguousIns", cSipStatsErrClient.CSipStatsClientAmbiguousIns})
    cSipStatsErrClient.EntityData.Leafs.Append("cSipStatsClientAmbiguousOuts", types.YLeaf{"CSipStatsClientAmbiguousOuts", cSipStatsErrClient.CSipStatsClientAmbiguousOuts})
    cSipStatsErrClient.EntityData.Leafs.Append("cSipStatsClientBusyHereIns", types.YLeaf{"CSipStatsClientBusyHereIns", cSipStatsErrClient.CSipStatsClientBusyHereIns})
    cSipStatsErrClient.EntityData.Leafs.Append("cSipStatsClientBusyHereOuts", types.YLeaf{"CSipStatsClientBusyHereOuts", cSipStatsErrClient.CSipStatsClientBusyHereOuts})
    cSipStatsErrClient.EntityData.Leafs.Append("cSipStatsClientReqTermIns", types.YLeaf{"CSipStatsClientReqTermIns", cSipStatsErrClient.CSipStatsClientReqTermIns})
    cSipStatsErrClient.EntityData.Leafs.Append("cSipStatsClientReqTermOuts", types.YLeaf{"CSipStatsClientReqTermOuts", cSipStatsErrClient.CSipStatsClientReqTermOuts})
    cSipStatsErrClient.EntityData.Leafs.Append("cSipStatsClientNoAcceptHereIns", types.YLeaf{"CSipStatsClientNoAcceptHereIns", cSipStatsErrClient.CSipStatsClientNoAcceptHereIns})
    cSipStatsErrClient.EntityData.Leafs.Append("cSipStatsClientNoAcceptHereOuts", types.YLeaf{"CSipStatsClientNoAcceptHereOuts", cSipStatsErrClient.CSipStatsClientNoAcceptHereOuts})
    cSipStatsErrClient.EntityData.Leafs.Append("cSipStatsClientBadEventIns", types.YLeaf{"CSipStatsClientBadEventIns", cSipStatsErrClient.CSipStatsClientBadEventIns})
    cSipStatsErrClient.EntityData.Leafs.Append("cSipStatsClientBadEventOuts", types.YLeaf{"CSipStatsClientBadEventOuts", cSipStatsErrClient.CSipStatsClientBadEventOuts})
    cSipStatsErrClient.EntityData.Leafs.Append("cSipStatsClientSTTooSmallIns", types.YLeaf{"CSipStatsClientSTTooSmallIns", cSipStatsErrClient.CSipStatsClientSTTooSmallIns})
    cSipStatsErrClient.EntityData.Leafs.Append("cSipStatsClientSTTooSmallOuts", types.YLeaf{"CSipStatsClientSTTooSmallOuts", cSipStatsErrClient.CSipStatsClientSTTooSmallOuts})
    cSipStatsErrClient.EntityData.Leafs.Append("cSipStatsClientReqPendingIns", types.YLeaf{"CSipStatsClientReqPendingIns", cSipStatsErrClient.CSipStatsClientReqPendingIns})
    cSipStatsErrClient.EntityData.Leafs.Append("cSipStatsClientReqPendingOuts", types.YLeaf{"CSipStatsClientReqPendingOuts", cSipStatsErrClient.CSipStatsClientReqPendingOuts})

    cSipStatsErrClient.EntityData.YListKeys = []string {}

    return &(cSipStatsErrClient.EntityData)
}

// CISCOSIPUAMIB_CSipStatsErrServer
type CISCOSIPUAMIB_CSipStatsErrServer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object reflects the total number of Internal Server Error (500)
    // responses received by the user agent since system startup. Inbound Internal
    // Server Error responses indicate that servers  to which this system is
    // sending requests have encountered  unexpected conditions that prevent them
    // from fulfilling the  requests. The type is interface{} with range:
    // 0..4294967295.
    CSipStatsServerIntErrorIns interface{}

    // This object reflects the total number of Internal Server Error (500)
    // responses sent by the user agent since system startup. Outbound Internal
    // Server Error responses indicate that this  system has encountered
    // unexpected conditions that prevent it  from fulfilling received requests.
    // The type is interface{} with range: 0..4294967295.
    CSipStatsServerIntErrorOuts interface{}

    // This object reflects the total number of Not Implemented  (501) responses
    // received by the user agent since system startup. Inbound Not Implemented
    // responses indicate that servers to  which this system is sending requests
    // do not support the  functionality required to fulfill the requests. The
    // type is interface{} with range: 0..4294967295.
    CSipStatsServerNotImplementedIns interface{}

    // This object reflects the total number of Not Implemented  (501) responses
    // sent by the user agent since system startup. Outbound Not Implemented
    // responses indicate that this system does not support the functionality
    // required to fulfill the  requests. The type is interface{} with range:
    // 0..4294967295.
    CSipStatsServerNotImplementedOuts interface{}

    // This object reflects the total number of Bad Gateway (502)  responses
    // received by the user agent since system startup. The type is interface{}
    // with range: 0..4294967295.
    CSipStatsServerBadGatewayIns interface{}

    // This object reflects the total number of Bad Gateway (502)  responses sent
    // by the user agent since system startup. The type is interface{} with range:
    // 0..4294967295.
    CSipStatsServerBadGatewayOuts interface{}

    // This object reflects the total number of Service Unavailable  (503)
    // responses received by the user agent since system startup. Inbound Service
    // Unavailable responses indicate that the server servicing this system's
    // request is temporarily unavailable to handle the request. The type is
    // interface{} with range: 0..4294967295.
    CSipStatsServerServiceUnavailIns interface{}

    // This object reflects the total number of Service Unavailable  (503)
    // responses sent by the user agent since system startup. Outbound Service
    // Unavailable responses indicate that this system is temporarily unable to
    // handle received requests. The type is interface{} with range:
    // 0..4294967295.
    CSipStatsServerServiceUnavailOuts interface{}

    // This object reflects the total number of Gateway Time-out  (504) responses
    // received by the user agent since system startup. Inbound Gateway Time-out
    // responses indicate that the server attempting to complete this system's
    // request did not receive a timely response from yet another system it was
    // accessing to complete the request. The type is interface{} with range:
    // 0..4294967295.
    CSipStatsServerGatewayTimeoutIns interface{}

    // This object reflects the total number of Gateway Time-out  (504) responses
    // sent by the user agent since system startup. Outbound Gateway Time-out
    // responses indicate that this system did not receive a timely response from
    // the system it had accessed to assist in completing a received request. The
    // type is interface{} with range: 0..4294967295.
    CSipStatsServerGatewayTimeoutOuts interface{}

    // This object reflects the total number of SIP Version Not  Supported (505)
    // responses received by the user agent since system startup. Inbound SIP
    // Version Not Supported responses indicate that  the server does not support,
    // or refuses to support, the SIP  protocol version that was used in the
    // request message. The type is interface{} with range: 0..4294967295.
    CSipStatsServerBadSipVersionIns interface{}

    // This object reflects the total number of SIP Version Not  Supported (505)
    // responses sent by the user agent since system startup. Outbound SIP Version
    // Not Supported responses indicate that  this system does not support, or
    // refuses to support, the SIP  protocol version used in received requests.
    // The type is interface{} with range: 0..4294967295.
    CSipStatsServerBadSipVersionOuts interface{}

    // This object reflects the total number of Precondition  Failure (580)
    // responses received by the user agent since system startup. This response is
    // returned by a UAS if it is unable to perform the mandatory preconditions
    // for the session. The type is interface{} with range: 0..4294967295.
    CSipStatsServerPrecondFailureIns interface{}

    // This object reflects the total number of Precondition  Failure (580)
    // responses sent by the user agent since system startup. This response is
    // returned by a UAS if it is unable to perform the mandatory preconditions
    // for the session. The type is interface{} with range: 0..4294967295.
    CSipStatsServerPrecondFailureOuts interface{}
}

func (cSipStatsErrServer *CISCOSIPUAMIB_CSipStatsErrServer) GetEntityData() *types.CommonEntityData {
    cSipStatsErrServer.EntityData.YFilter = cSipStatsErrServer.YFilter
    cSipStatsErrServer.EntityData.YangName = "cSipStatsErrServer"
    cSipStatsErrServer.EntityData.BundleName = "cisco_ios_xe"
    cSipStatsErrServer.EntityData.ParentYangName = "CISCO-SIP-UA-MIB"
    cSipStatsErrServer.EntityData.SegmentPath = "cSipStatsErrServer"
    cSipStatsErrServer.EntityData.AbsolutePath = "CISCO-SIP-UA-MIB:CISCO-SIP-UA-MIB/" + cSipStatsErrServer.EntityData.SegmentPath
    cSipStatsErrServer.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cSipStatsErrServer.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cSipStatsErrServer.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cSipStatsErrServer.EntityData.Children = types.NewOrderedMap()
    cSipStatsErrServer.EntityData.Leafs = types.NewOrderedMap()
    cSipStatsErrServer.EntityData.Leafs.Append("cSipStatsServerIntErrorIns", types.YLeaf{"CSipStatsServerIntErrorIns", cSipStatsErrServer.CSipStatsServerIntErrorIns})
    cSipStatsErrServer.EntityData.Leafs.Append("cSipStatsServerIntErrorOuts", types.YLeaf{"CSipStatsServerIntErrorOuts", cSipStatsErrServer.CSipStatsServerIntErrorOuts})
    cSipStatsErrServer.EntityData.Leafs.Append("cSipStatsServerNotImplementedIns", types.YLeaf{"CSipStatsServerNotImplementedIns", cSipStatsErrServer.CSipStatsServerNotImplementedIns})
    cSipStatsErrServer.EntityData.Leafs.Append("cSipStatsServerNotImplementedOuts", types.YLeaf{"CSipStatsServerNotImplementedOuts", cSipStatsErrServer.CSipStatsServerNotImplementedOuts})
    cSipStatsErrServer.EntityData.Leafs.Append("cSipStatsServerBadGatewayIns", types.YLeaf{"CSipStatsServerBadGatewayIns", cSipStatsErrServer.CSipStatsServerBadGatewayIns})
    cSipStatsErrServer.EntityData.Leafs.Append("cSipStatsServerBadGatewayOuts", types.YLeaf{"CSipStatsServerBadGatewayOuts", cSipStatsErrServer.CSipStatsServerBadGatewayOuts})
    cSipStatsErrServer.EntityData.Leafs.Append("cSipStatsServerServiceUnavailIns", types.YLeaf{"CSipStatsServerServiceUnavailIns", cSipStatsErrServer.CSipStatsServerServiceUnavailIns})
    cSipStatsErrServer.EntityData.Leafs.Append("cSipStatsServerServiceUnavailOuts", types.YLeaf{"CSipStatsServerServiceUnavailOuts", cSipStatsErrServer.CSipStatsServerServiceUnavailOuts})
    cSipStatsErrServer.EntityData.Leafs.Append("cSipStatsServerGatewayTimeoutIns", types.YLeaf{"CSipStatsServerGatewayTimeoutIns", cSipStatsErrServer.CSipStatsServerGatewayTimeoutIns})
    cSipStatsErrServer.EntityData.Leafs.Append("cSipStatsServerGatewayTimeoutOuts", types.YLeaf{"CSipStatsServerGatewayTimeoutOuts", cSipStatsErrServer.CSipStatsServerGatewayTimeoutOuts})
    cSipStatsErrServer.EntityData.Leafs.Append("cSipStatsServerBadSipVersionIns", types.YLeaf{"CSipStatsServerBadSipVersionIns", cSipStatsErrServer.CSipStatsServerBadSipVersionIns})
    cSipStatsErrServer.EntityData.Leafs.Append("cSipStatsServerBadSipVersionOuts", types.YLeaf{"CSipStatsServerBadSipVersionOuts", cSipStatsErrServer.CSipStatsServerBadSipVersionOuts})
    cSipStatsErrServer.EntityData.Leafs.Append("cSipStatsServerPrecondFailureIns", types.YLeaf{"CSipStatsServerPrecondFailureIns", cSipStatsErrServer.CSipStatsServerPrecondFailureIns})
    cSipStatsErrServer.EntityData.Leafs.Append("cSipStatsServerPrecondFailureOuts", types.YLeaf{"CSipStatsServerPrecondFailureOuts", cSipStatsErrServer.CSipStatsServerPrecondFailureOuts})

    cSipStatsErrServer.EntityData.YListKeys = []string {}

    return &(cSipStatsErrServer.EntityData)
}

// CISCOSIPUAMIB_CSipStatsGlobalFail
type CISCOSIPUAMIB_CSipStatsGlobalFail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object reflects the total number of Busy Everywhere (600) responses
    // received by the user agent since system startup. Inbound Busy Everywhere
    // responses indicate that the  called party's end system was contacted
    // successfully but the called party is busy and does not wish to take the
    // call at this time. The type is interface{} with range: 0..4294967295.
    CSipStatsGlobalBusyEverywhereIns interface{}

    // This object reflects the total number of Busy Everywhere (600) responses
    // sent by the user agent since system startup. Outbound Busy Everywhere
    // responses indicate that  this system has successfully contacted a called
    // party's end system and the called party does not wish to take the call at
    // this time. The type is interface{} with range: 0..4294967295.
    CSipStatsGlobalBusyEverywhereOuts interface{}

    // This object reflects the total number of Decline (603) responses received
    // by the user agent since system startup. Decline responses indicate that the
    // called party's end  system was contacted successfully but the called party 
    // explicitly does not wish to, or cannot, participate. The type is
    // interface{} with range: 0..4294967295.
    CSipStatsGlobalDeclineIns interface{}

    // This object reflects the total number of Decline (603) responses sent by
    // the user agent since system startup. Outbound Decline responses indicate
    // that this system has successfully contacted a called party's end system and
    // the called party explicitly does not wish to, or cannot, participate. The
    // type is interface{} with range: 0..4294967295.
    CSipStatsGlobalDeclineOuts interface{}

    // This object reflects the total number of Does Not Exist Anywhere (604)
    // responses received by the user agent since system startup. Inbound Does Not
    // Exist Anywhere responses indicate that the server handling this system's
    // request has authoritative information that the called party indicated in
    // the To request field does not exist anywhere. The type is interface{} with
    // range: 0..4294967295.
    CSipStatsGlobalNotAnywhereIns interface{}

    // This object reflects the total number of Does Not Exist Anywhere (604)
    // responses sent by the user agent since system startup. Outbound Does Not
    // Exist Anywhere responses indicate that this system has authoritative
    // information that the called party in the To field of received requests does
    // not exist anywhere. The type is interface{} with range: 0..4294967295.
    CSipStatsGlobalNotAnywhereOuts interface{}

    // This object reflects the total number of Not Acceptable (606) responses
    // received by the user agent since system startup. Inbound Not Acceptable
    // responses indicate that the called party's end system was contacted
    // successfully but some aspect of the session description is not acceptable.
    // The type is interface{} with range: 0..4294967295.
    CSipStatsGlobalNotAcceptableIns interface{}

    // This object reflects the total number of Not Acceptable (606) responses
    // sent by the user agent since system startup. Outbound Not Acceptable
    // responses indicate that the called party wishes to communicate, but cannot
    // adequately support the session described in the request. The type is
    // interface{} with range: 0..4294967295.
    CSipStatsGlobalNotAcceptableOuts interface{}
}

func (cSipStatsGlobalFail *CISCOSIPUAMIB_CSipStatsGlobalFail) GetEntityData() *types.CommonEntityData {
    cSipStatsGlobalFail.EntityData.YFilter = cSipStatsGlobalFail.YFilter
    cSipStatsGlobalFail.EntityData.YangName = "cSipStatsGlobalFail"
    cSipStatsGlobalFail.EntityData.BundleName = "cisco_ios_xe"
    cSipStatsGlobalFail.EntityData.ParentYangName = "CISCO-SIP-UA-MIB"
    cSipStatsGlobalFail.EntityData.SegmentPath = "cSipStatsGlobalFail"
    cSipStatsGlobalFail.EntityData.AbsolutePath = "CISCO-SIP-UA-MIB:CISCO-SIP-UA-MIB/" + cSipStatsGlobalFail.EntityData.SegmentPath
    cSipStatsGlobalFail.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cSipStatsGlobalFail.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cSipStatsGlobalFail.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cSipStatsGlobalFail.EntityData.Children = types.NewOrderedMap()
    cSipStatsGlobalFail.EntityData.Leafs = types.NewOrderedMap()
    cSipStatsGlobalFail.EntityData.Leafs.Append("cSipStatsGlobalBusyEverywhereIns", types.YLeaf{"CSipStatsGlobalBusyEverywhereIns", cSipStatsGlobalFail.CSipStatsGlobalBusyEverywhereIns})
    cSipStatsGlobalFail.EntityData.Leafs.Append("cSipStatsGlobalBusyEverywhereOuts", types.YLeaf{"CSipStatsGlobalBusyEverywhereOuts", cSipStatsGlobalFail.CSipStatsGlobalBusyEverywhereOuts})
    cSipStatsGlobalFail.EntityData.Leafs.Append("cSipStatsGlobalDeclineIns", types.YLeaf{"CSipStatsGlobalDeclineIns", cSipStatsGlobalFail.CSipStatsGlobalDeclineIns})
    cSipStatsGlobalFail.EntityData.Leafs.Append("cSipStatsGlobalDeclineOuts", types.YLeaf{"CSipStatsGlobalDeclineOuts", cSipStatsGlobalFail.CSipStatsGlobalDeclineOuts})
    cSipStatsGlobalFail.EntityData.Leafs.Append("cSipStatsGlobalNotAnywhereIns", types.YLeaf{"CSipStatsGlobalNotAnywhereIns", cSipStatsGlobalFail.CSipStatsGlobalNotAnywhereIns})
    cSipStatsGlobalFail.EntityData.Leafs.Append("cSipStatsGlobalNotAnywhereOuts", types.YLeaf{"CSipStatsGlobalNotAnywhereOuts", cSipStatsGlobalFail.CSipStatsGlobalNotAnywhereOuts})
    cSipStatsGlobalFail.EntityData.Leafs.Append("cSipStatsGlobalNotAcceptableIns", types.YLeaf{"CSipStatsGlobalNotAcceptableIns", cSipStatsGlobalFail.CSipStatsGlobalNotAcceptableIns})
    cSipStatsGlobalFail.EntityData.Leafs.Append("cSipStatsGlobalNotAcceptableOuts", types.YLeaf{"CSipStatsGlobalNotAcceptableOuts", cSipStatsGlobalFail.CSipStatsGlobalNotAcceptableOuts})

    cSipStatsGlobalFail.EntityData.YListKeys = []string {}

    return &(cSipStatsGlobalFail.EntityData)
}

// CISCOSIPUAMIB_CSipStatsTraffic
type CISCOSIPUAMIB_CSipStatsTraffic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object reflects the total number of INVITE requests  received by the
    // user agent since system startup. The type is interface{} with range:
    // 0..4294967295.
    CSipStatsTrafficInviteIns interface{}

    // This object reflects the total number of INVITE requests sent by the user
    // agent. The type is interface{} with range: 0..4294967295.
    CSipStatsTrafficInviteOuts interface{}

    // This object reflects the total number of ACK requests  received by the user
    // agent since system startup. The type is interface{} with range:
    // 0..4294967295.
    CSipStatsTrafficAckIns interface{}

    // This object reflects the total number of ACK requests sent by the user
    // agent. The type is interface{} with range: 0..4294967295.
    CSipStatsTrafficAckOuts interface{}

    // This object reflects the total number of BYE requests  received by the user
    // agent since system startup. The type is interface{} with range:
    // 0..4294967295.
    CSipStatsTrafficByeIns interface{}

    // This object reflects the total number of BYE requests sent by the user
    // agent. The type is interface{} with range: 0..4294967295.
    CSipStatsTrafficByeOuts interface{}

    // This object reflects the total number of CANCEL requests  received by the
    // user agent since system startup. The type is interface{} with range:
    // 0..4294967295.
    CSipStatsTrafficCancelIns interface{}

    // This object reflects the total number of CANCEL requests sent by the user
    // agent. The type is interface{} with range: 0..4294967295.
    CSipStatsTrafficCancelOuts interface{}

    // This object reflects the total number of OPTIONS requests  received by the
    // user agent since system startup. The type is interface{} with range:
    // 0..4294967295.
    CSipStatsTrafficOptionsIns interface{}

    // This object reflects the total number of OPTIONS requests sent by the user
    // agent. The type is interface{} with range: 0..4294967295.
    CSipStatsTrafficOptionsOuts interface{}

    // This object reflects the total number of REGISTER requests  received by the
    // user agent since system startup. The type is interface{} with range:
    // 0..4294967295.
    CSipStatsTrafficRegisterIns interface{}

    // This object reflects the total number of REGISTER requests  sent by the
    // user agent since system startup. The type is interface{} with range:
    // 0..4294967295.
    CSipStatsTrafficRegisterOuts interface{}

    // This object reflects the total number of COndition MET  requests received
    // by the user agent since system startup. The type is interface{} with range:
    // 0..4294967295.
    CSipStatsTrafficCometIns interface{}

    // This object reflects the total number of COndition MET  requests sent by
    // the user agent since system startup. The type is interface{} with range:
    // 0..4294967295.
    CSipStatsTrafficCometOuts interface{}

    // This object reflects the total number of PRovisonal  ACKnowledgement
    // requests received by the user agent since system startup. The type is
    // interface{} with range: 0..4294967295.
    CSipStatsTrafficPrackIns interface{}

    // This object reflects the total number of PRovisonal  ACKnowledgement
    // requests sent by the user agent since system startup. The type is
    // interface{} with range: 0..4294967295.
    CSipStatsTrafficPrackOuts interface{}

    // This object reflects the total number of Refer requests received by the
    // user agent since system startup. The type is interface{} with range:
    // 0..4294967295.
    CSipStatsTrafficReferIns interface{}

    // This object reflects the total number of Refer requests sent by the user
    // agent since system startup. The type is interface{} with range:
    // 0..4294967295.
    CSipStatsTrafficReferOuts interface{}

    // This object reflects the total number of Notify  requests received by the
    // user agent since system startup. The type is interface{} with range:
    // 0..4294967295.
    CSipStatsTrafficNotifyIns interface{}

    // This object reflects the total number of Notify  requests sent by the user
    // agent since system startup. The type is interface{} with range:
    // 0..4294967295.
    CSipStatsTrafficNotifyOuts interface{}

    // This object reflects the total number of Info  requests received by the
    // user agent since system startup. The type is interface{} with range:
    // 0..4294967295.
    CSipStatsTrafficInfoIns interface{}

    // This object reflects the total number of Info  requests sent by the user
    // agent since system startup. The type is interface{} with range:
    // 0..4294967295.
    CSipStatsTrafficInfoOuts interface{}

    // This object reflects the total number of Subscribe requests received by the
    // user agent since system startup. The type is interface{} with range:
    // 0..4294967295.
    CSipStatsTrafficSubscribeIns interface{}

    // This object reflects the total number of Subscribe requests sent by the
    // user agent since system startup. The type is interface{} with range:
    // 0..4294967295.
    CSipStatsTrafficSubscribeOuts interface{}

    // This object reflects the total number of Update requests received by the
    // user agent since system startup. The type is interface{} with range:
    // 0..4294967295.
    CSipStatsTrafficUpdateIns interface{}

    // This object reflects the total number of Update requests sent by the user
    // agent since system startup. The type is interface{} with range:
    // 0..4294967295.
    CSipStatsTrafficUpdateOuts interface{}
}

func (cSipStatsTraffic *CISCOSIPUAMIB_CSipStatsTraffic) GetEntityData() *types.CommonEntityData {
    cSipStatsTraffic.EntityData.YFilter = cSipStatsTraffic.YFilter
    cSipStatsTraffic.EntityData.YangName = "cSipStatsTraffic"
    cSipStatsTraffic.EntityData.BundleName = "cisco_ios_xe"
    cSipStatsTraffic.EntityData.ParentYangName = "CISCO-SIP-UA-MIB"
    cSipStatsTraffic.EntityData.SegmentPath = "cSipStatsTraffic"
    cSipStatsTraffic.EntityData.AbsolutePath = "CISCO-SIP-UA-MIB:CISCO-SIP-UA-MIB/" + cSipStatsTraffic.EntityData.SegmentPath
    cSipStatsTraffic.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cSipStatsTraffic.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cSipStatsTraffic.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cSipStatsTraffic.EntityData.Children = types.NewOrderedMap()
    cSipStatsTraffic.EntityData.Leafs = types.NewOrderedMap()
    cSipStatsTraffic.EntityData.Leafs.Append("cSipStatsTrafficInviteIns", types.YLeaf{"CSipStatsTrafficInviteIns", cSipStatsTraffic.CSipStatsTrafficInviteIns})
    cSipStatsTraffic.EntityData.Leafs.Append("cSipStatsTrafficInviteOuts", types.YLeaf{"CSipStatsTrafficInviteOuts", cSipStatsTraffic.CSipStatsTrafficInviteOuts})
    cSipStatsTraffic.EntityData.Leafs.Append("cSipStatsTrafficAckIns", types.YLeaf{"CSipStatsTrafficAckIns", cSipStatsTraffic.CSipStatsTrafficAckIns})
    cSipStatsTraffic.EntityData.Leafs.Append("cSipStatsTrafficAckOuts", types.YLeaf{"CSipStatsTrafficAckOuts", cSipStatsTraffic.CSipStatsTrafficAckOuts})
    cSipStatsTraffic.EntityData.Leafs.Append("cSipStatsTrafficByeIns", types.YLeaf{"CSipStatsTrafficByeIns", cSipStatsTraffic.CSipStatsTrafficByeIns})
    cSipStatsTraffic.EntityData.Leafs.Append("cSipStatsTrafficByeOuts", types.YLeaf{"CSipStatsTrafficByeOuts", cSipStatsTraffic.CSipStatsTrafficByeOuts})
    cSipStatsTraffic.EntityData.Leafs.Append("cSipStatsTrafficCancelIns", types.YLeaf{"CSipStatsTrafficCancelIns", cSipStatsTraffic.CSipStatsTrafficCancelIns})
    cSipStatsTraffic.EntityData.Leafs.Append("cSipStatsTrafficCancelOuts", types.YLeaf{"CSipStatsTrafficCancelOuts", cSipStatsTraffic.CSipStatsTrafficCancelOuts})
    cSipStatsTraffic.EntityData.Leafs.Append("cSipStatsTrafficOptionsIns", types.YLeaf{"CSipStatsTrafficOptionsIns", cSipStatsTraffic.CSipStatsTrafficOptionsIns})
    cSipStatsTraffic.EntityData.Leafs.Append("cSipStatsTrafficOptionsOuts", types.YLeaf{"CSipStatsTrafficOptionsOuts", cSipStatsTraffic.CSipStatsTrafficOptionsOuts})
    cSipStatsTraffic.EntityData.Leafs.Append("cSipStatsTrafficRegisterIns", types.YLeaf{"CSipStatsTrafficRegisterIns", cSipStatsTraffic.CSipStatsTrafficRegisterIns})
    cSipStatsTraffic.EntityData.Leafs.Append("cSipStatsTrafficRegisterOuts", types.YLeaf{"CSipStatsTrafficRegisterOuts", cSipStatsTraffic.CSipStatsTrafficRegisterOuts})
    cSipStatsTraffic.EntityData.Leafs.Append("cSipStatsTrafficCometIns", types.YLeaf{"CSipStatsTrafficCometIns", cSipStatsTraffic.CSipStatsTrafficCometIns})
    cSipStatsTraffic.EntityData.Leafs.Append("cSipStatsTrafficCometOuts", types.YLeaf{"CSipStatsTrafficCometOuts", cSipStatsTraffic.CSipStatsTrafficCometOuts})
    cSipStatsTraffic.EntityData.Leafs.Append("cSipStatsTrafficPrackIns", types.YLeaf{"CSipStatsTrafficPrackIns", cSipStatsTraffic.CSipStatsTrafficPrackIns})
    cSipStatsTraffic.EntityData.Leafs.Append("cSipStatsTrafficPrackOuts", types.YLeaf{"CSipStatsTrafficPrackOuts", cSipStatsTraffic.CSipStatsTrafficPrackOuts})
    cSipStatsTraffic.EntityData.Leafs.Append("cSipStatsTrafficReferIns", types.YLeaf{"CSipStatsTrafficReferIns", cSipStatsTraffic.CSipStatsTrafficReferIns})
    cSipStatsTraffic.EntityData.Leafs.Append("cSipStatsTrafficReferOuts", types.YLeaf{"CSipStatsTrafficReferOuts", cSipStatsTraffic.CSipStatsTrafficReferOuts})
    cSipStatsTraffic.EntityData.Leafs.Append("cSipStatsTrafficNotifyIns", types.YLeaf{"CSipStatsTrafficNotifyIns", cSipStatsTraffic.CSipStatsTrafficNotifyIns})
    cSipStatsTraffic.EntityData.Leafs.Append("cSipStatsTrafficNotifyOuts", types.YLeaf{"CSipStatsTrafficNotifyOuts", cSipStatsTraffic.CSipStatsTrafficNotifyOuts})
    cSipStatsTraffic.EntityData.Leafs.Append("cSipStatsTrafficInfoIns", types.YLeaf{"CSipStatsTrafficInfoIns", cSipStatsTraffic.CSipStatsTrafficInfoIns})
    cSipStatsTraffic.EntityData.Leafs.Append("cSipStatsTrafficInfoOuts", types.YLeaf{"CSipStatsTrafficInfoOuts", cSipStatsTraffic.CSipStatsTrafficInfoOuts})
    cSipStatsTraffic.EntityData.Leafs.Append("cSipStatsTrafficSubscribeIns", types.YLeaf{"CSipStatsTrafficSubscribeIns", cSipStatsTraffic.CSipStatsTrafficSubscribeIns})
    cSipStatsTraffic.EntityData.Leafs.Append("cSipStatsTrafficSubscribeOuts", types.YLeaf{"CSipStatsTrafficSubscribeOuts", cSipStatsTraffic.CSipStatsTrafficSubscribeOuts})
    cSipStatsTraffic.EntityData.Leafs.Append("cSipStatsTrafficUpdateIns", types.YLeaf{"CSipStatsTrafficUpdateIns", cSipStatsTraffic.CSipStatsTrafficUpdateIns})
    cSipStatsTraffic.EntityData.Leafs.Append("cSipStatsTrafficUpdateOuts", types.YLeaf{"CSipStatsTrafficUpdateOuts", cSipStatsTraffic.CSipStatsTrafficUpdateOuts})

    cSipStatsTraffic.EntityData.YListKeys = []string {}

    return &(cSipStatsTraffic.EntityData)
}

// CISCOSIPUAMIB_CSipStatsRetry
type CISCOSIPUAMIB_CSipStatsRetry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object reflects the total number of INVITE retries that  have been
    // sent by the user agent since system startup.   If the number of 'first 
    // attempt' INVITES is of interest, subtract the value of this  object from
    // cSipStatsTrafficInviteOut. The type is interface{} with range:
    // 0..4294967295.
    CSipStatsRetryInvites interface{}

    // This object reflects the total number of BYE retries that have been sent by
    // the user agent since system startup. The type is interface{} with range:
    // 0..4294967295.
    CSipStatsRetryByes interface{}

    // This object reflects the total number of CANCEL retries that  have been
    // sent by the user agent since system startup. The type is interface{} with
    // range: 0..4294967295.
    CSipStatsRetryCancels interface{}

    // This object reflects the total number of REGISTER retries that have been
    // sent by the user agent since system startup. The type is interface{} with
    // range: 0..4294967295.
    CSipStatsRetryRegisters interface{}

    // This object reflects the total number of Response (while  expecting a ACK)
    // retries that have been sent by the user agent. The type is interface{} with
    // range: 0..4294967295.
    CSipStatsRetryResponses interface{}

    // This object reflects the total number of PRovisional ACKnowledgement
    // retries sent by the user agent since system startup. The type is
    // interface{} with range: 0..4294967295.
    CSipStatsRetryPracks interface{}

    // This object reflects the total number of COndition MET retries sent by the
    // user agent since system startup. The type is interface{} with range:
    // 0..4294967295.
    CSipStatsRetryComets interface{}

    // This object reflects the total number of Reliable 1xx Response retries sent
    // by the user agent since system startup. The type is interface{} with range:
    // 0..4294967295.
    CSipStatsRetryReliable1xxRsps interface{}

    // This object reflects the total number of Notify  retries that have been
    // sent by the user agent since system startup. The type is interface{} with
    // range: 0..4294967295.
    CSipStatsRetryNotifys interface{}

    // This object reflects the total number of Refer retries that have been sent
    // by the user agent since system startup. The type is interface{} with range:
    // 0..4294967295.
    CSipStatsRetryRefers interface{}

    // This object reflects the total number of Info retries that have been sent
    // by the user agent since system startup. The type is interface{} with range:
    // 0..4294967295.
    CSipStatsRetryInfo interface{}

    // This object reflects the total number of Subscribe retries that have been
    // sent by the user agent since system startup. The type is interface{} with
    // range: 0..4294967295.
    CSipStatsRetrySubscribe interface{}
}

func (cSipStatsRetry *CISCOSIPUAMIB_CSipStatsRetry) GetEntityData() *types.CommonEntityData {
    cSipStatsRetry.EntityData.YFilter = cSipStatsRetry.YFilter
    cSipStatsRetry.EntityData.YangName = "cSipStatsRetry"
    cSipStatsRetry.EntityData.BundleName = "cisco_ios_xe"
    cSipStatsRetry.EntityData.ParentYangName = "CISCO-SIP-UA-MIB"
    cSipStatsRetry.EntityData.SegmentPath = "cSipStatsRetry"
    cSipStatsRetry.EntityData.AbsolutePath = "CISCO-SIP-UA-MIB:CISCO-SIP-UA-MIB/" + cSipStatsRetry.EntityData.SegmentPath
    cSipStatsRetry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cSipStatsRetry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cSipStatsRetry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cSipStatsRetry.EntityData.Children = types.NewOrderedMap()
    cSipStatsRetry.EntityData.Leafs = types.NewOrderedMap()
    cSipStatsRetry.EntityData.Leafs.Append("cSipStatsRetryInvites", types.YLeaf{"CSipStatsRetryInvites", cSipStatsRetry.CSipStatsRetryInvites})
    cSipStatsRetry.EntityData.Leafs.Append("cSipStatsRetryByes", types.YLeaf{"CSipStatsRetryByes", cSipStatsRetry.CSipStatsRetryByes})
    cSipStatsRetry.EntityData.Leafs.Append("cSipStatsRetryCancels", types.YLeaf{"CSipStatsRetryCancels", cSipStatsRetry.CSipStatsRetryCancels})
    cSipStatsRetry.EntityData.Leafs.Append("cSipStatsRetryRegisters", types.YLeaf{"CSipStatsRetryRegisters", cSipStatsRetry.CSipStatsRetryRegisters})
    cSipStatsRetry.EntityData.Leafs.Append("cSipStatsRetryResponses", types.YLeaf{"CSipStatsRetryResponses", cSipStatsRetry.CSipStatsRetryResponses})
    cSipStatsRetry.EntityData.Leafs.Append("cSipStatsRetryPracks", types.YLeaf{"CSipStatsRetryPracks", cSipStatsRetry.CSipStatsRetryPracks})
    cSipStatsRetry.EntityData.Leafs.Append("cSipStatsRetryComets", types.YLeaf{"CSipStatsRetryComets", cSipStatsRetry.CSipStatsRetryComets})
    cSipStatsRetry.EntityData.Leafs.Append("cSipStatsRetryReliable1xxRsps", types.YLeaf{"CSipStatsRetryReliable1xxRsps", cSipStatsRetry.CSipStatsRetryReliable1xxRsps})
    cSipStatsRetry.EntityData.Leafs.Append("cSipStatsRetryNotifys", types.YLeaf{"CSipStatsRetryNotifys", cSipStatsRetry.CSipStatsRetryNotifys})
    cSipStatsRetry.EntityData.Leafs.Append("cSipStatsRetryRefers", types.YLeaf{"CSipStatsRetryRefers", cSipStatsRetry.CSipStatsRetryRefers})
    cSipStatsRetry.EntityData.Leafs.Append("cSipStatsRetryInfo", types.YLeaf{"CSipStatsRetryInfo", cSipStatsRetry.CSipStatsRetryInfo})
    cSipStatsRetry.EntityData.Leafs.Append("cSipStatsRetrySubscribe", types.YLeaf{"CSipStatsRetrySubscribe", cSipStatsRetry.CSipStatsRetrySubscribe})

    cSipStatsRetry.EntityData.YListKeys = []string {}

    return &(cSipStatsRetry.EntityData)
}

// CISCOSIPUAMIB_CSipStatsMisc
type CISCOSIPUAMIB_CSipStatsMisc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object reflects the total number of incoming Redirect  (3xx) response
    // messages that have been mapped to Client  Error (4xx) response messages.
    // The type is interface{} with range: 0..4294967295.
    CSipStatsMisc3xxMappedTo4xxRsps interface{}
}

func (cSipStatsMisc *CISCOSIPUAMIB_CSipStatsMisc) GetEntityData() *types.CommonEntityData {
    cSipStatsMisc.EntityData.YFilter = cSipStatsMisc.YFilter
    cSipStatsMisc.EntityData.YangName = "cSipStatsMisc"
    cSipStatsMisc.EntityData.BundleName = "cisco_ios_xe"
    cSipStatsMisc.EntityData.ParentYangName = "CISCO-SIP-UA-MIB"
    cSipStatsMisc.EntityData.SegmentPath = "cSipStatsMisc"
    cSipStatsMisc.EntityData.AbsolutePath = "CISCO-SIP-UA-MIB:CISCO-SIP-UA-MIB/" + cSipStatsMisc.EntityData.SegmentPath
    cSipStatsMisc.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cSipStatsMisc.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cSipStatsMisc.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cSipStatsMisc.EntityData.Children = types.NewOrderedMap()
    cSipStatsMisc.EntityData.Leafs = types.NewOrderedMap()
    cSipStatsMisc.EntityData.Leafs.Append("cSipStatsMisc3xxMappedTo4xxRsps", types.YLeaf{"CSipStatsMisc3xxMappedTo4xxRsps", cSipStatsMisc.CSipStatsMisc3xxMappedTo4xxRsps})

    cSipStatsMisc.EntityData.YListKeys = []string {}

    return &(cSipStatsMisc.EntityData)
}

// CISCOSIPUAMIB_CSipStatsConnection
type CISCOSIPUAMIB_CSipStatsConnection struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object reflects the total number of TCP send failures since system
    // startup. The type is interface{} with range: 0..4294967295.
    CSipStatsConnTCPSendFailures interface{}

    // This object reflects the total number of UDP send failures since system
    // startup. The type is interface{} with range: 0..4294967295.
    CSipStatsConnUDPSendFailures interface{}

    // This object reflects the total number of Remote Closures  for TCP since
    // system startup. The type is interface{} with range: 0..4294967295.
    CSipStatsConnTCPRemoteClosures interface{}

    // This object reflects the total number of connection create failures for UDP
    // since system startup. The type is interface{} with range: 0..4294967295.
    CSipStatsConnUDPCreateFailures interface{}

    // This object reflects the total number of connection create failures for TCP
    // since system startup. The type is interface{} with range: 0..4294967295.
    CSipStatsConnTCPCreateFailures interface{}

    // This object reflects the total number of UDP connections that  timed out
    // due to inactivity since system startup. The type is interface{} with range:
    // 0..4294967295.
    CSipStatsConnUDPInactiveTimeouts interface{}

    // This object reflects the total number of TCP connections that timed out due
    // to inactivity since system startup. The type is interface{} with range:
    // 0..4294967295.
    CSipStatsConnTCPInactiveTimeouts interface{}

    // This object reflects the total number of active UDP connections since
    // system startup. The type is interface{} with range: 0..4294967295.
    CSipStatsActiveUDPConnections interface{}

    // This object reflects the total number of active TCP connections since
    // system startup. The type is interface{} with range: 0..4294967295.
    CSipStatsActiveTCPConnections interface{}
}

func (cSipStatsConnection *CISCOSIPUAMIB_CSipStatsConnection) GetEntityData() *types.CommonEntityData {
    cSipStatsConnection.EntityData.YFilter = cSipStatsConnection.YFilter
    cSipStatsConnection.EntityData.YangName = "cSipStatsConnection"
    cSipStatsConnection.EntityData.BundleName = "cisco_ios_xe"
    cSipStatsConnection.EntityData.ParentYangName = "CISCO-SIP-UA-MIB"
    cSipStatsConnection.EntityData.SegmentPath = "cSipStatsConnection"
    cSipStatsConnection.EntityData.AbsolutePath = "CISCO-SIP-UA-MIB:CISCO-SIP-UA-MIB/" + cSipStatsConnection.EntityData.SegmentPath
    cSipStatsConnection.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cSipStatsConnection.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cSipStatsConnection.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cSipStatsConnection.EntityData.Children = types.NewOrderedMap()
    cSipStatsConnection.EntityData.Leafs = types.NewOrderedMap()
    cSipStatsConnection.EntityData.Leafs.Append("cSipStatsConnTCPSendFailures", types.YLeaf{"CSipStatsConnTCPSendFailures", cSipStatsConnection.CSipStatsConnTCPSendFailures})
    cSipStatsConnection.EntityData.Leafs.Append("cSipStatsConnUDPSendFailures", types.YLeaf{"CSipStatsConnUDPSendFailures", cSipStatsConnection.CSipStatsConnUDPSendFailures})
    cSipStatsConnection.EntityData.Leafs.Append("cSipStatsConnTCPRemoteClosures", types.YLeaf{"CSipStatsConnTCPRemoteClosures", cSipStatsConnection.CSipStatsConnTCPRemoteClosures})
    cSipStatsConnection.EntityData.Leafs.Append("cSipStatsConnUDPCreateFailures", types.YLeaf{"CSipStatsConnUDPCreateFailures", cSipStatsConnection.CSipStatsConnUDPCreateFailures})
    cSipStatsConnection.EntityData.Leafs.Append("cSipStatsConnTCPCreateFailures", types.YLeaf{"CSipStatsConnTCPCreateFailures", cSipStatsConnection.CSipStatsConnTCPCreateFailures})
    cSipStatsConnection.EntityData.Leafs.Append("cSipStatsConnUDPInactiveTimeouts", types.YLeaf{"CSipStatsConnUDPInactiveTimeouts", cSipStatsConnection.CSipStatsConnUDPInactiveTimeouts})
    cSipStatsConnection.EntityData.Leafs.Append("cSipStatsConnTCPInactiveTimeouts", types.YLeaf{"CSipStatsConnTCPInactiveTimeouts", cSipStatsConnection.CSipStatsConnTCPInactiveTimeouts})
    cSipStatsConnection.EntityData.Leafs.Append("cSipStatsActiveUDPConnections", types.YLeaf{"CSipStatsActiveUDPConnections", cSipStatsConnection.CSipStatsActiveUDPConnections})
    cSipStatsConnection.EntityData.Leafs.Append("cSipStatsActiveTCPConnections", types.YLeaf{"CSipStatsActiveTCPConnections", cSipStatsConnection.CSipStatsActiveTCPConnections})

    cSipStatsConnection.EntityData.YListKeys = []string {}

    return &(cSipStatsConnection.EntityData)
}

// CISCOSIPUAMIB_CSipCfgEarlyMediaTable
// This table contains configuration for Early
// Media Cut Through.  The configuration controls
// how the SIP user agent will process 1xx
// (Provisional) SIP response messages that contain 
// Session Definition Protocol (SDP) payloads.
type CISCOSIPUAMIB_CSipCfgEarlyMediaTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A row in the cSipCfgEarlyMediaTable. A row is accessible with a Provisional
    // (1xx) status code value (eg, 180) and provides an object for the enabling
    // or disabling of the Early Media Cut Through functionality. The type is
    // slice of CISCOSIPUAMIB_CSipCfgEarlyMediaTable_CSipCfgEarlyMediaEntry.
    CSipCfgEarlyMediaEntry []*CISCOSIPUAMIB_CSipCfgEarlyMediaTable_CSipCfgEarlyMediaEntry
}

func (cSipCfgEarlyMediaTable *CISCOSIPUAMIB_CSipCfgEarlyMediaTable) GetEntityData() *types.CommonEntityData {
    cSipCfgEarlyMediaTable.EntityData.YFilter = cSipCfgEarlyMediaTable.YFilter
    cSipCfgEarlyMediaTable.EntityData.YangName = "cSipCfgEarlyMediaTable"
    cSipCfgEarlyMediaTable.EntityData.BundleName = "cisco_ios_xe"
    cSipCfgEarlyMediaTable.EntityData.ParentYangName = "CISCO-SIP-UA-MIB"
    cSipCfgEarlyMediaTable.EntityData.SegmentPath = "cSipCfgEarlyMediaTable"
    cSipCfgEarlyMediaTable.EntityData.AbsolutePath = "CISCO-SIP-UA-MIB:CISCO-SIP-UA-MIB/" + cSipCfgEarlyMediaTable.EntityData.SegmentPath
    cSipCfgEarlyMediaTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cSipCfgEarlyMediaTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cSipCfgEarlyMediaTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cSipCfgEarlyMediaTable.EntityData.Children = types.NewOrderedMap()
    cSipCfgEarlyMediaTable.EntityData.Children.Append("cSipCfgEarlyMediaEntry", types.YChild{"CSipCfgEarlyMediaEntry", nil})
    for i := range cSipCfgEarlyMediaTable.CSipCfgEarlyMediaEntry {
        cSipCfgEarlyMediaTable.EntityData.Children.Append(types.GetSegmentPath(cSipCfgEarlyMediaTable.CSipCfgEarlyMediaEntry[i]), types.YChild{"CSipCfgEarlyMediaEntry", cSipCfgEarlyMediaTable.CSipCfgEarlyMediaEntry[i]})
    }
    cSipCfgEarlyMediaTable.EntityData.Leafs = types.NewOrderedMap()

    cSipCfgEarlyMediaTable.EntityData.YListKeys = []string {}

    return &(cSipCfgEarlyMediaTable.EntityData)
}

// CISCOSIPUAMIB_CSipCfgEarlyMediaTable_CSipCfgEarlyMediaEntry
// A row in the cSipCfgEarlyMediaTable.
// A row is accessible with a Provisional (1xx)
// status code value (eg, 180) and provides
// an object for the enabling or disabling of
// the Early Media Cut Through functionality.
type CISCOSIPUAMIB_CSipCfgEarlyMediaTable_CSipCfgEarlyMediaEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. A unique identifier of a row in this table and a
    // valid SIP status code. The type is interface{} with range: 1..2147483647.
    CSipCfgEarlyMediaStatusCodeIndex interface{}

    // This object specifies whether Early Media  Cut Through is enabled or
    // disabled for the  SIP response messages with status codes that  match
    // cSipCfgEarlyMediaStatusCodeIndex.  If 'true', early media cut through is
    // disabled, and the user agent will process the message as though it did not
    // contain any SDP payload.  If 'false', early media cut through is enabled,
    // and the user agent will process the message similar to a 183 (Session
    // Progress) and cut through for early media.  The assumption being that the
    // SDP is an indication that the far end is going to send early media. The
    // type is bool.
    CSipCfgEarlyMediaCutThruDisabled interface{}
}

func (cSipCfgEarlyMediaEntry *CISCOSIPUAMIB_CSipCfgEarlyMediaTable_CSipCfgEarlyMediaEntry) GetEntityData() *types.CommonEntityData {
    cSipCfgEarlyMediaEntry.EntityData.YFilter = cSipCfgEarlyMediaEntry.YFilter
    cSipCfgEarlyMediaEntry.EntityData.YangName = "cSipCfgEarlyMediaEntry"
    cSipCfgEarlyMediaEntry.EntityData.BundleName = "cisco_ios_xe"
    cSipCfgEarlyMediaEntry.EntityData.ParentYangName = "cSipCfgEarlyMediaTable"
    cSipCfgEarlyMediaEntry.EntityData.SegmentPath = "cSipCfgEarlyMediaEntry" + types.AddKeyToken(cSipCfgEarlyMediaEntry.CSipCfgEarlyMediaStatusCodeIndex, "cSipCfgEarlyMediaStatusCodeIndex")
    cSipCfgEarlyMediaEntry.EntityData.AbsolutePath = "CISCO-SIP-UA-MIB:CISCO-SIP-UA-MIB/cSipCfgEarlyMediaTable/" + cSipCfgEarlyMediaEntry.EntityData.SegmentPath
    cSipCfgEarlyMediaEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cSipCfgEarlyMediaEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cSipCfgEarlyMediaEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cSipCfgEarlyMediaEntry.EntityData.Children = types.NewOrderedMap()
    cSipCfgEarlyMediaEntry.EntityData.Leafs = types.NewOrderedMap()
    cSipCfgEarlyMediaEntry.EntityData.Leafs.Append("cSipCfgEarlyMediaStatusCodeIndex", types.YLeaf{"CSipCfgEarlyMediaStatusCodeIndex", cSipCfgEarlyMediaEntry.CSipCfgEarlyMediaStatusCodeIndex})
    cSipCfgEarlyMediaEntry.EntityData.Leafs.Append("cSipCfgEarlyMediaCutThruDisabled", types.YLeaf{"CSipCfgEarlyMediaCutThruDisabled", cSipCfgEarlyMediaEntry.CSipCfgEarlyMediaCutThruDisabled})

    cSipCfgEarlyMediaEntry.EntityData.YListKeys = []string {"CSipCfgEarlyMediaStatusCodeIndex"}

    return &(cSipCfgEarlyMediaEntry.EntityData)
}

// CISCOSIPUAMIB_CSipCfgBindSourceAddrTable
// This table contains configuration for binding
// the scope of packets to the particular ethernet
// interface. The scope for the packets can be
// specified as either 'signalling' or 'media'
// packets. The ethernet interface shall be
// specified by the interface index. The table
// shall be indexed based on the scope.
type CISCOSIPUAMIB_CSipCfgBindSourceAddrTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A row in the cSipCfgBindSourceAddrTable. A row is accessible with the scope
    // of packets to which the source IP address of the interface designated by
    // cSipCfgBindSourceAddrInterface will be bound. The type is slice of
    // CISCOSIPUAMIB_CSipCfgBindSourceAddrTable_CSipCfgBindSourceAddrEntry.
    CSipCfgBindSourceAddrEntry []*CISCOSIPUAMIB_CSipCfgBindSourceAddrTable_CSipCfgBindSourceAddrEntry
}

func (cSipCfgBindSourceAddrTable *CISCOSIPUAMIB_CSipCfgBindSourceAddrTable) GetEntityData() *types.CommonEntityData {
    cSipCfgBindSourceAddrTable.EntityData.YFilter = cSipCfgBindSourceAddrTable.YFilter
    cSipCfgBindSourceAddrTable.EntityData.YangName = "cSipCfgBindSourceAddrTable"
    cSipCfgBindSourceAddrTable.EntityData.BundleName = "cisco_ios_xe"
    cSipCfgBindSourceAddrTable.EntityData.ParentYangName = "CISCO-SIP-UA-MIB"
    cSipCfgBindSourceAddrTable.EntityData.SegmentPath = "cSipCfgBindSourceAddrTable"
    cSipCfgBindSourceAddrTable.EntityData.AbsolutePath = "CISCO-SIP-UA-MIB:CISCO-SIP-UA-MIB/" + cSipCfgBindSourceAddrTable.EntityData.SegmentPath
    cSipCfgBindSourceAddrTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cSipCfgBindSourceAddrTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cSipCfgBindSourceAddrTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cSipCfgBindSourceAddrTable.EntityData.Children = types.NewOrderedMap()
    cSipCfgBindSourceAddrTable.EntityData.Children.Append("cSipCfgBindSourceAddrEntry", types.YChild{"CSipCfgBindSourceAddrEntry", nil})
    for i := range cSipCfgBindSourceAddrTable.CSipCfgBindSourceAddrEntry {
        cSipCfgBindSourceAddrTable.EntityData.Children.Append(types.GetSegmentPath(cSipCfgBindSourceAddrTable.CSipCfgBindSourceAddrEntry[i]), types.YChild{"CSipCfgBindSourceAddrEntry", cSipCfgBindSourceAddrTable.CSipCfgBindSourceAddrEntry[i]})
    }
    cSipCfgBindSourceAddrTable.EntityData.Leafs = types.NewOrderedMap()

    cSipCfgBindSourceAddrTable.EntityData.YListKeys = []string {}

    return &(cSipCfgBindSourceAddrTable.EntityData)
}

// CISCOSIPUAMIB_CSipCfgBindSourceAddrTable_CSipCfgBindSourceAddrEntry
// A row in the cSipCfgBindSourceAddrTable.
// A row is accessible with the scope of packets
// to which the source IP address of the interface
// designated by cSipCfgBindSourceAddrInterface
// will be bound.
type CISCOSIPUAMIB_CSipCfgBindSourceAddrTable_CSipCfgBindSourceAddrEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. A unique identifier of a row in this table and
    // specifies the scope of packets to which the source IP address of the
    // interface designated by cSipCfgBindSourceAddrInterface will be bound. The
    // type is CSipCfgBindSourceAddrScope.
    CSipCfgBindSourceAddrScope interface{}

    // This object may specify the interface where the source IP address used in
    // SIP signalling or media packets is configured.  A value of 0 means that
    // there is no specific source address configured and in this case the best
    // local IP address will be chosen by the system. The type is interface{} with
    // range: 0..2147483647.
    CSipCfgBindSourceAddrInterface interface{}
}

func (cSipCfgBindSourceAddrEntry *CISCOSIPUAMIB_CSipCfgBindSourceAddrTable_CSipCfgBindSourceAddrEntry) GetEntityData() *types.CommonEntityData {
    cSipCfgBindSourceAddrEntry.EntityData.YFilter = cSipCfgBindSourceAddrEntry.YFilter
    cSipCfgBindSourceAddrEntry.EntityData.YangName = "cSipCfgBindSourceAddrEntry"
    cSipCfgBindSourceAddrEntry.EntityData.BundleName = "cisco_ios_xe"
    cSipCfgBindSourceAddrEntry.EntityData.ParentYangName = "cSipCfgBindSourceAddrTable"
    cSipCfgBindSourceAddrEntry.EntityData.SegmentPath = "cSipCfgBindSourceAddrEntry" + types.AddKeyToken(cSipCfgBindSourceAddrEntry.CSipCfgBindSourceAddrScope, "cSipCfgBindSourceAddrScope")
    cSipCfgBindSourceAddrEntry.EntityData.AbsolutePath = "CISCO-SIP-UA-MIB:CISCO-SIP-UA-MIB/cSipCfgBindSourceAddrTable/" + cSipCfgBindSourceAddrEntry.EntityData.SegmentPath
    cSipCfgBindSourceAddrEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cSipCfgBindSourceAddrEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cSipCfgBindSourceAddrEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cSipCfgBindSourceAddrEntry.EntityData.Children = types.NewOrderedMap()
    cSipCfgBindSourceAddrEntry.EntityData.Leafs = types.NewOrderedMap()
    cSipCfgBindSourceAddrEntry.EntityData.Leafs.Append("cSipCfgBindSourceAddrScope", types.YLeaf{"CSipCfgBindSourceAddrScope", cSipCfgBindSourceAddrEntry.CSipCfgBindSourceAddrScope})
    cSipCfgBindSourceAddrEntry.EntityData.Leafs.Append("cSipCfgBindSourceAddrInterface", types.YLeaf{"CSipCfgBindSourceAddrInterface", cSipCfgBindSourceAddrEntry.CSipCfgBindSourceAddrInterface})

    cSipCfgBindSourceAddrEntry.EntityData.YListKeys = []string {"CSipCfgBindSourceAddrScope"}

    return &(cSipCfgBindSourceAddrEntry.EntityData)
}

// CISCOSIPUAMIB_CSipCfgBindSourceAddrTable_CSipCfgBindSourceAddrEntry_CSipCfgBindSourceAddrScope represents will be bound.
type CISCOSIPUAMIB_CSipCfgBindSourceAddrTable_CSipCfgBindSourceAddrEntry_CSipCfgBindSourceAddrScope string

const (
    CISCOSIPUAMIB_CSipCfgBindSourceAddrTable_CSipCfgBindSourceAddrEntry_CSipCfgBindSourceAddrScope_media CISCOSIPUAMIB_CSipCfgBindSourceAddrTable_CSipCfgBindSourceAddrEntry_CSipCfgBindSourceAddrScope = "media"

    CISCOSIPUAMIB_CSipCfgBindSourceAddrTable_CSipCfgBindSourceAddrEntry_CSipCfgBindSourceAddrScope_control CISCOSIPUAMIB_CSipCfgBindSourceAddrTable_CSipCfgBindSourceAddrEntry_CSipCfgBindSourceAddrScope = "control"
)

// CISCOSIPUAMIB_CSipCfgPeerTable
// This table contains per dial-peer SIP related 
// configuration.   
// 
// The table is a sparse table of dial-peer information.
// This means, it only reflects dial-peers being used 
// for SIP.  A dial-peer is being used for SIP if the 
// value of cvVoIPPeerCfgSessionProtocol 
// (CISCO-VOICE-DIAL-CONTROL-MIB) is 'sip'.
// 
// Dial-peers are not created or destroyed via this
// table.  Only SIP related configuration can be 
// performed via this table once the dial-peer exists
// in the system and is visible in this table.
type CISCOSIPUAMIB_CSipCfgPeerTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A row in the cSipCfgPeerTable. The type is slice of
    // CISCOSIPUAMIB_CSipCfgPeerTable_CSipCfgPeerEntry.
    CSipCfgPeerEntry []*CISCOSIPUAMIB_CSipCfgPeerTable_CSipCfgPeerEntry
}

func (cSipCfgPeerTable *CISCOSIPUAMIB_CSipCfgPeerTable) GetEntityData() *types.CommonEntityData {
    cSipCfgPeerTable.EntityData.YFilter = cSipCfgPeerTable.YFilter
    cSipCfgPeerTable.EntityData.YangName = "cSipCfgPeerTable"
    cSipCfgPeerTable.EntityData.BundleName = "cisco_ios_xe"
    cSipCfgPeerTable.EntityData.ParentYangName = "CISCO-SIP-UA-MIB"
    cSipCfgPeerTable.EntityData.SegmentPath = "cSipCfgPeerTable"
    cSipCfgPeerTable.EntityData.AbsolutePath = "CISCO-SIP-UA-MIB:CISCO-SIP-UA-MIB/" + cSipCfgPeerTable.EntityData.SegmentPath
    cSipCfgPeerTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cSipCfgPeerTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cSipCfgPeerTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cSipCfgPeerTable.EntityData.Children = types.NewOrderedMap()
    cSipCfgPeerTable.EntityData.Children.Append("cSipCfgPeerEntry", types.YChild{"CSipCfgPeerEntry", nil})
    for i := range cSipCfgPeerTable.CSipCfgPeerEntry {
        cSipCfgPeerTable.EntityData.Children.Append(types.GetSegmentPath(cSipCfgPeerTable.CSipCfgPeerEntry[i]), types.YChild{"CSipCfgPeerEntry", cSipCfgPeerTable.CSipCfgPeerEntry[i]})
    }
    cSipCfgPeerTable.EntityData.Leafs = types.NewOrderedMap()

    cSipCfgPeerTable.EntityData.YListKeys = []string {}

    return &(cSipCfgPeerTable.EntityData)
}

// CISCOSIPUAMIB_CSipCfgPeerTable_CSipCfgPeerEntry
// A row in the cSipCfgPeerTable.
type CISCOSIPUAMIB_CSipCfgPeerTable_CSipCfgPeerEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. An arbitrary index that uniquely identifies a 
    // dial-peer configured for SIP. The type is interface{} with range:
    // 1..2147483647.
    CSipCfgPeerIndex interface{}

    // This object specifies the session transport  protocol that will be used by
    // this dial-peer for outbound SIP messages.    The value 'system' is the
    // default and indicates  that this dial-peer should use the value set by 
    // cSipCfgOutSessionTransport instead. The type is
    // CSipCfgPeerOutSessionTransport.
    CSipCfgPeerOutSessionTransport interface{}

    // This object specifies the string that will be  placed in either the
    // Supported or Require SIP  header, as specified by
    // cSipCfgPeerReliable1xxRspHdr. The type is string.
    CSipCfgPeerReliable1xxRspStr interface{}

    // This object specifies behavior with respect to Support or Require headers
    // in SIP messages sent and received via this dial-peer.  If the originating
    // gateway is configured for 'require', the Require header is added to the
    // outgoing INVITEs with the value of cSipCfgPeerReliable1xxStr.  This
    // requires the use of reliable provisional responses by the terminating
    // gateway.  Sessions will be torn down if this use cannot be supported by
    // that gateway.  If the originating gateway is configured for 'supported',
    // the Supported header is added to the outgoing INVITEs with the value of
    // cSipCfgPeerReliable1xxStr.  This  requires that an attempt to use reliable
    // provisional responses be made, but sessions can continue without them.  If
    // the originating gateway is configured for 'disabled', the value of
    // cSipCfgReliable1xxStr will NOT be added to either the Require or Supported
    // headers of outgoing INVITEs.  The value 'system' is the default and
    // indicates that this  dial-peer should use the value of 
    // cSipCfgReliable1xxRspHdr instead. The type is CSipCfgPeerReliable1xxRspHdr.
    CSipCfgPeerReliable1xxRspHdr interface{}

    // This object specifies the URL type sent in outbound INVITES generated by
    // this device.  The value 'system' is the default and indicates that this 
    // dial-peer should use the value of cSipCfgUrlType instead. The type is
    // CSipCfgPeerUrlType.
    CSipCfgPeerUrlType interface{}

    // This object specifies if the functionality of switching between transports
    // from UDP to TCP if the message size of a Request is greater than 1300 bytes
    // is enabled or not. The type is bool.
    CSipCfgPeerSwitchTransEnabled interface{}
}

func (cSipCfgPeerEntry *CISCOSIPUAMIB_CSipCfgPeerTable_CSipCfgPeerEntry) GetEntityData() *types.CommonEntityData {
    cSipCfgPeerEntry.EntityData.YFilter = cSipCfgPeerEntry.YFilter
    cSipCfgPeerEntry.EntityData.YangName = "cSipCfgPeerEntry"
    cSipCfgPeerEntry.EntityData.BundleName = "cisco_ios_xe"
    cSipCfgPeerEntry.EntityData.ParentYangName = "cSipCfgPeerTable"
    cSipCfgPeerEntry.EntityData.SegmentPath = "cSipCfgPeerEntry" + types.AddKeyToken(cSipCfgPeerEntry.CSipCfgPeerIndex, "cSipCfgPeerIndex")
    cSipCfgPeerEntry.EntityData.AbsolutePath = "CISCO-SIP-UA-MIB:CISCO-SIP-UA-MIB/cSipCfgPeerTable/" + cSipCfgPeerEntry.EntityData.SegmentPath
    cSipCfgPeerEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cSipCfgPeerEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cSipCfgPeerEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cSipCfgPeerEntry.EntityData.Children = types.NewOrderedMap()
    cSipCfgPeerEntry.EntityData.Leafs = types.NewOrderedMap()
    cSipCfgPeerEntry.EntityData.Leafs.Append("cSipCfgPeerIndex", types.YLeaf{"CSipCfgPeerIndex", cSipCfgPeerEntry.CSipCfgPeerIndex})
    cSipCfgPeerEntry.EntityData.Leafs.Append("cSipCfgPeerOutSessionTransport", types.YLeaf{"CSipCfgPeerOutSessionTransport", cSipCfgPeerEntry.CSipCfgPeerOutSessionTransport})
    cSipCfgPeerEntry.EntityData.Leafs.Append("cSipCfgPeerReliable1xxRspStr", types.YLeaf{"CSipCfgPeerReliable1xxRspStr", cSipCfgPeerEntry.CSipCfgPeerReliable1xxRspStr})
    cSipCfgPeerEntry.EntityData.Leafs.Append("cSipCfgPeerReliable1xxRspHdr", types.YLeaf{"CSipCfgPeerReliable1xxRspHdr", cSipCfgPeerEntry.CSipCfgPeerReliable1xxRspHdr})
    cSipCfgPeerEntry.EntityData.Leafs.Append("cSipCfgPeerUrlType", types.YLeaf{"CSipCfgPeerUrlType", cSipCfgPeerEntry.CSipCfgPeerUrlType})
    cSipCfgPeerEntry.EntityData.Leafs.Append("cSipCfgPeerSwitchTransEnabled", types.YLeaf{"CSipCfgPeerSwitchTransEnabled", cSipCfgPeerEntry.CSipCfgPeerSwitchTransEnabled})

    cSipCfgPeerEntry.EntityData.YListKeys = []string {"CSipCfgPeerIndex"}

    return &(cSipCfgPeerEntry.EntityData)
}

// CISCOSIPUAMIB_CSipCfgPeerTable_CSipCfgPeerEntry_CSipCfgPeerOutSessionTransport represents cSipCfgOutSessionTransport instead.
type CISCOSIPUAMIB_CSipCfgPeerTable_CSipCfgPeerEntry_CSipCfgPeerOutSessionTransport string

const (
    CISCOSIPUAMIB_CSipCfgPeerTable_CSipCfgPeerEntry_CSipCfgPeerOutSessionTransport_system CISCOSIPUAMIB_CSipCfgPeerTable_CSipCfgPeerEntry_CSipCfgPeerOutSessionTransport = "system"

    CISCOSIPUAMIB_CSipCfgPeerTable_CSipCfgPeerEntry_CSipCfgPeerOutSessionTransport_udp CISCOSIPUAMIB_CSipCfgPeerTable_CSipCfgPeerEntry_CSipCfgPeerOutSessionTransport = "udp"

    CISCOSIPUAMIB_CSipCfgPeerTable_CSipCfgPeerEntry_CSipCfgPeerOutSessionTransport_tcp CISCOSIPUAMIB_CSipCfgPeerTable_CSipCfgPeerEntry_CSipCfgPeerOutSessionTransport = "tcp"
)

// CISCOSIPUAMIB_CSipCfgPeerTable_CSipCfgPeerEntry_CSipCfgPeerReliable1xxRspHdr represents instead.
type CISCOSIPUAMIB_CSipCfgPeerTable_CSipCfgPeerEntry_CSipCfgPeerReliable1xxRspHdr string

const (
    CISCOSIPUAMIB_CSipCfgPeerTable_CSipCfgPeerEntry_CSipCfgPeerReliable1xxRspHdr_system CISCOSIPUAMIB_CSipCfgPeerTable_CSipCfgPeerEntry_CSipCfgPeerReliable1xxRspHdr = "system"

    CISCOSIPUAMIB_CSipCfgPeerTable_CSipCfgPeerEntry_CSipCfgPeerReliable1xxRspHdr_supported CISCOSIPUAMIB_CSipCfgPeerTable_CSipCfgPeerEntry_CSipCfgPeerReliable1xxRspHdr = "supported"

    CISCOSIPUAMIB_CSipCfgPeerTable_CSipCfgPeerEntry_CSipCfgPeerReliable1xxRspHdr_require CISCOSIPUAMIB_CSipCfgPeerTable_CSipCfgPeerEntry_CSipCfgPeerReliable1xxRspHdr = "require"

    CISCOSIPUAMIB_CSipCfgPeerTable_CSipCfgPeerEntry_CSipCfgPeerReliable1xxRspHdr_disabled CISCOSIPUAMIB_CSipCfgPeerTable_CSipCfgPeerEntry_CSipCfgPeerReliable1xxRspHdr = "disabled"
)

// CISCOSIPUAMIB_CSipCfgPeerTable_CSipCfgPeerEntry_CSipCfgPeerUrlType represents dial-peer should use the value of cSipCfgUrlType instead.
type CISCOSIPUAMIB_CSipCfgPeerTable_CSipCfgPeerEntry_CSipCfgPeerUrlType string

const (
    CISCOSIPUAMIB_CSipCfgPeerTable_CSipCfgPeerEntry_CSipCfgPeerUrlType_system CISCOSIPUAMIB_CSipCfgPeerTable_CSipCfgPeerEntry_CSipCfgPeerUrlType = "system"

    CISCOSIPUAMIB_CSipCfgPeerTable_CSipCfgPeerEntry_CSipCfgPeerUrlType_sip CISCOSIPUAMIB_CSipCfgPeerTable_CSipCfgPeerEntry_CSipCfgPeerUrlType = "sip"

    CISCOSIPUAMIB_CSipCfgPeerTable_CSipCfgPeerEntry_CSipCfgPeerUrlType_tel CISCOSIPUAMIB_CSipCfgPeerTable_CSipCfgPeerEntry_CSipCfgPeerUrlType = "tel"
)

// CISCOSIPUAMIB_CSipCfgStatusCauseTable
// This table contains SIP status code to PSTN cause code
// mapping configuration.  Inbound SIP response messages 
// that will result in outbound PSTN signalling messages
// will have the SIP status codes transposed into the
// PSTN cause codes as prescribed by the contents of this 
// table.
type CISCOSIPUAMIB_CSipCfgStatusCauseTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A row in the cSipCfgStatusCauseTable.  Entries cannot be created or
    // destroyed by the actions of any NMS. The type is slice of
    // CISCOSIPUAMIB_CSipCfgStatusCauseTable_CSipCfgStatusCauseEntry.
    CSipCfgStatusCauseEntry []*CISCOSIPUAMIB_CSipCfgStatusCauseTable_CSipCfgStatusCauseEntry
}

func (cSipCfgStatusCauseTable *CISCOSIPUAMIB_CSipCfgStatusCauseTable) GetEntityData() *types.CommonEntityData {
    cSipCfgStatusCauseTable.EntityData.YFilter = cSipCfgStatusCauseTable.YFilter
    cSipCfgStatusCauseTable.EntityData.YangName = "cSipCfgStatusCauseTable"
    cSipCfgStatusCauseTable.EntityData.BundleName = "cisco_ios_xe"
    cSipCfgStatusCauseTable.EntityData.ParentYangName = "CISCO-SIP-UA-MIB"
    cSipCfgStatusCauseTable.EntityData.SegmentPath = "cSipCfgStatusCauseTable"
    cSipCfgStatusCauseTable.EntityData.AbsolutePath = "CISCO-SIP-UA-MIB:CISCO-SIP-UA-MIB/" + cSipCfgStatusCauseTable.EntityData.SegmentPath
    cSipCfgStatusCauseTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cSipCfgStatusCauseTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cSipCfgStatusCauseTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cSipCfgStatusCauseTable.EntityData.Children = types.NewOrderedMap()
    cSipCfgStatusCauseTable.EntityData.Children.Append("cSipCfgStatusCauseEntry", types.YChild{"CSipCfgStatusCauseEntry", nil})
    for i := range cSipCfgStatusCauseTable.CSipCfgStatusCauseEntry {
        cSipCfgStatusCauseTable.EntityData.Children.Append(types.GetSegmentPath(cSipCfgStatusCauseTable.CSipCfgStatusCauseEntry[i]), types.YChild{"CSipCfgStatusCauseEntry", cSipCfgStatusCauseTable.CSipCfgStatusCauseEntry[i]})
    }
    cSipCfgStatusCauseTable.EntityData.Leafs = types.NewOrderedMap()

    cSipCfgStatusCauseTable.EntityData.YListKeys = []string {}

    return &(cSipCfgStatusCauseTable.EntityData)
}

// CISCOSIPUAMIB_CSipCfgStatusCauseTable_CSipCfgStatusCauseEntry
// A row in the cSipCfgStatusCauseTable.  Entries cannot be
// created or destroyed by the actions of any NMS.
type CISCOSIPUAMIB_CSipCfgStatusCauseTable_CSipCfgStatusCauseEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. A unique identifier of a row in this table and a
    // valid SIP status code. The type is interface{} with range: 1..2147483647.
    CSipCfgStatusCodeIndex interface{}

    // The PSTN cause code to which the SIP status code given by
    // cSipCfgStatusCodeIndex is to be mapped for outbound PSTN signalling
    // messages. The type is interface{} with range: 1..2147483647.
    CSipCfgPstnCause interface{}

    // This object reflects the storage status of this table entry.  If the value
    // is 'volatile', then this entry only exists in RAM and the information would
    // be lost (reverting to defaults) on system reload.   If the value is
    // 'nonVolatile' then this entry has been  written to NVRAM and will persist
    // across system reloads. The type is StorageType.
    CSipCfgStatusCauseStoreStatus interface{}
}

func (cSipCfgStatusCauseEntry *CISCOSIPUAMIB_CSipCfgStatusCauseTable_CSipCfgStatusCauseEntry) GetEntityData() *types.CommonEntityData {
    cSipCfgStatusCauseEntry.EntityData.YFilter = cSipCfgStatusCauseEntry.YFilter
    cSipCfgStatusCauseEntry.EntityData.YangName = "cSipCfgStatusCauseEntry"
    cSipCfgStatusCauseEntry.EntityData.BundleName = "cisco_ios_xe"
    cSipCfgStatusCauseEntry.EntityData.ParentYangName = "cSipCfgStatusCauseTable"
    cSipCfgStatusCauseEntry.EntityData.SegmentPath = "cSipCfgStatusCauseEntry" + types.AddKeyToken(cSipCfgStatusCauseEntry.CSipCfgStatusCodeIndex, "cSipCfgStatusCodeIndex")
    cSipCfgStatusCauseEntry.EntityData.AbsolutePath = "CISCO-SIP-UA-MIB:CISCO-SIP-UA-MIB/cSipCfgStatusCauseTable/" + cSipCfgStatusCauseEntry.EntityData.SegmentPath
    cSipCfgStatusCauseEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cSipCfgStatusCauseEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cSipCfgStatusCauseEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cSipCfgStatusCauseEntry.EntityData.Children = types.NewOrderedMap()
    cSipCfgStatusCauseEntry.EntityData.Leafs = types.NewOrderedMap()
    cSipCfgStatusCauseEntry.EntityData.Leafs.Append("cSipCfgStatusCodeIndex", types.YLeaf{"CSipCfgStatusCodeIndex", cSipCfgStatusCauseEntry.CSipCfgStatusCodeIndex})
    cSipCfgStatusCauseEntry.EntityData.Leafs.Append("cSipCfgPstnCause", types.YLeaf{"CSipCfgPstnCause", cSipCfgStatusCauseEntry.CSipCfgPstnCause})
    cSipCfgStatusCauseEntry.EntityData.Leafs.Append("cSipCfgStatusCauseStoreStatus", types.YLeaf{"CSipCfgStatusCauseStoreStatus", cSipCfgStatusCauseEntry.CSipCfgStatusCauseStoreStatus})

    cSipCfgStatusCauseEntry.EntityData.YListKeys = []string {"CSipCfgStatusCodeIndex"}

    return &(cSipCfgStatusCauseEntry.EntityData)
}

// CISCOSIPUAMIB_CSipCfgCauseStatusTable
// This table contains PSTN cause code to SIP status code
// mapping configuration.   Inbound PSTN signalling messages
// that will result in outbound SIP response messages 
// will have the PSTN cause codes transposed into the
// SIP status codes as prescribed by the contents of this 
// table.
type CISCOSIPUAMIB_CSipCfgCauseStatusTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A row in the cSipCfgCauseStatusTable. Entries cannot be created or
    // destroyed by the actions of any NMS. The type is slice of
    // CISCOSIPUAMIB_CSipCfgCauseStatusTable_CSipCfgCauseStatusEntry.
    CSipCfgCauseStatusEntry []*CISCOSIPUAMIB_CSipCfgCauseStatusTable_CSipCfgCauseStatusEntry
}

func (cSipCfgCauseStatusTable *CISCOSIPUAMIB_CSipCfgCauseStatusTable) GetEntityData() *types.CommonEntityData {
    cSipCfgCauseStatusTable.EntityData.YFilter = cSipCfgCauseStatusTable.YFilter
    cSipCfgCauseStatusTable.EntityData.YangName = "cSipCfgCauseStatusTable"
    cSipCfgCauseStatusTable.EntityData.BundleName = "cisco_ios_xe"
    cSipCfgCauseStatusTable.EntityData.ParentYangName = "CISCO-SIP-UA-MIB"
    cSipCfgCauseStatusTable.EntityData.SegmentPath = "cSipCfgCauseStatusTable"
    cSipCfgCauseStatusTable.EntityData.AbsolutePath = "CISCO-SIP-UA-MIB:CISCO-SIP-UA-MIB/" + cSipCfgCauseStatusTable.EntityData.SegmentPath
    cSipCfgCauseStatusTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cSipCfgCauseStatusTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cSipCfgCauseStatusTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cSipCfgCauseStatusTable.EntityData.Children = types.NewOrderedMap()
    cSipCfgCauseStatusTable.EntityData.Children.Append("cSipCfgCauseStatusEntry", types.YChild{"CSipCfgCauseStatusEntry", nil})
    for i := range cSipCfgCauseStatusTable.CSipCfgCauseStatusEntry {
        cSipCfgCauseStatusTable.EntityData.Children.Append(types.GetSegmentPath(cSipCfgCauseStatusTable.CSipCfgCauseStatusEntry[i]), types.YChild{"CSipCfgCauseStatusEntry", cSipCfgCauseStatusTable.CSipCfgCauseStatusEntry[i]})
    }
    cSipCfgCauseStatusTable.EntityData.Leafs = types.NewOrderedMap()

    cSipCfgCauseStatusTable.EntityData.YListKeys = []string {}

    return &(cSipCfgCauseStatusTable.EntityData)
}

// CISCOSIPUAMIB_CSipCfgCauseStatusTable_CSipCfgCauseStatusEntry
// A row in the cSipCfgCauseStatusTable. Entries cannot be
// created or destroyed by the actions of any NMS.
type CISCOSIPUAMIB_CSipCfgCauseStatusTable_CSipCfgCauseStatusEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. A unique identifier of a row in this table and a
    // valid PSTN cause code. The type is interface{} with range: 1..2147483647.
    CSipCfgPstnCauseIndex interface{}

    // The SIP status code to which the PSTN cause code given by
    // cSipCfgPstnCauseIndex is to be mapped for outbound SIP response messages.
    // The type is interface{} with range: 1..2147483647.
    CSipCfgStatusCode interface{}

    // This object reflects the storage status of this table entry.  If the value
    // is 'volatile', then this entry only exists in RAM and the information would
    // be lost (reverting to defaults) on system reload.   If the value is
    // 'nonVolatile' then this entry has been  written to NVRAM and will persist
    // across system reloads. The type is StorageType.
    CSipCfgCauseStatusStoreStatus interface{}
}

func (cSipCfgCauseStatusEntry *CISCOSIPUAMIB_CSipCfgCauseStatusTable_CSipCfgCauseStatusEntry) GetEntityData() *types.CommonEntityData {
    cSipCfgCauseStatusEntry.EntityData.YFilter = cSipCfgCauseStatusEntry.YFilter
    cSipCfgCauseStatusEntry.EntityData.YangName = "cSipCfgCauseStatusEntry"
    cSipCfgCauseStatusEntry.EntityData.BundleName = "cisco_ios_xe"
    cSipCfgCauseStatusEntry.EntityData.ParentYangName = "cSipCfgCauseStatusTable"
    cSipCfgCauseStatusEntry.EntityData.SegmentPath = "cSipCfgCauseStatusEntry" + types.AddKeyToken(cSipCfgCauseStatusEntry.CSipCfgPstnCauseIndex, "cSipCfgPstnCauseIndex")
    cSipCfgCauseStatusEntry.EntityData.AbsolutePath = "CISCO-SIP-UA-MIB:CISCO-SIP-UA-MIB/cSipCfgCauseStatusTable/" + cSipCfgCauseStatusEntry.EntityData.SegmentPath
    cSipCfgCauseStatusEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cSipCfgCauseStatusEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cSipCfgCauseStatusEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cSipCfgCauseStatusEntry.EntityData.Children = types.NewOrderedMap()
    cSipCfgCauseStatusEntry.EntityData.Leafs = types.NewOrderedMap()
    cSipCfgCauseStatusEntry.EntityData.Leafs.Append("cSipCfgPstnCauseIndex", types.YLeaf{"CSipCfgPstnCauseIndex", cSipCfgCauseStatusEntry.CSipCfgPstnCauseIndex})
    cSipCfgCauseStatusEntry.EntityData.Leafs.Append("cSipCfgStatusCode", types.YLeaf{"CSipCfgStatusCode", cSipCfgCauseStatusEntry.CSipCfgStatusCode})
    cSipCfgCauseStatusEntry.EntityData.Leafs.Append("cSipCfgCauseStatusStoreStatus", types.YLeaf{"CSipCfgCauseStatusStoreStatus", cSipCfgCauseStatusEntry.CSipCfgCauseStatusStoreStatus})

    cSipCfgCauseStatusEntry.EntityData.YListKeys = []string {"CSipCfgPstnCauseIndex"}

    return &(cSipCfgCauseStatusEntry.EntityData)
}

// CISCOSIPUAMIB_CSipStatsSuccessOkTable
// This table contains statistics for sent and
// received 200 Ok response messages.  The 
// statistics are kept on per SIP method basis.
type CISCOSIPUAMIB_CSipStatsSuccessOkTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A row in the cSipStatsSuccessOkTable.  There is  an entry for each SIP
    // method for which 200 Ok  inbound and/or outbound statistics are kept. The
    // type is slice of
    // CISCOSIPUAMIB_CSipStatsSuccessOkTable_CSipStatsSuccessOkEntry.
    CSipStatsSuccessOkEntry []*CISCOSIPUAMIB_CSipStatsSuccessOkTable_CSipStatsSuccessOkEntry
}

func (cSipStatsSuccessOkTable *CISCOSIPUAMIB_CSipStatsSuccessOkTable) GetEntityData() *types.CommonEntityData {
    cSipStatsSuccessOkTable.EntityData.YFilter = cSipStatsSuccessOkTable.YFilter
    cSipStatsSuccessOkTable.EntityData.YangName = "cSipStatsSuccessOkTable"
    cSipStatsSuccessOkTable.EntityData.BundleName = "cisco_ios_xe"
    cSipStatsSuccessOkTable.EntityData.ParentYangName = "CISCO-SIP-UA-MIB"
    cSipStatsSuccessOkTable.EntityData.SegmentPath = "cSipStatsSuccessOkTable"
    cSipStatsSuccessOkTable.EntityData.AbsolutePath = "CISCO-SIP-UA-MIB:CISCO-SIP-UA-MIB/" + cSipStatsSuccessOkTable.EntityData.SegmentPath
    cSipStatsSuccessOkTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cSipStatsSuccessOkTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cSipStatsSuccessOkTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cSipStatsSuccessOkTable.EntityData.Children = types.NewOrderedMap()
    cSipStatsSuccessOkTable.EntityData.Children.Append("cSipStatsSuccessOkEntry", types.YChild{"CSipStatsSuccessOkEntry", nil})
    for i := range cSipStatsSuccessOkTable.CSipStatsSuccessOkEntry {
        cSipStatsSuccessOkTable.EntityData.Children.Append(types.GetSegmentPath(cSipStatsSuccessOkTable.CSipStatsSuccessOkEntry[i]), types.YChild{"CSipStatsSuccessOkEntry", cSipStatsSuccessOkTable.CSipStatsSuccessOkEntry[i]})
    }
    cSipStatsSuccessOkTable.EntityData.Leafs = types.NewOrderedMap()

    cSipStatsSuccessOkTable.EntityData.YListKeys = []string {}

    return &(cSipStatsSuccessOkTable.EntityData)
}

// CISCOSIPUAMIB_CSipStatsSuccessOkTable_CSipStatsSuccessOkEntry
// A row in the cSipStatsSuccessOkTable.  There is 
// an entry for each SIP method for which 200 Ok 
// inbound and/or outbound statistics are kept.
type CISCOSIPUAMIB_CSipStatsSuccessOkTable_CSipStatsSuccessOkEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. This object is used for instance identification of
    // objects in this table.  The value reflects a SIP method. The type is string
    // with length: 1..32.
    CSipStatsSuccessOkMethod interface{}

    // This object reflects the total number of Ok (200) responses sent by the
    // user agent, since system startup, that were associated with the SIP method
    // as specified by cSipStatsSuccessOkMethod. The type is interface{} with
    // range: 0..4294967295.
    CSipStatsSuccessOkInbounds interface{}

    // This object reflects the total number of Ok (200) responses received by the
    // user agent, since system startup, that were associated with the SIP method
    // as specified by cSipStatsSuccessOkMethod. The type is interface{} with
    // range: 0..4294967295.
    CSipStatsSuccessOkOutbounds interface{}
}

func (cSipStatsSuccessOkEntry *CISCOSIPUAMIB_CSipStatsSuccessOkTable_CSipStatsSuccessOkEntry) GetEntityData() *types.CommonEntityData {
    cSipStatsSuccessOkEntry.EntityData.YFilter = cSipStatsSuccessOkEntry.YFilter
    cSipStatsSuccessOkEntry.EntityData.YangName = "cSipStatsSuccessOkEntry"
    cSipStatsSuccessOkEntry.EntityData.BundleName = "cisco_ios_xe"
    cSipStatsSuccessOkEntry.EntityData.ParentYangName = "cSipStatsSuccessOkTable"
    cSipStatsSuccessOkEntry.EntityData.SegmentPath = "cSipStatsSuccessOkEntry" + types.AddKeyToken(cSipStatsSuccessOkEntry.CSipStatsSuccessOkMethod, "cSipStatsSuccessOkMethod")
    cSipStatsSuccessOkEntry.EntityData.AbsolutePath = "CISCO-SIP-UA-MIB:CISCO-SIP-UA-MIB/cSipStatsSuccessOkTable/" + cSipStatsSuccessOkEntry.EntityData.SegmentPath
    cSipStatsSuccessOkEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cSipStatsSuccessOkEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cSipStatsSuccessOkEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cSipStatsSuccessOkEntry.EntityData.Children = types.NewOrderedMap()
    cSipStatsSuccessOkEntry.EntityData.Leafs = types.NewOrderedMap()
    cSipStatsSuccessOkEntry.EntityData.Leafs.Append("cSipStatsSuccessOkMethod", types.YLeaf{"CSipStatsSuccessOkMethod", cSipStatsSuccessOkEntry.CSipStatsSuccessOkMethod})
    cSipStatsSuccessOkEntry.EntityData.Leafs.Append("cSipStatsSuccessOkInbounds", types.YLeaf{"CSipStatsSuccessOkInbounds", cSipStatsSuccessOkEntry.CSipStatsSuccessOkInbounds})
    cSipStatsSuccessOkEntry.EntityData.Leafs.Append("cSipStatsSuccessOkOutbounds", types.YLeaf{"CSipStatsSuccessOkOutbounds", cSipStatsSuccessOkEntry.CSipStatsSuccessOkOutbounds})

    cSipStatsSuccessOkEntry.EntityData.YListKeys = []string {"CSipStatsSuccessOkMethod"}

    return &(cSipStatsSuccessOkEntry.EntityData)
}

