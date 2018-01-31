// This MIB facilitates writing of configuration files
// of an SNMP Agent running Cisco's IOS in the 
// following ways: to and from the net, copying running 
// configurations to startup configurations and 
// vice-versa, and copying a configuration
// (running or startup) to and from the local 
// IOS file system.
package cisco_config_copy_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_config_copy_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-CONFIG-COPY-MIB CISCO-CONFIG-COPY-MIB}", reflect.TypeOf(CISCOCONFIGCOPYMIB{}))
    ydk.RegisterEntity("CISCO-CONFIG-COPY-MIB:CISCO-CONFIG-COPY-MIB", reflect.TypeOf(CISCOCONFIGCOPYMIB{}))
}

// ConfigCopyFailCause represents requestAborted: config copy operation aborted.
type ConfigCopyFailCause string

const (
    ConfigCopyFailCause_unknown ConfigCopyFailCause = "unknown"

    ConfigCopyFailCause_badFileName ConfigCopyFailCause = "badFileName"

    ConfigCopyFailCause_timeout ConfigCopyFailCause = "timeout"

    ConfigCopyFailCause_noMem ConfigCopyFailCause = "noMem"

    ConfigCopyFailCause_noConfig ConfigCopyFailCause = "noConfig"

    ConfigCopyFailCause_unsupportedProtocol ConfigCopyFailCause = "unsupportedProtocol"

    ConfigCopyFailCause_someConfigApplyFailed ConfigCopyFailCause = "someConfigApplyFailed"

    ConfigCopyFailCause_systemNotReady ConfigCopyFailCause = "systemNotReady"

    ConfigCopyFailCause_requestAborted ConfigCopyFailCause = "requestAborted"
)

// ConfigCopyState represents              unsuccessful.
type ConfigCopyState string

const (
    ConfigCopyState_waiting ConfigCopyState = "waiting"

    ConfigCopyState_running ConfigCopyState = "running"

    ConfigCopyState_successful ConfigCopyState = "successful"

    ConfigCopyState_failed ConfigCopyState = "failed"
)

// ConfigFileType represents                        or even a MAC-based fabric.
type ConfigFileType string

const (
    ConfigFileType_networkFile ConfigFileType = "networkFile"

    ConfigFileType_iosFile ConfigFileType = "iosFile"

    ConfigFileType_startupConfig ConfigFileType = "startupConfig"

    ConfigFileType_runningConfig ConfigFileType = "runningConfig"

    ConfigFileType_terminal ConfigFileType = "terminal"

    ConfigFileType_fabricStartupConfig ConfigFileType = "fabricStartupConfig"
)

// ConfigCopyProtocol represents sftp:   Secure File Transfer Protocol
type ConfigCopyProtocol string

const (
    ConfigCopyProtocol_tftp ConfigCopyProtocol = "tftp"

    ConfigCopyProtocol_ftp ConfigCopyProtocol = "ftp"

    ConfigCopyProtocol_rcp ConfigCopyProtocol = "rcp"

    ConfigCopyProtocol_scp ConfigCopyProtocol = "scp"

    ConfigCopyProtocol_sftp ConfigCopyProtocol = "sftp"
)

// CISCOCONFIGCOPYMIB
type CISCOCONFIGCOPYMIB struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A table of config-copy requests.
    Cccopytable CISCOCONFIGCOPYMIB_Cccopytable

    // A table containing information about the failure cause of the config copy
    // operation. An entry is created only when the value of ccCopyState changes
    // to 'failed' for a config copy operation.  Not all combinations of
    // ccCopySourceFileType and ccCopyDestFileType need to be supported.  For
    // example, an implementation may choose to support only the following
    // combination: ccCopySourceFileType = 'runningConfig' ccCopyDestFileType =
    // 'fabricStartupConfig'.   In the case where a fabric wide config copy 
    // operation is being performed, for example by selecting ccCopyDestFileType
    // value to be 'fabricStartupConfig', it is possible that the fabric could
    // have more than one device. In such cases this table would have one entry
    // for each device in the fabric. In this case even if the  operation
    // succeeded in one device and failed in  another, the operation as such has
    // failed, so the global state  represented by ccCopyState 'failed', but for
    // the device on which it was success,  ccCopyErrorDescription would have the 
    // distinguished value, 'success'.   Once the config copy operation completes
    // and if an entry gets instantiated, the management station  should retrieve
    // the values of the status objects of  interest. Once an entry in ccCopyTable
    // is deleted by management station, all the corresponding entries with the
    // same ccCopyIndex in this table are also  deleted.   In order to prevent old
    // entries from clogging the  table, entries age out at the same time as the 
    // corresponding entry with same ccCopyIndex in  ccCopyTable ages out.
    Cccopyerrortable CISCOCONFIGCOPYMIB_Cccopyerrortable
}

func (cISCOCONFIGCOPYMIB *CISCOCONFIGCOPYMIB) GetFilter() yfilter.YFilter { return cISCOCONFIGCOPYMIB.YFilter }

func (cISCOCONFIGCOPYMIB *CISCOCONFIGCOPYMIB) SetFilter(yf yfilter.YFilter) { cISCOCONFIGCOPYMIB.YFilter = yf }

func (cISCOCONFIGCOPYMIB *CISCOCONFIGCOPYMIB) GetGoName(yname string) string {
    if yname == "ccCopyTable" { return "Cccopytable" }
    if yname == "ccCopyErrorTable" { return "Cccopyerrortable" }
    return ""
}

func (cISCOCONFIGCOPYMIB *CISCOCONFIGCOPYMIB) GetSegmentPath() string {
    return "CISCO-CONFIG-COPY-MIB:CISCO-CONFIG-COPY-MIB"
}

func (cISCOCONFIGCOPYMIB *CISCOCONFIGCOPYMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ccCopyTable" {
        return &cISCOCONFIGCOPYMIB.Cccopytable
    }
    if childYangName == "ccCopyErrorTable" {
        return &cISCOCONFIGCOPYMIB.Cccopyerrortable
    }
    return nil
}

func (cISCOCONFIGCOPYMIB *CISCOCONFIGCOPYMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ccCopyTable"] = &cISCOCONFIGCOPYMIB.Cccopytable
    children["ccCopyErrorTable"] = &cISCOCONFIGCOPYMIB.Cccopyerrortable
    return children
}

func (cISCOCONFIGCOPYMIB *CISCOCONFIGCOPYMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOCONFIGCOPYMIB *CISCOCONFIGCOPYMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOCONFIGCOPYMIB *CISCOCONFIGCOPYMIB) GetYangName() string { return "CISCO-CONFIG-COPY-MIB" }

func (cISCOCONFIGCOPYMIB *CISCOCONFIGCOPYMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOCONFIGCOPYMIB *CISCOCONFIGCOPYMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOCONFIGCOPYMIB *CISCOCONFIGCOPYMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOCONFIGCOPYMIB *CISCOCONFIGCOPYMIB) SetParent(parent types.Entity) { cISCOCONFIGCOPYMIB.parent = parent }

func (cISCOCONFIGCOPYMIB *CISCOCONFIGCOPYMIB) GetParent() types.Entity { return cISCOCONFIGCOPYMIB.parent }

func (cISCOCONFIGCOPYMIB *CISCOCONFIGCOPYMIB) GetParentYangName() string { return "CISCO-CONFIG-COPY-MIB" }

// CISCOCONFIGCOPYMIB_Cccopytable
// A table of config-copy requests.
type CISCOCONFIGCOPYMIB_Cccopytable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A config-copy request.  A management station wishing to create an entry 
    // should first generate a random serial number to be used as the index to
    // this sparse table. The station  should then create the associated instance
    // of the row status and row index objects.  It must also,  either in the same
    // or in successive PDUs, create an instance of ccCopySourceFileType and 
    // ccCopyDestFileType.  At least one of the file types defined in 
    // ConfigFileType TC must be an agent-config file type (i.e. 'startupConfig'
    // or 'runningConfig'). If one of the file types is a 'networkFile', a valid
    // ccCopyFileName and ccCopyServerAddressType and  ccCopyServerAddressRev1 or
    // just ccCopyServerAddress must be created as well. If ccCopyServerAddress is
    // created then ccCopyServerAddressRev1 will store the same IP address and
    // ccCopyServerAddressType will  take the value 'ipv4'.  For a file type of
    // 'iosFile', only a valid ccCopyFileName needs to be created as an  extra
    // parameter.  It should also modify the default values for the  other
    // configuration objects if the defaults are not appropriate.  Once the
    // appropriate instance of all the  configuration objects have been created,
    // either by an explicit SNMP set request or by default, the row  status
    // should be set to active to initiate the  request. Note that this entire
    // procedure may be  initiated via a single set request which specifies a row
    // status of createAndGo as well as specifies valid values for the
    // non-defaulted  configuration objects.  Once the config-copy request has
    // been created  (i.e. the ccCopyEntryRowStatus has been made  active), the
    // entry cannot be modified - the only  operation possible after this is to
    // delete the row.  Once the request completes, the management station  should
    // retrieve the values of the status objects of  interest, and should then
    // delete the entry.  In order to prevent old entries from clogging the 
    // table, entries will be aged out, but an entry will  ever be deleted within
    // 5 minutes of completing. The type is slice of
    // CISCOCONFIGCOPYMIB_Cccopytable_Cccopyentry.
    Cccopyentry []CISCOCONFIGCOPYMIB_Cccopytable_Cccopyentry
}

func (cccopytable *CISCOCONFIGCOPYMIB_Cccopytable) GetFilter() yfilter.YFilter { return cccopytable.YFilter }

func (cccopytable *CISCOCONFIGCOPYMIB_Cccopytable) SetFilter(yf yfilter.YFilter) { cccopytable.YFilter = yf }

func (cccopytable *CISCOCONFIGCOPYMIB_Cccopytable) GetGoName(yname string) string {
    if yname == "ccCopyEntry" { return "Cccopyentry" }
    return ""
}

func (cccopytable *CISCOCONFIGCOPYMIB_Cccopytable) GetSegmentPath() string {
    return "ccCopyTable"
}

func (cccopytable *CISCOCONFIGCOPYMIB_Cccopytable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ccCopyEntry" {
        for _, c := range cccopytable.Cccopyentry {
            if cccopytable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOCONFIGCOPYMIB_Cccopytable_Cccopyentry{}
        cccopytable.Cccopyentry = append(cccopytable.Cccopyentry, child)
        return &cccopytable.Cccopyentry[len(cccopytable.Cccopyentry)-1]
    }
    return nil
}

func (cccopytable *CISCOCONFIGCOPYMIB_Cccopytable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cccopytable.Cccopyentry {
        children[cccopytable.Cccopyentry[i].GetSegmentPath()] = &cccopytable.Cccopyentry[i]
    }
    return children
}

func (cccopytable *CISCOCONFIGCOPYMIB_Cccopytable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cccopytable *CISCOCONFIGCOPYMIB_Cccopytable) GetBundleName() string { return "cisco_ios_xe" }

func (cccopytable *CISCOCONFIGCOPYMIB_Cccopytable) GetYangName() string { return "ccCopyTable" }

func (cccopytable *CISCOCONFIGCOPYMIB_Cccopytable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cccopytable *CISCOCONFIGCOPYMIB_Cccopytable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cccopytable *CISCOCONFIGCOPYMIB_Cccopytable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cccopytable *CISCOCONFIGCOPYMIB_Cccopytable) SetParent(parent types.Entity) { cccopytable.parent = parent }

func (cccopytable *CISCOCONFIGCOPYMIB_Cccopytable) GetParent() types.Entity { return cccopytable.parent }

func (cccopytable *CISCOCONFIGCOPYMIB_Cccopytable) GetParentYangName() string { return "CISCO-CONFIG-COPY-MIB" }

// CISCOCONFIGCOPYMIB_Cccopytable_Cccopyentry
// A config-copy request.
// 
// A management station wishing to create an entry 
// should first generate a random serial number to be
// used as the index to this sparse table. The station 
// should then create the associated instance of the
// row status and row index objects.  It must also, 
// either in the same or in successive PDUs, create an
// instance of ccCopySourceFileType and 
// ccCopyDestFileType.
// 
// At least one of the file types defined in 
// ConfigFileType TC must be an agent-config file type
// (i.e. 'startupConfig' or 'runningConfig').
// If one of the file types is a 'networkFile', a valid
// ccCopyFileName and ccCopyServerAddressType and 
// ccCopyServerAddressRev1 or just ccCopyServerAddress
// must be created as well. If ccCopyServerAddress is
// created then ccCopyServerAddressRev1 will store the
// same IP address and ccCopyServerAddressType will 
// take the value 'ipv4'.
// 
// For a file type of 'iosFile', only
// a valid ccCopyFileName needs to be created as an 
// extra parameter.
// 
// It should also modify the default values for the 
// other configuration objects if the defaults are not
// appropriate.
// 
// Once the appropriate instance of all the 
// configuration objects have been created, either by
// an explicit SNMP set request or by default, the row 
// status should be set to active to initiate the 
// request. Note that this entire procedure may be 
// initiated via a single set request which specifies
// a row status of createAndGo as well as
// specifies valid values for the non-defaulted 
// configuration objects.
// 
// Once the config-copy request has been created 
// (i.e. the ccCopyEntryRowStatus has been made 
// active), the entry cannot be modified - the only 
// operation possible after this is to delete the row.
// 
// Once the request completes, the management station 
// should retrieve the values of the status objects of 
// interest, and should then delete the entry.  In
// order to prevent old entries from clogging the 
// table, entries will be aged out, but an entry will 
// ever be deleted within 5 minutes of completing.
type CISCOCONFIGCOPYMIB_Cccopytable_Cccopyentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Object which specifies a unique entry in the
    // ccCopyTable.  A management station wishing to initiate a config-copy
    // operation should use a random value for this object when creating or
    // modifying an instance of a ccCopyEntry. The RowStatus semantics of the
    // ccCopyEntryRowStatus object will prevent access conflicts. The type is
    // interface{} with range: 1..2147483647.
    Cccopyindex interface{}

    // The protocol to be used for any copy.  If the copy operation occurs locally
    // on the SNMP  agent (e.g. 'runningConfig' to 'startupConfig'), this object
    // may be ignored by the implementation. The type is ConfigCopyProtocol.
    Cccopyprotocol interface{}

    // Specifies the type of file to copy from. Either the ccCopySourceFileType or
    // the ccCopyDestFileType  (or both) must be of type 'runningConfig' or 
    // 'startupConfig'. Also, the ccCopySourceFileType must be different from the
    // ccCopyDestFileType.  If the ccCopySourceFileType has the value of 
    // 'networkFile', the ccCopyServerAddress/ ccCopyServerAddressRev1 and
    // ccCopyServerAddressType and ccCopyFileName must also be created, and 3 
    // objects together (ccCopySourceFileType, ccCopyServerAddressRev1,
    // ccCopyFileName) will  uniquely identify the source file. If 
    // ccCopyServerAddress is created then  ccCopyServerAddressRev1 will store the
    // same IP address and ccCopyServerAddressType will  take the value 'ipv4'.  
    // If the ccCopySourceFileType is 'iosFile', the  ccCopyFileName must also be
    // created, and the  2 objects together (ccCopySourceFileType, ccCopyFileName)
    // will uniquely identify the source  file. The type is ConfigFileType.
    Cccopysourcefiletype interface{}

    // specifies the type of file to copy to. Either the ccCopySourceFileType or
    // the ccCopyDestFileType  (or both) must be of type 'runningConfig' or 
    // 'startupConfig'. Also, the ccCopySourceFileType  must be different from the
    // ccCopyDestFileType.  If the ccCopyDestFileType has the value of 
    // 'networkFile', the  ccCopyServerAddress/ccCopyServerAddressType and
    // ccCopyServerAddressRev1 and ccCopyFileName must also be created, and 3
    // objects together (ccCopyDestFileType, ccCopyServerAddressRev1,  
    // ccCopyFileName) will uniquely identify the  destination file. If
    // ccCopyServerAddress is created then ccCopyServerAddressRev1 will store the
    // same IP address and ccCopyServerAddressType will take the  value 'ipv4'.  
    // If the ccCopyDestFileType is 'iosFile', the  ccCopyFileName must also be
    // created, and the 2 objects together (ccCopyDestFileType,  ccCopyFileName)
    // will uniquely identify the  destination file. The type is ConfigFileType.
    Cccopydestfiletype interface{}

    // The IP address of the TFTP server from (or to) which to copy the
    // configuration file. This object  must be created when either the 
    // ccCopySourceFileType or ccCopyDestFileType has the value 'networkFile'. 
    // Values of 0.0.0.0 or FF.FF.FF.FF for ccCopyServerAddress are not allowed. 
    // Since this object can just hold only IPv4 Transport type, it is deprecated
    // and replaced by  ccCopyServerAddressRev1. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Cccopyserveraddress interface{}

    // The file name (including the path, if applicable) of the file. This object
    // must be created when either the ccCopySourceFileType or ccCopyDestFileType
    // has the value 'networkFile' or 'iosFile'. The type is string.
    Cccopyfilename interface{}

    // Remote username for copy via FTP, RCP, SFTP or SCP protocol. This object
    // must be created when the ccCopyProtocol is 'rcp', 'scp', 'ftp', or 'sftp'.
    // If the protocol is RCP, it will override the remote  username configured
    // through the          rcmd remote-username <username> configuration command.
    // The remote username is sent as the server username  in an RCP command
    // request sent by the system to a remote RCP server. The type is string with
    // length: 1..40.
    Cccopyusername interface{}

    // Password used by FTP, SFTP or SCP for copying a file to/from an
    // FTP/SFTP/SCP server. This object  must be created when the ccCopyProtocol
    // is FTP or SCP.  Reading it returns a zero-length string for security 
    // reasons. The type is string with length: 1..40.
    Cccopyuserpassword interface{}

    // Specifies whether or not a ccCopyCompletion notification should be issued
    // on completion of the TFTP transfer. If such a notification is desired,  it
    // is the responsibility of the management entity to ensure that the SNMP
    // administrative model is  configured in such a way as to allow the 
    // notification to be delivered. The type is bool.
    Cccopynotificationoncompletion interface{}

    // Specifies the state of this config-copy request. This value of this object
    // is instantiated only after  the row has been instantiated, i.e. after the 
    // ccCopyEntryRowStatus has been made active. The type is ConfigCopyState.
    Cccopystate interface{}

    // Specifies the time the ccCopyState last transitioned to 'running', or 0 if
    // the state has  never transitioned to 'running'(e.g., stuck in 'waiting'
    // state).  This object is instantiated only after the row has  been
    // instantiated. The type is interface{} with range: 0..4294967295.
    Cccopytimestarted interface{}

    // Specifies the time the ccCopyState last transitioned from 'running' to
    // 'successful' or  'failed' states. This object is instantiated only  after
    // the row has been instantiated. Its value will remain 0 until the request
    // has  completed. The type is interface{} with range: 0..4294967295.
    Cccopytimecompleted interface{}

    // The reason why the config-copy operation failed. This object is
    // instantiated only when the  ccCopyState for this entry is in the  'failed'
    // state. The type is ConfigCopyFailCause.
    Cccopyfailcause interface{}

    // The status of this table entry. Once the entry status is set to active, the
    // associated entry cannot  be modified until the request completes 
    // (ccCopyState transitions to 'successful' or 'failed' state). The type is
    // RowStatus.
    Cccopyentryrowstatus interface{}

    // This object indicates the transport type of the address contained in
    // ccCopyServerAddressRev1 object.  This must be created when either the
    // ccCopySourceFileType or ccCopyDestFileType has the value 'networkFile'. The
    // type is InetAddressType.
    Cccopyserveraddresstype interface{}

    // The IP address of the TFTP server from (or to) which to copy the
    // configuration file. This object must be created when either the 
    // ccCopySourceFileType or ccCopyDestFileType has the value 'networkFile'.   
    // All bits as 0s or 1s for ccCopyServerAddressRev1 are not allowed.  The
    // format of this address depends on the value of  the ccCopyServerAddressType
    // object. The type is string with length: 0..255.
    Cccopyserveraddressrev1 interface{}
}

func (cccopyentry *CISCOCONFIGCOPYMIB_Cccopytable_Cccopyentry) GetFilter() yfilter.YFilter { return cccopyentry.YFilter }

func (cccopyentry *CISCOCONFIGCOPYMIB_Cccopytable_Cccopyentry) SetFilter(yf yfilter.YFilter) { cccopyentry.YFilter = yf }

func (cccopyentry *CISCOCONFIGCOPYMIB_Cccopytable_Cccopyentry) GetGoName(yname string) string {
    if yname == "ccCopyIndex" { return "Cccopyindex" }
    if yname == "ccCopyProtocol" { return "Cccopyprotocol" }
    if yname == "ccCopySourceFileType" { return "Cccopysourcefiletype" }
    if yname == "ccCopyDestFileType" { return "Cccopydestfiletype" }
    if yname == "ccCopyServerAddress" { return "Cccopyserveraddress" }
    if yname == "ccCopyFileName" { return "Cccopyfilename" }
    if yname == "ccCopyUserName" { return "Cccopyusername" }
    if yname == "ccCopyUserPassword" { return "Cccopyuserpassword" }
    if yname == "ccCopyNotificationOnCompletion" { return "Cccopynotificationoncompletion" }
    if yname == "ccCopyState" { return "Cccopystate" }
    if yname == "ccCopyTimeStarted" { return "Cccopytimestarted" }
    if yname == "ccCopyTimeCompleted" { return "Cccopytimecompleted" }
    if yname == "ccCopyFailCause" { return "Cccopyfailcause" }
    if yname == "ccCopyEntryRowStatus" { return "Cccopyentryrowstatus" }
    if yname == "ccCopyServerAddressType" { return "Cccopyserveraddresstype" }
    if yname == "ccCopyServerAddressRev1" { return "Cccopyserveraddressrev1" }
    return ""
}

func (cccopyentry *CISCOCONFIGCOPYMIB_Cccopytable_Cccopyentry) GetSegmentPath() string {
    return "ccCopyEntry" + "[ccCopyIndex='" + fmt.Sprintf("%v", cccopyentry.Cccopyindex) + "']"
}

func (cccopyentry *CISCOCONFIGCOPYMIB_Cccopytable_Cccopyentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cccopyentry *CISCOCONFIGCOPYMIB_Cccopytable_Cccopyentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cccopyentry *CISCOCONFIGCOPYMIB_Cccopytable_Cccopyentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ccCopyIndex"] = cccopyentry.Cccopyindex
    leafs["ccCopyProtocol"] = cccopyentry.Cccopyprotocol
    leafs["ccCopySourceFileType"] = cccopyentry.Cccopysourcefiletype
    leafs["ccCopyDestFileType"] = cccopyentry.Cccopydestfiletype
    leafs["ccCopyServerAddress"] = cccopyentry.Cccopyserveraddress
    leafs["ccCopyFileName"] = cccopyentry.Cccopyfilename
    leafs["ccCopyUserName"] = cccopyentry.Cccopyusername
    leafs["ccCopyUserPassword"] = cccopyentry.Cccopyuserpassword
    leafs["ccCopyNotificationOnCompletion"] = cccopyentry.Cccopynotificationoncompletion
    leafs["ccCopyState"] = cccopyentry.Cccopystate
    leafs["ccCopyTimeStarted"] = cccopyentry.Cccopytimestarted
    leafs["ccCopyTimeCompleted"] = cccopyentry.Cccopytimecompleted
    leafs["ccCopyFailCause"] = cccopyentry.Cccopyfailcause
    leafs["ccCopyEntryRowStatus"] = cccopyentry.Cccopyentryrowstatus
    leafs["ccCopyServerAddressType"] = cccopyentry.Cccopyserveraddresstype
    leafs["ccCopyServerAddressRev1"] = cccopyentry.Cccopyserveraddressrev1
    return leafs
}

func (cccopyentry *CISCOCONFIGCOPYMIB_Cccopytable_Cccopyentry) GetBundleName() string { return "cisco_ios_xe" }

func (cccopyentry *CISCOCONFIGCOPYMIB_Cccopytable_Cccopyentry) GetYangName() string { return "ccCopyEntry" }

func (cccopyentry *CISCOCONFIGCOPYMIB_Cccopytable_Cccopyentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cccopyentry *CISCOCONFIGCOPYMIB_Cccopytable_Cccopyentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cccopyentry *CISCOCONFIGCOPYMIB_Cccopytable_Cccopyentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cccopyentry *CISCOCONFIGCOPYMIB_Cccopytable_Cccopyentry) SetParent(parent types.Entity) { cccopyentry.parent = parent }

func (cccopyentry *CISCOCONFIGCOPYMIB_Cccopytable_Cccopyentry) GetParent() types.Entity { return cccopyentry.parent }

func (cccopyentry *CISCOCONFIGCOPYMIB_Cccopytable_Cccopyentry) GetParentYangName() string { return "ccCopyTable" }

// CISCOCONFIGCOPYMIB_Cccopyerrortable
// A table containing information about the failure
// cause of the config copy operation. An entry is
// created only when the value of ccCopyState changes
// to 'failed' for a config copy operation.
// 
// Not all combinations of ccCopySourceFileType and
// ccCopyDestFileType need to be supported.  For
// example, an implementation may choose to support
// only the following combination:
// ccCopySourceFileType = 'runningConfig'
// ccCopyDestFileType = 'fabricStartupConfig'. 
// 
// In the case where a fabric wide config copy 
// operation is being performed, for example by
// selecting ccCopyDestFileType value to be
// 'fabricStartupConfig', it is possible that the
// fabric could have more than one device. In such
// cases this table would have one entry for each
// device in the fabric. In this case even if the 
// operation succeeded in one device and failed in 
// another, the operation as such has failed, so the
// global state  represented by ccCopyState 'failed',
// but for the device on which it was success, 
// ccCopyErrorDescription would have the 
// distinguished value, 'success'. 
// 
// Once the config copy operation completes and if an
// entry gets instantiated, the management station 
// should retrieve the values of the status objects of 
// interest. Once an entry in ccCopyTable is deleted
// by management station, all the corresponding entries
// with the same ccCopyIndex in this table are also 
// deleted. 
// 
// In order to prevent old entries from clogging the 
// table, entries age out at the same time as the 
// corresponding entry with same ccCopyIndex in 
// ccCopyTable ages out.
type CISCOCONFIGCOPYMIB_Cccopyerrortable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry containing information about the outcome at one destination of a
    // failed config copy operation. The type is slice of
    // CISCOCONFIGCOPYMIB_Cccopyerrortable_Cccopyerrorentry.
    Cccopyerrorentry []CISCOCONFIGCOPYMIB_Cccopyerrortable_Cccopyerrorentry
}

func (cccopyerrortable *CISCOCONFIGCOPYMIB_Cccopyerrortable) GetFilter() yfilter.YFilter { return cccopyerrortable.YFilter }

func (cccopyerrortable *CISCOCONFIGCOPYMIB_Cccopyerrortable) SetFilter(yf yfilter.YFilter) { cccopyerrortable.YFilter = yf }

func (cccopyerrortable *CISCOCONFIGCOPYMIB_Cccopyerrortable) GetGoName(yname string) string {
    if yname == "ccCopyErrorEntry" { return "Cccopyerrorentry" }
    return ""
}

func (cccopyerrortable *CISCOCONFIGCOPYMIB_Cccopyerrortable) GetSegmentPath() string {
    return "ccCopyErrorTable"
}

func (cccopyerrortable *CISCOCONFIGCOPYMIB_Cccopyerrortable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ccCopyErrorEntry" {
        for _, c := range cccopyerrortable.Cccopyerrorentry {
            if cccopyerrortable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOCONFIGCOPYMIB_Cccopyerrortable_Cccopyerrorentry{}
        cccopyerrortable.Cccopyerrorentry = append(cccopyerrortable.Cccopyerrorentry, child)
        return &cccopyerrortable.Cccopyerrorentry[len(cccopyerrortable.Cccopyerrorentry)-1]
    }
    return nil
}

func (cccopyerrortable *CISCOCONFIGCOPYMIB_Cccopyerrortable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cccopyerrortable.Cccopyerrorentry {
        children[cccopyerrortable.Cccopyerrorentry[i].GetSegmentPath()] = &cccopyerrortable.Cccopyerrorentry[i]
    }
    return children
}

func (cccopyerrortable *CISCOCONFIGCOPYMIB_Cccopyerrortable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cccopyerrortable *CISCOCONFIGCOPYMIB_Cccopyerrortable) GetBundleName() string { return "cisco_ios_xe" }

func (cccopyerrortable *CISCOCONFIGCOPYMIB_Cccopyerrortable) GetYangName() string { return "ccCopyErrorTable" }

func (cccopyerrortable *CISCOCONFIGCOPYMIB_Cccopyerrortable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cccopyerrortable *CISCOCONFIGCOPYMIB_Cccopyerrortable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cccopyerrortable *CISCOCONFIGCOPYMIB_Cccopyerrortable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cccopyerrortable *CISCOCONFIGCOPYMIB_Cccopyerrortable) SetParent(parent types.Entity) { cccopyerrortable.parent = parent }

func (cccopyerrortable *CISCOCONFIGCOPYMIB_Cccopyerrortable) GetParent() types.Entity { return cccopyerrortable.parent }

func (cccopyerrortable *CISCOCONFIGCOPYMIB_Cccopyerrortable) GetParentYangName() string { return "CISCO-CONFIG-COPY-MIB" }

// CISCOCONFIGCOPYMIB_Cccopyerrortable_Cccopyerrorentry
// An entry containing information about the
// outcome at one destination of a failed config
// copy operation.
type CISCOCONFIGCOPYMIB_Cccopyerrortable_Cccopyerrorentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // cisco_config_copy_mib.CISCOCONFIGCOPYMIB_Cccopytable_Cccopyentry_Cccopyindex
    Cccopyindex interface{}

    // This attribute is a key. A monotonically increasing integer for the sole
    // purpose of indexing entries in this table. When a config copy operation has
    // multiple  destinations, then this index value is used to  distinguish
    // between those multiple destinations. The type is interface{} with range:
    // 0..4294967295.
    Cccopyerrorindex interface{}

    // The type of Internet address for this destination device on which config
    // copy operation is performed. The type is InetAddressType.
    Cccopyerrordeviceipaddresstype interface{}

    // The IP address of this destination device on which config copy operation is
    // performed. The object value has to be consistent with the type specified in
    // ccCopyErrorDeviceIpAddressType. The type is string with length: 0..255.
    Cccopyerrordeviceipaddress interface{}

    // The World Wide Name (WWN) of this destination device on which config copy
    // operation is performed. The value of this object is zero-length string if 
    // WWN is unassigned or unknown. For example, devices  which do not support
    // fibre channel would not have WWN. The type is string with length: 0 | 8 |
    // 16.
    Cccopyerrordevicewwn interface{}

    // The error description for the error happened for this destination of this
    // config copy  operation. The type is string.
    Cccopyerrordescription interface{}
}

func (cccopyerrorentry *CISCOCONFIGCOPYMIB_Cccopyerrortable_Cccopyerrorentry) GetFilter() yfilter.YFilter { return cccopyerrorentry.YFilter }

func (cccopyerrorentry *CISCOCONFIGCOPYMIB_Cccopyerrortable_Cccopyerrorentry) SetFilter(yf yfilter.YFilter) { cccopyerrorentry.YFilter = yf }

func (cccopyerrorentry *CISCOCONFIGCOPYMIB_Cccopyerrortable_Cccopyerrorentry) GetGoName(yname string) string {
    if yname == "ccCopyIndex" { return "Cccopyindex" }
    if yname == "ccCopyErrorIndex" { return "Cccopyerrorindex" }
    if yname == "ccCopyErrorDeviceIpAddressType" { return "Cccopyerrordeviceipaddresstype" }
    if yname == "ccCopyErrorDeviceIpAddress" { return "Cccopyerrordeviceipaddress" }
    if yname == "ccCopyErrorDeviceWWN" { return "Cccopyerrordevicewwn" }
    if yname == "ccCopyErrorDescription" { return "Cccopyerrordescription" }
    return ""
}

func (cccopyerrorentry *CISCOCONFIGCOPYMIB_Cccopyerrortable_Cccopyerrorentry) GetSegmentPath() string {
    return "ccCopyErrorEntry" + "[ccCopyIndex='" + fmt.Sprintf("%v", cccopyerrorentry.Cccopyindex) + "']" + "[ccCopyErrorIndex='" + fmt.Sprintf("%v", cccopyerrorentry.Cccopyerrorindex) + "']"
}

func (cccopyerrorentry *CISCOCONFIGCOPYMIB_Cccopyerrortable_Cccopyerrorentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cccopyerrorentry *CISCOCONFIGCOPYMIB_Cccopyerrortable_Cccopyerrorentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cccopyerrorentry *CISCOCONFIGCOPYMIB_Cccopyerrortable_Cccopyerrorentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ccCopyIndex"] = cccopyerrorentry.Cccopyindex
    leafs["ccCopyErrorIndex"] = cccopyerrorentry.Cccopyerrorindex
    leafs["ccCopyErrorDeviceIpAddressType"] = cccopyerrorentry.Cccopyerrordeviceipaddresstype
    leafs["ccCopyErrorDeviceIpAddress"] = cccopyerrorentry.Cccopyerrordeviceipaddress
    leafs["ccCopyErrorDeviceWWN"] = cccopyerrorentry.Cccopyerrordevicewwn
    leafs["ccCopyErrorDescription"] = cccopyerrorentry.Cccopyerrordescription
    return leafs
}

func (cccopyerrorentry *CISCOCONFIGCOPYMIB_Cccopyerrortable_Cccopyerrorentry) GetBundleName() string { return "cisco_ios_xe" }

func (cccopyerrorentry *CISCOCONFIGCOPYMIB_Cccopyerrortable_Cccopyerrorentry) GetYangName() string { return "ccCopyErrorEntry" }

func (cccopyerrorentry *CISCOCONFIGCOPYMIB_Cccopyerrortable_Cccopyerrorentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cccopyerrorentry *CISCOCONFIGCOPYMIB_Cccopyerrortable_Cccopyerrorentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cccopyerrorentry *CISCOCONFIGCOPYMIB_Cccopyerrortable_Cccopyerrorentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cccopyerrorentry *CISCOCONFIGCOPYMIB_Cccopyerrortable_Cccopyerrorentry) SetParent(parent types.Entity) { cccopyerrorentry.parent = parent }

func (cccopyerrorentry *CISCOCONFIGCOPYMIB_Cccopyerrortable_Cccopyerrorentry) GetParent() types.Entity { return cccopyerrorentry.parent }

func (cccopyerrorentry *CISCOCONFIGCOPYMIB_Cccopyerrortable_Cccopyerrorentry) GetParentYangName() string { return "ccCopyErrorTable" }

