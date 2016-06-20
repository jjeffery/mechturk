package mechturk

import (
	"github.com/jjeffery/mechturk/soap"
	"github.com/pkg/errors"
)

// MechTurk is an AWS Mechanical Turk Requester client.
type MechTurk struct {
	Config *Config
}

// New returns a new AWS Mechanical Turk Requester client.
func New(configs ...*Config) *MechTurk {
	config := DefaultConfig.Clone()
	for _, c := range configs {
		config = config.Merge(c)
	}
	return &MechTurk{
		Config: config,
	}
}

// call contains common code for calling a SOAP web service.
func (mt *MechTurk) call(operation string, request interface{}, response interface{}) error {
	creds, err := mt.Config.getCredentials().Get()
	if err != nil {
		return errors.Wrap(err, operation)
	}
	payload := newPayload(operation, request, creds)
	if err != nil {
		return errors.Wrap(err, operation)
	}
	soapClient := soap.NewClient(mt.Config.getEndpoint())
	soapClient.Logger = mt.Config.Logger
	if err := soapClient.Call(soapAction, payload, response); err != nil {
		return errors.Wrap(err, operation)
	}
	return nil
}

// ApproveAssignment is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_ApproveAssignmentOperation.html.
func (mt *MechTurk) ApproveAssignment(request *ApproveAssignmentRequest) (*ApproveAssignmentResponse, error) {
	response := new(ApproveAssignmentResponse)
	err := mt.call("ApproveAssignment", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetAccountBalance is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_GetAccountBalanceOperation.html.
func (mt *MechTurk) GetAccountBalance(request *GetAccountBalanceRequest) (*GetAccountBalanceResponse, error) {
	const operation = "GetAccountBalance"
	response := &GetAccountBalanceResponse{}
	if err := mt.call(operation, request, response); err != nil {
		return nil, err
	}
	return response, nil
}

// CreateHIT is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_CreateHITOperation.html.
func (mt *MechTurk) CreateHIT(request *CreateHITRequest) (*CreateHITResponse, error) {
	response := new(CreateHITResponse)
	err := mt.call("CreateHIT", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// RegisterHITType is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_RegisterHITTypeOperation.html.
func (mt *MechTurk) RegisterHITType(request *RegisterHITTypeRequest) (*RegisterHITTypeResponse, error) {
	response := new(RegisterHITTypeResponse)
	err := mt.call("RegisterHITType", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// SetHITTypeNotification is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_SetHITTypeNotificationOperation.html.
func (mt *MechTurk) SetHITTypeNotification(request *SetHITTypeNotificationRequest) (*SetHITTypeNotificationResponse, error) {
	response := new(SetHITTypeNotificationResponse)
	err := mt.call("SetHITTypeNotification", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// SendTestEventNotification is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_SendTestEventNotificationOperation.html.
func (mt *MechTurk) SendTestEventNotification(request *SendTestEventNotificationRequest) (*SendTestEventNotificationResponse, error) {
	response := new(SendTestEventNotificationResponse)
	err := mt.call("SendTestEventNotification", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// DisposeHIT is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_DisposeHITOperation.html.
func (mt *MechTurk) DisposeHIT(request *DisposeHITRequest) (*DisposeHITResponse, error) {
	response := new(DisposeHITResponse)
	err := mt.call("DisposeHIT", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// DisableHIT is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_DisableHITOperation.html.
func (mt *MechTurk) DisableHIT(request *DisableHITRequest) (*DisableHITResponse, error) {
	response := new(DisableHITResponse)
	err := mt.call("DisableHIT", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetHIT is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_GetHITOperation.html.
func (mt *MechTurk) GetHIT(request *GetHITRequest) (*GetHITResponse, error) {
	response := new(GetHITResponse)
	err := mt.call("GetHIT", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetAssignment is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_GetAssignmentOperation.html.
func (mt *MechTurk) GetAssignment(request *GetAssignmentRequest) (*GetAssignmentResponse, error) {
	response := new(GetAssignmentResponse)
	err := mt.call("GetAssignment", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetReviewResultsForHIT is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_GetReviewResultsForHitOperation.html.
func (mt *MechTurk) GetReviewResultsForHIT(request *GetReviewResultsForHITRequest) (*GetReviewResultsForHITResponse, error) {
	response := new(GetReviewResultsForHITResponse)
	err := mt.call("GetReviewResultsForHIT", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetReviewableHITs is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_GetReviewableHITsOperation.html.
func (mt *MechTurk) GetReviewableHITs(request *GetReviewableHITsRequest) (*GetReviewableHITsResponse, error) {
	response := new(GetReviewableHITsResponse)
	err := mt.call("GetReviewableHITs", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetHITsForQualificationType is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_GetHITsForQualificationTypeOperation.html.
func (mt *MechTurk) GetHITsForQualificationType(request *GetHITsForQualificationTypeRequest) (*GetHITsForQualificationTypeResponse, error) {
	response := new(GetHITsForQualificationTypeResponse)
	err := mt.call("GetHITsForQualificationType", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetQualificationsForQualificationType is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_GetQualificationsForQualificationTypeOperation.html.
func (mt *MechTurk) GetQualificationsForQualificationType(request *GetQualificationsForQualificationTypeRequest) (*GetQualificationsForQualificationTypeResponse, error) {
	response := new(GetQualificationsForQualificationTypeResponse)
	err := mt.call("GetQualificationsForQualificationType", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// SetHITAsReviewing is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_SetHITAsReviewingOperation.html.
func (mt *MechTurk) SetHITAsReviewing(request *SetHITAsReviewingRequest) (*SetHITAsReviewingResponse, error) {
	response := new(SetHITAsReviewingResponse)
	err := mt.call("SetHITAsReviewing", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// ExtendHIT is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_ExtendHITOperation.html.
func (c *MechTurk) ExtendHIT(request *ExtendHITRequest) (*ExtendHITResponse, error) {
	response := new(ExtendHITResponse)
	err := c.call("ExtendHIT", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// ForceExpireHIT is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_ForceExpireHITOperation.html.
func (mt *MechTurk) ForceExpireHIT(request *ForceExpireHITRequest) (*ForceExpireHITResponse, error) {
	response := new(ForceExpireHITResponse)
	err := mt.call("ForceExpireHIT", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// RejectAssignment is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_RejectAssignmentOperation.html.
func (mt *MechTurk) RejectAssignment(request *RejectAssignmentRequest) (*RejectAssignmentResponse, error) {
	response := new(RejectAssignmentResponse)
	err := mt.call("RejectAssignment", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// ApproveRejectedAssignment is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_ApproveRejectedAssignmentOperation.html.
func (mt *MechTurk) ApproveRejectedAssignment(request *ApproveRejectedAssignmentRequest) (*ApproveRejectedAssignmentResponse, error) {
	response := new(ApproveRejectedAssignmentResponse)
	err := mt.call("ApproveRejectedAssignment", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetAssignmentsForHIT is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_GetAssignmentsForHITOperation.html.
func (mt *MechTurk) GetAssignmentsForHIT(request *GetAssignmentsForHITRequest) (*GetAssignmentsForHITResponse, error) {
	response := new(GetAssignmentsForHITResponse)
	err := mt.call("GetAssignmentsForHIT", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetFileUploadURL is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_GetFileUploadURLOperation.html.
func (mt *MechTurk) GetFileUploadURL(request *GetFileUploadURLRequest) (*GetFileUploadURLResponse, error) {
	response := new(GetFileUploadURLResponse)
	err := mt.call("GetFileUploadURL", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// SearchHITs is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_SearchHITsOperation.html.
func (mt *MechTurk) SearchHITs(request *SearchHITsRequest) (*SearchHITsResponse, error) {
	response := new(SearchHITsResponse)
	err := mt.call("SearchHITs", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GrantBonus is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_GrantBonusOperation.html.
func (mt *MechTurk) GrantBonus(request *GrantBonusRequest) (*GrantBonusResponse, error) {
	response := new(GrantBonusResponse)
	err := mt.call("GrantBonus", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetBonusPayments is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_GetBonusPaymentsOperation.html.
func (mt *MechTurk) GetBonusPayments(request *GetBonusPaymentsRequest) (*GetBonusPaymentsResponse, error) {
	response := new(GetBonusPaymentsResponse)
	err := mt.call("GetBonusPayments", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// ChangeHITTypeOfHIT is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_ChangeHITTypeOfHITOperation.html.
func (mt *MechTurk) ChangeHITTypeOfHIT(request *ChangeHITTypeOfHITRequest) (*ChangeHITTypeOfHITResponse, error) {
	response := new(ChangeHITTypeOfHITResponse)
	err := mt.call("ChangeHITTypeOfHIT", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// CreateQualificationType is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_CreateQualificationTypeOperation.html.
func (mt *MechTurk) CreateQualificationType(request *CreateQualificationTypeRequest) (*CreateQualificationTypeResponse, error) {
	response := new(CreateQualificationTypeResponse)
	err := mt.call("CreateQualificationType", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GrantQualification is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_GrantQualificationOperation.html.
func (mt *MechTurk) GrantQualification(request *GrantQualificationRequest) (*GrantQualificationResponse, error) {
	response := new(GrantQualificationResponse)
	err := mt.call("GrantQualification", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AssignQualification is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_AssignQualificationOperation.html.
func (mt *MechTurk) AssignQualification(request *AssignQualificationRequest) (*AssignQualificationResponse, error) {
	response := new(AssignQualificationResponse)
	err := mt.call("AssignQualification", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// RevokeQualification is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_RevokeQualificationOperation.html.
func (mt *MechTurk) RevokeQualification(request *RevokeQualificationRequest) (*RevokeQualificationResponse, error) {
	response := new(RevokeQualificationResponse)
	err := mt.call("RevokeQualification", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetQualificationType is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_GetQualificationTypeOperation.html.
func (mt *MechTurk) GetQualificationType(request *GetQualificationTypeRequest) (*GetQualificationTypeResponse, error) {
	response := new(GetQualificationTypeResponse)
	err := mt.call("GetQualificationType", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetQualificationRequests is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_GetQualificationRequestsOperation.html.
func (mt *MechTurk) GetQualificationRequests(request *GetQualificationRequestsRequest) (*GetQualificationRequestsResponse, error) {
	response := new(GetQualificationRequestsResponse)
	err := mt.call("GetQualificationRequests", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// RejectQualificationRequest is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_RejectQualificationRequestOperation.html.
func (mt *MechTurk) RejectQualificationRequest(request *RejectQualificationRequestRequest) (*RejectQualificationRequestResponse, error) {
	response := new(RejectQualificationRequestResponse)
	err := mt.call("RejectQualificationRequest", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// UpdateQualificationType is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_UpdateQualificationTypeOperation.html.
func (mt *MechTurk) UpdateQualificationType(request *UpdateQualificationTypeRequest) (*UpdateQualificationTypeResponse, error) {
	response := new(UpdateQualificationTypeResponse)
	err := mt.call("UpdateQualificationType", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// SearchQualificationTypes is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_SearchQualificationTypesOperation.html.
func (mt *MechTurk) SearchQualificationTypes(request *SearchQualificationTypesRequest) (*SearchQualificationTypesResponse, error) {
	response := new(SearchQualificationTypesResponse)
	err := mt.call("SearchQualificationTypes", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetQualificationScore is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_GetQualificationScoreOperation.html.
func (mt *MechTurk) GetQualificationScore(request *GetQualificationScoreRequest) (*GetQualificationScoreResponse, error) {
	response := new(GetQualificationScoreResponse)
	err := mt.call("GetQualificationScore", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// UpdateQualificationScore is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_UpdateQualificationScoreOperation.html.
func (mt *MechTurk) UpdateQualificationScore(request *UpdateQualificationScoreRequest) (*UpdateQualificationScoreResponse, error) {
	response := new(UpdateQualificationScoreResponse)
	err := mt.call("UpdateQualificationScore", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// DisposeQualificationType is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_DisposeQualificationTypeOperation.html.
func (mt *MechTurk) DisposeQualificationType(request *DisposeQualificationTypeRequest) (*DisposeQualificationTypeResponse, error) {
	response := new(DisposeQualificationTypeResponse)
	err := mt.call("DisposeQualificationType", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetRequesterStatistic is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_GetRequesterStatisticOperation.html.
func (mt *MechTurk) GetRequesterStatistic(request *GetRequesterStatisticRequest) (*GetRequesterStatisticResponse, error) {
	response := new(GetRequesterStatisticResponse)
	err := mt.call("GetRequesterStatistic", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetRequesterWorkerStatistic is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_GetRequesterWorkerStatisticOperation.html.
func (mt *MechTurk) GetRequesterWorkerStatistic(request *GetRequesterWorkerStatisticRequest) (*GetRequesterWorkerStatisticResponse, error) {
	response := new(GetRequesterWorkerStatisticResponse)
	err := mt.call("GetRequesterWorkerStatistic", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// NotifyWorkers is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_NotifyWorkersOperation.html.
func (mt *MechTurk) NotifyWorkers(request *NotifyWorkersRequest) (*NotifyWorkersResponse, error) {
	response := new(NotifyWorkersResponse)
	err := mt.call("NotifyWorkers", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetBlockedWorkers is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_GetBlockedWorkersOperation.html.
func (mt *MechTurk) GetBlockedWorkers(request *GetBlockedWorkersRequest) (*GetBlockedWorkersResponse, error) {
	response := new(GetBlockedWorkersResponse)
	err := mt.call("GetBlockedWorkers", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// BlockWorker is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_BlockWorkerOperation.html.
func (mt *MechTurk) BlockWorker(request *BlockWorkerRequest) (*BlockWorkerResponse, error) {
	response := new(BlockWorkerResponse)
	err := mt.call("BlockWorker", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// UnblockWorker is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_UnblockWorkerOperation.html.
func (mt *MechTurk) UnblockWorker(request *UnblockWorkerRequest) (*UnblockWorkerResponse, error) {
	response := new(UnblockWorkerResponse)
	err := mt.call("UnblockWorker", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Help is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_HelpOperation.html.
func (mt *MechTurk) Help(request *HelpRequest) (*HelpResponse, error) {
	response := new(HelpResponse)
	err := mt.call("Help", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
