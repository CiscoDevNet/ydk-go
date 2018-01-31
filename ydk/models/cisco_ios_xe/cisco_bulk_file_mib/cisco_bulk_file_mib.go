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
    parent types.Entity
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

func (cISCOBULKFILEMIB *CISCOBULKFILEMIB) GetFilter() yfilter.YFilter { return cISCOBULKFILEMIB.YFilter }

func (cISCOBULKFILEMIB *CISCOBULKFILEMIB) SetFilter(yf yfilter.YFilter) { cISCOBULKFILEMIB.YFilter = yf }

func (cISCOBULKFILEMIB *CISCOBULKFILEMIB) GetGoName(yname string) string {
    if yname == "cbfDefine" { return "Cbfdefine" }
    if yname == "cbfStatus" { return "Cbfstatus" }
    if yname == "cbfDefineFileTable" { return "Cbfdefinefiletable" }
    if yname == "cbfDefineObjectTable" { return "Cbfdefineobjecttable" }
    if yname == "cbfStatusFileTable" { return "Cbfstatusfiletable" }
    return ""
}

func (cISCOBULKFILEMIB *CISCOBULKFILEMIB) GetSegmentPath() string {
    return "CISCO-BULK-FILE-MIB:CISCO-BULK-FILE-MIB"
}

func (cISCOBULKFILEMIB *CISCOBULKFILEMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cbfDefine" {
        return &cISCOBULKFILEMIB.Cbfdefine
    }
    if childYangName == "cbfStatus" {
        return &cISCOBULKFILEMIB.Cbfstatus
    }
    if childYangName == "cbfDefineFileTable" {
        return &cISCOBULKFILEMIB.Cbfdefinefiletable
    }
    if childYangName == "cbfDefineObjectTable" {
        return &cISCOBULKFILEMIB.Cbfdefineobjecttable
    }
    if childYangName == "cbfStatusFileTable" {
        return &cISCOBULKFILEMIB.Cbfstatusfiletable
    }
    return nil
}

func (cISCOBULKFILEMIB *CISCOBULKFILEMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cbfDefine"] = &cISCOBULKFILEMIB.Cbfdefine
    children["cbfStatus"] = &cISCOBULKFILEMIB.Cbfstatus
    children["cbfDefineFileTable"] = &cISCOBULKFILEMIB.Cbfdefinefiletable
    children["cbfDefineObjectTable"] = &cISCOBULKFILEMIB.Cbfdefineobjecttable
    children["cbfStatusFileTable"] = &cISCOBULKFILEMIB.Cbfstatusfiletable
    return children
}

func (cISCOBULKFILEMIB *CISCOBULKFILEMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOBULKFILEMIB *CISCOBULKFILEMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOBULKFILEMIB *CISCOBULKFILEMIB) GetYangName() string { return "CISCO-BULK-FILE-MIB" }

func (cISCOBULKFILEMIB *CISCOBULKFILEMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOBULKFILEMIB *CISCOBULKFILEMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOBULKFILEMIB *CISCOBULKFILEMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOBULKFILEMIB *CISCOBULKFILEMIB) SetParent(parent types.Entity) { cISCOBULKFILEMIB.parent = parent }

func (cISCOBULKFILEMIB *CISCOBULKFILEMIB) GetParent() types.Entity { return cISCOBULKFILEMIB.parent }

func (cISCOBULKFILEMIB *CISCOBULKFILEMIB) GetParentYangName() string { return "CISCO-BULK-FILE-MIB" }

// CISCOBULKFILEMIB_Cbfdefine
type CISCOBULKFILEMIB_Cbfdefine struct {
    parent types.Entity
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

func (cbfdefine *CISCOBULKFILEMIB_Cbfdefine) GetFilter() yfilter.YFilter { return cbfdefine.YFilter }

func (cbfdefine *CISCOBULKFILEMIB_Cbfdefine) SetFilter(yf yfilter.YFilter) { cbfdefine.YFilter = yf }

func (cbfdefine *CISCOBULKFILEMIB_Cbfdefine) GetGoName(yname string) string {
    if yname == "cbfDefineMaxFiles" { return "Cbfdefinemaxfiles" }
    if yname == "cbfDefineFiles" { return "Cbfdefinefiles" }
    if yname == "cbfDefineHighFiles" { return "Cbfdefinehighfiles" }
    if yname == "cbfDefineFilesRefused" { return "Cbfdefinefilesrefused" }
    if yname == "cbfDefineMaxObjects" { return "Cbfdefinemaxobjects" }
    if yname == "cbfDefineObjects" { return "Cbfdefineobjects" }
    if yname == "cbfDefineHighObjects" { return "Cbfdefinehighobjects" }
    if yname == "cbfDefineObjectsRefused" { return "Cbfdefineobjectsrefused" }
    return ""
}

func (cbfdefine *CISCOBULKFILEMIB_Cbfdefine) GetSegmentPath() string {
    return "cbfDefine"
}

func (cbfdefine *CISCOBULKFILEMIB_Cbfdefine) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cbfdefine *CISCOBULKFILEMIB_Cbfdefine) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cbfdefine *CISCOBULKFILEMIB_Cbfdefine) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cbfDefineMaxFiles"] = cbfdefine.Cbfdefinemaxfiles
    leafs["cbfDefineFiles"] = cbfdefine.Cbfdefinefiles
    leafs["cbfDefineHighFiles"] = cbfdefine.Cbfdefinehighfiles
    leafs["cbfDefineFilesRefused"] = cbfdefine.Cbfdefinefilesrefused
    leafs["cbfDefineMaxObjects"] = cbfdefine.Cbfdefinemaxobjects
    leafs["cbfDefineObjects"] = cbfdefine.Cbfdefineobjects
    leafs["cbfDefineHighObjects"] = cbfdefine.Cbfdefinehighobjects
    leafs["cbfDefineObjectsRefused"] = cbfdefine.Cbfdefineobjectsrefused
    return leafs
}

func (cbfdefine *CISCOBULKFILEMIB_Cbfdefine) GetBundleName() string { return "cisco_ios_xe" }

func (cbfdefine *CISCOBULKFILEMIB_Cbfdefine) GetYangName() string { return "cbfDefine" }

func (cbfdefine *CISCOBULKFILEMIB_Cbfdefine) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cbfdefine *CISCOBULKFILEMIB_Cbfdefine) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cbfdefine *CISCOBULKFILEMIB_Cbfdefine) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cbfdefine *CISCOBULKFILEMIB_Cbfdefine) SetParent(parent types.Entity) { cbfdefine.parent = parent }

func (cbfdefine *CISCOBULKFILEMIB_Cbfdefine) GetParent() types.Entity { return cbfdefine.parent }

func (cbfdefine *CISCOBULKFILEMIB_Cbfdefine) GetParentYangName() string { return "CISCO-BULK-FILE-MIB" }

// CISCOBULKFILEMIB_Cbfstatus
type CISCOBULKFILEMIB_Cbfstatus struct {
    parent types.Entity
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

func (cbfstatus *CISCOBULKFILEMIB_Cbfstatus) GetFilter() yfilter.YFilter { return cbfstatus.YFilter }

func (cbfstatus *CISCOBULKFILEMIB_Cbfstatus) SetFilter(yf yfilter.YFilter) { cbfstatus.YFilter = yf }

func (cbfstatus *CISCOBULKFILEMIB_Cbfstatus) GetGoName(yname string) string {
    if yname == "cbfStatusMaxFiles" { return "Cbfstatusmaxfiles" }
    if yname == "cbfStatusFiles" { return "Cbfstatusfiles" }
    if yname == "cbfStatusHighFiles" { return "Cbfstatushighfiles" }
    if yname == "cbfStatusFilesBumped" { return "Cbfstatusfilesbumped" }
    return ""
}

func (cbfstatus *CISCOBULKFILEMIB_Cbfstatus) GetSegmentPath() string {
    return "cbfStatus"
}

func (cbfstatus *CISCOBULKFILEMIB_Cbfstatus) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cbfstatus *CISCOBULKFILEMIB_Cbfstatus) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cbfstatus *CISCOBULKFILEMIB_Cbfstatus) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cbfStatusMaxFiles"] = cbfstatus.Cbfstatusmaxfiles
    leafs["cbfStatusFiles"] = cbfstatus.Cbfstatusfiles
    leafs["cbfStatusHighFiles"] = cbfstatus.Cbfstatushighfiles
    leafs["cbfStatusFilesBumped"] = cbfstatus.Cbfstatusfilesbumped
    return leafs
}

func (cbfstatus *CISCOBULKFILEMIB_Cbfstatus) GetBundleName() string { return "cisco_ios_xe" }

func (cbfstatus *CISCOBULKFILEMIB_Cbfstatus) GetYangName() string { return "cbfStatus" }

func (cbfstatus *CISCOBULKFILEMIB_Cbfstatus) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cbfstatus *CISCOBULKFILEMIB_Cbfstatus) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cbfstatus *CISCOBULKFILEMIB_Cbfstatus) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cbfstatus *CISCOBULKFILEMIB_Cbfstatus) SetParent(parent types.Entity) { cbfstatus.parent = parent }

func (cbfstatus *CISCOBULKFILEMIB_Cbfstatus) GetParent() types.Entity { return cbfstatus.parent }

func (cbfstatus *CISCOBULKFILEMIB_Cbfstatus) GetParentYangName() string { return "CISCO-BULK-FILE-MIB" }

// CISCOBULKFILEMIB_Cbfdefinefiletable
// A table of bulk file definition and creation controls.
type CISCOBULKFILEMIB_Cbfdefinefiletable struct {
    parent types.Entity
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

func (cbfdefinefiletable *CISCOBULKFILEMIB_Cbfdefinefiletable) GetFilter() yfilter.YFilter { return cbfdefinefiletable.YFilter }

func (cbfdefinefiletable *CISCOBULKFILEMIB_Cbfdefinefiletable) SetFilter(yf yfilter.YFilter) { cbfdefinefiletable.YFilter = yf }

func (cbfdefinefiletable *CISCOBULKFILEMIB_Cbfdefinefiletable) GetGoName(yname string) string {
    if yname == "cbfDefineFileEntry" { return "Cbfdefinefileentry" }
    return ""
}

func (cbfdefinefiletable *CISCOBULKFILEMIB_Cbfdefinefiletable) GetSegmentPath() string {
    return "cbfDefineFileTable"
}

func (cbfdefinefiletable *CISCOBULKFILEMIB_Cbfdefinefiletable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cbfDefineFileEntry" {
        for _, c := range cbfdefinefiletable.Cbfdefinefileentry {
            if cbfdefinefiletable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOBULKFILEMIB_Cbfdefinefiletable_Cbfdefinefileentry{}
        cbfdefinefiletable.Cbfdefinefileentry = append(cbfdefinefiletable.Cbfdefinefileentry, child)
        return &cbfdefinefiletable.Cbfdefinefileentry[len(cbfdefinefiletable.Cbfdefinefileentry)-1]
    }
    return nil
}

func (cbfdefinefiletable *CISCOBULKFILEMIB_Cbfdefinefiletable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cbfdefinefiletable.Cbfdefinefileentry {
        children[cbfdefinefiletable.Cbfdefinefileentry[i].GetSegmentPath()] = &cbfdefinefiletable.Cbfdefinefileentry[i]
    }
    return children
}

func (cbfdefinefiletable *CISCOBULKFILEMIB_Cbfdefinefiletable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cbfdefinefiletable *CISCOBULKFILEMIB_Cbfdefinefiletable) GetBundleName() string { return "cisco_ios_xe" }

func (cbfdefinefiletable *CISCOBULKFILEMIB_Cbfdefinefiletable) GetYangName() string { return "cbfDefineFileTable" }

func (cbfdefinefiletable *CISCOBULKFILEMIB_Cbfdefinefiletable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cbfdefinefiletable *CISCOBULKFILEMIB_Cbfdefinefiletable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cbfdefinefiletable *CISCOBULKFILEMIB_Cbfdefinefiletable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cbfdefinefiletable *CISCOBULKFILEMIB_Cbfdefinefiletable) SetParent(parent types.Entity) { cbfdefinefiletable.parent = parent }

func (cbfdefinefiletable *CISCOBULKFILEMIB_Cbfdefinefiletable) GetParent() types.Entity { return cbfdefinefiletable.parent }

func (cbfdefinefiletable *CISCOBULKFILEMIB_Cbfdefinefiletable) GetParentYangName() string { return "CISCO-BULK-FILE-MIB" }

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
    parent types.Entity
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

func (cbfdefinefileentry *CISCOBULKFILEMIB_Cbfdefinefiletable_Cbfdefinefileentry) GetFilter() yfilter.YFilter { return cbfdefinefileentry.YFilter }

func (cbfdefinefileentry *CISCOBULKFILEMIB_Cbfdefinefiletable_Cbfdefinefileentry) SetFilter(yf yfilter.YFilter) { cbfdefinefileentry.YFilter = yf }

func (cbfdefinefileentry *CISCOBULKFILEMIB_Cbfdefinefiletable_Cbfdefinefileentry) GetGoName(yname string) string {
    if yname == "cbfDefineFileIndex" { return "Cbfdefinefileindex" }
    if yname == "cbfDefineFileName" { return "Cbfdefinefilename" }
    if yname == "cbfDefineFileStorage" { return "Cbfdefinefilestorage" }
    if yname == "cbfDefineFileFormat" { return "Cbfdefinefileformat" }
    if yname == "cbfDefineFileNow" { return "Cbfdefinefilenow" }
    if yname == "cbfDefineFileEntryStatus" { return "Cbfdefinefileentrystatus" }
    if yname == "cbfDefineFileNotifyOnCompletion" { return "Cbfdefinefilenotifyoncompletion" }
    return ""
}

func (cbfdefinefileentry *CISCOBULKFILEMIB_Cbfdefinefiletable_Cbfdefinefileentry) GetSegmentPath() string {
    return "cbfDefineFileEntry" + "[cbfDefineFileIndex='" + fmt.Sprintf("%v", cbfdefinefileentry.Cbfdefinefileindex) + "']"
}

func (cbfdefinefileentry *CISCOBULKFILEMIB_Cbfdefinefiletable_Cbfdefinefileentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cbfdefinefileentry *CISCOBULKFILEMIB_Cbfdefinefiletable_Cbfdefinefileentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cbfdefinefileentry *CISCOBULKFILEMIB_Cbfdefinefiletable_Cbfdefinefileentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cbfDefineFileIndex"] = cbfdefinefileentry.Cbfdefinefileindex
    leafs["cbfDefineFileName"] = cbfdefinefileentry.Cbfdefinefilename
    leafs["cbfDefineFileStorage"] = cbfdefinefileentry.Cbfdefinefilestorage
    leafs["cbfDefineFileFormat"] = cbfdefinefileentry.Cbfdefinefileformat
    leafs["cbfDefineFileNow"] = cbfdefinefileentry.Cbfdefinefilenow
    leafs["cbfDefineFileEntryStatus"] = cbfdefinefileentry.Cbfdefinefileentrystatus
    leafs["cbfDefineFileNotifyOnCompletion"] = cbfdefinefileentry.Cbfdefinefilenotifyoncompletion
    return leafs
}

func (cbfdefinefileentry *CISCOBULKFILEMIB_Cbfdefinefiletable_Cbfdefinefileentry) GetBundleName() string { return "cisco_ios_xe" }

func (cbfdefinefileentry *CISCOBULKFILEMIB_Cbfdefinefiletable_Cbfdefinefileentry) GetYangName() string { return "cbfDefineFileEntry" }

func (cbfdefinefileentry *CISCOBULKFILEMIB_Cbfdefinefiletable_Cbfdefinefileentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cbfdefinefileentry *CISCOBULKFILEMIB_Cbfdefinefiletable_Cbfdefinefileentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cbfdefinefileentry *CISCOBULKFILEMIB_Cbfdefinefiletable_Cbfdefinefileentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cbfdefinefileentry *CISCOBULKFILEMIB_Cbfdefinefiletable_Cbfdefinefileentry) SetParent(parent types.Entity) { cbfdefinefileentry.parent = parent }

func (cbfdefinefileentry *CISCOBULKFILEMIB_Cbfdefinefiletable_Cbfdefinefileentry) GetParent() types.Entity { return cbfdefinefileentry.parent }

func (cbfdefinefileentry *CISCOBULKFILEMIB_Cbfdefinefiletable_Cbfdefinefileentry) GetParentYangName() string { return "cbfDefineFileTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about one object for a particular file.  An application uses
    // cbfDefineObjectEntryStatus to create entries in this table in
    // correspondence with entries in cbfDefineFileTable, which must be created
    // first.  Entries in this table may not be changed, created or deleted while
    // the corresponding value of cbfDefineFileNow is 'running'. The type is slice
    // of CISCOBULKFILEMIB_Cbfdefineobjecttable_Cbfdefineobjectentry.
    Cbfdefineobjectentry []CISCOBULKFILEMIB_Cbfdefineobjecttable_Cbfdefineobjectentry
}

func (cbfdefineobjecttable *CISCOBULKFILEMIB_Cbfdefineobjecttable) GetFilter() yfilter.YFilter { return cbfdefineobjecttable.YFilter }

func (cbfdefineobjecttable *CISCOBULKFILEMIB_Cbfdefineobjecttable) SetFilter(yf yfilter.YFilter) { cbfdefineobjecttable.YFilter = yf }

func (cbfdefineobjecttable *CISCOBULKFILEMIB_Cbfdefineobjecttable) GetGoName(yname string) string {
    if yname == "cbfDefineObjectEntry" { return "Cbfdefineobjectentry" }
    return ""
}

func (cbfdefineobjecttable *CISCOBULKFILEMIB_Cbfdefineobjecttable) GetSegmentPath() string {
    return "cbfDefineObjectTable"
}

func (cbfdefineobjecttable *CISCOBULKFILEMIB_Cbfdefineobjecttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cbfDefineObjectEntry" {
        for _, c := range cbfdefineobjecttable.Cbfdefineobjectentry {
            if cbfdefineobjecttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOBULKFILEMIB_Cbfdefineobjecttable_Cbfdefineobjectentry{}
        cbfdefineobjecttable.Cbfdefineobjectentry = append(cbfdefineobjecttable.Cbfdefineobjectentry, child)
        return &cbfdefineobjecttable.Cbfdefineobjectentry[len(cbfdefineobjecttable.Cbfdefineobjectentry)-1]
    }
    return nil
}

func (cbfdefineobjecttable *CISCOBULKFILEMIB_Cbfdefineobjecttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cbfdefineobjecttable.Cbfdefineobjectentry {
        children[cbfdefineobjecttable.Cbfdefineobjectentry[i].GetSegmentPath()] = &cbfdefineobjecttable.Cbfdefineobjectentry[i]
    }
    return children
}

func (cbfdefineobjecttable *CISCOBULKFILEMIB_Cbfdefineobjecttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cbfdefineobjecttable *CISCOBULKFILEMIB_Cbfdefineobjecttable) GetBundleName() string { return "cisco_ios_xe" }

func (cbfdefineobjecttable *CISCOBULKFILEMIB_Cbfdefineobjecttable) GetYangName() string { return "cbfDefineObjectTable" }

func (cbfdefineobjecttable *CISCOBULKFILEMIB_Cbfdefineobjecttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cbfdefineobjecttable *CISCOBULKFILEMIB_Cbfdefineobjecttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cbfdefineobjecttable *CISCOBULKFILEMIB_Cbfdefineobjecttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cbfdefineobjecttable *CISCOBULKFILEMIB_Cbfdefineobjecttable) SetParent(parent types.Entity) { cbfdefineobjecttable.parent = parent }

func (cbfdefineobjecttable *CISCOBULKFILEMIB_Cbfdefineobjecttable) GetParent() types.Entity { return cbfdefineobjecttable.parent }

func (cbfdefineobjecttable *CISCOBULKFILEMIB_Cbfdefineobjecttable) GetParentYangName() string { return "CISCO-BULK-FILE-MIB" }

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
    parent types.Entity
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
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
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
    // pattern: (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
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
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    Cbfdefineobjectlastpolledinst interface{}
}

func (cbfdefineobjectentry *CISCOBULKFILEMIB_Cbfdefineobjecttable_Cbfdefineobjectentry) GetFilter() yfilter.YFilter { return cbfdefineobjectentry.YFilter }

func (cbfdefineobjectentry *CISCOBULKFILEMIB_Cbfdefineobjecttable_Cbfdefineobjectentry) SetFilter(yf yfilter.YFilter) { cbfdefineobjectentry.YFilter = yf }

func (cbfdefineobjectentry *CISCOBULKFILEMIB_Cbfdefineobjecttable_Cbfdefineobjectentry) GetGoName(yname string) string {
    if yname == "cbfDefineFileIndex" { return "Cbfdefinefileindex" }
    if yname == "cbfDefineObjectIndex" { return "Cbfdefineobjectindex" }
    if yname == "cbfDefineObjectClass" { return "Cbfdefineobjectclass" }
    if yname == "cbfDefineObjectID" { return "Cbfdefineobjectid" }
    if yname == "cbfDefineObjectEntryStatus" { return "Cbfdefineobjectentrystatus" }
    if yname == "cbfDefineObjectTableInstance" { return "Cbfdefineobjecttableinstance" }
    if yname == "cbfDefineObjectNumEntries" { return "Cbfdefineobjectnumentries" }
    if yname == "cbfDefineObjectLastPolledInst" { return "Cbfdefineobjectlastpolledinst" }
    return ""
}

func (cbfdefineobjectentry *CISCOBULKFILEMIB_Cbfdefineobjecttable_Cbfdefineobjectentry) GetSegmentPath() string {
    return "cbfDefineObjectEntry" + "[cbfDefineFileIndex='" + fmt.Sprintf("%v", cbfdefineobjectentry.Cbfdefinefileindex) + "']" + "[cbfDefineObjectIndex='" + fmt.Sprintf("%v", cbfdefineobjectentry.Cbfdefineobjectindex) + "']"
}

func (cbfdefineobjectentry *CISCOBULKFILEMIB_Cbfdefineobjecttable_Cbfdefineobjectentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cbfdefineobjectentry *CISCOBULKFILEMIB_Cbfdefineobjecttable_Cbfdefineobjectentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cbfdefineobjectentry *CISCOBULKFILEMIB_Cbfdefineobjecttable_Cbfdefineobjectentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cbfDefineFileIndex"] = cbfdefineobjectentry.Cbfdefinefileindex
    leafs["cbfDefineObjectIndex"] = cbfdefineobjectentry.Cbfdefineobjectindex
    leafs["cbfDefineObjectClass"] = cbfdefineobjectentry.Cbfdefineobjectclass
    leafs["cbfDefineObjectID"] = cbfdefineobjectentry.Cbfdefineobjectid
    leafs["cbfDefineObjectEntryStatus"] = cbfdefineobjectentry.Cbfdefineobjectentrystatus
    leafs["cbfDefineObjectTableInstance"] = cbfdefineobjectentry.Cbfdefineobjecttableinstance
    leafs["cbfDefineObjectNumEntries"] = cbfdefineobjectentry.Cbfdefineobjectnumentries
    leafs["cbfDefineObjectLastPolledInst"] = cbfdefineobjectentry.Cbfdefineobjectlastpolledinst
    return leafs
}

func (cbfdefineobjectentry *CISCOBULKFILEMIB_Cbfdefineobjecttable_Cbfdefineobjectentry) GetBundleName() string { return "cisco_ios_xe" }

func (cbfdefineobjectentry *CISCOBULKFILEMIB_Cbfdefineobjecttable_Cbfdefineobjectentry) GetYangName() string { return "cbfDefineObjectEntry" }

func (cbfdefineobjectentry *CISCOBULKFILEMIB_Cbfdefineobjecttable_Cbfdefineobjectentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cbfdefineobjectentry *CISCOBULKFILEMIB_Cbfdefineobjecttable_Cbfdefineobjectentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cbfdefineobjectentry *CISCOBULKFILEMIB_Cbfdefineobjecttable_Cbfdefineobjectentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cbfdefineobjectentry *CISCOBULKFILEMIB_Cbfdefineobjecttable_Cbfdefineobjectentry) SetParent(parent types.Entity) { cbfdefineobjectentry.parent = parent }

func (cbfdefineobjectentry *CISCOBULKFILEMIB_Cbfdefineobjecttable_Cbfdefineobjectentry) GetParent() types.Entity { return cbfdefineobjectentry.parent }

func (cbfdefineobjectentry *CISCOBULKFILEMIB_Cbfdefineobjecttable_Cbfdefineobjectentry) GetParentYangName() string { return "cbfDefineObjectTable" }

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
    parent types.Entity
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

func (cbfstatusfiletable *CISCOBULKFILEMIB_Cbfstatusfiletable) GetFilter() yfilter.YFilter { return cbfstatusfiletable.YFilter }

func (cbfstatusfiletable *CISCOBULKFILEMIB_Cbfstatusfiletable) SetFilter(yf yfilter.YFilter) { cbfstatusfiletable.YFilter = yf }

func (cbfstatusfiletable *CISCOBULKFILEMIB_Cbfstatusfiletable) GetGoName(yname string) string {
    if yname == "cbfStatusFileEntry" { return "Cbfstatusfileentry" }
    return ""
}

func (cbfstatusfiletable *CISCOBULKFILEMIB_Cbfstatusfiletable) GetSegmentPath() string {
    return "cbfStatusFileTable"
}

func (cbfstatusfiletable *CISCOBULKFILEMIB_Cbfstatusfiletable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cbfStatusFileEntry" {
        for _, c := range cbfstatusfiletable.Cbfstatusfileentry {
            if cbfstatusfiletable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOBULKFILEMIB_Cbfstatusfiletable_Cbfstatusfileentry{}
        cbfstatusfiletable.Cbfstatusfileentry = append(cbfstatusfiletable.Cbfstatusfileentry, child)
        return &cbfstatusfiletable.Cbfstatusfileentry[len(cbfstatusfiletable.Cbfstatusfileentry)-1]
    }
    return nil
}

func (cbfstatusfiletable *CISCOBULKFILEMIB_Cbfstatusfiletable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cbfstatusfiletable.Cbfstatusfileentry {
        children[cbfstatusfiletable.Cbfstatusfileentry[i].GetSegmentPath()] = &cbfstatusfiletable.Cbfstatusfileentry[i]
    }
    return children
}

func (cbfstatusfiletable *CISCOBULKFILEMIB_Cbfstatusfiletable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cbfstatusfiletable *CISCOBULKFILEMIB_Cbfstatusfiletable) GetBundleName() string { return "cisco_ios_xe" }

func (cbfstatusfiletable *CISCOBULKFILEMIB_Cbfstatusfiletable) GetYangName() string { return "cbfStatusFileTable" }

func (cbfstatusfiletable *CISCOBULKFILEMIB_Cbfstatusfiletable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cbfstatusfiletable *CISCOBULKFILEMIB_Cbfstatusfiletable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cbfstatusfiletable *CISCOBULKFILEMIB_Cbfstatusfiletable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cbfstatusfiletable *CISCOBULKFILEMIB_Cbfstatusfiletable) SetParent(parent types.Entity) { cbfstatusfiletable.parent = parent }

func (cbfstatusfiletable *CISCOBULKFILEMIB_Cbfstatusfiletable) GetParent() types.Entity { return cbfstatusfiletable.parent }

func (cbfstatusfiletable *CISCOBULKFILEMIB_Cbfstatusfiletable) GetParentYangName() string { return "CISCO-BULK-FILE-MIB" }

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
    parent types.Entity
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

func (cbfstatusfileentry *CISCOBULKFILEMIB_Cbfstatusfiletable_Cbfstatusfileentry) GetFilter() yfilter.YFilter { return cbfstatusfileentry.YFilter }

func (cbfstatusfileentry *CISCOBULKFILEMIB_Cbfstatusfiletable_Cbfstatusfileentry) SetFilter(yf yfilter.YFilter) { cbfstatusfileentry.YFilter = yf }

func (cbfstatusfileentry *CISCOBULKFILEMIB_Cbfstatusfiletable_Cbfstatusfileentry) GetGoName(yname string) string {
    if yname == "cbfDefineFileIndex" { return "Cbfdefinefileindex" }
    if yname == "cbfStatusFileIndex" { return "Cbfstatusfileindex" }
    if yname == "cbfStatusFileState" { return "Cbfstatusfilestate" }
    if yname == "cbfStatusFileCompletionTime" { return "Cbfstatusfilecompletiontime" }
    if yname == "cbfStatusFileEntryStatus" { return "Cbfstatusfileentrystatus" }
    return ""
}

func (cbfstatusfileentry *CISCOBULKFILEMIB_Cbfstatusfiletable_Cbfstatusfileentry) GetSegmentPath() string {
    return "cbfStatusFileEntry" + "[cbfDefineFileIndex='" + fmt.Sprintf("%v", cbfstatusfileentry.Cbfdefinefileindex) + "']" + "[cbfStatusFileIndex='" + fmt.Sprintf("%v", cbfstatusfileentry.Cbfstatusfileindex) + "']"
}

func (cbfstatusfileentry *CISCOBULKFILEMIB_Cbfstatusfiletable_Cbfstatusfileentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cbfstatusfileentry *CISCOBULKFILEMIB_Cbfstatusfiletable_Cbfstatusfileentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cbfstatusfileentry *CISCOBULKFILEMIB_Cbfstatusfiletable_Cbfstatusfileentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cbfDefineFileIndex"] = cbfstatusfileentry.Cbfdefinefileindex
    leafs["cbfStatusFileIndex"] = cbfstatusfileentry.Cbfstatusfileindex
    leafs["cbfStatusFileState"] = cbfstatusfileentry.Cbfstatusfilestate
    leafs["cbfStatusFileCompletionTime"] = cbfstatusfileentry.Cbfstatusfilecompletiontime
    leafs["cbfStatusFileEntryStatus"] = cbfstatusfileentry.Cbfstatusfileentrystatus
    return leafs
}

func (cbfstatusfileentry *CISCOBULKFILEMIB_Cbfstatusfiletable_Cbfstatusfileentry) GetBundleName() string { return "cisco_ios_xe" }

func (cbfstatusfileentry *CISCOBULKFILEMIB_Cbfstatusfiletable_Cbfstatusfileentry) GetYangName() string { return "cbfStatusFileEntry" }

func (cbfstatusfileentry *CISCOBULKFILEMIB_Cbfstatusfiletable_Cbfstatusfileentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cbfstatusfileentry *CISCOBULKFILEMIB_Cbfstatusfiletable_Cbfstatusfileentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cbfstatusfileentry *CISCOBULKFILEMIB_Cbfstatusfiletable_Cbfstatusfileentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cbfstatusfileentry *CISCOBULKFILEMIB_Cbfstatusfiletable_Cbfstatusfileentry) SetParent(parent types.Entity) { cbfstatusfileentry.parent = parent }

func (cbfstatusfileentry *CISCOBULKFILEMIB_Cbfstatusfiletable_Cbfstatusfileentry) GetParent() types.Entity { return cbfstatusfileentry.parent }

func (cbfstatusfileentry *CISCOBULKFILEMIB_Cbfstatusfiletable_Cbfstatusfileentry) GetParentYangName() string { return "cbfStatusFileTable" }

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

