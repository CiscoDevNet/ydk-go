// This MIB contains managed object definitions for Pseudo Wire
// emulation of ATM over Packet Switched Networks(PSN).
// 
// This MIB reports to the PW-MIB. The PW-MIB contains
// structures and MIB associations generic to Pseudo-Wire
// Virtual Circuit (VC) emulation. VC-specific MIBs (such as
// this) contain config and stats for specific VC types.
package cisco_ietf_pw_atm_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_ietf_pw_atm_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-IETF-PW-ATM-MIB CISCO-IETF-PW-ATM-MIB}", reflect.TypeOf(CISCOIETFPWATMMIB{}))
    ydk.RegisterEntity("CISCO-IETF-PW-ATM-MIB:CISCO-IETF-PW-ATM-MIB", reflect.TypeOf(CISCOIETFPWATMMIB{}))
}

// CISCOIETFPWATMMIB
type CISCOIETFPWATMMIB struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This table specifies the information for an ATM interface, VC, VP to be
    // carried over PSN.
    Cpwvcatmtable CISCOIETFPWATMMIB_Cpwvcatmtable
}

func (cISCOIETFPWATMMIB *CISCOIETFPWATMMIB) GetFilter() yfilter.YFilter { return cISCOIETFPWATMMIB.YFilter }

func (cISCOIETFPWATMMIB *CISCOIETFPWATMMIB) SetFilter(yf yfilter.YFilter) { cISCOIETFPWATMMIB.YFilter = yf }

func (cISCOIETFPWATMMIB *CISCOIETFPWATMMIB) GetGoName(yname string) string {
    if yname == "cpwVcAtmTable" { return "Cpwvcatmtable" }
    return ""
}

func (cISCOIETFPWATMMIB *CISCOIETFPWATMMIB) GetSegmentPath() string {
    return "CISCO-IETF-PW-ATM-MIB:CISCO-IETF-PW-ATM-MIB"
}

func (cISCOIETFPWATMMIB *CISCOIETFPWATMMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cpwVcAtmTable" {
        return &cISCOIETFPWATMMIB.Cpwvcatmtable
    }
    return nil
}

func (cISCOIETFPWATMMIB *CISCOIETFPWATMMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cpwVcAtmTable"] = &cISCOIETFPWATMMIB.Cpwvcatmtable
    return children
}

func (cISCOIETFPWATMMIB *CISCOIETFPWATMMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOIETFPWATMMIB *CISCOIETFPWATMMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOIETFPWATMMIB *CISCOIETFPWATMMIB) GetYangName() string { return "CISCO-IETF-PW-ATM-MIB" }

func (cISCOIETFPWATMMIB *CISCOIETFPWATMMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOIETFPWATMMIB *CISCOIETFPWATMMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOIETFPWATMMIB *CISCOIETFPWATMMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOIETFPWATMMIB *CISCOIETFPWATMMIB) SetParent(parent types.Entity) { cISCOIETFPWATMMIB.parent = parent }

func (cISCOIETFPWATMMIB *CISCOIETFPWATMMIB) GetParent() types.Entity { return cISCOIETFPWATMMIB.parent }

func (cISCOIETFPWATMMIB *CISCOIETFPWATMMIB) GetParentYangName() string { return "CISCO-IETF-PW-ATM-MIB" }

// CISCOIETFPWATMMIB_Cpwvcatmtable
// This table specifies the information for an ATM interface, VC,
// VP to be carried over PSN.
type CISCOIETFPWATMMIB_Cpwvcatmtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A row in this table represents an ATM interface, VC, VP that needs to be
    // adapted and carried over PSN. This table is indexed by CpwVcIndex in
    // CISCO-IETF-PW-MIB. The type is slice of
    // CISCOIETFPWATMMIB_Cpwvcatmtable_Cpwvcatmentry.
    Cpwvcatmentry []CISCOIETFPWATMMIB_Cpwvcatmtable_Cpwvcatmentry
}

func (cpwvcatmtable *CISCOIETFPWATMMIB_Cpwvcatmtable) GetFilter() yfilter.YFilter { return cpwvcatmtable.YFilter }

func (cpwvcatmtable *CISCOIETFPWATMMIB_Cpwvcatmtable) SetFilter(yf yfilter.YFilter) { cpwvcatmtable.YFilter = yf }

func (cpwvcatmtable *CISCOIETFPWATMMIB_Cpwvcatmtable) GetGoName(yname string) string {
    if yname == "cpwVcAtmEntry" { return "Cpwvcatmentry" }
    return ""
}

func (cpwvcatmtable *CISCOIETFPWATMMIB_Cpwvcatmtable) GetSegmentPath() string {
    return "cpwVcAtmTable"
}

func (cpwvcatmtable *CISCOIETFPWATMMIB_Cpwvcatmtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cpwVcAtmEntry" {
        for _, c := range cpwvcatmtable.Cpwvcatmentry {
            if cpwvcatmtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIETFPWATMMIB_Cpwvcatmtable_Cpwvcatmentry{}
        cpwvcatmtable.Cpwvcatmentry = append(cpwvcatmtable.Cpwvcatmentry, child)
        return &cpwvcatmtable.Cpwvcatmentry[len(cpwvcatmtable.Cpwvcatmentry)-1]
    }
    return nil
}

func (cpwvcatmtable *CISCOIETFPWATMMIB_Cpwvcatmtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cpwvcatmtable.Cpwvcatmentry {
        children[cpwvcatmtable.Cpwvcatmentry[i].GetSegmentPath()] = &cpwvcatmtable.Cpwvcatmentry[i]
    }
    return children
}

func (cpwvcatmtable *CISCOIETFPWATMMIB_Cpwvcatmtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cpwvcatmtable *CISCOIETFPWATMMIB_Cpwvcatmtable) GetBundleName() string { return "cisco_ios_xe" }

func (cpwvcatmtable *CISCOIETFPWATMMIB_Cpwvcatmtable) GetYangName() string { return "cpwVcAtmTable" }

func (cpwvcatmtable *CISCOIETFPWATMMIB_Cpwvcatmtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpwvcatmtable *CISCOIETFPWATMMIB_Cpwvcatmtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpwvcatmtable *CISCOIETFPWATMMIB_Cpwvcatmtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpwvcatmtable *CISCOIETFPWATMMIB_Cpwvcatmtable) SetParent(parent types.Entity) { cpwvcatmtable.parent = parent }

func (cpwvcatmtable *CISCOIETFPWATMMIB_Cpwvcatmtable) GetParent() types.Entity { return cpwvcatmtable.parent }

func (cpwvcatmtable *CISCOIETFPWATMMIB_Cpwvcatmtable) GetParentYangName() string { return "CISCO-IETF-PW-ATM-MIB" }

// CISCOIETFPWATMMIB_Cpwvcatmtable_Cpwvcatmentry
// A row in this table represents an ATM interface, VC, VP
// that needs to be adapted and carried over PSN. This table
// is indexed by CpwVcIndex in CISCO-IETF-PW-MIB.
type CISCOIETFPWATMMIB_Cpwvcatmtable_Cpwvcatmentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to cisco_ietf_pw_mib.CISCOIETFPWMIB_Cpwvctable_Cpwvcentry_Cpwvcindex
    Cpwvcindex interface{}

    // The ATM Interface that receives cells from the ATM network. The type is
    // interface{} with range: 1..2147483647.
    Cpwatmif interface{}

    // VPI value of this ATM VC. The type is interface{} with range: 0..4095.
    Cpwatmvpi interface{}

    // VCI value of this ATM VC. The type is interface{} with range: 0..65535.
    Cpwatmvci interface{}

    // This Object indicates whether the CLP bits are considered when determining
    // the value placed in the Quality of Service fields (e.g. EXP fields of the
    // MPLS Label Stack) of the encapsulating protocol. The type is bool.
    Cpwatmclpqosmapping interface{}

    // This Object is used to create, modify or delete a row in this table. The
    // type is RowStatus.
    Cpwatmrowstatus interface{}

    // This Object indicates whether OAM Cells are transported on this VC. The
    // type is bool.
    Cpwatmoamcellsupported interface{}

    // This Object represents the scaling factor (% value) to be applied to ATM
    // QoS rates when calculating QoS rates for the PSN domain . For example, in
    // the cell transport mode the bandwidth needed in the PSN domain will be
    // higher (since PSN Transport header, PW header, and optional control word
    // have to transmitted with every cell), whereas in the AAL5 mode the
    // bandwidth needed in PSN domain will be less since cell headers will be
    // removed after reassembly. The type is interface{} with range:
    // -2147483648..2147483647.
    Cpwatmqosscalingfactor interface{}

    // This object is used to identify if the VC is configured to do Cell Packing.
    // The type is bool.
    Cpwatmcellpacking interface{}

    // This object indicates the maximum number of cells that get packed in one
    // packet. The type is interface{} with range: -2147483648..2147483647.
    Cpwatmmncp interface{}

    // This Object represents the maximum number of cell that can be packed in one
    // packet for peer interface. The type is interface{} with range:
    // -2147483648..2147483647.
    Cpwatmpeermncp interface{}

    // This object indicates if the packet going on the pseudowire is mpls or
    // l2tpv3 encapsulated. The type is Cpwatmencap.
    Cpwatmencap interface{}

    // This Object represents which MCPT timeout value. The type is interface{}
    // with range: -2147483648..2147483647.
    Cpwatmmcpttimeout interface{}

    // This object can be used to obtain the information on the number of cells
    // that were received and sent to the PSN. The type is interface{} with range:
    // 0..4294967295.
    Cpwatmcellsreceived interface{}

    // This object can be used to obtain the information on the number of cells
    // that were received from the PSN and sent over the ATM network. The type is
    // interface{} with range: 0..4294967295.
    Cpwatmcellssent interface{}

    // This Object indicates the number of cells that were rejected by this VC
    // because of policing. The type is interface{} with range: 0..4294967295.
    Cpwatmcellsrejected interface{}

    // This Object indicates the number of cells that were Tagged. The type is
    // interface{} with range: 0..4294967295.
    Cpwatmcellstagged interface{}

    // High Capacity counter for the number of cells that were received by this
    // VC. The type is interface{} with range: 0..18446744073709551615.
    Cpwatmhccellsreceived interface{}

    // High Capacity counter for the number of cells that were rejected by this VC
    // because of policing. The type is interface{} with range:
    // 0..18446744073709551615.
    Cpwatmhccellsrejected interface{}

    // High Capacity counter for the number of cells that were tagged. The type is
    // interface{} with range: 0..18446744073709551615.
    Cpwatmhccellstagged interface{}

    // It indicates the Average number of cells that were received in one packet.
    // The type is interface{} with range: 0..4294967295.
    Cpwatmavgcellspacked interface{}

    // This object can be used to obtain the information on the number of packets
    // that were received and sent to the PSN. The type is interface{} with range:
    // 0..4294967295.
    Cpwatmpktsreceived interface{}

    // This object indicates the number of packets that were sent to the atm
    // network. The type is interface{} with range: 0..4294967295.
    Cpwatmpktssent interface{}

    // This object indicates the number of packets that were rejected because of
    // Policing. The type is interface{} with range: 0..4294967295.
    Cpwatmpktsrejected interface{}
}

func (cpwvcatmentry *CISCOIETFPWATMMIB_Cpwvcatmtable_Cpwvcatmentry) GetFilter() yfilter.YFilter { return cpwvcatmentry.YFilter }

func (cpwvcatmentry *CISCOIETFPWATMMIB_Cpwvcatmtable_Cpwvcatmentry) SetFilter(yf yfilter.YFilter) { cpwvcatmentry.YFilter = yf }

func (cpwvcatmentry *CISCOIETFPWATMMIB_Cpwvcatmtable_Cpwvcatmentry) GetGoName(yname string) string {
    if yname == "cpwVcIndex" { return "Cpwvcindex" }
    if yname == "cpwAtmIf" { return "Cpwatmif" }
    if yname == "cpwAtmVpi" { return "Cpwatmvpi" }
    if yname == "cpwAtmVci" { return "Cpwatmvci" }
    if yname == "cpwAtmClpQosMapping" { return "Cpwatmclpqosmapping" }
    if yname == "cpwAtmRowStatus" { return "Cpwatmrowstatus" }
    if yname == "cpwAtmOamCellSupported" { return "Cpwatmoamcellsupported" }
    if yname == "cpwAtmQosScalingFactor" { return "Cpwatmqosscalingfactor" }
    if yname == "cpwAtmCellPacking" { return "Cpwatmcellpacking" }
    if yname == "cpwAtmMncp" { return "Cpwatmmncp" }
    if yname == "cpwAtmPeerMncp" { return "Cpwatmpeermncp" }
    if yname == "cpwAtmEncap" { return "Cpwatmencap" }
    if yname == "cpwAtmMcptTimeout" { return "Cpwatmmcpttimeout" }
    if yname == "cpwAtmCellsReceived" { return "Cpwatmcellsreceived" }
    if yname == "cpwAtmCellsSent" { return "Cpwatmcellssent" }
    if yname == "cpwAtmCellsRejected" { return "Cpwatmcellsrejected" }
    if yname == "cpwAtmCellsTagged" { return "Cpwatmcellstagged" }
    if yname == "cpwAtmHCCellsReceived" { return "Cpwatmhccellsreceived" }
    if yname == "cpwAtmHCCellsRejected" { return "Cpwatmhccellsrejected" }
    if yname == "cpwAtmHCCellsTagged" { return "Cpwatmhccellstagged" }
    if yname == "cpwAtmAvgCellsPacked" { return "Cpwatmavgcellspacked" }
    if yname == "cpwAtmPktsReceived" { return "Cpwatmpktsreceived" }
    if yname == "cpwAtmPktsSent" { return "Cpwatmpktssent" }
    if yname == "cpwAtmPktsRejected" { return "Cpwatmpktsrejected" }
    return ""
}

func (cpwvcatmentry *CISCOIETFPWATMMIB_Cpwvcatmtable_Cpwvcatmentry) GetSegmentPath() string {
    return "cpwVcAtmEntry" + "[cpwVcIndex='" + fmt.Sprintf("%v", cpwvcatmentry.Cpwvcindex) + "']"
}

func (cpwvcatmentry *CISCOIETFPWATMMIB_Cpwvcatmtable_Cpwvcatmentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cpwvcatmentry *CISCOIETFPWATMMIB_Cpwvcatmtable_Cpwvcatmentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cpwvcatmentry *CISCOIETFPWATMMIB_Cpwvcatmtable_Cpwvcatmentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cpwVcIndex"] = cpwvcatmentry.Cpwvcindex
    leafs["cpwAtmIf"] = cpwvcatmentry.Cpwatmif
    leafs["cpwAtmVpi"] = cpwvcatmentry.Cpwatmvpi
    leafs["cpwAtmVci"] = cpwvcatmentry.Cpwatmvci
    leafs["cpwAtmClpQosMapping"] = cpwvcatmentry.Cpwatmclpqosmapping
    leafs["cpwAtmRowStatus"] = cpwvcatmentry.Cpwatmrowstatus
    leafs["cpwAtmOamCellSupported"] = cpwvcatmentry.Cpwatmoamcellsupported
    leafs["cpwAtmQosScalingFactor"] = cpwvcatmentry.Cpwatmqosscalingfactor
    leafs["cpwAtmCellPacking"] = cpwvcatmentry.Cpwatmcellpacking
    leafs["cpwAtmMncp"] = cpwvcatmentry.Cpwatmmncp
    leafs["cpwAtmPeerMncp"] = cpwvcatmentry.Cpwatmpeermncp
    leafs["cpwAtmEncap"] = cpwvcatmentry.Cpwatmencap
    leafs["cpwAtmMcptTimeout"] = cpwvcatmentry.Cpwatmmcpttimeout
    leafs["cpwAtmCellsReceived"] = cpwvcatmentry.Cpwatmcellsreceived
    leafs["cpwAtmCellsSent"] = cpwvcatmentry.Cpwatmcellssent
    leafs["cpwAtmCellsRejected"] = cpwvcatmentry.Cpwatmcellsrejected
    leafs["cpwAtmCellsTagged"] = cpwvcatmentry.Cpwatmcellstagged
    leafs["cpwAtmHCCellsReceived"] = cpwvcatmentry.Cpwatmhccellsreceived
    leafs["cpwAtmHCCellsRejected"] = cpwvcatmentry.Cpwatmhccellsrejected
    leafs["cpwAtmHCCellsTagged"] = cpwvcatmentry.Cpwatmhccellstagged
    leafs["cpwAtmAvgCellsPacked"] = cpwvcatmentry.Cpwatmavgcellspacked
    leafs["cpwAtmPktsReceived"] = cpwvcatmentry.Cpwatmpktsreceived
    leafs["cpwAtmPktsSent"] = cpwvcatmentry.Cpwatmpktssent
    leafs["cpwAtmPktsRejected"] = cpwvcatmentry.Cpwatmpktsrejected
    return leafs
}

func (cpwvcatmentry *CISCOIETFPWATMMIB_Cpwvcatmtable_Cpwvcatmentry) GetBundleName() string { return "cisco_ios_xe" }

func (cpwvcatmentry *CISCOIETFPWATMMIB_Cpwvcatmtable_Cpwvcatmentry) GetYangName() string { return "cpwVcAtmEntry" }

func (cpwvcatmentry *CISCOIETFPWATMMIB_Cpwvcatmtable_Cpwvcatmentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpwvcatmentry *CISCOIETFPWATMMIB_Cpwvcatmtable_Cpwvcatmentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpwvcatmentry *CISCOIETFPWATMMIB_Cpwvcatmtable_Cpwvcatmentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpwvcatmentry *CISCOIETFPWATMMIB_Cpwvcatmtable_Cpwvcatmentry) SetParent(parent types.Entity) { cpwvcatmentry.parent = parent }

func (cpwvcatmentry *CISCOIETFPWATMMIB_Cpwvcatmtable_Cpwvcatmentry) GetParent() types.Entity { return cpwvcatmentry.parent }

func (cpwvcatmentry *CISCOIETFPWATMMIB_Cpwvcatmtable_Cpwvcatmentry) GetParentYangName() string { return "cpwVcAtmTable" }

// CISCOIETFPWATMMIB_Cpwvcatmtable_Cpwvcatmentry_Cpwatmencap represents is mpls or l2tpv3 encapsulated.
type CISCOIETFPWATMMIB_Cpwvcatmtable_Cpwvcatmentry_Cpwatmencap string

const (
    CISCOIETFPWATMMIB_Cpwvcatmtable_Cpwvcatmentry_Cpwatmencap_mpls CISCOIETFPWATMMIB_Cpwvcatmtable_Cpwvcatmentry_Cpwatmencap = "mpls"

    CISCOIETFPWATMMIB_Cpwvcatmtable_Cpwvcatmentry_Cpwatmencap_l2tpv3 CISCOIETFPWATMMIB_Cpwvcatmtable_Cpwvcatmentry_Cpwatmencap = "l2tpv3"

    CISCOIETFPWATMMIB_Cpwvcatmtable_Cpwvcatmentry_Cpwatmencap_unknown CISCOIETFPWATMMIB_Cpwvcatmtable_Cpwvcatmentry_Cpwatmencap = "unknown"
)

