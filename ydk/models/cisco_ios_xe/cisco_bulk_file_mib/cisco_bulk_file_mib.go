// The MIB module for creating and deleting bulk files of
// SNMP data for file transfer.
package cisco_bulk_file_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_bulk_file_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-BULK-FILE-MIB CISCO-BULK-FILE-MIB}", reflect.TypeOf(CISCOBULKFILEMIB{}))
    ydk.RegisterEntity("CISCO-BULK-FILE-MIB:CISCO-BULK-FILE-MIB", reflect.TypeOf(CISCOBULKFILEMIB{}))
}

// CISCOBULKFILEMIB
type CISCOBULKFILEMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    CbfDefine CISCOBULKFILEMIB_CbfDefine

    
    CbfStatus CISCOBULKFILEMIB_CbfStatus

    // A table of bulk file definition and creation controls.
    CbfDefineFileTable CISCOBULKFILEMIB_CbfDefineFileTable

    // A table of objects to go in bulk files.
    CbfDefineObjectTable CISCOBULKFILEMIB_CbfDefineObjectTable

    // A table of bulk file status.
    CbfStatusFileTable CISCOBULKFILEMIB_CbfStatusFileTable
}

func (cISCOBULKFILEMIB *CISCOBULKFILEMIB) GetEntityData() *types.CommonEntityData {
    cISCOBULKFILEMIB.EntityData.YFilter = cISCOBULKFILEMIB.YFilter
    cISCOBULKFILEMIB.EntityData.YangName = "CISCO-BULK-FILE-MIB"
    cISCOBULKFILEMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOBULKFILEMIB.EntityData.ParentYangName = "CISCO-BULK-FILE-MIB"
    cISCOBULKFILEMIB.EntityData.SegmentPath = "CISCO-BULK-FILE-MIB:CISCO-BULK-FILE-MIB"
    cISCOBULKFILEMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOBULKFILEMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOBULKFILEMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOBULKFILEMIB.EntityData.Children = types.NewOrderedMap()
    cISCOBULKFILEMIB.EntityData.Children.Append("cbfDefine", types.YChild{"CbfDefine", &cISCOBULKFILEMIB.CbfDefine})
    cISCOBULKFILEMIB.EntityData.Children.Append("cbfStatus", types.YChild{"CbfStatus", &cISCOBULKFILEMIB.CbfStatus})
    cISCOBULKFILEMIB.EntityData.Children.Append("cbfDefineFileTable", types.YChild{"CbfDefineFileTable", &cISCOBULKFILEMIB.CbfDefineFileTable})
    cISCOBULKFILEMIB.EntityData.Children.Append("cbfDefineObjectTable", types.YChild{"CbfDefineObjectTable", &cISCOBULKFILEMIB.CbfDefineObjectTable})
    cISCOBULKFILEMIB.EntityData.Children.Append("cbfStatusFileTable", types.YChild{"CbfStatusFileTable", &cISCOBULKFILEMIB.CbfStatusFileTable})
    cISCOBULKFILEMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOBULKFILEMIB.EntityData.YListKeys = []string {}

    return &(cISCOBULKFILEMIB.EntityData)
}

// CISCOBULKFILEMIB_CbfDefine
type CISCOBULKFILEMIB_CbfDefine struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The maximum number of file definitions this system can hold in
    // cbfDefineFileTable.  A value of 0 indicates no configured limit.  This
    // object may be read-only on some systems.  Changing this number does not
    // disturb existing entries. The type is interface{} with range:
    // 0..4294967295.
    CbfDefineMaxFiles interface{}

    // The current number of file definitions in cbfDefineFileTable. The type is
    // interface{} with range: 0..4294967295.
    CbfDefineFiles interface{}

    // The maximum value of cbfDefineFiles since system  initialization. The type
    // is interface{} with range: 0..4294967295.
    CbfDefineHighFiles interface{}

    // The number of attempts to create a file definition that failed due to
    // exceeding cbfDefineMaxFiles. The type is interface{} with range:
    // 0..4294967295.
    CbfDefineFilesRefused interface{}

    // The maximum total number of object selections to go with file definitions
    // this system, that is, the total number of objects this system can hold in
    // cbfDefineObjectTable.  A value of 0 indicates no configured limit.  This
    // object may be read-only on some systems.  Changing this number does not
    // disturb existing entries. The type is interface{} with range:
    // 0..4294967295.
    CbfDefineMaxObjects interface{}

    // The current number of object selections in  cbfDefineObjectTable. The type
    // is interface{} with range: 0..4294967295.
    CbfDefineObjects interface{}

    // The maximum value of cbfDefineObjects since system  initialization. The
    // type is interface{} with range: 0..4294967295.
    CbfDefineHighObjects interface{}

    // The number of attempts to create an object selection that failed due to
    // exceeding cbfDefineMaxObjects. The type is interface{} with range:
    // 0..4294967295.
    CbfDefineObjectsRefused interface{}
}

func (cbfDefine *CISCOBULKFILEMIB_CbfDefine) GetEntityData() *types.CommonEntityData {
    cbfDefine.EntityData.YFilter = cbfDefine.YFilter
    cbfDefine.EntityData.YangName = "cbfDefine"
    cbfDefine.EntityData.BundleName = "cisco_ios_xe"
    cbfDefine.EntityData.ParentYangName = "CISCO-BULK-FILE-MIB"
    cbfDefine.EntityData.SegmentPath = "cbfDefine"
    cbfDefine.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cbfDefine.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cbfDefine.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cbfDefine.EntityData.Children = types.NewOrderedMap()
    cbfDefine.EntityData.Leafs = types.NewOrderedMap()
    cbfDefine.EntityData.Leafs.Append("cbfDefineMaxFiles", types.YLeaf{"CbfDefineMaxFiles", cbfDefine.CbfDefineMaxFiles})
    cbfDefine.EntityData.Leafs.Append("cbfDefineFiles", types.YLeaf{"CbfDefineFiles", cbfDefine.CbfDefineFiles})
    cbfDefine.EntityData.Leafs.Append("cbfDefineHighFiles", types.YLeaf{"CbfDefineHighFiles", cbfDefine.CbfDefineHighFiles})
    cbfDefine.EntityData.Leafs.Append("cbfDefineFilesRefused", types.YLeaf{"CbfDefineFilesRefused", cbfDefine.CbfDefineFilesRefused})
    cbfDefine.EntityData.Leafs.Append("cbfDefineMaxObjects", types.YLeaf{"CbfDefineMaxObjects", cbfDefine.CbfDefineMaxObjects})
    cbfDefine.EntityData.Leafs.Append("cbfDefineObjects", types.YLeaf{"CbfDefineObjects", cbfDefine.CbfDefineObjects})
    cbfDefine.EntityData.Leafs.Append("cbfDefineHighObjects", types.YLeaf{"CbfDefineHighObjects", cbfDefine.CbfDefineHighObjects})
    cbfDefine.EntityData.Leafs.Append("cbfDefineObjectsRefused", types.YLeaf{"CbfDefineObjectsRefused", cbfDefine.CbfDefineObjectsRefused})

    cbfDefine.EntityData.YListKeys = []string {}

    return &(cbfDefine.EntityData)
}

// CISCOBULKFILEMIB_CbfStatus
type CISCOBULKFILEMIB_CbfStatus struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The maximum number of file statuses this system can hold in
    // cbfStatusFileTable.  A value of 0 indicates no configured limit.  This
    // object may be read-only on some systems.  Changing this number deletes the
    // oldest finished entries until the new limit is satisfied. The type is
    // interface{} with range: 0..4294967295.
    CbfStatusMaxFiles interface{}

    // The current number of file statuses in cbfStatusFileTable. The type is
    // interface{} with range: 0..4294967295.
    CbfStatusFiles interface{}

    // The maximum value of cbfStatusFiles since system  initialization. The type
    // is interface{} with range: 0..4294967295.
    CbfStatusHighFiles interface{}

    // The number times the oldest entry was deleted due to exceeding
    // cbfStatusMaxFiles. The type is interface{} with range: 0..4294967295.
    CbfStatusFilesBumped interface{}
}

func (cbfStatus *CISCOBULKFILEMIB_CbfStatus) GetEntityData() *types.CommonEntityData {
    cbfStatus.EntityData.YFilter = cbfStatus.YFilter
    cbfStatus.EntityData.YangName = "cbfStatus"
    cbfStatus.EntityData.BundleName = "cisco_ios_xe"
    cbfStatus.EntityData.ParentYangName = "CISCO-BULK-FILE-MIB"
    cbfStatus.EntityData.SegmentPath = "cbfStatus"
    cbfStatus.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cbfStatus.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cbfStatus.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cbfStatus.EntityData.Children = types.NewOrderedMap()
    cbfStatus.EntityData.Leafs = types.NewOrderedMap()
    cbfStatus.EntityData.Leafs.Append("cbfStatusMaxFiles", types.YLeaf{"CbfStatusMaxFiles", cbfStatus.CbfStatusMaxFiles})
    cbfStatus.EntityData.Leafs.Append("cbfStatusFiles", types.YLeaf{"CbfStatusFiles", cbfStatus.CbfStatusFiles})
    cbfStatus.EntityData.Leafs.Append("cbfStatusHighFiles", types.YLeaf{"CbfStatusHighFiles", cbfStatus.CbfStatusHighFiles})
    cbfStatus.EntityData.Leafs.Append("cbfStatusFilesBumped", types.YLeaf{"CbfStatusFilesBumped", cbfStatus.CbfStatusFilesBumped})

    cbfStatus.EntityData.YListKeys = []string {}

    return &(cbfStatus.EntityData)
}

// CISCOBULKFILEMIB_CbfDefineFileTable
// A table of bulk file definition and creation controls.
type CISCOBULKFILEMIB_CbfDefineFileTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information for creation of a bulk file.  To creat a bulk file an
    // application creates an entry in this table and corresponding entries in
    // cbfDefineObjectTable.  When the entry in this table and the corresponding
    // entries in cbfDefineObjectTable are 'active' the appliction uses
    // cbfDefineFileNow to create the file and a corresponding entry in
    // cbfStatusFileTable.  Deleting an entry in cbfDefineFileTable deletes all
    // corresponding entries in cbfDefineObjectTable and cbfStatusFileTable.  An
    // entry may not be modified or deleted while its cbfDefineFileNow has the
    // value 'running'. The type is slice of
    // CISCOBULKFILEMIB_CbfDefineFileTable_CbfDefineFileEntry.
    CbfDefineFileEntry []*CISCOBULKFILEMIB_CbfDefineFileTable_CbfDefineFileEntry
}

func (cbfDefineFileTable *CISCOBULKFILEMIB_CbfDefineFileTable) GetEntityData() *types.CommonEntityData {
    cbfDefineFileTable.EntityData.YFilter = cbfDefineFileTable.YFilter
    cbfDefineFileTable.EntityData.YangName = "cbfDefineFileTable"
    cbfDefineFileTable.EntityData.BundleName = "cisco_ios_xe"
    cbfDefineFileTable.EntityData.ParentYangName = "CISCO-BULK-FILE-MIB"
    cbfDefineFileTable.EntityData.SegmentPath = "cbfDefineFileTable"
    cbfDefineFileTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cbfDefineFileTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cbfDefineFileTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cbfDefineFileTable.EntityData.Children = types.NewOrderedMap()
    cbfDefineFileTable.EntityData.Children.Append("cbfDefineFileEntry", types.YChild{"CbfDefineFileEntry", nil})
    for i := range cbfDefineFileTable.CbfDefineFileEntry {
        cbfDefineFileTable.EntityData.Children.Append(types.GetSegmentPath(cbfDefineFileTable.CbfDefineFileEntry[i]), types.YChild{"CbfDefineFileEntry", cbfDefineFileTable.CbfDefineFileEntry[i]})
    }
    cbfDefineFileTable.EntityData.Leafs = types.NewOrderedMap()

    cbfDefineFileTable.EntityData.YListKeys = []string {}

    return &(cbfDefineFileTable.EntityData)
}

// CISCOBULKFILEMIB_CbfDefineFileTable_CbfDefineFileEntry
// Information for creation of a bulk file.
// 
// To creat a bulk file an application creates an entry in this
// table and corresponding entries in cbfDefineObjectTable.
// 
// When the entry in this table and the corresponding
// entries in cbfDefineObjectTable are 'active' the
// appliction uses cbfDefineFileNow to create the file
// and a corresponding entry in cbfStatusFileTable.
// 
// Deleting an entry in cbfDefineFileTable deletes all
// corresponding entries in cbfDefineObjectTable and
// cbfStatusFileTable.
// 
// An entry may not be modified or deleted while its
// cbfDefineFileNow has the value 'running'.
type CISCOBULKFILEMIB_CbfDefineFileTable_CbfDefineFileEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An arbitrary integer to uniquely identify this
    // entry.  To create an entry a management application should pick a random
    // number. The type is interface{} with range: 1..4294967295.
    CbfDefineFileIndex interface{}

    // The file name which is to be created.  Explicit device or path choices in
    // the value of this object override cbfDefineFileStorage. The type is string
    // with length: 0..255.
    CbfDefineFileName interface{}

    // The type of file storage to use:  ephemeral        data exists in small
    // amounts until read volatile        data exists in volatile memory permanent
    // data survives reboot  An ephemeral file is suitable to be read only one
    // time.  Note that this value is taken as advisory and may be overridden by
    // explicit device or path choices in cbfDefineFile.  A given system may
    // support any or all of these. The type is CbfDefineFileStorage.
    CbfDefineFileStorage interface{}

    // The format of the data in the file:  StandardBER        standard SNMP ASN.1
    // BER bulkBinary        a binary format specified with this MIB bulkASCII    
    // a human-readable form of bulkBinary variantBERWithCksum ASN.1 BER encoding
    // with checksum variantBinWithCksum a binary format with checksum      A
    // given system may support any or all of these. The type is
    // CbfDefineFileFormat.
    CbfDefineFileFormat interface{}

    // The control to cause file creation.  The only values that can be set are
    // 'create' and 'forcedCreate'. These can be set only  when the value is
    // 'ready'.  Setting it to 'create' begins a  file creation and creates a
    // corresponding entry in  cbfStatusFileTable. The system may choose to use an
    // already  existing copy of the file instead of creating a new one. This may
    // happen if there has been no configuration change on the  system and a
    // request to recreate the file is received.  Setting this object to
    // 'forcedCreate' forces the system to  create a new copy of the file.  The
    // value is 'notActve' as long as cbfDefineFileEntryStatus or any
    // corresponding cbfDefineObjectEntryStatus is not active.  When
    // cbfDefineFileEntryStatus becomes active and all corresponding
    // cbfDefineObjectEntryStatuses are active this  object automatically goes to
    // 'ready'. The type is CbfDefineFileNow.
    CbfDefineFileNow interface{}

    // The control that allows creation, modification, and deletion of entries. 
    // For detailed rules see the DESCRIPTION for cbfDefineFileEntry. The type is
    // RowStatus.
    CbfDefineFileEntryStatus interface{}

    // This controls the cbfDefineFileCompletion notification.  If true,
    // cbfDefineFileCompletion notification will be generated. It is the
    // responsibility of the  management entity to ensure that the SNMP
    // administrative  model is configured in such a way as to allow the 
    // notification to be delivered. The type is bool.
    CbfDefineFileNotifyOnCompletion interface{}
}

func (cbfDefineFileEntry *CISCOBULKFILEMIB_CbfDefineFileTable_CbfDefineFileEntry) GetEntityData() *types.CommonEntityData {
    cbfDefineFileEntry.EntityData.YFilter = cbfDefineFileEntry.YFilter
    cbfDefineFileEntry.EntityData.YangName = "cbfDefineFileEntry"
    cbfDefineFileEntry.EntityData.BundleName = "cisco_ios_xe"
    cbfDefineFileEntry.EntityData.ParentYangName = "cbfDefineFileTable"
    cbfDefineFileEntry.EntityData.SegmentPath = "cbfDefineFileEntry" + types.AddKeyToken(cbfDefineFileEntry.CbfDefineFileIndex, "cbfDefineFileIndex")
    cbfDefineFileEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cbfDefineFileEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cbfDefineFileEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cbfDefineFileEntry.EntityData.Children = types.NewOrderedMap()
    cbfDefineFileEntry.EntityData.Leafs = types.NewOrderedMap()
    cbfDefineFileEntry.EntityData.Leafs.Append("cbfDefineFileIndex", types.YLeaf{"CbfDefineFileIndex", cbfDefineFileEntry.CbfDefineFileIndex})
    cbfDefineFileEntry.EntityData.Leafs.Append("cbfDefineFileName", types.YLeaf{"CbfDefineFileName", cbfDefineFileEntry.CbfDefineFileName})
    cbfDefineFileEntry.EntityData.Leafs.Append("cbfDefineFileStorage", types.YLeaf{"CbfDefineFileStorage", cbfDefineFileEntry.CbfDefineFileStorage})
    cbfDefineFileEntry.EntityData.Leafs.Append("cbfDefineFileFormat", types.YLeaf{"CbfDefineFileFormat", cbfDefineFileEntry.CbfDefineFileFormat})
    cbfDefineFileEntry.EntityData.Leafs.Append("cbfDefineFileNow", types.YLeaf{"CbfDefineFileNow", cbfDefineFileEntry.CbfDefineFileNow})
    cbfDefineFileEntry.EntityData.Leafs.Append("cbfDefineFileEntryStatus", types.YLeaf{"CbfDefineFileEntryStatus", cbfDefineFileEntry.CbfDefineFileEntryStatus})
    cbfDefineFileEntry.EntityData.Leafs.Append("cbfDefineFileNotifyOnCompletion", types.YLeaf{"CbfDefineFileNotifyOnCompletion", cbfDefineFileEntry.CbfDefineFileNotifyOnCompletion})

    cbfDefineFileEntry.EntityData.YListKeys = []string {"CbfDefineFileIndex"}

    return &(cbfDefineFileEntry.EntityData)
}

// CISCOBULKFILEMIB_CbfDefineFileTable_CbfDefineFileEntry_CbfDefineFileFormat represents     A given system may support any or all of these.
type CISCOBULKFILEMIB_CbfDefineFileTable_CbfDefineFileEntry_CbfDefineFileFormat string

const (
    CISCOBULKFILEMIB_CbfDefineFileTable_CbfDefineFileEntry_CbfDefineFileFormat_standardBER CISCOBULKFILEMIB_CbfDefineFileTable_CbfDefineFileEntry_CbfDefineFileFormat = "standardBER"

    CISCOBULKFILEMIB_CbfDefineFileTable_CbfDefineFileEntry_CbfDefineFileFormat_bulkBinary CISCOBULKFILEMIB_CbfDefineFileTable_CbfDefineFileEntry_CbfDefineFileFormat = "bulkBinary"

    CISCOBULKFILEMIB_CbfDefineFileTable_CbfDefineFileEntry_CbfDefineFileFormat_bulkASCII CISCOBULKFILEMIB_CbfDefineFileTable_CbfDefineFileEntry_CbfDefineFileFormat = "bulkASCII"

    CISCOBULKFILEMIB_CbfDefineFileTable_CbfDefineFileEntry_CbfDefineFileFormat_variantBERWithCksum CISCOBULKFILEMIB_CbfDefineFileTable_CbfDefineFileEntry_CbfDefineFileFormat = "variantBERWithCksum"

    CISCOBULKFILEMIB_CbfDefineFileTable_CbfDefineFileEntry_CbfDefineFileFormat_variantBinWithCksum CISCOBULKFILEMIB_CbfDefineFileTable_CbfDefineFileEntry_CbfDefineFileFormat = "variantBinWithCksum"
)

// CISCOBULKFILEMIB_CbfDefineFileTable_CbfDefineFileEntry_CbfDefineFileNow represents object automatically goes to 'ready'.
type CISCOBULKFILEMIB_CbfDefineFileTable_CbfDefineFileEntry_CbfDefineFileNow string

const (
    CISCOBULKFILEMIB_CbfDefineFileTable_CbfDefineFileEntry_CbfDefineFileNow_notActive CISCOBULKFILEMIB_CbfDefineFileTable_CbfDefineFileEntry_CbfDefineFileNow = "notActive"

    CISCOBULKFILEMIB_CbfDefineFileTable_CbfDefineFileEntry_CbfDefineFileNow_ready CISCOBULKFILEMIB_CbfDefineFileTable_CbfDefineFileEntry_CbfDefineFileNow = "ready"

    CISCOBULKFILEMIB_CbfDefineFileTable_CbfDefineFileEntry_CbfDefineFileNow_create CISCOBULKFILEMIB_CbfDefineFileTable_CbfDefineFileEntry_CbfDefineFileNow = "create"

    CISCOBULKFILEMIB_CbfDefineFileTable_CbfDefineFileEntry_CbfDefineFileNow_running CISCOBULKFILEMIB_CbfDefineFileTable_CbfDefineFileEntry_CbfDefineFileNow = "running"

    CISCOBULKFILEMIB_CbfDefineFileTable_CbfDefineFileEntry_CbfDefineFileNow_forcedCreate CISCOBULKFILEMIB_CbfDefineFileTable_CbfDefineFileEntry_CbfDefineFileNow = "forcedCreate"
)

// CISCOBULKFILEMIB_CbfDefineFileTable_CbfDefineFileEntry_CbfDefineFileStorage represents A given system may support any or all of these.
type CISCOBULKFILEMIB_CbfDefineFileTable_CbfDefineFileEntry_CbfDefineFileStorage string

const (
    CISCOBULKFILEMIB_CbfDefineFileTable_CbfDefineFileEntry_CbfDefineFileStorage_ephemeral CISCOBULKFILEMIB_CbfDefineFileTable_CbfDefineFileEntry_CbfDefineFileStorage = "ephemeral"

    CISCOBULKFILEMIB_CbfDefineFileTable_CbfDefineFileEntry_CbfDefineFileStorage_volatile CISCOBULKFILEMIB_CbfDefineFileTable_CbfDefineFileEntry_CbfDefineFileStorage = "volatile"

    CISCOBULKFILEMIB_CbfDefineFileTable_CbfDefineFileEntry_CbfDefineFileStorage_permanent CISCOBULKFILEMIB_CbfDefineFileTable_CbfDefineFileEntry_CbfDefineFileStorage = "permanent"
)

// CISCOBULKFILEMIB_CbfDefineObjectTable
// A table of objects to go in bulk files.
type CISCOBULKFILEMIB_CbfDefineObjectTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about one object for a particular file.  An application uses
    // cbfDefineObjectEntryStatus to create entries in this table in
    // correspondence with entries in cbfDefineFileTable, which must be created
    // first.  Entries in this table may not be changed, created or deleted while
    // the corresponding value of cbfDefineFileNow is 'running'. The type is slice
    // of CISCOBULKFILEMIB_CbfDefineObjectTable_CbfDefineObjectEntry.
    CbfDefineObjectEntry []*CISCOBULKFILEMIB_CbfDefineObjectTable_CbfDefineObjectEntry
}

func (cbfDefineObjectTable *CISCOBULKFILEMIB_CbfDefineObjectTable) GetEntityData() *types.CommonEntityData {
    cbfDefineObjectTable.EntityData.YFilter = cbfDefineObjectTable.YFilter
    cbfDefineObjectTable.EntityData.YangName = "cbfDefineObjectTable"
    cbfDefineObjectTable.EntityData.BundleName = "cisco_ios_xe"
    cbfDefineObjectTable.EntityData.ParentYangName = "CISCO-BULK-FILE-MIB"
    cbfDefineObjectTable.EntityData.SegmentPath = "cbfDefineObjectTable"
    cbfDefineObjectTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cbfDefineObjectTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cbfDefineObjectTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cbfDefineObjectTable.EntityData.Children = types.NewOrderedMap()
    cbfDefineObjectTable.EntityData.Children.Append("cbfDefineObjectEntry", types.YChild{"CbfDefineObjectEntry", nil})
    for i := range cbfDefineObjectTable.CbfDefineObjectEntry {
        cbfDefineObjectTable.EntityData.Children.Append(types.GetSegmentPath(cbfDefineObjectTable.CbfDefineObjectEntry[i]), types.YChild{"CbfDefineObjectEntry", cbfDefineObjectTable.CbfDefineObjectEntry[i]})
    }
    cbfDefineObjectTable.EntityData.Leafs = types.NewOrderedMap()

    cbfDefineObjectTable.EntityData.YListKeys = []string {}

    return &(cbfDefineObjectTable.EntityData)
}

// CISCOBULKFILEMIB_CbfDefineObjectTable_CbfDefineObjectEntry
// Information about one object for a particular file.
// 
// An application uses cbfDefineObjectEntryStatus to create entries
// in this table in correspondence with entries in
// cbfDefineFileTable, which must be created first.
// 
// Entries in this table may not be changed, created or deleted
// while the corresponding value of cbfDefineFileNow is 'running'.
type CISCOBULKFILEMIB_CbfDefineObjectTable_CbfDefineObjectEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // cisco_bulk_file_mib.CISCOBULKFILEMIB_CbfDefineFileTable_CbfDefineFileEntry_CbfDefineFileIndex
    CbfDefineFileIndex interface{}

    // This attribute is a key. An arbitrary integer to uniquely identify this
    // entry.  The numeric order of the entries controls the order of the objects
    // in the file. The type is interface{} with range: 1..4294967295.
    CbfDefineObjectIndex interface{}

    // The meaning of each object class is given below:  object          a single
    // MIB object is retrieved.  lexicalTable    an entire table or partial table 
    // is retrieved in lexical order of rows.  leastCpuTable   an entire table is
    // retrieved with                 lowest CPU utilization.                
    // Lexical ordering of rows may not be                  maintained and is
    // dependent upon                  individual MIB implementation. The type is
    // CbfDefineObjectClass.
    CbfDefineObjectClass interface{}

    // The object identifier of a MIB object to be included in the file.  If
    // cbfDefineObjectClass is 'object' this must be a full OID, including all
    // instance information.  If cbfDefineObjectClass is 'lexicalTable' or
    // 'leastCpuTable' this must be the OID of the table-defining SEQUENCE OF
    // registration point. The type is string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    CbfDefineObjectID interface{}

    // The control that allows creation, modification, and deletion of entries. 
    // For detailed rules see the DESCRIPTION for cbfDefineObjectEntry. The type
    // is RowStatus.
    CbfDefineObjectEntryStatus interface{}

    // If cbfDefineObjectClass is 'lexicalTable', then this object represents the
    // starting instance in the cbfDefineObjectID table. The file created will
    // have entries starting from the lexicographically next instance of the OID
    // represented by this object.   For Eg:  -------         Let us  assume we
    // are polling ifTable and we        have information till the second
    // row(ifIndex.2). Now        we may be interested in 10 rows lexically
    // following        the second row.                So, we set
    // cbfDefineObjectTableInstance as ifIndex.2         and
    // cbfDefineObjectNumEntries as 10.          We will get information for the
    // next 10 rows or        if there are less than 10 populated rows, we will   
    // receive information till the end of the table is         reached.  The
    // default value for this object is zeroDotZero.  If this object has the value
    // of zeroDotZero and  cbfDefineObjectNumEntries has value 0, then the whole
    // table(represented by cbfDefineObjectID) is retrieved.  If this object has
    // the value of zeroDotZero,   cbfDefineObjectNumEntries has value n (>0) and
    // there are  m(>0) entries in the table(represented by cbfDefineObjectID)
    // then the first n entries in the table are retrieved if n < m.  If n >= m,
    // then the whole table is retrieved.  When the value of
    // cbfDefineObjectNumEntries is 0,  it means all the entries in the
    // table(represented  by cbfDefineObjectID) which lexicographically follow 
    // cbfDefineObjectTableInstance are retrieved.  This object is irrelevent if
    // cbfDefineObjectClass is not 'lexicalTable'. The type is string with
    // pattern: (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    CbfDefineObjectTableInstance interface{}

    // If cbfDefineObjectClass is 'lexicalTable', then this object represents the
    // maximum number of entries which will be  populated in the file starting
    // from the lexicographically next instance of the OID represented by 
    // cbfDefineObjectTableInstance.   This object is irrelevent if
    // cbfDefineObjectClass is not 'lexicalTable'.  Refer to the description of
    // cbfDefineObjectTableInstance for examples and different scenarios relating
    // to this object. The type is interface{} with range: 0..4294967295.
    CbfDefineObjectNumEntries interface{}

    // This object represents the last polled instance in the table.  The value
    // represented by this object will be relevent only if the corresponding
    // cbfStatusFileState is emptied(3) for  ephemeral files or ready(2) for
    // volatile or permanent files.  A value of zeroDotZero indicates an absence
    // of last polled  object.  An NMS can use the value of this object and
    // populate the cbfDefineObjectTableInstance to retrieve a contiguous set of
    // rows in a table. The type is string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    CbfDefineObjectLastPolledInst interface{}
}

func (cbfDefineObjectEntry *CISCOBULKFILEMIB_CbfDefineObjectTable_CbfDefineObjectEntry) GetEntityData() *types.CommonEntityData {
    cbfDefineObjectEntry.EntityData.YFilter = cbfDefineObjectEntry.YFilter
    cbfDefineObjectEntry.EntityData.YangName = "cbfDefineObjectEntry"
    cbfDefineObjectEntry.EntityData.BundleName = "cisco_ios_xe"
    cbfDefineObjectEntry.EntityData.ParentYangName = "cbfDefineObjectTable"
    cbfDefineObjectEntry.EntityData.SegmentPath = "cbfDefineObjectEntry" + types.AddKeyToken(cbfDefineObjectEntry.CbfDefineFileIndex, "cbfDefineFileIndex") + types.AddKeyToken(cbfDefineObjectEntry.CbfDefineObjectIndex, "cbfDefineObjectIndex")
    cbfDefineObjectEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cbfDefineObjectEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cbfDefineObjectEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cbfDefineObjectEntry.EntityData.Children = types.NewOrderedMap()
    cbfDefineObjectEntry.EntityData.Leafs = types.NewOrderedMap()
    cbfDefineObjectEntry.EntityData.Leafs.Append("cbfDefineFileIndex", types.YLeaf{"CbfDefineFileIndex", cbfDefineObjectEntry.CbfDefineFileIndex})
    cbfDefineObjectEntry.EntityData.Leafs.Append("cbfDefineObjectIndex", types.YLeaf{"CbfDefineObjectIndex", cbfDefineObjectEntry.CbfDefineObjectIndex})
    cbfDefineObjectEntry.EntityData.Leafs.Append("cbfDefineObjectClass", types.YLeaf{"CbfDefineObjectClass", cbfDefineObjectEntry.CbfDefineObjectClass})
    cbfDefineObjectEntry.EntityData.Leafs.Append("cbfDefineObjectID", types.YLeaf{"CbfDefineObjectID", cbfDefineObjectEntry.CbfDefineObjectID})
    cbfDefineObjectEntry.EntityData.Leafs.Append("cbfDefineObjectEntryStatus", types.YLeaf{"CbfDefineObjectEntryStatus", cbfDefineObjectEntry.CbfDefineObjectEntryStatus})
    cbfDefineObjectEntry.EntityData.Leafs.Append("cbfDefineObjectTableInstance", types.YLeaf{"CbfDefineObjectTableInstance", cbfDefineObjectEntry.CbfDefineObjectTableInstance})
    cbfDefineObjectEntry.EntityData.Leafs.Append("cbfDefineObjectNumEntries", types.YLeaf{"CbfDefineObjectNumEntries", cbfDefineObjectEntry.CbfDefineObjectNumEntries})
    cbfDefineObjectEntry.EntityData.Leafs.Append("cbfDefineObjectLastPolledInst", types.YLeaf{"CbfDefineObjectLastPolledInst", cbfDefineObjectEntry.CbfDefineObjectLastPolledInst})

    cbfDefineObjectEntry.EntityData.YListKeys = []string {"CbfDefineFileIndex", "CbfDefineObjectIndex"}

    return &(cbfDefineObjectEntry.EntityData)
}

// CISCOBULKFILEMIB_CbfDefineObjectTable_CbfDefineObjectEntry_CbfDefineObjectClass represents                 individual MIB implementation.
type CISCOBULKFILEMIB_CbfDefineObjectTable_CbfDefineObjectEntry_CbfDefineObjectClass string

const (
    CISCOBULKFILEMIB_CbfDefineObjectTable_CbfDefineObjectEntry_CbfDefineObjectClass_object CISCOBULKFILEMIB_CbfDefineObjectTable_CbfDefineObjectEntry_CbfDefineObjectClass = "object"

    CISCOBULKFILEMIB_CbfDefineObjectTable_CbfDefineObjectEntry_CbfDefineObjectClass_lexicalTable CISCOBULKFILEMIB_CbfDefineObjectTable_CbfDefineObjectEntry_CbfDefineObjectClass = "lexicalTable"

    CISCOBULKFILEMIB_CbfDefineObjectTable_CbfDefineObjectEntry_CbfDefineObjectClass_leastCpuTable CISCOBULKFILEMIB_CbfDefineObjectTable_CbfDefineObjectEntry_CbfDefineObjectClass = "leastCpuTable"
)

// CISCOBULKFILEMIB_CbfStatusFileTable
// A table of bulk file status.
type CISCOBULKFILEMIB_CbfStatusFileTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Status for a particular file.  An entry exists in this table for each time
    // cbfDefineFileNow has been set to 'create' and the corresponding entry here
    // has not been explicitly deleted by the application or bumped to make room
    // for a new entry.  Deleting an entry with cbfStatusFileState 'running'
    // aborts the file creation attempt.  It is implementation and file-system
    // specific whether deleting the entry also deletes the file. The type is
    // slice of CISCOBULKFILEMIB_CbfStatusFileTable_CbfStatusFileEntry.
    CbfStatusFileEntry []*CISCOBULKFILEMIB_CbfStatusFileTable_CbfStatusFileEntry
}

func (cbfStatusFileTable *CISCOBULKFILEMIB_CbfStatusFileTable) GetEntityData() *types.CommonEntityData {
    cbfStatusFileTable.EntityData.YFilter = cbfStatusFileTable.YFilter
    cbfStatusFileTable.EntityData.YangName = "cbfStatusFileTable"
    cbfStatusFileTable.EntityData.BundleName = "cisco_ios_xe"
    cbfStatusFileTable.EntityData.ParentYangName = "CISCO-BULK-FILE-MIB"
    cbfStatusFileTable.EntityData.SegmentPath = "cbfStatusFileTable"
    cbfStatusFileTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cbfStatusFileTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cbfStatusFileTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cbfStatusFileTable.EntityData.Children = types.NewOrderedMap()
    cbfStatusFileTable.EntityData.Children.Append("cbfStatusFileEntry", types.YChild{"CbfStatusFileEntry", nil})
    for i := range cbfStatusFileTable.CbfStatusFileEntry {
        cbfStatusFileTable.EntityData.Children.Append(types.GetSegmentPath(cbfStatusFileTable.CbfStatusFileEntry[i]), types.YChild{"CbfStatusFileEntry", cbfStatusFileTable.CbfStatusFileEntry[i]})
    }
    cbfStatusFileTable.EntityData.Leafs = types.NewOrderedMap()

    cbfStatusFileTable.EntityData.YListKeys = []string {}

    return &(cbfStatusFileTable.EntityData)
}

// CISCOBULKFILEMIB_CbfStatusFileTable_CbfStatusFileEntry
// Status for a particular file.
// 
// An entry exists in this table for each time cbfDefineFileNow
// has been set to 'create' and the corresponding entry here
// has not been explicitly deleted by the application or bumped
// to make room for a new entry.
// 
// Deleting an entry with cbfStatusFileState 'running' aborts
// the file creation attempt.
// 
// It is implementation and file-system specific whether deleting
// the entry also deletes the file.
type CISCOBULKFILEMIB_CbfStatusFileTable_CbfStatusFileEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // cisco_bulk_file_mib.CISCOBULKFILEMIB_CbfDefineFileTable_CbfDefineFileEntry_CbfDefineFileIndex
    CbfDefineFileIndex interface{}

    // This attribute is a key. An arbitrary integer to uniquely identify this
    // file.  The numeric order of the entries implies the creation order of the
    // files. The type is interface{} with range: 1..4294967295.
    CbfStatusFileIndex interface{}

    // The file state:  running    data is being written to the file ready     
    // the file is ready to be read emptied    an ephemeral file was successfully
    // consumed noSpace    no data due to insufficient file space badName    no
    // data due to a name or path problem writeErr   no data due to fatal file
    // write error noMem      no data due to insufficient dynamic memory buffErr  
    // implementation buffer too small aborted    short terminated by operator
    // command  Only the 'ready' state implies that the file is available for
    // transfer.  The disposition of files after an error is implementation and
    // file-syste specific. The type is CbfStatusFileState.
    CbfStatusFileState interface{}

    // The value of sysUpTime when the creation attempt completed. A value of 0
    // indicates not complete.  For ephemeral files this is the time when
    // cbfStatusFileState goes to 'emptied'.  For others this is the time when the
    // state leaves 'running'. The type is interface{} with range: 0..4294967295.
    CbfStatusFileCompletionTime interface{}

    // The control that allows deletion of entries. For detailed rules see the
    // DESCRIPTION for cbfStatusFileEntry.  This object may not be set to any
    // value other than 'destroy'. The type is RowStatus.
    CbfStatusFileEntryStatus interface{}
}

func (cbfStatusFileEntry *CISCOBULKFILEMIB_CbfStatusFileTable_CbfStatusFileEntry) GetEntityData() *types.CommonEntityData {
    cbfStatusFileEntry.EntityData.YFilter = cbfStatusFileEntry.YFilter
    cbfStatusFileEntry.EntityData.YangName = "cbfStatusFileEntry"
    cbfStatusFileEntry.EntityData.BundleName = "cisco_ios_xe"
    cbfStatusFileEntry.EntityData.ParentYangName = "cbfStatusFileTable"
    cbfStatusFileEntry.EntityData.SegmentPath = "cbfStatusFileEntry" + types.AddKeyToken(cbfStatusFileEntry.CbfDefineFileIndex, "cbfDefineFileIndex") + types.AddKeyToken(cbfStatusFileEntry.CbfStatusFileIndex, "cbfStatusFileIndex")
    cbfStatusFileEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cbfStatusFileEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cbfStatusFileEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cbfStatusFileEntry.EntityData.Children = types.NewOrderedMap()
    cbfStatusFileEntry.EntityData.Leafs = types.NewOrderedMap()
    cbfStatusFileEntry.EntityData.Leafs.Append("cbfDefineFileIndex", types.YLeaf{"CbfDefineFileIndex", cbfStatusFileEntry.CbfDefineFileIndex})
    cbfStatusFileEntry.EntityData.Leafs.Append("cbfStatusFileIndex", types.YLeaf{"CbfStatusFileIndex", cbfStatusFileEntry.CbfStatusFileIndex})
    cbfStatusFileEntry.EntityData.Leafs.Append("cbfStatusFileState", types.YLeaf{"CbfStatusFileState", cbfStatusFileEntry.CbfStatusFileState})
    cbfStatusFileEntry.EntityData.Leafs.Append("cbfStatusFileCompletionTime", types.YLeaf{"CbfStatusFileCompletionTime", cbfStatusFileEntry.CbfStatusFileCompletionTime})
    cbfStatusFileEntry.EntityData.Leafs.Append("cbfStatusFileEntryStatus", types.YLeaf{"CbfStatusFileEntryStatus", cbfStatusFileEntry.CbfStatusFileEntryStatus})

    cbfStatusFileEntry.EntityData.YListKeys = []string {"CbfDefineFileIndex", "CbfStatusFileIndex"}

    return &(cbfStatusFileEntry.EntityData)
}

// CISCOBULKFILEMIB_CbfStatusFileTable_CbfStatusFileEntry_CbfStatusFileState represents and file-syste specific.
type CISCOBULKFILEMIB_CbfStatusFileTable_CbfStatusFileEntry_CbfStatusFileState string

const (
    CISCOBULKFILEMIB_CbfStatusFileTable_CbfStatusFileEntry_CbfStatusFileState_running CISCOBULKFILEMIB_CbfStatusFileTable_CbfStatusFileEntry_CbfStatusFileState = "running"

    CISCOBULKFILEMIB_CbfStatusFileTable_CbfStatusFileEntry_CbfStatusFileState_ready CISCOBULKFILEMIB_CbfStatusFileTable_CbfStatusFileEntry_CbfStatusFileState = "ready"

    CISCOBULKFILEMIB_CbfStatusFileTable_CbfStatusFileEntry_CbfStatusFileState_emptied CISCOBULKFILEMIB_CbfStatusFileTable_CbfStatusFileEntry_CbfStatusFileState = "emptied"

    CISCOBULKFILEMIB_CbfStatusFileTable_CbfStatusFileEntry_CbfStatusFileState_noSpace CISCOBULKFILEMIB_CbfStatusFileTable_CbfStatusFileEntry_CbfStatusFileState = "noSpace"

    CISCOBULKFILEMIB_CbfStatusFileTable_CbfStatusFileEntry_CbfStatusFileState_badName CISCOBULKFILEMIB_CbfStatusFileTable_CbfStatusFileEntry_CbfStatusFileState = "badName"

    CISCOBULKFILEMIB_CbfStatusFileTable_CbfStatusFileEntry_CbfStatusFileState_writeErr CISCOBULKFILEMIB_CbfStatusFileTable_CbfStatusFileEntry_CbfStatusFileState = "writeErr"

    CISCOBULKFILEMIB_CbfStatusFileTable_CbfStatusFileEntry_CbfStatusFileState_noMem CISCOBULKFILEMIB_CbfStatusFileTable_CbfStatusFileEntry_CbfStatusFileState = "noMem"

    CISCOBULKFILEMIB_CbfStatusFileTable_CbfStatusFileEntry_CbfStatusFileState_buffErr CISCOBULKFILEMIB_CbfStatusFileTable_CbfStatusFileEntry_CbfStatusFileState = "buffErr"

    CISCOBULKFILEMIB_CbfStatusFileTable_CbfStatusFileEntry_CbfStatusFileState_aborted CISCOBULKFILEMIB_CbfStatusFileTable_CbfStatusFileEntry_CbfStatusFileState = "aborted"
)

