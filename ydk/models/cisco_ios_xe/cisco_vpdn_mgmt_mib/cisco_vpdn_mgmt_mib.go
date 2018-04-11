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

// TunnelType represents tunnel.
type TunnelType string

const (
    TunnelType_l2f TunnelType = "l2f"

    TunnelType_l2tp TunnelType = "l2tp"

    TunnelType_pptp TunnelType = "pptp"
)

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

// CISCOVPDNMGMTMIB
type CISCOVPDNMGMTMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Ciscovpdnmgmtmibnotifs CISCOVPDNMGMTMIB_Ciscovpdnmgmtmibnotifs

    
    Cvpdnsysteminfo CISCOVPDNMGMTMIB_Cvpdnsysteminfo

    
    Cvpdnmultilinkinfo CISCOVPDNMGMTMIB_Cvpdnmultilinkinfo

    // Table of information about the VPDN system for all tunnel types.
    Cvpdnsystemtable CISCOVPDNMGMTMIB_Cvpdnsystemtable

    // Table of information about the active VPDN tunnels.
    Cvpdntunneltable CISCOVPDNMGMTMIB_Cvpdntunneltable

    // Table of information about the active VPDN tunnels.  An entry is added to
    // the table when a new tunnel is initiated and removed from the table when
    // the tunnel is terminated.
    Cvpdntunnelattrtable CISCOVPDNMGMTMIB_Cvpdntunnelattrtable

    // Table of information about individual user sessions within the active
    // tunnels.  Entry is added to the table when new user session is initiated
    // and be removed from the table when the user session is terminated.
    Cvpdntunnelsessiontable CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable

    // Table of information about individual sessions within the active tunnels. 
    // An entry is added to the table when a new session is initiated and removed
    // from the table when the session is terminated.
    Cvpdnsessionattrtable CISCOVPDNMGMTMIB_Cvpdnsessionattrtable

    // Table of the record of failure objects which can be referenced by an user
    // name.  Only a name that has a valid item in the Cisco IOS VPDN failure
    // history table will yield a valid entry in this table.  The table has a
    // maximum size of 50 entries.  Only the newest 50 entries will be kept in the
    // table.
    Cvpdnusertofailhistinfotable CISCOVPDNMGMTMIB_Cvpdnusertofailhistinfotable

    // Table of information about the VPDN templates.  The VPDN template is a
    // grouping mechanism that allows configuration settings to be shared among
    // multiple VPDN groups.  One such setting is a limit on the number of active
    // sessions across all VPDN groups associated with the template.  The template
    // table allows customers to monitor template-wide information such as
    // tracking the allocation of sessions across templates.
    Cvpdntemplatetable CISCOVPDNMGMTMIB_Cvpdntemplatetable

    // Table that describes the multilink PPP attributes of the active VPDN
    // sessions.
    Cvpdnbundletable CISCOVPDNMGMTMIB_Cvpdnbundletable

    // A table that exposes the containment relationship between a multilink PPP
    // bundle and a VPDN tunnel.
    Cvpdnbundlechildtable CISCOVPDNMGMTMIB_Cvpdnbundlechildtable
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

    cISCOVPDNMGMTMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOVPDNMGMTMIB.EntityData.Children["ciscoVpdnMgmtMIBNotifs"] = types.YChild{"Ciscovpdnmgmtmibnotifs", &cISCOVPDNMGMTMIB.Ciscovpdnmgmtmibnotifs}
    cISCOVPDNMGMTMIB.EntityData.Children["cvpdnSystemInfo"] = types.YChild{"Cvpdnsysteminfo", &cISCOVPDNMGMTMIB.Cvpdnsysteminfo}
    cISCOVPDNMGMTMIB.EntityData.Children["cvpdnMultilinkInfo"] = types.YChild{"Cvpdnmultilinkinfo", &cISCOVPDNMGMTMIB.Cvpdnmultilinkinfo}
    cISCOVPDNMGMTMIB.EntityData.Children["cvpdnSystemTable"] = types.YChild{"Cvpdnsystemtable", &cISCOVPDNMGMTMIB.Cvpdnsystemtable}
    cISCOVPDNMGMTMIB.EntityData.Children["cvpdnTunnelTable"] = types.YChild{"Cvpdntunneltable", &cISCOVPDNMGMTMIB.Cvpdntunneltable}
    cISCOVPDNMGMTMIB.EntityData.Children["cvpdnTunnelAttrTable"] = types.YChild{"Cvpdntunnelattrtable", &cISCOVPDNMGMTMIB.Cvpdntunnelattrtable}
    cISCOVPDNMGMTMIB.EntityData.Children["cvpdnTunnelSessionTable"] = types.YChild{"Cvpdntunnelsessiontable", &cISCOVPDNMGMTMIB.Cvpdntunnelsessiontable}
    cISCOVPDNMGMTMIB.EntityData.Children["cvpdnSessionAttrTable"] = types.YChild{"Cvpdnsessionattrtable", &cISCOVPDNMGMTMIB.Cvpdnsessionattrtable}
    cISCOVPDNMGMTMIB.EntityData.Children["cvpdnUserToFailHistInfoTable"] = types.YChild{"Cvpdnusertofailhistinfotable", &cISCOVPDNMGMTMIB.Cvpdnusertofailhistinfotable}
    cISCOVPDNMGMTMIB.EntityData.Children["cvpdnTemplateTable"] = types.YChild{"Cvpdntemplatetable", &cISCOVPDNMGMTMIB.Cvpdntemplatetable}
    cISCOVPDNMGMTMIB.EntityData.Children["cvpdnBundleTable"] = types.YChild{"Cvpdnbundletable", &cISCOVPDNMGMTMIB.Cvpdnbundletable}
    cISCOVPDNMGMTMIB.EntityData.Children["cvpdnBundleChildTable"] = types.YChild{"Cvpdnbundlechildtable", &cISCOVPDNMGMTMIB.Cvpdnbundlechildtable}
    cISCOVPDNMGMTMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOVPDNMGMTMIB.EntityData)
}

// CISCOVPDNMGMTMIB_Ciscovpdnmgmtmibnotifs
type CISCOVPDNMGMTMIB_Ciscovpdnmgmtmibnotifs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object contains the local session ID of the L2X session for which this
    // notification has been generated. The type is interface{} with range:
    // 0..65535.
    Cvpdnnotifsessionid interface{}

    // Indicates the event that generated the L2X session notification.  The
    // events are represented as follows:  up:     Session has come up.  down:  
    // Session has gone down.  pwUp:   Pseudowire associated with this         
    // session has come up.  pwDown: Pseudowire associated with this         
    // session has gone down. The type is Cvpdnnotifsessionevent.
    Cvpdnnotifsessionevent interface{}
}

func (ciscovpdnmgmtmibnotifs *CISCOVPDNMGMTMIB_Ciscovpdnmgmtmibnotifs) GetEntityData() *types.CommonEntityData {
    ciscovpdnmgmtmibnotifs.EntityData.YFilter = ciscovpdnmgmtmibnotifs.YFilter
    ciscovpdnmgmtmibnotifs.EntityData.YangName = "ciscoVpdnMgmtMIBNotifs"
    ciscovpdnmgmtmibnotifs.EntityData.BundleName = "cisco_ios_xe"
    ciscovpdnmgmtmibnotifs.EntityData.ParentYangName = "CISCO-VPDN-MGMT-MIB"
    ciscovpdnmgmtmibnotifs.EntityData.SegmentPath = "ciscoVpdnMgmtMIBNotifs"
    ciscovpdnmgmtmibnotifs.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscovpdnmgmtmibnotifs.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscovpdnmgmtmibnotifs.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscovpdnmgmtmibnotifs.EntityData.Children = make(map[string]types.YChild)
    ciscovpdnmgmtmibnotifs.EntityData.Leafs = make(map[string]types.YLeaf)
    ciscovpdnmgmtmibnotifs.EntityData.Leafs["cvpdnNotifSessionID"] = types.YLeaf{"Cvpdnnotifsessionid", ciscovpdnmgmtmibnotifs.Cvpdnnotifsessionid}
    ciscovpdnmgmtmibnotifs.EntityData.Leafs["cvpdnNotifSessionEvent"] = types.YLeaf{"Cvpdnnotifsessionevent", ciscovpdnmgmtmibnotifs.Cvpdnnotifsessionevent}
    return &(ciscovpdnmgmtmibnotifs.EntityData)
}

// CISCOVPDNMGMTMIB_Ciscovpdnmgmtmibnotifs_Cvpdnnotifsessionevent represents         session has gone down.
type CISCOVPDNMGMTMIB_Ciscovpdnmgmtmibnotifs_Cvpdnnotifsessionevent string

const (
    CISCOVPDNMGMTMIB_Ciscovpdnmgmtmibnotifs_Cvpdnnotifsessionevent_up CISCOVPDNMGMTMIB_Ciscovpdnmgmtmibnotifs_Cvpdnnotifsessionevent = "up"

    CISCOVPDNMGMTMIB_Ciscovpdnmgmtmibnotifs_Cvpdnnotifsessionevent_down CISCOVPDNMGMTMIB_Ciscovpdnmgmtmibnotifs_Cvpdnnotifsessionevent = "down"

    CISCOVPDNMGMTMIB_Ciscovpdnmgmtmibnotifs_Cvpdnnotifsessionevent_pwUp CISCOVPDNMGMTMIB_Ciscovpdnmgmtmibnotifs_Cvpdnnotifsessionevent = "pwUp"

    CISCOVPDNMGMTMIB_Ciscovpdnmgmtmibnotifs_Cvpdnnotifsessionevent_pwDown CISCOVPDNMGMTMIB_Ciscovpdnmgmtmibnotifs_Cvpdnnotifsessionevent = "pwDown"
)

// CISCOVPDNMGMTMIB_Cvpdnsysteminfo
type CISCOVPDNMGMTMIB_Cvpdnsysteminfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The total number of VPDN tunnels that are currently active within this
    // system. The type is interface{} with range: 0..4294967295. Units are
    // tunnels.
    Cvpdntunneltotal interface{}

    // The total number of active users in all the active VPDN tunnels within this
    // system. The type is interface{} with range: 0..4294967295. Units are users.
    Cvpdnsessiontotal interface{}

    // The total number of denied user attempts to all the active VPDN tunnels
    // within this system. The type is interface{} with range: 0..4294967295.
    // Units are attempts.
    Cvpdndenieduserstotal interface{}

    // Indicates whether Layer 2 VPN session notifications are enabled. The type
    // is bool.
    Cvpdnsystemnotifsessionenabled interface{}

    // Clears all the sessions in a given tunnel type.  When reading this object,
    // the value of 'none' will always be returned.  When setting these values,
    // the following operations will be performed:      none: no operation.     
    // all:  clears all the sessions in all the tunnels.      l2f:  clears all the
    // L2F sessions.      l2tp: clears all the L2TP sessions.      pptp: clears
    // all the PPTP sessions. The type is Cvpdnsystemclearsessions.
    Cvpdnsystemclearsessions interface{}
}

func (cvpdnsysteminfo *CISCOVPDNMGMTMIB_Cvpdnsysteminfo) GetEntityData() *types.CommonEntityData {
    cvpdnsysteminfo.EntityData.YFilter = cvpdnsysteminfo.YFilter
    cvpdnsysteminfo.EntityData.YangName = "cvpdnSystemInfo"
    cvpdnsysteminfo.EntityData.BundleName = "cisco_ios_xe"
    cvpdnsysteminfo.EntityData.ParentYangName = "CISCO-VPDN-MGMT-MIB"
    cvpdnsysteminfo.EntityData.SegmentPath = "cvpdnSystemInfo"
    cvpdnsysteminfo.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvpdnsysteminfo.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvpdnsysteminfo.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvpdnsysteminfo.EntityData.Children = make(map[string]types.YChild)
    cvpdnsysteminfo.EntityData.Leafs = make(map[string]types.YLeaf)
    cvpdnsysteminfo.EntityData.Leafs["cvpdnTunnelTotal"] = types.YLeaf{"Cvpdntunneltotal", cvpdnsysteminfo.Cvpdntunneltotal}
    cvpdnsysteminfo.EntityData.Leafs["cvpdnSessionTotal"] = types.YLeaf{"Cvpdnsessiontotal", cvpdnsysteminfo.Cvpdnsessiontotal}
    cvpdnsysteminfo.EntityData.Leafs["cvpdnDeniedUsersTotal"] = types.YLeaf{"Cvpdndenieduserstotal", cvpdnsysteminfo.Cvpdndenieduserstotal}
    cvpdnsysteminfo.EntityData.Leafs["cvpdnSystemNotifSessionEnabled"] = types.YLeaf{"Cvpdnsystemnotifsessionenabled", cvpdnsysteminfo.Cvpdnsystemnotifsessionenabled}
    cvpdnsysteminfo.EntityData.Leafs["cvpdnSystemClearSessions"] = types.YLeaf{"Cvpdnsystemclearsessions", cvpdnsysteminfo.Cvpdnsystemclearsessions}
    return &(cvpdnsysteminfo.EntityData)
}

// CISCOVPDNMGMTMIB_Cvpdnsysteminfo_Cvpdnsystemclearsessions represents     pptp: clears all the PPTP sessions.
type CISCOVPDNMGMTMIB_Cvpdnsysteminfo_Cvpdnsystemclearsessions string

const (
    CISCOVPDNMGMTMIB_Cvpdnsysteminfo_Cvpdnsystemclearsessions_none CISCOVPDNMGMTMIB_Cvpdnsysteminfo_Cvpdnsystemclearsessions = "none"

    CISCOVPDNMGMTMIB_Cvpdnsysteminfo_Cvpdnsystemclearsessions_all CISCOVPDNMGMTMIB_Cvpdnsysteminfo_Cvpdnsystemclearsessions = "all"

    CISCOVPDNMGMTMIB_Cvpdnsysteminfo_Cvpdnsystemclearsessions_l2f CISCOVPDNMGMTMIB_Cvpdnsysteminfo_Cvpdnsystemclearsessions = "l2f"

    CISCOVPDNMGMTMIB_Cvpdnsysteminfo_Cvpdnsystemclearsessions_l2tp CISCOVPDNMGMTMIB_Cvpdnsysteminfo_Cvpdnsystemclearsessions = "l2tp"

    CISCOVPDNMGMTMIB_Cvpdnsysteminfo_Cvpdnsystemclearsessions_pptp CISCOVPDNMGMTMIB_Cvpdnsysteminfo_Cvpdnsystemclearsessions = "pptp"
)

// CISCOVPDNMGMTMIB_Cvpdnmultilinkinfo
type CISCOVPDNMGMTMIB_Cvpdnmultilinkinfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The total number of bundles comprised of a single link. The type is
    // interface{} with range: 0..4294967295.
    Cvpdnbundleswithonelink interface{}

    // The total number of bundles comprised of two links. The type is interface{}
    // with range: 0..4294967295.
    Cvpdnbundleswithtwolinks interface{}

    // The total number of bundles comprised of more than two links. The type is
    // interface{} with range: 0..4294967295.
    Cvpdnbundleswithmorethantwolinks interface{}

    // The value of the sysUpTime object when the contents of cvpdnBundleTable
    // last changed. The type is interface{} with range: 0..4294967295.
    Cvpdnbundlelastchanged interface{}
}

func (cvpdnmultilinkinfo *CISCOVPDNMGMTMIB_Cvpdnmultilinkinfo) GetEntityData() *types.CommonEntityData {
    cvpdnmultilinkinfo.EntityData.YFilter = cvpdnmultilinkinfo.YFilter
    cvpdnmultilinkinfo.EntityData.YangName = "cvpdnMultilinkInfo"
    cvpdnmultilinkinfo.EntityData.BundleName = "cisco_ios_xe"
    cvpdnmultilinkinfo.EntityData.ParentYangName = "CISCO-VPDN-MGMT-MIB"
    cvpdnmultilinkinfo.EntityData.SegmentPath = "cvpdnMultilinkInfo"
    cvpdnmultilinkinfo.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvpdnmultilinkinfo.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvpdnmultilinkinfo.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvpdnmultilinkinfo.EntityData.Children = make(map[string]types.YChild)
    cvpdnmultilinkinfo.EntityData.Leafs = make(map[string]types.YLeaf)
    cvpdnmultilinkinfo.EntityData.Leafs["cvpdnBundlesWithOneLink"] = types.YLeaf{"Cvpdnbundleswithonelink", cvpdnmultilinkinfo.Cvpdnbundleswithonelink}
    cvpdnmultilinkinfo.EntityData.Leafs["cvpdnBundlesWithTwoLinks"] = types.YLeaf{"Cvpdnbundleswithtwolinks", cvpdnmultilinkinfo.Cvpdnbundleswithtwolinks}
    cvpdnmultilinkinfo.EntityData.Leafs["cvpdnBundlesWithMoreThanTwoLinks"] = types.YLeaf{"Cvpdnbundleswithmorethantwolinks", cvpdnmultilinkinfo.Cvpdnbundleswithmorethantwolinks}
    cvpdnmultilinkinfo.EntityData.Leafs["cvpdnBundleLastChanged"] = types.YLeaf{"Cvpdnbundlelastchanged", cvpdnmultilinkinfo.Cvpdnbundlelastchanged}
    return &(cvpdnmultilinkinfo.EntityData)
}

// CISCOVPDNMGMTMIB_Cvpdnsystemtable
// Table of information about the VPDN system for all tunnel
// types.
type CISCOVPDNMGMTMIB_Cvpdnsystemtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table, containing information about a single type of VPDN
    // tunnel. The type is slice of
    // CISCOVPDNMGMTMIB_Cvpdnsystemtable_Cvpdnsystementry.
    Cvpdnsystementry []CISCOVPDNMGMTMIB_Cvpdnsystemtable_Cvpdnsystementry
}

func (cvpdnsystemtable *CISCOVPDNMGMTMIB_Cvpdnsystemtable) GetEntityData() *types.CommonEntityData {
    cvpdnsystemtable.EntityData.YFilter = cvpdnsystemtable.YFilter
    cvpdnsystemtable.EntityData.YangName = "cvpdnSystemTable"
    cvpdnsystemtable.EntityData.BundleName = "cisco_ios_xe"
    cvpdnsystemtable.EntityData.ParentYangName = "CISCO-VPDN-MGMT-MIB"
    cvpdnsystemtable.EntityData.SegmentPath = "cvpdnSystemTable"
    cvpdnsystemtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvpdnsystemtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvpdnsystemtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvpdnsystemtable.EntityData.Children = make(map[string]types.YChild)
    cvpdnsystemtable.EntityData.Children["cvpdnSystemEntry"] = types.YChild{"Cvpdnsystementry", nil}
    for i := range cvpdnsystemtable.Cvpdnsystementry {
        cvpdnsystemtable.EntityData.Children[types.GetSegmentPath(&cvpdnsystemtable.Cvpdnsystementry[i])] = types.YChild{"Cvpdnsystementry", &cvpdnsystemtable.Cvpdnsystementry[i]}
    }
    cvpdnsystemtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cvpdnsystemtable.EntityData)
}

// CISCOVPDNMGMTMIB_Cvpdnsystemtable_Cvpdnsystementry
// An entry in the table, containing information about a
// single type of VPDN tunnel.
type CISCOVPDNMGMTMIB_Cvpdnsystemtable_Cvpdnsystementry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The tunnel type.  This is the tunnel protocol. The
    // type is TunnelType.
    Cvpdnsystemtunneltype interface{}

    // The total number of VPDN tunnels that are currently active of this tunnel
    // type. The type is interface{} with range: 0..4294967295. Units are tunnels.
    Cvpdnsystemtunneltotal interface{}

    // The total number of active sessions in all the active VPDN tunnels of this
    // tunnel type. The type is interface{} with range: 0..4294967295. Units are
    // sessions.
    Cvpdnsystemsessiontotal interface{}

    // The total number of denied user attempts to all the VPDN tunnels of this
    // tunnel type since last system re-initialization. The type is interface{}
    // with range: 0..4294967295. Units are attempts.
    Cvpdnsystemdenieduserstotal interface{}

    // The total number tunnel connection attempts on all the VPDN tunnels of this
    // tunnel type since last system re-initialization. The type is interface{}
    // with range: 0..4294967295. Units are attempts.
    Cvpdnsysteminitialconnreq interface{}

    // The total number tunnel Successful connection attempts in VPDN tunnels of
    // this tunnel type since last system re-initialization. The type is
    // interface{} with range: 0..4294967295. Units are attempts.
    Cvpdnsystemsuccessconnreq interface{}

    // The total number tunnel Failed connection attempts in VPDN tunnels of this
    // tunnel type since last system re-initialization. The type is interface{}
    // with range: 0..4294967295. Units are attempts.
    Cvpdnsystemfailedconnreq interface{}
}

func (cvpdnsystementry *CISCOVPDNMGMTMIB_Cvpdnsystemtable_Cvpdnsystementry) GetEntityData() *types.CommonEntityData {
    cvpdnsystementry.EntityData.YFilter = cvpdnsystementry.YFilter
    cvpdnsystementry.EntityData.YangName = "cvpdnSystemEntry"
    cvpdnsystementry.EntityData.BundleName = "cisco_ios_xe"
    cvpdnsystementry.EntityData.ParentYangName = "cvpdnSystemTable"
    cvpdnsystementry.EntityData.SegmentPath = "cvpdnSystemEntry" + "[cvpdnSystemTunnelType='" + fmt.Sprintf("%v", cvpdnsystementry.Cvpdnsystemtunneltype) + "']"
    cvpdnsystementry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvpdnsystementry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvpdnsystementry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvpdnsystementry.EntityData.Children = make(map[string]types.YChild)
    cvpdnsystementry.EntityData.Leafs = make(map[string]types.YLeaf)
    cvpdnsystementry.EntityData.Leafs["cvpdnSystemTunnelType"] = types.YLeaf{"Cvpdnsystemtunneltype", cvpdnsystementry.Cvpdnsystemtunneltype}
    cvpdnsystementry.EntityData.Leafs["cvpdnSystemTunnelTotal"] = types.YLeaf{"Cvpdnsystemtunneltotal", cvpdnsystementry.Cvpdnsystemtunneltotal}
    cvpdnsystementry.EntityData.Leafs["cvpdnSystemSessionTotal"] = types.YLeaf{"Cvpdnsystemsessiontotal", cvpdnsystementry.Cvpdnsystemsessiontotal}
    cvpdnsystementry.EntityData.Leafs["cvpdnSystemDeniedUsersTotal"] = types.YLeaf{"Cvpdnsystemdenieduserstotal", cvpdnsystementry.Cvpdnsystemdenieduserstotal}
    cvpdnsystementry.EntityData.Leafs["cvpdnSystemInitialConnReq"] = types.YLeaf{"Cvpdnsysteminitialconnreq", cvpdnsystementry.Cvpdnsysteminitialconnreq}
    cvpdnsystementry.EntityData.Leafs["cvpdnSystemSuccessConnReq"] = types.YLeaf{"Cvpdnsystemsuccessconnreq", cvpdnsystementry.Cvpdnsystemsuccessconnreq}
    cvpdnsystementry.EntityData.Leafs["cvpdnSystemFailedConnReq"] = types.YLeaf{"Cvpdnsystemfailedconnreq", cvpdnsystementry.Cvpdnsystemfailedconnreq}
    return &(cvpdnsystementry.EntityData)
}

// CISCOVPDNMGMTMIB_Cvpdntunneltable
// Table of information about the active VPDN tunnels.
type CISCOVPDNMGMTMIB_Cvpdntunneltable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table, containing information about a single active VPDN
    // tunnel. The type is slice of
    // CISCOVPDNMGMTMIB_Cvpdntunneltable_Cvpdntunnelentry.
    Cvpdntunnelentry []CISCOVPDNMGMTMIB_Cvpdntunneltable_Cvpdntunnelentry
}

func (cvpdntunneltable *CISCOVPDNMGMTMIB_Cvpdntunneltable) GetEntityData() *types.CommonEntityData {
    cvpdntunneltable.EntityData.YFilter = cvpdntunneltable.YFilter
    cvpdntunneltable.EntityData.YangName = "cvpdnTunnelTable"
    cvpdntunneltable.EntityData.BundleName = "cisco_ios_xe"
    cvpdntunneltable.EntityData.ParentYangName = "CISCO-VPDN-MGMT-MIB"
    cvpdntunneltable.EntityData.SegmentPath = "cvpdnTunnelTable"
    cvpdntunneltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvpdntunneltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvpdntunneltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvpdntunneltable.EntityData.Children = make(map[string]types.YChild)
    cvpdntunneltable.EntityData.Children["cvpdnTunnelEntry"] = types.YChild{"Cvpdntunnelentry", nil}
    for i := range cvpdntunneltable.Cvpdntunnelentry {
        cvpdntunneltable.EntityData.Children[types.GetSegmentPath(&cvpdntunneltable.Cvpdntunnelentry[i])] = types.YChild{"Cvpdntunnelentry", &cvpdntunneltable.Cvpdntunnelentry[i]}
    }
    cvpdntunneltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cvpdntunneltable.EntityData)
}

// CISCOVPDNMGMTMIB_Cvpdntunneltable_Cvpdntunnelentry
// An entry in the table, containing information about a
// single active VPDN tunnel.
type CISCOVPDNMGMTMIB_Cvpdntunneltable_Cvpdntunnelentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The Tunnel ID of an active VPDN tunnel.  If it is
    // the instigator of the tunnel, the ID is the HGW/LNS tunnel ID, otherwise it
    // is the NAS/LAC tunnel ID. The type is interface{} with range:
    // 0..4294967295.
    Cvpdntunneltunnelid interface{}

    // The remote Tunnel ID of an active VPDN tunnel.  If it is the instigator of
    // the tunnel, the ID is the NAS/LAC tunnel ID, otherwise it is the HGW/LNS
    // tunnel ID. The type is interface{} with range: 0..4294967295.
    Cvpdntunnelremotetunnelid interface{}

    // The local name of an active VPDN tunnel.  It will be the NAS/LAC name of
    // the tunnel if the router serves as the NAS/LAC, or the HGW/LNS name of the
    // tunnel if the system serves as the home gateway.  Typically, the local name
    // is the configured host name of the router. The type is string with length:
    // 1..255.
    Cvpdntunnellocalname interface{}

    // The remote name of an active VPDN tunnel.  It will be the home gateway name
    // of the tunnel if the system is a NAS/LAC, or the NAS/LAC name of the tunnel
    // if the system serves as the home gateway. The type is string with length:
    // 1..255.
    Cvpdntunnelremotename interface{}

    // The remote end point name of an active VPDN tunnel. This name is either the
    // domain name or the DNIS that this tunnel is projected with. The type is
    // string with length: 1..255.
    Cvpdntunnelremoteendpointname interface{}

    // This object indicates whether the tunnel was generated locally or not. The
    // type is bool.
    Cvpdntunnellocalinitconnection interface{}

    // The cause which originated an active VPDN tunnel.  The tunnel can be
    // projected via domain name, DNIS or a stack group (SGBP). The type is
    // Cvpdntunnelorigcause.
    Cvpdntunnelorigcause interface{}

    // The current state of an active VPDN tunnel.  Each state code is explained
    // below:         unknown: The current state of the tunnel is                
    // unknown.         opening: The tunnel has just been instigated and          
    // is pending for a remote end reply to                 complete the process. 
    // open:    The tunnel is active.         closing: The tunnel has just been
    // shut down and                 is pending for the remote end to reply       
    // to complete the process. The type is Cvpdntunnelstate.
    Cvpdntunnelstate interface{}

    // The total number of active session currently in the tunnel. The type is
    // interface{} with range: 0..4294967295. Units are sessions.
    Cvpdntunnelactivesessions interface{}

    // A count of the accumulated total of denied users for the tunnel. The type
    // is interface{} with range: 0..4294967295. Units are calls.
    Cvpdntunneldeniedusers interface{}

    // A VPDN tunnel can be put into a soft shut state to prevent any new user
    // session to be added.  This object specifies whether this tunnel has been
    // soft shut. The type is bool.
    Cvpdntunnelsoftshut interface{}

    // The type of network service used in the active tunnel. For now it is IP
    // only. The type is Cvpdntunnelnetworkservicetype.
    Cvpdntunnelnetworkservicetype interface{}

    // The local IP address of the tunnel.  This IP address is that of the
    // interface at the local end of the tunnel. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Cvpdntunnellocalipaddress interface{}

    // The source IP address of the tunnel.  This IP address is the user
    // configurable IP address for Stack Group Biding Protocol (SGBP) via the CLI
    // command: vpdn source-ip. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Cvpdntunnelsourceipaddress interface{}

    // The remote IP address of the tunnel.  This IP address is that of the
    // interface at the remote end of the tunnel. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Cvpdntunnelremoteipaddress interface{}
}

func (cvpdntunnelentry *CISCOVPDNMGMTMIB_Cvpdntunneltable_Cvpdntunnelentry) GetEntityData() *types.CommonEntityData {
    cvpdntunnelentry.EntityData.YFilter = cvpdntunnelentry.YFilter
    cvpdntunnelentry.EntityData.YangName = "cvpdnTunnelEntry"
    cvpdntunnelentry.EntityData.BundleName = "cisco_ios_xe"
    cvpdntunnelentry.EntityData.ParentYangName = "cvpdnTunnelTable"
    cvpdntunnelentry.EntityData.SegmentPath = "cvpdnTunnelEntry" + "[cvpdnTunnelTunnelId='" + fmt.Sprintf("%v", cvpdntunnelentry.Cvpdntunneltunnelid) + "']"
    cvpdntunnelentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvpdntunnelentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvpdntunnelentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvpdntunnelentry.EntityData.Children = make(map[string]types.YChild)
    cvpdntunnelentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cvpdntunnelentry.EntityData.Leafs["cvpdnTunnelTunnelId"] = types.YLeaf{"Cvpdntunneltunnelid", cvpdntunnelentry.Cvpdntunneltunnelid}
    cvpdntunnelentry.EntityData.Leafs["cvpdnTunnelRemoteTunnelId"] = types.YLeaf{"Cvpdntunnelremotetunnelid", cvpdntunnelentry.Cvpdntunnelremotetunnelid}
    cvpdntunnelentry.EntityData.Leafs["cvpdnTunnelLocalName"] = types.YLeaf{"Cvpdntunnellocalname", cvpdntunnelentry.Cvpdntunnellocalname}
    cvpdntunnelentry.EntityData.Leafs["cvpdnTunnelRemoteName"] = types.YLeaf{"Cvpdntunnelremotename", cvpdntunnelentry.Cvpdntunnelremotename}
    cvpdntunnelentry.EntityData.Leafs["cvpdnTunnelRemoteEndpointName"] = types.YLeaf{"Cvpdntunnelremoteendpointname", cvpdntunnelentry.Cvpdntunnelremoteendpointname}
    cvpdntunnelentry.EntityData.Leafs["cvpdnTunnelLocalInitConnection"] = types.YLeaf{"Cvpdntunnellocalinitconnection", cvpdntunnelentry.Cvpdntunnellocalinitconnection}
    cvpdntunnelentry.EntityData.Leafs["cvpdnTunnelOrigCause"] = types.YLeaf{"Cvpdntunnelorigcause", cvpdntunnelentry.Cvpdntunnelorigcause}
    cvpdntunnelentry.EntityData.Leafs["cvpdnTunnelState"] = types.YLeaf{"Cvpdntunnelstate", cvpdntunnelentry.Cvpdntunnelstate}
    cvpdntunnelentry.EntityData.Leafs["cvpdnTunnelActiveSessions"] = types.YLeaf{"Cvpdntunnelactivesessions", cvpdntunnelentry.Cvpdntunnelactivesessions}
    cvpdntunnelentry.EntityData.Leafs["cvpdnTunnelDeniedUsers"] = types.YLeaf{"Cvpdntunneldeniedusers", cvpdntunnelentry.Cvpdntunneldeniedusers}
    cvpdntunnelentry.EntityData.Leafs["cvpdnTunnelSoftshut"] = types.YLeaf{"Cvpdntunnelsoftshut", cvpdntunnelentry.Cvpdntunnelsoftshut}
    cvpdntunnelentry.EntityData.Leafs["cvpdnTunnelNetworkServiceType"] = types.YLeaf{"Cvpdntunnelnetworkservicetype", cvpdntunnelentry.Cvpdntunnelnetworkservicetype}
    cvpdntunnelentry.EntityData.Leafs["cvpdnTunnelLocalIpAddress"] = types.YLeaf{"Cvpdntunnellocalipaddress", cvpdntunnelentry.Cvpdntunnellocalipaddress}
    cvpdntunnelentry.EntityData.Leafs["cvpdnTunnelSourceIpAddress"] = types.YLeaf{"Cvpdntunnelsourceipaddress", cvpdntunnelentry.Cvpdntunnelsourceipaddress}
    cvpdntunnelentry.EntityData.Leafs["cvpdnTunnelRemoteIpAddress"] = types.YLeaf{"Cvpdntunnelremoteipaddress", cvpdntunnelentry.Cvpdntunnelremoteipaddress}
    return &(cvpdntunnelentry.EntityData)
}

// CISCOVPDNMGMTMIB_Cvpdntunneltable_Cvpdntunnelentry_Cvpdntunnelnetworkservicetype represents For now it is IP only.
type CISCOVPDNMGMTMIB_Cvpdntunneltable_Cvpdntunnelentry_Cvpdntunnelnetworkservicetype string

const (
    CISCOVPDNMGMTMIB_Cvpdntunneltable_Cvpdntunnelentry_Cvpdntunnelnetworkservicetype_ip CISCOVPDNMGMTMIB_Cvpdntunneltable_Cvpdntunnelentry_Cvpdntunnelnetworkservicetype = "ip"
)

// CISCOVPDNMGMTMIB_Cvpdntunneltable_Cvpdntunnelentry_Cvpdntunnelorigcause represents stack group (SGBP).
type CISCOVPDNMGMTMIB_Cvpdntunneltable_Cvpdntunnelentry_Cvpdntunnelorigcause string

const (
    CISCOVPDNMGMTMIB_Cvpdntunneltable_Cvpdntunnelentry_Cvpdntunnelorigcause_domain CISCOVPDNMGMTMIB_Cvpdntunneltable_Cvpdntunnelentry_Cvpdntunnelorigcause = "domain"

    CISCOVPDNMGMTMIB_Cvpdntunneltable_Cvpdntunnelentry_Cvpdntunnelorigcause_dnis CISCOVPDNMGMTMIB_Cvpdntunneltable_Cvpdntunnelentry_Cvpdntunnelorigcause = "dnis"

    CISCOVPDNMGMTMIB_Cvpdntunneltable_Cvpdntunnelentry_Cvpdntunnelorigcause_stack CISCOVPDNMGMTMIB_Cvpdntunneltable_Cvpdntunnelentry_Cvpdntunnelorigcause = "stack"
)

// CISCOVPDNMGMTMIB_Cvpdntunneltable_Cvpdntunnelentry_Cvpdntunnelstate represents                 to complete the process.
type CISCOVPDNMGMTMIB_Cvpdntunneltable_Cvpdntunnelentry_Cvpdntunnelstate string

const (
    CISCOVPDNMGMTMIB_Cvpdntunneltable_Cvpdntunnelentry_Cvpdntunnelstate_unknown CISCOVPDNMGMTMIB_Cvpdntunneltable_Cvpdntunnelentry_Cvpdntunnelstate = "unknown"

    CISCOVPDNMGMTMIB_Cvpdntunneltable_Cvpdntunnelentry_Cvpdntunnelstate_opening CISCOVPDNMGMTMIB_Cvpdntunneltable_Cvpdntunnelentry_Cvpdntunnelstate = "opening"

    CISCOVPDNMGMTMIB_Cvpdntunneltable_Cvpdntunnelentry_Cvpdntunnelstate_open CISCOVPDNMGMTMIB_Cvpdntunneltable_Cvpdntunnelentry_Cvpdntunnelstate = "open"

    CISCOVPDNMGMTMIB_Cvpdntunneltable_Cvpdntunnelentry_Cvpdntunnelstate_closing CISCOVPDNMGMTMIB_Cvpdntunneltable_Cvpdntunnelentry_Cvpdntunnelstate = "closing"
)

// CISCOVPDNMGMTMIB_Cvpdntunnelattrtable
// Table of information about the active VPDN tunnels.  An
// entry is added to the table when a new tunnel is initiated
// and removed from the table when the tunnel is terminated.
type CISCOVPDNMGMTMIB_Cvpdntunnelattrtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table, containing information about a single active VPDN
    // tunnel. The type is slice of
    // CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry.
    Cvpdntunnelattrentry []CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry
}

func (cvpdntunnelattrtable *CISCOVPDNMGMTMIB_Cvpdntunnelattrtable) GetEntityData() *types.CommonEntityData {
    cvpdntunnelattrtable.EntityData.YFilter = cvpdntunnelattrtable.YFilter
    cvpdntunnelattrtable.EntityData.YangName = "cvpdnTunnelAttrTable"
    cvpdntunnelattrtable.EntityData.BundleName = "cisco_ios_xe"
    cvpdntunnelattrtable.EntityData.ParentYangName = "CISCO-VPDN-MGMT-MIB"
    cvpdntunnelattrtable.EntityData.SegmentPath = "cvpdnTunnelAttrTable"
    cvpdntunnelattrtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvpdntunnelattrtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvpdntunnelattrtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvpdntunnelattrtable.EntityData.Children = make(map[string]types.YChild)
    cvpdntunnelattrtable.EntityData.Children["cvpdnTunnelAttrEntry"] = types.YChild{"Cvpdntunnelattrentry", nil}
    for i := range cvpdntunnelattrtable.Cvpdntunnelattrentry {
        cvpdntunnelattrtable.EntityData.Children[types.GetSegmentPath(&cvpdntunnelattrtable.Cvpdntunnelattrentry[i])] = types.YChild{"Cvpdntunnelattrentry", &cvpdntunnelattrtable.Cvpdntunnelattrentry[i]}
    }
    cvpdntunnelattrtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cvpdntunnelattrtable.EntityData)
}

// CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry
// An entry in the table, containing information about a
// single active VPDN tunnel.
type CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is TunnelType. Refers to
    // cisco_vpdn_mgmt_mib.CISCOVPDNMGMTMIB_Cvpdnsystemtable_Cvpdnsystementry_Cvpdnsystemtunneltype
    Cvpdnsystemtunneltype interface{}

    // This attribute is a key. The Tunnel ID of an active VPDN tunnel.  If this
    // end is the instigator of the tunnel, the ID is the TS tunnel ID, otherwise
    // it is the NAS tunnel ID.  Two distinct tunnels with the same tunnel ID may
    // exist, but with different tunnel types. The type is interface{} with range:
    // 0..65535.
    Cvpdntunnelattrtunnelid interface{}

    // The remote Tunnel ID of an active VPDN tunnel.  If this end is the
    // instigator of the tunnel, the ID is the NAS tunnel ID, otherwise it is the
    // TS tunnel ID. The type is interface{} with range: 0..65535.
    Cvpdntunnelattrremotetunnelid interface{}

    // The local name of an active VPDN tunnel.  It will be the NAS name of the
    // tunnel if the system serves as the NAS, or the TS name of the tunnel if the
    // system serves as the tunnel server.  Typically, the local name is the
    // configured host name of the system. The type is string with length: 1..255.
    Cvpdntunnelattrlocalname interface{}

    // The remote name of an active VPDN tunnel.  It will be the tunnel server
    // name of the tunnel if the system is a NAS, or the NAS name of the tunnel if
    // the system serves as the tunnel server. The type is string with length:
    // 1..255.
    Cvpdntunnelattrremotename interface{}

    // The remote end point name of an active VPDN tunnel.  This name is either
    // the domain name or the DNIS that this tunnel is projected with. The type is
    // string with length: 1..255.
    Cvpdntunnelattrremoteendpointname interface{}

    // This object indicates whether the tunnel was originated locally or not.  If
    // it's true, the tunnel was originated locally. The type is bool.
    Cvpdntunnelattrlocalinitconnection interface{}

    // The cause which originated an active VPDN tunnel.  The tunnel can be
    // projected via domain name, DNIS, stack group, or L2 Xconnect. The type is
    // Cvpdntunnelattrorigcause.
    Cvpdntunnelattrorigcause interface{}

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
    // been shut down. The type is Cvpdntunnelattrstate.
    Cvpdntunnelattrstate interface{}

    // The total number of active session currently in the tunnel. The type is
    // interface{} with range: 0..4294967295. Units are sessions.
    Cvpdntunnelattractivesessions interface{}

    // A count of the accumulated total of denied users for the tunnel. The type
    // is interface{} with range: 0..4294967295. Units are calls.
    Cvpdntunnelattrdeniedusers interface{}

    // A VPDN tunnel can be put into a soft shut state to prevent any new session
    // to be added.  This object specifies whether this tunnel has been soft shut.
    // If its true, it has been soft shut. The type is bool.
    Cvpdntunnelattrsoftshut interface{}

    // The type of network service used in the active tunnel. The type is
    // Cvpdntunnelattrnetworkservicetype.
    Cvpdntunnelattrnetworkservicetype interface{}

    // The local IP address of the tunnel.  This IP address is that of the
    // interface at the local end of the tunnel. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Cvpdntunnelattrlocalipaddress interface{}

    // The source IP address of the tunnel.  This IP address is the user
    // configurable IP address for Stack Group Biding Protocol. The type is string
    // with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Cvpdntunnelattrsourceipaddress interface{}

    // The remote IP address of the tunnel.  This IP address is that of the
    // interface at the remote end of the tunnel. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Cvpdntunnelattrremoteipaddress interface{}

    // Indicates the type of address contained in cvpdnTunnelAttrLocalInetAddress.
    // The type is InetAddressType.
    Cvpdntunnelattrlocalinetaddresstype interface{}

    // The local IP address of the tunnel.  This IP address is that of the
    // interface at the local end of the tunnel. The type of this address is
    // determined by the value of  cvpdnTunnelAttrLocalInetAddressType. The type
    // is string with length: 0..255.
    Cvpdntunnelattrlocalinetaddress interface{}

    // Indicates the type of address contained in
    // cvpdnTunnelAttrSourceInetAddress. The type is InetAddressType.
    Cvpdntunnelattrsourceinetaddresstype interface{}

    // The source IP address of the tunnel.  This IP address is the user
    // configurable IP address for Stack Group Biding Protocol.  The type of this
    // address is determined by the  value of
    // cvpdnTunnelAttrSourceInetAddressType. The type is string with length:
    // 0..255.
    Cvpdntunnelattrsourceinetaddress interface{}

    // Indicates the type of address contained in
    // cvpdnTunnelAttrRemoteInetAddress. The type is InetAddressType.
    Cvpdntunnelattrremoteinetaddresstype interface{}

    // The remote IP address of the tunnel.  This IP address is that of the
    // interface at the remote end of the tunnel. The type of this address is
    // determined by the value of  cvpdnTunnelAttrRemoteInetAddressType. The type
    // is string with length: 0..255.
    Cvpdntunnelattrremoteinetaddress interface{}
}

func (cvpdntunnelattrentry *CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry) GetEntityData() *types.CommonEntityData {
    cvpdntunnelattrentry.EntityData.YFilter = cvpdntunnelattrentry.YFilter
    cvpdntunnelattrentry.EntityData.YangName = "cvpdnTunnelAttrEntry"
    cvpdntunnelattrentry.EntityData.BundleName = "cisco_ios_xe"
    cvpdntunnelattrentry.EntityData.ParentYangName = "cvpdnTunnelAttrTable"
    cvpdntunnelattrentry.EntityData.SegmentPath = "cvpdnTunnelAttrEntry" + "[cvpdnSystemTunnelType='" + fmt.Sprintf("%v", cvpdntunnelattrentry.Cvpdnsystemtunneltype) + "']" + "[cvpdnTunnelAttrTunnelId='" + fmt.Sprintf("%v", cvpdntunnelattrentry.Cvpdntunnelattrtunnelid) + "']"
    cvpdntunnelattrentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvpdntunnelattrentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvpdntunnelattrentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvpdntunnelattrentry.EntityData.Children = make(map[string]types.YChild)
    cvpdntunnelattrentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cvpdntunnelattrentry.EntityData.Leafs["cvpdnSystemTunnelType"] = types.YLeaf{"Cvpdnsystemtunneltype", cvpdntunnelattrentry.Cvpdnsystemtunneltype}
    cvpdntunnelattrentry.EntityData.Leafs["cvpdnTunnelAttrTunnelId"] = types.YLeaf{"Cvpdntunnelattrtunnelid", cvpdntunnelattrentry.Cvpdntunnelattrtunnelid}
    cvpdntunnelattrentry.EntityData.Leafs["cvpdnTunnelAttrRemoteTunnelId"] = types.YLeaf{"Cvpdntunnelattrremotetunnelid", cvpdntunnelattrentry.Cvpdntunnelattrremotetunnelid}
    cvpdntunnelattrentry.EntityData.Leafs["cvpdnTunnelAttrLocalName"] = types.YLeaf{"Cvpdntunnelattrlocalname", cvpdntunnelattrentry.Cvpdntunnelattrlocalname}
    cvpdntunnelattrentry.EntityData.Leafs["cvpdnTunnelAttrRemoteName"] = types.YLeaf{"Cvpdntunnelattrremotename", cvpdntunnelattrentry.Cvpdntunnelattrremotename}
    cvpdntunnelattrentry.EntityData.Leafs["cvpdnTunnelAttrRemoteEndpointName"] = types.YLeaf{"Cvpdntunnelattrremoteendpointname", cvpdntunnelattrentry.Cvpdntunnelattrremoteendpointname}
    cvpdntunnelattrentry.EntityData.Leafs["cvpdnTunnelAttrLocalInitConnection"] = types.YLeaf{"Cvpdntunnelattrlocalinitconnection", cvpdntunnelattrentry.Cvpdntunnelattrlocalinitconnection}
    cvpdntunnelattrentry.EntityData.Leafs["cvpdnTunnelAttrOrigCause"] = types.YLeaf{"Cvpdntunnelattrorigcause", cvpdntunnelattrentry.Cvpdntunnelattrorigcause}
    cvpdntunnelattrentry.EntityData.Leafs["cvpdnTunnelAttrState"] = types.YLeaf{"Cvpdntunnelattrstate", cvpdntunnelattrentry.Cvpdntunnelattrstate}
    cvpdntunnelattrentry.EntityData.Leafs["cvpdnTunnelAttrActiveSessions"] = types.YLeaf{"Cvpdntunnelattractivesessions", cvpdntunnelattrentry.Cvpdntunnelattractivesessions}
    cvpdntunnelattrentry.EntityData.Leafs["cvpdnTunnelAttrDeniedUsers"] = types.YLeaf{"Cvpdntunnelattrdeniedusers", cvpdntunnelattrentry.Cvpdntunnelattrdeniedusers}
    cvpdntunnelattrentry.EntityData.Leafs["cvpdnTunnelAttrSoftshut"] = types.YLeaf{"Cvpdntunnelattrsoftshut", cvpdntunnelattrentry.Cvpdntunnelattrsoftshut}
    cvpdntunnelattrentry.EntityData.Leafs["cvpdnTunnelAttrNetworkServiceType"] = types.YLeaf{"Cvpdntunnelattrnetworkservicetype", cvpdntunnelattrentry.Cvpdntunnelattrnetworkservicetype}
    cvpdntunnelattrentry.EntityData.Leafs["cvpdnTunnelAttrLocalIpAddress"] = types.YLeaf{"Cvpdntunnelattrlocalipaddress", cvpdntunnelattrentry.Cvpdntunnelattrlocalipaddress}
    cvpdntunnelattrentry.EntityData.Leafs["cvpdnTunnelAttrSourceIpAddress"] = types.YLeaf{"Cvpdntunnelattrsourceipaddress", cvpdntunnelattrentry.Cvpdntunnelattrsourceipaddress}
    cvpdntunnelattrentry.EntityData.Leafs["cvpdnTunnelAttrRemoteIpAddress"] = types.YLeaf{"Cvpdntunnelattrremoteipaddress", cvpdntunnelattrentry.Cvpdntunnelattrremoteipaddress}
    cvpdntunnelattrentry.EntityData.Leafs["cvpdnTunnelAttrLocalInetAddressType"] = types.YLeaf{"Cvpdntunnelattrlocalinetaddresstype", cvpdntunnelattrentry.Cvpdntunnelattrlocalinetaddresstype}
    cvpdntunnelattrentry.EntityData.Leafs["cvpdnTunnelAttrLocalInetAddress"] = types.YLeaf{"Cvpdntunnelattrlocalinetaddress", cvpdntunnelattrentry.Cvpdntunnelattrlocalinetaddress}
    cvpdntunnelattrentry.EntityData.Leafs["cvpdnTunnelAttrSourceInetAddressType"] = types.YLeaf{"Cvpdntunnelattrsourceinetaddresstype", cvpdntunnelattrentry.Cvpdntunnelattrsourceinetaddresstype}
    cvpdntunnelattrentry.EntityData.Leafs["cvpdnTunnelAttrSourceInetAddress"] = types.YLeaf{"Cvpdntunnelattrsourceinetaddress", cvpdntunnelattrentry.Cvpdntunnelattrsourceinetaddress}
    cvpdntunnelattrentry.EntityData.Leafs["cvpdnTunnelAttrRemoteInetAddressType"] = types.YLeaf{"Cvpdntunnelattrremoteinetaddresstype", cvpdntunnelattrentry.Cvpdntunnelattrremoteinetaddresstype}
    cvpdntunnelattrentry.EntityData.Leafs["cvpdnTunnelAttrRemoteInetAddress"] = types.YLeaf{"Cvpdntunnelattrremoteinetaddress", cvpdntunnelattrentry.Cvpdntunnelattrremoteinetaddress}
    return &(cvpdntunnelattrentry.EntityData)
}

// CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry_Cvpdntunnelattrnetworkservicetype represents The type of network service used in the active tunnel.
type CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry_Cvpdntunnelattrnetworkservicetype string

const (
    CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry_Cvpdntunnelattrnetworkservicetype_ip CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry_Cvpdntunnelattrnetworkservicetype = "ip"
)

// CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry_Cvpdntunnelattrorigcause represents or L2 Xconnect.
type CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry_Cvpdntunnelattrorigcause string

const (
    CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry_Cvpdntunnelattrorigcause_domain CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry_Cvpdntunnelattrorigcause = "domain"

    CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry_Cvpdntunnelattrorigcause_dnis CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry_Cvpdntunnelattrorigcause = "dnis"

    CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry_Cvpdntunnelattrorigcause_stack CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry_Cvpdntunnelattrorigcause = "stack"

    CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry_Cvpdntunnelattrorigcause_xconnect CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry_Cvpdntunnelattrorigcause = "xconnect"
)

// CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry_Cvpdntunnelattrstate represents     pptpTerminal:       The tunnel has been shut down.
type CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry_Cvpdntunnelattrstate string

const (
    CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry_Cvpdntunnelattrstate_unknown CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry_Cvpdntunnelattrstate = "unknown"

    CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry_Cvpdntunnelattrstate_l2fOpening CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry_Cvpdntunnelattrstate = "l2fOpening"

    CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry_Cvpdntunnelattrstate_l2fOpenWait CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry_Cvpdntunnelattrstate = "l2fOpenWait"

    CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry_Cvpdntunnelattrstate_l2fOpen CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry_Cvpdntunnelattrstate = "l2fOpen"

    CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry_Cvpdntunnelattrstate_l2fClosing CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry_Cvpdntunnelattrstate = "l2fClosing"

    CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry_Cvpdntunnelattrstate_l2fCloseWait CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry_Cvpdntunnelattrstate = "l2fCloseWait"

    CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry_Cvpdntunnelattrstate_l2tpIdle CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry_Cvpdntunnelattrstate = "l2tpIdle"

    CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry_Cvpdntunnelattrstate_l2tpWaitCtlReply CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry_Cvpdntunnelattrstate = "l2tpWaitCtlReply"

    CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry_Cvpdntunnelattrstate_l2tpEstablished CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry_Cvpdntunnelattrstate = "l2tpEstablished"

    CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry_Cvpdntunnelattrstate_l2tpShuttingDown CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry_Cvpdntunnelattrstate = "l2tpShuttingDown"

    CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry_Cvpdntunnelattrstate_l2tpNoSessionLeft CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry_Cvpdntunnelattrstate = "l2tpNoSessionLeft"

    CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry_Cvpdntunnelattrstate_pptpIdle CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry_Cvpdntunnelattrstate = "pptpIdle"

    CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry_Cvpdntunnelattrstate_pptpWaitConnect CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry_Cvpdntunnelattrstate = "pptpWaitConnect"

    CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry_Cvpdntunnelattrstate_pptpWaitCtlRequest CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry_Cvpdntunnelattrstate = "pptpWaitCtlRequest"

    CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry_Cvpdntunnelattrstate_pptpWaitCtlReply CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry_Cvpdntunnelattrstate = "pptpWaitCtlReply"

    CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry_Cvpdntunnelattrstate_pptpEstablished CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry_Cvpdntunnelattrstate = "pptpEstablished"

    CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry_Cvpdntunnelattrstate_pptpWaitStopReply CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry_Cvpdntunnelattrstate = "pptpWaitStopReply"

    CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry_Cvpdntunnelattrstate_pptpTerminal CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry_Cvpdntunnelattrstate = "pptpTerminal"
)

// CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable
// Table of information about individual user sessions
// within the active tunnels.  Entry is added to the table
// when new user session is initiated and be removed from
// the table when the user session is terminated.
type CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table, containing information about a single user session
    // within the tunnel. The type is slice of
    // CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable_Cvpdntunnelsessionentry.
    Cvpdntunnelsessionentry []CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable_Cvpdntunnelsessionentry
}

func (cvpdntunnelsessiontable *CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable) GetEntityData() *types.CommonEntityData {
    cvpdntunnelsessiontable.EntityData.YFilter = cvpdntunnelsessiontable.YFilter
    cvpdntunnelsessiontable.EntityData.YangName = "cvpdnTunnelSessionTable"
    cvpdntunnelsessiontable.EntityData.BundleName = "cisco_ios_xe"
    cvpdntunnelsessiontable.EntityData.ParentYangName = "CISCO-VPDN-MGMT-MIB"
    cvpdntunnelsessiontable.EntityData.SegmentPath = "cvpdnTunnelSessionTable"
    cvpdntunnelsessiontable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvpdntunnelsessiontable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvpdntunnelsessiontable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvpdntunnelsessiontable.EntityData.Children = make(map[string]types.YChild)
    cvpdntunnelsessiontable.EntityData.Children["cvpdnTunnelSessionEntry"] = types.YChild{"Cvpdntunnelsessionentry", nil}
    for i := range cvpdntunnelsessiontable.Cvpdntunnelsessionentry {
        cvpdntunnelsessiontable.EntityData.Children[types.GetSegmentPath(&cvpdntunnelsessiontable.Cvpdntunnelsessionentry[i])] = types.YChild{"Cvpdntunnelsessionentry", &cvpdntunnelsessiontable.Cvpdntunnelsessionentry[i]}
    }
    cvpdntunnelsessiontable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cvpdntunnelsessiontable.EntityData)
}

// CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable_Cvpdntunnelsessionentry
// An entry in the table, containing information about a
// single user session within the tunnel.
type CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable_Cvpdntunnelsessionentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // cisco_vpdn_mgmt_mib.CISCOVPDNMGMTMIB_Cvpdntunneltable_Cvpdntunnelentry_Cvpdntunneltunnelid
    Cvpdntunneltunnelid interface{}

    // This attribute is a key. The ID of an active VPDN tunnel user session. The
    // type is interface{} with range: 0..4294967295.
    Cvpdntunnelsessionid interface{}

    // The name of the user of the user session. The type is string with length:
    // 1..255.
    Cvpdntunnelsessionusername interface{}

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
    // Cvpdntunnelsessionstate.
    Cvpdntunnelsessionstate interface{}

    // This object specifies the call duration of the current active user session
    // in value of system uptime. The type is interface{} with range:
    // 0..4294967295.
    Cvpdntunnelsessioncallduration interface{}

    // The total number of output packets through this active user session. The
    // type is interface{} with range: 0..4294967295. Units are packets.
    Cvpdntunnelsessionpacketsout interface{}

    // The total number of output bytes through this active user session. The type
    // is interface{} with range: 0..4294967295. Units are bytes.
    Cvpdntunnelsessionbytesout interface{}

    // The total number of input packets through this active user session. The
    // type is interface{} with range: 0..4294967295. Units are packets.
    Cvpdntunnelsessionpacketsin interface{}

    // The total number of input bytes through this active user session. The type
    // is interface{} with range: 0..4294967295. Units are bytes.
    Cvpdntunnelsessionbytesin interface{}

    // The type of physical devices that this user session is attached to for the
    // local end of the tunnel.  The meaning of each device type is explained
    // below:      other:              Any device that has not been               
    // defined.      asyncInternalModem: Modem Pool device of an access           
    // server.      async:              A regular asynchronous serial             
    // interface.      sync:               A regular synchronous serial           
    // interface.      bchan:              An ISDN call.      xdsl:              
    // Future application with xDSL                         devices.      cable:  
    // Future application with Cable                         modem devices. The
    // type is Cvpdntunnelsessiondevicetype.
    Cvpdntunnelsessiondevicetype interface{}

    // The incoming caller identification of the user.  It is the originating
    // number that called into the device that initiates the user session.  This
    // object can be empty since not all dial device can provide caller ID
    // information. The type is string.
    Cvpdntunnelsessiondevicecallerid interface{}

    // The device ID of the physical interface for the user session.  The object
    // is the the interface index which points to the ifTable.  For virtual
    // interface that is not in the ifTable, it will have zero value. The type is
    // interface{} with range: 0..2147483647.
    Cvpdntunnelsessiondevicephyid interface{}

    // This object indicates whether the session is part of a multilink or not.
    // The type is bool.
    Cvpdntunnelsessionmultilink interface{}

    // The Modem Pool database slot index if it is associated with this user
    // session.  Only a session with device of type asyncInternalModem will have a
    // valid value for this object. The type is interface{} with range:
    // 0..4294967295.
    Cvpdntunnelsessionmodemslotindex interface{}

    // The Modem Pool database port index if it is associated with this user
    // session.  Only a session with a device of type asyncInternalModem will have
    // a valid value for this object. The type is interface{} with range:
    // 0..4294967295.
    Cvpdntunnelsessionmodemportindex interface{}

    // The DS1 database slot index if it is associated with this user session. 
    // Only a session with a device of type asyncInternalModem will have a valid
    // value for this object. The type is interface{} with range: 0..4294967295.
    Cvpdntunnelsessionds1Slotindex interface{}

    // The DS1 database port index if it is associated with this user session. 
    // Only a session with a device of type asyncInternalModem will have a a valid
    // value for this object. The type is interface{} with range: 0..4294967295.
    Cvpdntunnelsessionds1Portindex interface{}

    // The DS1 database channel index if it is associated with this user session. 
    // Only a session over a device of type asyncInternalModem will have a valid
    // value for this object. The type is interface{} with range: 0..4294967295.
    Cvpdntunnelsessionds1Channelindex interface{}

    // The start time of the current modem call.  Only a session with a  device of
    // type asyncInternalModem will have a valid value for this object. The type
    // is interface{} with range: 0..4294967295.
    Cvpdntunnelsessionmodemcallstarttime interface{}

    // Arbitrary small integer to distinguish modem calls that occurred at the
    // same time tick.  Only session over device asyncInternalModem will have a
    // valid value for this object. The type is interface{} with range:
    // 0..4294967295.
    Cvpdntunnelsessionmodemcallstartindex interface{}
}

func (cvpdntunnelsessionentry *CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable_Cvpdntunnelsessionentry) GetEntityData() *types.CommonEntityData {
    cvpdntunnelsessionentry.EntityData.YFilter = cvpdntunnelsessionentry.YFilter
    cvpdntunnelsessionentry.EntityData.YangName = "cvpdnTunnelSessionEntry"
    cvpdntunnelsessionentry.EntityData.BundleName = "cisco_ios_xe"
    cvpdntunnelsessionentry.EntityData.ParentYangName = "cvpdnTunnelSessionTable"
    cvpdntunnelsessionentry.EntityData.SegmentPath = "cvpdnTunnelSessionEntry" + "[cvpdnTunnelTunnelId='" + fmt.Sprintf("%v", cvpdntunnelsessionentry.Cvpdntunneltunnelid) + "']" + "[cvpdnTunnelSessionId='" + fmt.Sprintf("%v", cvpdntunnelsessionentry.Cvpdntunnelsessionid) + "']"
    cvpdntunnelsessionentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvpdntunnelsessionentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvpdntunnelsessionentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvpdntunnelsessionentry.EntityData.Children = make(map[string]types.YChild)
    cvpdntunnelsessionentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cvpdntunnelsessionentry.EntityData.Leafs["cvpdnTunnelTunnelId"] = types.YLeaf{"Cvpdntunneltunnelid", cvpdntunnelsessionentry.Cvpdntunneltunnelid}
    cvpdntunnelsessionentry.EntityData.Leafs["cvpdnTunnelSessionId"] = types.YLeaf{"Cvpdntunnelsessionid", cvpdntunnelsessionentry.Cvpdntunnelsessionid}
    cvpdntunnelsessionentry.EntityData.Leafs["cvpdnTunnelSessionUserName"] = types.YLeaf{"Cvpdntunnelsessionusername", cvpdntunnelsessionentry.Cvpdntunnelsessionusername}
    cvpdntunnelsessionentry.EntityData.Leafs["cvpdnTunnelSessionState"] = types.YLeaf{"Cvpdntunnelsessionstate", cvpdntunnelsessionentry.Cvpdntunnelsessionstate}
    cvpdntunnelsessionentry.EntityData.Leafs["cvpdnTunnelSessionCallDuration"] = types.YLeaf{"Cvpdntunnelsessioncallduration", cvpdntunnelsessionentry.Cvpdntunnelsessioncallduration}
    cvpdntunnelsessionentry.EntityData.Leafs["cvpdnTunnelSessionPacketsOut"] = types.YLeaf{"Cvpdntunnelsessionpacketsout", cvpdntunnelsessionentry.Cvpdntunnelsessionpacketsout}
    cvpdntunnelsessionentry.EntityData.Leafs["cvpdnTunnelSessionBytesOut"] = types.YLeaf{"Cvpdntunnelsessionbytesout", cvpdntunnelsessionentry.Cvpdntunnelsessionbytesout}
    cvpdntunnelsessionentry.EntityData.Leafs["cvpdnTunnelSessionPacketsIn"] = types.YLeaf{"Cvpdntunnelsessionpacketsin", cvpdntunnelsessionentry.Cvpdntunnelsessionpacketsin}
    cvpdntunnelsessionentry.EntityData.Leafs["cvpdnTunnelSessionBytesIn"] = types.YLeaf{"Cvpdntunnelsessionbytesin", cvpdntunnelsessionentry.Cvpdntunnelsessionbytesin}
    cvpdntunnelsessionentry.EntityData.Leafs["cvpdnTunnelSessionDeviceType"] = types.YLeaf{"Cvpdntunnelsessiondevicetype", cvpdntunnelsessionentry.Cvpdntunnelsessiondevicetype}
    cvpdntunnelsessionentry.EntityData.Leafs["cvpdnTunnelSessionDeviceCallerId"] = types.YLeaf{"Cvpdntunnelsessiondevicecallerid", cvpdntunnelsessionentry.Cvpdntunnelsessiondevicecallerid}
    cvpdntunnelsessionentry.EntityData.Leafs["cvpdnTunnelSessionDevicePhyId"] = types.YLeaf{"Cvpdntunnelsessiondevicephyid", cvpdntunnelsessionentry.Cvpdntunnelsessiondevicephyid}
    cvpdntunnelsessionentry.EntityData.Leafs["cvpdnTunnelSessionMultilink"] = types.YLeaf{"Cvpdntunnelsessionmultilink", cvpdntunnelsessionentry.Cvpdntunnelsessionmultilink}
    cvpdntunnelsessionentry.EntityData.Leafs["cvpdnTunnelSessionModemSlotIndex"] = types.YLeaf{"Cvpdntunnelsessionmodemslotindex", cvpdntunnelsessionentry.Cvpdntunnelsessionmodemslotindex}
    cvpdntunnelsessionentry.EntityData.Leafs["cvpdnTunnelSessionModemPortIndex"] = types.YLeaf{"Cvpdntunnelsessionmodemportindex", cvpdntunnelsessionentry.Cvpdntunnelsessionmodemportindex}
    cvpdntunnelsessionentry.EntityData.Leafs["cvpdnTunnelSessionDS1SlotIndex"] = types.YLeaf{"Cvpdntunnelsessionds1Slotindex", cvpdntunnelsessionentry.Cvpdntunnelsessionds1Slotindex}
    cvpdntunnelsessionentry.EntityData.Leafs["cvpdnTunnelSessionDS1PortIndex"] = types.YLeaf{"Cvpdntunnelsessionds1Portindex", cvpdntunnelsessionentry.Cvpdntunnelsessionds1Portindex}
    cvpdntunnelsessionentry.EntityData.Leafs["cvpdnTunnelSessionDS1ChannelIndex"] = types.YLeaf{"Cvpdntunnelsessionds1Channelindex", cvpdntunnelsessionentry.Cvpdntunnelsessionds1Channelindex}
    cvpdntunnelsessionentry.EntityData.Leafs["cvpdnTunnelSessionModemCallStartTime"] = types.YLeaf{"Cvpdntunnelsessionmodemcallstarttime", cvpdntunnelsessionentry.Cvpdntunnelsessionmodemcallstarttime}
    cvpdntunnelsessionentry.EntityData.Leafs["cvpdnTunnelSessionModemCallStartIndex"] = types.YLeaf{"Cvpdntunnelsessionmodemcallstartindex", cvpdntunnelsessionentry.Cvpdntunnelsessionmodemcallstartindex}
    return &(cvpdntunnelsessionentry.EntityData)
}

// CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable_Cvpdntunnelsessionentry_Cvpdntunnelsessiondevicetype represents                         modem devices.
type CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable_Cvpdntunnelsessionentry_Cvpdntunnelsessiondevicetype string

const (
    CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable_Cvpdntunnelsessionentry_Cvpdntunnelsessiondevicetype_other CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable_Cvpdntunnelsessionentry_Cvpdntunnelsessiondevicetype = "other"

    CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable_Cvpdntunnelsessionentry_Cvpdntunnelsessiondevicetype_asyncInternalModem CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable_Cvpdntunnelsessionentry_Cvpdntunnelsessiondevicetype = "asyncInternalModem"

    CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable_Cvpdntunnelsessionentry_Cvpdntunnelsessiondevicetype_async CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable_Cvpdntunnelsessionentry_Cvpdntunnelsessiondevicetype = "async"

    CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable_Cvpdntunnelsessionentry_Cvpdntunnelsessiondevicetype_bchan CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable_Cvpdntunnelsessionentry_Cvpdntunnelsessiondevicetype = "bchan"

    CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable_Cvpdntunnelsessionentry_Cvpdntunnelsessiondevicetype_sync CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable_Cvpdntunnelsessionentry_Cvpdntunnelsessiondevicetype = "sync"

    CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable_Cvpdntunnelsessionentry_Cvpdntunnelsessiondevicetype_virtualAccess CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable_Cvpdntunnelsessionentry_Cvpdntunnelsessiondevicetype = "virtualAccess"

    CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable_Cvpdntunnelsessionentry_Cvpdntunnelsessiondevicetype_xdsl CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable_Cvpdntunnelsessionentry_Cvpdntunnelsessiondevicetype = "xdsl"

    CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable_Cvpdntunnelsessionentry_Cvpdntunnelsessiondevicetype_cable CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable_Cvpdntunnelsessionentry_Cvpdntunnelsessiondevicetype = "cable"
)

// CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable_Cvpdntunnelsessionentry_Cvpdntunnelsessionstate represents                       established.
type CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable_Cvpdntunnelsessionentry_Cvpdntunnelsessionstate string

const (
    CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable_Cvpdntunnelsessionentry_Cvpdntunnelsessionstate_unknown CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable_Cvpdntunnelsessionentry_Cvpdntunnelsessionstate = "unknown"

    CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable_Cvpdntunnelsessionentry_Cvpdntunnelsessionstate_opening CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable_Cvpdntunnelsessionentry_Cvpdntunnelsessionstate = "opening"

    CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable_Cvpdntunnelsessionentry_Cvpdntunnelsessionstate_open CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable_Cvpdntunnelsessionentry_Cvpdntunnelsessionstate = "open"

    CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable_Cvpdntunnelsessionentry_Cvpdntunnelsessionstate_closing CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable_Cvpdntunnelsessionentry_Cvpdntunnelsessionstate = "closing"

    CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable_Cvpdntunnelsessionentry_Cvpdntunnelsessionstate_waitingForTunnel CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable_Cvpdntunnelsessionentry_Cvpdntunnelsessionstate = "waitingForTunnel"
)

// CISCOVPDNMGMTMIB_Cvpdnsessionattrtable
// Table of information about individual sessions within the
// active tunnels.  An entry is added to the table when a new
// session is initiated and removed from the table when the
// session is terminated.
type CISCOVPDNMGMTMIB_Cvpdnsessionattrtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table, containing information about a single session within
    // the tunnel. The type is slice of
    // CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry.
    Cvpdnsessionattrentry []CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry
}

func (cvpdnsessionattrtable *CISCOVPDNMGMTMIB_Cvpdnsessionattrtable) GetEntityData() *types.CommonEntityData {
    cvpdnsessionattrtable.EntityData.YFilter = cvpdnsessionattrtable.YFilter
    cvpdnsessionattrtable.EntityData.YangName = "cvpdnSessionAttrTable"
    cvpdnsessionattrtable.EntityData.BundleName = "cisco_ios_xe"
    cvpdnsessionattrtable.EntityData.ParentYangName = "CISCO-VPDN-MGMT-MIB"
    cvpdnsessionattrtable.EntityData.SegmentPath = "cvpdnSessionAttrTable"
    cvpdnsessionattrtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvpdnsessionattrtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvpdnsessionattrtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvpdnsessionattrtable.EntityData.Children = make(map[string]types.YChild)
    cvpdnsessionattrtable.EntityData.Children["cvpdnSessionAttrEntry"] = types.YChild{"Cvpdnsessionattrentry", nil}
    for i := range cvpdnsessionattrtable.Cvpdnsessionattrentry {
        cvpdnsessionattrtable.EntityData.Children[types.GetSegmentPath(&cvpdnsessionattrtable.Cvpdnsessionattrentry[i])] = types.YChild{"Cvpdnsessionattrentry", &cvpdnsessionattrtable.Cvpdnsessionattrentry[i]}
    }
    cvpdnsessionattrtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cvpdnsessionattrtable.EntityData)
}

// CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry
// An entry in the table, containing information about a
// single session within the tunnel.
type CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is TunnelType. Refers to
    // cisco_vpdn_mgmt_mib.CISCOVPDNMGMTMIB_Cvpdnsystemtable_Cvpdnsystementry_Cvpdnsystemtunneltype
    Cvpdnsystemtunneltype interface{}

    // This attribute is a key. The type is string with range: 0..65535. Refers to
    // cisco_vpdn_mgmt_mib.CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry_Cvpdntunnelattrtunnelid
    Cvpdntunnelattrtunnelid interface{}

    // This attribute is a key. The ID of an active VPDN session. The type is
    // interface{} with range: 0..65535.
    Cvpdnsessionattrsessionid interface{}

    // The name of the user of the session. The type is string with length:
    // 1..255.
    Cvpdnsessionattrusername interface{}

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
    // Cvpdnsessionattrstate.
    Cvpdnsessionattrstate interface{}

    // This object specifies the call duration of the current active session. The
    // type is interface{} with range: 0..4294967295.
    Cvpdnsessionattrcallduration interface{}

    // The total number of output packets through this active session. The type is
    // interface{} with range: 0..4294967295. Units are packets.
    Cvpdnsessionattrpacketsout interface{}

    // The total number of output bytes through this active session. The type is
    // interface{} with range: 0..4294967295. Units are bytes.
    Cvpdnsessionattrbytesout interface{}

    // The total number of input packets through this active session. The type is
    // interface{} with range: 0..4294967295. Units are packets.
    Cvpdnsessionattrpacketsin interface{}

    // The total number of input bytes through this active session. The type is
    // interface{} with range: 0..4294967295. Units are bytes.
    Cvpdnsessionattrbytesin interface{}

    // The type of physical devices that this session is attached to for the local
    // end of the tunnel.  The meaning of each device type is explained below:    
    // other:              Any device that has not been                        
    // defined.      asyncInternalModem: Modem Pool device of an access           
    // server.      async:              A regular asynchronous serial             
    // interface.      sync:               A regular synchronous serial           
    // interface.      bchan:              An ISDN call.      xdsl:              
    // xDSL interface.      cable:              cable modem interface. The type is
    // Cvpdnsessionattrdevicetype.
    Cvpdnsessionattrdevicetype interface{}

    // The incoming caller identification of the user.  It is the originating
    // number that called into the device that initiates the session.  This object
    // can be empty since not all dial devices can provide caller ID information.
    // The type is string.
    Cvpdnsessionattrdevicecallerid interface{}

    // The device ID of the physical interface for the session. The object is the
    // the interface index which points to the ifTable.  For virtual interfaces
    // that are not in the ifTable, the value will be zero. The type is
    // interface{} with range: 0..2147483647.
    Cvpdnsessionattrdevicephyid interface{}

    // This object indicates whether the session is part of a multilink PPP
    // bundle, even if there is only one link or session in the bundle.  If it is
    // multilink PPP, the value is true. The type is bool.
    Cvpdnsessionattrmultilink interface{}

    // The Modem Pool database slot index if it is associated with this session. 
    // Only a session with device of type 'asyncInternalModem' will have a valid
    // value for this object; otherwise, it is not instantiated. The type is
    // interface{} with range: 0..4294967295.
    Cvpdnsessionattrmodemslotindex interface{}

    // The Modem Pool database port index if it is associated with this session. 
    // Only a session with a device of type 'asyncInternalModem' will have a valid
    // value for this object; otherwise, it is not instantiated. The type is
    // interface{} with range: 0..4294967295.
    Cvpdnsessionattrmodemportindex interface{}

    // The DS1 database slot index if it is associated with this session.  Only a
    // session with a device of type 'asyncInternalModem' will have a valid value
    // for this object; otherwise it is not instantiated. The type is interface{}
    // with range: 0..4294967295.
    Cvpdnsessionattrds1Slotindex interface{}

    // The DS1 database port index if it is associated with this session.  Only a
    // session with a device of type 'asyncInternalModem' will have a valid value
    // for this object; otherwise it is not instantiated. The type is interface{}
    // with range: 0..4294967295.
    Cvpdnsessionattrds1Portindex interface{}

    // The DS1 database channel index if it is associated with this session.  Only
    // a session over a device of type 'asyncInternalModem' will have a valid
    // value for this object; otherwise it is not instantiated. The type is
    // interface{} with range: 0..4294967295.
    Cvpdnsessionattrds1Channelindex interface{}

    // The start time of the current modem call.  Only a session with a device of
    // type 'asyncInternalModem' will have a valid value for this object;
    // otherwise, it is not instantiated. The type is interface{} with range:
    // 0..4294967295.
    Cvpdnsessionattrmodemcallstarttime interface{}

    // Arbitrary small integer to distinguish modem calls that occurred at the
    // same time tick.  Only session over device 'asyncInternalModem' will have a
    // valid value for this object; otherwise, it is not instantiated. The type is
    // interface{} with range: 0..4294967295.
    Cvpdnsessionattrmodemcallstartindex interface{}

    // The virtual circuit ID of an active Layer 2 VPN session. The type is
    // interface{} with range: 0..4294967295.
    Cvpdnsessionattrvirtualcircuitid interface{}

    // The total number of dropped packets that could not be sent into this active
    // session. The type is interface{} with range: 0..4294967295. Units are
    // packets.
    Cvpdnsessionattrsentpktsdropped interface{}

    // The total number of dropped packets that were received from this active
    // session. The type is interface{} with range: 0..4294967295. Units are
    // packets.
    Cvpdnsessionattrrecvpktsdropped interface{}

    // This object specifies the name of the multilink bundle to which this
    // session belongs.  The value of this object is only valid when the value of
    // cvpdnSessionAttrMultilink is 'true'.  If the value of
    // cvpdnSessionAttrMultilink is 'false', then the value of this object will be
    // the empty string. The type is string with length: 0..255.
    Cvpdnsessionattrmultilinkbundle interface{}

    // This object specifies the ifIndex of the multilink bundle to which this
    // session belongs.  The value of this object is only valid when the value of
    // cvpdnSessionAttrMultilink is 'true'.  If the value of
    // cvpdnSessionAttrMultilink is 'false', then the value of this object will be
    // zero. The type is interface{} with range: 0..2147483647.
    Cvpdnsessionattrmultilinkifindex interface{}
}

func (cvpdnsessionattrentry *CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry) GetEntityData() *types.CommonEntityData {
    cvpdnsessionattrentry.EntityData.YFilter = cvpdnsessionattrentry.YFilter
    cvpdnsessionattrentry.EntityData.YangName = "cvpdnSessionAttrEntry"
    cvpdnsessionattrentry.EntityData.BundleName = "cisco_ios_xe"
    cvpdnsessionattrentry.EntityData.ParentYangName = "cvpdnSessionAttrTable"
    cvpdnsessionattrentry.EntityData.SegmentPath = "cvpdnSessionAttrEntry" + "[cvpdnSystemTunnelType='" + fmt.Sprintf("%v", cvpdnsessionattrentry.Cvpdnsystemtunneltype) + "']" + "[cvpdnTunnelAttrTunnelId='" + fmt.Sprintf("%v", cvpdnsessionattrentry.Cvpdntunnelattrtunnelid) + "']" + "[cvpdnSessionAttrSessionId='" + fmt.Sprintf("%v", cvpdnsessionattrentry.Cvpdnsessionattrsessionid) + "']"
    cvpdnsessionattrentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvpdnsessionattrentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvpdnsessionattrentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvpdnsessionattrentry.EntityData.Children = make(map[string]types.YChild)
    cvpdnsessionattrentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cvpdnsessionattrentry.EntityData.Leafs["cvpdnSystemTunnelType"] = types.YLeaf{"Cvpdnsystemtunneltype", cvpdnsessionattrentry.Cvpdnsystemtunneltype}
    cvpdnsessionattrentry.EntityData.Leafs["cvpdnTunnelAttrTunnelId"] = types.YLeaf{"Cvpdntunnelattrtunnelid", cvpdnsessionattrentry.Cvpdntunnelattrtunnelid}
    cvpdnsessionattrentry.EntityData.Leafs["cvpdnSessionAttrSessionId"] = types.YLeaf{"Cvpdnsessionattrsessionid", cvpdnsessionattrentry.Cvpdnsessionattrsessionid}
    cvpdnsessionattrentry.EntityData.Leafs["cvpdnSessionAttrUserName"] = types.YLeaf{"Cvpdnsessionattrusername", cvpdnsessionattrentry.Cvpdnsessionattrusername}
    cvpdnsessionattrentry.EntityData.Leafs["cvpdnSessionAttrState"] = types.YLeaf{"Cvpdnsessionattrstate", cvpdnsessionattrentry.Cvpdnsessionattrstate}
    cvpdnsessionattrentry.EntityData.Leafs["cvpdnSessionAttrCallDuration"] = types.YLeaf{"Cvpdnsessionattrcallduration", cvpdnsessionattrentry.Cvpdnsessionattrcallduration}
    cvpdnsessionattrentry.EntityData.Leafs["cvpdnSessionAttrPacketsOut"] = types.YLeaf{"Cvpdnsessionattrpacketsout", cvpdnsessionattrentry.Cvpdnsessionattrpacketsout}
    cvpdnsessionattrentry.EntityData.Leafs["cvpdnSessionAttrBytesOut"] = types.YLeaf{"Cvpdnsessionattrbytesout", cvpdnsessionattrentry.Cvpdnsessionattrbytesout}
    cvpdnsessionattrentry.EntityData.Leafs["cvpdnSessionAttrPacketsIn"] = types.YLeaf{"Cvpdnsessionattrpacketsin", cvpdnsessionattrentry.Cvpdnsessionattrpacketsin}
    cvpdnsessionattrentry.EntityData.Leafs["cvpdnSessionAttrBytesIn"] = types.YLeaf{"Cvpdnsessionattrbytesin", cvpdnsessionattrentry.Cvpdnsessionattrbytesin}
    cvpdnsessionattrentry.EntityData.Leafs["cvpdnSessionAttrDeviceType"] = types.YLeaf{"Cvpdnsessionattrdevicetype", cvpdnsessionattrentry.Cvpdnsessionattrdevicetype}
    cvpdnsessionattrentry.EntityData.Leafs["cvpdnSessionAttrDeviceCallerId"] = types.YLeaf{"Cvpdnsessionattrdevicecallerid", cvpdnsessionattrentry.Cvpdnsessionattrdevicecallerid}
    cvpdnsessionattrentry.EntityData.Leafs["cvpdnSessionAttrDevicePhyId"] = types.YLeaf{"Cvpdnsessionattrdevicephyid", cvpdnsessionattrentry.Cvpdnsessionattrdevicephyid}
    cvpdnsessionattrentry.EntityData.Leafs["cvpdnSessionAttrMultilink"] = types.YLeaf{"Cvpdnsessionattrmultilink", cvpdnsessionattrentry.Cvpdnsessionattrmultilink}
    cvpdnsessionattrentry.EntityData.Leafs["cvpdnSessionAttrModemSlotIndex"] = types.YLeaf{"Cvpdnsessionattrmodemslotindex", cvpdnsessionattrentry.Cvpdnsessionattrmodemslotindex}
    cvpdnsessionattrentry.EntityData.Leafs["cvpdnSessionAttrModemPortIndex"] = types.YLeaf{"Cvpdnsessionattrmodemportindex", cvpdnsessionattrentry.Cvpdnsessionattrmodemportindex}
    cvpdnsessionattrentry.EntityData.Leafs["cvpdnSessionAttrDS1SlotIndex"] = types.YLeaf{"Cvpdnsessionattrds1Slotindex", cvpdnsessionattrentry.Cvpdnsessionattrds1Slotindex}
    cvpdnsessionattrentry.EntityData.Leafs["cvpdnSessionAttrDS1PortIndex"] = types.YLeaf{"Cvpdnsessionattrds1Portindex", cvpdnsessionattrentry.Cvpdnsessionattrds1Portindex}
    cvpdnsessionattrentry.EntityData.Leafs["cvpdnSessionAttrDS1ChannelIndex"] = types.YLeaf{"Cvpdnsessionattrds1Channelindex", cvpdnsessionattrentry.Cvpdnsessionattrds1Channelindex}
    cvpdnsessionattrentry.EntityData.Leafs["cvpdnSessionAttrModemCallStartTime"] = types.YLeaf{"Cvpdnsessionattrmodemcallstarttime", cvpdnsessionattrentry.Cvpdnsessionattrmodemcallstarttime}
    cvpdnsessionattrentry.EntityData.Leafs["cvpdnSessionAttrModemCallStartIndex"] = types.YLeaf{"Cvpdnsessionattrmodemcallstartindex", cvpdnsessionattrentry.Cvpdnsessionattrmodemcallstartindex}
    cvpdnsessionattrentry.EntityData.Leafs["cvpdnSessionAttrVirtualCircuitID"] = types.YLeaf{"Cvpdnsessionattrvirtualcircuitid", cvpdnsessionattrentry.Cvpdnsessionattrvirtualcircuitid}
    cvpdnsessionattrentry.EntityData.Leafs["cvpdnSessionAttrSentPktsDropped"] = types.YLeaf{"Cvpdnsessionattrsentpktsdropped", cvpdnsessionattrentry.Cvpdnsessionattrsentpktsdropped}
    cvpdnsessionattrentry.EntityData.Leafs["cvpdnSessionAttrRecvPktsDropped"] = types.YLeaf{"Cvpdnsessionattrrecvpktsdropped", cvpdnsessionattrentry.Cvpdnsessionattrrecvpktsdropped}
    cvpdnsessionattrentry.EntityData.Leafs["cvpdnSessionAttrMultilinkBundle"] = types.YLeaf{"Cvpdnsessionattrmultilinkbundle", cvpdnsessionattrentry.Cvpdnsessionattrmultilinkbundle}
    cvpdnsessionattrentry.EntityData.Leafs["cvpdnSessionAttrMultilinkIfIndex"] = types.YLeaf{"Cvpdnsessionattrmultilinkifindex", cvpdnsessionattrentry.Cvpdnsessionattrmultilinkifindex}
    return &(cvpdnsessionattrentry.EntityData)
}

// CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry_Cvpdnsessionattrdevicetype represents     cable:              cable modem interface.
type CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry_Cvpdnsessionattrdevicetype string

const (
    CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry_Cvpdnsessionattrdevicetype_other CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry_Cvpdnsessionattrdevicetype = "other"

    CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry_Cvpdnsessionattrdevicetype_asyncInternalModem CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry_Cvpdnsessionattrdevicetype = "asyncInternalModem"

    CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry_Cvpdnsessionattrdevicetype_async CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry_Cvpdnsessionattrdevicetype = "async"

    CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry_Cvpdnsessionattrdevicetype_bchan CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry_Cvpdnsessionattrdevicetype = "bchan"

    CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry_Cvpdnsessionattrdevicetype_sync CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry_Cvpdnsessionattrdevicetype = "sync"

    CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry_Cvpdnsessionattrdevicetype_virtualAccess CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry_Cvpdnsessionattrdevicetype = "virtualAccess"

    CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry_Cvpdnsessionattrdevicetype_xdsl CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry_Cvpdnsessionattrdevicetype = "xdsl"

    CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry_Cvpdnsessionattrdevicetype_cable CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry_Cvpdnsessionattrdevicetype = "cable"
)

// CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry_Cvpdnsessionattrstate represents     pptpWaitCallDisc:    Session shutdown is in progress.
type CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry_Cvpdnsessionattrstate string

const (
    CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry_Cvpdnsessionattrstate_unknown CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry_Cvpdnsessionattrstate = "unknown"

    CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry_Cvpdnsessionattrstate_l2fOpening CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry_Cvpdnsessionattrstate = "l2fOpening"

    CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry_Cvpdnsessionattrstate_l2fOpen CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry_Cvpdnsessionattrstate = "l2fOpen"

    CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry_Cvpdnsessionattrstate_l2fCloseWait CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry_Cvpdnsessionattrstate = "l2fCloseWait"

    CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry_Cvpdnsessionattrstate_l2fWaitingForTunnel CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry_Cvpdnsessionattrstate = "l2fWaitingForTunnel"

    CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry_Cvpdnsessionattrstate_l2tpIdle CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry_Cvpdnsessionattrstate = "l2tpIdle"

    CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry_Cvpdnsessionattrstate_l2tpWaitingTunnel CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry_Cvpdnsessionattrstate = "l2tpWaitingTunnel"

    CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry_Cvpdnsessionattrstate_l2tpWaitReply CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry_Cvpdnsessionattrstate = "l2tpWaitReply"

    CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry_Cvpdnsessionattrstate_l2tpWaitConnect CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry_Cvpdnsessionattrstate = "l2tpWaitConnect"

    CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry_Cvpdnsessionattrstate_l2tpEstablished CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry_Cvpdnsessionattrstate = "l2tpEstablished"

    CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry_Cvpdnsessionattrstate_l2tpShuttingDown CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry_Cvpdnsessionattrstate = "l2tpShuttingDown"

    CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry_Cvpdnsessionattrstate_pptpWaitVAccess CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry_Cvpdnsessionattrstate = "pptpWaitVAccess"

    CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry_Cvpdnsessionattrstate_pptpPacEstablished CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry_Cvpdnsessionattrstate = "pptpPacEstablished"

    CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry_Cvpdnsessionattrstate_pptpWaitTunnel CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry_Cvpdnsessionattrstate = "pptpWaitTunnel"

    CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry_Cvpdnsessionattrstate_pptpWaitOCRP CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry_Cvpdnsessionattrstate = "pptpWaitOCRP"

    CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry_Cvpdnsessionattrstate_pptpPnsEstablished CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry_Cvpdnsessionattrstate = "pptpPnsEstablished"

    CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry_Cvpdnsessionattrstate_pptpWaitCallDisc CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry_Cvpdnsessionattrstate = "pptpWaitCallDisc"

    CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry_Cvpdnsessionattrstate_pptpTerminal CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry_Cvpdnsessionattrstate = "pptpTerminal"
)

// CISCOVPDNMGMTMIB_Cvpdnusertofailhistinfotable
// Table of the record of failure objects which can be
// referenced by an user name.  Only a name that has a
// valid item in the Cisco IOS VPDN failure history table
// will yield a valid entry in this table.  The table has
// a maximum size of 50 entries.  Only the newest 50
// entries will be kept in the table.
type CISCOVPDNMGMTMIB_Cvpdnusertofailhistinfotable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table, containing failure history relevant to an user name.
    // The type is slice of
    // CISCOVPDNMGMTMIB_Cvpdnusertofailhistinfotable_Cvpdnusertofailhistinfoentry.
    Cvpdnusertofailhistinfoentry []CISCOVPDNMGMTMIB_Cvpdnusertofailhistinfotable_Cvpdnusertofailhistinfoentry
}

func (cvpdnusertofailhistinfotable *CISCOVPDNMGMTMIB_Cvpdnusertofailhistinfotable) GetEntityData() *types.CommonEntityData {
    cvpdnusertofailhistinfotable.EntityData.YFilter = cvpdnusertofailhistinfotable.YFilter
    cvpdnusertofailhistinfotable.EntityData.YangName = "cvpdnUserToFailHistInfoTable"
    cvpdnusertofailhistinfotable.EntityData.BundleName = "cisco_ios_xe"
    cvpdnusertofailhistinfotable.EntityData.ParentYangName = "CISCO-VPDN-MGMT-MIB"
    cvpdnusertofailhistinfotable.EntityData.SegmentPath = "cvpdnUserToFailHistInfoTable"
    cvpdnusertofailhistinfotable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvpdnusertofailhistinfotable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvpdnusertofailhistinfotable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvpdnusertofailhistinfotable.EntityData.Children = make(map[string]types.YChild)
    cvpdnusertofailhistinfotable.EntityData.Children["cvpdnUserToFailHistInfoEntry"] = types.YChild{"Cvpdnusertofailhistinfoentry", nil}
    for i := range cvpdnusertofailhistinfotable.Cvpdnusertofailhistinfoentry {
        cvpdnusertofailhistinfotable.EntityData.Children[types.GetSegmentPath(&cvpdnusertofailhistinfotable.Cvpdnusertofailhistinfoentry[i])] = types.YChild{"Cvpdnusertofailhistinfoentry", &cvpdnusertofailhistinfotable.Cvpdnusertofailhistinfoentry[i]}
    }
    cvpdnusertofailhistinfotable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cvpdnusertofailhistinfotable.EntityData)
}

// CISCOVPDNMGMTMIB_Cvpdnusertofailhistinfotable_Cvpdnusertofailhistinfoentry
// An entry in the table, containing failure history
// relevant to an user name.
type CISCOVPDNMGMTMIB_Cvpdnusertofailhistinfotable_Cvpdnusertofailhistinfoentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The user name for this failure entry. The type is
    // string with length: 1..255.
    Cvpdnunametofailhistuname interface{}

    // This attribute is a key. The Tunnel ID for this failure entry.  If it is
    // the instigator of the tunnel, the ID is the TS tunnel ID, otherwise it is
    // the NAS tunnel ID. The type is interface{} with range: 0..4294967295.
    Cvpdnunametofailhisttunnelid interface{}

    // The user ID of this failure entry. The type is interface{} with range:
    // 0..4294967295.
    Cvpdnunametofailhistuserid interface{}

    // This object indicates whether the tunnel in which the failure occurred was
    // generated locally or not. The type is bool.
    Cvpdnunametofailhistlocalinitconn interface{}

    // The local name of the VPDN tunnel in which the failure occurred.  It will
    // be the NAS name of the tunnel if the system serves as the NAS, or the TS
    // name of the tunnel if the system serves as the tunnel server.  The local
    // name is the configured host name of the system.  This object can be empty
    // if the failure occurred prior to successful tunnel projection, thus no
    // source name will be available. The type is string.
    Cvpdnunametofailhistlocalname interface{}

    // The remote name of the VPDN tunnel in which the failure occurred.  It will
    // be the tunnel server name of the tunnel if the system is a NAS, or the NAS
    // name of the tunnel if the system serves as the tunnel server.  This object
    // can be empty if the failure occurred prior to successful tunnel projection,
    // thus no source name will be available. The type is string.
    Cvpdnunametofailhistremotename interface{}

    // The source IP address of the tunnel in which the failure occurred.  This IP
    // address is that of the interface at the instigator end of the tunnel. The
    // type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Cvpdnunametofailhistsourceip interface{}

    // The destination IP address of the tunnel in which the failure occurred. 
    // This IP address is that of the interface at the receiver end of the tunnel.
    // The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Cvpdnunametofailhistdestip interface{}

    // This object is incremented when multiple failures has been experienced by
    // this user on this tunnel.  Seeing a delta of >1 is an indication that the
    // current failure record is the latest of a series of failures that has been
    // recorded. The type is interface{} with range: 0..4294967295. Units are
    // failures.
    Cvpdnunametofailhistcount interface{}

    // This object specifies the time when the failure is occurred. The type is
    // interface{} with range: 0..4294967295.
    Cvpdnunametofailhistfailtime interface{}

    // The type of failure for the current failure record.  It comes in a form of
    // a an ASCII string which describes the type of failure. The type is string
    // with length: 1..255.
    Cvpdnunametofailhistfailtype interface{}

    // The reason of failure for the current failure record. The type is string
    // with length: 1..255.
    Cvpdnunametofailhistfailreason interface{}

    // Indicates the type of address contained in
    // cvpdnUnameToFailHistSourceInetAddr. The type is InetAddressType.
    Cvpdnunametofailhistsourceinettype interface{}

    // The source IP address of the tunnel in which the failure occurred.  This IP
    // address is that of the interface at the instigator end of the tunnel.  The
    // instigator end is the peer which initiates the tunnel estalishment.  The
    // type of this address is determined by the value of
    // cvpdnUnameToFailHistSourceInetType. The type is string with length: 0..255.
    Cvpdnunametofailhistsourceinetaddr interface{}

    // Indicates the type of address contained in
    // cvpdnUnameToFailHistDestInetAddr. The type is InetAddressType.
    Cvpdnunametofailhistdestinettype interface{}

    // The destination IP address of the tunnel in which the failure occurred. 
    // This IP address is that of the interface at the receiver end of the tunnel.
    // The type  of this address is determined by the value of 
    // cvpdnUnameToFailHistDestInetType. The type is string with length: 0..255.
    Cvpdnunametofailhistdestinetaddr interface{}
}

func (cvpdnusertofailhistinfoentry *CISCOVPDNMGMTMIB_Cvpdnusertofailhistinfotable_Cvpdnusertofailhistinfoentry) GetEntityData() *types.CommonEntityData {
    cvpdnusertofailhistinfoentry.EntityData.YFilter = cvpdnusertofailhistinfoentry.YFilter
    cvpdnusertofailhistinfoentry.EntityData.YangName = "cvpdnUserToFailHistInfoEntry"
    cvpdnusertofailhistinfoentry.EntityData.BundleName = "cisco_ios_xe"
    cvpdnusertofailhistinfoentry.EntityData.ParentYangName = "cvpdnUserToFailHistInfoTable"
    cvpdnusertofailhistinfoentry.EntityData.SegmentPath = "cvpdnUserToFailHistInfoEntry" + "[cvpdnUnameToFailHistUname='" + fmt.Sprintf("%v", cvpdnusertofailhistinfoentry.Cvpdnunametofailhistuname) + "']" + "[cvpdnUnameToFailHistTunnelId='" + fmt.Sprintf("%v", cvpdnusertofailhistinfoentry.Cvpdnunametofailhisttunnelid) + "']"
    cvpdnusertofailhistinfoentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvpdnusertofailhistinfoentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvpdnusertofailhistinfoentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvpdnusertofailhistinfoentry.EntityData.Children = make(map[string]types.YChild)
    cvpdnusertofailhistinfoentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cvpdnusertofailhistinfoentry.EntityData.Leafs["cvpdnUnameToFailHistUname"] = types.YLeaf{"Cvpdnunametofailhistuname", cvpdnusertofailhistinfoentry.Cvpdnunametofailhistuname}
    cvpdnusertofailhistinfoentry.EntityData.Leafs["cvpdnUnameToFailHistTunnelId"] = types.YLeaf{"Cvpdnunametofailhisttunnelid", cvpdnusertofailhistinfoentry.Cvpdnunametofailhisttunnelid}
    cvpdnusertofailhistinfoentry.EntityData.Leafs["cvpdnUnameToFailHistUserId"] = types.YLeaf{"Cvpdnunametofailhistuserid", cvpdnusertofailhistinfoentry.Cvpdnunametofailhistuserid}
    cvpdnusertofailhistinfoentry.EntityData.Leafs["cvpdnUnameToFailHistLocalInitConn"] = types.YLeaf{"Cvpdnunametofailhistlocalinitconn", cvpdnusertofailhistinfoentry.Cvpdnunametofailhistlocalinitconn}
    cvpdnusertofailhistinfoentry.EntityData.Leafs["cvpdnUnameToFailHistLocalName"] = types.YLeaf{"Cvpdnunametofailhistlocalname", cvpdnusertofailhistinfoentry.Cvpdnunametofailhistlocalname}
    cvpdnusertofailhistinfoentry.EntityData.Leafs["cvpdnUnameToFailHistRemoteName"] = types.YLeaf{"Cvpdnunametofailhistremotename", cvpdnusertofailhistinfoentry.Cvpdnunametofailhistremotename}
    cvpdnusertofailhistinfoentry.EntityData.Leafs["cvpdnUnameToFailHistSourceIp"] = types.YLeaf{"Cvpdnunametofailhistsourceip", cvpdnusertofailhistinfoentry.Cvpdnunametofailhistsourceip}
    cvpdnusertofailhistinfoentry.EntityData.Leafs["cvpdnUnameToFailHistDestIp"] = types.YLeaf{"Cvpdnunametofailhistdestip", cvpdnusertofailhistinfoentry.Cvpdnunametofailhistdestip}
    cvpdnusertofailhistinfoentry.EntityData.Leafs["cvpdnUnameToFailHistCount"] = types.YLeaf{"Cvpdnunametofailhistcount", cvpdnusertofailhistinfoentry.Cvpdnunametofailhistcount}
    cvpdnusertofailhistinfoentry.EntityData.Leafs["cvpdnUnameToFailHistFailTime"] = types.YLeaf{"Cvpdnunametofailhistfailtime", cvpdnusertofailhistinfoentry.Cvpdnunametofailhistfailtime}
    cvpdnusertofailhistinfoentry.EntityData.Leafs["cvpdnUnameToFailHistFailType"] = types.YLeaf{"Cvpdnunametofailhistfailtype", cvpdnusertofailhistinfoentry.Cvpdnunametofailhistfailtype}
    cvpdnusertofailhistinfoentry.EntityData.Leafs["cvpdnUnameToFailHistFailReason"] = types.YLeaf{"Cvpdnunametofailhistfailreason", cvpdnusertofailhistinfoentry.Cvpdnunametofailhistfailreason}
    cvpdnusertofailhistinfoentry.EntityData.Leafs["cvpdnUnameToFailHistSourceInetType"] = types.YLeaf{"Cvpdnunametofailhistsourceinettype", cvpdnusertofailhistinfoentry.Cvpdnunametofailhistsourceinettype}
    cvpdnusertofailhistinfoentry.EntityData.Leafs["cvpdnUnameToFailHistSourceInetAddr"] = types.YLeaf{"Cvpdnunametofailhistsourceinetaddr", cvpdnusertofailhistinfoentry.Cvpdnunametofailhistsourceinetaddr}
    cvpdnusertofailhistinfoentry.EntityData.Leafs["cvpdnUnameToFailHistDestInetType"] = types.YLeaf{"Cvpdnunametofailhistdestinettype", cvpdnusertofailhistinfoentry.Cvpdnunametofailhistdestinettype}
    cvpdnusertofailhistinfoentry.EntityData.Leafs["cvpdnUnameToFailHistDestInetAddr"] = types.YLeaf{"Cvpdnunametofailhistdestinetaddr", cvpdnusertofailhistinfoentry.Cvpdnunametofailhistdestinetaddr}
    return &(cvpdnusertofailhistinfoentry.EntityData)
}

// CISCOVPDNMGMTMIB_Cvpdntemplatetable
// Table of information about the VPDN templates.  The
// VPDN template is a grouping mechanism that allows
// configuration settings to be shared among multiple VPDN
// groups.  One such setting is a limit on the number of
// active sessions across all VPDN groups associated with
// the template.  The template table allows customers to
// monitor template-wide information such as tracking the
// allocation of sessions across templates.
type CISCOVPDNMGMTMIB_Cvpdntemplatetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table, containing information about a single VPDN template.
    // The type is slice of
    // CISCOVPDNMGMTMIB_Cvpdntemplatetable_Cvpdntemplateentry.
    Cvpdntemplateentry []CISCOVPDNMGMTMIB_Cvpdntemplatetable_Cvpdntemplateentry
}

func (cvpdntemplatetable *CISCOVPDNMGMTMIB_Cvpdntemplatetable) GetEntityData() *types.CommonEntityData {
    cvpdntemplatetable.EntityData.YFilter = cvpdntemplatetable.YFilter
    cvpdntemplatetable.EntityData.YangName = "cvpdnTemplateTable"
    cvpdntemplatetable.EntityData.BundleName = "cisco_ios_xe"
    cvpdntemplatetable.EntityData.ParentYangName = "CISCO-VPDN-MGMT-MIB"
    cvpdntemplatetable.EntityData.SegmentPath = "cvpdnTemplateTable"
    cvpdntemplatetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvpdntemplatetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvpdntemplatetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvpdntemplatetable.EntityData.Children = make(map[string]types.YChild)
    cvpdntemplatetable.EntityData.Children["cvpdnTemplateEntry"] = types.YChild{"Cvpdntemplateentry", nil}
    for i := range cvpdntemplatetable.Cvpdntemplateentry {
        cvpdntemplatetable.EntityData.Children[types.GetSegmentPath(&cvpdntemplatetable.Cvpdntemplateentry[i])] = types.YChild{"Cvpdntemplateentry", &cvpdntemplatetable.Cvpdntemplateentry[i]}
    }
    cvpdntemplatetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cvpdntemplatetable.EntityData)
}

// CISCOVPDNMGMTMIB_Cvpdntemplatetable_Cvpdntemplateentry
// An entry in the table, containing information about a
// single VPDN template.
type CISCOVPDNMGMTMIB_Cvpdntemplatetable_Cvpdntemplateentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the VPDN template. The type is string
    // with length: 1..255.
    Cvpdntemplatename interface{}

    // The total number of active session in all groups associated with the
    // template. The type is interface{} with range: 0..4294967295. Units are
    // sessions.
    Cvpdntemplateactivesessions interface{}
}

func (cvpdntemplateentry *CISCOVPDNMGMTMIB_Cvpdntemplatetable_Cvpdntemplateentry) GetEntityData() *types.CommonEntityData {
    cvpdntemplateentry.EntityData.YFilter = cvpdntemplateentry.YFilter
    cvpdntemplateentry.EntityData.YangName = "cvpdnTemplateEntry"
    cvpdntemplateentry.EntityData.BundleName = "cisco_ios_xe"
    cvpdntemplateentry.EntityData.ParentYangName = "cvpdnTemplateTable"
    cvpdntemplateentry.EntityData.SegmentPath = "cvpdnTemplateEntry" + "[cvpdnTemplateName='" + fmt.Sprintf("%v", cvpdntemplateentry.Cvpdntemplatename) + "']"
    cvpdntemplateentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvpdntemplateentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvpdntemplateentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvpdntemplateentry.EntityData.Children = make(map[string]types.YChild)
    cvpdntemplateentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cvpdntemplateentry.EntityData.Leafs["cvpdnTemplateName"] = types.YLeaf{"Cvpdntemplatename", cvpdntemplateentry.Cvpdntemplatename}
    cvpdntemplateentry.EntityData.Leafs["cvpdnTemplateActiveSessions"] = types.YLeaf{"Cvpdntemplateactivesessions", cvpdntemplateentry.Cvpdntemplateactivesessions}
    return &(cvpdntemplateentry.EntityData)
}

// CISCOVPDNMGMTMIB_Cvpdnbundletable
// Table that describes the multilink PPP attributes of the
// active VPDN sessions.
type CISCOVPDNMGMTMIB_Cvpdnbundletable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table represents an active multilink PPP bundle that
    // belongs to a VPDN tunnel. The type is slice of
    // CISCOVPDNMGMTMIB_Cvpdnbundletable_Cvpdnbundleentry.
    Cvpdnbundleentry []CISCOVPDNMGMTMIB_Cvpdnbundletable_Cvpdnbundleentry
}

func (cvpdnbundletable *CISCOVPDNMGMTMIB_Cvpdnbundletable) GetEntityData() *types.CommonEntityData {
    cvpdnbundletable.EntityData.YFilter = cvpdnbundletable.YFilter
    cvpdnbundletable.EntityData.YangName = "cvpdnBundleTable"
    cvpdnbundletable.EntityData.BundleName = "cisco_ios_xe"
    cvpdnbundletable.EntityData.ParentYangName = "CISCO-VPDN-MGMT-MIB"
    cvpdnbundletable.EntityData.SegmentPath = "cvpdnBundleTable"
    cvpdnbundletable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvpdnbundletable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvpdnbundletable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvpdnbundletable.EntityData.Children = make(map[string]types.YChild)
    cvpdnbundletable.EntityData.Children["cvpdnBundleEntry"] = types.YChild{"Cvpdnbundleentry", nil}
    for i := range cvpdnbundletable.Cvpdnbundleentry {
        cvpdnbundletable.EntityData.Children[types.GetSegmentPath(&cvpdnbundletable.Cvpdnbundleentry[i])] = types.YChild{"Cvpdnbundleentry", &cvpdnbundletable.Cvpdnbundleentry[i]}
    }
    cvpdnbundletable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cvpdnbundletable.EntityData)
}

// CISCOVPDNMGMTMIB_Cvpdnbundletable_Cvpdnbundleentry
// An entry in this table represents an active multilink PPP
// bundle that belongs to a VPDN tunnel.
type CISCOVPDNMGMTMIB_Cvpdnbundletable_Cvpdnbundleentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the multilink PPP bundle associated
    // with a VPDN tunnel. The type is string with length: 1..64.
    Cvpdnbundlename interface{}

    // The total number of active links in a multilink PPP bundle associated with
    // a VPDN tunnel. The type is interface{} with range: 0..4294967295. Units are
    // links.
    Cvpdnbundlelinkcount interface{}

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
    // discriminator. The type is Cvpdnbundleendpointtype.
    Cvpdnbundleendpointtype interface{}

    // Indicates the discriminator used in this bundle that is associated with a
    // VPDN tunnel. The type is string with length: 0..255.
    Cvpdnbundleendpoint interface{}

    // Indicates the type of address contained in cvpdnBundlePeerIpAddr. The type
    // is InetAddressType.
    Cvpdnbundlepeeripaddrtype interface{}

    // The IP address of the multilink PPP peer in a VPDN tunnel. The type is
    // string with length: 0..255.
    Cvpdnbundlepeeripaddr interface{}

    // The multilink PPP bundle discriminator class associated with a VPDN tunnel.
    // The value of this object represents the type of discriminator used in
    // cvpdnBundleEndpoint. The type is EndpointClass.
    Cvpdnbundleendpointclass interface{}
}

func (cvpdnbundleentry *CISCOVPDNMGMTMIB_Cvpdnbundletable_Cvpdnbundleentry) GetEntityData() *types.CommonEntityData {
    cvpdnbundleentry.EntityData.YFilter = cvpdnbundleentry.YFilter
    cvpdnbundleentry.EntityData.YangName = "cvpdnBundleEntry"
    cvpdnbundleentry.EntityData.BundleName = "cisco_ios_xe"
    cvpdnbundleentry.EntityData.ParentYangName = "cvpdnBundleTable"
    cvpdnbundleentry.EntityData.SegmentPath = "cvpdnBundleEntry" + "[cvpdnBundleName='" + fmt.Sprintf("%v", cvpdnbundleentry.Cvpdnbundlename) + "']"
    cvpdnbundleentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvpdnbundleentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvpdnbundleentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvpdnbundleentry.EntityData.Children = make(map[string]types.YChild)
    cvpdnbundleentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cvpdnbundleentry.EntityData.Leafs["cvpdnBundleName"] = types.YLeaf{"Cvpdnbundlename", cvpdnbundleentry.Cvpdnbundlename}
    cvpdnbundleentry.EntityData.Leafs["cvpdnBundleLinkCount"] = types.YLeaf{"Cvpdnbundlelinkcount", cvpdnbundleentry.Cvpdnbundlelinkcount}
    cvpdnbundleentry.EntityData.Leafs["cvpdnBundleEndpointType"] = types.YLeaf{"Cvpdnbundleendpointtype", cvpdnbundleentry.Cvpdnbundleendpointtype}
    cvpdnbundleentry.EntityData.Leafs["cvpdnBundleEndpoint"] = types.YLeaf{"Cvpdnbundleendpoint", cvpdnbundleentry.Cvpdnbundleendpoint}
    cvpdnbundleentry.EntityData.Leafs["cvpdnBundlePeerIpAddrType"] = types.YLeaf{"Cvpdnbundlepeeripaddrtype", cvpdnbundleentry.Cvpdnbundlepeeripaddrtype}
    cvpdnbundleentry.EntityData.Leafs["cvpdnBundlePeerIpAddr"] = types.YLeaf{"Cvpdnbundlepeeripaddr", cvpdnbundleentry.Cvpdnbundlepeeripaddr}
    cvpdnbundleentry.EntityData.Leafs["cvpdnBundleEndpointClass"] = types.YLeaf{"Cvpdnbundleendpointclass", cvpdnbundleentry.Cvpdnbundleendpointclass}
    return &(cvpdnbundleentry.EntityData)
}

// CISCOVPDNMGMTMIB_Cvpdnbundletable_Cvpdnbundleentry_Cvpdnbundleendpointtype represents                  discriminator.
type CISCOVPDNMGMTMIB_Cvpdnbundletable_Cvpdnbundleentry_Cvpdnbundleendpointtype string

const (
    CISCOVPDNMGMTMIB_Cvpdnbundletable_Cvpdnbundleentry_Cvpdnbundleendpointtype_none CISCOVPDNMGMTMIB_Cvpdnbundletable_Cvpdnbundleentry_Cvpdnbundleendpointtype = "none"

    CISCOVPDNMGMTMIB_Cvpdnbundletable_Cvpdnbundleentry_Cvpdnbundleendpointtype_hostname CISCOVPDNMGMTMIB_Cvpdnbundletable_Cvpdnbundleentry_Cvpdnbundleendpointtype = "hostname"

    CISCOVPDNMGMTMIB_Cvpdnbundletable_Cvpdnbundleentry_Cvpdnbundleendpointtype_string_ CISCOVPDNMGMTMIB_Cvpdnbundletable_Cvpdnbundleentry_Cvpdnbundleendpointtype = "string"

    CISCOVPDNMGMTMIB_Cvpdnbundletable_Cvpdnbundleentry_Cvpdnbundleendpointtype_macAddress CISCOVPDNMGMTMIB_Cvpdnbundletable_Cvpdnbundleentry_Cvpdnbundleendpointtype = "macAddress"

    CISCOVPDNMGMTMIB_Cvpdnbundletable_Cvpdnbundleentry_Cvpdnbundleendpointtype_ipV4Address CISCOVPDNMGMTMIB_Cvpdnbundletable_Cvpdnbundleentry_Cvpdnbundleendpointtype = "ipV4Address"

    CISCOVPDNMGMTMIB_Cvpdnbundletable_Cvpdnbundleentry_Cvpdnbundleendpointtype_ipV6Address CISCOVPDNMGMTMIB_Cvpdnbundletable_Cvpdnbundleentry_Cvpdnbundleendpointtype = "ipV6Address"

    CISCOVPDNMGMTMIB_Cvpdnbundletable_Cvpdnbundleentry_Cvpdnbundleendpointtype_phone CISCOVPDNMGMTMIB_Cvpdnbundletable_Cvpdnbundleentry_Cvpdnbundleendpointtype = "phone"

    CISCOVPDNMGMTMIB_Cvpdnbundletable_Cvpdnbundleentry_Cvpdnbundleendpointtype_magicNumber CISCOVPDNMGMTMIB_Cvpdnbundletable_Cvpdnbundleentry_Cvpdnbundleendpointtype = "magicNumber"
)

// CISCOVPDNMGMTMIB_Cvpdnbundlechildtable
// A table that exposes the containment relationship between a
// multilink PPP bundle and a VPDN tunnel.
type CISCOVPDNMGMTMIB_Cvpdnbundlechildtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table represents a session that belongs to a VPDN tunnel
    // and to a multilink PPP bundle. The type is slice of
    // CISCOVPDNMGMTMIB_Cvpdnbundlechildtable_Cvpdnbundlechildentry.
    Cvpdnbundlechildentry []CISCOVPDNMGMTMIB_Cvpdnbundlechildtable_Cvpdnbundlechildentry
}

func (cvpdnbundlechildtable *CISCOVPDNMGMTMIB_Cvpdnbundlechildtable) GetEntityData() *types.CommonEntityData {
    cvpdnbundlechildtable.EntityData.YFilter = cvpdnbundlechildtable.YFilter
    cvpdnbundlechildtable.EntityData.YangName = "cvpdnBundleChildTable"
    cvpdnbundlechildtable.EntityData.BundleName = "cisco_ios_xe"
    cvpdnbundlechildtable.EntityData.ParentYangName = "CISCO-VPDN-MGMT-MIB"
    cvpdnbundlechildtable.EntityData.SegmentPath = "cvpdnBundleChildTable"
    cvpdnbundlechildtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvpdnbundlechildtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvpdnbundlechildtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvpdnbundlechildtable.EntityData.Children = make(map[string]types.YChild)
    cvpdnbundlechildtable.EntityData.Children["cvpdnBundleChildEntry"] = types.YChild{"Cvpdnbundlechildentry", nil}
    for i := range cvpdnbundlechildtable.Cvpdnbundlechildentry {
        cvpdnbundlechildtable.EntityData.Children[types.GetSegmentPath(&cvpdnbundlechildtable.Cvpdnbundlechildentry[i])] = types.YChild{"Cvpdnbundlechildentry", &cvpdnbundlechildtable.Cvpdnbundlechildentry[i]}
    }
    cvpdnbundlechildtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cvpdnbundlechildtable.EntityData)
}

// CISCOVPDNMGMTMIB_Cvpdnbundlechildtable_Cvpdnbundlechildentry
// An entry in this table represents a session that belongs to
// a VPDN tunnel and to a multilink PPP bundle.
type CISCOVPDNMGMTMIB_Cvpdnbundlechildtable_Cvpdnbundlechildentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with length: 1..64. Refers to
    // cisco_vpdn_mgmt_mib.CISCOVPDNMGMTMIB_Cvpdnbundletable_Cvpdnbundleentry_Cvpdnbundlename
    Cvpdnbundlename interface{}

    // This attribute is a key. The tunnel type.  This is the tunnel protocol of
    // an active VPDN session that is associated with a multilink PPP bundle. The
    // type is TunnelType.
    Cvpdnbundlechildtunneltype interface{}

    // This attribute is a key. The Tunnel ID of an active VPDN session that is
    // associated with a multilink PPP bundle. The type is interface{} with range:
    // 0..4294967295.
    Cvpdnbundlechildtunnelid interface{}

    // This attribute is a key. The ID of an active VPDN session that is
    // associated with a multilink PPP bundle. The type is interface{} with range:
    // 0..4294967295.
    Cvpdnbundlechildsessionid interface{}
}

func (cvpdnbundlechildentry *CISCOVPDNMGMTMIB_Cvpdnbundlechildtable_Cvpdnbundlechildentry) GetEntityData() *types.CommonEntityData {
    cvpdnbundlechildentry.EntityData.YFilter = cvpdnbundlechildentry.YFilter
    cvpdnbundlechildentry.EntityData.YangName = "cvpdnBundleChildEntry"
    cvpdnbundlechildentry.EntityData.BundleName = "cisco_ios_xe"
    cvpdnbundlechildentry.EntityData.ParentYangName = "cvpdnBundleChildTable"
    cvpdnbundlechildentry.EntityData.SegmentPath = "cvpdnBundleChildEntry" + "[cvpdnBundleName='" + fmt.Sprintf("%v", cvpdnbundlechildentry.Cvpdnbundlename) + "']" + "[cvpdnBundleChildTunnelType='" + fmt.Sprintf("%v", cvpdnbundlechildentry.Cvpdnbundlechildtunneltype) + "']" + "[cvpdnBundleChildTunnelId='" + fmt.Sprintf("%v", cvpdnbundlechildentry.Cvpdnbundlechildtunnelid) + "']" + "[cvpdnBundleChildSessionId='" + fmt.Sprintf("%v", cvpdnbundlechildentry.Cvpdnbundlechildsessionid) + "']"
    cvpdnbundlechildentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvpdnbundlechildentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvpdnbundlechildentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvpdnbundlechildentry.EntityData.Children = make(map[string]types.YChild)
    cvpdnbundlechildentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cvpdnbundlechildentry.EntityData.Leafs["cvpdnBundleName"] = types.YLeaf{"Cvpdnbundlename", cvpdnbundlechildentry.Cvpdnbundlename}
    cvpdnbundlechildentry.EntityData.Leafs["cvpdnBundleChildTunnelType"] = types.YLeaf{"Cvpdnbundlechildtunneltype", cvpdnbundlechildentry.Cvpdnbundlechildtunneltype}
    cvpdnbundlechildentry.EntityData.Leafs["cvpdnBundleChildTunnelId"] = types.YLeaf{"Cvpdnbundlechildtunnelid", cvpdnbundlechildentry.Cvpdnbundlechildtunnelid}
    cvpdnbundlechildentry.EntityData.Leafs["cvpdnBundleChildSessionId"] = types.YLeaf{"Cvpdnbundlechildsessionid", cvpdnbundlechildentry.Cvpdnbundlechildsessionid}
    return &(cvpdnbundlechildentry.EntityData)
}

