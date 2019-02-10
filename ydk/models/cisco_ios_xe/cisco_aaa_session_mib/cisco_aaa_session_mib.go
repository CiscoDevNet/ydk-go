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

    
    CasnActive CISCOAAASESSIONMIB_CasnActive

    
    CasnGeneral CISCOAAASESSIONMIB_CasnGeneral

    // This table contains entries for active AAA accounting sessions in the
    // system.
    CasnActiveTable CISCOAAASESSIONMIB_CasnActiveTable
}

func (cISCOAAASESSIONMIB *CISCOAAASESSIONMIB) GetEntityData() *types.CommonEntityData {
    cISCOAAASESSIONMIB.EntityData.YFilter = cISCOAAASESSIONMIB.YFilter
    cISCOAAASESSIONMIB.EntityData.YangName = "CISCO-AAA-SESSION-MIB"
    cISCOAAASESSIONMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOAAASESSIONMIB.EntityData.ParentYangName = "CISCO-AAA-SESSION-MIB"
    cISCOAAASESSIONMIB.EntityData.SegmentPath = "CISCO-AAA-SESSION-MIB:CISCO-AAA-SESSION-MIB"
    cISCOAAASESSIONMIB.EntityData.AbsolutePath = cISCOAAASESSIONMIB.EntityData.SegmentPath
    cISCOAAASESSIONMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOAAASESSIONMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOAAASESSIONMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOAAASESSIONMIB.EntityData.Children = types.NewOrderedMap()
    cISCOAAASESSIONMIB.EntityData.Children.Append("casnActive", types.YChild{"CasnActive", &cISCOAAASESSIONMIB.CasnActive})
    cISCOAAASESSIONMIB.EntityData.Children.Append("casnGeneral", types.YChild{"CasnGeneral", &cISCOAAASESSIONMIB.CasnGeneral})
    cISCOAAASESSIONMIB.EntityData.Children.Append("casnActiveTable", types.YChild{"CasnActiveTable", &cISCOAAASESSIONMIB.CasnActiveTable})
    cISCOAAASESSIONMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOAAASESSIONMIB.EntityData.YListKeys = []string {}

    return &(cISCOAAASESSIONMIB.EntityData)
}

// CISCOAAASESSIONMIB_CasnActive
type CISCOAAASESSIONMIB_CasnActive struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of entries currently in casnActiveTable. The type is interface{}
    // with range: 0..4294967295.
    CasnActiveTableEntries interface{}

    // Maximum number of entries present in casnActiveTable since last system
    // re-initialization.  This corresponds to the maximum value reported by
    // casnActiveTableEntries. The type is interface{} with range: 0..4294967295.
    CasnActiveTableHighWaterMark interface{}
}

func (casnActive *CISCOAAASESSIONMIB_CasnActive) GetEntityData() *types.CommonEntityData {
    casnActive.EntityData.YFilter = casnActive.YFilter
    casnActive.EntityData.YangName = "casnActive"
    casnActive.EntityData.BundleName = "cisco_ios_xe"
    casnActive.EntityData.ParentYangName = "CISCO-AAA-SESSION-MIB"
    casnActive.EntityData.SegmentPath = "casnActive"
    casnActive.EntityData.AbsolutePath = "CISCO-AAA-SESSION-MIB:CISCO-AAA-SESSION-MIB/" + casnActive.EntityData.SegmentPath
    casnActive.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    casnActive.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    casnActive.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    casnActive.EntityData.Children = types.NewOrderedMap()
    casnActive.EntityData.Leafs = types.NewOrderedMap()
    casnActive.EntityData.Leafs.Append("casnActiveTableEntries", types.YLeaf{"CasnActiveTableEntries", casnActive.CasnActiveTableEntries})
    casnActive.EntityData.Leafs.Append("casnActiveTableHighWaterMark", types.YLeaf{"CasnActiveTableHighWaterMark", casnActive.CasnActiveTableHighWaterMark})

    casnActive.EntityData.YListKeys = []string {}

    return &(casnActive.EntityData)
}

// CISCOAAASESSIONMIB_CasnGeneral
type CISCOAAASESSIONMIB_CasnGeneral struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Total number of sessions since last system re-initialization.  This value
    // includes all sessions currently in the casnActiveTable and all previous
    // sessions whether terminated via casnDisconnect or via other mechanisms. The
    // type is interface{} with range: 0..4294967295.
    CasnTotalSessions interface{}

    // Total number of sessions which have been disconnected using casnDisconnect
    // since last system re-initialization.  This value includes any sessions
    // still in the casnActiveTable with a casnDisconnect value of true(1) and all
    // previous sessions which terminated as a result of setting casnDisconnect.
    // The type is interface{} with range: 0..4294967295.
    CasnDisconnectedSessions interface{}
}

func (casnGeneral *CISCOAAASESSIONMIB_CasnGeneral) GetEntityData() *types.CommonEntityData {
    casnGeneral.EntityData.YFilter = casnGeneral.YFilter
    casnGeneral.EntityData.YangName = "casnGeneral"
    casnGeneral.EntityData.BundleName = "cisco_ios_xe"
    casnGeneral.EntityData.ParentYangName = "CISCO-AAA-SESSION-MIB"
    casnGeneral.EntityData.SegmentPath = "casnGeneral"
    casnGeneral.EntityData.AbsolutePath = "CISCO-AAA-SESSION-MIB:CISCO-AAA-SESSION-MIB/" + casnGeneral.EntityData.SegmentPath
    casnGeneral.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    casnGeneral.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    casnGeneral.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    casnGeneral.EntityData.Children = types.NewOrderedMap()
    casnGeneral.EntityData.Leafs = types.NewOrderedMap()
    casnGeneral.EntityData.Leafs.Append("casnTotalSessions", types.YLeaf{"CasnTotalSessions", casnGeneral.CasnTotalSessions})
    casnGeneral.EntityData.Leafs.Append("casnDisconnectedSessions", types.YLeaf{"CasnDisconnectedSessions", casnGeneral.CasnDisconnectedSessions})

    casnGeneral.EntityData.YListKeys = []string {}

    return &(casnGeneral.EntityData)
}

// CISCOAAASESSIONMIB_CasnActiveTable
// This table contains entries for active AAA accounting
// sessions in the system.
type CISCOAAASESSIONMIB_CasnActiveTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The information regarding a single accounting session.  Entries are created
    // when a new accounting session is begun.  Entries are removed when the
    // accounting session is ended.  Initiating termination of a session with the
    // object casnDisconnect will cause removal of the entry when the session
    // completes termination. The type is slice of
    // CISCOAAASESSIONMIB_CasnActiveTable_CasnActiveEntry.
    CasnActiveEntry []*CISCOAAASESSIONMIB_CasnActiveTable_CasnActiveEntry
}

func (casnActiveTable *CISCOAAASESSIONMIB_CasnActiveTable) GetEntityData() *types.CommonEntityData {
    casnActiveTable.EntityData.YFilter = casnActiveTable.YFilter
    casnActiveTable.EntityData.YangName = "casnActiveTable"
    casnActiveTable.EntityData.BundleName = "cisco_ios_xe"
    casnActiveTable.EntityData.ParentYangName = "CISCO-AAA-SESSION-MIB"
    casnActiveTable.EntityData.SegmentPath = "casnActiveTable"
    casnActiveTable.EntityData.AbsolutePath = "CISCO-AAA-SESSION-MIB:CISCO-AAA-SESSION-MIB/" + casnActiveTable.EntityData.SegmentPath
    casnActiveTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    casnActiveTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    casnActiveTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    casnActiveTable.EntityData.Children = types.NewOrderedMap()
    casnActiveTable.EntityData.Children.Append("casnActiveEntry", types.YChild{"CasnActiveEntry", nil})
    for i := range casnActiveTable.CasnActiveEntry {
        casnActiveTable.EntityData.Children.Append(types.GetSegmentPath(casnActiveTable.CasnActiveEntry[i]), types.YChild{"CasnActiveEntry", casnActiveTable.CasnActiveEntry[i]})
    }
    casnActiveTable.EntityData.Leafs = types.NewOrderedMap()

    casnActiveTable.EntityData.YListKeys = []string {}

    return &(casnActiveTable.EntityData)
}

// CISCOAAASESSIONMIB_CasnActiveTable_CasnActiveEntry
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
type CISCOAAASESSIONMIB_CasnActiveTable_CasnActiveEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. This is the session identification used by the
    // accounting protocol.  This value is unique to a session within the system,
    // even if multiple accounting protocols are in use.  The value of this object
    // corresponds to these accounting protocol attributes.    RADIUS:  attribute
    // 44, Acct-Session-Id    TACACS+: attribute 'task_id'. The type is
    // interface{} with range: 1..4294967295.
    CasnSessionId interface{}

    // The User login ID or zero length string if unavailable.  The value of this
    // object corresponds to these accounting protocol attributes.    RADIUS: 
    // attribute 1, User-Name    TACACS+: attribute 'user'. The type is string
    // with length: 0..255.
    CasnUserId interface{}

    // The IP address of the session or 0.0.0.0 if not applicable or unavailable. 
    // RADIUS:  attribute 8, Framed-IP-Address TACACS+: attribute 'addr'. The type
    // is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    CasnIpAddr interface{}

    // The elapsed time that this session has been idle.  This is the time since
    // the last user-level data has been received or transmitted. Protocol level
    // handshaking associated with the call is considered to be idle for this
    // object. The type is interface{} with range: 0..4294967295. Units are
    // seconds.
    CasnIdleTime interface{}

    // This object is used to terminate this session.  Setting the value to
    // true(1) will initiate termination of this session.  The entry will be
    // removed once the session has completed termination.  Once this object has
    // been set to true(1), the session termination process can not be cancelled
    // by setting the value false(2). The type is bool.
    CasnDisconnect interface{}

    // The value of this object is the entry index in the CISCO-CALL-TRACKER-MIB
    // cctActiveTable of the call corresponding to this accounting session.  Using
    // the value of this object to query the cctActiveTable will provide more
    // detailed data regarding the session represented by this casnActiveEntry.
    // The type is interface{} with range: 0..4294967295.
    CasnCallTrackerId interface{}

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
    CasnNasPort interface{}

    // The ifIndex of the Virtual Access Interface (VAI) that is associated with
    // the PPP session.  This interface may not be represented in the IF-MIB in
    // which case the value of this object will be zero. The type is interface{}
    // with range: 0..2147483647.
    CasnVaiIfIndex interface{}
}

func (casnActiveEntry *CISCOAAASESSIONMIB_CasnActiveTable_CasnActiveEntry) GetEntityData() *types.CommonEntityData {
    casnActiveEntry.EntityData.YFilter = casnActiveEntry.YFilter
    casnActiveEntry.EntityData.YangName = "casnActiveEntry"
    casnActiveEntry.EntityData.BundleName = "cisco_ios_xe"
    casnActiveEntry.EntityData.ParentYangName = "casnActiveTable"
    casnActiveEntry.EntityData.SegmentPath = "casnActiveEntry" + types.AddKeyToken(casnActiveEntry.CasnSessionId, "casnSessionId")
    casnActiveEntry.EntityData.AbsolutePath = "CISCO-AAA-SESSION-MIB:CISCO-AAA-SESSION-MIB/casnActiveTable/" + casnActiveEntry.EntityData.SegmentPath
    casnActiveEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    casnActiveEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    casnActiveEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    casnActiveEntry.EntityData.Children = types.NewOrderedMap()
    casnActiveEntry.EntityData.Leafs = types.NewOrderedMap()
    casnActiveEntry.EntityData.Leafs.Append("casnSessionId", types.YLeaf{"CasnSessionId", casnActiveEntry.CasnSessionId})
    casnActiveEntry.EntityData.Leafs.Append("casnUserId", types.YLeaf{"CasnUserId", casnActiveEntry.CasnUserId})
    casnActiveEntry.EntityData.Leafs.Append("casnIpAddr", types.YLeaf{"CasnIpAddr", casnActiveEntry.CasnIpAddr})
    casnActiveEntry.EntityData.Leafs.Append("casnIdleTime", types.YLeaf{"CasnIdleTime", casnActiveEntry.CasnIdleTime})
    casnActiveEntry.EntityData.Leafs.Append("casnDisconnect", types.YLeaf{"CasnDisconnect", casnActiveEntry.CasnDisconnect})
    casnActiveEntry.EntityData.Leafs.Append("casnCallTrackerId", types.YLeaf{"CasnCallTrackerId", casnActiveEntry.CasnCallTrackerId})
    casnActiveEntry.EntityData.Leafs.Append("casnNasPort", types.YLeaf{"CasnNasPort", casnActiveEntry.CasnNasPort})
    casnActiveEntry.EntityData.Leafs.Append("casnVaiIfIndex", types.YLeaf{"CasnVaiIfIndex", casnActiveEntry.CasnVaiIfIndex})

    casnActiveEntry.EntityData.YListKeys = []string {"CasnSessionId"}

    return &(casnActiveEntry.EntityData)
}

