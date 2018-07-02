// The MIB module for VPDN.
// 
// Overview of VPDN MIB
// 
// MIB description
// 
// This MIB is to support the Virtual Private Dialup Network (VPDN)
// feature of Cisco IOS.  VPDN handles the forwarding of PPP links
// from an Internet Provider (ISP) to a Home Gateway.
// 
// The VPDN MIB provides the operational information on Cisco's
// VPDN tunnelling implementation.  The following entities are
// managed:
// 1) Global VPDN information
// 2) VPDN tunnel information
// 3) VPDN tunnel's user information
// 4) Failure history per user
// 
// The following example configuration shows how the VPDN MIB
// returns VPDN information, from either CISCO A - Network Access
// Server (NAS) or CISCO B - Tunnel Server (TS).  The User call is
// projected by either domain name or Dialed Number Identification
// Service (DNIS).
// 
// The terms NAS and TS are generic terms refering to the VPDN
// systems.
// 
// The following table shows the corresponding technology-specific
// terms.
// 
//       Network Access Server            Tunnel Server
//       ------------------------------   -------------------------
// L2F   Network Access Server    (NAS)   Home Gateway (HGW)
// L2TP  L2TP Access Concentrator (LAC)   L2TP Network Server (LNS)
// PPTP  PPTP Access Concentrator (PAC)   PPTP Network Server (PNS)
// 
//            (NAS)                          (TS)
// User ===== Cisco A ===== IP Network ===== Cisco B ===== Server
//                 |                          |
//                 +========== VPDN ==========+
// 
// 1) The VPDN global entry identifies the system wide VPDN
//    information.
// 2) The VPDN tunnel table identifies the active VPDN tunnels on
//    Cisco A and Cisco B.  The table contains an entry for each
//    active tunnel on the system.
// 3) The VPDN tunnel user table identifies the active users for
//    each active tunnel on each system and provides relevant
//    information for each user.
// 4) The VPDN failure history table identifies the last failure
//    recorded per user and provides relevant information.
package cisco_vpdn_mgmt_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_vpdn_mgmt_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-VPDN-MGMT-MIB CISCO-VPDN-MGMT-MIB}", reflect.TypeOf(CISCOVPDNMGMTMIB{}))
    ydk.RegisterEntity("CISCO-VPDN-MGMT-MIB:CISCO-VPDN-MGMT-MIB", reflect.TypeOf(CISCOVPDNMGMTMIB{}))
}

// EndpointClass represents              SnmpAdminString (SIZE(1..15)) value.
type EndpointClass string

const (
    EndpointClass_none EndpointClass = "none"

    EndpointClass_local EndpointClass = "local"

    EndpointClass_ipV4Address EndpointClass = "ipV4Address"

    EndpointClass_macAddress EndpointClass = "macAddress"

    EndpointClass_magicNumber EndpointClass = "magicNumber"

    EndpointClass_phone EndpointClass = "phone"
)

// TunnelType represents tunnel.
type TunnelType string

const (
    TunnelType_l2f TunnelType = "l2f"

    TunnelType_l2tp TunnelType = "l2tp"

    TunnelType_pptp TunnelType = "pptp"
)

// CISCOVPDNMGMTMIB
type CISCOVPDNMGMTMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    CiscoVpdnMgmtMIBNotifs CISCOVPDNMGMTMIB_CiscoVpdnMgmtMIBNotifs

    
    CvpdnSystemInfo CISCOVPDNMGMTMIB_CvpdnSystemInfo

    
    CvpdnMultilinkInfo CISCOVPDNMGMTMIB_CvpdnMultilinkInfo

    // Table of information about the VPDN system for all tunnel types.
    CvpdnSystemTable CISCOVPDNMGMTMIB_CvpdnSystemTable

    // Table of information about the active VPDN tunnels.
    CvpdnTunnelTable CISCOVPDNMGMTMIB_CvpdnTunnelTable

    // Table of information about the active VPDN tunnels.  An entry is added to
    // the table when a new tunnel is initiated and removed from the table when
    // the tunnel is terminated.
    CvpdnTunnelAttrTable CISCOVPDNMGMTMIB_CvpdnTunnelAttrTable

    // Table of information about individual user sessions within the active
    // tunnels.  Entry is added to the table when new user session is initiated
    // and be removed from the table when the user session is terminated.
    CvpdnTunnelSessionTable CISCOVPDNMGMTMIB_CvpdnTunnelSessionTable

    // Table of information about individual sessions within the active tunnels. 
    // An entry is added to the table when a new session is initiated and removed
    // from the table when the session is terminated.
    CvpdnSessionAttrTable CISCOVPDNMGMTMIB_CvpdnSessionAttrTable

    // Table of the record of failure objects which can be referenced by an user
    // name.  Only a name that has a valid item in the Cisco IOS VPDN failure
    // history table will yield a valid entry in this table.  The table has a
    // maximum size of 50 entries.  Only the newest 50 entries will be kept in the
    // table.
    CvpdnUserToFailHistInfoTable CISCOVPDNMGMTMIB_CvpdnUserToFailHistInfoTable

    // Table of information about the VPDN templates.  The VPDN template is a
    // grouping mechanism that allows configuration settings to be shared among
    // multiple VPDN groups.  One such setting is a limit on the number of active
    // sessions across all VPDN groups associated with the template.  The template
    // table allows customers to monitor template-wide information such as
    // tracking the allocation of sessions across templates.
    CvpdnTemplateTable CISCOVPDNMGMTMIB_CvpdnTemplateTable

    // Table that describes the multilink PPP attributes of the active VPDN
    // sessions.
    CvpdnBundleTable CISCOVPDNMGMTMIB_CvpdnBundleTable

    // A table that exposes the containment relationship between a multilink PPP
    // bundle and a VPDN tunnel.
    CvpdnBundleChildTable CISCOVPDNMGMTMIB_CvpdnBundleChildTable
}

func (cISCOVPDNMGMTMIB *CISCOVPDNMGMTMIB) GetEntityData() *types.CommonEntityData {
    cISCOVPDNMGMTMIB.EntityData.YFilter = cISCOVPDNMGMTMIB.YFilter
    cISCOVPDNMGMTMIB.EntityData.YangName = "CISCO-VPDN-MGMT-MIB"
    cISCOVPDNMGMTMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOVPDNMGMTMIB.EntityData.ParentYangName = "CISCO-VPDN-MGMT-MIB"
    cISCOVPDNMGMTMIB.EntityData.SegmentPath = "CISCO-VPDN-MGMT-MIB:CISCO-VPDN-MGMT-MIB"
    cISCOVPDNMGMTMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOVPDNMGMTMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOVPDNMGMTMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOVPDNMGMTMIB.EntityData.Children = types.NewOrderedMap()
    cISCOVPDNMGMTMIB.EntityData.Children.Append("ciscoVpdnMgmtMIBNotifs", types.YChild{"CiscoVpdnMgmtMIBNotifs", &cISCOVPDNMGMTMIB.CiscoVpdnMgmtMIBNotifs})
    cISCOVPDNMGMTMIB.EntityData.Children.Append("cvpdnSystemInfo", types.YChild{"CvpdnSystemInfo", &cISCOVPDNMGMTMIB.CvpdnSystemInfo})
    cISCOVPDNMGMTMIB.EntityData.Children.Append("cvpdnMultilinkInfo", types.YChild{"CvpdnMultilinkInfo", &cISCOVPDNMGMTMIB.CvpdnMultilinkInfo})
    cISCOVPDNMGMTMIB.EntityData.Children.Append("cvpdnSystemTable", types.YChild{"CvpdnSystemTable", &cISCOVPDNMGMTMIB.CvpdnSystemTable})
    cISCOVPDNMGMTMIB.EntityData.Children.Append("cvpdnTunnelTable", types.YChild{"CvpdnTunnelTable", &cISCOVPDNMGMTMIB.CvpdnTunnelTable})
    cISCOVPDNMGMTMIB.EntityData.Children.Append("cvpdnTunnelAttrTable", types.YChild{"CvpdnTunnelAttrTable", &cISCOVPDNMGMTMIB.CvpdnTunnelAttrTable})
    cISCOVPDNMGMTMIB.EntityData.Children.Append("cvpdnTunnelSessionTable", types.YChild{"CvpdnTunnelSessionTable", &cISCOVPDNMGMTMIB.CvpdnTunnelSessionTable})
    cISCOVPDNMGMTMIB.EntityData.Children.Append("cvpdnSessionAttrTable", types.YChild{"CvpdnSessionAttrTable", &cISCOVPDNMGMTMIB.CvpdnSessionAttrTable})
    cISCOVPDNMGMTMIB.EntityData.Children.Append("cvpdnUserToFailHistInfoTable", types.YChild{"CvpdnUserToFailHistInfoTable", &cISCOVPDNMGMTMIB.CvpdnUserToFailHistInfoTable})
    cISCOVPDNMGMTMIB.EntityData.Children.Append("cvpdnTemplateTable", types.YChild{"CvpdnTemplateTable", &cISCOVPDNMGMTMIB.CvpdnTemplateTable})
    cISCOVPDNMGMTMIB.EntityData.Children.Append("cvpdnBundleTable", types.YChild{"CvpdnBundleTable", &cISCOVPDNMGMTMIB.CvpdnBundleTable})
    cISCOVPDNMGMTMIB.EntityData.Children.Append("cvpdnBundleChildTable", types.YChild{"CvpdnBundleChildTable", &cISCOVPDNMGMTMIB.CvpdnBundleChildTable})
    cISCOVPDNMGMTMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOVPDNMGMTMIB.EntityData.YListKeys = []string {}

    return &(cISCOVPDNMGMTMIB.EntityData)
}

// CISCOVPDNMGMTMIB_CiscoVpdnMgmtMIBNotifs
type CISCOVPDNMGMTMIB_CiscoVpdnMgmtMIBNotifs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object contains the local session ID of the L2X session for which this
    // notification has been generated. The type is interface{} with range:
    // 0..65535.
    CvpdnNotifSessionID interface{}

    // Indicates the event that generated the L2X session notification.  The
    // events are represented as follows:  up:     Session has come up.  down:  
    // Session has gone down.  pwUp:   Pseudowire associated with this         
    // session has come up.  pwDown: Pseudowire associated with this         
    // session has gone down. The type is CvpdnNotifSessionEvent.
    CvpdnNotifSessionEvent interface{}
}

func (ciscoVpdnMgmtMIBNotifs *CISCOVPDNMGMTMIB_CiscoVpdnMgmtMIBNotifs) GetEntityData() *types.CommonEntityData {
    ciscoVpdnMgmtMIBNotifs.EntityData.YFilter = ciscoVpdnMgmtMIBNotifs.YFilter
    ciscoVpdnMgmtMIBNotifs.EntityData.YangName = "ciscoVpdnMgmtMIBNotifs"
    ciscoVpdnMgmtMIBNotifs.EntityData.BundleName = "cisco_ios_xe"
    ciscoVpdnMgmtMIBNotifs.EntityData.ParentYangName = "CISCO-VPDN-MGMT-MIB"
    ciscoVpdnMgmtMIBNotifs.EntityData.SegmentPath = "ciscoVpdnMgmtMIBNotifs"
    ciscoVpdnMgmtMIBNotifs.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoVpdnMgmtMIBNotifs.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoVpdnMgmtMIBNotifs.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoVpdnMgmtMIBNotifs.EntityData.Children = types.NewOrderedMap()
    ciscoVpdnMgmtMIBNotifs.EntityData.Leafs = types.NewOrderedMap()
    ciscoVpdnMgmtMIBNotifs.EntityData.Leafs.Append("cvpdnNotifSessionID", types.YLeaf{"CvpdnNotifSessionID", ciscoVpdnMgmtMIBNotifs.CvpdnNotifSessionID})
    ciscoVpdnMgmtMIBNotifs.EntityData.Leafs.Append("cvpdnNotifSessionEvent", types.YLeaf{"CvpdnNotifSessionEvent", ciscoVpdnMgmtMIBNotifs.CvpdnNotifSessionEvent})

    ciscoVpdnMgmtMIBNotifs.EntityData.YListKeys = []string {}

    return &(ciscoVpdnMgmtMIBNotifs.EntityData)
}

// CISCOVPDNMGMTMIB_CiscoVpdnMgmtMIBNotifs_CvpdnNotifSessionEvent represents         session has gone down.
type CISCOVPDNMGMTMIB_CiscoVpdnMgmtMIBNotifs_CvpdnNotifSessionEvent string

const (
    CISCOVPDNMGMTMIB_CiscoVpdnMgmtMIBNotifs_CvpdnNotifSessionEvent_up CISCOVPDNMGMTMIB_CiscoVpdnMgmtMIBNotifs_CvpdnNotifSessionEvent = "up"

    CISCOVPDNMGMTMIB_CiscoVpdnMgmtMIBNotifs_CvpdnNotifSessionEvent_down CISCOVPDNMGMTMIB_CiscoVpdnMgmtMIBNotifs_CvpdnNotifSessionEvent = "down"

    CISCOVPDNMGMTMIB_CiscoVpdnMgmtMIBNotifs_CvpdnNotifSessionEvent_pwUp CISCOVPDNMGMTMIB_CiscoVpdnMgmtMIBNotifs_CvpdnNotifSessionEvent = "pwUp"

    CISCOVPDNMGMTMIB_CiscoVpdnMgmtMIBNotifs_CvpdnNotifSessionEvent_pwDown CISCOVPDNMGMTMIB_CiscoVpdnMgmtMIBNotifs_CvpdnNotifSessionEvent = "pwDown"
)

// CISCOVPDNMGMTMIB_CvpdnSystemInfo
type CISCOVPDNMGMTMIB_CvpdnSystemInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The total number of VPDN tunnels that are currently active within this
    // system. The type is interface{} with range: 0..4294967295. Units are
    // tunnels.
    CvpdnTunnelTotal interface{}

    // The total number of active users in all the active VPDN tunnels within this
    // system. The type is interface{} with range: 0..4294967295. Units are users.
    CvpdnSessionTotal interface{}

    // The total number of denied user attempts to all the active VPDN tunnels
    // within this system. The type is interface{} with range: 0..4294967295.
    // Units are attempts.
    CvpdnDeniedUsersTotal interface{}

    // Indicates whether Layer 2 VPN session notifications are enabled. The type
    // is bool.
    CvpdnSystemNotifSessionEnabled interface{}

    // Clears all the sessions in a given tunnel type.  When reading this object,
    // the value of 'none' will always be returned.  When setting these values,
    // the following operations will be performed:      none: no operation.     
    // all:  clears all the sessions in all the tunnels.      l2f:  clears all the
    // L2F sessions.      l2tp: clears all the L2TP sessions.      pptp: clears
    // all the PPTP sessions. The type is CvpdnSystemClearSessions.
    CvpdnSystemClearSessions interface{}
}

func (cvpdnSystemInfo *CISCOVPDNMGMTMIB_CvpdnSystemInfo) GetEntityData() *types.CommonEntityData {
    cvpdnSystemInfo.EntityData.YFilter = cvpdnSystemInfo.YFilter
    cvpdnSystemInfo.EntityData.YangName = "cvpdnSystemInfo"
    cvpdnSystemInfo.EntityData.BundleName = "cisco_ios_xe"
    cvpdnSystemInfo.EntityData.ParentYangName = "CISCO-VPDN-MGMT-MIB"
    cvpdnSystemInfo.EntityData.SegmentPath = "cvpdnSystemInfo"
    cvpdnSystemInfo.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvpdnSystemInfo.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvpdnSystemInfo.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvpdnSystemInfo.EntityData.Children = types.NewOrderedMap()
    cvpdnSystemInfo.EntityData.Leafs = types.NewOrderedMap()
    cvpdnSystemInfo.EntityData.Leafs.Append("cvpdnTunnelTotal", types.YLeaf{"CvpdnTunnelTotal", cvpdnSystemInfo.CvpdnTunnelTotal})
    cvpdnSystemInfo.EntityData.Leafs.Append("cvpdnSessionTotal", types.YLeaf{"CvpdnSessionTotal", cvpdnSystemInfo.CvpdnSessionTotal})
    cvpdnSystemInfo.EntityData.Leafs.Append("cvpdnDeniedUsersTotal", types.YLeaf{"CvpdnDeniedUsersTotal", cvpdnSystemInfo.CvpdnDeniedUsersTotal})
    cvpdnSystemInfo.EntityData.Leafs.Append("cvpdnSystemNotifSessionEnabled", types.YLeaf{"CvpdnSystemNotifSessionEnabled", cvpdnSystemInfo.CvpdnSystemNotifSessionEnabled})
    cvpdnSystemInfo.EntityData.Leafs.Append("cvpdnSystemClearSessions", types.YLeaf{"CvpdnSystemClearSessions", cvpdnSystemInfo.CvpdnSystemClearSessions})

    cvpdnSystemInfo.EntityData.YListKeys = []string {}

    return &(cvpdnSystemInfo.EntityData)
}

// CISCOVPDNMGMTMIB_CvpdnSystemInfo_CvpdnSystemClearSessions represents     pptp: clears all the PPTP sessions.
type CISCOVPDNMGMTMIB_CvpdnSystemInfo_CvpdnSystemClearSessions string

const (
    CISCOVPDNMGMTMIB_CvpdnSystemInfo_CvpdnSystemClearSessions_none CISCOVPDNMGMTMIB_CvpdnSystemInfo_CvpdnSystemClearSessions = "none"

    CISCOVPDNMGMTMIB_CvpdnSystemInfo_CvpdnSystemClearSessions_all CISCOVPDNMGMTMIB_CvpdnSystemInfo_CvpdnSystemClearSessions = "all"

    CISCOVPDNMGMTMIB_CvpdnSystemInfo_CvpdnSystemClearSessions_l2f CISCOVPDNMGMTMIB_CvpdnSystemInfo_CvpdnSystemClearSessions = "l2f"

    CISCOVPDNMGMTMIB_CvpdnSystemInfo_CvpdnSystemClearSessions_l2tp CISCOVPDNMGMTMIB_CvpdnSystemInfo_CvpdnSystemClearSessions = "l2tp"

    CISCOVPDNMGMTMIB_CvpdnSystemInfo_CvpdnSystemClearSessions_pptp CISCOVPDNMGMTMIB_CvpdnSystemInfo_CvpdnSystemClearSessions = "pptp"
)

// CISCOVPDNMGMTMIB_CvpdnMultilinkInfo
type CISCOVPDNMGMTMIB_CvpdnMultilinkInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The total number of bundles comprised of a single link. The type is
    // interface{} with range: 0..4294967295.
    CvpdnBundlesWithOneLink interface{}

    // The total number of bundles comprised of two links. The type is interface{}
    // with range: 0..4294967295.
    CvpdnBundlesWithTwoLinks interface{}

    // The total number of bundles comprised of more than two links. The type is
    // interface{} with range: 0..4294967295.
    CvpdnBundlesWithMoreThanTwoLinks interface{}

    // The value of the sysUpTime object when the contents of cvpdnBundleTable
    // last changed. The type is interface{} with range: 0..4294967295.
    CvpdnBundleLastChanged interface{}
}

func (cvpdnMultilinkInfo *CISCOVPDNMGMTMIB_CvpdnMultilinkInfo) GetEntityData() *types.CommonEntityData {
    cvpdnMultilinkInfo.EntityData.YFilter = cvpdnMultilinkInfo.YFilter
    cvpdnMultilinkInfo.EntityData.YangName = "cvpdnMultilinkInfo"
    cvpdnMultilinkInfo.EntityData.BundleName = "cisco_ios_xe"
    cvpdnMultilinkInfo.EntityData.ParentYangName = "CISCO-VPDN-MGMT-MIB"
    cvpdnMultilinkInfo.EntityData.SegmentPath = "cvpdnMultilinkInfo"
    cvpdnMultilinkInfo.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvpdnMultilinkInfo.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvpdnMultilinkInfo.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvpdnMultilinkInfo.EntityData.Children = types.NewOrderedMap()
    cvpdnMultilinkInfo.EntityData.Leafs = types.NewOrderedMap()
    cvpdnMultilinkInfo.EntityData.Leafs.Append("cvpdnBundlesWithOneLink", types.YLeaf{"CvpdnBundlesWithOneLink", cvpdnMultilinkInfo.CvpdnBundlesWithOneLink})
    cvpdnMultilinkInfo.EntityData.Leafs.Append("cvpdnBundlesWithTwoLinks", types.YLeaf{"CvpdnBundlesWithTwoLinks", cvpdnMultilinkInfo.CvpdnBundlesWithTwoLinks})
    cvpdnMultilinkInfo.EntityData.Leafs.Append("cvpdnBundlesWithMoreThanTwoLinks", types.YLeaf{"CvpdnBundlesWithMoreThanTwoLinks", cvpdnMultilinkInfo.CvpdnBundlesWithMoreThanTwoLinks})
    cvpdnMultilinkInfo.EntityData.Leafs.Append("cvpdnBundleLastChanged", types.YLeaf{"CvpdnBundleLastChanged", cvpdnMultilinkInfo.CvpdnBundleLastChanged})

    cvpdnMultilinkInfo.EntityData.YListKeys = []string {}

    return &(cvpdnMultilinkInfo.EntityData)
}

// CISCOVPDNMGMTMIB_CvpdnSystemTable
// Table of information about the VPDN system for all tunnel
// types.
type CISCOVPDNMGMTMIB_CvpdnSystemTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table, containing information about a single type of VPDN
    // tunnel. The type is slice of
    // CISCOVPDNMGMTMIB_CvpdnSystemTable_CvpdnSystemEntry.
    CvpdnSystemEntry []*CISCOVPDNMGMTMIB_CvpdnSystemTable_CvpdnSystemEntry
}

func (cvpdnSystemTable *CISCOVPDNMGMTMIB_CvpdnSystemTable) GetEntityData() *types.CommonEntityData {
    cvpdnSystemTable.EntityData.YFilter = cvpdnSystemTable.YFilter
    cvpdnSystemTable.EntityData.YangName = "cvpdnSystemTable"
    cvpdnSystemTable.EntityData.BundleName = "cisco_ios_xe"
    cvpdnSystemTable.EntityData.ParentYangName = "CISCO-VPDN-MGMT-MIB"
    cvpdnSystemTable.EntityData.SegmentPath = "cvpdnSystemTable"
    cvpdnSystemTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvpdnSystemTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvpdnSystemTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvpdnSystemTable.EntityData.Children = types.NewOrderedMap()
    cvpdnSystemTable.EntityData.Children.Append("cvpdnSystemEntry", types.YChild{"CvpdnSystemEntry", nil})
    for i := range cvpdnSystemTable.CvpdnSystemEntry {
        cvpdnSystemTable.EntityData.Children.Append(types.GetSegmentPath(cvpdnSystemTable.CvpdnSystemEntry[i]), types.YChild{"CvpdnSystemEntry", cvpdnSystemTable.CvpdnSystemEntry[i]})
    }
    cvpdnSystemTable.EntityData.Leafs = types.NewOrderedMap()

    cvpdnSystemTable.EntityData.YListKeys = []string {}

    return &(cvpdnSystemTable.EntityData)
}

// CISCOVPDNMGMTMIB_CvpdnSystemTable_CvpdnSystemEntry
// An entry in the table, containing information about a
// single type of VPDN tunnel.
type CISCOVPDNMGMTMIB_CvpdnSystemTable_CvpdnSystemEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The tunnel type.  This is the tunnel protocol. The
    // type is TunnelType.
    CvpdnSystemTunnelType interface{}

    // The total number of VPDN tunnels that are currently active of this tunnel
    // type. The type is interface{} with range: 0..4294967295. Units are tunnels.
    CvpdnSystemTunnelTotal interface{}

    // The total number of active sessions in all the active VPDN tunnels of this
    // tunnel type. The type is interface{} with range: 0..4294967295. Units are
    // sessions.
    CvpdnSystemSessionTotal interface{}

    // The total number of denied user attempts to all the VPDN tunnels of this
    // tunnel type since last system re-initialization. The type is interface{}
    // with range: 0..4294967295. Units are attempts.
    CvpdnSystemDeniedUsersTotal interface{}

    // The total number tunnel connection attempts on all the VPDN tunnels of this
    // tunnel type since last system re-initialization. The type is interface{}
    // with range: 0..4294967295. Units are attempts.
    CvpdnSystemInitialConnReq interface{}

    // The total number tunnel Successful connection attempts in VPDN tunnels of
    // this tunnel type since last system re-initialization. The type is
    // interface{} with range: 0..4294967295. Units are attempts.
    CvpdnSystemSuccessConnReq interface{}

    // The total number tunnel Failed connection attempts in VPDN tunnels of this
    // tunnel type since last system re-initialization. The type is interface{}
    // with range: 0..4294967295. Units are attempts.
    CvpdnSystemFailedConnReq interface{}
}

func (cvpdnSystemEntry *CISCOVPDNMGMTMIB_CvpdnSystemTable_CvpdnSystemEntry) GetEntityData() *types.CommonEntityData {
    cvpdnSystemEntry.EntityData.YFilter = cvpdnSystemEntry.YFilter
    cvpdnSystemEntry.EntityData.YangName = "cvpdnSystemEntry"
    cvpdnSystemEntry.EntityData.BundleName = "cisco_ios_xe"
    cvpdnSystemEntry.EntityData.ParentYangName = "cvpdnSystemTable"
    cvpdnSystemEntry.EntityData.SegmentPath = "cvpdnSystemEntry" + types.AddKeyToken(cvpdnSystemEntry.CvpdnSystemTunnelType, "cvpdnSystemTunnelType")
    cvpdnSystemEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvpdnSystemEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvpdnSystemEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvpdnSystemEntry.EntityData.Children = types.NewOrderedMap()
    cvpdnSystemEntry.EntityData.Leafs = types.NewOrderedMap()
    cvpdnSystemEntry.EntityData.Leafs.Append("cvpdnSystemTunnelType", types.YLeaf{"CvpdnSystemTunnelType", cvpdnSystemEntry.CvpdnSystemTunnelType})
    cvpdnSystemEntry.EntityData.Leafs.Append("cvpdnSystemTunnelTotal", types.YLeaf{"CvpdnSystemTunnelTotal", cvpdnSystemEntry.CvpdnSystemTunnelTotal})
    cvpdnSystemEntry.EntityData.Leafs.Append("cvpdnSystemSessionTotal", types.YLeaf{"CvpdnSystemSessionTotal", cvpdnSystemEntry.CvpdnSystemSessionTotal})
    cvpdnSystemEntry.EntityData.Leafs.Append("cvpdnSystemDeniedUsersTotal", types.YLeaf{"CvpdnSystemDeniedUsersTotal", cvpdnSystemEntry.CvpdnSystemDeniedUsersTotal})
    cvpdnSystemEntry.EntityData.Leafs.Append("cvpdnSystemInitialConnReq", types.YLeaf{"CvpdnSystemInitialConnReq", cvpdnSystemEntry.CvpdnSystemInitialConnReq})
    cvpdnSystemEntry.EntityData.Leafs.Append("cvpdnSystemSuccessConnReq", types.YLeaf{"CvpdnSystemSuccessConnReq", cvpdnSystemEntry.CvpdnSystemSuccessConnReq})
    cvpdnSystemEntry.EntityData.Leafs.Append("cvpdnSystemFailedConnReq", types.YLeaf{"CvpdnSystemFailedConnReq", cvpdnSystemEntry.CvpdnSystemFailedConnReq})

    cvpdnSystemEntry.EntityData.YListKeys = []string {"CvpdnSystemTunnelType"}

    return &(cvpdnSystemEntry.EntityData)
}

// CISCOVPDNMGMTMIB_CvpdnTunnelTable
// Table of information about the active VPDN tunnels.
type CISCOVPDNMGMTMIB_CvpdnTunnelTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table, containing information about a single active VPDN
    // tunnel. The type is slice of
    // CISCOVPDNMGMTMIB_CvpdnTunnelTable_CvpdnTunnelEntry.
    CvpdnTunnelEntry []*CISCOVPDNMGMTMIB_CvpdnTunnelTable_CvpdnTunnelEntry
}

func (cvpdnTunnelTable *CISCOVPDNMGMTMIB_CvpdnTunnelTable) GetEntityData() *types.CommonEntityData {
    cvpdnTunnelTable.EntityData.YFilter = cvpdnTunnelTable.YFilter
    cvpdnTunnelTable.EntityData.YangName = "cvpdnTunnelTable"
    cvpdnTunnelTable.EntityData.BundleName = "cisco_ios_xe"
    cvpdnTunnelTable.EntityData.ParentYangName = "CISCO-VPDN-MGMT-MIB"
    cvpdnTunnelTable.EntityData.SegmentPath = "cvpdnTunnelTable"
    cvpdnTunnelTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvpdnTunnelTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvpdnTunnelTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvpdnTunnelTable.EntityData.Children = types.NewOrderedMap()
    cvpdnTunnelTable.EntityData.Children.Append("cvpdnTunnelEntry", types.YChild{"CvpdnTunnelEntry", nil})
    for i := range cvpdnTunnelTable.CvpdnTunnelEntry {
        cvpdnTunnelTable.EntityData.Children.Append(types.GetSegmentPath(cvpdnTunnelTable.CvpdnTunnelEntry[i]), types.YChild{"CvpdnTunnelEntry", cvpdnTunnelTable.CvpdnTunnelEntry[i]})
    }
    cvpdnTunnelTable.EntityData.Leafs = types.NewOrderedMap()

    cvpdnTunnelTable.EntityData.YListKeys = []string {}

    return &(cvpdnTunnelTable.EntityData)
}

// CISCOVPDNMGMTMIB_CvpdnTunnelTable_CvpdnTunnelEntry
// An entry in the table, containing information about a
// single active VPDN tunnel.
type CISCOVPDNMGMTMIB_CvpdnTunnelTable_CvpdnTunnelEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The Tunnel ID of an active VPDN tunnel.  If it is
    // the instigator of the tunnel, the ID is the HGW/LNS tunnel ID, otherwise it
    // is the NAS/LAC tunnel ID. The type is interface{} with range:
    // 0..4294967295.
    CvpdnTunnelTunnelId interface{}

    // The remote Tunnel ID of an active VPDN tunnel.  If it is the instigator of
    // the tunnel, the ID is the NAS/LAC tunnel ID, otherwise it is the HGW/LNS
    // tunnel ID. The type is interface{} with range: 0..4294967295.
    CvpdnTunnelRemoteTunnelId interface{}

    // The local name of an active VPDN tunnel.  It will be the NAS/LAC name of
    // the tunnel if the router serves as the NAS/LAC, or the HGW/LNS name of the
    // tunnel if the system serves as the home gateway.  Typically, the local name
    // is the configured host name of the router. The type is string with length:
    // 1..255.
    CvpdnTunnelLocalName interface{}

    // The remote name of an active VPDN tunnel.  It will be the home gateway name
    // of the tunnel if the system is a NAS/LAC, or the NAS/LAC name of the tunnel
    // if the system serves as the home gateway. The type is string with length:
    // 1..255.
    CvpdnTunnelRemoteName interface{}

    // The remote end point name of an active VPDN tunnel. This name is either the
    // domain name or the DNIS that this tunnel is projected with. The type is
    // string with length: 1..255.
    CvpdnTunnelRemoteEndpointName interface{}

    // This object indicates whether the tunnel was generated locally or not. The
    // type is bool.
    CvpdnTunnelLocalInitConnection interface{}

    // The cause which originated an active VPDN tunnel.  The tunnel can be
    // projected via domain name, DNIS or a stack group (SGBP). The type is
    // CvpdnTunnelOrigCause.
    CvpdnTunnelOrigCause interface{}

    // The current state of an active VPDN tunnel.  Each state code is explained
    // below:         unknown: The current state of the tunnel is                
    // unknown.         opening: The tunnel has just been instigated and          
    // is pending for a remote end reply to                 complete the process. 
    // open:    The tunnel is active.         closing: The tunnel has just been
    // shut down and                 is pending for the remote end to reply       
    // to complete the process. The type is CvpdnTunnelState.
    CvpdnTunnelState interface{}

    // The total number of active session currently in the tunnel. The type is
    // interface{} with range: 0..4294967295. Units are sessions.
    CvpdnTunnelActiveSessions interface{}

    // A count of the accumulated total of denied users for the tunnel. The type
    // is interface{} with range: 0..4294967295. Units are calls.
    CvpdnTunnelDeniedUsers interface{}

    // A VPDN tunnel can be put into a soft shut state to prevent any new user
    // session to be added.  This object specifies whether this tunnel has been
    // soft shut. The type is bool.
    CvpdnTunnelSoftshut interface{}

    // The type of network service used in the active tunnel. For now it is IP
    // only. The type is CvpdnTunnelNetworkServiceType.
    CvpdnTunnelNetworkServiceType interface{}

    // The local IP address of the tunnel.  This IP address is that of the
    // interface at the local end of the tunnel. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    CvpdnTunnelLocalIpAddress interface{}

    // The source IP address of the tunnel.  This IP address is the user
    // configurable IP address for Stack Group Biding Protocol (SGBP) via the CLI
    // command: vpdn source-ip. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    CvpdnTunnelSourceIpAddress interface{}

    // The remote IP address of the tunnel.  This IP address is that of the
    // interface at the remote end of the tunnel. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    CvpdnTunnelRemoteIpAddress interface{}
}

func (cvpdnTunnelEntry *CISCOVPDNMGMTMIB_CvpdnTunnelTable_CvpdnTunnelEntry) GetEntityData() *types.CommonEntityData {
    cvpdnTunnelEntry.EntityData.YFilter = cvpdnTunnelEntry.YFilter
    cvpdnTunnelEntry.EntityData.YangName = "cvpdnTunnelEntry"
    cvpdnTunnelEntry.EntityData.BundleName = "cisco_ios_xe"
    cvpdnTunnelEntry.EntityData.ParentYangName = "cvpdnTunnelTable"
    cvpdnTunnelEntry.EntityData.SegmentPath = "cvpdnTunnelEntry" + types.AddKeyToken(cvpdnTunnelEntry.CvpdnTunnelTunnelId, "cvpdnTunnelTunnelId")
    cvpdnTunnelEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvpdnTunnelEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvpdnTunnelEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvpdnTunnelEntry.EntityData.Children = types.NewOrderedMap()
    cvpdnTunnelEntry.EntityData.Leafs = types.NewOrderedMap()
    cvpdnTunnelEntry.EntityData.Leafs.Append("cvpdnTunnelTunnelId", types.YLeaf{"CvpdnTunnelTunnelId", cvpdnTunnelEntry.CvpdnTunnelTunnelId})
    cvpdnTunnelEntry.EntityData.Leafs.Append("cvpdnTunnelRemoteTunnelId", types.YLeaf{"CvpdnTunnelRemoteTunnelId", cvpdnTunnelEntry.CvpdnTunnelRemoteTunnelId})
    cvpdnTunnelEntry.EntityData.Leafs.Append("cvpdnTunnelLocalName", types.YLeaf{"CvpdnTunnelLocalName", cvpdnTunnelEntry.CvpdnTunnelLocalName})
    cvpdnTunnelEntry.EntityData.Leafs.Append("cvpdnTunnelRemoteName", types.YLeaf{"CvpdnTunnelRemoteName", cvpdnTunnelEntry.CvpdnTunnelRemoteName})
    cvpdnTunnelEntry.EntityData.Leafs.Append("cvpdnTunnelRemoteEndpointName", types.YLeaf{"CvpdnTunnelRemoteEndpointName", cvpdnTunnelEntry.CvpdnTunnelRemoteEndpointName})
    cvpdnTunnelEntry.EntityData.Leafs.Append("cvpdnTunnelLocalInitConnection", types.YLeaf{"CvpdnTunnelLocalInitConnection", cvpdnTunnelEntry.CvpdnTunnelLocalInitConnection})
    cvpdnTunnelEntry.EntityData.Leafs.Append("cvpdnTunnelOrigCause", types.YLeaf{"CvpdnTunnelOrigCause", cvpdnTunnelEntry.CvpdnTunnelOrigCause})
    cvpdnTunnelEntry.EntityData.Leafs.Append("cvpdnTunnelState", types.YLeaf{"CvpdnTunnelState", cvpdnTunnelEntry.CvpdnTunnelState})
    cvpdnTunnelEntry.EntityData.Leafs.Append("cvpdnTunnelActiveSessions", types.YLeaf{"CvpdnTunnelActiveSessions", cvpdnTunnelEntry.CvpdnTunnelActiveSessions})
    cvpdnTunnelEntry.EntityData.Leafs.Append("cvpdnTunnelDeniedUsers", types.YLeaf{"CvpdnTunnelDeniedUsers", cvpdnTunnelEntry.CvpdnTunnelDeniedUsers})
    cvpdnTunnelEntry.EntityData.Leafs.Append("cvpdnTunnelSoftshut", types.YLeaf{"CvpdnTunnelSoftshut", cvpdnTunnelEntry.CvpdnTunnelSoftshut})
    cvpdnTunnelEntry.EntityData.Leafs.Append("cvpdnTunnelNetworkServiceType", types.YLeaf{"CvpdnTunnelNetworkServiceType", cvpdnTunnelEntry.CvpdnTunnelNetworkServiceType})
    cvpdnTunnelEntry.EntityData.Leafs.Append("cvpdnTunnelLocalIpAddress", types.YLeaf{"CvpdnTunnelLocalIpAddress", cvpdnTunnelEntry.CvpdnTunnelLocalIpAddress})
    cvpdnTunnelEntry.EntityData.Leafs.Append("cvpdnTunnelSourceIpAddress", types.YLeaf{"CvpdnTunnelSourceIpAddress", cvpdnTunnelEntry.CvpdnTunnelSourceIpAddress})
    cvpdnTunnelEntry.EntityData.Leafs.Append("cvpdnTunnelRemoteIpAddress", types.YLeaf{"CvpdnTunnelRemoteIpAddress", cvpdnTunnelEntry.CvpdnTunnelRemoteIpAddress})

    cvpdnTunnelEntry.EntityData.YListKeys = []string {"CvpdnTunnelTunnelId"}

    return &(cvpdnTunnelEntry.EntityData)
}

// CISCOVPDNMGMTMIB_CvpdnTunnelTable_CvpdnTunnelEntry_CvpdnTunnelNetworkServiceType represents For now it is IP only.
type CISCOVPDNMGMTMIB_CvpdnTunnelTable_CvpdnTunnelEntry_CvpdnTunnelNetworkServiceType string

const (
    CISCOVPDNMGMTMIB_CvpdnTunnelTable_CvpdnTunnelEntry_CvpdnTunnelNetworkServiceType_ip CISCOVPDNMGMTMIB_CvpdnTunnelTable_CvpdnTunnelEntry_CvpdnTunnelNetworkServiceType = "ip"
)

// CISCOVPDNMGMTMIB_CvpdnTunnelTable_CvpdnTunnelEntry_CvpdnTunnelOrigCause represents stack group (SGBP).
type CISCOVPDNMGMTMIB_CvpdnTunnelTable_CvpdnTunnelEntry_CvpdnTunnelOrigCause string

const (
    CISCOVPDNMGMTMIB_CvpdnTunnelTable_CvpdnTunnelEntry_CvpdnTunnelOrigCause_domain CISCOVPDNMGMTMIB_CvpdnTunnelTable_CvpdnTunnelEntry_CvpdnTunnelOrigCause = "domain"

    CISCOVPDNMGMTMIB_CvpdnTunnelTable_CvpdnTunnelEntry_CvpdnTunnelOrigCause_dnis CISCOVPDNMGMTMIB_CvpdnTunnelTable_CvpdnTunnelEntry_CvpdnTunnelOrigCause = "dnis"

    CISCOVPDNMGMTMIB_CvpdnTunnelTable_CvpdnTunnelEntry_CvpdnTunnelOrigCause_stack CISCOVPDNMGMTMIB_CvpdnTunnelTable_CvpdnTunnelEntry_CvpdnTunnelOrigCause = "stack"
)

// CISCOVPDNMGMTMIB_CvpdnTunnelTable_CvpdnTunnelEntry_CvpdnTunnelState represents                 to complete the process.
type CISCOVPDNMGMTMIB_CvpdnTunnelTable_CvpdnTunnelEntry_CvpdnTunnelState string

const (
    CISCOVPDNMGMTMIB_CvpdnTunnelTable_CvpdnTunnelEntry_CvpdnTunnelState_unknown CISCOVPDNMGMTMIB_CvpdnTunnelTable_CvpdnTunnelEntry_CvpdnTunnelState = "unknown"

    CISCOVPDNMGMTMIB_CvpdnTunnelTable_CvpdnTunnelEntry_CvpdnTunnelState_opening CISCOVPDNMGMTMIB_CvpdnTunnelTable_CvpdnTunnelEntry_CvpdnTunnelState = "opening"

    CISCOVPDNMGMTMIB_CvpdnTunnelTable_CvpdnTunnelEntry_CvpdnTunnelState_open CISCOVPDNMGMTMIB_CvpdnTunnelTable_CvpdnTunnelEntry_CvpdnTunnelState = "open"

    CISCOVPDNMGMTMIB_CvpdnTunnelTable_CvpdnTunnelEntry_CvpdnTunnelState_closing CISCOVPDNMGMTMIB_CvpdnTunnelTable_CvpdnTunnelEntry_CvpdnTunnelState = "closing"
)

// CISCOVPDNMGMTMIB_CvpdnTunnelAttrTable
// Table of information about the active VPDN tunnels.  An
// entry is added to the table when a new tunnel is initiated
// and removed from the table when the tunnel is terminated.
type CISCOVPDNMGMTMIB_CvpdnTunnelAttrTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table, containing information about a single active VPDN
    // tunnel. The type is slice of
    // CISCOVPDNMGMTMIB_CvpdnTunnelAttrTable_CvpdnTunnelAttrEntry.
    CvpdnTunnelAttrEntry []*CISCOVPDNMGMTMIB_CvpdnTunnelAttrTable_CvpdnTunnelAttrEntry
}

func (cvpdnTunnelAttrTable *CISCOVPDNMGMTMIB_CvpdnTunnelAttrTable) GetEntityData() *types.CommonEntityData {
    cvpdnTunnelAttrTable.EntityData.YFilter = cvpdnTunnelAttrTable.YFilter
    cvpdnTunnelAttrTable.EntityData.YangName = "cvpdnTunnelAttrTable"
    cvpdnTunnelAttrTable.EntityData.BundleName = "cisco_ios_xe"
    cvpdnTunnelAttrTable.EntityData.ParentYangName = "CISCO-VPDN-MGMT-MIB"
    cvpdnTunnelAttrTable.EntityData.SegmentPath = "cvpdnTunnelAttrTable"
    cvpdnTunnelAttrTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvpdnTunnelAttrTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvpdnTunnelAttrTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvpdnTunnelAttrTable.EntityData.Children = types.NewOrderedMap()
    cvpdnTunnelAttrTable.EntityData.Children.Append("cvpdnTunnelAttrEntry", types.YChild{"CvpdnTunnelAttrEntry", nil})
    for i := range cvpdnTunnelAttrTable.CvpdnTunnelAttrEntry {
        cvpdnTunnelAttrTable.EntityData.Children.Append(types.GetSegmentPath(cvpdnTunnelAttrTable.CvpdnTunnelAttrEntry[i]), types.YChild{"CvpdnTunnelAttrEntry", cvpdnTunnelAttrTable.CvpdnTunnelAttrEntry[i]})
    }
    cvpdnTunnelAttrTable.EntityData.Leafs = types.NewOrderedMap()

    cvpdnTunnelAttrTable.EntityData.YListKeys = []string {}

    return &(cvpdnTunnelAttrTable.EntityData)
}

// CISCOVPDNMGMTMIB_CvpdnTunnelAttrTable_CvpdnTunnelAttrEntry
// An entry in the table, containing information about a
// single active VPDN tunnel.
type CISCOVPDNMGMTMIB_CvpdnTunnelAttrTable_CvpdnTunnelAttrEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is TunnelType. Refers to
    // cisco_vpdn_mgmt_mib.CISCOVPDNMGMTMIB_CvpdnSystemTable_CvpdnSystemEntry_CvpdnSystemTunnelType
    CvpdnSystemTunnelType interface{}

    // This attribute is a key. The Tunnel ID of an active VPDN tunnel.  If this
    // end is the instigator of the tunnel, the ID is the TS tunnel ID, otherwise
    // it is the NAS tunnel ID.  Two distinct tunnels with the same tunnel ID may
    // exist, but with different tunnel types. The type is interface{} with range:
    // 0..65535.
    CvpdnTunnelAttrTunnelId interface{}

    // The remote Tunnel ID of an active VPDN tunnel.  If this end is the
    // instigator of the tunnel, the ID is the NAS tunnel ID, otherwise it is the
    // TS tunnel ID. The type is interface{} with range: 0..65535.
    CvpdnTunnelAttrRemoteTunnelId interface{}

    // The local name of an active VPDN tunnel.  It will be the NAS name of the
    // tunnel if the system serves as the NAS, or the TS name of the tunnel if the
    // system serves as the tunnel server.  Typically, the local name is the
    // configured host name of the system. The type is string with length: 1..255.
    CvpdnTunnelAttrLocalName interface{}

    // The remote name of an active VPDN tunnel.  It will be the tunnel server
    // name of the tunnel if the system is a NAS, or the NAS name of the tunnel if
    // the system serves as the tunnel server. The type is string with length:
    // 1..255.
    CvpdnTunnelAttrRemoteName interface{}

    // The remote end point name of an active VPDN tunnel.  This name is either
    // the domain name or the DNIS that this tunnel is projected with. The type is
    // string with length: 1..255.
    CvpdnTunnelAttrRemoteEndpointName interface{}

    // This object indicates whether the tunnel was originated locally or not.  If
    // it's true, the tunnel was originated locally. The type is bool.
    CvpdnTunnelAttrLocalInitConnection interface{}

    // The cause which originated an active VPDN tunnel.  The tunnel can be
    // projected via domain name, DNIS, stack group, or L2 Xconnect. The type is
    // CvpdnTunnelAttrOrigCause.
    CvpdnTunnelAttrOrigCause interface{}

    // The current state of an active VPDN tunnel. Tunnels of type l2f will have
    // states with the 'l2f' prefix. Tunnels of type l2tp will have states with
    // the 'l2tp' prefix. Tunnels of type pptp will have states with the 'pptp'
    // prefix.  Each state code is explained below:      unknown:            The
    // current state of the tunnel is                         unknown.     
    // l2fOpening:         The tunnel has just been initiated                     
    // and is pending for a remote end                         reply to complete
    // the process.      l2fOpenWait:        This end received a tunnel open      
    // request from the remote end and is                         waiting for the
    // tunnel to be                         established.      l2fOpen:           
    // The tunnel is active.      l2fClosing:         This end received a tunnel
    // close                         request.      l2fCloseWait:       The tunnel
    // has just been shut down                         and is pending for the
    // remote end                         to reply to complete the process.     
    // l2tpIdle:           No tunnel is initiated yet.      l2tpWaitCtlReply:  
    // The tunnel has been initiated and                         is pending for a
    // remote end reply                         to complete the process.     
    // l2tpEstablished:    The tunnel is active.      l2tpShuttingDown:   The
    // tunnel is in progress of                         shutting down.     
    // l2tpNoSessionLeft:  There is no session left in the                        
    // tunnel.      pptpIdle:           No tunnel is initiated yet.     
    // pptpWaitConnect:    The tunnel is waiting for a TCP                        
    // connection.      pptpWaitCtlRequest: The tunnel has been initiated and     
    // is pending for a remote end                         request.     
    // pptpWaitCtlReply:   The tunnel has been initiated and                      
    // is pending for a remote end reply.      pptpEstablished:    The tunnel is
    // active.      pptpWaitStopReply:  The tunnel is being shut down and         
    // is pending for a remote end reply.      pptpTerminal:       The tunnel has
    // been shut down. The type is CvpdnTunnelAttrState.
    CvpdnTunnelAttrState interface{}

    // The total number of active session currently in the tunnel. The type is
    // interface{} with range: 0..4294967295. Units are sessions.
    CvpdnTunnelAttrActiveSessions interface{}

    // A count of the accumulated total of denied users for the tunnel. The type
    // is interface{} with range: 0..4294967295. Units are calls.
    CvpdnTunnelAttrDeniedUsers interface{}

    // A VPDN tunnel can be put into a soft shut state to prevent any new session
    // to be added.  This object specifies whether this tunnel has been soft shut.
    // If its true, it has been soft shut. The type is bool.
    CvpdnTunnelAttrSoftshut interface{}

    // The type of network service used in the active tunnel. The type is
    // CvpdnTunnelAttrNetworkServiceType.
    CvpdnTunnelAttrNetworkServiceType interface{}

    // The local IP address of the tunnel.  This IP address is that of the
    // interface at the local end of the tunnel. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    CvpdnTunnelAttrLocalIpAddress interface{}

    // The source IP address of the tunnel.  This IP address is the user
    // configurable IP address for Stack Group Biding Protocol. The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    CvpdnTunnelAttrSourceIpAddress interface{}

    // The remote IP address of the tunnel.  This IP address is that of the
    // interface at the remote end of the tunnel. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    CvpdnTunnelAttrRemoteIpAddress interface{}

    // Indicates the type of address contained in cvpdnTunnelAttrLocalInetAddress.
    // The type is InetAddressType.
    CvpdnTunnelAttrLocalInetAddressType interface{}

    // The local IP address of the tunnel.  This IP address is that of the
    // interface at the local end of the tunnel. The type of this address is
    // determined by the value of  cvpdnTunnelAttrLocalInetAddressType. The type
    // is string with length: 0..255.
    CvpdnTunnelAttrLocalInetAddress interface{}

    // Indicates the type of address contained in
    // cvpdnTunnelAttrSourceInetAddress. The type is InetAddressType.
    CvpdnTunnelAttrSourceInetAddressType interface{}

    // The source IP address of the tunnel.  This IP address is the user
    // configurable IP address for Stack Group Biding Protocol.  The type of this
    // address is determined by the  value of
    // cvpdnTunnelAttrSourceInetAddressType. The type is string with length:
    // 0..255.
    CvpdnTunnelAttrSourceInetAddress interface{}

    // Indicates the type of address contained in
    // cvpdnTunnelAttrRemoteInetAddress. The type is InetAddressType.
    CvpdnTunnelAttrRemoteInetAddressType interface{}

    // The remote IP address of the tunnel.  This IP address is that of the
    // interface at the remote end of the tunnel. The type of this address is
    // determined by the value of  cvpdnTunnelAttrRemoteInetAddressType. The type
    // is string with length: 0..255.
    CvpdnTunnelAttrRemoteInetAddress interface{}
}

func (cvpdnTunnelAttrEntry *CISCOVPDNMGMTMIB_CvpdnTunnelAttrTable_CvpdnTunnelAttrEntry) GetEntityData() *types.CommonEntityData {
    cvpdnTunnelAttrEntry.EntityData.YFilter = cvpdnTunnelAttrEntry.YFilter
    cvpdnTunnelAttrEntry.EntityData.YangName = "cvpdnTunnelAttrEntry"
    cvpdnTunnelAttrEntry.EntityData.BundleName = "cisco_ios_xe"
    cvpdnTunnelAttrEntry.EntityData.ParentYangName = "cvpdnTunnelAttrTable"
    cvpdnTunnelAttrEntry.EntityData.SegmentPath = "cvpdnTunnelAttrEntry" + types.AddKeyToken(cvpdnTunnelAttrEntry.CvpdnSystemTunnelType, "cvpdnSystemTunnelType") + types.AddKeyToken(cvpdnTunnelAttrEntry.CvpdnTunnelAttrTunnelId, "cvpdnTunnelAttrTunnelId")
    cvpdnTunnelAttrEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvpdnTunnelAttrEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvpdnTunnelAttrEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvpdnTunnelAttrEntry.EntityData.Children = types.NewOrderedMap()
    cvpdnTunnelAttrEntry.EntityData.Leafs = types.NewOrderedMap()
    cvpdnTunnelAttrEntry.EntityData.Leafs.Append("cvpdnSystemTunnelType", types.YLeaf{"CvpdnSystemTunnelType", cvpdnTunnelAttrEntry.CvpdnSystemTunnelType})
    cvpdnTunnelAttrEntry.EntityData.Leafs.Append("cvpdnTunnelAttrTunnelId", types.YLeaf{"CvpdnTunnelAttrTunnelId", cvpdnTunnelAttrEntry.CvpdnTunnelAttrTunnelId})
    cvpdnTunnelAttrEntry.EntityData.Leafs.Append("cvpdnTunnelAttrRemoteTunnelId", types.YLeaf{"CvpdnTunnelAttrRemoteTunnelId", cvpdnTunnelAttrEntry.CvpdnTunnelAttrRemoteTunnelId})
    cvpdnTunnelAttrEntry.EntityData.Leafs.Append("cvpdnTunnelAttrLocalName", types.YLeaf{"CvpdnTunnelAttrLocalName", cvpdnTunnelAttrEntry.CvpdnTunnelAttrLocalName})
    cvpdnTunnelAttrEntry.EntityData.Leafs.Append("cvpdnTunnelAttrRemoteName", types.YLeaf{"CvpdnTunnelAttrRemoteName", cvpdnTunnelAttrEntry.CvpdnTunnelAttrRemoteName})
    cvpdnTunnelAttrEntry.EntityData.Leafs.Append("cvpdnTunnelAttrRemoteEndpointName", types.YLeaf{"CvpdnTunnelAttrRemoteEndpointName", cvpdnTunnelAttrEntry.CvpdnTunnelAttrRemoteEndpointName})
    cvpdnTunnelAttrEntry.EntityData.Leafs.Append("cvpdnTunnelAttrLocalInitConnection", types.YLeaf{"CvpdnTunnelAttrLocalInitConnection", cvpdnTunnelAttrEntry.CvpdnTunnelAttrLocalInitConnection})
    cvpdnTunnelAttrEntry.EntityData.Leafs.Append("cvpdnTunnelAttrOrigCause", types.YLeaf{"CvpdnTunnelAttrOrigCause", cvpdnTunnelAttrEntry.CvpdnTunnelAttrOrigCause})
    cvpdnTunnelAttrEntry.EntityData.Leafs.Append("cvpdnTunnelAttrState", types.YLeaf{"CvpdnTunnelAttrState", cvpdnTunnelAttrEntry.CvpdnTunnelAttrState})
    cvpdnTunnelAttrEntry.EntityData.Leafs.Append("cvpdnTunnelAttrActiveSessions", types.YLeaf{"CvpdnTunnelAttrActiveSessions", cvpdnTunnelAttrEntry.CvpdnTunnelAttrActiveSessions})
    cvpdnTunnelAttrEntry.EntityData.Leafs.Append("cvpdnTunnelAttrDeniedUsers", types.YLeaf{"CvpdnTunnelAttrDeniedUsers", cvpdnTunnelAttrEntry.CvpdnTunnelAttrDeniedUsers})
    cvpdnTunnelAttrEntry.EntityData.Leafs.Append("cvpdnTunnelAttrSoftshut", types.YLeaf{"CvpdnTunnelAttrSoftshut", cvpdnTunnelAttrEntry.CvpdnTunnelAttrSoftshut})
    cvpdnTunnelAttrEntry.EntityData.Leafs.Append("cvpdnTunnelAttrNetworkServiceType", types.YLeaf{"CvpdnTunnelAttrNetworkServiceType", cvpdnTunnelAttrEntry.CvpdnTunnelAttrNetworkServiceType})
    cvpdnTunnelAttrEntry.EntityData.Leafs.Append("cvpdnTunnelAttrLocalIpAddress", types.YLeaf{"CvpdnTunnelAttrLocalIpAddress", cvpdnTunnelAttrEntry.CvpdnTunnelAttrLocalIpAddress})
    cvpdnTunnelAttrEntry.EntityData.Leafs.Append("cvpdnTunnelAttrSourceIpAddress", types.YLeaf{"CvpdnTunnelAttrSourceIpAddress", cvpdnTunnelAttrEntry.CvpdnTunnelAttrSourceIpAddress})
    cvpdnTunnelAttrEntry.EntityData.Leafs.Append("cvpdnTunnelAttrRemoteIpAddress", types.YLeaf{"CvpdnTunnelAttrRemoteIpAddress", cvpdnTunnelAttrEntry.CvpdnTunnelAttrRemoteIpAddress})
    cvpdnTunnelAttrEntry.EntityData.Leafs.Append("cvpdnTunnelAttrLocalInetAddressType", types.YLeaf{"CvpdnTunnelAttrLocalInetAddressType", cvpdnTunnelAttrEntry.CvpdnTunnelAttrLocalInetAddressType})
    cvpdnTunnelAttrEntry.EntityData.Leafs.Append("cvpdnTunnelAttrLocalInetAddress", types.YLeaf{"CvpdnTunnelAttrLocalInetAddress", cvpdnTunnelAttrEntry.CvpdnTunnelAttrLocalInetAddress})
    cvpdnTunnelAttrEntry.EntityData.Leafs.Append("cvpdnTunnelAttrSourceInetAddressType", types.YLeaf{"CvpdnTunnelAttrSourceInetAddressType", cvpdnTunnelAttrEntry.CvpdnTunnelAttrSourceInetAddressType})
    cvpdnTunnelAttrEntry.EntityData.Leafs.Append("cvpdnTunnelAttrSourceInetAddress", types.YLeaf{"CvpdnTunnelAttrSourceInetAddress", cvpdnTunnelAttrEntry.CvpdnTunnelAttrSourceInetAddress})
    cvpdnTunnelAttrEntry.EntityData.Leafs.Append("cvpdnTunnelAttrRemoteInetAddressType", types.YLeaf{"CvpdnTunnelAttrRemoteInetAddressType", cvpdnTunnelAttrEntry.CvpdnTunnelAttrRemoteInetAddressType})
    cvpdnTunnelAttrEntry.EntityData.Leafs.Append("cvpdnTunnelAttrRemoteInetAddress", types.YLeaf{"CvpdnTunnelAttrRemoteInetAddress", cvpdnTunnelAttrEntry.CvpdnTunnelAttrRemoteInetAddress})

    cvpdnTunnelAttrEntry.EntityData.YListKeys = []string {"CvpdnSystemTunnelType", "CvpdnTunnelAttrTunnelId"}

    return &(cvpdnTunnelAttrEntry.EntityData)
}

// CISCOVPDNMGMTMIB_CvpdnTunnelAttrTable_CvpdnTunnelAttrEntry_CvpdnTunnelAttrNetworkServiceType represents The type of network service used in the active tunnel.
type CISCOVPDNMGMTMIB_CvpdnTunnelAttrTable_CvpdnTunnelAttrEntry_CvpdnTunnelAttrNetworkServiceType string

const (
    CISCOVPDNMGMTMIB_CvpdnTunnelAttrTable_CvpdnTunnelAttrEntry_CvpdnTunnelAttrNetworkServiceType_ip CISCOVPDNMGMTMIB_CvpdnTunnelAttrTable_CvpdnTunnelAttrEntry_CvpdnTunnelAttrNetworkServiceType = "ip"
)

// CISCOVPDNMGMTMIB_CvpdnTunnelAttrTable_CvpdnTunnelAttrEntry_CvpdnTunnelAttrOrigCause represents or L2 Xconnect.
type CISCOVPDNMGMTMIB_CvpdnTunnelAttrTable_CvpdnTunnelAttrEntry_CvpdnTunnelAttrOrigCause string

const (
    CISCOVPDNMGMTMIB_CvpdnTunnelAttrTable_CvpdnTunnelAttrEntry_CvpdnTunnelAttrOrigCause_domain CISCOVPDNMGMTMIB_CvpdnTunnelAttrTable_CvpdnTunnelAttrEntry_CvpdnTunnelAttrOrigCause = "domain"

    CISCOVPDNMGMTMIB_CvpdnTunnelAttrTable_CvpdnTunnelAttrEntry_CvpdnTunnelAttrOrigCause_dnis CISCOVPDNMGMTMIB_CvpdnTunnelAttrTable_CvpdnTunnelAttrEntry_CvpdnTunnelAttrOrigCause = "dnis"

    CISCOVPDNMGMTMIB_CvpdnTunnelAttrTable_CvpdnTunnelAttrEntry_CvpdnTunnelAttrOrigCause_stack CISCOVPDNMGMTMIB_CvpdnTunnelAttrTable_CvpdnTunnelAttrEntry_CvpdnTunnelAttrOrigCause = "stack"

    CISCOVPDNMGMTMIB_CvpdnTunnelAttrTable_CvpdnTunnelAttrEntry_CvpdnTunnelAttrOrigCause_xconnect CISCOVPDNMGMTMIB_CvpdnTunnelAttrTable_CvpdnTunnelAttrEntry_CvpdnTunnelAttrOrigCause = "xconnect"
)

// CISCOVPDNMGMTMIB_CvpdnTunnelAttrTable_CvpdnTunnelAttrEntry_CvpdnTunnelAttrState represents     pptpTerminal:       The tunnel has been shut down.
type CISCOVPDNMGMTMIB_CvpdnTunnelAttrTable_CvpdnTunnelAttrEntry_CvpdnTunnelAttrState string

const (
    CISCOVPDNMGMTMIB_CvpdnTunnelAttrTable_CvpdnTunnelAttrEntry_CvpdnTunnelAttrState_unknown CISCOVPDNMGMTMIB_CvpdnTunnelAttrTable_CvpdnTunnelAttrEntry_CvpdnTunnelAttrState = "unknown"

    CISCOVPDNMGMTMIB_CvpdnTunnelAttrTable_CvpdnTunnelAttrEntry_CvpdnTunnelAttrState_l2fOpening CISCOVPDNMGMTMIB_CvpdnTunnelAttrTable_CvpdnTunnelAttrEntry_CvpdnTunnelAttrState = "l2fOpening"

    CISCOVPDNMGMTMIB_CvpdnTunnelAttrTable_CvpdnTunnelAttrEntry_CvpdnTunnelAttrState_l2fOpenWait CISCOVPDNMGMTMIB_CvpdnTunnelAttrTable_CvpdnTunnelAttrEntry_CvpdnTunnelAttrState = "l2fOpenWait"

    CISCOVPDNMGMTMIB_CvpdnTunnelAttrTable_CvpdnTunnelAttrEntry_CvpdnTunnelAttrState_l2fOpen CISCOVPDNMGMTMIB_CvpdnTunnelAttrTable_CvpdnTunnelAttrEntry_CvpdnTunnelAttrState = "l2fOpen"

    CISCOVPDNMGMTMIB_CvpdnTunnelAttrTable_CvpdnTunnelAttrEntry_CvpdnTunnelAttrState_l2fClosing CISCOVPDNMGMTMIB_CvpdnTunnelAttrTable_CvpdnTunnelAttrEntry_CvpdnTunnelAttrState = "l2fClosing"

    CISCOVPDNMGMTMIB_CvpdnTunnelAttrTable_CvpdnTunnelAttrEntry_CvpdnTunnelAttrState_l2fCloseWait CISCOVPDNMGMTMIB_CvpdnTunnelAttrTable_CvpdnTunnelAttrEntry_CvpdnTunnelAttrState = "l2fCloseWait"

    CISCOVPDNMGMTMIB_CvpdnTunnelAttrTable_CvpdnTunnelAttrEntry_CvpdnTunnelAttrState_l2tpIdle CISCOVPDNMGMTMIB_CvpdnTunnelAttrTable_CvpdnTunnelAttrEntry_CvpdnTunnelAttrState = "l2tpIdle"

    CISCOVPDNMGMTMIB_CvpdnTunnelAttrTable_CvpdnTunnelAttrEntry_CvpdnTunnelAttrState_l2tpWaitCtlReply CISCOVPDNMGMTMIB_CvpdnTunnelAttrTable_CvpdnTunnelAttrEntry_CvpdnTunnelAttrState = "l2tpWaitCtlReply"

    CISCOVPDNMGMTMIB_CvpdnTunnelAttrTable_CvpdnTunnelAttrEntry_CvpdnTunnelAttrState_l2tpEstablished CISCOVPDNMGMTMIB_CvpdnTunnelAttrTable_CvpdnTunnelAttrEntry_CvpdnTunnelAttrState = "l2tpEstablished"

    CISCOVPDNMGMTMIB_CvpdnTunnelAttrTable_CvpdnTunnelAttrEntry_CvpdnTunnelAttrState_l2tpShuttingDown CISCOVPDNMGMTMIB_CvpdnTunnelAttrTable_CvpdnTunnelAttrEntry_CvpdnTunnelAttrState = "l2tpShuttingDown"

    CISCOVPDNMGMTMIB_CvpdnTunnelAttrTable_CvpdnTunnelAttrEntry_CvpdnTunnelAttrState_l2tpNoSessionLeft CISCOVPDNMGMTMIB_CvpdnTunnelAttrTable_CvpdnTunnelAttrEntry_CvpdnTunnelAttrState = "l2tpNoSessionLeft"

    CISCOVPDNMGMTMIB_CvpdnTunnelAttrTable_CvpdnTunnelAttrEntry_CvpdnTunnelAttrState_pptpIdle CISCOVPDNMGMTMIB_CvpdnTunnelAttrTable_CvpdnTunnelAttrEntry_CvpdnTunnelAttrState = "pptpIdle"

    CISCOVPDNMGMTMIB_CvpdnTunnelAttrTable_CvpdnTunnelAttrEntry_CvpdnTunnelAttrState_pptpWaitConnect CISCOVPDNMGMTMIB_CvpdnTunnelAttrTable_CvpdnTunnelAttrEntry_CvpdnTunnelAttrState = "pptpWaitConnect"

    CISCOVPDNMGMTMIB_CvpdnTunnelAttrTable_CvpdnTunnelAttrEntry_CvpdnTunnelAttrState_pptpWaitCtlRequest CISCOVPDNMGMTMIB_CvpdnTunnelAttrTable_CvpdnTunnelAttrEntry_CvpdnTunnelAttrState = "pptpWaitCtlRequest"

    CISCOVPDNMGMTMIB_CvpdnTunnelAttrTable_CvpdnTunnelAttrEntry_CvpdnTunnelAttrState_pptpWaitCtlReply CISCOVPDNMGMTMIB_CvpdnTunnelAttrTable_CvpdnTunnelAttrEntry_CvpdnTunnelAttrState = "pptpWaitCtlReply"

    CISCOVPDNMGMTMIB_CvpdnTunnelAttrTable_CvpdnTunnelAttrEntry_CvpdnTunnelAttrState_pptpEstablished CISCOVPDNMGMTMIB_CvpdnTunnelAttrTable_CvpdnTunnelAttrEntry_CvpdnTunnelAttrState = "pptpEstablished"

    CISCOVPDNMGMTMIB_CvpdnTunnelAttrTable_CvpdnTunnelAttrEntry_CvpdnTunnelAttrState_pptpWaitStopReply CISCOVPDNMGMTMIB_CvpdnTunnelAttrTable_CvpdnTunnelAttrEntry_CvpdnTunnelAttrState = "pptpWaitStopReply"

    CISCOVPDNMGMTMIB_CvpdnTunnelAttrTable_CvpdnTunnelAttrEntry_CvpdnTunnelAttrState_pptpTerminal CISCOVPDNMGMTMIB_CvpdnTunnelAttrTable_CvpdnTunnelAttrEntry_CvpdnTunnelAttrState = "pptpTerminal"
)

// CISCOVPDNMGMTMIB_CvpdnTunnelSessionTable
// Table of information about individual user sessions
// within the active tunnels.  Entry is added to the table
// when new user session is initiated and be removed from
// the table when the user session is terminated.
type CISCOVPDNMGMTMIB_CvpdnTunnelSessionTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table, containing information about a single user session
    // within the tunnel. The type is slice of
    // CISCOVPDNMGMTMIB_CvpdnTunnelSessionTable_CvpdnTunnelSessionEntry.
    CvpdnTunnelSessionEntry []*CISCOVPDNMGMTMIB_CvpdnTunnelSessionTable_CvpdnTunnelSessionEntry
}

func (cvpdnTunnelSessionTable *CISCOVPDNMGMTMIB_CvpdnTunnelSessionTable) GetEntityData() *types.CommonEntityData {
    cvpdnTunnelSessionTable.EntityData.YFilter = cvpdnTunnelSessionTable.YFilter
    cvpdnTunnelSessionTable.EntityData.YangName = "cvpdnTunnelSessionTable"
    cvpdnTunnelSessionTable.EntityData.BundleName = "cisco_ios_xe"
    cvpdnTunnelSessionTable.EntityData.ParentYangName = "CISCO-VPDN-MGMT-MIB"
    cvpdnTunnelSessionTable.EntityData.SegmentPath = "cvpdnTunnelSessionTable"
    cvpdnTunnelSessionTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvpdnTunnelSessionTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvpdnTunnelSessionTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvpdnTunnelSessionTable.EntityData.Children = types.NewOrderedMap()
    cvpdnTunnelSessionTable.EntityData.Children.Append("cvpdnTunnelSessionEntry", types.YChild{"CvpdnTunnelSessionEntry", nil})
    for i := range cvpdnTunnelSessionTable.CvpdnTunnelSessionEntry {
        cvpdnTunnelSessionTable.EntityData.Children.Append(types.GetSegmentPath(cvpdnTunnelSessionTable.CvpdnTunnelSessionEntry[i]), types.YChild{"CvpdnTunnelSessionEntry", cvpdnTunnelSessionTable.CvpdnTunnelSessionEntry[i]})
    }
    cvpdnTunnelSessionTable.EntityData.Leafs = types.NewOrderedMap()

    cvpdnTunnelSessionTable.EntityData.YListKeys = []string {}

    return &(cvpdnTunnelSessionTable.EntityData)
}

// CISCOVPDNMGMTMIB_CvpdnTunnelSessionTable_CvpdnTunnelSessionEntry
// An entry in the table, containing information about a
// single user session within the tunnel.
type CISCOVPDNMGMTMIB_CvpdnTunnelSessionTable_CvpdnTunnelSessionEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // cisco_vpdn_mgmt_mib.CISCOVPDNMGMTMIB_CvpdnTunnelTable_CvpdnTunnelEntry_CvpdnTunnelTunnelId
    CvpdnTunnelTunnelId interface{}

    // This attribute is a key. The ID of an active VPDN tunnel user session. The
    // type is interface{} with range: 0..4294967295.
    CvpdnTunnelSessionId interface{}

    // The name of the user of the user session. The type is string with length:
    // 1..255.
    CvpdnTunnelSessionUserName interface{}

    // The current state of an active user session.  Each state code is explained
    // below:      unknown:          The current state of the tunnel's            
    // session is unknown.      opening:          The user session has just been  
    // initiated through a tunnel and is                       pending for the
    // remote end reply                       to complete the process.      open: 
    // The user session is active.      closing:          The user session has
    // just been                       closed and is pending for the              
    // remote end reply to complete the                       process.     
    // waitingForTunnel: The user session is in this state                      
    // when the tunnel which this session                       is going through
    // is still in                       CLOSED state.  It waits for the          
    // tunnel to become OPEN before the                       session is allow to
    // be fully                       established. The type is
    // CvpdnTunnelSessionState.
    CvpdnTunnelSessionState interface{}

    // This object specifies the call duration of the current active user session
    // in value of system uptime. The type is interface{} with range:
    // 0..4294967295.
    CvpdnTunnelSessionCallDuration interface{}

    // The total number of output packets through this active user session. The
    // type is interface{} with range: 0..4294967295. Units are packets.
    CvpdnTunnelSessionPacketsOut interface{}

    // The total number of output bytes through this active user session. The type
    // is interface{} with range: 0..4294967295. Units are bytes.
    CvpdnTunnelSessionBytesOut interface{}

    // The total number of input packets through this active user session. The
    // type is interface{} with range: 0..4294967295. Units are packets.
    CvpdnTunnelSessionPacketsIn interface{}

    // The total number of input bytes through this active user session. The type
    // is interface{} with range: 0..4294967295. Units are bytes.
    CvpdnTunnelSessionBytesIn interface{}

    // The type of physical devices that this user session is attached to for the
    // local end of the tunnel.  The meaning of each device type is explained
    // below:      other:              Any device that has not been               
    // defined.      asyncInternalModem: Modem Pool device of an access           
    // server.      async:              A regular asynchronous serial             
    // interface.      sync:               A regular synchronous serial           
    // interface.      bchan:              An ISDN call.      xdsl:              
    // Future application with xDSL                         devices.      cable:  
    // Future application with Cable                         modem devices. The
    // type is CvpdnTunnelSessionDeviceType.
    CvpdnTunnelSessionDeviceType interface{}

    // The incoming caller identification of the user.  It is the originating
    // number that called into the device that initiates the user session.  This
    // object can be empty since not all dial device can provide caller ID
    // information. The type is string.
    CvpdnTunnelSessionDeviceCallerId interface{}

    // The device ID of the physical interface for the user session.  The object
    // is the the interface index which points to the ifTable.  For virtual
    // interface that is not in the ifTable, it will have zero value. The type is
    // interface{} with range: 0..2147483647.
    CvpdnTunnelSessionDevicePhyId interface{}

    // This object indicates whether the session is part of a multilink or not.
    // The type is bool.
    CvpdnTunnelSessionMultilink interface{}

    // The Modem Pool database slot index if it is associated with this user
    // session.  Only a session with device of type asyncInternalModem will have a
    // valid value for this object. The type is interface{} with range:
    // 0..4294967295.
    CvpdnTunnelSessionModemSlotIndex interface{}

    // The Modem Pool database port index if it is associated with this user
    // session.  Only a session with a device of type asyncInternalModem will have
    // a valid value for this object. The type is interface{} with range:
    // 0..4294967295.
    CvpdnTunnelSessionModemPortIndex interface{}

    // The DS1 database slot index if it is associated with this user session. 
    // Only a session with a device of type asyncInternalModem will have a valid
    // value for this object. The type is interface{} with range: 0..4294967295.
    CvpdnTunnelSessionDS1SlotIndex interface{}

    // The DS1 database port index if it is associated with this user session. 
    // Only a session with a device of type asyncInternalModem will have a a valid
    // value for this object. The type is interface{} with range: 0..4294967295.
    CvpdnTunnelSessionDS1PortIndex interface{}

    // The DS1 database channel index if it is associated with this user session. 
    // Only a session over a device of type asyncInternalModem will have a valid
    // value for this object. The type is interface{} with range: 0..4294967295.
    CvpdnTunnelSessionDS1ChannelIndex interface{}

    // The start time of the current modem call.  Only a session with a  device of
    // type asyncInternalModem will have a valid value for this object. The type
    // is interface{} with range: 0..4294967295.
    CvpdnTunnelSessionModemCallStartTime interface{}

    // Arbitrary small integer to distinguish modem calls that occurred at the
    // same time tick.  Only session over device asyncInternalModem will have a
    // valid value for this object. The type is interface{} with range:
    // 0..4294967295.
    CvpdnTunnelSessionModemCallStartIndex interface{}
}

func (cvpdnTunnelSessionEntry *CISCOVPDNMGMTMIB_CvpdnTunnelSessionTable_CvpdnTunnelSessionEntry) GetEntityData() *types.CommonEntityData {
    cvpdnTunnelSessionEntry.EntityData.YFilter = cvpdnTunnelSessionEntry.YFilter
    cvpdnTunnelSessionEntry.EntityData.YangName = "cvpdnTunnelSessionEntry"
    cvpdnTunnelSessionEntry.EntityData.BundleName = "cisco_ios_xe"
    cvpdnTunnelSessionEntry.EntityData.ParentYangName = "cvpdnTunnelSessionTable"
    cvpdnTunnelSessionEntry.EntityData.SegmentPath = "cvpdnTunnelSessionEntry" + types.AddKeyToken(cvpdnTunnelSessionEntry.CvpdnTunnelTunnelId, "cvpdnTunnelTunnelId") + types.AddKeyToken(cvpdnTunnelSessionEntry.CvpdnTunnelSessionId, "cvpdnTunnelSessionId")
    cvpdnTunnelSessionEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvpdnTunnelSessionEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvpdnTunnelSessionEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvpdnTunnelSessionEntry.EntityData.Children = types.NewOrderedMap()
    cvpdnTunnelSessionEntry.EntityData.Leafs = types.NewOrderedMap()
    cvpdnTunnelSessionEntry.EntityData.Leafs.Append("cvpdnTunnelTunnelId", types.YLeaf{"CvpdnTunnelTunnelId", cvpdnTunnelSessionEntry.CvpdnTunnelTunnelId})
    cvpdnTunnelSessionEntry.EntityData.Leafs.Append("cvpdnTunnelSessionId", types.YLeaf{"CvpdnTunnelSessionId", cvpdnTunnelSessionEntry.CvpdnTunnelSessionId})
    cvpdnTunnelSessionEntry.EntityData.Leafs.Append("cvpdnTunnelSessionUserName", types.YLeaf{"CvpdnTunnelSessionUserName", cvpdnTunnelSessionEntry.CvpdnTunnelSessionUserName})
    cvpdnTunnelSessionEntry.EntityData.Leafs.Append("cvpdnTunnelSessionState", types.YLeaf{"CvpdnTunnelSessionState", cvpdnTunnelSessionEntry.CvpdnTunnelSessionState})
    cvpdnTunnelSessionEntry.EntityData.Leafs.Append("cvpdnTunnelSessionCallDuration", types.YLeaf{"CvpdnTunnelSessionCallDuration", cvpdnTunnelSessionEntry.CvpdnTunnelSessionCallDuration})
    cvpdnTunnelSessionEntry.EntityData.Leafs.Append("cvpdnTunnelSessionPacketsOut", types.YLeaf{"CvpdnTunnelSessionPacketsOut", cvpdnTunnelSessionEntry.CvpdnTunnelSessionPacketsOut})
    cvpdnTunnelSessionEntry.EntityData.Leafs.Append("cvpdnTunnelSessionBytesOut", types.YLeaf{"CvpdnTunnelSessionBytesOut", cvpdnTunnelSessionEntry.CvpdnTunnelSessionBytesOut})
    cvpdnTunnelSessionEntry.EntityData.Leafs.Append("cvpdnTunnelSessionPacketsIn", types.YLeaf{"CvpdnTunnelSessionPacketsIn", cvpdnTunnelSessionEntry.CvpdnTunnelSessionPacketsIn})
    cvpdnTunnelSessionEntry.EntityData.Leafs.Append("cvpdnTunnelSessionBytesIn", types.YLeaf{"CvpdnTunnelSessionBytesIn", cvpdnTunnelSessionEntry.CvpdnTunnelSessionBytesIn})
    cvpdnTunnelSessionEntry.EntityData.Leafs.Append("cvpdnTunnelSessionDeviceType", types.YLeaf{"CvpdnTunnelSessionDeviceType", cvpdnTunnelSessionEntry.CvpdnTunnelSessionDeviceType})
    cvpdnTunnelSessionEntry.EntityData.Leafs.Append("cvpdnTunnelSessionDeviceCallerId", types.YLeaf{"CvpdnTunnelSessionDeviceCallerId", cvpdnTunnelSessionEntry.CvpdnTunnelSessionDeviceCallerId})
    cvpdnTunnelSessionEntry.EntityData.Leafs.Append("cvpdnTunnelSessionDevicePhyId", types.YLeaf{"CvpdnTunnelSessionDevicePhyId", cvpdnTunnelSessionEntry.CvpdnTunnelSessionDevicePhyId})
    cvpdnTunnelSessionEntry.EntityData.Leafs.Append("cvpdnTunnelSessionMultilink", types.YLeaf{"CvpdnTunnelSessionMultilink", cvpdnTunnelSessionEntry.CvpdnTunnelSessionMultilink})
    cvpdnTunnelSessionEntry.EntityData.Leafs.Append("cvpdnTunnelSessionModemSlotIndex", types.YLeaf{"CvpdnTunnelSessionModemSlotIndex", cvpdnTunnelSessionEntry.CvpdnTunnelSessionModemSlotIndex})
    cvpdnTunnelSessionEntry.EntityData.Leafs.Append("cvpdnTunnelSessionModemPortIndex", types.YLeaf{"CvpdnTunnelSessionModemPortIndex", cvpdnTunnelSessionEntry.CvpdnTunnelSessionModemPortIndex})
    cvpdnTunnelSessionEntry.EntityData.Leafs.Append("cvpdnTunnelSessionDS1SlotIndex", types.YLeaf{"CvpdnTunnelSessionDS1SlotIndex", cvpdnTunnelSessionEntry.CvpdnTunnelSessionDS1SlotIndex})
    cvpdnTunnelSessionEntry.EntityData.Leafs.Append("cvpdnTunnelSessionDS1PortIndex", types.YLeaf{"CvpdnTunnelSessionDS1PortIndex", cvpdnTunnelSessionEntry.CvpdnTunnelSessionDS1PortIndex})
    cvpdnTunnelSessionEntry.EntityData.Leafs.Append("cvpdnTunnelSessionDS1ChannelIndex", types.YLeaf{"CvpdnTunnelSessionDS1ChannelIndex", cvpdnTunnelSessionEntry.CvpdnTunnelSessionDS1ChannelIndex})
    cvpdnTunnelSessionEntry.EntityData.Leafs.Append("cvpdnTunnelSessionModemCallStartTime", types.YLeaf{"CvpdnTunnelSessionModemCallStartTime", cvpdnTunnelSessionEntry.CvpdnTunnelSessionModemCallStartTime})
    cvpdnTunnelSessionEntry.EntityData.Leafs.Append("cvpdnTunnelSessionModemCallStartIndex", types.YLeaf{"CvpdnTunnelSessionModemCallStartIndex", cvpdnTunnelSessionEntry.CvpdnTunnelSessionModemCallStartIndex})

    cvpdnTunnelSessionEntry.EntityData.YListKeys = []string {"CvpdnTunnelTunnelId", "CvpdnTunnelSessionId"}

    return &(cvpdnTunnelSessionEntry.EntityData)
}

// CISCOVPDNMGMTMIB_CvpdnTunnelSessionTable_CvpdnTunnelSessionEntry_CvpdnTunnelSessionDeviceType represents                         modem devices.
type CISCOVPDNMGMTMIB_CvpdnTunnelSessionTable_CvpdnTunnelSessionEntry_CvpdnTunnelSessionDeviceType string

const (
    CISCOVPDNMGMTMIB_CvpdnTunnelSessionTable_CvpdnTunnelSessionEntry_CvpdnTunnelSessionDeviceType_other CISCOVPDNMGMTMIB_CvpdnTunnelSessionTable_CvpdnTunnelSessionEntry_CvpdnTunnelSessionDeviceType = "other"

    CISCOVPDNMGMTMIB_CvpdnTunnelSessionTable_CvpdnTunnelSessionEntry_CvpdnTunnelSessionDeviceType_asyncInternalModem CISCOVPDNMGMTMIB_CvpdnTunnelSessionTable_CvpdnTunnelSessionEntry_CvpdnTunnelSessionDeviceType = "asyncInternalModem"

    CISCOVPDNMGMTMIB_CvpdnTunnelSessionTable_CvpdnTunnelSessionEntry_CvpdnTunnelSessionDeviceType_async CISCOVPDNMGMTMIB_CvpdnTunnelSessionTable_CvpdnTunnelSessionEntry_CvpdnTunnelSessionDeviceType = "async"

    CISCOVPDNMGMTMIB_CvpdnTunnelSessionTable_CvpdnTunnelSessionEntry_CvpdnTunnelSessionDeviceType_bchan CISCOVPDNMGMTMIB_CvpdnTunnelSessionTable_CvpdnTunnelSessionEntry_CvpdnTunnelSessionDeviceType = "bchan"

    CISCOVPDNMGMTMIB_CvpdnTunnelSessionTable_CvpdnTunnelSessionEntry_CvpdnTunnelSessionDeviceType_sync CISCOVPDNMGMTMIB_CvpdnTunnelSessionTable_CvpdnTunnelSessionEntry_CvpdnTunnelSessionDeviceType = "sync"

    CISCOVPDNMGMTMIB_CvpdnTunnelSessionTable_CvpdnTunnelSessionEntry_CvpdnTunnelSessionDeviceType_virtualAccess CISCOVPDNMGMTMIB_CvpdnTunnelSessionTable_CvpdnTunnelSessionEntry_CvpdnTunnelSessionDeviceType = "virtualAccess"

    CISCOVPDNMGMTMIB_CvpdnTunnelSessionTable_CvpdnTunnelSessionEntry_CvpdnTunnelSessionDeviceType_xdsl CISCOVPDNMGMTMIB_CvpdnTunnelSessionTable_CvpdnTunnelSessionEntry_CvpdnTunnelSessionDeviceType = "xdsl"

    CISCOVPDNMGMTMIB_CvpdnTunnelSessionTable_CvpdnTunnelSessionEntry_CvpdnTunnelSessionDeviceType_cable CISCOVPDNMGMTMIB_CvpdnTunnelSessionTable_CvpdnTunnelSessionEntry_CvpdnTunnelSessionDeviceType = "cable"
)

// CISCOVPDNMGMTMIB_CvpdnTunnelSessionTable_CvpdnTunnelSessionEntry_CvpdnTunnelSessionState represents                       established.
type CISCOVPDNMGMTMIB_CvpdnTunnelSessionTable_CvpdnTunnelSessionEntry_CvpdnTunnelSessionState string

const (
    CISCOVPDNMGMTMIB_CvpdnTunnelSessionTable_CvpdnTunnelSessionEntry_CvpdnTunnelSessionState_unknown CISCOVPDNMGMTMIB_CvpdnTunnelSessionTable_CvpdnTunnelSessionEntry_CvpdnTunnelSessionState = "unknown"

    CISCOVPDNMGMTMIB_CvpdnTunnelSessionTable_CvpdnTunnelSessionEntry_CvpdnTunnelSessionState_opening CISCOVPDNMGMTMIB_CvpdnTunnelSessionTable_CvpdnTunnelSessionEntry_CvpdnTunnelSessionState = "opening"

    CISCOVPDNMGMTMIB_CvpdnTunnelSessionTable_CvpdnTunnelSessionEntry_CvpdnTunnelSessionState_open CISCOVPDNMGMTMIB_CvpdnTunnelSessionTable_CvpdnTunnelSessionEntry_CvpdnTunnelSessionState = "open"

    CISCOVPDNMGMTMIB_CvpdnTunnelSessionTable_CvpdnTunnelSessionEntry_CvpdnTunnelSessionState_closing CISCOVPDNMGMTMIB_CvpdnTunnelSessionTable_CvpdnTunnelSessionEntry_CvpdnTunnelSessionState = "closing"

    CISCOVPDNMGMTMIB_CvpdnTunnelSessionTable_CvpdnTunnelSessionEntry_CvpdnTunnelSessionState_waitingForTunnel CISCOVPDNMGMTMIB_CvpdnTunnelSessionTable_CvpdnTunnelSessionEntry_CvpdnTunnelSessionState = "waitingForTunnel"
)

// CISCOVPDNMGMTMIB_CvpdnSessionAttrTable
// Table of information about individual sessions within the
// active tunnels.  An entry is added to the table when a new
// session is initiated and removed from the table when the
// session is terminated.
type CISCOVPDNMGMTMIB_CvpdnSessionAttrTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table, containing information about a single session within
    // the tunnel. The type is slice of
    // CISCOVPDNMGMTMIB_CvpdnSessionAttrTable_CvpdnSessionAttrEntry.
    CvpdnSessionAttrEntry []*CISCOVPDNMGMTMIB_CvpdnSessionAttrTable_CvpdnSessionAttrEntry
}

func (cvpdnSessionAttrTable *CISCOVPDNMGMTMIB_CvpdnSessionAttrTable) GetEntityData() *types.CommonEntityData {
    cvpdnSessionAttrTable.EntityData.YFilter = cvpdnSessionAttrTable.YFilter
    cvpdnSessionAttrTable.EntityData.YangName = "cvpdnSessionAttrTable"
    cvpdnSessionAttrTable.EntityData.BundleName = "cisco_ios_xe"
    cvpdnSessionAttrTable.EntityData.ParentYangName = "CISCO-VPDN-MGMT-MIB"
    cvpdnSessionAttrTable.EntityData.SegmentPath = "cvpdnSessionAttrTable"
    cvpdnSessionAttrTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvpdnSessionAttrTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvpdnSessionAttrTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvpdnSessionAttrTable.EntityData.Children = types.NewOrderedMap()
    cvpdnSessionAttrTable.EntityData.Children.Append("cvpdnSessionAttrEntry", types.YChild{"CvpdnSessionAttrEntry", nil})
    for i := range cvpdnSessionAttrTable.CvpdnSessionAttrEntry {
        cvpdnSessionAttrTable.EntityData.Children.Append(types.GetSegmentPath(cvpdnSessionAttrTable.CvpdnSessionAttrEntry[i]), types.YChild{"CvpdnSessionAttrEntry", cvpdnSessionAttrTable.CvpdnSessionAttrEntry[i]})
    }
    cvpdnSessionAttrTable.EntityData.Leafs = types.NewOrderedMap()

    cvpdnSessionAttrTable.EntityData.YListKeys = []string {}

    return &(cvpdnSessionAttrTable.EntityData)
}

// CISCOVPDNMGMTMIB_CvpdnSessionAttrTable_CvpdnSessionAttrEntry
// An entry in the table, containing information about a
// single session within the tunnel.
type CISCOVPDNMGMTMIB_CvpdnSessionAttrTable_CvpdnSessionAttrEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is TunnelType. Refers to
    // cisco_vpdn_mgmt_mib.CISCOVPDNMGMTMIB_CvpdnSystemTable_CvpdnSystemEntry_CvpdnSystemTunnelType
    CvpdnSystemTunnelType interface{}

    // This attribute is a key. The type is string with range: 0..65535. Refers to
    // cisco_vpdn_mgmt_mib.CISCOVPDNMGMTMIB_CvpdnTunnelAttrTable_CvpdnTunnelAttrEntry_CvpdnTunnelAttrTunnelId
    CvpdnTunnelAttrTunnelId interface{}

    // This attribute is a key. The ID of an active VPDN session. The type is
    // interface{} with range: 0..65535.
    CvpdnSessionAttrSessionId interface{}

    // The name of the user of the session. The type is string with length:
    // 1..255.
    CvpdnSessionAttrUserName interface{}

    // The current state of a tunnel session. L2F tunnel sessions will have states
    // with the 'l2f' prefix. L2TP tunnel sessions will have states with the
    // 'l2tp' prefix.  Each state code is explained below:      unknown:          
    // The current state of the tunnel's                          session is
    // unknown.      l2fOpening:          The session has just been               
    // initiated through a tunnel and is                          pending for the
    // remote end reply                          to complete the process.     
    // l2fOpen:             The session is active.      l2fCloseWait:        The
    // session has just been closed                          and is pending for
    // the remote end                          reply to complete the process.     
    // l2fWaitingForTunnel: The session is in this state when                     
    // the tunnel which this session is                          going through is
    // still in CLOSED                          state.  It waits for the tunnel to
    // become OPEN before the session is                          allowed to be
    // fully established.      l2tpIdle:            No session is initiated yet.  
    // l2tpWaitingTunnel:   The session is waiting for the                        
    // tunnel to be established.      l2tpWaitReply:       The session has been
    // initiated and                          is pending for the remote end       
    // reply to complete the process.      l2tpWaitConnect:     This end has
    // acknowledged a                          connection request and is waiting  
    // for the remote end to connect.      l2tpEstablished:     The session is
    // active.      l2tpShuttingDown:    The session is in progress of            
    // shutting down.      pptpWaitVAccess:     The session is waiting for the    
    // creation of a virtual access                          interface.     
    // pptpPacEstablished:  The session is active.      pptpWaitTunnel:      The
    // session is waiting for the                          tunnel to be
    // established.      pptpWaitOCRP:        The session has been initiated and  
    // is pending for the remote end                          reply to complete
    // the process.      pptpPnsEstablished:  The session is active.     
    // pptpWaitCallDisc:    Session shutdown is in progress. The type is
    // CvpdnSessionAttrState.
    CvpdnSessionAttrState interface{}

    // This object specifies the call duration of the current active session. The
    // type is interface{} with range: 0..4294967295.
    CvpdnSessionAttrCallDuration interface{}

    // The total number of output packets through this active session. The type is
    // interface{} with range: 0..4294967295. Units are packets.
    CvpdnSessionAttrPacketsOut interface{}

    // The total number of output bytes through this active session. The type is
    // interface{} with range: 0..4294967295. Units are bytes.
    CvpdnSessionAttrBytesOut interface{}

    // The total number of input packets through this active session. The type is
    // interface{} with range: 0..4294967295. Units are packets.
    CvpdnSessionAttrPacketsIn interface{}

    // The total number of input bytes through this active session. The type is
    // interface{} with range: 0..4294967295. Units are bytes.
    CvpdnSessionAttrBytesIn interface{}

    // The type of physical devices that this session is attached to for the local
    // end of the tunnel.  The meaning of each device type is explained below:    
    // other:              Any device that has not been                        
    // defined.      asyncInternalModem: Modem Pool device of an access           
    // server.      async:              A regular asynchronous serial             
    // interface.      sync:               A regular synchronous serial           
    // interface.      bchan:              An ISDN call.      xdsl:              
    // xDSL interface.      cable:              cable modem interface. The type is
    // CvpdnSessionAttrDeviceType.
    CvpdnSessionAttrDeviceType interface{}

    // The incoming caller identification of the user.  It is the originating
    // number that called into the device that initiates the session.  This object
    // can be empty since not all dial devices can provide caller ID information.
    // The type is string.
    CvpdnSessionAttrDeviceCallerId interface{}

    // The device ID of the physical interface for the session. The object is the
    // the interface index which points to the ifTable.  For virtual interfaces
    // that are not in the ifTable, the value will be zero. The type is
    // interface{} with range: 0..2147483647.
    CvpdnSessionAttrDevicePhyId interface{}

    // This object indicates whether the session is part of a multilink PPP
    // bundle, even if there is only one link or session in the bundle.  If it is
    // multilink PPP, the value is true. The type is bool.
    CvpdnSessionAttrMultilink interface{}

    // The Modem Pool database slot index if it is associated with this session. 
    // Only a session with device of type 'asyncInternalModem' will have a valid
    // value for this object; otherwise, it is not instantiated. The type is
    // interface{} with range: 0..4294967295.
    CvpdnSessionAttrModemSlotIndex interface{}

    // The Modem Pool database port index if it is associated with this session. 
    // Only a session with a device of type 'asyncInternalModem' will have a valid
    // value for this object; otherwise, it is not instantiated. The type is
    // interface{} with range: 0..4294967295.
    CvpdnSessionAttrModemPortIndex interface{}

    // The DS1 database slot index if it is associated with this session.  Only a
    // session with a device of type 'asyncInternalModem' will have a valid value
    // for this object; otherwise it is not instantiated. The type is interface{}
    // with range: 0..4294967295.
    CvpdnSessionAttrDS1SlotIndex interface{}

    // The DS1 database port index if it is associated with this session.  Only a
    // session with a device of type 'asyncInternalModem' will have a valid value
    // for this object; otherwise it is not instantiated. The type is interface{}
    // with range: 0..4294967295.
    CvpdnSessionAttrDS1PortIndex interface{}

    // The DS1 database channel index if it is associated with this session.  Only
    // a session over a device of type 'asyncInternalModem' will have a valid
    // value for this object; otherwise it is not instantiated. The type is
    // interface{} with range: 0..4294967295.
    CvpdnSessionAttrDS1ChannelIndex interface{}

    // The start time of the current modem call.  Only a session with a device of
    // type 'asyncInternalModem' will have a valid value for this object;
    // otherwise, it is not instantiated. The type is interface{} with range:
    // 0..4294967295.
    CvpdnSessionAttrModemCallStartTime interface{}

    // Arbitrary small integer to distinguish modem calls that occurred at the
    // same time tick.  Only session over device 'asyncInternalModem' will have a
    // valid value for this object; otherwise, it is not instantiated. The type is
    // interface{} with range: 0..4294967295.
    CvpdnSessionAttrModemCallStartIndex interface{}

    // The virtual circuit ID of an active Layer 2 VPN session. The type is
    // interface{} with range: 0..4294967295.
    CvpdnSessionAttrVirtualCircuitID interface{}

    // The total number of dropped packets that could not be sent into this active
    // session. The type is interface{} with range: 0..4294967295. Units are
    // packets.
    CvpdnSessionAttrSentPktsDropped interface{}

    // The total number of dropped packets that were received from this active
    // session. The type is interface{} with range: 0..4294967295. Units are
    // packets.
    CvpdnSessionAttrRecvPktsDropped interface{}

    // This object specifies the name of the multilink bundle to which this
    // session belongs.  The value of this object is only valid when the value of
    // cvpdnSessionAttrMultilink is 'true'.  If the value of
    // cvpdnSessionAttrMultilink is 'false', then the value of this object will be
    // the empty string. The type is string with length: 0..255.
    CvpdnSessionAttrMultilinkBundle interface{}

    // This object specifies the ifIndex of the multilink bundle to which this
    // session belongs.  The value of this object is only valid when the value of
    // cvpdnSessionAttrMultilink is 'true'.  If the value of
    // cvpdnSessionAttrMultilink is 'false', then the value of this object will be
    // zero. The type is interface{} with range: 0..2147483647.
    CvpdnSessionAttrMultilinkIfIndex interface{}
}

func (cvpdnSessionAttrEntry *CISCOVPDNMGMTMIB_CvpdnSessionAttrTable_CvpdnSessionAttrEntry) GetEntityData() *types.CommonEntityData {
    cvpdnSessionAttrEntry.EntityData.YFilter = cvpdnSessionAttrEntry.YFilter
    cvpdnSessionAttrEntry.EntityData.YangName = "cvpdnSessionAttrEntry"
    cvpdnSessionAttrEntry.EntityData.BundleName = "cisco_ios_xe"
    cvpdnSessionAttrEntry.EntityData.ParentYangName = "cvpdnSessionAttrTable"
    cvpdnSessionAttrEntry.EntityData.SegmentPath = "cvpdnSessionAttrEntry" + types.AddKeyToken(cvpdnSessionAttrEntry.CvpdnSystemTunnelType, "cvpdnSystemTunnelType") + types.AddKeyToken(cvpdnSessionAttrEntry.CvpdnTunnelAttrTunnelId, "cvpdnTunnelAttrTunnelId") + types.AddKeyToken(cvpdnSessionAttrEntry.CvpdnSessionAttrSessionId, "cvpdnSessionAttrSessionId")
    cvpdnSessionAttrEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvpdnSessionAttrEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvpdnSessionAttrEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvpdnSessionAttrEntry.EntityData.Children = types.NewOrderedMap()
    cvpdnSessionAttrEntry.EntityData.Leafs = types.NewOrderedMap()
    cvpdnSessionAttrEntry.EntityData.Leafs.Append("cvpdnSystemTunnelType", types.YLeaf{"CvpdnSystemTunnelType", cvpdnSessionAttrEntry.CvpdnSystemTunnelType})
    cvpdnSessionAttrEntry.EntityData.Leafs.Append("cvpdnTunnelAttrTunnelId", types.YLeaf{"CvpdnTunnelAttrTunnelId", cvpdnSessionAttrEntry.CvpdnTunnelAttrTunnelId})
    cvpdnSessionAttrEntry.EntityData.Leafs.Append("cvpdnSessionAttrSessionId", types.YLeaf{"CvpdnSessionAttrSessionId", cvpdnSessionAttrEntry.CvpdnSessionAttrSessionId})
    cvpdnSessionAttrEntry.EntityData.Leafs.Append("cvpdnSessionAttrUserName", types.YLeaf{"CvpdnSessionAttrUserName", cvpdnSessionAttrEntry.CvpdnSessionAttrUserName})
    cvpdnSessionAttrEntry.EntityData.Leafs.Append("cvpdnSessionAttrState", types.YLeaf{"CvpdnSessionAttrState", cvpdnSessionAttrEntry.CvpdnSessionAttrState})
    cvpdnSessionAttrEntry.EntityData.Leafs.Append("cvpdnSessionAttrCallDuration", types.YLeaf{"CvpdnSessionAttrCallDuration", cvpdnSessionAttrEntry.CvpdnSessionAttrCallDuration})
    cvpdnSessionAttrEntry.EntityData.Leafs.Append("cvpdnSessionAttrPacketsOut", types.YLeaf{"CvpdnSessionAttrPacketsOut", cvpdnSessionAttrEntry.CvpdnSessionAttrPacketsOut})
    cvpdnSessionAttrEntry.EntityData.Leafs.Append("cvpdnSessionAttrBytesOut", types.YLeaf{"CvpdnSessionAttrBytesOut", cvpdnSessionAttrEntry.CvpdnSessionAttrBytesOut})
    cvpdnSessionAttrEntry.EntityData.Leafs.Append("cvpdnSessionAttrPacketsIn", types.YLeaf{"CvpdnSessionAttrPacketsIn", cvpdnSessionAttrEntry.CvpdnSessionAttrPacketsIn})
    cvpdnSessionAttrEntry.EntityData.Leafs.Append("cvpdnSessionAttrBytesIn", types.YLeaf{"CvpdnSessionAttrBytesIn", cvpdnSessionAttrEntry.CvpdnSessionAttrBytesIn})
    cvpdnSessionAttrEntry.EntityData.Leafs.Append("cvpdnSessionAttrDeviceType", types.YLeaf{"CvpdnSessionAttrDeviceType", cvpdnSessionAttrEntry.CvpdnSessionAttrDeviceType})
    cvpdnSessionAttrEntry.EntityData.Leafs.Append("cvpdnSessionAttrDeviceCallerId", types.YLeaf{"CvpdnSessionAttrDeviceCallerId", cvpdnSessionAttrEntry.CvpdnSessionAttrDeviceCallerId})
    cvpdnSessionAttrEntry.EntityData.Leafs.Append("cvpdnSessionAttrDevicePhyId", types.YLeaf{"CvpdnSessionAttrDevicePhyId", cvpdnSessionAttrEntry.CvpdnSessionAttrDevicePhyId})
    cvpdnSessionAttrEntry.EntityData.Leafs.Append("cvpdnSessionAttrMultilink", types.YLeaf{"CvpdnSessionAttrMultilink", cvpdnSessionAttrEntry.CvpdnSessionAttrMultilink})
    cvpdnSessionAttrEntry.EntityData.Leafs.Append("cvpdnSessionAttrModemSlotIndex", types.YLeaf{"CvpdnSessionAttrModemSlotIndex", cvpdnSessionAttrEntry.CvpdnSessionAttrModemSlotIndex})
    cvpdnSessionAttrEntry.EntityData.Leafs.Append("cvpdnSessionAttrModemPortIndex", types.YLeaf{"CvpdnSessionAttrModemPortIndex", cvpdnSessionAttrEntry.CvpdnSessionAttrModemPortIndex})
    cvpdnSessionAttrEntry.EntityData.Leafs.Append("cvpdnSessionAttrDS1SlotIndex", types.YLeaf{"CvpdnSessionAttrDS1SlotIndex", cvpdnSessionAttrEntry.CvpdnSessionAttrDS1SlotIndex})
    cvpdnSessionAttrEntry.EntityData.Leafs.Append("cvpdnSessionAttrDS1PortIndex", types.YLeaf{"CvpdnSessionAttrDS1PortIndex", cvpdnSessionAttrEntry.CvpdnSessionAttrDS1PortIndex})
    cvpdnSessionAttrEntry.EntityData.Leafs.Append("cvpdnSessionAttrDS1ChannelIndex", types.YLeaf{"CvpdnSessionAttrDS1ChannelIndex", cvpdnSessionAttrEntry.CvpdnSessionAttrDS1ChannelIndex})
    cvpdnSessionAttrEntry.EntityData.Leafs.Append("cvpdnSessionAttrModemCallStartTime", types.YLeaf{"CvpdnSessionAttrModemCallStartTime", cvpdnSessionAttrEntry.CvpdnSessionAttrModemCallStartTime})
    cvpdnSessionAttrEntry.EntityData.Leafs.Append("cvpdnSessionAttrModemCallStartIndex", types.YLeaf{"CvpdnSessionAttrModemCallStartIndex", cvpdnSessionAttrEntry.CvpdnSessionAttrModemCallStartIndex})
    cvpdnSessionAttrEntry.EntityData.Leafs.Append("cvpdnSessionAttrVirtualCircuitID", types.YLeaf{"CvpdnSessionAttrVirtualCircuitID", cvpdnSessionAttrEntry.CvpdnSessionAttrVirtualCircuitID})
    cvpdnSessionAttrEntry.EntityData.Leafs.Append("cvpdnSessionAttrSentPktsDropped", types.YLeaf{"CvpdnSessionAttrSentPktsDropped", cvpdnSessionAttrEntry.CvpdnSessionAttrSentPktsDropped})
    cvpdnSessionAttrEntry.EntityData.Leafs.Append("cvpdnSessionAttrRecvPktsDropped", types.YLeaf{"CvpdnSessionAttrRecvPktsDropped", cvpdnSessionAttrEntry.CvpdnSessionAttrRecvPktsDropped})
    cvpdnSessionAttrEntry.EntityData.Leafs.Append("cvpdnSessionAttrMultilinkBundle", types.YLeaf{"CvpdnSessionAttrMultilinkBundle", cvpdnSessionAttrEntry.CvpdnSessionAttrMultilinkBundle})
    cvpdnSessionAttrEntry.EntityData.Leafs.Append("cvpdnSessionAttrMultilinkIfIndex", types.YLeaf{"CvpdnSessionAttrMultilinkIfIndex", cvpdnSessionAttrEntry.CvpdnSessionAttrMultilinkIfIndex})

    cvpdnSessionAttrEntry.EntityData.YListKeys = []string {"CvpdnSystemTunnelType", "CvpdnTunnelAttrTunnelId", "CvpdnSessionAttrSessionId"}

    return &(cvpdnSessionAttrEntry.EntityData)
}

// CISCOVPDNMGMTMIB_CvpdnSessionAttrTable_CvpdnSessionAttrEntry_CvpdnSessionAttrDeviceType represents     cable:              cable modem interface.
type CISCOVPDNMGMTMIB_CvpdnSessionAttrTable_CvpdnSessionAttrEntry_CvpdnSessionAttrDeviceType string

const (
    CISCOVPDNMGMTMIB_CvpdnSessionAttrTable_CvpdnSessionAttrEntry_CvpdnSessionAttrDeviceType_other CISCOVPDNMGMTMIB_CvpdnSessionAttrTable_CvpdnSessionAttrEntry_CvpdnSessionAttrDeviceType = "other"

    CISCOVPDNMGMTMIB_CvpdnSessionAttrTable_CvpdnSessionAttrEntry_CvpdnSessionAttrDeviceType_asyncInternalModem CISCOVPDNMGMTMIB_CvpdnSessionAttrTable_CvpdnSessionAttrEntry_CvpdnSessionAttrDeviceType = "asyncInternalModem"

    CISCOVPDNMGMTMIB_CvpdnSessionAttrTable_CvpdnSessionAttrEntry_CvpdnSessionAttrDeviceType_async CISCOVPDNMGMTMIB_CvpdnSessionAttrTable_CvpdnSessionAttrEntry_CvpdnSessionAttrDeviceType = "async"

    CISCOVPDNMGMTMIB_CvpdnSessionAttrTable_CvpdnSessionAttrEntry_CvpdnSessionAttrDeviceType_bchan CISCOVPDNMGMTMIB_CvpdnSessionAttrTable_CvpdnSessionAttrEntry_CvpdnSessionAttrDeviceType = "bchan"

    CISCOVPDNMGMTMIB_CvpdnSessionAttrTable_CvpdnSessionAttrEntry_CvpdnSessionAttrDeviceType_sync CISCOVPDNMGMTMIB_CvpdnSessionAttrTable_CvpdnSessionAttrEntry_CvpdnSessionAttrDeviceType = "sync"

    CISCOVPDNMGMTMIB_CvpdnSessionAttrTable_CvpdnSessionAttrEntry_CvpdnSessionAttrDeviceType_virtualAccess CISCOVPDNMGMTMIB_CvpdnSessionAttrTable_CvpdnSessionAttrEntry_CvpdnSessionAttrDeviceType = "virtualAccess"

    CISCOVPDNMGMTMIB_CvpdnSessionAttrTable_CvpdnSessionAttrEntry_CvpdnSessionAttrDeviceType_xdsl CISCOVPDNMGMTMIB_CvpdnSessionAttrTable_CvpdnSessionAttrEntry_CvpdnSessionAttrDeviceType = "xdsl"

    CISCOVPDNMGMTMIB_CvpdnSessionAttrTable_CvpdnSessionAttrEntry_CvpdnSessionAttrDeviceType_cable CISCOVPDNMGMTMIB_CvpdnSessionAttrTable_CvpdnSessionAttrEntry_CvpdnSessionAttrDeviceType = "cable"
)

// CISCOVPDNMGMTMIB_CvpdnSessionAttrTable_CvpdnSessionAttrEntry_CvpdnSessionAttrState represents     pptpWaitCallDisc:    Session shutdown is in progress.
type CISCOVPDNMGMTMIB_CvpdnSessionAttrTable_CvpdnSessionAttrEntry_CvpdnSessionAttrState string

const (
    CISCOVPDNMGMTMIB_CvpdnSessionAttrTable_CvpdnSessionAttrEntry_CvpdnSessionAttrState_unknown CISCOVPDNMGMTMIB_CvpdnSessionAttrTable_CvpdnSessionAttrEntry_CvpdnSessionAttrState = "unknown"

    CISCOVPDNMGMTMIB_CvpdnSessionAttrTable_CvpdnSessionAttrEntry_CvpdnSessionAttrState_l2fOpening CISCOVPDNMGMTMIB_CvpdnSessionAttrTable_CvpdnSessionAttrEntry_CvpdnSessionAttrState = "l2fOpening"

    CISCOVPDNMGMTMIB_CvpdnSessionAttrTable_CvpdnSessionAttrEntry_CvpdnSessionAttrState_l2fOpen CISCOVPDNMGMTMIB_CvpdnSessionAttrTable_CvpdnSessionAttrEntry_CvpdnSessionAttrState = "l2fOpen"

    CISCOVPDNMGMTMIB_CvpdnSessionAttrTable_CvpdnSessionAttrEntry_CvpdnSessionAttrState_l2fCloseWait CISCOVPDNMGMTMIB_CvpdnSessionAttrTable_CvpdnSessionAttrEntry_CvpdnSessionAttrState = "l2fCloseWait"

    CISCOVPDNMGMTMIB_CvpdnSessionAttrTable_CvpdnSessionAttrEntry_CvpdnSessionAttrState_l2fWaitingForTunnel CISCOVPDNMGMTMIB_CvpdnSessionAttrTable_CvpdnSessionAttrEntry_CvpdnSessionAttrState = "l2fWaitingForTunnel"

    CISCOVPDNMGMTMIB_CvpdnSessionAttrTable_CvpdnSessionAttrEntry_CvpdnSessionAttrState_l2tpIdle CISCOVPDNMGMTMIB_CvpdnSessionAttrTable_CvpdnSessionAttrEntry_CvpdnSessionAttrState = "l2tpIdle"

    CISCOVPDNMGMTMIB_CvpdnSessionAttrTable_CvpdnSessionAttrEntry_CvpdnSessionAttrState_l2tpWaitingTunnel CISCOVPDNMGMTMIB_CvpdnSessionAttrTable_CvpdnSessionAttrEntry_CvpdnSessionAttrState = "l2tpWaitingTunnel"

    CISCOVPDNMGMTMIB_CvpdnSessionAttrTable_CvpdnSessionAttrEntry_CvpdnSessionAttrState_l2tpWaitReply CISCOVPDNMGMTMIB_CvpdnSessionAttrTable_CvpdnSessionAttrEntry_CvpdnSessionAttrState = "l2tpWaitReply"

    CISCOVPDNMGMTMIB_CvpdnSessionAttrTable_CvpdnSessionAttrEntry_CvpdnSessionAttrState_l2tpWaitConnect CISCOVPDNMGMTMIB_CvpdnSessionAttrTable_CvpdnSessionAttrEntry_CvpdnSessionAttrState = "l2tpWaitConnect"

    CISCOVPDNMGMTMIB_CvpdnSessionAttrTable_CvpdnSessionAttrEntry_CvpdnSessionAttrState_l2tpEstablished CISCOVPDNMGMTMIB_CvpdnSessionAttrTable_CvpdnSessionAttrEntry_CvpdnSessionAttrState = "l2tpEstablished"

    CISCOVPDNMGMTMIB_CvpdnSessionAttrTable_CvpdnSessionAttrEntry_CvpdnSessionAttrState_l2tpShuttingDown CISCOVPDNMGMTMIB_CvpdnSessionAttrTable_CvpdnSessionAttrEntry_CvpdnSessionAttrState = "l2tpShuttingDown"

    CISCOVPDNMGMTMIB_CvpdnSessionAttrTable_CvpdnSessionAttrEntry_CvpdnSessionAttrState_pptpWaitVAccess CISCOVPDNMGMTMIB_CvpdnSessionAttrTable_CvpdnSessionAttrEntry_CvpdnSessionAttrState = "pptpWaitVAccess"

    CISCOVPDNMGMTMIB_CvpdnSessionAttrTable_CvpdnSessionAttrEntry_CvpdnSessionAttrState_pptpPacEstablished CISCOVPDNMGMTMIB_CvpdnSessionAttrTable_CvpdnSessionAttrEntry_CvpdnSessionAttrState = "pptpPacEstablished"

    CISCOVPDNMGMTMIB_CvpdnSessionAttrTable_CvpdnSessionAttrEntry_CvpdnSessionAttrState_pptpWaitTunnel CISCOVPDNMGMTMIB_CvpdnSessionAttrTable_CvpdnSessionAttrEntry_CvpdnSessionAttrState = "pptpWaitTunnel"

    CISCOVPDNMGMTMIB_CvpdnSessionAttrTable_CvpdnSessionAttrEntry_CvpdnSessionAttrState_pptpWaitOCRP CISCOVPDNMGMTMIB_CvpdnSessionAttrTable_CvpdnSessionAttrEntry_CvpdnSessionAttrState = "pptpWaitOCRP"

    CISCOVPDNMGMTMIB_CvpdnSessionAttrTable_CvpdnSessionAttrEntry_CvpdnSessionAttrState_pptpPnsEstablished CISCOVPDNMGMTMIB_CvpdnSessionAttrTable_CvpdnSessionAttrEntry_CvpdnSessionAttrState = "pptpPnsEstablished"

    CISCOVPDNMGMTMIB_CvpdnSessionAttrTable_CvpdnSessionAttrEntry_CvpdnSessionAttrState_pptpWaitCallDisc CISCOVPDNMGMTMIB_CvpdnSessionAttrTable_CvpdnSessionAttrEntry_CvpdnSessionAttrState = "pptpWaitCallDisc"

    CISCOVPDNMGMTMIB_CvpdnSessionAttrTable_CvpdnSessionAttrEntry_CvpdnSessionAttrState_pptpTerminal CISCOVPDNMGMTMIB_CvpdnSessionAttrTable_CvpdnSessionAttrEntry_CvpdnSessionAttrState = "pptpTerminal"
)

// CISCOVPDNMGMTMIB_CvpdnUserToFailHistInfoTable
// Table of the record of failure objects which can be
// referenced by an user name.  Only a name that has a
// valid item in the Cisco IOS VPDN failure history table
// will yield a valid entry in this table.  The table has
// a maximum size of 50 entries.  Only the newest 50
// entries will be kept in the table.
type CISCOVPDNMGMTMIB_CvpdnUserToFailHistInfoTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table, containing failure history relevant to an user name.
    // The type is slice of
    // CISCOVPDNMGMTMIB_CvpdnUserToFailHistInfoTable_CvpdnUserToFailHistInfoEntry.
    CvpdnUserToFailHistInfoEntry []*CISCOVPDNMGMTMIB_CvpdnUserToFailHistInfoTable_CvpdnUserToFailHistInfoEntry
}

func (cvpdnUserToFailHistInfoTable *CISCOVPDNMGMTMIB_CvpdnUserToFailHistInfoTable) GetEntityData() *types.CommonEntityData {
    cvpdnUserToFailHistInfoTable.EntityData.YFilter = cvpdnUserToFailHistInfoTable.YFilter
    cvpdnUserToFailHistInfoTable.EntityData.YangName = "cvpdnUserToFailHistInfoTable"
    cvpdnUserToFailHistInfoTable.EntityData.BundleName = "cisco_ios_xe"
    cvpdnUserToFailHistInfoTable.EntityData.ParentYangName = "CISCO-VPDN-MGMT-MIB"
    cvpdnUserToFailHistInfoTable.EntityData.SegmentPath = "cvpdnUserToFailHistInfoTable"
    cvpdnUserToFailHistInfoTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvpdnUserToFailHistInfoTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvpdnUserToFailHistInfoTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvpdnUserToFailHistInfoTable.EntityData.Children = types.NewOrderedMap()
    cvpdnUserToFailHistInfoTable.EntityData.Children.Append("cvpdnUserToFailHistInfoEntry", types.YChild{"CvpdnUserToFailHistInfoEntry", nil})
    for i := range cvpdnUserToFailHistInfoTable.CvpdnUserToFailHistInfoEntry {
        cvpdnUserToFailHistInfoTable.EntityData.Children.Append(types.GetSegmentPath(cvpdnUserToFailHistInfoTable.CvpdnUserToFailHistInfoEntry[i]), types.YChild{"CvpdnUserToFailHistInfoEntry", cvpdnUserToFailHistInfoTable.CvpdnUserToFailHistInfoEntry[i]})
    }
    cvpdnUserToFailHistInfoTable.EntityData.Leafs = types.NewOrderedMap()

    cvpdnUserToFailHistInfoTable.EntityData.YListKeys = []string {}

    return &(cvpdnUserToFailHistInfoTable.EntityData)
}

// CISCOVPDNMGMTMIB_CvpdnUserToFailHistInfoTable_CvpdnUserToFailHistInfoEntry
// An entry in the table, containing failure history
// relevant to an user name.
type CISCOVPDNMGMTMIB_CvpdnUserToFailHistInfoTable_CvpdnUserToFailHistInfoEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The user name for this failure entry. The type is
    // string with length: 1..255.
    CvpdnUnameToFailHistUname interface{}

    // This attribute is a key. The Tunnel ID for this failure entry.  If it is
    // the instigator of the tunnel, the ID is the TS tunnel ID, otherwise it is
    // the NAS tunnel ID. The type is interface{} with range: 0..4294967295.
    CvpdnUnameToFailHistTunnelId interface{}

    // The user ID of this failure entry. The type is interface{} with range:
    // 0..4294967295.
    CvpdnUnameToFailHistUserId interface{}

    // This object indicates whether the tunnel in which the failure occurred was
    // generated locally or not. The type is bool.
    CvpdnUnameToFailHistLocalInitConn interface{}

    // The local name of the VPDN tunnel in which the failure occurred.  It will
    // be the NAS name of the tunnel if the system serves as the NAS, or the TS
    // name of the tunnel if the system serves as the tunnel server.  The local
    // name is the configured host name of the system.  This object can be empty
    // if the failure occurred prior to successful tunnel projection, thus no
    // source name will be available. The type is string.
    CvpdnUnameToFailHistLocalName interface{}

    // The remote name of the VPDN tunnel in which the failure occurred.  It will
    // be the tunnel server name of the tunnel if the system is a NAS, or the NAS
    // name of the tunnel if the system serves as the tunnel server.  This object
    // can be empty if the failure occurred prior to successful tunnel projection,
    // thus no source name will be available. The type is string.
    CvpdnUnameToFailHistRemoteName interface{}

    // The source IP address of the tunnel in which the failure occurred.  This IP
    // address is that of the interface at the instigator end of the tunnel. The
    // type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    CvpdnUnameToFailHistSourceIp interface{}

    // The destination IP address of the tunnel in which the failure occurred. 
    // This IP address is that of the interface at the receiver end of the tunnel.
    // The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    CvpdnUnameToFailHistDestIp interface{}

    // This object is incremented when multiple failures has been experienced by
    // this user on this tunnel.  Seeing a delta of >1 is an indication that the
    // current failure record is the latest of a series of failures that has been
    // recorded. The type is interface{} with range: 0..4294967295. Units are
    // failures.
    CvpdnUnameToFailHistCount interface{}

    // This object specifies the time when the failure is occurred. The type is
    // interface{} with range: 0..4294967295.
    CvpdnUnameToFailHistFailTime interface{}

    // The type of failure for the current failure record.  It comes in a form of
    // a an ASCII string which describes the type of failure. The type is string
    // with length: 1..255.
    CvpdnUnameToFailHistFailType interface{}

    // The reason of failure for the current failure record. The type is string
    // with length: 1..255.
    CvpdnUnameToFailHistFailReason interface{}

    // Indicates the type of address contained in
    // cvpdnUnameToFailHistSourceInetAddr. The type is InetAddressType.
    CvpdnUnameToFailHistSourceInetType interface{}

    // The source IP address of the tunnel in which the failure occurred.  This IP
    // address is that of the interface at the instigator end of the tunnel.  The
    // instigator end is the peer which initiates the tunnel estalishment.  The
    // type of this address is determined by the value of
    // cvpdnUnameToFailHistSourceInetType. The type is string with length: 0..255.
    CvpdnUnameToFailHistSourceInetAddr interface{}

    // Indicates the type of address contained in
    // cvpdnUnameToFailHistDestInetAddr. The type is InetAddressType.
    CvpdnUnameToFailHistDestInetType interface{}

    // The destination IP address of the tunnel in which the failure occurred. 
    // This IP address is that of the interface at the receiver end of the tunnel.
    // The type  of this address is determined by the value of 
    // cvpdnUnameToFailHistDestInetType. The type is string with length: 0..255.
    CvpdnUnameToFailHistDestInetAddr interface{}
}

func (cvpdnUserToFailHistInfoEntry *CISCOVPDNMGMTMIB_CvpdnUserToFailHistInfoTable_CvpdnUserToFailHistInfoEntry) GetEntityData() *types.CommonEntityData {
    cvpdnUserToFailHistInfoEntry.EntityData.YFilter = cvpdnUserToFailHistInfoEntry.YFilter
    cvpdnUserToFailHistInfoEntry.EntityData.YangName = "cvpdnUserToFailHistInfoEntry"
    cvpdnUserToFailHistInfoEntry.EntityData.BundleName = "cisco_ios_xe"
    cvpdnUserToFailHistInfoEntry.EntityData.ParentYangName = "cvpdnUserToFailHistInfoTable"
    cvpdnUserToFailHistInfoEntry.EntityData.SegmentPath = "cvpdnUserToFailHistInfoEntry" + types.AddKeyToken(cvpdnUserToFailHistInfoEntry.CvpdnUnameToFailHistUname, "cvpdnUnameToFailHistUname") + types.AddKeyToken(cvpdnUserToFailHistInfoEntry.CvpdnUnameToFailHistTunnelId, "cvpdnUnameToFailHistTunnelId")
    cvpdnUserToFailHistInfoEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvpdnUserToFailHistInfoEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvpdnUserToFailHistInfoEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvpdnUserToFailHistInfoEntry.EntityData.Children = types.NewOrderedMap()
    cvpdnUserToFailHistInfoEntry.EntityData.Leafs = types.NewOrderedMap()
    cvpdnUserToFailHistInfoEntry.EntityData.Leafs.Append("cvpdnUnameToFailHistUname", types.YLeaf{"CvpdnUnameToFailHistUname", cvpdnUserToFailHistInfoEntry.CvpdnUnameToFailHistUname})
    cvpdnUserToFailHistInfoEntry.EntityData.Leafs.Append("cvpdnUnameToFailHistTunnelId", types.YLeaf{"CvpdnUnameToFailHistTunnelId", cvpdnUserToFailHistInfoEntry.CvpdnUnameToFailHistTunnelId})
    cvpdnUserToFailHistInfoEntry.EntityData.Leafs.Append("cvpdnUnameToFailHistUserId", types.YLeaf{"CvpdnUnameToFailHistUserId", cvpdnUserToFailHistInfoEntry.CvpdnUnameToFailHistUserId})
    cvpdnUserToFailHistInfoEntry.EntityData.Leafs.Append("cvpdnUnameToFailHistLocalInitConn", types.YLeaf{"CvpdnUnameToFailHistLocalInitConn", cvpdnUserToFailHistInfoEntry.CvpdnUnameToFailHistLocalInitConn})
    cvpdnUserToFailHistInfoEntry.EntityData.Leafs.Append("cvpdnUnameToFailHistLocalName", types.YLeaf{"CvpdnUnameToFailHistLocalName", cvpdnUserToFailHistInfoEntry.CvpdnUnameToFailHistLocalName})
    cvpdnUserToFailHistInfoEntry.EntityData.Leafs.Append("cvpdnUnameToFailHistRemoteName", types.YLeaf{"CvpdnUnameToFailHistRemoteName", cvpdnUserToFailHistInfoEntry.CvpdnUnameToFailHistRemoteName})
    cvpdnUserToFailHistInfoEntry.EntityData.Leafs.Append("cvpdnUnameToFailHistSourceIp", types.YLeaf{"CvpdnUnameToFailHistSourceIp", cvpdnUserToFailHistInfoEntry.CvpdnUnameToFailHistSourceIp})
    cvpdnUserToFailHistInfoEntry.EntityData.Leafs.Append("cvpdnUnameToFailHistDestIp", types.YLeaf{"CvpdnUnameToFailHistDestIp", cvpdnUserToFailHistInfoEntry.CvpdnUnameToFailHistDestIp})
    cvpdnUserToFailHistInfoEntry.EntityData.Leafs.Append("cvpdnUnameToFailHistCount", types.YLeaf{"CvpdnUnameToFailHistCount", cvpdnUserToFailHistInfoEntry.CvpdnUnameToFailHistCount})
    cvpdnUserToFailHistInfoEntry.EntityData.Leafs.Append("cvpdnUnameToFailHistFailTime", types.YLeaf{"CvpdnUnameToFailHistFailTime", cvpdnUserToFailHistInfoEntry.CvpdnUnameToFailHistFailTime})
    cvpdnUserToFailHistInfoEntry.EntityData.Leafs.Append("cvpdnUnameToFailHistFailType", types.YLeaf{"CvpdnUnameToFailHistFailType", cvpdnUserToFailHistInfoEntry.CvpdnUnameToFailHistFailType})
    cvpdnUserToFailHistInfoEntry.EntityData.Leafs.Append("cvpdnUnameToFailHistFailReason", types.YLeaf{"CvpdnUnameToFailHistFailReason", cvpdnUserToFailHistInfoEntry.CvpdnUnameToFailHistFailReason})
    cvpdnUserToFailHistInfoEntry.EntityData.Leafs.Append("cvpdnUnameToFailHistSourceInetType", types.YLeaf{"CvpdnUnameToFailHistSourceInetType", cvpdnUserToFailHistInfoEntry.CvpdnUnameToFailHistSourceInetType})
    cvpdnUserToFailHistInfoEntry.EntityData.Leafs.Append("cvpdnUnameToFailHistSourceInetAddr", types.YLeaf{"CvpdnUnameToFailHistSourceInetAddr", cvpdnUserToFailHistInfoEntry.CvpdnUnameToFailHistSourceInetAddr})
    cvpdnUserToFailHistInfoEntry.EntityData.Leafs.Append("cvpdnUnameToFailHistDestInetType", types.YLeaf{"CvpdnUnameToFailHistDestInetType", cvpdnUserToFailHistInfoEntry.CvpdnUnameToFailHistDestInetType})
    cvpdnUserToFailHistInfoEntry.EntityData.Leafs.Append("cvpdnUnameToFailHistDestInetAddr", types.YLeaf{"CvpdnUnameToFailHistDestInetAddr", cvpdnUserToFailHistInfoEntry.CvpdnUnameToFailHistDestInetAddr})

    cvpdnUserToFailHistInfoEntry.EntityData.YListKeys = []string {"CvpdnUnameToFailHistUname", "CvpdnUnameToFailHistTunnelId"}

    return &(cvpdnUserToFailHistInfoEntry.EntityData)
}

// CISCOVPDNMGMTMIB_CvpdnTemplateTable
// Table of information about the VPDN templates.  The
// VPDN template is a grouping mechanism that allows
// configuration settings to be shared among multiple VPDN
// groups.  One such setting is a limit on the number of
// active sessions across all VPDN groups associated with
// the template.  The template table allows customers to
// monitor template-wide information such as tracking the
// allocation of sessions across templates.
type CISCOVPDNMGMTMIB_CvpdnTemplateTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table, containing information about a single VPDN template.
    // The type is slice of
    // CISCOVPDNMGMTMIB_CvpdnTemplateTable_CvpdnTemplateEntry.
    CvpdnTemplateEntry []*CISCOVPDNMGMTMIB_CvpdnTemplateTable_CvpdnTemplateEntry
}

func (cvpdnTemplateTable *CISCOVPDNMGMTMIB_CvpdnTemplateTable) GetEntityData() *types.CommonEntityData {
    cvpdnTemplateTable.EntityData.YFilter = cvpdnTemplateTable.YFilter
    cvpdnTemplateTable.EntityData.YangName = "cvpdnTemplateTable"
    cvpdnTemplateTable.EntityData.BundleName = "cisco_ios_xe"
    cvpdnTemplateTable.EntityData.ParentYangName = "CISCO-VPDN-MGMT-MIB"
    cvpdnTemplateTable.EntityData.SegmentPath = "cvpdnTemplateTable"
    cvpdnTemplateTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvpdnTemplateTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvpdnTemplateTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvpdnTemplateTable.EntityData.Children = types.NewOrderedMap()
    cvpdnTemplateTable.EntityData.Children.Append("cvpdnTemplateEntry", types.YChild{"CvpdnTemplateEntry", nil})
    for i := range cvpdnTemplateTable.CvpdnTemplateEntry {
        cvpdnTemplateTable.EntityData.Children.Append(types.GetSegmentPath(cvpdnTemplateTable.CvpdnTemplateEntry[i]), types.YChild{"CvpdnTemplateEntry", cvpdnTemplateTable.CvpdnTemplateEntry[i]})
    }
    cvpdnTemplateTable.EntityData.Leafs = types.NewOrderedMap()

    cvpdnTemplateTable.EntityData.YListKeys = []string {}

    return &(cvpdnTemplateTable.EntityData)
}

// CISCOVPDNMGMTMIB_CvpdnTemplateTable_CvpdnTemplateEntry
// An entry in the table, containing information about a
// single VPDN template.
type CISCOVPDNMGMTMIB_CvpdnTemplateTable_CvpdnTemplateEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the VPDN template. The type is string
    // with length: 1..255.
    CvpdnTemplateName interface{}

    // The total number of active session in all groups associated with the
    // template. The type is interface{} with range: 0..4294967295. Units are
    // sessions.
    CvpdnTemplateActiveSessions interface{}
}

func (cvpdnTemplateEntry *CISCOVPDNMGMTMIB_CvpdnTemplateTable_CvpdnTemplateEntry) GetEntityData() *types.CommonEntityData {
    cvpdnTemplateEntry.EntityData.YFilter = cvpdnTemplateEntry.YFilter
    cvpdnTemplateEntry.EntityData.YangName = "cvpdnTemplateEntry"
    cvpdnTemplateEntry.EntityData.BundleName = "cisco_ios_xe"
    cvpdnTemplateEntry.EntityData.ParentYangName = "cvpdnTemplateTable"
    cvpdnTemplateEntry.EntityData.SegmentPath = "cvpdnTemplateEntry" + types.AddKeyToken(cvpdnTemplateEntry.CvpdnTemplateName, "cvpdnTemplateName")
    cvpdnTemplateEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvpdnTemplateEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvpdnTemplateEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvpdnTemplateEntry.EntityData.Children = types.NewOrderedMap()
    cvpdnTemplateEntry.EntityData.Leafs = types.NewOrderedMap()
    cvpdnTemplateEntry.EntityData.Leafs.Append("cvpdnTemplateName", types.YLeaf{"CvpdnTemplateName", cvpdnTemplateEntry.CvpdnTemplateName})
    cvpdnTemplateEntry.EntityData.Leafs.Append("cvpdnTemplateActiveSessions", types.YLeaf{"CvpdnTemplateActiveSessions", cvpdnTemplateEntry.CvpdnTemplateActiveSessions})

    cvpdnTemplateEntry.EntityData.YListKeys = []string {"CvpdnTemplateName"}

    return &(cvpdnTemplateEntry.EntityData)
}

// CISCOVPDNMGMTMIB_CvpdnBundleTable
// Table that describes the multilink PPP attributes of the
// active VPDN sessions.
type CISCOVPDNMGMTMIB_CvpdnBundleTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table represents an active multilink PPP bundle that
    // belongs to a VPDN tunnel. The type is slice of
    // CISCOVPDNMGMTMIB_CvpdnBundleTable_CvpdnBundleEntry.
    CvpdnBundleEntry []*CISCOVPDNMGMTMIB_CvpdnBundleTable_CvpdnBundleEntry
}

func (cvpdnBundleTable *CISCOVPDNMGMTMIB_CvpdnBundleTable) GetEntityData() *types.CommonEntityData {
    cvpdnBundleTable.EntityData.YFilter = cvpdnBundleTable.YFilter
    cvpdnBundleTable.EntityData.YangName = "cvpdnBundleTable"
    cvpdnBundleTable.EntityData.BundleName = "cisco_ios_xe"
    cvpdnBundleTable.EntityData.ParentYangName = "CISCO-VPDN-MGMT-MIB"
    cvpdnBundleTable.EntityData.SegmentPath = "cvpdnBundleTable"
    cvpdnBundleTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvpdnBundleTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvpdnBundleTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvpdnBundleTable.EntityData.Children = types.NewOrderedMap()
    cvpdnBundleTable.EntityData.Children.Append("cvpdnBundleEntry", types.YChild{"CvpdnBundleEntry", nil})
    for i := range cvpdnBundleTable.CvpdnBundleEntry {
        cvpdnBundleTable.EntityData.Children.Append(types.GetSegmentPath(cvpdnBundleTable.CvpdnBundleEntry[i]), types.YChild{"CvpdnBundleEntry", cvpdnBundleTable.CvpdnBundleEntry[i]})
    }
    cvpdnBundleTable.EntityData.Leafs = types.NewOrderedMap()

    cvpdnBundleTable.EntityData.YListKeys = []string {}

    return &(cvpdnBundleTable.EntityData)
}

// CISCOVPDNMGMTMIB_CvpdnBundleTable_CvpdnBundleEntry
// An entry in this table represents an active multilink PPP
// bundle that belongs to a VPDN tunnel.
type CISCOVPDNMGMTMIB_CvpdnBundleTable_CvpdnBundleEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the multilink PPP bundle associated
    // with a VPDN tunnel. The type is string with length: 1..64.
    CvpdnBundleName interface{}

    // The total number of active links in a multilink PPP bundle associated with
    // a VPDN tunnel. The type is interface{} with range: 0..4294967295. Units are
    // links.
    CvpdnBundleLinkCount interface{}

    // The multilink PPP bundle discriminator type associated with a VPDN tunnel. 
    // The value of this object represents the type of discriminator used in
    // cvpdnBundleEndpoint.      none:        No endpoint discriminator was
    // supplied, the                  default value is being used.      hostname: 
    // The router's hostname is being used as                  discriminator.     
    // string:      User specified string is being used as                 
    // discriminator.      macAddress:  A MAC address as defined by the MacAddress
    // textual convention is being used as                  discriminator.     
    // ipV4Address: An IP address as defined by the                 
    // InetAddressIPv4 textual convention is being                  used as
    // discriminator.      ipV6Address: An IP address as defined by the           
    // InetAddressIPv6 textual convention is being                  used as
    // discriminator.      phone:       The PSTN phone number is being used as    
    // discriminator.      magicNumber: A magic number is being used as           
    // discriminator. The type is CvpdnBundleEndpointType.
    CvpdnBundleEndpointType interface{}

    // Indicates the discriminator used in this bundle that is associated with a
    // VPDN tunnel. The type is string with length: 0..255.
    CvpdnBundleEndpoint interface{}

    // Indicates the type of address contained in cvpdnBundlePeerIpAddr. The type
    // is InetAddressType.
    CvpdnBundlePeerIpAddrType interface{}

    // The IP address of the multilink PPP peer in a VPDN tunnel. The type is
    // string with length: 0..255.
    CvpdnBundlePeerIpAddr interface{}

    // The multilink PPP bundle discriminator class associated with a VPDN tunnel.
    // The value of this object represents the type of discriminator used in
    // cvpdnBundleEndpoint. The type is EndpointClass.
    CvpdnBundleEndpointClass interface{}
}

func (cvpdnBundleEntry *CISCOVPDNMGMTMIB_CvpdnBundleTable_CvpdnBundleEntry) GetEntityData() *types.CommonEntityData {
    cvpdnBundleEntry.EntityData.YFilter = cvpdnBundleEntry.YFilter
    cvpdnBundleEntry.EntityData.YangName = "cvpdnBundleEntry"
    cvpdnBundleEntry.EntityData.BundleName = "cisco_ios_xe"
    cvpdnBundleEntry.EntityData.ParentYangName = "cvpdnBundleTable"
    cvpdnBundleEntry.EntityData.SegmentPath = "cvpdnBundleEntry" + types.AddKeyToken(cvpdnBundleEntry.CvpdnBundleName, "cvpdnBundleName")
    cvpdnBundleEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvpdnBundleEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvpdnBundleEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvpdnBundleEntry.EntityData.Children = types.NewOrderedMap()
    cvpdnBundleEntry.EntityData.Leafs = types.NewOrderedMap()
    cvpdnBundleEntry.EntityData.Leafs.Append("cvpdnBundleName", types.YLeaf{"CvpdnBundleName", cvpdnBundleEntry.CvpdnBundleName})
    cvpdnBundleEntry.EntityData.Leafs.Append("cvpdnBundleLinkCount", types.YLeaf{"CvpdnBundleLinkCount", cvpdnBundleEntry.CvpdnBundleLinkCount})
    cvpdnBundleEntry.EntityData.Leafs.Append("cvpdnBundleEndpointType", types.YLeaf{"CvpdnBundleEndpointType", cvpdnBundleEntry.CvpdnBundleEndpointType})
    cvpdnBundleEntry.EntityData.Leafs.Append("cvpdnBundleEndpoint", types.YLeaf{"CvpdnBundleEndpoint", cvpdnBundleEntry.CvpdnBundleEndpoint})
    cvpdnBundleEntry.EntityData.Leafs.Append("cvpdnBundlePeerIpAddrType", types.YLeaf{"CvpdnBundlePeerIpAddrType", cvpdnBundleEntry.CvpdnBundlePeerIpAddrType})
    cvpdnBundleEntry.EntityData.Leafs.Append("cvpdnBundlePeerIpAddr", types.YLeaf{"CvpdnBundlePeerIpAddr", cvpdnBundleEntry.CvpdnBundlePeerIpAddr})
    cvpdnBundleEntry.EntityData.Leafs.Append("cvpdnBundleEndpointClass", types.YLeaf{"CvpdnBundleEndpointClass", cvpdnBundleEntry.CvpdnBundleEndpointClass})

    cvpdnBundleEntry.EntityData.YListKeys = []string {"CvpdnBundleName"}

    return &(cvpdnBundleEntry.EntityData)
}

// CISCOVPDNMGMTMIB_CvpdnBundleTable_CvpdnBundleEntry_CvpdnBundleEndpointType represents                  discriminator.
type CISCOVPDNMGMTMIB_CvpdnBundleTable_CvpdnBundleEntry_CvpdnBundleEndpointType string

const (
    CISCOVPDNMGMTMIB_CvpdnBundleTable_CvpdnBundleEntry_CvpdnBundleEndpointType_none CISCOVPDNMGMTMIB_CvpdnBundleTable_CvpdnBundleEntry_CvpdnBundleEndpointType = "none"

    CISCOVPDNMGMTMIB_CvpdnBundleTable_CvpdnBundleEntry_CvpdnBundleEndpointType_hostname CISCOVPDNMGMTMIB_CvpdnBundleTable_CvpdnBundleEntry_CvpdnBundleEndpointType = "hostname"

    CISCOVPDNMGMTMIB_CvpdnBundleTable_CvpdnBundleEntry_CvpdnBundleEndpointType_string_ CISCOVPDNMGMTMIB_CvpdnBundleTable_CvpdnBundleEntry_CvpdnBundleEndpointType = "string"

    CISCOVPDNMGMTMIB_CvpdnBundleTable_CvpdnBundleEntry_CvpdnBundleEndpointType_macAddress CISCOVPDNMGMTMIB_CvpdnBundleTable_CvpdnBundleEntry_CvpdnBundleEndpointType = "macAddress"

    CISCOVPDNMGMTMIB_CvpdnBundleTable_CvpdnBundleEntry_CvpdnBundleEndpointType_ipV4Address CISCOVPDNMGMTMIB_CvpdnBundleTable_CvpdnBundleEntry_CvpdnBundleEndpointType = "ipV4Address"

    CISCOVPDNMGMTMIB_CvpdnBundleTable_CvpdnBundleEntry_CvpdnBundleEndpointType_ipV6Address CISCOVPDNMGMTMIB_CvpdnBundleTable_CvpdnBundleEntry_CvpdnBundleEndpointType = "ipV6Address"

    CISCOVPDNMGMTMIB_CvpdnBundleTable_CvpdnBundleEntry_CvpdnBundleEndpointType_phone CISCOVPDNMGMTMIB_CvpdnBundleTable_CvpdnBundleEntry_CvpdnBundleEndpointType = "phone"

    CISCOVPDNMGMTMIB_CvpdnBundleTable_CvpdnBundleEntry_CvpdnBundleEndpointType_magicNumber CISCOVPDNMGMTMIB_CvpdnBundleTable_CvpdnBundleEntry_CvpdnBundleEndpointType = "magicNumber"
)

// CISCOVPDNMGMTMIB_CvpdnBundleChildTable
// A table that exposes the containment relationship between a
// multilink PPP bundle and a VPDN tunnel.
type CISCOVPDNMGMTMIB_CvpdnBundleChildTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table represents a session that belongs to a VPDN tunnel
    // and to a multilink PPP bundle. The type is slice of
    // CISCOVPDNMGMTMIB_CvpdnBundleChildTable_CvpdnBundleChildEntry.
    CvpdnBundleChildEntry []*CISCOVPDNMGMTMIB_CvpdnBundleChildTable_CvpdnBundleChildEntry
}

func (cvpdnBundleChildTable *CISCOVPDNMGMTMIB_CvpdnBundleChildTable) GetEntityData() *types.CommonEntityData {
    cvpdnBundleChildTable.EntityData.YFilter = cvpdnBundleChildTable.YFilter
    cvpdnBundleChildTable.EntityData.YangName = "cvpdnBundleChildTable"
    cvpdnBundleChildTable.EntityData.BundleName = "cisco_ios_xe"
    cvpdnBundleChildTable.EntityData.ParentYangName = "CISCO-VPDN-MGMT-MIB"
    cvpdnBundleChildTable.EntityData.SegmentPath = "cvpdnBundleChildTable"
    cvpdnBundleChildTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvpdnBundleChildTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvpdnBundleChildTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvpdnBundleChildTable.EntityData.Children = types.NewOrderedMap()
    cvpdnBundleChildTable.EntityData.Children.Append("cvpdnBundleChildEntry", types.YChild{"CvpdnBundleChildEntry", nil})
    for i := range cvpdnBundleChildTable.CvpdnBundleChildEntry {
        cvpdnBundleChildTable.EntityData.Children.Append(types.GetSegmentPath(cvpdnBundleChildTable.CvpdnBundleChildEntry[i]), types.YChild{"CvpdnBundleChildEntry", cvpdnBundleChildTable.CvpdnBundleChildEntry[i]})
    }
    cvpdnBundleChildTable.EntityData.Leafs = types.NewOrderedMap()

    cvpdnBundleChildTable.EntityData.YListKeys = []string {}

    return &(cvpdnBundleChildTable.EntityData)
}

// CISCOVPDNMGMTMIB_CvpdnBundleChildTable_CvpdnBundleChildEntry
// An entry in this table represents a session that belongs to
// a VPDN tunnel and to a multilink PPP bundle.
type CISCOVPDNMGMTMIB_CvpdnBundleChildTable_CvpdnBundleChildEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with length: 1..64. Refers to
    // cisco_vpdn_mgmt_mib.CISCOVPDNMGMTMIB_CvpdnBundleTable_CvpdnBundleEntry_CvpdnBundleName
    CvpdnBundleName interface{}

    // This attribute is a key. The tunnel type.  This is the tunnel protocol of
    // an active VPDN session that is associated with a multilink PPP bundle. The
    // type is TunnelType.
    CvpdnBundleChildTunnelType interface{}

    // This attribute is a key. The Tunnel ID of an active VPDN session that is
    // associated with a multilink PPP bundle. The type is interface{} with range:
    // 0..4294967295.
    CvpdnBundleChildTunnelId interface{}

    // This attribute is a key. The ID of an active VPDN session that is
    // associated with a multilink PPP bundle. The type is interface{} with range:
    // 0..4294967295.
    CvpdnBundleChildSessionId interface{}
}

func (cvpdnBundleChildEntry *CISCOVPDNMGMTMIB_CvpdnBundleChildTable_CvpdnBundleChildEntry) GetEntityData() *types.CommonEntityData {
    cvpdnBundleChildEntry.EntityData.YFilter = cvpdnBundleChildEntry.YFilter
    cvpdnBundleChildEntry.EntityData.YangName = "cvpdnBundleChildEntry"
    cvpdnBundleChildEntry.EntityData.BundleName = "cisco_ios_xe"
    cvpdnBundleChildEntry.EntityData.ParentYangName = "cvpdnBundleChildTable"
    cvpdnBundleChildEntry.EntityData.SegmentPath = "cvpdnBundleChildEntry" + types.AddKeyToken(cvpdnBundleChildEntry.CvpdnBundleName, "cvpdnBundleName") + types.AddKeyToken(cvpdnBundleChildEntry.CvpdnBundleChildTunnelType, "cvpdnBundleChildTunnelType") + types.AddKeyToken(cvpdnBundleChildEntry.CvpdnBundleChildTunnelId, "cvpdnBundleChildTunnelId") + types.AddKeyToken(cvpdnBundleChildEntry.CvpdnBundleChildSessionId, "cvpdnBundleChildSessionId")
    cvpdnBundleChildEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvpdnBundleChildEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvpdnBundleChildEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvpdnBundleChildEntry.EntityData.Children = types.NewOrderedMap()
    cvpdnBundleChildEntry.EntityData.Leafs = types.NewOrderedMap()
    cvpdnBundleChildEntry.EntityData.Leafs.Append("cvpdnBundleName", types.YLeaf{"CvpdnBundleName", cvpdnBundleChildEntry.CvpdnBundleName})
    cvpdnBundleChildEntry.EntityData.Leafs.Append("cvpdnBundleChildTunnelType", types.YLeaf{"CvpdnBundleChildTunnelType", cvpdnBundleChildEntry.CvpdnBundleChildTunnelType})
    cvpdnBundleChildEntry.EntityData.Leafs.Append("cvpdnBundleChildTunnelId", types.YLeaf{"CvpdnBundleChildTunnelId", cvpdnBundleChildEntry.CvpdnBundleChildTunnelId})
    cvpdnBundleChildEntry.EntityData.Leafs.Append("cvpdnBundleChildSessionId", types.YLeaf{"CvpdnBundleChildSessionId", cvpdnBundleChildEntry.CvpdnBundleChildSessionId})

    cvpdnBundleChildEntry.EntityData.YListKeys = []string {"CvpdnBundleName", "CvpdnBundleChildTunnelType", "CvpdnBundleChildTunnelId", "CvpdnBundleChildSessionId"}

    return &(cvpdnBundleChildEntry.EntityData)
}

