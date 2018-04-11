// This module describes a YANG model for ISIS protocol configuration.
// It is a limited subset of all of the configuration parameters
// available in the variety of vendor implementations, hence it is
// expected that it would be augmented with vendor - specific configuration
// data as needed. Additional modules or submodules to handle other
// aspects of ISIS configuration, including policy, routing, types,
// LSDB and additional address families are also expected. This model
// supports the following ISIS configuration level hierarchy:
// 
// ISIS
// +-> { global ISIS configuration}
//    +-> levels +-> { level config}
//        +-> { system-level-counters }
//        +-> { level link-state-database}
//    +-> interface +-> { interface config }
//        +-> { circuit-counters }
//        +-> { levels config }
//        +-> { level adjacencies }
package isis

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package isis"))
}

// IsisMetricFlags represents Type definition for flags used in IS-IS metrics
type IsisMetricFlags string

const (
    // When this flag is not set, internal metrics are in use.
    IsisMetricFlags_INTERNAL IsisMetricFlags = "INTERNAL"

    // When this flag (referred to as the S-bit) is set, then
    // the metric is unsupported.
    IsisMetricFlags_UNSUPPORTED IsisMetricFlags = "UNSUPPORTED"
)

