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

type Ciscosipuamibnotificationprefix struct {
}

func (id Ciscosipuamibnotificationprefix) String() string {
	return "CISCO-SIP-UA-MIB:ciscoSipUaMIBNotificationPrefix"
}

type Ciscosipuamibnotifications struct {
}

func (id Ciscosipuamibnotifications) String() string {
	return "CISCO-SIP-UA-MIB:ciscoSipUaMIBNotifications"
}

// CISCOSIPUAMIB
type CISCOSIPUAMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Csipcfgbase CISCOSIPUAMIB_Csipcfgbase

    
    Csipcfgtimer CISCOSIPUAMIB_Csipcfgtimer

    
    Csipcfgretry CISCOSIPUAMIB_Csipcfgretry

    
    Csipcfgpeer CISCOSIPUAMIB_Csipcfgpeer

    
    Csipcfgaaa CISCOSIPUAMIB_Csipcfgaaa

    
    Csipcfgvoiceservicevoip CISCOSIPUAMIB_Csipcfgvoiceservicevoip

    
    Csipstatsinfo CISCOSIPUAMIB_Csipstatsinfo

    
    Csipstatssuccess CISCOSIPUAMIB_Csipstatssuccess

    
    Csipstatsredirect CISCOSIPUAMIB_Csipstatsredirect

    
    Csipstatserrclient CISCOSIPUAMIB_Csipstatserrclient

    
    Csipstatserrserver CISCOSIPUAMIB_Csipstatserrserver

    
    Csipstatsglobalfail CISCOSIPUAMIB_Csipstatsglobalfail

    
    Csipstatstraffic CISCOSIPUAMIB_Csipstatstraffic

    
    Csipstatsretry CISCOSIPUAMIB_Csipstatsretry

    
    Csipstatsmisc CISCOSIPUAMIB_Csipstatsmisc

    
    Csipstatsconnection CISCOSIPUAMIB_Csipstatsconnection

    // This table contains configuration for Early Media Cut Through.  The
    // configuration controls how the SIP user agent will process 1xx
    // (Provisional) SIP response messages that contain  Session Definition
    // Protocol (SDP) payloads.
    Csipcfgearlymediatable CISCOSIPUAMIB_Csipcfgearlymediatable

    // This table contains configuration for binding the scope of packets to the
    // particular ethernet interface. The scope for the packets can be specified
    // as either 'signalling' or 'media' packets. The ethernet interface shall be
    // specified by the interface index. The table shall be indexed based on the
    // scope.
    Csipcfgbindsourceaddrtable CISCOSIPUAMIB_Csipcfgbindsourceaddrtable

    // This table contains per dial-peer SIP related  configuration.     The table
    // is a sparse table of dial-peer information. This means, it only reflects
    // dial-peers being used  for SIP.  A dial-peer is being used for SIP if the 
    // value of cvVoIPPeerCfgSessionProtocol  (CISCO-VOICE-DIAL-CONTROL-MIB) is
    // 'sip'.  Dial-peers are not created or destroyed via this table.  Only SIP
    // related configuration can be  performed via this table once the dial-peer
    // exists in the system and is visible in this table.
    Csipcfgpeertable CISCOSIPUAMIB_Csipcfgpeertable

    // This table contains SIP status code to PSTN cause code mapping
    // configuration.  Inbound SIP response messages  that will result in outbound
    // PSTN signalling messages will have the SIP status codes transposed into the
    // PSTN cause codes as prescribed by the contents of this  table.
    Csipcfgstatuscausetable CISCOSIPUAMIB_Csipcfgstatuscausetable

    // This table contains PSTN cause code to SIP status code mapping
    // configuration.   Inbound PSTN signalling messages that will result in
    // outbound SIP response messages  will have the PSTN cause codes transposed
    // into the SIP status codes as prescribed by the contents of this  table.
    Csipcfgcausestatustable CISCOSIPUAMIB_Csipcfgcausestatustable

    // This table contains statistics for sent and received 200 Ok response
    // messages.  The  statistics are kept on per SIP method basis.
    Csipstatssuccessoktable CISCOSIPUAMIB_Csipstatssuccessoktable
}

func (cISCOSIPUAMIB *CISCOSIPUAMIB) GetEntityData() *types.CommonEntityData {
    cISCOSIPUAMIB.EntityData.YFilter = cISCOSIPUAMIB.YFilter
    cISCOSIPUAMIB.EntityData.YangName = "CISCO-SIP-UA-MIB"
    cISCOSIPUAMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOSIPUAMIB.EntityData.ParentYangName = "CISCO-SIP-UA-MIB"
    cISCOSIPUAMIB.EntityData.SegmentPath = "CISCO-SIP-UA-MIB:CISCO-SIP-UA-MIB"
    cISCOSIPUAMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOSIPUAMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOSIPUAMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOSIPUAMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOSIPUAMIB.EntityData.Children["cSipCfgBase"] = types.YChild{"Csipcfgbase", &cISCOSIPUAMIB.Csipcfgbase}
    cISCOSIPUAMIB.EntityData.Children["cSipCfgTimer"] = types.YChild{"Csipcfgtimer", &cISCOSIPUAMIB.Csipcfgtimer}
    cISCOSIPUAMIB.EntityData.Children["cSipCfgRetry"] = types.YChild{"Csipcfgretry", &cISCOSIPUAMIB.Csipcfgretry}
    cISCOSIPUAMIB.EntityData.Children["cSipCfgPeer"] = types.YChild{"Csipcfgpeer", &cISCOSIPUAMIB.Csipcfgpeer}
    cISCOSIPUAMIB.EntityData.Children["cSipCfgAaa"] = types.YChild{"Csipcfgaaa", &cISCOSIPUAMIB.Csipcfgaaa}
    cISCOSIPUAMIB.EntityData.Children["cSipCfgVoiceServiceVoip"] = types.YChild{"Csipcfgvoiceservicevoip", &cISCOSIPUAMIB.Csipcfgvoiceservicevoip}
    cISCOSIPUAMIB.EntityData.Children["cSipStatsInfo"] = types.YChild{"Csipstatsinfo", &cISCOSIPUAMIB.Csipstatsinfo}
    cISCOSIPUAMIB.EntityData.Children["cSipStatsSuccess"] = types.YChild{"Csipstatssuccess", &cISCOSIPUAMIB.Csipstatssuccess}
    cISCOSIPUAMIB.EntityData.Children["cSipStatsRedirect"] = types.YChild{"Csipstatsredirect", &cISCOSIPUAMIB.Csipstatsredirect}
    cISCOSIPUAMIB.EntityData.Children["cSipStatsErrClient"] = types.YChild{"Csipstatserrclient", &cISCOSIPUAMIB.Csipstatserrclient}
    cISCOSIPUAMIB.EntityData.Children["cSipStatsErrServer"] = types.YChild{"Csipstatserrserver", &cISCOSIPUAMIB.Csipstatserrserver}
    cISCOSIPUAMIB.EntityData.Children["cSipStatsGlobalFail"] = types.YChild{"Csipstatsglobalfail", &cISCOSIPUAMIB.Csipstatsglobalfail}
    cISCOSIPUAMIB.EntityData.Children["cSipStatsTraffic"] = types.YChild{"Csipstatstraffic", &cISCOSIPUAMIB.Csipstatstraffic}
    cISCOSIPUAMIB.EntityData.Children["cSipStatsRetry"] = types.YChild{"Csipstatsretry", &cISCOSIPUAMIB.Csipstatsretry}
    cISCOSIPUAMIB.EntityData.Children["cSipStatsMisc"] = types.YChild{"Csipstatsmisc", &cISCOSIPUAMIB.Csipstatsmisc}
    cISCOSIPUAMIB.EntityData.Children["cSipStatsConnection"] = types.YChild{"Csipstatsconnection", &cISCOSIPUAMIB.Csipstatsconnection}
    cISCOSIPUAMIB.EntityData.Children["cSipCfgEarlyMediaTable"] = types.YChild{"Csipcfgearlymediatable", &cISCOSIPUAMIB.Csipcfgearlymediatable}
    cISCOSIPUAMIB.EntityData.Children["cSipCfgBindSourceAddrTable"] = types.YChild{"Csipcfgbindsourceaddrtable", &cISCOSIPUAMIB.Csipcfgbindsourceaddrtable}
    cISCOSIPUAMIB.EntityData.Children["cSipCfgPeerTable"] = types.YChild{"Csipcfgpeertable", &cISCOSIPUAMIB.Csipcfgpeertable}
    cISCOSIPUAMIB.EntityData.Children["cSipCfgStatusCauseTable"] = types.YChild{"Csipcfgstatuscausetable", &cISCOSIPUAMIB.Csipcfgstatuscausetable}
    cISCOSIPUAMIB.EntityData.Children["cSipCfgCauseStatusTable"] = types.YChild{"Csipcfgcausestatustable", &cISCOSIPUAMIB.Csipcfgcausestatustable}
    cISCOSIPUAMIB.EntityData.Children["cSipStatsSuccessOkTable"] = types.YChild{"Csipstatssuccessoktable", &cISCOSIPUAMIB.Csipstatssuccessoktable}
    cISCOSIPUAMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOSIPUAMIB.EntityData)
}

// CISCOSIPUAMIB_Csipcfgbase
type CISCOSIPUAMIB_Csipcfgbase struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object will reflect the version of SIP supported by this  user agent. 
    // It will follow the same format as SIP version  information contained in the
    // SIP messages generated by this user agent.  For example, user agents
    // supporting SIP version 2 will return 'SIP/2.0' as dictated by RFC 2543. The
    // type is string.
    Csipcfgversion interface{}

    // This object specifies the transport protocol the SIP user  agent will use
    // to receive SIP messages.  A value of 'disabled' indicates that the UA will
    // not receive any SIP messages. The type is Csipcfgtransport.
    Csipcfgtransport interface{}

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
    Csipcfguserlocationserveraddr interface{}

    // This object may be used with any SIP method to limit the  number of proxies
    // that can forward the request to the next  downstream server. The type is
    // interface{} with range: 1..70.
    Csipcfgmaxforwards interface{}

    // This object may specify the interface where the source IP address used in
    // SIP signalling or media packets is configured.  A value of 0 means that 
    // there is no specific source address configured and  in this case the best
    // local IP address will be chosen  by the system. The type is interface{}
    // with range: 0..2147483647.
    Csipcfgbindsrcaddrinterface interface{}

    // This object specifies the scope of packets to which the source IP address
    // of the interface  designated by cSipCfgBindSrcAddrInterface will be bound. 
    // A value of 'all' means the IP address  will be bound to both SIP signalling
    // and media packets. A value of 'control' means the IP address will only be
    // bound to SIP signalling packets.   If cSipCfgBindSrcAddrInterface is set to
    // 0, the value of this object has no meaning. The type is
    // Csipcfgbindsrcaddrscope.
    Csipcfgbindsrcaddrscope interface{}

    // This object specifies the format of the prefix used  by the system for DNS
    // SRV queries.  v1  :  RFC 2052 format - 'protocol.transport.' v2  :  RFC
    // 2782 format - '_protocol._transport.'  This object allows backward
    // compatibility with systems only supporting RFC 2052 format. The type is
    // Csipcfgdnssrvquerystringformat.
    Csipcfgdnssrvquerystringformat interface{}

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
    Csipcfgredirectiondisabled interface{}

    // This object specifies whether remote media checks for Symmetric Network
    // Address Translation (NAT)  is enabled or disabled.  If 'true', remote media
    // checks are enabled.  The gateway will have the ability to open a Real Time 
    // Transport Protocol (RTP) session with the remote end and then update
    // (modify) the existing RTP session's remote address/port (raddr:rport) with
    // the source address and port of the actual media packet received.  This
    // would be triggered for only those calls where the Session Description
    // Protocol (SDP) received by the gateway has an indication to do so.  If
    // 'false', remote media checks are disabled. The type is bool.
    Csipcfgsymnatenabled interface{}

    // This object specifies the value of the 'a=direction:<role>' SDP attribute
    // supported by  the user agent.  The direction attribute is used  to describe
    // the role of the user agent (as an  endpoint for a connection-oriented media
    // stream)  in the connection setup procedure.  none    :  No role is
    // specified. passive :  The user agent will advertise itself            as a
    // 'passive' entity (inside the NAT)            in the SDP. active  :  The
    // user agent will advertise itself            as a 'active' entity (outside
    // the NAT)            in the SDP. The type is Csipcfgsymnatdirectionrole.
    Csipcfgsymnatdirectionrole interface{}

    // This object specifies if support for handling  Suspend/Resume events from
    // the switch is enabled or not.  If 'true', the user agent on getting a
    // Suspend will notify the remote agent by sending it a re-invite with a hold
    // SDP. Similarly, when the Gateway receives a Resume, it will initiate a
    // re-invite with the original SDP and unhold the call.  If 'false', the user
    // agent will not initiate any re-invites on receiving Suspend/Resume events,
    // basically it won't be putting the call on hold or off hold. The type is
    // bool.
    Csipcfgsuspendresumeenabled interface{}

    // This object specifies how the SIP gateway would initiate call hold
    // requests.  directionAttr: The user agent will use the direction            
    // attribute such as a=sendonly or a=inactive in                 the sdp to
    // initiate call hold requests.                    connAddr: The user agent
    // will use 0.0.0.0 connection address            to specify Call Hold. The
    // type is Csipcfgoffercallhold.
    Csipcfgoffercallhold interface{}

    // This object specifies that the Reason header overrides SIP  status code
    // mapping table. The type is bool.
    Csipcfgreasonheaderoveride interface{}

    // This object may be used with any SIP method to limit the  number of proxies
    // that can forward the request to the next  downstream server. The type is
    // interface{} with range: 1..70.
    Csipcfgmaximumforwards interface{}
}

func (csipcfgbase *CISCOSIPUAMIB_Csipcfgbase) GetEntityData() *types.CommonEntityData {
    csipcfgbase.EntityData.YFilter = csipcfgbase.YFilter
    csipcfgbase.EntityData.YangName = "cSipCfgBase"
    csipcfgbase.EntityData.BundleName = "cisco_ios_xe"
    csipcfgbase.EntityData.ParentYangName = "CISCO-SIP-UA-MIB"
    csipcfgbase.EntityData.SegmentPath = "cSipCfgBase"
    csipcfgbase.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csipcfgbase.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csipcfgbase.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csipcfgbase.EntityData.Children = make(map[string]types.YChild)
    csipcfgbase.EntityData.Leafs = make(map[string]types.YLeaf)
    csipcfgbase.EntityData.Leafs["cSipCfgVersion"] = types.YLeaf{"Csipcfgversion", csipcfgbase.Csipcfgversion}
    csipcfgbase.EntityData.Leafs["cSipCfgTransport"] = types.YLeaf{"Csipcfgtransport", csipcfgbase.Csipcfgtransport}
    csipcfgbase.EntityData.Leafs["cSipCfgUserLocationServerAddr"] = types.YLeaf{"Csipcfguserlocationserveraddr", csipcfgbase.Csipcfguserlocationserveraddr}
    csipcfgbase.EntityData.Leafs["cSipCfgMaxForwards"] = types.YLeaf{"Csipcfgmaxforwards", csipcfgbase.Csipcfgmaxforwards}
    csipcfgbase.EntityData.Leafs["cSipCfgBindSrcAddrInterface"] = types.YLeaf{"Csipcfgbindsrcaddrinterface", csipcfgbase.Csipcfgbindsrcaddrinterface}
    csipcfgbase.EntityData.Leafs["cSipCfgBindSrcAddrScope"] = types.YLeaf{"Csipcfgbindsrcaddrscope", csipcfgbase.Csipcfgbindsrcaddrscope}
    csipcfgbase.EntityData.Leafs["cSipCfgDnsSrvQueryStringFormat"] = types.YLeaf{"Csipcfgdnssrvquerystringformat", csipcfgbase.Csipcfgdnssrvquerystringformat}
    csipcfgbase.EntityData.Leafs["cSipCfgRedirectionDisabled"] = types.YLeaf{"Csipcfgredirectiondisabled", csipcfgbase.Csipcfgredirectiondisabled}
    csipcfgbase.EntityData.Leafs["cSipCfgSymNatEnabled"] = types.YLeaf{"Csipcfgsymnatenabled", csipcfgbase.Csipcfgsymnatenabled}
    csipcfgbase.EntityData.Leafs["cSipCfgSymNatDirectionRole"] = types.YLeaf{"Csipcfgsymnatdirectionrole", csipcfgbase.Csipcfgsymnatdirectionrole}
    csipcfgbase.EntityData.Leafs["cSipCfgSuspendResumeEnabled"] = types.YLeaf{"Csipcfgsuspendresumeenabled", csipcfgbase.Csipcfgsuspendresumeenabled}
    csipcfgbase.EntityData.Leafs["cSipCfgOfferCallHold"] = types.YLeaf{"Csipcfgoffercallhold", csipcfgbase.Csipcfgoffercallhold}
    csipcfgbase.EntityData.Leafs["cSipCfgReasonHeaderOveride"] = types.YLeaf{"Csipcfgreasonheaderoveride", csipcfgbase.Csipcfgreasonheaderoveride}
    csipcfgbase.EntityData.Leafs["cSipCfgMaximumForwards"] = types.YLeaf{"Csipcfgmaximumforwards", csipcfgbase.Csipcfgmaximumforwards}
    return &(csipcfgbase.EntityData)
}

// CISCOSIPUAMIB_Csipcfgbase_Csipcfgbindsrcaddrscope represents the value of this object has no meaning.
type CISCOSIPUAMIB_Csipcfgbase_Csipcfgbindsrcaddrscope string

const (
    CISCOSIPUAMIB_Csipcfgbase_Csipcfgbindsrcaddrscope_none CISCOSIPUAMIB_Csipcfgbase_Csipcfgbindsrcaddrscope = "none"

    CISCOSIPUAMIB_Csipcfgbase_Csipcfgbindsrcaddrscope_all CISCOSIPUAMIB_Csipcfgbase_Csipcfgbindsrcaddrscope = "all"

    CISCOSIPUAMIB_Csipcfgbase_Csipcfgbindsrcaddrscope_control CISCOSIPUAMIB_Csipcfgbase_Csipcfgbindsrcaddrscope = "control"
)

// CISCOSIPUAMIB_Csipcfgbase_Csipcfgdnssrvquerystringformat represents only supporting RFC 2052 format.
type CISCOSIPUAMIB_Csipcfgbase_Csipcfgdnssrvquerystringformat string

const (
    CISCOSIPUAMIB_Csipcfgbase_Csipcfgdnssrvquerystringformat_v1 CISCOSIPUAMIB_Csipcfgbase_Csipcfgdnssrvquerystringformat = "v1"

    CISCOSIPUAMIB_Csipcfgbase_Csipcfgdnssrvquerystringformat_v2 CISCOSIPUAMIB_Csipcfgbase_Csipcfgdnssrvquerystringformat = "v2"
)

// CISCOSIPUAMIB_Csipcfgbase_Csipcfgoffercallhold represents            to specify Call Hold.
type CISCOSIPUAMIB_Csipcfgbase_Csipcfgoffercallhold string

const (
    CISCOSIPUAMIB_Csipcfgbase_Csipcfgoffercallhold_directionAttr CISCOSIPUAMIB_Csipcfgbase_Csipcfgoffercallhold = "directionAttr"

    CISCOSIPUAMIB_Csipcfgbase_Csipcfgoffercallhold_connAddr CISCOSIPUAMIB_Csipcfgbase_Csipcfgoffercallhold = "connAddr"
)

// CISCOSIPUAMIB_Csipcfgbase_Csipcfgsymnatdirectionrole represents            in the SDP.
type CISCOSIPUAMIB_Csipcfgbase_Csipcfgsymnatdirectionrole string

const (
    CISCOSIPUAMIB_Csipcfgbase_Csipcfgsymnatdirectionrole_none CISCOSIPUAMIB_Csipcfgbase_Csipcfgsymnatdirectionrole = "none"

    CISCOSIPUAMIB_Csipcfgbase_Csipcfgsymnatdirectionrole_passive CISCOSIPUAMIB_Csipcfgbase_Csipcfgsymnatdirectionrole = "passive"

    CISCOSIPUAMIB_Csipcfgbase_Csipcfgsymnatdirectionrole_active CISCOSIPUAMIB_Csipcfgbase_Csipcfgsymnatdirectionrole = "active"
)

// CISCOSIPUAMIB_Csipcfgbase_Csipcfgtransport represents indicates that the UA will not receive any SIP messages.
type CISCOSIPUAMIB_Csipcfgbase_Csipcfgtransport string

const (
    CISCOSIPUAMIB_Csipcfgbase_Csipcfgtransport_udp CISCOSIPUAMIB_Csipcfgbase_Csipcfgtransport = "udp"

    CISCOSIPUAMIB_Csipcfgbase_Csipcfgtransport_tcp CISCOSIPUAMIB_Csipcfgbase_Csipcfgtransport = "tcp"

    CISCOSIPUAMIB_Csipcfgbase_Csipcfgtransport_udpAndTcp CISCOSIPUAMIB_Csipcfgbase_Csipcfgtransport = "udpAndTcp"

    CISCOSIPUAMIB_Csipcfgbase_Csipcfgtransport_disabled CISCOSIPUAMIB_Csipcfgbase_Csipcfgtransport = "disabled"
)

// CISCOSIPUAMIB_Csipcfgtimer
type CISCOSIPUAMIB_Csipcfgtimer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object specifies the time a user agent will wait to  receive a
    // provisional response to a INVITE before resending  the INVITE. The type is
    // interface{} with range: 100..1000. Units are milliseconds.
    Csipcfgtimertrying interface{}

    // This object specifies the time a user agent will wait to  receive a final
    // response to a INVITE before cancelling the  transaction. The type is
    // interface{} with range: 60000..300000. Units are milliseconds.
    Csipcfgtimerexpires interface{}

    // This object specifies the time a user agent will wait to  receive an ACK
    // confirmation a session is established. The type is interface{} with range:
    // 100..1000. Units are milliseconds.
    Csipcfgtimerconnect interface{}

    // This object specifies the time a user agent will wait to  receive an BYE
    // confirmation a session is disconnected. The type is interface{} with range:
    // 100..1000. Units are milliseconds.
    Csipcfgtimerdisconnect interface{}

    // This object specifies the time a user agent will wait for  a final response
    // before retransmitting the PRACK (PRovisional ACKnowledgment). The type is
    // interface{} with range: 100..1000. Units are milliseconds.
    Csipcfgtimerprack interface{}

    // This object specifies the time a user agent will wait  for a final response
    // before retransmitting the COMET  (COndition MET). The type is interface{}
    // with range: 100..1000. Units are milliseconds.
    Csipcfgtimercomet interface{}

    // This object specifies the amount of time to wait for a PRACK before
    // retransmitting the reliable 1xx response. The type is interface{} with
    // range: 100..1000. Units are milliseconds.
    Csipcfgtimerreliablersp interface{}

    // This object specifies the amount of time to wait for a final response
    // before retransmitting the Notify. The type is interface{} with range:
    // 100..1000. Units are milliseconds.
    Csipcfgtimernotify interface{}

    // This object specifies the amount of time to wait for a final response
    // before retransmitting the Refer. The type is interface{} with range:
    // 100..1000. Units are milliseconds.
    Csipcfgtimerrefer interface{}

    // This object specifies the value of the Min-SE  header for INVITE messages
    // originated by this  user agent and the minimum value of the 
    // Session-Expires headers for INVITE messages  received by this user agent. 
    // Any Session-Expires headers received with a  value below this object's
    // value will be rejected with a 422 client error response message.  Setting
    // this object to a value less than 600 is valid, but the possibly of
    // excessive re-INVITES  and the impact of those messages should be fully 
    // understood and considered an acceptable risk. The type is interface{} with
    // range: 60..86400. Units are seconds.
    Csipcfgtimersessiontimer interface{}

    // This object specifies the amount of time to wait before  disconnecting a
    // call already on hold. A value of 0 specifies that this functionality is
    // disabled. The type is interface{} with range: 0..None | 15..2880. Units are
    // minutes.
    Csipcfgtimerhold interface{}

    // This object specifies the amount of time to wait for a 200ok response
    // before retransmitting the Info. The type is interface{} with range:
    // 100..1000. Units are milliseconds.
    Csipcfgtimerinfo interface{}

    // This object specifies the amount of time to wait before  aging out a
    // TCP/UDP connection. The type is interface{} with range: 5..30. Units are
    // minutes.
    Csipcfgtimerconnectionaging interface{}

    // This object specifies the amount of time to buffer the INVITE  while
    // waiting for display name info in the Facility.  A value of 0 means that the
    // INVITE wouldn't be buffered waiting for the display name info in the
    // Facility. The type is interface{} with range: 0..None | 50..5000. Units are
    // milliseconds.
    Csipcfgtimerbufferinvite interface{}
}

func (csipcfgtimer *CISCOSIPUAMIB_Csipcfgtimer) GetEntityData() *types.CommonEntityData {
    csipcfgtimer.EntityData.YFilter = csipcfgtimer.YFilter
    csipcfgtimer.EntityData.YangName = "cSipCfgTimer"
    csipcfgtimer.EntityData.BundleName = "cisco_ios_xe"
    csipcfgtimer.EntityData.ParentYangName = "CISCO-SIP-UA-MIB"
    csipcfgtimer.EntityData.SegmentPath = "cSipCfgTimer"
    csipcfgtimer.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csipcfgtimer.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csipcfgtimer.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csipcfgtimer.EntityData.Children = make(map[string]types.YChild)
    csipcfgtimer.EntityData.Leafs = make(map[string]types.YLeaf)
    csipcfgtimer.EntityData.Leafs["cSipCfgTimerTrying"] = types.YLeaf{"Csipcfgtimertrying", csipcfgtimer.Csipcfgtimertrying}
    csipcfgtimer.EntityData.Leafs["cSipCfgTimerExpires"] = types.YLeaf{"Csipcfgtimerexpires", csipcfgtimer.Csipcfgtimerexpires}
    csipcfgtimer.EntityData.Leafs["cSipCfgTimerConnect"] = types.YLeaf{"Csipcfgtimerconnect", csipcfgtimer.Csipcfgtimerconnect}
    csipcfgtimer.EntityData.Leafs["cSipCfgTimerDisconnect"] = types.YLeaf{"Csipcfgtimerdisconnect", csipcfgtimer.Csipcfgtimerdisconnect}
    csipcfgtimer.EntityData.Leafs["cSipCfgTimerPrack"] = types.YLeaf{"Csipcfgtimerprack", csipcfgtimer.Csipcfgtimerprack}
    csipcfgtimer.EntityData.Leafs["cSipCfgTimerComet"] = types.YLeaf{"Csipcfgtimercomet", csipcfgtimer.Csipcfgtimercomet}
    csipcfgtimer.EntityData.Leafs["cSipCfgTimerReliableRsp"] = types.YLeaf{"Csipcfgtimerreliablersp", csipcfgtimer.Csipcfgtimerreliablersp}
    csipcfgtimer.EntityData.Leafs["cSipCfgTimerNotify"] = types.YLeaf{"Csipcfgtimernotify", csipcfgtimer.Csipcfgtimernotify}
    csipcfgtimer.EntityData.Leafs["cSipCfgTimerRefer"] = types.YLeaf{"Csipcfgtimerrefer", csipcfgtimer.Csipcfgtimerrefer}
    csipcfgtimer.EntityData.Leafs["cSipCfgTimerSessionTimer"] = types.YLeaf{"Csipcfgtimersessiontimer", csipcfgtimer.Csipcfgtimersessiontimer}
    csipcfgtimer.EntityData.Leafs["cSipCfgTimerHold"] = types.YLeaf{"Csipcfgtimerhold", csipcfgtimer.Csipcfgtimerhold}
    csipcfgtimer.EntityData.Leafs["cSipCfgTimerInfo"] = types.YLeaf{"Csipcfgtimerinfo", csipcfgtimer.Csipcfgtimerinfo}
    csipcfgtimer.EntityData.Leafs["cSipCfgTimerConnectionAging"] = types.YLeaf{"Csipcfgtimerconnectionaging", csipcfgtimer.Csipcfgtimerconnectionaging}
    csipcfgtimer.EntityData.Leafs["cSipCfgTimerBufferInvite"] = types.YLeaf{"Csipcfgtimerbufferinvite", csipcfgtimer.Csipcfgtimerbufferinvite}
    return &(csipcfgtimer.EntityData)
}

// CISCOSIPUAMIB_Csipcfgretry
type CISCOSIPUAMIB_Csipcfgretry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object specifies the number of times a user agent  will retry sending
    // a INVITE request. The type is interface{} with range: 1..10.
    Csipcfgretryinvite interface{}

    // This object specifies the number of times a user agent  will retry sending
    // a BYE request. The type is interface{} with range: 1..10.
    Csipcfgretrybye interface{}

    // This object specifies the number of times a user agent  will retry sending
    // a CANCEL request. The type is interface{} with range: 1..10.
    Csipcfgretrycancel interface{}

    // This object specifies the number of times a user agent  will retry sending
    // a REGISTER request. The type is interface{} with range: 1..10.
    Csipcfgretryregister interface{}

    // This object specifies the number of times a user agent  will retry sending
    // a Response and expecting a ACK. The type is interface{} with range: 1..10.
    Csipcfgretryresponse interface{}

    // This object specifies the number of times a user agent  will retry sending
    // a PRACK (PRovisional ACKnowledgement). The type is interface{} with range:
    // 1..10.
    Csipcfgretryprack interface{}

    // This object specifies the number of times a user agent  will retry sending
    // a COMET (COndition MET). The type is interface{} with range: 1..10.
    Csipcfgretrycomet interface{}

    // This object specifies the number of times a user agent  will retry sending
    // a reliable response. The type is interface{} with range: 1..10.
    Csipcfgretryreliablersp interface{}

    // This object specifies the number of times a user agent  will retry sending
    // a Notify request. The type is interface{} with range: 1..10.
    Csipcfgretrynotify interface{}

    // This object specifies the number of times a user agent  will retry sending
    // a Refer request. The type is interface{} with range: 1..10.
    Csipcfgretryrefer interface{}

    // This object specifies the number of times a user agent will retry sending a
    // Info request. The type is interface{} with range: 1..10.
    Csipcfgretryinfo interface{}

    // This object specifies the number of times a user agent will retry sending a
    // Subscribe request. The type is interface{} with range: 1..10.
    Csipcfgretrysubscribe interface{}
}

func (csipcfgretry *CISCOSIPUAMIB_Csipcfgretry) GetEntityData() *types.CommonEntityData {
    csipcfgretry.EntityData.YFilter = csipcfgretry.YFilter
    csipcfgretry.EntityData.YangName = "cSipCfgRetry"
    csipcfgretry.EntityData.BundleName = "cisco_ios_xe"
    csipcfgretry.EntityData.ParentYangName = "CISCO-SIP-UA-MIB"
    csipcfgretry.EntityData.SegmentPath = "cSipCfgRetry"
    csipcfgretry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csipcfgretry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csipcfgretry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csipcfgretry.EntityData.Children = make(map[string]types.YChild)
    csipcfgretry.EntityData.Leafs = make(map[string]types.YLeaf)
    csipcfgretry.EntityData.Leafs["cSipCfgRetryInvite"] = types.YLeaf{"Csipcfgretryinvite", csipcfgretry.Csipcfgretryinvite}
    csipcfgretry.EntityData.Leafs["cSipCfgRetryBye"] = types.YLeaf{"Csipcfgretrybye", csipcfgretry.Csipcfgretrybye}
    csipcfgretry.EntityData.Leafs["cSipCfgRetryCancel"] = types.YLeaf{"Csipcfgretrycancel", csipcfgretry.Csipcfgretrycancel}
    csipcfgretry.EntityData.Leafs["cSipCfgRetryRegister"] = types.YLeaf{"Csipcfgretryregister", csipcfgretry.Csipcfgretryregister}
    csipcfgretry.EntityData.Leafs["cSipCfgRetryResponse"] = types.YLeaf{"Csipcfgretryresponse", csipcfgretry.Csipcfgretryresponse}
    csipcfgretry.EntityData.Leafs["cSipCfgRetryPrack"] = types.YLeaf{"Csipcfgretryprack", csipcfgretry.Csipcfgretryprack}
    csipcfgretry.EntityData.Leafs["cSipCfgRetryComet"] = types.YLeaf{"Csipcfgretrycomet", csipcfgretry.Csipcfgretrycomet}
    csipcfgretry.EntityData.Leafs["cSipCfgRetryReliableRsp"] = types.YLeaf{"Csipcfgretryreliablersp", csipcfgretry.Csipcfgretryreliablersp}
    csipcfgretry.EntityData.Leafs["cSipCfgRetryNotify"] = types.YLeaf{"Csipcfgretrynotify", csipcfgretry.Csipcfgretrynotify}
    csipcfgretry.EntityData.Leafs["cSipCfgRetryRefer"] = types.YLeaf{"Csipcfgretryrefer", csipcfgretry.Csipcfgretryrefer}
    csipcfgretry.EntityData.Leafs["cSipCfgRetryInfo"] = types.YLeaf{"Csipcfgretryinfo", csipcfgretry.Csipcfgretryinfo}
    csipcfgretry.EntityData.Leafs["cSipCfgRetrySubscribe"] = types.YLeaf{"Csipcfgretrysubscribe", csipcfgretry.Csipcfgretrysubscribe}
    return &(csipcfgretry.EntityData)
}

// CISCOSIPUAMIB_Csipcfgpeer
type CISCOSIPUAMIB_Csipcfgpeer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object specifies the session transport  protocol that will be used for
    // outbound SIP  messages.  This configuration is applicable to all dial-peers
    // in the system having  cSipCfgPeerOutSessionTransport set to 'system'. The
    // type is Csipcfgoutsessiontransport.
    Csipcfgoutsessiontransport interface{}

    // This object specifies the string that will be  placed in either the
    // Supported or Require SIP  header, as specified by cSipCfgReliable1xxRspHdr.
    // The type is string.
    Csipcfgreliable1Xxrspstr interface{}

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
    // headers of outgoing INVITEs. The type is Csipcfgreliable1Xxrsphdr.
    Csipcfgreliable1Xxrsphdr interface{}

    // This object specifies the URL type sent in outbound INVITES generated by
    // this device. The type is Csipcfgurltype.
    Csipcfgurltype interface{}
}

func (csipcfgpeer *CISCOSIPUAMIB_Csipcfgpeer) GetEntityData() *types.CommonEntityData {
    csipcfgpeer.EntityData.YFilter = csipcfgpeer.YFilter
    csipcfgpeer.EntityData.YangName = "cSipCfgPeer"
    csipcfgpeer.EntityData.BundleName = "cisco_ios_xe"
    csipcfgpeer.EntityData.ParentYangName = "CISCO-SIP-UA-MIB"
    csipcfgpeer.EntityData.SegmentPath = "cSipCfgPeer"
    csipcfgpeer.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csipcfgpeer.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csipcfgpeer.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csipcfgpeer.EntityData.Children = make(map[string]types.YChild)
    csipcfgpeer.EntityData.Leafs = make(map[string]types.YLeaf)
    csipcfgpeer.EntityData.Leafs["cSipCfgOutSessionTransport"] = types.YLeaf{"Csipcfgoutsessiontransport", csipcfgpeer.Csipcfgoutsessiontransport}
    csipcfgpeer.EntityData.Leafs["cSipCfgReliable1xxRspStr"] = types.YLeaf{"Csipcfgreliable1Xxrspstr", csipcfgpeer.Csipcfgreliable1Xxrspstr}
    csipcfgpeer.EntityData.Leafs["cSipCfgReliable1xxRspHdr"] = types.YLeaf{"Csipcfgreliable1Xxrsphdr", csipcfgpeer.Csipcfgreliable1Xxrsphdr}
    csipcfgpeer.EntityData.Leafs["cSipCfgUrlType"] = types.YLeaf{"Csipcfgurltype", csipcfgpeer.Csipcfgurltype}
    return &(csipcfgpeer.EntityData)
}

// CISCOSIPUAMIB_Csipcfgpeer_Csipcfgoutsessiontransport represents cSipCfgPeerOutSessionTransport set to 'system'.
type CISCOSIPUAMIB_Csipcfgpeer_Csipcfgoutsessiontransport string

const (
    CISCOSIPUAMIB_Csipcfgpeer_Csipcfgoutsessiontransport_udp CISCOSIPUAMIB_Csipcfgpeer_Csipcfgoutsessiontransport = "udp"

    CISCOSIPUAMIB_Csipcfgpeer_Csipcfgoutsessiontransport_tcp CISCOSIPUAMIB_Csipcfgpeer_Csipcfgoutsessiontransport = "tcp"
)

// CISCOSIPUAMIB_Csipcfgpeer_Csipcfgreliable1Xxrsphdr represents either the Require or Supported headers of outgoing INVITEs.
type CISCOSIPUAMIB_Csipcfgpeer_Csipcfgreliable1Xxrsphdr string

const (
    CISCOSIPUAMIB_Csipcfgpeer_Csipcfgreliable1Xxrsphdr_supported CISCOSIPUAMIB_Csipcfgpeer_Csipcfgreliable1Xxrsphdr = "supported"

    CISCOSIPUAMIB_Csipcfgpeer_Csipcfgreliable1Xxrsphdr_require CISCOSIPUAMIB_Csipcfgpeer_Csipcfgreliable1Xxrsphdr = "require"

    CISCOSIPUAMIB_Csipcfgpeer_Csipcfgreliable1Xxrsphdr_disabled CISCOSIPUAMIB_Csipcfgpeer_Csipcfgreliable1Xxrsphdr = "disabled"
)

// CISCOSIPUAMIB_Csipcfgpeer_Csipcfgurltype represents INVITES generated by this device.
type CISCOSIPUAMIB_Csipcfgpeer_Csipcfgurltype string

const (
    CISCOSIPUAMIB_Csipcfgpeer_Csipcfgurltype_sip CISCOSIPUAMIB_Csipcfgpeer_Csipcfgurltype = "sip"

    CISCOSIPUAMIB_Csipcfgpeer_Csipcfgurltype_tel CISCOSIPUAMIB_Csipcfgpeer_Csipcfgurltype = "tel"
)

// CISCOSIPUAMIB_Csipcfgaaa
type CISCOSIPUAMIB_Csipcfgaaa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object specifies the source of the information used to populate the
    // username attribute of AAA billing records. The type is Csipcfgaaausername.
    Csipcfgaaausername interface{}
}

func (csipcfgaaa *CISCOSIPUAMIB_Csipcfgaaa) GetEntityData() *types.CommonEntityData {
    csipcfgaaa.EntityData.YFilter = csipcfgaaa.YFilter
    csipcfgaaa.EntityData.YangName = "cSipCfgAaa"
    csipcfgaaa.EntityData.BundleName = "cisco_ios_xe"
    csipcfgaaa.EntityData.ParentYangName = "CISCO-SIP-UA-MIB"
    csipcfgaaa.EntityData.SegmentPath = "cSipCfgAaa"
    csipcfgaaa.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csipcfgaaa.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csipcfgaaa.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csipcfgaaa.EntityData.Children = make(map[string]types.YChild)
    csipcfgaaa.EntityData.Leafs = make(map[string]types.YLeaf)
    csipcfgaaa.EntityData.Leafs["cSipCfgAaaUsername"] = types.YLeaf{"Csipcfgaaausername", csipcfgaaa.Csipcfgaaausername}
    return &(csipcfgaaa.EntityData)
}

// CISCOSIPUAMIB_Csipcfgaaa_Csipcfgaaausername represents populate the username attribute of AAA billing records.
type CISCOSIPUAMIB_Csipcfgaaa_Csipcfgaaausername string

const (
    CISCOSIPUAMIB_Csipcfgaaa_Csipcfgaaausername_callingNumber CISCOSIPUAMIB_Csipcfgaaa_Csipcfgaaausername = "callingNumber"

    CISCOSIPUAMIB_Csipcfgaaa_Csipcfgaaausername_proxyAuth CISCOSIPUAMIB_Csipcfgaaa_Csipcfgaaausername = "proxyAuth"
)

// CISCOSIPUAMIB_Csipcfgvoiceservicevoip
type CISCOSIPUAMIB_Csipcfgvoiceservicevoip struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object specifies if support for passing SIP headers from Invite,
    // Subscribe, Notify Request to the application is enabled.  If 'true', the
    // headers received in a message will be passed to the application.  If
    // 'false', the headers received in a message will not be passed to the
    // application. The type is bool.
    Csipcfgheaderpassingenabled interface{}

    // This object specifies the maximum number of concurrent SIP subscriptions a
    // SIP Gateway can accept. The type is interface{} with range: 0..4294967295.
    Csipcfgmaxsubscriptionaccept interface{}

    // This object specifies the maximum number of concurrent SIP subscriptions
    // that a SIP Gateway can originate. Default is Max Dialpeers on platform.
    // Maximum is 2*Max Dialpeers on Platform. The type is interface{} with range:
    // 0..4294967295.
    Csipcfgmaxsubscriptionoriginate interface{}

    // This object specifies if the functionality of switching between transports
    // from udp to tcp if the message size of a Request is greater than 1300 bytes
    // is enabled or not.  This configuration is at the global level, and will
    // only be  considered if there exists no voip dial-peer. The type is bool.
    Csipcfgswitchtransportenabled interface{}
}

func (csipcfgvoiceservicevoip *CISCOSIPUAMIB_Csipcfgvoiceservicevoip) GetEntityData() *types.CommonEntityData {
    csipcfgvoiceservicevoip.EntityData.YFilter = csipcfgvoiceservicevoip.YFilter
    csipcfgvoiceservicevoip.EntityData.YangName = "cSipCfgVoiceServiceVoip"
    csipcfgvoiceservicevoip.EntityData.BundleName = "cisco_ios_xe"
    csipcfgvoiceservicevoip.EntityData.ParentYangName = "CISCO-SIP-UA-MIB"
    csipcfgvoiceservicevoip.EntityData.SegmentPath = "cSipCfgVoiceServiceVoip"
    csipcfgvoiceservicevoip.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csipcfgvoiceservicevoip.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csipcfgvoiceservicevoip.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csipcfgvoiceservicevoip.EntityData.Children = make(map[string]types.YChild)
    csipcfgvoiceservicevoip.EntityData.Leafs = make(map[string]types.YLeaf)
    csipcfgvoiceservicevoip.EntityData.Leafs["cSipCfgHeaderPassingEnabled"] = types.YLeaf{"Csipcfgheaderpassingenabled", csipcfgvoiceservicevoip.Csipcfgheaderpassingenabled}
    csipcfgvoiceservicevoip.EntityData.Leafs["cSipCfgMaxSubscriptionAccept"] = types.YLeaf{"Csipcfgmaxsubscriptionaccept", csipcfgvoiceservicevoip.Csipcfgmaxsubscriptionaccept}
    csipcfgvoiceservicevoip.EntityData.Leafs["cSipCfgMaxSubscriptionOriginate"] = types.YLeaf{"Csipcfgmaxsubscriptionoriginate", csipcfgvoiceservicevoip.Csipcfgmaxsubscriptionoriginate}
    csipcfgvoiceservicevoip.EntityData.Leafs["cSipCfgSwitchTransportEnabled"] = types.YLeaf{"Csipcfgswitchtransportenabled", csipcfgvoiceservicevoip.Csipcfgswitchtransportenabled}
    return &(csipcfgvoiceservicevoip.EntityData)
}

// CISCOSIPUAMIB_Csipstatsinfo
type CISCOSIPUAMIB_Csipstatsinfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object reflects the total number of Trying (100) responses received by
    // the user agent since system startup.   Trying responses indicate that some
    // unspecified action is being taken on behalf of this call, but the user has
    // not yet been located.  Inbound Trying responses indicate that outbound
    // INVITE request  sent out by this system have been received and are
    // processed. The type is interface{} with range: 0..4294967295.
    Csipstatsinfotryingins interface{}

    // This object reflects the total number of Trying (100) responses sent by the
    // user agent since system startup. Trying responses indicate that some
    // unspecified action is being taken on behalf of this call, but the user has
    // not yet been located.  Outbound Trying responses indicate this system is
    // successfully  receiving INVITE requests and processing them on  behalf of
    // the system initiating the INVITE. The type is interface{} with range:
    // 0..4294967295.
    Csipstatsinfotryingouts interface{}

    // This object reflects the total number of Ringing (180) responses received
    // by the user agent since system startup. A inbound Ringing response
    // indicates that the UAS processing a INVITE initiated by this system has 
    // found a possible location where the desired end user  has registered
    // recently and is trying to alert the user. The type is interface{} with
    // range: 0..4294967295.
    Csipstatsinforingingins interface{}

    // This object reflects the total number of Ringing (180) responses sent by
    // the user agent since system startup. A outbound Ringing response indicates
    // that this system has processed an INVITE for a particular end user and
    // found a possible location where that user has registered recently.  The
    // system is trying to alert the end user and is conveying that status to the
    // system that originated the INVITE. The type is interface{} with range:
    // 0..4294967295.
    Csipstatsinforingingouts interface{}

    // This object reflects the total number of Call Is Being Forwarded (181)
    // responses received by the user agent since system startup. A proxy server
    // may use a Forwarded status code to indicate that the call is being
    // forwarded to a different set of destinations.  Inbound Forwarded responses
    // indicate  to this system that forwarding actions are taking place  with
    // regard to calls initiated by this system. The type is interface{} with
    // range: 0..4294967295.
    Csipstatsinfoforwardedins interface{}

    // This object reflects the total number of Call Is Being Forwarded (181)
    // responses sent by the user agent since system startup. A proxy server may
    // use a Forwarded status code to indicate that the call is being forwarded to
    // a different set of destinations.  Outbound Forwarded responses indicate
    // this system is taking some forwarding action for calls and conveying that
    // status to the system that initiated the calls. The type is interface{} with
    // range: 0..4294967295.
    Csipstatsinfoforwardedouts interface{}

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
    Csipstatsinfoqueuedins interface{}

    // This object reflects the total number of Queued (182) responses sent by the
    // user agent since system startup. Outbound Queued responses indicate this
    // system has determined that the called party is temporarily unavailable but
    // the call is not rejected.  Rather  the call is queued until the called
    // party becomes available.  Queued responses messages are sent to the system
    // originating the call request to convey the current status of a queued call.
    // The type is interface{} with range: 0..4294967295.
    Csipstatsinfoqueuedouts interface{}

    // This object reflects the total number of Session Progress (183) responses
    // received by the user agent since system startup. The type is interface{}
    // with range: 0..4294967295.
    Csipstatsinfosessionprogins interface{}

    // This object reflects the total number of Session Progress (183) responses
    // sent by the user agent since system startup. The type is interface{} with
    // range: 0..4294967295.
    Csipstatsinfosessionprogouts interface{}
}

func (csipstatsinfo *CISCOSIPUAMIB_Csipstatsinfo) GetEntityData() *types.CommonEntityData {
    csipstatsinfo.EntityData.YFilter = csipstatsinfo.YFilter
    csipstatsinfo.EntityData.YangName = "cSipStatsInfo"
    csipstatsinfo.EntityData.BundleName = "cisco_ios_xe"
    csipstatsinfo.EntityData.ParentYangName = "CISCO-SIP-UA-MIB"
    csipstatsinfo.EntityData.SegmentPath = "cSipStatsInfo"
    csipstatsinfo.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csipstatsinfo.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csipstatsinfo.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csipstatsinfo.EntityData.Children = make(map[string]types.YChild)
    csipstatsinfo.EntityData.Leafs = make(map[string]types.YLeaf)
    csipstatsinfo.EntityData.Leafs["cSipStatsInfoTryingIns"] = types.YLeaf{"Csipstatsinfotryingins", csipstatsinfo.Csipstatsinfotryingins}
    csipstatsinfo.EntityData.Leafs["cSipStatsInfoTryingOuts"] = types.YLeaf{"Csipstatsinfotryingouts", csipstatsinfo.Csipstatsinfotryingouts}
    csipstatsinfo.EntityData.Leafs["cSipStatsInfoRingingIns"] = types.YLeaf{"Csipstatsinforingingins", csipstatsinfo.Csipstatsinforingingins}
    csipstatsinfo.EntityData.Leafs["cSipStatsInfoRingingOuts"] = types.YLeaf{"Csipstatsinforingingouts", csipstatsinfo.Csipstatsinforingingouts}
    csipstatsinfo.EntityData.Leafs["cSipStatsInfoForwardedIns"] = types.YLeaf{"Csipstatsinfoforwardedins", csipstatsinfo.Csipstatsinfoforwardedins}
    csipstatsinfo.EntityData.Leafs["cSipStatsInfoForwardedOuts"] = types.YLeaf{"Csipstatsinfoforwardedouts", csipstatsinfo.Csipstatsinfoforwardedouts}
    csipstatsinfo.EntityData.Leafs["cSipStatsInfoQueuedIns"] = types.YLeaf{"Csipstatsinfoqueuedins", csipstatsinfo.Csipstatsinfoqueuedins}
    csipstatsinfo.EntityData.Leafs["cSipStatsInfoQueuedOuts"] = types.YLeaf{"Csipstatsinfoqueuedouts", csipstatsinfo.Csipstatsinfoqueuedouts}
    csipstatsinfo.EntityData.Leafs["cSipStatsInfoSessionProgIns"] = types.YLeaf{"Csipstatsinfosessionprogins", csipstatsinfo.Csipstatsinfosessionprogins}
    csipstatsinfo.EntityData.Leafs["cSipStatsInfoSessionProgOuts"] = types.YLeaf{"Csipstatsinfosessionprogouts", csipstatsinfo.Csipstatsinfosessionprogouts}
    return &(csipstatsinfo.EntityData)
}

// CISCOSIPUAMIB_Csipstatssuccess
type CISCOSIPUAMIB_Csipstatssuccess struct {
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
    Csipstatssuccessokins interface{}

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
    Csipstatssuccessokouts interface{}

    // This object reflects the total number of Accepted (202) responses received
    // by the user agent since system startup. The meaning of outbound 202 Ok
    // responses depends on the method used in the associated request. The type is
    // interface{} with range: 0..4294967295.
    Csipstatssuccessacceptedins interface{}

    // This object reflects the total number of Accepted (202) responses sent by
    // the user agent since system startup. The meaning of outbound 202 Ok
    // responses depends on the method used in the associated request. The type is
    // interface{} with range: 0..4294967295.
    Csipstatssuccessacceptedouts interface{}
}

func (csipstatssuccess *CISCOSIPUAMIB_Csipstatssuccess) GetEntityData() *types.CommonEntityData {
    csipstatssuccess.EntityData.YFilter = csipstatssuccess.YFilter
    csipstatssuccess.EntityData.YangName = "cSipStatsSuccess"
    csipstatssuccess.EntityData.BundleName = "cisco_ios_xe"
    csipstatssuccess.EntityData.ParentYangName = "CISCO-SIP-UA-MIB"
    csipstatssuccess.EntityData.SegmentPath = "cSipStatsSuccess"
    csipstatssuccess.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csipstatssuccess.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csipstatssuccess.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csipstatssuccess.EntityData.Children = make(map[string]types.YChild)
    csipstatssuccess.EntityData.Leafs = make(map[string]types.YLeaf)
    csipstatssuccess.EntityData.Leafs["cSipStatsSuccessOkIns"] = types.YLeaf{"Csipstatssuccessokins", csipstatssuccess.Csipstatssuccessokins}
    csipstatssuccess.EntityData.Leafs["cSipStatsSuccessOkOuts"] = types.YLeaf{"Csipstatssuccessokouts", csipstatssuccess.Csipstatssuccessokouts}
    csipstatssuccess.EntityData.Leafs["cSipStatsSuccessAcceptedIns"] = types.YLeaf{"Csipstatssuccessacceptedins", csipstatssuccess.Csipstatssuccessacceptedins}
    csipstatssuccess.EntityData.Leafs["cSipStatsSuccessAcceptedOuts"] = types.YLeaf{"Csipstatssuccessacceptedouts", csipstatssuccess.Csipstatssuccessacceptedouts}
    return &(csipstatssuccess.EntityData)
}

// CISCOSIPUAMIB_Csipstatsredirect
type CISCOSIPUAMIB_Csipstatsredirect struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object reflects the total number of Multiple Choices (300) responses
    // received by the user agent since system startup. Multiple Choices responses
    // indicate that the called party can be reached at several different
    // locations and the server cannot or prefers not to proxy the request. The
    // type is interface{} with range: 0..4294967295.
    Csipstatsredirmultiplechoices interface{}

    // This object reflects the total number of Moved  Permanently (301) responses
    // received by the user agent since system startup. Moved Permanently
    // responses indicate that the called party  can no longer be found at the
    // address offered in the request  and the requesting UAC should retry at the
    // new address given  by the Contact header field of the response. The type is
    // interface{} with range: 0..4294967295.
    Csipstatsredirmovedperms interface{}

    // This object reflects the total number of Moved  Temporarily (302) responses
    // received by the user agent since system startup. Moved Temporarily
    // responses indicate the UAC should retry the request directed at the new
    // address(es) given by the Contact header field of the response. The duration
    // of this redirection can be indicated through the Expires header.  If no
    // explicit expiration time is given, the new address(es) are only valid for
    // this call. The type is interface{} with range: 0..4294967295.
    Csipstatsredirmovedtemps interface{}

    // This object reflects the total number of See Other  (303) responses
    // received by the user agent since system startup. The type is interface{}
    // with range: 0..4294967295.
    Csipstatsredirseeothers interface{}

    // This object reflects the total number of Use Proxy  (305) responses
    // received by the user agent since system startup. See Other responses
    // indicate that requested resources must be accessed through the proxy given
    // by the  Contact header field of the response.  The recipient of this
    // response is expected to repeat this single request via the proxy. The type
    // is interface{} with range: 0..4294967295.
    Csipstatsrediruseproxys interface{}

    // This object reflects the total number of Alternative Service (380)
    // responses received by the user agent since system startup. Alternative
    // Service responses indicate that the call was not successful, but
    // alternative services are possible.  Those alternative services are
    // described in the message body of the response. The type is interface{} with
    // range: 0..4294967295.
    Csipstatsrediraltservices interface{}

    // This object reflects the total number of Moved Temporarily (302) responses
    // received by the user agent since system startup.  Moved Temporarily
    // responses indicate the UAC should retry the request directed at the new
    // address(es)  given by the Contact header field of the response. The
    // duration of this redirection can be indicated through the Expires header. 
    // If no explicit expiration time is given, the new address(es) are only valid
    // for this call. The type is interface{} with range: 0..4294967295.
    Csipstatsredirmovedtempsins interface{}

    // This object reflects the total number of Moved Temporarily (302) responses
    // sent by the user agent since system startup. The type is interface{} with
    // range: 0..4294967295.
    Csipstatsredirmovedtempsouts interface{}
}

func (csipstatsredirect *CISCOSIPUAMIB_Csipstatsredirect) GetEntityData() *types.CommonEntityData {
    csipstatsredirect.EntityData.YFilter = csipstatsredirect.YFilter
    csipstatsredirect.EntityData.YangName = "cSipStatsRedirect"
    csipstatsredirect.EntityData.BundleName = "cisco_ios_xe"
    csipstatsredirect.EntityData.ParentYangName = "CISCO-SIP-UA-MIB"
    csipstatsredirect.EntityData.SegmentPath = "cSipStatsRedirect"
    csipstatsredirect.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csipstatsredirect.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csipstatsredirect.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csipstatsredirect.EntityData.Children = make(map[string]types.YChild)
    csipstatsredirect.EntityData.Leafs = make(map[string]types.YLeaf)
    csipstatsredirect.EntityData.Leafs["cSipStatsRedirMultipleChoices"] = types.YLeaf{"Csipstatsredirmultiplechoices", csipstatsredirect.Csipstatsredirmultiplechoices}
    csipstatsredirect.EntityData.Leafs["cSipStatsRedirMovedPerms"] = types.YLeaf{"Csipstatsredirmovedperms", csipstatsredirect.Csipstatsredirmovedperms}
    csipstatsredirect.EntityData.Leafs["cSipStatsRedirMovedTemps"] = types.YLeaf{"Csipstatsredirmovedtemps", csipstatsredirect.Csipstatsredirmovedtemps}
    csipstatsredirect.EntityData.Leafs["cSipStatsRedirSeeOthers"] = types.YLeaf{"Csipstatsredirseeothers", csipstatsredirect.Csipstatsredirseeothers}
    csipstatsredirect.EntityData.Leafs["cSipStatsRedirUseProxys"] = types.YLeaf{"Csipstatsrediruseproxys", csipstatsredirect.Csipstatsrediruseproxys}
    csipstatsredirect.EntityData.Leafs["cSipStatsRedirAltServices"] = types.YLeaf{"Csipstatsrediraltservices", csipstatsredirect.Csipstatsrediraltservices}
    csipstatsredirect.EntityData.Leafs["cSipStatsRedirMovedTempsIns"] = types.YLeaf{"Csipstatsredirmovedtempsins", csipstatsredirect.Csipstatsredirmovedtempsins}
    csipstatsredirect.EntityData.Leafs["cSipStatsRedirMovedTempsOuts"] = types.YLeaf{"Csipstatsredirmovedtempsouts", csipstatsredirect.Csipstatsredirmovedtempsouts}
    return &(csipstatsredirect.EntityData)
}

// CISCOSIPUAMIB_Csipstatserrclient
type CISCOSIPUAMIB_Csipstatserrclient struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object reflects the total number of Bad Request (400)  responses
    // received by the user agent since system startup. Inbound Bad Request
    // responses indicate that requests issued  by this system could not be
    // understood due to malformed  syntax. The type is interface{} with range:
    // 0..4294967295.
    Csipstatsclientbadrequestins interface{}

    // This object reflects the total number of Bad Request (400)  responses sent
    // by the user agent since system startup. Outbound Bad Request responses
    // indicate that requests  received by this system could not be understood due
    // to  malformed syntax. The type is interface{} with range: 0..4294967295.
    Csipstatsclientbadrequestouts interface{}

    // This object reflects the total number of Unauthorized (401)  responses
    // received by the user agent since system startup.   Inbound Unauthorized
    // responses indicate that requests issued  by this system require user
    // authentication. The type is interface{} with range: 0..4294967295.
    Csipstatsclientunauthorizedins interface{}

    // This object reflects the total number of Unauthorized (401)  responses sent
    // by the user agent since system startup. Outbound Unauthorized responses
    // indicate that requests  received by this system require user
    // authentication. The type is interface{} with range: 0..4294967295.
    Csipstatsclientunauthorizedouts interface{}

    // This object reflects the total number of Payment Required  (402) responses
    // received by the user agent since system startup. The type is interface{}
    // with range: 0..4294967295.
    Csipstatsclientpaymentreqdins interface{}

    // This object reflects the total number of Payment Required  (402) responses
    // sent by the user agent since system startup. The type is interface{} with
    // range: 0..4294967295.
    Csipstatsclientpaymentreqdouts interface{}

    // This object reflects the total number of Forbidden (403)  responses
    // received by the user agent since system startup. Inbound Forbidden
    // responses indicate that requests issued by this system are understood by
    // the server but the server refuses to fulfill the request.  Authorization
    // will not help and the requests should not be repeated. The type is
    // interface{} with range: 0..4294967295.
    Csipstatsclientforbiddenins interface{}

    // This object reflects the total number of Forbidden (403)  responses sent by
    // the user agent since system startup. Outbound Forbidden responses indicate
    // that requests received by this system are understood but this system is
    // refusing to fulfill the requests. The type is interface{} with range:
    // 0..4294967295.
    Csipstatsclientforbiddenouts interface{}

    // This object reflects the total number of Not Found (404)  responses
    // received by the user agent since system startup. Inbound Not Found
    // responses indicate that the called party  does not exist at the domain
    // specified in the Request-URI  or the domain is not handled by the recipient
    // of the request. The type is interface{} with range: 0..4294967295.
    Csipstatsclientnotfoundins interface{}

    // This object reflects the total number of Not Found (404)  responses sent by
    // the user agent since system startup. Outbound Not Found responses indicate
    // that this system knows that the called party does not exist at the domain
    // specified in the Request-URI or the domain is not handled by this system.
    // The type is interface{} with range: 0..4294967295.
    Csipstatsclientnotfoundouts interface{}

    // This object reflects the total number of Method Not Allowed  (405) received
    // responses by the user agent. Inbound Method Not Allowed responses indicate
    // that requests  issued by this system have specified a SIP method in the 
    // Request-Line that is not allowed for the address identified  by the
    // Request-URI. The type is interface{} with range: 0..4294967295.
    Csipstatsclientmethnotallowedins interface{}

    // This object reflects the total number of Method Not Allowed  (405) received
    // sent by the user agent since system startup. Outbound Method Not Allowed
    // responses indicate that requests  received by this system have SIP methods
    // specified in the  Request-Line that are not allowed for the address
    // identified  by the Request-URI. The type is interface{} with range:
    // 0..4294967295.
    Csipstatsclientmethnotallowedouts interface{}

    // This object reflects the total number of Not Acceptable  (406) responses
    // received by the user agent since system startup. Inbound Not Acceptable
    // responses indicate the resources  identified by requests issued by this
    // system cannot generate  responses with content characteristics acceptable
    // to this  system according to the accept headers sent in the requests. The
    // type is interface{} with range: 0..4294967295.
    Csipstatsclientnotacceptableins interface{}

    // This object reflects the total number of Not Acceptable (406)  responses
    // sent by the user agent since system startup. Outbound Not Acceptable
    // responses indicate that the resources identified by requests received by
    // this system cannot generate responses with content characteristics
    // acceptable to the  system sending the requests. The type is interface{}
    // with range: 0..4294967295.
    Csipstatsclientnotacceptableouts interface{}

    // This object reflects the total number of Proxy Authentication  Required
    // (407) responses received by the user agent since system startup. Inbound
    // Proxy Authentication Required responses indicate that  this system must
    // authenticate itself with the proxy before  gaining access to the requested
    // resource. The type is interface{} with range: 0..4294967295.
    Csipstatsclientproxyauthreqdins interface{}

    // This object reflects the total number of Proxy Authentication  Required
    // (407) responses sent by the user agent since system startup. Outbound Proxy
    // Authentication Required responses indicate that the systems issuing
    // requests being processed by this system  must authenticate themselves with
    // this system before gaining  access to requested resources. The type is
    // interface{} with range: 0..4294967295.
    Csipstatsclientproxyauthreqdouts interface{}

    // This object reflects the total number of Request Timeout  (408) responses
    // received by the user agent since system startup. Inbound Request Timeout
    // responses indicate that requests  issued by this system are not being
    // processed by the server  within the time indicated in the Expires header of
    // the  request. The type is interface{} with range: 0..4294967295.
    Csipstatsclientreqtimeoutins interface{}

    // This object reflects the total number of Request Timeout  (408) responses
    // sent by the user agent since system startup. Outbound Request Timeout
    // responses indicate that this  system is not able to produce an appropriate
    // response within  the time indicated in the Expires header of the request.
    // The type is interface{} with range: 0..4294967295.
    Csipstatsclientreqtimeoutouts interface{}

    // This object reflects the total number of Conflict (409)  responses received
    // by the user agent since system startup. Inbound Conflict responses indicate
    // that requests issued by this system could not be completed due to a
    // conflict with the current state of a requested resource. The type is
    // interface{} with range: 0..4294967295.
    Csipstatsclientconflictins interface{}

    // This object reflects the total number of Conflict (409)  responses sent by
    // the user agent since system startup. Outbound Conflict responses indicate
    // that requests received by this system could not be completed due to a
    // conflict with the current state of a requested resource. The type is
    // interface{} with range: 0..4294967295.
    Csipstatsclientconflictouts interface{}

    // This object reflects the total number of Gone (410)  responses received by
    // the user agent since system startup. Inbound Gone responses indicate that
    // resources requested by this system are no longer available at the recipient
    // server and no forwarding address is known. The type is interface{} with
    // range: 0..4294967295.
    Csipstatsclientgoneins interface{}

    // This object reflects the total number of Gone (410)  responses sent by the
    // user agent since system startup. Outbound Gone responses indicate that the
    // requested resources are no longer available at this system and no
    // forwarding address is known. The type is interface{} with range:
    // 0..4294967295.
    Csipstatsclientgoneouts interface{}

    // This object reflects the total number of Length Required  (411) responses
    // received by the user agent since system startup. Inbound Length Required
    // responses indicate that requests  issued by this system are being refused
    // by servers because  of no defined Content-Length header field. The type is
    // interface{} with range: 0..4294967295.
    Csipstatsclientlengthrequiredins interface{}

    // This object reflects the total number of Length Required  (411) responses
    // sent by the user agent since system startup. Outbound Length Required
    // responses indicate that requests  received by this system are being refused
    // because of no  defined Content-Length header field. The type is interface{}
    // with range: 0..4294967295.
    Csipstatsclientlengthrequiredouts interface{}

    // This object reflects the total number of Request Entity Too  Large (413)
    // responses received by the user agent since system startup. Inbound Request
    // Entity Too Large responses indicate that  requests issued by this system
    // are being refused because  the request is larger than the server is willing
    // or able to  process. The type is interface{} with range: 0..4294967295.
    Csipstatsclientreqenttoolargeins interface{}

    // This object reflects the total number of Request Entity Too  Large (413)
    // responses sent by the user agent since system startup. Outbound Request
    // Entity Too Large responses indicate that  requests received by this system
    // are larger than this system  is willing or able to process. The type is
    // interface{} with range: 0..4294967295.
    Csipstatsclientreqenttoolargeouts interface{}

    // This object reflects the total number of Request-URI Too  Large (414)
    // responses received by the user agent since system startup. Inbound
    // Request-URI Too Large responses indicate that  requests issued by this
    // system are being refused because the  Request-URI is longer than the server
    // is willing or able to  interpret. The type is interface{} with range:
    // 0..4294967295.
    Csipstatsclientrequritoolargeins interface{}

    // This object reflects the total number of Request-URI Too  Large (414)
    // responses sent by the user agent since system startup. Outbound Request-URI
    // Too Large responses indicate that  Request-URIs received by this system are
    // longer than this  system is willing or able to interpret. The type is
    // interface{} with range: 0..4294967295.
    Csipstatsclientrequritoolargeouts interface{}

    // This object reflects the total number of Unsupported Media  Type (415)
    // responses received by the user agent since system startup. Inbound
    // Unsupported Media Type responses indicate that  requests issued by this
    // system are being refused because the  message body of the request is in a
    // format not supported by  the requested resource for the requested method.
    // The type is interface{} with range: 0..4294967295.
    Csipstatsclientnosupmediatypeins interface{}

    // This object reflects the total number of Unsupported Media  Type (415)
    // responses sent by the user agent since system startup. Outbound Unsupported
    // Media Type responses indicate that the  body of requests received by this
    // system are in a format not  supported by the requested resource for the
    // requested method. The type is interface{} with range: 0..4294967295.
    Csipstatsclientnosupmediatypeouts interface{}

    // This object reflects the total number of Bad Extension (420)  responses
    // received by the user agent since system startup. Inbound Bad Extension
    // responses indicate that the recipient  did not understand the protocol
    // extension specified in a  Require header field. The type is interface{}
    // with range: 0..4294967295.
    Csipstatsclientbadextensionins interface{}

    // This object reflects the total number of Bad Extension (420)  responses
    // sent by the user agent since system startup. Outbound Bad Extension
    // responses indicate that this system did not understand the protocol
    // extension specified in a Require header field of requests. The type is
    // interface{} with range: 0..4294967295.
    Csipstatsclientbadextensionouts interface{}

    // This object reflects the total number of Temporarily Not  Available (480)
    // responses received by the user agent since system startup. Inbound
    // Temporarily Not Available responses indicate that  the called party is
    // currently unavailable. The type is interface{} with range: 0..4294967295.
    Csipstatsclienttempnotavailins interface{}

    // This object reflects the total number of Temporarily Not  Available (480)
    // responses sent by the user agent since system startup. Outbound Temporarily
    // Not Available responses indicate that  the called party's end system was
    // contacted successfully but  the called party is currently unavailable. The
    // type is interface{} with range: 0..4294967295.
    Csipstatsclienttempnotavailouts interface{}

    // This object reflects the total number of Call Leg/Transaction  Does Not
    // Exist (481) responses received by the user agent since system startup.
    // Inbound Call Leg/Transaction Does Not Exist responses indicate that either
    // BYE or CANCEL requests issued by this system were  received by a server and
    // not matching call leg or transaction  existed. The type is interface{} with
    // range: 0..4294967295.
    Csipstatsclientcalllegnoexistins interface{}

    // This object reflects the total number of Call Leg/Transaction  Does Not
    // Exist (481) responses sent by the user agent since system startup. Outbound
    // Call Leg/Transaction Does Not Exist responses  indicate that BYE or CANCEL
    // requests have been received by  this system and not call leg or transaction
    // matching that  request exists. The type is interface{} with range:
    // 0..4294967295.
    Csipstatsclientcalllegnoexistouts interface{}

    // This object reflects the total number of Loop Detected (482)  responses
    // received by the user agent since system startup. Inbound Loop Detected
    // responses indicate that requests issued by this system were received at
    // servers and the server found  itself in the Via path more than once. The
    // type is interface{} with range: 0..4294967295.
    Csipstatsclientloopdetectedins interface{}

    // This object reflects the total number of Loop Detected (482)  responses
    // sent by the user agent since system startup. Outbound Loop Detected
    // responses indicate that requests  received by this system contain a Via
    // path with this system  appearing more than once. The type is interface{}
    // with range: 0..4294967295.
    Csipstatsclientloopdetectedouts interface{}

    // This object reflects the total number of Too Many Hops (483)  responses
    // received by the user agent since system startup. Inbound Too Many Hops
    // responses indicate that requests issued by this system contain more Via
    // entries (hops) than allowed by the Max-Forwards header field of the
    // requests. The type is interface{} with range: 0..4294967295.
    Csipstatsclienttoomanyhopsins interface{}

    // This object reflects the total number of Too Many Hops (483)  responses
    // sent by the user agent since system startup. Outbound Too Many Hops
    // responses indicate that requests received by this system contain more Via
    // entries (hops) than are allowed by the Max-Forwards header field of the
    // requests. The type is interface{} with range: 0..4294967295.
    Csipstatsclienttoomanyhopsouts interface{}

    // This object reflects the total number of Address Incomplete  (484)
    // responses received by the user agent since system startup. Inbound Address
    // Incomplete responses indicate that requests  issued by this system had To
    // addresses or Request-URIs that  were incomplete. The type is interface{}
    // with range: 0..4294967295.
    Csipstatsclientaddrincompleteins interface{}

    // This object reflects the total number of Address Incomplete  (484)
    // responses sent by the user agent since system startup. Outbound Address
    // Incomplete responses indicate that requests  received by this system had To
    // addresses or Request-URIs that  were incomplete. The type is interface{}
    // with range: 0..4294967295.
    Csipstatsclientaddrincompleteouts interface{}

    // This object reflects the total number of Ambiguous (485)  responses
    // received by the user agent since system startup. Inbound Ambiguous
    // responses indicate that requests issued by this system provided ambiguous
    // address information. The type is interface{} with range: 0..4294967295.
    Csipstatsclientambiguousins interface{}

    // This object reflects the total number of Ambiguous (485)  responses sent by
    // the user agent since system startup. Outbound Ambiguous responses indicate
    // that requests received by this system contained ambiguous address
    // information. The type is interface{} with range: 0..4294967295.
    Csipstatsclientambiguousouts interface{}

    // This object reflects the total number of Busy Here (486)  responses
    // received by the user agent since system startup. Inbound Busy Here
    // responses indicate that the called party is currently not willing or not
    // able to take additional calls. The type is interface{} with range:
    // 0..4294967295.
    Csipstatsclientbusyhereins interface{}

    // This object reflects the total number of Busy Here (486)  responses sent by
    // the user agent since system startup. Outbound Busy Here responses indicate
    // that the called party's end system was contacted successfully but the
    // called party is currently not willing or able to take  additional calls.
    // The type is interface{} with range: 0..4294967295.
    Csipstatsclientbusyhereouts interface{}

    // This object reflects the total number of Request Terminated  (487)
    // responses received by the user agent since system startup. Request
    // Terminated responses are issued if the original  request was terminated via
    // CANCEL or BYE. The type is interface{} with range: 0..4294967295.
    Csipstatsclientreqtermins interface{}

    // This object reflects the total number of Request Terminated  (487)
    // responses sent by the user agent since system startup. Request Terminated
    // responses are issued if the original  request was terminated via CANCEL or
    // BYE. The type is interface{} with range: 0..4294967295.
    Csipstatsclientreqtermouts interface{}

    // This object reflects the total number of Not Acceptable Here (488)
    // responses received by the user agent since system startup. The response has
    // the same meaning as 606 (Not Acceptable),  but only applies to the specific
    // entity addressed by the  Request-URI and the request may succeed elsewhere.
    // The type is interface{} with range: 0..4294967295.
    Csipstatsclientnoaccepthereins interface{}

    // This object reflects the total number of Not Acceptable Here (488)
    // responses sent by the user agent since system startup. The response has the
    // same meaning as 606 (Not Acceptable),  but only applies to the specific
    // entity addressed by the  Request-URI and the request may succeed elsewhere.
    // The type is interface{} with range: 0..4294967295.
    Csipstatsclientnoaccepthereouts interface{}

    // This object reflects the total number of BadEvent (489)  responses received
    // by the user agent since system startup. The type is interface{} with range:
    // 0..4294967295.
    Csipstatsclientbadeventins interface{}

    // This object reflects the total number of BadEvent (489)  responses sent by
    // the user agent since system startup. The type is interface{} with range:
    // 0..4294967295.
    Csipstatsclientbadeventouts interface{}

    // This object reflects the total number of SessionTimerTooSmall (422)
    // responses received by the user agent since system  startup. The type is
    // interface{} with range: 0..4294967295.
    Csipstatsclientsttoosmallins interface{}

    // This object reflects the total number of SessionTimerTooSmall (422)
    // responses sent by the user agent since system startup. The type is
    // interface{} with range: 0..4294967295.
    Csipstatsclientsttoosmallouts interface{}

    // This object reflects the total number of RequestPending (491) responses
    // received by the user agent since system startup. The type is interface{}
    // with range: 0..4294967295.
    Csipstatsclientreqpendingins interface{}

    // This object reflects the total number of RequestPending (491) responses
    // sent by the user agent since system startup. The type is interface{} with
    // range: 0..4294967295.
    Csipstatsclientreqpendingouts interface{}
}

func (csipstatserrclient *CISCOSIPUAMIB_Csipstatserrclient) GetEntityData() *types.CommonEntityData {
    csipstatserrclient.EntityData.YFilter = csipstatserrclient.YFilter
    csipstatserrclient.EntityData.YangName = "cSipStatsErrClient"
    csipstatserrclient.EntityData.BundleName = "cisco_ios_xe"
    csipstatserrclient.EntityData.ParentYangName = "CISCO-SIP-UA-MIB"
    csipstatserrclient.EntityData.SegmentPath = "cSipStatsErrClient"
    csipstatserrclient.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csipstatserrclient.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csipstatserrclient.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csipstatserrclient.EntityData.Children = make(map[string]types.YChild)
    csipstatserrclient.EntityData.Leafs = make(map[string]types.YLeaf)
    csipstatserrclient.EntityData.Leafs["cSipStatsClientBadRequestIns"] = types.YLeaf{"Csipstatsclientbadrequestins", csipstatserrclient.Csipstatsclientbadrequestins}
    csipstatserrclient.EntityData.Leafs["cSipStatsClientBadRequestOuts"] = types.YLeaf{"Csipstatsclientbadrequestouts", csipstatserrclient.Csipstatsclientbadrequestouts}
    csipstatserrclient.EntityData.Leafs["cSipStatsClientUnauthorizedIns"] = types.YLeaf{"Csipstatsclientunauthorizedins", csipstatserrclient.Csipstatsclientunauthorizedins}
    csipstatserrclient.EntityData.Leafs["cSipStatsClientUnauthorizedOuts"] = types.YLeaf{"Csipstatsclientunauthorizedouts", csipstatserrclient.Csipstatsclientunauthorizedouts}
    csipstatserrclient.EntityData.Leafs["cSipStatsClientPaymentReqdIns"] = types.YLeaf{"Csipstatsclientpaymentreqdins", csipstatserrclient.Csipstatsclientpaymentreqdins}
    csipstatserrclient.EntityData.Leafs["cSipStatsClientPaymentReqdOuts"] = types.YLeaf{"Csipstatsclientpaymentreqdouts", csipstatserrclient.Csipstatsclientpaymentreqdouts}
    csipstatserrclient.EntityData.Leafs["cSipStatsClientForbiddenIns"] = types.YLeaf{"Csipstatsclientforbiddenins", csipstatserrclient.Csipstatsclientforbiddenins}
    csipstatserrclient.EntityData.Leafs["cSipStatsClientForbiddenOuts"] = types.YLeaf{"Csipstatsclientforbiddenouts", csipstatserrclient.Csipstatsclientforbiddenouts}
    csipstatserrclient.EntityData.Leafs["cSipStatsClientNotFoundIns"] = types.YLeaf{"Csipstatsclientnotfoundins", csipstatserrclient.Csipstatsclientnotfoundins}
    csipstatserrclient.EntityData.Leafs["cSipStatsClientNotFoundOuts"] = types.YLeaf{"Csipstatsclientnotfoundouts", csipstatserrclient.Csipstatsclientnotfoundouts}
    csipstatserrclient.EntityData.Leafs["cSipStatsClientMethNotAllowedIns"] = types.YLeaf{"Csipstatsclientmethnotallowedins", csipstatserrclient.Csipstatsclientmethnotallowedins}
    csipstatserrclient.EntityData.Leafs["cSipStatsClientMethNotAllowedOuts"] = types.YLeaf{"Csipstatsclientmethnotallowedouts", csipstatserrclient.Csipstatsclientmethnotallowedouts}
    csipstatserrclient.EntityData.Leafs["cSipStatsClientNotAcceptableIns"] = types.YLeaf{"Csipstatsclientnotacceptableins", csipstatserrclient.Csipstatsclientnotacceptableins}
    csipstatserrclient.EntityData.Leafs["cSipStatsClientNotAcceptableOuts"] = types.YLeaf{"Csipstatsclientnotacceptableouts", csipstatserrclient.Csipstatsclientnotacceptableouts}
    csipstatserrclient.EntityData.Leafs["cSipStatsClientProxyAuthReqdIns"] = types.YLeaf{"Csipstatsclientproxyauthreqdins", csipstatserrclient.Csipstatsclientproxyauthreqdins}
    csipstatserrclient.EntityData.Leafs["cSipStatsClientProxyAuthReqdOuts"] = types.YLeaf{"Csipstatsclientproxyauthreqdouts", csipstatserrclient.Csipstatsclientproxyauthreqdouts}
    csipstatserrclient.EntityData.Leafs["cSipStatsClientReqTimeoutIns"] = types.YLeaf{"Csipstatsclientreqtimeoutins", csipstatserrclient.Csipstatsclientreqtimeoutins}
    csipstatserrclient.EntityData.Leafs["cSipStatsClientReqTimeoutOuts"] = types.YLeaf{"Csipstatsclientreqtimeoutouts", csipstatserrclient.Csipstatsclientreqtimeoutouts}
    csipstatserrclient.EntityData.Leafs["cSipStatsClientConflictIns"] = types.YLeaf{"Csipstatsclientconflictins", csipstatserrclient.Csipstatsclientconflictins}
    csipstatserrclient.EntityData.Leafs["cSipStatsClientConflictOuts"] = types.YLeaf{"Csipstatsclientconflictouts", csipstatserrclient.Csipstatsclientconflictouts}
    csipstatserrclient.EntityData.Leafs["cSipStatsClientGoneIns"] = types.YLeaf{"Csipstatsclientgoneins", csipstatserrclient.Csipstatsclientgoneins}
    csipstatserrclient.EntityData.Leafs["cSipStatsClientGoneOuts"] = types.YLeaf{"Csipstatsclientgoneouts", csipstatserrclient.Csipstatsclientgoneouts}
    csipstatserrclient.EntityData.Leafs["cSipStatsClientLengthRequiredIns"] = types.YLeaf{"Csipstatsclientlengthrequiredins", csipstatserrclient.Csipstatsclientlengthrequiredins}
    csipstatserrclient.EntityData.Leafs["cSipStatsClientLengthRequiredOuts"] = types.YLeaf{"Csipstatsclientlengthrequiredouts", csipstatserrclient.Csipstatsclientlengthrequiredouts}
    csipstatserrclient.EntityData.Leafs["cSipStatsClientReqEntTooLargeIns"] = types.YLeaf{"Csipstatsclientreqenttoolargeins", csipstatserrclient.Csipstatsclientreqenttoolargeins}
    csipstatserrclient.EntityData.Leafs["cSipStatsClientReqEntTooLargeOuts"] = types.YLeaf{"Csipstatsclientreqenttoolargeouts", csipstatserrclient.Csipstatsclientreqenttoolargeouts}
    csipstatserrclient.EntityData.Leafs["cSipStatsClientReqURITooLargeIns"] = types.YLeaf{"Csipstatsclientrequritoolargeins", csipstatserrclient.Csipstatsclientrequritoolargeins}
    csipstatserrclient.EntityData.Leafs["cSipStatsClientReqURITooLargeOuts"] = types.YLeaf{"Csipstatsclientrequritoolargeouts", csipstatserrclient.Csipstatsclientrequritoolargeouts}
    csipstatserrclient.EntityData.Leafs["cSipStatsClientNoSupMediaTypeIns"] = types.YLeaf{"Csipstatsclientnosupmediatypeins", csipstatserrclient.Csipstatsclientnosupmediatypeins}
    csipstatserrclient.EntityData.Leafs["cSipStatsClientNoSupMediaTypeOuts"] = types.YLeaf{"Csipstatsclientnosupmediatypeouts", csipstatserrclient.Csipstatsclientnosupmediatypeouts}
    csipstatserrclient.EntityData.Leafs["cSipStatsClientBadExtensionIns"] = types.YLeaf{"Csipstatsclientbadextensionins", csipstatserrclient.Csipstatsclientbadextensionins}
    csipstatserrclient.EntityData.Leafs["cSipStatsClientBadExtensionOuts"] = types.YLeaf{"Csipstatsclientbadextensionouts", csipstatserrclient.Csipstatsclientbadextensionouts}
    csipstatserrclient.EntityData.Leafs["cSipStatsClientTempNotAvailIns"] = types.YLeaf{"Csipstatsclienttempnotavailins", csipstatserrclient.Csipstatsclienttempnotavailins}
    csipstatserrclient.EntityData.Leafs["cSipStatsClientTempNotAvailOuts"] = types.YLeaf{"Csipstatsclienttempnotavailouts", csipstatserrclient.Csipstatsclienttempnotavailouts}
    csipstatserrclient.EntityData.Leafs["cSipStatsClientCallLegNoExistIns"] = types.YLeaf{"Csipstatsclientcalllegnoexistins", csipstatserrclient.Csipstatsclientcalllegnoexistins}
    csipstatserrclient.EntityData.Leafs["cSipStatsClientCallLegNoExistOuts"] = types.YLeaf{"Csipstatsclientcalllegnoexistouts", csipstatserrclient.Csipstatsclientcalllegnoexistouts}
    csipstatserrclient.EntityData.Leafs["cSipStatsClientLoopDetectedIns"] = types.YLeaf{"Csipstatsclientloopdetectedins", csipstatserrclient.Csipstatsclientloopdetectedins}
    csipstatserrclient.EntityData.Leafs["cSipStatsClientLoopDetectedOuts"] = types.YLeaf{"Csipstatsclientloopdetectedouts", csipstatserrclient.Csipstatsclientloopdetectedouts}
    csipstatserrclient.EntityData.Leafs["cSipStatsClientTooManyHopsIns"] = types.YLeaf{"Csipstatsclienttoomanyhopsins", csipstatserrclient.Csipstatsclienttoomanyhopsins}
    csipstatserrclient.EntityData.Leafs["cSipStatsClientTooManyHopsOuts"] = types.YLeaf{"Csipstatsclienttoomanyhopsouts", csipstatserrclient.Csipstatsclienttoomanyhopsouts}
    csipstatserrclient.EntityData.Leafs["cSipStatsClientAddrIncompleteIns"] = types.YLeaf{"Csipstatsclientaddrincompleteins", csipstatserrclient.Csipstatsclientaddrincompleteins}
    csipstatserrclient.EntityData.Leafs["cSipStatsClientAddrIncompleteOuts"] = types.YLeaf{"Csipstatsclientaddrincompleteouts", csipstatserrclient.Csipstatsclientaddrincompleteouts}
    csipstatserrclient.EntityData.Leafs["cSipStatsClientAmbiguousIns"] = types.YLeaf{"Csipstatsclientambiguousins", csipstatserrclient.Csipstatsclientambiguousins}
    csipstatserrclient.EntityData.Leafs["cSipStatsClientAmbiguousOuts"] = types.YLeaf{"Csipstatsclientambiguousouts", csipstatserrclient.Csipstatsclientambiguousouts}
    csipstatserrclient.EntityData.Leafs["cSipStatsClientBusyHereIns"] = types.YLeaf{"Csipstatsclientbusyhereins", csipstatserrclient.Csipstatsclientbusyhereins}
    csipstatserrclient.EntityData.Leafs["cSipStatsClientBusyHereOuts"] = types.YLeaf{"Csipstatsclientbusyhereouts", csipstatserrclient.Csipstatsclientbusyhereouts}
    csipstatserrclient.EntityData.Leafs["cSipStatsClientReqTermIns"] = types.YLeaf{"Csipstatsclientreqtermins", csipstatserrclient.Csipstatsclientreqtermins}
    csipstatserrclient.EntityData.Leafs["cSipStatsClientReqTermOuts"] = types.YLeaf{"Csipstatsclientreqtermouts", csipstatserrclient.Csipstatsclientreqtermouts}
    csipstatserrclient.EntityData.Leafs["cSipStatsClientNoAcceptHereIns"] = types.YLeaf{"Csipstatsclientnoaccepthereins", csipstatserrclient.Csipstatsclientnoaccepthereins}
    csipstatserrclient.EntityData.Leafs["cSipStatsClientNoAcceptHereOuts"] = types.YLeaf{"Csipstatsclientnoaccepthereouts", csipstatserrclient.Csipstatsclientnoaccepthereouts}
    csipstatserrclient.EntityData.Leafs["cSipStatsClientBadEventIns"] = types.YLeaf{"Csipstatsclientbadeventins", csipstatserrclient.Csipstatsclientbadeventins}
    csipstatserrclient.EntityData.Leafs["cSipStatsClientBadEventOuts"] = types.YLeaf{"Csipstatsclientbadeventouts", csipstatserrclient.Csipstatsclientbadeventouts}
    csipstatserrclient.EntityData.Leafs["cSipStatsClientSTTooSmallIns"] = types.YLeaf{"Csipstatsclientsttoosmallins", csipstatserrclient.Csipstatsclientsttoosmallins}
    csipstatserrclient.EntityData.Leafs["cSipStatsClientSTTooSmallOuts"] = types.YLeaf{"Csipstatsclientsttoosmallouts", csipstatserrclient.Csipstatsclientsttoosmallouts}
    csipstatserrclient.EntityData.Leafs["cSipStatsClientReqPendingIns"] = types.YLeaf{"Csipstatsclientreqpendingins", csipstatserrclient.Csipstatsclientreqpendingins}
    csipstatserrclient.EntityData.Leafs["cSipStatsClientReqPendingOuts"] = types.YLeaf{"Csipstatsclientreqpendingouts", csipstatserrclient.Csipstatsclientreqpendingouts}
    return &(csipstatserrclient.EntityData)
}

// CISCOSIPUAMIB_Csipstatserrserver
type CISCOSIPUAMIB_Csipstatserrserver struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object reflects the total number of Internal Server Error (500)
    // responses received by the user agent since system startup. Inbound Internal
    // Server Error responses indicate that servers  to which this system is
    // sending requests have encountered  unexpected conditions that prevent them
    // from fulfilling the  requests. The type is interface{} with range:
    // 0..4294967295.
    Csipstatsserverinterrorins interface{}

    // This object reflects the total number of Internal Server Error (500)
    // responses sent by the user agent since system startup. Outbound Internal
    // Server Error responses indicate that this  system has encountered
    // unexpected conditions that prevent it  from fulfilling received requests.
    // The type is interface{} with range: 0..4294967295.
    Csipstatsserverinterrorouts interface{}

    // This object reflects the total number of Not Implemented  (501) responses
    // received by the user agent since system startup. Inbound Not Implemented
    // responses indicate that servers to  which this system is sending requests
    // do not support the  functionality required to fulfill the requests. The
    // type is interface{} with range: 0..4294967295.
    Csipstatsservernotimplementedins interface{}

    // This object reflects the total number of Not Implemented  (501) responses
    // sent by the user agent since system startup. Outbound Not Implemented
    // responses indicate that this system does not support the functionality
    // required to fulfill the  requests. The type is interface{} with range:
    // 0..4294967295.
    Csipstatsservernotimplementedouts interface{}

    // This object reflects the total number of Bad Gateway (502)  responses
    // received by the user agent since system startup. The type is interface{}
    // with range: 0..4294967295.
    Csipstatsserverbadgatewayins interface{}

    // This object reflects the total number of Bad Gateway (502)  responses sent
    // by the user agent since system startup. The type is interface{} with range:
    // 0..4294967295.
    Csipstatsserverbadgatewayouts interface{}

    // This object reflects the total number of Service Unavailable  (503)
    // responses received by the user agent since system startup. Inbound Service
    // Unavailable responses indicate that the server servicing this system's
    // request is temporarily unavailable to handle the request. The type is
    // interface{} with range: 0..4294967295.
    Csipstatsserverserviceunavailins interface{}

    // This object reflects the total number of Service Unavailable  (503)
    // responses sent by the user agent since system startup. Outbound Service
    // Unavailable responses indicate that this system is temporarily unable to
    // handle received requests. The type is interface{} with range:
    // 0..4294967295.
    Csipstatsserverserviceunavailouts interface{}

    // This object reflects the total number of Gateway Time-out  (504) responses
    // received by the user agent since system startup. Inbound Gateway Time-out
    // responses indicate that the server attempting to complete this system's
    // request did not receive a timely response from yet another system it was
    // accessing to complete the request. The type is interface{} with range:
    // 0..4294967295.
    Csipstatsservergatewaytimeoutins interface{}

    // This object reflects the total number of Gateway Time-out  (504) responses
    // sent by the user agent since system startup. Outbound Gateway Time-out
    // responses indicate that this system did not receive a timely response from
    // the system it had accessed to assist in completing a received request. The
    // type is interface{} with range: 0..4294967295.
    Csipstatsservergatewaytimeoutouts interface{}

    // This object reflects the total number of SIP Version Not  Supported (505)
    // responses received by the user agent since system startup. Inbound SIP
    // Version Not Supported responses indicate that  the server does not support,
    // or refuses to support, the SIP  protocol version that was used in the
    // request message. The type is interface{} with range: 0..4294967295.
    Csipstatsserverbadsipversionins interface{}

    // This object reflects the total number of SIP Version Not  Supported (505)
    // responses sent by the user agent since system startup. Outbound SIP Version
    // Not Supported responses indicate that  this system does not support, or
    // refuses to support, the SIP  protocol version used in received requests.
    // The type is interface{} with range: 0..4294967295.
    Csipstatsserverbadsipversionouts interface{}

    // This object reflects the total number of Precondition  Failure (580)
    // responses received by the user agent since system startup. This response is
    // returned by a UAS if it is unable to perform the mandatory preconditions
    // for the session. The type is interface{} with range: 0..4294967295.
    Csipstatsserverprecondfailureins interface{}

    // This object reflects the total number of Precondition  Failure (580)
    // responses sent by the user agent since system startup. This response is
    // returned by a UAS if it is unable to perform the mandatory preconditions
    // for the session. The type is interface{} with range: 0..4294967295.
    Csipstatsserverprecondfailureouts interface{}
}

func (csipstatserrserver *CISCOSIPUAMIB_Csipstatserrserver) GetEntityData() *types.CommonEntityData {
    csipstatserrserver.EntityData.YFilter = csipstatserrserver.YFilter
    csipstatserrserver.EntityData.YangName = "cSipStatsErrServer"
    csipstatserrserver.EntityData.BundleName = "cisco_ios_xe"
    csipstatserrserver.EntityData.ParentYangName = "CISCO-SIP-UA-MIB"
    csipstatserrserver.EntityData.SegmentPath = "cSipStatsErrServer"
    csipstatserrserver.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csipstatserrserver.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csipstatserrserver.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csipstatserrserver.EntityData.Children = make(map[string]types.YChild)
    csipstatserrserver.EntityData.Leafs = make(map[string]types.YLeaf)
    csipstatserrserver.EntityData.Leafs["cSipStatsServerIntErrorIns"] = types.YLeaf{"Csipstatsserverinterrorins", csipstatserrserver.Csipstatsserverinterrorins}
    csipstatserrserver.EntityData.Leafs["cSipStatsServerIntErrorOuts"] = types.YLeaf{"Csipstatsserverinterrorouts", csipstatserrserver.Csipstatsserverinterrorouts}
    csipstatserrserver.EntityData.Leafs["cSipStatsServerNotImplementedIns"] = types.YLeaf{"Csipstatsservernotimplementedins", csipstatserrserver.Csipstatsservernotimplementedins}
    csipstatserrserver.EntityData.Leafs["cSipStatsServerNotImplementedOuts"] = types.YLeaf{"Csipstatsservernotimplementedouts", csipstatserrserver.Csipstatsservernotimplementedouts}
    csipstatserrserver.EntityData.Leafs["cSipStatsServerBadGatewayIns"] = types.YLeaf{"Csipstatsserverbadgatewayins", csipstatserrserver.Csipstatsserverbadgatewayins}
    csipstatserrserver.EntityData.Leafs["cSipStatsServerBadGatewayOuts"] = types.YLeaf{"Csipstatsserverbadgatewayouts", csipstatserrserver.Csipstatsserverbadgatewayouts}
    csipstatserrserver.EntityData.Leafs["cSipStatsServerServiceUnavailIns"] = types.YLeaf{"Csipstatsserverserviceunavailins", csipstatserrserver.Csipstatsserverserviceunavailins}
    csipstatserrserver.EntityData.Leafs["cSipStatsServerServiceUnavailOuts"] = types.YLeaf{"Csipstatsserverserviceunavailouts", csipstatserrserver.Csipstatsserverserviceunavailouts}
    csipstatserrserver.EntityData.Leafs["cSipStatsServerGatewayTimeoutIns"] = types.YLeaf{"Csipstatsservergatewaytimeoutins", csipstatserrserver.Csipstatsservergatewaytimeoutins}
    csipstatserrserver.EntityData.Leafs["cSipStatsServerGatewayTimeoutOuts"] = types.YLeaf{"Csipstatsservergatewaytimeoutouts", csipstatserrserver.Csipstatsservergatewaytimeoutouts}
    csipstatserrserver.EntityData.Leafs["cSipStatsServerBadSipVersionIns"] = types.YLeaf{"Csipstatsserverbadsipversionins", csipstatserrserver.Csipstatsserverbadsipversionins}
    csipstatserrserver.EntityData.Leafs["cSipStatsServerBadSipVersionOuts"] = types.YLeaf{"Csipstatsserverbadsipversionouts", csipstatserrserver.Csipstatsserverbadsipversionouts}
    csipstatserrserver.EntityData.Leafs["cSipStatsServerPrecondFailureIns"] = types.YLeaf{"Csipstatsserverprecondfailureins", csipstatserrserver.Csipstatsserverprecondfailureins}
    csipstatserrserver.EntityData.Leafs["cSipStatsServerPrecondFailureOuts"] = types.YLeaf{"Csipstatsserverprecondfailureouts", csipstatserrserver.Csipstatsserverprecondfailureouts}
    return &(csipstatserrserver.EntityData)
}

// CISCOSIPUAMIB_Csipstatsglobalfail
type CISCOSIPUAMIB_Csipstatsglobalfail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object reflects the total number of Busy Everywhere (600) responses
    // received by the user agent since system startup. Inbound Busy Everywhere
    // responses indicate that the  called party's end system was contacted
    // successfully but the called party is busy and does not wish to take the
    // call at this time. The type is interface{} with range: 0..4294967295.
    Csipstatsglobalbusyeverywhereins interface{}

    // This object reflects the total number of Busy Everywhere (600) responses
    // sent by the user agent since system startup. Outbound Busy Everywhere
    // responses indicate that  this system has successfully contacted a called
    // party's end system and the called party does not wish to take the call at
    // this time. The type is interface{} with range: 0..4294967295.
    Csipstatsglobalbusyeverywhereouts interface{}

    // This object reflects the total number of Decline (603) responses received
    // by the user agent since system startup. Decline responses indicate that the
    // called party's end  system was contacted successfully but the called party 
    // explicitly does not wish to, or cannot, participate. The type is
    // interface{} with range: 0..4294967295.
    Csipstatsglobaldeclineins interface{}

    // This object reflects the total number of Decline (603) responses sent by
    // the user agent since system startup. Outbound Decline responses indicate
    // that this system has successfully contacted a called party's end system and
    // the called party explicitly does not wish to, or cannot, participate. The
    // type is interface{} with range: 0..4294967295.
    Csipstatsglobaldeclineouts interface{}

    // This object reflects the total number of Does Not Exist Anywhere (604)
    // responses received by the user agent since system startup. Inbound Does Not
    // Exist Anywhere responses indicate that the server handling this system's
    // request has authoritative information that the called party indicated in
    // the To request field does not exist anywhere. The type is interface{} with
    // range: 0..4294967295.
    Csipstatsglobalnotanywhereins interface{}

    // This object reflects the total number of Does Not Exist Anywhere (604)
    // responses sent by the user agent since system startup. Outbound Does Not
    // Exist Anywhere responses indicate that this system has authoritative
    // information that the called party in the To field of received requests does
    // not exist anywhere. The type is interface{} with range: 0..4294967295.
    Csipstatsglobalnotanywhereouts interface{}

    // This object reflects the total number of Not Acceptable (606) responses
    // received by the user agent since system startup. Inbound Not Acceptable
    // responses indicate that the called party's end system was contacted
    // successfully but some aspect of the session description is not acceptable.
    // The type is interface{} with range: 0..4294967295.
    Csipstatsglobalnotacceptableins interface{}

    // This object reflects the total number of Not Acceptable (606) responses
    // sent by the user agent since system startup. Outbound Not Acceptable
    // responses indicate that the called party wishes to communicate, but cannot
    // adequately support the session described in the request. The type is
    // interface{} with range: 0..4294967295.
    Csipstatsglobalnotacceptableouts interface{}
}

func (csipstatsglobalfail *CISCOSIPUAMIB_Csipstatsglobalfail) GetEntityData() *types.CommonEntityData {
    csipstatsglobalfail.EntityData.YFilter = csipstatsglobalfail.YFilter
    csipstatsglobalfail.EntityData.YangName = "cSipStatsGlobalFail"
    csipstatsglobalfail.EntityData.BundleName = "cisco_ios_xe"
    csipstatsglobalfail.EntityData.ParentYangName = "CISCO-SIP-UA-MIB"
    csipstatsglobalfail.EntityData.SegmentPath = "cSipStatsGlobalFail"
    csipstatsglobalfail.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csipstatsglobalfail.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csipstatsglobalfail.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csipstatsglobalfail.EntityData.Children = make(map[string]types.YChild)
    csipstatsglobalfail.EntityData.Leafs = make(map[string]types.YLeaf)
    csipstatsglobalfail.EntityData.Leafs["cSipStatsGlobalBusyEverywhereIns"] = types.YLeaf{"Csipstatsglobalbusyeverywhereins", csipstatsglobalfail.Csipstatsglobalbusyeverywhereins}
    csipstatsglobalfail.EntityData.Leafs["cSipStatsGlobalBusyEverywhereOuts"] = types.YLeaf{"Csipstatsglobalbusyeverywhereouts", csipstatsglobalfail.Csipstatsglobalbusyeverywhereouts}
    csipstatsglobalfail.EntityData.Leafs["cSipStatsGlobalDeclineIns"] = types.YLeaf{"Csipstatsglobaldeclineins", csipstatsglobalfail.Csipstatsglobaldeclineins}
    csipstatsglobalfail.EntityData.Leafs["cSipStatsGlobalDeclineOuts"] = types.YLeaf{"Csipstatsglobaldeclineouts", csipstatsglobalfail.Csipstatsglobaldeclineouts}
    csipstatsglobalfail.EntityData.Leafs["cSipStatsGlobalNotAnywhereIns"] = types.YLeaf{"Csipstatsglobalnotanywhereins", csipstatsglobalfail.Csipstatsglobalnotanywhereins}
    csipstatsglobalfail.EntityData.Leafs["cSipStatsGlobalNotAnywhereOuts"] = types.YLeaf{"Csipstatsglobalnotanywhereouts", csipstatsglobalfail.Csipstatsglobalnotanywhereouts}
    csipstatsglobalfail.EntityData.Leafs["cSipStatsGlobalNotAcceptableIns"] = types.YLeaf{"Csipstatsglobalnotacceptableins", csipstatsglobalfail.Csipstatsglobalnotacceptableins}
    csipstatsglobalfail.EntityData.Leafs["cSipStatsGlobalNotAcceptableOuts"] = types.YLeaf{"Csipstatsglobalnotacceptableouts", csipstatsglobalfail.Csipstatsglobalnotacceptableouts}
    return &(csipstatsglobalfail.EntityData)
}

// CISCOSIPUAMIB_Csipstatstraffic
type CISCOSIPUAMIB_Csipstatstraffic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object reflects the total number of INVITE requests  received by the
    // user agent since system startup. The type is interface{} with range:
    // 0..4294967295.
    Csipstatstrafficinviteins interface{}

    // This object reflects the total number of INVITE requests sent by the user
    // agent. The type is interface{} with range: 0..4294967295.
    Csipstatstrafficinviteouts interface{}

    // This object reflects the total number of ACK requests  received by the user
    // agent since system startup. The type is interface{} with range:
    // 0..4294967295.
    Csipstatstrafficackins interface{}

    // This object reflects the total number of ACK requests sent by the user
    // agent. The type is interface{} with range: 0..4294967295.
    Csipstatstrafficackouts interface{}

    // This object reflects the total number of BYE requests  received by the user
    // agent since system startup. The type is interface{} with range:
    // 0..4294967295.
    Csipstatstrafficbyeins interface{}

    // This object reflects the total number of BYE requests sent by the user
    // agent. The type is interface{} with range: 0..4294967295.
    Csipstatstrafficbyeouts interface{}

    // This object reflects the total number of CANCEL requests  received by the
    // user agent since system startup. The type is interface{} with range:
    // 0..4294967295.
    Csipstatstrafficcancelins interface{}

    // This object reflects the total number of CANCEL requests sent by the user
    // agent. The type is interface{} with range: 0..4294967295.
    Csipstatstrafficcancelouts interface{}

    // This object reflects the total number of OPTIONS requests  received by the
    // user agent since system startup. The type is interface{} with range:
    // 0..4294967295.
    Csipstatstrafficoptionsins interface{}

    // This object reflects the total number of OPTIONS requests sent by the user
    // agent. The type is interface{} with range: 0..4294967295.
    Csipstatstrafficoptionsouts interface{}

    // This object reflects the total number of REGISTER requests  received by the
    // user agent since system startup. The type is interface{} with range:
    // 0..4294967295.
    Csipstatstrafficregisterins interface{}

    // This object reflects the total number of REGISTER requests  sent by the
    // user agent since system startup. The type is interface{} with range:
    // 0..4294967295.
    Csipstatstrafficregisterouts interface{}

    // This object reflects the total number of COndition MET  requests received
    // by the user agent since system startup. The type is interface{} with range:
    // 0..4294967295.
    Csipstatstrafficcometins interface{}

    // This object reflects the total number of COndition MET  requests sent by
    // the user agent since system startup. The type is interface{} with range:
    // 0..4294967295.
    Csipstatstrafficcometouts interface{}

    // This object reflects the total number of PRovisonal  ACKnowledgement
    // requests received by the user agent since system startup. The type is
    // interface{} with range: 0..4294967295.
    Csipstatstrafficprackins interface{}

    // This object reflects the total number of PRovisonal  ACKnowledgement
    // requests sent by the user agent since system startup. The type is
    // interface{} with range: 0..4294967295.
    Csipstatstrafficprackouts interface{}

    // This object reflects the total number of Refer requests received by the
    // user agent since system startup. The type is interface{} with range:
    // 0..4294967295.
    Csipstatstrafficreferins interface{}

    // This object reflects the total number of Refer requests sent by the user
    // agent since system startup. The type is interface{} with range:
    // 0..4294967295.
    Csipstatstrafficreferouts interface{}

    // This object reflects the total number of Notify  requests received by the
    // user agent since system startup. The type is interface{} with range:
    // 0..4294967295.
    Csipstatstrafficnotifyins interface{}

    // This object reflects the total number of Notify  requests sent by the user
    // agent since system startup. The type is interface{} with range:
    // 0..4294967295.
    Csipstatstrafficnotifyouts interface{}

    // This object reflects the total number of Info  requests received by the
    // user agent since system startup. The type is interface{} with range:
    // 0..4294967295.
    Csipstatstrafficinfoins interface{}

    // This object reflects the total number of Info  requests sent by the user
    // agent since system startup. The type is interface{} with range:
    // 0..4294967295.
    Csipstatstrafficinfoouts interface{}

    // This object reflects the total number of Subscribe requests received by the
    // user agent since system startup. The type is interface{} with range:
    // 0..4294967295.
    Csipstatstrafficsubscribeins interface{}

    // This object reflects the total number of Subscribe requests sent by the
    // user agent since system startup. The type is interface{} with range:
    // 0..4294967295.
    Csipstatstrafficsubscribeouts interface{}

    // This object reflects the total number of Update requests received by the
    // user agent since system startup. The type is interface{} with range:
    // 0..4294967295.
    Csipstatstrafficupdateins interface{}

    // This object reflects the total number of Update requests sent by the user
    // agent since system startup. The type is interface{} with range:
    // 0..4294967295.
    Csipstatstrafficupdateouts interface{}
}

func (csipstatstraffic *CISCOSIPUAMIB_Csipstatstraffic) GetEntityData() *types.CommonEntityData {
    csipstatstraffic.EntityData.YFilter = csipstatstraffic.YFilter
    csipstatstraffic.EntityData.YangName = "cSipStatsTraffic"
    csipstatstraffic.EntityData.BundleName = "cisco_ios_xe"
    csipstatstraffic.EntityData.ParentYangName = "CISCO-SIP-UA-MIB"
    csipstatstraffic.EntityData.SegmentPath = "cSipStatsTraffic"
    csipstatstraffic.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csipstatstraffic.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csipstatstraffic.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csipstatstraffic.EntityData.Children = make(map[string]types.YChild)
    csipstatstraffic.EntityData.Leafs = make(map[string]types.YLeaf)
    csipstatstraffic.EntityData.Leafs["cSipStatsTrafficInviteIns"] = types.YLeaf{"Csipstatstrafficinviteins", csipstatstraffic.Csipstatstrafficinviteins}
    csipstatstraffic.EntityData.Leafs["cSipStatsTrafficInviteOuts"] = types.YLeaf{"Csipstatstrafficinviteouts", csipstatstraffic.Csipstatstrafficinviteouts}
    csipstatstraffic.EntityData.Leafs["cSipStatsTrafficAckIns"] = types.YLeaf{"Csipstatstrafficackins", csipstatstraffic.Csipstatstrafficackins}
    csipstatstraffic.EntityData.Leafs["cSipStatsTrafficAckOuts"] = types.YLeaf{"Csipstatstrafficackouts", csipstatstraffic.Csipstatstrafficackouts}
    csipstatstraffic.EntityData.Leafs["cSipStatsTrafficByeIns"] = types.YLeaf{"Csipstatstrafficbyeins", csipstatstraffic.Csipstatstrafficbyeins}
    csipstatstraffic.EntityData.Leafs["cSipStatsTrafficByeOuts"] = types.YLeaf{"Csipstatstrafficbyeouts", csipstatstraffic.Csipstatstrafficbyeouts}
    csipstatstraffic.EntityData.Leafs["cSipStatsTrafficCancelIns"] = types.YLeaf{"Csipstatstrafficcancelins", csipstatstraffic.Csipstatstrafficcancelins}
    csipstatstraffic.EntityData.Leafs["cSipStatsTrafficCancelOuts"] = types.YLeaf{"Csipstatstrafficcancelouts", csipstatstraffic.Csipstatstrafficcancelouts}
    csipstatstraffic.EntityData.Leafs["cSipStatsTrafficOptionsIns"] = types.YLeaf{"Csipstatstrafficoptionsins", csipstatstraffic.Csipstatstrafficoptionsins}
    csipstatstraffic.EntityData.Leafs["cSipStatsTrafficOptionsOuts"] = types.YLeaf{"Csipstatstrafficoptionsouts", csipstatstraffic.Csipstatstrafficoptionsouts}
    csipstatstraffic.EntityData.Leafs["cSipStatsTrafficRegisterIns"] = types.YLeaf{"Csipstatstrafficregisterins", csipstatstraffic.Csipstatstrafficregisterins}
    csipstatstraffic.EntityData.Leafs["cSipStatsTrafficRegisterOuts"] = types.YLeaf{"Csipstatstrafficregisterouts", csipstatstraffic.Csipstatstrafficregisterouts}
    csipstatstraffic.EntityData.Leafs["cSipStatsTrafficCometIns"] = types.YLeaf{"Csipstatstrafficcometins", csipstatstraffic.Csipstatstrafficcometins}
    csipstatstraffic.EntityData.Leafs["cSipStatsTrafficCometOuts"] = types.YLeaf{"Csipstatstrafficcometouts", csipstatstraffic.Csipstatstrafficcometouts}
    csipstatstraffic.EntityData.Leafs["cSipStatsTrafficPrackIns"] = types.YLeaf{"Csipstatstrafficprackins", csipstatstraffic.Csipstatstrafficprackins}
    csipstatstraffic.EntityData.Leafs["cSipStatsTrafficPrackOuts"] = types.YLeaf{"Csipstatstrafficprackouts", csipstatstraffic.Csipstatstrafficprackouts}
    csipstatstraffic.EntityData.Leafs["cSipStatsTrafficReferIns"] = types.YLeaf{"Csipstatstrafficreferins", csipstatstraffic.Csipstatstrafficreferins}
    csipstatstraffic.EntityData.Leafs["cSipStatsTrafficReferOuts"] = types.YLeaf{"Csipstatstrafficreferouts", csipstatstraffic.Csipstatstrafficreferouts}
    csipstatstraffic.EntityData.Leafs["cSipStatsTrafficNotifyIns"] = types.YLeaf{"Csipstatstrafficnotifyins", csipstatstraffic.Csipstatstrafficnotifyins}
    csipstatstraffic.EntityData.Leafs["cSipStatsTrafficNotifyOuts"] = types.YLeaf{"Csipstatstrafficnotifyouts", csipstatstraffic.Csipstatstrafficnotifyouts}
    csipstatstraffic.EntityData.Leafs["cSipStatsTrafficInfoIns"] = types.YLeaf{"Csipstatstrafficinfoins", csipstatstraffic.Csipstatstrafficinfoins}
    csipstatstraffic.EntityData.Leafs["cSipStatsTrafficInfoOuts"] = types.YLeaf{"Csipstatstrafficinfoouts", csipstatstraffic.Csipstatstrafficinfoouts}
    csipstatstraffic.EntityData.Leafs["cSipStatsTrafficSubscribeIns"] = types.YLeaf{"Csipstatstrafficsubscribeins", csipstatstraffic.Csipstatstrafficsubscribeins}
    csipstatstraffic.EntityData.Leafs["cSipStatsTrafficSubscribeOuts"] = types.YLeaf{"Csipstatstrafficsubscribeouts", csipstatstraffic.Csipstatstrafficsubscribeouts}
    csipstatstraffic.EntityData.Leafs["cSipStatsTrafficUpdateIns"] = types.YLeaf{"Csipstatstrafficupdateins", csipstatstraffic.Csipstatstrafficupdateins}
    csipstatstraffic.EntityData.Leafs["cSipStatsTrafficUpdateOuts"] = types.YLeaf{"Csipstatstrafficupdateouts", csipstatstraffic.Csipstatstrafficupdateouts}
    return &(csipstatstraffic.EntityData)
}

// CISCOSIPUAMIB_Csipstatsretry
type CISCOSIPUAMIB_Csipstatsretry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object reflects the total number of INVITE retries that  have been
    // sent by the user agent since system startup.   If the number of 'first 
    // attempt' INVITES is of interest, subtract the value of this  object from
    // cSipStatsTrafficInviteOut. The type is interface{} with range:
    // 0..4294967295.
    Csipstatsretryinvites interface{}

    // This object reflects the total number of BYE retries that have been sent by
    // the user agent since system startup. The type is interface{} with range:
    // 0..4294967295.
    Csipstatsretrybyes interface{}

    // This object reflects the total number of CANCEL retries that  have been
    // sent by the user agent since system startup. The type is interface{} with
    // range: 0..4294967295.
    Csipstatsretrycancels interface{}

    // This object reflects the total number of REGISTER retries that have been
    // sent by the user agent since system startup. The type is interface{} with
    // range: 0..4294967295.
    Csipstatsretryregisters interface{}

    // This object reflects the total number of Response (while  expecting a ACK)
    // retries that have been sent by the user agent. The type is interface{} with
    // range: 0..4294967295.
    Csipstatsretryresponses interface{}

    // This object reflects the total number of PRovisional ACKnowledgement
    // retries sent by the user agent since system startup. The type is
    // interface{} with range: 0..4294967295.
    Csipstatsretrypracks interface{}

    // This object reflects the total number of COndition MET retries sent by the
    // user agent since system startup. The type is interface{} with range:
    // 0..4294967295.
    Csipstatsretrycomets interface{}

    // This object reflects the total number of Reliable 1xx Response retries sent
    // by the user agent since system startup. The type is interface{} with range:
    // 0..4294967295.
    Csipstatsretryreliable1Xxrsps interface{}

    // This object reflects the total number of Notify  retries that have been
    // sent by the user agent since system startup. The type is interface{} with
    // range: 0..4294967295.
    Csipstatsretrynotifys interface{}

    // This object reflects the total number of Refer retries that have been sent
    // by the user agent since system startup. The type is interface{} with range:
    // 0..4294967295.
    Csipstatsretryrefers interface{}

    // This object reflects the total number of Info retries that have been sent
    // by the user agent since system startup. The type is interface{} with range:
    // 0..4294967295.
    Csipstatsretryinfo interface{}

    // This object reflects the total number of Subscribe retries that have been
    // sent by the user agent since system startup. The type is interface{} with
    // range: 0..4294967295.
    Csipstatsretrysubscribe interface{}
}

func (csipstatsretry *CISCOSIPUAMIB_Csipstatsretry) GetEntityData() *types.CommonEntityData {
    csipstatsretry.EntityData.YFilter = csipstatsretry.YFilter
    csipstatsretry.EntityData.YangName = "cSipStatsRetry"
    csipstatsretry.EntityData.BundleName = "cisco_ios_xe"
    csipstatsretry.EntityData.ParentYangName = "CISCO-SIP-UA-MIB"
    csipstatsretry.EntityData.SegmentPath = "cSipStatsRetry"
    csipstatsretry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csipstatsretry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csipstatsretry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csipstatsretry.EntityData.Children = make(map[string]types.YChild)
    csipstatsretry.EntityData.Leafs = make(map[string]types.YLeaf)
    csipstatsretry.EntityData.Leafs["cSipStatsRetryInvites"] = types.YLeaf{"Csipstatsretryinvites", csipstatsretry.Csipstatsretryinvites}
    csipstatsretry.EntityData.Leafs["cSipStatsRetryByes"] = types.YLeaf{"Csipstatsretrybyes", csipstatsretry.Csipstatsretrybyes}
    csipstatsretry.EntityData.Leafs["cSipStatsRetryCancels"] = types.YLeaf{"Csipstatsretrycancels", csipstatsretry.Csipstatsretrycancels}
    csipstatsretry.EntityData.Leafs["cSipStatsRetryRegisters"] = types.YLeaf{"Csipstatsretryregisters", csipstatsretry.Csipstatsretryregisters}
    csipstatsretry.EntityData.Leafs["cSipStatsRetryResponses"] = types.YLeaf{"Csipstatsretryresponses", csipstatsretry.Csipstatsretryresponses}
    csipstatsretry.EntityData.Leafs["cSipStatsRetryPracks"] = types.YLeaf{"Csipstatsretrypracks", csipstatsretry.Csipstatsretrypracks}
    csipstatsretry.EntityData.Leafs["cSipStatsRetryComets"] = types.YLeaf{"Csipstatsretrycomets", csipstatsretry.Csipstatsretrycomets}
    csipstatsretry.EntityData.Leafs["cSipStatsRetryReliable1xxRsps"] = types.YLeaf{"Csipstatsretryreliable1Xxrsps", csipstatsretry.Csipstatsretryreliable1Xxrsps}
    csipstatsretry.EntityData.Leafs["cSipStatsRetryNotifys"] = types.YLeaf{"Csipstatsretrynotifys", csipstatsretry.Csipstatsretrynotifys}
    csipstatsretry.EntityData.Leafs["cSipStatsRetryRefers"] = types.YLeaf{"Csipstatsretryrefers", csipstatsretry.Csipstatsretryrefers}
    csipstatsretry.EntityData.Leafs["cSipStatsRetryInfo"] = types.YLeaf{"Csipstatsretryinfo", csipstatsretry.Csipstatsretryinfo}
    csipstatsretry.EntityData.Leafs["cSipStatsRetrySubscribe"] = types.YLeaf{"Csipstatsretrysubscribe", csipstatsretry.Csipstatsretrysubscribe}
    return &(csipstatsretry.EntityData)
}

// CISCOSIPUAMIB_Csipstatsmisc
type CISCOSIPUAMIB_Csipstatsmisc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object reflects the total number of incoming Redirect  (3xx) response
    // messages that have been mapped to Client  Error (4xx) response messages.
    // The type is interface{} with range: 0..4294967295.
    Csipstatsmisc3Xxmappedto4Xxrsps interface{}
}

func (csipstatsmisc *CISCOSIPUAMIB_Csipstatsmisc) GetEntityData() *types.CommonEntityData {
    csipstatsmisc.EntityData.YFilter = csipstatsmisc.YFilter
    csipstatsmisc.EntityData.YangName = "cSipStatsMisc"
    csipstatsmisc.EntityData.BundleName = "cisco_ios_xe"
    csipstatsmisc.EntityData.ParentYangName = "CISCO-SIP-UA-MIB"
    csipstatsmisc.EntityData.SegmentPath = "cSipStatsMisc"
    csipstatsmisc.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csipstatsmisc.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csipstatsmisc.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csipstatsmisc.EntityData.Children = make(map[string]types.YChild)
    csipstatsmisc.EntityData.Leafs = make(map[string]types.YLeaf)
    csipstatsmisc.EntityData.Leafs["cSipStatsMisc3xxMappedTo4xxRsps"] = types.YLeaf{"Csipstatsmisc3Xxmappedto4Xxrsps", csipstatsmisc.Csipstatsmisc3Xxmappedto4Xxrsps}
    return &(csipstatsmisc.EntityData)
}

// CISCOSIPUAMIB_Csipstatsconnection
type CISCOSIPUAMIB_Csipstatsconnection struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object reflects the total number of TCP send failures since system
    // startup. The type is interface{} with range: 0..4294967295.
    Csipstatsconntcpsendfailures interface{}

    // This object reflects the total number of UDP send failures since system
    // startup. The type is interface{} with range: 0..4294967295.
    Csipstatsconnudpsendfailures interface{}

    // This object reflects the total number of Remote Closures  for TCP since
    // system startup. The type is interface{} with range: 0..4294967295.
    Csipstatsconntcpremoteclosures interface{}

    // This object reflects the total number of connection create failures for UDP
    // since system startup. The type is interface{} with range: 0..4294967295.
    Csipstatsconnudpcreatefailures interface{}

    // This object reflects the total number of connection create failures for TCP
    // since system startup. The type is interface{} with range: 0..4294967295.
    Csipstatsconntcpcreatefailures interface{}

    // This object reflects the total number of UDP connections that  timed out
    // due to inactivity since system startup. The type is interface{} with range:
    // 0..4294967295.
    Csipstatsconnudpinactivetimeouts interface{}

    // This object reflects the total number of TCP connections that timed out due
    // to inactivity since system startup. The type is interface{} with range:
    // 0..4294967295.
    Csipstatsconntcpinactivetimeouts interface{}

    // This object reflects the total number of active UDP connections since
    // system startup. The type is interface{} with range: 0..4294967295.
    Csipstatsactiveudpconnections interface{}

    // This object reflects the total number of active TCP connections since
    // system startup. The type is interface{} with range: 0..4294967295.
    Csipstatsactivetcpconnections interface{}
}

func (csipstatsconnection *CISCOSIPUAMIB_Csipstatsconnection) GetEntityData() *types.CommonEntityData {
    csipstatsconnection.EntityData.YFilter = csipstatsconnection.YFilter
    csipstatsconnection.EntityData.YangName = "cSipStatsConnection"
    csipstatsconnection.EntityData.BundleName = "cisco_ios_xe"
    csipstatsconnection.EntityData.ParentYangName = "CISCO-SIP-UA-MIB"
    csipstatsconnection.EntityData.SegmentPath = "cSipStatsConnection"
    csipstatsconnection.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csipstatsconnection.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csipstatsconnection.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csipstatsconnection.EntityData.Children = make(map[string]types.YChild)
    csipstatsconnection.EntityData.Leafs = make(map[string]types.YLeaf)
    csipstatsconnection.EntityData.Leafs["cSipStatsConnTCPSendFailures"] = types.YLeaf{"Csipstatsconntcpsendfailures", csipstatsconnection.Csipstatsconntcpsendfailures}
    csipstatsconnection.EntityData.Leafs["cSipStatsConnUDPSendFailures"] = types.YLeaf{"Csipstatsconnudpsendfailures", csipstatsconnection.Csipstatsconnudpsendfailures}
    csipstatsconnection.EntityData.Leafs["cSipStatsConnTCPRemoteClosures"] = types.YLeaf{"Csipstatsconntcpremoteclosures", csipstatsconnection.Csipstatsconntcpremoteclosures}
    csipstatsconnection.EntityData.Leafs["cSipStatsConnUDPCreateFailures"] = types.YLeaf{"Csipstatsconnudpcreatefailures", csipstatsconnection.Csipstatsconnudpcreatefailures}
    csipstatsconnection.EntityData.Leafs["cSipStatsConnTCPCreateFailures"] = types.YLeaf{"Csipstatsconntcpcreatefailures", csipstatsconnection.Csipstatsconntcpcreatefailures}
    csipstatsconnection.EntityData.Leafs["cSipStatsConnUDPInactiveTimeouts"] = types.YLeaf{"Csipstatsconnudpinactivetimeouts", csipstatsconnection.Csipstatsconnudpinactivetimeouts}
    csipstatsconnection.EntityData.Leafs["cSipStatsConnTCPInactiveTimeouts"] = types.YLeaf{"Csipstatsconntcpinactivetimeouts", csipstatsconnection.Csipstatsconntcpinactivetimeouts}
    csipstatsconnection.EntityData.Leafs["cSipStatsActiveUDPConnections"] = types.YLeaf{"Csipstatsactiveudpconnections", csipstatsconnection.Csipstatsactiveudpconnections}
    csipstatsconnection.EntityData.Leafs["cSipStatsActiveTCPConnections"] = types.YLeaf{"Csipstatsactivetcpconnections", csipstatsconnection.Csipstatsactivetcpconnections}
    return &(csipstatsconnection.EntityData)
}

// CISCOSIPUAMIB_Csipcfgearlymediatable
// This table contains configuration for Early
// Media Cut Through.  The configuration controls
// how the SIP user agent will process 1xx
// (Provisional) SIP response messages that contain 
// Session Definition Protocol (SDP) payloads.
type CISCOSIPUAMIB_Csipcfgearlymediatable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A row in the cSipCfgEarlyMediaTable. A row is accessible with a Provisional
    // (1xx) status code value (eg, 180) and provides an object for the enabling
    // or disabling of the Early Media Cut Through functionality. The type is
    // slice of CISCOSIPUAMIB_Csipcfgearlymediatable_Csipcfgearlymediaentry.
    Csipcfgearlymediaentry []CISCOSIPUAMIB_Csipcfgearlymediatable_Csipcfgearlymediaentry
}

func (csipcfgearlymediatable *CISCOSIPUAMIB_Csipcfgearlymediatable) GetEntityData() *types.CommonEntityData {
    csipcfgearlymediatable.EntityData.YFilter = csipcfgearlymediatable.YFilter
    csipcfgearlymediatable.EntityData.YangName = "cSipCfgEarlyMediaTable"
    csipcfgearlymediatable.EntityData.BundleName = "cisco_ios_xe"
    csipcfgearlymediatable.EntityData.ParentYangName = "CISCO-SIP-UA-MIB"
    csipcfgearlymediatable.EntityData.SegmentPath = "cSipCfgEarlyMediaTable"
    csipcfgearlymediatable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csipcfgearlymediatable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csipcfgearlymediatable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csipcfgearlymediatable.EntityData.Children = make(map[string]types.YChild)
    csipcfgearlymediatable.EntityData.Children["cSipCfgEarlyMediaEntry"] = types.YChild{"Csipcfgearlymediaentry", nil}
    for i := range csipcfgearlymediatable.Csipcfgearlymediaentry {
        csipcfgearlymediatable.EntityData.Children[types.GetSegmentPath(&csipcfgearlymediatable.Csipcfgearlymediaentry[i])] = types.YChild{"Csipcfgearlymediaentry", &csipcfgearlymediatable.Csipcfgearlymediaentry[i]}
    }
    csipcfgearlymediatable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(csipcfgearlymediatable.EntityData)
}

// CISCOSIPUAMIB_Csipcfgearlymediatable_Csipcfgearlymediaentry
// A row in the cSipCfgEarlyMediaTable.
// A row is accessible with a Provisional (1xx)
// status code value (eg, 180) and provides
// an object for the enabling or disabling of
// the Early Media Cut Through functionality.
type CISCOSIPUAMIB_Csipcfgearlymediatable_Csipcfgearlymediaentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. A unique identifier of a row in this table and a
    // valid SIP status code. The type is interface{} with range: 1..2147483647.
    Csipcfgearlymediastatuscodeindex interface{}

    // This object specifies whether Early Media  Cut Through is enabled or
    // disabled for the  SIP response messages with status codes that  match
    // cSipCfgEarlyMediaStatusCodeIndex.  If 'true', early media cut through is
    // disabled, and the user agent will process the message as though it did not
    // contain any SDP payload.  If 'false', early media cut through is enabled,
    // and the user agent will process the message similar to a 183 (Session
    // Progress) and cut through for early media.  The assumption being that the
    // SDP is an indication that the far end is going to send early media. The
    // type is bool.
    Csipcfgearlymediacutthrudisabled interface{}
}

func (csipcfgearlymediaentry *CISCOSIPUAMIB_Csipcfgearlymediatable_Csipcfgearlymediaentry) GetEntityData() *types.CommonEntityData {
    csipcfgearlymediaentry.EntityData.YFilter = csipcfgearlymediaentry.YFilter
    csipcfgearlymediaentry.EntityData.YangName = "cSipCfgEarlyMediaEntry"
    csipcfgearlymediaentry.EntityData.BundleName = "cisco_ios_xe"
    csipcfgearlymediaentry.EntityData.ParentYangName = "cSipCfgEarlyMediaTable"
    csipcfgearlymediaentry.EntityData.SegmentPath = "cSipCfgEarlyMediaEntry" + "[cSipCfgEarlyMediaStatusCodeIndex='" + fmt.Sprintf("%v", csipcfgearlymediaentry.Csipcfgearlymediastatuscodeindex) + "']"
    csipcfgearlymediaentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csipcfgearlymediaentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csipcfgearlymediaentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csipcfgearlymediaentry.EntityData.Children = make(map[string]types.YChild)
    csipcfgearlymediaentry.EntityData.Leafs = make(map[string]types.YLeaf)
    csipcfgearlymediaentry.EntityData.Leafs["cSipCfgEarlyMediaStatusCodeIndex"] = types.YLeaf{"Csipcfgearlymediastatuscodeindex", csipcfgearlymediaentry.Csipcfgearlymediastatuscodeindex}
    csipcfgearlymediaentry.EntityData.Leafs["cSipCfgEarlyMediaCutThruDisabled"] = types.YLeaf{"Csipcfgearlymediacutthrudisabled", csipcfgearlymediaentry.Csipcfgearlymediacutthrudisabled}
    return &(csipcfgearlymediaentry.EntityData)
}

// CISCOSIPUAMIB_Csipcfgbindsourceaddrtable
// This table contains configuration for binding
// the scope of packets to the particular ethernet
// interface. The scope for the packets can be
// specified as either 'signalling' or 'media'
// packets. The ethernet interface shall be
// specified by the interface index. The table
// shall be indexed based on the scope.
type CISCOSIPUAMIB_Csipcfgbindsourceaddrtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A row in the cSipCfgBindSourceAddrTable. A row is accessible with the scope
    // of packets to which the source IP address of the interface designated by
    // cSipCfgBindSourceAddrInterface will be bound. The type is slice of
    // CISCOSIPUAMIB_Csipcfgbindsourceaddrtable_Csipcfgbindsourceaddrentry.
    Csipcfgbindsourceaddrentry []CISCOSIPUAMIB_Csipcfgbindsourceaddrtable_Csipcfgbindsourceaddrentry
}

func (csipcfgbindsourceaddrtable *CISCOSIPUAMIB_Csipcfgbindsourceaddrtable) GetEntityData() *types.CommonEntityData {
    csipcfgbindsourceaddrtable.EntityData.YFilter = csipcfgbindsourceaddrtable.YFilter
    csipcfgbindsourceaddrtable.EntityData.YangName = "cSipCfgBindSourceAddrTable"
    csipcfgbindsourceaddrtable.EntityData.BundleName = "cisco_ios_xe"
    csipcfgbindsourceaddrtable.EntityData.ParentYangName = "CISCO-SIP-UA-MIB"
    csipcfgbindsourceaddrtable.EntityData.SegmentPath = "cSipCfgBindSourceAddrTable"
    csipcfgbindsourceaddrtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csipcfgbindsourceaddrtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csipcfgbindsourceaddrtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csipcfgbindsourceaddrtable.EntityData.Children = make(map[string]types.YChild)
    csipcfgbindsourceaddrtable.EntityData.Children["cSipCfgBindSourceAddrEntry"] = types.YChild{"Csipcfgbindsourceaddrentry", nil}
    for i := range csipcfgbindsourceaddrtable.Csipcfgbindsourceaddrentry {
        csipcfgbindsourceaddrtable.EntityData.Children[types.GetSegmentPath(&csipcfgbindsourceaddrtable.Csipcfgbindsourceaddrentry[i])] = types.YChild{"Csipcfgbindsourceaddrentry", &csipcfgbindsourceaddrtable.Csipcfgbindsourceaddrentry[i]}
    }
    csipcfgbindsourceaddrtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(csipcfgbindsourceaddrtable.EntityData)
}

// CISCOSIPUAMIB_Csipcfgbindsourceaddrtable_Csipcfgbindsourceaddrentry
// A row in the cSipCfgBindSourceAddrTable.
// A row is accessible with the scope of packets
// to which the source IP address of the interface
// designated by cSipCfgBindSourceAddrInterface
// will be bound.
type CISCOSIPUAMIB_Csipcfgbindsourceaddrtable_Csipcfgbindsourceaddrentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. A unique identifier of a row in this table and
    // specifies the scope of packets to which the source IP address of the
    // interface designated by cSipCfgBindSourceAddrInterface will be bound. The
    // type is Csipcfgbindsourceaddrscope.
    Csipcfgbindsourceaddrscope interface{}

    // This object may specify the interface where the source IP address used in
    // SIP signalling or media packets is configured.  A value of 0 means that
    // there is no specific source address configured and in this case the best
    // local IP address will be chosen by the system. The type is interface{} with
    // range: 0..2147483647.
    Csipcfgbindsourceaddrinterface interface{}
}

func (csipcfgbindsourceaddrentry *CISCOSIPUAMIB_Csipcfgbindsourceaddrtable_Csipcfgbindsourceaddrentry) GetEntityData() *types.CommonEntityData {
    csipcfgbindsourceaddrentry.EntityData.YFilter = csipcfgbindsourceaddrentry.YFilter
    csipcfgbindsourceaddrentry.EntityData.YangName = "cSipCfgBindSourceAddrEntry"
    csipcfgbindsourceaddrentry.EntityData.BundleName = "cisco_ios_xe"
    csipcfgbindsourceaddrentry.EntityData.ParentYangName = "cSipCfgBindSourceAddrTable"
    csipcfgbindsourceaddrentry.EntityData.SegmentPath = "cSipCfgBindSourceAddrEntry" + "[cSipCfgBindSourceAddrScope='" + fmt.Sprintf("%v", csipcfgbindsourceaddrentry.Csipcfgbindsourceaddrscope) + "']"
    csipcfgbindsourceaddrentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csipcfgbindsourceaddrentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csipcfgbindsourceaddrentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csipcfgbindsourceaddrentry.EntityData.Children = make(map[string]types.YChild)
    csipcfgbindsourceaddrentry.EntityData.Leafs = make(map[string]types.YLeaf)
    csipcfgbindsourceaddrentry.EntityData.Leafs["cSipCfgBindSourceAddrScope"] = types.YLeaf{"Csipcfgbindsourceaddrscope", csipcfgbindsourceaddrentry.Csipcfgbindsourceaddrscope}
    csipcfgbindsourceaddrentry.EntityData.Leafs["cSipCfgBindSourceAddrInterface"] = types.YLeaf{"Csipcfgbindsourceaddrinterface", csipcfgbindsourceaddrentry.Csipcfgbindsourceaddrinterface}
    return &(csipcfgbindsourceaddrentry.EntityData)
}

// CISCOSIPUAMIB_Csipcfgbindsourceaddrtable_Csipcfgbindsourceaddrentry_Csipcfgbindsourceaddrscope represents will be bound.
type CISCOSIPUAMIB_Csipcfgbindsourceaddrtable_Csipcfgbindsourceaddrentry_Csipcfgbindsourceaddrscope string

const (
    CISCOSIPUAMIB_Csipcfgbindsourceaddrtable_Csipcfgbindsourceaddrentry_Csipcfgbindsourceaddrscope_media CISCOSIPUAMIB_Csipcfgbindsourceaddrtable_Csipcfgbindsourceaddrentry_Csipcfgbindsourceaddrscope = "media"

    CISCOSIPUAMIB_Csipcfgbindsourceaddrtable_Csipcfgbindsourceaddrentry_Csipcfgbindsourceaddrscope_control CISCOSIPUAMIB_Csipcfgbindsourceaddrtable_Csipcfgbindsourceaddrentry_Csipcfgbindsourceaddrscope = "control"
)

// CISCOSIPUAMIB_Csipcfgpeertable
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
type CISCOSIPUAMIB_Csipcfgpeertable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A row in the cSipCfgPeerTable. The type is slice of
    // CISCOSIPUAMIB_Csipcfgpeertable_Csipcfgpeerentry.
    Csipcfgpeerentry []CISCOSIPUAMIB_Csipcfgpeertable_Csipcfgpeerentry
}

func (csipcfgpeertable *CISCOSIPUAMIB_Csipcfgpeertable) GetEntityData() *types.CommonEntityData {
    csipcfgpeertable.EntityData.YFilter = csipcfgpeertable.YFilter
    csipcfgpeertable.EntityData.YangName = "cSipCfgPeerTable"
    csipcfgpeertable.EntityData.BundleName = "cisco_ios_xe"
    csipcfgpeertable.EntityData.ParentYangName = "CISCO-SIP-UA-MIB"
    csipcfgpeertable.EntityData.SegmentPath = "cSipCfgPeerTable"
    csipcfgpeertable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csipcfgpeertable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csipcfgpeertable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csipcfgpeertable.EntityData.Children = make(map[string]types.YChild)
    csipcfgpeertable.EntityData.Children["cSipCfgPeerEntry"] = types.YChild{"Csipcfgpeerentry", nil}
    for i := range csipcfgpeertable.Csipcfgpeerentry {
        csipcfgpeertable.EntityData.Children[types.GetSegmentPath(&csipcfgpeertable.Csipcfgpeerentry[i])] = types.YChild{"Csipcfgpeerentry", &csipcfgpeertable.Csipcfgpeerentry[i]}
    }
    csipcfgpeertable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(csipcfgpeertable.EntityData)
}

// CISCOSIPUAMIB_Csipcfgpeertable_Csipcfgpeerentry
// A row in the cSipCfgPeerTable.
type CISCOSIPUAMIB_Csipcfgpeertable_Csipcfgpeerentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An arbitrary index that uniquely identifies a 
    // dial-peer configured for SIP. The type is interface{} with range:
    // 1..2147483647.
    Csipcfgpeerindex interface{}

    // This object specifies the session transport  protocol that will be used by
    // this dial-peer for outbound SIP messages.    The value 'system' is the
    // default and indicates  that this dial-peer should use the value set by 
    // cSipCfgOutSessionTransport instead. The type is
    // Csipcfgpeeroutsessiontransport.
    Csipcfgpeeroutsessiontransport interface{}

    // This object specifies the string that will be  placed in either the
    // Supported or Require SIP  header, as specified by
    // cSipCfgPeerReliable1xxRspHdr. The type is string.
    Csipcfgpeerreliable1Xxrspstr interface{}

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
    // cSipCfgReliable1xxRspHdr instead. The type is Csipcfgpeerreliable1Xxrsphdr.
    Csipcfgpeerreliable1Xxrsphdr interface{}

    // This object specifies the URL type sent in outbound INVITES generated by
    // this device.  The value 'system' is the default and indicates that this 
    // dial-peer should use the value of cSipCfgUrlType instead. The type is
    // Csipcfgpeerurltype.
    Csipcfgpeerurltype interface{}

    // This object specifies if the functionality of switching between transports
    // from UDP to TCP if the message size of a Request is greater than 1300 bytes
    // is enabled or not. The type is bool.
    Csipcfgpeerswitchtransenabled interface{}
}

func (csipcfgpeerentry *CISCOSIPUAMIB_Csipcfgpeertable_Csipcfgpeerentry) GetEntityData() *types.CommonEntityData {
    csipcfgpeerentry.EntityData.YFilter = csipcfgpeerentry.YFilter
    csipcfgpeerentry.EntityData.YangName = "cSipCfgPeerEntry"
    csipcfgpeerentry.EntityData.BundleName = "cisco_ios_xe"
    csipcfgpeerentry.EntityData.ParentYangName = "cSipCfgPeerTable"
    csipcfgpeerentry.EntityData.SegmentPath = "cSipCfgPeerEntry" + "[cSipCfgPeerIndex='" + fmt.Sprintf("%v", csipcfgpeerentry.Csipcfgpeerindex) + "']"
    csipcfgpeerentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csipcfgpeerentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csipcfgpeerentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csipcfgpeerentry.EntityData.Children = make(map[string]types.YChild)
    csipcfgpeerentry.EntityData.Leafs = make(map[string]types.YLeaf)
    csipcfgpeerentry.EntityData.Leafs["cSipCfgPeerIndex"] = types.YLeaf{"Csipcfgpeerindex", csipcfgpeerentry.Csipcfgpeerindex}
    csipcfgpeerentry.EntityData.Leafs["cSipCfgPeerOutSessionTransport"] = types.YLeaf{"Csipcfgpeeroutsessiontransport", csipcfgpeerentry.Csipcfgpeeroutsessiontransport}
    csipcfgpeerentry.EntityData.Leafs["cSipCfgPeerReliable1xxRspStr"] = types.YLeaf{"Csipcfgpeerreliable1Xxrspstr", csipcfgpeerentry.Csipcfgpeerreliable1Xxrspstr}
    csipcfgpeerentry.EntityData.Leafs["cSipCfgPeerReliable1xxRspHdr"] = types.YLeaf{"Csipcfgpeerreliable1Xxrsphdr", csipcfgpeerentry.Csipcfgpeerreliable1Xxrsphdr}
    csipcfgpeerentry.EntityData.Leafs["cSipCfgPeerUrlType"] = types.YLeaf{"Csipcfgpeerurltype", csipcfgpeerentry.Csipcfgpeerurltype}
    csipcfgpeerentry.EntityData.Leafs["cSipCfgPeerSwitchTransEnabled"] = types.YLeaf{"Csipcfgpeerswitchtransenabled", csipcfgpeerentry.Csipcfgpeerswitchtransenabled}
    return &(csipcfgpeerentry.EntityData)
}

// CISCOSIPUAMIB_Csipcfgpeertable_Csipcfgpeerentry_Csipcfgpeeroutsessiontransport represents cSipCfgOutSessionTransport instead.
type CISCOSIPUAMIB_Csipcfgpeertable_Csipcfgpeerentry_Csipcfgpeeroutsessiontransport string

const (
    CISCOSIPUAMIB_Csipcfgpeertable_Csipcfgpeerentry_Csipcfgpeeroutsessiontransport_system CISCOSIPUAMIB_Csipcfgpeertable_Csipcfgpeerentry_Csipcfgpeeroutsessiontransport = "system"

    CISCOSIPUAMIB_Csipcfgpeertable_Csipcfgpeerentry_Csipcfgpeeroutsessiontransport_udp CISCOSIPUAMIB_Csipcfgpeertable_Csipcfgpeerentry_Csipcfgpeeroutsessiontransport = "udp"

    CISCOSIPUAMIB_Csipcfgpeertable_Csipcfgpeerentry_Csipcfgpeeroutsessiontransport_tcp CISCOSIPUAMIB_Csipcfgpeertable_Csipcfgpeerentry_Csipcfgpeeroutsessiontransport = "tcp"
)

// CISCOSIPUAMIB_Csipcfgpeertable_Csipcfgpeerentry_Csipcfgpeerreliable1Xxrsphdr represents instead.
type CISCOSIPUAMIB_Csipcfgpeertable_Csipcfgpeerentry_Csipcfgpeerreliable1Xxrsphdr string

const (
    CISCOSIPUAMIB_Csipcfgpeertable_Csipcfgpeerentry_Csipcfgpeerreliable1Xxrsphdr_system CISCOSIPUAMIB_Csipcfgpeertable_Csipcfgpeerentry_Csipcfgpeerreliable1Xxrsphdr = "system"

    CISCOSIPUAMIB_Csipcfgpeertable_Csipcfgpeerentry_Csipcfgpeerreliable1Xxrsphdr_supported CISCOSIPUAMIB_Csipcfgpeertable_Csipcfgpeerentry_Csipcfgpeerreliable1Xxrsphdr = "supported"

    CISCOSIPUAMIB_Csipcfgpeertable_Csipcfgpeerentry_Csipcfgpeerreliable1Xxrsphdr_require CISCOSIPUAMIB_Csipcfgpeertable_Csipcfgpeerentry_Csipcfgpeerreliable1Xxrsphdr = "require"

    CISCOSIPUAMIB_Csipcfgpeertable_Csipcfgpeerentry_Csipcfgpeerreliable1Xxrsphdr_disabled CISCOSIPUAMIB_Csipcfgpeertable_Csipcfgpeerentry_Csipcfgpeerreliable1Xxrsphdr = "disabled"
)

// CISCOSIPUAMIB_Csipcfgpeertable_Csipcfgpeerentry_Csipcfgpeerurltype represents dial-peer should use the value of cSipCfgUrlType instead.
type CISCOSIPUAMIB_Csipcfgpeertable_Csipcfgpeerentry_Csipcfgpeerurltype string

const (
    CISCOSIPUAMIB_Csipcfgpeertable_Csipcfgpeerentry_Csipcfgpeerurltype_system CISCOSIPUAMIB_Csipcfgpeertable_Csipcfgpeerentry_Csipcfgpeerurltype = "system"

    CISCOSIPUAMIB_Csipcfgpeertable_Csipcfgpeerentry_Csipcfgpeerurltype_sip CISCOSIPUAMIB_Csipcfgpeertable_Csipcfgpeerentry_Csipcfgpeerurltype = "sip"

    CISCOSIPUAMIB_Csipcfgpeertable_Csipcfgpeerentry_Csipcfgpeerurltype_tel CISCOSIPUAMIB_Csipcfgpeertable_Csipcfgpeerentry_Csipcfgpeerurltype = "tel"
)

// CISCOSIPUAMIB_Csipcfgstatuscausetable
// This table contains SIP status code to PSTN cause code
// mapping configuration.  Inbound SIP response messages 
// that will result in outbound PSTN signalling messages
// will have the SIP status codes transposed into the
// PSTN cause codes as prescribed by the contents of this 
// table.
type CISCOSIPUAMIB_Csipcfgstatuscausetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A row in the cSipCfgStatusCauseTable.  Entries cannot be created or
    // destroyed by the actions of any NMS. The type is slice of
    // CISCOSIPUAMIB_Csipcfgstatuscausetable_Csipcfgstatuscauseentry.
    Csipcfgstatuscauseentry []CISCOSIPUAMIB_Csipcfgstatuscausetable_Csipcfgstatuscauseentry
}

func (csipcfgstatuscausetable *CISCOSIPUAMIB_Csipcfgstatuscausetable) GetEntityData() *types.CommonEntityData {
    csipcfgstatuscausetable.EntityData.YFilter = csipcfgstatuscausetable.YFilter
    csipcfgstatuscausetable.EntityData.YangName = "cSipCfgStatusCauseTable"
    csipcfgstatuscausetable.EntityData.BundleName = "cisco_ios_xe"
    csipcfgstatuscausetable.EntityData.ParentYangName = "CISCO-SIP-UA-MIB"
    csipcfgstatuscausetable.EntityData.SegmentPath = "cSipCfgStatusCauseTable"
    csipcfgstatuscausetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csipcfgstatuscausetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csipcfgstatuscausetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csipcfgstatuscausetable.EntityData.Children = make(map[string]types.YChild)
    csipcfgstatuscausetable.EntityData.Children["cSipCfgStatusCauseEntry"] = types.YChild{"Csipcfgstatuscauseentry", nil}
    for i := range csipcfgstatuscausetable.Csipcfgstatuscauseentry {
        csipcfgstatuscausetable.EntityData.Children[types.GetSegmentPath(&csipcfgstatuscausetable.Csipcfgstatuscauseentry[i])] = types.YChild{"Csipcfgstatuscauseentry", &csipcfgstatuscausetable.Csipcfgstatuscauseentry[i]}
    }
    csipcfgstatuscausetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(csipcfgstatuscausetable.EntityData)
}

// CISCOSIPUAMIB_Csipcfgstatuscausetable_Csipcfgstatuscauseentry
// A row in the cSipCfgStatusCauseTable.  Entries cannot be
// created or destroyed by the actions of any NMS.
type CISCOSIPUAMIB_Csipcfgstatuscausetable_Csipcfgstatuscauseentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. A unique identifier of a row in this table and a
    // valid SIP status code. The type is interface{} with range: 1..2147483647.
    Csipcfgstatuscodeindex interface{}

    // The PSTN cause code to which the SIP status code given by
    // cSipCfgStatusCodeIndex is to be mapped for outbound PSTN signalling
    // messages. The type is interface{} with range: 1..2147483647.
    Csipcfgpstncause interface{}

    // This object reflects the storage status of this table entry.  If the value
    // is 'volatile', then this entry only exists in RAM and the information would
    // be lost (reverting to defaults) on system reload.   If the value is
    // 'nonVolatile' then this entry has been  written to NVRAM and will persist
    // across system reloads. The type is StorageType.
    Csipcfgstatuscausestorestatus interface{}
}

func (csipcfgstatuscauseentry *CISCOSIPUAMIB_Csipcfgstatuscausetable_Csipcfgstatuscauseentry) GetEntityData() *types.CommonEntityData {
    csipcfgstatuscauseentry.EntityData.YFilter = csipcfgstatuscauseentry.YFilter
    csipcfgstatuscauseentry.EntityData.YangName = "cSipCfgStatusCauseEntry"
    csipcfgstatuscauseentry.EntityData.BundleName = "cisco_ios_xe"
    csipcfgstatuscauseentry.EntityData.ParentYangName = "cSipCfgStatusCauseTable"
    csipcfgstatuscauseentry.EntityData.SegmentPath = "cSipCfgStatusCauseEntry" + "[cSipCfgStatusCodeIndex='" + fmt.Sprintf("%v", csipcfgstatuscauseentry.Csipcfgstatuscodeindex) + "']"
    csipcfgstatuscauseentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csipcfgstatuscauseentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csipcfgstatuscauseentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csipcfgstatuscauseentry.EntityData.Children = make(map[string]types.YChild)
    csipcfgstatuscauseentry.EntityData.Leafs = make(map[string]types.YLeaf)
    csipcfgstatuscauseentry.EntityData.Leafs["cSipCfgStatusCodeIndex"] = types.YLeaf{"Csipcfgstatuscodeindex", csipcfgstatuscauseentry.Csipcfgstatuscodeindex}
    csipcfgstatuscauseentry.EntityData.Leafs["cSipCfgPstnCause"] = types.YLeaf{"Csipcfgpstncause", csipcfgstatuscauseentry.Csipcfgpstncause}
    csipcfgstatuscauseentry.EntityData.Leafs["cSipCfgStatusCauseStoreStatus"] = types.YLeaf{"Csipcfgstatuscausestorestatus", csipcfgstatuscauseentry.Csipcfgstatuscausestorestatus}
    return &(csipcfgstatuscauseentry.EntityData)
}

// CISCOSIPUAMIB_Csipcfgcausestatustable
// This table contains PSTN cause code to SIP status code
// mapping configuration.   Inbound PSTN signalling messages
// that will result in outbound SIP response messages 
// will have the PSTN cause codes transposed into the
// SIP status codes as prescribed by the contents of this 
// table.
type CISCOSIPUAMIB_Csipcfgcausestatustable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A row in the cSipCfgCauseStatusTable. Entries cannot be created or
    // destroyed by the actions of any NMS. The type is slice of
    // CISCOSIPUAMIB_Csipcfgcausestatustable_Csipcfgcausestatusentry.
    Csipcfgcausestatusentry []CISCOSIPUAMIB_Csipcfgcausestatustable_Csipcfgcausestatusentry
}

func (csipcfgcausestatustable *CISCOSIPUAMIB_Csipcfgcausestatustable) GetEntityData() *types.CommonEntityData {
    csipcfgcausestatustable.EntityData.YFilter = csipcfgcausestatustable.YFilter
    csipcfgcausestatustable.EntityData.YangName = "cSipCfgCauseStatusTable"
    csipcfgcausestatustable.EntityData.BundleName = "cisco_ios_xe"
    csipcfgcausestatustable.EntityData.ParentYangName = "CISCO-SIP-UA-MIB"
    csipcfgcausestatustable.EntityData.SegmentPath = "cSipCfgCauseStatusTable"
    csipcfgcausestatustable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csipcfgcausestatustable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csipcfgcausestatustable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csipcfgcausestatustable.EntityData.Children = make(map[string]types.YChild)
    csipcfgcausestatustable.EntityData.Children["cSipCfgCauseStatusEntry"] = types.YChild{"Csipcfgcausestatusentry", nil}
    for i := range csipcfgcausestatustable.Csipcfgcausestatusentry {
        csipcfgcausestatustable.EntityData.Children[types.GetSegmentPath(&csipcfgcausestatustable.Csipcfgcausestatusentry[i])] = types.YChild{"Csipcfgcausestatusentry", &csipcfgcausestatustable.Csipcfgcausestatusentry[i]}
    }
    csipcfgcausestatustable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(csipcfgcausestatustable.EntityData)
}

// CISCOSIPUAMIB_Csipcfgcausestatustable_Csipcfgcausestatusentry
// A row in the cSipCfgCauseStatusTable. Entries cannot be
// created or destroyed by the actions of any NMS.
type CISCOSIPUAMIB_Csipcfgcausestatustable_Csipcfgcausestatusentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. A unique identifier of a row in this table and a
    // valid PSTN cause code. The type is interface{} with range: 1..2147483647.
    Csipcfgpstncauseindex interface{}

    // The SIP status code to which the PSTN cause code given by
    // cSipCfgPstnCauseIndex is to be mapped for outbound SIP response messages.
    // The type is interface{} with range: 1..2147483647.
    Csipcfgstatuscode interface{}

    // This object reflects the storage status of this table entry.  If the value
    // is 'volatile', then this entry only exists in RAM and the information would
    // be lost (reverting to defaults) on system reload.   If the value is
    // 'nonVolatile' then this entry has been  written to NVRAM and will persist
    // across system reloads. The type is StorageType.
    Csipcfgcausestatusstorestatus interface{}
}

func (csipcfgcausestatusentry *CISCOSIPUAMIB_Csipcfgcausestatustable_Csipcfgcausestatusentry) GetEntityData() *types.CommonEntityData {
    csipcfgcausestatusentry.EntityData.YFilter = csipcfgcausestatusentry.YFilter
    csipcfgcausestatusentry.EntityData.YangName = "cSipCfgCauseStatusEntry"
    csipcfgcausestatusentry.EntityData.BundleName = "cisco_ios_xe"
    csipcfgcausestatusentry.EntityData.ParentYangName = "cSipCfgCauseStatusTable"
    csipcfgcausestatusentry.EntityData.SegmentPath = "cSipCfgCauseStatusEntry" + "[cSipCfgPstnCauseIndex='" + fmt.Sprintf("%v", csipcfgcausestatusentry.Csipcfgpstncauseindex) + "']"
    csipcfgcausestatusentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csipcfgcausestatusentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csipcfgcausestatusentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csipcfgcausestatusentry.EntityData.Children = make(map[string]types.YChild)
    csipcfgcausestatusentry.EntityData.Leafs = make(map[string]types.YLeaf)
    csipcfgcausestatusentry.EntityData.Leafs["cSipCfgPstnCauseIndex"] = types.YLeaf{"Csipcfgpstncauseindex", csipcfgcausestatusentry.Csipcfgpstncauseindex}
    csipcfgcausestatusentry.EntityData.Leafs["cSipCfgStatusCode"] = types.YLeaf{"Csipcfgstatuscode", csipcfgcausestatusentry.Csipcfgstatuscode}
    csipcfgcausestatusentry.EntityData.Leafs["cSipCfgCauseStatusStoreStatus"] = types.YLeaf{"Csipcfgcausestatusstorestatus", csipcfgcausestatusentry.Csipcfgcausestatusstorestatus}
    return &(csipcfgcausestatusentry.EntityData)
}

// CISCOSIPUAMIB_Csipstatssuccessoktable
// This table contains statistics for sent and
// received 200 Ok response messages.  The 
// statistics are kept on per SIP method basis.
type CISCOSIPUAMIB_Csipstatssuccessoktable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A row in the cSipStatsSuccessOkTable.  There is  an entry for each SIP
    // method for which 200 Ok  inbound and/or outbound statistics are kept. The
    // type is slice of
    // CISCOSIPUAMIB_Csipstatssuccessoktable_Csipstatssuccessokentry.
    Csipstatssuccessokentry []CISCOSIPUAMIB_Csipstatssuccessoktable_Csipstatssuccessokentry
}

func (csipstatssuccessoktable *CISCOSIPUAMIB_Csipstatssuccessoktable) GetEntityData() *types.CommonEntityData {
    csipstatssuccessoktable.EntityData.YFilter = csipstatssuccessoktable.YFilter
    csipstatssuccessoktable.EntityData.YangName = "cSipStatsSuccessOkTable"
    csipstatssuccessoktable.EntityData.BundleName = "cisco_ios_xe"
    csipstatssuccessoktable.EntityData.ParentYangName = "CISCO-SIP-UA-MIB"
    csipstatssuccessoktable.EntityData.SegmentPath = "cSipStatsSuccessOkTable"
    csipstatssuccessoktable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csipstatssuccessoktable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csipstatssuccessoktable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csipstatssuccessoktable.EntityData.Children = make(map[string]types.YChild)
    csipstatssuccessoktable.EntityData.Children["cSipStatsSuccessOkEntry"] = types.YChild{"Csipstatssuccessokentry", nil}
    for i := range csipstatssuccessoktable.Csipstatssuccessokentry {
        csipstatssuccessoktable.EntityData.Children[types.GetSegmentPath(&csipstatssuccessoktable.Csipstatssuccessokentry[i])] = types.YChild{"Csipstatssuccessokentry", &csipstatssuccessoktable.Csipstatssuccessokentry[i]}
    }
    csipstatssuccessoktable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(csipstatssuccessoktable.EntityData)
}

// CISCOSIPUAMIB_Csipstatssuccessoktable_Csipstatssuccessokentry
// A row in the cSipStatsSuccessOkTable.  There is 
// an entry for each SIP method for which 200 Ok 
// inbound and/or outbound statistics are kept.
type CISCOSIPUAMIB_Csipstatssuccessoktable_Csipstatssuccessokentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This object is used for instance identification of
    // objects in this table.  The value reflects a SIP method. The type is string
    // with length: 1..32.
    Csipstatssuccessokmethod interface{}

    // This object reflects the total number of Ok (200) responses sent by the
    // user agent, since system startup, that were associated with the SIP method
    // as specified by cSipStatsSuccessOkMethod. The type is interface{} with
    // range: 0..4294967295.
    Csipstatssuccessokinbounds interface{}

    // This object reflects the total number of Ok (200) responses received by the
    // user agent, since system startup, that were associated with the SIP method
    // as specified by cSipStatsSuccessOkMethod. The type is interface{} with
    // range: 0..4294967295.
    Csipstatssuccessokoutbounds interface{}
}

func (csipstatssuccessokentry *CISCOSIPUAMIB_Csipstatssuccessoktable_Csipstatssuccessokentry) GetEntityData() *types.CommonEntityData {
    csipstatssuccessokentry.EntityData.YFilter = csipstatssuccessokentry.YFilter
    csipstatssuccessokentry.EntityData.YangName = "cSipStatsSuccessOkEntry"
    csipstatssuccessokentry.EntityData.BundleName = "cisco_ios_xe"
    csipstatssuccessokentry.EntityData.ParentYangName = "cSipStatsSuccessOkTable"
    csipstatssuccessokentry.EntityData.SegmentPath = "cSipStatsSuccessOkEntry" + "[cSipStatsSuccessOkMethod='" + fmt.Sprintf("%v", csipstatssuccessokentry.Csipstatssuccessokmethod) + "']"
    csipstatssuccessokentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csipstatssuccessokentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csipstatssuccessokentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csipstatssuccessokentry.EntityData.Children = make(map[string]types.YChild)
    csipstatssuccessokentry.EntityData.Leafs = make(map[string]types.YLeaf)
    csipstatssuccessokentry.EntityData.Leafs["cSipStatsSuccessOkMethod"] = types.YLeaf{"Csipstatssuccessokmethod", csipstatssuccessokentry.Csipstatssuccessokmethod}
    csipstatssuccessokentry.EntityData.Leafs["cSipStatsSuccessOkInbounds"] = types.YLeaf{"Csipstatssuccessokinbounds", csipstatssuccessokentry.Csipstatssuccessokinbounds}
    csipstatssuccessokentry.EntityData.Leafs["cSipStatsSuccessOkOutbounds"] = types.YLeaf{"Csipstatssuccessokoutbounds", csipstatssuccessokentry.Csipstatssuccessokoutbounds}
    return &(csipstatssuccessokentry.EntityData)
}

