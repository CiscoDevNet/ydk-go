// The MIB module for managing Trunk Media Gateway.  
// 
// A Media Gateway is a network element that provides conversion 
// between the audio signals carried on telephone circuits and 
// data packets carried over the Internet or over other packet 
// data networks. 
// 
// Trunk Media Gateway interface is between the telephone network 
// and a Voice over IP/ATM network. 
// The interface on a Trunk Gateway terminates a trunk connected 
// to PSTN switch (e.g., Class 5, Class 4, etc.).
// 
// Media Gateways use a call control architecture where the call
// control 'intelligence' is outside the gateways and handled by
// external call control elements, called Media Gateway 
// Controllers (MGCs). 
// The MGCs or Call Agents, synchronize with each other to 
// send coherent commands to the gateways under their control.
// 
// MGCs use master/slave protocols to command the gateways under 
// their control.  Examples of these protocols are:
//   *  Simple Gateway Control Protocol
//   *  Media Gateway Control Protocol
//   *  Megaco (H.248)
//   *  Simple Resource Control Protocol
// 
// To connect MG to MGCs using these control protocols through 
// an IP/UDP Ports which must be configured. To resolve IP 
// Addresses, DNS name services may be used.
package cisco_media_gateway_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_media_gateway_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-MEDIA-GATEWAY-MIB CISCO-MEDIA-GATEWAY-MIB}", reflect.TypeOf(CISCOMEDIAGATEWAYMIB{}))
    ydk.RegisterEntity("CISCO-MEDIA-GATEWAY-MIB:CISCO-MEDIA-GATEWAY-MIB", reflect.TypeOf(CISCOMEDIAGATEWAYMIB{}))
}

// CCallControlJitterDelayMode represents            which is specified by jitter nominal delay.
type CCallControlJitterDelayMode string

const (
    CCallControlJitterDelayMode_adaptive CCallControlJitterDelayMode = "adaptive"

    CCallControlJitterDelayMode_fixed CCallControlJitterDelayMode = "fixed"
)

// CGwAdminState represents       New connections would be blocked.
type CGwAdminState string

const (
    CGwAdminState_inService CGwAdminState = "inService"

    CGwAdminState_forcedOutOfService CGwAdminState = "forcedOutOfService"

    CGwAdminState_gracefulOutOfService CGwAdminState = "gracefulOutOfService"
)

// CGwServiceState represents     No new connections are allowed.
type CGwServiceState string

const (
    CGwServiceState_inService CGwServiceState = "inService"

    CGwServiceState_forcedOutOfService CGwServiceState = "forcedOutOfService"

    CGwServiceState_gracefulOutOfService CGwServiceState = "gracefulOutOfService"
)

// CISCOMEDIAGATEWAYMIB
type CISCOMEDIAGATEWAYMIB struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This table contains the global media gateway parameters information. It
    // supports the modification of the global media gateway  parameters.
    Cmediagwtable CISCOMEDIAGATEWAYMIB_Cmediagwtable

    // This table contains the available signaling protocols that are supported by
    // the media gateway for communication with MGCs.
    Cmgwsignalprotocoltable CISCOMEDIAGATEWAYMIB_Cmgwsignalprotocoltable

    // This table contains a list of media gateway IP address and the IP address
    // associated interface information.  If IP address associated interface is
    // PVC, only  aal5 control PVC or aal5 bearer PVC are valid.        When the
    // PVC is aal5 control, the IP address is used to  communicate to MGC; when
    // the PVC is aal5 bearer, the IP address is used to communicate to other
    // gateway. The PVC information is kept in cwAtmChanExtConfigTable: 
    // cwacChanPvcType:      aal5/aal2/aal1  cwacChanApplication: 
    // control/bearer/signaling  If IP address associated interface is not PVC,
    // refer to the  IP addresses associated interface table for the usage of IP
    // address.
    Cmediagwipconfigtable CISCOMEDIAGATEWAYMIB_Cmediagwipconfigtable

    // This table provides the domain names which are configured by  users.  The
    // domain names can be used to represent IP addresses  for:     gateway    
    // External DNS name server     MGC (call agent) .
    Cmediagwdomainnameconfigtable CISCOMEDIAGATEWAYMIB_Cmediagwdomainnameconfigtable

    // There is only one DNS name server on a gateway and the domain name of the
    // DNS name server is put on  cMediaGwDomainNameConfigTable with 'dnsServer
    // (2)'.  There could be multi IP addresses are associated with the DNS name
    // server, this table is used to store these IP  addresses.  If any domain
    // name using external resolution, the last entry of this table is not allowed
    // to be deleted.
    Cmediagwdnsipconfigtable CISCOMEDIAGATEWAYMIB_Cmediagwdnsipconfigtable

    // This table is for managing LIF (Logical Interface)  in a media gateway.  
    // LIF is a logical interface which groups the TDM  DSx1s associated with a
    // set of packet resource partitions  (PVCs) in a media gateway.  LIF is used
    // for: 1. VoIP switching  2. VoATM switching .
    Cmgwliftable CISCOMEDIAGATEWAYMIB_Cmgwliftable

    // This table defines general call control attributes for the media gateway.
    Cmediagwcallcontrolconfigtable CISCOMEDIAGATEWAYMIB_Cmediagwcallcontrolconfigtable

    // This table stores the gateway resource statistics information.
    Cmediagwrscstatstable CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable
}

func (cISCOMEDIAGATEWAYMIB *CISCOMEDIAGATEWAYMIB) GetFilter() yfilter.YFilter { return cISCOMEDIAGATEWAYMIB.YFilter }

func (cISCOMEDIAGATEWAYMIB *CISCOMEDIAGATEWAYMIB) SetFilter(yf yfilter.YFilter) { cISCOMEDIAGATEWAYMIB.YFilter = yf }

func (cISCOMEDIAGATEWAYMIB *CISCOMEDIAGATEWAYMIB) GetGoName(yname string) string {
    if yname == "cMediaGwTable" { return "Cmediagwtable" }
    if yname == "cmgwSignalProtocolTable" { return "Cmgwsignalprotocoltable" }
    if yname == "cMediaGwIpConfigTable" { return "Cmediagwipconfigtable" }
    if yname == "cMediaGwDomainNameConfigTable" { return "Cmediagwdomainnameconfigtable" }
    if yname == "cMediaGwDnsIpConfigTable" { return "Cmediagwdnsipconfigtable" }
    if yname == "cmgwLifTable" { return "Cmgwliftable" }
    if yname == "cMediaGwCallControlConfigTable" { return "Cmediagwcallcontrolconfigtable" }
    if yname == "cMediaGwRscStatsTable" { return "Cmediagwrscstatstable" }
    return ""
}

func (cISCOMEDIAGATEWAYMIB *CISCOMEDIAGATEWAYMIB) GetSegmentPath() string {
    return "CISCO-MEDIA-GATEWAY-MIB:CISCO-MEDIA-GATEWAY-MIB"
}

func (cISCOMEDIAGATEWAYMIB *CISCOMEDIAGATEWAYMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cMediaGwTable" {
        return &cISCOMEDIAGATEWAYMIB.Cmediagwtable
    }
    if childYangName == "cmgwSignalProtocolTable" {
        return &cISCOMEDIAGATEWAYMIB.Cmgwsignalprotocoltable
    }
    if childYangName == "cMediaGwIpConfigTable" {
        return &cISCOMEDIAGATEWAYMIB.Cmediagwipconfigtable
    }
    if childYangName == "cMediaGwDomainNameConfigTable" {
        return &cISCOMEDIAGATEWAYMIB.Cmediagwdomainnameconfigtable
    }
    if childYangName == "cMediaGwDnsIpConfigTable" {
        return &cISCOMEDIAGATEWAYMIB.Cmediagwdnsipconfigtable
    }
    if childYangName == "cmgwLifTable" {
        return &cISCOMEDIAGATEWAYMIB.Cmgwliftable
    }
    if childYangName == "cMediaGwCallControlConfigTable" {
        return &cISCOMEDIAGATEWAYMIB.Cmediagwcallcontrolconfigtable
    }
    if childYangName == "cMediaGwRscStatsTable" {
        return &cISCOMEDIAGATEWAYMIB.Cmediagwrscstatstable
    }
    return nil
}

func (cISCOMEDIAGATEWAYMIB *CISCOMEDIAGATEWAYMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cMediaGwTable"] = &cISCOMEDIAGATEWAYMIB.Cmediagwtable
    children["cmgwSignalProtocolTable"] = &cISCOMEDIAGATEWAYMIB.Cmgwsignalprotocoltable
    children["cMediaGwIpConfigTable"] = &cISCOMEDIAGATEWAYMIB.Cmediagwipconfigtable
    children["cMediaGwDomainNameConfigTable"] = &cISCOMEDIAGATEWAYMIB.Cmediagwdomainnameconfigtable
    children["cMediaGwDnsIpConfigTable"] = &cISCOMEDIAGATEWAYMIB.Cmediagwdnsipconfigtable
    children["cmgwLifTable"] = &cISCOMEDIAGATEWAYMIB.Cmgwliftable
    children["cMediaGwCallControlConfigTable"] = &cISCOMEDIAGATEWAYMIB.Cmediagwcallcontrolconfigtable
    children["cMediaGwRscStatsTable"] = &cISCOMEDIAGATEWAYMIB.Cmediagwrscstatstable
    return children
}

func (cISCOMEDIAGATEWAYMIB *CISCOMEDIAGATEWAYMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOMEDIAGATEWAYMIB *CISCOMEDIAGATEWAYMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOMEDIAGATEWAYMIB *CISCOMEDIAGATEWAYMIB) GetYangName() string { return "CISCO-MEDIA-GATEWAY-MIB" }

func (cISCOMEDIAGATEWAYMIB *CISCOMEDIAGATEWAYMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOMEDIAGATEWAYMIB *CISCOMEDIAGATEWAYMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOMEDIAGATEWAYMIB *CISCOMEDIAGATEWAYMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOMEDIAGATEWAYMIB *CISCOMEDIAGATEWAYMIB) SetParent(parent types.Entity) { cISCOMEDIAGATEWAYMIB.parent = parent }

func (cISCOMEDIAGATEWAYMIB *CISCOMEDIAGATEWAYMIB) GetParent() types.Entity { return cISCOMEDIAGATEWAYMIB.parent }

func (cISCOMEDIAGATEWAYMIB *CISCOMEDIAGATEWAYMIB) GetParentYangName() string { return "CISCO-MEDIA-GATEWAY-MIB" }

// CISCOMEDIAGATEWAYMIB_Cmediagwtable
// This table contains the global media gateway parameters
// information.
// It supports the modification of the global media gateway 
// parameters.
type CISCOMEDIAGATEWAYMIB_Cmediagwtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A Media Gateway Entry.   At system power-up, an entry is created by the
    // agent  if the system detects a media gateway module has been added  to the
    // system, and an entry is deleted if the entry associated media gateway
    // module has been removed from the system. The type is slice of
    // CISCOMEDIAGATEWAYMIB_Cmediagwtable_Cmediagwentry.
    Cmediagwentry []CISCOMEDIAGATEWAYMIB_Cmediagwtable_Cmediagwentry
}

func (cmediagwtable *CISCOMEDIAGATEWAYMIB_Cmediagwtable) GetFilter() yfilter.YFilter { return cmediagwtable.YFilter }

func (cmediagwtable *CISCOMEDIAGATEWAYMIB_Cmediagwtable) SetFilter(yf yfilter.YFilter) { cmediagwtable.YFilter = yf }

func (cmediagwtable *CISCOMEDIAGATEWAYMIB_Cmediagwtable) GetGoName(yname string) string {
    if yname == "cMediaGwEntry" { return "Cmediagwentry" }
    return ""
}

func (cmediagwtable *CISCOMEDIAGATEWAYMIB_Cmediagwtable) GetSegmentPath() string {
    return "cMediaGwTable"
}

func (cmediagwtable *CISCOMEDIAGATEWAYMIB_Cmediagwtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cMediaGwEntry" {
        for _, c := range cmediagwtable.Cmediagwentry {
            if cmediagwtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOMEDIAGATEWAYMIB_Cmediagwtable_Cmediagwentry{}
        cmediagwtable.Cmediagwentry = append(cmediagwtable.Cmediagwentry, child)
        return &cmediagwtable.Cmediagwentry[len(cmediagwtable.Cmediagwentry)-1]
    }
    return nil
}

func (cmediagwtable *CISCOMEDIAGATEWAYMIB_Cmediagwtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cmediagwtable.Cmediagwentry {
        children[cmediagwtable.Cmediagwentry[i].GetSegmentPath()] = &cmediagwtable.Cmediagwentry[i]
    }
    return children
}

func (cmediagwtable *CISCOMEDIAGATEWAYMIB_Cmediagwtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cmediagwtable *CISCOMEDIAGATEWAYMIB_Cmediagwtable) GetBundleName() string { return "cisco_ios_xe" }

func (cmediagwtable *CISCOMEDIAGATEWAYMIB_Cmediagwtable) GetYangName() string { return "cMediaGwTable" }

func (cmediagwtable *CISCOMEDIAGATEWAYMIB_Cmediagwtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cmediagwtable *CISCOMEDIAGATEWAYMIB_Cmediagwtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cmediagwtable *CISCOMEDIAGATEWAYMIB_Cmediagwtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cmediagwtable *CISCOMEDIAGATEWAYMIB_Cmediagwtable) SetParent(parent types.Entity) { cmediagwtable.parent = parent }

func (cmediagwtable *CISCOMEDIAGATEWAYMIB_Cmediagwtable) GetParent() types.Entity { return cmediagwtable.parent }

func (cmediagwtable *CISCOMEDIAGATEWAYMIB_Cmediagwtable) GetParentYangName() string { return "CISCO-MEDIA-GATEWAY-MIB" }

// CISCOMEDIAGATEWAYMIB_Cmediagwtable_Cmediagwentry
// A Media Gateway Entry.  
// At system power-up, an entry is created by the agent 
// if the system detects a media gateway module has been added 
// to the system, and an entry is deleted if the entry associated
// media gateway module has been removed from the system.
type CISCOMEDIAGATEWAYMIB_Cmediagwtable_Cmediagwentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. An index that uniquely identifies an entry in the 
    // cMediaGwTable. The type is interface{} with range: 1..2147483647.
    Cmgwindex interface{}

    // This object is used to represent a domain name under which    the Media
    // Gateway could also be registered in a DNS name server.   The value of this
    // object reflects the value of  cmgwConfigDomainName from the entry with a
    // value of  'gateway(1)' for object cmgwConfigDomainNameEntity of 
    // cMediaGwDomainNameConfigTable.  If there is no entry in
    // cMediaGwDomainNameConfigTable with 'gateway(1)' of
    // cmgwConfigDomainNameEntity, then the value of this object will be empty
    // string. The type is string with length: 0..64.
    Cmgwdomainname interface{}

    // This object represents the entPhysicalIndex of the card in which media
    // gateway is running. It will contain value 0 if the entPhysicalIndex value
    // is not available or  not applicable. The type is interface{} with range:
    // 0..2147483647.
    Cmgwphysicalindex interface{}

    // This object indicates the current service state of the Media  Gateway. This
    // object is controlled by 'cmgwAdminState'  object. The type is
    // CGwServiceState.
    Cmgwservicestate interface{}

    // This object is used to change the service state of  the Media Gateway from
    // inService to outOfService or from  outOfService to inService.  The
    // resulting service state of the gateway is represented   by
    // 'cmgwServiceState'. The type is CGwAdminState.
    Cmgwadminstate interface{}

    // This object is used to represent grace period. The grace period (restart
    // delay in RSIP message) is   expressed in a number of seconds.  It means how
    // soon the gateway will be taken out of service. The value -1 indicates that
    // the grace period time is disabled. The type is interface{} with range:
    // -1..65535. Units are seconds.
    Cmgwgracetime interface{}

    // This object is used to represent the VT (sonet Virtual Tributary) counting.
    // standard - standard counting (based on Bellcore TR253) titan    - TITAN5500
    // counting (based on Tellabs TITAN 5500)  Note: 'titan' is valid only if
    // sonet line medium type        (sonetMediumType of SONET-MIB) is 'sonet' and
    // sonet path payload type (cspSonetPathPayload of       CISCO-SONET-MIB) is
    // 'vt15vc11'. The type is Cmgwvtmappingmode.
    Cmgwvtmappingmode interface{}

    // This object is used to enable or disable the source IP and port filtering
    // with MGC for security consideration as follows:   'true'  - source IP and
    // port filter is enabled    'false' - source IP and port filter is disable .
    // The type is bool.
    Cmgwsrcfilterenabled interface{}

    // This object is used to enable or disable the lawful intercept for
    // government. as follows:   'true'  - enable lawful intercept   'false' -
    // disable lawful intercept. The type is bool.
    Cmgwlawinterceptenabled interface{}

    // This object is to enable or disable V23 tone. Setting the object value to
    // 'true', will cause VXSM (Voice Switching Service Module) to detect V23
    // tone. The type is bool.
    Cmgwv23Enabled interface{}
}

func (cmediagwentry *CISCOMEDIAGATEWAYMIB_Cmediagwtable_Cmediagwentry) GetFilter() yfilter.YFilter { return cmediagwentry.YFilter }

func (cmediagwentry *CISCOMEDIAGATEWAYMIB_Cmediagwtable_Cmediagwentry) SetFilter(yf yfilter.YFilter) { cmediagwentry.YFilter = yf }

func (cmediagwentry *CISCOMEDIAGATEWAYMIB_Cmediagwtable_Cmediagwentry) GetGoName(yname string) string {
    if yname == "cmgwIndex" { return "Cmgwindex" }
    if yname == "cmgwDomainName" { return "Cmgwdomainname" }
    if yname == "cmgwPhysicalIndex" { return "Cmgwphysicalindex" }
    if yname == "cmgwServiceState" { return "Cmgwservicestate" }
    if yname == "cmgwAdminState" { return "Cmgwadminstate" }
    if yname == "cmgwGraceTime" { return "Cmgwgracetime" }
    if yname == "cmgwVtMappingMode" { return "Cmgwvtmappingmode" }
    if yname == "cmgwSrcFilterEnabled" { return "Cmgwsrcfilterenabled" }
    if yname == "cmgwLawInterceptEnabled" { return "Cmgwlawinterceptenabled" }
    if yname == "cmgwV23Enabled" { return "Cmgwv23Enabled" }
    return ""
}

func (cmediagwentry *CISCOMEDIAGATEWAYMIB_Cmediagwtable_Cmediagwentry) GetSegmentPath() string {
    return "cMediaGwEntry" + "[cmgwIndex='" + fmt.Sprintf("%v", cmediagwentry.Cmgwindex) + "']"
}

func (cmediagwentry *CISCOMEDIAGATEWAYMIB_Cmediagwtable_Cmediagwentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cmediagwentry *CISCOMEDIAGATEWAYMIB_Cmediagwtable_Cmediagwentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cmediagwentry *CISCOMEDIAGATEWAYMIB_Cmediagwtable_Cmediagwentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cmgwIndex"] = cmediagwentry.Cmgwindex
    leafs["cmgwDomainName"] = cmediagwentry.Cmgwdomainname
    leafs["cmgwPhysicalIndex"] = cmediagwentry.Cmgwphysicalindex
    leafs["cmgwServiceState"] = cmediagwentry.Cmgwservicestate
    leafs["cmgwAdminState"] = cmediagwentry.Cmgwadminstate
    leafs["cmgwGraceTime"] = cmediagwentry.Cmgwgracetime
    leafs["cmgwVtMappingMode"] = cmediagwentry.Cmgwvtmappingmode
    leafs["cmgwSrcFilterEnabled"] = cmediagwentry.Cmgwsrcfilterenabled
    leafs["cmgwLawInterceptEnabled"] = cmediagwentry.Cmgwlawinterceptenabled
    leafs["cmgwV23Enabled"] = cmediagwentry.Cmgwv23Enabled
    return leafs
}

func (cmediagwentry *CISCOMEDIAGATEWAYMIB_Cmediagwtable_Cmediagwentry) GetBundleName() string { return "cisco_ios_xe" }

func (cmediagwentry *CISCOMEDIAGATEWAYMIB_Cmediagwtable_Cmediagwentry) GetYangName() string { return "cMediaGwEntry" }

func (cmediagwentry *CISCOMEDIAGATEWAYMIB_Cmediagwtable_Cmediagwentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cmediagwentry *CISCOMEDIAGATEWAYMIB_Cmediagwtable_Cmediagwentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cmediagwentry *CISCOMEDIAGATEWAYMIB_Cmediagwtable_Cmediagwentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cmediagwentry *CISCOMEDIAGATEWAYMIB_Cmediagwtable_Cmediagwentry) SetParent(parent types.Entity) { cmediagwentry.parent = parent }

func (cmediagwentry *CISCOMEDIAGATEWAYMIB_Cmediagwtable_Cmediagwentry) GetParent() types.Entity { return cmediagwentry.parent }

func (cmediagwentry *CISCOMEDIAGATEWAYMIB_Cmediagwtable_Cmediagwentry) GetParentYangName() string { return "cMediaGwTable" }

// CISCOMEDIAGATEWAYMIB_Cmediagwtable_Cmediagwentry_Cmgwvtmappingmode represents       CISCO-SONET-MIB) is 'vt15vc11'.
type CISCOMEDIAGATEWAYMIB_Cmediagwtable_Cmediagwentry_Cmgwvtmappingmode string

const (
    CISCOMEDIAGATEWAYMIB_Cmediagwtable_Cmediagwentry_Cmgwvtmappingmode_standard CISCOMEDIAGATEWAYMIB_Cmediagwtable_Cmediagwentry_Cmgwvtmappingmode = "standard"

    CISCOMEDIAGATEWAYMIB_Cmediagwtable_Cmediagwentry_Cmgwvtmappingmode_titan CISCOMEDIAGATEWAYMIB_Cmediagwtable_Cmediagwentry_Cmgwvtmappingmode = "titan"
)

// CISCOMEDIAGATEWAYMIB_Cmgwsignalprotocoltable
// This table contains the available signaling protocols that
// are supported by the media gateway for communication with
// MGCs.
type CISCOMEDIAGATEWAYMIB_Cmgwsignalprotocoltable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry represents an signaling protocol supported by the media gateway.
    // The type is slice of
    // CISCOMEDIAGATEWAYMIB_Cmgwsignalprotocoltable_Cmgwsignalprotocolentry.
    Cmgwsignalprotocolentry []CISCOMEDIAGATEWAYMIB_Cmgwsignalprotocoltable_Cmgwsignalprotocolentry
}

func (cmgwsignalprotocoltable *CISCOMEDIAGATEWAYMIB_Cmgwsignalprotocoltable) GetFilter() yfilter.YFilter { return cmgwsignalprotocoltable.YFilter }

func (cmgwsignalprotocoltable *CISCOMEDIAGATEWAYMIB_Cmgwsignalprotocoltable) SetFilter(yf yfilter.YFilter) { cmgwsignalprotocoltable.YFilter = yf }

func (cmgwsignalprotocoltable *CISCOMEDIAGATEWAYMIB_Cmgwsignalprotocoltable) GetGoName(yname string) string {
    if yname == "cmgwSignalProtocolEntry" { return "Cmgwsignalprotocolentry" }
    return ""
}

func (cmgwsignalprotocoltable *CISCOMEDIAGATEWAYMIB_Cmgwsignalprotocoltable) GetSegmentPath() string {
    return "cmgwSignalProtocolTable"
}

func (cmgwsignalprotocoltable *CISCOMEDIAGATEWAYMIB_Cmgwsignalprotocoltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cmgwSignalProtocolEntry" {
        for _, c := range cmgwsignalprotocoltable.Cmgwsignalprotocolentry {
            if cmgwsignalprotocoltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOMEDIAGATEWAYMIB_Cmgwsignalprotocoltable_Cmgwsignalprotocolentry{}
        cmgwsignalprotocoltable.Cmgwsignalprotocolentry = append(cmgwsignalprotocoltable.Cmgwsignalprotocolentry, child)
        return &cmgwsignalprotocoltable.Cmgwsignalprotocolentry[len(cmgwsignalprotocoltable.Cmgwsignalprotocolentry)-1]
    }
    return nil
}

func (cmgwsignalprotocoltable *CISCOMEDIAGATEWAYMIB_Cmgwsignalprotocoltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cmgwsignalprotocoltable.Cmgwsignalprotocolentry {
        children[cmgwsignalprotocoltable.Cmgwsignalprotocolentry[i].GetSegmentPath()] = &cmgwsignalprotocoltable.Cmgwsignalprotocolentry[i]
    }
    return children
}

func (cmgwsignalprotocoltable *CISCOMEDIAGATEWAYMIB_Cmgwsignalprotocoltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cmgwsignalprotocoltable *CISCOMEDIAGATEWAYMIB_Cmgwsignalprotocoltable) GetBundleName() string { return "cisco_ios_xe" }

func (cmgwsignalprotocoltable *CISCOMEDIAGATEWAYMIB_Cmgwsignalprotocoltable) GetYangName() string { return "cmgwSignalProtocolTable" }

func (cmgwsignalprotocoltable *CISCOMEDIAGATEWAYMIB_Cmgwsignalprotocoltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cmgwsignalprotocoltable *CISCOMEDIAGATEWAYMIB_Cmgwsignalprotocoltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cmgwsignalprotocoltable *CISCOMEDIAGATEWAYMIB_Cmgwsignalprotocoltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cmgwsignalprotocoltable *CISCOMEDIAGATEWAYMIB_Cmgwsignalprotocoltable) SetParent(parent types.Entity) { cmgwsignalprotocoltable.parent = parent }

func (cmgwsignalprotocoltable *CISCOMEDIAGATEWAYMIB_Cmgwsignalprotocoltable) GetParent() types.Entity { return cmgwsignalprotocoltable.parent }

func (cmgwsignalprotocoltable *CISCOMEDIAGATEWAYMIB_Cmgwsignalprotocoltable) GetParentYangName() string { return "CISCO-MEDIA-GATEWAY-MIB" }

// CISCOMEDIAGATEWAYMIB_Cmgwsignalprotocoltable_Cmgwsignalprotocolentry
// Each entry represents an signaling protocol supported
// by the media gateway.
type CISCOMEDIAGATEWAYMIB_Cmgwsignalprotocoltable_Cmgwsignalprotocolentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // cisco_media_gateway_mib.CISCOMEDIAGATEWAYMIB_Cmediagwtable_Cmediagwentry_Cmgwindex
    Cmgwindex interface{}

    // This attribute is a key. An index that uniquely identifies an entry in
    // cmgwSignalProtocolTable. The type is interface{} with range: 1..65535.
    Cmgwsignalprotocolindex interface{}

    // This object is used to represent the protocol type. other - None of the
    // following types. mgcp  - Media Gateway Control Protocol h248 - Media
    // Gateway Control (ITU H.248) tgcp - Trunking Gateway Control Protocol. The
    // type is Cmgwsignalprotocol.
    Cmgwsignalprotocol interface{}

    // This object is used to represent the protocol version.  For example
    // cmgwSignalProtocol is 'mgcp(2)' and this object is string '1.0'.
    // cmgwSignalProtocol is  'h248(3)' and this object is set to '2.0'. The type
    // is string with length: 1..16.
    Cmgwsignalprotocolversion interface{}

    // This object is used to represent the UDP port associated  with the
    // protocol. If the value of cmgwSignalProtocol is 'mgcp(2)' and the value of
    // cmgwSignalProtcolVersion is '1.0', the default value of this object is
    // '2727'.  If the value of cmgwSignalProtocol is 'h248(3)' and the value of
    // cmgwSignalProtcolVersion is '1.0', the default value of this object is
    // '2944'. The type is interface{} with range: 0..65535.
    Cmgwsignalprotocolport interface{}

    // This object specifies the protocol port of the Media Gateway Controller
    // (MGC). If the value of cmgwSignalProtocol is 'mgcp(2)' or 'tgcp(4)' and the
    // value of cmgwSignalProtcolVersion is '1.0', the default value of this
    // object is '2427'. If the value of cmgwSignalProtocol is 'h248(3)' and the
    // value of cmgwSignalProtcolVersion is '1.0', the default value of this
    // object is '2944'. The type is interface{} with range: 0..65535.
    Cmgwsignalmgcprotocolport interface{}

    // This object specifies the preference of the signal protocol  supported in
    // the media gateway.  If this object is set to 0, the corresponding signal
    // protocol will not be used by the gateway.   The value of this object is
    // unique within the corresponding gateway. The entry with lower value has
    // higher preference. The type is interface{} with range: 0..255.
    Cmgwsignalprotocolpreference interface{}

    // This object specifies the protocol version used by the gateway in the
    // messages to MGC in order to exchange the service capabilities.  For example
    // cmgwSignalProtocol is 'h248(3)' and this object can be string '1' or '1.0',
    // '2' or '2.0'.   'MAX' is a special string indicating the gateway will use
    // the highest protocol version supported in the  gateway, but it can be
    // changed to lower version after  it negotiates with MGC. The final
    // negotiated protocol version will be indicated in cmgwSignalProtocolVersion.
    // The version strings other than 'MAX' can be specified for the gateway to
    // communicate with the MGC which doesn't support service capabilities
    // negotiation. For example if a MGC supports only version 1.0 MGCP, this
    // object should be set to '1' to instruct the gateway using MGCP  version 1.0
    // format messages to communicate with MGC. . The type is string with length:
    // 1..16.
    Cmgwsignalprotocolconfigver interface{}
}

func (cmgwsignalprotocolentry *CISCOMEDIAGATEWAYMIB_Cmgwsignalprotocoltable_Cmgwsignalprotocolentry) GetFilter() yfilter.YFilter { return cmgwsignalprotocolentry.YFilter }

func (cmgwsignalprotocolentry *CISCOMEDIAGATEWAYMIB_Cmgwsignalprotocoltable_Cmgwsignalprotocolentry) SetFilter(yf yfilter.YFilter) { cmgwsignalprotocolentry.YFilter = yf }

func (cmgwsignalprotocolentry *CISCOMEDIAGATEWAYMIB_Cmgwsignalprotocoltable_Cmgwsignalprotocolentry) GetGoName(yname string) string {
    if yname == "cmgwIndex" { return "Cmgwindex" }
    if yname == "cmgwSignalProtocolIndex" { return "Cmgwsignalprotocolindex" }
    if yname == "cmgwSignalProtocol" { return "Cmgwsignalprotocol" }
    if yname == "cmgwSignalProtocolVersion" { return "Cmgwsignalprotocolversion" }
    if yname == "cmgwSignalProtocolPort" { return "Cmgwsignalprotocolport" }
    if yname == "cmgwSignalMgcProtocolPort" { return "Cmgwsignalmgcprotocolport" }
    if yname == "cmgwSignalProtocolPreference" { return "Cmgwsignalprotocolpreference" }
    if yname == "cmgwSignalProtocolConfigVer" { return "Cmgwsignalprotocolconfigver" }
    return ""
}

func (cmgwsignalprotocolentry *CISCOMEDIAGATEWAYMIB_Cmgwsignalprotocoltable_Cmgwsignalprotocolentry) GetSegmentPath() string {
    return "cmgwSignalProtocolEntry" + "[cmgwIndex='" + fmt.Sprintf("%v", cmgwsignalprotocolentry.Cmgwindex) + "']" + "[cmgwSignalProtocolIndex='" + fmt.Sprintf("%v", cmgwsignalprotocolentry.Cmgwsignalprotocolindex) + "']"
}

func (cmgwsignalprotocolentry *CISCOMEDIAGATEWAYMIB_Cmgwsignalprotocoltable_Cmgwsignalprotocolentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cmgwsignalprotocolentry *CISCOMEDIAGATEWAYMIB_Cmgwsignalprotocoltable_Cmgwsignalprotocolentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cmgwsignalprotocolentry *CISCOMEDIAGATEWAYMIB_Cmgwsignalprotocoltable_Cmgwsignalprotocolentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cmgwIndex"] = cmgwsignalprotocolentry.Cmgwindex
    leafs["cmgwSignalProtocolIndex"] = cmgwsignalprotocolentry.Cmgwsignalprotocolindex
    leafs["cmgwSignalProtocol"] = cmgwsignalprotocolentry.Cmgwsignalprotocol
    leafs["cmgwSignalProtocolVersion"] = cmgwsignalprotocolentry.Cmgwsignalprotocolversion
    leafs["cmgwSignalProtocolPort"] = cmgwsignalprotocolentry.Cmgwsignalprotocolport
    leafs["cmgwSignalMgcProtocolPort"] = cmgwsignalprotocolentry.Cmgwsignalmgcprotocolport
    leafs["cmgwSignalProtocolPreference"] = cmgwsignalprotocolentry.Cmgwsignalprotocolpreference
    leafs["cmgwSignalProtocolConfigVer"] = cmgwsignalprotocolentry.Cmgwsignalprotocolconfigver
    return leafs
}

func (cmgwsignalprotocolentry *CISCOMEDIAGATEWAYMIB_Cmgwsignalprotocoltable_Cmgwsignalprotocolentry) GetBundleName() string { return "cisco_ios_xe" }

func (cmgwsignalprotocolentry *CISCOMEDIAGATEWAYMIB_Cmgwsignalprotocoltable_Cmgwsignalprotocolentry) GetYangName() string { return "cmgwSignalProtocolEntry" }

func (cmgwsignalprotocolentry *CISCOMEDIAGATEWAYMIB_Cmgwsignalprotocoltable_Cmgwsignalprotocolentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cmgwsignalprotocolentry *CISCOMEDIAGATEWAYMIB_Cmgwsignalprotocoltable_Cmgwsignalprotocolentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cmgwsignalprotocolentry *CISCOMEDIAGATEWAYMIB_Cmgwsignalprotocoltable_Cmgwsignalprotocolentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cmgwsignalprotocolentry *CISCOMEDIAGATEWAYMIB_Cmgwsignalprotocoltable_Cmgwsignalprotocolentry) SetParent(parent types.Entity) { cmgwsignalprotocolentry.parent = parent }

func (cmgwsignalprotocolentry *CISCOMEDIAGATEWAYMIB_Cmgwsignalprotocoltable_Cmgwsignalprotocolentry) GetParent() types.Entity { return cmgwsignalprotocolentry.parent }

func (cmgwsignalprotocolentry *CISCOMEDIAGATEWAYMIB_Cmgwsignalprotocoltable_Cmgwsignalprotocolentry) GetParentYangName() string { return "cmgwSignalProtocolTable" }

// CISCOMEDIAGATEWAYMIB_Cmgwsignalprotocoltable_Cmgwsignalprotocolentry_Cmgwsignalprotocol represents tgcp - Trunking Gateway Control Protocol
type CISCOMEDIAGATEWAYMIB_Cmgwsignalprotocoltable_Cmgwsignalprotocolentry_Cmgwsignalprotocol string

const (
    CISCOMEDIAGATEWAYMIB_Cmgwsignalprotocoltable_Cmgwsignalprotocolentry_Cmgwsignalprotocol_other CISCOMEDIAGATEWAYMIB_Cmgwsignalprotocoltable_Cmgwsignalprotocolentry_Cmgwsignalprotocol = "other"

    CISCOMEDIAGATEWAYMIB_Cmgwsignalprotocoltable_Cmgwsignalprotocolentry_Cmgwsignalprotocol_mgcp CISCOMEDIAGATEWAYMIB_Cmgwsignalprotocoltable_Cmgwsignalprotocolentry_Cmgwsignalprotocol = "mgcp"

    CISCOMEDIAGATEWAYMIB_Cmgwsignalprotocoltable_Cmgwsignalprotocolentry_Cmgwsignalprotocol_h248 CISCOMEDIAGATEWAYMIB_Cmgwsignalprotocoltable_Cmgwsignalprotocolentry_Cmgwsignalprotocol = "h248"

    CISCOMEDIAGATEWAYMIB_Cmgwsignalprotocoltable_Cmgwsignalprotocolentry_Cmgwsignalprotocol_tgcp CISCOMEDIAGATEWAYMIB_Cmgwsignalprotocoltable_Cmgwsignalprotocolentry_Cmgwsignalprotocol = "tgcp"
)

// CISCOMEDIAGATEWAYMIB_Cmediagwipconfigtable
// This table contains a list of media gateway IP address and
// the IP address associated interface information.
// 
// If IP address associated interface is PVC, only 
// aal5 control PVC or aal5 bearer PVC are valid.       
// When the PVC is aal5 control, the IP address is used to 
// communicate to MGC; when the PVC is aal5 bearer, the IP
// address is used to communicate to other gateway.
// The PVC information is kept in cwAtmChanExtConfigTable:
//  cwacChanPvcType:      aal5/aal2/aal1
//  cwacChanApplication:  control/bearer/signaling
// 
// If IP address associated interface is not PVC, refer to the 
// IP addresses associated interface table for the usage
// of IP address.
type CISCOMEDIAGATEWAYMIB_Cmediagwipconfigtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A Media Gateway IP configuration entry.  Each entry represents a media
    // gateway IP address for MGCs to communicate with the media gateway. The type
    // is slice of
    // CISCOMEDIAGATEWAYMIB_Cmediagwipconfigtable_Cmediagwipconfigentry.
    Cmediagwipconfigentry []CISCOMEDIAGATEWAYMIB_Cmediagwipconfigtable_Cmediagwipconfigentry
}

func (cmediagwipconfigtable *CISCOMEDIAGATEWAYMIB_Cmediagwipconfigtable) GetFilter() yfilter.YFilter { return cmediagwipconfigtable.YFilter }

func (cmediagwipconfigtable *CISCOMEDIAGATEWAYMIB_Cmediagwipconfigtable) SetFilter(yf yfilter.YFilter) { cmediagwipconfigtable.YFilter = yf }

func (cmediagwipconfigtable *CISCOMEDIAGATEWAYMIB_Cmediagwipconfigtable) GetGoName(yname string) string {
    if yname == "cMediaGwIpConfigEntry" { return "Cmediagwipconfigentry" }
    return ""
}

func (cmediagwipconfigtable *CISCOMEDIAGATEWAYMIB_Cmediagwipconfigtable) GetSegmentPath() string {
    return "cMediaGwIpConfigTable"
}

func (cmediagwipconfigtable *CISCOMEDIAGATEWAYMIB_Cmediagwipconfigtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cMediaGwIpConfigEntry" {
        for _, c := range cmediagwipconfigtable.Cmediagwipconfigentry {
            if cmediagwipconfigtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOMEDIAGATEWAYMIB_Cmediagwipconfigtable_Cmediagwipconfigentry{}
        cmediagwipconfigtable.Cmediagwipconfigentry = append(cmediagwipconfigtable.Cmediagwipconfigentry, child)
        return &cmediagwipconfigtable.Cmediagwipconfigentry[len(cmediagwipconfigtable.Cmediagwipconfigentry)-1]
    }
    return nil
}

func (cmediagwipconfigtable *CISCOMEDIAGATEWAYMIB_Cmediagwipconfigtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cmediagwipconfigtable.Cmediagwipconfigentry {
        children[cmediagwipconfigtable.Cmediagwipconfigentry[i].GetSegmentPath()] = &cmediagwipconfigtable.Cmediagwipconfigentry[i]
    }
    return children
}

func (cmediagwipconfigtable *CISCOMEDIAGATEWAYMIB_Cmediagwipconfigtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cmediagwipconfigtable *CISCOMEDIAGATEWAYMIB_Cmediagwipconfigtable) GetBundleName() string { return "cisco_ios_xe" }

func (cmediagwipconfigtable *CISCOMEDIAGATEWAYMIB_Cmediagwipconfigtable) GetYangName() string { return "cMediaGwIpConfigTable" }

func (cmediagwipconfigtable *CISCOMEDIAGATEWAYMIB_Cmediagwipconfigtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cmediagwipconfigtable *CISCOMEDIAGATEWAYMIB_Cmediagwipconfigtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cmediagwipconfigtable *CISCOMEDIAGATEWAYMIB_Cmediagwipconfigtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cmediagwipconfigtable *CISCOMEDIAGATEWAYMIB_Cmediagwipconfigtable) SetParent(parent types.Entity) { cmediagwipconfigtable.parent = parent }

func (cmediagwipconfigtable *CISCOMEDIAGATEWAYMIB_Cmediagwipconfigtable) GetParent() types.Entity { return cmediagwipconfigtable.parent }

func (cmediagwipconfigtable *CISCOMEDIAGATEWAYMIB_Cmediagwipconfigtable) GetParentYangName() string { return "CISCO-MEDIA-GATEWAY-MIB" }

// CISCOMEDIAGATEWAYMIB_Cmediagwipconfigtable_Cmediagwipconfigentry
// A Media Gateway IP configuration entry. 
// Each entry represents a media gateway IP address for MGCs
// to communicate with the media gateway.
type CISCOMEDIAGATEWAYMIB_Cmediagwipconfigtable_Cmediagwipconfigentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // cisco_media_gateway_mib.CISCOMEDIAGATEWAYMIB_Cmediagwtable_Cmediagwentry_Cmgwindex
    Cmgwindex interface{}

    // This attribute is a key. A unique index to identify each media gateway IP
    // address. The type is interface{} with range: 1..64.
    Cmgwipconfigindex interface{}

    // This object is ifIndex of the interface which is associated to the media
    // gateway IP address.  For ATM interface, the IP address should be associated
    // to an existing PVC:    cmgwIpConfigIfIndex represents port of the PVC   
    // cmgwIpConfigVpi represents VPI of the PVC    cmgwIpConfigVci represents VCI
    // of the PVC And one PVC only can be associated with one IP address.  If this
    // object is set to zero which means the IP address is not associated to any
    // interface. The type is interface{} with range: 0..2147483647.
    Cmgwipconfigifindex interface{}

    // This object represents VPI of the PVC which is associated to the IP
    // address. If the IP address is not associated to PVC, the value  of this
    // object is set to -1. The type is interface{} with range: -1..4095.
    Cmgwipconfigvpi interface{}

    // This object represents VCI of the PVC which is associated to the IP
    // address. If the IP address is not associated to PVC, the value of this
    // object is set to -1. The type is interface{} with range: -1..65535.
    Cmgwipconfigvci interface{}

    // This object is the IP address type. The type is InetAddressType.
    Cmgwipconfigaddrtype interface{}

    // The configured IP address of media gateway. This object can not be
    // modified. The type is string with length: 0..255.
    Cmgwipconfigaddress interface{}

    // This object is used to specify the number of leading one    bits which from
    // the mask to be logical-ANDed with the media   gateway address before being
    // compared to the value in the  cmgwIpCofigAddress.  Any assignment (implicit
    // or otherwise) of an instance of this object to a value x must be rejected
    // if the bitwise logical-AND of the mask formed from x with the value  of the
    // corresponding instance of the cmgwIpCofigAddress  object is not equal to
    // cmgwIpCofigAddress. The type is interface{} with range: 0..2040.
    Cmgwipconfigsubnetmask interface{}

    // This object specifies cmgwIpConfigAddress of the entry will become the
    // default gateway address. This object can be set to 'true' for only one
    // entry in the table. The type is bool.
    Cmgwipconfigdefaultgwip interface{}

    // This object specifies whether the address defined in cmgwIpConfigAddress is
    // the address mapping at the remote end of this PVC.   If this object is set
    // to 'true', the address defined in cmgwIpConfigAddress is for the remote end
    // of the PVC. If this object is set to 'false', the address defined in
    // cmgwIpConfigAddress is for the local end of the PVC. The type is bool.
    Cmgwipconfigforremotemapping interface{}

    // This object is used to add and delete an entry.  When an entry of the table
    // is created, the following  objects are mandatory:     cmgwIpConfigIfIndex  
    // cmgwIpConfigVpi     cmgwIpConfigVci     cmgwIpConfigAddress    
    // cmgwIpConfigSubnetMask  These objects can not be modified after the value
    // of this object is set to 'active'.  Modification can only be done by
    // deleting and re-adding the  entry again.  After the system verify the
    // validity of the data, it will set the cmgwIpConfigRowStatus to 'active'.
    // The type is RowStatus.
    Cmgwipconfigrowstatus interface{}
}

func (cmediagwipconfigentry *CISCOMEDIAGATEWAYMIB_Cmediagwipconfigtable_Cmediagwipconfigentry) GetFilter() yfilter.YFilter { return cmediagwipconfigentry.YFilter }

func (cmediagwipconfigentry *CISCOMEDIAGATEWAYMIB_Cmediagwipconfigtable_Cmediagwipconfigentry) SetFilter(yf yfilter.YFilter) { cmediagwipconfigentry.YFilter = yf }

func (cmediagwipconfigentry *CISCOMEDIAGATEWAYMIB_Cmediagwipconfigtable_Cmediagwipconfigentry) GetGoName(yname string) string {
    if yname == "cmgwIndex" { return "Cmgwindex" }
    if yname == "cmgwIpConfigIndex" { return "Cmgwipconfigindex" }
    if yname == "cmgwIpConfigIfIndex" { return "Cmgwipconfigifindex" }
    if yname == "cmgwIpConfigVpi" { return "Cmgwipconfigvpi" }
    if yname == "cmgwIpConfigVci" { return "Cmgwipconfigvci" }
    if yname == "cmgwIpConfigAddrType" { return "Cmgwipconfigaddrtype" }
    if yname == "cmgwIpConfigAddress" { return "Cmgwipconfigaddress" }
    if yname == "cmgwIpConfigSubnetMask" { return "Cmgwipconfigsubnetmask" }
    if yname == "cmgwIpConfigDefaultGwIp" { return "Cmgwipconfigdefaultgwip" }
    if yname == "cmgwIpConfigForRemoteMapping" { return "Cmgwipconfigforremotemapping" }
    if yname == "cmgwIpConfigRowStatus" { return "Cmgwipconfigrowstatus" }
    return ""
}

func (cmediagwipconfigentry *CISCOMEDIAGATEWAYMIB_Cmediagwipconfigtable_Cmediagwipconfigentry) GetSegmentPath() string {
    return "cMediaGwIpConfigEntry" + "[cmgwIndex='" + fmt.Sprintf("%v", cmediagwipconfigentry.Cmgwindex) + "']" + "[cmgwIpConfigIndex='" + fmt.Sprintf("%v", cmediagwipconfigentry.Cmgwipconfigindex) + "']"
}

func (cmediagwipconfigentry *CISCOMEDIAGATEWAYMIB_Cmediagwipconfigtable_Cmediagwipconfigentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cmediagwipconfigentry *CISCOMEDIAGATEWAYMIB_Cmediagwipconfigtable_Cmediagwipconfigentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cmediagwipconfigentry *CISCOMEDIAGATEWAYMIB_Cmediagwipconfigtable_Cmediagwipconfigentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cmgwIndex"] = cmediagwipconfigentry.Cmgwindex
    leafs["cmgwIpConfigIndex"] = cmediagwipconfigentry.Cmgwipconfigindex
    leafs["cmgwIpConfigIfIndex"] = cmediagwipconfigentry.Cmgwipconfigifindex
    leafs["cmgwIpConfigVpi"] = cmediagwipconfigentry.Cmgwipconfigvpi
    leafs["cmgwIpConfigVci"] = cmediagwipconfigentry.Cmgwipconfigvci
    leafs["cmgwIpConfigAddrType"] = cmediagwipconfigentry.Cmgwipconfigaddrtype
    leafs["cmgwIpConfigAddress"] = cmediagwipconfigentry.Cmgwipconfigaddress
    leafs["cmgwIpConfigSubnetMask"] = cmediagwipconfigentry.Cmgwipconfigsubnetmask
    leafs["cmgwIpConfigDefaultGwIp"] = cmediagwipconfigentry.Cmgwipconfigdefaultgwip
    leafs["cmgwIpConfigForRemoteMapping"] = cmediagwipconfigentry.Cmgwipconfigforremotemapping
    leafs["cmgwIpConfigRowStatus"] = cmediagwipconfigentry.Cmgwipconfigrowstatus
    return leafs
}

func (cmediagwipconfigentry *CISCOMEDIAGATEWAYMIB_Cmediagwipconfigtable_Cmediagwipconfigentry) GetBundleName() string { return "cisco_ios_xe" }

func (cmediagwipconfigentry *CISCOMEDIAGATEWAYMIB_Cmediagwipconfigtable_Cmediagwipconfigentry) GetYangName() string { return "cMediaGwIpConfigEntry" }

func (cmediagwipconfigentry *CISCOMEDIAGATEWAYMIB_Cmediagwipconfigtable_Cmediagwipconfigentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cmediagwipconfigentry *CISCOMEDIAGATEWAYMIB_Cmediagwipconfigtable_Cmediagwipconfigentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cmediagwipconfigentry *CISCOMEDIAGATEWAYMIB_Cmediagwipconfigtable_Cmediagwipconfigentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cmediagwipconfigentry *CISCOMEDIAGATEWAYMIB_Cmediagwipconfigtable_Cmediagwipconfigentry) SetParent(parent types.Entity) { cmediagwipconfigentry.parent = parent }

func (cmediagwipconfigentry *CISCOMEDIAGATEWAYMIB_Cmediagwipconfigtable_Cmediagwipconfigentry) GetParent() types.Entity { return cmediagwipconfigentry.parent }

func (cmediagwipconfigentry *CISCOMEDIAGATEWAYMIB_Cmediagwipconfigtable_Cmediagwipconfigentry) GetParentYangName() string { return "cMediaGwIpConfigTable" }

// CISCOMEDIAGATEWAYMIB_Cmediagwdomainnameconfigtable
// This table provides the domain names which are configured by 
// users. 
// The domain names can be used to represent IP addresses 
// for:
//     gateway
//     External DNS name server
//     MGC (call agent) 
type CISCOMEDIAGATEWAYMIB_Cmediagwdomainnameconfigtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry represents a domain name used in the system.  Creation and
    // deletion are supported. Modification is prohibited. The type is slice of
    // CISCOMEDIAGATEWAYMIB_Cmediagwdomainnameconfigtable_Cmediagwdomainnameconfigentry.
    Cmediagwdomainnameconfigentry []CISCOMEDIAGATEWAYMIB_Cmediagwdomainnameconfigtable_Cmediagwdomainnameconfigentry
}

func (cmediagwdomainnameconfigtable *CISCOMEDIAGATEWAYMIB_Cmediagwdomainnameconfigtable) GetFilter() yfilter.YFilter { return cmediagwdomainnameconfigtable.YFilter }

func (cmediagwdomainnameconfigtable *CISCOMEDIAGATEWAYMIB_Cmediagwdomainnameconfigtable) SetFilter(yf yfilter.YFilter) { cmediagwdomainnameconfigtable.YFilter = yf }

func (cmediagwdomainnameconfigtable *CISCOMEDIAGATEWAYMIB_Cmediagwdomainnameconfigtable) GetGoName(yname string) string {
    if yname == "cMediaGwDomainNameConfigEntry" { return "Cmediagwdomainnameconfigentry" }
    return ""
}

func (cmediagwdomainnameconfigtable *CISCOMEDIAGATEWAYMIB_Cmediagwdomainnameconfigtable) GetSegmentPath() string {
    return "cMediaGwDomainNameConfigTable"
}

func (cmediagwdomainnameconfigtable *CISCOMEDIAGATEWAYMIB_Cmediagwdomainnameconfigtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cMediaGwDomainNameConfigEntry" {
        for _, c := range cmediagwdomainnameconfigtable.Cmediagwdomainnameconfigentry {
            if cmediagwdomainnameconfigtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOMEDIAGATEWAYMIB_Cmediagwdomainnameconfigtable_Cmediagwdomainnameconfigentry{}
        cmediagwdomainnameconfigtable.Cmediagwdomainnameconfigentry = append(cmediagwdomainnameconfigtable.Cmediagwdomainnameconfigentry, child)
        return &cmediagwdomainnameconfigtable.Cmediagwdomainnameconfigentry[len(cmediagwdomainnameconfigtable.Cmediagwdomainnameconfigentry)-1]
    }
    return nil
}

func (cmediagwdomainnameconfigtable *CISCOMEDIAGATEWAYMIB_Cmediagwdomainnameconfigtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cmediagwdomainnameconfigtable.Cmediagwdomainnameconfigentry {
        children[cmediagwdomainnameconfigtable.Cmediagwdomainnameconfigentry[i].GetSegmentPath()] = &cmediagwdomainnameconfigtable.Cmediagwdomainnameconfigentry[i]
    }
    return children
}

func (cmediagwdomainnameconfigtable *CISCOMEDIAGATEWAYMIB_Cmediagwdomainnameconfigtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cmediagwdomainnameconfigtable *CISCOMEDIAGATEWAYMIB_Cmediagwdomainnameconfigtable) GetBundleName() string { return "cisco_ios_xe" }

func (cmediagwdomainnameconfigtable *CISCOMEDIAGATEWAYMIB_Cmediagwdomainnameconfigtable) GetYangName() string { return "cMediaGwDomainNameConfigTable" }

func (cmediagwdomainnameconfigtable *CISCOMEDIAGATEWAYMIB_Cmediagwdomainnameconfigtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cmediagwdomainnameconfigtable *CISCOMEDIAGATEWAYMIB_Cmediagwdomainnameconfigtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cmediagwdomainnameconfigtable *CISCOMEDIAGATEWAYMIB_Cmediagwdomainnameconfigtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cmediagwdomainnameconfigtable *CISCOMEDIAGATEWAYMIB_Cmediagwdomainnameconfigtable) SetParent(parent types.Entity) { cmediagwdomainnameconfigtable.parent = parent }

func (cmediagwdomainnameconfigtable *CISCOMEDIAGATEWAYMIB_Cmediagwdomainnameconfigtable) GetParent() types.Entity { return cmediagwdomainnameconfigtable.parent }

func (cmediagwdomainnameconfigtable *CISCOMEDIAGATEWAYMIB_Cmediagwdomainnameconfigtable) GetParentYangName() string { return "CISCO-MEDIA-GATEWAY-MIB" }

// CISCOMEDIAGATEWAYMIB_Cmediagwdomainnameconfigtable_Cmediagwdomainnameconfigentry
// Each entry represents a domain name used in the system.
// 
// Creation and deletion are supported. Modification
// is prohibited.
type CISCOMEDIAGATEWAYMIB_Cmediagwdomainnameconfigtable_Cmediagwdomainnameconfigentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // cisco_media_gateway_mib.CISCOMEDIAGATEWAYMIB_Cmediagwtable_Cmediagwentry_Cmgwindex
    Cmgwindex interface{}

    // This attribute is a key. An index that is uniquely identifies a domain name
    // configured in the system. The type is interface{} with range: 1..128.
    Cmgwconfigdomainnameindex interface{}

    // This object indicates which entity to use this domain name.  gateway(1)   -
    // The domain name of media gateway.                With the same cmgwIndex,
    // there is one and                 only one entry allowed with the value     
    // 'gateway(1)' of this object.  dnsServer(2) - The domain name of DNS name
    // server that is used                 by Media gateway to find Internet
    // Network                 Address from a DNS name.  mgc(3)       - The domain
    // name of a MGC (Media Gateway                Controller) associated with the
    // media                 gateway. . The type is Cmgwconfigdomainnameentity.
    Cmgwconfigdomainnameentity interface{}

    // This object specifies the domain name.  The domain name should be unique if
    // there are more than one entries having the same value in the object 
    // cmgwConfigDomainNameEntity. For example, the gateway domain name should be
    // unique  if the cmgwConfigDomainNameEntity has the value of  'gateway(1)'.
    // The type is string with length: 1..64.
    Cmgwconfigdomainname interface{}

    // This object is used to add and delete an entry.  When an entry is created,
    // the following objects are mandatory:      cmgwConfigDomainName     
    // cmgwConfigDomainNameEntity  When deleting domain name of DNS name server
    // (cmgwConfigDomainNameEntity is dnsServer (2)), the 
    // cMediaGwDnsIpConfigTable should be empty.  Adding/deleting entry with
    // cmgwConfigDomainNameEntity of 'mgc' will cause adding/deleting entry in 
    // cMgcConfigTable (CISCO-MGC-MIB) automatically.  The cmgwConfigDomainName
    // and cmgwConfigDomainNameEntity can not be modified if the value of this
    // object is 'active'. . The type is RowStatus.
    Cmgwconfigdomainnamerowstatus interface{}
}

func (cmediagwdomainnameconfigentry *CISCOMEDIAGATEWAYMIB_Cmediagwdomainnameconfigtable_Cmediagwdomainnameconfigentry) GetFilter() yfilter.YFilter { return cmediagwdomainnameconfigentry.YFilter }

func (cmediagwdomainnameconfigentry *CISCOMEDIAGATEWAYMIB_Cmediagwdomainnameconfigtable_Cmediagwdomainnameconfigentry) SetFilter(yf yfilter.YFilter) { cmediagwdomainnameconfigentry.YFilter = yf }

func (cmediagwdomainnameconfigentry *CISCOMEDIAGATEWAYMIB_Cmediagwdomainnameconfigtable_Cmediagwdomainnameconfigentry) GetGoName(yname string) string {
    if yname == "cmgwIndex" { return "Cmgwindex" }
    if yname == "cmgwConfigDomainNameIndex" { return "Cmgwconfigdomainnameindex" }
    if yname == "cmgwConfigDomainNameEntity" { return "Cmgwconfigdomainnameentity" }
    if yname == "cmgwConfigDomainName" { return "Cmgwconfigdomainname" }
    if yname == "cmgwConfigDomainNameRowStatus" { return "Cmgwconfigdomainnamerowstatus" }
    return ""
}

func (cmediagwdomainnameconfigentry *CISCOMEDIAGATEWAYMIB_Cmediagwdomainnameconfigtable_Cmediagwdomainnameconfigentry) GetSegmentPath() string {
    return "cMediaGwDomainNameConfigEntry" + "[cmgwIndex='" + fmt.Sprintf("%v", cmediagwdomainnameconfigentry.Cmgwindex) + "']" + "[cmgwConfigDomainNameIndex='" + fmt.Sprintf("%v", cmediagwdomainnameconfigentry.Cmgwconfigdomainnameindex) + "']"
}

func (cmediagwdomainnameconfigentry *CISCOMEDIAGATEWAYMIB_Cmediagwdomainnameconfigtable_Cmediagwdomainnameconfigentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cmediagwdomainnameconfigentry *CISCOMEDIAGATEWAYMIB_Cmediagwdomainnameconfigtable_Cmediagwdomainnameconfigentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cmediagwdomainnameconfigentry *CISCOMEDIAGATEWAYMIB_Cmediagwdomainnameconfigtable_Cmediagwdomainnameconfigentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cmgwIndex"] = cmediagwdomainnameconfigentry.Cmgwindex
    leafs["cmgwConfigDomainNameIndex"] = cmediagwdomainnameconfigentry.Cmgwconfigdomainnameindex
    leafs["cmgwConfigDomainNameEntity"] = cmediagwdomainnameconfigentry.Cmgwconfigdomainnameentity
    leafs["cmgwConfigDomainName"] = cmediagwdomainnameconfigentry.Cmgwconfigdomainname
    leafs["cmgwConfigDomainNameRowStatus"] = cmediagwdomainnameconfigentry.Cmgwconfigdomainnamerowstatus
    return leafs
}

func (cmediagwdomainnameconfigentry *CISCOMEDIAGATEWAYMIB_Cmediagwdomainnameconfigtable_Cmediagwdomainnameconfigentry) GetBundleName() string { return "cisco_ios_xe" }

func (cmediagwdomainnameconfigentry *CISCOMEDIAGATEWAYMIB_Cmediagwdomainnameconfigtable_Cmediagwdomainnameconfigentry) GetYangName() string { return "cMediaGwDomainNameConfigEntry" }

func (cmediagwdomainnameconfigentry *CISCOMEDIAGATEWAYMIB_Cmediagwdomainnameconfigtable_Cmediagwdomainnameconfigentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cmediagwdomainnameconfigentry *CISCOMEDIAGATEWAYMIB_Cmediagwdomainnameconfigtable_Cmediagwdomainnameconfigentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cmediagwdomainnameconfigentry *CISCOMEDIAGATEWAYMIB_Cmediagwdomainnameconfigtable_Cmediagwdomainnameconfigentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cmediagwdomainnameconfigentry *CISCOMEDIAGATEWAYMIB_Cmediagwdomainnameconfigtable_Cmediagwdomainnameconfigentry) SetParent(parent types.Entity) { cmediagwdomainnameconfigentry.parent = parent }

func (cmediagwdomainnameconfigentry *CISCOMEDIAGATEWAYMIB_Cmediagwdomainnameconfigtable_Cmediagwdomainnameconfigentry) GetParent() types.Entity { return cmediagwdomainnameconfigentry.parent }

func (cmediagwdomainnameconfigentry *CISCOMEDIAGATEWAYMIB_Cmediagwdomainnameconfigtable_Cmediagwdomainnameconfigentry) GetParentYangName() string { return "cMediaGwDomainNameConfigTable" }

// CISCOMEDIAGATEWAYMIB_Cmediagwdomainnameconfigtable_Cmediagwdomainnameconfigentry_Cmgwconfigdomainnameentity represents                gateway. 
type CISCOMEDIAGATEWAYMIB_Cmediagwdomainnameconfigtable_Cmediagwdomainnameconfigentry_Cmgwconfigdomainnameentity string

const (
    CISCOMEDIAGATEWAYMIB_Cmediagwdomainnameconfigtable_Cmediagwdomainnameconfigentry_Cmgwconfigdomainnameentity_gateway CISCOMEDIAGATEWAYMIB_Cmediagwdomainnameconfigtable_Cmediagwdomainnameconfigentry_Cmgwconfigdomainnameentity = "gateway"

    CISCOMEDIAGATEWAYMIB_Cmediagwdomainnameconfigtable_Cmediagwdomainnameconfigentry_Cmgwconfigdomainnameentity_dnsServer CISCOMEDIAGATEWAYMIB_Cmediagwdomainnameconfigtable_Cmediagwdomainnameconfigentry_Cmgwconfigdomainnameentity = "dnsServer"

    CISCOMEDIAGATEWAYMIB_Cmediagwdomainnameconfigtable_Cmediagwdomainnameconfigentry_Cmgwconfigdomainnameentity_mgc CISCOMEDIAGATEWAYMIB_Cmediagwdomainnameconfigtable_Cmediagwdomainnameconfigentry_Cmgwconfigdomainnameentity = "mgc"
)

// CISCOMEDIAGATEWAYMIB_Cmediagwdnsipconfigtable
// There is only one DNS name server on a gateway
// and the domain name of the DNS name server is put on 
// cMediaGwDomainNameConfigTable with 'dnsServer (2)'.
// 
// There could be multi IP addresses are associated with the
// DNS name server, this table is used to store these IP 
// addresses.
// 
// If any domain name using external resolution, the last entry
// of this table is not allowed to be deleted.
type CISCOMEDIAGATEWAYMIB_Cmediagwdnsipconfigtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry represents an IP address of the DNS name  server. The type is
    // slice of
    // CISCOMEDIAGATEWAYMIB_Cmediagwdnsipconfigtable_Cmediagwdnsipconfigentry.
    Cmediagwdnsipconfigentry []CISCOMEDIAGATEWAYMIB_Cmediagwdnsipconfigtable_Cmediagwdnsipconfigentry
}

func (cmediagwdnsipconfigtable *CISCOMEDIAGATEWAYMIB_Cmediagwdnsipconfigtable) GetFilter() yfilter.YFilter { return cmediagwdnsipconfigtable.YFilter }

func (cmediagwdnsipconfigtable *CISCOMEDIAGATEWAYMIB_Cmediagwdnsipconfigtable) SetFilter(yf yfilter.YFilter) { cmediagwdnsipconfigtable.YFilter = yf }

func (cmediagwdnsipconfigtable *CISCOMEDIAGATEWAYMIB_Cmediagwdnsipconfigtable) GetGoName(yname string) string {
    if yname == "cMediaGwDnsIpConfigEntry" { return "Cmediagwdnsipconfigentry" }
    return ""
}

func (cmediagwdnsipconfigtable *CISCOMEDIAGATEWAYMIB_Cmediagwdnsipconfigtable) GetSegmentPath() string {
    return "cMediaGwDnsIpConfigTable"
}

func (cmediagwdnsipconfigtable *CISCOMEDIAGATEWAYMIB_Cmediagwdnsipconfigtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cMediaGwDnsIpConfigEntry" {
        for _, c := range cmediagwdnsipconfigtable.Cmediagwdnsipconfigentry {
            if cmediagwdnsipconfigtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOMEDIAGATEWAYMIB_Cmediagwdnsipconfigtable_Cmediagwdnsipconfigentry{}
        cmediagwdnsipconfigtable.Cmediagwdnsipconfigentry = append(cmediagwdnsipconfigtable.Cmediagwdnsipconfigentry, child)
        return &cmediagwdnsipconfigtable.Cmediagwdnsipconfigentry[len(cmediagwdnsipconfigtable.Cmediagwdnsipconfigentry)-1]
    }
    return nil
}

func (cmediagwdnsipconfigtable *CISCOMEDIAGATEWAYMIB_Cmediagwdnsipconfigtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cmediagwdnsipconfigtable.Cmediagwdnsipconfigentry {
        children[cmediagwdnsipconfigtable.Cmediagwdnsipconfigentry[i].GetSegmentPath()] = &cmediagwdnsipconfigtable.Cmediagwdnsipconfigentry[i]
    }
    return children
}

func (cmediagwdnsipconfigtable *CISCOMEDIAGATEWAYMIB_Cmediagwdnsipconfigtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cmediagwdnsipconfigtable *CISCOMEDIAGATEWAYMIB_Cmediagwdnsipconfigtable) GetBundleName() string { return "cisco_ios_xe" }

func (cmediagwdnsipconfigtable *CISCOMEDIAGATEWAYMIB_Cmediagwdnsipconfigtable) GetYangName() string { return "cMediaGwDnsIpConfigTable" }

func (cmediagwdnsipconfigtable *CISCOMEDIAGATEWAYMIB_Cmediagwdnsipconfigtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cmediagwdnsipconfigtable *CISCOMEDIAGATEWAYMIB_Cmediagwdnsipconfigtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cmediagwdnsipconfigtable *CISCOMEDIAGATEWAYMIB_Cmediagwdnsipconfigtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cmediagwdnsipconfigtable *CISCOMEDIAGATEWAYMIB_Cmediagwdnsipconfigtable) SetParent(parent types.Entity) { cmediagwdnsipconfigtable.parent = parent }

func (cmediagwdnsipconfigtable *CISCOMEDIAGATEWAYMIB_Cmediagwdnsipconfigtable) GetParent() types.Entity { return cmediagwdnsipconfigtable.parent }

func (cmediagwdnsipconfigtable *CISCOMEDIAGATEWAYMIB_Cmediagwdnsipconfigtable) GetParentYangName() string { return "CISCO-MEDIA-GATEWAY-MIB" }

// CISCOMEDIAGATEWAYMIB_Cmediagwdnsipconfigtable_Cmediagwdnsipconfigentry
// Each entry represents an IP address of the DNS name 
// server.
type CISCOMEDIAGATEWAYMIB_Cmediagwdnsipconfigtable_Cmediagwdnsipconfigentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // cisco_media_gateway_mib.CISCOMEDIAGATEWAYMIB_Cmediagwtable_Cmediagwentry_Cmgwindex
    Cmgwindex interface{}

    // This attribute is a key. An index that uniquely identifies an IP address of
    // DNS name server. The type is interface{} with range: 1..6.
    Cmgwdnsipindex interface{}

    // The domain name of DNS name server.  The value of this object reflects the
    // value of  cmgwConfigDomainName from the entry with a value of 
    // 'dnsServer(2)' for object cmgwConfigDomainNameEntity of 
    // cMediaGwDomainNameConfigTable.  If there is no entry in
    // cMediaGwDomainNameConfigTable with 'dnsServer(2)' of
    // cmgwConfigDomainNameEntity, then the value of this object will be empty
    // string. The type is string with length: 1..64.
    Cmgwdnsdomainname interface{}

    // DNS name server IP address type. The type is InetAddressType.
    Cmgwdnsiptype interface{}

    // The IP address of DNS name server. The IP address of DNS name server must
    // be unique in this table. The type is string with length: 0..255.
    Cmgwdnsip interface{}

    // This object is used to add and delete an entry.  When an entry of the table
    // is created, the value of this object should be set to 'createAndGo' and the
    // following objects are mandatory:     cmgwDnsIp  When the user wants to
    // delete the entry, the value of this object should be set to 'destroy'.  The
    // entry can not be modified if the value of this  object is 'active'. The
    // type is RowStatus.
    Cmgwdnsiprowstatus interface{}
}

func (cmediagwdnsipconfigentry *CISCOMEDIAGATEWAYMIB_Cmediagwdnsipconfigtable_Cmediagwdnsipconfigentry) GetFilter() yfilter.YFilter { return cmediagwdnsipconfigentry.YFilter }

func (cmediagwdnsipconfigentry *CISCOMEDIAGATEWAYMIB_Cmediagwdnsipconfigtable_Cmediagwdnsipconfigentry) SetFilter(yf yfilter.YFilter) { cmediagwdnsipconfigentry.YFilter = yf }

func (cmediagwdnsipconfigentry *CISCOMEDIAGATEWAYMIB_Cmediagwdnsipconfigtable_Cmediagwdnsipconfigentry) GetGoName(yname string) string {
    if yname == "cmgwIndex" { return "Cmgwindex" }
    if yname == "cmgwDnsIpIndex" { return "Cmgwdnsipindex" }
    if yname == "cmgwDnsDomainName" { return "Cmgwdnsdomainname" }
    if yname == "cmgwDnsIpType" { return "Cmgwdnsiptype" }
    if yname == "cmgwDnsIp" { return "Cmgwdnsip" }
    if yname == "cmgwDnsIpRowStatus" { return "Cmgwdnsiprowstatus" }
    return ""
}

func (cmediagwdnsipconfigentry *CISCOMEDIAGATEWAYMIB_Cmediagwdnsipconfigtable_Cmediagwdnsipconfigentry) GetSegmentPath() string {
    return "cMediaGwDnsIpConfigEntry" + "[cmgwIndex='" + fmt.Sprintf("%v", cmediagwdnsipconfigentry.Cmgwindex) + "']" + "[cmgwDnsIpIndex='" + fmt.Sprintf("%v", cmediagwdnsipconfigentry.Cmgwdnsipindex) + "']"
}

func (cmediagwdnsipconfigentry *CISCOMEDIAGATEWAYMIB_Cmediagwdnsipconfigtable_Cmediagwdnsipconfigentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cmediagwdnsipconfigentry *CISCOMEDIAGATEWAYMIB_Cmediagwdnsipconfigtable_Cmediagwdnsipconfigentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cmediagwdnsipconfigentry *CISCOMEDIAGATEWAYMIB_Cmediagwdnsipconfigtable_Cmediagwdnsipconfigentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cmgwIndex"] = cmediagwdnsipconfigentry.Cmgwindex
    leafs["cmgwDnsIpIndex"] = cmediagwdnsipconfigentry.Cmgwdnsipindex
    leafs["cmgwDnsDomainName"] = cmediagwdnsipconfigentry.Cmgwdnsdomainname
    leafs["cmgwDnsIpType"] = cmediagwdnsipconfigentry.Cmgwdnsiptype
    leafs["cmgwDnsIp"] = cmediagwdnsipconfigentry.Cmgwdnsip
    leafs["cmgwDnsIpRowStatus"] = cmediagwdnsipconfigentry.Cmgwdnsiprowstatus
    return leafs
}

func (cmediagwdnsipconfigentry *CISCOMEDIAGATEWAYMIB_Cmediagwdnsipconfigtable_Cmediagwdnsipconfigentry) GetBundleName() string { return "cisco_ios_xe" }

func (cmediagwdnsipconfigentry *CISCOMEDIAGATEWAYMIB_Cmediagwdnsipconfigtable_Cmediagwdnsipconfigentry) GetYangName() string { return "cMediaGwDnsIpConfigEntry" }

func (cmediagwdnsipconfigentry *CISCOMEDIAGATEWAYMIB_Cmediagwdnsipconfigtable_Cmediagwdnsipconfigentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cmediagwdnsipconfigentry *CISCOMEDIAGATEWAYMIB_Cmediagwdnsipconfigtable_Cmediagwdnsipconfigentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cmediagwdnsipconfigentry *CISCOMEDIAGATEWAYMIB_Cmediagwdnsipconfigtable_Cmediagwdnsipconfigentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cmediagwdnsipconfigentry *CISCOMEDIAGATEWAYMIB_Cmediagwdnsipconfigtable_Cmediagwdnsipconfigentry) SetParent(parent types.Entity) { cmediagwdnsipconfigentry.parent = parent }

func (cmediagwdnsipconfigentry *CISCOMEDIAGATEWAYMIB_Cmediagwdnsipconfigtable_Cmediagwdnsipconfigentry) GetParent() types.Entity { return cmediagwdnsipconfigentry.parent }

func (cmediagwdnsipconfigentry *CISCOMEDIAGATEWAYMIB_Cmediagwdnsipconfigtable_Cmediagwdnsipconfigentry) GetParentYangName() string { return "cMediaGwDnsIpConfigTable" }

// CISCOMEDIAGATEWAYMIB_Cmgwliftable
// This table is for managing LIF (Logical Interface) 
// in a media gateway. 
// 
// LIF is a logical interface which groups the TDM 
// DSx1s associated with a set of packet resource partitions 
// (PVCs) in a media gateway.
// 
// LIF is used for:
// 1. VoIP switching 
// 2. VoATM switching 
type CISCOMEDIAGATEWAYMIB_Cmgwliftable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry of this table is created by the media gateway when it supports the
    // VoIP/VoATM application. The type is slice of
    // CISCOMEDIAGATEWAYMIB_Cmgwliftable_Cmgwlifentry.
    Cmgwlifentry []CISCOMEDIAGATEWAYMIB_Cmgwliftable_Cmgwlifentry
}

func (cmgwliftable *CISCOMEDIAGATEWAYMIB_Cmgwliftable) GetFilter() yfilter.YFilter { return cmgwliftable.YFilter }

func (cmgwliftable *CISCOMEDIAGATEWAYMIB_Cmgwliftable) SetFilter(yf yfilter.YFilter) { cmgwliftable.YFilter = yf }

func (cmgwliftable *CISCOMEDIAGATEWAYMIB_Cmgwliftable) GetGoName(yname string) string {
    if yname == "cmgwLifEntry" { return "Cmgwlifentry" }
    return ""
}

func (cmgwliftable *CISCOMEDIAGATEWAYMIB_Cmgwliftable) GetSegmentPath() string {
    return "cmgwLifTable"
}

func (cmgwliftable *CISCOMEDIAGATEWAYMIB_Cmgwliftable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cmgwLifEntry" {
        for _, c := range cmgwliftable.Cmgwlifentry {
            if cmgwliftable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOMEDIAGATEWAYMIB_Cmgwliftable_Cmgwlifentry{}
        cmgwliftable.Cmgwlifentry = append(cmgwliftable.Cmgwlifentry, child)
        return &cmgwliftable.Cmgwlifentry[len(cmgwliftable.Cmgwlifentry)-1]
    }
    return nil
}

func (cmgwliftable *CISCOMEDIAGATEWAYMIB_Cmgwliftable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cmgwliftable.Cmgwlifentry {
        children[cmgwliftable.Cmgwlifentry[i].GetSegmentPath()] = &cmgwliftable.Cmgwlifentry[i]
    }
    return children
}

func (cmgwliftable *CISCOMEDIAGATEWAYMIB_Cmgwliftable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cmgwliftable *CISCOMEDIAGATEWAYMIB_Cmgwliftable) GetBundleName() string { return "cisco_ios_xe" }

func (cmgwliftable *CISCOMEDIAGATEWAYMIB_Cmgwliftable) GetYangName() string { return "cmgwLifTable" }

func (cmgwliftable *CISCOMEDIAGATEWAYMIB_Cmgwliftable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cmgwliftable *CISCOMEDIAGATEWAYMIB_Cmgwliftable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cmgwliftable *CISCOMEDIAGATEWAYMIB_Cmgwliftable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cmgwliftable *CISCOMEDIAGATEWAYMIB_Cmgwliftable) SetParent(parent types.Entity) { cmgwliftable.parent = parent }

func (cmgwliftable *CISCOMEDIAGATEWAYMIB_Cmgwliftable) GetParent() types.Entity { return cmgwliftable.parent }

func (cmgwliftable *CISCOMEDIAGATEWAYMIB_Cmgwliftable) GetParentYangName() string { return "CISCO-MEDIA-GATEWAY-MIB" }

// CISCOMEDIAGATEWAYMIB_Cmgwliftable_Cmgwlifentry
// An entry of this table is created by the media gateway
// when it supports the VoIP/VoATM application.
type CISCOMEDIAGATEWAYMIB_Cmgwliftable_Cmgwlifentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // cisco_media_gateway_mib.CISCOMEDIAGATEWAYMIB_Cmediagwtable_Cmediagwentry_Cmgwindex
    Cmgwindex interface{}

    // This attribute is a key. An index that uniquely identifies a LIF in the 
    // media gateway. The type is interface{} with range: 1..255.
    Cmgwlifnumber interface{}

    // This object represents the total number of PVC within  this LIF.  When
    // users associate/disassociate a PVC with a LIF  by giving a non-zero/zero
    // value of cwacChanLifNum in cwAtmChanExtConfigTable, the value of this
    // object  will be incremented/decremented accordingly.  The value zero means
    // there is no PVC associated with  the LIF. The type is interface{} with
    // range: 0..10000.
    Cmgwlifpvccount interface{}

    // This object represents the total number of Voice Interfaces within this
    // LIF.  When users associate/disassociate a Voice Interface with a LIF by
    // giving a non-zero/zero value of  ccasVoiceCfgLifNumber for the DS0 group in
    // ccasVoiceExtCfgTable, the value of this object will be 
    // incremented/decremented accordingly.   The value zero means there is no
    // Voice Interface associated with the LIF. The type is interface{} with
    // range: 0..1000.
    Cmgwlifvoiceifcount interface{}
}

func (cmgwlifentry *CISCOMEDIAGATEWAYMIB_Cmgwliftable_Cmgwlifentry) GetFilter() yfilter.YFilter { return cmgwlifentry.YFilter }

func (cmgwlifentry *CISCOMEDIAGATEWAYMIB_Cmgwliftable_Cmgwlifentry) SetFilter(yf yfilter.YFilter) { cmgwlifentry.YFilter = yf }

func (cmgwlifentry *CISCOMEDIAGATEWAYMIB_Cmgwliftable_Cmgwlifentry) GetGoName(yname string) string {
    if yname == "cmgwIndex" { return "Cmgwindex" }
    if yname == "cmgwLifNumber" { return "Cmgwlifnumber" }
    if yname == "cmgwLifPvcCount" { return "Cmgwlifpvccount" }
    if yname == "cmgwLifVoiceIfCount" { return "Cmgwlifvoiceifcount" }
    return ""
}

func (cmgwlifentry *CISCOMEDIAGATEWAYMIB_Cmgwliftable_Cmgwlifentry) GetSegmentPath() string {
    return "cmgwLifEntry" + "[cmgwIndex='" + fmt.Sprintf("%v", cmgwlifentry.Cmgwindex) + "']" + "[cmgwLifNumber='" + fmt.Sprintf("%v", cmgwlifentry.Cmgwlifnumber) + "']"
}

func (cmgwlifentry *CISCOMEDIAGATEWAYMIB_Cmgwliftable_Cmgwlifentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cmgwlifentry *CISCOMEDIAGATEWAYMIB_Cmgwliftable_Cmgwlifentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cmgwlifentry *CISCOMEDIAGATEWAYMIB_Cmgwliftable_Cmgwlifentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cmgwIndex"] = cmgwlifentry.Cmgwindex
    leafs["cmgwLifNumber"] = cmgwlifentry.Cmgwlifnumber
    leafs["cmgwLifPvcCount"] = cmgwlifentry.Cmgwlifpvccount
    leafs["cmgwLifVoiceIfCount"] = cmgwlifentry.Cmgwlifvoiceifcount
    return leafs
}

func (cmgwlifentry *CISCOMEDIAGATEWAYMIB_Cmgwliftable_Cmgwlifentry) GetBundleName() string { return "cisco_ios_xe" }

func (cmgwlifentry *CISCOMEDIAGATEWAYMIB_Cmgwliftable_Cmgwlifentry) GetYangName() string { return "cmgwLifEntry" }

func (cmgwlifentry *CISCOMEDIAGATEWAYMIB_Cmgwliftable_Cmgwlifentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cmgwlifentry *CISCOMEDIAGATEWAYMIB_Cmgwliftable_Cmgwlifentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cmgwlifentry *CISCOMEDIAGATEWAYMIB_Cmgwliftable_Cmgwlifentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cmgwlifentry *CISCOMEDIAGATEWAYMIB_Cmgwliftable_Cmgwlifentry) SetParent(parent types.Entity) { cmgwlifentry.parent = parent }

func (cmgwlifentry *CISCOMEDIAGATEWAYMIB_Cmgwliftable_Cmgwlifentry) GetParent() types.Entity { return cmgwlifentry.parent }

func (cmgwlifentry *CISCOMEDIAGATEWAYMIB_Cmgwliftable_Cmgwlifentry) GetParentYangName() string { return "cmgwLifTable" }

// CISCOMEDIAGATEWAYMIB_Cmediagwcallcontrolconfigtable
// This table defines general call control attributes for
// the media gateway.
type CISCOMEDIAGATEWAYMIB_Cmediagwcallcontrolconfigtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // One entry for each media gateway which supports call control  protocol. The
    // type is slice of
    // CISCOMEDIAGATEWAYMIB_Cmediagwcallcontrolconfigtable_Cmediagwcallcontrolconfigentry.
    Cmediagwcallcontrolconfigentry []CISCOMEDIAGATEWAYMIB_Cmediagwcallcontrolconfigtable_Cmediagwcallcontrolconfigentry
}

func (cmediagwcallcontrolconfigtable *CISCOMEDIAGATEWAYMIB_Cmediagwcallcontrolconfigtable) GetFilter() yfilter.YFilter { return cmediagwcallcontrolconfigtable.YFilter }

func (cmediagwcallcontrolconfigtable *CISCOMEDIAGATEWAYMIB_Cmediagwcallcontrolconfigtable) SetFilter(yf yfilter.YFilter) { cmediagwcallcontrolconfigtable.YFilter = yf }

func (cmediagwcallcontrolconfigtable *CISCOMEDIAGATEWAYMIB_Cmediagwcallcontrolconfigtable) GetGoName(yname string) string {
    if yname == "cMediaGwCallControlConfigEntry" { return "Cmediagwcallcontrolconfigentry" }
    return ""
}

func (cmediagwcallcontrolconfigtable *CISCOMEDIAGATEWAYMIB_Cmediagwcallcontrolconfigtable) GetSegmentPath() string {
    return "cMediaGwCallControlConfigTable"
}

func (cmediagwcallcontrolconfigtable *CISCOMEDIAGATEWAYMIB_Cmediagwcallcontrolconfigtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cMediaGwCallControlConfigEntry" {
        for _, c := range cmediagwcallcontrolconfigtable.Cmediagwcallcontrolconfigentry {
            if cmediagwcallcontrolconfigtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOMEDIAGATEWAYMIB_Cmediagwcallcontrolconfigtable_Cmediagwcallcontrolconfigentry{}
        cmediagwcallcontrolconfigtable.Cmediagwcallcontrolconfigentry = append(cmediagwcallcontrolconfigtable.Cmediagwcallcontrolconfigentry, child)
        return &cmediagwcallcontrolconfigtable.Cmediagwcallcontrolconfigentry[len(cmediagwcallcontrolconfigtable.Cmediagwcallcontrolconfigentry)-1]
    }
    return nil
}

func (cmediagwcallcontrolconfigtable *CISCOMEDIAGATEWAYMIB_Cmediagwcallcontrolconfigtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cmediagwcallcontrolconfigtable.Cmediagwcallcontrolconfigentry {
        children[cmediagwcallcontrolconfigtable.Cmediagwcallcontrolconfigentry[i].GetSegmentPath()] = &cmediagwcallcontrolconfigtable.Cmediagwcallcontrolconfigentry[i]
    }
    return children
}

func (cmediagwcallcontrolconfigtable *CISCOMEDIAGATEWAYMIB_Cmediagwcallcontrolconfigtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cmediagwcallcontrolconfigtable *CISCOMEDIAGATEWAYMIB_Cmediagwcallcontrolconfigtable) GetBundleName() string { return "cisco_ios_xe" }

func (cmediagwcallcontrolconfigtable *CISCOMEDIAGATEWAYMIB_Cmediagwcallcontrolconfigtable) GetYangName() string { return "cMediaGwCallControlConfigTable" }

func (cmediagwcallcontrolconfigtable *CISCOMEDIAGATEWAYMIB_Cmediagwcallcontrolconfigtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cmediagwcallcontrolconfigtable *CISCOMEDIAGATEWAYMIB_Cmediagwcallcontrolconfigtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cmediagwcallcontrolconfigtable *CISCOMEDIAGATEWAYMIB_Cmediagwcallcontrolconfigtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cmediagwcallcontrolconfigtable *CISCOMEDIAGATEWAYMIB_Cmediagwcallcontrolconfigtable) SetParent(parent types.Entity) { cmediagwcallcontrolconfigtable.parent = parent }

func (cmediagwcallcontrolconfigtable *CISCOMEDIAGATEWAYMIB_Cmediagwcallcontrolconfigtable) GetParent() types.Entity { return cmediagwcallcontrolconfigtable.parent }

func (cmediagwcallcontrolconfigtable *CISCOMEDIAGATEWAYMIB_Cmediagwcallcontrolconfigtable) GetParentYangName() string { return "CISCO-MEDIA-GATEWAY-MIB" }

// CISCOMEDIAGATEWAYMIB_Cmediagwcallcontrolconfigtable_Cmediagwcallcontrolconfigentry
// One entry for each media gateway which supports call control 
// protocol.
type CISCOMEDIAGATEWAYMIB_Cmediagwcallcontrolconfigtable_Cmediagwcallcontrolconfigentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // cisco_media_gateway_mib.CISCOMEDIAGATEWAYMIB_Cmediagwtable_Cmediagwentry_Cmgwindex
    Cmgwindex interface{}

    // This object specifies Type Of Service (TOS) field of IP header for the
    // signaling control packet in VoIP application. The type is interface{} with
    // range: 0..255.
    Cmediagwcccfgcontroltos interface{}

    // This object specifies Type Of Service (TOS) field of IP header for the
    // voice payload packet in VoIP application. The type is interface{} with
    // range: 0..255.
    Cmediagwcccfgbearertos interface{}

    // This object specifies NTE (Named Telephony Events) payload type. The type
    // is interface{} with range: 96..127.
    Cmediagwcccfgntepayload interface{}

    // This object specifies NSE (Network Signaling Events) payload type. The type
    // is interface{} with range: 98..117.
    Cmediagwcccfgnsepayload interface{}

    // This object specifies Network Signaling Event (NSE) timeout value. The type
    // is interface{} with range: 250..10000. Units are milliseconds.
    Cmediagwcccfgnseresptimer interface{}

    // The object specifies the jitter buffer mode applied to a VBD (Voice Band
    // Data) call connection.  adaptive - means to use
    // cMediaGwCcCfgVbdJitterNomDelay as            the initial jitter buffers
    // size and let the DSP            pick the optimal value of the jitter buffer
    // size between the range of            cMediaGwCcCfgVbcJitterMaxDelay and    
    // cMediaGwCcCfgVbcJitterMinDelay.  fixed - means to use a constant jitter
    // buffer size         which is specified by cMediaGwCcCfgVbdJitterNomDelay.
    // The type is CCallControlJitterDelayMode.
    Cmediagwcccfgvbdjitterdelaymode interface{}

    // This object specifies the maximum jitter buffer size  in VBD (Voice Band
    // Data). The type is interface{} with range: 20..135. Units are milliseconds.
    Cmediagwcccfgvbdjittermaxdelay interface{}

    // This object specifies the nominal jitter buffer size  in VBD (Voice Band
    // Data). The type is interface{} with range: 5..135. Units are milliseconds.
    Cmediagwcccfgvbdjitternomdelay interface{}

    // This object specifies the nominal jitter buffer size  in VBD (Voice Band
    // Data). The type is interface{} with range: 0..135. Units are milliseconds.
    Cmediagwcccfgvbdjittermindelay interface{}

    // This object specifies the default tone plan index (the value of
    // cvtcTonePlanId) for the media gateway. The type is interface{} with range:
    // 1..65535.
    Cmediagwcccfgdefaulttoneplanid interface{}

    // This object specifies whether the media gateway supports descriptive suffix
    // of the name schema for terminations.  There are two parts in name schema of
    // termination, prefix and suffix. For example the name schema for a DS
    // (Digital Subscriber) termination, can be 'DS/OC3_2/DS1_6/DS0_24'. It
    // represents DS type termination in 2nd OC3 line,  6th DS1 and 24th DS0
    // channel. In this example, 'DS' is  the prefix, 'OC3_2/DS1_6/DS0_24' is the
    // suffix.  The name schema in above example has a descriptive suffix. The
    // non-descriptive suffix for the same termination is  '2/6/24' and name
    // schema becomes 'DS/2/6/24'.  This object can not be modified if there is
    // any termination existing in the media gateway. The type is bool.
    Cmediagwcccfgdescrinfoenabled interface{}

    // This object specifies the prefix of the name schema for DS (Digital
    // Subscriber) terminations. The value of this object must be unique among the
    // following objects:        cMediaGwCcCfgDsNamePrefix       
    // cMediaGwCcCfgRtpNamePrefix        cMediaGwCcCfgAal2SvcNamePrefix       
    // cMediaGwCcCfgAal2SvcNamePrefix        cMediaGwCcCfgDefRtpNamePrefix This
    // object can not be modified when there is any DS termination existing in the
    // media gateway. It is default to 'DS'. The type is string with length:
    // 1..64.
    Cmediagwcccfgdsnameprefix interface{}

    // This object specifies the prefix of the name schema for RTP (Real-Time
    // Transport Protocol) terminations. The value of this object must be unique
    // among the  following objects:        cMediaGwCcCfgDsNamePrefix       
    // cMediaGwCcCfgRtpNamePrefix        cMediaGwCcCfgAal2SvcNamePrefix       
    // cMediaGwCcCfgAal2SvcNamePrefix        cMediaGwCcCfgDefRtpNamePrefix This
    // object can not be modified when there is any RTP termination type existing
    // in the media gateway. It is default to 'RTP'. The type is string with
    // length: 1..64.
    Cmediagwcccfgrtpnameprefix interface{}

    // This object specifies the prefix of the name schema for voice over AAL1 SVC
    // (Switched Virtual Circuit) terminations. The value of this object must be
    // unique among the  following objects:        cMediaGwCcCfgDsNamePrefix      
    // cMediaGwCcCfgRtpNamePrefix        cMediaGwCcCfgAal2SvcNamePrefix       
    // cMediaGwCcCfgAal2SvcNamePrefix        cMediaGwCcCfgDefRtpNamePrefix This
    // object can not be modified when there is any AAL1 SVC termination type
    // existing in the media gateway. It is default to 'AAL1/SVC'. The type is
    // string with length: 1..64.
    Cmediagwcccfgaal1Svcnameprefix interface{}

    // This object specifies the prefix of the name schema for voice over AAL2 SVC
    // (Switched Virtual Circuit) terminations. The value of this object must be
    // unique among the  following objects:        cMediaGwCcCfgDsNamePrefix      
    // cMediaGwCcCfgRtpNamePrefix        cMediaGwCcCfgAal2SvcNamePrefix       
    // cMediaGwCcCfgAal2SvcNamePrefix        cMediaGwCcCfgDefRtpNamePrefix This
    // object can not be modified when there is any AAL2 SVC termination type
    // existing in the media gateway. It is default to 'AAL2/SVC'. The type is
    // string with length: 1..64.
    Cmediagwcccfgaal2Svcnameprefix interface{}

    // This object specifies the condition of the cluster generation in the call
    // control.  A cluster is a group of endpoints that share a particular bearer
    // possibility for connections among each other.  disabled(1) - The generation
    // of the cluster attribute               is disabled. enabled(2) -
    // Unconditionally generate the cluster              attribute.
    // conditionalEnabled(3) - The generation of the cluster               
    // attribute is upon MGC request. The type is Cmediagwcccfgclusterenabled.
    Cmediagwcccfgclusterenabled interface{}

    // This object specifies the combination of the network type (IP/ATM), virtual
    // circuit type (PVC/SVC) and ATM adaptation layer type (AAL1/AAL2/AAL5) for
    // the connection used in transporting bearer traffic.      ipPvcAal5 (1) -
    // The bearer traffic is transported in                     IP network,
    // through Permanent Virtual                     Circuit(PVC) over AAL5
    // adaptation layer.     atmPvcAal2 (2) - The bearer traffic is transported in
    // ATM network, through Permanent Virtual                      Circuit(PVC)
    // over AAL2 adaptation layer.     atmSvcAal2 (3) - The bearer traffic is
    // transported in                      ATM network, through Switching Virtual 
    // Circuit(SVC) over AAL2 adaptation layer.     atmSvcAal1 (4) - The bearer
    // traffic is transported in                      ATM network, through
    // Switching Virtual                      Circuit(SVC) over AAL1 adaptation
    // layer.  In MGCP, if the call agent specifies the bear traffic type  in the
    // local connection options (CRCX request),  configuration of this object will
    // have no effect,  otherwise the value of this object will be used when 
    // media gateway sending CRCX response. The type is
    // Cmediagwcccfgdefbearertraffic.
    Cmediagwcccfgdefbearertraffic interface{}

    // This object specifies the prefix of the name schema for default RTP
    // terminations. The value of this object must be unique among the  following
    // objects:        cMediaGwCcCfgDsNamePrefix        cMediaGwCcCfgRtpNamePrefix
    // cMediaGwCcCfgAal1SvcNamePrefix        cMediaGwCcCfgAal2SvcNamePrefix  It is
    // defaulted to 'TGWRTP'. The type is string with length: 1..64.
    Cmediagwcccfgdefrtpnameprefix interface{}
}

func (cmediagwcallcontrolconfigentry *CISCOMEDIAGATEWAYMIB_Cmediagwcallcontrolconfigtable_Cmediagwcallcontrolconfigentry) GetFilter() yfilter.YFilter { return cmediagwcallcontrolconfigentry.YFilter }

func (cmediagwcallcontrolconfigentry *CISCOMEDIAGATEWAYMIB_Cmediagwcallcontrolconfigtable_Cmediagwcallcontrolconfigentry) SetFilter(yf yfilter.YFilter) { cmediagwcallcontrolconfigentry.YFilter = yf }

func (cmediagwcallcontrolconfigentry *CISCOMEDIAGATEWAYMIB_Cmediagwcallcontrolconfigtable_Cmediagwcallcontrolconfigentry) GetGoName(yname string) string {
    if yname == "cmgwIndex" { return "Cmgwindex" }
    if yname == "cMediaGwCcCfgControlTos" { return "Cmediagwcccfgcontroltos" }
    if yname == "cMediaGwCcCfgBearerTos" { return "Cmediagwcccfgbearertos" }
    if yname == "cMediaGwCcCfgNtePayload" { return "Cmediagwcccfgntepayload" }
    if yname == "cMediaGwCcCfgNsePayload" { return "Cmediagwcccfgnsepayload" }
    if yname == "cMediaGwCcCfgNseRespTimer" { return "Cmediagwcccfgnseresptimer" }
    if yname == "cMediaGwCcCfgVbdJitterDelayMode" { return "Cmediagwcccfgvbdjitterdelaymode" }
    if yname == "cMediaGwCcCfgVbdJitterMaxDelay" { return "Cmediagwcccfgvbdjittermaxdelay" }
    if yname == "cMediaGwCcCfgVbdJitterNomDelay" { return "Cmediagwcccfgvbdjitternomdelay" }
    if yname == "cMediaGwCcCfgVbdJitterMinDelay" { return "Cmediagwcccfgvbdjittermindelay" }
    if yname == "cMediaGwCcCfgDefaultTonePlanId" { return "Cmediagwcccfgdefaulttoneplanid" }
    if yname == "cMediaGwCcCfgDescrInfoEnabled" { return "Cmediagwcccfgdescrinfoenabled" }
    if yname == "cMediaGwCcCfgDsNamePrefix" { return "Cmediagwcccfgdsnameprefix" }
    if yname == "cMediaGwCcCfgRtpNamePrefix" { return "Cmediagwcccfgrtpnameprefix" }
    if yname == "cMediaGwCcCfgAal1SvcNamePrefix" { return "Cmediagwcccfgaal1Svcnameprefix" }
    if yname == "cMediaGwCcCfgAal2SvcNamePrefix" { return "Cmediagwcccfgaal2Svcnameprefix" }
    if yname == "cMediaGwCcCfgClusterEnabled" { return "Cmediagwcccfgclusterenabled" }
    if yname == "cMediaGwCcCfgDefBearerTraffic" { return "Cmediagwcccfgdefbearertraffic" }
    if yname == "cMediaGwCcCfgDefRtpNamePrefix" { return "Cmediagwcccfgdefrtpnameprefix" }
    return ""
}

func (cmediagwcallcontrolconfigentry *CISCOMEDIAGATEWAYMIB_Cmediagwcallcontrolconfigtable_Cmediagwcallcontrolconfigentry) GetSegmentPath() string {
    return "cMediaGwCallControlConfigEntry" + "[cmgwIndex='" + fmt.Sprintf("%v", cmediagwcallcontrolconfigentry.Cmgwindex) + "']"
}

func (cmediagwcallcontrolconfigentry *CISCOMEDIAGATEWAYMIB_Cmediagwcallcontrolconfigtable_Cmediagwcallcontrolconfigentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cmediagwcallcontrolconfigentry *CISCOMEDIAGATEWAYMIB_Cmediagwcallcontrolconfigtable_Cmediagwcallcontrolconfigentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cmediagwcallcontrolconfigentry *CISCOMEDIAGATEWAYMIB_Cmediagwcallcontrolconfigtable_Cmediagwcallcontrolconfigentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cmgwIndex"] = cmediagwcallcontrolconfigentry.Cmgwindex
    leafs["cMediaGwCcCfgControlTos"] = cmediagwcallcontrolconfigentry.Cmediagwcccfgcontroltos
    leafs["cMediaGwCcCfgBearerTos"] = cmediagwcallcontrolconfigentry.Cmediagwcccfgbearertos
    leafs["cMediaGwCcCfgNtePayload"] = cmediagwcallcontrolconfigentry.Cmediagwcccfgntepayload
    leafs["cMediaGwCcCfgNsePayload"] = cmediagwcallcontrolconfigentry.Cmediagwcccfgnsepayload
    leafs["cMediaGwCcCfgNseRespTimer"] = cmediagwcallcontrolconfigentry.Cmediagwcccfgnseresptimer
    leafs["cMediaGwCcCfgVbdJitterDelayMode"] = cmediagwcallcontrolconfigentry.Cmediagwcccfgvbdjitterdelaymode
    leafs["cMediaGwCcCfgVbdJitterMaxDelay"] = cmediagwcallcontrolconfigentry.Cmediagwcccfgvbdjittermaxdelay
    leafs["cMediaGwCcCfgVbdJitterNomDelay"] = cmediagwcallcontrolconfigentry.Cmediagwcccfgvbdjitternomdelay
    leafs["cMediaGwCcCfgVbdJitterMinDelay"] = cmediagwcallcontrolconfigentry.Cmediagwcccfgvbdjittermindelay
    leafs["cMediaGwCcCfgDefaultTonePlanId"] = cmediagwcallcontrolconfigentry.Cmediagwcccfgdefaulttoneplanid
    leafs["cMediaGwCcCfgDescrInfoEnabled"] = cmediagwcallcontrolconfigentry.Cmediagwcccfgdescrinfoenabled
    leafs["cMediaGwCcCfgDsNamePrefix"] = cmediagwcallcontrolconfigentry.Cmediagwcccfgdsnameprefix
    leafs["cMediaGwCcCfgRtpNamePrefix"] = cmediagwcallcontrolconfigentry.Cmediagwcccfgrtpnameprefix
    leafs["cMediaGwCcCfgAal1SvcNamePrefix"] = cmediagwcallcontrolconfigentry.Cmediagwcccfgaal1Svcnameprefix
    leafs["cMediaGwCcCfgAal2SvcNamePrefix"] = cmediagwcallcontrolconfigentry.Cmediagwcccfgaal2Svcnameprefix
    leafs["cMediaGwCcCfgClusterEnabled"] = cmediagwcallcontrolconfigentry.Cmediagwcccfgclusterenabled
    leafs["cMediaGwCcCfgDefBearerTraffic"] = cmediagwcallcontrolconfigentry.Cmediagwcccfgdefbearertraffic
    leafs["cMediaGwCcCfgDefRtpNamePrefix"] = cmediagwcallcontrolconfigentry.Cmediagwcccfgdefrtpnameprefix
    return leafs
}

func (cmediagwcallcontrolconfigentry *CISCOMEDIAGATEWAYMIB_Cmediagwcallcontrolconfigtable_Cmediagwcallcontrolconfigentry) GetBundleName() string { return "cisco_ios_xe" }

func (cmediagwcallcontrolconfigentry *CISCOMEDIAGATEWAYMIB_Cmediagwcallcontrolconfigtable_Cmediagwcallcontrolconfigentry) GetYangName() string { return "cMediaGwCallControlConfigEntry" }

func (cmediagwcallcontrolconfigentry *CISCOMEDIAGATEWAYMIB_Cmediagwcallcontrolconfigtable_Cmediagwcallcontrolconfigentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cmediagwcallcontrolconfigentry *CISCOMEDIAGATEWAYMIB_Cmediagwcallcontrolconfigtable_Cmediagwcallcontrolconfigentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cmediagwcallcontrolconfigentry *CISCOMEDIAGATEWAYMIB_Cmediagwcallcontrolconfigtable_Cmediagwcallcontrolconfigentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cmediagwcallcontrolconfigentry *CISCOMEDIAGATEWAYMIB_Cmediagwcallcontrolconfigtable_Cmediagwcallcontrolconfigentry) SetParent(parent types.Entity) { cmediagwcallcontrolconfigentry.parent = parent }

func (cmediagwcallcontrolconfigentry *CISCOMEDIAGATEWAYMIB_Cmediagwcallcontrolconfigtable_Cmediagwcallcontrolconfigentry) GetParent() types.Entity { return cmediagwcallcontrolconfigentry.parent }

func (cmediagwcallcontrolconfigentry *CISCOMEDIAGATEWAYMIB_Cmediagwcallcontrolconfigtable_Cmediagwcallcontrolconfigentry) GetParentYangName() string { return "cMediaGwCallControlConfigTable" }

// CISCOMEDIAGATEWAYMIB_Cmediagwcallcontrolconfigtable_Cmediagwcallcontrolconfigentry_Cmediagwcccfgclusterenabled represents               attribute is upon MGC request.
type CISCOMEDIAGATEWAYMIB_Cmediagwcallcontrolconfigtable_Cmediagwcallcontrolconfigentry_Cmediagwcccfgclusterenabled string

const (
    CISCOMEDIAGATEWAYMIB_Cmediagwcallcontrolconfigtable_Cmediagwcallcontrolconfigentry_Cmediagwcccfgclusterenabled_disabled CISCOMEDIAGATEWAYMIB_Cmediagwcallcontrolconfigtable_Cmediagwcallcontrolconfigentry_Cmediagwcccfgclusterenabled = "disabled"

    CISCOMEDIAGATEWAYMIB_Cmediagwcallcontrolconfigtable_Cmediagwcallcontrolconfigentry_Cmediagwcccfgclusterenabled_enabled CISCOMEDIAGATEWAYMIB_Cmediagwcallcontrolconfigtable_Cmediagwcallcontrolconfigentry_Cmediagwcccfgclusterenabled = "enabled"

    CISCOMEDIAGATEWAYMIB_Cmediagwcallcontrolconfigtable_Cmediagwcallcontrolconfigentry_Cmediagwcccfgclusterenabled_conditionalEnabled CISCOMEDIAGATEWAYMIB_Cmediagwcallcontrolconfigtable_Cmediagwcallcontrolconfigentry_Cmediagwcccfgclusterenabled = "conditionalEnabled"
)

// CISCOMEDIAGATEWAYMIB_Cmediagwcallcontrolconfigtable_Cmediagwcallcontrolconfigentry_Cmediagwcccfgdefbearertraffic represents media gateway sending CRCX response.
type CISCOMEDIAGATEWAYMIB_Cmediagwcallcontrolconfigtable_Cmediagwcallcontrolconfigentry_Cmediagwcccfgdefbearertraffic string

const (
    CISCOMEDIAGATEWAYMIB_Cmediagwcallcontrolconfigtable_Cmediagwcallcontrolconfigentry_Cmediagwcccfgdefbearertraffic_ipPvcAal5 CISCOMEDIAGATEWAYMIB_Cmediagwcallcontrolconfigtable_Cmediagwcallcontrolconfigentry_Cmediagwcccfgdefbearertraffic = "ipPvcAal5"

    CISCOMEDIAGATEWAYMIB_Cmediagwcallcontrolconfigtable_Cmediagwcallcontrolconfigentry_Cmediagwcccfgdefbearertraffic_atmPvcAal2 CISCOMEDIAGATEWAYMIB_Cmediagwcallcontrolconfigtable_Cmediagwcallcontrolconfigentry_Cmediagwcccfgdefbearertraffic = "atmPvcAal2"

    CISCOMEDIAGATEWAYMIB_Cmediagwcallcontrolconfigtable_Cmediagwcallcontrolconfigentry_Cmediagwcccfgdefbearertraffic_atmSvcAal2 CISCOMEDIAGATEWAYMIB_Cmediagwcallcontrolconfigtable_Cmediagwcallcontrolconfigentry_Cmediagwcccfgdefbearertraffic = "atmSvcAal2"

    CISCOMEDIAGATEWAYMIB_Cmediagwcallcontrolconfigtable_Cmediagwcallcontrolconfigentry_Cmediagwcccfgdefbearertraffic_atmSvcAal1 CISCOMEDIAGATEWAYMIB_Cmediagwcallcontrolconfigtable_Cmediagwcallcontrolconfigentry_Cmediagwcccfgdefbearertraffic = "atmSvcAal1"
)

// CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable
// This table stores the gateway resource statistics
// information.
type CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry stores the statistics information for a specific resource. The
    // type is slice of
    // CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable_Cmediagwrscstatsentry.
    Cmediagwrscstatsentry []CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable_Cmediagwrscstatsentry
}

func (cmediagwrscstatstable *CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable) GetFilter() yfilter.YFilter { return cmediagwrscstatstable.YFilter }

func (cmediagwrscstatstable *CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable) SetFilter(yf yfilter.YFilter) { cmediagwrscstatstable.YFilter = yf }

func (cmediagwrscstatstable *CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable) GetGoName(yname string) string {
    if yname == "cMediaGwRscStatsEntry" { return "Cmediagwrscstatsentry" }
    return ""
}

func (cmediagwrscstatstable *CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable) GetSegmentPath() string {
    return "cMediaGwRscStatsTable"
}

func (cmediagwrscstatstable *CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cMediaGwRscStatsEntry" {
        for _, c := range cmediagwrscstatstable.Cmediagwrscstatsentry {
            if cmediagwrscstatstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable_Cmediagwrscstatsentry{}
        cmediagwrscstatstable.Cmediagwrscstatsentry = append(cmediagwrscstatstable.Cmediagwrscstatsentry, child)
        return &cmediagwrscstatstable.Cmediagwrscstatsentry[len(cmediagwrscstatstable.Cmediagwrscstatsentry)-1]
    }
    return nil
}

func (cmediagwrscstatstable *CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cmediagwrscstatstable.Cmediagwrscstatsentry {
        children[cmediagwrscstatstable.Cmediagwrscstatsentry[i].GetSegmentPath()] = &cmediagwrscstatstable.Cmediagwrscstatsentry[i]
    }
    return children
}

func (cmediagwrscstatstable *CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cmediagwrscstatstable *CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable) GetBundleName() string { return "cisco_ios_xe" }

func (cmediagwrscstatstable *CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable) GetYangName() string { return "cMediaGwRscStatsTable" }

func (cmediagwrscstatstable *CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cmediagwrscstatstable *CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cmediagwrscstatstable *CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cmediagwrscstatstable *CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable) SetParent(parent types.Entity) { cmediagwrscstatstable.parent = parent }

func (cmediagwrscstatstable *CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable) GetParent() types.Entity { return cmediagwrscstatstable.parent }

func (cmediagwrscstatstable *CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable) GetParentYangName() string { return "CISCO-MEDIA-GATEWAY-MIB" }

// CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable_Cmediagwrscstatsentry
// Each entry stores the statistics
// information for a specific resource.
type CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable_Cmediagwrscstatsentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // cisco_media_gateway_mib.CISCOMEDIAGATEWAYMIB_Cmediagwtable_Cmediagwentry_Cmgwindex
    Cmgwindex interface{}

    // This attribute is a key. An index that uniquely identifies a specific
    // gateway resource. The type is Cmgwrscstatsindex.
    Cmgwrscstatsindex interface{}

    // This object indicates the maximum utilization of the resource over the
    // interval specified by the 'cmgwRscSinceLastReset'. The type is interface{}
    // with range: 0..4294967295.
    Cmgwrscmaximumutilization interface{}

    // This object indicates the minimum utilization of the resource over the
    // interval specified by the 'cmgwRscSinceLastReset'. The type is interface{}
    // with range: 0..4294967295.
    Cmgwrscminimumutilization interface{}

    // This object indicates the average utilization of the resource over the
    // interval specified by the 'cmgwRscSinceLastReset'. The type is interface{}
    // with range: 0..4294967295.
    Cmgwrscaverageutilization interface{}

    // The elapsed time (in seconds) from the last periodic reset.  The following
    // objects are reset at the last reset:      'cmgwRscMaximumUtilization'    
    // 'cmgwRscMinimumUtilization'     'cmgwRscAverageUtilization'. The type is
    // interface{} with range: 0..4294967295. Units are seconds.
    Cmgwrscsincelastreset interface{}
}

func (cmediagwrscstatsentry *CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable_Cmediagwrscstatsentry) GetFilter() yfilter.YFilter { return cmediagwrscstatsentry.YFilter }

func (cmediagwrscstatsentry *CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable_Cmediagwrscstatsentry) SetFilter(yf yfilter.YFilter) { cmediagwrscstatsentry.YFilter = yf }

func (cmediagwrscstatsentry *CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable_Cmediagwrscstatsentry) GetGoName(yname string) string {
    if yname == "cmgwIndex" { return "Cmgwindex" }
    if yname == "cmgwRscStatsIndex" { return "Cmgwrscstatsindex" }
    if yname == "cmgwRscMaximumUtilization" { return "Cmgwrscmaximumutilization" }
    if yname == "cmgwRscMinimumUtilization" { return "Cmgwrscminimumutilization" }
    if yname == "cmgwRscAverageUtilization" { return "Cmgwrscaverageutilization" }
    if yname == "cmgwRscSinceLastReset" { return "Cmgwrscsincelastreset" }
    return ""
}

func (cmediagwrscstatsentry *CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable_Cmediagwrscstatsentry) GetSegmentPath() string {
    return "cMediaGwRscStatsEntry" + "[cmgwIndex='" + fmt.Sprintf("%v", cmediagwrscstatsentry.Cmgwindex) + "']" + "[cmgwRscStatsIndex='" + fmt.Sprintf("%v", cmediagwrscstatsentry.Cmgwrscstatsindex) + "']"
}

func (cmediagwrscstatsentry *CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable_Cmediagwrscstatsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cmediagwrscstatsentry *CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable_Cmediagwrscstatsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cmediagwrscstatsentry *CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable_Cmediagwrscstatsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cmgwIndex"] = cmediagwrscstatsentry.Cmgwindex
    leafs["cmgwRscStatsIndex"] = cmediagwrscstatsentry.Cmgwrscstatsindex
    leafs["cmgwRscMaximumUtilization"] = cmediagwrscstatsentry.Cmgwrscmaximumutilization
    leafs["cmgwRscMinimumUtilization"] = cmediagwrscstatsentry.Cmgwrscminimumutilization
    leafs["cmgwRscAverageUtilization"] = cmediagwrscstatsentry.Cmgwrscaverageutilization
    leafs["cmgwRscSinceLastReset"] = cmediagwrscstatsentry.Cmgwrscsincelastreset
    return leafs
}

func (cmediagwrscstatsentry *CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable_Cmediagwrscstatsentry) GetBundleName() string { return "cisco_ios_xe" }

func (cmediagwrscstatsentry *CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable_Cmediagwrscstatsentry) GetYangName() string { return "cMediaGwRscStatsEntry" }

func (cmediagwrscstatsentry *CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable_Cmediagwrscstatsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cmediagwrscstatsentry *CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable_Cmediagwrscstatsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cmediagwrscstatsentry *CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable_Cmediagwrscstatsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cmediagwrscstatsentry *CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable_Cmediagwrscstatsentry) SetParent(parent types.Entity) { cmediagwrscstatsentry.parent = parent }

func (cmediagwrscstatsentry *CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable_Cmediagwrscstatsentry) GetParent() types.Entity { return cmediagwrscstatsentry.parent }

func (cmediagwrscstatsentry *CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable_Cmediagwrscstatsentry) GetParentYangName() string { return "cMediaGwRscStatsTable" }

// CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable_Cmediagwrscstatsentry_Cmgwrscstatsindex represents resource.
type CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable_Cmediagwrscstatsentry_Cmgwrscstatsindex string

const (
    CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable_Cmediagwrscstatsentry_Cmgwrscstatsindex_cpu CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable_Cmediagwrscstatsentry_Cmgwrscstatsindex = "cpu"

    CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable_Cmediagwrscstatsentry_Cmgwrscstatsindex_staticmemory CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable_Cmediagwrscstatsentry_Cmgwrscstatsindex = "staticmemory"

    CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable_Cmediagwrscstatsentry_Cmgwrscstatsindex_dynamicmemory CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable_Cmediagwrscstatsentry_Cmgwrscstatsindex = "dynamicmemory"

    CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable_Cmediagwrscstatsentry_Cmgwrscstatsindex_sysmemory CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable_Cmediagwrscstatsentry_Cmgwrscstatsindex = "sysmemory"

    CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable_Cmediagwrscstatsentry_Cmgwrscstatsindex_commbuffer CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable_Cmediagwrscstatsentry_Cmgwrscstatsindex = "commbuffer"

    CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable_Cmediagwrscstatsentry_Cmgwrscstatsindex_msgq CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable_Cmediagwrscstatsentry_Cmgwrscstatsindex = "msgq"

    CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable_Cmediagwrscstatsentry_Cmgwrscstatsindex_atmq CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable_Cmediagwrscstatsentry_Cmgwrscstatsindex = "atmq"

    CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable_Cmediagwrscstatsentry_Cmgwrscstatsindex_svccongestion CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable_Cmediagwrscstatsentry_Cmgwrscstatsindex = "svccongestion"

    CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable_Cmediagwrscstatsentry_Cmgwrscstatsindex_rsvpq CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable_Cmediagwrscstatsentry_Cmgwrscstatsindex = "rsvpq"

    CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable_Cmediagwrscstatsentry_Cmgwrscstatsindex_dspq CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable_Cmediagwrscstatsentry_Cmgwrscstatsindex = "dspq"

    CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable_Cmediagwrscstatsentry_Cmgwrscstatsindex_h248congestion CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable_Cmediagwrscstatsentry_Cmgwrscstatsindex = "h248congestion"

    CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable_Cmediagwrscstatsentry_Cmgwrscstatsindex_callpersec CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable_Cmediagwrscstatsentry_Cmgwrscstatsindex = "callpersec"

    CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable_Cmediagwrscstatsentry_Cmgwrscstatsindex_smallipcbuffer CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable_Cmediagwrscstatsentry_Cmgwrscstatsindex = "smallipcbuffer"

    CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable_Cmediagwrscstatsentry_Cmgwrscstatsindex_mediumipcbuffer CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable_Cmediagwrscstatsentry_Cmgwrscstatsindex = "mediumipcbuffer"

    CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable_Cmediagwrscstatsentry_Cmgwrscstatsindex_largeipcbuffer CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable_Cmediagwrscstatsentry_Cmgwrscstatsindex = "largeipcbuffer"

    CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable_Cmediagwrscstatsentry_Cmgwrscstatsindex_hugeipcbuffer CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable_Cmediagwrscstatsentry_Cmgwrscstatsindex = "hugeipcbuffer"

    CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable_Cmediagwrscstatsentry_Cmgwrscstatsindex_mblkipcbuffer CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable_Cmediagwrscstatsentry_Cmgwrscstatsindex = "mblkipcbuffer"
)

