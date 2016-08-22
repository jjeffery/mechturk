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
func (mt *MechTurk) ApproveAssignment(request *ApproveAssignmentRequest) (*ApproveAssignmentResult, error) {
	response := new(approveAssignmentResponse)
	err := mt.call("ApproveAssignment", request, response)
	if err != nil {
		return nil, err
	}
	if err := checkOperationRequest(response.OperationRequest); err != nil {
		return nil, err
	}
	if response.ApproveAssignmentResult == nil {
		return nil, errMissingResult
	}
	if err := checkRequest(response.ApproveAssignmentResult.Request); err != nil {
		return nil, err
	}
	return response.ApproveAssignmentResult, nil
}

// GetAccountBalance is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_GetAccountBalanceOperation.html.
func (mt *MechTurk) GetAccountBalance(request *GetAccountBalanceRequest) (*GetAccountBalanceResult, error) {
	const operation = "GetAccountBalance"
	response := &getAccountBalanceResponse{}
	if err := mt.call(operation, request, response); err != nil {
		return nil, err
	}
	if err := checkOperationRequest(response.OperationRequest); err != nil {
		return nil, err
	}
	if response.GetAccountBalanceResult == nil {
		return nil, errMissingResult
	}
	if err := checkRequest(response.GetAccountBalanceResult.Request); err != nil {
		return nil, err
	}
	return response.GetAccountBalanceResult, nil
}

// CreateHIT is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_CreateHITOperation.html.
func (mt *MechTurk) CreateHIT(request *CreateHITRequest) (*HIT, error) {
	response := new(createHITResponse)
	err := mt.call("CreateHIT", request, response)
	if err != nil {
		return nil, err
	}
	if err = checkOperationRequest(response.OperationRequest); err != nil {
		return nil, err
	}
	if err := checkRequest(response.HIT.Request); err != nil {
		return nil, err
	}

	return response.HIT, nil
}

// RegisterHITType is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_RegisterHITTypeOperation.html.
func (mt *MechTurk) RegisterHITType(request *RegisterHITTypeRequest) (*RegisterHITTypeResult, error) {
	response := new(registerHITTypeResponse)
	err := mt.call("RegisterHITType", request, response)
	if err != nil {
		return nil, err
	}
	if err := checkOperationRequest(response.OperationRequest); err != nil {
		return nil, err
	}
	if response.RegisterHITTypeResult == nil {
		return nil, errMissingResult
	}
	if err := checkRequest(response.RegisterHITTypeResult.Request); err != nil {
		return nil, err
	}

	return response.RegisterHITTypeResult, nil
}

// SetHITTypeNotification is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_SetHITTypeNotificationOperation.html.
func (mt *MechTurk) SetHITTypeNotification(request *SetHITTypeNotificationRequest) (*SetHITTypeNotificationResult, error) {
	response := new(setHITTypeNotificationResponse)
	err := mt.call("SetHITTypeNotification", request, response)
	if err != nil {
		return nil, err
	}
	if err := checkOperationRequest(response.OperationRequest); err != nil {
		return nil, err
	}
	if response.SetHITTypeNotificationResult == nil {
		return nil, errMissingResult
	}
	if err := checkRequest(response.SetHITTypeNotificationResult.Request); err != nil {
		return nil, err
	}

	return response.SetHITTypeNotificationResult, nil
}

// SendTestEventNotification is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_SendTestEventNotificationOperation.html.
func (mt *MechTurk) SendTestEventNotification(request *SendTestEventNotificationRequest) (*SendTestEventNotificationResult, error) {
	response := new(sendTestEventNotificationResponse)
	err := mt.call("SendTestEventNotification", request, response)
	if err != nil {
		return nil, err
	}
	if err := checkOperationRequest(response.OperationRequest); err != nil {
		return nil, err
	}
	if response.SendTestEventNotificationResult == nil {
		return nil, errMissingResult
	}
	if err := checkRequest(response.SendTestEventNotificationResult.Request); err != nil {
		return nil, err
	}

	return response.SendTestEventNotificationResult, nil
}

// DisposeHIT is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_DisposeHITOperation.html.
func (mt *MechTurk) DisposeHIT(request *DisposeHITRequest) (*DisposeHITResult, error) {
	response := new(disposeHITResponse)
	err := mt.call("DisposeHIT", request, response)
	if err != nil {
		return nil, err
	}
	if err := checkOperationRequest(response.OperationRequest); err != nil {
		return nil, err
	}
	if response.DisposeHITResult == nil {
		return nil, errMissingResult
	}
	if err := checkRequest(response.DisposeHITResult.Request); err != nil {
		return nil, err
	}

	return response.DisposeHITResult, nil
}

// DisableHIT is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_DisableHITOperation.html.
func (mt *MechTurk) DisableHIT(request *DisableHITRequest) (*DisableHITResult, error) {
	response := new(disableHITResponse)
	err := mt.call("DisableHIT", request, response)
	if err != nil {
		return nil, err
	}
	if err := checkOperationRequest(response.OperationRequest); err != nil {
		return nil, err
	}
	if response.DisableHITResult == nil {
		return nil, errMissingResult
	}
	if err := checkRequest(response.DisableHITResult.Request); err != nil {
		return nil, err
	}

	return response.DisableHITResult, nil
}

// GetHIT is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_GetHITOperation.html.
func (mt *MechTurk) GetHIT(request *GetHITRequest) (*HIT, error) {
	response := new(getHITResponse)
	err := mt.call("GetHIT", request, response)
	if err != nil {
		return nil, err
	}
	if err := checkOperationRequest(response.OperationRequest); err != nil {
		return nil, err
	}
	if response.HIT == nil {
		return nil, errMissingResult
	}
	if err := checkRequest(response.HIT.Request); err != nil {
		return nil, err
	}

	return response.HIT, nil
}

// GetAssignment is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_GetAssignmentOperation.html.
func (mt *MechTurk) GetAssignment(request *GetAssignmentRequest) (*GetAssignmentResult, error) {
	response := new(getAssignmentResponse)
	err := mt.call("GetAssignment", request, response)
	if err != nil {
		return nil, err
	}
	if err := checkOperationRequest(response.OperationRequest); err != nil {
		return nil, err
	}
	if response.GetAssignmentResult == nil {
		return nil, errMissingResult
	}
	if err := checkRequest(response.GetAssignmentResult.Request); err != nil {
		return nil, err
	}

	return response.GetAssignmentResult, nil
}

// GetReviewResultsForHIT is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_GetReviewResultsForHitOperation.html.
func (mt *MechTurk) GetReviewResultsForHIT(request *GetReviewResultsForHITRequest) (*GetReviewResultsForHITResult, error) {
	response := new(getReviewResultsForHITResponse)
	err := mt.call("GetReviewResultsForHIT", request, response)
	if err != nil {
		return nil, err
	}
	if err := checkOperationRequest(response.OperationRequest); err != nil {
		return nil, err
	}
	if response.GetReviewResultsForHITResult == nil {
		return nil, errMissingResult
	}
	if err := checkRequest(response.GetReviewResultsForHITResult.Request); err != nil {
		return nil, err
	}

	return response.GetReviewResultsForHITResult, nil
}

// GetReviewableHITs is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_GetReviewableHITsOperation.html.
func (mt *MechTurk) GetReviewableHITs(request *GetReviewableHITsRequest) (*GetReviewableHITsResult, error) {
	response := new(getReviewableHITsResponse)
	err := mt.call("GetReviewableHITs", request, response)
	if err != nil {
		return nil, err
	}
	if err := checkOperationRequest(response.OperationRequest); err != nil {
		return nil, err
	}
	if response.GetReviewableHITsResult == nil {
		return nil, errMissingResult
	}
	if err := checkRequest(response.GetReviewableHITsResult.Request); err != nil {
		return nil, err
	}

	return response.GetReviewableHITsResult, nil
}

// GetHITsForQualificationType is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_GetHITsForQualificationTypeOperation.html.
func (mt *MechTurk) GetHITsForQualificationType(request *GetHITsForQualificationTypeRequest) (*GetHITsForQualificationTypeResult, error) {
	response := new(getHITsForQualificationTypeResponse)
	err := mt.call("GetHITsForQualificationType", request, response)
	if err != nil {
		return nil, err
	}
	if err := checkOperationRequest(response.OperationRequest); err != nil {
		return nil, err
	}
	if response.GetHITsForQualificationTypeResult == nil {
		return nil, errMissingResult
	}
	if err := checkRequest(response.GetHITsForQualificationTypeResult.Request); err != nil {
		return nil, err
	}

	return response.GetHITsForQualificationTypeResult, nil
}

// GetQualificationsForQualificationType is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_GetQualificationsForQualificationTypeOperation.html.
func (mt *MechTurk) GetQualificationsForQualificationType(request *GetQualificationsForQualificationTypeRequest) (*GetQualificationsForQualificationTypeResult, error) {
	response := new(getQualificationsForQualificationTypeResponse)
	err := mt.call("GetQualificationsForQualificationType", request, response)
	if err != nil {
		return nil, err
	}
	if err := checkOperationRequest(response.OperationRequest); err != nil {
		return nil, err
	}
	if response.GetQualificationsForQualificationTypeResult == nil {
		return nil, errMissingResult
	}
	if err := checkRequest(response.GetQualificationsForQualificationTypeResult.Request); err != nil {
		return nil, err
	}

	return response.GetQualificationsForQualificationTypeResult, nil
}

// SetHITAsReviewing is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_SetHITAsReviewingOperation.html.
func (mt *MechTurk) SetHITAsReviewing(request *SetHITAsReviewingRequest) (*SetHITAsReviewingResult, error) {
	response := new(setHITAsReviewingResponse)
	err := mt.call("SetHITAsReviewing", request, response)
	if err != nil {
		return nil, err
	}
	if err := checkOperationRequest(response.OperationRequest); err != nil {
		return nil, err
	}
	if response.SetHITAsReviewingResult == nil {
		return nil, errMissingResult
	}
	if err := checkRequest(response.SetHITAsReviewingResult.Request); err != nil {
		return nil, err
	}

	return response.SetHITAsReviewingResult, nil
}

// ExtendHIT is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_ExtendHITOperation.html.
func (c *MechTurk) ExtendHIT(request *ExtendHITRequest) (*ExtendHITResult, error) {
	response := new(extendHITResponse)
	err := c.call("ExtendHIT", request, response)
	if err != nil {
		return nil, err
	}
	if err := checkOperationRequest(response.OperationRequest); err != nil {
		return nil, err
	}
	if response.ExtendHITResult == nil {
		return nil, errMissingResult
	}
	if err := checkRequest(response.ExtendHITResult.Request); err != nil {
		return nil, err
	}

	return response.ExtendHITResult, nil
}

// ForceExpireHIT is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_ForceExpireHITOperation.html.
func (mt *MechTurk) ForceExpireHIT(request *ForceExpireHITRequest) (*ForceExpireHITResult, error) {
	response := new(forceExpireHITResponse)
	err := mt.call("ForceExpireHIT", request, response)
	if err != nil {
		return nil, err
	}
	if err := checkOperationRequest(response.OperationRequest); err != nil {
		return nil, err
	}
	if response.ForceExpireHITResult == nil {
		return nil, errMissingResult
	}
	if err := checkRequest(response.ForceExpireHITResult.Request); err != nil {
		return nil, err
	}

	return response.ForceExpireHITResult, nil
}

// RejectAssignment is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_RejectAssignmentOperation.html.
func (mt *MechTurk) RejectAssignment(request *RejectAssignmentRequest) (*RejectAssignmentResult, error) {
	response := new(rejectAssignmentResponse)
	err := mt.call("RejectAssignment", request, response)
	if err != nil {
		return nil, err
	}
	if err := checkOperationRequest(response.OperationRequest); err != nil {
		return nil, err
	}
	if response.RejectAssignmentResult == nil {
		return nil, errMissingResult
	}
	if err := checkRequest(response.RejectAssignmentResult.Request); err != nil {
		return nil, err
	}

	return response.RejectAssignmentResult, nil
}

// ApproveRejectedAssignment is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_ApproveRejectedAssignmentOperation.html.
func (mt *MechTurk) ApproveRejectedAssignment(request *ApproveRejectedAssignmentRequest) (*ApproveRejectedAssignmentResult, error) {
	response := new(approveRejectedAssignmentResponse)
	err := mt.call("ApproveRejectedAssignment", request, response)
	if err != nil {
		return nil, err
	}
	if err := checkOperationRequest(response.OperationRequest); err != nil {
		return nil, err
	}
	if response.ApproveRejectedAssignmentResult == nil {
		return nil, errMissingResult
	}
	if err := checkRequest(response.ApproveRejectedAssignmentResult.Request); err != nil {
		return nil, err
	}

	return response.ApproveRejectedAssignmentResult, nil
}

// GetAssignmentsForHIT is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_GetAssignmentsForHITOperation.html.
func (mt *MechTurk) GetAssignmentsForHIT(request *GetAssignmentsForHITRequest) (*GetAssignmentsForHITResult, error) {
	response := new(getAssignmentsForHITResponse)
	err := mt.call("GetAssignmentsForHIT", request, response)
	if err != nil {
		return nil, err
	}
	if err := checkOperationRequest(response.OperationRequest); err != nil {
		return nil, err
	}
	if response.GetAssignmentsForHITResult == nil {
		return nil, errMissingResult
	}
	if err := checkRequest(response.GetAssignmentsForHITResult.Request); err != nil {
		return nil, err
	}

	return response.GetAssignmentsForHITResult, nil
}

// GetFileUploadURL is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_GetFileUploadURLOperation.html.
func (mt *MechTurk) GetFileUploadURL(request *GetFileUploadURLRequest) (*GetFileUploadURLResult, error) {
	response := new(getFileUploadURLResponse)
	err := mt.call("GetFileUploadURL", request, response)
	if err != nil {
		return nil, err
	}
	if err := checkOperationRequest(response.OperationRequest); err != nil {
		return nil, err
	}
	if response.GetFileUploadURLResult == nil {
		return nil, errMissingResult
	}
	if err := checkRequest(response.GetFileUploadURLResult.Request); err != nil {
		return nil, err
	}

	return response.GetFileUploadURLResult, nil
}

// SearchHITs is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_SearchHITsOperation.html.
func (mt *MechTurk) SearchHITs(request *SearchHITsRequest) (*SearchHITsResult, error) {
	response := new(searchHITsResponse)
	err := mt.call("SearchHITs", request, response)
	if err != nil {
		return nil, err
	}
	if err := checkOperationRequest(response.OperationRequest); err != nil {
		return nil, err
	}
	if response.SearchHITsResult == nil {
		return nil, errMissingResult
	}
	if err := checkRequest(response.SearchHITsResult.Request); err != nil {
		return nil, err
	}

	return response.SearchHITsResult, nil
}

// GrantBonus is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_GrantBonusOperation.html.
func (mt *MechTurk) GrantBonus(request *GrantBonusRequest) (*GrantBonusResult, error) {
	response := new(grantBonusResponse)
	err := mt.call("GrantBonus", request, response)
	if err != nil {
		return nil, err
	}
	if err := checkOperationRequest(response.OperationRequest); err != nil {
		return nil, err
	}
	if response.GrantBonusResult == nil {
		return nil, errMissingResult
	}
	if err := checkRequest(response.GrantBonusResult.Request); err != nil {
		return nil, err
	}

	return response.GrantBonusResult, nil
}

// GetBonusPayments is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_GetBonusPaymentsOperation.html.
func (mt *MechTurk) GetBonusPayments(request *GetBonusPaymentsRequest) (*GetBonusPaymentsResult, error) {
	response := new(getBonusPaymentsResponse)
	err := mt.call("GetBonusPayments", request, response)
	if err != nil {
		return nil, err
	}
	if err := checkOperationRequest(response.OperationRequest); err != nil {
		return nil, err
	}
	if response.GetBonusPaymentsResult == nil {
		return nil, errMissingResult
	}
	if err := checkRequest(response.GetBonusPaymentsResult.Request); err != nil {
		return nil, err
	}

	return response.GetBonusPaymentsResult, nil
}

// ChangeHITTypeOfHIT is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_ChangeHITTypeOfHITOperation.html.
func (mt *MechTurk) ChangeHITTypeOfHIT(request *ChangeHITTypeOfHITRequest) (*ChangeHITTypeOfHITResult, error) {
	response := new(changeHITTypeOfHITResponse)
	err := mt.call("ChangeHITTypeOfHIT", request, response)
	if err != nil {
		return nil, err
	}
	if err := checkOperationRequest(response.OperationRequest); err != nil {
		return nil, err
	}
	if response.ChangeHITTypeOfHITResult == nil {
		return nil, errMissingResult
	}
	if err := checkRequest(response.ChangeHITTypeOfHITResult.Request); err != nil {
		return nil, err
	}

	return response.ChangeHITTypeOfHITResult, nil
}

// CreateQualificationType is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_CreateQualificationTypeOperation.html.
func (mt *MechTurk) CreateQualificationType(request *CreateQualificationTypeRequest) (*QualificationType, error) {
	response := new(createQualificationTypeResponse)
	err := mt.call("CreateQualificationType", request, response)
	if err != nil {
		return nil, err
	}
	if err := checkOperationRequest(response.OperationRequest); err != nil {
		return nil, err
	}
	if response.QualificationType == nil {
		return nil, errMissingResult
	}
	if err := checkRequest(response.QualificationType.Request); err != nil {
		return nil, err
	}

	return response.QualificationType, nil
}

// GrantQualification is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_GrantQualificationOperation.html.
func (mt *MechTurk) GrantQualification(request *GrantQualificationRequest) (*GrantQualificationResult, error) {
	response := new(grantQualificationResponse)
	err := mt.call("GrantQualification", request, response)
	if err != nil {
		return nil, err
	}
	if err := checkOperationRequest(response.OperationRequest); err != nil {
		return nil, err
	}
	if response.GrantQualificationResult == nil {
		return nil, errMissingResult
	}
	if err := checkRequest(response.GrantQualificationResult.Request); err != nil {
		return nil, err
	}

	return response.GrantQualificationResult, nil
}

// AssignQualification is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_AssignQualificationOperation.html.
func (mt *MechTurk) AssignQualification(request *AssignQualificationRequest) (*AssignQualificationResult, error) {
	response := new(assignQualificationResponse)
	err := mt.call("AssignQualification", request, response)
	if err != nil {
		return nil, err
	}
	if err := checkOperationRequest(response.OperationRequest); err != nil {
		return nil, err
	}
	if response.AssignQualificationResult == nil {
		return nil, errMissingResult
	}
	if err := checkRequest(response.AssignQualificationResult.Request); err != nil {
		return nil, err
	}

	return response.AssignQualificationResult, nil
}

// RevokeQualification is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_RevokeQualificationOperation.html.
func (mt *MechTurk) RevokeQualification(request *RevokeQualificationRequest) (*RevokeQualificationResult, error) {
	response := new(revokeQualificationResponse)
	err := mt.call("RevokeQualification", request, response)
	if err != nil {
		return nil, err
	}
	if err := checkOperationRequest(response.OperationRequest); err != nil {
		return nil, err
	}
	if response.RevokeQualificationResult == nil {
		return nil, errMissingResult
	}
	if err := checkRequest(response.RevokeQualificationResult.Request); err != nil {
		return nil, err
	}

	return response.RevokeQualificationResult, nil
}

// GetQualificationType is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_GetQualificationTypeOperation.html.
func (mt *MechTurk) GetQualificationType(request *GetQualificationTypeRequest) (*QualificationType, error) {
	response := new(getQualificationTypeResponse)
	err := mt.call("GetQualificationType", request, response)
	if err != nil {
		return nil, err
	}
	if err := checkOperationRequest(response.OperationRequest); err != nil {
		return nil, err
	}
	if response.QualificationType == nil {
		return nil, errMissingResult
	}
	if err := checkRequest(response.QualificationType.Request); err != nil {
		return nil, err
	}

	return response.QualificationType, nil
}

// GetQualificationRequests is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_GetQualificationRequestsOperation.html.
func (mt *MechTurk) GetQualificationRequests(request *GetQualificationRequestsRequest) (*GetQualificationRequestsResult, error) {
	response := new(getQualificationRequestsResponse)
	err := mt.call("GetQualificationRequests", request, response)
	if err != nil {
		return nil, err
	}
	if err := checkOperationRequest(response.OperationRequest); err != nil {
		return nil, err
	}
	if response.GetQualificationRequestsResult == nil {
		return nil, errMissingResult
	}
	if err := checkRequest(response.GetQualificationRequestsResult.Request); err != nil {
		return nil, err
	}

	return response.GetQualificationRequestsResult, nil
}

// RejectQualificationRequest is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_RejectQualificationRequestOperation.html.
func (mt *MechTurk) RejectQualificationRequest(request *RejectQualificationRequestRequest) (*RejectQualificationRequestResult, error) {
	response := new(rejectQualificationRequestResponse)
	err := mt.call("RejectQualificationRequest", request, response)
	if err != nil {
		return nil, err
	}
	if err := checkOperationRequest(response.OperationRequest); err != nil {
		return nil, err
	}
	if response.RejectQualificationRequestResult == nil {
		return nil, errMissingResult
	}
	if err := checkRequest(response.RejectQualificationRequestResult.Request); err != nil {
		return nil, err
	}

	return response.RejectQualificationRequestResult, nil
}

// UpdateQualificationType is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_UpdateQualificationTypeOperation.html.
func (mt *MechTurk) UpdateQualificationType(request *UpdateQualificationTypeRequest) (*QualificationType, error) {
	response := new(updateQualificationTypeResponse)
	err := mt.call("UpdateQualificationType", request, response)
	if err != nil {
		return nil, err
	}
	if err := checkOperationRequest(response.OperationRequest); err != nil {
		return nil, err
	}
	if response.QualificationType == nil {
		return nil, errMissingResult
	}
	if err := checkRequest(response.QualificationType.Request); err != nil {
		return nil, err
	}

	return response.QualificationType, nil
}

// SearchQualificationTypes is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_SearchQualificationTypesOperation.html.
func (mt *MechTurk) SearchQualificationTypes(request *SearchQualificationTypesRequest) (*SearchQualificationTypesResult, error) {
	response := new(searchQualificationTypesResponse)
	err := mt.call("SearchQualificationTypes", request, response)
	if err != nil {
		return nil, err
	}
	if err := checkOperationRequest(response.OperationRequest); err != nil {
		return nil, err
	}
	if response.SearchQualificationTypesResult == nil {
		return nil, errMissingResult
	}
	if err := checkRequest(response.SearchQualificationTypesResult.Request); err != nil {
		return nil, err
	}

	return response.SearchQualificationTypesResult, nil
}

// GetQualificationScore is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_GetQualificationScoreOperation.html.
func (mt *MechTurk) GetQualificationScore(request *GetQualificationScoreRequest) (*Qualification, error) {
	response := new(getQualificationScoreResponse)
	err := mt.call("GetQualificationScore", request, response)
	if err != nil {
		return nil, err
	}
	if err := checkOperationRequest(response.OperationRequest); err != nil {
		return nil, err
	}
	if response.Qualification == nil {
		return nil, errMissingResult
	}
	if err := checkRequest(response.Qualification.Request); err != nil {
		return nil, err
	}

	return response.Qualification, nil
}

// UpdateQualificationScore is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_UpdateQualificationScoreOperation.html.
func (mt *MechTurk) UpdateQualificationScore(request *UpdateQualificationScoreRequest) (*UpdateQualificationScoreResult, error) {
	response := new(updateQualificationScoreResponse)
	err := mt.call("UpdateQualificationScore", request, response)
	if err != nil {
		return nil, err
	}
	if err := checkOperationRequest(response.OperationRequest); err != nil {
		return nil, err
	}
	if response.UpdateQualificationScoreResult == nil {
		return nil, errMissingResult
	}
	if err := checkRequest(response.UpdateQualificationScoreResult.Request); err != nil {
		return nil, err
	}

	return response.UpdateQualificationScoreResult, nil
}

// DisposeQualificationType is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_DisposeQualificationTypeOperation.html.
func (mt *MechTurk) DisposeQualificationType(request *DisposeQualificationTypeRequest) (*DisposeQualificationTypeResult, error) {
	response := new(disposeQualificationTypeResponse)
	err := mt.call("DisposeQualificationType", request, response)
	if err != nil {
		return nil, err
	}
	if err := checkOperationRequest(response.OperationRequest); err != nil {
		return nil, err
	}
	if response.DisposeQualificationTypeResult == nil {
		return nil, errMissingResult
	}
	if err := checkRequest(response.DisposeQualificationTypeResult.Request); err != nil {
		return nil, err
	}

	return response.DisposeQualificationTypeResult, nil
}

// GetRequesterStatistic is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_GetRequesterStatisticOperation.html.
func (mt *MechTurk) GetRequesterStatistic(request *GetRequesterStatisticRequest) (*GetStatisticResult, error) {
	response := new(getRequesterStatisticResponse)
	err := mt.call("GetRequesterStatistic", request, response)
	if err != nil {
		return nil, err
	}
	if err := checkOperationRequest(response.OperationRequest); err != nil {
		return nil, err
	}
	if response.GetStatisticResult == nil {
		return nil, errMissingResult
	}
	if err := checkRequest(response.GetStatisticResult.Request); err != nil {
		return nil, err
	}

	return response.GetStatisticResult, nil
}

// GetRequesterWorkerStatistic is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_GetRequesterWorkerStatisticOperation.html.
func (mt *MechTurk) GetRequesterWorkerStatistic(request *GetRequesterWorkerStatisticRequest) (*GetStatisticResult, error) {
	response := new(getRequesterWorkerStatisticResponse)
	err := mt.call("GetRequesterWorkerStatistic", request, response)
	if err != nil {
		return nil, err
	}
	if err := checkOperationRequest(response.OperationRequest); err != nil {
		return nil, err
	}
	if response.GetStatisticResult == nil {
		return nil, errMissingResult
	}
	if err := checkRequest(response.GetStatisticResult.Request); err != nil {
		return nil, err
	}

	return response.GetStatisticResult, nil
}

// NotifyWorkers is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_NotifyWorkersOperation.html.
func (mt *MechTurk) NotifyWorkers(request *NotifyWorkersRequest) (*NotifyWorkersResult, error) {
	response := new(notifyWorkersResponse)
	err := mt.call("NotifyWorkers", request, response)
	if err != nil {
		return nil, err
	}
	if err := checkOperationRequest(response.OperationRequest); err != nil {
		return nil, err
	}
	if response.NotifyWorkersResult == nil {
		return nil, errMissingResult
	}
	if err := checkRequest(response.NotifyWorkersResult.Request); err != nil {
		return nil, err
	}

	return response.NotifyWorkersResult, nil
}

// GetBlockedWorkers is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_GetBlockedWorkersOperation.html.
func (mt *MechTurk) GetBlockedWorkers(request *GetBlockedWorkersRequest) (*GetBlockedWorkersResult, error) {
	response := new(getBlockedWorkersResponse)
	err := mt.call("GetBlockedWorkers", request, response)
	if err != nil {
		return nil, err
	}
	if err := checkOperationRequest(response.OperationRequest); err != nil {
		return nil, err
	}
	if response.GetBlockedWorkersResult == nil {
		return nil, errMissingResult
	}
	if err := checkRequest(response.GetBlockedWorkersResult.Request); err != nil {
		return nil, err
	}

	return response.GetBlockedWorkersResult, nil
}

// BlockWorker is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_BlockWorkerOperation.html.
func (mt *MechTurk) BlockWorker(request *BlockWorkerRequest) (*BlockWorkerResult, error) {
	response := new(blockWorkerResponse)
	err := mt.call("BlockWorker", request, response)
	if err != nil {
		return nil, err
	}
	if err := checkOperationRequest(response.OperationRequest); err != nil {
		return nil, err
	}
	if response.BlockWorkerResult == nil {
		return nil, errMissingResult
	}
	if err := checkRequest(response.BlockWorkerResult.Request); err != nil {
		return nil, err
	}

	return response.BlockWorkerResult, nil
}

// UnblockWorker is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_UnblockWorkerOperation.html.
func (mt *MechTurk) UnblockWorker(request *UnblockWorkerRequest) (*UnblockWorkerResult, error) {
	response := new(unblockWorkerResponse)
	err := mt.call("UnblockWorker", request, response)
	if err != nil {
		return nil, err
	}
	if err := checkOperationRequest(response.OperationRequest); err != nil {
		return nil, err
	}
	if response.UnblockWorkerResult == nil {
		return nil, errMissingResult
	}
	if err := checkRequest(response.UnblockWorkerResult.Request); err != nil {
		return nil, err
	}

	return response.UnblockWorkerResult, nil
}

// Help is documented at http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_HelpOperation.html.
func (mt *MechTurk) Help(request *HelpRequest) (*Information, error) {
	response := new(helpResponse)
	err := mt.call("Help", request, response)
	if err != nil {
		return nil, err
	}
	if err := checkOperationRequest(response.OperationRequest); err != nil {
		return nil, err
	}
	if response.Information == nil {
		return nil, errMissingResult
	}
	if err := checkRequest(response.Information.Request); err != nil {
		return nil, err
	}

	return response.Information, nil
}
