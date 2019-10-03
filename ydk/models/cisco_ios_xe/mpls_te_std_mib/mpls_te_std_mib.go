// Copyright (C) The Internet Society (2004). The
// initial version of this MIB module was published
// in RFC 3812. For full legal notices see the RFC
// itself or see: http://www.ietf.org/copyrights/ianamib.html
// 
// This MIB module contains managed object definitions
//  for MPLS Traffic Engineering (TE) as defined in:
// 1. Extensions to RSVP for LSP Tunnels, Awduche et
//  al, RFC 3209, December 2001
// 2. Constraint-Based LSP Setup using LDP, Jamoussi
//  (Editor), RFC 3212, January 2002
// 3. Requirements for Traffic Engineering Over MPLS,
//  Awduche, D., Malcolm, J., Agogbua, J., O'Dell, M.,
//  and J. McManus, [RFC2702], September 1999
package mpls_te_std_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package mpls_te_std_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:MPLS-TE-STD-MIB MPLS-TE-STD-MIB}", reflect.TypeOf(MPLSTESTDMIB{}))
    ydk.RegisterEntity("MPLS-TE-STD-MIB:MPLS-TE-STD-MIB", reflect.TypeOf(MPLSTESTDMIB{}))
}

// MPLSTESTDMIB
type MPLSTESTDMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    MplsTeScalars MPLSTESTDMIB_MplsTeScalars

    
    MplsTeObjects MPLSTESTDMIB_MplsTeObjects

    // The mplsTunnelTable allows new MPLS tunnels to be created between an LSR
    // and a remote endpoint, and existing tunnels to be reconfigured or removed.
    // Note that only point-to-point tunnel segments are supported, although
    // multipoint-to-point and point- to-multipoint connections are supported by
    // an LSR acting as a cross-connect.  Each MPLS tunnel can thus have one
    // out-segment originating at this LSR and/or one in-segment terminating at
    // this LSR.
    MplsTunnelTable MPLSTESTDMIB_MplsTunnelTable

    // The mplsTunnelHopTable is used to indicate the hops, strict or loose, for
    // an instance of an MPLS tunnel defined in mplsTunnelTable, when it is
    // established via signalling, for the outgoing direction of the tunnel. Thus
    // at a transit LSR, this table contains the desired path of the tunnel from
    // this LSR onwards. Each row in this table is indexed by
    // mplsTunnelHopListIndex which corresponds to a group of hop lists or path
    // options.  Each row also has a secondary index mplsTunnelHopIndex, which
    // indicates a group of hops (also known as a path option). Finally, the third
    // index, mplsTunnelHopIndex indicates the specific hop information for a path
    // option. In case we want to specify a particular interface on the
    // originating LSR of an outgoing tunnel by which we want packets to exit the
    // LSR, we specify this as the first hop for this tunnel in
    // mplsTunnelHopTable.
    MplsTunnelHopTable MPLSTESTDMIB_MplsTunnelHopTable

    // The mplsTunnelResourceTable allows a manager to specify which resources are
    // desired for an MPLS tunnel.  This table also allows several tunnels to
    // point to a single entry in this table, implying that these tunnels should
    // share resources.
    MplsTunnelResourceTable MPLSTESTDMIB_MplsTunnelResourceTable

    // The mplsTunnelARHopTable is used to indicate the hops for an MPLS tunnel
    // defined in mplsTunnelTable, as reported by the MPLS signalling protocol.
    // Thus at a transit LSR, this table (if the table is supported and if the
    // signaling protocol is recording actual route information) contains the
    // actual route of the whole tunnel. If the signaling protocol is not
    // recording the actual route, this table MAY report the information from the
    // mplsTunnelHopTable or the mplsTunnelCHopTable.  Each row in this table is
    // indexed by mplsTunnelARHopListIndex. Each row also has a secondary index
    // mplsTunnelARHopIndex, corresponding to the next hop that this row
    // corresponds to.  Please note that since the information necessary to build
    // entries within this table is not provided by some MPLS signalling
    // protocols, implementation of this table is optional. Furthermore, since the
    // information in this table is actually provided by the MPLS signalling
    // protocol after the path has been set-up, the entries in this table are
    // provided only for observation, and hence, all variables in this table are
    // accessible exclusively as read- only.  Note also that the contents of this
    // table may change while it is being read because of re-routing activities. A
    // network administrator may verify that the actual route read is consistent
    // by reference to the mplsTunnelLastPathChange object.
    MplsTunnelARHopTable MPLSTESTDMIB_MplsTunnelARHopTable

    // The mplsTunnelCHopTable is used to indicate the hops, strict or loose, for
    // an MPLS tunnel defined in mplsTunnelTable, as computed by a constraint-
    // based routing protocol, based on the mplsTunnelHopTable for the outgoing
    // direction of the tunnel. Thus at a transit LSR, this table (if the table is
    // supported) MAY contain the path computed by the CSPF engine on (or on
    // behalf of) this LSR. Each row in this table is indexed by
    // mplsTunnelCHopListIndex.  Each row also has a secondary index
    // mplsTunnelCHopIndex, corresponding to the next hop that this row
    // corresponds to. In case we want to specify a particular interface on the
    // originating LSR of an outgoing tunnel by which we want packets to exit the
    // LSR, we specify this as the first hop for this tunnel in
    // mplsTunnelCHopTable.  Please note that since the information necessary to
    // build entries within this table may not be supported by some LSRs,
    // implementation of this table is optional. Furthermore, since the
    // information in this table describes the path computed by the CSPF engine
    // the entries in this table are read-only.
    MplsTunnelCHopTable MPLSTESTDMIB_MplsTunnelCHopTable

    // The mplsTunnelCRLDPResTable allows a manager to specify which
    // CR-LDP-specific resources are desired for an MPLS tunnel if that tunnel is
    // signaled using CR-LDP. Note that these attributes are in addition to those
    // specified in mplsTunnelResourceTable. This table also allows several
    // tunnels to point to a single entry in this table, implying that these
    // tunnels should share resources.
    MplsTunnelCRLDPResTable MPLSTESTDMIB_MplsTunnelCRLDPResTable
}

func (mPLSTESTDMIB *MPLSTESTDMIB) GetEntityData() *types.CommonEntityData {
    mPLSTESTDMIB.EntityData.YFilter = mPLSTESTDMIB.YFilter
    mPLSTESTDMIB.EntityData.YangName = "MPLS-TE-STD-MIB"
    mPLSTESTDMIB.EntityData.BundleName = "cisco_ios_xe"
    mPLSTESTDMIB.EntityData.ParentYangName = "MPLS-TE-STD-MIB"
    mPLSTESTDMIB.EntityData.SegmentPath = "MPLS-TE-STD-MIB:MPLS-TE-STD-MIB"
    mPLSTESTDMIB.EntityData.AbsolutePath = mPLSTESTDMIB.EntityData.SegmentPath
    mPLSTESTDMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mPLSTESTDMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mPLSTESTDMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mPLSTESTDMIB.EntityData.Children = types.NewOrderedMap()
    mPLSTESTDMIB.EntityData.Children.Append("mplsTeScalars", types.YChild{"MplsTeScalars", &mPLSTESTDMIB.MplsTeScalars})
    mPLSTESTDMIB.EntityData.Children.Append("mplsTeObjects", types.YChild{"MplsTeObjects", &mPLSTESTDMIB.MplsTeObjects})
    mPLSTESTDMIB.EntityData.Children.Append("mplsTunnelTable", types.YChild{"MplsTunnelTable", &mPLSTESTDMIB.MplsTunnelTable})
    mPLSTESTDMIB.EntityData.Children.Append("mplsTunnelHopTable", types.YChild{"MplsTunnelHopTable", &mPLSTESTDMIB.MplsTunnelHopTable})
    mPLSTESTDMIB.EntityData.Children.Append("mplsTunnelResourceTable", types.YChild{"MplsTunnelResourceTable", &mPLSTESTDMIB.MplsTunnelResourceTable})
    mPLSTESTDMIB.EntityData.Children.Append("mplsTunnelARHopTable", types.YChild{"MplsTunnelARHopTable", &mPLSTESTDMIB.MplsTunnelARHopTable})
    mPLSTESTDMIB.EntityData.Children.Append("mplsTunnelCHopTable", types.YChild{"MplsTunnelCHopTable", &mPLSTESTDMIB.MplsTunnelCHopTable})
    mPLSTESTDMIB.EntityData.Children.Append("mplsTunnelCRLDPResTable", types.YChild{"MplsTunnelCRLDPResTable", &mPLSTESTDMIB.MplsTunnelCRLDPResTable})
    mPLSTESTDMIB.EntityData.Leafs = types.NewOrderedMap()

    mPLSTESTDMIB.EntityData.YListKeys = []string {}

    return &(mPLSTESTDMIB.EntityData)
}

// MPLSTESTDMIB_MplsTeScalars
type MPLSTESTDMIB_MplsTeScalars struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The number of tunnels configured on this device. A tunnel is considered
    // configured if the mplsTunnelRowStatus is active(1). The type is interface{}
    // with range: 0..4294967295.
    MplsTunnelConfigured interface{}

    // The number of tunnels active on this device. A tunnel is considered active
    // if the mplsTunnelOperStatus is up(1). The type is interface{} with range:
    // 0..4294967295.
    MplsTunnelActive interface{}

    // The traffic engineering distribution protocol(s) used by this LSR. Note
    // that an LSR may support more than one distribution protocol simultaneously.
    // The type is map[string]bool.
    MplsTunnelTEDistProto interface{}

    // The maximum number of hops that can be specified for a tunnel on this
    // device. The type is interface{} with range: 0..4294967295.
    MplsTunnelMaxHops interface{}

    // This variable indicates the maximum number of notifications issued per
    // second. If events occur more rapidly, the implementation may simply fail to
    // emit these notifications during that period, or may queue them until an
    // appropriate time. A value of 0 means no throttling is applied and events
    // may be notified at the rate at which they occur. The type is interface{}
    // with range: 0..4294967295.
    MplsTunnelNotificationMaxRate interface{}
}

func (mplsTeScalars *MPLSTESTDMIB_MplsTeScalars) GetEntityData() *types.CommonEntityData {
    mplsTeScalars.EntityData.YFilter = mplsTeScalars.YFilter
    mplsTeScalars.EntityData.YangName = "mplsTeScalars"
    mplsTeScalars.EntityData.BundleName = "cisco_ios_xe"
    mplsTeScalars.EntityData.ParentYangName = "MPLS-TE-STD-MIB"
    mplsTeScalars.EntityData.SegmentPath = "mplsTeScalars"
    mplsTeScalars.EntityData.AbsolutePath = "MPLS-TE-STD-MIB:MPLS-TE-STD-MIB/" + mplsTeScalars.EntityData.SegmentPath
    mplsTeScalars.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsTeScalars.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsTeScalars.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsTeScalars.EntityData.Children = types.NewOrderedMap()
    mplsTeScalars.EntityData.Leafs = types.NewOrderedMap()
    mplsTeScalars.EntityData.Leafs.Append("mplsTunnelConfigured", types.YLeaf{"MplsTunnelConfigured", mplsTeScalars.MplsTunnelConfigured})
    mplsTeScalars.EntityData.Leafs.Append("mplsTunnelActive", types.YLeaf{"MplsTunnelActive", mplsTeScalars.MplsTunnelActive})
    mplsTeScalars.EntityData.Leafs.Append("mplsTunnelTEDistProto", types.YLeaf{"MplsTunnelTEDistProto", mplsTeScalars.MplsTunnelTEDistProto})
    mplsTeScalars.EntityData.Leafs.Append("mplsTunnelMaxHops", types.YLeaf{"MplsTunnelMaxHops", mplsTeScalars.MplsTunnelMaxHops})
    mplsTeScalars.EntityData.Leafs.Append("mplsTunnelNotificationMaxRate", types.YLeaf{"MplsTunnelNotificationMaxRate", mplsTeScalars.MplsTunnelNotificationMaxRate})

    mplsTeScalars.EntityData.YListKeys = []string {}

    return &(mplsTeScalars.EntityData)
}

// MPLSTESTDMIB_MplsTeObjects
type MPLSTESTDMIB_MplsTeObjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object contains an unused value for mplsTunnelIndex, or a zero to
    // indicate that none exist. Negative values are not allowed, as they do not
    // correspond to valid values of mplsTunnelIndex.  Note that this object
    // offers an unused value for an mplsTunnelIndex value at the ingress side of
    // a tunnel. At other LSRs the value of mplsTunnelIndex SHOULD be taken from
    // the value signaled by the MPLS signaling protocol. The type is interface{}
    // with range: 0..65535.
    MplsTunnelIndexNext interface{}

    // This object contains an appropriate value to be used for
    // mplsTunnelHopListIndex when creating entries in the mplsTunnelHopTable.  If
    // the number of unassigned entries is exhausted, a retrieval operation will
    // return a value of 0.  This object may also return a value of 0 when the LSR
    // is unable to accept conceptual row creation, for example, if the
    // mplsTunnelHopTable is implemented as read-only. To obtain the value of
    // mplsTunnelHopListIndex for a new entry in the mplsTunnelHopTable, the
    // manager issues a management protocol retrieval operation to obtain the
    // current value of mplsTunnelHopIndex.  When the SET is performed to create a
    // row in the mplsTunnelHopTable, the Command Responder (agent) must determine
    // whether the value is indeed still unused; Two Network Management
    // Applications may attempt to create a row (configuration entry)
    // simultaneously and use the same value. If it is currently unused, the SET
    // succeeds and the Command Responder (agent) changes the value of this
    // object, according to an implementation-specific algorithm. If the value is
    // in use, however, the SET fails.  The Network Management Application must
    // then re-read this variable to obtain a new usable value. The type is
    // interface{} with range: 0..4294967295.
    MplsTunnelHopListIndexNext interface{}

    // This object contains the next appropriate value to be used for
    // mplsTunnelResourceIndex when creating entries in the
    // mplsTunnelResourceTable. If the number of unassigned entries is exhausted,
    // a retrieval operation will return a value of 0.  This object may also
    // return a value of 0 when the LSR is unable to accept conceptual row
    // creation, for example, if the mplsTunnelTable is implemented as read-only. 
    // To obtain the mplsTunnelResourceIndex value for a new entry, the manager
    // must first issue a management protocol retrieval operation to obtain the
    // current value of this object.  When the SET is performed to create a row in
    // the mplsTunnelResourceTable, the Command Responder (agent) must determine
    // whether the value is indeed still unused; Two Network Management
    // Applications may attempt to create a row (configuration entry)
    // simultaneously and use the same value. If it is currently unused, the SET
    // succeeds and the Command Responder (agent) changes the value of this
    // object, according to an implementation-specific algorithm. If the value is
    // in use, however, the SET fails.  The Network Management Application must
    // then re-read this variable to obtain a new usable value. The type is
    // interface{} with range: 0..2147483647.
    MplsTunnelResourceIndexNext interface{}

    // If this object is true, then it enables the generation of mplsTunnelUp and
    // mplsTunnelDown traps, otherwise these traps are not emitted. The type is
    // bool.
    MplsTunnelNotificationEnable interface{}
}

func (mplsTeObjects *MPLSTESTDMIB_MplsTeObjects) GetEntityData() *types.CommonEntityData {
    mplsTeObjects.EntityData.YFilter = mplsTeObjects.YFilter
    mplsTeObjects.EntityData.YangName = "mplsTeObjects"
    mplsTeObjects.EntityData.BundleName = "cisco_ios_xe"
    mplsTeObjects.EntityData.ParentYangName = "MPLS-TE-STD-MIB"
    mplsTeObjects.EntityData.SegmentPath = "mplsTeObjects"
    mplsTeObjects.EntityData.AbsolutePath = "MPLS-TE-STD-MIB:MPLS-TE-STD-MIB/" + mplsTeObjects.EntityData.SegmentPath
    mplsTeObjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsTeObjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsTeObjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsTeObjects.EntityData.Children = types.NewOrderedMap()
    mplsTeObjects.EntityData.Leafs = types.NewOrderedMap()
    mplsTeObjects.EntityData.Leafs.Append("mplsTunnelIndexNext", types.YLeaf{"MplsTunnelIndexNext", mplsTeObjects.MplsTunnelIndexNext})
    mplsTeObjects.EntityData.Leafs.Append("mplsTunnelHopListIndexNext", types.YLeaf{"MplsTunnelHopListIndexNext", mplsTeObjects.MplsTunnelHopListIndexNext})
    mplsTeObjects.EntityData.Leafs.Append("mplsTunnelResourceIndexNext", types.YLeaf{"MplsTunnelResourceIndexNext", mplsTeObjects.MplsTunnelResourceIndexNext})
    mplsTeObjects.EntityData.Leafs.Append("mplsTunnelNotificationEnable", types.YLeaf{"MplsTunnelNotificationEnable", mplsTeObjects.MplsTunnelNotificationEnable})

    mplsTeObjects.EntityData.YListKeys = []string {}

    return &(mplsTeObjects.EntityData)
}

// MPLSTESTDMIB_MplsTunnelTable
// The mplsTunnelTable allows new MPLS tunnels to be
// created between an LSR and a remote endpoint, and
// existing tunnels to be reconfigured or removed.
// Note that only point-to-point tunnel segments are
// supported, although multipoint-to-point and point-
// to-multipoint connections are supported by an LSR
// acting as a cross-connect.  Each MPLS tunnel can
// thus have one out-segment originating at this LSR
// and/or one in-segment terminating at this LSR.
type MPLSTESTDMIB_MplsTunnelTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table represents an MPLS tunnel. An entry can be created
    // by a network administrator or by an SNMP agent as instructed by an MPLS
    // signalling protocol. Whenever a new entry is created with mplsTunnelIsIf
    // set to true(1), then a corresponding entry is created in ifTable as well
    // (see RFC 2863). The ifType of this entry is mplsTunnel(150).  A tunnel
    // entry needs to be uniquely identified across a MPLS network. Indices
    // mplsTunnelIndex and mplsTunnelInstance uniquely identify a tunnel on the
    // LSR originating the tunnel.  To uniquely identify a tunnel across an MPLS
    // network requires index mplsTunnelIngressLSRId.  The last index
    // mplsTunnelEgressLSRId is useful in identifying all instances of a tunnel
    // that terminate on the same egress LSR. The type is slice of
    // MPLSTESTDMIB_MplsTunnelTable_MplsTunnelEntry.
    MplsTunnelEntry []*MPLSTESTDMIB_MplsTunnelTable_MplsTunnelEntry
}

func (mplsTunnelTable *MPLSTESTDMIB_MplsTunnelTable) GetEntityData() *types.CommonEntityData {
    mplsTunnelTable.EntityData.YFilter = mplsTunnelTable.YFilter
    mplsTunnelTable.EntityData.YangName = "mplsTunnelTable"
    mplsTunnelTable.EntityData.BundleName = "cisco_ios_xe"
    mplsTunnelTable.EntityData.ParentYangName = "MPLS-TE-STD-MIB"
    mplsTunnelTable.EntityData.SegmentPath = "mplsTunnelTable"
    mplsTunnelTable.EntityData.AbsolutePath = "MPLS-TE-STD-MIB:MPLS-TE-STD-MIB/" + mplsTunnelTable.EntityData.SegmentPath
    mplsTunnelTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsTunnelTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsTunnelTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsTunnelTable.EntityData.Children = types.NewOrderedMap()
    mplsTunnelTable.EntityData.Children.Append("mplsTunnelEntry", types.YChild{"MplsTunnelEntry", nil})
    for i := range mplsTunnelTable.MplsTunnelEntry {
        mplsTunnelTable.EntityData.Children.Append(types.GetSegmentPath(mplsTunnelTable.MplsTunnelEntry[i]), types.YChild{"MplsTunnelEntry", mplsTunnelTable.MplsTunnelEntry[i]})
    }
    mplsTunnelTable.EntityData.Leafs = types.NewOrderedMap()

    mplsTunnelTable.EntityData.YListKeys = []string {}

    return &(mplsTunnelTable.EntityData)
}

// MPLSTESTDMIB_MplsTunnelTable_MplsTunnelEntry
// An entry in this table represents an MPLS tunnel.
// An entry can be created by a network administrator
// or by an SNMP agent as instructed by an MPLS
// signalling protocol. Whenever a new entry is
// created with mplsTunnelIsIf set to true(1), then a
// corresponding entry is created in ifTable as well
// (see RFC 2863). The ifType of this entry is
// mplsTunnel(150).
// 
// A tunnel entry needs to be uniquely identified across
// a MPLS network. Indices mplsTunnelIndex and
// mplsTunnelInstance uniquely identify a tunnel on
// the LSR originating the tunnel.  To uniquely
// identify a tunnel across an MPLS network requires
// index mplsTunnelIngressLSRId.  The last index
// mplsTunnelEgressLSRId is useful in identifying all
// instances of a tunnel that terminate on the same
// egress LSR.
type MPLSTESTDMIB_MplsTunnelTable_MplsTunnelEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Uniquely identifies a set of tunnel instances
    // between a pair of ingress and egress LSRs. Managers should obtain new
    // values for row creation in this table by reading mplsTunnelIndexNext. When
    // the MPLS signalling protocol is rsvp(2) this value SHOULD be equal to the
    // value signaled in the Tunnel Id of the Session object. When the MPLS
    // signalling protocol is crldp(3) this value SHOULD be equal to the value
    // signaled in the LSP ID. The type is interface{} with range: 0..65535.
    MplsTunnelIndex interface{}

    // This attribute is a key. Uniquely identifies a particular instance of a
    // tunnel between a pair of ingress and egress LSRs. It is useful to identify
    // multiple instances of tunnels for the purposes of backup and parallel
    // tunnels. When the MPLS signaling protocol is rsvp(2) this value SHOULD be
    // equal to the LSP Id of the Sender Template object. When the signaling
    // protocol is crldp(3) there is no equivalent signaling object. The type is
    // interface{} with range: 0..4294967295.
    MplsTunnelInstance interface{}

    // This attribute is a key. Identity of the ingress LSR associated with this
    // tunnel instance. When the MPLS signalling protocol is rsvp(2) this value
    // SHOULD be equal to the Tunnel Sender Address in the Sender Template object
    // and MAY be equal to the Extended Tunnel Id field in the SESSION object.
    // When the MPLS signalling protocol is crldp(3) this value SHOULD be equal to
    // the Ingress LSR Router ID field in the LSPID TLV object. The type is
    // interface{} with range: 0..4294967295.
    MplsTunnelIngressLSRId interface{}

    // This attribute is a key. Identity of the egress LSR associated with this
    // tunnel instance. The type is interface{} with range: 0..4294967295.
    MplsTunnelEgressLSRId interface{}

    // The canonical name assigned to the tunnel. This name can be used to refer
    // to the tunnel on the LSR's console port.  If mplsTunnelIsIf is set to true
    // then the ifName of the interface corresponding to this tunnel should have a
    // value equal to mplsTunnelName.  Also see the description of ifName in RFC
    // 2863. The type is string.
    MplsTunnelName interface{}

    // A textual string containing information about the tunnel.  If there is no
    // description this object contains a zero length string. This object is may
    // not be signaled by MPLS signaling protocols, consequentally the value of
    // this object at transit and egress LSRs MAY be automatically generated or
    // absent. The type is string.
    MplsTunnelDescr interface{}

    // Denotes whether or not this tunnel corresponds to an interface represented
    // in the interfaces group table. Note that if this variable is set to true
    // then the ifName of the interface corresponding to this tunnel should have a
    // value equal to mplsTunnelName.  Also see the description of ifName in RFC
    // 2863.  This object is meaningful only at the ingress and egress LSRs. The
    // type is bool.
    MplsTunnelIsIf interface{}

    // If mplsTunnelIsIf is set to true, then this value contains the LSR-assigned
    // ifIndex which corresponds to an entry in the interfaces table.  Otherwise
    // this variable should contain the value of zero indicating that a valid
    // ifIndex was not assigned to this tunnel interface. The type is interface{}
    // with range: 0..2147483647.
    MplsTunnelIfIndex interface{}

    // Denotes the entity that created and is responsible for managing this
    // tunnel. This column is automatically filled by the agent on creation of a
    // row. The type is MplsOwner.
    MplsTunnelOwner interface{}

    // This value signifies the role that this tunnel entry/instance represents.
    // This value MUST be set to head(1) at the originating point of the tunnel.
    // This value MUST be set to transit(2) at transit points along the tunnel, if
    // transit points are supported. This value MUST be set to tail(3) at the
    // terminating point of the tunnel if tunnel tails are supported.  The value
    // headTail(4) is provided for tunnels that begin and end on the same LSR. The
    // type is MplsTunnelRole.
    MplsTunnelRole interface{}

    // This variable points to a row in the mplsXCTable. This table identifies the
    // segments that compose this tunnel, their characteristics, and relationships
    // to each other. A value of zeroDotZero indicates that no LSP has been
    // associated with this tunnel yet. The type is string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    MplsTunnelXCPointer interface{}

    // The signalling protocol, if any, used to setup this tunnel. The type is
    // MplsTunnelSignallingProto.
    MplsTunnelSignallingProto interface{}

    // Indicates the setup priority of this tunnel. The type is interface{} with
    // range: 0..7.
    MplsTunnelSetupPrio interface{}

    // Indicates the holding priority for this tunnel. The type is interface{}
    // with range: 0..7.
    MplsTunnelHoldingPrio interface{}

    // This bit mask indicates optional session values for this tunnel. The
    // following describes these bit fields:  fastRerouteThis flag indicates that
    // the any tunnel hop may choose to reroute this tunnel without tearing it
    // down.  This flag permits transit routers to use a local repair mechanism
    // which may result in violation of the explicit routing of this tunnel. When
    // a fault is detected on an adjacent downstream link or node, a transit
    // router can re-route traffic for fast service restoration.  mergingPermitted
    // This flag permits transit routers to merge this session with other RSVP
    // sessions for the purpose of reducing resource overhead on downstream
    // transit routers, thereby providing better network scaling.  isPersistent 
    // Indicates whether this tunnel should be restored automatically after a
    // failure occurs.  isPinned   This flag indicates whether the loose- routed
    // hops of this tunnel are to be pinned.  recordRouteThis flag indicates
    // whether or not the signalling protocol should remember the tunnel path
    // after it has been signaled. The type is map[string]bool.
    MplsTunnelSessionAttributes interface{}

    // Indicates that the local repair mechanism is in use to maintain this tunnel
    // (usually in the face of an outage of the link it was previously routed
    // over). The type is bool.
    MplsTunnelLocalProtectInUse interface{}

    // This variable represents a pointer to the traffic parameter specification
    // for this tunnel.  This value may point at an entry in the
    // mplsTunnelResourceEntry to indicate which mplsTunnelResourceEntry is to be
    // assigned to this LSP instance.  This value may optionally point at an
    // externally defined traffic parameter specification table.  A value of
    // zeroDotZero indicates best-effort treatment.  By having the same value of
    // this object, two or more LSPs can indicate resource sharing. The type is
    // string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    MplsTunnelResourcePointer interface{}

    // Specifies the instance index of the primary instance of this tunnel. More
    // details of the definition of tunnel instances and the primary tunnel
    // instance can be found in the description of the TEXTUAL-CONVENTION
    // MplsTunnelInstanceIndex. The type is interface{} with range: 0..4294967295.
    MplsTunnelPrimaryInstance interface{}

    // This value indicates which priority, in descending order, with 0 indicating
    // the lowest priority, within a group of tunnel instances. A group of tunnel
    // instances is defined as a set of LSPs with the same mplsTunnelIndex in this
    // table, but with a different mplsTunnelInstance. Tunnel instance priorities
    // are used to denote the priority at which a particular tunnel instance will
    // supercede another. Instances of tunnels containing the same
    // mplsTunnelInstancePriority will be used for load sharing. The type is
    // interface{} with range: 0..4294967295.
    MplsTunnelInstancePriority interface{}

    // Index into the mplsTunnelHopTable entry that specifies the explicit route
    // hops for this tunnel. This object is meaningful only at the head-end of the
    // tunnel. The type is interface{} with range: 0..4294967295.
    MplsTunnelHopTableIndex interface{}

    // This value denotes the configured path that was chosen for this tunnel.
    // This value reflects the secondary index into mplsTunnelHopTable. This path
    // may not exactly match the one in mplsTunnelARHopTable due to the fact that
    // some CSPF modification may have taken place. See mplsTunnelARHopTable for
    // the actual path being taken by the tunnel. A value of zero denotes that no
    // path is currently in use or available. The type is interface{} with range:
    // 0..4294967295.
    MplsTunnelPathInUse interface{}

    // Index into the mplsTunnelARHopTable entry that specifies the actual hops
    // traversed by the tunnel. This is automatically updated by the agent when
    // the actual hops becomes available. The type is interface{} with range:
    // 0..4294967295.
    MplsTunnelARHopTableIndex interface{}

    // Index into the mplsTunnelCHopTable entry that specifies the computed hops
    // traversed by the tunnel. This is automatically updated by the agent when
    // computed hops become available or when computed hops get modified. The type
    // is interface{} with range: 0..4294967295.
    MplsTunnelCHopTableIndex interface{}

    // A link satisfies the include-any constraint if and only if the constraint
    // is zero, or the link and the constraint have a resource class in common.
    // The type is interface{} with range: 0..4294967295.
    MplsTunnelIncludeAnyAffinity interface{}

    // A link satisfies the include-all constraint if and only if the link
    // contains all of the administrative groups specified in the constraint. The
    // type is interface{} with range: 0..4294967295.
    MplsTunnelIncludeAllAffinity interface{}

    // A link satisfies the exclude-any constraint if and only if the link
    // contains none of the administrative groups specified in the constraint. The
    // type is interface{} with range: 0..4294967295.
    MplsTunnelExcludeAnyAffinity interface{}

    // This value represents the aggregate up time for all instances of this
    // tunnel, if available. If this value is unavailable, it MUST return a value
    // of 0. The type is interface{} with range: 0..4294967295.
    MplsTunnelTotalUpTime interface{}

    // This value identifies the total time that this tunnel instance's operStatus
    // has been Up(1). The type is interface{} with range: 0..4294967295.
    MplsTunnelInstanceUpTime interface{}

    // Specifies the total time the primary instance of this tunnel has been
    // active. The primary instance of this tunnel is defined in
    // mplsTunnelPrimaryInstance. The type is interface{} with range:
    // 0..4294967295.
    MplsTunnelPrimaryUpTime interface{}

    // Specifies the number of times the actual path for this tunnel instance has
    // changed. The type is interface{} with range: 0..4294967295.
    MplsTunnelPathChanges interface{}

    // Specifies the time since the last change to the actual path for this tunnel
    // instance. The type is interface{} with range: 0..4294967295.
    MplsTunnelLastPathChange interface{}

    // Specifies the value of SysUpTime when the first instance of this tunnel
    // came into existence. That is, when the value of mplsTunnelOperStatus was
    // first set to up(1). The type is interface{} with range: 0..4294967295.
    MplsTunnelCreationTime interface{}

    // Specifies the number of times the state (mplsTunnelOperStatus) of this
    // tunnel instance has changed. The type is interface{} with range:
    // 0..4294967295.
    MplsTunnelStateTransitions interface{}

    // Indicates the desired operational status of this tunnel. The type is
    // MplsTunnelAdminStatus.
    MplsTunnelAdminStatus interface{}

    // Indicates the actual operational status of this tunnel, which is typically
    // but not limited to, a function of the state of individual segments of this
    // tunnel. The type is MplsTunnelOperStatus.
    MplsTunnelOperStatus interface{}

    // This variable is used to create, modify, and/or delete a row in this table.
    // When a row in this table is in active(1) state, no objects in that row can
    // be modified by the agent except mplsTunnelAdminStatus, mplsTunnelRowStatus
    // and mplsTunnelStorageType. The type is RowStatus.
    MplsTunnelRowStatus interface{}

    // The storage type for this tunnel entry. Conceptual rows having the value
    // 'permanent' need not allow write-access to any columnar objects in the row.
    // The type is StorageType.
    MplsTunnelStorageType interface{}

    // Number of packets forwarded by the tunnel. This object should represents
    // the 32-bit value of the least significant part of the 64-bit value if both
    // mplsTunnelPerfHCPackets is returned. The type is interface{} with range:
    // 0..4294967295.
    MplsTunnelPerfPackets interface{}

    // High capacity counter for number of packets forwarded by the tunnel. . The
    // type is interface{} with range: 0..18446744073709551615.
    MplsTunnelPerfHCPackets interface{}

    // Number of packets dropped because of errors or for other reasons. The type
    // is interface{} with range: 0..4294967295.
    MplsTunnelPerfErrors interface{}

    // Number of bytes forwarded by the tunnel. This object should represents the
    // 32-bit value of the least significant part of the 64-bit value if both
    // mplsTunnelPerfHCBytes is returned. The type is interface{} with range:
    // 0..4294967295.
    MplsTunnelPerfBytes interface{}

    // High capacity counter for number of bytes forwarded by the tunnel. The type
    // is interface{} with range: 0..18446744073709551615.
    MplsTunnelPerfHCBytes interface{}
}

func (mplsTunnelEntry *MPLSTESTDMIB_MplsTunnelTable_MplsTunnelEntry) GetEntityData() *types.CommonEntityData {
    mplsTunnelEntry.EntityData.YFilter = mplsTunnelEntry.YFilter
    mplsTunnelEntry.EntityData.YangName = "mplsTunnelEntry"
    mplsTunnelEntry.EntityData.BundleName = "cisco_ios_xe"
    mplsTunnelEntry.EntityData.ParentYangName = "mplsTunnelTable"
    mplsTunnelEntry.EntityData.SegmentPath = "mplsTunnelEntry" + types.AddKeyToken(mplsTunnelEntry.MplsTunnelIndex, "mplsTunnelIndex") + types.AddKeyToken(mplsTunnelEntry.MplsTunnelInstance, "mplsTunnelInstance") + types.AddKeyToken(mplsTunnelEntry.MplsTunnelIngressLSRId, "mplsTunnelIngressLSRId") + types.AddKeyToken(mplsTunnelEntry.MplsTunnelEgressLSRId, "mplsTunnelEgressLSRId")
    mplsTunnelEntry.EntityData.AbsolutePath = "MPLS-TE-STD-MIB:MPLS-TE-STD-MIB/mplsTunnelTable/" + mplsTunnelEntry.EntityData.SegmentPath
    mplsTunnelEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsTunnelEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsTunnelEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsTunnelEntry.EntityData.Children = types.NewOrderedMap()
    mplsTunnelEntry.EntityData.Leafs = types.NewOrderedMap()
    mplsTunnelEntry.EntityData.Leafs.Append("mplsTunnelIndex", types.YLeaf{"MplsTunnelIndex", mplsTunnelEntry.MplsTunnelIndex})
    mplsTunnelEntry.EntityData.Leafs.Append("mplsTunnelInstance", types.YLeaf{"MplsTunnelInstance", mplsTunnelEntry.MplsTunnelInstance})
    mplsTunnelEntry.EntityData.Leafs.Append("mplsTunnelIngressLSRId", types.YLeaf{"MplsTunnelIngressLSRId", mplsTunnelEntry.MplsTunnelIngressLSRId})
    mplsTunnelEntry.EntityData.Leafs.Append("mplsTunnelEgressLSRId", types.YLeaf{"MplsTunnelEgressLSRId", mplsTunnelEntry.MplsTunnelEgressLSRId})
    mplsTunnelEntry.EntityData.Leafs.Append("mplsTunnelName", types.YLeaf{"MplsTunnelName", mplsTunnelEntry.MplsTunnelName})
    mplsTunnelEntry.EntityData.Leafs.Append("mplsTunnelDescr", types.YLeaf{"MplsTunnelDescr", mplsTunnelEntry.MplsTunnelDescr})
    mplsTunnelEntry.EntityData.Leafs.Append("mplsTunnelIsIf", types.YLeaf{"MplsTunnelIsIf", mplsTunnelEntry.MplsTunnelIsIf})
    mplsTunnelEntry.EntityData.Leafs.Append("mplsTunnelIfIndex", types.YLeaf{"MplsTunnelIfIndex", mplsTunnelEntry.MplsTunnelIfIndex})
    mplsTunnelEntry.EntityData.Leafs.Append("mplsTunnelOwner", types.YLeaf{"MplsTunnelOwner", mplsTunnelEntry.MplsTunnelOwner})
    mplsTunnelEntry.EntityData.Leafs.Append("mplsTunnelRole", types.YLeaf{"MplsTunnelRole", mplsTunnelEntry.MplsTunnelRole})
    mplsTunnelEntry.EntityData.Leafs.Append("mplsTunnelXCPointer", types.YLeaf{"MplsTunnelXCPointer", mplsTunnelEntry.MplsTunnelXCPointer})
    mplsTunnelEntry.EntityData.Leafs.Append("mplsTunnelSignallingProto", types.YLeaf{"MplsTunnelSignallingProto", mplsTunnelEntry.MplsTunnelSignallingProto})
    mplsTunnelEntry.EntityData.Leafs.Append("mplsTunnelSetupPrio", types.YLeaf{"MplsTunnelSetupPrio", mplsTunnelEntry.MplsTunnelSetupPrio})
    mplsTunnelEntry.EntityData.Leafs.Append("mplsTunnelHoldingPrio", types.YLeaf{"MplsTunnelHoldingPrio", mplsTunnelEntry.MplsTunnelHoldingPrio})
    mplsTunnelEntry.EntityData.Leafs.Append("mplsTunnelSessionAttributes", types.YLeaf{"MplsTunnelSessionAttributes", mplsTunnelEntry.MplsTunnelSessionAttributes})
    mplsTunnelEntry.EntityData.Leafs.Append("mplsTunnelLocalProtectInUse", types.YLeaf{"MplsTunnelLocalProtectInUse", mplsTunnelEntry.MplsTunnelLocalProtectInUse})
    mplsTunnelEntry.EntityData.Leafs.Append("mplsTunnelResourcePointer", types.YLeaf{"MplsTunnelResourcePointer", mplsTunnelEntry.MplsTunnelResourcePointer})
    mplsTunnelEntry.EntityData.Leafs.Append("mplsTunnelPrimaryInstance", types.YLeaf{"MplsTunnelPrimaryInstance", mplsTunnelEntry.MplsTunnelPrimaryInstance})
    mplsTunnelEntry.EntityData.Leafs.Append("mplsTunnelInstancePriority", types.YLeaf{"MplsTunnelInstancePriority", mplsTunnelEntry.MplsTunnelInstancePriority})
    mplsTunnelEntry.EntityData.Leafs.Append("mplsTunnelHopTableIndex", types.YLeaf{"MplsTunnelHopTableIndex", mplsTunnelEntry.MplsTunnelHopTableIndex})
    mplsTunnelEntry.EntityData.Leafs.Append("mplsTunnelPathInUse", types.YLeaf{"MplsTunnelPathInUse", mplsTunnelEntry.MplsTunnelPathInUse})
    mplsTunnelEntry.EntityData.Leafs.Append("mplsTunnelARHopTableIndex", types.YLeaf{"MplsTunnelARHopTableIndex", mplsTunnelEntry.MplsTunnelARHopTableIndex})
    mplsTunnelEntry.EntityData.Leafs.Append("mplsTunnelCHopTableIndex", types.YLeaf{"MplsTunnelCHopTableIndex", mplsTunnelEntry.MplsTunnelCHopTableIndex})
    mplsTunnelEntry.EntityData.Leafs.Append("mplsTunnelIncludeAnyAffinity", types.YLeaf{"MplsTunnelIncludeAnyAffinity", mplsTunnelEntry.MplsTunnelIncludeAnyAffinity})
    mplsTunnelEntry.EntityData.Leafs.Append("mplsTunnelIncludeAllAffinity", types.YLeaf{"MplsTunnelIncludeAllAffinity", mplsTunnelEntry.MplsTunnelIncludeAllAffinity})
    mplsTunnelEntry.EntityData.Leafs.Append("mplsTunnelExcludeAnyAffinity", types.YLeaf{"MplsTunnelExcludeAnyAffinity", mplsTunnelEntry.MplsTunnelExcludeAnyAffinity})
    mplsTunnelEntry.EntityData.Leafs.Append("mplsTunnelTotalUpTime", types.YLeaf{"MplsTunnelTotalUpTime", mplsTunnelEntry.MplsTunnelTotalUpTime})
    mplsTunnelEntry.EntityData.Leafs.Append("mplsTunnelInstanceUpTime", types.YLeaf{"MplsTunnelInstanceUpTime", mplsTunnelEntry.MplsTunnelInstanceUpTime})
    mplsTunnelEntry.EntityData.Leafs.Append("mplsTunnelPrimaryUpTime", types.YLeaf{"MplsTunnelPrimaryUpTime", mplsTunnelEntry.MplsTunnelPrimaryUpTime})
    mplsTunnelEntry.EntityData.Leafs.Append("mplsTunnelPathChanges", types.YLeaf{"MplsTunnelPathChanges", mplsTunnelEntry.MplsTunnelPathChanges})
    mplsTunnelEntry.EntityData.Leafs.Append("mplsTunnelLastPathChange", types.YLeaf{"MplsTunnelLastPathChange", mplsTunnelEntry.MplsTunnelLastPathChange})
    mplsTunnelEntry.EntityData.Leafs.Append("mplsTunnelCreationTime", types.YLeaf{"MplsTunnelCreationTime", mplsTunnelEntry.MplsTunnelCreationTime})
    mplsTunnelEntry.EntityData.Leafs.Append("mplsTunnelStateTransitions", types.YLeaf{"MplsTunnelStateTransitions", mplsTunnelEntry.MplsTunnelStateTransitions})
    mplsTunnelEntry.EntityData.Leafs.Append("mplsTunnelAdminStatus", types.YLeaf{"MplsTunnelAdminStatus", mplsTunnelEntry.MplsTunnelAdminStatus})
    mplsTunnelEntry.EntityData.Leafs.Append("mplsTunnelOperStatus", types.YLeaf{"MplsTunnelOperStatus", mplsTunnelEntry.MplsTunnelOperStatus})
    mplsTunnelEntry.EntityData.Leafs.Append("mplsTunnelRowStatus", types.YLeaf{"MplsTunnelRowStatus", mplsTunnelEntry.MplsTunnelRowStatus})
    mplsTunnelEntry.EntityData.Leafs.Append("mplsTunnelStorageType", types.YLeaf{"MplsTunnelStorageType", mplsTunnelEntry.MplsTunnelStorageType})
    mplsTunnelEntry.EntityData.Leafs.Append("mplsTunnelPerfPackets", types.YLeaf{"MplsTunnelPerfPackets", mplsTunnelEntry.MplsTunnelPerfPackets})
    mplsTunnelEntry.EntityData.Leafs.Append("mplsTunnelPerfHCPackets", types.YLeaf{"MplsTunnelPerfHCPackets", mplsTunnelEntry.MplsTunnelPerfHCPackets})
    mplsTunnelEntry.EntityData.Leafs.Append("mplsTunnelPerfErrors", types.YLeaf{"MplsTunnelPerfErrors", mplsTunnelEntry.MplsTunnelPerfErrors})
    mplsTunnelEntry.EntityData.Leafs.Append("mplsTunnelPerfBytes", types.YLeaf{"MplsTunnelPerfBytes", mplsTunnelEntry.MplsTunnelPerfBytes})
    mplsTunnelEntry.EntityData.Leafs.Append("mplsTunnelPerfHCBytes", types.YLeaf{"MplsTunnelPerfHCBytes", mplsTunnelEntry.MplsTunnelPerfHCBytes})

    mplsTunnelEntry.EntityData.YListKeys = []string {"MplsTunnelIndex", "MplsTunnelInstance", "MplsTunnelIngressLSRId", "MplsTunnelEgressLSRId"}

    return &(mplsTunnelEntry.EntityData)
}

// MPLSTESTDMIB_MplsTunnelTable_MplsTunnelEntry_MplsTunnelAdminStatus represents tunnel.
type MPLSTESTDMIB_MplsTunnelTable_MplsTunnelEntry_MplsTunnelAdminStatus string

const (
    MPLSTESTDMIB_MplsTunnelTable_MplsTunnelEntry_MplsTunnelAdminStatus_up MPLSTESTDMIB_MplsTunnelTable_MplsTunnelEntry_MplsTunnelAdminStatus = "up"

    MPLSTESTDMIB_MplsTunnelTable_MplsTunnelEntry_MplsTunnelAdminStatus_down MPLSTESTDMIB_MplsTunnelTable_MplsTunnelEntry_MplsTunnelAdminStatus = "down"

    MPLSTESTDMIB_MplsTunnelTable_MplsTunnelEntry_MplsTunnelAdminStatus_testing MPLSTESTDMIB_MplsTunnelTable_MplsTunnelEntry_MplsTunnelAdminStatus = "testing"
)

// MPLSTESTDMIB_MplsTunnelTable_MplsTunnelEntry_MplsTunnelOperStatus represents this tunnel.
type MPLSTESTDMIB_MplsTunnelTable_MplsTunnelEntry_MplsTunnelOperStatus string

const (
    MPLSTESTDMIB_MplsTunnelTable_MplsTunnelEntry_MplsTunnelOperStatus_up MPLSTESTDMIB_MplsTunnelTable_MplsTunnelEntry_MplsTunnelOperStatus = "up"

    MPLSTESTDMIB_MplsTunnelTable_MplsTunnelEntry_MplsTunnelOperStatus_down MPLSTESTDMIB_MplsTunnelTable_MplsTunnelEntry_MplsTunnelOperStatus = "down"

    MPLSTESTDMIB_MplsTunnelTable_MplsTunnelEntry_MplsTunnelOperStatus_testing MPLSTESTDMIB_MplsTunnelTable_MplsTunnelEntry_MplsTunnelOperStatus = "testing"

    MPLSTESTDMIB_MplsTunnelTable_MplsTunnelEntry_MplsTunnelOperStatus_unknown MPLSTESTDMIB_MplsTunnelTable_MplsTunnelEntry_MplsTunnelOperStatus = "unknown"

    MPLSTESTDMIB_MplsTunnelTable_MplsTunnelEntry_MplsTunnelOperStatus_dormant MPLSTESTDMIB_MplsTunnelTable_MplsTunnelEntry_MplsTunnelOperStatus = "dormant"

    MPLSTESTDMIB_MplsTunnelTable_MplsTunnelEntry_MplsTunnelOperStatus_notPresent MPLSTESTDMIB_MplsTunnelTable_MplsTunnelEntry_MplsTunnelOperStatus = "notPresent"

    MPLSTESTDMIB_MplsTunnelTable_MplsTunnelEntry_MplsTunnelOperStatus_lowerLayerDown MPLSTESTDMIB_MplsTunnelTable_MplsTunnelEntry_MplsTunnelOperStatus = "lowerLayerDown"
)

// MPLSTESTDMIB_MplsTunnelTable_MplsTunnelEntry_MplsTunnelRole represents begin and end on the same LSR.
type MPLSTESTDMIB_MplsTunnelTable_MplsTunnelEntry_MplsTunnelRole string

const (
    MPLSTESTDMIB_MplsTunnelTable_MplsTunnelEntry_MplsTunnelRole_head MPLSTESTDMIB_MplsTunnelTable_MplsTunnelEntry_MplsTunnelRole = "head"

    MPLSTESTDMIB_MplsTunnelTable_MplsTunnelEntry_MplsTunnelRole_transit MPLSTESTDMIB_MplsTunnelTable_MplsTunnelEntry_MplsTunnelRole = "transit"

    MPLSTESTDMIB_MplsTunnelTable_MplsTunnelEntry_MplsTunnelRole_tail MPLSTESTDMIB_MplsTunnelTable_MplsTunnelEntry_MplsTunnelRole = "tail"

    MPLSTESTDMIB_MplsTunnelTable_MplsTunnelEntry_MplsTunnelRole_headTail MPLSTESTDMIB_MplsTunnelTable_MplsTunnelEntry_MplsTunnelRole = "headTail"
)

// MPLSTESTDMIB_MplsTunnelTable_MplsTunnelEntry_MplsTunnelSignallingProto represents tunnel.
type MPLSTESTDMIB_MplsTunnelTable_MplsTunnelEntry_MplsTunnelSignallingProto string

const (
    MPLSTESTDMIB_MplsTunnelTable_MplsTunnelEntry_MplsTunnelSignallingProto_none MPLSTESTDMIB_MplsTunnelTable_MplsTunnelEntry_MplsTunnelSignallingProto = "none"

    MPLSTESTDMIB_MplsTunnelTable_MplsTunnelEntry_MplsTunnelSignallingProto_rsvp MPLSTESTDMIB_MplsTunnelTable_MplsTunnelEntry_MplsTunnelSignallingProto = "rsvp"

    MPLSTESTDMIB_MplsTunnelTable_MplsTunnelEntry_MplsTunnelSignallingProto_crldp MPLSTESTDMIB_MplsTunnelTable_MplsTunnelEntry_MplsTunnelSignallingProto = "crldp"

    MPLSTESTDMIB_MplsTunnelTable_MplsTunnelEntry_MplsTunnelSignallingProto_other MPLSTESTDMIB_MplsTunnelTable_MplsTunnelEntry_MplsTunnelSignallingProto = "other"
)

// MPLSTESTDMIB_MplsTunnelHopTable
// The mplsTunnelHopTable is used to indicate the hops,
// strict or loose, for an instance of an MPLS tunnel
// defined in mplsTunnelTable, when it is established
// via signalling, for the outgoing direction of the
// tunnel. Thus at a transit LSR, this table contains
// the desired path of the tunnel from this LSR
// onwards. Each row in this table is indexed by
// mplsTunnelHopListIndex which corresponds to a group
// of hop lists or path options.  Each row also has a
// secondary index mplsTunnelHopIndex, which indicates
// a group of hops (also known as a path option).
// Finally, the third index, mplsTunnelHopIndex
// indicates the specific hop information for a path
// option. In case we want to specify a particular
// interface on the originating LSR of an outgoing
// tunnel by which we want packets to exit the LSR,
// we specify this as the first hop for this tunnel in
// mplsTunnelHopTable.
type MPLSTESTDMIB_MplsTunnelHopTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table represents a tunnel hop.  An entry is created by a
    // network administrator for signaled ERLSP set up by an MPLS signalling
    // protocol. The type is slice of
    // MPLSTESTDMIB_MplsTunnelHopTable_MplsTunnelHopEntry.
    MplsTunnelHopEntry []*MPLSTESTDMIB_MplsTunnelHopTable_MplsTunnelHopEntry
}

func (mplsTunnelHopTable *MPLSTESTDMIB_MplsTunnelHopTable) GetEntityData() *types.CommonEntityData {
    mplsTunnelHopTable.EntityData.YFilter = mplsTunnelHopTable.YFilter
    mplsTunnelHopTable.EntityData.YangName = "mplsTunnelHopTable"
    mplsTunnelHopTable.EntityData.BundleName = "cisco_ios_xe"
    mplsTunnelHopTable.EntityData.ParentYangName = "MPLS-TE-STD-MIB"
    mplsTunnelHopTable.EntityData.SegmentPath = "mplsTunnelHopTable"
    mplsTunnelHopTable.EntityData.AbsolutePath = "MPLS-TE-STD-MIB:MPLS-TE-STD-MIB/" + mplsTunnelHopTable.EntityData.SegmentPath
    mplsTunnelHopTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsTunnelHopTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsTunnelHopTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsTunnelHopTable.EntityData.Children = types.NewOrderedMap()
    mplsTunnelHopTable.EntityData.Children.Append("mplsTunnelHopEntry", types.YChild{"MplsTunnelHopEntry", nil})
    for i := range mplsTunnelHopTable.MplsTunnelHopEntry {
        mplsTunnelHopTable.EntityData.Children.Append(types.GetSegmentPath(mplsTunnelHopTable.MplsTunnelHopEntry[i]), types.YChild{"MplsTunnelHopEntry", mplsTunnelHopTable.MplsTunnelHopEntry[i]})
    }
    mplsTunnelHopTable.EntityData.Leafs = types.NewOrderedMap()

    mplsTunnelHopTable.EntityData.YListKeys = []string {}

    return &(mplsTunnelHopTable.EntityData)
}

// MPLSTESTDMIB_MplsTunnelHopTable_MplsTunnelHopEntry
// An entry in this table represents a tunnel hop.  An
// entry is created by a network administrator for
// signaled ERLSP set up by an MPLS signalling
// protocol.
type MPLSTESTDMIB_MplsTunnelHopTable_MplsTunnelHopEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Primary index into this table identifying a
    // particular explicit route object. The type is interface{} with range:
    // 1..4294967295.
    MplsTunnelHopListIndex interface{}

    // This attribute is a key. Secondary index into this table identifying a
    // particular group of hops representing a particular configured path. This is
    // otherwise known as a path option. The type is interface{} with range:
    // 1..4294967295.
    MplsTunnelHopPathOptionIndex interface{}

    // This attribute is a key. Tertiary index into this table identifying a
    // particular hop. The type is interface{} with range: 1..4294967295.
    MplsTunnelHopIndex interface{}

    // The Hop Address Type of this tunnel hop.  The value of this object cannot
    // be changed if the value of the corresponding mplsTunnelHopRowStatus object
    // is 'active'.  Note that lspid(5) is a valid option only for tunnels
    // signaled via CRLDP. The type is TeHopAddressType.
    MplsTunnelHopAddrType interface{}

    // The Tunnel Hop Address for this tunnel hop.  The type of this address is
    // determined by the value of the corresponding mplsTunnelHopAddrType.  The
    // value of this object cannot be changed if the value of the corresponding
    // mplsTunnelHopRowStatus object is 'active'. The type is string with length:
    // 0..32.
    MplsTunnelHopIpAddr interface{}

    // If mplsTunnelHopAddrType is set to ipv4(1) or ipv6(2), then this value will
    // contain an appropriate prefix length for the IP address in object
    // mplsTunnelHopIpAddr. Otherwise this value is irrelevant and should be
    // ignored. The type is interface{} with range: 0..2040.
    MplsTunnelHopIpPrefixLen interface{}

    // If mplsTunnelHopAddrType is set to asnumber(3), then this value will
    // contain the AS number of this hop. Otherwise the agent should set this
    // object to zero- length string and the manager should ignore this. The type
    // is string with length: 4..4.
    MplsTunnelHopAsNumber interface{}

    // If mplsTunnelHopAddrType is set to unnum(4), then this value will contain
    // the interface identifier of the unnumbered interface for this hop. This
    // object should be used in conjunction with mplsTunnelHopIpAddress which
    // would contain the LSR Router ID in this case. Otherwise the agent should
    // set this object to zero-length string and the manager should ignore this.
    // The type is string with length: 4..4.
    MplsTunnelHopAddrUnnum interface{}

    // If mplsTunnelHopAddrType is set to lspid(5), then this value will contain
    // the LSPID of a tunnel of this hop. The present tunnel being configured is
    // tunneled through this hop (using label stacking). This object is otherwise
    // insignificant and should contain a value of 0 to indicate this fact. The
    // type is string with length: 2..2 | 6..6.
    MplsTunnelHopLspId interface{}

    // Denotes whether this tunnel hop is routed in a strict or loose fashion. The
    // value of this object has no meaning if the mplsTunnelHopInclude object is
    // set to 'false'. The type is MplsTunnelHopType.
    MplsTunnelHopType interface{}

    // If this value is set to true, then this indicates that this hop must be
    // included in the tunnel's path. If this value is set to 'false', then this
    // hop must be avoided when calculating the path for this tunnel. The default
    // value of this object is 'true', so that by default all indicated hops are
    // included in the CSPF path computation. If this object is set to 'false' the
    // value of mplsTunnelHopType should be ignored. The type is bool.
    MplsTunnelHopInclude interface{}

    // The description of this series of hops as they relate to the specified path
    // option. The value of this object SHOULD be the same for each hop in the
    // series that comprises a path option. The type is string.
    MplsTunnelHopPathOptionName interface{}

    // If this value is set to dynamic, then the user should only specify the
    // source and destination of the path and expect that the CSPF will calculate
    // the remainder of the path.  If this value is set to explicit, the user
    // should specify the entire path for the tunnel to take.  This path may
    // contain strict or loose hops.  Each hop along a specific path SHOULD have
    // this object set to the same value. The type is MplsTunnelHopEntryPathComp.
    MplsTunnelHopEntryPathComp interface{}

    // This variable is used to create, modify, and/or delete a row in this table.
    // When a row in this table is in active(1) state, no objects in that row can
    // be modified by the agent except mplsTunnelHopRowStatus and
    // mplsTunnelHopStorageType. The type is RowStatus.
    MplsTunnelHopRowStatus interface{}

    // The storage type for this Hop entry. Conceptual rows having the value
    // 'permanent' need not allow write-access to any columnar objects in the row.
    // The type is StorageType.
    MplsTunnelHopStorageType interface{}
}

func (mplsTunnelHopEntry *MPLSTESTDMIB_MplsTunnelHopTable_MplsTunnelHopEntry) GetEntityData() *types.CommonEntityData {
    mplsTunnelHopEntry.EntityData.YFilter = mplsTunnelHopEntry.YFilter
    mplsTunnelHopEntry.EntityData.YangName = "mplsTunnelHopEntry"
    mplsTunnelHopEntry.EntityData.BundleName = "cisco_ios_xe"
    mplsTunnelHopEntry.EntityData.ParentYangName = "mplsTunnelHopTable"
    mplsTunnelHopEntry.EntityData.SegmentPath = "mplsTunnelHopEntry" + types.AddKeyToken(mplsTunnelHopEntry.MplsTunnelHopListIndex, "mplsTunnelHopListIndex") + types.AddKeyToken(mplsTunnelHopEntry.MplsTunnelHopPathOptionIndex, "mplsTunnelHopPathOptionIndex") + types.AddKeyToken(mplsTunnelHopEntry.MplsTunnelHopIndex, "mplsTunnelHopIndex")
    mplsTunnelHopEntry.EntityData.AbsolutePath = "MPLS-TE-STD-MIB:MPLS-TE-STD-MIB/mplsTunnelHopTable/" + mplsTunnelHopEntry.EntityData.SegmentPath
    mplsTunnelHopEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsTunnelHopEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsTunnelHopEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsTunnelHopEntry.EntityData.Children = types.NewOrderedMap()
    mplsTunnelHopEntry.EntityData.Leafs = types.NewOrderedMap()
    mplsTunnelHopEntry.EntityData.Leafs.Append("mplsTunnelHopListIndex", types.YLeaf{"MplsTunnelHopListIndex", mplsTunnelHopEntry.MplsTunnelHopListIndex})
    mplsTunnelHopEntry.EntityData.Leafs.Append("mplsTunnelHopPathOptionIndex", types.YLeaf{"MplsTunnelHopPathOptionIndex", mplsTunnelHopEntry.MplsTunnelHopPathOptionIndex})
    mplsTunnelHopEntry.EntityData.Leafs.Append("mplsTunnelHopIndex", types.YLeaf{"MplsTunnelHopIndex", mplsTunnelHopEntry.MplsTunnelHopIndex})
    mplsTunnelHopEntry.EntityData.Leafs.Append("mplsTunnelHopAddrType", types.YLeaf{"MplsTunnelHopAddrType", mplsTunnelHopEntry.MplsTunnelHopAddrType})
    mplsTunnelHopEntry.EntityData.Leafs.Append("mplsTunnelHopIpAddr", types.YLeaf{"MplsTunnelHopIpAddr", mplsTunnelHopEntry.MplsTunnelHopIpAddr})
    mplsTunnelHopEntry.EntityData.Leafs.Append("mplsTunnelHopIpPrefixLen", types.YLeaf{"MplsTunnelHopIpPrefixLen", mplsTunnelHopEntry.MplsTunnelHopIpPrefixLen})
    mplsTunnelHopEntry.EntityData.Leafs.Append("mplsTunnelHopAsNumber", types.YLeaf{"MplsTunnelHopAsNumber", mplsTunnelHopEntry.MplsTunnelHopAsNumber})
    mplsTunnelHopEntry.EntityData.Leafs.Append("mplsTunnelHopAddrUnnum", types.YLeaf{"MplsTunnelHopAddrUnnum", mplsTunnelHopEntry.MplsTunnelHopAddrUnnum})
    mplsTunnelHopEntry.EntityData.Leafs.Append("mplsTunnelHopLspId", types.YLeaf{"MplsTunnelHopLspId", mplsTunnelHopEntry.MplsTunnelHopLspId})
    mplsTunnelHopEntry.EntityData.Leafs.Append("mplsTunnelHopType", types.YLeaf{"MplsTunnelHopType", mplsTunnelHopEntry.MplsTunnelHopType})
    mplsTunnelHopEntry.EntityData.Leafs.Append("mplsTunnelHopInclude", types.YLeaf{"MplsTunnelHopInclude", mplsTunnelHopEntry.MplsTunnelHopInclude})
    mplsTunnelHopEntry.EntityData.Leafs.Append("mplsTunnelHopPathOptionName", types.YLeaf{"MplsTunnelHopPathOptionName", mplsTunnelHopEntry.MplsTunnelHopPathOptionName})
    mplsTunnelHopEntry.EntityData.Leafs.Append("mplsTunnelHopEntryPathComp", types.YLeaf{"MplsTunnelHopEntryPathComp", mplsTunnelHopEntry.MplsTunnelHopEntryPathComp})
    mplsTunnelHopEntry.EntityData.Leafs.Append("mplsTunnelHopRowStatus", types.YLeaf{"MplsTunnelHopRowStatus", mplsTunnelHopEntry.MplsTunnelHopRowStatus})
    mplsTunnelHopEntry.EntityData.Leafs.Append("mplsTunnelHopStorageType", types.YLeaf{"MplsTunnelHopStorageType", mplsTunnelHopEntry.MplsTunnelHopStorageType})

    mplsTunnelHopEntry.EntityData.YListKeys = []string {"MplsTunnelHopListIndex", "MplsTunnelHopPathOptionIndex", "MplsTunnelHopIndex"}

    return &(mplsTunnelHopEntry.EntityData)
}

// MPLSTESTDMIB_MplsTunnelHopTable_MplsTunnelHopEntry_MplsTunnelHopEntryPathComp represents path SHOULD have this object set to the same value
type MPLSTESTDMIB_MplsTunnelHopTable_MplsTunnelHopEntry_MplsTunnelHopEntryPathComp string

const (
    MPLSTESTDMIB_MplsTunnelHopTable_MplsTunnelHopEntry_MplsTunnelHopEntryPathComp_dynamic MPLSTESTDMIB_MplsTunnelHopTable_MplsTunnelHopEntry_MplsTunnelHopEntryPathComp = "dynamic"

    MPLSTESTDMIB_MplsTunnelHopTable_MplsTunnelHopEntry_MplsTunnelHopEntryPathComp_explicit MPLSTESTDMIB_MplsTunnelHopTable_MplsTunnelHopEntry_MplsTunnelHopEntryPathComp = "explicit"
)

// MPLSTESTDMIB_MplsTunnelHopTable_MplsTunnelHopEntry_MplsTunnelHopType represents is set to 'false'.
type MPLSTESTDMIB_MplsTunnelHopTable_MplsTunnelHopEntry_MplsTunnelHopType string

const (
    MPLSTESTDMIB_MplsTunnelHopTable_MplsTunnelHopEntry_MplsTunnelHopType_strict MPLSTESTDMIB_MplsTunnelHopTable_MplsTunnelHopEntry_MplsTunnelHopType = "strict"

    MPLSTESTDMIB_MplsTunnelHopTable_MplsTunnelHopEntry_MplsTunnelHopType_loose MPLSTESTDMIB_MplsTunnelHopTable_MplsTunnelHopEntry_MplsTunnelHopType = "loose"
)

// MPLSTESTDMIB_MplsTunnelResourceTable
// The mplsTunnelResourceTable allows a manager to
// specify which resources are desired for an MPLS
// tunnel.  This table also allows several tunnels to
// point to a single entry in this table, implying
// that these tunnels should share resources.
type MPLSTESTDMIB_MplsTunnelResourceTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table represents a set of resources for an MPLS tunnel. 
    // An entry can be created by a network administrator or by an SNMP agent as
    // instructed by any MPLS signalling protocol. An entry in this table
    // referenced by a tunnel instance with zero mplsTunnelInstance value
    // indicates a configured set of resource parameter. An entry referenced by a
    // tunnel instance with a non-zero mplsTunnelInstance reflects the in-use
    // resource parameters for the tunnel instance which may have been negotiated
    // or modified by the MPLS signaling protocols. The type is slice of
    // MPLSTESTDMIB_MplsTunnelResourceTable_MplsTunnelResourceEntry.
    MplsTunnelResourceEntry []*MPLSTESTDMIB_MplsTunnelResourceTable_MplsTunnelResourceEntry
}

func (mplsTunnelResourceTable *MPLSTESTDMIB_MplsTunnelResourceTable) GetEntityData() *types.CommonEntityData {
    mplsTunnelResourceTable.EntityData.YFilter = mplsTunnelResourceTable.YFilter
    mplsTunnelResourceTable.EntityData.YangName = "mplsTunnelResourceTable"
    mplsTunnelResourceTable.EntityData.BundleName = "cisco_ios_xe"
    mplsTunnelResourceTable.EntityData.ParentYangName = "MPLS-TE-STD-MIB"
    mplsTunnelResourceTable.EntityData.SegmentPath = "mplsTunnelResourceTable"
    mplsTunnelResourceTable.EntityData.AbsolutePath = "MPLS-TE-STD-MIB:MPLS-TE-STD-MIB/" + mplsTunnelResourceTable.EntityData.SegmentPath
    mplsTunnelResourceTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsTunnelResourceTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsTunnelResourceTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsTunnelResourceTable.EntityData.Children = types.NewOrderedMap()
    mplsTunnelResourceTable.EntityData.Children.Append("mplsTunnelResourceEntry", types.YChild{"MplsTunnelResourceEntry", nil})
    for i := range mplsTunnelResourceTable.MplsTunnelResourceEntry {
        mplsTunnelResourceTable.EntityData.Children.Append(types.GetSegmentPath(mplsTunnelResourceTable.MplsTunnelResourceEntry[i]), types.YChild{"MplsTunnelResourceEntry", mplsTunnelResourceTable.MplsTunnelResourceEntry[i]})
    }
    mplsTunnelResourceTable.EntityData.Leafs = types.NewOrderedMap()

    mplsTunnelResourceTable.EntityData.YListKeys = []string {}

    return &(mplsTunnelResourceTable.EntityData)
}

// MPLSTESTDMIB_MplsTunnelResourceTable_MplsTunnelResourceEntry
// An entry in this table represents a set of resources
// for an MPLS tunnel.  An entry can be created by a
// network administrator or by an SNMP agent as
// instructed by any MPLS signalling protocol.
// An entry in this table referenced by a tunnel instance
// with zero mplsTunnelInstance value indicates a
// configured set of resource parameter. An entry
// referenced by a tunnel instance with a non-zero
// mplsTunnelInstance reflects the in-use resource
// parameters for the tunnel instance which may have
// been negotiated or modified by the MPLS signaling
// protocols.
type MPLSTESTDMIB_MplsTunnelResourceTable_MplsTunnelResourceEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Uniquely identifies this row. The type is
    // interface{} with range: 1..2147483647.
    MplsTunnelResourceIndex interface{}

    // The maximum rate in bits/second.  Note that setting
    // mplsTunnelResourceMaxRate, mplsTunnelResourceMeanRate, and
    // mplsTunnelResourceMaxBurstSize to 0 indicates best- effort treatment. The
    // type is interface{} with range: 0..4294967295. Units are kilobits per
    // second.
    MplsTunnelResourceMaxRate interface{}

    // This object is copied into an instance of mplsTrafficParamMeanRate in the
    // mplsTrafficParamTable. The OID of this table entry is then copied into the
    // corresponding mplsInSegmentTrafficParamPtr. The type is interface{} with
    // range: 0..4294967295. Units are kilobits per second.
    MplsTunnelResourceMeanRate interface{}

    // The maximum burst size in bytes. The type is interface{} with range:
    // 0..4294967295. Units are bytes.
    MplsTunnelResourceMaxBurstSize interface{}

    // The mean burst size in bytes.  The implementations which do not implement
    // this variable must return a noSuchObject exception for this object and must
    // not allow a user to set this object. The type is interface{} with range:
    // 0..4294967295. Units are bytes.
    MplsTunnelResourceMeanBurstSize interface{}

    // The Excess burst size in bytes.  The implementations which do not implement
    // this variable must return noSuchObject exception for this object and must
    // not allow a user to set this value. The type is interface{} with range:
    // 0..4294967295. Units are bytes.
    MplsTunnelResourceExBurstSize interface{}

    // The granularity of the availability of committed rate.  The implementations
    // which do not implement this variable must return unspecified(1) for this
    // value and must not allow a user to set this value. The type is
    // MplsTunnelResourceFrequency.
    MplsTunnelResourceFrequency interface{}

    // The relative weight for using excess bandwidth above its committed rate. 
    // The value of 0 means that weight is not applicable for the CR-LSP. The type
    // is interface{} with range: 0..255.
    MplsTunnelResourceWeight interface{}

    // This variable is used to create, modify, and/or delete a row in this table.
    // When a row in this table is in active(1) state, no objects in that row can
    // be modified by the agent except mplsTunnelResourceRowStatus and
    // mplsTunnelResourceStorageType. The type is RowStatus.
    MplsTunnelResourceRowStatus interface{}

    // The storage type for this Hop entry. Conceptual rows having the value
    // 'permanent' need not allow write-access to any columnar objects in the row.
    // The type is StorageType.
    MplsTunnelResourceStorageType interface{}
}

func (mplsTunnelResourceEntry *MPLSTESTDMIB_MplsTunnelResourceTable_MplsTunnelResourceEntry) GetEntityData() *types.CommonEntityData {
    mplsTunnelResourceEntry.EntityData.YFilter = mplsTunnelResourceEntry.YFilter
    mplsTunnelResourceEntry.EntityData.YangName = "mplsTunnelResourceEntry"
    mplsTunnelResourceEntry.EntityData.BundleName = "cisco_ios_xe"
    mplsTunnelResourceEntry.EntityData.ParentYangName = "mplsTunnelResourceTable"
    mplsTunnelResourceEntry.EntityData.SegmentPath = "mplsTunnelResourceEntry" + types.AddKeyToken(mplsTunnelResourceEntry.MplsTunnelResourceIndex, "mplsTunnelResourceIndex")
    mplsTunnelResourceEntry.EntityData.AbsolutePath = "MPLS-TE-STD-MIB:MPLS-TE-STD-MIB/mplsTunnelResourceTable/" + mplsTunnelResourceEntry.EntityData.SegmentPath
    mplsTunnelResourceEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsTunnelResourceEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsTunnelResourceEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsTunnelResourceEntry.EntityData.Children = types.NewOrderedMap()
    mplsTunnelResourceEntry.EntityData.Leafs = types.NewOrderedMap()
    mplsTunnelResourceEntry.EntityData.Leafs.Append("mplsTunnelResourceIndex", types.YLeaf{"MplsTunnelResourceIndex", mplsTunnelResourceEntry.MplsTunnelResourceIndex})
    mplsTunnelResourceEntry.EntityData.Leafs.Append("mplsTunnelResourceMaxRate", types.YLeaf{"MplsTunnelResourceMaxRate", mplsTunnelResourceEntry.MplsTunnelResourceMaxRate})
    mplsTunnelResourceEntry.EntityData.Leafs.Append("mplsTunnelResourceMeanRate", types.YLeaf{"MplsTunnelResourceMeanRate", mplsTunnelResourceEntry.MplsTunnelResourceMeanRate})
    mplsTunnelResourceEntry.EntityData.Leafs.Append("mplsTunnelResourceMaxBurstSize", types.YLeaf{"MplsTunnelResourceMaxBurstSize", mplsTunnelResourceEntry.MplsTunnelResourceMaxBurstSize})
    mplsTunnelResourceEntry.EntityData.Leafs.Append("mplsTunnelResourceMeanBurstSize", types.YLeaf{"MplsTunnelResourceMeanBurstSize", mplsTunnelResourceEntry.MplsTunnelResourceMeanBurstSize})
    mplsTunnelResourceEntry.EntityData.Leafs.Append("mplsTunnelResourceExBurstSize", types.YLeaf{"MplsTunnelResourceExBurstSize", mplsTunnelResourceEntry.MplsTunnelResourceExBurstSize})
    mplsTunnelResourceEntry.EntityData.Leafs.Append("mplsTunnelResourceFrequency", types.YLeaf{"MplsTunnelResourceFrequency", mplsTunnelResourceEntry.MplsTunnelResourceFrequency})
    mplsTunnelResourceEntry.EntityData.Leafs.Append("mplsTunnelResourceWeight", types.YLeaf{"MplsTunnelResourceWeight", mplsTunnelResourceEntry.MplsTunnelResourceWeight})
    mplsTunnelResourceEntry.EntityData.Leafs.Append("mplsTunnelResourceRowStatus", types.YLeaf{"MplsTunnelResourceRowStatus", mplsTunnelResourceEntry.MplsTunnelResourceRowStatus})
    mplsTunnelResourceEntry.EntityData.Leafs.Append("mplsTunnelResourceStorageType", types.YLeaf{"MplsTunnelResourceStorageType", mplsTunnelResourceEntry.MplsTunnelResourceStorageType})

    mplsTunnelResourceEntry.EntityData.YListKeys = []string {"MplsTunnelResourceIndex"}

    return &(mplsTunnelResourceEntry.EntityData)
}

// MPLSTESTDMIB_MplsTunnelResourceTable_MplsTunnelResourceEntry_MplsTunnelResourceFrequency represents value and must not allow a user to set this value.
type MPLSTESTDMIB_MplsTunnelResourceTable_MplsTunnelResourceEntry_MplsTunnelResourceFrequency string

const (
    MPLSTESTDMIB_MplsTunnelResourceTable_MplsTunnelResourceEntry_MplsTunnelResourceFrequency_unspecified MPLSTESTDMIB_MplsTunnelResourceTable_MplsTunnelResourceEntry_MplsTunnelResourceFrequency = "unspecified"

    MPLSTESTDMIB_MplsTunnelResourceTable_MplsTunnelResourceEntry_MplsTunnelResourceFrequency_frequent MPLSTESTDMIB_MplsTunnelResourceTable_MplsTunnelResourceEntry_MplsTunnelResourceFrequency = "frequent"

    MPLSTESTDMIB_MplsTunnelResourceTable_MplsTunnelResourceEntry_MplsTunnelResourceFrequency_veryFrequent MPLSTESTDMIB_MplsTunnelResourceTable_MplsTunnelResourceEntry_MplsTunnelResourceFrequency = "veryFrequent"
)

// MPLSTESTDMIB_MplsTunnelARHopTable
// The mplsTunnelARHopTable is used to indicate the
// hops for an MPLS tunnel defined in mplsTunnelTable,
// as reported by the MPLS signalling protocol. Thus at
// a transit LSR, this table (if the table is supported
// and if the signaling protocol is recording actual
// route information) contains the actual route of the
// whole tunnel. If the signaling protocol is not
// recording the actual route, this table MAY report
// the information from the mplsTunnelHopTable or the
// mplsTunnelCHopTable.
// 
// Each row in this table is indexed by
// mplsTunnelARHopListIndex. Each row also has a
// secondary index mplsTunnelARHopIndex, corresponding
// to the next hop that this row corresponds to.
// 
// Please note that since the information necessary to
// build entries within this table is not provided by
// some MPLS signalling protocols, implementation of
// this table is optional. Furthermore, since the
// information in this table is actually provided by
// the MPLS signalling protocol after the path has
// been set-up, the entries in this table are provided
// only for observation, and hence, all variables in
// this table are accessible exclusively as read-
// only.
// 
// Note also that the contents of this table may change
// while it is being read because of re-routing
// activities. A network administrator may verify that
// the actual route read is consistent by reference to
// the mplsTunnelLastPathChange object.
type MPLSTESTDMIB_MplsTunnelARHopTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table represents a tunnel hop.  An entry is created by the
    // agent for signaled ERLSP set up by an MPLS signalling protocol. The type is
    // slice of MPLSTESTDMIB_MplsTunnelARHopTable_MplsTunnelARHopEntry.
    MplsTunnelARHopEntry []*MPLSTESTDMIB_MplsTunnelARHopTable_MplsTunnelARHopEntry
}

func (mplsTunnelARHopTable *MPLSTESTDMIB_MplsTunnelARHopTable) GetEntityData() *types.CommonEntityData {
    mplsTunnelARHopTable.EntityData.YFilter = mplsTunnelARHopTable.YFilter
    mplsTunnelARHopTable.EntityData.YangName = "mplsTunnelARHopTable"
    mplsTunnelARHopTable.EntityData.BundleName = "cisco_ios_xe"
    mplsTunnelARHopTable.EntityData.ParentYangName = "MPLS-TE-STD-MIB"
    mplsTunnelARHopTable.EntityData.SegmentPath = "mplsTunnelARHopTable"
    mplsTunnelARHopTable.EntityData.AbsolutePath = "MPLS-TE-STD-MIB:MPLS-TE-STD-MIB/" + mplsTunnelARHopTable.EntityData.SegmentPath
    mplsTunnelARHopTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsTunnelARHopTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsTunnelARHopTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsTunnelARHopTable.EntityData.Children = types.NewOrderedMap()
    mplsTunnelARHopTable.EntityData.Children.Append("mplsTunnelARHopEntry", types.YChild{"MplsTunnelARHopEntry", nil})
    for i := range mplsTunnelARHopTable.MplsTunnelARHopEntry {
        mplsTunnelARHopTable.EntityData.Children.Append(types.GetSegmentPath(mplsTunnelARHopTable.MplsTunnelARHopEntry[i]), types.YChild{"MplsTunnelARHopEntry", mplsTunnelARHopTable.MplsTunnelARHopEntry[i]})
    }
    mplsTunnelARHopTable.EntityData.Leafs = types.NewOrderedMap()

    mplsTunnelARHopTable.EntityData.YListKeys = []string {}

    return &(mplsTunnelARHopTable.EntityData)
}

// MPLSTESTDMIB_MplsTunnelARHopTable_MplsTunnelARHopEntry
// An entry in this table represents a tunnel hop.  An
// entry is created by the agent for signaled ERLSP
// set up by an MPLS signalling protocol.
type MPLSTESTDMIB_MplsTunnelARHopTable_MplsTunnelARHopEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Primary index into this table identifying a
    // particular recorded hop list. The type is interface{} with range:
    // 1..4294967295.
    MplsTunnelARHopListIndex interface{}

    // This attribute is a key. Secondary index into this table identifying the
    // particular hop. The type is interface{} with range: 1..4294967295.
    MplsTunnelARHopIndex interface{}

    // The Hop Address Type of this tunnel hop.  Note that lspid(5) is a valid
    // option only for tunnels signaled via CRLDP. The type is TeHopAddressType.
    MplsTunnelARHopAddrType interface{}

    // The Tunnel Hop Address for this tunnel hop.  The type of this address is
    // determined by the value of the corresponding mplsTunnelARHopAddrType. If
    // mplsTunnelARHopAddrType is set to unnum(4),  then this value contains the
    // LSR Router ID of the  unnumbered interface. Otherwise the agent SHOULD  set
    // this object to the zero-length string and the  manager should ignore this
    // object. The type is string with length: 0..32.
    MplsTunnelARHopIpAddr interface{}

    // If mplsTunnelARHopAddrType is set to unnum(4), then this value will contain
    // the interface identifier of the unnumbered interface for this hop. This
    // object should be used in conjunction with mplsTunnelARHopIpAddr which would
    // contain the LSR Router ID in this case. Otherwise the agent should set this
    // object to zero-length string and the manager should ignore this. The type
    // is string with length: 4..4.
    MplsTunnelARHopAddrUnnum interface{}

    // If mplsTunnelARHopAddrType is set to lspid(5), then this value will contain
    // the LSP ID of this hop. This object is otherwise insignificant and should
    // contain a value of 0 to indicate this fact. The type is string with length:
    // 2..2 | 6..6.
    MplsTunnelARHopLspId interface{}
}

func (mplsTunnelARHopEntry *MPLSTESTDMIB_MplsTunnelARHopTable_MplsTunnelARHopEntry) GetEntityData() *types.CommonEntityData {
    mplsTunnelARHopEntry.EntityData.YFilter = mplsTunnelARHopEntry.YFilter
    mplsTunnelARHopEntry.EntityData.YangName = "mplsTunnelARHopEntry"
    mplsTunnelARHopEntry.EntityData.BundleName = "cisco_ios_xe"
    mplsTunnelARHopEntry.EntityData.ParentYangName = "mplsTunnelARHopTable"
    mplsTunnelARHopEntry.EntityData.SegmentPath = "mplsTunnelARHopEntry" + types.AddKeyToken(mplsTunnelARHopEntry.MplsTunnelARHopListIndex, "mplsTunnelARHopListIndex") + types.AddKeyToken(mplsTunnelARHopEntry.MplsTunnelARHopIndex, "mplsTunnelARHopIndex")
    mplsTunnelARHopEntry.EntityData.AbsolutePath = "MPLS-TE-STD-MIB:MPLS-TE-STD-MIB/mplsTunnelARHopTable/" + mplsTunnelARHopEntry.EntityData.SegmentPath
    mplsTunnelARHopEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsTunnelARHopEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsTunnelARHopEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsTunnelARHopEntry.EntityData.Children = types.NewOrderedMap()
    mplsTunnelARHopEntry.EntityData.Leafs = types.NewOrderedMap()
    mplsTunnelARHopEntry.EntityData.Leafs.Append("mplsTunnelARHopListIndex", types.YLeaf{"MplsTunnelARHopListIndex", mplsTunnelARHopEntry.MplsTunnelARHopListIndex})
    mplsTunnelARHopEntry.EntityData.Leafs.Append("mplsTunnelARHopIndex", types.YLeaf{"MplsTunnelARHopIndex", mplsTunnelARHopEntry.MplsTunnelARHopIndex})
    mplsTunnelARHopEntry.EntityData.Leafs.Append("mplsTunnelARHopAddrType", types.YLeaf{"MplsTunnelARHopAddrType", mplsTunnelARHopEntry.MplsTunnelARHopAddrType})
    mplsTunnelARHopEntry.EntityData.Leafs.Append("mplsTunnelARHopIpAddr", types.YLeaf{"MplsTunnelARHopIpAddr", mplsTunnelARHopEntry.MplsTunnelARHopIpAddr})
    mplsTunnelARHopEntry.EntityData.Leafs.Append("mplsTunnelARHopAddrUnnum", types.YLeaf{"MplsTunnelARHopAddrUnnum", mplsTunnelARHopEntry.MplsTunnelARHopAddrUnnum})
    mplsTunnelARHopEntry.EntityData.Leafs.Append("mplsTunnelARHopLspId", types.YLeaf{"MplsTunnelARHopLspId", mplsTunnelARHopEntry.MplsTunnelARHopLspId})

    mplsTunnelARHopEntry.EntityData.YListKeys = []string {"MplsTunnelARHopListIndex", "MplsTunnelARHopIndex"}

    return &(mplsTunnelARHopEntry.EntityData)
}

// MPLSTESTDMIB_MplsTunnelCHopTable
// The mplsTunnelCHopTable is used to indicate the
// hops, strict or loose, for an MPLS tunnel defined
// in mplsTunnelTable, as computed by a constraint-
// based routing protocol, based on the
// mplsTunnelHopTable for the outgoing direction of
// the tunnel. Thus at a transit LSR, this table (if
// the table is supported) MAY contain the path
// computed by the CSPF engine on (or on behalf of)
// this LSR. Each row in this table is indexed by
// mplsTunnelCHopListIndex.  Each row also has a
// secondary index mplsTunnelCHopIndex, corresponding
// to the next hop that this row corresponds to. In
// case we want to specify a particular interface on
// the originating LSR of an outgoing tunnel by which
// we want packets to exit the LSR, we specify this as
// the first hop for this tunnel in
// mplsTunnelCHopTable.
// 
// Please note that since the information necessary to
// build entries within this table may not be
// supported by some LSRs, implementation of this
// table is optional. Furthermore, since the
// information in this table describes the path
// computed by the CSPF engine the entries in this
// table are read-only.
type MPLSTESTDMIB_MplsTunnelCHopTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table represents a tunnel hop.  An entry in this table is
    // created by a path computation engine using CSPF techniques applied to the
    // information collected by routing protocols and the hops specified in the
    // corresponding mplsTunnelHopTable. The type is slice of
    // MPLSTESTDMIB_MplsTunnelCHopTable_MplsTunnelCHopEntry.
    MplsTunnelCHopEntry []*MPLSTESTDMIB_MplsTunnelCHopTable_MplsTunnelCHopEntry
}

func (mplsTunnelCHopTable *MPLSTESTDMIB_MplsTunnelCHopTable) GetEntityData() *types.CommonEntityData {
    mplsTunnelCHopTable.EntityData.YFilter = mplsTunnelCHopTable.YFilter
    mplsTunnelCHopTable.EntityData.YangName = "mplsTunnelCHopTable"
    mplsTunnelCHopTable.EntityData.BundleName = "cisco_ios_xe"
    mplsTunnelCHopTable.EntityData.ParentYangName = "MPLS-TE-STD-MIB"
    mplsTunnelCHopTable.EntityData.SegmentPath = "mplsTunnelCHopTable"
    mplsTunnelCHopTable.EntityData.AbsolutePath = "MPLS-TE-STD-MIB:MPLS-TE-STD-MIB/" + mplsTunnelCHopTable.EntityData.SegmentPath
    mplsTunnelCHopTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsTunnelCHopTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsTunnelCHopTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsTunnelCHopTable.EntityData.Children = types.NewOrderedMap()
    mplsTunnelCHopTable.EntityData.Children.Append("mplsTunnelCHopEntry", types.YChild{"MplsTunnelCHopEntry", nil})
    for i := range mplsTunnelCHopTable.MplsTunnelCHopEntry {
        mplsTunnelCHopTable.EntityData.Children.Append(types.GetSegmentPath(mplsTunnelCHopTable.MplsTunnelCHopEntry[i]), types.YChild{"MplsTunnelCHopEntry", mplsTunnelCHopTable.MplsTunnelCHopEntry[i]})
    }
    mplsTunnelCHopTable.EntityData.Leafs = types.NewOrderedMap()

    mplsTunnelCHopTable.EntityData.YListKeys = []string {}

    return &(mplsTunnelCHopTable.EntityData)
}

// MPLSTESTDMIB_MplsTunnelCHopTable_MplsTunnelCHopEntry
// An entry in this table represents a tunnel hop.  An
// entry in this table is created by a path
// computation engine using CSPF techniques applied to
// the information collected by routing protocols and
// the hops specified in the corresponding
// mplsTunnelHopTable.
type MPLSTESTDMIB_MplsTunnelCHopTable_MplsTunnelCHopEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Primary index into this table identifying a
    // particular computed hop list. The type is interface{} with range:
    // 1..4294967295.
    MplsTunnelCHopListIndex interface{}

    // This attribute is a key. Secondary index into this table identifying the
    // particular hop. The type is interface{} with range: 1..4294967295.
    MplsTunnelCHopIndex interface{}

    // The Hop Address Type of this tunnel hop.  Note that lspid(5) is a valid
    // option only for tunnels signaled via CRLDP. The type is TeHopAddressType.
    MplsTunnelCHopAddrType interface{}

    // The Tunnel Hop Address for this tunnel hop. The type of this address is
    // determined by the  value of the corresponding mplsTunnelCHopAddrType.  If
    // mplsTunnelCHopAddrType is set to unnum(4), then  this value will contain
    // the LSR Router ID of the  unnumbered interface. Otherwise the agent should 
    // set this object to the zero-length string and the  manager SHOULD ignore
    // this object. The type is string with length: 0..32.
    MplsTunnelCHopIpAddr interface{}

    // If mplsTunnelCHopAddrType is set to ipv4(1) or ipv6(2), then this value
    // will contain an appropriate prefix length for the IP address in object
    // mplsTunnelCHopIpAddr. Otherwise this value is irrelevant and should be
    // ignored. The type is interface{} with range: 0..2040.
    MplsTunnelCHopIpPrefixLen interface{}

    // If mplsTunnelCHopAddrType is set to asnumber(3), then this value will
    // contain the AS number of this hop. Otherwise the agent should set this
    // object to zero-length string and the manager should ignore this. The type
    // is string with length: 4..4.
    MplsTunnelCHopAsNumber interface{}

    // If mplsTunnelCHopAddrType is set to unnum(4), then this value will contain
    // the unnumbered interface identifier of this hop. This object should be used
    // in conjunction with mplsTunnelCHopIpAddr which would contain the LSR Router
    // ID in this case. Otherwise the agent should set this object to zero- length
    // string and the manager should ignore this. The type is string with length:
    // 4..4.
    MplsTunnelCHopAddrUnnum interface{}

    // If mplsTunnelCHopAddrType is set to lspid(5), then this value will contain
    // the LSP ID of this hop. This object is otherwise insignificant and should
    // contain a value of 0 to indicate this fact. The type is string with length:
    // 2..2 | 6..6.
    MplsTunnelCHopLspId interface{}

    // Denotes whether this is tunnel hop is routed in a strict or loose fashion.
    // The type is MplsTunnelCHopType.
    MplsTunnelCHopType interface{}
}

func (mplsTunnelCHopEntry *MPLSTESTDMIB_MplsTunnelCHopTable_MplsTunnelCHopEntry) GetEntityData() *types.CommonEntityData {
    mplsTunnelCHopEntry.EntityData.YFilter = mplsTunnelCHopEntry.YFilter
    mplsTunnelCHopEntry.EntityData.YangName = "mplsTunnelCHopEntry"
    mplsTunnelCHopEntry.EntityData.BundleName = "cisco_ios_xe"
    mplsTunnelCHopEntry.EntityData.ParentYangName = "mplsTunnelCHopTable"
    mplsTunnelCHopEntry.EntityData.SegmentPath = "mplsTunnelCHopEntry" + types.AddKeyToken(mplsTunnelCHopEntry.MplsTunnelCHopListIndex, "mplsTunnelCHopListIndex") + types.AddKeyToken(mplsTunnelCHopEntry.MplsTunnelCHopIndex, "mplsTunnelCHopIndex")
    mplsTunnelCHopEntry.EntityData.AbsolutePath = "MPLS-TE-STD-MIB:MPLS-TE-STD-MIB/mplsTunnelCHopTable/" + mplsTunnelCHopEntry.EntityData.SegmentPath
    mplsTunnelCHopEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsTunnelCHopEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsTunnelCHopEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsTunnelCHopEntry.EntityData.Children = types.NewOrderedMap()
    mplsTunnelCHopEntry.EntityData.Leafs = types.NewOrderedMap()
    mplsTunnelCHopEntry.EntityData.Leafs.Append("mplsTunnelCHopListIndex", types.YLeaf{"MplsTunnelCHopListIndex", mplsTunnelCHopEntry.MplsTunnelCHopListIndex})
    mplsTunnelCHopEntry.EntityData.Leafs.Append("mplsTunnelCHopIndex", types.YLeaf{"MplsTunnelCHopIndex", mplsTunnelCHopEntry.MplsTunnelCHopIndex})
    mplsTunnelCHopEntry.EntityData.Leafs.Append("mplsTunnelCHopAddrType", types.YLeaf{"MplsTunnelCHopAddrType", mplsTunnelCHopEntry.MplsTunnelCHopAddrType})
    mplsTunnelCHopEntry.EntityData.Leafs.Append("mplsTunnelCHopIpAddr", types.YLeaf{"MplsTunnelCHopIpAddr", mplsTunnelCHopEntry.MplsTunnelCHopIpAddr})
    mplsTunnelCHopEntry.EntityData.Leafs.Append("mplsTunnelCHopIpPrefixLen", types.YLeaf{"MplsTunnelCHopIpPrefixLen", mplsTunnelCHopEntry.MplsTunnelCHopIpPrefixLen})
    mplsTunnelCHopEntry.EntityData.Leafs.Append("mplsTunnelCHopAsNumber", types.YLeaf{"MplsTunnelCHopAsNumber", mplsTunnelCHopEntry.MplsTunnelCHopAsNumber})
    mplsTunnelCHopEntry.EntityData.Leafs.Append("mplsTunnelCHopAddrUnnum", types.YLeaf{"MplsTunnelCHopAddrUnnum", mplsTunnelCHopEntry.MplsTunnelCHopAddrUnnum})
    mplsTunnelCHopEntry.EntityData.Leafs.Append("mplsTunnelCHopLspId", types.YLeaf{"MplsTunnelCHopLspId", mplsTunnelCHopEntry.MplsTunnelCHopLspId})
    mplsTunnelCHopEntry.EntityData.Leafs.Append("mplsTunnelCHopType", types.YLeaf{"MplsTunnelCHopType", mplsTunnelCHopEntry.MplsTunnelCHopType})

    mplsTunnelCHopEntry.EntityData.YListKeys = []string {"MplsTunnelCHopListIndex", "MplsTunnelCHopIndex"}

    return &(mplsTunnelCHopEntry.EntityData)
}

// MPLSTESTDMIB_MplsTunnelCHopTable_MplsTunnelCHopEntry_MplsTunnelCHopType represents strict or loose fashion.
type MPLSTESTDMIB_MplsTunnelCHopTable_MplsTunnelCHopEntry_MplsTunnelCHopType string

const (
    MPLSTESTDMIB_MplsTunnelCHopTable_MplsTunnelCHopEntry_MplsTunnelCHopType_strict MPLSTESTDMIB_MplsTunnelCHopTable_MplsTunnelCHopEntry_MplsTunnelCHopType = "strict"

    MPLSTESTDMIB_MplsTunnelCHopTable_MplsTunnelCHopEntry_MplsTunnelCHopType_loose MPLSTESTDMIB_MplsTunnelCHopTable_MplsTunnelCHopEntry_MplsTunnelCHopType = "loose"
)

// MPLSTESTDMIB_MplsTunnelCRLDPResTable
// The mplsTunnelCRLDPResTable allows a manager to
// specify which CR-LDP-specific resources are desired
// for an MPLS tunnel if that tunnel is signaled using
// CR-LDP. Note that these attributes are in addition
// to those specified in mplsTunnelResourceTable. This
// table also allows several tunnels to point to a
// single entry in this table, implying that these
// tunnels should share resources.
type MPLSTESTDMIB_MplsTunnelCRLDPResTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table represents a set of resources for an MPLS tunnel
    // established using CRLDP (mplsTunnelSignallingProto equal to crldp (3)). An
    // entry can be created by a network administrator or by an SNMP agent as
    // instructed by any MPLS signalling protocol. The type is slice of
    // MPLSTESTDMIB_MplsTunnelCRLDPResTable_MplsTunnelCRLDPResEntry.
    MplsTunnelCRLDPResEntry []*MPLSTESTDMIB_MplsTunnelCRLDPResTable_MplsTunnelCRLDPResEntry
}

func (mplsTunnelCRLDPResTable *MPLSTESTDMIB_MplsTunnelCRLDPResTable) GetEntityData() *types.CommonEntityData {
    mplsTunnelCRLDPResTable.EntityData.YFilter = mplsTunnelCRLDPResTable.YFilter
    mplsTunnelCRLDPResTable.EntityData.YangName = "mplsTunnelCRLDPResTable"
    mplsTunnelCRLDPResTable.EntityData.BundleName = "cisco_ios_xe"
    mplsTunnelCRLDPResTable.EntityData.ParentYangName = "MPLS-TE-STD-MIB"
    mplsTunnelCRLDPResTable.EntityData.SegmentPath = "mplsTunnelCRLDPResTable"
    mplsTunnelCRLDPResTable.EntityData.AbsolutePath = "MPLS-TE-STD-MIB:MPLS-TE-STD-MIB/" + mplsTunnelCRLDPResTable.EntityData.SegmentPath
    mplsTunnelCRLDPResTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsTunnelCRLDPResTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsTunnelCRLDPResTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsTunnelCRLDPResTable.EntityData.Children = types.NewOrderedMap()
    mplsTunnelCRLDPResTable.EntityData.Children.Append("mplsTunnelCRLDPResEntry", types.YChild{"MplsTunnelCRLDPResEntry", nil})
    for i := range mplsTunnelCRLDPResTable.MplsTunnelCRLDPResEntry {
        mplsTunnelCRLDPResTable.EntityData.Children.Append(types.GetSegmentPath(mplsTunnelCRLDPResTable.MplsTunnelCRLDPResEntry[i]), types.YChild{"MplsTunnelCRLDPResEntry", mplsTunnelCRLDPResTable.MplsTunnelCRLDPResEntry[i]})
    }
    mplsTunnelCRLDPResTable.EntityData.Leafs = types.NewOrderedMap()

    mplsTunnelCRLDPResTable.EntityData.YListKeys = []string {}

    return &(mplsTunnelCRLDPResTable.EntityData)
}

// MPLSTESTDMIB_MplsTunnelCRLDPResTable_MplsTunnelCRLDPResEntry
// An entry in this table represents a set of resources
// for an MPLS tunnel established using CRLDP
// (mplsTunnelSignallingProto equal to crldp (3)). An
// entry can be created by a network administrator or
// by an SNMP agent as instructed by any MPLS
// signalling protocol.
type MPLSTESTDMIB_MplsTunnelCRLDPResTable_MplsTunnelCRLDPResEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // mpls_te_std_mib.MPLSTESTDMIB_MplsTunnelResourceTable_MplsTunnelResourceEntry_MplsTunnelResourceIndex
    MplsTunnelResourceIndex interface{}

    // The mean burst size in bytes. The type is interface{} with range:
    // 0..4294967295. Units are bytes.
    MplsTunnelCRLDPResMeanBurstSize interface{}

    // The Excess burst size in bytes. The type is interface{} with range:
    // 0..4294967295. Units are bytes.
    MplsTunnelCRLDPResExBurstSize interface{}

    // The granularity of the availability of committed rate. The type is
    // MplsTunnelCRLDPResFrequency.
    MplsTunnelCRLDPResFrequency interface{}

    // The relative weight for using excess bandwidth above its committed rate. 
    // The value of 0 means that weight is not applicable for the CR-LSP. The type
    // is interface{} with range: 0..255.
    MplsTunnelCRLDPResWeight interface{}

    // The value of the 1 byte Flags conveyed as part of the traffic parameters
    // during the establishment of the CRLSP. The bits in this object are to be
    // interpreted as follows.  +--+--+--+--+--+--+--+--+ | Res
    // |F6|F5|F4|F3|F2|F1| +--+--+--+--+--+--+--+--+  Res - These bits are
    // reserved. Zero on transmission. Ignored on receipt. F1 - Corresponds to the
    // PDR. F2 - Corresponds to the PBS. F3 - Corresponds to the CDR. F4 -
    // Corresponds to the CBS. F5 - Corresponds to the EBS. F6 - Corresponds to
    // the Weight.  Each flag if is a Negotiable Flag corresponding to a Traffic
    // Parameter. The Negotiable Flag value zero denotes Not Negotiable and value
    // one denotes Negotiable. The type is interface{} with range: 0..63.
    MplsTunnelCRLDPResFlags interface{}

    // This variable is used to create, modify, and/or delete a row in this table.
    // When a row in this table is in active(1) state, no objects in that row can
    // be modified by the agent except mplsTunnelCRLDPResRowStatus and
    // mplsTunnelCRLDPResStorageType. The type is RowStatus.
    MplsTunnelCRLDPResRowStatus interface{}

    // The storage type for this CR-LDP Resource entry. Conceptual rows having the
    // value 'permanent' need not allow write-access to any columnar objects in
    // the row. The type is StorageType.
    MplsTunnelCRLDPResStorageType interface{}
}

func (mplsTunnelCRLDPResEntry *MPLSTESTDMIB_MplsTunnelCRLDPResTable_MplsTunnelCRLDPResEntry) GetEntityData() *types.CommonEntityData {
    mplsTunnelCRLDPResEntry.EntityData.YFilter = mplsTunnelCRLDPResEntry.YFilter
    mplsTunnelCRLDPResEntry.EntityData.YangName = "mplsTunnelCRLDPResEntry"
    mplsTunnelCRLDPResEntry.EntityData.BundleName = "cisco_ios_xe"
    mplsTunnelCRLDPResEntry.EntityData.ParentYangName = "mplsTunnelCRLDPResTable"
    mplsTunnelCRLDPResEntry.EntityData.SegmentPath = "mplsTunnelCRLDPResEntry" + types.AddKeyToken(mplsTunnelCRLDPResEntry.MplsTunnelResourceIndex, "mplsTunnelResourceIndex")
    mplsTunnelCRLDPResEntry.EntityData.AbsolutePath = "MPLS-TE-STD-MIB:MPLS-TE-STD-MIB/mplsTunnelCRLDPResTable/" + mplsTunnelCRLDPResEntry.EntityData.SegmentPath
    mplsTunnelCRLDPResEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsTunnelCRLDPResEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsTunnelCRLDPResEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsTunnelCRLDPResEntry.EntityData.Children = types.NewOrderedMap()
    mplsTunnelCRLDPResEntry.EntityData.Leafs = types.NewOrderedMap()
    mplsTunnelCRLDPResEntry.EntityData.Leafs.Append("mplsTunnelResourceIndex", types.YLeaf{"MplsTunnelResourceIndex", mplsTunnelCRLDPResEntry.MplsTunnelResourceIndex})
    mplsTunnelCRLDPResEntry.EntityData.Leafs.Append("mplsTunnelCRLDPResMeanBurstSize", types.YLeaf{"MplsTunnelCRLDPResMeanBurstSize", mplsTunnelCRLDPResEntry.MplsTunnelCRLDPResMeanBurstSize})
    mplsTunnelCRLDPResEntry.EntityData.Leafs.Append("mplsTunnelCRLDPResExBurstSize", types.YLeaf{"MplsTunnelCRLDPResExBurstSize", mplsTunnelCRLDPResEntry.MplsTunnelCRLDPResExBurstSize})
    mplsTunnelCRLDPResEntry.EntityData.Leafs.Append("mplsTunnelCRLDPResFrequency", types.YLeaf{"MplsTunnelCRLDPResFrequency", mplsTunnelCRLDPResEntry.MplsTunnelCRLDPResFrequency})
    mplsTunnelCRLDPResEntry.EntityData.Leafs.Append("mplsTunnelCRLDPResWeight", types.YLeaf{"MplsTunnelCRLDPResWeight", mplsTunnelCRLDPResEntry.MplsTunnelCRLDPResWeight})
    mplsTunnelCRLDPResEntry.EntityData.Leafs.Append("mplsTunnelCRLDPResFlags", types.YLeaf{"MplsTunnelCRLDPResFlags", mplsTunnelCRLDPResEntry.MplsTunnelCRLDPResFlags})
    mplsTunnelCRLDPResEntry.EntityData.Leafs.Append("mplsTunnelCRLDPResRowStatus", types.YLeaf{"MplsTunnelCRLDPResRowStatus", mplsTunnelCRLDPResEntry.MplsTunnelCRLDPResRowStatus})
    mplsTunnelCRLDPResEntry.EntityData.Leafs.Append("mplsTunnelCRLDPResStorageType", types.YLeaf{"MplsTunnelCRLDPResStorageType", mplsTunnelCRLDPResEntry.MplsTunnelCRLDPResStorageType})

    mplsTunnelCRLDPResEntry.EntityData.YListKeys = []string {"MplsTunnelResourceIndex"}

    return &(mplsTunnelCRLDPResEntry.EntityData)
}

// MPLSTESTDMIB_MplsTunnelCRLDPResTable_MplsTunnelCRLDPResEntry_MplsTunnelCRLDPResFrequency represents rate.
type MPLSTESTDMIB_MplsTunnelCRLDPResTable_MplsTunnelCRLDPResEntry_MplsTunnelCRLDPResFrequency string

const (
    MPLSTESTDMIB_MplsTunnelCRLDPResTable_MplsTunnelCRLDPResEntry_MplsTunnelCRLDPResFrequency_unspecified MPLSTESTDMIB_MplsTunnelCRLDPResTable_MplsTunnelCRLDPResEntry_MplsTunnelCRLDPResFrequency = "unspecified"

    MPLSTESTDMIB_MplsTunnelCRLDPResTable_MplsTunnelCRLDPResEntry_MplsTunnelCRLDPResFrequency_frequent MPLSTESTDMIB_MplsTunnelCRLDPResTable_MplsTunnelCRLDPResEntry_MplsTunnelCRLDPResFrequency = "frequent"

    MPLSTESTDMIB_MplsTunnelCRLDPResTable_MplsTunnelCRLDPResEntry_MplsTunnelCRLDPResFrequency_veryFrequent MPLSTESTDMIB_MplsTunnelCRLDPResTable_MplsTunnelCRLDPResEntry_MplsTunnelCRLDPResFrequency = "veryFrequent"
)

