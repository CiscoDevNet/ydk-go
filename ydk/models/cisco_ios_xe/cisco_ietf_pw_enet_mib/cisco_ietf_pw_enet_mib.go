// This MIB describes a model for managing Ethernet  
// point-to-point pseudo wire services over a Packet  
// Switched Network (PSN).
package cisco_ietf_pw_enet_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_ietf_pw_enet_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-IETF-PW-ENET-MIB CISCO-IETF-PW-ENET-MIB}", reflect.TypeOf(CISCOIETFPWENETMIB{}))
    ydk.RegisterEntity("CISCO-IETF-PW-ENET-MIB:CISCO-IETF-PW-ENET-MIB", reflect.TypeOf(CISCOIETFPWENETMIB{}))
}

// CISCOIETFPWENETMIB
type CISCOIETFPWENETMIB struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This table contains the index to the Ethernet tables   associated with this
    // ETH VC, the VLAN configuration and   VLAN mode.
    Cpwvcenettable CISCOIETFPWENETMIB_Cpwvcenettable

    // This table may be used for MPLS PSNs if there is a need   to hold multiple
    // VC, each with different COS, for the same  user service (port + PW VLAN).
    // Such a need may arise if the  MPLS network is capable of L-LSP or E-LSP
    // without multiple  COS capabilities.  Each row is indexed by the cpwVcIndex 
    // and indicate the PRI bits on the packet recieved from the   user port (or
    // VPLS virtual port) that are  classified to this VC. Note that the EXP bit
    // value of the VC  is configured in the CISCO-IETF-PW-MPLS-MIB.
    Cpwvcenetmplsprimappingtable CISCOIETFPWENETMIB_Cpwvcenetmplsprimappingtable

    // This table contains statistical counters specific for   Ethernet PW.
    Cpwvcenetstatstable CISCOIETFPWENETMIB_Cpwvcenetstatstable
}

func (cISCOIETFPWENETMIB *CISCOIETFPWENETMIB) GetFilter() yfilter.YFilter { return cISCOIETFPWENETMIB.YFilter }

func (cISCOIETFPWENETMIB *CISCOIETFPWENETMIB) SetFilter(yf yfilter.YFilter) { cISCOIETFPWENETMIB.YFilter = yf }

func (cISCOIETFPWENETMIB *CISCOIETFPWENETMIB) GetGoName(yname string) string {
    if yname == "cpwVcEnetTable" { return "Cpwvcenettable" }
    if yname == "cpwVcEnetMplsPriMappingTable" { return "Cpwvcenetmplsprimappingtable" }
    if yname == "cpwVcEnetStatsTable" { return "Cpwvcenetstatstable" }
    return ""
}

func (cISCOIETFPWENETMIB *CISCOIETFPWENETMIB) GetSegmentPath() string {
    return "CISCO-IETF-PW-ENET-MIB:CISCO-IETF-PW-ENET-MIB"
}

func (cISCOIETFPWENETMIB *CISCOIETFPWENETMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cpwVcEnetTable" {
        return &cISCOIETFPWENETMIB.Cpwvcenettable
    }
    if childYangName == "cpwVcEnetMplsPriMappingTable" {
        return &cISCOIETFPWENETMIB.Cpwvcenetmplsprimappingtable
    }
    if childYangName == "cpwVcEnetStatsTable" {
        return &cISCOIETFPWENETMIB.Cpwvcenetstatstable
    }
    return nil
}

func (cISCOIETFPWENETMIB *CISCOIETFPWENETMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cpwVcEnetTable"] = &cISCOIETFPWENETMIB.Cpwvcenettable
    children["cpwVcEnetMplsPriMappingTable"] = &cISCOIETFPWENETMIB.Cpwvcenetmplsprimappingtable
    children["cpwVcEnetStatsTable"] = &cISCOIETFPWENETMIB.Cpwvcenetstatstable
    return children
}

func (cISCOIETFPWENETMIB *CISCOIETFPWENETMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOIETFPWENETMIB *CISCOIETFPWENETMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOIETFPWENETMIB *CISCOIETFPWENETMIB) GetYangName() string { return "CISCO-IETF-PW-ENET-MIB" }

func (cISCOIETFPWENETMIB *CISCOIETFPWENETMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOIETFPWENETMIB *CISCOIETFPWENETMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOIETFPWENETMIB *CISCOIETFPWENETMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOIETFPWENETMIB *CISCOIETFPWENETMIB) SetParent(parent types.Entity) { cISCOIETFPWENETMIB.parent = parent }

func (cISCOIETFPWENETMIB *CISCOIETFPWENETMIB) GetParent() types.Entity { return cISCOIETFPWENETMIB.parent }

func (cISCOIETFPWENETMIB *CISCOIETFPWENETMIB) GetParentYangName() string { return "CISCO-IETF-PW-ENET-MIB" }

// CISCOIETFPWENETMIB_Cpwvcenettable
// This table contains the index to the Ethernet tables  
// associated with this ETH VC, the VLAN configuration and  
// VLAN mode.
type CISCOIETFPWENETMIB_Cpwvcenettable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This table is indexed by the same index that was created   for the
    // associated entry in the PW VC Table in the  CISCO-IETF-PW-MIB.  The
    // CpwVcIndex and the cpwVcEnetPwVlan  are used as indexes to allow multiple
    // VLANs to exist on  the same PW.   An entry is created in this table by the
    // agent for every   entry in the cpwVc table with a VcType of 'ethernetVLAN',
    // 'ethernet' or 'ethernetVPLS'. Additional rows may be   created by the
    // operator or the agent if multiple entries  are required for the same VC.  
    // This table provides Ethernet port mapping and VLAN   configuration for each
    // Ethernet VC. The type is slice of
    // CISCOIETFPWENETMIB_Cpwvcenettable_Cpwvcenetentry.
    Cpwvcenetentry []CISCOIETFPWENETMIB_Cpwvcenettable_Cpwvcenetentry
}

func (cpwvcenettable *CISCOIETFPWENETMIB_Cpwvcenettable) GetFilter() yfilter.YFilter { return cpwvcenettable.YFilter }

func (cpwvcenettable *CISCOIETFPWENETMIB_Cpwvcenettable) SetFilter(yf yfilter.YFilter) { cpwvcenettable.YFilter = yf }

func (cpwvcenettable *CISCOIETFPWENETMIB_Cpwvcenettable) GetGoName(yname string) string {
    if yname == "cpwVcEnetEntry" { return "Cpwvcenetentry" }
    return ""
}

func (cpwvcenettable *CISCOIETFPWENETMIB_Cpwvcenettable) GetSegmentPath() string {
    return "cpwVcEnetTable"
}

func (cpwvcenettable *CISCOIETFPWENETMIB_Cpwvcenettable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cpwVcEnetEntry" {
        for _, c := range cpwvcenettable.Cpwvcenetentry {
            if cpwvcenettable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIETFPWENETMIB_Cpwvcenettable_Cpwvcenetentry{}
        cpwvcenettable.Cpwvcenetentry = append(cpwvcenettable.Cpwvcenetentry, child)
        return &cpwvcenettable.Cpwvcenetentry[len(cpwvcenettable.Cpwvcenetentry)-1]
    }
    return nil
}

func (cpwvcenettable *CISCOIETFPWENETMIB_Cpwvcenettable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cpwvcenettable.Cpwvcenetentry {
        children[cpwvcenettable.Cpwvcenetentry[i].GetSegmentPath()] = &cpwvcenettable.Cpwvcenetentry[i]
    }
    return children
}

func (cpwvcenettable *CISCOIETFPWENETMIB_Cpwvcenettable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cpwvcenettable *CISCOIETFPWENETMIB_Cpwvcenettable) GetBundleName() string { return "cisco_ios_xe" }

func (cpwvcenettable *CISCOIETFPWENETMIB_Cpwvcenettable) GetYangName() string { return "cpwVcEnetTable" }

func (cpwvcenettable *CISCOIETFPWENETMIB_Cpwvcenettable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpwvcenettable *CISCOIETFPWENETMIB_Cpwvcenettable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpwvcenettable *CISCOIETFPWENETMIB_Cpwvcenettable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpwvcenettable *CISCOIETFPWENETMIB_Cpwvcenettable) SetParent(parent types.Entity) { cpwvcenettable.parent = parent }

func (cpwvcenettable *CISCOIETFPWENETMIB_Cpwvcenettable) GetParent() types.Entity { return cpwvcenettable.parent }

func (cpwvcenettable *CISCOIETFPWENETMIB_Cpwvcenettable) GetParentYangName() string { return "CISCO-IETF-PW-ENET-MIB" }

// CISCOIETFPWENETMIB_Cpwvcenettable_Cpwvcenetentry
// This table is indexed by the same index that was created  
// for the associated entry in the PW VC Table in the 
// CISCO-IETF-PW-MIB.  The CpwVcIndex and the cpwVcEnetPwVlan 
// are used as indexes to allow multiple VLANs to exist on 
// the same PW. 
// 
// An entry is created in this table by the agent for every  
// entry in the cpwVc table with a VcType of 'ethernetVLAN', 
// 'ethernet' or 'ethernetVPLS'. Additional rows may be  
// created by the operator or the agent if multiple entries 
// are required for the same VC. 
// 
// This table provides Ethernet port mapping and VLAN  
// configuration for each Ethernet VC.
type CISCOIETFPWENETMIB_Cpwvcenettable_Cpwvcenetentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to cisco_ietf_pw_mib.CISCOIETFPWMIB_Cpwvctable_Cpwvcentry_Cpwvcindex
    Cpwvcindex interface{}

    // This attribute is a key. This Object defines the VLAN on the VC. The value
    // of 4097  is used if the object is not applicable, for example when  mapping
    // all packets from an Ethernet port to this VC.  The value of 4096 is used to
    // indicate untagged frames (at   least from the PW point of view), for
    // example if   cpwVcEnetVlanMode is equal 'removeVLAN' or when  
    // cpwVcEnetVlanMode equal 'noChange' and cpwVcEnetPortVlan  is equal 4096.
    // The type is interface{} with range: 0..4097.
    Cpwvcenetpwvlan interface{}

    // Indicate the mode of VLAN handling between the port   associated to the VC
    // and the VC encapsulation itself.   - 'other' indicate operation that is not
    // defined by    this MIB.   - 'portBased' indicates that the forwarder will
    // forward    packets between the port and the PW independent of their   
    // structure.   - 'noChange' indicates that the VC contains the original    
    // user VLAN, as specified in cpwVcEnetPortVlan.   - 'changeVlan' indicates
    // that the VLAN field on the VC     may be different than the VLAN field on
    // the user's     port.   - 'removeVlan' indicates that the encapsulation on
    // the     VC does not include the original VLAN field. Note     that PRI bits
    // transparency is lost in this case.   - 'addVlan' indicate that a VLAN field
    // will be added    on the PSN bound direction. cpwVcEnetPwVlan indicate   
    // the value that will be added.    - 'removeVlan', 'addVlan' and 'changeVlan'
    // implementation    is not required. . The type is Cpwvcenetvlanmode.
    Cpwvcenetvlanmode interface{}

    // This object define the VLAN value on the physical port (or   VPLS virtual
    // port) if a change is required to the VLAN value  between the VC and the
    // physical/virtual port.   The value of this object can be ignored if the
    // whole traffic   from the port is forwarded to one VC independent of the  
    // tagging on the port, but it is RECOMENDED that the value in  this case will
    // be '4097' indicating not relevant.   It MUST be equal to cpwVcEnetPwVlan if
    // 'noChange' mode  is used.   The value 4096 indicate that no VLAN (i.e.
    // untagged   frames) on the port are associated to this VC. This   allows the
    // same behaviors as assigning 'Default VLAN'   to un-tagged frames. . The
    // type is interface{} with range: 0..4097.
    Cpwvcenetportvlan interface{}

    // It is sometimes convenient to model the VC PW as a   virtual interface in
    // the ifTable. In these cases this   object hold the value of the ifIndex in
    // the ifTable   representing this VC PW. A value of zero indicate no such  
    // association or association is not yet known. The type is interface{} with
    // range: 0..2147483647.
    Cpwvcenetvcifindex interface{}

    // This object is used to specify the ifIndex of the ETHERNET  port associated
    // with this VC for point-to-point Ethernet   service, or the ifIndex of the
    // virtual interface of the VPLS   instance associated with the PW if the
    // service is VPLS. Two   rows in this table can point to the same ifIndex
    // only if:   1) It is required to support multiple COS on a MPLS PSN      for
    // the same service (i.e.: a combination of ports and      VLANs) by the use
    // of multiple VC, each with a different     COS.   2) There is no overlap of
    // VLAN values specified in      cpwVcEnetPortVlan that are associated with
    // this port.   A value of zero indicate that association to an ifIndex is 
    // not yet known. The type is interface{} with range: 0..2147483647.
    Cpwvcenetportifindex interface{}

    // Enable creating, deleting and modifying this row. The type is RowStatus.
    Cpwvcenetrowstatus interface{}

    // Indicates the storage type of this row. The type is StorageType.
    Cpwvcenetstoragetype interface{}
}

func (cpwvcenetentry *CISCOIETFPWENETMIB_Cpwvcenettable_Cpwvcenetentry) GetFilter() yfilter.YFilter { return cpwvcenetentry.YFilter }

func (cpwvcenetentry *CISCOIETFPWENETMIB_Cpwvcenettable_Cpwvcenetentry) SetFilter(yf yfilter.YFilter) { cpwvcenetentry.YFilter = yf }

func (cpwvcenetentry *CISCOIETFPWENETMIB_Cpwvcenettable_Cpwvcenetentry) GetGoName(yname string) string {
    if yname == "cpwVcIndex" { return "Cpwvcindex" }
    if yname == "cpwVcEnetPwVlan" { return "Cpwvcenetpwvlan" }
    if yname == "cpwVcEnetVlanMode" { return "Cpwvcenetvlanmode" }
    if yname == "cpwVcEnetPortVlan" { return "Cpwvcenetportvlan" }
    if yname == "cpwVcEnetVcIfIndex" { return "Cpwvcenetvcifindex" }
    if yname == "cpwVcEnetPortIfIndex" { return "Cpwvcenetportifindex" }
    if yname == "cpwVcEnetRowStatus" { return "Cpwvcenetrowstatus" }
    if yname == "cpwVcEnetStorageType" { return "Cpwvcenetstoragetype" }
    return ""
}

func (cpwvcenetentry *CISCOIETFPWENETMIB_Cpwvcenettable_Cpwvcenetentry) GetSegmentPath() string {
    return "cpwVcEnetEntry" + "[cpwVcIndex='" + fmt.Sprintf("%v", cpwvcenetentry.Cpwvcindex) + "']" + "[cpwVcEnetPwVlan='" + fmt.Sprintf("%v", cpwvcenetentry.Cpwvcenetpwvlan) + "']"
}

func (cpwvcenetentry *CISCOIETFPWENETMIB_Cpwvcenettable_Cpwvcenetentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cpwvcenetentry *CISCOIETFPWENETMIB_Cpwvcenettable_Cpwvcenetentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cpwvcenetentry *CISCOIETFPWENETMIB_Cpwvcenettable_Cpwvcenetentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cpwVcIndex"] = cpwvcenetentry.Cpwvcindex
    leafs["cpwVcEnetPwVlan"] = cpwvcenetentry.Cpwvcenetpwvlan
    leafs["cpwVcEnetVlanMode"] = cpwvcenetentry.Cpwvcenetvlanmode
    leafs["cpwVcEnetPortVlan"] = cpwvcenetentry.Cpwvcenetportvlan
    leafs["cpwVcEnetVcIfIndex"] = cpwvcenetentry.Cpwvcenetvcifindex
    leafs["cpwVcEnetPortIfIndex"] = cpwvcenetentry.Cpwvcenetportifindex
    leafs["cpwVcEnetRowStatus"] = cpwvcenetentry.Cpwvcenetrowstatus
    leafs["cpwVcEnetStorageType"] = cpwvcenetentry.Cpwvcenetstoragetype
    return leafs
}

func (cpwvcenetentry *CISCOIETFPWENETMIB_Cpwvcenettable_Cpwvcenetentry) GetBundleName() string { return "cisco_ios_xe" }

func (cpwvcenetentry *CISCOIETFPWENETMIB_Cpwvcenettable_Cpwvcenetentry) GetYangName() string { return "cpwVcEnetEntry" }

func (cpwvcenetentry *CISCOIETFPWENETMIB_Cpwvcenettable_Cpwvcenetentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpwvcenetentry *CISCOIETFPWENETMIB_Cpwvcenettable_Cpwvcenetentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpwvcenetentry *CISCOIETFPWENETMIB_Cpwvcenettable_Cpwvcenetentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpwvcenetentry *CISCOIETFPWENETMIB_Cpwvcenettable_Cpwvcenetentry) SetParent(parent types.Entity) { cpwvcenetentry.parent = parent }

func (cpwvcenetentry *CISCOIETFPWENETMIB_Cpwvcenettable_Cpwvcenetentry) GetParent() types.Entity { return cpwvcenetentry.parent }

func (cpwvcenetentry *CISCOIETFPWENETMIB_Cpwvcenettable_Cpwvcenetentry) GetParentYangName() string { return "cpwVcEnetTable" }

// CISCOIETFPWENETMIB_Cpwvcenettable_Cpwvcenetentry_Cpwvcenetvlanmode represents   is not required. 
type CISCOIETFPWENETMIB_Cpwvcenettable_Cpwvcenetentry_Cpwvcenetvlanmode string

const (
    CISCOIETFPWENETMIB_Cpwvcenettable_Cpwvcenetentry_Cpwvcenetvlanmode_other CISCOIETFPWENETMIB_Cpwvcenettable_Cpwvcenetentry_Cpwvcenetvlanmode = "other"

    CISCOIETFPWENETMIB_Cpwvcenettable_Cpwvcenetentry_Cpwvcenetvlanmode_portBased CISCOIETFPWENETMIB_Cpwvcenettable_Cpwvcenetentry_Cpwvcenetvlanmode = "portBased"

    CISCOIETFPWENETMIB_Cpwvcenettable_Cpwvcenetentry_Cpwvcenetvlanmode_noChange CISCOIETFPWENETMIB_Cpwvcenettable_Cpwvcenetentry_Cpwvcenetvlanmode = "noChange"

    CISCOIETFPWENETMIB_Cpwvcenettable_Cpwvcenetentry_Cpwvcenetvlanmode_changeVlan CISCOIETFPWENETMIB_Cpwvcenettable_Cpwvcenetentry_Cpwvcenetvlanmode = "changeVlan"

    CISCOIETFPWENETMIB_Cpwvcenettable_Cpwvcenetentry_Cpwvcenetvlanmode_addVlan CISCOIETFPWENETMIB_Cpwvcenettable_Cpwvcenetentry_Cpwvcenetvlanmode = "addVlan"

    CISCOIETFPWENETMIB_Cpwvcenettable_Cpwvcenetentry_Cpwvcenetvlanmode_removeVlan CISCOIETFPWENETMIB_Cpwvcenettable_Cpwvcenetentry_Cpwvcenetvlanmode = "removeVlan"
)

// CISCOIETFPWENETMIB_Cpwvcenetmplsprimappingtable
// This table may be used for MPLS PSNs if there is a need  
// to hold multiple VC, each with different COS, for the same 
// user service (port + PW VLAN). Such a need may arise if the 
// MPLS network is capable of L-LSP or E-LSP without multiple 
// COS capabilities.  Each row is indexed by the cpwVcIndex  
// and indicate the PRI bits on the packet recieved from the  
// user port (or VPLS virtual port) that are 
// classified to this VC. Note that the EXP bit value of the VC 
// is configured in the CISCO-IETF-PW-MPLS-MIB.
type CISCOIETFPWENETMIB_Cpwvcenetmplsprimappingtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry is created if special classification based on   the PRI bits is
    // required for this VC. The type is slice of
    // CISCOIETFPWENETMIB_Cpwvcenetmplsprimappingtable_Cpwvcenetmplsprimappingtableentry.
    Cpwvcenetmplsprimappingtableentry []CISCOIETFPWENETMIB_Cpwvcenetmplsprimappingtable_Cpwvcenetmplsprimappingtableentry
}

func (cpwvcenetmplsprimappingtable *CISCOIETFPWENETMIB_Cpwvcenetmplsprimappingtable) GetFilter() yfilter.YFilter { return cpwvcenetmplsprimappingtable.YFilter }

func (cpwvcenetmplsprimappingtable *CISCOIETFPWENETMIB_Cpwvcenetmplsprimappingtable) SetFilter(yf yfilter.YFilter) { cpwvcenetmplsprimappingtable.YFilter = yf }

func (cpwvcenetmplsprimappingtable *CISCOIETFPWENETMIB_Cpwvcenetmplsprimappingtable) GetGoName(yname string) string {
    if yname == "cpwVcEnetMplsPriMappingTableEntry" { return "Cpwvcenetmplsprimappingtableentry" }
    return ""
}

func (cpwvcenetmplsprimappingtable *CISCOIETFPWENETMIB_Cpwvcenetmplsprimappingtable) GetSegmentPath() string {
    return "cpwVcEnetMplsPriMappingTable"
}

func (cpwvcenetmplsprimappingtable *CISCOIETFPWENETMIB_Cpwvcenetmplsprimappingtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cpwVcEnetMplsPriMappingTableEntry" {
        for _, c := range cpwvcenetmplsprimappingtable.Cpwvcenetmplsprimappingtableentry {
            if cpwvcenetmplsprimappingtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIETFPWENETMIB_Cpwvcenetmplsprimappingtable_Cpwvcenetmplsprimappingtableentry{}
        cpwvcenetmplsprimappingtable.Cpwvcenetmplsprimappingtableentry = append(cpwvcenetmplsprimappingtable.Cpwvcenetmplsprimappingtableentry, child)
        return &cpwvcenetmplsprimappingtable.Cpwvcenetmplsprimappingtableentry[len(cpwvcenetmplsprimappingtable.Cpwvcenetmplsprimappingtableentry)-1]
    }
    return nil
}

func (cpwvcenetmplsprimappingtable *CISCOIETFPWENETMIB_Cpwvcenetmplsprimappingtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cpwvcenetmplsprimappingtable.Cpwvcenetmplsprimappingtableentry {
        children[cpwvcenetmplsprimappingtable.Cpwvcenetmplsprimappingtableentry[i].GetSegmentPath()] = &cpwvcenetmplsprimappingtable.Cpwvcenetmplsprimappingtableentry[i]
    }
    return children
}

func (cpwvcenetmplsprimappingtable *CISCOIETFPWENETMIB_Cpwvcenetmplsprimappingtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cpwvcenetmplsprimappingtable *CISCOIETFPWENETMIB_Cpwvcenetmplsprimappingtable) GetBundleName() string { return "cisco_ios_xe" }

func (cpwvcenetmplsprimappingtable *CISCOIETFPWENETMIB_Cpwvcenetmplsprimappingtable) GetYangName() string { return "cpwVcEnetMplsPriMappingTable" }

func (cpwvcenetmplsprimappingtable *CISCOIETFPWENETMIB_Cpwvcenetmplsprimappingtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpwvcenetmplsprimappingtable *CISCOIETFPWENETMIB_Cpwvcenetmplsprimappingtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpwvcenetmplsprimappingtable *CISCOIETFPWENETMIB_Cpwvcenetmplsprimappingtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpwvcenetmplsprimappingtable *CISCOIETFPWENETMIB_Cpwvcenetmplsprimappingtable) SetParent(parent types.Entity) { cpwvcenetmplsprimappingtable.parent = parent }

func (cpwvcenetmplsprimappingtable *CISCOIETFPWENETMIB_Cpwvcenetmplsprimappingtable) GetParent() types.Entity { return cpwvcenetmplsprimappingtable.parent }

func (cpwvcenetmplsprimappingtable *CISCOIETFPWENETMIB_Cpwvcenetmplsprimappingtable) GetParentYangName() string { return "CISCO-IETF-PW-ENET-MIB" }

// CISCOIETFPWENETMIB_Cpwvcenetmplsprimappingtable_Cpwvcenetmplsprimappingtableentry
// Each entry is created if special classification based on  
// the PRI bits is required for this VC.
type CISCOIETFPWENETMIB_Cpwvcenetmplsprimappingtable_Cpwvcenetmplsprimappingtableentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to cisco_ietf_pw_mib.CISCOIETFPWMIB_Cpwvctable_Cpwvcentry_Cpwvcindex
    Cpwvcindex interface{}

    // This object defines the groups of user PRI mapped into  this VC. Each bit
    // set indicates that this user priority   is assigned to this VC.   The value
    // 'untagged' is used to indicate that untagged   frames are also associated
    // to this VC.   This object allow the use of different PSN COS based on  
    // user marking of PRI bits in MPLS PSN with L-LSP or   E-LSP without multiple
    // COS support. In all other cases,   the default value MUST be used.   It is
    // REQUIRED that there is no overlap on this object   between rows serving the
    // same service (port+ PW VLAN).   In case of missing BIT configuration
    // between rows to   the same service, incoming packets with PRI marking not  
    // configured should be handled by the VC with the lowest   COS. . The type is
    // map[string]bool.
    Cpwvcenetmplsprimapping interface{}

    // Enable creating, deleting and modifying this row. The type is RowStatus.
    Cpwvcenetmplsprimappingrowstatus interface{}

    // Indicates the storage type of this row. The type is StorageType.
    Cpwvcenetmplsprimappingstoragetype interface{}
}

func (cpwvcenetmplsprimappingtableentry *CISCOIETFPWENETMIB_Cpwvcenetmplsprimappingtable_Cpwvcenetmplsprimappingtableentry) GetFilter() yfilter.YFilter { return cpwvcenetmplsprimappingtableentry.YFilter }

func (cpwvcenetmplsprimappingtableentry *CISCOIETFPWENETMIB_Cpwvcenetmplsprimappingtable_Cpwvcenetmplsprimappingtableentry) SetFilter(yf yfilter.YFilter) { cpwvcenetmplsprimappingtableentry.YFilter = yf }

func (cpwvcenetmplsprimappingtableentry *CISCOIETFPWENETMIB_Cpwvcenetmplsprimappingtable_Cpwvcenetmplsprimappingtableentry) GetGoName(yname string) string {
    if yname == "cpwVcIndex" { return "Cpwvcindex" }
    if yname == "cpwVcEnetMplsPriMapping" { return "Cpwvcenetmplsprimapping" }
    if yname == "cpwVcEnetMplsPriMappingRowStatus" { return "Cpwvcenetmplsprimappingrowstatus" }
    if yname == "cpwVcEnetMplsPriMappingStorageType" { return "Cpwvcenetmplsprimappingstoragetype" }
    return ""
}

func (cpwvcenetmplsprimappingtableentry *CISCOIETFPWENETMIB_Cpwvcenetmplsprimappingtable_Cpwvcenetmplsprimappingtableentry) GetSegmentPath() string {
    return "cpwVcEnetMplsPriMappingTableEntry" + "[cpwVcIndex='" + fmt.Sprintf("%v", cpwvcenetmplsprimappingtableentry.Cpwvcindex) + "']"
}

func (cpwvcenetmplsprimappingtableentry *CISCOIETFPWENETMIB_Cpwvcenetmplsprimappingtable_Cpwvcenetmplsprimappingtableentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cpwvcenetmplsprimappingtableentry *CISCOIETFPWENETMIB_Cpwvcenetmplsprimappingtable_Cpwvcenetmplsprimappingtableentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cpwvcenetmplsprimappingtableentry *CISCOIETFPWENETMIB_Cpwvcenetmplsprimappingtable_Cpwvcenetmplsprimappingtableentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cpwVcIndex"] = cpwvcenetmplsprimappingtableentry.Cpwvcindex
    leafs["cpwVcEnetMplsPriMapping"] = cpwvcenetmplsprimappingtableentry.Cpwvcenetmplsprimapping
    leafs["cpwVcEnetMplsPriMappingRowStatus"] = cpwvcenetmplsprimappingtableentry.Cpwvcenetmplsprimappingrowstatus
    leafs["cpwVcEnetMplsPriMappingStorageType"] = cpwvcenetmplsprimappingtableentry.Cpwvcenetmplsprimappingstoragetype
    return leafs
}

func (cpwvcenetmplsprimappingtableentry *CISCOIETFPWENETMIB_Cpwvcenetmplsprimappingtable_Cpwvcenetmplsprimappingtableentry) GetBundleName() string { return "cisco_ios_xe" }

func (cpwvcenetmplsprimappingtableentry *CISCOIETFPWENETMIB_Cpwvcenetmplsprimappingtable_Cpwvcenetmplsprimappingtableentry) GetYangName() string { return "cpwVcEnetMplsPriMappingTableEntry" }

func (cpwvcenetmplsprimappingtableentry *CISCOIETFPWENETMIB_Cpwvcenetmplsprimappingtable_Cpwvcenetmplsprimappingtableentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpwvcenetmplsprimappingtableentry *CISCOIETFPWENETMIB_Cpwvcenetmplsprimappingtable_Cpwvcenetmplsprimappingtableentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpwvcenetmplsprimappingtableentry *CISCOIETFPWENETMIB_Cpwvcenetmplsprimappingtable_Cpwvcenetmplsprimappingtableentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpwvcenetmplsprimappingtableentry *CISCOIETFPWENETMIB_Cpwvcenetmplsprimappingtable_Cpwvcenetmplsprimappingtableentry) SetParent(parent types.Entity) { cpwvcenetmplsprimappingtableentry.parent = parent }

func (cpwvcenetmplsprimappingtableentry *CISCOIETFPWENETMIB_Cpwvcenetmplsprimappingtable_Cpwvcenetmplsprimappingtableentry) GetParent() types.Entity { return cpwvcenetmplsprimappingtableentry.parent }

func (cpwvcenetmplsprimappingtableentry *CISCOIETFPWENETMIB_Cpwvcenetmplsprimappingtable_Cpwvcenetmplsprimappingtableentry) GetParentYangName() string { return "cpwVcEnetMplsPriMappingTable" }

// CISCOIETFPWENETMIB_Cpwvcenetstatstable
// This table contains statistical counters specific for  
// Ethernet PW.
type CISCOIETFPWENETMIB_Cpwvcenetstatstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry represents the statistics gathered for the   VC carrying the
    // Ethernet packets since this VC was   first created in the cpwVcEnetTable.
    // The type is slice of
    // CISCOIETFPWENETMIB_Cpwvcenetstatstable_Cpwvcenetstatsentry.
    Cpwvcenetstatsentry []CISCOIETFPWENETMIB_Cpwvcenetstatstable_Cpwvcenetstatsentry
}

func (cpwvcenetstatstable *CISCOIETFPWENETMIB_Cpwvcenetstatstable) GetFilter() yfilter.YFilter { return cpwvcenetstatstable.YFilter }

func (cpwvcenetstatstable *CISCOIETFPWENETMIB_Cpwvcenetstatstable) SetFilter(yf yfilter.YFilter) { cpwvcenetstatstable.YFilter = yf }

func (cpwvcenetstatstable *CISCOIETFPWENETMIB_Cpwvcenetstatstable) GetGoName(yname string) string {
    if yname == "cpwVcEnetStatsEntry" { return "Cpwvcenetstatsentry" }
    return ""
}

func (cpwvcenetstatstable *CISCOIETFPWENETMIB_Cpwvcenetstatstable) GetSegmentPath() string {
    return "cpwVcEnetStatsTable"
}

func (cpwvcenetstatstable *CISCOIETFPWENETMIB_Cpwvcenetstatstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cpwVcEnetStatsEntry" {
        for _, c := range cpwvcenetstatstable.Cpwvcenetstatsentry {
            if cpwvcenetstatstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIETFPWENETMIB_Cpwvcenetstatstable_Cpwvcenetstatsentry{}
        cpwvcenetstatstable.Cpwvcenetstatsentry = append(cpwvcenetstatstable.Cpwvcenetstatsentry, child)
        return &cpwvcenetstatstable.Cpwvcenetstatsentry[len(cpwvcenetstatstable.Cpwvcenetstatsentry)-1]
    }
    return nil
}

func (cpwvcenetstatstable *CISCOIETFPWENETMIB_Cpwvcenetstatstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cpwvcenetstatstable.Cpwvcenetstatsentry {
        children[cpwvcenetstatstable.Cpwvcenetstatsentry[i].GetSegmentPath()] = &cpwvcenetstatstable.Cpwvcenetstatsentry[i]
    }
    return children
}

func (cpwvcenetstatstable *CISCOIETFPWENETMIB_Cpwvcenetstatstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cpwvcenetstatstable *CISCOIETFPWENETMIB_Cpwvcenetstatstable) GetBundleName() string { return "cisco_ios_xe" }

func (cpwvcenetstatstable *CISCOIETFPWENETMIB_Cpwvcenetstatstable) GetYangName() string { return "cpwVcEnetStatsTable" }

func (cpwvcenetstatstable *CISCOIETFPWENETMIB_Cpwvcenetstatstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpwvcenetstatstable *CISCOIETFPWENETMIB_Cpwvcenetstatstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpwvcenetstatstable *CISCOIETFPWENETMIB_Cpwvcenetstatstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpwvcenetstatstable *CISCOIETFPWENETMIB_Cpwvcenetstatstable) SetParent(parent types.Entity) { cpwvcenetstatstable.parent = parent }

func (cpwvcenetstatstable *CISCOIETFPWENETMIB_Cpwvcenetstatstable) GetParent() types.Entity { return cpwvcenetstatstable.parent }

func (cpwvcenetstatstable *CISCOIETFPWENETMIB_Cpwvcenetstatstable) GetParentYangName() string { return "CISCO-IETF-PW-ENET-MIB" }

// CISCOIETFPWENETMIB_Cpwvcenetstatstable_Cpwvcenetstatsentry
// Each entry represents the statistics gathered for the  
// VC carrying the Ethernet packets since this VC was  
// first created in the cpwVcEnetTable.
type CISCOIETFPWENETMIB_Cpwvcenetstatstable_Cpwvcenetstatsentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to cisco_ietf_pw_mib.CISCOIETFPWMIB_Cpwvctable_Cpwvcentry_Cpwvcindex
    Cpwvcindex interface{}

    // The number of packets received (from the PSN) on this VC with   an illegal
    // VLAN field, missing VLAN field that was expected, or   A VLAN field when it
    // was not expected. This counter is not   relevant if the VC type is
    // 'ethernet' (i.e. raw mode), and   should be set to 0 by the agent to
    // indicate this. The type is interface{} with range: 0..18446744073709551615.
    Cpwvcenetstatsillegalvlan interface{}

    // The number of packets that were received with an illegal   Ethernet packet
    // length on this VC. An illegal length is defined  as being greater than the
    // value in the advertised maximum MTU   supported, or shorter than the
    // allowed Ethernet packet size. The type is interface{} with range:
    // 0..18446744073709551615.
    Cpwvcenetstatsillegallength interface{}
}

func (cpwvcenetstatsentry *CISCOIETFPWENETMIB_Cpwvcenetstatstable_Cpwvcenetstatsentry) GetFilter() yfilter.YFilter { return cpwvcenetstatsentry.YFilter }

func (cpwvcenetstatsentry *CISCOIETFPWENETMIB_Cpwvcenetstatstable_Cpwvcenetstatsentry) SetFilter(yf yfilter.YFilter) { cpwvcenetstatsentry.YFilter = yf }

func (cpwvcenetstatsentry *CISCOIETFPWENETMIB_Cpwvcenetstatstable_Cpwvcenetstatsentry) GetGoName(yname string) string {
    if yname == "cpwVcIndex" { return "Cpwvcindex" }
    if yname == "cpwVcEnetStatsIllegalVlan" { return "Cpwvcenetstatsillegalvlan" }
    if yname == "cpwVcEnetStatsIllegalLength" { return "Cpwvcenetstatsillegallength" }
    return ""
}

func (cpwvcenetstatsentry *CISCOIETFPWENETMIB_Cpwvcenetstatstable_Cpwvcenetstatsentry) GetSegmentPath() string {
    return "cpwVcEnetStatsEntry" + "[cpwVcIndex='" + fmt.Sprintf("%v", cpwvcenetstatsentry.Cpwvcindex) + "']"
}

func (cpwvcenetstatsentry *CISCOIETFPWENETMIB_Cpwvcenetstatstable_Cpwvcenetstatsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cpwvcenetstatsentry *CISCOIETFPWENETMIB_Cpwvcenetstatstable_Cpwvcenetstatsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cpwvcenetstatsentry *CISCOIETFPWENETMIB_Cpwvcenetstatstable_Cpwvcenetstatsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cpwVcIndex"] = cpwvcenetstatsentry.Cpwvcindex
    leafs["cpwVcEnetStatsIllegalVlan"] = cpwvcenetstatsentry.Cpwvcenetstatsillegalvlan
    leafs["cpwVcEnetStatsIllegalLength"] = cpwvcenetstatsentry.Cpwvcenetstatsillegallength
    return leafs
}

func (cpwvcenetstatsentry *CISCOIETFPWENETMIB_Cpwvcenetstatstable_Cpwvcenetstatsentry) GetBundleName() string { return "cisco_ios_xe" }

func (cpwvcenetstatsentry *CISCOIETFPWENETMIB_Cpwvcenetstatstable_Cpwvcenetstatsentry) GetYangName() string { return "cpwVcEnetStatsEntry" }

func (cpwvcenetstatsentry *CISCOIETFPWENETMIB_Cpwvcenetstatstable_Cpwvcenetstatsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpwvcenetstatsentry *CISCOIETFPWENETMIB_Cpwvcenetstatstable_Cpwvcenetstatsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpwvcenetstatsentry *CISCOIETFPWENETMIB_Cpwvcenetstatstable_Cpwvcenetstatsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpwvcenetstatsentry *CISCOIETFPWENETMIB_Cpwvcenetstatstable_Cpwvcenetstatsentry) SetParent(parent types.Entity) { cpwvcenetstatsentry.parent = parent }

func (cpwvcenetstatsentry *CISCOIETFPWENETMIB_Cpwvcenetstatstable_Cpwvcenetstatsentry) GetParent() types.Entity { return cpwvcenetstatsentry.parent }

func (cpwvcenetstatsentry *CISCOIETFPWENETMIB_Cpwvcenetstatstable_Cpwvcenetstatsentry) GetParentYangName() string { return "cpwVcEnetStatsTable" }

