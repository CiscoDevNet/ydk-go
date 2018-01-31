// The MIB is created to provide ATM QoS information in
// these areas:
// 
// 1. Traffic shaping on a per-VC basis
// 2. Traffic shaping on a per-VP basis 
// 3. Per-VC queuing/buffering
// 
// Although the initial requirements of the MIB are
// driven to support the GSR TAZ line card,
// CISCO-ATM-QOS-MIB is designed as a generic MIB to
// support ATM interfaces across all platforms.
// 
// Here are the tables defined in this MIB:
// ciscoAtmQosVccTable 
//      - to provide information on traffic shaping on 
//        a per-VC basis.
//                       
// ciscoAtmQosVpcTable 
//      - to provide information on traffic shaping on
//        a per-VP basis.
//     
// ciscoAtmQosVcQueuingTable
// ciscoAtmQosVcQueuingClassTable
//      - to provide information on per-VC
//        queuing/buffering.
package cisco_atm_qos_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_atm_qos_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-ATM-QOS-MIB CISCO-ATM-QOS-MIB}", reflect.TypeOf(CISCOATMQOSMIB{}))
    ydk.RegisterEntity("CISCO-ATM-QOS-MIB:CISCO-ATM-QOS-MIB", reflect.TypeOf(CISCOATMQOSMIB{}))
}

// VcParamConfigLocation represents 5 - VC-class configured on main-interface.
type VcParamConfigLocation string

const (
    VcParamConfigLocation_configDefault VcParamConfigLocation = "configDefault"

    VcParamConfigLocation_configVcDirect VcParamConfigLocation = "configVcDirect"

    VcParamConfigLocation_configVcClass VcParamConfigLocation = "configVcClass"

    VcParamConfigLocation_configVcClassSubInterface VcParamConfigLocation = "configVcClassSubInterface"

    VcParamConfigLocation_configVcClassInterface VcParamConfigLocation = "configVcClassInterface"
)

// VpState represents 2 - Active
type VpState string

const (
    VpState_vpStateInactive VpState = "vpStateInactive"

    VpState_vpStateActive VpState = "vpStateActive"
)

// CISCOATMQOSMIB
type CISCOATMQOSMIB struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This table is defined to provide QoS information for each active ATM VC
    // existing on the interface.
    Caqvccparamstable CISCOATMQOSMIB_Caqvccparamstable

    // This table is defined to provide QoS information for each active ATM VP
    // existing on the interface.
    Caqvpcparamstable CISCOATMQOSMIB_Caqvpcparamstable

    // This table provides queuing related information for a VC existing on an ATM
    // interface.
    Caqqueuingparamstable CISCOATMQOSMIB_Caqqueuingparamstable

    // This table provides queuing information for all  queuing classes
    // associating with a VC.
    Caqqueuingparamsclasstable CISCOATMQOSMIB_Caqqueuingparamsclasstable
}

func (cISCOATMQOSMIB *CISCOATMQOSMIB) GetFilter() yfilter.YFilter { return cISCOATMQOSMIB.YFilter }

func (cISCOATMQOSMIB *CISCOATMQOSMIB) SetFilter(yf yfilter.YFilter) { cISCOATMQOSMIB.YFilter = yf }

func (cISCOATMQOSMIB *CISCOATMQOSMIB) GetGoName(yname string) string {
    if yname == "caqVccParamsTable" { return "Caqvccparamstable" }
    if yname == "caqVpcParamsTable" { return "Caqvpcparamstable" }
    if yname == "caqQueuingParamsTable" { return "Caqqueuingparamstable" }
    if yname == "caqQueuingParamsClassTable" { return "Caqqueuingparamsclasstable" }
    return ""
}

func (cISCOATMQOSMIB *CISCOATMQOSMIB) GetSegmentPath() string {
    return "CISCO-ATM-QOS-MIB:CISCO-ATM-QOS-MIB"
}

func (cISCOATMQOSMIB *CISCOATMQOSMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "caqVccParamsTable" {
        return &cISCOATMQOSMIB.Caqvccparamstable
    }
    if childYangName == "caqVpcParamsTable" {
        return &cISCOATMQOSMIB.Caqvpcparamstable
    }
    if childYangName == "caqQueuingParamsTable" {
        return &cISCOATMQOSMIB.Caqqueuingparamstable
    }
    if childYangName == "caqQueuingParamsClassTable" {
        return &cISCOATMQOSMIB.Caqqueuingparamsclasstable
    }
    return nil
}

func (cISCOATMQOSMIB *CISCOATMQOSMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["caqVccParamsTable"] = &cISCOATMQOSMIB.Caqvccparamstable
    children["caqVpcParamsTable"] = &cISCOATMQOSMIB.Caqvpcparamstable
    children["caqQueuingParamsTable"] = &cISCOATMQOSMIB.Caqqueuingparamstable
    children["caqQueuingParamsClassTable"] = &cISCOATMQOSMIB.Caqqueuingparamsclasstable
    return children
}

func (cISCOATMQOSMIB *CISCOATMQOSMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOATMQOSMIB *CISCOATMQOSMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOATMQOSMIB *CISCOATMQOSMIB) GetYangName() string { return "CISCO-ATM-QOS-MIB" }

func (cISCOATMQOSMIB *CISCOATMQOSMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOATMQOSMIB *CISCOATMQOSMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOATMQOSMIB *CISCOATMQOSMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOATMQOSMIB *CISCOATMQOSMIB) SetParent(parent types.Entity) { cISCOATMQOSMIB.parent = parent }

func (cISCOATMQOSMIB *CISCOATMQOSMIB) GetParent() types.Entity { return cISCOATMQOSMIB.parent }

func (cISCOATMQOSMIB *CISCOATMQOSMIB) GetParentYangName() string { return "CISCO-ATM-QOS-MIB" }

// CISCOATMQOSMIB_Caqvccparamstable
// This table is defined to provide QoS information for
// each active ATM VC existing on the interface.
type CISCOATMQOSMIB_Caqvccparamstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This list contains the ATM QoS parameters provided by ciscoAtmQosVccEntry.
    // The type is slice of CISCOATMQOSMIB_Caqvccparamstable_Caqvccparamsentry.
    Caqvccparamsentry []CISCOATMQOSMIB_Caqvccparamstable_Caqvccparamsentry
}

func (caqvccparamstable *CISCOATMQOSMIB_Caqvccparamstable) GetFilter() yfilter.YFilter { return caqvccparamstable.YFilter }

func (caqvccparamstable *CISCOATMQOSMIB_Caqvccparamstable) SetFilter(yf yfilter.YFilter) { caqvccparamstable.YFilter = yf }

func (caqvccparamstable *CISCOATMQOSMIB_Caqvccparamstable) GetGoName(yname string) string {
    if yname == "caqVccParamsEntry" { return "Caqvccparamsentry" }
    return ""
}

func (caqvccparamstable *CISCOATMQOSMIB_Caqvccparamstable) GetSegmentPath() string {
    return "caqVccParamsTable"
}

func (caqvccparamstable *CISCOATMQOSMIB_Caqvccparamstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "caqVccParamsEntry" {
        for _, c := range caqvccparamstable.Caqvccparamsentry {
            if caqvccparamstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOATMQOSMIB_Caqvccparamstable_Caqvccparamsentry{}
        caqvccparamstable.Caqvccparamsentry = append(caqvccparamstable.Caqvccparamsentry, child)
        return &caqvccparamstable.Caqvccparamsentry[len(caqvccparamstable.Caqvccparamsentry)-1]
    }
    return nil
}

func (caqvccparamstable *CISCOATMQOSMIB_Caqvccparamstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range caqvccparamstable.Caqvccparamsentry {
        children[caqvccparamstable.Caqvccparamsentry[i].GetSegmentPath()] = &caqvccparamstable.Caqvccparamsentry[i]
    }
    return children
}

func (caqvccparamstable *CISCOATMQOSMIB_Caqvccparamstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (caqvccparamstable *CISCOATMQOSMIB_Caqvccparamstable) GetBundleName() string { return "cisco_ios_xe" }

func (caqvccparamstable *CISCOATMQOSMIB_Caqvccparamstable) GetYangName() string { return "caqVccParamsTable" }

func (caqvccparamstable *CISCOATMQOSMIB_Caqvccparamstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (caqvccparamstable *CISCOATMQOSMIB_Caqvccparamstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (caqvccparamstable *CISCOATMQOSMIB_Caqvccparamstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (caqvccparamstable *CISCOATMQOSMIB_Caqvccparamstable) SetParent(parent types.Entity) { caqvccparamstable.parent = parent }

func (caqvccparamstable *CISCOATMQOSMIB_Caqvccparamstable) GetParent() types.Entity { return caqvccparamstable.parent }

func (caqvccparamstable *CISCOATMQOSMIB_Caqvccparamstable) GetParentYangName() string { return "CISCO-ATM-QOS-MIB" }

// CISCOATMQOSMIB_Caqvccparamstable_Caqvccparamsentry
// This list contains the ATM QoS parameters provided by
// ciscoAtmQosVccEntry.
type CISCOATMQOSMIB_Caqvccparamstable_Caqvccparamsentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This attribute is a key. The type is string with range: 0..4095. Refers to
    // atm_mib.ATMMIB_Atmvcltable_Atmvclentry_Atmvclvpi
    Atmvclvpi interface{}

    // This attribute is a key. The type is string with range: 0..65535. Refers to
    // atm_mib.ATMMIB_Atmvcltable_Atmvclentry_Atmvclvci
    Atmvclvci interface{}

    // The service category of this virtual circuit connection. The type is
    // AtmServiceCategory.
    Caqvccparamstype interface{}

    // Input Peak Cell Rate (PCR) in kbps with  Cell Loss Priority bit set to 0
    // (clp0). The type is interface{} with range: 0..4294967295.
    Caqvccparamspcrin0 interface{}

    // Number of OAM F5 end to end loopback cells sent through the VCC. The type
    // is interface{} with range: 0..4294967295.
    Caqvccparamspcrin01 interface{}

    // Output Peak Cell Rate (PCR) in kbps with Cell Loss Priority bit set to 0
    // (clp0). The type is interface{} with range: 0..4294967295.
    Caqvccparamspcrout0 interface{}

    // Output Peak Cell Rate (PCR) in kbps with Cell Loss Priority bit set to 1
    // (clp01). The type is interface{} with range: 0..4294967295.
    Caqvccparamspcrout01 interface{}

    // Input Sustained Cell Rate (SCR) in kbps for connection with VBR type of QoS
    // and Cell Loss Priority bit set to 0 (clp0). The type is interface{} with
    // range: 0..4294967295.
    Caqvccparamsscrin0 interface{}

    // Input Sustained Cell Rate (SCR) in kbps for connection with VBR type of QoS
    // and Cell Loss Priority bit set to 1 (clp01). The type is interface{} with
    // range: 0..4294967295.
    Caqvccparamsscrin01 interface{}

    // Output Sustained Cell Rate (SCR) in kbps for connection with VBR type of
    // QoS and Cell Loss Priority bit set to 0 (clp0). The type is interface{}
    // with range: 0..4294967295.
    Caqvccparamsscrout0 interface{}

    // Output Sustained Cell Rate (SCR) in kbps for connection with VBR type of
    // QoS and Cell Loss Priority bit set to 1 (clp01). The type is interface{}
    // with range: 0..4294967295.
    Caqvccparamsscrout01 interface{}

    // Input Burst Cell Size (BCS) for connection with VBR type of QoS and Cell
    // Loss Priority bit set to 0 (clp0). The type is interface{} with range:
    // 0..4294967295.
    Caqvccparamsbcsin0 interface{}

    // Input Burst Cell Size (BCS) for connection with VBR type of QoS and Cell
    // Loss Priority bit set to 1 (clp01). The type is interface{} with range:
    // 0..4294967295.
    Caqvccparamsbcsin01 interface{}

    // Output Burst Cell Size (BCS) for connection with VBR type of QoS and  Cell
    // Loss Priority bit set to 0 (clp0). The type is interface{} with range:
    // 0..4294967295.
    Caqvccparamsbcsout0 interface{}

    // Output Burst Cell Size (BCS) for connection with VBR type of QoS and Cell
    // Loss Priority bit set to 1 (clp01). The type is interface{} with range:
    // 0..4294967295.
    Caqvccparamsbcsout01 interface{}

    // The source of configuration for peak cell rate. The type is
    // VcParamConfigLocation.
    Caqvccparamsinheritlevel interface{}

    // Input Minimum Cell Rate (MCR) in kbps for connection with VBR-nrt or ABR
    // type of QoS. The type is interface{} with range: 0..4294967295.
    Caqvccparamsmcrin interface{}

    // Output Minimum Cell Rate (MCR) in kbps for connection with VBR-nrt or ABR
    // type of QoS. The type is interface{} with range: 0..4294967295.
    Caqvccparamsmcrout interface{}

    // Inverse of rate decrease factor. The type is interface{} with range:
    // 0..4294967295.
    Caqvccparamsinvrdf interface{}

    // Inverse of rate increase factor. The type is interface{} with range:
    // 0..4294967295.
    Caqvccparamsinvrif interface{}

    // The source of configuration for rate factor. The type is
    // VcParamConfigLocation.
    Caqvccparamsrfl interface{}

    // Cell delay variation. The type is interface{} with range: 0..4294967295.
    Caqvccparamscdv interface{}

    // Cell delay variation tolerance. The type is interface{} with range:
    // 0..4294967295.
    Caqvccparamscdvt interface{}

    // Initial cell rate. The type is interface{} with range: 0..4294967295.
    Caqvccparamsicr interface{}

    // Transient buffer exposure. The type is interface{} with range:
    // 0..4294967295.
    Caqvccparamstbe interface{}

    // Fixed round-trip time. The type is interface{} with range: 0..4294967295.
    Caqvccparamsfrtt interface{}

    // Maximum number of tx cells for each forward rm cell. The type is
    // interface{} with range: 0..4294967295.
    Caqvccparamsnrm interface{}

    // Maximum time between forward rm cells. The type is interface{} with range:
    // 0..4294967295.
    Caqvccparamsinvtrm interface{}

    // Inverse of cutoff decrease factor. The type is interface{} with range:
    // 0..4294967295.
    Caqvccparamsinvcdf interface{}

    // Allowed cell rate decrease time factor. The type is interface{} with range:
    // 0..4294967295.
    Caqvccparamsadtf interface{}
}

func (caqvccparamsentry *CISCOATMQOSMIB_Caqvccparamstable_Caqvccparamsentry) GetFilter() yfilter.YFilter { return caqvccparamsentry.YFilter }

func (caqvccparamsentry *CISCOATMQOSMIB_Caqvccparamstable_Caqvccparamsentry) SetFilter(yf yfilter.YFilter) { caqvccparamsentry.YFilter = yf }

func (caqvccparamsentry *CISCOATMQOSMIB_Caqvccparamstable_Caqvccparamsentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "atmVclVpi" { return "Atmvclvpi" }
    if yname == "atmVclVci" { return "Atmvclvci" }
    if yname == "caqVccParamsType" { return "Caqvccparamstype" }
    if yname == "caqVccParamsPcrIn0" { return "Caqvccparamspcrin0" }
    if yname == "caqVccParamsPcrIn01" { return "Caqvccparamspcrin01" }
    if yname == "caqVccParamsPcrOut0" { return "Caqvccparamspcrout0" }
    if yname == "caqVccParamsPcrOut01" { return "Caqvccparamspcrout01" }
    if yname == "caqVccParamsScrIn0" { return "Caqvccparamsscrin0" }
    if yname == "caqVccParamsScrIn01" { return "Caqvccparamsscrin01" }
    if yname == "caqVccParamsScrOut0" { return "Caqvccparamsscrout0" }
    if yname == "caqVccParamsScrOut01" { return "Caqvccparamsscrout01" }
    if yname == "caqVccParamsBcsIn0" { return "Caqvccparamsbcsin0" }
    if yname == "caqVccParamsBcsIn01" { return "Caqvccparamsbcsin01" }
    if yname == "caqVccParamsBcsOut0" { return "Caqvccparamsbcsout0" }
    if yname == "caqVccParamsBcsOut01" { return "Caqvccparamsbcsout01" }
    if yname == "caqVccParamsInheritLevel" { return "Caqvccparamsinheritlevel" }
    if yname == "caqVccParamsMcrIn" { return "Caqvccparamsmcrin" }
    if yname == "caqVccParamsMcrOut" { return "Caqvccparamsmcrout" }
    if yname == "caqVccParamsInvRdf" { return "Caqvccparamsinvrdf" }
    if yname == "caqVccParamsInvRif" { return "Caqvccparamsinvrif" }
    if yname == "caqVccParamsRfl" { return "Caqvccparamsrfl" }
    if yname == "caqVccParamsCdv" { return "Caqvccparamscdv" }
    if yname == "caqVccParamsCdvt" { return "Caqvccparamscdvt" }
    if yname == "caqVccParamsIcr" { return "Caqvccparamsicr" }
    if yname == "caqVccParamsTbe" { return "Caqvccparamstbe" }
    if yname == "caqVccParamsFrtt" { return "Caqvccparamsfrtt" }
    if yname == "caqVccParamsNrm" { return "Caqvccparamsnrm" }
    if yname == "caqVccParamsInvTrm" { return "Caqvccparamsinvtrm" }
    if yname == "caqVccParamsInvCdf" { return "Caqvccparamsinvcdf" }
    if yname == "caqVccParamsAdtf" { return "Caqvccparamsadtf" }
    return ""
}

func (caqvccparamsentry *CISCOATMQOSMIB_Caqvccparamstable_Caqvccparamsentry) GetSegmentPath() string {
    return "caqVccParamsEntry" + "[ifIndex='" + fmt.Sprintf("%v", caqvccparamsentry.Ifindex) + "']" + "[atmVclVpi='" + fmt.Sprintf("%v", caqvccparamsentry.Atmvclvpi) + "']" + "[atmVclVci='" + fmt.Sprintf("%v", caqvccparamsentry.Atmvclvci) + "']"
}

func (caqvccparamsentry *CISCOATMQOSMIB_Caqvccparamstable_Caqvccparamsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (caqvccparamsentry *CISCOATMQOSMIB_Caqvccparamstable_Caqvccparamsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (caqvccparamsentry *CISCOATMQOSMIB_Caqvccparamstable_Caqvccparamsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = caqvccparamsentry.Ifindex
    leafs["atmVclVpi"] = caqvccparamsentry.Atmvclvpi
    leafs["atmVclVci"] = caqvccparamsentry.Atmvclvci
    leafs["caqVccParamsType"] = caqvccparamsentry.Caqvccparamstype
    leafs["caqVccParamsPcrIn0"] = caqvccparamsentry.Caqvccparamspcrin0
    leafs["caqVccParamsPcrIn01"] = caqvccparamsentry.Caqvccparamspcrin01
    leafs["caqVccParamsPcrOut0"] = caqvccparamsentry.Caqvccparamspcrout0
    leafs["caqVccParamsPcrOut01"] = caqvccparamsentry.Caqvccparamspcrout01
    leafs["caqVccParamsScrIn0"] = caqvccparamsentry.Caqvccparamsscrin0
    leafs["caqVccParamsScrIn01"] = caqvccparamsentry.Caqvccparamsscrin01
    leafs["caqVccParamsScrOut0"] = caqvccparamsentry.Caqvccparamsscrout0
    leafs["caqVccParamsScrOut01"] = caqvccparamsentry.Caqvccparamsscrout01
    leafs["caqVccParamsBcsIn0"] = caqvccparamsentry.Caqvccparamsbcsin0
    leafs["caqVccParamsBcsIn01"] = caqvccparamsentry.Caqvccparamsbcsin01
    leafs["caqVccParamsBcsOut0"] = caqvccparamsentry.Caqvccparamsbcsout0
    leafs["caqVccParamsBcsOut01"] = caqvccparamsentry.Caqvccparamsbcsout01
    leafs["caqVccParamsInheritLevel"] = caqvccparamsentry.Caqvccparamsinheritlevel
    leafs["caqVccParamsMcrIn"] = caqvccparamsentry.Caqvccparamsmcrin
    leafs["caqVccParamsMcrOut"] = caqvccparamsentry.Caqvccparamsmcrout
    leafs["caqVccParamsInvRdf"] = caqvccparamsentry.Caqvccparamsinvrdf
    leafs["caqVccParamsInvRif"] = caqvccparamsentry.Caqvccparamsinvrif
    leafs["caqVccParamsRfl"] = caqvccparamsentry.Caqvccparamsrfl
    leafs["caqVccParamsCdv"] = caqvccparamsentry.Caqvccparamscdv
    leafs["caqVccParamsCdvt"] = caqvccparamsentry.Caqvccparamscdvt
    leafs["caqVccParamsIcr"] = caqvccparamsentry.Caqvccparamsicr
    leafs["caqVccParamsTbe"] = caqvccparamsentry.Caqvccparamstbe
    leafs["caqVccParamsFrtt"] = caqvccparamsentry.Caqvccparamsfrtt
    leafs["caqVccParamsNrm"] = caqvccparamsentry.Caqvccparamsnrm
    leafs["caqVccParamsInvTrm"] = caqvccparamsentry.Caqvccparamsinvtrm
    leafs["caqVccParamsInvCdf"] = caqvccparamsentry.Caqvccparamsinvcdf
    leafs["caqVccParamsAdtf"] = caqvccparamsentry.Caqvccparamsadtf
    return leafs
}

func (caqvccparamsentry *CISCOATMQOSMIB_Caqvccparamstable_Caqvccparamsentry) GetBundleName() string { return "cisco_ios_xe" }

func (caqvccparamsentry *CISCOATMQOSMIB_Caqvccparamstable_Caqvccparamsentry) GetYangName() string { return "caqVccParamsEntry" }

func (caqvccparamsentry *CISCOATMQOSMIB_Caqvccparamstable_Caqvccparamsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (caqvccparamsentry *CISCOATMQOSMIB_Caqvccparamstable_Caqvccparamsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (caqvccparamsentry *CISCOATMQOSMIB_Caqvccparamstable_Caqvccparamsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (caqvccparamsentry *CISCOATMQOSMIB_Caqvccparamstable_Caqvccparamsentry) SetParent(parent types.Entity) { caqvccparamsentry.parent = parent }

func (caqvccparamsentry *CISCOATMQOSMIB_Caqvccparamstable_Caqvccparamsentry) GetParent() types.Entity { return caqvccparamsentry.parent }

func (caqvccparamsentry *CISCOATMQOSMIB_Caqvccparamstable_Caqvccparamsentry) GetParentYangName() string { return "caqVccParamsTable" }

// CISCOATMQOSMIB_Caqvpcparamstable
// This table is defined to provide QoS information for
// each active ATM VP existing on the interface.
type CISCOATMQOSMIB_Caqvpcparamstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This list contains the ATM QoS parameters provided by ciscoAtmQosVpcEntry.
    // The type is slice of CISCOATMQOSMIB_Caqvpcparamstable_Caqvpcparamsentry.
    Caqvpcparamsentry []CISCOATMQOSMIB_Caqvpcparamstable_Caqvpcparamsentry
}

func (caqvpcparamstable *CISCOATMQOSMIB_Caqvpcparamstable) GetFilter() yfilter.YFilter { return caqvpcparamstable.YFilter }

func (caqvpcparamstable *CISCOATMQOSMIB_Caqvpcparamstable) SetFilter(yf yfilter.YFilter) { caqvpcparamstable.YFilter = yf }

func (caqvpcparamstable *CISCOATMQOSMIB_Caqvpcparamstable) GetGoName(yname string) string {
    if yname == "caqVpcParamsEntry" { return "Caqvpcparamsentry" }
    return ""
}

func (caqvpcparamstable *CISCOATMQOSMIB_Caqvpcparamstable) GetSegmentPath() string {
    return "caqVpcParamsTable"
}

func (caqvpcparamstable *CISCOATMQOSMIB_Caqvpcparamstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "caqVpcParamsEntry" {
        for _, c := range caqvpcparamstable.Caqvpcparamsentry {
            if caqvpcparamstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOATMQOSMIB_Caqvpcparamstable_Caqvpcparamsentry{}
        caqvpcparamstable.Caqvpcparamsentry = append(caqvpcparamstable.Caqvpcparamsentry, child)
        return &caqvpcparamstable.Caqvpcparamsentry[len(caqvpcparamstable.Caqvpcparamsentry)-1]
    }
    return nil
}

func (caqvpcparamstable *CISCOATMQOSMIB_Caqvpcparamstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range caqvpcparamstable.Caqvpcparamsentry {
        children[caqvpcparamstable.Caqvpcparamsentry[i].GetSegmentPath()] = &caqvpcparamstable.Caqvpcparamsentry[i]
    }
    return children
}

func (caqvpcparamstable *CISCOATMQOSMIB_Caqvpcparamstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (caqvpcparamstable *CISCOATMQOSMIB_Caqvpcparamstable) GetBundleName() string { return "cisco_ios_xe" }

func (caqvpcparamstable *CISCOATMQOSMIB_Caqvpcparamstable) GetYangName() string { return "caqVpcParamsTable" }

func (caqvpcparamstable *CISCOATMQOSMIB_Caqvpcparamstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (caqvpcparamstable *CISCOATMQOSMIB_Caqvpcparamstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (caqvpcparamstable *CISCOATMQOSMIB_Caqvpcparamstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (caqvpcparamstable *CISCOATMQOSMIB_Caqvpcparamstable) SetParent(parent types.Entity) { caqvpcparamstable.parent = parent }

func (caqvpcparamstable *CISCOATMQOSMIB_Caqvpcparamstable) GetParent() types.Entity { return caqvpcparamstable.parent }

func (caqvpcparamstable *CISCOATMQOSMIB_Caqvpcparamstable) GetParentYangName() string { return "CISCO-ATM-QOS-MIB" }

// CISCOATMQOSMIB_Caqvpcparamstable_Caqvpcparamsentry
// This list contains the ATM QoS parameters provided by
// ciscoAtmQosVpcEntry.
type CISCOATMQOSMIB_Caqvpcparamstable_Caqvpcparamsentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This attribute is a key. The type is string with range: 0..4095. Refers to
    // atm_mib.ATMMIB_Atmvpltable_Atmvplentry_Atmvplvpi
    Atmvplvpi interface{}

    // VP state of the current permanent virtual path. The type is VpState.
    Caqvpcparamsvpstate interface{}

    // Maximum rate in kbps at which the associated permanent virtual path can
    // transmit data. The type is interface{} with range: 0..4294967295.
    Caqvpcparamspeakrate interface{}

    // Maximum rate in kbps at which CES VCs can transmit data with the associated
    // permanent virtual path. The type is interface{} with range: 0..4294967295.
    Caqvpcparamscesrate interface{}

    // Number of data VCs currently associated with the permanent virtual path.
    // The type is interface{} with range: 0..65535.
    Caqvpcparamsdatavccount interface{}

    // Number of CES VCs currently associated with the permanent virtual path. The
    // type is interface{} with range: 0..65535.
    Caqvpcparamscesvccount interface{}

    // Vcd for F4 OAM segment processing. The type is interface{} with range:
    // 0..65535.
    Caqvpcparamsvcdf4Seg interface{}

    // Vcd for F4 OAM end to end processing. The type is interface{} with range:
    // 0..65535.
    Caqvpcparamsvcdf4Ete interface{}

    // Sustained cell rate associated with the PVP. The type is interface{} with
    // range: 0..4294967295.
    Caqvpcparamsscr interface{}

    // Maximum burst size associated with the PVP. The type is interface{} with
    // range: 0..4294967295.
    Caqvpcparamsmbs interface{}

    // Bandwidth in Kbps currently currently available on this PVP. The type is
    // interface{} with range: 0..4294967295.
    Caqvpcparamsavailbw interface{}
}

func (caqvpcparamsentry *CISCOATMQOSMIB_Caqvpcparamstable_Caqvpcparamsentry) GetFilter() yfilter.YFilter { return caqvpcparamsentry.YFilter }

func (caqvpcparamsentry *CISCOATMQOSMIB_Caqvpcparamstable_Caqvpcparamsentry) SetFilter(yf yfilter.YFilter) { caqvpcparamsentry.YFilter = yf }

func (caqvpcparamsentry *CISCOATMQOSMIB_Caqvpcparamstable_Caqvpcparamsentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "atmVplVpi" { return "Atmvplvpi" }
    if yname == "caqVpcParamsVpState" { return "Caqvpcparamsvpstate" }
    if yname == "caqVpcParamsPeakRate" { return "Caqvpcparamspeakrate" }
    if yname == "caqVpcParamsCesRate" { return "Caqvpcparamscesrate" }
    if yname == "caqVpcParamsDataVcCount" { return "Caqvpcparamsdatavccount" }
    if yname == "caqVpcParamsCesVcCount" { return "Caqvpcparamscesvccount" }
    if yname == "caqVpcParamsVcdF4Seg" { return "Caqvpcparamsvcdf4Seg" }
    if yname == "caqVpcParamsVcdF4Ete" { return "Caqvpcparamsvcdf4Ete" }
    if yname == "caqVpcParamsScr" { return "Caqvpcparamsscr" }
    if yname == "caqVpcParamsMbs" { return "Caqvpcparamsmbs" }
    if yname == "caqVpcParamsAvailBw" { return "Caqvpcparamsavailbw" }
    return ""
}

func (caqvpcparamsentry *CISCOATMQOSMIB_Caqvpcparamstable_Caqvpcparamsentry) GetSegmentPath() string {
    return "caqVpcParamsEntry" + "[ifIndex='" + fmt.Sprintf("%v", caqvpcparamsentry.Ifindex) + "']" + "[atmVplVpi='" + fmt.Sprintf("%v", caqvpcparamsentry.Atmvplvpi) + "']"
}

func (caqvpcparamsentry *CISCOATMQOSMIB_Caqvpcparamstable_Caqvpcparamsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (caqvpcparamsentry *CISCOATMQOSMIB_Caqvpcparamstable_Caqvpcparamsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (caqvpcparamsentry *CISCOATMQOSMIB_Caqvpcparamstable_Caqvpcparamsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = caqvpcparamsentry.Ifindex
    leafs["atmVplVpi"] = caqvpcparamsentry.Atmvplvpi
    leafs["caqVpcParamsVpState"] = caqvpcparamsentry.Caqvpcparamsvpstate
    leafs["caqVpcParamsPeakRate"] = caqvpcparamsentry.Caqvpcparamspeakrate
    leafs["caqVpcParamsCesRate"] = caqvpcparamsentry.Caqvpcparamscesrate
    leafs["caqVpcParamsDataVcCount"] = caqvpcparamsentry.Caqvpcparamsdatavccount
    leafs["caqVpcParamsCesVcCount"] = caqvpcparamsentry.Caqvpcparamscesvccount
    leafs["caqVpcParamsVcdF4Seg"] = caqvpcparamsentry.Caqvpcparamsvcdf4Seg
    leafs["caqVpcParamsVcdF4Ete"] = caqvpcparamsentry.Caqvpcparamsvcdf4Ete
    leafs["caqVpcParamsScr"] = caqvpcparamsentry.Caqvpcparamsscr
    leafs["caqVpcParamsMbs"] = caqvpcparamsentry.Caqvpcparamsmbs
    leafs["caqVpcParamsAvailBw"] = caqvpcparamsentry.Caqvpcparamsavailbw
    return leafs
}

func (caqvpcparamsentry *CISCOATMQOSMIB_Caqvpcparamstable_Caqvpcparamsentry) GetBundleName() string { return "cisco_ios_xe" }

func (caqvpcparamsentry *CISCOATMQOSMIB_Caqvpcparamstable_Caqvpcparamsentry) GetYangName() string { return "caqVpcParamsEntry" }

func (caqvpcparamsentry *CISCOATMQOSMIB_Caqvpcparamstable_Caqvpcparamsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (caqvpcparamsentry *CISCOATMQOSMIB_Caqvpcparamstable_Caqvpcparamsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (caqvpcparamsentry *CISCOATMQOSMIB_Caqvpcparamstable_Caqvpcparamsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (caqvpcparamsentry *CISCOATMQOSMIB_Caqvpcparamstable_Caqvpcparamsentry) SetParent(parent types.Entity) { caqvpcparamsentry.parent = parent }

func (caqvpcparamsentry *CISCOATMQOSMIB_Caqvpcparamstable_Caqvpcparamsentry) GetParent() types.Entity { return caqvpcparamsentry.parent }

func (caqvpcparamsentry *CISCOATMQOSMIB_Caqvpcparamstable_Caqvpcparamsentry) GetParentYangName() string { return "caqVpcParamsTable" }

// CISCOATMQOSMIB_Caqqueuingparamstable
// This table provides queuing related information
// for a VC existing on an ATM interface.
type CISCOATMQOSMIB_Caqqueuingparamstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This is defined as an entry in caqQueuingParamsTable. The type is slice of
    // CISCOATMQOSMIB_Caqqueuingparamstable_Caqqueuingparamsentry.
    Caqqueuingparamsentry []CISCOATMQOSMIB_Caqqueuingparamstable_Caqqueuingparamsentry
}

func (caqqueuingparamstable *CISCOATMQOSMIB_Caqqueuingparamstable) GetFilter() yfilter.YFilter { return caqqueuingparamstable.YFilter }

func (caqqueuingparamstable *CISCOATMQOSMIB_Caqqueuingparamstable) SetFilter(yf yfilter.YFilter) { caqqueuingparamstable.YFilter = yf }

func (caqqueuingparamstable *CISCOATMQOSMIB_Caqqueuingparamstable) GetGoName(yname string) string {
    if yname == "caqQueuingParamsEntry" { return "Caqqueuingparamsentry" }
    return ""
}

func (caqqueuingparamstable *CISCOATMQOSMIB_Caqqueuingparamstable) GetSegmentPath() string {
    return "caqQueuingParamsTable"
}

func (caqqueuingparamstable *CISCOATMQOSMIB_Caqqueuingparamstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "caqQueuingParamsEntry" {
        for _, c := range caqqueuingparamstable.Caqqueuingparamsentry {
            if caqqueuingparamstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOATMQOSMIB_Caqqueuingparamstable_Caqqueuingparamsentry{}
        caqqueuingparamstable.Caqqueuingparamsentry = append(caqqueuingparamstable.Caqqueuingparamsentry, child)
        return &caqqueuingparamstable.Caqqueuingparamsentry[len(caqqueuingparamstable.Caqqueuingparamsentry)-1]
    }
    return nil
}

func (caqqueuingparamstable *CISCOATMQOSMIB_Caqqueuingparamstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range caqqueuingparamstable.Caqqueuingparamsentry {
        children[caqqueuingparamstable.Caqqueuingparamsentry[i].GetSegmentPath()] = &caqqueuingparamstable.Caqqueuingparamsentry[i]
    }
    return children
}

func (caqqueuingparamstable *CISCOATMQOSMIB_Caqqueuingparamstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (caqqueuingparamstable *CISCOATMQOSMIB_Caqqueuingparamstable) GetBundleName() string { return "cisco_ios_xe" }

func (caqqueuingparamstable *CISCOATMQOSMIB_Caqqueuingparamstable) GetYangName() string { return "caqQueuingParamsTable" }

func (caqqueuingparamstable *CISCOATMQOSMIB_Caqqueuingparamstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (caqqueuingparamstable *CISCOATMQOSMIB_Caqqueuingparamstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (caqqueuingparamstable *CISCOATMQOSMIB_Caqqueuingparamstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (caqqueuingparamstable *CISCOATMQOSMIB_Caqqueuingparamstable) SetParent(parent types.Entity) { caqqueuingparamstable.parent = parent }

func (caqqueuingparamstable *CISCOATMQOSMIB_Caqqueuingparamstable) GetParent() types.Entity { return caqqueuingparamstable.parent }

func (caqqueuingparamstable *CISCOATMQOSMIB_Caqqueuingparamstable) GetParentYangName() string { return "CISCO-ATM-QOS-MIB" }

// CISCOATMQOSMIB_Caqqueuingparamstable_Caqqueuingparamsentry
// This is defined as an entry in caqQueuingParamsTable.
type CISCOATMQOSMIB_Caqqueuingparamstable_Caqqueuingparamsentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This attribute is a key. The type is string with range: 0..4095. Refers to
    // atm_mib.ATMMIB_Atmvcltable_Atmvclentry_Atmvclvpi
    Atmvclvpi interface{}

    // This attribute is a key. The type is string with range: 0..65535. Refers to
    // atm_mib.ATMMIB_Atmvcltable_Atmvclentry_Atmvclvci
    Atmvclvci interface{}

    // Mean Queue Depth associated with the vc.  This value is calculated based on
    // the actual queue depth on the interface and the exponential weighting
    // constant. The type is interface{} with range: 0..4294967295.
    Caqqueuingparamsmeanqdepth interface{}
}

func (caqqueuingparamsentry *CISCOATMQOSMIB_Caqqueuingparamstable_Caqqueuingparamsentry) GetFilter() yfilter.YFilter { return caqqueuingparamsentry.YFilter }

func (caqqueuingparamsentry *CISCOATMQOSMIB_Caqqueuingparamstable_Caqqueuingparamsentry) SetFilter(yf yfilter.YFilter) { caqqueuingparamsentry.YFilter = yf }

func (caqqueuingparamsentry *CISCOATMQOSMIB_Caqqueuingparamstable_Caqqueuingparamsentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "atmVclVpi" { return "Atmvclvpi" }
    if yname == "atmVclVci" { return "Atmvclvci" }
    if yname == "caqQueuingParamsMeanQDepth" { return "Caqqueuingparamsmeanqdepth" }
    return ""
}

func (caqqueuingparamsentry *CISCOATMQOSMIB_Caqqueuingparamstable_Caqqueuingparamsentry) GetSegmentPath() string {
    return "caqQueuingParamsEntry" + "[ifIndex='" + fmt.Sprintf("%v", caqqueuingparamsentry.Ifindex) + "']" + "[atmVclVpi='" + fmt.Sprintf("%v", caqqueuingparamsentry.Atmvclvpi) + "']" + "[atmVclVci='" + fmt.Sprintf("%v", caqqueuingparamsentry.Atmvclvci) + "']"
}

func (caqqueuingparamsentry *CISCOATMQOSMIB_Caqqueuingparamstable_Caqqueuingparamsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (caqqueuingparamsentry *CISCOATMQOSMIB_Caqqueuingparamstable_Caqqueuingparamsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (caqqueuingparamsentry *CISCOATMQOSMIB_Caqqueuingparamstable_Caqqueuingparamsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = caqqueuingparamsentry.Ifindex
    leafs["atmVclVpi"] = caqqueuingparamsentry.Atmvclvpi
    leafs["atmVclVci"] = caqqueuingparamsentry.Atmvclvci
    leafs["caqQueuingParamsMeanQDepth"] = caqqueuingparamsentry.Caqqueuingparamsmeanqdepth
    return leafs
}

func (caqqueuingparamsentry *CISCOATMQOSMIB_Caqqueuingparamstable_Caqqueuingparamsentry) GetBundleName() string { return "cisco_ios_xe" }

func (caqqueuingparamsentry *CISCOATMQOSMIB_Caqqueuingparamstable_Caqqueuingparamsentry) GetYangName() string { return "caqQueuingParamsEntry" }

func (caqqueuingparamsentry *CISCOATMQOSMIB_Caqqueuingparamstable_Caqqueuingparamsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (caqqueuingparamsentry *CISCOATMQOSMIB_Caqqueuingparamstable_Caqqueuingparamsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (caqqueuingparamsentry *CISCOATMQOSMIB_Caqqueuingparamstable_Caqqueuingparamsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (caqqueuingparamsentry *CISCOATMQOSMIB_Caqqueuingparamstable_Caqqueuingparamsentry) SetParent(parent types.Entity) { caqqueuingparamsentry.parent = parent }

func (caqqueuingparamsentry *CISCOATMQOSMIB_Caqqueuingparamstable_Caqqueuingparamsentry) GetParent() types.Entity { return caqqueuingparamsentry.parent }

func (caqqueuingparamsentry *CISCOATMQOSMIB_Caqqueuingparamstable_Caqqueuingparamsentry) GetParentYangName() string { return "caqQueuingParamsTable" }

// CISCOATMQOSMIB_Caqqueuingparamsclasstable
// This table provides queuing information for all 
// queuing classes associating with a VC.
type CISCOATMQOSMIB_Caqqueuingparamsclasstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This is defined as an entry in ciscoAtmQosVcQueuingClassTable to provide
    // queuing information of a specific class. The type is slice of
    // CISCOATMQOSMIB_Caqqueuingparamsclasstable_Caqqueuingparamsclassentry.
    Caqqueuingparamsclassentry []CISCOATMQOSMIB_Caqqueuingparamsclasstable_Caqqueuingparamsclassentry
}

func (caqqueuingparamsclasstable *CISCOATMQOSMIB_Caqqueuingparamsclasstable) GetFilter() yfilter.YFilter { return caqqueuingparamsclasstable.YFilter }

func (caqqueuingparamsclasstable *CISCOATMQOSMIB_Caqqueuingparamsclasstable) SetFilter(yf yfilter.YFilter) { caqqueuingparamsclasstable.YFilter = yf }

func (caqqueuingparamsclasstable *CISCOATMQOSMIB_Caqqueuingparamsclasstable) GetGoName(yname string) string {
    if yname == "caqQueuingParamsClassEntry" { return "Caqqueuingparamsclassentry" }
    return ""
}

func (caqqueuingparamsclasstable *CISCOATMQOSMIB_Caqqueuingparamsclasstable) GetSegmentPath() string {
    return "caqQueuingParamsClassTable"
}

func (caqqueuingparamsclasstable *CISCOATMQOSMIB_Caqqueuingparamsclasstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "caqQueuingParamsClassEntry" {
        for _, c := range caqqueuingparamsclasstable.Caqqueuingparamsclassentry {
            if caqqueuingparamsclasstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOATMQOSMIB_Caqqueuingparamsclasstable_Caqqueuingparamsclassentry{}
        caqqueuingparamsclasstable.Caqqueuingparamsclassentry = append(caqqueuingparamsclasstable.Caqqueuingparamsclassentry, child)
        return &caqqueuingparamsclasstable.Caqqueuingparamsclassentry[len(caqqueuingparamsclasstable.Caqqueuingparamsclassentry)-1]
    }
    return nil
}

func (caqqueuingparamsclasstable *CISCOATMQOSMIB_Caqqueuingparamsclasstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range caqqueuingparamsclasstable.Caqqueuingparamsclassentry {
        children[caqqueuingparamsclasstable.Caqqueuingparamsclassentry[i].GetSegmentPath()] = &caqqueuingparamsclasstable.Caqqueuingparamsclassentry[i]
    }
    return children
}

func (caqqueuingparamsclasstable *CISCOATMQOSMIB_Caqqueuingparamsclasstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (caqqueuingparamsclasstable *CISCOATMQOSMIB_Caqqueuingparamsclasstable) GetBundleName() string { return "cisco_ios_xe" }

func (caqqueuingparamsclasstable *CISCOATMQOSMIB_Caqqueuingparamsclasstable) GetYangName() string { return "caqQueuingParamsClassTable" }

func (caqqueuingparamsclasstable *CISCOATMQOSMIB_Caqqueuingparamsclasstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (caqqueuingparamsclasstable *CISCOATMQOSMIB_Caqqueuingparamsclasstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (caqqueuingparamsclasstable *CISCOATMQOSMIB_Caqqueuingparamsclasstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (caqqueuingparamsclasstable *CISCOATMQOSMIB_Caqqueuingparamsclasstable) SetParent(parent types.Entity) { caqqueuingparamsclasstable.parent = parent }

func (caqqueuingparamsclasstable *CISCOATMQOSMIB_Caqqueuingparamsclasstable) GetParent() types.Entity { return caqqueuingparamsclasstable.parent }

func (caqqueuingparamsclasstable *CISCOATMQOSMIB_Caqqueuingparamsclasstable) GetParentYangName() string { return "CISCO-ATM-QOS-MIB" }

// CISCOATMQOSMIB_Caqqueuingparamsclasstable_Caqqueuingparamsclassentry
// This is defined as an entry in ciscoAtmQosVcQueuingClassTable
// to provide queuing information of a specific class.
type CISCOATMQOSMIB_Caqqueuingparamsclasstable_Caqqueuingparamsclassentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This attribute is a key. The type is string with range: 0..4095. Refers to
    // atm_mib.ATMMIB_Atmvcltable_Atmvclentry_Atmvclvpi
    Atmvclvpi interface{}

    // This attribute is a key. The type is string with range: 0..65535. Refers to
    // atm_mib.ATMMIB_Atmvcltable_Atmvclentry_Atmvclvci
    Atmvclvci interface{}

    // This attribute is a key. A class index, which associates with an IP
    // precedence (0 to 8), is defined to reference individual
    // caqQueuingParamsClassEntry. The type is interface{} with range: 0..8.
    Caqqueuingparamsclassindex interface{}

    // Number of packets dropped when Mean Queue Length is between Minimum
    // Threshold and Maximum Threshold range. The type is interface{} with range:
    // 0..4294967295.
    Caqqueuingparamsclassranddrp interface{}

    // Number of packets dropped because the Mean Queue Depth exceeds the Maximum
    // Threshold value. The type is interface{} with range: 0..4294967295.
    Caqqueuingparamsclasstaildrp interface{}

    // Minimum Threshold value in kbps. The type is interface{} with range:
    // 0..4294967295.
    Caqqueuingparamsclassminthre interface{}

    // Maximum Threshold value in kbps. The type is interface{} with range:
    // 0..4294967295.
    Caqqueuingparamsclassmaxthre interface{}

    // Mark probability denominator.  This is the value used in the calculation of
    // a packet being dropped when the average queue size is between the minimum
    // threshold and the maximum threshold. The type is interface{} with range:
    // 0..4294967295.
    Caqqueuingparamsclassmrkprob interface{}
}

func (caqqueuingparamsclassentry *CISCOATMQOSMIB_Caqqueuingparamsclasstable_Caqqueuingparamsclassentry) GetFilter() yfilter.YFilter { return caqqueuingparamsclassentry.YFilter }

func (caqqueuingparamsclassentry *CISCOATMQOSMIB_Caqqueuingparamsclasstable_Caqqueuingparamsclassentry) SetFilter(yf yfilter.YFilter) { caqqueuingparamsclassentry.YFilter = yf }

func (caqqueuingparamsclassentry *CISCOATMQOSMIB_Caqqueuingparamsclasstable_Caqqueuingparamsclassentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "atmVclVpi" { return "Atmvclvpi" }
    if yname == "atmVclVci" { return "Atmvclvci" }
    if yname == "caqQueuingParamsClassIndex" { return "Caqqueuingparamsclassindex" }
    if yname == "caqQueuingParamsClassRandDrp" { return "Caqqueuingparamsclassranddrp" }
    if yname == "caqQueuingParamsClassTailDrp" { return "Caqqueuingparamsclasstaildrp" }
    if yname == "caqQueuingParamsClassMinThre" { return "Caqqueuingparamsclassminthre" }
    if yname == "caqQueuingParamsClassMaxThre" { return "Caqqueuingparamsclassmaxthre" }
    if yname == "caqQueuingParamsClassMrkProb" { return "Caqqueuingparamsclassmrkprob" }
    return ""
}

func (caqqueuingparamsclassentry *CISCOATMQOSMIB_Caqqueuingparamsclasstable_Caqqueuingparamsclassentry) GetSegmentPath() string {
    return "caqQueuingParamsClassEntry" + "[ifIndex='" + fmt.Sprintf("%v", caqqueuingparamsclassentry.Ifindex) + "']" + "[atmVclVpi='" + fmt.Sprintf("%v", caqqueuingparamsclassentry.Atmvclvpi) + "']" + "[atmVclVci='" + fmt.Sprintf("%v", caqqueuingparamsclassentry.Atmvclvci) + "']" + "[caqQueuingParamsClassIndex='" + fmt.Sprintf("%v", caqqueuingparamsclassentry.Caqqueuingparamsclassindex) + "']"
}

func (caqqueuingparamsclassentry *CISCOATMQOSMIB_Caqqueuingparamsclasstable_Caqqueuingparamsclassentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (caqqueuingparamsclassentry *CISCOATMQOSMIB_Caqqueuingparamsclasstable_Caqqueuingparamsclassentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (caqqueuingparamsclassentry *CISCOATMQOSMIB_Caqqueuingparamsclasstable_Caqqueuingparamsclassentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = caqqueuingparamsclassentry.Ifindex
    leafs["atmVclVpi"] = caqqueuingparamsclassentry.Atmvclvpi
    leafs["atmVclVci"] = caqqueuingparamsclassentry.Atmvclvci
    leafs["caqQueuingParamsClassIndex"] = caqqueuingparamsclassentry.Caqqueuingparamsclassindex
    leafs["caqQueuingParamsClassRandDrp"] = caqqueuingparamsclassentry.Caqqueuingparamsclassranddrp
    leafs["caqQueuingParamsClassTailDrp"] = caqqueuingparamsclassentry.Caqqueuingparamsclasstaildrp
    leafs["caqQueuingParamsClassMinThre"] = caqqueuingparamsclassentry.Caqqueuingparamsclassminthre
    leafs["caqQueuingParamsClassMaxThre"] = caqqueuingparamsclassentry.Caqqueuingparamsclassmaxthre
    leafs["caqQueuingParamsClassMrkProb"] = caqqueuingparamsclassentry.Caqqueuingparamsclassmrkprob
    return leafs
}

func (caqqueuingparamsclassentry *CISCOATMQOSMIB_Caqqueuingparamsclasstable_Caqqueuingparamsclassentry) GetBundleName() string { return "cisco_ios_xe" }

func (caqqueuingparamsclassentry *CISCOATMQOSMIB_Caqqueuingparamsclasstable_Caqqueuingparamsclassentry) GetYangName() string { return "caqQueuingParamsClassEntry" }

func (caqqueuingparamsclassentry *CISCOATMQOSMIB_Caqqueuingparamsclasstable_Caqqueuingparamsclassentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (caqqueuingparamsclassentry *CISCOATMQOSMIB_Caqqueuingparamsclasstable_Caqqueuingparamsclassentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (caqqueuingparamsclassentry *CISCOATMQOSMIB_Caqqueuingparamsclasstable_Caqqueuingparamsclassentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (caqqueuingparamsclassentry *CISCOATMQOSMIB_Caqqueuingparamsclasstable_Caqqueuingparamsclassentry) SetParent(parent types.Entity) { caqqueuingparamsclassentry.parent = parent }

func (caqqueuingparamsclassentry *CISCOATMQOSMIB_Caqqueuingparamsclasstable_Caqqueuingparamsclassentry) GetParent() types.Entity { return caqqueuingparamsclassentry.parent }

func (caqqueuingparamsclassentry *CISCOATMQOSMIB_Caqqueuingparamsclasstable_Caqqueuingparamsclassentry) GetParentYangName() string { return "caqQueuingParamsClassTable" }

