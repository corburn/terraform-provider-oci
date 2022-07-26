// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OperatorAccessControl API
//
// Operator Access Control enables you to control the time duration and the actions an Oracle operator can perform on your Exadata Cloud@Customer infrastructure.
// Using logging service, you can view a near real-time audit report of all actions performed by an Oracle operator.
// Use the table of contents and search tool to explore the OperatorAccessControl API.
//

package operatoraccesscontrol

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AccessRequest An Oracle operator raises access request when they need access to any infrastructure resource governed by Operator Access Control.
// The access request identifies the target resource and the set of operator actions. Access request handling depends upon the Operator Control
// that governs the target resource, and the set of operator actions listed for approval in the access request. If all of the operator actions
// listed in the access request are in the pre-approved list in the Operator Control that governs the target resource, then the access request is
// automatically approved. If not, then the access request requires explicit approval from the approver group specified by the Operator Control governing the target resource.
// You can approve or reject an access request. You can also revoke the approval of an already approved access request. While creating an access request,
// the operator specifies the duration of access. You have the option to approve the entire duration or reduce or even increase the time duration.
// An operator can also request for an extension. The approval for such an extension is processed the same way the original access request was processed.
type AccessRequest struct {

	// The OCID of the access request.
	Id *string `mandatory:"true" json:"id"`

	// Summary comment by the operator creating the access request.
	AccessReasonSummary *string `mandatory:"true" json:"accessReasonSummary"`

	// The OCID of the target resource associated with the access request. The operator raises an access request to get approval to
	// access the target resource.
	ResourceId *string `mandatory:"true" json:"resourceId"`

	// This is an automatic identifier generated by the system which is easier for human comprehension.
	RequestId *string `mandatory:"false" json:"requestId"`

	// A unique identifier associated with the operator who raised the request. This identifier can not be used directly to identify the operator.
	// You need to provide this identifier if you would like Oracle to provide additional information about the operator action within Oracle tenancy.
	OperatorId *string `mandatory:"false" json:"operatorId"`

	// The name of the target resource.
	ResourceName *string `mandatory:"false" json:"resourceName"`

	// The OCID of the compartment that contains the access request.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// resourceType for which the AccessRequest is applicable
	ResourceType ResourceTypesEnum `mandatory:"false" json:"resourceType,omitempty"`

	// List of operator actions for which approval is sought by the operator user.
	ActionRequestsList []string `mandatory:"false" json:"actionRequestsList"`

	// Summary reason for which the operator is requesting access on the target resource.
	Reason *string `mandatory:"false" json:"reason"`

	// Priority assigned to the access request by the operator
	Severity AccessRequestSeveritiesEnum `mandatory:"false" json:"severity,omitempty"`

	// Duration in hours for which access is sought on the target resource.
	Duration *int `mandatory:"false" json:"duration"`

	// Duration in hours for which extension access is sought on the target resource.
	ExtendDuration *int `mandatory:"false" json:"extendDuration"`

	// The OCID of the workflow associated with the access request. This is needed if you want to contact Oracle Support for a stuck access request
	// or for an access request that encounters an internal error.
	WorkflowId []string `mandatory:"false" json:"workflowId"`

	// Whether the access request was automatically approved.
	IsAutoApproved *bool `mandatory:"false" json:"isAutoApproved"`

	// The current state of the AccessRequest.
	LifecycleState AccessRequestLifecycleStatesEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// more in detail about the lifeCycleState.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Time when the access request was created in RFC 3339 (https://tools.ietf.org/html/rfc3339)timestamp format. Example: '2020-05-22T21:10:29.600Z'
	TimeOfCreation *common.SDKTime `mandatory:"false" json:"timeOfCreation"`

	// Time when the access request was last modified in RFC 3339 (https://tools.ietf.org/html/rfc3339)timestamp format. Example: '2020-05-22T21:10:29.600Z'
	TimeOfModification *common.SDKTime `mandatory:"false" json:"timeOfModification"`

	// The time when access request is scheduled to be approved in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.Example: '2020-05-22T21:10:29.600Z'
	TimeOfUserCreation *common.SDKTime `mandatory:"false" json:"timeOfUserCreation"`

	// The OCID of the user that last modified the access request.
	UserId *string `mandatory:"false" json:"userId"`

	// The last recent Comment entered by the approver of the request.
	ApproverComment *string `mandatory:"false" json:"approverComment"`

	// The comment entered by the operator while closing the request.
	ClosureComment *string `mandatory:"false" json:"closureComment"`

	// The OCID of the operator control governing the target resource.
	OpctlId *string `mandatory:"false" json:"opctlId"`

	// Name of the Operator control governing the target resource.
	OpctlName *string `mandatory:"false" json:"opctlName"`

	// System message that will be displayed to the operator at login to the target resource.
	SystemMessage *string `mandatory:"false" json:"systemMessage"`

	// Additional message specific to the access request that can be specified by the approver at the time of approval.
	OpctlAdditionalMessage *string `mandatory:"false" json:"opctlAdditionalMessage"`

	// Specifies the type of auditing to be enabled. There are two levels of auditing: command-level and keystroke-level.
	// By default, auditing is enabled at the command level i.e., each command issued by the operator is audited. When keystroke-level is chosen,
	// in addition to command level logging, key strokes are also logged.
	AuditType []string `mandatory:"false" json:"auditType"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m AccessRequest) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AccessRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingResourceTypesEnum(string(m.ResourceType)); !ok && m.ResourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResourceType: %s. Supported values are: %s.", m.ResourceType, strings.Join(GetResourceTypesEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAccessRequestSeveritiesEnum(string(m.Severity)); !ok && m.Severity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Severity: %s. Supported values are: %s.", m.Severity, strings.Join(GetAccessRequestSeveritiesEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAccessRequestLifecycleStatesEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAccessRequestLifecycleStatesEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
