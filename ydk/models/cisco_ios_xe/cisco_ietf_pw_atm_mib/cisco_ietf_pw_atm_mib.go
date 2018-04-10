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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This table specifies the information for an ATM interface, VC, VP to be
    // carried over PSN.
    Cpwvcatmtable CISCOIETFPWATMMIB_Cpwvcatmtable
}

func (cISCOIETFPWATMMIB *CISCOIETFPWATMMIB) GetEntityData() *types.CommonEntityData {
    cISCOIETFPWATMMIB.EntityData.YFilter = cISCOIETFPWATMMIB.YFilter
    cISCOIETFPWATMMIB.EntityData.YangName = "CISCO-IETF-PW-ATM-MIB"
    cISCOIETFPWATMMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOIETFPWATMMIB.EntityData.ParentYangName = "CISCO-IETF-PW-ATM-MIB"
    cISCOIETFPWATMMIB.EntityData.SegmentPath = "CISCO-IETF-PW-ATM-MIB:CISCO-IETF-PW-ATM-MIB"
    cISCOIETFPWATMMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOIETFPWATMMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOIETFPWATMMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOIETFPWATMMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOIETFPWATMMIB.EntityData.Children["cpwVcAtmTable"] = types.YChild{"Cpwvcatmtable", &cISCOIETFPWATMMIB.Cpwvcatmtable}
    cISCOIETFPWATMMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOIETFPWATMMIB.EntityData)
}

// CISCOIETFPWATMMIB_Cpwvcatmtable
// This table specifies the information for an ATM interface, VC,
// VP to be carried over PSN.
type CISCOIETFPWATMMIB_Cpwvcatmtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A row in this table represents an ATM interface, VC, VP that needs to be
    // adapted and carried over PSN. This table is indexed by CpwVcIndex in
    // CISCO-IETF-PW-MIB. The type is slice of
    // CISCOIETFPWATMMIB_Cpwvcatmtable_Cpwvcatmentry.
    Cpwvcatmentry []CISCOIETFPWATMMIB_Cpwvcatmtable_Cpwvcatmentry
}

func (cpwvcatmtable *CISCOIETFPWATMMIB_Cpwvcatmtable) GetEntityData() *types.CommonEntityData {
    cpwvcatmtable.EntityData.YFilter = cpwvcatmtable.YFilter
    cpwvcatmtable.EntityData.YangName = "cpwVcAtmTable"
    cpwvcatmtable.EntityData.BundleName = "cisco_ios_xe"
    cpwvcatmtable.EntityData.ParentYangName = "CISCO-IETF-PW-ATM-MIB"
    cpwvcatmtable.EntityData.SegmentPath = "cpwVcAtmTable"
    cpwvcatmtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwvcatmtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwvcatmtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwvcatmtable.EntityData.Children = make(map[string]types.YChild)
    cpwvcatmtable.EntityData.Children["cpwVcAtmEntry"] = types.YChild{"Cpwvcatmentry", nil}
    for i := range cpwvcatmtable.Cpwvcatmentry {
        cpwvcatmtable.EntityData.Children[types.GetSegmentPath(&cpwvcatmtable.Cpwvcatmentry[i])] = types.YChild{"Cpwvcatmentry", &cpwvcatmtable.Cpwvcatmentry[i]}
    }
    cpwvcatmtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cpwvcatmtable.EntityData)
}

// CISCOIETFPWATMMIB_Cpwvcatmtable_Cpwvcatmentry
// A row in this table represents an ATM interface, VC, VP
// that needs to be adapted and carried over PSN. This table
// is indexed by CpwVcIndex in CISCO-IETF-PW-MIB.
type CISCOIETFPWATMMIB_Cpwvcatmtable_Cpwvcatmentry struct {
    EntityData types.CommonEntityData
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

func (cpwvcatmentry *CISCOIETFPWATMMIB_Cpwvcatmtable_Cpwvcatmentry) GetEntityData() *types.CommonEntityData {
    cpwvcatmentry.EntityData.YFilter = cpwvcatmentry.YFilter
    cpwvcatmentry.EntityData.YangName = "cpwVcAtmEntry"
    cpwvcatmentry.EntityData.BundleName = "cisco_ios_xe"
    cpwvcatmentry.EntityData.ParentYangName = "cpwVcAtmTable"
    cpwvcatmentry.EntityData.SegmentPath = "cpwVcAtmEntry" + "[cpwVcIndex='" + fmt.Sprintf("%v", cpwvcatmentry.Cpwvcindex) + "']"
    cpwvcatmentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwvcatmentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwvcatmentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwvcatmentry.EntityData.Children = make(map[string]types.YChild)
    cpwvcatmentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cpwvcatmentry.EntityData.Leafs["cpwVcIndex"] = types.YLeaf{"Cpwvcindex", cpwvcatmentry.Cpwvcindex}
    cpwvcatmentry.EntityData.Leafs["cpwAtmIf"] = types.YLeaf{"Cpwatmif", cpwvcatmentry.Cpwatmif}
    cpwvcatmentry.EntityData.Leafs["cpwAtmVpi"] = types.YLeaf{"Cpwatmvpi", cpwvcatmentry.Cpwatmvpi}
    cpwvcatmentry.EntityData.Leafs["cpwAtmVci"] = types.YLeaf{"Cpwatmvci", cpwvcatmentry.Cpwatmvci}
    cpwvcatmentry.EntityData.Leafs["cpwAtmClpQosMapping"] = types.YLeaf{"Cpwatmclpqosmapping", cpwvcatmentry.Cpwatmclpqosmapping}
    cpwvcatmentry.EntityData.Leafs["cpwAtmRowStatus"] = types.YLeaf{"Cpwatmrowstatus", cpwvcatmentry.Cpwatmrowstatus}
    cpwvcatmentry.EntityData.Leafs["cpwAtmOamCellSupported"] = types.YLeaf{"Cpwatmoamcellsupported", cpwvcatmentry.Cpwatmoamcellsupported}
    cpwvcatmentry.EntityData.Leafs["cpwAtmQosScalingFactor"] = types.YLeaf{"Cpwatmqosscalingfactor", cpwvcatmentry.Cpwatmqosscalingfactor}
    cpwvcatmentry.EntityData.Leafs["cpwAtmCellPacking"] = types.YLeaf{"Cpwatmcellpacking", cpwvcatmentry.Cpwatmcellpacking}
    cpwvcatmentry.EntityData.Leafs["cpwAtmMncp"] = types.YLeaf{"Cpwatmmncp", cpwvcatmentry.Cpwatmmncp}
    cpwvcatmentry.EntityData.Leafs["cpwAtmPeerMncp"] = types.YLeaf{"Cpwatmpeermncp", cpwvcatmentry.Cpwatmpeermncp}
    cpwvcatmentry.EntityData.Leafs["cpwAtmEncap"] = types.YLeaf{"Cpwatmencap", cpwvcatmentry.Cpwatmencap}
    cpwvcatmentry.EntityData.Leafs["cpwAtmMcptTimeout"] = types.YLeaf{"Cpwatmmcpttimeout", cpwvcatmentry.Cpwatmmcpttimeout}
    cpwvcatmentry.EntityData.Leafs["cpwAtmCellsReceived"] = types.YLeaf{"Cpwatmcellsreceived", cpwvcatmentry.Cpwatmcellsreceived}
    cpwvcatmentry.EntityData.Leafs["cpwAtmCellsSent"] = types.YLeaf{"Cpwatmcellssent", cpwvcatmentry.Cpwatmcellssent}
    cpwvcatmentry.EntityData.Leafs["cpwAtmCellsRejected"] = types.YLeaf{"Cpwatmcellsrejected", cpwvcatmentry.Cpwatmcellsrejected}
    cpwvcatmentry.EntityData.Leafs["cpwAtmCellsTagged"] = types.YLeaf{"Cpwatmcellstagged", cpwvcatmentry.Cpwatmcellstagged}
    cpwvcatmentry.EntityData.Leafs["cpwAtmHCCellsReceived"] = types.YLeaf{"Cpwatmhccellsreceived", cpwvcatmentry.Cpwatmhccellsreceived}
    cpwvcatmentry.EntityData.Leafs["cpwAtmHCCellsRejected"] = types.YLeaf{"Cpwatmhccellsrejected", cpwvcatmentry.Cpwatmhccellsrejected}
    cpwvcatmentry.EntityData.Leafs["cpwAtmHCCellsTagged"] = types.YLeaf{"Cpwatmhccellstagged", cpwvcatmentry.Cpwatmhccellstagged}
    cpwvcatmentry.EntityData.Leafs["cpwAtmAvgCellsPacked"] = types.YLeaf{"Cpwatmavgcellspacked", cpwvcatmentry.Cpwatmavgcellspacked}
    cpwvcatmentry.EntityData.Leafs["cpwAtmPktsReceived"] = types.YLeaf{"Cpwatmpktsreceived", cpwvcatmentry.Cpwatmpktsreceived}
    cpwvcatmentry.EntityData.Leafs["cpwAtmPktsSent"] = types.YLeaf{"Cpwatmpktssent", cpwvcatmentry.Cpwatmpktssent}
    cpwvcatmentry.EntityData.Leafs["cpwAtmPktsRejected"] = types.YLeaf{"Cpwatmpktsrejected", cpwvcatmentry.Cpwatmpktsrejected}
    return &(cpwvcatmentry.EntityData)
}

// CISCOIETFPWATMMIB_Cpwvcatmtable_Cpwvcatmentry_Cpwatmencap represents is mpls or l2tpv3 encapsulated.
type CISCOIETFPWATMMIB_Cpwvcatmtable_Cpwvcatmentry_Cpwatmencap string

const (
    CISCOIETFPWATMMIB_Cpwvcatmtable_Cpwvcatmentry_Cpwatmencap_mpls CISCOIETFPWATMMIB_Cpwvcatmtable_Cpwvcatmentry_Cpwatmencap = "mpls"

    CISCOIETFPWATMMIB_Cpwvcatmtable_Cpwvcatmentry_Cpwatmencap_l2tpv3 CISCOIETFPWATMMIB_Cpwvcatmtable_Cpwvcatmentry_Cpwatmencap = "l2tpv3"

    CISCOIETFPWATMMIB_Cpwvcatmtable_Cpwvcatmentry_Cpwatmencap_unknown CISCOIETFPWATMMIB_Cpwvcatmtable_Cpwvcatmentry_Cpwatmencap = "unknown"
)

