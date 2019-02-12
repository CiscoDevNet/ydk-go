// This is the management MIB for devices complying to the  
// DOCSIS L2VPN Feature.
package docs_l2vpn_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package docs_l2vpn_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:DOCS-L2VPN-MIB DOCS-L2VPN-MIB}", reflect.TypeOf(DOCSL2VPNMIB{}))
    ydk.RegisterEntity("DOCS-L2VPN-MIB:DOCS-L2VPN-MIB", reflect.TypeOf(DOCSL2VPNMIB{}))
}

// DocsNsiEncapSubtype represents other than ieee802.1q(2).
type DocsNsiEncapSubtype string

const (
    DocsNsiEncapSubtype_other DocsNsiEncapSubtype = "other"

    DocsNsiEncapSubtype_ieee8021q DocsNsiEncapSubtype = "ieee8021q"

    DocsNsiEncapSubtype_ieee8021ad DocsNsiEncapSubtype = "ieee8021ad"

    DocsNsiEncapSubtype_mpls DocsNsiEncapSubtype = "mpls"

    DocsNsiEncapSubtype_l2tpv3 DocsNsiEncapSubtype = "l2tpv3"
)

// DOCSL2VPNMIB
type DOCSL2VPNMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table indexed by the octet string DocsL2vpnIdentifier that  provides the
    // local agent's internally assigned docsL2vpnIdx  value for that
    // DocsL2vpnIdentifier value. The mapping of   DocsL2vpnIdentifier to
    // docsL2vpnIdx is 1-1. The agent   must instantiate a row in both
    // docsL2vpnIndexToIdTable and  docsL2vpnIdToIndexTable for each known L2VPN
    // Identifier.
    DocsL2vpnIdToIndexTable DOCSL2VPNMIB_DocsL2vpnIdToIndexTable

    // Table indexed by agent's local docsL2vpnIdx that provides  the global L2VPN
    // Identifier. The mapping of docsL2vpnIdx to  DocsL2vpnIdentifier is 1-1. The
    // agent must instantiate a   row in both docsL2vpnIndexToIdTable and  
    // docsL2vpnIdToIndexTable for each known L2VPN.
    DocsL2vpnIndexToIdTable DOCSL2VPNMIB_DocsL2vpnIndexToIdTable

    // This table describes L2VPN per-CM information that  is in common with all
    // L2VPNs for the CM, regardless  of forwarding mode.
    DocsL2vpnCmTable DOCSL2VPNMIB_DocsL2vpnCmTable

    // This table describes the operation of L2VPN forwarding  on each CM.
    DocsL2vpnVpnCmTable DOCSL2VPNMIB_DocsL2vpnVpnCmTable

    // This table contains statistics for forwarding of   packets to and from a CM
    // on each VPN.
    DocsL2vpnVpnCmStatsTable DOCSL2VPNMIB_DocsL2vpnVpnCmStatsTable

    // This table displays summary information for the  run-time state of each VPN
    // that is currently operating   on each bridge port.
    DocsL2vpnPortStatusTable DOCSL2VPNMIB_DocsL2vpnPortStatusTable

    // This table displays SF-specific L2VPN forwarding status   for each upstream
    // service flow configured with a per-SF    L2VPN Encoding.    Objects which
    // were signaled in a per-SF L2VPN Encoding but   apply for the entire CM are
    // shown in the    docsL2vpnVpnCmTable.
    DocsL2vpnSfStatusTable DOCSL2VPNMIB_DocsL2vpnSfStatusTable

    // This table provides the L2VPN-specific objects for  packet classifiers that
    // apply to only L2VPN traffic.   The indices of this table are a subset of
    // the  indices of classifiers in docsQosPktClassTable.
    DocsL2vpnPktClassTable DOCSL2VPNMIB_DocsL2vpnPktClassTable

    // This table describes the NSI configuration for a single  CM when operating
    // in point-to-point forwarding mode for an  L2VPN.
    DocsL2vpnCmNsiTable DOCSL2VPNMIB_DocsL2vpnCmNsiTable

    // This table is a list of CPEs, indexed by the VPNs on a   Cable Modem.
    DocsL2vpnCmVpnCpeTable DOCSL2VPNMIB_DocsL2vpnCmVpnCpeTable

    // This table contains a list of bridging CPEs, indexed by  L2VPN Index and
    // the corresponding CMs on that VPN.
    DocsL2vpnVpnCmCpeTable DOCSL2VPNMIB_DocsL2vpnVpnCmCpeTable

    // This table contains packet counters for   Unicast MAC Addresses within a
    // VPN.
    DocsL2vpnDot1qTpFdbExtTable DOCSL2VPNMIB_DocsL2vpnDot1qTpFdbExtTable

    // This table contains packet counters for   Multicast MAC Addresses within a
    // VPN.
    DocsL2vpnDot1qTpGroupExtTable DOCSL2VPNMIB_DocsL2vpnDot1qTpGroupExtTable
}

func (dOCSL2VPNMIB *DOCSL2VPNMIB) GetEntityData() *types.CommonEntityData {
    dOCSL2VPNMIB.EntityData.YFilter = dOCSL2VPNMIB.YFilter
    dOCSL2VPNMIB.EntityData.YangName = "DOCS-L2VPN-MIB"
    dOCSL2VPNMIB.EntityData.BundleName = "cisco_ios_xe"
    dOCSL2VPNMIB.EntityData.ParentYangName = "DOCS-L2VPN-MIB"
    dOCSL2VPNMIB.EntityData.SegmentPath = "DOCS-L2VPN-MIB:DOCS-L2VPN-MIB"
    dOCSL2VPNMIB.EntityData.AbsolutePath = dOCSL2VPNMIB.EntityData.SegmentPath
    dOCSL2VPNMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dOCSL2VPNMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dOCSL2VPNMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dOCSL2VPNMIB.EntityData.Children = types.NewOrderedMap()
    dOCSL2VPNMIB.EntityData.Children.Append("docsL2vpnIdToIndexTable", types.YChild{"DocsL2vpnIdToIndexTable", &dOCSL2VPNMIB.DocsL2vpnIdToIndexTable})
    dOCSL2VPNMIB.EntityData.Children.Append("docsL2vpnIndexToIdTable", types.YChild{"DocsL2vpnIndexToIdTable", &dOCSL2VPNMIB.DocsL2vpnIndexToIdTable})
    dOCSL2VPNMIB.EntityData.Children.Append("docsL2vpnCmTable", types.YChild{"DocsL2vpnCmTable", &dOCSL2VPNMIB.DocsL2vpnCmTable})
    dOCSL2VPNMIB.EntityData.Children.Append("docsL2vpnVpnCmTable", types.YChild{"DocsL2vpnVpnCmTable", &dOCSL2VPNMIB.DocsL2vpnVpnCmTable})
    dOCSL2VPNMIB.EntityData.Children.Append("docsL2vpnVpnCmStatsTable", types.YChild{"DocsL2vpnVpnCmStatsTable", &dOCSL2VPNMIB.DocsL2vpnVpnCmStatsTable})
    dOCSL2VPNMIB.EntityData.Children.Append("docsL2vpnPortStatusTable", types.YChild{"DocsL2vpnPortStatusTable", &dOCSL2VPNMIB.DocsL2vpnPortStatusTable})
    dOCSL2VPNMIB.EntityData.Children.Append("docsL2vpnSfStatusTable", types.YChild{"DocsL2vpnSfStatusTable", &dOCSL2VPNMIB.DocsL2vpnSfStatusTable})
    dOCSL2VPNMIB.EntityData.Children.Append("docsL2vpnPktClassTable", types.YChild{"DocsL2vpnPktClassTable", &dOCSL2VPNMIB.DocsL2vpnPktClassTable})
    dOCSL2VPNMIB.EntityData.Children.Append("docsL2vpnCmNsiTable", types.YChild{"DocsL2vpnCmNsiTable", &dOCSL2VPNMIB.DocsL2vpnCmNsiTable})
    dOCSL2VPNMIB.EntityData.Children.Append("docsL2vpnCmVpnCpeTable", types.YChild{"DocsL2vpnCmVpnCpeTable", &dOCSL2VPNMIB.DocsL2vpnCmVpnCpeTable})
    dOCSL2VPNMIB.EntityData.Children.Append("docsL2vpnVpnCmCpeTable", types.YChild{"DocsL2vpnVpnCmCpeTable", &dOCSL2VPNMIB.DocsL2vpnVpnCmCpeTable})
    dOCSL2VPNMIB.EntityData.Children.Append("docsL2vpnDot1qTpFdbExtTable", types.YChild{"DocsL2vpnDot1qTpFdbExtTable", &dOCSL2VPNMIB.DocsL2vpnDot1qTpFdbExtTable})
    dOCSL2VPNMIB.EntityData.Children.Append("docsL2vpnDot1qTpGroupExtTable", types.YChild{"DocsL2vpnDot1qTpGroupExtTable", &dOCSL2VPNMIB.DocsL2vpnDot1qTpGroupExtTable})
    dOCSL2VPNMIB.EntityData.Leafs = types.NewOrderedMap()

    dOCSL2VPNMIB.EntityData.YListKeys = []string {}

    return &(dOCSL2VPNMIB.EntityData)
}

// DOCSL2VPNMIB_DocsL2vpnIdToIndexTable
// Table indexed by the octet string DocsL2vpnIdentifier that 
// provides the local agent's internally assigned docsL2vpnIdx 
// value for that DocsL2vpnIdentifier value. The mapping of  
// DocsL2vpnIdentifier to docsL2vpnIdx is 1-1. The agent  
// must instantiate a row in both docsL2vpnIndexToIdTable and 
// docsL2vpnIdToIndexTable for each known L2VPN Identifier.
type DOCSL2VPNMIB_DocsL2vpnIdToIndexTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maps a DocsL2vpnIdentifier octet string into the local   agent's locally
    // assigned docsL2vpnIdx value. The type is slice of
    // DOCSL2VPNMIB_DocsL2vpnIdToIndexTable_DocsL2vpnIdToIndexEntry.
    DocsL2vpnIdToIndexEntry []*DOCSL2VPNMIB_DocsL2vpnIdToIndexTable_DocsL2vpnIdToIndexEntry
}

func (docsL2vpnIdToIndexTable *DOCSL2VPNMIB_DocsL2vpnIdToIndexTable) GetEntityData() *types.CommonEntityData {
    docsL2vpnIdToIndexTable.EntityData.YFilter = docsL2vpnIdToIndexTable.YFilter
    docsL2vpnIdToIndexTable.EntityData.YangName = "docsL2vpnIdToIndexTable"
    docsL2vpnIdToIndexTable.EntityData.BundleName = "cisco_ios_xe"
    docsL2vpnIdToIndexTable.EntityData.ParentYangName = "DOCS-L2VPN-MIB"
    docsL2vpnIdToIndexTable.EntityData.SegmentPath = "docsL2vpnIdToIndexTable"
    docsL2vpnIdToIndexTable.EntityData.AbsolutePath = "DOCS-L2VPN-MIB:DOCS-L2VPN-MIB/" + docsL2vpnIdToIndexTable.EntityData.SegmentPath
    docsL2vpnIdToIndexTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsL2vpnIdToIndexTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsL2vpnIdToIndexTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsL2vpnIdToIndexTable.EntityData.Children = types.NewOrderedMap()
    docsL2vpnIdToIndexTable.EntityData.Children.Append("docsL2vpnIdToIndexEntry", types.YChild{"DocsL2vpnIdToIndexEntry", nil})
    for i := range docsL2vpnIdToIndexTable.DocsL2vpnIdToIndexEntry {
        docsL2vpnIdToIndexTable.EntityData.Children.Append(types.GetSegmentPath(docsL2vpnIdToIndexTable.DocsL2vpnIdToIndexEntry[i]), types.YChild{"DocsL2vpnIdToIndexEntry", docsL2vpnIdToIndexTable.DocsL2vpnIdToIndexEntry[i]})
    }
    docsL2vpnIdToIndexTable.EntityData.Leafs = types.NewOrderedMap()

    docsL2vpnIdToIndexTable.EntityData.YListKeys = []string {}

    return &(docsL2vpnIdToIndexTable.EntityData)
}

// DOCSL2VPNMIB_DocsL2vpnIdToIndexTable_DocsL2vpnIdToIndexEntry
// Maps a DocsL2vpnIdentifier octet string into the local  
// agent's locally assigned docsL2vpnIdx value.
type DOCSL2VPNMIB_DocsL2vpnIdToIndexTable_DocsL2vpnIdToIndexEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. An externally configured octet string that
    // identifies an  L2VPN. The type is string.
    DocsL2vpnId interface{}

    // An internally assigned index value for a known L2VPN. The type is
    // interface{} with range: 0..4294967295.
    DocsL2vpnIdToIndexIdx interface{}
}

func (docsL2vpnIdToIndexEntry *DOCSL2VPNMIB_DocsL2vpnIdToIndexTable_DocsL2vpnIdToIndexEntry) GetEntityData() *types.CommonEntityData {
    docsL2vpnIdToIndexEntry.EntityData.YFilter = docsL2vpnIdToIndexEntry.YFilter
    docsL2vpnIdToIndexEntry.EntityData.YangName = "docsL2vpnIdToIndexEntry"
    docsL2vpnIdToIndexEntry.EntityData.BundleName = "cisco_ios_xe"
    docsL2vpnIdToIndexEntry.EntityData.ParentYangName = "docsL2vpnIdToIndexTable"
    docsL2vpnIdToIndexEntry.EntityData.SegmentPath = "docsL2vpnIdToIndexEntry" + types.AddKeyToken(docsL2vpnIdToIndexEntry.DocsL2vpnId, "docsL2vpnId")
    docsL2vpnIdToIndexEntry.EntityData.AbsolutePath = "DOCS-L2VPN-MIB:DOCS-L2VPN-MIB/docsL2vpnIdToIndexTable/" + docsL2vpnIdToIndexEntry.EntityData.SegmentPath
    docsL2vpnIdToIndexEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsL2vpnIdToIndexEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsL2vpnIdToIndexEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsL2vpnIdToIndexEntry.EntityData.Children = types.NewOrderedMap()
    docsL2vpnIdToIndexEntry.EntityData.Leafs = types.NewOrderedMap()
    docsL2vpnIdToIndexEntry.EntityData.Leafs.Append("docsL2vpnId", types.YLeaf{"DocsL2vpnId", docsL2vpnIdToIndexEntry.DocsL2vpnId})
    docsL2vpnIdToIndexEntry.EntityData.Leafs.Append("docsL2vpnIdToIndexIdx", types.YLeaf{"DocsL2vpnIdToIndexIdx", docsL2vpnIdToIndexEntry.DocsL2vpnIdToIndexIdx})

    docsL2vpnIdToIndexEntry.EntityData.YListKeys = []string {"DocsL2vpnId"}

    return &(docsL2vpnIdToIndexEntry.EntityData)
}

// DOCSL2VPNMIB_DocsL2vpnIndexToIdTable
// Table indexed by agent's local docsL2vpnIdx that provides 
// the global L2VPN Identifier. The mapping of docsL2vpnIdx to 
// DocsL2vpnIdentifier is 1-1. The agent must instantiate a  
// row in both docsL2vpnIndexToIdTable and  
// docsL2vpnIdToIndexTable for each known L2VPN.
type DOCSL2VPNMIB_DocsL2vpnIndexToIdTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Provides the L2VPN Identifier for each locally-assigned   L2vpn Index. The
    // type is slice of
    // DOCSL2VPNMIB_DocsL2vpnIndexToIdTable_DocsL2vpnIndexToIdEntry.
    DocsL2vpnIndexToIdEntry []*DOCSL2VPNMIB_DocsL2vpnIndexToIdTable_DocsL2vpnIndexToIdEntry
}

func (docsL2vpnIndexToIdTable *DOCSL2VPNMIB_DocsL2vpnIndexToIdTable) GetEntityData() *types.CommonEntityData {
    docsL2vpnIndexToIdTable.EntityData.YFilter = docsL2vpnIndexToIdTable.YFilter
    docsL2vpnIndexToIdTable.EntityData.YangName = "docsL2vpnIndexToIdTable"
    docsL2vpnIndexToIdTable.EntityData.BundleName = "cisco_ios_xe"
    docsL2vpnIndexToIdTable.EntityData.ParentYangName = "DOCS-L2VPN-MIB"
    docsL2vpnIndexToIdTable.EntityData.SegmentPath = "docsL2vpnIndexToIdTable"
    docsL2vpnIndexToIdTable.EntityData.AbsolutePath = "DOCS-L2VPN-MIB:DOCS-L2VPN-MIB/" + docsL2vpnIndexToIdTable.EntityData.SegmentPath
    docsL2vpnIndexToIdTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsL2vpnIndexToIdTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsL2vpnIndexToIdTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsL2vpnIndexToIdTable.EntityData.Children = types.NewOrderedMap()
    docsL2vpnIndexToIdTable.EntityData.Children.Append("docsL2vpnIndexToIdEntry", types.YChild{"DocsL2vpnIndexToIdEntry", nil})
    for i := range docsL2vpnIndexToIdTable.DocsL2vpnIndexToIdEntry {
        docsL2vpnIndexToIdTable.EntityData.Children.Append(types.GetSegmentPath(docsL2vpnIndexToIdTable.DocsL2vpnIndexToIdEntry[i]), types.YChild{"DocsL2vpnIndexToIdEntry", docsL2vpnIndexToIdTable.DocsL2vpnIndexToIdEntry[i]})
    }
    docsL2vpnIndexToIdTable.EntityData.Leafs = types.NewOrderedMap()

    docsL2vpnIndexToIdTable.EntityData.YListKeys = []string {}

    return &(docsL2vpnIndexToIdTable.EntityData)
}

// DOCSL2VPNMIB_DocsL2vpnIndexToIdTable_DocsL2vpnIndexToIdEntry
// Provides the L2VPN Identifier for each locally-assigned  
// L2vpn Index.
type DOCSL2VPNMIB_DocsL2vpnIndexToIdTable_DocsL2vpnIndexToIdEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. An internally assigned index value for a known
    // L2VPN. The type is interface{} with range: 0..4294967295.
    DocsL2vpnIdx interface{}

    // An administered octet string that externally identifies an  L2VPN. The type
    // is string.
    DocsL2vpnIndexToIdId interface{}
}

func (docsL2vpnIndexToIdEntry *DOCSL2VPNMIB_DocsL2vpnIndexToIdTable_DocsL2vpnIndexToIdEntry) GetEntityData() *types.CommonEntityData {
    docsL2vpnIndexToIdEntry.EntityData.YFilter = docsL2vpnIndexToIdEntry.YFilter
    docsL2vpnIndexToIdEntry.EntityData.YangName = "docsL2vpnIndexToIdEntry"
    docsL2vpnIndexToIdEntry.EntityData.BundleName = "cisco_ios_xe"
    docsL2vpnIndexToIdEntry.EntityData.ParentYangName = "docsL2vpnIndexToIdTable"
    docsL2vpnIndexToIdEntry.EntityData.SegmentPath = "docsL2vpnIndexToIdEntry" + types.AddKeyToken(docsL2vpnIndexToIdEntry.DocsL2vpnIdx, "docsL2vpnIdx")
    docsL2vpnIndexToIdEntry.EntityData.AbsolutePath = "DOCS-L2VPN-MIB:DOCS-L2VPN-MIB/docsL2vpnIndexToIdTable/" + docsL2vpnIndexToIdEntry.EntityData.SegmentPath
    docsL2vpnIndexToIdEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsL2vpnIndexToIdEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsL2vpnIndexToIdEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsL2vpnIndexToIdEntry.EntityData.Children = types.NewOrderedMap()
    docsL2vpnIndexToIdEntry.EntityData.Leafs = types.NewOrderedMap()
    docsL2vpnIndexToIdEntry.EntityData.Leafs.Append("docsL2vpnIdx", types.YLeaf{"DocsL2vpnIdx", docsL2vpnIndexToIdEntry.DocsL2vpnIdx})
    docsL2vpnIndexToIdEntry.EntityData.Leafs.Append("docsL2vpnIndexToIdId", types.YLeaf{"DocsL2vpnIndexToIdId", docsL2vpnIndexToIdEntry.DocsL2vpnIndexToIdId})

    docsL2vpnIndexToIdEntry.EntityData.YListKeys = []string {"DocsL2vpnIdx"}

    return &(docsL2vpnIndexToIdEntry.EntityData)
}

// DOCSL2VPNMIB_DocsL2vpnCmTable
// This table describes L2VPN per-CM information that 
// is in common with all L2VPNs for the CM, regardless 
// of forwarding mode.
type DOCSL2VPNMIB_DocsL2vpnCmTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry is indexed by Cable Modem Index that  describes L2VPN information
    // for a single CM that is in  common with all L2VPNs implemented by the CM, 
    // regardless of the L2VPN forwarding mode.   An entry in this table is
    // created for every CM that   registers with a forwarding L2VPN encoding. The
    // type is slice of DOCSL2VPNMIB_DocsL2vpnCmTable_DocsL2vpnCmEntry.
    DocsL2vpnCmEntry []*DOCSL2VPNMIB_DocsL2vpnCmTable_DocsL2vpnCmEntry
}

func (docsL2vpnCmTable *DOCSL2VPNMIB_DocsL2vpnCmTable) GetEntityData() *types.CommonEntityData {
    docsL2vpnCmTable.EntityData.YFilter = docsL2vpnCmTable.YFilter
    docsL2vpnCmTable.EntityData.YangName = "docsL2vpnCmTable"
    docsL2vpnCmTable.EntityData.BundleName = "cisco_ios_xe"
    docsL2vpnCmTable.EntityData.ParentYangName = "DOCS-L2VPN-MIB"
    docsL2vpnCmTable.EntityData.SegmentPath = "docsL2vpnCmTable"
    docsL2vpnCmTable.EntityData.AbsolutePath = "DOCS-L2VPN-MIB:DOCS-L2VPN-MIB/" + docsL2vpnCmTable.EntityData.SegmentPath
    docsL2vpnCmTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsL2vpnCmTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsL2vpnCmTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsL2vpnCmTable.EntityData.Children = types.NewOrderedMap()
    docsL2vpnCmTable.EntityData.Children.Append("docsL2vpnCmEntry", types.YChild{"DocsL2vpnCmEntry", nil})
    for i := range docsL2vpnCmTable.DocsL2vpnCmEntry {
        docsL2vpnCmTable.EntityData.Children.Append(types.GetSegmentPath(docsL2vpnCmTable.DocsL2vpnCmEntry[i]), types.YChild{"DocsL2vpnCmEntry", docsL2vpnCmTable.DocsL2vpnCmEntry[i]})
    }
    docsL2vpnCmTable.EntityData.Leafs = types.NewOrderedMap()

    docsL2vpnCmTable.EntityData.YListKeys = []string {}

    return &(docsL2vpnCmTable.EntityData)
}

// DOCSL2VPNMIB_DocsL2vpnCmTable_DocsL2vpnCmEntry
// An entry is indexed by Cable Modem Index that 
// describes L2VPN information for a single CM that is in 
// common with all L2VPNs implemented by the CM, 
// regardless of the L2VPN forwarding mode. 
// 
// An entry in this table is created for every CM that  
// registers with a forwarding L2VPN encoding.
type DOCSL2VPNMIB_DocsL2vpnCmTable_DocsL2vpnCmEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // docs_if_mib.DOCSIFMIB_DocsIfCmtsCmStatusTable_DocsIfCmtsCmStatusEntry_DocsIfCmtsCmStatusIndex
    DocsIfCmtsCmStatusIndex interface{}

    // This object reports whether an L2VPN forwarding CM is  compliant with the
    // DOCSIS L2VPN specification, as reported  in the L2VPN Capability encoding
    // in the CM's registration  request message.   If the capability encoding was
    // omitted, this object must  report the value false(2). The type is bool.
    DocsL2vpnCmCompliantCapability interface{}

    // This object reports whether an L2VPN forwarding CM is  capable of
    // Downstream Unencrypted Traffic (DUT) Filtering,  as reported in the CM's
    // registration request message.   If the capability encoding was omitted,
    // this object must  report the value false(2). The type is bool.
    DocsL2vpnCmDutFilteringCapability interface{}

    // This object reports the value configured in a per-CM   L2VPN Encoding for
    // Downstream Unencrypted Traffic (DUT)  Cable Modem Interface Mask (CMIM).   
    // The DUT CMIM is a bit string with a '1' for each bit   position K for an
    // internal or external CM interface with   ifIndex K to which the CM permits
    // DUT to be forwarded. A CM  capable of DUT filtering MUST discard DUT to
    // interfaces  with a '0' in the DUT CMIM.    If a CM's top-level registration
    // request L2VPN Encoding  contained no DUT CMIM subtype, this object is
    // reported  with its default value of a '1' in bit position 0  
    // (corresponding to the eCM's own 'self' host) and a '1' in   each bit
    // position K for which an eSAFE interface exists at  ifIndex K. In other
    // words, the default DUT CMIM includes   the eCM and all eSAFE interfaces.  
    // This value is reported independently of whether the CM is  actually capable
    // of performing DUT filtering. The type is string.
    DocsL2vpnCmDutCMIM interface{}

    // This object reports the value of the Enable DHCP Snooping  subtype of a
    // top-level L2VPN Encoding.    It has the syntax of a CM Interface List
    // bitmask and  represents a set of CM MAC bridge interfaces   corresponding
    // to eSAFE hosts for which the CMTS is enabled   to snoop DHCP traffic in
    // order to learn the eSAFE host MAC  address on that interface.     Only bits
    // corresponding to eSAFE host MAC addresses may be  validly set in this
    // object, including cpe(1) for ePS  and the eSAFE interfaces in bits
    // positions 16 through 31. The type is map[string]bool.
    DocsL2vpnCmDhcpSnooping interface{}
}

func (docsL2vpnCmEntry *DOCSL2VPNMIB_DocsL2vpnCmTable_DocsL2vpnCmEntry) GetEntityData() *types.CommonEntityData {
    docsL2vpnCmEntry.EntityData.YFilter = docsL2vpnCmEntry.YFilter
    docsL2vpnCmEntry.EntityData.YangName = "docsL2vpnCmEntry"
    docsL2vpnCmEntry.EntityData.BundleName = "cisco_ios_xe"
    docsL2vpnCmEntry.EntityData.ParentYangName = "docsL2vpnCmTable"
    docsL2vpnCmEntry.EntityData.SegmentPath = "docsL2vpnCmEntry" + types.AddKeyToken(docsL2vpnCmEntry.DocsIfCmtsCmStatusIndex, "docsIfCmtsCmStatusIndex")
    docsL2vpnCmEntry.EntityData.AbsolutePath = "DOCS-L2VPN-MIB:DOCS-L2VPN-MIB/docsL2vpnCmTable/" + docsL2vpnCmEntry.EntityData.SegmentPath
    docsL2vpnCmEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsL2vpnCmEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsL2vpnCmEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsL2vpnCmEntry.EntityData.Children = types.NewOrderedMap()
    docsL2vpnCmEntry.EntityData.Leafs = types.NewOrderedMap()
    docsL2vpnCmEntry.EntityData.Leafs.Append("docsIfCmtsCmStatusIndex", types.YLeaf{"DocsIfCmtsCmStatusIndex", docsL2vpnCmEntry.DocsIfCmtsCmStatusIndex})
    docsL2vpnCmEntry.EntityData.Leafs.Append("docsL2vpnCmCompliantCapability", types.YLeaf{"DocsL2vpnCmCompliantCapability", docsL2vpnCmEntry.DocsL2vpnCmCompliantCapability})
    docsL2vpnCmEntry.EntityData.Leafs.Append("docsL2vpnCmDutFilteringCapability", types.YLeaf{"DocsL2vpnCmDutFilteringCapability", docsL2vpnCmEntry.DocsL2vpnCmDutFilteringCapability})
    docsL2vpnCmEntry.EntityData.Leafs.Append("docsL2vpnCmDutCMIM", types.YLeaf{"DocsL2vpnCmDutCMIM", docsL2vpnCmEntry.DocsL2vpnCmDutCMIM})
    docsL2vpnCmEntry.EntityData.Leafs.Append("docsL2vpnCmDhcpSnooping", types.YLeaf{"DocsL2vpnCmDhcpSnooping", docsL2vpnCmEntry.DocsL2vpnCmDhcpSnooping})

    docsL2vpnCmEntry.EntityData.YListKeys = []string {"DocsIfCmtsCmStatusIndex"}

    return &(docsL2vpnCmEntry.EntityData)
}

// DOCSL2VPNMIB_DocsL2vpnVpnCmTable
// This table describes the operation of L2VPN forwarding 
// on each CM.
type DOCSL2VPNMIB_DocsL2vpnVpnCmTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry is indexed by VPN ID and Cable Modem Index that  describes the
    // operation of L2VPN forwarding for a single  L2VPN on a single CM. The type
    // is slice of DOCSL2VPNMIB_DocsL2vpnVpnCmTable_DocsL2vpnVpnCmEntry.
    DocsL2vpnVpnCmEntry []*DOCSL2VPNMIB_DocsL2vpnVpnCmTable_DocsL2vpnVpnCmEntry
}

func (docsL2vpnVpnCmTable *DOCSL2VPNMIB_DocsL2vpnVpnCmTable) GetEntityData() *types.CommonEntityData {
    docsL2vpnVpnCmTable.EntityData.YFilter = docsL2vpnVpnCmTable.YFilter
    docsL2vpnVpnCmTable.EntityData.YangName = "docsL2vpnVpnCmTable"
    docsL2vpnVpnCmTable.EntityData.BundleName = "cisco_ios_xe"
    docsL2vpnVpnCmTable.EntityData.ParentYangName = "DOCS-L2VPN-MIB"
    docsL2vpnVpnCmTable.EntityData.SegmentPath = "docsL2vpnVpnCmTable"
    docsL2vpnVpnCmTable.EntityData.AbsolutePath = "DOCS-L2VPN-MIB:DOCS-L2VPN-MIB/" + docsL2vpnVpnCmTable.EntityData.SegmentPath
    docsL2vpnVpnCmTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsL2vpnVpnCmTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsL2vpnVpnCmTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsL2vpnVpnCmTable.EntityData.Children = types.NewOrderedMap()
    docsL2vpnVpnCmTable.EntityData.Children.Append("docsL2vpnVpnCmEntry", types.YChild{"DocsL2vpnVpnCmEntry", nil})
    for i := range docsL2vpnVpnCmTable.DocsL2vpnVpnCmEntry {
        docsL2vpnVpnCmTable.EntityData.Children.Append(types.GetSegmentPath(docsL2vpnVpnCmTable.DocsL2vpnVpnCmEntry[i]), types.YChild{"DocsL2vpnVpnCmEntry", docsL2vpnVpnCmTable.DocsL2vpnVpnCmEntry[i]})
    }
    docsL2vpnVpnCmTable.EntityData.Leafs = types.NewOrderedMap()

    docsL2vpnVpnCmTable.EntityData.YListKeys = []string {}

    return &(docsL2vpnVpnCmTable.EntityData)
}

// DOCSL2VPNMIB_DocsL2vpnVpnCmTable_DocsL2vpnVpnCmEntry
// An entry is indexed by VPN ID and Cable Modem Index that 
// describes the operation of L2VPN forwarding for a single 
// L2VPN on a single CM.
type DOCSL2VPNMIB_DocsL2vpnVpnCmTable_DocsL2vpnVpnCmEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // docs_l2vpn_mib.DOCSL2VPNMIB_DocsL2vpnIndexToIdTable_DocsL2vpnIndexToIdEntry_DocsL2vpnIdx
    DocsL2vpnIdx interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // docs_if_mib.DOCSIFMIB_DocsIfCmtsCmStatusTable_DocsIfCmtsCmStatusEntry_DocsIfCmtsCmStatusIndex
    DocsIfCmtsCmStatusIndex interface{}

    // A Cable Modem Interface Mask represents a set of   MAC bridge interfaces
    // within the CM. This object  represents the CMIM within a forwarding per-SF
    // L2VPN   encoding, which specifies a set of CM MAC bridge   interfaces to
    // which L2VPN forwarding is restricted.   If the CMIM Subtype is omitted from
    // a forwarding  per-SF encoding, its default value includes only 
    // cpePrimary(1) and cableMac(2), which can be encoded  with a single octet
    // with the value 0x60. The type is map[string]bool.
    DocsL2vpnVpnCmCMIM interface{}

    // The BPI+ Security Association ID in which traffic intended  for
    // point-to-point forwarding through an individual CM is   forwarded.    If
    // the CMTS does not allocate an individual SAID for  multipoint forwarding
    // (as is recommended),it MUST   report this object as zero. The type is
    // interface{} with range: 0..16383.
    DocsL2vpnVpnCmIndividualSAId interface{}

    // This object encodes the concatenation of all Vendor   Specific Subtype
    // encodings that appeared in any   registration per-CM L2VPN Encoding
    // associated with this   entry. The type is string.
    DocsL2vpnVpnCmVendorSpecific interface{}
}

func (docsL2vpnVpnCmEntry *DOCSL2VPNMIB_DocsL2vpnVpnCmTable_DocsL2vpnVpnCmEntry) GetEntityData() *types.CommonEntityData {
    docsL2vpnVpnCmEntry.EntityData.YFilter = docsL2vpnVpnCmEntry.YFilter
    docsL2vpnVpnCmEntry.EntityData.YangName = "docsL2vpnVpnCmEntry"
    docsL2vpnVpnCmEntry.EntityData.BundleName = "cisco_ios_xe"
    docsL2vpnVpnCmEntry.EntityData.ParentYangName = "docsL2vpnVpnCmTable"
    docsL2vpnVpnCmEntry.EntityData.SegmentPath = "docsL2vpnVpnCmEntry" + types.AddKeyToken(docsL2vpnVpnCmEntry.DocsL2vpnIdx, "docsL2vpnIdx") + types.AddKeyToken(docsL2vpnVpnCmEntry.DocsIfCmtsCmStatusIndex, "docsIfCmtsCmStatusIndex")
    docsL2vpnVpnCmEntry.EntityData.AbsolutePath = "DOCS-L2VPN-MIB:DOCS-L2VPN-MIB/docsL2vpnVpnCmTable/" + docsL2vpnVpnCmEntry.EntityData.SegmentPath
    docsL2vpnVpnCmEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsL2vpnVpnCmEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsL2vpnVpnCmEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsL2vpnVpnCmEntry.EntityData.Children = types.NewOrderedMap()
    docsL2vpnVpnCmEntry.EntityData.Leafs = types.NewOrderedMap()
    docsL2vpnVpnCmEntry.EntityData.Leafs.Append("docsL2vpnIdx", types.YLeaf{"DocsL2vpnIdx", docsL2vpnVpnCmEntry.DocsL2vpnIdx})
    docsL2vpnVpnCmEntry.EntityData.Leafs.Append("docsIfCmtsCmStatusIndex", types.YLeaf{"DocsIfCmtsCmStatusIndex", docsL2vpnVpnCmEntry.DocsIfCmtsCmStatusIndex})
    docsL2vpnVpnCmEntry.EntityData.Leafs.Append("docsL2vpnVpnCmCMIM", types.YLeaf{"DocsL2vpnVpnCmCMIM", docsL2vpnVpnCmEntry.DocsL2vpnVpnCmCMIM})
    docsL2vpnVpnCmEntry.EntityData.Leafs.Append("docsL2vpnVpnCmIndividualSAId", types.YLeaf{"DocsL2vpnVpnCmIndividualSAId", docsL2vpnVpnCmEntry.DocsL2vpnVpnCmIndividualSAId})
    docsL2vpnVpnCmEntry.EntityData.Leafs.Append("docsL2vpnVpnCmVendorSpecific", types.YLeaf{"DocsL2vpnVpnCmVendorSpecific", docsL2vpnVpnCmEntry.DocsL2vpnVpnCmVendorSpecific})

    docsL2vpnVpnCmEntry.EntityData.YListKeys = []string {"DocsL2vpnIdx", "DocsIfCmtsCmStatusIndex"}

    return &(docsL2vpnVpnCmEntry.EntityData)
}

// DOCSL2VPNMIB_DocsL2vpnVpnCmStatsTable
// This table contains statistics for forwarding of  
// packets to and from a CM on each VPN.
type DOCSL2VPNMIB_DocsL2vpnVpnCmStatsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry is indexed by VPN ID and Cable Modem Index. The type is slice of
    // DOCSL2VPNMIB_DocsL2vpnVpnCmStatsTable_DocsL2vpnVpnCmStatsEntry.
    DocsL2vpnVpnCmStatsEntry []*DOCSL2VPNMIB_DocsL2vpnVpnCmStatsTable_DocsL2vpnVpnCmStatsEntry
}

func (docsL2vpnVpnCmStatsTable *DOCSL2VPNMIB_DocsL2vpnVpnCmStatsTable) GetEntityData() *types.CommonEntityData {
    docsL2vpnVpnCmStatsTable.EntityData.YFilter = docsL2vpnVpnCmStatsTable.YFilter
    docsL2vpnVpnCmStatsTable.EntityData.YangName = "docsL2vpnVpnCmStatsTable"
    docsL2vpnVpnCmStatsTable.EntityData.BundleName = "cisco_ios_xe"
    docsL2vpnVpnCmStatsTable.EntityData.ParentYangName = "DOCS-L2VPN-MIB"
    docsL2vpnVpnCmStatsTable.EntityData.SegmentPath = "docsL2vpnVpnCmStatsTable"
    docsL2vpnVpnCmStatsTable.EntityData.AbsolutePath = "DOCS-L2VPN-MIB:DOCS-L2VPN-MIB/" + docsL2vpnVpnCmStatsTable.EntityData.SegmentPath
    docsL2vpnVpnCmStatsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsL2vpnVpnCmStatsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsL2vpnVpnCmStatsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsL2vpnVpnCmStatsTable.EntityData.Children = types.NewOrderedMap()
    docsL2vpnVpnCmStatsTable.EntityData.Children.Append("docsL2vpnVpnCmStatsEntry", types.YChild{"DocsL2vpnVpnCmStatsEntry", nil})
    for i := range docsL2vpnVpnCmStatsTable.DocsL2vpnVpnCmStatsEntry {
        docsL2vpnVpnCmStatsTable.EntityData.Children.Append(types.GetSegmentPath(docsL2vpnVpnCmStatsTable.DocsL2vpnVpnCmStatsEntry[i]), types.YChild{"DocsL2vpnVpnCmStatsEntry", docsL2vpnVpnCmStatsTable.DocsL2vpnVpnCmStatsEntry[i]})
    }
    docsL2vpnVpnCmStatsTable.EntityData.Leafs = types.NewOrderedMap()

    docsL2vpnVpnCmStatsTable.EntityData.YListKeys = []string {}

    return &(docsL2vpnVpnCmStatsTable.EntityData)
}

// DOCSL2VPNMIB_DocsL2vpnVpnCmStatsTable_DocsL2vpnVpnCmStatsEntry
// An entry is indexed by VPN ID and Cable Modem Index.
type DOCSL2VPNMIB_DocsL2vpnVpnCmStatsTable_DocsL2vpnVpnCmStatsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // docs_l2vpn_mib.DOCSL2VPNMIB_DocsL2vpnIndexToIdTable_DocsL2vpnIndexToIdEntry_DocsL2vpnIdx
    DocsL2vpnIdx interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // docs_if_mib.DOCSIFMIB_DocsIfCmtsCmStatusTable_DocsIfCmtsCmStatusEntry_DocsIfCmtsCmStatusIndex
    DocsIfCmtsCmStatusIndex interface{}

    // The number of L2vpn-forwarded packets received  from this instance's Cable
    // Modem on  this instance's L2VPN. The type is interface{} with range:
    // 0..4294967295.
    DocsL2vpnVpnCmStatsUpstreamPkts interface{}

    // The number of L2vpn-forwarded bytes received  from this instance's Cable
    // Modem on  this instance's L2VPN. The type is interface{} with range:
    // 0..4294967295.
    DocsL2vpnVpnCmStatsUpstreamBytes interface{}

    // The number of L2-forwarded packets   discarded from this instance's   Cable
    // Modem on this instance's VPN. The type is interface{} with range:
    // 0..4294967295.
    DocsL2vpnVpnCmStatsUpstreamDiscards interface{}

    // The number of L2-forwarded packets  transmitted to this instance's  Cable
    // Modem on this instance's VPN. The type is interface{} with range:
    // 0..4294967295.
    DocsL2vpnVpnCmStatsDownstreamPkts interface{}

    // The number of L2-forwarded bytes  transmitted to this instance's  Cable
    // Modem on this instance's VPN. The type is interface{} with range:
    // 0..4294967295.
    DocsL2vpnVpnCmStatsDownstreamBytes interface{}

    // The number of L2-forwarded packets that were discarded   before they could
    // be transmitted to this instance's   Cable Modem on this instance's VPN. The
    // type is interface{} with range: 0..4294967295.
    DocsL2vpnVpnCmStatsDownstreamDiscards interface{}
}

func (docsL2vpnVpnCmStatsEntry *DOCSL2VPNMIB_DocsL2vpnVpnCmStatsTable_DocsL2vpnVpnCmStatsEntry) GetEntityData() *types.CommonEntityData {
    docsL2vpnVpnCmStatsEntry.EntityData.YFilter = docsL2vpnVpnCmStatsEntry.YFilter
    docsL2vpnVpnCmStatsEntry.EntityData.YangName = "docsL2vpnVpnCmStatsEntry"
    docsL2vpnVpnCmStatsEntry.EntityData.BundleName = "cisco_ios_xe"
    docsL2vpnVpnCmStatsEntry.EntityData.ParentYangName = "docsL2vpnVpnCmStatsTable"
    docsL2vpnVpnCmStatsEntry.EntityData.SegmentPath = "docsL2vpnVpnCmStatsEntry" + types.AddKeyToken(docsL2vpnVpnCmStatsEntry.DocsL2vpnIdx, "docsL2vpnIdx") + types.AddKeyToken(docsL2vpnVpnCmStatsEntry.DocsIfCmtsCmStatusIndex, "docsIfCmtsCmStatusIndex")
    docsL2vpnVpnCmStatsEntry.EntityData.AbsolutePath = "DOCS-L2VPN-MIB:DOCS-L2VPN-MIB/docsL2vpnVpnCmStatsTable/" + docsL2vpnVpnCmStatsEntry.EntityData.SegmentPath
    docsL2vpnVpnCmStatsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsL2vpnVpnCmStatsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsL2vpnVpnCmStatsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsL2vpnVpnCmStatsEntry.EntityData.Children = types.NewOrderedMap()
    docsL2vpnVpnCmStatsEntry.EntityData.Leafs = types.NewOrderedMap()
    docsL2vpnVpnCmStatsEntry.EntityData.Leafs.Append("docsL2vpnIdx", types.YLeaf{"DocsL2vpnIdx", docsL2vpnVpnCmStatsEntry.DocsL2vpnIdx})
    docsL2vpnVpnCmStatsEntry.EntityData.Leafs.Append("docsIfCmtsCmStatusIndex", types.YLeaf{"DocsIfCmtsCmStatusIndex", docsL2vpnVpnCmStatsEntry.DocsIfCmtsCmStatusIndex})
    docsL2vpnVpnCmStatsEntry.EntityData.Leafs.Append("docsL2vpnVpnCmStatsUpstreamPkts", types.YLeaf{"DocsL2vpnVpnCmStatsUpstreamPkts", docsL2vpnVpnCmStatsEntry.DocsL2vpnVpnCmStatsUpstreamPkts})
    docsL2vpnVpnCmStatsEntry.EntityData.Leafs.Append("docsL2vpnVpnCmStatsUpstreamBytes", types.YLeaf{"DocsL2vpnVpnCmStatsUpstreamBytes", docsL2vpnVpnCmStatsEntry.DocsL2vpnVpnCmStatsUpstreamBytes})
    docsL2vpnVpnCmStatsEntry.EntityData.Leafs.Append("docsL2vpnVpnCmStatsUpstreamDiscards", types.YLeaf{"DocsL2vpnVpnCmStatsUpstreamDiscards", docsL2vpnVpnCmStatsEntry.DocsL2vpnVpnCmStatsUpstreamDiscards})
    docsL2vpnVpnCmStatsEntry.EntityData.Leafs.Append("docsL2vpnVpnCmStatsDownstreamPkts", types.YLeaf{"DocsL2vpnVpnCmStatsDownstreamPkts", docsL2vpnVpnCmStatsEntry.DocsL2vpnVpnCmStatsDownstreamPkts})
    docsL2vpnVpnCmStatsEntry.EntityData.Leafs.Append("docsL2vpnVpnCmStatsDownstreamBytes", types.YLeaf{"DocsL2vpnVpnCmStatsDownstreamBytes", docsL2vpnVpnCmStatsEntry.DocsL2vpnVpnCmStatsDownstreamBytes})
    docsL2vpnVpnCmStatsEntry.EntityData.Leafs.Append("docsL2vpnVpnCmStatsDownstreamDiscards", types.YLeaf{"DocsL2vpnVpnCmStatsDownstreamDiscards", docsL2vpnVpnCmStatsEntry.DocsL2vpnVpnCmStatsDownstreamDiscards})

    docsL2vpnVpnCmStatsEntry.EntityData.YListKeys = []string {"DocsL2vpnIdx", "DocsIfCmtsCmStatusIndex"}

    return &(docsL2vpnVpnCmStatsEntry.EntityData)
}

// DOCSL2VPNMIB_DocsL2vpnPortStatusTable
// This table displays summary information for the 
// run-time state of each VPN that is currently operating  
// on each bridge port.
type DOCSL2VPNMIB_DocsL2vpnPortStatusTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information specific to the operation of L2VPN forwarding  on a particular
    // CMTS 'bridge port'. A CMTS 'bridge port'   may be defined by the CMTS
    // vendor, but is advantageously a  single DOCSIS MAC Domain. The type is
    // slice of DOCSL2VPNMIB_DocsL2vpnPortStatusTable_DocsL2vpnPortStatusEntry.
    DocsL2vpnPortStatusEntry []*DOCSL2VPNMIB_DocsL2vpnPortStatusTable_DocsL2vpnPortStatusEntry
}

func (docsL2vpnPortStatusTable *DOCSL2VPNMIB_DocsL2vpnPortStatusTable) GetEntityData() *types.CommonEntityData {
    docsL2vpnPortStatusTable.EntityData.YFilter = docsL2vpnPortStatusTable.YFilter
    docsL2vpnPortStatusTable.EntityData.YangName = "docsL2vpnPortStatusTable"
    docsL2vpnPortStatusTable.EntityData.BundleName = "cisco_ios_xe"
    docsL2vpnPortStatusTable.EntityData.ParentYangName = "DOCS-L2VPN-MIB"
    docsL2vpnPortStatusTable.EntityData.SegmentPath = "docsL2vpnPortStatusTable"
    docsL2vpnPortStatusTable.EntityData.AbsolutePath = "DOCS-L2VPN-MIB:DOCS-L2VPN-MIB/" + docsL2vpnPortStatusTable.EntityData.SegmentPath
    docsL2vpnPortStatusTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsL2vpnPortStatusTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsL2vpnPortStatusTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsL2vpnPortStatusTable.EntityData.Children = types.NewOrderedMap()
    docsL2vpnPortStatusTable.EntityData.Children.Append("docsL2vpnPortStatusEntry", types.YChild{"DocsL2vpnPortStatusEntry", nil})
    for i := range docsL2vpnPortStatusTable.DocsL2vpnPortStatusEntry {
        docsL2vpnPortStatusTable.EntityData.Children.Append(types.GetSegmentPath(docsL2vpnPortStatusTable.DocsL2vpnPortStatusEntry[i]), types.YChild{"DocsL2vpnPortStatusEntry", docsL2vpnPortStatusTable.DocsL2vpnPortStatusEntry[i]})
    }
    docsL2vpnPortStatusTable.EntityData.Leafs = types.NewOrderedMap()

    docsL2vpnPortStatusTable.EntityData.YListKeys = []string {}

    return &(docsL2vpnPortStatusTable.EntityData)
}

// DOCSL2VPNMIB_DocsL2vpnPortStatusTable_DocsL2vpnPortStatusEntry
// Information specific to the operation of L2VPN forwarding 
// on a particular CMTS 'bridge port'. A CMTS 'bridge port'  
// may be defined by the CMTS vendor, but is advantageously a 
// single DOCSIS MAC Domain.
type DOCSL2VPNMIB_DocsL2vpnPortStatusTable_DocsL2vpnPortStatusEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..65535. Refers to
    // bridge_mib.BRIDGEMIB_Dot1dBasePortTable_Dot1dBasePortEntry_Dot1dBasePort
    Dot1dBasePort interface{}

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // docs_l2vpn_mib.DOCSL2VPNMIB_DocsL2vpnIndexToIdTable_DocsL2vpnIndexToIdEntry_DocsL2vpnIdx
    DocsL2vpnIdx interface{}

    // The Group SAID associated with this VPN on a   particular CMTS MAC domain.
    // This SAID is used to encrypt  all downstream flooded bridge traffic sent to
    // CMs on   this VPN and CMTS MAC domain bridge port.    A value of '0' means
    // there is no associated Group SAID for  this VPN and bridge port, e.g., if
    // the L2VPN uses  point-to-point individual SAIDs only for forwarding.   A
    // bridge port that is not a CMTS MAC  domain will report a value of '0'. The
    // type is interface{} with range: 0..16383.
    DocsL2vpnPortStatusGroupSAId interface{}
}

func (docsL2vpnPortStatusEntry *DOCSL2VPNMIB_DocsL2vpnPortStatusTable_DocsL2vpnPortStatusEntry) GetEntityData() *types.CommonEntityData {
    docsL2vpnPortStatusEntry.EntityData.YFilter = docsL2vpnPortStatusEntry.YFilter
    docsL2vpnPortStatusEntry.EntityData.YangName = "docsL2vpnPortStatusEntry"
    docsL2vpnPortStatusEntry.EntityData.BundleName = "cisco_ios_xe"
    docsL2vpnPortStatusEntry.EntityData.ParentYangName = "docsL2vpnPortStatusTable"
    docsL2vpnPortStatusEntry.EntityData.SegmentPath = "docsL2vpnPortStatusEntry" + types.AddKeyToken(docsL2vpnPortStatusEntry.Dot1dBasePort, "dot1dBasePort") + types.AddKeyToken(docsL2vpnPortStatusEntry.DocsL2vpnIdx, "docsL2vpnIdx")
    docsL2vpnPortStatusEntry.EntityData.AbsolutePath = "DOCS-L2VPN-MIB:DOCS-L2VPN-MIB/docsL2vpnPortStatusTable/" + docsL2vpnPortStatusEntry.EntityData.SegmentPath
    docsL2vpnPortStatusEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsL2vpnPortStatusEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsL2vpnPortStatusEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsL2vpnPortStatusEntry.EntityData.Children = types.NewOrderedMap()
    docsL2vpnPortStatusEntry.EntityData.Leafs = types.NewOrderedMap()
    docsL2vpnPortStatusEntry.EntityData.Leafs.Append("dot1dBasePort", types.YLeaf{"Dot1dBasePort", docsL2vpnPortStatusEntry.Dot1dBasePort})
    docsL2vpnPortStatusEntry.EntityData.Leafs.Append("docsL2vpnIdx", types.YLeaf{"DocsL2vpnIdx", docsL2vpnPortStatusEntry.DocsL2vpnIdx})
    docsL2vpnPortStatusEntry.EntityData.Leafs.Append("docsL2vpnPortStatusGroupSAId", types.YLeaf{"DocsL2vpnPortStatusGroupSAId", docsL2vpnPortStatusEntry.DocsL2vpnPortStatusGroupSAId})

    docsL2vpnPortStatusEntry.EntityData.YListKeys = []string {"Dot1dBasePort", "DocsL2vpnIdx"}

    return &(docsL2vpnPortStatusEntry.EntityData)
}

// DOCSL2VPNMIB_DocsL2vpnSfStatusTable
// This table displays SF-specific L2VPN forwarding status  
// for each upstream service flow configured with a per-SF  
//  L2VPN Encoding. 
// 
//  Objects which were signaled in a per-SF L2VPN Encoding but 
//  apply for the entire CM are shown in the  
//  docsL2vpnVpnCmTable.
type DOCSL2VPNMIB_DocsL2vpnSfStatusTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SF-specific L2VPN forwarding status information for each  upstream service
    // flow configured with a per-SF L2VPN   Encoding. The ifIndex is of type
    // docsCableMacLayer(127). The type is slice of
    // DOCSL2VPNMIB_DocsL2vpnSfStatusTable_DocsL2vpnSfStatusEntry.
    DocsL2vpnSfStatusEntry []*DOCSL2VPNMIB_DocsL2vpnSfStatusTable_DocsL2vpnSfStatusEntry
}

func (docsL2vpnSfStatusTable *DOCSL2VPNMIB_DocsL2vpnSfStatusTable) GetEntityData() *types.CommonEntityData {
    docsL2vpnSfStatusTable.EntityData.YFilter = docsL2vpnSfStatusTable.YFilter
    docsL2vpnSfStatusTable.EntityData.YangName = "docsL2vpnSfStatusTable"
    docsL2vpnSfStatusTable.EntityData.BundleName = "cisco_ios_xe"
    docsL2vpnSfStatusTable.EntityData.ParentYangName = "DOCS-L2VPN-MIB"
    docsL2vpnSfStatusTable.EntityData.SegmentPath = "docsL2vpnSfStatusTable"
    docsL2vpnSfStatusTable.EntityData.AbsolutePath = "DOCS-L2VPN-MIB:DOCS-L2VPN-MIB/" + docsL2vpnSfStatusTable.EntityData.SegmentPath
    docsL2vpnSfStatusTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsL2vpnSfStatusTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsL2vpnSfStatusTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsL2vpnSfStatusTable.EntityData.Children = types.NewOrderedMap()
    docsL2vpnSfStatusTable.EntityData.Children.Append("docsL2vpnSfStatusEntry", types.YChild{"DocsL2vpnSfStatusEntry", nil})
    for i := range docsL2vpnSfStatusTable.DocsL2vpnSfStatusEntry {
        docsL2vpnSfStatusTable.EntityData.Children.Append(types.GetSegmentPath(docsL2vpnSfStatusTable.DocsL2vpnSfStatusEntry[i]), types.YChild{"DocsL2vpnSfStatusEntry", docsL2vpnSfStatusTable.DocsL2vpnSfStatusEntry[i]})
    }
    docsL2vpnSfStatusTable.EntityData.Leafs = types.NewOrderedMap()

    docsL2vpnSfStatusTable.EntityData.YListKeys = []string {}

    return &(docsL2vpnSfStatusTable.EntityData)
}

// DOCSL2VPNMIB_DocsL2vpnSfStatusTable_DocsL2vpnSfStatusEntry
// SF-specific L2VPN forwarding status information for each 
// upstream service flow configured with a per-SF L2VPN  
// Encoding. The ifIndex is of type docsCableMacLayer(127).
type DOCSL2VPNMIB_DocsL2vpnSfStatusTable_DocsL2vpnSfStatusEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // docs_qos_mib.DOCSQOSMIB_DocsQosServiceFlowTable_DocsQosServiceFlowEntry_DocsQosServiceFlowId
    DocsQosServiceFlowId interface{}

    // This object represents the value of the L2VPN Identifier  subtype of a
    // per-SF L2VPN Encoding. The type is string.
    DocsL2vpnSfStatusL2vpnId interface{}

    // This object provides the configured Ingress User Priority  subtype of a
    // per-SF L2VPN Encoding for this CM. If the   subtype was omitted, this
    // object's value is zero. The type is interface{} with range: 0..7.
    DocsL2vpnSfStatusIngressUserPriority interface{}

    // This object provides the set of configured Vendor Specific  subtypes within
    // a per-SF L2VPN Encoding for a CM. If no   Vendor Specific subtype was
    // specified, this object is a   zero length octet string. If one or more
    // Vendor Specific   subtype parameters was specified, this object represents 
    // the concatenation of all such subtypes. The type is string.
    DocsL2vpnSfStatusVendorSpecific interface{}
}

func (docsL2vpnSfStatusEntry *DOCSL2VPNMIB_DocsL2vpnSfStatusTable_DocsL2vpnSfStatusEntry) GetEntityData() *types.CommonEntityData {
    docsL2vpnSfStatusEntry.EntityData.YFilter = docsL2vpnSfStatusEntry.YFilter
    docsL2vpnSfStatusEntry.EntityData.YangName = "docsL2vpnSfStatusEntry"
    docsL2vpnSfStatusEntry.EntityData.BundleName = "cisco_ios_xe"
    docsL2vpnSfStatusEntry.EntityData.ParentYangName = "docsL2vpnSfStatusTable"
    docsL2vpnSfStatusEntry.EntityData.SegmentPath = "docsL2vpnSfStatusEntry" + types.AddKeyToken(docsL2vpnSfStatusEntry.IfIndex, "ifIndex") + types.AddKeyToken(docsL2vpnSfStatusEntry.DocsQosServiceFlowId, "docsQosServiceFlowId")
    docsL2vpnSfStatusEntry.EntityData.AbsolutePath = "DOCS-L2VPN-MIB:DOCS-L2VPN-MIB/docsL2vpnSfStatusTable/" + docsL2vpnSfStatusEntry.EntityData.SegmentPath
    docsL2vpnSfStatusEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsL2vpnSfStatusEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsL2vpnSfStatusEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsL2vpnSfStatusEntry.EntityData.Children = types.NewOrderedMap()
    docsL2vpnSfStatusEntry.EntityData.Leafs = types.NewOrderedMap()
    docsL2vpnSfStatusEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", docsL2vpnSfStatusEntry.IfIndex})
    docsL2vpnSfStatusEntry.EntityData.Leafs.Append("docsQosServiceFlowId", types.YLeaf{"DocsQosServiceFlowId", docsL2vpnSfStatusEntry.DocsQosServiceFlowId})
    docsL2vpnSfStatusEntry.EntityData.Leafs.Append("docsL2vpnSfStatusL2vpnId", types.YLeaf{"DocsL2vpnSfStatusL2vpnId", docsL2vpnSfStatusEntry.DocsL2vpnSfStatusL2vpnId})
    docsL2vpnSfStatusEntry.EntityData.Leafs.Append("docsL2vpnSfStatusIngressUserPriority", types.YLeaf{"DocsL2vpnSfStatusIngressUserPriority", docsL2vpnSfStatusEntry.DocsL2vpnSfStatusIngressUserPriority})
    docsL2vpnSfStatusEntry.EntityData.Leafs.Append("docsL2vpnSfStatusVendorSpecific", types.YLeaf{"DocsL2vpnSfStatusVendorSpecific", docsL2vpnSfStatusEntry.DocsL2vpnSfStatusVendorSpecific})

    docsL2vpnSfStatusEntry.EntityData.YListKeys = []string {"IfIndex", "DocsQosServiceFlowId"}

    return &(docsL2vpnSfStatusEntry.EntityData)
}

// DOCSL2VPNMIB_DocsL2vpnPktClassTable
// This table provides the L2VPN-specific objects for 
// packet classifiers that apply to only L2VPN traffic.  
// The indices of this table are a subset of the 
// indices of classifiers in docsQosPktClassTable.
type DOCSL2VPNMIB_DocsL2vpnPktClassTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table extends a single row  of docsQosPktClassTable for a
    // rule that applies only to  downstream L2VPN forwarded packets.  The index
    // ifIndex is an ifType of docsCableMaclayer(127). The type is slice of
    // DOCSL2VPNMIB_DocsL2vpnPktClassTable_DocsL2vpnPktClassEntry.
    DocsL2vpnPktClassEntry []*DOCSL2VPNMIB_DocsL2vpnPktClassTable_DocsL2vpnPktClassEntry
}

func (docsL2vpnPktClassTable *DOCSL2VPNMIB_DocsL2vpnPktClassTable) GetEntityData() *types.CommonEntityData {
    docsL2vpnPktClassTable.EntityData.YFilter = docsL2vpnPktClassTable.YFilter
    docsL2vpnPktClassTable.EntityData.YangName = "docsL2vpnPktClassTable"
    docsL2vpnPktClassTable.EntityData.BundleName = "cisco_ios_xe"
    docsL2vpnPktClassTable.EntityData.ParentYangName = "DOCS-L2VPN-MIB"
    docsL2vpnPktClassTable.EntityData.SegmentPath = "docsL2vpnPktClassTable"
    docsL2vpnPktClassTable.EntityData.AbsolutePath = "DOCS-L2VPN-MIB:DOCS-L2VPN-MIB/" + docsL2vpnPktClassTable.EntityData.SegmentPath
    docsL2vpnPktClassTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsL2vpnPktClassTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsL2vpnPktClassTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsL2vpnPktClassTable.EntityData.Children = types.NewOrderedMap()
    docsL2vpnPktClassTable.EntityData.Children.Append("docsL2vpnPktClassEntry", types.YChild{"DocsL2vpnPktClassEntry", nil})
    for i := range docsL2vpnPktClassTable.DocsL2vpnPktClassEntry {
        docsL2vpnPktClassTable.EntityData.Children.Append(types.GetSegmentPath(docsL2vpnPktClassTable.DocsL2vpnPktClassEntry[i]), types.YChild{"DocsL2vpnPktClassEntry", docsL2vpnPktClassTable.DocsL2vpnPktClassEntry[i]})
    }
    docsL2vpnPktClassTable.EntityData.Leafs = types.NewOrderedMap()

    docsL2vpnPktClassTable.EntityData.YListKeys = []string {}

    return &(docsL2vpnPktClassTable.EntityData)
}

// DOCSL2VPNMIB_DocsL2vpnPktClassTable_DocsL2vpnPktClassEntry
// An entry in this table extends a single row 
// of docsQosPktClassTable for a rule that applies only to 
// downstream L2VPN forwarded packets. 
// The index ifIndex is an ifType of docsCableMaclayer(127).
type DOCSL2VPNMIB_DocsL2vpnPktClassTable_DocsL2vpnPktClassEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // docs_qos_mib.DOCSQOSMIB_DocsQosServiceFlowTable_DocsQosServiceFlowEntry_DocsQosServiceFlowId
    DocsQosServiceFlowId interface{}

    // This attribute is a key. The type is string with range: 1..65535. Refers to
    // docs_qos_mib.DOCSQOSMIB_DocsQosPktClassTable_DocsQosPktClassEntry_DocsQosPktClassId
    DocsQosPktClassId interface{}

    // The locally assigned L2VPN index corresponding to the VPN  Identifier
    // subtype of a Downstream Classifier L2VPN   Encoding. The type is string.
    DocsL2vpnPktClassL2vpnId interface{}

    // The lower priority of the user Priority Range subtype  of a Downstream
    // Classifier L2VPN Encoding. If the subtype  was omitted, this object has
    // value 0. The type is interface{} with range: 0..7.
    DocsL2vpnPktClassUserPriRangeLow interface{}

    // The higher priority of the user Priority Range subtype  of a Downstream
    // Classifier L2VPN Encoding. If the subtype  was omitted, this object has
    // value 7. The type is interface{} with range: 0..7.
    DocsL2vpnPktClassUserPriRangeHigh interface{}

    // The Cable Modem Interface Mask (CMIM) signaled in a   Packet Classifier
    // Encoding.  In a Downstream Packet   Classifier Encoding, a specified CMIM
    // value restricts the   classifier to match packets with a Destination MAC
    // address  corresponding to the interfaces indicated in the CMIM mask.  The
    // eCM self and any eSAFE interface bits correspond to  the individual eCM and
    // eSAFE host MAC addresses.   In an Upstream Packet Classifier encoding, a
    // specified CMIM  value restricts the classifier to match packets with an  
    // ingress bridge port interface matching the bits in the   CMIM value.   If
    // the CMIM subtype was omitted, this object should be   reported as a zero
    // length octet string. The type is map[string]bool.
    DocsL2vpnPktClassCMIM interface{}

    // This object provides the set of configured   Vendor Specific subtypes
    // within a Packet Classifier   Encoding for a CM. If no Vendor Specific
    // subtype was   specified, this object is a zero length octet string.   If
    // one or more Vendor Specific subtype parameters was   specified, this object
    // represents the concatenation of all  such subtypes. The type is string.
    DocsL2vpnPktClassVendorSpecific interface{}
}

func (docsL2vpnPktClassEntry *DOCSL2VPNMIB_DocsL2vpnPktClassTable_DocsL2vpnPktClassEntry) GetEntityData() *types.CommonEntityData {
    docsL2vpnPktClassEntry.EntityData.YFilter = docsL2vpnPktClassEntry.YFilter
    docsL2vpnPktClassEntry.EntityData.YangName = "docsL2vpnPktClassEntry"
    docsL2vpnPktClassEntry.EntityData.BundleName = "cisco_ios_xe"
    docsL2vpnPktClassEntry.EntityData.ParentYangName = "docsL2vpnPktClassTable"
    docsL2vpnPktClassEntry.EntityData.SegmentPath = "docsL2vpnPktClassEntry" + types.AddKeyToken(docsL2vpnPktClassEntry.IfIndex, "ifIndex") + types.AddKeyToken(docsL2vpnPktClassEntry.DocsQosServiceFlowId, "docsQosServiceFlowId") + types.AddKeyToken(docsL2vpnPktClassEntry.DocsQosPktClassId, "docsQosPktClassId")
    docsL2vpnPktClassEntry.EntityData.AbsolutePath = "DOCS-L2VPN-MIB:DOCS-L2VPN-MIB/docsL2vpnPktClassTable/" + docsL2vpnPktClassEntry.EntityData.SegmentPath
    docsL2vpnPktClassEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsL2vpnPktClassEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsL2vpnPktClassEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsL2vpnPktClassEntry.EntityData.Children = types.NewOrderedMap()
    docsL2vpnPktClassEntry.EntityData.Leafs = types.NewOrderedMap()
    docsL2vpnPktClassEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", docsL2vpnPktClassEntry.IfIndex})
    docsL2vpnPktClassEntry.EntityData.Leafs.Append("docsQosServiceFlowId", types.YLeaf{"DocsQosServiceFlowId", docsL2vpnPktClassEntry.DocsQosServiceFlowId})
    docsL2vpnPktClassEntry.EntityData.Leafs.Append("docsQosPktClassId", types.YLeaf{"DocsQosPktClassId", docsL2vpnPktClassEntry.DocsQosPktClassId})
    docsL2vpnPktClassEntry.EntityData.Leafs.Append("docsL2vpnPktClassL2vpnId", types.YLeaf{"DocsL2vpnPktClassL2vpnId", docsL2vpnPktClassEntry.DocsL2vpnPktClassL2vpnId})
    docsL2vpnPktClassEntry.EntityData.Leafs.Append("docsL2vpnPktClassUserPriRangeLow", types.YLeaf{"DocsL2vpnPktClassUserPriRangeLow", docsL2vpnPktClassEntry.DocsL2vpnPktClassUserPriRangeLow})
    docsL2vpnPktClassEntry.EntityData.Leafs.Append("docsL2vpnPktClassUserPriRangeHigh", types.YLeaf{"DocsL2vpnPktClassUserPriRangeHigh", docsL2vpnPktClassEntry.DocsL2vpnPktClassUserPriRangeHigh})
    docsL2vpnPktClassEntry.EntityData.Leafs.Append("docsL2vpnPktClassCMIM", types.YLeaf{"DocsL2vpnPktClassCMIM", docsL2vpnPktClassEntry.DocsL2vpnPktClassCMIM})
    docsL2vpnPktClassEntry.EntityData.Leafs.Append("docsL2vpnPktClassVendorSpecific", types.YLeaf{"DocsL2vpnPktClassVendorSpecific", docsL2vpnPktClassEntry.DocsL2vpnPktClassVendorSpecific})

    docsL2vpnPktClassEntry.EntityData.YListKeys = []string {"IfIndex", "DocsQosServiceFlowId", "DocsQosPktClassId"}

    return &(docsL2vpnPktClassEntry.EntityData)
}

// DOCSL2VPNMIB_DocsL2vpnCmNsiTable
// This table describes the NSI configuration for a single 
// CM when operating in point-to-point forwarding mode for an 
// L2VPN.
type DOCSL2VPNMIB_DocsL2vpnCmNsiTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry indexed by VPN ID and Cable Modem Index that  describes the
    // point-to-point forwarding between a single  NSI encapsulation and a single
    // CM. This table is   implemented only for a CM forwarding an L2VPN on a  
    // point-to-point basis. It is associated with a single   per-CM L2VPN
    // encoding. The type is slice of
    // DOCSL2VPNMIB_DocsL2vpnCmNsiTable_DocsL2vpnCmNsiEntry.
    DocsL2vpnCmNsiEntry []*DOCSL2VPNMIB_DocsL2vpnCmNsiTable_DocsL2vpnCmNsiEntry
}

func (docsL2vpnCmNsiTable *DOCSL2VPNMIB_DocsL2vpnCmNsiTable) GetEntityData() *types.CommonEntityData {
    docsL2vpnCmNsiTable.EntityData.YFilter = docsL2vpnCmNsiTable.YFilter
    docsL2vpnCmNsiTable.EntityData.YangName = "docsL2vpnCmNsiTable"
    docsL2vpnCmNsiTable.EntityData.BundleName = "cisco_ios_xe"
    docsL2vpnCmNsiTable.EntityData.ParentYangName = "DOCS-L2VPN-MIB"
    docsL2vpnCmNsiTable.EntityData.SegmentPath = "docsL2vpnCmNsiTable"
    docsL2vpnCmNsiTable.EntityData.AbsolutePath = "DOCS-L2VPN-MIB:DOCS-L2VPN-MIB/" + docsL2vpnCmNsiTable.EntityData.SegmentPath
    docsL2vpnCmNsiTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsL2vpnCmNsiTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsL2vpnCmNsiTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsL2vpnCmNsiTable.EntityData.Children = types.NewOrderedMap()
    docsL2vpnCmNsiTable.EntityData.Children.Append("docsL2vpnCmNsiEntry", types.YChild{"DocsL2vpnCmNsiEntry", nil})
    for i := range docsL2vpnCmNsiTable.DocsL2vpnCmNsiEntry {
        docsL2vpnCmNsiTable.EntityData.Children.Append(types.GetSegmentPath(docsL2vpnCmNsiTable.DocsL2vpnCmNsiEntry[i]), types.YChild{"DocsL2vpnCmNsiEntry", docsL2vpnCmNsiTable.DocsL2vpnCmNsiEntry[i]})
    }
    docsL2vpnCmNsiTable.EntityData.Leafs = types.NewOrderedMap()

    docsL2vpnCmNsiTable.EntityData.YListKeys = []string {}

    return &(docsL2vpnCmNsiTable.EntityData)
}

// DOCSL2VPNMIB_DocsL2vpnCmNsiTable_DocsL2vpnCmNsiEntry
// An entry indexed by VPN ID and Cable Modem Index that 
// describes the point-to-point forwarding between a single 
// NSI encapsulation and a single CM. This table is  
// implemented only for a CM forwarding an L2VPN on a  
// point-to-point basis. It is associated with a single  
// per-CM L2VPN encoding.
type DOCSL2VPNMIB_DocsL2vpnCmNsiTable_DocsL2vpnCmNsiEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // docs_l2vpn_mib.DOCSL2VPNMIB_DocsL2vpnIndexToIdTable_DocsL2vpnIndexToIdEntry_DocsL2vpnIdx
    DocsL2vpnIdx interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // docs_if_mib.DOCSIFMIB_DocsIfCmtsCmStatusTable_DocsIfCmtsCmStatusEntry_DocsIfCmtsCmStatusIndex
    DocsIfCmtsCmStatusIndex interface{}

    // The General Encapsulation Information (GEI) subtype of the  Network System
    // Interface (NSI) encapsulation configured  for the CM. The type is
    // DocsNsiEncapSubtype.
    DocsL2vpnCmNsiEncapSubtype interface{}

    // The encapsulation value for L2VPN forwarded packets on NSI  ports. The type
    // is string.
    DocsL2vpnCmNsiEncapValue interface{}

    // This object is the configuration of any Attachment Group   Identifier
    // subtype in the per-SF L2VPN Encoding   represented by this row. If the
    // subtype was omitted, this   object's value is a zero length string. The
    // type is string.
    DocsL2vpnCmNsiAGI interface{}

    // This object is the configuration of any Source   Attachment Individual ID
    // subtype in the L2VPN Encoding   represented by this row. If the subtype was
    // omitted, this  object's value is a zero length string. The type is string.
    DocsL2vpnCmNsiSAII interface{}

    // This object is the configuration of any Target  Attachment Individual ID
    // subtype in the L2VPN Encoding  represented by this row. If the subtype was
    // omitted, this  object's value is a zero length string. The type is string.
    DocsL2vpnCmNsiTAII interface{}
}

func (docsL2vpnCmNsiEntry *DOCSL2VPNMIB_DocsL2vpnCmNsiTable_DocsL2vpnCmNsiEntry) GetEntityData() *types.CommonEntityData {
    docsL2vpnCmNsiEntry.EntityData.YFilter = docsL2vpnCmNsiEntry.YFilter
    docsL2vpnCmNsiEntry.EntityData.YangName = "docsL2vpnCmNsiEntry"
    docsL2vpnCmNsiEntry.EntityData.BundleName = "cisco_ios_xe"
    docsL2vpnCmNsiEntry.EntityData.ParentYangName = "docsL2vpnCmNsiTable"
    docsL2vpnCmNsiEntry.EntityData.SegmentPath = "docsL2vpnCmNsiEntry" + types.AddKeyToken(docsL2vpnCmNsiEntry.DocsL2vpnIdx, "docsL2vpnIdx") + types.AddKeyToken(docsL2vpnCmNsiEntry.DocsIfCmtsCmStatusIndex, "docsIfCmtsCmStatusIndex")
    docsL2vpnCmNsiEntry.EntityData.AbsolutePath = "DOCS-L2VPN-MIB:DOCS-L2VPN-MIB/docsL2vpnCmNsiTable/" + docsL2vpnCmNsiEntry.EntityData.SegmentPath
    docsL2vpnCmNsiEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsL2vpnCmNsiEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsL2vpnCmNsiEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsL2vpnCmNsiEntry.EntityData.Children = types.NewOrderedMap()
    docsL2vpnCmNsiEntry.EntityData.Leafs = types.NewOrderedMap()
    docsL2vpnCmNsiEntry.EntityData.Leafs.Append("docsL2vpnIdx", types.YLeaf{"DocsL2vpnIdx", docsL2vpnCmNsiEntry.DocsL2vpnIdx})
    docsL2vpnCmNsiEntry.EntityData.Leafs.Append("docsIfCmtsCmStatusIndex", types.YLeaf{"DocsIfCmtsCmStatusIndex", docsL2vpnCmNsiEntry.DocsIfCmtsCmStatusIndex})
    docsL2vpnCmNsiEntry.EntityData.Leafs.Append("docsL2vpnCmNsiEncapSubtype", types.YLeaf{"DocsL2vpnCmNsiEncapSubtype", docsL2vpnCmNsiEntry.DocsL2vpnCmNsiEncapSubtype})
    docsL2vpnCmNsiEntry.EntityData.Leafs.Append("docsL2vpnCmNsiEncapValue", types.YLeaf{"DocsL2vpnCmNsiEncapValue", docsL2vpnCmNsiEntry.DocsL2vpnCmNsiEncapValue})
    docsL2vpnCmNsiEntry.EntityData.Leafs.Append("docsL2vpnCmNsiAGI", types.YLeaf{"DocsL2vpnCmNsiAGI", docsL2vpnCmNsiEntry.DocsL2vpnCmNsiAGI})
    docsL2vpnCmNsiEntry.EntityData.Leafs.Append("docsL2vpnCmNsiSAII", types.YLeaf{"DocsL2vpnCmNsiSAII", docsL2vpnCmNsiEntry.DocsL2vpnCmNsiSAII})
    docsL2vpnCmNsiEntry.EntityData.Leafs.Append("docsL2vpnCmNsiTAII", types.YLeaf{"DocsL2vpnCmNsiTAII", docsL2vpnCmNsiEntry.DocsL2vpnCmNsiTAII})

    docsL2vpnCmNsiEntry.EntityData.YListKeys = []string {"DocsL2vpnIdx", "DocsIfCmtsCmStatusIndex"}

    return &(docsL2vpnCmNsiEntry.EntityData)
}

// DOCSL2VPNMIB_DocsL2vpnCmVpnCpeTable
// This table is a list of CPEs, indexed by the VPNs on a  
// Cable Modem.
type DOCSL2VPNMIB_DocsL2vpnCmVpnCpeTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This table is a list of CPEs, indexed by the VPNs on a   Cable Modem. The
    // type is slice of
    // DOCSL2VPNMIB_DocsL2vpnCmVpnCpeTable_DocsL2vpnCmVpnCpeEntry.
    DocsL2vpnCmVpnCpeEntry []*DOCSL2VPNMIB_DocsL2vpnCmVpnCpeTable_DocsL2vpnCmVpnCpeEntry
}

func (docsL2vpnCmVpnCpeTable *DOCSL2VPNMIB_DocsL2vpnCmVpnCpeTable) GetEntityData() *types.CommonEntityData {
    docsL2vpnCmVpnCpeTable.EntityData.YFilter = docsL2vpnCmVpnCpeTable.YFilter
    docsL2vpnCmVpnCpeTable.EntityData.YangName = "docsL2vpnCmVpnCpeTable"
    docsL2vpnCmVpnCpeTable.EntityData.BundleName = "cisco_ios_xe"
    docsL2vpnCmVpnCpeTable.EntityData.ParentYangName = "DOCS-L2VPN-MIB"
    docsL2vpnCmVpnCpeTable.EntityData.SegmentPath = "docsL2vpnCmVpnCpeTable"
    docsL2vpnCmVpnCpeTable.EntityData.AbsolutePath = "DOCS-L2VPN-MIB:DOCS-L2VPN-MIB/" + docsL2vpnCmVpnCpeTable.EntityData.SegmentPath
    docsL2vpnCmVpnCpeTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsL2vpnCmVpnCpeTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsL2vpnCmVpnCpeTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsL2vpnCmVpnCpeTable.EntityData.Children = types.NewOrderedMap()
    docsL2vpnCmVpnCpeTable.EntityData.Children.Append("docsL2vpnCmVpnCpeEntry", types.YChild{"DocsL2vpnCmVpnCpeEntry", nil})
    for i := range docsL2vpnCmVpnCpeTable.DocsL2vpnCmVpnCpeEntry {
        docsL2vpnCmVpnCpeTable.EntityData.Children.Append(types.GetSegmentPath(docsL2vpnCmVpnCpeTable.DocsL2vpnCmVpnCpeEntry[i]), types.YChild{"DocsL2vpnCmVpnCpeEntry", docsL2vpnCmVpnCpeTable.DocsL2vpnCmVpnCpeEntry[i]})
    }
    docsL2vpnCmVpnCpeTable.EntityData.Leafs = types.NewOrderedMap()

    docsL2vpnCmVpnCpeTable.EntityData.YListKeys = []string {}

    return &(docsL2vpnCmVpnCpeTable.EntityData)
}

// DOCSL2VPNMIB_DocsL2vpnCmVpnCpeTable_DocsL2vpnCmVpnCpeEntry
// This table is a list of CPEs, indexed by the VPNs on a  
// Cable Modem.
type DOCSL2VPNMIB_DocsL2vpnCmVpnCpeTable_DocsL2vpnCmVpnCpeEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // docs_if_mib.DOCSIFMIB_DocsIfCmtsCmStatusTable_DocsIfCmtsCmStatusEntry_DocsIfCmtsCmStatusIndex
    DocsIfCmtsCmStatusIndex interface{}

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // docs_l2vpn_mib.DOCSL2VPNMIB_DocsL2vpnIndexToIdTable_DocsL2vpnIndexToIdEntry_DocsL2vpnIdx
    DocsL2vpnIdx interface{}

    // This attribute is a key. The Customer Premise Equipment (CPE) Mac Address 
    // that is attached to this instances Cable Modem  and bridging on this
    // instance's VPN Id. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    DocsL2vpnCmVpnCpeMacAddress interface{}
}

func (docsL2vpnCmVpnCpeEntry *DOCSL2VPNMIB_DocsL2vpnCmVpnCpeTable_DocsL2vpnCmVpnCpeEntry) GetEntityData() *types.CommonEntityData {
    docsL2vpnCmVpnCpeEntry.EntityData.YFilter = docsL2vpnCmVpnCpeEntry.YFilter
    docsL2vpnCmVpnCpeEntry.EntityData.YangName = "docsL2vpnCmVpnCpeEntry"
    docsL2vpnCmVpnCpeEntry.EntityData.BundleName = "cisco_ios_xe"
    docsL2vpnCmVpnCpeEntry.EntityData.ParentYangName = "docsL2vpnCmVpnCpeTable"
    docsL2vpnCmVpnCpeEntry.EntityData.SegmentPath = "docsL2vpnCmVpnCpeEntry" + types.AddKeyToken(docsL2vpnCmVpnCpeEntry.DocsIfCmtsCmStatusIndex, "docsIfCmtsCmStatusIndex") + types.AddKeyToken(docsL2vpnCmVpnCpeEntry.DocsL2vpnIdx, "docsL2vpnIdx") + types.AddKeyToken(docsL2vpnCmVpnCpeEntry.DocsL2vpnCmVpnCpeMacAddress, "docsL2vpnCmVpnCpeMacAddress")
    docsL2vpnCmVpnCpeEntry.EntityData.AbsolutePath = "DOCS-L2VPN-MIB:DOCS-L2VPN-MIB/docsL2vpnCmVpnCpeTable/" + docsL2vpnCmVpnCpeEntry.EntityData.SegmentPath
    docsL2vpnCmVpnCpeEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsL2vpnCmVpnCpeEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsL2vpnCmVpnCpeEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsL2vpnCmVpnCpeEntry.EntityData.Children = types.NewOrderedMap()
    docsL2vpnCmVpnCpeEntry.EntityData.Leafs = types.NewOrderedMap()
    docsL2vpnCmVpnCpeEntry.EntityData.Leafs.Append("docsIfCmtsCmStatusIndex", types.YLeaf{"DocsIfCmtsCmStatusIndex", docsL2vpnCmVpnCpeEntry.DocsIfCmtsCmStatusIndex})
    docsL2vpnCmVpnCpeEntry.EntityData.Leafs.Append("docsL2vpnIdx", types.YLeaf{"DocsL2vpnIdx", docsL2vpnCmVpnCpeEntry.DocsL2vpnIdx})
    docsL2vpnCmVpnCpeEntry.EntityData.Leafs.Append("docsL2vpnCmVpnCpeMacAddress", types.YLeaf{"DocsL2vpnCmVpnCpeMacAddress", docsL2vpnCmVpnCpeEntry.DocsL2vpnCmVpnCpeMacAddress})

    docsL2vpnCmVpnCpeEntry.EntityData.YListKeys = []string {"DocsIfCmtsCmStatusIndex", "DocsL2vpnIdx", "DocsL2vpnCmVpnCpeMacAddress"}

    return &(docsL2vpnCmVpnCpeEntry.EntityData)
}

// DOCSL2VPNMIB_DocsL2vpnVpnCmCpeTable
// This table contains a list of bridging CPEs, indexed by 
// L2VPN Index and the corresponding CMs on that VPN.
type DOCSL2VPNMIB_DocsL2vpnVpnCmCpeTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This table contains a list of bridging CPEs, indexed by  VPN and the
    // corresponding CMs on that VPN. The type is slice of
    // DOCSL2VPNMIB_DocsL2vpnVpnCmCpeTable_DocsL2vpnVpnCmCpeEntry.
    DocsL2vpnVpnCmCpeEntry []*DOCSL2VPNMIB_DocsL2vpnVpnCmCpeTable_DocsL2vpnVpnCmCpeEntry
}

func (docsL2vpnVpnCmCpeTable *DOCSL2VPNMIB_DocsL2vpnVpnCmCpeTable) GetEntityData() *types.CommonEntityData {
    docsL2vpnVpnCmCpeTable.EntityData.YFilter = docsL2vpnVpnCmCpeTable.YFilter
    docsL2vpnVpnCmCpeTable.EntityData.YangName = "docsL2vpnVpnCmCpeTable"
    docsL2vpnVpnCmCpeTable.EntityData.BundleName = "cisco_ios_xe"
    docsL2vpnVpnCmCpeTable.EntityData.ParentYangName = "DOCS-L2VPN-MIB"
    docsL2vpnVpnCmCpeTable.EntityData.SegmentPath = "docsL2vpnVpnCmCpeTable"
    docsL2vpnVpnCmCpeTable.EntityData.AbsolutePath = "DOCS-L2VPN-MIB:DOCS-L2VPN-MIB/" + docsL2vpnVpnCmCpeTable.EntityData.SegmentPath
    docsL2vpnVpnCmCpeTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsL2vpnVpnCmCpeTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsL2vpnVpnCmCpeTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsL2vpnVpnCmCpeTable.EntityData.Children = types.NewOrderedMap()
    docsL2vpnVpnCmCpeTable.EntityData.Children.Append("docsL2vpnVpnCmCpeEntry", types.YChild{"DocsL2vpnVpnCmCpeEntry", nil})
    for i := range docsL2vpnVpnCmCpeTable.DocsL2vpnVpnCmCpeEntry {
        docsL2vpnVpnCmCpeTable.EntityData.Children.Append(types.GetSegmentPath(docsL2vpnVpnCmCpeTable.DocsL2vpnVpnCmCpeEntry[i]), types.YChild{"DocsL2vpnVpnCmCpeEntry", docsL2vpnVpnCmCpeTable.DocsL2vpnVpnCmCpeEntry[i]})
    }
    docsL2vpnVpnCmCpeTable.EntityData.Leafs = types.NewOrderedMap()

    docsL2vpnVpnCmCpeTable.EntityData.YListKeys = []string {}

    return &(docsL2vpnVpnCmCpeTable.EntityData)
}

// DOCSL2VPNMIB_DocsL2vpnVpnCmCpeTable_DocsL2vpnVpnCmCpeEntry
// This table contains a list of bridging CPEs, indexed by 
// VPN and the corresponding CMs on that VPN.
type DOCSL2VPNMIB_DocsL2vpnVpnCmCpeTable_DocsL2vpnVpnCmCpeEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // docs_l2vpn_mib.DOCSL2VPNMIB_DocsL2vpnIndexToIdTable_DocsL2vpnIndexToIdEntry_DocsL2vpnIdx
    DocsL2vpnIdx interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // docs_if_mib.DOCSIFMIB_DocsIfCmtsCmStatusTable_DocsIfCmtsCmStatusEntry_DocsIfCmtsCmStatusIndex
    DocsIfCmtsCmStatusIndex interface{}

    // This attribute is a key. The Customer Premise Equipment (CPE) Mac Address 
    // that is attached to this instances Cable Modem  and bridging on this
    // instance's L2vpn Index. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    DocsL2vpnVpnCmCpeMacAddress interface{}
}

func (docsL2vpnVpnCmCpeEntry *DOCSL2VPNMIB_DocsL2vpnVpnCmCpeTable_DocsL2vpnVpnCmCpeEntry) GetEntityData() *types.CommonEntityData {
    docsL2vpnVpnCmCpeEntry.EntityData.YFilter = docsL2vpnVpnCmCpeEntry.YFilter
    docsL2vpnVpnCmCpeEntry.EntityData.YangName = "docsL2vpnVpnCmCpeEntry"
    docsL2vpnVpnCmCpeEntry.EntityData.BundleName = "cisco_ios_xe"
    docsL2vpnVpnCmCpeEntry.EntityData.ParentYangName = "docsL2vpnVpnCmCpeTable"
    docsL2vpnVpnCmCpeEntry.EntityData.SegmentPath = "docsL2vpnVpnCmCpeEntry" + types.AddKeyToken(docsL2vpnVpnCmCpeEntry.DocsL2vpnIdx, "docsL2vpnIdx") + types.AddKeyToken(docsL2vpnVpnCmCpeEntry.DocsIfCmtsCmStatusIndex, "docsIfCmtsCmStatusIndex") + types.AddKeyToken(docsL2vpnVpnCmCpeEntry.DocsL2vpnVpnCmCpeMacAddress, "docsL2vpnVpnCmCpeMacAddress")
    docsL2vpnVpnCmCpeEntry.EntityData.AbsolutePath = "DOCS-L2VPN-MIB:DOCS-L2VPN-MIB/docsL2vpnVpnCmCpeTable/" + docsL2vpnVpnCmCpeEntry.EntityData.SegmentPath
    docsL2vpnVpnCmCpeEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsL2vpnVpnCmCpeEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsL2vpnVpnCmCpeEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsL2vpnVpnCmCpeEntry.EntityData.Children = types.NewOrderedMap()
    docsL2vpnVpnCmCpeEntry.EntityData.Leafs = types.NewOrderedMap()
    docsL2vpnVpnCmCpeEntry.EntityData.Leafs.Append("docsL2vpnIdx", types.YLeaf{"DocsL2vpnIdx", docsL2vpnVpnCmCpeEntry.DocsL2vpnIdx})
    docsL2vpnVpnCmCpeEntry.EntityData.Leafs.Append("docsIfCmtsCmStatusIndex", types.YLeaf{"DocsIfCmtsCmStatusIndex", docsL2vpnVpnCmCpeEntry.DocsIfCmtsCmStatusIndex})
    docsL2vpnVpnCmCpeEntry.EntityData.Leafs.Append("docsL2vpnVpnCmCpeMacAddress", types.YLeaf{"DocsL2vpnVpnCmCpeMacAddress", docsL2vpnVpnCmCpeEntry.DocsL2vpnVpnCmCpeMacAddress})

    docsL2vpnVpnCmCpeEntry.EntityData.YListKeys = []string {"DocsL2vpnIdx", "DocsIfCmtsCmStatusIndex", "DocsL2vpnVpnCmCpeMacAddress"}

    return &(docsL2vpnVpnCmCpeEntry.EntityData)
}

// DOCSL2VPNMIB_DocsL2vpnDot1qTpFdbExtTable
// This table contains packet counters for  
// Unicast MAC Addresses within a VPN.
type DOCSL2VPNMIB_DocsL2vpnDot1qTpFdbExtTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This table extends the dot1qTpFdbTable only for RF network  bridge port
    // entries.  It is implemented by an agent only  if the agent implements
    // dot1qTpFdbTable for RF network  bridge ports. The type is slice of
    // DOCSL2VPNMIB_DocsL2vpnDot1qTpFdbExtTable_DocsL2vpnDot1qTpFdbExtEntry.
    DocsL2vpnDot1qTpFdbExtEntry []*DOCSL2VPNMIB_DocsL2vpnDot1qTpFdbExtTable_DocsL2vpnDot1qTpFdbExtEntry
}

func (docsL2vpnDot1qTpFdbExtTable *DOCSL2VPNMIB_DocsL2vpnDot1qTpFdbExtTable) GetEntityData() *types.CommonEntityData {
    docsL2vpnDot1qTpFdbExtTable.EntityData.YFilter = docsL2vpnDot1qTpFdbExtTable.YFilter
    docsL2vpnDot1qTpFdbExtTable.EntityData.YangName = "docsL2vpnDot1qTpFdbExtTable"
    docsL2vpnDot1qTpFdbExtTable.EntityData.BundleName = "cisco_ios_xe"
    docsL2vpnDot1qTpFdbExtTable.EntityData.ParentYangName = "DOCS-L2VPN-MIB"
    docsL2vpnDot1qTpFdbExtTable.EntityData.SegmentPath = "docsL2vpnDot1qTpFdbExtTable"
    docsL2vpnDot1qTpFdbExtTable.EntityData.AbsolutePath = "DOCS-L2VPN-MIB:DOCS-L2VPN-MIB/" + docsL2vpnDot1qTpFdbExtTable.EntityData.SegmentPath
    docsL2vpnDot1qTpFdbExtTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsL2vpnDot1qTpFdbExtTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsL2vpnDot1qTpFdbExtTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsL2vpnDot1qTpFdbExtTable.EntityData.Children = types.NewOrderedMap()
    docsL2vpnDot1qTpFdbExtTable.EntityData.Children.Append("docsL2vpnDot1qTpFdbExtEntry", types.YChild{"DocsL2vpnDot1qTpFdbExtEntry", nil})
    for i := range docsL2vpnDot1qTpFdbExtTable.DocsL2vpnDot1qTpFdbExtEntry {
        docsL2vpnDot1qTpFdbExtTable.EntityData.Children.Append(types.GetSegmentPath(docsL2vpnDot1qTpFdbExtTable.DocsL2vpnDot1qTpFdbExtEntry[i]), types.YChild{"DocsL2vpnDot1qTpFdbExtEntry", docsL2vpnDot1qTpFdbExtTable.DocsL2vpnDot1qTpFdbExtEntry[i]})
    }
    docsL2vpnDot1qTpFdbExtTable.EntityData.Leafs = types.NewOrderedMap()

    docsL2vpnDot1qTpFdbExtTable.EntityData.YListKeys = []string {}

    return &(docsL2vpnDot1qTpFdbExtTable.EntityData)
}

// DOCSL2VPNMIB_DocsL2vpnDot1qTpFdbExtTable_DocsL2vpnDot1qTpFdbExtEntry
// This table extends the dot1qTpFdbTable only for RF network 
// bridge port entries.  It is implemented by an agent only 
// if the agent implements dot1qTpFdbTable for RF network 
// bridge ports.
type DOCSL2VPNMIB_DocsL2vpnDot1qTpFdbExtTable_DocsL2vpnDot1qTpFdbExtEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to q_bridge_mib.QBRIDGEMIB_Dot1qFdbTable_Dot1qFdbEntry_Dot1qFdbId
    Dot1qFdbId interface{}

    // This attribute is a key. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}. Refers to
    // q_bridge_mib.QBRIDGEMIB_Dot1qTpFdbTable_Dot1qTpFdbEntry_Dot1qTpFdbAddress
    Dot1qTpFdbAddress interface{}

    // The number of packets where the Destination   MAC Address matched this
    // instance   dot1qTpFdbAddress and packet was bridged on  a VPN, where the
    // VPN ID matched this   instance's dot1qFdbId. The type is interface{} with
    // range: 0..4294967295.
    DocsL2vpnDot1qTpFdbExtTransmitPkts interface{}

    // The number of packets where the Source MAC   Address matched this instance
    // dot1qTpFdbAddress  and the packet was bridged on a VPN,  where the
    // docsL2vpnIdx matched this instance's   dot1qFdbId. The type is interface{}
    // with range: 0..4294967295.
    DocsL2vpnDot1qTpFdbExtReceivePkts interface{}
}

func (docsL2vpnDot1qTpFdbExtEntry *DOCSL2VPNMIB_DocsL2vpnDot1qTpFdbExtTable_DocsL2vpnDot1qTpFdbExtEntry) GetEntityData() *types.CommonEntityData {
    docsL2vpnDot1qTpFdbExtEntry.EntityData.YFilter = docsL2vpnDot1qTpFdbExtEntry.YFilter
    docsL2vpnDot1qTpFdbExtEntry.EntityData.YangName = "docsL2vpnDot1qTpFdbExtEntry"
    docsL2vpnDot1qTpFdbExtEntry.EntityData.BundleName = "cisco_ios_xe"
    docsL2vpnDot1qTpFdbExtEntry.EntityData.ParentYangName = "docsL2vpnDot1qTpFdbExtTable"
    docsL2vpnDot1qTpFdbExtEntry.EntityData.SegmentPath = "docsL2vpnDot1qTpFdbExtEntry" + types.AddKeyToken(docsL2vpnDot1qTpFdbExtEntry.Dot1qFdbId, "dot1qFdbId") + types.AddKeyToken(docsL2vpnDot1qTpFdbExtEntry.Dot1qTpFdbAddress, "dot1qTpFdbAddress")
    docsL2vpnDot1qTpFdbExtEntry.EntityData.AbsolutePath = "DOCS-L2VPN-MIB:DOCS-L2VPN-MIB/docsL2vpnDot1qTpFdbExtTable/" + docsL2vpnDot1qTpFdbExtEntry.EntityData.SegmentPath
    docsL2vpnDot1qTpFdbExtEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsL2vpnDot1qTpFdbExtEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsL2vpnDot1qTpFdbExtEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsL2vpnDot1qTpFdbExtEntry.EntityData.Children = types.NewOrderedMap()
    docsL2vpnDot1qTpFdbExtEntry.EntityData.Leafs = types.NewOrderedMap()
    docsL2vpnDot1qTpFdbExtEntry.EntityData.Leafs.Append("dot1qFdbId", types.YLeaf{"Dot1qFdbId", docsL2vpnDot1qTpFdbExtEntry.Dot1qFdbId})
    docsL2vpnDot1qTpFdbExtEntry.EntityData.Leafs.Append("dot1qTpFdbAddress", types.YLeaf{"Dot1qTpFdbAddress", docsL2vpnDot1qTpFdbExtEntry.Dot1qTpFdbAddress})
    docsL2vpnDot1qTpFdbExtEntry.EntityData.Leafs.Append("docsL2vpnDot1qTpFdbExtTransmitPkts", types.YLeaf{"DocsL2vpnDot1qTpFdbExtTransmitPkts", docsL2vpnDot1qTpFdbExtEntry.DocsL2vpnDot1qTpFdbExtTransmitPkts})
    docsL2vpnDot1qTpFdbExtEntry.EntityData.Leafs.Append("docsL2vpnDot1qTpFdbExtReceivePkts", types.YLeaf{"DocsL2vpnDot1qTpFdbExtReceivePkts", docsL2vpnDot1qTpFdbExtEntry.DocsL2vpnDot1qTpFdbExtReceivePkts})

    docsL2vpnDot1qTpFdbExtEntry.EntityData.YListKeys = []string {"Dot1qFdbId", "Dot1qTpFdbAddress"}

    return &(docsL2vpnDot1qTpFdbExtEntry.EntityData)
}

// DOCSL2VPNMIB_DocsL2vpnDot1qTpGroupExtTable
// This table contains packet counters for  
// Multicast MAC Addresses within a VPN.
type DOCSL2VPNMIB_DocsL2vpnDot1qTpGroupExtTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This table extends the dot1qTpGroupTable only for RF   Network bridge port
    // entries.  It is implemented by an agent  Only if the agent implements
    // dot1qTpGroupTable for RF   network bridge ports. The type is slice of
    // DOCSL2VPNMIB_DocsL2vpnDot1qTpGroupExtTable_DocsL2vpnDot1qTpGroupExtEntry.
    DocsL2vpnDot1qTpGroupExtEntry []*DOCSL2VPNMIB_DocsL2vpnDot1qTpGroupExtTable_DocsL2vpnDot1qTpGroupExtEntry
}

func (docsL2vpnDot1qTpGroupExtTable *DOCSL2VPNMIB_DocsL2vpnDot1qTpGroupExtTable) GetEntityData() *types.CommonEntityData {
    docsL2vpnDot1qTpGroupExtTable.EntityData.YFilter = docsL2vpnDot1qTpGroupExtTable.YFilter
    docsL2vpnDot1qTpGroupExtTable.EntityData.YangName = "docsL2vpnDot1qTpGroupExtTable"
    docsL2vpnDot1qTpGroupExtTable.EntityData.BundleName = "cisco_ios_xe"
    docsL2vpnDot1qTpGroupExtTable.EntityData.ParentYangName = "DOCS-L2VPN-MIB"
    docsL2vpnDot1qTpGroupExtTable.EntityData.SegmentPath = "docsL2vpnDot1qTpGroupExtTable"
    docsL2vpnDot1qTpGroupExtTable.EntityData.AbsolutePath = "DOCS-L2VPN-MIB:DOCS-L2VPN-MIB/" + docsL2vpnDot1qTpGroupExtTable.EntityData.SegmentPath
    docsL2vpnDot1qTpGroupExtTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsL2vpnDot1qTpGroupExtTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsL2vpnDot1qTpGroupExtTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsL2vpnDot1qTpGroupExtTable.EntityData.Children = types.NewOrderedMap()
    docsL2vpnDot1qTpGroupExtTable.EntityData.Children.Append("docsL2vpnDot1qTpGroupExtEntry", types.YChild{"DocsL2vpnDot1qTpGroupExtEntry", nil})
    for i := range docsL2vpnDot1qTpGroupExtTable.DocsL2vpnDot1qTpGroupExtEntry {
        docsL2vpnDot1qTpGroupExtTable.EntityData.Children.Append(types.GetSegmentPath(docsL2vpnDot1qTpGroupExtTable.DocsL2vpnDot1qTpGroupExtEntry[i]), types.YChild{"DocsL2vpnDot1qTpGroupExtEntry", docsL2vpnDot1qTpGroupExtTable.DocsL2vpnDot1qTpGroupExtEntry[i]})
    }
    docsL2vpnDot1qTpGroupExtTable.EntityData.Leafs = types.NewOrderedMap()

    docsL2vpnDot1qTpGroupExtTable.EntityData.YListKeys = []string {}

    return &(docsL2vpnDot1qTpGroupExtTable.EntityData)
}

// DOCSL2VPNMIB_DocsL2vpnDot1qTpGroupExtTable_DocsL2vpnDot1qTpGroupExtEntry
// This table extends the dot1qTpGroupTable only for RF  
// Network bridge port entries.  It is implemented by an agent 
// Only if the agent implements dot1qTpGroupTable for RF  
// network bridge ports.
type DOCSL2VPNMIB_DocsL2vpnDot1qTpGroupExtTable_DocsL2vpnDot1qTpGroupExtEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // q_bridge_mib.QBRIDGEMIB_Dot1qVlanCurrentTable_Dot1qVlanCurrentEntry_Dot1qVlanIndex
    Dot1qVlanIndex interface{}

    // This attribute is a key. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}. Refers to
    // q_bridge_mib.QBRIDGEMIB_Dot1qTpGroupTable_Dot1qTpGroupEntry_Dot1qTpGroupAddress
    Dot1qTpGroupAddress interface{}

    // The number of packets where the Destination   MAC Address matched this
    // instance   dot1qTpGroupAddress and packet was bridged on  a VPN, where the
    // docsL2vpnIdx matched this   instance's dot1qVlanIndex. The type is
    // interface{} with range: 0..4294967295.
    DocsL2vpnDot1qTpGroupExtTransmitPkts interface{}

    // The number of packets where the Source MAC   Address matched this instance
    // dot1qTpGroupAddress  and the packet was bridged on a VPN,  where the
    // docsL2vpnIdx matched this instance's   dot1qVlanIndex. The type is
    // interface{} with range: 0..4294967295.
    DocsL2vpnDot1qTpGroupExtReceivePkts interface{}
}

func (docsL2vpnDot1qTpGroupExtEntry *DOCSL2VPNMIB_DocsL2vpnDot1qTpGroupExtTable_DocsL2vpnDot1qTpGroupExtEntry) GetEntityData() *types.CommonEntityData {
    docsL2vpnDot1qTpGroupExtEntry.EntityData.YFilter = docsL2vpnDot1qTpGroupExtEntry.YFilter
    docsL2vpnDot1qTpGroupExtEntry.EntityData.YangName = "docsL2vpnDot1qTpGroupExtEntry"
    docsL2vpnDot1qTpGroupExtEntry.EntityData.BundleName = "cisco_ios_xe"
    docsL2vpnDot1qTpGroupExtEntry.EntityData.ParentYangName = "docsL2vpnDot1qTpGroupExtTable"
    docsL2vpnDot1qTpGroupExtEntry.EntityData.SegmentPath = "docsL2vpnDot1qTpGroupExtEntry" + types.AddKeyToken(docsL2vpnDot1qTpGroupExtEntry.Dot1qVlanIndex, "dot1qVlanIndex") + types.AddKeyToken(docsL2vpnDot1qTpGroupExtEntry.Dot1qTpGroupAddress, "dot1qTpGroupAddress")
    docsL2vpnDot1qTpGroupExtEntry.EntityData.AbsolutePath = "DOCS-L2VPN-MIB:DOCS-L2VPN-MIB/docsL2vpnDot1qTpGroupExtTable/" + docsL2vpnDot1qTpGroupExtEntry.EntityData.SegmentPath
    docsL2vpnDot1qTpGroupExtEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsL2vpnDot1qTpGroupExtEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsL2vpnDot1qTpGroupExtEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsL2vpnDot1qTpGroupExtEntry.EntityData.Children = types.NewOrderedMap()
    docsL2vpnDot1qTpGroupExtEntry.EntityData.Leafs = types.NewOrderedMap()
    docsL2vpnDot1qTpGroupExtEntry.EntityData.Leafs.Append("dot1qVlanIndex", types.YLeaf{"Dot1qVlanIndex", docsL2vpnDot1qTpGroupExtEntry.Dot1qVlanIndex})
    docsL2vpnDot1qTpGroupExtEntry.EntityData.Leafs.Append("dot1qTpGroupAddress", types.YLeaf{"Dot1qTpGroupAddress", docsL2vpnDot1qTpGroupExtEntry.Dot1qTpGroupAddress})
    docsL2vpnDot1qTpGroupExtEntry.EntityData.Leafs.Append("docsL2vpnDot1qTpGroupExtTransmitPkts", types.YLeaf{"DocsL2vpnDot1qTpGroupExtTransmitPkts", docsL2vpnDot1qTpGroupExtEntry.DocsL2vpnDot1qTpGroupExtTransmitPkts})
    docsL2vpnDot1qTpGroupExtEntry.EntityData.Leafs.Append("docsL2vpnDot1qTpGroupExtReceivePkts", types.YLeaf{"DocsL2vpnDot1qTpGroupExtReceivePkts", docsL2vpnDot1qTpGroupExtEntry.DocsL2vpnDot1qTpGroupExtReceivePkts})

    docsL2vpnDot1qTpGroupExtEntry.EntityData.YListKeys = []string {"Dot1qVlanIndex", "Dot1qTpGroupAddress"}

    return &(docsL2vpnDot1qTpGroupExtEntry.EntityData)
}

