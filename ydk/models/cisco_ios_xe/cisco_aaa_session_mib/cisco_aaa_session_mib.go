// This MIB module provides data for accounting sessions
// based on Authentication, Authorization, Accounting
// (AAA) protocols.
// 
// 
// References:
//     RFC 2139 RADIUS Accounting
//     The TACACS+ Protocol Version 1.78, Internet Draft
package cisco_aaa_session_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_aaa_session_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-AAA-SESSION-MIB CISCO-AAA-SESSION-MIB}", reflect.TypeOf(CISCOAAASESSIONMIB{}))
    ydk.RegisterEntity("CISCO-AAA-SESSION-MIB:CISCO-AAA-SESSION-MIB", reflect.TypeOf(CISCOAAASESSIONMIB{}))
}

// CISCOAAASESSIONMIB
type CISCOAAASESSIONMIB struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Casnactive CISCOAAASESSIONMIB_Casnactive

    
    Casngeneral CISCOAAASESSIONMIB_Casngeneral

    // This table contains entries for active AAA accounting sessions in the
    // system.
    Casnactivetable CISCOAAASESSIONMIB_Casnactivetable
}

func (cISCOAAASESSIONMIB *CISCOAAASESSIONMIB) GetFilter() yfilter.YFilter { return cISCOAAASESSIONMIB.YFilter }

func (cISCOAAASESSIONMIB *CISCOAAASESSIONMIB) SetFilter(yf yfilter.YFilter) { cISCOAAASESSIONMIB.YFilter = yf }

func (cISCOAAASESSIONMIB *CISCOAAASESSIONMIB) GetGoName(yname string) string {
    if yname == "casnActive" { return "Casnactive" }
    if yname == "casnGeneral" { return "Casngeneral" }
    if yname == "casnActiveTable" { return "Casnactivetable" }
    return ""
}

func (cISCOAAASESSIONMIB *CISCOAAASESSIONMIB) GetSegmentPath() string {
    return "CISCO-AAA-SESSION-MIB:CISCO-AAA-SESSION-MIB"
}

func (cISCOAAASESSIONMIB *CISCOAAASESSIONMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "casnActive" {
        return &cISCOAAASESSIONMIB.Casnactive
    }
    if childYangName == "casnGeneral" {
        return &cISCOAAASESSIONMIB.Casngeneral
    }
    if childYangName == "casnActiveTable" {
        return &cISCOAAASESSIONMIB.Casnactivetable
    }
    return nil
}

func (cISCOAAASESSIONMIB *CISCOAAASESSIONMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["casnActive"] = &cISCOAAASESSIONMIB.Casnactive
    children["casnGeneral"] = &cISCOAAASESSIONMIB.Casngeneral
    children["casnActiveTable"] = &cISCOAAASESSIONMIB.Casnactivetable
    return children
}

func (cISCOAAASESSIONMIB *CISCOAAASESSIONMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOAAASESSIONMIB *CISCOAAASESSIONMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOAAASESSIONMIB *CISCOAAASESSIONMIB) GetYangName() string { return "CISCO-AAA-SESSION-MIB" }

func (cISCOAAASESSIONMIB *CISCOAAASESSIONMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOAAASESSIONMIB *CISCOAAASESSIONMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOAAASESSIONMIB *CISCOAAASESSIONMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOAAASESSIONMIB *CISCOAAASESSIONMIB) SetParent(parent types.Entity) { cISCOAAASESSIONMIB.parent = parent }

func (cISCOAAASESSIONMIB *CISCOAAASESSIONMIB) GetParent() types.Entity { return cISCOAAASESSIONMIB.parent }

func (cISCOAAASESSIONMIB *CISCOAAASESSIONMIB) GetParentYangName() string { return "CISCO-AAA-SESSION-MIB" }

// CISCOAAASESSIONMIB_Casnactive
type CISCOAAASESSIONMIB_Casnactive struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of entries currently in casnActiveTable. The type is interface{}
    // with range: 0..4294967295.
    Casnactivetableentries interface{}

    // Maximum number of entries present in casnActiveTable since last system
    // re-initialization.  This corresponds to the maximum value reported by
    // casnActiveTableEntries. The type is interface{} with range: 0..4294967295.
    Casnactivetablehighwatermark interface{}
}

func (casnactive *CISCOAAASESSIONMIB_Casnactive) GetFilter() yfilter.YFilter { return casnactive.YFilter }

func (casnactive *CISCOAAASESSIONMIB_Casnactive) SetFilter(yf yfilter.YFilter) { casnactive.YFilter = yf }

func (casnactive *CISCOAAASESSIONMIB_Casnactive) GetGoName(yname string) string {
    if yname == "casnActiveTableEntries" { return "Casnactivetableentries" }
    if yname == "casnActiveTableHighWaterMark" { return "Casnactivetablehighwatermark" }
    return ""
}

func (casnactive *CISCOAAASESSIONMIB_Casnactive) GetSegmentPath() string {
    return "casnActive"
}

func (casnactive *CISCOAAASESSIONMIB_Casnactive) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (casnactive *CISCOAAASESSIONMIB_Casnactive) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (casnactive *CISCOAAASESSIONMIB_Casnactive) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["casnActiveTableEntries"] = casnactive.Casnactivetableentries
    leafs["casnActiveTableHighWaterMark"] = casnactive.Casnactivetablehighwatermark
    return leafs
}

func (casnactive *CISCOAAASESSIONMIB_Casnactive) GetBundleName() string { return "cisco_ios_xe" }

func (casnactive *CISCOAAASESSIONMIB_Casnactive) GetYangName() string { return "casnActive" }

func (casnactive *CISCOAAASESSIONMIB_Casnactive) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (casnactive *CISCOAAASESSIONMIB_Casnactive) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (casnactive *CISCOAAASESSIONMIB_Casnactive) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (casnactive *CISCOAAASESSIONMIB_Casnactive) SetParent(parent types.Entity) { casnactive.parent = parent }

func (casnactive *CISCOAAASESSIONMIB_Casnactive) GetParent() types.Entity { return casnactive.parent }

func (casnactive *CISCOAAASESSIONMIB_Casnactive) GetParentYangName() string { return "CISCO-AAA-SESSION-MIB" }

// CISCOAAASESSIONMIB_Casngeneral
type CISCOAAASESSIONMIB_Casngeneral struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Total number of sessions since last system re-initialization.  This value
    // includes all sessions currently in the casnActiveTable and all previous
    // sessions whether terminated via casnDisconnect or via other mechanisms. The
    // type is interface{} with range: 0..4294967295.
    Casntotalsessions interface{}

    // Total number of sessions which have been disconnected using casnDisconnect
    // since last system re-initialization.  This value includes any sessions
    // still in the casnActiveTable with a casnDisconnect value of true(1) and all
    // previous sessions which terminated as a result of setting casnDisconnect.
    // The type is interface{} with range: 0..4294967295.
    Casndisconnectedsessions interface{}
}

func (casngeneral *CISCOAAASESSIONMIB_Casngeneral) GetFilter() yfilter.YFilter { return casngeneral.YFilter }

func (casngeneral *CISCOAAASESSIONMIB_Casngeneral) SetFilter(yf yfilter.YFilter) { casngeneral.YFilter = yf }

func (casngeneral *CISCOAAASESSIONMIB_Casngeneral) GetGoName(yname string) string {
    if yname == "casnTotalSessions" { return "Casntotalsessions" }
    if yname == "casnDisconnectedSessions" { return "Casndisconnectedsessions" }
    return ""
}

func (casngeneral *CISCOAAASESSIONMIB_Casngeneral) GetSegmentPath() string {
    return "casnGeneral"
}

func (casngeneral *CISCOAAASESSIONMIB_Casngeneral) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (casngeneral *CISCOAAASESSIONMIB_Casngeneral) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (casngeneral *CISCOAAASESSIONMIB_Casngeneral) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["casnTotalSessions"] = casngeneral.Casntotalsessions
    leafs["casnDisconnectedSessions"] = casngeneral.Casndisconnectedsessions
    return leafs
}

func (casngeneral *CISCOAAASESSIONMIB_Casngeneral) GetBundleName() string { return "cisco_ios_xe" }

func (casngeneral *CISCOAAASESSIONMIB_Casngeneral) GetYangName() string { return "casnGeneral" }

func (casngeneral *CISCOAAASESSIONMIB_Casngeneral) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (casngeneral *CISCOAAASESSIONMIB_Casngeneral) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (casngeneral *CISCOAAASESSIONMIB_Casngeneral) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (casngeneral *CISCOAAASESSIONMIB_Casngeneral) SetParent(parent types.Entity) { casngeneral.parent = parent }

func (casngeneral *CISCOAAASESSIONMIB_Casngeneral) GetParent() types.Entity { return casngeneral.parent }

func (casngeneral *CISCOAAASESSIONMIB_Casngeneral) GetParentYangName() string { return "CISCO-AAA-SESSION-MIB" }

// CISCOAAASESSIONMIB_Casnactivetable
// This table contains entries for active AAA accounting
// sessions in the system.
type CISCOAAASESSIONMIB_Casnactivetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The information regarding a single accounting session.  Entries are created
    // when a new accounting session is begun.  Entries are removed when the
    // accounting session is ended.  Initiating termination of a session with the
    // object casnDisconnect will cause removal of the entry when the session
    // completes termination. The type is slice of
    // CISCOAAASESSIONMIB_Casnactivetable_Casnactiveentry.
    Casnactiveentry []CISCOAAASESSIONMIB_Casnactivetable_Casnactiveentry
}

func (casnactivetable *CISCOAAASESSIONMIB_Casnactivetable) GetFilter() yfilter.YFilter { return casnactivetable.YFilter }

func (casnactivetable *CISCOAAASESSIONMIB_Casnactivetable) SetFilter(yf yfilter.YFilter) { casnactivetable.YFilter = yf }

func (casnactivetable *CISCOAAASESSIONMIB_Casnactivetable) GetGoName(yname string) string {
    if yname == "casnActiveEntry" { return "Casnactiveentry" }
    return ""
}

func (casnactivetable *CISCOAAASESSIONMIB_Casnactivetable) GetSegmentPath() string {
    return "casnActiveTable"
}

func (casnactivetable *CISCOAAASESSIONMIB_Casnactivetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "casnActiveEntry" {
        for _, c := range casnactivetable.Casnactiveentry {
            if casnactivetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOAAASESSIONMIB_Casnactivetable_Casnactiveentry{}
        casnactivetable.Casnactiveentry = append(casnactivetable.Casnactiveentry, child)
        return &casnactivetable.Casnactiveentry[len(casnactivetable.Casnactiveentry)-1]
    }
    return nil
}

func (casnactivetable *CISCOAAASESSIONMIB_Casnactivetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range casnactivetable.Casnactiveentry {
        children[casnactivetable.Casnactiveentry[i].GetSegmentPath()] = &casnactivetable.Casnactiveentry[i]
    }
    return children
}

func (casnactivetable *CISCOAAASESSIONMIB_Casnactivetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (casnactivetable *CISCOAAASESSIONMIB_Casnactivetable) GetBundleName() string { return "cisco_ios_xe" }

func (casnactivetable *CISCOAAASESSIONMIB_Casnactivetable) GetYangName() string { return "casnActiveTable" }

func (casnactivetable *CISCOAAASESSIONMIB_Casnactivetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (casnactivetable *CISCOAAASESSIONMIB_Casnactivetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (casnactivetable *CISCOAAASESSIONMIB_Casnactivetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (casnactivetable *CISCOAAASESSIONMIB_Casnactivetable) SetParent(parent types.Entity) { casnactivetable.parent = parent }

func (casnactivetable *CISCOAAASESSIONMIB_Casnactivetable) GetParent() types.Entity { return casnactivetable.parent }

func (casnactivetable *CISCOAAASESSIONMIB_Casnactivetable) GetParentYangName() string { return "CISCO-AAA-SESSION-MIB" }

// CISCOAAASESSIONMIB_Casnactivetable_Casnactiveentry
// The information regarding a single accounting session.
// 
// Entries are created when a new accounting session
// is begun.
// 
// Entries are removed when the accounting session
// is ended.
// 
// Initiating termination of a session with the object
// casnDisconnect will cause removal of the entry when
// the session completes termination.
type CISCOAAASESSIONMIB_Casnactivetable_Casnactiveentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. This is the session identification used by the
    // accounting protocol.  This value is unique to a session within the system,
    // even if multiple accounting protocols are in use.  The value of this object
    // corresponds to these accounting protocol attributes.    RADIUS:  attribute
    // 44, Acct-Session-Id    TACACS+: attribute 'task_id'. The type is
    // interface{} with range: 1..4294967295.
    Casnsessionid interface{}

    // The User login ID or zero length string if unavailable.  The value of this
    // object corresponds to these accounting protocol attributes.    RADIUS: 
    // attribute 1, User-Name    TACACS+: attribute 'user'. The type is string
    // with length: 0..255.
    Casnuserid interface{}

    // The IP address of the session or 0.0.0.0 if not applicable or unavailable. 
    // RADIUS:  attribute 8, Framed-IP-Address TACACS+: attribute 'addr'. The type
    // is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Casnipaddr interface{}

    // The elapsed time that this session has been idle.  This is the time since
    // the last user-level data has been received or transmitted. Protocol level
    // handshaking associated with the call is considered to be idle for this
    // object. The type is interface{} with range: 0..4294967295. Units are
    // seconds.
    Casnidletime interface{}

    // This object is used to terminate this session.  Setting the value to
    // true(1) will initiate termination of this session.  The entry will be
    // removed once the session has completed termination.  Once this object has
    // been set to true(1), the session termination process can not be cancelled
    // by setting the value false(2). The type is bool.
    Casndisconnect interface{}

    // The value of this object is the entry index in the CISCO-CALL-TRACKER-MIB
    // cctActiveTable of the call corresponding to this accounting session.  Using
    // the value of this object to query the cctActiveTable will provide more
    // detailed data regarding the session represented by this casnActiveEntry.
    // The type is interface{} with range: 0..4294967295.
    Casncalltrackerid interface{}

    // The value of this object identifies a particular conceptual row associated
    // with the session identified by casnSessionId.  The conceptual row that this
    // object points to represents a port that is used to transport a session.  If
    // the port transporting the session cannot be determined, the value of this
    // object will be zeroDotZero.  For example, suppose a session is established
    // using an ATM PVC.  If the ifIndex of the ATM interface is 7, and the 
    // VPI/VCI values of the PVC are 1, 100 respectively, then the value of this
    // object might be as follows:         casnNasPort.15 =
    // atmVclAdminStatus.7.1.100                    ^                      ^ ^  ^ 
    // |                      | |  |    casnSessionId --+                      | |
    // |          ifIndex -------------------------+ |  |        atmVclVpi
    // ---------------------------+  |        atmVclVci
    // ------------------------------+  where atmVclAdminStatus is the first
    // accessible object of the atmVclTable of the ATM-MIB. The type is string
    // with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    Casnnasport interface{}

    // The ifIndex of the Virtual Access Interface (VAI) that is associated with
    // the PPP session.  This interface may not be represented in the IF-MIB in
    // which case the value of this object will be zero. The type is interface{}
    // with range: 0..2147483647.
    Casnvaiifindex interface{}
}

func (casnactiveentry *CISCOAAASESSIONMIB_Casnactivetable_Casnactiveentry) GetFilter() yfilter.YFilter { return casnactiveentry.YFilter }

func (casnactiveentry *CISCOAAASESSIONMIB_Casnactivetable_Casnactiveentry) SetFilter(yf yfilter.YFilter) { casnactiveentry.YFilter = yf }

func (casnactiveentry *CISCOAAASESSIONMIB_Casnactivetable_Casnactiveentry) GetGoName(yname string) string {
    if yname == "casnSessionId" { return "Casnsessionid" }
    if yname == "casnUserId" { return "Casnuserid" }
    if yname == "casnIpAddr" { return "Casnipaddr" }
    if yname == "casnIdleTime" { return "Casnidletime" }
    if yname == "casnDisconnect" { return "Casndisconnect" }
    if yname == "casnCallTrackerId" { return "Casncalltrackerid" }
    if yname == "casnNasPort" { return "Casnnasport" }
    if yname == "casnVaiIfIndex" { return "Casnvaiifindex" }
    return ""
}

func (casnactiveentry *CISCOAAASESSIONMIB_Casnactivetable_Casnactiveentry) GetSegmentPath() string {
    return "casnActiveEntry" + "[casnSessionId='" + fmt.Sprintf("%v", casnactiveentry.Casnsessionid) + "']"
}

func (casnactiveentry *CISCOAAASESSIONMIB_Casnactivetable_Casnactiveentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (casnactiveentry *CISCOAAASESSIONMIB_Casnactivetable_Casnactiveentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (casnactiveentry *CISCOAAASESSIONMIB_Casnactivetable_Casnactiveentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["casnSessionId"] = casnactiveentry.Casnsessionid
    leafs["casnUserId"] = casnactiveentry.Casnuserid
    leafs["casnIpAddr"] = casnactiveentry.Casnipaddr
    leafs["casnIdleTime"] = casnactiveentry.Casnidletime
    leafs["casnDisconnect"] = casnactiveentry.Casndisconnect
    leafs["casnCallTrackerId"] = casnactiveentry.Casncalltrackerid
    leafs["casnNasPort"] = casnactiveentry.Casnnasport
    leafs["casnVaiIfIndex"] = casnactiveentry.Casnvaiifindex
    return leafs
}

func (casnactiveentry *CISCOAAASESSIONMIB_Casnactivetable_Casnactiveentry) GetBundleName() string { return "cisco_ios_xe" }

func (casnactiveentry *CISCOAAASESSIONMIB_Casnactivetable_Casnactiveentry) GetYangName() string { return "casnActiveEntry" }

func (casnactiveentry *CISCOAAASESSIONMIB_Casnactivetable_Casnactiveentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (casnactiveentry *CISCOAAASESSIONMIB_Casnactivetable_Casnactiveentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (casnactiveentry *CISCOAAASESSIONMIB_Casnactivetable_Casnactiveentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (casnactiveentry *CISCOAAASESSIONMIB_Casnactivetable_Casnactiveentry) SetParent(parent types.Entity) { casnactiveentry.parent = parent }

func (casnactiveentry *CISCOAAASESSIONMIB_Casnactivetable_Casnactiveentry) GetParent() types.Entity { return casnactiveentry.parent }

func (casnactiveentry *CISCOAAASESSIONMIB_Casnactivetable_Casnactiveentry) GetParentYangName() string { return "casnActiveTable" }

