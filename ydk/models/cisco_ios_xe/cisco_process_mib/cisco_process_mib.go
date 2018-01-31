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
    parent types.Entity
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

func (cISCOPROCESSMIB *CISCOPROCESSMIB) GetFilter() yfilter.YFilter { return cISCOPROCESSMIB.YFilter }

func (cISCOPROCESSMIB *CISCOPROCESSMIB) SetFilter(yf yfilter.YFilter) { cISCOPROCESSMIB.YFilter = yf }

func (cISCOPROCESSMIB *CISCOPROCESSMIB) GetGoName(yname string) string {
    if yname == "cpmCPUHistory" { return "Cpmcpuhistory" }
    if yname == "cpmCPUTotalTable" { return "Cpmcputotaltable" }
    if yname == "cpmCoreTable" { return "Cpmcoretable" }
    if yname == "cpmProcessTable" { return "Cpmprocesstable" }
    if yname == "cpmProcessExtRevTable" { return "Cpmprocessextrevtable" }
    if yname == "cpmCPUThresholdTable" { return "Cpmcputhresholdtable" }
    if yname == "cpmCPUHistoryTable" { return "Cpmcpuhistorytable" }
    if yname == "cpmCPUProcessHistoryTable" { return "Cpmcpuprocesshistorytable" }
    if yname == "cpmThreadTable" { return "Cpmthreadtable" }
    if yname == "cpmVirtualProcessTable" { return "Cpmvirtualprocesstable" }
    return ""
}

func (cISCOPROCESSMIB *CISCOPROCESSMIB) GetSegmentPath() string {
    return "CISCO-PROCESS-MIB:CISCO-PROCESS-MIB"
}

func (cISCOPROCESSMIB *CISCOPROCESSMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cpmCPUHistory" {
        return &cISCOPROCESSMIB.Cpmcpuhistory
    }
    if childYangName == "cpmCPUTotalTable" {
        return &cISCOPROCESSMIB.Cpmcputotaltable
    }
    if childYangName == "cpmCoreTable" {
        return &cISCOPROCESSMIB.Cpmcoretable
    }
    if childYangName == "cpmProcessTable" {
        return &cISCOPROCESSMIB.Cpmprocesstable
    }
    if childYangName == "cpmProcessExtRevTable" {
        return &cISCOPROCESSMIB.Cpmprocessextrevtable
    }
    if childYangName == "cpmCPUThresholdTable" {
        return &cISCOPROCESSMIB.Cpmcputhresholdtable
    }
    if childYangName == "cpmCPUHistoryTable" {
        return &cISCOPROCESSMIB.Cpmcpuhistorytable
    }
    if childYangName == "cpmCPUProcessHistoryTable" {
        return &cISCOPROCESSMIB.Cpmcpuprocesshistorytable
    }
    if childYangName == "cpmThreadTable" {
        return &cISCOPROCESSMIB.Cpmthreadtable
    }
    if childYangName == "cpmVirtualProcessTable" {
        return &cISCOPROCESSMIB.Cpmvirtualprocesstable
    }
    return nil
}

func (cISCOPROCESSMIB *CISCOPROCESSMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cpmCPUHistory"] = &cISCOPROCESSMIB.Cpmcpuhistory
    children["cpmCPUTotalTable"] = &cISCOPROCESSMIB.Cpmcputotaltable
    children["cpmCoreTable"] = &cISCOPROCESSMIB.Cpmcoretable
    children["cpmProcessTable"] = &cISCOPROCESSMIB.Cpmprocesstable
    children["cpmProcessExtRevTable"] = &cISCOPROCESSMIB.Cpmprocessextrevtable
    children["cpmCPUThresholdTable"] = &cISCOPROCESSMIB.Cpmcputhresholdtable
    children["cpmCPUHistoryTable"] = &cISCOPROCESSMIB.Cpmcpuhistorytable
    children["cpmCPUProcessHistoryTable"] = &cISCOPROCESSMIB.Cpmcpuprocesshistorytable
    children["cpmThreadTable"] = &cISCOPROCESSMIB.Cpmthreadtable
    children["cpmVirtualProcessTable"] = &cISCOPROCESSMIB.Cpmvirtualprocesstable
    return children
}

func (cISCOPROCESSMIB *CISCOPROCESSMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOPROCESSMIB *CISCOPROCESSMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOPROCESSMIB *CISCOPROCESSMIB) GetYangName() string { return "CISCO-PROCESS-MIB" }

func (cISCOPROCESSMIB *CISCOPROCESSMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOPROCESSMIB *CISCOPROCESSMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOPROCESSMIB *CISCOPROCESSMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOPROCESSMIB *CISCOPROCESSMIB) SetParent(parent types.Entity) { cISCOPROCESSMIB.parent = parent }

func (cISCOPROCESSMIB *CISCOPROCESSMIB) GetParent() types.Entity { return cISCOPROCESSMIB.parent }

func (cISCOPROCESSMIB *CISCOPROCESSMIB) GetParentYangName() string { return "CISCO-PROCESS-MIB" }

// CISCOPROCESSMIB_Cpmcpuhistory
type CISCOPROCESSMIB_Cpmcpuhistory struct {
    parent types.Entity
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

func (cpmcpuhistory *CISCOPROCESSMIB_Cpmcpuhistory) GetFilter() yfilter.YFilter { return cpmcpuhistory.YFilter }

func (cpmcpuhistory *CISCOPROCESSMIB_Cpmcpuhistory) SetFilter(yf yfilter.YFilter) { cpmcpuhistory.YFilter = yf }

func (cpmcpuhistory *CISCOPROCESSMIB_Cpmcpuhistory) GetGoName(yname string) string {
    if yname == "cpmCPUHistoryThreshold" { return "Cpmcpuhistorythreshold" }
    if yname == "cpmCPUHistorySize" { return "Cpmcpuhistorysize" }
    return ""
}

func (cpmcpuhistory *CISCOPROCESSMIB_Cpmcpuhistory) GetSegmentPath() string {
    return "cpmCPUHistory"
}

func (cpmcpuhistory *CISCOPROCESSMIB_Cpmcpuhistory) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cpmcpuhistory *CISCOPROCESSMIB_Cpmcpuhistory) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cpmcpuhistory *CISCOPROCESSMIB_Cpmcpuhistory) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cpmCPUHistoryThreshold"] = cpmcpuhistory.Cpmcpuhistorythreshold
    leafs["cpmCPUHistorySize"] = cpmcpuhistory.Cpmcpuhistorysize
    return leafs
}

func (cpmcpuhistory *CISCOPROCESSMIB_Cpmcpuhistory) GetBundleName() string { return "cisco_ios_xe" }

func (cpmcpuhistory *CISCOPROCESSMIB_Cpmcpuhistory) GetYangName() string { return "cpmCPUHistory" }

func (cpmcpuhistory *CISCOPROCESSMIB_Cpmcpuhistory) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpmcpuhistory *CISCOPROCESSMIB_Cpmcpuhistory) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpmcpuhistory *CISCOPROCESSMIB_Cpmcpuhistory) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpmcpuhistory *CISCOPROCESSMIB_Cpmcpuhistory) SetParent(parent types.Entity) { cpmcpuhistory.parent = parent }

func (cpmcpuhistory *CISCOPROCESSMIB_Cpmcpuhistory) GetParent() types.Entity { return cpmcpuhistory.parent }

func (cpmcpuhistory *CISCOPROCESSMIB_Cpmcpuhistory) GetParentYangName() string { return "CISCO-PROCESS-MIB" }

// CISCOPROCESSMIB_Cpmcputotaltable
// A table of overall CPU statistics.
type CISCOPROCESSMIB_Cpmcputotaltable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Overall information about the CPU load. Entries in this table come and go
    // as CPUs are added and removed from the system. The type is slice of
    // CISCOPROCESSMIB_Cpmcputotaltable_Cpmcputotalentry.
    Cpmcputotalentry []CISCOPROCESSMIB_Cpmcputotaltable_Cpmcputotalentry
}

func (cpmcputotaltable *CISCOPROCESSMIB_Cpmcputotaltable) GetFilter() yfilter.YFilter { return cpmcputotaltable.YFilter }

func (cpmcputotaltable *CISCOPROCESSMIB_Cpmcputotaltable) SetFilter(yf yfilter.YFilter) { cpmcputotaltable.YFilter = yf }

func (cpmcputotaltable *CISCOPROCESSMIB_Cpmcputotaltable) GetGoName(yname string) string {
    if yname == "cpmCPUTotalEntry" { return "Cpmcputotalentry" }
    return ""
}

func (cpmcputotaltable *CISCOPROCESSMIB_Cpmcputotaltable) GetSegmentPath() string {
    return "cpmCPUTotalTable"
}

func (cpmcputotaltable *CISCOPROCESSMIB_Cpmcputotaltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cpmCPUTotalEntry" {
        for _, c := range cpmcputotaltable.Cpmcputotalentry {
            if cpmcputotaltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOPROCESSMIB_Cpmcputotaltable_Cpmcputotalentry{}
        cpmcputotaltable.Cpmcputotalentry = append(cpmcputotaltable.Cpmcputotalentry, child)
        return &cpmcputotaltable.Cpmcputotalentry[len(cpmcputotaltable.Cpmcputotalentry)-1]
    }
    return nil
}

func (cpmcputotaltable *CISCOPROCESSMIB_Cpmcputotaltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cpmcputotaltable.Cpmcputotalentry {
        children[cpmcputotaltable.Cpmcputotalentry[i].GetSegmentPath()] = &cpmcputotaltable.Cpmcputotalentry[i]
    }
    return children
}

func (cpmcputotaltable *CISCOPROCESSMIB_Cpmcputotaltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cpmcputotaltable *CISCOPROCESSMIB_Cpmcputotaltable) GetBundleName() string { return "cisco_ios_xe" }

func (cpmcputotaltable *CISCOPROCESSMIB_Cpmcputotaltable) GetYangName() string { return "cpmCPUTotalTable" }

func (cpmcputotaltable *CISCOPROCESSMIB_Cpmcputotaltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpmcputotaltable *CISCOPROCESSMIB_Cpmcputotaltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpmcputotaltable *CISCOPROCESSMIB_Cpmcputotaltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpmcputotaltable *CISCOPROCESSMIB_Cpmcputotaltable) SetParent(parent types.Entity) { cpmcputotaltable.parent = parent }

func (cpmcputotaltable *CISCOPROCESSMIB_Cpmcputotaltable) GetParent() types.Entity { return cpmcputotaltable.parent }

func (cpmcputotaltable *CISCOPROCESSMIB_Cpmcputotaltable) GetParentYangName() string { return "CISCO-PROCESS-MIB" }

// CISCOPROCESSMIB_Cpmcputotaltable_Cpmcputotalentry
// Overall information about the CPU load. Entries in this
// table come and go as CPUs are added and removed from the
// system.
type CISCOPROCESSMIB_Cpmcputotaltable_Cpmcputotalentry struct {
    parent types.Entity
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

func (cpmcputotalentry *CISCOPROCESSMIB_Cpmcputotaltable_Cpmcputotalentry) GetFilter() yfilter.YFilter { return cpmcputotalentry.YFilter }

func (cpmcputotalentry *CISCOPROCESSMIB_Cpmcputotaltable_Cpmcputotalentry) SetFilter(yf yfilter.YFilter) { cpmcputotalentry.YFilter = yf }

func (cpmcputotalentry *CISCOPROCESSMIB_Cpmcputotaltable_Cpmcputotalentry) GetGoName(yname string) string {
    if yname == "cpmCPUTotalIndex" { return "Cpmcputotalindex" }
    if yname == "cpmCPUTotalPhysicalIndex" { return "Cpmcputotalphysicalindex" }
    if yname == "cpmCPUTotal5sec" { return "Cpmcputotal5Sec" }
    if yname == "cpmCPUTotal1min" { return "Cpmcputotal1Min" }
    if yname == "cpmCPUTotal5min" { return "Cpmcputotal5Min" }
    if yname == "cpmCPUTotal5secRev" { return "Cpmcputotal5Secrev" }
    if yname == "cpmCPUTotal1minRev" { return "Cpmcputotal1Minrev" }
    if yname == "cpmCPUTotal5minRev" { return "Cpmcputotal5Minrev" }
    if yname == "cpmCPUMonInterval" { return "Cpmcpumoninterval" }
    if yname == "cpmCPUTotalMonIntervalValue" { return "Cpmcputotalmonintervalvalue" }
    if yname == "cpmCPUInterruptMonIntervalValue" { return "Cpmcpuinterruptmonintervalvalue" }
    if yname == "cpmCPUMemoryUsed" { return "Cpmcpumemoryused" }
    if yname == "cpmCPUMemoryFree" { return "Cpmcpumemoryfree" }
    if yname == "cpmCPUMemoryKernelReserved" { return "Cpmcpumemorykernelreserved" }
    if yname == "cpmCPUMemoryLowest" { return "Cpmcpumemorylowest" }
    if yname == "cpmCPUMemoryUsedOvrflw" { return "Cpmcpumemoryusedovrflw" }
    if yname == "cpmCPUMemoryHCUsed" { return "Cpmcpumemoryhcused" }
    if yname == "cpmCPUMemoryFreeOvrflw" { return "Cpmcpumemoryfreeovrflw" }
    if yname == "cpmCPUMemoryHCFree" { return "Cpmcpumemoryhcfree" }
    if yname == "cpmCPUMemoryKernelReservedOvrflw" { return "Cpmcpumemorykernelreservedovrflw" }
    if yname == "cpmCPUMemoryHCKernelReserved" { return "Cpmcpumemoryhckernelreserved" }
    if yname == "cpmCPUMemoryLowestOvrflw" { return "Cpmcpumemorylowestovrflw" }
    if yname == "cpmCPUMemoryHCLowest" { return "Cpmcpumemoryhclowest" }
    if yname == "cpmCPULoadAvg1min" { return "Cpmcpuloadavg1Min" }
    if yname == "cpmCPULoadAvg5min" { return "Cpmcpuloadavg5Min" }
    if yname == "cpmCPULoadAvg15min" { return "Cpmcpuloadavg15Min" }
    if yname == "cpmCPUMemoryCommitted" { return "Cpmcpumemorycommitted" }
    if yname == "cpmCPUMemoryCommittedOvrflw" { return "Cpmcpumemorycommittedovrflw" }
    if yname == "cpmCPUMemoryHCCommitted" { return "Cpmcpumemoryhccommitted" }
    return ""
}

func (cpmcputotalentry *CISCOPROCESSMIB_Cpmcputotaltable_Cpmcputotalentry) GetSegmentPath() string {
    return "cpmCPUTotalEntry" + "[cpmCPUTotalIndex='" + fmt.Sprintf("%v", cpmcputotalentry.Cpmcputotalindex) + "']"
}

func (cpmcputotalentry *CISCOPROCESSMIB_Cpmcputotaltable_Cpmcputotalentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cpmcputotalentry *CISCOPROCESSMIB_Cpmcputotaltable_Cpmcputotalentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cpmcputotalentry *CISCOPROCESSMIB_Cpmcputotaltable_Cpmcputotalentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cpmCPUTotalIndex"] = cpmcputotalentry.Cpmcputotalindex
    leafs["cpmCPUTotalPhysicalIndex"] = cpmcputotalentry.Cpmcputotalphysicalindex
    leafs["cpmCPUTotal5sec"] = cpmcputotalentry.Cpmcputotal5Sec
    leafs["cpmCPUTotal1min"] = cpmcputotalentry.Cpmcputotal1Min
    leafs["cpmCPUTotal5min"] = cpmcputotalentry.Cpmcputotal5Min
    leafs["cpmCPUTotal5secRev"] = cpmcputotalentry.Cpmcputotal5Secrev
    leafs["cpmCPUTotal1minRev"] = cpmcputotalentry.Cpmcputotal1Minrev
    leafs["cpmCPUTotal5minRev"] = cpmcputotalentry.Cpmcputotal5Minrev
    leafs["cpmCPUMonInterval"] = cpmcputotalentry.Cpmcpumoninterval
    leafs["cpmCPUTotalMonIntervalValue"] = cpmcputotalentry.Cpmcputotalmonintervalvalue
    leafs["cpmCPUInterruptMonIntervalValue"] = cpmcputotalentry.Cpmcpuinterruptmonintervalvalue
    leafs["cpmCPUMemoryUsed"] = cpmcputotalentry.Cpmcpumemoryused
    leafs["cpmCPUMemoryFree"] = cpmcputotalentry.Cpmcpumemoryfree
    leafs["cpmCPUMemoryKernelReserved"] = cpmcputotalentry.Cpmcpumemorykernelreserved
    leafs["cpmCPUMemoryLowest"] = cpmcputotalentry.Cpmcpumemorylowest
    leafs["cpmCPUMemoryUsedOvrflw"] = cpmcputotalentry.Cpmcpumemoryusedovrflw
    leafs["cpmCPUMemoryHCUsed"] = cpmcputotalentry.Cpmcpumemoryhcused
    leafs["cpmCPUMemoryFreeOvrflw"] = cpmcputotalentry.Cpmcpumemoryfreeovrflw
    leafs["cpmCPUMemoryHCFree"] = cpmcputotalentry.Cpmcpumemoryhcfree
    leafs["cpmCPUMemoryKernelReservedOvrflw"] = cpmcputotalentry.Cpmcpumemorykernelreservedovrflw
    leafs["cpmCPUMemoryHCKernelReserved"] = cpmcputotalentry.Cpmcpumemoryhckernelreserved
    leafs["cpmCPUMemoryLowestOvrflw"] = cpmcputotalentry.Cpmcpumemorylowestovrflw
    leafs["cpmCPUMemoryHCLowest"] = cpmcputotalentry.Cpmcpumemoryhclowest
    leafs["cpmCPULoadAvg1min"] = cpmcputotalentry.Cpmcpuloadavg1Min
    leafs["cpmCPULoadAvg5min"] = cpmcputotalentry.Cpmcpuloadavg5Min
    leafs["cpmCPULoadAvg15min"] = cpmcputotalentry.Cpmcpuloadavg15Min
    leafs["cpmCPUMemoryCommitted"] = cpmcputotalentry.Cpmcpumemorycommitted
    leafs["cpmCPUMemoryCommittedOvrflw"] = cpmcputotalentry.Cpmcpumemorycommittedovrflw
    leafs["cpmCPUMemoryHCCommitted"] = cpmcputotalentry.Cpmcpumemoryhccommitted
    return leafs
}

func (cpmcputotalentry *CISCOPROCESSMIB_Cpmcputotaltable_Cpmcputotalentry) GetBundleName() string { return "cisco_ios_xe" }

func (cpmcputotalentry *CISCOPROCESSMIB_Cpmcputotaltable_Cpmcputotalentry) GetYangName() string { return "cpmCPUTotalEntry" }

func (cpmcputotalentry *CISCOPROCESSMIB_Cpmcputotaltable_Cpmcputotalentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpmcputotalentry *CISCOPROCESSMIB_Cpmcputotaltable_Cpmcputotalentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpmcputotalentry *CISCOPROCESSMIB_Cpmcputotaltable_Cpmcputotalentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpmcputotalentry *CISCOPROCESSMIB_Cpmcputotaltable_Cpmcputotalentry) SetParent(parent types.Entity) { cpmcputotalentry.parent = parent }

func (cpmcputotalentry *CISCOPROCESSMIB_Cpmcputotaltable_Cpmcputotalentry) GetParent() types.Entity { return cpmcputotalentry.parent }

func (cpmcputotalentry *CISCOPROCESSMIB_Cpmcputotaltable_Cpmcputotalentry) GetParentYangName() string { return "cpmCPUTotalTable" }

// CISCOPROCESSMIB_Cpmcoretable
// A table of per-Core statistics.
type CISCOPROCESSMIB_Cpmcoretable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Overall information about the Core load. Entries in this table could come
    // and go as Cores go online or offline. The type is slice of
    // CISCOPROCESSMIB_Cpmcoretable_Cpmcoreentry.
    Cpmcoreentry []CISCOPROCESSMIB_Cpmcoretable_Cpmcoreentry
}

func (cpmcoretable *CISCOPROCESSMIB_Cpmcoretable) GetFilter() yfilter.YFilter { return cpmcoretable.YFilter }

func (cpmcoretable *CISCOPROCESSMIB_Cpmcoretable) SetFilter(yf yfilter.YFilter) { cpmcoretable.YFilter = yf }

func (cpmcoretable *CISCOPROCESSMIB_Cpmcoretable) GetGoName(yname string) string {
    if yname == "cpmCoreEntry" { return "Cpmcoreentry" }
    return ""
}

func (cpmcoretable *CISCOPROCESSMIB_Cpmcoretable) GetSegmentPath() string {
    return "cpmCoreTable"
}

func (cpmcoretable *CISCOPROCESSMIB_Cpmcoretable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cpmCoreEntry" {
        for _, c := range cpmcoretable.Cpmcoreentry {
            if cpmcoretable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOPROCESSMIB_Cpmcoretable_Cpmcoreentry{}
        cpmcoretable.Cpmcoreentry = append(cpmcoretable.Cpmcoreentry, child)
        return &cpmcoretable.Cpmcoreentry[len(cpmcoretable.Cpmcoreentry)-1]
    }
    return nil
}

func (cpmcoretable *CISCOPROCESSMIB_Cpmcoretable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cpmcoretable.Cpmcoreentry {
        children[cpmcoretable.Cpmcoreentry[i].GetSegmentPath()] = &cpmcoretable.Cpmcoreentry[i]
    }
    return children
}

func (cpmcoretable *CISCOPROCESSMIB_Cpmcoretable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cpmcoretable *CISCOPROCESSMIB_Cpmcoretable) GetBundleName() string { return "cisco_ios_xe" }

func (cpmcoretable *CISCOPROCESSMIB_Cpmcoretable) GetYangName() string { return "cpmCoreTable" }

func (cpmcoretable *CISCOPROCESSMIB_Cpmcoretable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpmcoretable *CISCOPROCESSMIB_Cpmcoretable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpmcoretable *CISCOPROCESSMIB_Cpmcoretable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpmcoretable *CISCOPROCESSMIB_Cpmcoretable) SetParent(parent types.Entity) { cpmcoretable.parent = parent }

func (cpmcoretable *CISCOPROCESSMIB_Cpmcoretable) GetParent() types.Entity { return cpmcoretable.parent }

func (cpmcoretable *CISCOPROCESSMIB_Cpmcoretable) GetParentYangName() string { return "CISCO-PROCESS-MIB" }

// CISCOPROCESSMIB_Cpmcoretable_Cpmcoreentry
// Overall information about the Core load. Entries in this
// table could come and go as Cores go online or offline.
type CISCOPROCESSMIB_Cpmcoretable_Cpmcoreentry struct {
    parent types.Entity
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

func (cpmcoreentry *CISCOPROCESSMIB_Cpmcoretable_Cpmcoreentry) GetFilter() yfilter.YFilter { return cpmcoreentry.YFilter }

func (cpmcoreentry *CISCOPROCESSMIB_Cpmcoretable_Cpmcoreentry) SetFilter(yf yfilter.YFilter) { cpmcoreentry.YFilter = yf }

func (cpmcoreentry *CISCOPROCESSMIB_Cpmcoretable_Cpmcoreentry) GetGoName(yname string) string {
    if yname == "cpmCPUTotalIndex" { return "Cpmcputotalindex" }
    if yname == "cpmCoreIndex" { return "Cpmcoreindex" }
    if yname == "cpmCorePhysicalIndex" { return "Cpmcorephysicalindex" }
    if yname == "cpmCore5sec" { return "Cpmcore5Sec" }
    if yname == "cpmCore1min" { return "Cpmcore1Min" }
    if yname == "cpmCore5min" { return "Cpmcore5Min" }
    if yname == "cpmCoreLoadAvg1min" { return "Cpmcoreloadavg1Min" }
    if yname == "cpmCoreLoadAvg5min" { return "Cpmcoreloadavg5Min" }
    if yname == "cpmCoreLoadAvg15min" { return "Cpmcoreloadavg15Min" }
    return ""
}

func (cpmcoreentry *CISCOPROCESSMIB_Cpmcoretable_Cpmcoreentry) GetSegmentPath() string {
    return "cpmCoreEntry" + "[cpmCPUTotalIndex='" + fmt.Sprintf("%v", cpmcoreentry.Cpmcputotalindex) + "']" + "[cpmCoreIndex='" + fmt.Sprintf("%v", cpmcoreentry.Cpmcoreindex) + "']"
}

func (cpmcoreentry *CISCOPROCESSMIB_Cpmcoretable_Cpmcoreentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cpmcoreentry *CISCOPROCESSMIB_Cpmcoretable_Cpmcoreentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cpmcoreentry *CISCOPROCESSMIB_Cpmcoretable_Cpmcoreentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cpmCPUTotalIndex"] = cpmcoreentry.Cpmcputotalindex
    leafs["cpmCoreIndex"] = cpmcoreentry.Cpmcoreindex
    leafs["cpmCorePhysicalIndex"] = cpmcoreentry.Cpmcorephysicalindex
    leafs["cpmCore5sec"] = cpmcoreentry.Cpmcore5Sec
    leafs["cpmCore1min"] = cpmcoreentry.Cpmcore1Min
    leafs["cpmCore5min"] = cpmcoreentry.Cpmcore5Min
    leafs["cpmCoreLoadAvg1min"] = cpmcoreentry.Cpmcoreloadavg1Min
    leafs["cpmCoreLoadAvg5min"] = cpmcoreentry.Cpmcoreloadavg5Min
    leafs["cpmCoreLoadAvg15min"] = cpmcoreentry.Cpmcoreloadavg15Min
    return leafs
}

func (cpmcoreentry *CISCOPROCESSMIB_Cpmcoretable_Cpmcoreentry) GetBundleName() string { return "cisco_ios_xe" }

func (cpmcoreentry *CISCOPROCESSMIB_Cpmcoretable_Cpmcoreentry) GetYangName() string { return "cpmCoreEntry" }

func (cpmcoreentry *CISCOPROCESSMIB_Cpmcoretable_Cpmcoreentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpmcoreentry *CISCOPROCESSMIB_Cpmcoretable_Cpmcoreentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpmcoreentry *CISCOPROCESSMIB_Cpmcoretable_Cpmcoreentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpmcoreentry *CISCOPROCESSMIB_Cpmcoretable_Cpmcoreentry) SetParent(parent types.Entity) { cpmcoreentry.parent = parent }

func (cpmcoreentry *CISCOPROCESSMIB_Cpmcoretable_Cpmcoreentry) GetParent() types.Entity { return cpmcoreentry.parent }

func (cpmcoreentry *CISCOPROCESSMIB_Cpmcoretable_Cpmcoreentry) GetParentYangName() string { return "cpmCoreTable" }

// CISCOPROCESSMIB_Cpmprocesstable
// A table of generic information on all active
// processes on this device.
type CISCOPROCESSMIB_Cpmprocesstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Generic information about an active process on this device. Entries in this
    // table come and go as processes are  created and destroyed by the device.
    // The type is slice of CISCOPROCESSMIB_Cpmprocesstable_Cpmprocessentry.
    Cpmprocessentry []CISCOPROCESSMIB_Cpmprocesstable_Cpmprocessentry
}

func (cpmprocesstable *CISCOPROCESSMIB_Cpmprocesstable) GetFilter() yfilter.YFilter { return cpmprocesstable.YFilter }

func (cpmprocesstable *CISCOPROCESSMIB_Cpmprocesstable) SetFilter(yf yfilter.YFilter) { cpmprocesstable.YFilter = yf }

func (cpmprocesstable *CISCOPROCESSMIB_Cpmprocesstable) GetGoName(yname string) string {
    if yname == "cpmProcessEntry" { return "Cpmprocessentry" }
    return ""
}

func (cpmprocesstable *CISCOPROCESSMIB_Cpmprocesstable) GetSegmentPath() string {
    return "cpmProcessTable"
}

func (cpmprocesstable *CISCOPROCESSMIB_Cpmprocesstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cpmProcessEntry" {
        for _, c := range cpmprocesstable.Cpmprocessentry {
            if cpmprocesstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOPROCESSMIB_Cpmprocesstable_Cpmprocessentry{}
        cpmprocesstable.Cpmprocessentry = append(cpmprocesstable.Cpmprocessentry, child)
        return &cpmprocesstable.Cpmprocessentry[len(cpmprocesstable.Cpmprocessentry)-1]
    }
    return nil
}

func (cpmprocesstable *CISCOPROCESSMIB_Cpmprocesstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cpmprocesstable.Cpmprocessentry {
        children[cpmprocesstable.Cpmprocessentry[i].GetSegmentPath()] = &cpmprocesstable.Cpmprocessentry[i]
    }
    return children
}

func (cpmprocesstable *CISCOPROCESSMIB_Cpmprocesstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cpmprocesstable *CISCOPROCESSMIB_Cpmprocesstable) GetBundleName() string { return "cisco_ios_xe" }

func (cpmprocesstable *CISCOPROCESSMIB_Cpmprocesstable) GetYangName() string { return "cpmProcessTable" }

func (cpmprocesstable *CISCOPROCESSMIB_Cpmprocesstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpmprocesstable *CISCOPROCESSMIB_Cpmprocesstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpmprocesstable *CISCOPROCESSMIB_Cpmprocesstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpmprocesstable *CISCOPROCESSMIB_Cpmprocesstable) SetParent(parent types.Entity) { cpmprocesstable.parent = parent }

func (cpmprocesstable *CISCOPROCESSMIB_Cpmprocesstable) GetParent() types.Entity { return cpmprocesstable.parent }

func (cpmprocesstable *CISCOPROCESSMIB_Cpmprocesstable) GetParentYangName() string { return "CISCO-PROCESS-MIB" }

// CISCOPROCESSMIB_Cpmprocesstable_Cpmprocessentry
// Generic information about an active process on this
// device. Entries in this table come and go as processes are 
// created and destroyed by the device.
type CISCOPROCESSMIB_Cpmprocesstable_Cpmprocessentry struct {
    parent types.Entity
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

func (cpmprocessentry *CISCOPROCESSMIB_Cpmprocesstable_Cpmprocessentry) GetFilter() yfilter.YFilter { return cpmprocessentry.YFilter }

func (cpmprocessentry *CISCOPROCESSMIB_Cpmprocesstable_Cpmprocessentry) SetFilter(yf yfilter.YFilter) { cpmprocessentry.YFilter = yf }

func (cpmprocessentry *CISCOPROCESSMIB_Cpmprocesstable_Cpmprocessentry) GetGoName(yname string) string {
    if yname == "cpmCPUTotalIndex" { return "Cpmcputotalindex" }
    if yname == "cpmProcessPID" { return "Cpmprocesspid" }
    if yname == "cpmProcessName" { return "Cpmprocessname" }
    if yname == "cpmProcessuSecs" { return "Cpmprocessusecs" }
    if yname == "cpmProcessTimeCreated" { return "Cpmprocesstimecreated" }
    if yname == "cpmProcessAverageUSecs" { return "Cpmprocessaverageusecs" }
    if yname == "cpmProcExtMemAllocated" { return "Cpmprocextmemallocated" }
    if yname == "cpmProcExtMemFreed" { return "Cpmprocextmemfreed" }
    if yname == "cpmProcExtInvoked" { return "Cpmprocextinvoked" }
    if yname == "cpmProcExtRuntime" { return "Cpmprocextruntime" }
    if yname == "cpmProcExtUtil5Sec" { return "Cpmprocextutil5Sec" }
    if yname == "cpmProcExtUtil1Min" { return "Cpmprocextutil1Min" }
    if yname == "cpmProcExtUtil5Min" { return "Cpmprocextutil5Min" }
    if yname == "cpmProcExtPriority" { return "Cpmprocextpriority" }
    return ""
}

func (cpmprocessentry *CISCOPROCESSMIB_Cpmprocesstable_Cpmprocessentry) GetSegmentPath() string {
    return "cpmProcessEntry" + "[cpmCPUTotalIndex='" + fmt.Sprintf("%v", cpmprocessentry.Cpmcputotalindex) + "']" + "[cpmProcessPID='" + fmt.Sprintf("%v", cpmprocessentry.Cpmprocesspid) + "']"
}

func (cpmprocessentry *CISCOPROCESSMIB_Cpmprocesstable_Cpmprocessentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cpmprocessentry *CISCOPROCESSMIB_Cpmprocesstable_Cpmprocessentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cpmprocessentry *CISCOPROCESSMIB_Cpmprocesstable_Cpmprocessentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cpmCPUTotalIndex"] = cpmprocessentry.Cpmcputotalindex
    leafs["cpmProcessPID"] = cpmprocessentry.Cpmprocesspid
    leafs["cpmProcessName"] = cpmprocessentry.Cpmprocessname
    leafs["cpmProcessuSecs"] = cpmprocessentry.Cpmprocessusecs
    leafs["cpmProcessTimeCreated"] = cpmprocessentry.Cpmprocesstimecreated
    leafs["cpmProcessAverageUSecs"] = cpmprocessentry.Cpmprocessaverageusecs
    leafs["cpmProcExtMemAllocated"] = cpmprocessentry.Cpmprocextmemallocated
    leafs["cpmProcExtMemFreed"] = cpmprocessentry.Cpmprocextmemfreed
    leafs["cpmProcExtInvoked"] = cpmprocessentry.Cpmprocextinvoked
    leafs["cpmProcExtRuntime"] = cpmprocessentry.Cpmprocextruntime
    leafs["cpmProcExtUtil5Sec"] = cpmprocessentry.Cpmprocextutil5Sec
    leafs["cpmProcExtUtil1Min"] = cpmprocessentry.Cpmprocextutil1Min
    leafs["cpmProcExtUtil5Min"] = cpmprocessentry.Cpmprocextutil5Min
    leafs["cpmProcExtPriority"] = cpmprocessentry.Cpmprocextpriority
    return leafs
}

func (cpmprocessentry *CISCOPROCESSMIB_Cpmprocesstable_Cpmprocessentry) GetBundleName() string { return "cisco_ios_xe" }

func (cpmprocessentry *CISCOPROCESSMIB_Cpmprocesstable_Cpmprocessentry) GetYangName() string { return "cpmProcessEntry" }

func (cpmprocessentry *CISCOPROCESSMIB_Cpmprocesstable_Cpmprocessentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpmprocessentry *CISCOPROCESSMIB_Cpmprocesstable_Cpmprocessentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpmprocessentry *CISCOPROCESSMIB_Cpmprocesstable_Cpmprocessentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpmprocessentry *CISCOPROCESSMIB_Cpmprocesstable_Cpmprocessentry) SetParent(parent types.Entity) { cpmprocessentry.parent = parent }

func (cpmprocessentry *CISCOPROCESSMIB_Cpmprocesstable_Cpmprocessentry) GetParent() types.Entity { return cpmprocessentry.parent }

func (cpmprocessentry *CISCOPROCESSMIB_Cpmprocesstable_Cpmprocessentry) GetParentYangName() string { return "cpmProcessTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry containing additional information for a particular process. This
    // object deprecates  cpmProcessExtEntry. The type is slice of
    // CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry.
    Cpmprocessextreventry []CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry
}

func (cpmprocessextrevtable *CISCOPROCESSMIB_Cpmprocessextrevtable) GetFilter() yfilter.YFilter { return cpmprocessextrevtable.YFilter }

func (cpmprocessextrevtable *CISCOPROCESSMIB_Cpmprocessextrevtable) SetFilter(yf yfilter.YFilter) { cpmprocessextrevtable.YFilter = yf }

func (cpmprocessextrevtable *CISCOPROCESSMIB_Cpmprocessextrevtable) GetGoName(yname string) string {
    if yname == "cpmProcessExtRevEntry" { return "Cpmprocessextreventry" }
    return ""
}

func (cpmprocessextrevtable *CISCOPROCESSMIB_Cpmprocessextrevtable) GetSegmentPath() string {
    return "cpmProcessExtRevTable"
}

func (cpmprocessextrevtable *CISCOPROCESSMIB_Cpmprocessextrevtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cpmProcessExtRevEntry" {
        for _, c := range cpmprocessextrevtable.Cpmprocessextreventry {
            if cpmprocessextrevtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry{}
        cpmprocessextrevtable.Cpmprocessextreventry = append(cpmprocessextrevtable.Cpmprocessextreventry, child)
        return &cpmprocessextrevtable.Cpmprocessextreventry[len(cpmprocessextrevtable.Cpmprocessextreventry)-1]
    }
    return nil
}

func (cpmprocessextrevtable *CISCOPROCESSMIB_Cpmprocessextrevtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cpmprocessextrevtable.Cpmprocessextreventry {
        children[cpmprocessextrevtable.Cpmprocessextreventry[i].GetSegmentPath()] = &cpmprocessextrevtable.Cpmprocessextreventry[i]
    }
    return children
}

func (cpmprocessextrevtable *CISCOPROCESSMIB_Cpmprocessextrevtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cpmprocessextrevtable *CISCOPROCESSMIB_Cpmprocessextrevtable) GetBundleName() string { return "cisco_ios_xe" }

func (cpmprocessextrevtable *CISCOPROCESSMIB_Cpmprocessextrevtable) GetYangName() string { return "cpmProcessExtRevTable" }

func (cpmprocessextrevtable *CISCOPROCESSMIB_Cpmprocessextrevtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpmprocessextrevtable *CISCOPROCESSMIB_Cpmprocessextrevtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpmprocessextrevtable *CISCOPROCESSMIB_Cpmprocessextrevtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpmprocessextrevtable *CISCOPROCESSMIB_Cpmprocessextrevtable) SetParent(parent types.Entity) { cpmprocessextrevtable.parent = parent }

func (cpmprocessextrevtable *CISCOPROCESSMIB_Cpmprocessextrevtable) GetParent() types.Entity { return cpmprocessextrevtable.parent }

func (cpmprocessextrevtable *CISCOPROCESSMIB_Cpmprocessextrevtable) GetParentYangName() string { return "CISCO-PROCESS-MIB" }

// CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry
// An entry containing additional information for
// a particular process. This object deprecates 
// cpmProcessExtEntry.
type CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry struct {
    parent types.Entity
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

func (cpmprocessextreventry *CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry) GetFilter() yfilter.YFilter { return cpmprocessextreventry.YFilter }

func (cpmprocessextreventry *CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry) SetFilter(yf yfilter.YFilter) { cpmprocessextreventry.YFilter = yf }

func (cpmprocessextreventry *CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry) GetGoName(yname string) string {
    if yname == "cpmCPUTotalIndex" { return "Cpmcputotalindex" }
    if yname == "cpmProcessPID" { return "Cpmprocesspid" }
    if yname == "cpmProcExtMemAllocatedRev" { return "Cpmprocextmemallocatedrev" }
    if yname == "cpmProcExtMemFreedRev" { return "Cpmprocextmemfreedrev" }
    if yname == "cpmProcExtInvokedRev" { return "Cpmprocextinvokedrev" }
    if yname == "cpmProcExtRuntimeRev" { return "Cpmprocextruntimerev" }
    if yname == "cpmProcExtUtil5SecRev" { return "Cpmprocextutil5Secrev" }
    if yname == "cpmProcExtUtil1MinRev" { return "Cpmprocextutil1Minrev" }
    if yname == "cpmProcExtUtil5MinRev" { return "Cpmprocextutil5Minrev" }
    if yname == "cpmProcExtPriorityRev" { return "Cpmprocextpriorityrev" }
    if yname == "cpmProcessType" { return "Cpmprocesstype" }
    if yname == "cpmProcessRespawn" { return "Cpmprocessrespawn" }
    if yname == "cpmProcessRespawnCount" { return "Cpmprocessrespawncount" }
    if yname == "cpmProcessRespawnAfterLastPatch" { return "Cpmprocessrespawnafterlastpatch" }
    if yname == "cpmProcessMemoryCore" { return "Cpmprocessmemorycore" }
    if yname == "cpmProcessLastRestartUser" { return "Cpmprocesslastrestartuser" }
    if yname == "cpmProcessTextSegmentSize" { return "Cpmprocesstextsegmentsize" }
    if yname == "cpmProcessDataSegmentSize" { return "Cpmprocessdatasegmentsize" }
    if yname == "cpmProcessStackSize" { return "Cpmprocessstacksize" }
    if yname == "cpmProcessDynamicMemorySize" { return "Cpmprocessdynamicmemorysize" }
    if yname == "cpmProcExtMemAllocatedRevOvrflw" { return "Cpmprocextmemallocatedrevovrflw" }
    if yname == "cpmProcExtHCMemAllocatedRev" { return "Cpmprocexthcmemallocatedrev" }
    if yname == "cpmProcExtMemFreedRevOvrflw" { return "Cpmprocextmemfreedrevovrflw" }
    if yname == "cpmProcExtHCMemFreedRev" { return "Cpmprocexthcmemfreedrev" }
    if yname == "cpmProcessTextSegmentSizeOvrflw" { return "Cpmprocesstextsegmentsizeovrflw" }
    if yname == "cpmProcessHCTextSegmentSize" { return "Cpmprocesshctextsegmentsize" }
    if yname == "cpmProcessDataSegmentSizeOvrflw" { return "Cpmprocessdatasegmentsizeovrflw" }
    if yname == "cpmProcessHCDataSegmentSize" { return "Cpmprocesshcdatasegmentsize" }
    if yname == "cpmProcessStackSizeOvrflw" { return "Cpmprocessstacksizeovrflw" }
    if yname == "cpmProcessHCStackSize" { return "Cpmprocesshcstacksize" }
    if yname == "cpmProcessDynamicMemorySizeOvrflw" { return "Cpmprocessdynamicmemorysizeovrflw" }
    if yname == "cpmProcessHCDynamicMemorySize" { return "Cpmprocesshcdynamicmemorysize" }
    return ""
}

func (cpmprocessextreventry *CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry) GetSegmentPath() string {
    return "cpmProcessExtRevEntry" + "[cpmCPUTotalIndex='" + fmt.Sprintf("%v", cpmprocessextreventry.Cpmcputotalindex) + "']" + "[cpmProcessPID='" + fmt.Sprintf("%v", cpmprocessextreventry.Cpmprocesspid) + "']"
}

func (cpmprocessextreventry *CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cpmprocessextreventry *CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cpmprocessextreventry *CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cpmCPUTotalIndex"] = cpmprocessextreventry.Cpmcputotalindex
    leafs["cpmProcessPID"] = cpmprocessextreventry.Cpmprocesspid
    leafs["cpmProcExtMemAllocatedRev"] = cpmprocessextreventry.Cpmprocextmemallocatedrev
    leafs["cpmProcExtMemFreedRev"] = cpmprocessextreventry.Cpmprocextmemfreedrev
    leafs["cpmProcExtInvokedRev"] = cpmprocessextreventry.Cpmprocextinvokedrev
    leafs["cpmProcExtRuntimeRev"] = cpmprocessextreventry.Cpmprocextruntimerev
    leafs["cpmProcExtUtil5SecRev"] = cpmprocessextreventry.Cpmprocextutil5Secrev
    leafs["cpmProcExtUtil1MinRev"] = cpmprocessextreventry.Cpmprocextutil1Minrev
    leafs["cpmProcExtUtil5MinRev"] = cpmprocessextreventry.Cpmprocextutil5Minrev
    leafs["cpmProcExtPriorityRev"] = cpmprocessextreventry.Cpmprocextpriorityrev
    leafs["cpmProcessType"] = cpmprocessextreventry.Cpmprocesstype
    leafs["cpmProcessRespawn"] = cpmprocessextreventry.Cpmprocessrespawn
    leafs["cpmProcessRespawnCount"] = cpmprocessextreventry.Cpmprocessrespawncount
    leafs["cpmProcessRespawnAfterLastPatch"] = cpmprocessextreventry.Cpmprocessrespawnafterlastpatch
    leafs["cpmProcessMemoryCore"] = cpmprocessextreventry.Cpmprocessmemorycore
    leafs["cpmProcessLastRestartUser"] = cpmprocessextreventry.Cpmprocesslastrestartuser
    leafs["cpmProcessTextSegmentSize"] = cpmprocessextreventry.Cpmprocesstextsegmentsize
    leafs["cpmProcessDataSegmentSize"] = cpmprocessextreventry.Cpmprocessdatasegmentsize
    leafs["cpmProcessStackSize"] = cpmprocessextreventry.Cpmprocessstacksize
    leafs["cpmProcessDynamicMemorySize"] = cpmprocessextreventry.Cpmprocessdynamicmemorysize
    leafs["cpmProcExtMemAllocatedRevOvrflw"] = cpmprocessextreventry.Cpmprocextmemallocatedrevovrflw
    leafs["cpmProcExtHCMemAllocatedRev"] = cpmprocessextreventry.Cpmprocexthcmemallocatedrev
    leafs["cpmProcExtMemFreedRevOvrflw"] = cpmprocessextreventry.Cpmprocextmemfreedrevovrflw
    leafs["cpmProcExtHCMemFreedRev"] = cpmprocessextreventry.Cpmprocexthcmemfreedrev
    leafs["cpmProcessTextSegmentSizeOvrflw"] = cpmprocessextreventry.Cpmprocesstextsegmentsizeovrflw
    leafs["cpmProcessHCTextSegmentSize"] = cpmprocessextreventry.Cpmprocesshctextsegmentsize
    leafs["cpmProcessDataSegmentSizeOvrflw"] = cpmprocessextreventry.Cpmprocessdatasegmentsizeovrflw
    leafs["cpmProcessHCDataSegmentSize"] = cpmprocessextreventry.Cpmprocesshcdatasegmentsize
    leafs["cpmProcessStackSizeOvrflw"] = cpmprocessextreventry.Cpmprocessstacksizeovrflw
    leafs["cpmProcessHCStackSize"] = cpmprocessextreventry.Cpmprocesshcstacksize
    leafs["cpmProcessDynamicMemorySizeOvrflw"] = cpmprocessextreventry.Cpmprocessdynamicmemorysizeovrflw
    leafs["cpmProcessHCDynamicMemorySize"] = cpmprocessextreventry.Cpmprocesshcdynamicmemorysize
    return leafs
}

func (cpmprocessextreventry *CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry) GetBundleName() string { return "cisco_ios_xe" }

func (cpmprocessextreventry *CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry) GetYangName() string { return "cpmProcessExtRevEntry" }

func (cpmprocessextreventry *CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpmprocessextreventry *CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpmprocessextreventry *CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpmprocessextreventry *CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry) SetParent(parent types.Entity) { cpmprocessextreventry.parent = parent }

func (cpmprocessextreventry *CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry) GetParent() types.Entity { return cpmprocessextreventry.parent }

func (cpmprocessextreventry *CISCOPROCESSMIB_Cpmprocessextrevtable_Cpmprocessextreventry) GetParentYangName() string { return "cpmProcessExtRevTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry containing information about CPU thresholding parameters.
    // cpmCPUTotalIndex identifies the CPU (or group of CPUs) for which this
    // configuration applies. The type is slice of
    // CISCOPROCESSMIB_Cpmcputhresholdtable_Cpmcputhresholdentry.
    Cpmcputhresholdentry []CISCOPROCESSMIB_Cpmcputhresholdtable_Cpmcputhresholdentry
}

func (cpmcputhresholdtable *CISCOPROCESSMIB_Cpmcputhresholdtable) GetFilter() yfilter.YFilter { return cpmcputhresholdtable.YFilter }

func (cpmcputhresholdtable *CISCOPROCESSMIB_Cpmcputhresholdtable) SetFilter(yf yfilter.YFilter) { cpmcputhresholdtable.YFilter = yf }

func (cpmcputhresholdtable *CISCOPROCESSMIB_Cpmcputhresholdtable) GetGoName(yname string) string {
    if yname == "cpmCPUThresholdEntry" { return "Cpmcputhresholdentry" }
    return ""
}

func (cpmcputhresholdtable *CISCOPROCESSMIB_Cpmcputhresholdtable) GetSegmentPath() string {
    return "cpmCPUThresholdTable"
}

func (cpmcputhresholdtable *CISCOPROCESSMIB_Cpmcputhresholdtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cpmCPUThresholdEntry" {
        for _, c := range cpmcputhresholdtable.Cpmcputhresholdentry {
            if cpmcputhresholdtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOPROCESSMIB_Cpmcputhresholdtable_Cpmcputhresholdentry{}
        cpmcputhresholdtable.Cpmcputhresholdentry = append(cpmcputhresholdtable.Cpmcputhresholdentry, child)
        return &cpmcputhresholdtable.Cpmcputhresholdentry[len(cpmcputhresholdtable.Cpmcputhresholdentry)-1]
    }
    return nil
}

func (cpmcputhresholdtable *CISCOPROCESSMIB_Cpmcputhresholdtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cpmcputhresholdtable.Cpmcputhresholdentry {
        children[cpmcputhresholdtable.Cpmcputhresholdentry[i].GetSegmentPath()] = &cpmcputhresholdtable.Cpmcputhresholdentry[i]
    }
    return children
}

func (cpmcputhresholdtable *CISCOPROCESSMIB_Cpmcputhresholdtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cpmcputhresholdtable *CISCOPROCESSMIB_Cpmcputhresholdtable) GetBundleName() string { return "cisco_ios_xe" }

func (cpmcputhresholdtable *CISCOPROCESSMIB_Cpmcputhresholdtable) GetYangName() string { return "cpmCPUThresholdTable" }

func (cpmcputhresholdtable *CISCOPROCESSMIB_Cpmcputhresholdtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpmcputhresholdtable *CISCOPROCESSMIB_Cpmcputhresholdtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpmcputhresholdtable *CISCOPROCESSMIB_Cpmcputhresholdtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpmcputhresholdtable *CISCOPROCESSMIB_Cpmcputhresholdtable) SetParent(parent types.Entity) { cpmcputhresholdtable.parent = parent }

func (cpmcputhresholdtable *CISCOPROCESSMIB_Cpmcputhresholdtable) GetParent() types.Entity { return cpmcputhresholdtable.parent }

func (cpmcputhresholdtable *CISCOPROCESSMIB_Cpmcputhresholdtable) GetParentYangName() string { return "CISCO-PROCESS-MIB" }

// CISCOPROCESSMIB_Cpmcputhresholdtable_Cpmcputhresholdentry
// An entry containing information about
// CPU thresholding parameters. cpmCPUTotalIndex
// identifies the CPU (or group of CPUs) for which this
// configuration applies.
type CISCOPROCESSMIB_Cpmcputhresholdtable_Cpmcputhresholdentry struct {
    parent types.Entity
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

func (cpmcputhresholdentry *CISCOPROCESSMIB_Cpmcputhresholdtable_Cpmcputhresholdentry) GetFilter() yfilter.YFilter { return cpmcputhresholdentry.YFilter }

func (cpmcputhresholdentry *CISCOPROCESSMIB_Cpmcputhresholdtable_Cpmcputhresholdentry) SetFilter(yf yfilter.YFilter) { cpmcputhresholdentry.YFilter = yf }

func (cpmcputhresholdentry *CISCOPROCESSMIB_Cpmcputhresholdtable_Cpmcputhresholdentry) GetGoName(yname string) string {
    if yname == "cpmCPUTotalIndex" { return "Cpmcputotalindex" }
    if yname == "cpmCPUThresholdClass" { return "Cpmcputhresholdclass" }
    if yname == "cpmCPURisingThresholdValue" { return "Cpmcpurisingthresholdvalue" }
    if yname == "cpmCPURisingThresholdPeriod" { return "Cpmcpurisingthresholdperiod" }
    if yname == "cpmCPUFallingThresholdValue" { return "Cpmcpufallingthresholdvalue" }
    if yname == "cpmCPUFallingThresholdPeriod" { return "Cpmcpufallingthresholdperiod" }
    if yname == "cpmCPUThresholdEntryStatus" { return "Cpmcputhresholdentrystatus" }
    return ""
}

func (cpmcputhresholdentry *CISCOPROCESSMIB_Cpmcputhresholdtable_Cpmcputhresholdentry) GetSegmentPath() string {
    return "cpmCPUThresholdEntry" + "[cpmCPUTotalIndex='" + fmt.Sprintf("%v", cpmcputhresholdentry.Cpmcputotalindex) + "']" + "[cpmCPUThresholdClass='" + fmt.Sprintf("%v", cpmcputhresholdentry.Cpmcputhresholdclass) + "']"
}

func (cpmcputhresholdentry *CISCOPROCESSMIB_Cpmcputhresholdtable_Cpmcputhresholdentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cpmcputhresholdentry *CISCOPROCESSMIB_Cpmcputhresholdtable_Cpmcputhresholdentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cpmcputhresholdentry *CISCOPROCESSMIB_Cpmcputhresholdtable_Cpmcputhresholdentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cpmCPUTotalIndex"] = cpmcputhresholdentry.Cpmcputotalindex
    leafs["cpmCPUThresholdClass"] = cpmcputhresholdentry.Cpmcputhresholdclass
    leafs["cpmCPURisingThresholdValue"] = cpmcputhresholdentry.Cpmcpurisingthresholdvalue
    leafs["cpmCPURisingThresholdPeriod"] = cpmcputhresholdentry.Cpmcpurisingthresholdperiod
    leafs["cpmCPUFallingThresholdValue"] = cpmcputhresholdentry.Cpmcpufallingthresholdvalue
    leafs["cpmCPUFallingThresholdPeriod"] = cpmcputhresholdentry.Cpmcpufallingthresholdperiod
    leafs["cpmCPUThresholdEntryStatus"] = cpmcputhresholdentry.Cpmcputhresholdentrystatus
    return leafs
}

func (cpmcputhresholdentry *CISCOPROCESSMIB_Cpmcputhresholdtable_Cpmcputhresholdentry) GetBundleName() string { return "cisco_ios_xe" }

func (cpmcputhresholdentry *CISCOPROCESSMIB_Cpmcputhresholdtable_Cpmcputhresholdentry) GetYangName() string { return "cpmCPUThresholdEntry" }

func (cpmcputhresholdentry *CISCOPROCESSMIB_Cpmcputhresholdtable_Cpmcputhresholdentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpmcputhresholdentry *CISCOPROCESSMIB_Cpmcputhresholdtable_Cpmcputhresholdentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpmcputhresholdentry *CISCOPROCESSMIB_Cpmcputhresholdtable_Cpmcputhresholdentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpmcputhresholdentry *CISCOPROCESSMIB_Cpmcputhresholdtable_Cpmcputhresholdentry) SetParent(parent types.Entity) { cpmcputhresholdentry.parent = parent }

func (cpmcputhresholdentry *CISCOPROCESSMIB_Cpmcputhresholdtable_Cpmcputhresholdentry) GetParent() types.Entity { return cpmcputhresholdentry.parent }

func (cpmcputhresholdentry *CISCOPROCESSMIB_Cpmcputhresholdtable_Cpmcputhresholdentry) GetParentYangName() string { return "cpmCPUThresholdTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // A historical sample of CPU utilization statistics. cpmCPUTotalIndex
    // identifies the CPU (or group of CPUs) for which this history is collected. 
    // When the cpmCPUHistorySize is reached the least recent entry is lost. The
    // type is slice of CISCOPROCESSMIB_Cpmcpuhistorytable_Cpmcpuhistoryentry.
    Cpmcpuhistoryentry []CISCOPROCESSMIB_Cpmcpuhistorytable_Cpmcpuhistoryentry
}

func (cpmcpuhistorytable *CISCOPROCESSMIB_Cpmcpuhistorytable) GetFilter() yfilter.YFilter { return cpmcpuhistorytable.YFilter }

func (cpmcpuhistorytable *CISCOPROCESSMIB_Cpmcpuhistorytable) SetFilter(yf yfilter.YFilter) { cpmcpuhistorytable.YFilter = yf }

func (cpmcpuhistorytable *CISCOPROCESSMIB_Cpmcpuhistorytable) GetGoName(yname string) string {
    if yname == "cpmCPUHistoryEntry" { return "Cpmcpuhistoryentry" }
    return ""
}

func (cpmcpuhistorytable *CISCOPROCESSMIB_Cpmcpuhistorytable) GetSegmentPath() string {
    return "cpmCPUHistoryTable"
}

func (cpmcpuhistorytable *CISCOPROCESSMIB_Cpmcpuhistorytable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cpmCPUHistoryEntry" {
        for _, c := range cpmcpuhistorytable.Cpmcpuhistoryentry {
            if cpmcpuhistorytable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOPROCESSMIB_Cpmcpuhistorytable_Cpmcpuhistoryentry{}
        cpmcpuhistorytable.Cpmcpuhistoryentry = append(cpmcpuhistorytable.Cpmcpuhistoryentry, child)
        return &cpmcpuhistorytable.Cpmcpuhistoryentry[len(cpmcpuhistorytable.Cpmcpuhistoryentry)-1]
    }
    return nil
}

func (cpmcpuhistorytable *CISCOPROCESSMIB_Cpmcpuhistorytable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cpmcpuhistorytable.Cpmcpuhistoryentry {
        children[cpmcpuhistorytable.Cpmcpuhistoryentry[i].GetSegmentPath()] = &cpmcpuhistorytable.Cpmcpuhistoryentry[i]
    }
    return children
}

func (cpmcpuhistorytable *CISCOPROCESSMIB_Cpmcpuhistorytable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cpmcpuhistorytable *CISCOPROCESSMIB_Cpmcpuhistorytable) GetBundleName() string { return "cisco_ios_xe" }

func (cpmcpuhistorytable *CISCOPROCESSMIB_Cpmcpuhistorytable) GetYangName() string { return "cpmCPUHistoryTable" }

func (cpmcpuhistorytable *CISCOPROCESSMIB_Cpmcpuhistorytable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpmcpuhistorytable *CISCOPROCESSMIB_Cpmcpuhistorytable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpmcpuhistorytable *CISCOPROCESSMIB_Cpmcpuhistorytable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpmcpuhistorytable *CISCOPROCESSMIB_Cpmcpuhistorytable) SetParent(parent types.Entity) { cpmcpuhistorytable.parent = parent }

func (cpmcpuhistorytable *CISCOPROCESSMIB_Cpmcpuhistorytable) GetParent() types.Entity { return cpmcpuhistorytable.parent }

func (cpmcpuhistorytable *CISCOPROCESSMIB_Cpmcpuhistorytable) GetParentYangName() string { return "CISCO-PROCESS-MIB" }

// CISCOPROCESSMIB_Cpmcpuhistorytable_Cpmcpuhistoryentry
// A historical sample of CPU utilization statistics.
// cpmCPUTotalIndex identifies the CPU (or group of CPUs)
// for which this history is collected. 
// When the cpmCPUHistorySize is
// reached the least recent entry is lost.
type CISCOPROCESSMIB_Cpmcpuhistorytable_Cpmcpuhistoryentry struct {
    parent types.Entity
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

func (cpmcpuhistoryentry *CISCOPROCESSMIB_Cpmcpuhistorytable_Cpmcpuhistoryentry) GetFilter() yfilter.YFilter { return cpmcpuhistoryentry.YFilter }

func (cpmcpuhistoryentry *CISCOPROCESSMIB_Cpmcpuhistorytable_Cpmcpuhistoryentry) SetFilter(yf yfilter.YFilter) { cpmcpuhistoryentry.YFilter = yf }

func (cpmcpuhistoryentry *CISCOPROCESSMIB_Cpmcpuhistorytable_Cpmcpuhistoryentry) GetGoName(yname string) string {
    if yname == "cpmCPUTotalIndex" { return "Cpmcputotalindex" }
    if yname == "cpmCPUHistoryReportId" { return "Cpmcpuhistoryreportid" }
    if yname == "cpmCPUHistoryReportSize" { return "Cpmcpuhistoryreportsize" }
    if yname == "cpmCPUHistoryTotalUtil" { return "Cpmcpuhistorytotalutil" }
    if yname == "cpmCPUHistoryInterruptUtil" { return "Cpmcpuhistoryinterruptutil" }
    if yname == "cpmCPUHistoryCreatedTime" { return "Cpmcpuhistorycreatedtime" }
    return ""
}

func (cpmcpuhistoryentry *CISCOPROCESSMIB_Cpmcpuhistorytable_Cpmcpuhistoryentry) GetSegmentPath() string {
    return "cpmCPUHistoryEntry" + "[cpmCPUTotalIndex='" + fmt.Sprintf("%v", cpmcpuhistoryentry.Cpmcputotalindex) + "']" + "[cpmCPUHistoryReportId='" + fmt.Sprintf("%v", cpmcpuhistoryentry.Cpmcpuhistoryreportid) + "']"
}

func (cpmcpuhistoryentry *CISCOPROCESSMIB_Cpmcpuhistorytable_Cpmcpuhistoryentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cpmcpuhistoryentry *CISCOPROCESSMIB_Cpmcpuhistorytable_Cpmcpuhistoryentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cpmcpuhistoryentry *CISCOPROCESSMIB_Cpmcpuhistorytable_Cpmcpuhistoryentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cpmCPUTotalIndex"] = cpmcpuhistoryentry.Cpmcputotalindex
    leafs["cpmCPUHistoryReportId"] = cpmcpuhistoryentry.Cpmcpuhistoryreportid
    leafs["cpmCPUHistoryReportSize"] = cpmcpuhistoryentry.Cpmcpuhistoryreportsize
    leafs["cpmCPUHistoryTotalUtil"] = cpmcpuhistoryentry.Cpmcpuhistorytotalutil
    leafs["cpmCPUHistoryInterruptUtil"] = cpmcpuhistoryentry.Cpmcpuhistoryinterruptutil
    leafs["cpmCPUHistoryCreatedTime"] = cpmcpuhistoryentry.Cpmcpuhistorycreatedtime
    return leafs
}

func (cpmcpuhistoryentry *CISCOPROCESSMIB_Cpmcpuhistorytable_Cpmcpuhistoryentry) GetBundleName() string { return "cisco_ios_xe" }

func (cpmcpuhistoryentry *CISCOPROCESSMIB_Cpmcpuhistorytable_Cpmcpuhistoryentry) GetYangName() string { return "cpmCPUHistoryEntry" }

func (cpmcpuhistoryentry *CISCOPROCESSMIB_Cpmcpuhistorytable_Cpmcpuhistoryentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpmcpuhistoryentry *CISCOPROCESSMIB_Cpmcpuhistorytable_Cpmcpuhistoryentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpmcpuhistoryentry *CISCOPROCESSMIB_Cpmcpuhistorytable_Cpmcpuhistoryentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpmcpuhistoryentry *CISCOPROCESSMIB_Cpmcpuhistorytable_Cpmcpuhistoryentry) SetParent(parent types.Entity) { cpmcpuhistoryentry.parent = parent }

func (cpmcpuhistoryentry *CISCOPROCESSMIB_Cpmcpuhistorytable_Cpmcpuhistoryentry) GetParent() types.Entity { return cpmcpuhistoryentry.parent }

func (cpmcpuhistoryentry *CISCOPROCESSMIB_Cpmcpuhistorytable_Cpmcpuhistoryentry) GetParentYangName() string { return "cpmCPUHistoryTable" }

// CISCOPROCESSMIB_Cpmcpuprocesshistorytable
// A list of process history entries. This table contains
// CPU utilization of processes which crossed the 
// cpmCPUHistoryThreshold.
type CISCOPROCESSMIB_Cpmcpuprocesshistorytable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A historical sample of process utilization statistics. The entries in this
    // table will have corresponding entires in the cpmCPUHistoryTable. The
    // entries in this table get deleted when the entry associated with this entry
    // in the cpmCPUHistoryTable  gets deleted. The type is slice of
    // CISCOPROCESSMIB_Cpmcpuprocesshistorytable_Cpmcpuprocesshistoryentry.
    Cpmcpuprocesshistoryentry []CISCOPROCESSMIB_Cpmcpuprocesshistorytable_Cpmcpuprocesshistoryentry
}

func (cpmcpuprocesshistorytable *CISCOPROCESSMIB_Cpmcpuprocesshistorytable) GetFilter() yfilter.YFilter { return cpmcpuprocesshistorytable.YFilter }

func (cpmcpuprocesshistorytable *CISCOPROCESSMIB_Cpmcpuprocesshistorytable) SetFilter(yf yfilter.YFilter) { cpmcpuprocesshistorytable.YFilter = yf }

func (cpmcpuprocesshistorytable *CISCOPROCESSMIB_Cpmcpuprocesshistorytable) GetGoName(yname string) string {
    if yname == "cpmCPUProcessHistoryEntry" { return "Cpmcpuprocesshistoryentry" }
    return ""
}

func (cpmcpuprocesshistorytable *CISCOPROCESSMIB_Cpmcpuprocesshistorytable) GetSegmentPath() string {
    return "cpmCPUProcessHistoryTable"
}

func (cpmcpuprocesshistorytable *CISCOPROCESSMIB_Cpmcpuprocesshistorytable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cpmCPUProcessHistoryEntry" {
        for _, c := range cpmcpuprocesshistorytable.Cpmcpuprocesshistoryentry {
            if cpmcpuprocesshistorytable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOPROCESSMIB_Cpmcpuprocesshistorytable_Cpmcpuprocesshistoryentry{}
        cpmcpuprocesshistorytable.Cpmcpuprocesshistoryentry = append(cpmcpuprocesshistorytable.Cpmcpuprocesshistoryentry, child)
        return &cpmcpuprocesshistorytable.Cpmcpuprocesshistoryentry[len(cpmcpuprocesshistorytable.Cpmcpuprocesshistoryentry)-1]
    }
    return nil
}

func (cpmcpuprocesshistorytable *CISCOPROCESSMIB_Cpmcpuprocesshistorytable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cpmcpuprocesshistorytable.Cpmcpuprocesshistoryentry {
        children[cpmcpuprocesshistorytable.Cpmcpuprocesshistoryentry[i].GetSegmentPath()] = &cpmcpuprocesshistorytable.Cpmcpuprocesshistoryentry[i]
    }
    return children
}

func (cpmcpuprocesshistorytable *CISCOPROCESSMIB_Cpmcpuprocesshistorytable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cpmcpuprocesshistorytable *CISCOPROCESSMIB_Cpmcpuprocesshistorytable) GetBundleName() string { return "cisco_ios_xe" }

func (cpmcpuprocesshistorytable *CISCOPROCESSMIB_Cpmcpuprocesshistorytable) GetYangName() string { return "cpmCPUProcessHistoryTable" }

func (cpmcpuprocesshistorytable *CISCOPROCESSMIB_Cpmcpuprocesshistorytable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpmcpuprocesshistorytable *CISCOPROCESSMIB_Cpmcpuprocesshistorytable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpmcpuprocesshistorytable *CISCOPROCESSMIB_Cpmcpuprocesshistorytable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpmcpuprocesshistorytable *CISCOPROCESSMIB_Cpmcpuprocesshistorytable) SetParent(parent types.Entity) { cpmcpuprocesshistorytable.parent = parent }

func (cpmcpuprocesshistorytable *CISCOPROCESSMIB_Cpmcpuprocesshistorytable) GetParent() types.Entity { return cpmcpuprocesshistorytable.parent }

func (cpmcpuprocesshistorytable *CISCOPROCESSMIB_Cpmcpuprocesshistorytable) GetParentYangName() string { return "CISCO-PROCESS-MIB" }

// CISCOPROCESSMIB_Cpmcpuprocesshistorytable_Cpmcpuprocesshistoryentry
// A historical sample of process utilization
// statistics. The entries in this table will have
// corresponding entires in the cpmCPUHistoryTable.
// The entries in this table get deleted when the entry
// associated with this entry in the cpmCPUHistoryTable 
// gets deleted.
type CISCOPROCESSMIB_Cpmcpuprocesshistorytable_Cpmcpuprocesshistoryentry struct {
    parent types.Entity
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

func (cpmcpuprocesshistoryentry *CISCOPROCESSMIB_Cpmcpuprocesshistorytable_Cpmcpuprocesshistoryentry) GetFilter() yfilter.YFilter { return cpmcpuprocesshistoryentry.YFilter }

func (cpmcpuprocesshistoryentry *CISCOPROCESSMIB_Cpmcpuprocesshistorytable_Cpmcpuprocesshistoryentry) SetFilter(yf yfilter.YFilter) { cpmcpuprocesshistoryentry.YFilter = yf }

func (cpmcpuprocesshistoryentry *CISCOPROCESSMIB_Cpmcpuprocesshistorytable_Cpmcpuprocesshistoryentry) GetGoName(yname string) string {
    if yname == "cpmCPUTotalIndex" { return "Cpmcputotalindex" }
    if yname == "cpmCPUHistoryReportId" { return "Cpmcpuhistoryreportid" }
    if yname == "cpmCPUProcessHistoryIndex" { return "Cpmcpuprocesshistoryindex" }
    if yname == "cpmCPUHistoryProcId" { return "Cpmcpuhistoryprocid" }
    if yname == "cpmCPUHistoryProcName" { return "Cpmcpuhistoryprocname" }
    if yname == "cpmCPUHistoryProcCreated" { return "Cpmcpuhistoryproccreated" }
    if yname == "cpmCPUHistoryProcUtil" { return "Cpmcpuhistoryprocutil" }
    return ""
}

func (cpmcpuprocesshistoryentry *CISCOPROCESSMIB_Cpmcpuprocesshistorytable_Cpmcpuprocesshistoryentry) GetSegmentPath() string {
    return "cpmCPUProcessHistoryEntry" + "[cpmCPUTotalIndex='" + fmt.Sprintf("%v", cpmcpuprocesshistoryentry.Cpmcputotalindex) + "']" + "[cpmCPUHistoryReportId='" + fmt.Sprintf("%v", cpmcpuprocesshistoryentry.Cpmcpuhistoryreportid) + "']" + "[cpmCPUProcessHistoryIndex='" + fmt.Sprintf("%v", cpmcpuprocesshistoryentry.Cpmcpuprocesshistoryindex) + "']"
}

func (cpmcpuprocesshistoryentry *CISCOPROCESSMIB_Cpmcpuprocesshistorytable_Cpmcpuprocesshistoryentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cpmcpuprocesshistoryentry *CISCOPROCESSMIB_Cpmcpuprocesshistorytable_Cpmcpuprocesshistoryentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cpmcpuprocesshistoryentry *CISCOPROCESSMIB_Cpmcpuprocesshistorytable_Cpmcpuprocesshistoryentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cpmCPUTotalIndex"] = cpmcpuprocesshistoryentry.Cpmcputotalindex
    leafs["cpmCPUHistoryReportId"] = cpmcpuprocesshistoryentry.Cpmcpuhistoryreportid
    leafs["cpmCPUProcessHistoryIndex"] = cpmcpuprocesshistoryentry.Cpmcpuprocesshistoryindex
    leafs["cpmCPUHistoryProcId"] = cpmcpuprocesshistoryentry.Cpmcpuhistoryprocid
    leafs["cpmCPUHistoryProcName"] = cpmcpuprocesshistoryentry.Cpmcpuhistoryprocname
    leafs["cpmCPUHistoryProcCreated"] = cpmcpuprocesshistoryentry.Cpmcpuhistoryproccreated
    leafs["cpmCPUHistoryProcUtil"] = cpmcpuprocesshistoryentry.Cpmcpuhistoryprocutil
    return leafs
}

func (cpmcpuprocesshistoryentry *CISCOPROCESSMIB_Cpmcpuprocesshistorytable_Cpmcpuprocesshistoryentry) GetBundleName() string { return "cisco_ios_xe" }

func (cpmcpuprocesshistoryentry *CISCOPROCESSMIB_Cpmcpuprocesshistorytable_Cpmcpuprocesshistoryentry) GetYangName() string { return "cpmCPUProcessHistoryEntry" }

func (cpmcpuprocesshistoryentry *CISCOPROCESSMIB_Cpmcpuprocesshistorytable_Cpmcpuprocesshistoryentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpmcpuprocesshistoryentry *CISCOPROCESSMIB_Cpmcpuprocesshistorytable_Cpmcpuprocesshistoryentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpmcpuprocesshistoryentry *CISCOPROCESSMIB_Cpmcpuprocesshistorytable_Cpmcpuprocesshistoryentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpmcpuprocesshistoryentry *CISCOPROCESSMIB_Cpmcpuprocesshistorytable_Cpmcpuprocesshistoryentry) SetParent(parent types.Entity) { cpmcpuprocesshistoryentry.parent = parent }

func (cpmcpuprocesshistoryentry *CISCOPROCESSMIB_Cpmcpuprocesshistorytable_Cpmcpuprocesshistoryentry) GetParent() types.Entity { return cpmcpuprocesshistoryentry.parent }

func (cpmcpuprocesshistoryentry *CISCOPROCESSMIB_Cpmcpuprocesshistorytable_Cpmcpuprocesshistoryentry) GetParentYangName() string { return "cpmCPUProcessHistoryTable" }

// CISCOPROCESSMIB_Cpmthreadtable
// This table contains generic information about
// POSIX threads in the device.
type CISCOPROCESSMIB_Cpmthreadtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry containing the general statistics of a POSIX thread. The type is
    // slice of CISCOPROCESSMIB_Cpmthreadtable_Cpmthreadentry.
    Cpmthreadentry []CISCOPROCESSMIB_Cpmthreadtable_Cpmthreadentry
}

func (cpmthreadtable *CISCOPROCESSMIB_Cpmthreadtable) GetFilter() yfilter.YFilter { return cpmthreadtable.YFilter }

func (cpmthreadtable *CISCOPROCESSMIB_Cpmthreadtable) SetFilter(yf yfilter.YFilter) { cpmthreadtable.YFilter = yf }

func (cpmthreadtable *CISCOPROCESSMIB_Cpmthreadtable) GetGoName(yname string) string {
    if yname == "cpmThreadEntry" { return "Cpmthreadentry" }
    return ""
}

func (cpmthreadtable *CISCOPROCESSMIB_Cpmthreadtable) GetSegmentPath() string {
    return "cpmThreadTable"
}

func (cpmthreadtable *CISCOPROCESSMIB_Cpmthreadtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cpmThreadEntry" {
        for _, c := range cpmthreadtable.Cpmthreadentry {
            if cpmthreadtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOPROCESSMIB_Cpmthreadtable_Cpmthreadentry{}
        cpmthreadtable.Cpmthreadentry = append(cpmthreadtable.Cpmthreadentry, child)
        return &cpmthreadtable.Cpmthreadentry[len(cpmthreadtable.Cpmthreadentry)-1]
    }
    return nil
}

func (cpmthreadtable *CISCOPROCESSMIB_Cpmthreadtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cpmthreadtable.Cpmthreadentry {
        children[cpmthreadtable.Cpmthreadentry[i].GetSegmentPath()] = &cpmthreadtable.Cpmthreadentry[i]
    }
    return children
}

func (cpmthreadtable *CISCOPROCESSMIB_Cpmthreadtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cpmthreadtable *CISCOPROCESSMIB_Cpmthreadtable) GetBundleName() string { return "cisco_ios_xe" }

func (cpmthreadtable *CISCOPROCESSMIB_Cpmthreadtable) GetYangName() string { return "cpmThreadTable" }

func (cpmthreadtable *CISCOPROCESSMIB_Cpmthreadtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpmthreadtable *CISCOPROCESSMIB_Cpmthreadtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpmthreadtable *CISCOPROCESSMIB_Cpmthreadtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpmthreadtable *CISCOPROCESSMIB_Cpmthreadtable) SetParent(parent types.Entity) { cpmthreadtable.parent = parent }

func (cpmthreadtable *CISCOPROCESSMIB_Cpmthreadtable) GetParent() types.Entity { return cpmthreadtable.parent }

func (cpmthreadtable *CISCOPROCESSMIB_Cpmthreadtable) GetParentYangName() string { return "CISCO-PROCESS-MIB" }

// CISCOPROCESSMIB_Cpmthreadtable_Cpmthreadentry
// An entry containing the general statistics
// of a POSIX thread.
type CISCOPROCESSMIB_Cpmthreadtable_Cpmthreadentry struct {
    parent types.Entity
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
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
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

func (cpmthreadentry *CISCOPROCESSMIB_Cpmthreadtable_Cpmthreadentry) GetFilter() yfilter.YFilter { return cpmthreadentry.YFilter }

func (cpmthreadentry *CISCOPROCESSMIB_Cpmthreadtable_Cpmthreadentry) SetFilter(yf yfilter.YFilter) { cpmthreadentry.YFilter = yf }

func (cpmthreadentry *CISCOPROCESSMIB_Cpmthreadtable_Cpmthreadentry) GetGoName(yname string) string {
    if yname == "cpmCPUTotalIndex" { return "Cpmcputotalindex" }
    if yname == "cpmProcessPID" { return "Cpmprocesspid" }
    if yname == "cpmThreadID" { return "Cpmthreadid" }
    if yname == "cpmThreadName" { return "Cpmthreadname" }
    if yname == "cpmThreadPriority" { return "Cpmthreadpriority" }
    if yname == "cpmThreadState" { return "Cpmthreadstate" }
    if yname == "cpmThreadBlockingProcess" { return "Cpmthreadblockingprocess" }
    if yname == "cpmThreadCpuUtilization" { return "Cpmthreadcpuutilization" }
    if yname == "cpmThreadStackSize" { return "Cpmthreadstacksize" }
    if yname == "cpmThreadStackSizeOvrflw" { return "Cpmthreadstacksizeovrflw" }
    if yname == "cpmThreadHCStackSize" { return "Cpmthreadhcstacksize" }
    return ""
}

func (cpmthreadentry *CISCOPROCESSMIB_Cpmthreadtable_Cpmthreadentry) GetSegmentPath() string {
    return "cpmThreadEntry" + "[cpmCPUTotalIndex='" + fmt.Sprintf("%v", cpmthreadentry.Cpmcputotalindex) + "']" + "[cpmProcessPID='" + fmt.Sprintf("%v", cpmthreadentry.Cpmprocesspid) + "']" + "[cpmThreadID='" + fmt.Sprintf("%v", cpmthreadentry.Cpmthreadid) + "']"
}

func (cpmthreadentry *CISCOPROCESSMIB_Cpmthreadtable_Cpmthreadentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cpmthreadentry *CISCOPROCESSMIB_Cpmthreadtable_Cpmthreadentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cpmthreadentry *CISCOPROCESSMIB_Cpmthreadtable_Cpmthreadentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cpmCPUTotalIndex"] = cpmthreadentry.Cpmcputotalindex
    leafs["cpmProcessPID"] = cpmthreadentry.Cpmprocesspid
    leafs["cpmThreadID"] = cpmthreadentry.Cpmthreadid
    leafs["cpmThreadName"] = cpmthreadentry.Cpmthreadname
    leafs["cpmThreadPriority"] = cpmthreadentry.Cpmthreadpriority
    leafs["cpmThreadState"] = cpmthreadentry.Cpmthreadstate
    leafs["cpmThreadBlockingProcess"] = cpmthreadentry.Cpmthreadblockingprocess
    leafs["cpmThreadCpuUtilization"] = cpmthreadentry.Cpmthreadcpuutilization
    leafs["cpmThreadStackSize"] = cpmthreadentry.Cpmthreadstacksize
    leafs["cpmThreadStackSizeOvrflw"] = cpmthreadentry.Cpmthreadstacksizeovrflw
    leafs["cpmThreadHCStackSize"] = cpmthreadentry.Cpmthreadhcstacksize
    return leafs
}

func (cpmthreadentry *CISCOPROCESSMIB_Cpmthreadtable_Cpmthreadentry) GetBundleName() string { return "cisco_ios_xe" }

func (cpmthreadentry *CISCOPROCESSMIB_Cpmthreadtable_Cpmthreadentry) GetYangName() string { return "cpmThreadEntry" }

func (cpmthreadentry *CISCOPROCESSMIB_Cpmthreadtable_Cpmthreadentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpmthreadentry *CISCOPROCESSMIB_Cpmthreadtable_Cpmthreadentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpmthreadentry *CISCOPROCESSMIB_Cpmthreadtable_Cpmthreadentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpmthreadentry *CISCOPROCESSMIB_Cpmthreadtable_Cpmthreadentry) SetParent(parent types.Entity) { cpmthreadentry.parent = parent }

func (cpmthreadentry *CISCOPROCESSMIB_Cpmthreadtable_Cpmthreadentry) GetParent() types.Entity { return cpmthreadentry.parent }

func (cpmthreadentry *CISCOPROCESSMIB_Cpmthreadtable_Cpmthreadentry) GetParentYangName() string { return "cpmThreadTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry containing the general statistics of a virtual process in a
    // virtual machine. The type is slice of
    // CISCOPROCESSMIB_Cpmvirtualprocesstable_Cpmvirtualprocessentry.
    Cpmvirtualprocessentry []CISCOPROCESSMIB_Cpmvirtualprocesstable_Cpmvirtualprocessentry
}

func (cpmvirtualprocesstable *CISCOPROCESSMIB_Cpmvirtualprocesstable) GetFilter() yfilter.YFilter { return cpmvirtualprocesstable.YFilter }

func (cpmvirtualprocesstable *CISCOPROCESSMIB_Cpmvirtualprocesstable) SetFilter(yf yfilter.YFilter) { cpmvirtualprocesstable.YFilter = yf }

func (cpmvirtualprocesstable *CISCOPROCESSMIB_Cpmvirtualprocesstable) GetGoName(yname string) string {
    if yname == "cpmVirtualProcessEntry" { return "Cpmvirtualprocessentry" }
    return ""
}

func (cpmvirtualprocesstable *CISCOPROCESSMIB_Cpmvirtualprocesstable) GetSegmentPath() string {
    return "cpmVirtualProcessTable"
}

func (cpmvirtualprocesstable *CISCOPROCESSMIB_Cpmvirtualprocesstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cpmVirtualProcessEntry" {
        for _, c := range cpmvirtualprocesstable.Cpmvirtualprocessentry {
            if cpmvirtualprocesstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOPROCESSMIB_Cpmvirtualprocesstable_Cpmvirtualprocessentry{}
        cpmvirtualprocesstable.Cpmvirtualprocessentry = append(cpmvirtualprocesstable.Cpmvirtualprocessentry, child)
        return &cpmvirtualprocesstable.Cpmvirtualprocessentry[len(cpmvirtualprocesstable.Cpmvirtualprocessentry)-1]
    }
    return nil
}

func (cpmvirtualprocesstable *CISCOPROCESSMIB_Cpmvirtualprocesstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cpmvirtualprocesstable.Cpmvirtualprocessentry {
        children[cpmvirtualprocesstable.Cpmvirtualprocessentry[i].GetSegmentPath()] = &cpmvirtualprocesstable.Cpmvirtualprocessentry[i]
    }
    return children
}

func (cpmvirtualprocesstable *CISCOPROCESSMIB_Cpmvirtualprocesstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cpmvirtualprocesstable *CISCOPROCESSMIB_Cpmvirtualprocesstable) GetBundleName() string { return "cisco_ios_xe" }

func (cpmvirtualprocesstable *CISCOPROCESSMIB_Cpmvirtualprocesstable) GetYangName() string { return "cpmVirtualProcessTable" }

func (cpmvirtualprocesstable *CISCOPROCESSMIB_Cpmvirtualprocesstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpmvirtualprocesstable *CISCOPROCESSMIB_Cpmvirtualprocesstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpmvirtualprocesstable *CISCOPROCESSMIB_Cpmvirtualprocesstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpmvirtualprocesstable *CISCOPROCESSMIB_Cpmvirtualprocesstable) SetParent(parent types.Entity) { cpmvirtualprocesstable.parent = parent }

func (cpmvirtualprocesstable *CISCOPROCESSMIB_Cpmvirtualprocesstable) GetParent() types.Entity { return cpmvirtualprocesstable.parent }

func (cpmvirtualprocesstable *CISCOPROCESSMIB_Cpmvirtualprocesstable) GetParentYangName() string { return "CISCO-PROCESS-MIB" }

// CISCOPROCESSMIB_Cpmvirtualprocesstable_Cpmvirtualprocessentry
// An entry containing the general statistics of a
// virtual process in a virtual machine.
type CISCOPROCESSMIB_Cpmvirtualprocesstable_Cpmvirtualprocessentry struct {
    parent types.Entity
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

func (cpmvirtualprocessentry *CISCOPROCESSMIB_Cpmvirtualprocesstable_Cpmvirtualprocessentry) GetFilter() yfilter.YFilter { return cpmvirtualprocessentry.YFilter }

func (cpmvirtualprocessentry *CISCOPROCESSMIB_Cpmvirtualprocesstable_Cpmvirtualprocessentry) SetFilter(yf yfilter.YFilter) { cpmvirtualprocessentry.YFilter = yf }

func (cpmvirtualprocessentry *CISCOPROCESSMIB_Cpmvirtualprocesstable_Cpmvirtualprocessentry) GetGoName(yname string) string {
    if yname == "cpmCPUTotalIndex" { return "Cpmcputotalindex" }
    if yname == "cpmProcessPID" { return "Cpmprocesspid" }
    if yname == "cpmVirtualProcessID" { return "Cpmvirtualprocessid" }
    if yname == "cpmVirtualProcessName" { return "Cpmvirtualprocessname" }
    if yname == "cpmVirtualProcessUtil5Sec" { return "Cpmvirtualprocessutil5Sec" }
    if yname == "cpmVirtualProcessUtil1Min" { return "Cpmvirtualprocessutil1Min" }
    if yname == "cpmVirtualProcessUtil5Min" { return "Cpmvirtualprocessutil5Min" }
    if yname == "cpmVirtualProcessMemAllocated" { return "Cpmvirtualprocessmemallocated" }
    if yname == "cpmVirtualProcessMemFreed" { return "Cpmvirtualprocessmemfreed" }
    if yname == "cpmVirtualProcessInvokeCount" { return "Cpmvirtualprocessinvokecount" }
    if yname == "cpmVirtualProcessRuntime" { return "Cpmvirtualprocessruntime" }
    if yname == "cpmVirtualProcessMemAllocatedOvrflw" { return "Cpmvirtualprocessmemallocatedovrflw" }
    if yname == "cpmVirtualProcessHCMemAllocated" { return "Cpmvirtualprocesshcmemallocated" }
    if yname == "cpmVirtualProcessMemFreedOvrflw" { return "Cpmvirtualprocessmemfreedovrflw" }
    if yname == "cpmVirtualProcessHCMemFreed" { return "Cpmvirtualprocesshcmemfreed" }
    return ""
}

func (cpmvirtualprocessentry *CISCOPROCESSMIB_Cpmvirtualprocesstable_Cpmvirtualprocessentry) GetSegmentPath() string {
    return "cpmVirtualProcessEntry" + "[cpmCPUTotalIndex='" + fmt.Sprintf("%v", cpmvirtualprocessentry.Cpmcputotalindex) + "']" + "[cpmProcessPID='" + fmt.Sprintf("%v", cpmvirtualprocessentry.Cpmprocesspid) + "']" + "[cpmVirtualProcessID='" + fmt.Sprintf("%v", cpmvirtualprocessentry.Cpmvirtualprocessid) + "']"
}

func (cpmvirtualprocessentry *CISCOPROCESSMIB_Cpmvirtualprocesstable_Cpmvirtualprocessentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cpmvirtualprocessentry *CISCOPROCESSMIB_Cpmvirtualprocesstable_Cpmvirtualprocessentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cpmvirtualprocessentry *CISCOPROCESSMIB_Cpmvirtualprocesstable_Cpmvirtualprocessentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cpmCPUTotalIndex"] = cpmvirtualprocessentry.Cpmcputotalindex
    leafs["cpmProcessPID"] = cpmvirtualprocessentry.Cpmprocesspid
    leafs["cpmVirtualProcessID"] = cpmvirtualprocessentry.Cpmvirtualprocessid
    leafs["cpmVirtualProcessName"] = cpmvirtualprocessentry.Cpmvirtualprocessname
    leafs["cpmVirtualProcessUtil5Sec"] = cpmvirtualprocessentry.Cpmvirtualprocessutil5Sec
    leafs["cpmVirtualProcessUtil1Min"] = cpmvirtualprocessentry.Cpmvirtualprocessutil1Min
    leafs["cpmVirtualProcessUtil5Min"] = cpmvirtualprocessentry.Cpmvirtualprocessutil5Min
    leafs["cpmVirtualProcessMemAllocated"] = cpmvirtualprocessentry.Cpmvirtualprocessmemallocated
    leafs["cpmVirtualProcessMemFreed"] = cpmvirtualprocessentry.Cpmvirtualprocessmemfreed
    leafs["cpmVirtualProcessInvokeCount"] = cpmvirtualprocessentry.Cpmvirtualprocessinvokecount
    leafs["cpmVirtualProcessRuntime"] = cpmvirtualprocessentry.Cpmvirtualprocessruntime
    leafs["cpmVirtualProcessMemAllocatedOvrflw"] = cpmvirtualprocessentry.Cpmvirtualprocessmemallocatedovrflw
    leafs["cpmVirtualProcessHCMemAllocated"] = cpmvirtualprocessentry.Cpmvirtualprocesshcmemallocated
    leafs["cpmVirtualProcessMemFreedOvrflw"] = cpmvirtualprocessentry.Cpmvirtualprocessmemfreedovrflw
    leafs["cpmVirtualProcessHCMemFreed"] = cpmvirtualprocessentry.Cpmvirtualprocesshcmemfreed
    return leafs
}

func (cpmvirtualprocessentry *CISCOPROCESSMIB_Cpmvirtualprocesstable_Cpmvirtualprocessentry) GetBundleName() string { return "cisco_ios_xe" }

func (cpmvirtualprocessentry *CISCOPROCESSMIB_Cpmvirtualprocesstable_Cpmvirtualprocessentry) GetYangName() string { return "cpmVirtualProcessEntry" }

func (cpmvirtualprocessentry *CISCOPROCESSMIB_Cpmvirtualprocesstable_Cpmvirtualprocessentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpmvirtualprocessentry *CISCOPROCESSMIB_Cpmvirtualprocesstable_Cpmvirtualprocessentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpmvirtualprocessentry *CISCOPROCESSMIB_Cpmvirtualprocesstable_Cpmvirtualprocessentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpmvirtualprocessentry *CISCOPROCESSMIB_Cpmvirtualprocesstable_Cpmvirtualprocessentry) SetParent(parent types.Entity) { cpmvirtualprocessentry.parent = parent }

func (cpmvirtualprocessentry *CISCOPROCESSMIB_Cpmvirtualprocesstable_Cpmvirtualprocessentry) GetParent() types.Entity { return cpmvirtualprocessentry.parent }

func (cpmvirtualprocessentry *CISCOPROCESSMIB_Cpmvirtualprocesstable_Cpmvirtualprocessentry) GetParentYangName() string { return "cpmVirtualProcessTable" }

