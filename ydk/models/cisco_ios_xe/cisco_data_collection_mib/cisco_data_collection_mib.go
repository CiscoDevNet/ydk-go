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

// CdcFileFormat represents  immediately precedes the next tag or the end of file.
type CdcFileFormat string

const (
    CdcFileFormat_cdcBulkASCII CdcFileFormat = "cdcBulkASCII"

    CdcFileFormat_cdcBulkBinary CdcFileFormat = "cdcBulkBinary"

    CdcFileFormat_cdcSchemaASCII CdcFileFormat = "cdcSchemaASCII"
)

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

// CISCODATACOLLECTIONMIB
type CISCODATACOLLECTIONMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Cdcvfile CISCODATACOLLECTIONMIB_Cdcvfile

    // A table for setting up VFiles for collecting data.
    Cdcvfiletable CISCODATACOLLECTIONMIB_Cdcvfiletable

    // A table to manage frozen VFiles.
    Cdcvfilemgmttable CISCODATACOLLECTIONMIB_Cdcvfilemgmttable

    // A table for specifying data groups.
    Cdcdgtable CISCODATACOLLECTIONMIB_Cdcdgtable

    // A table specifying the base objects of a 'table' type data group.
    Cdcdgbaseobjecttable CISCODATACOLLECTIONMIB_Cdcdgbaseobjecttable

    // Identifies the instances of the base objects that need to be fetched for a
    // 'table' type data group.  The agent is not responsible for verifying that
    // the instances specified for a data group do not overlap.
    Cdcdginstancetable CISCODATACOLLECTIONMIB_Cdcdginstancetable

    // A table for configuring file transfer operations.
    Cdcfilexferconftable CISCODATACOLLECTIONMIB_Cdcfilexferconftable
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

    cISCODATACOLLECTIONMIB.EntityData.Children = make(map[string]types.YChild)
    cISCODATACOLLECTIONMIB.EntityData.Children["cdcVFile"] = types.YChild{"Cdcvfile", &cISCODATACOLLECTIONMIB.Cdcvfile}
    cISCODATACOLLECTIONMIB.EntityData.Children["cdcVFileTable"] = types.YChild{"Cdcvfiletable", &cISCODATACOLLECTIONMIB.Cdcvfiletable}
    cISCODATACOLLECTIONMIB.EntityData.Children["cdcVFileMgmtTable"] = types.YChild{"Cdcvfilemgmttable", &cISCODATACOLLECTIONMIB.Cdcvfilemgmttable}
    cISCODATACOLLECTIONMIB.EntityData.Children["cdcDGTable"] = types.YChild{"Cdcdgtable", &cISCODATACOLLECTIONMIB.Cdcdgtable}
    cISCODATACOLLECTIONMIB.EntityData.Children["cdcDGBaseObjectTable"] = types.YChild{"Cdcdgbaseobjecttable", &cISCODATACOLLECTIONMIB.Cdcdgbaseobjecttable}
    cISCODATACOLLECTIONMIB.EntityData.Children["cdcDGInstanceTable"] = types.YChild{"Cdcdginstancetable", &cISCODATACOLLECTIONMIB.Cdcdginstancetable}
    cISCODATACOLLECTIONMIB.EntityData.Children["cdcFileXferConfTable"] = types.YChild{"Cdcfilexferconftable", &cISCODATACOLLECTIONMIB.Cdcfilexferconftable}
    cISCODATACOLLECTIONMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCODATACOLLECTIONMIB.EntityData)
}

// CISCODATACOLLECTIONMIB_Cdcvfile
type CISCODATACOLLECTIONMIB_Cdcvfile struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object's value reads 'true', if the agent implementation of this MIB
    // supports placement of VFiles in application specified persistent storage
    // locations. Otherwise  the value is 'false'. The type is bool.
    Cdcvfilepersistentstorage interface{}

    // A global limit for the number of consecutive times the maximum VFile size
    // (cdcVFileMaxSize) is exceeded for a given VFile. When this limit is
    // exceeded the offending cdcVFileEntry is moved to the error state (see
    // cdcVFileOperStatus). The type is interface{} with range: 1..4294967295.
    Cdcvfilemaxsizehitslimit interface{}
}

func (cdcvfile *CISCODATACOLLECTIONMIB_Cdcvfile) GetEntityData() *types.CommonEntityData {
    cdcvfile.EntityData.YFilter = cdcvfile.YFilter
    cdcvfile.EntityData.YangName = "cdcVFile"
    cdcvfile.EntityData.BundleName = "cisco_ios_xe"
    cdcvfile.EntityData.ParentYangName = "CISCO-DATA-COLLECTION-MIB"
    cdcvfile.EntityData.SegmentPath = "cdcVFile"
    cdcvfile.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdcvfile.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdcvfile.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdcvfile.EntityData.Children = make(map[string]types.YChild)
    cdcvfile.EntityData.Leafs = make(map[string]types.YLeaf)
    cdcvfile.EntityData.Leafs["cdcVFilePersistentStorage"] = types.YLeaf{"Cdcvfilepersistentstorage", cdcvfile.Cdcvfilepersistentstorage}
    cdcvfile.EntityData.Leafs["cdcVFileMaxSizeHitsLimit"] = types.YLeaf{"Cdcvfilemaxsizehitslimit", cdcvfile.Cdcvfilemaxsizehitslimit}
    return &(cdcvfile.EntityData)
}

// CISCODATACOLLECTIONMIB_Cdcvfiletable
// A table for setting up VFiles for collecting data.
type CISCODATACOLLECTIONMIB_Cdcvfiletable struct {
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
    // CISCODATACOLLECTIONMIB_Cdcvfiletable_Cdcvfileentry.
    Cdcvfileentry []CISCODATACOLLECTIONMIB_Cdcvfiletable_Cdcvfileentry
}

func (cdcvfiletable *CISCODATACOLLECTIONMIB_Cdcvfiletable) GetEntityData() *types.CommonEntityData {
    cdcvfiletable.EntityData.YFilter = cdcvfiletable.YFilter
    cdcvfiletable.EntityData.YangName = "cdcVFileTable"
    cdcvfiletable.EntityData.BundleName = "cisco_ios_xe"
    cdcvfiletable.EntityData.ParentYangName = "CISCO-DATA-COLLECTION-MIB"
    cdcvfiletable.EntityData.SegmentPath = "cdcVFileTable"
    cdcvfiletable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdcvfiletable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdcvfiletable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdcvfiletable.EntityData.Children = make(map[string]types.YChild)
    cdcvfiletable.EntityData.Children["cdcVFileEntry"] = types.YChild{"Cdcvfileentry", nil}
    for i := range cdcvfiletable.Cdcvfileentry {
        cdcvfiletable.EntityData.Children[types.GetSegmentPath(&cdcvfiletable.Cdcvfileentry[i])] = types.YChild{"Cdcvfileentry", &cdcvfiletable.Cdcvfileentry[i]}
    }
    cdcvfiletable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cdcvfiletable.EntityData)
}

// CISCODATACOLLECTIONMIB_Cdcvfiletable_Cdcvfileentry
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
type CISCODATACOLLECTIONMIB_Cdcvfiletable_Cdcvfileentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An arbitrary integer for uniquely identifying this
    // entry. When creating a row, the application should pick a random number.  
    // If the configuration in this entry is persisted across system/agent
    // restarts then the same value of cdcVFileIndex must be assigned to this
    // entry after the restart. The type is interface{} with range: 1..4294967295.
    Cdcvfileindex interface{}

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
    Cdcvfilename interface{}

    // A string that can be used for administrative purposes.  This object's value
    // may be modified at any time. The type is string.
    Cdcvfiledescription interface{}

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
    // cdcVFileCollectMode). The type is Cdcvfilecommand.
    Cdcvfilecommand interface{}

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
    Cdcvfilemaxsize interface{}

    // The size of the current VFile. The type is interface{} with range:
    // 0..4294967295. Units are bytes.
    Cdcvfilecurrentsize interface{}

    // The format in which data is stored into the current VFile.  This object's
    // value cannot be modified while the entry is in the 'activated' state. The
    // type is CdcFileFormat.
    Cdcvfileformat interface{}

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
    // Cdcvfilecollectmode.
    Cdcvfilecollectmode interface{}

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
    Cdcvfilecollectionperiod interface{}

    // The time for which a frozen VFile is retained by the agent. When a VFile is
    // frozen, a timer (one per frozen VFile) is started to keep track of the
    // retention period for the  VFile. Once this timer expires, the VFile is
    // deleted. Till the expiry of the retention period, information  about frozen
    // VFiles is maintained in  cdcVFileMgmtTable.  This object's value may be
    // modified at any time, however the new value will take effect only for new
    // frozen VFiles. The type is interface{} with range: 60..86400. Units are
    // seconds.
    Cdcvfileretentionperiod interface{}

    // A control object to indicate the administratively desired state of data
    // collection for this entry. On setting the value to 'disabled' data
    // collection operations for this  entry are stopped, the current VFile is
    // frozen and it's  transfer is initiated.   Modifying the value of
    // cdcVFileAdminStatus to 'disabled' does not remove or change the current
    // configuration as represented by the active rows in this table. The type is
    // Cdcvfileadminstatus.
    Cdcvfileadminstatus interface{}

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
    // this entry. The type is Cdcvfileoperstatus.
    Cdcvfileoperstatus interface{}

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
    // dynamic memory. The type is Cdcvfileerrorcode.
    Cdcvfileerrorcode interface{}

    // When set to 'true', cdcVFileCollectionError notification will be sent out
    // in the event of a data collection error. The type is bool.
    Cdcvfilecollectionerrorenable interface{}

    // The status of this conceptual row.  A valid cdcVFileName is only mandatory
    // object for setting this object to 'active'. But collection of data in to
    // active vfile starts only on setting cdcVFileAdminStatus  to 'enabled'.
    // Setting this object to 'destroy' stops all data collection operations for
    // this entry, deletes all VFiles and removes this entry from cdcVFileTable.
    // The type is RowStatus.
    Cdcvfilerowstatus interface{}
}

func (cdcvfileentry *CISCODATACOLLECTIONMIB_Cdcvfiletable_Cdcvfileentry) GetEntityData() *types.CommonEntityData {
    cdcvfileentry.EntityData.YFilter = cdcvfileentry.YFilter
    cdcvfileentry.EntityData.YangName = "cdcVFileEntry"
    cdcvfileentry.EntityData.BundleName = "cisco_ios_xe"
    cdcvfileentry.EntityData.ParentYangName = "cdcVFileTable"
    cdcvfileentry.EntityData.SegmentPath = "cdcVFileEntry" + "[cdcVFileIndex='" + fmt.Sprintf("%v", cdcvfileentry.Cdcvfileindex) + "']"
    cdcvfileentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdcvfileentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdcvfileentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdcvfileentry.EntityData.Children = make(map[string]types.YChild)
    cdcvfileentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cdcvfileentry.EntityData.Leafs["cdcVFileIndex"] = types.YLeaf{"Cdcvfileindex", cdcvfileentry.Cdcvfileindex}
    cdcvfileentry.EntityData.Leafs["cdcVFileName"] = types.YLeaf{"Cdcvfilename", cdcvfileentry.Cdcvfilename}
    cdcvfileentry.EntityData.Leafs["cdcVFileDescription"] = types.YLeaf{"Cdcvfiledescription", cdcvfileentry.Cdcvfiledescription}
    cdcvfileentry.EntityData.Leafs["cdcVFileCommand"] = types.YLeaf{"Cdcvfilecommand", cdcvfileentry.Cdcvfilecommand}
    cdcvfileentry.EntityData.Leafs["cdcVFileMaxSize"] = types.YLeaf{"Cdcvfilemaxsize", cdcvfileentry.Cdcvfilemaxsize}
    cdcvfileentry.EntityData.Leafs["cdcVFileCurrentSize"] = types.YLeaf{"Cdcvfilecurrentsize", cdcvfileentry.Cdcvfilecurrentsize}
    cdcvfileentry.EntityData.Leafs["cdcVFileFormat"] = types.YLeaf{"Cdcvfileformat", cdcvfileentry.Cdcvfileformat}
    cdcvfileentry.EntityData.Leafs["cdcVFileCollectMode"] = types.YLeaf{"Cdcvfilecollectmode", cdcvfileentry.Cdcvfilecollectmode}
    cdcvfileentry.EntityData.Leafs["cdcVFileCollectionPeriod"] = types.YLeaf{"Cdcvfilecollectionperiod", cdcvfileentry.Cdcvfilecollectionperiod}
    cdcvfileentry.EntityData.Leafs["cdcVFileRetentionPeriod"] = types.YLeaf{"Cdcvfileretentionperiod", cdcvfileentry.Cdcvfileretentionperiod}
    cdcvfileentry.EntityData.Leafs["cdcVFileAdminStatus"] = types.YLeaf{"Cdcvfileadminstatus", cdcvfileentry.Cdcvfileadminstatus}
    cdcvfileentry.EntityData.Leafs["cdcVFileOperStatus"] = types.YLeaf{"Cdcvfileoperstatus", cdcvfileentry.Cdcvfileoperstatus}
    cdcvfileentry.EntityData.Leafs["cdcVFileErrorCode"] = types.YLeaf{"Cdcvfileerrorcode", cdcvfileentry.Cdcvfileerrorcode}
    cdcvfileentry.EntityData.Leafs["cdcVFileCollectionErrorEnable"] = types.YLeaf{"Cdcvfilecollectionerrorenable", cdcvfileentry.Cdcvfilecollectionerrorenable}
    cdcvfileentry.EntityData.Leafs["cdcVFileRowStatus"] = types.YLeaf{"Cdcvfilerowstatus", cdcvfileentry.Cdcvfilerowstatus}
    return &(cdcvfileentry.EntityData)
}

// CISCODATACOLLECTIONMIB_Cdcvfiletable_Cdcvfileentry_Cdcvfileadminstatus represents by the active rows in this table.
type CISCODATACOLLECTIONMIB_Cdcvfiletable_Cdcvfileentry_Cdcvfileadminstatus string

const (
    CISCODATACOLLECTIONMIB_Cdcvfiletable_Cdcvfileentry_Cdcvfileadminstatus_enabled CISCODATACOLLECTIONMIB_Cdcvfiletable_Cdcvfileentry_Cdcvfileadminstatus = "enabled"

    CISCODATACOLLECTIONMIB_Cdcvfiletable_Cdcvfileentry_Cdcvfileadminstatus_disabled CISCODATACOLLECTIONMIB_Cdcvfiletable_Cdcvfileentry_Cdcvfileadminstatus = "disabled"
)

// CISCODATACOLLECTIONMIB_Cdcvfiletable_Cdcvfileentry_Cdcvfilecollectmode represents is in the 'activated' state.
type CISCODATACOLLECTIONMIB_Cdcvfiletable_Cdcvfileentry_Cdcvfilecollectmode string

const (
    CISCODATACOLLECTIONMIB_Cdcvfiletable_Cdcvfileentry_Cdcvfilecollectmode_auto CISCODATACOLLECTIONMIB_Cdcvfiletable_Cdcvfileentry_Cdcvfilecollectmode = "auto"

    CISCODATACOLLECTIONMIB_Cdcvfiletable_Cdcvfileentry_Cdcvfilecollectmode_manual CISCODATACOLLECTIONMIB_Cdcvfiletable_Cdcvfileentry_Cdcvfilecollectmode = "manual"
)

// CISCODATACOLLECTIONMIB_Cdcvfiletable_Cdcvfileentry_Cdcvfilecommand represents                   cdcVFileCollectMode).
type CISCODATACOLLECTIONMIB_Cdcvfiletable_Cdcvfileentry_Cdcvfilecommand string

const (
    CISCODATACOLLECTIONMIB_Cdcvfiletable_Cdcvfileentry_Cdcvfilecommand_idle CISCODATACOLLECTIONMIB_Cdcvfiletable_Cdcvfileentry_Cdcvfilecommand = "idle"

    CISCODATACOLLECTIONMIB_Cdcvfiletable_Cdcvfileentry_Cdcvfilecommand_swapToNewFile CISCODATACOLLECTIONMIB_Cdcvfiletable_Cdcvfileentry_Cdcvfilecommand = "swapToNewFile"

    CISCODATACOLLECTIONMIB_Cdcvfiletable_Cdcvfileentry_Cdcvfilecommand_collectNow CISCODATACOLLECTIONMIB_Cdcvfiletable_Cdcvfileentry_Cdcvfilecommand = "collectNow"
)

// CISCODATACOLLECTIONMIB_Cdcvfiletable_Cdcvfileentry_Cdcvfileerrorcode represents                        e.g. unavailability of dynamic memory.
type CISCODATACOLLECTIONMIB_Cdcvfiletable_Cdcvfileentry_Cdcvfileerrorcode string

const (
    CISCODATACOLLECTIONMIB_Cdcvfiletable_Cdcvfileentry_Cdcvfileerrorcode_noError CISCODATACOLLECTIONMIB_Cdcvfiletable_Cdcvfileentry_Cdcvfileerrorcode = "noError"

    CISCODATACOLLECTIONMIB_Cdcvfiletable_Cdcvfileentry_Cdcvfileerrorcode_otherError CISCODATACOLLECTIONMIB_Cdcvfiletable_Cdcvfileentry_Cdcvfileerrorcode = "otherError"

    CISCODATACOLLECTIONMIB_Cdcvfiletable_Cdcvfileentry_Cdcvfileerrorcode_noSpace CISCODATACOLLECTIONMIB_Cdcvfiletable_Cdcvfileentry_Cdcvfileerrorcode = "noSpace"

    CISCODATACOLLECTIONMIB_Cdcvfiletable_Cdcvfileentry_Cdcvfileerrorcode_openError CISCODATACOLLECTIONMIB_Cdcvfiletable_Cdcvfileentry_Cdcvfileerrorcode = "openError"

    CISCODATACOLLECTIONMIB_Cdcvfiletable_Cdcvfileentry_Cdcvfileerrorcode_tooSmallMaxSize CISCODATACOLLECTIONMIB_Cdcvfiletable_Cdcvfileentry_Cdcvfileerrorcode = "tooSmallMaxSize"

    CISCODATACOLLECTIONMIB_Cdcvfiletable_Cdcvfileentry_Cdcvfileerrorcode_tooManyMaxSizeHits CISCODATACOLLECTIONMIB_Cdcvfiletable_Cdcvfileentry_Cdcvfileerrorcode = "tooManyMaxSizeHits"

    CISCODATACOLLECTIONMIB_Cdcvfiletable_Cdcvfileentry_Cdcvfileerrorcode_noResource CISCODATACOLLECTIONMIB_Cdcvfiletable_Cdcvfileentry_Cdcvfileerrorcode = "noResource"
)

// CISCODATACOLLECTIONMIB_Cdcvfiletable_Cdcvfileentry_Cdcvfileoperstatus represents this entry.
type CISCODATACOLLECTIONMIB_Cdcvfiletable_Cdcvfileentry_Cdcvfileoperstatus string

const (
    CISCODATACOLLECTIONMIB_Cdcvfiletable_Cdcvfileentry_Cdcvfileoperstatus_enabled CISCODATACOLLECTIONMIB_Cdcvfiletable_Cdcvfileentry_Cdcvfileoperstatus = "enabled"

    CISCODATACOLLECTIONMIB_Cdcvfiletable_Cdcvfileentry_Cdcvfileoperstatus_disabled CISCODATACOLLECTIONMIB_Cdcvfiletable_Cdcvfileentry_Cdcvfileoperstatus = "disabled"

    CISCODATACOLLECTIONMIB_Cdcvfiletable_Cdcvfileentry_Cdcvfileoperstatus_error CISCODATACOLLECTIONMIB_Cdcvfiletable_Cdcvfileentry_Cdcvfileoperstatus = "error"
)

// CISCODATACOLLECTIONMIB_Cdcvfilemgmttable
// A table to manage frozen VFiles.
type CISCODATACOLLECTIONMIB_Cdcvfilemgmttable struct {
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
    // CISCODATACOLLECTIONMIB_Cdcvfilemgmttable_Cdcvfilemgmtentry.
    Cdcvfilemgmtentry []CISCODATACOLLECTIONMIB_Cdcvfilemgmttable_Cdcvfilemgmtentry
}

func (cdcvfilemgmttable *CISCODATACOLLECTIONMIB_Cdcvfilemgmttable) GetEntityData() *types.CommonEntityData {
    cdcvfilemgmttable.EntityData.YFilter = cdcvfilemgmttable.YFilter
    cdcvfilemgmttable.EntityData.YangName = "cdcVFileMgmtTable"
    cdcvfilemgmttable.EntityData.BundleName = "cisco_ios_xe"
    cdcvfilemgmttable.EntityData.ParentYangName = "CISCO-DATA-COLLECTION-MIB"
    cdcvfilemgmttable.EntityData.SegmentPath = "cdcVFileMgmtTable"
    cdcvfilemgmttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdcvfilemgmttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdcvfilemgmttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdcvfilemgmttable.EntityData.Children = make(map[string]types.YChild)
    cdcvfilemgmttable.EntityData.Children["cdcVFileMgmtEntry"] = types.YChild{"Cdcvfilemgmtentry", nil}
    for i := range cdcvfilemgmttable.Cdcvfilemgmtentry {
        cdcvfilemgmttable.EntityData.Children[types.GetSegmentPath(&cdcvfilemgmttable.Cdcvfilemgmtentry[i])] = types.YChild{"Cdcvfilemgmtentry", &cdcvfilemgmttable.Cdcvfilemgmtentry[i]}
    }
    cdcvfilemgmttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cdcvfilemgmttable.EntityData)
}

// CISCODATACOLLECTIONMIB_Cdcvfilemgmttable_Cdcvfilemgmtentry
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
type CISCODATACOLLECTIONMIB_Cdcvfilemgmttable_Cdcvfilemgmtentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // cisco_data_collection_mib.CISCODATACOLLECTIONMIB_Cdcvfiletable_Cdcvfileentry_Cdcvfileindex
    Cdcvfileindex interface{}

    // This attribute is a key. This value is a running counter starting at 1,
    // generated by the agent so that the combination of  cdcVFileIndex and
    // cdcVFileMgmtIndex uniquely identifies a frozen VFile. The deleted file
    // indicies do not get reused.  This object's value needs to be unique only
    // across the set of frozen VFiles corresponding to a cdcVFileEntry
    // (identified by cdcVFileIndex). The type is interface{} with range:
    // 1..4294967295.
    Cdcvfilemgmtindex interface{}

    // The full name of the VFile. If the VFile is stored as a file in the agent's
    // filesystem, then this value also contains the absolute path of the file.
    // The type is string.
    Cdcvfilemgmtname interface{}

    // The timestamp when this VFile was created, in the date-time format. The
    // type is string.
    Cdcvfilemgmttimestamp interface{}

    // The time left before this VFile is deleted (see cdcVFileRetentionPeriod).
    // The type is interface{} with range: 60..86400. Units are seconds.
    Cdcvfilemgmttimetolive interface{}

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
    // Cdcvfilemgmtcommand.
    Cdcvfilemgmtcommand interface{}

    // The complete URL of the destination to which this VFile will be transferred
    // in the next attempt. The URL also includes the complete filename of the
    // remote file that will be created. When the default value of this object is 
    // retained this VFile will be transferred to the URL   specified in
    // cdcFileXferConfPriUrl or cdcFileXferConfSecUrl, as the case may be.  
    // However an application can specify a different URL, in which case the VFile
    // will be transferred to this new URL the next time transfer is initiated.  
    // This object's value may be modified at any time. The type is string with
    // length: 0..255.
    Cdcvfilemgmtxferurl interface{}

    // Indicates the status of the last completed transfer. The type is
    // CdcFileXferStatus.
    Cdcvfilemgmtlastxferstatus interface{}

    // Indicates the URL of the destination to which the last (completed) transfer
    // was initiated. The type is string with length: 0..255.
    Cdcvfilemgmtlastxferurl interface{}
}

func (cdcvfilemgmtentry *CISCODATACOLLECTIONMIB_Cdcvfilemgmttable_Cdcvfilemgmtentry) GetEntityData() *types.CommonEntityData {
    cdcvfilemgmtentry.EntityData.YFilter = cdcvfilemgmtentry.YFilter
    cdcvfilemgmtentry.EntityData.YangName = "cdcVFileMgmtEntry"
    cdcvfilemgmtentry.EntityData.BundleName = "cisco_ios_xe"
    cdcvfilemgmtentry.EntityData.ParentYangName = "cdcVFileMgmtTable"
    cdcvfilemgmtentry.EntityData.SegmentPath = "cdcVFileMgmtEntry" + "[cdcVFileIndex='" + fmt.Sprintf("%v", cdcvfilemgmtentry.Cdcvfileindex) + "']" + "[cdcVFileMgmtIndex='" + fmt.Sprintf("%v", cdcvfilemgmtentry.Cdcvfilemgmtindex) + "']"
    cdcvfilemgmtentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdcvfilemgmtentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdcvfilemgmtentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdcvfilemgmtentry.EntityData.Children = make(map[string]types.YChild)
    cdcvfilemgmtentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cdcvfilemgmtentry.EntityData.Leafs["cdcVFileIndex"] = types.YLeaf{"Cdcvfileindex", cdcvfilemgmtentry.Cdcvfileindex}
    cdcvfilemgmtentry.EntityData.Leafs["cdcVFileMgmtIndex"] = types.YLeaf{"Cdcvfilemgmtindex", cdcvfilemgmtentry.Cdcvfilemgmtindex}
    cdcvfilemgmtentry.EntityData.Leafs["cdcVFileMgmtName"] = types.YLeaf{"Cdcvfilemgmtname", cdcvfilemgmtentry.Cdcvfilemgmtname}
    cdcvfilemgmtentry.EntityData.Leafs["cdcVFileMgmtTimestamp"] = types.YLeaf{"Cdcvfilemgmttimestamp", cdcvfilemgmtentry.Cdcvfilemgmttimestamp}
    cdcvfilemgmtentry.EntityData.Leafs["cdcVFileMgmtTimeToLive"] = types.YLeaf{"Cdcvfilemgmttimetolive", cdcvfilemgmtentry.Cdcvfilemgmttimetolive}
    cdcvfilemgmtentry.EntityData.Leafs["cdcVFileMgmtCommand"] = types.YLeaf{"Cdcvfilemgmtcommand", cdcvfilemgmtentry.Cdcvfilemgmtcommand}
    cdcvfilemgmtentry.EntityData.Leafs["cdcVFileMgmtXferURL"] = types.YLeaf{"Cdcvfilemgmtxferurl", cdcvfilemgmtentry.Cdcvfilemgmtxferurl}
    cdcvfilemgmtentry.EntityData.Leafs["cdcVFileMgmtLastXferStatus"] = types.YLeaf{"Cdcvfilemgmtlastxferstatus", cdcvfilemgmtentry.Cdcvfilemgmtlastxferstatus}
    cdcvfilemgmtentry.EntityData.Leafs["cdcVFileMgmtLastXferURL"] = types.YLeaf{"Cdcvfilemgmtlastxferurl", cdcvfilemgmtentry.Cdcvfilemgmtlastxferurl}
    return &(cdcvfilemgmtentry.EntityData)
}

// CISCODATACOLLECTIONMIB_Cdcvfilemgmttable_Cdcvfilemgmtentry_Cdcvfilemgmtcommand represents                to abort an ongoing transfer.
type CISCODATACOLLECTIONMIB_Cdcvfilemgmttable_Cdcvfilemgmtentry_Cdcvfilemgmtcommand string

const (
    CISCODATACOLLECTIONMIB_Cdcvfilemgmttable_Cdcvfilemgmtentry_Cdcvfilemgmtcommand_idle CISCODATACOLLECTIONMIB_Cdcvfilemgmttable_Cdcvfilemgmtentry_Cdcvfilemgmtcommand = "idle"

    CISCODATACOLLECTIONMIB_Cdcvfilemgmttable_Cdcvfilemgmtentry_Cdcvfilemgmtcommand_delete CISCODATACOLLECTIONMIB_Cdcvfilemgmttable_Cdcvfilemgmtentry_Cdcvfilemgmtcommand = "delete"

    CISCODATACOLLECTIONMIB_Cdcvfilemgmttable_Cdcvfilemgmtentry_Cdcvfilemgmtcommand_transfer CISCODATACOLLECTIONMIB_Cdcvfilemgmttable_Cdcvfilemgmtentry_Cdcvfilemgmtcommand = "transfer"

    CISCODATACOLLECTIONMIB_Cdcvfilemgmttable_Cdcvfilemgmtentry_Cdcvfilemgmtcommand_abortTransfer CISCODATACOLLECTIONMIB_Cdcvfilemgmttable_Cdcvfilemgmtentry_Cdcvfilemgmtcommand = "abortTransfer"
)

// CISCODATACOLLECTIONMIB_Cdcdgtable
// A table for specifying data groups.
type CISCODATACOLLECTIONMIB_Cdcdgtable struct {
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
    // CISCODATACOLLECTIONMIB_Cdcdgtable_Cdcdgentry.
    Cdcdgentry []CISCODATACOLLECTIONMIB_Cdcdgtable_Cdcdgentry
}

func (cdcdgtable *CISCODATACOLLECTIONMIB_Cdcdgtable) GetEntityData() *types.CommonEntityData {
    cdcdgtable.EntityData.YFilter = cdcdgtable.YFilter
    cdcdgtable.EntityData.YangName = "cdcDGTable"
    cdcdgtable.EntityData.BundleName = "cisco_ios_xe"
    cdcdgtable.EntityData.ParentYangName = "CISCO-DATA-COLLECTION-MIB"
    cdcdgtable.EntityData.SegmentPath = "cdcDGTable"
    cdcdgtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdcdgtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdcdgtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdcdgtable.EntityData.Children = make(map[string]types.YChild)
    cdcdgtable.EntityData.Children["cdcDGEntry"] = types.YChild{"Cdcdgentry", nil}
    for i := range cdcdgtable.Cdcdgentry {
        cdcdgtable.EntityData.Children[types.GetSegmentPath(&cdcdgtable.Cdcdgentry[i])] = types.YChild{"Cdcdgentry", &cdcdgtable.Cdcdgentry[i]}
    }
    cdcdgtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cdcdgtable.EntityData)
}

// CISCODATACOLLECTIONMIB_Cdcdgtable_Cdcdgentry
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
type CISCODATACOLLECTIONMIB_Cdcdgtable_Cdcdgentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An arbitrary integer used to uniquely identify
    // this entry. When creating an entry, a management application should pick a
    // random number. The type is interface{} with range: 1..4294967295.
    Cdcdgindex interface{}

    // A descriptive string. This object's value may be modified at any time. The
    // type is string.
    Cdcdgcomment interface{}

    // Identifies the type of this data group. object         Data is a single MIB
    // object. The fully                instantiated OID is specified in          
    // cdcDGBaseObject.  table          Data is a logical table. The columns of   
    // this table correspond to the base objects                specified in
    // cdcDGBaseObjectTable, and the                rows correspond to the values
    // of the instances                specified in cdcDGInstanceTable.  This
    // object's value cannot be modified while the value of cdcDGRowStatus is
    // 'active'. The type is Cdcdgtype.
    Cdcdgtype interface{}

    // Corresponds to a value of cdcVFileIndex. It is used to associate this data
    // group with a cdcVFileEntry. The values of the base objects for  this data
    // group are stored into the current VFile of the associated cdcVFileEntry.  
    // This object's value may be modified at any time. The change takes effect
    // the next time data is fetched for this data group. The type is interface{}
    // with range: 1..4294967295.
    Cdcdgvfileindex interface{}

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
    Cdcdgtargettag interface{}

    // The management context from which to obtain data for this data group.  This
    // object's value may be modified at any time. The change takes effect the
    // next time data is fetched for this data group. The type is string.
    Cdcdgcontextname interface{}

    // The fully instantiated name of the MIB object whose value needs to be
    // fetched. This object's value is used only when cdcDGType is of type
    // 'object'.   This object's value may be modified at any time. The change
    // takes effect the next time data is fetched for this data group. The type is
    // string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Cdcdgobject interface{}

    // Corresponds to a value of cdcDGBaseObjectGrpIndex, thus identifying a set
    // of entries in cdcDGBaseObjectTable, having this value of
    // cdcDGBaseObjectGrpIndex. This object's value is used only when cdcDGType is
    // of type 'table'.   This set of entries in cdcDGBaseObjectTable in turn 
    // identifies the set of base objects, that makes up the columns  of this
    // 'table' type data group.     This object's value may be modified at any
    // time. The change takes effect the next time data is fetched for this data
    // group. The type is interface{} with range: 1..4294967295.
    Cdcdgobjectgrpindex interface{}

    // Corresponds to a value of cdcDGInstanceGrpIndex, thus identifying a set of
    // entries in cdcDGInstanceTable, having this value of cdcDGInstanceGrpIndex.
    // This object's value is used only when cdcDGType is of type 'table'.   The
    // set of entries in cdcDGInstanceTable, in turn identifies the set of
    // instances of the base objects, whose values need to fetched. If the value
    // is 0, then all instances of the base objects will be fetched.    This
    // object's value may be modified at any time. The change takes effect the
    // next time data is fetched for this data group. The type is interface{} with
    // range: 0..4294967295.
    Cdcdginstgrpindex interface{}

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
    Cdcdgpollperiod interface{}

    // The status of this conceptual row.  This object cannot be set to 'active'
    // until values have been assigned to  cdcDGVFileIndex & cdcDGColGrpIndex. The
    // type is RowStatus.
    Cdcdgrowstatus interface{}
}

func (cdcdgentry *CISCODATACOLLECTIONMIB_Cdcdgtable_Cdcdgentry) GetEntityData() *types.CommonEntityData {
    cdcdgentry.EntityData.YFilter = cdcdgentry.YFilter
    cdcdgentry.EntityData.YangName = "cdcDGEntry"
    cdcdgentry.EntityData.BundleName = "cisco_ios_xe"
    cdcdgentry.EntityData.ParentYangName = "cdcDGTable"
    cdcdgentry.EntityData.SegmentPath = "cdcDGEntry" + "[cdcDGIndex='" + fmt.Sprintf("%v", cdcdgentry.Cdcdgindex) + "']"
    cdcdgentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdcdgentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdcdgentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdcdgentry.EntityData.Children = make(map[string]types.YChild)
    cdcdgentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cdcdgentry.EntityData.Leafs["cdcDGIndex"] = types.YLeaf{"Cdcdgindex", cdcdgentry.Cdcdgindex}
    cdcdgentry.EntityData.Leafs["cdcDGComment"] = types.YLeaf{"Cdcdgcomment", cdcdgentry.Cdcdgcomment}
    cdcdgentry.EntityData.Leafs["cdcDGType"] = types.YLeaf{"Cdcdgtype", cdcdgentry.Cdcdgtype}
    cdcdgentry.EntityData.Leafs["cdcDGVFileIndex"] = types.YLeaf{"Cdcdgvfileindex", cdcdgentry.Cdcdgvfileindex}
    cdcdgentry.EntityData.Leafs["cdcDGTargetTag"] = types.YLeaf{"Cdcdgtargettag", cdcdgentry.Cdcdgtargettag}
    cdcdgentry.EntityData.Leafs["cdcDGContextName"] = types.YLeaf{"Cdcdgcontextname", cdcdgentry.Cdcdgcontextname}
    cdcdgentry.EntityData.Leafs["cdcDGObject"] = types.YLeaf{"Cdcdgobject", cdcdgentry.Cdcdgobject}
    cdcdgentry.EntityData.Leafs["cdcDGObjectGrpIndex"] = types.YLeaf{"Cdcdgobjectgrpindex", cdcdgentry.Cdcdgobjectgrpindex}
    cdcdgentry.EntityData.Leafs["cdcDGInstGrpIndex"] = types.YLeaf{"Cdcdginstgrpindex", cdcdgentry.Cdcdginstgrpindex}
    cdcdgentry.EntityData.Leafs["cdcDGPollPeriod"] = types.YLeaf{"Cdcdgpollperiod", cdcdgentry.Cdcdgpollperiod}
    cdcdgentry.EntityData.Leafs["cdcDGRowStatus"] = types.YLeaf{"Cdcdgrowstatus", cdcdgentry.Cdcdgrowstatus}
    return &(cdcdgentry.EntityData)
}

// CISCODATACOLLECTIONMIB_Cdcdgtable_Cdcdgentry_Cdcdgtype represents cdcDGRowStatus is 'active'.
type CISCODATACOLLECTIONMIB_Cdcdgtable_Cdcdgentry_Cdcdgtype string

const (
    CISCODATACOLLECTIONMIB_Cdcdgtable_Cdcdgentry_Cdcdgtype_object CISCODATACOLLECTIONMIB_Cdcdgtable_Cdcdgentry_Cdcdgtype = "object"

    CISCODATACOLLECTIONMIB_Cdcdgtable_Cdcdgentry_Cdcdgtype_table CISCODATACOLLECTIONMIB_Cdcdgtable_Cdcdgentry_Cdcdgtype = "table"
)

// CISCODATACOLLECTIONMIB_Cdcdgbaseobjecttable
// A table specifying the base objects of a 'table' type
// data group.
type CISCODATACOLLECTIONMIB_Cdcdgbaseobjecttable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An individual entry in this table. Each entry is a  {subtree, list} tuple.
    // Each tuple identifies a set of  base objects for the associated data group.
    // The type is slice of
    // CISCODATACOLLECTIONMIB_Cdcdgbaseobjecttable_Cdcdgbaseobjectentry.
    Cdcdgbaseobjectentry []CISCODATACOLLECTIONMIB_Cdcdgbaseobjecttable_Cdcdgbaseobjectentry
}

func (cdcdgbaseobjecttable *CISCODATACOLLECTIONMIB_Cdcdgbaseobjecttable) GetEntityData() *types.CommonEntityData {
    cdcdgbaseobjecttable.EntityData.YFilter = cdcdgbaseobjecttable.YFilter
    cdcdgbaseobjecttable.EntityData.YangName = "cdcDGBaseObjectTable"
    cdcdgbaseobjecttable.EntityData.BundleName = "cisco_ios_xe"
    cdcdgbaseobjecttable.EntityData.ParentYangName = "CISCO-DATA-COLLECTION-MIB"
    cdcdgbaseobjecttable.EntityData.SegmentPath = "cdcDGBaseObjectTable"
    cdcdgbaseobjecttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdcdgbaseobjecttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdcdgbaseobjecttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdcdgbaseobjecttable.EntityData.Children = make(map[string]types.YChild)
    cdcdgbaseobjecttable.EntityData.Children["cdcDGBaseObjectEntry"] = types.YChild{"Cdcdgbaseobjectentry", nil}
    for i := range cdcdgbaseobjecttable.Cdcdgbaseobjectentry {
        cdcdgbaseobjecttable.EntityData.Children[types.GetSegmentPath(&cdcdgbaseobjecttable.Cdcdgbaseobjectentry[i])] = types.YChild{"Cdcdgbaseobjectentry", &cdcdgbaseobjecttable.Cdcdgbaseobjectentry[i]}
    }
    cdcdgbaseobjecttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cdcdgbaseobjecttable.EntityData)
}

// CISCODATACOLLECTIONMIB_Cdcdgbaseobjecttable_Cdcdgbaseobjectentry
// An individual entry in this table. Each entry is a 
// {subtree, list} tuple. Each tuple identifies a set of 
// base objects for the associated data group.
type CISCODATACOLLECTIONMIB_Cdcdgbaseobjecttable_Cdcdgbaseobjectentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This object's value when combined with the value
    // of cdcDGBaseObjectIndex uniquely identifies an entry in this table. An
    // application must use the same value (can  be randomly picked) for this
    // object while creating a group of entries that collectively identifies the
    // set of base objects for a data group. The type is interface{} with range:
    // 1..4294967295.
    Cdcdgbaseobjectgrpindex interface{}

    // This attribute is a key. This object's value when combined with the value
    // of cdcDGBaseObjectGrpIndex uniquely identifies an entry in this table.  A
    // managment application can assign incremental values starting from one, when
    // creating each entry in a group of entries (as identified by the value of
    // cdcDGBaseObjectGrpIndex). The type is interface{} with range:
    // 1..4294967295.
    Cdcdgbaseobjectindex interface{}

    // The subtree component of a {subtree, list} tuple.  This object's value may
    // be modified at any time. The change takes effect the next time data is
    // fetched for this data group. The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Cdcdgbaseobjectsubtree interface{}

    // The list component of a {subtree, list} tuple.  This object's value may be
    // modified at any time. The change takes effect the next time data is fetched
    // for this data group. The type is string with length: 0..16.
    Cdcdgbaseobjectlist interface{}

    // The status of this conceptual row.  This object cannot be set to 'active'
    // until values have been assigned to cdcDGBaseObjectSubtree &
    // cdcDGBaseObjectList. The type is RowStatus.
    Cdcdgbaseobjectrowstatus interface{}
}

func (cdcdgbaseobjectentry *CISCODATACOLLECTIONMIB_Cdcdgbaseobjecttable_Cdcdgbaseobjectentry) GetEntityData() *types.CommonEntityData {
    cdcdgbaseobjectentry.EntityData.YFilter = cdcdgbaseobjectentry.YFilter
    cdcdgbaseobjectentry.EntityData.YangName = "cdcDGBaseObjectEntry"
    cdcdgbaseobjectentry.EntityData.BundleName = "cisco_ios_xe"
    cdcdgbaseobjectentry.EntityData.ParentYangName = "cdcDGBaseObjectTable"
    cdcdgbaseobjectentry.EntityData.SegmentPath = "cdcDGBaseObjectEntry" + "[cdcDGBaseObjectGrpIndex='" + fmt.Sprintf("%v", cdcdgbaseobjectentry.Cdcdgbaseobjectgrpindex) + "']" + "[cdcDGBaseObjectIndex='" + fmt.Sprintf("%v", cdcdgbaseobjectentry.Cdcdgbaseobjectindex) + "']"
    cdcdgbaseobjectentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdcdgbaseobjectentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdcdgbaseobjectentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdcdgbaseobjectentry.EntityData.Children = make(map[string]types.YChild)
    cdcdgbaseobjectentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cdcdgbaseobjectentry.EntityData.Leafs["cdcDGBaseObjectGrpIndex"] = types.YLeaf{"Cdcdgbaseobjectgrpindex", cdcdgbaseobjectentry.Cdcdgbaseobjectgrpindex}
    cdcdgbaseobjectentry.EntityData.Leafs["cdcDGBaseObjectIndex"] = types.YLeaf{"Cdcdgbaseobjectindex", cdcdgbaseobjectentry.Cdcdgbaseobjectindex}
    cdcdgbaseobjectentry.EntityData.Leafs["cdcDGBaseObjectSubtree"] = types.YLeaf{"Cdcdgbaseobjectsubtree", cdcdgbaseobjectentry.Cdcdgbaseobjectsubtree}
    cdcdgbaseobjectentry.EntityData.Leafs["cdcDGBaseObjectList"] = types.YLeaf{"Cdcdgbaseobjectlist", cdcdgbaseobjectentry.Cdcdgbaseobjectlist}
    cdcdgbaseobjectentry.EntityData.Leafs["cdcDGBaseObjectRowStatus"] = types.YLeaf{"Cdcdgbaseobjectrowstatus", cdcdgbaseobjectentry.Cdcdgbaseobjectrowstatus}
    return &(cdcdgbaseobjectentry.EntityData)
}

// CISCODATACOLLECTIONMIB_Cdcdginstancetable
// Identifies the instances of the base objects that need to
// be fetched for a 'table' type data group.
// 
// The agent is not responsible for verifying that the instances
// specified for a data group do not overlap.
type CISCODATACOLLECTIONMIB_Cdcdginstancetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table. Each entry identifies one or more instances of the
    // base objects that need to be fetched. An instance is represented by an OID
    // fragment. The type is slice of
    // CISCODATACOLLECTIONMIB_Cdcdginstancetable_Cdcdginstanceentry.
    Cdcdginstanceentry []CISCODATACOLLECTIONMIB_Cdcdginstancetable_Cdcdginstanceentry
}

func (cdcdginstancetable *CISCODATACOLLECTIONMIB_Cdcdginstancetable) GetEntityData() *types.CommonEntityData {
    cdcdginstancetable.EntityData.YFilter = cdcdginstancetable.YFilter
    cdcdginstancetable.EntityData.YangName = "cdcDGInstanceTable"
    cdcdginstancetable.EntityData.BundleName = "cisco_ios_xe"
    cdcdginstancetable.EntityData.ParentYangName = "CISCO-DATA-COLLECTION-MIB"
    cdcdginstancetable.EntityData.SegmentPath = "cdcDGInstanceTable"
    cdcdginstancetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdcdginstancetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdcdginstancetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdcdginstancetable.EntityData.Children = make(map[string]types.YChild)
    cdcdginstancetable.EntityData.Children["cdcDGInstanceEntry"] = types.YChild{"Cdcdginstanceentry", nil}
    for i := range cdcdginstancetable.Cdcdginstanceentry {
        cdcdginstancetable.EntityData.Children[types.GetSegmentPath(&cdcdginstancetable.Cdcdginstanceentry[i])] = types.YChild{"Cdcdginstanceentry", &cdcdginstancetable.Cdcdginstanceentry[i]}
    }
    cdcdginstancetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cdcdginstancetable.EntityData)
}

// CISCODATACOLLECTIONMIB_Cdcdginstancetable_Cdcdginstanceentry
// An entry in this table. Each entry identifies one or more
// instances of the base objects that need to be fetched.
// An instance is represented by an OID fragment.
type CISCODATACOLLECTIONMIB_Cdcdginstancetable_Cdcdginstanceentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This object's value when combined with the value
    // of cdcDGInstanceIndex uniquely identifies an entry in this table. An
    // application must use the same value (can  be randomly picked) for this
    // object while creating a group of entries that collectively identifies the
    // set of instances for a data group. The type is interface{} with range:
    // 1..4294967295.
    Cdcdginstancegrpindex interface{}

    // This attribute is a key. This object's value when combined with the value
    // of cdcDGInstanceGrpIndex uniquely identifies an entry in this table.  A
    // managment application can assign incremental values starting from one, when
    // creating each entry within a group of entries (as identified by the value
    // of cdcDGInstanceGrpIndex). The type is interface{} with range:
    // 1..4294967295.
    Cdcdginstanceindex interface{}

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
    // value of cdcDGInstanceStatus is 'active'. The type is Cdcdginstancetype.
    Cdcdginstancetype interface{}

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
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Cdcdginstanceoid interface{}

    // Contains the OID fragment that, when appended to each base object gives the
    // end of the range of object instances that needs to be fetched.  This value
    // is used only when the value of cdcDGInstanceType is of type 'range'.   This
    // object's value may be modified at any time. The change takes effect the
    // next time data is fetched for this data group. The type is string with
    // pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Cdcdginstanceoidend interface{}

    // Specifies the number of lexicographically consecutive object instances to
    // fetch.  This value is used only when the value of cdcDGInstanceType is of
    // type 'repititions'.    This object's value may be modified at any time. The
    // change takes effect the next time data is fetched for this data group. The
    // type is interface{} with range: 0..4294967295.
    Cdcdginstancenumrepititions interface{}

    // Contains a pointer to a row in another MIB table that contains MIB specific
    // criteria for selecting instances.  This value is used only when the value
    // of cdcDGInstanceType is of type 'other'.   This object's value may be
    // modified at any time. The change takes effect the next time data is fetched
    // for this data group. The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Cdcdginstanceotherptr interface{}

    // The status of this conceptual row. The type is RowStatus.
    Cdcdginstancerowstatus interface{}
}

func (cdcdginstanceentry *CISCODATACOLLECTIONMIB_Cdcdginstancetable_Cdcdginstanceentry) GetEntityData() *types.CommonEntityData {
    cdcdginstanceentry.EntityData.YFilter = cdcdginstanceentry.YFilter
    cdcdginstanceentry.EntityData.YangName = "cdcDGInstanceEntry"
    cdcdginstanceentry.EntityData.BundleName = "cisco_ios_xe"
    cdcdginstanceentry.EntityData.ParentYangName = "cdcDGInstanceTable"
    cdcdginstanceentry.EntityData.SegmentPath = "cdcDGInstanceEntry" + "[cdcDGInstanceGrpIndex='" + fmt.Sprintf("%v", cdcdginstanceentry.Cdcdginstancegrpindex) + "']" + "[cdcDGInstanceIndex='" + fmt.Sprintf("%v", cdcdginstanceentry.Cdcdginstanceindex) + "']"
    cdcdginstanceentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdcdginstanceentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdcdginstanceentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdcdginstanceentry.EntityData.Children = make(map[string]types.YChild)
    cdcdginstanceentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cdcdginstanceentry.EntityData.Leafs["cdcDGInstanceGrpIndex"] = types.YLeaf{"Cdcdginstancegrpindex", cdcdginstanceentry.Cdcdginstancegrpindex}
    cdcdginstanceentry.EntityData.Leafs["cdcDGInstanceIndex"] = types.YLeaf{"Cdcdginstanceindex", cdcdginstanceentry.Cdcdginstanceindex}
    cdcdginstanceentry.EntityData.Leafs["cdcDGInstanceType"] = types.YLeaf{"Cdcdginstancetype", cdcdginstanceentry.Cdcdginstancetype}
    cdcdginstanceentry.EntityData.Leafs["cdcDGInstanceOid"] = types.YLeaf{"Cdcdginstanceoid", cdcdginstanceentry.Cdcdginstanceoid}
    cdcdginstanceentry.EntityData.Leafs["cdcDGInstanceOidEnd"] = types.YLeaf{"Cdcdginstanceoidend", cdcdginstanceentry.Cdcdginstanceoidend}
    cdcdginstanceentry.EntityData.Leafs["cdcDGInstanceNumRepititions"] = types.YLeaf{"Cdcdginstancenumrepititions", cdcdginstanceentry.Cdcdginstancenumrepititions}
    cdcdginstanceentry.EntityData.Leafs["cdcDGInstanceOtherPtr"] = types.YLeaf{"Cdcdginstanceotherptr", cdcdginstanceentry.Cdcdginstanceotherptr}
    cdcdginstanceentry.EntityData.Leafs["cdcDGInstanceRowStatus"] = types.YLeaf{"Cdcdginstancerowstatus", cdcdginstanceentry.Cdcdginstancerowstatus}
    return &(cdcdginstanceentry.EntityData)
}

// CISCODATACOLLECTIONMIB_Cdcdginstancetable_Cdcdginstanceentry_Cdcdginstancetype represents cdcDGInstanceStatus is 'active'.
type CISCODATACOLLECTIONMIB_Cdcdginstancetable_Cdcdginstanceentry_Cdcdginstancetype string

const (
    CISCODATACOLLECTIONMIB_Cdcdginstancetable_Cdcdginstanceentry_Cdcdginstancetype_individual CISCODATACOLLECTIONMIB_Cdcdginstancetable_Cdcdginstanceentry_Cdcdginstancetype = "individual"

    CISCODATACOLLECTIONMIB_Cdcdginstancetable_Cdcdginstanceentry_Cdcdginstancetype_range_ CISCODATACOLLECTIONMIB_Cdcdginstancetable_Cdcdginstanceentry_Cdcdginstancetype = "range"

    CISCODATACOLLECTIONMIB_Cdcdginstancetable_Cdcdginstanceentry_Cdcdginstancetype_repititions CISCODATACOLLECTIONMIB_Cdcdginstancetable_Cdcdginstanceentry_Cdcdginstancetype = "repititions"

    CISCODATACOLLECTIONMIB_Cdcdginstancetable_Cdcdginstanceentry_Cdcdginstancetype_subTree CISCODATACOLLECTIONMIB_Cdcdginstancetable_Cdcdginstanceentry_Cdcdginstancetype = "subTree"

    CISCODATACOLLECTIONMIB_Cdcdginstancetable_Cdcdginstanceentry_Cdcdginstancetype_other CISCODATACOLLECTIONMIB_Cdcdginstancetable_Cdcdginstanceentry_Cdcdginstancetype = "other"
)

// CISCODATACOLLECTIONMIB_Cdcfilexferconftable
// A table for configuring file transfer operations.
type CISCODATACOLLECTIONMIB_Cdcfilexferconftable struct {
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
    // CISCODATACOLLECTIONMIB_Cdcfilexferconftable_Cdcfilexferconfentry.
    Cdcfilexferconfentry []CISCODATACOLLECTIONMIB_Cdcfilexferconftable_Cdcfilexferconfentry
}

func (cdcfilexferconftable *CISCODATACOLLECTIONMIB_Cdcfilexferconftable) GetEntityData() *types.CommonEntityData {
    cdcfilexferconftable.EntityData.YFilter = cdcfilexferconftable.YFilter
    cdcfilexferconftable.EntityData.YangName = "cdcFileXferConfTable"
    cdcfilexferconftable.EntityData.BundleName = "cisco_ios_xe"
    cdcfilexferconftable.EntityData.ParentYangName = "CISCO-DATA-COLLECTION-MIB"
    cdcfilexferconftable.EntityData.SegmentPath = "cdcFileXferConfTable"
    cdcfilexferconftable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdcfilexferconftable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdcfilexferconftable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdcfilexferconftable.EntityData.Children = make(map[string]types.YChild)
    cdcfilexferconftable.EntityData.Children["cdcFileXferConfEntry"] = types.YChild{"Cdcfilexferconfentry", nil}
    for i := range cdcfilexferconftable.Cdcfilexferconfentry {
        cdcfilexferconftable.EntityData.Children[types.GetSegmentPath(&cdcfilexferconftable.Cdcfilexferconfentry[i])] = types.YChild{"Cdcfilexferconfentry", &cdcfilexferconftable.Cdcfilexferconfentry[i]}
    }
    cdcfilexferconftable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cdcfilexferconftable.EntityData)
}

// CISCODATACOLLECTIONMIB_Cdcfilexferconftable_Cdcfilexferconfentry
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
type CISCODATACOLLECTIONMIB_Cdcfilexferconftable_Cdcfilexferconfentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // cisco_data_collection_mib.CISCODATACOLLECTIONMIB_Cdcvfiletable_Cdcvfileentry_Cdcvfileindex
    Cdcvfileindex interface{}

    // The URL which specifies the primary destination to which the file has to be
    // transferred. The URL should contain the base-name of the remote file, the
    // suffix will be carried over from the name of the VFile being tranferred,
    // and will be automatically appended by the agent.  This object's value may
    // be modified at any time. The type is string with length: 0..255.
    Cdcfilexferconfpriurl interface{}

    // The URL which specifies the secondary destination to which the file has to
    // be transferred if the transfer to the  primary destination fails. Failure
    // occurs when the file  cannot be transferred in it's entirety to the
    // specified  destination for some reason. Some common reasons for such 
    // failures are listed out in CdcFileXferStatus.    The specified URL should
    // contain the base-name of the remote file, the suffix will be carried over
    // from the name of the VFile being tranferred, and will be automatically
    // appended by the agent.   This object's value may be modified at any time.
    // The type is string with length: 0..255.
    Cdcfilexferconfsecurl interface{}

    // Specifies the time interval after which transfer has to be retried.
    // Transfer needs to be retried only if in a previous  attempt the file could
    // not be successfully transferred to  either the primary destination or the
    // secondary destination.  This object's value may be modified at any time.
    // The type is interface{} with range: 60..86400. Units are seconds.
    Cdcfilexferconfretryperiod interface{}

    // Maximum number of times, transfer has to be retried. If the retry count
    // exceeds this value, then no further attempts will be made.   This object's
    // value may be modified at any time. The type is interface{} with range:
    // 0..256. Units are seconds.
    Cdcfilexferconfretrycount interface{}

    // When set to 'true', cdcFileXferComplete notification will be sent out in
    // the event of a successful file transfer. The type is bool.
    Cdcfilexferconfsuccessenable interface{}

    // When set to 'true', cdcFileXferComplete notification will be sent out in
    // the event of a file transfer failure. The type is bool.
    Cdcfilexferconffailureenable interface{}
}

func (cdcfilexferconfentry *CISCODATACOLLECTIONMIB_Cdcfilexferconftable_Cdcfilexferconfentry) GetEntityData() *types.CommonEntityData {
    cdcfilexferconfentry.EntityData.YFilter = cdcfilexferconfentry.YFilter
    cdcfilexferconfentry.EntityData.YangName = "cdcFileXferConfEntry"
    cdcfilexferconfentry.EntityData.BundleName = "cisco_ios_xe"
    cdcfilexferconfentry.EntityData.ParentYangName = "cdcFileXferConfTable"
    cdcfilexferconfentry.EntityData.SegmentPath = "cdcFileXferConfEntry" + "[cdcVFileIndex='" + fmt.Sprintf("%v", cdcfilexferconfentry.Cdcvfileindex) + "']"
    cdcfilexferconfentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdcfilexferconfentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdcfilexferconfentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdcfilexferconfentry.EntityData.Children = make(map[string]types.YChild)
    cdcfilexferconfentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cdcfilexferconfentry.EntityData.Leafs["cdcVFileIndex"] = types.YLeaf{"Cdcvfileindex", cdcfilexferconfentry.Cdcvfileindex}
    cdcfilexferconfentry.EntityData.Leafs["cdcFileXferConfPriUrl"] = types.YLeaf{"Cdcfilexferconfpriurl", cdcfilexferconfentry.Cdcfilexferconfpriurl}
    cdcfilexferconfentry.EntityData.Leafs["cdcFileXferConfSecUrl"] = types.YLeaf{"Cdcfilexferconfsecurl", cdcfilexferconfentry.Cdcfilexferconfsecurl}
    cdcfilexferconfentry.EntityData.Leafs["cdcFileXferConfRetryPeriod"] = types.YLeaf{"Cdcfilexferconfretryperiod", cdcfilexferconfentry.Cdcfilexferconfretryperiod}
    cdcfilexferconfentry.EntityData.Leafs["cdcFileXferConfRetryCount"] = types.YLeaf{"Cdcfilexferconfretrycount", cdcfilexferconfentry.Cdcfilexferconfretrycount}
    cdcfilexferconfentry.EntityData.Leafs["cdcFileXferConfSuccessEnable"] = types.YLeaf{"Cdcfilexferconfsuccessenable", cdcfilexferconfentry.Cdcfilexferconfsuccessenable}
    cdcfilexferconfentry.EntityData.Leafs["cdcFileXferConfFailureEnable"] = types.YLeaf{"Cdcfilexferconffailureenable", cdcfilexferconfentry.Cdcfilexferconffailureenable}
    return &(cdcfilexferconfentry.EntityData)
}

