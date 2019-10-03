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

    
    CIgmpFilterGeneral CISCOIGMPFILTERMIB_CIgmpFilterGeneral

    
    CIgmpFilterEditor CISCOIGMPFILTERMIB_CIgmpFilterEditor

    // This table contains entries that identify lists of IP Multicast groups for
    // each profile configured on the device.   Each entry contains a range of
    // contiguous IP Multicast groups associated to a profile index. Multiple
    // cIgmpFilterEntry may be associated with the same cIgmpFilterProfileIndex. 
    // All the cIgmpFilterEntry with  the same profile index  are subjected to the
    // same IGMP filtering action as  defined in cIgmpFilterProfileAction.  Each
    // profile index may be associated with zero or more  interfaces through
    // cIgmpFilterInterfaceIfIndex object in cIgmpFilterInterfaceTable. The
    // maximum number of entries is determined by cIgmpFilterMaxProfiles.
    CIgmpFilterTable CISCOIGMPFILTERMIB_CIgmpFilterTable

    // This table contains the list of interfaces that can support IGMP filter
    // feature.
    CIgmpFilterInterfaceTable CISCOIGMPFILTERMIB_CIgmpFilterInterfaceTable
}

func (cISCOIGMPFILTERMIB *CISCOIGMPFILTERMIB) GetEntityData() *types.CommonEntityData {
    cISCOIGMPFILTERMIB.EntityData.YFilter = cISCOIGMPFILTERMIB.YFilter
    cISCOIGMPFILTERMIB.EntityData.YangName = "CISCO-IGMP-FILTER-MIB"
    cISCOIGMPFILTERMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOIGMPFILTERMIB.EntityData.ParentYangName = "CISCO-IGMP-FILTER-MIB"
    cISCOIGMPFILTERMIB.EntityData.SegmentPath = "CISCO-IGMP-FILTER-MIB:CISCO-IGMP-FILTER-MIB"
    cISCOIGMPFILTERMIB.EntityData.AbsolutePath = cISCOIGMPFILTERMIB.EntityData.SegmentPath
    cISCOIGMPFILTERMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOIGMPFILTERMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOIGMPFILTERMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOIGMPFILTERMIB.EntityData.Children = types.NewOrderedMap()
    cISCOIGMPFILTERMIB.EntityData.Children.Append("cIgmpFilterGeneral", types.YChild{"CIgmpFilterGeneral", &cISCOIGMPFILTERMIB.CIgmpFilterGeneral})
    cISCOIGMPFILTERMIB.EntityData.Children.Append("cIgmpFilterEditor", types.YChild{"CIgmpFilterEditor", &cISCOIGMPFILTERMIB.CIgmpFilterEditor})
    cISCOIGMPFILTERMIB.EntityData.Children.Append("cIgmpFilterTable", types.YChild{"CIgmpFilterTable", &cISCOIGMPFILTERMIB.CIgmpFilterTable})
    cISCOIGMPFILTERMIB.EntityData.Children.Append("cIgmpFilterInterfaceTable", types.YChild{"CIgmpFilterInterfaceTable", &cISCOIGMPFILTERMIB.CIgmpFilterInterfaceTable})
    cISCOIGMPFILTERMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOIGMPFILTERMIB.EntityData.YListKeys = []string {}

    return &(cISCOIGMPFILTERMIB.EntityData)
}

// CISCOIGMPFILTERMIB_CIgmpFilterGeneral
type CISCOIGMPFILTERMIB_CIgmpFilterGeneral struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object controls whether the IGMP filtering is enabled on the device. A
    // false(2) value will  prevent the IGMP filtering action on the device. The
    // type is bool.
    CIgmpFilterEnable interface{}

    // Indicates the maximum number of profiles supported by this device.  A value
    // of zero indicates no limitation on the number of profiles. The type is
    // interface{} with range: 0..4294967295. Units are profiles.
    CIgmpFilterMaxProfiles interface{}
}

func (cIgmpFilterGeneral *CISCOIGMPFILTERMIB_CIgmpFilterGeneral) GetEntityData() *types.CommonEntityData {
    cIgmpFilterGeneral.EntityData.YFilter = cIgmpFilterGeneral.YFilter
    cIgmpFilterGeneral.EntityData.YangName = "cIgmpFilterGeneral"
    cIgmpFilterGeneral.EntityData.BundleName = "cisco_ios_xe"
    cIgmpFilterGeneral.EntityData.ParentYangName = "CISCO-IGMP-FILTER-MIB"
    cIgmpFilterGeneral.EntityData.SegmentPath = "cIgmpFilterGeneral"
    cIgmpFilterGeneral.EntityData.AbsolutePath = "CISCO-IGMP-FILTER-MIB:CISCO-IGMP-FILTER-MIB/" + cIgmpFilterGeneral.EntityData.SegmentPath
    cIgmpFilterGeneral.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cIgmpFilterGeneral.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cIgmpFilterGeneral.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cIgmpFilterGeneral.EntityData.Children = types.NewOrderedMap()
    cIgmpFilterGeneral.EntityData.Leafs = types.NewOrderedMap()
    cIgmpFilterGeneral.EntityData.Leafs.Append("cIgmpFilterEnable", types.YLeaf{"CIgmpFilterEnable", cIgmpFilterGeneral.CIgmpFilterEnable})
    cIgmpFilterGeneral.EntityData.Leafs.Append("cIgmpFilterMaxProfiles", types.YLeaf{"CIgmpFilterMaxProfiles", cIgmpFilterGeneral.CIgmpFilterMaxProfiles})

    cIgmpFilterGeneral.EntityData.YListKeys = []string {}

    return &(cIgmpFilterGeneral.EntityData)
}

// CISCOIGMPFILTERMIB_CIgmpFilterEditor
type CISCOIGMPFILTERMIB_CIgmpFilterEditor struct {
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
    CIgmpFilterEditSpinLock interface{}

    // Buffer object to edit corresponding object cIgmpFilterProfileIndex in
    // cIgmpFilterTable.  Maximum value this object can be set to is  determined
    // by cIgmpFilterMaxProfiles object. The type is interface{} with range:
    // 0..4294967295.
    CIgmpFilterEditProfileIndex interface{}

    // Buffer object to edit corresponding object cIgmpFilterStartAddressType in
    // cIgmpFilterTable.  This object describes the type of Internet   address
    // used to determine the start address  of IP multicast group for a profile.
    // The type is InetAddressType.
    CIgmpFilterEditStartAddressType interface{}

    // Buffer object to edit corresponding object cIgmpFilterStartAddress in
    // cIgmpFilterTable   to establish start address of filtering range for a
    // profile. The type is string with length: 1..64.
    CIgmpFilterEditStartAddress interface{}

    // Buffer object to edit corresponding object cIgmpFilterEndAddressType in
    // cIgmpFilterTable.  This object describes the type of Internet      address
    // used to determine the start address  of IP multicast group for a profile.
    // The type is InetAddressType.
    CIgmpFilterEditEndAddressType interface{}

    // Buffer object to edit corresponding object cIgmpFilterEndAddress in
    // cIgmpFilterTable  to establish end address of filtering  range for a
    // profile. The type is string with length: 0..255.
    CIgmpFilterEditEndAddress interface{}

    // Buffer object to edit corresponding object in cIgmpFilterTable to determine
    // filtering action of each profile. The type is CIgmpFilterEditProfileAction.
    CIgmpFilterEditProfileAction interface{}

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
    // CIgmpFilterEditOperation.
    CIgmpFilterEditOperation interface{}

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
    // cIgmpFilterTable. The type is CIgmpFilterApplyStatus.
    CIgmpFilterApplyStatus interface{}
}

func (cIgmpFilterEditor *CISCOIGMPFILTERMIB_CIgmpFilterEditor) GetEntityData() *types.CommonEntityData {
    cIgmpFilterEditor.EntityData.YFilter = cIgmpFilterEditor.YFilter
    cIgmpFilterEditor.EntityData.YangName = "cIgmpFilterEditor"
    cIgmpFilterEditor.EntityData.BundleName = "cisco_ios_xe"
    cIgmpFilterEditor.EntityData.ParentYangName = "CISCO-IGMP-FILTER-MIB"
    cIgmpFilterEditor.EntityData.SegmentPath = "cIgmpFilterEditor"
    cIgmpFilterEditor.EntityData.AbsolutePath = "CISCO-IGMP-FILTER-MIB:CISCO-IGMP-FILTER-MIB/" + cIgmpFilterEditor.EntityData.SegmentPath
    cIgmpFilterEditor.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cIgmpFilterEditor.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cIgmpFilterEditor.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cIgmpFilterEditor.EntityData.Children = types.NewOrderedMap()
    cIgmpFilterEditor.EntityData.Leafs = types.NewOrderedMap()
    cIgmpFilterEditor.EntityData.Leafs.Append("cIgmpFilterEditSpinLock", types.YLeaf{"CIgmpFilterEditSpinLock", cIgmpFilterEditor.CIgmpFilterEditSpinLock})
    cIgmpFilterEditor.EntityData.Leafs.Append("cIgmpFilterEditProfileIndex", types.YLeaf{"CIgmpFilterEditProfileIndex", cIgmpFilterEditor.CIgmpFilterEditProfileIndex})
    cIgmpFilterEditor.EntityData.Leafs.Append("cIgmpFilterEditStartAddressType", types.YLeaf{"CIgmpFilterEditStartAddressType", cIgmpFilterEditor.CIgmpFilterEditStartAddressType})
    cIgmpFilterEditor.EntityData.Leafs.Append("cIgmpFilterEditStartAddress", types.YLeaf{"CIgmpFilterEditStartAddress", cIgmpFilterEditor.CIgmpFilterEditStartAddress})
    cIgmpFilterEditor.EntityData.Leafs.Append("cIgmpFilterEditEndAddressType", types.YLeaf{"CIgmpFilterEditEndAddressType", cIgmpFilterEditor.CIgmpFilterEditEndAddressType})
    cIgmpFilterEditor.EntityData.Leafs.Append("cIgmpFilterEditEndAddress", types.YLeaf{"CIgmpFilterEditEndAddress", cIgmpFilterEditor.CIgmpFilterEditEndAddress})
    cIgmpFilterEditor.EntityData.Leafs.Append("cIgmpFilterEditProfileAction", types.YLeaf{"CIgmpFilterEditProfileAction", cIgmpFilterEditor.CIgmpFilterEditProfileAction})
    cIgmpFilterEditor.EntityData.Leafs.Append("cIgmpFilterEditOperation", types.YLeaf{"CIgmpFilterEditOperation", cIgmpFilterEditor.CIgmpFilterEditOperation})
    cIgmpFilterEditor.EntityData.Leafs.Append("cIgmpFilterApplyStatus", types.YLeaf{"CIgmpFilterApplyStatus", cIgmpFilterEditor.CIgmpFilterApplyStatus})

    cIgmpFilterEditor.EntityData.YListKeys = []string {}

    return &(cIgmpFilterEditor.EntityData)
}

// CISCOIGMPFILTERMIB_CIgmpFilterEditor_CIgmpFilterApplyStatus represents    as no corresponding entry exists in cIgmpFilterTable.
type CISCOIGMPFILTERMIB_CIgmpFilterEditor_CIgmpFilterApplyStatus string

const (
    CISCOIGMPFILTERMIB_CIgmpFilterEditor_CIgmpFilterApplyStatus_someOtherError CISCOIGMPFILTERMIB_CIgmpFilterEditor_CIgmpFilterApplyStatus = "someOtherError"

    CISCOIGMPFILTERMIB_CIgmpFilterEditor_CIgmpFilterApplyStatus_succeeded CISCOIGMPFILTERMIB_CIgmpFilterEditor_CIgmpFilterApplyStatus = "succeeded"

    CISCOIGMPFILTERMIB_CIgmpFilterEditor_CIgmpFilterApplyStatus_inconsistentEdit CISCOIGMPFILTERMIB_CIgmpFilterEditor_CIgmpFilterApplyStatus = "inconsistentEdit"

    CISCOIGMPFILTERMIB_CIgmpFilterEditor_CIgmpFilterApplyStatus_entryPresentError CISCOIGMPFILTERMIB_CIgmpFilterEditor_CIgmpFilterApplyStatus = "entryPresentError"

    CISCOIGMPFILTERMIB_CIgmpFilterEditor_CIgmpFilterApplyStatus_entryNotPresentError CISCOIGMPFILTERMIB_CIgmpFilterEditor_CIgmpFilterApplyStatus = "entryNotPresentError"
)

// CISCOIGMPFILTERMIB_CIgmpFilterEditor_CIgmpFilterEditOperation represents 'none' - no operation is performed.
type CISCOIGMPFILTERMIB_CIgmpFilterEditor_CIgmpFilterEditOperation string

const (
    CISCOIGMPFILTERMIB_CIgmpFilterEditor_CIgmpFilterEditOperation_none CISCOIGMPFILTERMIB_CIgmpFilterEditor_CIgmpFilterEditOperation = "none"

    CISCOIGMPFILTERMIB_CIgmpFilterEditor_CIgmpFilterEditOperation_add CISCOIGMPFILTERMIB_CIgmpFilterEditor_CIgmpFilterEditOperation = "add"

    CISCOIGMPFILTERMIB_CIgmpFilterEditor_CIgmpFilterEditOperation_delete_ CISCOIGMPFILTERMIB_CIgmpFilterEditor_CIgmpFilterEditOperation = "delete"

    CISCOIGMPFILTERMIB_CIgmpFilterEditor_CIgmpFilterEditOperation_modify CISCOIGMPFILTERMIB_CIgmpFilterEditor_CIgmpFilterEditOperation = "modify"
)

// CISCOIGMPFILTERMIB_CIgmpFilterEditor_CIgmpFilterEditProfileAction represents of each profile.
type CISCOIGMPFILTERMIB_CIgmpFilterEditor_CIgmpFilterEditProfileAction string

const (
    CISCOIGMPFILTERMIB_CIgmpFilterEditor_CIgmpFilterEditProfileAction_unSpecified CISCOIGMPFILTERMIB_CIgmpFilterEditor_CIgmpFilterEditProfileAction = "unSpecified"

    CISCOIGMPFILTERMIB_CIgmpFilterEditor_CIgmpFilterEditProfileAction_permit CISCOIGMPFILTERMIB_CIgmpFilterEditor_CIgmpFilterEditProfileAction = "permit"

    CISCOIGMPFILTERMIB_CIgmpFilterEditor_CIgmpFilterEditProfileAction_deny CISCOIGMPFILTERMIB_CIgmpFilterEditor_CIgmpFilterEditProfileAction = "deny"
)

// CISCOIGMPFILTERMIB_CIgmpFilterTable
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
type CISCOIGMPFILTERMIB_CIgmpFilterTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the cIgmpFilterTable.  The creation, deletion
    // or modification of an entry is controlled through the MIB objects defined
    // under cIgmpFilterEditor group. The type is slice of
    // CISCOIGMPFILTERMIB_CIgmpFilterTable_CIgmpFilterEntry.
    CIgmpFilterEntry []*CISCOIGMPFILTERMIB_CIgmpFilterTable_CIgmpFilterEntry
}

func (cIgmpFilterTable *CISCOIGMPFILTERMIB_CIgmpFilterTable) GetEntityData() *types.CommonEntityData {
    cIgmpFilterTable.EntityData.YFilter = cIgmpFilterTable.YFilter
    cIgmpFilterTable.EntityData.YangName = "cIgmpFilterTable"
    cIgmpFilterTable.EntityData.BundleName = "cisco_ios_xe"
    cIgmpFilterTable.EntityData.ParentYangName = "CISCO-IGMP-FILTER-MIB"
    cIgmpFilterTable.EntityData.SegmentPath = "cIgmpFilterTable"
    cIgmpFilterTable.EntityData.AbsolutePath = "CISCO-IGMP-FILTER-MIB:CISCO-IGMP-FILTER-MIB/" + cIgmpFilterTable.EntityData.SegmentPath
    cIgmpFilterTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cIgmpFilterTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cIgmpFilterTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cIgmpFilterTable.EntityData.Children = types.NewOrderedMap()
    cIgmpFilterTable.EntityData.Children.Append("cIgmpFilterEntry", types.YChild{"CIgmpFilterEntry", nil})
    for i := range cIgmpFilterTable.CIgmpFilterEntry {
        cIgmpFilterTable.EntityData.Children.Append(types.GetSegmentPath(cIgmpFilterTable.CIgmpFilterEntry[i]), types.YChild{"CIgmpFilterEntry", cIgmpFilterTable.CIgmpFilterEntry[i]})
    }
    cIgmpFilterTable.EntityData.Leafs = types.NewOrderedMap()

    cIgmpFilterTable.EntityData.YListKeys = []string {}

    return &(cIgmpFilterTable.EntityData)
}

// CISCOIGMPFILTERMIB_CIgmpFilterTable_CIgmpFilterEntry
// An entry (conceptual row) in the cIgmpFilterTable.
// 
// The creation, deletion or modification of an entry
// is controlled through the MIB objects defined under
// cIgmpFilterEditor group.
type CISCOIGMPFILTERMIB_CIgmpFilterTable_CIgmpFilterEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Index identifying this entry. The type is
    // interface{} with range: 0..4294967295.
    CIgmpFilterProfileIndex interface{}

    // This attribute is a key. This object describes the type of Internet address
    // used to determine the start address  of IP multicast group for a profile.
    // The type is InetAddressType.
    CIgmpFilterStartAddressType interface{}

    // This attribute is a key. This object describes the start of the IP
    // multicast group address of a contiguous range which will be subjected to
    // filtering operation. The type is string with length: 1..64.
    CIgmpFilterStartAddress interface{}

    // This object indicates the type of Internet address used to determine the
    // end address  of IP multicast group for a profile. The type is
    // InetAddressType.
    CIgmpFilterEndAddressType interface{}

    // This object indicates the end of the IP multicast group address of a
    // contiguous range which will be  subjected to filtering operation. The type
    // is string with length: 0..255.
    CIgmpFilterEndAddress interface{}

    // This object defines the action for filtering IGMP reports for this profile.
    // If the object is set to deny(2): then all IGMP reports associated to IP
    // multicast groups included in the profile identified by
    // cIgmpFilterInterfaceProfileIndex will be dropped.  If the object is set to
    // permit(1): then all IGMP reports associated to IP multicast groups not
    // included in the profile identified by cIgmpFilterInterfaceProfileIndex will
    // be dropped. The type is CIgmpFilterProfileAction.
    CIgmpFilterProfileAction interface{}
}

func (cIgmpFilterEntry *CISCOIGMPFILTERMIB_CIgmpFilterTable_CIgmpFilterEntry) GetEntityData() *types.CommonEntityData {
    cIgmpFilterEntry.EntityData.YFilter = cIgmpFilterEntry.YFilter
    cIgmpFilterEntry.EntityData.YangName = "cIgmpFilterEntry"
    cIgmpFilterEntry.EntityData.BundleName = "cisco_ios_xe"
    cIgmpFilterEntry.EntityData.ParentYangName = "cIgmpFilterTable"
    cIgmpFilterEntry.EntityData.SegmentPath = "cIgmpFilterEntry" + types.AddKeyToken(cIgmpFilterEntry.CIgmpFilterProfileIndex, "cIgmpFilterProfileIndex") + types.AddKeyToken(cIgmpFilterEntry.CIgmpFilterStartAddressType, "cIgmpFilterStartAddressType") + types.AddKeyToken(cIgmpFilterEntry.CIgmpFilterStartAddress, "cIgmpFilterStartAddress")
    cIgmpFilterEntry.EntityData.AbsolutePath = "CISCO-IGMP-FILTER-MIB:CISCO-IGMP-FILTER-MIB/cIgmpFilterTable/" + cIgmpFilterEntry.EntityData.SegmentPath
    cIgmpFilterEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cIgmpFilterEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cIgmpFilterEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cIgmpFilterEntry.EntityData.Children = types.NewOrderedMap()
    cIgmpFilterEntry.EntityData.Leafs = types.NewOrderedMap()
    cIgmpFilterEntry.EntityData.Leafs.Append("cIgmpFilterProfileIndex", types.YLeaf{"CIgmpFilterProfileIndex", cIgmpFilterEntry.CIgmpFilterProfileIndex})
    cIgmpFilterEntry.EntityData.Leafs.Append("cIgmpFilterStartAddressType", types.YLeaf{"CIgmpFilterStartAddressType", cIgmpFilterEntry.CIgmpFilterStartAddressType})
    cIgmpFilterEntry.EntityData.Leafs.Append("cIgmpFilterStartAddress", types.YLeaf{"CIgmpFilterStartAddress", cIgmpFilterEntry.CIgmpFilterStartAddress})
    cIgmpFilterEntry.EntityData.Leafs.Append("cIgmpFilterEndAddressType", types.YLeaf{"CIgmpFilterEndAddressType", cIgmpFilterEntry.CIgmpFilterEndAddressType})
    cIgmpFilterEntry.EntityData.Leafs.Append("cIgmpFilterEndAddress", types.YLeaf{"CIgmpFilterEndAddress", cIgmpFilterEntry.CIgmpFilterEndAddress})
    cIgmpFilterEntry.EntityData.Leafs.Append("cIgmpFilterProfileAction", types.YLeaf{"CIgmpFilterProfileAction", cIgmpFilterEntry.CIgmpFilterProfileAction})

    cIgmpFilterEntry.EntityData.YListKeys = []string {"CIgmpFilterProfileIndex", "CIgmpFilterStartAddressType", "CIgmpFilterStartAddress"}

    return &(cIgmpFilterEntry.EntityData)
}

// CISCOIGMPFILTERMIB_CIgmpFilterTable_CIgmpFilterEntry_CIgmpFilterProfileAction represents cIgmpFilterInterfaceProfileIndex will be dropped.
type CISCOIGMPFILTERMIB_CIgmpFilterTable_CIgmpFilterEntry_CIgmpFilterProfileAction string

const (
    CISCOIGMPFILTERMIB_CIgmpFilterTable_CIgmpFilterEntry_CIgmpFilterProfileAction_permit CISCOIGMPFILTERMIB_CIgmpFilterTable_CIgmpFilterEntry_CIgmpFilterProfileAction = "permit"

    CISCOIGMPFILTERMIB_CIgmpFilterTable_CIgmpFilterEntry_CIgmpFilterProfileAction_deny CISCOIGMPFILTERMIB_CIgmpFilterTable_CIgmpFilterEntry_CIgmpFilterProfileAction = "deny"
)

// CISCOIGMPFILTERMIB_CIgmpFilterInterfaceTable
// This table contains the list of interfaces that can
// support IGMP filter feature.
type CISCOIGMPFILTERMIB_CIgmpFilterInterfaceTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry contains the configuration for associating the IGMP filter
    // profile index with the interface.  An entry is created for each of the IGMP
    // filter capable  interface on the system.  The entry is removed on removal
    // of corresponding  interface from system. The type is slice of
    // CISCOIGMPFILTERMIB_CIgmpFilterInterfaceTable_CIgmpFilterInterfaceEntry.
    CIgmpFilterInterfaceEntry []*CISCOIGMPFILTERMIB_CIgmpFilterInterfaceTable_CIgmpFilterInterfaceEntry
}

func (cIgmpFilterInterfaceTable *CISCOIGMPFILTERMIB_CIgmpFilterInterfaceTable) GetEntityData() *types.CommonEntityData {
    cIgmpFilterInterfaceTable.EntityData.YFilter = cIgmpFilterInterfaceTable.YFilter
    cIgmpFilterInterfaceTable.EntityData.YangName = "cIgmpFilterInterfaceTable"
    cIgmpFilterInterfaceTable.EntityData.BundleName = "cisco_ios_xe"
    cIgmpFilterInterfaceTable.EntityData.ParentYangName = "CISCO-IGMP-FILTER-MIB"
    cIgmpFilterInterfaceTable.EntityData.SegmentPath = "cIgmpFilterInterfaceTable"
    cIgmpFilterInterfaceTable.EntityData.AbsolutePath = "CISCO-IGMP-FILTER-MIB:CISCO-IGMP-FILTER-MIB/" + cIgmpFilterInterfaceTable.EntityData.SegmentPath
    cIgmpFilterInterfaceTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cIgmpFilterInterfaceTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cIgmpFilterInterfaceTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cIgmpFilterInterfaceTable.EntityData.Children = types.NewOrderedMap()
    cIgmpFilterInterfaceTable.EntityData.Children.Append("cIgmpFilterInterfaceEntry", types.YChild{"CIgmpFilterInterfaceEntry", nil})
    for i := range cIgmpFilterInterfaceTable.CIgmpFilterInterfaceEntry {
        cIgmpFilterInterfaceTable.EntityData.Children.Append(types.GetSegmentPath(cIgmpFilterInterfaceTable.CIgmpFilterInterfaceEntry[i]), types.YChild{"CIgmpFilterInterfaceEntry", cIgmpFilterInterfaceTable.CIgmpFilterInterfaceEntry[i]})
    }
    cIgmpFilterInterfaceTable.EntityData.Leafs = types.NewOrderedMap()

    cIgmpFilterInterfaceTable.EntityData.YListKeys = []string {}

    return &(cIgmpFilterInterfaceTable.EntityData)
}

// CISCOIGMPFILTERMIB_CIgmpFilterInterfaceTable_CIgmpFilterInterfaceEntry
// Each entry contains the configuration for associating
// the IGMP filter profile index with the interface.
// 
// An entry is created for each of the IGMP filter capable 
// interface on the system.
// 
// The entry is removed on removal of corresponding 
// interface from system.
type CISCOIGMPFILTERMIB_CIgmpFilterInterfaceTable_CIgmpFilterInterfaceEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // Specifies which IGMP filter profile applies to this interface. If the value
    // of this MIB object matches the  value of cIgmpFilterProfileIndex in
    // cIgmpFilterTable,  the corresponding profile configuration will apply to
    // this interface.   Agent returns inconsistentValue if this profile  does not
    // exist in cIgmpFilterTable.  A value of zero indicates no profile is
    // associated with corresponding interface.  The filtering action on each
    // interface is also defined by the associated profile. The type is
    // interface{} with range: 0..4294967295.
    CIgmpFilterInterfaceProfileIndex interface{}
}

func (cIgmpFilterInterfaceEntry *CISCOIGMPFILTERMIB_CIgmpFilterInterfaceTable_CIgmpFilterInterfaceEntry) GetEntityData() *types.CommonEntityData {
    cIgmpFilterInterfaceEntry.EntityData.YFilter = cIgmpFilterInterfaceEntry.YFilter
    cIgmpFilterInterfaceEntry.EntityData.YangName = "cIgmpFilterInterfaceEntry"
    cIgmpFilterInterfaceEntry.EntityData.BundleName = "cisco_ios_xe"
    cIgmpFilterInterfaceEntry.EntityData.ParentYangName = "cIgmpFilterInterfaceTable"
    cIgmpFilterInterfaceEntry.EntityData.SegmentPath = "cIgmpFilterInterfaceEntry" + types.AddKeyToken(cIgmpFilterInterfaceEntry.IfIndex, "ifIndex")
    cIgmpFilterInterfaceEntry.EntityData.AbsolutePath = "CISCO-IGMP-FILTER-MIB:CISCO-IGMP-FILTER-MIB/cIgmpFilterInterfaceTable/" + cIgmpFilterInterfaceEntry.EntityData.SegmentPath
    cIgmpFilterInterfaceEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cIgmpFilterInterfaceEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cIgmpFilterInterfaceEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cIgmpFilterInterfaceEntry.EntityData.Children = types.NewOrderedMap()
    cIgmpFilterInterfaceEntry.EntityData.Leafs = types.NewOrderedMap()
    cIgmpFilterInterfaceEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", cIgmpFilterInterfaceEntry.IfIndex})
    cIgmpFilterInterfaceEntry.EntityData.Leafs.Append("cIgmpFilterInterfaceProfileIndex", types.YLeaf{"CIgmpFilterInterfaceProfileIndex", cIgmpFilterInterfaceEntry.CIgmpFilterInterfaceProfileIndex})

    cIgmpFilterInterfaceEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(cIgmpFilterInterfaceEntry.EntityData)
}

