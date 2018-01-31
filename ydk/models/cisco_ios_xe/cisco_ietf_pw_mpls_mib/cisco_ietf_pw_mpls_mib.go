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
    parent types.Entity
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

func (cISCOIETFPWMPLSMIB *CISCOIETFPWMPLSMIB) GetFilter() yfilter.YFilter { return cISCOIETFPWMPLSMIB.YFilter }

func (cISCOIETFPWMPLSMIB *CISCOIETFPWMPLSMIB) SetFilter(yf yfilter.YFilter) { cISCOIETFPWMPLSMIB.YFilter = yf }

func (cISCOIETFPWMPLSMIB *CISCOIETFPWMPLSMIB) GetGoName(yname string) string {
    if yname == "cpwVcMplsObjects" { return "Cpwvcmplsobjects" }
    if yname == "cpwVcMplsTable" { return "Cpwvcmplstable" }
    if yname == "cpwVcMplsOutboundTable" { return "Cpwvcmplsoutboundtable" }
    if yname == "cpwVcMplsInboundTable" { return "Cpwvcmplsinboundtable" }
    if yname == "cpwVcMplsNonTeMappingTable" { return "Cpwvcmplsnontemappingtable" }
    if yname == "cpwVcMplsTeMappingTable" { return "Cpwvcmplstemappingtable" }
    return ""
}

func (cISCOIETFPWMPLSMIB *CISCOIETFPWMPLSMIB) GetSegmentPath() string {
    return "CISCO-IETF-PW-MPLS-MIB:CISCO-IETF-PW-MPLS-MIB"
}

func (cISCOIETFPWMPLSMIB *CISCOIETFPWMPLSMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cpwVcMplsObjects" {
        return &cISCOIETFPWMPLSMIB.Cpwvcmplsobjects
    }
    if childYangName == "cpwVcMplsTable" {
        return &cISCOIETFPWMPLSMIB.Cpwvcmplstable
    }
    if childYangName == "cpwVcMplsOutboundTable" {
        return &cISCOIETFPWMPLSMIB.Cpwvcmplsoutboundtable
    }
    if childYangName == "cpwVcMplsInboundTable" {
        return &cISCOIETFPWMPLSMIB.Cpwvcmplsinboundtable
    }
    if childYangName == "cpwVcMplsNonTeMappingTable" {
        return &cISCOIETFPWMPLSMIB.Cpwvcmplsnontemappingtable
    }
    if childYangName == "cpwVcMplsTeMappingTable" {
        return &cISCOIETFPWMPLSMIB.Cpwvcmplstemappingtable
    }
    return nil
}

func (cISCOIETFPWMPLSMIB *CISCOIETFPWMPLSMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cpwVcMplsObjects"] = &cISCOIETFPWMPLSMIB.Cpwvcmplsobjects
    children["cpwVcMplsTable"] = &cISCOIETFPWMPLSMIB.Cpwvcmplstable
    children["cpwVcMplsOutboundTable"] = &cISCOIETFPWMPLSMIB.Cpwvcmplsoutboundtable
    children["cpwVcMplsInboundTable"] = &cISCOIETFPWMPLSMIB.Cpwvcmplsinboundtable
    children["cpwVcMplsNonTeMappingTable"] = &cISCOIETFPWMPLSMIB.Cpwvcmplsnontemappingtable
    children["cpwVcMplsTeMappingTable"] = &cISCOIETFPWMPLSMIB.Cpwvcmplstemappingtable
    return children
}

func (cISCOIETFPWMPLSMIB *CISCOIETFPWMPLSMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOIETFPWMPLSMIB *CISCOIETFPWMPLSMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOIETFPWMPLSMIB *CISCOIETFPWMPLSMIB) GetYangName() string { return "CISCO-IETF-PW-MPLS-MIB" }

func (cISCOIETFPWMPLSMIB *CISCOIETFPWMPLSMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOIETFPWMPLSMIB *CISCOIETFPWMPLSMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOIETFPWMPLSMIB *CISCOIETFPWMPLSMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOIETFPWMPLSMIB *CISCOIETFPWMPLSMIB) SetParent(parent types.Entity) { cISCOIETFPWMPLSMIB.parent = parent }

func (cISCOIETFPWMPLSMIB *CISCOIETFPWMPLSMIB) GetParent() types.Entity { return cISCOIETFPWMPLSMIB.parent }

func (cISCOIETFPWMPLSMIB *CISCOIETFPWMPLSMIB) GetParentYangName() string { return "CISCO-IETF-PW-MPLS-MIB" }

// CISCOIETFPWMPLSMIB_Cpwvcmplsobjects
type CISCOIETFPWMPLSMIB_Cpwvcmplsobjects struct {
    parent types.Entity
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

func (cpwvcmplsobjects *CISCOIETFPWMPLSMIB_Cpwvcmplsobjects) GetFilter() yfilter.YFilter { return cpwvcmplsobjects.YFilter }

func (cpwvcmplsobjects *CISCOIETFPWMPLSMIB_Cpwvcmplsobjects) SetFilter(yf yfilter.YFilter) { cpwvcmplsobjects.YFilter = yf }

func (cpwvcmplsobjects *CISCOIETFPWMPLSMIB_Cpwvcmplsobjects) GetGoName(yname string) string {
    if yname == "cpwVcMplsOutboundIndexNext" { return "Cpwvcmplsoutboundindexnext" }
    if yname == "cpwVcMplsInboundIndexNext" { return "Cpwvcmplsinboundindexnext" }
    return ""
}

func (cpwvcmplsobjects *CISCOIETFPWMPLSMIB_Cpwvcmplsobjects) GetSegmentPath() string {
    return "cpwVcMplsObjects"
}

func (cpwvcmplsobjects *CISCOIETFPWMPLSMIB_Cpwvcmplsobjects) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cpwvcmplsobjects *CISCOIETFPWMPLSMIB_Cpwvcmplsobjects) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cpwvcmplsobjects *CISCOIETFPWMPLSMIB_Cpwvcmplsobjects) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cpwVcMplsOutboundIndexNext"] = cpwvcmplsobjects.Cpwvcmplsoutboundindexnext
    leafs["cpwVcMplsInboundIndexNext"] = cpwvcmplsobjects.Cpwvcmplsinboundindexnext
    return leafs
}

func (cpwvcmplsobjects *CISCOIETFPWMPLSMIB_Cpwvcmplsobjects) GetBundleName() string { return "cisco_ios_xe" }

func (cpwvcmplsobjects *CISCOIETFPWMPLSMIB_Cpwvcmplsobjects) GetYangName() string { return "cpwVcMplsObjects" }

func (cpwvcmplsobjects *CISCOIETFPWMPLSMIB_Cpwvcmplsobjects) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpwvcmplsobjects *CISCOIETFPWMPLSMIB_Cpwvcmplsobjects) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpwvcmplsobjects *CISCOIETFPWMPLSMIB_Cpwvcmplsobjects) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpwvcmplsobjects *CISCOIETFPWMPLSMIB_Cpwvcmplsobjects) SetParent(parent types.Entity) { cpwvcmplsobjects.parent = parent }

func (cpwvcmplsobjects *CISCOIETFPWMPLSMIB_Cpwvcmplsobjects) GetParent() types.Entity { return cpwvcmplsobjects.parent }

func (cpwvcmplsobjects *CISCOIETFPWMPLSMIB_Cpwvcmplsobjects) GetParentYangName() string { return "CISCO-IETF-PW-MPLS-MIB" }

// CISCOIETFPWMPLSMIB_Cpwvcmplstable
// This table specifies information for VC to be carried over  
// MPLS PSN.
type CISCOIETFPWMPLSMIB_Cpwvcmplstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A row in this table represents parameters specific to MPLS   PSN for a
    // pseudo wire connection (VC). The row is created   automatically by the
    // local agent if the cpwVcPsnType is   MPLS. It is indexed by cpwVcIndex,
    // which uniquely   identifying a singular connection. . The type is slice of
    // CISCOIETFPWMPLSMIB_Cpwvcmplstable_Cpwvcmplsentry.
    Cpwvcmplsentry []CISCOIETFPWMPLSMIB_Cpwvcmplstable_Cpwvcmplsentry
}

func (cpwvcmplstable *CISCOIETFPWMPLSMIB_Cpwvcmplstable) GetFilter() yfilter.YFilter { return cpwvcmplstable.YFilter }

func (cpwvcmplstable *CISCOIETFPWMPLSMIB_Cpwvcmplstable) SetFilter(yf yfilter.YFilter) { cpwvcmplstable.YFilter = yf }

func (cpwvcmplstable *CISCOIETFPWMPLSMIB_Cpwvcmplstable) GetGoName(yname string) string {
    if yname == "cpwVcMplsEntry" { return "Cpwvcmplsentry" }
    return ""
}

func (cpwvcmplstable *CISCOIETFPWMPLSMIB_Cpwvcmplstable) GetSegmentPath() string {
    return "cpwVcMplsTable"
}

func (cpwvcmplstable *CISCOIETFPWMPLSMIB_Cpwvcmplstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cpwVcMplsEntry" {
        for _, c := range cpwvcmplstable.Cpwvcmplsentry {
            if cpwvcmplstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIETFPWMPLSMIB_Cpwvcmplstable_Cpwvcmplsentry{}
        cpwvcmplstable.Cpwvcmplsentry = append(cpwvcmplstable.Cpwvcmplsentry, child)
        return &cpwvcmplstable.Cpwvcmplsentry[len(cpwvcmplstable.Cpwvcmplsentry)-1]
    }
    return nil
}

func (cpwvcmplstable *CISCOIETFPWMPLSMIB_Cpwvcmplstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cpwvcmplstable.Cpwvcmplsentry {
        children[cpwvcmplstable.Cpwvcmplsentry[i].GetSegmentPath()] = &cpwvcmplstable.Cpwvcmplsentry[i]
    }
    return children
}

func (cpwvcmplstable *CISCOIETFPWMPLSMIB_Cpwvcmplstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cpwvcmplstable *CISCOIETFPWMPLSMIB_Cpwvcmplstable) GetBundleName() string { return "cisco_ios_xe" }

func (cpwvcmplstable *CISCOIETFPWMPLSMIB_Cpwvcmplstable) GetYangName() string { return "cpwVcMplsTable" }

func (cpwvcmplstable *CISCOIETFPWMPLSMIB_Cpwvcmplstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpwvcmplstable *CISCOIETFPWMPLSMIB_Cpwvcmplstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpwvcmplstable *CISCOIETFPWMPLSMIB_Cpwvcmplstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpwvcmplstable *CISCOIETFPWMPLSMIB_Cpwvcmplstable) SetParent(parent types.Entity) { cpwvcmplstable.parent = parent }

func (cpwvcmplstable *CISCOIETFPWMPLSMIB_Cpwvcmplstable) GetParent() types.Entity { return cpwvcmplstable.parent }

func (cpwvcmplstable *CISCOIETFPWMPLSMIB_Cpwvcmplstable) GetParentYangName() string { return "CISCO-IETF-PW-MPLS-MIB" }

// CISCOIETFPWMPLSMIB_Cpwvcmplstable_Cpwvcmplsentry
// A row in this table represents parameters specific to MPLS  
// PSN for a pseudo wire connection (VC). The row is created  
// automatically by the local agent if the cpwVcPsnType is  
// MPLS. It is indexed by cpwVcIndex, which uniquely  
// identifying a singular connection. 
type CISCOIETFPWMPLSMIB_Cpwvcmplstable_Cpwvcmplsentry struct {
    parent types.Entity
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

func (cpwvcmplsentry *CISCOIETFPWMPLSMIB_Cpwvcmplstable_Cpwvcmplsentry) GetFilter() yfilter.YFilter { return cpwvcmplsentry.YFilter }

func (cpwvcmplsentry *CISCOIETFPWMPLSMIB_Cpwvcmplstable_Cpwvcmplsentry) SetFilter(yf yfilter.YFilter) { cpwvcmplsentry.YFilter = yf }

func (cpwvcmplsentry *CISCOIETFPWMPLSMIB_Cpwvcmplstable_Cpwvcmplsentry) GetGoName(yname string) string {
    if yname == "cpwVcIndex" { return "Cpwvcindex" }
    if yname == "cpwVcMplsMplsType" { return "Cpwvcmplsmplstype" }
    if yname == "cpwVcMplsExpBitsMode" { return "Cpwvcmplsexpbitsmode" }
    if yname == "cpwVcMplsExpBits" { return "Cpwvcmplsexpbits" }
    if yname == "cpwVcMplsTtl" { return "Cpwvcmplsttl" }
    if yname == "cpwVcMplsLocalLdpID" { return "Cpwvcmplslocalldpid" }
    if yname == "cpwVcMplsLocalLdpEntityID" { return "Cpwvcmplslocalldpentityid" }
    if yname == "cpwVcMplsPeerLdpID" { return "Cpwvcmplspeerldpid" }
    if yname == "cpwVcMplsStorageType" { return "Cpwvcmplsstoragetype" }
    return ""
}

func (cpwvcmplsentry *CISCOIETFPWMPLSMIB_Cpwvcmplstable_Cpwvcmplsentry) GetSegmentPath() string {
    return "cpwVcMplsEntry" + "[cpwVcIndex='" + fmt.Sprintf("%v", cpwvcmplsentry.Cpwvcindex) + "']"
}

func (cpwvcmplsentry *CISCOIETFPWMPLSMIB_Cpwvcmplstable_Cpwvcmplsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cpwvcmplsentry *CISCOIETFPWMPLSMIB_Cpwvcmplstable_Cpwvcmplsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cpwvcmplsentry *CISCOIETFPWMPLSMIB_Cpwvcmplstable_Cpwvcmplsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cpwVcIndex"] = cpwvcmplsentry.Cpwvcindex
    leafs["cpwVcMplsMplsType"] = cpwvcmplsentry.Cpwvcmplsmplstype
    leafs["cpwVcMplsExpBitsMode"] = cpwvcmplsentry.Cpwvcmplsexpbitsmode
    leafs["cpwVcMplsExpBits"] = cpwvcmplsentry.Cpwvcmplsexpbits
    leafs["cpwVcMplsTtl"] = cpwvcmplsentry.Cpwvcmplsttl
    leafs["cpwVcMplsLocalLdpID"] = cpwvcmplsentry.Cpwvcmplslocalldpid
    leafs["cpwVcMplsLocalLdpEntityID"] = cpwvcmplsentry.Cpwvcmplslocalldpentityid
    leafs["cpwVcMplsPeerLdpID"] = cpwvcmplsentry.Cpwvcmplspeerldpid
    leafs["cpwVcMplsStorageType"] = cpwvcmplsentry.Cpwvcmplsstoragetype
    return leafs
}

func (cpwvcmplsentry *CISCOIETFPWMPLSMIB_Cpwvcmplstable_Cpwvcmplsentry) GetBundleName() string { return "cisco_ios_xe" }

func (cpwvcmplsentry *CISCOIETFPWMPLSMIB_Cpwvcmplstable_Cpwvcmplsentry) GetYangName() string { return "cpwVcMplsEntry" }

func (cpwvcmplsentry *CISCOIETFPWMPLSMIB_Cpwvcmplstable_Cpwvcmplsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpwvcmplsentry *CISCOIETFPWMPLSMIB_Cpwvcmplstable_Cpwvcmplsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpwvcmplsentry *CISCOIETFPWMPLSMIB_Cpwvcmplstable_Cpwvcmplsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpwvcmplsentry *CISCOIETFPWMPLSMIB_Cpwvcmplstable_Cpwvcmplsentry) SetParent(parent types.Entity) { cpwvcmplsentry.parent = parent }

func (cpwvcmplsentry *CISCOIETFPWMPLSMIB_Cpwvcmplstable_Cpwvcmplsentry) GetParent() types.Entity { return cpwvcmplsentry.parent }

func (cpwvcmplsentry *CISCOIETFPWMPLSMIB_Cpwvcmplstable_Cpwvcmplsentry) GetParentYangName() string { return "cpwVcMplsTable" }

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
    parent types.Entity
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

func (cpwvcmplsoutboundtable *CISCOIETFPWMPLSMIB_Cpwvcmplsoutboundtable) GetFilter() yfilter.YFilter { return cpwvcmplsoutboundtable.YFilter }

func (cpwvcmplsoutboundtable *CISCOIETFPWMPLSMIB_Cpwvcmplsoutboundtable) SetFilter(yf yfilter.YFilter) { cpwvcmplsoutboundtable.YFilter = yf }

func (cpwvcmplsoutboundtable *CISCOIETFPWMPLSMIB_Cpwvcmplsoutboundtable) GetGoName(yname string) string {
    if yname == "cpwVcMplsOutboundEntry" { return "Cpwvcmplsoutboundentry" }
    return ""
}

func (cpwvcmplsoutboundtable *CISCOIETFPWMPLSMIB_Cpwvcmplsoutboundtable) GetSegmentPath() string {
    return "cpwVcMplsOutboundTable"
}

func (cpwvcmplsoutboundtable *CISCOIETFPWMPLSMIB_Cpwvcmplsoutboundtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cpwVcMplsOutboundEntry" {
        for _, c := range cpwvcmplsoutboundtable.Cpwvcmplsoutboundentry {
            if cpwvcmplsoutboundtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIETFPWMPLSMIB_Cpwvcmplsoutboundtable_Cpwvcmplsoutboundentry{}
        cpwvcmplsoutboundtable.Cpwvcmplsoutboundentry = append(cpwvcmplsoutboundtable.Cpwvcmplsoutboundentry, child)
        return &cpwvcmplsoutboundtable.Cpwvcmplsoutboundentry[len(cpwvcmplsoutboundtable.Cpwvcmplsoutboundentry)-1]
    }
    return nil
}

func (cpwvcmplsoutboundtable *CISCOIETFPWMPLSMIB_Cpwvcmplsoutboundtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cpwvcmplsoutboundtable.Cpwvcmplsoutboundentry {
        children[cpwvcmplsoutboundtable.Cpwvcmplsoutboundentry[i].GetSegmentPath()] = &cpwvcmplsoutboundtable.Cpwvcmplsoutboundentry[i]
    }
    return children
}

func (cpwvcmplsoutboundtable *CISCOIETFPWMPLSMIB_Cpwvcmplsoutboundtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cpwvcmplsoutboundtable *CISCOIETFPWMPLSMIB_Cpwvcmplsoutboundtable) GetBundleName() string { return "cisco_ios_xe" }

func (cpwvcmplsoutboundtable *CISCOIETFPWMPLSMIB_Cpwvcmplsoutboundtable) GetYangName() string { return "cpwVcMplsOutboundTable" }

func (cpwvcmplsoutboundtable *CISCOIETFPWMPLSMIB_Cpwvcmplsoutboundtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpwvcmplsoutboundtable *CISCOIETFPWMPLSMIB_Cpwvcmplsoutboundtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpwvcmplsoutboundtable *CISCOIETFPWMPLSMIB_Cpwvcmplsoutboundtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpwvcmplsoutboundtable *CISCOIETFPWMPLSMIB_Cpwvcmplsoutboundtable) SetParent(parent types.Entity) { cpwvcmplsoutboundtable.parent = parent }

func (cpwvcmplsoutboundtable *CISCOIETFPWMPLSMIB_Cpwvcmplsoutboundtable) GetParent() types.Entity { return cpwvcmplsoutboundtable.parent }

func (cpwvcmplsoutboundtable *CISCOIETFPWMPLSMIB_Cpwvcmplsoutboundtable) GetParentYangName() string { return "CISCO-IETF-PW-MPLS-MIB" }

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
    parent types.Entity
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

func (cpwvcmplsoutboundentry *CISCOIETFPWMPLSMIB_Cpwvcmplsoutboundtable_Cpwvcmplsoutboundentry) GetFilter() yfilter.YFilter { return cpwvcmplsoutboundentry.YFilter }

func (cpwvcmplsoutboundentry *CISCOIETFPWMPLSMIB_Cpwvcmplsoutboundtable_Cpwvcmplsoutboundentry) SetFilter(yf yfilter.YFilter) { cpwvcmplsoutboundentry.YFilter = yf }

func (cpwvcmplsoutboundentry *CISCOIETFPWMPLSMIB_Cpwvcmplsoutboundtable_Cpwvcmplsoutboundentry) GetGoName(yname string) string {
    if yname == "cpwVcIndex" { return "Cpwvcindex" }
    if yname == "cpwVcMplsOutboundIndex" { return "Cpwvcmplsoutboundindex" }
    if yname == "cpwVcMplsOutboundLsrXcIndex" { return "Cpwvcmplsoutboundlsrxcindex" }
    if yname == "cpwVcMplsOutboundTunnelIndex" { return "Cpwvcmplsoutboundtunnelindex" }
    if yname == "cpwVcMplsOutboundTunnelInstance" { return "Cpwvcmplsoutboundtunnelinstance" }
    if yname == "cpwVcMplsOutboundTunnelLclLSR" { return "Cpwvcmplsoutboundtunnellcllsr" }
    if yname == "cpwVcMplsOutboundTunnelPeerLSR" { return "Cpwvcmplsoutboundtunnelpeerlsr" }
    if yname == "cpwVcMplsOutboundIfIndex" { return "Cpwvcmplsoutboundifindex" }
    if yname == "cpwVcMplsOutboundRowStatus" { return "Cpwvcmplsoutboundrowstatus" }
    if yname == "cpwVcMplsOutboundStorageType" { return "Cpwvcmplsoutboundstoragetype" }
    return ""
}

func (cpwvcmplsoutboundentry *CISCOIETFPWMPLSMIB_Cpwvcmplsoutboundtable_Cpwvcmplsoutboundentry) GetSegmentPath() string {
    return "cpwVcMplsOutboundEntry" + "[cpwVcIndex='" + fmt.Sprintf("%v", cpwvcmplsoutboundentry.Cpwvcindex) + "']" + "[cpwVcMplsOutboundIndex='" + fmt.Sprintf("%v", cpwvcmplsoutboundentry.Cpwvcmplsoutboundindex) + "']"
}

func (cpwvcmplsoutboundentry *CISCOIETFPWMPLSMIB_Cpwvcmplsoutboundtable_Cpwvcmplsoutboundentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cpwvcmplsoutboundentry *CISCOIETFPWMPLSMIB_Cpwvcmplsoutboundtable_Cpwvcmplsoutboundentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cpwvcmplsoutboundentry *CISCOIETFPWMPLSMIB_Cpwvcmplsoutboundtable_Cpwvcmplsoutboundentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cpwVcIndex"] = cpwvcmplsoutboundentry.Cpwvcindex
    leafs["cpwVcMplsOutboundIndex"] = cpwvcmplsoutboundentry.Cpwvcmplsoutboundindex
    leafs["cpwVcMplsOutboundLsrXcIndex"] = cpwvcmplsoutboundentry.Cpwvcmplsoutboundlsrxcindex
    leafs["cpwVcMplsOutboundTunnelIndex"] = cpwvcmplsoutboundentry.Cpwvcmplsoutboundtunnelindex
    leafs["cpwVcMplsOutboundTunnelInstance"] = cpwvcmplsoutboundentry.Cpwvcmplsoutboundtunnelinstance
    leafs["cpwVcMplsOutboundTunnelLclLSR"] = cpwvcmplsoutboundentry.Cpwvcmplsoutboundtunnellcllsr
    leafs["cpwVcMplsOutboundTunnelPeerLSR"] = cpwvcmplsoutboundentry.Cpwvcmplsoutboundtunnelpeerlsr
    leafs["cpwVcMplsOutboundIfIndex"] = cpwvcmplsoutboundentry.Cpwvcmplsoutboundifindex
    leafs["cpwVcMplsOutboundRowStatus"] = cpwvcmplsoutboundentry.Cpwvcmplsoutboundrowstatus
    leafs["cpwVcMplsOutboundStorageType"] = cpwvcmplsoutboundentry.Cpwvcmplsoutboundstoragetype
    return leafs
}

func (cpwvcmplsoutboundentry *CISCOIETFPWMPLSMIB_Cpwvcmplsoutboundtable_Cpwvcmplsoutboundentry) GetBundleName() string { return "cisco_ios_xe" }

func (cpwvcmplsoutboundentry *CISCOIETFPWMPLSMIB_Cpwvcmplsoutboundtable_Cpwvcmplsoutboundentry) GetYangName() string { return "cpwVcMplsOutboundEntry" }

func (cpwvcmplsoutboundentry *CISCOIETFPWMPLSMIB_Cpwvcmplsoutboundtable_Cpwvcmplsoutboundentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpwvcmplsoutboundentry *CISCOIETFPWMPLSMIB_Cpwvcmplsoutboundtable_Cpwvcmplsoutboundentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpwvcmplsoutboundentry *CISCOIETFPWMPLSMIB_Cpwvcmplsoutboundtable_Cpwvcmplsoutboundentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpwvcmplsoutboundentry *CISCOIETFPWMPLSMIB_Cpwvcmplsoutboundtable_Cpwvcmplsoutboundentry) SetParent(parent types.Entity) { cpwvcmplsoutboundentry.parent = parent }

func (cpwvcmplsoutboundentry *CISCOIETFPWMPLSMIB_Cpwvcmplsoutboundtable_Cpwvcmplsoutboundentry) GetParent() types.Entity { return cpwvcmplsoutboundentry.parent }

func (cpwvcmplsoutboundentry *CISCOIETFPWMPLSMIB_Cpwvcmplsoutboundtable_Cpwvcmplsoutboundentry) GetParentYangName() string { return "cpwVcMplsOutboundTable" }

// CISCOIETFPWMPLSMIB_Cpwvcmplsinboundtable
// This table associates VCs using MPLS PSN with the inbound 
// MPLS tunnels (i.e. for packets coming from the PSN),  
// if such association is desired (mainly for security  
// reasons).
type CISCOIETFPWMPLSMIB_Cpwvcmplsinboundtable struct {
    parent types.Entity
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

func (cpwvcmplsinboundtable *CISCOIETFPWMPLSMIB_Cpwvcmplsinboundtable) GetFilter() yfilter.YFilter { return cpwvcmplsinboundtable.YFilter }

func (cpwvcmplsinboundtable *CISCOIETFPWMPLSMIB_Cpwvcmplsinboundtable) SetFilter(yf yfilter.YFilter) { cpwvcmplsinboundtable.YFilter = yf }

func (cpwvcmplsinboundtable *CISCOIETFPWMPLSMIB_Cpwvcmplsinboundtable) GetGoName(yname string) string {
    if yname == "cpwVcMplsInboundEntry" { return "Cpwvcmplsinboundentry" }
    return ""
}

func (cpwvcmplsinboundtable *CISCOIETFPWMPLSMIB_Cpwvcmplsinboundtable) GetSegmentPath() string {
    return "cpwVcMplsInboundTable"
}

func (cpwvcmplsinboundtable *CISCOIETFPWMPLSMIB_Cpwvcmplsinboundtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cpwVcMplsInboundEntry" {
        for _, c := range cpwvcmplsinboundtable.Cpwvcmplsinboundentry {
            if cpwvcmplsinboundtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIETFPWMPLSMIB_Cpwvcmplsinboundtable_Cpwvcmplsinboundentry{}
        cpwvcmplsinboundtable.Cpwvcmplsinboundentry = append(cpwvcmplsinboundtable.Cpwvcmplsinboundentry, child)
        return &cpwvcmplsinboundtable.Cpwvcmplsinboundentry[len(cpwvcmplsinboundtable.Cpwvcmplsinboundentry)-1]
    }
    return nil
}

func (cpwvcmplsinboundtable *CISCOIETFPWMPLSMIB_Cpwvcmplsinboundtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cpwvcmplsinboundtable.Cpwvcmplsinboundentry {
        children[cpwvcmplsinboundtable.Cpwvcmplsinboundentry[i].GetSegmentPath()] = &cpwvcmplsinboundtable.Cpwvcmplsinboundentry[i]
    }
    return children
}

func (cpwvcmplsinboundtable *CISCOIETFPWMPLSMIB_Cpwvcmplsinboundtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cpwvcmplsinboundtable *CISCOIETFPWMPLSMIB_Cpwvcmplsinboundtable) GetBundleName() string { return "cisco_ios_xe" }

func (cpwvcmplsinboundtable *CISCOIETFPWMPLSMIB_Cpwvcmplsinboundtable) GetYangName() string { return "cpwVcMplsInboundTable" }

func (cpwvcmplsinboundtable *CISCOIETFPWMPLSMIB_Cpwvcmplsinboundtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpwvcmplsinboundtable *CISCOIETFPWMPLSMIB_Cpwvcmplsinboundtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpwvcmplsinboundtable *CISCOIETFPWMPLSMIB_Cpwvcmplsinboundtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpwvcmplsinboundtable *CISCOIETFPWMPLSMIB_Cpwvcmplsinboundtable) SetParent(parent types.Entity) { cpwvcmplsinboundtable.parent = parent }

func (cpwvcmplsinboundtable *CISCOIETFPWMPLSMIB_Cpwvcmplsinboundtable) GetParent() types.Entity { return cpwvcmplsinboundtable.parent }

func (cpwvcmplsinboundtable *CISCOIETFPWMPLSMIB_Cpwvcmplsinboundtable) GetParentYangName() string { return "CISCO-IETF-PW-MPLS-MIB" }

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
    parent types.Entity
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

func (cpwvcmplsinboundentry *CISCOIETFPWMPLSMIB_Cpwvcmplsinboundtable_Cpwvcmplsinboundentry) GetFilter() yfilter.YFilter { return cpwvcmplsinboundentry.YFilter }

func (cpwvcmplsinboundentry *CISCOIETFPWMPLSMIB_Cpwvcmplsinboundtable_Cpwvcmplsinboundentry) SetFilter(yf yfilter.YFilter) { cpwvcmplsinboundentry.YFilter = yf }

func (cpwvcmplsinboundentry *CISCOIETFPWMPLSMIB_Cpwvcmplsinboundtable_Cpwvcmplsinboundentry) GetGoName(yname string) string {
    if yname == "cpwVcIndex" { return "Cpwvcindex" }
    if yname == "cpwVcMplsInboundIndex" { return "Cpwvcmplsinboundindex" }
    if yname == "cpwVcMplsInboundLsrXcIndex" { return "Cpwvcmplsinboundlsrxcindex" }
    if yname == "cpwVcMplsInboundTunnelIndex" { return "Cpwvcmplsinboundtunnelindex" }
    if yname == "cpwVcMplsInboundTunnelInstance" { return "Cpwvcmplsinboundtunnelinstance" }
    if yname == "cpwVcMplsInboundTunnelLclLSR" { return "Cpwvcmplsinboundtunnellcllsr" }
    if yname == "cpwVcMplsInboundTunnelPeerLSR" { return "Cpwvcmplsinboundtunnelpeerlsr" }
    if yname == "cpwVcMplsInboundIfIndex" { return "Cpwvcmplsinboundifindex" }
    if yname == "cpwVcMplsInboundRowStatus" { return "Cpwvcmplsinboundrowstatus" }
    if yname == "cpwVcMplsInboundStorageType" { return "Cpwvcmplsinboundstoragetype" }
    return ""
}

func (cpwvcmplsinboundentry *CISCOIETFPWMPLSMIB_Cpwvcmplsinboundtable_Cpwvcmplsinboundentry) GetSegmentPath() string {
    return "cpwVcMplsInboundEntry" + "[cpwVcIndex='" + fmt.Sprintf("%v", cpwvcmplsinboundentry.Cpwvcindex) + "']" + "[cpwVcMplsInboundIndex='" + fmt.Sprintf("%v", cpwvcmplsinboundentry.Cpwvcmplsinboundindex) + "']"
}

func (cpwvcmplsinboundentry *CISCOIETFPWMPLSMIB_Cpwvcmplsinboundtable_Cpwvcmplsinboundentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cpwvcmplsinboundentry *CISCOIETFPWMPLSMIB_Cpwvcmplsinboundtable_Cpwvcmplsinboundentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cpwvcmplsinboundentry *CISCOIETFPWMPLSMIB_Cpwvcmplsinboundtable_Cpwvcmplsinboundentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cpwVcIndex"] = cpwvcmplsinboundentry.Cpwvcindex
    leafs["cpwVcMplsInboundIndex"] = cpwvcmplsinboundentry.Cpwvcmplsinboundindex
    leafs["cpwVcMplsInboundLsrXcIndex"] = cpwvcmplsinboundentry.Cpwvcmplsinboundlsrxcindex
    leafs["cpwVcMplsInboundTunnelIndex"] = cpwvcmplsinboundentry.Cpwvcmplsinboundtunnelindex
    leafs["cpwVcMplsInboundTunnelInstance"] = cpwvcmplsinboundentry.Cpwvcmplsinboundtunnelinstance
    leafs["cpwVcMplsInboundTunnelLclLSR"] = cpwvcmplsinboundentry.Cpwvcmplsinboundtunnellcllsr
    leafs["cpwVcMplsInboundTunnelPeerLSR"] = cpwvcmplsinboundentry.Cpwvcmplsinboundtunnelpeerlsr
    leafs["cpwVcMplsInboundIfIndex"] = cpwvcmplsinboundentry.Cpwvcmplsinboundifindex
    leafs["cpwVcMplsInboundRowStatus"] = cpwvcmplsinboundentry.Cpwvcmplsinboundrowstatus
    leafs["cpwVcMplsInboundStorageType"] = cpwvcmplsinboundentry.Cpwvcmplsinboundstoragetype
    return leafs
}

func (cpwvcmplsinboundentry *CISCOIETFPWMPLSMIB_Cpwvcmplsinboundtable_Cpwvcmplsinboundentry) GetBundleName() string { return "cisco_ios_xe" }

func (cpwvcmplsinboundentry *CISCOIETFPWMPLSMIB_Cpwvcmplsinboundtable_Cpwvcmplsinboundentry) GetYangName() string { return "cpwVcMplsInboundEntry" }

func (cpwvcmplsinboundentry *CISCOIETFPWMPLSMIB_Cpwvcmplsinboundtable_Cpwvcmplsinboundentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpwvcmplsinboundentry *CISCOIETFPWMPLSMIB_Cpwvcmplsinboundtable_Cpwvcmplsinboundentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpwvcmplsinboundentry *CISCOIETFPWMPLSMIB_Cpwvcmplsinboundtable_Cpwvcmplsinboundentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpwvcmplsinboundentry *CISCOIETFPWMPLSMIB_Cpwvcmplsinboundtable_Cpwvcmplsinboundentry) SetParent(parent types.Entity) { cpwvcmplsinboundentry.parent = parent }

func (cpwvcmplsinboundentry *CISCOIETFPWMPLSMIB_Cpwvcmplsinboundtable_Cpwvcmplsinboundentry) GetParent() types.Entity { return cpwvcmplsinboundentry.parent }

func (cpwvcmplsinboundentry *CISCOIETFPWMPLSMIB_Cpwvcmplsinboundtable_Cpwvcmplsinboundentry) GetParentYangName() string { return "cpwVcMplsInboundTable" }

// CISCOIETFPWMPLSMIB_Cpwvcmplsnontemappingtable
// This table maps an inbound/outbound Tunnel to a VC in non- 
// TE applications.
type CISCOIETFPWMPLSMIB_Cpwvcmplsnontemappingtable struct {
    parent types.Entity
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

func (cpwvcmplsnontemappingtable *CISCOIETFPWMPLSMIB_Cpwvcmplsnontemappingtable) GetFilter() yfilter.YFilter { return cpwvcmplsnontemappingtable.YFilter }

func (cpwvcmplsnontemappingtable *CISCOIETFPWMPLSMIB_Cpwvcmplsnontemappingtable) SetFilter(yf yfilter.YFilter) { cpwvcmplsnontemappingtable.YFilter = yf }

func (cpwvcmplsnontemappingtable *CISCOIETFPWMPLSMIB_Cpwvcmplsnontemappingtable) GetGoName(yname string) string {
    if yname == "cpwVcMplsNonTeMappingEntry" { return "Cpwvcmplsnontemappingentry" }
    return ""
}

func (cpwvcmplsnontemappingtable *CISCOIETFPWMPLSMIB_Cpwvcmplsnontemappingtable) GetSegmentPath() string {
    return "cpwVcMplsNonTeMappingTable"
}

func (cpwvcmplsnontemappingtable *CISCOIETFPWMPLSMIB_Cpwvcmplsnontemappingtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cpwVcMplsNonTeMappingEntry" {
        for _, c := range cpwvcmplsnontemappingtable.Cpwvcmplsnontemappingentry {
            if cpwvcmplsnontemappingtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIETFPWMPLSMIB_Cpwvcmplsnontemappingtable_Cpwvcmplsnontemappingentry{}
        cpwvcmplsnontemappingtable.Cpwvcmplsnontemappingentry = append(cpwvcmplsnontemappingtable.Cpwvcmplsnontemappingentry, child)
        return &cpwvcmplsnontemappingtable.Cpwvcmplsnontemappingentry[len(cpwvcmplsnontemappingtable.Cpwvcmplsnontemappingentry)-1]
    }
    return nil
}

func (cpwvcmplsnontemappingtable *CISCOIETFPWMPLSMIB_Cpwvcmplsnontemappingtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cpwvcmplsnontemappingtable.Cpwvcmplsnontemappingentry {
        children[cpwvcmplsnontemappingtable.Cpwvcmplsnontemappingentry[i].GetSegmentPath()] = &cpwvcmplsnontemappingtable.Cpwvcmplsnontemappingentry[i]
    }
    return children
}

func (cpwvcmplsnontemappingtable *CISCOIETFPWMPLSMIB_Cpwvcmplsnontemappingtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cpwvcmplsnontemappingtable *CISCOIETFPWMPLSMIB_Cpwvcmplsnontemappingtable) GetBundleName() string { return "cisco_ios_xe" }

func (cpwvcmplsnontemappingtable *CISCOIETFPWMPLSMIB_Cpwvcmplsnontemappingtable) GetYangName() string { return "cpwVcMplsNonTeMappingTable" }

func (cpwvcmplsnontemappingtable *CISCOIETFPWMPLSMIB_Cpwvcmplsnontemappingtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpwvcmplsnontemappingtable *CISCOIETFPWMPLSMIB_Cpwvcmplsnontemappingtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpwvcmplsnontemappingtable *CISCOIETFPWMPLSMIB_Cpwvcmplsnontemappingtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpwvcmplsnontemappingtable *CISCOIETFPWMPLSMIB_Cpwvcmplsnontemappingtable) SetParent(parent types.Entity) { cpwvcmplsnontemappingtable.parent = parent }

func (cpwvcmplsnontemappingtable *CISCOIETFPWMPLSMIB_Cpwvcmplsnontemappingtable) GetParent() types.Entity { return cpwvcmplsnontemappingtable.parent }

func (cpwvcmplsnontemappingtable *CISCOIETFPWMPLSMIB_Cpwvcmplsnontemappingtable) GetParentYangName() string { return "CISCO-IETF-PW-MPLS-MIB" }

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
    parent types.Entity
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

func (cpwvcmplsnontemappingentry *CISCOIETFPWMPLSMIB_Cpwvcmplsnontemappingtable_Cpwvcmplsnontemappingentry) GetFilter() yfilter.YFilter { return cpwvcmplsnontemappingentry.YFilter }

func (cpwvcmplsnontemappingentry *CISCOIETFPWMPLSMIB_Cpwvcmplsnontemappingtable_Cpwvcmplsnontemappingentry) SetFilter(yf yfilter.YFilter) { cpwvcmplsnontemappingentry.YFilter = yf }

func (cpwvcmplsnontemappingentry *CISCOIETFPWMPLSMIB_Cpwvcmplsnontemappingtable_Cpwvcmplsnontemappingentry) GetGoName(yname string) string {
    if yname == "cpwVcMplsNonTeMappingTunnelDirection" { return "Cpwvcmplsnontemappingtunneldirection" }
    if yname == "cpwVcMplsNonTeMappingXcTunnelIndex" { return "Cpwvcmplsnontemappingxctunnelindex" }
    if yname == "cpwVcMplsNonTeMappingIfIndex" { return "Cpwvcmplsnontemappingifindex" }
    if yname == "cpwVcMplsNonTeMappingVcIndex" { return "Cpwvcmplsnontemappingvcindex" }
    return ""
}

func (cpwvcmplsnontemappingentry *CISCOIETFPWMPLSMIB_Cpwvcmplsnontemappingtable_Cpwvcmplsnontemappingentry) GetSegmentPath() string {
    return "cpwVcMplsNonTeMappingEntry" + "[cpwVcMplsNonTeMappingTunnelDirection='" + fmt.Sprintf("%v", cpwvcmplsnontemappingentry.Cpwvcmplsnontemappingtunneldirection) + "']" + "[cpwVcMplsNonTeMappingXcTunnelIndex='" + fmt.Sprintf("%v", cpwvcmplsnontemappingentry.Cpwvcmplsnontemappingxctunnelindex) + "']" + "[cpwVcMplsNonTeMappingIfIndex='" + fmt.Sprintf("%v", cpwvcmplsnontemappingentry.Cpwvcmplsnontemappingifindex) + "']" + "[cpwVcMplsNonTeMappingVcIndex='" + fmt.Sprintf("%v", cpwvcmplsnontemappingentry.Cpwvcmplsnontemappingvcindex) + "']"
}

func (cpwvcmplsnontemappingentry *CISCOIETFPWMPLSMIB_Cpwvcmplsnontemappingtable_Cpwvcmplsnontemappingentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cpwvcmplsnontemappingentry *CISCOIETFPWMPLSMIB_Cpwvcmplsnontemappingtable_Cpwvcmplsnontemappingentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cpwvcmplsnontemappingentry *CISCOIETFPWMPLSMIB_Cpwvcmplsnontemappingtable_Cpwvcmplsnontemappingentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cpwVcMplsNonTeMappingTunnelDirection"] = cpwvcmplsnontemappingentry.Cpwvcmplsnontemappingtunneldirection
    leafs["cpwVcMplsNonTeMappingXcTunnelIndex"] = cpwvcmplsnontemappingentry.Cpwvcmplsnontemappingxctunnelindex
    leafs["cpwVcMplsNonTeMappingIfIndex"] = cpwvcmplsnontemappingentry.Cpwvcmplsnontemappingifindex
    leafs["cpwVcMplsNonTeMappingVcIndex"] = cpwvcmplsnontemappingentry.Cpwvcmplsnontemappingvcindex
    return leafs
}

func (cpwvcmplsnontemappingentry *CISCOIETFPWMPLSMIB_Cpwvcmplsnontemappingtable_Cpwvcmplsnontemappingentry) GetBundleName() string { return "cisco_ios_xe" }

func (cpwvcmplsnontemappingentry *CISCOIETFPWMPLSMIB_Cpwvcmplsnontemappingtable_Cpwvcmplsnontemappingentry) GetYangName() string { return "cpwVcMplsNonTeMappingEntry" }

func (cpwvcmplsnontemappingentry *CISCOIETFPWMPLSMIB_Cpwvcmplsnontemappingtable_Cpwvcmplsnontemappingentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpwvcmplsnontemappingentry *CISCOIETFPWMPLSMIB_Cpwvcmplsnontemappingtable_Cpwvcmplsnontemappingentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpwvcmplsnontemappingentry *CISCOIETFPWMPLSMIB_Cpwvcmplsnontemappingtable_Cpwvcmplsnontemappingentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpwvcmplsnontemappingentry *CISCOIETFPWMPLSMIB_Cpwvcmplsnontemappingtable_Cpwvcmplsnontemappingentry) SetParent(parent types.Entity) { cpwvcmplsnontemappingentry.parent = parent }

func (cpwvcmplsnontemappingentry *CISCOIETFPWMPLSMIB_Cpwvcmplsnontemappingtable_Cpwvcmplsnontemappingentry) GetParent() types.Entity { return cpwvcmplsnontemappingentry.parent }

func (cpwvcmplsnontemappingentry *CISCOIETFPWMPLSMIB_Cpwvcmplsnontemappingtable_Cpwvcmplsnontemappingentry) GetParentYangName() string { return "cpwVcMplsNonTeMappingTable" }

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
    parent types.Entity
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

func (cpwvcmplstemappingtable *CISCOIETFPWMPLSMIB_Cpwvcmplstemappingtable) GetFilter() yfilter.YFilter { return cpwvcmplstemappingtable.YFilter }

func (cpwvcmplstemappingtable *CISCOIETFPWMPLSMIB_Cpwvcmplstemappingtable) SetFilter(yf yfilter.YFilter) { cpwvcmplstemappingtable.YFilter = yf }

func (cpwvcmplstemappingtable *CISCOIETFPWMPLSMIB_Cpwvcmplstemappingtable) GetGoName(yname string) string {
    if yname == "cpwVcMplsTeMappingEntry" { return "Cpwvcmplstemappingentry" }
    return ""
}

func (cpwvcmplstemappingtable *CISCOIETFPWMPLSMIB_Cpwvcmplstemappingtable) GetSegmentPath() string {
    return "cpwVcMplsTeMappingTable"
}

func (cpwvcmplstemappingtable *CISCOIETFPWMPLSMIB_Cpwvcmplstemappingtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cpwVcMplsTeMappingEntry" {
        for _, c := range cpwvcmplstemappingtable.Cpwvcmplstemappingentry {
            if cpwvcmplstemappingtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIETFPWMPLSMIB_Cpwvcmplstemappingtable_Cpwvcmplstemappingentry{}
        cpwvcmplstemappingtable.Cpwvcmplstemappingentry = append(cpwvcmplstemappingtable.Cpwvcmplstemappingentry, child)
        return &cpwvcmplstemappingtable.Cpwvcmplstemappingentry[len(cpwvcmplstemappingtable.Cpwvcmplstemappingentry)-1]
    }
    return nil
}

func (cpwvcmplstemappingtable *CISCOIETFPWMPLSMIB_Cpwvcmplstemappingtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cpwvcmplstemappingtable.Cpwvcmplstemappingentry {
        children[cpwvcmplstemappingtable.Cpwvcmplstemappingentry[i].GetSegmentPath()] = &cpwvcmplstemappingtable.Cpwvcmplstemappingentry[i]
    }
    return children
}

func (cpwvcmplstemappingtable *CISCOIETFPWMPLSMIB_Cpwvcmplstemappingtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cpwvcmplstemappingtable *CISCOIETFPWMPLSMIB_Cpwvcmplstemappingtable) GetBundleName() string { return "cisco_ios_xe" }

func (cpwvcmplstemappingtable *CISCOIETFPWMPLSMIB_Cpwvcmplstemappingtable) GetYangName() string { return "cpwVcMplsTeMappingTable" }

func (cpwvcmplstemappingtable *CISCOIETFPWMPLSMIB_Cpwvcmplstemappingtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpwvcmplstemappingtable *CISCOIETFPWMPLSMIB_Cpwvcmplstemappingtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpwvcmplstemappingtable *CISCOIETFPWMPLSMIB_Cpwvcmplstemappingtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpwvcmplstemappingtable *CISCOIETFPWMPLSMIB_Cpwvcmplstemappingtable) SetParent(parent types.Entity) { cpwvcmplstemappingtable.parent = parent }

func (cpwvcmplstemappingtable *CISCOIETFPWMPLSMIB_Cpwvcmplstemappingtable) GetParent() types.Entity { return cpwvcmplstemappingtable.parent }

func (cpwvcmplstemappingtable *CISCOIETFPWMPLSMIB_Cpwvcmplstemappingtable) GetParentYangName() string { return "CISCO-IETF-PW-MPLS-MIB" }

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
    parent types.Entity
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

func (cpwvcmplstemappingentry *CISCOIETFPWMPLSMIB_Cpwvcmplstemappingtable_Cpwvcmplstemappingentry) GetFilter() yfilter.YFilter { return cpwvcmplstemappingentry.YFilter }

func (cpwvcmplstemappingentry *CISCOIETFPWMPLSMIB_Cpwvcmplstemappingtable_Cpwvcmplstemappingentry) SetFilter(yf yfilter.YFilter) { cpwvcmplstemappingentry.YFilter = yf }

func (cpwvcmplstemappingentry *CISCOIETFPWMPLSMIB_Cpwvcmplstemappingtable_Cpwvcmplstemappingentry) GetGoName(yname string) string {
    if yname == "cpwVcMplsTeMappingTunnelDirection" { return "Cpwvcmplstemappingtunneldirection" }
    if yname == "cpwVcMplsTeMappingTunnelIndex" { return "Cpwvcmplstemappingtunnelindex" }
    if yname == "cpwVcMplsTeMappingTunnelInstance" { return "Cpwvcmplstemappingtunnelinstance" }
    if yname == "cpwVcMplsTeMappingTunnelPeerLsrID" { return "Cpwvcmplstemappingtunnelpeerlsrid" }
    if yname == "cpwVcMplsTeMappingTunnelLocalLsrID" { return "Cpwvcmplstemappingtunnellocallsrid" }
    if yname == "cpwVcMplsTeMappingVcIndex" { return "Cpwvcmplstemappingvcindex" }
    return ""
}

func (cpwvcmplstemappingentry *CISCOIETFPWMPLSMIB_Cpwvcmplstemappingtable_Cpwvcmplstemappingentry) GetSegmentPath() string {
    return "cpwVcMplsTeMappingEntry" + "[cpwVcMplsTeMappingTunnelDirection='" + fmt.Sprintf("%v", cpwvcmplstemappingentry.Cpwvcmplstemappingtunneldirection) + "']" + "[cpwVcMplsTeMappingTunnelIndex='" + fmt.Sprintf("%v", cpwvcmplstemappingentry.Cpwvcmplstemappingtunnelindex) + "']" + "[cpwVcMplsTeMappingTunnelInstance='" + fmt.Sprintf("%v", cpwvcmplstemappingentry.Cpwvcmplstemappingtunnelinstance) + "']" + "[cpwVcMplsTeMappingTunnelPeerLsrID='" + fmt.Sprintf("%v", cpwvcmplstemappingentry.Cpwvcmplstemappingtunnelpeerlsrid) + "']" + "[cpwVcMplsTeMappingTunnelLocalLsrID='" + fmt.Sprintf("%v", cpwvcmplstemappingentry.Cpwvcmplstemappingtunnellocallsrid) + "']" + "[cpwVcMplsTeMappingVcIndex='" + fmt.Sprintf("%v", cpwvcmplstemappingentry.Cpwvcmplstemappingvcindex) + "']"
}

func (cpwvcmplstemappingentry *CISCOIETFPWMPLSMIB_Cpwvcmplstemappingtable_Cpwvcmplstemappingentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cpwvcmplstemappingentry *CISCOIETFPWMPLSMIB_Cpwvcmplstemappingtable_Cpwvcmplstemappingentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cpwvcmplstemappingentry *CISCOIETFPWMPLSMIB_Cpwvcmplstemappingtable_Cpwvcmplstemappingentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cpwVcMplsTeMappingTunnelDirection"] = cpwvcmplstemappingentry.Cpwvcmplstemappingtunneldirection
    leafs["cpwVcMplsTeMappingTunnelIndex"] = cpwvcmplstemappingentry.Cpwvcmplstemappingtunnelindex
    leafs["cpwVcMplsTeMappingTunnelInstance"] = cpwvcmplstemappingentry.Cpwvcmplstemappingtunnelinstance
    leafs["cpwVcMplsTeMappingTunnelPeerLsrID"] = cpwvcmplstemappingentry.Cpwvcmplstemappingtunnelpeerlsrid
    leafs["cpwVcMplsTeMappingTunnelLocalLsrID"] = cpwvcmplstemappingentry.Cpwvcmplstemappingtunnellocallsrid
    leafs["cpwVcMplsTeMappingVcIndex"] = cpwvcmplstemappingentry.Cpwvcmplstemappingvcindex
    return leafs
}

func (cpwvcmplstemappingentry *CISCOIETFPWMPLSMIB_Cpwvcmplstemappingtable_Cpwvcmplstemappingentry) GetBundleName() string { return "cisco_ios_xe" }

func (cpwvcmplstemappingentry *CISCOIETFPWMPLSMIB_Cpwvcmplstemappingtable_Cpwvcmplstemappingentry) GetYangName() string { return "cpwVcMplsTeMappingEntry" }

func (cpwvcmplstemappingentry *CISCOIETFPWMPLSMIB_Cpwvcmplstemappingtable_Cpwvcmplstemappingentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpwvcmplstemappingentry *CISCOIETFPWMPLSMIB_Cpwvcmplstemappingtable_Cpwvcmplstemappingentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpwvcmplstemappingentry *CISCOIETFPWMPLSMIB_Cpwvcmplstemappingtable_Cpwvcmplstemappingentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpwvcmplstemappingentry *CISCOIETFPWMPLSMIB_Cpwvcmplstemappingtable_Cpwvcmplstemappingentry) SetParent(parent types.Entity) { cpwvcmplstemappingentry.parent = parent }

func (cpwvcmplstemappingentry *CISCOIETFPWMPLSMIB_Cpwvcmplstemappingtable_Cpwvcmplstemappingentry) GetParent() types.Entity { return cpwvcmplstemappingentry.parent }

func (cpwvcmplstemappingentry *CISCOIETFPWMPLSMIB_Cpwvcmplstemappingtable_Cpwvcmplstemappingentry) GetParentYangName() string { return "cpwVcMplsTeMappingTable" }

// CISCOIETFPWMPLSMIB_Cpwvcmplstemappingtable_Cpwvcmplstemappingentry_Cpwvcmplstemappingtunneldirection represents mapping.
type CISCOIETFPWMPLSMIB_Cpwvcmplstemappingtable_Cpwvcmplstemappingentry_Cpwvcmplstemappingtunneldirection string

const (
    CISCOIETFPWMPLSMIB_Cpwvcmplstemappingtable_Cpwvcmplstemappingentry_Cpwvcmplstemappingtunneldirection_outbound CISCOIETFPWMPLSMIB_Cpwvcmplstemappingtable_Cpwvcmplstemappingentry_Cpwvcmplstemappingtunneldirection = "outbound"

    CISCOIETFPWMPLSMIB_Cpwvcmplstemappingtable_Cpwvcmplstemappingentry_Cpwvcmplstemappingtunneldirection_inbound CISCOIETFPWMPLSMIB_Cpwvcmplstemappingtable_Cpwvcmplstemappingentry_Cpwvcmplstemappingtunneldirection = "inbound"
)

