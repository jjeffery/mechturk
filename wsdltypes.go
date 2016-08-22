package mechturk

// NOTE: the types in this file were initially generated
// using gowsdl (github.com/hooklift/gowsdl). The file has,
// however, been extensively hand-edited.
//
// * Operations removed
// * Namespaces removed
// * Comments added for Godoc
// * Arrays changed to scalars where only one item is possible
//   for example, GetAccountBalanceResponse contains only one
//   result where the auto-generated WSDL allows for multiples.
//
// Ideally it would be a good idea to fork gowsdl and create
// a generator that produces code, but this is unlikely to
// happen.

import (
	"time"
)

type SortDirection string

// Sort directions
const (
	SortDirectionAscending  SortDirection = "Ascending"
	SortDirectionDescending SortDirection = "Descending"
)

type Comparator string

const (
	ComparatorLessThan             Comparator = "LessThan"
	ComparatorLessThanOrEqualTo    Comparator = "LessThanOrEqualTo"
	ComparatorGreaterThan          Comparator = "GreaterThan"
	ComparatorGreaterThanOrEqualTo Comparator = "GreaterThanOrEqualTo"
	ComparatorEqualTo              Comparator = "EqualTo"
	ComparatorNotEqualTo           Comparator = "NotEqualTo"
	ComparatorExists               Comparator = "Exists"
)

type EventType string

const (
	EventTypeAssignmentAccepted  EventType = "AssignmentAccepted"
	EventTypeAssignmentAbandoned EventType = "AssignmentAbandoned"
	EventTypeAssignmentReturned  EventType = "AssignmentReturned"
	EventTypeAssignmentSubmitted EventType = "AssignmentSubmitted"
	EventTypeHITReviewable       EventType = "HITReviewable"
	EventTypeHITExpired          EventType = "HITExpired"
	EventTypePing                EventType = "Ping"
)

type NotificationTransport string

const (
	NotificationTransportSOAP  NotificationTransport = "SOAP"
	NotificationTransportREST  NotificationTransport = "REST"
	NotificationTransportEmail NotificationTransport = "Email"
	NotificationTransportSQS   NotificationTransport = "SQS"
)

type HITStatus string

const (
	HITStatusAssignable   HITStatus = "Assignable"
	HITStatusUnassignable HITStatus = "Unassignable"
	HITStatusReviewable   HITStatus = "Reviewable"
	HITStatusReviewing    HITStatus = "Reviewing"
	HITStatusDisposed     HITStatus = "Disposed"
)

type HITReviewStatus string

const (
	HITReviewStatusNotReviewed           HITReviewStatus = "NotReviewed"
	HITReviewStatusMarkedForReview       HITReviewStatus = "MarkedForReview"
	HITReviewStatusReviewedAppropriate   HITReviewStatus = "ReviewedAppropriate"
	HITReviewStatusReviewedInappropriate HITReviewStatus = "ReviewedInappropriate"
)

type ReviewableHITStatus string

const (
	ReviewableHITStatusReviewable ReviewableHITStatus = "Reviewable"
	ReviewableHITStatusReviewing  ReviewableHITStatus = "Reviewing"
)

type GetReviewableHITsSortProperty string

const (
	GetReviewableHITsSortPropertyTitle        GetReviewableHITsSortProperty = "Title"
	GetReviewableHITsSortPropertyReward       GetReviewableHITsSortProperty = "Reward"
	GetReviewableHITsSortPropertyExpiration   GetReviewableHITsSortProperty = "Expiration"
	GetReviewableHITsSortPropertyCreationTime GetReviewableHITsSortProperty = "CreationTime"
	GetReviewableHITsSortPropertyEnumeration  GetReviewableHITsSortProperty = "Enumeration"
)

type ReviewPolicyLevel string

const (
	ReviewPolicyLevelAssignment ReviewPolicyLevel = "Assignment"
	ReviewPolicyLevelHIT        ReviewPolicyLevel = "HIT"
)

type ReviewActionStatus string

const (
	ReviewActionStatusIntended  ReviewActionStatus = "Intended"
	ReviewActionStatusSucceeded ReviewActionStatus = "Succeeded"
	ReviewActionStatusFailed    ReviewActionStatus = "Failed"
	ReviewActionStatusCancelled ReviewActionStatus = "Cancelled"
)

type AssignmentStatus string

const (
	AssignmentStatusSubmitted AssignmentStatus = "Submitted"
	AssignmentStatusApproved  AssignmentStatus = "Approved"
	AssignmentStatusRejected  AssignmentStatus = "Rejected"
)

type GetAssignmentsForHITSortProperty string

const (
	GetAssignmentsForHITSortPropertyAcceptTime       GetAssignmentsForHITSortProperty = "AcceptTime"
	GetAssignmentsForHITSortPropertySubmitTime       GetAssignmentsForHITSortProperty = "SubmitTime"
	GetAssignmentsForHITSortPropertyAnswer           GetAssignmentsForHITSortProperty = "Answer"
	GetAssignmentsForHITSortPropertyAssignmentStatus GetAssignmentsForHITSortProperty = "AssignmentStatus"
)

type SearchHITsSortProperty string

const (
	SearchHITsSortPropertyTitle        SearchHITsSortProperty = "Title"
	SearchHITsSortPropertyReward       SearchHITsSortProperty = "Reward"
	SearchHITsSortPropertyExpiration   SearchHITsSortProperty = "Expiration"
	SearchHITsSortPropertyCreationTime SearchHITsSortProperty = "CreationTime"
	SearchHITsSortPropertyEnumeration  SearchHITsSortProperty = "Enumeration"
)

type QualificationTypeStatus string

const (
	QualificationTypeStatusActive   QualificationTypeStatus = "Active"
	QualificationTypeStatusInactive QualificationTypeStatus = "Inactive"
)

type GetQualificationRequestsSortProperty string

const (
	GetQualificationRequestsSortPropertyQualificationTypeId GetQualificationRequestsSortProperty = "QualificationTypeId"
	GetQualificationRequestsSortPropertySubmitTime          GetQualificationRequestsSortProperty = "SubmitTime"
)

type QualificationStatus string

const (
	QualificationStatusGranted QualificationStatus = "Granted"
	QualificationStatusRevoked QualificationStatus = "Revoked"
)

type SearchQualificationTypesSortProperty string

const (
	SearchQualificationTypesSortPropertyName SearchQualificationTypesSortProperty = "Name"
)

type RequesterStatistic string

const (
	RequesterStatisticNumberHITsAssignable            RequesterStatistic = "NumberHITsAssignable"
	RequesterStatisticNumberHITsReviewable            RequesterStatistic = "NumberHITsReviewable"
	RequesterStatisticNumberHITsCreated               RequesterStatistic = "NumberHITsCreated"
	RequesterStatisticNumberHITsCompleted             RequesterStatistic = "NumberHITsCompleted"
	RequesterStatisticTotalRewardPayout               RequesterStatistic = "TotalRewardPayout"
	RequesterStatisticTotalRewardFeePayout            RequesterStatistic = "TotalRewardFeePayout"
	RequesterStatisticTotalFeePayout                  RequesterStatistic = "TotalFeePayout"
	RequesterStatisticTotalRewardAndFeePayout         RequesterStatistic = "TotalRewardAndFeePayout"
	RequesterStatisticTotalBonusPayout                RequesterStatistic = "TotalBonusPayout"
	RequesterStatisticTotalBonusFeePayout             RequesterStatistic = "TotalBonusFeePayout"
	RequesterStatisticEstimatedFeeLiability           RequesterStatistic = "EstimatedFeeLiability"
	RequesterStatisticEstimatedRewardLiability        RequesterStatistic = "EstimatedRewardLiability"
	RequesterStatisticEstimatedTotalLiability         RequesterStatistic = "EstimatedTotalLiability"
	RequesterStatisticNumberAssignmentsAvailable      RequesterStatistic = "NumberAssignmentsAvailable"
	RequesterStatisticNumberAssignmentsAccepted       RequesterStatistic = "NumberAssignmentsAccepted"
	RequesterStatisticNumberAssignmentsPending        RequesterStatistic = "NumberAssignmentsPending"
	RequesterStatisticNumberAssignmentsApproved       RequesterStatistic = "NumberAssignmentsApproved"
	RequesterStatisticNumberAssignmentsRejected       RequesterStatistic = "NumberAssignmentsRejected"
	RequesterStatisticNumberAssignmentsReturned       RequesterStatistic = "NumberAssignmentsReturned"
	RequesterStatisticNumberAssignmentsAbandoned      RequesterStatistic = "NumberAssignmentsAbandoned"
	RequesterStatisticAverageRewardAmount             RequesterStatistic = "AverageRewardAmount"
	RequesterStatisticPercentAssignmentsApproved      RequesterStatistic = "PercentAssignmentsApproved"
	RequesterStatisticPercentAssignmentsRejected      RequesterStatistic = "PercentAssignmentsRejected"
	RequesterStatisticNumberKnownAnswersCorrect       RequesterStatistic = "NumberKnownAnswersCorrect"
	RequesterStatisticNumberKnownAnswersIncorrect     RequesterStatistic = "NumberKnownAnswersIncorrect"
	RequesterStatisticNumberKnownAnswersEvaluated     RequesterStatistic = "NumberKnownAnswersEvaluated"
	RequesterStatisticPercentKnownAnswersCorrect      RequesterStatistic = "PercentKnownAnswersCorrect"
	RequesterStatisticNumberPluralityAnswersCorrect   RequesterStatistic = "NumberPluralityAnswersCorrect"
	RequesterStatisticNumberPluralityAnswersIncorrect RequesterStatistic = "NumberPluralityAnswersIncorrect"
	RequesterStatisticNumberPluralityAnswersEvaluated RequesterStatistic = "NumberPluralityAnswersEvaluated"
	RequesterStatisticPercentPluralityAnswersCorrect  RequesterStatistic = "PercentPluralityAnswersCorrect"
)

type NotifyWorkersFailureCode string

const (
	NotifyWorkersFailureCodeSoftFailure NotifyWorkersFailureCode = "SoftFailure"
	NotifyWorkersFailureCodeHardFailure NotifyWorkersFailureCode = "HardFailure"
)

type TimePeriod string

const (
	TimePeriodOneDay     TimePeriod = "OneDay"
	TimePeriodSevenDays  TimePeriod = "SevenDays"
	TimePeriodThirtyDays TimePeriod = "ThirtyDays"
	TimePeriodLifeToDate TimePeriod = "LifeToDate"
)

// CreateHITResponse is the output from CreateHIT, which is documented at
// http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_CreateHITOperation.html.
type createHITResponse struct {
	OperationRequest *OperationRequest `xml:"OperationRequest,omitempty"`
	HIT              *HIT              `xml:"HIT,omitempty"`
}

type registerHITTypeResponse struct {
	OperationRequest      *OperationRequest      `xml:"OperationRequest,omitempty"`
	RegisterHITTypeResult *RegisterHITTypeResult `xml:"RegisterHITTypeResult,omitempty"`
}

type setHITTypeNotificationResponse struct {
	OperationRequest             *OperationRequest             `xml:"OperationRequest,omitempty"`
	SetHITTypeNotificationResult *SetHITTypeNotificationResult `xml:"SetHITTypeNotificationResult,omitempty"`
}

type sendTestEventNotificationResponse struct {
	OperationRequest                *OperationRequest                `xml:"OperationRequest,omitempty"`
	SendTestEventNotificationResult *SendTestEventNotificationResult `xml:"SendTestEventNotificationResult,omitempty"`
}

type disposeHITResponse struct {
	OperationRequest *OperationRequest `xml:"OperationRequest,omitempty"`
	DisposeHITResult *DisposeHITResult `xml:"DisposeHITResult,omitempty"`
}

type disableHITResponse struct {
	OperationRequest *OperationRequest `xml:"OperationRequest,omitempty"`
	DisableHITResult *DisableHITResult `xml:"DisableHITResult,omitempty"`
}

type getHITResponse struct {
	OperationRequest *OperationRequest `xml:"OperationRequest,omitempty"`
	HIT              *HIT              `xml:"HIT,omitempty"`
}

type getAssignmentResponse struct {
	OperationRequest    *OperationRequest    `xml:"OperationRequest,omitempty"`
	GetAssignmentResult *GetAssignmentResult `xml:"GetAssignmentResult,omitempty"`
}

type getReviewableHITsResponse struct {
	OperationRequest        *OperationRequest        `xml:"OperationRequest,omitempty"`
	GetReviewableHITsResult *GetReviewableHITsResult `xml:"GetReviewableHITsResult,omitempty"`
}

type getReviewResultsForHITResponse struct {
	OperationRequest             *OperationRequest             `xml:"OperationRequest,omitempty"`
	GetReviewResultsForHITResult *GetReviewResultsForHITResult `xml:"GetReviewResultsForHITResult,omitempty"`
}

type getHITsForQualificationTypeResponse struct {
	OperationRequest                  *OperationRequest                  `xml:"OperationRequest,omitempty"`
	GetHITsForQualificationTypeResult *GetHITsForQualificationTypeResult `xml:"GetHITsForQualificationTypeResult,omitempty"`
}

type getQualificationsForQualificationTypeResponse struct {
	OperationRequest                            *OperationRequest                            `xml:"OperationRequest,omitempty"`
	GetQualificationsForQualificationTypeResult *GetQualificationsForQualificationTypeResult `xml:"GetQualificationsForQualificationTypeResult,omitempty"`
}

type setHITAsReviewingResponse struct {
	OperationRequest        *OperationRequest        `xml:"OperationRequest,omitempty"`
	SetHITAsReviewingResult *SetHITAsReviewingResult `xml:"SetHITAsReviewingResult,omitempty"`
}

type extendHITResponse struct {
	OperationRequest *OperationRequest `xml:"OperationRequest,omitempty"`
	ExtendHITResult  *ExtendHITResult  `xml:"ExtendHITResult,omitempty"`
}

type forceExpireHITResponse struct {
	OperationRequest     *OperationRequest     `xml:"OperationRequest,omitempty"`
	ForceExpireHITResult *ForceExpireHITResult `xml:"ForceExpireHITResult,omitempty"`
}

type approveAssignmentResponse struct {
	OperationRequest        *OperationRequest        `xml:"OperationRequest,omitempty"`
	ApproveAssignmentResult *ApproveAssignmentResult `xml:"ApproveAssignmentResult,omitempty"`
}

type rejectAssignmentResponse struct {
	OperationRequest       *OperationRequest       `xml:"OperationRequest,omitempty"`
	RejectAssignmentResult *RejectAssignmentResult `xml:"RejectAssignmentResult,omitempty"`
}

type approveRejectedAssignmentResponse struct {
	OperationRequest                *OperationRequest                `xml:"OperationRequest,omitempty"`
	ApproveRejectedAssignmentResult *ApproveRejectedAssignmentResult `xml:"ApproveRejectedAssignmentResult,omitempty"`
}

type getAssignmentsForHITResponse struct {
	OperationRequest           *OperationRequest           `xml:"OperationRequest,omitempty"`
	GetAssignmentsForHITResult *GetAssignmentsForHITResult `xml:"GetAssignmentsForHITResult,omitempty"`
}

type getFileUploadURLResponse struct {
	OperationRequest       *OperationRequest       `xml:"OperationRequest,omitempty"`
	GetFileUploadURLResult *GetFileUploadURLResult `xml:"GetFileUploadURLResult,omitempty"`
}

type searchHITsResponse struct {
	OperationRequest *OperationRequest `xml:"OperationRequest,omitempty"`
	SearchHITsResult *SearchHITsResult `xml:"SearchHITsResult,omitempty"`
}

type grantBonusResponse struct {
	OperationRequest *OperationRequest `xml:"OperationRequest,omitempty"`
	GrantBonusResult *GrantBonusResult `xml:"GrantBonusResult,omitempty"`
}

type getBonusPaymentsResponse struct {
	OperationRequest       *OperationRequest       `xml:"OperationRequest,omitempty"`
	GetBonusPaymentsResult *GetBonusPaymentsResult `xml:"GetBonusPaymentsResult,omitempty"`
}

type changeHITTypeOfHITResponse struct {
	OperationRequest         *OperationRequest         `xml:"OperationRequest,omitempty"`
	ChangeHITTypeOfHITResult *ChangeHITTypeOfHITResult `xml:"ChangeHITTypeOfHITResult,omitempty"`
}

type createQualificationTypeResponse struct {
	OperationRequest  *OperationRequest  `xml:"OperationRequest,omitempty"`
	QualificationType *QualificationType `xml:"QualificationType,omitempty"`
}

type disposeQualificationTypeResponse struct {
	OperationRequest               *OperationRequest               `xml:"OperationRequest,omitempty"`
	DisposeQualificationTypeResult *DisposeQualificationTypeResult `xml:"DisposeQualificationTypeResult,omitempty"`
}

type getQualificationRequestsResponse struct {
	OperationRequest               *OperationRequest               `xml:"OperationRequest,omitempty"`
	GetQualificationRequestsResult *GetQualificationRequestsResult `xml:"GetQualificationRequestsResult,omitempty"`
}

type rejectQualificationRequestResponse struct {
	OperationRequest                 *OperationRequest                 `xml:"OperationRequest,omitempty"`
	RejectQualificationRequestResult *RejectQualificationRequestResult `xml:"RejectQualificationRequestResult,omitempty"`
}

type grantQualificationResponse struct {
	OperationRequest         *OperationRequest         `xml:"OperationRequest,omitempty"`
	GrantQualificationResult *GrantQualificationResult `xml:"GrantQualificationResult,omitempty"`
}

type assignQualificationResponse struct {
	OperationRequest          *OperationRequest          `xml:"OperationRequest,omitempty"`
	AssignQualificationResult *AssignQualificationResult `xml:"AssignQualificationResult,omitempty"`
}

type revokeQualificationResponse struct {
	OperationRequest          *OperationRequest          `xml:"OperationRequest,omitempty"`
	RevokeQualificationResult *RevokeQualificationResult `xml:"RevokeQualificationResult,omitempty"`
}

type updateQualificationScoreResponse struct {
	OperationRequest               *OperationRequest               `xml:"OperationRequest,omitempty"`
	UpdateQualificationScoreResult *UpdateQualificationScoreResult `xml:"UpdateQualificationScoreResult,omitempty"`
}

type getQualificationTypeResponse struct {
	OperationRequest  *OperationRequest  `xml:"OperationRequest,omitempty"`
	QualificationType *QualificationType `xml:"QualificationType,omitempty"`
}

type getQualificationScoreResponse struct {
	OperationRequest *OperationRequest `xml:"OperationRequest,omitempty"`
	Qualification    *Qualification    `xml:"Qualification,omitempty"`
}

type searchQualificationTypesResponse struct {
	OperationRequest               *OperationRequest               `xml:"OperationRequest,omitempty"`
	SearchQualificationTypesResult *SearchQualificationTypesResult `xml:"SearchQualificationTypesResult,omitempty"`
}

type updateQualificationTypeResponse struct {
	OperationRequest  *OperationRequest  `xml:"OperationRequest,omitempty"`
	QualificationType *QualificationType `xml:"QualificationType,omitempty"`
}

type getAccountBalanceResponse struct {
	OperationRequest        *OperationRequest        `xml:"OperationRequest,omitempty"`
	GetAccountBalanceResult *GetAccountBalanceResult `xml:"GetAccountBalanceResult,omitempty"`
}

type getRequesterStatisticResponse struct {
	OperationRequest   *OperationRequest   `xml:"OperationRequest,omitempty"`
	GetStatisticResult *GetStatisticResult `xml:"GetStatisticResult,omitempty"`
}

type getRequesterWorkerStatisticResponse struct {
	OperationRequest   *OperationRequest   `xml:"OperationRequest,omitempty"`
	GetStatisticResult *GetStatisticResult `xml:"GetStatisticResult,omitempty"`
}

type notifyWorkersResponse struct {
	OperationRequest    *OperationRequest    `xml:"OperationRequest,omitempty"`
	NotifyWorkersResult *NotifyWorkersResult `xml:"NotifyWorkersResult,omitempty"`
}

type getBlockedWorkersResponse struct {
	OperationRequest        *OperationRequest        `xml:"OperationRequest,omitempty"`
	GetBlockedWorkersResult *GetBlockedWorkersResult `xml:"GetBlockedWorkersResult,omitempty"`
}

type blockWorkerResponse struct {
	OperationRequest  *OperationRequest  `xml:"OperationRequest,omitempty"`
	BlockWorkerResult *BlockWorkerResult `xml:"BlockWorkerResult,omitempty"`
}

type unblockWorkerResponse struct {
	OperationRequest    *OperationRequest    `xml:"OperationRequest,omitempty"`
	UnblockWorkerResult *UnblockWorkerResult `xml:"UnblockWorkerResult,omitempty"`
}

type OperationRequest struct {
	HTTPHeaders *HTTPHeaders `xml:"HTTPHeaders,omitempty"`
	RequestId   string       `xml:"RequestId,omitempty"`
	Arguments   *Arguments   `xml:"Arguments,omitempty"`
	Errors      *Errors      `xml:"Errors,omitempty"`
}

type Request struct {
	IsValid                                      string                                        `xml:"IsValid,omitempty"`
	CreateHITRequest                             *CreateHITRequest                             `xml:"CreateHITRequest,omitempty"`
	RegisterHITTypeRequest                       *RegisterHITTypeRequest                       `xml:"RegisterHITTypeRequest,omitempty"`
	DisposeHITRequest                            *DisposeHITRequest                            `xml:"DisposeHITRequest,omitempty"`
	DisableHITRequest                            *DisableHITRequest                            `xml:"DisableHITRequest,omitempty"`
	GetHITRequest                                *GetHITRequest                                `xml:"GetHITRequest,omitempty"`
	GetAssignmentRequest                         *GetAssignmentRequest                         `xml:"GetAssignmentRequest,omitempty"`
	GetReviewResultsForHITRequest                *GetReviewResultsForHITRequest                `xml:"GetReviewResultsForHITRequest,omitempty"`
	GetReviewableHITsRequest                     *GetReviewableHITsRequest                     `xml:"GetReviewableHITsRequest,omitempty"`
	GetHITsForQualificationTypeRequest           *GetHITsForQualificationTypeRequest           `xml:"GetHITsForQualificationTypeRequest,omitempty"`
	GetQualificationsForQualificationTypeRequest *GetQualificationsForQualificationTypeRequest `xml:"GetQualificationsForQualificationTypeRequest,omitempty"`
	SetHITAsReviewingRequest                     *SetHITAsReviewingRequest                     `xml:"SetHITAsReviewingRequest,omitempty"`
	SearchHITsRequest                            *SearchHITsRequest                            `xml:"SearchHITsRequest,omitempty"`
	ExtendHITRequest                             *ExtendHITRequest                             `xml:"ExtendHITRequest,omitempty"`
	ForceExpireHITRequest                        *ForceExpireHITRequest                        `xml:"ForceExpireHITRequest,omitempty"`
	ChangeHITTypeOfHITRequest                    *ChangeHITTypeOfHITRequest                    `xml:"ChangeHITTypeOfHITRequest,omitempty"`
	CreateQualificationTypeRequest               *CreateQualificationTypeRequest               `xml:"CreateQualificationTypeRequest,omitempty"`
	DisposeQualificationTypeRequest              *DisposeQualificationTypeRequest              `xml:"DisposeQualificationTypeRequest,omitempty"`
	GrantQualificationRequest                    *GrantQualificationRequest                    `xml:"GrantQualificationRequest,omitempty"`
	AssignQualificationRequest                   *AssignQualificationRequest                   `xml:"AssignQualificationRequest,omitempty"`
	RevokeQualificationRequest                   *RevokeQualificationRequest                   `xml:"RevokeQualificationRequest,omitempty"`
	GetQualificationRequestsRequest              *GetQualificationRequestsRequest              `xml:"GetQualificationRequestsRequest,omitempty"`
	RejectQualificationRequestRequest            *RejectQualificationRequestRequest            `xml:"RejectQualificationRequestRequest,omitempty"`
	GetQualificationTypeRequest                  *GetQualificationTypeRequest                  `xml:"GetQualificationTypeRequest,omitempty"`
	SearchQualificationTypesRequest              *SearchQualificationTypesRequest              `xml:"SearchQualificationTypesRequest,omitempty"`
	UpdateQualificationTypeRequest               *UpdateQualificationTypeRequest               `xml:"UpdateQualificationTypeRequest,omitempty"`
	ApproveAssignmentRequest                     *ApproveAssignmentRequest                     `xml:"ApproveAssignmentRequest,omitempty"`
	RejectAssignmentRequest                      *RejectAssignmentRequest                      `xml:"RejectAssignmentRequest,omitempty"`
	ApproveRejectedAssignmentRequest             *ApproveRejectedAssignmentRequest             `xml:"ApproveRejectedAssignmentRequest,omitempty"`
	GetAssignmentsForHIT                         *GetAssignmentsForHITRequest                  `xml:"GetAssignmentsForHIT,omitempty"`
	GetFileUploadURL                             *GetFileUploadURLRequest                      `xml:"GetFileUploadURL,omitempty"`
	GrantBonusRequest                            *GrantBonusRequest                            `xml:"GrantBonusRequest,omitempty"`
	GetBonusPaymentsRequest                      *GetBonusPaymentsRequest                      `xml:"GetBonusPaymentsRequest,omitempty"`
	GetAccountBalanceRequest                     *GetAccountBalanceRequest                     `xml:"GetAccountBalanceRequest,omitempty"`
	NotifyWorkersRequest                         *NotifyWorkersRequest                         `xml:"NotifyWorkersRequest,omitempty"`
	GetBlockedWorkersRequest                     *GetBlockedWorkersRequest                     `xml:"GetBlockedWorkersRequest,omitempty"`
	BlockWorkerRequest                           *BlockWorkerRequest                           `xml:"BlockWorkerRequest,omitempty"`
	UnblockWorkerRequest                         *UnblockWorkerRequest                         `xml:"UnblockWorkerRequest,omitempty"`
	GetRequesterStatistic                        *GetRequesterStatisticRequest                 `xml:"GetRequesterStatistic,omitempty"`
	GetRequesterWorkerStatistic                  *GetRequesterWorkerStatisticRequest           `xml:"GetRequesterWorkerStatistic,omitempty"`
	HelpRequest                                  *HelpRequest                                  `xml:"HelpRequest,omitempty"`
	Errors                                       *Errors                                       `xml:"Errors,omitempty"`
}

type HTTPHeaders struct {
	Header struct {
		Name  string `xml:"Name,attr,omitempty"`
		Value string `xml:"Value,attr,omitempty"`
	} `xml:"Header,omitempty"`
}

type Arguments struct {
	Argument struct {
		Name  string `xml:"Name,attr,omitempty"`
		Value string `xml:"Value,attr,omitempty"`
	} `xml:"Argument,omitempty"`
}

type Errors struct {
	Error struct {
		Code    string          `xml:"Code,omitempty"`
		Message string          `xml:"Message,omitempty"`
		Data    []*KeyValuePair `xml:"Data,omitempty"`
	} `xml:"Error,omitempty"`
}

type helpResponse struct {
	OperationRequest *OperationRequest `xml:"OperationRequest,omitempty"`
	Information      *Information      `xml:"Information,omitempty"`
}

type Information struct {
	Request                  *Request                    `xml:"Request,omitempty"`
	OperationInformation     []*OperationInformation     `xml:"OperationInformation,omitempty"`
	ResponseGroupInformation []*ResponseGroupInformation `xml:"ResponseGroupInformation,omitempty"`
}

type OperationInformation struct {
	Name               string `xml:"Name,omitempty"`
	Description        string `xml:"Description,omitempty"`
	RequiredParameters struct {
		Parameter []string `xml:"Parameter,omitempty"`
	} `xml:"RequiredParameters,omitempty"`
	AvailableParameters struct {
		Parameter []string `xml:"Parameter,omitempty"`
	} `xml:"AvailableParameters,omitempty"`
	DefaultResponseGroups struct {
		ResponseGroup []string `xml:"ResponseGroup,omitempty"`
	} `xml:"DefaultResponseGroups,omitempty"`
	AvailableResponseGroups struct {
		ResponseGroup []string `xml:"ResponseGroup,omitempty"`
	} `xml:"AvailableResponseGroups,omitempty"`
}

type ResponseGroupInformation struct {
	Name            string `xml:"Name,omitempty"`
	CreationDate    string `xml:"CreationDate,omitempty"`
	ValidOperations struct {
		Operation []string `xml:"Operation,omitempty"`
	} `xml:"ValidOperations,omitempty"`
	Elements struct {
		Element []string `xml:"Element,omitempty"`
	} `xml:"Elements,omitempty"`
}

type Claim struct {
	Type string `xml:"Type,omitempty"`
	Body string `xml:"Body,omitempty"`
}

type Price struct {
	Amount         float64 `xml:"Amount,omitempty"`
	CurrencyCode   string  `xml:"CurrencyCode,omitempty"`
	FormattedPrice string  `xml:"FormattedPrice,omitempty"`
}

type QualificationRequirement struct {
	QualificationTypeId string      `xml:"QualificationTypeId,omitempty"`
	Comparator          *Comparator `xml:"Comparator,omitempty"`
	IntegerValue        int32       `xml:"IntegerValue,omitempty"`
	LocaleValue         *Locale     `xml:"LocaleValue,omitempty"`
	RequiredToPreview   bool        `xml:"RequiredToPreview,omitempty"`
}

type Locale struct {
	Country string `xml:"Country,omitempty"`
}

type NotificationSpecification struct {
	Destination string                 `xml:"Destination,omitempty"`
	Transport   *NotificationTransport `xml:"Transport,omitempty"`
	Version     string                 `xml:"Version,omitempty"`
	EventType   []*EventType           `xml:"EventType,omitempty"`
}

type HIT struct {
	Request                      *Request                    `xml:"Request,omitempty"`
	HITId                        string                      `xml:"HITId,omitempty"`
	HITTypeId                    string                      `xml:"HITTypeId,omitempty"`
	HITGroupId                   string                      `xml:"HITGroupId,omitempty"`
	HITLayoutId                  string                      `xml:"HITLayoutId,omitempty"`
	CreationTime                 time.Time                   `xml:"CreationTime,omitempty"`
	Title                        string                      `xml:"Title,omitempty"`
	Description                  string                      `xml:"Description,omitempty"`
	Question                     string                      `xml:"Question,omitempty"`
	Keywords                     string                      `xml:"Keywords,omitempty"`
	HITStatus                    *HITStatus                  `xml:"HITStatus,omitempty"`
	MaxAssignments               int32                       `xml:"MaxAssignments,omitempty"`
	Reward                       *Price                      `xml:"Reward,omitempty"`
	AutoApprovalDelayInSeconds   int64                       `xml:"AutoApprovalDelayInSeconds,omitempty"`
	Expiration                   time.Time                   `xml:"Expiration,omitempty"`
	AssignmentDurationInSeconds  int64                       `xml:"AssignmentDurationInSeconds,omitempty"`
	RequesterAnnotation          string                      `xml:"RequesterAnnotation,omitempty"`
	QualificationRequirement     []*QualificationRequirement `xml:"QualificationRequirement,omitempty"`
	HITReviewStatus              *HITReviewStatus            `xml:"HITReviewStatus,omitempty"`
	NumberOfAssignmentsPending   int32                       `xml:"NumberOfAssignmentsPending,omitempty"`
	NumberOfAssignmentsAvailable int32                       `xml:"NumberOfAssignmentsAvailable,omitempty"`
	NumberOfAssignmentsCompleted int32                       `xml:"NumberOfAssignmentsCompleted,omitempty"`
}

type ReviewPolicy struct {
	PolicyName string             `xml:"PolicyName,omitempty"`
	Parameter  []*PolicyParameter `xml:"Parameter,omitempty"`
}

type PolicyParameter struct {
	Key      string               `xml:"Key,omitempty"`
	Value    []string             `xml:"Value,omitempty"`
	MapEntry []*ParameterMapEntry `xml:"MapEntry,omitempty"`
}

type ParameterMapEntry struct {
	Key   string   `xml:"Key,omitempty"`
	Value []string `xml:"Value,omitempty"`
}

type HITLayoutParameter struct {
	Name  string `xml:"Name,omitempty"`
	Value string `xml:"Value,omitempty"`
}

// CreateHITRequest is the input to CreateHIT, which is documented at
// http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_CreateHITOperation.html.
type CreateHITRequest struct {
	HITTypeId                   string                      `xml:"HITTypeId,omitempty"`
	MaxAssignments              int32                       `xml:"MaxAssignments,omitempty"`
	AutoApprovalDelayInSeconds  int64                       `xml:"AutoApprovalDelayInSeconds,omitempty"`
	LifetimeInSeconds           int64                       `xml:"LifetimeInSeconds,omitempty"`
	AssignmentDurationInSeconds int64                       `xml:"AssignmentDurationInSeconds,omitempty"`
	Reward                      *Price                      `xml:"Reward,omitempty"`
	Title                       string                      `xml:"Title,omitempty"`
	Keywords                    string                      `xml:"Keywords,omitempty"`
	Description                 string                      `xml:"Description,omitempty"`
	Question                    string                      `xml:"Question,omitempty"`
	RequesterAnnotation         string                      `xml:"RequesterAnnotation,omitempty"`
	QualificationRequirement    []*QualificationRequirement `xml:"QualificationRequirement,omitempty"`
	UniqueRequestToken          string                      `xml:"UniqueRequestToken,omitempty"`
	AssignmentReviewPolicy      *ReviewPolicy               `xml:"AssignmentReviewPolicy,omitempty"`
	HITReviewPolicy             *ReviewPolicy               `xml:"HITReviewPolicy,omitempty"`
	HITLayoutId                 string                      `xml:"HITLayoutId,omitempty"`
	HITLayoutParameter          []*HITLayoutParameter       `xml:"HITLayoutParameter,omitempty"`
	ResponseGroup               []string                    `xml:"ResponseGroup,omitempty"`
}

type RegisterHITTypeRequest struct {
	AutoApprovalDelayInSeconds  int64                       `xml:"AutoApprovalDelayInSeconds,omitempty"`
	AssignmentDurationInSeconds int64                       `xml:"AssignmentDurationInSeconds,omitempty"`
	Reward                      *Price                      `xml:"Reward,omitempty"`
	Title                       string                      `xml:"Title,omitempty"`
	Keywords                    string                      `xml:"Keywords,omitempty"`
	Description                 string                      `xml:"Description,omitempty"`
	QualificationRequirement    []*QualificationRequirement `xml:"QualificationRequirement,omitempty"`
	ResponseGroup               []string                    `xml:"ResponseGroup,omitempty"`
}

type RegisterHITTypeResult struct {
	Request   *Request `xml:"Request,omitempty"`
	HITTypeId string   `xml:"HITTypeId,omitempty"`
}

type SetHITTypeNotificationRequest struct {
	HITTypeId    string                     `xml:"HITTypeId,omitempty"`
	Notification *NotificationSpecification `xml:"Notification,omitempty"`
	Active       bool                       `xml:"Active,omitempty"`
}

type SetHITTypeNotificationResult struct {
	Request *Request `xml:"Request,omitempty"`
}

type SendTestEventNotificationRequest struct {
	Notification  *NotificationSpecification `xml:"Notification,omitempty"`
	TestEventType *EventType                 `xml:"TestEventType,omitempty"`
}

type SendTestEventNotificationResult struct {
	Request *Request `xml:"Request,omitempty"`
}

type DisposeHITRequest struct {
	HITId         string   `xml:"HITId,omitempty"`
	ResponseGroup []string `xml:"ResponseGroup,omitempty"`
}

type DisposeHITResult struct {
	Request *Request `xml:"Request,omitempty"`
}

type DisableHITRequest struct {
	HITId         string   `xml:"HITId,omitempty"`
	ResponseGroup []string `xml:"ResponseGroup,omitempty"`
}

type DisableHITResult struct {
	Request *Request `xml:"Request,omitempty"`
}

type GetHITRequest struct {
	HITId         string   `xml:"HITId,omitempty"`
	ResponseGroup []string `xml:"ResponseGroup,omitempty"`
}

type GetAssignmentRequest struct {
	AssignmentId  string   `xml:"AssignmentId,omitempty"`
	ResponseGroup []string `xml:"ResponseGroup,omitempty"`
}

type GetAssignmentResult struct {
	Request    *Request    `xml:"Request,omitempty"`
	Assignment *Assignment `xml:"Assignment,omitempty"`
	HIT        *HIT        `xml:"HIT,omitempty"`
}

type GetReviewableHITsRequest struct {
	HITTypeId     string                         `xml:"HITTypeId,omitempty"`
	Status        *ReviewableHITStatus           `xml:"Status,omitempty"`
	SortDirection *SortDirection                 `xml:"SortDirection,omitempty"`
	SortProperty  *GetReviewableHITsSortProperty `xml:"SortProperty,omitempty"`
	PageNumber    int32                          `xml:"PageNumber,omitempty"`
	PageSize      int32                          `xml:"PageSize,omitempty"`
	ResponseGroup []string                       `xml:"ResponseGroup,omitempty"`
}

type GetReviewableHITsResult struct {
	Request         *Request `xml:"Request,omitempty"`
	PageNumber      int32    `xml:"PageNumber,omitempty"`
	NumResults      int32    `xml:"NumResults,omitempty"`
	TotalNumResults int32    `xml:"TotalNumResults,omitempty"`
	HIT             []*HIT   `xml:"HIT,omitempty"`
}

type ReviewReport struct {
	PageNumber      int32                 `xml:"PageNumber,omitempty"`
	NumResults      int32                 `xml:"NumResults,omitempty"`
	TotalNumResults int32                 `xml:"TotalNumResults,omitempty"`
	ReviewResult    []*ReviewResultDetail `xml:"ReviewResult,omitempty"`
	ReviewAction    []*ReviewActionDetail `xml:"ReviewAction,omitempty"`
}

type ReviewResultDetail struct {
	ActionId    string `xml:"ActionId,omitempty"`
	SubjectId   string `xml:"SubjectId,omitempty"`
	SubjectType string `xml:"SubjectType,omitempty"`
	QuestionId  string `xml:"QuestionId,omitempty"`
	Key         string `xml:"Key,omitempty"`
	Value       string `xml:"Value,omitempty"`
}

type ReviewActionDetail struct {
	ActionId     string              `xml:"ActionId,omitempty"`
	ActionName   string              `xml:"ActionName,omitempty"`
	ObjectId     string              `xml:"ObjectId,omitempty"`
	ObjectType   string              `xml:"ObjectType,omitempty"`
	Status       *ReviewActionStatus `xml:"Status,omitempty"`
	CompleteTime time.Time           `xml:"CompleteTime,omitempty"`
	Result       string              `xml:"Result,omitempty"`
	ErrorCode    string              `xml:"ErrorCode,omitempty"`
}

type GetReviewResultsForHITRequest struct {
	HITId           string               `xml:"HITId,omitempty"`
	PolicyLevel     []*ReviewPolicyLevel `xml:"PolicyLevel,omitempty"`
	RetrieveActions bool                 `xml:"RetrieveActions,omitempty"`
	RetrieveResults bool                 `xml:"RetrieveResults,omitempty"`
	PageNumber      int32                `xml:"PageNumber,omitempty"`
	PageSize        int32                `xml:"PageSize,omitempty"`
	ResponseGroup   []string             `xml:"ResponseGroup,omitempty"`
}

type GetReviewResultsForHITResult struct {
	Request                *Request      `xml:"Request,omitempty"`
	HITId                  string        `xml:"HITId,omitempty"`
	AssignmentReviewPolicy *ReviewPolicy `xml:"AssignmentReviewPolicy,omitempty"`
	HITReviewPolicy        *ReviewPolicy `xml:"HITReviewPolicy,omitempty"`
	AssignmentReviewReport *ReviewReport `xml:"AssignmentReviewReport,omitempty"`
	HITReviewReport        *ReviewReport `xml:"HITReviewReport,omitempty"`
}

type GetHITsForQualificationTypeRequest struct {
	QualificationTypeId string   `xml:"QualificationTypeId,omitempty"`
	PageNumber          int32    `xml:"PageNumber,omitempty"`
	PageSize            int32    `xml:"PageSize,omitempty"`
	ResponseGroup       []string `xml:"ResponseGroup,omitempty"`
}

type GetHITsForQualificationTypeResult struct {
	Request         *Request `xml:"Request,omitempty"`
	PageNumber      int32    `xml:"PageNumber,omitempty"`
	NumResults      int32    `xml:"NumResults,omitempty"`
	TotalNumResults int32    `xml:"TotalNumResults,omitempty"`
	HIT             []*HIT   `xml:"HIT,omitempty"`
}

type GetQualificationsForQualificationTypeRequest struct {
	QualificationTypeId string               `xml:"QualificationTypeId,omitempty"`
	Status              *QualificationStatus `xml:"Status,omitempty"`
	PageNumber          int32                `xml:"PageNumber,omitempty"`
	PageSize            int32                `xml:"PageSize,omitempty"`
	ResponseGroup       []string             `xml:"ResponseGroup,omitempty"`
}

type GetQualificationsForQualificationTypeResult struct {
	Request         *Request         `xml:"Request,omitempty"`
	PageNumber      int32            `xml:"PageNumber,omitempty"`
	NumResults      int32            `xml:"NumResults,omitempty"`
	TotalNumResults int32            `xml:"TotalNumResults,omitempty"`
	Qualification   []*Qualification `xml:"Qualification,omitempty"`
}

type SetHITAsReviewingRequest struct {
	HITId         string   `xml:"HITId,omitempty"`
	Revert        bool     `xml:"Revert,omitempty"`
	ResponseGroup []string `xml:"ResponseGroup,omitempty"`
}

type SetHITAsReviewingResult struct {
	Request *Request `xml:"Request,omitempty"`
}

type ExtendHITRequest struct {
	HITId                        string   `xml:"HITId,omitempty"`
	MaxAssignmentsIncrement      int32    `xml:"MaxAssignmentsIncrement,omitempty"`
	ExpirationIncrementInSeconds int64    `xml:"ExpirationIncrementInSeconds,omitempty"`
	UniqueRequestToken           string   `xml:"UniqueRequestToken,omitempty"`
	ResponseGroup                []string `xml:"ResponseGroup,omitempty"`
}

type ExtendHITResult struct {
	Request *Request `xml:"Request,omitempty"`
}

type ForceExpireHITRequest struct {
	HITId         string   `xml:"HITId,omitempty"`
	ResponseGroup []string `xml:"ResponseGroup,omitempty"`
}

type ForceExpireHITResult struct {
	Request *Request `xml:"Request,omitempty"`
}

type Assignment struct {
	Request           *Request          `xml:"Request,omitempty"`
	AssignmentId      string            `xml:"AssignmentId,omitempty"`
	WorkerId          string            `xml:"WorkerId,omitempty"`
	HITId             string            `xml:"HITId,omitempty"`
	AssignmentStatus  *AssignmentStatus `xml:"AssignmentStatus,omitempty"`
	AutoApprovalTime  time.Time         `xml:"AutoApprovalTime,omitempty"`
	AcceptTime        time.Time         `xml:"AcceptTime,omitempty"`
	SubmitTime        time.Time         `xml:"SubmitTime,omitempty"`
	ApprovalTime      time.Time         `xml:"ApprovalTime,omitempty"`
	RejectionTime     time.Time         `xml:"RejectionTime,omitempty"`
	Deadline          time.Time         `xml:"Deadline,omitempty"`
	Answer            string            `xml:"Answer,omitempty"`
	RequesterFeedback string            `xml:"RequesterFeedback,omitempty"`
}

type ApproveAssignmentRequest struct {
	AssignmentId      string   `xml:"AssignmentId,omitempty"`
	ResponseGroup     []string `xml:"ResponseGroup,omitempty"`
	RequesterFeedback string   `xml:"RequesterFeedback,omitempty"`
}

type ApproveAssignmentResult struct {
	Request *Request `xml:"Request,omitempty"`
}

type RejectAssignmentRequest struct {
	AssignmentId      string   `xml:"AssignmentId,omitempty"`
	ResponseGroup     []string `xml:"ResponseGroup,omitempty"`
	RequesterFeedback string   `xml:"RequesterFeedback,omitempty"`
}

type RejectAssignmentResult struct {
	Request *Request `xml:"Request,omitempty"`
}

type ApproveRejectedAssignmentRequest struct {
	AssignmentId      string   `xml:"AssignmentId,omitempty"`
	ResponseGroup     []string `xml:"ResponseGroup,omitempty"`
	RequesterFeedback string   `xml:"RequesterFeedback,omitempty"`
}

type ApproveRejectedAssignmentResult struct {
	Request *Request `xml:"Request,omitempty"`
}

type GetAssignmentsForHITRequest struct {
	HITId            string                            `xml:"HITId,omitempty"`
	SortDirection    *SortDirection                    `xml:"SortDirection,omitempty"`
	SortProperty     *GetAssignmentsForHITSortProperty `xml:"SortProperty,omitempty"`
	PageNumber       int32                             `xml:"PageNumber,omitempty"`
	PageSize         int32                             `xml:"PageSize,omitempty"`
	AssignmentStatus []*AssignmentStatus               `xml:"AssignmentStatus,omitempty"`
	ResponseGroup    []string                          `xml:"ResponseGroup,omitempty"`
}

type GetAssignmentsForHITResult struct {
	Request         *Request      `xml:"Request,omitempty"`
	PageNumber      int32         `xml:"PageNumber,omitempty"`
	NumResults      int32         `xml:"NumResults,omitempty"`
	TotalNumResults int32         `xml:"TotalNumResults,omitempty"`
	Assignment      []*Assignment `xml:"Assignment,omitempty"`
}

type GetFileUploadURLRequest struct {
	AssignmentId       string `xml:"AssignmentId,omitempty"`
	QuestionIdentifier string `xml:"QuestionIdentifier,omitempty"`
}

type GetFileUploadURLResult struct {
	Request       *Request `xml:"Request,omitempty"`
	FileUploadURL string   `xml:"FileUploadURL,omitempty"`
}

type SearchHITsRequest struct {
	SortDirection *SortDirection          `xml:"SortDirection,omitempty"`
	SortProperty  *SearchHITsSortProperty `xml:"SortProperty,omitempty"`
	PageNumber    int32                   `xml:"PageNumber,omitempty"`
	PageSize      int32                   `xml:"PageSize,omitempty"`
	ResponseGroup []string                `xml:"ResponseGroup,omitempty"`
}

type SearchHITsResult struct {
	Request         *Request `xml:"Request,omitempty"`
	PageNumber      int32    `xml:"PageNumber,omitempty"`
	NumResults      int32    `xml:"NumResults,omitempty"`
	TotalNumResults int32    `xml:"TotalNumResults,omitempty"`
	HIT             []*HIT   `xml:"HIT,omitempty"`
}

type GrantBonusRequest struct {
	WorkerId           string   `xml:"WorkerId,omitempty"`
	BonusAmount        *Price   `xml:"BonusAmount,omitempty"`
	AssignmentId       string   `xml:"AssignmentId,omitempty"`
	Reason             string   `xml:"Reason,omitempty"`
	UniqueRequestToken string   `xml:"UniqueRequestToken,omitempty"`
	ResponseGroup      []string `xml:"ResponseGroup,omitempty"`
}

type GrantBonusResult struct {
	Request *Request `xml:"Request,omitempty"`
}

type BonusPayment struct {
	Request      *Request  `xml:"Request,omitempty"`
	WorkerId     string    `xml:"WorkerId,omitempty"`
	BonusAmount  *Price    `xml:"BonusAmount,omitempty"`
	AssignmentId string    `xml:"AssignmentId,omitempty"`
	Reason       string    `xml:"Reason,omitempty"`
	GrantTime    time.Time `xml:"GrantTime,omitempty"`
}

type GetBonusPaymentsRequest struct {
	HITId         string   `xml:"HITId,omitempty"`
	AssignmentId  string   `xml:"AssignmentId,omitempty"`
	PageNumber    int32    `xml:"PageNumber,omitempty"`
	PageSize      int32    `xml:"PageSize,omitempty"`
	ResponseGroup []string `xml:"ResponseGroup,omitempty"`
}

type GetBonusPaymentsResult struct {
	Request         *Request        `xml:"Request,omitempty"`
	NumResults      int32           `xml:"NumResults,omitempty"`
	PageNumber      int32           `xml:"PageNumber,omitempty"`
	TotalNumResults int32           `xml:"TotalNumResults,omitempty"`
	BonusPayment    []*BonusPayment `xml:"BonusPayment,omitempty"`
}

type ChangeHITTypeOfHITRequest struct {
	HITId         string   `xml:"HITId,omitempty"`
	HITTypeId     string   `xml:"HITTypeId,omitempty"`
	ResponseGroup []string `xml:"ResponseGroup,omitempty"`
}

type ChangeHITTypeOfHITResult struct {
	Request *Request `xml:"Request,omitempty"`
}

type QualificationType struct {
	Request                 *Request                 `xml:"Request,omitempty"`
	QualificationTypeId     string                   `xml:"QualificationTypeId,omitempty"`
	CreationTime            time.Time                `xml:"CreationTime,omitempty"`
	Name                    string                   `xml:"Name,omitempty"`
	Description             string                   `xml:"Description,omitempty"`
	Keywords                string                   `xml:"Keywords,omitempty"`
	QualificationTypeStatus *QualificationTypeStatus `xml:"QualificationTypeStatus,omitempty"`
	Test                    string                   `xml:"Test,omitempty"`
	TestDurationInSeconds   int64                    `xml:"TestDurationInSeconds,omitempty"`
	AnswerKey               string                   `xml:"AnswerKey,omitempty"`
	RetryDelayInSeconds     int64                    `xml:"RetryDelayInSeconds,omitempty"`
	IsRequestable           bool                     `xml:"IsRequestable,omitempty"`
	AutoGranted             bool                     `xml:"AutoGranted,omitempty"`
	AutoGrantedValue        int32                    `xml:"AutoGrantedValue,omitempty"`
}

type QualificationRequest struct {
	QualificationRequestId string    `xml:"QualificationRequestId,omitempty"`
	QualificationTypeId    string    `xml:"QualificationTypeId,omitempty"`
	SubjectId              string    `xml:"SubjectId,omitempty"`
	Test                   string    `xml:"Test,omitempty"`
	Answer                 string    `xml:"Answer,omitempty"`
	SubmitTime             time.Time `xml:"SubmitTime,omitempty"`
}

type CreateQualificationTypeRequest struct {
	Name                    string                   `xml:"Name,omitempty"`
	Keywords                string                   `xml:"Keywords,omitempty"`
	Description             string                   `xml:"Description,omitempty"`
	QualificationTypeStatus *QualificationTypeStatus `xml:"QualificationTypeStatus,omitempty"`
	RetryDelayInSeconds     int64                    `xml:"RetryDelayInSeconds,omitempty"`
	Test                    string                   `xml:"Test,omitempty"`
	AnswerKey               string                   `xml:"AnswerKey,omitempty"`
	TestDurationInSeconds   int64                    `xml:"TestDurationInSeconds,omitempty"`
	AutoGranted             bool                     `xml:"AutoGranted,omitempty"`
	AutoGrantedValue        int32                    `xml:"AutoGrantedValue,omitempty"`
	ResponseGroup           []string                 `xml:"ResponseGroup,omitempty"`
}

type DisposeQualificationTypeRequest struct {
	QualificationTypeId string `xml:"QualificationTypeId,omitempty"`
}

type DisposeQualificationTypeResult struct {
	Request *Request `xml:"Request,omitempty"`
}

type GetQualificationRequestsRequest struct {
	QualificationTypeId string                                `xml:"QualificationTypeId,omitempty"`
	SortDirection       *SortDirection                        `xml:"SortDirection,omitempty"`
	SortProperty        *GetQualificationRequestsSortProperty `xml:"SortProperty,omitempty"`
	PageNumber          int32                                 `xml:"PageNumber,omitempty"`
	PageSize            int32                                 `xml:"PageSize,omitempty"`
	ResponseGroup       []string                              `xml:"ResponseGroup,omitempty"`
}

type GetQualificationRequestsResult struct {
	Request              *Request                `xml:"Request,omitempty"`
	NumResults           int32                   `xml:"NumResults,omitempty"`
	TotalNumResults      int32                   `xml:"TotalNumResults,omitempty"`
	PageNumber           int32                   `xml:"PageNumber,omitempty"`
	QualificationRequest []*QualificationRequest `xml:"QualificationRequest,omitempty"`
}

type RejectQualificationRequestRequest struct {
	QualificationRequestId string   `xml:"QualificationRequestId,omitempty"`
	Reason                 string   `xml:"Reason,omitempty"`
	ResponseGroup          []string `xml:"ResponseGroup,omitempty"`
}

type RejectQualificationRequestResult struct {
	Request *Request `xml:"Request,omitempty"`
}

type GrantQualificationRequest struct {
	QualificationRequestId string   `xml:"QualificationRequestId,omitempty"`
	IntegerValue           int32    `xml:"IntegerValue,omitempty"`
	ResponseGroup          []string `xml:"ResponseGroup,omitempty"`
}

type GrantQualificationResult struct {
	Request *Request `xml:"Request,omitempty"`
}

type AssignQualificationRequest struct {
	QualificationTypeId string   `xml:"QualificationTypeId,omitempty"`
	WorkerId            string   `xml:"WorkerId,omitempty"`
	IntegerValue        int32    `xml:"IntegerValue,omitempty"`
	SendNotification    bool     `xml:"SendNotification,omitempty"`
	ResponseGroup       []string `xml:"ResponseGroup,omitempty"`
}

type AssignQualificationResult struct {
	Request *Request `xml:"Request,omitempty"`
}

type RevokeQualificationRequest struct {
	SubjectId           string   `xml:"SubjectId,omitempty"`
	QualificationTypeId string   `xml:"QualificationTypeId,omitempty"`
	Reason              string   `xml:"Reason,omitempty"`
	ResponseGroup       []string `xml:"ResponseGroup,omitempty"`
}

type RevokeQualificationResult struct {
	Request *Request `xml:"Request,omitempty"`
}

type UpdateQualificationScoreRequest struct {
	QualificationTypeId string   `xml:"QualificationTypeId,omitempty"`
	SubjectId           string   `xml:"SubjectId,omitempty"`
	IntegerValue        int32    `xml:"IntegerValue,omitempty"`
	ResponseGroup       []string `xml:"ResponseGroup,omitempty"`
}

type UpdateQualificationScoreResult struct {
	Request *Request `xml:"Request,omitempty"`
}

type GetQualificationTypeRequest struct {
	QualificationTypeId string   `xml:"QualificationTypeId,omitempty"`
	ResponseGroup       []string `xml:"ResponseGroup,omitempty"`
}

type GetQualificationScoreRequest struct {
	QualificationTypeId string   `xml:"QualificationTypeId,omitempty"`
	SubjectId           string   `xml:"SubjectId,omitempty"`
	ResponseGroup       []string `xml:"ResponseGroup,omitempty"`
}

type Qualification struct {
	Request             *Request             `xml:"Request,omitempty"`
	QualificationTypeId string               `xml:"QualificationTypeId,omitempty"`
	SubjectId           string               `xml:"SubjectId,omitempty"`
	GrantTime           time.Time            `xml:"GrantTime,omitempty"`
	IntegerValue        int32                `xml:"IntegerValue,omitempty"`
	LocaleValue         *Locale              `xml:"LocaleValue,omitempty"`
	Status              *QualificationStatus `xml:"Status,omitempty"`
}

type SearchQualificationTypesRequest struct {
	Query               string                                `xml:"Query,omitempty"`
	MustBeRequestable   bool                                  `xml:"MustBeRequestable,omitempty"`
	MustBeOwnedByCaller bool                                  `xml:"MustBeOwnedByCaller,omitempty"`
	SortDirection       *SortDirection                        `xml:"SortDirection,omitempty"`
	SortProperty        *SearchQualificationTypesSortProperty `xml:"SortProperty,omitempty"`
	PageNumber          int32                                 `xml:"PageNumber,omitempty"`
	PageSize            int32                                 `xml:"PageSize,omitempty"`
	ResponseGroup       []string                              `xml:"ResponseGroup,omitempty"`
}

type SearchQualificationTypesResult struct {
	Request           *Request             `xml:"Request,omitempty"`
	NumResults        int32                `xml:"NumResults,omitempty"`
	TotalNumResults   int32                `xml:"TotalNumResults,omitempty"`
	PageNumber        int32                `xml:"PageNumber,omitempty"`
	QualificationType []*QualificationType `xml:"QualificationType,omitempty"`
}

type UpdateQualificationTypeRequest struct {
	QualificationTypeId     string                   `xml:"QualificationTypeId,omitempty"`
	Description             string                   `xml:"Description,omitempty"`
	QualificationTypeStatus *QualificationTypeStatus `xml:"QualificationTypeStatus,omitempty"`
	Test                    string                   `xml:"Test,omitempty"`
	AnswerKey               string                   `xml:"AnswerKey,omitempty"`
	TestDurationInSeconds   int64                    `xml:"TestDurationInSeconds,omitempty"`
	RetryDelayInSeconds     int64                    `xml:"RetryDelayInSeconds,omitempty"`
	AutoGranted             bool                     `xml:"AutoGranted,omitempty"`
	AutoGrantedValue        int32                    `xml:"AutoGrantedValue,omitempty"`
	ResponseGroup           []string                 `xml:"ResponseGroup,omitempty"`
}

type GetAccountBalanceRequest struct {
	ResponseGroup []string `xml:"ResponseGroup,omitempty"`
	Unused        string   `xml:"Unused,omitempty"`
}

type GetAccountBalanceResult struct {
	Request          *Request `xml:"Request,omitempty"`
	AvailableBalance *Price   `xml:"AvailableBalance,omitempty"`
	OnHoldBalance    *Price   `xml:"OnHoldBalance,omitempty"`
}

type GetRequesterStatisticRequest struct {
	Statistic     *RequesterStatistic `xml:"Statistic,omitempty"`
	TimePeriod    *TimePeriod         `xml:"TimePeriod,omitempty"`
	Count         int32               `xml:"Count,omitempty"`
	ResponseGroup []string            `xml:"ResponseGroup,omitempty"`
}

type GetRequesterWorkerStatisticRequest struct {
	Statistic     *RequesterStatistic `xml:"Statistic,omitempty"`
	TimePeriod    *TimePeriod         `xml:"TimePeriod,omitempty"`
	WorkerId      string              `xml:"WorkerId,omitempty"`
	Count         int32               `xml:"Count,omitempty"`
	ResponseGroup []string            `xml:"ResponseGroup,omitempty"`
}

type NotifyWorkersFailureStatus struct {
	NotifyWorkersFailureCode    *NotifyWorkersFailureCode `xml:"NotifyWorkersFailureCode,omitempty"`
	NotifyWorkersFailureMessage string                    `xml:"NotifyWorkersFailureMessage,omitempty"`
	WorkerId                    string                    `xml:"WorkerId,omitempty"`
}

type NotifyWorkersRequest struct {
	Subject       string   `xml:"Subject,omitempty"`
	MessageText   string   `xml:"MessageText,omitempty"`
	WorkerId      string   `xml:"WorkerId,omitempty"`
	ResponseGroup []string `xml:"ResponseGroup,omitempty"`
}

type NotifyWorkersResult struct {
	Request                    *Request                    `xml:"Request,omitempty"`
	NotifyWorkersFailureStatus *NotifyWorkersFailureStatus `xml:"NotifyWorkersFailureStatus,omitempty"`
}

type GetStatisticResult struct {
	Request    *Request            `xml:"Request,omitempty"`
	DataPoint  []*DataPoint        `xml:"DataPoint,omitempty"`
	Statistic  *RequesterStatistic `xml:"Statistic,omitempty"`
	TimePeriod *TimePeriod         `xml:"TimePeriod,omitempty"`
	WorkerId   string              `xml:"WorkerId,omitempty"`
}

type DataPoint struct {
	Date        time.Time `xml:"Date,omitempty"`
	LongValue   int64     `xml:"LongValue,omitempty"`
	DoubleValue float64   `xml:"DoubleValue,omitempty"`
}

type GetBlockedWorkersRequest struct {
	PageNumber    int32    `xml:"PageNumber,omitempty"`
	PageSize      int32    `xml:"PageSize,omitempty"`
	ResponseGroup []string `xml:"ResponseGroup,omitempty"`
}

type GetBlockedWorkersResult struct {
	Request         *Request       `xml:"Request,omitempty"`
	PageNumber      int32          `xml:"PageNumber,omitempty"`
	NumResults      int32          `xml:"NumResults,omitempty"`
	TotalNumResults int32          `xml:"TotalNumResults,omitempty"`
	WorkerBlock     []*WorkerBlock `xml:"WorkerBlock,omitempty"`
}

type WorkerBlock struct {
	WorkerId string `xml:"WorkerId,omitempty"`
	Reason   string `xml:"Reason,omitempty"`
}

type BlockWorkerRequest struct {
	WorkerId      string   `xml:"WorkerId,omitempty"`
	Reason        string   `xml:"Reason,omitempty"`
	ResponseGroup []string `xml:"ResponseGroup,omitempty"`
}

type BlockWorkerResult struct {
	Request *Request `xml:"Request,omitempty"`
}

type UnblockWorkerRequest struct {
	WorkerId      string   `xml:"WorkerId,omitempty"`
	Reason        string   `xml:"Reason,omitempty"`
	ResponseGroup []string `xml:"ResponseGroup,omitempty"`
}

type UnblockWorkerResult struct {
	Request *Request `xml:"Request,omitempty"`
}

type KeyValuePair struct {
	Key   string `xml:"Key,omitempty"`
	Value string `xml:"Value,omitempty"`
}

type HelpRequest struct {
	About    []string `xml:"About,omitempty"`
	HelpType struct {
	} `xml:"HelpType,omitempty"`
	ResponseGroup []string `xml:"ResponseGroup,omitempty"`
}
