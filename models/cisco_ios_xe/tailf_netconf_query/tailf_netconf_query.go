// This module introduces four new rpc operations to run
// advanced search queries.
// 
// The operation 'start-query' starts a query, with some search
// conditions and control parameters for how to return the results.
// This operation returns a query handle, to be used in subsequent
// operations.
// 
// The operation 'fetch-query-result' is repeatedly to get result
// chunks from the query evaluation.
// 
// The operation 'immediate-query' is used to start a new query, and return
// the entire result.
// 
// The operation 'reset-query' can be used to restart the query.
// 
// Finally 'stop-query' is used to clean up query resources on the
// server.
package tailf_netconf_query

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package tailf_netconf_query"))
    ydk.RegisterEntity("{http://tail-f.com/ns/netconf/query start-query}", reflect.TypeOf(StartQuery{}))
    ydk.RegisterEntity("tailf-netconf-query:start-query", reflect.TypeOf(StartQuery{}))
    ydk.RegisterEntity("{http://tail-f.com/ns/netconf/query fetch-query-result}", reflect.TypeOf(FetchQueryResult{}))
    ydk.RegisterEntity("tailf-netconf-query:fetch-query-result", reflect.TypeOf(FetchQueryResult{}))
    ydk.RegisterEntity("{http://tail-f.com/ns/netconf/query immediate-query}", reflect.TypeOf(ImmediateQuery{}))
    ydk.RegisterEntity("tailf-netconf-query:immediate-query", reflect.TypeOf(ImmediateQuery{}))
    ydk.RegisterEntity("{http://tail-f.com/ns/netconf/query reset-query}", reflect.TypeOf(ResetQuery{}))
    ydk.RegisterEntity("tailf-netconf-query:reset-query", reflect.TypeOf(ResetQuery{}))
    ydk.RegisterEntity("{http://tail-f.com/ns/netconf/query stop-query}", reflect.TypeOf(StopQuery{}))
    ydk.RegisterEntity("tailf-netconf-query:stop-query", reflect.TypeOf(StopQuery{}))
}

// StartQuery
type StartQuery struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input StartQuery_Input

    
    Output StartQuery_Output
}

func (startQuery *StartQuery) GetEntityData() *types.CommonEntityData {
    startQuery.EntityData.YFilter = startQuery.YFilter
    startQuery.EntityData.YangName = "start-query"
    startQuery.EntityData.BundleName = "cisco_ios_xe"
    startQuery.EntityData.ParentYangName = "tailf-netconf-query"
    startQuery.EntityData.SegmentPath = "tailf-netconf-query:start-query"
    startQuery.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    startQuery.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    startQuery.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    startQuery.EntityData.Children = types.NewOrderedMap()
    startQuery.EntityData.Children.Append("input", types.YChild{"Input", &startQuery.Input})
    startQuery.EntityData.Children.Append("output", types.YChild{"Output", &startQuery.Output})
    startQuery.EntityData.Leafs = types.NewOrderedMap()

    startQuery.EntityData.YListKeys = []string {}

    return &(startQuery.EntityData)
}

// StartQuery_Input
type StartQuery_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An XPath 1.0 expression that returns a node set.  For each node in this
    // node set, a 'result' entry is constructed.  For each such node all
    // 'select/expression's are evaluated, and stored in 'result/select'.  The
    // resulting entries are returned from the 'fetch-query-result' function. 
    // When this XPath expression is evaluated, the context node is the root node
    // of the requested data store. The type is string. This attribute is
    // mandatory.
    Foreach interface{}

    // It is possible to sort the result using an ordered list of XPath
    // expressions.  For each node in the node set returned by 'foreach', all
    // 'sort-by' expressions are evaluated, in order, with the node from the
    // 'foreach' evaluation as context node, and the result is stored in a tuple. 
    // Thus, this tuple has as many elements as entries in the 'sort-by' leaf
    // list.  Each expression should return a node set where the first node should
    // be a leaf.  The value of this leaf is used in the tuple.  If the expression
    // returns something else, the value in the tuple is undefined.  When the
    // 'result' list is fetched, is is sorted according to the associated tuple.
    // The type is slice of string.
    SortBy []interface{}

    // The maximum number of 'result' entries to return in each call to
    // 'fetch-query-result'.  If this parameter is not given, all entries are
    // returned. The type is interface{} with range: 1..4294967295.
    Limit interface{}

    // The type is interface{} with range: 1..4294967295. The default value is 1.
    Offset interface{}

    // The maximum time (in seconds) before a query times out. Resets every new
    // request, i.e. subsequent function calls starts a new timeout timer. The
    // type is interface{} with range: 1..4294967295. The default value is 600.
    Timeout interface{}

    // A list of expressions that define what to return from each node in the node
    // set returned by the 'foreach' expression. The type is slice of
    // StartQuery_Input_Select.
    Select []*StartQuery_Input_Select
}

func (input *StartQuery_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xe"
    input.EntityData.ParentYangName = "start-query"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Children.Append("select", types.YChild{"Select", nil})
    for i := range input.Select {
        input.EntityData.Children.Append(types.GetSegmentPath(input.Select[i]), types.YChild{"Select", input.Select[i]})
    }
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("foreach", types.YLeaf{"Foreach", input.Foreach})
    input.EntityData.Leafs.Append("sort-by", types.YLeaf{"SortBy", input.SortBy})
    input.EntityData.Leafs.Append("limit", types.YLeaf{"Limit", input.Limit})
    input.EntityData.Leafs.Append("offset", types.YLeaf{"Offset", input.Offset})
    input.EntityData.Leafs.Append("timeout", types.YLeaf{"Timeout", input.Timeout})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// StartQuery_Input_Select
// A list of expressions that define what to return from each
// node in the node set returned by the 'foreach' expression.
type StartQuery_Input_Select struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Optional label which is copied as is to the 'result' list; can be used for
    // easy labeling of the returned node(s). The type is string.
    Label interface{}

    // Declare what node(s) you want to retrieve.  This XPath expression is
    // evaluated once for every node in the node set returned by the 'foreach'
    // expression.  That node is the inital context node when this expression is
    // evaluated. The type is string. This attribute is mandatory.
    Expression interface{}

    // Controls how the result of the select expression is returned in
    // 'fetch-query-result'. The type is slice of ResultType.
    ResultType []interface{}
}

func (self *StartQuery_Input_Select) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "select"
    self.EntityData.BundleName = "cisco_ios_xe"
    self.EntityData.ParentYangName = "input"
    self.EntityData.SegmentPath = "select"
    self.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("label", types.YLeaf{"Label", self.Label})
    self.EntityData.Leafs.Append("expression", types.YLeaf{"Expression", self.Expression})
    self.EntityData.Leafs.Append("result-type", types.YLeaf{"ResultType", self.ResultType})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// StartQuery_Input_Select_ResultType represents in 'fetch-query-result'.
type StartQuery_Input_Select_ResultType string

const (
    // Return the result of evaluating the expression as if it
    // was surrounded by a call to the xpath function string().
    StartQuery_Input_Select_ResultType_string_ StartQuery_Input_Select_ResultType = "string"

    // If the result is a node set, return the path to the
    // first node in the node set as an instance-identifier.
    // 
    // If the result is not a node set, nothing is returned
    // for this expression.
    StartQuery_Input_Select_ResultType_path StartQuery_Input_Select_ResultType = "path"

    // If the result is a node set, return the value of the
    // first node in the node set, if the first node is a leaf.
    // Otherwise, nothing is returned for this expression.
    StartQuery_Input_Select_ResultType_leaf_value StartQuery_Input_Select_ResultType = "leaf-value"

    // The result is returned inline, i.e., a deep structure
    // of XML (or other API dependent format, e.g., JSON)
    StartQuery_Input_Select_ResultType_inline StartQuery_Input_Select_ResultType = "inline"
)

// StartQuery_Output
type StartQuery_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    QueryHandle interface{}
}

func (output *StartQuery_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "cisco_ios_xe"
    output.EntityData.ParentYangName = "start-query"
    output.EntityData.SegmentPath = "output"
    output.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    output.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    output.EntityData.Children = types.NewOrderedMap()
    output.EntityData.Leafs = types.NewOrderedMap()
    output.EntityData.Leafs.Append("query-handle", types.YLeaf{"QueryHandle", output.QueryHandle})

    output.EntityData.YListKeys = []string {}

    return &(output.EntityData)
}

// FetchQueryResult
type FetchQueryResult struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input FetchQueryResult_Input

    
    Output FetchQueryResult_Output
}

func (fetchQueryResult *FetchQueryResult) GetEntityData() *types.CommonEntityData {
    fetchQueryResult.EntityData.YFilter = fetchQueryResult.YFilter
    fetchQueryResult.EntityData.YangName = "fetch-query-result"
    fetchQueryResult.EntityData.BundleName = "cisco_ios_xe"
    fetchQueryResult.EntityData.ParentYangName = "tailf-netconf-query"
    fetchQueryResult.EntityData.SegmentPath = "tailf-netconf-query:fetch-query-result"
    fetchQueryResult.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    fetchQueryResult.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    fetchQueryResult.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    fetchQueryResult.EntityData.Children = types.NewOrderedMap()
    fetchQueryResult.EntityData.Children.Append("input", types.YChild{"Input", &fetchQueryResult.Input})
    fetchQueryResult.EntityData.Children.Append("output", types.YChild{"Output", &fetchQueryResult.Output})
    fetchQueryResult.EntityData.Leafs = types.NewOrderedMap()

    fetchQueryResult.EntityData.YListKeys = []string {}

    return &(fetchQueryResult.EntityData)
}

// FetchQueryResult_Input
type FetchQueryResult_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    QueryHandle interface{}
}

func (input *FetchQueryResult_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xe"
    input.EntityData.ParentYangName = "fetch-query-result"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("query-handle", types.YLeaf{"QueryHandle", input.QueryHandle})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// FetchQueryResult_Output
type FetchQueryResult_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    QueryResult FetchQueryResult_Output_QueryResult
}

func (output *FetchQueryResult_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "cisco_ios_xe"
    output.EntityData.ParentYangName = "fetch-query-result"
    output.EntityData.SegmentPath = "output"
    output.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    output.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    output.EntityData.Children = types.NewOrderedMap()
    output.EntityData.Children.Append("query-result", types.YChild{"QueryResult", &output.QueryResult})
    output.EntityData.Leafs = types.NewOrderedMap()

    output.EntityData.YListKeys = []string {}

    return &(output.EntityData)
}

// FetchQueryResult_Output_QueryResult
type FetchQueryResult_Output_QueryResult struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // There will be one result for each node in the node set produced by
    // evaluating the 'foreach' expression. The type is slice of
    // FetchQueryResult_Output_QueryResult_Result.
    Result []*FetchQueryResult_Output_QueryResult_Result
}

func (queryResult *FetchQueryResult_Output_QueryResult) GetEntityData() *types.CommonEntityData {
    queryResult.EntityData.YFilter = queryResult.YFilter
    queryResult.EntityData.YangName = "query-result"
    queryResult.EntityData.BundleName = "cisco_ios_xe"
    queryResult.EntityData.ParentYangName = "output"
    queryResult.EntityData.SegmentPath = "query-result"
    queryResult.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    queryResult.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    queryResult.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    queryResult.EntityData.Children = types.NewOrderedMap()
    queryResult.EntityData.Children.Append("result", types.YChild{"Result", nil})
    for i := range queryResult.Result {
        queryResult.EntityData.Children.Append(types.GetSegmentPath(queryResult.Result[i]), types.YChild{"Result", queryResult.Result[i]})
    }
    queryResult.EntityData.Leafs = types.NewOrderedMap()

    queryResult.EntityData.YListKeys = []string {}

    return &(queryResult.EntityData)
}

// FetchQueryResult_Output_QueryResult_Result
// There will be one result for each node in the node set
// produced by evaluating the 'foreach' expression.
type FetchQueryResult_Output_QueryResult_Result struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of FetchQueryResult_Output_QueryResult_Result_Select.
    Select []*FetchQueryResult_Output_QueryResult_Result_Select
}

func (result *FetchQueryResult_Output_QueryResult_Result) GetEntityData() *types.CommonEntityData {
    result.EntityData.YFilter = result.YFilter
    result.EntityData.YangName = "result"
    result.EntityData.BundleName = "cisco_ios_xe"
    result.EntityData.ParentYangName = "query-result"
    result.EntityData.SegmentPath = "result"
    result.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    result.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    result.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    result.EntityData.Children = types.NewOrderedMap()
    result.EntityData.Children.Append("select", types.YChild{"Select", nil})
    for i := range result.Select {
        result.EntityData.Children.Append(types.GetSegmentPath(result.Select[i]), types.YChild{"Select", result.Select[i]})
    }
    result.EntityData.Leafs = types.NewOrderedMap()

    result.EntityData.YListKeys = []string {}

    return &(result.EntityData)
}

// FetchQueryResult_Output_QueryResult_Result_Select
type FetchQueryResult_Output_QueryResult_Result_Select struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Present if the label was given in the input select entry. The type is
    // string.
    Label interface{}

    // The type is string.
    Path interface{}

    // The type is string.
    Value interface{}

    // A deep structure of XML (or other API dependent format, e.g., JSON). The
    // type is string.
    Data interface{}
}

func (self *FetchQueryResult_Output_QueryResult_Result_Select) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "select"
    self.EntityData.BundleName = "cisco_ios_xe"
    self.EntityData.ParentYangName = "result"
    self.EntityData.SegmentPath = "select"
    self.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("label", types.YLeaf{"Label", self.Label})
    self.EntityData.Leafs.Append("path", types.YLeaf{"Path", self.Path})
    self.EntityData.Leafs.Append("value", types.YLeaf{"Value", self.Value})
    self.EntityData.Leafs.Append("data", types.YLeaf{"Data", self.Data})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// ImmediateQuery
type ImmediateQuery struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input ImmediateQuery_Input

    
    Output ImmediateQuery_Output
}

func (immediateQuery *ImmediateQuery) GetEntityData() *types.CommonEntityData {
    immediateQuery.EntityData.YFilter = immediateQuery.YFilter
    immediateQuery.EntityData.YangName = "immediate-query"
    immediateQuery.EntityData.BundleName = "cisco_ios_xe"
    immediateQuery.EntityData.ParentYangName = "tailf-netconf-query"
    immediateQuery.EntityData.SegmentPath = "tailf-netconf-query:immediate-query"
    immediateQuery.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    immediateQuery.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    immediateQuery.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    immediateQuery.EntityData.Children = types.NewOrderedMap()
    immediateQuery.EntityData.Children.Append("input", types.YChild{"Input", &immediateQuery.Input})
    immediateQuery.EntityData.Children.Append("output", types.YChild{"Output", &immediateQuery.Output})
    immediateQuery.EntityData.Leafs = types.NewOrderedMap()

    immediateQuery.EntityData.YListKeys = []string {}

    return &(immediateQuery.EntityData)
}

// ImmediateQuery_Input
type ImmediateQuery_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An XPath 1.0 expression that returns a node set.  For each node in this
    // node set, a 'result' entry is constructed.  For each such node all
    // 'select/expression's are evaluated, and stored in 'result/select'.  The
    // resulting entries are returned from the 'fetch-query-result' function. 
    // When this XPath expression is evaluated, the context node is the root node
    // of the requested data store. The type is string. This attribute is
    // mandatory.
    Foreach interface{}

    // It is possible to sort the result using an ordered list of XPath
    // expressions.  For each node in the node set returned by 'foreach', all
    // 'sort-by' expressions are evaluated, in order, with the node from the
    // 'foreach' evaluation as context node, and the result is stored in a tuple. 
    // Thus, this tuple has as many elements as entries in the 'sort-by' leaf
    // list.  Each expression should return a node set where the first node should
    // be a leaf.  The value of this leaf is used in the tuple.  If the expression
    // returns something else, the value in the tuple is undefined.  When the
    // 'result' list is fetched, is is sorted according to the associated tuple.
    // The type is slice of string.
    SortBy []interface{}

    // The maximum number of 'result' entries to return in each call to
    // 'fetch-query-result'.  If this parameter is not given, all entries are
    // returned. The type is interface{} with range: 1..4294967295.
    Limit interface{}

    // The type is interface{} with range: 1..4294967295. The default value is 1.
    Offset interface{}

    // The maximum time (in seconds) before a query times out. Resets every new
    // request, i.e. subsequent function calls starts a new timeout timer. The
    // type is interface{} with range: 1..4294967295. The default value is 600.
    Timeout interface{}

    // A list of expressions that define what to return from each node in the node
    // set returned by the 'foreach' expression. The type is slice of
    // ImmediateQuery_Input_Select.
    Select []*ImmediateQuery_Input_Select
}

func (input *ImmediateQuery_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xe"
    input.EntityData.ParentYangName = "immediate-query"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Children.Append("select", types.YChild{"Select", nil})
    for i := range input.Select {
        input.EntityData.Children.Append(types.GetSegmentPath(input.Select[i]), types.YChild{"Select", input.Select[i]})
    }
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("foreach", types.YLeaf{"Foreach", input.Foreach})
    input.EntityData.Leafs.Append("sort-by", types.YLeaf{"SortBy", input.SortBy})
    input.EntityData.Leafs.Append("limit", types.YLeaf{"Limit", input.Limit})
    input.EntityData.Leafs.Append("offset", types.YLeaf{"Offset", input.Offset})
    input.EntityData.Leafs.Append("timeout", types.YLeaf{"Timeout", input.Timeout})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// ImmediateQuery_Input_Select
// A list of expressions that define what to return from each
// node in the node set returned by the 'foreach' expression.
type ImmediateQuery_Input_Select struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Optional label which is copied as is to the 'result' list; can be used for
    // easy labeling of the returned node(s). The type is string.
    Label interface{}

    // Declare what node(s) you want to retrieve.  This XPath expression is
    // evaluated once for every node in the node set returned by the 'foreach'
    // expression.  That node is the inital context node when this expression is
    // evaluated. The type is string. This attribute is mandatory.
    Expression interface{}

    // Controls how the result of the select expression is returned in
    // 'fetch-query-result'. The type is slice of ResultType.
    ResultType []interface{}
}

func (self *ImmediateQuery_Input_Select) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "select"
    self.EntityData.BundleName = "cisco_ios_xe"
    self.EntityData.ParentYangName = "input"
    self.EntityData.SegmentPath = "select"
    self.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("label", types.YLeaf{"Label", self.Label})
    self.EntityData.Leafs.Append("expression", types.YLeaf{"Expression", self.Expression})
    self.EntityData.Leafs.Append("result-type", types.YLeaf{"ResultType", self.ResultType})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// ImmediateQuery_Input_Select_ResultType represents in 'fetch-query-result'.
type ImmediateQuery_Input_Select_ResultType string

const (
    // Return the result of evaluating the expression as if it
    // was surrounded by a call to the xpath function string().
    ImmediateQuery_Input_Select_ResultType_string_ ImmediateQuery_Input_Select_ResultType = "string"

    // If the result is a node set, return the path to the
    // first node in the node set as an instance-identifier.
    // 
    // If the result is not a node set, nothing is returned
    // for this expression.
    ImmediateQuery_Input_Select_ResultType_path ImmediateQuery_Input_Select_ResultType = "path"

    // If the result is a node set, return the value of the
    // first node in the node set, if the first node is a leaf.
    // Otherwise, nothing is returned for this expression.
    ImmediateQuery_Input_Select_ResultType_leaf_value ImmediateQuery_Input_Select_ResultType = "leaf-value"

    // The result is returned inline, i.e., a deep structure
    // of XML (or other API dependent format, e.g., JSON)
    ImmediateQuery_Input_Select_ResultType_inline ImmediateQuery_Input_Select_ResultType = "inline"
)

// ImmediateQuery_Output
type ImmediateQuery_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    QueryResult ImmediateQuery_Output_QueryResult
}

func (output *ImmediateQuery_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "cisco_ios_xe"
    output.EntityData.ParentYangName = "immediate-query"
    output.EntityData.SegmentPath = "output"
    output.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    output.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    output.EntityData.Children = types.NewOrderedMap()
    output.EntityData.Children.Append("query-result", types.YChild{"QueryResult", &output.QueryResult})
    output.EntityData.Leafs = types.NewOrderedMap()

    output.EntityData.YListKeys = []string {}

    return &(output.EntityData)
}

// ImmediateQuery_Output_QueryResult
type ImmediateQuery_Output_QueryResult struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // There will be one result for each node in the node set produced by
    // evaluating the 'foreach' expression. The type is slice of
    // ImmediateQuery_Output_QueryResult_Result.
    Result []*ImmediateQuery_Output_QueryResult_Result
}

func (queryResult *ImmediateQuery_Output_QueryResult) GetEntityData() *types.CommonEntityData {
    queryResult.EntityData.YFilter = queryResult.YFilter
    queryResult.EntityData.YangName = "query-result"
    queryResult.EntityData.BundleName = "cisco_ios_xe"
    queryResult.EntityData.ParentYangName = "output"
    queryResult.EntityData.SegmentPath = "query-result"
    queryResult.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    queryResult.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    queryResult.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    queryResult.EntityData.Children = types.NewOrderedMap()
    queryResult.EntityData.Children.Append("result", types.YChild{"Result", nil})
    for i := range queryResult.Result {
        queryResult.EntityData.Children.Append(types.GetSegmentPath(queryResult.Result[i]), types.YChild{"Result", queryResult.Result[i]})
    }
    queryResult.EntityData.Leafs = types.NewOrderedMap()

    queryResult.EntityData.YListKeys = []string {}

    return &(queryResult.EntityData)
}

// ImmediateQuery_Output_QueryResult_Result
// There will be one result for each node in the node set
// produced by evaluating the 'foreach' expression.
type ImmediateQuery_Output_QueryResult_Result struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of ImmediateQuery_Output_QueryResult_Result_Select.
    Select []*ImmediateQuery_Output_QueryResult_Result_Select
}

func (result *ImmediateQuery_Output_QueryResult_Result) GetEntityData() *types.CommonEntityData {
    result.EntityData.YFilter = result.YFilter
    result.EntityData.YangName = "result"
    result.EntityData.BundleName = "cisco_ios_xe"
    result.EntityData.ParentYangName = "query-result"
    result.EntityData.SegmentPath = "result"
    result.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    result.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    result.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    result.EntityData.Children = types.NewOrderedMap()
    result.EntityData.Children.Append("select", types.YChild{"Select", nil})
    for i := range result.Select {
        result.EntityData.Children.Append(types.GetSegmentPath(result.Select[i]), types.YChild{"Select", result.Select[i]})
    }
    result.EntityData.Leafs = types.NewOrderedMap()

    result.EntityData.YListKeys = []string {}

    return &(result.EntityData)
}

// ImmediateQuery_Output_QueryResult_Result_Select
type ImmediateQuery_Output_QueryResult_Result_Select struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Present if the label was given in the input select entry. The type is
    // string.
    Label interface{}

    // The type is string.
    Path interface{}

    // The type is string.
    Value interface{}

    // A deep structure of XML (or other API dependent format, e.g., JSON). The
    // type is string.
    Data interface{}
}

func (self *ImmediateQuery_Output_QueryResult_Result_Select) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "select"
    self.EntityData.BundleName = "cisco_ios_xe"
    self.EntityData.ParentYangName = "result"
    self.EntityData.SegmentPath = "select"
    self.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("label", types.YLeaf{"Label", self.Label})
    self.EntityData.Leafs.Append("path", types.YLeaf{"Path", self.Path})
    self.EntityData.Leafs.Append("value", types.YLeaf{"Value", self.Value})
    self.EntityData.Leafs.Append("data", types.YLeaf{"Data", self.Data})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// ResetQuery
type ResetQuery struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input ResetQuery_Input
}

func (resetQuery *ResetQuery) GetEntityData() *types.CommonEntityData {
    resetQuery.EntityData.YFilter = resetQuery.YFilter
    resetQuery.EntityData.YangName = "reset-query"
    resetQuery.EntityData.BundleName = "cisco_ios_xe"
    resetQuery.EntityData.ParentYangName = "tailf-netconf-query"
    resetQuery.EntityData.SegmentPath = "tailf-netconf-query:reset-query"
    resetQuery.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    resetQuery.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    resetQuery.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    resetQuery.EntityData.Children = types.NewOrderedMap()
    resetQuery.EntityData.Children.Append("input", types.YChild{"Input", &resetQuery.Input})
    resetQuery.EntityData.Leafs = types.NewOrderedMap()

    resetQuery.EntityData.YListKeys = []string {}

    return &(resetQuery.EntityData)
}

// ResetQuery_Input
type ResetQuery_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    QueryHandle interface{}

    // The type is interface{} with range: 1..4294967295. The default value is 1.
    Offset interface{}

    // The maximum time (in seconds) before a query times out. Resets every new
    // request, i.e. subsequent function calls starts a new timeout timer. The
    // type is interface{} with range: 1..4294967295. The default value is 600.
    Timeout interface{}
}

func (input *ResetQuery_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xe"
    input.EntityData.ParentYangName = "reset-query"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("query-handle", types.YLeaf{"QueryHandle", input.QueryHandle})
    input.EntityData.Leafs.Append("offset", types.YLeaf{"Offset", input.Offset})
    input.EntityData.Leafs.Append("timeout", types.YLeaf{"Timeout", input.Timeout})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// StopQuery
type StopQuery struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input StopQuery_Input
}

func (stopQuery *StopQuery) GetEntityData() *types.CommonEntityData {
    stopQuery.EntityData.YFilter = stopQuery.YFilter
    stopQuery.EntityData.YangName = "stop-query"
    stopQuery.EntityData.BundleName = "cisco_ios_xe"
    stopQuery.EntityData.ParentYangName = "tailf-netconf-query"
    stopQuery.EntityData.SegmentPath = "tailf-netconf-query:stop-query"
    stopQuery.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stopQuery.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stopQuery.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stopQuery.EntityData.Children = types.NewOrderedMap()
    stopQuery.EntityData.Children.Append("input", types.YChild{"Input", &stopQuery.Input})
    stopQuery.EntityData.Leafs = types.NewOrderedMap()

    stopQuery.EntityData.YListKeys = []string {}

    return &(stopQuery.EntityData)
}

// StopQuery_Input
type StopQuery_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    QueryHandle interface{}
}

func (input *StopQuery_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xe"
    input.EntityData.ParentYangName = "stop-query"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("query-handle", types.YLeaf{"QueryHandle", input.QueryHandle})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

