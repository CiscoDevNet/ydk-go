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
    parent types.Entity
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

func (cISCOSIPUAMIB *CISCOSIPUAMIB) GetFilter() yfilter.YFilter { return cISCOSIPUAMIB.YFilter }

func (cISCOSIPUAMIB *CISCOSIPUAMIB) SetFilter(yf yfilter.YFilter) { cISCOSIPUAMIB.YFilter = yf }

func (cISCOSIPUAMIB *CISCOSIPUAMIB) GetGoName(yname string) string {
    if yname == "cSipCfgBase" { return "Csipcfgbase" }
    if yname == "cSipCfgTimer" { return "Csipcfgtimer" }
    if yname == "cSipCfgRetry" { return "Csipcfgretry" }
    if yname == "cSipCfgPeer" { return "Csipcfgpeer" }
    if yname == "cSipCfgAaa" { return "Csipcfgaaa" }
    if yname == "cSipCfgVoiceServiceVoip" { return "Csipcfgvoiceservicevoip" }
    if yname == "cSipStatsInfo" { return "Csipstatsinfo" }
    if yname == "cSipStatsSuccess" { return "Csipstatssuccess" }
    if yname == "cSipStatsRedirect" { return "Csipstatsredirect" }
    if yname == "cSipStatsErrClient" { return "Csipstatserrclient" }
    if yname == "cSipStatsErrServer" { return "Csipstatserrserver" }
    if yname == "cSipStatsGlobalFail" { return "Csipstatsglobalfail" }
    if yname == "cSipStatsTraffic" { return "Csipstatstraffic" }
    if yname == "cSipStatsRetry" { return "Csipstatsretry" }
    if yname == "cSipStatsMisc" { return "Csipstatsmisc" }
    if yname == "cSipStatsConnection" { return "Csipstatsconnection" }
    if yname == "cSipCfgEarlyMediaTable" { return "Csipcfgearlymediatable" }
    if yname == "cSipCfgBindSourceAddrTable" { return "Csipcfgbindsourceaddrtable" }
    if yname == "cSipCfgPeerTable" { return "Csipcfgpeertable" }
    if yname == "cSipCfgStatusCauseTable" { return "Csipcfgstatuscausetable" }
    if yname == "cSipCfgCauseStatusTable" { return "Csipcfgcausestatustable" }
    if yname == "cSipStatsSuccessOkTable" { return "Csipstatssuccessoktable" }
    return ""
}

func (cISCOSIPUAMIB *CISCOSIPUAMIB) GetSegmentPath() string {
    return "CISCO-SIP-UA-MIB:CISCO-SIP-UA-MIB"
}

func (cISCOSIPUAMIB *CISCOSIPUAMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cSipCfgBase" {
        return &cISCOSIPUAMIB.Csipcfgbase
    }
    if childYangName == "cSipCfgTimer" {
        return &cISCOSIPUAMIB.Csipcfgtimer
    }
    if childYangName == "cSipCfgRetry" {
        return &cISCOSIPUAMIB.Csipcfgretry
    }
    if childYangName == "cSipCfgPeer" {
        return &cISCOSIPUAMIB.Csipcfgpeer
    }
    if childYangName == "cSipCfgAaa" {
        return &cISCOSIPUAMIB.Csipcfgaaa
    }
    if childYangName == "cSipCfgVoiceServiceVoip" {
        return &cISCOSIPUAMIB.Csipcfgvoiceservicevoip
    }
    if childYangName == "cSipStatsInfo" {
        return &cISCOSIPUAMIB.Csipstatsinfo
    }
    if childYangName == "cSipStatsSuccess" {
        return &cISCOSIPUAMIB.Csipstatssuccess
    }
    if childYangName == "cSipStatsRedirect" {
        return &cISCOSIPUAMIB.Csipstatsredirect
    }
    if childYangName == "cSipStatsErrClient" {
        return &cISCOSIPUAMIB.Csipstatserrclient
    }
    if childYangName == "cSipStatsErrServer" {
        return &cISCOSIPUAMIB.Csipstatserrserver
    }
    if childYangName == "cSipStatsGlobalFail" {
        return &cISCOSIPUAMIB.Csipstatsglobalfail
    }
    if childYangName == "cSipStatsTraffic" {
        return &cISCOSIPUAMIB.Csipstatstraffic
    }
    if childYangName == "cSipStatsRetry" {
        return &cISCOSIPUAMIB.Csipstatsretry
    }
    if childYangName == "cSipStatsMisc" {
        return &cISCOSIPUAMIB.Csipstatsmisc
    }
    if childYangName == "cSipStatsConnection" {
        return &cISCOSIPUAMIB.Csipstatsconnection
    }
    if childYangName == "cSipCfgEarlyMediaTable" {
        return &cISCOSIPUAMIB.Csipcfgearlymediatable
    }
    if childYangName == "cSipCfgBindSourceAddrTable" {
        return &cISCOSIPUAMIB.Csipcfgbindsourceaddrtable
    }
    if childYangName == "cSipCfgPeerTable" {
        return &cISCOSIPUAMIB.Csipcfgpeertable
    }
    if childYangName == "cSipCfgStatusCauseTable" {
        return &cISCOSIPUAMIB.Csipcfgstatuscausetable
    }
    if childYangName == "cSipCfgCauseStatusTable" {
        return &cISCOSIPUAMIB.Csipcfgcausestatustable
    }
    if childYangName == "cSipStatsSuccessOkTable" {
        return &cISCOSIPUAMIB.Csipstatssuccessoktable
    }
    return nil
}

func (cISCOSIPUAMIB *CISCOSIPUAMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cSipCfgBase"] = &cISCOSIPUAMIB.Csipcfgbase
    children["cSipCfgTimer"] = &cISCOSIPUAMIB.Csipcfgtimer
    children["cSipCfgRetry"] = &cISCOSIPUAMIB.Csipcfgretry
    children["cSipCfgPeer"] = &cISCOSIPUAMIB.Csipcfgpeer
    children["cSipCfgAaa"] = &cISCOSIPUAMIB.Csipcfgaaa
    children["cSipCfgVoiceServiceVoip"] = &cISCOSIPUAMIB.Csipcfgvoiceservicevoip
    children["cSipStatsInfo"] = &cISCOSIPUAMIB.Csipstatsinfo
    children["cSipStatsSuccess"] = &cISCOSIPUAMIB.Csipstatssuccess
    children["cSipStatsRedirect"] = &cISCOSIPUAMIB.Csipstatsredirect
    children["cSipStatsErrClient"] = &cISCOSIPUAMIB.Csipstatserrclient
    children["cSipStatsErrServer"] = &cISCOSIPUAMIB.Csipstatserrserver
    children["cSipStatsGlobalFail"] = &cISCOSIPUAMIB.Csipstatsglobalfail
    children["cSipStatsTraffic"] = &cISCOSIPUAMIB.Csipstatstraffic
    children["cSipStatsRetry"] = &cISCOSIPUAMIB.Csipstatsretry
    children["cSipStatsMisc"] = &cISCOSIPUAMIB.Csipstatsmisc
    children["cSipStatsConnection"] = &cISCOSIPUAMIB.Csipstatsconnection
    children["cSipCfgEarlyMediaTable"] = &cISCOSIPUAMIB.Csipcfgearlymediatable
    children["cSipCfgBindSourceAddrTable"] = &cISCOSIPUAMIB.Csipcfgbindsourceaddrtable
    children["cSipCfgPeerTable"] = &cISCOSIPUAMIB.Csipcfgpeertable
    children["cSipCfgStatusCauseTable"] = &cISCOSIPUAMIB.Csipcfgstatuscausetable
    children["cSipCfgCauseStatusTable"] = &cISCOSIPUAMIB.Csipcfgcausestatustable
    children["cSipStatsSuccessOkTable"] = &cISCOSIPUAMIB.Csipstatssuccessoktable
    return children
}

func (cISCOSIPUAMIB *CISCOSIPUAMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOSIPUAMIB *CISCOSIPUAMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOSIPUAMIB *CISCOSIPUAMIB) GetYangName() string { return "CISCO-SIP-UA-MIB" }

func (cISCOSIPUAMIB *CISCOSIPUAMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOSIPUAMIB *CISCOSIPUAMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOSIPUAMIB *CISCOSIPUAMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOSIPUAMIB *CISCOSIPUAMIB) SetParent(parent types.Entity) { cISCOSIPUAMIB.parent = parent }

func (cISCOSIPUAMIB *CISCOSIPUAMIB) GetParent() types.Entity { return cISCOSIPUAMIB.parent }

func (cISCOSIPUAMIB *CISCOSIPUAMIB) GetParentYangName() string { return "CISCO-SIP-UA-MIB" }

// CISCOSIPUAMIB_Csipcfgbase
type CISCOSIPUAMIB_Csipcfgbase struct {
    parent types.Entity
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

func (csipcfgbase *CISCOSIPUAMIB_Csipcfgbase) GetFilter() yfilter.YFilter { return csipcfgbase.YFilter }

func (csipcfgbase *CISCOSIPUAMIB_Csipcfgbase) SetFilter(yf yfilter.YFilter) { csipcfgbase.YFilter = yf }

func (csipcfgbase *CISCOSIPUAMIB_Csipcfgbase) GetGoName(yname string) string {
    if yname == "cSipCfgVersion" { return "Csipcfgversion" }
    if yname == "cSipCfgTransport" { return "Csipcfgtransport" }
    if yname == "cSipCfgUserLocationServerAddr" { return "Csipcfguserlocationserveraddr" }
    if yname == "cSipCfgMaxForwards" { return "Csipcfgmaxforwards" }
    if yname == "cSipCfgBindSrcAddrInterface" { return "Csipcfgbindsrcaddrinterface" }
    if yname == "cSipCfgBindSrcAddrScope" { return "Csipcfgbindsrcaddrscope" }
    if yname == "cSipCfgDnsSrvQueryStringFormat" { return "Csipcfgdnssrvquerystringformat" }
    if yname == "cSipCfgRedirectionDisabled" { return "Csipcfgredirectiondisabled" }
    if yname == "cSipCfgSymNatEnabled" { return "Csipcfgsymnatenabled" }
    if yname == "cSipCfgSymNatDirectionRole" { return "Csipcfgsymnatdirectionrole" }
    if yname == "cSipCfgSuspendResumeEnabled" { return "Csipcfgsuspendresumeenabled" }
    if yname == "cSipCfgOfferCallHold" { return "Csipcfgoffercallhold" }
    if yname == "cSipCfgReasonHeaderOveride" { return "Csipcfgreasonheaderoveride" }
    if yname == "cSipCfgMaximumForwards" { return "Csipcfgmaximumforwards" }
    return ""
}

func (csipcfgbase *CISCOSIPUAMIB_Csipcfgbase) GetSegmentPath() string {
    return "cSipCfgBase"
}

func (csipcfgbase *CISCOSIPUAMIB_Csipcfgbase) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csipcfgbase *CISCOSIPUAMIB_Csipcfgbase) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csipcfgbase *CISCOSIPUAMIB_Csipcfgbase) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cSipCfgVersion"] = csipcfgbase.Csipcfgversion
    leafs["cSipCfgTransport"] = csipcfgbase.Csipcfgtransport
    leafs["cSipCfgUserLocationServerAddr"] = csipcfgbase.Csipcfguserlocationserveraddr
    leafs["cSipCfgMaxForwards"] = csipcfgbase.Csipcfgmaxforwards
    leafs["cSipCfgBindSrcAddrInterface"] = csipcfgbase.Csipcfgbindsrcaddrinterface
    leafs["cSipCfgBindSrcAddrScope"] = csipcfgbase.Csipcfgbindsrcaddrscope
    leafs["cSipCfgDnsSrvQueryStringFormat"] = csipcfgbase.Csipcfgdnssrvquerystringformat
    leafs["cSipCfgRedirectionDisabled"] = csipcfgbase.Csipcfgredirectiondisabled
    leafs["cSipCfgSymNatEnabled"] = csipcfgbase.Csipcfgsymnatenabled
    leafs["cSipCfgSymNatDirectionRole"] = csipcfgbase.Csipcfgsymnatdirectionrole
    leafs["cSipCfgSuspendResumeEnabled"] = csipcfgbase.Csipcfgsuspendresumeenabled
    leafs["cSipCfgOfferCallHold"] = csipcfgbase.Csipcfgoffercallhold
    leafs["cSipCfgReasonHeaderOveride"] = csipcfgbase.Csipcfgreasonheaderoveride
    leafs["cSipCfgMaximumForwards"] = csipcfgbase.Csipcfgmaximumforwards
    return leafs
}

func (csipcfgbase *CISCOSIPUAMIB_Csipcfgbase) GetBundleName() string { return "cisco_ios_xe" }

func (csipcfgbase *CISCOSIPUAMIB_Csipcfgbase) GetYangName() string { return "cSipCfgBase" }

func (csipcfgbase *CISCOSIPUAMIB_Csipcfgbase) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csipcfgbase *CISCOSIPUAMIB_Csipcfgbase) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csipcfgbase *CISCOSIPUAMIB_Csipcfgbase) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csipcfgbase *CISCOSIPUAMIB_Csipcfgbase) SetParent(parent types.Entity) { csipcfgbase.parent = parent }

func (csipcfgbase *CISCOSIPUAMIB_Csipcfgbase) GetParent() types.Entity { return csipcfgbase.parent }

func (csipcfgbase *CISCOSIPUAMIB_Csipcfgbase) GetParentYangName() string { return "CISCO-SIP-UA-MIB" }

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
    parent types.Entity
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

func (csipcfgtimer *CISCOSIPUAMIB_Csipcfgtimer) GetFilter() yfilter.YFilter { return csipcfgtimer.YFilter }

func (csipcfgtimer *CISCOSIPUAMIB_Csipcfgtimer) SetFilter(yf yfilter.YFilter) { csipcfgtimer.YFilter = yf }

func (csipcfgtimer *CISCOSIPUAMIB_Csipcfgtimer) GetGoName(yname string) string {
    if yname == "cSipCfgTimerTrying" { return "Csipcfgtimertrying" }
    if yname == "cSipCfgTimerExpires" { return "Csipcfgtimerexpires" }
    if yname == "cSipCfgTimerConnect" { return "Csipcfgtimerconnect" }
    if yname == "cSipCfgTimerDisconnect" { return "Csipcfgtimerdisconnect" }
    if yname == "cSipCfgTimerPrack" { return "Csipcfgtimerprack" }
    if yname == "cSipCfgTimerComet" { return "Csipcfgtimercomet" }
    if yname == "cSipCfgTimerReliableRsp" { return "Csipcfgtimerreliablersp" }
    if yname == "cSipCfgTimerNotify" { return "Csipcfgtimernotify" }
    if yname == "cSipCfgTimerRefer" { return "Csipcfgtimerrefer" }
    if yname == "cSipCfgTimerSessionTimer" { return "Csipcfgtimersessiontimer" }
    if yname == "cSipCfgTimerHold" { return "Csipcfgtimerhold" }
    if yname == "cSipCfgTimerInfo" { return "Csipcfgtimerinfo" }
    if yname == "cSipCfgTimerConnectionAging" { return "Csipcfgtimerconnectionaging" }
    if yname == "cSipCfgTimerBufferInvite" { return "Csipcfgtimerbufferinvite" }
    return ""
}

func (csipcfgtimer *CISCOSIPUAMIB_Csipcfgtimer) GetSegmentPath() string {
    return "cSipCfgTimer"
}

func (csipcfgtimer *CISCOSIPUAMIB_Csipcfgtimer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csipcfgtimer *CISCOSIPUAMIB_Csipcfgtimer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csipcfgtimer *CISCOSIPUAMIB_Csipcfgtimer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cSipCfgTimerTrying"] = csipcfgtimer.Csipcfgtimertrying
    leafs["cSipCfgTimerExpires"] = csipcfgtimer.Csipcfgtimerexpires
    leafs["cSipCfgTimerConnect"] = csipcfgtimer.Csipcfgtimerconnect
    leafs["cSipCfgTimerDisconnect"] = csipcfgtimer.Csipcfgtimerdisconnect
    leafs["cSipCfgTimerPrack"] = csipcfgtimer.Csipcfgtimerprack
    leafs["cSipCfgTimerComet"] = csipcfgtimer.Csipcfgtimercomet
    leafs["cSipCfgTimerReliableRsp"] = csipcfgtimer.Csipcfgtimerreliablersp
    leafs["cSipCfgTimerNotify"] = csipcfgtimer.Csipcfgtimernotify
    leafs["cSipCfgTimerRefer"] = csipcfgtimer.Csipcfgtimerrefer
    leafs["cSipCfgTimerSessionTimer"] = csipcfgtimer.Csipcfgtimersessiontimer
    leafs["cSipCfgTimerHold"] = csipcfgtimer.Csipcfgtimerhold
    leafs["cSipCfgTimerInfo"] = csipcfgtimer.Csipcfgtimerinfo
    leafs["cSipCfgTimerConnectionAging"] = csipcfgtimer.Csipcfgtimerconnectionaging
    leafs["cSipCfgTimerBufferInvite"] = csipcfgtimer.Csipcfgtimerbufferinvite
    return leafs
}

func (csipcfgtimer *CISCOSIPUAMIB_Csipcfgtimer) GetBundleName() string { return "cisco_ios_xe" }

func (csipcfgtimer *CISCOSIPUAMIB_Csipcfgtimer) GetYangName() string { return "cSipCfgTimer" }

func (csipcfgtimer *CISCOSIPUAMIB_Csipcfgtimer) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csipcfgtimer *CISCOSIPUAMIB_Csipcfgtimer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csipcfgtimer *CISCOSIPUAMIB_Csipcfgtimer) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csipcfgtimer *CISCOSIPUAMIB_Csipcfgtimer) SetParent(parent types.Entity) { csipcfgtimer.parent = parent }

func (csipcfgtimer *CISCOSIPUAMIB_Csipcfgtimer) GetParent() types.Entity { return csipcfgtimer.parent }

func (csipcfgtimer *CISCOSIPUAMIB_Csipcfgtimer) GetParentYangName() string { return "CISCO-SIP-UA-MIB" }

// CISCOSIPUAMIB_Csipcfgretry
type CISCOSIPUAMIB_Csipcfgretry struct {
    parent types.Entity
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

func (csipcfgretry *CISCOSIPUAMIB_Csipcfgretry) GetFilter() yfilter.YFilter { return csipcfgretry.YFilter }

func (csipcfgretry *CISCOSIPUAMIB_Csipcfgretry) SetFilter(yf yfilter.YFilter) { csipcfgretry.YFilter = yf }

func (csipcfgretry *CISCOSIPUAMIB_Csipcfgretry) GetGoName(yname string) string {
    if yname == "cSipCfgRetryInvite" { return "Csipcfgretryinvite" }
    if yname == "cSipCfgRetryBye" { return "Csipcfgretrybye" }
    if yname == "cSipCfgRetryCancel" { return "Csipcfgretrycancel" }
    if yname == "cSipCfgRetryRegister" { return "Csipcfgretryregister" }
    if yname == "cSipCfgRetryResponse" { return "Csipcfgretryresponse" }
    if yname == "cSipCfgRetryPrack" { return "Csipcfgretryprack" }
    if yname == "cSipCfgRetryComet" { return "Csipcfgretrycomet" }
    if yname == "cSipCfgRetryReliableRsp" { return "Csipcfgretryreliablersp" }
    if yname == "cSipCfgRetryNotify" { return "Csipcfgretrynotify" }
    if yname == "cSipCfgRetryRefer" { return "Csipcfgretryrefer" }
    if yname == "cSipCfgRetryInfo" { return "Csipcfgretryinfo" }
    if yname == "cSipCfgRetrySubscribe" { return "Csipcfgretrysubscribe" }
    return ""
}

func (csipcfgretry *CISCOSIPUAMIB_Csipcfgretry) GetSegmentPath() string {
    return "cSipCfgRetry"
}

func (csipcfgretry *CISCOSIPUAMIB_Csipcfgretry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csipcfgretry *CISCOSIPUAMIB_Csipcfgretry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csipcfgretry *CISCOSIPUAMIB_Csipcfgretry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cSipCfgRetryInvite"] = csipcfgretry.Csipcfgretryinvite
    leafs["cSipCfgRetryBye"] = csipcfgretry.Csipcfgretrybye
    leafs["cSipCfgRetryCancel"] = csipcfgretry.Csipcfgretrycancel
    leafs["cSipCfgRetryRegister"] = csipcfgretry.Csipcfgretryregister
    leafs["cSipCfgRetryResponse"] = csipcfgretry.Csipcfgretryresponse
    leafs["cSipCfgRetryPrack"] = csipcfgretry.Csipcfgretryprack
    leafs["cSipCfgRetryComet"] = csipcfgretry.Csipcfgretrycomet
    leafs["cSipCfgRetryReliableRsp"] = csipcfgretry.Csipcfgretryreliablersp
    leafs["cSipCfgRetryNotify"] = csipcfgretry.Csipcfgretrynotify
    leafs["cSipCfgRetryRefer"] = csipcfgretry.Csipcfgretryrefer
    leafs["cSipCfgRetryInfo"] = csipcfgretry.Csipcfgretryinfo
    leafs["cSipCfgRetrySubscribe"] = csipcfgretry.Csipcfgretrysubscribe
    return leafs
}

func (csipcfgretry *CISCOSIPUAMIB_Csipcfgretry) GetBundleName() string { return "cisco_ios_xe" }

func (csipcfgretry *CISCOSIPUAMIB_Csipcfgretry) GetYangName() string { return "cSipCfgRetry" }

func (csipcfgretry *CISCOSIPUAMIB_Csipcfgretry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csipcfgretry *CISCOSIPUAMIB_Csipcfgretry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csipcfgretry *CISCOSIPUAMIB_Csipcfgretry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csipcfgretry *CISCOSIPUAMIB_Csipcfgretry) SetParent(parent types.Entity) { csipcfgretry.parent = parent }

func (csipcfgretry *CISCOSIPUAMIB_Csipcfgretry) GetParent() types.Entity { return csipcfgretry.parent }

func (csipcfgretry *CISCOSIPUAMIB_Csipcfgretry) GetParentYangName() string { return "CISCO-SIP-UA-MIB" }

// CISCOSIPUAMIB_Csipcfgpeer
type CISCOSIPUAMIB_Csipcfgpeer struct {
    parent types.Entity
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

func (csipcfgpeer *CISCOSIPUAMIB_Csipcfgpeer) GetFilter() yfilter.YFilter { return csipcfgpeer.YFilter }

func (csipcfgpeer *CISCOSIPUAMIB_Csipcfgpeer) SetFilter(yf yfilter.YFilter) { csipcfgpeer.YFilter = yf }

func (csipcfgpeer *CISCOSIPUAMIB_Csipcfgpeer) GetGoName(yname string) string {
    if yname == "cSipCfgOutSessionTransport" { return "Csipcfgoutsessiontransport" }
    if yname == "cSipCfgReliable1xxRspStr" { return "Csipcfgreliable1Xxrspstr" }
    if yname == "cSipCfgReliable1xxRspHdr" { return "Csipcfgreliable1Xxrsphdr" }
    if yname == "cSipCfgUrlType" { return "Csipcfgurltype" }
    return ""
}

func (csipcfgpeer *CISCOSIPUAMIB_Csipcfgpeer) GetSegmentPath() string {
    return "cSipCfgPeer"
}

func (csipcfgpeer *CISCOSIPUAMIB_Csipcfgpeer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csipcfgpeer *CISCOSIPUAMIB_Csipcfgpeer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csipcfgpeer *CISCOSIPUAMIB_Csipcfgpeer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cSipCfgOutSessionTransport"] = csipcfgpeer.Csipcfgoutsessiontransport
    leafs["cSipCfgReliable1xxRspStr"] = csipcfgpeer.Csipcfgreliable1Xxrspstr
    leafs["cSipCfgReliable1xxRspHdr"] = csipcfgpeer.Csipcfgreliable1Xxrsphdr
    leafs["cSipCfgUrlType"] = csipcfgpeer.Csipcfgurltype
    return leafs
}

func (csipcfgpeer *CISCOSIPUAMIB_Csipcfgpeer) GetBundleName() string { return "cisco_ios_xe" }

func (csipcfgpeer *CISCOSIPUAMIB_Csipcfgpeer) GetYangName() string { return "cSipCfgPeer" }

func (csipcfgpeer *CISCOSIPUAMIB_Csipcfgpeer) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csipcfgpeer *CISCOSIPUAMIB_Csipcfgpeer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csipcfgpeer *CISCOSIPUAMIB_Csipcfgpeer) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csipcfgpeer *CISCOSIPUAMIB_Csipcfgpeer) SetParent(parent types.Entity) { csipcfgpeer.parent = parent }

func (csipcfgpeer *CISCOSIPUAMIB_Csipcfgpeer) GetParent() types.Entity { return csipcfgpeer.parent }

func (csipcfgpeer *CISCOSIPUAMIB_Csipcfgpeer) GetParentYangName() string { return "CISCO-SIP-UA-MIB" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // This object specifies the source of the information used to populate the
    // username attribute of AAA billing records. The type is Csipcfgaaausername.
    Csipcfgaaausername interface{}
}

func (csipcfgaaa *CISCOSIPUAMIB_Csipcfgaaa) GetFilter() yfilter.YFilter { return csipcfgaaa.YFilter }

func (csipcfgaaa *CISCOSIPUAMIB_Csipcfgaaa) SetFilter(yf yfilter.YFilter) { csipcfgaaa.YFilter = yf }

func (csipcfgaaa *CISCOSIPUAMIB_Csipcfgaaa) GetGoName(yname string) string {
    if yname == "cSipCfgAaaUsername" { return "Csipcfgaaausername" }
    return ""
}

func (csipcfgaaa *CISCOSIPUAMIB_Csipcfgaaa) GetSegmentPath() string {
    return "cSipCfgAaa"
}

func (csipcfgaaa *CISCOSIPUAMIB_Csipcfgaaa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csipcfgaaa *CISCOSIPUAMIB_Csipcfgaaa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csipcfgaaa *CISCOSIPUAMIB_Csipcfgaaa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cSipCfgAaaUsername"] = csipcfgaaa.Csipcfgaaausername
    return leafs
}

func (csipcfgaaa *CISCOSIPUAMIB_Csipcfgaaa) GetBundleName() string { return "cisco_ios_xe" }

func (csipcfgaaa *CISCOSIPUAMIB_Csipcfgaaa) GetYangName() string { return "cSipCfgAaa" }

func (csipcfgaaa *CISCOSIPUAMIB_Csipcfgaaa) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csipcfgaaa *CISCOSIPUAMIB_Csipcfgaaa) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csipcfgaaa *CISCOSIPUAMIB_Csipcfgaaa) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csipcfgaaa *CISCOSIPUAMIB_Csipcfgaaa) SetParent(parent types.Entity) { csipcfgaaa.parent = parent }

func (csipcfgaaa *CISCOSIPUAMIB_Csipcfgaaa) GetParent() types.Entity { return csipcfgaaa.parent }

func (csipcfgaaa *CISCOSIPUAMIB_Csipcfgaaa) GetParentYangName() string { return "CISCO-SIP-UA-MIB" }

// CISCOSIPUAMIB_Csipcfgaaa_Csipcfgaaausername represents populate the username attribute of AAA billing records.
type CISCOSIPUAMIB_Csipcfgaaa_Csipcfgaaausername string

const (
    CISCOSIPUAMIB_Csipcfgaaa_Csipcfgaaausername_callingNumber CISCOSIPUAMIB_Csipcfgaaa_Csipcfgaaausername = "callingNumber"

    CISCOSIPUAMIB_Csipcfgaaa_Csipcfgaaausername_proxyAuth CISCOSIPUAMIB_Csipcfgaaa_Csipcfgaaausername = "proxyAuth"
)

// CISCOSIPUAMIB_Csipcfgvoiceservicevoip
type CISCOSIPUAMIB_Csipcfgvoiceservicevoip struct {
    parent types.Entity
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

func (csipcfgvoiceservicevoip *CISCOSIPUAMIB_Csipcfgvoiceservicevoip) GetFilter() yfilter.YFilter { return csipcfgvoiceservicevoip.YFilter }

func (csipcfgvoiceservicevoip *CISCOSIPUAMIB_Csipcfgvoiceservicevoip) SetFilter(yf yfilter.YFilter) { csipcfgvoiceservicevoip.YFilter = yf }

func (csipcfgvoiceservicevoip *CISCOSIPUAMIB_Csipcfgvoiceservicevoip) GetGoName(yname string) string {
    if yname == "cSipCfgHeaderPassingEnabled" { return "Csipcfgheaderpassingenabled" }
    if yname == "cSipCfgMaxSubscriptionAccept" { return "Csipcfgmaxsubscriptionaccept" }
    if yname == "cSipCfgMaxSubscriptionOriginate" { return "Csipcfgmaxsubscriptionoriginate" }
    if yname == "cSipCfgSwitchTransportEnabled" { return "Csipcfgswitchtransportenabled" }
    return ""
}

func (csipcfgvoiceservicevoip *CISCOSIPUAMIB_Csipcfgvoiceservicevoip) GetSegmentPath() string {
    return "cSipCfgVoiceServiceVoip"
}

func (csipcfgvoiceservicevoip *CISCOSIPUAMIB_Csipcfgvoiceservicevoip) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csipcfgvoiceservicevoip *CISCOSIPUAMIB_Csipcfgvoiceservicevoip) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csipcfgvoiceservicevoip *CISCOSIPUAMIB_Csipcfgvoiceservicevoip) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cSipCfgHeaderPassingEnabled"] = csipcfgvoiceservicevoip.Csipcfgheaderpassingenabled
    leafs["cSipCfgMaxSubscriptionAccept"] = csipcfgvoiceservicevoip.Csipcfgmaxsubscriptionaccept
    leafs["cSipCfgMaxSubscriptionOriginate"] = csipcfgvoiceservicevoip.Csipcfgmaxsubscriptionoriginate
    leafs["cSipCfgSwitchTransportEnabled"] = csipcfgvoiceservicevoip.Csipcfgswitchtransportenabled
    return leafs
}

func (csipcfgvoiceservicevoip *CISCOSIPUAMIB_Csipcfgvoiceservicevoip) GetBundleName() string { return "cisco_ios_xe" }

func (csipcfgvoiceservicevoip *CISCOSIPUAMIB_Csipcfgvoiceservicevoip) GetYangName() string { return "cSipCfgVoiceServiceVoip" }

func (csipcfgvoiceservicevoip *CISCOSIPUAMIB_Csipcfgvoiceservicevoip) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csipcfgvoiceservicevoip *CISCOSIPUAMIB_Csipcfgvoiceservicevoip) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csipcfgvoiceservicevoip *CISCOSIPUAMIB_Csipcfgvoiceservicevoip) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csipcfgvoiceservicevoip *CISCOSIPUAMIB_Csipcfgvoiceservicevoip) SetParent(parent types.Entity) { csipcfgvoiceservicevoip.parent = parent }

func (csipcfgvoiceservicevoip *CISCOSIPUAMIB_Csipcfgvoiceservicevoip) GetParent() types.Entity { return csipcfgvoiceservicevoip.parent }

func (csipcfgvoiceservicevoip *CISCOSIPUAMIB_Csipcfgvoiceservicevoip) GetParentYangName() string { return "CISCO-SIP-UA-MIB" }

// CISCOSIPUAMIB_Csipstatsinfo
type CISCOSIPUAMIB_Csipstatsinfo struct {
    parent types.Entity
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

func (csipstatsinfo *CISCOSIPUAMIB_Csipstatsinfo) GetFilter() yfilter.YFilter { return csipstatsinfo.YFilter }

func (csipstatsinfo *CISCOSIPUAMIB_Csipstatsinfo) SetFilter(yf yfilter.YFilter) { csipstatsinfo.YFilter = yf }

func (csipstatsinfo *CISCOSIPUAMIB_Csipstatsinfo) GetGoName(yname string) string {
    if yname == "cSipStatsInfoTryingIns" { return "Csipstatsinfotryingins" }
    if yname == "cSipStatsInfoTryingOuts" { return "Csipstatsinfotryingouts" }
    if yname == "cSipStatsInfoRingingIns" { return "Csipstatsinforingingins" }
    if yname == "cSipStatsInfoRingingOuts" { return "Csipstatsinforingingouts" }
    if yname == "cSipStatsInfoForwardedIns" { return "Csipstatsinfoforwardedins" }
    if yname == "cSipStatsInfoForwardedOuts" { return "Csipstatsinfoforwardedouts" }
    if yname == "cSipStatsInfoQueuedIns" { return "Csipstatsinfoqueuedins" }
    if yname == "cSipStatsInfoQueuedOuts" { return "Csipstatsinfoqueuedouts" }
    if yname == "cSipStatsInfoSessionProgIns" { return "Csipstatsinfosessionprogins" }
    if yname == "cSipStatsInfoSessionProgOuts" { return "Csipstatsinfosessionprogouts" }
    return ""
}

func (csipstatsinfo *CISCOSIPUAMIB_Csipstatsinfo) GetSegmentPath() string {
    return "cSipStatsInfo"
}

func (csipstatsinfo *CISCOSIPUAMIB_Csipstatsinfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csipstatsinfo *CISCOSIPUAMIB_Csipstatsinfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csipstatsinfo *CISCOSIPUAMIB_Csipstatsinfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cSipStatsInfoTryingIns"] = csipstatsinfo.Csipstatsinfotryingins
    leafs["cSipStatsInfoTryingOuts"] = csipstatsinfo.Csipstatsinfotryingouts
    leafs["cSipStatsInfoRingingIns"] = csipstatsinfo.Csipstatsinforingingins
    leafs["cSipStatsInfoRingingOuts"] = csipstatsinfo.Csipstatsinforingingouts
    leafs["cSipStatsInfoForwardedIns"] = csipstatsinfo.Csipstatsinfoforwardedins
    leafs["cSipStatsInfoForwardedOuts"] = csipstatsinfo.Csipstatsinfoforwardedouts
    leafs["cSipStatsInfoQueuedIns"] = csipstatsinfo.Csipstatsinfoqueuedins
    leafs["cSipStatsInfoQueuedOuts"] = csipstatsinfo.Csipstatsinfoqueuedouts
    leafs["cSipStatsInfoSessionProgIns"] = csipstatsinfo.Csipstatsinfosessionprogins
    leafs["cSipStatsInfoSessionProgOuts"] = csipstatsinfo.Csipstatsinfosessionprogouts
    return leafs
}

func (csipstatsinfo *CISCOSIPUAMIB_Csipstatsinfo) GetBundleName() string { return "cisco_ios_xe" }

func (csipstatsinfo *CISCOSIPUAMIB_Csipstatsinfo) GetYangName() string { return "cSipStatsInfo" }

func (csipstatsinfo *CISCOSIPUAMIB_Csipstatsinfo) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csipstatsinfo *CISCOSIPUAMIB_Csipstatsinfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csipstatsinfo *CISCOSIPUAMIB_Csipstatsinfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csipstatsinfo *CISCOSIPUAMIB_Csipstatsinfo) SetParent(parent types.Entity) { csipstatsinfo.parent = parent }

func (csipstatsinfo *CISCOSIPUAMIB_Csipstatsinfo) GetParent() types.Entity { return csipstatsinfo.parent }

func (csipstatsinfo *CISCOSIPUAMIB_Csipstatsinfo) GetParentYangName() string { return "CISCO-SIP-UA-MIB" }

// CISCOSIPUAMIB_Csipstatssuccess
type CISCOSIPUAMIB_Csipstatssuccess struct {
    parent types.Entity
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

func (csipstatssuccess *CISCOSIPUAMIB_Csipstatssuccess) GetFilter() yfilter.YFilter { return csipstatssuccess.YFilter }

func (csipstatssuccess *CISCOSIPUAMIB_Csipstatssuccess) SetFilter(yf yfilter.YFilter) { csipstatssuccess.YFilter = yf }

func (csipstatssuccess *CISCOSIPUAMIB_Csipstatssuccess) GetGoName(yname string) string {
    if yname == "cSipStatsSuccessOkIns" { return "Csipstatssuccessokins" }
    if yname == "cSipStatsSuccessOkOuts" { return "Csipstatssuccessokouts" }
    if yname == "cSipStatsSuccessAcceptedIns" { return "Csipstatssuccessacceptedins" }
    if yname == "cSipStatsSuccessAcceptedOuts" { return "Csipstatssuccessacceptedouts" }
    return ""
}

func (csipstatssuccess *CISCOSIPUAMIB_Csipstatssuccess) GetSegmentPath() string {
    return "cSipStatsSuccess"
}

func (csipstatssuccess *CISCOSIPUAMIB_Csipstatssuccess) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csipstatssuccess *CISCOSIPUAMIB_Csipstatssuccess) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csipstatssuccess *CISCOSIPUAMIB_Csipstatssuccess) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cSipStatsSuccessOkIns"] = csipstatssuccess.Csipstatssuccessokins
    leafs["cSipStatsSuccessOkOuts"] = csipstatssuccess.Csipstatssuccessokouts
    leafs["cSipStatsSuccessAcceptedIns"] = csipstatssuccess.Csipstatssuccessacceptedins
    leafs["cSipStatsSuccessAcceptedOuts"] = csipstatssuccess.Csipstatssuccessacceptedouts
    return leafs
}

func (csipstatssuccess *CISCOSIPUAMIB_Csipstatssuccess) GetBundleName() string { return "cisco_ios_xe" }

func (csipstatssuccess *CISCOSIPUAMIB_Csipstatssuccess) GetYangName() string { return "cSipStatsSuccess" }

func (csipstatssuccess *CISCOSIPUAMIB_Csipstatssuccess) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csipstatssuccess *CISCOSIPUAMIB_Csipstatssuccess) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csipstatssuccess *CISCOSIPUAMIB_Csipstatssuccess) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csipstatssuccess *CISCOSIPUAMIB_Csipstatssuccess) SetParent(parent types.Entity) { csipstatssuccess.parent = parent }

func (csipstatssuccess *CISCOSIPUAMIB_Csipstatssuccess) GetParent() types.Entity { return csipstatssuccess.parent }

func (csipstatssuccess *CISCOSIPUAMIB_Csipstatssuccess) GetParentYangName() string { return "CISCO-SIP-UA-MIB" }

// CISCOSIPUAMIB_Csipstatsredirect
type CISCOSIPUAMIB_Csipstatsredirect struct {
    parent types.Entity
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

func (csipstatsredirect *CISCOSIPUAMIB_Csipstatsredirect) GetFilter() yfilter.YFilter { return csipstatsredirect.YFilter }

func (csipstatsredirect *CISCOSIPUAMIB_Csipstatsredirect) SetFilter(yf yfilter.YFilter) { csipstatsredirect.YFilter = yf }

func (csipstatsredirect *CISCOSIPUAMIB_Csipstatsredirect) GetGoName(yname string) string {
    if yname == "cSipStatsRedirMultipleChoices" { return "Csipstatsredirmultiplechoices" }
    if yname == "cSipStatsRedirMovedPerms" { return "Csipstatsredirmovedperms" }
    if yname == "cSipStatsRedirMovedTemps" { return "Csipstatsredirmovedtemps" }
    if yname == "cSipStatsRedirSeeOthers" { return "Csipstatsredirseeothers" }
    if yname == "cSipStatsRedirUseProxys" { return "Csipstatsrediruseproxys" }
    if yname == "cSipStatsRedirAltServices" { return "Csipstatsrediraltservices" }
    if yname == "cSipStatsRedirMovedTempsIns" { return "Csipstatsredirmovedtempsins" }
    if yname == "cSipStatsRedirMovedTempsOuts" { return "Csipstatsredirmovedtempsouts" }
    return ""
}

func (csipstatsredirect *CISCOSIPUAMIB_Csipstatsredirect) GetSegmentPath() string {
    return "cSipStatsRedirect"
}

func (csipstatsredirect *CISCOSIPUAMIB_Csipstatsredirect) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csipstatsredirect *CISCOSIPUAMIB_Csipstatsredirect) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csipstatsredirect *CISCOSIPUAMIB_Csipstatsredirect) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cSipStatsRedirMultipleChoices"] = csipstatsredirect.Csipstatsredirmultiplechoices
    leafs["cSipStatsRedirMovedPerms"] = csipstatsredirect.Csipstatsredirmovedperms
    leafs["cSipStatsRedirMovedTemps"] = csipstatsredirect.Csipstatsredirmovedtemps
    leafs["cSipStatsRedirSeeOthers"] = csipstatsredirect.Csipstatsredirseeothers
    leafs["cSipStatsRedirUseProxys"] = csipstatsredirect.Csipstatsrediruseproxys
    leafs["cSipStatsRedirAltServices"] = csipstatsredirect.Csipstatsrediraltservices
    leafs["cSipStatsRedirMovedTempsIns"] = csipstatsredirect.Csipstatsredirmovedtempsins
    leafs["cSipStatsRedirMovedTempsOuts"] = csipstatsredirect.Csipstatsredirmovedtempsouts
    return leafs
}

func (csipstatsredirect *CISCOSIPUAMIB_Csipstatsredirect) GetBundleName() string { return "cisco_ios_xe" }

func (csipstatsredirect *CISCOSIPUAMIB_Csipstatsredirect) GetYangName() string { return "cSipStatsRedirect" }

func (csipstatsredirect *CISCOSIPUAMIB_Csipstatsredirect) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csipstatsredirect *CISCOSIPUAMIB_Csipstatsredirect) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csipstatsredirect *CISCOSIPUAMIB_Csipstatsredirect) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csipstatsredirect *CISCOSIPUAMIB_Csipstatsredirect) SetParent(parent types.Entity) { csipstatsredirect.parent = parent }

func (csipstatsredirect *CISCOSIPUAMIB_Csipstatsredirect) GetParent() types.Entity { return csipstatsredirect.parent }

func (csipstatsredirect *CISCOSIPUAMIB_Csipstatsredirect) GetParentYangName() string { return "CISCO-SIP-UA-MIB" }

// CISCOSIPUAMIB_Csipstatserrclient
type CISCOSIPUAMIB_Csipstatserrclient struct {
    parent types.Entity
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

func (csipstatserrclient *CISCOSIPUAMIB_Csipstatserrclient) GetFilter() yfilter.YFilter { return csipstatserrclient.YFilter }

func (csipstatserrclient *CISCOSIPUAMIB_Csipstatserrclient) SetFilter(yf yfilter.YFilter) { csipstatserrclient.YFilter = yf }

func (csipstatserrclient *CISCOSIPUAMIB_Csipstatserrclient) GetGoName(yname string) string {
    if yname == "cSipStatsClientBadRequestIns" { return "Csipstatsclientbadrequestins" }
    if yname == "cSipStatsClientBadRequestOuts" { return "Csipstatsclientbadrequestouts" }
    if yname == "cSipStatsClientUnauthorizedIns" { return "Csipstatsclientunauthorizedins" }
    if yname == "cSipStatsClientUnauthorizedOuts" { return "Csipstatsclientunauthorizedouts" }
    if yname == "cSipStatsClientPaymentReqdIns" { return "Csipstatsclientpaymentreqdins" }
    if yname == "cSipStatsClientPaymentReqdOuts" { return "Csipstatsclientpaymentreqdouts" }
    if yname == "cSipStatsClientForbiddenIns" { return "Csipstatsclientforbiddenins" }
    if yname == "cSipStatsClientForbiddenOuts" { return "Csipstatsclientforbiddenouts" }
    if yname == "cSipStatsClientNotFoundIns" { return "Csipstatsclientnotfoundins" }
    if yname == "cSipStatsClientNotFoundOuts" { return "Csipstatsclientnotfoundouts" }
    if yname == "cSipStatsClientMethNotAllowedIns" { return "Csipstatsclientmethnotallowedins" }
    if yname == "cSipStatsClientMethNotAllowedOuts" { return "Csipstatsclientmethnotallowedouts" }
    if yname == "cSipStatsClientNotAcceptableIns" { return "Csipstatsclientnotacceptableins" }
    if yname == "cSipStatsClientNotAcceptableOuts" { return "Csipstatsclientnotacceptableouts" }
    if yname == "cSipStatsClientProxyAuthReqdIns" { return "Csipstatsclientproxyauthreqdins" }
    if yname == "cSipStatsClientProxyAuthReqdOuts" { return "Csipstatsclientproxyauthreqdouts" }
    if yname == "cSipStatsClientReqTimeoutIns" { return "Csipstatsclientreqtimeoutins" }
    if yname == "cSipStatsClientReqTimeoutOuts" { return "Csipstatsclientreqtimeoutouts" }
    if yname == "cSipStatsClientConflictIns" { return "Csipstatsclientconflictins" }
    if yname == "cSipStatsClientConflictOuts" { return "Csipstatsclientconflictouts" }
    if yname == "cSipStatsClientGoneIns" { return "Csipstatsclientgoneins" }
    if yname == "cSipStatsClientGoneOuts" { return "Csipstatsclientgoneouts" }
    if yname == "cSipStatsClientLengthRequiredIns" { return "Csipstatsclientlengthrequiredins" }
    if yname == "cSipStatsClientLengthRequiredOuts" { return "Csipstatsclientlengthrequiredouts" }
    if yname == "cSipStatsClientReqEntTooLargeIns" { return "Csipstatsclientreqenttoolargeins" }
    if yname == "cSipStatsClientReqEntTooLargeOuts" { return "Csipstatsclientreqenttoolargeouts" }
    if yname == "cSipStatsClientReqURITooLargeIns" { return "Csipstatsclientrequritoolargeins" }
    if yname == "cSipStatsClientReqURITooLargeOuts" { return "Csipstatsclientrequritoolargeouts" }
    if yname == "cSipStatsClientNoSupMediaTypeIns" { return "Csipstatsclientnosupmediatypeins" }
    if yname == "cSipStatsClientNoSupMediaTypeOuts" { return "Csipstatsclientnosupmediatypeouts" }
    if yname == "cSipStatsClientBadExtensionIns" { return "Csipstatsclientbadextensionins" }
    if yname == "cSipStatsClientBadExtensionOuts" { return "Csipstatsclientbadextensionouts" }
    if yname == "cSipStatsClientTempNotAvailIns" { return "Csipstatsclienttempnotavailins" }
    if yname == "cSipStatsClientTempNotAvailOuts" { return "Csipstatsclienttempnotavailouts" }
    if yname == "cSipStatsClientCallLegNoExistIns" { return "Csipstatsclientcalllegnoexistins" }
    if yname == "cSipStatsClientCallLegNoExistOuts" { return "Csipstatsclientcalllegnoexistouts" }
    if yname == "cSipStatsClientLoopDetectedIns" { return "Csipstatsclientloopdetectedins" }
    if yname == "cSipStatsClientLoopDetectedOuts" { return "Csipstatsclientloopdetectedouts" }
    if yname == "cSipStatsClientTooManyHopsIns" { return "Csipstatsclienttoomanyhopsins" }
    if yname == "cSipStatsClientTooManyHopsOuts" { return "Csipstatsclienttoomanyhopsouts" }
    if yname == "cSipStatsClientAddrIncompleteIns" { return "Csipstatsclientaddrincompleteins" }
    if yname == "cSipStatsClientAddrIncompleteOuts" { return "Csipstatsclientaddrincompleteouts" }
    if yname == "cSipStatsClientAmbiguousIns" { return "Csipstatsclientambiguousins" }
    if yname == "cSipStatsClientAmbiguousOuts" { return "Csipstatsclientambiguousouts" }
    if yname == "cSipStatsClientBusyHereIns" { return "Csipstatsclientbusyhereins" }
    if yname == "cSipStatsClientBusyHereOuts" { return "Csipstatsclientbusyhereouts" }
    if yname == "cSipStatsClientReqTermIns" { return "Csipstatsclientreqtermins" }
    if yname == "cSipStatsClientReqTermOuts" { return "Csipstatsclientreqtermouts" }
    if yname == "cSipStatsClientNoAcceptHereIns" { return "Csipstatsclientnoaccepthereins" }
    if yname == "cSipStatsClientNoAcceptHereOuts" { return "Csipstatsclientnoaccepthereouts" }
    if yname == "cSipStatsClientBadEventIns" { return "Csipstatsclientbadeventins" }
    if yname == "cSipStatsClientBadEventOuts" { return "Csipstatsclientbadeventouts" }
    if yname == "cSipStatsClientSTTooSmallIns" { return "Csipstatsclientsttoosmallins" }
    if yname == "cSipStatsClientSTTooSmallOuts" { return "Csipstatsclientsttoosmallouts" }
    if yname == "cSipStatsClientReqPendingIns" { return "Csipstatsclientreqpendingins" }
    if yname == "cSipStatsClientReqPendingOuts" { return "Csipstatsclientreqpendingouts" }
    return ""
}

func (csipstatserrclient *CISCOSIPUAMIB_Csipstatserrclient) GetSegmentPath() string {
    return "cSipStatsErrClient"
}

func (csipstatserrclient *CISCOSIPUAMIB_Csipstatserrclient) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csipstatserrclient *CISCOSIPUAMIB_Csipstatserrclient) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csipstatserrclient *CISCOSIPUAMIB_Csipstatserrclient) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cSipStatsClientBadRequestIns"] = csipstatserrclient.Csipstatsclientbadrequestins
    leafs["cSipStatsClientBadRequestOuts"] = csipstatserrclient.Csipstatsclientbadrequestouts
    leafs["cSipStatsClientUnauthorizedIns"] = csipstatserrclient.Csipstatsclientunauthorizedins
    leafs["cSipStatsClientUnauthorizedOuts"] = csipstatserrclient.Csipstatsclientunauthorizedouts
    leafs["cSipStatsClientPaymentReqdIns"] = csipstatserrclient.Csipstatsclientpaymentreqdins
    leafs["cSipStatsClientPaymentReqdOuts"] = csipstatserrclient.Csipstatsclientpaymentreqdouts
    leafs["cSipStatsClientForbiddenIns"] = csipstatserrclient.Csipstatsclientforbiddenins
    leafs["cSipStatsClientForbiddenOuts"] = csipstatserrclient.Csipstatsclientforbiddenouts
    leafs["cSipStatsClientNotFoundIns"] = csipstatserrclient.Csipstatsclientnotfoundins
    leafs["cSipStatsClientNotFoundOuts"] = csipstatserrclient.Csipstatsclientnotfoundouts
    leafs["cSipStatsClientMethNotAllowedIns"] = csipstatserrclient.Csipstatsclientmethnotallowedins
    leafs["cSipStatsClientMethNotAllowedOuts"] = csipstatserrclient.Csipstatsclientmethnotallowedouts
    leafs["cSipStatsClientNotAcceptableIns"] = csipstatserrclient.Csipstatsclientnotacceptableins
    leafs["cSipStatsClientNotAcceptableOuts"] = csipstatserrclient.Csipstatsclientnotacceptableouts
    leafs["cSipStatsClientProxyAuthReqdIns"] = csipstatserrclient.Csipstatsclientproxyauthreqdins
    leafs["cSipStatsClientProxyAuthReqdOuts"] = csipstatserrclient.Csipstatsclientproxyauthreqdouts
    leafs["cSipStatsClientReqTimeoutIns"] = csipstatserrclient.Csipstatsclientreqtimeoutins
    leafs["cSipStatsClientReqTimeoutOuts"] = csipstatserrclient.Csipstatsclientreqtimeoutouts
    leafs["cSipStatsClientConflictIns"] = csipstatserrclient.Csipstatsclientconflictins
    leafs["cSipStatsClientConflictOuts"] = csipstatserrclient.Csipstatsclientconflictouts
    leafs["cSipStatsClientGoneIns"] = csipstatserrclient.Csipstatsclientgoneins
    leafs["cSipStatsClientGoneOuts"] = csipstatserrclient.Csipstatsclientgoneouts
    leafs["cSipStatsClientLengthRequiredIns"] = csipstatserrclient.Csipstatsclientlengthrequiredins
    leafs["cSipStatsClientLengthRequiredOuts"] = csipstatserrclient.Csipstatsclientlengthrequiredouts
    leafs["cSipStatsClientReqEntTooLargeIns"] = csipstatserrclient.Csipstatsclientreqenttoolargeins
    leafs["cSipStatsClientReqEntTooLargeOuts"] = csipstatserrclient.Csipstatsclientreqenttoolargeouts
    leafs["cSipStatsClientReqURITooLargeIns"] = csipstatserrclient.Csipstatsclientrequritoolargeins
    leafs["cSipStatsClientReqURITooLargeOuts"] = csipstatserrclient.Csipstatsclientrequritoolargeouts
    leafs["cSipStatsClientNoSupMediaTypeIns"] = csipstatserrclient.Csipstatsclientnosupmediatypeins
    leafs["cSipStatsClientNoSupMediaTypeOuts"] = csipstatserrclient.Csipstatsclientnosupmediatypeouts
    leafs["cSipStatsClientBadExtensionIns"] = csipstatserrclient.Csipstatsclientbadextensionins
    leafs["cSipStatsClientBadExtensionOuts"] = csipstatserrclient.Csipstatsclientbadextensionouts
    leafs["cSipStatsClientTempNotAvailIns"] = csipstatserrclient.Csipstatsclienttempnotavailins
    leafs["cSipStatsClientTempNotAvailOuts"] = csipstatserrclient.Csipstatsclienttempnotavailouts
    leafs["cSipStatsClientCallLegNoExistIns"] = csipstatserrclient.Csipstatsclientcalllegnoexistins
    leafs["cSipStatsClientCallLegNoExistOuts"] = csipstatserrclient.Csipstatsclientcalllegnoexistouts
    leafs["cSipStatsClientLoopDetectedIns"] = csipstatserrclient.Csipstatsclientloopdetectedins
    leafs["cSipStatsClientLoopDetectedOuts"] = csipstatserrclient.Csipstatsclientloopdetectedouts
    leafs["cSipStatsClientTooManyHopsIns"] = csipstatserrclient.Csipstatsclienttoomanyhopsins
    leafs["cSipStatsClientTooManyHopsOuts"] = csipstatserrclient.Csipstatsclienttoomanyhopsouts
    leafs["cSipStatsClientAddrIncompleteIns"] = csipstatserrclient.Csipstatsclientaddrincompleteins
    leafs["cSipStatsClientAddrIncompleteOuts"] = csipstatserrclient.Csipstatsclientaddrincompleteouts
    leafs["cSipStatsClientAmbiguousIns"] = csipstatserrclient.Csipstatsclientambiguousins
    leafs["cSipStatsClientAmbiguousOuts"] = csipstatserrclient.Csipstatsclientambiguousouts
    leafs["cSipStatsClientBusyHereIns"] = csipstatserrclient.Csipstatsclientbusyhereins
    leafs["cSipStatsClientBusyHereOuts"] = csipstatserrclient.Csipstatsclientbusyhereouts
    leafs["cSipStatsClientReqTermIns"] = csipstatserrclient.Csipstatsclientreqtermins
    leafs["cSipStatsClientReqTermOuts"] = csipstatserrclient.Csipstatsclientreqtermouts
    leafs["cSipStatsClientNoAcceptHereIns"] = csipstatserrclient.Csipstatsclientnoaccepthereins
    leafs["cSipStatsClientNoAcceptHereOuts"] = csipstatserrclient.Csipstatsclientnoaccepthereouts
    leafs["cSipStatsClientBadEventIns"] = csipstatserrclient.Csipstatsclientbadeventins
    leafs["cSipStatsClientBadEventOuts"] = csipstatserrclient.Csipstatsclientbadeventouts
    leafs["cSipStatsClientSTTooSmallIns"] = csipstatserrclient.Csipstatsclientsttoosmallins
    leafs["cSipStatsClientSTTooSmallOuts"] = csipstatserrclient.Csipstatsclientsttoosmallouts
    leafs["cSipStatsClientReqPendingIns"] = csipstatserrclient.Csipstatsclientreqpendingins
    leafs["cSipStatsClientReqPendingOuts"] = csipstatserrclient.Csipstatsclientreqpendingouts
    return leafs
}

func (csipstatserrclient *CISCOSIPUAMIB_Csipstatserrclient) GetBundleName() string { return "cisco_ios_xe" }

func (csipstatserrclient *CISCOSIPUAMIB_Csipstatserrclient) GetYangName() string { return "cSipStatsErrClient" }

func (csipstatserrclient *CISCOSIPUAMIB_Csipstatserrclient) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csipstatserrclient *CISCOSIPUAMIB_Csipstatserrclient) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csipstatserrclient *CISCOSIPUAMIB_Csipstatserrclient) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csipstatserrclient *CISCOSIPUAMIB_Csipstatserrclient) SetParent(parent types.Entity) { csipstatserrclient.parent = parent }

func (csipstatserrclient *CISCOSIPUAMIB_Csipstatserrclient) GetParent() types.Entity { return csipstatserrclient.parent }

func (csipstatserrclient *CISCOSIPUAMIB_Csipstatserrclient) GetParentYangName() string { return "CISCO-SIP-UA-MIB" }

// CISCOSIPUAMIB_Csipstatserrserver
type CISCOSIPUAMIB_Csipstatserrserver struct {
    parent types.Entity
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

func (csipstatserrserver *CISCOSIPUAMIB_Csipstatserrserver) GetFilter() yfilter.YFilter { return csipstatserrserver.YFilter }

func (csipstatserrserver *CISCOSIPUAMIB_Csipstatserrserver) SetFilter(yf yfilter.YFilter) { csipstatserrserver.YFilter = yf }

func (csipstatserrserver *CISCOSIPUAMIB_Csipstatserrserver) GetGoName(yname string) string {
    if yname == "cSipStatsServerIntErrorIns" { return "Csipstatsserverinterrorins" }
    if yname == "cSipStatsServerIntErrorOuts" { return "Csipstatsserverinterrorouts" }
    if yname == "cSipStatsServerNotImplementedIns" { return "Csipstatsservernotimplementedins" }
    if yname == "cSipStatsServerNotImplementedOuts" { return "Csipstatsservernotimplementedouts" }
    if yname == "cSipStatsServerBadGatewayIns" { return "Csipstatsserverbadgatewayins" }
    if yname == "cSipStatsServerBadGatewayOuts" { return "Csipstatsserverbadgatewayouts" }
    if yname == "cSipStatsServerServiceUnavailIns" { return "Csipstatsserverserviceunavailins" }
    if yname == "cSipStatsServerServiceUnavailOuts" { return "Csipstatsserverserviceunavailouts" }
    if yname == "cSipStatsServerGatewayTimeoutIns" { return "Csipstatsservergatewaytimeoutins" }
    if yname == "cSipStatsServerGatewayTimeoutOuts" { return "Csipstatsservergatewaytimeoutouts" }
    if yname == "cSipStatsServerBadSipVersionIns" { return "Csipstatsserverbadsipversionins" }
    if yname == "cSipStatsServerBadSipVersionOuts" { return "Csipstatsserverbadsipversionouts" }
    if yname == "cSipStatsServerPrecondFailureIns" { return "Csipstatsserverprecondfailureins" }
    if yname == "cSipStatsServerPrecondFailureOuts" { return "Csipstatsserverprecondfailureouts" }
    return ""
}

func (csipstatserrserver *CISCOSIPUAMIB_Csipstatserrserver) GetSegmentPath() string {
    return "cSipStatsErrServer"
}

func (csipstatserrserver *CISCOSIPUAMIB_Csipstatserrserver) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csipstatserrserver *CISCOSIPUAMIB_Csipstatserrserver) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csipstatserrserver *CISCOSIPUAMIB_Csipstatserrserver) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cSipStatsServerIntErrorIns"] = csipstatserrserver.Csipstatsserverinterrorins
    leafs["cSipStatsServerIntErrorOuts"] = csipstatserrserver.Csipstatsserverinterrorouts
    leafs["cSipStatsServerNotImplementedIns"] = csipstatserrserver.Csipstatsservernotimplementedins
    leafs["cSipStatsServerNotImplementedOuts"] = csipstatserrserver.Csipstatsservernotimplementedouts
    leafs["cSipStatsServerBadGatewayIns"] = csipstatserrserver.Csipstatsserverbadgatewayins
    leafs["cSipStatsServerBadGatewayOuts"] = csipstatserrserver.Csipstatsserverbadgatewayouts
    leafs["cSipStatsServerServiceUnavailIns"] = csipstatserrserver.Csipstatsserverserviceunavailins
    leafs["cSipStatsServerServiceUnavailOuts"] = csipstatserrserver.Csipstatsserverserviceunavailouts
    leafs["cSipStatsServerGatewayTimeoutIns"] = csipstatserrserver.Csipstatsservergatewaytimeoutins
    leafs["cSipStatsServerGatewayTimeoutOuts"] = csipstatserrserver.Csipstatsservergatewaytimeoutouts
    leafs["cSipStatsServerBadSipVersionIns"] = csipstatserrserver.Csipstatsserverbadsipversionins
    leafs["cSipStatsServerBadSipVersionOuts"] = csipstatserrserver.Csipstatsserverbadsipversionouts
    leafs["cSipStatsServerPrecondFailureIns"] = csipstatserrserver.Csipstatsserverprecondfailureins
    leafs["cSipStatsServerPrecondFailureOuts"] = csipstatserrserver.Csipstatsserverprecondfailureouts
    return leafs
}

func (csipstatserrserver *CISCOSIPUAMIB_Csipstatserrserver) GetBundleName() string { return "cisco_ios_xe" }

func (csipstatserrserver *CISCOSIPUAMIB_Csipstatserrserver) GetYangName() string { return "cSipStatsErrServer" }

func (csipstatserrserver *CISCOSIPUAMIB_Csipstatserrserver) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csipstatserrserver *CISCOSIPUAMIB_Csipstatserrserver) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csipstatserrserver *CISCOSIPUAMIB_Csipstatserrserver) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csipstatserrserver *CISCOSIPUAMIB_Csipstatserrserver) SetParent(parent types.Entity) { csipstatserrserver.parent = parent }

func (csipstatserrserver *CISCOSIPUAMIB_Csipstatserrserver) GetParent() types.Entity { return csipstatserrserver.parent }

func (csipstatserrserver *CISCOSIPUAMIB_Csipstatserrserver) GetParentYangName() string { return "CISCO-SIP-UA-MIB" }

// CISCOSIPUAMIB_Csipstatsglobalfail
type CISCOSIPUAMIB_Csipstatsglobalfail struct {
    parent types.Entity
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

func (csipstatsglobalfail *CISCOSIPUAMIB_Csipstatsglobalfail) GetFilter() yfilter.YFilter { return csipstatsglobalfail.YFilter }

func (csipstatsglobalfail *CISCOSIPUAMIB_Csipstatsglobalfail) SetFilter(yf yfilter.YFilter) { csipstatsglobalfail.YFilter = yf }

func (csipstatsglobalfail *CISCOSIPUAMIB_Csipstatsglobalfail) GetGoName(yname string) string {
    if yname == "cSipStatsGlobalBusyEverywhereIns" { return "Csipstatsglobalbusyeverywhereins" }
    if yname == "cSipStatsGlobalBusyEverywhereOuts" { return "Csipstatsglobalbusyeverywhereouts" }
    if yname == "cSipStatsGlobalDeclineIns" { return "Csipstatsglobaldeclineins" }
    if yname == "cSipStatsGlobalDeclineOuts" { return "Csipstatsglobaldeclineouts" }
    if yname == "cSipStatsGlobalNotAnywhereIns" { return "Csipstatsglobalnotanywhereins" }
    if yname == "cSipStatsGlobalNotAnywhereOuts" { return "Csipstatsglobalnotanywhereouts" }
    if yname == "cSipStatsGlobalNotAcceptableIns" { return "Csipstatsglobalnotacceptableins" }
    if yname == "cSipStatsGlobalNotAcceptableOuts" { return "Csipstatsglobalnotacceptableouts" }
    return ""
}

func (csipstatsglobalfail *CISCOSIPUAMIB_Csipstatsglobalfail) GetSegmentPath() string {
    return "cSipStatsGlobalFail"
}

func (csipstatsglobalfail *CISCOSIPUAMIB_Csipstatsglobalfail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csipstatsglobalfail *CISCOSIPUAMIB_Csipstatsglobalfail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csipstatsglobalfail *CISCOSIPUAMIB_Csipstatsglobalfail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cSipStatsGlobalBusyEverywhereIns"] = csipstatsglobalfail.Csipstatsglobalbusyeverywhereins
    leafs["cSipStatsGlobalBusyEverywhereOuts"] = csipstatsglobalfail.Csipstatsglobalbusyeverywhereouts
    leafs["cSipStatsGlobalDeclineIns"] = csipstatsglobalfail.Csipstatsglobaldeclineins
    leafs["cSipStatsGlobalDeclineOuts"] = csipstatsglobalfail.Csipstatsglobaldeclineouts
    leafs["cSipStatsGlobalNotAnywhereIns"] = csipstatsglobalfail.Csipstatsglobalnotanywhereins
    leafs["cSipStatsGlobalNotAnywhereOuts"] = csipstatsglobalfail.Csipstatsglobalnotanywhereouts
    leafs["cSipStatsGlobalNotAcceptableIns"] = csipstatsglobalfail.Csipstatsglobalnotacceptableins
    leafs["cSipStatsGlobalNotAcceptableOuts"] = csipstatsglobalfail.Csipstatsglobalnotacceptableouts
    return leafs
}

func (csipstatsglobalfail *CISCOSIPUAMIB_Csipstatsglobalfail) GetBundleName() string { return "cisco_ios_xe" }

func (csipstatsglobalfail *CISCOSIPUAMIB_Csipstatsglobalfail) GetYangName() string { return "cSipStatsGlobalFail" }

func (csipstatsglobalfail *CISCOSIPUAMIB_Csipstatsglobalfail) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csipstatsglobalfail *CISCOSIPUAMIB_Csipstatsglobalfail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csipstatsglobalfail *CISCOSIPUAMIB_Csipstatsglobalfail) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csipstatsglobalfail *CISCOSIPUAMIB_Csipstatsglobalfail) SetParent(parent types.Entity) { csipstatsglobalfail.parent = parent }

func (csipstatsglobalfail *CISCOSIPUAMIB_Csipstatsglobalfail) GetParent() types.Entity { return csipstatsglobalfail.parent }

func (csipstatsglobalfail *CISCOSIPUAMIB_Csipstatsglobalfail) GetParentYangName() string { return "CISCO-SIP-UA-MIB" }

// CISCOSIPUAMIB_Csipstatstraffic
type CISCOSIPUAMIB_Csipstatstraffic struct {
    parent types.Entity
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

func (csipstatstraffic *CISCOSIPUAMIB_Csipstatstraffic) GetFilter() yfilter.YFilter { return csipstatstraffic.YFilter }

func (csipstatstraffic *CISCOSIPUAMIB_Csipstatstraffic) SetFilter(yf yfilter.YFilter) { csipstatstraffic.YFilter = yf }

func (csipstatstraffic *CISCOSIPUAMIB_Csipstatstraffic) GetGoName(yname string) string {
    if yname == "cSipStatsTrafficInviteIns" { return "Csipstatstrafficinviteins" }
    if yname == "cSipStatsTrafficInviteOuts" { return "Csipstatstrafficinviteouts" }
    if yname == "cSipStatsTrafficAckIns" { return "Csipstatstrafficackins" }
    if yname == "cSipStatsTrafficAckOuts" { return "Csipstatstrafficackouts" }
    if yname == "cSipStatsTrafficByeIns" { return "Csipstatstrafficbyeins" }
    if yname == "cSipStatsTrafficByeOuts" { return "Csipstatstrafficbyeouts" }
    if yname == "cSipStatsTrafficCancelIns" { return "Csipstatstrafficcancelins" }
    if yname == "cSipStatsTrafficCancelOuts" { return "Csipstatstrafficcancelouts" }
    if yname == "cSipStatsTrafficOptionsIns" { return "Csipstatstrafficoptionsins" }
    if yname == "cSipStatsTrafficOptionsOuts" { return "Csipstatstrafficoptionsouts" }
    if yname == "cSipStatsTrafficRegisterIns" { return "Csipstatstrafficregisterins" }
    if yname == "cSipStatsTrafficRegisterOuts" { return "Csipstatstrafficregisterouts" }
    if yname == "cSipStatsTrafficCometIns" { return "Csipstatstrafficcometins" }
    if yname == "cSipStatsTrafficCometOuts" { return "Csipstatstrafficcometouts" }
    if yname == "cSipStatsTrafficPrackIns" { return "Csipstatstrafficprackins" }
    if yname == "cSipStatsTrafficPrackOuts" { return "Csipstatstrafficprackouts" }
    if yname == "cSipStatsTrafficReferIns" { return "Csipstatstrafficreferins" }
    if yname == "cSipStatsTrafficReferOuts" { return "Csipstatstrafficreferouts" }
    if yname == "cSipStatsTrafficNotifyIns" { return "Csipstatstrafficnotifyins" }
    if yname == "cSipStatsTrafficNotifyOuts" { return "Csipstatstrafficnotifyouts" }
    if yname == "cSipStatsTrafficInfoIns" { return "Csipstatstrafficinfoins" }
    if yname == "cSipStatsTrafficInfoOuts" { return "Csipstatstrafficinfoouts" }
    if yname == "cSipStatsTrafficSubscribeIns" { return "Csipstatstrafficsubscribeins" }
    if yname == "cSipStatsTrafficSubscribeOuts" { return "Csipstatstrafficsubscribeouts" }
    if yname == "cSipStatsTrafficUpdateIns" { return "Csipstatstrafficupdateins" }
    if yname == "cSipStatsTrafficUpdateOuts" { return "Csipstatstrafficupdateouts" }
    return ""
}

func (csipstatstraffic *CISCOSIPUAMIB_Csipstatstraffic) GetSegmentPath() string {
    return "cSipStatsTraffic"
}

func (csipstatstraffic *CISCOSIPUAMIB_Csipstatstraffic) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csipstatstraffic *CISCOSIPUAMIB_Csipstatstraffic) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csipstatstraffic *CISCOSIPUAMIB_Csipstatstraffic) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cSipStatsTrafficInviteIns"] = csipstatstraffic.Csipstatstrafficinviteins
    leafs["cSipStatsTrafficInviteOuts"] = csipstatstraffic.Csipstatstrafficinviteouts
    leafs["cSipStatsTrafficAckIns"] = csipstatstraffic.Csipstatstrafficackins
    leafs["cSipStatsTrafficAckOuts"] = csipstatstraffic.Csipstatstrafficackouts
    leafs["cSipStatsTrafficByeIns"] = csipstatstraffic.Csipstatstrafficbyeins
    leafs["cSipStatsTrafficByeOuts"] = csipstatstraffic.Csipstatstrafficbyeouts
    leafs["cSipStatsTrafficCancelIns"] = csipstatstraffic.Csipstatstrafficcancelins
    leafs["cSipStatsTrafficCancelOuts"] = csipstatstraffic.Csipstatstrafficcancelouts
    leafs["cSipStatsTrafficOptionsIns"] = csipstatstraffic.Csipstatstrafficoptionsins
    leafs["cSipStatsTrafficOptionsOuts"] = csipstatstraffic.Csipstatstrafficoptionsouts
    leafs["cSipStatsTrafficRegisterIns"] = csipstatstraffic.Csipstatstrafficregisterins
    leafs["cSipStatsTrafficRegisterOuts"] = csipstatstraffic.Csipstatstrafficregisterouts
    leafs["cSipStatsTrafficCometIns"] = csipstatstraffic.Csipstatstrafficcometins
    leafs["cSipStatsTrafficCometOuts"] = csipstatstraffic.Csipstatstrafficcometouts
    leafs["cSipStatsTrafficPrackIns"] = csipstatstraffic.Csipstatstrafficprackins
    leafs["cSipStatsTrafficPrackOuts"] = csipstatstraffic.Csipstatstrafficprackouts
    leafs["cSipStatsTrafficReferIns"] = csipstatstraffic.Csipstatstrafficreferins
    leafs["cSipStatsTrafficReferOuts"] = csipstatstraffic.Csipstatstrafficreferouts
    leafs["cSipStatsTrafficNotifyIns"] = csipstatstraffic.Csipstatstrafficnotifyins
    leafs["cSipStatsTrafficNotifyOuts"] = csipstatstraffic.Csipstatstrafficnotifyouts
    leafs["cSipStatsTrafficInfoIns"] = csipstatstraffic.Csipstatstrafficinfoins
    leafs["cSipStatsTrafficInfoOuts"] = csipstatstraffic.Csipstatstrafficinfoouts
    leafs["cSipStatsTrafficSubscribeIns"] = csipstatstraffic.Csipstatstrafficsubscribeins
    leafs["cSipStatsTrafficSubscribeOuts"] = csipstatstraffic.Csipstatstrafficsubscribeouts
    leafs["cSipStatsTrafficUpdateIns"] = csipstatstraffic.Csipstatstrafficupdateins
    leafs["cSipStatsTrafficUpdateOuts"] = csipstatstraffic.Csipstatstrafficupdateouts
    return leafs
}

func (csipstatstraffic *CISCOSIPUAMIB_Csipstatstraffic) GetBundleName() string { return "cisco_ios_xe" }

func (csipstatstraffic *CISCOSIPUAMIB_Csipstatstraffic) GetYangName() string { return "cSipStatsTraffic" }

func (csipstatstraffic *CISCOSIPUAMIB_Csipstatstraffic) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csipstatstraffic *CISCOSIPUAMIB_Csipstatstraffic) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csipstatstraffic *CISCOSIPUAMIB_Csipstatstraffic) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csipstatstraffic *CISCOSIPUAMIB_Csipstatstraffic) SetParent(parent types.Entity) { csipstatstraffic.parent = parent }

func (csipstatstraffic *CISCOSIPUAMIB_Csipstatstraffic) GetParent() types.Entity { return csipstatstraffic.parent }

func (csipstatstraffic *CISCOSIPUAMIB_Csipstatstraffic) GetParentYangName() string { return "CISCO-SIP-UA-MIB" }

// CISCOSIPUAMIB_Csipstatsretry
type CISCOSIPUAMIB_Csipstatsretry struct {
    parent types.Entity
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

func (csipstatsretry *CISCOSIPUAMIB_Csipstatsretry) GetFilter() yfilter.YFilter { return csipstatsretry.YFilter }

func (csipstatsretry *CISCOSIPUAMIB_Csipstatsretry) SetFilter(yf yfilter.YFilter) { csipstatsretry.YFilter = yf }

func (csipstatsretry *CISCOSIPUAMIB_Csipstatsretry) GetGoName(yname string) string {
    if yname == "cSipStatsRetryInvites" { return "Csipstatsretryinvites" }
    if yname == "cSipStatsRetryByes" { return "Csipstatsretrybyes" }
    if yname == "cSipStatsRetryCancels" { return "Csipstatsretrycancels" }
    if yname == "cSipStatsRetryRegisters" { return "Csipstatsretryregisters" }
    if yname == "cSipStatsRetryResponses" { return "Csipstatsretryresponses" }
    if yname == "cSipStatsRetryPracks" { return "Csipstatsretrypracks" }
    if yname == "cSipStatsRetryComets" { return "Csipstatsretrycomets" }
    if yname == "cSipStatsRetryReliable1xxRsps" { return "Csipstatsretryreliable1Xxrsps" }
    if yname == "cSipStatsRetryNotifys" { return "Csipstatsretrynotifys" }
    if yname == "cSipStatsRetryRefers" { return "Csipstatsretryrefers" }
    if yname == "cSipStatsRetryInfo" { return "Csipstatsretryinfo" }
    if yname == "cSipStatsRetrySubscribe" { return "Csipstatsretrysubscribe" }
    return ""
}

func (csipstatsretry *CISCOSIPUAMIB_Csipstatsretry) GetSegmentPath() string {
    return "cSipStatsRetry"
}

func (csipstatsretry *CISCOSIPUAMIB_Csipstatsretry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csipstatsretry *CISCOSIPUAMIB_Csipstatsretry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csipstatsretry *CISCOSIPUAMIB_Csipstatsretry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cSipStatsRetryInvites"] = csipstatsretry.Csipstatsretryinvites
    leafs["cSipStatsRetryByes"] = csipstatsretry.Csipstatsretrybyes
    leafs["cSipStatsRetryCancels"] = csipstatsretry.Csipstatsretrycancels
    leafs["cSipStatsRetryRegisters"] = csipstatsretry.Csipstatsretryregisters
    leafs["cSipStatsRetryResponses"] = csipstatsretry.Csipstatsretryresponses
    leafs["cSipStatsRetryPracks"] = csipstatsretry.Csipstatsretrypracks
    leafs["cSipStatsRetryComets"] = csipstatsretry.Csipstatsretrycomets
    leafs["cSipStatsRetryReliable1xxRsps"] = csipstatsretry.Csipstatsretryreliable1Xxrsps
    leafs["cSipStatsRetryNotifys"] = csipstatsretry.Csipstatsretrynotifys
    leafs["cSipStatsRetryRefers"] = csipstatsretry.Csipstatsretryrefers
    leafs["cSipStatsRetryInfo"] = csipstatsretry.Csipstatsretryinfo
    leafs["cSipStatsRetrySubscribe"] = csipstatsretry.Csipstatsretrysubscribe
    return leafs
}

func (csipstatsretry *CISCOSIPUAMIB_Csipstatsretry) GetBundleName() string { return "cisco_ios_xe" }

func (csipstatsretry *CISCOSIPUAMIB_Csipstatsretry) GetYangName() string { return "cSipStatsRetry" }

func (csipstatsretry *CISCOSIPUAMIB_Csipstatsretry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csipstatsretry *CISCOSIPUAMIB_Csipstatsretry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csipstatsretry *CISCOSIPUAMIB_Csipstatsretry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csipstatsretry *CISCOSIPUAMIB_Csipstatsretry) SetParent(parent types.Entity) { csipstatsretry.parent = parent }

func (csipstatsretry *CISCOSIPUAMIB_Csipstatsretry) GetParent() types.Entity { return csipstatsretry.parent }

func (csipstatsretry *CISCOSIPUAMIB_Csipstatsretry) GetParentYangName() string { return "CISCO-SIP-UA-MIB" }

// CISCOSIPUAMIB_Csipstatsmisc
type CISCOSIPUAMIB_Csipstatsmisc struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This object reflects the total number of incoming Redirect  (3xx) response
    // messages that have been mapped to Client  Error (4xx) response messages.
    // The type is interface{} with range: 0..4294967295.
    Csipstatsmisc3Xxmappedto4Xxrsps interface{}
}

func (csipstatsmisc *CISCOSIPUAMIB_Csipstatsmisc) GetFilter() yfilter.YFilter { return csipstatsmisc.YFilter }

func (csipstatsmisc *CISCOSIPUAMIB_Csipstatsmisc) SetFilter(yf yfilter.YFilter) { csipstatsmisc.YFilter = yf }

func (csipstatsmisc *CISCOSIPUAMIB_Csipstatsmisc) GetGoName(yname string) string {
    if yname == "cSipStatsMisc3xxMappedTo4xxRsps" { return "Csipstatsmisc3Xxmappedto4Xxrsps" }
    return ""
}

func (csipstatsmisc *CISCOSIPUAMIB_Csipstatsmisc) GetSegmentPath() string {
    return "cSipStatsMisc"
}

func (csipstatsmisc *CISCOSIPUAMIB_Csipstatsmisc) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csipstatsmisc *CISCOSIPUAMIB_Csipstatsmisc) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csipstatsmisc *CISCOSIPUAMIB_Csipstatsmisc) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cSipStatsMisc3xxMappedTo4xxRsps"] = csipstatsmisc.Csipstatsmisc3Xxmappedto4Xxrsps
    return leafs
}

func (csipstatsmisc *CISCOSIPUAMIB_Csipstatsmisc) GetBundleName() string { return "cisco_ios_xe" }

func (csipstatsmisc *CISCOSIPUAMIB_Csipstatsmisc) GetYangName() string { return "cSipStatsMisc" }

func (csipstatsmisc *CISCOSIPUAMIB_Csipstatsmisc) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csipstatsmisc *CISCOSIPUAMIB_Csipstatsmisc) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csipstatsmisc *CISCOSIPUAMIB_Csipstatsmisc) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csipstatsmisc *CISCOSIPUAMIB_Csipstatsmisc) SetParent(parent types.Entity) { csipstatsmisc.parent = parent }

func (csipstatsmisc *CISCOSIPUAMIB_Csipstatsmisc) GetParent() types.Entity { return csipstatsmisc.parent }

func (csipstatsmisc *CISCOSIPUAMIB_Csipstatsmisc) GetParentYangName() string { return "CISCO-SIP-UA-MIB" }

// CISCOSIPUAMIB_Csipstatsconnection
type CISCOSIPUAMIB_Csipstatsconnection struct {
    parent types.Entity
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

func (csipstatsconnection *CISCOSIPUAMIB_Csipstatsconnection) GetFilter() yfilter.YFilter { return csipstatsconnection.YFilter }

func (csipstatsconnection *CISCOSIPUAMIB_Csipstatsconnection) SetFilter(yf yfilter.YFilter) { csipstatsconnection.YFilter = yf }

func (csipstatsconnection *CISCOSIPUAMIB_Csipstatsconnection) GetGoName(yname string) string {
    if yname == "cSipStatsConnTCPSendFailures" { return "Csipstatsconntcpsendfailures" }
    if yname == "cSipStatsConnUDPSendFailures" { return "Csipstatsconnudpsendfailures" }
    if yname == "cSipStatsConnTCPRemoteClosures" { return "Csipstatsconntcpremoteclosures" }
    if yname == "cSipStatsConnUDPCreateFailures" { return "Csipstatsconnudpcreatefailures" }
    if yname == "cSipStatsConnTCPCreateFailures" { return "Csipstatsconntcpcreatefailures" }
    if yname == "cSipStatsConnUDPInactiveTimeouts" { return "Csipstatsconnudpinactivetimeouts" }
    if yname == "cSipStatsConnTCPInactiveTimeouts" { return "Csipstatsconntcpinactivetimeouts" }
    if yname == "cSipStatsActiveUDPConnections" { return "Csipstatsactiveudpconnections" }
    if yname == "cSipStatsActiveTCPConnections" { return "Csipstatsactivetcpconnections" }
    return ""
}

func (csipstatsconnection *CISCOSIPUAMIB_Csipstatsconnection) GetSegmentPath() string {
    return "cSipStatsConnection"
}

func (csipstatsconnection *CISCOSIPUAMIB_Csipstatsconnection) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csipstatsconnection *CISCOSIPUAMIB_Csipstatsconnection) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csipstatsconnection *CISCOSIPUAMIB_Csipstatsconnection) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cSipStatsConnTCPSendFailures"] = csipstatsconnection.Csipstatsconntcpsendfailures
    leafs["cSipStatsConnUDPSendFailures"] = csipstatsconnection.Csipstatsconnudpsendfailures
    leafs["cSipStatsConnTCPRemoteClosures"] = csipstatsconnection.Csipstatsconntcpremoteclosures
    leafs["cSipStatsConnUDPCreateFailures"] = csipstatsconnection.Csipstatsconnudpcreatefailures
    leafs["cSipStatsConnTCPCreateFailures"] = csipstatsconnection.Csipstatsconntcpcreatefailures
    leafs["cSipStatsConnUDPInactiveTimeouts"] = csipstatsconnection.Csipstatsconnudpinactivetimeouts
    leafs["cSipStatsConnTCPInactiveTimeouts"] = csipstatsconnection.Csipstatsconntcpinactivetimeouts
    leafs["cSipStatsActiveUDPConnections"] = csipstatsconnection.Csipstatsactiveudpconnections
    leafs["cSipStatsActiveTCPConnections"] = csipstatsconnection.Csipstatsactivetcpconnections
    return leafs
}

func (csipstatsconnection *CISCOSIPUAMIB_Csipstatsconnection) GetBundleName() string { return "cisco_ios_xe" }

func (csipstatsconnection *CISCOSIPUAMIB_Csipstatsconnection) GetYangName() string { return "cSipStatsConnection" }

func (csipstatsconnection *CISCOSIPUAMIB_Csipstatsconnection) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csipstatsconnection *CISCOSIPUAMIB_Csipstatsconnection) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csipstatsconnection *CISCOSIPUAMIB_Csipstatsconnection) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csipstatsconnection *CISCOSIPUAMIB_Csipstatsconnection) SetParent(parent types.Entity) { csipstatsconnection.parent = parent }

func (csipstatsconnection *CISCOSIPUAMIB_Csipstatsconnection) GetParent() types.Entity { return csipstatsconnection.parent }

func (csipstatsconnection *CISCOSIPUAMIB_Csipstatsconnection) GetParentYangName() string { return "CISCO-SIP-UA-MIB" }

// CISCOSIPUAMIB_Csipcfgearlymediatable
// This table contains configuration for Early
// Media Cut Through.  The configuration controls
// how the SIP user agent will process 1xx
// (Provisional) SIP response messages that contain 
// Session Definition Protocol (SDP) payloads.
type CISCOSIPUAMIB_Csipcfgearlymediatable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A row in the cSipCfgEarlyMediaTable. A row is accessible with a Provisional
    // (1xx) status code value (eg, 180) and provides an object for the enabling
    // or disabling of the Early Media Cut Through functionality. The type is
    // slice of CISCOSIPUAMIB_Csipcfgearlymediatable_Csipcfgearlymediaentry.
    Csipcfgearlymediaentry []CISCOSIPUAMIB_Csipcfgearlymediatable_Csipcfgearlymediaentry
}

func (csipcfgearlymediatable *CISCOSIPUAMIB_Csipcfgearlymediatable) GetFilter() yfilter.YFilter { return csipcfgearlymediatable.YFilter }

func (csipcfgearlymediatable *CISCOSIPUAMIB_Csipcfgearlymediatable) SetFilter(yf yfilter.YFilter) { csipcfgearlymediatable.YFilter = yf }

func (csipcfgearlymediatable *CISCOSIPUAMIB_Csipcfgearlymediatable) GetGoName(yname string) string {
    if yname == "cSipCfgEarlyMediaEntry" { return "Csipcfgearlymediaentry" }
    return ""
}

func (csipcfgearlymediatable *CISCOSIPUAMIB_Csipcfgearlymediatable) GetSegmentPath() string {
    return "cSipCfgEarlyMediaTable"
}

func (csipcfgearlymediatable *CISCOSIPUAMIB_Csipcfgearlymediatable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cSipCfgEarlyMediaEntry" {
        for _, c := range csipcfgearlymediatable.Csipcfgearlymediaentry {
            if csipcfgearlymediatable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOSIPUAMIB_Csipcfgearlymediatable_Csipcfgearlymediaentry{}
        csipcfgearlymediatable.Csipcfgearlymediaentry = append(csipcfgearlymediatable.Csipcfgearlymediaentry, child)
        return &csipcfgearlymediatable.Csipcfgearlymediaentry[len(csipcfgearlymediatable.Csipcfgearlymediaentry)-1]
    }
    return nil
}

func (csipcfgearlymediatable *CISCOSIPUAMIB_Csipcfgearlymediatable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range csipcfgearlymediatable.Csipcfgearlymediaentry {
        children[csipcfgearlymediatable.Csipcfgearlymediaentry[i].GetSegmentPath()] = &csipcfgearlymediatable.Csipcfgearlymediaentry[i]
    }
    return children
}

func (csipcfgearlymediatable *CISCOSIPUAMIB_Csipcfgearlymediatable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (csipcfgearlymediatable *CISCOSIPUAMIB_Csipcfgearlymediatable) GetBundleName() string { return "cisco_ios_xe" }

func (csipcfgearlymediatable *CISCOSIPUAMIB_Csipcfgearlymediatable) GetYangName() string { return "cSipCfgEarlyMediaTable" }

func (csipcfgearlymediatable *CISCOSIPUAMIB_Csipcfgearlymediatable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csipcfgearlymediatable *CISCOSIPUAMIB_Csipcfgearlymediatable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csipcfgearlymediatable *CISCOSIPUAMIB_Csipcfgearlymediatable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csipcfgearlymediatable *CISCOSIPUAMIB_Csipcfgearlymediatable) SetParent(parent types.Entity) { csipcfgearlymediatable.parent = parent }

func (csipcfgearlymediatable *CISCOSIPUAMIB_Csipcfgearlymediatable) GetParent() types.Entity { return csipcfgearlymediatable.parent }

func (csipcfgearlymediatable *CISCOSIPUAMIB_Csipcfgearlymediatable) GetParentYangName() string { return "CISCO-SIP-UA-MIB" }

// CISCOSIPUAMIB_Csipcfgearlymediatable_Csipcfgearlymediaentry
// A row in the cSipCfgEarlyMediaTable.
// A row is accessible with a Provisional (1xx)
// status code value (eg, 180) and provides
// an object for the enabling or disabling of
// the Early Media Cut Through functionality.
type CISCOSIPUAMIB_Csipcfgearlymediatable_Csipcfgearlymediaentry struct {
    parent types.Entity
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

func (csipcfgearlymediaentry *CISCOSIPUAMIB_Csipcfgearlymediatable_Csipcfgearlymediaentry) GetFilter() yfilter.YFilter { return csipcfgearlymediaentry.YFilter }

func (csipcfgearlymediaentry *CISCOSIPUAMIB_Csipcfgearlymediatable_Csipcfgearlymediaentry) SetFilter(yf yfilter.YFilter) { csipcfgearlymediaentry.YFilter = yf }

func (csipcfgearlymediaentry *CISCOSIPUAMIB_Csipcfgearlymediatable_Csipcfgearlymediaentry) GetGoName(yname string) string {
    if yname == "cSipCfgEarlyMediaStatusCodeIndex" { return "Csipcfgearlymediastatuscodeindex" }
    if yname == "cSipCfgEarlyMediaCutThruDisabled" { return "Csipcfgearlymediacutthrudisabled" }
    return ""
}

func (csipcfgearlymediaentry *CISCOSIPUAMIB_Csipcfgearlymediatable_Csipcfgearlymediaentry) GetSegmentPath() string {
    return "cSipCfgEarlyMediaEntry" + "[cSipCfgEarlyMediaStatusCodeIndex='" + fmt.Sprintf("%v", csipcfgearlymediaentry.Csipcfgearlymediastatuscodeindex) + "']"
}

func (csipcfgearlymediaentry *CISCOSIPUAMIB_Csipcfgearlymediatable_Csipcfgearlymediaentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csipcfgearlymediaentry *CISCOSIPUAMIB_Csipcfgearlymediatable_Csipcfgearlymediaentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csipcfgearlymediaentry *CISCOSIPUAMIB_Csipcfgearlymediatable_Csipcfgearlymediaentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cSipCfgEarlyMediaStatusCodeIndex"] = csipcfgearlymediaentry.Csipcfgearlymediastatuscodeindex
    leafs["cSipCfgEarlyMediaCutThruDisabled"] = csipcfgearlymediaentry.Csipcfgearlymediacutthrudisabled
    return leafs
}

func (csipcfgearlymediaentry *CISCOSIPUAMIB_Csipcfgearlymediatable_Csipcfgearlymediaentry) GetBundleName() string { return "cisco_ios_xe" }

func (csipcfgearlymediaentry *CISCOSIPUAMIB_Csipcfgearlymediatable_Csipcfgearlymediaentry) GetYangName() string { return "cSipCfgEarlyMediaEntry" }

func (csipcfgearlymediaentry *CISCOSIPUAMIB_Csipcfgearlymediatable_Csipcfgearlymediaentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csipcfgearlymediaentry *CISCOSIPUAMIB_Csipcfgearlymediatable_Csipcfgearlymediaentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csipcfgearlymediaentry *CISCOSIPUAMIB_Csipcfgearlymediatable_Csipcfgearlymediaentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csipcfgearlymediaentry *CISCOSIPUAMIB_Csipcfgearlymediatable_Csipcfgearlymediaentry) SetParent(parent types.Entity) { csipcfgearlymediaentry.parent = parent }

func (csipcfgearlymediaentry *CISCOSIPUAMIB_Csipcfgearlymediatable_Csipcfgearlymediaentry) GetParent() types.Entity { return csipcfgearlymediaentry.parent }

func (csipcfgearlymediaentry *CISCOSIPUAMIB_Csipcfgearlymediatable_Csipcfgearlymediaentry) GetParentYangName() string { return "cSipCfgEarlyMediaTable" }

// CISCOSIPUAMIB_Csipcfgbindsourceaddrtable
// This table contains configuration for binding
// the scope of packets to the particular ethernet
// interface. The scope for the packets can be
// specified as either 'signalling' or 'media'
// packets. The ethernet interface shall be
// specified by the interface index. The table
// shall be indexed based on the scope.
type CISCOSIPUAMIB_Csipcfgbindsourceaddrtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A row in the cSipCfgBindSourceAddrTable. A row is accessible with the scope
    // of packets to which the source IP address of the interface designated by
    // cSipCfgBindSourceAddrInterface will be bound. The type is slice of
    // CISCOSIPUAMIB_Csipcfgbindsourceaddrtable_Csipcfgbindsourceaddrentry.
    Csipcfgbindsourceaddrentry []CISCOSIPUAMIB_Csipcfgbindsourceaddrtable_Csipcfgbindsourceaddrentry
}

func (csipcfgbindsourceaddrtable *CISCOSIPUAMIB_Csipcfgbindsourceaddrtable) GetFilter() yfilter.YFilter { return csipcfgbindsourceaddrtable.YFilter }

func (csipcfgbindsourceaddrtable *CISCOSIPUAMIB_Csipcfgbindsourceaddrtable) SetFilter(yf yfilter.YFilter) { csipcfgbindsourceaddrtable.YFilter = yf }

func (csipcfgbindsourceaddrtable *CISCOSIPUAMIB_Csipcfgbindsourceaddrtable) GetGoName(yname string) string {
    if yname == "cSipCfgBindSourceAddrEntry" { return "Csipcfgbindsourceaddrentry" }
    return ""
}

func (csipcfgbindsourceaddrtable *CISCOSIPUAMIB_Csipcfgbindsourceaddrtable) GetSegmentPath() string {
    return "cSipCfgBindSourceAddrTable"
}

func (csipcfgbindsourceaddrtable *CISCOSIPUAMIB_Csipcfgbindsourceaddrtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cSipCfgBindSourceAddrEntry" {
        for _, c := range csipcfgbindsourceaddrtable.Csipcfgbindsourceaddrentry {
            if csipcfgbindsourceaddrtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOSIPUAMIB_Csipcfgbindsourceaddrtable_Csipcfgbindsourceaddrentry{}
        csipcfgbindsourceaddrtable.Csipcfgbindsourceaddrentry = append(csipcfgbindsourceaddrtable.Csipcfgbindsourceaddrentry, child)
        return &csipcfgbindsourceaddrtable.Csipcfgbindsourceaddrentry[len(csipcfgbindsourceaddrtable.Csipcfgbindsourceaddrentry)-1]
    }
    return nil
}

func (csipcfgbindsourceaddrtable *CISCOSIPUAMIB_Csipcfgbindsourceaddrtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range csipcfgbindsourceaddrtable.Csipcfgbindsourceaddrentry {
        children[csipcfgbindsourceaddrtable.Csipcfgbindsourceaddrentry[i].GetSegmentPath()] = &csipcfgbindsourceaddrtable.Csipcfgbindsourceaddrentry[i]
    }
    return children
}

func (csipcfgbindsourceaddrtable *CISCOSIPUAMIB_Csipcfgbindsourceaddrtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (csipcfgbindsourceaddrtable *CISCOSIPUAMIB_Csipcfgbindsourceaddrtable) GetBundleName() string { return "cisco_ios_xe" }

func (csipcfgbindsourceaddrtable *CISCOSIPUAMIB_Csipcfgbindsourceaddrtable) GetYangName() string { return "cSipCfgBindSourceAddrTable" }

func (csipcfgbindsourceaddrtable *CISCOSIPUAMIB_Csipcfgbindsourceaddrtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csipcfgbindsourceaddrtable *CISCOSIPUAMIB_Csipcfgbindsourceaddrtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csipcfgbindsourceaddrtable *CISCOSIPUAMIB_Csipcfgbindsourceaddrtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csipcfgbindsourceaddrtable *CISCOSIPUAMIB_Csipcfgbindsourceaddrtable) SetParent(parent types.Entity) { csipcfgbindsourceaddrtable.parent = parent }

func (csipcfgbindsourceaddrtable *CISCOSIPUAMIB_Csipcfgbindsourceaddrtable) GetParent() types.Entity { return csipcfgbindsourceaddrtable.parent }

func (csipcfgbindsourceaddrtable *CISCOSIPUAMIB_Csipcfgbindsourceaddrtable) GetParentYangName() string { return "CISCO-SIP-UA-MIB" }

// CISCOSIPUAMIB_Csipcfgbindsourceaddrtable_Csipcfgbindsourceaddrentry
// A row in the cSipCfgBindSourceAddrTable.
// A row is accessible with the scope of packets
// to which the source IP address of the interface
// designated by cSipCfgBindSourceAddrInterface
// will be bound.
type CISCOSIPUAMIB_Csipcfgbindsourceaddrtable_Csipcfgbindsourceaddrentry struct {
    parent types.Entity
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

func (csipcfgbindsourceaddrentry *CISCOSIPUAMIB_Csipcfgbindsourceaddrtable_Csipcfgbindsourceaddrentry) GetFilter() yfilter.YFilter { return csipcfgbindsourceaddrentry.YFilter }

func (csipcfgbindsourceaddrentry *CISCOSIPUAMIB_Csipcfgbindsourceaddrtable_Csipcfgbindsourceaddrentry) SetFilter(yf yfilter.YFilter) { csipcfgbindsourceaddrentry.YFilter = yf }

func (csipcfgbindsourceaddrentry *CISCOSIPUAMIB_Csipcfgbindsourceaddrtable_Csipcfgbindsourceaddrentry) GetGoName(yname string) string {
    if yname == "cSipCfgBindSourceAddrScope" { return "Csipcfgbindsourceaddrscope" }
    if yname == "cSipCfgBindSourceAddrInterface" { return "Csipcfgbindsourceaddrinterface" }
    return ""
}

func (csipcfgbindsourceaddrentry *CISCOSIPUAMIB_Csipcfgbindsourceaddrtable_Csipcfgbindsourceaddrentry) GetSegmentPath() string {
    return "cSipCfgBindSourceAddrEntry" + "[cSipCfgBindSourceAddrScope='" + fmt.Sprintf("%v", csipcfgbindsourceaddrentry.Csipcfgbindsourceaddrscope) + "']"
}

func (csipcfgbindsourceaddrentry *CISCOSIPUAMIB_Csipcfgbindsourceaddrtable_Csipcfgbindsourceaddrentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csipcfgbindsourceaddrentry *CISCOSIPUAMIB_Csipcfgbindsourceaddrtable_Csipcfgbindsourceaddrentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csipcfgbindsourceaddrentry *CISCOSIPUAMIB_Csipcfgbindsourceaddrtable_Csipcfgbindsourceaddrentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cSipCfgBindSourceAddrScope"] = csipcfgbindsourceaddrentry.Csipcfgbindsourceaddrscope
    leafs["cSipCfgBindSourceAddrInterface"] = csipcfgbindsourceaddrentry.Csipcfgbindsourceaddrinterface
    return leafs
}

func (csipcfgbindsourceaddrentry *CISCOSIPUAMIB_Csipcfgbindsourceaddrtable_Csipcfgbindsourceaddrentry) GetBundleName() string { return "cisco_ios_xe" }

func (csipcfgbindsourceaddrentry *CISCOSIPUAMIB_Csipcfgbindsourceaddrtable_Csipcfgbindsourceaddrentry) GetYangName() string { return "cSipCfgBindSourceAddrEntry" }

func (csipcfgbindsourceaddrentry *CISCOSIPUAMIB_Csipcfgbindsourceaddrtable_Csipcfgbindsourceaddrentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csipcfgbindsourceaddrentry *CISCOSIPUAMIB_Csipcfgbindsourceaddrtable_Csipcfgbindsourceaddrentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csipcfgbindsourceaddrentry *CISCOSIPUAMIB_Csipcfgbindsourceaddrtable_Csipcfgbindsourceaddrentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csipcfgbindsourceaddrentry *CISCOSIPUAMIB_Csipcfgbindsourceaddrtable_Csipcfgbindsourceaddrentry) SetParent(parent types.Entity) { csipcfgbindsourceaddrentry.parent = parent }

func (csipcfgbindsourceaddrentry *CISCOSIPUAMIB_Csipcfgbindsourceaddrtable_Csipcfgbindsourceaddrentry) GetParent() types.Entity { return csipcfgbindsourceaddrentry.parent }

func (csipcfgbindsourceaddrentry *CISCOSIPUAMIB_Csipcfgbindsourceaddrtable_Csipcfgbindsourceaddrentry) GetParentYangName() string { return "cSipCfgBindSourceAddrTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // A row in the cSipCfgPeerTable. The type is slice of
    // CISCOSIPUAMIB_Csipcfgpeertable_Csipcfgpeerentry.
    Csipcfgpeerentry []CISCOSIPUAMIB_Csipcfgpeertable_Csipcfgpeerentry
}

func (csipcfgpeertable *CISCOSIPUAMIB_Csipcfgpeertable) GetFilter() yfilter.YFilter { return csipcfgpeertable.YFilter }

func (csipcfgpeertable *CISCOSIPUAMIB_Csipcfgpeertable) SetFilter(yf yfilter.YFilter) { csipcfgpeertable.YFilter = yf }

func (csipcfgpeertable *CISCOSIPUAMIB_Csipcfgpeertable) GetGoName(yname string) string {
    if yname == "cSipCfgPeerEntry" { return "Csipcfgpeerentry" }
    return ""
}

func (csipcfgpeertable *CISCOSIPUAMIB_Csipcfgpeertable) GetSegmentPath() string {
    return "cSipCfgPeerTable"
}

func (csipcfgpeertable *CISCOSIPUAMIB_Csipcfgpeertable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cSipCfgPeerEntry" {
        for _, c := range csipcfgpeertable.Csipcfgpeerentry {
            if csipcfgpeertable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOSIPUAMIB_Csipcfgpeertable_Csipcfgpeerentry{}
        csipcfgpeertable.Csipcfgpeerentry = append(csipcfgpeertable.Csipcfgpeerentry, child)
        return &csipcfgpeertable.Csipcfgpeerentry[len(csipcfgpeertable.Csipcfgpeerentry)-1]
    }
    return nil
}

func (csipcfgpeertable *CISCOSIPUAMIB_Csipcfgpeertable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range csipcfgpeertable.Csipcfgpeerentry {
        children[csipcfgpeertable.Csipcfgpeerentry[i].GetSegmentPath()] = &csipcfgpeertable.Csipcfgpeerentry[i]
    }
    return children
}

func (csipcfgpeertable *CISCOSIPUAMIB_Csipcfgpeertable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (csipcfgpeertable *CISCOSIPUAMIB_Csipcfgpeertable) GetBundleName() string { return "cisco_ios_xe" }

func (csipcfgpeertable *CISCOSIPUAMIB_Csipcfgpeertable) GetYangName() string { return "cSipCfgPeerTable" }

func (csipcfgpeertable *CISCOSIPUAMIB_Csipcfgpeertable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csipcfgpeertable *CISCOSIPUAMIB_Csipcfgpeertable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csipcfgpeertable *CISCOSIPUAMIB_Csipcfgpeertable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csipcfgpeertable *CISCOSIPUAMIB_Csipcfgpeertable) SetParent(parent types.Entity) { csipcfgpeertable.parent = parent }

func (csipcfgpeertable *CISCOSIPUAMIB_Csipcfgpeertable) GetParent() types.Entity { return csipcfgpeertable.parent }

func (csipcfgpeertable *CISCOSIPUAMIB_Csipcfgpeertable) GetParentYangName() string { return "CISCO-SIP-UA-MIB" }

// CISCOSIPUAMIB_Csipcfgpeertable_Csipcfgpeerentry
// A row in the cSipCfgPeerTable.
type CISCOSIPUAMIB_Csipcfgpeertable_Csipcfgpeerentry struct {
    parent types.Entity
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

func (csipcfgpeerentry *CISCOSIPUAMIB_Csipcfgpeertable_Csipcfgpeerentry) GetFilter() yfilter.YFilter { return csipcfgpeerentry.YFilter }

func (csipcfgpeerentry *CISCOSIPUAMIB_Csipcfgpeertable_Csipcfgpeerentry) SetFilter(yf yfilter.YFilter) { csipcfgpeerentry.YFilter = yf }

func (csipcfgpeerentry *CISCOSIPUAMIB_Csipcfgpeertable_Csipcfgpeerentry) GetGoName(yname string) string {
    if yname == "cSipCfgPeerIndex" { return "Csipcfgpeerindex" }
    if yname == "cSipCfgPeerOutSessionTransport" { return "Csipcfgpeeroutsessiontransport" }
    if yname == "cSipCfgPeerReliable1xxRspStr" { return "Csipcfgpeerreliable1Xxrspstr" }
    if yname == "cSipCfgPeerReliable1xxRspHdr" { return "Csipcfgpeerreliable1Xxrsphdr" }
    if yname == "cSipCfgPeerUrlType" { return "Csipcfgpeerurltype" }
    if yname == "cSipCfgPeerSwitchTransEnabled" { return "Csipcfgpeerswitchtransenabled" }
    return ""
}

func (csipcfgpeerentry *CISCOSIPUAMIB_Csipcfgpeertable_Csipcfgpeerentry) GetSegmentPath() string {
    return "cSipCfgPeerEntry" + "[cSipCfgPeerIndex='" + fmt.Sprintf("%v", csipcfgpeerentry.Csipcfgpeerindex) + "']"
}

func (csipcfgpeerentry *CISCOSIPUAMIB_Csipcfgpeertable_Csipcfgpeerentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csipcfgpeerentry *CISCOSIPUAMIB_Csipcfgpeertable_Csipcfgpeerentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csipcfgpeerentry *CISCOSIPUAMIB_Csipcfgpeertable_Csipcfgpeerentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cSipCfgPeerIndex"] = csipcfgpeerentry.Csipcfgpeerindex
    leafs["cSipCfgPeerOutSessionTransport"] = csipcfgpeerentry.Csipcfgpeeroutsessiontransport
    leafs["cSipCfgPeerReliable1xxRspStr"] = csipcfgpeerentry.Csipcfgpeerreliable1Xxrspstr
    leafs["cSipCfgPeerReliable1xxRspHdr"] = csipcfgpeerentry.Csipcfgpeerreliable1Xxrsphdr
    leafs["cSipCfgPeerUrlType"] = csipcfgpeerentry.Csipcfgpeerurltype
    leafs["cSipCfgPeerSwitchTransEnabled"] = csipcfgpeerentry.Csipcfgpeerswitchtransenabled
    return leafs
}

func (csipcfgpeerentry *CISCOSIPUAMIB_Csipcfgpeertable_Csipcfgpeerentry) GetBundleName() string { return "cisco_ios_xe" }

func (csipcfgpeerentry *CISCOSIPUAMIB_Csipcfgpeertable_Csipcfgpeerentry) GetYangName() string { return "cSipCfgPeerEntry" }

func (csipcfgpeerentry *CISCOSIPUAMIB_Csipcfgpeertable_Csipcfgpeerentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csipcfgpeerentry *CISCOSIPUAMIB_Csipcfgpeertable_Csipcfgpeerentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csipcfgpeerentry *CISCOSIPUAMIB_Csipcfgpeertable_Csipcfgpeerentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csipcfgpeerentry *CISCOSIPUAMIB_Csipcfgpeertable_Csipcfgpeerentry) SetParent(parent types.Entity) { csipcfgpeerentry.parent = parent }

func (csipcfgpeerentry *CISCOSIPUAMIB_Csipcfgpeertable_Csipcfgpeerentry) GetParent() types.Entity { return csipcfgpeerentry.parent }

func (csipcfgpeerentry *CISCOSIPUAMIB_Csipcfgpeertable_Csipcfgpeerentry) GetParentYangName() string { return "cSipCfgPeerTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // A row in the cSipCfgStatusCauseTable.  Entries cannot be created or
    // destroyed by the actions of any NMS. The type is slice of
    // CISCOSIPUAMIB_Csipcfgstatuscausetable_Csipcfgstatuscauseentry.
    Csipcfgstatuscauseentry []CISCOSIPUAMIB_Csipcfgstatuscausetable_Csipcfgstatuscauseentry
}

func (csipcfgstatuscausetable *CISCOSIPUAMIB_Csipcfgstatuscausetable) GetFilter() yfilter.YFilter { return csipcfgstatuscausetable.YFilter }

func (csipcfgstatuscausetable *CISCOSIPUAMIB_Csipcfgstatuscausetable) SetFilter(yf yfilter.YFilter) { csipcfgstatuscausetable.YFilter = yf }

func (csipcfgstatuscausetable *CISCOSIPUAMIB_Csipcfgstatuscausetable) GetGoName(yname string) string {
    if yname == "cSipCfgStatusCauseEntry" { return "Csipcfgstatuscauseentry" }
    return ""
}

func (csipcfgstatuscausetable *CISCOSIPUAMIB_Csipcfgstatuscausetable) GetSegmentPath() string {
    return "cSipCfgStatusCauseTable"
}

func (csipcfgstatuscausetable *CISCOSIPUAMIB_Csipcfgstatuscausetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cSipCfgStatusCauseEntry" {
        for _, c := range csipcfgstatuscausetable.Csipcfgstatuscauseentry {
            if csipcfgstatuscausetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOSIPUAMIB_Csipcfgstatuscausetable_Csipcfgstatuscauseentry{}
        csipcfgstatuscausetable.Csipcfgstatuscauseentry = append(csipcfgstatuscausetable.Csipcfgstatuscauseentry, child)
        return &csipcfgstatuscausetable.Csipcfgstatuscauseentry[len(csipcfgstatuscausetable.Csipcfgstatuscauseentry)-1]
    }
    return nil
}

func (csipcfgstatuscausetable *CISCOSIPUAMIB_Csipcfgstatuscausetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range csipcfgstatuscausetable.Csipcfgstatuscauseentry {
        children[csipcfgstatuscausetable.Csipcfgstatuscauseentry[i].GetSegmentPath()] = &csipcfgstatuscausetable.Csipcfgstatuscauseentry[i]
    }
    return children
}

func (csipcfgstatuscausetable *CISCOSIPUAMIB_Csipcfgstatuscausetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (csipcfgstatuscausetable *CISCOSIPUAMIB_Csipcfgstatuscausetable) GetBundleName() string { return "cisco_ios_xe" }

func (csipcfgstatuscausetable *CISCOSIPUAMIB_Csipcfgstatuscausetable) GetYangName() string { return "cSipCfgStatusCauseTable" }

func (csipcfgstatuscausetable *CISCOSIPUAMIB_Csipcfgstatuscausetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csipcfgstatuscausetable *CISCOSIPUAMIB_Csipcfgstatuscausetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csipcfgstatuscausetable *CISCOSIPUAMIB_Csipcfgstatuscausetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csipcfgstatuscausetable *CISCOSIPUAMIB_Csipcfgstatuscausetable) SetParent(parent types.Entity) { csipcfgstatuscausetable.parent = parent }

func (csipcfgstatuscausetable *CISCOSIPUAMIB_Csipcfgstatuscausetable) GetParent() types.Entity { return csipcfgstatuscausetable.parent }

func (csipcfgstatuscausetable *CISCOSIPUAMIB_Csipcfgstatuscausetable) GetParentYangName() string { return "CISCO-SIP-UA-MIB" }

// CISCOSIPUAMIB_Csipcfgstatuscausetable_Csipcfgstatuscauseentry
// A row in the cSipCfgStatusCauseTable.  Entries cannot be
// created or destroyed by the actions of any NMS.
type CISCOSIPUAMIB_Csipcfgstatuscausetable_Csipcfgstatuscauseentry struct {
    parent types.Entity
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

func (csipcfgstatuscauseentry *CISCOSIPUAMIB_Csipcfgstatuscausetable_Csipcfgstatuscauseentry) GetFilter() yfilter.YFilter { return csipcfgstatuscauseentry.YFilter }

func (csipcfgstatuscauseentry *CISCOSIPUAMIB_Csipcfgstatuscausetable_Csipcfgstatuscauseentry) SetFilter(yf yfilter.YFilter) { csipcfgstatuscauseentry.YFilter = yf }

func (csipcfgstatuscauseentry *CISCOSIPUAMIB_Csipcfgstatuscausetable_Csipcfgstatuscauseentry) GetGoName(yname string) string {
    if yname == "cSipCfgStatusCodeIndex" { return "Csipcfgstatuscodeindex" }
    if yname == "cSipCfgPstnCause" { return "Csipcfgpstncause" }
    if yname == "cSipCfgStatusCauseStoreStatus" { return "Csipcfgstatuscausestorestatus" }
    return ""
}

func (csipcfgstatuscauseentry *CISCOSIPUAMIB_Csipcfgstatuscausetable_Csipcfgstatuscauseentry) GetSegmentPath() string {
    return "cSipCfgStatusCauseEntry" + "[cSipCfgStatusCodeIndex='" + fmt.Sprintf("%v", csipcfgstatuscauseentry.Csipcfgstatuscodeindex) + "']"
}

func (csipcfgstatuscauseentry *CISCOSIPUAMIB_Csipcfgstatuscausetable_Csipcfgstatuscauseentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csipcfgstatuscauseentry *CISCOSIPUAMIB_Csipcfgstatuscausetable_Csipcfgstatuscauseentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csipcfgstatuscauseentry *CISCOSIPUAMIB_Csipcfgstatuscausetable_Csipcfgstatuscauseentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cSipCfgStatusCodeIndex"] = csipcfgstatuscauseentry.Csipcfgstatuscodeindex
    leafs["cSipCfgPstnCause"] = csipcfgstatuscauseentry.Csipcfgpstncause
    leafs["cSipCfgStatusCauseStoreStatus"] = csipcfgstatuscauseentry.Csipcfgstatuscausestorestatus
    return leafs
}

func (csipcfgstatuscauseentry *CISCOSIPUAMIB_Csipcfgstatuscausetable_Csipcfgstatuscauseentry) GetBundleName() string { return "cisco_ios_xe" }

func (csipcfgstatuscauseentry *CISCOSIPUAMIB_Csipcfgstatuscausetable_Csipcfgstatuscauseentry) GetYangName() string { return "cSipCfgStatusCauseEntry" }

func (csipcfgstatuscauseentry *CISCOSIPUAMIB_Csipcfgstatuscausetable_Csipcfgstatuscauseentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csipcfgstatuscauseentry *CISCOSIPUAMIB_Csipcfgstatuscausetable_Csipcfgstatuscauseentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csipcfgstatuscauseentry *CISCOSIPUAMIB_Csipcfgstatuscausetable_Csipcfgstatuscauseentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csipcfgstatuscauseentry *CISCOSIPUAMIB_Csipcfgstatuscausetable_Csipcfgstatuscauseentry) SetParent(parent types.Entity) { csipcfgstatuscauseentry.parent = parent }

func (csipcfgstatuscauseentry *CISCOSIPUAMIB_Csipcfgstatuscausetable_Csipcfgstatuscauseentry) GetParent() types.Entity { return csipcfgstatuscauseentry.parent }

func (csipcfgstatuscauseentry *CISCOSIPUAMIB_Csipcfgstatuscausetable_Csipcfgstatuscauseentry) GetParentYangName() string { return "cSipCfgStatusCauseTable" }

// CISCOSIPUAMIB_Csipcfgcausestatustable
// This table contains PSTN cause code to SIP status code
// mapping configuration.   Inbound PSTN signalling messages
// that will result in outbound SIP response messages 
// will have the PSTN cause codes transposed into the
// SIP status codes as prescribed by the contents of this 
// table.
type CISCOSIPUAMIB_Csipcfgcausestatustable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A row in the cSipCfgCauseStatusTable. Entries cannot be created or
    // destroyed by the actions of any NMS. The type is slice of
    // CISCOSIPUAMIB_Csipcfgcausestatustable_Csipcfgcausestatusentry.
    Csipcfgcausestatusentry []CISCOSIPUAMIB_Csipcfgcausestatustable_Csipcfgcausestatusentry
}

func (csipcfgcausestatustable *CISCOSIPUAMIB_Csipcfgcausestatustable) GetFilter() yfilter.YFilter { return csipcfgcausestatustable.YFilter }

func (csipcfgcausestatustable *CISCOSIPUAMIB_Csipcfgcausestatustable) SetFilter(yf yfilter.YFilter) { csipcfgcausestatustable.YFilter = yf }

func (csipcfgcausestatustable *CISCOSIPUAMIB_Csipcfgcausestatustable) GetGoName(yname string) string {
    if yname == "cSipCfgCauseStatusEntry" { return "Csipcfgcausestatusentry" }
    return ""
}

func (csipcfgcausestatustable *CISCOSIPUAMIB_Csipcfgcausestatustable) GetSegmentPath() string {
    return "cSipCfgCauseStatusTable"
}

func (csipcfgcausestatustable *CISCOSIPUAMIB_Csipcfgcausestatustable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cSipCfgCauseStatusEntry" {
        for _, c := range csipcfgcausestatustable.Csipcfgcausestatusentry {
            if csipcfgcausestatustable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOSIPUAMIB_Csipcfgcausestatustable_Csipcfgcausestatusentry{}
        csipcfgcausestatustable.Csipcfgcausestatusentry = append(csipcfgcausestatustable.Csipcfgcausestatusentry, child)
        return &csipcfgcausestatustable.Csipcfgcausestatusentry[len(csipcfgcausestatustable.Csipcfgcausestatusentry)-1]
    }
    return nil
}

func (csipcfgcausestatustable *CISCOSIPUAMIB_Csipcfgcausestatustable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range csipcfgcausestatustable.Csipcfgcausestatusentry {
        children[csipcfgcausestatustable.Csipcfgcausestatusentry[i].GetSegmentPath()] = &csipcfgcausestatustable.Csipcfgcausestatusentry[i]
    }
    return children
}

func (csipcfgcausestatustable *CISCOSIPUAMIB_Csipcfgcausestatustable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (csipcfgcausestatustable *CISCOSIPUAMIB_Csipcfgcausestatustable) GetBundleName() string { return "cisco_ios_xe" }

func (csipcfgcausestatustable *CISCOSIPUAMIB_Csipcfgcausestatustable) GetYangName() string { return "cSipCfgCauseStatusTable" }

func (csipcfgcausestatustable *CISCOSIPUAMIB_Csipcfgcausestatustable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csipcfgcausestatustable *CISCOSIPUAMIB_Csipcfgcausestatustable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csipcfgcausestatustable *CISCOSIPUAMIB_Csipcfgcausestatustable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csipcfgcausestatustable *CISCOSIPUAMIB_Csipcfgcausestatustable) SetParent(parent types.Entity) { csipcfgcausestatustable.parent = parent }

func (csipcfgcausestatustable *CISCOSIPUAMIB_Csipcfgcausestatustable) GetParent() types.Entity { return csipcfgcausestatustable.parent }

func (csipcfgcausestatustable *CISCOSIPUAMIB_Csipcfgcausestatustable) GetParentYangName() string { return "CISCO-SIP-UA-MIB" }

// CISCOSIPUAMIB_Csipcfgcausestatustable_Csipcfgcausestatusentry
// A row in the cSipCfgCauseStatusTable. Entries cannot be
// created or destroyed by the actions of any NMS.
type CISCOSIPUAMIB_Csipcfgcausestatustable_Csipcfgcausestatusentry struct {
    parent types.Entity
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

func (csipcfgcausestatusentry *CISCOSIPUAMIB_Csipcfgcausestatustable_Csipcfgcausestatusentry) GetFilter() yfilter.YFilter { return csipcfgcausestatusentry.YFilter }

func (csipcfgcausestatusentry *CISCOSIPUAMIB_Csipcfgcausestatustable_Csipcfgcausestatusentry) SetFilter(yf yfilter.YFilter) { csipcfgcausestatusentry.YFilter = yf }

func (csipcfgcausestatusentry *CISCOSIPUAMIB_Csipcfgcausestatustable_Csipcfgcausestatusentry) GetGoName(yname string) string {
    if yname == "cSipCfgPstnCauseIndex" { return "Csipcfgpstncauseindex" }
    if yname == "cSipCfgStatusCode" { return "Csipcfgstatuscode" }
    if yname == "cSipCfgCauseStatusStoreStatus" { return "Csipcfgcausestatusstorestatus" }
    return ""
}

func (csipcfgcausestatusentry *CISCOSIPUAMIB_Csipcfgcausestatustable_Csipcfgcausestatusentry) GetSegmentPath() string {
    return "cSipCfgCauseStatusEntry" + "[cSipCfgPstnCauseIndex='" + fmt.Sprintf("%v", csipcfgcausestatusentry.Csipcfgpstncauseindex) + "']"
}

func (csipcfgcausestatusentry *CISCOSIPUAMIB_Csipcfgcausestatustable_Csipcfgcausestatusentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csipcfgcausestatusentry *CISCOSIPUAMIB_Csipcfgcausestatustable_Csipcfgcausestatusentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csipcfgcausestatusentry *CISCOSIPUAMIB_Csipcfgcausestatustable_Csipcfgcausestatusentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cSipCfgPstnCauseIndex"] = csipcfgcausestatusentry.Csipcfgpstncauseindex
    leafs["cSipCfgStatusCode"] = csipcfgcausestatusentry.Csipcfgstatuscode
    leafs["cSipCfgCauseStatusStoreStatus"] = csipcfgcausestatusentry.Csipcfgcausestatusstorestatus
    return leafs
}

func (csipcfgcausestatusentry *CISCOSIPUAMIB_Csipcfgcausestatustable_Csipcfgcausestatusentry) GetBundleName() string { return "cisco_ios_xe" }

func (csipcfgcausestatusentry *CISCOSIPUAMIB_Csipcfgcausestatustable_Csipcfgcausestatusentry) GetYangName() string { return "cSipCfgCauseStatusEntry" }

func (csipcfgcausestatusentry *CISCOSIPUAMIB_Csipcfgcausestatustable_Csipcfgcausestatusentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csipcfgcausestatusentry *CISCOSIPUAMIB_Csipcfgcausestatustable_Csipcfgcausestatusentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csipcfgcausestatusentry *CISCOSIPUAMIB_Csipcfgcausestatustable_Csipcfgcausestatusentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csipcfgcausestatusentry *CISCOSIPUAMIB_Csipcfgcausestatustable_Csipcfgcausestatusentry) SetParent(parent types.Entity) { csipcfgcausestatusentry.parent = parent }

func (csipcfgcausestatusentry *CISCOSIPUAMIB_Csipcfgcausestatustable_Csipcfgcausestatusentry) GetParent() types.Entity { return csipcfgcausestatusentry.parent }

func (csipcfgcausestatusentry *CISCOSIPUAMIB_Csipcfgcausestatustable_Csipcfgcausestatusentry) GetParentYangName() string { return "cSipCfgCauseStatusTable" }

// CISCOSIPUAMIB_Csipstatssuccessoktable
// This table contains statistics for sent and
// received 200 Ok response messages.  The 
// statistics are kept on per SIP method basis.
type CISCOSIPUAMIB_Csipstatssuccessoktable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A row in the cSipStatsSuccessOkTable.  There is  an entry for each SIP
    // method for which 200 Ok  inbound and/or outbound statistics are kept. The
    // type is slice of
    // CISCOSIPUAMIB_Csipstatssuccessoktable_Csipstatssuccessokentry.
    Csipstatssuccessokentry []CISCOSIPUAMIB_Csipstatssuccessoktable_Csipstatssuccessokentry
}

func (csipstatssuccessoktable *CISCOSIPUAMIB_Csipstatssuccessoktable) GetFilter() yfilter.YFilter { return csipstatssuccessoktable.YFilter }

func (csipstatssuccessoktable *CISCOSIPUAMIB_Csipstatssuccessoktable) SetFilter(yf yfilter.YFilter) { csipstatssuccessoktable.YFilter = yf }

func (csipstatssuccessoktable *CISCOSIPUAMIB_Csipstatssuccessoktable) GetGoName(yname string) string {
    if yname == "cSipStatsSuccessOkEntry" { return "Csipstatssuccessokentry" }
    return ""
}

func (csipstatssuccessoktable *CISCOSIPUAMIB_Csipstatssuccessoktable) GetSegmentPath() string {
    return "cSipStatsSuccessOkTable"
}

func (csipstatssuccessoktable *CISCOSIPUAMIB_Csipstatssuccessoktable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cSipStatsSuccessOkEntry" {
        for _, c := range csipstatssuccessoktable.Csipstatssuccessokentry {
            if csipstatssuccessoktable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOSIPUAMIB_Csipstatssuccessoktable_Csipstatssuccessokentry{}
        csipstatssuccessoktable.Csipstatssuccessokentry = append(csipstatssuccessoktable.Csipstatssuccessokentry, child)
        return &csipstatssuccessoktable.Csipstatssuccessokentry[len(csipstatssuccessoktable.Csipstatssuccessokentry)-1]
    }
    return nil
}

func (csipstatssuccessoktable *CISCOSIPUAMIB_Csipstatssuccessoktable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range csipstatssuccessoktable.Csipstatssuccessokentry {
        children[csipstatssuccessoktable.Csipstatssuccessokentry[i].GetSegmentPath()] = &csipstatssuccessoktable.Csipstatssuccessokentry[i]
    }
    return children
}

func (csipstatssuccessoktable *CISCOSIPUAMIB_Csipstatssuccessoktable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (csipstatssuccessoktable *CISCOSIPUAMIB_Csipstatssuccessoktable) GetBundleName() string { return "cisco_ios_xe" }

func (csipstatssuccessoktable *CISCOSIPUAMIB_Csipstatssuccessoktable) GetYangName() string { return "cSipStatsSuccessOkTable" }

func (csipstatssuccessoktable *CISCOSIPUAMIB_Csipstatssuccessoktable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csipstatssuccessoktable *CISCOSIPUAMIB_Csipstatssuccessoktable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csipstatssuccessoktable *CISCOSIPUAMIB_Csipstatssuccessoktable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csipstatssuccessoktable *CISCOSIPUAMIB_Csipstatssuccessoktable) SetParent(parent types.Entity) { csipstatssuccessoktable.parent = parent }

func (csipstatssuccessoktable *CISCOSIPUAMIB_Csipstatssuccessoktable) GetParent() types.Entity { return csipstatssuccessoktable.parent }

func (csipstatssuccessoktable *CISCOSIPUAMIB_Csipstatssuccessoktable) GetParentYangName() string { return "CISCO-SIP-UA-MIB" }

// CISCOSIPUAMIB_Csipstatssuccessoktable_Csipstatssuccessokentry
// A row in the cSipStatsSuccessOkTable.  There is 
// an entry for each SIP method for which 200 Ok 
// inbound and/or outbound statistics are kept.
type CISCOSIPUAMIB_Csipstatssuccessoktable_Csipstatssuccessokentry struct {
    parent types.Entity
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

func (csipstatssuccessokentry *CISCOSIPUAMIB_Csipstatssuccessoktable_Csipstatssuccessokentry) GetFilter() yfilter.YFilter { return csipstatssuccessokentry.YFilter }

func (csipstatssuccessokentry *CISCOSIPUAMIB_Csipstatssuccessoktable_Csipstatssuccessokentry) SetFilter(yf yfilter.YFilter) { csipstatssuccessokentry.YFilter = yf }

func (csipstatssuccessokentry *CISCOSIPUAMIB_Csipstatssuccessoktable_Csipstatssuccessokentry) GetGoName(yname string) string {
    if yname == "cSipStatsSuccessOkMethod" { return "Csipstatssuccessokmethod" }
    if yname == "cSipStatsSuccessOkInbounds" { return "Csipstatssuccessokinbounds" }
    if yname == "cSipStatsSuccessOkOutbounds" { return "Csipstatssuccessokoutbounds" }
    return ""
}

func (csipstatssuccessokentry *CISCOSIPUAMIB_Csipstatssuccessoktable_Csipstatssuccessokentry) GetSegmentPath() string {
    return "cSipStatsSuccessOkEntry" + "[cSipStatsSuccessOkMethod='" + fmt.Sprintf("%v", csipstatssuccessokentry.Csipstatssuccessokmethod) + "']"
}

func (csipstatssuccessokentry *CISCOSIPUAMIB_Csipstatssuccessoktable_Csipstatssuccessokentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csipstatssuccessokentry *CISCOSIPUAMIB_Csipstatssuccessoktable_Csipstatssuccessokentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csipstatssuccessokentry *CISCOSIPUAMIB_Csipstatssuccessoktable_Csipstatssuccessokentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cSipStatsSuccessOkMethod"] = csipstatssuccessokentry.Csipstatssuccessokmethod
    leafs["cSipStatsSuccessOkInbounds"] = csipstatssuccessokentry.Csipstatssuccessokinbounds
    leafs["cSipStatsSuccessOkOutbounds"] = csipstatssuccessokentry.Csipstatssuccessokoutbounds
    return leafs
}

func (csipstatssuccessokentry *CISCOSIPUAMIB_Csipstatssuccessoktable_Csipstatssuccessokentry) GetBundleName() string { return "cisco_ios_xe" }

func (csipstatssuccessokentry *CISCOSIPUAMIB_Csipstatssuccessoktable_Csipstatssuccessokentry) GetYangName() string { return "cSipStatsSuccessOkEntry" }

func (csipstatssuccessokentry *CISCOSIPUAMIB_Csipstatssuccessoktable_Csipstatssuccessokentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csipstatssuccessokentry *CISCOSIPUAMIB_Csipstatssuccessoktable_Csipstatssuccessokentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csipstatssuccessokentry *CISCOSIPUAMIB_Csipstatssuccessoktable_Csipstatssuccessokentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csipstatssuccessokentry *CISCOSIPUAMIB_Csipstatssuccessoktable_Csipstatssuccessokentry) SetParent(parent types.Entity) { csipstatssuccessokentry.parent = parent }

func (csipstatssuccessokentry *CISCOSIPUAMIB_Csipstatssuccessoktable_Csipstatssuccessokentry) GetParent() types.Entity { return csipstatssuccessokentry.parent }

func (csipstatssuccessokentry *CISCOSIPUAMIB_Csipstatssuccessoktable_Csipstatssuccessokentry) GetParentYangName() string { return "cSipStatsSuccessOkTable" }

