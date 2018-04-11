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

    
    Cbfdefine CISCOBULKFILEMIB_Cbfdefine

    
    Cbfstatus CISCOBULKFILEMIB_Cbfstatus

    // A table of bulk file definition and creation controls.
    Cbfdefinefiletable CISCOBULKFILEMIB_Cbfdefinefiletable

    // A table of objects to go in bulk files.
    Cbfdefineobjecttable CISCOBULKFILEMIB_Cbfdefineobjecttable

    // A table of bulk file status.
    Cbfstatusfiletable CISCOBULKFILEMIB_Cbfstatusfiletable
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

    cISCOBULKFILEMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOBULKFILEMIB.EntityData.Children["cbfDefine"] = types.YChild{"Cbfdefine", &cISCOBULKFILEMIB.Cbfdefine}
    cISCOBULKFILEMIB.EntityData.Children["cbfStatus"] = types.YChild{"Cbfstatus", &cISCOBULKFILEMIB.Cbfstatus}
    cISCOBULKFILEMIB.EntityData.Children["cbfDefineFileTable"] = types.YChild{"Cbfdefinefiletable", &cISCOBULKFILEMIB.Cbfdefinefiletable}
    cISCOBULKFILEMIB.EntityData.Children["cbfDefineObjectTable"] = types.YChild{"Cbfdefineobjecttable", &cISCOBULKFILEMIB.Cbfdefineobjecttable}
    cISCOBULKFILEMIB.EntityData.Children["cbfStatusFileTable"] = types.YChild{"Cbfstatusfiletable", &cISCOBULKFILEMIB.Cbfstatusfiletable}
    cISCOBULKFILEMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOBULKFILEMIB.EntityData)
}

// CISCOBULKFILEMIB_Cbfdefine
type CISCOBULKFILEMIB_Cbfdefine struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The maximum number of file definitions this system can hold in
    // cbfDefineFileTable.  A value of 0 indicates no configured limit.  This
    // object may be read-only on some systems.  Changing this number does not
    // disturb existing entries. The type is interface{} with range:
    // 0..4294967295.
    Cbfdefinemaxfiles interface{}

    // The current number of file definitions in cbfDefineFileTable. The type is
    // interface{} with range: 0..4294967295.
    Cbfdefinefiles interface{}

    // The maximum value of cbfDefineFiles since system  initialization. The type
    // is interface{} with range: 0..4294967295.
    Cbfdefinehighfiles interface{}

    // The number of attempts to create a file definition that failed due to
    // exceeding cbfDefineMaxFiles. The type is interface{} with range:
    // 0..4294967295.
    Cbfdefinefilesrefused interface{}

    // The maximum total number of object selections to go with file definitions
    // this system, that is, the total number of objects this system can hold in
    // cbfDefineObjectTable.  A value of 0 indicates no configured limit.  This
    // object may be read-only on some systems.  Changing this number does not
    // disturb existing entries. The type is interface{} with range:
    // 0..4294967295.
    Cbfdefinemaxobjects interface{}

    // The current number of object selections in  cbfDefineObjectTable. The type
    // is interface{} with range: 0..4294967295.
    Cbfdefineobjects interface{}

    // The maximum value of cbfDefineObjects since system  initialization. The
    // type is interface{} with range: 0..4294967295.
    Cbfdefinehighobjects interface{}

    // The number of attempts to create an object selection that failed due to
    // exceeding cbfDefineMaxObjects. The type is interface{} with range:
    // 0..4294967295.
    Cbfdefineobjectsrefused interface{}
}

func (cbfdefine *CISCOBULKFILEMIB_Cbfdefine) GetEntityData() *types.CommonEntityData {
    cbfdefine.EntityData.YFilter = cbfdefine.YFilter
    cbfdefine.EntityData.YangName = "cbfDefine"
    cbfdefine.EntityData.BundleName = "cisco_ios_xe"
    cbfdefine.EntityData.ParentYangName = "CISCO-BULK-FILE-MIB"
    cbfdefine.EntityData.SegmentPath = "cbfDefine"
    cbfdefine.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cbfdefine.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cbfdefine.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cbfdefine.EntityData.Children = make(map[string]types.YChild)
    cbfdefine.EntityData.Leafs = make(map[string]types.YLeaf)
    cbfdefine.EntityData.Leafs["cbfDefineMaxFiles"] = types.YLeaf{"Cbfdefinemaxfiles", cbfdefine.Cbfdefinemaxfiles}
    cbfdefine.EntityData.Leafs["cbfDefineFiles"] = types.YLeaf{"Cbfdefinefiles", cbfdefine.Cbfdefinefiles}
    cbfdefine.EntityData.Leafs["cbfDefineHighFiles"] = types.YLeaf{"Cbfdefinehighfiles", cbfdefine.Cbfdefinehighfiles}
    cbfdefine.EntityData.Leafs["cbfDefineFilesRefused"] = types.YLeaf{"Cbfdefinefilesrefused", cbfdefine.Cbfdefinefilesrefused}
    cbfdefine.EntityData.Leafs["cbfDefineMaxObjects"] = types.YLeaf{"Cbfdefinemaxobjects", cbfdefine.Cbfdefinemaxobjects}
    cbfdefine.EntityData.Leafs["cbfDefineObjects"] = types.YLeaf{"Cbfdefineobjects", cbfdefine.Cbfdefineobjects}
    cbfdefine.EntityData.Leafs["cbfDefineHighObjects"] = types.YLeaf{"Cbfdefinehighobjects", cbfdefine.Cbfdefinehighobjects}
    cbfdefine.EntityData.Leafs["cbfDefineObjectsRefused"] = types.YLeaf{"Cbfdefineobjectsrefused", cbfdefine.Cbfdefineobjectsrefused}
    return &(cbfdefine.EntityData)
}

// CISCOBULKFILEMIB_Cbfstatus
type CISCOBULKFILEMIB_Cbfstatus struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The maximum number of file statuses this system can hold in
    // cbfStatusFileTable.  A value of 0 indicates no configured limit.  This
    // object may be read-only on some systems.  Changing this number deletes the
    // oldest finished entries until the new limit is satisfied. The type is
    // interface{} with range: 0..4294967295.
    Cbfstatusmaxfiles interface{}

    // The current number of file statuses in cbfStatusFileTable. The type is
    // interface{} with range: 0..4294967295.
    Cbfstatusfiles interface{}

    // The maximum value of cbfStatusFiles since system  initialization. The type
    // is interface{} with range: 0..4294967295.
    Cbfstatushighfiles interface{}

    // The number times the oldest entry was deleted due to exceeding
    // cbfStatusMaxFiles. The type is interface{} with range: 0..4294967295.
    Cbfstatusfilesbumped interface{}
}

func (cbfstatus *CISCOBULKFILEMIB_Cbfstatus) GetEntityData() *types.CommonEntityData {
    cbfstatus.EntityData.YFilter = cbfstatus.YFilter
    cbfstatus.EntityData.YangName = "cbfStatus"
    cbfstatus.EntityData.BundleName = "cisco_ios_xe"
    cbfstatus.EntityData.ParentYangName = "CISCO-BULK-FILE-MIB"
    cbfstatus.EntityData.SegmentPath = "cbfStatus"
    cbfstatus.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cbfstatus.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cbfstatus.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cbfstatus.EntityData.Children = make(map[string]types.YChild)
    cbfstatus.EntityData.Leafs = make(map[string]types.YLeaf)
    cbfstatus.EntityData.Leafs["cbfStatusMaxFiles"] = types.YLeaf{"Cbfstatusmaxfiles", cbfstatus.Cbfstatusmaxfiles}
    cbfstatus.EntityData.Leafs["cbfStatusFiles"] = types.YLeaf{"Cbfstatusfiles", cbfstatus.Cbfstatusfiles}
    cbfstatus.EntityData.Leafs["cbfStatusHighFiles"] = types.YLeaf{"Cbfstatushighfiles", cbfstatus.Cbfstatushighfiles}
    cbfstatus.EntityData.Leafs["cbfStatusFilesBumped"] = types.YLeaf{"Cbfstatusfilesbumped", cbfstatus.Cbfstatusfilesbumped}
    return &(cbfstatus.EntityData)
}

// CISCOBULKFILEMIB_Cbfdefinefiletable
// A table of bulk file definition and creation controls.
type CISCOBULKFILEMIB_Cbfdefinefiletable struct {
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
    // CISCOBULKFILEMIB_Cbfdefinefiletable_Cbfdefinefileentry.
    Cbfdefinefileentry []CISCOBULKFILEMIB_Cbfdefinefiletable_Cbfdefinefileentry
}

func (cbfdefinefiletable *CISCOBULKFILEMIB_Cbfdefinefiletable) GetEntityData() *types.CommonEntityData {
    cbfdefinefiletable.EntityData.YFilter = cbfdefinefiletable.YFilter
    cbfdefinefiletable.EntityData.YangName = "cbfDefineFileTable"
    cbfdefinefiletable.EntityData.BundleName = "cisco_ios_xe"
    cbfdefinefiletable.EntityData.ParentYangName = "CISCO-BULK-FILE-MIB"
    cbfdefinefiletable.EntityData.SegmentPath = "cbfDefineFileTable"
    cbfdefinefiletable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cbfdefinefiletable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cbfdefinefiletable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cbfdefinefiletable.EntityData.Children = make(map[string]types.YChild)
    cbfdefinefiletable.EntityData.Children["cbfDefineFileEntry"] = types.YChild{"Cbfdefinefileentry", nil}
    for i := range cbfdefinefiletable.Cbfdefinefileentry {
        cbfdefinefiletable.EntityData.Children[types.GetSegmentPath(&cbfdefinefiletable.Cbfdefinefileentry[i])] = types.YChild{"Cbfdefinefileentry", &cbfdefinefiletable.Cbfdefinefileentry[i]}
    }
    cbfdefinefiletable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cbfdefinefiletable.EntityData)
}

// CISCOBULKFILEMIB_Cbfdefinefiletable_Cbfdefinefileentry
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
type CISCOBULKFILEMIB_Cbfdefinefiletable_Cbfdefinefileentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An arbitrary integer to uniquely identify this
    // entry.  To create an entry a management application should pick a random
    // number. The type is interface{} with range: 1..4294967295.
    Cbfdefinefileindex interface{}

    // The file name which is to be created.  Explicit device or path choices in
    // the value of this object override cbfDefineFileStorage. The type is string
    // with length: 0..255.
    Cbfdefinefilename interface{}

    // The type of file storage to use:  ephemeral        data exists in small
    // amounts until read volatile        data exists in volatile memory permanent
    // data survives reboot  An ephemeral file is suitable to be read only one
    // time.  Note that this value is taken as advisory and may be overridden by
    // explicit device or path choices in cbfDefineFile.  A given system may
    // support any or all of these. The type is Cbfdefinefilestorage.
    Cbfdefinefilestorage interface{}

    // The format of the data in the file:  StandardBER        standard SNMP ASN.1
    // BER bulkBinary        a binary format specified with this MIB bulkASCII    
    // a human-readable form of bulkBinary variantBERWithCksum ASN.1 BER encoding
    // with checksum variantBinWithCksum a binary format with checksum      A
    // given system may support any or all of these. The type is
    // Cbfdefinefileformat.
    Cbfdefinefileformat interface{}

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
    // 'ready'. The type is Cbfdefinefilenow.
    Cbfdefinefilenow interface{}

    // The control that allows creation, modification, and deletion of entries. 
    // For detailed rules see the DESCRIPTION for cbfDefineFileEntry. The type is
    // RowStatus.
    Cbfdefinefileentrystatus interface{}

    // This controls the cbfDefineFileCompletion notification.  If true,
    // cbfDefineFileCompletion notification will be generated. It is the
    // responsibility of the  management entity to ensure that the SNMP
    // administrative  model is configured in such a way as to allow the 
    // notification to be delivered. The type is bool.
    Cbfdefinefilenotifyoncompletion interface{}
}

func (cbfdefinefileentry *CISCOBULKFILEMIB_Cbfdefinefiletable_Cbfdefinefileentry) GetEntityData() *types.CommonEntityData {
    cbfdefinefileentry.EntityData.YFilter = cbfdefinefileentry.YFilter
    cbfdefinefileentry.EntityData.YangName = "cbfDefineFileEntry"
    cbfdefinefileentry.EntityData.BundleName = "cisco_ios_xe"
    cbfdefinefileentry.EntityData.ParentYangName = "cbfDefineFileTable"
    cbfdefinefileentry.EntityData.SegmentPath = "cbfDefineFileEntry" + "[cbfDefineFileIndex='" + fmt.Sprintf("%v", cbfdefinefileentry.Cbfdefinefileindex) + "']"
    cbfdefinefileentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cbfdefinefileentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cbfdefinefileentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cbfdefinefileentry.EntityData.Children = make(map[string]types.YChild)
    cbfdefinefileentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cbfdefinefileentry.EntityData.Leafs["cbfDefineFileIndex"] = types.YLeaf{"Cbfdefinefileindex", cbfdefinefileentry.Cbfdefinefileindex}
    cbfdefinefileentry.EntityData.Leafs["cbfDefineFileName"] = types.YLeaf{"Cbfdefinefilename", cbfdefinefileentry.Cbfdefinefilename}
    cbfdefinefileentry.EntityData.Leafs["cbfDefineFileStorage"] = types.YLeaf{"Cbfdefinefilestorage", cbfdefinefileentry.Cbfdefinefilestorage}
    cbfdefinefileentry.EntityData.Leafs["cbfDefineFileFormat"] = types.YLeaf{"Cbfdefinefileformat", cbfdefinefileentry.Cbfdefinefileformat}
    cbfdefinefileentry.EntityData.Leafs["cbfDefineFileNow"] = types.YLeaf{"Cbfdefinefilenow", cbfdefinefileentry.Cbfdefinefilenow}
    cbfdefinefileentry.EntityData.Leafs["cbfDefineFileEntryStatus"] = types.YLeaf{"Cbfdefinefileentrystatus", cbfdefinefileentry.Cbfdefinefileentrystatus}
    cbfdefinefileentry.EntityData.Leafs["cbfDefineFileNotifyOnCompletion"] = types.YLeaf{"Cbfdefinefilenotifyoncompletion", cbfdefinefileentry.Cbfdefinefilenotifyoncompletion}
    return &(cbfdefinefileentry.EntityData)
}

// CISCOBULKFILEMIB_Cbfdefinefiletable_Cbfdefinefileentry_Cbfdefinefileformat represents     A given system may support any or all of these.
type CISCOBULKFILEMIB_Cbfdefinefiletable_Cbfdefinefileentry_Cbfdefinefileformat string

const (
    CISCOBULKFILEMIB_Cbfdefinefiletable_Cbfdefinefileentry_Cbfdefinefileformat_standardBER CISCOBULKFILEMIB_Cbfdefinefiletable_Cbfdefinefileentry_Cbfdefinefileformat = "standardBER"

    CISCOBULKFILEMIB_Cbfdefinefiletable_Cbfdefinefileentry_Cbfdefinefileformat_bulkBinary CISCOBULKFILEMIB_Cbfdefinefiletable_Cbfdefinefileentry_Cbfdefinefileformat = "bulkBinary"

    CISCOBULKFILEMIB_Cbfdefinefiletable_Cbfdefinefileentry_Cbfdefinefileformat_bulkASCII CISCOBULKFILEMIB_Cbfdefinefiletable_Cbfdefinefileentry_Cbfdefinefileformat = "bulkASCII"

    CISCOBULKFILEMIB_Cbfdefinefiletable_Cbfdefinefileentry_Cbfdefinefileformat_variantBERWithCksum CISCOBULKFILEMIB_Cbfdefinefiletable_Cbfdefinefileentry_Cbfdefinefileformat = "variantBERWithCksum"

    CISCOBULKFILEMIB_Cbfdefinefiletable_Cbfdefinefileentry_Cbfdefinefileformat_variantBinWithCksum CISCOBULKFILEMIB_Cbfdefinefiletable_Cbfdefinefileentry_Cbfdefinefileformat = "variantBinWithCksum"
)

// CISCOBULKFILEMIB_Cbfdefinefiletable_Cbfdefinefileentry_Cbfdefinefilenow represents object automatically goes to 'ready'.
type CISCOBULKFILEMIB_Cbfdefinefiletable_Cbfdefinefileentry_Cbfdefinefilenow string

const (
    CISCOBULKFILEMIB_Cbfdefinefiletable_Cbfdefinefileentry_Cbfdefinefilenow_notActive CISCOBULKFILEMIB_Cbfdefinefiletable_Cbfdefinefileentry_Cbfdefinefilenow = "notActive"

    CISCOBULKFILEMIB_Cbfdefinefiletable_Cbfdefinefileentry_Cbfdefinefilenow_ready CISCOBULKFILEMIB_Cbfdefinefiletable_Cbfdefinefileentry_Cbfdefinefilenow = "ready"

    CISCOBULKFILEMIB_Cbfdefinefiletable_Cbfdefinefileentry_Cbfdefinefilenow_create CISCOBULKFILEMIB_Cbfdefinefiletable_Cbfdefinefileentry_Cbfdefinefilenow = "create"

    CISCOBULKFILEMIB_Cbfdefinefiletable_Cbfdefinefileentry_Cbfdefinefilenow_running CISCOBULKFILEMIB_Cbfdefinefiletable_Cbfdefinefileentry_Cbfdefinefilenow = "running"

    CISCOBULKFILEMIB_Cbfdefinefiletable_Cbfdefinefileentry_Cbfdefinefilenow_forcedCreate CISCOBULKFILEMIB_Cbfdefinefiletable_Cbfdefinefileentry_Cbfdefinefilenow = "forcedCreate"
)

// CISCOBULKFILEMIB_Cbfdefinefiletable_Cbfdefinefileentry_Cbfdefinefilestorage represents A given system may support any or all of these.
type CISCOBULKFILEMIB_Cbfdefinefiletable_Cbfdefinefileentry_Cbfdefinefilestorage string

const (
    CISCOBULKFILEMIB_Cbfdefinefiletable_Cbfdefinefileentry_Cbfdefinefilestorage_ephemeral CISCOBULKFILEMIB_Cbfdefinefiletable_Cbfdefinefileentry_Cbfdefinefilestorage = "ephemeral"

    CISCOBULKFILEMIB_Cbfdefinefiletable_Cbfdefinefileentry_Cbfdefinefilestorage_volatile CISCOBULKFILEMIB_Cbfdefinefiletable_Cbfdefinefileentry_Cbfdefinefilestorage = "volatile"

    CISCOBULKFILEMIB_Cbfdefinefiletable_Cbfdefinefileentry_Cbfdefinefilestorage_permanent CISCOBULKFILEMIB_Cbfdefinefiletable_Cbfdefinefileentry_Cbfdefinefilestorage = "permanent"
)

// CISCOBULKFILEMIB_Cbfdefineobjecttable
// A table of objects to go in bulk files.
type CISCOBULKFILEMIB_Cbfdefineobjecttable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about one object for a particular file.  An application uses
    // cbfDefineObjectEntryStatus to create entries in this table in
    // correspondence with entries in cbfDefineFileTable, which must be created
    // first.  Entries in this table may not be changed, created or deleted while
    // the corresponding value of cbfDefineFileNow is 'running'. The type is slice
    // of CISCOBULKFILEMIB_Cbfdefineobjecttable_Cbfdefineobjectentry.
    Cbfdefineobjectentry []CISCOBULKFILEMIB_Cbfdefineobjecttable_Cbfdefineobjectentry
}

func (cbfdefineobjecttable *CISCOBULKFILEMIB_Cbfdefineobjecttable) GetEntityData() *types.CommonEntityData {
    cbfdefineobjecttable.EntityData.YFilter = cbfdefineobjecttable.YFilter
    cbfdefineobjecttable.EntityData.YangName = "cbfDefineObjectTable"
    cbfdefineobjecttable.EntityData.BundleName = "cisco_ios_xe"
    cbfdefineobjecttable.EntityData.ParentYangName = "CISCO-BULK-FILE-MIB"
    cbfdefineobjecttable.EntityData.SegmentPath = "cbfDefineObjectTable"
    cbfdefineobjecttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cbfdefineobjecttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cbfdefineobjecttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cbfdefineobjecttable.EntityData.Children = make(map[string]types.YChild)
    cbfdefineobjecttable.EntityData.Children["cbfDefineObjectEntry"] = types.YChild{"Cbfdefineobjectentry", nil}
    for i := range cbfdefineobjecttable.Cbfdefineobjectentry {
        cbfdefineobjecttable.EntityData.Children[types.GetSegmentPath(&cbfdefineobjecttable.Cbfdefineobjectentry[i])] = types.YChild{"Cbfdefineobjectentry", &cbfdefineobjecttable.Cbfdefineobjectentry[i]}
    }
    cbfdefineobjecttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cbfdefineobjecttable.EntityData)
}

// CISCOBULKFILEMIB_Cbfdefineobjecttable_Cbfdefineobjectentry
// Information about one object for a particular file.
// 
// An application uses cbfDefineObjectEntryStatus to create entries
// in this table in correspondence with entries in
// cbfDefineFileTable, which must be created first.
// 
// Entries in this table may not be changed, created or deleted
// while the corresponding value of cbfDefineFileNow is 'running'.
type CISCOBULKFILEMIB_Cbfdefineobjecttable_Cbfdefineobjectentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // cisco_bulk_file_mib.CISCOBULKFILEMIB_Cbfdefinefiletable_Cbfdefinefileentry_Cbfdefinefileindex
    Cbfdefinefileindex interface{}

    // This attribute is a key. An arbitrary integer to uniquely identify this
    // entry.  The numeric order of the entries controls the order of the objects
    // in the file. The type is interface{} with range: 1..4294967295.
    Cbfdefineobjectindex interface{}

    // The meaning of each object class is given below:  object          a single
    // MIB object is retrieved.  lexicalTable    an entire table or partial table 
    // is retrieved in lexical order of rows.  leastCpuTable   an entire table is
    // retrieved with                 lowest CPU utilization.                
    // Lexical ordering of rows may not be                  maintained and is
    // dependent upon                  individual MIB implementation. The type is
    // Cbfdefineobjectclass.
    Cbfdefineobjectclass interface{}

    // The object identifier of a MIB object to be included in the file.  If
    // cbfDefineObjectClass is 'object' this must be a full OID, including all
    // instance information.  If cbfDefineObjectClass is 'lexicalTable' or
    // 'leastCpuTable' this must be the OID of the table-defining SEQUENCE OF
    // registration point. The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Cbfdefineobjectid interface{}

    // The control that allows creation, modification, and deletion of entries. 
    // For detailed rules see the DESCRIPTION for cbfDefineObjectEntry. The type
    // is RowStatus.
    Cbfdefineobjectentrystatus interface{}

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
    // pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Cbfdefineobjecttableinstance interface{}

    // If cbfDefineObjectClass is 'lexicalTable', then this object represents the
    // maximum number of entries which will be  populated in the file starting
    // from the lexicographically next instance of the OID represented by 
    // cbfDefineObjectTableInstance.   This object is irrelevent if
    // cbfDefineObjectClass is not 'lexicalTable'.  Refer to the description of
    // cbfDefineObjectTableInstance for examples and different scenarios relating
    // to this object. The type is interface{} with range: 0..4294967295.
    Cbfdefineobjectnumentries interface{}

    // This object represents the last polled instance in the table.  The value
    // represented by this object will be relevent only if the corresponding
    // cbfStatusFileState is emptied(3) for  ephemeral files or ready(2) for
    // volatile or permanent files.  A value of zeroDotZero indicates an absence
    // of last polled  object.  An NMS can use the value of this object and
    // populate the cbfDefineObjectTableInstance to retrieve a contiguous set of
    // rows in a table. The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Cbfdefineobjectlastpolledinst interface{}
}

func (cbfdefineobjectentry *CISCOBULKFILEMIB_Cbfdefineobjecttable_Cbfdefineobjectentry) GetEntityData() *types.CommonEntityData {
    cbfdefineobjectentry.EntityData.YFilter = cbfdefineobjectentry.YFilter
    cbfdefineobjectentry.EntityData.YangName = "cbfDefineObjectEntry"
    cbfdefineobjectentry.EntityData.BundleName = "cisco_ios_xe"
    cbfdefineobjectentry.EntityData.ParentYangName = "cbfDefineObjectTable"
    cbfdefineobjectentry.EntityData.SegmentPath = "cbfDefineObjectEntry" + "[cbfDefineFileIndex='" + fmt.Sprintf("%v", cbfdefineobjectentry.Cbfdefinefileindex) + "']" + "[cbfDefineObjectIndex='" + fmt.Sprintf("%v", cbfdefineobjectentry.Cbfdefineobjectindex) + "']"
    cbfdefineobjectentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cbfdefineobjectentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cbfdefineobjectentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cbfdefineobjectentry.EntityData.Children = make(map[string]types.YChild)
    cbfdefineobjectentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cbfdefineobjectentry.EntityData.Leafs["cbfDefineFileIndex"] = types.YLeaf{"Cbfdefinefileindex", cbfdefineobjectentry.Cbfdefinefileindex}
    cbfdefineobjectentry.EntityData.Leafs["cbfDefineObjectIndex"] = types.YLeaf{"Cbfdefineobjectindex", cbfdefineobjectentry.Cbfdefineobjectindex}
    cbfdefineobjectentry.EntityData.Leafs["cbfDefineObjectClass"] = types.YLeaf{"Cbfdefineobjectclass", cbfdefineobjectentry.Cbfdefineobjectclass}
    cbfdefineobjectentry.EntityData.Leafs["cbfDefineObjectID"] = types.YLeaf{"Cbfdefineobjectid", cbfdefineobjectentry.Cbfdefineobjectid}
    cbfdefineobjectentry.EntityData.Leafs["cbfDefineObjectEntryStatus"] = types.YLeaf{"Cbfdefineobjectentrystatus", cbfdefineobjectentry.Cbfdefineobjectentrystatus}
    cbfdefineobjectentry.EntityData.Leafs["cbfDefineObjectTableInstance"] = types.YLeaf{"Cbfdefineobjecttableinstance", cbfdefineobjectentry.Cbfdefineobjecttableinstance}
    cbfdefineobjectentry.EntityData.Leafs["cbfDefineObjectNumEntries"] = types.YLeaf{"Cbfdefineobjectnumentries", cbfdefineobjectentry.Cbfdefineobjectnumentries}
    cbfdefineobjectentry.EntityData.Leafs["cbfDefineObjectLastPolledInst"] = types.YLeaf{"Cbfdefineobjectlastpolledinst", cbfdefineobjectentry.Cbfdefineobjectlastpolledinst}
    return &(cbfdefineobjectentry.EntityData)
}

// CISCOBULKFILEMIB_Cbfdefineobjecttable_Cbfdefineobjectentry_Cbfdefineobjectclass represents                 individual MIB implementation.
type CISCOBULKFILEMIB_Cbfdefineobjecttable_Cbfdefineobjectentry_Cbfdefineobjectclass string

const (
    CISCOBULKFILEMIB_Cbfdefineobjecttable_Cbfdefineobjectentry_Cbfdefineobjectclass_object CISCOBULKFILEMIB_Cbfdefineobjecttable_Cbfdefineobjectentry_Cbfdefineobjectclass = "object"

    CISCOBULKFILEMIB_Cbfdefineobjecttable_Cbfdefineobjectentry_Cbfdefineobjectclass_lexicalTable CISCOBULKFILEMIB_Cbfdefineobjecttable_Cbfdefineobjectentry_Cbfdefineobjectclass = "lexicalTable"

    CISCOBULKFILEMIB_Cbfdefineobjecttable_Cbfdefineobjectentry_Cbfdefineobjectclass_leastCpuTable CISCOBULKFILEMIB_Cbfdefineobjecttable_Cbfdefineobjectentry_Cbfdefineobjectclass = "leastCpuTable"
)

// CISCOBULKFILEMIB_Cbfstatusfiletable
// A table of bulk file status.
type CISCOBULKFILEMIB_Cbfstatusfiletable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Status for a particular file.  An entry exists in this table for each time
    // cbfDefineFileNow has been set to 'create' and the corresponding entry here
    // has not been explicitly deleted by the application or bumped to make room
    // for a new entry.  Deleting an entry with cbfStatusFileState 'running'
    // aborts the file creation attempt.  It is implementation and file-system
    // specific whether deleting the entry also deletes the file. The type is
    // slice of CISCOBULKFILEMIB_Cbfstatusfiletable_Cbfstatusfileentry.
    Cbfstatusfileentry []CISCOBULKFILEMIB_Cbfstatusfiletable_Cbfstatusfileentry
}

func (cbfstatusfiletable *CISCOBULKFILEMIB_Cbfstatusfiletable) GetEntityData() *types.CommonEntityData {
    cbfstatusfiletable.EntityData.YFilter = cbfstatusfiletable.YFilter
    cbfstatusfiletable.EntityData.YangName = "cbfStatusFileTable"
    cbfstatusfiletable.EntityData.BundleName = "cisco_ios_xe"
    cbfstatusfiletable.EntityData.ParentYangName = "CISCO-BULK-FILE-MIB"
    cbfstatusfiletable.EntityData.SegmentPath = "cbfStatusFileTable"
    cbfstatusfiletable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cbfstatusfiletable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cbfstatusfiletable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cbfstatusfiletable.EntityData.Children = make(map[string]types.YChild)
    cbfstatusfiletable.EntityData.Children["cbfStatusFileEntry"] = types.YChild{"Cbfstatusfileentry", nil}
    for i := range cbfstatusfiletable.Cbfstatusfileentry {
        cbfstatusfiletable.EntityData.Children[types.GetSegmentPath(&cbfstatusfiletable.Cbfstatusfileentry[i])] = types.YChild{"Cbfstatusfileentry", &cbfstatusfiletable.Cbfstatusfileentry[i]}
    }
    cbfstatusfiletable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cbfstatusfiletable.EntityData)
}

// CISCOBULKFILEMIB_Cbfstatusfiletable_Cbfstatusfileentry
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
type CISCOBULKFILEMIB_Cbfstatusfiletable_Cbfstatusfileentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // cisco_bulk_file_mib.CISCOBULKFILEMIB_Cbfdefinefiletable_Cbfdefinefileentry_Cbfdefinefileindex
    Cbfdefinefileindex interface{}

    // This attribute is a key. An arbitrary integer to uniquely identify this
    // file.  The numeric order of the entries implies the creation order of the
    // files. The type is interface{} with range: 1..4294967295.
    Cbfstatusfileindex interface{}

    // The file state:  running    data is being written to the file ready     
    // the file is ready to be read emptied    an ephemeral file was successfully
    // consumed noSpace    no data due to insufficient file space badName    no
    // data due to a name or path problem writeErr   no data due to fatal file
    // write error noMem      no data due to insufficient dynamic memory buffErr  
    // implementation buffer too small aborted    short terminated by operator
    // command  Only the 'ready' state implies that the file is available for
    // transfer.  The disposition of files after an error is implementation and
    // file-syste specific. The type is Cbfstatusfilestate.
    Cbfstatusfilestate interface{}

    // The value of sysUpTime when the creation attempt completed. A value of 0
    // indicates not complete.  For ephemeral files this is the time when
    // cbfStatusFileState goes to 'emptied'.  For others this is the time when the
    // state leaves 'running'. The type is interface{} with range: 0..4294967295.
    Cbfstatusfilecompletiontime interface{}

    // The control that allows deletion of entries. For detailed rules see the
    // DESCRIPTION for cbfStatusFileEntry.  This object may not be set to any
    // value other than 'destroy'. The type is RowStatus.
    Cbfstatusfileentrystatus interface{}
}

func (cbfstatusfileentry *CISCOBULKFILEMIB_Cbfstatusfiletable_Cbfstatusfileentry) GetEntityData() *types.CommonEntityData {
    cbfstatusfileentry.EntityData.YFilter = cbfstatusfileentry.YFilter
    cbfstatusfileentry.EntityData.YangName = "cbfStatusFileEntry"
    cbfstatusfileentry.EntityData.BundleName = "cisco_ios_xe"
    cbfstatusfileentry.EntityData.ParentYangName = "cbfStatusFileTable"
    cbfstatusfileentry.EntityData.SegmentPath = "cbfStatusFileEntry" + "[cbfDefineFileIndex='" + fmt.Sprintf("%v", cbfstatusfileentry.Cbfdefinefileindex) + "']" + "[cbfStatusFileIndex='" + fmt.Sprintf("%v", cbfstatusfileentry.Cbfstatusfileindex) + "']"
    cbfstatusfileentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cbfstatusfileentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cbfstatusfileentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cbfstatusfileentry.EntityData.Children = make(map[string]types.YChild)
    cbfstatusfileentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cbfstatusfileentry.EntityData.Leafs["cbfDefineFileIndex"] = types.YLeaf{"Cbfdefinefileindex", cbfstatusfileentry.Cbfdefinefileindex}
    cbfstatusfileentry.EntityData.Leafs["cbfStatusFileIndex"] = types.YLeaf{"Cbfstatusfileindex", cbfstatusfileentry.Cbfstatusfileindex}
    cbfstatusfileentry.EntityData.Leafs["cbfStatusFileState"] = types.YLeaf{"Cbfstatusfilestate", cbfstatusfileentry.Cbfstatusfilestate}
    cbfstatusfileentry.EntityData.Leafs["cbfStatusFileCompletionTime"] = types.YLeaf{"Cbfstatusfilecompletiontime", cbfstatusfileentry.Cbfstatusfilecompletiontime}
    cbfstatusfileentry.EntityData.Leafs["cbfStatusFileEntryStatus"] = types.YLeaf{"Cbfstatusfileentrystatus", cbfstatusfileentry.Cbfstatusfileentrystatus}
    return &(cbfstatusfileentry.EntityData)
}

// CISCOBULKFILEMIB_Cbfstatusfiletable_Cbfstatusfileentry_Cbfstatusfilestate represents and file-syste specific.
type CISCOBULKFILEMIB_Cbfstatusfiletable_Cbfstatusfileentry_Cbfstatusfilestate string

const (
    CISCOBULKFILEMIB_Cbfstatusfiletable_Cbfstatusfileentry_Cbfstatusfilestate_running CISCOBULKFILEMIB_Cbfstatusfiletable_Cbfstatusfileentry_Cbfstatusfilestate = "running"

    CISCOBULKFILEMIB_Cbfstatusfiletable_Cbfstatusfileentry_Cbfstatusfilestate_ready CISCOBULKFILEMIB_Cbfstatusfiletable_Cbfstatusfileentry_Cbfstatusfilestate = "ready"

    CISCOBULKFILEMIB_Cbfstatusfiletable_Cbfstatusfileentry_Cbfstatusfilestate_emptied CISCOBULKFILEMIB_Cbfstatusfiletable_Cbfstatusfileentry_Cbfstatusfilestate = "emptied"

    CISCOBULKFILEMIB_Cbfstatusfiletable_Cbfstatusfileentry_Cbfstatusfilestate_noSpace CISCOBULKFILEMIB_Cbfstatusfiletable_Cbfstatusfileentry_Cbfstatusfilestate = "noSpace"

    CISCOBULKFILEMIB_Cbfstatusfiletable_Cbfstatusfileentry_Cbfstatusfilestate_badName CISCOBULKFILEMIB_Cbfstatusfiletable_Cbfstatusfileentry_Cbfstatusfilestate = "badName"

    CISCOBULKFILEMIB_Cbfstatusfiletable_Cbfstatusfileentry_Cbfstatusfilestate_writeErr CISCOBULKFILEMIB_Cbfstatusfiletable_Cbfstatusfileentry_Cbfstatusfilestate = "writeErr"

    CISCOBULKFILEMIB_Cbfstatusfiletable_Cbfstatusfileentry_Cbfstatusfilestate_noMem CISCOBULKFILEMIB_Cbfstatusfiletable_Cbfstatusfileentry_Cbfstatusfilestate = "noMem"

    CISCOBULKFILEMIB_Cbfstatusfiletable_Cbfstatusfileentry_Cbfstatusfilestate_buffErr CISCOBULKFILEMIB_Cbfstatusfiletable_Cbfstatusfileentry_Cbfstatusfilestate = "buffErr"

    CISCOBULKFILEMIB_Cbfstatusfiletable_Cbfstatusfileentry_Cbfstatusfilestate_aborted CISCOBULKFILEMIB_Cbfstatusfiletable_Cbfstatusfileentry_Cbfstatusfilestate = "aborted"
)

