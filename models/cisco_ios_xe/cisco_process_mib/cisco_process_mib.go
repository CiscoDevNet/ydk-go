// The MIB module to describe active system processes.
// Virtual Machine refers to those OS which can run the 
// code or process of a different executional model OS.
// Virtual Process assume the executional model 
// of a OS which is different from Native OS. Virtual
// Processes are also referred as Tasks.
// Thread is a sequence of instructions to be executed
// within a program. Thread which adhere to POSIX standard
// is referred as a POSIX thread.
package cisco_process_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_process_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-PROCESS-MIB CISCO-PROCESS-MIB}", reflect.TypeOf(CISCOPROCESSMIB{}))
    ydk.RegisterEntity("CISCO-PROCESS-MIB:CISCO-PROCESS-MIB", reflect.TypeOf(CISCOPROCESSMIB{}))
}

// CISCOPROCESSMIB
type CISCOPROCESSMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    CpmCPUHistory CISCOPROCESSMIB_CpmCPUHistory

    // A table of overall CPU statistics.
    CpmCPUTotalTable CISCOPROCESSMIB_CpmCPUTotalTable

    // A table of per-Core statistics.
    CpmCoreTable CISCOPROCESSMIB_CpmCoreTable

    // A table of generic information on all active processes on this device.
    CpmProcessTable CISCOPROCESSMIB_CpmProcessTable

    // This table contains information that may or may not be available on all
    // cisco devices. It contains additional objects for the more general
    // cpmProcessTable. This object deprecates  cpmProcessExtTable.
    CpmProcessExtRevTable CISCOPROCESSMIB_CpmProcessExtRevTable

    // This table contains the information about the thresholding values for CPU ,
    // configured by the user.
    CpmCPUThresholdTable CISCOPROCESSMIB_CpmCPUThresholdTable

    // A list of CPU utilization history entries.
    CpmCPUHistoryTable CISCOPROCESSMIB_CpmCPUHistoryTable

    // A list of process history entries. This table contains CPU utilization of
    // processes which crossed the  cpmCPUHistoryThreshold.
    CpmCPUProcessHistoryTable CISCOPROCESSMIB_CpmCPUProcessHistoryTable

    // This table contains generic information about POSIX threads in the device.
    CpmThreadTable CISCOPROCESSMIB_CpmThreadTable

    // This table contains information about virtual processes in a virtual
    // machine.
    CpmVirtualProcessTable CISCOPROCESSMIB_CpmVirtualProcessTable
}

func (cISCOPROCESSMIB *CISCOPROCESSMIB) GetEntityData() *types.CommonEntityData {
    cISCOPROCESSMIB.EntityData.YFilter = cISCOPROCESSMIB.YFilter
    cISCOPROCESSMIB.EntityData.YangName = "CISCO-PROCESS-MIB"
    cISCOPROCESSMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOPROCESSMIB.EntityData.ParentYangName = "CISCO-PROCESS-MIB"
    cISCOPROCESSMIB.EntityData.SegmentPath = "CISCO-PROCESS-MIB:CISCO-PROCESS-MIB"
    cISCOPROCESSMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOPROCESSMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOPROCESSMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOPROCESSMIB.EntityData.Children = types.NewOrderedMap()
    cISCOPROCESSMIB.EntityData.Children.Append("cpmCPUHistory", types.YChild{"CpmCPUHistory", &cISCOPROCESSMIB.CpmCPUHistory})
    cISCOPROCESSMIB.EntityData.Children.Append("cpmCPUTotalTable", types.YChild{"CpmCPUTotalTable", &cISCOPROCESSMIB.CpmCPUTotalTable})
    cISCOPROCESSMIB.EntityData.Children.Append("cpmCoreTable", types.YChild{"CpmCoreTable", &cISCOPROCESSMIB.CpmCoreTable})
    cISCOPROCESSMIB.EntityData.Children.Append("cpmProcessTable", types.YChild{"CpmProcessTable", &cISCOPROCESSMIB.CpmProcessTable})
    cISCOPROCESSMIB.EntityData.Children.Append("cpmProcessExtRevTable", types.YChild{"CpmProcessExtRevTable", &cISCOPROCESSMIB.CpmProcessExtRevTable})
    cISCOPROCESSMIB.EntityData.Children.Append("cpmCPUThresholdTable", types.YChild{"CpmCPUThresholdTable", &cISCOPROCESSMIB.CpmCPUThresholdTable})
    cISCOPROCESSMIB.EntityData.Children.Append("cpmCPUHistoryTable", types.YChild{"CpmCPUHistoryTable", &cISCOPROCESSMIB.CpmCPUHistoryTable})
    cISCOPROCESSMIB.EntityData.Children.Append("cpmCPUProcessHistoryTable", types.YChild{"CpmCPUProcessHistoryTable", &cISCOPROCESSMIB.CpmCPUProcessHistoryTable})
    cISCOPROCESSMIB.EntityData.Children.Append("cpmThreadTable", types.YChild{"CpmThreadTable", &cISCOPROCESSMIB.CpmThreadTable})
    cISCOPROCESSMIB.EntityData.Children.Append("cpmVirtualProcessTable", types.YChild{"CpmVirtualProcessTable", &cISCOPROCESSMIB.CpmVirtualProcessTable})
    cISCOPROCESSMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOPROCESSMIB.EntityData.YListKeys = []string {}

    return &(cISCOPROCESSMIB.EntityData)
}

// CISCOPROCESSMIB_CpmCPUHistory
type CISCOPROCESSMIB_CpmCPUHistory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The user  configured value of this object gives the minimum percent CPU
    // utilization of a process in the last cpmCPUMonInterval duration required to
    // be a  member of history table. When this object is changed the new value
    // will have effect in the next interval. The type is interface{} with range:
    // 1..100.
    CpmCPUHistoryThreshold interface{}

    // A value configured by the user which specifies the number of reports in the
    // history table.  A report contains set of processes which crossed the
    // cpmCPUHistoryThreshold  in the last cpmCPUMonInterval along with  the time
    // at which this report is created, total and interrupt CPU utilizations. 
    // When this object is changed the new value will have effect in the next
    // interval. The type is interface{} with range: 1..4294967295.
    CpmCPUHistorySize interface{}
}

func (cpmCPUHistory *CISCOPROCESSMIB_CpmCPUHistory) GetEntityData() *types.CommonEntityData {
    cpmCPUHistory.EntityData.YFilter = cpmCPUHistory.YFilter
    cpmCPUHistory.EntityData.YangName = "cpmCPUHistory"
    cpmCPUHistory.EntityData.BundleName = "cisco_ios_xe"
    cpmCPUHistory.EntityData.ParentYangName = "CISCO-PROCESS-MIB"
    cpmCPUHistory.EntityData.SegmentPath = "cpmCPUHistory"
    cpmCPUHistory.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpmCPUHistory.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpmCPUHistory.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpmCPUHistory.EntityData.Children = types.NewOrderedMap()
    cpmCPUHistory.EntityData.Leafs = types.NewOrderedMap()
    cpmCPUHistory.EntityData.Leafs.Append("cpmCPUHistoryThreshold", types.YLeaf{"CpmCPUHistoryThreshold", cpmCPUHistory.CpmCPUHistoryThreshold})
    cpmCPUHistory.EntityData.Leafs.Append("cpmCPUHistorySize", types.YLeaf{"CpmCPUHistorySize", cpmCPUHistory.CpmCPUHistorySize})

    cpmCPUHistory.EntityData.YListKeys = []string {}

    return &(cpmCPUHistory.EntityData)
}

// CISCOPROCESSMIB_CpmCPUTotalTable
// A table of overall CPU statistics.
type CISCOPROCESSMIB_CpmCPUTotalTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Overall information about the CPU load. Entries in this table come and go
    // as CPUs are added and removed from the system. The type is slice of
    // CISCOPROCESSMIB_CpmCPUTotalTable_CpmCPUTotalEntry.
    CpmCPUTotalEntry []*CISCOPROCESSMIB_CpmCPUTotalTable_CpmCPUTotalEntry
}

func (cpmCPUTotalTable *CISCOPROCESSMIB_CpmCPUTotalTable) GetEntityData() *types.CommonEntityData {
    cpmCPUTotalTable.EntityData.YFilter = cpmCPUTotalTable.YFilter
    cpmCPUTotalTable.EntityData.YangName = "cpmCPUTotalTable"
    cpmCPUTotalTable.EntityData.BundleName = "cisco_ios_xe"
    cpmCPUTotalTable.EntityData.ParentYangName = "CISCO-PROCESS-MIB"
    cpmCPUTotalTable.EntityData.SegmentPath = "cpmCPUTotalTable"
    cpmCPUTotalTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpmCPUTotalTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpmCPUTotalTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpmCPUTotalTable.EntityData.Children = types.NewOrderedMap()
    cpmCPUTotalTable.EntityData.Children.Append("cpmCPUTotalEntry", types.YChild{"CpmCPUTotalEntry", nil})
    for i := range cpmCPUTotalTable.CpmCPUTotalEntry {
        cpmCPUTotalTable.EntityData.Children.Append(types.GetSegmentPath(cpmCPUTotalTable.CpmCPUTotalEntry[i]), types.YChild{"CpmCPUTotalEntry", cpmCPUTotalTable.CpmCPUTotalEntry[i]})
    }
    cpmCPUTotalTable.EntityData.Leafs = types.NewOrderedMap()

    cpmCPUTotalTable.EntityData.YListKeys = []string {}

    return &(cpmCPUTotalTable.EntityData)
}

// CISCOPROCESSMIB_CpmCPUTotalTable_CpmCPUTotalEntry
// Overall information about the CPU load. Entries in this
// table come and go as CPUs are added and removed from the
// system.
type CISCOPROCESSMIB_CpmCPUTotalTable_CpmCPUTotalEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An index that uniquely represents a CPU (or group
    // of CPUs) whose CPU load information is reported by a row in this table.
    // This index is assigned arbitrarily by the engine and is not saved over
    // reboots. The type is interface{} with range: 1..4294967295.
    CpmCPUTotalIndex interface{}

    // The entPhysicalIndex of the physical entity for which the CPU statistics in
    // this entry are maintained. The physical entity can be a CPU chip, a group
    // of CPUs, a CPU card etc. The exact type of this entity is described by its
    // entPhysicalVendorType value. If the CPU statistics in this entry correspond
    // to more than one physical entity (or to no physical entity), or if the
    // entPhysicalTable is not supported on the SNMP agent, the value of this
    // object must be zero. The type is interface{} with range: 0..2147483647.
    CpmCPUTotalPhysicalIndex interface{}

    // The overall CPU busy percentage in the last 5 second period. This object
    // obsoletes the busyPer object from  the OLD-CISCO-SYSTEM-MIB. This object is
    // deprecated by cpmCPUTotal5secRev which has the changed range of value
    // (0..100). The type is interface{} with range: 0..100.
    CpmCPUTotal5sec interface{}

    // The overall CPU busy percentage in the last 1 minute period. This object
    // obsoletes the avgBusy1 object from  the OLD-CISCO-SYSTEM-MIB. This object
    // is deprecated by cpmCPUTotal1minRev which has the changed range of value
    // (0..100). The type is interface{} with range: 0..100.
    CpmCPUTotal1min interface{}

    // The overall CPU busy percentage in the last 5 minute period. This object
    // deprecates the avgBusy5 object from  the OLD-CISCO-SYSTEM-MIB. This object
    // is deprecated by cpmCPUTotal5minRev which has the changed range  of value
    // (0..100). The type is interface{} with range: 0..100.
    CpmCPUTotal5min interface{}

    // The overall CPU busy percentage in the last 5 second period. This object
    // deprecates the object cpmCPUTotal5sec  and increases the value range to
    // (0..100). This object is deprecated by cpmCPUTotalMonIntervalValue. The
    // type is interface{} with range: 0..100. Units are percent.
    CpmCPUTotal5secRev interface{}

    // The overall CPU busy percentage in the last 1 minute period. This object
    // deprecates the object cpmCPUTotal1min  and increases the value range to
    // (0..100). The type is interface{} with range: 0..100. Units are percent.
    CpmCPUTotal1minRev interface{}

    // The overall CPU busy percentage in the last 5 minute period. This object
    // deprecates the object cpmCPUTotal5min  and increases the value range to
    // (0..100). The type is interface{} with range: 0..100. Units are percent.
    CpmCPUTotal5minRev interface{}

    // CPU usage monitoring interval. The value of this object in seconds
    // indicates the how often the  CPU utilization is calculated and monitored.
    // The type is interface{} with range: 0..4294967295. Units are seconds.
    CpmCPUMonInterval interface{}

    // The overall CPU busy percentage in the last cpmCPUMonInterval period.  This
    // object deprecates the object cpmCPUTotal5secRev. The type is interface{}
    // with range: 0..100. Units are percent.
    CpmCPUTotalMonIntervalValue interface{}

    // The overall CPU busy percentage in the interrupt context in the last
    // cpmCPUMonInterval period. The type is interface{} with range: 0..100. Units
    // are percent.
    CpmCPUInterruptMonIntervalValue interface{}

    // The overall CPU wide system memory which is currently under use. The type
    // is interface{} with range: 0..4294967295. Units are kilo-bytes.
    CpmCPUMemoryUsed interface{}

    // The overall CPU wide system memory which is currently free. The type is
    // interface{} with range: 0..4294967295. Units are kilo-bytes.
    CpmCPUMemoryFree interface{}

    // The overall CPU wide system memory which is reserved for kernel usage. The
    // type is interface{} with range: 0..4294967295. Units are kilo-bytes.
    CpmCPUMemoryKernelReserved interface{}

    // The lowest free memory that has been recorded since device has booted. The
    // type is interface{} with range: 0..4294967295. Units are bytes.
    CpmCPUMemoryLowest interface{}

    // This object represents the upper 32-bit of cpmCPUMemoryUsed. This object
    // needs to be supported only when the value of cpmCPUMemoryUsed exceeds
    // 32-bit, otherwise this object value would be set to 0. The type is
    // interface{} with range: 0..4294967295. Units are kilo-bytes.
    CpmCPUMemoryUsedOvrflw interface{}

    // The overall CPU wide system memory which is currently under use. This
    // object is a 64-bit version of cpmCPUMemoryUsed. The type is interface{}
    // with range: 0..18446744073709551615. Units are kilo-bytes.
    CpmCPUMemoryHCUsed interface{}

    // This object represents the upper 32-bit of cpmCPUMemoryFree. This object
    // needs to be supported only when the value of cpmCPUMemoryFree exceeds
    // 32-bit, otherwise this object value would be set to 0. The type is
    // interface{} with range: 0..4294967295. Units are kilo-bytes.
    CpmCPUMemoryFreeOvrflw interface{}

    // The overall CPU wide system memory which is currently free. This object is
    // a 64-bit version of cpmCPUMemoryFree. The type is interface{} with range:
    // 0..18446744073709551615. Units are kilo-bytes.
    CpmCPUMemoryHCFree interface{}

    // This object represents the upper 32-bit of cpmCPUMemoryKernelReserved. This
    // object needs  to be supported only when the value of 
    // cpmCPUMemoryKernelReserved exceeds 32-bit, otherwise this object value
    // would be set to 0. The type is interface{} with range: 0..4294967295. Units
    // are kilo-bytes.
    CpmCPUMemoryKernelReservedOvrflw interface{}

    // The overall CPU wide system memory which is reserved for kernel usage. This
    // object is a 64-bit version of cpmCPUMemoryKernelReserved. The type is
    // interface{} with range: 0..18446744073709551615. Units are kilo-bytes.
    CpmCPUMemoryHCKernelReserved interface{}

    // This object represents the upper 32-bit of cpmCPUMemoryLowest. This object
    // needs to be supported only when the value of cpmCPUMemoryLowest exceeds
    // 32-bit, otherwise this object value would be set to 0. The type is
    // interface{} with range: 0..4294967295. Units are kilo-bytes.
    CpmCPUMemoryLowestOvrflw interface{}

    // The lowest free memory that has been recorded since device has booted. This
    // object is a 64-bit version of cpmCPUMemoryLowest. The type is interface{}
    // with range: 0..18446744073709551615. Units are kilo-bytes.
    CpmCPUMemoryHCLowest interface{}

    // The overall CPU load Average in the last 1 minute period. The type is
    // interface{} with range: 0..4294967295. Units are hundredths of processes.
    CpmCPULoadAvg1min interface{}

    // The overall CPU load Average in the last 5 minutes period. The type is
    // interface{} with range: 0..4294967295. Units are hundredths of processes.
    CpmCPULoadAvg5min interface{}

    // The overall CPU load Average in the last 15 minutes period. The type is
    // interface{} with range: 0..4294967295. Units are hundredths of processes.
    CpmCPULoadAvg15min interface{}

    // The overall CPU wide system memory which is currently Committed. The type
    // is interface{} with range: 0..4294967295.
    CpmCPUMemoryCommitted interface{}

    // This object represents the upper 32-bit of cpmCPUMemoryCommitted. This
    // object needs to be supported only when the value of cpmCPUMemoryCommitted
    // exceeds 32-bit, otherwise this object value would be set to 0. The type is
    // interface{} with range: 0..4294967295.
    CpmCPUMemoryCommittedOvrflw interface{}

    // The overall CPU wide system memory which is currently committed. This
    // object is a 64-bit version of cpmCPUMemoryCommitted. The type is
    // interface{} with range: 0..18446744073709551615.
    CpmCPUMemoryHCCommitted interface{}
}

func (cpmCPUTotalEntry *CISCOPROCESSMIB_CpmCPUTotalTable_CpmCPUTotalEntry) GetEntityData() *types.CommonEntityData {
    cpmCPUTotalEntry.EntityData.YFilter = cpmCPUTotalEntry.YFilter
    cpmCPUTotalEntry.EntityData.YangName = "cpmCPUTotalEntry"
    cpmCPUTotalEntry.EntityData.BundleName = "cisco_ios_xe"
    cpmCPUTotalEntry.EntityData.ParentYangName = "cpmCPUTotalTable"
    cpmCPUTotalEntry.EntityData.SegmentPath = "cpmCPUTotalEntry" + types.AddKeyToken(cpmCPUTotalEntry.CpmCPUTotalIndex, "cpmCPUTotalIndex")
    cpmCPUTotalEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpmCPUTotalEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpmCPUTotalEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpmCPUTotalEntry.EntityData.Children = types.NewOrderedMap()
    cpmCPUTotalEntry.EntityData.Leafs = types.NewOrderedMap()
    cpmCPUTotalEntry.EntityData.Leafs.Append("cpmCPUTotalIndex", types.YLeaf{"CpmCPUTotalIndex", cpmCPUTotalEntry.CpmCPUTotalIndex})
    cpmCPUTotalEntry.EntityData.Leafs.Append("cpmCPUTotalPhysicalIndex", types.YLeaf{"CpmCPUTotalPhysicalIndex", cpmCPUTotalEntry.CpmCPUTotalPhysicalIndex})
    cpmCPUTotalEntry.EntityData.Leafs.Append("cpmCPUTotal5sec", types.YLeaf{"CpmCPUTotal5sec", cpmCPUTotalEntry.CpmCPUTotal5sec})
    cpmCPUTotalEntry.EntityData.Leafs.Append("cpmCPUTotal1min", types.YLeaf{"CpmCPUTotal1min", cpmCPUTotalEntry.CpmCPUTotal1min})
    cpmCPUTotalEntry.EntityData.Leafs.Append("cpmCPUTotal5min", types.YLeaf{"CpmCPUTotal5min", cpmCPUTotalEntry.CpmCPUTotal5min})
    cpmCPUTotalEntry.EntityData.Leafs.Append("cpmCPUTotal5secRev", types.YLeaf{"CpmCPUTotal5secRev", cpmCPUTotalEntry.CpmCPUTotal5secRev})
    cpmCPUTotalEntry.EntityData.Leafs.Append("cpmCPUTotal1minRev", types.YLeaf{"CpmCPUTotal1minRev", cpmCPUTotalEntry.CpmCPUTotal1minRev})
    cpmCPUTotalEntry.EntityData.Leafs.Append("cpmCPUTotal5minRev", types.YLeaf{"CpmCPUTotal5minRev", cpmCPUTotalEntry.CpmCPUTotal5minRev})
    cpmCPUTotalEntry.EntityData.Leafs.Append("cpmCPUMonInterval", types.YLeaf{"CpmCPUMonInterval", cpmCPUTotalEntry.CpmCPUMonInterval})
    cpmCPUTotalEntry.EntityData.Leafs.Append("cpmCPUTotalMonIntervalValue", types.YLeaf{"CpmCPUTotalMonIntervalValue", cpmCPUTotalEntry.CpmCPUTotalMonIntervalValue})
    cpmCPUTotalEntry.EntityData.Leafs.Append("cpmCPUInterruptMonIntervalValue", types.YLeaf{"CpmCPUInterruptMonIntervalValue", cpmCPUTotalEntry.CpmCPUInterruptMonIntervalValue})
    cpmCPUTotalEntry.EntityData.Leafs.Append("cpmCPUMemoryUsed", types.YLeaf{"CpmCPUMemoryUsed", cpmCPUTotalEntry.CpmCPUMemoryUsed})
    cpmCPUTotalEntry.EntityData.Leafs.Append("cpmCPUMemoryFree", types.YLeaf{"CpmCPUMemoryFree", cpmCPUTotalEntry.CpmCPUMemoryFree})
    cpmCPUTotalEntry.EntityData.Leafs.Append("cpmCPUMemoryKernelReserved", types.YLeaf{"CpmCPUMemoryKernelReserved", cpmCPUTotalEntry.CpmCPUMemoryKernelReserved})
    cpmCPUTotalEntry.EntityData.Leafs.Append("cpmCPUMemoryLowest", types.YLeaf{"CpmCPUMemoryLowest", cpmCPUTotalEntry.CpmCPUMemoryLowest})
    cpmCPUTotalEntry.EntityData.Leafs.Append("cpmCPUMemoryUsedOvrflw", types.YLeaf{"CpmCPUMemoryUsedOvrflw", cpmCPUTotalEntry.CpmCPUMemoryUsedOvrflw})
    cpmCPUTotalEntry.EntityData.Leafs.Append("cpmCPUMemoryHCUsed", types.YLeaf{"CpmCPUMemoryHCUsed", cpmCPUTotalEntry.CpmCPUMemoryHCUsed})
    cpmCPUTotalEntry.EntityData.Leafs.Append("cpmCPUMemoryFreeOvrflw", types.YLeaf{"CpmCPUMemoryFreeOvrflw", cpmCPUTotalEntry.CpmCPUMemoryFreeOvrflw})
    cpmCPUTotalEntry.EntityData.Leafs.Append("cpmCPUMemoryHCFree", types.YLeaf{"CpmCPUMemoryHCFree", cpmCPUTotalEntry.CpmCPUMemoryHCFree})
    cpmCPUTotalEntry.EntityData.Leafs.Append("cpmCPUMemoryKernelReservedOvrflw", types.YLeaf{"CpmCPUMemoryKernelReservedOvrflw", cpmCPUTotalEntry.CpmCPUMemoryKernelReservedOvrflw})
    cpmCPUTotalEntry.EntityData.Leafs.Append("cpmCPUMemoryHCKernelReserved", types.YLeaf{"CpmCPUMemoryHCKernelReserved", cpmCPUTotalEntry.CpmCPUMemoryHCKernelReserved})
    cpmCPUTotalEntry.EntityData.Leafs.Append("cpmCPUMemoryLowestOvrflw", types.YLeaf{"CpmCPUMemoryLowestOvrflw", cpmCPUTotalEntry.CpmCPUMemoryLowestOvrflw})
    cpmCPUTotalEntry.EntityData.Leafs.Append("cpmCPUMemoryHCLowest", types.YLeaf{"CpmCPUMemoryHCLowest", cpmCPUTotalEntry.CpmCPUMemoryHCLowest})
    cpmCPUTotalEntry.EntityData.Leafs.Append("cpmCPULoadAvg1min", types.YLeaf{"CpmCPULoadAvg1min", cpmCPUTotalEntry.CpmCPULoadAvg1min})
    cpmCPUTotalEntry.EntityData.Leafs.Append("cpmCPULoadAvg5min", types.YLeaf{"CpmCPULoadAvg5min", cpmCPUTotalEntry.CpmCPULoadAvg5min})
    cpmCPUTotalEntry.EntityData.Leafs.Append("cpmCPULoadAvg15min", types.YLeaf{"CpmCPULoadAvg15min", cpmCPUTotalEntry.CpmCPULoadAvg15min})
    cpmCPUTotalEntry.EntityData.Leafs.Append("cpmCPUMemoryCommitted", types.YLeaf{"CpmCPUMemoryCommitted", cpmCPUTotalEntry.CpmCPUMemoryCommitted})
    cpmCPUTotalEntry.EntityData.Leafs.Append("cpmCPUMemoryCommittedOvrflw", types.YLeaf{"CpmCPUMemoryCommittedOvrflw", cpmCPUTotalEntry.CpmCPUMemoryCommittedOvrflw})
    cpmCPUTotalEntry.EntityData.Leafs.Append("cpmCPUMemoryHCCommitted", types.YLeaf{"CpmCPUMemoryHCCommitted", cpmCPUTotalEntry.CpmCPUMemoryHCCommitted})

    cpmCPUTotalEntry.EntityData.YListKeys = []string {"CpmCPUTotalIndex"}

    return &(cpmCPUTotalEntry.EntityData)
}

// CISCOPROCESSMIB_CpmCoreTable
// A table of per-Core statistics.
type CISCOPROCESSMIB_CpmCoreTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Overall information about the Core load. Entries in this table could come
    // and go as Cores go online or offline. The type is slice of
    // CISCOPROCESSMIB_CpmCoreTable_CpmCoreEntry.
    CpmCoreEntry []*CISCOPROCESSMIB_CpmCoreTable_CpmCoreEntry
}

func (cpmCoreTable *CISCOPROCESSMIB_CpmCoreTable) GetEntityData() *types.CommonEntityData {
    cpmCoreTable.EntityData.YFilter = cpmCoreTable.YFilter
    cpmCoreTable.EntityData.YangName = "cpmCoreTable"
    cpmCoreTable.EntityData.BundleName = "cisco_ios_xe"
    cpmCoreTable.EntityData.ParentYangName = "CISCO-PROCESS-MIB"
    cpmCoreTable.EntityData.SegmentPath = "cpmCoreTable"
    cpmCoreTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpmCoreTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpmCoreTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpmCoreTable.EntityData.Children = types.NewOrderedMap()
    cpmCoreTable.EntityData.Children.Append("cpmCoreEntry", types.YChild{"CpmCoreEntry", nil})
    for i := range cpmCoreTable.CpmCoreEntry {
        cpmCoreTable.EntityData.Children.Append(types.GetSegmentPath(cpmCoreTable.CpmCoreEntry[i]), types.YChild{"CpmCoreEntry", cpmCoreTable.CpmCoreEntry[i]})
    }
    cpmCoreTable.EntityData.Leafs = types.NewOrderedMap()

    cpmCoreTable.EntityData.YListKeys = []string {}

    return &(cpmCoreTable.EntityData)
}

// CISCOPROCESSMIB_CpmCoreTable_CpmCoreEntry
// Overall information about the Core load. Entries in this
// table could come and go as Cores go online or offline.
type CISCOPROCESSMIB_CpmCoreTable_CpmCoreEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // cisco_process_mib.CISCOPROCESSMIB_CpmCPUTotalTable_CpmCPUTotalEntry_CpmCPUTotalIndex
    CpmCPUTotalIndex interface{}

    // This attribute is a key. An index that uniquely represents a Core (or group
    // of Cores) whose Core load information is reported by a row in this table.
    // This index is assigned arbitrarily by the engine and is not saved over
    // reboots. The type is interface{} with range: 0..4294967295.
    CpmCoreIndex interface{}

    // The entCorePhysicalIndex of the physical entity for which the Core
    // statistics in this entry are maintained. The physical entity can be a CPU
    // chip, a group of CPUs, a CPU card etc. The exact type of this entity is
    // described by its entPhysicalVendorType value. If the Core statistics in
    // this entry correspond to more than one physical entity (or to no physical
    // entity), or if the entPhysicalTable is not supported on the SNMP agent, the
    // value of this object must be zero. The type is interface{} with range:
    // 0..2147483647.
    CpmCorePhysicalIndex interface{}

    // The overall Core busy percentage in the last 5 second period. The type is
    // interface{} with range: 0..100.
    CpmCore5sec interface{}

    // The overall Core busy percentage in the last 1 minute period. The type is
    // interface{} with range: 0..100.
    CpmCore1min interface{}

    // The overall Core busy percentage in the last 5 minute period. The type is
    // interface{} with range: 0..100.
    CpmCore5min interface{}

    // The overall Core load Average in the last 1 minute period. The type is
    // interface{} with range: 0..4294967295. Units are hundredths of processes.
    CpmCoreLoadAvg1min interface{}

    // The overall Core load Average in the last 5 minutes period. The type is
    // interface{} with range: 0..4294967295. Units are hundredths of processes.
    CpmCoreLoadAvg5min interface{}

    // The overall Core load Average in the last 15 minutes period. The type is
    // interface{} with range: 0..4294967295. Units are hundredths of processes.
    CpmCoreLoadAvg15min interface{}
}

func (cpmCoreEntry *CISCOPROCESSMIB_CpmCoreTable_CpmCoreEntry) GetEntityData() *types.CommonEntityData {
    cpmCoreEntry.EntityData.YFilter = cpmCoreEntry.YFilter
    cpmCoreEntry.EntityData.YangName = "cpmCoreEntry"
    cpmCoreEntry.EntityData.BundleName = "cisco_ios_xe"
    cpmCoreEntry.EntityData.ParentYangName = "cpmCoreTable"
    cpmCoreEntry.EntityData.SegmentPath = "cpmCoreEntry" + types.AddKeyToken(cpmCoreEntry.CpmCPUTotalIndex, "cpmCPUTotalIndex") + types.AddKeyToken(cpmCoreEntry.CpmCoreIndex, "cpmCoreIndex")
    cpmCoreEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpmCoreEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpmCoreEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpmCoreEntry.EntityData.Children = types.NewOrderedMap()
    cpmCoreEntry.EntityData.Leafs = types.NewOrderedMap()
    cpmCoreEntry.EntityData.Leafs.Append("cpmCPUTotalIndex", types.YLeaf{"CpmCPUTotalIndex", cpmCoreEntry.CpmCPUTotalIndex})
    cpmCoreEntry.EntityData.Leafs.Append("cpmCoreIndex", types.YLeaf{"CpmCoreIndex", cpmCoreEntry.CpmCoreIndex})
    cpmCoreEntry.EntityData.Leafs.Append("cpmCorePhysicalIndex", types.YLeaf{"CpmCorePhysicalIndex", cpmCoreEntry.CpmCorePhysicalIndex})
    cpmCoreEntry.EntityData.Leafs.Append("cpmCore5sec", types.YLeaf{"CpmCore5sec", cpmCoreEntry.CpmCore5sec})
    cpmCoreEntry.EntityData.Leafs.Append("cpmCore1min", types.YLeaf{"CpmCore1min", cpmCoreEntry.CpmCore1min})
    cpmCoreEntry.EntityData.Leafs.Append("cpmCore5min", types.YLeaf{"CpmCore5min", cpmCoreEntry.CpmCore5min})
    cpmCoreEntry.EntityData.Leafs.Append("cpmCoreLoadAvg1min", types.YLeaf{"CpmCoreLoadAvg1min", cpmCoreEntry.CpmCoreLoadAvg1min})
    cpmCoreEntry.EntityData.Leafs.Append("cpmCoreLoadAvg5min", types.YLeaf{"CpmCoreLoadAvg5min", cpmCoreEntry.CpmCoreLoadAvg5min})
    cpmCoreEntry.EntityData.Leafs.Append("cpmCoreLoadAvg15min", types.YLeaf{"CpmCoreLoadAvg15min", cpmCoreEntry.CpmCoreLoadAvg15min})

    cpmCoreEntry.EntityData.YListKeys = []string {"CpmCPUTotalIndex", "CpmCoreIndex"}

    return &(cpmCoreEntry.EntityData)
}

// CISCOPROCESSMIB_CpmProcessTable
// A table of generic information on all active
// processes on this device.
type CISCOPROCESSMIB_CpmProcessTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Generic information about an active process on this device. Entries in this
    // table come and go as processes are  created and destroyed by the device.
    // The type is slice of CISCOPROCESSMIB_CpmProcessTable_CpmProcessEntry.
    CpmProcessEntry []*CISCOPROCESSMIB_CpmProcessTable_CpmProcessEntry
}

func (cpmProcessTable *CISCOPROCESSMIB_CpmProcessTable) GetEntityData() *types.CommonEntityData {
    cpmProcessTable.EntityData.YFilter = cpmProcessTable.YFilter
    cpmProcessTable.EntityData.YangName = "cpmProcessTable"
    cpmProcessTable.EntityData.BundleName = "cisco_ios_xe"
    cpmProcessTable.EntityData.ParentYangName = "CISCO-PROCESS-MIB"
    cpmProcessTable.EntityData.SegmentPath = "cpmProcessTable"
    cpmProcessTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpmProcessTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpmProcessTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpmProcessTable.EntityData.Children = types.NewOrderedMap()
    cpmProcessTable.EntityData.Children.Append("cpmProcessEntry", types.YChild{"CpmProcessEntry", nil})
    for i := range cpmProcessTable.CpmProcessEntry {
        cpmProcessTable.EntityData.Children.Append(types.GetSegmentPath(cpmProcessTable.CpmProcessEntry[i]), types.YChild{"CpmProcessEntry", cpmProcessTable.CpmProcessEntry[i]})
    }
    cpmProcessTable.EntityData.Leafs = types.NewOrderedMap()

    cpmProcessTable.EntityData.YListKeys = []string {}

    return &(cpmProcessTable.EntityData)
}

// CISCOPROCESSMIB_CpmProcessTable_CpmProcessEntry
// Generic information about an active process on this
// device. Entries in this table come and go as processes are 
// created and destroyed by the device.
type CISCOPROCESSMIB_CpmProcessTable_CpmProcessEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // cisco_process_mib.CISCOPROCESSMIB_CpmCPUTotalTable_CpmCPUTotalEntry_CpmCPUTotalIndex
    CpmCPUTotalIndex interface{}

    // This attribute is a key. This object contains the process ID.
    // cpmTimeCreated should be checked against the last time it was polled, and
    // if it has changed the PID has been reused and the entire entry should be
    // polled again. The type is interface{} with range: 0..4294967295.
    CpmProcessPID interface{}

    // The name associated with this process. If the name is longer than 32
    // characters, it will be truncated to the first 31 characters, and a `*' will
    // be appended as the last character to imply this is a truncated process
    // name. The type is string with length: 1..32.
    CpmProcessName interface{}

    // Average elapsed CPU time in microseconds when the process was active. This
    // object is deprecated by cpmProcessAverageUSecs. The type is interface{}
    // with range: 0..4294967295. Units are microseconds.
    CpmProcessuSecs interface{}

    // The time when the process was created. The process ID and the time when the
    // process was created, uniquely  identifies a process. The type is
    // interface{} with range: 0..4294967295.
    CpmProcessTimeCreated interface{}

    // Average elapsed CPU time in microseconds when the process was active. This
    // object deprecates the object cpmProcessuSecs. The type is interface{} with
    // range: 0..4294967295. Units are microseconds.
    CpmProcessAverageUSecs interface{}

    // The sum of all the dynamically allocated memory that this process has
    // received from the system. This includes memory that may have been returned.
    // The sum of freed memory is provided by cpmProcExtMemFreed. This object is
    // deprecated by cpmProcExtMemAllocatedRev. The type is interface{} with
    // range: 0..4294967295. Units are bytes.
    CpmProcExtMemAllocated interface{}

    // The sum of all memory that this process has returned to the system. This
    // object is deprecated by  cpmProcExtMemFreedRev. The type is interface{}
    // with range: 0..4294967295. Units are bytes.
    CpmProcExtMemFreed interface{}

    // The number of times since cpmTimeCreated that the process has been invoked.
    // This object is deprecated by cpmProcExtInvokedRev. The type is interface{}
    // with range: 0..4294967295.
    CpmProcExtInvoked interface{}

    // The amount of CPU time the process has used, in microseconds. This object
    // is deprecated by cpmProcExtRuntimeRev. The type is interface{} with range:
    // 0..4294967295. Units are microseconds.
    CpmProcExtRuntime interface{}

    // This object provides a general idea of how busy a process caused the
    // processor to be over a 5  second period. It is determined as a weighted 
    // decaying average of the current idle time over  the longest idle time. Note
    // that this information  should be used as an estimate only. This object is 
    // deprecated by cpmProcExtUtil5SecRev which has the  changed range of value
    // (0..100). The type is interface{} with range: 0..100.
    CpmProcExtUtil5Sec interface{}

    // This object provides a general idea of how busy a process caused the
    // processor to be over a 1  minute period. It is determined as a weighted 
    // decaying average of the current idle time over the  longest idle time. Note
    // that this information  should be used as an estimate only. This object is 
    // deprecated by cpmProcExtUtil1MinRev which has the changed range of value
    // (0..100). The type is interface{} with range: 0..100.
    CpmProcExtUtil1Min interface{}

    // This object provides a general idea of how busy a process caused the
    // processor to be over a 5  minute period. It is determined as a weighted 
    // decaying average of the current idle time over  the longest idle time. Note
    // that this information  should be used as an estimate only. This object is
    // deprecated by cpmProcExtUtil5MinRev which has the changed range of value
    // (0..100). The type is interface{} with range: 0..100.
    CpmProcExtUtil5Min interface{}

    // The priority level at which the process is running. This object is
    // deprecated by cpmProcExtPriorityRev. The type is CpmProcExtPriority.
    CpmProcExtPriority interface{}
}

func (cpmProcessEntry *CISCOPROCESSMIB_CpmProcessTable_CpmProcessEntry) GetEntityData() *types.CommonEntityData {
    cpmProcessEntry.EntityData.YFilter = cpmProcessEntry.YFilter
    cpmProcessEntry.EntityData.YangName = "cpmProcessEntry"
    cpmProcessEntry.EntityData.BundleName = "cisco_ios_xe"
    cpmProcessEntry.EntityData.ParentYangName = "cpmProcessTable"
    cpmProcessEntry.EntityData.SegmentPath = "cpmProcessEntry" + types.AddKeyToken(cpmProcessEntry.CpmCPUTotalIndex, "cpmCPUTotalIndex") + types.AddKeyToken(cpmProcessEntry.CpmProcessPID, "cpmProcessPID")
    cpmProcessEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpmProcessEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpmProcessEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpmProcessEntry.EntityData.Children = types.NewOrderedMap()
    cpmProcessEntry.EntityData.Leafs = types.NewOrderedMap()
    cpmProcessEntry.EntityData.Leafs.Append("cpmCPUTotalIndex", types.YLeaf{"CpmCPUTotalIndex", cpmProcessEntry.CpmCPUTotalIndex})
    cpmProcessEntry.EntityData.Leafs.Append("cpmProcessPID", types.YLeaf{"CpmProcessPID", cpmProcessEntry.CpmProcessPID})
    cpmProcessEntry.EntityData.Leafs.Append("cpmProcessName", types.YLeaf{"CpmProcessName", cpmProcessEntry.CpmProcessName})
    cpmProcessEntry.EntityData.Leafs.Append("cpmProcessuSecs", types.YLeaf{"CpmProcessuSecs", cpmProcessEntry.CpmProcessuSecs})
    cpmProcessEntry.EntityData.Leafs.Append("cpmProcessTimeCreated", types.YLeaf{"CpmProcessTimeCreated", cpmProcessEntry.CpmProcessTimeCreated})
    cpmProcessEntry.EntityData.Leafs.Append("cpmProcessAverageUSecs", types.YLeaf{"CpmProcessAverageUSecs", cpmProcessEntry.CpmProcessAverageUSecs})
    cpmProcessEntry.EntityData.Leafs.Append("cpmProcExtMemAllocated", types.YLeaf{"CpmProcExtMemAllocated", cpmProcessEntry.CpmProcExtMemAllocated})
    cpmProcessEntry.EntityData.Leafs.Append("cpmProcExtMemFreed", types.YLeaf{"CpmProcExtMemFreed", cpmProcessEntry.CpmProcExtMemFreed})
    cpmProcessEntry.EntityData.Leafs.Append("cpmProcExtInvoked", types.YLeaf{"CpmProcExtInvoked", cpmProcessEntry.CpmProcExtInvoked})
    cpmProcessEntry.EntityData.Leafs.Append("cpmProcExtRuntime", types.YLeaf{"CpmProcExtRuntime", cpmProcessEntry.CpmProcExtRuntime})
    cpmProcessEntry.EntityData.Leafs.Append("cpmProcExtUtil5Sec", types.YLeaf{"CpmProcExtUtil5Sec", cpmProcessEntry.CpmProcExtUtil5Sec})
    cpmProcessEntry.EntityData.Leafs.Append("cpmProcExtUtil1Min", types.YLeaf{"CpmProcExtUtil1Min", cpmProcessEntry.CpmProcExtUtil1Min})
    cpmProcessEntry.EntityData.Leafs.Append("cpmProcExtUtil5Min", types.YLeaf{"CpmProcExtUtil5Min", cpmProcessEntry.CpmProcExtUtil5Min})
    cpmProcessEntry.EntityData.Leafs.Append("cpmProcExtPriority", types.YLeaf{"CpmProcExtPriority", cpmProcessEntry.CpmProcExtPriority})

    cpmProcessEntry.EntityData.YListKeys = []string {"CpmCPUTotalIndex", "CpmProcessPID"}

    return &(cpmProcessEntry.EntityData)
}

// CISCOPROCESSMIB_CpmProcessTable_CpmProcessEntry_CpmProcExtPriority represents cpmProcExtPriorityRev.
type CISCOPROCESSMIB_CpmProcessTable_CpmProcessEntry_CpmProcExtPriority string

const (
    CISCOPROCESSMIB_CpmProcessTable_CpmProcessEntry_CpmProcExtPriority_critical CISCOPROCESSMIB_CpmProcessTable_CpmProcessEntry_CpmProcExtPriority = "critical"

    CISCOPROCESSMIB_CpmProcessTable_CpmProcessEntry_CpmProcExtPriority_high CISCOPROCESSMIB_CpmProcessTable_CpmProcessEntry_CpmProcExtPriority = "high"

    CISCOPROCESSMIB_CpmProcessTable_CpmProcessEntry_CpmProcExtPriority_normal CISCOPROCESSMIB_CpmProcessTable_CpmProcessEntry_CpmProcExtPriority = "normal"

    CISCOPROCESSMIB_CpmProcessTable_CpmProcessEntry_CpmProcExtPriority_low CISCOPROCESSMIB_CpmProcessTable_CpmProcessEntry_CpmProcExtPriority = "low"

    CISCOPROCESSMIB_CpmProcessTable_CpmProcessEntry_CpmProcExtPriority_notAssigned CISCOPROCESSMIB_CpmProcessTable_CpmProcessEntry_CpmProcExtPriority = "notAssigned"
)

// CISCOPROCESSMIB_CpmProcessExtRevTable
// This table contains information that may or may
// not be available on all cisco devices. It contains
// additional objects for the more general
// cpmProcessTable. This object deprecates 
// cpmProcessExtTable.
type CISCOPROCESSMIB_CpmProcessExtRevTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry containing additional information for a particular process. This
    // object deprecates  cpmProcessExtEntry. The type is slice of
    // CISCOPROCESSMIB_CpmProcessExtRevTable_CpmProcessExtRevEntry.
    CpmProcessExtRevEntry []*CISCOPROCESSMIB_CpmProcessExtRevTable_CpmProcessExtRevEntry
}

func (cpmProcessExtRevTable *CISCOPROCESSMIB_CpmProcessExtRevTable) GetEntityData() *types.CommonEntityData {
    cpmProcessExtRevTable.EntityData.YFilter = cpmProcessExtRevTable.YFilter
    cpmProcessExtRevTable.EntityData.YangName = "cpmProcessExtRevTable"
    cpmProcessExtRevTable.EntityData.BundleName = "cisco_ios_xe"
    cpmProcessExtRevTable.EntityData.ParentYangName = "CISCO-PROCESS-MIB"
    cpmProcessExtRevTable.EntityData.SegmentPath = "cpmProcessExtRevTable"
    cpmProcessExtRevTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpmProcessExtRevTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpmProcessExtRevTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpmProcessExtRevTable.EntityData.Children = types.NewOrderedMap()
    cpmProcessExtRevTable.EntityData.Children.Append("cpmProcessExtRevEntry", types.YChild{"CpmProcessExtRevEntry", nil})
    for i := range cpmProcessExtRevTable.CpmProcessExtRevEntry {
        cpmProcessExtRevTable.EntityData.Children.Append(types.GetSegmentPath(cpmProcessExtRevTable.CpmProcessExtRevEntry[i]), types.YChild{"CpmProcessExtRevEntry", cpmProcessExtRevTable.CpmProcessExtRevEntry[i]})
    }
    cpmProcessExtRevTable.EntityData.Leafs = types.NewOrderedMap()

    cpmProcessExtRevTable.EntityData.YListKeys = []string {}

    return &(cpmProcessExtRevTable.EntityData)
}

// CISCOPROCESSMIB_CpmProcessExtRevTable_CpmProcessExtRevEntry
// An entry containing additional information for
// a particular process. This object deprecates 
// cpmProcessExtEntry.
type CISCOPROCESSMIB_CpmProcessExtRevTable_CpmProcessExtRevEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // cisco_process_mib.CISCOPROCESSMIB_CpmCPUTotalTable_CpmCPUTotalEntry_CpmCPUTotalIndex
    CpmCPUTotalIndex interface{}

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // cisco_process_mib.CISCOPROCESSMIB_CpmProcessTable_CpmProcessEntry_CpmProcessPID
    CpmProcessPID interface{}

    // The sum of all the dynamically allocated memory that this process has
    // received from the system. This includes memory that may have been returned.
    // The sum of freed memory is provided by cpmProcExtMemFreedRev. This object
    // deprecates cpmProcExtMemAllocated. The type is interface{} with range:
    // 0..4294967295. Units are bytes.
    CpmProcExtMemAllocatedRev interface{}

    // The sum of all memory that this process has returned to the system. This
    // object  deprecates  cpmProcExtMemFreed. The type is interface{} with range:
    // 0..4294967295. Units are bytes.
    CpmProcExtMemFreedRev interface{}

    // The number of times since cpmTimeCreated that the process has been invoked.
    // This object  deprecates cpmProcExtInvoked. The type is interface{} with
    // range: 0..4294967295.
    CpmProcExtInvokedRev interface{}

    // The amount of CPU time the process has used, in microseconds. This object
    // deprecates cpmProcExtRuntime. The type is interface{} with range:
    // 0..4294967295. Units are microseconds.
    CpmProcExtRuntimeRev interface{}

    // This object provides a general idea of how busy a process caused the
    // processor to be over a 5  second period. It is determined as a weighted 
    // decaying average of the current idle time over  the longest idle time. Note
    // that this information  should be used as an estimate only. This object
    // deprecates cpmProcExtUtil5Sec and increases the  value range to (0..100).
    // The type is interface{} with range: 0..100. Units are percent.
    CpmProcExtUtil5SecRev interface{}

    // This object provides a general idea of how busy a process caused the
    // processor to be over a 1  minute period. It is determined as a weighted 
    // decaying average of the current idle time over the  longest idle time. Note
    // that this information  should be used as an estimate only. This object 
    // deprecates cpmProcExtUtil1Min and increases the value range to (0..100).
    // The type is interface{} with range: 0..100. Units are percent.
    CpmProcExtUtil1MinRev interface{}

    // This object provides a general idea of how busy a process caused the
    // processor to be over a 5  minute period. It is determined as a weighted 
    // decaying average of the current idle time over  the longest idle time. Note
    // that this information  should be used as an estimate only. This object
    // deprecates cpmProcExtUtil5Min and increases the value range to (0..100).
    // The type is interface{} with range: 0..100. Units are percent.
    CpmProcExtUtil5MinRev interface{}

    // The priority level at  which the process is running. This object deprecates
    // cpmProcExtPriority. The type is CpmProcExtPriorityRev.
    CpmProcExtPriorityRev interface{}

    // This indicates the kind of process in context. The type is CpmProcessType.
    CpmProcessType interface{}

    // This indicates whether respawn of a process is enabled or not. If enabled
    // the process in context repawns after it has crashed/stopped. The type is
    // interface{} with range: 0..2.
    CpmProcessRespawn interface{}

    // This indicates the number of times the process has respawned/restarted. The
    // type is interface{} with range: 0..4294967295.
    CpmProcessRespawnCount interface{}

    // This indicates the number of times a process has restarted after the last
    // patch is applied. This is to  determine the stability of the last patch.
    // The type is interface{} with range: 0..4294967295.
    CpmProcessRespawnAfterLastPatch interface{}

    // This indicates the part of process memory to be dumped when a process
    // crashes. The process  memory is used for debugging purposes to trace the 
    // root cause of the crash. sparse        - Some operating systems support
    // minimal                 dump of process core like register                
    // info, partial stack, partial memory                 pages especially for
    // critical process                 to facilitate faster process restart. The
    // type is CpmProcessMemoryCore.
    CpmProcessMemoryCore interface{}

    // This indicate the user that has last restarted the process or has taken
    // running coredump of the process. The type is string.
    CpmProcessLastRestartUser interface{}

    // This indicates the text memory of a process and all its shared objects. The
    // type is interface{} with range: 0..4294967295. Units are kilo-bytes.
    CpmProcessTextSegmentSize interface{}

    // This indicates the data segment of a process and all its shared objects.
    // The type is interface{} with range: 0..4294967295. Units are kilo-bytes.
    CpmProcessDataSegmentSize interface{}

    // This indicates the amount of stack memory used by the process. The type is
    // interface{} with range: 0..4294967295. Units are kilo-bytes.
    CpmProcessStackSize interface{}

    // This indicates the amount of dynamic memory being used by the process. The
    // type is interface{} with range: 0..4294967295. Units are kilo-bytes.
    CpmProcessDynamicMemorySize interface{}

    // This object represents the upper 32-bit of cpmProcExtMemAllocatedRev. This
    // object needs to be supported only when the value of 
    // cpmProcExtMemAllocatedRev exceeds 32-bit,  otherwise this object value
    // would be set to 0. The type is interface{} with range: 0..4294967295. Units
    // are bytes.
    CpmProcExtMemAllocatedRevOvrflw interface{}

    // The sum of all the dynamically allocated memory that this process has
    // received from the system. This includes memory that may have been returned.
    // This object is a 64-bit version of cpmProcExtMemAllocatedRev. The type is
    // interface{} with range: 0..18446744073709551615. Units are bytes.
    CpmProcExtHCMemAllocatedRev interface{}

    // This object represents the upper 32-bit of cpmProcExtMemFreedRev. This
    // object needs to  be supported only when the value of cpmProcExtMemFreedRev
    // exceeds 32-bit,otherwise this object value would be set to 0. The type is
    // interface{} with range: 0..4294967295. Units are bytes.
    CpmProcExtMemFreedRevOvrflw interface{}

    // The sum of all memory that this process has returned to the system. This
    // object is a 64-bit version of cpmProcExtMemFreedRev. The type is
    // interface{} with range: 0..18446744073709551615. Units are bytes.
    CpmProcExtHCMemFreedRev interface{}

    // This object represents the upper 32-bit of cpmProcessTextSegmentSize. This
    // object needs to be supported only when the value of 
    // cpmProcessTextSegmentSize exceeds 32-bit,  otherwise this object value
    // would be set to 0. The type is interface{} with range: 0..4294967295. Units
    // are kilo-bytes.
    CpmProcessTextSegmentSizeOvrflw interface{}

    // This indicates the text memory of a process and all its shared objects.
    // This object is a 64-bit version of cpmProcessTextSegmentSize. The type is
    // interface{} with range: 0..18446744073709551615. Units are kilo-bytes.
    CpmProcessHCTextSegmentSize interface{}

    // This object represents the upper 32-bit of cpmProcessDataSegmentSize. This
    // object needs to be supported only when the value of 
    // cpmProcessDataSegmentSize exceeds 32-bit,  otherwise this object value
    // would be set to 0. The type is interface{} with range: 0..4294967295. Units
    // are kilo-bytes.
    CpmProcessDataSegmentSizeOvrflw interface{}

    // This indicates the data segment of a process and all its shared objects..
    // This object is a 64-bit version of cpmProcessDataSegmentSize. The type is
    // interface{} with range: 0..18446744073709551615. Units are kilo-bytes.
    CpmProcessHCDataSegmentSize interface{}

    // This object represents the upper 32-bit of cpmProcessStackSize. This object
    // needs to be supported only when the value of cpmProcessStackSize exceeds
    // 32-bit, otherwise this object value would be set to 0. The type is
    // interface{} with range: 0..4294967295. Units are kilo-bytes.
    CpmProcessStackSizeOvrflw interface{}

    // This indicates the amount of stack memory used by the process. This object
    // is a 64-bit version of cpmProcessStackSize. The type is interface{} with
    // range: 0..18446744073709551615. Units are kilo-bytes.
    CpmProcessHCStackSize interface{}

    // This object represents the upper 32-bit of cpmProcessDynamicMemorySize.
    // This object needs to be supported only when the value of 
    // cpmProcessDynamicMemorySize exceeds 32-bit,  otherwise this object value
    // would be set to 0. The type is interface{} with range: 0..4294967295. Units
    // are kilo-bytes.
    CpmProcessDynamicMemorySizeOvrflw interface{}

    // This indicates the amount of dynamic memory being used by the process. This
    // object is a 64-bit version of cpmProcessDynamicMemorySize. The type is
    // interface{} with range: 0..18446744073709551615. Units are kilo-bytes.
    CpmProcessHCDynamicMemorySize interface{}
}

func (cpmProcessExtRevEntry *CISCOPROCESSMIB_CpmProcessExtRevTable_CpmProcessExtRevEntry) GetEntityData() *types.CommonEntityData {
    cpmProcessExtRevEntry.EntityData.YFilter = cpmProcessExtRevEntry.YFilter
    cpmProcessExtRevEntry.EntityData.YangName = "cpmProcessExtRevEntry"
    cpmProcessExtRevEntry.EntityData.BundleName = "cisco_ios_xe"
    cpmProcessExtRevEntry.EntityData.ParentYangName = "cpmProcessExtRevTable"
    cpmProcessExtRevEntry.EntityData.SegmentPath = "cpmProcessExtRevEntry" + types.AddKeyToken(cpmProcessExtRevEntry.CpmCPUTotalIndex, "cpmCPUTotalIndex") + types.AddKeyToken(cpmProcessExtRevEntry.CpmProcessPID, "cpmProcessPID")
    cpmProcessExtRevEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpmProcessExtRevEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpmProcessExtRevEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpmProcessExtRevEntry.EntityData.Children = types.NewOrderedMap()
    cpmProcessExtRevEntry.EntityData.Leafs = types.NewOrderedMap()
    cpmProcessExtRevEntry.EntityData.Leafs.Append("cpmCPUTotalIndex", types.YLeaf{"CpmCPUTotalIndex", cpmProcessExtRevEntry.CpmCPUTotalIndex})
    cpmProcessExtRevEntry.EntityData.Leafs.Append("cpmProcessPID", types.YLeaf{"CpmProcessPID", cpmProcessExtRevEntry.CpmProcessPID})
    cpmProcessExtRevEntry.EntityData.Leafs.Append("cpmProcExtMemAllocatedRev", types.YLeaf{"CpmProcExtMemAllocatedRev", cpmProcessExtRevEntry.CpmProcExtMemAllocatedRev})
    cpmProcessExtRevEntry.EntityData.Leafs.Append("cpmProcExtMemFreedRev", types.YLeaf{"CpmProcExtMemFreedRev", cpmProcessExtRevEntry.CpmProcExtMemFreedRev})
    cpmProcessExtRevEntry.EntityData.Leafs.Append("cpmProcExtInvokedRev", types.YLeaf{"CpmProcExtInvokedRev", cpmProcessExtRevEntry.CpmProcExtInvokedRev})
    cpmProcessExtRevEntry.EntityData.Leafs.Append("cpmProcExtRuntimeRev", types.YLeaf{"CpmProcExtRuntimeRev", cpmProcessExtRevEntry.CpmProcExtRuntimeRev})
    cpmProcessExtRevEntry.EntityData.Leafs.Append("cpmProcExtUtil5SecRev", types.YLeaf{"CpmProcExtUtil5SecRev", cpmProcessExtRevEntry.CpmProcExtUtil5SecRev})
    cpmProcessExtRevEntry.EntityData.Leafs.Append("cpmProcExtUtil1MinRev", types.YLeaf{"CpmProcExtUtil1MinRev", cpmProcessExtRevEntry.CpmProcExtUtil1MinRev})
    cpmProcessExtRevEntry.EntityData.Leafs.Append("cpmProcExtUtil5MinRev", types.YLeaf{"CpmProcExtUtil5MinRev", cpmProcessExtRevEntry.CpmProcExtUtil5MinRev})
    cpmProcessExtRevEntry.EntityData.Leafs.Append("cpmProcExtPriorityRev", types.YLeaf{"CpmProcExtPriorityRev", cpmProcessExtRevEntry.CpmProcExtPriorityRev})
    cpmProcessExtRevEntry.EntityData.Leafs.Append("cpmProcessType", types.YLeaf{"CpmProcessType", cpmProcessExtRevEntry.CpmProcessType})
    cpmProcessExtRevEntry.EntityData.Leafs.Append("cpmProcessRespawn", types.YLeaf{"CpmProcessRespawn", cpmProcessExtRevEntry.CpmProcessRespawn})
    cpmProcessExtRevEntry.EntityData.Leafs.Append("cpmProcessRespawnCount", types.YLeaf{"CpmProcessRespawnCount", cpmProcessExtRevEntry.CpmProcessRespawnCount})
    cpmProcessExtRevEntry.EntityData.Leafs.Append("cpmProcessRespawnAfterLastPatch", types.YLeaf{"CpmProcessRespawnAfterLastPatch", cpmProcessExtRevEntry.CpmProcessRespawnAfterLastPatch})
    cpmProcessExtRevEntry.EntityData.Leafs.Append("cpmProcessMemoryCore", types.YLeaf{"CpmProcessMemoryCore", cpmProcessExtRevEntry.CpmProcessMemoryCore})
    cpmProcessExtRevEntry.EntityData.Leafs.Append("cpmProcessLastRestartUser", types.YLeaf{"CpmProcessLastRestartUser", cpmProcessExtRevEntry.CpmProcessLastRestartUser})
    cpmProcessExtRevEntry.EntityData.Leafs.Append("cpmProcessTextSegmentSize", types.YLeaf{"CpmProcessTextSegmentSize", cpmProcessExtRevEntry.CpmProcessTextSegmentSize})
    cpmProcessExtRevEntry.EntityData.Leafs.Append("cpmProcessDataSegmentSize", types.YLeaf{"CpmProcessDataSegmentSize", cpmProcessExtRevEntry.CpmProcessDataSegmentSize})
    cpmProcessExtRevEntry.EntityData.Leafs.Append("cpmProcessStackSize", types.YLeaf{"CpmProcessStackSize", cpmProcessExtRevEntry.CpmProcessStackSize})
    cpmProcessExtRevEntry.EntityData.Leafs.Append("cpmProcessDynamicMemorySize", types.YLeaf{"CpmProcessDynamicMemorySize", cpmProcessExtRevEntry.CpmProcessDynamicMemorySize})
    cpmProcessExtRevEntry.EntityData.Leafs.Append("cpmProcExtMemAllocatedRevOvrflw", types.YLeaf{"CpmProcExtMemAllocatedRevOvrflw", cpmProcessExtRevEntry.CpmProcExtMemAllocatedRevOvrflw})
    cpmProcessExtRevEntry.EntityData.Leafs.Append("cpmProcExtHCMemAllocatedRev", types.YLeaf{"CpmProcExtHCMemAllocatedRev", cpmProcessExtRevEntry.CpmProcExtHCMemAllocatedRev})
    cpmProcessExtRevEntry.EntityData.Leafs.Append("cpmProcExtMemFreedRevOvrflw", types.YLeaf{"CpmProcExtMemFreedRevOvrflw", cpmProcessExtRevEntry.CpmProcExtMemFreedRevOvrflw})
    cpmProcessExtRevEntry.EntityData.Leafs.Append("cpmProcExtHCMemFreedRev", types.YLeaf{"CpmProcExtHCMemFreedRev", cpmProcessExtRevEntry.CpmProcExtHCMemFreedRev})
    cpmProcessExtRevEntry.EntityData.Leafs.Append("cpmProcessTextSegmentSizeOvrflw", types.YLeaf{"CpmProcessTextSegmentSizeOvrflw", cpmProcessExtRevEntry.CpmProcessTextSegmentSizeOvrflw})
    cpmProcessExtRevEntry.EntityData.Leafs.Append("cpmProcessHCTextSegmentSize", types.YLeaf{"CpmProcessHCTextSegmentSize", cpmProcessExtRevEntry.CpmProcessHCTextSegmentSize})
    cpmProcessExtRevEntry.EntityData.Leafs.Append("cpmProcessDataSegmentSizeOvrflw", types.YLeaf{"CpmProcessDataSegmentSizeOvrflw", cpmProcessExtRevEntry.CpmProcessDataSegmentSizeOvrflw})
    cpmProcessExtRevEntry.EntityData.Leafs.Append("cpmProcessHCDataSegmentSize", types.YLeaf{"CpmProcessHCDataSegmentSize", cpmProcessExtRevEntry.CpmProcessHCDataSegmentSize})
    cpmProcessExtRevEntry.EntityData.Leafs.Append("cpmProcessStackSizeOvrflw", types.YLeaf{"CpmProcessStackSizeOvrflw", cpmProcessExtRevEntry.CpmProcessStackSizeOvrflw})
    cpmProcessExtRevEntry.EntityData.Leafs.Append("cpmProcessHCStackSize", types.YLeaf{"CpmProcessHCStackSize", cpmProcessExtRevEntry.CpmProcessHCStackSize})
    cpmProcessExtRevEntry.EntityData.Leafs.Append("cpmProcessDynamicMemorySizeOvrflw", types.YLeaf{"CpmProcessDynamicMemorySizeOvrflw", cpmProcessExtRevEntry.CpmProcessDynamicMemorySizeOvrflw})
    cpmProcessExtRevEntry.EntityData.Leafs.Append("cpmProcessHCDynamicMemorySize", types.YLeaf{"CpmProcessHCDynamicMemorySize", cpmProcessExtRevEntry.CpmProcessHCDynamicMemorySize})

    cpmProcessExtRevEntry.EntityData.YListKeys = []string {"CpmCPUTotalIndex", "CpmProcessPID"}

    return &(cpmProcessExtRevEntry.EntityData)
}

// CISCOPROCESSMIB_CpmProcessExtRevTable_CpmProcessExtRevEntry_CpmProcExtPriorityRev represents cpmProcExtPriority.
type CISCOPROCESSMIB_CpmProcessExtRevTable_CpmProcessExtRevEntry_CpmProcExtPriorityRev string

const (
    CISCOPROCESSMIB_CpmProcessExtRevTable_CpmProcessExtRevEntry_CpmProcExtPriorityRev_critical CISCOPROCESSMIB_CpmProcessExtRevTable_CpmProcessExtRevEntry_CpmProcExtPriorityRev = "critical"

    CISCOPROCESSMIB_CpmProcessExtRevTable_CpmProcessExtRevEntry_CpmProcExtPriorityRev_high CISCOPROCESSMIB_CpmProcessExtRevTable_CpmProcessExtRevEntry_CpmProcExtPriorityRev = "high"

    CISCOPROCESSMIB_CpmProcessExtRevTable_CpmProcessExtRevEntry_CpmProcExtPriorityRev_normal CISCOPROCESSMIB_CpmProcessExtRevTable_CpmProcessExtRevEntry_CpmProcExtPriorityRev = "normal"

    CISCOPROCESSMIB_CpmProcessExtRevTable_CpmProcessExtRevEntry_CpmProcExtPriorityRev_low CISCOPROCESSMIB_CpmProcessExtRevTable_CpmProcessExtRevEntry_CpmProcExtPriorityRev = "low"

    CISCOPROCESSMIB_CpmProcessExtRevTable_CpmProcessExtRevEntry_CpmProcExtPriorityRev_notAssigned CISCOPROCESSMIB_CpmProcessExtRevTable_CpmProcessExtRevEntry_CpmProcExtPriorityRev = "notAssigned"
)

// CISCOPROCESSMIB_CpmProcessExtRevTable_CpmProcessExtRevEntry_CpmProcessMemoryCore represents                 to facilitate faster process restart.
type CISCOPROCESSMIB_CpmProcessExtRevTable_CpmProcessExtRevEntry_CpmProcessMemoryCore string

const (
    CISCOPROCESSMIB_CpmProcessExtRevTable_CpmProcessExtRevEntry_CpmProcessMemoryCore_none CISCOPROCESSMIB_CpmProcessExtRevTable_CpmProcessExtRevEntry_CpmProcessMemoryCore = "none"

    CISCOPROCESSMIB_CpmProcessExtRevTable_CpmProcessExtRevEntry_CpmProcessMemoryCore_other CISCOPROCESSMIB_CpmProcessExtRevTable_CpmProcessExtRevEntry_CpmProcessMemoryCore = "other"

    CISCOPROCESSMIB_CpmProcessExtRevTable_CpmProcessExtRevEntry_CpmProcessMemoryCore_mainmem CISCOPROCESSMIB_CpmProcessExtRevTable_CpmProcessExtRevEntry_CpmProcessMemoryCore = "mainmem"

    CISCOPROCESSMIB_CpmProcessExtRevTable_CpmProcessExtRevEntry_CpmProcessMemoryCore_mainmemSharedmem CISCOPROCESSMIB_CpmProcessExtRevTable_CpmProcessExtRevEntry_CpmProcessMemoryCore = "mainmemSharedmem"

    CISCOPROCESSMIB_CpmProcessExtRevTable_CpmProcessExtRevEntry_CpmProcessMemoryCore_mainmemText CISCOPROCESSMIB_CpmProcessExtRevTable_CpmProcessExtRevEntry_CpmProcessMemoryCore = "mainmemText"

    CISCOPROCESSMIB_CpmProcessExtRevTable_CpmProcessExtRevEntry_CpmProcessMemoryCore_mainmemTextSharedmem CISCOPROCESSMIB_CpmProcessExtRevTable_CpmProcessExtRevEntry_CpmProcessMemoryCore = "mainmemTextSharedmem"

    CISCOPROCESSMIB_CpmProcessExtRevTable_CpmProcessExtRevEntry_CpmProcessMemoryCore_sharedmem CISCOPROCESSMIB_CpmProcessExtRevTable_CpmProcessExtRevEntry_CpmProcessMemoryCore = "sharedmem"

    CISCOPROCESSMIB_CpmProcessExtRevTable_CpmProcessExtRevEntry_CpmProcessMemoryCore_sparse CISCOPROCESSMIB_CpmProcessExtRevTable_CpmProcessExtRevEntry_CpmProcessMemoryCore = "sparse"

    CISCOPROCESSMIB_CpmProcessExtRevTable_CpmProcessExtRevEntry_CpmProcessMemoryCore_off CISCOPROCESSMIB_CpmProcessExtRevTable_CpmProcessExtRevEntry_CpmProcessMemoryCore = "off"
)

// CISCOPROCESSMIB_CpmProcessExtRevTable_CpmProcessExtRevEntry_CpmProcessType represents This indicates the kind of process in context.
type CISCOPROCESSMIB_CpmProcessExtRevTable_CpmProcessExtRevEntry_CpmProcessType string

const (
    CISCOPROCESSMIB_CpmProcessExtRevTable_CpmProcessExtRevEntry_CpmProcessType_none CISCOPROCESSMIB_CpmProcessExtRevTable_CpmProcessExtRevEntry_CpmProcessType = "none"

    CISCOPROCESSMIB_CpmProcessExtRevTable_CpmProcessExtRevEntry_CpmProcessType_other CISCOPROCESSMIB_CpmProcessExtRevTable_CpmProcessExtRevEntry_CpmProcessType = "other"

    CISCOPROCESSMIB_CpmProcessExtRevTable_CpmProcessExtRevEntry_CpmProcessType_posix CISCOPROCESSMIB_CpmProcessExtRevTable_CpmProcessExtRevEntry_CpmProcessType = "posix"

    CISCOPROCESSMIB_CpmProcessExtRevTable_CpmProcessExtRevEntry_CpmProcessType_ios CISCOPROCESSMIB_CpmProcessExtRevTable_CpmProcessExtRevEntry_CpmProcessType = "ios"
)

// CISCOPROCESSMIB_CpmCPUThresholdTable
// This table contains the information about the
// thresholding values for CPU , configured by the user.
type CISCOPROCESSMIB_CpmCPUThresholdTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry containing information about CPU thresholding parameters.
    // cpmCPUTotalIndex identifies the CPU (or group of CPUs) for which this
    // configuration applies. The type is slice of
    // CISCOPROCESSMIB_CpmCPUThresholdTable_CpmCPUThresholdEntry.
    CpmCPUThresholdEntry []*CISCOPROCESSMIB_CpmCPUThresholdTable_CpmCPUThresholdEntry
}

func (cpmCPUThresholdTable *CISCOPROCESSMIB_CpmCPUThresholdTable) GetEntityData() *types.CommonEntityData {
    cpmCPUThresholdTable.EntityData.YFilter = cpmCPUThresholdTable.YFilter
    cpmCPUThresholdTable.EntityData.YangName = "cpmCPUThresholdTable"
    cpmCPUThresholdTable.EntityData.BundleName = "cisco_ios_xe"
    cpmCPUThresholdTable.EntityData.ParentYangName = "CISCO-PROCESS-MIB"
    cpmCPUThresholdTable.EntityData.SegmentPath = "cpmCPUThresholdTable"
    cpmCPUThresholdTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpmCPUThresholdTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpmCPUThresholdTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpmCPUThresholdTable.EntityData.Children = types.NewOrderedMap()
    cpmCPUThresholdTable.EntityData.Children.Append("cpmCPUThresholdEntry", types.YChild{"CpmCPUThresholdEntry", nil})
    for i := range cpmCPUThresholdTable.CpmCPUThresholdEntry {
        cpmCPUThresholdTable.EntityData.Children.Append(types.GetSegmentPath(cpmCPUThresholdTable.CpmCPUThresholdEntry[i]), types.YChild{"CpmCPUThresholdEntry", cpmCPUThresholdTable.CpmCPUThresholdEntry[i]})
    }
    cpmCPUThresholdTable.EntityData.Leafs = types.NewOrderedMap()

    cpmCPUThresholdTable.EntityData.YListKeys = []string {}

    return &(cpmCPUThresholdTable.EntityData)
}

// CISCOPROCESSMIB_CpmCPUThresholdTable_CpmCPUThresholdEntry
// An entry containing information about
// CPU thresholding parameters. cpmCPUTotalIndex
// identifies the CPU (or group of CPUs) for which this
// configuration applies.
type CISCOPROCESSMIB_CpmCPUThresholdTable_CpmCPUThresholdEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // cisco_process_mib.CISCOPROCESSMIB_CpmCPUTotalTable_CpmCPUTotalEntry_CpmCPUTotalIndex
    CpmCPUTotalIndex interface{}

    // This attribute is a key. Value of this object indicates the type of
    // utilization, which is monitored. The total(1) indicates the total CPU
    // utilization, interrupt(2) indicates the the CPU utilization in interrupt
    // context and process(3) indicates the CPU utilization in the process level
    // execution context. The type is CpmCPUThresholdClass.
    CpmCPUThresholdClass interface{}

    // The percentage rising threshold value configured by the user. The value
    // indicates,  if the percentage CPU utilization is equal to or above this
    // value for cpmCPURisingThresholdPeriod duration  then send a
    // cpmCPURisingThreshold notification to the NMS. The type is interface{} with
    // range: 1..100.
    CpmCPURisingThresholdValue interface{}

    // This is an observation interval. The value of this object indicates that 
    // the CPU utilization should be above cpmCPURisingThresholdValue for this
    // duration to send a  cpmCPURisingThreshold notification to the NMS. The type
    // is interface{} with range: 5..4294967295. Units are seconds.
    CpmCPURisingThresholdPeriod interface{}

    // The percentage falling threshold value configured by the user. The value
    // indicates, if the percentage  CPU utilization is equal to or below this
    // value for  cpmCPUFallingThresholdPeriod duration then send a
    // cpmCPUFallingThreshold notification  to the NMS. The type is interface{}
    // with range: 1..100.
    CpmCPUFallingThresholdValue interface{}

    // This is an observation interval. The value of this object indicates that
    // CPU utilization should be below cpmCPUFallingThresholdValue for this
    // duration to send a  cpmCPURisingThreshold notification to the NMS. The type
    // is interface{} with range: 5..4294967295. Units are seconds.
    CpmCPUFallingThresholdPeriod interface{}

    // The status of this table entry. The type is RowStatus.
    CpmCPUThresholdEntryStatus interface{}
}

func (cpmCPUThresholdEntry *CISCOPROCESSMIB_CpmCPUThresholdTable_CpmCPUThresholdEntry) GetEntityData() *types.CommonEntityData {
    cpmCPUThresholdEntry.EntityData.YFilter = cpmCPUThresholdEntry.YFilter
    cpmCPUThresholdEntry.EntityData.YangName = "cpmCPUThresholdEntry"
    cpmCPUThresholdEntry.EntityData.BundleName = "cisco_ios_xe"
    cpmCPUThresholdEntry.EntityData.ParentYangName = "cpmCPUThresholdTable"
    cpmCPUThresholdEntry.EntityData.SegmentPath = "cpmCPUThresholdEntry" + types.AddKeyToken(cpmCPUThresholdEntry.CpmCPUTotalIndex, "cpmCPUTotalIndex") + types.AddKeyToken(cpmCPUThresholdEntry.CpmCPUThresholdClass, "cpmCPUThresholdClass")
    cpmCPUThresholdEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpmCPUThresholdEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpmCPUThresholdEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpmCPUThresholdEntry.EntityData.Children = types.NewOrderedMap()
    cpmCPUThresholdEntry.EntityData.Leafs = types.NewOrderedMap()
    cpmCPUThresholdEntry.EntityData.Leafs.Append("cpmCPUTotalIndex", types.YLeaf{"CpmCPUTotalIndex", cpmCPUThresholdEntry.CpmCPUTotalIndex})
    cpmCPUThresholdEntry.EntityData.Leafs.Append("cpmCPUThresholdClass", types.YLeaf{"CpmCPUThresholdClass", cpmCPUThresholdEntry.CpmCPUThresholdClass})
    cpmCPUThresholdEntry.EntityData.Leafs.Append("cpmCPURisingThresholdValue", types.YLeaf{"CpmCPURisingThresholdValue", cpmCPUThresholdEntry.CpmCPURisingThresholdValue})
    cpmCPUThresholdEntry.EntityData.Leafs.Append("cpmCPURisingThresholdPeriod", types.YLeaf{"CpmCPURisingThresholdPeriod", cpmCPUThresholdEntry.CpmCPURisingThresholdPeriod})
    cpmCPUThresholdEntry.EntityData.Leafs.Append("cpmCPUFallingThresholdValue", types.YLeaf{"CpmCPUFallingThresholdValue", cpmCPUThresholdEntry.CpmCPUFallingThresholdValue})
    cpmCPUThresholdEntry.EntityData.Leafs.Append("cpmCPUFallingThresholdPeriod", types.YLeaf{"CpmCPUFallingThresholdPeriod", cpmCPUThresholdEntry.CpmCPUFallingThresholdPeriod})
    cpmCPUThresholdEntry.EntityData.Leafs.Append("cpmCPUThresholdEntryStatus", types.YLeaf{"CpmCPUThresholdEntryStatus", cpmCPUThresholdEntry.CpmCPUThresholdEntryStatus})

    cpmCPUThresholdEntry.EntityData.YListKeys = []string {"CpmCPUTotalIndex", "CpmCPUThresholdClass"}

    return &(cpmCPUThresholdEntry.EntityData)
}

// CISCOPROCESSMIB_CpmCPUThresholdTable_CpmCPUThresholdEntry_CpmCPUThresholdClass represents execution context.
type CISCOPROCESSMIB_CpmCPUThresholdTable_CpmCPUThresholdEntry_CpmCPUThresholdClass string

const (
    CISCOPROCESSMIB_CpmCPUThresholdTable_CpmCPUThresholdEntry_CpmCPUThresholdClass_total CISCOPROCESSMIB_CpmCPUThresholdTable_CpmCPUThresholdEntry_CpmCPUThresholdClass = "total"

    CISCOPROCESSMIB_CpmCPUThresholdTable_CpmCPUThresholdEntry_CpmCPUThresholdClass_interrupt CISCOPROCESSMIB_CpmCPUThresholdTable_CpmCPUThresholdEntry_CpmCPUThresholdClass = "interrupt"

    CISCOPROCESSMIB_CpmCPUThresholdTable_CpmCPUThresholdEntry_CpmCPUThresholdClass_process CISCOPROCESSMIB_CpmCPUThresholdTable_CpmCPUThresholdEntry_CpmCPUThresholdClass = "process"
)

// CISCOPROCESSMIB_CpmCPUHistoryTable
// A list of CPU utilization history entries.
type CISCOPROCESSMIB_CpmCPUHistoryTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A historical sample of CPU utilization statistics. cpmCPUTotalIndex
    // identifies the CPU (or group of CPUs) for which this history is collected. 
    // When the cpmCPUHistorySize is reached the least recent entry is lost. The
    // type is slice of CISCOPROCESSMIB_CpmCPUHistoryTable_CpmCPUHistoryEntry.
    CpmCPUHistoryEntry []*CISCOPROCESSMIB_CpmCPUHistoryTable_CpmCPUHistoryEntry
}

func (cpmCPUHistoryTable *CISCOPROCESSMIB_CpmCPUHistoryTable) GetEntityData() *types.CommonEntityData {
    cpmCPUHistoryTable.EntityData.YFilter = cpmCPUHistoryTable.YFilter
    cpmCPUHistoryTable.EntityData.YangName = "cpmCPUHistoryTable"
    cpmCPUHistoryTable.EntityData.BundleName = "cisco_ios_xe"
    cpmCPUHistoryTable.EntityData.ParentYangName = "CISCO-PROCESS-MIB"
    cpmCPUHistoryTable.EntityData.SegmentPath = "cpmCPUHistoryTable"
    cpmCPUHistoryTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpmCPUHistoryTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpmCPUHistoryTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpmCPUHistoryTable.EntityData.Children = types.NewOrderedMap()
    cpmCPUHistoryTable.EntityData.Children.Append("cpmCPUHistoryEntry", types.YChild{"CpmCPUHistoryEntry", nil})
    for i := range cpmCPUHistoryTable.CpmCPUHistoryEntry {
        cpmCPUHistoryTable.EntityData.Children.Append(types.GetSegmentPath(cpmCPUHistoryTable.CpmCPUHistoryEntry[i]), types.YChild{"CpmCPUHistoryEntry", cpmCPUHistoryTable.CpmCPUHistoryEntry[i]})
    }
    cpmCPUHistoryTable.EntityData.Leafs = types.NewOrderedMap()

    cpmCPUHistoryTable.EntityData.YListKeys = []string {}

    return &(cpmCPUHistoryTable.EntityData)
}

// CISCOPROCESSMIB_CpmCPUHistoryTable_CpmCPUHistoryEntry
// A historical sample of CPU utilization statistics.
// cpmCPUTotalIndex identifies the CPU (or group of CPUs)
// for which this history is collected. 
// When the cpmCPUHistorySize is
// reached the least recent entry is lost.
type CISCOPROCESSMIB_CpmCPUHistoryTable_CpmCPUHistoryEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // cisco_process_mib.CISCOPROCESSMIB_CpmCPUTotalTable_CpmCPUTotalEntry_CpmCPUTotalIndex
    CpmCPUTotalIndex interface{}

    // This attribute is a key. All the entries which are created at the same time
    // will have same value for this object. When the configured threshold for
    // being a part of History table is reached then the qualified processes
    // become the part of history table. The entries which became the  part of
    // history table at one instant will have the same value for this object. When
    // this object reaches the max index value then it will wrap around. The type
    // is interface{} with range: 0..4294967295.
    CpmCPUHistoryReportId interface{}

    // The number of process entries in a report. This object gives information
    // about how many processes  became a part of history table at one instant.
    // The type is interface{} with range: 0..4294967295.
    CpmCPUHistoryReportSize interface{}

    // Total percentage of CPU utilization at cpmCPUHistoryCreated. The type is
    // interface{} with range: 0..100. Units are percent.
    CpmCPUHistoryTotalUtil interface{}

    // Percentage of CPU utilization in the interrupt context at
    // cpmCPUHistoryCreated. The type is interface{} with range: 0..100. Units are
    // percent.
    CpmCPUHistoryInterruptUtil interface{}

    // Time stamp with respect to sysUpTime indicating the time at which this
    // report is created. The type is interface{} with range: 0..4294967295.
    CpmCPUHistoryCreatedTime interface{}
}

func (cpmCPUHistoryEntry *CISCOPROCESSMIB_CpmCPUHistoryTable_CpmCPUHistoryEntry) GetEntityData() *types.CommonEntityData {
    cpmCPUHistoryEntry.EntityData.YFilter = cpmCPUHistoryEntry.YFilter
    cpmCPUHistoryEntry.EntityData.YangName = "cpmCPUHistoryEntry"
    cpmCPUHistoryEntry.EntityData.BundleName = "cisco_ios_xe"
    cpmCPUHistoryEntry.EntityData.ParentYangName = "cpmCPUHistoryTable"
    cpmCPUHistoryEntry.EntityData.SegmentPath = "cpmCPUHistoryEntry" + types.AddKeyToken(cpmCPUHistoryEntry.CpmCPUTotalIndex, "cpmCPUTotalIndex") + types.AddKeyToken(cpmCPUHistoryEntry.CpmCPUHistoryReportId, "cpmCPUHistoryReportId")
    cpmCPUHistoryEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpmCPUHistoryEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpmCPUHistoryEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpmCPUHistoryEntry.EntityData.Children = types.NewOrderedMap()
    cpmCPUHistoryEntry.EntityData.Leafs = types.NewOrderedMap()
    cpmCPUHistoryEntry.EntityData.Leafs.Append("cpmCPUTotalIndex", types.YLeaf{"CpmCPUTotalIndex", cpmCPUHistoryEntry.CpmCPUTotalIndex})
    cpmCPUHistoryEntry.EntityData.Leafs.Append("cpmCPUHistoryReportId", types.YLeaf{"CpmCPUHistoryReportId", cpmCPUHistoryEntry.CpmCPUHistoryReportId})
    cpmCPUHistoryEntry.EntityData.Leafs.Append("cpmCPUHistoryReportSize", types.YLeaf{"CpmCPUHistoryReportSize", cpmCPUHistoryEntry.CpmCPUHistoryReportSize})
    cpmCPUHistoryEntry.EntityData.Leafs.Append("cpmCPUHistoryTotalUtil", types.YLeaf{"CpmCPUHistoryTotalUtil", cpmCPUHistoryEntry.CpmCPUHistoryTotalUtil})
    cpmCPUHistoryEntry.EntityData.Leafs.Append("cpmCPUHistoryInterruptUtil", types.YLeaf{"CpmCPUHistoryInterruptUtil", cpmCPUHistoryEntry.CpmCPUHistoryInterruptUtil})
    cpmCPUHistoryEntry.EntityData.Leafs.Append("cpmCPUHistoryCreatedTime", types.YLeaf{"CpmCPUHistoryCreatedTime", cpmCPUHistoryEntry.CpmCPUHistoryCreatedTime})

    cpmCPUHistoryEntry.EntityData.YListKeys = []string {"CpmCPUTotalIndex", "CpmCPUHistoryReportId"}

    return &(cpmCPUHistoryEntry.EntityData)
}

// CISCOPROCESSMIB_CpmCPUProcessHistoryTable
// A list of process history entries. This table contains
// CPU utilization of processes which crossed the 
// cpmCPUHistoryThreshold.
type CISCOPROCESSMIB_CpmCPUProcessHistoryTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A historical sample of process utilization statistics. The entries in this
    // table will have corresponding entires in the cpmCPUHistoryTable. The
    // entries in this table get deleted when the entry associated with this entry
    // in the cpmCPUHistoryTable  gets deleted. The type is slice of
    // CISCOPROCESSMIB_CpmCPUProcessHistoryTable_CpmCPUProcessHistoryEntry.
    CpmCPUProcessHistoryEntry []*CISCOPROCESSMIB_CpmCPUProcessHistoryTable_CpmCPUProcessHistoryEntry
}

func (cpmCPUProcessHistoryTable *CISCOPROCESSMIB_CpmCPUProcessHistoryTable) GetEntityData() *types.CommonEntityData {
    cpmCPUProcessHistoryTable.EntityData.YFilter = cpmCPUProcessHistoryTable.YFilter
    cpmCPUProcessHistoryTable.EntityData.YangName = "cpmCPUProcessHistoryTable"
    cpmCPUProcessHistoryTable.EntityData.BundleName = "cisco_ios_xe"
    cpmCPUProcessHistoryTable.EntityData.ParentYangName = "CISCO-PROCESS-MIB"
    cpmCPUProcessHistoryTable.EntityData.SegmentPath = "cpmCPUProcessHistoryTable"
    cpmCPUProcessHistoryTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpmCPUProcessHistoryTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpmCPUProcessHistoryTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpmCPUProcessHistoryTable.EntityData.Children = types.NewOrderedMap()
    cpmCPUProcessHistoryTable.EntityData.Children.Append("cpmCPUProcessHistoryEntry", types.YChild{"CpmCPUProcessHistoryEntry", nil})
    for i := range cpmCPUProcessHistoryTable.CpmCPUProcessHistoryEntry {
        cpmCPUProcessHistoryTable.EntityData.Children.Append(types.GetSegmentPath(cpmCPUProcessHistoryTable.CpmCPUProcessHistoryEntry[i]), types.YChild{"CpmCPUProcessHistoryEntry", cpmCPUProcessHistoryTable.CpmCPUProcessHistoryEntry[i]})
    }
    cpmCPUProcessHistoryTable.EntityData.Leafs = types.NewOrderedMap()

    cpmCPUProcessHistoryTable.EntityData.YListKeys = []string {}

    return &(cpmCPUProcessHistoryTable.EntityData)
}

// CISCOPROCESSMIB_CpmCPUProcessHistoryTable_CpmCPUProcessHistoryEntry
// A historical sample of process utilization
// statistics. The entries in this table will have
// corresponding entires in the cpmCPUHistoryTable.
// The entries in this table get deleted when the entry
// associated with this entry in the cpmCPUHistoryTable 
// gets deleted.
type CISCOPROCESSMIB_CpmCPUProcessHistoryTable_CpmCPUProcessHistoryEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // cisco_process_mib.CISCOPROCESSMIB_CpmCPUTotalTable_CpmCPUTotalEntry_CpmCPUTotalIndex
    CpmCPUTotalIndex interface{}

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // cisco_process_mib.CISCOPROCESSMIB_CpmCPUHistoryTable_CpmCPUHistoryEntry_CpmCPUHistoryReportId
    CpmCPUHistoryReportId interface{}

    // This attribute is a key. An index that uniquely identifies an entry in the
    // cmpCPUProcessHistory table among those in the  same report. This index is
    // between 1 to N,  where N is the cpmCPUHistoryReportSize. The type is
    // interface{} with range: 1..4294967295.
    CpmCPUProcessHistoryIndex interface{}

    // The process Id associated with this entry. The type is interface{} with
    // range: 1..2147483647.
    CpmCPUHistoryProcId interface{}

    // The process name associated with this entry. The type is string.
    CpmCPUHistoryProcName interface{}

    // The time when the process was created. The process ID and the time when the
    // process was created, uniquely  identifies a process. The type is
    // interface{} with range: 0..4294967295.
    CpmCPUHistoryProcCreated interface{}

    // The percentage CPU utilization of a process at cpmCPUHistoryCreatedTime.
    // The type is interface{} with range: 0..100. Units are percent.
    CpmCPUHistoryProcUtil interface{}
}

func (cpmCPUProcessHistoryEntry *CISCOPROCESSMIB_CpmCPUProcessHistoryTable_CpmCPUProcessHistoryEntry) GetEntityData() *types.CommonEntityData {
    cpmCPUProcessHistoryEntry.EntityData.YFilter = cpmCPUProcessHistoryEntry.YFilter
    cpmCPUProcessHistoryEntry.EntityData.YangName = "cpmCPUProcessHistoryEntry"
    cpmCPUProcessHistoryEntry.EntityData.BundleName = "cisco_ios_xe"
    cpmCPUProcessHistoryEntry.EntityData.ParentYangName = "cpmCPUProcessHistoryTable"
    cpmCPUProcessHistoryEntry.EntityData.SegmentPath = "cpmCPUProcessHistoryEntry" + types.AddKeyToken(cpmCPUProcessHistoryEntry.CpmCPUTotalIndex, "cpmCPUTotalIndex") + types.AddKeyToken(cpmCPUProcessHistoryEntry.CpmCPUHistoryReportId, "cpmCPUHistoryReportId") + types.AddKeyToken(cpmCPUProcessHistoryEntry.CpmCPUProcessHistoryIndex, "cpmCPUProcessHistoryIndex")
    cpmCPUProcessHistoryEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpmCPUProcessHistoryEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpmCPUProcessHistoryEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpmCPUProcessHistoryEntry.EntityData.Children = types.NewOrderedMap()
    cpmCPUProcessHistoryEntry.EntityData.Leafs = types.NewOrderedMap()
    cpmCPUProcessHistoryEntry.EntityData.Leafs.Append("cpmCPUTotalIndex", types.YLeaf{"CpmCPUTotalIndex", cpmCPUProcessHistoryEntry.CpmCPUTotalIndex})
    cpmCPUProcessHistoryEntry.EntityData.Leafs.Append("cpmCPUHistoryReportId", types.YLeaf{"CpmCPUHistoryReportId", cpmCPUProcessHistoryEntry.CpmCPUHistoryReportId})
    cpmCPUProcessHistoryEntry.EntityData.Leafs.Append("cpmCPUProcessHistoryIndex", types.YLeaf{"CpmCPUProcessHistoryIndex", cpmCPUProcessHistoryEntry.CpmCPUProcessHistoryIndex})
    cpmCPUProcessHistoryEntry.EntityData.Leafs.Append("cpmCPUHistoryProcId", types.YLeaf{"CpmCPUHistoryProcId", cpmCPUProcessHistoryEntry.CpmCPUHistoryProcId})
    cpmCPUProcessHistoryEntry.EntityData.Leafs.Append("cpmCPUHistoryProcName", types.YLeaf{"CpmCPUHistoryProcName", cpmCPUProcessHistoryEntry.CpmCPUHistoryProcName})
    cpmCPUProcessHistoryEntry.EntityData.Leafs.Append("cpmCPUHistoryProcCreated", types.YLeaf{"CpmCPUHistoryProcCreated", cpmCPUProcessHistoryEntry.CpmCPUHistoryProcCreated})
    cpmCPUProcessHistoryEntry.EntityData.Leafs.Append("cpmCPUHistoryProcUtil", types.YLeaf{"CpmCPUHistoryProcUtil", cpmCPUProcessHistoryEntry.CpmCPUHistoryProcUtil})

    cpmCPUProcessHistoryEntry.EntityData.YListKeys = []string {"CpmCPUTotalIndex", "CpmCPUHistoryReportId", "CpmCPUProcessHistoryIndex"}

    return &(cpmCPUProcessHistoryEntry.EntityData)
}

// CISCOPROCESSMIB_CpmThreadTable
// This table contains generic information about
// POSIX threads in the device.
type CISCOPROCESSMIB_CpmThreadTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry containing the general statistics of a POSIX thread. The type is
    // slice of CISCOPROCESSMIB_CpmThreadTable_CpmThreadEntry.
    CpmThreadEntry []*CISCOPROCESSMIB_CpmThreadTable_CpmThreadEntry
}

func (cpmThreadTable *CISCOPROCESSMIB_CpmThreadTable) GetEntityData() *types.CommonEntityData {
    cpmThreadTable.EntityData.YFilter = cpmThreadTable.YFilter
    cpmThreadTable.EntityData.YangName = "cpmThreadTable"
    cpmThreadTable.EntityData.BundleName = "cisco_ios_xe"
    cpmThreadTable.EntityData.ParentYangName = "CISCO-PROCESS-MIB"
    cpmThreadTable.EntityData.SegmentPath = "cpmThreadTable"
    cpmThreadTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpmThreadTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpmThreadTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpmThreadTable.EntityData.Children = types.NewOrderedMap()
    cpmThreadTable.EntityData.Children.Append("cpmThreadEntry", types.YChild{"CpmThreadEntry", nil})
    for i := range cpmThreadTable.CpmThreadEntry {
        cpmThreadTable.EntityData.Children.Append(types.GetSegmentPath(cpmThreadTable.CpmThreadEntry[i]), types.YChild{"CpmThreadEntry", cpmThreadTable.CpmThreadEntry[i]})
    }
    cpmThreadTable.EntityData.Leafs = types.NewOrderedMap()

    cpmThreadTable.EntityData.YListKeys = []string {}

    return &(cpmThreadTable.EntityData)
}

// CISCOPROCESSMIB_CpmThreadTable_CpmThreadEntry
// An entry containing the general statistics
// of a POSIX thread.
type CISCOPROCESSMIB_CpmThreadTable_CpmThreadEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // cisco_process_mib.CISCOPROCESSMIB_CpmCPUTotalTable_CpmCPUTotalEntry_CpmCPUTotalIndex
    CpmCPUTotalIndex interface{}

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // cisco_process_mib.CISCOPROCESSMIB_CpmProcessTable_CpmProcessEntry_CpmProcessPID
    CpmProcessPID interface{}

    // This attribute is a key. This object contains the thread ID. ThreadID is
    // Unique per process. The type is interface{} with range: 0..4294967295.
    CpmThreadID interface{}

    // This object represents the name of the thread. Thread names need not be
    // unique. Hence statistics  should be analyzed against thread ID. The type is
    // string.
    CpmThreadName interface{}

    // This object indicates the priority of a POSIX thread. The higher the
    // number, the higher the priority of the  thread over other threads. The type
    // is interface{} with range: 0..63.
    CpmThreadPriority interface{}

    // This object indicates the current state of a thread. Running state means
    // that the thread is actively  consumig CPU. All the other states are just
    // waiting  states. The valid states are: other         - Any other state
    // apart from the listed                  ones. dead          - Kernel is
    // waiting to release the                  thread's resources. running       -
    // Actively running on a CPU. ready         - Not running on a CPU, but is
    // ready to                  run (one or more higher or equal                 
    // priority threads are running). stopped       - Suspended (SIGSTOP signal).
    // send          - Waiting for a server to receive                  a message.
    // receive       - Waiting for a client to send a message. reply         -
    // Waiting for a server to reply to a                  message. stack        
    // - Waiting for more stack to be allocated. waitpage      - Waiting for
    // process manager to                  resolve a fault on a page. sigsuspend  
    // - Suspended for a signal. sigwaitinfo   - Waiting for a signal. nanosleep  
    // - Sleeping for a period of time. mutex         - Waiting to acquire a mutex
    // condvar       - Waiting for a condition variable to be                 
    // signalled. join          - Waiting for the completion of another           
    // thread. intr          - Waiting for an interrupt. sem           - Waiting
    // to acquire a semaphore. The type is CpmThreadState.
    CpmThreadState interface{}

    // This object identifies the process on which the current thread is blocked
    // on. This points to the  cpmProcessTable of the process on which the thread 
    // in context is blocked. This is valid only to threads which are either in
    // send/reply states. For the  rest of the threads it is returned as 0.0. The
    // type is string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    CpmThreadBlockingProcess interface{}

    // This object provides a general idea on how busy the thread in context
    // caused the processor to be. The type is interface{} with range:
    // 0..4294967295. Units are milliseconds.
    CpmThreadCpuUtilization interface{}

    // This object indicates the stack size allocated to the thread in context.
    // The type is interface{} with range: 0..4294967295. Units are bytes.
    CpmThreadStackSize interface{}

    // This object represents the upper 32-bit of cpmThreadStackSize. This object
    // needs to be supported only when the value of cpmThreadStackSize exceeds
    // 32-bit, otherwise this object value would be set to 0. The type is
    // interface{} with range: 0..4294967295. Units are bytes.
    CpmThreadStackSizeOvrflw interface{}

    // This object indicates the stack size allocated to the thread in context.
    // This object is a 64-bit version of cpmThreadStackSize. The type is
    // interface{} with range: 0..18446744073709551615. Units are bytes.
    CpmThreadHCStackSize interface{}
}

func (cpmThreadEntry *CISCOPROCESSMIB_CpmThreadTable_CpmThreadEntry) GetEntityData() *types.CommonEntityData {
    cpmThreadEntry.EntityData.YFilter = cpmThreadEntry.YFilter
    cpmThreadEntry.EntityData.YangName = "cpmThreadEntry"
    cpmThreadEntry.EntityData.BundleName = "cisco_ios_xe"
    cpmThreadEntry.EntityData.ParentYangName = "cpmThreadTable"
    cpmThreadEntry.EntityData.SegmentPath = "cpmThreadEntry" + types.AddKeyToken(cpmThreadEntry.CpmCPUTotalIndex, "cpmCPUTotalIndex") + types.AddKeyToken(cpmThreadEntry.CpmProcessPID, "cpmProcessPID") + types.AddKeyToken(cpmThreadEntry.CpmThreadID, "cpmThreadID")
    cpmThreadEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpmThreadEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpmThreadEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpmThreadEntry.EntityData.Children = types.NewOrderedMap()
    cpmThreadEntry.EntityData.Leafs = types.NewOrderedMap()
    cpmThreadEntry.EntityData.Leafs.Append("cpmCPUTotalIndex", types.YLeaf{"CpmCPUTotalIndex", cpmThreadEntry.CpmCPUTotalIndex})
    cpmThreadEntry.EntityData.Leafs.Append("cpmProcessPID", types.YLeaf{"CpmProcessPID", cpmThreadEntry.CpmProcessPID})
    cpmThreadEntry.EntityData.Leafs.Append("cpmThreadID", types.YLeaf{"CpmThreadID", cpmThreadEntry.CpmThreadID})
    cpmThreadEntry.EntityData.Leafs.Append("cpmThreadName", types.YLeaf{"CpmThreadName", cpmThreadEntry.CpmThreadName})
    cpmThreadEntry.EntityData.Leafs.Append("cpmThreadPriority", types.YLeaf{"CpmThreadPriority", cpmThreadEntry.CpmThreadPriority})
    cpmThreadEntry.EntityData.Leafs.Append("cpmThreadState", types.YLeaf{"CpmThreadState", cpmThreadEntry.CpmThreadState})
    cpmThreadEntry.EntityData.Leafs.Append("cpmThreadBlockingProcess", types.YLeaf{"CpmThreadBlockingProcess", cpmThreadEntry.CpmThreadBlockingProcess})
    cpmThreadEntry.EntityData.Leafs.Append("cpmThreadCpuUtilization", types.YLeaf{"CpmThreadCpuUtilization", cpmThreadEntry.CpmThreadCpuUtilization})
    cpmThreadEntry.EntityData.Leafs.Append("cpmThreadStackSize", types.YLeaf{"CpmThreadStackSize", cpmThreadEntry.CpmThreadStackSize})
    cpmThreadEntry.EntityData.Leafs.Append("cpmThreadStackSizeOvrflw", types.YLeaf{"CpmThreadStackSizeOvrflw", cpmThreadEntry.CpmThreadStackSizeOvrflw})
    cpmThreadEntry.EntityData.Leafs.Append("cpmThreadHCStackSize", types.YLeaf{"CpmThreadHCStackSize", cpmThreadEntry.CpmThreadHCStackSize})

    cpmThreadEntry.EntityData.YListKeys = []string {"CpmCPUTotalIndex", "CpmProcessPID", "CpmThreadID"}

    return &(cpmThreadEntry.EntityData)
}

// CISCOPROCESSMIB_CpmThreadTable_CpmThreadEntry_CpmThreadState represents sem           - Waiting to acquire a semaphore.
type CISCOPROCESSMIB_CpmThreadTable_CpmThreadEntry_CpmThreadState string

const (
    CISCOPROCESSMIB_CpmThreadTable_CpmThreadEntry_CpmThreadState_other CISCOPROCESSMIB_CpmThreadTable_CpmThreadEntry_CpmThreadState = "other"

    CISCOPROCESSMIB_CpmThreadTable_CpmThreadEntry_CpmThreadState_dead CISCOPROCESSMIB_CpmThreadTable_CpmThreadEntry_CpmThreadState = "dead"

    CISCOPROCESSMIB_CpmThreadTable_CpmThreadEntry_CpmThreadState_running CISCOPROCESSMIB_CpmThreadTable_CpmThreadEntry_CpmThreadState = "running"

    CISCOPROCESSMIB_CpmThreadTable_CpmThreadEntry_CpmThreadState_ready CISCOPROCESSMIB_CpmThreadTable_CpmThreadEntry_CpmThreadState = "ready"

    CISCOPROCESSMIB_CpmThreadTable_CpmThreadEntry_CpmThreadState_stopped CISCOPROCESSMIB_CpmThreadTable_CpmThreadEntry_CpmThreadState = "stopped"

    CISCOPROCESSMIB_CpmThreadTable_CpmThreadEntry_CpmThreadState_send CISCOPROCESSMIB_CpmThreadTable_CpmThreadEntry_CpmThreadState = "send"

    CISCOPROCESSMIB_CpmThreadTable_CpmThreadEntry_CpmThreadState_receive CISCOPROCESSMIB_CpmThreadTable_CpmThreadEntry_CpmThreadState = "receive"

    CISCOPROCESSMIB_CpmThreadTable_CpmThreadEntry_CpmThreadState_reply CISCOPROCESSMIB_CpmThreadTable_CpmThreadEntry_CpmThreadState = "reply"

    CISCOPROCESSMIB_CpmThreadTable_CpmThreadEntry_CpmThreadState_stack CISCOPROCESSMIB_CpmThreadTable_CpmThreadEntry_CpmThreadState = "stack"

    CISCOPROCESSMIB_CpmThreadTable_CpmThreadEntry_CpmThreadState_waitpage CISCOPROCESSMIB_CpmThreadTable_CpmThreadEntry_CpmThreadState = "waitpage"

    CISCOPROCESSMIB_CpmThreadTable_CpmThreadEntry_CpmThreadState_sigsuspend CISCOPROCESSMIB_CpmThreadTable_CpmThreadEntry_CpmThreadState = "sigsuspend"

    CISCOPROCESSMIB_CpmThreadTable_CpmThreadEntry_CpmThreadState_sigwaitinfo CISCOPROCESSMIB_CpmThreadTable_CpmThreadEntry_CpmThreadState = "sigwaitinfo"

    CISCOPROCESSMIB_CpmThreadTable_CpmThreadEntry_CpmThreadState_nanosleep CISCOPROCESSMIB_CpmThreadTable_CpmThreadEntry_CpmThreadState = "nanosleep"

    CISCOPROCESSMIB_CpmThreadTable_CpmThreadEntry_CpmThreadState_mutex CISCOPROCESSMIB_CpmThreadTable_CpmThreadEntry_CpmThreadState = "mutex"

    CISCOPROCESSMIB_CpmThreadTable_CpmThreadEntry_CpmThreadState_condvar CISCOPROCESSMIB_CpmThreadTable_CpmThreadEntry_CpmThreadState = "condvar"

    CISCOPROCESSMIB_CpmThreadTable_CpmThreadEntry_CpmThreadState_join CISCOPROCESSMIB_CpmThreadTable_CpmThreadEntry_CpmThreadState = "join"

    CISCOPROCESSMIB_CpmThreadTable_CpmThreadEntry_CpmThreadState_intr CISCOPROCESSMIB_CpmThreadTable_CpmThreadEntry_CpmThreadState = "intr"

    CISCOPROCESSMIB_CpmThreadTable_CpmThreadEntry_CpmThreadState_sem CISCOPROCESSMIB_CpmThreadTable_CpmThreadEntry_CpmThreadState = "sem"
)

// CISCOPROCESSMIB_CpmVirtualProcessTable
// This table contains information about virtual
// processes in a virtual machine.
type CISCOPROCESSMIB_CpmVirtualProcessTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry containing the general statistics of a virtual process in a
    // virtual machine. The type is slice of
    // CISCOPROCESSMIB_CpmVirtualProcessTable_CpmVirtualProcessEntry.
    CpmVirtualProcessEntry []*CISCOPROCESSMIB_CpmVirtualProcessTable_CpmVirtualProcessEntry
}

func (cpmVirtualProcessTable *CISCOPROCESSMIB_CpmVirtualProcessTable) GetEntityData() *types.CommonEntityData {
    cpmVirtualProcessTable.EntityData.YFilter = cpmVirtualProcessTable.YFilter
    cpmVirtualProcessTable.EntityData.YangName = "cpmVirtualProcessTable"
    cpmVirtualProcessTable.EntityData.BundleName = "cisco_ios_xe"
    cpmVirtualProcessTable.EntityData.ParentYangName = "CISCO-PROCESS-MIB"
    cpmVirtualProcessTable.EntityData.SegmentPath = "cpmVirtualProcessTable"
    cpmVirtualProcessTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpmVirtualProcessTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpmVirtualProcessTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpmVirtualProcessTable.EntityData.Children = types.NewOrderedMap()
    cpmVirtualProcessTable.EntityData.Children.Append("cpmVirtualProcessEntry", types.YChild{"CpmVirtualProcessEntry", nil})
    for i := range cpmVirtualProcessTable.CpmVirtualProcessEntry {
        cpmVirtualProcessTable.EntityData.Children.Append(types.GetSegmentPath(cpmVirtualProcessTable.CpmVirtualProcessEntry[i]), types.YChild{"CpmVirtualProcessEntry", cpmVirtualProcessTable.CpmVirtualProcessEntry[i]})
    }
    cpmVirtualProcessTable.EntityData.Leafs = types.NewOrderedMap()

    cpmVirtualProcessTable.EntityData.YListKeys = []string {}

    return &(cpmVirtualProcessTable.EntityData)
}

// CISCOPROCESSMIB_CpmVirtualProcessTable_CpmVirtualProcessEntry
// An entry containing the general statistics of a
// virtual process in a virtual machine.
type CISCOPROCESSMIB_CpmVirtualProcessTable_CpmVirtualProcessEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // cisco_process_mib.CISCOPROCESSMIB_CpmCPUTotalTable_CpmCPUTotalEntry_CpmCPUTotalIndex
    CpmCPUTotalIndex interface{}

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // cisco_process_mib.CISCOPROCESSMIB_CpmProcessTable_CpmProcessEntry_CpmProcessPID
    CpmProcessPID interface{}

    // This attribute is a key. This object indicates the process ID of a virtual
    // process. PID is unique only inside one address space. Virtual process PID
    // should be considered along with  Parent process cpmProcessPID. The type is
    // interface{} with range: 0..4294967295.
    CpmVirtualProcessID interface{}

    // This object indicates the name of a virtual process. If the name is longer
    // than 32 characters, it will be truncated to the first 31 characters, and a
    // `*' will be appended as the last character to imply this is a truncated
    // process name. The type is string with length: 1..32.
    CpmVirtualProcessName interface{}

    // This indicates an estimated CPU utilization by a virtual process over the
    // last 5 seconds. The type is interface{} with range: 0..100. Units are
    // percent.
    CpmVirtualProcessUtil5Sec interface{}

    // This indicates an estimated CPU utilization by a virtual process over the
    // last one minute. The type is interface{} with range: 0..100. Units are
    // percent.
    CpmVirtualProcessUtil1Min interface{}

    // This indicates an estimated CPU utilization by a virtual process over the
    // last 5 minutes. The type is interface{} with range: 0..100. Units are
    // percent.
    CpmVirtualProcessUtil5Min interface{}

    // This object indicates the memory allocated by the virtual process inside
    // the address space of a  process running on Native OS. The type is
    // interface{} with range: 0..4294967295. Units are bytes.
    CpmVirtualProcessMemAllocated interface{}

    // This object indicates the memory freed by the virtual process inside the
    // address space of a process running  on Native OS. The type is interface{}
    // with range: 0..4294967295. Units are bytes.
    CpmVirtualProcessMemFreed interface{}

    // The number of times a virtual process is invoked. The type is interface{}
    // with range: 0..4294967295.
    CpmVirtualProcessInvokeCount interface{}

    // The amount of CPU time a virtual process has used in microseconds. The type
    // is interface{} with range: 0..4294967295. Units are microseconds.
    CpmVirtualProcessRuntime interface{}

    // This object represents the upper 32-bit of cpmVirtualProcessMemAllocated.
    // This object  needs to be supported only when the value of
    // cpmVirtualProcessMemAllocated exceeds 32-bit, otherwise this object value
    // would be set to 0. The type is interface{} with range: 0..4294967295. Units
    // are bytes.
    CpmVirtualProcessMemAllocatedOvrflw interface{}

    // This object indicates the memory allocated by the virtual process inside
    // the address space of a process running on Native OS. This object is a
    // 64-bit version of cpmVirtualProcessMemAllocated. The type is interface{}
    // with range: 0..18446744073709551615. Units are bytes.
    CpmVirtualProcessHCMemAllocated interface{}

    // This object represents the upper 32-bit of cpmVirtualProcessMemFreed. This
    // object needs to be supported only when the value of 
    // cpmVirtualProcessMemFreed exceeds 32-bit,  otherwise this object value
    // would be set to 0. The type is interface{} with range: 0..4294967295. Units
    // are bytes.
    CpmVirtualProcessMemFreedOvrflw interface{}

    // This object indicates the memory freed by the virtual process inside the
    // address space of a process running on Native OS.This object is a 64-bit
    // version of cpmVirtualProcessMemAllocated. The type is interface{} with
    // range: 0..18446744073709551615. Units are bytes.
    CpmVirtualProcessHCMemFreed interface{}
}

func (cpmVirtualProcessEntry *CISCOPROCESSMIB_CpmVirtualProcessTable_CpmVirtualProcessEntry) GetEntityData() *types.CommonEntityData {
    cpmVirtualProcessEntry.EntityData.YFilter = cpmVirtualProcessEntry.YFilter
    cpmVirtualProcessEntry.EntityData.YangName = "cpmVirtualProcessEntry"
    cpmVirtualProcessEntry.EntityData.BundleName = "cisco_ios_xe"
    cpmVirtualProcessEntry.EntityData.ParentYangName = "cpmVirtualProcessTable"
    cpmVirtualProcessEntry.EntityData.SegmentPath = "cpmVirtualProcessEntry" + types.AddKeyToken(cpmVirtualProcessEntry.CpmCPUTotalIndex, "cpmCPUTotalIndex") + types.AddKeyToken(cpmVirtualProcessEntry.CpmProcessPID, "cpmProcessPID") + types.AddKeyToken(cpmVirtualProcessEntry.CpmVirtualProcessID, "cpmVirtualProcessID")
    cpmVirtualProcessEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpmVirtualProcessEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpmVirtualProcessEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpmVirtualProcessEntry.EntityData.Children = types.NewOrderedMap()
    cpmVirtualProcessEntry.EntityData.Leafs = types.NewOrderedMap()
    cpmVirtualProcessEntry.EntityData.Leafs.Append("cpmCPUTotalIndex", types.YLeaf{"CpmCPUTotalIndex", cpmVirtualProcessEntry.CpmCPUTotalIndex})
    cpmVirtualProcessEntry.EntityData.Leafs.Append("cpmProcessPID", types.YLeaf{"CpmProcessPID", cpmVirtualProcessEntry.CpmProcessPID})
    cpmVirtualProcessEntry.EntityData.Leafs.Append("cpmVirtualProcessID", types.YLeaf{"CpmVirtualProcessID", cpmVirtualProcessEntry.CpmVirtualProcessID})
    cpmVirtualProcessEntry.EntityData.Leafs.Append("cpmVirtualProcessName", types.YLeaf{"CpmVirtualProcessName", cpmVirtualProcessEntry.CpmVirtualProcessName})
    cpmVirtualProcessEntry.EntityData.Leafs.Append("cpmVirtualProcessUtil5Sec", types.YLeaf{"CpmVirtualProcessUtil5Sec", cpmVirtualProcessEntry.CpmVirtualProcessUtil5Sec})
    cpmVirtualProcessEntry.EntityData.Leafs.Append("cpmVirtualProcessUtil1Min", types.YLeaf{"CpmVirtualProcessUtil1Min", cpmVirtualProcessEntry.CpmVirtualProcessUtil1Min})
    cpmVirtualProcessEntry.EntityData.Leafs.Append("cpmVirtualProcessUtil5Min", types.YLeaf{"CpmVirtualProcessUtil5Min", cpmVirtualProcessEntry.CpmVirtualProcessUtil5Min})
    cpmVirtualProcessEntry.EntityData.Leafs.Append("cpmVirtualProcessMemAllocated", types.YLeaf{"CpmVirtualProcessMemAllocated", cpmVirtualProcessEntry.CpmVirtualProcessMemAllocated})
    cpmVirtualProcessEntry.EntityData.Leafs.Append("cpmVirtualProcessMemFreed", types.YLeaf{"CpmVirtualProcessMemFreed", cpmVirtualProcessEntry.CpmVirtualProcessMemFreed})
    cpmVirtualProcessEntry.EntityData.Leafs.Append("cpmVirtualProcessInvokeCount", types.YLeaf{"CpmVirtualProcessInvokeCount", cpmVirtualProcessEntry.CpmVirtualProcessInvokeCount})
    cpmVirtualProcessEntry.EntityData.Leafs.Append("cpmVirtualProcessRuntime", types.YLeaf{"CpmVirtualProcessRuntime", cpmVirtualProcessEntry.CpmVirtualProcessRuntime})
    cpmVirtualProcessEntry.EntityData.Leafs.Append("cpmVirtualProcessMemAllocatedOvrflw", types.YLeaf{"CpmVirtualProcessMemAllocatedOvrflw", cpmVirtualProcessEntry.CpmVirtualProcessMemAllocatedOvrflw})
    cpmVirtualProcessEntry.EntityData.Leafs.Append("cpmVirtualProcessHCMemAllocated", types.YLeaf{"CpmVirtualProcessHCMemAllocated", cpmVirtualProcessEntry.CpmVirtualProcessHCMemAllocated})
    cpmVirtualProcessEntry.EntityData.Leafs.Append("cpmVirtualProcessMemFreedOvrflw", types.YLeaf{"CpmVirtualProcessMemFreedOvrflw", cpmVirtualProcessEntry.CpmVirtualProcessMemFreedOvrflw})
    cpmVirtualProcessEntry.EntityData.Leafs.Append("cpmVirtualProcessHCMemFreed", types.YLeaf{"CpmVirtualProcessHCMemFreed", cpmVirtualProcessEntry.CpmVirtualProcessHCMemFreed})

    cpmVirtualProcessEntry.EntityData.YListKeys = []string {"CpmCPUTotalIndex", "CpmProcessPID", "CpmVirtualProcessID"}

    return &(cpmVirtualProcessEntry.EntityData)
}

