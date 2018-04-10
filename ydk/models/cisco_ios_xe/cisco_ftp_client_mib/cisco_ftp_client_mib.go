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

    
    Cfcrequest CISCOFTPCLIENTMIB_Cfcrequest

    // A table of FTP client requests.
    Cfcrequesttable CISCOFTPCLIENTMIB_Cfcrequesttable
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

    cISCOFTPCLIENTMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOFTPCLIENTMIB.EntityData.Children["cfcRequest"] = types.YChild{"Cfcrequest", &cISCOFTPCLIENTMIB.Cfcrequest}
    cISCOFTPCLIENTMIB.EntityData.Children["cfcRequestTable"] = types.YChild{"Cfcrequesttable", &cISCOFTPCLIENTMIB.Cfcrequesttable}
    cISCOFTPCLIENTMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOFTPCLIENTMIB.EntityData)
}

// CISCOFTPCLIENTMIB_Cfcrequest
type CISCOFTPCLIENTMIB_Cfcrequest struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The maximum number of requests this system can hold in cfcRequestTable.  A
    // value of 0 indicates no configured limit.  This object may be read-only on
    // some systems.  When an attempt is made to create a new entry but the table
    // is full, the oldest completed entry is bumped out and cfcRequestsBumped is
    // incremented.  Changing this number does not disturb existing requests that
    // are not completed and bumps completed requests as necessary. The type is
    // interface{} with range: 0..4294967295.
    Cfcrequestmaximum interface{}

    // The current number of requests in cfcRequestTable. The type is interface{}
    // with range: 0..4294967295.
    Cfcrequests interface{}

    // The highest number of requests in cfcRequestTable since this system was
    // last initialized. The type is interface{} with range: 0..4294967295.
    Cfcrequestshigh interface{}

    // The number of requests in cfcRequestTable that were bumped out to make room
    // for a new request. The type is interface{} with range: 0..4294967295.
    Cfcrequestsbumped interface{}
}

func (cfcrequest *CISCOFTPCLIENTMIB_Cfcrequest) GetEntityData() *types.CommonEntityData {
    cfcrequest.EntityData.YFilter = cfcrequest.YFilter
    cfcrequest.EntityData.YangName = "cfcRequest"
    cfcrequest.EntityData.BundleName = "cisco_ios_xe"
    cfcrequest.EntityData.ParentYangName = "CISCO-FTP-CLIENT-MIB"
    cfcrequest.EntityData.SegmentPath = "cfcRequest"
    cfcrequest.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cfcrequest.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cfcrequest.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cfcrequest.EntityData.Children = make(map[string]types.YChild)
    cfcrequest.EntityData.Leafs = make(map[string]types.YLeaf)
    cfcrequest.EntityData.Leafs["cfcRequestMaximum"] = types.YLeaf{"Cfcrequestmaximum", cfcrequest.Cfcrequestmaximum}
    cfcrequest.EntityData.Leafs["cfcRequests"] = types.YLeaf{"Cfcrequests", cfcrequest.Cfcrequests}
    cfcrequest.EntityData.Leafs["cfcRequestsHigh"] = types.YLeaf{"Cfcrequestshigh", cfcrequest.Cfcrequestshigh}
    cfcrequest.EntityData.Leafs["cfcRequestsBumped"] = types.YLeaf{"Cfcrequestsbumped", cfcrequest.Cfcrequestsbumped}
    return &(cfcrequest.EntityData)
}

// CISCOFTPCLIENTMIB_Cfcrequesttable
// A table of FTP client requests.
type CISCOFTPCLIENTMIB_Cfcrequesttable struct {
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
    // The type is slice of CISCOFTPCLIENTMIB_Cfcrequesttable_Cfcrequestentry.
    Cfcrequestentry []CISCOFTPCLIENTMIB_Cfcrequesttable_Cfcrequestentry
}

func (cfcrequesttable *CISCOFTPCLIENTMIB_Cfcrequesttable) GetEntityData() *types.CommonEntityData {
    cfcrequesttable.EntityData.YFilter = cfcrequesttable.YFilter
    cfcrequesttable.EntityData.YangName = "cfcRequestTable"
    cfcrequesttable.EntityData.BundleName = "cisco_ios_xe"
    cfcrequesttable.EntityData.ParentYangName = "CISCO-FTP-CLIENT-MIB"
    cfcrequesttable.EntityData.SegmentPath = "cfcRequestTable"
    cfcrequesttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cfcrequesttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cfcrequesttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cfcrequesttable.EntityData.Children = make(map[string]types.YChild)
    cfcrequesttable.EntityData.Children["cfcRequestEntry"] = types.YChild{"Cfcrequestentry", nil}
    for i := range cfcrequesttable.Cfcrequestentry {
        cfcrequesttable.EntityData.Children[types.GetSegmentPath(&cfcrequesttable.Cfcrequestentry[i])] = types.YChild{"Cfcrequestentry", &cfcrequesttable.Cfcrequestentry[i]}
    }
    cfcrequesttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cfcrequesttable.EntityData)
}

// CISCOFTPCLIENTMIB_Cfcrequesttable_Cfcrequestentry
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
type CISCOFTPCLIENTMIB_Cfcrequesttable_Cfcrequestentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An arbitrary integer to uniquely identify this
    // entry.  To create an entry a management application should pick a random
    // number. The type is interface{} with range: 1..4294967295.
    Cfcrequestindex interface{}

    // The FTP operation to be performed. The type is Cfcrequestoperation.
    Cfcrequestoperation interface{}

    // The local file on which the operation is to be performed. The type is
    // string with length: 1..255.
    Cfcrequestlocalfile interface{}

    // The remote file on which the operation is to be performed. The type is
    // string with length: 1..255.
    Cfcrequestremotefile interface{}

    // The domain name or IP address of the FTP server to use. The type is string
    // with length: 1..64.
    Cfcrequestserver interface{}

    // The user name to use at the FTP server. The type is string with length:
    // 1..32.
    Cfcrequestuser interface{}

    // The password to use at the FTP server.  When read this object always
    // returns a zero-length string. The type is string with length: 0..16.
    Cfcrequestpassword interface{}

    // The result of the FTP operation. The type is Cfcrequestresult.
    Cfcrequestresult interface{}

    // The value of sysUpTime when the operation completed.  For an incomplete
    // operation this value is zero. The type is interface{} with range:
    // 0..4294967295.
    Cfcrequestcompletiontime interface{}

    // The action control to stop a running request.  Setting this to 'stop' will
    // begin the process of stopping the request.  Setting it to 'ready' or
    // setting it to 'stop' more than once have no effect.  When read this object
    // always returns ready. The type is Cfcrequeststop.
    Cfcrequeststop interface{}

    // The operational state of the file transfer.  To short-terminate the
    // transfer set cfcRequestStop to 'stop'. The type is
    // Cfcrequestoperationstate.
    Cfcrequestoperationstate interface{}

    // The control that allows modification, creation, and deletion of entries. 
    // For detailed rules see the DESCRIPTION for cfcRequestEntry. The type is
    // RowStatus.
    Cfcrequestentrystatus interface{}
}

func (cfcrequestentry *CISCOFTPCLIENTMIB_Cfcrequesttable_Cfcrequestentry) GetEntityData() *types.CommonEntityData {
    cfcrequestentry.EntityData.YFilter = cfcrequestentry.YFilter
    cfcrequestentry.EntityData.YangName = "cfcRequestEntry"
    cfcrequestentry.EntityData.BundleName = "cisco_ios_xe"
    cfcrequestentry.EntityData.ParentYangName = "cfcRequestTable"
    cfcrequestentry.EntityData.SegmentPath = "cfcRequestEntry" + "[cfcRequestIndex='" + fmt.Sprintf("%v", cfcrequestentry.Cfcrequestindex) + "']"
    cfcrequestentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cfcrequestentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cfcrequestentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cfcrequestentry.EntityData.Children = make(map[string]types.YChild)
    cfcrequestentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cfcrequestentry.EntityData.Leafs["cfcRequestIndex"] = types.YLeaf{"Cfcrequestindex", cfcrequestentry.Cfcrequestindex}
    cfcrequestentry.EntityData.Leafs["cfcRequestOperation"] = types.YLeaf{"Cfcrequestoperation", cfcrequestentry.Cfcrequestoperation}
    cfcrequestentry.EntityData.Leafs["cfcRequestLocalFile"] = types.YLeaf{"Cfcrequestlocalfile", cfcrequestentry.Cfcrequestlocalfile}
    cfcrequestentry.EntityData.Leafs["cfcRequestRemoteFile"] = types.YLeaf{"Cfcrequestremotefile", cfcrequestentry.Cfcrequestremotefile}
    cfcrequestentry.EntityData.Leafs["cfcRequestServer"] = types.YLeaf{"Cfcrequestserver", cfcrequestentry.Cfcrequestserver}
    cfcrequestentry.EntityData.Leafs["cfcRequestUser"] = types.YLeaf{"Cfcrequestuser", cfcrequestentry.Cfcrequestuser}
    cfcrequestentry.EntityData.Leafs["cfcRequestPassword"] = types.YLeaf{"Cfcrequestpassword", cfcrequestentry.Cfcrequestpassword}
    cfcrequestentry.EntityData.Leafs["cfcRequestResult"] = types.YLeaf{"Cfcrequestresult", cfcrequestentry.Cfcrequestresult}
    cfcrequestentry.EntityData.Leafs["cfcRequestCompletionTime"] = types.YLeaf{"Cfcrequestcompletiontime", cfcrequestentry.Cfcrequestcompletiontime}
    cfcrequestentry.EntityData.Leafs["cfcRequestStop"] = types.YLeaf{"Cfcrequeststop", cfcrequestentry.Cfcrequeststop}
    cfcrequestentry.EntityData.Leafs["cfcRequestOperationState"] = types.YLeaf{"Cfcrequestoperationstate", cfcrequestentry.Cfcrequestoperationstate}
    cfcrequestentry.EntityData.Leafs["cfcRequestEntryStatus"] = types.YLeaf{"Cfcrequestentrystatus", cfcrequestentry.Cfcrequestentrystatus}
    return &(cfcrequestentry.EntityData)
}

// CISCOFTPCLIENTMIB_Cfcrequesttable_Cfcrequestentry_Cfcrequestoperation represents The FTP operation to be performed.
type CISCOFTPCLIENTMIB_Cfcrequesttable_Cfcrequestentry_Cfcrequestoperation string

const (
    CISCOFTPCLIENTMIB_Cfcrequesttable_Cfcrequestentry_Cfcrequestoperation_putBinary CISCOFTPCLIENTMIB_Cfcrequesttable_Cfcrequestentry_Cfcrequestoperation = "putBinary"

    CISCOFTPCLIENTMIB_Cfcrequesttable_Cfcrequestentry_Cfcrequestoperation_putASCII CISCOFTPCLIENTMIB_Cfcrequesttable_Cfcrequestentry_Cfcrequestoperation = "putASCII"
)

// CISCOFTPCLIENTMIB_Cfcrequesttable_Cfcrequestentry_Cfcrequestoperationstate represents the transfer set cfcRequestStop to 'stop'.
type CISCOFTPCLIENTMIB_Cfcrequesttable_Cfcrequestentry_Cfcrequestoperationstate string

const (
    CISCOFTPCLIENTMIB_Cfcrequesttable_Cfcrequestentry_Cfcrequestoperationstate_running CISCOFTPCLIENTMIB_Cfcrequesttable_Cfcrequestentry_Cfcrequestoperationstate = "running"

    CISCOFTPCLIENTMIB_Cfcrequesttable_Cfcrequestentry_Cfcrequestoperationstate_stopping CISCOFTPCLIENTMIB_Cfcrequesttable_Cfcrequestentry_Cfcrequestoperationstate = "stopping"

    CISCOFTPCLIENTMIB_Cfcrequesttable_Cfcrequestentry_Cfcrequestoperationstate_stopped CISCOFTPCLIENTMIB_Cfcrequesttable_Cfcrequestentry_Cfcrequestoperationstate = "stopped"
)

// CISCOFTPCLIENTMIB_Cfcrequesttable_Cfcrequestentry_Cfcrequestresult represents The result of the FTP operation.
type CISCOFTPCLIENTMIB_Cfcrequesttable_Cfcrequestentry_Cfcrequestresult string

const (
    CISCOFTPCLIENTMIB_Cfcrequesttable_Cfcrequestentry_Cfcrequestresult_pending CISCOFTPCLIENTMIB_Cfcrequesttable_Cfcrequestentry_Cfcrequestresult = "pending"

    CISCOFTPCLIENTMIB_Cfcrequesttable_Cfcrequestentry_Cfcrequestresult_success CISCOFTPCLIENTMIB_Cfcrequesttable_Cfcrequestentry_Cfcrequestresult = "success"

    CISCOFTPCLIENTMIB_Cfcrequesttable_Cfcrequestentry_Cfcrequestresult_aborted CISCOFTPCLIENTMIB_Cfcrequesttable_Cfcrequestentry_Cfcrequestresult = "aborted"

    CISCOFTPCLIENTMIB_Cfcrequesttable_Cfcrequestentry_Cfcrequestresult_fileOpenFailLocal CISCOFTPCLIENTMIB_Cfcrequesttable_Cfcrequestentry_Cfcrequestresult = "fileOpenFailLocal"

    CISCOFTPCLIENTMIB_Cfcrequesttable_Cfcrequestentry_Cfcrequestresult_fileOpenFailRemote CISCOFTPCLIENTMIB_Cfcrequesttable_Cfcrequestentry_Cfcrequestresult = "fileOpenFailRemote"

    CISCOFTPCLIENTMIB_Cfcrequesttable_Cfcrequestentry_Cfcrequestresult_badDomainName CISCOFTPCLIENTMIB_Cfcrequesttable_Cfcrequestentry_Cfcrequestresult = "badDomainName"

    CISCOFTPCLIENTMIB_Cfcrequesttable_Cfcrequestentry_Cfcrequestresult_unreachableIpAddress CISCOFTPCLIENTMIB_Cfcrequesttable_Cfcrequestentry_Cfcrequestresult = "unreachableIpAddress"

    CISCOFTPCLIENTMIB_Cfcrequesttable_Cfcrequestentry_Cfcrequestresult_linkFailed CISCOFTPCLIENTMIB_Cfcrequesttable_Cfcrequestentry_Cfcrequestresult = "linkFailed"

    CISCOFTPCLIENTMIB_Cfcrequesttable_Cfcrequestentry_Cfcrequestresult_fileReadFailed CISCOFTPCLIENTMIB_Cfcrequesttable_Cfcrequestentry_Cfcrequestresult = "fileReadFailed"

    CISCOFTPCLIENTMIB_Cfcrequesttable_Cfcrequestentry_Cfcrequestresult_fileWriteFailed CISCOFTPCLIENTMIB_Cfcrequesttable_Cfcrequestentry_Cfcrequestresult = "fileWriteFailed"
)

// CISCOFTPCLIENTMIB_Cfcrequesttable_Cfcrequestentry_Cfcrequeststop represents effect.  When read this object always returns ready.
type CISCOFTPCLIENTMIB_Cfcrequesttable_Cfcrequestentry_Cfcrequeststop string

const (
    CISCOFTPCLIENTMIB_Cfcrequesttable_Cfcrequestentry_Cfcrequeststop_ready CISCOFTPCLIENTMIB_Cfcrequesttable_Cfcrequestentry_Cfcrequeststop = "ready"

    CISCOFTPCLIENTMIB_Cfcrequesttable_Cfcrequestentry_Cfcrequeststop_stop CISCOFTPCLIENTMIB_Cfcrequesttable_Cfcrequestentry_Cfcrequeststop = "stop"
)

