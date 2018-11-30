// The MIB module for invoking Internet File Transfer Protocol
// (FTP) operations for network management purposes.
package cisco_ftp_client_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_ftp_client_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-FTP-CLIENT-MIB CISCO-FTP-CLIENT-MIB}", reflect.TypeOf(CISCOFTPCLIENTMIB{}))
    ydk.RegisterEntity("CISCO-FTP-CLIENT-MIB:CISCO-FTP-CLIENT-MIB", reflect.TypeOf(CISCOFTPCLIENTMIB{}))
}

// CISCOFTPCLIENTMIB
type CISCOFTPCLIENTMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    CfcRequest CISCOFTPCLIENTMIB_CfcRequest

    // A table of FTP client requests.
    CfcRequestTable CISCOFTPCLIENTMIB_CfcRequestTable
}

func (cISCOFTPCLIENTMIB *CISCOFTPCLIENTMIB) GetEntityData() *types.CommonEntityData {
    cISCOFTPCLIENTMIB.EntityData.YFilter = cISCOFTPCLIENTMIB.YFilter
    cISCOFTPCLIENTMIB.EntityData.YangName = "CISCO-FTP-CLIENT-MIB"
    cISCOFTPCLIENTMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOFTPCLIENTMIB.EntityData.ParentYangName = "CISCO-FTP-CLIENT-MIB"
    cISCOFTPCLIENTMIB.EntityData.SegmentPath = "CISCO-FTP-CLIENT-MIB:CISCO-FTP-CLIENT-MIB"
    cISCOFTPCLIENTMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOFTPCLIENTMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOFTPCLIENTMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOFTPCLIENTMIB.EntityData.Children = types.NewOrderedMap()
    cISCOFTPCLIENTMIB.EntityData.Children.Append("cfcRequest", types.YChild{"CfcRequest", &cISCOFTPCLIENTMIB.CfcRequest})
    cISCOFTPCLIENTMIB.EntityData.Children.Append("cfcRequestTable", types.YChild{"CfcRequestTable", &cISCOFTPCLIENTMIB.CfcRequestTable})
    cISCOFTPCLIENTMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOFTPCLIENTMIB.EntityData.YListKeys = []string {}

    return &(cISCOFTPCLIENTMIB.EntityData)
}

// CISCOFTPCLIENTMIB_CfcRequest
type CISCOFTPCLIENTMIB_CfcRequest struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The maximum number of requests this system can hold in cfcRequestTable.  A
    // value of 0 indicates no configured limit.  This object may be read-only on
    // some systems.  When an attempt is made to create a new entry but the table
    // is full, the oldest completed entry is bumped out and cfcRequestsBumped is
    // incremented.  Changing this number does not disturb existing requests that
    // are not completed and bumps completed requests as necessary. The type is
    // interface{} with range: 0..4294967295.
    CfcRequestMaximum interface{}

    // The current number of requests in cfcRequestTable. The type is interface{}
    // with range: 0..4294967295.
    CfcRequests interface{}

    // The highest number of requests in cfcRequestTable since this system was
    // last initialized. The type is interface{} with range: 0..4294967295.
    CfcRequestsHigh interface{}

    // The number of requests in cfcRequestTable that were bumped out to make room
    // for a new request. The type is interface{} with range: 0..4294967295.
    CfcRequestsBumped interface{}
}

func (cfcRequest *CISCOFTPCLIENTMIB_CfcRequest) GetEntityData() *types.CommonEntityData {
    cfcRequest.EntityData.YFilter = cfcRequest.YFilter
    cfcRequest.EntityData.YangName = "cfcRequest"
    cfcRequest.EntityData.BundleName = "cisco_ios_xe"
    cfcRequest.EntityData.ParentYangName = "CISCO-FTP-CLIENT-MIB"
    cfcRequest.EntityData.SegmentPath = "cfcRequest"
    cfcRequest.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cfcRequest.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cfcRequest.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cfcRequest.EntityData.Children = types.NewOrderedMap()
    cfcRequest.EntityData.Leafs = types.NewOrderedMap()
    cfcRequest.EntityData.Leafs.Append("cfcRequestMaximum", types.YLeaf{"CfcRequestMaximum", cfcRequest.CfcRequestMaximum})
    cfcRequest.EntityData.Leafs.Append("cfcRequests", types.YLeaf{"CfcRequests", cfcRequest.CfcRequests})
    cfcRequest.EntityData.Leafs.Append("cfcRequestsHigh", types.YLeaf{"CfcRequestsHigh", cfcRequest.CfcRequestsHigh})
    cfcRequest.EntityData.Leafs.Append("cfcRequestsBumped", types.YLeaf{"CfcRequestsBumped", cfcRequest.CfcRequestsBumped})

    cfcRequest.EntityData.YListKeys = []string {}

    return &(cfcRequest.EntityData)
}

// CISCOFTPCLIENTMIB_CfcRequestTable
// A table of FTP client requests.
type CISCOFTPCLIENTMIB_CfcRequestTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an FTP client request.  Management applications use
    // cfcRequestEntryStatus to control entry modification, creation, and
    // deletion.  Setting cfcRequestEntryStatus to 'active' from any state
    // including 'active' causes the operation to be started.  The entry may be
    // modified only when cfcRequestOperationState is 'stopped'.  The value of
    // cfcRequestEntryStatus may be set to 'destroy' at any time.  Doing so will
    // abort a running request.  Entries may not be created without explicitly
    // setting cfcRequestEntryStatus to either 'createAndGo' or 'createAndWait'.
    // The type is slice of CISCOFTPCLIENTMIB_CfcRequestTable_CfcRequestEntry.
    CfcRequestEntry []*CISCOFTPCLIENTMIB_CfcRequestTable_CfcRequestEntry
}

func (cfcRequestTable *CISCOFTPCLIENTMIB_CfcRequestTable) GetEntityData() *types.CommonEntityData {
    cfcRequestTable.EntityData.YFilter = cfcRequestTable.YFilter
    cfcRequestTable.EntityData.YangName = "cfcRequestTable"
    cfcRequestTable.EntityData.BundleName = "cisco_ios_xe"
    cfcRequestTable.EntityData.ParentYangName = "CISCO-FTP-CLIENT-MIB"
    cfcRequestTable.EntityData.SegmentPath = "cfcRequestTable"
    cfcRequestTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cfcRequestTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cfcRequestTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cfcRequestTable.EntityData.Children = types.NewOrderedMap()
    cfcRequestTable.EntityData.Children.Append("cfcRequestEntry", types.YChild{"CfcRequestEntry", nil})
    for i := range cfcRequestTable.CfcRequestEntry {
        cfcRequestTable.EntityData.Children.Append(types.GetSegmentPath(cfcRequestTable.CfcRequestEntry[i]), types.YChild{"CfcRequestEntry", cfcRequestTable.CfcRequestEntry[i]})
    }
    cfcRequestTable.EntityData.Leafs = types.NewOrderedMap()

    cfcRequestTable.EntityData.YListKeys = []string {}

    return &(cfcRequestTable.EntityData)
}

// CISCOFTPCLIENTMIB_CfcRequestTable_CfcRequestEntry
// Information about an FTP client request.  Management applications
// use cfcRequestEntryStatus to control entry modification, creation,
// and deletion.
// 
// Setting cfcRequestEntryStatus to 'active' from any state including
// 'active' causes the operation to be started.
// 
// The entry may be modified only when cfcRequestOperationState is
// 'stopped'.
// 
// The value of cfcRequestEntryStatus may be set to 'destroy' at any
// time.  Doing so will abort a running request.
// 
// Entries may not be created without explicitly setting
// cfcRequestEntryStatus to either 'createAndGo' or 'createAndWait'.
type CISCOFTPCLIENTMIB_CfcRequestTable_CfcRequestEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An arbitrary integer to uniquely identify this
    // entry.  To create an entry a management application should pick a random
    // number. The type is interface{} with range: 1..4294967295.
    CfcRequestIndex interface{}

    // The FTP operation to be performed. The type is CfcRequestOperation.
    CfcRequestOperation interface{}

    // The local file on which the operation is to be performed. The type is
    // string with length: 1..255.
    CfcRequestLocalFile interface{}

    // The remote file on which the operation is to be performed. The type is
    // string with length: 1..255.
    CfcRequestRemoteFile interface{}

    // The domain name or IP address of the FTP server to use. The type is string
    // with length: 1..64.
    CfcRequestServer interface{}

    // The user name to use at the FTP server. The type is string with length:
    // 1..32.
    CfcRequestUser interface{}

    // The password to use at the FTP server.  When read this object always
    // returns a zero-length string. The type is string with length: 0..16.
    CfcRequestPassword interface{}

    // The result of the FTP operation. The type is CfcRequestResult.
    CfcRequestResult interface{}

    // The value of sysUpTime when the operation completed.  For an incomplete
    // operation this value is zero. The type is interface{} with range:
    // 0..4294967295.
    CfcRequestCompletionTime interface{}

    // The action control to stop a running request.  Setting this to 'stop' will
    // begin the process of stopping the request.  Setting it to 'ready' or
    // setting it to 'stop' more than once have no effect.  When read this object
    // always returns ready. The type is CfcRequestStop.
    CfcRequestStop interface{}

    // The operational state of the file transfer.  To short-terminate the
    // transfer set cfcRequestStop to 'stop'. The type is
    // CfcRequestOperationState.
    CfcRequestOperationState interface{}

    // The control that allows modification, creation, and deletion of entries. 
    // For detailed rules see the DESCRIPTION for cfcRequestEntry. The type is
    // RowStatus.
    CfcRequestEntryStatus interface{}
}

func (cfcRequestEntry *CISCOFTPCLIENTMIB_CfcRequestTable_CfcRequestEntry) GetEntityData() *types.CommonEntityData {
    cfcRequestEntry.EntityData.YFilter = cfcRequestEntry.YFilter
    cfcRequestEntry.EntityData.YangName = "cfcRequestEntry"
    cfcRequestEntry.EntityData.BundleName = "cisco_ios_xe"
    cfcRequestEntry.EntityData.ParentYangName = "cfcRequestTable"
    cfcRequestEntry.EntityData.SegmentPath = "cfcRequestEntry" + types.AddKeyToken(cfcRequestEntry.CfcRequestIndex, "cfcRequestIndex")
    cfcRequestEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cfcRequestEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cfcRequestEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cfcRequestEntry.EntityData.Children = types.NewOrderedMap()
    cfcRequestEntry.EntityData.Leafs = types.NewOrderedMap()
    cfcRequestEntry.EntityData.Leafs.Append("cfcRequestIndex", types.YLeaf{"CfcRequestIndex", cfcRequestEntry.CfcRequestIndex})
    cfcRequestEntry.EntityData.Leafs.Append("cfcRequestOperation", types.YLeaf{"CfcRequestOperation", cfcRequestEntry.CfcRequestOperation})
    cfcRequestEntry.EntityData.Leafs.Append("cfcRequestLocalFile", types.YLeaf{"CfcRequestLocalFile", cfcRequestEntry.CfcRequestLocalFile})
    cfcRequestEntry.EntityData.Leafs.Append("cfcRequestRemoteFile", types.YLeaf{"CfcRequestRemoteFile", cfcRequestEntry.CfcRequestRemoteFile})
    cfcRequestEntry.EntityData.Leafs.Append("cfcRequestServer", types.YLeaf{"CfcRequestServer", cfcRequestEntry.CfcRequestServer})
    cfcRequestEntry.EntityData.Leafs.Append("cfcRequestUser", types.YLeaf{"CfcRequestUser", cfcRequestEntry.CfcRequestUser})
    cfcRequestEntry.EntityData.Leafs.Append("cfcRequestPassword", types.YLeaf{"CfcRequestPassword", cfcRequestEntry.CfcRequestPassword})
    cfcRequestEntry.EntityData.Leafs.Append("cfcRequestResult", types.YLeaf{"CfcRequestResult", cfcRequestEntry.CfcRequestResult})
    cfcRequestEntry.EntityData.Leafs.Append("cfcRequestCompletionTime", types.YLeaf{"CfcRequestCompletionTime", cfcRequestEntry.CfcRequestCompletionTime})
    cfcRequestEntry.EntityData.Leafs.Append("cfcRequestStop", types.YLeaf{"CfcRequestStop", cfcRequestEntry.CfcRequestStop})
    cfcRequestEntry.EntityData.Leafs.Append("cfcRequestOperationState", types.YLeaf{"CfcRequestOperationState", cfcRequestEntry.CfcRequestOperationState})
    cfcRequestEntry.EntityData.Leafs.Append("cfcRequestEntryStatus", types.YLeaf{"CfcRequestEntryStatus", cfcRequestEntry.CfcRequestEntryStatus})

    cfcRequestEntry.EntityData.YListKeys = []string {"CfcRequestIndex"}

    return &(cfcRequestEntry.EntityData)
}

// CISCOFTPCLIENTMIB_CfcRequestTable_CfcRequestEntry_CfcRequestOperation represents The FTP operation to be performed.
type CISCOFTPCLIENTMIB_CfcRequestTable_CfcRequestEntry_CfcRequestOperation string

const (
    CISCOFTPCLIENTMIB_CfcRequestTable_CfcRequestEntry_CfcRequestOperation_putBinary CISCOFTPCLIENTMIB_CfcRequestTable_CfcRequestEntry_CfcRequestOperation = "putBinary"

    CISCOFTPCLIENTMIB_CfcRequestTable_CfcRequestEntry_CfcRequestOperation_putASCII CISCOFTPCLIENTMIB_CfcRequestTable_CfcRequestEntry_CfcRequestOperation = "putASCII"
)

// CISCOFTPCLIENTMIB_CfcRequestTable_CfcRequestEntry_CfcRequestOperationState represents the transfer set cfcRequestStop to 'stop'.
type CISCOFTPCLIENTMIB_CfcRequestTable_CfcRequestEntry_CfcRequestOperationState string

const (
    CISCOFTPCLIENTMIB_CfcRequestTable_CfcRequestEntry_CfcRequestOperationState_running CISCOFTPCLIENTMIB_CfcRequestTable_CfcRequestEntry_CfcRequestOperationState = "running"

    CISCOFTPCLIENTMIB_CfcRequestTable_CfcRequestEntry_CfcRequestOperationState_stopping CISCOFTPCLIENTMIB_CfcRequestTable_CfcRequestEntry_CfcRequestOperationState = "stopping"

    CISCOFTPCLIENTMIB_CfcRequestTable_CfcRequestEntry_CfcRequestOperationState_stopped CISCOFTPCLIENTMIB_CfcRequestTable_CfcRequestEntry_CfcRequestOperationState = "stopped"
)

// CISCOFTPCLIENTMIB_CfcRequestTable_CfcRequestEntry_CfcRequestResult represents The result of the FTP operation.
type CISCOFTPCLIENTMIB_CfcRequestTable_CfcRequestEntry_CfcRequestResult string

const (
    CISCOFTPCLIENTMIB_CfcRequestTable_CfcRequestEntry_CfcRequestResult_pending CISCOFTPCLIENTMIB_CfcRequestTable_CfcRequestEntry_CfcRequestResult = "pending"

    CISCOFTPCLIENTMIB_CfcRequestTable_CfcRequestEntry_CfcRequestResult_success CISCOFTPCLIENTMIB_CfcRequestTable_CfcRequestEntry_CfcRequestResult = "success"

    CISCOFTPCLIENTMIB_CfcRequestTable_CfcRequestEntry_CfcRequestResult_aborted CISCOFTPCLIENTMIB_CfcRequestTable_CfcRequestEntry_CfcRequestResult = "aborted"

    CISCOFTPCLIENTMIB_CfcRequestTable_CfcRequestEntry_CfcRequestResult_fileOpenFailLocal CISCOFTPCLIENTMIB_CfcRequestTable_CfcRequestEntry_CfcRequestResult = "fileOpenFailLocal"

    CISCOFTPCLIENTMIB_CfcRequestTable_CfcRequestEntry_CfcRequestResult_fileOpenFailRemote CISCOFTPCLIENTMIB_CfcRequestTable_CfcRequestEntry_CfcRequestResult = "fileOpenFailRemote"

    CISCOFTPCLIENTMIB_CfcRequestTable_CfcRequestEntry_CfcRequestResult_badDomainName CISCOFTPCLIENTMIB_CfcRequestTable_CfcRequestEntry_CfcRequestResult = "badDomainName"

    CISCOFTPCLIENTMIB_CfcRequestTable_CfcRequestEntry_CfcRequestResult_unreachableIpAddress CISCOFTPCLIENTMIB_CfcRequestTable_CfcRequestEntry_CfcRequestResult = "unreachableIpAddress"

    CISCOFTPCLIENTMIB_CfcRequestTable_CfcRequestEntry_CfcRequestResult_linkFailed CISCOFTPCLIENTMIB_CfcRequestTable_CfcRequestEntry_CfcRequestResult = "linkFailed"

    CISCOFTPCLIENTMIB_CfcRequestTable_CfcRequestEntry_CfcRequestResult_fileReadFailed CISCOFTPCLIENTMIB_CfcRequestTable_CfcRequestEntry_CfcRequestResult = "fileReadFailed"

    CISCOFTPCLIENTMIB_CfcRequestTable_CfcRequestEntry_CfcRequestResult_fileWriteFailed CISCOFTPCLIENTMIB_CfcRequestTable_CfcRequestEntry_CfcRequestResult = "fileWriteFailed"
)

// CISCOFTPCLIENTMIB_CfcRequestTable_CfcRequestEntry_CfcRequestStop represents effect.  When read this object always returns ready.
type CISCOFTPCLIENTMIB_CfcRequestTable_CfcRequestEntry_CfcRequestStop string

const (
    CISCOFTPCLIENTMIB_CfcRequestTable_CfcRequestEntry_CfcRequestStop_ready CISCOFTPCLIENTMIB_CfcRequestTable_CfcRequestEntry_CfcRequestStop = "ready"

    CISCOFTPCLIENTMIB_CfcRequestTable_CfcRequestEntry_CfcRequestStop_stop CISCOFTPCLIENTMIB_CfcRequestTable_CfcRequestEntry_CfcRequestStop = "stop"
)

