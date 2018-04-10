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

// CGwServiceState represents     No new connections are allowed.
type CGwServiceState string

const (
    CGwServiceState_inService CGwServiceState = "inService"

    CGwServiceState_forcedOutOfService CGwServiceState = "forcedOutOfService"

    CGwServiceState_gracefulOutOfService CGwServiceState = "gracefulOutOfService"
)

// CGwAdminState represents       New connections would be blocked.
type CGwAdminState string

const (
    CGwAdminState_inService CGwAdminState = "inService"

    CGwAdminState_forcedOutOfService CGwAdminState = "forcedOutOfService"

    CGwAdminState_gracefulOutOfService CGwAdminState = "gracefulOutOfService"
)

// CCallControlJitterDelayMode represents            which is specified by jitter nominal delay.
type CCallControlJitterDelayMode string

const (
    CCallControlJitterDelayMode_adaptive CCallControlJitterDelayMode = "adaptive"

    CCallControlJitterDelayMode_fixed CCallControlJitterDelayMode = "fixed"
)

// CISCOMEDIAGATEWAYMIB
type CISCOMEDIAGATEWAYMIB struct {
    EntityData types.CommonEntityData
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

func (cISCOMEDIAGATEWAYMIB *CISCOMEDIAGATEWAYMIB) GetEntityData() *types.CommonEntityData {
    cISCOMEDIAGATEWAYMIB.EntityData.YFilter = cISCOMEDIAGATEWAYMIB.YFilter
    cISCOMEDIAGATEWAYMIB.EntityData.YangName = "CISCO-MEDIA-GATEWAY-MIB"
    cISCOMEDIAGATEWAYMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOMEDIAGATEWAYMIB.EntityData.ParentYangName = "CISCO-MEDIA-GATEWAY-MIB"
    cISCOMEDIAGATEWAYMIB.EntityData.SegmentPath = "CISCO-MEDIA-GATEWAY-MIB:CISCO-MEDIA-GATEWAY-MIB"
    cISCOMEDIAGATEWAYMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOMEDIAGATEWAYMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOMEDIAGATEWAYMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOMEDIAGATEWAYMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOMEDIAGATEWAYMIB.EntityData.Children["cMediaGwTable"] = types.YChild{"Cmediagwtable", &cISCOMEDIAGATEWAYMIB.Cmediagwtable}
    cISCOMEDIAGATEWAYMIB.EntityData.Children["cmgwSignalProtocolTable"] = types.YChild{"Cmgwsignalprotocoltable", &cISCOMEDIAGATEWAYMIB.Cmgwsignalprotocoltable}
    cISCOMEDIAGATEWAYMIB.EntityData.Children["cMediaGwIpConfigTable"] = types.YChild{"Cmediagwipconfigtable", &cISCOMEDIAGATEWAYMIB.Cmediagwipconfigtable}
    cISCOMEDIAGATEWAYMIB.EntityData.Children["cMediaGwDomainNameConfigTable"] = types.YChild{"Cmediagwdomainnameconfigtable", &cISCOMEDIAGATEWAYMIB.Cmediagwdomainnameconfigtable}
    cISCOMEDIAGATEWAYMIB.EntityData.Children["cMediaGwDnsIpConfigTable"] = types.YChild{"Cmediagwdnsipconfigtable", &cISCOMEDIAGATEWAYMIB.Cmediagwdnsipconfigtable}
    cISCOMEDIAGATEWAYMIB.EntityData.Children["cmgwLifTable"] = types.YChild{"Cmgwliftable", &cISCOMEDIAGATEWAYMIB.Cmgwliftable}
    cISCOMEDIAGATEWAYMIB.EntityData.Children["cMediaGwCallControlConfigTable"] = types.YChild{"Cmediagwcallcontrolconfigtable", &cISCOMEDIAGATEWAYMIB.Cmediagwcallcontrolconfigtable}
    cISCOMEDIAGATEWAYMIB.EntityData.Children["cMediaGwRscStatsTable"] = types.YChild{"Cmediagwrscstatstable", &cISCOMEDIAGATEWAYMIB.Cmediagwrscstatstable}
    cISCOMEDIAGATEWAYMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOMEDIAGATEWAYMIB.EntityData)
}

// CISCOMEDIAGATEWAYMIB_Cmediagwtable
// This table contains the global media gateway parameters
// information.
// It supports the modification of the global media gateway 
// parameters.
type CISCOMEDIAGATEWAYMIB_Cmediagwtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A Media Gateway Entry.   At system power-up, an entry is created by the
    // agent  if the system detects a media gateway module has been added  to the
    // system, and an entry is deleted if the entry associated media gateway
    // module has been removed from the system. The type is slice of
    // CISCOMEDIAGATEWAYMIB_Cmediagwtable_Cmediagwentry.
    Cmediagwentry []CISCOMEDIAGATEWAYMIB_Cmediagwtable_Cmediagwentry
}

func (cmediagwtable *CISCOMEDIAGATEWAYMIB_Cmediagwtable) GetEntityData() *types.CommonEntityData {
    cmediagwtable.EntityData.YFilter = cmediagwtable.YFilter
    cmediagwtable.EntityData.YangName = "cMediaGwTable"
    cmediagwtable.EntityData.BundleName = "cisco_ios_xe"
    cmediagwtable.EntityData.ParentYangName = "CISCO-MEDIA-GATEWAY-MIB"
    cmediagwtable.EntityData.SegmentPath = "cMediaGwTable"
    cmediagwtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cmediagwtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cmediagwtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cmediagwtable.EntityData.Children = make(map[string]types.YChild)
    cmediagwtable.EntityData.Children["cMediaGwEntry"] = types.YChild{"Cmediagwentry", nil}
    for i := range cmediagwtable.Cmediagwentry {
        cmediagwtable.EntityData.Children[types.GetSegmentPath(&cmediagwtable.Cmediagwentry[i])] = types.YChild{"Cmediagwentry", &cmediagwtable.Cmediagwentry[i]}
    }
    cmediagwtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cmediagwtable.EntityData)
}

// CISCOMEDIAGATEWAYMIB_Cmediagwtable_Cmediagwentry
// A Media Gateway Entry.  
// At system power-up, an entry is created by the agent 
// if the system detects a media gateway module has been added 
// to the system, and an entry is deleted if the entry associated
// media gateway module has been removed from the system.
type CISCOMEDIAGATEWAYMIB_Cmediagwtable_Cmediagwentry struct {
    EntityData types.CommonEntityData
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

func (cmediagwentry *CISCOMEDIAGATEWAYMIB_Cmediagwtable_Cmediagwentry) GetEntityData() *types.CommonEntityData {
    cmediagwentry.EntityData.YFilter = cmediagwentry.YFilter
    cmediagwentry.EntityData.YangName = "cMediaGwEntry"
    cmediagwentry.EntityData.BundleName = "cisco_ios_xe"
    cmediagwentry.EntityData.ParentYangName = "cMediaGwTable"
    cmediagwentry.EntityData.SegmentPath = "cMediaGwEntry" + "[cmgwIndex='" + fmt.Sprintf("%v", cmediagwentry.Cmgwindex) + "']"
    cmediagwentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cmediagwentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cmediagwentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cmediagwentry.EntityData.Children = make(map[string]types.YChild)
    cmediagwentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cmediagwentry.EntityData.Leafs["cmgwIndex"] = types.YLeaf{"Cmgwindex", cmediagwentry.Cmgwindex}
    cmediagwentry.EntityData.Leafs["cmgwDomainName"] = types.YLeaf{"Cmgwdomainname", cmediagwentry.Cmgwdomainname}
    cmediagwentry.EntityData.Leafs["cmgwPhysicalIndex"] = types.YLeaf{"Cmgwphysicalindex", cmediagwentry.Cmgwphysicalindex}
    cmediagwentry.EntityData.Leafs["cmgwServiceState"] = types.YLeaf{"Cmgwservicestate", cmediagwentry.Cmgwservicestate}
    cmediagwentry.EntityData.Leafs["cmgwAdminState"] = types.YLeaf{"Cmgwadminstate", cmediagwentry.Cmgwadminstate}
    cmediagwentry.EntityData.Leafs["cmgwGraceTime"] = types.YLeaf{"Cmgwgracetime", cmediagwentry.Cmgwgracetime}
    cmediagwentry.EntityData.Leafs["cmgwVtMappingMode"] = types.YLeaf{"Cmgwvtmappingmode", cmediagwentry.Cmgwvtmappingmode}
    cmediagwentry.EntityData.Leafs["cmgwSrcFilterEnabled"] = types.YLeaf{"Cmgwsrcfilterenabled", cmediagwentry.Cmgwsrcfilterenabled}
    cmediagwentry.EntityData.Leafs["cmgwLawInterceptEnabled"] = types.YLeaf{"Cmgwlawinterceptenabled", cmediagwentry.Cmgwlawinterceptenabled}
    cmediagwentry.EntityData.Leafs["cmgwV23Enabled"] = types.YLeaf{"Cmgwv23Enabled", cmediagwentry.Cmgwv23Enabled}
    return &(cmediagwentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry represents an signaling protocol supported by the media gateway.
    // The type is slice of
    // CISCOMEDIAGATEWAYMIB_Cmgwsignalprotocoltable_Cmgwsignalprotocolentry.
    Cmgwsignalprotocolentry []CISCOMEDIAGATEWAYMIB_Cmgwsignalprotocoltable_Cmgwsignalprotocolentry
}

func (cmgwsignalprotocoltable *CISCOMEDIAGATEWAYMIB_Cmgwsignalprotocoltable) GetEntityData() *types.CommonEntityData {
    cmgwsignalprotocoltable.EntityData.YFilter = cmgwsignalprotocoltable.YFilter
    cmgwsignalprotocoltable.EntityData.YangName = "cmgwSignalProtocolTable"
    cmgwsignalprotocoltable.EntityData.BundleName = "cisco_ios_xe"
    cmgwsignalprotocoltable.EntityData.ParentYangName = "CISCO-MEDIA-GATEWAY-MIB"
    cmgwsignalprotocoltable.EntityData.SegmentPath = "cmgwSignalProtocolTable"
    cmgwsignalprotocoltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cmgwsignalprotocoltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cmgwsignalprotocoltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cmgwsignalprotocoltable.EntityData.Children = make(map[string]types.YChild)
    cmgwsignalprotocoltable.EntityData.Children["cmgwSignalProtocolEntry"] = types.YChild{"Cmgwsignalprotocolentry", nil}
    for i := range cmgwsignalprotocoltable.Cmgwsignalprotocolentry {
        cmgwsignalprotocoltable.EntityData.Children[types.GetSegmentPath(&cmgwsignalprotocoltable.Cmgwsignalprotocolentry[i])] = types.YChild{"Cmgwsignalprotocolentry", &cmgwsignalprotocoltable.Cmgwsignalprotocolentry[i]}
    }
    cmgwsignalprotocoltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cmgwsignalprotocoltable.EntityData)
}

// CISCOMEDIAGATEWAYMIB_Cmgwsignalprotocoltable_Cmgwsignalprotocolentry
// Each entry represents an signaling protocol supported
// by the media gateway.
type CISCOMEDIAGATEWAYMIB_Cmgwsignalprotocoltable_Cmgwsignalprotocolentry struct {
    EntityData types.CommonEntityData
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

func (cmgwsignalprotocolentry *CISCOMEDIAGATEWAYMIB_Cmgwsignalprotocoltable_Cmgwsignalprotocolentry) GetEntityData() *types.CommonEntityData {
    cmgwsignalprotocolentry.EntityData.YFilter = cmgwsignalprotocolentry.YFilter
    cmgwsignalprotocolentry.EntityData.YangName = "cmgwSignalProtocolEntry"
    cmgwsignalprotocolentry.EntityData.BundleName = "cisco_ios_xe"
    cmgwsignalprotocolentry.EntityData.ParentYangName = "cmgwSignalProtocolTable"
    cmgwsignalprotocolentry.EntityData.SegmentPath = "cmgwSignalProtocolEntry" + "[cmgwIndex='" + fmt.Sprintf("%v", cmgwsignalprotocolentry.Cmgwindex) + "']" + "[cmgwSignalProtocolIndex='" + fmt.Sprintf("%v", cmgwsignalprotocolentry.Cmgwsignalprotocolindex) + "']"
    cmgwsignalprotocolentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cmgwsignalprotocolentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cmgwsignalprotocolentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cmgwsignalprotocolentry.EntityData.Children = make(map[string]types.YChild)
    cmgwsignalprotocolentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cmgwsignalprotocolentry.EntityData.Leafs["cmgwIndex"] = types.YLeaf{"Cmgwindex", cmgwsignalprotocolentry.Cmgwindex}
    cmgwsignalprotocolentry.EntityData.Leafs["cmgwSignalProtocolIndex"] = types.YLeaf{"Cmgwsignalprotocolindex", cmgwsignalprotocolentry.Cmgwsignalprotocolindex}
    cmgwsignalprotocolentry.EntityData.Leafs["cmgwSignalProtocol"] = types.YLeaf{"Cmgwsignalprotocol", cmgwsignalprotocolentry.Cmgwsignalprotocol}
    cmgwsignalprotocolentry.EntityData.Leafs["cmgwSignalProtocolVersion"] = types.YLeaf{"Cmgwsignalprotocolversion", cmgwsignalprotocolentry.Cmgwsignalprotocolversion}
    cmgwsignalprotocolentry.EntityData.Leafs["cmgwSignalProtocolPort"] = types.YLeaf{"Cmgwsignalprotocolport", cmgwsignalprotocolentry.Cmgwsignalprotocolport}
    cmgwsignalprotocolentry.EntityData.Leafs["cmgwSignalMgcProtocolPort"] = types.YLeaf{"Cmgwsignalmgcprotocolport", cmgwsignalprotocolentry.Cmgwsignalmgcprotocolport}
    cmgwsignalprotocolentry.EntityData.Leafs["cmgwSignalProtocolPreference"] = types.YLeaf{"Cmgwsignalprotocolpreference", cmgwsignalprotocolentry.Cmgwsignalprotocolpreference}
    cmgwsignalprotocolentry.EntityData.Leafs["cmgwSignalProtocolConfigVer"] = types.YLeaf{"Cmgwsignalprotocolconfigver", cmgwsignalprotocolentry.Cmgwsignalprotocolconfigver}
    return &(cmgwsignalprotocolentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A Media Gateway IP configuration entry.  Each entry represents a media
    // gateway IP address for MGCs to communicate with the media gateway. The type
    // is slice of
    // CISCOMEDIAGATEWAYMIB_Cmediagwipconfigtable_Cmediagwipconfigentry.
    Cmediagwipconfigentry []CISCOMEDIAGATEWAYMIB_Cmediagwipconfigtable_Cmediagwipconfigentry
}

func (cmediagwipconfigtable *CISCOMEDIAGATEWAYMIB_Cmediagwipconfigtable) GetEntityData() *types.CommonEntityData {
    cmediagwipconfigtable.EntityData.YFilter = cmediagwipconfigtable.YFilter
    cmediagwipconfigtable.EntityData.YangName = "cMediaGwIpConfigTable"
    cmediagwipconfigtable.EntityData.BundleName = "cisco_ios_xe"
    cmediagwipconfigtable.EntityData.ParentYangName = "CISCO-MEDIA-GATEWAY-MIB"
    cmediagwipconfigtable.EntityData.SegmentPath = "cMediaGwIpConfigTable"
    cmediagwipconfigtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cmediagwipconfigtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cmediagwipconfigtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cmediagwipconfigtable.EntityData.Children = make(map[string]types.YChild)
    cmediagwipconfigtable.EntityData.Children["cMediaGwIpConfigEntry"] = types.YChild{"Cmediagwipconfigentry", nil}
    for i := range cmediagwipconfigtable.Cmediagwipconfigentry {
        cmediagwipconfigtable.EntityData.Children[types.GetSegmentPath(&cmediagwipconfigtable.Cmediagwipconfigentry[i])] = types.YChild{"Cmediagwipconfigentry", &cmediagwipconfigtable.Cmediagwipconfigentry[i]}
    }
    cmediagwipconfigtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cmediagwipconfigtable.EntityData)
}

// CISCOMEDIAGATEWAYMIB_Cmediagwipconfigtable_Cmediagwipconfigentry
// A Media Gateway IP configuration entry. 
// Each entry represents a media gateway IP address for MGCs
// to communicate with the media gateway.
type CISCOMEDIAGATEWAYMIB_Cmediagwipconfigtable_Cmediagwipconfigentry struct {
    EntityData types.CommonEntityData
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

func (cmediagwipconfigentry *CISCOMEDIAGATEWAYMIB_Cmediagwipconfigtable_Cmediagwipconfigentry) GetEntityData() *types.CommonEntityData {
    cmediagwipconfigentry.EntityData.YFilter = cmediagwipconfigentry.YFilter
    cmediagwipconfigentry.EntityData.YangName = "cMediaGwIpConfigEntry"
    cmediagwipconfigentry.EntityData.BundleName = "cisco_ios_xe"
    cmediagwipconfigentry.EntityData.ParentYangName = "cMediaGwIpConfigTable"
    cmediagwipconfigentry.EntityData.SegmentPath = "cMediaGwIpConfigEntry" + "[cmgwIndex='" + fmt.Sprintf("%v", cmediagwipconfigentry.Cmgwindex) + "']" + "[cmgwIpConfigIndex='" + fmt.Sprintf("%v", cmediagwipconfigentry.Cmgwipconfigindex) + "']"
    cmediagwipconfigentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cmediagwipconfigentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cmediagwipconfigentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cmediagwipconfigentry.EntityData.Children = make(map[string]types.YChild)
    cmediagwipconfigentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cmediagwipconfigentry.EntityData.Leafs["cmgwIndex"] = types.YLeaf{"Cmgwindex", cmediagwipconfigentry.Cmgwindex}
    cmediagwipconfigentry.EntityData.Leafs["cmgwIpConfigIndex"] = types.YLeaf{"Cmgwipconfigindex", cmediagwipconfigentry.Cmgwipconfigindex}
    cmediagwipconfigentry.EntityData.Leafs["cmgwIpConfigIfIndex"] = types.YLeaf{"Cmgwipconfigifindex", cmediagwipconfigentry.Cmgwipconfigifindex}
    cmediagwipconfigentry.EntityData.Leafs["cmgwIpConfigVpi"] = types.YLeaf{"Cmgwipconfigvpi", cmediagwipconfigentry.Cmgwipconfigvpi}
    cmediagwipconfigentry.EntityData.Leafs["cmgwIpConfigVci"] = types.YLeaf{"Cmgwipconfigvci", cmediagwipconfigentry.Cmgwipconfigvci}
    cmediagwipconfigentry.EntityData.Leafs["cmgwIpConfigAddrType"] = types.YLeaf{"Cmgwipconfigaddrtype", cmediagwipconfigentry.Cmgwipconfigaddrtype}
    cmediagwipconfigentry.EntityData.Leafs["cmgwIpConfigAddress"] = types.YLeaf{"Cmgwipconfigaddress", cmediagwipconfigentry.Cmgwipconfigaddress}
    cmediagwipconfigentry.EntityData.Leafs["cmgwIpConfigSubnetMask"] = types.YLeaf{"Cmgwipconfigsubnetmask", cmediagwipconfigentry.Cmgwipconfigsubnetmask}
    cmediagwipconfigentry.EntityData.Leafs["cmgwIpConfigDefaultGwIp"] = types.YLeaf{"Cmgwipconfigdefaultgwip", cmediagwipconfigentry.Cmgwipconfigdefaultgwip}
    cmediagwipconfigentry.EntityData.Leafs["cmgwIpConfigForRemoteMapping"] = types.YLeaf{"Cmgwipconfigforremotemapping", cmediagwipconfigentry.Cmgwipconfigforremotemapping}
    cmediagwipconfigentry.EntityData.Leafs["cmgwIpConfigRowStatus"] = types.YLeaf{"Cmgwipconfigrowstatus", cmediagwipconfigentry.Cmgwipconfigrowstatus}
    return &(cmediagwipconfigentry.EntityData)
}

// CISCOMEDIAGATEWAYMIB_Cmediagwdomainnameconfigtable
// This table provides the domain names which are configured by 
// users. 
// The domain names can be used to represent IP addresses 
// for:
//     gateway
//     External DNS name server
//     MGC (call agent) 
type CISCOMEDIAGATEWAYMIB_Cmediagwdomainnameconfigtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry represents a domain name used in the system.  Creation and
    // deletion are supported. Modification is prohibited. The type is slice of
    // CISCOMEDIAGATEWAYMIB_Cmediagwdomainnameconfigtable_Cmediagwdomainnameconfigentry.
    Cmediagwdomainnameconfigentry []CISCOMEDIAGATEWAYMIB_Cmediagwdomainnameconfigtable_Cmediagwdomainnameconfigentry
}

func (cmediagwdomainnameconfigtable *CISCOMEDIAGATEWAYMIB_Cmediagwdomainnameconfigtable) GetEntityData() *types.CommonEntityData {
    cmediagwdomainnameconfigtable.EntityData.YFilter = cmediagwdomainnameconfigtable.YFilter
    cmediagwdomainnameconfigtable.EntityData.YangName = "cMediaGwDomainNameConfigTable"
    cmediagwdomainnameconfigtable.EntityData.BundleName = "cisco_ios_xe"
    cmediagwdomainnameconfigtable.EntityData.ParentYangName = "CISCO-MEDIA-GATEWAY-MIB"
    cmediagwdomainnameconfigtable.EntityData.SegmentPath = "cMediaGwDomainNameConfigTable"
    cmediagwdomainnameconfigtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cmediagwdomainnameconfigtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cmediagwdomainnameconfigtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cmediagwdomainnameconfigtable.EntityData.Children = make(map[string]types.YChild)
    cmediagwdomainnameconfigtable.EntityData.Children["cMediaGwDomainNameConfigEntry"] = types.YChild{"Cmediagwdomainnameconfigentry", nil}
    for i := range cmediagwdomainnameconfigtable.Cmediagwdomainnameconfigentry {
        cmediagwdomainnameconfigtable.EntityData.Children[types.GetSegmentPath(&cmediagwdomainnameconfigtable.Cmediagwdomainnameconfigentry[i])] = types.YChild{"Cmediagwdomainnameconfigentry", &cmediagwdomainnameconfigtable.Cmediagwdomainnameconfigentry[i]}
    }
    cmediagwdomainnameconfigtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cmediagwdomainnameconfigtable.EntityData)
}

// CISCOMEDIAGATEWAYMIB_Cmediagwdomainnameconfigtable_Cmediagwdomainnameconfigentry
// Each entry represents a domain name used in the system.
// 
// Creation and deletion are supported. Modification
// is prohibited.
type CISCOMEDIAGATEWAYMIB_Cmediagwdomainnameconfigtable_Cmediagwdomainnameconfigentry struct {
    EntityData types.CommonEntityData
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

func (cmediagwdomainnameconfigentry *CISCOMEDIAGATEWAYMIB_Cmediagwdomainnameconfigtable_Cmediagwdomainnameconfigentry) GetEntityData() *types.CommonEntityData {
    cmediagwdomainnameconfigentry.EntityData.YFilter = cmediagwdomainnameconfigentry.YFilter
    cmediagwdomainnameconfigentry.EntityData.YangName = "cMediaGwDomainNameConfigEntry"
    cmediagwdomainnameconfigentry.EntityData.BundleName = "cisco_ios_xe"
    cmediagwdomainnameconfigentry.EntityData.ParentYangName = "cMediaGwDomainNameConfigTable"
    cmediagwdomainnameconfigentry.EntityData.SegmentPath = "cMediaGwDomainNameConfigEntry" + "[cmgwIndex='" + fmt.Sprintf("%v", cmediagwdomainnameconfigentry.Cmgwindex) + "']" + "[cmgwConfigDomainNameIndex='" + fmt.Sprintf("%v", cmediagwdomainnameconfigentry.Cmgwconfigdomainnameindex) + "']"
    cmediagwdomainnameconfigentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cmediagwdomainnameconfigentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cmediagwdomainnameconfigentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cmediagwdomainnameconfigentry.EntityData.Children = make(map[string]types.YChild)
    cmediagwdomainnameconfigentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cmediagwdomainnameconfigentry.EntityData.Leafs["cmgwIndex"] = types.YLeaf{"Cmgwindex", cmediagwdomainnameconfigentry.Cmgwindex}
    cmediagwdomainnameconfigentry.EntityData.Leafs["cmgwConfigDomainNameIndex"] = types.YLeaf{"Cmgwconfigdomainnameindex", cmediagwdomainnameconfigentry.Cmgwconfigdomainnameindex}
    cmediagwdomainnameconfigentry.EntityData.Leafs["cmgwConfigDomainNameEntity"] = types.YLeaf{"Cmgwconfigdomainnameentity", cmediagwdomainnameconfigentry.Cmgwconfigdomainnameentity}
    cmediagwdomainnameconfigentry.EntityData.Leafs["cmgwConfigDomainName"] = types.YLeaf{"Cmgwconfigdomainname", cmediagwdomainnameconfigentry.Cmgwconfigdomainname}
    cmediagwdomainnameconfigentry.EntityData.Leafs["cmgwConfigDomainNameRowStatus"] = types.YLeaf{"Cmgwconfigdomainnamerowstatus", cmediagwdomainnameconfigentry.Cmgwconfigdomainnamerowstatus}
    return &(cmediagwdomainnameconfigentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry represents an IP address of the DNS name  server. The type is
    // slice of
    // CISCOMEDIAGATEWAYMIB_Cmediagwdnsipconfigtable_Cmediagwdnsipconfigentry.
    Cmediagwdnsipconfigentry []CISCOMEDIAGATEWAYMIB_Cmediagwdnsipconfigtable_Cmediagwdnsipconfigentry
}

func (cmediagwdnsipconfigtable *CISCOMEDIAGATEWAYMIB_Cmediagwdnsipconfigtable) GetEntityData() *types.CommonEntityData {
    cmediagwdnsipconfigtable.EntityData.YFilter = cmediagwdnsipconfigtable.YFilter
    cmediagwdnsipconfigtable.EntityData.YangName = "cMediaGwDnsIpConfigTable"
    cmediagwdnsipconfigtable.EntityData.BundleName = "cisco_ios_xe"
    cmediagwdnsipconfigtable.EntityData.ParentYangName = "CISCO-MEDIA-GATEWAY-MIB"
    cmediagwdnsipconfigtable.EntityData.SegmentPath = "cMediaGwDnsIpConfigTable"
    cmediagwdnsipconfigtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cmediagwdnsipconfigtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cmediagwdnsipconfigtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cmediagwdnsipconfigtable.EntityData.Children = make(map[string]types.YChild)
    cmediagwdnsipconfigtable.EntityData.Children["cMediaGwDnsIpConfigEntry"] = types.YChild{"Cmediagwdnsipconfigentry", nil}
    for i := range cmediagwdnsipconfigtable.Cmediagwdnsipconfigentry {
        cmediagwdnsipconfigtable.EntityData.Children[types.GetSegmentPath(&cmediagwdnsipconfigtable.Cmediagwdnsipconfigentry[i])] = types.YChild{"Cmediagwdnsipconfigentry", &cmediagwdnsipconfigtable.Cmediagwdnsipconfigentry[i]}
    }
    cmediagwdnsipconfigtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cmediagwdnsipconfigtable.EntityData)
}

// CISCOMEDIAGATEWAYMIB_Cmediagwdnsipconfigtable_Cmediagwdnsipconfigentry
// Each entry represents an IP address of the DNS name 
// server.
type CISCOMEDIAGATEWAYMIB_Cmediagwdnsipconfigtable_Cmediagwdnsipconfigentry struct {
    EntityData types.CommonEntityData
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

func (cmediagwdnsipconfigentry *CISCOMEDIAGATEWAYMIB_Cmediagwdnsipconfigtable_Cmediagwdnsipconfigentry) GetEntityData() *types.CommonEntityData {
    cmediagwdnsipconfigentry.EntityData.YFilter = cmediagwdnsipconfigentry.YFilter
    cmediagwdnsipconfigentry.EntityData.YangName = "cMediaGwDnsIpConfigEntry"
    cmediagwdnsipconfigentry.EntityData.BundleName = "cisco_ios_xe"
    cmediagwdnsipconfigentry.EntityData.ParentYangName = "cMediaGwDnsIpConfigTable"
    cmediagwdnsipconfigentry.EntityData.SegmentPath = "cMediaGwDnsIpConfigEntry" + "[cmgwIndex='" + fmt.Sprintf("%v", cmediagwdnsipconfigentry.Cmgwindex) + "']" + "[cmgwDnsIpIndex='" + fmt.Sprintf("%v", cmediagwdnsipconfigentry.Cmgwdnsipindex) + "']"
    cmediagwdnsipconfigentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cmediagwdnsipconfigentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cmediagwdnsipconfigentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cmediagwdnsipconfigentry.EntityData.Children = make(map[string]types.YChild)
    cmediagwdnsipconfigentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cmediagwdnsipconfigentry.EntityData.Leafs["cmgwIndex"] = types.YLeaf{"Cmgwindex", cmediagwdnsipconfigentry.Cmgwindex}
    cmediagwdnsipconfigentry.EntityData.Leafs["cmgwDnsIpIndex"] = types.YLeaf{"Cmgwdnsipindex", cmediagwdnsipconfigentry.Cmgwdnsipindex}
    cmediagwdnsipconfigentry.EntityData.Leafs["cmgwDnsDomainName"] = types.YLeaf{"Cmgwdnsdomainname", cmediagwdnsipconfigentry.Cmgwdnsdomainname}
    cmediagwdnsipconfigentry.EntityData.Leafs["cmgwDnsIpType"] = types.YLeaf{"Cmgwdnsiptype", cmediagwdnsipconfigentry.Cmgwdnsiptype}
    cmediagwdnsipconfigentry.EntityData.Leafs["cmgwDnsIp"] = types.YLeaf{"Cmgwdnsip", cmediagwdnsipconfigentry.Cmgwdnsip}
    cmediagwdnsipconfigentry.EntityData.Leafs["cmgwDnsIpRowStatus"] = types.YLeaf{"Cmgwdnsiprowstatus", cmediagwdnsipconfigentry.Cmgwdnsiprowstatus}
    return &(cmediagwdnsipconfigentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry of this table is created by the media gateway when it supports the
    // VoIP/VoATM application. The type is slice of
    // CISCOMEDIAGATEWAYMIB_Cmgwliftable_Cmgwlifentry.
    Cmgwlifentry []CISCOMEDIAGATEWAYMIB_Cmgwliftable_Cmgwlifentry
}

func (cmgwliftable *CISCOMEDIAGATEWAYMIB_Cmgwliftable) GetEntityData() *types.CommonEntityData {
    cmgwliftable.EntityData.YFilter = cmgwliftable.YFilter
    cmgwliftable.EntityData.YangName = "cmgwLifTable"
    cmgwliftable.EntityData.BundleName = "cisco_ios_xe"
    cmgwliftable.EntityData.ParentYangName = "CISCO-MEDIA-GATEWAY-MIB"
    cmgwliftable.EntityData.SegmentPath = "cmgwLifTable"
    cmgwliftable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cmgwliftable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cmgwliftable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cmgwliftable.EntityData.Children = make(map[string]types.YChild)
    cmgwliftable.EntityData.Children["cmgwLifEntry"] = types.YChild{"Cmgwlifentry", nil}
    for i := range cmgwliftable.Cmgwlifentry {
        cmgwliftable.EntityData.Children[types.GetSegmentPath(&cmgwliftable.Cmgwlifentry[i])] = types.YChild{"Cmgwlifentry", &cmgwliftable.Cmgwlifentry[i]}
    }
    cmgwliftable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cmgwliftable.EntityData)
}

// CISCOMEDIAGATEWAYMIB_Cmgwliftable_Cmgwlifentry
// An entry of this table is created by the media gateway
// when it supports the VoIP/VoATM application.
type CISCOMEDIAGATEWAYMIB_Cmgwliftable_Cmgwlifentry struct {
    EntityData types.CommonEntityData
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

func (cmgwlifentry *CISCOMEDIAGATEWAYMIB_Cmgwliftable_Cmgwlifentry) GetEntityData() *types.CommonEntityData {
    cmgwlifentry.EntityData.YFilter = cmgwlifentry.YFilter
    cmgwlifentry.EntityData.YangName = "cmgwLifEntry"
    cmgwlifentry.EntityData.BundleName = "cisco_ios_xe"
    cmgwlifentry.EntityData.ParentYangName = "cmgwLifTable"
    cmgwlifentry.EntityData.SegmentPath = "cmgwLifEntry" + "[cmgwIndex='" + fmt.Sprintf("%v", cmgwlifentry.Cmgwindex) + "']" + "[cmgwLifNumber='" + fmt.Sprintf("%v", cmgwlifentry.Cmgwlifnumber) + "']"
    cmgwlifentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cmgwlifentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cmgwlifentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cmgwlifentry.EntityData.Children = make(map[string]types.YChild)
    cmgwlifentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cmgwlifentry.EntityData.Leafs["cmgwIndex"] = types.YLeaf{"Cmgwindex", cmgwlifentry.Cmgwindex}
    cmgwlifentry.EntityData.Leafs["cmgwLifNumber"] = types.YLeaf{"Cmgwlifnumber", cmgwlifentry.Cmgwlifnumber}
    cmgwlifentry.EntityData.Leafs["cmgwLifPvcCount"] = types.YLeaf{"Cmgwlifpvccount", cmgwlifentry.Cmgwlifpvccount}
    cmgwlifentry.EntityData.Leafs["cmgwLifVoiceIfCount"] = types.YLeaf{"Cmgwlifvoiceifcount", cmgwlifentry.Cmgwlifvoiceifcount}
    return &(cmgwlifentry.EntityData)
}

// CISCOMEDIAGATEWAYMIB_Cmediagwcallcontrolconfigtable
// This table defines general call control attributes for
// the media gateway.
type CISCOMEDIAGATEWAYMIB_Cmediagwcallcontrolconfigtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // One entry for each media gateway which supports call control  protocol. The
    // type is slice of
    // CISCOMEDIAGATEWAYMIB_Cmediagwcallcontrolconfigtable_Cmediagwcallcontrolconfigentry.
    Cmediagwcallcontrolconfigentry []CISCOMEDIAGATEWAYMIB_Cmediagwcallcontrolconfigtable_Cmediagwcallcontrolconfigentry
}

func (cmediagwcallcontrolconfigtable *CISCOMEDIAGATEWAYMIB_Cmediagwcallcontrolconfigtable) GetEntityData() *types.CommonEntityData {
    cmediagwcallcontrolconfigtable.EntityData.YFilter = cmediagwcallcontrolconfigtable.YFilter
    cmediagwcallcontrolconfigtable.EntityData.YangName = "cMediaGwCallControlConfigTable"
    cmediagwcallcontrolconfigtable.EntityData.BundleName = "cisco_ios_xe"
    cmediagwcallcontrolconfigtable.EntityData.ParentYangName = "CISCO-MEDIA-GATEWAY-MIB"
    cmediagwcallcontrolconfigtable.EntityData.SegmentPath = "cMediaGwCallControlConfigTable"
    cmediagwcallcontrolconfigtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cmediagwcallcontrolconfigtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cmediagwcallcontrolconfigtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cmediagwcallcontrolconfigtable.EntityData.Children = make(map[string]types.YChild)
    cmediagwcallcontrolconfigtable.EntityData.Children["cMediaGwCallControlConfigEntry"] = types.YChild{"Cmediagwcallcontrolconfigentry", nil}
    for i := range cmediagwcallcontrolconfigtable.Cmediagwcallcontrolconfigentry {
        cmediagwcallcontrolconfigtable.EntityData.Children[types.GetSegmentPath(&cmediagwcallcontrolconfigtable.Cmediagwcallcontrolconfigentry[i])] = types.YChild{"Cmediagwcallcontrolconfigentry", &cmediagwcallcontrolconfigtable.Cmediagwcallcontrolconfigentry[i]}
    }
    cmediagwcallcontrolconfigtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cmediagwcallcontrolconfigtable.EntityData)
}

// CISCOMEDIAGATEWAYMIB_Cmediagwcallcontrolconfigtable_Cmediagwcallcontrolconfigentry
// One entry for each media gateway which supports call control 
// protocol.
type CISCOMEDIAGATEWAYMIB_Cmediagwcallcontrolconfigtable_Cmediagwcallcontrolconfigentry struct {
    EntityData types.CommonEntityData
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

func (cmediagwcallcontrolconfigentry *CISCOMEDIAGATEWAYMIB_Cmediagwcallcontrolconfigtable_Cmediagwcallcontrolconfigentry) GetEntityData() *types.CommonEntityData {
    cmediagwcallcontrolconfigentry.EntityData.YFilter = cmediagwcallcontrolconfigentry.YFilter
    cmediagwcallcontrolconfigentry.EntityData.YangName = "cMediaGwCallControlConfigEntry"
    cmediagwcallcontrolconfigentry.EntityData.BundleName = "cisco_ios_xe"
    cmediagwcallcontrolconfigentry.EntityData.ParentYangName = "cMediaGwCallControlConfigTable"
    cmediagwcallcontrolconfigentry.EntityData.SegmentPath = "cMediaGwCallControlConfigEntry" + "[cmgwIndex='" + fmt.Sprintf("%v", cmediagwcallcontrolconfigentry.Cmgwindex) + "']"
    cmediagwcallcontrolconfigentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cmediagwcallcontrolconfigentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cmediagwcallcontrolconfigentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cmediagwcallcontrolconfigentry.EntityData.Children = make(map[string]types.YChild)
    cmediagwcallcontrolconfigentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cmediagwcallcontrolconfigentry.EntityData.Leafs["cmgwIndex"] = types.YLeaf{"Cmgwindex", cmediagwcallcontrolconfigentry.Cmgwindex}
    cmediagwcallcontrolconfigentry.EntityData.Leafs["cMediaGwCcCfgControlTos"] = types.YLeaf{"Cmediagwcccfgcontroltos", cmediagwcallcontrolconfigentry.Cmediagwcccfgcontroltos}
    cmediagwcallcontrolconfigentry.EntityData.Leafs["cMediaGwCcCfgBearerTos"] = types.YLeaf{"Cmediagwcccfgbearertos", cmediagwcallcontrolconfigentry.Cmediagwcccfgbearertos}
    cmediagwcallcontrolconfigentry.EntityData.Leafs["cMediaGwCcCfgNtePayload"] = types.YLeaf{"Cmediagwcccfgntepayload", cmediagwcallcontrolconfigentry.Cmediagwcccfgntepayload}
    cmediagwcallcontrolconfigentry.EntityData.Leafs["cMediaGwCcCfgNsePayload"] = types.YLeaf{"Cmediagwcccfgnsepayload", cmediagwcallcontrolconfigentry.Cmediagwcccfgnsepayload}
    cmediagwcallcontrolconfigentry.EntityData.Leafs["cMediaGwCcCfgNseRespTimer"] = types.YLeaf{"Cmediagwcccfgnseresptimer", cmediagwcallcontrolconfigentry.Cmediagwcccfgnseresptimer}
    cmediagwcallcontrolconfigentry.EntityData.Leafs["cMediaGwCcCfgVbdJitterDelayMode"] = types.YLeaf{"Cmediagwcccfgvbdjitterdelaymode", cmediagwcallcontrolconfigentry.Cmediagwcccfgvbdjitterdelaymode}
    cmediagwcallcontrolconfigentry.EntityData.Leafs["cMediaGwCcCfgVbdJitterMaxDelay"] = types.YLeaf{"Cmediagwcccfgvbdjittermaxdelay", cmediagwcallcontrolconfigentry.Cmediagwcccfgvbdjittermaxdelay}
    cmediagwcallcontrolconfigentry.EntityData.Leafs["cMediaGwCcCfgVbdJitterNomDelay"] = types.YLeaf{"Cmediagwcccfgvbdjitternomdelay", cmediagwcallcontrolconfigentry.Cmediagwcccfgvbdjitternomdelay}
    cmediagwcallcontrolconfigentry.EntityData.Leafs["cMediaGwCcCfgVbdJitterMinDelay"] = types.YLeaf{"Cmediagwcccfgvbdjittermindelay", cmediagwcallcontrolconfigentry.Cmediagwcccfgvbdjittermindelay}
    cmediagwcallcontrolconfigentry.EntityData.Leafs["cMediaGwCcCfgDefaultTonePlanId"] = types.YLeaf{"Cmediagwcccfgdefaulttoneplanid", cmediagwcallcontrolconfigentry.Cmediagwcccfgdefaulttoneplanid}
    cmediagwcallcontrolconfigentry.EntityData.Leafs["cMediaGwCcCfgDescrInfoEnabled"] = types.YLeaf{"Cmediagwcccfgdescrinfoenabled", cmediagwcallcontrolconfigentry.Cmediagwcccfgdescrinfoenabled}
    cmediagwcallcontrolconfigentry.EntityData.Leafs["cMediaGwCcCfgDsNamePrefix"] = types.YLeaf{"Cmediagwcccfgdsnameprefix", cmediagwcallcontrolconfigentry.Cmediagwcccfgdsnameprefix}
    cmediagwcallcontrolconfigentry.EntityData.Leafs["cMediaGwCcCfgRtpNamePrefix"] = types.YLeaf{"Cmediagwcccfgrtpnameprefix", cmediagwcallcontrolconfigentry.Cmediagwcccfgrtpnameprefix}
    cmediagwcallcontrolconfigentry.EntityData.Leafs["cMediaGwCcCfgAal1SvcNamePrefix"] = types.YLeaf{"Cmediagwcccfgaal1Svcnameprefix", cmediagwcallcontrolconfigentry.Cmediagwcccfgaal1Svcnameprefix}
    cmediagwcallcontrolconfigentry.EntityData.Leafs["cMediaGwCcCfgAal2SvcNamePrefix"] = types.YLeaf{"Cmediagwcccfgaal2Svcnameprefix", cmediagwcallcontrolconfigentry.Cmediagwcccfgaal2Svcnameprefix}
    cmediagwcallcontrolconfigentry.EntityData.Leafs["cMediaGwCcCfgClusterEnabled"] = types.YLeaf{"Cmediagwcccfgclusterenabled", cmediagwcallcontrolconfigentry.Cmediagwcccfgclusterenabled}
    cmediagwcallcontrolconfigentry.EntityData.Leafs["cMediaGwCcCfgDefBearerTraffic"] = types.YLeaf{"Cmediagwcccfgdefbearertraffic", cmediagwcallcontrolconfigentry.Cmediagwcccfgdefbearertraffic}
    cmediagwcallcontrolconfigentry.EntityData.Leafs["cMediaGwCcCfgDefRtpNamePrefix"] = types.YLeaf{"Cmediagwcccfgdefrtpnameprefix", cmediagwcallcontrolconfigentry.Cmediagwcccfgdefrtpnameprefix}
    return &(cmediagwcallcontrolconfigentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry stores the statistics information for a specific resource. The
    // type is slice of
    // CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable_Cmediagwrscstatsentry.
    Cmediagwrscstatsentry []CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable_Cmediagwrscstatsentry
}

func (cmediagwrscstatstable *CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable) GetEntityData() *types.CommonEntityData {
    cmediagwrscstatstable.EntityData.YFilter = cmediagwrscstatstable.YFilter
    cmediagwrscstatstable.EntityData.YangName = "cMediaGwRscStatsTable"
    cmediagwrscstatstable.EntityData.BundleName = "cisco_ios_xe"
    cmediagwrscstatstable.EntityData.ParentYangName = "CISCO-MEDIA-GATEWAY-MIB"
    cmediagwrscstatstable.EntityData.SegmentPath = "cMediaGwRscStatsTable"
    cmediagwrscstatstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cmediagwrscstatstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cmediagwrscstatstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cmediagwrscstatstable.EntityData.Children = make(map[string]types.YChild)
    cmediagwrscstatstable.EntityData.Children["cMediaGwRscStatsEntry"] = types.YChild{"Cmediagwrscstatsentry", nil}
    for i := range cmediagwrscstatstable.Cmediagwrscstatsentry {
        cmediagwrscstatstable.EntityData.Children[types.GetSegmentPath(&cmediagwrscstatstable.Cmediagwrscstatsentry[i])] = types.YChild{"Cmediagwrscstatsentry", &cmediagwrscstatstable.Cmediagwrscstatsentry[i]}
    }
    cmediagwrscstatstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cmediagwrscstatstable.EntityData)
}

// CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable_Cmediagwrscstatsentry
// Each entry stores the statistics
// information for a specific resource.
type CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable_Cmediagwrscstatsentry struct {
    EntityData types.CommonEntityData
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

func (cmediagwrscstatsentry *CISCOMEDIAGATEWAYMIB_Cmediagwrscstatstable_Cmediagwrscstatsentry) GetEntityData() *types.CommonEntityData {
    cmediagwrscstatsentry.EntityData.YFilter = cmediagwrscstatsentry.YFilter
    cmediagwrscstatsentry.EntityData.YangName = "cMediaGwRscStatsEntry"
    cmediagwrscstatsentry.EntityData.BundleName = "cisco_ios_xe"
    cmediagwrscstatsentry.EntityData.ParentYangName = "cMediaGwRscStatsTable"
    cmediagwrscstatsentry.EntityData.SegmentPath = "cMediaGwRscStatsEntry" + "[cmgwIndex='" + fmt.Sprintf("%v", cmediagwrscstatsentry.Cmgwindex) + "']" + "[cmgwRscStatsIndex='" + fmt.Sprintf("%v", cmediagwrscstatsentry.Cmgwrscstatsindex) + "']"
    cmediagwrscstatsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cmediagwrscstatsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cmediagwrscstatsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cmediagwrscstatsentry.EntityData.Children = make(map[string]types.YChild)
    cmediagwrscstatsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cmediagwrscstatsentry.EntityData.Leafs["cmgwIndex"] = types.YLeaf{"Cmgwindex", cmediagwrscstatsentry.Cmgwindex}
    cmediagwrscstatsentry.EntityData.Leafs["cmgwRscStatsIndex"] = types.YLeaf{"Cmgwrscstatsindex", cmediagwrscstatsentry.Cmgwrscstatsindex}
    cmediagwrscstatsentry.EntityData.Leafs["cmgwRscMaximumUtilization"] = types.YLeaf{"Cmgwrscmaximumutilization", cmediagwrscstatsentry.Cmgwrscmaximumutilization}
    cmediagwrscstatsentry.EntityData.Leafs["cmgwRscMinimumUtilization"] = types.YLeaf{"Cmgwrscminimumutilization", cmediagwrscstatsentry.Cmgwrscminimumutilization}
    cmediagwrscstatsentry.EntityData.Leafs["cmgwRscAverageUtilization"] = types.YLeaf{"Cmgwrscaverageutilization", cmediagwrscstatsentry.Cmgwrscaverageutilization}
    cmediagwrscstatsentry.EntityData.Leafs["cmgwRscSinceLastReset"] = types.YLeaf{"Cmgwrscsincelastreset", cmediagwrscstatsentry.Cmgwrscsincelastreset}
    return &(cmediagwrscstatsentry.EntityData)
}

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

