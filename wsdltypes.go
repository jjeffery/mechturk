package mturk

// NOTE: this file has been generated using gowsdl
// (github.com/hooklift/gowsdl). The file has, however,
// ben hand-edited to remove the operations and the SOAP
// code.
//
// TODO(jpj): Fork gowsdl and add an option to generate
// only type definitions. Could possibly add an option
// to supply a template file to use for generating options.
//
// TODO(jpj): Come up with a way to insert documentation
// for public types from the AWS documentation set.

import (
	"encoding/xml"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

type SortDirection string

const (
	SortDirectionAscending SortDirection = "Ascending"

	SortDirectionDescending SortDirection = "Descending"
)

type Comparator string

const (
	ComparatorLessThan Comparator = "LessThan"

	ComparatorLessThanOrEqualTo Comparator = "LessThanOrEqualTo"

	ComparatorGreaterThan Comparator = "GreaterThan"

	ComparatorGreaterThanOrEqualTo Comparator = "GreaterThanOrEqualTo"

	ComparatorEqualTo Comparator = "EqualTo"

	ComparatorNotEqualTo Comparator = "NotEqualTo"

	ComparatorExists Comparator = "Exists"
)

type EventType string

const (
	EventTypeAssignmentAccepted EventType = "AssignmentAccepted"

	EventTypeAssignmentAbandoned EventType = "AssignmentAbandoned"

	EventTypeAssignmentReturned EventType = "AssignmentReturned"

	EventTypeAssignmentSubmitted EventType = "AssignmentSubmitted"

	EventTypeHITReviewable EventType = "HITReviewable"

	EventTypeHITExpired EventType = "HITExpired"

	EventTypePing EventType = "Ping"
)

type NotificationTransport string

const (
	NotificationTransportSOAP NotificationTransport = "SOAP"

	NotificationTransportREST NotificationTransport = "REST"

	NotificationTransportEmail NotificationTransport = "Email"

	NotificationTransportSQS NotificationTransport = "SQS"
)

type HITStatus string

const (
	HITStatusAssignable HITStatus = "Assignable"

	HITStatusUnassignable HITStatus = "Unassignable"

	HITStatusReviewable HITStatus = "Reviewable"

	HITStatusReviewing HITStatus = "Reviewing"

	HITStatusDisposed HITStatus = "Disposed"
)

type HITReviewStatus string

const (
	HITReviewStatusNotReviewed HITReviewStatus = "NotReviewed"

	HITReviewStatusMarkedForReview HITReviewStatus = "MarkedForReview"

	HITReviewStatusReviewedAppropriate HITReviewStatus = "ReviewedAppropriate"

	HITReviewStatusReviewedInappropriate HITReviewStatus = "ReviewedInappropriate"
)

type ReviewableHITStatus string

const (
	ReviewableHITStatusReviewable ReviewableHITStatus = "Reviewable"

	ReviewableHITStatusReviewing ReviewableHITStatus = "Reviewing"
)

type GetReviewableHITsSortProperty string

const (
	GetReviewableHITsSortPropertyTitle GetReviewableHITsSortProperty = "Title"

	GetReviewableHITsSortPropertyReward GetReviewableHITsSortProperty = "Reward"

	GetReviewableHITsSortPropertyExpiration GetReviewableHITsSortProperty = "Expiration"

	GetReviewableHITsSortPropertyCreationTime GetReviewableHITsSortProperty = "CreationTime"

	GetReviewableHITsSortPropertyEnumeration GetReviewableHITsSortProperty = "Enumeration"
)

type ReviewPolicyLevel string

const (
	ReviewPolicyLevelAssignment ReviewPolicyLevel = "Assignment"

	ReviewPolicyLevelHIT ReviewPolicyLevel = "HIT"
)

type ReviewActionStatus string

const (
	ReviewActionStatusIntended ReviewActionStatus = "Intended"

	ReviewActionStatusSucceeded ReviewActionStatus = "Succeeded"

	ReviewActionStatusFailed ReviewActionStatus = "Failed"

	ReviewActionStatusCancelled ReviewActionStatus = "Cancelled"
)

type AssignmentStatus string

const (
	AssignmentStatusSubmitted AssignmentStatus = "Submitted"

	AssignmentStatusApproved AssignmentStatus = "Approved"

	AssignmentStatusRejected AssignmentStatus = "Rejected"
)

type GetAssignmentsForHITSortProperty string

const (
	GetAssignmentsForHITSortPropertyAcceptTime GetAssignmentsForHITSortProperty = "AcceptTime"

	GetAssignmentsForHITSortPropertySubmitTime GetAssignmentsForHITSortProperty = "SubmitTime"

	GetAssignmentsForHITSortPropertyAnswer GetAssignmentsForHITSortProperty = "Answer"

	GetAssignmentsForHITSortPropertyAssignmentStatus GetAssignmentsForHITSortProperty = "AssignmentStatus"
)

type SearchHITsSortProperty string

const (
	SearchHITsSortPropertyTitle SearchHITsSortProperty = "Title"

	SearchHITsSortPropertyReward SearchHITsSortProperty = "Reward"

	SearchHITsSortPropertyExpiration SearchHITsSortProperty = "Expiration"

	SearchHITsSortPropertyCreationTime SearchHITsSortProperty = "CreationTime"

	SearchHITsSortPropertyEnumeration SearchHITsSortProperty = "Enumeration"
)

type QualificationTypeStatus string

const (
	QualificationTypeStatusActive QualificationTypeStatus = "Active"

	QualificationTypeStatusInactive QualificationTypeStatus = "Inactive"
)

type GetQualificationRequestsSortProperty string

const (
	GetQualificationRequestsSortPropertyQualificationTypeId GetQualificationRequestsSortProperty = "QualificationTypeId"

	GetQualificationRequestsSortPropertySubmitTime GetQualificationRequestsSortProperty = "SubmitTime"
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
	RequesterStatisticNumberHITsAssignable RequesterStatistic = "NumberHITsAssignable"

	RequesterStatisticNumberHITsReviewable RequesterStatistic = "NumberHITsReviewable"

	RequesterStatisticNumberHITsCreated RequesterStatistic = "NumberHITsCreated"

	RequesterStatisticNumberHITsCompleted RequesterStatistic = "NumberHITsCompleted"

	RequesterStatisticTotalRewardPayout RequesterStatistic = "TotalRewardPayout"

	RequesterStatisticTotalRewardFeePayout RequesterStatistic = "TotalRewardFeePayout"

	RequesterStatisticTotalFeePayout RequesterStatistic = "TotalFeePayout"

	RequesterStatisticTotalRewardAndFeePayout RequesterStatistic = "TotalRewardAndFeePayout"

	RequesterStatisticTotalBonusPayout RequesterStatistic = "TotalBonusPayout"

	RequesterStatisticTotalBonusFeePayout RequesterStatistic = "TotalBonusFeePayout"

	RequesterStatisticEstimatedFeeLiability RequesterStatistic = "EstimatedFeeLiability"

	RequesterStatisticEstimatedRewardLiability RequesterStatistic = "EstimatedRewardLiability"

	RequesterStatisticEstimatedTotalLiability RequesterStatistic = "EstimatedTotalLiability"

	RequesterStatisticNumberAssignmentsAvailable RequesterStatistic = "NumberAssignmentsAvailable"

	RequesterStatisticNumberAssignmentsAccepted RequesterStatistic = "NumberAssignmentsAccepted"

	RequesterStatisticNumberAssignmentsPending RequesterStatistic = "NumberAssignmentsPending"

	RequesterStatisticNumberAssignmentsApproved RequesterStatistic = "NumberAssignmentsApproved"

	RequesterStatisticNumberAssignmentsRejected RequesterStatistic = "NumberAssignmentsRejected"

	RequesterStatisticNumberAssignmentsReturned RequesterStatistic = "NumberAssignmentsReturned"

	RequesterStatisticNumberAssignmentsAbandoned RequesterStatistic = "NumberAssignmentsAbandoned"

	RequesterStatisticAverageRewardAmount RequesterStatistic = "AverageRewardAmount"

	RequesterStatisticPercentAssignmentsApproved RequesterStatistic = "PercentAssignmentsApproved"

	RequesterStatisticPercentAssignmentsRejected RequesterStatistic = "PercentAssignmentsRejected"

	RequesterStatisticNumberKnownAnswersCorrect RequesterStatistic = "NumberKnownAnswersCorrect"

	RequesterStatisticNumberKnownAnswersIncorrect RequesterStatistic = "NumberKnownAnswersIncorrect"

	RequesterStatisticNumberKnownAnswersEvaluated RequesterStatistic = "NumberKnownAnswersEvaluated"

	RequesterStatisticPercentKnownAnswersCorrect RequesterStatistic = "PercentKnownAnswersCorrect"

	RequesterStatisticNumberPluralityAnswersCorrect RequesterStatistic = "NumberPluralityAnswersCorrect"

	RequesterStatisticNumberPluralityAnswersIncorrect RequesterStatistic = "NumberPluralityAnswersIncorrect"

	RequesterStatisticNumberPluralityAnswersEvaluated RequesterStatistic = "NumberPluralityAnswersEvaluated"

	RequesterStatisticPercentPluralityAnswersCorrect RequesterStatistic = "PercentPluralityAnswersCorrect"
)

type NotifyWorkersFailureCode string

const (
	NotifyWorkersFailureCodeSoftFailure NotifyWorkersFailureCode = "SoftFailure"

	NotifyWorkersFailureCodeHardFailure NotifyWorkersFailureCode = "HardFailure"
)

type TimePeriod string

const (
	TimePeriodOneDay TimePeriod = "OneDay"

	TimePeriodSevenDays TimePeriod = "SevenDays"

	TimePeriodThirtyDays TimePeriod = "ThirtyDays"

	TimePeriodLifeToDate TimePeriod = "LifeToDate"
)

type CreateHIT struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 CreateHIT"`

	AWSAccessKeyId string `xml:"AWSAccessKeyId,omitempty"`

	Timestamp time.Time `xml:"Timestamp,omitempty"`

	Signature string `xml:"Signature,omitempty"`

	Validate string `xml:"Validate,omitempty"`

	Credential string `xml:"Credential,omitempty"`

	Request []*CreateHITRequest `xml:"Request,omitempty"`
}

type CreateHITResponse struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 CreateHITResponse"`

	OperationRequest *OperationRequest `xml:"OperationRequest,omitempty"`

	HIT []*HIT `xml:"HIT,omitempty"`
}

type RegisterHITType struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 RegisterHITType"`

	AWSAccessKeyId string `xml:"AWSAccessKeyId,omitempty"`

	Timestamp time.Time `xml:"Timestamp,omitempty"`

	Signature string `xml:"Signature,omitempty"`

	Validate string `xml:"Validate,omitempty"`

	Credential string `xml:"Credential,omitempty"`

	Request []*RegisterHITTypeRequest `xml:"Request,omitempty"`
}

type RegisterHITTypeResponse struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 RegisterHITTypeResponse"`

	OperationRequest *OperationRequest `xml:"OperationRequest,omitempty"`

	RegisterHITTypeResult []*RegisterHITTypeResult `xml:"RegisterHITTypeResult,omitempty"`
}

type SetHITTypeNotification struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 SetHITTypeNotification"`

	AWSAccessKeyId string `xml:"AWSAccessKeyId,omitempty"`

	Timestamp time.Time `xml:"Timestamp,omitempty"`

	Signature string `xml:"Signature,omitempty"`

	Validate string `xml:"Validate,omitempty"`

	Credential string `xml:"Credential,omitempty"`

	Request []*SetHITTypeNotificationRequest `xml:"Request,omitempty"`
}

type SetHITTypeNotificationResponse struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 SetHITTypeNotificationResponse"`

	OperationRequest *OperationRequest `xml:"OperationRequest,omitempty"`

	SetHITTypeNotificationResult []*SetHITTypeNotificationResult `xml:"SetHITTypeNotificationResult,omitempty"`
}

type SendTestEventNotification struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 SendTestEventNotification"`

	AWSAccessKeyId string `xml:"AWSAccessKeyId,omitempty"`

	Timestamp time.Time `xml:"Timestamp,omitempty"`

	Signature string `xml:"Signature,omitempty"`

	Validate string `xml:"Validate,omitempty"`

	Credential string `xml:"Credential,omitempty"`

	Request []*SendTestEventNotificationRequest `xml:"Request,omitempty"`
}

type SendTestEventNotificationResponse struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 SendTestEventNotificationResponse"`

	OperationRequest *OperationRequest `xml:"OperationRequest,omitempty"`

	SendTestEventNotificationResult []*SendTestEventNotificationResult `xml:"SendTestEventNotificationResult,omitempty"`
}

type DisposeHIT struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 DisposeHIT"`

	AWSAccessKeyId string `xml:"AWSAccessKeyId,omitempty"`

	Timestamp time.Time `xml:"Timestamp,omitempty"`

	Signature string `xml:"Signature,omitempty"`

	Validate string `xml:"Validate,omitempty"`

	Credential string `xml:"Credential,omitempty"`

	Request []*DisposeHITRequest `xml:"Request,omitempty"`
}

type DisposeHITResponse struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 DisposeHITResponse"`

	OperationRequest *OperationRequest `xml:"OperationRequest,omitempty"`

	DisposeHITResult []*DisposeHITResult `xml:"DisposeHITResult,omitempty"`
}

type DisableHIT struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 DisableHIT"`

	AWSAccessKeyId string `xml:"AWSAccessKeyId,omitempty"`

	Timestamp time.Time `xml:"Timestamp,omitempty"`

	Signature string `xml:"Signature,omitempty"`

	Validate string `xml:"Validate,omitempty"`

	Credential string `xml:"Credential,omitempty"`

	Request []*DisableHITRequest `xml:"Request,omitempty"`
}

type DisableHITResponse struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 DisableHITResponse"`

	OperationRequest *OperationRequest `xml:"OperationRequest,omitempty"`

	DisableHITResult []*DisableHITResult `xml:"DisableHITResult,omitempty"`
}

type GetHIT struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GetHIT"`

	AWSAccessKeyId string `xml:"AWSAccessKeyId,omitempty"`

	Timestamp time.Time `xml:"Timestamp,omitempty"`

	Signature string `xml:"Signature,omitempty"`

	Validate string `xml:"Validate,omitempty"`

	Credential string `xml:"Credential,omitempty"`

	Request []*GetHITRequest `xml:"Request,omitempty"`
}

type GetHITResponse struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GetHITResponse"`

	OperationRequest *OperationRequest `xml:"OperationRequest,omitempty"`

	HIT []*HIT `xml:"HIT,omitempty"`
}

type GetAssignment struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GetAssignment"`

	AWSAccessKeyId string `xml:"AWSAccessKeyId,omitempty"`

	Timestamp time.Time `xml:"Timestamp,omitempty"`

	Signature string `xml:"Signature,omitempty"`

	Validate string `xml:"Validate,omitempty"`

	Credential string `xml:"Credential,omitempty"`

	Request []*GetAssignmentRequest `xml:"Request,omitempty"`
}

type GetAssignmentResponse struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GetAssignmentResponse"`

	OperationRequest *OperationRequest `xml:"OperationRequest,omitempty"`

	GetAssignmentResult []*GetAssignmentResult `xml:"GetAssignmentResult,omitempty"`
}

type GetReviewableHITs struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GetReviewableHITs"`

	AWSAccessKeyId string `xml:"AWSAccessKeyId,omitempty"`

	Timestamp time.Time `xml:"Timestamp,omitempty"`

	Signature string `xml:"Signature,omitempty"`

	Validate string `xml:"Validate,omitempty"`

	Credential string `xml:"Credential,omitempty"`

	Request []*GetReviewableHITsRequest `xml:"Request,omitempty"`
}

type GetReviewableHITsResponse struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GetReviewableHITsResponse"`

	OperationRequest *OperationRequest `xml:"OperationRequest,omitempty"`

	GetReviewableHITsResult []*GetReviewableHITsResult `xml:"GetReviewableHITsResult,omitempty"`
}

type GetReviewResultsForHIT struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GetReviewResultsForHIT"`

	AWSAccessKeyId string `xml:"AWSAccessKeyId,omitempty"`

	Timestamp time.Time `xml:"Timestamp,omitempty"`

	Signature string `xml:"Signature,omitempty"`

	Validate string `xml:"Validate,omitempty"`

	Credential string `xml:"Credential,omitempty"`

	Request []*GetReviewResultsForHITRequest `xml:"Request,omitempty"`
}

type GetReviewResultsForHITResponse struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GetReviewResultsForHITResponse"`

	OperationRequest *OperationRequest `xml:"OperationRequest,omitempty"`

	GetReviewResultsForHITResult []*GetReviewResultsForHITResult `xml:"GetReviewResultsForHITResult,omitempty"`
}

type GetHITsForQualificationType struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GetHITsForQualificationType"`

	AWSAccessKeyId string `xml:"AWSAccessKeyId,omitempty"`

	Timestamp time.Time `xml:"Timestamp,omitempty"`

	Signature string `xml:"Signature,omitempty"`

	Validate string `xml:"Validate,omitempty"`

	Credential string `xml:"Credential,omitempty"`

	Request []*GetHITsForQualificationTypeRequest `xml:"Request,omitempty"`
}

type GetHITsForQualificationTypeResponse struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GetHITsForQualificationTypeResponse"`

	OperationRequest *OperationRequest `xml:"OperationRequest,omitempty"`

	GetHITsForQualificationTypeResult []*GetHITsForQualificationTypeResult `xml:"GetHITsForQualificationTypeResult,omitempty"`
}

type GetQualificationsForQualificationType struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GetQualificationsForQualificationType"`

	AWSAccessKeyId string `xml:"AWSAccessKeyId,omitempty"`

	Timestamp time.Time `xml:"Timestamp,omitempty"`

	Signature string `xml:"Signature,omitempty"`

	Validate string `xml:"Validate,omitempty"`

	Credential string `xml:"Credential,omitempty"`

	Request []*GetQualificationsForQualificationTypeRequest `xml:"Request,omitempty"`
}

type GetQualificationsForQualificationTypeResponse struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GetQualificationsForQualificationTypeResponse"`

	OperationRequest *OperationRequest `xml:"OperationRequest,omitempty"`

	GetQualificationsForQualificationTypeResult []*GetQualificationsForQualificationTypeResult `xml:"GetQualificationsForQualificationTypeResult,omitempty"`
}

type SetHITAsReviewing struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 SetHITAsReviewing"`

	AWSAccessKeyId string `xml:"AWSAccessKeyId,omitempty"`

	Timestamp time.Time `xml:"Timestamp,omitempty"`

	Signature string `xml:"Signature,omitempty"`

	Validate string `xml:"Validate,omitempty"`

	Credential string `xml:"Credential,omitempty"`

	Request []*SetHITAsReviewingRequest `xml:"Request,omitempty"`
}

type SetHITAsReviewingResponse struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 SetHITAsReviewingResponse"`

	OperationRequest *OperationRequest `xml:"OperationRequest,omitempty"`

	SetHITAsReviewingResult []*SetHITAsReviewingResult `xml:"SetHITAsReviewingResult,omitempty"`
}

type ExtendHIT struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 ExtendHIT"`

	AWSAccessKeyId string `xml:"AWSAccessKeyId,omitempty"`

	Timestamp time.Time `xml:"Timestamp,omitempty"`

	Signature string `xml:"Signature,omitempty"`

	Validate string `xml:"Validate,omitempty"`

	Credential string `xml:"Credential,omitempty"`

	Request []*ExtendHITRequest `xml:"Request,omitempty"`
}

type ExtendHITResponse struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 ExtendHITResponse"`

	OperationRequest *OperationRequest `xml:"OperationRequest,omitempty"`

	ExtendHITResult []*ExtendHITResult `xml:"ExtendHITResult,omitempty"`
}

type ForceExpireHIT struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 ForceExpireHIT"`

	AWSAccessKeyId string `xml:"AWSAccessKeyId,omitempty"`

	Timestamp time.Time `xml:"Timestamp,omitempty"`

	Signature string `xml:"Signature,omitempty"`

	Validate string `xml:"Validate,omitempty"`

	Credential string `xml:"Credential,omitempty"`

	Request []*ForceExpireHITRequest `xml:"Request,omitempty"`
}

type ForceExpireHITResponse struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 ForceExpireHITResponse"`

	OperationRequest *OperationRequest `xml:"OperationRequest,omitempty"`

	ForceExpireHITResult []*ForceExpireHITResult `xml:"ForceExpireHITResult,omitempty"`
}

type ApproveAssignment struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 ApproveAssignment"`

	AWSAccessKeyId string `xml:"AWSAccessKeyId,omitempty"`

	Timestamp time.Time `xml:"Timestamp,omitempty"`

	Signature string `xml:"Signature,omitempty"`

	Validate string `xml:"Validate,omitempty"`

	Credential string `xml:"Credential,omitempty"`

	Request []*ApproveAssignmentRequest `xml:"Request,omitempty"`
}

type ApproveAssignmentResponse struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 ApproveAssignmentResponse"`

	OperationRequest *OperationRequest `xml:"OperationRequest,omitempty"`

	ApproveAssignmentResult []*ApproveAssignmentResult `xml:"ApproveAssignmentResult,omitempty"`
}

type RejectAssignment struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 RejectAssignment"`

	AWSAccessKeyId string `xml:"AWSAccessKeyId,omitempty"`

	Timestamp time.Time `xml:"Timestamp,omitempty"`

	Signature string `xml:"Signature,omitempty"`

	Validate string `xml:"Validate,omitempty"`

	Credential string `xml:"Credential,omitempty"`

	Request []*RejectAssignmentRequest `xml:"Request,omitempty"`
}

type RejectAssignmentResponse struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 RejectAssignmentResponse"`

	OperationRequest *OperationRequest `xml:"OperationRequest,omitempty"`

	RejectAssignmentResult []*RejectAssignmentResult `xml:"RejectAssignmentResult,omitempty"`
}

type ApproveRejectedAssignment struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 ApproveRejectedAssignment"`

	AWSAccessKeyId string `xml:"AWSAccessKeyId,omitempty"`

	Timestamp time.Time `xml:"Timestamp,omitempty"`

	Signature string `xml:"Signature,omitempty"`

	Validate string `xml:"Validate,omitempty"`

	Credential string `xml:"Credential,omitempty"`

	Request []*ApproveRejectedAssignmentRequest `xml:"Request,omitempty"`
}

type ApproveRejectedAssignmentResponse struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 ApproveRejectedAssignmentResponse"`

	OperationRequest *OperationRequest `xml:"OperationRequest,omitempty"`

	ApproveRejectedAssignmentResult []*ApproveRejectedAssignmentResult `xml:"ApproveRejectedAssignmentResult,omitempty"`
}

type GetAssignmentsForHIT struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GetAssignmentsForHIT"`

	AWSAccessKeyId string `xml:"AWSAccessKeyId,omitempty"`

	Timestamp time.Time `xml:"Timestamp,omitempty"`

	Signature string `xml:"Signature,omitempty"`

	Validate string `xml:"Validate,omitempty"`

	Credential string `xml:"Credential,omitempty"`

	Request []*GetAssignmentsForHITRequest `xml:"Request,omitempty"`
}

type GetAssignmentsForHITResponse struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GetAssignmentsForHITResponse"`

	OperationRequest *OperationRequest `xml:"OperationRequest,omitempty"`

	GetAssignmentsForHITResult []*GetAssignmentsForHITResult `xml:"GetAssignmentsForHITResult,omitempty"`
}

type GetFileUploadURL struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GetFileUploadURL"`

	AWSAccessKeyId string `xml:"AWSAccessKeyId,omitempty"`

	Timestamp time.Time `xml:"Timestamp,omitempty"`

	Signature string `xml:"Signature,omitempty"`

	Validate string `xml:"Validate,omitempty"`

	Credential string `xml:"Credential,omitempty"`

	Request []*GetFileUploadURLRequest `xml:"Request,omitempty"`
}

type GetFileUploadURLResponse struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GetFileUploadURLResponse"`

	OperationRequest *OperationRequest `xml:"OperationRequest,omitempty"`

	GetFileUploadURLResult []*GetFileUploadURLResult `xml:"GetFileUploadURLResult,omitempty"`
}

type SearchHITs struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 SearchHITs"`

	AWSAccessKeyId string `xml:"AWSAccessKeyId,omitempty"`

	Timestamp time.Time `xml:"Timestamp,omitempty"`

	Signature string `xml:"Signature,omitempty"`

	Validate string `xml:"Validate,omitempty"`

	Credential string `xml:"Credential,omitempty"`

	Request []*SearchHITsRequest `xml:"Request,omitempty"`
}

type SearchHITsResponse struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 SearchHITsResponse"`

	OperationRequest *OperationRequest `xml:"OperationRequest,omitempty"`

	SearchHITsResult []*SearchHITsResult `xml:"SearchHITsResult,omitempty"`
}

type GrantBonus struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GrantBonus"`

	AWSAccessKeyId string `xml:"AWSAccessKeyId,omitempty"`

	Timestamp time.Time `xml:"Timestamp,omitempty"`

	Signature string `xml:"Signature,omitempty"`

	Validate string `xml:"Validate,omitempty"`

	Credential string `xml:"Credential,omitempty"`

	Request []*GrantBonusRequest `xml:"Request,omitempty"`
}

type GrantBonusResponse struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GrantBonusResponse"`

	OperationRequest *OperationRequest `xml:"OperationRequest,omitempty"`

	GrantBonusResult []*GrantBonusResult `xml:"GrantBonusResult,omitempty"`
}

type GetBonusPayments struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GetBonusPayments"`

	AWSAccessKeyId string `xml:"AWSAccessKeyId,omitempty"`

	Timestamp time.Time `xml:"Timestamp,omitempty"`

	Signature string `xml:"Signature,omitempty"`

	Validate string `xml:"Validate,omitempty"`

	Credential string `xml:"Credential,omitempty"`

	Request []*GetBonusPaymentsRequest `xml:"Request,omitempty"`
}

type GetBonusPaymentsResponse struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GetBonusPaymentsResponse"`

	OperationRequest *OperationRequest `xml:"OperationRequest,omitempty"`

	GetBonusPaymentsResult []*GetBonusPaymentsResult `xml:"GetBonusPaymentsResult,omitempty"`
}

type ChangeHITTypeOfHIT struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 ChangeHITTypeOfHIT"`

	AWSAccessKeyId string `xml:"AWSAccessKeyId,omitempty"`

	Timestamp time.Time `xml:"Timestamp,omitempty"`

	Signature string `xml:"Signature,omitempty"`

	Validate string `xml:"Validate,omitempty"`

	Credential string `xml:"Credential,omitempty"`

	Request []*ChangeHITTypeOfHITRequest `xml:"Request,omitempty"`
}

type ChangeHITTypeOfHITResponse struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 ChangeHITTypeOfHITResponse"`

	OperationRequest *OperationRequest `xml:"OperationRequest,omitempty"`

	ChangeHITTypeOfHITResult []*ChangeHITTypeOfHITResult `xml:"ChangeHITTypeOfHITResult,omitempty"`
}

type CreateQualificationType struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 CreateQualificationType"`

	AWSAccessKeyId string `xml:"AWSAccessKeyId,omitempty"`

	Timestamp time.Time `xml:"Timestamp,omitempty"`

	Signature string `xml:"Signature,omitempty"`

	Validate string `xml:"Validate,omitempty"`

	Credential string `xml:"Credential,omitempty"`

	Request []*CreateQualificationTypeRequest `xml:"Request,omitempty"`
}

type CreateQualificationTypeResponse struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 CreateQualificationTypeResponse"`

	OperationRequest *OperationRequest `xml:"OperationRequest,omitempty"`

	QualificationType []*QualificationType `xml:"QualificationType,omitempty"`
}

type DisposeQualificationType struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 DisposeQualificationType"`

	AWSAccessKeyId string `xml:"AWSAccessKeyId,omitempty"`

	Timestamp time.Time `xml:"Timestamp,omitempty"`

	Signature string `xml:"Signature,omitempty"`

	Validate string `xml:"Validate,omitempty"`

	Credential string `xml:"Credential,omitempty"`

	Request []*DisposeQualificationTypeRequest `xml:"Request,omitempty"`
}

type DisposeQualificationTypeResponse struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 DisposeQualificationTypeResponse"`

	OperationRequest *OperationRequest `xml:"OperationRequest,omitempty"`

	DisposeQualificationTypeResult []*DisposeQualificationTypeResult `xml:"DisposeQualificationTypeResult,omitempty"`
}

type GetQualificationRequests struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GetQualificationRequests"`

	AWSAccessKeyId string `xml:"AWSAccessKeyId,omitempty"`

	Timestamp time.Time `xml:"Timestamp,omitempty"`

	Signature string `xml:"Signature,omitempty"`

	Validate string `xml:"Validate,omitempty"`

	Credential string `xml:"Credential,omitempty"`

	Request []*GetQualificationRequestsRequest `xml:"Request,omitempty"`
}

type GetQualificationRequestsResponse struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GetQualificationRequestsResponse"`

	OperationRequest *OperationRequest `xml:"OperationRequest,omitempty"`

	GetQualificationRequestsResult []*GetQualificationRequestsResult `xml:"GetQualificationRequestsResult,omitempty"`
}

type RejectQualificationRequest struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 RejectQualificationRequest"`

	AWSAccessKeyId string `xml:"AWSAccessKeyId,omitempty"`

	Timestamp time.Time `xml:"Timestamp,omitempty"`

	Signature string `xml:"Signature,omitempty"`

	Validate string `xml:"Validate,omitempty"`

	Credential string `xml:"Credential,omitempty"`

	Request []*RejectQualificationRequestRequest `xml:"Request,omitempty"`
}

type RejectQualificationRequestResponse struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 RejectQualificationRequestResponse"`

	OperationRequest *OperationRequest `xml:"OperationRequest,omitempty"`

	RejectQualificationRequestResult []*RejectQualificationRequestResult `xml:"RejectQualificationRequestResult,omitempty"`
}

type GrantQualification struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GrantQualification"`

	AWSAccessKeyId string `xml:"AWSAccessKeyId,omitempty"`

	Timestamp time.Time `xml:"Timestamp,omitempty"`

	Signature string `xml:"Signature,omitempty"`

	Validate string `xml:"Validate,omitempty"`

	Credential string `xml:"Credential,omitempty"`

	Request []*GrantQualificationRequest `xml:"Request,omitempty"`
}

type GrantQualificationResponse struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GrantQualificationResponse"`

	OperationRequest *OperationRequest `xml:"OperationRequest,omitempty"`

	GrantQualificationResult []*GrantQualificationResult `xml:"GrantQualificationResult,omitempty"`
}

type AssignQualification struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 AssignQualification"`

	AWSAccessKeyId string `xml:"AWSAccessKeyId,omitempty"`

	Timestamp time.Time `xml:"Timestamp,omitempty"`

	Signature string `xml:"Signature,omitempty"`

	Validate string `xml:"Validate,omitempty"`

	Credential string `xml:"Credential,omitempty"`

	Request []*AssignQualificationRequest `xml:"Request,omitempty"`
}

type AssignQualificationResponse struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 AssignQualificationResponse"`

	OperationRequest *OperationRequest `xml:"OperationRequest,omitempty"`

	AssignQualificationResult []*AssignQualificationResult `xml:"AssignQualificationResult,omitempty"`
}

type RevokeQualification struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 RevokeQualification"`

	AWSAccessKeyId string `xml:"AWSAccessKeyId,omitempty"`

	Timestamp time.Time `xml:"Timestamp,omitempty"`

	Signature string `xml:"Signature,omitempty"`

	Validate string `xml:"Validate,omitempty"`

	Credential string `xml:"Credential,omitempty"`

	Request []*RevokeQualificationRequest `xml:"Request,omitempty"`
}

type RevokeQualificationResponse struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 RevokeQualificationResponse"`

	OperationRequest *OperationRequest `xml:"OperationRequest,omitempty"`

	RevokeQualificationResult []*RevokeQualificationResult `xml:"RevokeQualificationResult,omitempty"`
}

type UpdateQualificationScore struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 UpdateQualificationScore"`

	AWSAccessKeyId string `xml:"AWSAccessKeyId,omitempty"`

	Timestamp time.Time `xml:"Timestamp,omitempty"`

	Signature string `xml:"Signature,omitempty"`

	Validate string `xml:"Validate,omitempty"`

	Credential string `xml:"Credential,omitempty"`

	Request []*UpdateQualificationScoreRequest `xml:"Request,omitempty"`
}

type UpdateQualificationScoreResponse struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 UpdateQualificationScoreResponse"`

	OperationRequest *OperationRequest `xml:"OperationRequest,omitempty"`

	UpdateQualificationScoreResult []*UpdateQualificationScoreResult `xml:"UpdateQualificationScoreResult,omitempty"`
}

type GetQualificationType struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GetQualificationType"`

	AWSAccessKeyId string `xml:"AWSAccessKeyId,omitempty"`

	Timestamp time.Time `xml:"Timestamp,omitempty"`

	Signature string `xml:"Signature,omitempty"`

	Validate string `xml:"Validate,omitempty"`

	Credential string `xml:"Credential,omitempty"`

	Request []*GetQualificationTypeRequest `xml:"Request,omitempty"`
}

type GetQualificationTypeResponse struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GetQualificationTypeResponse"`

	OperationRequest *OperationRequest `xml:"OperationRequest,omitempty"`

	QualificationType []*QualificationType `xml:"QualificationType,omitempty"`
}

type GetQualificationScore struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GetQualificationScore"`

	AWSAccessKeyId string `xml:"AWSAccessKeyId,omitempty"`

	Timestamp time.Time `xml:"Timestamp,omitempty"`

	Signature string `xml:"Signature,omitempty"`

	Validate string `xml:"Validate,omitempty"`

	Credential string `xml:"Credential,omitempty"`

	Request []*GetQualificationScoreRequest `xml:"Request,omitempty"`
}

type GetQualificationScoreResponse struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GetQualificationScoreResponse"`

	OperationRequest *OperationRequest `xml:"OperationRequest,omitempty"`

	Qualification []*Qualification `xml:"Qualification,omitempty"`
}

type SearchQualificationTypes struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 SearchQualificationTypes"`

	AWSAccessKeyId string `xml:"AWSAccessKeyId,omitempty"`

	Timestamp time.Time `xml:"Timestamp,omitempty"`

	Signature string `xml:"Signature,omitempty"`

	Validate string `xml:"Validate,omitempty"`

	Credential string `xml:"Credential,omitempty"`

	Request []*SearchQualificationTypesRequest `xml:"Request,omitempty"`
}

type SearchQualificationTypesResponse struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 SearchQualificationTypesResponse"`

	OperationRequest *OperationRequest `xml:"OperationRequest,omitempty"`

	SearchQualificationTypesResult []*SearchQualificationTypesResult `xml:"SearchQualificationTypesResult,omitempty"`
}

type UpdateQualificationType struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 UpdateQualificationType"`

	AWSAccessKeyId string `xml:"AWSAccessKeyId,omitempty"`

	Timestamp time.Time `xml:"Timestamp,omitempty"`

	Signature string `xml:"Signature,omitempty"`

	Validate string `xml:"Validate,omitempty"`

	Credential string `xml:"Credential,omitempty"`

	Request []*UpdateQualificationTypeRequest `xml:"Request,omitempty"`
}

type UpdateQualificationTypeResponse struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 UpdateQualificationTypeResponse"`

	OperationRequest *OperationRequest `xml:"OperationRequest,omitempty"`

	QualificationType []*QualificationType `xml:"QualificationType,omitempty"`
}

type GetAccountBalance struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GetAccountBalance"`

	AWSAccessKeyId string `xml:"AWSAccessKeyId,omitempty"`

	Timestamp time.Time `xml:"Timestamp,omitempty"`

	Signature string `xml:"Signature,omitempty"`

	Validate string `xml:"Validate,omitempty"`

	Credential string `xml:"Credential,omitempty"`

	Request []*GetAccountBalanceRequest `xml:"Request,omitempty"`
}

type GetAccountBalanceResponse struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GetAccountBalanceResponse"`

	OperationRequest *OperationRequest `xml:"OperationRequest,omitempty"`

	GetAccountBalanceResult []*GetAccountBalanceResult `xml:"GetAccountBalanceResult,omitempty"`
}

type GetRequesterStatistic struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GetRequesterStatistic"`

	AWSAccessKeyId string `xml:"AWSAccessKeyId,omitempty"`

	Timestamp time.Time `xml:"Timestamp,omitempty"`

	Signature string `xml:"Signature,omitempty"`

	Validate string `xml:"Validate,omitempty"`

	Credential string `xml:"Credential,omitempty"`

	Request []*GetRequesterStatisticRequest `xml:"Request,omitempty"`
}

type GetRequesterStatisticResponse struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GetRequesterStatisticResponse"`

	OperationRequest *OperationRequest `xml:"OperationRequest,omitempty"`

	GetStatisticResult []*GetStatisticResult `xml:"GetStatisticResult,omitempty"`
}

type GetRequesterWorkerStatistic struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GetRequesterWorkerStatistic"`

	AWSAccessKeyId string `xml:"AWSAccessKeyId,omitempty"`

	Timestamp time.Time `xml:"Timestamp,omitempty"`

	Signature string `xml:"Signature,omitempty"`

	Validate string `xml:"Validate,omitempty"`

	Credential string `xml:"Credential,omitempty"`

	Request []*GetRequesterWorkerStatisticRequest `xml:"Request,omitempty"`
}

type GetRequesterWorkerStatisticResponse struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GetRequesterWorkerStatisticResponse"`

	OperationRequest *OperationRequest `xml:"OperationRequest,omitempty"`

	GetStatisticResult []*GetStatisticResult `xml:"GetStatisticResult,omitempty"`
}

type NotifyWorkers struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 NotifyWorkers"`

	AWSAccessKeyId string `xml:"AWSAccessKeyId,omitempty"`

	Timestamp time.Time `xml:"Timestamp,omitempty"`

	Signature string `xml:"Signature,omitempty"`

	Validate string `xml:"Validate,omitempty"`

	Credential string `xml:"Credential,omitempty"`

	Request []*NotifyWorkersRequest `xml:"Request,omitempty"`
}

type NotifyWorkersResponse struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 NotifyWorkersResponse"`

	OperationRequest *OperationRequest `xml:"OperationRequest,omitempty"`

	NotifyWorkersResult []*NotifyWorkersResult `xml:"NotifyWorkersResult,omitempty"`
}

type GetBlockedWorkers struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GetBlockedWorkers"`

	AWSAccessKeyId string `xml:"AWSAccessKeyId,omitempty"`

	Timestamp time.Time `xml:"Timestamp,omitempty"`

	Signature string `xml:"Signature,omitempty"`

	Validate string `xml:"Validate,omitempty"`

	Credential string `xml:"Credential,omitempty"`

	Request []*GetBlockedWorkersRequest `xml:"Request,omitempty"`
}

type GetBlockedWorkersResponse struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GetBlockedWorkersResponse"`

	OperationRequest *OperationRequest `xml:"OperationRequest,omitempty"`

	GetBlockedWorkersResult []*GetBlockedWorkersResult `xml:"GetBlockedWorkersResult,omitempty"`
}

type BlockWorker struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 BlockWorker"`

	AWSAccessKeyId string `xml:"AWSAccessKeyId,omitempty"`

	Timestamp time.Time `xml:"Timestamp,omitempty"`

	Signature string `xml:"Signature,omitempty"`

	Validate string `xml:"Validate,omitempty"`

	Credential string `xml:"Credential,omitempty"`

	Request []*BlockWorkerRequest `xml:"Request,omitempty"`
}

type BlockWorkerResponse struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 BlockWorkerResponse"`

	OperationRequest *OperationRequest `xml:"OperationRequest,omitempty"`

	BlockWorkerResult []*BlockWorkerResult `xml:"BlockWorkerResult,omitempty"`
}

type UnblockWorker struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 UnblockWorker"`

	AWSAccessKeyId string `xml:"AWSAccessKeyId,omitempty"`

	Timestamp time.Time `xml:"Timestamp,omitempty"`

	Signature string `xml:"Signature,omitempty"`

	Validate string `xml:"Validate,omitempty"`

	Credential string `xml:"Credential,omitempty"`

	Request []*UnblockWorkerRequest `xml:"Request,omitempty"`
}

type UnblockWorkerResponse struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 UnblockWorkerResponse"`

	OperationRequest *OperationRequest `xml:"OperationRequest,omitempty"`

	UnblockWorkerResult []*UnblockWorkerResult `xml:"UnblockWorkerResult,omitempty"`
}

type OperationRequest struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 OperationRequest"`

	HTTPHeaders *HTTPHeaders `xml:"HTTPHeaders,omitempty"`

	RequestId string `xml:"RequestId,omitempty"`

	Arguments *Arguments `xml:"Arguments,omitempty"`

	Errors *Errors `xml:"Errors,omitempty"`
}

type Request struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 Request"`

	IsValid string `xml:"IsValid,omitempty"`

	CreateHITRequest *CreateHITRequest `xml:"CreateHITRequest,omitempty"`

	RegisterHITTypeRequest *RegisterHITTypeRequest `xml:"RegisterHITTypeRequest,omitempty"`

	DisposeHITRequest *DisposeHITRequest `xml:"DisposeHITRequest,omitempty"`

	DisableHITRequest *DisableHITRequest `xml:"DisableHITRequest,omitempty"`

	GetHITRequest *GetHITRequest `xml:"GetHITRequest,omitempty"`

	GetAssignmentRequest *GetAssignmentRequest `xml:"GetAssignmentRequest,omitempty"`

	GetReviewResultsForHITRequest *GetReviewResultsForHITRequest `xml:"GetReviewResultsForHITRequest,omitempty"`

	GetReviewableHITsRequest *GetReviewableHITsRequest `xml:"GetReviewableHITsRequest,omitempty"`

	GetHITsForQualificationTypeRequest *GetHITsForQualificationTypeRequest `xml:"GetHITsForQualificationTypeRequest,omitempty"`

	GetQualificationsForQualificationTypeRequest *GetQualificationsForQualificationTypeRequest `xml:"GetQualificationsForQualificationTypeRequest,omitempty"`

	SetHITAsReviewingRequest *SetHITAsReviewingRequest `xml:"SetHITAsReviewingRequest,omitempty"`

	SearchHITsRequest *SearchHITsRequest `xml:"SearchHITsRequest,omitempty"`

	ExtendHITRequest *ExtendHITRequest `xml:"ExtendHITRequest,omitempty"`

	ForceExpireHITRequest *ForceExpireHITRequest `xml:"ForceExpireHITRequest,omitempty"`

	ChangeHITTypeOfHITRequest *ChangeHITTypeOfHITRequest `xml:"ChangeHITTypeOfHITRequest,omitempty"`

	CreateQualificationTypeRequest *CreateQualificationTypeRequest `xml:"CreateQualificationTypeRequest,omitempty"`

	DisposeQualificationTypeRequest *DisposeQualificationTypeRequest `xml:"DisposeQualificationTypeRequest,omitempty"`

	GrantQualificationRequest *GrantQualificationRequest `xml:"GrantQualificationRequest,omitempty"`

	AssignQualificationRequest *AssignQualificationRequest `xml:"AssignQualificationRequest,omitempty"`

	RevokeQualificationRequest *RevokeQualificationRequest `xml:"RevokeQualificationRequest,omitempty"`

	GetQualificationRequestsRequest *GetQualificationRequestsRequest `xml:"GetQualificationRequestsRequest,omitempty"`

	RejectQualificationRequestRequest *RejectQualificationRequestRequest `xml:"RejectQualificationRequestRequest,omitempty"`

	GetQualificationTypeRequest *GetQualificationTypeRequest `xml:"GetQualificationTypeRequest,omitempty"`

	SearchQualificationTypesRequest *SearchQualificationTypesRequest `xml:"SearchQualificationTypesRequest,omitempty"`

	UpdateQualificationTypeRequest *UpdateQualificationTypeRequest `xml:"UpdateQualificationTypeRequest,omitempty"`

	ApproveAssignmentRequest *ApproveAssignmentRequest `xml:"ApproveAssignmentRequest,omitempty"`

	RejectAssignmentRequest *RejectAssignmentRequest `xml:"RejectAssignmentRequest,omitempty"`

	ApproveRejectedAssignmentRequest *ApproveRejectedAssignmentRequest `xml:"ApproveRejectedAssignmentRequest,omitempty"`

	GetAssignmentsForHIT *GetAssignmentsForHITRequest `xml:"GetAssignmentsForHIT,omitempty"`

	GetFileUploadURL *GetFileUploadURLRequest `xml:"GetFileUploadURL,omitempty"`

	GrantBonusRequest *GrantBonusRequest `xml:"GrantBonusRequest,omitempty"`

	GetBonusPaymentsRequest *GetBonusPaymentsRequest `xml:"GetBonusPaymentsRequest,omitempty"`

	GetAccountBalanceRequest *GetAccountBalanceRequest `xml:"GetAccountBalanceRequest,omitempty"`

	NotifyWorkersRequest *NotifyWorkersRequest `xml:"NotifyWorkersRequest,omitempty"`

	GetBlockedWorkersRequest *GetBlockedWorkersRequest `xml:"GetBlockedWorkersRequest,omitempty"`

	BlockWorkerRequest *BlockWorkerRequest `xml:"BlockWorkerRequest,omitempty"`

	UnblockWorkerRequest *UnblockWorkerRequest `xml:"UnblockWorkerRequest,omitempty"`

	GetRequesterStatistic *GetRequesterStatisticRequest `xml:"GetRequesterStatistic,omitempty"`

	GetRequesterWorkerStatistic *GetRequesterWorkerStatisticRequest `xml:"GetRequesterWorkerStatistic,omitempty"`

	HelpRequest *HelpRequest `xml:"HelpRequest,omitempty"`

	Errors *Errors `xml:"Errors,omitempty"`
}

type HTTPHeaders struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 HTTPHeaders"`

	Header struct {
		Name string `xml:"Name,attr,omitempty"`

		Value string `xml:"Value,attr,omitempty"`
	} `xml:"Header,omitempty"`
}

type Arguments struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 Arguments"`

	Argument struct {
		Name string `xml:"Name,attr,omitempty"`

		Value string `xml:"Value,attr,omitempty"`
	} `xml:"Argument,omitempty"`
}

type Errors struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 Errors"`

	Error struct {
		Code string `xml:"Code,omitempty"`

		Message string `xml:"Message,omitempty"`

		Data []*KeyValuePair `xml:"Data,omitempty"`
	} `xml:"Error,omitempty"`
}

type Help struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 Help"`

	AWSAccessKeyId string `xml:"AWSAccessKeyId,omitempty"`

	Timestamp time.Time `xml:"Timestamp,omitempty"`

	Signature string `xml:"Signature,omitempty"`

	Validate string `xml:"Validate,omitempty"`

	Request []*HelpRequest `xml:"Request,omitempty"`
}

type HelpResponse struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 HelpResponse"`

	OperationRequest *OperationRequest `xml:"OperationRequest,omitempty"`

	Information []*Information `xml:"Information,omitempty"`
}

type Information struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 Information"`

	Request *Request `xml:"Request,omitempty"`

	OperationInformation []*OperationInformation `xml:"OperationInformation,omitempty"`

	ResponseGroupInformation []*ResponseGroupInformation `xml:"ResponseGroupInformation,omitempty"`
}

type OperationInformation struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 OperationInformation"`

	Name string `xml:"Name,omitempty"`

	Description string `xml:"Description,omitempty"`

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
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 ResponseGroupInformation"`

	Name string `xml:"Name,omitempty"`

	CreationDate string `xml:"CreationDate,omitempty"`

	ValidOperations struct {
		Operation []string `xml:"Operation,omitempty"`
	} `xml:"ValidOperations,omitempty"`

	Elements struct {
		Element []string `xml:"Element,omitempty"`
	} `xml:"Elements,omitempty"`
}

type Claim struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 Claim"`

	Type string `xml:"Type,omitempty"`

	Body string `xml:"Body,omitempty"`
}

type Price struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 Price"`

	Amount float64 `xml:"Amount,omitempty"`

	CurrencyCode string `xml:"CurrencyCode,omitempty"`

	FormattedPrice string `xml:"FormattedPrice,omitempty"`
}

type QualificationRequirement struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 QualificationRequirement"`

	QualificationTypeId string `xml:"QualificationTypeId,omitempty"`

	Comparator *Comparator `xml:"Comparator,omitempty"`

	IntegerValue int32 `xml:"IntegerValue,omitempty"`

	LocaleValue *Locale `xml:"LocaleValue,omitempty"`

	RequiredToPreview bool `xml:"RequiredToPreview,omitempty"`
}

type Locale struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 Locale"`

	Country string `xml:"Country,omitempty"`
}

type NotificationSpecification struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 NotificationSpecification"`

	Destination string `xml:"Destination,omitempty"`

	Transport *NotificationTransport `xml:"Transport,omitempty"`

	Version string `xml:"Version,omitempty"`

	EventType []*EventType `xml:"EventType,omitempty"`
}

type HIT struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 HIT"`

	Request *Request `xml:"Request,omitempty"`

	HITId string `xml:"HITId,omitempty"`

	HITTypeId string `xml:"HITTypeId,omitempty"`

	HITGroupId string `xml:"HITGroupId,omitempty"`

	HITLayoutId string `xml:"HITLayoutId,omitempty"`

	CreationTime time.Time `xml:"CreationTime,omitempty"`

	Title string `xml:"Title,omitempty"`

	Description string `xml:"Description,omitempty"`

	Question string `xml:"Question,omitempty"`

	Keywords string `xml:"Keywords,omitempty"`

	HITStatus *HITStatus `xml:"HITStatus,omitempty"`

	MaxAssignments int32 `xml:"MaxAssignments,omitempty"`

	Reward *Price `xml:"Reward,omitempty"`

	AutoApprovalDelayInSeconds int64 `xml:"AutoApprovalDelayInSeconds,omitempty"`

	Expiration time.Time `xml:"Expiration,omitempty"`

	AssignmentDurationInSeconds int64 `xml:"AssignmentDurationInSeconds,omitempty"`

	RequesterAnnotation string `xml:"RequesterAnnotation,omitempty"`

	QualificationRequirement []*QualificationRequirement `xml:"QualificationRequirement,omitempty"`

	HITReviewStatus *HITReviewStatus `xml:"HITReviewStatus,omitempty"`

	NumberOfAssignmentsPending int32 `xml:"NumberOfAssignmentsPending,omitempty"`

	NumberOfAssignmentsAvailable int32 `xml:"NumberOfAssignmentsAvailable,omitempty"`

	NumberOfAssignmentsCompleted int32 `xml:"NumberOfAssignmentsCompleted,omitempty"`
}

type ReviewPolicy struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 ReviewPolicy"`

	PolicyName string `xml:"PolicyName,omitempty"`

	Parameter []*PolicyParameter `xml:"Parameter,omitempty"`
}

type PolicyParameter struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 PolicyParameter"`

	Key string `xml:"Key,omitempty"`

	Value []string `xml:"Value,omitempty"`

	MapEntry []*ParameterMapEntry `xml:"MapEntry,omitempty"`
}

type ParameterMapEntry struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 ParameterMapEntry"`

	Key string `xml:"Key,omitempty"`

	Value []string `xml:"Value,omitempty"`
}

type HITLayoutParameter struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 HITLayoutParameter"`

	Name string `xml:"Name,omitempty"`

	Value string `xml:"Value,omitempty"`
}

type CreateHITRequest struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 CreateHITRequest"`

	HITTypeId string `xml:"HITTypeId,omitempty"`

	MaxAssignments int32 `xml:"MaxAssignments,omitempty"`

	AutoApprovalDelayInSeconds int64 `xml:"AutoApprovalDelayInSeconds,omitempty"`

	LifetimeInSeconds int64 `xml:"LifetimeInSeconds,omitempty"`

	AssignmentDurationInSeconds int64 `xml:"AssignmentDurationInSeconds,omitempty"`

	Reward *Price `xml:"Reward,omitempty"`

	Title string `xml:"Title,omitempty"`

	Keywords string `xml:"Keywords,omitempty"`

	Description string `xml:"Description,omitempty"`

	Question string `xml:"Question,omitempty"`

	RequesterAnnotation string `xml:"RequesterAnnotation,omitempty"`

	QualificationRequirement []*QualificationRequirement `xml:"QualificationRequirement,omitempty"`

	UniqueRequestToken string `xml:"UniqueRequestToken,omitempty"`

	AssignmentReviewPolicy *ReviewPolicy `xml:"AssignmentReviewPolicy,omitempty"`

	HITReviewPolicy *ReviewPolicy `xml:"HITReviewPolicy,omitempty"`

	HITLayoutId string `xml:"HITLayoutId,omitempty"`

	HITLayoutParameter []*HITLayoutParameter `xml:"HITLayoutParameter,omitempty"`

	ResponseGroup []string `xml:"ResponseGroup,omitempty"`
}

type RegisterHITTypeRequest struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 RegisterHITTypeRequest"`

	AutoApprovalDelayInSeconds int64 `xml:"AutoApprovalDelayInSeconds,omitempty"`

	AssignmentDurationInSeconds int64 `xml:"AssignmentDurationInSeconds,omitempty"`

	Reward *Price `xml:"Reward,omitempty"`

	Title string `xml:"Title,omitempty"`

	Keywords string `xml:"Keywords,omitempty"`

	Description string `xml:"Description,omitempty"`

	QualificationRequirement []*QualificationRequirement `xml:"QualificationRequirement,omitempty"`

	ResponseGroup []string `xml:"ResponseGroup,omitempty"`
}

type RegisterHITTypeResult struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 RegisterHITTypeResult"`

	Request *Request `xml:"Request,omitempty"`

	HITTypeId string `xml:"HITTypeId,omitempty"`
}

type SetHITTypeNotificationRequest struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 SetHITTypeNotificationRequest"`

	HITTypeId string `xml:"HITTypeId,omitempty"`

	Notification *NotificationSpecification `xml:"Notification,omitempty"`

	Active bool `xml:"Active,omitempty"`
}

type SetHITTypeNotificationResult struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 SetHITTypeNotificationResult"`

	Request *Request `xml:"Request,omitempty"`
}

type SendTestEventNotificationRequest struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 SendTestEventNotificationRequest"`

	Notification *NotificationSpecification `xml:"Notification,omitempty"`

	TestEventType *EventType `xml:"TestEventType,omitempty"`
}

type SendTestEventNotificationResult struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 SendTestEventNotificationResult"`

	Request *Request `xml:"Request,omitempty"`
}

type DisposeHITRequest struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 DisposeHITRequest"`

	HITId string `xml:"HITId,omitempty"`

	ResponseGroup []string `xml:"ResponseGroup,omitempty"`
}

type DisposeHITResult struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 DisposeHITResult"`

	Request *Request `xml:"Request,omitempty"`
}

type DisableHITRequest struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 DisableHITRequest"`

	HITId string `xml:"HITId,omitempty"`

	ResponseGroup []string `xml:"ResponseGroup,omitempty"`
}

type DisableHITResult struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 DisableHITResult"`

	Request *Request `xml:"Request,omitempty"`
}

type GetHITRequest struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GetHITRequest"`

	HITId string `xml:"HITId,omitempty"`

	ResponseGroup []string `xml:"ResponseGroup,omitempty"`
}

type GetAssignmentRequest struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GetAssignmentRequest"`

	AssignmentId string `xml:"AssignmentId,omitempty"`

	ResponseGroup []string `xml:"ResponseGroup,omitempty"`
}

type GetAssignmentResult struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GetAssignmentResult"`

	Request *Request `xml:"Request,omitempty"`

	Assignment *Assignment `xml:"Assignment,omitempty"`

	HIT *HIT `xml:"HIT,omitempty"`
}

type GetReviewableHITsRequest struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GetReviewableHITsRequest"`

	HITTypeId string `xml:"HITTypeId,omitempty"`

	Status *ReviewableHITStatus `xml:"Status,omitempty"`

	SortDirection *SortDirection `xml:"SortDirection,omitempty"`

	SortProperty *GetReviewableHITsSortProperty `xml:"SortProperty,omitempty"`

	PageNumber int32 `xml:"PageNumber,omitempty"`

	PageSize int32 `xml:"PageSize,omitempty"`

	ResponseGroup []string `xml:"ResponseGroup,omitempty"`
}

type GetReviewableHITsResult struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GetReviewableHITsResult"`

	Request *Request `xml:"Request,omitempty"`

	PageNumber int32 `xml:"PageNumber,omitempty"`

	NumResults int32 `xml:"NumResults,omitempty"`

	TotalNumResults int32 `xml:"TotalNumResults,omitempty"`

	HIT []*HIT `xml:"HIT,omitempty"`
}

type ReviewReport struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 ReviewReport"`

	PageNumber int32 `xml:"PageNumber,omitempty"`

	NumResults int32 `xml:"NumResults,omitempty"`

	TotalNumResults int32 `xml:"TotalNumResults,omitempty"`

	ReviewResult []*ReviewResultDetail `xml:"ReviewResult,omitempty"`

	ReviewAction []*ReviewActionDetail `xml:"ReviewAction,omitempty"`
}

type ReviewResultDetail struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 ReviewResultDetail"`

	ActionId string `xml:"ActionId,omitempty"`

	SubjectId string `xml:"SubjectId,omitempty"`

	SubjectType string `xml:"SubjectType,omitempty"`

	QuestionId string `xml:"QuestionId,omitempty"`

	Key string `xml:"Key,omitempty"`

	Value string `xml:"Value,omitempty"`
}

type ReviewActionDetail struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 ReviewActionDetail"`

	ActionId string `xml:"ActionId,omitempty"`

	ActionName string `xml:"ActionName,omitempty"`

	ObjectId string `xml:"ObjectId,omitempty"`

	ObjectType string `xml:"ObjectType,omitempty"`

	Status *ReviewActionStatus `xml:"Status,omitempty"`

	CompleteTime time.Time `xml:"CompleteTime,omitempty"`

	Result string `xml:"Result,omitempty"`

	ErrorCode string `xml:"ErrorCode,omitempty"`
}

type GetReviewResultsForHITRequest struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GetReviewResultsForHITRequest"`

	HITId string `xml:"HITId,omitempty"`

	PolicyLevel []*ReviewPolicyLevel `xml:"PolicyLevel,omitempty"`

	RetrieveActions bool `xml:"RetrieveActions,omitempty"`

	RetrieveResults bool `xml:"RetrieveResults,omitempty"`

	PageNumber int32 `xml:"PageNumber,omitempty"`

	PageSize int32 `xml:"PageSize,omitempty"`

	ResponseGroup []string `xml:"ResponseGroup,omitempty"`
}

type GetReviewResultsForHITResult struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GetReviewResultsForHITResult"`

	Request *Request `xml:"Request,omitempty"`

	HITId string `xml:"HITId,omitempty"`

	AssignmentReviewPolicy *ReviewPolicy `xml:"AssignmentReviewPolicy,omitempty"`

	HITReviewPolicy *ReviewPolicy `xml:"HITReviewPolicy,omitempty"`

	AssignmentReviewReport *ReviewReport `xml:"AssignmentReviewReport,omitempty"`

	HITReviewReport *ReviewReport `xml:"HITReviewReport,omitempty"`
}

type GetHITsForQualificationTypeRequest struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GetHITsForQualificationTypeRequest"`

	QualificationTypeId string `xml:"QualificationTypeId,omitempty"`

	PageNumber int32 `xml:"PageNumber,omitempty"`

	PageSize int32 `xml:"PageSize,omitempty"`

	ResponseGroup []string `xml:"ResponseGroup,omitempty"`
}

type GetHITsForQualificationTypeResult struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GetHITsForQualificationTypeResult"`

	Request *Request `xml:"Request,omitempty"`

	PageNumber int32 `xml:"PageNumber,omitempty"`

	NumResults int32 `xml:"NumResults,omitempty"`

	TotalNumResults int32 `xml:"TotalNumResults,omitempty"`

	HIT []*HIT `xml:"HIT,omitempty"`
}

type GetQualificationsForQualificationTypeRequest struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GetQualificationsForQualificationTypeRequest"`

	QualificationTypeId string `xml:"QualificationTypeId,omitempty"`

	Status *QualificationStatus `xml:"Status,omitempty"`

	PageNumber int32 `xml:"PageNumber,omitempty"`

	PageSize int32 `xml:"PageSize,omitempty"`

	ResponseGroup []string `xml:"ResponseGroup,omitempty"`
}

type GetQualificationsForQualificationTypeResult struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GetQualificationsForQualificationTypeResult"`

	Request *Request `xml:"Request,omitempty"`

	PageNumber int32 `xml:"PageNumber,omitempty"`

	NumResults int32 `xml:"NumResults,omitempty"`

	TotalNumResults int32 `xml:"TotalNumResults,omitempty"`

	Qualification []*Qualification `xml:"Qualification,omitempty"`
}

type SetHITAsReviewingRequest struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 SetHITAsReviewingRequest"`

	HITId string `xml:"HITId,omitempty"`

	Revert bool `xml:"Revert,omitempty"`

	ResponseGroup []string `xml:"ResponseGroup,omitempty"`
}

type SetHITAsReviewingResult struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 SetHITAsReviewingResult"`

	Request *Request `xml:"Request,omitempty"`
}

type ExtendHITRequest struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 ExtendHITRequest"`

	HITId string `xml:"HITId,omitempty"`

	MaxAssignmentsIncrement int32 `xml:"MaxAssignmentsIncrement,omitempty"`

	ExpirationIncrementInSeconds int64 `xml:"ExpirationIncrementInSeconds,omitempty"`

	UniqueRequestToken string `xml:"UniqueRequestToken,omitempty"`

	ResponseGroup []string `xml:"ResponseGroup,omitempty"`
}

type ExtendHITResult struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 ExtendHITResult"`

	Request *Request `xml:"Request,omitempty"`
}

type ForceExpireHITRequest struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 ForceExpireHITRequest"`

	HITId string `xml:"HITId,omitempty"`

	ResponseGroup []string `xml:"ResponseGroup,omitempty"`
}

type ForceExpireHITResult struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 ForceExpireHITResult"`

	Request *Request `xml:"Request,omitempty"`
}

type Assignment struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 Assignment"`

	Request *Request `xml:"Request,omitempty"`

	AssignmentId string `xml:"AssignmentId,omitempty"`

	WorkerId string `xml:"WorkerId,omitempty"`

	HITId string `xml:"HITId,omitempty"`

	AssignmentStatus *AssignmentStatus `xml:"AssignmentStatus,omitempty"`

	AutoApprovalTime time.Time `xml:"AutoApprovalTime,omitempty"`

	AcceptTime time.Time `xml:"AcceptTime,omitempty"`

	SubmitTime time.Time `xml:"SubmitTime,omitempty"`

	ApprovalTime time.Time `xml:"ApprovalTime,omitempty"`

	RejectionTime time.Time `xml:"RejectionTime,omitempty"`

	Deadline time.Time `xml:"Deadline,omitempty"`

	Answer string `xml:"Answer,omitempty"`

	RequesterFeedback string `xml:"RequesterFeedback,omitempty"`
}

type ApproveAssignmentRequest struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 ApproveAssignmentRequest"`

	AssignmentId string `xml:"AssignmentId,omitempty"`

	ResponseGroup []string `xml:"ResponseGroup,omitempty"`

	RequesterFeedback string `xml:"RequesterFeedback,omitempty"`
}

type ApproveAssignmentResult struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 ApproveAssignmentResult"`

	Request *Request `xml:"Request,omitempty"`
}

type RejectAssignmentRequest struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 RejectAssignmentRequest"`

	AssignmentId string `xml:"AssignmentId,omitempty"`

	ResponseGroup []string `xml:"ResponseGroup,omitempty"`

	RequesterFeedback string `xml:"RequesterFeedback,omitempty"`
}

type RejectAssignmentResult struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 RejectAssignmentResult"`

	Request *Request `xml:"Request,omitempty"`
}

type ApproveRejectedAssignmentRequest struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 ApproveRejectedAssignmentRequest"`

	AssignmentId string `xml:"AssignmentId,omitempty"`

	ResponseGroup []string `xml:"ResponseGroup,omitempty"`

	RequesterFeedback string `xml:"RequesterFeedback,omitempty"`
}

type ApproveRejectedAssignmentResult struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 ApproveRejectedAssignmentResult"`

	Request *Request `xml:"Request,omitempty"`
}

type GetAssignmentsForHITRequest struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GetAssignmentsForHITRequest"`

	HITId string `xml:"HITId,omitempty"`

	SortDirection *SortDirection `xml:"SortDirection,omitempty"`

	SortProperty *GetAssignmentsForHITSortProperty `xml:"SortProperty,omitempty"`

	PageNumber int32 `xml:"PageNumber,omitempty"`

	PageSize int32 `xml:"PageSize,omitempty"`

	AssignmentStatus []*AssignmentStatus `xml:"AssignmentStatus,omitempty"`

	ResponseGroup []string `xml:"ResponseGroup,omitempty"`
}

type GetAssignmentsForHITResult struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GetAssignmentsForHITResult"`

	Request *Request `xml:"Request,omitempty"`

	PageNumber int32 `xml:"PageNumber,omitempty"`

	NumResults int32 `xml:"NumResults,omitempty"`

	TotalNumResults int32 `xml:"TotalNumResults,omitempty"`

	Assignment []*Assignment `xml:"Assignment,omitempty"`
}

type GetFileUploadURLRequest struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GetFileUploadURLRequest"`

	AssignmentId string `xml:"AssignmentId,omitempty"`

	QuestionIdentifier string `xml:"QuestionIdentifier,omitempty"`
}

type GetFileUploadURLResult struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GetFileUploadURLResult"`

	Request *Request `xml:"Request,omitempty"`

	FileUploadURL string `xml:"FileUploadURL,omitempty"`
}

type SearchHITsRequest struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 SearchHITsRequest"`

	SortDirection *SortDirection `xml:"SortDirection,omitempty"`

	SortProperty *SearchHITsSortProperty `xml:"SortProperty,omitempty"`

	PageNumber int32 `xml:"PageNumber,omitempty"`

	PageSize int32 `xml:"PageSize,omitempty"`

	ResponseGroup []string `xml:"ResponseGroup,omitempty"`
}

type SearchHITsResult struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 SearchHITsResult"`

	Request *Request `xml:"Request,omitempty"`

	PageNumber int32 `xml:"PageNumber,omitempty"`

	NumResults int32 `xml:"NumResults,omitempty"`

	TotalNumResults int32 `xml:"TotalNumResults,omitempty"`

	HIT []*HIT `xml:"HIT,omitempty"`
}

type GrantBonusRequest struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GrantBonusRequest"`

	WorkerId string `xml:"WorkerId,omitempty"`

	BonusAmount *Price `xml:"BonusAmount,omitempty"`

	AssignmentId string `xml:"AssignmentId,omitempty"`

	Reason string `xml:"Reason,omitempty"`

	UniqueRequestToken string `xml:"UniqueRequestToken,omitempty"`

	ResponseGroup []string `xml:"ResponseGroup,omitempty"`
}

type GrantBonusResult struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GrantBonusResult"`

	Request *Request `xml:"Request,omitempty"`
}

type BonusPayment struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 BonusPayment"`

	Request *Request `xml:"Request,omitempty"`

	WorkerId string `xml:"WorkerId,omitempty"`

	BonusAmount *Price `xml:"BonusAmount,omitempty"`

	AssignmentId string `xml:"AssignmentId,omitempty"`

	Reason string `xml:"Reason,omitempty"`

	GrantTime time.Time `xml:"GrantTime,omitempty"`
}

type GetBonusPaymentsRequest struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GetBonusPaymentsRequest"`

	HITId string `xml:"HITId,omitempty"`

	AssignmentId string `xml:"AssignmentId,omitempty"`

	PageNumber int32 `xml:"PageNumber,omitempty"`

	PageSize int32 `xml:"PageSize,omitempty"`

	ResponseGroup []string `xml:"ResponseGroup,omitempty"`
}

type GetBonusPaymentsResult struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GetBonusPaymentsResult"`

	Request *Request `xml:"Request,omitempty"`

	NumResults int32 `xml:"NumResults,omitempty"`

	PageNumber int32 `xml:"PageNumber,omitempty"`

	TotalNumResults int32 `xml:"TotalNumResults,omitempty"`

	BonusPayment []*BonusPayment `xml:"BonusPayment,omitempty"`
}

type ChangeHITTypeOfHITRequest struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 ChangeHITTypeOfHITRequest"`

	HITId string `xml:"HITId,omitempty"`

	HITTypeId string `xml:"HITTypeId,omitempty"`

	ResponseGroup []string `xml:"ResponseGroup,omitempty"`
}

type ChangeHITTypeOfHITResult struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 ChangeHITTypeOfHITResult"`

	Request *Request `xml:"Request,omitempty"`
}

type QualificationType struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 QualificationType"`

	Request *Request `xml:"Request,omitempty"`

	QualificationTypeId string `xml:"QualificationTypeId,omitempty"`

	CreationTime time.Time `xml:"CreationTime,omitempty"`

	Name string `xml:"Name,omitempty"`

	Description string `xml:"Description,omitempty"`

	Keywords string `xml:"Keywords,omitempty"`

	QualificationTypeStatus *QualificationTypeStatus `xml:"QualificationTypeStatus,omitempty"`

	Test string `xml:"Test,omitempty"`

	TestDurationInSeconds int64 `xml:"TestDurationInSeconds,omitempty"`

	AnswerKey string `xml:"AnswerKey,omitempty"`

	RetryDelayInSeconds int64 `xml:"RetryDelayInSeconds,omitempty"`

	IsRequestable bool `xml:"IsRequestable,omitempty"`

	AutoGranted bool `xml:"AutoGranted,omitempty"`

	AutoGrantedValue int32 `xml:"AutoGrantedValue,omitempty"`
}

type QualificationRequest struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 QualificationRequest"`

	QualificationRequestId string `xml:"QualificationRequestId,omitempty"`

	QualificationTypeId string `xml:"QualificationTypeId,omitempty"`

	SubjectId string `xml:"SubjectId,omitempty"`

	Test string `xml:"Test,omitempty"`

	Answer string `xml:"Answer,omitempty"`

	SubmitTime time.Time `xml:"SubmitTime,omitempty"`
}

type CreateQualificationTypeRequest struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 CreateQualificationTypeRequest"`

	Name string `xml:"Name,omitempty"`

	Keywords string `xml:"Keywords,omitempty"`

	Description string `xml:"Description,omitempty"`

	QualificationTypeStatus *QualificationTypeStatus `xml:"QualificationTypeStatus,omitempty"`

	RetryDelayInSeconds int64 `xml:"RetryDelayInSeconds,omitempty"`

	Test string `xml:"Test,omitempty"`

	AnswerKey string `xml:"AnswerKey,omitempty"`

	TestDurationInSeconds int64 `xml:"TestDurationInSeconds,omitempty"`

	AutoGranted bool `xml:"AutoGranted,omitempty"`

	AutoGrantedValue int32 `xml:"AutoGrantedValue,omitempty"`

	ResponseGroup []string `xml:"ResponseGroup,omitempty"`
}

type DisposeQualificationTypeRequest struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 DisposeQualificationTypeRequest"`

	QualificationTypeId string `xml:"QualificationTypeId,omitempty"`
}

type DisposeQualificationTypeResult struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 DisposeQualificationTypeResult"`

	Request *Request `xml:"Request,omitempty"`
}

type GetQualificationRequestsRequest struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GetQualificationRequestsRequest"`

	QualificationTypeId string `xml:"QualificationTypeId,omitempty"`

	SortDirection *SortDirection `xml:"SortDirection,omitempty"`

	SortProperty *GetQualificationRequestsSortProperty `xml:"SortProperty,omitempty"`

	PageNumber int32 `xml:"PageNumber,omitempty"`

	PageSize int32 `xml:"PageSize,omitempty"`

	ResponseGroup []string `xml:"ResponseGroup,omitempty"`
}

type GetQualificationRequestsResult struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GetQualificationRequestsResult"`

	Request *Request `xml:"Request,omitempty"`

	NumResults int32 `xml:"NumResults,omitempty"`

	TotalNumResults int32 `xml:"TotalNumResults,omitempty"`

	PageNumber int32 `xml:"PageNumber,omitempty"`

	QualificationRequest []*QualificationRequest `xml:"QualificationRequest,omitempty"`
}

type RejectQualificationRequestRequest struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 RejectQualificationRequestRequest"`

	QualificationRequestId string `xml:"QualificationRequestId,omitempty"`

	Reason string `xml:"Reason,omitempty"`

	ResponseGroup []string `xml:"ResponseGroup,omitempty"`
}

type RejectQualificationRequestResult struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 RejectQualificationRequestResult"`

	Request *Request `xml:"Request,omitempty"`
}

type GrantQualificationRequest struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GrantQualificationRequest"`

	QualificationRequestId string `xml:"QualificationRequestId,omitempty"`

	IntegerValue int32 `xml:"IntegerValue,omitempty"`

	ResponseGroup []string `xml:"ResponseGroup,omitempty"`
}

type GrantQualificationResult struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GrantQualificationResult"`

	Request *Request `xml:"Request,omitempty"`
}

type AssignQualificationRequest struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 AssignQualificationRequest"`

	QualificationTypeId string `xml:"QualificationTypeId,omitempty"`

	WorkerId string `xml:"WorkerId,omitempty"`

	IntegerValue int32 `xml:"IntegerValue,omitempty"`

	SendNotification bool `xml:"SendNotification,omitempty"`

	ResponseGroup []string `xml:"ResponseGroup,omitempty"`
}

type AssignQualificationResult struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 AssignQualificationResult"`

	Request *Request `xml:"Request,omitempty"`
}

type RevokeQualificationRequest struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 RevokeQualificationRequest"`

	SubjectId string `xml:"SubjectId,omitempty"`

	QualificationTypeId string `xml:"QualificationTypeId,omitempty"`

	Reason string `xml:"Reason,omitempty"`

	ResponseGroup []string `xml:"ResponseGroup,omitempty"`
}

type RevokeQualificationResult struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 RevokeQualificationResult"`

	Request *Request `xml:"Request,omitempty"`
}

type UpdateQualificationScoreRequest struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 UpdateQualificationScoreRequest"`

	QualificationTypeId string `xml:"QualificationTypeId,omitempty"`

	SubjectId string `xml:"SubjectId,omitempty"`

	IntegerValue int32 `xml:"IntegerValue,omitempty"`

	ResponseGroup []string `xml:"ResponseGroup,omitempty"`
}

type UpdateQualificationScoreResult struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 UpdateQualificationScoreResult"`

	Request *Request `xml:"Request,omitempty"`
}

type GetQualificationTypeRequest struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GetQualificationTypeRequest"`

	QualificationTypeId string `xml:"QualificationTypeId,omitempty"`

	ResponseGroup []string `xml:"ResponseGroup,omitempty"`
}

type GetQualificationScoreRequest struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GetQualificationScoreRequest"`

	QualificationTypeId string `xml:"QualificationTypeId,omitempty"`

	SubjectId string `xml:"SubjectId,omitempty"`

	ResponseGroup []string `xml:"ResponseGroup,omitempty"`
}

type Qualification struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 Qualification"`

	Request *Request `xml:"Request,omitempty"`

	QualificationTypeId string `xml:"QualificationTypeId,omitempty"`

	SubjectId string `xml:"SubjectId,omitempty"`

	GrantTime time.Time `xml:"GrantTime,omitempty"`

	IntegerValue int32 `xml:"IntegerValue,omitempty"`

	LocaleValue *Locale `xml:"LocaleValue,omitempty"`

	Status *QualificationStatus `xml:"Status,omitempty"`
}

type SearchQualificationTypesRequest struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 SearchQualificationTypesRequest"`

	Query string `xml:"Query,omitempty"`

	MustBeRequestable bool `xml:"MustBeRequestable,omitempty"`

	MustBeOwnedByCaller bool `xml:"MustBeOwnedByCaller,omitempty"`

	SortDirection *SortDirection `xml:"SortDirection,omitempty"`

	SortProperty *SearchQualificationTypesSortProperty `xml:"SortProperty,omitempty"`

	PageNumber int32 `xml:"PageNumber,omitempty"`

	PageSize int32 `xml:"PageSize,omitempty"`

	ResponseGroup []string `xml:"ResponseGroup,omitempty"`
}

type SearchQualificationTypesResult struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 SearchQualificationTypesResult"`

	Request *Request `xml:"Request,omitempty"`

	NumResults int32 `xml:"NumResults,omitempty"`

	TotalNumResults int32 `xml:"TotalNumResults,omitempty"`

	PageNumber int32 `xml:"PageNumber,omitempty"`

	QualificationType []*QualificationType `xml:"QualificationType,omitempty"`
}

type UpdateQualificationTypeRequest struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 UpdateQualificationTypeRequest"`

	QualificationTypeId string `xml:"QualificationTypeId,omitempty"`

	Description string `xml:"Description,omitempty"`

	QualificationTypeStatus *QualificationTypeStatus `xml:"QualificationTypeStatus,omitempty"`

	Test string `xml:"Test,omitempty"`

	AnswerKey string `xml:"AnswerKey,omitempty"`

	TestDurationInSeconds int64 `xml:"TestDurationInSeconds,omitempty"`

	RetryDelayInSeconds int64 `xml:"RetryDelayInSeconds,omitempty"`

	AutoGranted bool `xml:"AutoGranted,omitempty"`

	AutoGrantedValue int32 `xml:"AutoGrantedValue,omitempty"`

	ResponseGroup []string `xml:"ResponseGroup,omitempty"`
}

type GetAccountBalanceRequest struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GetAccountBalanceRequest"`

	ResponseGroup []string `xml:"ResponseGroup,omitempty"`

	Unused string `xml:"Unused,omitempty"`
}

type GetAccountBalanceResult struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GetAccountBalanceResult"`

	Request *Request `xml:"Request,omitempty"`

	AvailableBalance *Price `xml:"AvailableBalance,omitempty"`

	OnHoldBalance *Price `xml:"OnHoldBalance,omitempty"`
}

type GetRequesterStatisticRequest struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GetRequesterStatisticRequest"`

	Statistic *RequesterStatistic `xml:"Statistic,omitempty"`

	TimePeriod *TimePeriod `xml:"TimePeriod,omitempty"`

	Count int32 `xml:"Count,omitempty"`

	ResponseGroup []string `xml:"ResponseGroup,omitempty"`
}

type GetRequesterWorkerStatisticRequest struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GetRequesterWorkerStatisticRequest"`

	Statistic *RequesterStatistic `xml:"Statistic,omitempty"`

	TimePeriod *TimePeriod `xml:"TimePeriod,omitempty"`

	WorkerId string `xml:"WorkerId,omitempty"`

	Count int32 `xml:"Count,omitempty"`

	ResponseGroup []string `xml:"ResponseGroup,omitempty"`
}

type NotifyWorkersFailureStatus struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 NotifyWorkersFailureStatus"`

	NotifyWorkersFailureCode *NotifyWorkersFailureCode `xml:"NotifyWorkersFailureCode,omitempty"`

	NotifyWorkersFailureMessage string `xml:"NotifyWorkersFailureMessage,omitempty"`

	WorkerId string `xml:"WorkerId,omitempty"`
}

type NotifyWorkersRequest struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 NotifyWorkersRequest"`

	Subject string `xml:"Subject,omitempty"`

	MessageText string `xml:"MessageText,omitempty"`

	WorkerId string `xml:"WorkerId,omitempty"`

	ResponseGroup []string `xml:"ResponseGroup,omitempty"`
}

type NotifyWorkersResult struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 NotifyWorkersResult"`

	Request *Request `xml:"Request,omitempty"`

	NotifyWorkersFailureStatus *NotifyWorkersFailureStatus `xml:"NotifyWorkersFailureStatus,omitempty"`
}

type GetStatisticResult struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GetStatisticResult"`

	Request *Request `xml:"Request,omitempty"`

	DataPoint []*DataPoint `xml:"DataPoint,omitempty"`

	Statistic *RequesterStatistic `xml:"Statistic,omitempty"`

	TimePeriod *TimePeriod `xml:"TimePeriod,omitempty"`

	WorkerId string `xml:"WorkerId,omitempty"`
}

type DataPoint struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 DataPoint"`

	Date time.Time `xml:"Date,omitempty"`

	LongValue int64 `xml:"LongValue,omitempty"`

	DoubleValue float64 `xml:"DoubleValue,omitempty"`
}

type GetBlockedWorkersRequest struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GetBlockedWorkersRequest"`

	PageNumber int32 `xml:"PageNumber,omitempty"`

	PageSize int32 `xml:"PageSize,omitempty"`

	ResponseGroup []string `xml:"ResponseGroup,omitempty"`
}

type GetBlockedWorkersResult struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 GetBlockedWorkersResult"`

	Request *Request `xml:"Request,omitempty"`

	PageNumber int32 `xml:"PageNumber,omitempty"`

	NumResults int32 `xml:"NumResults,omitempty"`

	TotalNumResults int32 `xml:"TotalNumResults,omitempty"`

	WorkerBlock []*WorkerBlock `xml:"WorkerBlock,omitempty"`
}

type WorkerBlock struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 WorkerBlock"`

	WorkerId string `xml:"WorkerId,omitempty"`

	Reason string `xml:"Reason,omitempty"`
}

type BlockWorkerRequest struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 BlockWorkerRequest"`

	WorkerId string `xml:"WorkerId,omitempty"`

	Reason string `xml:"Reason,omitempty"`

	ResponseGroup []string `xml:"ResponseGroup,omitempty"`
}

type BlockWorkerResult struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 BlockWorkerResult"`

	Request *Request `xml:"Request,omitempty"`
}

type UnblockWorkerRequest struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 UnblockWorkerRequest"`

	WorkerId string `xml:"WorkerId,omitempty"`

	Reason string `xml:"Reason,omitempty"`

	ResponseGroup []string `xml:"ResponseGroup,omitempty"`
}

type UnblockWorkerResult struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 UnblockWorkerResult"`

	Request *Request `xml:"Request,omitempty"`
}

type KeyValuePair struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 KeyValuePair"`

	Key string `xml:"Key,omitempty"`

	Value string `xml:"Value,omitempty"`
}

type HelpRequest struct {
	XMLName xml.Name `xml:"http://requester.mturk.amazonaws.com/doc/2012-03-25 HelpRequest"`

	About []string `xml:"About,omitempty"`

	HelpType struct {
	} `xml:"HelpType,omitempty"`

	ResponseGroup []string `xml:"ResponseGroup,omitempty"`
}
