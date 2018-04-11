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
    Msdprequeststable DRAFTMSDPMIB_Msdprequeststable

    // The (conceptual) table listing the MSDP speaker's peers.
    Msdppeertable DRAFTMSDPMIB_Msdppeertable

    // The (conceptual) table listing the MSDP SA advertisements currently in the
    // MSDP speaker's cache.
    Msdpsacachetable DRAFTMSDPMIB_Msdpsacachetable
}

func (dRAFTMSDPMIB *DRAFTMSDPMIB) GetEntityData() *types.CommonEntityData {
    dRAFTMSDPMIB.EntityData.YFilter = dRAFTMSDPMIB.YFilter
    dRAFTMSDPMIB.EntityData.YangName = "DRAFT-MSDP-MIB"
    dRAFTMSDPMIB.EntityData.BundleName = "cisco_ios_xe"
    dRAFTMSDPMIB.EntityData.ParentYangName = "DRAFT-MSDP-MIB"
    dRAFTMSDPMIB.EntityData.SegmentPath = "DRAFT-MSDP-MIB:DRAFT-MSDP-MIB"
    dRAFTMSDPMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dRAFTMSDPMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dRAFTMSDPMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dRAFTMSDPMIB.EntityData.Children = make(map[string]types.YChild)
    dRAFTMSDPMIB.EntityData.Children["msdp"] = types.YChild{"Msdp", &dRAFTMSDPMIB.Msdp}
    dRAFTMSDPMIB.EntityData.Children["msdpRequestsTable"] = types.YChild{"Msdprequeststable", &dRAFTMSDPMIB.Msdprequeststable}
    dRAFTMSDPMIB.EntityData.Children["msdpPeerTable"] = types.YChild{"Msdppeertable", &dRAFTMSDPMIB.Msdppeertable}
    dRAFTMSDPMIB.EntityData.Children["msdpSACacheTable"] = types.YChild{"Msdpsacachetable", &dRAFTMSDPMIB.Msdpsacachetable}
    dRAFTMSDPMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dRAFTMSDPMIB.EntityData)
}

// DRAFTMSDPMIB_Msdp
type DRAFTMSDPMIB_Msdp struct {
    EntityData types.CommonEntityData
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

func (msdp *DRAFTMSDPMIB_Msdp) GetEntityData() *types.CommonEntityData {
    msdp.EntityData.YFilter = msdp.YFilter
    msdp.EntityData.YangName = "msdp"
    msdp.EntityData.BundleName = "cisco_ios_xe"
    msdp.EntityData.ParentYangName = "DRAFT-MSDP-MIB"
    msdp.EntityData.SegmentPath = "msdp"
    msdp.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    msdp.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    msdp.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    msdp.EntityData.Children = make(map[string]types.YChild)
    msdp.EntityData.Leafs = make(map[string]types.YLeaf)
    msdp.EntityData.Leafs["msdpEnabled"] = types.YLeaf{"Msdpenabled", msdp.Msdpenabled}
    msdp.EntityData.Leafs["msdpCacheLifetime"] = types.YLeaf{"Msdpcachelifetime", msdp.Msdpcachelifetime}
    msdp.EntityData.Leafs["msdpNumSACacheEntries"] = types.YLeaf{"Msdpnumsacacheentries", msdp.Msdpnumsacacheentries}
    msdp.EntityData.Leafs["msdpSAHoldDownPeriod"] = types.YLeaf{"Msdpsaholddownperiod", msdp.Msdpsaholddownperiod}
    return &(msdp.EntityData)
}

// DRAFTMSDPMIB_Msdprequeststable
// The (conceptual) table listing group ranges and MSDP
// peers used when deciding where to send an SA Request
// message when required.  If SA Caching is enabled, this
// table may be empty.
type DRAFTMSDPMIB_Msdprequeststable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry (conceptual row) representing a group range used when deciding
    // where to send an SA Request message. The type is slice of
    // DRAFTMSDPMIB_Msdprequeststable_Msdprequestsentry.
    Msdprequestsentry []DRAFTMSDPMIB_Msdprequeststable_Msdprequestsentry
}

func (msdprequeststable *DRAFTMSDPMIB_Msdprequeststable) GetEntityData() *types.CommonEntityData {
    msdprequeststable.EntityData.YFilter = msdprequeststable.YFilter
    msdprequeststable.EntityData.YangName = "msdpRequestsTable"
    msdprequeststable.EntityData.BundleName = "cisco_ios_xe"
    msdprequeststable.EntityData.ParentYangName = "DRAFT-MSDP-MIB"
    msdprequeststable.EntityData.SegmentPath = "msdpRequestsTable"
    msdprequeststable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    msdprequeststable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    msdprequeststable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    msdprequeststable.EntityData.Children = make(map[string]types.YChild)
    msdprequeststable.EntityData.Children["msdpRequestsEntry"] = types.YChild{"Msdprequestsentry", nil}
    for i := range msdprequeststable.Msdprequestsentry {
        msdprequeststable.EntityData.Children[types.GetSegmentPath(&msdprequeststable.Msdprequestsentry[i])] = types.YChild{"Msdprequestsentry", &msdprequeststable.Msdprequestsentry[i]}
    }
    msdprequeststable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(msdprequeststable.EntityData)
}

// DRAFTMSDPMIB_Msdprequeststable_Msdprequestsentry
// An entry (conceptual row) representing a group range
// used when deciding where to send an SA Request
// message.
type DRAFTMSDPMIB_Msdprequeststable_Msdprequestsentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The group address that, when combined with the
    // mask in this entry, represents the group range for which this peer will
    // service MSDP SA Requests. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Msdprequestsgroupaddress interface{}

    // This attribute is a key. The mask that, when combined with the group
    // address in this entry, represents the group range for which this peer will
    // service MSDP SA Requests. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Msdprequestsgroupmask interface{}

    // The peer to which MSDP SA Requests for groups matching this entry's group
    // range will be sent.  Must match the INDEX of a row in the msdpPeerTable.
    // The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Msdprequestspeer interface{}

    // The status of this row, by which new rows may be added to the table. The
    // type is RowStatus.
    Msdprequestsstatus interface{}
}

func (msdprequestsentry *DRAFTMSDPMIB_Msdprequeststable_Msdprequestsentry) GetEntityData() *types.CommonEntityData {
    msdprequestsentry.EntityData.YFilter = msdprequestsentry.YFilter
    msdprequestsentry.EntityData.YangName = "msdpRequestsEntry"
    msdprequestsentry.EntityData.BundleName = "cisco_ios_xe"
    msdprequestsentry.EntityData.ParentYangName = "msdpRequestsTable"
    msdprequestsentry.EntityData.SegmentPath = "msdpRequestsEntry" + "[msdpRequestsGroupAddress='" + fmt.Sprintf("%v", msdprequestsentry.Msdprequestsgroupaddress) + "']" + "[msdpRequestsGroupMask='" + fmt.Sprintf("%v", msdprequestsentry.Msdprequestsgroupmask) + "']"
    msdprequestsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    msdprequestsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    msdprequestsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    msdprequestsentry.EntityData.Children = make(map[string]types.YChild)
    msdprequestsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    msdprequestsentry.EntityData.Leafs["msdpRequestsGroupAddress"] = types.YLeaf{"Msdprequestsgroupaddress", msdprequestsentry.Msdprequestsgroupaddress}
    msdprequestsentry.EntityData.Leafs["msdpRequestsGroupMask"] = types.YLeaf{"Msdprequestsgroupmask", msdprequestsentry.Msdprequestsgroupmask}
    msdprequestsentry.EntityData.Leafs["msdpRequestsPeer"] = types.YLeaf{"Msdprequestspeer", msdprequestsentry.Msdprequestspeer}
    msdprequestsentry.EntityData.Leafs["msdpRequestsStatus"] = types.YLeaf{"Msdprequestsstatus", msdprequestsentry.Msdprequestsstatus}
    return &(msdprequestsentry.EntityData)
}

// DRAFTMSDPMIB_Msdppeertable
// The (conceptual) table listing the MSDP speaker's
// peers.
type DRAFTMSDPMIB_Msdppeertable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry (conceptual row) representing an MSDP peer. The type is slice of
    // DRAFTMSDPMIB_Msdppeertable_Msdppeerentry.
    Msdppeerentry []DRAFTMSDPMIB_Msdppeertable_Msdppeerentry
}

func (msdppeertable *DRAFTMSDPMIB_Msdppeertable) GetEntityData() *types.CommonEntityData {
    msdppeertable.EntityData.YFilter = msdppeertable.YFilter
    msdppeertable.EntityData.YangName = "msdpPeerTable"
    msdppeertable.EntityData.BundleName = "cisco_ios_xe"
    msdppeertable.EntityData.ParentYangName = "DRAFT-MSDP-MIB"
    msdppeertable.EntityData.SegmentPath = "msdpPeerTable"
    msdppeertable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    msdppeertable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    msdppeertable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    msdppeertable.EntityData.Children = make(map[string]types.YChild)
    msdppeertable.EntityData.Children["msdpPeerEntry"] = types.YChild{"Msdppeerentry", nil}
    for i := range msdppeertable.Msdppeerentry {
        msdppeertable.EntityData.Children[types.GetSegmentPath(&msdppeertable.Msdppeerentry[i])] = types.YChild{"Msdppeerentry", &msdppeertable.Msdppeerentry[i]}
    }
    msdppeertable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(msdppeertable.EntityData)
}

// DRAFTMSDPMIB_Msdppeertable_Msdppeerentry
// An entry (conceptual row) representing an MSDP peer.
type DRAFTMSDPMIB_Msdppeertable_Msdppeerentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The address of the remote MSDP peer. The type is
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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

func (msdppeerentry *DRAFTMSDPMIB_Msdppeertable_Msdppeerentry) GetEntityData() *types.CommonEntityData {
    msdppeerentry.EntityData.YFilter = msdppeerentry.YFilter
    msdppeerentry.EntityData.YangName = "msdpPeerEntry"
    msdppeerentry.EntityData.BundleName = "cisco_ios_xe"
    msdppeerentry.EntityData.ParentYangName = "msdpPeerTable"
    msdppeerentry.EntityData.SegmentPath = "msdpPeerEntry" + "[msdpPeerRemoteAddress='" + fmt.Sprintf("%v", msdppeerentry.Msdppeerremoteaddress) + "']"
    msdppeerentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    msdppeerentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    msdppeerentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    msdppeerentry.EntityData.Children = make(map[string]types.YChild)
    msdppeerentry.EntityData.Leafs = make(map[string]types.YLeaf)
    msdppeerentry.EntityData.Leafs["msdpPeerRemoteAddress"] = types.YLeaf{"Msdppeerremoteaddress", msdppeerentry.Msdppeerremoteaddress}
    msdppeerentry.EntityData.Leafs["msdpPeerState"] = types.YLeaf{"Msdppeerstate", msdppeerentry.Msdppeerstate}
    msdppeerentry.EntityData.Leafs["msdpPeerRPFFailures"] = types.YLeaf{"Msdppeerrpffailures", msdppeerentry.Msdppeerrpffailures}
    msdppeerentry.EntityData.Leafs["msdpPeerInSAs"] = types.YLeaf{"Msdppeerinsas", msdppeerentry.Msdppeerinsas}
    msdppeerentry.EntityData.Leafs["msdpPeerOutSAs"] = types.YLeaf{"Msdppeeroutsas", msdppeerentry.Msdppeeroutsas}
    msdppeerentry.EntityData.Leafs["msdpPeerInSARequests"] = types.YLeaf{"Msdppeerinsarequests", msdppeerentry.Msdppeerinsarequests}
    msdppeerentry.EntityData.Leafs["msdpPeerOutSARequests"] = types.YLeaf{"Msdppeeroutsarequests", msdppeerentry.Msdppeeroutsarequests}
    msdppeerentry.EntityData.Leafs["msdpPeerInSAResponses"] = types.YLeaf{"Msdppeerinsaresponses", msdppeerentry.Msdppeerinsaresponses}
    msdppeerentry.EntityData.Leafs["msdpPeerOutSAResponses"] = types.YLeaf{"Msdppeeroutsaresponses", msdppeerentry.Msdppeeroutsaresponses}
    msdppeerentry.EntityData.Leafs["msdpPeerInControlMessages"] = types.YLeaf{"Msdppeerincontrolmessages", msdppeerentry.Msdppeerincontrolmessages}
    msdppeerentry.EntityData.Leafs["msdpPeerOutControlMessages"] = types.YLeaf{"Msdppeeroutcontrolmessages", msdppeerentry.Msdppeeroutcontrolmessages}
    msdppeerentry.EntityData.Leafs["msdpPeerInDataPackets"] = types.YLeaf{"Msdppeerindatapackets", msdppeerentry.Msdppeerindatapackets}
    msdppeerentry.EntityData.Leafs["msdpPeerOutDataPackets"] = types.YLeaf{"Msdppeeroutdatapackets", msdppeerentry.Msdppeeroutdatapackets}
    msdppeerentry.EntityData.Leafs["msdpPeerFsmEstablishedTransitions"] = types.YLeaf{"Msdppeerfsmestablishedtransitions", msdppeerentry.Msdppeerfsmestablishedtransitions}
    msdppeerentry.EntityData.Leafs["msdpPeerFsmEstablishedTime"] = types.YLeaf{"Msdppeerfsmestablishedtime", msdppeerentry.Msdppeerfsmestablishedtime}
    msdppeerentry.EntityData.Leafs["msdpPeerInMessageElapsedTime"] = types.YLeaf{"Msdppeerinmessageelapsedtime", msdppeerentry.Msdppeerinmessageelapsedtime}
    msdppeerentry.EntityData.Leafs["msdpPeerLocalAddress"] = types.YLeaf{"Msdppeerlocaladdress", msdppeerentry.Msdppeerlocaladdress}
    msdppeerentry.EntityData.Leafs["msdpPeerSAAdvPeriod"] = types.YLeaf{"Msdppeersaadvperiod", msdppeerentry.Msdppeersaadvperiod}
    msdppeerentry.EntityData.Leafs["msdpPeerConnectRetryInterval"] = types.YLeaf{"Msdppeerconnectretryinterval", msdppeerentry.Msdppeerconnectretryinterval}
    msdppeerentry.EntityData.Leafs["msdpPeerHoldTimeConfigured"] = types.YLeaf{"Msdppeerholdtimeconfigured", msdppeerentry.Msdppeerholdtimeconfigured}
    msdppeerentry.EntityData.Leafs["msdpPeerKeepAliveConfigured"] = types.YLeaf{"Msdppeerkeepaliveconfigured", msdppeerentry.Msdppeerkeepaliveconfigured}
    msdppeerentry.EntityData.Leafs["msdpPeerDataTtl"] = types.YLeaf{"Msdppeerdatattl", msdppeerentry.Msdppeerdatattl}
    msdppeerentry.EntityData.Leafs["msdpPeerProcessRequestsFrom"] = types.YLeaf{"Msdppeerprocessrequestsfrom", msdppeerentry.Msdppeerprocessrequestsfrom}
    msdppeerentry.EntityData.Leafs["msdpPeerStatus"] = types.YLeaf{"Msdppeerstatus", msdppeerentry.Msdppeerstatus}
    msdppeerentry.EntityData.Leafs["msdpPeerRemotePort"] = types.YLeaf{"Msdppeerremoteport", msdppeerentry.Msdppeerremoteport}
    msdppeerentry.EntityData.Leafs["msdpPeerLocalPort"] = types.YLeaf{"Msdppeerlocalport", msdppeerentry.Msdppeerlocalport}
    msdppeerentry.EntityData.Leafs["msdpPeerEncapsulationState"] = types.YLeaf{"Msdppeerencapsulationstate", msdppeerentry.Msdppeerencapsulationstate}
    msdppeerentry.EntityData.Leafs["msdpPeerEncapsulationType"] = types.YLeaf{"Msdppeerencapsulationtype", msdppeerentry.Msdppeerencapsulationtype}
    msdppeerentry.EntityData.Leafs["msdpPeerConnectionAttempts"] = types.YLeaf{"Msdppeerconnectionattempts", msdppeerentry.Msdppeerconnectionattempts}
    msdppeerentry.EntityData.Leafs["msdpPeerInNotifications"] = types.YLeaf{"Msdppeerinnotifications", msdppeerentry.Msdppeerinnotifications}
    msdppeerentry.EntityData.Leafs["msdpPeerOutNotifications"] = types.YLeaf{"Msdppeeroutnotifications", msdppeerentry.Msdppeeroutnotifications}
    msdppeerentry.EntityData.Leafs["msdpPeerLastError"] = types.YLeaf{"Msdppeerlasterror", msdppeerentry.Msdppeerlasterror}
    return &(msdppeerentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry (conceptual row) representing an MSDP SA advert. The type is slice
    // of DRAFTMSDPMIB_Msdpsacachetable_Msdpsacacheentry.
    Msdpsacacheentry []DRAFTMSDPMIB_Msdpsacachetable_Msdpsacacheentry
}

func (msdpsacachetable *DRAFTMSDPMIB_Msdpsacachetable) GetEntityData() *types.CommonEntityData {
    msdpsacachetable.EntityData.YFilter = msdpsacachetable.YFilter
    msdpsacachetable.EntityData.YangName = "msdpSACacheTable"
    msdpsacachetable.EntityData.BundleName = "cisco_ios_xe"
    msdpsacachetable.EntityData.ParentYangName = "DRAFT-MSDP-MIB"
    msdpsacachetable.EntityData.SegmentPath = "msdpSACacheTable"
    msdpsacachetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    msdpsacachetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    msdpsacachetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    msdpsacachetable.EntityData.Children = make(map[string]types.YChild)
    msdpsacachetable.EntityData.Children["msdpSACacheEntry"] = types.YChild{"Msdpsacacheentry", nil}
    for i := range msdpsacachetable.Msdpsacacheentry {
        msdpsacachetable.EntityData.Children[types.GetSegmentPath(&msdpsacachetable.Msdpsacacheentry[i])] = types.YChild{"Msdpsacacheentry", &msdpsacachetable.Msdpsacacheentry[i]}
    }
    msdpsacachetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(msdpsacachetable.EntityData)
}

// DRAFTMSDPMIB_Msdpsacachetable_Msdpsacacheentry
// An entry (conceptual row) representing an MSDP SA
// advert.
type DRAFTMSDPMIB_Msdpsacachetable_Msdpsacacheentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The group address of the SA Cache entry. The type
    // is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Msdpsacachegroupaddr interface{}

    // This attribute is a key. The source address of the SA Cache entry. The type
    // is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Msdpsacachesourceaddr interface{}

    // This attribute is a key. The address of the RP which originated the last SA
    // message accepted for this entry. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Msdpsacacheoriginrp interface{}

    // The peer from which this SA Cache entry was last accepted.  This address
    // must correspond to the msdpPeerRemoteAddress value for a row in the MSDP
    // Peer Table. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Msdpsacachepeerlearnedfrom interface{}

    // The peer from which an SA message corresponding to this cache entry would
    // be accepted (i.e. the RPF peer for msdpSACacheOriginRP).  This may be
    // different than msdpSACachePeerLearnedFrom if this entry was created by an
    // MSDP SA-Response.  This address must correspond to the
    // msdpPeerRemoteAddress value for a row in the MSDP Peer Table. The type is
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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

func (msdpsacacheentry *DRAFTMSDPMIB_Msdpsacachetable_Msdpsacacheentry) GetEntityData() *types.CommonEntityData {
    msdpsacacheentry.EntityData.YFilter = msdpsacacheentry.YFilter
    msdpsacacheentry.EntityData.YangName = "msdpSACacheEntry"
    msdpsacacheentry.EntityData.BundleName = "cisco_ios_xe"
    msdpsacacheentry.EntityData.ParentYangName = "msdpSACacheTable"
    msdpsacacheentry.EntityData.SegmentPath = "msdpSACacheEntry" + "[msdpSACacheGroupAddr='" + fmt.Sprintf("%v", msdpsacacheentry.Msdpsacachegroupaddr) + "']" + "[msdpSACacheSourceAddr='" + fmt.Sprintf("%v", msdpsacacheentry.Msdpsacachesourceaddr) + "']" + "[msdpSACacheOriginRP='" + fmt.Sprintf("%v", msdpsacacheentry.Msdpsacacheoriginrp) + "']"
    msdpsacacheentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    msdpsacacheentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    msdpsacacheentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    msdpsacacheentry.EntityData.Children = make(map[string]types.YChild)
    msdpsacacheentry.EntityData.Leafs = make(map[string]types.YLeaf)
    msdpsacacheentry.EntityData.Leafs["msdpSACacheGroupAddr"] = types.YLeaf{"Msdpsacachegroupaddr", msdpsacacheentry.Msdpsacachegroupaddr}
    msdpsacacheentry.EntityData.Leafs["msdpSACacheSourceAddr"] = types.YLeaf{"Msdpsacachesourceaddr", msdpsacacheentry.Msdpsacachesourceaddr}
    msdpsacacheentry.EntityData.Leafs["msdpSACacheOriginRP"] = types.YLeaf{"Msdpsacacheoriginrp", msdpsacacheentry.Msdpsacacheoriginrp}
    msdpsacacheentry.EntityData.Leafs["msdpSACachePeerLearnedFrom"] = types.YLeaf{"Msdpsacachepeerlearnedfrom", msdpsacacheentry.Msdpsacachepeerlearnedfrom}
    msdpsacacheentry.EntityData.Leafs["msdpSACacheRPFPeer"] = types.YLeaf{"Msdpsacacherpfpeer", msdpsacacheentry.Msdpsacacherpfpeer}
    msdpsacacheentry.EntityData.Leafs["msdpSACacheInSAs"] = types.YLeaf{"Msdpsacacheinsas", msdpsacacheentry.Msdpsacacheinsas}
    msdpsacacheentry.EntityData.Leafs["msdpSACacheInDataPackets"] = types.YLeaf{"Msdpsacacheindatapackets", msdpsacacheentry.Msdpsacacheindatapackets}
    msdpsacacheentry.EntityData.Leafs["msdpSACacheUpTime"] = types.YLeaf{"Msdpsacacheuptime", msdpsacacheentry.Msdpsacacheuptime}
    msdpsacacheentry.EntityData.Leafs["msdpSACacheExpiryTime"] = types.YLeaf{"Msdpsacacheexpirytime", msdpsacacheentry.Msdpsacacheexpirytime}
    msdpsacacheentry.EntityData.Leafs["msdpSACacheStatus"] = types.YLeaf{"Msdpsacachestatus", msdpsacacheentry.Msdpsacachestatus}
    return &(msdpsacacheentry.EntityData)
}

