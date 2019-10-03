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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A table of config-copy requests.
    CcCopyTable CISCOCONFIGCOPYMIB_CcCopyTable

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
    CcCopyErrorTable CISCOCONFIGCOPYMIB_CcCopyErrorTable
}

func (cISCOCONFIGCOPYMIB *CISCOCONFIGCOPYMIB) GetEntityData() *types.CommonEntityData {
    cISCOCONFIGCOPYMIB.EntityData.YFilter = cISCOCONFIGCOPYMIB.YFilter
    cISCOCONFIGCOPYMIB.EntityData.YangName = "CISCO-CONFIG-COPY-MIB"
    cISCOCONFIGCOPYMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOCONFIGCOPYMIB.EntityData.ParentYangName = "CISCO-CONFIG-COPY-MIB"
    cISCOCONFIGCOPYMIB.EntityData.SegmentPath = "CISCO-CONFIG-COPY-MIB:CISCO-CONFIG-COPY-MIB"
    cISCOCONFIGCOPYMIB.EntityData.AbsolutePath = cISCOCONFIGCOPYMIB.EntityData.SegmentPath
    cISCOCONFIGCOPYMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOCONFIGCOPYMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOCONFIGCOPYMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOCONFIGCOPYMIB.EntityData.Children = types.NewOrderedMap()
    cISCOCONFIGCOPYMIB.EntityData.Children.Append("ccCopyTable", types.YChild{"CcCopyTable", &cISCOCONFIGCOPYMIB.CcCopyTable})
    cISCOCONFIGCOPYMIB.EntityData.Children.Append("ccCopyErrorTable", types.YChild{"CcCopyErrorTable", &cISCOCONFIGCOPYMIB.CcCopyErrorTable})
    cISCOCONFIGCOPYMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOCONFIGCOPYMIB.EntityData.YListKeys = []string {}

    return &(cISCOCONFIGCOPYMIB.EntityData)
}

// CISCOCONFIGCOPYMIB_CcCopyTable
// A table of config-copy requests.
type CISCOCONFIGCOPYMIB_CcCopyTable struct {
    EntityData types.CommonEntityData
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
    // CISCOCONFIGCOPYMIB_CcCopyTable_CcCopyEntry.
    CcCopyEntry []*CISCOCONFIGCOPYMIB_CcCopyTable_CcCopyEntry
}

func (ccCopyTable *CISCOCONFIGCOPYMIB_CcCopyTable) GetEntityData() *types.CommonEntityData {
    ccCopyTable.EntityData.YFilter = ccCopyTable.YFilter
    ccCopyTable.EntityData.YangName = "ccCopyTable"
    ccCopyTable.EntityData.BundleName = "cisco_ios_xe"
    ccCopyTable.EntityData.ParentYangName = "CISCO-CONFIG-COPY-MIB"
    ccCopyTable.EntityData.SegmentPath = "ccCopyTable"
    ccCopyTable.EntityData.AbsolutePath = "CISCO-CONFIG-COPY-MIB:CISCO-CONFIG-COPY-MIB/" + ccCopyTable.EntityData.SegmentPath
    ccCopyTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccCopyTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccCopyTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccCopyTable.EntityData.Children = types.NewOrderedMap()
    ccCopyTable.EntityData.Children.Append("ccCopyEntry", types.YChild{"CcCopyEntry", nil})
    for i := range ccCopyTable.CcCopyEntry {
        ccCopyTable.EntityData.Children.Append(types.GetSegmentPath(ccCopyTable.CcCopyEntry[i]), types.YChild{"CcCopyEntry", ccCopyTable.CcCopyEntry[i]})
    }
    ccCopyTable.EntityData.Leafs = types.NewOrderedMap()

    ccCopyTable.EntityData.YListKeys = []string {}

    return &(ccCopyTable.EntityData)
}

// CISCOCONFIGCOPYMIB_CcCopyTable_CcCopyEntry
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
type CISCOCONFIGCOPYMIB_CcCopyTable_CcCopyEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Object which specifies a unique entry in the
    // ccCopyTable.  A management station wishing to initiate a config-copy
    // operation should use a random value for this object when creating or
    // modifying an instance of a ccCopyEntry. The RowStatus semantics of the
    // ccCopyEntryRowStatus object will prevent access conflicts. The type is
    // interface{} with range: 1..2147483647.
    CcCopyIndex interface{}

    // The protocol to be used for any copy.  If the copy operation occurs locally
    // on the SNMP  agent (e.g. 'runningConfig' to 'startupConfig'), this object
    // may be ignored by the implementation. The type is ConfigCopyProtocol.
    CcCopyProtocol interface{}

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
    CcCopySourceFileType interface{}

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
    CcCopyDestFileType interface{}

    // The IP address of the TFTP server from (or to) which to copy the
    // configuration file. This object  must be created when either the 
    // ccCopySourceFileType or ccCopyDestFileType has the value 'networkFile'. 
    // Values of 0.0.0.0 or FF.FF.FF.FF for ccCopyServerAddress are not allowed. 
    // Since this object can just hold only IPv4 Transport type, it is deprecated
    // and replaced by  ccCopyServerAddressRev1. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    CcCopyServerAddress interface{}

    // The file name (including the path, if applicable) of the file. This object
    // must be created when either the ccCopySourceFileType or ccCopyDestFileType
    // has the value 'networkFile' or 'iosFile'. The type is string.
    CcCopyFileName interface{}

    // Remote username for copy via FTP, RCP, SFTP or SCP protocol. This object
    // must be created when the ccCopyProtocol is 'rcp', 'scp', 'ftp', or 'sftp'.
    // If the protocol is RCP, it will override the remote  username configured
    // through the          rcmd remote-username <username> configuration command.
    // The remote username is sent as the server username  in an RCP command
    // request sent by the system to a remote RCP server. The type is string with
    // length: 1..40.
    CcCopyUserName interface{}

    // Password used by FTP, SFTP or SCP for copying a file to/from an
    // FTP/SFTP/SCP server. This object  must be created when the ccCopyProtocol
    // is FTP or SCP.  Reading it returns a zero-length string for security 
    // reasons. The type is string with length: 1..40.
    CcCopyUserPassword interface{}

    // Specifies whether or not a ccCopyCompletion notification should be issued
    // on completion of the TFTP transfer. If such a notification is desired,  it
    // is the responsibility of the management entity to ensure that the SNMP
    // administrative model is  configured in such a way as to allow the 
    // notification to be delivered. The type is bool.
    CcCopyNotificationOnCompletion interface{}

    // Specifies the state of this config-copy request. This value of this object
    // is instantiated only after  the row has been instantiated, i.e. after the 
    // ccCopyEntryRowStatus has been made active. The type is ConfigCopyState.
    CcCopyState interface{}

    // Specifies the time the ccCopyState last transitioned to 'running', or 0 if
    // the state has  never transitioned to 'running'(e.g., stuck in 'waiting'
    // state).  This object is instantiated only after the row has  been
    // instantiated. The type is interface{} with range: 0..4294967295.
    CcCopyTimeStarted interface{}

    // Specifies the time the ccCopyState last transitioned from 'running' to
    // 'successful' or  'failed' states. This object is instantiated only  after
    // the row has been instantiated. Its value will remain 0 until the request
    // has  completed. The type is interface{} with range: 0..4294967295.
    CcCopyTimeCompleted interface{}

    // The reason why the config-copy operation failed. This object is
    // instantiated only when the  ccCopyState for this entry is in the  'failed'
    // state. The type is ConfigCopyFailCause.
    CcCopyFailCause interface{}

    // The status of this table entry. Once the entry status is set to active, the
    // associated entry cannot  be modified until the request completes 
    // (ccCopyState transitions to 'successful' or 'failed' state). The type is
    // RowStatus.
    CcCopyEntryRowStatus interface{}

    // This object indicates the transport type of the address contained in
    // ccCopyServerAddressRev1 object.  This must be created when either the
    // ccCopySourceFileType or ccCopyDestFileType has the value 'networkFile'. The
    // type is InetAddressType.
    CcCopyServerAddressType interface{}

    // The IP address of the TFTP server from (or to) which to copy the
    // configuration file. This object must be created when either the 
    // ccCopySourceFileType or ccCopyDestFileType has the value 'networkFile'.   
    // All bits as 0s or 1s for ccCopyServerAddressRev1 are not allowed.  The
    // format of this address depends on the value of  the ccCopyServerAddressType
    // object. The type is string with length: 0..255.
    CcCopyServerAddressRev1 interface{}
}

func (ccCopyEntry *CISCOCONFIGCOPYMIB_CcCopyTable_CcCopyEntry) GetEntityData() *types.CommonEntityData {
    ccCopyEntry.EntityData.YFilter = ccCopyEntry.YFilter
    ccCopyEntry.EntityData.YangName = "ccCopyEntry"
    ccCopyEntry.EntityData.BundleName = "cisco_ios_xe"
    ccCopyEntry.EntityData.ParentYangName = "ccCopyTable"
    ccCopyEntry.EntityData.SegmentPath = "ccCopyEntry" + types.AddKeyToken(ccCopyEntry.CcCopyIndex, "ccCopyIndex")
    ccCopyEntry.EntityData.AbsolutePath = "CISCO-CONFIG-COPY-MIB:CISCO-CONFIG-COPY-MIB/ccCopyTable/" + ccCopyEntry.EntityData.SegmentPath
    ccCopyEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccCopyEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccCopyEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccCopyEntry.EntityData.Children = types.NewOrderedMap()
    ccCopyEntry.EntityData.Leafs = types.NewOrderedMap()
    ccCopyEntry.EntityData.Leafs.Append("ccCopyIndex", types.YLeaf{"CcCopyIndex", ccCopyEntry.CcCopyIndex})
    ccCopyEntry.EntityData.Leafs.Append("ccCopyProtocol", types.YLeaf{"CcCopyProtocol", ccCopyEntry.CcCopyProtocol})
    ccCopyEntry.EntityData.Leafs.Append("ccCopySourceFileType", types.YLeaf{"CcCopySourceFileType", ccCopyEntry.CcCopySourceFileType})
    ccCopyEntry.EntityData.Leafs.Append("ccCopyDestFileType", types.YLeaf{"CcCopyDestFileType", ccCopyEntry.CcCopyDestFileType})
    ccCopyEntry.EntityData.Leafs.Append("ccCopyServerAddress", types.YLeaf{"CcCopyServerAddress", ccCopyEntry.CcCopyServerAddress})
    ccCopyEntry.EntityData.Leafs.Append("ccCopyFileName", types.YLeaf{"CcCopyFileName", ccCopyEntry.CcCopyFileName})
    ccCopyEntry.EntityData.Leafs.Append("ccCopyUserName", types.YLeaf{"CcCopyUserName", ccCopyEntry.CcCopyUserName})
    ccCopyEntry.EntityData.Leafs.Append("ccCopyUserPassword", types.YLeaf{"CcCopyUserPassword", ccCopyEntry.CcCopyUserPassword})
    ccCopyEntry.EntityData.Leafs.Append("ccCopyNotificationOnCompletion", types.YLeaf{"CcCopyNotificationOnCompletion", ccCopyEntry.CcCopyNotificationOnCompletion})
    ccCopyEntry.EntityData.Leafs.Append("ccCopyState", types.YLeaf{"CcCopyState", ccCopyEntry.CcCopyState})
    ccCopyEntry.EntityData.Leafs.Append("ccCopyTimeStarted", types.YLeaf{"CcCopyTimeStarted", ccCopyEntry.CcCopyTimeStarted})
    ccCopyEntry.EntityData.Leafs.Append("ccCopyTimeCompleted", types.YLeaf{"CcCopyTimeCompleted", ccCopyEntry.CcCopyTimeCompleted})
    ccCopyEntry.EntityData.Leafs.Append("ccCopyFailCause", types.YLeaf{"CcCopyFailCause", ccCopyEntry.CcCopyFailCause})
    ccCopyEntry.EntityData.Leafs.Append("ccCopyEntryRowStatus", types.YLeaf{"CcCopyEntryRowStatus", ccCopyEntry.CcCopyEntryRowStatus})
    ccCopyEntry.EntityData.Leafs.Append("ccCopyServerAddressType", types.YLeaf{"CcCopyServerAddressType", ccCopyEntry.CcCopyServerAddressType})
    ccCopyEntry.EntityData.Leafs.Append("ccCopyServerAddressRev1", types.YLeaf{"CcCopyServerAddressRev1", ccCopyEntry.CcCopyServerAddressRev1})

    ccCopyEntry.EntityData.YListKeys = []string {"CcCopyIndex"}

    return &(ccCopyEntry.EntityData)
}

// CISCOCONFIGCOPYMIB_CcCopyErrorTable
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
type CISCOCONFIGCOPYMIB_CcCopyErrorTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry containing information about the outcome at one destination of a
    // failed config copy operation. The type is slice of
    // CISCOCONFIGCOPYMIB_CcCopyErrorTable_CcCopyErrorEntry.
    CcCopyErrorEntry []*CISCOCONFIGCOPYMIB_CcCopyErrorTable_CcCopyErrorEntry
}

func (ccCopyErrorTable *CISCOCONFIGCOPYMIB_CcCopyErrorTable) GetEntityData() *types.CommonEntityData {
    ccCopyErrorTable.EntityData.YFilter = ccCopyErrorTable.YFilter
    ccCopyErrorTable.EntityData.YangName = "ccCopyErrorTable"
    ccCopyErrorTable.EntityData.BundleName = "cisco_ios_xe"
    ccCopyErrorTable.EntityData.ParentYangName = "CISCO-CONFIG-COPY-MIB"
    ccCopyErrorTable.EntityData.SegmentPath = "ccCopyErrorTable"
    ccCopyErrorTable.EntityData.AbsolutePath = "CISCO-CONFIG-COPY-MIB:CISCO-CONFIG-COPY-MIB/" + ccCopyErrorTable.EntityData.SegmentPath
    ccCopyErrorTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccCopyErrorTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccCopyErrorTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccCopyErrorTable.EntityData.Children = types.NewOrderedMap()
    ccCopyErrorTable.EntityData.Children.Append("ccCopyErrorEntry", types.YChild{"CcCopyErrorEntry", nil})
    for i := range ccCopyErrorTable.CcCopyErrorEntry {
        ccCopyErrorTable.EntityData.Children.Append(types.GetSegmentPath(ccCopyErrorTable.CcCopyErrorEntry[i]), types.YChild{"CcCopyErrorEntry", ccCopyErrorTable.CcCopyErrorEntry[i]})
    }
    ccCopyErrorTable.EntityData.Leafs = types.NewOrderedMap()

    ccCopyErrorTable.EntityData.YListKeys = []string {}

    return &(ccCopyErrorTable.EntityData)
}

// CISCOCONFIGCOPYMIB_CcCopyErrorTable_CcCopyErrorEntry
// An entry containing information about the
// outcome at one destination of a failed config
// copy operation.
type CISCOCONFIGCOPYMIB_CcCopyErrorTable_CcCopyErrorEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // cisco_config_copy_mib.CISCOCONFIGCOPYMIB_CcCopyTable_CcCopyEntry_CcCopyIndex
    CcCopyIndex interface{}

    // This attribute is a key. A monotonically increasing integer for the sole
    // purpose of indexing entries in this table. When a config copy operation has
    // multiple  destinations, then this index value is used to  distinguish
    // between those multiple destinations. The type is interface{} with range:
    // 0..4294967295.
    CcCopyErrorIndex interface{}

    // The type of Internet address for this destination device on which config
    // copy operation is performed. The type is InetAddressType.
    CcCopyErrorDeviceIpAddressType interface{}

    // The IP address of this destination device on which config copy operation is
    // performed. The object value has to be consistent with the type specified in
    // ccCopyErrorDeviceIpAddressType. The type is string with length: 0..255.
    CcCopyErrorDeviceIpAddress interface{}

    // The World Wide Name (WWN) of this destination device on which config copy
    // operation is performed. The value of this object is zero-length string if 
    // WWN is unassigned or unknown. For example, devices  which do not support
    // fibre channel would not have WWN. The type is string with length: 0..0 |
    // 8..8 | 16..16.
    CcCopyErrorDeviceWWN interface{}

    // The error description for the error happened for this destination of this
    // config copy  operation. The type is string.
    CcCopyErrorDescription interface{}
}

func (ccCopyErrorEntry *CISCOCONFIGCOPYMIB_CcCopyErrorTable_CcCopyErrorEntry) GetEntityData() *types.CommonEntityData {
    ccCopyErrorEntry.EntityData.YFilter = ccCopyErrorEntry.YFilter
    ccCopyErrorEntry.EntityData.YangName = "ccCopyErrorEntry"
    ccCopyErrorEntry.EntityData.BundleName = "cisco_ios_xe"
    ccCopyErrorEntry.EntityData.ParentYangName = "ccCopyErrorTable"
    ccCopyErrorEntry.EntityData.SegmentPath = "ccCopyErrorEntry" + types.AddKeyToken(ccCopyErrorEntry.CcCopyIndex, "ccCopyIndex") + types.AddKeyToken(ccCopyErrorEntry.CcCopyErrorIndex, "ccCopyErrorIndex")
    ccCopyErrorEntry.EntityData.AbsolutePath = "CISCO-CONFIG-COPY-MIB:CISCO-CONFIG-COPY-MIB/ccCopyErrorTable/" + ccCopyErrorEntry.EntityData.SegmentPath
    ccCopyErrorEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccCopyErrorEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccCopyErrorEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccCopyErrorEntry.EntityData.Children = types.NewOrderedMap()
    ccCopyErrorEntry.EntityData.Leafs = types.NewOrderedMap()
    ccCopyErrorEntry.EntityData.Leafs.Append("ccCopyIndex", types.YLeaf{"CcCopyIndex", ccCopyErrorEntry.CcCopyIndex})
    ccCopyErrorEntry.EntityData.Leafs.Append("ccCopyErrorIndex", types.YLeaf{"CcCopyErrorIndex", ccCopyErrorEntry.CcCopyErrorIndex})
    ccCopyErrorEntry.EntityData.Leafs.Append("ccCopyErrorDeviceIpAddressType", types.YLeaf{"CcCopyErrorDeviceIpAddressType", ccCopyErrorEntry.CcCopyErrorDeviceIpAddressType})
    ccCopyErrorEntry.EntityData.Leafs.Append("ccCopyErrorDeviceIpAddress", types.YLeaf{"CcCopyErrorDeviceIpAddress", ccCopyErrorEntry.CcCopyErrorDeviceIpAddress})
    ccCopyErrorEntry.EntityData.Leafs.Append("ccCopyErrorDeviceWWN", types.YLeaf{"CcCopyErrorDeviceWWN", ccCopyErrorEntry.CcCopyErrorDeviceWWN})
    ccCopyErrorEntry.EntityData.Leafs.Append("ccCopyErrorDescription", types.YLeaf{"CcCopyErrorDescription", ccCopyErrorEntry.CcCopyErrorDescription})

    ccCopyErrorEntry.EntityData.YListKeys = []string {"CcCopyIndex", "CcCopyErrorIndex"}

    return &(ccCopyErrorEntry.EntityData)
}

