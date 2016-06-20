package mturk

import (
	"github.com/jjeffery/mechturk/soap"
	"github.com/pkg/errors"
)

// Client is an AWS Mechanical Turk Requester client.
type Client struct {
	Config *Config
}

// NewClient returns a new AWS Mechanical Turk Requester client.
func NewClient(configs ...*Config) *Client {
	config := DefaultConfig.Clone()
	for _, c := range configs {
		config = config.Merge(c)
	}
	return &Client{
		Config: config,
	}
}

func (c *Client) call(operation string, request interface{}, response interface{}) error {
	creds, err := c.Config.getCredentials().Get()
	if err != nil {
		return errors.Wrap(err, operation)
	}
	payload := newPayload(operation, request, creds)
	if err != nil {
		return errors.Wrap(err, operation)
	}
	soap := soap.NewClient(c.Config.getEndpoint())
	soap.Logger = c.Config.Logger
	if err := soap.Call(soapAction, payload, response); err != nil {
		return errors.Wrap(err, operation)
	}
	return nil
}

// ApproveAssignment is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_ApproveAssignmentOperation.html.
func (c *Client) ApproveAssignment(request *ApproveAssignmentRequest) (*ApproveAssignmentResponse, error) {
	response := new(ApproveAssignmentResponse)
	err := c.call("ApproveAssignment", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetAccountBalance is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_GetAccountBalanceOperation.html.
func (c *Client) GetAccountBalance(request *GetAccountBalanceRequest) (*GetAccountBalanceResponse, error) {
	const operation = "GetAccountBalance"
	response := &GetAccountBalanceResponse{}
	if err := c.call(operation, request, response); err != nil {
		return nil, err
	}
	return response, nil
}

// CreateHIT is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_CreateHITOperation.html.
func (c *Client) CreateHIT(request *CreateHITRequest) (*CreateHITResponse, error) {
	response := new(CreateHITResponse)
	err := c.call("CreateHIT", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// RegisterHITType is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_RegisterHITTypeOperation.html.
func (c *Client) RegisterHITType(request *RegisterHITTypeRequest) (*RegisterHITTypeResponse, error) {
	response := new(RegisterHITTypeResponse)
	err := c.call("RegisterHITType", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// SetHITTypeNotification is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_SetHITTypeNotificationOperation.html.
func (c *Client) SetHITTypeNotification(request *SetHITTypeNotificationRequest) (*SetHITTypeNotificationResponse, error) {
	response := new(SetHITTypeNotificationResponse)
	err := c.call("SetHITTypeNotification", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// SendTestEventNotification is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_SendTestEventNotificationOperation.html.
func (c *Client) SendTestEventNotification(request *SendTestEventNotificationRequest) (*SendTestEventNotificationResponse, error) {
	response := new(SendTestEventNotificationResponse)
	err := c.call("SendTestEventNotification", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// DisposeHIT is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_DisposeHITOperation.html.
func (c *Client) DisposeHIT(request *DisposeHITRequest) (*DisposeHITResponse, error) {
	response := new(DisposeHITResponse)
	err := c.call("DisposeHIT", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// DisableHIT is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_DisableHITOperation.html.
func (c *Client) DisableHIT(request *DisableHITRequest) (*DisableHITResponse, error) {
	response := new(DisableHITResponse)
	err := c.call("DisableHIT", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetHIT is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_GetHITOperation.html.
func (c *Client) GetHIT(request *GetHITRequest) (*GetHITResponse, error) {
	response := new(GetHITResponse)
	err := c.call("GetHIT", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetAssignment is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_GetAssignmentOperation.html.
func (c *Client) GetAssignment(request *GetAssignmentRequest) (*GetAssignmentResponse, error) {
	response := new(GetAssignmentResponse)
	err := c.call("GetAssignment", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetReviewResultsForHIT is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_GetReviewResultsForHitOperation.html.
func (c *Client) GetReviewResultsForHIT(request *GetReviewResultsForHITRequest) (*GetReviewResultsForHITResponse, error) {
	response := new(GetReviewResultsForHITResponse)
	err := c.call("GetReviewResultsForHIT", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetReviewableHITs is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_GetReviewableHITsOperation.html.
func (c *Client) GetReviewableHITs(request *GetReviewableHITsRequest) (*GetReviewableHITsResponse, error) {
	response := new(GetReviewableHITsResponse)
	err := c.call("GetReviewableHITs", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetHITsForQualificationType is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_GetHITsForQualificationTypeOperation.html.
func (c *Client) GetHITsForQualificationType(request *GetHITsForQualificationTypeRequest) (*GetHITsForQualificationTypeResponse, error) {
	response := new(GetHITsForQualificationTypeResponse)
	err := c.call("GetHITsForQualificationType", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetQualificationsForQualificationType is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_GetQualificationsForQualificationTypeOperation.html.
func (c *Client) GetQualificationsForQualificationType(request *GetQualificationsForQualificationTypeRequest) (*GetQualificationsForQualificationTypeResponse, error) {
	response := new(GetQualificationsForQualificationTypeResponse)
	err := c.call("GetQualificationsForQualificationType", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// SetHITAsReviewing is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_SetHITAsReviewingOperation.html.
func (c *Client) SetHITAsReviewing(request *SetHITAsReviewingRequest) (*SetHITAsReviewingResponse, error) {
	response := new(SetHITAsReviewingResponse)
	err := c.call("SetHITAsReviewing", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// ExtendHIT is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_ExtendHITOperation.html.
func (c *Client) ExtendHIT(request *ExtendHITRequest) (*ExtendHITResponse, error) {
	response := new(ExtendHITResponse)
	err := c.call("ExtendHIT", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// ForceExpireHIT is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_ForceExpireHITOperation.html.
func (c *Client) ForceExpireHIT(request *ForceExpireHITRequest) (*ForceExpireHITResponse, error) {
	response := new(ForceExpireHITResponse)
	err := c.call("ForceExpireHIT", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// RejectAssignment is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_RejectAssignmentOperation.html.
func (c *Client) RejectAssignment(request *RejectAssignmentRequest) (*RejectAssignmentResponse, error) {
	response := new(RejectAssignmentResponse)
	err := c.call("RejectAssignment", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// ApproveRejectedAssignment is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_ApproveRejectedAssignmentOperation.html.
func (c *Client) ApproveRejectedAssignment(request *ApproveRejectedAssignmentRequest) (*ApproveRejectedAssignmentResponse, error) {
	response := new(ApproveRejectedAssignmentResponse)
	err := c.call("ApproveRejectedAssignment", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetAssignmentsForHIT is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_GetAssignmentsForHITOperation.html.
func (c *Client) GetAssignmentsForHIT(request *GetAssignmentsForHITRequest) (*GetAssignmentsForHITResponse, error) {
	response := new(GetAssignmentsForHITResponse)
	err := c.call("GetAssignmentsForHIT", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetFileUploadURL is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_GetFileUploadURLOperation.html.
func (c *Client) GetFileUploadURL(request *GetFileUploadURLRequest) (*GetFileUploadURLResponse, error) {
	response := new(GetFileUploadURLResponse)
	err := c.call("GetFileUploadURL", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// SearchHITs is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_SearchHITsOperation.html.
func (c *Client) SearchHITs(request *SearchHITsRequest) (*SearchHITsResponse, error) {
	response := new(SearchHITsResponse)
	err := c.call("SearchHITs", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GrantBonus is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_GrantBonusOperation.html.
func (c *Client) GrantBonus(request *GrantBonusRequest) (*GrantBonusResponse, error) {
	response := new(GrantBonusResponse)
	err := c.call("GrantBonus", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetBonusPayments is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_GetBonusPaymentsOperation.html.
func (c *Client) GetBonusPayments(request *GetBonusPaymentsRequest) (*GetBonusPaymentsResponse, error) {
	response := new(GetBonusPaymentsResponse)
	err := c.call("GetBonusPayments", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// ChangeHITTypeOfHIT is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_ChangeHITTypeOfHITOperation.html.
func (c *Client) ChangeHITTypeOfHIT(request *ChangeHITTypeOfHITRequest) (*ChangeHITTypeOfHITResponse, error) {
	response := new(ChangeHITTypeOfHITResponse)
	err := c.call("ChangeHITTypeOfHIT", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// CreateQualificationType is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_CreateQualificationTypeOperation.html.
func (c *Client) CreateQualificationType(request *CreateQualificationTypeRequest) (*CreateQualificationTypeResponse, error) {
	response := new(CreateQualificationTypeResponse)
	err := c.call("CreateQualificationType", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GrantQualification is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_GrantQualificationOperation.html.
func (c *Client) GrantQualification(request *GrantQualificationRequest) (*GrantQualificationResponse, error) {
	response := new(GrantQualificationResponse)
	err := c.call("GrantQualification", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AssignQualification is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_AssignQualificationOperation.html.
func (c *Client) AssignQualification(request *AssignQualificationRequest) (*AssignQualificationResponse, error) {
	response := new(AssignQualificationResponse)
	err := c.call("AssignQualification", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// RevokeQualification is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_RevokeQualificationOperation.html.
func (c *Client) RevokeQualification(request *RevokeQualificationRequest) (*RevokeQualificationResponse, error) {
	response := new(RevokeQualificationResponse)
	err := c.call("RevokeQualification", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetQualificationType is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_GetQualificationTypeOperation.html.
func (c *Client) GetQualificationType(request *GetQualificationTypeRequest) (*GetQualificationTypeResponse, error) {
	response := new(GetQualificationTypeResponse)
	err := c.call("GetQualificationType", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetQualificationRequests is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_GetQualificationRequestsOperation.html.
func (c *Client) GetQualificationRequests(request *GetQualificationRequestsRequest) (*GetQualificationRequestsResponse, error) {
	response := new(GetQualificationRequestsResponse)
	err := c.call("GetQualificationRequests", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// RejectQualificationRequest is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_RejectQualificationRequestOperation.html.
func (c *Client) RejectQualificationRequest(request *RejectQualificationRequestRequest) (*RejectQualificationRequestResponse, error) {
	response := new(RejectQualificationRequestResponse)
	err := c.call("RejectQualificationRequest", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// UpdateQualificationType is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_UpdateQualificationTypeOperation.html.
func (c *Client) UpdateQualificationType(request *UpdateQualificationTypeRequest) (*UpdateQualificationTypeResponse, error) {
	response := new(UpdateQualificationTypeResponse)
	err := c.call("UpdateQualificationType", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// SearchQualificationTypes is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_SearchQualificationTypesOperation.html.
func (c *Client) SearchQualificationTypes(request *SearchQualificationTypesRequest) (*SearchQualificationTypesResponse, error) {
	response := new(SearchQualificationTypesResponse)
	err := c.call("SearchQualificationTypes", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetQualificationScore is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_GetQualificationScoreOperation.html.
func (c *Client) GetQualificationScore(request *GetQualificationScoreRequest) (*GetQualificationScoreResponse, error) {
	response := new(GetQualificationScoreResponse)
	err := c.call("GetQualificationScore", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// UpdateQualificationScore is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_UpdateQualificationScoreOperation.html.
func (c *Client) UpdateQualificationScore(request *UpdateQualificationScoreRequest) (*UpdateQualificationScoreResponse, error) {
	response := new(UpdateQualificationScoreResponse)
	err := c.call("UpdateQualificationScore", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// DisposeQualificationType is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_DisposeQualificationTypeOperation.html.
func (c *Client) DisposeQualificationType(request *DisposeQualificationTypeRequest) (*DisposeQualificationTypeResponse, error) {
	response := new(DisposeQualificationTypeResponse)
	err := c.call("DisposeQualificationType", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetRequesterStatistic is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_GetRequesterStatisticOperation.html.
func (c *Client) GetRequesterStatistic(request *GetRequesterStatisticRequest) (*GetRequesterStatisticResponse, error) {
	response := new(GetRequesterStatisticResponse)
	err := c.call("GetRequesterStatistic", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetRequesterWorkerStatistic is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_GetRequesterWorkerStatisticOperation.html.
func (c *Client) GetRequesterWorkerStatistic(request *GetRequesterWorkerStatisticRequest) (*GetRequesterWorkerStatisticResponse, error) {
	response := new(GetRequesterWorkerStatisticResponse)
	err := c.call("GetRequesterWorkerStatistic", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// NotifyWorkers is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_NotifyWorkersOperation.html.
func (c *Client) NotifyWorkers(request *NotifyWorkersRequest) (*NotifyWorkersResponse, error) {
	response := new(NotifyWorkersResponse)
	err := c.call("NotifyWorkers", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetBlockedWorkers is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_GetBlockedWorkersOperation.html.
func (c *Client) GetBlockedWorkers(request *GetBlockedWorkersRequest) (*GetBlockedWorkersResponse, error) {
	response := new(GetBlockedWorkersResponse)
	err := c.call("GetBlockedWorkers", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// BlockWorker is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_BlockWorkerOperation.html.
func (c *Client) BlockWorker(request *BlockWorkerRequest) (*BlockWorkerResponse, error) {
	response := new(BlockWorkerResponse)
	err := c.call("BlockWorker", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// UnblockWorker is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_UnblockWorkerOperation.html.
func (c *Client) UnblockWorker(request *UnblockWorkerRequest) (*UnblockWorkerResponse, error) {
	response := new(UnblockWorkerResponse)
	err := c.call("UnblockWorker", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Help is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_HelpOperation.html.
func (c *Client) Help(request *HelpRequest) (*HelpResponse, error) {
	response := new(HelpResponse)
	err := c.call("Help", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
