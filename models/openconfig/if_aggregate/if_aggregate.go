// Model for managing aggregated (aka bundle, LAG) interfaces.
package if_aggregate

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package if_aggregate"))
}

// AggregationType represents defined and managed
type AggregationType string

const (
    // LAG managed by LACP
    AggregationType_LACP AggregationType = "LACP"

    // Statically configured bundle / LAG
    AggregationType_STATIC AggregationType = "STATIC"
)

