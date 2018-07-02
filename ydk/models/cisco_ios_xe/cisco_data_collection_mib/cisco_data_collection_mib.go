// The MIB module allows a management application to
// select a set of MIB object instances whose values need 
// to be collected on a periodic basis. The term 'data' in 
// the context of this MIB is used to generically refer to 
// the values of the selected set of object instances. 
// 
// Once the required setup is done, the MIB implementation 
// carries out the following periodic tasks:
//   - collects the required object values into local
//     file-like entities called VFiles (virtual files).
//   - transfers the VFiles to specified locations.
//   - carries out VFile management operations.
// 
// Some of the key features of this MIB are:
//   a) Allows grouping of MIB objects into groups called 
//      data groups. The constraint is that the MIB objects 
//      grouped into a data group, need to have the same 
//      semantic MIB index. So it is possible to group MIB 
//      objects belonging to different MIB tables into a 
//      single data group as long as the above constraint is 
//      met.
//      For e.g. it is possible to group ifInOctets from 
//      ifTable, ifHCInOctets from ifXTable, 
//      dot3StatsExcessiveCollisions from dot3StatsTable 
//      into a single data group.
// 
//   b) Allows the application to specify a set of instances 
//      (of the MIB objects in a data group) whose values 
//      need to be collected. 
// 
//   c) The required data can be collected for each such 
//      data group on a periodic basis into a virtual file
//      (VFile). A VFile is an abstraction of a file. 
// 
//   d) The format of the contents of the VFile, can be 
//      specified by the application. 
// 
//   e) An application can also specify a collection period. 
//      A collection period is an interval of time during 
//      which data is collected into a VFile. After the 
//      collection period ends, the VFile is frozen, and a 
//      new VFile is created for storing data. The frozen 
//      VFile is then transferred to a specified destination. 
//      An application can choose to retain such frozen 
//      VFiles on the device for a certain period of time, 
//      called the retention period. 
// 
//          Data Collection MIB vs Bulkfile MIB
//          ***********************************
//    The data collection MIB can pretty much do what the
//    CISCO-BULK-FILE-MIB (Bulkfile MIB) can do. The 'manual' 
//    mode of the Data collection MIB is similar to the way 
//    in which the Bulkfile MIB operates.
// 
//    However the data collection MIB is mainly targetted 
//    for medium to high-end platforms which have sufficient
//    local storage (volatile or permanent) to store VFiles.
//    Locally storing VFiles, helps minimize loss of data 
//    during temporary network outages. If the local store 
//    is permament, then the collected data is also available 
//    across agent restarts.  
// 
//    The data collection MIB has more powerful data 
//    selection features than the Bulkfile MIB. It allows 
//    grouping of MIB objects from different tables into 
//    data groups. It also incorporates a more flexible 
//    instance selection mechanism, where the application 
//    is not restricted to fetching an entire MIB table. 
// 
//                 Definitions:
//                 ************
//     Base objects: 
//     *************
//     MIB objects whose values an application needs to 
//     collect.
// 
//     Data group:
//     ***********
//     A group of base objects. Can be of 2 types: 'object' 
//     and 'table'. An 'object' type data group can consist 
//     of only one fully instantiated base object. A 'table'
//     type data group can consist of more than one base
//     objects, each a columnar object in a conceptual
//     table. In addition a 'table' type data group can
//     specify the instances of the base objects whose 
//     values need to be collected. In the context of this 
//     MIB, collecting data for a data group means fetching 
//     the values of the associated base object instances 
//     and storing them into VFiles.
// 
//     Virtual File (VFile):
//     *********************
//     A VFile is a file like entity used to collect data. 
//     An agent might choose to implement a VFile as a 
//     simple in-memory buffer, or it might choose to
//     use a file in it's filesystem. An application does
//     not really need to know the location of a VFile 
//     - the MIB provides mechanisms to transfer the 
//     VFile to application specified locations. However 
//     if the implementation supports it, the application 
//     can specify the location of the VFiles.
// 
//     Current VFile:
//     **************
//     The VFile into which data is currently being 
//      collected.
// 
//     Frozen VFile:
//     *************
//     A VFile which is no longer used for collecting 
//     data. Only frozen VFiles can be transferred to 
//     specified destinations.
// 
//     Collection interval:
//     ********************
//     A collection interval is associated with a VFile.
//     It is the interval of time over which a VFile 
//     is used to collect data. 
//     This interval of time can be specified by the 
//     application. However there are conditions under 
//     which a collection interval can be shorter than 
//     the specified time. For e.g. a collection 
//     interval is prematurely terminated when the 
//     maximum size of a VFile is exceeded, or when 
//     there is an error condition.
// 
//     Polling period:
//     ***************
//     A polling period is associated with a data 
//     group. It determines the frequency at which 
//     the base objects of a data group should
//     be fetched and stored into a VFile.
// 
//     Data collection operations:
//     **************************
//     A generic term used to refer to operations 
//     that are carried out while collecting data. 
//     These include:
//        - Periodically creating new VFiles for 
//          collecting data.
//        - Transferring frozen VFiles either 
//          automatically or on demand.
//        - Fetching base object values and storing 
//          them into current VFiles, either periodically 
//          or on demand.
//        - Deleting frozen VFiles, either periodically 
//          or on demand.
package cisco_data_collection_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_data_collection_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-DATA-COLLECTION-MIB CISCO-DATA-COLLECTION-MIB}", reflect.TypeOf(CISCODATACOLLECTIONMIB{}))
    ydk.RegisterEntity("CISCO-DATA-COLLECTION-MIB:CISCO-DATA-COLLECTION-MIB", reflect.TypeOf(CISCODATACOLLECTIONMIB{}))
}

// CdcFileXferStatus represents                        of FTP(File Transfer Protocol).
type CdcFileXferStatus string

const (
    CdcFileXferStatus_notStarted CdcFileXferStatus = "notStarted"

    CdcFileXferStatus_success CdcFileXferStatus = "success"

    CdcFileXferStatus_aborted CdcFileXferStatus = "aborted"

    CdcFileXferStatus_fileOpenFailRemote CdcFileXferStatus = "fileOpenFailRemote"

    CdcFileXferStatus_badDomainName CdcFileXferStatus = "badDomainName"

    CdcFileXferStatus_unreachableIpAddress CdcFileXferStatus = "unreachableIpAddress"

    CdcFileXferStatus_networkFailed CdcFileXferStatus = "networkFailed"

    CdcFileXferStatus_fileWriteFailed CdcFileXferStatus = "fileWriteFailed"

    CdcFileXferStatus_authFailed CdcFileXferStatus = "authFailed"
)

// CdcFileFormat represents  immediately precedes the next tag or the end of file.
type CdcFileFormat string

const (
    CdcFileFormat_cdcBulkASCII CdcFileFormat = "cdcBulkASCII"

    CdcFileFormat_cdcBulkBinary CdcFileFormat = "cdcBulkBinary"

    CdcFileFormat_cdcSchemaASCII CdcFileFormat = "cdcSchemaASCII"
)

// CISCODATACOLLECTIONMIB
type CISCODATACOLLECTIONMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    CdcVFile CISCODATACOLLECTIONMIB_CdcVFile

    // A table for setting up VFiles for collecting data.
    CdcVFileTable CISCODATACOLLECTIONMIB_CdcVFileTable

    // A table to manage frozen VFiles.
    CdcVFileMgmtTable CISCODATACOLLECTIONMIB_CdcVFileMgmtTable

    // A table for specifying data groups.
    CdcDGTable CISCODATACOLLECTIONMIB_CdcDGTable

    // A table specifying the base objects of a 'table' type data group.
    CdcDGBaseObjectTable CISCODATACOLLECTIONMIB_CdcDGBaseObjectTable

    // Identifies the instances of the base objects that need to be fetched for a
    // 'table' type data group.  The agent is not responsible for verifying that
    // the instances specified for a data group do not overlap.
    CdcDGInstanceTable CISCODATACOLLECTIONMIB_CdcDGInstanceTable

    // A table for configuring file transfer operations.
    CdcFileXferConfTable CISCODATACOLLECTIONMIB_CdcFileXferConfTable
}

func (cISCODATACOLLECTIONMIB *CISCODATACOLLECTIONMIB) GetEntityData() *types.CommonEntityData {
    cISCODATACOLLECTIONMIB.EntityData.YFilter = cISCODATACOLLECTIONMIB.YFilter
    cISCODATACOLLECTIONMIB.EntityData.YangName = "CISCO-DATA-COLLECTION-MIB"
    cISCODATACOLLECTIONMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCODATACOLLECTIONMIB.EntityData.ParentYangName = "CISCO-DATA-COLLECTION-MIB"
    cISCODATACOLLECTIONMIB.EntityData.SegmentPath = "CISCO-DATA-COLLECTION-MIB:CISCO-DATA-COLLECTION-MIB"
    cISCODATACOLLECTIONMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCODATACOLLECTIONMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCODATACOLLECTIONMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCODATACOLLECTIONMIB.EntityData.Children = types.NewOrderedMap()
    cISCODATACOLLECTIONMIB.EntityData.Children.Append("cdcVFile", types.YChild{"CdcVFile", &cISCODATACOLLECTIONMIB.CdcVFile})
    cISCODATACOLLECTIONMIB.EntityData.Children.Append("cdcVFileTable", types.YChild{"CdcVFileTable", &cISCODATACOLLECTIONMIB.CdcVFileTable})
    cISCODATACOLLECTIONMIB.EntityData.Children.Append("cdcVFileMgmtTable", types.YChild{"CdcVFileMgmtTable", &cISCODATACOLLECTIONMIB.CdcVFileMgmtTable})
    cISCODATACOLLECTIONMIB.EntityData.Children.Append("cdcDGTable", types.YChild{"CdcDGTable", &cISCODATACOLLECTIONMIB.CdcDGTable})
    cISCODATACOLLECTIONMIB.EntityData.Children.Append("cdcDGBaseObjectTable", types.YChild{"CdcDGBaseObjectTable", &cISCODATACOLLECTIONMIB.CdcDGBaseObjectTable})
    cISCODATACOLLECTIONMIB.EntityData.Children.Append("cdcDGInstanceTable", types.YChild{"CdcDGInstanceTable", &cISCODATACOLLECTIONMIB.CdcDGInstanceTable})
    cISCODATACOLLECTIONMIB.EntityData.Children.Append("cdcFileXferConfTable", types.YChild{"CdcFileXferConfTable", &cISCODATACOLLECTIONMIB.CdcFileXferConfTable})
    cISCODATACOLLECTIONMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCODATACOLLECTIONMIB.EntityData.YListKeys = []string {}

    return &(cISCODATACOLLECTIONMIB.EntityData)
}

// CISCODATACOLLECTIONMIB_CdcVFile
type CISCODATACOLLECTIONMIB_CdcVFile struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object's value reads 'true', if the agent implementation of this MIB
    // supports placement of VFiles in application specified persistent storage
    // locations. Otherwise  the value is 'false'. The type is bool.
    CdcVFilePersistentStorage interface{}

    // A global limit for the number of consecutive times the maximum VFile size
    // (cdcVFileMaxSize) is exceeded for a given VFile. When this limit is
    // exceeded the offending cdcVFileEntry is moved to the error state (see
    // cdcVFileOperStatus). The type is interface{} with range: 1..4294967295.
    CdcVFileMaxSizeHitsLimit interface{}
}

func (cdcVFile *CISCODATACOLLECTIONMIB_CdcVFile) GetEntityData() *types.CommonEntityData {
    cdcVFile.EntityData.YFilter = cdcVFile.YFilter
    cdcVFile.EntityData.YangName = "cdcVFile"
    cdcVFile.EntityData.BundleName = "cisco_ios_xe"
    cdcVFile.EntityData.ParentYangName = "CISCO-DATA-COLLECTION-MIB"
    cdcVFile.EntityData.SegmentPath = "cdcVFile"
    cdcVFile.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdcVFile.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdcVFile.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdcVFile.EntityData.Children = types.NewOrderedMap()
    cdcVFile.EntityData.Leafs = types.NewOrderedMap()
    cdcVFile.EntityData.Leafs.Append("cdcVFilePersistentStorage", types.YLeaf{"CdcVFilePersistentStorage", cdcVFile.CdcVFilePersistentStorage})
    cdcVFile.EntityData.Leafs.Append("cdcVFileMaxSizeHitsLimit", types.YLeaf{"CdcVFileMaxSizeHitsLimit", cdcVFile.CdcVFileMaxSizeHitsLimit})

    cdcVFile.EntityData.YListKeys = []string {}

    return &(cdcVFile.EntityData)
}

// CISCODATACOLLECTIONMIB_CdcVFileTable
// A table for setting up VFiles for collecting data.
type CISCODATACOLLECTIONMIB_CdcVFileTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the cdcVFileTable. Each entry contains application specified
    // configuration that is used to create  virtual files (VFile) and start data
    // collection operations.   A VFile is used to store data (values of base
    // object instances) as selected by entities called data groups. A data group
    // is defined in cdcDGTable.   An entry in this table is said to be
    // 'activated' when the corresponding instances of cdcVFileRowStatus is
    // 'active' AND cdcVFileOperStatus is 'enabled'. The value of sysUpTime.0 when
    // the condition evaluates to 'true' is called the activation time of the
    // entry. The activation time for each entry is maintained internally by the
    // agent. The type is slice of
    // CISCODATACOLLECTIONMIB_CdcVFileTable_CdcVFileEntry.
    CdcVFileEntry []*CISCODATACOLLECTIONMIB_CdcVFileTable_CdcVFileEntry
}

func (cdcVFileTable *CISCODATACOLLECTIONMIB_CdcVFileTable) GetEntityData() *types.CommonEntityData {
    cdcVFileTable.EntityData.YFilter = cdcVFileTable.YFilter
    cdcVFileTable.EntityData.YangName = "cdcVFileTable"
    cdcVFileTable.EntityData.BundleName = "cisco_ios_xe"
    cdcVFileTable.EntityData.ParentYangName = "CISCO-DATA-COLLECTION-MIB"
    cdcVFileTable.EntityData.SegmentPath = "cdcVFileTable"
    cdcVFileTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdcVFileTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdcVFileTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdcVFileTable.EntityData.Children = types.NewOrderedMap()
    cdcVFileTable.EntityData.Children.Append("cdcVFileEntry", types.YChild{"CdcVFileEntry", nil})
    for i := range cdcVFileTable.CdcVFileEntry {
        cdcVFileTable.EntityData.Children.Append(types.GetSegmentPath(cdcVFileTable.CdcVFileEntry[i]), types.YChild{"CdcVFileEntry", cdcVFileTable.CdcVFileEntry[i]})
    }
    cdcVFileTable.EntityData.Leafs = types.NewOrderedMap()

    cdcVFileTable.EntityData.YListKeys = []string {}

    return &(cdcVFileTable.EntityData)
}

// CISCODATACOLLECTIONMIB_CdcVFileTable_CdcVFileEntry
// An entry in the cdcVFileTable. Each entry contains
// application specified configuration that is used to create
//  virtual files (VFile) and start data collection operations.
// 
//  A VFile is used to store data (values of base object
// instances) as selected by entities called data groups.
// A data group is defined in cdcDGTable. 
// 
// An entry in this table is said to be 'activated' when the
// corresponding instances of cdcVFileRowStatus is 'active' AND
// cdcVFileOperStatus is 'enabled'. The value of sysUpTime.0 when
// the condition evaluates to 'true' is called the activation
// time of the entry. The activation time for each entry is
// maintained internally by the agent.
type CISCODATACOLLECTIONMIB_CdcVFileTable_CdcVFileEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An arbitrary integer for uniquely identifying this
    // entry. When creating a row, the application should pick a random number.  
    // If the configuration in this entry is persisted across system/agent
    // restarts then the same value of cdcVFileIndex must be assigned to this
    // entry after the restart. The type is interface{} with range: 1..4294967295.
    CdcVFileIndex interface{}

    // The base-name of the VFiles (created by data collection operations
    // corresponding to this entry) into which data is to be collected.   When a
    // VFile is created, it's full name is obtained by the concatentation of a
    // suffix to this value. The suffix will be chosen by the agent such that the
    // VFiles created for this entry have unique names. For e.g. the suffix could
    // be a string representation of the date and time when the VFile was created.
    // If VFiles are to be placed in the agent's local filesystem (provided the
    // agent supports it) then this value should also contain the absolute path of
    // the location as a prefix to the base name.  An agent will respond with
    // inconsistentValue to a management set operation which attempts to modify
    // the value of this object to the same value as already held by another
    // instance of cdcVFileName, or wrongValue if the new value  is invalid for
    // use as a file name on the local file  system (e.g., many file systems do
    // not support white  space embedded in file names).  This object's value may
    // be modified at any time. However the new name will be used only when the
    // next VFile is created for this entry. The type is string with length:
    // 0..255.
    CdcVFileName interface{}

    // A string that can be used for administrative purposes.  This object's value
    // may be modified at any time. The type is string.
    CdcVFileDescription interface{}

    // An object for controlling collection of data.  'idle'            Indicates
    // that no command is in progress.  'swapToNewFile'   When written, the
    // current VFile is frozen,                   and a new VFile is created for
    // collecting                   data. 		   If the data collection mode is
    // automatic                   (see cdcVFileCollectMode), then the current    
    // collection interval is stopped and a new                   collection
    // interval is started  		   (see cdcVFileCollectPeriod).                     
    // 'collectNow'      When written, base object values for                  
    // all associated data groups are fetched                    and stored into
    // the current VFile. This                   value can only be written when
    // the                   collection mode is 'manual' (see                   
    // cdcVFileCollectMode). The type is CdcVFileCommand.
    CdcVFileCommand interface{}

    // The maximum size of a VFile.   The agent maintains an internal counter for
    // each cdcVFileEntry. This counter counts the number of consecutive times the
    // size of a VFile has exceeded the value of this object. When the value of
    // this counter exceeds the value of cdcVFileMaxSizeHitsLimit, this entry is
    // moved to the 'error' state (see cdcVFileOperStatus). However if the value
    // of cdcVFileMaxSizeHitsLimit is not exceeded, then the current VFile is
    // frozen, and a new VFile is created for data collection.  If the data
    // collection mode is automatic (see cdcVFileCollectMode), then the current
    // collection interval is stopped and a new collection interval is started. 
    // This object's value may be modified at any time. The new size limit MUST be
    // checked against the size of the current VFile at the time of modification,
    // and appropriate action taken. The type is interface{} with range:
    // 512..4294967295. Units are bytes.
    CdcVFileMaxSize interface{}

    // The size of the current VFile. The type is interface{} with range:
    // 0..4294967295. Units are bytes.
    CdcVFileCurrentSize interface{}

    // The format in which data is stored into the current VFile.  This object's
    // value cannot be modified while the entry is in the 'activated' state. The
    // type is CdcFileFormat.
    CdcVFileFormat interface{}

    // Determines the mode of data collection.  'auto'         Data is
    // periodically fetched for all data                groups associated with
    // this entry. This is                done at data group specific periodic
    // intervals                (cdcDGPollPeriod).                The data thus
    // collected, is formatted and                stored into the current VFile.  
    // In addition at regular intervals (see                cdcVFileCollectPeriod)
    // a new VFile                 is created to store data, and the current      
    // VFile is frozen and transferred.  'manual'       Data for all data groups
    // is fetched and                collected into the current VFile only when   
    // cdcVFileCommand is set to 'collectNow'.   This object's value cannot be
    // modified while the entry is in the 'activated' state. The type is
    // CdcVFileCollectMode.
    CdcVFileCollectMode interface{}

    // Specifies the period of a collection interval. The value of this object is
    // used only when the data collection mode is  automatic (see
    // cdcVFileCollectMode).  A periodic timer (one per entry) is started when
    // this entry is 'activated'. The time at which this entry is 'activated' is
    // called the 'activation time' for this entry, and is internally maintained
    // by the agent.  When this periodic timer expires, the current VFile is
    // frozen and a new VFile is created for data collection.  Transfer is then
    // initiated for the frozen VFile.   In addition, the internally maintained
    // counter for counting the number of consecutive times the size of a VFile
    // has exceeded the maximum limit, is reset to zero. (See cdcVFileMaxSize)  
    // This object's value may be modified at any time, and the  new value takes
    // effect immediately. i.e setting a new  value can cause the current
    // collection interval to terminate  and a new collection interval to start.
    // The type is interface{} with range: 60..604800. Units are seconds.
    CdcVFileCollectionPeriod interface{}

    // The time for which a frozen VFile is retained by the agent. When a VFile is
    // frozen, a timer (one per frozen VFile) is started to keep track of the
    // retention period for the  VFile. Once this timer expires, the VFile is
    // deleted. Till the expiry of the retention period, information  about frozen
    // VFiles is maintained in  cdcVFileMgmtTable.  This object's value may be
    // modified at any time, however the new value will take effect only for new
    // frozen VFiles. The type is interface{} with range: 60..86400. Units are
    // seconds.
    CdcVFileRetentionPeriod interface{}

    // A control object to indicate the administratively desired state of data
    // collection for this entry. On setting the value to 'disabled' data
    // collection operations for this  entry are stopped, the current VFile is
    // frozen and it's  transfer is initiated.   Modifying the value of
    // cdcVFileAdminStatus to 'disabled' does not remove or change the current
    // configuration as represented by the active rows in this table. The type is
    // CdcVFileAdminStatus.
    CdcVFileAdminStatus interface{}

    // A status object to indicate the operational state of collection for this
    // entry.  When the value of cdcVFileAdminStatus is modified to be 'enabled',
    // the value of this object will change to 'enabled' providing it is possible
    // to begin collecting data. If at any point of time data cannot be collected
    // because of some error, then the value of this object is changed to 'error'
    // and all collection operations stop, as if cdcVFileAdminStatus has been set
    // to 'disabled'. More information about the nature of the error can be
    // obtained by retrieving the value of cdcVFileErrorCode.   When the value of
    // cdcVFileAdminStatus is modified to be 'disabled', the value of this object
    // will change to 'disabled' and data collection operations are stopped for
    // this entry. The type is CdcVFileOperStatus.
    CdcVFileOperStatus interface{}

    // A value indicating the type of error that has occurred during data
    // collection operations for this entry.  noError                The value is
    // 'noError' when                        the corresponding value of           
    // cdcVFileOperStatus is not 'error'.  otherError             Any error other
    // than one of the                         following listed errors.  noSpace  
    // There is no space left to write into                        the current
    // VFile.   openError              Could not open VFile for writing. One      
    // possible reason could be the existence                        of another
    // file by the same name in                        the agent's filesystem.  
    // tooSmallMaxSize        Indicates that cdcVFileMaxSize is                   
    // too small for data collection. The                         cdcVFileMaxSize
    // configured for this                        VFile is not sufficient even to
    // hold                         the data collected in one poll. 
    // tooManyMaxSizeHits     Indicates that data collection                      
    // operations are stopped because                        the value of
    // cdcVFileMaxSizeHitsLimit                        has been exceeded.  
    // noResource             Some kind of resource was unavailable               
    // while collecting data. For                        e.g. unavailability of
    // dynamic memory. The type is CdcVFileErrorCode.
    CdcVFileErrorCode interface{}

    // When set to 'true', cdcVFileCollectionError notification will be sent out
    // in the event of a data collection error. The type is bool.
    CdcVFileCollectionErrorEnable interface{}

    // The status of this conceptual row.  A valid cdcVFileName is only mandatory
    // object for setting this object to 'active'. But collection of data in to
    // active vfile starts only on setting cdcVFileAdminStatus  to 'enabled'.
    // Setting this object to 'destroy' stops all data collection operations for
    // this entry, deletes all VFiles and removes this entry from cdcVFileTable.
    // The type is RowStatus.
    CdcVFileRowStatus interface{}
}

func (cdcVFileEntry *CISCODATACOLLECTIONMIB_CdcVFileTable_CdcVFileEntry) GetEntityData() *types.CommonEntityData {
    cdcVFileEntry.EntityData.YFilter = cdcVFileEntry.YFilter
    cdcVFileEntry.EntityData.YangName = "cdcVFileEntry"
    cdcVFileEntry.EntityData.BundleName = "cisco_ios_xe"
    cdcVFileEntry.EntityData.ParentYangName = "cdcVFileTable"
    cdcVFileEntry.EntityData.SegmentPath = "cdcVFileEntry" + types.AddKeyToken(cdcVFileEntry.CdcVFileIndex, "cdcVFileIndex")
    cdcVFileEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdcVFileEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdcVFileEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdcVFileEntry.EntityData.Children = types.NewOrderedMap()
    cdcVFileEntry.EntityData.Leafs = types.NewOrderedMap()
    cdcVFileEntry.EntityData.Leafs.Append("cdcVFileIndex", types.YLeaf{"CdcVFileIndex", cdcVFileEntry.CdcVFileIndex})
    cdcVFileEntry.EntityData.Leafs.Append("cdcVFileName", types.YLeaf{"CdcVFileName", cdcVFileEntry.CdcVFileName})
    cdcVFileEntry.EntityData.Leafs.Append("cdcVFileDescription", types.YLeaf{"CdcVFileDescription", cdcVFileEntry.CdcVFileDescription})
    cdcVFileEntry.EntityData.Leafs.Append("cdcVFileCommand", types.YLeaf{"CdcVFileCommand", cdcVFileEntry.CdcVFileCommand})
    cdcVFileEntry.EntityData.Leafs.Append("cdcVFileMaxSize", types.YLeaf{"CdcVFileMaxSize", cdcVFileEntry.CdcVFileMaxSize})
    cdcVFileEntry.EntityData.Leafs.Append("cdcVFileCurrentSize", types.YLeaf{"CdcVFileCurrentSize", cdcVFileEntry.CdcVFileCurrentSize})
    cdcVFileEntry.EntityData.Leafs.Append("cdcVFileFormat", types.YLeaf{"CdcVFileFormat", cdcVFileEntry.CdcVFileFormat})
    cdcVFileEntry.EntityData.Leafs.Append("cdcVFileCollectMode", types.YLeaf{"CdcVFileCollectMode", cdcVFileEntry.CdcVFileCollectMode})
    cdcVFileEntry.EntityData.Leafs.Append("cdcVFileCollectionPeriod", types.YLeaf{"CdcVFileCollectionPeriod", cdcVFileEntry.CdcVFileCollectionPeriod})
    cdcVFileEntry.EntityData.Leafs.Append("cdcVFileRetentionPeriod", types.YLeaf{"CdcVFileRetentionPeriod", cdcVFileEntry.CdcVFileRetentionPeriod})
    cdcVFileEntry.EntityData.Leafs.Append("cdcVFileAdminStatus", types.YLeaf{"CdcVFileAdminStatus", cdcVFileEntry.CdcVFileAdminStatus})
    cdcVFileEntry.EntityData.Leafs.Append("cdcVFileOperStatus", types.YLeaf{"CdcVFileOperStatus", cdcVFileEntry.CdcVFileOperStatus})
    cdcVFileEntry.EntityData.Leafs.Append("cdcVFileErrorCode", types.YLeaf{"CdcVFileErrorCode", cdcVFileEntry.CdcVFileErrorCode})
    cdcVFileEntry.EntityData.Leafs.Append("cdcVFileCollectionErrorEnable", types.YLeaf{"CdcVFileCollectionErrorEnable", cdcVFileEntry.CdcVFileCollectionErrorEnable})
    cdcVFileEntry.EntityData.Leafs.Append("cdcVFileRowStatus", types.YLeaf{"CdcVFileRowStatus", cdcVFileEntry.CdcVFileRowStatus})

    cdcVFileEntry.EntityData.YListKeys = []string {"CdcVFileIndex"}

    return &(cdcVFileEntry.EntityData)
}

// CISCODATACOLLECTIONMIB_CdcVFileTable_CdcVFileEntry_CdcVFileAdminStatus represents by the active rows in this table.
type CISCODATACOLLECTIONMIB_CdcVFileTable_CdcVFileEntry_CdcVFileAdminStatus string

const (
    CISCODATACOLLECTIONMIB_CdcVFileTable_CdcVFileEntry_CdcVFileAdminStatus_enabled CISCODATACOLLECTIONMIB_CdcVFileTable_CdcVFileEntry_CdcVFileAdminStatus = "enabled"

    CISCODATACOLLECTIONMIB_CdcVFileTable_CdcVFileEntry_CdcVFileAdminStatus_disabled CISCODATACOLLECTIONMIB_CdcVFileTable_CdcVFileEntry_CdcVFileAdminStatus = "disabled"
)

// CISCODATACOLLECTIONMIB_CdcVFileTable_CdcVFileEntry_CdcVFileCollectMode represents is in the 'activated' state.
type CISCODATACOLLECTIONMIB_CdcVFileTable_CdcVFileEntry_CdcVFileCollectMode string

const (
    CISCODATACOLLECTIONMIB_CdcVFileTable_CdcVFileEntry_CdcVFileCollectMode_auto CISCODATACOLLECTIONMIB_CdcVFileTable_CdcVFileEntry_CdcVFileCollectMode = "auto"

    CISCODATACOLLECTIONMIB_CdcVFileTable_CdcVFileEntry_CdcVFileCollectMode_manual CISCODATACOLLECTIONMIB_CdcVFileTable_CdcVFileEntry_CdcVFileCollectMode = "manual"
)

// CISCODATACOLLECTIONMIB_CdcVFileTable_CdcVFileEntry_CdcVFileCommand represents                   cdcVFileCollectMode).
type CISCODATACOLLECTIONMIB_CdcVFileTable_CdcVFileEntry_CdcVFileCommand string

const (
    CISCODATACOLLECTIONMIB_CdcVFileTable_CdcVFileEntry_CdcVFileCommand_idle CISCODATACOLLECTIONMIB_CdcVFileTable_CdcVFileEntry_CdcVFileCommand = "idle"

    CISCODATACOLLECTIONMIB_CdcVFileTable_CdcVFileEntry_CdcVFileCommand_swapToNewFile CISCODATACOLLECTIONMIB_CdcVFileTable_CdcVFileEntry_CdcVFileCommand = "swapToNewFile"

    CISCODATACOLLECTIONMIB_CdcVFileTable_CdcVFileEntry_CdcVFileCommand_collectNow CISCODATACOLLECTIONMIB_CdcVFileTable_CdcVFileEntry_CdcVFileCommand = "collectNow"
)

// CISCODATACOLLECTIONMIB_CdcVFileTable_CdcVFileEntry_CdcVFileErrorCode represents                        e.g. unavailability of dynamic memory.
type CISCODATACOLLECTIONMIB_CdcVFileTable_CdcVFileEntry_CdcVFileErrorCode string

const (
    CISCODATACOLLECTIONMIB_CdcVFileTable_CdcVFileEntry_CdcVFileErrorCode_noError CISCODATACOLLECTIONMIB_CdcVFileTable_CdcVFileEntry_CdcVFileErrorCode = "noError"

    CISCODATACOLLECTIONMIB_CdcVFileTable_CdcVFileEntry_CdcVFileErrorCode_otherError CISCODATACOLLECTIONMIB_CdcVFileTable_CdcVFileEntry_CdcVFileErrorCode = "otherError"

    CISCODATACOLLECTIONMIB_CdcVFileTable_CdcVFileEntry_CdcVFileErrorCode_noSpace CISCODATACOLLECTIONMIB_CdcVFileTable_CdcVFileEntry_CdcVFileErrorCode = "noSpace"

    CISCODATACOLLECTIONMIB_CdcVFileTable_CdcVFileEntry_CdcVFileErrorCode_openError CISCODATACOLLECTIONMIB_CdcVFileTable_CdcVFileEntry_CdcVFileErrorCode = "openError"

    CISCODATACOLLECTIONMIB_CdcVFileTable_CdcVFileEntry_CdcVFileErrorCode_tooSmallMaxSize CISCODATACOLLECTIONMIB_CdcVFileTable_CdcVFileEntry_CdcVFileErrorCode = "tooSmallMaxSize"

    CISCODATACOLLECTIONMIB_CdcVFileTable_CdcVFileEntry_CdcVFileErrorCode_tooManyMaxSizeHits CISCODATACOLLECTIONMIB_CdcVFileTable_CdcVFileEntry_CdcVFileErrorCode = "tooManyMaxSizeHits"

    CISCODATACOLLECTIONMIB_CdcVFileTable_CdcVFileEntry_CdcVFileErrorCode_noResource CISCODATACOLLECTIONMIB_CdcVFileTable_CdcVFileEntry_CdcVFileErrorCode = "noResource"
)

// CISCODATACOLLECTIONMIB_CdcVFileTable_CdcVFileEntry_CdcVFileOperStatus represents this entry.
type CISCODATACOLLECTIONMIB_CdcVFileTable_CdcVFileEntry_CdcVFileOperStatus string

const (
    CISCODATACOLLECTIONMIB_CdcVFileTable_CdcVFileEntry_CdcVFileOperStatus_enabled CISCODATACOLLECTIONMIB_CdcVFileTable_CdcVFileEntry_CdcVFileOperStatus = "enabled"

    CISCODATACOLLECTIONMIB_CdcVFileTable_CdcVFileEntry_CdcVFileOperStatus_disabled CISCODATACOLLECTIONMIB_CdcVFileTable_CdcVFileEntry_CdcVFileOperStatus = "disabled"

    CISCODATACOLLECTIONMIB_CdcVFileTable_CdcVFileEntry_CdcVFileOperStatus_error_ CISCODATACOLLECTIONMIB_CdcVFileTable_CdcVFileEntry_CdcVFileOperStatus = "error"
)

// CISCODATACOLLECTIONMIB_CdcVFileMgmtTable
// A table to manage frozen VFiles.
type CISCODATACOLLECTIONMIB_CdcVFileMgmtTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in cdcVFileMgmtTable. Each entry corresponds to a frozen VFile. An
    // entry is created in this table, whenever a VFile is frozen. An entry is
    // removed from this table whenever a frozen VFile is deleted either because
    // the retention period elapsed or because it was adminstratively deleted.  If
    // the configuration specified in cdcVFileEntry is persisted across
    // system/agent restarts AND the VFiles created as a result of that
    // configuration are persisted across restarts, then this table must be
    // populated with entries corresponding to those persisted VFiles. However any
    // state related to an entry, like time to live etc. need not be maintained
    // across restarts. The type is slice of
    // CISCODATACOLLECTIONMIB_CdcVFileMgmtTable_CdcVFileMgmtEntry.
    CdcVFileMgmtEntry []*CISCODATACOLLECTIONMIB_CdcVFileMgmtTable_CdcVFileMgmtEntry
}

func (cdcVFileMgmtTable *CISCODATACOLLECTIONMIB_CdcVFileMgmtTable) GetEntityData() *types.CommonEntityData {
    cdcVFileMgmtTable.EntityData.YFilter = cdcVFileMgmtTable.YFilter
    cdcVFileMgmtTable.EntityData.YangName = "cdcVFileMgmtTable"
    cdcVFileMgmtTable.EntityData.BundleName = "cisco_ios_xe"
    cdcVFileMgmtTable.EntityData.ParentYangName = "CISCO-DATA-COLLECTION-MIB"
    cdcVFileMgmtTable.EntityData.SegmentPath = "cdcVFileMgmtTable"
    cdcVFileMgmtTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdcVFileMgmtTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdcVFileMgmtTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdcVFileMgmtTable.EntityData.Children = types.NewOrderedMap()
    cdcVFileMgmtTable.EntityData.Children.Append("cdcVFileMgmtEntry", types.YChild{"CdcVFileMgmtEntry", nil})
    for i := range cdcVFileMgmtTable.CdcVFileMgmtEntry {
        cdcVFileMgmtTable.EntityData.Children.Append(types.GetSegmentPath(cdcVFileMgmtTable.CdcVFileMgmtEntry[i]), types.YChild{"CdcVFileMgmtEntry", cdcVFileMgmtTable.CdcVFileMgmtEntry[i]})
    }
    cdcVFileMgmtTable.EntityData.Leafs = types.NewOrderedMap()

    cdcVFileMgmtTable.EntityData.YListKeys = []string {}

    return &(cdcVFileMgmtTable.EntityData)
}

// CISCODATACOLLECTIONMIB_CdcVFileMgmtTable_CdcVFileMgmtEntry
// An entry in cdcVFileMgmtTable. Each entry corresponds to a
// frozen VFile. An entry is created in this table, whenever a
// VFile is frozen. An entry is removed from this table whenever
// a frozen VFile is deleted either because the retention period
// elapsed or because it was adminstratively deleted.
// 
// If the configuration specified in cdcVFileEntry is persisted
// across system/agent restarts AND the VFiles created as a
// result of that configuration are persisted across restarts,
// then this table must be populated with entries corresponding
// to those persisted VFiles. However any state related to an
// entry, like time to live etc. need not be maintained
// across restarts.
type CISCODATACOLLECTIONMIB_CdcVFileMgmtTable_CdcVFileMgmtEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // cisco_data_collection_mib.CISCODATACOLLECTIONMIB_CdcVFileTable_CdcVFileEntry_CdcVFileIndex
    CdcVFileIndex interface{}

    // This attribute is a key. This value is a running counter starting at 1,
    // generated by the agent so that the combination of  cdcVFileIndex and
    // cdcVFileMgmtIndex uniquely identifies a frozen VFile. The deleted file
    // indicies do not get reused.  This object's value needs to be unique only
    // across the set of frozen VFiles corresponding to a cdcVFileEntry
    // (identified by cdcVFileIndex). The type is interface{} with range:
    // 1..4294967295.
    CdcVFileMgmtIndex interface{}

    // The full name of the VFile. If the VFile is stored as a file in the agent's
    // filesystem, then this value also contains the absolute path of the file.
    // The type is string.
    CdcVFileMgmtName interface{}

    // The timestamp when this VFile was created, in the date-time format. The
    // type is string.
    CdcVFileMgmtTimestamp interface{}

    // The time left before this VFile is deleted (see cdcVFileRetentionPeriod).
    // The type is interface{} with range: 60..86400. Units are seconds.
    CdcVFileMgmtTimeToLive interface{}

    // A control to manage VFiles.   idle           This value can be only be
    // read. It indicates                that no management action is currently
    // being                performed on this VFile.  delete         This value is
    // only written, and is used to                delete the frozen VFile.
    // Writing this value                will cause this entry to be removed from
    // this                table.   transfer       This value can be both read and
    // written.                When read it means that the VFile is in the        
    // process of being transferred. When written, it                initiates a
    // transfer for the VFile.  abortTransfer  This value can only be written, and
    // is used                to abort an ongoing transfer. The type is
    // CdcVFileMgmtCommand.
    CdcVFileMgmtCommand interface{}

    // The complete URL of the destination to which this VFile will be transferred
    // in the next attempt. The URL also includes the complete filename of the
    // remote file that will be created. When the default value of this object is 
    // retained this VFile will be transferred to the URL   specified in
    // cdcFileXferConfPriUrl or cdcFileXferConfSecUrl, as the case may be.  
    // However an application can specify a different URL, in which case the VFile
    // will be transferred to this new URL the next time transfer is initiated.  
    // This object's value may be modified at any time. The type is string with
    // length: 0..255.
    CdcVFileMgmtXferURL interface{}

    // Indicates the status of the last completed transfer. The type is
    // CdcFileXferStatus.
    CdcVFileMgmtLastXferStatus interface{}

    // Indicates the URL of the destination to which the last (completed) transfer
    // was initiated. The type is string with length: 0..255.
    CdcVFileMgmtLastXferURL interface{}
}

func (cdcVFileMgmtEntry *CISCODATACOLLECTIONMIB_CdcVFileMgmtTable_CdcVFileMgmtEntry) GetEntityData() *types.CommonEntityData {
    cdcVFileMgmtEntry.EntityData.YFilter = cdcVFileMgmtEntry.YFilter
    cdcVFileMgmtEntry.EntityData.YangName = "cdcVFileMgmtEntry"
    cdcVFileMgmtEntry.EntityData.BundleName = "cisco_ios_xe"
    cdcVFileMgmtEntry.EntityData.ParentYangName = "cdcVFileMgmtTable"
    cdcVFileMgmtEntry.EntityData.SegmentPath = "cdcVFileMgmtEntry" + types.AddKeyToken(cdcVFileMgmtEntry.CdcVFileIndex, "cdcVFileIndex") + types.AddKeyToken(cdcVFileMgmtEntry.CdcVFileMgmtIndex, "cdcVFileMgmtIndex")
    cdcVFileMgmtEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdcVFileMgmtEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdcVFileMgmtEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdcVFileMgmtEntry.EntityData.Children = types.NewOrderedMap()
    cdcVFileMgmtEntry.EntityData.Leafs = types.NewOrderedMap()
    cdcVFileMgmtEntry.EntityData.Leafs.Append("cdcVFileIndex", types.YLeaf{"CdcVFileIndex", cdcVFileMgmtEntry.CdcVFileIndex})
    cdcVFileMgmtEntry.EntityData.Leafs.Append("cdcVFileMgmtIndex", types.YLeaf{"CdcVFileMgmtIndex", cdcVFileMgmtEntry.CdcVFileMgmtIndex})
    cdcVFileMgmtEntry.EntityData.Leafs.Append("cdcVFileMgmtName", types.YLeaf{"CdcVFileMgmtName", cdcVFileMgmtEntry.CdcVFileMgmtName})
    cdcVFileMgmtEntry.EntityData.Leafs.Append("cdcVFileMgmtTimestamp", types.YLeaf{"CdcVFileMgmtTimestamp", cdcVFileMgmtEntry.CdcVFileMgmtTimestamp})
    cdcVFileMgmtEntry.EntityData.Leafs.Append("cdcVFileMgmtTimeToLive", types.YLeaf{"CdcVFileMgmtTimeToLive", cdcVFileMgmtEntry.CdcVFileMgmtTimeToLive})
    cdcVFileMgmtEntry.EntityData.Leafs.Append("cdcVFileMgmtCommand", types.YLeaf{"CdcVFileMgmtCommand", cdcVFileMgmtEntry.CdcVFileMgmtCommand})
    cdcVFileMgmtEntry.EntityData.Leafs.Append("cdcVFileMgmtXferURL", types.YLeaf{"CdcVFileMgmtXferURL", cdcVFileMgmtEntry.CdcVFileMgmtXferURL})
    cdcVFileMgmtEntry.EntityData.Leafs.Append("cdcVFileMgmtLastXferStatus", types.YLeaf{"CdcVFileMgmtLastXferStatus", cdcVFileMgmtEntry.CdcVFileMgmtLastXferStatus})
    cdcVFileMgmtEntry.EntityData.Leafs.Append("cdcVFileMgmtLastXferURL", types.YLeaf{"CdcVFileMgmtLastXferURL", cdcVFileMgmtEntry.CdcVFileMgmtLastXferURL})

    cdcVFileMgmtEntry.EntityData.YListKeys = []string {"CdcVFileIndex", "CdcVFileMgmtIndex"}

    return &(cdcVFileMgmtEntry.EntityData)
}

// CISCODATACOLLECTIONMIB_CdcVFileMgmtTable_CdcVFileMgmtEntry_CdcVFileMgmtCommand represents                to abort an ongoing transfer.
type CISCODATACOLLECTIONMIB_CdcVFileMgmtTable_CdcVFileMgmtEntry_CdcVFileMgmtCommand string

const (
    CISCODATACOLLECTIONMIB_CdcVFileMgmtTable_CdcVFileMgmtEntry_CdcVFileMgmtCommand_idle CISCODATACOLLECTIONMIB_CdcVFileMgmtTable_CdcVFileMgmtEntry_CdcVFileMgmtCommand = "idle"

    CISCODATACOLLECTIONMIB_CdcVFileMgmtTable_CdcVFileMgmtEntry_CdcVFileMgmtCommand_delete_ CISCODATACOLLECTIONMIB_CdcVFileMgmtTable_CdcVFileMgmtEntry_CdcVFileMgmtCommand = "delete"

    CISCODATACOLLECTIONMIB_CdcVFileMgmtTable_CdcVFileMgmtEntry_CdcVFileMgmtCommand_transfer CISCODATACOLLECTIONMIB_CdcVFileMgmtTable_CdcVFileMgmtEntry_CdcVFileMgmtCommand = "transfer"

    CISCODATACOLLECTIONMIB_CdcVFileMgmtTable_CdcVFileMgmtEntry_CdcVFileMgmtCommand_abortTransfer CISCODATACOLLECTIONMIB_CdcVFileMgmtTable_CdcVFileMgmtEntry_CdcVFileMgmtCommand = "abortTransfer"
)

// CISCODATACOLLECTIONMIB_CdcDGTable
// A table for specifying data groups.
type CISCODATACOLLECTIONMIB_CdcDGTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table. Each entry corresponds to a data group. A data
    // group is used to select data that needs to be collected into VFiles. The
    // selection is done by specifying the base objects and their instances for
    // which the values need to be fetched.  Data is collected only for those data
    // groups, that have the corresponding instance of cdcDGRowStatus set to
    // 'active'.   In order for data to be collected, each data group has to be
    // associated with a cdcVFileEntry (see cdcDGVFileIndex). If the data
    // collection mode of the associated cdcVFileEntry is automatic, then data is
    // fetched and stored into the current VFile of the associated cdcVFileEntry
    // at periodic intervals (cdcDGPollPeriod). The type is slice of
    // CISCODATACOLLECTIONMIB_CdcDGTable_CdcDGEntry.
    CdcDGEntry []*CISCODATACOLLECTIONMIB_CdcDGTable_CdcDGEntry
}

func (cdcDGTable *CISCODATACOLLECTIONMIB_CdcDGTable) GetEntityData() *types.CommonEntityData {
    cdcDGTable.EntityData.YFilter = cdcDGTable.YFilter
    cdcDGTable.EntityData.YangName = "cdcDGTable"
    cdcDGTable.EntityData.BundleName = "cisco_ios_xe"
    cdcDGTable.EntityData.ParentYangName = "CISCO-DATA-COLLECTION-MIB"
    cdcDGTable.EntityData.SegmentPath = "cdcDGTable"
    cdcDGTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdcDGTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdcDGTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdcDGTable.EntityData.Children = types.NewOrderedMap()
    cdcDGTable.EntityData.Children.Append("cdcDGEntry", types.YChild{"CdcDGEntry", nil})
    for i := range cdcDGTable.CdcDGEntry {
        cdcDGTable.EntityData.Children.Append(types.GetSegmentPath(cdcDGTable.CdcDGEntry[i]), types.YChild{"CdcDGEntry", cdcDGTable.CdcDGEntry[i]})
    }
    cdcDGTable.EntityData.Leafs = types.NewOrderedMap()

    cdcDGTable.EntityData.YListKeys = []string {}

    return &(cdcDGTable.EntityData)
}

// CISCODATACOLLECTIONMIB_CdcDGTable_CdcDGEntry
// An entry in this table. Each entry corresponds to a data
// group. A data group is used to select data that needs to be
// collected into VFiles. The selection is done by specifying
// the base objects and their instances for which the values
// need to be fetched.
// 
// Data is collected only for those data groups, that have
// the corresponding instance of cdcDGRowStatus set to
// 'active'. 
// 
// In order for data to be collected, each data group has to
// be associated with a cdcVFileEntry (see cdcDGVFileIndex). If
// the data collection mode of the associated cdcVFileEntry is
// automatic, then data is fetched and stored into the current
// VFile of the associated cdcVFileEntry at periodic
// intervals (cdcDGPollPeriod).
type CISCODATACOLLECTIONMIB_CdcDGTable_CdcDGEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An arbitrary integer used to uniquely identify
    // this entry. When creating an entry, a management application should pick a
    // random number. The type is interface{} with range: 1..4294967295.
    CdcDGIndex interface{}

    // A descriptive string. This object's value may be modified at any time. The
    // type is string.
    CdcDGComment interface{}

    // Identifies the type of this data group. object         Data is a single MIB
    // object. The fully                instantiated OID is specified in          
    // cdcDGBaseObject.  table          Data is a logical table. The columns of   
    // this table correspond to the base objects                specified in
    // cdcDGBaseObjectTable, and the                rows correspond to the values
    // of the instances                specified in cdcDGInstanceTable.  This
    // object's value cannot be modified while the value of cdcDGRowStatus is
    // 'active'. The type is CdcDGType.
    CdcDGType interface{}

    // Corresponds to a value of cdcVFileIndex. It is used to associate this data
    // group with a cdcVFileEntry. The values of the base objects for  this data
    // group are stored into the current VFile of the associated cdcVFileEntry.  
    // This object's value may be modified at any time. The change takes effect
    // the next time data is fetched for this data group. The type is interface{}
    // with range: 1..4294967295.
    CdcDGVFileIndex interface{}

    // The tag for the target from which to obtain the data for this data group. 
    // A length of 0 indicates the local system.  In this case, access to the
    // objects of this data group is under the security credentials of the
    // requester that set cdcDGRowStatus to 'active'. Those credentials are the
    // input parameters for isAccessAllowed from the Architecture for Describing
    // SNMP Management Frameworks.  Otherwise a search is carried out for an entry
    // in the snmpTargetAddrTable whose snmpTargetAddrTagList contains the tag
    // specified by the value of this object. The security credentials
    // (snmpTargetParamsEntry) of the first entry that satisfies the above
    // criteria, are passed as input parameters for isAccessAllowed.   This
    // object's value may be modified at any time. The change takes effect the
    // next time data is fetched for this data group. The type is string.
    CdcDGTargetTag interface{}

    // The management context from which to obtain data for this data group.  This
    // object's value may be modified at any time. The change takes effect the
    // next time data is fetched for this data group. The type is string.
    CdcDGContextName interface{}

    // The fully instantiated name of the MIB object whose value needs to be
    // fetched. This object's value is used only when cdcDGType is of type
    // 'object'.   This object's value may be modified at any time. The change
    // takes effect the next time data is fetched for this data group. The type is
    // string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    CdcDGObject interface{}

    // Corresponds to a value of cdcDGBaseObjectGrpIndex, thus identifying a set
    // of entries in cdcDGBaseObjectTable, having this value of
    // cdcDGBaseObjectGrpIndex. This object's value is used only when cdcDGType is
    // of type 'table'.   This set of entries in cdcDGBaseObjectTable in turn 
    // identifies the set of base objects, that makes up the columns  of this
    // 'table' type data group.     This object's value may be modified at any
    // time. The change takes effect the next time data is fetched for this data
    // group. The type is interface{} with range: 1..4294967295.
    CdcDGObjectGrpIndex interface{}

    // Corresponds to a value of cdcDGInstanceGrpIndex, thus identifying a set of
    // entries in cdcDGInstanceTable, having this value of cdcDGInstanceGrpIndex.
    // This object's value is used only when cdcDGType is of type 'table'.   The
    // set of entries in cdcDGInstanceTable, in turn identifies the set of
    // instances of the base objects, whose values need to fetched. If the value
    // is 0, then all instances of the base objects will be fetched.    This
    // object's value may be modified at any time. The change takes effect the
    // next time data is fetched for this data group. The type is interface{} with
    // range: 0..4294967295.
    CdcDGInstGrpIndex interface{}

    // Specifies the time intervals at which the data should be fetched for this
    // data group. This object's value is used only when the collection mode of
    // the associated cdcVFileEntry is automatic (see cdcVFileCollectMode).   A
    // periodic timer is started for this data group when cdcDGRowStatus is set to
    // 'active', provided the associated cdcVFileEntry has already been
    // 'activated', otherwise it is started when the associated cdcVFileEntry is
    // finally activated.   The time interval after which the first expiration of
    // this timer should occur, is calculated as follows:  (value of sysUpTime.0)
    // +  (value of cdcPollPeriod for this entry -     (value of sysUpTime.0 -
    // VFile activation time for the     associated cdcVFileEntry) %
    // cdcPollPeriod)  Subsequent expirations of the periodic timer can occur as
    // per the value specified in cdcDGPollPeriod. This helps in synchronizing
    // periodic polling of the data groups with respect to the VFile activation
    // time.  This object's value may be modified at any time, and the change must
    // take effect immediately. i.e. if the periodic timer has been started, it's
    // expiry time may need to be re-adjusted. The type is interface{} with range:
    // 1..86400. Units are seconds.
    CdcDGPollPeriod interface{}

    // The status of this conceptual row.  This object cannot be set to 'active'
    // until values have been assigned to  cdcDGVFileIndex & cdcDGColGrpIndex. The
    // type is RowStatus.
    CdcDGRowStatus interface{}
}

func (cdcDGEntry *CISCODATACOLLECTIONMIB_CdcDGTable_CdcDGEntry) GetEntityData() *types.CommonEntityData {
    cdcDGEntry.EntityData.YFilter = cdcDGEntry.YFilter
    cdcDGEntry.EntityData.YangName = "cdcDGEntry"
    cdcDGEntry.EntityData.BundleName = "cisco_ios_xe"
    cdcDGEntry.EntityData.ParentYangName = "cdcDGTable"
    cdcDGEntry.EntityData.SegmentPath = "cdcDGEntry" + types.AddKeyToken(cdcDGEntry.CdcDGIndex, "cdcDGIndex")
    cdcDGEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdcDGEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdcDGEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdcDGEntry.EntityData.Children = types.NewOrderedMap()
    cdcDGEntry.EntityData.Leafs = types.NewOrderedMap()
    cdcDGEntry.EntityData.Leafs.Append("cdcDGIndex", types.YLeaf{"CdcDGIndex", cdcDGEntry.CdcDGIndex})
    cdcDGEntry.EntityData.Leafs.Append("cdcDGComment", types.YLeaf{"CdcDGComment", cdcDGEntry.CdcDGComment})
    cdcDGEntry.EntityData.Leafs.Append("cdcDGType", types.YLeaf{"CdcDGType", cdcDGEntry.CdcDGType})
    cdcDGEntry.EntityData.Leafs.Append("cdcDGVFileIndex", types.YLeaf{"CdcDGVFileIndex", cdcDGEntry.CdcDGVFileIndex})
    cdcDGEntry.EntityData.Leafs.Append("cdcDGTargetTag", types.YLeaf{"CdcDGTargetTag", cdcDGEntry.CdcDGTargetTag})
    cdcDGEntry.EntityData.Leafs.Append("cdcDGContextName", types.YLeaf{"CdcDGContextName", cdcDGEntry.CdcDGContextName})
    cdcDGEntry.EntityData.Leafs.Append("cdcDGObject", types.YLeaf{"CdcDGObject", cdcDGEntry.CdcDGObject})
    cdcDGEntry.EntityData.Leafs.Append("cdcDGObjectGrpIndex", types.YLeaf{"CdcDGObjectGrpIndex", cdcDGEntry.CdcDGObjectGrpIndex})
    cdcDGEntry.EntityData.Leafs.Append("cdcDGInstGrpIndex", types.YLeaf{"CdcDGInstGrpIndex", cdcDGEntry.CdcDGInstGrpIndex})
    cdcDGEntry.EntityData.Leafs.Append("cdcDGPollPeriod", types.YLeaf{"CdcDGPollPeriod", cdcDGEntry.CdcDGPollPeriod})
    cdcDGEntry.EntityData.Leafs.Append("cdcDGRowStatus", types.YLeaf{"CdcDGRowStatus", cdcDGEntry.CdcDGRowStatus})

    cdcDGEntry.EntityData.YListKeys = []string {"CdcDGIndex"}

    return &(cdcDGEntry.EntityData)
}

// CISCODATACOLLECTIONMIB_CdcDGTable_CdcDGEntry_CdcDGType represents cdcDGRowStatus is 'active'.
type CISCODATACOLLECTIONMIB_CdcDGTable_CdcDGEntry_CdcDGType string

const (
    CISCODATACOLLECTIONMIB_CdcDGTable_CdcDGEntry_CdcDGType_object CISCODATACOLLECTIONMIB_CdcDGTable_CdcDGEntry_CdcDGType = "object"

    CISCODATACOLLECTIONMIB_CdcDGTable_CdcDGEntry_CdcDGType_table CISCODATACOLLECTIONMIB_CdcDGTable_CdcDGEntry_CdcDGType = "table"
)

// CISCODATACOLLECTIONMIB_CdcDGBaseObjectTable
// A table specifying the base objects of a 'table' type
// data group.
type CISCODATACOLLECTIONMIB_CdcDGBaseObjectTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An individual entry in this table. Each entry is a  {subtree, list} tuple.
    // Each tuple identifies a set of  base objects for the associated data group.
    // The type is slice of
    // CISCODATACOLLECTIONMIB_CdcDGBaseObjectTable_CdcDGBaseObjectEntry.
    CdcDGBaseObjectEntry []*CISCODATACOLLECTIONMIB_CdcDGBaseObjectTable_CdcDGBaseObjectEntry
}

func (cdcDGBaseObjectTable *CISCODATACOLLECTIONMIB_CdcDGBaseObjectTable) GetEntityData() *types.CommonEntityData {
    cdcDGBaseObjectTable.EntityData.YFilter = cdcDGBaseObjectTable.YFilter
    cdcDGBaseObjectTable.EntityData.YangName = "cdcDGBaseObjectTable"
    cdcDGBaseObjectTable.EntityData.BundleName = "cisco_ios_xe"
    cdcDGBaseObjectTable.EntityData.ParentYangName = "CISCO-DATA-COLLECTION-MIB"
    cdcDGBaseObjectTable.EntityData.SegmentPath = "cdcDGBaseObjectTable"
    cdcDGBaseObjectTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdcDGBaseObjectTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdcDGBaseObjectTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdcDGBaseObjectTable.EntityData.Children = types.NewOrderedMap()
    cdcDGBaseObjectTable.EntityData.Children.Append("cdcDGBaseObjectEntry", types.YChild{"CdcDGBaseObjectEntry", nil})
    for i := range cdcDGBaseObjectTable.CdcDGBaseObjectEntry {
        cdcDGBaseObjectTable.EntityData.Children.Append(types.GetSegmentPath(cdcDGBaseObjectTable.CdcDGBaseObjectEntry[i]), types.YChild{"CdcDGBaseObjectEntry", cdcDGBaseObjectTable.CdcDGBaseObjectEntry[i]})
    }
    cdcDGBaseObjectTable.EntityData.Leafs = types.NewOrderedMap()

    cdcDGBaseObjectTable.EntityData.YListKeys = []string {}

    return &(cdcDGBaseObjectTable.EntityData)
}

// CISCODATACOLLECTIONMIB_CdcDGBaseObjectTable_CdcDGBaseObjectEntry
// An individual entry in this table. Each entry is a 
// {subtree, list} tuple. Each tuple identifies a set of 
// base objects for the associated data group.
type CISCODATACOLLECTIONMIB_CdcDGBaseObjectTable_CdcDGBaseObjectEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This object's value when combined with the value
    // of cdcDGBaseObjectIndex uniquely identifies an entry in this table. An
    // application must use the same value (can  be randomly picked) for this
    // object while creating a group of entries that collectively identifies the
    // set of base objects for a data group. The type is interface{} with range:
    // 1..4294967295.
    CdcDGBaseObjectGrpIndex interface{}

    // This attribute is a key. This object's value when combined with the value
    // of cdcDGBaseObjectGrpIndex uniquely identifies an entry in this table.  A
    // managment application can assign incremental values starting from one, when
    // creating each entry in a group of entries (as identified by the value of
    // cdcDGBaseObjectGrpIndex). The type is interface{} with range:
    // 1..4294967295.
    CdcDGBaseObjectIndex interface{}

    // The subtree component of a {subtree, list} tuple.  This object's value may
    // be modified at any time. The change takes effect the next time data is
    // fetched for this data group. The type is string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    CdcDGBaseObjectSubtree interface{}

    // The list component of a {subtree, list} tuple.  This object's value may be
    // modified at any time. The change takes effect the next time data is fetched
    // for this data group. The type is string with length: 0..16.
    CdcDGBaseObjectList interface{}

    // The status of this conceptual row.  This object cannot be set to 'active'
    // until values have been assigned to cdcDGBaseObjectSubtree &
    // cdcDGBaseObjectList. The type is RowStatus.
    CdcDGBaseObjectRowStatus interface{}
}

func (cdcDGBaseObjectEntry *CISCODATACOLLECTIONMIB_CdcDGBaseObjectTable_CdcDGBaseObjectEntry) GetEntityData() *types.CommonEntityData {
    cdcDGBaseObjectEntry.EntityData.YFilter = cdcDGBaseObjectEntry.YFilter
    cdcDGBaseObjectEntry.EntityData.YangName = "cdcDGBaseObjectEntry"
    cdcDGBaseObjectEntry.EntityData.BundleName = "cisco_ios_xe"
    cdcDGBaseObjectEntry.EntityData.ParentYangName = "cdcDGBaseObjectTable"
    cdcDGBaseObjectEntry.EntityData.SegmentPath = "cdcDGBaseObjectEntry" + types.AddKeyToken(cdcDGBaseObjectEntry.CdcDGBaseObjectGrpIndex, "cdcDGBaseObjectGrpIndex") + types.AddKeyToken(cdcDGBaseObjectEntry.CdcDGBaseObjectIndex, "cdcDGBaseObjectIndex")
    cdcDGBaseObjectEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdcDGBaseObjectEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdcDGBaseObjectEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdcDGBaseObjectEntry.EntityData.Children = types.NewOrderedMap()
    cdcDGBaseObjectEntry.EntityData.Leafs = types.NewOrderedMap()
    cdcDGBaseObjectEntry.EntityData.Leafs.Append("cdcDGBaseObjectGrpIndex", types.YLeaf{"CdcDGBaseObjectGrpIndex", cdcDGBaseObjectEntry.CdcDGBaseObjectGrpIndex})
    cdcDGBaseObjectEntry.EntityData.Leafs.Append("cdcDGBaseObjectIndex", types.YLeaf{"CdcDGBaseObjectIndex", cdcDGBaseObjectEntry.CdcDGBaseObjectIndex})
    cdcDGBaseObjectEntry.EntityData.Leafs.Append("cdcDGBaseObjectSubtree", types.YLeaf{"CdcDGBaseObjectSubtree", cdcDGBaseObjectEntry.CdcDGBaseObjectSubtree})
    cdcDGBaseObjectEntry.EntityData.Leafs.Append("cdcDGBaseObjectList", types.YLeaf{"CdcDGBaseObjectList", cdcDGBaseObjectEntry.CdcDGBaseObjectList})
    cdcDGBaseObjectEntry.EntityData.Leafs.Append("cdcDGBaseObjectRowStatus", types.YLeaf{"CdcDGBaseObjectRowStatus", cdcDGBaseObjectEntry.CdcDGBaseObjectRowStatus})

    cdcDGBaseObjectEntry.EntityData.YListKeys = []string {"CdcDGBaseObjectGrpIndex", "CdcDGBaseObjectIndex"}

    return &(cdcDGBaseObjectEntry.EntityData)
}

// CISCODATACOLLECTIONMIB_CdcDGInstanceTable
// Identifies the instances of the base objects that need to
// be fetched for a 'table' type data group.
// 
// The agent is not responsible for verifying that the instances
// specified for a data group do not overlap.
type CISCODATACOLLECTIONMIB_CdcDGInstanceTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table. Each entry identifies one or more instances of the
    // base objects that need to be fetched. An instance is represented by an OID
    // fragment. The type is slice of
    // CISCODATACOLLECTIONMIB_CdcDGInstanceTable_CdcDGInstanceEntry.
    CdcDGInstanceEntry []*CISCODATACOLLECTIONMIB_CdcDGInstanceTable_CdcDGInstanceEntry
}

func (cdcDGInstanceTable *CISCODATACOLLECTIONMIB_CdcDGInstanceTable) GetEntityData() *types.CommonEntityData {
    cdcDGInstanceTable.EntityData.YFilter = cdcDGInstanceTable.YFilter
    cdcDGInstanceTable.EntityData.YangName = "cdcDGInstanceTable"
    cdcDGInstanceTable.EntityData.BundleName = "cisco_ios_xe"
    cdcDGInstanceTable.EntityData.ParentYangName = "CISCO-DATA-COLLECTION-MIB"
    cdcDGInstanceTable.EntityData.SegmentPath = "cdcDGInstanceTable"
    cdcDGInstanceTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdcDGInstanceTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdcDGInstanceTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdcDGInstanceTable.EntityData.Children = types.NewOrderedMap()
    cdcDGInstanceTable.EntityData.Children.Append("cdcDGInstanceEntry", types.YChild{"CdcDGInstanceEntry", nil})
    for i := range cdcDGInstanceTable.CdcDGInstanceEntry {
        cdcDGInstanceTable.EntityData.Children.Append(types.GetSegmentPath(cdcDGInstanceTable.CdcDGInstanceEntry[i]), types.YChild{"CdcDGInstanceEntry", cdcDGInstanceTable.CdcDGInstanceEntry[i]})
    }
    cdcDGInstanceTable.EntityData.Leafs = types.NewOrderedMap()

    cdcDGInstanceTable.EntityData.YListKeys = []string {}

    return &(cdcDGInstanceTable.EntityData)
}

// CISCODATACOLLECTIONMIB_CdcDGInstanceTable_CdcDGInstanceEntry
// An entry in this table. Each entry identifies one or more
// instances of the base objects that need to be fetched.
// An instance is represented by an OID fragment.
type CISCODATACOLLECTIONMIB_CdcDGInstanceTable_CdcDGInstanceEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This object's value when combined with the value
    // of cdcDGInstanceIndex uniquely identifies an entry in this table. An
    // application must use the same value (can  be randomly picked) for this
    // object while creating a group of entries that collectively identifies the
    // set of instances for a data group. The type is interface{} with range:
    // 1..4294967295.
    CdcDGInstanceGrpIndex interface{}

    // This attribute is a key. This object's value when combined with the value
    // of cdcDGInstanceGrpIndex uniquely identifies an entry in this table.  A
    // managment application can assign incremental values starting from one, when
    // creating each entry within a group of entries (as identified by the value
    // of cdcDGInstanceGrpIndex). The type is interface{} with range:
    // 1..4294967295.
    CdcDGInstanceIndex interface{}

    // Specifies the way in which the instances are to be used while collecting
    // data.   individual     The value of cdcDGInstanceOid is               
    // appended to each base object of the                associated data group,
    // thus giving the exact                instance of the objects to be
    // collected.  range          The value of cdcDGInstanceOid is               
    // appended to each base object in the                associated data group,
    // thus giving the                starting object instance of the range.      
    // The value of cdcDGInstanceEndOid                is appended to to each base
    // object in the                associated data group, thus giving the        
    // last object instances of the range.   repititions      The value of
    // cdcDGInstanceOid is                appended to each base object in the     
    // associated data group, thus giving the                first object instance
    // of the next 'n'                instances that need to be collected.        
    // The value of 'n' is set in                cdcDGInstanceNumRepititions. 
    // subTree        The value of cdcDGInstanceOid is                appended to
    // each base object in the                associated data group, thus
    // identifying the                OBJECT IDENTFIFIER sub-tree, whose leaf     
    // instances need to be collected.  other          The value of
    // cdcDGInstanceOtherPtr points to a                row (in another MIB table)
    // that contains MIB                specific instance selection criteria. A
    // MIB                defined for such purposes should describe               
    // the selection criteria.  This object's value cannot be modified while the
    // value of cdcDGInstanceStatus is 'active'. The type is CdcDGInstanceType.
    CdcDGInstanceType interface{}

    // Contains the OBJECT IDENTIFIER fragment that identifies the instances of
    // the base objects that need to be collected.  If cdcDGInstanceType is
    // 'individual' then this value should be the OID fragment that, when appended
    // to each base MIB object gives the fully instantiated OID to be fetched.  If
    // cdcDGInstanceType is 'range' then this value should be the OID fragment
    // that, when appended to each base MIB object gives the start of a range of
    // object instances that needs to be fetched.  If cdcDGInstanceType is
    // 'subTree' then this value should be the OID fragment that, when appended to
    // each base MIB gives the sub-tree under which all leaf object instances need
    // to be fetched.  This object's value may be modified at any time. The change
    // takes effect the next time data is fetched for this data group. The type is
    // string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    CdcDGInstanceOid interface{}

    // Contains the OID fragment that, when appended to each base object gives the
    // end of the range of object instances that needs to be fetched.  This value
    // is used only when the value of cdcDGInstanceType is of type 'range'.   This
    // object's value may be modified at any time. The change takes effect the
    // next time data is fetched for this data group. The type is string with
    // pattern: (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    CdcDGInstanceOidEnd interface{}

    // Specifies the number of lexicographically consecutive object instances to
    // fetch.  This value is used only when the value of cdcDGInstanceType is of
    // type 'repititions'.    This object's value may be modified at any time. The
    // change takes effect the next time data is fetched for this data group. The
    // type is interface{} with range: 0..4294967295.
    CdcDGInstanceNumRepititions interface{}

    // Contains a pointer to a row in another MIB table that contains MIB specific
    // criteria for selecting instances.  This value is used only when the value
    // of cdcDGInstanceType is of type 'other'.   This object's value may be
    // modified at any time. The change takes effect the next time data is fetched
    // for this data group. The type is string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    CdcDGInstanceOtherPtr interface{}

    // The status of this conceptual row. The type is RowStatus.
    CdcDGInstanceRowStatus interface{}
}

func (cdcDGInstanceEntry *CISCODATACOLLECTIONMIB_CdcDGInstanceTable_CdcDGInstanceEntry) GetEntityData() *types.CommonEntityData {
    cdcDGInstanceEntry.EntityData.YFilter = cdcDGInstanceEntry.YFilter
    cdcDGInstanceEntry.EntityData.YangName = "cdcDGInstanceEntry"
    cdcDGInstanceEntry.EntityData.BundleName = "cisco_ios_xe"
    cdcDGInstanceEntry.EntityData.ParentYangName = "cdcDGInstanceTable"
    cdcDGInstanceEntry.EntityData.SegmentPath = "cdcDGInstanceEntry" + types.AddKeyToken(cdcDGInstanceEntry.CdcDGInstanceGrpIndex, "cdcDGInstanceGrpIndex") + types.AddKeyToken(cdcDGInstanceEntry.CdcDGInstanceIndex, "cdcDGInstanceIndex")
    cdcDGInstanceEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdcDGInstanceEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdcDGInstanceEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdcDGInstanceEntry.EntityData.Children = types.NewOrderedMap()
    cdcDGInstanceEntry.EntityData.Leafs = types.NewOrderedMap()
    cdcDGInstanceEntry.EntityData.Leafs.Append("cdcDGInstanceGrpIndex", types.YLeaf{"CdcDGInstanceGrpIndex", cdcDGInstanceEntry.CdcDGInstanceGrpIndex})
    cdcDGInstanceEntry.EntityData.Leafs.Append("cdcDGInstanceIndex", types.YLeaf{"CdcDGInstanceIndex", cdcDGInstanceEntry.CdcDGInstanceIndex})
    cdcDGInstanceEntry.EntityData.Leafs.Append("cdcDGInstanceType", types.YLeaf{"CdcDGInstanceType", cdcDGInstanceEntry.CdcDGInstanceType})
    cdcDGInstanceEntry.EntityData.Leafs.Append("cdcDGInstanceOid", types.YLeaf{"CdcDGInstanceOid", cdcDGInstanceEntry.CdcDGInstanceOid})
    cdcDGInstanceEntry.EntityData.Leafs.Append("cdcDGInstanceOidEnd", types.YLeaf{"CdcDGInstanceOidEnd", cdcDGInstanceEntry.CdcDGInstanceOidEnd})
    cdcDGInstanceEntry.EntityData.Leafs.Append("cdcDGInstanceNumRepititions", types.YLeaf{"CdcDGInstanceNumRepititions", cdcDGInstanceEntry.CdcDGInstanceNumRepititions})
    cdcDGInstanceEntry.EntityData.Leafs.Append("cdcDGInstanceOtherPtr", types.YLeaf{"CdcDGInstanceOtherPtr", cdcDGInstanceEntry.CdcDGInstanceOtherPtr})
    cdcDGInstanceEntry.EntityData.Leafs.Append("cdcDGInstanceRowStatus", types.YLeaf{"CdcDGInstanceRowStatus", cdcDGInstanceEntry.CdcDGInstanceRowStatus})

    cdcDGInstanceEntry.EntityData.YListKeys = []string {"CdcDGInstanceGrpIndex", "CdcDGInstanceIndex"}

    return &(cdcDGInstanceEntry.EntityData)
}

// CISCODATACOLLECTIONMIB_CdcDGInstanceTable_CdcDGInstanceEntry_CdcDGInstanceType represents cdcDGInstanceStatus is 'active'.
type CISCODATACOLLECTIONMIB_CdcDGInstanceTable_CdcDGInstanceEntry_CdcDGInstanceType string

const (
    CISCODATACOLLECTIONMIB_CdcDGInstanceTable_CdcDGInstanceEntry_CdcDGInstanceType_individual CISCODATACOLLECTIONMIB_CdcDGInstanceTable_CdcDGInstanceEntry_CdcDGInstanceType = "individual"

    CISCODATACOLLECTIONMIB_CdcDGInstanceTable_CdcDGInstanceEntry_CdcDGInstanceType_range_ CISCODATACOLLECTIONMIB_CdcDGInstanceTable_CdcDGInstanceEntry_CdcDGInstanceType = "range"

    CISCODATACOLLECTIONMIB_CdcDGInstanceTable_CdcDGInstanceEntry_CdcDGInstanceType_repititions CISCODATACOLLECTIONMIB_CdcDGInstanceTable_CdcDGInstanceEntry_CdcDGInstanceType = "repititions"

    CISCODATACOLLECTIONMIB_CdcDGInstanceTable_CdcDGInstanceEntry_CdcDGInstanceType_subTree CISCODATACOLLECTIONMIB_CdcDGInstanceTable_CdcDGInstanceEntry_CdcDGInstanceType = "subTree"

    CISCODATACOLLECTIONMIB_CdcDGInstanceTable_CdcDGInstanceEntry_CdcDGInstanceType_other CISCODATACOLLECTIONMIB_CdcDGInstanceTable_CdcDGInstanceEntry_CdcDGInstanceType = "other"
)

// CISCODATACOLLECTIONMIB_CdcFileXferConfTable
// A table for configuring file transfer operations.
type CISCODATACOLLECTIONMIB_CdcFileXferConfTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An individual entry in the cdcFileXferConfTable. Each entry identifies a
    // primary and an optional secondary destination.  An entry is automatically
    // created in this table, whenever an entry is created in the cdcVFileTable.
    // The application needs to specify the URLs of the destination to which
    // frozen VFiles are transferred.   When a VFile is frozen, transfer will be
    // first initiated to the primary destination, if the transfer fails, then
    // transfer is initiated to the secondary destination. If this too fails, then
    // the cycle is repeated again after a specified time period (value of
    // cdcFileXferConfRetryPeriod) elapses. The type is slice of
    // CISCODATACOLLECTIONMIB_CdcFileXferConfTable_CdcFileXferConfEntry.
    CdcFileXferConfEntry []*CISCODATACOLLECTIONMIB_CdcFileXferConfTable_CdcFileXferConfEntry
}

func (cdcFileXferConfTable *CISCODATACOLLECTIONMIB_CdcFileXferConfTable) GetEntityData() *types.CommonEntityData {
    cdcFileXferConfTable.EntityData.YFilter = cdcFileXferConfTable.YFilter
    cdcFileXferConfTable.EntityData.YangName = "cdcFileXferConfTable"
    cdcFileXferConfTable.EntityData.BundleName = "cisco_ios_xe"
    cdcFileXferConfTable.EntityData.ParentYangName = "CISCO-DATA-COLLECTION-MIB"
    cdcFileXferConfTable.EntityData.SegmentPath = "cdcFileXferConfTable"
    cdcFileXferConfTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdcFileXferConfTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdcFileXferConfTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdcFileXferConfTable.EntityData.Children = types.NewOrderedMap()
    cdcFileXferConfTable.EntityData.Children.Append("cdcFileXferConfEntry", types.YChild{"CdcFileXferConfEntry", nil})
    for i := range cdcFileXferConfTable.CdcFileXferConfEntry {
        cdcFileXferConfTable.EntityData.Children.Append(types.GetSegmentPath(cdcFileXferConfTable.CdcFileXferConfEntry[i]), types.YChild{"CdcFileXferConfEntry", cdcFileXferConfTable.CdcFileXferConfEntry[i]})
    }
    cdcFileXferConfTable.EntityData.Leafs = types.NewOrderedMap()

    cdcFileXferConfTable.EntityData.YListKeys = []string {}

    return &(cdcFileXferConfTable.EntityData)
}

// CISCODATACOLLECTIONMIB_CdcFileXferConfTable_CdcFileXferConfEntry
// An individual entry in the cdcFileXferConfTable. Each entry
// identifies a primary and an optional secondary destination.
// 
// An entry is automatically created in this table, whenever an
// entry is created in the cdcVFileTable. The application needs
// to specify the URLs of the destination to which frozen VFiles
// are transferred. 
// 
// When a VFile is frozen, transfer will be first initiated to
// the primary destination, if the transfer fails, then transfer
// is initiated to the secondary destination. If this too fails,
// then the cycle is repeated again after a specified time
// period (value of cdcFileXferConfRetryPeriod) elapses.
type CISCODATACOLLECTIONMIB_CdcFileXferConfTable_CdcFileXferConfEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // cisco_data_collection_mib.CISCODATACOLLECTIONMIB_CdcVFileTable_CdcVFileEntry_CdcVFileIndex
    CdcVFileIndex interface{}

    // The URL which specifies the primary destination to which the file has to be
    // transferred. The URL should contain the base-name of the remote file, the
    // suffix will be carried over from the name of the VFile being tranferred,
    // and will be automatically appended by the agent.  This object's value may
    // be modified at any time. The type is string with length: 0..255.
    CdcFileXferConfPriUrl interface{}

    // The URL which specifies the secondary destination to which the file has to
    // be transferred if the transfer to the  primary destination fails. Failure
    // occurs when the file  cannot be transferred in it's entirety to the
    // specified  destination for some reason. Some common reasons for such 
    // failures are listed out in CdcFileXferStatus.    The specified URL should
    // contain the base-name of the remote file, the suffix will be carried over
    // from the name of the VFile being tranferred, and will be automatically
    // appended by the agent.   This object's value may be modified at any time.
    // The type is string with length: 0..255.
    CdcFileXferConfSecUrl interface{}

    // Specifies the time interval after which transfer has to be retried.
    // Transfer needs to be retried only if in a previous  attempt the file could
    // not be successfully transferred to  either the primary destination or the
    // secondary destination.  This object's value may be modified at any time.
    // The type is interface{} with range: 60..86400. Units are seconds.
    CdcFileXferConfRetryPeriod interface{}

    // Maximum number of times, transfer has to be retried. If the retry count
    // exceeds this value, then no further attempts will be made.   This object's
    // value may be modified at any time. The type is interface{} with range:
    // 0..256. Units are seconds.
    CdcFileXferConfRetryCount interface{}

    // When set to 'true', cdcFileXferComplete notification will be sent out in
    // the event of a successful file transfer. The type is bool.
    CdcFileXferConfSuccessEnable interface{}

    // When set to 'true', cdcFileXferComplete notification will be sent out in
    // the event of a file transfer failure. The type is bool.
    CdcFileXferConfFailureEnable interface{}
}

func (cdcFileXferConfEntry *CISCODATACOLLECTIONMIB_CdcFileXferConfTable_CdcFileXferConfEntry) GetEntityData() *types.CommonEntityData {
    cdcFileXferConfEntry.EntityData.YFilter = cdcFileXferConfEntry.YFilter
    cdcFileXferConfEntry.EntityData.YangName = "cdcFileXferConfEntry"
    cdcFileXferConfEntry.EntityData.BundleName = "cisco_ios_xe"
    cdcFileXferConfEntry.EntityData.ParentYangName = "cdcFileXferConfTable"
    cdcFileXferConfEntry.EntityData.SegmentPath = "cdcFileXferConfEntry" + types.AddKeyToken(cdcFileXferConfEntry.CdcVFileIndex, "cdcVFileIndex")
    cdcFileXferConfEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdcFileXferConfEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdcFileXferConfEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdcFileXferConfEntry.EntityData.Children = types.NewOrderedMap()
    cdcFileXferConfEntry.EntityData.Leafs = types.NewOrderedMap()
    cdcFileXferConfEntry.EntityData.Leafs.Append("cdcVFileIndex", types.YLeaf{"CdcVFileIndex", cdcFileXferConfEntry.CdcVFileIndex})
    cdcFileXferConfEntry.EntityData.Leafs.Append("cdcFileXferConfPriUrl", types.YLeaf{"CdcFileXferConfPriUrl", cdcFileXferConfEntry.CdcFileXferConfPriUrl})
    cdcFileXferConfEntry.EntityData.Leafs.Append("cdcFileXferConfSecUrl", types.YLeaf{"CdcFileXferConfSecUrl", cdcFileXferConfEntry.CdcFileXferConfSecUrl})
    cdcFileXferConfEntry.EntityData.Leafs.Append("cdcFileXferConfRetryPeriod", types.YLeaf{"CdcFileXferConfRetryPeriod", cdcFileXferConfEntry.CdcFileXferConfRetryPeriod})
    cdcFileXferConfEntry.EntityData.Leafs.Append("cdcFileXferConfRetryCount", types.YLeaf{"CdcFileXferConfRetryCount", cdcFileXferConfEntry.CdcFileXferConfRetryCount})
    cdcFileXferConfEntry.EntityData.Leafs.Append("cdcFileXferConfSuccessEnable", types.YLeaf{"CdcFileXferConfSuccessEnable", cdcFileXferConfEntry.CdcFileXferConfSuccessEnable})
    cdcFileXferConfEntry.EntityData.Leafs.Append("cdcFileXferConfFailureEnable", types.YLeaf{"CdcFileXferConfFailureEnable", cdcFileXferConfEntry.CdcFileXferConfFailureEnable})

    cdcFileXferConfEntry.EntityData.YListKeys = []string {"CdcVFileIndex"}

    return &(cdcFileXferConfEntry.EntityData)
}

