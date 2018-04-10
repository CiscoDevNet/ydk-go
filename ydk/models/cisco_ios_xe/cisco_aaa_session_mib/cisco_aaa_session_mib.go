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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Casnactive CISCOAAASESSIONMIB_Casnactive

    
    Casngeneral CISCOAAASESSIONMIB_Casngeneral

    // This table contains entries for active AAA accounting sessions in the
    // system.
    Casnactivetable CISCOAAASESSIONMIB_Casnactivetable
}

func (cISCOAAASESSIONMIB *CISCOAAASESSIONMIB) GetEntityData() *types.CommonEntityData {
    cISCOAAASESSIONMIB.EntityData.YFilter = cISCOAAASESSIONMIB.YFilter
    cISCOAAASESSIONMIB.EntityData.YangName = "CISCO-AAA-SESSION-MIB"
    cISCOAAASESSIONMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOAAASESSIONMIB.EntityData.ParentYangName = "CISCO-AAA-SESSION-MIB"
    cISCOAAASESSIONMIB.EntityData.SegmentPath = "CISCO-AAA-SESSION-MIB:CISCO-AAA-SESSION-MIB"
    cISCOAAASESSIONMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOAAASESSIONMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOAAASESSIONMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOAAASESSIONMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOAAASESSIONMIB.EntityData.Children["casnActive"] = types.YChild{"Casnactive", &cISCOAAASESSIONMIB.Casnactive}
    cISCOAAASESSIONMIB.EntityData.Children["casnGeneral"] = types.YChild{"Casngeneral", &cISCOAAASESSIONMIB.Casngeneral}
    cISCOAAASESSIONMIB.EntityData.Children["casnActiveTable"] = types.YChild{"Casnactivetable", &cISCOAAASESSIONMIB.Casnactivetable}
    cISCOAAASESSIONMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOAAASESSIONMIB.EntityData)
}

// CISCOAAASESSIONMIB_Casnactive
type CISCOAAASESSIONMIB_Casnactive struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of entries currently in casnActiveTable. The type is interface{}
    // with range: 0..4294967295.
    Casnactivetableentries interface{}

    // Maximum number of entries present in casnActiveTable since last system
    // re-initialization.  This corresponds to the maximum value reported by
    // casnActiveTableEntries. The type is interface{} with range: 0..4294967295.
    Casnactivetablehighwatermark interface{}
}

func (casnactive *CISCOAAASESSIONMIB_Casnactive) GetEntityData() *types.CommonEntityData {
    casnactive.EntityData.YFilter = casnactive.YFilter
    casnactive.EntityData.YangName = "casnActive"
    casnactive.EntityData.BundleName = "cisco_ios_xe"
    casnactive.EntityData.ParentYangName = "CISCO-AAA-SESSION-MIB"
    casnactive.EntityData.SegmentPath = "casnActive"
    casnactive.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    casnactive.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    casnactive.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    casnactive.EntityData.Children = make(map[string]types.YChild)
    casnactive.EntityData.Leafs = make(map[string]types.YLeaf)
    casnactive.EntityData.Leafs["casnActiveTableEntries"] = types.YLeaf{"Casnactivetableentries", casnactive.Casnactivetableentries}
    casnactive.EntityData.Leafs["casnActiveTableHighWaterMark"] = types.YLeaf{"Casnactivetablehighwatermark", casnactive.Casnactivetablehighwatermark}
    return &(casnactive.EntityData)
}

// CISCOAAASESSIONMIB_Casngeneral
type CISCOAAASESSIONMIB_Casngeneral struct {
    EntityData types.CommonEntityData
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

func (casngeneral *CISCOAAASESSIONMIB_Casngeneral) GetEntityData() *types.CommonEntityData {
    casngeneral.EntityData.YFilter = casngeneral.YFilter
    casngeneral.EntityData.YangName = "casnGeneral"
    casngeneral.EntityData.BundleName = "cisco_ios_xe"
    casngeneral.EntityData.ParentYangName = "CISCO-AAA-SESSION-MIB"
    casngeneral.EntityData.SegmentPath = "casnGeneral"
    casngeneral.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    casngeneral.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    casngeneral.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    casngeneral.EntityData.Children = make(map[string]types.YChild)
    casngeneral.EntityData.Leafs = make(map[string]types.YLeaf)
    casngeneral.EntityData.Leafs["casnTotalSessions"] = types.YLeaf{"Casntotalsessions", casngeneral.Casntotalsessions}
    casngeneral.EntityData.Leafs["casnDisconnectedSessions"] = types.YLeaf{"Casndisconnectedsessions", casngeneral.Casndisconnectedsessions}
    return &(casngeneral.EntityData)
}

// CISCOAAASESSIONMIB_Casnactivetable
// This table contains entries for active AAA accounting
// sessions in the system.
type CISCOAAASESSIONMIB_Casnactivetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The information regarding a single accounting session.  Entries are created
    // when a new accounting session is begun.  Entries are removed when the
    // accounting session is ended.  Initiating termination of a session with the
    // object casnDisconnect will cause removal of the entry when the session
    // completes termination. The type is slice of
    // CISCOAAASESSIONMIB_Casnactivetable_Casnactiveentry.
    Casnactiveentry []CISCOAAASESSIONMIB_Casnactivetable_Casnactiveentry
}

func (casnactivetable *CISCOAAASESSIONMIB_Casnactivetable) GetEntityData() *types.CommonEntityData {
    casnactivetable.EntityData.YFilter = casnactivetable.YFilter
    casnactivetable.EntityData.YangName = "casnActiveTable"
    casnactivetable.EntityData.BundleName = "cisco_ios_xe"
    casnactivetable.EntityData.ParentYangName = "CISCO-AAA-SESSION-MIB"
    casnactivetable.EntityData.SegmentPath = "casnActiveTable"
    casnactivetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    casnactivetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    casnactivetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    casnactivetable.EntityData.Children = make(map[string]types.YChild)
    casnactivetable.EntityData.Children["casnActiveEntry"] = types.YChild{"Casnactiveentry", nil}
    for i := range casnactivetable.Casnactiveentry {
        casnactivetable.EntityData.Children[types.GetSegmentPath(&casnactivetable.Casnactiveentry[i])] = types.YChild{"Casnactiveentry", &casnactivetable.Casnactiveentry[i]}
    }
    casnactivetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(casnactivetable.EntityData)
}

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
    EntityData types.CommonEntityData
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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Casnnasport interface{}

    // The ifIndex of the Virtual Access Interface (VAI) that is associated with
    // the PPP session.  This interface may not be represented in the IF-MIB in
    // which case the value of this object will be zero. The type is interface{}
    // with range: 0..2147483647.
    Casnvaiifindex interface{}
}

func (casnactiveentry *CISCOAAASESSIONMIB_Casnactivetable_Casnactiveentry) GetEntityData() *types.CommonEntityData {
    casnactiveentry.EntityData.YFilter = casnactiveentry.YFilter
    casnactiveentry.EntityData.YangName = "casnActiveEntry"
    casnactiveentry.EntityData.BundleName = "cisco_ios_xe"
    casnactiveentry.EntityData.ParentYangName = "casnActiveTable"
    casnactiveentry.EntityData.SegmentPath = "casnActiveEntry" + "[casnSessionId='" + fmt.Sprintf("%v", casnactiveentry.Casnsessionid) + "']"
    casnactiveentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    casnactiveentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    casnactiveentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    casnactiveentry.EntityData.Children = make(map[string]types.YChild)
    casnactiveentry.EntityData.Leafs = make(map[string]types.YLeaf)
    casnactiveentry.EntityData.Leafs["casnSessionId"] = types.YLeaf{"Casnsessionid", casnactiveentry.Casnsessionid}
    casnactiveentry.EntityData.Leafs["casnUserId"] = types.YLeaf{"Casnuserid", casnactiveentry.Casnuserid}
    casnactiveentry.EntityData.Leafs["casnIpAddr"] = types.YLeaf{"Casnipaddr", casnactiveentry.Casnipaddr}
    casnactiveentry.EntityData.Leafs["casnIdleTime"] = types.YLeaf{"Casnidletime", casnactiveentry.Casnidletime}
    casnactiveentry.EntityData.Leafs["casnDisconnect"] = types.YLeaf{"Casndisconnect", casnactiveentry.Casndisconnect}
    casnactiveentry.EntityData.Leafs["casnCallTrackerId"] = types.YLeaf{"Casncalltrackerid", casnactiveentry.Casncalltrackerid}
    casnactiveentry.EntityData.Leafs["casnNasPort"] = types.YLeaf{"Casnnasport", casnactiveentry.Casnnasport}
    casnactiveentry.EntityData.Leafs["casnVaiIfIndex"] = types.YLeaf{"Casnvaiifindex", casnactiveentry.Casnvaiifindex}
    return &(casnactiveentry.EntityData)
}

