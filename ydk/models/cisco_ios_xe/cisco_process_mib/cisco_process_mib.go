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

    
    Cpmcpuhistory CISCOPROCESSMIB_Cpmcpuhistory

    // A table of overall CPU statistics.
    Cpmcputotaltable CISCOPROCESSMIB_Cpmcputotaltable

    // A table of per-Core statistics.
    Cpmcoretable CISCOPROCESSMIB_Cpmcoretable

    // A table of generic information on all active processes on this device.
    Cpmprocesstable CISCOPROCESSMIB_Cpmprocesstable

    // This table contains information that may or may not be available on all
    // cisco devices. It contains additional objects for the more general
    // cpmProcessTable. This object deprecates  cpmProcessExtTable.
    Cpmprocessextrevtable CISCOPROCESSMIB_Cpmprocessextrevtable

    // This table contains the information about the thresholding values for CPU ,
    // configured by the user.
    Cpmcputhresholdtable CISCOPROCESSMIB_Cpmcputhresholdtable

    // A list of CPU utilization history entries.
    Cpmcpuhistorytable CISCOPROCESSMIB_Cpmcpuhistorytable

    // A list of process history entries. This table contains CPU utilization of
    // processes which crossed the  cpmCPUHistoryThreshold.
    Cpmcpuprocesshistorytable CISCOPROCESSMIB_Cpmcpuprocesshistorytable

    // This table contains generic information about POSIX threads in the device.
    Cpmthreadtable CISCOPROCESSMIB_Cpmthreadtable

    // This table contains information about virtual processes in a virtual
    // machine.
    Cpmvirtualprocesstable CISCOPROCESSMIB_Cpmvirtualprocesstable
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

    cISCOPROCESSMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOPROCESSMIB.EntityData.Children["cpmCPUHistory"] = types.YChild{"Cpmcpuhistory", &cISCOPROCESSMIB.Cpmcpuhistory}
    cISCOPROCESSMIB.EntityData.Children["cpmCPUTotalTable"] = types.YChild{"Cpmcputotaltable", &cISCOPROCESSMIB.Cpmcputotaltable}
    cISCOPROCESSMIB.EntityData.Children["cpmCoreTable"] = types.YChild{"Cpmcoretable", &cISCOPROCESSMIB.Cpmcoretable}
    cISCOPROCESSMIB.EntityData.Children["cpmProcessTable"] = types.YChild{"Cpmprocesstable", &cISCOPROCESSMIB.Cpmprocesstable}
    cISCOPROCESSMIB.EntityData.Children["cpmProcessExtRevTable"] = types.YChild{"Cpmprocessextrevtable", &cISCOPROCESSMIB.Cpmprocessextrevtable}
    cISCOPROCESSMIB.EntityData.Children["cpmCPUThresholdTable"] = types.YChild{"Cpmcputhresholdtable", &cISCOPROCESSMIB.Cpmcputhresholdtable}
    cISCOPROCESSMIB.EntityData.Children["cpmCPUHistoryTable"] = types.YChild{"Cpmcpuhistorytable", &cISCOPROCESSMIB.Cpmcpuhistorytable}
    cISCOPROCESSMIB.EntityData.Children["cpmCPUProcessHistoryTable"] = types.YChild{"Cpmcpuprocesshistorytable", &cISCOPROCESSMIB.Cpmcpuprocesshistorytable}
    cISCOPROCESSMIB.EntityData.Children["cpmThreadTable"] = types.YChild{"Cpmthreadtable", &cISCOPROCESSMIB.Cpmthreadtable}
    cISCOPROCESSMIB.EntityData.Children["cpmVirtualProcessTable"] = types.YChild{"Cpmvirtualprocesstable", &cISCOPROCESSMIB.Cpmvirtualprocesstable}
    cISCOPROCESSMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOPROCESSMIB.EntityData)
}

// CISCOPROCESSMIB_Cpmcpuhistory
type CISCOPROCESSMIB_Cpmcpuhistory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The user  configured value of this object gives the minimum percent CPU
    // utilization of a process in the last cpmCPUMonInterval duration required to
    // be a  member of history table. When this object is changed the new value
    // will have effect in the next interval. The type is interface{} with range:
    // 1..100.
    Cpmcpuhistorythreshold interface{}

    // A value configured by the user which specifies the number of reports in the
    // history table.  A report contains set of processes which crossed the
    // cpmCPUHistoryThreshold  in the last cpmCPUMonInterval along with  the time
    // at which this report is created, total and interrupt CPU utilizations. 
    // When this object is changed the new value will have effect in the next
    // interval. The type is interface{} with range: 1..4294967295.
    Cpmcpuhistorysize interface{}
}

func (cpmcpuhistory *CISCOPROCESSMIB_Cpmcpuhistory) GetEntityData() *types.CommonEntityData {
    cpmcpuhistory.EntityData.YFilter = cpmcpuhistory.YFilter
    cpmcpuhistory.EntityData.YangName = "cpmCPUHistory"
    cpmcpuhistory.EntityData.BundleName = "cisco_ios_xe"
    cpmcpuhistory.EntityData.ParentYangName = "CISCO-PROCESS-MIB"
    cpmcpuhistory.EntityData.SegmentPath = "cpmCPUHistory"
    cpmcpuhistory.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpmcpuhistory.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpmcpuhistory.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpmcpuhistory.EntityData.Children = make(map[string]types.YChild)
    cpmcpuhistory.EntityData.Leafs = make(map[string]types.YLeaf)
    cpmcpuhistory.EntityData.Leafs["cpmCPUHistoryThreshold"] = types.YLeaf{"Cpmcpuhistorythreshold", cpmcpuhistory.Cpmcpuhistorythreshold}
    cpmcpuhistory.EntityData.Leafs["cpmCPUHistorySize"] = types.YLeaf{"Cpmcpuhistorysize", cpmcpuhistory.Cpmcpuhistorysize}
    return &(cpmcpuhistory.EntityData)
}

// CISCOPROCESSMIB_Cpmcputotaltable
// A table of overall CPU statistics.
type CISCOPROCESSMIB_Cpmcputotaltable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Overall information about the CPU load. Entries in this table come and go
    // as CPUs are added and removed from the system. The type is slice of
    // CISCOPROCESSMIB_Cpmcputotaltable_Cpmcputotalentry.
    Cpmcputotalentry []CISCOPROCESSMIB_Cpmcputotaltable_Cpmcputotalentry
}

func (cpmcputotaltable *CISCOPROCESSMIB_Cpmcputotaltable) GetEntityData() *types.CommonEntityData {
    cpmcputotaltable.EntityData.YFilter = cpmcputotaltable.YFilter
    cpmcputotaltable.EntityData.YangName = "cpmCPUTotalTable"
    cpmcputotaltable.EntityData.BundleName = "cisco_ios_xe"
    cpmcputotaltable.EntityData.ParentYangName = "CISCO-PROCESS-MIB"
    cpmcputotaltable.EntityData.SegmentPath = "cpmCPUTotalTable"
    cpmcputotaltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpmcputotaltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpmcputotaltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpmcputotaltable.EntityData.Children = make(map[string]types.YChild)
    cpmcputotaltable.EntityData.Children["cpmCPUTotalEntry"] = types.YChild{"Cpmcputotalentry", nil}
    for i := range cpmcputotaltable.Cpmcputotalentry {
        cpmcputotaltable.EntityData.Children[types.GetSegmentPath(&cpmcputotaltable.Cpmcputotalentry[i])] = types.YChild{"Cpmcputotalentry", &cpmcputotaltable.Cpmcputotalentry[i]}
    }
    cpmcputotaltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cpmcputotaltable.EntityData)
}

// CISCOPROCESSMIB_Cpmcputotaltable_Cpmcputotalentry
// Overall information about the CPU load. Entries in this
// table come and go as CPUs are added and removed from the
// system.
type CISCOPROCESSMIB_Cpmcputotaltable_Cpmcputotalentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An index that uniquely represents a CPU (or group
    // of CPUs) whose CPU load information is reported by a row in this table.
    // This index is assigned arbitrarily by the engine and is not saved over
    // reboots. The type is interface{} with range: 1..4294967295.
    Cpmcputotalindex interface{}

    // The entPhysicalIndex of the physical entity for which the CPU statistics in
    // this entry are maintained. The physical entity can be a CPU chip, a group
    // of CPUs, a CPU card etc. The exact type of this entity is described by its
    // entPhysicalVendorType value. If the CPU statistics in this entry correspond
    // to more than one physical entity (or to no physical entity), or if the
    // entPhysicalTable is not supported on the SNMP agent, the value of this
    // object must be zero. The type is interface{} with range: 0..2147483647.
    Cpmcputotalphysicalindex interface{}

    // The overall CPU busy percentage in the last 5 second period. This object
    // obsoletes the busyPer object from  the OLD-CISCO-SYSTEM-MIB. This object is
    // deprecated by cpmCPUTotal5secRev which has the changed range of value
    // (0..100). The type is interface{} with range: 0..100.
    Cpmcputotal5Sec interface{}

    // The overall CPU busy percentage in the last 1 minute period. This object
    // obsoletes the avgBusy1 object from  the OLD-CISCO-SYSTEM-MIB. This object
    // is deprecated by cpmCPUTotal1minRev which has the changed range of value
    // (0..100). The type is interface{} with range: 0..100.
    Cpmcputotal1Min interface{}

    // The overall CPU busy percentage in the last 5 minute period. This object
    // deprecates the avgBusy5 object from  the OLD-CISCO-SYSTEM-MIB. This object
    // is deprecated by cpmCPUTotal5minRev which has the changed range  of value
    // (0..100). The type is interface{} with range: 0..100.
    Cpmcputotal5Min interface{}

    // The overall CPU busy percentage in the last 5 second period. This object
    // deprecates the object cpmCPUTotal5sec  and increases the value range to
    // (0..100). This object is deprecated by cpmCPUTotalMonIntervalValue. The
    // type is interface{} with range: 0..100. Units are percent.
    Cpmcputotal5Secrev interface{}

    // The overall CPU busy percentage in the last 1 minute period. This object
    // deprecates the object cpmCPUTotal1min  and increases the value range to
    // (0..100). The type is interface{} with range: 0..100. Units are percent.
    Cpmcputotal1Minrev interface{}

    // The overall CPU busy percentage in the last 5 minute period. This object
    // deprecates the object cpmCPUTotal5min  and increases the value range to
    // (0..100). The type is interface{} with range: 0..100. Units are percent.
    Cpmcputotal5Minrev interface{}

    // CPU usage monitoring interval. The value of this object in seconds
    // indicates the how often the  CPU utilization is calculated and monitored.
    // The type is interface{} with range: 0..4294967295. Units are seconds.
    Cpmcpumoninterval interface{}

    // The overall CPU busy percentage in the last cpmCPUMonInterval period.  This
    // object deprecates the object cpmCPUTotal5secRev. The type is interface{}
    // with range: 0..100. Units are percent.
    Cpmcputotalmonintervalvalue interface{}

    // The overall CPU busy percentage in the interrupt context in the last
    // cpmCPUMonInterval period. The type is interface{} with range: 0..100. Units
    // are percent.
    Cpmcpuinterruptmonintervalvalue interface{}

    // The overall CPU wide system memory which is currently under use. The type
    // is interface{} with range: 0..4294967295. Units are kilo-bytes.
    Cpmcpumemoryused interface{}

    // The overall CPU wide system memory which is currently free. The type is
    // interface{} with range: 0..4294967295. Units are kilo-bytes.
    Cpmcpumemoryfree interface{}

    // The overall CPU wide system memory which is reserved for kernel usage. The
    // type is interface{} with range: 0..4294967295. Units are kilo-bytes.
    Cpmcpumemorykernelreserved interface{}

    // The lowest free memory that has been recorded since device has booted. The
    // type is interface{} with range: 0..4294967295. Units are bytes.
    Cpmcpumemorylowest interface{}

    // This object represents the upper 32-bit of cpmCPUMemoryUsed. This object
    // needs to be supported only when the value of cpmCPUMemoryUsed exceeds
    // 32-bit, otherwise this object value would be set to 0. The type is
    // interface{} with range: 0..4294967295. Units are kilo-bytes.
    Cpmcpumemoryusedovrflw interface{}

    // The overall CPU wide system memory which is currently under use. This
    // object is a 64-bit version of cpmCPUMemoryUsed. The type is interface{}
    // with range: 0..18446744073709551615. Units are kilo-bytes.
    Cpmcpumemoryhcused interface{}

    // This object represents the upper 32-bit of cpmCPUMemoryFree. This object
    // needs to be supported only when the value of cpmCPUMemoryFree exceeds
    // 32-bit, otherwise this object value would be set to 0. The type is
    // interface{} with range: 0..4294967295. Units are kilo-bytes.
    Cpmcpumemoryfreeovrflw interface{}

    // The overall CPU wide system memory which is currently free. This object is
    // a 64-bit version of cpmCPUMemoryFree. The type is interface{} with range:
    // 0..18446744073709551615. Units are kilo-bytes.
    Cpmcpumemoryhcfree interface{}

    // This object represents the upper 32-bit of cpmCPUMemoryKernelReserved. This
    // object needs  to be supported only when the value of 
    // cpmCPUMemoryKernelReserved exceeds 32-bit, otherwise this object value
    // would be set to 0. The type is interface{} with range: 0..4294967295. Units
    // are kilo-bytes.
    Cpmcpumemorykernelreservedovrflw interface{}

    // The overall CPU wide system memory which is reserved for kernel usage. This
    // object is a 64-bit version of cpmCPUMemoryKernelReserved. The type is
    // interface{} with range: 0..18446744073709551615. Units are kilo-bytes.
    Cpmcpumemoryhckernelreserved interface{}

    // This object represents the upper 32-bit of cpmCPUMemoryLowest. This object
    // needs to be supported only when the value of cpmCPUMemoryLowest exceeds
    // 32-bit, otherwise this object value would be set to 0. The type is
    // interface{} with range: 0..4294967295. Units are kilo-bytes.
    Cpmcpumemorylowestovrflw interface{}

    // The lowest free memory that has been recorded since device has booted. This
    // object is a 64-bit version of cpmCPUMemoryLowest. The type is interface{}
    // with range: 0..18446744073709551615. Units are kilo-bytes.
    Cpmcpumemoryhclowest interface{}

    // The overall CPU load Average in the last 1 minute period. The type is
    // interface{} with range: 0..4294967295. Units are hundredths of processes.
    Cpmcpuloadavg1Min interface{}

    // The overall CPU load Average in the last 5 minutes period. The type is
    // interface{} with range: 0..4294967295. Units are hundredths of processes.
    Cpmcpuloadavg5Min interface{}

    // The overall CPU load Average in the last 15 minutes period. The type is
    // interface{} with range: 0..4294967295. Units are hundredths of processes.
    Cpmcpuloadavg15Min interface{}

    // The overall CPU wide system memory which is currently Committed. The type
    // is interface{} with range: 0..4294967295.
    Cpmcpumemorycommitted interface{}

    // This object represents the upper 32-bit of cpmCPUMemoryCommitted. This
    // object needs to be supported only when the value of cpmCPUMemoryCommitted
    // exceeds 32-bit, otherwise this object value would be set to 0. The type is
    // interface{} with range: 0..4294967295.
    Cpmcpumemorycommittedovrflw interface{}

    // The overall CPU wide system memory which is currently committed. This
    // object is a 64-bit version of cpmCPUMemoryCommitted. The type is
    // interface{} with range: 0..18446744073709551615.
    Cpmcpumemoryhccommitted interface{}
}

func (cpmcputotalentry *CISCOPROCESSMIB_Cpmcputotaltable_Cpmcputotalentry) GetEntityData() *types.CommonEntityData {
    cpmcputotalentry.EntityData.YFilter = cpmcputotalentry.YFilter
    cpmcputotalentry.EntityData.YangName = "cpmCPUTotalEntry"
    cpmcputotalentry.EntityData.BundleName = "cisco_ios_xe"
    cpmcputotalentry.EntityData.ParentYangName = "cpmCPUTotalTable"
    cpmcputotalentry.EntityData.SegmentPath = "cpmCPUTotalEntry" + "[cpmCPUTotalIndex='" + fmt.Sprintf("%v", cpmcputotalentry.Cpmcputotalindex) + "']"
    cpmcputotalentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpmcputotalentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpmcputotalentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpmcputotalentry.EntityData.Children = make(map[string]types.YChild)
    cpmcputotalentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cpmcputotalentry.EntityData.Leafs["cpmCPUTotalIndex"] = types.YLeaf{"Cpmcputotalindex", cpmcputotalentry.Cpmcputotalindex}
    cpmcputotalentry.EntityData.Leafs["cpmCPUTotalPhysicalIndex"] = types.YLeaf{"Cpmcputotalphysicalindex", cpmcputotalentry.Cpmcputotalphysicalindex}
    cpmcputotalentry.EntityData.Leafs["cpmCPUTotal5sec"] = types.YLeaf{"Cpmcputotal5Sec", cpmcputotalentry.Cpmcputotal5Sec}
    cpmcputotalentry.EntityData.Leafs["cpmCPUTotal1min"] = types.YLeaf{"Cpmcputotal1Min", cpmcputotalentry.Cpmcputotal1Min}
    cpmcputotalentry.EntityData.Leafs["cpmCPUTotal5min"] = types.YLeaf{"Cpmcputotal5Min", cpmcputotalentry.Cpmcputotal5Min}
    cpmcputotalentry.EntityData.Leafs["cpmCPUTotal5secRev"] = types.YLeaf{"Cpmcputotal5Secrev", cpmcputotalentry.Cpmcputotal5Secrev}
    cpmcputotalentry.EntityData.Leafs["cpmCPUTotal1minRev"] = types.YLeaf{"Cpmcputotal1Minrev", cpmcputotalentry.Cpmcputotal1Minrev}
    cpmcputotalentry.EntityData.Leafs["cpmCPUTotal5minRev"] = types.YLeaf{"Cpmcputotal5Minrev", cpmcputotalentry.Cpmcputotal5Minrev}
    cpmcputotalentry.EntityData.Leafs["cpmCPUMonInterval"] = types.YLeaf{"Cpmcpumoninterval", cpmcputotalentry.Cpmcpumoninterval}
    cpmcputotalentry.EntityData.Leafs["cpmCPUTotalMonIntervalValue"] = types.YLeaf{"Cpmcputotalmonintervalvalue", cpmcputotalentry.Cpmcputotalmonintervalvalue}
    cpmcputotalentry.EntityData.Leafs["cpmCPUInterruptMonIntervalValue"] = types.YLeaf{"Cpmcpuinterruptmonintervalvalue", cpmcputotalentry.Cpmcpuinterruptmonintervalvalue}
    cpmcputotalentry.EntityData.Leafs["cpmCPUMemoryUsed"] = types.YLeaf{"Cpmcpumemoryused", cpmcputotalentry.Cpmcpumemoryused}
    cpmcputotalentry.EntityData.Leafs["cpmCPUMemoryFree"] = types.YLeaf{"Cpmcpumemoryfree", cpmcputotalentry.Cpmcpumemoryfree}
    cpmcputotalentry.EntityData.Leafs["cpmCPUMemoryKernelReserved"] = types.YLeaf{"Cpmcpumemorykernelreserved", cpmcputotalentry.Cpmcpumemorykernelreserved}
    cpmcputotalentry.EntityData.Leafs["cpmCPUMemoryLowest"] = types.YLeaf{"Cpmcpumemorylowest", cpmcputotalentry.Cpmcpumemorylowest}
    cpmcputotalentry.EntityData.Leafs["cpmCPUMemoryUsedOvrflw"] = types.YLeaf{"Cpmcpumemoryusedovrflw", cpmcputotalentry.Cpmcpumemoryusedovrflw}
    cpmcputotalentry.EntityData.Leafs["cpmCPUMemoryHCUsed"] = types.YLeaf{"Cpmcpumemoryhcused", cpmcputotalentry.Cpmcpumemoryhcused}
    cpmcputotalentry.EntityData.Leafs["cpmCPUMemoryFreeOvrflw"] = types.YLeaf{"Cpmcpumemoryfreeovrflw", cpmcputotalentry.Cpmcpumemoryfreeovrflw}
    cpmcputotalentry.EntityData.Leafs["cpmCPUMemoryHCFree"] = types.YLeaf{"Cpmcpumemoryhcfree", cpmcputotalentry.Cpmcpumemoryhcfree}
    cpmcputotalentry.EntityData.Leafs["cpmCPUMemoryKernelReservedOvrflw"] = types.YLeaf{"Cpmcpumemorykernelreservedovrflw", cpmcputotalentry.Cpmcpumemorykernelreservedovrflw}
    cpmcputotalentry.EntityData.Leafs["cpmCPUMemoryHCKernelReserved"] = types.YLeaf{"Cpmcpumemoryhckernelreserved", cpmcputotalentry.Cpmcpumemoryhckernelreserved}
    cpmcputotalentry.EntityData.Leafs["cpmCPUMemoryLowestOvrflw"] = types.YLeaf{"Cpmcpumemorylowestovrflw", cpmcputotalentry.Cpmcpumemorylowestovrflw}
    cpmcputotalentry.EntityData.Leafs["cpmCPUMemoryHCLowest"] = types.YLeaf{"Cpmcpumemoryhclowest", cpmcputotalentry.Cpmcpumemoryhclowest}
    cpmcputotalentry.EntityData.Leafs["cpmCPULoadAvg1min"] = types.YLeaf{"Cpmcpuloadavg1Min", cpmcputotalentry.Cpmcpuloadavg1Min}
    cpmcputotalentry.EntityData.Leafs["cpmCPULoadAvg5min"] = types.YLeaf{"Cpmcpuloadavg5Min", cpmcputotalentry.Cpmcpuloadavg5Min}
    cpmcputotalentry.EntityData.Leafs["cpmCPULoadAvg15min"] = types.YLeaf{"Cpmcpuloadavg15Min", cpmcputotalentry.Cpmcpuloadavg15Min}
    cpmcputotalentry.EntityData.Leafs["cpmCPUMemoryCommitted"] = types.YLeaf{"Cpmcpumemorycommitted", cpmcputotalentry.Cpmcpumemorycommitted}
    cpmcputotalentry.EntityData.Leafs["cpmCPUMemoryCommittedOvrflw"] = types.YLeaf{"Cpmcpumemorycommittedovrflw", cpmcputotalentry.Cpmcpumemorycommittedovrflw}
    cpmcputotalentry.EntityData.Leafs["cpmCPUMemoryHCCommitted"] = types.YLeaf{"Cpmcpumemoryhccommitted", cpmcputotalentry.Cpmcpumemoryhccommitted}
    return &(cpmcputotalentry.EntityData)
}

// CISCOPROCESSMIB_Cpmcoretable
// A table of per-Core statistics.
type CISCOPROCESSMIB_Cpmcoretable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Overall information about the Core load. Entries in this table could come
    // and go as Cores go online or offline. The type is slice of
    // CISCOPROCESSMIB_Cpmcoretable_Cpmcoreentry.
    Cpmcoreentry []CISCOPROCESSMIB_Cpmcoretable_Cpmcoreentry
}

func (cpmcoretable *CISCOPROCESSMIB_Cpmcoretable) GetEntityData() *types.CommonEntityData {
    cpmcoretable.EntityData.YFilter = cpmcoretable.YFilter
    cpmcoretable.EntityData.YangName = "cpmCoreTable"
    cpmcoretable.EntityData.BundleName = "cisco_ios_xe"
    cpmcoretable.EntityData.ParentYangName = "CISCO-PROCESS-MIB"
    cpmcoretable.EntityData.SegmentPath = "cpmCoreTable"
    cpmcoretable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpmcoretable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpmcoretable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpmcoretable.EntityData.Children = make(map[string]types.YChild)
    cpmcoretable.EntityData.Children["cpmCoreEntry"] = types.YChild{"Cpmcoreentry", nil}
    for i := range cpmcoretable.Cpmcoreentry {
        cpmcoretable.EntityData.Children[types.GetSegmentPath(&cpmcoretable.Cpmcoreentry[i])] = types.YChild{"Cpmcoreentry", &cpmcoretable.Cpmcoreentry[i]}
    }
    cpmcoretable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cpmcoretable.EntityData)
}

// CISCOPROCESSMIB_Cpmcoretable_Cpmcoreentry
// Overall information about the Core load. Entries in this
// table could come and go as Cores go online or offline.
type CISCOPROCESSMIB_Cpmcoretable_Cpmcoreentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // cisco_process_mib.CISCOPROCESSMIB_Cpmcputotaltable_Cpmcputotalentry_Cpmcputotalindex
    Cpmcputotalindex interface{}

    // This attribute is a key. An index that uniquely represents a Core (or group
    // of Cores) whose Core load information is reported by a row in this table.
    // This index is assigned arbitrarily by the engine and is not saved over
    // reboots. The type is interface{} with range: 0..4294967295.
    Cpmcoreindex interface{}

    // The entCorePhysicalIndex of the physical entity for which the Core
    // statistics in this entry are maintained. The physical entity can be a CPU
    // chip, a group of CPUs, a CPU card etc. The exact type of this entity is
    // described by its entPhysicalVendorType value. If the Core statistics in
    // this entry correspond to more than one physical entity (or to no physical
    // entity), or if the entPhysicalTable is not supported on the SNMP agent, the
    // value of this object must be zero. The type is interface{} with range:
    // 0..2147483647.
    Cpmcorephysicalindex interface{}

    // The overall Core busy percentage in the last 5 second period. The type is
    // interface{} with range: 0..100.
    Cpmcore5Sec interface{}

    // The overall Core busy percentage in the last 1 minute period. The type is
    // interface{} with range: 0..100.
    Cpmcore1Min interface{}

    // The overall Core busy percentage in the last 5 minute period. The type is
    // interface{} with range: 0..100.
    Cpmcore5Min interface{}

    // The overall Core load Average in the last 1 minute period. The type is
    // interface{} with range: 0..4294967295. Units are hundredths of processes.
    Cpmcoreloadavg1Min interface{}

    // The overall Core load Average in the last 5 minutes period. The type is
    // interface{} with range: 0..4294967295. Units are hundredths of processes.
    Cpmcoreloadavg5Min interface{}

    // The overall Core load Average in the last 15 minutes period. The type is
    // interface{} with range: 0..4294967295. Units are hundredths of processes.
    Cpmcoreloadavg15Min interface{}
}

func (cpmcoreentry *CISCOPROCESSMIB_Cpmcoretable_Cpmcoreentry) GetEntityData() *types.CommonEntityData {
    cpmcoreentry.EntityData.YFilter = cpmcoreentry.YFilter
    cpmcoreentry.EntityData.YangName = "cpmCoreEntry"
    cpmcoreentry.EntityData.BundleName = "cisco_ios_xe"
    cpmcoreentry.EntityData.ParentYangName = "cpmCoreTable"
    cpmcoreentry.EntityData.SegmentPath = "cpmCoreEntry" + "[cpmCPUTotalIndex='" + fmt.Sprintf("%v", cpmcoreentry.Cpmcputotalindex) + "']" + "[cpmCoreIndex='" + fmt.Sprintf("%v", cpmcoreentry.Cpmcoreindex) + "']"
    cpmcoreentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpmcoreentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpmcoreentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpmcoreentry.EntityData.Children = make(map[string]types.YChild)
    cpmcoreentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cpmcoreentry.EntityData.Leafs["cpmCPUTotalIndex"] = types.YLeaf{"Cpmcputotalindex", cpmcoreentry.Cpmcputotalindex}
    cpmcoreentry.EntityData.Leafs["cpmCoreIndex"] = types.YLeaf{"Cpmcoreindex", cpmcoreentry.Cpmcoreindex}
    cpmcoreentry.EntityData.Leafs["cpmCorePhysicalIndex"] = types.YLeaf{"Cpmcorephysicalindex", cpmcoreentry.Cpmcorephysicalindex}
    cpmcoreentry.EntityData.Leafs["cpmCore5sec"] = types.YLeaf{"Cpmcore5Sec", cpmcoreentry.Cpmcore5Sec}
    cpmcoreentry.EntityData.Leafs["cpmCore1min"] = types.YLeaf{"Cpmcore1Min", cpmcoreentry.Cpmcore1Min}
    cpmcoreentry.EntityData.Leafs["cpmCore5min"] = types.YLeaf{"Cpmcore5Min", cpmcoreentry.Cpmcore5Min}
    cpmcoreentry.EntityData.Leafs["cpmCoreLoadAvg1min"] = types.YLeaf{"Cpmcoreloadavg1Min", cpmcoreentry.Cpmcoreloadavg1Min}
    cpmcoreentry.EntityData.Leafs["cpmCoreLoadAvg5min"] = types.YLeaf{"Cpmcoreloadavg5Min", cpmcoreentry.Cpmcoreloadavg5Min}
    cpmcoreentry.EntityData.Leafs["cpmCoreLoadAvg15min"] = types.YLeaf{"Cpmcoreloadavg15Min", cpmcoreentry.Cpmcoreloadavg15Min}
    return &(cpmcoreentry.EntityData)
}

// CISCOPROCESSMIB_Cpmprocesstable
// A table of generic information on all active
// processes on this device.
type CISCOPROCESSMIB_Cpmprocesstable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Generic information about an active process on this device. Entries in this
    // table come and go as processes are  created and destroyed by the device.
    // The type is slice of CISCOPROCESSMIB_Cpmprocesstable_Cpmprocessentry.
    Cpmprocessentry []CISCOPROCESSMIB_Cpmprocesstable_Cpmprocessentry
}

func (cpmprocesstable *CISCOPROCESSMIB_Cpmprocesstable) GetEntityData() *types.CommonEntityData {
    cpmprocesstable.EntityData.YFilter = cpmprocesstable.YFilter
    cpmprocesstable.EntityData.YangName = "cpmProcessTable"
    cpmprocesstable.EntityData.BundleName = "cisco_ios_xe"
    cpmprocesstable.EntityData.ParentYangName = "CISCO-PROCESS-MIB"
    cpmprocesstable.EntityData.SegmentPath = "cpmProcessTable"
    cpmprocesstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpmprocesstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpmprocesstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpmprocesstable.EntityData.Children = make(map[string]types.YChild)
    cpmprocesstable.EntityData.Children["cpmProcessEntry"] = types.YChild{"Cpmprocessentry", nil}
    for i := range cpmprocesstable.Cpmprocessentry {
        cpmprocesstable.EntityData.Children[types.GetSegmentPath(&cpmprocesstable.Cpmprocessentry[i])] = types.YChild{"Cpmprocessentry", &cpmprocesstable.Cpmprocessentry[i]}
    }
    cpmprocesstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cpmprocesstable.EntityData)
}

// CISCOPROCESSMIB_Cpmprocesstable_Cpmprocessentry
// Generic information about an active process on this
// device. Entries in this table come and go as processes are 
// created and destroyed by the device.
type CISCOPROCESSMIB_Cpmprocesstable_Cpmprocessentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // cisco_process_mib.CISCOPROCESSMIB_Cpmcputotaltable_Cpmcputotalentry_Cpmcputotalindex
    Cpmcputotalindex interface{}

    // This attribute is a key. This object contains the process ID.
    // cpmTimeCreated should be checked against the last time it was polled, and
    // if it has changed the PID has been reused and the entire entry should be
    // polled again. The type is interface{} with range: 0..4294967295.
    Cpmprocesspid interface{}

    // The name associated with this process. If the name is longer than 32
    // characters, it will be truncated to the first 31 characters, and a `*' will
    // be appended as the last character to imply this is a truncated process
    // name. The type is string with length: 1..32.
    Cpmprocessname interface{}

    // Average elapsed CPU time in microseconds when the process was active. This
    // object is deprecated by cpmProcessAverageUSecs. The type is interface{}
    // with range: 0..4294967295. Units are microseconds.
    Cpmprocessusecs interface{}

    // The time when the process was created. The process ID and the time when the
    // process was created, uniquely  identifies a process. The type is
    // interface{} with range: 0..4294967295.
    Cpmprocesstimecreated interface{}

    // Average elapsed CPU time in microseconds when the process was active. This
    // object deprecates the object cpmProcessuSecs. The type is interface{} with
    // range: 0..4294967295. Units are microseconds.
    Cpmprocessaverageusecs interface{}

    // The sum of all the dynamically allocated memory that this process has
    // received from the system. This includes memory that may have been returned.
    // The sum of freed memory is provided by cpmProcExtMemFreed. This object is
    // deprecated by cpmProcExtMemAllocatedRev. The type is interface{} with
    // range: 0..4294967295. Units are bytes.
    Cpmprocextmemallocated interface{}

    // The sum of all memory that this process has returned to the system. This
    // object is deprecated by  cpmProcExtMemFreedRev. The type is interface{}
    // with range: 0..4294967295. Units are bytes.
    Cpmprocextmemfreed interface{}

    // The number of times since cpmTimeCreated that the process has been invoked.
    // This object is deprecated by cpmProcExtInvokedRev. The type is interface{}
    // with range: 0..4294967295.
    Cpmprocextinvoked interface{}

    // The amount of CPU time the process has used, in microseconds. This object
    // is deprecated by cpmProcExtRuntimeRev. The type is interface{} with range:
    // 0..4294967295. Units are microseconds.
    Cpmprocextruntime interface{}

    // This object provides a general idea of how busy a process caused the
    // processor to be over a 5  second period. It is determined as a weighted 
    // decaying average of the current idle time over  the longest idle time. Note
    // that this information  should be used as an estimate only. This object is 
    // deprecated by cpmProcExtUtil5SecRev which has the  changed range of value
    // (0..100). The type is interface{} with range: 0..100.
    Cpmprocextutil5Sec interface{}

    // This object provides a general idea of how busy a process caused the
    // processor to be over a 1  minute period. It is determined as a weighted 
    // decaying average of the current idle time over the  longest idle time. Note
    // that this information  should be used as an estimate only. This object is 
    // deprecated by cpmProcExtUtil1MinRev which has the changed range of value
    // (0..100). The type is interface{} with range: 0..100.
    Cpmprocextutil1Min interface{}

    // This object provides a general idea of how busy a process caused the
    // processor to be over a 5  minute period. It is determined as a weighted 
    // decaying average of the current idle time over  the longest idle time. Note
    // that this information  should be used as an estimate only. This object is
    // deprecated by cpmProcExtUtil5MinRev which has the changed range of value
    // (0..100). The type is interface{} with range: 0..100.
    Cpmprocextutil5Min interface{}

    // The priority level at which the process is running. This object is
    // deprecated by cpmProcExtPriorityRev. The type is Cpmprocextpriority.
    Cpmprocextpriority interface{}
}

func (cpmprocessentry *CISCOPROCESSMIB_Cpmprocesstable_Cpmprocessentry) GetEntityData() *types.CommonEntityData {
    cpmprocessentry.EntityData.YFilter = cpmprocessentry.YFilter
    cpmprocessentry.EntityData.YangName = "cpmProcessEntry"
    cpmprocessentry.EntityData.BundleName = "cisco_ios_xe"
    cpmprocessentry.EntityData.ParentYangName = "cpmProcessTable"
    cpmprocessentry.EntityData.SegmentPath = "cpmProcessEntry" + "[cpmCPUTotalIndex='" + fmt.Sprintf("%v", cpmprocessentry.Cpmcputotalindex) + "']" + "[cpmProcessPID='" + fmt.Sprintf("%v", cpmprocessentry.Cpmprocesspid) + "']"
    cpmprocessentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpmprocessentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpmprocessentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpmprocessentry.EntityData.Children = make(map[string]types.YChild)
    cpmprocessentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cpmprocessentry.EntityData.Leafs["cpmCPUTotalIndex"] = types.YLeaf{"Cpmcputotalindex", cpmprocessentry.Cpmcputotalindex}
    cpmprocessentry.EntityData.Leafs["cpmProcessPID"] = types.YLeaf{"Cpmprocesspid", cpmprocessentry.Cpmprocesspid}
    cpmprocessentry.EntityData.Leafs["cpmProcessName"] = types.YLeaf{"Cpmprocessname", cpmprocessentry.Cpmprocessname}
    cpmprocessentry.EntityData.Leafs["cpmProcessuSecs"] = types.YLeaf{"Cpmprocessusecs", cpmprocessentry.Cpmprocessusecs}
    cpmprocessentry.EntityData.Leafs["cpmProcessTimeCreated"] = types.YLeaf{"Cpmprocesstimecreated", cpmprocessentry.Cpmprocesstimecreated}
    cpmprocessentry.EntityData.Leafs["cpmProcessAverageUSecs"] = types.YLeaf{"Cpmprocessaverageusecs", cpmprocessentry.Cpmprocessaverageusecs}
    cpmprocessentry.EntityData.Leafs["cpmProcExtMemAllocated"] = types.YLeaf{"Cpmprocextmemallocated", cpmprocessentry.Cpmprocextmemallocated}
    cpmprocessentry.EntityData.Leafs["cpmProcExtMemFreed"] = types.YLeaf{"Cpmprocextmemfreed", cpmprocessentry.Cpmprocextmemfreed}
    cpmprocessentry.EntityData.Leafs["cpmProcExtInvoked"] = types.YLeaf{"Cpmprocextinvoked", cpmprocessentry.Cpmprocextinvoked}
    cpmprocessentry.EntityData.Leafs["cpmProcExtRuntime"] = types.YLeaf{"Cpmprocextruntime", cpmprocessentry.Cpmprocextruntime}
    cpmprocessentry.EntityData.Leafs["cpmProcExtUtil5Sec"] = types.YLeaf{"Cpmprocextutil5Sec", cpmprocessentry.Cpmprocextutil5Sec}
    cpmprocessentry.EntityData.Leafs["cpmProcExtUtil1Min"] = types.YLeaf{"Cpmprocextutil1Min", cpmprocessentry.Cpmprocextutil1Min}
    cpmprocessentry.EntityData.Leafs["cpmProcExtUtil5Min"] = types.YLeaf{"Cpmprocextutil5Min", cpmprocessentry.Cpmprocextutil5Min}
    cpmprocessentry.EntityData.Leafs["cpmProcExtPriority"] = types.YLeaf{"Cpmprocextpriority", cpmprocessentry.Cpmprocextpriority}
    return &(cpmprocessentry.EntityData)
}

// CISCOPROCESSMIB_Cpmprocesstable_Cpmprocessentry_Cpmprocextpriority represents cpmProcExtPriorityRev.
type CISCOPROCESSMIB_Cpmprocesstable_Cpmprocessentry_Cpmprocextpriority string

const (
    CISCOPROCESSMIB_Cpmprocesstable_Cpmprocessentry_Cpmprocextpriority_critical CISCOPROCESSMIB_Cpmprocesstable_Cpmprocessentry_Cpmprocextpriority = "critical"

    CISCOPROCESSMIB_Cpmprocesstable_Cpmprocessentry_Cpmprocextpriority_high CISCOPROCESSMIB_Cpmprocesstable_Cpmprocessentry_Cpmprocextpriority = "high"

    CISCOPROCESSMIB_Cpmprocesstable_Cpmprocessentry_Cpmprocextpriority_normal CISCOPROCESSMIB_Cpmprocesstable_Cpmprocessentry_Cpmprocextpriority = "normal"

    CISCOPROCESSMIB_Cpmprocesstable_Cpmprocessentry_Cpmprocextpriority_low CISCOPROCESSMIB_Cpmprocesstable_Cpmprocessentry_Cpmprocextpriority = "low"

    CISCOPROCESSMIB_Cpmprocesstable_Cpmprocessentry_Cpmprocextpriority_notAssigned CISCOPROCESSMIB_Cpmprocesstable_Cpmprocessentry_Cpmprocextpriority = "notAssigned"
)

// CISCOPROCESSMIB_Cpmprocessextrevtable
// This table contains information that may or may
// not be available on all cisco devices. It contains
// additional objects for the more general
// cpmProcessTable. This object deprecates 
// cpmProcessExtTable.
type CISCOPROCESSMIB_Cpmprocessextrevtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry containing additional information for a particular process. This
    // object deprecates  cpmProcessExtEntry. The type is slice of
    // CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry.
    Cpmprocessextreventry []CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry
}

func (cpmprocessextrevtable *CISCOPROCESSMIB_Cpmprocessextrevtable) GetEntityData() *types.CommonEntityData {
    cpmprocessextrevtable.EntityData.YFilter = cpmprocessextrevtable.YFilter
    cpmprocessextrevtable.EntityData.YangName = "cpmProcessExtRevTable"
    cpmprocessextrevtable.EntityData.BundleName = "cisco_ios_xe"
    cpmprocessextrevtable.EntityData.ParentYangName = "CISCO-PROCESS-MIB"
    cpmprocessextrevtable.EntityData.SegmentPath = "cpmProcessExtRevTable"
    cpmprocessextrevtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpmprocessextrevtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpmprocessextrevtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpmprocessextrevtable.EntityData.Children = make(map[string]types.YChild)
    cpmprocessextrevtable.EntityData.Children["cpmProcessExtRevEntry"] = types.YChild{"Cpmprocessextreventry", nil}
    for i := range cpmprocessextrevtable.Cpmprocessextreventry {
        cpmprocessextrevtable.EntityData.Children[types.GetSegmentPath(&cpmprocessextrevtable.Cpmprocessextreventry[i])] = types.YChild{"Cpmprocessextreventry", &cpmprocessextrevtable.Cpmprocessextreventry[i]}
    }
    cpmprocessextrevtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cpmprocessextrevtable.EntityData)
}

// CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry
// An entry containing additional information for
// a particular process. This object deprecates 
// cpmProcessExtEntry.
type CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // cisco_process_mib.CISCOPROCESSMIB_Cpmcputotaltable_Cpmcputotalentry_Cpmcputotalindex
    Cpmcputotalindex interface{}

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // cisco_process_mib.CISCOPROCESSMIB_Cpmprocesstable_Cpmprocessentry_Cpmprocesspid
    Cpmprocesspid interface{}

    // The sum of all the dynamically allocated memory that this process has
    // received from the system. This includes memory that may have been returned.
    // The sum of freed memory is provided by cpmProcExtMemFreedRev. This object
    // deprecates cpmProcExtMemAllocated. The type is interface{} with range:
    // 0..4294967295. Units are bytes.
    Cpmprocextmemallocatedrev interface{}

    // The sum of all memory that this process has returned to the system. This
    // object  deprecates  cpmProcExtMemFreed. The type is interface{} with range:
    // 0..4294967295. Units are bytes.
    Cpmprocextmemfreedrev interface{}

    // The number of times since cpmTimeCreated that the process has been invoked.
    // This object  deprecates cpmProcExtInvoked. The type is interface{} with
    // range: 0..4294967295.
    Cpmprocextinvokedrev interface{}

    // The amount of CPU time the process has used, in microseconds. This object
    // deprecates cpmProcExtRuntime. The type is interface{} with range:
    // 0..4294967295. Units are microseconds.
    Cpmprocextruntimerev interface{}

    // This object provides a general idea of how busy a process caused the
    // processor to be over a 5  second period. It is determined as a weighted 
    // decaying average of the current idle time over  the longest idle time. Note
    // that this information  should be used as an estimate only. This object
    // deprecates cpmProcExtUtil5Sec and increases the  value range to (0..100).
    // The type is interface{} with range: 0..100. Units are percent.
    Cpmprocextutil5Secrev interface{}

    // This object provides a general idea of how busy a process caused the
    // processor to be over a 1  minute period. It is determined as a weighted 
    // decaying average of the current idle time over the  longest idle time. Note
    // that this information  should be used as an estimate only. This object 
    // deprecates cpmProcExtUtil1Min and increases the value range to (0..100).
    // The type is interface{} with range: 0..100. Units are percent.
    Cpmprocextutil1Minrev interface{}

    // This object provides a general idea of how busy a process caused the
    // processor to be over a 5  minute period. It is determined as a weighted 
    // decaying average of the current idle time over  the longest idle time. Note
    // that this information  should be used as an estimate only. This object
    // deprecates cpmProcExtUtil5Min and increases the value range to (0..100).
    // The type is interface{} with range: 0..100. Units are percent.
    Cpmprocextutil5Minrev interface{}

    // The priority level at  which the process is running. This object deprecates
    // cpmProcExtPriority. The type is Cpmprocextpriorityrev.
    Cpmprocextpriorityrev interface{}

    // This indicates the kind of process in context. The type is Cpmprocesstype.
    Cpmprocesstype interface{}

    // This indicates whether respawn of a process is enabled or not. If enabled
    // the process in context repawns after it has crashed/stopped. The type is
    // interface{} with range: 0..2.
    Cpmprocessrespawn interface{}

    // This indicates the number of times the process has respawned/restarted. The
    // type is interface{} with range: 0..4294967295.
    Cpmprocessrespawncount interface{}

    // This indicates the number of times a process has restarted after the last
    // patch is applied. This is to  determine the stability of the last patch.
    // The type is interface{} with range: 0..4294967295.
    Cpmprocessrespawnafterlastpatch interface{}

    // This indicates the part of process memory to be dumped when a process
    // crashes. The process  memory is used for debugging purposes to trace the 
    // root cause of the crash. sparse        - Some operating systems support
    // minimal                 dump of process core like register                
    // info, partial stack, partial memory                 pages especially for
    // critical process                 to facilitate faster process restart. The
    // type is Cpmprocessmemorycore.
    Cpmprocessmemorycore interface{}

    // This indicate the user that has last restarted the process or has taken
    // running coredump of the process. The type is string.
    Cpmprocesslastrestartuser interface{}

    // This indicates the text memory of a process and all its shared objects. The
    // type is interface{} with range: 0..4294967295. Units are kilo-bytes.
    Cpmprocesstextsegmentsize interface{}

    // This indicates the data segment of a process and all its shared objects.
    // The type is interface{} with range: 0..4294967295. Units are kilo-bytes.
    Cpmprocessdatasegmentsize interface{}

    // This indicates the amount of stack memory used by the process. The type is
    // interface{} with range: 0..4294967295. Units are kilo-bytes.
    Cpmprocessstacksize interface{}

    // This indicates the amount of dynamic memory being used by the process. The
    // type is interface{} with range: 0..4294967295. Units are kilo-bytes.
    Cpmprocessdynamicmemorysize interface{}

    // This object represents the upper 32-bit of cpmProcExtMemAllocatedRev. This
    // object needs to be supported only when the value of 
    // cpmProcExtMemAllocatedRev exceeds 32-bit,  otherwise this object value
    // would be set to 0. The type is interface{} with range: 0..4294967295. Units
    // are bytes.
    Cpmprocextmemallocatedrevovrflw interface{}

    // The sum of all the dynamically allocated memory that this process has
    // received from the system. This includes memory that may have been returned.
    // This object is a 64-bit version of cpmProcExtMemAllocatedRev. The type is
    // interface{} with range: 0..18446744073709551615. Units are bytes.
    Cpmprocexthcmemallocatedrev interface{}

    // This object represents the upper 32-bit of cpmProcExtMemFreedRev. This
    // object needs to  be supported only when the value of cpmProcExtMemFreedRev
    // exceeds 32-bit,otherwise this object value would be set to 0. The type is
    // interface{} with range: 0..4294967295. Units are bytes.
    Cpmprocextmemfreedrevovrflw interface{}

    // The sum of all memory that this process has returned to the system. This
    // object is a 64-bit version of cpmProcExtMemFreedRev. The type is
    // interface{} with range: 0..18446744073709551615. Units are bytes.
    Cpmprocexthcmemfreedrev interface{}

    // This object represents the upper 32-bit of cpmProcessTextSegmentSize. This
    // object needs to be supported only when the value of 
    // cpmProcessTextSegmentSize exceeds 32-bit,  otherwise this object value
    // would be set to 0. The type is interface{} with range: 0..4294967295. Units
    // are kilo-bytes.
    Cpmprocesstextsegmentsizeovrflw interface{}

    // This indicates the text memory of a process and all its shared objects.
    // This object is a 64-bit version of cpmProcessTextSegmentSize. The type is
    // interface{} with range: 0..18446744073709551615. Units are kilo-bytes.
    Cpmprocesshctextsegmentsize interface{}

    // This object represents the upper 32-bit of cpmProcessDataSegmentSize. This
    // object needs to be supported only when the value of 
    // cpmProcessDataSegmentSize exceeds 32-bit,  otherwise this object value
    // would be set to 0. The type is interface{} with range: 0..4294967295. Units
    // are kilo-bytes.
    Cpmprocessdatasegmentsizeovrflw interface{}

    // This indicates the data segment of a process and all its shared objects..
    // This object is a 64-bit version of cpmProcessDataSegmentSize. The type is
    // interface{} with range: 0..18446744073709551615. Units are kilo-bytes.
    Cpmprocesshcdatasegmentsize interface{}

    // This object represents the upper 32-bit of cpmProcessStackSize. This object
    // needs to be supported only when the value of cpmProcessStackSize exceeds
    // 32-bit, otherwise this object value would be set to 0. The type is
    // interface{} with range: 0..4294967295. Units are kilo-bytes.
    Cpmprocessstacksizeovrflw interface{}

    // This indicates the amount of stack memory used by the process. This object
    // is a 64-bit version of cpmProcessStackSize. The type is interface{} with
    // range: 0..18446744073709551615. Units are kilo-bytes.
    Cpmprocesshcstacksize interface{}

    // This object represents the upper 32-bit of cpmProcessDynamicMemorySize.
    // This object needs to be supported only when the value of 
    // cpmProcessDynamicMemorySize exceeds 32-bit,  otherwise this object value
    // would be set to 0. The type is interface{} with range: 0..4294967295. Units
    // are kilo-bytes.
    Cpmprocessdynamicmemorysizeovrflw interface{}

    // This indicates the amount of dynamic memory being used by the process. This
    // object is a 64-bit version of cpmProcessDynamicMemorySize. The type is
    // interface{} with range: 0..18446744073709551615. Units are kilo-bytes.
    Cpmprocesshcdynamicmemorysize interface{}
}

func (cpmprocessextreventry *CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry) GetEntityData() *types.CommonEntityData {
    cpmprocessextreventry.EntityData.YFilter = cpmprocessextreventry.YFilter
    cpmprocessextreventry.EntityData.YangName = "cpmProcessExtRevEntry"
    cpmprocessextreventry.EntityData.BundleName = "cisco_ios_xe"
    cpmprocessextreventry.EntityData.ParentYangName = "cpmProcessExtRevTable"
    cpmprocessextreventry.EntityData.SegmentPath = "cpmProcessExtRevEntry" + "[cpmCPUTotalIndex='" + fmt.Sprintf("%v", cpmprocessextreventry.Cpmcputotalindex) + "']" + "[cpmProcessPID='" + fmt.Sprintf("%v", cpmprocessextreventry.Cpmprocesspid) + "']"
    cpmprocessextreventry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpmprocessextreventry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpmprocessextreventry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpmprocessextreventry.EntityData.Children = make(map[string]types.YChild)
    cpmprocessextreventry.EntityData.Leafs = make(map[string]types.YLeaf)
    cpmprocessextreventry.EntityData.Leafs["cpmCPUTotalIndex"] = types.YLeaf{"Cpmcputotalindex", cpmprocessextreventry.Cpmcputotalindex}
    cpmprocessextreventry.EntityData.Leafs["cpmProcessPID"] = types.YLeaf{"Cpmprocesspid", cpmprocessextreventry.Cpmprocesspid}
    cpmprocessextreventry.EntityData.Leafs["cpmProcExtMemAllocatedRev"] = types.YLeaf{"Cpmprocextmemallocatedrev", cpmprocessextreventry.Cpmprocextmemallocatedrev}
    cpmprocessextreventry.EntityData.Leafs["cpmProcExtMemFreedRev"] = types.YLeaf{"Cpmprocextmemfreedrev", cpmprocessextreventry.Cpmprocextmemfreedrev}
    cpmprocessextreventry.EntityData.Leafs["cpmProcExtInvokedRev"] = types.YLeaf{"Cpmprocextinvokedrev", cpmprocessextreventry.Cpmprocextinvokedrev}
    cpmprocessextreventry.EntityData.Leafs["cpmProcExtRuntimeRev"] = types.YLeaf{"Cpmprocextruntimerev", cpmprocessextreventry.Cpmprocextruntimerev}
    cpmprocessextreventry.EntityData.Leafs["cpmProcExtUtil5SecRev"] = types.YLeaf{"Cpmprocextutil5Secrev", cpmprocessextreventry.Cpmprocextutil5Secrev}
    cpmprocessextreventry.EntityData.Leafs["cpmProcExtUtil1MinRev"] = types.YLeaf{"Cpmprocextutil1Minrev", cpmprocessextreventry.Cpmprocextutil1Minrev}
    cpmprocessextreventry.EntityData.Leafs["cpmProcExtUtil5MinRev"] = types.YLeaf{"Cpmprocextutil5Minrev", cpmprocessextreventry.Cpmprocextutil5Minrev}
    cpmprocessextreventry.EntityData.Leafs["cpmProcExtPriorityRev"] = types.YLeaf{"Cpmprocextpriorityrev", cpmprocessextreventry.Cpmprocextpriorityrev}
    cpmprocessextreventry.EntityData.Leafs["cpmProcessType"] = types.YLeaf{"Cpmprocesstype", cpmprocessextreventry.Cpmprocesstype}
    cpmprocessextreventry.EntityData.Leafs["cpmProcessRespawn"] = types.YLeaf{"Cpmprocessrespawn", cpmprocessextreventry.Cpmprocessrespawn}
    cpmprocessextreventry.EntityData.Leafs["cpmProcessRespawnCount"] = types.YLeaf{"Cpmprocessrespawncount", cpmprocessextreventry.Cpmprocessrespawncount}
    cpmprocessextreventry.EntityData.Leafs["cpmProcessRespawnAfterLastPatch"] = types.YLeaf{"Cpmprocessrespawnafterlastpatch", cpmprocessextreventry.Cpmprocessrespawnafterlastpatch}
    cpmprocessextreventry.EntityData.Leafs["cpmProcessMemoryCore"] = types.YLeaf{"Cpmprocessmemorycore", cpmprocessextreventry.Cpmprocessmemorycore}
    cpmprocessextreventry.EntityData.Leafs["cpmProcessLastRestartUser"] = types.YLeaf{"Cpmprocesslastrestartuser", cpmprocessextreventry.Cpmprocesslastrestartuser}
    cpmprocessextreventry.EntityData.Leafs["cpmProcessTextSegmentSize"] = types.YLeaf{"Cpmprocesstextsegmentsize", cpmprocessextreventry.Cpmprocesstextsegmentsize}
    cpmprocessextreventry.EntityData.Leafs["cpmProcessDataSegmentSize"] = types.YLeaf{"Cpmprocessdatasegmentsize", cpmprocessextreventry.Cpmprocessdatasegmentsize}
    cpmprocessextreventry.EntityData.Leafs["cpmProcessStackSize"] = types.YLeaf{"Cpmprocessstacksize", cpmprocessextreventry.Cpmprocessstacksize}
    cpmprocessextreventry.EntityData.Leafs["cpmProcessDynamicMemorySize"] = types.YLeaf{"Cpmprocessdynamicmemorysize", cpmprocessextreventry.Cpmprocessdynamicmemorysize}
    cpmprocessextreventry.EntityData.Leafs["cpmProcExtMemAllocatedRevOvrflw"] = types.YLeaf{"Cpmprocextmemallocatedrevovrflw", cpmprocessextreventry.Cpmprocextmemallocatedrevovrflw}
    cpmprocessextreventry.EntityData.Leafs["cpmProcExtHCMemAllocatedRev"] = types.YLeaf{"Cpmprocexthcmemallocatedrev", cpmprocessextreventry.Cpmprocexthcmemallocatedrev}
    cpmprocessextreventry.EntityData.Leafs["cpmProcExtMemFreedRevOvrflw"] = types.YLeaf{"Cpmprocextmemfreedrevovrflw", cpmprocessextreventry.Cpmprocextmemfreedrevovrflw}
    cpmprocessextreventry.EntityData.Leafs["cpmProcExtHCMemFreedRev"] = types.YLeaf{"Cpmprocexthcmemfreedrev", cpmprocessextreventry.Cpmprocexthcmemfreedrev}
    cpmprocessextreventry.EntityData.Leafs["cpmProcessTextSegmentSizeOvrflw"] = types.YLeaf{"Cpmprocesstextsegmentsizeovrflw", cpmprocessextreventry.Cpmprocesstextsegmentsizeovrflw}
    cpmprocessextreventry.EntityData.Leafs["cpmProcessHCTextSegmentSize"] = types.YLeaf{"Cpmprocesshctextsegmentsize", cpmprocessextreventry.Cpmprocesshctextsegmentsize}
    cpmprocessextreventry.EntityData.Leafs["cpmProcessDataSegmentSizeOvrflw"] = types.YLeaf{"Cpmprocessdatasegmentsizeovrflw", cpmprocessextreventry.Cpmprocessdatasegmentsizeovrflw}
    cpmprocessextreventry.EntityData.Leafs["cpmProcessHCDataSegmentSize"] = types.YLeaf{"Cpmprocesshcdatasegmentsize", cpmprocessextreventry.Cpmprocesshcdatasegmentsize}
    cpmprocessextreventry.EntityData.Leafs["cpmProcessStackSizeOvrflw"] = types.YLeaf{"Cpmprocessstacksizeovrflw", cpmprocessextreventry.Cpmprocessstacksizeovrflw}
    cpmprocessextreventry.EntityData.Leafs["cpmProcessHCStackSize"] = types.YLeaf{"Cpmprocesshcstacksize", cpmprocessextreventry.Cpmprocesshcstacksize}
    cpmprocessextreventry.EntityData.Leafs["cpmProcessDynamicMemorySizeOvrflw"] = types.YLeaf{"Cpmprocessdynamicmemorysizeovrflw", cpmprocessextreventry.Cpmprocessdynamicmemorysizeovrflw}
    cpmprocessextreventry.EntityData.Leafs["cpmProcessHCDynamicMemorySize"] = types.YLeaf{"Cpmprocesshcdynamicmemorysize", cpmprocessextreventry.Cpmprocesshcdynamicmemorysize}
    return &(cpmprocessextreventry.EntityData)
}

// CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry_Cpmprocessmemorycore represents                 to facilitate faster process restart.
type CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry_Cpmprocessmemorycore string

const (
    CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry_Cpmprocessmemorycore_none CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry_Cpmprocessmemorycore = "none"

    CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry_Cpmprocessmemorycore_other CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry_Cpmprocessmemorycore = "other"

    CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry_Cpmprocessmemorycore_mainmem CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry_Cpmprocessmemorycore = "mainmem"

    CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry_Cpmprocessmemorycore_mainmemSharedmem CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry_Cpmprocessmemorycore = "mainmemSharedmem"

    CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry_Cpmprocessmemorycore_mainmemText CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry_Cpmprocessmemorycore = "mainmemText"

    CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry_Cpmprocessmemorycore_mainmemTextSharedmem CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry_Cpmprocessmemorycore = "mainmemTextSharedmem"

    CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry_Cpmprocessmemorycore_sharedmem CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry_Cpmprocessmemorycore = "sharedmem"

    CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry_Cpmprocessmemorycore_sparse CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry_Cpmprocessmemorycore = "sparse"

    CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry_Cpmprocessmemorycore_off CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry_Cpmprocessmemorycore = "off"
)

// CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry_Cpmprocesstype represents This indicates the kind of process in context.
type CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry_Cpmprocesstype string

const (
    CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry_Cpmprocesstype_none CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry_Cpmprocesstype = "none"

    CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry_Cpmprocesstype_other CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry_Cpmprocesstype = "other"

    CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry_Cpmprocesstype_posix CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry_Cpmprocesstype = "posix"

    CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry_Cpmprocesstype_ios CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry_Cpmprocesstype = "ios"
)

// CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry_Cpmprocextpriorityrev represents cpmProcExtPriority.
type CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry_Cpmprocextpriorityrev string

const (
    CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry_Cpmprocextpriorityrev_critical CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry_Cpmprocextpriorityrev = "critical"

    CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry_Cpmprocextpriorityrev_high CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry_Cpmprocextpriorityrev = "high"

    CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry_Cpmprocextpriorityrev_normal CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry_Cpmprocextpriorityrev = "normal"

    CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry_Cpmprocextpriorityrev_low CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry_Cpmprocextpriorityrev = "low"

    CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry_Cpmprocextpriorityrev_notAssigned CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry_Cpmprocextpriorityrev = "notAssigned"
)

// CISCOPROCESSMIB_Cpmcputhresholdtable
// This table contains the information about the
// thresholding values for CPU , configured by the user.
type CISCOPROCESSMIB_Cpmcputhresholdtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry containing information about CPU thresholding parameters.
    // cpmCPUTotalIndex identifies the CPU (or group of CPUs) for which this
    // configuration applies. The type is slice of
    // CISCOPROCESSMIB_Cpmcputhresholdtable_Cpmcputhresholdentry.
    Cpmcputhresholdentry []CISCOPROCESSMIB_Cpmcputhresholdtable_Cpmcputhresholdentry
}

func (cpmcputhresholdtable *CISCOPROCESSMIB_Cpmcputhresholdtable) GetEntityData() *types.CommonEntityData {
    cpmcputhresholdtable.EntityData.YFilter = cpmcputhresholdtable.YFilter
    cpmcputhresholdtable.EntityData.YangName = "cpmCPUThresholdTable"
    cpmcputhresholdtable.EntityData.BundleName = "cisco_ios_xe"
    cpmcputhresholdtable.EntityData.ParentYangName = "CISCO-PROCESS-MIB"
    cpmcputhresholdtable.EntityData.SegmentPath = "cpmCPUThresholdTable"
    cpmcputhresholdtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpmcputhresholdtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpmcputhresholdtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpmcputhresholdtable.EntityData.Children = make(map[string]types.YChild)
    cpmcputhresholdtable.EntityData.Children["cpmCPUThresholdEntry"] = types.YChild{"Cpmcputhresholdentry", nil}
    for i := range cpmcputhresholdtable.Cpmcputhresholdentry {
        cpmcputhresholdtable.EntityData.Children[types.GetSegmentPath(&cpmcputhresholdtable.Cpmcputhresholdentry[i])] = types.YChild{"Cpmcputhresholdentry", &cpmcputhresholdtable.Cpmcputhresholdentry[i]}
    }
    cpmcputhresholdtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cpmcputhresholdtable.EntityData)
}

// CISCOPROCESSMIB_Cpmcputhresholdtable_Cpmcputhresholdentry
// An entry containing information about
// CPU thresholding parameters. cpmCPUTotalIndex
// identifies the CPU (or group of CPUs) for which this
// configuration applies.
type CISCOPROCESSMIB_Cpmcputhresholdtable_Cpmcputhresholdentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // cisco_process_mib.CISCOPROCESSMIB_Cpmcputotaltable_Cpmcputotalentry_Cpmcputotalindex
    Cpmcputotalindex interface{}

    // This attribute is a key. Value of this object indicates the type of
    // utilization, which is monitored. The total(1) indicates the total CPU
    // utilization, interrupt(2) indicates the the CPU utilization in interrupt
    // context and process(3) indicates the CPU utilization in the process level
    // execution context. The type is Cpmcputhresholdclass.
    Cpmcputhresholdclass interface{}

    // The percentage rising threshold value configured by the user. The value
    // indicates,  if the percentage CPU utilization is equal to or above this
    // value for cpmCPURisingThresholdPeriod duration  then send a
    // cpmCPURisingThreshold notification to the NMS. The type is interface{} with
    // range: 1..100.
    Cpmcpurisingthresholdvalue interface{}

    // This is an observation interval. The value of this object indicates that 
    // the CPU utilization should be above cpmCPURisingThresholdValue for this
    // duration to send a  cpmCPURisingThreshold notification to the NMS. The type
    // is interface{} with range: 5..4294967295. Units are seconds.
    Cpmcpurisingthresholdperiod interface{}

    // The percentage falling threshold value configured by the user. The value
    // indicates, if the percentage  CPU utilization is equal to or below this
    // value for  cpmCPUFallingThresholdPeriod duration then send a
    // cpmCPUFallingThreshold notification  to the NMS. The type is interface{}
    // with range: 1..100.
    Cpmcpufallingthresholdvalue interface{}

    // This is an observation interval. The value of this object indicates that
    // CPU utilization should be below cpmCPUFallingThresholdValue for this
    // duration to send a  cpmCPURisingThreshold notification to the NMS. The type
    // is interface{} with range: 5..4294967295. Units are seconds.
    Cpmcpufallingthresholdperiod interface{}

    // The status of this table entry. The type is RowStatus.
    Cpmcputhresholdentrystatus interface{}
}

func (cpmcputhresholdentry *CISCOPROCESSMIB_Cpmcputhresholdtable_Cpmcputhresholdentry) GetEntityData() *types.CommonEntityData {
    cpmcputhresholdentry.EntityData.YFilter = cpmcputhresholdentry.YFilter
    cpmcputhresholdentry.EntityData.YangName = "cpmCPUThresholdEntry"
    cpmcputhresholdentry.EntityData.BundleName = "cisco_ios_xe"
    cpmcputhresholdentry.EntityData.ParentYangName = "cpmCPUThresholdTable"
    cpmcputhresholdentry.EntityData.SegmentPath = "cpmCPUThresholdEntry" + "[cpmCPUTotalIndex='" + fmt.Sprintf("%v", cpmcputhresholdentry.Cpmcputotalindex) + "']" + "[cpmCPUThresholdClass='" + fmt.Sprintf("%v", cpmcputhresholdentry.Cpmcputhresholdclass) + "']"
    cpmcputhresholdentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpmcputhresholdentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpmcputhresholdentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpmcputhresholdentry.EntityData.Children = make(map[string]types.YChild)
    cpmcputhresholdentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cpmcputhresholdentry.EntityData.Leafs["cpmCPUTotalIndex"] = types.YLeaf{"Cpmcputotalindex", cpmcputhresholdentry.Cpmcputotalindex}
    cpmcputhresholdentry.EntityData.Leafs["cpmCPUThresholdClass"] = types.YLeaf{"Cpmcputhresholdclass", cpmcputhresholdentry.Cpmcputhresholdclass}
    cpmcputhresholdentry.EntityData.Leafs["cpmCPURisingThresholdValue"] = types.YLeaf{"Cpmcpurisingthresholdvalue", cpmcputhresholdentry.Cpmcpurisingthresholdvalue}
    cpmcputhresholdentry.EntityData.Leafs["cpmCPURisingThresholdPeriod"] = types.YLeaf{"Cpmcpurisingthresholdperiod", cpmcputhresholdentry.Cpmcpurisingthresholdperiod}
    cpmcputhresholdentry.EntityData.Leafs["cpmCPUFallingThresholdValue"] = types.YLeaf{"Cpmcpufallingthresholdvalue", cpmcputhresholdentry.Cpmcpufallingthresholdvalue}
    cpmcputhresholdentry.EntityData.Leafs["cpmCPUFallingThresholdPeriod"] = types.YLeaf{"Cpmcpufallingthresholdperiod", cpmcputhresholdentry.Cpmcpufallingthresholdperiod}
    cpmcputhresholdentry.EntityData.Leafs["cpmCPUThresholdEntryStatus"] = types.YLeaf{"Cpmcputhresholdentrystatus", cpmcputhresholdentry.Cpmcputhresholdentrystatus}
    return &(cpmcputhresholdentry.EntityData)
}

// CISCOPROCESSMIB_Cpmcputhresholdtable_Cpmcputhresholdentry_Cpmcputhresholdclass represents execution context.
type CISCOPROCESSMIB_Cpmcputhresholdtable_Cpmcputhresholdentry_Cpmcputhresholdclass string

const (
    CISCOPROCESSMIB_Cpmcputhresholdtable_Cpmcputhresholdentry_Cpmcputhresholdclass_total CISCOPROCESSMIB_Cpmcputhresholdtable_Cpmcputhresholdentry_Cpmcputhresholdclass = "total"

    CISCOPROCESSMIB_Cpmcputhresholdtable_Cpmcputhresholdentry_Cpmcputhresholdclass_interrupt CISCOPROCESSMIB_Cpmcputhresholdtable_Cpmcputhresholdentry_Cpmcputhresholdclass = "interrupt"

    CISCOPROCESSMIB_Cpmcputhresholdtable_Cpmcputhresholdentry_Cpmcputhresholdclass_process CISCOPROCESSMIB_Cpmcputhresholdtable_Cpmcputhresholdentry_Cpmcputhresholdclass = "process"
)

// CISCOPROCESSMIB_Cpmcpuhistorytable
// A list of CPU utilization history entries.
type CISCOPROCESSMIB_Cpmcpuhistorytable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A historical sample of CPU utilization statistics. cpmCPUTotalIndex
    // identifies the CPU (or group of CPUs) for which this history is collected. 
    // When the cpmCPUHistorySize is reached the least recent entry is lost. The
    // type is slice of CISCOPROCESSMIB_Cpmcpuhistorytable_Cpmcpuhistoryentry.
    Cpmcpuhistoryentry []CISCOPROCESSMIB_Cpmcpuhistorytable_Cpmcpuhistoryentry
}

func (cpmcpuhistorytable *CISCOPROCESSMIB_Cpmcpuhistorytable) GetEntityData() *types.CommonEntityData {
    cpmcpuhistorytable.EntityData.YFilter = cpmcpuhistorytable.YFilter
    cpmcpuhistorytable.EntityData.YangName = "cpmCPUHistoryTable"
    cpmcpuhistorytable.EntityData.BundleName = "cisco_ios_xe"
    cpmcpuhistorytable.EntityData.ParentYangName = "CISCO-PROCESS-MIB"
    cpmcpuhistorytable.EntityData.SegmentPath = "cpmCPUHistoryTable"
    cpmcpuhistorytable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpmcpuhistorytable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpmcpuhistorytable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpmcpuhistorytable.EntityData.Children = make(map[string]types.YChild)
    cpmcpuhistorytable.EntityData.Children["cpmCPUHistoryEntry"] = types.YChild{"Cpmcpuhistoryentry", nil}
    for i := range cpmcpuhistorytable.Cpmcpuhistoryentry {
        cpmcpuhistorytable.EntityData.Children[types.GetSegmentPath(&cpmcpuhistorytable.Cpmcpuhistoryentry[i])] = types.YChild{"Cpmcpuhistoryentry", &cpmcpuhistorytable.Cpmcpuhistoryentry[i]}
    }
    cpmcpuhistorytable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cpmcpuhistorytable.EntityData)
}

// CISCOPROCESSMIB_Cpmcpuhistorytable_Cpmcpuhistoryentry
// A historical sample of CPU utilization statistics.
// cpmCPUTotalIndex identifies the CPU (or group of CPUs)
// for which this history is collected. 
// When the cpmCPUHistorySize is
// reached the least recent entry is lost.
type CISCOPROCESSMIB_Cpmcpuhistorytable_Cpmcpuhistoryentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // cisco_process_mib.CISCOPROCESSMIB_Cpmcputotaltable_Cpmcputotalentry_Cpmcputotalindex
    Cpmcputotalindex interface{}

    // This attribute is a key. All the entries which are created at the same time
    // will have same value for this object. When the configured threshold for
    // being a part of History table is reached then the qualified processes
    // become the part of history table. The entries which became the  part of
    // history table at one instant will have the same value for this object. When
    // this object reaches the max index value then it will wrap around. The type
    // is interface{} with range: 0..4294967295.
    Cpmcpuhistoryreportid interface{}

    // The number of process entries in a report. This object gives information
    // about how many processes  became a part of history table at one instant.
    // The type is interface{} with range: 0..4294967295.
    Cpmcpuhistoryreportsize interface{}

    // Total percentage of CPU utilization at cpmCPUHistoryCreated. The type is
    // interface{} with range: 0..100. Units are percent.
    Cpmcpuhistorytotalutil interface{}

    // Percentage of CPU utilization in the interrupt context at
    // cpmCPUHistoryCreated. The type is interface{} with range: 0..100. Units are
    // percent.
    Cpmcpuhistoryinterruptutil interface{}

    // Time stamp with respect to sysUpTime indicating the time at which this
    // report is created. The type is interface{} with range: 0..4294967295.
    Cpmcpuhistorycreatedtime interface{}
}

func (cpmcpuhistoryentry *CISCOPROCESSMIB_Cpmcpuhistorytable_Cpmcpuhistoryentry) GetEntityData() *types.CommonEntityData {
    cpmcpuhistoryentry.EntityData.YFilter = cpmcpuhistoryentry.YFilter
    cpmcpuhistoryentry.EntityData.YangName = "cpmCPUHistoryEntry"
    cpmcpuhistoryentry.EntityData.BundleName = "cisco_ios_xe"
    cpmcpuhistoryentry.EntityData.ParentYangName = "cpmCPUHistoryTable"
    cpmcpuhistoryentry.EntityData.SegmentPath = "cpmCPUHistoryEntry" + "[cpmCPUTotalIndex='" + fmt.Sprintf("%v", cpmcpuhistoryentry.Cpmcputotalindex) + "']" + "[cpmCPUHistoryReportId='" + fmt.Sprintf("%v", cpmcpuhistoryentry.Cpmcpuhistoryreportid) + "']"
    cpmcpuhistoryentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpmcpuhistoryentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpmcpuhistoryentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpmcpuhistoryentry.EntityData.Children = make(map[string]types.YChild)
    cpmcpuhistoryentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cpmcpuhistoryentry.EntityData.Leafs["cpmCPUTotalIndex"] = types.YLeaf{"Cpmcputotalindex", cpmcpuhistoryentry.Cpmcputotalindex}
    cpmcpuhistoryentry.EntityData.Leafs["cpmCPUHistoryReportId"] = types.YLeaf{"Cpmcpuhistoryreportid", cpmcpuhistoryentry.Cpmcpuhistoryreportid}
    cpmcpuhistoryentry.EntityData.Leafs["cpmCPUHistoryReportSize"] = types.YLeaf{"Cpmcpuhistoryreportsize", cpmcpuhistoryentry.Cpmcpuhistoryreportsize}
    cpmcpuhistoryentry.EntityData.Leafs["cpmCPUHistoryTotalUtil"] = types.YLeaf{"Cpmcpuhistorytotalutil", cpmcpuhistoryentry.Cpmcpuhistorytotalutil}
    cpmcpuhistoryentry.EntityData.Leafs["cpmCPUHistoryInterruptUtil"] = types.YLeaf{"Cpmcpuhistoryinterruptutil", cpmcpuhistoryentry.Cpmcpuhistoryinterruptutil}
    cpmcpuhistoryentry.EntityData.Leafs["cpmCPUHistoryCreatedTime"] = types.YLeaf{"Cpmcpuhistorycreatedtime", cpmcpuhistoryentry.Cpmcpuhistorycreatedtime}
    return &(cpmcpuhistoryentry.EntityData)
}

// CISCOPROCESSMIB_Cpmcpuprocesshistorytable
// A list of process history entries. This table contains
// CPU utilization of processes which crossed the 
// cpmCPUHistoryThreshold.
type CISCOPROCESSMIB_Cpmcpuprocesshistorytable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A historical sample of process utilization statistics. The entries in this
    // table will have corresponding entires in the cpmCPUHistoryTable. The
    // entries in this table get deleted when the entry associated with this entry
    // in the cpmCPUHistoryTable  gets deleted. The type is slice of
    // CISCOPROCESSMIB_Cpmcpuprocesshistorytable_Cpmcpuprocesshistoryentry.
    Cpmcpuprocesshistoryentry []CISCOPROCESSMIB_Cpmcpuprocesshistorytable_Cpmcpuprocesshistoryentry
}

func (cpmcpuprocesshistorytable *CISCOPROCESSMIB_Cpmcpuprocesshistorytable) GetEntityData() *types.CommonEntityData {
    cpmcpuprocesshistorytable.EntityData.YFilter = cpmcpuprocesshistorytable.YFilter
    cpmcpuprocesshistorytable.EntityData.YangName = "cpmCPUProcessHistoryTable"
    cpmcpuprocesshistorytable.EntityData.BundleName = "cisco_ios_xe"
    cpmcpuprocesshistorytable.EntityData.ParentYangName = "CISCO-PROCESS-MIB"
    cpmcpuprocesshistorytable.EntityData.SegmentPath = "cpmCPUProcessHistoryTable"
    cpmcpuprocesshistorytable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpmcpuprocesshistorytable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpmcpuprocesshistorytable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpmcpuprocesshistorytable.EntityData.Children = make(map[string]types.YChild)
    cpmcpuprocesshistorytable.EntityData.Children["cpmCPUProcessHistoryEntry"] = types.YChild{"Cpmcpuprocesshistoryentry", nil}
    for i := range cpmcpuprocesshistorytable.Cpmcpuprocesshistoryentry {
        cpmcpuprocesshistorytable.EntityData.Children[types.GetSegmentPath(&cpmcpuprocesshistorytable.Cpmcpuprocesshistoryentry[i])] = types.YChild{"Cpmcpuprocesshistoryentry", &cpmcpuprocesshistorytable.Cpmcpuprocesshistoryentry[i]}
    }
    cpmcpuprocesshistorytable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cpmcpuprocesshistorytable.EntityData)
}

// CISCOPROCESSMIB_Cpmcpuprocesshistorytable_Cpmcpuprocesshistoryentry
// A historical sample of process utilization
// statistics. The entries in this table will have
// corresponding entires in the cpmCPUHistoryTable.
// The entries in this table get deleted when the entry
// associated with this entry in the cpmCPUHistoryTable 
// gets deleted.
type CISCOPROCESSMIB_Cpmcpuprocesshistorytable_Cpmcpuprocesshistoryentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // cisco_process_mib.CISCOPROCESSMIB_Cpmcputotaltable_Cpmcputotalentry_Cpmcputotalindex
    Cpmcputotalindex interface{}

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // cisco_process_mib.CISCOPROCESSMIB_Cpmcpuhistorytable_Cpmcpuhistoryentry_Cpmcpuhistoryreportid
    Cpmcpuhistoryreportid interface{}

    // This attribute is a key. An index that uniquely identifies an entry in the
    // cmpCPUProcessHistory table among those in the  same report. This index is
    // between 1 to N,  where N is the cpmCPUHistoryReportSize. The type is
    // interface{} with range: 1..4294967295.
    Cpmcpuprocesshistoryindex interface{}

    // The process Id associated with this entry. The type is interface{} with
    // range: 1..2147483647.
    Cpmcpuhistoryprocid interface{}

    // The process name associated with this entry. The type is string.
    Cpmcpuhistoryprocname interface{}

    // The time when the process was created. The process ID and the time when the
    // process was created, uniquely  identifies a process. The type is
    // interface{} with range: 0..4294967295.
    Cpmcpuhistoryproccreated interface{}

    // The percentage CPU utilization of a process at cpmCPUHistoryCreatedTime.
    // The type is interface{} with range: 0..100. Units are percent.
    Cpmcpuhistoryprocutil interface{}
}

func (cpmcpuprocesshistoryentry *CISCOPROCESSMIB_Cpmcpuprocesshistorytable_Cpmcpuprocesshistoryentry) GetEntityData() *types.CommonEntityData {
    cpmcpuprocesshistoryentry.EntityData.YFilter = cpmcpuprocesshistoryentry.YFilter
    cpmcpuprocesshistoryentry.EntityData.YangName = "cpmCPUProcessHistoryEntry"
    cpmcpuprocesshistoryentry.EntityData.BundleName = "cisco_ios_xe"
    cpmcpuprocesshistoryentry.EntityData.ParentYangName = "cpmCPUProcessHistoryTable"
    cpmcpuprocesshistoryentry.EntityData.SegmentPath = "cpmCPUProcessHistoryEntry" + "[cpmCPUTotalIndex='" + fmt.Sprintf("%v", cpmcpuprocesshistoryentry.Cpmcputotalindex) + "']" + "[cpmCPUHistoryReportId='" + fmt.Sprintf("%v", cpmcpuprocesshistoryentry.Cpmcpuhistoryreportid) + "']" + "[cpmCPUProcessHistoryIndex='" + fmt.Sprintf("%v", cpmcpuprocesshistoryentry.Cpmcpuprocesshistoryindex) + "']"
    cpmcpuprocesshistoryentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpmcpuprocesshistoryentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpmcpuprocesshistoryentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpmcpuprocesshistoryentry.EntityData.Children = make(map[string]types.YChild)
    cpmcpuprocesshistoryentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cpmcpuprocesshistoryentry.EntityData.Leafs["cpmCPUTotalIndex"] = types.YLeaf{"Cpmcputotalindex", cpmcpuprocesshistoryentry.Cpmcputotalindex}
    cpmcpuprocesshistoryentry.EntityData.Leafs["cpmCPUHistoryReportId"] = types.YLeaf{"Cpmcpuhistoryreportid", cpmcpuprocesshistoryentry.Cpmcpuhistoryreportid}
    cpmcpuprocesshistoryentry.EntityData.Leafs["cpmCPUProcessHistoryIndex"] = types.YLeaf{"Cpmcpuprocesshistoryindex", cpmcpuprocesshistoryentry.Cpmcpuprocesshistoryindex}
    cpmcpuprocesshistoryentry.EntityData.Leafs["cpmCPUHistoryProcId"] = types.YLeaf{"Cpmcpuhistoryprocid", cpmcpuprocesshistoryentry.Cpmcpuhistoryprocid}
    cpmcpuprocesshistoryentry.EntityData.Leafs["cpmCPUHistoryProcName"] = types.YLeaf{"Cpmcpuhistoryprocname", cpmcpuprocesshistoryentry.Cpmcpuhistoryprocname}
    cpmcpuprocesshistoryentry.EntityData.Leafs["cpmCPUHistoryProcCreated"] = types.YLeaf{"Cpmcpuhistoryproccreated", cpmcpuprocesshistoryentry.Cpmcpuhistoryproccreated}
    cpmcpuprocesshistoryentry.EntityData.Leafs["cpmCPUHistoryProcUtil"] = types.YLeaf{"Cpmcpuhistoryprocutil", cpmcpuprocesshistoryentry.Cpmcpuhistoryprocutil}
    return &(cpmcpuprocesshistoryentry.EntityData)
}

// CISCOPROCESSMIB_Cpmthreadtable
// This table contains generic information about
// POSIX threads in the device.
type CISCOPROCESSMIB_Cpmthreadtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry containing the general statistics of a POSIX thread. The type is
    // slice of CISCOPROCESSMIB_Cpmthreadtable_Cpmthreadentry.
    Cpmthreadentry []CISCOPROCESSMIB_Cpmthreadtable_Cpmthreadentry
}

func (cpmthreadtable *CISCOPROCESSMIB_Cpmthreadtable) GetEntityData() *types.CommonEntityData {
    cpmthreadtable.EntityData.YFilter = cpmthreadtable.YFilter
    cpmthreadtable.EntityData.YangName = "cpmThreadTable"
    cpmthreadtable.EntityData.BundleName = "cisco_ios_xe"
    cpmthreadtable.EntityData.ParentYangName = "CISCO-PROCESS-MIB"
    cpmthreadtable.EntityData.SegmentPath = "cpmThreadTable"
    cpmthreadtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpmthreadtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpmthreadtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpmthreadtable.EntityData.Children = make(map[string]types.YChild)
    cpmthreadtable.EntityData.Children["cpmThreadEntry"] = types.YChild{"Cpmthreadentry", nil}
    for i := range cpmthreadtable.Cpmthreadentry {
        cpmthreadtable.EntityData.Children[types.GetSegmentPath(&cpmthreadtable.Cpmthreadentry[i])] = types.YChild{"Cpmthreadentry", &cpmthreadtable.Cpmthreadentry[i]}
    }
    cpmthreadtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cpmthreadtable.EntityData)
}

// CISCOPROCESSMIB_Cpmthreadtable_Cpmthreadentry
// An entry containing the general statistics
// of a POSIX thread.
type CISCOPROCESSMIB_Cpmthreadtable_Cpmthreadentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // cisco_process_mib.CISCOPROCESSMIB_Cpmcputotaltable_Cpmcputotalentry_Cpmcputotalindex
    Cpmcputotalindex interface{}

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // cisco_process_mib.CISCOPROCESSMIB_Cpmprocesstable_Cpmprocessentry_Cpmprocesspid
    Cpmprocesspid interface{}

    // This attribute is a key. This object contains the thread ID. ThreadID is
    // Unique per process. The type is interface{} with range: 0..4294967295.
    Cpmthreadid interface{}

    // This object represents the name of the thread. Thread names need not be
    // unique. Hence statistics  should be analyzed against thread ID. The type is
    // string.
    Cpmthreadname interface{}

    // This object indicates the priority of a POSIX thread. The higher the
    // number, the higher the priority of the  thread over other threads. The type
    // is interface{} with range: 0..63.
    Cpmthreadpriority interface{}

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
    // to acquire a semaphore. The type is Cpmthreadstate.
    Cpmthreadstate interface{}

    // This object identifies the process on which the current thread is blocked
    // on. This points to the  cpmProcessTable of the process on which the thread 
    // in context is blocked. This is valid only to threads which are either in
    // send/reply states. For the  rest of the threads it is returned as 0.0. The
    // type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Cpmthreadblockingprocess interface{}

    // This object provides a general idea on how busy the thread in context
    // caused the processor to be. The type is interface{} with range:
    // 0..4294967295. Units are milliseconds.
    Cpmthreadcpuutilization interface{}

    // This object indicates the stack size allocated to the thread in context.
    // The type is interface{} with range: 0..4294967295. Units are bytes.
    Cpmthreadstacksize interface{}

    // This object represents the upper 32-bit of cpmThreadStackSize. This object
    // needs to be supported only when the value of cpmThreadStackSize exceeds
    // 32-bit, otherwise this object value would be set to 0. The type is
    // interface{} with range: 0..4294967295. Units are bytes.
    Cpmthreadstacksizeovrflw interface{}

    // This object indicates the stack size allocated to the thread in context.
    // This object is a 64-bit version of cpmThreadStackSize. The type is
    // interface{} with range: 0..18446744073709551615. Units are bytes.
    Cpmthreadhcstacksize interface{}
}

func (cpmthreadentry *CISCOPROCESSMIB_Cpmthreadtable_Cpmthreadentry) GetEntityData() *types.CommonEntityData {
    cpmthreadentry.EntityData.YFilter = cpmthreadentry.YFilter
    cpmthreadentry.EntityData.YangName = "cpmThreadEntry"
    cpmthreadentry.EntityData.BundleName = "cisco_ios_xe"
    cpmthreadentry.EntityData.ParentYangName = "cpmThreadTable"
    cpmthreadentry.EntityData.SegmentPath = "cpmThreadEntry" + "[cpmCPUTotalIndex='" + fmt.Sprintf("%v", cpmthreadentry.Cpmcputotalindex) + "']" + "[cpmProcessPID='" + fmt.Sprintf("%v", cpmthreadentry.Cpmprocesspid) + "']" + "[cpmThreadID='" + fmt.Sprintf("%v", cpmthreadentry.Cpmthreadid) + "']"
    cpmthreadentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpmthreadentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpmthreadentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpmthreadentry.EntityData.Children = make(map[string]types.YChild)
    cpmthreadentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cpmthreadentry.EntityData.Leafs["cpmCPUTotalIndex"] = types.YLeaf{"Cpmcputotalindex", cpmthreadentry.Cpmcputotalindex}
    cpmthreadentry.EntityData.Leafs["cpmProcessPID"] = types.YLeaf{"Cpmprocesspid", cpmthreadentry.Cpmprocesspid}
    cpmthreadentry.EntityData.Leafs["cpmThreadID"] = types.YLeaf{"Cpmthreadid", cpmthreadentry.Cpmthreadid}
    cpmthreadentry.EntityData.Leafs["cpmThreadName"] = types.YLeaf{"Cpmthreadname", cpmthreadentry.Cpmthreadname}
    cpmthreadentry.EntityData.Leafs["cpmThreadPriority"] = types.YLeaf{"Cpmthreadpriority", cpmthreadentry.Cpmthreadpriority}
    cpmthreadentry.EntityData.Leafs["cpmThreadState"] = types.YLeaf{"Cpmthreadstate", cpmthreadentry.Cpmthreadstate}
    cpmthreadentry.EntityData.Leafs["cpmThreadBlockingProcess"] = types.YLeaf{"Cpmthreadblockingprocess", cpmthreadentry.Cpmthreadblockingprocess}
    cpmthreadentry.EntityData.Leafs["cpmThreadCpuUtilization"] = types.YLeaf{"Cpmthreadcpuutilization", cpmthreadentry.Cpmthreadcpuutilization}
    cpmthreadentry.EntityData.Leafs["cpmThreadStackSize"] = types.YLeaf{"Cpmthreadstacksize", cpmthreadentry.Cpmthreadstacksize}
    cpmthreadentry.EntityData.Leafs["cpmThreadStackSizeOvrflw"] = types.YLeaf{"Cpmthreadstacksizeovrflw", cpmthreadentry.Cpmthreadstacksizeovrflw}
    cpmthreadentry.EntityData.Leafs["cpmThreadHCStackSize"] = types.YLeaf{"Cpmthreadhcstacksize", cpmthreadentry.Cpmthreadhcstacksize}
    return &(cpmthreadentry.EntityData)
}

// CISCOPROCESSMIB_Cpmthreadtable_Cpmthreadentry_Cpmthreadstate represents sem           - Waiting to acquire a semaphore.
type CISCOPROCESSMIB_Cpmthreadtable_Cpmthreadentry_Cpmthreadstate string

const (
    CISCOPROCESSMIB_Cpmthreadtable_Cpmthreadentry_Cpmthreadstate_other CISCOPROCESSMIB_Cpmthreadtable_Cpmthreadentry_Cpmthreadstate = "other"

    CISCOPROCESSMIB_Cpmthreadtable_Cpmthreadentry_Cpmthreadstate_dead CISCOPROCESSMIB_Cpmthreadtable_Cpmthreadentry_Cpmthreadstate = "dead"

    CISCOPROCESSMIB_Cpmthreadtable_Cpmthreadentry_Cpmthreadstate_running CISCOPROCESSMIB_Cpmthreadtable_Cpmthreadentry_Cpmthreadstate = "running"

    CISCOPROCESSMIB_Cpmthreadtable_Cpmthreadentry_Cpmthreadstate_ready CISCOPROCESSMIB_Cpmthreadtable_Cpmthreadentry_Cpmthreadstate = "ready"

    CISCOPROCESSMIB_Cpmthreadtable_Cpmthreadentry_Cpmthreadstate_stopped CISCOPROCESSMIB_Cpmthreadtable_Cpmthreadentry_Cpmthreadstate = "stopped"

    CISCOPROCESSMIB_Cpmthreadtable_Cpmthreadentry_Cpmthreadstate_send CISCOPROCESSMIB_Cpmthreadtable_Cpmthreadentry_Cpmthreadstate = "send"

    CISCOPROCESSMIB_Cpmthreadtable_Cpmthreadentry_Cpmthreadstate_receive CISCOPROCESSMIB_Cpmthreadtable_Cpmthreadentry_Cpmthreadstate = "receive"

    CISCOPROCESSMIB_Cpmthreadtable_Cpmthreadentry_Cpmthreadstate_reply CISCOPROCESSMIB_Cpmthreadtable_Cpmthreadentry_Cpmthreadstate = "reply"

    CISCOPROCESSMIB_Cpmthreadtable_Cpmthreadentry_Cpmthreadstate_stack CISCOPROCESSMIB_Cpmthreadtable_Cpmthreadentry_Cpmthreadstate = "stack"

    CISCOPROCESSMIB_Cpmthreadtable_Cpmthreadentry_Cpmthreadstate_waitpage CISCOPROCESSMIB_Cpmthreadtable_Cpmthreadentry_Cpmthreadstate = "waitpage"

    CISCOPROCESSMIB_Cpmthreadtable_Cpmthreadentry_Cpmthreadstate_sigsuspend CISCOPROCESSMIB_Cpmthreadtable_Cpmthreadentry_Cpmthreadstate = "sigsuspend"

    CISCOPROCESSMIB_Cpmthreadtable_Cpmthreadentry_Cpmthreadstate_sigwaitinfo CISCOPROCESSMIB_Cpmthreadtable_Cpmthreadentry_Cpmthreadstate = "sigwaitinfo"

    CISCOPROCESSMIB_Cpmthreadtable_Cpmthreadentry_Cpmthreadstate_nanosleep CISCOPROCESSMIB_Cpmthreadtable_Cpmthreadentry_Cpmthreadstate = "nanosleep"

    CISCOPROCESSMIB_Cpmthreadtable_Cpmthreadentry_Cpmthreadstate_mutex CISCOPROCESSMIB_Cpmthreadtable_Cpmthreadentry_Cpmthreadstate = "mutex"

    CISCOPROCESSMIB_Cpmthreadtable_Cpmthreadentry_Cpmthreadstate_condvar CISCOPROCESSMIB_Cpmthreadtable_Cpmthreadentry_Cpmthreadstate = "condvar"

    CISCOPROCESSMIB_Cpmthreadtable_Cpmthreadentry_Cpmthreadstate_join CISCOPROCESSMIB_Cpmthreadtable_Cpmthreadentry_Cpmthreadstate = "join"

    CISCOPROCESSMIB_Cpmthreadtable_Cpmthreadentry_Cpmthreadstate_intr CISCOPROCESSMIB_Cpmthreadtable_Cpmthreadentry_Cpmthreadstate = "intr"

    CISCOPROCESSMIB_Cpmthreadtable_Cpmthreadentry_Cpmthreadstate_sem CISCOPROCESSMIB_Cpmthreadtable_Cpmthreadentry_Cpmthreadstate = "sem"
)

// CISCOPROCESSMIB_Cpmvirtualprocesstable
// This table contains information about virtual
// processes in a virtual machine.
type CISCOPROCESSMIB_Cpmvirtualprocesstable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry containing the general statistics of a virtual process in a
    // virtual machine. The type is slice of
    // CISCOPROCESSMIB_Cpmvirtualprocesstable_Cpmvirtualprocessentry.
    Cpmvirtualprocessentry []CISCOPROCESSMIB_Cpmvirtualprocesstable_Cpmvirtualprocessentry
}

func (cpmvirtualprocesstable *CISCOPROCESSMIB_Cpmvirtualprocesstable) GetEntityData() *types.CommonEntityData {
    cpmvirtualprocesstable.EntityData.YFilter = cpmvirtualprocesstable.YFilter
    cpmvirtualprocesstable.EntityData.YangName = "cpmVirtualProcessTable"
    cpmvirtualprocesstable.EntityData.BundleName = "cisco_ios_xe"
    cpmvirtualprocesstable.EntityData.ParentYangName = "CISCO-PROCESS-MIB"
    cpmvirtualprocesstable.EntityData.SegmentPath = "cpmVirtualProcessTable"
    cpmvirtualprocesstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpmvirtualprocesstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpmvirtualprocesstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpmvirtualprocesstable.EntityData.Children = make(map[string]types.YChild)
    cpmvirtualprocesstable.EntityData.Children["cpmVirtualProcessEntry"] = types.YChild{"Cpmvirtualprocessentry", nil}
    for i := range cpmvirtualprocesstable.Cpmvirtualprocessentry {
        cpmvirtualprocesstable.EntityData.Children[types.GetSegmentPath(&cpmvirtualprocesstable.Cpmvirtualprocessentry[i])] = types.YChild{"Cpmvirtualprocessentry", &cpmvirtualprocesstable.Cpmvirtualprocessentry[i]}
    }
    cpmvirtualprocesstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cpmvirtualprocesstable.EntityData)
}

// CISCOPROCESSMIB_Cpmvirtualprocesstable_Cpmvirtualprocessentry
// An entry containing the general statistics of a
// virtual process in a virtual machine.
type CISCOPROCESSMIB_Cpmvirtualprocesstable_Cpmvirtualprocessentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // cisco_process_mib.CISCOPROCESSMIB_Cpmcputotaltable_Cpmcputotalentry_Cpmcputotalindex
    Cpmcputotalindex interface{}

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // cisco_process_mib.CISCOPROCESSMIB_Cpmprocesstable_Cpmprocessentry_Cpmprocesspid
    Cpmprocesspid interface{}

    // This attribute is a key. This object indicates the process ID of a virtual
    // process. PID is unique only inside one address space. Virtual process PID
    // should be considered along with  Parent process cpmProcessPID. The type is
    // interface{} with range: 0..4294967295.
    Cpmvirtualprocessid interface{}

    // This object indicates the name of a virtual process. If the name is longer
    // than 32 characters, it will be truncated to the first 31 characters, and a
    // `*' will be appended as the last character to imply this is a truncated
    // process name. The type is string with length: 1..32.
    Cpmvirtualprocessname interface{}

    // This indicates an estimated CPU utilization by a virtual process over the
    // last 5 seconds. The type is interface{} with range: 0..100. Units are
    // percent.
    Cpmvirtualprocessutil5Sec interface{}

    // This indicates an estimated CPU utilization by a virtual process over the
    // last one minute. The type is interface{} with range: 0..100. Units are
    // percent.
    Cpmvirtualprocessutil1Min interface{}

    // This indicates an estimated CPU utilization by a virtual process over the
    // last 5 minutes. The type is interface{} with range: 0..100. Units are
    // percent.
    Cpmvirtualprocessutil5Min interface{}

    // This object indicates the memory allocated by the virtual process inside
    // the address space of a  process running on Native OS. The type is
    // interface{} with range: 0..4294967295. Units are bytes.
    Cpmvirtualprocessmemallocated interface{}

    // This object indicates the memory freed by the virtual process inside the
    // address space of a process running  on Native OS. The type is interface{}
    // with range: 0..4294967295. Units are bytes.
    Cpmvirtualprocessmemfreed interface{}

    // The number of times a virtual process is invoked. The type is interface{}
    // with range: 0..4294967295.
    Cpmvirtualprocessinvokecount interface{}

    // The amount of CPU time a virtual process has used in microseconds. The type
    // is interface{} with range: 0..4294967295. Units are microseconds.
    Cpmvirtualprocessruntime interface{}

    // This object represents the upper 32-bit of cpmVirtualProcessMemAllocated.
    // This object  needs to be supported only when the value of
    // cpmVirtualProcessMemAllocated exceeds 32-bit, otherwise this object value
    // would be set to 0. The type is interface{} with range: 0..4294967295. Units
    // are bytes.
    Cpmvirtualprocessmemallocatedovrflw interface{}

    // This object indicates the memory allocated by the virtual process inside
    // the address space of a process running on Native OS. This object is a
    // 64-bit version of cpmVirtualProcessMemAllocated. The type is interface{}
    // with range: 0..18446744073709551615. Units are bytes.
    Cpmvirtualprocesshcmemallocated interface{}

    // This object represents the upper 32-bit of cpmVirtualProcessMemFreed. This
    // object needs to be supported only when the value of 
    // cpmVirtualProcessMemFreed exceeds 32-bit,  otherwise this object value
    // would be set to 0. The type is interface{} with range: 0..4294967295. Units
    // are bytes.
    Cpmvirtualprocessmemfreedovrflw interface{}

    // This object indicates the memory freed by the virtual process inside the
    // address space of a process running on Native OS.This object is a 64-bit
    // version of cpmVirtualProcessMemAllocated. The type is interface{} with
    // range: 0..18446744073709551615. Units are bytes.
    Cpmvirtualprocesshcmemfreed interface{}
}

func (cpmvirtualprocessentry *CISCOPROCESSMIB_Cpmvirtualprocesstable_Cpmvirtualprocessentry) GetEntityData() *types.CommonEntityData {
    cpmvirtualprocessentry.EntityData.YFilter = cpmvirtualprocessentry.YFilter
    cpmvirtualprocessentry.EntityData.YangName = "cpmVirtualProcessEntry"
    cpmvirtualprocessentry.EntityData.BundleName = "cisco_ios_xe"
    cpmvirtualprocessentry.EntityData.ParentYangName = "cpmVirtualProcessTable"
    cpmvirtualprocessentry.EntityData.SegmentPath = "cpmVirtualProcessEntry" + "[cpmCPUTotalIndex='" + fmt.Sprintf("%v", cpmvirtualprocessentry.Cpmcputotalindex) + "']" + "[cpmProcessPID='" + fmt.Sprintf("%v", cpmvirtualprocessentry.Cpmprocesspid) + "']" + "[cpmVirtualProcessID='" + fmt.Sprintf("%v", cpmvirtualprocessentry.Cpmvirtualprocessid) + "']"
    cpmvirtualprocessentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpmvirtualprocessentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpmvirtualprocessentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpmvirtualprocessentry.EntityData.Children = make(map[string]types.YChild)
    cpmvirtualprocessentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cpmvirtualprocessentry.EntityData.Leafs["cpmCPUTotalIndex"] = types.YLeaf{"Cpmcputotalindex", cpmvirtualprocessentry.Cpmcputotalindex}
    cpmvirtualprocessentry.EntityData.Leafs["cpmProcessPID"] = types.YLeaf{"Cpmprocesspid", cpmvirtualprocessentry.Cpmprocesspid}
    cpmvirtualprocessentry.EntityData.Leafs["cpmVirtualProcessID"] = types.YLeaf{"Cpmvirtualprocessid", cpmvirtualprocessentry.Cpmvirtualprocessid}
    cpmvirtualprocessentry.EntityData.Leafs["cpmVirtualProcessName"] = types.YLeaf{"Cpmvirtualprocessname", cpmvirtualprocessentry.Cpmvirtualprocessname}
    cpmvirtualprocessentry.EntityData.Leafs["cpmVirtualProcessUtil5Sec"] = types.YLeaf{"Cpmvirtualprocessutil5Sec", cpmvirtualprocessentry.Cpmvirtualprocessutil5Sec}
    cpmvirtualprocessentry.EntityData.Leafs["cpmVirtualProcessUtil1Min"] = types.YLeaf{"Cpmvirtualprocessutil1Min", cpmvirtualprocessentry.Cpmvirtualprocessutil1Min}
    cpmvirtualprocessentry.EntityData.Leafs["cpmVirtualProcessUtil5Min"] = types.YLeaf{"Cpmvirtualprocessutil5Min", cpmvirtualprocessentry.Cpmvirtualprocessutil5Min}
    cpmvirtualprocessentry.EntityData.Leafs["cpmVirtualProcessMemAllocated"] = types.YLeaf{"Cpmvirtualprocessmemallocated", cpmvirtualprocessentry.Cpmvirtualprocessmemallocated}
    cpmvirtualprocessentry.EntityData.Leafs["cpmVirtualProcessMemFreed"] = types.YLeaf{"Cpmvirtualprocessmemfreed", cpmvirtualprocessentry.Cpmvirtualprocessmemfreed}
    cpmvirtualprocessentry.EntityData.Leafs["cpmVirtualProcessInvokeCount"] = types.YLeaf{"Cpmvirtualprocessinvokecount", cpmvirtualprocessentry.Cpmvirtualprocessinvokecount}
    cpmvirtualprocessentry.EntityData.Leafs["cpmVirtualProcessRuntime"] = types.YLeaf{"Cpmvirtualprocessruntime", cpmvirtualprocessentry.Cpmvirtualprocessruntime}
    cpmvirtualprocessentry.EntityData.Leafs["cpmVirtualProcessMemAllocatedOvrflw"] = types.YLeaf{"Cpmvirtualprocessmemallocatedovrflw", cpmvirtualprocessentry.Cpmvirtualprocessmemallocatedovrflw}
    cpmvirtualprocessentry.EntityData.Leafs["cpmVirtualProcessHCMemAllocated"] = types.YLeaf{"Cpmvirtualprocesshcmemallocated", cpmvirtualprocessentry.Cpmvirtualprocesshcmemallocated}
    cpmvirtualprocessentry.EntityData.Leafs["cpmVirtualProcessMemFreedOvrflw"] = types.YLeaf{"Cpmvirtualprocessmemfreedovrflw", cpmvirtualprocessentry.Cpmvirtualprocessmemfreedovrflw}
    cpmvirtualprocessentry.EntityData.Leafs["cpmVirtualProcessHCMemFreed"] = types.YLeaf{"Cpmvirtualprocesshcmemfreed", cpmvirtualprocessentry.Cpmvirtualprocesshcmemfreed}
    return &(cpmvirtualprocessentry.EntityData)
}

