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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This table contains the global media gateway parameters information. It
    // supports the modification of the global media gateway  parameters.
    CMediaGwTable CISCOMEDIAGATEWAYMIB_CMediaGwTable

    // This table contains the available signaling protocols that are supported by
    // the media gateway for communication with MGCs.
    CmgwSignalProtocolTable CISCOMEDIAGATEWAYMIB_CmgwSignalProtocolTable

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
    CMediaGwIpConfigTable CISCOMEDIAGATEWAYMIB_CMediaGwIpConfigTable

    // This table provides the domain names which are configured by  users.  The
    // domain names can be used to represent IP addresses  for:     gateway    
    // External DNS name server     MGC (call agent) .
    CMediaGwDomainNameConfigTable CISCOMEDIAGATEWAYMIB_CMediaGwDomainNameConfigTable

    // There is only one DNS name server on a gateway and the domain name of the
    // DNS name server is put on  cMediaGwDomainNameConfigTable with 'dnsServer
    // (2)'.  There could be multi IP addresses are associated with the DNS name
    // server, this table is used to store these IP  addresses.  If any domain
    // name using external resolution, the last entry of this table is not allowed
    // to be deleted.
    CMediaGwDnsIpConfigTable CISCOMEDIAGATEWAYMIB_CMediaGwDnsIpConfigTable

    // This table is for managing LIF (Logical Interface)  in a media gateway.  
    // LIF is a logical interface which groups the TDM  DSx1s associated with a
    // set of packet resource partitions  (PVCs) in a media gateway.  LIF is used
    // for: 1. VoIP switching  2. VoATM switching .
    CmgwLifTable CISCOMEDIAGATEWAYMIB_CmgwLifTable

    // This table defines general call control attributes for the media gateway.
    CMediaGwCallControlConfigTable CISCOMEDIAGATEWAYMIB_CMediaGwCallControlConfigTable

    // This table stores the gateway resource statistics information.
    CMediaGwRscStatsTable CISCOMEDIAGATEWAYMIB_CMediaGwRscStatsTable
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

    cISCOMEDIAGATEWAYMIB.EntityData.Children = types.NewOrderedMap()
    cISCOMEDIAGATEWAYMIB.EntityData.Children.Append("cMediaGwTable", types.YChild{"CMediaGwTable", &cISCOMEDIAGATEWAYMIB.CMediaGwTable})
    cISCOMEDIAGATEWAYMIB.EntityData.Children.Append("cmgwSignalProtocolTable", types.YChild{"CmgwSignalProtocolTable", &cISCOMEDIAGATEWAYMIB.CmgwSignalProtocolTable})
    cISCOMEDIAGATEWAYMIB.EntityData.Children.Append("cMediaGwIpConfigTable", types.YChild{"CMediaGwIpConfigTable", &cISCOMEDIAGATEWAYMIB.CMediaGwIpConfigTable})
    cISCOMEDIAGATEWAYMIB.EntityData.Children.Append("cMediaGwDomainNameConfigTable", types.YChild{"CMediaGwDomainNameConfigTable", &cISCOMEDIAGATEWAYMIB.CMediaGwDomainNameConfigTable})
    cISCOMEDIAGATEWAYMIB.EntityData.Children.Append("cMediaGwDnsIpConfigTable", types.YChild{"CMediaGwDnsIpConfigTable", &cISCOMEDIAGATEWAYMIB.CMediaGwDnsIpConfigTable})
    cISCOMEDIAGATEWAYMIB.EntityData.Children.Append("cmgwLifTable", types.YChild{"CmgwLifTable", &cISCOMEDIAGATEWAYMIB.CmgwLifTable})
    cISCOMEDIAGATEWAYMIB.EntityData.Children.Append("cMediaGwCallControlConfigTable", types.YChild{"CMediaGwCallControlConfigTable", &cISCOMEDIAGATEWAYMIB.CMediaGwCallControlConfigTable})
    cISCOMEDIAGATEWAYMIB.EntityData.Children.Append("cMediaGwRscStatsTable", types.YChild{"CMediaGwRscStatsTable", &cISCOMEDIAGATEWAYMIB.CMediaGwRscStatsTable})
    cISCOMEDIAGATEWAYMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOMEDIAGATEWAYMIB.EntityData.YListKeys = []string {}

    return &(cISCOMEDIAGATEWAYMIB.EntityData)
}

// CISCOMEDIAGATEWAYMIB_CMediaGwTable
// This table contains the global media gateway parameters
// information.
// It supports the modification of the global media gateway 
// parameters.
type CISCOMEDIAGATEWAYMIB_CMediaGwTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A Media Gateway Entry.   At system power-up, an entry is created by the
    // agent  if the system detects a media gateway module has been added  to the
    // system, and an entry is deleted if the entry associated media gateway
    // module has been removed from the system. The type is slice of
    // CISCOMEDIAGATEWAYMIB_CMediaGwTable_CMediaGwEntry.
    CMediaGwEntry []*CISCOMEDIAGATEWAYMIB_CMediaGwTable_CMediaGwEntry
}

func (cMediaGwTable *CISCOMEDIAGATEWAYMIB_CMediaGwTable) GetEntityData() *types.CommonEntityData {
    cMediaGwTable.EntityData.YFilter = cMediaGwTable.YFilter
    cMediaGwTable.EntityData.YangName = "cMediaGwTable"
    cMediaGwTable.EntityData.BundleName = "cisco_ios_xe"
    cMediaGwTable.EntityData.ParentYangName = "CISCO-MEDIA-GATEWAY-MIB"
    cMediaGwTable.EntityData.SegmentPath = "cMediaGwTable"
    cMediaGwTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cMediaGwTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cMediaGwTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cMediaGwTable.EntityData.Children = types.NewOrderedMap()
    cMediaGwTable.EntityData.Children.Append("cMediaGwEntry", types.YChild{"CMediaGwEntry", nil})
    for i := range cMediaGwTable.CMediaGwEntry {
        cMediaGwTable.EntityData.Children.Append(types.GetSegmentPath(cMediaGwTable.CMediaGwEntry[i]), types.YChild{"CMediaGwEntry", cMediaGwTable.CMediaGwEntry[i]})
    }
    cMediaGwTable.EntityData.Leafs = types.NewOrderedMap()

    cMediaGwTable.EntityData.YListKeys = []string {}

    return &(cMediaGwTable.EntityData)
}

// CISCOMEDIAGATEWAYMIB_CMediaGwTable_CMediaGwEntry
// A Media Gateway Entry.  
// At system power-up, an entry is created by the agent 
// if the system detects a media gateway module has been added 
// to the system, and an entry is deleted if the entry associated
// media gateway module has been removed from the system.
type CISCOMEDIAGATEWAYMIB_CMediaGwTable_CMediaGwEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An index that uniquely identifies an entry in the 
    // cMediaGwTable. The type is interface{} with range: 1..2147483647.
    CmgwIndex interface{}

    // This object is used to represent a domain name under which    the Media
    // Gateway could also be registered in a DNS name server.   The value of this
    // object reflects the value of  cmgwConfigDomainName from the entry with a
    // value of  'gateway(1)' for object cmgwConfigDomainNameEntity of 
    // cMediaGwDomainNameConfigTable.  If there is no entry in
    // cMediaGwDomainNameConfigTable with 'gateway(1)' of
    // cmgwConfigDomainNameEntity, then the value of this object will be empty
    // string. The type is string with length: 0..64.
    CmgwDomainName interface{}

    // This object represents the entPhysicalIndex of the card in which media
    // gateway is running. It will contain value 0 if the entPhysicalIndex value
    // is not available or  not applicable. The type is interface{} with range:
    // 0..2147483647.
    CmgwPhysicalIndex interface{}

    // This object indicates the current service state of the Media  Gateway. This
    // object is controlled by 'cmgwAdminState'  object. The type is
    // CGwServiceState.
    CmgwServiceState interface{}

    // This object is used to change the service state of  the Media Gateway from
    // inService to outOfService or from  outOfService to inService.  The
    // resulting service state of the gateway is represented   by
    // 'cmgwServiceState'. The type is CGwAdminState.
    CmgwAdminState interface{}

    // This object is used to represent grace period. The grace period (restart
    // delay in RSIP message) is   expressed in a number of seconds.  It means how
    // soon the gateway will be taken out of service. The value -1 indicates that
    // the grace period time is disabled. The type is interface{} with range:
    // -1..65535. Units are seconds.
    CmgwGraceTime interface{}

    // This object is used to represent the VT (sonet Virtual Tributary) counting.
    // standard - standard counting (based on Bellcore TR253) titan    - TITAN5500
    // counting (based on Tellabs TITAN 5500)  Note: 'titan' is valid only if
    // sonet line medium type        (sonetMediumType of SONET-MIB) is 'sonet' and
    // sonet path payload type (cspSonetPathPayload of       CISCO-SONET-MIB) is
    // 'vt15vc11'. The type is CmgwVtMappingMode.
    CmgwVtMappingMode interface{}

    // This object is used to enable or disable the source IP and port filtering
    // with MGC for security consideration as follows:   'true'  - source IP and
    // port filter is enabled    'false' - source IP and port filter is disable .
    // The type is bool.
    CmgwSrcFilterEnabled interface{}

    // This object is used to enable or disable the lawful intercept for
    // government. as follows:   'true'  - enable lawful intercept   'false' -
    // disable lawful intercept. The type is bool.
    CmgwLawInterceptEnabled interface{}

    // This object is to enable or disable V23 tone. Setting the object value to
    // 'true', will cause VXSM (Voice Switching Service Module) to detect V23
    // tone. The type is bool.
    CmgwV23Enabled interface{}
}

func (cMediaGwEntry *CISCOMEDIAGATEWAYMIB_CMediaGwTable_CMediaGwEntry) GetEntityData() *types.CommonEntityData {
    cMediaGwEntry.EntityData.YFilter = cMediaGwEntry.YFilter
    cMediaGwEntry.EntityData.YangName = "cMediaGwEntry"
    cMediaGwEntry.EntityData.BundleName = "cisco_ios_xe"
    cMediaGwEntry.EntityData.ParentYangName = "cMediaGwTable"
    cMediaGwEntry.EntityData.SegmentPath = "cMediaGwEntry" + types.AddKeyToken(cMediaGwEntry.CmgwIndex, "cmgwIndex")
    cMediaGwEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cMediaGwEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cMediaGwEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cMediaGwEntry.EntityData.Children = types.NewOrderedMap()
    cMediaGwEntry.EntityData.Leafs = types.NewOrderedMap()
    cMediaGwEntry.EntityData.Leafs.Append("cmgwIndex", types.YLeaf{"CmgwIndex", cMediaGwEntry.CmgwIndex})
    cMediaGwEntry.EntityData.Leafs.Append("cmgwDomainName", types.YLeaf{"CmgwDomainName", cMediaGwEntry.CmgwDomainName})
    cMediaGwEntry.EntityData.Leafs.Append("cmgwPhysicalIndex", types.YLeaf{"CmgwPhysicalIndex", cMediaGwEntry.CmgwPhysicalIndex})
    cMediaGwEntry.EntityData.Leafs.Append("cmgwServiceState", types.YLeaf{"CmgwServiceState", cMediaGwEntry.CmgwServiceState})
    cMediaGwEntry.EntityData.Leafs.Append("cmgwAdminState", types.YLeaf{"CmgwAdminState", cMediaGwEntry.CmgwAdminState})
    cMediaGwEntry.EntityData.Leafs.Append("cmgwGraceTime", types.YLeaf{"CmgwGraceTime", cMediaGwEntry.CmgwGraceTime})
    cMediaGwEntry.EntityData.Leafs.Append("cmgwVtMappingMode", types.YLeaf{"CmgwVtMappingMode", cMediaGwEntry.CmgwVtMappingMode})
    cMediaGwEntry.EntityData.Leafs.Append("cmgwSrcFilterEnabled", types.YLeaf{"CmgwSrcFilterEnabled", cMediaGwEntry.CmgwSrcFilterEnabled})
    cMediaGwEntry.EntityData.Leafs.Append("cmgwLawInterceptEnabled", types.YLeaf{"CmgwLawInterceptEnabled", cMediaGwEntry.CmgwLawInterceptEnabled})
    cMediaGwEntry.EntityData.Leafs.Append("cmgwV23Enabled", types.YLeaf{"CmgwV23Enabled", cMediaGwEntry.CmgwV23Enabled})

    cMediaGwEntry.EntityData.YListKeys = []string {"CmgwIndex"}

    return &(cMediaGwEntry.EntityData)
}

// CISCOMEDIAGATEWAYMIB_CMediaGwTable_CMediaGwEntry_CmgwVtMappingMode represents       CISCO-SONET-MIB) is 'vt15vc11'.
type CISCOMEDIAGATEWAYMIB_CMediaGwTable_CMediaGwEntry_CmgwVtMappingMode string

const (
    CISCOMEDIAGATEWAYMIB_CMediaGwTable_CMediaGwEntry_CmgwVtMappingMode_standard CISCOMEDIAGATEWAYMIB_CMediaGwTable_CMediaGwEntry_CmgwVtMappingMode = "standard"

    CISCOMEDIAGATEWAYMIB_CMediaGwTable_CMediaGwEntry_CmgwVtMappingMode_titan CISCOMEDIAGATEWAYMIB_CMediaGwTable_CMediaGwEntry_CmgwVtMappingMode = "titan"
)

// CISCOMEDIAGATEWAYMIB_CmgwSignalProtocolTable
// This table contains the available signaling protocols that
// are supported by the media gateway for communication with
// MGCs.
type CISCOMEDIAGATEWAYMIB_CmgwSignalProtocolTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry represents an signaling protocol supported by the media gateway.
    // The type is slice of
    // CISCOMEDIAGATEWAYMIB_CmgwSignalProtocolTable_CmgwSignalProtocolEntry.
    CmgwSignalProtocolEntry []*CISCOMEDIAGATEWAYMIB_CmgwSignalProtocolTable_CmgwSignalProtocolEntry
}

func (cmgwSignalProtocolTable *CISCOMEDIAGATEWAYMIB_CmgwSignalProtocolTable) GetEntityData() *types.CommonEntityData {
    cmgwSignalProtocolTable.EntityData.YFilter = cmgwSignalProtocolTable.YFilter
    cmgwSignalProtocolTable.EntityData.YangName = "cmgwSignalProtocolTable"
    cmgwSignalProtocolTable.EntityData.BundleName = "cisco_ios_xe"
    cmgwSignalProtocolTable.EntityData.ParentYangName = "CISCO-MEDIA-GATEWAY-MIB"
    cmgwSignalProtocolTable.EntityData.SegmentPath = "cmgwSignalProtocolTable"
    cmgwSignalProtocolTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cmgwSignalProtocolTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cmgwSignalProtocolTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cmgwSignalProtocolTable.EntityData.Children = types.NewOrderedMap()
    cmgwSignalProtocolTable.EntityData.Children.Append("cmgwSignalProtocolEntry", types.YChild{"CmgwSignalProtocolEntry", nil})
    for i := range cmgwSignalProtocolTable.CmgwSignalProtocolEntry {
        cmgwSignalProtocolTable.EntityData.Children.Append(types.GetSegmentPath(cmgwSignalProtocolTable.CmgwSignalProtocolEntry[i]), types.YChild{"CmgwSignalProtocolEntry", cmgwSignalProtocolTable.CmgwSignalProtocolEntry[i]})
    }
    cmgwSignalProtocolTable.EntityData.Leafs = types.NewOrderedMap()

    cmgwSignalProtocolTable.EntityData.YListKeys = []string {}

    return &(cmgwSignalProtocolTable.EntityData)
}

// CISCOMEDIAGATEWAYMIB_CmgwSignalProtocolTable_CmgwSignalProtocolEntry
// Each entry represents an signaling protocol supported
// by the media gateway.
type CISCOMEDIAGATEWAYMIB_CmgwSignalProtocolTable_CmgwSignalProtocolEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // cisco_media_gateway_mib.CISCOMEDIAGATEWAYMIB_CMediaGwTable_CMediaGwEntry_CmgwIndex
    CmgwIndex interface{}

    // This attribute is a key. An index that uniquely identifies an entry in
    // cmgwSignalProtocolTable. The type is interface{} with range: 1..65535.
    CmgwSignalProtocolIndex interface{}

    // This object is used to represent the protocol type. other - None of the
    // following types. mgcp  - Media Gateway Control Protocol h248 - Media
    // Gateway Control (ITU H.248) tgcp - Trunking Gateway Control Protocol. The
    // type is CmgwSignalProtocol.
    CmgwSignalProtocol interface{}

    // This object is used to represent the protocol version.  For example
    // cmgwSignalProtocol is 'mgcp(2)' and this object is string '1.0'.
    // cmgwSignalProtocol is  'h248(3)' and this object is set to '2.0'. The type
    // is string with length: 1..16.
    CmgwSignalProtocolVersion interface{}

    // This object is used to represent the UDP port associated  with the
    // protocol. If the value of cmgwSignalProtocol is 'mgcp(2)' and the value of
    // cmgwSignalProtcolVersion is '1.0', the default value of this object is
    // '2727'.  If the value of cmgwSignalProtocol is 'h248(3)' and the value of
    // cmgwSignalProtcolVersion is '1.0', the default value of this object is
    // '2944'. The type is interface{} with range: 0..65535.
    CmgwSignalProtocolPort interface{}

    // This object specifies the protocol port of the Media Gateway Controller
    // (MGC). If the value of cmgwSignalProtocol is 'mgcp(2)' or 'tgcp(4)' and the
    // value of cmgwSignalProtcolVersion is '1.0', the default value of this
    // object is '2427'. If the value of cmgwSignalProtocol is 'h248(3)' and the
    // value of cmgwSignalProtcolVersion is '1.0', the default value of this
    // object is '2944'. The type is interface{} with range: 0..65535.
    CmgwSignalMgcProtocolPort interface{}

    // This object specifies the preference of the signal protocol  supported in
    // the media gateway.  If this object is set to 0, the corresponding signal
    // protocol will not be used by the gateway.   The value of this object is
    // unique within the corresponding gateway. The entry with lower value has
    // higher preference. The type is interface{} with range: 0..255.
    CmgwSignalProtocolPreference interface{}

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
    CmgwSignalProtocolConfigVer interface{}
}

func (cmgwSignalProtocolEntry *CISCOMEDIAGATEWAYMIB_CmgwSignalProtocolTable_CmgwSignalProtocolEntry) GetEntityData() *types.CommonEntityData {
    cmgwSignalProtocolEntry.EntityData.YFilter = cmgwSignalProtocolEntry.YFilter
    cmgwSignalProtocolEntry.EntityData.YangName = "cmgwSignalProtocolEntry"
    cmgwSignalProtocolEntry.EntityData.BundleName = "cisco_ios_xe"
    cmgwSignalProtocolEntry.EntityData.ParentYangName = "cmgwSignalProtocolTable"
    cmgwSignalProtocolEntry.EntityData.SegmentPath = "cmgwSignalProtocolEntry" + types.AddKeyToken(cmgwSignalProtocolEntry.CmgwIndex, "cmgwIndex") + types.AddKeyToken(cmgwSignalProtocolEntry.CmgwSignalProtocolIndex, "cmgwSignalProtocolIndex")
    cmgwSignalProtocolEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cmgwSignalProtocolEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cmgwSignalProtocolEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cmgwSignalProtocolEntry.EntityData.Children = types.NewOrderedMap()
    cmgwSignalProtocolEntry.EntityData.Leafs = types.NewOrderedMap()
    cmgwSignalProtocolEntry.EntityData.Leafs.Append("cmgwIndex", types.YLeaf{"CmgwIndex", cmgwSignalProtocolEntry.CmgwIndex})
    cmgwSignalProtocolEntry.EntityData.Leafs.Append("cmgwSignalProtocolIndex", types.YLeaf{"CmgwSignalProtocolIndex", cmgwSignalProtocolEntry.CmgwSignalProtocolIndex})
    cmgwSignalProtocolEntry.EntityData.Leafs.Append("cmgwSignalProtocol", types.YLeaf{"CmgwSignalProtocol", cmgwSignalProtocolEntry.CmgwSignalProtocol})
    cmgwSignalProtocolEntry.EntityData.Leafs.Append("cmgwSignalProtocolVersion", types.YLeaf{"CmgwSignalProtocolVersion", cmgwSignalProtocolEntry.CmgwSignalProtocolVersion})
    cmgwSignalProtocolEntry.EntityData.Leafs.Append("cmgwSignalProtocolPort", types.YLeaf{"CmgwSignalProtocolPort", cmgwSignalProtocolEntry.CmgwSignalProtocolPort})
    cmgwSignalProtocolEntry.EntityData.Leafs.Append("cmgwSignalMgcProtocolPort", types.YLeaf{"CmgwSignalMgcProtocolPort", cmgwSignalProtocolEntry.CmgwSignalMgcProtocolPort})
    cmgwSignalProtocolEntry.EntityData.Leafs.Append("cmgwSignalProtocolPreference", types.YLeaf{"CmgwSignalProtocolPreference", cmgwSignalProtocolEntry.CmgwSignalProtocolPreference})
    cmgwSignalProtocolEntry.EntityData.Leafs.Append("cmgwSignalProtocolConfigVer", types.YLeaf{"CmgwSignalProtocolConfigVer", cmgwSignalProtocolEntry.CmgwSignalProtocolConfigVer})

    cmgwSignalProtocolEntry.EntityData.YListKeys = []string {"CmgwIndex", "CmgwSignalProtocolIndex"}

    return &(cmgwSignalProtocolEntry.EntityData)
}

// CISCOMEDIAGATEWAYMIB_CmgwSignalProtocolTable_CmgwSignalProtocolEntry_CmgwSignalProtocol represents tgcp - Trunking Gateway Control Protocol
type CISCOMEDIAGATEWAYMIB_CmgwSignalProtocolTable_CmgwSignalProtocolEntry_CmgwSignalProtocol string

const (
    CISCOMEDIAGATEWAYMIB_CmgwSignalProtocolTable_CmgwSignalProtocolEntry_CmgwSignalProtocol_other CISCOMEDIAGATEWAYMIB_CmgwSignalProtocolTable_CmgwSignalProtocolEntry_CmgwSignalProtocol = "other"

    CISCOMEDIAGATEWAYMIB_CmgwSignalProtocolTable_CmgwSignalProtocolEntry_CmgwSignalProtocol_mgcp CISCOMEDIAGATEWAYMIB_CmgwSignalProtocolTable_CmgwSignalProtocolEntry_CmgwSignalProtocol = "mgcp"

    CISCOMEDIAGATEWAYMIB_CmgwSignalProtocolTable_CmgwSignalProtocolEntry_CmgwSignalProtocol_h248 CISCOMEDIAGATEWAYMIB_CmgwSignalProtocolTable_CmgwSignalProtocolEntry_CmgwSignalProtocol = "h248"

    CISCOMEDIAGATEWAYMIB_CmgwSignalProtocolTable_CmgwSignalProtocolEntry_CmgwSignalProtocol_tgcp CISCOMEDIAGATEWAYMIB_CmgwSignalProtocolTable_CmgwSignalProtocolEntry_CmgwSignalProtocol = "tgcp"
)

// CISCOMEDIAGATEWAYMIB_CMediaGwIpConfigTable
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
type CISCOMEDIAGATEWAYMIB_CMediaGwIpConfigTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A Media Gateway IP configuration entry.  Each entry represents a media
    // gateway IP address for MGCs to communicate with the media gateway. The type
    // is slice of
    // CISCOMEDIAGATEWAYMIB_CMediaGwIpConfigTable_CMediaGwIpConfigEntry.
    CMediaGwIpConfigEntry []*CISCOMEDIAGATEWAYMIB_CMediaGwIpConfigTable_CMediaGwIpConfigEntry
}

func (cMediaGwIpConfigTable *CISCOMEDIAGATEWAYMIB_CMediaGwIpConfigTable) GetEntityData() *types.CommonEntityData {
    cMediaGwIpConfigTable.EntityData.YFilter = cMediaGwIpConfigTable.YFilter
    cMediaGwIpConfigTable.EntityData.YangName = "cMediaGwIpConfigTable"
    cMediaGwIpConfigTable.EntityData.BundleName = "cisco_ios_xe"
    cMediaGwIpConfigTable.EntityData.ParentYangName = "CISCO-MEDIA-GATEWAY-MIB"
    cMediaGwIpConfigTable.EntityData.SegmentPath = "cMediaGwIpConfigTable"
    cMediaGwIpConfigTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cMediaGwIpConfigTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cMediaGwIpConfigTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cMediaGwIpConfigTable.EntityData.Children = types.NewOrderedMap()
    cMediaGwIpConfigTable.EntityData.Children.Append("cMediaGwIpConfigEntry", types.YChild{"CMediaGwIpConfigEntry", nil})
    for i := range cMediaGwIpConfigTable.CMediaGwIpConfigEntry {
        cMediaGwIpConfigTable.EntityData.Children.Append(types.GetSegmentPath(cMediaGwIpConfigTable.CMediaGwIpConfigEntry[i]), types.YChild{"CMediaGwIpConfigEntry", cMediaGwIpConfigTable.CMediaGwIpConfigEntry[i]})
    }
    cMediaGwIpConfigTable.EntityData.Leafs = types.NewOrderedMap()

    cMediaGwIpConfigTable.EntityData.YListKeys = []string {}

    return &(cMediaGwIpConfigTable.EntityData)
}

// CISCOMEDIAGATEWAYMIB_CMediaGwIpConfigTable_CMediaGwIpConfigEntry
// A Media Gateway IP configuration entry. 
// Each entry represents a media gateway IP address for MGCs
// to communicate with the media gateway.
type CISCOMEDIAGATEWAYMIB_CMediaGwIpConfigTable_CMediaGwIpConfigEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // cisco_media_gateway_mib.CISCOMEDIAGATEWAYMIB_CMediaGwTable_CMediaGwEntry_CmgwIndex
    CmgwIndex interface{}

    // This attribute is a key. A unique index to identify each media gateway IP
    // address. The type is interface{} with range: 1..64.
    CmgwIpConfigIndex interface{}

    // This object is ifIndex of the interface which is associated to the media
    // gateway IP address.  For ATM interface, the IP address should be associated
    // to an existing PVC:    cmgwIpConfigIfIndex represents port of the PVC   
    // cmgwIpConfigVpi represents VPI of the PVC    cmgwIpConfigVci represents VCI
    // of the PVC And one PVC only can be associated with one IP address.  If this
    // object is set to zero which means the IP address is not associated to any
    // interface. The type is interface{} with range: 0..2147483647.
    CmgwIpConfigIfIndex interface{}

    // This object represents VPI of the PVC which is associated to the IP
    // address. If the IP address is not associated to PVC, the value  of this
    // object is set to -1. The type is interface{} with range: -1..4095.
    CmgwIpConfigVpi interface{}

    // This object represents VCI of the PVC which is associated to the IP
    // address. If the IP address is not associated to PVC, the value of this
    // object is set to -1. The type is interface{} with range: -1..65535.
    CmgwIpConfigVci interface{}

    // This object is the IP address type. The type is InetAddressType.
    CmgwIpConfigAddrType interface{}

    // The configured IP address of media gateway. This object can not be
    // modified. The type is string with length: 0..255.
    CmgwIpConfigAddress interface{}

    // This object is used to specify the number of leading one    bits which from
    // the mask to be logical-ANDed with the media   gateway address before being
    // compared to the value in the  cmgwIpCofigAddress.  Any assignment (implicit
    // or otherwise) of an instance of this object to a value x must be rejected
    // if the bitwise logical-AND of the mask formed from x with the value  of the
    // corresponding instance of the cmgwIpCofigAddress  object is not equal to
    // cmgwIpCofigAddress. The type is interface{} with range: 0..2040.
    CmgwIpConfigSubnetMask interface{}

    // This object specifies cmgwIpConfigAddress of the entry will become the
    // default gateway address. This object can be set to 'true' for only one
    // entry in the table. The type is bool.
    CmgwIpConfigDefaultGwIp interface{}

    // This object specifies whether the address defined in cmgwIpConfigAddress is
    // the address mapping at the remote end of this PVC.   If this object is set
    // to 'true', the address defined in cmgwIpConfigAddress is for the remote end
    // of the PVC. If this object is set to 'false', the address defined in
    // cmgwIpConfigAddress is for the local end of the PVC. The type is bool.
    CmgwIpConfigForRemoteMapping interface{}

    // This object is used to add and delete an entry.  When an entry of the table
    // is created, the following  objects are mandatory:     cmgwIpConfigIfIndex  
    // cmgwIpConfigVpi     cmgwIpConfigVci     cmgwIpConfigAddress    
    // cmgwIpConfigSubnetMask  These objects can not be modified after the value
    // of this object is set to 'active'.  Modification can only be done by
    // deleting and re-adding the  entry again.  After the system verify the
    // validity of the data, it will set the cmgwIpConfigRowStatus to 'active'.
    // The type is RowStatus.
    CmgwIpConfigRowStatus interface{}
}

func (cMediaGwIpConfigEntry *CISCOMEDIAGATEWAYMIB_CMediaGwIpConfigTable_CMediaGwIpConfigEntry) GetEntityData() *types.CommonEntityData {
    cMediaGwIpConfigEntry.EntityData.YFilter = cMediaGwIpConfigEntry.YFilter
    cMediaGwIpConfigEntry.EntityData.YangName = "cMediaGwIpConfigEntry"
    cMediaGwIpConfigEntry.EntityData.BundleName = "cisco_ios_xe"
    cMediaGwIpConfigEntry.EntityData.ParentYangName = "cMediaGwIpConfigTable"
    cMediaGwIpConfigEntry.EntityData.SegmentPath = "cMediaGwIpConfigEntry" + types.AddKeyToken(cMediaGwIpConfigEntry.CmgwIndex, "cmgwIndex") + types.AddKeyToken(cMediaGwIpConfigEntry.CmgwIpConfigIndex, "cmgwIpConfigIndex")
    cMediaGwIpConfigEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cMediaGwIpConfigEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cMediaGwIpConfigEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cMediaGwIpConfigEntry.EntityData.Children = types.NewOrderedMap()
    cMediaGwIpConfigEntry.EntityData.Leafs = types.NewOrderedMap()
    cMediaGwIpConfigEntry.EntityData.Leafs.Append("cmgwIndex", types.YLeaf{"CmgwIndex", cMediaGwIpConfigEntry.CmgwIndex})
    cMediaGwIpConfigEntry.EntityData.Leafs.Append("cmgwIpConfigIndex", types.YLeaf{"CmgwIpConfigIndex", cMediaGwIpConfigEntry.CmgwIpConfigIndex})
    cMediaGwIpConfigEntry.EntityData.Leafs.Append("cmgwIpConfigIfIndex", types.YLeaf{"CmgwIpConfigIfIndex", cMediaGwIpConfigEntry.CmgwIpConfigIfIndex})
    cMediaGwIpConfigEntry.EntityData.Leafs.Append("cmgwIpConfigVpi", types.YLeaf{"CmgwIpConfigVpi", cMediaGwIpConfigEntry.CmgwIpConfigVpi})
    cMediaGwIpConfigEntry.EntityData.Leafs.Append("cmgwIpConfigVci", types.YLeaf{"CmgwIpConfigVci", cMediaGwIpConfigEntry.CmgwIpConfigVci})
    cMediaGwIpConfigEntry.EntityData.Leafs.Append("cmgwIpConfigAddrType", types.YLeaf{"CmgwIpConfigAddrType", cMediaGwIpConfigEntry.CmgwIpConfigAddrType})
    cMediaGwIpConfigEntry.EntityData.Leafs.Append("cmgwIpConfigAddress", types.YLeaf{"CmgwIpConfigAddress", cMediaGwIpConfigEntry.CmgwIpConfigAddress})
    cMediaGwIpConfigEntry.EntityData.Leafs.Append("cmgwIpConfigSubnetMask", types.YLeaf{"CmgwIpConfigSubnetMask", cMediaGwIpConfigEntry.CmgwIpConfigSubnetMask})
    cMediaGwIpConfigEntry.EntityData.Leafs.Append("cmgwIpConfigDefaultGwIp", types.YLeaf{"CmgwIpConfigDefaultGwIp", cMediaGwIpConfigEntry.CmgwIpConfigDefaultGwIp})
    cMediaGwIpConfigEntry.EntityData.Leafs.Append("cmgwIpConfigForRemoteMapping", types.YLeaf{"CmgwIpConfigForRemoteMapping", cMediaGwIpConfigEntry.CmgwIpConfigForRemoteMapping})
    cMediaGwIpConfigEntry.EntityData.Leafs.Append("cmgwIpConfigRowStatus", types.YLeaf{"CmgwIpConfigRowStatus", cMediaGwIpConfigEntry.CmgwIpConfigRowStatus})

    cMediaGwIpConfigEntry.EntityData.YListKeys = []string {"CmgwIndex", "CmgwIpConfigIndex"}

    return &(cMediaGwIpConfigEntry.EntityData)
}

// CISCOMEDIAGATEWAYMIB_CMediaGwDomainNameConfigTable
// This table provides the domain names which are configured by 
// users. 
// The domain names can be used to represent IP addresses 
// for:
//     gateway
//     External DNS name server
//     MGC (call agent) 
type CISCOMEDIAGATEWAYMIB_CMediaGwDomainNameConfigTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry represents a domain name used in the system.  Creation and
    // deletion are supported. Modification is prohibited. The type is slice of
    // CISCOMEDIAGATEWAYMIB_CMediaGwDomainNameConfigTable_CMediaGwDomainNameConfigEntry.
    CMediaGwDomainNameConfigEntry []*CISCOMEDIAGATEWAYMIB_CMediaGwDomainNameConfigTable_CMediaGwDomainNameConfigEntry
}

func (cMediaGwDomainNameConfigTable *CISCOMEDIAGATEWAYMIB_CMediaGwDomainNameConfigTable) GetEntityData() *types.CommonEntityData {
    cMediaGwDomainNameConfigTable.EntityData.YFilter = cMediaGwDomainNameConfigTable.YFilter
    cMediaGwDomainNameConfigTable.EntityData.YangName = "cMediaGwDomainNameConfigTable"
    cMediaGwDomainNameConfigTable.EntityData.BundleName = "cisco_ios_xe"
    cMediaGwDomainNameConfigTable.EntityData.ParentYangName = "CISCO-MEDIA-GATEWAY-MIB"
    cMediaGwDomainNameConfigTable.EntityData.SegmentPath = "cMediaGwDomainNameConfigTable"
    cMediaGwDomainNameConfigTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cMediaGwDomainNameConfigTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cMediaGwDomainNameConfigTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cMediaGwDomainNameConfigTable.EntityData.Children = types.NewOrderedMap()
    cMediaGwDomainNameConfigTable.EntityData.Children.Append("cMediaGwDomainNameConfigEntry", types.YChild{"CMediaGwDomainNameConfigEntry", nil})
    for i := range cMediaGwDomainNameConfigTable.CMediaGwDomainNameConfigEntry {
        cMediaGwDomainNameConfigTable.EntityData.Children.Append(types.GetSegmentPath(cMediaGwDomainNameConfigTable.CMediaGwDomainNameConfigEntry[i]), types.YChild{"CMediaGwDomainNameConfigEntry", cMediaGwDomainNameConfigTable.CMediaGwDomainNameConfigEntry[i]})
    }
    cMediaGwDomainNameConfigTable.EntityData.Leafs = types.NewOrderedMap()

    cMediaGwDomainNameConfigTable.EntityData.YListKeys = []string {}

    return &(cMediaGwDomainNameConfigTable.EntityData)
}

// CISCOMEDIAGATEWAYMIB_CMediaGwDomainNameConfigTable_CMediaGwDomainNameConfigEntry
// Each entry represents a domain name used in the system.
// 
// Creation and deletion are supported. Modification
// is prohibited.
type CISCOMEDIAGATEWAYMIB_CMediaGwDomainNameConfigTable_CMediaGwDomainNameConfigEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // cisco_media_gateway_mib.CISCOMEDIAGATEWAYMIB_CMediaGwTable_CMediaGwEntry_CmgwIndex
    CmgwIndex interface{}

    // This attribute is a key. An index that is uniquely identifies a domain name
    // configured in the system. The type is interface{} with range: 1..128.
    CmgwConfigDomainNameIndex interface{}

    // This object indicates which entity to use this domain name.  gateway(1)   -
    // The domain name of media gateway.                With the same cmgwIndex,
    // there is one and                 only one entry allowed with the value     
    // 'gateway(1)' of this object.  dnsServer(2) - The domain name of DNS name
    // server that is used                 by Media gateway to find Internet
    // Network                 Address from a DNS name.  mgc(3)       - The domain
    // name of a MGC (Media Gateway                Controller) associated with the
    // media                 gateway. . The type is CmgwConfigDomainNameEntity.
    CmgwConfigDomainNameEntity interface{}

    // This object specifies the domain name.  The domain name should be unique if
    // there are more than one entries having the same value in the object 
    // cmgwConfigDomainNameEntity. For example, the gateway domain name should be
    // unique  if the cmgwConfigDomainNameEntity has the value of  'gateway(1)'.
    // The type is string with length: 1..64.
    CmgwConfigDomainName interface{}

    // This object is used to add and delete an entry.  When an entry is created,
    // the following objects are mandatory:      cmgwConfigDomainName     
    // cmgwConfigDomainNameEntity  When deleting domain name of DNS name server
    // (cmgwConfigDomainNameEntity is dnsServer (2)), the 
    // cMediaGwDnsIpConfigTable should be empty.  Adding/deleting entry with
    // cmgwConfigDomainNameEntity of 'mgc' will cause adding/deleting entry in 
    // cMgcConfigTable (CISCO-MGC-MIB) automatically.  The cmgwConfigDomainName
    // and cmgwConfigDomainNameEntity can not be modified if the value of this
    // object is 'active'. . The type is RowStatus.
    CmgwConfigDomainNameRowStatus interface{}
}

func (cMediaGwDomainNameConfigEntry *CISCOMEDIAGATEWAYMIB_CMediaGwDomainNameConfigTable_CMediaGwDomainNameConfigEntry) GetEntityData() *types.CommonEntityData {
    cMediaGwDomainNameConfigEntry.EntityData.YFilter = cMediaGwDomainNameConfigEntry.YFilter
    cMediaGwDomainNameConfigEntry.EntityData.YangName = "cMediaGwDomainNameConfigEntry"
    cMediaGwDomainNameConfigEntry.EntityData.BundleName = "cisco_ios_xe"
    cMediaGwDomainNameConfigEntry.EntityData.ParentYangName = "cMediaGwDomainNameConfigTable"
    cMediaGwDomainNameConfigEntry.EntityData.SegmentPath = "cMediaGwDomainNameConfigEntry" + types.AddKeyToken(cMediaGwDomainNameConfigEntry.CmgwIndex, "cmgwIndex") + types.AddKeyToken(cMediaGwDomainNameConfigEntry.CmgwConfigDomainNameIndex, "cmgwConfigDomainNameIndex")
    cMediaGwDomainNameConfigEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cMediaGwDomainNameConfigEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cMediaGwDomainNameConfigEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cMediaGwDomainNameConfigEntry.EntityData.Children = types.NewOrderedMap()
    cMediaGwDomainNameConfigEntry.EntityData.Leafs = types.NewOrderedMap()
    cMediaGwDomainNameConfigEntry.EntityData.Leafs.Append("cmgwIndex", types.YLeaf{"CmgwIndex", cMediaGwDomainNameConfigEntry.CmgwIndex})
    cMediaGwDomainNameConfigEntry.EntityData.Leafs.Append("cmgwConfigDomainNameIndex", types.YLeaf{"CmgwConfigDomainNameIndex", cMediaGwDomainNameConfigEntry.CmgwConfigDomainNameIndex})
    cMediaGwDomainNameConfigEntry.EntityData.Leafs.Append("cmgwConfigDomainNameEntity", types.YLeaf{"CmgwConfigDomainNameEntity", cMediaGwDomainNameConfigEntry.CmgwConfigDomainNameEntity})
    cMediaGwDomainNameConfigEntry.EntityData.Leafs.Append("cmgwConfigDomainName", types.YLeaf{"CmgwConfigDomainName", cMediaGwDomainNameConfigEntry.CmgwConfigDomainName})
    cMediaGwDomainNameConfigEntry.EntityData.Leafs.Append("cmgwConfigDomainNameRowStatus", types.YLeaf{"CmgwConfigDomainNameRowStatus", cMediaGwDomainNameConfigEntry.CmgwConfigDomainNameRowStatus})

    cMediaGwDomainNameConfigEntry.EntityData.YListKeys = []string {"CmgwIndex", "CmgwConfigDomainNameIndex"}

    return &(cMediaGwDomainNameConfigEntry.EntityData)
}

// CISCOMEDIAGATEWAYMIB_CMediaGwDomainNameConfigTable_CMediaGwDomainNameConfigEntry_CmgwConfigDomainNameEntity represents                gateway. 
type CISCOMEDIAGATEWAYMIB_CMediaGwDomainNameConfigTable_CMediaGwDomainNameConfigEntry_CmgwConfigDomainNameEntity string

const (
    CISCOMEDIAGATEWAYMIB_CMediaGwDomainNameConfigTable_CMediaGwDomainNameConfigEntry_CmgwConfigDomainNameEntity_gateway CISCOMEDIAGATEWAYMIB_CMediaGwDomainNameConfigTable_CMediaGwDomainNameConfigEntry_CmgwConfigDomainNameEntity = "gateway"

    CISCOMEDIAGATEWAYMIB_CMediaGwDomainNameConfigTable_CMediaGwDomainNameConfigEntry_CmgwConfigDomainNameEntity_dnsServer CISCOMEDIAGATEWAYMIB_CMediaGwDomainNameConfigTable_CMediaGwDomainNameConfigEntry_CmgwConfigDomainNameEntity = "dnsServer"

    CISCOMEDIAGATEWAYMIB_CMediaGwDomainNameConfigTable_CMediaGwDomainNameConfigEntry_CmgwConfigDomainNameEntity_mgc CISCOMEDIAGATEWAYMIB_CMediaGwDomainNameConfigTable_CMediaGwDomainNameConfigEntry_CmgwConfigDomainNameEntity = "mgc"
)

// CISCOMEDIAGATEWAYMIB_CMediaGwDnsIpConfigTable
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
type CISCOMEDIAGATEWAYMIB_CMediaGwDnsIpConfigTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry represents an IP address of the DNS name  server. The type is
    // slice of
    // CISCOMEDIAGATEWAYMIB_CMediaGwDnsIpConfigTable_CMediaGwDnsIpConfigEntry.
    CMediaGwDnsIpConfigEntry []*CISCOMEDIAGATEWAYMIB_CMediaGwDnsIpConfigTable_CMediaGwDnsIpConfigEntry
}

func (cMediaGwDnsIpConfigTable *CISCOMEDIAGATEWAYMIB_CMediaGwDnsIpConfigTable) GetEntityData() *types.CommonEntityData {
    cMediaGwDnsIpConfigTable.EntityData.YFilter = cMediaGwDnsIpConfigTable.YFilter
    cMediaGwDnsIpConfigTable.EntityData.YangName = "cMediaGwDnsIpConfigTable"
    cMediaGwDnsIpConfigTable.EntityData.BundleName = "cisco_ios_xe"
    cMediaGwDnsIpConfigTable.EntityData.ParentYangName = "CISCO-MEDIA-GATEWAY-MIB"
    cMediaGwDnsIpConfigTable.EntityData.SegmentPath = "cMediaGwDnsIpConfigTable"
    cMediaGwDnsIpConfigTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cMediaGwDnsIpConfigTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cMediaGwDnsIpConfigTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cMediaGwDnsIpConfigTable.EntityData.Children = types.NewOrderedMap()
    cMediaGwDnsIpConfigTable.EntityData.Children.Append("cMediaGwDnsIpConfigEntry", types.YChild{"CMediaGwDnsIpConfigEntry", nil})
    for i := range cMediaGwDnsIpConfigTable.CMediaGwDnsIpConfigEntry {
        cMediaGwDnsIpConfigTable.EntityData.Children.Append(types.GetSegmentPath(cMediaGwDnsIpConfigTable.CMediaGwDnsIpConfigEntry[i]), types.YChild{"CMediaGwDnsIpConfigEntry", cMediaGwDnsIpConfigTable.CMediaGwDnsIpConfigEntry[i]})
    }
    cMediaGwDnsIpConfigTable.EntityData.Leafs = types.NewOrderedMap()

    cMediaGwDnsIpConfigTable.EntityData.YListKeys = []string {}

    return &(cMediaGwDnsIpConfigTable.EntityData)
}

// CISCOMEDIAGATEWAYMIB_CMediaGwDnsIpConfigTable_CMediaGwDnsIpConfigEntry
// Each entry represents an IP address of the DNS name 
// server.
type CISCOMEDIAGATEWAYMIB_CMediaGwDnsIpConfigTable_CMediaGwDnsIpConfigEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // cisco_media_gateway_mib.CISCOMEDIAGATEWAYMIB_CMediaGwTable_CMediaGwEntry_CmgwIndex
    CmgwIndex interface{}

    // This attribute is a key. An index that uniquely identifies an IP address of
    // DNS name server. The type is interface{} with range: 1..6.
    CmgwDnsIpIndex interface{}

    // The domain name of DNS name server.  The value of this object reflects the
    // value of  cmgwConfigDomainName from the entry with a value of 
    // 'dnsServer(2)' for object cmgwConfigDomainNameEntity of 
    // cMediaGwDomainNameConfigTable.  If there is no entry in
    // cMediaGwDomainNameConfigTable with 'dnsServer(2)' of
    // cmgwConfigDomainNameEntity, then the value of this object will be empty
    // string. The type is string with length: 1..64.
    CmgwDnsDomainName interface{}

    // DNS name server IP address type. The type is InetAddressType.
    CmgwDnsIpType interface{}

    // The IP address of DNS name server. The IP address of DNS name server must
    // be unique in this table. The type is string with length: 0..255.
    CmgwDnsIp interface{}

    // This object is used to add and delete an entry.  When an entry of the table
    // is created, the value of this object should be set to 'createAndGo' and the
    // following objects are mandatory:     cmgwDnsIp  When the user wants to
    // delete the entry, the value of this object should be set to 'destroy'.  The
    // entry can not be modified if the value of this  object is 'active'. The
    // type is RowStatus.
    CmgwDnsIpRowStatus interface{}
}

func (cMediaGwDnsIpConfigEntry *CISCOMEDIAGATEWAYMIB_CMediaGwDnsIpConfigTable_CMediaGwDnsIpConfigEntry) GetEntityData() *types.CommonEntityData {
    cMediaGwDnsIpConfigEntry.EntityData.YFilter = cMediaGwDnsIpConfigEntry.YFilter
    cMediaGwDnsIpConfigEntry.EntityData.YangName = "cMediaGwDnsIpConfigEntry"
    cMediaGwDnsIpConfigEntry.EntityData.BundleName = "cisco_ios_xe"
    cMediaGwDnsIpConfigEntry.EntityData.ParentYangName = "cMediaGwDnsIpConfigTable"
    cMediaGwDnsIpConfigEntry.EntityData.SegmentPath = "cMediaGwDnsIpConfigEntry" + types.AddKeyToken(cMediaGwDnsIpConfigEntry.CmgwIndex, "cmgwIndex") + types.AddKeyToken(cMediaGwDnsIpConfigEntry.CmgwDnsIpIndex, "cmgwDnsIpIndex")
    cMediaGwDnsIpConfigEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cMediaGwDnsIpConfigEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cMediaGwDnsIpConfigEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cMediaGwDnsIpConfigEntry.EntityData.Children = types.NewOrderedMap()
    cMediaGwDnsIpConfigEntry.EntityData.Leafs = types.NewOrderedMap()
    cMediaGwDnsIpConfigEntry.EntityData.Leafs.Append("cmgwIndex", types.YLeaf{"CmgwIndex", cMediaGwDnsIpConfigEntry.CmgwIndex})
    cMediaGwDnsIpConfigEntry.EntityData.Leafs.Append("cmgwDnsIpIndex", types.YLeaf{"CmgwDnsIpIndex", cMediaGwDnsIpConfigEntry.CmgwDnsIpIndex})
    cMediaGwDnsIpConfigEntry.EntityData.Leafs.Append("cmgwDnsDomainName", types.YLeaf{"CmgwDnsDomainName", cMediaGwDnsIpConfigEntry.CmgwDnsDomainName})
    cMediaGwDnsIpConfigEntry.EntityData.Leafs.Append("cmgwDnsIpType", types.YLeaf{"CmgwDnsIpType", cMediaGwDnsIpConfigEntry.CmgwDnsIpType})
    cMediaGwDnsIpConfigEntry.EntityData.Leafs.Append("cmgwDnsIp", types.YLeaf{"CmgwDnsIp", cMediaGwDnsIpConfigEntry.CmgwDnsIp})
    cMediaGwDnsIpConfigEntry.EntityData.Leafs.Append("cmgwDnsIpRowStatus", types.YLeaf{"CmgwDnsIpRowStatus", cMediaGwDnsIpConfigEntry.CmgwDnsIpRowStatus})

    cMediaGwDnsIpConfigEntry.EntityData.YListKeys = []string {"CmgwIndex", "CmgwDnsIpIndex"}

    return &(cMediaGwDnsIpConfigEntry.EntityData)
}

// CISCOMEDIAGATEWAYMIB_CmgwLifTable
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
type CISCOMEDIAGATEWAYMIB_CmgwLifTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry of this table is created by the media gateway when it supports the
    // VoIP/VoATM application. The type is slice of
    // CISCOMEDIAGATEWAYMIB_CmgwLifTable_CmgwLifEntry.
    CmgwLifEntry []*CISCOMEDIAGATEWAYMIB_CmgwLifTable_CmgwLifEntry
}

func (cmgwLifTable *CISCOMEDIAGATEWAYMIB_CmgwLifTable) GetEntityData() *types.CommonEntityData {
    cmgwLifTable.EntityData.YFilter = cmgwLifTable.YFilter
    cmgwLifTable.EntityData.YangName = "cmgwLifTable"
    cmgwLifTable.EntityData.BundleName = "cisco_ios_xe"
    cmgwLifTable.EntityData.ParentYangName = "CISCO-MEDIA-GATEWAY-MIB"
    cmgwLifTable.EntityData.SegmentPath = "cmgwLifTable"
    cmgwLifTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cmgwLifTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cmgwLifTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cmgwLifTable.EntityData.Children = types.NewOrderedMap()
    cmgwLifTable.EntityData.Children.Append("cmgwLifEntry", types.YChild{"CmgwLifEntry", nil})
    for i := range cmgwLifTable.CmgwLifEntry {
        cmgwLifTable.EntityData.Children.Append(types.GetSegmentPath(cmgwLifTable.CmgwLifEntry[i]), types.YChild{"CmgwLifEntry", cmgwLifTable.CmgwLifEntry[i]})
    }
    cmgwLifTable.EntityData.Leafs = types.NewOrderedMap()

    cmgwLifTable.EntityData.YListKeys = []string {}

    return &(cmgwLifTable.EntityData)
}

// CISCOMEDIAGATEWAYMIB_CmgwLifTable_CmgwLifEntry
// An entry of this table is created by the media gateway
// when it supports the VoIP/VoATM application.
type CISCOMEDIAGATEWAYMIB_CmgwLifTable_CmgwLifEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // cisco_media_gateway_mib.CISCOMEDIAGATEWAYMIB_CMediaGwTable_CMediaGwEntry_CmgwIndex
    CmgwIndex interface{}

    // This attribute is a key. An index that uniquely identifies a LIF in the 
    // media gateway. The type is interface{} with range: 1..255.
    CmgwLifNumber interface{}

    // This object represents the total number of PVC within  this LIF.  When
    // users associate/disassociate a PVC with a LIF  by giving a non-zero/zero
    // value of cwacChanLifNum in cwAtmChanExtConfigTable, the value of this
    // object  will be incremented/decremented accordingly.  The value zero means
    // there is no PVC associated with  the LIF. The type is interface{} with
    // range: 0..10000.
    CmgwLifPvcCount interface{}

    // This object represents the total number of Voice Interfaces within this
    // LIF.  When users associate/disassociate a Voice Interface with a LIF by
    // giving a non-zero/zero value of  ccasVoiceCfgLifNumber for the DS0 group in
    // ccasVoiceExtCfgTable, the value of this object will be 
    // incremented/decremented accordingly.   The value zero means there is no
    // Voice Interface associated with the LIF. The type is interface{} with
    // range: 0..1000.
    CmgwLifVoiceIfCount interface{}
}

func (cmgwLifEntry *CISCOMEDIAGATEWAYMIB_CmgwLifTable_CmgwLifEntry) GetEntityData() *types.CommonEntityData {
    cmgwLifEntry.EntityData.YFilter = cmgwLifEntry.YFilter
    cmgwLifEntry.EntityData.YangName = "cmgwLifEntry"
    cmgwLifEntry.EntityData.BundleName = "cisco_ios_xe"
    cmgwLifEntry.EntityData.ParentYangName = "cmgwLifTable"
    cmgwLifEntry.EntityData.SegmentPath = "cmgwLifEntry" + types.AddKeyToken(cmgwLifEntry.CmgwIndex, "cmgwIndex") + types.AddKeyToken(cmgwLifEntry.CmgwLifNumber, "cmgwLifNumber")
    cmgwLifEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cmgwLifEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cmgwLifEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cmgwLifEntry.EntityData.Children = types.NewOrderedMap()
    cmgwLifEntry.EntityData.Leafs = types.NewOrderedMap()
    cmgwLifEntry.EntityData.Leafs.Append("cmgwIndex", types.YLeaf{"CmgwIndex", cmgwLifEntry.CmgwIndex})
    cmgwLifEntry.EntityData.Leafs.Append("cmgwLifNumber", types.YLeaf{"CmgwLifNumber", cmgwLifEntry.CmgwLifNumber})
    cmgwLifEntry.EntityData.Leafs.Append("cmgwLifPvcCount", types.YLeaf{"CmgwLifPvcCount", cmgwLifEntry.CmgwLifPvcCount})
    cmgwLifEntry.EntityData.Leafs.Append("cmgwLifVoiceIfCount", types.YLeaf{"CmgwLifVoiceIfCount", cmgwLifEntry.CmgwLifVoiceIfCount})

    cmgwLifEntry.EntityData.YListKeys = []string {"CmgwIndex", "CmgwLifNumber"}

    return &(cmgwLifEntry.EntityData)
}

// CISCOMEDIAGATEWAYMIB_CMediaGwCallControlConfigTable
// This table defines general call control attributes for
// the media gateway.
type CISCOMEDIAGATEWAYMIB_CMediaGwCallControlConfigTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // One entry for each media gateway which supports call control  protocol. The
    // type is slice of
    // CISCOMEDIAGATEWAYMIB_CMediaGwCallControlConfigTable_CMediaGwCallControlConfigEntry.
    CMediaGwCallControlConfigEntry []*CISCOMEDIAGATEWAYMIB_CMediaGwCallControlConfigTable_CMediaGwCallControlConfigEntry
}

func (cMediaGwCallControlConfigTable *CISCOMEDIAGATEWAYMIB_CMediaGwCallControlConfigTable) GetEntityData() *types.CommonEntityData {
    cMediaGwCallControlConfigTable.EntityData.YFilter = cMediaGwCallControlConfigTable.YFilter
    cMediaGwCallControlConfigTable.EntityData.YangName = "cMediaGwCallControlConfigTable"
    cMediaGwCallControlConfigTable.EntityData.BundleName = "cisco_ios_xe"
    cMediaGwCallControlConfigTable.EntityData.ParentYangName = "CISCO-MEDIA-GATEWAY-MIB"
    cMediaGwCallControlConfigTable.EntityData.SegmentPath = "cMediaGwCallControlConfigTable"
    cMediaGwCallControlConfigTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cMediaGwCallControlConfigTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cMediaGwCallControlConfigTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cMediaGwCallControlConfigTable.EntityData.Children = types.NewOrderedMap()
    cMediaGwCallControlConfigTable.EntityData.Children.Append("cMediaGwCallControlConfigEntry", types.YChild{"CMediaGwCallControlConfigEntry", nil})
    for i := range cMediaGwCallControlConfigTable.CMediaGwCallControlConfigEntry {
        cMediaGwCallControlConfigTable.EntityData.Children.Append(types.GetSegmentPath(cMediaGwCallControlConfigTable.CMediaGwCallControlConfigEntry[i]), types.YChild{"CMediaGwCallControlConfigEntry", cMediaGwCallControlConfigTable.CMediaGwCallControlConfigEntry[i]})
    }
    cMediaGwCallControlConfigTable.EntityData.Leafs = types.NewOrderedMap()

    cMediaGwCallControlConfigTable.EntityData.YListKeys = []string {}

    return &(cMediaGwCallControlConfigTable.EntityData)
}

// CISCOMEDIAGATEWAYMIB_CMediaGwCallControlConfigTable_CMediaGwCallControlConfigEntry
// One entry for each media gateway which supports call control 
// protocol.
type CISCOMEDIAGATEWAYMIB_CMediaGwCallControlConfigTable_CMediaGwCallControlConfigEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // cisco_media_gateway_mib.CISCOMEDIAGATEWAYMIB_CMediaGwTable_CMediaGwEntry_CmgwIndex
    CmgwIndex interface{}

    // This object specifies Type Of Service (TOS) field of IP header for the
    // signaling control packet in VoIP application. The type is interface{} with
    // range: 0..255.
    CMediaGwCcCfgControlTos interface{}

    // This object specifies Type Of Service (TOS) field of IP header for the
    // voice payload packet in VoIP application. The type is interface{} with
    // range: 0..255.
    CMediaGwCcCfgBearerTos interface{}

    // This object specifies NTE (Named Telephony Events) payload type. The type
    // is interface{} with range: 96..127.
    CMediaGwCcCfgNtePayload interface{}

    // This object specifies NSE (Network Signaling Events) payload type. The type
    // is interface{} with range: 98..117.
    CMediaGwCcCfgNsePayload interface{}

    // This object specifies Network Signaling Event (NSE) timeout value. The type
    // is interface{} with range: 250..10000. Units are milliseconds.
    CMediaGwCcCfgNseRespTimer interface{}

    // The object specifies the jitter buffer mode applied to a VBD (Voice Band
    // Data) call connection.  adaptive - means to use
    // cMediaGwCcCfgVbdJitterNomDelay as            the initial jitter buffers
    // size and let the DSP            pick the optimal value of the jitter buffer
    // size between the range of            cMediaGwCcCfgVbcJitterMaxDelay and    
    // cMediaGwCcCfgVbcJitterMinDelay.  fixed - means to use a constant jitter
    // buffer size         which is specified by cMediaGwCcCfgVbdJitterNomDelay.
    // The type is CCallControlJitterDelayMode.
    CMediaGwCcCfgVbdJitterDelayMode interface{}

    // This object specifies the maximum jitter buffer size  in VBD (Voice Band
    // Data). The type is interface{} with range: 20..135. Units are milliseconds.
    CMediaGwCcCfgVbdJitterMaxDelay interface{}

    // This object specifies the nominal jitter buffer size  in VBD (Voice Band
    // Data). The type is interface{} with range: 5..135. Units are milliseconds.
    CMediaGwCcCfgVbdJitterNomDelay interface{}

    // This object specifies the nominal jitter buffer size  in VBD (Voice Band
    // Data). The type is interface{} with range: 0..135. Units are milliseconds.
    CMediaGwCcCfgVbdJitterMinDelay interface{}

    // This object specifies the default tone plan index (the value of
    // cvtcTonePlanId) for the media gateway. The type is interface{} with range:
    // 1..65535.
    CMediaGwCcCfgDefaultTonePlanId interface{}

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
    CMediaGwCcCfgDescrInfoEnabled interface{}

    // This object specifies the prefix of the name schema for DS (Digital
    // Subscriber) terminations. The value of this object must be unique among the
    // following objects:        cMediaGwCcCfgDsNamePrefix       
    // cMediaGwCcCfgRtpNamePrefix        cMediaGwCcCfgAal2SvcNamePrefix       
    // cMediaGwCcCfgAal2SvcNamePrefix        cMediaGwCcCfgDefRtpNamePrefix This
    // object can not be modified when there is any DS termination existing in the
    // media gateway. It is default to 'DS'. The type is string with length:
    // 1..64.
    CMediaGwCcCfgDsNamePrefix interface{}

    // This object specifies the prefix of the name schema for RTP (Real-Time
    // Transport Protocol) terminations. The value of this object must be unique
    // among the  following objects:        cMediaGwCcCfgDsNamePrefix       
    // cMediaGwCcCfgRtpNamePrefix        cMediaGwCcCfgAal2SvcNamePrefix       
    // cMediaGwCcCfgAal2SvcNamePrefix        cMediaGwCcCfgDefRtpNamePrefix This
    // object can not be modified when there is any RTP termination type existing
    // in the media gateway. It is default to 'RTP'. The type is string with
    // length: 1..64.
    CMediaGwCcCfgRtpNamePrefix interface{}

    // This object specifies the prefix of the name schema for voice over AAL1 SVC
    // (Switched Virtual Circuit) terminations. The value of this object must be
    // unique among the  following objects:        cMediaGwCcCfgDsNamePrefix      
    // cMediaGwCcCfgRtpNamePrefix        cMediaGwCcCfgAal2SvcNamePrefix       
    // cMediaGwCcCfgAal2SvcNamePrefix        cMediaGwCcCfgDefRtpNamePrefix This
    // object can not be modified when there is any AAL1 SVC termination type
    // existing in the media gateway. It is default to 'AAL1/SVC'. The type is
    // string with length: 1..64.
    CMediaGwCcCfgAal1SvcNamePrefix interface{}

    // This object specifies the prefix of the name schema for voice over AAL2 SVC
    // (Switched Virtual Circuit) terminations. The value of this object must be
    // unique among the  following objects:        cMediaGwCcCfgDsNamePrefix      
    // cMediaGwCcCfgRtpNamePrefix        cMediaGwCcCfgAal2SvcNamePrefix       
    // cMediaGwCcCfgAal2SvcNamePrefix        cMediaGwCcCfgDefRtpNamePrefix This
    // object can not be modified when there is any AAL2 SVC termination type
    // existing in the media gateway. It is default to 'AAL2/SVC'. The type is
    // string with length: 1..64.
    CMediaGwCcCfgAal2SvcNamePrefix interface{}

    // This object specifies the condition of the cluster generation in the call
    // control.  A cluster is a group of endpoints that share a particular bearer
    // possibility for connections among each other.  disabled(1) - The generation
    // of the cluster attribute               is disabled. enabled(2) -
    // Unconditionally generate the cluster              attribute.
    // conditionalEnabled(3) - The generation of the cluster               
    // attribute is upon MGC request. The type is CMediaGwCcCfgClusterEnabled.
    CMediaGwCcCfgClusterEnabled interface{}

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
    // CMediaGwCcCfgDefBearerTraffic.
    CMediaGwCcCfgDefBearerTraffic interface{}

    // This object specifies the prefix of the name schema for default RTP
    // terminations. The value of this object must be unique among the  following
    // objects:        cMediaGwCcCfgDsNamePrefix        cMediaGwCcCfgRtpNamePrefix
    // cMediaGwCcCfgAal1SvcNamePrefix        cMediaGwCcCfgAal2SvcNamePrefix  It is
    // defaulted to 'TGWRTP'. The type is string with length: 1..64.
    CMediaGwCcCfgDefRtpNamePrefix interface{}
}

func (cMediaGwCallControlConfigEntry *CISCOMEDIAGATEWAYMIB_CMediaGwCallControlConfigTable_CMediaGwCallControlConfigEntry) GetEntityData() *types.CommonEntityData {
    cMediaGwCallControlConfigEntry.EntityData.YFilter = cMediaGwCallControlConfigEntry.YFilter
    cMediaGwCallControlConfigEntry.EntityData.YangName = "cMediaGwCallControlConfigEntry"
    cMediaGwCallControlConfigEntry.EntityData.BundleName = "cisco_ios_xe"
    cMediaGwCallControlConfigEntry.EntityData.ParentYangName = "cMediaGwCallControlConfigTable"
    cMediaGwCallControlConfigEntry.EntityData.SegmentPath = "cMediaGwCallControlConfigEntry" + types.AddKeyToken(cMediaGwCallControlConfigEntry.CmgwIndex, "cmgwIndex")
    cMediaGwCallControlConfigEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cMediaGwCallControlConfigEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cMediaGwCallControlConfigEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cMediaGwCallControlConfigEntry.EntityData.Children = types.NewOrderedMap()
    cMediaGwCallControlConfigEntry.EntityData.Leafs = types.NewOrderedMap()
    cMediaGwCallControlConfigEntry.EntityData.Leafs.Append("cmgwIndex", types.YLeaf{"CmgwIndex", cMediaGwCallControlConfigEntry.CmgwIndex})
    cMediaGwCallControlConfigEntry.EntityData.Leafs.Append("cMediaGwCcCfgControlTos", types.YLeaf{"CMediaGwCcCfgControlTos", cMediaGwCallControlConfigEntry.CMediaGwCcCfgControlTos})
    cMediaGwCallControlConfigEntry.EntityData.Leafs.Append("cMediaGwCcCfgBearerTos", types.YLeaf{"CMediaGwCcCfgBearerTos", cMediaGwCallControlConfigEntry.CMediaGwCcCfgBearerTos})
    cMediaGwCallControlConfigEntry.EntityData.Leafs.Append("cMediaGwCcCfgNtePayload", types.YLeaf{"CMediaGwCcCfgNtePayload", cMediaGwCallControlConfigEntry.CMediaGwCcCfgNtePayload})
    cMediaGwCallControlConfigEntry.EntityData.Leafs.Append("cMediaGwCcCfgNsePayload", types.YLeaf{"CMediaGwCcCfgNsePayload", cMediaGwCallControlConfigEntry.CMediaGwCcCfgNsePayload})
    cMediaGwCallControlConfigEntry.EntityData.Leafs.Append("cMediaGwCcCfgNseRespTimer", types.YLeaf{"CMediaGwCcCfgNseRespTimer", cMediaGwCallControlConfigEntry.CMediaGwCcCfgNseRespTimer})
    cMediaGwCallControlConfigEntry.EntityData.Leafs.Append("cMediaGwCcCfgVbdJitterDelayMode", types.YLeaf{"CMediaGwCcCfgVbdJitterDelayMode", cMediaGwCallControlConfigEntry.CMediaGwCcCfgVbdJitterDelayMode})
    cMediaGwCallControlConfigEntry.EntityData.Leafs.Append("cMediaGwCcCfgVbdJitterMaxDelay", types.YLeaf{"CMediaGwCcCfgVbdJitterMaxDelay", cMediaGwCallControlConfigEntry.CMediaGwCcCfgVbdJitterMaxDelay})
    cMediaGwCallControlConfigEntry.EntityData.Leafs.Append("cMediaGwCcCfgVbdJitterNomDelay", types.YLeaf{"CMediaGwCcCfgVbdJitterNomDelay", cMediaGwCallControlConfigEntry.CMediaGwCcCfgVbdJitterNomDelay})
    cMediaGwCallControlConfigEntry.EntityData.Leafs.Append("cMediaGwCcCfgVbdJitterMinDelay", types.YLeaf{"CMediaGwCcCfgVbdJitterMinDelay", cMediaGwCallControlConfigEntry.CMediaGwCcCfgVbdJitterMinDelay})
    cMediaGwCallControlConfigEntry.EntityData.Leafs.Append("cMediaGwCcCfgDefaultTonePlanId", types.YLeaf{"CMediaGwCcCfgDefaultTonePlanId", cMediaGwCallControlConfigEntry.CMediaGwCcCfgDefaultTonePlanId})
    cMediaGwCallControlConfigEntry.EntityData.Leafs.Append("cMediaGwCcCfgDescrInfoEnabled", types.YLeaf{"CMediaGwCcCfgDescrInfoEnabled", cMediaGwCallControlConfigEntry.CMediaGwCcCfgDescrInfoEnabled})
    cMediaGwCallControlConfigEntry.EntityData.Leafs.Append("cMediaGwCcCfgDsNamePrefix", types.YLeaf{"CMediaGwCcCfgDsNamePrefix", cMediaGwCallControlConfigEntry.CMediaGwCcCfgDsNamePrefix})
    cMediaGwCallControlConfigEntry.EntityData.Leafs.Append("cMediaGwCcCfgRtpNamePrefix", types.YLeaf{"CMediaGwCcCfgRtpNamePrefix", cMediaGwCallControlConfigEntry.CMediaGwCcCfgRtpNamePrefix})
    cMediaGwCallControlConfigEntry.EntityData.Leafs.Append("cMediaGwCcCfgAal1SvcNamePrefix", types.YLeaf{"CMediaGwCcCfgAal1SvcNamePrefix", cMediaGwCallControlConfigEntry.CMediaGwCcCfgAal1SvcNamePrefix})
    cMediaGwCallControlConfigEntry.EntityData.Leafs.Append("cMediaGwCcCfgAal2SvcNamePrefix", types.YLeaf{"CMediaGwCcCfgAal2SvcNamePrefix", cMediaGwCallControlConfigEntry.CMediaGwCcCfgAal2SvcNamePrefix})
    cMediaGwCallControlConfigEntry.EntityData.Leafs.Append("cMediaGwCcCfgClusterEnabled", types.YLeaf{"CMediaGwCcCfgClusterEnabled", cMediaGwCallControlConfigEntry.CMediaGwCcCfgClusterEnabled})
    cMediaGwCallControlConfigEntry.EntityData.Leafs.Append("cMediaGwCcCfgDefBearerTraffic", types.YLeaf{"CMediaGwCcCfgDefBearerTraffic", cMediaGwCallControlConfigEntry.CMediaGwCcCfgDefBearerTraffic})
    cMediaGwCallControlConfigEntry.EntityData.Leafs.Append("cMediaGwCcCfgDefRtpNamePrefix", types.YLeaf{"CMediaGwCcCfgDefRtpNamePrefix", cMediaGwCallControlConfigEntry.CMediaGwCcCfgDefRtpNamePrefix})

    cMediaGwCallControlConfigEntry.EntityData.YListKeys = []string {"CmgwIndex"}

    return &(cMediaGwCallControlConfigEntry.EntityData)
}

// CISCOMEDIAGATEWAYMIB_CMediaGwCallControlConfigTable_CMediaGwCallControlConfigEntry_CMediaGwCcCfgClusterEnabled represents               attribute is upon MGC request.
type CISCOMEDIAGATEWAYMIB_CMediaGwCallControlConfigTable_CMediaGwCallControlConfigEntry_CMediaGwCcCfgClusterEnabled string

const (
    CISCOMEDIAGATEWAYMIB_CMediaGwCallControlConfigTable_CMediaGwCallControlConfigEntry_CMediaGwCcCfgClusterEnabled_disabled CISCOMEDIAGATEWAYMIB_CMediaGwCallControlConfigTable_CMediaGwCallControlConfigEntry_CMediaGwCcCfgClusterEnabled = "disabled"

    CISCOMEDIAGATEWAYMIB_CMediaGwCallControlConfigTable_CMediaGwCallControlConfigEntry_CMediaGwCcCfgClusterEnabled_enabled CISCOMEDIAGATEWAYMIB_CMediaGwCallControlConfigTable_CMediaGwCallControlConfigEntry_CMediaGwCcCfgClusterEnabled = "enabled"

    CISCOMEDIAGATEWAYMIB_CMediaGwCallControlConfigTable_CMediaGwCallControlConfigEntry_CMediaGwCcCfgClusterEnabled_conditionalEnabled CISCOMEDIAGATEWAYMIB_CMediaGwCallControlConfigTable_CMediaGwCallControlConfigEntry_CMediaGwCcCfgClusterEnabled = "conditionalEnabled"
)

// CISCOMEDIAGATEWAYMIB_CMediaGwCallControlConfigTable_CMediaGwCallControlConfigEntry_CMediaGwCcCfgDefBearerTraffic represents media gateway sending CRCX response.
type CISCOMEDIAGATEWAYMIB_CMediaGwCallControlConfigTable_CMediaGwCallControlConfigEntry_CMediaGwCcCfgDefBearerTraffic string

const (
    CISCOMEDIAGATEWAYMIB_CMediaGwCallControlConfigTable_CMediaGwCallControlConfigEntry_CMediaGwCcCfgDefBearerTraffic_ipPvcAal5 CISCOMEDIAGATEWAYMIB_CMediaGwCallControlConfigTable_CMediaGwCallControlConfigEntry_CMediaGwCcCfgDefBearerTraffic = "ipPvcAal5"

    CISCOMEDIAGATEWAYMIB_CMediaGwCallControlConfigTable_CMediaGwCallControlConfigEntry_CMediaGwCcCfgDefBearerTraffic_atmPvcAal2 CISCOMEDIAGATEWAYMIB_CMediaGwCallControlConfigTable_CMediaGwCallControlConfigEntry_CMediaGwCcCfgDefBearerTraffic = "atmPvcAal2"

    CISCOMEDIAGATEWAYMIB_CMediaGwCallControlConfigTable_CMediaGwCallControlConfigEntry_CMediaGwCcCfgDefBearerTraffic_atmSvcAal2 CISCOMEDIAGATEWAYMIB_CMediaGwCallControlConfigTable_CMediaGwCallControlConfigEntry_CMediaGwCcCfgDefBearerTraffic = "atmSvcAal2"

    CISCOMEDIAGATEWAYMIB_CMediaGwCallControlConfigTable_CMediaGwCallControlConfigEntry_CMediaGwCcCfgDefBearerTraffic_atmSvcAal1 CISCOMEDIAGATEWAYMIB_CMediaGwCallControlConfigTable_CMediaGwCallControlConfigEntry_CMediaGwCcCfgDefBearerTraffic = "atmSvcAal1"
)

// CISCOMEDIAGATEWAYMIB_CMediaGwRscStatsTable
// This table stores the gateway resource statistics
// information.
type CISCOMEDIAGATEWAYMIB_CMediaGwRscStatsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry stores the statistics information for a specific resource. The
    // type is slice of
    // CISCOMEDIAGATEWAYMIB_CMediaGwRscStatsTable_CMediaGwRscStatsEntry.
    CMediaGwRscStatsEntry []*CISCOMEDIAGATEWAYMIB_CMediaGwRscStatsTable_CMediaGwRscStatsEntry
}

func (cMediaGwRscStatsTable *CISCOMEDIAGATEWAYMIB_CMediaGwRscStatsTable) GetEntityData() *types.CommonEntityData {
    cMediaGwRscStatsTable.EntityData.YFilter = cMediaGwRscStatsTable.YFilter
    cMediaGwRscStatsTable.EntityData.YangName = "cMediaGwRscStatsTable"
    cMediaGwRscStatsTable.EntityData.BundleName = "cisco_ios_xe"
    cMediaGwRscStatsTable.EntityData.ParentYangName = "CISCO-MEDIA-GATEWAY-MIB"
    cMediaGwRscStatsTable.EntityData.SegmentPath = "cMediaGwRscStatsTable"
    cMediaGwRscStatsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cMediaGwRscStatsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cMediaGwRscStatsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cMediaGwRscStatsTable.EntityData.Children = types.NewOrderedMap()
    cMediaGwRscStatsTable.EntityData.Children.Append("cMediaGwRscStatsEntry", types.YChild{"CMediaGwRscStatsEntry", nil})
    for i := range cMediaGwRscStatsTable.CMediaGwRscStatsEntry {
        cMediaGwRscStatsTable.EntityData.Children.Append(types.GetSegmentPath(cMediaGwRscStatsTable.CMediaGwRscStatsEntry[i]), types.YChild{"CMediaGwRscStatsEntry", cMediaGwRscStatsTable.CMediaGwRscStatsEntry[i]})
    }
    cMediaGwRscStatsTable.EntityData.Leafs = types.NewOrderedMap()

    cMediaGwRscStatsTable.EntityData.YListKeys = []string {}

    return &(cMediaGwRscStatsTable.EntityData)
}

// CISCOMEDIAGATEWAYMIB_CMediaGwRscStatsTable_CMediaGwRscStatsEntry
// Each entry stores the statistics
// information for a specific resource.
type CISCOMEDIAGATEWAYMIB_CMediaGwRscStatsTable_CMediaGwRscStatsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // cisco_media_gateway_mib.CISCOMEDIAGATEWAYMIB_CMediaGwTable_CMediaGwEntry_CmgwIndex
    CmgwIndex interface{}

    // This attribute is a key. An index that uniquely identifies a specific
    // gateway resource. The type is CmgwRscStatsIndex.
    CmgwRscStatsIndex interface{}

    // This object indicates the maximum utilization of the resource over the
    // interval specified by the 'cmgwRscSinceLastReset'. The type is interface{}
    // with range: 0..4294967295.
    CmgwRscMaximumUtilization interface{}

    // This object indicates the minimum utilization of the resource over the
    // interval specified by the 'cmgwRscSinceLastReset'. The type is interface{}
    // with range: 0..4294967295.
    CmgwRscMinimumUtilization interface{}

    // This object indicates the average utilization of the resource over the
    // interval specified by the 'cmgwRscSinceLastReset'. The type is interface{}
    // with range: 0..4294967295.
    CmgwRscAverageUtilization interface{}

    // The elapsed time (in seconds) from the last periodic reset.  The following
    // objects are reset at the last reset:      'cmgwRscMaximumUtilization'    
    // 'cmgwRscMinimumUtilization'     'cmgwRscAverageUtilization'. The type is
    // interface{} with range: 0..4294967295. Units are seconds.
    CmgwRscSinceLastReset interface{}
}

func (cMediaGwRscStatsEntry *CISCOMEDIAGATEWAYMIB_CMediaGwRscStatsTable_CMediaGwRscStatsEntry) GetEntityData() *types.CommonEntityData {
    cMediaGwRscStatsEntry.EntityData.YFilter = cMediaGwRscStatsEntry.YFilter
    cMediaGwRscStatsEntry.EntityData.YangName = "cMediaGwRscStatsEntry"
    cMediaGwRscStatsEntry.EntityData.BundleName = "cisco_ios_xe"
    cMediaGwRscStatsEntry.EntityData.ParentYangName = "cMediaGwRscStatsTable"
    cMediaGwRscStatsEntry.EntityData.SegmentPath = "cMediaGwRscStatsEntry" + types.AddKeyToken(cMediaGwRscStatsEntry.CmgwIndex, "cmgwIndex") + types.AddKeyToken(cMediaGwRscStatsEntry.CmgwRscStatsIndex, "cmgwRscStatsIndex")
    cMediaGwRscStatsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cMediaGwRscStatsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cMediaGwRscStatsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cMediaGwRscStatsEntry.EntityData.Children = types.NewOrderedMap()
    cMediaGwRscStatsEntry.EntityData.Leafs = types.NewOrderedMap()
    cMediaGwRscStatsEntry.EntityData.Leafs.Append("cmgwIndex", types.YLeaf{"CmgwIndex", cMediaGwRscStatsEntry.CmgwIndex})
    cMediaGwRscStatsEntry.EntityData.Leafs.Append("cmgwRscStatsIndex", types.YLeaf{"CmgwRscStatsIndex", cMediaGwRscStatsEntry.CmgwRscStatsIndex})
    cMediaGwRscStatsEntry.EntityData.Leafs.Append("cmgwRscMaximumUtilization", types.YLeaf{"CmgwRscMaximumUtilization", cMediaGwRscStatsEntry.CmgwRscMaximumUtilization})
    cMediaGwRscStatsEntry.EntityData.Leafs.Append("cmgwRscMinimumUtilization", types.YLeaf{"CmgwRscMinimumUtilization", cMediaGwRscStatsEntry.CmgwRscMinimumUtilization})
    cMediaGwRscStatsEntry.EntityData.Leafs.Append("cmgwRscAverageUtilization", types.YLeaf{"CmgwRscAverageUtilization", cMediaGwRscStatsEntry.CmgwRscAverageUtilization})
    cMediaGwRscStatsEntry.EntityData.Leafs.Append("cmgwRscSinceLastReset", types.YLeaf{"CmgwRscSinceLastReset", cMediaGwRscStatsEntry.CmgwRscSinceLastReset})

    cMediaGwRscStatsEntry.EntityData.YListKeys = []string {"CmgwIndex", "CmgwRscStatsIndex"}

    return &(cMediaGwRscStatsEntry.EntityData)
}

// CISCOMEDIAGATEWAYMIB_CMediaGwRscStatsTable_CMediaGwRscStatsEntry_CmgwRscStatsIndex represents resource.
type CISCOMEDIAGATEWAYMIB_CMediaGwRscStatsTable_CMediaGwRscStatsEntry_CmgwRscStatsIndex string

const (
    CISCOMEDIAGATEWAYMIB_CMediaGwRscStatsTable_CMediaGwRscStatsEntry_CmgwRscStatsIndex_cpu CISCOMEDIAGATEWAYMIB_CMediaGwRscStatsTable_CMediaGwRscStatsEntry_CmgwRscStatsIndex = "cpu"

    CISCOMEDIAGATEWAYMIB_CMediaGwRscStatsTable_CMediaGwRscStatsEntry_CmgwRscStatsIndex_staticmemory CISCOMEDIAGATEWAYMIB_CMediaGwRscStatsTable_CMediaGwRscStatsEntry_CmgwRscStatsIndex = "staticmemory"

    CISCOMEDIAGATEWAYMIB_CMediaGwRscStatsTable_CMediaGwRscStatsEntry_CmgwRscStatsIndex_dynamicmemory CISCOMEDIAGATEWAYMIB_CMediaGwRscStatsTable_CMediaGwRscStatsEntry_CmgwRscStatsIndex = "dynamicmemory"

    CISCOMEDIAGATEWAYMIB_CMediaGwRscStatsTable_CMediaGwRscStatsEntry_CmgwRscStatsIndex_sysmemory CISCOMEDIAGATEWAYMIB_CMediaGwRscStatsTable_CMediaGwRscStatsEntry_CmgwRscStatsIndex = "sysmemory"

    CISCOMEDIAGATEWAYMIB_CMediaGwRscStatsTable_CMediaGwRscStatsEntry_CmgwRscStatsIndex_commbuffer CISCOMEDIAGATEWAYMIB_CMediaGwRscStatsTable_CMediaGwRscStatsEntry_CmgwRscStatsIndex = "commbuffer"

    CISCOMEDIAGATEWAYMIB_CMediaGwRscStatsTable_CMediaGwRscStatsEntry_CmgwRscStatsIndex_msgq CISCOMEDIAGATEWAYMIB_CMediaGwRscStatsTable_CMediaGwRscStatsEntry_CmgwRscStatsIndex = "msgq"

    CISCOMEDIAGATEWAYMIB_CMediaGwRscStatsTable_CMediaGwRscStatsEntry_CmgwRscStatsIndex_atmq CISCOMEDIAGATEWAYMIB_CMediaGwRscStatsTable_CMediaGwRscStatsEntry_CmgwRscStatsIndex = "atmq"

    CISCOMEDIAGATEWAYMIB_CMediaGwRscStatsTable_CMediaGwRscStatsEntry_CmgwRscStatsIndex_svccongestion CISCOMEDIAGATEWAYMIB_CMediaGwRscStatsTable_CMediaGwRscStatsEntry_CmgwRscStatsIndex = "svccongestion"

    CISCOMEDIAGATEWAYMIB_CMediaGwRscStatsTable_CMediaGwRscStatsEntry_CmgwRscStatsIndex_rsvpq CISCOMEDIAGATEWAYMIB_CMediaGwRscStatsTable_CMediaGwRscStatsEntry_CmgwRscStatsIndex = "rsvpq"

    CISCOMEDIAGATEWAYMIB_CMediaGwRscStatsTable_CMediaGwRscStatsEntry_CmgwRscStatsIndex_dspq CISCOMEDIAGATEWAYMIB_CMediaGwRscStatsTable_CMediaGwRscStatsEntry_CmgwRscStatsIndex = "dspq"

    CISCOMEDIAGATEWAYMIB_CMediaGwRscStatsTable_CMediaGwRscStatsEntry_CmgwRscStatsIndex_h248congestion CISCOMEDIAGATEWAYMIB_CMediaGwRscStatsTable_CMediaGwRscStatsEntry_CmgwRscStatsIndex = "h248congestion"

    CISCOMEDIAGATEWAYMIB_CMediaGwRscStatsTable_CMediaGwRscStatsEntry_CmgwRscStatsIndex_callpersec CISCOMEDIAGATEWAYMIB_CMediaGwRscStatsTable_CMediaGwRscStatsEntry_CmgwRscStatsIndex = "callpersec"

    CISCOMEDIAGATEWAYMIB_CMediaGwRscStatsTable_CMediaGwRscStatsEntry_CmgwRscStatsIndex_smallipcbuffer CISCOMEDIAGATEWAYMIB_CMediaGwRscStatsTable_CMediaGwRscStatsEntry_CmgwRscStatsIndex = "smallipcbuffer"

    CISCOMEDIAGATEWAYMIB_CMediaGwRscStatsTable_CMediaGwRscStatsEntry_CmgwRscStatsIndex_mediumipcbuffer CISCOMEDIAGATEWAYMIB_CMediaGwRscStatsTable_CMediaGwRscStatsEntry_CmgwRscStatsIndex = "mediumipcbuffer"

    CISCOMEDIAGATEWAYMIB_CMediaGwRscStatsTable_CMediaGwRscStatsEntry_CmgwRscStatsIndex_largeipcbuffer CISCOMEDIAGATEWAYMIB_CMediaGwRscStatsTable_CMediaGwRscStatsEntry_CmgwRscStatsIndex = "largeipcbuffer"

    CISCOMEDIAGATEWAYMIB_CMediaGwRscStatsTable_CMediaGwRscStatsEntry_CmgwRscStatsIndex_hugeipcbuffer CISCOMEDIAGATEWAYMIB_CMediaGwRscStatsTable_CMediaGwRscStatsEntry_CmgwRscStatsIndex = "hugeipcbuffer"

    CISCOMEDIAGATEWAYMIB_CMediaGwRscStatsTable_CMediaGwRscStatsEntry_CmgwRscStatsIndex_mblkipcbuffer CISCOMEDIAGATEWAYMIB_CMediaGwRscStatsTable_CMediaGwRscStatsEntry_CmgwRscStatsIndex = "mblkipcbuffer"
)

