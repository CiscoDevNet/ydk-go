// An extension to the Cisco ATM MIB module for managing
// ATM implementations.
// 
// Acronyms and terms used in the MIB module:
// 
// AAL5        -- ATM adaptation layer 5.
// AIS         -- Alarm Indication Signal.
// CC          -- Continuity Check.
// End-to-end  -- End-to-end continuity checking.
//                Monitoring occurs on the entire VC
//                between two ATM end stations.
// F5 OAM      -- OAM information flow between 
//                network elements (NEs) used within 
//                virtual connections to report degraded 
//                virtual channel performance.
// OAM         -- Operations, Administration 
//                and Maintenance.
// RDI         -- Remote Detection Indication.
// Segment     -- Segment continuity checking. 
//                Monitoring occurs on a VC segment
//                between a router and a first-hop 
//                ATM switch.
// VC          -- Virtual Channel.
// VCC         -- Virtual Channel Connection.
// VCL         -- Virtual Channel Link.
package cisco_atm_ext_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_atm_ext_mib"))
}

// OamCCStatus represents                           confirm.
type OamCCStatus string

const (
    OamCCStatus_ready OamCCStatus = "ready"

    OamCCStatus_waitActiveResponse OamCCStatus = "waitActiveResponse"

    OamCCStatus_waitActiveConfirm OamCCStatus = "waitActiveConfirm"

    OamCCStatus_active OamCCStatus = "active"

    OamCCStatus_waitDeactiveConfirm OamCCStatus = "waitDeactiveConfirm"
)

// OamCCVcState represents notManaged(3)          -- VC is not managed by CC.
type OamCCVcState string

const (
    OamCCVcState_verified OamCCVcState = "verified"

    OamCCVcState_aisrdi OamCCVcState = "aisrdi"

    OamCCVcState_notManaged OamCCVcState = "notManaged"
)

