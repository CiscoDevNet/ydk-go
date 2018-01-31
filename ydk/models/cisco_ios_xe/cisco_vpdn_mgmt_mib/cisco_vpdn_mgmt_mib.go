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
    parent types.Entity
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

func (cISCOVPDNMGMTMIB *CISCOVPDNMGMTMIB) GetFilter() yfilter.YFilter { return cISCOVPDNMGMTMIB.YFilter }

func (cISCOVPDNMGMTMIB *CISCOVPDNMGMTMIB) SetFilter(yf yfilter.YFilter) { cISCOVPDNMGMTMIB.YFilter = yf }

func (cISCOVPDNMGMTMIB *CISCOVPDNMGMTMIB) GetGoName(yname string) string {
    if yname == "ciscoVpdnMgmtMIBNotifs" { return "Ciscovpdnmgmtmibnotifs" }
    if yname == "cvpdnSystemInfo" { return "Cvpdnsysteminfo" }
    if yname == "cvpdnMultilinkInfo" { return "Cvpdnmultilinkinfo" }
    if yname == "cvpdnSystemTable" { return "Cvpdnsystemtable" }
    if yname == "cvpdnTunnelTable" { return "Cvpdntunneltable" }
    if yname == "cvpdnTunnelAttrTable" { return "Cvpdntunnelattrtable" }
    if yname == "cvpdnTunnelSessionTable" { return "Cvpdntunnelsessiontable" }
    if yname == "cvpdnSessionAttrTable" { return "Cvpdnsessionattrtable" }
    if yname == "cvpdnUserToFailHistInfoTable" { return "Cvpdnusertofailhistinfotable" }
    if yname == "cvpdnTemplateTable" { return "Cvpdntemplatetable" }
    if yname == "cvpdnBundleTable" { return "Cvpdnbundletable" }
    if yname == "cvpdnBundleChildTable" { return "Cvpdnbundlechildtable" }
    return ""
}

func (cISCOVPDNMGMTMIB *CISCOVPDNMGMTMIB) GetSegmentPath() string {
    return "CISCO-VPDN-MGMT-MIB:CISCO-VPDN-MGMT-MIB"
}

func (cISCOVPDNMGMTMIB *CISCOVPDNMGMTMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ciscoVpdnMgmtMIBNotifs" {
        return &cISCOVPDNMGMTMIB.Ciscovpdnmgmtmibnotifs
    }
    if childYangName == "cvpdnSystemInfo" {
        return &cISCOVPDNMGMTMIB.Cvpdnsysteminfo
    }
    if childYangName == "cvpdnMultilinkInfo" {
        return &cISCOVPDNMGMTMIB.Cvpdnmultilinkinfo
    }
    if childYangName == "cvpdnSystemTable" {
        return &cISCOVPDNMGMTMIB.Cvpdnsystemtable
    }
    if childYangName == "cvpdnTunnelTable" {
        return &cISCOVPDNMGMTMIB.Cvpdntunneltable
    }
    if childYangName == "cvpdnTunnelAttrTable" {
        return &cISCOVPDNMGMTMIB.Cvpdntunnelattrtable
    }
    if childYangName == "cvpdnTunnelSessionTable" {
        return &cISCOVPDNMGMTMIB.Cvpdntunnelsessiontable
    }
    if childYangName == "cvpdnSessionAttrTable" {
        return &cISCOVPDNMGMTMIB.Cvpdnsessionattrtable
    }
    if childYangName == "cvpdnUserToFailHistInfoTable" {
        return &cISCOVPDNMGMTMIB.Cvpdnusertofailhistinfotable
    }
    if childYangName == "cvpdnTemplateTable" {
        return &cISCOVPDNMGMTMIB.Cvpdntemplatetable
    }
    if childYangName == "cvpdnBundleTable" {
        return &cISCOVPDNMGMTMIB.Cvpdnbundletable
    }
    if childYangName == "cvpdnBundleChildTable" {
        return &cISCOVPDNMGMTMIB.Cvpdnbundlechildtable
    }
    return nil
}

func (cISCOVPDNMGMTMIB *CISCOVPDNMGMTMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ciscoVpdnMgmtMIBNotifs"] = &cISCOVPDNMGMTMIB.Ciscovpdnmgmtmibnotifs
    children["cvpdnSystemInfo"] = &cISCOVPDNMGMTMIB.Cvpdnsysteminfo
    children["cvpdnMultilinkInfo"] = &cISCOVPDNMGMTMIB.Cvpdnmultilinkinfo
    children["cvpdnSystemTable"] = &cISCOVPDNMGMTMIB.Cvpdnsystemtable
    children["cvpdnTunnelTable"] = &cISCOVPDNMGMTMIB.Cvpdntunneltable
    children["cvpdnTunnelAttrTable"] = &cISCOVPDNMGMTMIB.Cvpdntunnelattrtable
    children["cvpdnTunnelSessionTable"] = &cISCOVPDNMGMTMIB.Cvpdntunnelsessiontable
    children["cvpdnSessionAttrTable"] = &cISCOVPDNMGMTMIB.Cvpdnsessionattrtable
    children["cvpdnUserToFailHistInfoTable"] = &cISCOVPDNMGMTMIB.Cvpdnusertofailhistinfotable
    children["cvpdnTemplateTable"] = &cISCOVPDNMGMTMIB.Cvpdntemplatetable
    children["cvpdnBundleTable"] = &cISCOVPDNMGMTMIB.Cvpdnbundletable
    children["cvpdnBundleChildTable"] = &cISCOVPDNMGMTMIB.Cvpdnbundlechildtable
    return children
}

func (cISCOVPDNMGMTMIB *CISCOVPDNMGMTMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOVPDNMGMTMIB *CISCOVPDNMGMTMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOVPDNMGMTMIB *CISCOVPDNMGMTMIB) GetYangName() string { return "CISCO-VPDN-MGMT-MIB" }

func (cISCOVPDNMGMTMIB *CISCOVPDNMGMTMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOVPDNMGMTMIB *CISCOVPDNMGMTMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOVPDNMGMTMIB *CISCOVPDNMGMTMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOVPDNMGMTMIB *CISCOVPDNMGMTMIB) SetParent(parent types.Entity) { cISCOVPDNMGMTMIB.parent = parent }

func (cISCOVPDNMGMTMIB *CISCOVPDNMGMTMIB) GetParent() types.Entity { return cISCOVPDNMGMTMIB.parent }

func (cISCOVPDNMGMTMIB *CISCOVPDNMGMTMIB) GetParentYangName() string { return "CISCO-VPDN-MGMT-MIB" }

// CISCOVPDNMGMTMIB_Ciscovpdnmgmtmibnotifs
type CISCOVPDNMGMTMIB_Ciscovpdnmgmtmibnotifs struct {
    parent types.Entity
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

func (ciscovpdnmgmtmibnotifs *CISCOVPDNMGMTMIB_Ciscovpdnmgmtmibnotifs) GetFilter() yfilter.YFilter { return ciscovpdnmgmtmibnotifs.YFilter }

func (ciscovpdnmgmtmibnotifs *CISCOVPDNMGMTMIB_Ciscovpdnmgmtmibnotifs) SetFilter(yf yfilter.YFilter) { ciscovpdnmgmtmibnotifs.YFilter = yf }

func (ciscovpdnmgmtmibnotifs *CISCOVPDNMGMTMIB_Ciscovpdnmgmtmibnotifs) GetGoName(yname string) string {
    if yname == "cvpdnNotifSessionID" { return "Cvpdnnotifsessionid" }
    if yname == "cvpdnNotifSessionEvent" { return "Cvpdnnotifsessionevent" }
    return ""
}

func (ciscovpdnmgmtmibnotifs *CISCOVPDNMGMTMIB_Ciscovpdnmgmtmibnotifs) GetSegmentPath() string {
    return "ciscoVpdnMgmtMIBNotifs"
}

func (ciscovpdnmgmtmibnotifs *CISCOVPDNMGMTMIB_Ciscovpdnmgmtmibnotifs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ciscovpdnmgmtmibnotifs *CISCOVPDNMGMTMIB_Ciscovpdnmgmtmibnotifs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ciscovpdnmgmtmibnotifs *CISCOVPDNMGMTMIB_Ciscovpdnmgmtmibnotifs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cvpdnNotifSessionID"] = ciscovpdnmgmtmibnotifs.Cvpdnnotifsessionid
    leafs["cvpdnNotifSessionEvent"] = ciscovpdnmgmtmibnotifs.Cvpdnnotifsessionevent
    return leafs
}

func (ciscovpdnmgmtmibnotifs *CISCOVPDNMGMTMIB_Ciscovpdnmgmtmibnotifs) GetBundleName() string { return "cisco_ios_xe" }

func (ciscovpdnmgmtmibnotifs *CISCOVPDNMGMTMIB_Ciscovpdnmgmtmibnotifs) GetYangName() string { return "ciscoVpdnMgmtMIBNotifs" }

func (ciscovpdnmgmtmibnotifs *CISCOVPDNMGMTMIB_Ciscovpdnmgmtmibnotifs) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciscovpdnmgmtmibnotifs *CISCOVPDNMGMTMIB_Ciscovpdnmgmtmibnotifs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciscovpdnmgmtmibnotifs *CISCOVPDNMGMTMIB_Ciscovpdnmgmtmibnotifs) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciscovpdnmgmtmibnotifs *CISCOVPDNMGMTMIB_Ciscovpdnmgmtmibnotifs) SetParent(parent types.Entity) { ciscovpdnmgmtmibnotifs.parent = parent }

func (ciscovpdnmgmtmibnotifs *CISCOVPDNMGMTMIB_Ciscovpdnmgmtmibnotifs) GetParent() types.Entity { return ciscovpdnmgmtmibnotifs.parent }

func (ciscovpdnmgmtmibnotifs *CISCOVPDNMGMTMIB_Ciscovpdnmgmtmibnotifs) GetParentYangName() string { return "CISCO-VPDN-MGMT-MIB" }

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
    parent types.Entity
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

func (cvpdnsysteminfo *CISCOVPDNMGMTMIB_Cvpdnsysteminfo) GetFilter() yfilter.YFilter { return cvpdnsysteminfo.YFilter }

func (cvpdnsysteminfo *CISCOVPDNMGMTMIB_Cvpdnsysteminfo) SetFilter(yf yfilter.YFilter) { cvpdnsysteminfo.YFilter = yf }

func (cvpdnsysteminfo *CISCOVPDNMGMTMIB_Cvpdnsysteminfo) GetGoName(yname string) string {
    if yname == "cvpdnTunnelTotal" { return "Cvpdntunneltotal" }
    if yname == "cvpdnSessionTotal" { return "Cvpdnsessiontotal" }
    if yname == "cvpdnDeniedUsersTotal" { return "Cvpdndenieduserstotal" }
    if yname == "cvpdnSystemNotifSessionEnabled" { return "Cvpdnsystemnotifsessionenabled" }
    if yname == "cvpdnSystemClearSessions" { return "Cvpdnsystemclearsessions" }
    return ""
}

func (cvpdnsysteminfo *CISCOVPDNMGMTMIB_Cvpdnsysteminfo) GetSegmentPath() string {
    return "cvpdnSystemInfo"
}

func (cvpdnsysteminfo *CISCOVPDNMGMTMIB_Cvpdnsysteminfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cvpdnsysteminfo *CISCOVPDNMGMTMIB_Cvpdnsysteminfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cvpdnsysteminfo *CISCOVPDNMGMTMIB_Cvpdnsysteminfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cvpdnTunnelTotal"] = cvpdnsysteminfo.Cvpdntunneltotal
    leafs["cvpdnSessionTotal"] = cvpdnsysteminfo.Cvpdnsessiontotal
    leafs["cvpdnDeniedUsersTotal"] = cvpdnsysteminfo.Cvpdndenieduserstotal
    leafs["cvpdnSystemNotifSessionEnabled"] = cvpdnsysteminfo.Cvpdnsystemnotifsessionenabled
    leafs["cvpdnSystemClearSessions"] = cvpdnsysteminfo.Cvpdnsystemclearsessions
    return leafs
}

func (cvpdnsysteminfo *CISCOVPDNMGMTMIB_Cvpdnsysteminfo) GetBundleName() string { return "cisco_ios_xe" }

func (cvpdnsysteminfo *CISCOVPDNMGMTMIB_Cvpdnsysteminfo) GetYangName() string { return "cvpdnSystemInfo" }

func (cvpdnsysteminfo *CISCOVPDNMGMTMIB_Cvpdnsysteminfo) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvpdnsysteminfo *CISCOVPDNMGMTMIB_Cvpdnsysteminfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvpdnsysteminfo *CISCOVPDNMGMTMIB_Cvpdnsysteminfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvpdnsysteminfo *CISCOVPDNMGMTMIB_Cvpdnsysteminfo) SetParent(parent types.Entity) { cvpdnsysteminfo.parent = parent }

func (cvpdnsysteminfo *CISCOVPDNMGMTMIB_Cvpdnsysteminfo) GetParent() types.Entity { return cvpdnsysteminfo.parent }

func (cvpdnsysteminfo *CISCOVPDNMGMTMIB_Cvpdnsysteminfo) GetParentYangName() string { return "CISCO-VPDN-MGMT-MIB" }

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
    parent types.Entity
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

func (cvpdnmultilinkinfo *CISCOVPDNMGMTMIB_Cvpdnmultilinkinfo) GetFilter() yfilter.YFilter { return cvpdnmultilinkinfo.YFilter }

func (cvpdnmultilinkinfo *CISCOVPDNMGMTMIB_Cvpdnmultilinkinfo) SetFilter(yf yfilter.YFilter) { cvpdnmultilinkinfo.YFilter = yf }

func (cvpdnmultilinkinfo *CISCOVPDNMGMTMIB_Cvpdnmultilinkinfo) GetGoName(yname string) string {
    if yname == "cvpdnBundlesWithOneLink" { return "Cvpdnbundleswithonelink" }
    if yname == "cvpdnBundlesWithTwoLinks" { return "Cvpdnbundleswithtwolinks" }
    if yname == "cvpdnBundlesWithMoreThanTwoLinks" { return "Cvpdnbundleswithmorethantwolinks" }
    if yname == "cvpdnBundleLastChanged" { return "Cvpdnbundlelastchanged" }
    return ""
}

func (cvpdnmultilinkinfo *CISCOVPDNMGMTMIB_Cvpdnmultilinkinfo) GetSegmentPath() string {
    return "cvpdnMultilinkInfo"
}

func (cvpdnmultilinkinfo *CISCOVPDNMGMTMIB_Cvpdnmultilinkinfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cvpdnmultilinkinfo *CISCOVPDNMGMTMIB_Cvpdnmultilinkinfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cvpdnmultilinkinfo *CISCOVPDNMGMTMIB_Cvpdnmultilinkinfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cvpdnBundlesWithOneLink"] = cvpdnmultilinkinfo.Cvpdnbundleswithonelink
    leafs["cvpdnBundlesWithTwoLinks"] = cvpdnmultilinkinfo.Cvpdnbundleswithtwolinks
    leafs["cvpdnBundlesWithMoreThanTwoLinks"] = cvpdnmultilinkinfo.Cvpdnbundleswithmorethantwolinks
    leafs["cvpdnBundleLastChanged"] = cvpdnmultilinkinfo.Cvpdnbundlelastchanged
    return leafs
}

func (cvpdnmultilinkinfo *CISCOVPDNMGMTMIB_Cvpdnmultilinkinfo) GetBundleName() string { return "cisco_ios_xe" }

func (cvpdnmultilinkinfo *CISCOVPDNMGMTMIB_Cvpdnmultilinkinfo) GetYangName() string { return "cvpdnMultilinkInfo" }

func (cvpdnmultilinkinfo *CISCOVPDNMGMTMIB_Cvpdnmultilinkinfo) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvpdnmultilinkinfo *CISCOVPDNMGMTMIB_Cvpdnmultilinkinfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvpdnmultilinkinfo *CISCOVPDNMGMTMIB_Cvpdnmultilinkinfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvpdnmultilinkinfo *CISCOVPDNMGMTMIB_Cvpdnmultilinkinfo) SetParent(parent types.Entity) { cvpdnmultilinkinfo.parent = parent }

func (cvpdnmultilinkinfo *CISCOVPDNMGMTMIB_Cvpdnmultilinkinfo) GetParent() types.Entity { return cvpdnmultilinkinfo.parent }

func (cvpdnmultilinkinfo *CISCOVPDNMGMTMIB_Cvpdnmultilinkinfo) GetParentYangName() string { return "CISCO-VPDN-MGMT-MIB" }

// CISCOVPDNMGMTMIB_Cvpdnsystemtable
// Table of information about the VPDN system for all tunnel
// types.
type CISCOVPDNMGMTMIB_Cvpdnsystemtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the table, containing information about a single type of VPDN
    // tunnel. The type is slice of
    // CISCOVPDNMGMTMIB_Cvpdnsystemtable_Cvpdnsystementry.
    Cvpdnsystementry []CISCOVPDNMGMTMIB_Cvpdnsystemtable_Cvpdnsystementry
}

func (cvpdnsystemtable *CISCOVPDNMGMTMIB_Cvpdnsystemtable) GetFilter() yfilter.YFilter { return cvpdnsystemtable.YFilter }

func (cvpdnsystemtable *CISCOVPDNMGMTMIB_Cvpdnsystemtable) SetFilter(yf yfilter.YFilter) { cvpdnsystemtable.YFilter = yf }

func (cvpdnsystemtable *CISCOVPDNMGMTMIB_Cvpdnsystemtable) GetGoName(yname string) string {
    if yname == "cvpdnSystemEntry" { return "Cvpdnsystementry" }
    return ""
}

func (cvpdnsystemtable *CISCOVPDNMGMTMIB_Cvpdnsystemtable) GetSegmentPath() string {
    return "cvpdnSystemTable"
}

func (cvpdnsystemtable *CISCOVPDNMGMTMIB_Cvpdnsystemtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cvpdnSystemEntry" {
        for _, c := range cvpdnsystemtable.Cvpdnsystementry {
            if cvpdnsystemtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOVPDNMGMTMIB_Cvpdnsystemtable_Cvpdnsystementry{}
        cvpdnsystemtable.Cvpdnsystementry = append(cvpdnsystemtable.Cvpdnsystementry, child)
        return &cvpdnsystemtable.Cvpdnsystementry[len(cvpdnsystemtable.Cvpdnsystementry)-1]
    }
    return nil
}

func (cvpdnsystemtable *CISCOVPDNMGMTMIB_Cvpdnsystemtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cvpdnsystemtable.Cvpdnsystementry {
        children[cvpdnsystemtable.Cvpdnsystementry[i].GetSegmentPath()] = &cvpdnsystemtable.Cvpdnsystementry[i]
    }
    return children
}

func (cvpdnsystemtable *CISCOVPDNMGMTMIB_Cvpdnsystemtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cvpdnsystemtable *CISCOVPDNMGMTMIB_Cvpdnsystemtable) GetBundleName() string { return "cisco_ios_xe" }

func (cvpdnsystemtable *CISCOVPDNMGMTMIB_Cvpdnsystemtable) GetYangName() string { return "cvpdnSystemTable" }

func (cvpdnsystemtable *CISCOVPDNMGMTMIB_Cvpdnsystemtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvpdnsystemtable *CISCOVPDNMGMTMIB_Cvpdnsystemtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvpdnsystemtable *CISCOVPDNMGMTMIB_Cvpdnsystemtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvpdnsystemtable *CISCOVPDNMGMTMIB_Cvpdnsystemtable) SetParent(parent types.Entity) { cvpdnsystemtable.parent = parent }

func (cvpdnsystemtable *CISCOVPDNMGMTMIB_Cvpdnsystemtable) GetParent() types.Entity { return cvpdnsystemtable.parent }

func (cvpdnsystemtable *CISCOVPDNMGMTMIB_Cvpdnsystemtable) GetParentYangName() string { return "CISCO-VPDN-MGMT-MIB" }

// CISCOVPDNMGMTMIB_Cvpdnsystemtable_Cvpdnsystementry
// An entry in the table, containing information about a
// single type of VPDN tunnel.
type CISCOVPDNMGMTMIB_Cvpdnsystemtable_Cvpdnsystementry struct {
    parent types.Entity
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

func (cvpdnsystementry *CISCOVPDNMGMTMIB_Cvpdnsystemtable_Cvpdnsystementry) GetFilter() yfilter.YFilter { return cvpdnsystementry.YFilter }

func (cvpdnsystementry *CISCOVPDNMGMTMIB_Cvpdnsystemtable_Cvpdnsystementry) SetFilter(yf yfilter.YFilter) { cvpdnsystementry.YFilter = yf }

func (cvpdnsystementry *CISCOVPDNMGMTMIB_Cvpdnsystemtable_Cvpdnsystementry) GetGoName(yname string) string {
    if yname == "cvpdnSystemTunnelType" { return "Cvpdnsystemtunneltype" }
    if yname == "cvpdnSystemTunnelTotal" { return "Cvpdnsystemtunneltotal" }
    if yname == "cvpdnSystemSessionTotal" { return "Cvpdnsystemsessiontotal" }
    if yname == "cvpdnSystemDeniedUsersTotal" { return "Cvpdnsystemdenieduserstotal" }
    if yname == "cvpdnSystemInitialConnReq" { return "Cvpdnsysteminitialconnreq" }
    if yname == "cvpdnSystemSuccessConnReq" { return "Cvpdnsystemsuccessconnreq" }
    if yname == "cvpdnSystemFailedConnReq" { return "Cvpdnsystemfailedconnreq" }
    return ""
}

func (cvpdnsystementry *CISCOVPDNMGMTMIB_Cvpdnsystemtable_Cvpdnsystementry) GetSegmentPath() string {
    return "cvpdnSystemEntry" + "[cvpdnSystemTunnelType='" + fmt.Sprintf("%v", cvpdnsystementry.Cvpdnsystemtunneltype) + "']"
}

func (cvpdnsystementry *CISCOVPDNMGMTMIB_Cvpdnsystemtable_Cvpdnsystementry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cvpdnsystementry *CISCOVPDNMGMTMIB_Cvpdnsystemtable_Cvpdnsystementry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cvpdnsystementry *CISCOVPDNMGMTMIB_Cvpdnsystemtable_Cvpdnsystementry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cvpdnSystemTunnelType"] = cvpdnsystementry.Cvpdnsystemtunneltype
    leafs["cvpdnSystemTunnelTotal"] = cvpdnsystementry.Cvpdnsystemtunneltotal
    leafs["cvpdnSystemSessionTotal"] = cvpdnsystementry.Cvpdnsystemsessiontotal
    leafs["cvpdnSystemDeniedUsersTotal"] = cvpdnsystementry.Cvpdnsystemdenieduserstotal
    leafs["cvpdnSystemInitialConnReq"] = cvpdnsystementry.Cvpdnsysteminitialconnreq
    leafs["cvpdnSystemSuccessConnReq"] = cvpdnsystementry.Cvpdnsystemsuccessconnreq
    leafs["cvpdnSystemFailedConnReq"] = cvpdnsystementry.Cvpdnsystemfailedconnreq
    return leafs
}

func (cvpdnsystementry *CISCOVPDNMGMTMIB_Cvpdnsystemtable_Cvpdnsystementry) GetBundleName() string { return "cisco_ios_xe" }

func (cvpdnsystementry *CISCOVPDNMGMTMIB_Cvpdnsystemtable_Cvpdnsystementry) GetYangName() string { return "cvpdnSystemEntry" }

func (cvpdnsystementry *CISCOVPDNMGMTMIB_Cvpdnsystemtable_Cvpdnsystementry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvpdnsystementry *CISCOVPDNMGMTMIB_Cvpdnsystemtable_Cvpdnsystementry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvpdnsystementry *CISCOVPDNMGMTMIB_Cvpdnsystemtable_Cvpdnsystementry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvpdnsystementry *CISCOVPDNMGMTMIB_Cvpdnsystemtable_Cvpdnsystementry) SetParent(parent types.Entity) { cvpdnsystementry.parent = parent }

func (cvpdnsystementry *CISCOVPDNMGMTMIB_Cvpdnsystemtable_Cvpdnsystementry) GetParent() types.Entity { return cvpdnsystementry.parent }

func (cvpdnsystementry *CISCOVPDNMGMTMIB_Cvpdnsystemtable_Cvpdnsystementry) GetParentYangName() string { return "cvpdnSystemTable" }

// CISCOVPDNMGMTMIB_Cvpdntunneltable
// Table of information about the active VPDN tunnels.
type CISCOVPDNMGMTMIB_Cvpdntunneltable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the table, containing information about a single active VPDN
    // tunnel. The type is slice of
    // CISCOVPDNMGMTMIB_Cvpdntunneltable_Cvpdntunnelentry.
    Cvpdntunnelentry []CISCOVPDNMGMTMIB_Cvpdntunneltable_Cvpdntunnelentry
}

func (cvpdntunneltable *CISCOVPDNMGMTMIB_Cvpdntunneltable) GetFilter() yfilter.YFilter { return cvpdntunneltable.YFilter }

func (cvpdntunneltable *CISCOVPDNMGMTMIB_Cvpdntunneltable) SetFilter(yf yfilter.YFilter) { cvpdntunneltable.YFilter = yf }

func (cvpdntunneltable *CISCOVPDNMGMTMIB_Cvpdntunneltable) GetGoName(yname string) string {
    if yname == "cvpdnTunnelEntry" { return "Cvpdntunnelentry" }
    return ""
}

func (cvpdntunneltable *CISCOVPDNMGMTMIB_Cvpdntunneltable) GetSegmentPath() string {
    return "cvpdnTunnelTable"
}

func (cvpdntunneltable *CISCOVPDNMGMTMIB_Cvpdntunneltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cvpdnTunnelEntry" {
        for _, c := range cvpdntunneltable.Cvpdntunnelentry {
            if cvpdntunneltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOVPDNMGMTMIB_Cvpdntunneltable_Cvpdntunnelentry{}
        cvpdntunneltable.Cvpdntunnelentry = append(cvpdntunneltable.Cvpdntunnelentry, child)
        return &cvpdntunneltable.Cvpdntunnelentry[len(cvpdntunneltable.Cvpdntunnelentry)-1]
    }
    return nil
}

func (cvpdntunneltable *CISCOVPDNMGMTMIB_Cvpdntunneltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cvpdntunneltable.Cvpdntunnelentry {
        children[cvpdntunneltable.Cvpdntunnelentry[i].GetSegmentPath()] = &cvpdntunneltable.Cvpdntunnelentry[i]
    }
    return children
}

func (cvpdntunneltable *CISCOVPDNMGMTMIB_Cvpdntunneltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cvpdntunneltable *CISCOVPDNMGMTMIB_Cvpdntunneltable) GetBundleName() string { return "cisco_ios_xe" }

func (cvpdntunneltable *CISCOVPDNMGMTMIB_Cvpdntunneltable) GetYangName() string { return "cvpdnTunnelTable" }

func (cvpdntunneltable *CISCOVPDNMGMTMIB_Cvpdntunneltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvpdntunneltable *CISCOVPDNMGMTMIB_Cvpdntunneltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvpdntunneltable *CISCOVPDNMGMTMIB_Cvpdntunneltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvpdntunneltable *CISCOVPDNMGMTMIB_Cvpdntunneltable) SetParent(parent types.Entity) { cvpdntunneltable.parent = parent }

func (cvpdntunneltable *CISCOVPDNMGMTMIB_Cvpdntunneltable) GetParent() types.Entity { return cvpdntunneltable.parent }

func (cvpdntunneltable *CISCOVPDNMGMTMIB_Cvpdntunneltable) GetParentYangName() string { return "CISCO-VPDN-MGMT-MIB" }

// CISCOVPDNMGMTMIB_Cvpdntunneltable_Cvpdntunnelentry
// An entry in the table, containing information about a
// single active VPDN tunnel.
type CISCOVPDNMGMTMIB_Cvpdntunneltable_Cvpdntunnelentry struct {
    parent types.Entity
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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Cvpdntunnellocalipaddress interface{}

    // The source IP address of the tunnel.  This IP address is the user
    // configurable IP address for Stack Group Biding Protocol (SGBP) via the CLI
    // command: vpdn source-ip. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Cvpdntunnelsourceipaddress interface{}

    // The remote IP address of the tunnel.  This IP address is that of the
    // interface at the remote end of the tunnel. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Cvpdntunnelremoteipaddress interface{}
}

func (cvpdntunnelentry *CISCOVPDNMGMTMIB_Cvpdntunneltable_Cvpdntunnelentry) GetFilter() yfilter.YFilter { return cvpdntunnelentry.YFilter }

func (cvpdntunnelentry *CISCOVPDNMGMTMIB_Cvpdntunneltable_Cvpdntunnelentry) SetFilter(yf yfilter.YFilter) { cvpdntunnelentry.YFilter = yf }

func (cvpdntunnelentry *CISCOVPDNMGMTMIB_Cvpdntunneltable_Cvpdntunnelentry) GetGoName(yname string) string {
    if yname == "cvpdnTunnelTunnelId" { return "Cvpdntunneltunnelid" }
    if yname == "cvpdnTunnelRemoteTunnelId" { return "Cvpdntunnelremotetunnelid" }
    if yname == "cvpdnTunnelLocalName" { return "Cvpdntunnellocalname" }
    if yname == "cvpdnTunnelRemoteName" { return "Cvpdntunnelremotename" }
    if yname == "cvpdnTunnelRemoteEndpointName" { return "Cvpdntunnelremoteendpointname" }
    if yname == "cvpdnTunnelLocalInitConnection" { return "Cvpdntunnellocalinitconnection" }
    if yname == "cvpdnTunnelOrigCause" { return "Cvpdntunnelorigcause" }
    if yname == "cvpdnTunnelState" { return "Cvpdntunnelstate" }
    if yname == "cvpdnTunnelActiveSessions" { return "Cvpdntunnelactivesessions" }
    if yname == "cvpdnTunnelDeniedUsers" { return "Cvpdntunneldeniedusers" }
    if yname == "cvpdnTunnelSoftshut" { return "Cvpdntunnelsoftshut" }
    if yname == "cvpdnTunnelNetworkServiceType" { return "Cvpdntunnelnetworkservicetype" }
    if yname == "cvpdnTunnelLocalIpAddress" { return "Cvpdntunnellocalipaddress" }
    if yname == "cvpdnTunnelSourceIpAddress" { return "Cvpdntunnelsourceipaddress" }
    if yname == "cvpdnTunnelRemoteIpAddress" { return "Cvpdntunnelremoteipaddress" }
    return ""
}

func (cvpdntunnelentry *CISCOVPDNMGMTMIB_Cvpdntunneltable_Cvpdntunnelentry) GetSegmentPath() string {
    return "cvpdnTunnelEntry" + "[cvpdnTunnelTunnelId='" + fmt.Sprintf("%v", cvpdntunnelentry.Cvpdntunneltunnelid) + "']"
}

func (cvpdntunnelentry *CISCOVPDNMGMTMIB_Cvpdntunneltable_Cvpdntunnelentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cvpdntunnelentry *CISCOVPDNMGMTMIB_Cvpdntunneltable_Cvpdntunnelentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cvpdntunnelentry *CISCOVPDNMGMTMIB_Cvpdntunneltable_Cvpdntunnelentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cvpdnTunnelTunnelId"] = cvpdntunnelentry.Cvpdntunneltunnelid
    leafs["cvpdnTunnelRemoteTunnelId"] = cvpdntunnelentry.Cvpdntunnelremotetunnelid
    leafs["cvpdnTunnelLocalName"] = cvpdntunnelentry.Cvpdntunnellocalname
    leafs["cvpdnTunnelRemoteName"] = cvpdntunnelentry.Cvpdntunnelremotename
    leafs["cvpdnTunnelRemoteEndpointName"] = cvpdntunnelentry.Cvpdntunnelremoteendpointname
    leafs["cvpdnTunnelLocalInitConnection"] = cvpdntunnelentry.Cvpdntunnellocalinitconnection
    leafs["cvpdnTunnelOrigCause"] = cvpdntunnelentry.Cvpdntunnelorigcause
    leafs["cvpdnTunnelState"] = cvpdntunnelentry.Cvpdntunnelstate
    leafs["cvpdnTunnelActiveSessions"] = cvpdntunnelentry.Cvpdntunnelactivesessions
    leafs["cvpdnTunnelDeniedUsers"] = cvpdntunnelentry.Cvpdntunneldeniedusers
    leafs["cvpdnTunnelSoftshut"] = cvpdntunnelentry.Cvpdntunnelsoftshut
    leafs["cvpdnTunnelNetworkServiceType"] = cvpdntunnelentry.Cvpdntunnelnetworkservicetype
    leafs["cvpdnTunnelLocalIpAddress"] = cvpdntunnelentry.Cvpdntunnellocalipaddress
    leafs["cvpdnTunnelSourceIpAddress"] = cvpdntunnelentry.Cvpdntunnelsourceipaddress
    leafs["cvpdnTunnelRemoteIpAddress"] = cvpdntunnelentry.Cvpdntunnelremoteipaddress
    return leafs
}

func (cvpdntunnelentry *CISCOVPDNMGMTMIB_Cvpdntunneltable_Cvpdntunnelentry) GetBundleName() string { return "cisco_ios_xe" }

func (cvpdntunnelentry *CISCOVPDNMGMTMIB_Cvpdntunneltable_Cvpdntunnelentry) GetYangName() string { return "cvpdnTunnelEntry" }

func (cvpdntunnelentry *CISCOVPDNMGMTMIB_Cvpdntunneltable_Cvpdntunnelentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvpdntunnelentry *CISCOVPDNMGMTMIB_Cvpdntunneltable_Cvpdntunnelentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvpdntunnelentry *CISCOVPDNMGMTMIB_Cvpdntunneltable_Cvpdntunnelentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvpdntunnelentry *CISCOVPDNMGMTMIB_Cvpdntunneltable_Cvpdntunnelentry) SetParent(parent types.Entity) { cvpdntunnelentry.parent = parent }

func (cvpdntunnelentry *CISCOVPDNMGMTMIB_Cvpdntunneltable_Cvpdntunnelentry) GetParent() types.Entity { return cvpdntunnelentry.parent }

func (cvpdntunnelentry *CISCOVPDNMGMTMIB_Cvpdntunneltable_Cvpdntunnelentry) GetParentYangName() string { return "cvpdnTunnelTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the table, containing information about a single active VPDN
    // tunnel. The type is slice of
    // CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry.
    Cvpdntunnelattrentry []CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry
}

func (cvpdntunnelattrtable *CISCOVPDNMGMTMIB_Cvpdntunnelattrtable) GetFilter() yfilter.YFilter { return cvpdntunnelattrtable.YFilter }

func (cvpdntunnelattrtable *CISCOVPDNMGMTMIB_Cvpdntunnelattrtable) SetFilter(yf yfilter.YFilter) { cvpdntunnelattrtable.YFilter = yf }

func (cvpdntunnelattrtable *CISCOVPDNMGMTMIB_Cvpdntunnelattrtable) GetGoName(yname string) string {
    if yname == "cvpdnTunnelAttrEntry" { return "Cvpdntunnelattrentry" }
    return ""
}

func (cvpdntunnelattrtable *CISCOVPDNMGMTMIB_Cvpdntunnelattrtable) GetSegmentPath() string {
    return "cvpdnTunnelAttrTable"
}

func (cvpdntunnelattrtable *CISCOVPDNMGMTMIB_Cvpdntunnelattrtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cvpdnTunnelAttrEntry" {
        for _, c := range cvpdntunnelattrtable.Cvpdntunnelattrentry {
            if cvpdntunnelattrtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry{}
        cvpdntunnelattrtable.Cvpdntunnelattrentry = append(cvpdntunnelattrtable.Cvpdntunnelattrentry, child)
        return &cvpdntunnelattrtable.Cvpdntunnelattrentry[len(cvpdntunnelattrtable.Cvpdntunnelattrentry)-1]
    }
    return nil
}

func (cvpdntunnelattrtable *CISCOVPDNMGMTMIB_Cvpdntunnelattrtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cvpdntunnelattrtable.Cvpdntunnelattrentry {
        children[cvpdntunnelattrtable.Cvpdntunnelattrentry[i].GetSegmentPath()] = &cvpdntunnelattrtable.Cvpdntunnelattrentry[i]
    }
    return children
}

func (cvpdntunnelattrtable *CISCOVPDNMGMTMIB_Cvpdntunnelattrtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cvpdntunnelattrtable *CISCOVPDNMGMTMIB_Cvpdntunnelattrtable) GetBundleName() string { return "cisco_ios_xe" }

func (cvpdntunnelattrtable *CISCOVPDNMGMTMIB_Cvpdntunnelattrtable) GetYangName() string { return "cvpdnTunnelAttrTable" }

func (cvpdntunnelattrtable *CISCOVPDNMGMTMIB_Cvpdntunnelattrtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvpdntunnelattrtable *CISCOVPDNMGMTMIB_Cvpdntunnelattrtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvpdntunnelattrtable *CISCOVPDNMGMTMIB_Cvpdntunnelattrtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvpdntunnelattrtable *CISCOVPDNMGMTMIB_Cvpdntunnelattrtable) SetParent(parent types.Entity) { cvpdntunnelattrtable.parent = parent }

func (cvpdntunnelattrtable *CISCOVPDNMGMTMIB_Cvpdntunnelattrtable) GetParent() types.Entity { return cvpdntunnelattrtable.parent }

func (cvpdntunnelattrtable *CISCOVPDNMGMTMIB_Cvpdntunnelattrtable) GetParentYangName() string { return "CISCO-VPDN-MGMT-MIB" }

// CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry
// An entry in the table, containing information about a
// single active VPDN tunnel.
type CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry struct {
    parent types.Entity
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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Cvpdntunnelattrlocalipaddress interface{}

    // The source IP address of the tunnel.  This IP address is the user
    // configurable IP address for Stack Group Biding Protocol. The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Cvpdntunnelattrsourceipaddress interface{}

    // The remote IP address of the tunnel.  This IP address is that of the
    // interface at the remote end of the tunnel. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
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

func (cvpdntunnelattrentry *CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry) GetFilter() yfilter.YFilter { return cvpdntunnelattrentry.YFilter }

func (cvpdntunnelattrentry *CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry) SetFilter(yf yfilter.YFilter) { cvpdntunnelattrentry.YFilter = yf }

func (cvpdntunnelattrentry *CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry) GetGoName(yname string) string {
    if yname == "cvpdnSystemTunnelType" { return "Cvpdnsystemtunneltype" }
    if yname == "cvpdnTunnelAttrTunnelId" { return "Cvpdntunnelattrtunnelid" }
    if yname == "cvpdnTunnelAttrRemoteTunnelId" { return "Cvpdntunnelattrremotetunnelid" }
    if yname == "cvpdnTunnelAttrLocalName" { return "Cvpdntunnelattrlocalname" }
    if yname == "cvpdnTunnelAttrRemoteName" { return "Cvpdntunnelattrremotename" }
    if yname == "cvpdnTunnelAttrRemoteEndpointName" { return "Cvpdntunnelattrremoteendpointname" }
    if yname == "cvpdnTunnelAttrLocalInitConnection" { return "Cvpdntunnelattrlocalinitconnection" }
    if yname == "cvpdnTunnelAttrOrigCause" { return "Cvpdntunnelattrorigcause" }
    if yname == "cvpdnTunnelAttrState" { return "Cvpdntunnelattrstate" }
    if yname == "cvpdnTunnelAttrActiveSessions" { return "Cvpdntunnelattractivesessions" }
    if yname == "cvpdnTunnelAttrDeniedUsers" { return "Cvpdntunnelattrdeniedusers" }
    if yname == "cvpdnTunnelAttrSoftshut" { return "Cvpdntunnelattrsoftshut" }
    if yname == "cvpdnTunnelAttrNetworkServiceType" { return "Cvpdntunnelattrnetworkservicetype" }
    if yname == "cvpdnTunnelAttrLocalIpAddress" { return "Cvpdntunnelattrlocalipaddress" }
    if yname == "cvpdnTunnelAttrSourceIpAddress" { return "Cvpdntunnelattrsourceipaddress" }
    if yname == "cvpdnTunnelAttrRemoteIpAddress" { return "Cvpdntunnelattrremoteipaddress" }
    if yname == "cvpdnTunnelAttrLocalInetAddressType" { return "Cvpdntunnelattrlocalinetaddresstype" }
    if yname == "cvpdnTunnelAttrLocalInetAddress" { return "Cvpdntunnelattrlocalinetaddress" }
    if yname == "cvpdnTunnelAttrSourceInetAddressType" { return "Cvpdntunnelattrsourceinetaddresstype" }
    if yname == "cvpdnTunnelAttrSourceInetAddress" { return "Cvpdntunnelattrsourceinetaddress" }
    if yname == "cvpdnTunnelAttrRemoteInetAddressType" { return "Cvpdntunnelattrremoteinetaddresstype" }
    if yname == "cvpdnTunnelAttrRemoteInetAddress" { return "Cvpdntunnelattrremoteinetaddress" }
    return ""
}

func (cvpdntunnelattrentry *CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry) GetSegmentPath() string {
    return "cvpdnTunnelAttrEntry" + "[cvpdnSystemTunnelType='" + fmt.Sprintf("%v", cvpdntunnelattrentry.Cvpdnsystemtunneltype) + "']" + "[cvpdnTunnelAttrTunnelId='" + fmt.Sprintf("%v", cvpdntunnelattrentry.Cvpdntunnelattrtunnelid) + "']"
}

func (cvpdntunnelattrentry *CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cvpdntunnelattrentry *CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cvpdntunnelattrentry *CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cvpdnSystemTunnelType"] = cvpdntunnelattrentry.Cvpdnsystemtunneltype
    leafs["cvpdnTunnelAttrTunnelId"] = cvpdntunnelattrentry.Cvpdntunnelattrtunnelid
    leafs["cvpdnTunnelAttrRemoteTunnelId"] = cvpdntunnelattrentry.Cvpdntunnelattrremotetunnelid
    leafs["cvpdnTunnelAttrLocalName"] = cvpdntunnelattrentry.Cvpdntunnelattrlocalname
    leafs["cvpdnTunnelAttrRemoteName"] = cvpdntunnelattrentry.Cvpdntunnelattrremotename
    leafs["cvpdnTunnelAttrRemoteEndpointName"] = cvpdntunnelattrentry.Cvpdntunnelattrremoteendpointname
    leafs["cvpdnTunnelAttrLocalInitConnection"] = cvpdntunnelattrentry.Cvpdntunnelattrlocalinitconnection
    leafs["cvpdnTunnelAttrOrigCause"] = cvpdntunnelattrentry.Cvpdntunnelattrorigcause
    leafs["cvpdnTunnelAttrState"] = cvpdntunnelattrentry.Cvpdntunnelattrstate
    leafs["cvpdnTunnelAttrActiveSessions"] = cvpdntunnelattrentry.Cvpdntunnelattractivesessions
    leafs["cvpdnTunnelAttrDeniedUsers"] = cvpdntunnelattrentry.Cvpdntunnelattrdeniedusers
    leafs["cvpdnTunnelAttrSoftshut"] = cvpdntunnelattrentry.Cvpdntunnelattrsoftshut
    leafs["cvpdnTunnelAttrNetworkServiceType"] = cvpdntunnelattrentry.Cvpdntunnelattrnetworkservicetype
    leafs["cvpdnTunnelAttrLocalIpAddress"] = cvpdntunnelattrentry.Cvpdntunnelattrlocalipaddress
    leafs["cvpdnTunnelAttrSourceIpAddress"] = cvpdntunnelattrentry.Cvpdntunnelattrsourceipaddress
    leafs["cvpdnTunnelAttrRemoteIpAddress"] = cvpdntunnelattrentry.Cvpdntunnelattrremoteipaddress
    leafs["cvpdnTunnelAttrLocalInetAddressType"] = cvpdntunnelattrentry.Cvpdntunnelattrlocalinetaddresstype
    leafs["cvpdnTunnelAttrLocalInetAddress"] = cvpdntunnelattrentry.Cvpdntunnelattrlocalinetaddress
    leafs["cvpdnTunnelAttrSourceInetAddressType"] = cvpdntunnelattrentry.Cvpdntunnelattrsourceinetaddresstype
    leafs["cvpdnTunnelAttrSourceInetAddress"] = cvpdntunnelattrentry.Cvpdntunnelattrsourceinetaddress
    leafs["cvpdnTunnelAttrRemoteInetAddressType"] = cvpdntunnelattrentry.Cvpdntunnelattrremoteinetaddresstype
    leafs["cvpdnTunnelAttrRemoteInetAddress"] = cvpdntunnelattrentry.Cvpdntunnelattrremoteinetaddress
    return leafs
}

func (cvpdntunnelattrentry *CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry) GetBundleName() string { return "cisco_ios_xe" }

func (cvpdntunnelattrentry *CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry) GetYangName() string { return "cvpdnTunnelAttrEntry" }

func (cvpdntunnelattrentry *CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvpdntunnelattrentry *CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvpdntunnelattrentry *CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvpdntunnelattrentry *CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry) SetParent(parent types.Entity) { cvpdntunnelattrentry.parent = parent }

func (cvpdntunnelattrentry *CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry) GetParent() types.Entity { return cvpdntunnelattrentry.parent }

func (cvpdntunnelattrentry *CISCOVPDNMGMTMIB_Cvpdntunnelattrtable_Cvpdntunnelattrentry) GetParentYangName() string { return "cvpdnTunnelAttrTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the table, containing information about a single user session
    // within the tunnel. The type is slice of
    // CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable_Cvpdntunnelsessionentry.
    Cvpdntunnelsessionentry []CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable_Cvpdntunnelsessionentry
}

func (cvpdntunnelsessiontable *CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable) GetFilter() yfilter.YFilter { return cvpdntunnelsessiontable.YFilter }

func (cvpdntunnelsessiontable *CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable) SetFilter(yf yfilter.YFilter) { cvpdntunnelsessiontable.YFilter = yf }

func (cvpdntunnelsessiontable *CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable) GetGoName(yname string) string {
    if yname == "cvpdnTunnelSessionEntry" { return "Cvpdntunnelsessionentry" }
    return ""
}

func (cvpdntunnelsessiontable *CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable) GetSegmentPath() string {
    return "cvpdnTunnelSessionTable"
}

func (cvpdntunnelsessiontable *CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cvpdnTunnelSessionEntry" {
        for _, c := range cvpdntunnelsessiontable.Cvpdntunnelsessionentry {
            if cvpdntunnelsessiontable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable_Cvpdntunnelsessionentry{}
        cvpdntunnelsessiontable.Cvpdntunnelsessionentry = append(cvpdntunnelsessiontable.Cvpdntunnelsessionentry, child)
        return &cvpdntunnelsessiontable.Cvpdntunnelsessionentry[len(cvpdntunnelsessiontable.Cvpdntunnelsessionentry)-1]
    }
    return nil
}

func (cvpdntunnelsessiontable *CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cvpdntunnelsessiontable.Cvpdntunnelsessionentry {
        children[cvpdntunnelsessiontable.Cvpdntunnelsessionentry[i].GetSegmentPath()] = &cvpdntunnelsessiontable.Cvpdntunnelsessionentry[i]
    }
    return children
}

func (cvpdntunnelsessiontable *CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cvpdntunnelsessiontable *CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable) GetBundleName() string { return "cisco_ios_xe" }

func (cvpdntunnelsessiontable *CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable) GetYangName() string { return "cvpdnTunnelSessionTable" }

func (cvpdntunnelsessiontable *CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvpdntunnelsessiontable *CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvpdntunnelsessiontable *CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvpdntunnelsessiontable *CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable) SetParent(parent types.Entity) { cvpdntunnelsessiontable.parent = parent }

func (cvpdntunnelsessiontable *CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable) GetParent() types.Entity { return cvpdntunnelsessiontable.parent }

func (cvpdntunnelsessiontable *CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable) GetParentYangName() string { return "CISCO-VPDN-MGMT-MIB" }

// CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable_Cvpdntunnelsessionentry
// An entry in the table, containing information about a
// single user session within the tunnel.
type CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable_Cvpdntunnelsessionentry struct {
    parent types.Entity
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

func (cvpdntunnelsessionentry *CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable_Cvpdntunnelsessionentry) GetFilter() yfilter.YFilter { return cvpdntunnelsessionentry.YFilter }

func (cvpdntunnelsessionentry *CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable_Cvpdntunnelsessionentry) SetFilter(yf yfilter.YFilter) { cvpdntunnelsessionentry.YFilter = yf }

func (cvpdntunnelsessionentry *CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable_Cvpdntunnelsessionentry) GetGoName(yname string) string {
    if yname == "cvpdnTunnelTunnelId" { return "Cvpdntunneltunnelid" }
    if yname == "cvpdnTunnelSessionId" { return "Cvpdntunnelsessionid" }
    if yname == "cvpdnTunnelSessionUserName" { return "Cvpdntunnelsessionusername" }
    if yname == "cvpdnTunnelSessionState" { return "Cvpdntunnelsessionstate" }
    if yname == "cvpdnTunnelSessionCallDuration" { return "Cvpdntunnelsessioncallduration" }
    if yname == "cvpdnTunnelSessionPacketsOut" { return "Cvpdntunnelsessionpacketsout" }
    if yname == "cvpdnTunnelSessionBytesOut" { return "Cvpdntunnelsessionbytesout" }
    if yname == "cvpdnTunnelSessionPacketsIn" { return "Cvpdntunnelsessionpacketsin" }
    if yname == "cvpdnTunnelSessionBytesIn" { return "Cvpdntunnelsessionbytesin" }
    if yname == "cvpdnTunnelSessionDeviceType" { return "Cvpdntunnelsessiondevicetype" }
    if yname == "cvpdnTunnelSessionDeviceCallerId" { return "Cvpdntunnelsessiondevicecallerid" }
    if yname == "cvpdnTunnelSessionDevicePhyId" { return "Cvpdntunnelsessiondevicephyid" }
    if yname == "cvpdnTunnelSessionMultilink" { return "Cvpdntunnelsessionmultilink" }
    if yname == "cvpdnTunnelSessionModemSlotIndex" { return "Cvpdntunnelsessionmodemslotindex" }
    if yname == "cvpdnTunnelSessionModemPortIndex" { return "Cvpdntunnelsessionmodemportindex" }
    if yname == "cvpdnTunnelSessionDS1SlotIndex" { return "Cvpdntunnelsessionds1Slotindex" }
    if yname == "cvpdnTunnelSessionDS1PortIndex" { return "Cvpdntunnelsessionds1Portindex" }
    if yname == "cvpdnTunnelSessionDS1ChannelIndex" { return "Cvpdntunnelsessionds1Channelindex" }
    if yname == "cvpdnTunnelSessionModemCallStartTime" { return "Cvpdntunnelsessionmodemcallstarttime" }
    if yname == "cvpdnTunnelSessionModemCallStartIndex" { return "Cvpdntunnelsessionmodemcallstartindex" }
    return ""
}

func (cvpdntunnelsessionentry *CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable_Cvpdntunnelsessionentry) GetSegmentPath() string {
    return "cvpdnTunnelSessionEntry" + "[cvpdnTunnelTunnelId='" + fmt.Sprintf("%v", cvpdntunnelsessionentry.Cvpdntunneltunnelid) + "']" + "[cvpdnTunnelSessionId='" + fmt.Sprintf("%v", cvpdntunnelsessionentry.Cvpdntunnelsessionid) + "']"
}

func (cvpdntunnelsessionentry *CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable_Cvpdntunnelsessionentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cvpdntunnelsessionentry *CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable_Cvpdntunnelsessionentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cvpdntunnelsessionentry *CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable_Cvpdntunnelsessionentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cvpdnTunnelTunnelId"] = cvpdntunnelsessionentry.Cvpdntunneltunnelid
    leafs["cvpdnTunnelSessionId"] = cvpdntunnelsessionentry.Cvpdntunnelsessionid
    leafs["cvpdnTunnelSessionUserName"] = cvpdntunnelsessionentry.Cvpdntunnelsessionusername
    leafs["cvpdnTunnelSessionState"] = cvpdntunnelsessionentry.Cvpdntunnelsessionstate
    leafs["cvpdnTunnelSessionCallDuration"] = cvpdntunnelsessionentry.Cvpdntunnelsessioncallduration
    leafs["cvpdnTunnelSessionPacketsOut"] = cvpdntunnelsessionentry.Cvpdntunnelsessionpacketsout
    leafs["cvpdnTunnelSessionBytesOut"] = cvpdntunnelsessionentry.Cvpdntunnelsessionbytesout
    leafs["cvpdnTunnelSessionPacketsIn"] = cvpdntunnelsessionentry.Cvpdntunnelsessionpacketsin
    leafs["cvpdnTunnelSessionBytesIn"] = cvpdntunnelsessionentry.Cvpdntunnelsessionbytesin
    leafs["cvpdnTunnelSessionDeviceType"] = cvpdntunnelsessionentry.Cvpdntunnelsessiondevicetype
    leafs["cvpdnTunnelSessionDeviceCallerId"] = cvpdntunnelsessionentry.Cvpdntunnelsessiondevicecallerid
    leafs["cvpdnTunnelSessionDevicePhyId"] = cvpdntunnelsessionentry.Cvpdntunnelsessiondevicephyid
    leafs["cvpdnTunnelSessionMultilink"] = cvpdntunnelsessionentry.Cvpdntunnelsessionmultilink
    leafs["cvpdnTunnelSessionModemSlotIndex"] = cvpdntunnelsessionentry.Cvpdntunnelsessionmodemslotindex
    leafs["cvpdnTunnelSessionModemPortIndex"] = cvpdntunnelsessionentry.Cvpdntunnelsessionmodemportindex
    leafs["cvpdnTunnelSessionDS1SlotIndex"] = cvpdntunnelsessionentry.Cvpdntunnelsessionds1Slotindex
    leafs["cvpdnTunnelSessionDS1PortIndex"] = cvpdntunnelsessionentry.Cvpdntunnelsessionds1Portindex
    leafs["cvpdnTunnelSessionDS1ChannelIndex"] = cvpdntunnelsessionentry.Cvpdntunnelsessionds1Channelindex
    leafs["cvpdnTunnelSessionModemCallStartTime"] = cvpdntunnelsessionentry.Cvpdntunnelsessionmodemcallstarttime
    leafs["cvpdnTunnelSessionModemCallStartIndex"] = cvpdntunnelsessionentry.Cvpdntunnelsessionmodemcallstartindex
    return leafs
}

func (cvpdntunnelsessionentry *CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable_Cvpdntunnelsessionentry) GetBundleName() string { return "cisco_ios_xe" }

func (cvpdntunnelsessionentry *CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable_Cvpdntunnelsessionentry) GetYangName() string { return "cvpdnTunnelSessionEntry" }

func (cvpdntunnelsessionentry *CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable_Cvpdntunnelsessionentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvpdntunnelsessionentry *CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable_Cvpdntunnelsessionentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvpdntunnelsessionentry *CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable_Cvpdntunnelsessionentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvpdntunnelsessionentry *CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable_Cvpdntunnelsessionentry) SetParent(parent types.Entity) { cvpdntunnelsessionentry.parent = parent }

func (cvpdntunnelsessionentry *CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable_Cvpdntunnelsessionentry) GetParent() types.Entity { return cvpdntunnelsessionentry.parent }

func (cvpdntunnelsessionentry *CISCOVPDNMGMTMIB_Cvpdntunnelsessiontable_Cvpdntunnelsessionentry) GetParentYangName() string { return "cvpdnTunnelSessionTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the table, containing information about a single session within
    // the tunnel. The type is slice of
    // CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry.
    Cvpdnsessionattrentry []CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry
}

func (cvpdnsessionattrtable *CISCOVPDNMGMTMIB_Cvpdnsessionattrtable) GetFilter() yfilter.YFilter { return cvpdnsessionattrtable.YFilter }

func (cvpdnsessionattrtable *CISCOVPDNMGMTMIB_Cvpdnsessionattrtable) SetFilter(yf yfilter.YFilter) { cvpdnsessionattrtable.YFilter = yf }

func (cvpdnsessionattrtable *CISCOVPDNMGMTMIB_Cvpdnsessionattrtable) GetGoName(yname string) string {
    if yname == "cvpdnSessionAttrEntry" { return "Cvpdnsessionattrentry" }
    return ""
}

func (cvpdnsessionattrtable *CISCOVPDNMGMTMIB_Cvpdnsessionattrtable) GetSegmentPath() string {
    return "cvpdnSessionAttrTable"
}

func (cvpdnsessionattrtable *CISCOVPDNMGMTMIB_Cvpdnsessionattrtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cvpdnSessionAttrEntry" {
        for _, c := range cvpdnsessionattrtable.Cvpdnsessionattrentry {
            if cvpdnsessionattrtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry{}
        cvpdnsessionattrtable.Cvpdnsessionattrentry = append(cvpdnsessionattrtable.Cvpdnsessionattrentry, child)
        return &cvpdnsessionattrtable.Cvpdnsessionattrentry[len(cvpdnsessionattrtable.Cvpdnsessionattrentry)-1]
    }
    return nil
}

func (cvpdnsessionattrtable *CISCOVPDNMGMTMIB_Cvpdnsessionattrtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cvpdnsessionattrtable.Cvpdnsessionattrentry {
        children[cvpdnsessionattrtable.Cvpdnsessionattrentry[i].GetSegmentPath()] = &cvpdnsessionattrtable.Cvpdnsessionattrentry[i]
    }
    return children
}

func (cvpdnsessionattrtable *CISCOVPDNMGMTMIB_Cvpdnsessionattrtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cvpdnsessionattrtable *CISCOVPDNMGMTMIB_Cvpdnsessionattrtable) GetBundleName() string { return "cisco_ios_xe" }

func (cvpdnsessionattrtable *CISCOVPDNMGMTMIB_Cvpdnsessionattrtable) GetYangName() string { return "cvpdnSessionAttrTable" }

func (cvpdnsessionattrtable *CISCOVPDNMGMTMIB_Cvpdnsessionattrtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvpdnsessionattrtable *CISCOVPDNMGMTMIB_Cvpdnsessionattrtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvpdnsessionattrtable *CISCOVPDNMGMTMIB_Cvpdnsessionattrtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvpdnsessionattrtable *CISCOVPDNMGMTMIB_Cvpdnsessionattrtable) SetParent(parent types.Entity) { cvpdnsessionattrtable.parent = parent }

func (cvpdnsessionattrtable *CISCOVPDNMGMTMIB_Cvpdnsessionattrtable) GetParent() types.Entity { return cvpdnsessionattrtable.parent }

func (cvpdnsessionattrtable *CISCOVPDNMGMTMIB_Cvpdnsessionattrtable) GetParentYangName() string { return "CISCO-VPDN-MGMT-MIB" }

// CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry
// An entry in the table, containing information about a
// single session within the tunnel.
type CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry struct {
    parent types.Entity
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

func (cvpdnsessionattrentry *CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry) GetFilter() yfilter.YFilter { return cvpdnsessionattrentry.YFilter }

func (cvpdnsessionattrentry *CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry) SetFilter(yf yfilter.YFilter) { cvpdnsessionattrentry.YFilter = yf }

func (cvpdnsessionattrentry *CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry) GetGoName(yname string) string {
    if yname == "cvpdnSystemTunnelType" { return "Cvpdnsystemtunneltype" }
    if yname == "cvpdnTunnelAttrTunnelId" { return "Cvpdntunnelattrtunnelid" }
    if yname == "cvpdnSessionAttrSessionId" { return "Cvpdnsessionattrsessionid" }
    if yname == "cvpdnSessionAttrUserName" { return "Cvpdnsessionattrusername" }
    if yname == "cvpdnSessionAttrState" { return "Cvpdnsessionattrstate" }
    if yname == "cvpdnSessionAttrCallDuration" { return "Cvpdnsessionattrcallduration" }
    if yname == "cvpdnSessionAttrPacketsOut" { return "Cvpdnsessionattrpacketsout" }
    if yname == "cvpdnSessionAttrBytesOut" { return "Cvpdnsessionattrbytesout" }
    if yname == "cvpdnSessionAttrPacketsIn" { return "Cvpdnsessionattrpacketsin" }
    if yname == "cvpdnSessionAttrBytesIn" { return "Cvpdnsessionattrbytesin" }
    if yname == "cvpdnSessionAttrDeviceType" { return "Cvpdnsessionattrdevicetype" }
    if yname == "cvpdnSessionAttrDeviceCallerId" { return "Cvpdnsessionattrdevicecallerid" }
    if yname == "cvpdnSessionAttrDevicePhyId" { return "Cvpdnsessionattrdevicephyid" }
    if yname == "cvpdnSessionAttrMultilink" { return "Cvpdnsessionattrmultilink" }
    if yname == "cvpdnSessionAttrModemSlotIndex" { return "Cvpdnsessionattrmodemslotindex" }
    if yname == "cvpdnSessionAttrModemPortIndex" { return "Cvpdnsessionattrmodemportindex" }
    if yname == "cvpdnSessionAttrDS1SlotIndex" { return "Cvpdnsessionattrds1Slotindex" }
    if yname == "cvpdnSessionAttrDS1PortIndex" { return "Cvpdnsessionattrds1Portindex" }
    if yname == "cvpdnSessionAttrDS1ChannelIndex" { return "Cvpdnsessionattrds1Channelindex" }
    if yname == "cvpdnSessionAttrModemCallStartTime" { return "Cvpdnsessionattrmodemcallstarttime" }
    if yname == "cvpdnSessionAttrModemCallStartIndex" { return "Cvpdnsessionattrmodemcallstartindex" }
    if yname == "cvpdnSessionAttrVirtualCircuitID" { return "Cvpdnsessionattrvirtualcircuitid" }
    if yname == "cvpdnSessionAttrSentPktsDropped" { return "Cvpdnsessionattrsentpktsdropped" }
    if yname == "cvpdnSessionAttrRecvPktsDropped" { return "Cvpdnsessionattrrecvpktsdropped" }
    if yname == "cvpdnSessionAttrMultilinkBundle" { return "Cvpdnsessionattrmultilinkbundle" }
    if yname == "cvpdnSessionAttrMultilinkIfIndex" { return "Cvpdnsessionattrmultilinkifindex" }
    return ""
}

func (cvpdnsessionattrentry *CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry) GetSegmentPath() string {
    return "cvpdnSessionAttrEntry" + "[cvpdnSystemTunnelType='" + fmt.Sprintf("%v", cvpdnsessionattrentry.Cvpdnsystemtunneltype) + "']" + "[cvpdnTunnelAttrTunnelId='" + fmt.Sprintf("%v", cvpdnsessionattrentry.Cvpdntunnelattrtunnelid) + "']" + "[cvpdnSessionAttrSessionId='" + fmt.Sprintf("%v", cvpdnsessionattrentry.Cvpdnsessionattrsessionid) + "']"
}

func (cvpdnsessionattrentry *CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cvpdnsessionattrentry *CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cvpdnsessionattrentry *CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cvpdnSystemTunnelType"] = cvpdnsessionattrentry.Cvpdnsystemtunneltype
    leafs["cvpdnTunnelAttrTunnelId"] = cvpdnsessionattrentry.Cvpdntunnelattrtunnelid
    leafs["cvpdnSessionAttrSessionId"] = cvpdnsessionattrentry.Cvpdnsessionattrsessionid
    leafs["cvpdnSessionAttrUserName"] = cvpdnsessionattrentry.Cvpdnsessionattrusername
    leafs["cvpdnSessionAttrState"] = cvpdnsessionattrentry.Cvpdnsessionattrstate
    leafs["cvpdnSessionAttrCallDuration"] = cvpdnsessionattrentry.Cvpdnsessionattrcallduration
    leafs["cvpdnSessionAttrPacketsOut"] = cvpdnsessionattrentry.Cvpdnsessionattrpacketsout
    leafs["cvpdnSessionAttrBytesOut"] = cvpdnsessionattrentry.Cvpdnsessionattrbytesout
    leafs["cvpdnSessionAttrPacketsIn"] = cvpdnsessionattrentry.Cvpdnsessionattrpacketsin
    leafs["cvpdnSessionAttrBytesIn"] = cvpdnsessionattrentry.Cvpdnsessionattrbytesin
    leafs["cvpdnSessionAttrDeviceType"] = cvpdnsessionattrentry.Cvpdnsessionattrdevicetype
    leafs["cvpdnSessionAttrDeviceCallerId"] = cvpdnsessionattrentry.Cvpdnsessionattrdevicecallerid
    leafs["cvpdnSessionAttrDevicePhyId"] = cvpdnsessionattrentry.Cvpdnsessionattrdevicephyid
    leafs["cvpdnSessionAttrMultilink"] = cvpdnsessionattrentry.Cvpdnsessionattrmultilink
    leafs["cvpdnSessionAttrModemSlotIndex"] = cvpdnsessionattrentry.Cvpdnsessionattrmodemslotindex
    leafs["cvpdnSessionAttrModemPortIndex"] = cvpdnsessionattrentry.Cvpdnsessionattrmodemportindex
    leafs["cvpdnSessionAttrDS1SlotIndex"] = cvpdnsessionattrentry.Cvpdnsessionattrds1Slotindex
    leafs["cvpdnSessionAttrDS1PortIndex"] = cvpdnsessionattrentry.Cvpdnsessionattrds1Portindex
    leafs["cvpdnSessionAttrDS1ChannelIndex"] = cvpdnsessionattrentry.Cvpdnsessionattrds1Channelindex
    leafs["cvpdnSessionAttrModemCallStartTime"] = cvpdnsessionattrentry.Cvpdnsessionattrmodemcallstarttime
    leafs["cvpdnSessionAttrModemCallStartIndex"] = cvpdnsessionattrentry.Cvpdnsessionattrmodemcallstartindex
    leafs["cvpdnSessionAttrVirtualCircuitID"] = cvpdnsessionattrentry.Cvpdnsessionattrvirtualcircuitid
    leafs["cvpdnSessionAttrSentPktsDropped"] = cvpdnsessionattrentry.Cvpdnsessionattrsentpktsdropped
    leafs["cvpdnSessionAttrRecvPktsDropped"] = cvpdnsessionattrentry.Cvpdnsessionattrrecvpktsdropped
    leafs["cvpdnSessionAttrMultilinkBundle"] = cvpdnsessionattrentry.Cvpdnsessionattrmultilinkbundle
    leafs["cvpdnSessionAttrMultilinkIfIndex"] = cvpdnsessionattrentry.Cvpdnsessionattrmultilinkifindex
    return leafs
}

func (cvpdnsessionattrentry *CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry) GetBundleName() string { return "cisco_ios_xe" }

func (cvpdnsessionattrentry *CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry) GetYangName() string { return "cvpdnSessionAttrEntry" }

func (cvpdnsessionattrentry *CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvpdnsessionattrentry *CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvpdnsessionattrentry *CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvpdnsessionattrentry *CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry) SetParent(parent types.Entity) { cvpdnsessionattrentry.parent = parent }

func (cvpdnsessionattrentry *CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry) GetParent() types.Entity { return cvpdnsessionattrentry.parent }

func (cvpdnsessionattrentry *CISCOVPDNMGMTMIB_Cvpdnsessionattrtable_Cvpdnsessionattrentry) GetParentYangName() string { return "cvpdnSessionAttrTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the table, containing failure history relevant to an user name.
    // The type is slice of
    // CISCOVPDNMGMTMIB_Cvpdnusertofailhistinfotable_Cvpdnusertofailhistinfoentry.
    Cvpdnusertofailhistinfoentry []CISCOVPDNMGMTMIB_Cvpdnusertofailhistinfotable_Cvpdnusertofailhistinfoentry
}

func (cvpdnusertofailhistinfotable *CISCOVPDNMGMTMIB_Cvpdnusertofailhistinfotable) GetFilter() yfilter.YFilter { return cvpdnusertofailhistinfotable.YFilter }

func (cvpdnusertofailhistinfotable *CISCOVPDNMGMTMIB_Cvpdnusertofailhistinfotable) SetFilter(yf yfilter.YFilter) { cvpdnusertofailhistinfotable.YFilter = yf }

func (cvpdnusertofailhistinfotable *CISCOVPDNMGMTMIB_Cvpdnusertofailhistinfotable) GetGoName(yname string) string {
    if yname == "cvpdnUserToFailHistInfoEntry" { return "Cvpdnusertofailhistinfoentry" }
    return ""
}

func (cvpdnusertofailhistinfotable *CISCOVPDNMGMTMIB_Cvpdnusertofailhistinfotable) GetSegmentPath() string {
    return "cvpdnUserToFailHistInfoTable"
}

func (cvpdnusertofailhistinfotable *CISCOVPDNMGMTMIB_Cvpdnusertofailhistinfotable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cvpdnUserToFailHistInfoEntry" {
        for _, c := range cvpdnusertofailhistinfotable.Cvpdnusertofailhistinfoentry {
            if cvpdnusertofailhistinfotable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOVPDNMGMTMIB_Cvpdnusertofailhistinfotable_Cvpdnusertofailhistinfoentry{}
        cvpdnusertofailhistinfotable.Cvpdnusertofailhistinfoentry = append(cvpdnusertofailhistinfotable.Cvpdnusertofailhistinfoentry, child)
        return &cvpdnusertofailhistinfotable.Cvpdnusertofailhistinfoentry[len(cvpdnusertofailhistinfotable.Cvpdnusertofailhistinfoentry)-1]
    }
    return nil
}

func (cvpdnusertofailhistinfotable *CISCOVPDNMGMTMIB_Cvpdnusertofailhistinfotable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cvpdnusertofailhistinfotable.Cvpdnusertofailhistinfoentry {
        children[cvpdnusertofailhistinfotable.Cvpdnusertofailhistinfoentry[i].GetSegmentPath()] = &cvpdnusertofailhistinfotable.Cvpdnusertofailhistinfoentry[i]
    }
    return children
}

func (cvpdnusertofailhistinfotable *CISCOVPDNMGMTMIB_Cvpdnusertofailhistinfotable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cvpdnusertofailhistinfotable *CISCOVPDNMGMTMIB_Cvpdnusertofailhistinfotable) GetBundleName() string { return "cisco_ios_xe" }

func (cvpdnusertofailhistinfotable *CISCOVPDNMGMTMIB_Cvpdnusertofailhistinfotable) GetYangName() string { return "cvpdnUserToFailHistInfoTable" }

func (cvpdnusertofailhistinfotable *CISCOVPDNMGMTMIB_Cvpdnusertofailhistinfotable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvpdnusertofailhistinfotable *CISCOVPDNMGMTMIB_Cvpdnusertofailhistinfotable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvpdnusertofailhistinfotable *CISCOVPDNMGMTMIB_Cvpdnusertofailhistinfotable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvpdnusertofailhistinfotable *CISCOVPDNMGMTMIB_Cvpdnusertofailhistinfotable) SetParent(parent types.Entity) { cvpdnusertofailhistinfotable.parent = parent }

func (cvpdnusertofailhistinfotable *CISCOVPDNMGMTMIB_Cvpdnusertofailhistinfotable) GetParent() types.Entity { return cvpdnusertofailhistinfotable.parent }

func (cvpdnusertofailhistinfotable *CISCOVPDNMGMTMIB_Cvpdnusertofailhistinfotable) GetParentYangName() string { return "CISCO-VPDN-MGMT-MIB" }

// CISCOVPDNMGMTMIB_Cvpdnusertofailhistinfotable_Cvpdnusertofailhistinfoentry
// An entry in the table, containing failure history
// relevant to an user name.
type CISCOVPDNMGMTMIB_Cvpdnusertofailhistinfotable_Cvpdnusertofailhistinfoentry struct {
    parent types.Entity
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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Cvpdnunametofailhistsourceip interface{}

    // The destination IP address of the tunnel in which the failure occurred. 
    // This IP address is that of the interface at the receiver end of the tunnel.
    // The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
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

func (cvpdnusertofailhistinfoentry *CISCOVPDNMGMTMIB_Cvpdnusertofailhistinfotable_Cvpdnusertofailhistinfoentry) GetFilter() yfilter.YFilter { return cvpdnusertofailhistinfoentry.YFilter }

func (cvpdnusertofailhistinfoentry *CISCOVPDNMGMTMIB_Cvpdnusertofailhistinfotable_Cvpdnusertofailhistinfoentry) SetFilter(yf yfilter.YFilter) { cvpdnusertofailhistinfoentry.YFilter = yf }

func (cvpdnusertofailhistinfoentry *CISCOVPDNMGMTMIB_Cvpdnusertofailhistinfotable_Cvpdnusertofailhistinfoentry) GetGoName(yname string) string {
    if yname == "cvpdnUnameToFailHistUname" { return "Cvpdnunametofailhistuname" }
    if yname == "cvpdnUnameToFailHistTunnelId" { return "Cvpdnunametofailhisttunnelid" }
    if yname == "cvpdnUnameToFailHistUserId" { return "Cvpdnunametofailhistuserid" }
    if yname == "cvpdnUnameToFailHistLocalInitConn" { return "Cvpdnunametofailhistlocalinitconn" }
    if yname == "cvpdnUnameToFailHistLocalName" { return "Cvpdnunametofailhistlocalname" }
    if yname == "cvpdnUnameToFailHistRemoteName" { return "Cvpdnunametofailhistremotename" }
    if yname == "cvpdnUnameToFailHistSourceIp" { return "Cvpdnunametofailhistsourceip" }
    if yname == "cvpdnUnameToFailHistDestIp" { return "Cvpdnunametofailhistdestip" }
    if yname == "cvpdnUnameToFailHistCount" { return "Cvpdnunametofailhistcount" }
    if yname == "cvpdnUnameToFailHistFailTime" { return "Cvpdnunametofailhistfailtime" }
    if yname == "cvpdnUnameToFailHistFailType" { return "Cvpdnunametofailhistfailtype" }
    if yname == "cvpdnUnameToFailHistFailReason" { return "Cvpdnunametofailhistfailreason" }
    if yname == "cvpdnUnameToFailHistSourceInetType" { return "Cvpdnunametofailhistsourceinettype" }
    if yname == "cvpdnUnameToFailHistSourceInetAddr" { return "Cvpdnunametofailhistsourceinetaddr" }
    if yname == "cvpdnUnameToFailHistDestInetType" { return "Cvpdnunametofailhistdestinettype" }
    if yname == "cvpdnUnameToFailHistDestInetAddr" { return "Cvpdnunametofailhistdestinetaddr" }
    return ""
}

func (cvpdnusertofailhistinfoentry *CISCOVPDNMGMTMIB_Cvpdnusertofailhistinfotable_Cvpdnusertofailhistinfoentry) GetSegmentPath() string {
    return "cvpdnUserToFailHistInfoEntry" + "[cvpdnUnameToFailHistUname='" + fmt.Sprintf("%v", cvpdnusertofailhistinfoentry.Cvpdnunametofailhistuname) + "']" + "[cvpdnUnameToFailHistTunnelId='" + fmt.Sprintf("%v", cvpdnusertofailhistinfoentry.Cvpdnunametofailhisttunnelid) + "']"
}

func (cvpdnusertofailhistinfoentry *CISCOVPDNMGMTMIB_Cvpdnusertofailhistinfotable_Cvpdnusertofailhistinfoentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cvpdnusertofailhistinfoentry *CISCOVPDNMGMTMIB_Cvpdnusertofailhistinfotable_Cvpdnusertofailhistinfoentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cvpdnusertofailhistinfoentry *CISCOVPDNMGMTMIB_Cvpdnusertofailhistinfotable_Cvpdnusertofailhistinfoentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cvpdnUnameToFailHistUname"] = cvpdnusertofailhistinfoentry.Cvpdnunametofailhistuname
    leafs["cvpdnUnameToFailHistTunnelId"] = cvpdnusertofailhistinfoentry.Cvpdnunametofailhisttunnelid
    leafs["cvpdnUnameToFailHistUserId"] = cvpdnusertofailhistinfoentry.Cvpdnunametofailhistuserid
    leafs["cvpdnUnameToFailHistLocalInitConn"] = cvpdnusertofailhistinfoentry.Cvpdnunametofailhistlocalinitconn
    leafs["cvpdnUnameToFailHistLocalName"] = cvpdnusertofailhistinfoentry.Cvpdnunametofailhistlocalname
    leafs["cvpdnUnameToFailHistRemoteName"] = cvpdnusertofailhistinfoentry.Cvpdnunametofailhistremotename
    leafs["cvpdnUnameToFailHistSourceIp"] = cvpdnusertofailhistinfoentry.Cvpdnunametofailhistsourceip
    leafs["cvpdnUnameToFailHistDestIp"] = cvpdnusertofailhistinfoentry.Cvpdnunametofailhistdestip
    leafs["cvpdnUnameToFailHistCount"] = cvpdnusertofailhistinfoentry.Cvpdnunametofailhistcount
    leafs["cvpdnUnameToFailHistFailTime"] = cvpdnusertofailhistinfoentry.Cvpdnunametofailhistfailtime
    leafs["cvpdnUnameToFailHistFailType"] = cvpdnusertofailhistinfoentry.Cvpdnunametofailhistfailtype
    leafs["cvpdnUnameToFailHistFailReason"] = cvpdnusertofailhistinfoentry.Cvpdnunametofailhistfailreason
    leafs["cvpdnUnameToFailHistSourceInetType"] = cvpdnusertofailhistinfoentry.Cvpdnunametofailhistsourceinettype
    leafs["cvpdnUnameToFailHistSourceInetAddr"] = cvpdnusertofailhistinfoentry.Cvpdnunametofailhistsourceinetaddr
    leafs["cvpdnUnameToFailHistDestInetType"] = cvpdnusertofailhistinfoentry.Cvpdnunametofailhistdestinettype
    leafs["cvpdnUnameToFailHistDestInetAddr"] = cvpdnusertofailhistinfoentry.Cvpdnunametofailhistdestinetaddr
    return leafs
}

func (cvpdnusertofailhistinfoentry *CISCOVPDNMGMTMIB_Cvpdnusertofailhistinfotable_Cvpdnusertofailhistinfoentry) GetBundleName() string { return "cisco_ios_xe" }

func (cvpdnusertofailhistinfoentry *CISCOVPDNMGMTMIB_Cvpdnusertofailhistinfotable_Cvpdnusertofailhistinfoentry) GetYangName() string { return "cvpdnUserToFailHistInfoEntry" }

func (cvpdnusertofailhistinfoentry *CISCOVPDNMGMTMIB_Cvpdnusertofailhistinfotable_Cvpdnusertofailhistinfoentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvpdnusertofailhistinfoentry *CISCOVPDNMGMTMIB_Cvpdnusertofailhistinfotable_Cvpdnusertofailhistinfoentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvpdnusertofailhistinfoentry *CISCOVPDNMGMTMIB_Cvpdnusertofailhistinfotable_Cvpdnusertofailhistinfoentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvpdnusertofailhistinfoentry *CISCOVPDNMGMTMIB_Cvpdnusertofailhistinfotable_Cvpdnusertofailhistinfoentry) SetParent(parent types.Entity) { cvpdnusertofailhistinfoentry.parent = parent }

func (cvpdnusertofailhistinfoentry *CISCOVPDNMGMTMIB_Cvpdnusertofailhistinfotable_Cvpdnusertofailhistinfoentry) GetParent() types.Entity { return cvpdnusertofailhistinfoentry.parent }

func (cvpdnusertofailhistinfoentry *CISCOVPDNMGMTMIB_Cvpdnusertofailhistinfotable_Cvpdnusertofailhistinfoentry) GetParentYangName() string { return "cvpdnUserToFailHistInfoTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the table, containing information about a single VPDN template.
    // The type is slice of
    // CISCOVPDNMGMTMIB_Cvpdntemplatetable_Cvpdntemplateentry.
    Cvpdntemplateentry []CISCOVPDNMGMTMIB_Cvpdntemplatetable_Cvpdntemplateentry
}

func (cvpdntemplatetable *CISCOVPDNMGMTMIB_Cvpdntemplatetable) GetFilter() yfilter.YFilter { return cvpdntemplatetable.YFilter }

func (cvpdntemplatetable *CISCOVPDNMGMTMIB_Cvpdntemplatetable) SetFilter(yf yfilter.YFilter) { cvpdntemplatetable.YFilter = yf }

func (cvpdntemplatetable *CISCOVPDNMGMTMIB_Cvpdntemplatetable) GetGoName(yname string) string {
    if yname == "cvpdnTemplateEntry" { return "Cvpdntemplateentry" }
    return ""
}

func (cvpdntemplatetable *CISCOVPDNMGMTMIB_Cvpdntemplatetable) GetSegmentPath() string {
    return "cvpdnTemplateTable"
}

func (cvpdntemplatetable *CISCOVPDNMGMTMIB_Cvpdntemplatetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cvpdnTemplateEntry" {
        for _, c := range cvpdntemplatetable.Cvpdntemplateentry {
            if cvpdntemplatetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOVPDNMGMTMIB_Cvpdntemplatetable_Cvpdntemplateentry{}
        cvpdntemplatetable.Cvpdntemplateentry = append(cvpdntemplatetable.Cvpdntemplateentry, child)
        return &cvpdntemplatetable.Cvpdntemplateentry[len(cvpdntemplatetable.Cvpdntemplateentry)-1]
    }
    return nil
}

func (cvpdntemplatetable *CISCOVPDNMGMTMIB_Cvpdntemplatetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cvpdntemplatetable.Cvpdntemplateentry {
        children[cvpdntemplatetable.Cvpdntemplateentry[i].GetSegmentPath()] = &cvpdntemplatetable.Cvpdntemplateentry[i]
    }
    return children
}

func (cvpdntemplatetable *CISCOVPDNMGMTMIB_Cvpdntemplatetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cvpdntemplatetable *CISCOVPDNMGMTMIB_Cvpdntemplatetable) GetBundleName() string { return "cisco_ios_xe" }

func (cvpdntemplatetable *CISCOVPDNMGMTMIB_Cvpdntemplatetable) GetYangName() string { return "cvpdnTemplateTable" }

func (cvpdntemplatetable *CISCOVPDNMGMTMIB_Cvpdntemplatetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvpdntemplatetable *CISCOVPDNMGMTMIB_Cvpdntemplatetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvpdntemplatetable *CISCOVPDNMGMTMIB_Cvpdntemplatetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvpdntemplatetable *CISCOVPDNMGMTMIB_Cvpdntemplatetable) SetParent(parent types.Entity) { cvpdntemplatetable.parent = parent }

func (cvpdntemplatetable *CISCOVPDNMGMTMIB_Cvpdntemplatetable) GetParent() types.Entity { return cvpdntemplatetable.parent }

func (cvpdntemplatetable *CISCOVPDNMGMTMIB_Cvpdntemplatetable) GetParentYangName() string { return "CISCO-VPDN-MGMT-MIB" }

// CISCOVPDNMGMTMIB_Cvpdntemplatetable_Cvpdntemplateentry
// An entry in the table, containing information about a
// single VPDN template.
type CISCOVPDNMGMTMIB_Cvpdntemplatetable_Cvpdntemplateentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the VPDN template. The type is string
    // with length: 1..255.
    Cvpdntemplatename interface{}

    // The total number of active session in all groups associated with the
    // template. The type is interface{} with range: 0..4294967295. Units are
    // sessions.
    Cvpdntemplateactivesessions interface{}
}

func (cvpdntemplateentry *CISCOVPDNMGMTMIB_Cvpdntemplatetable_Cvpdntemplateentry) GetFilter() yfilter.YFilter { return cvpdntemplateentry.YFilter }

func (cvpdntemplateentry *CISCOVPDNMGMTMIB_Cvpdntemplatetable_Cvpdntemplateentry) SetFilter(yf yfilter.YFilter) { cvpdntemplateentry.YFilter = yf }

func (cvpdntemplateentry *CISCOVPDNMGMTMIB_Cvpdntemplatetable_Cvpdntemplateentry) GetGoName(yname string) string {
    if yname == "cvpdnTemplateName" { return "Cvpdntemplatename" }
    if yname == "cvpdnTemplateActiveSessions" { return "Cvpdntemplateactivesessions" }
    return ""
}

func (cvpdntemplateentry *CISCOVPDNMGMTMIB_Cvpdntemplatetable_Cvpdntemplateentry) GetSegmentPath() string {
    return "cvpdnTemplateEntry" + "[cvpdnTemplateName='" + fmt.Sprintf("%v", cvpdntemplateentry.Cvpdntemplatename) + "']"
}

func (cvpdntemplateentry *CISCOVPDNMGMTMIB_Cvpdntemplatetable_Cvpdntemplateentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cvpdntemplateentry *CISCOVPDNMGMTMIB_Cvpdntemplatetable_Cvpdntemplateentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cvpdntemplateentry *CISCOVPDNMGMTMIB_Cvpdntemplatetable_Cvpdntemplateentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cvpdnTemplateName"] = cvpdntemplateentry.Cvpdntemplatename
    leafs["cvpdnTemplateActiveSessions"] = cvpdntemplateentry.Cvpdntemplateactivesessions
    return leafs
}

func (cvpdntemplateentry *CISCOVPDNMGMTMIB_Cvpdntemplatetable_Cvpdntemplateentry) GetBundleName() string { return "cisco_ios_xe" }

func (cvpdntemplateentry *CISCOVPDNMGMTMIB_Cvpdntemplatetable_Cvpdntemplateentry) GetYangName() string { return "cvpdnTemplateEntry" }

func (cvpdntemplateentry *CISCOVPDNMGMTMIB_Cvpdntemplatetable_Cvpdntemplateentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvpdntemplateentry *CISCOVPDNMGMTMIB_Cvpdntemplatetable_Cvpdntemplateentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvpdntemplateentry *CISCOVPDNMGMTMIB_Cvpdntemplatetable_Cvpdntemplateentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvpdntemplateentry *CISCOVPDNMGMTMIB_Cvpdntemplatetable_Cvpdntemplateentry) SetParent(parent types.Entity) { cvpdntemplateentry.parent = parent }

func (cvpdntemplateentry *CISCOVPDNMGMTMIB_Cvpdntemplatetable_Cvpdntemplateentry) GetParent() types.Entity { return cvpdntemplateentry.parent }

func (cvpdntemplateentry *CISCOVPDNMGMTMIB_Cvpdntemplatetable_Cvpdntemplateentry) GetParentYangName() string { return "cvpdnTemplateTable" }

// CISCOVPDNMGMTMIB_Cvpdnbundletable
// Table that describes the multilink PPP attributes of the
// active VPDN sessions.
type CISCOVPDNMGMTMIB_Cvpdnbundletable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in this table represents an active multilink PPP bundle that
    // belongs to a VPDN tunnel. The type is slice of
    // CISCOVPDNMGMTMIB_Cvpdnbundletable_Cvpdnbundleentry.
    Cvpdnbundleentry []CISCOVPDNMGMTMIB_Cvpdnbundletable_Cvpdnbundleentry
}

func (cvpdnbundletable *CISCOVPDNMGMTMIB_Cvpdnbundletable) GetFilter() yfilter.YFilter { return cvpdnbundletable.YFilter }

func (cvpdnbundletable *CISCOVPDNMGMTMIB_Cvpdnbundletable) SetFilter(yf yfilter.YFilter) { cvpdnbundletable.YFilter = yf }

func (cvpdnbundletable *CISCOVPDNMGMTMIB_Cvpdnbundletable) GetGoName(yname string) string {
    if yname == "cvpdnBundleEntry" { return "Cvpdnbundleentry" }
    return ""
}

func (cvpdnbundletable *CISCOVPDNMGMTMIB_Cvpdnbundletable) GetSegmentPath() string {
    return "cvpdnBundleTable"
}

func (cvpdnbundletable *CISCOVPDNMGMTMIB_Cvpdnbundletable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cvpdnBundleEntry" {
        for _, c := range cvpdnbundletable.Cvpdnbundleentry {
            if cvpdnbundletable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOVPDNMGMTMIB_Cvpdnbundletable_Cvpdnbundleentry{}
        cvpdnbundletable.Cvpdnbundleentry = append(cvpdnbundletable.Cvpdnbundleentry, child)
        return &cvpdnbundletable.Cvpdnbundleentry[len(cvpdnbundletable.Cvpdnbundleentry)-1]
    }
    return nil
}

func (cvpdnbundletable *CISCOVPDNMGMTMIB_Cvpdnbundletable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cvpdnbundletable.Cvpdnbundleentry {
        children[cvpdnbundletable.Cvpdnbundleentry[i].GetSegmentPath()] = &cvpdnbundletable.Cvpdnbundleentry[i]
    }
    return children
}

func (cvpdnbundletable *CISCOVPDNMGMTMIB_Cvpdnbundletable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cvpdnbundletable *CISCOVPDNMGMTMIB_Cvpdnbundletable) GetBundleName() string { return "cisco_ios_xe" }

func (cvpdnbundletable *CISCOVPDNMGMTMIB_Cvpdnbundletable) GetYangName() string { return "cvpdnBundleTable" }

func (cvpdnbundletable *CISCOVPDNMGMTMIB_Cvpdnbundletable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvpdnbundletable *CISCOVPDNMGMTMIB_Cvpdnbundletable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvpdnbundletable *CISCOVPDNMGMTMIB_Cvpdnbundletable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvpdnbundletable *CISCOVPDNMGMTMIB_Cvpdnbundletable) SetParent(parent types.Entity) { cvpdnbundletable.parent = parent }

func (cvpdnbundletable *CISCOVPDNMGMTMIB_Cvpdnbundletable) GetParent() types.Entity { return cvpdnbundletable.parent }

func (cvpdnbundletable *CISCOVPDNMGMTMIB_Cvpdnbundletable) GetParentYangName() string { return "CISCO-VPDN-MGMT-MIB" }

// CISCOVPDNMGMTMIB_Cvpdnbundletable_Cvpdnbundleentry
// An entry in this table represents an active multilink PPP
// bundle that belongs to a VPDN tunnel.
type CISCOVPDNMGMTMIB_Cvpdnbundletable_Cvpdnbundleentry struct {
    parent types.Entity
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

func (cvpdnbundleentry *CISCOVPDNMGMTMIB_Cvpdnbundletable_Cvpdnbundleentry) GetFilter() yfilter.YFilter { return cvpdnbundleentry.YFilter }

func (cvpdnbundleentry *CISCOVPDNMGMTMIB_Cvpdnbundletable_Cvpdnbundleentry) SetFilter(yf yfilter.YFilter) { cvpdnbundleentry.YFilter = yf }

func (cvpdnbundleentry *CISCOVPDNMGMTMIB_Cvpdnbundletable_Cvpdnbundleentry) GetGoName(yname string) string {
    if yname == "cvpdnBundleName" { return "Cvpdnbundlename" }
    if yname == "cvpdnBundleLinkCount" { return "Cvpdnbundlelinkcount" }
    if yname == "cvpdnBundleEndpointType" { return "Cvpdnbundleendpointtype" }
    if yname == "cvpdnBundleEndpoint" { return "Cvpdnbundleendpoint" }
    if yname == "cvpdnBundlePeerIpAddrType" { return "Cvpdnbundlepeeripaddrtype" }
    if yname == "cvpdnBundlePeerIpAddr" { return "Cvpdnbundlepeeripaddr" }
    if yname == "cvpdnBundleEndpointClass" { return "Cvpdnbundleendpointclass" }
    return ""
}

func (cvpdnbundleentry *CISCOVPDNMGMTMIB_Cvpdnbundletable_Cvpdnbundleentry) GetSegmentPath() string {
    return "cvpdnBundleEntry" + "[cvpdnBundleName='" + fmt.Sprintf("%v", cvpdnbundleentry.Cvpdnbundlename) + "']"
}

func (cvpdnbundleentry *CISCOVPDNMGMTMIB_Cvpdnbundletable_Cvpdnbundleentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cvpdnbundleentry *CISCOVPDNMGMTMIB_Cvpdnbundletable_Cvpdnbundleentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cvpdnbundleentry *CISCOVPDNMGMTMIB_Cvpdnbundletable_Cvpdnbundleentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cvpdnBundleName"] = cvpdnbundleentry.Cvpdnbundlename
    leafs["cvpdnBundleLinkCount"] = cvpdnbundleentry.Cvpdnbundlelinkcount
    leafs["cvpdnBundleEndpointType"] = cvpdnbundleentry.Cvpdnbundleendpointtype
    leafs["cvpdnBundleEndpoint"] = cvpdnbundleentry.Cvpdnbundleendpoint
    leafs["cvpdnBundlePeerIpAddrType"] = cvpdnbundleentry.Cvpdnbundlepeeripaddrtype
    leafs["cvpdnBundlePeerIpAddr"] = cvpdnbundleentry.Cvpdnbundlepeeripaddr
    leafs["cvpdnBundleEndpointClass"] = cvpdnbundleentry.Cvpdnbundleendpointclass
    return leafs
}

func (cvpdnbundleentry *CISCOVPDNMGMTMIB_Cvpdnbundletable_Cvpdnbundleentry) GetBundleName() string { return "cisco_ios_xe" }

func (cvpdnbundleentry *CISCOVPDNMGMTMIB_Cvpdnbundletable_Cvpdnbundleentry) GetYangName() string { return "cvpdnBundleEntry" }

func (cvpdnbundleentry *CISCOVPDNMGMTMIB_Cvpdnbundletable_Cvpdnbundleentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvpdnbundleentry *CISCOVPDNMGMTMIB_Cvpdnbundletable_Cvpdnbundleentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvpdnbundleentry *CISCOVPDNMGMTMIB_Cvpdnbundletable_Cvpdnbundleentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvpdnbundleentry *CISCOVPDNMGMTMIB_Cvpdnbundletable_Cvpdnbundleentry) SetParent(parent types.Entity) { cvpdnbundleentry.parent = parent }

func (cvpdnbundleentry *CISCOVPDNMGMTMIB_Cvpdnbundletable_Cvpdnbundleentry) GetParent() types.Entity { return cvpdnbundleentry.parent }

func (cvpdnbundleentry *CISCOVPDNMGMTMIB_Cvpdnbundletable_Cvpdnbundleentry) GetParentYangName() string { return "cvpdnBundleTable" }

// CISCOVPDNMGMTMIB_Cvpdnbundletable_Cvpdnbundleentry_Cvpdnbundleendpointtype represents                  discriminator.
type CISCOVPDNMGMTMIB_Cvpdnbundletable_Cvpdnbundleentry_Cvpdnbundleendpointtype string

const (
    CISCOVPDNMGMTMIB_Cvpdnbundletable_Cvpdnbundleentry_Cvpdnbundleendpointtype_none CISCOVPDNMGMTMIB_Cvpdnbundletable_Cvpdnbundleentry_Cvpdnbundleendpointtype = "none"

    CISCOVPDNMGMTMIB_Cvpdnbundletable_Cvpdnbundleentry_Cvpdnbundleendpointtype_hostname CISCOVPDNMGMTMIB_Cvpdnbundletable_Cvpdnbundleentry_Cvpdnbundleendpointtype = "hostname"

    CISCOVPDNMGMTMIB_Cvpdnbundletable_Cvpdnbundleentry_Cvpdnbundleendpointtype_string CISCOVPDNMGMTMIB_Cvpdnbundletable_Cvpdnbundleentry_Cvpdnbundleendpointtype = "string"

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
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in this table represents a session that belongs to a VPDN tunnel
    // and to a multilink PPP bundle. The type is slice of
    // CISCOVPDNMGMTMIB_Cvpdnbundlechildtable_Cvpdnbundlechildentry.
    Cvpdnbundlechildentry []CISCOVPDNMGMTMIB_Cvpdnbundlechildtable_Cvpdnbundlechildentry
}

func (cvpdnbundlechildtable *CISCOVPDNMGMTMIB_Cvpdnbundlechildtable) GetFilter() yfilter.YFilter { return cvpdnbundlechildtable.YFilter }

func (cvpdnbundlechildtable *CISCOVPDNMGMTMIB_Cvpdnbundlechildtable) SetFilter(yf yfilter.YFilter) { cvpdnbundlechildtable.YFilter = yf }

func (cvpdnbundlechildtable *CISCOVPDNMGMTMIB_Cvpdnbundlechildtable) GetGoName(yname string) string {
    if yname == "cvpdnBundleChildEntry" { return "Cvpdnbundlechildentry" }
    return ""
}

func (cvpdnbundlechildtable *CISCOVPDNMGMTMIB_Cvpdnbundlechildtable) GetSegmentPath() string {
    return "cvpdnBundleChildTable"
}

func (cvpdnbundlechildtable *CISCOVPDNMGMTMIB_Cvpdnbundlechildtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cvpdnBundleChildEntry" {
        for _, c := range cvpdnbundlechildtable.Cvpdnbundlechildentry {
            if cvpdnbundlechildtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOVPDNMGMTMIB_Cvpdnbundlechildtable_Cvpdnbundlechildentry{}
        cvpdnbundlechildtable.Cvpdnbundlechildentry = append(cvpdnbundlechildtable.Cvpdnbundlechildentry, child)
        return &cvpdnbundlechildtable.Cvpdnbundlechildentry[len(cvpdnbundlechildtable.Cvpdnbundlechildentry)-1]
    }
    return nil
}

func (cvpdnbundlechildtable *CISCOVPDNMGMTMIB_Cvpdnbundlechildtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cvpdnbundlechildtable.Cvpdnbundlechildentry {
        children[cvpdnbundlechildtable.Cvpdnbundlechildentry[i].GetSegmentPath()] = &cvpdnbundlechildtable.Cvpdnbundlechildentry[i]
    }
    return children
}

func (cvpdnbundlechildtable *CISCOVPDNMGMTMIB_Cvpdnbundlechildtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cvpdnbundlechildtable *CISCOVPDNMGMTMIB_Cvpdnbundlechildtable) GetBundleName() string { return "cisco_ios_xe" }

func (cvpdnbundlechildtable *CISCOVPDNMGMTMIB_Cvpdnbundlechildtable) GetYangName() string { return "cvpdnBundleChildTable" }

func (cvpdnbundlechildtable *CISCOVPDNMGMTMIB_Cvpdnbundlechildtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvpdnbundlechildtable *CISCOVPDNMGMTMIB_Cvpdnbundlechildtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvpdnbundlechildtable *CISCOVPDNMGMTMIB_Cvpdnbundlechildtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvpdnbundlechildtable *CISCOVPDNMGMTMIB_Cvpdnbundlechildtable) SetParent(parent types.Entity) { cvpdnbundlechildtable.parent = parent }

func (cvpdnbundlechildtable *CISCOVPDNMGMTMIB_Cvpdnbundlechildtable) GetParent() types.Entity { return cvpdnbundlechildtable.parent }

func (cvpdnbundlechildtable *CISCOVPDNMGMTMIB_Cvpdnbundlechildtable) GetParentYangName() string { return "CISCO-VPDN-MGMT-MIB" }

// CISCOVPDNMGMTMIB_Cvpdnbundlechildtable_Cvpdnbundlechildentry
// An entry in this table represents a session that belongs to
// a VPDN tunnel and to a multilink PPP bundle.
type CISCOVPDNMGMTMIB_Cvpdnbundlechildtable_Cvpdnbundlechildentry struct {
    parent types.Entity
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

func (cvpdnbundlechildentry *CISCOVPDNMGMTMIB_Cvpdnbundlechildtable_Cvpdnbundlechildentry) GetFilter() yfilter.YFilter { return cvpdnbundlechildentry.YFilter }

func (cvpdnbundlechildentry *CISCOVPDNMGMTMIB_Cvpdnbundlechildtable_Cvpdnbundlechildentry) SetFilter(yf yfilter.YFilter) { cvpdnbundlechildentry.YFilter = yf }

func (cvpdnbundlechildentry *CISCOVPDNMGMTMIB_Cvpdnbundlechildtable_Cvpdnbundlechildentry) GetGoName(yname string) string {
    if yname == "cvpdnBundleName" { return "Cvpdnbundlename" }
    if yname == "cvpdnBundleChildTunnelType" { return "Cvpdnbundlechildtunneltype" }
    if yname == "cvpdnBundleChildTunnelId" { return "Cvpdnbundlechildtunnelid" }
    if yname == "cvpdnBundleChildSessionId" { return "Cvpdnbundlechildsessionid" }
    return ""
}

func (cvpdnbundlechildentry *CISCOVPDNMGMTMIB_Cvpdnbundlechildtable_Cvpdnbundlechildentry) GetSegmentPath() string {
    return "cvpdnBundleChildEntry" + "[cvpdnBundleName='" + fmt.Sprintf("%v", cvpdnbundlechildentry.Cvpdnbundlename) + "']" + "[cvpdnBundleChildTunnelType='" + fmt.Sprintf("%v", cvpdnbundlechildentry.Cvpdnbundlechildtunneltype) + "']" + "[cvpdnBundleChildTunnelId='" + fmt.Sprintf("%v", cvpdnbundlechildentry.Cvpdnbundlechildtunnelid) + "']" + "[cvpdnBundleChildSessionId='" + fmt.Sprintf("%v", cvpdnbundlechildentry.Cvpdnbundlechildsessionid) + "']"
}

func (cvpdnbundlechildentry *CISCOVPDNMGMTMIB_Cvpdnbundlechildtable_Cvpdnbundlechildentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cvpdnbundlechildentry *CISCOVPDNMGMTMIB_Cvpdnbundlechildtable_Cvpdnbundlechildentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cvpdnbundlechildentry *CISCOVPDNMGMTMIB_Cvpdnbundlechildtable_Cvpdnbundlechildentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cvpdnBundleName"] = cvpdnbundlechildentry.Cvpdnbundlename
    leafs["cvpdnBundleChildTunnelType"] = cvpdnbundlechildentry.Cvpdnbundlechildtunneltype
    leafs["cvpdnBundleChildTunnelId"] = cvpdnbundlechildentry.Cvpdnbundlechildtunnelid
    leafs["cvpdnBundleChildSessionId"] = cvpdnbundlechildentry.Cvpdnbundlechildsessionid
    return leafs
}

func (cvpdnbundlechildentry *CISCOVPDNMGMTMIB_Cvpdnbundlechildtable_Cvpdnbundlechildentry) GetBundleName() string { return "cisco_ios_xe" }

func (cvpdnbundlechildentry *CISCOVPDNMGMTMIB_Cvpdnbundlechildtable_Cvpdnbundlechildentry) GetYangName() string { return "cvpdnBundleChildEntry" }

func (cvpdnbundlechildentry *CISCOVPDNMGMTMIB_Cvpdnbundlechildtable_Cvpdnbundlechildentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvpdnbundlechildentry *CISCOVPDNMGMTMIB_Cvpdnbundlechildtable_Cvpdnbundlechildentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvpdnbundlechildentry *CISCOVPDNMGMTMIB_Cvpdnbundlechildtable_Cvpdnbundlechildentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvpdnbundlechildentry *CISCOVPDNMGMTMIB_Cvpdnbundlechildtable_Cvpdnbundlechildentry) SetParent(parent types.Entity) { cvpdnbundlechildentry.parent = parent }

func (cvpdnbundlechildentry *CISCOVPDNMGMTMIB_Cvpdnbundlechildtable_Cvpdnbundlechildentry) GetParent() types.Entity { return cvpdnbundlechildentry.parent }

func (cvpdnbundlechildentry *CISCOVPDNMGMTMIB_Cvpdnbundlechildtable_Cvpdnbundlechildentry) GetParentYangName() string { return "cvpdnBundleChildTable" }

