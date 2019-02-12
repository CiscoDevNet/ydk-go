// This MIB complements the CISCO-IETF-PW-MIB for PW operation 
// over MPLS. 
package cisco_ietf_pw_mpls_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_ietf_pw_mpls_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-IETF-PW-MPLS-MIB CISCO-IETF-PW-MPLS-MIB}", reflect.TypeOf(CISCOIETFPWMPLSMIB{}))
    ydk.RegisterEntity("CISCO-IETF-PW-MPLS-MIB:CISCO-IETF-PW-MPLS-MIB", reflect.TypeOf(CISCOIETFPWMPLSMIB{}))
}

// CISCOIETFPWMPLSMIB
type CISCOIETFPWMPLSMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    CpwVcMplsObjects CISCOIETFPWMPLSMIB_CpwVcMplsObjects

    // This table specifies information for VC to be carried over   MPLS PSN.
    CpwVcMplsTable CISCOIETFPWMPLSMIB_CpwVcMplsTable

    // This table associates VCs using MPLS PSN with the outbound  MPLS tunnels
    // (i.e. toward the PSN) or the physical   interface in case of VC only.
    CpwVcMplsOutboundTable CISCOIETFPWMPLSMIB_CpwVcMplsOutboundTable

    // This table associates VCs using MPLS PSN with the inbound  MPLS tunnels
    // (i.e. for packets coming from the PSN),   if such association is desired
    // (mainly for security   reasons).
    CpwVcMplsInboundTable CISCOIETFPWMPLSMIB_CpwVcMplsInboundTable

    // This table maps an inbound/outbound Tunnel to a VC in non-  TE
    // applications.
    CpwVcMplsNonTeMappingTable CISCOIETFPWMPLSMIB_CpwVcMplsNonTeMappingTable

    // This table maps an inbound/outbound Tunnel to a VC in   MPLS-TE
    // applications.
    CpwVcMplsTeMappingTable CISCOIETFPWMPLSMIB_CpwVcMplsTeMappingTable
}

func (cISCOIETFPWMPLSMIB *CISCOIETFPWMPLSMIB) GetEntityData() *types.CommonEntityData {
    cISCOIETFPWMPLSMIB.EntityData.YFilter = cISCOIETFPWMPLSMIB.YFilter
    cISCOIETFPWMPLSMIB.EntityData.YangName = "CISCO-IETF-PW-MPLS-MIB"
    cISCOIETFPWMPLSMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOIETFPWMPLSMIB.EntityData.ParentYangName = "CISCO-IETF-PW-MPLS-MIB"
    cISCOIETFPWMPLSMIB.EntityData.SegmentPath = "CISCO-IETF-PW-MPLS-MIB:CISCO-IETF-PW-MPLS-MIB"
    cISCOIETFPWMPLSMIB.EntityData.AbsolutePath = cISCOIETFPWMPLSMIB.EntityData.SegmentPath
    cISCOIETFPWMPLSMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOIETFPWMPLSMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOIETFPWMPLSMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOIETFPWMPLSMIB.EntityData.Children = types.NewOrderedMap()
    cISCOIETFPWMPLSMIB.EntityData.Children.Append("cpwVcMplsObjects", types.YChild{"CpwVcMplsObjects", &cISCOIETFPWMPLSMIB.CpwVcMplsObjects})
    cISCOIETFPWMPLSMIB.EntityData.Children.Append("cpwVcMplsTable", types.YChild{"CpwVcMplsTable", &cISCOIETFPWMPLSMIB.CpwVcMplsTable})
    cISCOIETFPWMPLSMIB.EntityData.Children.Append("cpwVcMplsOutboundTable", types.YChild{"CpwVcMplsOutboundTable", &cISCOIETFPWMPLSMIB.CpwVcMplsOutboundTable})
    cISCOIETFPWMPLSMIB.EntityData.Children.Append("cpwVcMplsInboundTable", types.YChild{"CpwVcMplsInboundTable", &cISCOIETFPWMPLSMIB.CpwVcMplsInboundTable})
    cISCOIETFPWMPLSMIB.EntityData.Children.Append("cpwVcMplsNonTeMappingTable", types.YChild{"CpwVcMplsNonTeMappingTable", &cISCOIETFPWMPLSMIB.CpwVcMplsNonTeMappingTable})
    cISCOIETFPWMPLSMIB.EntityData.Children.Append("cpwVcMplsTeMappingTable", types.YChild{"CpwVcMplsTeMappingTable", &cISCOIETFPWMPLSMIB.CpwVcMplsTeMappingTable})
    cISCOIETFPWMPLSMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOIETFPWMPLSMIB.EntityData.YListKeys = []string {}

    return &(cISCOIETFPWMPLSMIB.EntityData)
}

// CISCOIETFPWMPLSMIB_CpwVcMplsObjects
type CISCOIETFPWMPLSMIB_CpwVcMplsObjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object contains an appropriate value to  be used for
    // cpwVcMplsOutboundIndex when creating  entries in the
    // cpwVcMplsOutboundTable. The value  0 indicates that no unassigned entries
    // are  available. To obtain the cpwVcMplsOutboundIndex  value for a new
    // entry, the manager issues a  management protocol retrieval operation to
    // obtain  the current value of this object.  After each  retrieval, the agent
    // should modify the value to  the next unassigned index, however the agent
    // MUST  NOT assume such retrieval will be done for each   row created. The
    // type is interface{} with range: 0..4294967295.
    CpwVcMplsOutboundIndexNext interface{}

    // This object contains an appropriate value to  be used for
    // cpwVcMplsInboundIndex when creating  entries in the cpwVcMplsInboundTable.
    // The value  0 indicates that no unassigned entries are  available. To obtain
    // the cpwVcMplsInboundIndex  value for a new entry, the manager issues a 
    // management protocol retrieval operation to obtain  the current value of
    // this object.  After each  retrieval, the agent should modify the value to 
    // the next unassigned index, however the agent MUST  NOT assume such
    // retrieval will be done for each   row created. The type is interface{} with
    // range: 0..4294967295.
    CpwVcMplsInboundIndexNext interface{}
}

func (cpwVcMplsObjects *CISCOIETFPWMPLSMIB_CpwVcMplsObjects) GetEntityData() *types.CommonEntityData {
    cpwVcMplsObjects.EntityData.YFilter = cpwVcMplsObjects.YFilter
    cpwVcMplsObjects.EntityData.YangName = "cpwVcMplsObjects"
    cpwVcMplsObjects.EntityData.BundleName = "cisco_ios_xe"
    cpwVcMplsObjects.EntityData.ParentYangName = "CISCO-IETF-PW-MPLS-MIB"
    cpwVcMplsObjects.EntityData.SegmentPath = "cpwVcMplsObjects"
    cpwVcMplsObjects.EntityData.AbsolutePath = "CISCO-IETF-PW-MPLS-MIB:CISCO-IETF-PW-MPLS-MIB/" + cpwVcMplsObjects.EntityData.SegmentPath
    cpwVcMplsObjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwVcMplsObjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwVcMplsObjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwVcMplsObjects.EntityData.Children = types.NewOrderedMap()
    cpwVcMplsObjects.EntityData.Leafs = types.NewOrderedMap()
    cpwVcMplsObjects.EntityData.Leafs.Append("cpwVcMplsOutboundIndexNext", types.YLeaf{"CpwVcMplsOutboundIndexNext", cpwVcMplsObjects.CpwVcMplsOutboundIndexNext})
    cpwVcMplsObjects.EntityData.Leafs.Append("cpwVcMplsInboundIndexNext", types.YLeaf{"CpwVcMplsInboundIndexNext", cpwVcMplsObjects.CpwVcMplsInboundIndexNext})

    cpwVcMplsObjects.EntityData.YListKeys = []string {}

    return &(cpwVcMplsObjects.EntityData)
}

// CISCOIETFPWMPLSMIB_CpwVcMplsTable
// This table specifies information for VC to be carried over  
// MPLS PSN.
type CISCOIETFPWMPLSMIB_CpwVcMplsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A row in this table represents parameters specific to MPLS   PSN for a
    // pseudo wire connection (VC). The row is created   automatically by the
    // local agent if the cpwVcPsnType is   MPLS. It is indexed by cpwVcIndex,
    // which uniquely   identifying a singular connection. . The type is slice of
    // CISCOIETFPWMPLSMIB_CpwVcMplsTable_CpwVcMplsEntry.
    CpwVcMplsEntry []*CISCOIETFPWMPLSMIB_CpwVcMplsTable_CpwVcMplsEntry
}

func (cpwVcMplsTable *CISCOIETFPWMPLSMIB_CpwVcMplsTable) GetEntityData() *types.CommonEntityData {
    cpwVcMplsTable.EntityData.YFilter = cpwVcMplsTable.YFilter
    cpwVcMplsTable.EntityData.YangName = "cpwVcMplsTable"
    cpwVcMplsTable.EntityData.BundleName = "cisco_ios_xe"
    cpwVcMplsTable.EntityData.ParentYangName = "CISCO-IETF-PW-MPLS-MIB"
    cpwVcMplsTable.EntityData.SegmentPath = "cpwVcMplsTable"
    cpwVcMplsTable.EntityData.AbsolutePath = "CISCO-IETF-PW-MPLS-MIB:CISCO-IETF-PW-MPLS-MIB/" + cpwVcMplsTable.EntityData.SegmentPath
    cpwVcMplsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwVcMplsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwVcMplsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwVcMplsTable.EntityData.Children = types.NewOrderedMap()
    cpwVcMplsTable.EntityData.Children.Append("cpwVcMplsEntry", types.YChild{"CpwVcMplsEntry", nil})
    for i := range cpwVcMplsTable.CpwVcMplsEntry {
        cpwVcMplsTable.EntityData.Children.Append(types.GetSegmentPath(cpwVcMplsTable.CpwVcMplsEntry[i]), types.YChild{"CpwVcMplsEntry", cpwVcMplsTable.CpwVcMplsEntry[i]})
    }
    cpwVcMplsTable.EntityData.Leafs = types.NewOrderedMap()

    cpwVcMplsTable.EntityData.YListKeys = []string {}

    return &(cpwVcMplsTable.EntityData)
}

// CISCOIETFPWMPLSMIB_CpwVcMplsTable_CpwVcMplsEntry
// A row in this table represents parameters specific to MPLS  
// PSN for a pseudo wire connection (VC). The row is created  
// automatically by the local agent if the cpwVcPsnType is  
// MPLS. It is indexed by cpwVcIndex, which uniquely  
// identifying a singular connection. 
type CISCOIETFPWMPLSMIB_CpwVcMplsTable_CpwVcMplsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to cisco_ietf_pw_mib.CISCOIETFPWMIB_CpwVcTable_CpwVcEntry_CpwVcIndex
    CpwVcIndex interface{}

    // Set by the operator to indicate the outer tunnel types, if  exists. mplsTe
    // is used if the outer tunnel was set-up by   MPLS-TE, and mplsNonTe is used
    // the outer tunnel was set up  by LDP or manually. Combination of mplsTe and
    // mplsNonTe   may exist in case of outer tunnel protection.  vcOnly is used
    // if there is no outer tunnel label. vcOnly   cannot be combined with
    // mplsNonTe or mplsTe. The type is map[string]bool.
    CpwVcMplsMplsType interface{}

    // Set by the operator to indicate the way the VC shim label  EXP bits are to
    // be determined. The value of outerTunnel(1)  is used where there is an outer
    // tunnel - cpwVcMplsMplsType   is mplsTe or mplsNonTe. Note that in this case
    // there is no  need to mark the VC label with the EXP bits since the VC  
    // label is not visible to the intermediate nodes.  If there is no outer
    // tunnel, specifiedValue(2) indicate   that the value is specified by
    // cpwVcMplsExpBits, and   serviceDependant(3) indicate that the EXP bits are
    // setup   based on a rule specified in the emulated service specific  
    // tables, for example when the EXP bits are a function of   802.1p marking
    // for Ethernet emulated service. The type is CpwVcMplsExpBitsMode.
    CpwVcMplsExpBitsMode interface{}

    // Set by the operator to indicate the MPLS EXP bits to be   used on the VC
    // shim label if cpwVcMplsExpBitsMode is    specifiedValue(2), zero otherwise.
    // The type is interface{} with range: 0..7.
    CpwVcMplsExpBits interface{}

    // Set by the operator to indicate the VC TTL bits to be used  on the VC shim
    // label. The type is interface{} with range: 0..255.
    CpwVcMplsTtl interface{}

    // The local LDP identifier of the LDP entity creating  this VC in the local
    // node. As the VC labels are always  set from the per platform label space,
    // the last two octets   in the LDP ID MUST be always both zeros. The type is
    // string.
    CpwVcMplsLocalLdpID interface{}

    // The local LDP Entity index of the LDP entity to be used   for this VC on
    // the local node. Should be set to all zeros   if not used. The type is
    // interface{} with range: 0..4294967295.
    CpwVcMplsLocalLdpEntityID interface{}

    // The peer LDP identifier as identified from the LDP   session. Should be
    // zero if not relevant or not known yet. The type is string.
    CpwVcMplsPeerLdpID interface{}

    // This variable indicates the storage type for this row. The type is
    // StorageType.
    CpwVcMplsStorageType interface{}
}

func (cpwVcMplsEntry *CISCOIETFPWMPLSMIB_CpwVcMplsTable_CpwVcMplsEntry) GetEntityData() *types.CommonEntityData {
    cpwVcMplsEntry.EntityData.YFilter = cpwVcMplsEntry.YFilter
    cpwVcMplsEntry.EntityData.YangName = "cpwVcMplsEntry"
    cpwVcMplsEntry.EntityData.BundleName = "cisco_ios_xe"
    cpwVcMplsEntry.EntityData.ParentYangName = "cpwVcMplsTable"
    cpwVcMplsEntry.EntityData.SegmentPath = "cpwVcMplsEntry" + types.AddKeyToken(cpwVcMplsEntry.CpwVcIndex, "cpwVcIndex")
    cpwVcMplsEntry.EntityData.AbsolutePath = "CISCO-IETF-PW-MPLS-MIB:CISCO-IETF-PW-MPLS-MIB/cpwVcMplsTable/" + cpwVcMplsEntry.EntityData.SegmentPath
    cpwVcMplsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwVcMplsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwVcMplsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwVcMplsEntry.EntityData.Children = types.NewOrderedMap()
    cpwVcMplsEntry.EntityData.Leafs = types.NewOrderedMap()
    cpwVcMplsEntry.EntityData.Leafs.Append("cpwVcIndex", types.YLeaf{"CpwVcIndex", cpwVcMplsEntry.CpwVcIndex})
    cpwVcMplsEntry.EntityData.Leafs.Append("cpwVcMplsMplsType", types.YLeaf{"CpwVcMplsMplsType", cpwVcMplsEntry.CpwVcMplsMplsType})
    cpwVcMplsEntry.EntityData.Leafs.Append("cpwVcMplsExpBitsMode", types.YLeaf{"CpwVcMplsExpBitsMode", cpwVcMplsEntry.CpwVcMplsExpBitsMode})
    cpwVcMplsEntry.EntityData.Leafs.Append("cpwVcMplsExpBits", types.YLeaf{"CpwVcMplsExpBits", cpwVcMplsEntry.CpwVcMplsExpBits})
    cpwVcMplsEntry.EntityData.Leafs.Append("cpwVcMplsTtl", types.YLeaf{"CpwVcMplsTtl", cpwVcMplsEntry.CpwVcMplsTtl})
    cpwVcMplsEntry.EntityData.Leafs.Append("cpwVcMplsLocalLdpID", types.YLeaf{"CpwVcMplsLocalLdpID", cpwVcMplsEntry.CpwVcMplsLocalLdpID})
    cpwVcMplsEntry.EntityData.Leafs.Append("cpwVcMplsLocalLdpEntityID", types.YLeaf{"CpwVcMplsLocalLdpEntityID", cpwVcMplsEntry.CpwVcMplsLocalLdpEntityID})
    cpwVcMplsEntry.EntityData.Leafs.Append("cpwVcMplsPeerLdpID", types.YLeaf{"CpwVcMplsPeerLdpID", cpwVcMplsEntry.CpwVcMplsPeerLdpID})
    cpwVcMplsEntry.EntityData.Leafs.Append("cpwVcMplsStorageType", types.YLeaf{"CpwVcMplsStorageType", cpwVcMplsEntry.CpwVcMplsStorageType})

    cpwVcMplsEntry.EntityData.YListKeys = []string {"CpwVcIndex"}

    return &(cpwVcMplsEntry.EntityData)
}

// CISCOIETFPWMPLSMIB_CpwVcMplsTable_CpwVcMplsEntry_CpwVcMplsExpBitsMode represents 802.1p marking for Ethernet emulated service.
type CISCOIETFPWMPLSMIB_CpwVcMplsTable_CpwVcMplsEntry_CpwVcMplsExpBitsMode string

const (
    CISCOIETFPWMPLSMIB_CpwVcMplsTable_CpwVcMplsEntry_CpwVcMplsExpBitsMode_outerTunnel CISCOIETFPWMPLSMIB_CpwVcMplsTable_CpwVcMplsEntry_CpwVcMplsExpBitsMode = "outerTunnel"

    CISCOIETFPWMPLSMIB_CpwVcMplsTable_CpwVcMplsEntry_CpwVcMplsExpBitsMode_specifiedValue CISCOIETFPWMPLSMIB_CpwVcMplsTable_CpwVcMplsEntry_CpwVcMplsExpBitsMode = "specifiedValue"

    CISCOIETFPWMPLSMIB_CpwVcMplsTable_CpwVcMplsEntry_CpwVcMplsExpBitsMode_serviceDependant CISCOIETFPWMPLSMIB_CpwVcMplsTable_CpwVcMplsEntry_CpwVcMplsExpBitsMode = "serviceDependant"
)

// CISCOIETFPWMPLSMIB_CpwVcMplsOutboundTable
// This table associates VCs using MPLS PSN with the outbound 
// MPLS tunnels (i.e. toward the PSN) or the physical  
// interface in case of VC only.
type CISCOIETFPWMPLSMIB_CpwVcMplsOutboundTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A row in this table represents a link between PW VC (that  require MPLS
    // tunnels) and MPLS tunnel toward the PSN.  In the case of VC only, it
    // associate the VC with the   interface that shall carry the VC.  This table
    // is indexed by the pwVcIndex and an additional  index enabling multiple rows
    // for the same VC index.   At least one entry is created in this table by the
    // operator   for each PW VC that requires MPLS PSN. Note that the first 
    // entry for each VC can be indexed by cpwVcMplsOutboundIndex   equal zero
    // without a need for retrieval of   cpwVcMplsOutboundIndexNext.   This table
    // points to the appropriate MPLS MIB. In the case   of MPLS-TE, the 4
    // variables relevant to the indexing of   a TE MPLS tunnel are set as in
    // Srinivasan, et al, <draft-  ietf-mpls-te-mib>.  In case of Non-TE MPLS (an
    // outer tunnel label assigned by   LDP or manually) the table points to the
    // XC entry in the   LSR MIB as in Srinivasan, et al,
    // <draft-ietf-mpls-lsr-mib>.  In case of VC only (no outer tunnel) the
    // ifIndex of the  port to carry the VC is configured.    Each VC may have
    // multiple rows in this tables if protection   is available at the outer
    // tunnel level, each row may be of  different type except for VC only, on
    // which only rows with  ifIndex of the port are allowed. . The type is slice
    // of CISCOIETFPWMPLSMIB_CpwVcMplsOutboundTable_CpwVcMplsOutboundEntry.
    CpwVcMplsOutboundEntry []*CISCOIETFPWMPLSMIB_CpwVcMplsOutboundTable_CpwVcMplsOutboundEntry
}

func (cpwVcMplsOutboundTable *CISCOIETFPWMPLSMIB_CpwVcMplsOutboundTable) GetEntityData() *types.CommonEntityData {
    cpwVcMplsOutboundTable.EntityData.YFilter = cpwVcMplsOutboundTable.YFilter
    cpwVcMplsOutboundTable.EntityData.YangName = "cpwVcMplsOutboundTable"
    cpwVcMplsOutboundTable.EntityData.BundleName = "cisco_ios_xe"
    cpwVcMplsOutboundTable.EntityData.ParentYangName = "CISCO-IETF-PW-MPLS-MIB"
    cpwVcMplsOutboundTable.EntityData.SegmentPath = "cpwVcMplsOutboundTable"
    cpwVcMplsOutboundTable.EntityData.AbsolutePath = "CISCO-IETF-PW-MPLS-MIB:CISCO-IETF-PW-MPLS-MIB/" + cpwVcMplsOutboundTable.EntityData.SegmentPath
    cpwVcMplsOutboundTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwVcMplsOutboundTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwVcMplsOutboundTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwVcMplsOutboundTable.EntityData.Children = types.NewOrderedMap()
    cpwVcMplsOutboundTable.EntityData.Children.Append("cpwVcMplsOutboundEntry", types.YChild{"CpwVcMplsOutboundEntry", nil})
    for i := range cpwVcMplsOutboundTable.CpwVcMplsOutboundEntry {
        cpwVcMplsOutboundTable.EntityData.Children.Append(types.GetSegmentPath(cpwVcMplsOutboundTable.CpwVcMplsOutboundEntry[i]), types.YChild{"CpwVcMplsOutboundEntry", cpwVcMplsOutboundTable.CpwVcMplsOutboundEntry[i]})
    }
    cpwVcMplsOutboundTable.EntityData.Leafs = types.NewOrderedMap()

    cpwVcMplsOutboundTable.EntityData.YListKeys = []string {}

    return &(cpwVcMplsOutboundTable.EntityData)
}

// CISCOIETFPWMPLSMIB_CpwVcMplsOutboundTable_CpwVcMplsOutboundEntry
// A row in this table represents a link between PW VC (that 
// require MPLS tunnels) and MPLS tunnel toward the PSN. 
// In the case of VC only, it associate the VC with the  
// interface that shall carry the VC. 
// This table is indexed by the pwVcIndex and an additional 
// index enabling multiple rows for the same VC index. 
// 
// At least one entry is created in this table by the operator  
// for each PW VC that requires MPLS PSN. Note that the first 
// entry for each VC can be indexed by cpwVcMplsOutboundIndex  
// equal zero without a need for retrieval of  
// cpwVcMplsOutboundIndexNext. 
// 
// This table points to the appropriate MPLS MIB. In the case  
// of MPLS-TE, the 4 variables relevant to the indexing of  
// a TE MPLS tunnel are set as in Srinivasan, et al, <draft- 
// ietf-mpls-te-mib>. 
// In case of Non-TE MPLS (an outer tunnel label assigned by  
// LDP or manually) the table points to the XC entry in the  
// LSR MIB as in Srinivasan, et al, <draft-ietf-mpls-lsr-mib>. 
// In case of VC only (no outer tunnel) the ifIndex of the 
// port to carry the VC is configured.  
// 
// Each VC may have multiple rows in this tables if protection  
// is available at the outer tunnel level, each row may be of 
// different type except for VC only, on which only rows with 
// ifIndex of the port are allowed. 
type CISCOIETFPWMPLSMIB_CpwVcMplsOutboundTable_CpwVcMplsOutboundEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to cisco_ietf_pw_mib.CISCOIETFPWMIB_CpwVcTable_CpwVcEntry_CpwVcIndex
    CpwVcIndex interface{}

    // This attribute is a key. Arbitrary index for enabling multiple rows per VC
    // in  this table. Next available free index can be retrieved   using
    // cpwVcMplsOutboundIndexNext. . The type is interface{} with range:
    // 0..4294967295.
    CpwVcMplsOutboundIndex interface{}

    // This object will be set by the operator. If the outer  label is defined in
    // the MPL-LSR-MIB, i.e. set by LDP  or manually, this object points to the XC
    // index   of the outer tunnel. Otherwise, it is set to zero. The type is
    // interface{} with range: 0..4294967295.
    CpwVcMplsOutboundLsrXcIndex interface{}

    // Part of set of indexes for outbound tunnel in the case of   MPLS-TE outer
    // tunnel, otherwise set to zero. The type is interface{} with range:
    // 0..65535.
    CpwVcMplsOutboundTunnelIndex interface{}

    // Part of set of indexes for outbound tunnel in the case of   MPLS-TE outer
    // tunnel, otherwise set to zero. The type is interface{} with range:
    // 0..4294967295.
    CpwVcMplsOutboundTunnelInstance interface{}

    // Part of set of indexes for outbound tunnel in the case of   MPLS-TE outer
    // tunnel, otherwise set to zero. The type is string with length: 4.
    CpwVcMplsOutboundTunnelLclLSR interface{}

    // Part of set of indexes for outbound tunnel in the case of   MPLS-TE outer
    // tunnel, otherwise set to zero. The type is string with length: 4.
    CpwVcMplsOutboundTunnelPeerLSR interface{}

    // In case of VC only (no outer tunnel), this object holds  the ifIndex of the
    // outbound port, otherwise set to zero. The type is interface{} with range:
    // 0..2147483647.
    CpwVcMplsOutboundIfIndex interface{}

    // For creating, modifying, and deleting this row. The type is RowStatus.
    CpwVcMplsOutboundRowStatus interface{}

    // This variable indicates the storage type for this object. The type is
    // StorageType.
    CpwVcMplsOutboundStorageType interface{}
}

func (cpwVcMplsOutboundEntry *CISCOIETFPWMPLSMIB_CpwVcMplsOutboundTable_CpwVcMplsOutboundEntry) GetEntityData() *types.CommonEntityData {
    cpwVcMplsOutboundEntry.EntityData.YFilter = cpwVcMplsOutboundEntry.YFilter
    cpwVcMplsOutboundEntry.EntityData.YangName = "cpwVcMplsOutboundEntry"
    cpwVcMplsOutboundEntry.EntityData.BundleName = "cisco_ios_xe"
    cpwVcMplsOutboundEntry.EntityData.ParentYangName = "cpwVcMplsOutboundTable"
    cpwVcMplsOutboundEntry.EntityData.SegmentPath = "cpwVcMplsOutboundEntry" + types.AddKeyToken(cpwVcMplsOutboundEntry.CpwVcIndex, "cpwVcIndex") + types.AddKeyToken(cpwVcMplsOutboundEntry.CpwVcMplsOutboundIndex, "cpwVcMplsOutboundIndex")
    cpwVcMplsOutboundEntry.EntityData.AbsolutePath = "CISCO-IETF-PW-MPLS-MIB:CISCO-IETF-PW-MPLS-MIB/cpwVcMplsOutboundTable/" + cpwVcMplsOutboundEntry.EntityData.SegmentPath
    cpwVcMplsOutboundEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwVcMplsOutboundEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwVcMplsOutboundEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwVcMplsOutboundEntry.EntityData.Children = types.NewOrderedMap()
    cpwVcMplsOutboundEntry.EntityData.Leafs = types.NewOrderedMap()
    cpwVcMplsOutboundEntry.EntityData.Leafs.Append("cpwVcIndex", types.YLeaf{"CpwVcIndex", cpwVcMplsOutboundEntry.CpwVcIndex})
    cpwVcMplsOutboundEntry.EntityData.Leafs.Append("cpwVcMplsOutboundIndex", types.YLeaf{"CpwVcMplsOutboundIndex", cpwVcMplsOutboundEntry.CpwVcMplsOutboundIndex})
    cpwVcMplsOutboundEntry.EntityData.Leafs.Append("cpwVcMplsOutboundLsrXcIndex", types.YLeaf{"CpwVcMplsOutboundLsrXcIndex", cpwVcMplsOutboundEntry.CpwVcMplsOutboundLsrXcIndex})
    cpwVcMplsOutboundEntry.EntityData.Leafs.Append("cpwVcMplsOutboundTunnelIndex", types.YLeaf{"CpwVcMplsOutboundTunnelIndex", cpwVcMplsOutboundEntry.CpwVcMplsOutboundTunnelIndex})
    cpwVcMplsOutboundEntry.EntityData.Leafs.Append("cpwVcMplsOutboundTunnelInstance", types.YLeaf{"CpwVcMplsOutboundTunnelInstance", cpwVcMplsOutboundEntry.CpwVcMplsOutboundTunnelInstance})
    cpwVcMplsOutboundEntry.EntityData.Leafs.Append("cpwVcMplsOutboundTunnelLclLSR", types.YLeaf{"CpwVcMplsOutboundTunnelLclLSR", cpwVcMplsOutboundEntry.CpwVcMplsOutboundTunnelLclLSR})
    cpwVcMplsOutboundEntry.EntityData.Leafs.Append("cpwVcMplsOutboundTunnelPeerLSR", types.YLeaf{"CpwVcMplsOutboundTunnelPeerLSR", cpwVcMplsOutboundEntry.CpwVcMplsOutboundTunnelPeerLSR})
    cpwVcMplsOutboundEntry.EntityData.Leafs.Append("cpwVcMplsOutboundIfIndex", types.YLeaf{"CpwVcMplsOutboundIfIndex", cpwVcMplsOutboundEntry.CpwVcMplsOutboundIfIndex})
    cpwVcMplsOutboundEntry.EntityData.Leafs.Append("cpwVcMplsOutboundRowStatus", types.YLeaf{"CpwVcMplsOutboundRowStatus", cpwVcMplsOutboundEntry.CpwVcMplsOutboundRowStatus})
    cpwVcMplsOutboundEntry.EntityData.Leafs.Append("cpwVcMplsOutboundStorageType", types.YLeaf{"CpwVcMplsOutboundStorageType", cpwVcMplsOutboundEntry.CpwVcMplsOutboundStorageType})

    cpwVcMplsOutboundEntry.EntityData.YListKeys = []string {"CpwVcIndex", "CpwVcMplsOutboundIndex"}

    return &(cpwVcMplsOutboundEntry.EntityData)
}

// CISCOIETFPWMPLSMIB_CpwVcMplsInboundTable
// This table associates VCs using MPLS PSN with the inbound 
// MPLS tunnels (i.e. for packets coming from the PSN),  
// if such association is desired (mainly for security  
// reasons).
type CISCOIETFPWMPLSMIB_CpwVcMplsInboundTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A row in this table represents a link between PW VCs (that  require MPLS
    // tunnels) and MPLS tunnel for packets arriving  from the PSN.  This table is
    // indexed by the set of indexes used to  identify the VC - cpwVcIndex and an
    // additional  index enabling multiple rows for the same VC index.   Note that
    // the first entry for each VC can be indexed by   cpwVcMplsOutboundIndex
    // equal zero without a need for   retrieval of cpwVcMplsInboundIndexNext.  
    // An entry is created in this table either automatically by   the local agent
    // or created manually by the operator in   cases that strict mode is
    // required.   Note that the control messages contain VC ID and VC type,  
    // which together with the remote IP address identify the  cpwVcIndex in the
    // local node.  This table points to the appropriate MPLS MIB. In the case  
    // of MPLS-TE, the 4 variables relevant to the indexing of a  TE MPLS tunnel
    // are set as in Srinivasan, et al, <draft-  ietf-mpls-te-mib>.   In case of
    // non-TE MPLS tunnel (an outer tunnel label   assigned by LDP or manually)
    // the table points to the XC   entry in the MPLS-LSR-MIB as in Srinivasan, et
    // al, <draft-  ietf-mpls-lsr-mib>.   Each VC may have multiple rows in this
    // tables if protection   is available at the outer tunnel level, each row may
    // be of  different type except for VC only, on which only rows with  ifIndex
    // of the port are allowed. . The type is slice of
    // CISCOIETFPWMPLSMIB_CpwVcMplsInboundTable_CpwVcMplsInboundEntry.
    CpwVcMplsInboundEntry []*CISCOIETFPWMPLSMIB_CpwVcMplsInboundTable_CpwVcMplsInboundEntry
}

func (cpwVcMplsInboundTable *CISCOIETFPWMPLSMIB_CpwVcMplsInboundTable) GetEntityData() *types.CommonEntityData {
    cpwVcMplsInboundTable.EntityData.YFilter = cpwVcMplsInboundTable.YFilter
    cpwVcMplsInboundTable.EntityData.YangName = "cpwVcMplsInboundTable"
    cpwVcMplsInboundTable.EntityData.BundleName = "cisco_ios_xe"
    cpwVcMplsInboundTable.EntityData.ParentYangName = "CISCO-IETF-PW-MPLS-MIB"
    cpwVcMplsInboundTable.EntityData.SegmentPath = "cpwVcMplsInboundTable"
    cpwVcMplsInboundTable.EntityData.AbsolutePath = "CISCO-IETF-PW-MPLS-MIB:CISCO-IETF-PW-MPLS-MIB/" + cpwVcMplsInboundTable.EntityData.SegmentPath
    cpwVcMplsInboundTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwVcMplsInboundTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwVcMplsInboundTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwVcMplsInboundTable.EntityData.Children = types.NewOrderedMap()
    cpwVcMplsInboundTable.EntityData.Children.Append("cpwVcMplsInboundEntry", types.YChild{"CpwVcMplsInboundEntry", nil})
    for i := range cpwVcMplsInboundTable.CpwVcMplsInboundEntry {
        cpwVcMplsInboundTable.EntityData.Children.Append(types.GetSegmentPath(cpwVcMplsInboundTable.CpwVcMplsInboundEntry[i]), types.YChild{"CpwVcMplsInboundEntry", cpwVcMplsInboundTable.CpwVcMplsInboundEntry[i]})
    }
    cpwVcMplsInboundTable.EntityData.Leafs = types.NewOrderedMap()

    cpwVcMplsInboundTable.EntityData.YListKeys = []string {}

    return &(cpwVcMplsInboundTable.EntityData)
}

// CISCOIETFPWMPLSMIB_CpwVcMplsInboundTable_CpwVcMplsInboundEntry
// A row in this table represents a link between PW VCs (that 
// require MPLS tunnels) and MPLS tunnel for packets arriving 
// from the PSN. 
// This table is indexed by the set of indexes used to 
// identify the VC - cpwVcIndex and an additional 
// index enabling multiple rows for the same VC index. 
// 
// Note that the first entry for each VC can be indexed by  
// cpwVcMplsOutboundIndex equal zero without a need for  
// retrieval of cpwVcMplsInboundIndexNext. 
// 
// An entry is created in this table either automatically by  
// the local agent or created manually by the operator in  
// cases that strict mode is required. 
// 
// Note that the control messages contain VC ID and VC type,  
// which together with the remote IP address identify the 
// cpwVcIndex in the local node. 
// This table points to the appropriate MPLS MIB. In the case  
// of MPLS-TE, the 4 variables relevant to the indexing of a 
// TE MPLS tunnel are set as in Srinivasan, et al, <draft- 
// ietf-mpls-te-mib>. 
// 
// In case of non-TE MPLS tunnel (an outer tunnel label  
// assigned by LDP or manually) the table points to the XC  
// entry in the MPLS-LSR-MIB as in Srinivasan, et al, <draft- 
// ietf-mpls-lsr-mib>. 
// 
// Each VC may have multiple rows in this tables if protection  
// is available at the outer tunnel level, each row may be of 
// different type except for VC only, on which only rows with 
// ifIndex of the port are allowed. 
type CISCOIETFPWMPLSMIB_CpwVcMplsInboundTable_CpwVcMplsInboundEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to cisco_ietf_pw_mib.CISCOIETFPWMIB_CpwVcTable_CpwVcEntry_CpwVcIndex
    CpwVcIndex interface{}

    // This attribute is a key. Arbitrary index for enabling multiple rows per VC
    // in  this table. Next available free index can be retrieved  using
    // cpwVcMplsInboundIndexNext. . The type is interface{} with range:
    // 0..4294967295.
    CpwVcMplsInboundIndex interface{}

    // If the outer label is defined in the MPL-LSR-MIB, i.e. set   by LDP or
    // manually, this object points to the XC index   of the outer tunnel.
    // Otherwise, it is set to zero. The type is interface{} with range:
    // 0..4294967295.
    CpwVcMplsInboundLsrXcIndex interface{}

    // Part of set of indexes for outbound tunnel in the case of   MPLS-TE outer
    // tunnel, otherwise set to zero. The type is interface{} with range:
    // 0..65535.
    CpwVcMplsInboundTunnelIndex interface{}

    // Part of set of indexes for outbound tunnel in the case of   MPLS-TE outer
    // tunnel, otherwise set to zero. The type is interface{} with range:
    // 0..4294967295.
    CpwVcMplsInboundTunnelInstance interface{}

    // Part of set of indexes for outbound tunnel in the case of   MPLS-TE outer
    // tunnel, otherwise set to zero. The type is string with length: 4.
    CpwVcMplsInboundTunnelLclLSR interface{}

    // Part of set of indexes for outbound tunnel in the case of   MPLS-TE outer
    // tunnel, otherwise set to zero. The type is string with length: 4.
    CpwVcMplsInboundTunnelPeerLSR interface{}

    // In case of VC only (no outer tunnel), this object holds the  ifIndex of the
    // inbound port, otherwise set to zero. The type is interface{} with range:
    // 0..2147483647.
    CpwVcMplsInboundIfIndex interface{}

    // For creating, modifying, and deleting this row. The type is RowStatus.
    CpwVcMplsInboundRowStatus interface{}

    // This variable indicates the storage type for this row. The type is
    // StorageType.
    CpwVcMplsInboundStorageType interface{}
}

func (cpwVcMplsInboundEntry *CISCOIETFPWMPLSMIB_CpwVcMplsInboundTable_CpwVcMplsInboundEntry) GetEntityData() *types.CommonEntityData {
    cpwVcMplsInboundEntry.EntityData.YFilter = cpwVcMplsInboundEntry.YFilter
    cpwVcMplsInboundEntry.EntityData.YangName = "cpwVcMplsInboundEntry"
    cpwVcMplsInboundEntry.EntityData.BundleName = "cisco_ios_xe"
    cpwVcMplsInboundEntry.EntityData.ParentYangName = "cpwVcMplsInboundTable"
    cpwVcMplsInboundEntry.EntityData.SegmentPath = "cpwVcMplsInboundEntry" + types.AddKeyToken(cpwVcMplsInboundEntry.CpwVcIndex, "cpwVcIndex") + types.AddKeyToken(cpwVcMplsInboundEntry.CpwVcMplsInboundIndex, "cpwVcMplsInboundIndex")
    cpwVcMplsInboundEntry.EntityData.AbsolutePath = "CISCO-IETF-PW-MPLS-MIB:CISCO-IETF-PW-MPLS-MIB/cpwVcMplsInboundTable/" + cpwVcMplsInboundEntry.EntityData.SegmentPath
    cpwVcMplsInboundEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwVcMplsInboundEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwVcMplsInboundEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwVcMplsInboundEntry.EntityData.Children = types.NewOrderedMap()
    cpwVcMplsInboundEntry.EntityData.Leafs = types.NewOrderedMap()
    cpwVcMplsInboundEntry.EntityData.Leafs.Append("cpwVcIndex", types.YLeaf{"CpwVcIndex", cpwVcMplsInboundEntry.CpwVcIndex})
    cpwVcMplsInboundEntry.EntityData.Leafs.Append("cpwVcMplsInboundIndex", types.YLeaf{"CpwVcMplsInboundIndex", cpwVcMplsInboundEntry.CpwVcMplsInboundIndex})
    cpwVcMplsInboundEntry.EntityData.Leafs.Append("cpwVcMplsInboundLsrXcIndex", types.YLeaf{"CpwVcMplsInboundLsrXcIndex", cpwVcMplsInboundEntry.CpwVcMplsInboundLsrXcIndex})
    cpwVcMplsInboundEntry.EntityData.Leafs.Append("cpwVcMplsInboundTunnelIndex", types.YLeaf{"CpwVcMplsInboundTunnelIndex", cpwVcMplsInboundEntry.CpwVcMplsInboundTunnelIndex})
    cpwVcMplsInboundEntry.EntityData.Leafs.Append("cpwVcMplsInboundTunnelInstance", types.YLeaf{"CpwVcMplsInboundTunnelInstance", cpwVcMplsInboundEntry.CpwVcMplsInboundTunnelInstance})
    cpwVcMplsInboundEntry.EntityData.Leafs.Append("cpwVcMplsInboundTunnelLclLSR", types.YLeaf{"CpwVcMplsInboundTunnelLclLSR", cpwVcMplsInboundEntry.CpwVcMplsInboundTunnelLclLSR})
    cpwVcMplsInboundEntry.EntityData.Leafs.Append("cpwVcMplsInboundTunnelPeerLSR", types.YLeaf{"CpwVcMplsInboundTunnelPeerLSR", cpwVcMplsInboundEntry.CpwVcMplsInboundTunnelPeerLSR})
    cpwVcMplsInboundEntry.EntityData.Leafs.Append("cpwVcMplsInboundIfIndex", types.YLeaf{"CpwVcMplsInboundIfIndex", cpwVcMplsInboundEntry.CpwVcMplsInboundIfIndex})
    cpwVcMplsInboundEntry.EntityData.Leafs.Append("cpwVcMplsInboundRowStatus", types.YLeaf{"CpwVcMplsInboundRowStatus", cpwVcMplsInboundEntry.CpwVcMplsInboundRowStatus})
    cpwVcMplsInboundEntry.EntityData.Leafs.Append("cpwVcMplsInboundStorageType", types.YLeaf{"CpwVcMplsInboundStorageType", cpwVcMplsInboundEntry.CpwVcMplsInboundStorageType})

    cpwVcMplsInboundEntry.EntityData.YListKeys = []string {"CpwVcIndex", "CpwVcMplsInboundIndex"}

    return &(cpwVcMplsInboundEntry.EntityData)
}

// CISCOIETFPWMPLSMIB_CpwVcMplsNonTeMappingTable
// This table maps an inbound/outbound Tunnel to a VC in non- 
// TE applications.
type CISCOIETFPWMPLSMIB_CpwVcMplsNonTeMappingTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A row in this table represents the association  between the PW VC and it's
    // non TE MPLS outer Tunnel  it's physical interface if there is no outer
    // tunnel   (VC only).   An application can use this table to quickly retrieve
    // the   PW carried over specific non-TE MPLS outer tunnel or   physical
    // interface.   The table in indexed by the XC index for MPLS Non-TE   tunnel,
    // or ifIndex of the port in VC only case, the   direction of the VC in the
    // specific entry and the VCIndex.   The same table is used in both inbound
    // and outbound  directions, but in a different row for each direction. If  
    // the inbound association is not known, no rows should exist   for it.   Rows
    // are created by the local agent when all the   association data is available
    // for display. The type is slice of
    // CISCOIETFPWMPLSMIB_CpwVcMplsNonTeMappingTable_CpwVcMplsNonTeMappingEntry.
    CpwVcMplsNonTeMappingEntry []*CISCOIETFPWMPLSMIB_CpwVcMplsNonTeMappingTable_CpwVcMplsNonTeMappingEntry
}

func (cpwVcMplsNonTeMappingTable *CISCOIETFPWMPLSMIB_CpwVcMplsNonTeMappingTable) GetEntityData() *types.CommonEntityData {
    cpwVcMplsNonTeMappingTable.EntityData.YFilter = cpwVcMplsNonTeMappingTable.YFilter
    cpwVcMplsNonTeMappingTable.EntityData.YangName = "cpwVcMplsNonTeMappingTable"
    cpwVcMplsNonTeMappingTable.EntityData.BundleName = "cisco_ios_xe"
    cpwVcMplsNonTeMappingTable.EntityData.ParentYangName = "CISCO-IETF-PW-MPLS-MIB"
    cpwVcMplsNonTeMappingTable.EntityData.SegmentPath = "cpwVcMplsNonTeMappingTable"
    cpwVcMplsNonTeMappingTable.EntityData.AbsolutePath = "CISCO-IETF-PW-MPLS-MIB:CISCO-IETF-PW-MPLS-MIB/" + cpwVcMplsNonTeMappingTable.EntityData.SegmentPath
    cpwVcMplsNonTeMappingTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwVcMplsNonTeMappingTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwVcMplsNonTeMappingTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwVcMplsNonTeMappingTable.EntityData.Children = types.NewOrderedMap()
    cpwVcMplsNonTeMappingTable.EntityData.Children.Append("cpwVcMplsNonTeMappingEntry", types.YChild{"CpwVcMplsNonTeMappingEntry", nil})
    for i := range cpwVcMplsNonTeMappingTable.CpwVcMplsNonTeMappingEntry {
        cpwVcMplsNonTeMappingTable.EntityData.Children.Append(types.GetSegmentPath(cpwVcMplsNonTeMappingTable.CpwVcMplsNonTeMappingEntry[i]), types.YChild{"CpwVcMplsNonTeMappingEntry", cpwVcMplsNonTeMappingTable.CpwVcMplsNonTeMappingEntry[i]})
    }
    cpwVcMplsNonTeMappingTable.EntityData.Leafs = types.NewOrderedMap()

    cpwVcMplsNonTeMappingTable.EntityData.YListKeys = []string {}

    return &(cpwVcMplsNonTeMappingTable.EntityData)
}

// CISCOIETFPWMPLSMIB_CpwVcMplsNonTeMappingTable_CpwVcMplsNonTeMappingEntry
// A row in this table represents the association 
// between the PW VC and it's non TE MPLS outer Tunnel 
// it's physical interface if there is no outer tunnel  
// (VC only). 
// 
// An application can use this table to quickly retrieve the  
// PW carried over specific non-TE MPLS outer tunnel or  
// physical interface. 
// 
// The table in indexed by the XC index for MPLS Non-TE  
// tunnel, or ifIndex of the port in VC only case, the  
// direction of the VC in the specific entry and the VCIndex. 
// 
// The same table is used in both inbound and outbound 
// directions, but in a different row for each direction. If  
// the inbound association is not known, no rows should exist  
// for it. 
// 
// Rows are created by the local agent when all the  
// association data is available for display.
type CISCOIETFPWMPLSMIB_CpwVcMplsNonTeMappingTable_CpwVcMplsNonTeMappingEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Identifies if the row represent an outbound or
    // inbound   mapping. The type is CpwVcMplsNonTeMappingTunnelDirection.
    CpwVcMplsNonTeMappingTunnelDirection interface{}

    // This attribute is a key. Index for the conceptual XC row identifying Tunnel
    // to VC   mappings when the outer tunnel is created by the MPLS-LSR-  MIB,
    // Zero otherwise. The type is interface{} with range: 0..4294967295.
    CpwVcMplsNonTeMappingXcTunnelIndex interface{}

    // This attribute is a key. Identify the port on which the VC is carried for
    // VC only   case. The type is interface{} with range: 0..2147483647.
    CpwVcMplsNonTeMappingIfIndex interface{}

    // This attribute is a key. The value that represent the VC in the cpwVcTable.
    // The type is interface{} with range: 0..4294967295.
    CpwVcMplsNonTeMappingVcIndex interface{}
}

func (cpwVcMplsNonTeMappingEntry *CISCOIETFPWMPLSMIB_CpwVcMplsNonTeMappingTable_CpwVcMplsNonTeMappingEntry) GetEntityData() *types.CommonEntityData {
    cpwVcMplsNonTeMappingEntry.EntityData.YFilter = cpwVcMplsNonTeMappingEntry.YFilter
    cpwVcMplsNonTeMappingEntry.EntityData.YangName = "cpwVcMplsNonTeMappingEntry"
    cpwVcMplsNonTeMappingEntry.EntityData.BundleName = "cisco_ios_xe"
    cpwVcMplsNonTeMappingEntry.EntityData.ParentYangName = "cpwVcMplsNonTeMappingTable"
    cpwVcMplsNonTeMappingEntry.EntityData.SegmentPath = "cpwVcMplsNonTeMappingEntry" + types.AddKeyToken(cpwVcMplsNonTeMappingEntry.CpwVcMplsNonTeMappingTunnelDirection, "cpwVcMplsNonTeMappingTunnelDirection") + types.AddKeyToken(cpwVcMplsNonTeMappingEntry.CpwVcMplsNonTeMappingXcTunnelIndex, "cpwVcMplsNonTeMappingXcTunnelIndex") + types.AddKeyToken(cpwVcMplsNonTeMappingEntry.CpwVcMplsNonTeMappingIfIndex, "cpwVcMplsNonTeMappingIfIndex") + types.AddKeyToken(cpwVcMplsNonTeMappingEntry.CpwVcMplsNonTeMappingVcIndex, "cpwVcMplsNonTeMappingVcIndex")
    cpwVcMplsNonTeMappingEntry.EntityData.AbsolutePath = "CISCO-IETF-PW-MPLS-MIB:CISCO-IETF-PW-MPLS-MIB/cpwVcMplsNonTeMappingTable/" + cpwVcMplsNonTeMappingEntry.EntityData.SegmentPath
    cpwVcMplsNonTeMappingEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwVcMplsNonTeMappingEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwVcMplsNonTeMappingEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwVcMplsNonTeMappingEntry.EntityData.Children = types.NewOrderedMap()
    cpwVcMplsNonTeMappingEntry.EntityData.Leafs = types.NewOrderedMap()
    cpwVcMplsNonTeMappingEntry.EntityData.Leafs.Append("cpwVcMplsNonTeMappingTunnelDirection", types.YLeaf{"CpwVcMplsNonTeMappingTunnelDirection", cpwVcMplsNonTeMappingEntry.CpwVcMplsNonTeMappingTunnelDirection})
    cpwVcMplsNonTeMappingEntry.EntityData.Leafs.Append("cpwVcMplsNonTeMappingXcTunnelIndex", types.YLeaf{"CpwVcMplsNonTeMappingXcTunnelIndex", cpwVcMplsNonTeMappingEntry.CpwVcMplsNonTeMappingXcTunnelIndex})
    cpwVcMplsNonTeMappingEntry.EntityData.Leafs.Append("cpwVcMplsNonTeMappingIfIndex", types.YLeaf{"CpwVcMplsNonTeMappingIfIndex", cpwVcMplsNonTeMappingEntry.CpwVcMplsNonTeMappingIfIndex})
    cpwVcMplsNonTeMappingEntry.EntityData.Leafs.Append("cpwVcMplsNonTeMappingVcIndex", types.YLeaf{"CpwVcMplsNonTeMappingVcIndex", cpwVcMplsNonTeMappingEntry.CpwVcMplsNonTeMappingVcIndex})

    cpwVcMplsNonTeMappingEntry.EntityData.YListKeys = []string {"CpwVcMplsNonTeMappingTunnelDirection", "CpwVcMplsNonTeMappingXcTunnelIndex", "CpwVcMplsNonTeMappingIfIndex", "CpwVcMplsNonTeMappingVcIndex"}

    return &(cpwVcMplsNonTeMappingEntry.EntityData)
}

// CISCOIETFPWMPLSMIB_CpwVcMplsNonTeMappingTable_CpwVcMplsNonTeMappingEntry_CpwVcMplsNonTeMappingTunnelDirection represents mapping.
type CISCOIETFPWMPLSMIB_CpwVcMplsNonTeMappingTable_CpwVcMplsNonTeMappingEntry_CpwVcMplsNonTeMappingTunnelDirection string

const (
    CISCOIETFPWMPLSMIB_CpwVcMplsNonTeMappingTable_CpwVcMplsNonTeMappingEntry_CpwVcMplsNonTeMappingTunnelDirection_outbound CISCOIETFPWMPLSMIB_CpwVcMplsNonTeMappingTable_CpwVcMplsNonTeMappingEntry_CpwVcMplsNonTeMappingTunnelDirection = "outbound"

    CISCOIETFPWMPLSMIB_CpwVcMplsNonTeMappingTable_CpwVcMplsNonTeMappingEntry_CpwVcMplsNonTeMappingTunnelDirection_inbound CISCOIETFPWMPLSMIB_CpwVcMplsNonTeMappingTable_CpwVcMplsNonTeMappingEntry_CpwVcMplsNonTeMappingTunnelDirection = "inbound"
)

// CISCOIETFPWMPLSMIB_CpwVcMplsTeMappingTable
// This table maps an inbound/outbound Tunnel to a VC in  
// MPLS-TE applications.
type CISCOIETFPWMPLSMIB_CpwVcMplsTeMappingTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A row in this table represents the association  between a PW VC and it's
    // MPLS-TE outer Tunnel.   An application can use this table to quickly
    // retrieve the   PW carried over specific TE MPLS outer tunnel.   The table
    // in indexed by the 4 indexes of a TE tunnel,  the direction of the VC
    // specific entry and the VcIndex.   The same table is used in both inbound
    // and outbound  directions, a different row for each direction. If the  
    // inbound association is not known, no rows should exist for   it.   Rows are
    // created by the local agent when all the   association data is available for
    // display. The type is slice of
    // CISCOIETFPWMPLSMIB_CpwVcMplsTeMappingTable_CpwVcMplsTeMappingEntry.
    CpwVcMplsTeMappingEntry []*CISCOIETFPWMPLSMIB_CpwVcMplsTeMappingTable_CpwVcMplsTeMappingEntry
}

func (cpwVcMplsTeMappingTable *CISCOIETFPWMPLSMIB_CpwVcMplsTeMappingTable) GetEntityData() *types.CommonEntityData {
    cpwVcMplsTeMappingTable.EntityData.YFilter = cpwVcMplsTeMappingTable.YFilter
    cpwVcMplsTeMappingTable.EntityData.YangName = "cpwVcMplsTeMappingTable"
    cpwVcMplsTeMappingTable.EntityData.BundleName = "cisco_ios_xe"
    cpwVcMplsTeMappingTable.EntityData.ParentYangName = "CISCO-IETF-PW-MPLS-MIB"
    cpwVcMplsTeMappingTable.EntityData.SegmentPath = "cpwVcMplsTeMappingTable"
    cpwVcMplsTeMappingTable.EntityData.AbsolutePath = "CISCO-IETF-PW-MPLS-MIB:CISCO-IETF-PW-MPLS-MIB/" + cpwVcMplsTeMappingTable.EntityData.SegmentPath
    cpwVcMplsTeMappingTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwVcMplsTeMappingTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwVcMplsTeMappingTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwVcMplsTeMappingTable.EntityData.Children = types.NewOrderedMap()
    cpwVcMplsTeMappingTable.EntityData.Children.Append("cpwVcMplsTeMappingEntry", types.YChild{"CpwVcMplsTeMappingEntry", nil})
    for i := range cpwVcMplsTeMappingTable.CpwVcMplsTeMappingEntry {
        cpwVcMplsTeMappingTable.EntityData.Children.Append(types.GetSegmentPath(cpwVcMplsTeMappingTable.CpwVcMplsTeMappingEntry[i]), types.YChild{"CpwVcMplsTeMappingEntry", cpwVcMplsTeMappingTable.CpwVcMplsTeMappingEntry[i]})
    }
    cpwVcMplsTeMappingTable.EntityData.Leafs = types.NewOrderedMap()

    cpwVcMplsTeMappingTable.EntityData.YListKeys = []string {}

    return &(cpwVcMplsTeMappingTable.EntityData)
}

// CISCOIETFPWMPLSMIB_CpwVcMplsTeMappingTable_CpwVcMplsTeMappingEntry
// A row in this table represents the association 
// between a PW VC and it's MPLS-TE outer Tunnel. 
// 
// An application can use this table to quickly retrieve the  
// PW carried over specific TE MPLS outer tunnel. 
// 
// The table in indexed by the 4 indexes of a TE tunnel, 
// the direction of the VC specific entry and the VcIndex. 
// 
// The same table is used in both inbound and outbound 
// directions, a different row for each direction. If the  
// inbound association is not known, no rows should exist for  
// it. 
// 
// Rows are created by the local agent when all the  
// association data is available for display.
type CISCOIETFPWMPLSMIB_CpwVcMplsTeMappingTable_CpwVcMplsTeMappingEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Identifies if the row represent an outbound or
    // inbound   mapping. The type is CpwVcMplsTeMappingTunnelDirection.
    CpwVcMplsTeMappingTunnelDirection interface{}

    // This attribute is a key. Primary index for the conceptual row identifying
    // the   MPLS-TE tunnel. The type is interface{} with range: 0..65535.
    CpwVcMplsTeMappingTunnelIndex interface{}

    // This attribute is a key. Identifies an instance of the MPLS-TE tunnel. The
    // type is interface{} with range: 0..4294967295.
    CpwVcMplsTeMappingTunnelInstance interface{}

    // This attribute is a key. Identifies an Peer LSR when the outer tunnel is
    // MPLS-TE   based. The type is string with length: 4.
    CpwVcMplsTeMappingTunnelPeerLsrID interface{}

    // This attribute is a key. Identifies the local LSR. The type is string with
    // length: 4.
    CpwVcMplsTeMappingTunnelLocalLsrID interface{}

    // This attribute is a key. The value that represent the VC in the cpwVcTable.
    // The type is interface{} with range: 0..4294967295.
    CpwVcMplsTeMappingVcIndex interface{}
}

func (cpwVcMplsTeMappingEntry *CISCOIETFPWMPLSMIB_CpwVcMplsTeMappingTable_CpwVcMplsTeMappingEntry) GetEntityData() *types.CommonEntityData {
    cpwVcMplsTeMappingEntry.EntityData.YFilter = cpwVcMplsTeMappingEntry.YFilter
    cpwVcMplsTeMappingEntry.EntityData.YangName = "cpwVcMplsTeMappingEntry"
    cpwVcMplsTeMappingEntry.EntityData.BundleName = "cisco_ios_xe"
    cpwVcMplsTeMappingEntry.EntityData.ParentYangName = "cpwVcMplsTeMappingTable"
    cpwVcMplsTeMappingEntry.EntityData.SegmentPath = "cpwVcMplsTeMappingEntry" + types.AddKeyToken(cpwVcMplsTeMappingEntry.CpwVcMplsTeMappingTunnelDirection, "cpwVcMplsTeMappingTunnelDirection") + types.AddKeyToken(cpwVcMplsTeMappingEntry.CpwVcMplsTeMappingTunnelIndex, "cpwVcMplsTeMappingTunnelIndex") + types.AddKeyToken(cpwVcMplsTeMappingEntry.CpwVcMplsTeMappingTunnelInstance, "cpwVcMplsTeMappingTunnelInstance") + types.AddKeyToken(cpwVcMplsTeMappingEntry.CpwVcMplsTeMappingTunnelPeerLsrID, "cpwVcMplsTeMappingTunnelPeerLsrID") + types.AddKeyToken(cpwVcMplsTeMappingEntry.CpwVcMplsTeMappingTunnelLocalLsrID, "cpwVcMplsTeMappingTunnelLocalLsrID") + types.AddKeyToken(cpwVcMplsTeMappingEntry.CpwVcMplsTeMappingVcIndex, "cpwVcMplsTeMappingVcIndex")
    cpwVcMplsTeMappingEntry.EntityData.AbsolutePath = "CISCO-IETF-PW-MPLS-MIB:CISCO-IETF-PW-MPLS-MIB/cpwVcMplsTeMappingTable/" + cpwVcMplsTeMappingEntry.EntityData.SegmentPath
    cpwVcMplsTeMappingEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwVcMplsTeMappingEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwVcMplsTeMappingEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwVcMplsTeMappingEntry.EntityData.Children = types.NewOrderedMap()
    cpwVcMplsTeMappingEntry.EntityData.Leafs = types.NewOrderedMap()
    cpwVcMplsTeMappingEntry.EntityData.Leafs.Append("cpwVcMplsTeMappingTunnelDirection", types.YLeaf{"CpwVcMplsTeMappingTunnelDirection", cpwVcMplsTeMappingEntry.CpwVcMplsTeMappingTunnelDirection})
    cpwVcMplsTeMappingEntry.EntityData.Leafs.Append("cpwVcMplsTeMappingTunnelIndex", types.YLeaf{"CpwVcMplsTeMappingTunnelIndex", cpwVcMplsTeMappingEntry.CpwVcMplsTeMappingTunnelIndex})
    cpwVcMplsTeMappingEntry.EntityData.Leafs.Append("cpwVcMplsTeMappingTunnelInstance", types.YLeaf{"CpwVcMplsTeMappingTunnelInstance", cpwVcMplsTeMappingEntry.CpwVcMplsTeMappingTunnelInstance})
    cpwVcMplsTeMappingEntry.EntityData.Leafs.Append("cpwVcMplsTeMappingTunnelPeerLsrID", types.YLeaf{"CpwVcMplsTeMappingTunnelPeerLsrID", cpwVcMplsTeMappingEntry.CpwVcMplsTeMappingTunnelPeerLsrID})
    cpwVcMplsTeMappingEntry.EntityData.Leafs.Append("cpwVcMplsTeMappingTunnelLocalLsrID", types.YLeaf{"CpwVcMplsTeMappingTunnelLocalLsrID", cpwVcMplsTeMappingEntry.CpwVcMplsTeMappingTunnelLocalLsrID})
    cpwVcMplsTeMappingEntry.EntityData.Leafs.Append("cpwVcMplsTeMappingVcIndex", types.YLeaf{"CpwVcMplsTeMappingVcIndex", cpwVcMplsTeMappingEntry.CpwVcMplsTeMappingVcIndex})

    cpwVcMplsTeMappingEntry.EntityData.YListKeys = []string {"CpwVcMplsTeMappingTunnelDirection", "CpwVcMplsTeMappingTunnelIndex", "CpwVcMplsTeMappingTunnelInstance", "CpwVcMplsTeMappingTunnelPeerLsrID", "CpwVcMplsTeMappingTunnelLocalLsrID", "CpwVcMplsTeMappingVcIndex"}

    return &(cpwVcMplsTeMappingEntry.EntityData)
}

// CISCOIETFPWMPLSMIB_CpwVcMplsTeMappingTable_CpwVcMplsTeMappingEntry_CpwVcMplsTeMappingTunnelDirection represents mapping.
type CISCOIETFPWMPLSMIB_CpwVcMplsTeMappingTable_CpwVcMplsTeMappingEntry_CpwVcMplsTeMappingTunnelDirection string

const (
    CISCOIETFPWMPLSMIB_CpwVcMplsTeMappingTable_CpwVcMplsTeMappingEntry_CpwVcMplsTeMappingTunnelDirection_outbound CISCOIETFPWMPLSMIB_CpwVcMplsTeMappingTable_CpwVcMplsTeMappingEntry_CpwVcMplsTeMappingTunnelDirection = "outbound"

    CISCOIETFPWMPLSMIB_CpwVcMplsTeMappingTable_CpwVcMplsTeMappingEntry_CpwVcMplsTeMappingTunnelDirection_inbound CISCOIETFPWMPLSMIB_CpwVcMplsTeMappingTable_CpwVcMplsTeMappingEntry_CpwVcMplsTeMappingTunnelDirection = "inbound"
)

