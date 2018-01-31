// IGMP Filter configuration MIB. This MIB provides a
// mechanism for users to configure the system to intercept 
// IGMP joins for IP Multicast groups identified in this 
// MIB and only allow certain ports to join certain 
// multicast groups.
package cisco_igmp_filter_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_igmp_filter_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-IGMP-FILTER-MIB CISCO-IGMP-FILTER-MIB}", reflect.TypeOf(CISCOIGMPFILTERMIB{}))
    ydk.RegisterEntity("CISCO-IGMP-FILTER-MIB:CISCO-IGMP-FILTER-MIB", reflect.TypeOf(CISCOIGMPFILTERMIB{}))
}

// CISCOIGMPFILTERMIB
type CISCOIGMPFILTERMIB struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Cigmpfiltergeneral CISCOIGMPFILTERMIB_Cigmpfiltergeneral

    
    Cigmpfiltereditor CISCOIGMPFILTERMIB_Cigmpfiltereditor

    // This table contains entries that identify lists of IP Multicast groups for
    // each profile configured on the device.   Each entry contains a range of
    // contiguous IP Multicast groups associated to a profile index. Multiple
    // cIgmpFilterEntry may be associated with the same cIgmpFilterProfileIndex. 
    // All the cIgmpFilterEntry with  the same profile index  are subjected to the
    // same IGMP filtering action as  defined in cIgmpFilterProfileAction.  Each
    // profile index may be associated with zero or more  interfaces through
    // cIgmpFilterInterfaceIfIndex object in cIgmpFilterInterfaceTable. The
    // maximum number of entries is determined by cIgmpFilterMaxProfiles.
    Cigmpfiltertable CISCOIGMPFILTERMIB_Cigmpfiltertable

    // This table contains the list of interfaces that can support IGMP filter
    // feature.
    Cigmpfilterinterfacetable CISCOIGMPFILTERMIB_Cigmpfilterinterfacetable
}

func (cISCOIGMPFILTERMIB *CISCOIGMPFILTERMIB) GetFilter() yfilter.YFilter { return cISCOIGMPFILTERMIB.YFilter }

func (cISCOIGMPFILTERMIB *CISCOIGMPFILTERMIB) SetFilter(yf yfilter.YFilter) { cISCOIGMPFILTERMIB.YFilter = yf }

func (cISCOIGMPFILTERMIB *CISCOIGMPFILTERMIB) GetGoName(yname string) string {
    if yname == "cIgmpFilterGeneral" { return "Cigmpfiltergeneral" }
    if yname == "cIgmpFilterEditor" { return "Cigmpfiltereditor" }
    if yname == "cIgmpFilterTable" { return "Cigmpfiltertable" }
    if yname == "cIgmpFilterInterfaceTable" { return "Cigmpfilterinterfacetable" }
    return ""
}

func (cISCOIGMPFILTERMIB *CISCOIGMPFILTERMIB) GetSegmentPath() string {
    return "CISCO-IGMP-FILTER-MIB:CISCO-IGMP-FILTER-MIB"
}

func (cISCOIGMPFILTERMIB *CISCOIGMPFILTERMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cIgmpFilterGeneral" {
        return &cISCOIGMPFILTERMIB.Cigmpfiltergeneral
    }
    if childYangName == "cIgmpFilterEditor" {
        return &cISCOIGMPFILTERMIB.Cigmpfiltereditor
    }
    if childYangName == "cIgmpFilterTable" {
        return &cISCOIGMPFILTERMIB.Cigmpfiltertable
    }
    if childYangName == "cIgmpFilterInterfaceTable" {
        return &cISCOIGMPFILTERMIB.Cigmpfilterinterfacetable
    }
    return nil
}

func (cISCOIGMPFILTERMIB *CISCOIGMPFILTERMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cIgmpFilterGeneral"] = &cISCOIGMPFILTERMIB.Cigmpfiltergeneral
    children["cIgmpFilterEditor"] = &cISCOIGMPFILTERMIB.Cigmpfiltereditor
    children["cIgmpFilterTable"] = &cISCOIGMPFILTERMIB.Cigmpfiltertable
    children["cIgmpFilterInterfaceTable"] = &cISCOIGMPFILTERMIB.Cigmpfilterinterfacetable
    return children
}

func (cISCOIGMPFILTERMIB *CISCOIGMPFILTERMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOIGMPFILTERMIB *CISCOIGMPFILTERMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOIGMPFILTERMIB *CISCOIGMPFILTERMIB) GetYangName() string { return "CISCO-IGMP-FILTER-MIB" }

func (cISCOIGMPFILTERMIB *CISCOIGMPFILTERMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOIGMPFILTERMIB *CISCOIGMPFILTERMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOIGMPFILTERMIB *CISCOIGMPFILTERMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOIGMPFILTERMIB *CISCOIGMPFILTERMIB) SetParent(parent types.Entity) { cISCOIGMPFILTERMIB.parent = parent }

func (cISCOIGMPFILTERMIB *CISCOIGMPFILTERMIB) GetParent() types.Entity { return cISCOIGMPFILTERMIB.parent }

func (cISCOIGMPFILTERMIB *CISCOIGMPFILTERMIB) GetParentYangName() string { return "CISCO-IGMP-FILTER-MIB" }

// CISCOIGMPFILTERMIB_Cigmpfiltergeneral
type CISCOIGMPFILTERMIB_Cigmpfiltergeneral struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This object controls whether the IGMP filtering is enabled on the device. A
    // false(2) value will  prevent the IGMP filtering action on the device. The
    // type is bool.
    Cigmpfilterenable interface{}

    // Indicates the maximum number of profiles supported by this device.  A value
    // of zero indicates no limitation on the number of profiles. The type is
    // interface{} with range: 0..4294967295. Units are profiles.
    Cigmpfiltermaxprofiles interface{}
}

func (cigmpfiltergeneral *CISCOIGMPFILTERMIB_Cigmpfiltergeneral) GetFilter() yfilter.YFilter { return cigmpfiltergeneral.YFilter }

func (cigmpfiltergeneral *CISCOIGMPFILTERMIB_Cigmpfiltergeneral) SetFilter(yf yfilter.YFilter) { cigmpfiltergeneral.YFilter = yf }

func (cigmpfiltergeneral *CISCOIGMPFILTERMIB_Cigmpfiltergeneral) GetGoName(yname string) string {
    if yname == "cIgmpFilterEnable" { return "Cigmpfilterenable" }
    if yname == "cIgmpFilterMaxProfiles" { return "Cigmpfiltermaxprofiles" }
    return ""
}

func (cigmpfiltergeneral *CISCOIGMPFILTERMIB_Cigmpfiltergeneral) GetSegmentPath() string {
    return "cIgmpFilterGeneral"
}

func (cigmpfiltergeneral *CISCOIGMPFILTERMIB_Cigmpfiltergeneral) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cigmpfiltergeneral *CISCOIGMPFILTERMIB_Cigmpfiltergeneral) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cigmpfiltergeneral *CISCOIGMPFILTERMIB_Cigmpfiltergeneral) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cIgmpFilterEnable"] = cigmpfiltergeneral.Cigmpfilterenable
    leafs["cIgmpFilterMaxProfiles"] = cigmpfiltergeneral.Cigmpfiltermaxprofiles
    return leafs
}

func (cigmpfiltergeneral *CISCOIGMPFILTERMIB_Cigmpfiltergeneral) GetBundleName() string { return "cisco_ios_xe" }

func (cigmpfiltergeneral *CISCOIGMPFILTERMIB_Cigmpfiltergeneral) GetYangName() string { return "cIgmpFilterGeneral" }

func (cigmpfiltergeneral *CISCOIGMPFILTERMIB_Cigmpfiltergeneral) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cigmpfiltergeneral *CISCOIGMPFILTERMIB_Cigmpfiltergeneral) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cigmpfiltergeneral *CISCOIGMPFILTERMIB_Cigmpfiltergeneral) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cigmpfiltergeneral *CISCOIGMPFILTERMIB_Cigmpfiltergeneral) SetParent(parent types.Entity) { cigmpfiltergeneral.parent = parent }

func (cigmpfiltergeneral *CISCOIGMPFILTERMIB_Cigmpfiltergeneral) GetParent() types.Entity { return cigmpfiltergeneral.parent }

func (cigmpfiltergeneral *CISCOIGMPFILTERMIB_Cigmpfiltergeneral) GetParentYangName() string { return "CISCO-IGMP-FILTER-MIB" }

// CISCOIGMPFILTERMIB_Cigmpfiltereditor
type CISCOIGMPFILTERMIB_Cigmpfiltereditor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This object is used to facilitate modification of IGMP Filter Editor Group
    // in the CISCO-IGMP-FILTER-MIB  module by multiple managers.  In particular,
    // it is  useful when modifying the value of the cIgmpFilterEditor  object.  
    // The procedure for modifying the cIgmpFilterEditor  object is as follows:   
    // 1.  Retrieve the value of cIgmpFilterEditSpinLock and          of object
    // within cIgmpFilterEditor group.       2.  Generate new values for
    // 'writeable' objects         in the cIgmpFilterEditor group except for      
    // cIgmpFilterEditSpinLock object.      3.  Set the value of
    // cIgmpFilterEditSpinLock to the          retrieved value, and the value of
    // objects in          cIgmpFilterEditor to the new value. If the set        
    // fails for the cIgmpFilterEditSpinLock object,         go back to step 1.  
    // The cIgmpFilterApplyStatus and cIgmpFilterEditSpinLock  should be read
    // together by NMS to make sure that the  result is associated with its
    // configuration request.  To add/delete or modify a profile ( through
    // cIgmpFilterEditor objects ) following procedure may be followed as an
    // example.:      1.  GET(cIgmpFilterEditSpinLock.0) and save in sValue.    
    // 2.  SET(cIgmpFilterEditSpinLock.0 = sValue,            
    // cIgmpFilterEditProfileIndex.0 = validProfileNumber,            
    // cIgmpFilterEditStartAddress.0 = validStartAddress,            
    // cIgmpFilterEditEndAddress.0 = validEndAddress,            
    // cIgmpFilterEditOperation.0 =  validOperation)     3.  If the SET in step 2
    // is not successful go back         to step 1.     4.  If the SET in step 2
    // is successful, user should          GET(cIgmpFilterEditSpinLock.0) and     
    // GET(cIgmpFilterApplyStatus.0) simultaneously.      5.  The
    // cIgmpFilterApplyStatus.0 reflects the outcome of         step 2 only if    
    // cIgmpFilterEditSpinLock.0 = sValue + 1 or         cIgmpFilterEditSpinLock.0
    // = 0 if sValue reaches         value at which cIgmpFilterEditSpinLock wraps 
    // around. The type is interface{} with range: 0..2147483647.
    Cigmpfiltereditspinlock interface{}

    // Buffer object to edit corresponding object cIgmpFilterProfileIndex in
    // cIgmpFilterTable.  Maximum value this object can be set to is  determined
    // by cIgmpFilterMaxProfiles object. The type is interface{} with range:
    // 0..4294967295.
    Cigmpfiltereditprofileindex interface{}

    // Buffer object to edit corresponding object cIgmpFilterStartAddressType in
    // cIgmpFilterTable.  This object describes the type of Internet   address
    // used to determine the start address  of IP multicast group for a profile.
    // The type is InetAddressType.
    Cigmpfiltereditstartaddresstype interface{}

    // Buffer object to edit corresponding object cIgmpFilterStartAddress in
    // cIgmpFilterTable   to establish start address of filtering range for a
    // profile. The type is string with length: 1..64.
    Cigmpfiltereditstartaddress interface{}

    // Buffer object to edit corresponding object cIgmpFilterEndAddressType in
    // cIgmpFilterTable.  This object describes the type of Internet      address
    // used to determine the start address  of IP multicast group for a profile.
    // The type is InetAddressType.
    Cigmpfiltereditendaddresstype interface{}

    // Buffer object to edit corresponding object cIgmpFilterEndAddress in
    // cIgmpFilterTable  to establish end address of filtering  range for a
    // profile. The type is string with length: 0..255.
    Cigmpfiltereditendaddress interface{}

    // Buffer object to edit corresponding object in cIgmpFilterTable to determine
    // filtering action of each profile. The type is Cigmpfiltereditprofileaction.
    Cigmpfiltereditprofileaction interface{}

    // The function of this object is to allow user to apply the changes in
    // cIgmpFilterEditor objects to  cIgmpFilterTable.   This object always has
    // the value 'none' when read. When written each value causes the appropriate
    // action:  'add' - tries to insert the information contained  in
    // cIgmpFilterEditor objects into  cIgmpFilterTable. If the entry already
    // exists in the table the 'add'  fails.          'delete' - tries to delete
    // corresponding entry from cIgmpFilterTable. If a user completely deletes a
    // profile that has corresponding entries in the cIgmpFilterInterfaceTable,
    // then all the interfaces mapped to corresponding profile will be cleared and
    // set to zero.  'modify' - Mode of operation used to edit an entry in 
    // cIgmpFilterTable. If the corresponding entry does not   exist, modify
    // operation fails. This mode allows user to  extend/truncate a contiguous
    // filtered range, type of  Internet addressing and filtering action for a
    // profile.   'none' - no operation is performed. The type is
    // Cigmpfiltereditoperation.
    Cigmpfiltereditoperation interface{}

    // The current status of an 'add', 'delete' or 'modify' operation. If no apply
    // is currently active, the status  represented is that of the most recently
    // completed 'add',  'delete' or 'modify' operation.  The value of this
    // objects indicates succeeded(2) state  initially when no 'add', 'delete',
    // 'modify' operation has been carried out.  The possible values are:   
    // someOtherError - the 'add', 'delete' or 'modify' failed     due to
    // undefined error(s) in apply operation.    (e.g., insufficient memory).     
    // succeeded - the 'add', 'delete' or 'modify' operation    was successful.
    // (This value is also used when no    apply has been invoked since the last
    // time the local     system restarted.)     inconsistentEdit - the 'add',
    // 'delete' or 'modify' failed    due to inconsistency of the data.    
    // entryPresentError - the 'add' operation failed  as the     corresponding
    // entry already exists in cIgmpFilterTable.     entryNotPresentError - the
    // 'modify' operation failed     as no corresponding entry exists in
    // cIgmpFilterTable. The type is Cigmpfilterapplystatus.
    Cigmpfilterapplystatus interface{}
}

func (cigmpfiltereditor *CISCOIGMPFILTERMIB_Cigmpfiltereditor) GetFilter() yfilter.YFilter { return cigmpfiltereditor.YFilter }

func (cigmpfiltereditor *CISCOIGMPFILTERMIB_Cigmpfiltereditor) SetFilter(yf yfilter.YFilter) { cigmpfiltereditor.YFilter = yf }

func (cigmpfiltereditor *CISCOIGMPFILTERMIB_Cigmpfiltereditor) GetGoName(yname string) string {
    if yname == "cIgmpFilterEditSpinLock" { return "Cigmpfiltereditspinlock" }
    if yname == "cIgmpFilterEditProfileIndex" { return "Cigmpfiltereditprofileindex" }
    if yname == "cIgmpFilterEditStartAddressType" { return "Cigmpfiltereditstartaddresstype" }
    if yname == "cIgmpFilterEditStartAddress" { return "Cigmpfiltereditstartaddress" }
    if yname == "cIgmpFilterEditEndAddressType" { return "Cigmpfiltereditendaddresstype" }
    if yname == "cIgmpFilterEditEndAddress" { return "Cigmpfiltereditendaddress" }
    if yname == "cIgmpFilterEditProfileAction" { return "Cigmpfiltereditprofileaction" }
    if yname == "cIgmpFilterEditOperation" { return "Cigmpfiltereditoperation" }
    if yname == "cIgmpFilterApplyStatus" { return "Cigmpfilterapplystatus" }
    return ""
}

func (cigmpfiltereditor *CISCOIGMPFILTERMIB_Cigmpfiltereditor) GetSegmentPath() string {
    return "cIgmpFilterEditor"
}

func (cigmpfiltereditor *CISCOIGMPFILTERMIB_Cigmpfiltereditor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cigmpfiltereditor *CISCOIGMPFILTERMIB_Cigmpfiltereditor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cigmpfiltereditor *CISCOIGMPFILTERMIB_Cigmpfiltereditor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cIgmpFilterEditSpinLock"] = cigmpfiltereditor.Cigmpfiltereditspinlock
    leafs["cIgmpFilterEditProfileIndex"] = cigmpfiltereditor.Cigmpfiltereditprofileindex
    leafs["cIgmpFilterEditStartAddressType"] = cigmpfiltereditor.Cigmpfiltereditstartaddresstype
    leafs["cIgmpFilterEditStartAddress"] = cigmpfiltereditor.Cigmpfiltereditstartaddress
    leafs["cIgmpFilterEditEndAddressType"] = cigmpfiltereditor.Cigmpfiltereditendaddresstype
    leafs["cIgmpFilterEditEndAddress"] = cigmpfiltereditor.Cigmpfiltereditendaddress
    leafs["cIgmpFilterEditProfileAction"] = cigmpfiltereditor.Cigmpfiltereditprofileaction
    leafs["cIgmpFilterEditOperation"] = cigmpfiltereditor.Cigmpfiltereditoperation
    leafs["cIgmpFilterApplyStatus"] = cigmpfiltereditor.Cigmpfilterapplystatus
    return leafs
}

func (cigmpfiltereditor *CISCOIGMPFILTERMIB_Cigmpfiltereditor) GetBundleName() string { return "cisco_ios_xe" }

func (cigmpfiltereditor *CISCOIGMPFILTERMIB_Cigmpfiltereditor) GetYangName() string { return "cIgmpFilterEditor" }

func (cigmpfiltereditor *CISCOIGMPFILTERMIB_Cigmpfiltereditor) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cigmpfiltereditor *CISCOIGMPFILTERMIB_Cigmpfiltereditor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cigmpfiltereditor *CISCOIGMPFILTERMIB_Cigmpfiltereditor) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cigmpfiltereditor *CISCOIGMPFILTERMIB_Cigmpfiltereditor) SetParent(parent types.Entity) { cigmpfiltereditor.parent = parent }

func (cigmpfiltereditor *CISCOIGMPFILTERMIB_Cigmpfiltereditor) GetParent() types.Entity { return cigmpfiltereditor.parent }

func (cigmpfiltereditor *CISCOIGMPFILTERMIB_Cigmpfiltereditor) GetParentYangName() string { return "CISCO-IGMP-FILTER-MIB" }

// CISCOIGMPFILTERMIB_Cigmpfiltereditor_Cigmpfilterapplystatus represents    as no corresponding entry exists in cIgmpFilterTable.
type CISCOIGMPFILTERMIB_Cigmpfiltereditor_Cigmpfilterapplystatus string

const (
    CISCOIGMPFILTERMIB_Cigmpfiltereditor_Cigmpfilterapplystatus_someOtherError CISCOIGMPFILTERMIB_Cigmpfiltereditor_Cigmpfilterapplystatus = "someOtherError"

    CISCOIGMPFILTERMIB_Cigmpfiltereditor_Cigmpfilterapplystatus_succeeded CISCOIGMPFILTERMIB_Cigmpfiltereditor_Cigmpfilterapplystatus = "succeeded"

    CISCOIGMPFILTERMIB_Cigmpfiltereditor_Cigmpfilterapplystatus_inconsistentEdit CISCOIGMPFILTERMIB_Cigmpfiltereditor_Cigmpfilterapplystatus = "inconsistentEdit"

    CISCOIGMPFILTERMIB_Cigmpfiltereditor_Cigmpfilterapplystatus_entryPresentError CISCOIGMPFILTERMIB_Cigmpfiltereditor_Cigmpfilterapplystatus = "entryPresentError"

    CISCOIGMPFILTERMIB_Cigmpfiltereditor_Cigmpfilterapplystatus_entryNotPresentError CISCOIGMPFILTERMIB_Cigmpfiltereditor_Cigmpfilterapplystatus = "entryNotPresentError"
)

// CISCOIGMPFILTERMIB_Cigmpfiltereditor_Cigmpfiltereditoperation represents 'none' - no operation is performed.
type CISCOIGMPFILTERMIB_Cigmpfiltereditor_Cigmpfiltereditoperation string

const (
    CISCOIGMPFILTERMIB_Cigmpfiltereditor_Cigmpfiltereditoperation_none CISCOIGMPFILTERMIB_Cigmpfiltereditor_Cigmpfiltereditoperation = "none"

    CISCOIGMPFILTERMIB_Cigmpfiltereditor_Cigmpfiltereditoperation_add CISCOIGMPFILTERMIB_Cigmpfiltereditor_Cigmpfiltereditoperation = "add"

    CISCOIGMPFILTERMIB_Cigmpfiltereditor_Cigmpfiltereditoperation_delete CISCOIGMPFILTERMIB_Cigmpfiltereditor_Cigmpfiltereditoperation = "delete"

    CISCOIGMPFILTERMIB_Cigmpfiltereditor_Cigmpfiltereditoperation_modify CISCOIGMPFILTERMIB_Cigmpfiltereditor_Cigmpfiltereditoperation = "modify"
)

// CISCOIGMPFILTERMIB_Cigmpfiltereditor_Cigmpfiltereditprofileaction represents of each profile.
type CISCOIGMPFILTERMIB_Cigmpfiltereditor_Cigmpfiltereditprofileaction string

const (
    CISCOIGMPFILTERMIB_Cigmpfiltereditor_Cigmpfiltereditprofileaction_unSpecified CISCOIGMPFILTERMIB_Cigmpfiltereditor_Cigmpfiltereditprofileaction = "unSpecified"

    CISCOIGMPFILTERMIB_Cigmpfiltereditor_Cigmpfiltereditprofileaction_permit CISCOIGMPFILTERMIB_Cigmpfiltereditor_Cigmpfiltereditprofileaction = "permit"

    CISCOIGMPFILTERMIB_Cigmpfiltereditor_Cigmpfiltereditprofileaction_deny CISCOIGMPFILTERMIB_Cigmpfiltereditor_Cigmpfiltereditprofileaction = "deny"
)

// CISCOIGMPFILTERMIB_Cigmpfiltertable
// This table contains entries that identify lists of
// IP Multicast groups for each profile configured on
// the device. 
// 
// Each entry contains a range of contiguous IP
// Multicast groups associated to a profile index.
// Multiple cIgmpFilterEntry may be associated
// with the same cIgmpFilterProfileIndex.
// 
// All the cIgmpFilterEntry with  the same profile index 
// are subjected to the same IGMP filtering action as 
// defined in cIgmpFilterProfileAction.
// 
// Each profile index may be associated with zero or more 
// interfaces through cIgmpFilterInterfaceIfIndex
// object in cIgmpFilterInterfaceTable.
// The maximum number of entries is determined by
// cIgmpFilterMaxProfiles.
type CISCOIGMPFILTERMIB_Cigmpfiltertable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the cIgmpFilterTable.  The creation, deletion
    // or modification of an entry is controlled through the MIB objects defined
    // under cIgmpFilterEditor group. The type is slice of
    // CISCOIGMPFILTERMIB_Cigmpfiltertable_Cigmpfilterentry.
    Cigmpfilterentry []CISCOIGMPFILTERMIB_Cigmpfiltertable_Cigmpfilterentry
}

func (cigmpfiltertable *CISCOIGMPFILTERMIB_Cigmpfiltertable) GetFilter() yfilter.YFilter { return cigmpfiltertable.YFilter }

func (cigmpfiltertable *CISCOIGMPFILTERMIB_Cigmpfiltertable) SetFilter(yf yfilter.YFilter) { cigmpfiltertable.YFilter = yf }

func (cigmpfiltertable *CISCOIGMPFILTERMIB_Cigmpfiltertable) GetGoName(yname string) string {
    if yname == "cIgmpFilterEntry" { return "Cigmpfilterentry" }
    return ""
}

func (cigmpfiltertable *CISCOIGMPFILTERMIB_Cigmpfiltertable) GetSegmentPath() string {
    return "cIgmpFilterTable"
}

func (cigmpfiltertable *CISCOIGMPFILTERMIB_Cigmpfiltertable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cIgmpFilterEntry" {
        for _, c := range cigmpfiltertable.Cigmpfilterentry {
            if cigmpfiltertable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIGMPFILTERMIB_Cigmpfiltertable_Cigmpfilterentry{}
        cigmpfiltertable.Cigmpfilterentry = append(cigmpfiltertable.Cigmpfilterentry, child)
        return &cigmpfiltertable.Cigmpfilterentry[len(cigmpfiltertable.Cigmpfilterentry)-1]
    }
    return nil
}

func (cigmpfiltertable *CISCOIGMPFILTERMIB_Cigmpfiltertable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cigmpfiltertable.Cigmpfilterentry {
        children[cigmpfiltertable.Cigmpfilterentry[i].GetSegmentPath()] = &cigmpfiltertable.Cigmpfilterentry[i]
    }
    return children
}

func (cigmpfiltertable *CISCOIGMPFILTERMIB_Cigmpfiltertable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cigmpfiltertable *CISCOIGMPFILTERMIB_Cigmpfiltertable) GetBundleName() string { return "cisco_ios_xe" }

func (cigmpfiltertable *CISCOIGMPFILTERMIB_Cigmpfiltertable) GetYangName() string { return "cIgmpFilterTable" }

func (cigmpfiltertable *CISCOIGMPFILTERMIB_Cigmpfiltertable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cigmpfiltertable *CISCOIGMPFILTERMIB_Cigmpfiltertable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cigmpfiltertable *CISCOIGMPFILTERMIB_Cigmpfiltertable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cigmpfiltertable *CISCOIGMPFILTERMIB_Cigmpfiltertable) SetParent(parent types.Entity) { cigmpfiltertable.parent = parent }

func (cigmpfiltertable *CISCOIGMPFILTERMIB_Cigmpfiltertable) GetParent() types.Entity { return cigmpfiltertable.parent }

func (cigmpfiltertable *CISCOIGMPFILTERMIB_Cigmpfiltertable) GetParentYangName() string { return "CISCO-IGMP-FILTER-MIB" }

// CISCOIGMPFILTERMIB_Cigmpfiltertable_Cigmpfilterentry
// An entry (conceptual row) in the cIgmpFilterTable.
// 
// The creation, deletion or modification of an entry
// is controlled through the MIB objects defined under
// cIgmpFilterEditor group.
type CISCOIGMPFILTERMIB_Cigmpfiltertable_Cigmpfilterentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Index identifying this entry. The type is
    // interface{} with range: 0..4294967295.
    Cigmpfilterprofileindex interface{}

    // This attribute is a key. This object describes the type of Internet address
    // used to determine the start address  of IP multicast group for a profile.
    // The type is InetAddressType.
    Cigmpfilterstartaddresstype interface{}

    // This attribute is a key. This object describes the start of the IP
    // multicast group address of a contiguous range which will be subjected to
    // filtering operation. The type is string with length: 1..64.
    Cigmpfilterstartaddress interface{}

    // This object indicates the type of Internet address used to determine the
    // end address  of IP multicast group for a profile. The type is
    // InetAddressType.
    Cigmpfilterendaddresstype interface{}

    // This object indicates the end of the IP multicast group address of a
    // contiguous range which will be  subjected to filtering operation. The type
    // is string with length: 0..255.
    Cigmpfilterendaddress interface{}

    // This object defines the action for filtering IGMP reports for this profile.
    // If the object is set to deny(2): then all IGMP reports associated to IP
    // multicast groups included in the profile identified by
    // cIgmpFilterInterfaceProfileIndex will be dropped.  If the object is set to
    // permit(1): then all IGMP reports associated to IP multicast groups not
    // included in the profile identified by cIgmpFilterInterfaceProfileIndex will
    // be dropped. The type is Cigmpfilterprofileaction.
    Cigmpfilterprofileaction interface{}
}

func (cigmpfilterentry *CISCOIGMPFILTERMIB_Cigmpfiltertable_Cigmpfilterentry) GetFilter() yfilter.YFilter { return cigmpfilterentry.YFilter }

func (cigmpfilterentry *CISCOIGMPFILTERMIB_Cigmpfiltertable_Cigmpfilterentry) SetFilter(yf yfilter.YFilter) { cigmpfilterentry.YFilter = yf }

func (cigmpfilterentry *CISCOIGMPFILTERMIB_Cigmpfiltertable_Cigmpfilterentry) GetGoName(yname string) string {
    if yname == "cIgmpFilterProfileIndex" { return "Cigmpfilterprofileindex" }
    if yname == "cIgmpFilterStartAddressType" { return "Cigmpfilterstartaddresstype" }
    if yname == "cIgmpFilterStartAddress" { return "Cigmpfilterstartaddress" }
    if yname == "cIgmpFilterEndAddressType" { return "Cigmpfilterendaddresstype" }
    if yname == "cIgmpFilterEndAddress" { return "Cigmpfilterendaddress" }
    if yname == "cIgmpFilterProfileAction" { return "Cigmpfilterprofileaction" }
    return ""
}

func (cigmpfilterentry *CISCOIGMPFILTERMIB_Cigmpfiltertable_Cigmpfilterentry) GetSegmentPath() string {
    return "cIgmpFilterEntry" + "[cIgmpFilterProfileIndex='" + fmt.Sprintf("%v", cigmpfilterentry.Cigmpfilterprofileindex) + "']" + "[cIgmpFilterStartAddressType='" + fmt.Sprintf("%v", cigmpfilterentry.Cigmpfilterstartaddresstype) + "']" + "[cIgmpFilterStartAddress='" + fmt.Sprintf("%v", cigmpfilterentry.Cigmpfilterstartaddress) + "']"
}

func (cigmpfilterentry *CISCOIGMPFILTERMIB_Cigmpfiltertable_Cigmpfilterentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cigmpfilterentry *CISCOIGMPFILTERMIB_Cigmpfiltertable_Cigmpfilterentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cigmpfilterentry *CISCOIGMPFILTERMIB_Cigmpfiltertable_Cigmpfilterentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cIgmpFilterProfileIndex"] = cigmpfilterentry.Cigmpfilterprofileindex
    leafs["cIgmpFilterStartAddressType"] = cigmpfilterentry.Cigmpfilterstartaddresstype
    leafs["cIgmpFilterStartAddress"] = cigmpfilterentry.Cigmpfilterstartaddress
    leafs["cIgmpFilterEndAddressType"] = cigmpfilterentry.Cigmpfilterendaddresstype
    leafs["cIgmpFilterEndAddress"] = cigmpfilterentry.Cigmpfilterendaddress
    leafs["cIgmpFilterProfileAction"] = cigmpfilterentry.Cigmpfilterprofileaction
    return leafs
}

func (cigmpfilterentry *CISCOIGMPFILTERMIB_Cigmpfiltertable_Cigmpfilterentry) GetBundleName() string { return "cisco_ios_xe" }

func (cigmpfilterentry *CISCOIGMPFILTERMIB_Cigmpfiltertable_Cigmpfilterentry) GetYangName() string { return "cIgmpFilterEntry" }

func (cigmpfilterentry *CISCOIGMPFILTERMIB_Cigmpfiltertable_Cigmpfilterentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cigmpfilterentry *CISCOIGMPFILTERMIB_Cigmpfiltertable_Cigmpfilterentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cigmpfilterentry *CISCOIGMPFILTERMIB_Cigmpfiltertable_Cigmpfilterentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cigmpfilterentry *CISCOIGMPFILTERMIB_Cigmpfiltertable_Cigmpfilterentry) SetParent(parent types.Entity) { cigmpfilterentry.parent = parent }

func (cigmpfilterentry *CISCOIGMPFILTERMIB_Cigmpfiltertable_Cigmpfilterentry) GetParent() types.Entity { return cigmpfilterentry.parent }

func (cigmpfilterentry *CISCOIGMPFILTERMIB_Cigmpfiltertable_Cigmpfilterentry) GetParentYangName() string { return "cIgmpFilterTable" }

// CISCOIGMPFILTERMIB_Cigmpfiltertable_Cigmpfilterentry_Cigmpfilterprofileaction represents cIgmpFilterInterfaceProfileIndex will be dropped.
type CISCOIGMPFILTERMIB_Cigmpfiltertable_Cigmpfilterentry_Cigmpfilterprofileaction string

const (
    CISCOIGMPFILTERMIB_Cigmpfiltertable_Cigmpfilterentry_Cigmpfilterprofileaction_permit CISCOIGMPFILTERMIB_Cigmpfiltertable_Cigmpfilterentry_Cigmpfilterprofileaction = "permit"

    CISCOIGMPFILTERMIB_Cigmpfiltertable_Cigmpfilterentry_Cigmpfilterprofileaction_deny CISCOIGMPFILTERMIB_Cigmpfiltertable_Cigmpfilterentry_Cigmpfilterprofileaction = "deny"
)

// CISCOIGMPFILTERMIB_Cigmpfilterinterfacetable
// This table contains the list of interfaces that can
// support IGMP filter feature.
type CISCOIGMPFILTERMIB_Cigmpfilterinterfacetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry contains the configuration for associating the IGMP filter
    // profile index with the interface.  An entry is created for each of the IGMP
    // filter capable  interface on the system.  The entry is removed on removal
    // of corresponding  interface from system. The type is slice of
    // CISCOIGMPFILTERMIB_Cigmpfilterinterfacetable_Cigmpfilterinterfaceentry.
    Cigmpfilterinterfaceentry []CISCOIGMPFILTERMIB_Cigmpfilterinterfacetable_Cigmpfilterinterfaceentry
}

func (cigmpfilterinterfacetable *CISCOIGMPFILTERMIB_Cigmpfilterinterfacetable) GetFilter() yfilter.YFilter { return cigmpfilterinterfacetable.YFilter }

func (cigmpfilterinterfacetable *CISCOIGMPFILTERMIB_Cigmpfilterinterfacetable) SetFilter(yf yfilter.YFilter) { cigmpfilterinterfacetable.YFilter = yf }

func (cigmpfilterinterfacetable *CISCOIGMPFILTERMIB_Cigmpfilterinterfacetable) GetGoName(yname string) string {
    if yname == "cIgmpFilterInterfaceEntry" { return "Cigmpfilterinterfaceentry" }
    return ""
}

func (cigmpfilterinterfacetable *CISCOIGMPFILTERMIB_Cigmpfilterinterfacetable) GetSegmentPath() string {
    return "cIgmpFilterInterfaceTable"
}

func (cigmpfilterinterfacetable *CISCOIGMPFILTERMIB_Cigmpfilterinterfacetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cIgmpFilterInterfaceEntry" {
        for _, c := range cigmpfilterinterfacetable.Cigmpfilterinterfaceentry {
            if cigmpfilterinterfacetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIGMPFILTERMIB_Cigmpfilterinterfacetable_Cigmpfilterinterfaceentry{}
        cigmpfilterinterfacetable.Cigmpfilterinterfaceentry = append(cigmpfilterinterfacetable.Cigmpfilterinterfaceentry, child)
        return &cigmpfilterinterfacetable.Cigmpfilterinterfaceentry[len(cigmpfilterinterfacetable.Cigmpfilterinterfaceentry)-1]
    }
    return nil
}

func (cigmpfilterinterfacetable *CISCOIGMPFILTERMIB_Cigmpfilterinterfacetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cigmpfilterinterfacetable.Cigmpfilterinterfaceentry {
        children[cigmpfilterinterfacetable.Cigmpfilterinterfaceentry[i].GetSegmentPath()] = &cigmpfilterinterfacetable.Cigmpfilterinterfaceentry[i]
    }
    return children
}

func (cigmpfilterinterfacetable *CISCOIGMPFILTERMIB_Cigmpfilterinterfacetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cigmpfilterinterfacetable *CISCOIGMPFILTERMIB_Cigmpfilterinterfacetable) GetBundleName() string { return "cisco_ios_xe" }

func (cigmpfilterinterfacetable *CISCOIGMPFILTERMIB_Cigmpfilterinterfacetable) GetYangName() string { return "cIgmpFilterInterfaceTable" }

func (cigmpfilterinterfacetable *CISCOIGMPFILTERMIB_Cigmpfilterinterfacetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cigmpfilterinterfacetable *CISCOIGMPFILTERMIB_Cigmpfilterinterfacetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cigmpfilterinterfacetable *CISCOIGMPFILTERMIB_Cigmpfilterinterfacetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cigmpfilterinterfacetable *CISCOIGMPFILTERMIB_Cigmpfilterinterfacetable) SetParent(parent types.Entity) { cigmpfilterinterfacetable.parent = parent }

func (cigmpfilterinterfacetable *CISCOIGMPFILTERMIB_Cigmpfilterinterfacetable) GetParent() types.Entity { return cigmpfilterinterfacetable.parent }

func (cigmpfilterinterfacetable *CISCOIGMPFILTERMIB_Cigmpfilterinterfacetable) GetParentYangName() string { return "CISCO-IGMP-FILTER-MIB" }

// CISCOIGMPFILTERMIB_Cigmpfilterinterfacetable_Cigmpfilterinterfaceentry
// Each entry contains the configuration for associating
// the IGMP filter profile index with the interface.
// 
// An entry is created for each of the IGMP filter capable 
// interface on the system.
// 
// The entry is removed on removal of corresponding 
// interface from system.
type CISCOIGMPFILTERMIB_Cigmpfilterinterfacetable_Cigmpfilterinterfaceentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // Specifies which IGMP filter profile applies to this interface. If the value
    // of this MIB object matches the  value of cIgmpFilterProfileIndex in
    // cIgmpFilterTable,  the corresponding profile configuration will apply to
    // this interface.   Agent returns inconsistentValue if this profile  does not
    // exist in cIgmpFilterTable.  A value of zero indicates no profile is
    // associated with corresponding interface.  The filtering action on each
    // interface is also defined by the associated profile. The type is
    // interface{} with range: 0..4294967295.
    Cigmpfilterinterfaceprofileindex interface{}
}

func (cigmpfilterinterfaceentry *CISCOIGMPFILTERMIB_Cigmpfilterinterfacetable_Cigmpfilterinterfaceentry) GetFilter() yfilter.YFilter { return cigmpfilterinterfaceentry.YFilter }

func (cigmpfilterinterfaceentry *CISCOIGMPFILTERMIB_Cigmpfilterinterfacetable_Cigmpfilterinterfaceentry) SetFilter(yf yfilter.YFilter) { cigmpfilterinterfaceentry.YFilter = yf }

func (cigmpfilterinterfaceentry *CISCOIGMPFILTERMIB_Cigmpfilterinterfacetable_Cigmpfilterinterfaceentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "cIgmpFilterInterfaceProfileIndex" { return "Cigmpfilterinterfaceprofileindex" }
    return ""
}

func (cigmpfilterinterfaceentry *CISCOIGMPFILTERMIB_Cigmpfilterinterfacetable_Cigmpfilterinterfaceentry) GetSegmentPath() string {
    return "cIgmpFilterInterfaceEntry" + "[ifIndex='" + fmt.Sprintf("%v", cigmpfilterinterfaceentry.Ifindex) + "']"
}

func (cigmpfilterinterfaceentry *CISCOIGMPFILTERMIB_Cigmpfilterinterfacetable_Cigmpfilterinterfaceentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cigmpfilterinterfaceentry *CISCOIGMPFILTERMIB_Cigmpfilterinterfacetable_Cigmpfilterinterfaceentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cigmpfilterinterfaceentry *CISCOIGMPFILTERMIB_Cigmpfilterinterfacetable_Cigmpfilterinterfaceentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = cigmpfilterinterfaceentry.Ifindex
    leafs["cIgmpFilterInterfaceProfileIndex"] = cigmpfilterinterfaceentry.Cigmpfilterinterfaceprofileindex
    return leafs
}

func (cigmpfilterinterfaceentry *CISCOIGMPFILTERMIB_Cigmpfilterinterfacetable_Cigmpfilterinterfaceentry) GetBundleName() string { return "cisco_ios_xe" }

func (cigmpfilterinterfaceentry *CISCOIGMPFILTERMIB_Cigmpfilterinterfacetable_Cigmpfilterinterfaceentry) GetYangName() string { return "cIgmpFilterInterfaceEntry" }

func (cigmpfilterinterfaceentry *CISCOIGMPFILTERMIB_Cigmpfilterinterfacetable_Cigmpfilterinterfaceentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cigmpfilterinterfaceentry *CISCOIGMPFILTERMIB_Cigmpfilterinterfacetable_Cigmpfilterinterfaceentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cigmpfilterinterfaceentry *CISCOIGMPFILTERMIB_Cigmpfilterinterfacetable_Cigmpfilterinterfaceentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cigmpfilterinterfaceentry *CISCOIGMPFILTERMIB_Cigmpfilterinterfacetable_Cigmpfilterinterfaceentry) SetParent(parent types.Entity) { cigmpfilterinterfaceentry.parent = parent }

func (cigmpfilterinterfaceentry *CISCOIGMPFILTERMIB_Cigmpfilterinterfacetable_Cigmpfilterinterfaceentry) GetParent() types.Entity { return cigmpfilterinterfaceentry.parent }

func (cigmpfilterinterfaceentry *CISCOIGMPFILTERMIB_Cigmpfilterinterfacetable_Cigmpfilterinterfaceentry) GetParentYangName() string { return "cIgmpFilterInterfaceTable" }

