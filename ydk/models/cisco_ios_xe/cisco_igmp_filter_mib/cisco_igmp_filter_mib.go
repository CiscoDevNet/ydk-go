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
    EntityData types.CommonEntityData
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

func (cISCOIGMPFILTERMIB *CISCOIGMPFILTERMIB) GetEntityData() *types.CommonEntityData {
    cISCOIGMPFILTERMIB.EntityData.YFilter = cISCOIGMPFILTERMIB.YFilter
    cISCOIGMPFILTERMIB.EntityData.YangName = "CISCO-IGMP-FILTER-MIB"
    cISCOIGMPFILTERMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOIGMPFILTERMIB.EntityData.ParentYangName = "CISCO-IGMP-FILTER-MIB"
    cISCOIGMPFILTERMIB.EntityData.SegmentPath = "CISCO-IGMP-FILTER-MIB:CISCO-IGMP-FILTER-MIB"
    cISCOIGMPFILTERMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOIGMPFILTERMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOIGMPFILTERMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOIGMPFILTERMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOIGMPFILTERMIB.EntityData.Children["cIgmpFilterGeneral"] = types.YChild{"Cigmpfiltergeneral", &cISCOIGMPFILTERMIB.Cigmpfiltergeneral}
    cISCOIGMPFILTERMIB.EntityData.Children["cIgmpFilterEditor"] = types.YChild{"Cigmpfiltereditor", &cISCOIGMPFILTERMIB.Cigmpfiltereditor}
    cISCOIGMPFILTERMIB.EntityData.Children["cIgmpFilterTable"] = types.YChild{"Cigmpfiltertable", &cISCOIGMPFILTERMIB.Cigmpfiltertable}
    cISCOIGMPFILTERMIB.EntityData.Children["cIgmpFilterInterfaceTable"] = types.YChild{"Cigmpfilterinterfacetable", &cISCOIGMPFILTERMIB.Cigmpfilterinterfacetable}
    cISCOIGMPFILTERMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOIGMPFILTERMIB.EntityData)
}

// CISCOIGMPFILTERMIB_Cigmpfiltergeneral
type CISCOIGMPFILTERMIB_Cigmpfiltergeneral struct {
    EntityData types.CommonEntityData
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

func (cigmpfiltergeneral *CISCOIGMPFILTERMIB_Cigmpfiltergeneral) GetEntityData() *types.CommonEntityData {
    cigmpfiltergeneral.EntityData.YFilter = cigmpfiltergeneral.YFilter
    cigmpfiltergeneral.EntityData.YangName = "cIgmpFilterGeneral"
    cigmpfiltergeneral.EntityData.BundleName = "cisco_ios_xe"
    cigmpfiltergeneral.EntityData.ParentYangName = "CISCO-IGMP-FILTER-MIB"
    cigmpfiltergeneral.EntityData.SegmentPath = "cIgmpFilterGeneral"
    cigmpfiltergeneral.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cigmpfiltergeneral.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cigmpfiltergeneral.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cigmpfiltergeneral.EntityData.Children = make(map[string]types.YChild)
    cigmpfiltergeneral.EntityData.Leafs = make(map[string]types.YLeaf)
    cigmpfiltergeneral.EntityData.Leafs["cIgmpFilterEnable"] = types.YLeaf{"Cigmpfilterenable", cigmpfiltergeneral.Cigmpfilterenable}
    cigmpfiltergeneral.EntityData.Leafs["cIgmpFilterMaxProfiles"] = types.YLeaf{"Cigmpfiltermaxprofiles", cigmpfiltergeneral.Cigmpfiltermaxprofiles}
    return &(cigmpfiltergeneral.EntityData)
}

// CISCOIGMPFILTERMIB_Cigmpfiltereditor
type CISCOIGMPFILTERMIB_Cigmpfiltereditor struct {
    EntityData types.CommonEntityData
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

func (cigmpfiltereditor *CISCOIGMPFILTERMIB_Cigmpfiltereditor) GetEntityData() *types.CommonEntityData {
    cigmpfiltereditor.EntityData.YFilter = cigmpfiltereditor.YFilter
    cigmpfiltereditor.EntityData.YangName = "cIgmpFilterEditor"
    cigmpfiltereditor.EntityData.BundleName = "cisco_ios_xe"
    cigmpfiltereditor.EntityData.ParentYangName = "CISCO-IGMP-FILTER-MIB"
    cigmpfiltereditor.EntityData.SegmentPath = "cIgmpFilterEditor"
    cigmpfiltereditor.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cigmpfiltereditor.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cigmpfiltereditor.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cigmpfiltereditor.EntityData.Children = make(map[string]types.YChild)
    cigmpfiltereditor.EntityData.Leafs = make(map[string]types.YLeaf)
    cigmpfiltereditor.EntityData.Leafs["cIgmpFilterEditSpinLock"] = types.YLeaf{"Cigmpfiltereditspinlock", cigmpfiltereditor.Cigmpfiltereditspinlock}
    cigmpfiltereditor.EntityData.Leafs["cIgmpFilterEditProfileIndex"] = types.YLeaf{"Cigmpfiltereditprofileindex", cigmpfiltereditor.Cigmpfiltereditprofileindex}
    cigmpfiltereditor.EntityData.Leafs["cIgmpFilterEditStartAddressType"] = types.YLeaf{"Cigmpfiltereditstartaddresstype", cigmpfiltereditor.Cigmpfiltereditstartaddresstype}
    cigmpfiltereditor.EntityData.Leafs["cIgmpFilterEditStartAddress"] = types.YLeaf{"Cigmpfiltereditstartaddress", cigmpfiltereditor.Cigmpfiltereditstartaddress}
    cigmpfiltereditor.EntityData.Leafs["cIgmpFilterEditEndAddressType"] = types.YLeaf{"Cigmpfiltereditendaddresstype", cigmpfiltereditor.Cigmpfiltereditendaddresstype}
    cigmpfiltereditor.EntityData.Leafs["cIgmpFilterEditEndAddress"] = types.YLeaf{"Cigmpfiltereditendaddress", cigmpfiltereditor.Cigmpfiltereditendaddress}
    cigmpfiltereditor.EntityData.Leafs["cIgmpFilterEditProfileAction"] = types.YLeaf{"Cigmpfiltereditprofileaction", cigmpfiltereditor.Cigmpfiltereditprofileaction}
    cigmpfiltereditor.EntityData.Leafs["cIgmpFilterEditOperation"] = types.YLeaf{"Cigmpfiltereditoperation", cigmpfiltereditor.Cigmpfiltereditoperation}
    cigmpfiltereditor.EntityData.Leafs["cIgmpFilterApplyStatus"] = types.YLeaf{"Cigmpfilterapplystatus", cigmpfiltereditor.Cigmpfilterapplystatus}
    return &(cigmpfiltereditor.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the cIgmpFilterTable.  The creation, deletion
    // or modification of an entry is controlled through the MIB objects defined
    // under cIgmpFilterEditor group. The type is slice of
    // CISCOIGMPFILTERMIB_Cigmpfiltertable_Cigmpfilterentry.
    Cigmpfilterentry []CISCOIGMPFILTERMIB_Cigmpfiltertable_Cigmpfilterentry
}

func (cigmpfiltertable *CISCOIGMPFILTERMIB_Cigmpfiltertable) GetEntityData() *types.CommonEntityData {
    cigmpfiltertable.EntityData.YFilter = cigmpfiltertable.YFilter
    cigmpfiltertable.EntityData.YangName = "cIgmpFilterTable"
    cigmpfiltertable.EntityData.BundleName = "cisco_ios_xe"
    cigmpfiltertable.EntityData.ParentYangName = "CISCO-IGMP-FILTER-MIB"
    cigmpfiltertable.EntityData.SegmentPath = "cIgmpFilterTable"
    cigmpfiltertable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cigmpfiltertable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cigmpfiltertable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cigmpfiltertable.EntityData.Children = make(map[string]types.YChild)
    cigmpfiltertable.EntityData.Children["cIgmpFilterEntry"] = types.YChild{"Cigmpfilterentry", nil}
    for i := range cigmpfiltertable.Cigmpfilterentry {
        cigmpfiltertable.EntityData.Children[types.GetSegmentPath(&cigmpfiltertable.Cigmpfilterentry[i])] = types.YChild{"Cigmpfilterentry", &cigmpfiltertable.Cigmpfilterentry[i]}
    }
    cigmpfiltertable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cigmpfiltertable.EntityData)
}

// CISCOIGMPFILTERMIB_Cigmpfiltertable_Cigmpfilterentry
// An entry (conceptual row) in the cIgmpFilterTable.
// 
// The creation, deletion or modification of an entry
// is controlled through the MIB objects defined under
// cIgmpFilterEditor group.
type CISCOIGMPFILTERMIB_Cigmpfiltertable_Cigmpfilterentry struct {
    EntityData types.CommonEntityData
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

func (cigmpfilterentry *CISCOIGMPFILTERMIB_Cigmpfiltertable_Cigmpfilterentry) GetEntityData() *types.CommonEntityData {
    cigmpfilterentry.EntityData.YFilter = cigmpfilterentry.YFilter
    cigmpfilterentry.EntityData.YangName = "cIgmpFilterEntry"
    cigmpfilterentry.EntityData.BundleName = "cisco_ios_xe"
    cigmpfilterentry.EntityData.ParentYangName = "cIgmpFilterTable"
    cigmpfilterentry.EntityData.SegmentPath = "cIgmpFilterEntry" + "[cIgmpFilterProfileIndex='" + fmt.Sprintf("%v", cigmpfilterentry.Cigmpfilterprofileindex) + "']" + "[cIgmpFilterStartAddressType='" + fmt.Sprintf("%v", cigmpfilterentry.Cigmpfilterstartaddresstype) + "']" + "[cIgmpFilterStartAddress='" + fmt.Sprintf("%v", cigmpfilterentry.Cigmpfilterstartaddress) + "']"
    cigmpfilterentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cigmpfilterentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cigmpfilterentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cigmpfilterentry.EntityData.Children = make(map[string]types.YChild)
    cigmpfilterentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cigmpfilterentry.EntityData.Leafs["cIgmpFilterProfileIndex"] = types.YLeaf{"Cigmpfilterprofileindex", cigmpfilterentry.Cigmpfilterprofileindex}
    cigmpfilterentry.EntityData.Leafs["cIgmpFilterStartAddressType"] = types.YLeaf{"Cigmpfilterstartaddresstype", cigmpfilterentry.Cigmpfilterstartaddresstype}
    cigmpfilterentry.EntityData.Leafs["cIgmpFilterStartAddress"] = types.YLeaf{"Cigmpfilterstartaddress", cigmpfilterentry.Cigmpfilterstartaddress}
    cigmpfilterentry.EntityData.Leafs["cIgmpFilterEndAddressType"] = types.YLeaf{"Cigmpfilterendaddresstype", cigmpfilterentry.Cigmpfilterendaddresstype}
    cigmpfilterentry.EntityData.Leafs["cIgmpFilterEndAddress"] = types.YLeaf{"Cigmpfilterendaddress", cigmpfilterentry.Cigmpfilterendaddress}
    cigmpfilterentry.EntityData.Leafs["cIgmpFilterProfileAction"] = types.YLeaf{"Cigmpfilterprofileaction", cigmpfilterentry.Cigmpfilterprofileaction}
    return &(cigmpfilterentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry contains the configuration for associating the IGMP filter
    // profile index with the interface.  An entry is created for each of the IGMP
    // filter capable  interface on the system.  The entry is removed on removal
    // of corresponding  interface from system. The type is slice of
    // CISCOIGMPFILTERMIB_Cigmpfilterinterfacetable_Cigmpfilterinterfaceentry.
    Cigmpfilterinterfaceentry []CISCOIGMPFILTERMIB_Cigmpfilterinterfacetable_Cigmpfilterinterfaceentry
}

func (cigmpfilterinterfacetable *CISCOIGMPFILTERMIB_Cigmpfilterinterfacetable) GetEntityData() *types.CommonEntityData {
    cigmpfilterinterfacetable.EntityData.YFilter = cigmpfilterinterfacetable.YFilter
    cigmpfilterinterfacetable.EntityData.YangName = "cIgmpFilterInterfaceTable"
    cigmpfilterinterfacetable.EntityData.BundleName = "cisco_ios_xe"
    cigmpfilterinterfacetable.EntityData.ParentYangName = "CISCO-IGMP-FILTER-MIB"
    cigmpfilterinterfacetable.EntityData.SegmentPath = "cIgmpFilterInterfaceTable"
    cigmpfilterinterfacetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cigmpfilterinterfacetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cigmpfilterinterfacetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cigmpfilterinterfacetable.EntityData.Children = make(map[string]types.YChild)
    cigmpfilterinterfacetable.EntityData.Children["cIgmpFilterInterfaceEntry"] = types.YChild{"Cigmpfilterinterfaceentry", nil}
    for i := range cigmpfilterinterfacetable.Cigmpfilterinterfaceentry {
        cigmpfilterinterfacetable.EntityData.Children[types.GetSegmentPath(&cigmpfilterinterfacetable.Cigmpfilterinterfaceentry[i])] = types.YChild{"Cigmpfilterinterfaceentry", &cigmpfilterinterfacetable.Cigmpfilterinterfaceentry[i]}
    }
    cigmpfilterinterfacetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cigmpfilterinterfacetable.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (cigmpfilterinterfaceentry *CISCOIGMPFILTERMIB_Cigmpfilterinterfacetable_Cigmpfilterinterfaceentry) GetEntityData() *types.CommonEntityData {
    cigmpfilterinterfaceentry.EntityData.YFilter = cigmpfilterinterfaceentry.YFilter
    cigmpfilterinterfaceentry.EntityData.YangName = "cIgmpFilterInterfaceEntry"
    cigmpfilterinterfaceentry.EntityData.BundleName = "cisco_ios_xe"
    cigmpfilterinterfaceentry.EntityData.ParentYangName = "cIgmpFilterInterfaceTable"
    cigmpfilterinterfaceentry.EntityData.SegmentPath = "cIgmpFilterInterfaceEntry" + "[ifIndex='" + fmt.Sprintf("%v", cigmpfilterinterfaceentry.Ifindex) + "']"
    cigmpfilterinterfaceentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cigmpfilterinterfaceentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cigmpfilterinterfaceentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cigmpfilterinterfaceentry.EntityData.Children = make(map[string]types.YChild)
    cigmpfilterinterfaceentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cigmpfilterinterfaceentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", cigmpfilterinterfaceentry.Ifindex}
    cigmpfilterinterfaceentry.EntityData.Leafs["cIgmpFilterInterfaceProfileIndex"] = types.YLeaf{"Cigmpfilterinterfaceprofileindex", cigmpfilterinterfaceentry.Cigmpfilterinterfaceprofileindex}
    return &(cigmpfilterinterfaceentry.EntityData)
}

