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

    
    Mplstescalars MPLSTESTDMIB_Mplstescalars

    
    Mplsteobjects MPLSTESTDMIB_Mplsteobjects

    // The mplsTunnelTable allows new MPLS tunnels to be created between an LSR
    // and a remote endpoint, and existing tunnels to be reconfigured or removed.
    // Note that only point-to-point tunnel segments are supported, although
    // multipoint-to-point and point- to-multipoint connections are supported by
    // an LSR acting as a cross-connect.  Each MPLS tunnel can thus have one
    // out-segment originating at this LSR and/or one in-segment terminating at
    // this LSR.
    Mplstunneltable MPLSTESTDMIB_Mplstunneltable

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
    Mplstunnelhoptable MPLSTESTDMIB_Mplstunnelhoptable

    // The mplsTunnelResourceTable allows a manager to specify which resources are
    // desired for an MPLS tunnel.  This table also allows several tunnels to
    // point to a single entry in this table, implying that these tunnels should
    // share resources.
    Mplstunnelresourcetable MPLSTESTDMIB_Mplstunnelresourcetable

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
    Mplstunnelarhoptable MPLSTESTDMIB_Mplstunnelarhoptable

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
    Mplstunnelchoptable MPLSTESTDMIB_Mplstunnelchoptable

    // The mplsTunnelCRLDPResTable allows a manager to specify which
    // CR-LDP-specific resources are desired for an MPLS tunnel if that tunnel is
    // signaled using CR-LDP. Note that these attributes are in addition to those
    // specified in mplsTunnelResourceTable. This table also allows several
    // tunnels to point to a single entry in this table, implying that these
    // tunnels should share resources.
    Mplstunnelcrldprestable MPLSTESTDMIB_Mplstunnelcrldprestable
}

func (mPLSTESTDMIB *MPLSTESTDMIB) GetEntityData() *types.CommonEntityData {
    mPLSTESTDMIB.EntityData.YFilter = mPLSTESTDMIB.YFilter
    mPLSTESTDMIB.EntityData.YangName = "MPLS-TE-STD-MIB"
    mPLSTESTDMIB.EntityData.BundleName = "cisco_ios_xe"
    mPLSTESTDMIB.EntityData.ParentYangName = "MPLS-TE-STD-MIB"
    mPLSTESTDMIB.EntityData.SegmentPath = "MPLS-TE-STD-MIB:MPLS-TE-STD-MIB"
    mPLSTESTDMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mPLSTESTDMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mPLSTESTDMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mPLSTESTDMIB.EntityData.Children = make(map[string]types.YChild)
    mPLSTESTDMIB.EntityData.Children["mplsTeScalars"] = types.YChild{"Mplstescalars", &mPLSTESTDMIB.Mplstescalars}
    mPLSTESTDMIB.EntityData.Children["mplsTeObjects"] = types.YChild{"Mplsteobjects", &mPLSTESTDMIB.Mplsteobjects}
    mPLSTESTDMIB.EntityData.Children["mplsTunnelTable"] = types.YChild{"Mplstunneltable", &mPLSTESTDMIB.Mplstunneltable}
    mPLSTESTDMIB.EntityData.Children["mplsTunnelHopTable"] = types.YChild{"Mplstunnelhoptable", &mPLSTESTDMIB.Mplstunnelhoptable}
    mPLSTESTDMIB.EntityData.Children["mplsTunnelResourceTable"] = types.YChild{"Mplstunnelresourcetable", &mPLSTESTDMIB.Mplstunnelresourcetable}
    mPLSTESTDMIB.EntityData.Children["mplsTunnelARHopTable"] = types.YChild{"Mplstunnelarhoptable", &mPLSTESTDMIB.Mplstunnelarhoptable}
    mPLSTESTDMIB.EntityData.Children["mplsTunnelCHopTable"] = types.YChild{"Mplstunnelchoptable", &mPLSTESTDMIB.Mplstunnelchoptable}
    mPLSTESTDMIB.EntityData.Children["mplsTunnelCRLDPResTable"] = types.YChild{"Mplstunnelcrldprestable", &mPLSTESTDMIB.Mplstunnelcrldprestable}
    mPLSTESTDMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mPLSTESTDMIB.EntityData)
}

// MPLSTESTDMIB_Mplstescalars
type MPLSTESTDMIB_Mplstescalars struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The number of tunnels configured on this device. A tunnel is considered
    // configured if the mplsTunnelRowStatus is active(1). The type is interface{}
    // with range: 0..4294967295.
    Mplstunnelconfigured interface{}

    // The number of tunnels active on this device. A tunnel is considered active
    // if the mplsTunnelOperStatus is up(1). The type is interface{} with range:
    // 0..4294967295.
    Mplstunnelactive interface{}

    // The traffic engineering distribution protocol(s) used by this LSR. Note
    // that an LSR may support more than one distribution protocol simultaneously.
    // The type is map[string]bool.
    Mplstunneltedistproto interface{}

    // The maximum number of hops that can be specified for a tunnel on this
    // device. The type is interface{} with range: 0..4294967295.
    Mplstunnelmaxhops interface{}

    // This variable indicates the maximum number of notifications issued per
    // second. If events occur more rapidly, the implementation may simply fail to
    // emit these notifications during that period, or may queue them until an
    // appropriate time. A value of 0 means no throttling is applied and events
    // may be notified at the rate at which they occur. The type is interface{}
    // with range: 0..4294967295.
    Mplstunnelnotificationmaxrate interface{}
}

func (mplstescalars *MPLSTESTDMIB_Mplstescalars) GetEntityData() *types.CommonEntityData {
    mplstescalars.EntityData.YFilter = mplstescalars.YFilter
    mplstescalars.EntityData.YangName = "mplsTeScalars"
    mplstescalars.EntityData.BundleName = "cisco_ios_xe"
    mplstescalars.EntityData.ParentYangName = "MPLS-TE-STD-MIB"
    mplstescalars.EntityData.SegmentPath = "mplsTeScalars"
    mplstescalars.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplstescalars.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplstescalars.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplstescalars.EntityData.Children = make(map[string]types.YChild)
    mplstescalars.EntityData.Leafs = make(map[string]types.YLeaf)
    mplstescalars.EntityData.Leafs["mplsTunnelConfigured"] = types.YLeaf{"Mplstunnelconfigured", mplstescalars.Mplstunnelconfigured}
    mplstescalars.EntityData.Leafs["mplsTunnelActive"] = types.YLeaf{"Mplstunnelactive", mplstescalars.Mplstunnelactive}
    mplstescalars.EntityData.Leafs["mplsTunnelTEDistProto"] = types.YLeaf{"Mplstunneltedistproto", mplstescalars.Mplstunneltedistproto}
    mplstescalars.EntityData.Leafs["mplsTunnelMaxHops"] = types.YLeaf{"Mplstunnelmaxhops", mplstescalars.Mplstunnelmaxhops}
    mplstescalars.EntityData.Leafs["mplsTunnelNotificationMaxRate"] = types.YLeaf{"Mplstunnelnotificationmaxrate", mplstescalars.Mplstunnelnotificationmaxrate}
    return &(mplstescalars.EntityData)
}

// MPLSTESTDMIB_Mplsteobjects
type MPLSTESTDMIB_Mplsteobjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object contains an unused value for mplsTunnelIndex, or a zero to
    // indicate that none exist. Negative values are not allowed, as they do not
    // correspond to valid values of mplsTunnelIndex.  Note that this object
    // offers an unused value for an mplsTunnelIndex value at the ingress side of
    // a tunnel. At other LSRs the value of mplsTunnelIndex SHOULD be taken from
    // the value signaled by the MPLS signaling protocol. The type is interface{}
    // with range: 0..65535.
    Mplstunnelindexnext interface{}

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
    Mplstunnelhoplistindexnext interface{}

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
    Mplstunnelresourceindexnext interface{}

    // If this object is true, then it enables the generation of mplsTunnelUp and
    // mplsTunnelDown traps, otherwise these traps are not emitted. The type is
    // bool.
    Mplstunnelnotificationenable interface{}
}

func (mplsteobjects *MPLSTESTDMIB_Mplsteobjects) GetEntityData() *types.CommonEntityData {
    mplsteobjects.EntityData.YFilter = mplsteobjects.YFilter
    mplsteobjects.EntityData.YangName = "mplsTeObjects"
    mplsteobjects.EntityData.BundleName = "cisco_ios_xe"
    mplsteobjects.EntityData.ParentYangName = "MPLS-TE-STD-MIB"
    mplsteobjects.EntityData.SegmentPath = "mplsTeObjects"
    mplsteobjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsteobjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsteobjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsteobjects.EntityData.Children = make(map[string]types.YChild)
    mplsteobjects.EntityData.Leafs = make(map[string]types.YLeaf)
    mplsteobjects.EntityData.Leafs["mplsTunnelIndexNext"] = types.YLeaf{"Mplstunnelindexnext", mplsteobjects.Mplstunnelindexnext}
    mplsteobjects.EntityData.Leafs["mplsTunnelHopListIndexNext"] = types.YLeaf{"Mplstunnelhoplistindexnext", mplsteobjects.Mplstunnelhoplistindexnext}
    mplsteobjects.EntityData.Leafs["mplsTunnelResourceIndexNext"] = types.YLeaf{"Mplstunnelresourceindexnext", mplsteobjects.Mplstunnelresourceindexnext}
    mplsteobjects.EntityData.Leafs["mplsTunnelNotificationEnable"] = types.YLeaf{"Mplstunnelnotificationenable", mplsteobjects.Mplstunnelnotificationenable}
    return &(mplsteobjects.EntityData)
}

// MPLSTESTDMIB_Mplstunneltable
// The mplsTunnelTable allows new MPLS tunnels to be
// created between an LSR and a remote endpoint, and
// existing tunnels to be reconfigured or removed.
// Note that only point-to-point tunnel segments are
// supported, although multipoint-to-point and point-
// to-multipoint connections are supported by an LSR
// acting as a cross-connect.  Each MPLS tunnel can
// thus have one out-segment originating at this LSR
// and/or one in-segment terminating at this LSR.
type MPLSTESTDMIB_Mplstunneltable struct {
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
    // MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry.
    Mplstunnelentry []MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry
}

func (mplstunneltable *MPLSTESTDMIB_Mplstunneltable) GetEntityData() *types.CommonEntityData {
    mplstunneltable.EntityData.YFilter = mplstunneltable.YFilter
    mplstunneltable.EntityData.YangName = "mplsTunnelTable"
    mplstunneltable.EntityData.BundleName = "cisco_ios_xe"
    mplstunneltable.EntityData.ParentYangName = "MPLS-TE-STD-MIB"
    mplstunneltable.EntityData.SegmentPath = "mplsTunnelTable"
    mplstunneltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplstunneltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplstunneltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplstunneltable.EntityData.Children = make(map[string]types.YChild)
    mplstunneltable.EntityData.Children["mplsTunnelEntry"] = types.YChild{"Mplstunnelentry", nil}
    for i := range mplstunneltable.Mplstunnelentry {
        mplstunneltable.EntityData.Children[types.GetSegmentPath(&mplstunneltable.Mplstunnelentry[i])] = types.YChild{"Mplstunnelentry", &mplstunneltable.Mplstunnelentry[i]}
    }
    mplstunneltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mplstunneltable.EntityData)
}

// MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry
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
type MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Uniquely identifies a set of tunnel instances
    // between a pair of ingress and egress LSRs. Managers should obtain new
    // values for row creation in this table by reading mplsTunnelIndexNext. When
    // the MPLS signalling protocol is rsvp(2) this value SHOULD be equal to the
    // value signaled in the Tunnel Id of the Session object. When the MPLS
    // signalling protocol is crldp(3) this value SHOULD be equal to the value
    // signaled in the LSP ID. The type is interface{} with range: 0..65535.
    Mplstunnelindex interface{}

    // This attribute is a key. Uniquely identifies a particular instance of a
    // tunnel between a pair of ingress and egress LSRs. It is useful to identify
    // multiple instances of tunnels for the purposes of backup and parallel
    // tunnels. When the MPLS signaling protocol is rsvp(2) this value SHOULD be
    // equal to the LSP Id of the Sender Template object. When the signaling
    // protocol is crldp(3) there is no equivalent signaling object. The type is
    // interface{} with range: 0..4294967295.
    Mplstunnelinstance interface{}

    // This attribute is a key. Identity of the ingress LSR associated with this
    // tunnel instance. When the MPLS signalling protocol is rsvp(2) this value
    // SHOULD be equal to the Tunnel Sender Address in the Sender Template object
    // and MAY be equal to the Extended Tunnel Id field in the SESSION object.
    // When the MPLS signalling protocol is crldp(3) this value SHOULD be equal to
    // the Ingress LSR Router ID field in the LSPID TLV object. The type is
    // interface{} with range: 0..4294967295.
    Mplstunnelingresslsrid interface{}

    // This attribute is a key. Identity of the egress LSR associated with this
    // tunnel instance. The type is interface{} with range: 0..4294967295.
    Mplstunnelegresslsrid interface{}

    // The canonical name assigned to the tunnel. This name can be used to refer
    // to the tunnel on the LSR's console port.  If mplsTunnelIsIf is set to true
    // then the ifName of the interface corresponding to this tunnel should have a
    // value equal to mplsTunnelName.  Also see the description of ifName in RFC
    // 2863. The type is string.
    Mplstunnelname interface{}

    // A textual string containing information about the tunnel.  If there is no
    // description this object contains a zero length string. This object is may
    // not be signaled by MPLS signaling protocols, consequentally the value of
    // this object at transit and egress LSRs MAY be automatically generated or
    // absent. The type is string.
    Mplstunneldescr interface{}

    // Denotes whether or not this tunnel corresponds to an interface represented
    // in the interfaces group table. Note that if this variable is set to true
    // then the ifName of the interface corresponding to this tunnel should have a
    // value equal to mplsTunnelName.  Also see the description of ifName in RFC
    // 2863.  This object is meaningful only at the ingress and egress LSRs. The
    // type is bool.
    Mplstunnelisif interface{}

    // If mplsTunnelIsIf is set to true, then this value contains the LSR-assigned
    // ifIndex which corresponds to an entry in the interfaces table.  Otherwise
    // this variable should contain the value of zero indicating that a valid
    // ifIndex was not assigned to this tunnel interface. The type is interface{}
    // with range: 0..2147483647.
    Mplstunnelifindex interface{}

    // Denotes the entity that created and is responsible for managing this
    // tunnel. This column is automatically filled by the agent on creation of a
    // row. The type is MplsOwner.
    Mplstunnelowner interface{}

    // This value signifies the role that this tunnel entry/instance represents.
    // This value MUST be set to head(1) at the originating point of the tunnel.
    // This value MUST be set to transit(2) at transit points along the tunnel, if
    // transit points are supported. This value MUST be set to tail(3) at the
    // terminating point of the tunnel if tunnel tails are supported.  The value
    // headTail(4) is provided for tunnels that begin and end on the same LSR. The
    // type is Mplstunnelrole.
    Mplstunnelrole interface{}

    // This variable points to a row in the mplsXCTable. This table identifies the
    // segments that compose this tunnel, their characteristics, and relationships
    // to each other. A value of zeroDotZero indicates that no LSP has been
    // associated with this tunnel yet. The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Mplstunnelxcpointer interface{}

    // The signalling protocol, if any, used to setup this tunnel. The type is
    // Mplstunnelsignallingproto.
    Mplstunnelsignallingproto interface{}

    // Indicates the setup priority of this tunnel. The type is interface{} with
    // range: 0..7.
    Mplstunnelsetupprio interface{}

    // Indicates the holding priority for this tunnel. The type is interface{}
    // with range: 0..7.
    Mplstunnelholdingprio interface{}

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
    Mplstunnelsessionattributes interface{}

    // Indicates that the local repair mechanism is in use to maintain this tunnel
    // (usually in the face of an outage of the link it was previously routed
    // over). The type is bool.
    Mplstunnellocalprotectinuse interface{}

    // This variable represents a pointer to the traffic parameter specification
    // for this tunnel.  This value may point at an entry in the
    // mplsTunnelResourceEntry to indicate which mplsTunnelResourceEntry is to be
    // assigned to this LSP instance.  This value may optionally point at an
    // externally defined traffic parameter specification table.  A value of
    // zeroDotZero indicates best-effort treatment.  By having the same value of
    // this object, two or more LSPs can indicate resource sharing. The type is
    // string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Mplstunnelresourcepointer interface{}

    // Specifies the instance index of the primary instance of this tunnel. More
    // details of the definition of tunnel instances and the primary tunnel
    // instance can be found in the description of the TEXTUAL-CONVENTION
    // MplsTunnelInstanceIndex. The type is interface{} with range: 0..4294967295.
    Mplstunnelprimaryinstance interface{}

    // This value indicates which priority, in descending order, with 0 indicating
    // the lowest priority, within a group of tunnel instances. A group of tunnel
    // instances is defined as a set of LSPs with the same mplsTunnelIndex in this
    // table, but with a different mplsTunnelInstance. Tunnel instance priorities
    // are used to denote the priority at which a particular tunnel instance will
    // supercede another. Instances of tunnels containing the same
    // mplsTunnelInstancePriority will be used for load sharing. The type is
    // interface{} with range: 0..4294967295.
    Mplstunnelinstancepriority interface{}

    // Index into the mplsTunnelHopTable entry that specifies the explicit route
    // hops for this tunnel. This object is meaningful only at the head-end of the
    // tunnel. The type is interface{} with range: 0..4294967295.
    Mplstunnelhoptableindex interface{}

    // This value denotes the configured path that was chosen for this tunnel.
    // This value reflects the secondary index into mplsTunnelHopTable. This path
    // may not exactly match the one in mplsTunnelARHopTable due to the fact that
    // some CSPF modification may have taken place. See mplsTunnelARHopTable for
    // the actual path being taken by the tunnel. A value of zero denotes that no
    // path is currently in use or available. The type is interface{} with range:
    // 0..4294967295.
    Mplstunnelpathinuse interface{}

    // Index into the mplsTunnelARHopTable entry that specifies the actual hops
    // traversed by the tunnel. This is automatically updated by the agent when
    // the actual hops becomes available. The type is interface{} with range:
    // 0..4294967295.
    Mplstunnelarhoptableindex interface{}

    // Index into the mplsTunnelCHopTable entry that specifies the computed hops
    // traversed by the tunnel. This is automatically updated by the agent when
    // computed hops become available or when computed hops get modified. The type
    // is interface{} with range: 0..4294967295.
    Mplstunnelchoptableindex interface{}

    // A link satisfies the include-any constraint if and only if the constraint
    // is zero, or the link and the constraint have a resource class in common.
    // The type is interface{} with range: 0..4294967295.
    Mplstunnelincludeanyaffinity interface{}

    // A link satisfies the include-all constraint if and only if the link
    // contains all of the administrative groups specified in the constraint. The
    // type is interface{} with range: 0..4294967295.
    Mplstunnelincludeallaffinity interface{}

    // A link satisfies the exclude-any constraint if and only if the link
    // contains none of the administrative groups specified in the constraint. The
    // type is interface{} with range: 0..4294967295.
    Mplstunnelexcludeanyaffinity interface{}

    // This value represents the aggregate up time for all instances of this
    // tunnel, if available. If this value is unavailable, it MUST return a value
    // of 0. The type is interface{} with range: 0..4294967295.
    Mplstunneltotaluptime interface{}

    // This value identifies the total time that this tunnel instance's operStatus
    // has been Up(1). The type is interface{} with range: 0..4294967295.
    Mplstunnelinstanceuptime interface{}

    // Specifies the total time the primary instance of this tunnel has been
    // active. The primary instance of this tunnel is defined in
    // mplsTunnelPrimaryInstance. The type is interface{} with range:
    // 0..4294967295.
    Mplstunnelprimaryuptime interface{}

    // Specifies the number of times the actual path for this tunnel instance has
    // changed. The type is interface{} with range: 0..4294967295.
    Mplstunnelpathchanges interface{}

    // Specifies the time since the last change to the actual path for this tunnel
    // instance. The type is interface{} with range: 0..4294967295.
    Mplstunnellastpathchange interface{}

    // Specifies the value of SysUpTime when the first instance of this tunnel
    // came into existence. That is, when the value of mplsTunnelOperStatus was
    // first set to up(1). The type is interface{} with range: 0..4294967295.
    Mplstunnelcreationtime interface{}

    // Specifies the number of times the state (mplsTunnelOperStatus) of this
    // tunnel instance has changed. The type is interface{} with range:
    // 0..4294967295.
    Mplstunnelstatetransitions interface{}

    // Indicates the desired operational status of this tunnel. The type is
    // Mplstunneladminstatus.
    Mplstunneladminstatus interface{}

    // Indicates the actual operational status of this tunnel, which is typically
    // but not limited to, a function of the state of individual segments of this
    // tunnel. The type is Mplstunneloperstatus.
    Mplstunneloperstatus interface{}

    // This variable is used to create, modify, and/or delete a row in this table.
    // When a row in this table is in active(1) state, no objects in that row can
    // be modified by the agent except mplsTunnelAdminStatus, mplsTunnelRowStatus
    // and mplsTunnelStorageType. The type is RowStatus.
    Mplstunnelrowstatus interface{}

    // The storage type for this tunnel entry. Conceptual rows having the value
    // 'permanent' need not allow write-access to any columnar objects in the row.
    // The type is StorageType.
    Mplstunnelstoragetype interface{}

    // Number of packets forwarded by the tunnel. This object should represents
    // the 32-bit value of the least significant part of the 64-bit value if both
    // mplsTunnelPerfHCPackets is returned. The type is interface{} with range:
    // 0..4294967295.
    Mplstunnelperfpackets interface{}

    // High capacity counter for number of packets forwarded by the tunnel. . The
    // type is interface{} with range: 0..18446744073709551615.
    Mplstunnelperfhcpackets interface{}

    // Number of packets dropped because of errors or for other reasons. The type
    // is interface{} with range: 0..4294967295.
    Mplstunnelperferrors interface{}

    // Number of bytes forwarded by the tunnel. This object should represents the
    // 32-bit value of the least significant part of the 64-bit value if both
    // mplsTunnelPerfHCBytes is returned. The type is interface{} with range:
    // 0..4294967295.
    Mplstunnelperfbytes interface{}

    // High capacity counter for number of bytes forwarded by the tunnel. The type
    // is interface{} with range: 0..18446744073709551615.
    Mplstunnelperfhcbytes interface{}
}

func (mplstunnelentry *MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry) GetEntityData() *types.CommonEntityData {
    mplstunnelentry.EntityData.YFilter = mplstunnelentry.YFilter
    mplstunnelentry.EntityData.YangName = "mplsTunnelEntry"
    mplstunnelentry.EntityData.BundleName = "cisco_ios_xe"
    mplstunnelentry.EntityData.ParentYangName = "mplsTunnelTable"
    mplstunnelentry.EntityData.SegmentPath = "mplsTunnelEntry" + "[mplsTunnelIndex='" + fmt.Sprintf("%v", mplstunnelentry.Mplstunnelindex) + "']" + "[mplsTunnelInstance='" + fmt.Sprintf("%v", mplstunnelentry.Mplstunnelinstance) + "']" + "[mplsTunnelIngressLSRId='" + fmt.Sprintf("%v", mplstunnelentry.Mplstunnelingresslsrid) + "']" + "[mplsTunnelEgressLSRId='" + fmt.Sprintf("%v", mplstunnelentry.Mplstunnelegresslsrid) + "']"
    mplstunnelentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplstunnelentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplstunnelentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplstunnelentry.EntityData.Children = make(map[string]types.YChild)
    mplstunnelentry.EntityData.Leafs = make(map[string]types.YLeaf)
    mplstunnelentry.EntityData.Leafs["mplsTunnelIndex"] = types.YLeaf{"Mplstunnelindex", mplstunnelentry.Mplstunnelindex}
    mplstunnelentry.EntityData.Leafs["mplsTunnelInstance"] = types.YLeaf{"Mplstunnelinstance", mplstunnelentry.Mplstunnelinstance}
    mplstunnelentry.EntityData.Leafs["mplsTunnelIngressLSRId"] = types.YLeaf{"Mplstunnelingresslsrid", mplstunnelentry.Mplstunnelingresslsrid}
    mplstunnelentry.EntityData.Leafs["mplsTunnelEgressLSRId"] = types.YLeaf{"Mplstunnelegresslsrid", mplstunnelentry.Mplstunnelegresslsrid}
    mplstunnelentry.EntityData.Leafs["mplsTunnelName"] = types.YLeaf{"Mplstunnelname", mplstunnelentry.Mplstunnelname}
    mplstunnelentry.EntityData.Leafs["mplsTunnelDescr"] = types.YLeaf{"Mplstunneldescr", mplstunnelentry.Mplstunneldescr}
    mplstunnelentry.EntityData.Leafs["mplsTunnelIsIf"] = types.YLeaf{"Mplstunnelisif", mplstunnelentry.Mplstunnelisif}
    mplstunnelentry.EntityData.Leafs["mplsTunnelIfIndex"] = types.YLeaf{"Mplstunnelifindex", mplstunnelentry.Mplstunnelifindex}
    mplstunnelentry.EntityData.Leafs["mplsTunnelOwner"] = types.YLeaf{"Mplstunnelowner", mplstunnelentry.Mplstunnelowner}
    mplstunnelentry.EntityData.Leafs["mplsTunnelRole"] = types.YLeaf{"Mplstunnelrole", mplstunnelentry.Mplstunnelrole}
    mplstunnelentry.EntityData.Leafs["mplsTunnelXCPointer"] = types.YLeaf{"Mplstunnelxcpointer", mplstunnelentry.Mplstunnelxcpointer}
    mplstunnelentry.EntityData.Leafs["mplsTunnelSignallingProto"] = types.YLeaf{"Mplstunnelsignallingproto", mplstunnelentry.Mplstunnelsignallingproto}
    mplstunnelentry.EntityData.Leafs["mplsTunnelSetupPrio"] = types.YLeaf{"Mplstunnelsetupprio", mplstunnelentry.Mplstunnelsetupprio}
    mplstunnelentry.EntityData.Leafs["mplsTunnelHoldingPrio"] = types.YLeaf{"Mplstunnelholdingprio", mplstunnelentry.Mplstunnelholdingprio}
    mplstunnelentry.EntityData.Leafs["mplsTunnelSessionAttributes"] = types.YLeaf{"Mplstunnelsessionattributes", mplstunnelentry.Mplstunnelsessionattributes}
    mplstunnelentry.EntityData.Leafs["mplsTunnelLocalProtectInUse"] = types.YLeaf{"Mplstunnellocalprotectinuse", mplstunnelentry.Mplstunnellocalprotectinuse}
    mplstunnelentry.EntityData.Leafs["mplsTunnelResourcePointer"] = types.YLeaf{"Mplstunnelresourcepointer", mplstunnelentry.Mplstunnelresourcepointer}
    mplstunnelentry.EntityData.Leafs["mplsTunnelPrimaryInstance"] = types.YLeaf{"Mplstunnelprimaryinstance", mplstunnelentry.Mplstunnelprimaryinstance}
    mplstunnelentry.EntityData.Leafs["mplsTunnelInstancePriority"] = types.YLeaf{"Mplstunnelinstancepriority", mplstunnelentry.Mplstunnelinstancepriority}
    mplstunnelentry.EntityData.Leafs["mplsTunnelHopTableIndex"] = types.YLeaf{"Mplstunnelhoptableindex", mplstunnelentry.Mplstunnelhoptableindex}
    mplstunnelentry.EntityData.Leafs["mplsTunnelPathInUse"] = types.YLeaf{"Mplstunnelpathinuse", mplstunnelentry.Mplstunnelpathinuse}
    mplstunnelentry.EntityData.Leafs["mplsTunnelARHopTableIndex"] = types.YLeaf{"Mplstunnelarhoptableindex", mplstunnelentry.Mplstunnelarhoptableindex}
    mplstunnelentry.EntityData.Leafs["mplsTunnelCHopTableIndex"] = types.YLeaf{"Mplstunnelchoptableindex", mplstunnelentry.Mplstunnelchoptableindex}
    mplstunnelentry.EntityData.Leafs["mplsTunnelIncludeAnyAffinity"] = types.YLeaf{"Mplstunnelincludeanyaffinity", mplstunnelentry.Mplstunnelincludeanyaffinity}
    mplstunnelentry.EntityData.Leafs["mplsTunnelIncludeAllAffinity"] = types.YLeaf{"Mplstunnelincludeallaffinity", mplstunnelentry.Mplstunnelincludeallaffinity}
    mplstunnelentry.EntityData.Leafs["mplsTunnelExcludeAnyAffinity"] = types.YLeaf{"Mplstunnelexcludeanyaffinity", mplstunnelentry.Mplstunnelexcludeanyaffinity}
    mplstunnelentry.EntityData.Leafs["mplsTunnelTotalUpTime"] = types.YLeaf{"Mplstunneltotaluptime", mplstunnelentry.Mplstunneltotaluptime}
    mplstunnelentry.EntityData.Leafs["mplsTunnelInstanceUpTime"] = types.YLeaf{"Mplstunnelinstanceuptime", mplstunnelentry.Mplstunnelinstanceuptime}
    mplstunnelentry.EntityData.Leafs["mplsTunnelPrimaryUpTime"] = types.YLeaf{"Mplstunnelprimaryuptime", mplstunnelentry.Mplstunnelprimaryuptime}
    mplstunnelentry.EntityData.Leafs["mplsTunnelPathChanges"] = types.YLeaf{"Mplstunnelpathchanges", mplstunnelentry.Mplstunnelpathchanges}
    mplstunnelentry.EntityData.Leafs["mplsTunnelLastPathChange"] = types.YLeaf{"Mplstunnellastpathchange", mplstunnelentry.Mplstunnellastpathchange}
    mplstunnelentry.EntityData.Leafs["mplsTunnelCreationTime"] = types.YLeaf{"Mplstunnelcreationtime", mplstunnelentry.Mplstunnelcreationtime}
    mplstunnelentry.EntityData.Leafs["mplsTunnelStateTransitions"] = types.YLeaf{"Mplstunnelstatetransitions", mplstunnelentry.Mplstunnelstatetransitions}
    mplstunnelentry.EntityData.Leafs["mplsTunnelAdminStatus"] = types.YLeaf{"Mplstunneladminstatus", mplstunnelentry.Mplstunneladminstatus}
    mplstunnelentry.EntityData.Leafs["mplsTunnelOperStatus"] = types.YLeaf{"Mplstunneloperstatus", mplstunnelentry.Mplstunneloperstatus}
    mplstunnelentry.EntityData.Leafs["mplsTunnelRowStatus"] = types.YLeaf{"Mplstunnelrowstatus", mplstunnelentry.Mplstunnelrowstatus}
    mplstunnelentry.EntityData.Leafs["mplsTunnelStorageType"] = types.YLeaf{"Mplstunnelstoragetype", mplstunnelentry.Mplstunnelstoragetype}
    mplstunnelentry.EntityData.Leafs["mplsTunnelPerfPackets"] = types.YLeaf{"Mplstunnelperfpackets", mplstunnelentry.Mplstunnelperfpackets}
    mplstunnelentry.EntityData.Leafs["mplsTunnelPerfHCPackets"] = types.YLeaf{"Mplstunnelperfhcpackets", mplstunnelentry.Mplstunnelperfhcpackets}
    mplstunnelentry.EntityData.Leafs["mplsTunnelPerfErrors"] = types.YLeaf{"Mplstunnelperferrors", mplstunnelentry.Mplstunnelperferrors}
    mplstunnelentry.EntityData.Leafs["mplsTunnelPerfBytes"] = types.YLeaf{"Mplstunnelperfbytes", mplstunnelentry.Mplstunnelperfbytes}
    mplstunnelentry.EntityData.Leafs["mplsTunnelPerfHCBytes"] = types.YLeaf{"Mplstunnelperfhcbytes", mplstunnelentry.Mplstunnelperfhcbytes}
    return &(mplstunnelentry.EntityData)
}

// MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry_Mplstunneladminstatus represents tunnel.
type MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry_Mplstunneladminstatus string

const (
    MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry_Mplstunneladminstatus_up MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry_Mplstunneladminstatus = "up"

    MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry_Mplstunneladminstatus_down MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry_Mplstunneladminstatus = "down"

    MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry_Mplstunneladminstatus_testing MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry_Mplstunneladminstatus = "testing"
)

// MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry_Mplstunneloperstatus represents this tunnel.
type MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry_Mplstunneloperstatus string

const (
    MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry_Mplstunneloperstatus_up MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry_Mplstunneloperstatus = "up"

    MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry_Mplstunneloperstatus_down MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry_Mplstunneloperstatus = "down"

    MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry_Mplstunneloperstatus_testing MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry_Mplstunneloperstatus = "testing"

    MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry_Mplstunneloperstatus_unknown MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry_Mplstunneloperstatus = "unknown"

    MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry_Mplstunneloperstatus_dormant MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry_Mplstunneloperstatus = "dormant"

    MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry_Mplstunneloperstatus_notPresent MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry_Mplstunneloperstatus = "notPresent"

    MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry_Mplstunneloperstatus_lowerLayerDown MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry_Mplstunneloperstatus = "lowerLayerDown"
)

// MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry_Mplstunnelrole represents begin and end on the same LSR.
type MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry_Mplstunnelrole string

const (
    MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry_Mplstunnelrole_head MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry_Mplstunnelrole = "head"

    MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry_Mplstunnelrole_transit MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry_Mplstunnelrole = "transit"

    MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry_Mplstunnelrole_tail MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry_Mplstunnelrole = "tail"

    MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry_Mplstunnelrole_headTail MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry_Mplstunnelrole = "headTail"
)

// MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry_Mplstunnelsignallingproto represents tunnel.
type MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry_Mplstunnelsignallingproto string

const (
    MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry_Mplstunnelsignallingproto_none MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry_Mplstunnelsignallingproto = "none"

    MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry_Mplstunnelsignallingproto_rsvp MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry_Mplstunnelsignallingproto = "rsvp"

    MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry_Mplstunnelsignallingproto_crldp MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry_Mplstunnelsignallingproto = "crldp"

    MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry_Mplstunnelsignallingproto_other MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry_Mplstunnelsignallingproto = "other"
)

// MPLSTESTDMIB_Mplstunnelhoptable
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
type MPLSTESTDMIB_Mplstunnelhoptable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table represents a tunnel hop.  An entry is created by a
    // network administrator for signaled ERLSP set up by an MPLS signalling
    // protocol. The type is slice of
    // MPLSTESTDMIB_Mplstunnelhoptable_Mplstunnelhopentry.
    Mplstunnelhopentry []MPLSTESTDMIB_Mplstunnelhoptable_Mplstunnelhopentry
}

func (mplstunnelhoptable *MPLSTESTDMIB_Mplstunnelhoptable) GetEntityData() *types.CommonEntityData {
    mplstunnelhoptable.EntityData.YFilter = mplstunnelhoptable.YFilter
    mplstunnelhoptable.EntityData.YangName = "mplsTunnelHopTable"
    mplstunnelhoptable.EntityData.BundleName = "cisco_ios_xe"
    mplstunnelhoptable.EntityData.ParentYangName = "MPLS-TE-STD-MIB"
    mplstunnelhoptable.EntityData.SegmentPath = "mplsTunnelHopTable"
    mplstunnelhoptable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplstunnelhoptable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplstunnelhoptable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplstunnelhoptable.EntityData.Children = make(map[string]types.YChild)
    mplstunnelhoptable.EntityData.Children["mplsTunnelHopEntry"] = types.YChild{"Mplstunnelhopentry", nil}
    for i := range mplstunnelhoptable.Mplstunnelhopentry {
        mplstunnelhoptable.EntityData.Children[types.GetSegmentPath(&mplstunnelhoptable.Mplstunnelhopentry[i])] = types.YChild{"Mplstunnelhopentry", &mplstunnelhoptable.Mplstunnelhopentry[i]}
    }
    mplstunnelhoptable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mplstunnelhoptable.EntityData)
}

// MPLSTESTDMIB_Mplstunnelhoptable_Mplstunnelhopentry
// An entry in this table represents a tunnel hop.  An
// entry is created by a network administrator for
// signaled ERLSP set up by an MPLS signalling
// protocol.
type MPLSTESTDMIB_Mplstunnelhoptable_Mplstunnelhopentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Primary index into this table identifying a
    // particular explicit route object. The type is interface{} with range:
    // 1..4294967295.
    Mplstunnelhoplistindex interface{}

    // This attribute is a key. Secondary index into this table identifying a
    // particular group of hops representing a particular configured path. This is
    // otherwise known as a path option. The type is interface{} with range:
    // 1..4294967295.
    Mplstunnelhoppathoptionindex interface{}

    // This attribute is a key. Tertiary index into this table identifying a
    // particular hop. The type is interface{} with range: 1..4294967295.
    Mplstunnelhopindex interface{}

    // The Hop Address Type of this tunnel hop.  The value of this object cannot
    // be changed if the value of the corresponding mplsTunnelHopRowStatus object
    // is 'active'.  Note that lspid(5) is a valid option only for tunnels
    // signaled via CRLDP. The type is TeHopAddressType.
    Mplstunnelhopaddrtype interface{}

    // The Tunnel Hop Address for this tunnel hop.  The type of this address is
    // determined by the value of the corresponding mplsTunnelHopAddrType.  The
    // value of this object cannot be changed if the value of the corresponding
    // mplsTunnelHopRowStatus object is 'active'. The type is string with length:
    // 0..32.
    Mplstunnelhopipaddr interface{}

    // If mplsTunnelHopAddrType is set to ipv4(1) or ipv6(2), then this value will
    // contain an appropriate prefix length for the IP address in object
    // mplsTunnelHopIpAddr. Otherwise this value is irrelevant and should be
    // ignored. The type is interface{} with range: 0..2040.
    Mplstunnelhopipprefixlen interface{}

    // If mplsTunnelHopAddrType is set to asnumber(3), then this value will
    // contain the AS number of this hop. Otherwise the agent should set this
    // object to zero- length string and the manager should ignore this. The type
    // is string with length: 4.
    Mplstunnelhopasnumber interface{}

    // If mplsTunnelHopAddrType is set to unnum(4), then this value will contain
    // the interface identifier of the unnumbered interface for this hop. This
    // object should be used in conjunction with mplsTunnelHopIpAddress which
    // would contain the LSR Router ID in this case. Otherwise the agent should
    // set this object to zero-length string and the manager should ignore this.
    // The type is string with length: 4.
    Mplstunnelhopaddrunnum interface{}

    // If mplsTunnelHopAddrType is set to lspid(5), then this value will contain
    // the LSPID of a tunnel of this hop. The present tunnel being configured is
    // tunneled through this hop (using label stacking). This object is otherwise
    // insignificant and should contain a value of 0 to indicate this fact. The
    // type is string with length: 2 | 6.
    Mplstunnelhoplspid interface{}

    // Denotes whether this tunnel hop is routed in a strict or loose fashion. The
    // value of this object has no meaning if the mplsTunnelHopInclude object is
    // set to 'false'. The type is Mplstunnelhoptype.
    Mplstunnelhoptype interface{}

    // If this value is set to true, then this indicates that this hop must be
    // included in the tunnel's path. If this value is set to 'false', then this
    // hop must be avoided when calculating the path for this tunnel. The default
    // value of this object is 'true', so that by default all indicated hops are
    // included in the CSPF path computation. If this object is set to 'false' the
    // value of mplsTunnelHopType should be ignored. The type is bool.
    Mplstunnelhopinclude interface{}

    // The description of this series of hops as they relate to the specified path
    // option. The value of this object SHOULD be the same for each hop in the
    // series that comprises a path option. The type is string.
    Mplstunnelhoppathoptionname interface{}

    // If this value is set to dynamic, then the user should only specify the
    // source and destination of the path and expect that the CSPF will calculate
    // the remainder of the path.  If this value is set to explicit, the user
    // should specify the entire path for the tunnel to take.  This path may
    // contain strict or loose hops.  Each hop along a specific path SHOULD have
    // this object set to the same value. The type is Mplstunnelhopentrypathcomp.
    Mplstunnelhopentrypathcomp interface{}

    // This variable is used to create, modify, and/or delete a row in this table.
    // When a row in this table is in active(1) state, no objects in that row can
    // be modified by the agent except mplsTunnelHopRowStatus and
    // mplsTunnelHopStorageType. The type is RowStatus.
    Mplstunnelhoprowstatus interface{}

    // The storage type for this Hop entry. Conceptual rows having the value
    // 'permanent' need not allow write-access to any columnar objects in the row.
    // The type is StorageType.
    Mplstunnelhopstoragetype interface{}
}

func (mplstunnelhopentry *MPLSTESTDMIB_Mplstunnelhoptable_Mplstunnelhopentry) GetEntityData() *types.CommonEntityData {
    mplstunnelhopentry.EntityData.YFilter = mplstunnelhopentry.YFilter
    mplstunnelhopentry.EntityData.YangName = "mplsTunnelHopEntry"
    mplstunnelhopentry.EntityData.BundleName = "cisco_ios_xe"
    mplstunnelhopentry.EntityData.ParentYangName = "mplsTunnelHopTable"
    mplstunnelhopentry.EntityData.SegmentPath = "mplsTunnelHopEntry" + "[mplsTunnelHopListIndex='" + fmt.Sprintf("%v", mplstunnelhopentry.Mplstunnelhoplistindex) + "']" + "[mplsTunnelHopPathOptionIndex='" + fmt.Sprintf("%v", mplstunnelhopentry.Mplstunnelhoppathoptionindex) + "']" + "[mplsTunnelHopIndex='" + fmt.Sprintf("%v", mplstunnelhopentry.Mplstunnelhopindex) + "']"
    mplstunnelhopentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplstunnelhopentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplstunnelhopentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplstunnelhopentry.EntityData.Children = make(map[string]types.YChild)
    mplstunnelhopentry.EntityData.Leafs = make(map[string]types.YLeaf)
    mplstunnelhopentry.EntityData.Leafs["mplsTunnelHopListIndex"] = types.YLeaf{"Mplstunnelhoplistindex", mplstunnelhopentry.Mplstunnelhoplistindex}
    mplstunnelhopentry.EntityData.Leafs["mplsTunnelHopPathOptionIndex"] = types.YLeaf{"Mplstunnelhoppathoptionindex", mplstunnelhopentry.Mplstunnelhoppathoptionindex}
    mplstunnelhopentry.EntityData.Leafs["mplsTunnelHopIndex"] = types.YLeaf{"Mplstunnelhopindex", mplstunnelhopentry.Mplstunnelhopindex}
    mplstunnelhopentry.EntityData.Leafs["mplsTunnelHopAddrType"] = types.YLeaf{"Mplstunnelhopaddrtype", mplstunnelhopentry.Mplstunnelhopaddrtype}
    mplstunnelhopentry.EntityData.Leafs["mplsTunnelHopIpAddr"] = types.YLeaf{"Mplstunnelhopipaddr", mplstunnelhopentry.Mplstunnelhopipaddr}
    mplstunnelhopentry.EntityData.Leafs["mplsTunnelHopIpPrefixLen"] = types.YLeaf{"Mplstunnelhopipprefixlen", mplstunnelhopentry.Mplstunnelhopipprefixlen}
    mplstunnelhopentry.EntityData.Leafs["mplsTunnelHopAsNumber"] = types.YLeaf{"Mplstunnelhopasnumber", mplstunnelhopentry.Mplstunnelhopasnumber}
    mplstunnelhopentry.EntityData.Leafs["mplsTunnelHopAddrUnnum"] = types.YLeaf{"Mplstunnelhopaddrunnum", mplstunnelhopentry.Mplstunnelhopaddrunnum}
    mplstunnelhopentry.EntityData.Leafs["mplsTunnelHopLspId"] = types.YLeaf{"Mplstunnelhoplspid", mplstunnelhopentry.Mplstunnelhoplspid}
    mplstunnelhopentry.EntityData.Leafs["mplsTunnelHopType"] = types.YLeaf{"Mplstunnelhoptype", mplstunnelhopentry.Mplstunnelhoptype}
    mplstunnelhopentry.EntityData.Leafs["mplsTunnelHopInclude"] = types.YLeaf{"Mplstunnelhopinclude", mplstunnelhopentry.Mplstunnelhopinclude}
    mplstunnelhopentry.EntityData.Leafs["mplsTunnelHopPathOptionName"] = types.YLeaf{"Mplstunnelhoppathoptionname", mplstunnelhopentry.Mplstunnelhoppathoptionname}
    mplstunnelhopentry.EntityData.Leafs["mplsTunnelHopEntryPathComp"] = types.YLeaf{"Mplstunnelhopentrypathcomp", mplstunnelhopentry.Mplstunnelhopentrypathcomp}
    mplstunnelhopentry.EntityData.Leafs["mplsTunnelHopRowStatus"] = types.YLeaf{"Mplstunnelhoprowstatus", mplstunnelhopentry.Mplstunnelhoprowstatus}
    mplstunnelhopentry.EntityData.Leafs["mplsTunnelHopStorageType"] = types.YLeaf{"Mplstunnelhopstoragetype", mplstunnelhopentry.Mplstunnelhopstoragetype}
    return &(mplstunnelhopentry.EntityData)
}

// MPLSTESTDMIB_Mplstunnelhoptable_Mplstunnelhopentry_Mplstunnelhopentrypathcomp represents path SHOULD have this object set to the same value
type MPLSTESTDMIB_Mplstunnelhoptable_Mplstunnelhopentry_Mplstunnelhopentrypathcomp string

const (
    MPLSTESTDMIB_Mplstunnelhoptable_Mplstunnelhopentry_Mplstunnelhopentrypathcomp_dynamic MPLSTESTDMIB_Mplstunnelhoptable_Mplstunnelhopentry_Mplstunnelhopentrypathcomp = "dynamic"

    MPLSTESTDMIB_Mplstunnelhoptable_Mplstunnelhopentry_Mplstunnelhopentrypathcomp_explicit MPLSTESTDMIB_Mplstunnelhoptable_Mplstunnelhopentry_Mplstunnelhopentrypathcomp = "explicit"
)

// MPLSTESTDMIB_Mplstunnelhoptable_Mplstunnelhopentry_Mplstunnelhoptype represents is set to 'false'.
type MPLSTESTDMIB_Mplstunnelhoptable_Mplstunnelhopentry_Mplstunnelhoptype string

const (
    MPLSTESTDMIB_Mplstunnelhoptable_Mplstunnelhopentry_Mplstunnelhoptype_strict MPLSTESTDMIB_Mplstunnelhoptable_Mplstunnelhopentry_Mplstunnelhoptype = "strict"

    MPLSTESTDMIB_Mplstunnelhoptable_Mplstunnelhopentry_Mplstunnelhoptype_loose MPLSTESTDMIB_Mplstunnelhoptable_Mplstunnelhopentry_Mplstunnelhoptype = "loose"
)

// MPLSTESTDMIB_Mplstunnelresourcetable
// The mplsTunnelResourceTable allows a manager to
// specify which resources are desired for an MPLS
// tunnel.  This table also allows several tunnels to
// point to a single entry in this table, implying
// that these tunnels should share resources.
type MPLSTESTDMIB_Mplstunnelresourcetable struct {
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
    // MPLSTESTDMIB_Mplstunnelresourcetable_Mplstunnelresourceentry.
    Mplstunnelresourceentry []MPLSTESTDMIB_Mplstunnelresourcetable_Mplstunnelresourceentry
}

func (mplstunnelresourcetable *MPLSTESTDMIB_Mplstunnelresourcetable) GetEntityData() *types.CommonEntityData {
    mplstunnelresourcetable.EntityData.YFilter = mplstunnelresourcetable.YFilter
    mplstunnelresourcetable.EntityData.YangName = "mplsTunnelResourceTable"
    mplstunnelresourcetable.EntityData.BundleName = "cisco_ios_xe"
    mplstunnelresourcetable.EntityData.ParentYangName = "MPLS-TE-STD-MIB"
    mplstunnelresourcetable.EntityData.SegmentPath = "mplsTunnelResourceTable"
    mplstunnelresourcetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplstunnelresourcetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplstunnelresourcetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplstunnelresourcetable.EntityData.Children = make(map[string]types.YChild)
    mplstunnelresourcetable.EntityData.Children["mplsTunnelResourceEntry"] = types.YChild{"Mplstunnelresourceentry", nil}
    for i := range mplstunnelresourcetable.Mplstunnelresourceentry {
        mplstunnelresourcetable.EntityData.Children[types.GetSegmentPath(&mplstunnelresourcetable.Mplstunnelresourceentry[i])] = types.YChild{"Mplstunnelresourceentry", &mplstunnelresourcetable.Mplstunnelresourceentry[i]}
    }
    mplstunnelresourcetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mplstunnelresourcetable.EntityData)
}

// MPLSTESTDMIB_Mplstunnelresourcetable_Mplstunnelresourceentry
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
type MPLSTESTDMIB_Mplstunnelresourcetable_Mplstunnelresourceentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Uniquely identifies this row. The type is
    // interface{} with range: 1..2147483647.
    Mplstunnelresourceindex interface{}

    // The maximum rate in bits/second.  Note that setting
    // mplsTunnelResourceMaxRate, mplsTunnelResourceMeanRate, and
    // mplsTunnelResourceMaxBurstSize to 0 indicates best- effort treatment. The
    // type is interface{} with range: 0..4294967295. Units are kilobits per
    // second.
    Mplstunnelresourcemaxrate interface{}

    // This object is copied into an instance of mplsTrafficParamMeanRate in the
    // mplsTrafficParamTable. The OID of this table entry is then copied into the
    // corresponding mplsInSegmentTrafficParamPtr. The type is interface{} with
    // range: 0..4294967295. Units are kilobits per second.
    Mplstunnelresourcemeanrate interface{}

    // The maximum burst size in bytes. The type is interface{} with range:
    // 0..4294967295. Units are bytes.
    Mplstunnelresourcemaxburstsize interface{}

    // The mean burst size in bytes.  The implementations which do not implement
    // this variable must return a noSuchObject exception for this object and must
    // not allow a user to set this object. The type is interface{} with range:
    // 0..4294967295. Units are bytes.
    Mplstunnelresourcemeanburstsize interface{}

    // The Excess burst size in bytes.  The implementations which do not implement
    // this variable must return noSuchObject exception for this object and must
    // not allow a user to set this value. The type is interface{} with range:
    // 0..4294967295. Units are bytes.
    Mplstunnelresourceexburstsize interface{}

    // The granularity of the availability of committed rate.  The implementations
    // which do not implement this variable must return unspecified(1) for this
    // value and must not allow a user to set this value. The type is
    // Mplstunnelresourcefrequency.
    Mplstunnelresourcefrequency interface{}

    // The relative weight for using excess bandwidth above its committed rate. 
    // The value of 0 means that weight is not applicable for the CR-LSP. The type
    // is interface{} with range: 0..255.
    Mplstunnelresourceweight interface{}

    // This variable is used to create, modify, and/or delete a row in this table.
    // When a row in this table is in active(1) state, no objects in that row can
    // be modified by the agent except mplsTunnelResourceRowStatus and
    // mplsTunnelResourceStorageType. The type is RowStatus.
    Mplstunnelresourcerowstatus interface{}

    // The storage type for this Hop entry. Conceptual rows having the value
    // 'permanent' need not allow write-access to any columnar objects in the row.
    // The type is StorageType.
    Mplstunnelresourcestoragetype interface{}
}

func (mplstunnelresourceentry *MPLSTESTDMIB_Mplstunnelresourcetable_Mplstunnelresourceentry) GetEntityData() *types.CommonEntityData {
    mplstunnelresourceentry.EntityData.YFilter = mplstunnelresourceentry.YFilter
    mplstunnelresourceentry.EntityData.YangName = "mplsTunnelResourceEntry"
    mplstunnelresourceentry.EntityData.BundleName = "cisco_ios_xe"
    mplstunnelresourceentry.EntityData.ParentYangName = "mplsTunnelResourceTable"
    mplstunnelresourceentry.EntityData.SegmentPath = "mplsTunnelResourceEntry" + "[mplsTunnelResourceIndex='" + fmt.Sprintf("%v", mplstunnelresourceentry.Mplstunnelresourceindex) + "']"
    mplstunnelresourceentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplstunnelresourceentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplstunnelresourceentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplstunnelresourceentry.EntityData.Children = make(map[string]types.YChild)
    mplstunnelresourceentry.EntityData.Leafs = make(map[string]types.YLeaf)
    mplstunnelresourceentry.EntityData.Leafs["mplsTunnelResourceIndex"] = types.YLeaf{"Mplstunnelresourceindex", mplstunnelresourceentry.Mplstunnelresourceindex}
    mplstunnelresourceentry.EntityData.Leafs["mplsTunnelResourceMaxRate"] = types.YLeaf{"Mplstunnelresourcemaxrate", mplstunnelresourceentry.Mplstunnelresourcemaxrate}
    mplstunnelresourceentry.EntityData.Leafs["mplsTunnelResourceMeanRate"] = types.YLeaf{"Mplstunnelresourcemeanrate", mplstunnelresourceentry.Mplstunnelresourcemeanrate}
    mplstunnelresourceentry.EntityData.Leafs["mplsTunnelResourceMaxBurstSize"] = types.YLeaf{"Mplstunnelresourcemaxburstsize", mplstunnelresourceentry.Mplstunnelresourcemaxburstsize}
    mplstunnelresourceentry.EntityData.Leafs["mplsTunnelResourceMeanBurstSize"] = types.YLeaf{"Mplstunnelresourcemeanburstsize", mplstunnelresourceentry.Mplstunnelresourcemeanburstsize}
    mplstunnelresourceentry.EntityData.Leafs["mplsTunnelResourceExBurstSize"] = types.YLeaf{"Mplstunnelresourceexburstsize", mplstunnelresourceentry.Mplstunnelresourceexburstsize}
    mplstunnelresourceentry.EntityData.Leafs["mplsTunnelResourceFrequency"] = types.YLeaf{"Mplstunnelresourcefrequency", mplstunnelresourceentry.Mplstunnelresourcefrequency}
    mplstunnelresourceentry.EntityData.Leafs["mplsTunnelResourceWeight"] = types.YLeaf{"Mplstunnelresourceweight", mplstunnelresourceentry.Mplstunnelresourceweight}
    mplstunnelresourceentry.EntityData.Leafs["mplsTunnelResourceRowStatus"] = types.YLeaf{"Mplstunnelresourcerowstatus", mplstunnelresourceentry.Mplstunnelresourcerowstatus}
    mplstunnelresourceentry.EntityData.Leafs["mplsTunnelResourceStorageType"] = types.YLeaf{"Mplstunnelresourcestoragetype", mplstunnelresourceentry.Mplstunnelresourcestoragetype}
    return &(mplstunnelresourceentry.EntityData)
}

// MPLSTESTDMIB_Mplstunnelresourcetable_Mplstunnelresourceentry_Mplstunnelresourcefrequency represents value and must not allow a user to set this value.
type MPLSTESTDMIB_Mplstunnelresourcetable_Mplstunnelresourceentry_Mplstunnelresourcefrequency string

const (
    MPLSTESTDMIB_Mplstunnelresourcetable_Mplstunnelresourceentry_Mplstunnelresourcefrequency_unspecified MPLSTESTDMIB_Mplstunnelresourcetable_Mplstunnelresourceentry_Mplstunnelresourcefrequency = "unspecified"

    MPLSTESTDMIB_Mplstunnelresourcetable_Mplstunnelresourceentry_Mplstunnelresourcefrequency_frequent MPLSTESTDMIB_Mplstunnelresourcetable_Mplstunnelresourceentry_Mplstunnelresourcefrequency = "frequent"

    MPLSTESTDMIB_Mplstunnelresourcetable_Mplstunnelresourceentry_Mplstunnelresourcefrequency_veryFrequent MPLSTESTDMIB_Mplstunnelresourcetable_Mplstunnelresourceentry_Mplstunnelresourcefrequency = "veryFrequent"
)

// MPLSTESTDMIB_Mplstunnelarhoptable
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
type MPLSTESTDMIB_Mplstunnelarhoptable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table represents a tunnel hop.  An entry is created by the
    // agent for signaled ERLSP set up by an MPLS signalling protocol. The type is
    // slice of MPLSTESTDMIB_Mplstunnelarhoptable_Mplstunnelarhopentry.
    Mplstunnelarhopentry []MPLSTESTDMIB_Mplstunnelarhoptable_Mplstunnelarhopentry
}

func (mplstunnelarhoptable *MPLSTESTDMIB_Mplstunnelarhoptable) GetEntityData() *types.CommonEntityData {
    mplstunnelarhoptable.EntityData.YFilter = mplstunnelarhoptable.YFilter
    mplstunnelarhoptable.EntityData.YangName = "mplsTunnelARHopTable"
    mplstunnelarhoptable.EntityData.BundleName = "cisco_ios_xe"
    mplstunnelarhoptable.EntityData.ParentYangName = "MPLS-TE-STD-MIB"
    mplstunnelarhoptable.EntityData.SegmentPath = "mplsTunnelARHopTable"
    mplstunnelarhoptable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplstunnelarhoptable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplstunnelarhoptable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplstunnelarhoptable.EntityData.Children = make(map[string]types.YChild)
    mplstunnelarhoptable.EntityData.Children["mplsTunnelARHopEntry"] = types.YChild{"Mplstunnelarhopentry", nil}
    for i := range mplstunnelarhoptable.Mplstunnelarhopentry {
        mplstunnelarhoptable.EntityData.Children[types.GetSegmentPath(&mplstunnelarhoptable.Mplstunnelarhopentry[i])] = types.YChild{"Mplstunnelarhopentry", &mplstunnelarhoptable.Mplstunnelarhopentry[i]}
    }
    mplstunnelarhoptable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mplstunnelarhoptable.EntityData)
}

// MPLSTESTDMIB_Mplstunnelarhoptable_Mplstunnelarhopentry
// An entry in this table represents a tunnel hop.  An
// entry is created by the agent for signaled ERLSP
// set up by an MPLS signalling protocol.
type MPLSTESTDMIB_Mplstunnelarhoptable_Mplstunnelarhopentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Primary index into this table identifying a
    // particular recorded hop list. The type is interface{} with range:
    // 1..4294967295.
    Mplstunnelarhoplistindex interface{}

    // This attribute is a key. Secondary index into this table identifying the
    // particular hop. The type is interface{} with range: 1..4294967295.
    Mplstunnelarhopindex interface{}

    // The Hop Address Type of this tunnel hop.  Note that lspid(5) is a valid
    // option only for tunnels signaled via CRLDP. The type is TeHopAddressType.
    Mplstunnelarhopaddrtype interface{}

    // The Tunnel Hop Address for this tunnel hop.  The type of this address is
    // determined by the value of the corresponding mplsTunnelARHopAddrType. If
    // mplsTunnelARHopAddrType is set to unnum(4),  then this value contains the
    // LSR Router ID of the  unnumbered interface. Otherwise the agent SHOULD  set
    // this object to the zero-length string and the  manager should ignore this
    // object. The type is string with length: 0..32.
    Mplstunnelarhopipaddr interface{}

    // If mplsTunnelARHopAddrType is set to unnum(4), then this value will contain
    // the interface identifier of the unnumbered interface for this hop. This
    // object should be used in conjunction with mplsTunnelARHopIpAddr which would
    // contain the LSR Router ID in this case. Otherwise the agent should set this
    // object to zero-length string and the manager should ignore this. The type
    // is string with length: 4.
    Mplstunnelarhopaddrunnum interface{}

    // If mplsTunnelARHopAddrType is set to lspid(5), then this value will contain
    // the LSP ID of this hop. This object is otherwise insignificant and should
    // contain a value of 0 to indicate this fact. The type is string with length:
    // 2 | 6.
    Mplstunnelarhoplspid interface{}
}

func (mplstunnelarhopentry *MPLSTESTDMIB_Mplstunnelarhoptable_Mplstunnelarhopentry) GetEntityData() *types.CommonEntityData {
    mplstunnelarhopentry.EntityData.YFilter = mplstunnelarhopentry.YFilter
    mplstunnelarhopentry.EntityData.YangName = "mplsTunnelARHopEntry"
    mplstunnelarhopentry.EntityData.BundleName = "cisco_ios_xe"
    mplstunnelarhopentry.EntityData.ParentYangName = "mplsTunnelARHopTable"
    mplstunnelarhopentry.EntityData.SegmentPath = "mplsTunnelARHopEntry" + "[mplsTunnelARHopListIndex='" + fmt.Sprintf("%v", mplstunnelarhopentry.Mplstunnelarhoplistindex) + "']" + "[mplsTunnelARHopIndex='" + fmt.Sprintf("%v", mplstunnelarhopentry.Mplstunnelarhopindex) + "']"
    mplstunnelarhopentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplstunnelarhopentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplstunnelarhopentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplstunnelarhopentry.EntityData.Children = make(map[string]types.YChild)
    mplstunnelarhopentry.EntityData.Leafs = make(map[string]types.YLeaf)
    mplstunnelarhopentry.EntityData.Leafs["mplsTunnelARHopListIndex"] = types.YLeaf{"Mplstunnelarhoplistindex", mplstunnelarhopentry.Mplstunnelarhoplistindex}
    mplstunnelarhopentry.EntityData.Leafs["mplsTunnelARHopIndex"] = types.YLeaf{"Mplstunnelarhopindex", mplstunnelarhopentry.Mplstunnelarhopindex}
    mplstunnelarhopentry.EntityData.Leafs["mplsTunnelARHopAddrType"] = types.YLeaf{"Mplstunnelarhopaddrtype", mplstunnelarhopentry.Mplstunnelarhopaddrtype}
    mplstunnelarhopentry.EntityData.Leafs["mplsTunnelARHopIpAddr"] = types.YLeaf{"Mplstunnelarhopipaddr", mplstunnelarhopentry.Mplstunnelarhopipaddr}
    mplstunnelarhopentry.EntityData.Leafs["mplsTunnelARHopAddrUnnum"] = types.YLeaf{"Mplstunnelarhopaddrunnum", mplstunnelarhopentry.Mplstunnelarhopaddrunnum}
    mplstunnelarhopentry.EntityData.Leafs["mplsTunnelARHopLspId"] = types.YLeaf{"Mplstunnelarhoplspid", mplstunnelarhopentry.Mplstunnelarhoplspid}
    return &(mplstunnelarhopentry.EntityData)
}

// MPLSTESTDMIB_Mplstunnelchoptable
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
type MPLSTESTDMIB_Mplstunnelchoptable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table represents a tunnel hop.  An entry in this table is
    // created by a path computation engine using CSPF techniques applied to the
    // information collected by routing protocols and the hops specified in the
    // corresponding mplsTunnelHopTable. The type is slice of
    // MPLSTESTDMIB_Mplstunnelchoptable_Mplstunnelchopentry.
    Mplstunnelchopentry []MPLSTESTDMIB_Mplstunnelchoptable_Mplstunnelchopentry
}

func (mplstunnelchoptable *MPLSTESTDMIB_Mplstunnelchoptable) GetEntityData() *types.CommonEntityData {
    mplstunnelchoptable.EntityData.YFilter = mplstunnelchoptable.YFilter
    mplstunnelchoptable.EntityData.YangName = "mplsTunnelCHopTable"
    mplstunnelchoptable.EntityData.BundleName = "cisco_ios_xe"
    mplstunnelchoptable.EntityData.ParentYangName = "MPLS-TE-STD-MIB"
    mplstunnelchoptable.EntityData.SegmentPath = "mplsTunnelCHopTable"
    mplstunnelchoptable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplstunnelchoptable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplstunnelchoptable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplstunnelchoptable.EntityData.Children = make(map[string]types.YChild)
    mplstunnelchoptable.EntityData.Children["mplsTunnelCHopEntry"] = types.YChild{"Mplstunnelchopentry", nil}
    for i := range mplstunnelchoptable.Mplstunnelchopentry {
        mplstunnelchoptable.EntityData.Children[types.GetSegmentPath(&mplstunnelchoptable.Mplstunnelchopentry[i])] = types.YChild{"Mplstunnelchopentry", &mplstunnelchoptable.Mplstunnelchopentry[i]}
    }
    mplstunnelchoptable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mplstunnelchoptable.EntityData)
}

// MPLSTESTDMIB_Mplstunnelchoptable_Mplstunnelchopentry
// An entry in this table represents a tunnel hop.  An
// entry in this table is created by a path
// computation engine using CSPF techniques applied to
// the information collected by routing protocols and
// the hops specified in the corresponding
// mplsTunnelHopTable.
type MPLSTESTDMIB_Mplstunnelchoptable_Mplstunnelchopentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Primary index into this table identifying a
    // particular computed hop list. The type is interface{} with range:
    // 1..4294967295.
    Mplstunnelchoplistindex interface{}

    // This attribute is a key. Secondary index into this table identifying the
    // particular hop. The type is interface{} with range: 1..4294967295.
    Mplstunnelchopindex interface{}

    // The Hop Address Type of this tunnel hop.  Note that lspid(5) is a valid
    // option only for tunnels signaled via CRLDP. The type is TeHopAddressType.
    Mplstunnelchopaddrtype interface{}

    // The Tunnel Hop Address for this tunnel hop. The type of this address is
    // determined by the  value of the corresponding mplsTunnelCHopAddrType.  If
    // mplsTunnelCHopAddrType is set to unnum(4), then  this value will contain
    // the LSR Router ID of the  unnumbered interface. Otherwise the agent should 
    // set this object to the zero-length string and the  manager SHOULD ignore
    // this object. The type is string with length: 0..32.
    Mplstunnelchopipaddr interface{}

    // If mplsTunnelCHopAddrType is set to ipv4(1) or ipv6(2), then this value
    // will contain an appropriate prefix length for the IP address in object
    // mplsTunnelCHopIpAddr. Otherwise this value is irrelevant and should be
    // ignored. The type is interface{} with range: 0..2040.
    Mplstunnelchopipprefixlen interface{}

    // If mplsTunnelCHopAddrType is set to asnumber(3), then this value will
    // contain the AS number of this hop. Otherwise the agent should set this
    // object to zero-length string and the manager should ignore this. The type
    // is string with length: 4.
    Mplstunnelchopasnumber interface{}

    // If mplsTunnelCHopAddrType is set to unnum(4), then this value will contain
    // the unnumbered interface identifier of this hop. This object should be used
    // in conjunction with mplsTunnelCHopIpAddr which would contain the LSR Router
    // ID in this case. Otherwise the agent should set this object to zero- length
    // string and the manager should ignore this. The type is string with length:
    // 4.
    Mplstunnelchopaddrunnum interface{}

    // If mplsTunnelCHopAddrType is set to lspid(5), then this value will contain
    // the LSP ID of this hop. This object is otherwise insignificant and should
    // contain a value of 0 to indicate this fact. The type is string with length:
    // 2 | 6.
    Mplstunnelchoplspid interface{}

    // Denotes whether this is tunnel hop is routed in a strict or loose fashion.
    // The type is Mplstunnelchoptype.
    Mplstunnelchoptype interface{}
}

func (mplstunnelchopentry *MPLSTESTDMIB_Mplstunnelchoptable_Mplstunnelchopentry) GetEntityData() *types.CommonEntityData {
    mplstunnelchopentry.EntityData.YFilter = mplstunnelchopentry.YFilter
    mplstunnelchopentry.EntityData.YangName = "mplsTunnelCHopEntry"
    mplstunnelchopentry.EntityData.BundleName = "cisco_ios_xe"
    mplstunnelchopentry.EntityData.ParentYangName = "mplsTunnelCHopTable"
    mplstunnelchopentry.EntityData.SegmentPath = "mplsTunnelCHopEntry" + "[mplsTunnelCHopListIndex='" + fmt.Sprintf("%v", mplstunnelchopentry.Mplstunnelchoplistindex) + "']" + "[mplsTunnelCHopIndex='" + fmt.Sprintf("%v", mplstunnelchopentry.Mplstunnelchopindex) + "']"
    mplstunnelchopentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplstunnelchopentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplstunnelchopentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplstunnelchopentry.EntityData.Children = make(map[string]types.YChild)
    mplstunnelchopentry.EntityData.Leafs = make(map[string]types.YLeaf)
    mplstunnelchopentry.EntityData.Leafs["mplsTunnelCHopListIndex"] = types.YLeaf{"Mplstunnelchoplistindex", mplstunnelchopentry.Mplstunnelchoplistindex}
    mplstunnelchopentry.EntityData.Leafs["mplsTunnelCHopIndex"] = types.YLeaf{"Mplstunnelchopindex", mplstunnelchopentry.Mplstunnelchopindex}
    mplstunnelchopentry.EntityData.Leafs["mplsTunnelCHopAddrType"] = types.YLeaf{"Mplstunnelchopaddrtype", mplstunnelchopentry.Mplstunnelchopaddrtype}
    mplstunnelchopentry.EntityData.Leafs["mplsTunnelCHopIpAddr"] = types.YLeaf{"Mplstunnelchopipaddr", mplstunnelchopentry.Mplstunnelchopipaddr}
    mplstunnelchopentry.EntityData.Leafs["mplsTunnelCHopIpPrefixLen"] = types.YLeaf{"Mplstunnelchopipprefixlen", mplstunnelchopentry.Mplstunnelchopipprefixlen}
    mplstunnelchopentry.EntityData.Leafs["mplsTunnelCHopAsNumber"] = types.YLeaf{"Mplstunnelchopasnumber", mplstunnelchopentry.Mplstunnelchopasnumber}
    mplstunnelchopentry.EntityData.Leafs["mplsTunnelCHopAddrUnnum"] = types.YLeaf{"Mplstunnelchopaddrunnum", mplstunnelchopentry.Mplstunnelchopaddrunnum}
    mplstunnelchopentry.EntityData.Leafs["mplsTunnelCHopLspId"] = types.YLeaf{"Mplstunnelchoplspid", mplstunnelchopentry.Mplstunnelchoplspid}
    mplstunnelchopentry.EntityData.Leafs["mplsTunnelCHopType"] = types.YLeaf{"Mplstunnelchoptype", mplstunnelchopentry.Mplstunnelchoptype}
    return &(mplstunnelchopentry.EntityData)
}

// MPLSTESTDMIB_Mplstunnelchoptable_Mplstunnelchopentry_Mplstunnelchoptype represents strict or loose fashion.
type MPLSTESTDMIB_Mplstunnelchoptable_Mplstunnelchopentry_Mplstunnelchoptype string

const (
    MPLSTESTDMIB_Mplstunnelchoptable_Mplstunnelchopentry_Mplstunnelchoptype_strict MPLSTESTDMIB_Mplstunnelchoptable_Mplstunnelchopentry_Mplstunnelchoptype = "strict"

    MPLSTESTDMIB_Mplstunnelchoptable_Mplstunnelchopentry_Mplstunnelchoptype_loose MPLSTESTDMIB_Mplstunnelchoptable_Mplstunnelchopentry_Mplstunnelchoptype = "loose"
)

// MPLSTESTDMIB_Mplstunnelcrldprestable
// The mplsTunnelCRLDPResTable allows a manager to
// specify which CR-LDP-specific resources are desired
// for an MPLS tunnel if that tunnel is signaled using
// CR-LDP. Note that these attributes are in addition
// to those specified in mplsTunnelResourceTable. This
// table also allows several tunnels to point to a
// single entry in this table, implying that these
// tunnels should share resources.
type MPLSTESTDMIB_Mplstunnelcrldprestable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table represents a set of resources for an MPLS tunnel
    // established using CRLDP (mplsTunnelSignallingProto equal to crldp (3)). An
    // entry can be created by a network administrator or by an SNMP agent as
    // instructed by any MPLS signalling protocol. The type is slice of
    // MPLSTESTDMIB_Mplstunnelcrldprestable_Mplstunnelcrldpresentry.
    Mplstunnelcrldpresentry []MPLSTESTDMIB_Mplstunnelcrldprestable_Mplstunnelcrldpresentry
}

func (mplstunnelcrldprestable *MPLSTESTDMIB_Mplstunnelcrldprestable) GetEntityData() *types.CommonEntityData {
    mplstunnelcrldprestable.EntityData.YFilter = mplstunnelcrldprestable.YFilter
    mplstunnelcrldprestable.EntityData.YangName = "mplsTunnelCRLDPResTable"
    mplstunnelcrldprestable.EntityData.BundleName = "cisco_ios_xe"
    mplstunnelcrldprestable.EntityData.ParentYangName = "MPLS-TE-STD-MIB"
    mplstunnelcrldprestable.EntityData.SegmentPath = "mplsTunnelCRLDPResTable"
    mplstunnelcrldprestable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplstunnelcrldprestable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplstunnelcrldprestable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplstunnelcrldprestable.EntityData.Children = make(map[string]types.YChild)
    mplstunnelcrldprestable.EntityData.Children["mplsTunnelCRLDPResEntry"] = types.YChild{"Mplstunnelcrldpresentry", nil}
    for i := range mplstunnelcrldprestable.Mplstunnelcrldpresentry {
        mplstunnelcrldprestable.EntityData.Children[types.GetSegmentPath(&mplstunnelcrldprestable.Mplstunnelcrldpresentry[i])] = types.YChild{"Mplstunnelcrldpresentry", &mplstunnelcrldprestable.Mplstunnelcrldpresentry[i]}
    }
    mplstunnelcrldprestable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mplstunnelcrldprestable.EntityData)
}

// MPLSTESTDMIB_Mplstunnelcrldprestable_Mplstunnelcrldpresentry
// An entry in this table represents a set of resources
// for an MPLS tunnel established using CRLDP
// (mplsTunnelSignallingProto equal to crldp (3)). An
// entry can be created by a network administrator or
// by an SNMP agent as instructed by any MPLS
// signalling protocol.
type MPLSTESTDMIB_Mplstunnelcrldprestable_Mplstunnelcrldpresentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // mpls_te_std_mib.MPLSTESTDMIB_Mplstunnelresourcetable_Mplstunnelresourceentry_Mplstunnelresourceindex
    Mplstunnelresourceindex interface{}

    // The mean burst size in bytes. The type is interface{} with range:
    // 0..4294967295. Units are bytes.
    Mplstunnelcrldpresmeanburstsize interface{}

    // The Excess burst size in bytes. The type is interface{} with range:
    // 0..4294967295. Units are bytes.
    Mplstunnelcrldpresexburstsize interface{}

    // The granularity of the availability of committed rate. The type is
    // Mplstunnelcrldpresfrequency.
    Mplstunnelcrldpresfrequency interface{}

    // The relative weight for using excess bandwidth above its committed rate. 
    // The value of 0 means that weight is not applicable for the CR-LSP. The type
    // is interface{} with range: 0..255.
    Mplstunnelcrldpresweight interface{}

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
    Mplstunnelcrldpresflags interface{}

    // This variable is used to create, modify, and/or delete a row in this table.
    // When a row in this table is in active(1) state, no objects in that row can
    // be modified by the agent except mplsTunnelCRLDPResRowStatus and
    // mplsTunnelCRLDPResStorageType. The type is RowStatus.
    Mplstunnelcrldpresrowstatus interface{}

    // The storage type for this CR-LDP Resource entry. Conceptual rows having the
    // value 'permanent' need not allow write-access to any columnar objects in
    // the row. The type is StorageType.
    Mplstunnelcrldpresstoragetype interface{}
}

func (mplstunnelcrldpresentry *MPLSTESTDMIB_Mplstunnelcrldprestable_Mplstunnelcrldpresentry) GetEntityData() *types.CommonEntityData {
    mplstunnelcrldpresentry.EntityData.YFilter = mplstunnelcrldpresentry.YFilter
    mplstunnelcrldpresentry.EntityData.YangName = "mplsTunnelCRLDPResEntry"
    mplstunnelcrldpresentry.EntityData.BundleName = "cisco_ios_xe"
    mplstunnelcrldpresentry.EntityData.ParentYangName = "mplsTunnelCRLDPResTable"
    mplstunnelcrldpresentry.EntityData.SegmentPath = "mplsTunnelCRLDPResEntry" + "[mplsTunnelResourceIndex='" + fmt.Sprintf("%v", mplstunnelcrldpresentry.Mplstunnelresourceindex) + "']"
    mplstunnelcrldpresentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplstunnelcrldpresentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplstunnelcrldpresentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplstunnelcrldpresentry.EntityData.Children = make(map[string]types.YChild)
    mplstunnelcrldpresentry.EntityData.Leafs = make(map[string]types.YLeaf)
    mplstunnelcrldpresentry.EntityData.Leafs["mplsTunnelResourceIndex"] = types.YLeaf{"Mplstunnelresourceindex", mplstunnelcrldpresentry.Mplstunnelresourceindex}
    mplstunnelcrldpresentry.EntityData.Leafs["mplsTunnelCRLDPResMeanBurstSize"] = types.YLeaf{"Mplstunnelcrldpresmeanburstsize", mplstunnelcrldpresentry.Mplstunnelcrldpresmeanburstsize}
    mplstunnelcrldpresentry.EntityData.Leafs["mplsTunnelCRLDPResExBurstSize"] = types.YLeaf{"Mplstunnelcrldpresexburstsize", mplstunnelcrldpresentry.Mplstunnelcrldpresexburstsize}
    mplstunnelcrldpresentry.EntityData.Leafs["mplsTunnelCRLDPResFrequency"] = types.YLeaf{"Mplstunnelcrldpresfrequency", mplstunnelcrldpresentry.Mplstunnelcrldpresfrequency}
    mplstunnelcrldpresentry.EntityData.Leafs["mplsTunnelCRLDPResWeight"] = types.YLeaf{"Mplstunnelcrldpresweight", mplstunnelcrldpresentry.Mplstunnelcrldpresweight}
    mplstunnelcrldpresentry.EntityData.Leafs["mplsTunnelCRLDPResFlags"] = types.YLeaf{"Mplstunnelcrldpresflags", mplstunnelcrldpresentry.Mplstunnelcrldpresflags}
    mplstunnelcrldpresentry.EntityData.Leafs["mplsTunnelCRLDPResRowStatus"] = types.YLeaf{"Mplstunnelcrldpresrowstatus", mplstunnelcrldpresentry.Mplstunnelcrldpresrowstatus}
    mplstunnelcrldpresentry.EntityData.Leafs["mplsTunnelCRLDPResStorageType"] = types.YLeaf{"Mplstunnelcrldpresstoragetype", mplstunnelcrldpresentry.Mplstunnelcrldpresstoragetype}
    return &(mplstunnelcrldpresentry.EntityData)
}

// MPLSTESTDMIB_Mplstunnelcrldprestable_Mplstunnelcrldpresentry_Mplstunnelcrldpresfrequency represents rate.
type MPLSTESTDMIB_Mplstunnelcrldprestable_Mplstunnelcrldpresentry_Mplstunnelcrldpresfrequency string

const (
    MPLSTESTDMIB_Mplstunnelcrldprestable_Mplstunnelcrldpresentry_Mplstunnelcrldpresfrequency_unspecified MPLSTESTDMIB_Mplstunnelcrldprestable_Mplstunnelcrldpresentry_Mplstunnelcrldpresfrequency = "unspecified"

    MPLSTESTDMIB_Mplstunnelcrldprestable_Mplstunnelcrldpresentry_Mplstunnelcrldpresfrequency_frequent MPLSTESTDMIB_Mplstunnelcrldprestable_Mplstunnelcrldpresentry_Mplstunnelcrldpresfrequency = "frequent"

    MPLSTESTDMIB_Mplstunnelcrldprestable_Mplstunnelcrldpresentry_Mplstunnelcrldpresfrequency_veryFrequent MPLSTESTDMIB_Mplstunnelcrldprestable_Mplstunnelcrldpresentry_Mplstunnelcrldpresfrequency = "veryFrequent"
)

