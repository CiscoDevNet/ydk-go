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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Msdp DRAFTMSDPMIB_Msdp

    // The (conceptual) table listing group ranges and MSDP peers used when
    // deciding where to send an SA Request message when required.  If SA Caching
    // is enabled, this table may be empty.
    MsdpRequestsTable DRAFTMSDPMIB_MsdpRequestsTable

    // The (conceptual) table listing the MSDP speaker's peers.
    MsdpPeerTable DRAFTMSDPMIB_MsdpPeerTable

    // The (conceptual) table listing the MSDP SA advertisements currently in the
    // MSDP speaker's cache.
    MsdpSACacheTable DRAFTMSDPMIB_MsdpSACacheTable
}

func (dRAFTMSDPMIB *DRAFTMSDPMIB) GetEntityData() *types.CommonEntityData {
    dRAFTMSDPMIB.EntityData.YFilter = dRAFTMSDPMIB.YFilter
    dRAFTMSDPMIB.EntityData.YangName = "DRAFT-MSDP-MIB"
    dRAFTMSDPMIB.EntityData.BundleName = "cisco_ios_xe"
    dRAFTMSDPMIB.EntityData.ParentYangName = "DRAFT-MSDP-MIB"
    dRAFTMSDPMIB.EntityData.SegmentPath = "DRAFT-MSDP-MIB:DRAFT-MSDP-MIB"
    dRAFTMSDPMIB.EntityData.AbsolutePath = dRAFTMSDPMIB.EntityData.SegmentPath
    dRAFTMSDPMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dRAFTMSDPMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dRAFTMSDPMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dRAFTMSDPMIB.EntityData.Children = types.NewOrderedMap()
    dRAFTMSDPMIB.EntityData.Children.Append("msdp", types.YChild{"Msdp", &dRAFTMSDPMIB.Msdp})
    dRAFTMSDPMIB.EntityData.Children.Append("msdpRequestsTable", types.YChild{"MsdpRequestsTable", &dRAFTMSDPMIB.MsdpRequestsTable})
    dRAFTMSDPMIB.EntityData.Children.Append("msdpPeerTable", types.YChild{"MsdpPeerTable", &dRAFTMSDPMIB.MsdpPeerTable})
    dRAFTMSDPMIB.EntityData.Children.Append("msdpSACacheTable", types.YChild{"MsdpSACacheTable", &dRAFTMSDPMIB.MsdpSACacheTable})
    dRAFTMSDPMIB.EntityData.Leafs = types.NewOrderedMap()

    dRAFTMSDPMIB.EntityData.YListKeys = []string {}

    return &(dRAFTMSDPMIB.EntityData)
}

// DRAFTMSDPMIB_Msdp
type DRAFTMSDPMIB_Msdp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The state of MSDP on this MSDP speaker - globally enabled or disabled. The
    // type is bool.
    MsdpEnabled interface{}

    // The lifetime given to SA cache entries when created or refreshed.  A value
    // of 0 means no SA caching is done by this MSDP speaker. The type is
    // interface{} with range: 0..4294967295.
    MsdpCacheLifetime interface{}

    // The total number of entries in the SA Cache table. The type is interface{}
    // with range: 0..4294967295.
    MsdpNumSACacheEntries interface{}

    // The number of seconds in the MSDP SA Hold-down period. The type is
    // interface{} with range: 1..2147483647. Units are seconds.
    MsdpSAHoldDownPeriod interface{}
}

func (msdp *DRAFTMSDPMIB_Msdp) GetEntityData() *types.CommonEntityData {
    msdp.EntityData.YFilter = msdp.YFilter
    msdp.EntityData.YangName = "msdp"
    msdp.EntityData.BundleName = "cisco_ios_xe"
    msdp.EntityData.ParentYangName = "DRAFT-MSDP-MIB"
    msdp.EntityData.SegmentPath = "msdp"
    msdp.EntityData.AbsolutePath = "DRAFT-MSDP-MIB:DRAFT-MSDP-MIB/" + msdp.EntityData.SegmentPath
    msdp.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    msdp.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    msdp.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    msdp.EntityData.Children = types.NewOrderedMap()
    msdp.EntityData.Leafs = types.NewOrderedMap()
    msdp.EntityData.Leafs.Append("msdpEnabled", types.YLeaf{"MsdpEnabled", msdp.MsdpEnabled})
    msdp.EntityData.Leafs.Append("msdpCacheLifetime", types.YLeaf{"MsdpCacheLifetime", msdp.MsdpCacheLifetime})
    msdp.EntityData.Leafs.Append("msdpNumSACacheEntries", types.YLeaf{"MsdpNumSACacheEntries", msdp.MsdpNumSACacheEntries})
    msdp.EntityData.Leafs.Append("msdpSAHoldDownPeriod", types.YLeaf{"MsdpSAHoldDownPeriod", msdp.MsdpSAHoldDownPeriod})

    msdp.EntityData.YListKeys = []string {}

    return &(msdp.EntityData)
}

// DRAFTMSDPMIB_MsdpRequestsTable
// The (conceptual) table listing group ranges and MSDP
// peers used when deciding where to send an SA Request
// message when required.  If SA Caching is enabled, this
// table may be empty.
type DRAFTMSDPMIB_MsdpRequestsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry (conceptual row) representing a group range used when deciding
    // where to send an SA Request message. The type is slice of
    // DRAFTMSDPMIB_MsdpRequestsTable_MsdpRequestsEntry.
    MsdpRequestsEntry []*DRAFTMSDPMIB_MsdpRequestsTable_MsdpRequestsEntry
}

func (msdpRequestsTable *DRAFTMSDPMIB_MsdpRequestsTable) GetEntityData() *types.CommonEntityData {
    msdpRequestsTable.EntityData.YFilter = msdpRequestsTable.YFilter
    msdpRequestsTable.EntityData.YangName = "msdpRequestsTable"
    msdpRequestsTable.EntityData.BundleName = "cisco_ios_xe"
    msdpRequestsTable.EntityData.ParentYangName = "DRAFT-MSDP-MIB"
    msdpRequestsTable.EntityData.SegmentPath = "msdpRequestsTable"
    msdpRequestsTable.EntityData.AbsolutePath = "DRAFT-MSDP-MIB:DRAFT-MSDP-MIB/" + msdpRequestsTable.EntityData.SegmentPath
    msdpRequestsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    msdpRequestsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    msdpRequestsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    msdpRequestsTable.EntityData.Children = types.NewOrderedMap()
    msdpRequestsTable.EntityData.Children.Append("msdpRequestsEntry", types.YChild{"MsdpRequestsEntry", nil})
    for i := range msdpRequestsTable.MsdpRequestsEntry {
        msdpRequestsTable.EntityData.Children.Append(types.GetSegmentPath(msdpRequestsTable.MsdpRequestsEntry[i]), types.YChild{"MsdpRequestsEntry", msdpRequestsTable.MsdpRequestsEntry[i]})
    }
    msdpRequestsTable.EntityData.Leafs = types.NewOrderedMap()

    msdpRequestsTable.EntityData.YListKeys = []string {}

    return &(msdpRequestsTable.EntityData)
}

// DRAFTMSDPMIB_MsdpRequestsTable_MsdpRequestsEntry
// An entry (conceptual row) representing a group range
// used when deciding where to send an SA Request
// message.
type DRAFTMSDPMIB_MsdpRequestsTable_MsdpRequestsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The group address that, when combined with the
    // mask in this entry, represents the group range for which this peer will
    // service MSDP SA Requests. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    MsdpRequestsGroupAddress interface{}

    // This attribute is a key. The mask that, when combined with the group
    // address in this entry, represents the group range for which this peer will
    // service MSDP SA Requests. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    MsdpRequestsGroupMask interface{}

    // The peer to which MSDP SA Requests for groups matching this entry's group
    // range will be sent.  Must match the INDEX of a row in the msdpPeerTable.
    // The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    MsdpRequestsPeer interface{}

    // The status of this row, by which new rows may be added to the table. The
    // type is RowStatus.
    MsdpRequestsStatus interface{}
}

func (msdpRequestsEntry *DRAFTMSDPMIB_MsdpRequestsTable_MsdpRequestsEntry) GetEntityData() *types.CommonEntityData {
    msdpRequestsEntry.EntityData.YFilter = msdpRequestsEntry.YFilter
    msdpRequestsEntry.EntityData.YangName = "msdpRequestsEntry"
    msdpRequestsEntry.EntityData.BundleName = "cisco_ios_xe"
    msdpRequestsEntry.EntityData.ParentYangName = "msdpRequestsTable"
    msdpRequestsEntry.EntityData.SegmentPath = "msdpRequestsEntry" + types.AddKeyToken(msdpRequestsEntry.MsdpRequestsGroupAddress, "msdpRequestsGroupAddress") + types.AddKeyToken(msdpRequestsEntry.MsdpRequestsGroupMask, "msdpRequestsGroupMask")
    msdpRequestsEntry.EntityData.AbsolutePath = "DRAFT-MSDP-MIB:DRAFT-MSDP-MIB/msdpRequestsTable/" + msdpRequestsEntry.EntityData.SegmentPath
    msdpRequestsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    msdpRequestsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    msdpRequestsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    msdpRequestsEntry.EntityData.Children = types.NewOrderedMap()
    msdpRequestsEntry.EntityData.Leafs = types.NewOrderedMap()
    msdpRequestsEntry.EntityData.Leafs.Append("msdpRequestsGroupAddress", types.YLeaf{"MsdpRequestsGroupAddress", msdpRequestsEntry.MsdpRequestsGroupAddress})
    msdpRequestsEntry.EntityData.Leafs.Append("msdpRequestsGroupMask", types.YLeaf{"MsdpRequestsGroupMask", msdpRequestsEntry.MsdpRequestsGroupMask})
    msdpRequestsEntry.EntityData.Leafs.Append("msdpRequestsPeer", types.YLeaf{"MsdpRequestsPeer", msdpRequestsEntry.MsdpRequestsPeer})
    msdpRequestsEntry.EntityData.Leafs.Append("msdpRequestsStatus", types.YLeaf{"MsdpRequestsStatus", msdpRequestsEntry.MsdpRequestsStatus})

    msdpRequestsEntry.EntityData.YListKeys = []string {"MsdpRequestsGroupAddress", "MsdpRequestsGroupMask"}

    return &(msdpRequestsEntry.EntityData)
}

// DRAFTMSDPMIB_MsdpPeerTable
// The (conceptual) table listing the MSDP speaker's
// peers.
type DRAFTMSDPMIB_MsdpPeerTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry (conceptual row) representing an MSDP peer. The type is slice of
    // DRAFTMSDPMIB_MsdpPeerTable_MsdpPeerEntry.
    MsdpPeerEntry []*DRAFTMSDPMIB_MsdpPeerTable_MsdpPeerEntry
}

func (msdpPeerTable *DRAFTMSDPMIB_MsdpPeerTable) GetEntityData() *types.CommonEntityData {
    msdpPeerTable.EntityData.YFilter = msdpPeerTable.YFilter
    msdpPeerTable.EntityData.YangName = "msdpPeerTable"
    msdpPeerTable.EntityData.BundleName = "cisco_ios_xe"
    msdpPeerTable.EntityData.ParentYangName = "DRAFT-MSDP-MIB"
    msdpPeerTable.EntityData.SegmentPath = "msdpPeerTable"
    msdpPeerTable.EntityData.AbsolutePath = "DRAFT-MSDP-MIB:DRAFT-MSDP-MIB/" + msdpPeerTable.EntityData.SegmentPath
    msdpPeerTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    msdpPeerTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    msdpPeerTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    msdpPeerTable.EntityData.Children = types.NewOrderedMap()
    msdpPeerTable.EntityData.Children.Append("msdpPeerEntry", types.YChild{"MsdpPeerEntry", nil})
    for i := range msdpPeerTable.MsdpPeerEntry {
        msdpPeerTable.EntityData.Children.Append(types.GetSegmentPath(msdpPeerTable.MsdpPeerEntry[i]), types.YChild{"MsdpPeerEntry", msdpPeerTable.MsdpPeerEntry[i]})
    }
    msdpPeerTable.EntityData.Leafs = types.NewOrderedMap()

    msdpPeerTable.EntityData.YListKeys = []string {}

    return &(msdpPeerTable.EntityData)
}

// DRAFTMSDPMIB_MsdpPeerTable_MsdpPeerEntry
// An entry (conceptual row) representing an MSDP peer.
type DRAFTMSDPMIB_MsdpPeerTable_MsdpPeerEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The address of the remote MSDP peer. The type is
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    MsdpPeerRemoteAddress interface{}

    // The state of the MSDP TCP connection with this peer. The type is
    // MsdpPeerState.
    MsdpPeerState interface{}

    // The number of RPF failures on SA messages received from this peer. The type
    // is interface{} with range: 0..4294967295.
    MsdpPeerRPFFailures interface{}

    // The number of MSDP SA messages received on this connection.  This object
    // should be initialized to zero when the connection is established. The type
    // is interface{} with range: 0..4294967295.
    MsdpPeerInSAs interface{}

    // The number of MSDP SA messages transmitted on this connection.  This object
    // should be initialized to zero when the connection is established. The type
    // is interface{} with range: 0..4294967295.
    MsdpPeerOutSAs interface{}

    // The number of MSDP SA-Request messages received on this connection.  This
    // object should be initialized to zero when the connection is established.
    // The type is interface{} with range: 0..4294967295.
    MsdpPeerInSARequests interface{}

    // The number of MSDP SA-Request messages transmitted on this connection. 
    // This object should be initialized to zero when the connection is
    // established. The type is interface{} with range: 0..4294967295.
    MsdpPeerOutSARequests interface{}

    // The number of MSDP SA-Response messages received on this connection.  This
    // object should be initialized to zero when the connection is established.
    // The type is interface{} with range: 0..4294967295.
    MsdpPeerInSAResponses interface{}

    // The number of MSDP SA Response messages transmitted on this TCP connection.
    // This object should be initialized to zero when the connection is
    // established. The type is interface{} with range: 0..4294967295.
    MsdpPeerOutSAResponses interface{}

    // The total number of MSDP messages received on this TCP connection.  This
    // object should be initialized to zero when the connection is established.
    // The type is interface{} with range: 0..4294967295.
    MsdpPeerInControlMessages interface{}

    // The total number of MSDP messages transmitted on this TCP connection.  This
    // object should be initialized to zero when the connection is established.
    // The type is interface{} with range: 0..4294967295.
    MsdpPeerOutControlMessages interface{}

    // The total number of encapsulated data packets received from this peer. 
    // This object should be initialized to zero when the connection is
    // established. The type is interface{} with range: 0..4294967295.
    MsdpPeerInDataPackets interface{}

    // The total number of encapsulated data packets sent to this peer.  This
    // object should be initialized to zero when the connection is established.
    // The type is interface{} with range: 0..4294967295.
    MsdpPeerOutDataPackets interface{}

    // The total number of times the MSDP FSM transitioned into the established
    // state. The type is interface{} with range: 0..4294967295.
    MsdpPeerFsmEstablishedTransitions interface{}

    // This timer indicates how long (in seconds) this peer has been in the
    // Established state or how long since this peer was last in the Established
    // state.  It is set to zero when a new peer is configured or the MSDP speaker
    // is booted. The type is interface{} with range: 0..4294967295. Units are
    // seconds.
    MsdpPeerFsmEstablishedTime interface{}

    // Elapsed time in seconds since the last MSDP message was received from the
    // peer.  Each time msdpPeerInControlMessages is incremented, the value of
    // this object is set to zero (0). The type is interface{} with range:
    // 0..4294967295. Units are seconds.
    MsdpPeerInMessageElapsedTime interface{}

    // The local IP address of this entry's MSDP connection. The type is string
    // with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    MsdpPeerLocalAddress interface{}

    // Time interval in seconds for the MinSAAdvertisementInterval MSDP timer. The
    // type is interface{} with range: 1..2147483647. Units are seconds.
    MsdpPeerSAAdvPeriod interface{}

    // Time interval in seconds for the ConnectRetry timer. The type is
    // interface{} with range: 1..65535. Units are seconds.
    MsdpPeerConnectRetryInterval interface{}

    // Time interval in seconds for the Hold Timer configured for this MSDP
    // speaker with this peer. The type is interface{} with range: 0..0 |
    // 3..65535. Units are seconds.
    MsdpPeerHoldTimeConfigured interface{}

    // Time interval in seconds for the KeepAlive timer configured for this MSDP
    // speaker with this peer.  A reasonable maximum value for this timer would be
    // configured to be one third of that of msdpPeerHoldTimeConfigured.  If the
    // value of this object is zero (0), no periodic KEEPALIVE messages are sent
    // to the peer after the MSDP connection has been established. The type is
    // interface{} with range: 0..21845. Units are seconds.
    MsdpPeerKeepAliveConfigured interface{}

    // The minimum TTL a packet is required to have before it may be forwarded
    // using SA encapsulation to this peer. The type is interface{} with range:
    // 0..255.
    MsdpPeerDataTtl interface{}

    // This object indicates whether or not to process MSDP SA Request messages
    // from this peer.  If True(1), MSDP SA Request messages from this peer are
    // processed and replied to (if appropriate) with SA Response messages. If
    // False(2), MSDP SA Request messages from this peer are silently ignored.  It
    // defaults to False when msdpCacheLifetime is 0 and True when
    // msdpCacheLifetime is non-0. The type is bool.
    MsdpPeerProcessRequestsFrom interface{}

    // The RowStatus object by which peers can be added and deleted.  A transition
    // to 'active' will cause the MSDP Start Event to be generated.  A transition
    // out of the 'active' state will cause the MSDP Stop Event to be generated. 
    // Care should be used in providing write access to this object without
    // adequate authentication. The type is RowStatus.
    MsdpPeerStatus interface{}

    // The remote port for the TCP connection between the MSDP peers. The type is
    // interface{} with range: 0..65535.
    MsdpPeerRemotePort interface{}

    // The local port for the TCP connection between the MSDP peers. The type is
    // interface{} with range: 0..65535.
    MsdpPeerLocalPort interface{}

    // The status of the encapsulation negotiation state machine. The type is
    // MsdpPeerEncapsulationState.
    MsdpPeerEncapsulationState interface{}

    // The encapsulation in use when encapsulating data in SA messages to this
    // peer. The type is MsdpPeerEncapsulationType.
    MsdpPeerEncapsulationType interface{}

    // The number of times the state machine has transitioned from inactive to
    // connecting. The type is interface{} with range: 0..4294967295.
    MsdpPeerConnectionAttempts interface{}

    // The number of MSDP Notification messages received on this connection.  This
    // object should be initialized to zero when the connection is established.
    // The type is interface{} with range: 0..4294967295.
    MsdpPeerInNotifications interface{}

    // The number of MSDP Notification messages transmitted on this connection. 
    // This object should be initialized to zero when the connection is
    // established. The type is interface{} with range: 0..4294967295.
    MsdpPeerOutNotifications interface{}

    // The last error code and subcode seen by this peer on this connection.  If
    // no error has occurred, this field is zero.  Otherwise, the first byte of
    // this two byte OCTET STRING contains the error code, and the second byte
    // contains the subcode. The type is string with length: 2..2.
    MsdpPeerLastError interface{}
}

func (msdpPeerEntry *DRAFTMSDPMIB_MsdpPeerTable_MsdpPeerEntry) GetEntityData() *types.CommonEntityData {
    msdpPeerEntry.EntityData.YFilter = msdpPeerEntry.YFilter
    msdpPeerEntry.EntityData.YangName = "msdpPeerEntry"
    msdpPeerEntry.EntityData.BundleName = "cisco_ios_xe"
    msdpPeerEntry.EntityData.ParentYangName = "msdpPeerTable"
    msdpPeerEntry.EntityData.SegmentPath = "msdpPeerEntry" + types.AddKeyToken(msdpPeerEntry.MsdpPeerRemoteAddress, "msdpPeerRemoteAddress")
    msdpPeerEntry.EntityData.AbsolutePath = "DRAFT-MSDP-MIB:DRAFT-MSDP-MIB/msdpPeerTable/" + msdpPeerEntry.EntityData.SegmentPath
    msdpPeerEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    msdpPeerEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    msdpPeerEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    msdpPeerEntry.EntityData.Children = types.NewOrderedMap()
    msdpPeerEntry.EntityData.Leafs = types.NewOrderedMap()
    msdpPeerEntry.EntityData.Leafs.Append("msdpPeerRemoteAddress", types.YLeaf{"MsdpPeerRemoteAddress", msdpPeerEntry.MsdpPeerRemoteAddress})
    msdpPeerEntry.EntityData.Leafs.Append("msdpPeerState", types.YLeaf{"MsdpPeerState", msdpPeerEntry.MsdpPeerState})
    msdpPeerEntry.EntityData.Leafs.Append("msdpPeerRPFFailures", types.YLeaf{"MsdpPeerRPFFailures", msdpPeerEntry.MsdpPeerRPFFailures})
    msdpPeerEntry.EntityData.Leafs.Append("msdpPeerInSAs", types.YLeaf{"MsdpPeerInSAs", msdpPeerEntry.MsdpPeerInSAs})
    msdpPeerEntry.EntityData.Leafs.Append("msdpPeerOutSAs", types.YLeaf{"MsdpPeerOutSAs", msdpPeerEntry.MsdpPeerOutSAs})
    msdpPeerEntry.EntityData.Leafs.Append("msdpPeerInSARequests", types.YLeaf{"MsdpPeerInSARequests", msdpPeerEntry.MsdpPeerInSARequests})
    msdpPeerEntry.EntityData.Leafs.Append("msdpPeerOutSARequests", types.YLeaf{"MsdpPeerOutSARequests", msdpPeerEntry.MsdpPeerOutSARequests})
    msdpPeerEntry.EntityData.Leafs.Append("msdpPeerInSAResponses", types.YLeaf{"MsdpPeerInSAResponses", msdpPeerEntry.MsdpPeerInSAResponses})
    msdpPeerEntry.EntityData.Leafs.Append("msdpPeerOutSAResponses", types.YLeaf{"MsdpPeerOutSAResponses", msdpPeerEntry.MsdpPeerOutSAResponses})
    msdpPeerEntry.EntityData.Leafs.Append("msdpPeerInControlMessages", types.YLeaf{"MsdpPeerInControlMessages", msdpPeerEntry.MsdpPeerInControlMessages})
    msdpPeerEntry.EntityData.Leafs.Append("msdpPeerOutControlMessages", types.YLeaf{"MsdpPeerOutControlMessages", msdpPeerEntry.MsdpPeerOutControlMessages})
    msdpPeerEntry.EntityData.Leafs.Append("msdpPeerInDataPackets", types.YLeaf{"MsdpPeerInDataPackets", msdpPeerEntry.MsdpPeerInDataPackets})
    msdpPeerEntry.EntityData.Leafs.Append("msdpPeerOutDataPackets", types.YLeaf{"MsdpPeerOutDataPackets", msdpPeerEntry.MsdpPeerOutDataPackets})
    msdpPeerEntry.EntityData.Leafs.Append("msdpPeerFsmEstablishedTransitions", types.YLeaf{"MsdpPeerFsmEstablishedTransitions", msdpPeerEntry.MsdpPeerFsmEstablishedTransitions})
    msdpPeerEntry.EntityData.Leafs.Append("msdpPeerFsmEstablishedTime", types.YLeaf{"MsdpPeerFsmEstablishedTime", msdpPeerEntry.MsdpPeerFsmEstablishedTime})
    msdpPeerEntry.EntityData.Leafs.Append("msdpPeerInMessageElapsedTime", types.YLeaf{"MsdpPeerInMessageElapsedTime", msdpPeerEntry.MsdpPeerInMessageElapsedTime})
    msdpPeerEntry.EntityData.Leafs.Append("msdpPeerLocalAddress", types.YLeaf{"MsdpPeerLocalAddress", msdpPeerEntry.MsdpPeerLocalAddress})
    msdpPeerEntry.EntityData.Leafs.Append("msdpPeerSAAdvPeriod", types.YLeaf{"MsdpPeerSAAdvPeriod", msdpPeerEntry.MsdpPeerSAAdvPeriod})
    msdpPeerEntry.EntityData.Leafs.Append("msdpPeerConnectRetryInterval", types.YLeaf{"MsdpPeerConnectRetryInterval", msdpPeerEntry.MsdpPeerConnectRetryInterval})
    msdpPeerEntry.EntityData.Leafs.Append("msdpPeerHoldTimeConfigured", types.YLeaf{"MsdpPeerHoldTimeConfigured", msdpPeerEntry.MsdpPeerHoldTimeConfigured})
    msdpPeerEntry.EntityData.Leafs.Append("msdpPeerKeepAliveConfigured", types.YLeaf{"MsdpPeerKeepAliveConfigured", msdpPeerEntry.MsdpPeerKeepAliveConfigured})
    msdpPeerEntry.EntityData.Leafs.Append("msdpPeerDataTtl", types.YLeaf{"MsdpPeerDataTtl", msdpPeerEntry.MsdpPeerDataTtl})
    msdpPeerEntry.EntityData.Leafs.Append("msdpPeerProcessRequestsFrom", types.YLeaf{"MsdpPeerProcessRequestsFrom", msdpPeerEntry.MsdpPeerProcessRequestsFrom})
    msdpPeerEntry.EntityData.Leafs.Append("msdpPeerStatus", types.YLeaf{"MsdpPeerStatus", msdpPeerEntry.MsdpPeerStatus})
    msdpPeerEntry.EntityData.Leafs.Append("msdpPeerRemotePort", types.YLeaf{"MsdpPeerRemotePort", msdpPeerEntry.MsdpPeerRemotePort})
    msdpPeerEntry.EntityData.Leafs.Append("msdpPeerLocalPort", types.YLeaf{"MsdpPeerLocalPort", msdpPeerEntry.MsdpPeerLocalPort})
    msdpPeerEntry.EntityData.Leafs.Append("msdpPeerEncapsulationState", types.YLeaf{"MsdpPeerEncapsulationState", msdpPeerEntry.MsdpPeerEncapsulationState})
    msdpPeerEntry.EntityData.Leafs.Append("msdpPeerEncapsulationType", types.YLeaf{"MsdpPeerEncapsulationType", msdpPeerEntry.MsdpPeerEncapsulationType})
    msdpPeerEntry.EntityData.Leafs.Append("msdpPeerConnectionAttempts", types.YLeaf{"MsdpPeerConnectionAttempts", msdpPeerEntry.MsdpPeerConnectionAttempts})
    msdpPeerEntry.EntityData.Leafs.Append("msdpPeerInNotifications", types.YLeaf{"MsdpPeerInNotifications", msdpPeerEntry.MsdpPeerInNotifications})
    msdpPeerEntry.EntityData.Leafs.Append("msdpPeerOutNotifications", types.YLeaf{"MsdpPeerOutNotifications", msdpPeerEntry.MsdpPeerOutNotifications})
    msdpPeerEntry.EntityData.Leafs.Append("msdpPeerLastError", types.YLeaf{"MsdpPeerLastError", msdpPeerEntry.MsdpPeerLastError})

    msdpPeerEntry.EntityData.YListKeys = []string {"MsdpPeerRemoteAddress"}

    return &(msdpPeerEntry.EntityData)
}

// DRAFTMSDPMIB_MsdpPeerTable_MsdpPeerEntry_MsdpPeerEncapsulationState represents machine.
type DRAFTMSDPMIB_MsdpPeerTable_MsdpPeerEntry_MsdpPeerEncapsulationState string

const (
    DRAFTMSDPMIB_MsdpPeerTable_MsdpPeerEntry_MsdpPeerEncapsulationState_default_ DRAFTMSDPMIB_MsdpPeerTable_MsdpPeerEntry_MsdpPeerEncapsulationState = "default"

    DRAFTMSDPMIB_MsdpPeerTable_MsdpPeerEntry_MsdpPeerEncapsulationState_received DRAFTMSDPMIB_MsdpPeerTable_MsdpPeerEntry_MsdpPeerEncapsulationState = "received"

    DRAFTMSDPMIB_MsdpPeerTable_MsdpPeerEntry_MsdpPeerEncapsulationState_advertising DRAFTMSDPMIB_MsdpPeerTable_MsdpPeerEntry_MsdpPeerEncapsulationState = "advertising"

    DRAFTMSDPMIB_MsdpPeerTable_MsdpPeerEntry_MsdpPeerEncapsulationState_sent DRAFTMSDPMIB_MsdpPeerTable_MsdpPeerEntry_MsdpPeerEncapsulationState = "sent"

    DRAFTMSDPMIB_MsdpPeerTable_MsdpPeerEntry_MsdpPeerEncapsulationState_agreed DRAFTMSDPMIB_MsdpPeerTable_MsdpPeerEntry_MsdpPeerEncapsulationState = "agreed"

    DRAFTMSDPMIB_MsdpPeerTable_MsdpPeerEntry_MsdpPeerEncapsulationState_failed DRAFTMSDPMIB_MsdpPeerTable_MsdpPeerEntry_MsdpPeerEncapsulationState = "failed"
)

// DRAFTMSDPMIB_MsdpPeerTable_MsdpPeerEntry_MsdpPeerEncapsulationType represents SA messages to this peer.
type DRAFTMSDPMIB_MsdpPeerTable_MsdpPeerEntry_MsdpPeerEncapsulationType string

const (
    DRAFTMSDPMIB_MsdpPeerTable_MsdpPeerEntry_MsdpPeerEncapsulationType_tcp DRAFTMSDPMIB_MsdpPeerTable_MsdpPeerEntry_MsdpPeerEncapsulationType = "tcp"

    DRAFTMSDPMIB_MsdpPeerTable_MsdpPeerEntry_MsdpPeerEncapsulationType_udp DRAFTMSDPMIB_MsdpPeerTable_MsdpPeerEntry_MsdpPeerEncapsulationType = "udp"

    DRAFTMSDPMIB_MsdpPeerTable_MsdpPeerEntry_MsdpPeerEncapsulationType_gre DRAFTMSDPMIB_MsdpPeerTable_MsdpPeerEntry_MsdpPeerEncapsulationType = "gre"
)

// DRAFTMSDPMIB_MsdpPeerTable_MsdpPeerEntry_MsdpPeerState represents The state of the MSDP TCP connection with this peer.
type DRAFTMSDPMIB_MsdpPeerTable_MsdpPeerEntry_MsdpPeerState string

const (
    DRAFTMSDPMIB_MsdpPeerTable_MsdpPeerEntry_MsdpPeerState_inactive DRAFTMSDPMIB_MsdpPeerTable_MsdpPeerEntry_MsdpPeerState = "inactive"

    DRAFTMSDPMIB_MsdpPeerTable_MsdpPeerEntry_MsdpPeerState_listen DRAFTMSDPMIB_MsdpPeerTable_MsdpPeerEntry_MsdpPeerState = "listen"

    DRAFTMSDPMIB_MsdpPeerTable_MsdpPeerEntry_MsdpPeerState_connecting DRAFTMSDPMIB_MsdpPeerTable_MsdpPeerEntry_MsdpPeerState = "connecting"

    DRAFTMSDPMIB_MsdpPeerTable_MsdpPeerEntry_MsdpPeerState_established DRAFTMSDPMIB_MsdpPeerTable_MsdpPeerEntry_MsdpPeerState = "established"

    DRAFTMSDPMIB_MsdpPeerTable_MsdpPeerEntry_MsdpPeerState_disabled DRAFTMSDPMIB_MsdpPeerTable_MsdpPeerEntry_MsdpPeerState = "disabled"
)

// DRAFTMSDPMIB_MsdpSACacheTable
// The (conceptual) table listing the MSDP SA
// advertisements currently in the MSDP speaker's cache.
type DRAFTMSDPMIB_MsdpSACacheTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry (conceptual row) representing an MSDP SA advert. The type is slice
    // of DRAFTMSDPMIB_MsdpSACacheTable_MsdpSACacheEntry.
    MsdpSACacheEntry []*DRAFTMSDPMIB_MsdpSACacheTable_MsdpSACacheEntry
}

func (msdpSACacheTable *DRAFTMSDPMIB_MsdpSACacheTable) GetEntityData() *types.CommonEntityData {
    msdpSACacheTable.EntityData.YFilter = msdpSACacheTable.YFilter
    msdpSACacheTable.EntityData.YangName = "msdpSACacheTable"
    msdpSACacheTable.EntityData.BundleName = "cisco_ios_xe"
    msdpSACacheTable.EntityData.ParentYangName = "DRAFT-MSDP-MIB"
    msdpSACacheTable.EntityData.SegmentPath = "msdpSACacheTable"
    msdpSACacheTable.EntityData.AbsolutePath = "DRAFT-MSDP-MIB:DRAFT-MSDP-MIB/" + msdpSACacheTable.EntityData.SegmentPath
    msdpSACacheTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    msdpSACacheTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    msdpSACacheTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    msdpSACacheTable.EntityData.Children = types.NewOrderedMap()
    msdpSACacheTable.EntityData.Children.Append("msdpSACacheEntry", types.YChild{"MsdpSACacheEntry", nil})
    for i := range msdpSACacheTable.MsdpSACacheEntry {
        msdpSACacheTable.EntityData.Children.Append(types.GetSegmentPath(msdpSACacheTable.MsdpSACacheEntry[i]), types.YChild{"MsdpSACacheEntry", msdpSACacheTable.MsdpSACacheEntry[i]})
    }
    msdpSACacheTable.EntityData.Leafs = types.NewOrderedMap()

    msdpSACacheTable.EntityData.YListKeys = []string {}

    return &(msdpSACacheTable.EntityData)
}

// DRAFTMSDPMIB_MsdpSACacheTable_MsdpSACacheEntry
// An entry (conceptual row) representing an MSDP SA
// advert.
type DRAFTMSDPMIB_MsdpSACacheTable_MsdpSACacheEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The group address of the SA Cache entry. The type
    // is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    MsdpSACacheGroupAddr interface{}

    // This attribute is a key. The source address of the SA Cache entry. The type
    // is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    MsdpSACacheSourceAddr interface{}

    // This attribute is a key. The address of the RP which originated the last SA
    // message accepted for this entry. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    MsdpSACacheOriginRP interface{}

    // The peer from which this SA Cache entry was last accepted.  This address
    // must correspond to the msdpPeerRemoteAddress value for a row in the MSDP
    // Peer Table. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    MsdpSACachePeerLearnedFrom interface{}

    // The peer from which an SA message corresponding to this cache entry would
    // be accepted (i.e. the RPF peer for msdpSACacheOriginRP).  This may be
    // different than msdpSACachePeerLearnedFrom if this entry was created by an
    // MSDP SA-Response.  This address must correspond to the
    // msdpPeerRemoteAddress value for a row in the MSDP Peer Table. The type is
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    MsdpSACacheRPFPeer interface{}

    // The number of MSDP SA messages received relevant to this cache entry.  This
    // object must be initialized to zero when creating a cache entry. The type is
    // interface{} with range: 0..4294967295.
    MsdpSACacheInSAs interface{}

    // The number of MSDP encapsulated data packets received relevant to this
    // cache entry.  This object must be initialized to zero when creating a cache
    // entry. The type is interface{} with range: 0..4294967295.
    MsdpSACacheInDataPackets interface{}

    // The time since this entry was placed in the SA cache. The type is
    // interface{} with range: 0..4294967295.
    MsdpSACacheUpTime interface{}

    // The time remaining before this entry will expire from the SA cache. The
    // type is interface{} with range: 0..4294967295.
    MsdpSACacheExpiryTime interface{}

    // The status of this row in the table.  The only allowable actions are to
    // retreive the status, which will be `active', or to set the status to
    // `destroy' in order to remove this entry from the cache. The type is
    // RowStatus.
    MsdpSACacheStatus interface{}
}

func (msdpSACacheEntry *DRAFTMSDPMIB_MsdpSACacheTable_MsdpSACacheEntry) GetEntityData() *types.CommonEntityData {
    msdpSACacheEntry.EntityData.YFilter = msdpSACacheEntry.YFilter
    msdpSACacheEntry.EntityData.YangName = "msdpSACacheEntry"
    msdpSACacheEntry.EntityData.BundleName = "cisco_ios_xe"
    msdpSACacheEntry.EntityData.ParentYangName = "msdpSACacheTable"
    msdpSACacheEntry.EntityData.SegmentPath = "msdpSACacheEntry" + types.AddKeyToken(msdpSACacheEntry.MsdpSACacheGroupAddr, "msdpSACacheGroupAddr") + types.AddKeyToken(msdpSACacheEntry.MsdpSACacheSourceAddr, "msdpSACacheSourceAddr") + types.AddKeyToken(msdpSACacheEntry.MsdpSACacheOriginRP, "msdpSACacheOriginRP")
    msdpSACacheEntry.EntityData.AbsolutePath = "DRAFT-MSDP-MIB:DRAFT-MSDP-MIB/msdpSACacheTable/" + msdpSACacheEntry.EntityData.SegmentPath
    msdpSACacheEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    msdpSACacheEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    msdpSACacheEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    msdpSACacheEntry.EntityData.Children = types.NewOrderedMap()
    msdpSACacheEntry.EntityData.Leafs = types.NewOrderedMap()
    msdpSACacheEntry.EntityData.Leafs.Append("msdpSACacheGroupAddr", types.YLeaf{"MsdpSACacheGroupAddr", msdpSACacheEntry.MsdpSACacheGroupAddr})
    msdpSACacheEntry.EntityData.Leafs.Append("msdpSACacheSourceAddr", types.YLeaf{"MsdpSACacheSourceAddr", msdpSACacheEntry.MsdpSACacheSourceAddr})
    msdpSACacheEntry.EntityData.Leafs.Append("msdpSACacheOriginRP", types.YLeaf{"MsdpSACacheOriginRP", msdpSACacheEntry.MsdpSACacheOriginRP})
    msdpSACacheEntry.EntityData.Leafs.Append("msdpSACachePeerLearnedFrom", types.YLeaf{"MsdpSACachePeerLearnedFrom", msdpSACacheEntry.MsdpSACachePeerLearnedFrom})
    msdpSACacheEntry.EntityData.Leafs.Append("msdpSACacheRPFPeer", types.YLeaf{"MsdpSACacheRPFPeer", msdpSACacheEntry.MsdpSACacheRPFPeer})
    msdpSACacheEntry.EntityData.Leafs.Append("msdpSACacheInSAs", types.YLeaf{"MsdpSACacheInSAs", msdpSACacheEntry.MsdpSACacheInSAs})
    msdpSACacheEntry.EntityData.Leafs.Append("msdpSACacheInDataPackets", types.YLeaf{"MsdpSACacheInDataPackets", msdpSACacheEntry.MsdpSACacheInDataPackets})
    msdpSACacheEntry.EntityData.Leafs.Append("msdpSACacheUpTime", types.YLeaf{"MsdpSACacheUpTime", msdpSACacheEntry.MsdpSACacheUpTime})
    msdpSACacheEntry.EntityData.Leafs.Append("msdpSACacheExpiryTime", types.YLeaf{"MsdpSACacheExpiryTime", msdpSACacheEntry.MsdpSACacheExpiryTime})
    msdpSACacheEntry.EntityData.Leafs.Append("msdpSACacheStatus", types.YLeaf{"MsdpSACacheStatus", msdpSACacheEntry.MsdpSACacheStatus})

    msdpSACacheEntry.EntityData.YListKeys = []string {"MsdpSACacheGroupAddr", "MsdpSACacheSourceAddr", "MsdpSACacheOriginRP"}

    return &(msdpSACacheEntry.EntityData)
}

