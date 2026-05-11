# CurrentBrandFlags

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CoreDefaultAppAccess** | **bool** | Configure if applications without any policy/group/user bindings should be accessible to any user. | 
**EnterpriseAuditIncludeExpandedDiff** | **bool** | Include additional information in audit logs, may incur a performance penalty. | 
**FlowsContinuousLogin** | **bool** | Upon successful authentication, re-start authentication in other open tabs. | 
**FlowsRefreshOthers** | **bool** | Refresh other tabs after successful authentication. | 

## Methods

### NewCurrentBrandFlags

`func NewCurrentBrandFlags(coreDefaultAppAccess bool, enterpriseAuditIncludeExpandedDiff bool, flowsContinuousLogin bool, flowsRefreshOthers bool, ) *CurrentBrandFlags`

NewCurrentBrandFlags instantiates a new CurrentBrandFlags object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrentBrandFlagsWithDefaults

`func NewCurrentBrandFlagsWithDefaults() *CurrentBrandFlags`

NewCurrentBrandFlagsWithDefaults instantiates a new CurrentBrandFlags object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoreDefaultAppAccess

`func (o *CurrentBrandFlags) GetCoreDefaultAppAccess() bool`

GetCoreDefaultAppAccess returns the CoreDefaultAppAccess field if non-nil, zero value otherwise.

### GetCoreDefaultAppAccessOk

`func (o *CurrentBrandFlags) GetCoreDefaultAppAccessOk() (*bool, bool)`

GetCoreDefaultAppAccessOk returns a tuple with the CoreDefaultAppAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreDefaultAppAccess

`func (o *CurrentBrandFlags) SetCoreDefaultAppAccess(v bool)`

SetCoreDefaultAppAccess sets CoreDefaultAppAccess field to given value.


### GetEnterpriseAuditIncludeExpandedDiff

`func (o *CurrentBrandFlags) GetEnterpriseAuditIncludeExpandedDiff() bool`

GetEnterpriseAuditIncludeExpandedDiff returns the EnterpriseAuditIncludeExpandedDiff field if non-nil, zero value otherwise.

### GetEnterpriseAuditIncludeExpandedDiffOk

`func (o *CurrentBrandFlags) GetEnterpriseAuditIncludeExpandedDiffOk() (*bool, bool)`

GetEnterpriseAuditIncludeExpandedDiffOk returns a tuple with the EnterpriseAuditIncludeExpandedDiff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseAuditIncludeExpandedDiff

`func (o *CurrentBrandFlags) SetEnterpriseAuditIncludeExpandedDiff(v bool)`

SetEnterpriseAuditIncludeExpandedDiff sets EnterpriseAuditIncludeExpandedDiff field to given value.


### GetFlowsContinuousLogin

`func (o *CurrentBrandFlags) GetFlowsContinuousLogin() bool`

GetFlowsContinuousLogin returns the FlowsContinuousLogin field if non-nil, zero value otherwise.

### GetFlowsContinuousLoginOk

`func (o *CurrentBrandFlags) GetFlowsContinuousLoginOk() (*bool, bool)`

GetFlowsContinuousLoginOk returns a tuple with the FlowsContinuousLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowsContinuousLogin

`func (o *CurrentBrandFlags) SetFlowsContinuousLogin(v bool)`

SetFlowsContinuousLogin sets FlowsContinuousLogin field to given value.


### GetFlowsRefreshOthers

`func (o *CurrentBrandFlags) GetFlowsRefreshOthers() bool`

GetFlowsRefreshOthers returns the FlowsRefreshOthers field if non-nil, zero value otherwise.

### GetFlowsRefreshOthersOk

`func (o *CurrentBrandFlags) GetFlowsRefreshOthersOk() (*bool, bool)`

GetFlowsRefreshOthersOk returns a tuple with the FlowsRefreshOthers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowsRefreshOthers

`func (o *CurrentBrandFlags) SetFlowsRefreshOthers(v bool)`

SetFlowsRefreshOthers sets FlowsRefreshOthers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


