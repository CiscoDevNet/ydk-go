// This module introduces four new rpc methods that are used to
// control a two-phase commit transaction on the NETCONF server.
// The normal <edit-config> operation is used to write data in the
// transaction, but the modifications are not applied until an
// explicit <commit-transaction> is sent.
// 
// A typical sequence of operations looks like this:
// 
// 
//           C                           S
//           |                           |
//           |  capability exchange      |
//           |-------------------------->|
//           |<------------------------->|
//           |                           |
//           |   <start-transaction>     |
//           |-------------------------->|
//           |<--------------------------|
//           |         <ok/>             |
//           |                           |
//           |     <edit-config>         |
//           |-------------------------->|
//           |<--------------------------|
//           |         <ok/>             |
//           |                           |
//           |  <prepare-transaction>    |
//           |-------------------------->|
//           |<--------------------------|
//           |         <ok/>             |
//           |                           |
//           |   <commit-transaction>    |
//           |-------------------------->|
//           |<--------------------------|
//           |         <ok/>             |
//           |                           |
// 
package tailf_netconf_transactions

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package tailf_netconf_transactions"))
    ydk.RegisterEntity("{http://tail-f.com/ns/netconf/transactions/1.0 start-transaction}", reflect.TypeOf(StartTransaction{}))
    ydk.RegisterEntity("tailf-netconf-transactions:start-transaction", reflect.TypeOf(StartTransaction{}))
    ydk.RegisterEntity("{http://tail-f.com/ns/netconf/transactions/1.0 prepare-transaction}", reflect.TypeOf(PrepareTransaction{}))
    ydk.RegisterEntity("tailf-netconf-transactions:prepare-transaction", reflect.TypeOf(PrepareTransaction{}))
    ydk.RegisterEntity("{http://tail-f.com/ns/netconf/transactions/1.0 commit-transaction}", reflect.TypeOf(CommitTransaction{}))
    ydk.RegisterEntity("tailf-netconf-transactions:commit-transaction", reflect.TypeOf(CommitTransaction{}))
    ydk.RegisterEntity("{http://tail-f.com/ns/netconf/transactions/1.0 abort-transaction}", reflect.TypeOf(AbortTransaction{}))
    ydk.RegisterEntity("tailf-netconf-transactions:abort-transaction", reflect.TypeOf(AbortTransaction{}))
}

// StartTransaction
// Starts a transaction towards a configuration datastore.  There
// can be a single ongoing transaction per session at any time.
// 
// When a transaction has been started, the client can send any
// NETCONF operation, but any <edit-config> or <copy-config>
// operation sent from the client MUST specify the same <target>
// as the <start-transaction>, and any <get-config> MUST specify
// the same <source> as <start-transaction>.
// 
// If the server receives an <edit-config> or <copy-config> with
// another <target>, or a <get-config> with another <source>, an
// error MUST be returned with an <error-tag> set to 'invalid-value'.
// 
// The modifications sent in the <edit-config> operations are not
// immediately applied to the configuration datastore.  Instead
// they are kept in the transaction state of the server.  The
// transaction state is only applied when a <commit-transaction>
// is received.
// 
// The client sends a <prepare-transaction> when all modifications
// have been sent.
// 
// If there is an ongoing transaction for this session already, an
// error MUST be returned with <error-app-tag> set to 'bad-state'.
type StartTransaction struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input StartTransaction_Input
}

func (startTransaction *StartTransaction) GetFilter() yfilter.YFilter { return startTransaction.YFilter }

func (startTransaction *StartTransaction) SetFilter(yf yfilter.YFilter) { startTransaction.YFilter = yf }

func (startTransaction *StartTransaction) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (startTransaction *StartTransaction) GetSegmentPath() string {
    return "tailf-netconf-transactions:start-transaction"
}

func (startTransaction *StartTransaction) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &startTransaction.Input
    }
    return nil
}

func (startTransaction *StartTransaction) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &startTransaction.Input
    return children
}

func (startTransaction *StartTransaction) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (startTransaction *StartTransaction) GetBundleName() string { return "cisco_ios_xe" }

func (startTransaction *StartTransaction) GetYangName() string { return "start-transaction" }

func (startTransaction *StartTransaction) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (startTransaction *StartTransaction) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (startTransaction *StartTransaction) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (startTransaction *StartTransaction) SetParent(parent types.Entity) { startTransaction.parent = parent }

func (startTransaction *StartTransaction) GetParent() types.Entity { return startTransaction.parent }

func (startTransaction *StartTransaction) GetParentYangName() string { return "tailf-netconf-transactions" }

// StartTransaction_Input
type StartTransaction_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // If the parameter is present in <start-transaction>, it MUST also be present
    // in any <edit-config>, <copy-config>, <get>, or <get-config> operations
    // within the transaction.  If it is not present in <start-transaction>, it
    // MUST NOT be present in any <edit-config> operation within the transaction.
    // The type is interface{}.
    WithInactive interface{}

    // Name of the configuration datastore towards which the transaction is
    // started.
    Target StartTransaction_Input_Target
}

func (input *StartTransaction_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *StartTransaction_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *StartTransaction_Input) GetGoName(yname string) string {
    if yname == "with-inactive" { return "WithInactive" }
    if yname == "target" { return "Target" }
    return ""
}

func (input *StartTransaction_Input) GetSegmentPath() string {
    return "input"
}

func (input *StartTransaction_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "target" {
        return &input.Target
    }
    return nil
}

func (input *StartTransaction_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["target"] = &input.Target
    return children
}

func (input *StartTransaction_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["with-inactive"] = input.WithInactive
    return leafs
}

func (input *StartTransaction_Input) GetBundleName() string { return "cisco_ios_xe" }

func (input *StartTransaction_Input) GetYangName() string { return "input" }

func (input *StartTransaction_Input) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (input *StartTransaction_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (input *StartTransaction_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (input *StartTransaction_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *StartTransaction_Input) GetParent() types.Entity { return input.parent }

func (input *StartTransaction_Input) GetParentYangName() string { return "start-transaction" }

// StartTransaction_Input_Target
// Name of the configuration datastore towards which the
// transaction is started.
type StartTransaction_Input_Target struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is interface{}.
    Startup interface{}

    // The type is interface{}.
    Running interface{}

    // The type is interface{}.
    Candidate interface{}
}

func (target *StartTransaction_Input_Target) GetFilter() yfilter.YFilter { return target.YFilter }

func (target *StartTransaction_Input_Target) SetFilter(yf yfilter.YFilter) { target.YFilter = yf }

func (target *StartTransaction_Input_Target) GetGoName(yname string) string {
    if yname == "startup" { return "Startup" }
    if yname == "running" { return "Running" }
    if yname == "candidate" { return "Candidate" }
    return ""
}

func (target *StartTransaction_Input_Target) GetSegmentPath() string {
    return "target"
}

func (target *StartTransaction_Input_Target) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (target *StartTransaction_Input_Target) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (target *StartTransaction_Input_Target) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["startup"] = target.Startup
    leafs["running"] = target.Running
    leafs["candidate"] = target.Candidate
    return leafs
}

func (target *StartTransaction_Input_Target) GetBundleName() string { return "cisco_ios_xe" }

func (target *StartTransaction_Input_Target) GetYangName() string { return "target" }

func (target *StartTransaction_Input_Target) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (target *StartTransaction_Input_Target) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (target *StartTransaction_Input_Target) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (target *StartTransaction_Input_Target) SetParent(parent types.Entity) { target.parent = parent }

func (target *StartTransaction_Input_Target) GetParent() types.Entity { return target.parent }

func (target *StartTransaction_Input_Target) GetParentYangName() string { return "input" }

// PrepareTransaction
// Prepares the transaction state for commit.  The server may reject
// the prepare request for any reason, for example due to lack of
// resources or if the combined changes would result in an invalid
// configuration datastore.
// 
// After a successful <prepare-transaction>, the next transaction
// related rpc operation must be <commit-transaction> or
// <abort-transaction>.  Note that an <edit-config> cannot be sent
// before the transaction is either committed or aborted.
// 
// Care must be taken by the server to make sure that if
// <prepare-transaction> succeeds then the <commit-transaction>
// SHOULD not fail, since this might result in an inconsistent
// distributed state.  Thus, <prepare-transaction> should allocate
// any resources needed to make sure the <commit-transaction> will
// succeed.
// 
// If there is no ongoing transaction in this session, or if the
// ongoing transaction already has been prepared, an error MUST be
// returned with <error-app-tag> set to 'bad-state'.
type PrepareTransaction struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (prepareTransaction *PrepareTransaction) GetFilter() yfilter.YFilter { return prepareTransaction.YFilter }

func (prepareTransaction *PrepareTransaction) SetFilter(yf yfilter.YFilter) { prepareTransaction.YFilter = yf }

func (prepareTransaction *PrepareTransaction) GetGoName(yname string) string {
    return ""
}

func (prepareTransaction *PrepareTransaction) GetSegmentPath() string {
    return "tailf-netconf-transactions:prepare-transaction"
}

func (prepareTransaction *PrepareTransaction) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (prepareTransaction *PrepareTransaction) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (prepareTransaction *PrepareTransaction) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (prepareTransaction *PrepareTransaction) GetBundleName() string { return "cisco_ios_xe" }

func (prepareTransaction *PrepareTransaction) GetYangName() string { return "prepare-transaction" }

func (prepareTransaction *PrepareTransaction) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (prepareTransaction *PrepareTransaction) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (prepareTransaction *PrepareTransaction) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (prepareTransaction *PrepareTransaction) SetParent(parent types.Entity) { prepareTransaction.parent = parent }

func (prepareTransaction *PrepareTransaction) GetParent() types.Entity { return prepareTransaction.parent }

func (prepareTransaction *PrepareTransaction) GetParentYangName() string { return "tailf-netconf-transactions" }

// CommitTransaction
// Applies the changes made in the transaction to the configuration
// datatore.  The transaction is closed after a <commit-transaction>.
// 
// If there is no ongoing transaction in this session, or if the
// ongoing transaction already has not been prepared, an error
// MUST be returned with <error-app-tag> set to 'bad-state'.
type CommitTransaction struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (commitTransaction *CommitTransaction) GetFilter() yfilter.YFilter { return commitTransaction.YFilter }

func (commitTransaction *CommitTransaction) SetFilter(yf yfilter.YFilter) { commitTransaction.YFilter = yf }

func (commitTransaction *CommitTransaction) GetGoName(yname string) string {
    return ""
}

func (commitTransaction *CommitTransaction) GetSegmentPath() string {
    return "tailf-netconf-transactions:commit-transaction"
}

func (commitTransaction *CommitTransaction) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (commitTransaction *CommitTransaction) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (commitTransaction *CommitTransaction) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (commitTransaction *CommitTransaction) GetBundleName() string { return "cisco_ios_xe" }

func (commitTransaction *CommitTransaction) GetYangName() string { return "commit-transaction" }

func (commitTransaction *CommitTransaction) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (commitTransaction *CommitTransaction) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (commitTransaction *CommitTransaction) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (commitTransaction *CommitTransaction) SetParent(parent types.Entity) { commitTransaction.parent = parent }

func (commitTransaction *CommitTransaction) GetParent() types.Entity { return commitTransaction.parent }

func (commitTransaction *CommitTransaction) GetParentYangName() string { return "tailf-netconf-transactions" }

// AbortTransaction
// Aborts the ongoing transaction, and all pending changes are
// discarded.  <abort-transaction> can be given at any time during an
// ongoing transaction.
// 
// If there is no ongoing transaction in this session, an error MUST
// be returned with <error-app-tag> set to 'bad-state'.
type AbortTransaction struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (abortTransaction *AbortTransaction) GetFilter() yfilter.YFilter { return abortTransaction.YFilter }

func (abortTransaction *AbortTransaction) SetFilter(yf yfilter.YFilter) { abortTransaction.YFilter = yf }

func (abortTransaction *AbortTransaction) GetGoName(yname string) string {
    return ""
}

func (abortTransaction *AbortTransaction) GetSegmentPath() string {
    return "tailf-netconf-transactions:abort-transaction"
}

func (abortTransaction *AbortTransaction) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (abortTransaction *AbortTransaction) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (abortTransaction *AbortTransaction) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (abortTransaction *AbortTransaction) GetBundleName() string { return "cisco_ios_xe" }

func (abortTransaction *AbortTransaction) GetYangName() string { return "abort-transaction" }

func (abortTransaction *AbortTransaction) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (abortTransaction *AbortTransaction) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (abortTransaction *AbortTransaction) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (abortTransaction *AbortTransaction) SetParent(parent types.Entity) { abortTransaction.parent = parent }

func (abortTransaction *AbortTransaction) GetParent() types.Entity { return abortTransaction.parent }

func (abortTransaction *AbortTransaction) GetParentYangName() string { return "tailf-netconf-transactions" }

