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
    parent types.Entity
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

func (mPLSTESTDMIB *MPLSTESTDMIB) GetFilter() yfilter.YFilter { return mPLSTESTDMIB.YFilter }

func (mPLSTESTDMIB *MPLSTESTDMIB) SetFilter(yf yfilter.YFilter) { mPLSTESTDMIB.YFilter = yf }

func (mPLSTESTDMIB *MPLSTESTDMIB) GetGoName(yname string) string {
    if yname == "mplsTeScalars" { return "Mplstescalars" }
    if yname == "mplsTeObjects" { return "Mplsteobjects" }
    if yname == "mplsTunnelTable" { return "Mplstunneltable" }
    if yname == "mplsTunnelHopTable" { return "Mplstunnelhoptable" }
    if yname == "mplsTunnelResourceTable" { return "Mplstunnelresourcetable" }
    if yname == "mplsTunnelARHopTable" { return "Mplstunnelarhoptable" }
    if yname == "mplsTunnelCHopTable" { return "Mplstunnelchoptable" }
    if yname == "mplsTunnelCRLDPResTable" { return "Mplstunnelcrldprestable" }
    return ""
}

func (mPLSTESTDMIB *MPLSTESTDMIB) GetSegmentPath() string {
    return "MPLS-TE-STD-MIB:MPLS-TE-STD-MIB"
}

func (mPLSTESTDMIB *MPLSTESTDMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mplsTeScalars" {
        return &mPLSTESTDMIB.Mplstescalars
    }
    if childYangName == "mplsTeObjects" {
        return &mPLSTESTDMIB.Mplsteobjects
    }
    if childYangName == "mplsTunnelTable" {
        return &mPLSTESTDMIB.Mplstunneltable
    }
    if childYangName == "mplsTunnelHopTable" {
        return &mPLSTESTDMIB.Mplstunnelhoptable
    }
    if childYangName == "mplsTunnelResourceTable" {
        return &mPLSTESTDMIB.Mplstunnelresourcetable
    }
    if childYangName == "mplsTunnelARHopTable" {
        return &mPLSTESTDMIB.Mplstunnelarhoptable
    }
    if childYangName == "mplsTunnelCHopTable" {
        return &mPLSTESTDMIB.Mplstunnelchoptable
    }
    if childYangName == "mplsTunnelCRLDPResTable" {
        return &mPLSTESTDMIB.Mplstunnelcrldprestable
    }
    return nil
}

func (mPLSTESTDMIB *MPLSTESTDMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["mplsTeScalars"] = &mPLSTESTDMIB.Mplstescalars
    children["mplsTeObjects"] = &mPLSTESTDMIB.Mplsteobjects
    children["mplsTunnelTable"] = &mPLSTESTDMIB.Mplstunneltable
    children["mplsTunnelHopTable"] = &mPLSTESTDMIB.Mplstunnelhoptable
    children["mplsTunnelResourceTable"] = &mPLSTESTDMIB.Mplstunnelresourcetable
    children["mplsTunnelARHopTable"] = &mPLSTESTDMIB.Mplstunnelarhoptable
    children["mplsTunnelCHopTable"] = &mPLSTESTDMIB.Mplstunnelchoptable
    children["mplsTunnelCRLDPResTable"] = &mPLSTESTDMIB.Mplstunnelcrldprestable
    return children
}

func (mPLSTESTDMIB *MPLSTESTDMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mPLSTESTDMIB *MPLSTESTDMIB) GetBundleName() string { return "cisco_ios_xe" }

func (mPLSTESTDMIB *MPLSTESTDMIB) GetYangName() string { return "MPLS-TE-STD-MIB" }

func (mPLSTESTDMIB *MPLSTESTDMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mPLSTESTDMIB *MPLSTESTDMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mPLSTESTDMIB *MPLSTESTDMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mPLSTESTDMIB *MPLSTESTDMIB) SetParent(parent types.Entity) { mPLSTESTDMIB.parent = parent }

func (mPLSTESTDMIB *MPLSTESTDMIB) GetParent() types.Entity { return mPLSTESTDMIB.parent }

func (mPLSTESTDMIB *MPLSTESTDMIB) GetParentYangName() string { return "MPLS-TE-STD-MIB" }

// MPLSTESTDMIB_Mplstescalars
type MPLSTESTDMIB_Mplstescalars struct {
    parent types.Entity
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

func (mplstescalars *MPLSTESTDMIB_Mplstescalars) GetFilter() yfilter.YFilter { return mplstescalars.YFilter }

func (mplstescalars *MPLSTESTDMIB_Mplstescalars) SetFilter(yf yfilter.YFilter) { mplstescalars.YFilter = yf }

func (mplstescalars *MPLSTESTDMIB_Mplstescalars) GetGoName(yname string) string {
    if yname == "mplsTunnelConfigured" { return "Mplstunnelconfigured" }
    if yname == "mplsTunnelActive" { return "Mplstunnelactive" }
    if yname == "mplsTunnelTEDistProto" { return "Mplstunneltedistproto" }
    if yname == "mplsTunnelMaxHops" { return "Mplstunnelmaxhops" }
    if yname == "mplsTunnelNotificationMaxRate" { return "Mplstunnelnotificationmaxrate" }
    return ""
}

func (mplstescalars *MPLSTESTDMIB_Mplstescalars) GetSegmentPath() string {
    return "mplsTeScalars"
}

func (mplstescalars *MPLSTESTDMIB_Mplstescalars) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mplstescalars *MPLSTESTDMIB_Mplstescalars) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mplstescalars *MPLSTESTDMIB_Mplstescalars) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mplsTunnelConfigured"] = mplstescalars.Mplstunnelconfigured
    leafs["mplsTunnelActive"] = mplstescalars.Mplstunnelactive
    leafs["mplsTunnelTEDistProto"] = mplstescalars.Mplstunneltedistproto
    leafs["mplsTunnelMaxHops"] = mplstescalars.Mplstunnelmaxhops
    leafs["mplsTunnelNotificationMaxRate"] = mplstescalars.Mplstunnelnotificationmaxrate
    return leafs
}

func (mplstescalars *MPLSTESTDMIB_Mplstescalars) GetBundleName() string { return "cisco_ios_xe" }

func (mplstescalars *MPLSTESTDMIB_Mplstescalars) GetYangName() string { return "mplsTeScalars" }

func (mplstescalars *MPLSTESTDMIB_Mplstescalars) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplstescalars *MPLSTESTDMIB_Mplstescalars) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplstescalars *MPLSTESTDMIB_Mplstescalars) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplstescalars *MPLSTESTDMIB_Mplstescalars) SetParent(parent types.Entity) { mplstescalars.parent = parent }

func (mplstescalars *MPLSTESTDMIB_Mplstescalars) GetParent() types.Entity { return mplstescalars.parent }

func (mplstescalars *MPLSTESTDMIB_Mplstescalars) GetParentYangName() string { return "MPLS-TE-STD-MIB" }

// MPLSTESTDMIB_Mplsteobjects
type MPLSTESTDMIB_Mplsteobjects struct {
    parent types.Entity
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

func (mplsteobjects *MPLSTESTDMIB_Mplsteobjects) GetFilter() yfilter.YFilter { return mplsteobjects.YFilter }

func (mplsteobjects *MPLSTESTDMIB_Mplsteobjects) SetFilter(yf yfilter.YFilter) { mplsteobjects.YFilter = yf }

func (mplsteobjects *MPLSTESTDMIB_Mplsteobjects) GetGoName(yname string) string {
    if yname == "mplsTunnelIndexNext" { return "Mplstunnelindexnext" }
    if yname == "mplsTunnelHopListIndexNext" { return "Mplstunnelhoplistindexnext" }
    if yname == "mplsTunnelResourceIndexNext" { return "Mplstunnelresourceindexnext" }
    if yname == "mplsTunnelNotificationEnable" { return "Mplstunnelnotificationenable" }
    return ""
}

func (mplsteobjects *MPLSTESTDMIB_Mplsteobjects) GetSegmentPath() string {
    return "mplsTeObjects"
}

func (mplsteobjects *MPLSTESTDMIB_Mplsteobjects) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mplsteobjects *MPLSTESTDMIB_Mplsteobjects) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mplsteobjects *MPLSTESTDMIB_Mplsteobjects) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mplsTunnelIndexNext"] = mplsteobjects.Mplstunnelindexnext
    leafs["mplsTunnelHopListIndexNext"] = mplsteobjects.Mplstunnelhoplistindexnext
    leafs["mplsTunnelResourceIndexNext"] = mplsteobjects.Mplstunnelresourceindexnext
    leafs["mplsTunnelNotificationEnable"] = mplsteobjects.Mplstunnelnotificationenable
    return leafs
}

func (mplsteobjects *MPLSTESTDMIB_Mplsteobjects) GetBundleName() string { return "cisco_ios_xe" }

func (mplsteobjects *MPLSTESTDMIB_Mplsteobjects) GetYangName() string { return "mplsTeObjects" }

func (mplsteobjects *MPLSTESTDMIB_Mplsteobjects) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplsteobjects *MPLSTESTDMIB_Mplsteobjects) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplsteobjects *MPLSTESTDMIB_Mplsteobjects) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplsteobjects *MPLSTESTDMIB_Mplsteobjects) SetParent(parent types.Entity) { mplsteobjects.parent = parent }

func (mplsteobjects *MPLSTESTDMIB_Mplsteobjects) GetParent() types.Entity { return mplsteobjects.parent }

func (mplsteobjects *MPLSTESTDMIB_Mplsteobjects) GetParentYangName() string { return "MPLS-TE-STD-MIB" }

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
    parent types.Entity
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

func (mplstunneltable *MPLSTESTDMIB_Mplstunneltable) GetFilter() yfilter.YFilter { return mplstunneltable.YFilter }

func (mplstunneltable *MPLSTESTDMIB_Mplstunneltable) SetFilter(yf yfilter.YFilter) { mplstunneltable.YFilter = yf }

func (mplstunneltable *MPLSTESTDMIB_Mplstunneltable) GetGoName(yname string) string {
    if yname == "mplsTunnelEntry" { return "Mplstunnelentry" }
    return ""
}

func (mplstunneltable *MPLSTESTDMIB_Mplstunneltable) GetSegmentPath() string {
    return "mplsTunnelTable"
}

func (mplstunneltable *MPLSTESTDMIB_Mplstunneltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mplsTunnelEntry" {
        for _, c := range mplstunneltable.Mplstunnelentry {
            if mplstunneltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry{}
        mplstunneltable.Mplstunnelentry = append(mplstunneltable.Mplstunnelentry, child)
        return &mplstunneltable.Mplstunnelentry[len(mplstunneltable.Mplstunnelentry)-1]
    }
    return nil
}

func (mplstunneltable *MPLSTESTDMIB_Mplstunneltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mplstunneltable.Mplstunnelentry {
        children[mplstunneltable.Mplstunnelentry[i].GetSegmentPath()] = &mplstunneltable.Mplstunnelentry[i]
    }
    return children
}

func (mplstunneltable *MPLSTESTDMIB_Mplstunneltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mplstunneltable *MPLSTESTDMIB_Mplstunneltable) GetBundleName() string { return "cisco_ios_xe" }

func (mplstunneltable *MPLSTESTDMIB_Mplstunneltable) GetYangName() string { return "mplsTunnelTable" }

func (mplstunneltable *MPLSTESTDMIB_Mplstunneltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplstunneltable *MPLSTESTDMIB_Mplstunneltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplstunneltable *MPLSTESTDMIB_Mplstunneltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplstunneltable *MPLSTESTDMIB_Mplstunneltable) SetParent(parent types.Entity) { mplstunneltable.parent = parent }

func (mplstunneltable *MPLSTESTDMIB_Mplstunneltable) GetParent() types.Entity { return mplstunneltable.parent }

func (mplstunneltable *MPLSTESTDMIB_Mplstunneltable) GetParentYangName() string { return "MPLS-TE-STD-MIB" }

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
    parent types.Entity
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
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
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
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
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

func (mplstunnelentry *MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry) GetFilter() yfilter.YFilter { return mplstunnelentry.YFilter }

func (mplstunnelentry *MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry) SetFilter(yf yfilter.YFilter) { mplstunnelentry.YFilter = yf }

func (mplstunnelentry *MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry) GetGoName(yname string) string {
    if yname == "mplsTunnelIndex" { return "Mplstunnelindex" }
    if yname == "mplsTunnelInstance" { return "Mplstunnelinstance" }
    if yname == "mplsTunnelIngressLSRId" { return "Mplstunnelingresslsrid" }
    if yname == "mplsTunnelEgressLSRId" { return "Mplstunnelegresslsrid" }
    if yname == "mplsTunnelName" { return "Mplstunnelname" }
    if yname == "mplsTunnelDescr" { return "Mplstunneldescr" }
    if yname == "mplsTunnelIsIf" { return "Mplstunnelisif" }
    if yname == "mplsTunnelIfIndex" { return "Mplstunnelifindex" }
    if yname == "mplsTunnelOwner" { return "Mplstunnelowner" }
    if yname == "mplsTunnelRole" { return "Mplstunnelrole" }
    if yname == "mplsTunnelXCPointer" { return "Mplstunnelxcpointer" }
    if yname == "mplsTunnelSignallingProto" { return "Mplstunnelsignallingproto" }
    if yname == "mplsTunnelSetupPrio" { return "Mplstunnelsetupprio" }
    if yname == "mplsTunnelHoldingPrio" { return "Mplstunnelholdingprio" }
    if yname == "mplsTunnelSessionAttributes" { return "Mplstunnelsessionattributes" }
    if yname == "mplsTunnelLocalProtectInUse" { return "Mplstunnellocalprotectinuse" }
    if yname == "mplsTunnelResourcePointer" { return "Mplstunnelresourcepointer" }
    if yname == "mplsTunnelPrimaryInstance" { return "Mplstunnelprimaryinstance" }
    if yname == "mplsTunnelInstancePriority" { return "Mplstunnelinstancepriority" }
    if yname == "mplsTunnelHopTableIndex" { return "Mplstunnelhoptableindex" }
    if yname == "mplsTunnelPathInUse" { return "Mplstunnelpathinuse" }
    if yname == "mplsTunnelARHopTableIndex" { return "Mplstunnelarhoptableindex" }
    if yname == "mplsTunnelCHopTableIndex" { return "Mplstunnelchoptableindex" }
    if yname == "mplsTunnelIncludeAnyAffinity" { return "Mplstunnelincludeanyaffinity" }
    if yname == "mplsTunnelIncludeAllAffinity" { return "Mplstunnelincludeallaffinity" }
    if yname == "mplsTunnelExcludeAnyAffinity" { return "Mplstunnelexcludeanyaffinity" }
    if yname == "mplsTunnelTotalUpTime" { return "Mplstunneltotaluptime" }
    if yname == "mplsTunnelInstanceUpTime" { return "Mplstunnelinstanceuptime" }
    if yname == "mplsTunnelPrimaryUpTime" { return "Mplstunnelprimaryuptime" }
    if yname == "mplsTunnelPathChanges" { return "Mplstunnelpathchanges" }
    if yname == "mplsTunnelLastPathChange" { return "Mplstunnellastpathchange" }
    if yname == "mplsTunnelCreationTime" { return "Mplstunnelcreationtime" }
    if yname == "mplsTunnelStateTransitions" { return "Mplstunnelstatetransitions" }
    if yname == "mplsTunnelAdminStatus" { return "Mplstunneladminstatus" }
    if yname == "mplsTunnelOperStatus" { return "Mplstunneloperstatus" }
    if yname == "mplsTunnelRowStatus" { return "Mplstunnelrowstatus" }
    if yname == "mplsTunnelStorageType" { return "Mplstunnelstoragetype" }
    if yname == "mplsTunnelPerfPackets" { return "Mplstunnelperfpackets" }
    if yname == "mplsTunnelPerfHCPackets" { return "Mplstunnelperfhcpackets" }
    if yname == "mplsTunnelPerfErrors" { return "Mplstunnelperferrors" }
    if yname == "mplsTunnelPerfBytes" { return "Mplstunnelperfbytes" }
    if yname == "mplsTunnelPerfHCBytes" { return "Mplstunnelperfhcbytes" }
    return ""
}

func (mplstunnelentry *MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry) GetSegmentPath() string {
    return "mplsTunnelEntry" + "[mplsTunnelIndex='" + fmt.Sprintf("%v", mplstunnelentry.Mplstunnelindex) + "']" + "[mplsTunnelInstance='" + fmt.Sprintf("%v", mplstunnelentry.Mplstunnelinstance) + "']" + "[mplsTunnelIngressLSRId='" + fmt.Sprintf("%v", mplstunnelentry.Mplstunnelingresslsrid) + "']" + "[mplsTunnelEgressLSRId='" + fmt.Sprintf("%v", mplstunnelentry.Mplstunnelegresslsrid) + "']"
}

func (mplstunnelentry *MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mplstunnelentry *MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mplstunnelentry *MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mplsTunnelIndex"] = mplstunnelentry.Mplstunnelindex
    leafs["mplsTunnelInstance"] = mplstunnelentry.Mplstunnelinstance
    leafs["mplsTunnelIngressLSRId"] = mplstunnelentry.Mplstunnelingresslsrid
    leafs["mplsTunnelEgressLSRId"] = mplstunnelentry.Mplstunnelegresslsrid
    leafs["mplsTunnelName"] = mplstunnelentry.Mplstunnelname
    leafs["mplsTunnelDescr"] = mplstunnelentry.Mplstunneldescr
    leafs["mplsTunnelIsIf"] = mplstunnelentry.Mplstunnelisif
    leafs["mplsTunnelIfIndex"] = mplstunnelentry.Mplstunnelifindex
    leafs["mplsTunnelOwner"] = mplstunnelentry.Mplstunnelowner
    leafs["mplsTunnelRole"] = mplstunnelentry.Mplstunnelrole
    leafs["mplsTunnelXCPointer"] = mplstunnelentry.Mplstunnelxcpointer
    leafs["mplsTunnelSignallingProto"] = mplstunnelentry.Mplstunnelsignallingproto
    leafs["mplsTunnelSetupPrio"] = mplstunnelentry.Mplstunnelsetupprio
    leafs["mplsTunnelHoldingPrio"] = mplstunnelentry.Mplstunnelholdingprio
    leafs["mplsTunnelSessionAttributes"] = mplstunnelentry.Mplstunnelsessionattributes
    leafs["mplsTunnelLocalProtectInUse"] = mplstunnelentry.Mplstunnellocalprotectinuse
    leafs["mplsTunnelResourcePointer"] = mplstunnelentry.Mplstunnelresourcepointer
    leafs["mplsTunnelPrimaryInstance"] = mplstunnelentry.Mplstunnelprimaryinstance
    leafs["mplsTunnelInstancePriority"] = mplstunnelentry.Mplstunnelinstancepriority
    leafs["mplsTunnelHopTableIndex"] = mplstunnelentry.Mplstunnelhoptableindex
    leafs["mplsTunnelPathInUse"] = mplstunnelentry.Mplstunnelpathinuse
    leafs["mplsTunnelARHopTableIndex"] = mplstunnelentry.Mplstunnelarhoptableindex
    leafs["mplsTunnelCHopTableIndex"] = mplstunnelentry.Mplstunnelchoptableindex
    leafs["mplsTunnelIncludeAnyAffinity"] = mplstunnelentry.Mplstunnelincludeanyaffinity
    leafs["mplsTunnelIncludeAllAffinity"] = mplstunnelentry.Mplstunnelincludeallaffinity
    leafs["mplsTunnelExcludeAnyAffinity"] = mplstunnelentry.Mplstunnelexcludeanyaffinity
    leafs["mplsTunnelTotalUpTime"] = mplstunnelentry.Mplstunneltotaluptime
    leafs["mplsTunnelInstanceUpTime"] = mplstunnelentry.Mplstunnelinstanceuptime
    leafs["mplsTunnelPrimaryUpTime"] = mplstunnelentry.Mplstunnelprimaryuptime
    leafs["mplsTunnelPathChanges"] = mplstunnelentry.Mplstunnelpathchanges
    leafs["mplsTunnelLastPathChange"] = mplstunnelentry.Mplstunnellastpathchange
    leafs["mplsTunnelCreationTime"] = mplstunnelentry.Mplstunnelcreationtime
    leafs["mplsTunnelStateTransitions"] = mplstunnelentry.Mplstunnelstatetransitions
    leafs["mplsTunnelAdminStatus"] = mplstunnelentry.Mplstunneladminstatus
    leafs["mplsTunnelOperStatus"] = mplstunnelentry.Mplstunneloperstatus
    leafs["mplsTunnelRowStatus"] = mplstunnelentry.Mplstunnelrowstatus
    leafs["mplsTunnelStorageType"] = mplstunnelentry.Mplstunnelstoragetype
    leafs["mplsTunnelPerfPackets"] = mplstunnelentry.Mplstunnelperfpackets
    leafs["mplsTunnelPerfHCPackets"] = mplstunnelentry.Mplstunnelperfhcpackets
    leafs["mplsTunnelPerfErrors"] = mplstunnelentry.Mplstunnelperferrors
    leafs["mplsTunnelPerfBytes"] = mplstunnelentry.Mplstunnelperfbytes
    leafs["mplsTunnelPerfHCBytes"] = mplstunnelentry.Mplstunnelperfhcbytes
    return leafs
}

func (mplstunnelentry *MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry) GetBundleName() string { return "cisco_ios_xe" }

func (mplstunnelentry *MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry) GetYangName() string { return "mplsTunnelEntry" }

func (mplstunnelentry *MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplstunnelentry *MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplstunnelentry *MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplstunnelentry *MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry) SetParent(parent types.Entity) { mplstunnelentry.parent = parent }

func (mplstunnelentry *MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry) GetParent() types.Entity { return mplstunnelentry.parent }

func (mplstunnelentry *MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry) GetParentYangName() string { return "mplsTunnelTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in this table represents a tunnel hop.  An entry is created by a
    // network administrator for signaled ERLSP set up by an MPLS signalling
    // protocol. The type is slice of
    // MPLSTESTDMIB_Mplstunnelhoptable_Mplstunnelhopentry.
    Mplstunnelhopentry []MPLSTESTDMIB_Mplstunnelhoptable_Mplstunnelhopentry
}

func (mplstunnelhoptable *MPLSTESTDMIB_Mplstunnelhoptable) GetFilter() yfilter.YFilter { return mplstunnelhoptable.YFilter }

func (mplstunnelhoptable *MPLSTESTDMIB_Mplstunnelhoptable) SetFilter(yf yfilter.YFilter) { mplstunnelhoptable.YFilter = yf }

func (mplstunnelhoptable *MPLSTESTDMIB_Mplstunnelhoptable) GetGoName(yname string) string {
    if yname == "mplsTunnelHopEntry" { return "Mplstunnelhopentry" }
    return ""
}

func (mplstunnelhoptable *MPLSTESTDMIB_Mplstunnelhoptable) GetSegmentPath() string {
    return "mplsTunnelHopTable"
}

func (mplstunnelhoptable *MPLSTESTDMIB_Mplstunnelhoptable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mplsTunnelHopEntry" {
        for _, c := range mplstunnelhoptable.Mplstunnelhopentry {
            if mplstunnelhoptable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MPLSTESTDMIB_Mplstunnelhoptable_Mplstunnelhopentry{}
        mplstunnelhoptable.Mplstunnelhopentry = append(mplstunnelhoptable.Mplstunnelhopentry, child)
        return &mplstunnelhoptable.Mplstunnelhopentry[len(mplstunnelhoptable.Mplstunnelhopentry)-1]
    }
    return nil
}

func (mplstunnelhoptable *MPLSTESTDMIB_Mplstunnelhoptable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mplstunnelhoptable.Mplstunnelhopentry {
        children[mplstunnelhoptable.Mplstunnelhopentry[i].GetSegmentPath()] = &mplstunnelhoptable.Mplstunnelhopentry[i]
    }
    return children
}

func (mplstunnelhoptable *MPLSTESTDMIB_Mplstunnelhoptable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mplstunnelhoptable *MPLSTESTDMIB_Mplstunnelhoptable) GetBundleName() string { return "cisco_ios_xe" }

func (mplstunnelhoptable *MPLSTESTDMIB_Mplstunnelhoptable) GetYangName() string { return "mplsTunnelHopTable" }

func (mplstunnelhoptable *MPLSTESTDMIB_Mplstunnelhoptable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplstunnelhoptable *MPLSTESTDMIB_Mplstunnelhoptable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplstunnelhoptable *MPLSTESTDMIB_Mplstunnelhoptable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplstunnelhoptable *MPLSTESTDMIB_Mplstunnelhoptable) SetParent(parent types.Entity) { mplstunnelhoptable.parent = parent }

func (mplstunnelhoptable *MPLSTESTDMIB_Mplstunnelhoptable) GetParent() types.Entity { return mplstunnelhoptable.parent }

func (mplstunnelhoptable *MPLSTESTDMIB_Mplstunnelhoptable) GetParentYangName() string { return "MPLS-TE-STD-MIB" }

// MPLSTESTDMIB_Mplstunnelhoptable_Mplstunnelhopentry
// An entry in this table represents a tunnel hop.  An
// entry is created by a network administrator for
// signaled ERLSP set up by an MPLS signalling
// protocol.
type MPLSTESTDMIB_Mplstunnelhoptable_Mplstunnelhopentry struct {
    parent types.Entity
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

func (mplstunnelhopentry *MPLSTESTDMIB_Mplstunnelhoptable_Mplstunnelhopentry) GetFilter() yfilter.YFilter { return mplstunnelhopentry.YFilter }

func (mplstunnelhopentry *MPLSTESTDMIB_Mplstunnelhoptable_Mplstunnelhopentry) SetFilter(yf yfilter.YFilter) { mplstunnelhopentry.YFilter = yf }

func (mplstunnelhopentry *MPLSTESTDMIB_Mplstunnelhoptable_Mplstunnelhopentry) GetGoName(yname string) string {
    if yname == "mplsTunnelHopListIndex" { return "Mplstunnelhoplistindex" }
    if yname == "mplsTunnelHopPathOptionIndex" { return "Mplstunnelhoppathoptionindex" }
    if yname == "mplsTunnelHopIndex" { return "Mplstunnelhopindex" }
    if yname == "mplsTunnelHopAddrType" { return "Mplstunnelhopaddrtype" }
    if yname == "mplsTunnelHopIpAddr" { return "Mplstunnelhopipaddr" }
    if yname == "mplsTunnelHopIpPrefixLen" { return "Mplstunnelhopipprefixlen" }
    if yname == "mplsTunnelHopAsNumber" { return "Mplstunnelhopasnumber" }
    if yname == "mplsTunnelHopAddrUnnum" { return "Mplstunnelhopaddrunnum" }
    if yname == "mplsTunnelHopLspId" { return "Mplstunnelhoplspid" }
    if yname == "mplsTunnelHopType" { return "Mplstunnelhoptype" }
    if yname == "mplsTunnelHopInclude" { return "Mplstunnelhopinclude" }
    if yname == "mplsTunnelHopPathOptionName" { return "Mplstunnelhoppathoptionname" }
    if yname == "mplsTunnelHopEntryPathComp" { return "Mplstunnelhopentrypathcomp" }
    if yname == "mplsTunnelHopRowStatus" { return "Mplstunnelhoprowstatus" }
    if yname == "mplsTunnelHopStorageType" { return "Mplstunnelhopstoragetype" }
    return ""
}

func (mplstunnelhopentry *MPLSTESTDMIB_Mplstunnelhoptable_Mplstunnelhopentry) GetSegmentPath() string {
    return "mplsTunnelHopEntry" + "[mplsTunnelHopListIndex='" + fmt.Sprintf("%v", mplstunnelhopentry.Mplstunnelhoplistindex) + "']" + "[mplsTunnelHopPathOptionIndex='" + fmt.Sprintf("%v", mplstunnelhopentry.Mplstunnelhoppathoptionindex) + "']" + "[mplsTunnelHopIndex='" + fmt.Sprintf("%v", mplstunnelhopentry.Mplstunnelhopindex) + "']"
}

func (mplstunnelhopentry *MPLSTESTDMIB_Mplstunnelhoptable_Mplstunnelhopentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mplstunnelhopentry *MPLSTESTDMIB_Mplstunnelhoptable_Mplstunnelhopentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mplstunnelhopentry *MPLSTESTDMIB_Mplstunnelhoptable_Mplstunnelhopentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mplsTunnelHopListIndex"] = mplstunnelhopentry.Mplstunnelhoplistindex
    leafs["mplsTunnelHopPathOptionIndex"] = mplstunnelhopentry.Mplstunnelhoppathoptionindex
    leafs["mplsTunnelHopIndex"] = mplstunnelhopentry.Mplstunnelhopindex
    leafs["mplsTunnelHopAddrType"] = mplstunnelhopentry.Mplstunnelhopaddrtype
    leafs["mplsTunnelHopIpAddr"] = mplstunnelhopentry.Mplstunnelhopipaddr
    leafs["mplsTunnelHopIpPrefixLen"] = mplstunnelhopentry.Mplstunnelhopipprefixlen
    leafs["mplsTunnelHopAsNumber"] = mplstunnelhopentry.Mplstunnelhopasnumber
    leafs["mplsTunnelHopAddrUnnum"] = mplstunnelhopentry.Mplstunnelhopaddrunnum
    leafs["mplsTunnelHopLspId"] = mplstunnelhopentry.Mplstunnelhoplspid
    leafs["mplsTunnelHopType"] = mplstunnelhopentry.Mplstunnelhoptype
    leafs["mplsTunnelHopInclude"] = mplstunnelhopentry.Mplstunnelhopinclude
    leafs["mplsTunnelHopPathOptionName"] = mplstunnelhopentry.Mplstunnelhoppathoptionname
    leafs["mplsTunnelHopEntryPathComp"] = mplstunnelhopentry.Mplstunnelhopentrypathcomp
    leafs["mplsTunnelHopRowStatus"] = mplstunnelhopentry.Mplstunnelhoprowstatus
    leafs["mplsTunnelHopStorageType"] = mplstunnelhopentry.Mplstunnelhopstoragetype
    return leafs
}

func (mplstunnelhopentry *MPLSTESTDMIB_Mplstunnelhoptable_Mplstunnelhopentry) GetBundleName() string { return "cisco_ios_xe" }

func (mplstunnelhopentry *MPLSTESTDMIB_Mplstunnelhoptable_Mplstunnelhopentry) GetYangName() string { return "mplsTunnelHopEntry" }

func (mplstunnelhopentry *MPLSTESTDMIB_Mplstunnelhoptable_Mplstunnelhopentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplstunnelhopentry *MPLSTESTDMIB_Mplstunnelhoptable_Mplstunnelhopentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplstunnelhopentry *MPLSTESTDMIB_Mplstunnelhoptable_Mplstunnelhopentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplstunnelhopentry *MPLSTESTDMIB_Mplstunnelhoptable_Mplstunnelhopentry) SetParent(parent types.Entity) { mplstunnelhopentry.parent = parent }

func (mplstunnelhopentry *MPLSTESTDMIB_Mplstunnelhoptable_Mplstunnelhopentry) GetParent() types.Entity { return mplstunnelhopentry.parent }

func (mplstunnelhopentry *MPLSTESTDMIB_Mplstunnelhoptable_Mplstunnelhopentry) GetParentYangName() string { return "mplsTunnelHopTable" }

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
    parent types.Entity
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

func (mplstunnelresourcetable *MPLSTESTDMIB_Mplstunnelresourcetable) GetFilter() yfilter.YFilter { return mplstunnelresourcetable.YFilter }

func (mplstunnelresourcetable *MPLSTESTDMIB_Mplstunnelresourcetable) SetFilter(yf yfilter.YFilter) { mplstunnelresourcetable.YFilter = yf }

func (mplstunnelresourcetable *MPLSTESTDMIB_Mplstunnelresourcetable) GetGoName(yname string) string {
    if yname == "mplsTunnelResourceEntry" { return "Mplstunnelresourceentry" }
    return ""
}

func (mplstunnelresourcetable *MPLSTESTDMIB_Mplstunnelresourcetable) GetSegmentPath() string {
    return "mplsTunnelResourceTable"
}

func (mplstunnelresourcetable *MPLSTESTDMIB_Mplstunnelresourcetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mplsTunnelResourceEntry" {
        for _, c := range mplstunnelresourcetable.Mplstunnelresourceentry {
            if mplstunnelresourcetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MPLSTESTDMIB_Mplstunnelresourcetable_Mplstunnelresourceentry{}
        mplstunnelresourcetable.Mplstunnelresourceentry = append(mplstunnelresourcetable.Mplstunnelresourceentry, child)
        return &mplstunnelresourcetable.Mplstunnelresourceentry[len(mplstunnelresourcetable.Mplstunnelresourceentry)-1]
    }
    return nil
}

func (mplstunnelresourcetable *MPLSTESTDMIB_Mplstunnelresourcetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mplstunnelresourcetable.Mplstunnelresourceentry {
        children[mplstunnelresourcetable.Mplstunnelresourceentry[i].GetSegmentPath()] = &mplstunnelresourcetable.Mplstunnelresourceentry[i]
    }
    return children
}

func (mplstunnelresourcetable *MPLSTESTDMIB_Mplstunnelresourcetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mplstunnelresourcetable *MPLSTESTDMIB_Mplstunnelresourcetable) GetBundleName() string { return "cisco_ios_xe" }

func (mplstunnelresourcetable *MPLSTESTDMIB_Mplstunnelresourcetable) GetYangName() string { return "mplsTunnelResourceTable" }

func (mplstunnelresourcetable *MPLSTESTDMIB_Mplstunnelresourcetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplstunnelresourcetable *MPLSTESTDMIB_Mplstunnelresourcetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplstunnelresourcetable *MPLSTESTDMIB_Mplstunnelresourcetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplstunnelresourcetable *MPLSTESTDMIB_Mplstunnelresourcetable) SetParent(parent types.Entity) { mplstunnelresourcetable.parent = parent }

func (mplstunnelresourcetable *MPLSTESTDMIB_Mplstunnelresourcetable) GetParent() types.Entity { return mplstunnelresourcetable.parent }

func (mplstunnelresourcetable *MPLSTESTDMIB_Mplstunnelresourcetable) GetParentYangName() string { return "MPLS-TE-STD-MIB" }

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
    parent types.Entity
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

func (mplstunnelresourceentry *MPLSTESTDMIB_Mplstunnelresourcetable_Mplstunnelresourceentry) GetFilter() yfilter.YFilter { return mplstunnelresourceentry.YFilter }

func (mplstunnelresourceentry *MPLSTESTDMIB_Mplstunnelresourcetable_Mplstunnelresourceentry) SetFilter(yf yfilter.YFilter) { mplstunnelresourceentry.YFilter = yf }

func (mplstunnelresourceentry *MPLSTESTDMIB_Mplstunnelresourcetable_Mplstunnelresourceentry) GetGoName(yname string) string {
    if yname == "mplsTunnelResourceIndex" { return "Mplstunnelresourceindex" }
    if yname == "mplsTunnelResourceMaxRate" { return "Mplstunnelresourcemaxrate" }
    if yname == "mplsTunnelResourceMeanRate" { return "Mplstunnelresourcemeanrate" }
    if yname == "mplsTunnelResourceMaxBurstSize" { return "Mplstunnelresourcemaxburstsize" }
    if yname == "mplsTunnelResourceMeanBurstSize" { return "Mplstunnelresourcemeanburstsize" }
    if yname == "mplsTunnelResourceExBurstSize" { return "Mplstunnelresourceexburstsize" }
    if yname == "mplsTunnelResourceFrequency" { return "Mplstunnelresourcefrequency" }
    if yname == "mplsTunnelResourceWeight" { return "Mplstunnelresourceweight" }
    if yname == "mplsTunnelResourceRowStatus" { return "Mplstunnelresourcerowstatus" }
    if yname == "mplsTunnelResourceStorageType" { return "Mplstunnelresourcestoragetype" }
    return ""
}

func (mplstunnelresourceentry *MPLSTESTDMIB_Mplstunnelresourcetable_Mplstunnelresourceentry) GetSegmentPath() string {
    return "mplsTunnelResourceEntry" + "[mplsTunnelResourceIndex='" + fmt.Sprintf("%v", mplstunnelresourceentry.Mplstunnelresourceindex) + "']"
}

func (mplstunnelresourceentry *MPLSTESTDMIB_Mplstunnelresourcetable_Mplstunnelresourceentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mplstunnelresourceentry *MPLSTESTDMIB_Mplstunnelresourcetable_Mplstunnelresourceentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mplstunnelresourceentry *MPLSTESTDMIB_Mplstunnelresourcetable_Mplstunnelresourceentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mplsTunnelResourceIndex"] = mplstunnelresourceentry.Mplstunnelresourceindex
    leafs["mplsTunnelResourceMaxRate"] = mplstunnelresourceentry.Mplstunnelresourcemaxrate
    leafs["mplsTunnelResourceMeanRate"] = mplstunnelresourceentry.Mplstunnelresourcemeanrate
    leafs["mplsTunnelResourceMaxBurstSize"] = mplstunnelresourceentry.Mplstunnelresourcemaxburstsize
    leafs["mplsTunnelResourceMeanBurstSize"] = mplstunnelresourceentry.Mplstunnelresourcemeanburstsize
    leafs["mplsTunnelResourceExBurstSize"] = mplstunnelresourceentry.Mplstunnelresourceexburstsize
    leafs["mplsTunnelResourceFrequency"] = mplstunnelresourceentry.Mplstunnelresourcefrequency
    leafs["mplsTunnelResourceWeight"] = mplstunnelresourceentry.Mplstunnelresourceweight
    leafs["mplsTunnelResourceRowStatus"] = mplstunnelresourceentry.Mplstunnelresourcerowstatus
    leafs["mplsTunnelResourceStorageType"] = mplstunnelresourceentry.Mplstunnelresourcestoragetype
    return leafs
}

func (mplstunnelresourceentry *MPLSTESTDMIB_Mplstunnelresourcetable_Mplstunnelresourceentry) GetBundleName() string { return "cisco_ios_xe" }

func (mplstunnelresourceentry *MPLSTESTDMIB_Mplstunnelresourcetable_Mplstunnelresourceentry) GetYangName() string { return "mplsTunnelResourceEntry" }

func (mplstunnelresourceentry *MPLSTESTDMIB_Mplstunnelresourcetable_Mplstunnelresourceentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplstunnelresourceentry *MPLSTESTDMIB_Mplstunnelresourcetable_Mplstunnelresourceentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplstunnelresourceentry *MPLSTESTDMIB_Mplstunnelresourcetable_Mplstunnelresourceentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplstunnelresourceentry *MPLSTESTDMIB_Mplstunnelresourcetable_Mplstunnelresourceentry) SetParent(parent types.Entity) { mplstunnelresourceentry.parent = parent }

func (mplstunnelresourceentry *MPLSTESTDMIB_Mplstunnelresourcetable_Mplstunnelresourceentry) GetParent() types.Entity { return mplstunnelresourceentry.parent }

func (mplstunnelresourceentry *MPLSTESTDMIB_Mplstunnelresourcetable_Mplstunnelresourceentry) GetParentYangName() string { return "mplsTunnelResourceTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in this table represents a tunnel hop.  An entry is created by the
    // agent for signaled ERLSP set up by an MPLS signalling protocol. The type is
    // slice of MPLSTESTDMIB_Mplstunnelarhoptable_Mplstunnelarhopentry.
    Mplstunnelarhopentry []MPLSTESTDMIB_Mplstunnelarhoptable_Mplstunnelarhopentry
}

func (mplstunnelarhoptable *MPLSTESTDMIB_Mplstunnelarhoptable) GetFilter() yfilter.YFilter { return mplstunnelarhoptable.YFilter }

func (mplstunnelarhoptable *MPLSTESTDMIB_Mplstunnelarhoptable) SetFilter(yf yfilter.YFilter) { mplstunnelarhoptable.YFilter = yf }

func (mplstunnelarhoptable *MPLSTESTDMIB_Mplstunnelarhoptable) GetGoName(yname string) string {
    if yname == "mplsTunnelARHopEntry" { return "Mplstunnelarhopentry" }
    return ""
}

func (mplstunnelarhoptable *MPLSTESTDMIB_Mplstunnelarhoptable) GetSegmentPath() string {
    return "mplsTunnelARHopTable"
}

func (mplstunnelarhoptable *MPLSTESTDMIB_Mplstunnelarhoptable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mplsTunnelARHopEntry" {
        for _, c := range mplstunnelarhoptable.Mplstunnelarhopentry {
            if mplstunnelarhoptable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MPLSTESTDMIB_Mplstunnelarhoptable_Mplstunnelarhopentry{}
        mplstunnelarhoptable.Mplstunnelarhopentry = append(mplstunnelarhoptable.Mplstunnelarhopentry, child)
        return &mplstunnelarhoptable.Mplstunnelarhopentry[len(mplstunnelarhoptable.Mplstunnelarhopentry)-1]
    }
    return nil
}

func (mplstunnelarhoptable *MPLSTESTDMIB_Mplstunnelarhoptable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mplstunnelarhoptable.Mplstunnelarhopentry {
        children[mplstunnelarhoptable.Mplstunnelarhopentry[i].GetSegmentPath()] = &mplstunnelarhoptable.Mplstunnelarhopentry[i]
    }
    return children
}

func (mplstunnelarhoptable *MPLSTESTDMIB_Mplstunnelarhoptable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mplstunnelarhoptable *MPLSTESTDMIB_Mplstunnelarhoptable) GetBundleName() string { return "cisco_ios_xe" }

func (mplstunnelarhoptable *MPLSTESTDMIB_Mplstunnelarhoptable) GetYangName() string { return "mplsTunnelARHopTable" }

func (mplstunnelarhoptable *MPLSTESTDMIB_Mplstunnelarhoptable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplstunnelarhoptable *MPLSTESTDMIB_Mplstunnelarhoptable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplstunnelarhoptable *MPLSTESTDMIB_Mplstunnelarhoptable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplstunnelarhoptable *MPLSTESTDMIB_Mplstunnelarhoptable) SetParent(parent types.Entity) { mplstunnelarhoptable.parent = parent }

func (mplstunnelarhoptable *MPLSTESTDMIB_Mplstunnelarhoptable) GetParent() types.Entity { return mplstunnelarhoptable.parent }

func (mplstunnelarhoptable *MPLSTESTDMIB_Mplstunnelarhoptable) GetParentYangName() string { return "MPLS-TE-STD-MIB" }

// MPLSTESTDMIB_Mplstunnelarhoptable_Mplstunnelarhopentry
// An entry in this table represents a tunnel hop.  An
// entry is created by the agent for signaled ERLSP
// set up by an MPLS signalling protocol.
type MPLSTESTDMIB_Mplstunnelarhoptable_Mplstunnelarhopentry struct {
    parent types.Entity
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

func (mplstunnelarhopentry *MPLSTESTDMIB_Mplstunnelarhoptable_Mplstunnelarhopentry) GetFilter() yfilter.YFilter { return mplstunnelarhopentry.YFilter }

func (mplstunnelarhopentry *MPLSTESTDMIB_Mplstunnelarhoptable_Mplstunnelarhopentry) SetFilter(yf yfilter.YFilter) { mplstunnelarhopentry.YFilter = yf }

func (mplstunnelarhopentry *MPLSTESTDMIB_Mplstunnelarhoptable_Mplstunnelarhopentry) GetGoName(yname string) string {
    if yname == "mplsTunnelARHopListIndex" { return "Mplstunnelarhoplistindex" }
    if yname == "mplsTunnelARHopIndex" { return "Mplstunnelarhopindex" }
    if yname == "mplsTunnelARHopAddrType" { return "Mplstunnelarhopaddrtype" }
    if yname == "mplsTunnelARHopIpAddr" { return "Mplstunnelarhopipaddr" }
    if yname == "mplsTunnelARHopAddrUnnum" { return "Mplstunnelarhopaddrunnum" }
    if yname == "mplsTunnelARHopLspId" { return "Mplstunnelarhoplspid" }
    return ""
}

func (mplstunnelarhopentry *MPLSTESTDMIB_Mplstunnelarhoptable_Mplstunnelarhopentry) GetSegmentPath() string {
    return "mplsTunnelARHopEntry" + "[mplsTunnelARHopListIndex='" + fmt.Sprintf("%v", mplstunnelarhopentry.Mplstunnelarhoplistindex) + "']" + "[mplsTunnelARHopIndex='" + fmt.Sprintf("%v", mplstunnelarhopentry.Mplstunnelarhopindex) + "']"
}

func (mplstunnelarhopentry *MPLSTESTDMIB_Mplstunnelarhoptable_Mplstunnelarhopentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mplstunnelarhopentry *MPLSTESTDMIB_Mplstunnelarhoptable_Mplstunnelarhopentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mplstunnelarhopentry *MPLSTESTDMIB_Mplstunnelarhoptable_Mplstunnelarhopentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mplsTunnelARHopListIndex"] = mplstunnelarhopentry.Mplstunnelarhoplistindex
    leafs["mplsTunnelARHopIndex"] = mplstunnelarhopentry.Mplstunnelarhopindex
    leafs["mplsTunnelARHopAddrType"] = mplstunnelarhopentry.Mplstunnelarhopaddrtype
    leafs["mplsTunnelARHopIpAddr"] = mplstunnelarhopentry.Mplstunnelarhopipaddr
    leafs["mplsTunnelARHopAddrUnnum"] = mplstunnelarhopentry.Mplstunnelarhopaddrunnum
    leafs["mplsTunnelARHopLspId"] = mplstunnelarhopentry.Mplstunnelarhoplspid
    return leafs
}

func (mplstunnelarhopentry *MPLSTESTDMIB_Mplstunnelarhoptable_Mplstunnelarhopentry) GetBundleName() string { return "cisco_ios_xe" }

func (mplstunnelarhopentry *MPLSTESTDMIB_Mplstunnelarhoptable_Mplstunnelarhopentry) GetYangName() string { return "mplsTunnelARHopEntry" }

func (mplstunnelarhopentry *MPLSTESTDMIB_Mplstunnelarhoptable_Mplstunnelarhopentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplstunnelarhopentry *MPLSTESTDMIB_Mplstunnelarhoptable_Mplstunnelarhopentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplstunnelarhopentry *MPLSTESTDMIB_Mplstunnelarhoptable_Mplstunnelarhopentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplstunnelarhopentry *MPLSTESTDMIB_Mplstunnelarhoptable_Mplstunnelarhopentry) SetParent(parent types.Entity) { mplstunnelarhopentry.parent = parent }

func (mplstunnelarhopentry *MPLSTESTDMIB_Mplstunnelarhoptable_Mplstunnelarhopentry) GetParent() types.Entity { return mplstunnelarhopentry.parent }

func (mplstunnelarhopentry *MPLSTESTDMIB_Mplstunnelarhoptable_Mplstunnelarhopentry) GetParentYangName() string { return "mplsTunnelARHopTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in this table represents a tunnel hop.  An entry in this table is
    // created by a path computation engine using CSPF techniques applied to the
    // information collected by routing protocols and the hops specified in the
    // corresponding mplsTunnelHopTable. The type is slice of
    // MPLSTESTDMIB_Mplstunnelchoptable_Mplstunnelchopentry.
    Mplstunnelchopentry []MPLSTESTDMIB_Mplstunnelchoptable_Mplstunnelchopentry
}

func (mplstunnelchoptable *MPLSTESTDMIB_Mplstunnelchoptable) GetFilter() yfilter.YFilter { return mplstunnelchoptable.YFilter }

func (mplstunnelchoptable *MPLSTESTDMIB_Mplstunnelchoptable) SetFilter(yf yfilter.YFilter) { mplstunnelchoptable.YFilter = yf }

func (mplstunnelchoptable *MPLSTESTDMIB_Mplstunnelchoptable) GetGoName(yname string) string {
    if yname == "mplsTunnelCHopEntry" { return "Mplstunnelchopentry" }
    return ""
}

func (mplstunnelchoptable *MPLSTESTDMIB_Mplstunnelchoptable) GetSegmentPath() string {
    return "mplsTunnelCHopTable"
}

func (mplstunnelchoptable *MPLSTESTDMIB_Mplstunnelchoptable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mplsTunnelCHopEntry" {
        for _, c := range mplstunnelchoptable.Mplstunnelchopentry {
            if mplstunnelchoptable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MPLSTESTDMIB_Mplstunnelchoptable_Mplstunnelchopentry{}
        mplstunnelchoptable.Mplstunnelchopentry = append(mplstunnelchoptable.Mplstunnelchopentry, child)
        return &mplstunnelchoptable.Mplstunnelchopentry[len(mplstunnelchoptable.Mplstunnelchopentry)-1]
    }
    return nil
}

func (mplstunnelchoptable *MPLSTESTDMIB_Mplstunnelchoptable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mplstunnelchoptable.Mplstunnelchopentry {
        children[mplstunnelchoptable.Mplstunnelchopentry[i].GetSegmentPath()] = &mplstunnelchoptable.Mplstunnelchopentry[i]
    }
    return children
}

func (mplstunnelchoptable *MPLSTESTDMIB_Mplstunnelchoptable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mplstunnelchoptable *MPLSTESTDMIB_Mplstunnelchoptable) GetBundleName() string { return "cisco_ios_xe" }

func (mplstunnelchoptable *MPLSTESTDMIB_Mplstunnelchoptable) GetYangName() string { return "mplsTunnelCHopTable" }

func (mplstunnelchoptable *MPLSTESTDMIB_Mplstunnelchoptable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplstunnelchoptable *MPLSTESTDMIB_Mplstunnelchoptable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplstunnelchoptable *MPLSTESTDMIB_Mplstunnelchoptable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplstunnelchoptable *MPLSTESTDMIB_Mplstunnelchoptable) SetParent(parent types.Entity) { mplstunnelchoptable.parent = parent }

func (mplstunnelchoptable *MPLSTESTDMIB_Mplstunnelchoptable) GetParent() types.Entity { return mplstunnelchoptable.parent }

func (mplstunnelchoptable *MPLSTESTDMIB_Mplstunnelchoptable) GetParentYangName() string { return "MPLS-TE-STD-MIB" }

// MPLSTESTDMIB_Mplstunnelchoptable_Mplstunnelchopentry
// An entry in this table represents a tunnel hop.  An
// entry in this table is created by a path
// computation engine using CSPF techniques applied to
// the information collected by routing protocols and
// the hops specified in the corresponding
// mplsTunnelHopTable.
type MPLSTESTDMIB_Mplstunnelchoptable_Mplstunnelchopentry struct {
    parent types.Entity
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

func (mplstunnelchopentry *MPLSTESTDMIB_Mplstunnelchoptable_Mplstunnelchopentry) GetFilter() yfilter.YFilter { return mplstunnelchopentry.YFilter }

func (mplstunnelchopentry *MPLSTESTDMIB_Mplstunnelchoptable_Mplstunnelchopentry) SetFilter(yf yfilter.YFilter) { mplstunnelchopentry.YFilter = yf }

func (mplstunnelchopentry *MPLSTESTDMIB_Mplstunnelchoptable_Mplstunnelchopentry) GetGoName(yname string) string {
    if yname == "mplsTunnelCHopListIndex" { return "Mplstunnelchoplistindex" }
    if yname == "mplsTunnelCHopIndex" { return "Mplstunnelchopindex" }
    if yname == "mplsTunnelCHopAddrType" { return "Mplstunnelchopaddrtype" }
    if yname == "mplsTunnelCHopIpAddr" { return "Mplstunnelchopipaddr" }
    if yname == "mplsTunnelCHopIpPrefixLen" { return "Mplstunnelchopipprefixlen" }
    if yname == "mplsTunnelCHopAsNumber" { return "Mplstunnelchopasnumber" }
    if yname == "mplsTunnelCHopAddrUnnum" { return "Mplstunnelchopaddrunnum" }
    if yname == "mplsTunnelCHopLspId" { return "Mplstunnelchoplspid" }
    if yname == "mplsTunnelCHopType" { return "Mplstunnelchoptype" }
    return ""
}

func (mplstunnelchopentry *MPLSTESTDMIB_Mplstunnelchoptable_Mplstunnelchopentry) GetSegmentPath() string {
    return "mplsTunnelCHopEntry" + "[mplsTunnelCHopListIndex='" + fmt.Sprintf("%v", mplstunnelchopentry.Mplstunnelchoplistindex) + "']" + "[mplsTunnelCHopIndex='" + fmt.Sprintf("%v", mplstunnelchopentry.Mplstunnelchopindex) + "']"
}

func (mplstunnelchopentry *MPLSTESTDMIB_Mplstunnelchoptable_Mplstunnelchopentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mplstunnelchopentry *MPLSTESTDMIB_Mplstunnelchoptable_Mplstunnelchopentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mplstunnelchopentry *MPLSTESTDMIB_Mplstunnelchoptable_Mplstunnelchopentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mplsTunnelCHopListIndex"] = mplstunnelchopentry.Mplstunnelchoplistindex
    leafs["mplsTunnelCHopIndex"] = mplstunnelchopentry.Mplstunnelchopindex
    leafs["mplsTunnelCHopAddrType"] = mplstunnelchopentry.Mplstunnelchopaddrtype
    leafs["mplsTunnelCHopIpAddr"] = mplstunnelchopentry.Mplstunnelchopipaddr
    leafs["mplsTunnelCHopIpPrefixLen"] = mplstunnelchopentry.Mplstunnelchopipprefixlen
    leafs["mplsTunnelCHopAsNumber"] = mplstunnelchopentry.Mplstunnelchopasnumber
    leafs["mplsTunnelCHopAddrUnnum"] = mplstunnelchopentry.Mplstunnelchopaddrunnum
    leafs["mplsTunnelCHopLspId"] = mplstunnelchopentry.Mplstunnelchoplspid
    leafs["mplsTunnelCHopType"] = mplstunnelchopentry.Mplstunnelchoptype
    return leafs
}

func (mplstunnelchopentry *MPLSTESTDMIB_Mplstunnelchoptable_Mplstunnelchopentry) GetBundleName() string { return "cisco_ios_xe" }

func (mplstunnelchopentry *MPLSTESTDMIB_Mplstunnelchoptable_Mplstunnelchopentry) GetYangName() string { return "mplsTunnelCHopEntry" }

func (mplstunnelchopentry *MPLSTESTDMIB_Mplstunnelchoptable_Mplstunnelchopentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplstunnelchopentry *MPLSTESTDMIB_Mplstunnelchoptable_Mplstunnelchopentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplstunnelchopentry *MPLSTESTDMIB_Mplstunnelchoptable_Mplstunnelchopentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplstunnelchopentry *MPLSTESTDMIB_Mplstunnelchoptable_Mplstunnelchopentry) SetParent(parent types.Entity) { mplstunnelchopentry.parent = parent }

func (mplstunnelchopentry *MPLSTESTDMIB_Mplstunnelchoptable_Mplstunnelchopentry) GetParent() types.Entity { return mplstunnelchopentry.parent }

func (mplstunnelchopentry *MPLSTESTDMIB_Mplstunnelchoptable_Mplstunnelchopentry) GetParentYangName() string { return "mplsTunnelCHopTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in this table represents a set of resources for an MPLS tunnel
    // established using CRLDP (mplsTunnelSignallingProto equal to crldp (3)). An
    // entry can be created by a network administrator or by an SNMP agent as
    // instructed by any MPLS signalling protocol. The type is slice of
    // MPLSTESTDMIB_Mplstunnelcrldprestable_Mplstunnelcrldpresentry.
    Mplstunnelcrldpresentry []MPLSTESTDMIB_Mplstunnelcrldprestable_Mplstunnelcrldpresentry
}

func (mplstunnelcrldprestable *MPLSTESTDMIB_Mplstunnelcrldprestable) GetFilter() yfilter.YFilter { return mplstunnelcrldprestable.YFilter }

func (mplstunnelcrldprestable *MPLSTESTDMIB_Mplstunnelcrldprestable) SetFilter(yf yfilter.YFilter) { mplstunnelcrldprestable.YFilter = yf }

func (mplstunnelcrldprestable *MPLSTESTDMIB_Mplstunnelcrldprestable) GetGoName(yname string) string {
    if yname == "mplsTunnelCRLDPResEntry" { return "Mplstunnelcrldpresentry" }
    return ""
}

func (mplstunnelcrldprestable *MPLSTESTDMIB_Mplstunnelcrldprestable) GetSegmentPath() string {
    return "mplsTunnelCRLDPResTable"
}

func (mplstunnelcrldprestable *MPLSTESTDMIB_Mplstunnelcrldprestable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mplsTunnelCRLDPResEntry" {
        for _, c := range mplstunnelcrldprestable.Mplstunnelcrldpresentry {
            if mplstunnelcrldprestable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MPLSTESTDMIB_Mplstunnelcrldprestable_Mplstunnelcrldpresentry{}
        mplstunnelcrldprestable.Mplstunnelcrldpresentry = append(mplstunnelcrldprestable.Mplstunnelcrldpresentry, child)
        return &mplstunnelcrldprestable.Mplstunnelcrldpresentry[len(mplstunnelcrldprestable.Mplstunnelcrldpresentry)-1]
    }
    return nil
}

func (mplstunnelcrldprestable *MPLSTESTDMIB_Mplstunnelcrldprestable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mplstunnelcrldprestable.Mplstunnelcrldpresentry {
        children[mplstunnelcrldprestable.Mplstunnelcrldpresentry[i].GetSegmentPath()] = &mplstunnelcrldprestable.Mplstunnelcrldpresentry[i]
    }
    return children
}

func (mplstunnelcrldprestable *MPLSTESTDMIB_Mplstunnelcrldprestable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mplstunnelcrldprestable *MPLSTESTDMIB_Mplstunnelcrldprestable) GetBundleName() string { return "cisco_ios_xe" }

func (mplstunnelcrldprestable *MPLSTESTDMIB_Mplstunnelcrldprestable) GetYangName() string { return "mplsTunnelCRLDPResTable" }

func (mplstunnelcrldprestable *MPLSTESTDMIB_Mplstunnelcrldprestable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplstunnelcrldprestable *MPLSTESTDMIB_Mplstunnelcrldprestable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplstunnelcrldprestable *MPLSTESTDMIB_Mplstunnelcrldprestable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplstunnelcrldprestable *MPLSTESTDMIB_Mplstunnelcrldprestable) SetParent(parent types.Entity) { mplstunnelcrldprestable.parent = parent }

func (mplstunnelcrldprestable *MPLSTESTDMIB_Mplstunnelcrldprestable) GetParent() types.Entity { return mplstunnelcrldprestable.parent }

func (mplstunnelcrldprestable *MPLSTESTDMIB_Mplstunnelcrldprestable) GetParentYangName() string { return "MPLS-TE-STD-MIB" }

// MPLSTESTDMIB_Mplstunnelcrldprestable_Mplstunnelcrldpresentry
// An entry in this table represents a set of resources
// for an MPLS tunnel established using CRLDP
// (mplsTunnelSignallingProto equal to crldp (3)). An
// entry can be created by a network administrator or
// by an SNMP agent as instructed by any MPLS
// signalling protocol.
type MPLSTESTDMIB_Mplstunnelcrldprestable_Mplstunnelcrldpresentry struct {
    parent types.Entity
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

func (mplstunnelcrldpresentry *MPLSTESTDMIB_Mplstunnelcrldprestable_Mplstunnelcrldpresentry) GetFilter() yfilter.YFilter { return mplstunnelcrldpresentry.YFilter }

func (mplstunnelcrldpresentry *MPLSTESTDMIB_Mplstunnelcrldprestable_Mplstunnelcrldpresentry) SetFilter(yf yfilter.YFilter) { mplstunnelcrldpresentry.YFilter = yf }

func (mplstunnelcrldpresentry *MPLSTESTDMIB_Mplstunnelcrldprestable_Mplstunnelcrldpresentry) GetGoName(yname string) string {
    if yname == "mplsTunnelResourceIndex" { return "Mplstunnelresourceindex" }
    if yname == "mplsTunnelCRLDPResMeanBurstSize" { return "Mplstunnelcrldpresmeanburstsize" }
    if yname == "mplsTunnelCRLDPResExBurstSize" { return "Mplstunnelcrldpresexburstsize" }
    if yname == "mplsTunnelCRLDPResFrequency" { return "Mplstunnelcrldpresfrequency" }
    if yname == "mplsTunnelCRLDPResWeight" { return "Mplstunnelcrldpresweight" }
    if yname == "mplsTunnelCRLDPResFlags" { return "Mplstunnelcrldpresflags" }
    if yname == "mplsTunnelCRLDPResRowStatus" { return "Mplstunnelcrldpresrowstatus" }
    if yname == "mplsTunnelCRLDPResStorageType" { return "Mplstunnelcrldpresstoragetype" }
    return ""
}

func (mplstunnelcrldpresentry *MPLSTESTDMIB_Mplstunnelcrldprestable_Mplstunnelcrldpresentry) GetSegmentPath() string {
    return "mplsTunnelCRLDPResEntry" + "[mplsTunnelResourceIndex='" + fmt.Sprintf("%v", mplstunnelcrldpresentry.Mplstunnelresourceindex) + "']"
}

func (mplstunnelcrldpresentry *MPLSTESTDMIB_Mplstunnelcrldprestable_Mplstunnelcrldpresentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mplstunnelcrldpresentry *MPLSTESTDMIB_Mplstunnelcrldprestable_Mplstunnelcrldpresentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mplstunnelcrldpresentry *MPLSTESTDMIB_Mplstunnelcrldprestable_Mplstunnelcrldpresentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mplsTunnelResourceIndex"] = mplstunnelcrldpresentry.Mplstunnelresourceindex
    leafs["mplsTunnelCRLDPResMeanBurstSize"] = mplstunnelcrldpresentry.Mplstunnelcrldpresmeanburstsize
    leafs["mplsTunnelCRLDPResExBurstSize"] = mplstunnelcrldpresentry.Mplstunnelcrldpresexburstsize
    leafs["mplsTunnelCRLDPResFrequency"] = mplstunnelcrldpresentry.Mplstunnelcrldpresfrequency
    leafs["mplsTunnelCRLDPResWeight"] = mplstunnelcrldpresentry.Mplstunnelcrldpresweight
    leafs["mplsTunnelCRLDPResFlags"] = mplstunnelcrldpresentry.Mplstunnelcrldpresflags
    leafs["mplsTunnelCRLDPResRowStatus"] = mplstunnelcrldpresentry.Mplstunnelcrldpresrowstatus
    leafs["mplsTunnelCRLDPResStorageType"] = mplstunnelcrldpresentry.Mplstunnelcrldpresstoragetype
    return leafs
}

func (mplstunnelcrldpresentry *MPLSTESTDMIB_Mplstunnelcrldprestable_Mplstunnelcrldpresentry) GetBundleName() string { return "cisco_ios_xe" }

func (mplstunnelcrldpresentry *MPLSTESTDMIB_Mplstunnelcrldprestable_Mplstunnelcrldpresentry) GetYangName() string { return "mplsTunnelCRLDPResEntry" }

func (mplstunnelcrldpresentry *MPLSTESTDMIB_Mplstunnelcrldprestable_Mplstunnelcrldpresentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplstunnelcrldpresentry *MPLSTESTDMIB_Mplstunnelcrldprestable_Mplstunnelcrldpresentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplstunnelcrldpresentry *MPLSTESTDMIB_Mplstunnelcrldprestable_Mplstunnelcrldpresentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplstunnelcrldpresentry *MPLSTESTDMIB_Mplstunnelcrldprestable_Mplstunnelcrldpresentry) SetParent(parent types.Entity) { mplstunnelcrldpresentry.parent = parent }

func (mplstunnelcrldpresentry *MPLSTESTDMIB_Mplstunnelcrldprestable_Mplstunnelcrldpresentry) GetParent() types.Entity { return mplstunnelcrldpresentry.parent }

func (mplstunnelcrldpresentry *MPLSTESTDMIB_Mplstunnelcrldprestable_Mplstunnelcrldpresentry) GetParentYangName() string { return "mplsTunnelCRLDPResTable" }

// MPLSTESTDMIB_Mplstunnelcrldprestable_Mplstunnelcrldpresentry_Mplstunnelcrldpresfrequency represents rate.
type MPLSTESTDMIB_Mplstunnelcrldprestable_Mplstunnelcrldpresentry_Mplstunnelcrldpresfrequency string

const (
    MPLSTESTDMIB_Mplstunnelcrldprestable_Mplstunnelcrldpresentry_Mplstunnelcrldpresfrequency_unspecified MPLSTESTDMIB_Mplstunnelcrldprestable_Mplstunnelcrldpresentry_Mplstunnelcrldpresfrequency = "unspecified"

    MPLSTESTDMIB_Mplstunnelcrldprestable_Mplstunnelcrldpresentry_Mplstunnelcrldpresfrequency_frequent MPLSTESTDMIB_Mplstunnelcrldprestable_Mplstunnelcrldpresentry_Mplstunnelcrldpresfrequency = "frequent"

    MPLSTESTDMIB_Mplstunnelcrldprestable_Mplstunnelcrldpresentry_Mplstunnelcrldpresfrequency_veryFrequent MPLSTESTDMIB_Mplstunnelcrldprestable_Mplstunnelcrldpresentry_Mplstunnelcrldpresfrequency = "veryFrequent"
)

