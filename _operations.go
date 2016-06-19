package mturk

// NOTE: This code has been adapted from the generated output
// of gowsdl (github.com/hooklift/gowsdl). See comments in wsdltypes.go
// re possible changes to gowsdl. Eventual goal is to auto-generate
// code that can be auto-generated, but for now I'm just focussing on
// getting some code to work.

import "github.com/jjeffery/mturk/soap"

type MTurk struct {
	client *soap.Client
}

func NewMTurk(url string, tls bool, auth *soap.BasicAuth) *MTurk {
	if url == "" {
		url = ""
	}
	client := soap.NewClient(url, tls, auth)

	return &MTurk{
		client: client,
	}
}

func (service *MTurk) CreateHIT(request *CreateHIT) (*CreateHITResponse, error) {
	response := new(CreateHITResponse)
	err := service.client.Call("http://soap.amazon.com", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MTurk) RegisterHITType(request *RegisterHITType) (*RegisterHITTypeResponse, error) {
	response := new(RegisterHITTypeResponse)
	err := service.client.Call("http://soap.amazon.com", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MTurk) SetHITTypeNotification(request *SetHITTypeNotification) (*SetHITTypeNotificationResponse, error) {
	response := new(SetHITTypeNotificationResponse)
	err := service.client.Call("http://soap.amazon.com", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MTurk) SendTestEventNotification(request *SendTestEventNotification) (*SendTestEventNotificationResponse, error) {
	response := new(SendTestEventNotificationResponse)
	err := service.client.Call("http://soap.amazon.com", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MTurk) DisposeHIT(request *DisposeHIT) (*DisposeHITResponse, error) {
	response := new(DisposeHITResponse)
	err := service.client.Call("http://soap.amazon.com", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MTurk) DisableHIT(request *DisableHIT) (*DisableHITResponse, error) {
	response := new(DisableHITResponse)
	err := service.client.Call("http://soap.amazon.com", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MTurk) GetHIT(request *GetHIT) (*GetHITResponse, error) {
	response := new(GetHITResponse)
	err := service.client.Call("http://soap.amazon.com", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MTurk) GetAssignment(request *GetAssignment) (*GetAssignmentResponse, error) {
	response := new(GetAssignmentResponse)
	err := service.client.Call("http://soap.amazon.com", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MTurk) GetReviewResultsForHIT(request *GetReviewResultsForHIT) (*GetReviewResultsForHITResponse, error) {
	response := new(GetReviewResultsForHITResponse)
	err := service.client.Call("http://soap.amazon.com", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MTurk) GetReviewableHITs(request *GetReviewableHITs) (*GetReviewableHITsResponse, error) {
	response := new(GetReviewableHITsResponse)
	err := service.client.Call("http://soap.amazon.com", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MTurk) GetHITsForQualificationType(request *GetHITsForQualificationType) (*GetHITsForQualificationTypeResponse, error) {
	response := new(GetHITsForQualificationTypeResponse)
	err := service.client.Call("http://soap.amazon.com", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MTurk) GetQualificationsForQualificationType(request *GetQualificationsForQualificationType) (*GetQualificationsForQualificationTypeResponse, error) {
	response := new(GetQualificationsForQualificationTypeResponse)
	err := service.client.Call("http://soap.amazon.com", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MTurk) SetHITAsReviewing(request *SetHITAsReviewing) (*SetHITAsReviewingResponse, error) {
	response := new(SetHITAsReviewingResponse)
	err := service.client.Call("http://soap.amazon.com", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MTurk) ExtendHIT(request *ExtendHIT) (*ExtendHITResponse, error) {
	response := new(ExtendHITResponse)
	err := service.client.Call("http://soap.amazon.com", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MTurk) ForceExpireHIT(request *ForceExpireHIT) (*ForceExpireHITResponse, error) {
	response := new(ForceExpireHITResponse)
	err := service.client.Call("http://soap.amazon.com", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MTurk) ApproveAssignment(request *ApproveAssignment) (*ApproveAssignmentResponse, error) {
	response := new(ApproveAssignmentResponse)
	err := service.client.Call("http://soap.amazon.com", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MTurk) RejectAssignment(request *RejectAssignment) (*RejectAssignmentResponse, error) {
	response := new(RejectAssignmentResponse)
	err := service.client.Call("http://soap.amazon.com", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MTurk) ApproveRejectedAssignment(request *ApproveRejectedAssignment) (*ApproveRejectedAssignmentResponse, error) {
	response := new(ApproveRejectedAssignmentResponse)
	err := service.client.Call("http://soap.amazon.com", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MTurk) GetAssignmentsForHIT(request *GetAssignmentsForHIT) (*GetAssignmentsForHITResponse, error) {
	response := new(GetAssignmentsForHITResponse)
	err := service.client.Call("http://soap.amazon.com", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MTurk) GetFileUploadURL(request *GetFileUploadURL) (*GetFileUploadURLResponse, error) {
	response := new(GetFileUploadURLResponse)
	err := service.client.Call("http://soap.amazon.com", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MTurk) SearchHITs(request *SearchHITs) (*SearchHITsResponse, error) {
	response := new(SearchHITsResponse)
	err := service.client.Call("http://soap.amazon.com", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MTurk) GrantBonus(request *GrantBonus) (*GrantBonusResponse, error) {
	response := new(GrantBonusResponse)
	err := service.client.Call("http://soap.amazon.com", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MTurk) GetBonusPayments(request *GetBonusPayments) (*GetBonusPaymentsResponse, error) {
	response := new(GetBonusPaymentsResponse)
	err := service.client.Call("http://soap.amazon.com", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MTurk) ChangeHITTypeOfHIT(request *ChangeHITTypeOfHIT) (*ChangeHITTypeOfHITResponse, error) {
	response := new(ChangeHITTypeOfHITResponse)
	err := service.client.Call("http://soap.amazon.com", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MTurk) CreateQualificationType(request *CreateQualificationType) (*CreateQualificationTypeResponse, error) {
	response := new(CreateQualificationTypeResponse)
	err := service.client.Call("http://soap.amazon.com", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MTurk) GrantQualification(request *GrantQualification) (*GrantQualificationResponse, error) {
	response := new(GrantQualificationResponse)
	err := service.client.Call("http://soap.amazon.com", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MTurk) AssignQualification(request *AssignQualification) (*AssignQualificationResponse, error) {
	response := new(AssignQualificationResponse)
	err := service.client.Call("http://soap.amazon.com", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MTurk) RevokeQualification(request *RevokeQualification) (*RevokeQualificationResponse, error) {
	response := new(RevokeQualificationResponse)
	err := service.client.Call("http://soap.amazon.com", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MTurk) GetQualificationType(request *GetQualificationType) (*GetQualificationTypeResponse, error) {
	response := new(GetQualificationTypeResponse)
	err := service.client.Call("http://soap.amazon.com", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MTurk) GetQualificationRequests(request *GetQualificationRequests) (*GetQualificationRequestsResponse, error) {
	response := new(GetQualificationRequestsResponse)
	err := service.client.Call("http://soap.amazon.com", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MTurk) RejectQualificationRequest(request *RejectQualificationRequest) (*RejectQualificationRequestResponse, error) {
	response := new(RejectQualificationRequestResponse)
	err := service.client.Call("http://soap.amazon.com", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MTurk) UpdateQualificationType(request *UpdateQualificationType) (*UpdateQualificationTypeResponse, error) {
	response := new(UpdateQualificationTypeResponse)
	err := service.client.Call("http://soap.amazon.com", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MTurk) SearchQualificationTypes(request *SearchQualificationTypes) (*SearchQualificationTypesResponse, error) {
	response := new(SearchQualificationTypesResponse)
	err := service.client.Call("http://soap.amazon.com", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MTurk) GetQualificationScore(request *GetQualificationScore) (*GetQualificationScoreResponse, error) {
	response := new(GetQualificationScoreResponse)
	err := service.client.Call("http://soap.amazon.com", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MTurk) UpdateQualificationScore(request *UpdateQualificationScore) (*UpdateQualificationScoreResponse, error) {
	response := new(UpdateQualificationScoreResponse)
	err := service.client.Call("http://soap.amazon.com", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MTurk) DisposeQualificationType(request *DisposeQualificationType) (*DisposeQualificationTypeResponse, error) {
	response := new(DisposeQualificationTypeResponse)
	err := service.client.Call("http://soap.amazon.com", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MTurk) GetRequesterStatistic(request *GetRequesterStatistic) (*GetRequesterStatisticResponse, error) {
	response := new(GetRequesterStatisticResponse)
	err := service.client.Call("http://soap.amazon.com", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MTurk) GetRequesterWorkerStatistic(request *GetRequesterWorkerStatistic) (*GetRequesterWorkerStatisticResponse, error) {
	response := new(GetRequesterWorkerStatisticResponse)
	err := service.client.Call("http://soap.amazon.com", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MTurk) NotifyWorkers(request *NotifyWorkers) (*NotifyWorkersResponse, error) {
	response := new(NotifyWorkersResponse)
	err := service.client.Call("http://soap.amazon.com", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MTurk) GetAccountBalance(request *GetAccountBalance) (*GetAccountBalanceResponse, error) {
	response := new(GetAccountBalanceResponse)
	err := service.client.Call("http://soap.amazon.com", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MTurk) GetBlockedWorkers(request *GetBlockedWorkers) (*GetBlockedWorkersResponse, error) {
	response := new(GetBlockedWorkersResponse)
	err := service.client.Call("http://soap.amazon.com", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MTurk) BlockWorker(request *BlockWorker) (*BlockWorkerResponse, error) {
	response := new(BlockWorkerResponse)
	err := service.client.Call("http://soap.amazon.com", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MTurk) UnblockWorker(request *UnblockWorker) (*UnblockWorkerResponse, error) {
	response := new(UnblockWorkerResponse)
	err := service.client.Call("http://soap.amazon.com", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MTurk) Help(request *Help) (*HelpResponse, error) {
	response := new(HelpResponse)
	err := service.client.Call("http://soap.amazon.com", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
