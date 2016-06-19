package mturk

import (
	"github.com/jjeffery/mturk/soap"
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
	soap := soap.NewClient(c.Config.getEndpoint(), false, nil)
	soap.Logger = c.Config.Logger
	if err := soap.Call(soapAction, payload, response); err != nil {
		return errors.Wrap(err, operation)
	}
	return nil
}

func (c *Client) GetAccountBalance(request *GetAccountBalanceRequest) (*GetAccountBalanceResponse, error) {
	const operation = "GetAccountBalance"
	response := &GetAccountBalanceResponse{}
	if err := c.call(operation, request, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) CreateHIT(request *CreateHITRequest) (*CreateHITResponse, error) {
	response := new(CreateHITResponse)
	err := c.call("CreateHIT", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) RegisterHITType(request *RegisterHITTypeRequest) (*RegisterHITTypeResponse, error) {
	response := new(RegisterHITTypeResponse)
	err := c.call("RegisterHITType", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) SetHITTypeNotification(request *SetHITTypeNotificationRequest) (*SetHITTypeNotificationResponse, error) {
	response := new(SetHITTypeNotificationResponse)
	err := c.call("SetHITTypeNotification", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) SendTestEventNotification(request *SendTestEventNotificationRequest) (*SendTestEventNotificationResponse, error) {
	response := new(SendTestEventNotificationResponse)
	err := c.call("SendTestEventNotification", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) DisposeHIT(request *DisposeHITRequest) (*DisposeHITResponse, error) {
	response := new(DisposeHITResponse)
	err := c.call("DisposeHIT", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) DisableHIT(request *DisableHITRequest) (*DisableHITResponse, error) {
	response := new(DisableHITResponse)
	err := c.call("DisableHIT", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) GetHIT(request *GetHITRequest) (*GetHITResponse, error) {
	response := new(GetHITResponse)
	err := c.call("GetHIT", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) GetAssignment(request *GetAssignmentRequest) (*GetAssignmentResponse, error) {
	response := new(GetAssignmentResponse)
	err := c.call("GetAssignment", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) GetReviewResultsForHIT(request *GetReviewResultsForHITRequest) (*GetReviewResultsForHITResponse, error) {
	response := new(GetReviewResultsForHITResponse)
	err := c.call("GetReviewResultsForHIT", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) GetReviewableHITs(request *GetReviewableHITsRequest) (*GetReviewableHITsResponse, error) {
	response := new(GetReviewableHITsResponse)
	err := c.call("GetReviewableHITs", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) GetHITsForQualificationType(request *GetHITsForQualificationTypeRequest) (*GetHITsForQualificationTypeResponse, error) {
	response := new(GetHITsForQualificationTypeResponse)
	err := c.call("GetHITsForQualificationType", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) GetQualificationsForQualificationType(request *GetQualificationsForQualificationTypeRequest) (*GetQualificationsForQualificationTypeResponse, error) {
	response := new(GetQualificationsForQualificationTypeResponse)
	err := c.call("GetQualificationsForQualificationType", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) SetHITAsReviewing(request *SetHITAsReviewingRequest) (*SetHITAsReviewingResponse, error) {
	response := new(SetHITAsReviewingResponse)
	err := c.call("SetHITAsReviewing", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) ExtendHIT(request *ExtendHITRequest) (*ExtendHITResponse, error) {
	response := new(ExtendHITResponse)
	err := c.call("ExtendHIT", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) ForceExpireHIT(request *ForceExpireHITRequest) (*ForceExpireHITResponse, error) {
	response := new(ForceExpireHITResponse)
	err := c.call("ForceExpireHIT", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) ApproveAssignment(request *ApproveAssignmentRequest) (*ApproveAssignmentResponse, error) {
	response := new(ApproveAssignmentResponse)
	err := c.call("ApproveAssignment", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) RejectAssignment(request *RejectAssignmentRequest) (*RejectAssignmentResponse, error) {
	response := new(RejectAssignmentResponse)
	err := c.call("RejectAssignment", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) ApproveRejectedAssignment(request *ApproveRejectedAssignmentRequest) (*ApproveRejectedAssignmentResponse, error) {
	response := new(ApproveRejectedAssignmentResponse)
	err := c.call("ApproveRejectedAssignment", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) GetAssignmentsForHIT(request *GetAssignmentsForHITRequest) (*GetAssignmentsForHITResponse, error) {
	response := new(GetAssignmentsForHITResponse)
	err := c.call("GetAssignmentsForHIT", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) GetFileUploadURL(request *GetFileUploadURLRequest) (*GetFileUploadURLResponse, error) {
	response := new(GetFileUploadURLResponse)
	err := c.call("GetFileUploadURL", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) SearchHITs(request *SearchHITsRequest) (*SearchHITsResponse, error) {
	response := new(SearchHITsResponse)
	err := c.call("SearchHITs", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) GrantBonus(request *GrantBonusRequest) (*GrantBonusResponse, error) {
	response := new(GrantBonusResponse)
	err := c.call("GrantBonus", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) GetBonusPayments(request *GetBonusPaymentsRequest) (*GetBonusPaymentsResponse, error) {
	response := new(GetBonusPaymentsResponse)
	err := c.call("GetBonusPayments", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) ChangeHITTypeOfHIT(request *ChangeHITTypeOfHITRequest) (*ChangeHITTypeOfHITResponse, error) {
	response := new(ChangeHITTypeOfHITResponse)
	err := c.call("ChangeHITTypeOfHIT", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) CreateQualificationType(request *CreateQualificationTypeRequest) (*CreateQualificationTypeResponse, error) {
	response := new(CreateQualificationTypeResponse)
	err := c.call("CreateQualificationType", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) GrantQualification(request *GrantQualificationRequest) (*GrantQualificationResponse, error) {
	response := new(GrantQualificationResponse)
	err := c.call("GrantQualification", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) AssignQualification(request *AssignQualificationRequest) (*AssignQualificationResponse, error) {
	response := new(AssignQualificationResponse)
	err := c.call("AssignQualification", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) RevokeQualification(request *RevokeQualificationRequest) (*RevokeQualificationResponse, error) {
	response := new(RevokeQualificationResponse)
	err := c.call("RevokeQualification", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) GetQualificationType(request *GetQualificationTypeRequest) (*GetQualificationTypeResponse, error) {
	response := new(GetQualificationTypeResponse)
	err := c.call("GetQualificationType", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) GetQualificationRequests(request *GetQualificationRequestsRequest) (*GetQualificationRequestsResponse, error) {
	response := new(GetQualificationRequestsResponse)
	err := c.call("GetQualificationRequests", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) RejectQualificationRequest(request *RejectQualificationRequestRequest) (*RejectQualificationRequestResponse, error) {
	response := new(RejectQualificationRequestResponse)
	err := c.call("RejectQualificationRequest", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) UpdateQualificationType(request *UpdateQualificationTypeRequest) (*UpdateQualificationTypeResponse, error) {
	response := new(UpdateQualificationTypeResponse)
	err := c.call("UpdateQualificationType", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) SearchQualificationTypes(request *SearchQualificationTypesRequest) (*SearchQualificationTypesResponse, error) {
	response := new(SearchQualificationTypesResponse)
	err := c.call("SearchQualificationTypes", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) GetQualificationScore(request *GetQualificationScoreRequest) (*GetQualificationScoreResponse, error) {
	response := new(GetQualificationScoreResponse)
	err := c.call("GetQualificationScore", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) UpdateQualificationScore(request *UpdateQualificationScoreRequest) (*UpdateQualificationScoreResponse, error) {
	response := new(UpdateQualificationScoreResponse)
	err := c.call("UpdateQualificationScore", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) DisposeQualificationType(request *DisposeQualificationTypeRequest) (*DisposeQualificationTypeResponse, error) {
	response := new(DisposeQualificationTypeResponse)
	err := c.call("DisposeQualificationType", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) GetRequesterStatistic(request *GetRequesterStatisticRequest) (*GetRequesterStatisticResponse, error) {
	response := new(GetRequesterStatisticResponse)
	err := c.call("GetRequesterStatistic", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) GetRequesterWorkerStatistic(request *GetRequesterWorkerStatisticRequest) (*GetRequesterWorkerStatisticResponse, error) {
	response := new(GetRequesterWorkerStatisticResponse)
	err := c.call("GetRequesterWorkerStatistic", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) NotifyWorkers(request *NotifyWorkersRequest) (*NotifyWorkersResponse, error) {
	response := new(NotifyWorkersResponse)
	err := c.call("NotifyWorkers", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) GetBlockedWorkers(request *GetBlockedWorkersRequest) (*GetBlockedWorkersResponse, error) {
	response := new(GetBlockedWorkersResponse)
	err := c.call("GetBlockedWorkers", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) BlockWorker(request *BlockWorkerRequest) (*BlockWorkerResponse, error) {
	response := new(BlockWorkerResponse)
	err := c.call("BlockWorker", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) UnblockWorker(request *UnblockWorkerRequest) (*UnblockWorkerResponse, error) {
	response := new(UnblockWorkerResponse)
	err := c.call("UnblockWorker", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) Help(request *HelpRequest) (*HelpResponse, error) {
	response := new(HelpResponse)
	err := c.call("Help", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
