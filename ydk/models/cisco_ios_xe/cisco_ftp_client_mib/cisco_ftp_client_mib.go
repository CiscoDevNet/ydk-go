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
    parent types.Entity
    YFilter yfilter.YFilter

    
    Cfcrequest CISCOFTPCLIENTMIB_Cfcrequest

    // A table of FTP client requests.
    Cfcrequesttable CISCOFTPCLIENTMIB_Cfcrequesttable
}

func (cISCOFTPCLIENTMIB *CISCOFTPCLIENTMIB) GetFilter() yfilter.YFilter { return cISCOFTPCLIENTMIB.YFilter }

func (cISCOFTPCLIENTMIB *CISCOFTPCLIENTMIB) SetFilter(yf yfilter.YFilter) { cISCOFTPCLIENTMIB.YFilter = yf }

func (cISCOFTPCLIENTMIB *CISCOFTPCLIENTMIB) GetGoName(yname string) string {
    if yname == "cfcRequest" { return "Cfcrequest" }
    if yname == "cfcRequestTable" { return "Cfcrequesttable" }
    return ""
}

func (cISCOFTPCLIENTMIB *CISCOFTPCLIENTMIB) GetSegmentPath() string {
    return "CISCO-FTP-CLIENT-MIB:CISCO-FTP-CLIENT-MIB"
}

func (cISCOFTPCLIENTMIB *CISCOFTPCLIENTMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cfcRequest" {
        return &cISCOFTPCLIENTMIB.Cfcrequest
    }
    if childYangName == "cfcRequestTable" {
        return &cISCOFTPCLIENTMIB.Cfcrequesttable
    }
    return nil
}

func (cISCOFTPCLIENTMIB *CISCOFTPCLIENTMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cfcRequest"] = &cISCOFTPCLIENTMIB.Cfcrequest
    children["cfcRequestTable"] = &cISCOFTPCLIENTMIB.Cfcrequesttable
    return children
}

func (cISCOFTPCLIENTMIB *CISCOFTPCLIENTMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOFTPCLIENTMIB *CISCOFTPCLIENTMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOFTPCLIENTMIB *CISCOFTPCLIENTMIB) GetYangName() string { return "CISCO-FTP-CLIENT-MIB" }

func (cISCOFTPCLIENTMIB *CISCOFTPCLIENTMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOFTPCLIENTMIB *CISCOFTPCLIENTMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOFTPCLIENTMIB *CISCOFTPCLIENTMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOFTPCLIENTMIB *CISCOFTPCLIENTMIB) SetParent(parent types.Entity) { cISCOFTPCLIENTMIB.parent = parent }

func (cISCOFTPCLIENTMIB *CISCOFTPCLIENTMIB) GetParent() types.Entity { return cISCOFTPCLIENTMIB.parent }

func (cISCOFTPCLIENTMIB *CISCOFTPCLIENTMIB) GetParentYangName() string { return "CISCO-FTP-CLIENT-MIB" }

// CISCOFTPCLIENTMIB_Cfcrequest
type CISCOFTPCLIENTMIB_Cfcrequest struct {
    parent types.Entity
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

func (cfcrequest *CISCOFTPCLIENTMIB_Cfcrequest) GetFilter() yfilter.YFilter { return cfcrequest.YFilter }

func (cfcrequest *CISCOFTPCLIENTMIB_Cfcrequest) SetFilter(yf yfilter.YFilter) { cfcrequest.YFilter = yf }

func (cfcrequest *CISCOFTPCLIENTMIB_Cfcrequest) GetGoName(yname string) string {
    if yname == "cfcRequestMaximum" { return "Cfcrequestmaximum" }
    if yname == "cfcRequests" { return "Cfcrequests" }
    if yname == "cfcRequestsHigh" { return "Cfcrequestshigh" }
    if yname == "cfcRequestsBumped" { return "Cfcrequestsbumped" }
    return ""
}

func (cfcrequest *CISCOFTPCLIENTMIB_Cfcrequest) GetSegmentPath() string {
    return "cfcRequest"
}

func (cfcrequest *CISCOFTPCLIENTMIB_Cfcrequest) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cfcrequest *CISCOFTPCLIENTMIB_Cfcrequest) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cfcrequest *CISCOFTPCLIENTMIB_Cfcrequest) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cfcRequestMaximum"] = cfcrequest.Cfcrequestmaximum
    leafs["cfcRequests"] = cfcrequest.Cfcrequests
    leafs["cfcRequestsHigh"] = cfcrequest.Cfcrequestshigh
    leafs["cfcRequestsBumped"] = cfcrequest.Cfcrequestsbumped
    return leafs
}

func (cfcrequest *CISCOFTPCLIENTMIB_Cfcrequest) GetBundleName() string { return "cisco_ios_xe" }

func (cfcrequest *CISCOFTPCLIENTMIB_Cfcrequest) GetYangName() string { return "cfcRequest" }

func (cfcrequest *CISCOFTPCLIENTMIB_Cfcrequest) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cfcrequest *CISCOFTPCLIENTMIB_Cfcrequest) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cfcrequest *CISCOFTPCLIENTMIB_Cfcrequest) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cfcrequest *CISCOFTPCLIENTMIB_Cfcrequest) SetParent(parent types.Entity) { cfcrequest.parent = parent }

func (cfcrequest *CISCOFTPCLIENTMIB_Cfcrequest) GetParent() types.Entity { return cfcrequest.parent }

func (cfcrequest *CISCOFTPCLIENTMIB_Cfcrequest) GetParentYangName() string { return "CISCO-FTP-CLIENT-MIB" }

// CISCOFTPCLIENTMIB_Cfcrequesttable
// A table of FTP client requests.
type CISCOFTPCLIENTMIB_Cfcrequesttable struct {
    parent types.Entity
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

func (cfcrequesttable *CISCOFTPCLIENTMIB_Cfcrequesttable) GetFilter() yfilter.YFilter { return cfcrequesttable.YFilter }

func (cfcrequesttable *CISCOFTPCLIENTMIB_Cfcrequesttable) SetFilter(yf yfilter.YFilter) { cfcrequesttable.YFilter = yf }

func (cfcrequesttable *CISCOFTPCLIENTMIB_Cfcrequesttable) GetGoName(yname string) string {
    if yname == "cfcRequestEntry" { return "Cfcrequestentry" }
    return ""
}

func (cfcrequesttable *CISCOFTPCLIENTMIB_Cfcrequesttable) GetSegmentPath() string {
    return "cfcRequestTable"
}

func (cfcrequesttable *CISCOFTPCLIENTMIB_Cfcrequesttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cfcRequestEntry" {
        for _, c := range cfcrequesttable.Cfcrequestentry {
            if cfcrequesttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOFTPCLIENTMIB_Cfcrequesttable_Cfcrequestentry{}
        cfcrequesttable.Cfcrequestentry = append(cfcrequesttable.Cfcrequestentry, child)
        return &cfcrequesttable.Cfcrequestentry[len(cfcrequesttable.Cfcrequestentry)-1]
    }
    return nil
}

func (cfcrequesttable *CISCOFTPCLIENTMIB_Cfcrequesttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cfcrequesttable.Cfcrequestentry {
        children[cfcrequesttable.Cfcrequestentry[i].GetSegmentPath()] = &cfcrequesttable.Cfcrequestentry[i]
    }
    return children
}

func (cfcrequesttable *CISCOFTPCLIENTMIB_Cfcrequesttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cfcrequesttable *CISCOFTPCLIENTMIB_Cfcrequesttable) GetBundleName() string { return "cisco_ios_xe" }

func (cfcrequesttable *CISCOFTPCLIENTMIB_Cfcrequesttable) GetYangName() string { return "cfcRequestTable" }

func (cfcrequesttable *CISCOFTPCLIENTMIB_Cfcrequesttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cfcrequesttable *CISCOFTPCLIENTMIB_Cfcrequesttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cfcrequesttable *CISCOFTPCLIENTMIB_Cfcrequesttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cfcrequesttable *CISCOFTPCLIENTMIB_Cfcrequesttable) SetParent(parent types.Entity) { cfcrequesttable.parent = parent }

func (cfcrequesttable *CISCOFTPCLIENTMIB_Cfcrequesttable) GetParent() types.Entity { return cfcrequesttable.parent }

func (cfcrequesttable *CISCOFTPCLIENTMIB_Cfcrequesttable) GetParentYangName() string { return "CISCO-FTP-CLIENT-MIB" }

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
    parent types.Entity
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

func (cfcrequestentry *CISCOFTPCLIENTMIB_Cfcrequesttable_Cfcrequestentry) GetFilter() yfilter.YFilter { return cfcrequestentry.YFilter }

func (cfcrequestentry *CISCOFTPCLIENTMIB_Cfcrequesttable_Cfcrequestentry) SetFilter(yf yfilter.YFilter) { cfcrequestentry.YFilter = yf }

func (cfcrequestentry *CISCOFTPCLIENTMIB_Cfcrequesttable_Cfcrequestentry) GetGoName(yname string) string {
    if yname == "cfcRequestIndex" { return "Cfcrequestindex" }
    if yname == "cfcRequestOperation" { return "Cfcrequestoperation" }
    if yname == "cfcRequestLocalFile" { return "Cfcrequestlocalfile" }
    if yname == "cfcRequestRemoteFile" { return "Cfcrequestremotefile" }
    if yname == "cfcRequestServer" { return "Cfcrequestserver" }
    if yname == "cfcRequestUser" { return "Cfcrequestuser" }
    if yname == "cfcRequestPassword" { return "Cfcrequestpassword" }
    if yname == "cfcRequestResult" { return "Cfcrequestresult" }
    if yname == "cfcRequestCompletionTime" { return "Cfcrequestcompletiontime" }
    if yname == "cfcRequestStop" { return "Cfcrequeststop" }
    if yname == "cfcRequestOperationState" { return "Cfcrequestoperationstate" }
    if yname == "cfcRequestEntryStatus" { return "Cfcrequestentrystatus" }
    return ""
}

func (cfcrequestentry *CISCOFTPCLIENTMIB_Cfcrequesttable_Cfcrequestentry) GetSegmentPath() string {
    return "cfcRequestEntry" + "[cfcRequestIndex='" + fmt.Sprintf("%v", cfcrequestentry.Cfcrequestindex) + "']"
}

func (cfcrequestentry *CISCOFTPCLIENTMIB_Cfcrequesttable_Cfcrequestentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cfcrequestentry *CISCOFTPCLIENTMIB_Cfcrequesttable_Cfcrequestentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cfcrequestentry *CISCOFTPCLIENTMIB_Cfcrequesttable_Cfcrequestentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cfcRequestIndex"] = cfcrequestentry.Cfcrequestindex
    leafs["cfcRequestOperation"] = cfcrequestentry.Cfcrequestoperation
    leafs["cfcRequestLocalFile"] = cfcrequestentry.Cfcrequestlocalfile
    leafs["cfcRequestRemoteFile"] = cfcrequestentry.Cfcrequestremotefile
    leafs["cfcRequestServer"] = cfcrequestentry.Cfcrequestserver
    leafs["cfcRequestUser"] = cfcrequestentry.Cfcrequestuser
    leafs["cfcRequestPassword"] = cfcrequestentry.Cfcrequestpassword
    leafs["cfcRequestResult"] = cfcrequestentry.Cfcrequestresult
    leafs["cfcRequestCompletionTime"] = cfcrequestentry.Cfcrequestcompletiontime
    leafs["cfcRequestStop"] = cfcrequestentry.Cfcrequeststop
    leafs["cfcRequestOperationState"] = cfcrequestentry.Cfcrequestoperationstate
    leafs["cfcRequestEntryStatus"] = cfcrequestentry.Cfcrequestentrystatus
    return leafs
}

func (cfcrequestentry *CISCOFTPCLIENTMIB_Cfcrequesttable_Cfcrequestentry) GetBundleName() string { return "cisco_ios_xe" }

func (cfcrequestentry *CISCOFTPCLIENTMIB_Cfcrequesttable_Cfcrequestentry) GetYangName() string { return "cfcRequestEntry" }

func (cfcrequestentry *CISCOFTPCLIENTMIB_Cfcrequesttable_Cfcrequestentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cfcrequestentry *CISCOFTPCLIENTMIB_Cfcrequesttable_Cfcrequestentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cfcrequestentry *CISCOFTPCLIENTMIB_Cfcrequesttable_Cfcrequestentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cfcrequestentry *CISCOFTPCLIENTMIB_Cfcrequesttable_Cfcrequestentry) SetParent(parent types.Entity) { cfcrequestentry.parent = parent }

func (cfcrequestentry *CISCOFTPCLIENTMIB_Cfcrequesttable_Cfcrequestentry) GetParent() types.Entity { return cfcrequestentry.parent }

func (cfcrequestentry *CISCOFTPCLIENTMIB_Cfcrequesttable_Cfcrequestentry) GetParentYangName() string { return "cfcRequestTable" }

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

