# PatchedSettingsRequestFlags

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CoreDefaultAppAccess** | **bool** | Applications with no policies bound can be accessed by any user. | 
**EnterpriseAuditIncludeExpandedDiff** | **bool** | Include additional information in audit logs, may incur a performance penalty. | 
**FlowsContinuousLogin** | **bool** | Upon successful authentication, re-start authentication in other open tabs. | 
**RbacCustomPermissions** | **bool** | Let permission checks decide API access. While enabled, every API request costs an additional database query. | 

## Methods

### NewPatchedSettingsRequestFlags

`func NewPatchedSettingsRequestFlags(coreDefaultAppAccess bool, enterpriseAuditIncludeExpandedDiff bool, flowsContinuousLogin bool, rbacCustomPermissions bool, ) *PatchedSettingsRequestFlags`

NewPatchedSettingsRequestFlags instantiates a new PatchedSettingsRequestFlags object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedSettingsRequestFlagsWithDefaults

`func NewPatchedSettingsRequestFlagsWithDefaults() *PatchedSettingsRequestFlags`

NewPatchedSettingsRequestFlagsWithDefaults instantiates a new PatchedSettingsRequestFlags object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoreDefaultAppAccess

`func (o *PatchedSettingsRequestFlags) GetCoreDefaultAppAccess() bool`

GetCoreDefaultAppAccess returns the CoreDefaultAppAccess field if non-nil, zero value otherwise.

### GetCoreDefaultAppAccessOk

`func (o *PatchedSettingsRequestFlags) GetCoreDefaultAppAccessOk() (*bool, bool)`

GetCoreDefaultAppAccessOk returns a tuple with the CoreDefaultAppAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreDefaultAppAccess

`func (o *PatchedSettingsRequestFlags) SetCoreDefaultAppAccess(v bool)`

SetCoreDefaultAppAccess sets CoreDefaultAppAccess field to given value.


### GetEnterpriseAuditIncludeExpandedDiff

`func (o *PatchedSettingsRequestFlags) GetEnterpriseAuditIncludeExpandedDiff() bool`

GetEnterpriseAuditIncludeExpandedDiff returns the EnterpriseAuditIncludeExpandedDiff field if non-nil, zero value otherwise.

### GetEnterpriseAuditIncludeExpandedDiffOk

`func (o *PatchedSettingsRequestFlags) GetEnterpriseAuditIncludeExpandedDiffOk() (*bool, bool)`

GetEnterpriseAuditIncludeExpandedDiffOk returns a tuple with the EnterpriseAuditIncludeExpandedDiff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseAuditIncludeExpandedDiff

`func (o *PatchedSettingsRequestFlags) SetEnterpriseAuditIncludeExpandedDiff(v bool)`

SetEnterpriseAuditIncludeExpandedDiff sets EnterpriseAuditIncludeExpandedDiff field to given value.


### GetFlowsContinuousLogin

`func (o *PatchedSettingsRequestFlags) GetFlowsContinuousLogin() bool`

GetFlowsContinuousLogin returns the FlowsContinuousLogin field if non-nil, zero value otherwise.

### GetFlowsContinuousLoginOk

`func (o *PatchedSettingsRequestFlags) GetFlowsContinuousLoginOk() (*bool, bool)`

GetFlowsContinuousLoginOk returns a tuple with the FlowsContinuousLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowsContinuousLogin

`func (o *PatchedSettingsRequestFlags) SetFlowsContinuousLogin(v bool)`

SetFlowsContinuousLogin sets FlowsContinuousLogin field to given value.


### GetRbacCustomPermissions

`func (o *PatchedSettingsRequestFlags) GetRbacCustomPermissions() bool`

GetRbacCustomPermissions returns the RbacCustomPermissions field if non-nil, zero value otherwise.

### GetRbacCustomPermissionsOk

`func (o *PatchedSettingsRequestFlags) GetRbacCustomPermissionsOk() (*bool, bool)`

GetRbacCustomPermissionsOk returns a tuple with the RbacCustomPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRbacCustomPermissions

`func (o *PatchedSettingsRequestFlags) SetRbacCustomPermissions(v bool)`

SetRbacCustomPermissions sets RbacCustomPermissions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


