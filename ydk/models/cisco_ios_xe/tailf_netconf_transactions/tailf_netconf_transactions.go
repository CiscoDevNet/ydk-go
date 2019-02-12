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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input StartTransaction_Input
}

func (startTransaction *StartTransaction) GetEntityData() *types.CommonEntityData {
    startTransaction.EntityData.YFilter = startTransaction.YFilter
    startTransaction.EntityData.YangName = "start-transaction"
    startTransaction.EntityData.BundleName = "cisco_ios_xe"
    startTransaction.EntityData.ParentYangName = "tailf-netconf-transactions"
    startTransaction.EntityData.SegmentPath = "tailf-netconf-transactions:start-transaction"
    startTransaction.EntityData.AbsolutePath = startTransaction.EntityData.SegmentPath
    startTransaction.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    startTransaction.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    startTransaction.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    startTransaction.EntityData.Children = types.NewOrderedMap()
    startTransaction.EntityData.Children.Append("input", types.YChild{"Input", &startTransaction.Input})
    startTransaction.EntityData.Leafs = types.NewOrderedMap()

    startTransaction.EntityData.YListKeys = []string {}

    return &(startTransaction.EntityData)
}

// StartTransaction_Input
type StartTransaction_Input struct {
    EntityData types.CommonEntityData
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

func (input *StartTransaction_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xe"
    input.EntityData.ParentYangName = "start-transaction"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "tailf-netconf-transactions:start-transaction/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Children.Append("target", types.YChild{"Target", &input.Target})
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("with-inactive", types.YLeaf{"WithInactive", input.WithInactive})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// StartTransaction_Input_Target
// Name of the configuration datastore towards which the
// transaction is started.
type StartTransaction_Input_Target struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{}.
    Startup interface{}

    // The type is interface{}.
    Running interface{}

    // The type is interface{}.
    Candidate interface{}
}

func (target *StartTransaction_Input_Target) GetEntityData() *types.CommonEntityData {
    target.EntityData.YFilter = target.YFilter
    target.EntityData.YangName = "target"
    target.EntityData.BundleName = "cisco_ios_xe"
    target.EntityData.ParentYangName = "input"
    target.EntityData.SegmentPath = "target"
    target.EntityData.AbsolutePath = "tailf-netconf-transactions:start-transaction/input/" + target.EntityData.SegmentPath
    target.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    target.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    target.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    target.EntityData.Children = types.NewOrderedMap()
    target.EntityData.Leafs = types.NewOrderedMap()
    target.EntityData.Leafs.Append("startup", types.YLeaf{"Startup", target.Startup})
    target.EntityData.Leafs.Append("running", types.YLeaf{"Running", target.Running})
    target.EntityData.Leafs.Append("candidate", types.YLeaf{"Candidate", target.Candidate})

    target.EntityData.YListKeys = []string {}

    return &(target.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (prepareTransaction *PrepareTransaction) GetEntityData() *types.CommonEntityData {
    prepareTransaction.EntityData.YFilter = prepareTransaction.YFilter
    prepareTransaction.EntityData.YangName = "prepare-transaction"
    prepareTransaction.EntityData.BundleName = "cisco_ios_xe"
    prepareTransaction.EntityData.ParentYangName = "tailf-netconf-transactions"
    prepareTransaction.EntityData.SegmentPath = "tailf-netconf-transactions:prepare-transaction"
    prepareTransaction.EntityData.AbsolutePath = prepareTransaction.EntityData.SegmentPath
    prepareTransaction.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    prepareTransaction.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    prepareTransaction.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    prepareTransaction.EntityData.Children = types.NewOrderedMap()
    prepareTransaction.EntityData.Leafs = types.NewOrderedMap()

    prepareTransaction.EntityData.YListKeys = []string {}

    return &(prepareTransaction.EntityData)
}

// CommitTransaction
// Applies the changes made in the transaction to the configuration
// datatore.  The transaction is closed after a <commit-transaction>.
// 
// If there is no ongoing transaction in this session, or if the
// ongoing transaction already has not been prepared, an error
// MUST be returned with <error-app-tag> set to 'bad-state'.
type CommitTransaction struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (commitTransaction *CommitTransaction) GetEntityData() *types.CommonEntityData {
    commitTransaction.EntityData.YFilter = commitTransaction.YFilter
    commitTransaction.EntityData.YangName = "commit-transaction"
    commitTransaction.EntityData.BundleName = "cisco_ios_xe"
    commitTransaction.EntityData.ParentYangName = "tailf-netconf-transactions"
    commitTransaction.EntityData.SegmentPath = "tailf-netconf-transactions:commit-transaction"
    commitTransaction.EntityData.AbsolutePath = commitTransaction.EntityData.SegmentPath
    commitTransaction.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    commitTransaction.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    commitTransaction.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    commitTransaction.EntityData.Children = types.NewOrderedMap()
    commitTransaction.EntityData.Leafs = types.NewOrderedMap()

    commitTransaction.EntityData.YListKeys = []string {}

    return &(commitTransaction.EntityData)
}

// AbortTransaction
// Aborts the ongoing transaction, and all pending changes are
// discarded.  <abort-transaction> can be given at any time during an
// ongoing transaction.
// 
// If there is no ongoing transaction in this session, an error MUST
// be returned with <error-app-tag> set to 'bad-state'.
type AbortTransaction struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (abortTransaction *AbortTransaction) GetEntityData() *types.CommonEntityData {
    abortTransaction.EntityData.YFilter = abortTransaction.YFilter
    abortTransaction.EntityData.YangName = "abort-transaction"
    abortTransaction.EntityData.BundleName = "cisco_ios_xe"
    abortTransaction.EntityData.ParentYangName = "tailf-netconf-transactions"
    abortTransaction.EntityData.SegmentPath = "tailf-netconf-transactions:abort-transaction"
    abortTransaction.EntityData.AbsolutePath = abortTransaction.EntityData.SegmentPath
    abortTransaction.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    abortTransaction.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    abortTransaction.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    abortTransaction.EntityData.Children = types.NewOrderedMap()
    abortTransaction.EntityData.Leafs = types.NewOrderedMap()

    abortTransaction.EntityData.YListKeys = []string {}

    return &(abortTransaction.EntityData)
}

