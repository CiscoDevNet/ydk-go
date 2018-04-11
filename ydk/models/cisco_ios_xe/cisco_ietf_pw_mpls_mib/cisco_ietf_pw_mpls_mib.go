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

    
    Cpwvcmplsobjects CISCOIETFPWMPLSMIB_Cpwvcmplsobjects

    // This table specifies information for VC to be carried over   MPLS PSN.
    Cpwvcmplstable CISCOIETFPWMPLSMIB_Cpwvcmplstable

    // This table associates VCs using MPLS PSN with the outbound  MPLS tunnels
    // (i.e. toward the PSN) or the physical   interface in case of VC only.
    Cpwvcmplsoutboundtable CISCOIETFPWMPLSMIB_Cpwvcmplsoutboundtable

    // This table associates VCs using MPLS PSN with the inbound  MPLS tunnels
    // (i.e. for packets coming from the PSN),   if such association is desired
    // (mainly for security   reasons).
    Cpwvcmplsinboundtable CISCOIETFPWMPLSMIB_Cpwvcmplsinboundtable

    // This table maps an inbound/outbound Tunnel to a VC in non-  TE
    // applications.
    Cpwvcmplsnontemappingtable CISCOIETFPWMPLSMIB_Cpwvcmplsnontemappingtable

    // This table maps an inbound/outbound Tunnel to a VC in   MPLS-TE
    // applications.
    Cpwvcmplstemappingtable CISCOIETFPWMPLSMIB_Cpwvcmplstemappingtable
}

func (cISCOIETFPWMPLSMIB *CISCOIETFPWMPLSMIB) GetEntityData() *types.CommonEntityData {
    cISCOIETFPWMPLSMIB.EntityData.YFilter = cISCOIETFPWMPLSMIB.YFilter
    cISCOIETFPWMPLSMIB.EntityData.YangName = "CISCO-IETF-PW-MPLS-MIB"
    cISCOIETFPWMPLSMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOIETFPWMPLSMIB.EntityData.ParentYangName = "CISCO-IETF-PW-MPLS-MIB"
    cISCOIETFPWMPLSMIB.EntityData.SegmentPath = "CISCO-IETF-PW-MPLS-MIB:CISCO-IETF-PW-MPLS-MIB"
    cISCOIETFPWMPLSMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOIETFPWMPLSMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOIETFPWMPLSMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOIETFPWMPLSMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOIETFPWMPLSMIB.EntityData.Children["cpwVcMplsObjects"] = types.YChild{"Cpwvcmplsobjects", &cISCOIETFPWMPLSMIB.Cpwvcmplsobjects}
    cISCOIETFPWMPLSMIB.EntityData.Children["cpwVcMplsTable"] = types.YChild{"Cpwvcmplstable", &cISCOIETFPWMPLSMIB.Cpwvcmplstable}
    cISCOIETFPWMPLSMIB.EntityData.Children["cpwVcMplsOutboundTable"] = types.YChild{"Cpwvcmplsoutboundtable", &cISCOIETFPWMPLSMIB.Cpwvcmplsoutboundtable}
    cISCOIETFPWMPLSMIB.EntityData.Children["cpwVcMplsInboundTable"] = types.YChild{"Cpwvcmplsinboundtable", &cISCOIETFPWMPLSMIB.Cpwvcmplsinboundtable}
    cISCOIETFPWMPLSMIB.EntityData.Children["cpwVcMplsNonTeMappingTable"] = types.YChild{"Cpwvcmplsnontemappingtable", &cISCOIETFPWMPLSMIB.Cpwvcmplsnontemappingtable}
    cISCOIETFPWMPLSMIB.EntityData.Children["cpwVcMplsTeMappingTable"] = types.YChild{"Cpwvcmplstemappingtable", &cISCOIETFPWMPLSMIB.Cpwvcmplstemappingtable}
    cISCOIETFPWMPLSMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOIETFPWMPLSMIB.EntityData)
}

// CISCOIETFPWMPLSMIB_Cpwvcmplsobjects
type CISCOIETFPWMPLSMIB_Cpwvcmplsobjects struct {
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
    Cpwvcmplsoutboundindexnext interface{}

    // This object contains an appropriate value to  be used for
    // cpwVcMplsInboundIndex when creating  entries in the cpwVcMplsInboundTable.
    // The value  0 indicates that no unassigned entries are  available. To obtain
    // the cpwVcMplsInboundIndex  value for a new entry, the manager issues a 
    // management protocol retrieval operation to obtain  the current value of
    // this object.  After each  retrieval, the agent should modify the value to 
    // the next unassigned index, however the agent MUST  NOT assume such
    // retrieval will be done for each   row created. The type is interface{} with
    // range: 0..4294967295.
    Cpwvcmplsinboundindexnext interface{}
}

func (cpwvcmplsobjects *CISCOIETFPWMPLSMIB_Cpwvcmplsobjects) GetEntityData() *types.CommonEntityData {
    cpwvcmplsobjects.EntityData.YFilter = cpwvcmplsobjects.YFilter
    cpwvcmplsobjects.EntityData.YangName = "cpwVcMplsObjects"
    cpwvcmplsobjects.EntityData.BundleName = "cisco_ios_xe"
    cpwvcmplsobjects.EntityData.ParentYangName = "CISCO-IETF-PW-MPLS-MIB"
    cpwvcmplsobjects.EntityData.SegmentPath = "cpwVcMplsObjects"
    cpwvcmplsobjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwvcmplsobjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwvcmplsobjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwvcmplsobjects.EntityData.Children = make(map[string]types.YChild)
    cpwvcmplsobjects.EntityData.Leafs = make(map[string]types.YLeaf)
    cpwvcmplsobjects.EntityData.Leafs["cpwVcMplsOutboundIndexNext"] = types.YLeaf{"Cpwvcmplsoutboundindexnext", cpwvcmplsobjects.Cpwvcmplsoutboundindexnext}
    cpwvcmplsobjects.EntityData.Leafs["cpwVcMplsInboundIndexNext"] = types.YLeaf{"Cpwvcmplsinboundindexnext", cpwvcmplsobjects.Cpwvcmplsinboundindexnext}
    return &(cpwvcmplsobjects.EntityData)
}

// CISCOIETFPWMPLSMIB_Cpwvcmplstable
// This table specifies information for VC to be carried over  
// MPLS PSN.
type CISCOIETFPWMPLSMIB_Cpwvcmplstable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A row in this table represents parameters specific to MPLS   PSN for a
    // pseudo wire connection (VC). The row is created   automatically by the
    // local agent if the cpwVcPsnType is   MPLS. It is indexed by cpwVcIndex,
    // which uniquely   identifying a singular connection. . The type is slice of
    // CISCOIETFPWMPLSMIB_Cpwvcmplstable_Cpwvcmplsentry.
    Cpwvcmplsentry []CISCOIETFPWMPLSMIB_Cpwvcmplstable_Cpwvcmplsentry
}

func (cpwvcmplstable *CISCOIETFPWMPLSMIB_Cpwvcmplstable) GetEntityData() *types.CommonEntityData {
    cpwvcmplstable.EntityData.YFilter = cpwvcmplstable.YFilter
    cpwvcmplstable.EntityData.YangName = "cpwVcMplsTable"
    cpwvcmplstable.EntityData.BundleName = "cisco_ios_xe"
    cpwvcmplstable.EntityData.ParentYangName = "CISCO-IETF-PW-MPLS-MIB"
    cpwvcmplstable.EntityData.SegmentPath = "cpwVcMplsTable"
    cpwvcmplstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwvcmplstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwvcmplstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwvcmplstable.EntityData.Children = make(map[string]types.YChild)
    cpwvcmplstable.EntityData.Children["cpwVcMplsEntry"] = types.YChild{"Cpwvcmplsentry", nil}
    for i := range cpwvcmplstable.Cpwvcmplsentry {
        cpwvcmplstable.EntityData.Children[types.GetSegmentPath(&cpwvcmplstable.Cpwvcmplsentry[i])] = types.YChild{"Cpwvcmplsentry", &cpwvcmplstable.Cpwvcmplsentry[i]}
    }
    cpwvcmplstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cpwvcmplstable.EntityData)
}

// CISCOIETFPWMPLSMIB_Cpwvcmplstable_Cpwvcmplsentry
// A row in this table represents parameters specific to MPLS  
// PSN for a pseudo wire connection (VC). The row is created  
// automatically by the local agent if the cpwVcPsnType is  
// MPLS. It is indexed by cpwVcIndex, which uniquely  
// identifying a singular connection. 
type CISCOIETFPWMPLSMIB_Cpwvcmplstable_Cpwvcmplsentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to cisco_ietf_pw_mib.CISCOIETFPWMIB_Cpwvctable_Cpwvcentry_Cpwvcindex
    Cpwvcindex interface{}

    // Set by the operator to indicate the outer tunnel types, if  exists. mplsTe
    // is used if the outer tunnel was set-up by   MPLS-TE, and mplsNonTe is used
    // the outer tunnel was set up  by LDP or manually. Combination of mplsTe and
    // mplsNonTe   may exist in case of outer tunnel protection.  vcOnly is used
    // if there is no outer tunnel label. vcOnly   cannot be combined with
    // mplsNonTe or mplsTe. The type is map[string]bool.
    Cpwvcmplsmplstype interface{}

    // Set by the operator to indicate the way the VC shim label  EXP bits are to
    // be determined. The value of outerTunnel(1)  is used where there is an outer
    // tunnel - cpwVcMplsMplsType   is mplsTe or mplsNonTe. Note that in this case
    // there is no  need to mark the VC label with the EXP bits since the VC  
    // label is not visible to the intermediate nodes.  If there is no outer
    // tunnel, specifiedValue(2) indicate   that the value is specified by
    // cpwVcMplsExpBits, and   serviceDependant(3) indicate that the EXP bits are
    // setup   based on a rule specified in the emulated service specific  
    // tables, for example when the EXP bits are a function of   802.1p marking
    // for Ethernet emulated service. The type is Cpwvcmplsexpbitsmode.
    Cpwvcmplsexpbitsmode interface{}

    // Set by the operator to indicate the MPLS EXP bits to be   used on the VC
    // shim label if cpwVcMplsExpBitsMode is    specifiedValue(2), zero otherwise.
    // The type is interface{} with range: 0..7.
    Cpwvcmplsexpbits interface{}

    // Set by the operator to indicate the VC TTL bits to be used  on the VC shim
    // label. The type is interface{} with range: 0..255.
    Cpwvcmplsttl interface{}

    // The local LDP identifier of the LDP entity creating  this VC in the local
    // node. As the VC labels are always  set from the per platform label space,
    // the last two octets   in the LDP ID MUST be always both zeros. The type is
    // string.
    Cpwvcmplslocalldpid interface{}

    // The local LDP Entity index of the LDP entity to be used   for this VC on
    // the local node. Should be set to all zeros   if not used. The type is
    // interface{} with range: 0..4294967295.
    Cpwvcmplslocalldpentityid interface{}

    // The peer LDP identifier as identified from the LDP   session. Should be
    // zero if not relevant or not known yet. The type is string.
    Cpwvcmplspeerldpid interface{}

    // This variable indicates the storage type for this row. The type is
    // StorageType.
    Cpwvcmplsstoragetype interface{}
}

func (cpwvcmplsentry *CISCOIETFPWMPLSMIB_Cpwvcmplstable_Cpwvcmplsentry) GetEntityData() *types.CommonEntityData {
    cpwvcmplsentry.EntityData.YFilter = cpwvcmplsentry.YFilter
    cpwvcmplsentry.EntityData.YangName = "cpwVcMplsEntry"
    cpwvcmplsentry.EntityData.BundleName = "cisco_ios_xe"
    cpwvcmplsentry.EntityData.ParentYangName = "cpwVcMplsTable"
    cpwvcmplsentry.EntityData.SegmentPath = "cpwVcMplsEntry" + "[cpwVcIndex='" + fmt.Sprintf("%v", cpwvcmplsentry.Cpwvcindex) + "']"
    cpwvcmplsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwvcmplsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwvcmplsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwvcmplsentry.EntityData.Children = make(map[string]types.YChild)
    cpwvcmplsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cpwvcmplsentry.EntityData.Leafs["cpwVcIndex"] = types.YLeaf{"Cpwvcindex", cpwvcmplsentry.Cpwvcindex}
    cpwvcmplsentry.EntityData.Leafs["cpwVcMplsMplsType"] = types.YLeaf{"Cpwvcmplsmplstype", cpwvcmplsentry.Cpwvcmplsmplstype}
    cpwvcmplsentry.EntityData.Leafs["cpwVcMplsExpBitsMode"] = types.YLeaf{"Cpwvcmplsexpbitsmode", cpwvcmplsentry.Cpwvcmplsexpbitsmode}
    cpwvcmplsentry.EntityData.Leafs["cpwVcMplsExpBits"] = types.YLeaf{"Cpwvcmplsexpbits", cpwvcmplsentry.Cpwvcmplsexpbits}
    cpwvcmplsentry.EntityData.Leafs["cpwVcMplsTtl"] = types.YLeaf{"Cpwvcmplsttl", cpwvcmplsentry.Cpwvcmplsttl}
    cpwvcmplsentry.EntityData.Leafs["cpwVcMplsLocalLdpID"] = types.YLeaf{"Cpwvcmplslocalldpid", cpwvcmplsentry.Cpwvcmplslocalldpid}
    cpwvcmplsentry.EntityData.Leafs["cpwVcMplsLocalLdpEntityID"] = types.YLeaf{"Cpwvcmplslocalldpentityid", cpwvcmplsentry.Cpwvcmplslocalldpentityid}
    cpwvcmplsentry.EntityData.Leafs["cpwVcMplsPeerLdpID"] = types.YLeaf{"Cpwvcmplspeerldpid", cpwvcmplsentry.Cpwvcmplspeerldpid}
    cpwvcmplsentry.EntityData.Leafs["cpwVcMplsStorageType"] = types.YLeaf{"Cpwvcmplsstoragetype", cpwvcmplsentry.Cpwvcmplsstoragetype}
    return &(cpwvcmplsentry.EntityData)
}

// CISCOIETFPWMPLSMIB_Cpwvcmplstable_Cpwvcmplsentry_Cpwvcmplsexpbitsmode represents 802.1p marking for Ethernet emulated service.
type CISCOIETFPWMPLSMIB_Cpwvcmplstable_Cpwvcmplsentry_Cpwvcmplsexpbitsmode string

const (
    CISCOIETFPWMPLSMIB_Cpwvcmplstable_Cpwvcmplsentry_Cpwvcmplsexpbitsmode_outerTunnel CISCOIETFPWMPLSMIB_Cpwvcmplstable_Cpwvcmplsentry_Cpwvcmplsexpbitsmode = "outerTunnel"

    CISCOIETFPWMPLSMIB_Cpwvcmplstable_Cpwvcmplsentry_Cpwvcmplsexpbitsmode_specifiedValue CISCOIETFPWMPLSMIB_Cpwvcmplstable_Cpwvcmplsentry_Cpwvcmplsexpbitsmode = "specifiedValue"

    CISCOIETFPWMPLSMIB_Cpwvcmplstable_Cpwvcmplsentry_Cpwvcmplsexpbitsmode_serviceDependant CISCOIETFPWMPLSMIB_Cpwvcmplstable_Cpwvcmplsentry_Cpwvcmplsexpbitsmode = "serviceDependant"
)

// CISCOIETFPWMPLSMIB_Cpwvcmplsoutboundtable
// This table associates VCs using MPLS PSN with the outbound 
// MPLS tunnels (i.e. toward the PSN) or the physical  
// interface in case of VC only.
type CISCOIETFPWMPLSMIB_Cpwvcmplsoutboundtable struct {
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
    // of CISCOIETFPWMPLSMIB_Cpwvcmplsoutboundtable_Cpwvcmplsoutboundentry.
    Cpwvcmplsoutboundentry []CISCOIETFPWMPLSMIB_Cpwvcmplsoutboundtable_Cpwvcmplsoutboundentry
}

func (cpwvcmplsoutboundtable *CISCOIETFPWMPLSMIB_Cpwvcmplsoutboundtable) GetEntityData() *types.CommonEntityData {
    cpwvcmplsoutboundtable.EntityData.YFilter = cpwvcmplsoutboundtable.YFilter
    cpwvcmplsoutboundtable.EntityData.YangName = "cpwVcMplsOutboundTable"
    cpwvcmplsoutboundtable.EntityData.BundleName = "cisco_ios_xe"
    cpwvcmplsoutboundtable.EntityData.ParentYangName = "CISCO-IETF-PW-MPLS-MIB"
    cpwvcmplsoutboundtable.EntityData.SegmentPath = "cpwVcMplsOutboundTable"
    cpwvcmplsoutboundtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwvcmplsoutboundtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwvcmplsoutboundtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwvcmplsoutboundtable.EntityData.Children = make(map[string]types.YChild)
    cpwvcmplsoutboundtable.EntityData.Children["cpwVcMplsOutboundEntry"] = types.YChild{"Cpwvcmplsoutboundentry", nil}
    for i := range cpwvcmplsoutboundtable.Cpwvcmplsoutboundentry {
        cpwvcmplsoutboundtable.EntityData.Children[types.GetSegmentPath(&cpwvcmplsoutboundtable.Cpwvcmplsoutboundentry[i])] = types.YChild{"Cpwvcmplsoutboundentry", &cpwvcmplsoutboundtable.Cpwvcmplsoutboundentry[i]}
    }
    cpwvcmplsoutboundtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cpwvcmplsoutboundtable.EntityData)
}

// CISCOIETFPWMPLSMIB_Cpwvcmplsoutboundtable_Cpwvcmplsoutboundentry
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
type CISCOIETFPWMPLSMIB_Cpwvcmplsoutboundtable_Cpwvcmplsoutboundentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to cisco_ietf_pw_mib.CISCOIETFPWMIB_Cpwvctable_Cpwvcentry_Cpwvcindex
    Cpwvcindex interface{}

    // This attribute is a key. Arbitrary index for enabling multiple rows per VC
    // in  this table. Next available free index can be retrieved   using
    // cpwVcMplsOutboundIndexNext. . The type is interface{} with range:
    // 0..4294967295.
    Cpwvcmplsoutboundindex interface{}

    // This object will be set by the operator. If the outer  label is defined in
    // the MPL-LSR-MIB, i.e. set by LDP  or manually, this object points to the XC
    // index   of the outer tunnel. Otherwise, it is set to zero. The type is
    // interface{} with range: 0..4294967295.
    Cpwvcmplsoutboundlsrxcindex interface{}

    // Part of set of indexes for outbound tunnel in the case of   MPLS-TE outer
    // tunnel, otherwise set to zero. The type is interface{} with range:
    // 0..65535.
    Cpwvcmplsoutboundtunnelindex interface{}

    // Part of set of indexes for outbound tunnel in the case of   MPLS-TE outer
    // tunnel, otherwise set to zero. The type is interface{} with range:
    // 0..4294967295.
    Cpwvcmplsoutboundtunnelinstance interface{}

    // Part of set of indexes for outbound tunnel in the case of   MPLS-TE outer
    // tunnel, otherwise set to zero. The type is string with length: 4.
    Cpwvcmplsoutboundtunnellcllsr interface{}

    // Part of set of indexes for outbound tunnel in the case of   MPLS-TE outer
    // tunnel, otherwise set to zero. The type is string with length: 4.
    Cpwvcmplsoutboundtunnelpeerlsr interface{}

    // In case of VC only (no outer tunnel), this object holds  the ifIndex of the
    // outbound port, otherwise set to zero. The type is interface{} with range:
    // 0..2147483647.
    Cpwvcmplsoutboundifindex interface{}

    // For creating, modifying, and deleting this row. The type is RowStatus.
    Cpwvcmplsoutboundrowstatus interface{}

    // This variable indicates the storage type for this object. The type is
    // StorageType.
    Cpwvcmplsoutboundstoragetype interface{}
}

func (cpwvcmplsoutboundentry *CISCOIETFPWMPLSMIB_Cpwvcmplsoutboundtable_Cpwvcmplsoutboundentry) GetEntityData() *types.CommonEntityData {
    cpwvcmplsoutboundentry.EntityData.YFilter = cpwvcmplsoutboundentry.YFilter
    cpwvcmplsoutboundentry.EntityData.YangName = "cpwVcMplsOutboundEntry"
    cpwvcmplsoutboundentry.EntityData.BundleName = "cisco_ios_xe"
    cpwvcmplsoutboundentry.EntityData.ParentYangName = "cpwVcMplsOutboundTable"
    cpwvcmplsoutboundentry.EntityData.SegmentPath = "cpwVcMplsOutboundEntry" + "[cpwVcIndex='" + fmt.Sprintf("%v", cpwvcmplsoutboundentry.Cpwvcindex) + "']" + "[cpwVcMplsOutboundIndex='" + fmt.Sprintf("%v", cpwvcmplsoutboundentry.Cpwvcmplsoutboundindex) + "']"
    cpwvcmplsoutboundentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwvcmplsoutboundentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwvcmplsoutboundentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwvcmplsoutboundentry.EntityData.Children = make(map[string]types.YChild)
    cpwvcmplsoutboundentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cpwvcmplsoutboundentry.EntityData.Leafs["cpwVcIndex"] = types.YLeaf{"Cpwvcindex", cpwvcmplsoutboundentry.Cpwvcindex}
    cpwvcmplsoutboundentry.EntityData.Leafs["cpwVcMplsOutboundIndex"] = types.YLeaf{"Cpwvcmplsoutboundindex", cpwvcmplsoutboundentry.Cpwvcmplsoutboundindex}
    cpwvcmplsoutboundentry.EntityData.Leafs["cpwVcMplsOutboundLsrXcIndex"] = types.YLeaf{"Cpwvcmplsoutboundlsrxcindex", cpwvcmplsoutboundentry.Cpwvcmplsoutboundlsrxcindex}
    cpwvcmplsoutboundentry.EntityData.Leafs["cpwVcMplsOutboundTunnelIndex"] = types.YLeaf{"Cpwvcmplsoutboundtunnelindex", cpwvcmplsoutboundentry.Cpwvcmplsoutboundtunnelindex}
    cpwvcmplsoutboundentry.EntityData.Leafs["cpwVcMplsOutboundTunnelInstance"] = types.YLeaf{"Cpwvcmplsoutboundtunnelinstance", cpwvcmplsoutboundentry.Cpwvcmplsoutboundtunnelinstance}
    cpwvcmplsoutboundentry.EntityData.Leafs["cpwVcMplsOutboundTunnelLclLSR"] = types.YLeaf{"Cpwvcmplsoutboundtunnellcllsr", cpwvcmplsoutboundentry.Cpwvcmplsoutboundtunnellcllsr}
    cpwvcmplsoutboundentry.EntityData.Leafs["cpwVcMplsOutboundTunnelPeerLSR"] = types.YLeaf{"Cpwvcmplsoutboundtunnelpeerlsr", cpwvcmplsoutboundentry.Cpwvcmplsoutboundtunnelpeerlsr}
    cpwvcmplsoutboundentry.EntityData.Leafs["cpwVcMplsOutboundIfIndex"] = types.YLeaf{"Cpwvcmplsoutboundifindex", cpwvcmplsoutboundentry.Cpwvcmplsoutboundifindex}
    cpwvcmplsoutboundentry.EntityData.Leafs["cpwVcMplsOutboundRowStatus"] = types.YLeaf{"Cpwvcmplsoutboundrowstatus", cpwvcmplsoutboundentry.Cpwvcmplsoutboundrowstatus}
    cpwvcmplsoutboundentry.EntityData.Leafs["cpwVcMplsOutboundStorageType"] = types.YLeaf{"Cpwvcmplsoutboundstoragetype", cpwvcmplsoutboundentry.Cpwvcmplsoutboundstoragetype}
    return &(cpwvcmplsoutboundentry.EntityData)
}

// CISCOIETFPWMPLSMIB_Cpwvcmplsinboundtable
// This table associates VCs using MPLS PSN with the inbound 
// MPLS tunnels (i.e. for packets coming from the PSN),  
// if such association is desired (mainly for security  
// reasons).
type CISCOIETFPWMPLSMIB_Cpwvcmplsinboundtable struct {
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
    // CISCOIETFPWMPLSMIB_Cpwvcmplsinboundtable_Cpwvcmplsinboundentry.
    Cpwvcmplsinboundentry []CISCOIETFPWMPLSMIB_Cpwvcmplsinboundtable_Cpwvcmplsinboundentry
}

func (cpwvcmplsinboundtable *CISCOIETFPWMPLSMIB_Cpwvcmplsinboundtable) GetEntityData() *types.CommonEntityData {
    cpwvcmplsinboundtable.EntityData.YFilter = cpwvcmplsinboundtable.YFilter
    cpwvcmplsinboundtable.EntityData.YangName = "cpwVcMplsInboundTable"
    cpwvcmplsinboundtable.EntityData.BundleName = "cisco_ios_xe"
    cpwvcmplsinboundtable.EntityData.ParentYangName = "CISCO-IETF-PW-MPLS-MIB"
    cpwvcmplsinboundtable.EntityData.SegmentPath = "cpwVcMplsInboundTable"
    cpwvcmplsinboundtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwvcmplsinboundtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwvcmplsinboundtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwvcmplsinboundtable.EntityData.Children = make(map[string]types.YChild)
    cpwvcmplsinboundtable.EntityData.Children["cpwVcMplsInboundEntry"] = types.YChild{"Cpwvcmplsinboundentry", nil}
    for i := range cpwvcmplsinboundtable.Cpwvcmplsinboundentry {
        cpwvcmplsinboundtable.EntityData.Children[types.GetSegmentPath(&cpwvcmplsinboundtable.Cpwvcmplsinboundentry[i])] = types.YChild{"Cpwvcmplsinboundentry", &cpwvcmplsinboundtable.Cpwvcmplsinboundentry[i]}
    }
    cpwvcmplsinboundtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cpwvcmplsinboundtable.EntityData)
}

// CISCOIETFPWMPLSMIB_Cpwvcmplsinboundtable_Cpwvcmplsinboundentry
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
type CISCOIETFPWMPLSMIB_Cpwvcmplsinboundtable_Cpwvcmplsinboundentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to cisco_ietf_pw_mib.CISCOIETFPWMIB_Cpwvctable_Cpwvcentry_Cpwvcindex
    Cpwvcindex interface{}

    // This attribute is a key. Arbitrary index for enabling multiple rows per VC
    // in  this table. Next available free index can be retrieved  using
    // cpwVcMplsInboundIndexNext. . The type is interface{} with range:
    // 0..4294967295.
    Cpwvcmplsinboundindex interface{}

    // If the outer label is defined in the MPL-LSR-MIB, i.e. set   by LDP or
    // manually, this object points to the XC index   of the outer tunnel.
    // Otherwise, it is set to zero. The type is interface{} with range:
    // 0..4294967295.
    Cpwvcmplsinboundlsrxcindex interface{}

    // Part of set of indexes for outbound tunnel in the case of   MPLS-TE outer
    // tunnel, otherwise set to zero. The type is interface{} with range:
    // 0..65535.
    Cpwvcmplsinboundtunnelindex interface{}

    // Part of set of indexes for outbound tunnel in the case of   MPLS-TE outer
    // tunnel, otherwise set to zero. The type is interface{} with range:
    // 0..4294967295.
    Cpwvcmplsinboundtunnelinstance interface{}

    // Part of set of indexes for outbound tunnel in the case of   MPLS-TE outer
    // tunnel, otherwise set to zero. The type is string with length: 4.
    Cpwvcmplsinboundtunnellcllsr interface{}

    // Part of set of indexes for outbound tunnel in the case of   MPLS-TE outer
    // tunnel, otherwise set to zero. The type is string with length: 4.
    Cpwvcmplsinboundtunnelpeerlsr interface{}

    // In case of VC only (no outer tunnel), this object holds the  ifIndex of the
    // inbound port, otherwise set to zero. The type is interface{} with range:
    // 0..2147483647.
    Cpwvcmplsinboundifindex interface{}

    // For creating, modifying, and deleting this row. The type is RowStatus.
    Cpwvcmplsinboundrowstatus interface{}

    // This variable indicates the storage type for this row. The type is
    // StorageType.
    Cpwvcmplsinboundstoragetype interface{}
}

func (cpwvcmplsinboundentry *CISCOIETFPWMPLSMIB_Cpwvcmplsinboundtable_Cpwvcmplsinboundentry) GetEntityData() *types.CommonEntityData {
    cpwvcmplsinboundentry.EntityData.YFilter = cpwvcmplsinboundentry.YFilter
    cpwvcmplsinboundentry.EntityData.YangName = "cpwVcMplsInboundEntry"
    cpwvcmplsinboundentry.EntityData.BundleName = "cisco_ios_xe"
    cpwvcmplsinboundentry.EntityData.ParentYangName = "cpwVcMplsInboundTable"
    cpwvcmplsinboundentry.EntityData.SegmentPath = "cpwVcMplsInboundEntry" + "[cpwVcIndex='" + fmt.Sprintf("%v", cpwvcmplsinboundentry.Cpwvcindex) + "']" + "[cpwVcMplsInboundIndex='" + fmt.Sprintf("%v", cpwvcmplsinboundentry.Cpwvcmplsinboundindex) + "']"
    cpwvcmplsinboundentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwvcmplsinboundentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwvcmplsinboundentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwvcmplsinboundentry.EntityData.Children = make(map[string]types.YChild)
    cpwvcmplsinboundentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cpwvcmplsinboundentry.EntityData.Leafs["cpwVcIndex"] = types.YLeaf{"Cpwvcindex", cpwvcmplsinboundentry.Cpwvcindex}
    cpwvcmplsinboundentry.EntityData.Leafs["cpwVcMplsInboundIndex"] = types.YLeaf{"Cpwvcmplsinboundindex", cpwvcmplsinboundentry.Cpwvcmplsinboundindex}
    cpwvcmplsinboundentry.EntityData.Leafs["cpwVcMplsInboundLsrXcIndex"] = types.YLeaf{"Cpwvcmplsinboundlsrxcindex", cpwvcmplsinboundentry.Cpwvcmplsinboundlsrxcindex}
    cpwvcmplsinboundentry.EntityData.Leafs["cpwVcMplsInboundTunnelIndex"] = types.YLeaf{"Cpwvcmplsinboundtunnelindex", cpwvcmplsinboundentry.Cpwvcmplsinboundtunnelindex}
    cpwvcmplsinboundentry.EntityData.Leafs["cpwVcMplsInboundTunnelInstance"] = types.YLeaf{"Cpwvcmplsinboundtunnelinstance", cpwvcmplsinboundentry.Cpwvcmplsinboundtunnelinstance}
    cpwvcmplsinboundentry.EntityData.Leafs["cpwVcMplsInboundTunnelLclLSR"] = types.YLeaf{"Cpwvcmplsinboundtunnellcllsr", cpwvcmplsinboundentry.Cpwvcmplsinboundtunnellcllsr}
    cpwvcmplsinboundentry.EntityData.Leafs["cpwVcMplsInboundTunnelPeerLSR"] = types.YLeaf{"Cpwvcmplsinboundtunnelpeerlsr", cpwvcmplsinboundentry.Cpwvcmplsinboundtunnelpeerlsr}
    cpwvcmplsinboundentry.EntityData.Leafs["cpwVcMplsInboundIfIndex"] = types.YLeaf{"Cpwvcmplsinboundifindex", cpwvcmplsinboundentry.Cpwvcmplsinboundifindex}
    cpwvcmplsinboundentry.EntityData.Leafs["cpwVcMplsInboundRowStatus"] = types.YLeaf{"Cpwvcmplsinboundrowstatus", cpwvcmplsinboundentry.Cpwvcmplsinboundrowstatus}
    cpwvcmplsinboundentry.EntityData.Leafs["cpwVcMplsInboundStorageType"] = types.YLeaf{"Cpwvcmplsinboundstoragetype", cpwvcmplsinboundentry.Cpwvcmplsinboundstoragetype}
    return &(cpwvcmplsinboundentry.EntityData)
}

// CISCOIETFPWMPLSMIB_Cpwvcmplsnontemappingtable
// This table maps an inbound/outbound Tunnel to a VC in non- 
// TE applications.
type CISCOIETFPWMPLSMIB_Cpwvcmplsnontemappingtable struct {
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
    // CISCOIETFPWMPLSMIB_Cpwvcmplsnontemappingtable_Cpwvcmplsnontemappingentry.
    Cpwvcmplsnontemappingentry []CISCOIETFPWMPLSMIB_Cpwvcmplsnontemappingtable_Cpwvcmplsnontemappingentry
}

func (cpwvcmplsnontemappingtable *CISCOIETFPWMPLSMIB_Cpwvcmplsnontemappingtable) GetEntityData() *types.CommonEntityData {
    cpwvcmplsnontemappingtable.EntityData.YFilter = cpwvcmplsnontemappingtable.YFilter
    cpwvcmplsnontemappingtable.EntityData.YangName = "cpwVcMplsNonTeMappingTable"
    cpwvcmplsnontemappingtable.EntityData.BundleName = "cisco_ios_xe"
    cpwvcmplsnontemappingtable.EntityData.ParentYangName = "CISCO-IETF-PW-MPLS-MIB"
    cpwvcmplsnontemappingtable.EntityData.SegmentPath = "cpwVcMplsNonTeMappingTable"
    cpwvcmplsnontemappingtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwvcmplsnontemappingtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwvcmplsnontemappingtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwvcmplsnontemappingtable.EntityData.Children = make(map[string]types.YChild)
    cpwvcmplsnontemappingtable.EntityData.Children["cpwVcMplsNonTeMappingEntry"] = types.YChild{"Cpwvcmplsnontemappingentry", nil}
    for i := range cpwvcmplsnontemappingtable.Cpwvcmplsnontemappingentry {
        cpwvcmplsnontemappingtable.EntityData.Children[types.GetSegmentPath(&cpwvcmplsnontemappingtable.Cpwvcmplsnontemappingentry[i])] = types.YChild{"Cpwvcmplsnontemappingentry", &cpwvcmplsnontemappingtable.Cpwvcmplsnontemappingentry[i]}
    }
    cpwvcmplsnontemappingtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cpwvcmplsnontemappingtable.EntityData)
}

// CISCOIETFPWMPLSMIB_Cpwvcmplsnontemappingtable_Cpwvcmplsnontemappingentry
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
type CISCOIETFPWMPLSMIB_Cpwvcmplsnontemappingtable_Cpwvcmplsnontemappingentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Identifies if the row represent an outbound or
    // inbound   mapping. The type is Cpwvcmplsnontemappingtunneldirection.
    Cpwvcmplsnontemappingtunneldirection interface{}

    // This attribute is a key. Index for the conceptual XC row identifying Tunnel
    // to VC   mappings when the outer tunnel is created by the MPLS-LSR-  MIB,
    // Zero otherwise. The type is interface{} with range: 0..4294967295.
    Cpwvcmplsnontemappingxctunnelindex interface{}

    // This attribute is a key. Identify the port on which the VC is carried for
    // VC only   case. The type is interface{} with range: 0..2147483647.
    Cpwvcmplsnontemappingifindex interface{}

    // This attribute is a key. The value that represent the VC in the cpwVcTable.
    // The type is interface{} with range: 0..4294967295.
    Cpwvcmplsnontemappingvcindex interface{}
}

func (cpwvcmplsnontemappingentry *CISCOIETFPWMPLSMIB_Cpwvcmplsnontemappingtable_Cpwvcmplsnontemappingentry) GetEntityData() *types.CommonEntityData {
    cpwvcmplsnontemappingentry.EntityData.YFilter = cpwvcmplsnontemappingentry.YFilter
    cpwvcmplsnontemappingentry.EntityData.YangName = "cpwVcMplsNonTeMappingEntry"
    cpwvcmplsnontemappingentry.EntityData.BundleName = "cisco_ios_xe"
    cpwvcmplsnontemappingentry.EntityData.ParentYangName = "cpwVcMplsNonTeMappingTable"
    cpwvcmplsnontemappingentry.EntityData.SegmentPath = "cpwVcMplsNonTeMappingEntry" + "[cpwVcMplsNonTeMappingTunnelDirection='" + fmt.Sprintf("%v", cpwvcmplsnontemappingentry.Cpwvcmplsnontemappingtunneldirection) + "']" + "[cpwVcMplsNonTeMappingXcTunnelIndex='" + fmt.Sprintf("%v", cpwvcmplsnontemappingentry.Cpwvcmplsnontemappingxctunnelindex) + "']" + "[cpwVcMplsNonTeMappingIfIndex='" + fmt.Sprintf("%v", cpwvcmplsnontemappingentry.Cpwvcmplsnontemappingifindex) + "']" + "[cpwVcMplsNonTeMappingVcIndex='" + fmt.Sprintf("%v", cpwvcmplsnontemappingentry.Cpwvcmplsnontemappingvcindex) + "']"
    cpwvcmplsnontemappingentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwvcmplsnontemappingentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwvcmplsnontemappingentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwvcmplsnontemappingentry.EntityData.Children = make(map[string]types.YChild)
    cpwvcmplsnontemappingentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cpwvcmplsnontemappingentry.EntityData.Leafs["cpwVcMplsNonTeMappingTunnelDirection"] = types.YLeaf{"Cpwvcmplsnontemappingtunneldirection", cpwvcmplsnontemappingentry.Cpwvcmplsnontemappingtunneldirection}
    cpwvcmplsnontemappingentry.EntityData.Leafs["cpwVcMplsNonTeMappingXcTunnelIndex"] = types.YLeaf{"Cpwvcmplsnontemappingxctunnelindex", cpwvcmplsnontemappingentry.Cpwvcmplsnontemappingxctunnelindex}
    cpwvcmplsnontemappingentry.EntityData.Leafs["cpwVcMplsNonTeMappingIfIndex"] = types.YLeaf{"Cpwvcmplsnontemappingifindex", cpwvcmplsnontemappingentry.Cpwvcmplsnontemappingifindex}
    cpwvcmplsnontemappingentry.EntityData.Leafs["cpwVcMplsNonTeMappingVcIndex"] = types.YLeaf{"Cpwvcmplsnontemappingvcindex", cpwvcmplsnontemappingentry.Cpwvcmplsnontemappingvcindex}
    return &(cpwvcmplsnontemappingentry.EntityData)
}

// CISCOIETFPWMPLSMIB_Cpwvcmplsnontemappingtable_Cpwvcmplsnontemappingentry_Cpwvcmplsnontemappingtunneldirection represents mapping.
type CISCOIETFPWMPLSMIB_Cpwvcmplsnontemappingtable_Cpwvcmplsnontemappingentry_Cpwvcmplsnontemappingtunneldirection string

const (
    CISCOIETFPWMPLSMIB_Cpwvcmplsnontemappingtable_Cpwvcmplsnontemappingentry_Cpwvcmplsnontemappingtunneldirection_outbound CISCOIETFPWMPLSMIB_Cpwvcmplsnontemappingtable_Cpwvcmplsnontemappingentry_Cpwvcmplsnontemappingtunneldirection = "outbound"

    CISCOIETFPWMPLSMIB_Cpwvcmplsnontemappingtable_Cpwvcmplsnontemappingentry_Cpwvcmplsnontemappingtunneldirection_inbound CISCOIETFPWMPLSMIB_Cpwvcmplsnontemappingtable_Cpwvcmplsnontemappingentry_Cpwvcmplsnontemappingtunneldirection = "inbound"
)

// CISCOIETFPWMPLSMIB_Cpwvcmplstemappingtable
// This table maps an inbound/outbound Tunnel to a VC in  
// MPLS-TE applications.
type CISCOIETFPWMPLSMIB_Cpwvcmplstemappingtable struct {
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
    // CISCOIETFPWMPLSMIB_Cpwvcmplstemappingtable_Cpwvcmplstemappingentry.
    Cpwvcmplstemappingentry []CISCOIETFPWMPLSMIB_Cpwvcmplstemappingtable_Cpwvcmplstemappingentry
}

func (cpwvcmplstemappingtable *CISCOIETFPWMPLSMIB_Cpwvcmplstemappingtable) GetEntityData() *types.CommonEntityData {
    cpwvcmplstemappingtable.EntityData.YFilter = cpwvcmplstemappingtable.YFilter
    cpwvcmplstemappingtable.EntityData.YangName = "cpwVcMplsTeMappingTable"
    cpwvcmplstemappingtable.EntityData.BundleName = "cisco_ios_xe"
    cpwvcmplstemappingtable.EntityData.ParentYangName = "CISCO-IETF-PW-MPLS-MIB"
    cpwvcmplstemappingtable.EntityData.SegmentPath = "cpwVcMplsTeMappingTable"
    cpwvcmplstemappingtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwvcmplstemappingtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwvcmplstemappingtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwvcmplstemappingtable.EntityData.Children = make(map[string]types.YChild)
    cpwvcmplstemappingtable.EntityData.Children["cpwVcMplsTeMappingEntry"] = types.YChild{"Cpwvcmplstemappingentry", nil}
    for i := range cpwvcmplstemappingtable.Cpwvcmplstemappingentry {
        cpwvcmplstemappingtable.EntityData.Children[types.GetSegmentPath(&cpwvcmplstemappingtable.Cpwvcmplstemappingentry[i])] = types.YChild{"Cpwvcmplstemappingentry", &cpwvcmplstemappingtable.Cpwvcmplstemappingentry[i]}
    }
    cpwvcmplstemappingtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cpwvcmplstemappingtable.EntityData)
}

// CISCOIETFPWMPLSMIB_Cpwvcmplstemappingtable_Cpwvcmplstemappingentry
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
type CISCOIETFPWMPLSMIB_Cpwvcmplstemappingtable_Cpwvcmplstemappingentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Identifies if the row represent an outbound or
    // inbound   mapping. The type is Cpwvcmplstemappingtunneldirection.
    Cpwvcmplstemappingtunneldirection interface{}

    // This attribute is a key. Primary index for the conceptual row identifying
    // the   MPLS-TE tunnel. The type is interface{} with range: 0..65535.
    Cpwvcmplstemappingtunnelindex interface{}

    // This attribute is a key. Identifies an instance of the MPLS-TE tunnel. The
    // type is interface{} with range: 0..4294967295.
    Cpwvcmplstemappingtunnelinstance interface{}

    // This attribute is a key. Identifies an Peer LSR when the outer tunnel is
    // MPLS-TE   based. The type is string with length: 4.
    Cpwvcmplstemappingtunnelpeerlsrid interface{}

    // This attribute is a key. Identifies the local LSR. The type is string with
    // length: 4.
    Cpwvcmplstemappingtunnellocallsrid interface{}

    // This attribute is a key. The value that represent the VC in the cpwVcTable.
    // The type is interface{} with range: 0..4294967295.
    Cpwvcmplstemappingvcindex interface{}
}

func (cpwvcmplstemappingentry *CISCOIETFPWMPLSMIB_Cpwvcmplstemappingtable_Cpwvcmplstemappingentry) GetEntityData() *types.CommonEntityData {
    cpwvcmplstemappingentry.EntityData.YFilter = cpwvcmplstemappingentry.YFilter
    cpwvcmplstemappingentry.EntityData.YangName = "cpwVcMplsTeMappingEntry"
    cpwvcmplstemappingentry.EntityData.BundleName = "cisco_ios_xe"
    cpwvcmplstemappingentry.EntityData.ParentYangName = "cpwVcMplsTeMappingTable"
    cpwvcmplstemappingentry.EntityData.SegmentPath = "cpwVcMplsTeMappingEntry" + "[cpwVcMplsTeMappingTunnelDirection='" + fmt.Sprintf("%v", cpwvcmplstemappingentry.Cpwvcmplstemappingtunneldirection) + "']" + "[cpwVcMplsTeMappingTunnelIndex='" + fmt.Sprintf("%v", cpwvcmplstemappingentry.Cpwvcmplstemappingtunnelindex) + "']" + "[cpwVcMplsTeMappingTunnelInstance='" + fmt.Sprintf("%v", cpwvcmplstemappingentry.Cpwvcmplstemappingtunnelinstance) + "']" + "[cpwVcMplsTeMappingTunnelPeerLsrID='" + fmt.Sprintf("%v", cpwvcmplstemappingentry.Cpwvcmplstemappingtunnelpeerlsrid) + "']" + "[cpwVcMplsTeMappingTunnelLocalLsrID='" + fmt.Sprintf("%v", cpwvcmplstemappingentry.Cpwvcmplstemappingtunnellocallsrid) + "']" + "[cpwVcMplsTeMappingVcIndex='" + fmt.Sprintf("%v", cpwvcmplstemappingentry.Cpwvcmplstemappingvcindex) + "']"
    cpwvcmplstemappingentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwvcmplstemappingentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwvcmplstemappingentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwvcmplstemappingentry.EntityData.Children = make(map[string]types.YChild)
    cpwvcmplstemappingentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cpwvcmplstemappingentry.EntityData.Leafs["cpwVcMplsTeMappingTunnelDirection"] = types.YLeaf{"Cpwvcmplstemappingtunneldirection", cpwvcmplstemappingentry.Cpwvcmplstemappingtunneldirection}
    cpwvcmplstemappingentry.EntityData.Leafs["cpwVcMplsTeMappingTunnelIndex"] = types.YLeaf{"Cpwvcmplstemappingtunnelindex", cpwvcmplstemappingentry.Cpwvcmplstemappingtunnelindex}
    cpwvcmplstemappingentry.EntityData.Leafs["cpwVcMplsTeMappingTunnelInstance"] = types.YLeaf{"Cpwvcmplstemappingtunnelinstance", cpwvcmplstemappingentry.Cpwvcmplstemappingtunnelinstance}
    cpwvcmplstemappingentry.EntityData.Leafs["cpwVcMplsTeMappingTunnelPeerLsrID"] = types.YLeaf{"Cpwvcmplstemappingtunnelpeerlsrid", cpwvcmplstemappingentry.Cpwvcmplstemappingtunnelpeerlsrid}
    cpwvcmplstemappingentry.EntityData.Leafs["cpwVcMplsTeMappingTunnelLocalLsrID"] = types.YLeaf{"Cpwvcmplstemappingtunnellocallsrid", cpwvcmplstemappingentry.Cpwvcmplstemappingtunnellocallsrid}
    cpwvcmplstemappingentry.EntityData.Leafs["cpwVcMplsTeMappingVcIndex"] = types.YLeaf{"Cpwvcmplstemappingvcindex", cpwvcmplstemappingentry.Cpwvcmplstemappingvcindex}
    return &(cpwvcmplstemappingentry.EntityData)
}

// CISCOIETFPWMPLSMIB_Cpwvcmplstemappingtable_Cpwvcmplstemappingentry_Cpwvcmplstemappingtunneldirection represents mapping.
type CISCOIETFPWMPLSMIB_Cpwvcmplstemappingtable_Cpwvcmplstemappingentry_Cpwvcmplstemappingtunneldirection string

const (
    CISCOIETFPWMPLSMIB_Cpwvcmplstemappingtable_Cpwvcmplstemappingentry_Cpwvcmplstemappingtunneldirection_outbound CISCOIETFPWMPLSMIB_Cpwvcmplstemappingtable_Cpwvcmplstemappingentry_Cpwvcmplstemappingtunneldirection = "outbound"

    CISCOIETFPWMPLSMIB_Cpwvcmplstemappingtable_Cpwvcmplstemappingentry_Cpwvcmplstemappingtunneldirection_inbound CISCOIETFPWMPLSMIB_Cpwvcmplstemappingtable_Cpwvcmplstemappingentry_Cpwvcmplstemappingtunneldirection = "inbound"
)

