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
    ydk.RegisterEntity("{http://tail-f.com/ns/netconf/query reset-query}", reflect.TypeOf(ResetQuery{}))
    ydk.RegisterEntity("tailf-netconf-query:reset-query", reflect.TypeOf(ResetQuery{}))
    ydk.RegisterEntity("{http://tail-f.com/ns/netconf/query stop-query}", reflect.TypeOf(StopQuery{}))
    ydk.RegisterEntity("tailf-netconf-query:stop-query", reflect.TypeOf(StopQuery{}))
}

// StartQuery
type StartQuery struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input StartQuery_Input

    
    Output StartQuery_Output
}

func (startQuery *StartQuery) GetFilter() yfilter.YFilter { return startQuery.YFilter }

func (startQuery *StartQuery) SetFilter(yf yfilter.YFilter) { startQuery.YFilter = yf }

func (startQuery *StartQuery) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    if yname == "output" { return "Output" }
    return ""
}

func (startQuery *StartQuery) GetSegmentPath() string {
    return "tailf-netconf-query:start-query"
}

func (startQuery *StartQuery) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &startQuery.Input
    }
    if childYangName == "output" {
        return &startQuery.Output
    }
    return nil
}

func (startQuery *StartQuery) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &startQuery.Input
    children["output"] = &startQuery.Output
    return children
}

func (startQuery *StartQuery) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (startQuery *StartQuery) GetBundleName() string { return "cisco_ios_xe" }

func (startQuery *StartQuery) GetYangName() string { return "start-query" }

func (startQuery *StartQuery) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (startQuery *StartQuery) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (startQuery *StartQuery) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (startQuery *StartQuery) SetParent(parent types.Entity) { startQuery.parent = parent }

func (startQuery *StartQuery) GetParent() types.Entity { return startQuery.parent }

func (startQuery *StartQuery) GetParentYangName() string { return "tailf-netconf-query" }

// StartQuery_Input
type StartQuery_Input struct {
    parent types.Entity
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
    Select []StartQuery_Input_Select
}

func (input *StartQuery_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *StartQuery_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *StartQuery_Input) GetGoName(yname string) string {
    if yname == "foreach" { return "Foreach" }
    if yname == "sort-by" { return "SortBy" }
    if yname == "limit" { return "Limit" }
    if yname == "offset" { return "Offset" }
    if yname == "timeout" { return "Timeout" }
    if yname == "select" { return "Select" }
    return ""
}

func (input *StartQuery_Input) GetSegmentPath() string {
    return "input"
}

func (input *StartQuery_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "select" {
        for _, c := range input.Select {
            if input.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := StartQuery_Input_Select{}
        input.Select = append(input.Select, child)
        return &input.Select[len(input.Select)-1]
    }
    return nil
}

func (input *StartQuery_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range input.Select {
        children[input.Select[i].GetSegmentPath()] = &input.Select[i]
    }
    return children
}

func (input *StartQuery_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["foreach"] = input.Foreach
    leafs["sort-by"] = input.SortBy
    leafs["limit"] = input.Limit
    leafs["offset"] = input.Offset
    leafs["timeout"] = input.Timeout
    return leafs
}

func (input *StartQuery_Input) GetBundleName() string { return "cisco_ios_xe" }

func (input *StartQuery_Input) GetYangName() string { return "input" }

func (input *StartQuery_Input) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (input *StartQuery_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (input *StartQuery_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (input *StartQuery_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *StartQuery_Input) GetParent() types.Entity { return input.parent }

func (input *StartQuery_Input) GetParentYangName() string { return "start-query" }

// StartQuery_Input_Select
// A list of expressions that define what to return from each
// node in the node set returned by the 'foreach' expression.
type StartQuery_Input_Select struct {
    parent types.Entity
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

func (self *StartQuery_Input_Select) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *StartQuery_Input_Select) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *StartQuery_Input_Select) GetGoName(yname string) string {
    if yname == "label" { return "Label" }
    if yname == "expression" { return "Expression" }
    if yname == "result-type" { return "ResultType" }
    return ""
}

func (self *StartQuery_Input_Select) GetSegmentPath() string {
    return "select"
}

func (self *StartQuery_Input_Select) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *StartQuery_Input_Select) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *StartQuery_Input_Select) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["label"] = self.Label
    leafs["expression"] = self.Expression
    leafs["result-type"] = self.ResultType
    return leafs
}

func (self *StartQuery_Input_Select) GetBundleName() string { return "cisco_ios_xe" }

func (self *StartQuery_Input_Select) GetYangName() string { return "select" }

func (self *StartQuery_Input_Select) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (self *StartQuery_Input_Select) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (self *StartQuery_Input_Select) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (self *StartQuery_Input_Select) SetParent(parent types.Entity) { self.parent = parent }

func (self *StartQuery_Input_Select) GetParent() types.Entity { return self.parent }

func (self *StartQuery_Input_Select) GetParentYangName() string { return "input" }

// StartQuery_Input_Select_ResultType represents in 'fetch-query-result'.
type StartQuery_Input_Select_ResultType string

const (
    // Return the result of evaluating the expression as if it
    // was surrounded by a call to the xpath function string().
    StartQuery_Input_Select_ResultType_string StartQuery_Input_Select_ResultType = "string"

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
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    QueryHandle interface{}
}

func (output *StartQuery_Output) GetFilter() yfilter.YFilter { return output.YFilter }

func (output *StartQuery_Output) SetFilter(yf yfilter.YFilter) { output.YFilter = yf }

func (output *StartQuery_Output) GetGoName(yname string) string {
    if yname == "query-handle" { return "QueryHandle" }
    return ""
}

func (output *StartQuery_Output) GetSegmentPath() string {
    return "output"
}

func (output *StartQuery_Output) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (output *StartQuery_Output) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (output *StartQuery_Output) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["query-handle"] = output.QueryHandle
    return leafs
}

func (output *StartQuery_Output) GetBundleName() string { return "cisco_ios_xe" }

func (output *StartQuery_Output) GetYangName() string { return "output" }

func (output *StartQuery_Output) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (output *StartQuery_Output) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (output *StartQuery_Output) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (output *StartQuery_Output) SetParent(parent types.Entity) { output.parent = parent }

func (output *StartQuery_Output) GetParent() types.Entity { return output.parent }

func (output *StartQuery_Output) GetParentYangName() string { return "start-query" }

// FetchQueryResult
type FetchQueryResult struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input FetchQueryResult_Input

    
    Output FetchQueryResult_Output
}

func (fetchQueryResult *FetchQueryResult) GetFilter() yfilter.YFilter { return fetchQueryResult.YFilter }

func (fetchQueryResult *FetchQueryResult) SetFilter(yf yfilter.YFilter) { fetchQueryResult.YFilter = yf }

func (fetchQueryResult *FetchQueryResult) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    if yname == "output" { return "Output" }
    return ""
}

func (fetchQueryResult *FetchQueryResult) GetSegmentPath() string {
    return "tailf-netconf-query:fetch-query-result"
}

func (fetchQueryResult *FetchQueryResult) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &fetchQueryResult.Input
    }
    if childYangName == "output" {
        return &fetchQueryResult.Output
    }
    return nil
}

func (fetchQueryResult *FetchQueryResult) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &fetchQueryResult.Input
    children["output"] = &fetchQueryResult.Output
    return children
}

func (fetchQueryResult *FetchQueryResult) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (fetchQueryResult *FetchQueryResult) GetBundleName() string { return "cisco_ios_xe" }

func (fetchQueryResult *FetchQueryResult) GetYangName() string { return "fetch-query-result" }

func (fetchQueryResult *FetchQueryResult) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (fetchQueryResult *FetchQueryResult) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (fetchQueryResult *FetchQueryResult) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (fetchQueryResult *FetchQueryResult) SetParent(parent types.Entity) { fetchQueryResult.parent = parent }

func (fetchQueryResult *FetchQueryResult) GetParent() types.Entity { return fetchQueryResult.parent }

func (fetchQueryResult *FetchQueryResult) GetParentYangName() string { return "tailf-netconf-query" }

// FetchQueryResult_Input
type FetchQueryResult_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    QueryHandle interface{}
}

func (input *FetchQueryResult_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *FetchQueryResult_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *FetchQueryResult_Input) GetGoName(yname string) string {
    if yname == "query-handle" { return "QueryHandle" }
    return ""
}

func (input *FetchQueryResult_Input) GetSegmentPath() string {
    return "input"
}

func (input *FetchQueryResult_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *FetchQueryResult_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *FetchQueryResult_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["query-handle"] = input.QueryHandle
    return leafs
}

func (input *FetchQueryResult_Input) GetBundleName() string { return "cisco_ios_xe" }

func (input *FetchQueryResult_Input) GetYangName() string { return "input" }

func (input *FetchQueryResult_Input) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (input *FetchQueryResult_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (input *FetchQueryResult_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (input *FetchQueryResult_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *FetchQueryResult_Input) GetParent() types.Entity { return input.parent }

func (input *FetchQueryResult_Input) GetParentYangName() string { return "fetch-query-result" }

// FetchQueryResult_Output
type FetchQueryResult_Output struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    QueryResult FetchQueryResult_Output_QueryResult
}

func (output *FetchQueryResult_Output) GetFilter() yfilter.YFilter { return output.YFilter }

func (output *FetchQueryResult_Output) SetFilter(yf yfilter.YFilter) { output.YFilter = yf }

func (output *FetchQueryResult_Output) GetGoName(yname string) string {
    if yname == "query-result" { return "QueryResult" }
    return ""
}

func (output *FetchQueryResult_Output) GetSegmentPath() string {
    return "output"
}

func (output *FetchQueryResult_Output) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "query-result" {
        return &output.QueryResult
    }
    return nil
}

func (output *FetchQueryResult_Output) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["query-result"] = &output.QueryResult
    return children
}

func (output *FetchQueryResult_Output) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (output *FetchQueryResult_Output) GetBundleName() string { return "cisco_ios_xe" }

func (output *FetchQueryResult_Output) GetYangName() string { return "output" }

func (output *FetchQueryResult_Output) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (output *FetchQueryResult_Output) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (output *FetchQueryResult_Output) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (output *FetchQueryResult_Output) SetParent(parent types.Entity) { output.parent = parent }

func (output *FetchQueryResult_Output) GetParent() types.Entity { return output.parent }

func (output *FetchQueryResult_Output) GetParentYangName() string { return "fetch-query-result" }

// FetchQueryResult_Output_QueryResult
type FetchQueryResult_Output_QueryResult struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // There will be one result for each node in the node set produced by
    // evaluating the 'foreach' expression. The type is slice of
    // FetchQueryResult_Output_QueryResult_Result.
    Result []FetchQueryResult_Output_QueryResult_Result
}

func (queryResult *FetchQueryResult_Output_QueryResult) GetFilter() yfilter.YFilter { return queryResult.YFilter }

func (queryResult *FetchQueryResult_Output_QueryResult) SetFilter(yf yfilter.YFilter) { queryResult.YFilter = yf }

func (queryResult *FetchQueryResult_Output_QueryResult) GetGoName(yname string) string {
    if yname == "result" { return "Result" }
    return ""
}

func (queryResult *FetchQueryResult_Output_QueryResult) GetSegmentPath() string {
    return "query-result"
}

func (queryResult *FetchQueryResult_Output_QueryResult) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "result" {
        for _, c := range queryResult.Result {
            if queryResult.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := FetchQueryResult_Output_QueryResult_Result{}
        queryResult.Result = append(queryResult.Result, child)
        return &queryResult.Result[len(queryResult.Result)-1]
    }
    return nil
}

func (queryResult *FetchQueryResult_Output_QueryResult) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range queryResult.Result {
        children[queryResult.Result[i].GetSegmentPath()] = &queryResult.Result[i]
    }
    return children
}

func (queryResult *FetchQueryResult_Output_QueryResult) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (queryResult *FetchQueryResult_Output_QueryResult) GetBundleName() string { return "cisco_ios_xe" }

func (queryResult *FetchQueryResult_Output_QueryResult) GetYangName() string { return "query-result" }

func (queryResult *FetchQueryResult_Output_QueryResult) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (queryResult *FetchQueryResult_Output_QueryResult) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (queryResult *FetchQueryResult_Output_QueryResult) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (queryResult *FetchQueryResult_Output_QueryResult) SetParent(parent types.Entity) { queryResult.parent = parent }

func (queryResult *FetchQueryResult_Output_QueryResult) GetParent() types.Entity { return queryResult.parent }

func (queryResult *FetchQueryResult_Output_QueryResult) GetParentYangName() string { return "output" }

// FetchQueryResult_Output_QueryResult_Result
// There will be one result for each node in the node set
// produced by evaluating the 'foreach' expression.
type FetchQueryResult_Output_QueryResult_Result struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is slice of FetchQueryResult_Output_QueryResult_Result_Select.
    Select []FetchQueryResult_Output_QueryResult_Result_Select
}

func (result *FetchQueryResult_Output_QueryResult_Result) GetFilter() yfilter.YFilter { return result.YFilter }

func (result *FetchQueryResult_Output_QueryResult_Result) SetFilter(yf yfilter.YFilter) { result.YFilter = yf }

func (result *FetchQueryResult_Output_QueryResult_Result) GetGoName(yname string) string {
    if yname == "select" { return "Select" }
    return ""
}

func (result *FetchQueryResult_Output_QueryResult_Result) GetSegmentPath() string {
    return "result"
}

func (result *FetchQueryResult_Output_QueryResult_Result) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "select" {
        for _, c := range result.Select {
            if result.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := FetchQueryResult_Output_QueryResult_Result_Select{}
        result.Select = append(result.Select, child)
        return &result.Select[len(result.Select)-1]
    }
    return nil
}

func (result *FetchQueryResult_Output_QueryResult_Result) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range result.Select {
        children[result.Select[i].GetSegmentPath()] = &result.Select[i]
    }
    return children
}

func (result *FetchQueryResult_Output_QueryResult_Result) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (result *FetchQueryResult_Output_QueryResult_Result) GetBundleName() string { return "cisco_ios_xe" }

func (result *FetchQueryResult_Output_QueryResult_Result) GetYangName() string { return "result" }

func (result *FetchQueryResult_Output_QueryResult_Result) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (result *FetchQueryResult_Output_QueryResult_Result) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (result *FetchQueryResult_Output_QueryResult_Result) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (result *FetchQueryResult_Output_QueryResult_Result) SetParent(parent types.Entity) { result.parent = parent }

func (result *FetchQueryResult_Output_QueryResult_Result) GetParent() types.Entity { return result.parent }

func (result *FetchQueryResult_Output_QueryResult_Result) GetParentYangName() string { return "query-result" }

// FetchQueryResult_Output_QueryResult_Result_Select
type FetchQueryResult_Output_QueryResult_Result_Select struct {
    parent types.Entity
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

func (self *FetchQueryResult_Output_QueryResult_Result_Select) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *FetchQueryResult_Output_QueryResult_Result_Select) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *FetchQueryResult_Output_QueryResult_Result_Select) GetGoName(yname string) string {
    if yname == "label" { return "Label" }
    if yname == "path" { return "Path" }
    if yname == "value" { return "Value" }
    if yname == "data" { return "Data" }
    return ""
}

func (self *FetchQueryResult_Output_QueryResult_Result_Select) GetSegmentPath() string {
    return "select"
}

func (self *FetchQueryResult_Output_QueryResult_Result_Select) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *FetchQueryResult_Output_QueryResult_Result_Select) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *FetchQueryResult_Output_QueryResult_Result_Select) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["label"] = self.Label
    leafs["path"] = self.Path
    leafs["value"] = self.Value
    leafs["data"] = self.Data
    return leafs
}

func (self *FetchQueryResult_Output_QueryResult_Result_Select) GetBundleName() string { return "cisco_ios_xe" }

func (self *FetchQueryResult_Output_QueryResult_Result_Select) GetYangName() string { return "select" }

func (self *FetchQueryResult_Output_QueryResult_Result_Select) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (self *FetchQueryResult_Output_QueryResult_Result_Select) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (self *FetchQueryResult_Output_QueryResult_Result_Select) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (self *FetchQueryResult_Output_QueryResult_Result_Select) SetParent(parent types.Entity) { self.parent = parent }

func (self *FetchQueryResult_Output_QueryResult_Result_Select) GetParent() types.Entity { return self.parent }

func (self *FetchQueryResult_Output_QueryResult_Result_Select) GetParentYangName() string { return "result" }

// ResetQuery
type ResetQuery struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input ResetQuery_Input
}

func (resetQuery *ResetQuery) GetFilter() yfilter.YFilter { return resetQuery.YFilter }

func (resetQuery *ResetQuery) SetFilter(yf yfilter.YFilter) { resetQuery.YFilter = yf }

func (resetQuery *ResetQuery) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (resetQuery *ResetQuery) GetSegmentPath() string {
    return "tailf-netconf-query:reset-query"
}

func (resetQuery *ResetQuery) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &resetQuery.Input
    }
    return nil
}

func (resetQuery *ResetQuery) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &resetQuery.Input
    return children
}

func (resetQuery *ResetQuery) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (resetQuery *ResetQuery) GetBundleName() string { return "cisco_ios_xe" }

func (resetQuery *ResetQuery) GetYangName() string { return "reset-query" }

func (resetQuery *ResetQuery) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (resetQuery *ResetQuery) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (resetQuery *ResetQuery) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (resetQuery *ResetQuery) SetParent(parent types.Entity) { resetQuery.parent = parent }

func (resetQuery *ResetQuery) GetParent() types.Entity { return resetQuery.parent }

func (resetQuery *ResetQuery) GetParentYangName() string { return "tailf-netconf-query" }

// ResetQuery_Input
type ResetQuery_Input struct {
    parent types.Entity
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

func (input *ResetQuery_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *ResetQuery_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *ResetQuery_Input) GetGoName(yname string) string {
    if yname == "query-handle" { return "QueryHandle" }
    if yname == "offset" { return "Offset" }
    if yname == "timeout" { return "Timeout" }
    return ""
}

func (input *ResetQuery_Input) GetSegmentPath() string {
    return "input"
}

func (input *ResetQuery_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *ResetQuery_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *ResetQuery_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["query-handle"] = input.QueryHandle
    leafs["offset"] = input.Offset
    leafs["timeout"] = input.Timeout
    return leafs
}

func (input *ResetQuery_Input) GetBundleName() string { return "cisco_ios_xe" }

func (input *ResetQuery_Input) GetYangName() string { return "input" }

func (input *ResetQuery_Input) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (input *ResetQuery_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (input *ResetQuery_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (input *ResetQuery_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *ResetQuery_Input) GetParent() types.Entity { return input.parent }

func (input *ResetQuery_Input) GetParentYangName() string { return "reset-query" }

// StopQuery
type StopQuery struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input StopQuery_Input
}

func (stopQuery *StopQuery) GetFilter() yfilter.YFilter { return stopQuery.YFilter }

func (stopQuery *StopQuery) SetFilter(yf yfilter.YFilter) { stopQuery.YFilter = yf }

func (stopQuery *StopQuery) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (stopQuery *StopQuery) GetSegmentPath() string {
    return "tailf-netconf-query:stop-query"
}

func (stopQuery *StopQuery) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &stopQuery.Input
    }
    return nil
}

func (stopQuery *StopQuery) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &stopQuery.Input
    return children
}

func (stopQuery *StopQuery) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (stopQuery *StopQuery) GetBundleName() string { return "cisco_ios_xe" }

func (stopQuery *StopQuery) GetYangName() string { return "stop-query" }

func (stopQuery *StopQuery) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stopQuery *StopQuery) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stopQuery *StopQuery) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stopQuery *StopQuery) SetParent(parent types.Entity) { stopQuery.parent = parent }

func (stopQuery *StopQuery) GetParent() types.Entity { return stopQuery.parent }

func (stopQuery *StopQuery) GetParentYangName() string { return "tailf-netconf-query" }

// StopQuery_Input
type StopQuery_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    QueryHandle interface{}
}

func (input *StopQuery_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *StopQuery_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *StopQuery_Input) GetGoName(yname string) string {
    if yname == "query-handle" { return "QueryHandle" }
    return ""
}

func (input *StopQuery_Input) GetSegmentPath() string {
    return "input"
}

func (input *StopQuery_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *StopQuery_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *StopQuery_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["query-handle"] = input.QueryHandle
    return leafs
}

func (input *StopQuery_Input) GetBundleName() string { return "cisco_ios_xe" }

func (input *StopQuery_Input) GetYangName() string { return "input" }

func (input *StopQuery_Input) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (input *StopQuery_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (input *StopQuery_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (input *StopQuery_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *StopQuery_Input) GetParent() types.Entity { return input.parent }

func (input *StopQuery_Input) GetParentYangName() string { return "stop-query" }

