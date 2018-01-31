// An experimental MIB module for MSDP Management.
package draft_msdp_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package draft_msdp_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:DRAFT-MSDP-MIB DRAFT-MSDP-MIB}", reflect.TypeOf(DRAFTMSDPMIB{}))
    ydk.RegisterEntity("DRAFT-MSDP-MIB:DRAFT-MSDP-MIB", reflect.TypeOf(DRAFTMSDPMIB{}))
}

// DRAFTMSDPMIB
type DRAFTMSDPMIB struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Msdp DRAFTMSDPMIB_Msdp

    // The (conceptual) table listing group ranges and MSDP peers used when
    // deciding where to send an SA Request message when required.  If SA Caching
    // is enabled, this table may be empty.
    Msdprequeststable DRAFTMSDPMIB_Msdprequeststable

    // The (conceptual) table listing the MSDP speaker's peers.
    Msdppeertable DRAFTMSDPMIB_Msdppeertable

    // The (conceptual) table listing the MSDP SA advertisements currently in the
    // MSDP speaker's cache.
    Msdpsacachetable DRAFTMSDPMIB_Msdpsacachetable
}

func (dRAFTMSDPMIB *DRAFTMSDPMIB) GetFilter() yfilter.YFilter { return dRAFTMSDPMIB.YFilter }

func (dRAFTMSDPMIB *DRAFTMSDPMIB) SetFilter(yf yfilter.YFilter) { dRAFTMSDPMIB.YFilter = yf }

func (dRAFTMSDPMIB *DRAFTMSDPMIB) GetGoName(yname string) string {
    if yname == "msdp" { return "Msdp" }
    if yname == "msdpRequestsTable" { return "Msdprequeststable" }
    if yname == "msdpPeerTable" { return "Msdppeertable" }
    if yname == "msdpSACacheTable" { return "Msdpsacachetable" }
    return ""
}

func (dRAFTMSDPMIB *DRAFTMSDPMIB) GetSegmentPath() string {
    return "DRAFT-MSDP-MIB:DRAFT-MSDP-MIB"
}

func (dRAFTMSDPMIB *DRAFTMSDPMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "msdp" {
        return &dRAFTMSDPMIB.Msdp
    }
    if childYangName == "msdpRequestsTable" {
        return &dRAFTMSDPMIB.Msdprequeststable
    }
    if childYangName == "msdpPeerTable" {
        return &dRAFTMSDPMIB.Msdppeertable
    }
    if childYangName == "msdpSACacheTable" {
        return &dRAFTMSDPMIB.Msdpsacachetable
    }
    return nil
}

func (dRAFTMSDPMIB *DRAFTMSDPMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["msdp"] = &dRAFTMSDPMIB.Msdp
    children["msdpRequestsTable"] = &dRAFTMSDPMIB.Msdprequeststable
    children["msdpPeerTable"] = &dRAFTMSDPMIB.Msdppeertable
    children["msdpSACacheTable"] = &dRAFTMSDPMIB.Msdpsacachetable
    return children
}

func (dRAFTMSDPMIB *DRAFTMSDPMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dRAFTMSDPMIB *DRAFTMSDPMIB) GetBundleName() string { return "cisco_ios_xe" }

func (dRAFTMSDPMIB *DRAFTMSDPMIB) GetYangName() string { return "DRAFT-MSDP-MIB" }

func (dRAFTMSDPMIB *DRAFTMSDPMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dRAFTMSDPMIB *DRAFTMSDPMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dRAFTMSDPMIB *DRAFTMSDPMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dRAFTMSDPMIB *DRAFTMSDPMIB) SetParent(parent types.Entity) { dRAFTMSDPMIB.parent = parent }

func (dRAFTMSDPMIB *DRAFTMSDPMIB) GetParent() types.Entity { return dRAFTMSDPMIB.parent }

func (dRAFTMSDPMIB *DRAFTMSDPMIB) GetParentYangName() string { return "DRAFT-MSDP-MIB" }

// DRAFTMSDPMIB_Msdp
type DRAFTMSDPMIB_Msdp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The state of MSDP on this MSDP speaker - globally enabled or disabled. The
    // type is bool.
    Msdpenabled interface{}

    // The lifetime given to SA cache entries when created or refreshed.  A value
    // of 0 means no SA caching is done by this MSDP speaker. The type is
    // interface{} with range: 0..4294967295.
    Msdpcachelifetime interface{}

    // The total number of entries in the SA Cache table. The type is interface{}
    // with range: 0..4294967295.
    Msdpnumsacacheentries interface{}

    // The number of seconds in the MSDP SA Hold-down period. The type is
    // interface{} with range: 1..2147483647. Units are seconds.
    Msdpsaholddownperiod interface{}
}

func (msdp *DRAFTMSDPMIB_Msdp) GetFilter() yfilter.YFilter { return msdp.YFilter }

func (msdp *DRAFTMSDPMIB_Msdp) SetFilter(yf yfilter.YFilter) { msdp.YFilter = yf }

func (msdp *DRAFTMSDPMIB_Msdp) GetGoName(yname string) string {
    if yname == "msdpEnabled" { return "Msdpenabled" }
    if yname == "msdpCacheLifetime" { return "Msdpcachelifetime" }
    if yname == "msdpNumSACacheEntries" { return "Msdpnumsacacheentries" }
    if yname == "msdpSAHoldDownPeriod" { return "Msdpsaholddownperiod" }
    return ""
}

func (msdp *DRAFTMSDPMIB_Msdp) GetSegmentPath() string {
    return "msdp"
}

func (msdp *DRAFTMSDPMIB_Msdp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (msdp *DRAFTMSDPMIB_Msdp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (msdp *DRAFTMSDPMIB_Msdp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["msdpEnabled"] = msdp.Msdpenabled
    leafs["msdpCacheLifetime"] = msdp.Msdpcachelifetime
    leafs["msdpNumSACacheEntries"] = msdp.Msdpnumsacacheentries
    leafs["msdpSAHoldDownPeriod"] = msdp.Msdpsaholddownperiod
    return leafs
}

func (msdp *DRAFTMSDPMIB_Msdp) GetBundleName() string { return "cisco_ios_xe" }

func (msdp *DRAFTMSDPMIB_Msdp) GetYangName() string { return "msdp" }

func (msdp *DRAFTMSDPMIB_Msdp) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (msdp *DRAFTMSDPMIB_Msdp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (msdp *DRAFTMSDPMIB_Msdp) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (msdp *DRAFTMSDPMIB_Msdp) SetParent(parent types.Entity) { msdp.parent = parent }

func (msdp *DRAFTMSDPMIB_Msdp) GetParent() types.Entity { return msdp.parent }

func (msdp *DRAFTMSDPMIB_Msdp) GetParentYangName() string { return "DRAFT-MSDP-MIB" }

// DRAFTMSDPMIB_Msdprequeststable
// The (conceptual) table listing group ranges and MSDP
// peers used when deciding where to send an SA Request
// message when required.  If SA Caching is enabled, this
// table may be empty.
type DRAFTMSDPMIB_Msdprequeststable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry (conceptual row) representing a group range used when deciding
    // where to send an SA Request message. The type is slice of
    // DRAFTMSDPMIB_Msdprequeststable_Msdprequestsentry.
    Msdprequestsentry []DRAFTMSDPMIB_Msdprequeststable_Msdprequestsentry
}

func (msdprequeststable *DRAFTMSDPMIB_Msdprequeststable) GetFilter() yfilter.YFilter { return msdprequeststable.YFilter }

func (msdprequeststable *DRAFTMSDPMIB_Msdprequeststable) SetFilter(yf yfilter.YFilter) { msdprequeststable.YFilter = yf }

func (msdprequeststable *DRAFTMSDPMIB_Msdprequeststable) GetGoName(yname string) string {
    if yname == "msdpRequestsEntry" { return "Msdprequestsentry" }
    return ""
}

func (msdprequeststable *DRAFTMSDPMIB_Msdprequeststable) GetSegmentPath() string {
    return "msdpRequestsTable"
}

func (msdprequeststable *DRAFTMSDPMIB_Msdprequeststable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "msdpRequestsEntry" {
        for _, c := range msdprequeststable.Msdprequestsentry {
            if msdprequeststable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DRAFTMSDPMIB_Msdprequeststable_Msdprequestsentry{}
        msdprequeststable.Msdprequestsentry = append(msdprequeststable.Msdprequestsentry, child)
        return &msdprequeststable.Msdprequestsentry[len(msdprequeststable.Msdprequestsentry)-1]
    }
    return nil
}

func (msdprequeststable *DRAFTMSDPMIB_Msdprequeststable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range msdprequeststable.Msdprequestsentry {
        children[msdprequeststable.Msdprequestsentry[i].GetSegmentPath()] = &msdprequeststable.Msdprequestsentry[i]
    }
    return children
}

func (msdprequeststable *DRAFTMSDPMIB_Msdprequeststable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (msdprequeststable *DRAFTMSDPMIB_Msdprequeststable) GetBundleName() string { return "cisco_ios_xe" }

func (msdprequeststable *DRAFTMSDPMIB_Msdprequeststable) GetYangName() string { return "msdpRequestsTable" }

func (msdprequeststable *DRAFTMSDPMIB_Msdprequeststable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (msdprequeststable *DRAFTMSDPMIB_Msdprequeststable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (msdprequeststable *DRAFTMSDPMIB_Msdprequeststable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (msdprequeststable *DRAFTMSDPMIB_Msdprequeststable) SetParent(parent types.Entity) { msdprequeststable.parent = parent }

func (msdprequeststable *DRAFTMSDPMIB_Msdprequeststable) GetParent() types.Entity { return msdprequeststable.parent }

func (msdprequeststable *DRAFTMSDPMIB_Msdprequeststable) GetParentYangName() string { return "DRAFT-MSDP-MIB" }

// DRAFTMSDPMIB_Msdprequeststable_Msdprequestsentry
// An entry (conceptual row) representing a group range
// used when deciding where to send an SA Request
// message.
type DRAFTMSDPMIB_Msdprequeststable_Msdprequestsentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The group address that, when combined with the
    // mask in this entry, represents the group range for which this peer will
    // service MSDP SA Requests. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Msdprequestsgroupaddress interface{}

    // This attribute is a key. The mask that, when combined with the group
    // address in this entry, represents the group range for which this peer will
    // service MSDP SA Requests. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Msdprequestsgroupmask interface{}

    // The peer to which MSDP SA Requests for groups matching this entry's group
    // range will be sent.  Must match the INDEX of a row in the msdpPeerTable.
    // The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Msdprequestspeer interface{}

    // The status of this row, by which new rows may be added to the table. The
    // type is RowStatus.
    Msdprequestsstatus interface{}
}

func (msdprequestsentry *DRAFTMSDPMIB_Msdprequeststable_Msdprequestsentry) GetFilter() yfilter.YFilter { return msdprequestsentry.YFilter }

func (msdprequestsentry *DRAFTMSDPMIB_Msdprequeststable_Msdprequestsentry) SetFilter(yf yfilter.YFilter) { msdprequestsentry.YFilter = yf }

func (msdprequestsentry *DRAFTMSDPMIB_Msdprequeststable_Msdprequestsentry) GetGoName(yname string) string {
    if yname == "msdpRequestsGroupAddress" { return "Msdprequestsgroupaddress" }
    if yname == "msdpRequestsGroupMask" { return "Msdprequestsgroupmask" }
    if yname == "msdpRequestsPeer" { return "Msdprequestspeer" }
    if yname == "msdpRequestsStatus" { return "Msdprequestsstatus" }
    return ""
}

func (msdprequestsentry *DRAFTMSDPMIB_Msdprequeststable_Msdprequestsentry) GetSegmentPath() string {
    return "msdpRequestsEntry" + "[msdpRequestsGroupAddress='" + fmt.Sprintf("%v", msdprequestsentry.Msdprequestsgroupaddress) + "']" + "[msdpRequestsGroupMask='" + fmt.Sprintf("%v", msdprequestsentry.Msdprequestsgroupmask) + "']"
}

func (msdprequestsentry *DRAFTMSDPMIB_Msdprequeststable_Msdprequestsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (msdprequestsentry *DRAFTMSDPMIB_Msdprequeststable_Msdprequestsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (msdprequestsentry *DRAFTMSDPMIB_Msdprequeststable_Msdprequestsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["msdpRequestsGroupAddress"] = msdprequestsentry.Msdprequestsgroupaddress
    leafs["msdpRequestsGroupMask"] = msdprequestsentry.Msdprequestsgroupmask
    leafs["msdpRequestsPeer"] = msdprequestsentry.Msdprequestspeer
    leafs["msdpRequestsStatus"] = msdprequestsentry.Msdprequestsstatus
    return leafs
}

func (msdprequestsentry *DRAFTMSDPMIB_Msdprequeststable_Msdprequestsentry) GetBundleName() string { return "cisco_ios_xe" }

func (msdprequestsentry *DRAFTMSDPMIB_Msdprequeststable_Msdprequestsentry) GetYangName() string { return "msdpRequestsEntry" }

func (msdprequestsentry *DRAFTMSDPMIB_Msdprequeststable_Msdprequestsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (msdprequestsentry *DRAFTMSDPMIB_Msdprequeststable_Msdprequestsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (msdprequestsentry *DRAFTMSDPMIB_Msdprequeststable_Msdprequestsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (msdprequestsentry *DRAFTMSDPMIB_Msdprequeststable_Msdprequestsentry) SetParent(parent types.Entity) { msdprequestsentry.parent = parent }

func (msdprequestsentry *DRAFTMSDPMIB_Msdprequeststable_Msdprequestsentry) GetParent() types.Entity { return msdprequestsentry.parent }

func (msdprequestsentry *DRAFTMSDPMIB_Msdprequeststable_Msdprequestsentry) GetParentYangName() string { return "msdpRequestsTable" }

// DRAFTMSDPMIB_Msdppeertable
// The (conceptual) table listing the MSDP speaker's
// peers.
type DRAFTMSDPMIB_Msdppeertable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry (conceptual row) representing an MSDP peer. The type is slice of
    // DRAFTMSDPMIB_Msdppeertable_Msdppeerentry.
    Msdppeerentry []DRAFTMSDPMIB_Msdppeertable_Msdppeerentry
}

func (msdppeertable *DRAFTMSDPMIB_Msdppeertable) GetFilter() yfilter.YFilter { return msdppeertable.YFilter }

func (msdppeertable *DRAFTMSDPMIB_Msdppeertable) SetFilter(yf yfilter.YFilter) { msdppeertable.YFilter = yf }

func (msdppeertable *DRAFTMSDPMIB_Msdppeertable) GetGoName(yname string) string {
    if yname == "msdpPeerEntry" { return "Msdppeerentry" }
    return ""
}

func (msdppeertable *DRAFTMSDPMIB_Msdppeertable) GetSegmentPath() string {
    return "msdpPeerTable"
}

func (msdppeertable *DRAFTMSDPMIB_Msdppeertable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "msdpPeerEntry" {
        for _, c := range msdppeertable.Msdppeerentry {
            if msdppeertable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DRAFTMSDPMIB_Msdppeertable_Msdppeerentry{}
        msdppeertable.Msdppeerentry = append(msdppeertable.Msdppeerentry, child)
        return &msdppeertable.Msdppeerentry[len(msdppeertable.Msdppeerentry)-1]
    }
    return nil
}

func (msdppeertable *DRAFTMSDPMIB_Msdppeertable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range msdppeertable.Msdppeerentry {
        children[msdppeertable.Msdppeerentry[i].GetSegmentPath()] = &msdppeertable.Msdppeerentry[i]
    }
    return children
}

func (msdppeertable *DRAFTMSDPMIB_Msdppeertable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (msdppeertable *DRAFTMSDPMIB_Msdppeertable) GetBundleName() string { return "cisco_ios_xe" }

func (msdppeertable *DRAFTMSDPMIB_Msdppeertable) GetYangName() string { return "msdpPeerTable" }

func (msdppeertable *DRAFTMSDPMIB_Msdppeertable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (msdppeertable *DRAFTMSDPMIB_Msdppeertable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (msdppeertable *DRAFTMSDPMIB_Msdppeertable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (msdppeertable *DRAFTMSDPMIB_Msdppeertable) SetParent(parent types.Entity) { msdppeertable.parent = parent }

func (msdppeertable *DRAFTMSDPMIB_Msdppeertable) GetParent() types.Entity { return msdppeertable.parent }

func (msdppeertable *DRAFTMSDPMIB_Msdppeertable) GetParentYangName() string { return "DRAFT-MSDP-MIB" }

// DRAFTMSDPMIB_Msdppeertable_Msdppeerentry
// An entry (conceptual row) representing an MSDP peer.
type DRAFTMSDPMIB_Msdppeertable_Msdppeerentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The address of the remote MSDP peer. The type is
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Msdppeerremoteaddress interface{}

    // The state of the MSDP TCP connection with this peer. The type is
    // Msdppeerstate.
    Msdppeerstate interface{}

    // The number of RPF failures on SA messages received from this peer. The type
    // is interface{} with range: 0..4294967295.
    Msdppeerrpffailures interface{}

    // The number of MSDP SA messages received on this connection.  This object
    // should be initialized to zero when the connection is established. The type
    // is interface{} with range: 0..4294967295.
    Msdppeerinsas interface{}

    // The number of MSDP SA messages transmitted on this connection.  This object
    // should be initialized to zero when the connection is established. The type
    // is interface{} with range: 0..4294967295.
    Msdppeeroutsas interface{}

    // The number of MSDP SA-Request messages received on this connection.  This
    // object should be initialized to zero when the connection is established.
    // The type is interface{} with range: 0..4294967295.
    Msdppeerinsarequests interface{}

    // The number of MSDP SA-Request messages transmitted on this connection. 
    // This object should be initialized to zero when the connection is
    // established. The type is interface{} with range: 0..4294967295.
    Msdppeeroutsarequests interface{}

    // The number of MSDP SA-Response messages received on this connection.  This
    // object should be initialized to zero when the connection is established.
    // The type is interface{} with range: 0..4294967295.
    Msdppeerinsaresponses interface{}

    // The number of MSDP SA Response messages transmitted on this TCP connection.
    // This object should be initialized to zero when the connection is
    // established. The type is interface{} with range: 0..4294967295.
    Msdppeeroutsaresponses interface{}

    // The total number of MSDP messages received on this TCP connection.  This
    // object should be initialized to zero when the connection is established.
    // The type is interface{} with range: 0..4294967295.
    Msdppeerincontrolmessages interface{}

    // The total number of MSDP messages transmitted on this TCP connection.  This
    // object should be initialized to zero when the connection is established.
    // The type is interface{} with range: 0..4294967295.
    Msdppeeroutcontrolmessages interface{}

    // The total number of encapsulated data packets received from this peer. 
    // This object should be initialized to zero when the connection is
    // established. The type is interface{} with range: 0..4294967295.
    Msdppeerindatapackets interface{}

    // The total number of encapsulated data packets sent to this peer.  This
    // object should be initialized to zero when the connection is established.
    // The type is interface{} with range: 0..4294967295.
    Msdppeeroutdatapackets interface{}

    // The total number of times the MSDP FSM transitioned into the established
    // state. The type is interface{} with range: 0..4294967295.
    Msdppeerfsmestablishedtransitions interface{}

    // This timer indicates how long (in seconds) this peer has been in the
    // Established state or how long since this peer was last in the Established
    // state.  It is set to zero when a new peer is configured or the MSDP speaker
    // is booted. The type is interface{} with range: 0..4294967295. Units are
    // seconds.
    Msdppeerfsmestablishedtime interface{}

    // Elapsed time in seconds since the last MSDP message was received from the
    // peer.  Each time msdpPeerInControlMessages is incremented, the value of
    // this object is set to zero (0). The type is interface{} with range:
    // 0..4294967295. Units are seconds.
    Msdppeerinmessageelapsedtime interface{}

    // The local IP address of this entry's MSDP connection. The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Msdppeerlocaladdress interface{}

    // Time interval in seconds for the MinSAAdvertisementInterval MSDP timer. The
    // type is interface{} with range: 1..2147483647. Units are seconds.
    Msdppeersaadvperiod interface{}

    // Time interval in seconds for the ConnectRetry timer. The type is
    // interface{} with range: 1..65535. Units are seconds.
    Msdppeerconnectretryinterval interface{}

    // Time interval in seconds for the Hold Timer configured for this MSDP
    // speaker with this peer. The type is interface{} with range: 0..None |
    // 3..65535. Units are seconds.
    Msdppeerholdtimeconfigured interface{}

    // Time interval in seconds for the KeepAlive timer configured for this MSDP
    // speaker with this peer.  A reasonable maximum value for this timer would be
    // configured to be one third of that of msdpPeerHoldTimeConfigured.  If the
    // value of this object is zero (0), no periodic KEEPALIVE messages are sent
    // to the peer after the MSDP connection has been established. The type is
    // interface{} with range: 0..21845. Units are seconds.
    Msdppeerkeepaliveconfigured interface{}

    // The minimum TTL a packet is required to have before it may be forwarded
    // using SA encapsulation to this peer. The type is interface{} with range:
    // 0..255.
    Msdppeerdatattl interface{}

    // This object indicates whether or not to process MSDP SA Request messages
    // from this peer.  If True(1), MSDP SA Request messages from this peer are
    // processed and replied to (if appropriate) with SA Response messages. If
    // False(2), MSDP SA Request messages from this peer are silently ignored.  It
    // defaults to False when msdpCacheLifetime is 0 and True when
    // msdpCacheLifetime is non-0. The type is bool.
    Msdppeerprocessrequestsfrom interface{}

    // The RowStatus object by which peers can be added and deleted.  A transition
    // to 'active' will cause the MSDP Start Event to be generated.  A transition
    // out of the 'active' state will cause the MSDP Stop Event to be generated. 
    // Care should be used in providing write access to this object without
    // adequate authentication. The type is RowStatus.
    Msdppeerstatus interface{}

    // The remote port for the TCP connection between the MSDP peers. The type is
    // interface{} with range: 0..65535.
    Msdppeerremoteport interface{}

    // The local port for the TCP connection between the MSDP peers. The type is
    // interface{} with range: 0..65535.
    Msdppeerlocalport interface{}

    // The status of the encapsulation negotiation state machine. The type is
    // Msdppeerencapsulationstate.
    Msdppeerencapsulationstate interface{}

    // The encapsulation in use when encapsulating data in SA messages to this
    // peer. The type is Msdppeerencapsulationtype.
    Msdppeerencapsulationtype interface{}

    // The number of times the state machine has transitioned from inactive to
    // connecting. The type is interface{} with range: 0..4294967295.
    Msdppeerconnectionattempts interface{}

    // The number of MSDP Notification messages received on this connection.  This
    // object should be initialized to zero when the connection is established.
    // The type is interface{} with range: 0..4294967295.
    Msdppeerinnotifications interface{}

    // The number of MSDP Notification messages transmitted on this connection. 
    // This object should be initialized to zero when the connection is
    // established. The type is interface{} with range: 0..4294967295.
    Msdppeeroutnotifications interface{}

    // The last error code and subcode seen by this peer on this connection.  If
    // no error has occurred, this field is zero.  Otherwise, the first byte of
    // this two byte OCTET STRING contains the error code, and the second byte
    // contains the subcode. The type is string with length: 2.
    Msdppeerlasterror interface{}
}

func (msdppeerentry *DRAFTMSDPMIB_Msdppeertable_Msdppeerentry) GetFilter() yfilter.YFilter { return msdppeerentry.YFilter }

func (msdppeerentry *DRAFTMSDPMIB_Msdppeertable_Msdppeerentry) SetFilter(yf yfilter.YFilter) { msdppeerentry.YFilter = yf }

func (msdppeerentry *DRAFTMSDPMIB_Msdppeertable_Msdppeerentry) GetGoName(yname string) string {
    if yname == "msdpPeerRemoteAddress" { return "Msdppeerremoteaddress" }
    if yname == "msdpPeerState" { return "Msdppeerstate" }
    if yname == "msdpPeerRPFFailures" { return "Msdppeerrpffailures" }
    if yname == "msdpPeerInSAs" { return "Msdppeerinsas" }
    if yname == "msdpPeerOutSAs" { return "Msdppeeroutsas" }
    if yname == "msdpPeerInSARequests" { return "Msdppeerinsarequests" }
    if yname == "msdpPeerOutSARequests" { return "Msdppeeroutsarequests" }
    if yname == "msdpPeerInSAResponses" { return "Msdppeerinsaresponses" }
    if yname == "msdpPeerOutSAResponses" { return "Msdppeeroutsaresponses" }
    if yname == "msdpPeerInControlMessages" { return "Msdppeerincontrolmessages" }
    if yname == "msdpPeerOutControlMessages" { return "Msdppeeroutcontrolmessages" }
    if yname == "msdpPeerInDataPackets" { return "Msdppeerindatapackets" }
    if yname == "msdpPeerOutDataPackets" { return "Msdppeeroutdatapackets" }
    if yname == "msdpPeerFsmEstablishedTransitions" { return "Msdppeerfsmestablishedtransitions" }
    if yname == "msdpPeerFsmEstablishedTime" { return "Msdppeerfsmestablishedtime" }
    if yname == "msdpPeerInMessageElapsedTime" { return "Msdppeerinmessageelapsedtime" }
    if yname == "msdpPeerLocalAddress" { return "Msdppeerlocaladdress" }
    if yname == "msdpPeerSAAdvPeriod" { return "Msdppeersaadvperiod" }
    if yname == "msdpPeerConnectRetryInterval" { return "Msdppeerconnectretryinterval" }
    if yname == "msdpPeerHoldTimeConfigured" { return "Msdppeerholdtimeconfigured" }
    if yname == "msdpPeerKeepAliveConfigured" { return "Msdppeerkeepaliveconfigured" }
    if yname == "msdpPeerDataTtl" { return "Msdppeerdatattl" }
    if yname == "msdpPeerProcessRequestsFrom" { return "Msdppeerprocessrequestsfrom" }
    if yname == "msdpPeerStatus" { return "Msdppeerstatus" }
    if yname == "msdpPeerRemotePort" { return "Msdppeerremoteport" }
    if yname == "msdpPeerLocalPort" { return "Msdppeerlocalport" }
    if yname == "msdpPeerEncapsulationState" { return "Msdppeerencapsulationstate" }
    if yname == "msdpPeerEncapsulationType" { return "Msdppeerencapsulationtype" }
    if yname == "msdpPeerConnectionAttempts" { return "Msdppeerconnectionattempts" }
    if yname == "msdpPeerInNotifications" { return "Msdppeerinnotifications" }
    if yname == "msdpPeerOutNotifications" { return "Msdppeeroutnotifications" }
    if yname == "msdpPeerLastError" { return "Msdppeerlasterror" }
    return ""
}

func (msdppeerentry *DRAFTMSDPMIB_Msdppeertable_Msdppeerentry) GetSegmentPath() string {
    return "msdpPeerEntry" + "[msdpPeerRemoteAddress='" + fmt.Sprintf("%v", msdppeerentry.Msdppeerremoteaddress) + "']"
}

func (msdppeerentry *DRAFTMSDPMIB_Msdppeertable_Msdppeerentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (msdppeerentry *DRAFTMSDPMIB_Msdppeertable_Msdppeerentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (msdppeerentry *DRAFTMSDPMIB_Msdppeertable_Msdppeerentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["msdpPeerRemoteAddress"] = msdppeerentry.Msdppeerremoteaddress
    leafs["msdpPeerState"] = msdppeerentry.Msdppeerstate
    leafs["msdpPeerRPFFailures"] = msdppeerentry.Msdppeerrpffailures
    leafs["msdpPeerInSAs"] = msdppeerentry.Msdppeerinsas
    leafs["msdpPeerOutSAs"] = msdppeerentry.Msdppeeroutsas
    leafs["msdpPeerInSARequests"] = msdppeerentry.Msdppeerinsarequests
    leafs["msdpPeerOutSARequests"] = msdppeerentry.Msdppeeroutsarequests
    leafs["msdpPeerInSAResponses"] = msdppeerentry.Msdppeerinsaresponses
    leafs["msdpPeerOutSAResponses"] = msdppeerentry.Msdppeeroutsaresponses
    leafs["msdpPeerInControlMessages"] = msdppeerentry.Msdppeerincontrolmessages
    leafs["msdpPeerOutControlMessages"] = msdppeerentry.Msdppeeroutcontrolmessages
    leafs["msdpPeerInDataPackets"] = msdppeerentry.Msdppeerindatapackets
    leafs["msdpPeerOutDataPackets"] = msdppeerentry.Msdppeeroutdatapackets
    leafs["msdpPeerFsmEstablishedTransitions"] = msdppeerentry.Msdppeerfsmestablishedtransitions
    leafs["msdpPeerFsmEstablishedTime"] = msdppeerentry.Msdppeerfsmestablishedtime
    leafs["msdpPeerInMessageElapsedTime"] = msdppeerentry.Msdppeerinmessageelapsedtime
    leafs["msdpPeerLocalAddress"] = msdppeerentry.Msdppeerlocaladdress
    leafs["msdpPeerSAAdvPeriod"] = msdppeerentry.Msdppeersaadvperiod
    leafs["msdpPeerConnectRetryInterval"] = msdppeerentry.Msdppeerconnectretryinterval
    leafs["msdpPeerHoldTimeConfigured"] = msdppeerentry.Msdppeerholdtimeconfigured
    leafs["msdpPeerKeepAliveConfigured"] = msdppeerentry.Msdppeerkeepaliveconfigured
    leafs["msdpPeerDataTtl"] = msdppeerentry.Msdppeerdatattl
    leafs["msdpPeerProcessRequestsFrom"] = msdppeerentry.Msdppeerprocessrequestsfrom
    leafs["msdpPeerStatus"] = msdppeerentry.Msdppeerstatus
    leafs["msdpPeerRemotePort"] = msdppeerentry.Msdppeerremoteport
    leafs["msdpPeerLocalPort"] = msdppeerentry.Msdppeerlocalport
    leafs["msdpPeerEncapsulationState"] = msdppeerentry.Msdppeerencapsulationstate
    leafs["msdpPeerEncapsulationType"] = msdppeerentry.Msdppeerencapsulationtype
    leafs["msdpPeerConnectionAttempts"] = msdppeerentry.Msdppeerconnectionattempts
    leafs["msdpPeerInNotifications"] = msdppeerentry.Msdppeerinnotifications
    leafs["msdpPeerOutNotifications"] = msdppeerentry.Msdppeeroutnotifications
    leafs["msdpPeerLastError"] = msdppeerentry.Msdppeerlasterror
    return leafs
}

func (msdppeerentry *DRAFTMSDPMIB_Msdppeertable_Msdppeerentry) GetBundleName() string { return "cisco_ios_xe" }

func (msdppeerentry *DRAFTMSDPMIB_Msdppeertable_Msdppeerentry) GetYangName() string { return "msdpPeerEntry" }

func (msdppeerentry *DRAFTMSDPMIB_Msdppeertable_Msdppeerentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (msdppeerentry *DRAFTMSDPMIB_Msdppeertable_Msdppeerentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (msdppeerentry *DRAFTMSDPMIB_Msdppeertable_Msdppeerentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (msdppeerentry *DRAFTMSDPMIB_Msdppeertable_Msdppeerentry) SetParent(parent types.Entity) { msdppeerentry.parent = parent }

func (msdppeerentry *DRAFTMSDPMIB_Msdppeertable_Msdppeerentry) GetParent() types.Entity { return msdppeerentry.parent }

func (msdppeerentry *DRAFTMSDPMIB_Msdppeertable_Msdppeerentry) GetParentYangName() string { return "msdpPeerTable" }

// DRAFTMSDPMIB_Msdppeertable_Msdppeerentry_Msdppeerencapsulationstate represents machine.
type DRAFTMSDPMIB_Msdppeertable_Msdppeerentry_Msdppeerencapsulationstate string

const (
    DRAFTMSDPMIB_Msdppeertable_Msdppeerentry_Msdppeerencapsulationstate_default_ DRAFTMSDPMIB_Msdppeertable_Msdppeerentry_Msdppeerencapsulationstate = "default"

    DRAFTMSDPMIB_Msdppeertable_Msdppeerentry_Msdppeerencapsulationstate_received DRAFTMSDPMIB_Msdppeertable_Msdppeerentry_Msdppeerencapsulationstate = "received"

    DRAFTMSDPMIB_Msdppeertable_Msdppeerentry_Msdppeerencapsulationstate_advertising DRAFTMSDPMIB_Msdppeertable_Msdppeerentry_Msdppeerencapsulationstate = "advertising"

    DRAFTMSDPMIB_Msdppeertable_Msdppeerentry_Msdppeerencapsulationstate_sent DRAFTMSDPMIB_Msdppeertable_Msdppeerentry_Msdppeerencapsulationstate = "sent"

    DRAFTMSDPMIB_Msdppeertable_Msdppeerentry_Msdppeerencapsulationstate_agreed DRAFTMSDPMIB_Msdppeertable_Msdppeerentry_Msdppeerencapsulationstate = "agreed"

    DRAFTMSDPMIB_Msdppeertable_Msdppeerentry_Msdppeerencapsulationstate_failed DRAFTMSDPMIB_Msdppeertable_Msdppeerentry_Msdppeerencapsulationstate = "failed"
)

// DRAFTMSDPMIB_Msdppeertable_Msdppeerentry_Msdppeerencapsulationtype represents SA messages to this peer.
type DRAFTMSDPMIB_Msdppeertable_Msdppeerentry_Msdppeerencapsulationtype string

const (
    DRAFTMSDPMIB_Msdppeertable_Msdppeerentry_Msdppeerencapsulationtype_tcp DRAFTMSDPMIB_Msdppeertable_Msdppeerentry_Msdppeerencapsulationtype = "tcp"

    DRAFTMSDPMIB_Msdppeertable_Msdppeerentry_Msdppeerencapsulationtype_udp DRAFTMSDPMIB_Msdppeertable_Msdppeerentry_Msdppeerencapsulationtype = "udp"

    DRAFTMSDPMIB_Msdppeertable_Msdppeerentry_Msdppeerencapsulationtype_gre DRAFTMSDPMIB_Msdppeertable_Msdppeerentry_Msdppeerencapsulationtype = "gre"
)

// DRAFTMSDPMIB_Msdppeertable_Msdppeerentry_Msdppeerstate represents The state of the MSDP TCP connection with this peer.
type DRAFTMSDPMIB_Msdppeertable_Msdppeerentry_Msdppeerstate string

const (
    DRAFTMSDPMIB_Msdppeertable_Msdppeerentry_Msdppeerstate_inactive DRAFTMSDPMIB_Msdppeertable_Msdppeerentry_Msdppeerstate = "inactive"

    DRAFTMSDPMIB_Msdppeertable_Msdppeerentry_Msdppeerstate_listen DRAFTMSDPMIB_Msdppeertable_Msdppeerentry_Msdppeerstate = "listen"

    DRAFTMSDPMIB_Msdppeertable_Msdppeerentry_Msdppeerstate_connecting DRAFTMSDPMIB_Msdppeertable_Msdppeerentry_Msdppeerstate = "connecting"

    DRAFTMSDPMIB_Msdppeertable_Msdppeerentry_Msdppeerstate_established DRAFTMSDPMIB_Msdppeertable_Msdppeerentry_Msdppeerstate = "established"

    DRAFTMSDPMIB_Msdppeertable_Msdppeerentry_Msdppeerstate_disabled DRAFTMSDPMIB_Msdppeertable_Msdppeerentry_Msdppeerstate = "disabled"
)

// DRAFTMSDPMIB_Msdpsacachetable
// The (conceptual) table listing the MSDP SA
// advertisements currently in the MSDP speaker's cache.
type DRAFTMSDPMIB_Msdpsacachetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry (conceptual row) representing an MSDP SA advert. The type is slice
    // of DRAFTMSDPMIB_Msdpsacachetable_Msdpsacacheentry.
    Msdpsacacheentry []DRAFTMSDPMIB_Msdpsacachetable_Msdpsacacheentry
}

func (msdpsacachetable *DRAFTMSDPMIB_Msdpsacachetable) GetFilter() yfilter.YFilter { return msdpsacachetable.YFilter }

func (msdpsacachetable *DRAFTMSDPMIB_Msdpsacachetable) SetFilter(yf yfilter.YFilter) { msdpsacachetable.YFilter = yf }

func (msdpsacachetable *DRAFTMSDPMIB_Msdpsacachetable) GetGoName(yname string) string {
    if yname == "msdpSACacheEntry" { return "Msdpsacacheentry" }
    return ""
}

func (msdpsacachetable *DRAFTMSDPMIB_Msdpsacachetable) GetSegmentPath() string {
    return "msdpSACacheTable"
}

func (msdpsacachetable *DRAFTMSDPMIB_Msdpsacachetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "msdpSACacheEntry" {
        for _, c := range msdpsacachetable.Msdpsacacheentry {
            if msdpsacachetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DRAFTMSDPMIB_Msdpsacachetable_Msdpsacacheentry{}
        msdpsacachetable.Msdpsacacheentry = append(msdpsacachetable.Msdpsacacheentry, child)
        return &msdpsacachetable.Msdpsacacheentry[len(msdpsacachetable.Msdpsacacheentry)-1]
    }
    return nil
}

func (msdpsacachetable *DRAFTMSDPMIB_Msdpsacachetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range msdpsacachetable.Msdpsacacheentry {
        children[msdpsacachetable.Msdpsacacheentry[i].GetSegmentPath()] = &msdpsacachetable.Msdpsacacheentry[i]
    }
    return children
}

func (msdpsacachetable *DRAFTMSDPMIB_Msdpsacachetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (msdpsacachetable *DRAFTMSDPMIB_Msdpsacachetable) GetBundleName() string { return "cisco_ios_xe" }

func (msdpsacachetable *DRAFTMSDPMIB_Msdpsacachetable) GetYangName() string { return "msdpSACacheTable" }

func (msdpsacachetable *DRAFTMSDPMIB_Msdpsacachetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (msdpsacachetable *DRAFTMSDPMIB_Msdpsacachetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (msdpsacachetable *DRAFTMSDPMIB_Msdpsacachetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (msdpsacachetable *DRAFTMSDPMIB_Msdpsacachetable) SetParent(parent types.Entity) { msdpsacachetable.parent = parent }

func (msdpsacachetable *DRAFTMSDPMIB_Msdpsacachetable) GetParent() types.Entity { return msdpsacachetable.parent }

func (msdpsacachetable *DRAFTMSDPMIB_Msdpsacachetable) GetParentYangName() string { return "DRAFT-MSDP-MIB" }

// DRAFTMSDPMIB_Msdpsacachetable_Msdpsacacheentry
// An entry (conceptual row) representing an MSDP SA
// advert.
type DRAFTMSDPMIB_Msdpsacachetable_Msdpsacacheentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The group address of the SA Cache entry. The type
    // is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Msdpsacachegroupaddr interface{}

    // This attribute is a key. The source address of the SA Cache entry. The type
    // is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Msdpsacachesourceaddr interface{}

    // This attribute is a key. The address of the RP which originated the last SA
    // message accepted for this entry. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Msdpsacacheoriginrp interface{}

    // The peer from which this SA Cache entry was last accepted.  This address
    // must correspond to the msdpPeerRemoteAddress value for a row in the MSDP
    // Peer Table. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Msdpsacachepeerlearnedfrom interface{}

    // The peer from which an SA message corresponding to this cache entry would
    // be accepted (i.e. the RPF peer for msdpSACacheOriginRP).  This may be
    // different than msdpSACachePeerLearnedFrom if this entry was created by an
    // MSDP SA-Response.  This address must correspond to the
    // msdpPeerRemoteAddress value for a row in the MSDP Peer Table. The type is
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Msdpsacacherpfpeer interface{}

    // The number of MSDP SA messages received relevant to this cache entry.  This
    // object must be initialized to zero when creating a cache entry. The type is
    // interface{} with range: 0..4294967295.
    Msdpsacacheinsas interface{}

    // The number of MSDP encapsulated data packets received relevant to this
    // cache entry.  This object must be initialized to zero when creating a cache
    // entry. The type is interface{} with range: 0..4294967295.
    Msdpsacacheindatapackets interface{}

    // The time since this entry was placed in the SA cache. The type is
    // interface{} with range: 0..4294967295.
    Msdpsacacheuptime interface{}

    // The time remaining before this entry will expire from the SA cache. The
    // type is interface{} with range: 0..4294967295.
    Msdpsacacheexpirytime interface{}

    // The status of this row in the table.  The only allowable actions are to
    // retreive the status, which will be `active', or to set the status to
    // `destroy' in order to remove this entry from the cache. The type is
    // RowStatus.
    Msdpsacachestatus interface{}
}

func (msdpsacacheentry *DRAFTMSDPMIB_Msdpsacachetable_Msdpsacacheentry) GetFilter() yfilter.YFilter { return msdpsacacheentry.YFilter }

func (msdpsacacheentry *DRAFTMSDPMIB_Msdpsacachetable_Msdpsacacheentry) SetFilter(yf yfilter.YFilter) { msdpsacacheentry.YFilter = yf }

func (msdpsacacheentry *DRAFTMSDPMIB_Msdpsacachetable_Msdpsacacheentry) GetGoName(yname string) string {
    if yname == "msdpSACacheGroupAddr" { return "Msdpsacachegroupaddr" }
    if yname == "msdpSACacheSourceAddr" { return "Msdpsacachesourceaddr" }
    if yname == "msdpSACacheOriginRP" { return "Msdpsacacheoriginrp" }
    if yname == "msdpSACachePeerLearnedFrom" { return "Msdpsacachepeerlearnedfrom" }
    if yname == "msdpSACacheRPFPeer" { return "Msdpsacacherpfpeer" }
    if yname == "msdpSACacheInSAs" { return "Msdpsacacheinsas" }
    if yname == "msdpSACacheInDataPackets" { return "Msdpsacacheindatapackets" }
    if yname == "msdpSACacheUpTime" { return "Msdpsacacheuptime" }
    if yname == "msdpSACacheExpiryTime" { return "Msdpsacacheexpirytime" }
    if yname == "msdpSACacheStatus" { return "Msdpsacachestatus" }
    return ""
}

func (msdpsacacheentry *DRAFTMSDPMIB_Msdpsacachetable_Msdpsacacheentry) GetSegmentPath() string {
    return "msdpSACacheEntry" + "[msdpSACacheGroupAddr='" + fmt.Sprintf("%v", msdpsacacheentry.Msdpsacachegroupaddr) + "']" + "[msdpSACacheSourceAddr='" + fmt.Sprintf("%v", msdpsacacheentry.Msdpsacachesourceaddr) + "']" + "[msdpSACacheOriginRP='" + fmt.Sprintf("%v", msdpsacacheentry.Msdpsacacheoriginrp) + "']"
}

func (msdpsacacheentry *DRAFTMSDPMIB_Msdpsacachetable_Msdpsacacheentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (msdpsacacheentry *DRAFTMSDPMIB_Msdpsacachetable_Msdpsacacheentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (msdpsacacheentry *DRAFTMSDPMIB_Msdpsacachetable_Msdpsacacheentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["msdpSACacheGroupAddr"] = msdpsacacheentry.Msdpsacachegroupaddr
    leafs["msdpSACacheSourceAddr"] = msdpsacacheentry.Msdpsacachesourceaddr
    leafs["msdpSACacheOriginRP"] = msdpsacacheentry.Msdpsacacheoriginrp
    leafs["msdpSACachePeerLearnedFrom"] = msdpsacacheentry.Msdpsacachepeerlearnedfrom
    leafs["msdpSACacheRPFPeer"] = msdpsacacheentry.Msdpsacacherpfpeer
    leafs["msdpSACacheInSAs"] = msdpsacacheentry.Msdpsacacheinsas
    leafs["msdpSACacheInDataPackets"] = msdpsacacheentry.Msdpsacacheindatapackets
    leafs["msdpSACacheUpTime"] = msdpsacacheentry.Msdpsacacheuptime
    leafs["msdpSACacheExpiryTime"] = msdpsacacheentry.Msdpsacacheexpirytime
    leafs["msdpSACacheStatus"] = msdpsacacheentry.Msdpsacachestatus
    return leafs
}

func (msdpsacacheentry *DRAFTMSDPMIB_Msdpsacachetable_Msdpsacacheentry) GetBundleName() string { return "cisco_ios_xe" }

func (msdpsacacheentry *DRAFTMSDPMIB_Msdpsacachetable_Msdpsacacheentry) GetYangName() string { return "msdpSACacheEntry" }

func (msdpsacacheentry *DRAFTMSDPMIB_Msdpsacachetable_Msdpsacacheentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (msdpsacacheentry *DRAFTMSDPMIB_Msdpsacachetable_Msdpsacacheentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (msdpsacacheentry *DRAFTMSDPMIB_Msdpsacachetable_Msdpsacacheentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (msdpsacacheentry *DRAFTMSDPMIB_Msdpsacachetable_Msdpsacacheentry) SetParent(parent types.Entity) { msdpsacacheentry.parent = parent }

func (msdpsacacheentry *DRAFTMSDPMIB_Msdpsacachetable_Msdpsacacheentry) GetParent() types.Entity { return msdpsacacheentry.parent }

func (msdpsacacheentry *DRAFTMSDPMIB_Msdpsacachetable_Msdpsacacheentry) GetParentYangName() string { return "msdpSACacheTable" }

