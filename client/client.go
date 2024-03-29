// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type EnterpriseDingtalkGroupMember struct {
	// 代表资源名称的资源属性字段
	IsAdmin *bool `json:"IsAdmin,omitempty" xml:"IsAdmin,omitempty"`
	// 代表资源组的资源属性字段
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// 代表创建时间的资源属性字段
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s EnterpriseDingtalkGroupMember) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseDingtalkGroupMember) GoString() string {
	return s.String()
}

func (s *EnterpriseDingtalkGroupMember) SetIsAdmin(v bool) *EnterpriseDingtalkGroupMember {
	s.IsAdmin = &v
	return s
}

func (s *EnterpriseDingtalkGroupMember) SetMobile(v string) *EnterpriseDingtalkGroupMember {
	s.Mobile = &v
	return s
}

func (s *EnterpriseDingtalkGroupMember) SetName(v string) *EnterpriseDingtalkGroupMember {
	s.Name = &v
	return s
}

type CloseTaskOrderRequest struct {
	OrderId  *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s CloseTaskOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CloseTaskOrderRequest) GoString() string {
	return s.String()
}

func (s *CloseTaskOrderRequest) SetOrderId(v string) *CloseTaskOrderRequest {
	s.OrderId = &v
	return s
}

func (s *CloseTaskOrderRequest) SetUserName(v string) *CloseTaskOrderRequest {
	s.UserName = &v
	return s
}

type CloseTaskOrderResponseBody struct {
	// code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// data
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// msg
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CloseTaskOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CloseTaskOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CloseTaskOrderResponseBody) SetCode(v string) *CloseTaskOrderResponseBody {
	s.Code = &v
	return s
}

func (s *CloseTaskOrderResponseBody) SetData(v string) *CloseTaskOrderResponseBody {
	s.Data = &v
	return s
}

func (s *CloseTaskOrderResponseBody) SetMessage(v string) *CloseTaskOrderResponseBody {
	s.Message = &v
	return s
}

func (s *CloseTaskOrderResponseBody) SetRequestId(v string) *CloseTaskOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloseTaskOrderResponseBody) SetSuccess(v bool) *CloseTaskOrderResponseBody {
	s.Success = &v
	return s
}

type CloseTaskOrderResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CloseTaskOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CloseTaskOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CloseTaskOrderResponse) GoString() string {
	return s.String()
}

func (s *CloseTaskOrderResponse) SetHeaders(v map[string]*string) *CloseTaskOrderResponse {
	s.Headers = v
	return s
}

func (s *CloseTaskOrderResponse) SetStatusCode(v int32) *CloseTaskOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CloseTaskOrderResponse) SetBody(v *CloseTaskOrderResponseBody) *CloseTaskOrderResponse {
	s.Body = v
	return s
}

type CreateTaskOrderRequest struct {
	CustomerRealName     *string `json:"CustomerRealName,omitempty" xml:"CustomerRealName,omitempty"`
	CustomerUserId       *string `json:"CustomerUserId,omitempty" xml:"CustomerUserId,omitempty"`
	ImportantDescription *string `json:"ImportantDescription,omitempty" xml:"ImportantDescription,omitempty"`
	IsImportant          *string `json:"IsImportant,omitempty" xml:"IsImportant,omitempty"`
	OpenGroupId          *string `json:"OpenGroupId,omitempty" xml:"OpenGroupId,omitempty"`
	// productType
	ProductType     *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	ProductTypeName *string `json:"ProductTypeName,omitempty" xml:"ProductTypeName,omitempty"`
	// taskTitle
	TaskTitle *string `json:"TaskTitle,omitempty" xml:"TaskTitle,omitempty"`
}

func (s CreateTaskOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateTaskOrderRequest) SetCustomerRealName(v string) *CreateTaskOrderRequest {
	s.CustomerRealName = &v
	return s
}

func (s *CreateTaskOrderRequest) SetCustomerUserId(v string) *CreateTaskOrderRequest {
	s.CustomerUserId = &v
	return s
}

func (s *CreateTaskOrderRequest) SetImportantDescription(v string) *CreateTaskOrderRequest {
	s.ImportantDescription = &v
	return s
}

func (s *CreateTaskOrderRequest) SetIsImportant(v string) *CreateTaskOrderRequest {
	s.IsImportant = &v
	return s
}

func (s *CreateTaskOrderRequest) SetOpenGroupId(v string) *CreateTaskOrderRequest {
	s.OpenGroupId = &v
	return s
}

func (s *CreateTaskOrderRequest) SetProductType(v string) *CreateTaskOrderRequest {
	s.ProductType = &v
	return s
}

func (s *CreateTaskOrderRequest) SetProductTypeName(v string) *CreateTaskOrderRequest {
	s.ProductTypeName = &v
	return s
}

func (s *CreateTaskOrderRequest) SetTaskTitle(v string) *CreateTaskOrderRequest {
	s.TaskTitle = &v
	return s
}

type CreateTaskOrderResponseBody struct {
	// code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// msg
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateTaskOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTaskOrderResponseBody) SetCode(v string) *CreateTaskOrderResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTaskOrderResponseBody) SetData(v string) *CreateTaskOrderResponseBody {
	s.Data = &v
	return s
}

func (s *CreateTaskOrderResponseBody) SetMessage(v string) *CreateTaskOrderResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTaskOrderResponseBody) SetRequestId(v string) *CreateTaskOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTaskOrderResponseBody) SetSuccess(v bool) *CreateTaskOrderResponseBody {
	s.Success = &v
	return s
}

type CreateTaskOrderResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateTaskOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTaskOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateTaskOrderResponse) SetHeaders(v map[string]*string) *CreateTaskOrderResponse {
	s.Headers = v
	return s
}

func (s *CreateTaskOrderResponse) SetStatusCode(v int32) *CreateTaskOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTaskOrderResponse) SetBody(v *CreateTaskOrderResponseBody) *CreateTaskOrderResponse {
	s.Body = v
	return s
}

type CreateTaskOrderByEventReportRequest struct {
	Business              *string                                       `json:"Business,omitempty" xml:"Business,omitempty"`
	CreateRealName        *string                                       `json:"CreateRealName,omitempty" xml:"CreateRealName,omitempty"`
	CreateUserId          *string                                       `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	EventBody             *CreateTaskOrderByEventReportRequestEventBody `json:"EventBody,omitempty" xml:"EventBody,omitempty" type:"Struct"`
	Extinfo               []*CreateTaskOrderByEventReportRequestExtinfo `json:"Extinfo,omitempty" xml:"Extinfo,omitempty" type:"Repeated"`
	ImportantDesc         *string                                       `json:"ImportantDesc,omitempty" xml:"ImportantDesc,omitempty"`
	JoinChildGroupUserIds *string                                       `json:"JoinChildGroupUserIds,omitempty" xml:"JoinChildGroupUserIds,omitempty"`
	MonitorCongregation   *string                                       `json:"MonitorCongregation,omitempty" xml:"MonitorCongregation,omitempty"`
	OpenGroupId           *string                                       `json:"OpenGroupId,omitempty" xml:"OpenGroupId,omitempty"`
	ProductType           *string                                       `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s CreateTaskOrderByEventReportRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskOrderByEventReportRequest) GoString() string {
	return s.String()
}

func (s *CreateTaskOrderByEventReportRequest) SetBusiness(v string) *CreateTaskOrderByEventReportRequest {
	s.Business = &v
	return s
}

func (s *CreateTaskOrderByEventReportRequest) SetCreateRealName(v string) *CreateTaskOrderByEventReportRequest {
	s.CreateRealName = &v
	return s
}

func (s *CreateTaskOrderByEventReportRequest) SetCreateUserId(v string) *CreateTaskOrderByEventReportRequest {
	s.CreateUserId = &v
	return s
}

func (s *CreateTaskOrderByEventReportRequest) SetEventBody(v *CreateTaskOrderByEventReportRequestEventBody) *CreateTaskOrderByEventReportRequest {
	s.EventBody = v
	return s
}

func (s *CreateTaskOrderByEventReportRequest) SetExtinfo(v []*CreateTaskOrderByEventReportRequestExtinfo) *CreateTaskOrderByEventReportRequest {
	s.Extinfo = v
	return s
}

func (s *CreateTaskOrderByEventReportRequest) SetImportantDesc(v string) *CreateTaskOrderByEventReportRequest {
	s.ImportantDesc = &v
	return s
}

func (s *CreateTaskOrderByEventReportRequest) SetJoinChildGroupUserIds(v string) *CreateTaskOrderByEventReportRequest {
	s.JoinChildGroupUserIds = &v
	return s
}

func (s *CreateTaskOrderByEventReportRequest) SetMonitorCongregation(v string) *CreateTaskOrderByEventReportRequest {
	s.MonitorCongregation = &v
	return s
}

func (s *CreateTaskOrderByEventReportRequest) SetOpenGroupId(v string) *CreateTaskOrderByEventReportRequest {
	s.OpenGroupId = &v
	return s
}

func (s *CreateTaskOrderByEventReportRequest) SetProductType(v string) *CreateTaskOrderByEventReportRequest {
	s.ProductType = &v
	return s
}

type CreateTaskOrderByEventReportRequestEventBody struct {
	EventDesc     *string                                                    `json:"EventDesc,omitempty" xml:"EventDesc,omitempty"`
	EventId       *string                                                    `json:"EventId,omitempty" xml:"EventId,omitempty"`
	EventLevel    *string                                                    `json:"EventLevel,omitempty" xml:"EventLevel,omitempty"`
	EventLocation *CreateTaskOrderByEventReportRequestEventBodyEventLocation `json:"EventLocation,omitempty" xml:"EventLocation,omitempty" type:"Struct"`
	EventTime     *string                                                    `json:"EventTime,omitempty" xml:"EventTime,omitempty"`
}

func (s CreateTaskOrderByEventReportRequestEventBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskOrderByEventReportRequestEventBody) GoString() string {
	return s.String()
}

func (s *CreateTaskOrderByEventReportRequestEventBody) SetEventDesc(v string) *CreateTaskOrderByEventReportRequestEventBody {
	s.EventDesc = &v
	return s
}

func (s *CreateTaskOrderByEventReportRequestEventBody) SetEventId(v string) *CreateTaskOrderByEventReportRequestEventBody {
	s.EventId = &v
	return s
}

func (s *CreateTaskOrderByEventReportRequestEventBody) SetEventLevel(v string) *CreateTaskOrderByEventReportRequestEventBody {
	s.EventLevel = &v
	return s
}

func (s *CreateTaskOrderByEventReportRequestEventBody) SetEventLocation(v *CreateTaskOrderByEventReportRequestEventBodyEventLocation) *CreateTaskOrderByEventReportRequestEventBody {
	s.EventLocation = v
	return s
}

func (s *CreateTaskOrderByEventReportRequestEventBody) SetEventTime(v string) *CreateTaskOrderByEventReportRequestEventBody {
	s.EventTime = &v
	return s
}

type CreateTaskOrderByEventReportRequestEventBodyEventLocation struct {
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s CreateTaskOrderByEventReportRequestEventBodyEventLocation) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskOrderByEventReportRequestEventBodyEventLocation) GoString() string {
	return s.String()
}

func (s *CreateTaskOrderByEventReportRequestEventBodyEventLocation) SetDomain(v string) *CreateTaskOrderByEventReportRequestEventBodyEventLocation {
	s.Domain = &v
	return s
}

type CreateTaskOrderByEventReportRequestExtinfo struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateTaskOrderByEventReportRequestExtinfo) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskOrderByEventReportRequestExtinfo) GoString() string {
	return s.String()
}

func (s *CreateTaskOrderByEventReportRequestExtinfo) SetName(v string) *CreateTaskOrderByEventReportRequestExtinfo {
	s.Name = &v
	return s
}

func (s *CreateTaskOrderByEventReportRequestExtinfo) SetValue(v string) *CreateTaskOrderByEventReportRequestExtinfo {
	s.Value = &v
	return s
}

type CreateTaskOrderByEventReportShrinkRequest struct {
	Business              *string `json:"Business,omitempty" xml:"Business,omitempty"`
	CreateRealName        *string `json:"CreateRealName,omitempty" xml:"CreateRealName,omitempty"`
	CreateUserId          *string `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	EventBodyShrink       *string `json:"EventBody,omitempty" xml:"EventBody,omitempty"`
	ExtinfoShrink         *string `json:"Extinfo,omitempty" xml:"Extinfo,omitempty"`
	ImportantDesc         *string `json:"ImportantDesc,omitempty" xml:"ImportantDesc,omitempty"`
	JoinChildGroupUserIds *string `json:"JoinChildGroupUserIds,omitempty" xml:"JoinChildGroupUserIds,omitempty"`
	MonitorCongregation   *string `json:"MonitorCongregation,omitempty" xml:"MonitorCongregation,omitempty"`
	OpenGroupId           *string `json:"OpenGroupId,omitempty" xml:"OpenGroupId,omitempty"`
	ProductType           *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s CreateTaskOrderByEventReportShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskOrderByEventReportShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateTaskOrderByEventReportShrinkRequest) SetBusiness(v string) *CreateTaskOrderByEventReportShrinkRequest {
	s.Business = &v
	return s
}

func (s *CreateTaskOrderByEventReportShrinkRequest) SetCreateRealName(v string) *CreateTaskOrderByEventReportShrinkRequest {
	s.CreateRealName = &v
	return s
}

func (s *CreateTaskOrderByEventReportShrinkRequest) SetCreateUserId(v string) *CreateTaskOrderByEventReportShrinkRequest {
	s.CreateUserId = &v
	return s
}

func (s *CreateTaskOrderByEventReportShrinkRequest) SetEventBodyShrink(v string) *CreateTaskOrderByEventReportShrinkRequest {
	s.EventBodyShrink = &v
	return s
}

func (s *CreateTaskOrderByEventReportShrinkRequest) SetExtinfoShrink(v string) *CreateTaskOrderByEventReportShrinkRequest {
	s.ExtinfoShrink = &v
	return s
}

func (s *CreateTaskOrderByEventReportShrinkRequest) SetImportantDesc(v string) *CreateTaskOrderByEventReportShrinkRequest {
	s.ImportantDesc = &v
	return s
}

func (s *CreateTaskOrderByEventReportShrinkRequest) SetJoinChildGroupUserIds(v string) *CreateTaskOrderByEventReportShrinkRequest {
	s.JoinChildGroupUserIds = &v
	return s
}

func (s *CreateTaskOrderByEventReportShrinkRequest) SetMonitorCongregation(v string) *CreateTaskOrderByEventReportShrinkRequest {
	s.MonitorCongregation = &v
	return s
}

func (s *CreateTaskOrderByEventReportShrinkRequest) SetOpenGroupId(v string) *CreateTaskOrderByEventReportShrinkRequest {
	s.OpenGroupId = &v
	return s
}

func (s *CreateTaskOrderByEventReportShrinkRequest) SetProductType(v string) *CreateTaskOrderByEventReportShrinkRequest {
	s.ProductType = &v
	return s
}

type CreateTaskOrderByEventReportResponseBody struct {
	// code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// msg
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateTaskOrderByEventReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskOrderByEventReportResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTaskOrderByEventReportResponseBody) SetCode(v string) *CreateTaskOrderByEventReportResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTaskOrderByEventReportResponseBody) SetData(v string) *CreateTaskOrderByEventReportResponseBody {
	s.Data = &v
	return s
}

func (s *CreateTaskOrderByEventReportResponseBody) SetMessage(v string) *CreateTaskOrderByEventReportResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTaskOrderByEventReportResponseBody) SetRequestId(v string) *CreateTaskOrderByEventReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTaskOrderByEventReportResponseBody) SetSuccess(v bool) *CreateTaskOrderByEventReportResponseBody {
	s.Success = &v
	return s
}

type CreateTaskOrderByEventReportResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateTaskOrderByEventReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTaskOrderByEventReportResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskOrderByEventReportResponse) GoString() string {
	return s.String()
}

func (s *CreateTaskOrderByEventReportResponse) SetHeaders(v map[string]*string) *CreateTaskOrderByEventReportResponse {
	s.Headers = v
	return s
}

func (s *CreateTaskOrderByEventReportResponse) SetStatusCode(v int32) *CreateTaskOrderByEventReportResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTaskOrderByEventReportResponse) SetBody(v *CreateTaskOrderByEventReportResponseBody) *CreateTaskOrderByEventReportResponse {
	s.Body = v
	return s
}

type DeleteEnterpriseDingtalkGroupCustomerMemberRequest struct {
	Mobiles     []*string `json:"Mobiles,omitempty" xml:"Mobiles,omitempty" type:"Repeated"`
	OpenGroupId *string   `json:"OpenGroupId,omitempty" xml:"OpenGroupId,omitempty"`
}

func (s DeleteEnterpriseDingtalkGroupCustomerMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteEnterpriseDingtalkGroupCustomerMemberRequest) GoString() string {
	return s.String()
}

func (s *DeleteEnterpriseDingtalkGroupCustomerMemberRequest) SetMobiles(v []*string) *DeleteEnterpriseDingtalkGroupCustomerMemberRequest {
	s.Mobiles = v
	return s
}

func (s *DeleteEnterpriseDingtalkGroupCustomerMemberRequest) SetOpenGroupId(v string) *DeleteEnterpriseDingtalkGroupCustomerMemberRequest {
	s.OpenGroupId = &v
	return s
}

type DeleteEnterpriseDingtalkGroupCustomerMemberShrinkRequest struct {
	MobilesShrink *string `json:"Mobiles,omitempty" xml:"Mobiles,omitempty"`
	OpenGroupId   *string `json:"OpenGroupId,omitempty" xml:"OpenGroupId,omitempty"`
}

func (s DeleteEnterpriseDingtalkGroupCustomerMemberShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteEnterpriseDingtalkGroupCustomerMemberShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteEnterpriseDingtalkGroupCustomerMemberShrinkRequest) SetMobilesShrink(v string) *DeleteEnterpriseDingtalkGroupCustomerMemberShrinkRequest {
	s.MobilesShrink = &v
	return s
}

func (s *DeleteEnterpriseDingtalkGroupCustomerMemberShrinkRequest) SetOpenGroupId(v string) *DeleteEnterpriseDingtalkGroupCustomerMemberShrinkRequest {
	s.OpenGroupId = &v
	return s
}

type DeleteEnterpriseDingtalkGroupCustomerMemberResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteEnterpriseDingtalkGroupCustomerMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteEnterpriseDingtalkGroupCustomerMemberResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEnterpriseDingtalkGroupCustomerMemberResponseBody) SetCode(v string) *DeleteEnterpriseDingtalkGroupCustomerMemberResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteEnterpriseDingtalkGroupCustomerMemberResponseBody) SetMessage(v string) *DeleteEnterpriseDingtalkGroupCustomerMemberResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteEnterpriseDingtalkGroupCustomerMemberResponseBody) SetRequestId(v string) *DeleteEnterpriseDingtalkGroupCustomerMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEnterpriseDingtalkGroupCustomerMemberResponseBody) SetSuccess(v bool) *DeleteEnterpriseDingtalkGroupCustomerMemberResponseBody {
	s.Success = &v
	return s
}

type DeleteEnterpriseDingtalkGroupCustomerMemberResponse struct {
	Headers    map[string]*string                                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteEnterpriseDingtalkGroupCustomerMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteEnterpriseDingtalkGroupCustomerMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteEnterpriseDingtalkGroupCustomerMemberResponse) GoString() string {
	return s.String()
}

func (s *DeleteEnterpriseDingtalkGroupCustomerMemberResponse) SetHeaders(v map[string]*string) *DeleteEnterpriseDingtalkGroupCustomerMemberResponse {
	s.Headers = v
	return s
}

func (s *DeleteEnterpriseDingtalkGroupCustomerMemberResponse) SetStatusCode(v int32) *DeleteEnterpriseDingtalkGroupCustomerMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEnterpriseDingtalkGroupCustomerMemberResponse) SetBody(v *DeleteEnterpriseDingtalkGroupCustomerMemberResponseBody) *DeleteEnterpriseDingtalkGroupCustomerMemberResponse {
	s.Body = v
	return s
}

type GetEnterpriseDingtalkGroupRequest struct {
	OpenGroupId *string `json:"OpenGroupId,omitempty" xml:"OpenGroupId,omitempty"`
}

func (s GetEnterpriseDingtalkGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s GetEnterpriseDingtalkGroupRequest) GoString() string {
	return s.String()
}

func (s *GetEnterpriseDingtalkGroupRequest) SetOpenGroupId(v string) *GetEnterpriseDingtalkGroupRequest {
	s.OpenGroupId = &v
	return s
}

type GetEnterpriseDingtalkGroupResponseBody struct {
	Code      *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetEnterpriseDingtalkGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetEnterpriseDingtalkGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEnterpriseDingtalkGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetEnterpriseDingtalkGroupResponseBody) SetCode(v string) *GetEnterpriseDingtalkGroupResponseBody {
	s.Code = &v
	return s
}

func (s *GetEnterpriseDingtalkGroupResponseBody) SetData(v *GetEnterpriseDingtalkGroupResponseBodyData) *GetEnterpriseDingtalkGroupResponseBody {
	s.Data = v
	return s
}

func (s *GetEnterpriseDingtalkGroupResponseBody) SetMessage(v string) *GetEnterpriseDingtalkGroupResponseBody {
	s.Message = &v
	return s
}

func (s *GetEnterpriseDingtalkGroupResponseBody) SetRequestId(v string) *GetEnterpriseDingtalkGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEnterpriseDingtalkGroupResponseBody) SetSuccess(v bool) *GetEnterpriseDingtalkGroupResponseBody {
	s.Success = &v
	return s
}

type GetEnterpriseDingtalkGroupResponseBodyData struct {
	GroupName   *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	OpenGroupId *string `json:"OpenGroupId,omitempty" xml:"OpenGroupId,omitempty"`
}

func (s GetEnterpriseDingtalkGroupResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetEnterpriseDingtalkGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetEnterpriseDingtalkGroupResponseBodyData) SetGroupName(v string) *GetEnterpriseDingtalkGroupResponseBodyData {
	s.GroupName = &v
	return s
}

func (s *GetEnterpriseDingtalkGroupResponseBodyData) SetOpenGroupId(v string) *GetEnterpriseDingtalkGroupResponseBodyData {
	s.OpenGroupId = &v
	return s
}

type GetEnterpriseDingtalkGroupResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetEnterpriseDingtalkGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetEnterpriseDingtalkGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEnterpriseDingtalkGroupResponse) GoString() string {
	return s.String()
}

func (s *GetEnterpriseDingtalkGroupResponse) SetHeaders(v map[string]*string) *GetEnterpriseDingtalkGroupResponse {
	s.Headers = v
	return s
}

func (s *GetEnterpriseDingtalkGroupResponse) SetStatusCode(v int32) *GetEnterpriseDingtalkGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEnterpriseDingtalkGroupResponse) SetBody(v *GetEnterpriseDingtalkGroupResponseBody) *GetEnterpriseDingtalkGroupResponse {
	s.Body = v
	return s
}

type GetEnterpriseDingtalkGroupCustomerMemberRequest struct {
	Mobile      *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	OpenGroupId *string `json:"OpenGroupId,omitempty" xml:"OpenGroupId,omitempty"`
}

func (s GetEnterpriseDingtalkGroupCustomerMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s GetEnterpriseDingtalkGroupCustomerMemberRequest) GoString() string {
	return s.String()
}

func (s *GetEnterpriseDingtalkGroupCustomerMemberRequest) SetMobile(v string) *GetEnterpriseDingtalkGroupCustomerMemberRequest {
	s.Mobile = &v
	return s
}

func (s *GetEnterpriseDingtalkGroupCustomerMemberRequest) SetOpenGroupId(v string) *GetEnterpriseDingtalkGroupCustomerMemberRequest {
	s.OpenGroupId = &v
	return s
}

type GetEnterpriseDingtalkGroupCustomerMemberResponseBody struct {
	Code      *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *EnterpriseDingtalkGroupMember `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetEnterpriseDingtalkGroupCustomerMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEnterpriseDingtalkGroupCustomerMemberResponseBody) GoString() string {
	return s.String()
}

func (s *GetEnterpriseDingtalkGroupCustomerMemberResponseBody) SetCode(v string) *GetEnterpriseDingtalkGroupCustomerMemberResponseBody {
	s.Code = &v
	return s
}

func (s *GetEnterpriseDingtalkGroupCustomerMemberResponseBody) SetData(v *EnterpriseDingtalkGroupMember) *GetEnterpriseDingtalkGroupCustomerMemberResponseBody {
	s.Data = v
	return s
}

func (s *GetEnterpriseDingtalkGroupCustomerMemberResponseBody) SetMessage(v string) *GetEnterpriseDingtalkGroupCustomerMemberResponseBody {
	s.Message = &v
	return s
}

func (s *GetEnterpriseDingtalkGroupCustomerMemberResponseBody) SetRequestId(v string) *GetEnterpriseDingtalkGroupCustomerMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEnterpriseDingtalkGroupCustomerMemberResponseBody) SetSuccess(v bool) *GetEnterpriseDingtalkGroupCustomerMemberResponseBody {
	s.Success = &v
	return s
}

type GetEnterpriseDingtalkGroupCustomerMemberResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetEnterpriseDingtalkGroupCustomerMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetEnterpriseDingtalkGroupCustomerMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEnterpriseDingtalkGroupCustomerMemberResponse) GoString() string {
	return s.String()
}

func (s *GetEnterpriseDingtalkGroupCustomerMemberResponse) SetHeaders(v map[string]*string) *GetEnterpriseDingtalkGroupCustomerMemberResponse {
	s.Headers = v
	return s
}

func (s *GetEnterpriseDingtalkGroupCustomerMemberResponse) SetStatusCode(v int32) *GetEnterpriseDingtalkGroupCustomerMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEnterpriseDingtalkGroupCustomerMemberResponse) SetBody(v *GetEnterpriseDingtalkGroupCustomerMemberResponseBody) *GetEnterpriseDingtalkGroupCustomerMemberResponse {
	s.Body = v
	return s
}

type ListDdTaskOrderRequest struct {
	// callerParentId
	CallerParentId *int64 `json:"CallerParentId,omitempty" xml:"CallerParentId,omitempty"`
	// callerType
	CallerType *string `json:"CallerType,omitempty" xml:"CallerType,omitempty"`
	// callerUid
	CallerUid *int64 `json:"CallerUid,omitempty" xml:"CallerUid,omitempty"`
	// openGroupId
	OpenGroupId *string `json:"OpenGroupId,omitempty" xml:"OpenGroupId,omitempty"`
	// orderId
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// taskStatus
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
}

func (s ListDdTaskOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDdTaskOrderRequest) GoString() string {
	return s.String()
}

func (s *ListDdTaskOrderRequest) SetCallerParentId(v int64) *ListDdTaskOrderRequest {
	s.CallerParentId = &v
	return s
}

func (s *ListDdTaskOrderRequest) SetCallerType(v string) *ListDdTaskOrderRequest {
	s.CallerType = &v
	return s
}

func (s *ListDdTaskOrderRequest) SetCallerUid(v int64) *ListDdTaskOrderRequest {
	s.CallerUid = &v
	return s
}

func (s *ListDdTaskOrderRequest) SetOpenGroupId(v string) *ListDdTaskOrderRequest {
	s.OpenGroupId = &v
	return s
}

func (s *ListDdTaskOrderRequest) SetOrderId(v string) *ListDdTaskOrderRequest {
	s.OrderId = &v
	return s
}

func (s *ListDdTaskOrderRequest) SetRequestId(v string) *ListDdTaskOrderRequest {
	s.RequestId = &v
	return s
}

func (s *ListDdTaskOrderRequest) SetTaskStatus(v string) *ListDdTaskOrderRequest {
	s.TaskStatus = &v
	return s
}

type ListDdTaskOrderResponseBody struct {
	// code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// data
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// msg
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDdTaskOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDdTaskOrderResponseBody) GoString() string {
	return s.String()
}

func (s *ListDdTaskOrderResponseBody) SetCode(v string) *ListDdTaskOrderResponseBody {
	s.Code = &v
	return s
}

func (s *ListDdTaskOrderResponseBody) SetData(v string) *ListDdTaskOrderResponseBody {
	s.Data = &v
	return s
}

func (s *ListDdTaskOrderResponseBody) SetMessage(v string) *ListDdTaskOrderResponseBody {
	s.Message = &v
	return s
}

func (s *ListDdTaskOrderResponseBody) SetRequestId(v string) *ListDdTaskOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDdTaskOrderResponseBody) SetSuccess(v bool) *ListDdTaskOrderResponseBody {
	s.Success = &v
	return s
}

type ListDdTaskOrderResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDdTaskOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDdTaskOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDdTaskOrderResponse) GoString() string {
	return s.String()
}

func (s *ListDdTaskOrderResponse) SetHeaders(v map[string]*string) *ListDdTaskOrderResponse {
	s.Headers = v
	return s
}

func (s *ListDdTaskOrderResponse) SetStatusCode(v int32) *ListDdTaskOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDdTaskOrderResponse) SetBody(v *ListDdTaskOrderResponseBody) *ListDdTaskOrderResponse {
	s.Body = v
	return s
}

type ListEnterpriseDingtalkGroupCustomerMembersRequest struct {
	OpenGroupId *string `json:"OpenGroupId,omitempty" xml:"OpenGroupId,omitempty"`
}

func (s ListEnterpriseDingtalkGroupCustomerMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEnterpriseDingtalkGroupCustomerMembersRequest) GoString() string {
	return s.String()
}

func (s *ListEnterpriseDingtalkGroupCustomerMembersRequest) SetOpenGroupId(v string) *ListEnterpriseDingtalkGroupCustomerMembersRequest {
	s.OpenGroupId = &v
	return s
}

type ListEnterpriseDingtalkGroupCustomerMembersResponseBody struct {
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*EnterpriseDingtalkGroupMember `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListEnterpriseDingtalkGroupCustomerMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEnterpriseDingtalkGroupCustomerMembersResponseBody) GoString() string {
	return s.String()
}

func (s *ListEnterpriseDingtalkGroupCustomerMembersResponseBody) SetCode(v string) *ListEnterpriseDingtalkGroupCustomerMembersResponseBody {
	s.Code = &v
	return s
}

func (s *ListEnterpriseDingtalkGroupCustomerMembersResponseBody) SetData(v []*EnterpriseDingtalkGroupMember) *ListEnterpriseDingtalkGroupCustomerMembersResponseBody {
	s.Data = v
	return s
}

func (s *ListEnterpriseDingtalkGroupCustomerMembersResponseBody) SetMessage(v string) *ListEnterpriseDingtalkGroupCustomerMembersResponseBody {
	s.Message = &v
	return s
}

func (s *ListEnterpriseDingtalkGroupCustomerMembersResponseBody) SetRequestId(v string) *ListEnterpriseDingtalkGroupCustomerMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEnterpriseDingtalkGroupCustomerMembersResponseBody) SetSuccess(v bool) *ListEnterpriseDingtalkGroupCustomerMembersResponseBody {
	s.Success = &v
	return s
}

type ListEnterpriseDingtalkGroupCustomerMembersResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListEnterpriseDingtalkGroupCustomerMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListEnterpriseDingtalkGroupCustomerMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEnterpriseDingtalkGroupCustomerMembersResponse) GoString() string {
	return s.String()
}

func (s *ListEnterpriseDingtalkGroupCustomerMembersResponse) SetHeaders(v map[string]*string) *ListEnterpriseDingtalkGroupCustomerMembersResponse {
	s.Headers = v
	return s
}

func (s *ListEnterpriseDingtalkGroupCustomerMembersResponse) SetStatusCode(v int32) *ListEnterpriseDingtalkGroupCustomerMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEnterpriseDingtalkGroupCustomerMembersResponse) SetBody(v *ListEnterpriseDingtalkGroupCustomerMembersResponseBody) *ListEnterpriseDingtalkGroupCustomerMembersResponse {
	s.Body = v
	return s
}

type ListEnterpriseDingtalkGroupsResponseBody struct {
	Code      *string                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*ListEnterpriseDingtalkGroupsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListEnterpriseDingtalkGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEnterpriseDingtalkGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEnterpriseDingtalkGroupsResponseBody) SetCode(v string) *ListEnterpriseDingtalkGroupsResponseBody {
	s.Code = &v
	return s
}

func (s *ListEnterpriseDingtalkGroupsResponseBody) SetData(v []*ListEnterpriseDingtalkGroupsResponseBodyData) *ListEnterpriseDingtalkGroupsResponseBody {
	s.Data = v
	return s
}

func (s *ListEnterpriseDingtalkGroupsResponseBody) SetMessage(v string) *ListEnterpriseDingtalkGroupsResponseBody {
	s.Message = &v
	return s
}

func (s *ListEnterpriseDingtalkGroupsResponseBody) SetRequestId(v string) *ListEnterpriseDingtalkGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEnterpriseDingtalkGroupsResponseBody) SetSuccess(v bool) *ListEnterpriseDingtalkGroupsResponseBody {
	s.Success = &v
	return s
}

type ListEnterpriseDingtalkGroupsResponseBodyData struct {
	GroupName   *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	OpenGroupId *string `json:"OpenGroupId,omitempty" xml:"OpenGroupId,omitempty"`
}

func (s ListEnterpriseDingtalkGroupsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListEnterpriseDingtalkGroupsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListEnterpriseDingtalkGroupsResponseBodyData) SetGroupName(v string) *ListEnterpriseDingtalkGroupsResponseBodyData {
	s.GroupName = &v
	return s
}

func (s *ListEnterpriseDingtalkGroupsResponseBodyData) SetOpenGroupId(v string) *ListEnterpriseDingtalkGroupsResponseBodyData {
	s.OpenGroupId = &v
	return s
}

type ListEnterpriseDingtalkGroupsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListEnterpriseDingtalkGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListEnterpriseDingtalkGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEnterpriseDingtalkGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListEnterpriseDingtalkGroupsResponse) SetHeaders(v map[string]*string) *ListEnterpriseDingtalkGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListEnterpriseDingtalkGroupsResponse) SetStatusCode(v int32) *ListEnterpriseDingtalkGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEnterpriseDingtalkGroupsResponse) SetBody(v *ListEnterpriseDingtalkGroupsResponseBody) *ListEnterpriseDingtalkGroupsResponse {
	s.Body = v
	return s
}

type ListProductByGroupRequest struct {
	OpenGroupId *string `json:"OpenGroupId,omitempty" xml:"OpenGroupId,omitempty"`
}

func (s ListProductByGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProductByGroupRequest) GoString() string {
	return s.String()
}

func (s *ListProductByGroupRequest) SetOpenGroupId(v string) *ListProductByGroupRequest {
	s.OpenGroupId = &v
	return s
}

type ListProductByGroupResponseBody struct {
	// code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// data
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// msg
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListProductByGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProductByGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListProductByGroupResponseBody) SetCode(v string) *ListProductByGroupResponseBody {
	s.Code = &v
	return s
}

func (s *ListProductByGroupResponseBody) SetData(v string) *ListProductByGroupResponseBody {
	s.Data = &v
	return s
}

func (s *ListProductByGroupResponseBody) SetMessage(v string) *ListProductByGroupResponseBody {
	s.Message = &v
	return s
}

func (s *ListProductByGroupResponseBody) SetRequestId(v string) *ListProductByGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProductByGroupResponseBody) SetSuccess(v bool) *ListProductByGroupResponseBody {
	s.Success = &v
	return s
}

type ListProductByGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListProductByGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListProductByGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProductByGroupResponse) GoString() string {
	return s.String()
}

func (s *ListProductByGroupResponse) SetHeaders(v map[string]*string) *ListProductByGroupResponse {
	s.Headers = v
	return s
}

func (s *ListProductByGroupResponse) SetStatusCode(v int32) *ListProductByGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProductByGroupResponse) SetBody(v *ListProductByGroupResponseBody) *ListProductByGroupResponse {
	s.Body = v
	return s
}

type QueryTaskInfoRequest struct {
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s QueryTaskInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryTaskInfoRequest) SetOrderId(v string) *QueryTaskInfoRequest {
	s.OrderId = &v
	return s
}

type QueryTaskInfoResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryTaskInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTaskInfoResponseBody) SetCode(v string) *QueryTaskInfoResponseBody {
	s.Code = &v
	return s
}

func (s *QueryTaskInfoResponseBody) SetData(v string) *QueryTaskInfoResponseBody {
	s.Data = &v
	return s
}

func (s *QueryTaskInfoResponseBody) SetMessage(v string) *QueryTaskInfoResponseBody {
	s.Message = &v
	return s
}

func (s *QueryTaskInfoResponseBody) SetRequestId(v string) *QueryTaskInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTaskInfoResponseBody) SetSuccess(v bool) *QueryTaskInfoResponseBody {
	s.Success = &v
	return s
}

type QueryTaskInfoResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryTaskInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryTaskInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryTaskInfoResponse) SetHeaders(v map[string]*string) *QueryTaskInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryTaskInfoResponse) SetStatusCode(v int32) *QueryTaskInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTaskInfoResponse) SetBody(v *QueryTaskInfoResponseBody) *QueryTaskInfoResponse {
	s.Body = v
	return s
}

type ReplyMessageApiRequest struct {
	MsgContent  *string `json:"MsgContent,omitempty" xml:"MsgContent,omitempty"`
	MsgType     *string `json:"MsgType,omitempty" xml:"MsgType,omitempty"`
	OpenGroupId *string `json:"OpenGroupId,omitempty" xml:"OpenGroupId,omitempty"`
	OrderId     *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName    *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ReplyMessageApiRequest) String() string {
	return tea.Prettify(s)
}

func (s ReplyMessageApiRequest) GoString() string {
	return s.String()
}

func (s *ReplyMessageApiRequest) SetMsgContent(v string) *ReplyMessageApiRequest {
	s.MsgContent = &v
	return s
}

func (s *ReplyMessageApiRequest) SetMsgType(v string) *ReplyMessageApiRequest {
	s.MsgType = &v
	return s
}

func (s *ReplyMessageApiRequest) SetOpenGroupId(v string) *ReplyMessageApiRequest {
	s.OpenGroupId = &v
	return s
}

func (s *ReplyMessageApiRequest) SetOrderId(v string) *ReplyMessageApiRequest {
	s.OrderId = &v
	return s
}

func (s *ReplyMessageApiRequest) SetUserId(v string) *ReplyMessageApiRequest {
	s.UserId = &v
	return s
}

func (s *ReplyMessageApiRequest) SetUserName(v string) *ReplyMessageApiRequest {
	s.UserName = &v
	return s
}

type ReplyMessageApiResponseBody struct {
	// code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// data
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// msg
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReplyMessageApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReplyMessageApiResponseBody) GoString() string {
	return s.String()
}

func (s *ReplyMessageApiResponseBody) SetCode(v string) *ReplyMessageApiResponseBody {
	s.Code = &v
	return s
}

func (s *ReplyMessageApiResponseBody) SetData(v string) *ReplyMessageApiResponseBody {
	s.Data = &v
	return s
}

func (s *ReplyMessageApiResponseBody) SetMessage(v string) *ReplyMessageApiResponseBody {
	s.Message = &v
	return s
}

func (s *ReplyMessageApiResponseBody) SetRequestId(v string) *ReplyMessageApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReplyMessageApiResponseBody) SetSuccess(v bool) *ReplyMessageApiResponseBody {
	s.Success = &v
	return s
}

type ReplyMessageApiResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ReplyMessageApiResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReplyMessageApiResponse) String() string {
	return tea.Prettify(s)
}

func (s ReplyMessageApiResponse) GoString() string {
	return s.String()
}

func (s *ReplyMessageApiResponse) SetHeaders(v map[string]*string) *ReplyMessageApiResponse {
	s.Headers = v
	return s
}

func (s *ReplyMessageApiResponse) SetStatusCode(v int32) *ReplyMessageApiResponse {
	s.StatusCode = &v
	return s
}

func (s *ReplyMessageApiResponse) SetBody(v *ReplyMessageApiResponseBody) *ReplyMessageApiResponse {
	s.Body = v
	return s
}

type RestOpenTaskOrderRequest struct {
	OpenGroupId  *string `json:"OpenGroupId,omitempty" xml:"OpenGroupId,omitempty"`
	OrderId      *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	ResetContent *string `json:"ResetContent,omitempty" xml:"ResetContent,omitempty"`
	ResetType    *string `json:"ResetType,omitempty" xml:"ResetType,omitempty"`
	UserName     *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s RestOpenTaskOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s RestOpenTaskOrderRequest) GoString() string {
	return s.String()
}

func (s *RestOpenTaskOrderRequest) SetOpenGroupId(v string) *RestOpenTaskOrderRequest {
	s.OpenGroupId = &v
	return s
}

func (s *RestOpenTaskOrderRequest) SetOrderId(v string) *RestOpenTaskOrderRequest {
	s.OrderId = &v
	return s
}

func (s *RestOpenTaskOrderRequest) SetResetContent(v string) *RestOpenTaskOrderRequest {
	s.ResetContent = &v
	return s
}

func (s *RestOpenTaskOrderRequest) SetResetType(v string) *RestOpenTaskOrderRequest {
	s.ResetType = &v
	return s
}

func (s *RestOpenTaskOrderRequest) SetUserName(v string) *RestOpenTaskOrderRequest {
	s.UserName = &v
	return s
}

type RestOpenTaskOrderResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RestOpenTaskOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RestOpenTaskOrderResponseBody) GoString() string {
	return s.String()
}

func (s *RestOpenTaskOrderResponseBody) SetCode(v string) *RestOpenTaskOrderResponseBody {
	s.Code = &v
	return s
}

func (s *RestOpenTaskOrderResponseBody) SetData(v string) *RestOpenTaskOrderResponseBody {
	s.Data = &v
	return s
}

func (s *RestOpenTaskOrderResponseBody) SetMessage(v string) *RestOpenTaskOrderResponseBody {
	s.Message = &v
	return s
}

func (s *RestOpenTaskOrderResponseBody) SetRequestId(v string) *RestOpenTaskOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestOpenTaskOrderResponseBody) SetSuccess(v bool) *RestOpenTaskOrderResponseBody {
	s.Success = &v
	return s
}

type RestOpenTaskOrderResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RestOpenTaskOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RestOpenTaskOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s RestOpenTaskOrderResponse) GoString() string {
	return s.String()
}

func (s *RestOpenTaskOrderResponse) SetHeaders(v map[string]*string) *RestOpenTaskOrderResponse {
	s.Headers = v
	return s
}

func (s *RestOpenTaskOrderResponse) SetStatusCode(v int32) *RestOpenTaskOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *RestOpenTaskOrderResponse) SetBody(v *RestOpenTaskOrderResponseBody) *RestOpenTaskOrderResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("support-plan"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CloseTaskOrderWithOptions(request *CloseTaskOrderRequest, runtime *util.RuntimeOptions) (_result *CloseTaskOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		query["OrderId"] = request.OrderId
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CloseTaskOrder"),
		Version:     tea.String("2021-07-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CloseTaskOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CloseTaskOrder(request *CloseTaskOrderRequest) (_result *CloseTaskOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CloseTaskOrderResponse{}
	_body, _err := client.CloseTaskOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTaskOrderWithOptions(request *CreateTaskOrderRequest, runtime *util.RuntimeOptions) (_result *CreateTaskOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomerRealName)) {
		query["CustomerRealName"] = request.CustomerRealName
	}

	if !tea.BoolValue(util.IsUnset(request.CustomerUserId)) {
		query["CustomerUserId"] = request.CustomerUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ImportantDescription)) {
		query["ImportantDescription"] = request.ImportantDescription
	}

	if !tea.BoolValue(util.IsUnset(request.IsImportant)) {
		query["IsImportant"] = request.IsImportant
	}

	if !tea.BoolValue(util.IsUnset(request.OpenGroupId)) {
		query["OpenGroupId"] = request.OpenGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		query["ProductType"] = request.ProductType
	}

	if !tea.BoolValue(util.IsUnset(request.ProductTypeName)) {
		query["ProductTypeName"] = request.ProductTypeName
	}

	if !tea.BoolValue(util.IsUnset(request.TaskTitle)) {
		query["TaskTitle"] = request.TaskTitle
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTaskOrder"),
		Version:     tea.String("2021-07-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTaskOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTaskOrder(request *CreateTaskOrderRequest) (_result *CreateTaskOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTaskOrderResponse{}
	_body, _err := client.CreateTaskOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTaskOrderByEventReportWithOptions(tmpReq *CreateTaskOrderByEventReportRequest, runtime *util.RuntimeOptions) (_result *CreateTaskOrderByEventReportResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateTaskOrderByEventReportShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.EventBody)) {
		request.EventBodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EventBody, tea.String("EventBody"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Extinfo)) {
		request.ExtinfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Extinfo, tea.String("Extinfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Business)) {
		query["Business"] = request.Business
	}

	if !tea.BoolValue(util.IsUnset(request.CreateRealName)) {
		query["CreateRealName"] = request.CreateRealName
	}

	if !tea.BoolValue(util.IsUnset(request.CreateUserId)) {
		query["CreateUserId"] = request.CreateUserId
	}

	if !tea.BoolValue(util.IsUnset(request.EventBodyShrink)) {
		query["EventBody"] = request.EventBodyShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ExtinfoShrink)) {
		query["Extinfo"] = request.ExtinfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ImportantDesc)) {
		query["ImportantDesc"] = request.ImportantDesc
	}

	if !tea.BoolValue(util.IsUnset(request.JoinChildGroupUserIds)) {
		query["JoinChildGroupUserIds"] = request.JoinChildGroupUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.MonitorCongregation)) {
		query["MonitorCongregation"] = request.MonitorCongregation
	}

	if !tea.BoolValue(util.IsUnset(request.OpenGroupId)) {
		query["OpenGroupId"] = request.OpenGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		query["ProductType"] = request.ProductType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTaskOrderByEventReport"),
		Version:     tea.String("2021-07-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTaskOrderByEventReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTaskOrderByEventReport(request *CreateTaskOrderByEventReportRequest) (_result *CreateTaskOrderByEventReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTaskOrderByEventReportResponse{}
	_body, _err := client.CreateTaskOrderByEventReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteEnterpriseDingtalkGroupCustomerMemberWithOptions(tmpReq *DeleteEnterpriseDingtalkGroupCustomerMemberRequest, runtime *util.RuntimeOptions) (_result *DeleteEnterpriseDingtalkGroupCustomerMemberResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DeleteEnterpriseDingtalkGroupCustomerMemberShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Mobiles)) {
		request.MobilesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Mobiles, tea.String("Mobiles"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MobilesShrink)) {
		body["Mobiles"] = request.MobilesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.OpenGroupId)) {
		body["OpenGroupId"] = request.OpenGroupId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteEnterpriseDingtalkGroupCustomerMember"),
		Version:     tea.String("2021-07-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteEnterpriseDingtalkGroupCustomerMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteEnterpriseDingtalkGroupCustomerMember(request *DeleteEnterpriseDingtalkGroupCustomerMemberRequest) (_result *DeleteEnterpriseDingtalkGroupCustomerMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteEnterpriseDingtalkGroupCustomerMemberResponse{}
	_body, _err := client.DeleteEnterpriseDingtalkGroupCustomerMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEnterpriseDingtalkGroupWithOptions(request *GetEnterpriseDingtalkGroupRequest, runtime *util.RuntimeOptions) (_result *GetEnterpriseDingtalkGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenGroupId)) {
		body["OpenGroupId"] = request.OpenGroupId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetEnterpriseDingtalkGroup"),
		Version:     tea.String("2021-07-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetEnterpriseDingtalkGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetEnterpriseDingtalkGroup(request *GetEnterpriseDingtalkGroupRequest) (_result *GetEnterpriseDingtalkGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetEnterpriseDingtalkGroupResponse{}
	_body, _err := client.GetEnterpriseDingtalkGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEnterpriseDingtalkGroupCustomerMemberWithOptions(request *GetEnterpriseDingtalkGroupCustomerMemberRequest, runtime *util.RuntimeOptions) (_result *GetEnterpriseDingtalkGroupCustomerMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Mobile)) {
		body["Mobile"] = request.Mobile
	}

	if !tea.BoolValue(util.IsUnset(request.OpenGroupId)) {
		body["OpenGroupId"] = request.OpenGroupId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetEnterpriseDingtalkGroupCustomerMember"),
		Version:     tea.String("2021-07-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetEnterpriseDingtalkGroupCustomerMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetEnterpriseDingtalkGroupCustomerMember(request *GetEnterpriseDingtalkGroupCustomerMemberRequest) (_result *GetEnterpriseDingtalkGroupCustomerMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetEnterpriseDingtalkGroupCustomerMemberResponse{}
	_body, _err := client.GetEnterpriseDingtalkGroupCustomerMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDdTaskOrderWithOptions(request *ListDdTaskOrderRequest, runtime *util.RuntimeOptions) (_result *ListDdTaskOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CallerParentId)) {
		query["CallerParentId"] = request.CallerParentId
	}

	if !tea.BoolValue(util.IsUnset(request.CallerType)) {
		query["CallerType"] = request.CallerType
	}

	if !tea.BoolValue(util.IsUnset(request.CallerUid)) {
		query["CallerUid"] = request.CallerUid
	}

	if !tea.BoolValue(util.IsUnset(request.OpenGroupId)) {
		query["OpenGroupId"] = request.OpenGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		query["OrderId"] = request.OrderId
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		query["RequestId"] = request.RequestId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskStatus)) {
		query["TaskStatus"] = request.TaskStatus
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDdTaskOrder"),
		Version:     tea.String("2021-07-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDdTaskOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDdTaskOrder(request *ListDdTaskOrderRequest) (_result *ListDdTaskOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDdTaskOrderResponse{}
	_body, _err := client.ListDdTaskOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListEnterpriseDingtalkGroupCustomerMembersWithOptions(request *ListEnterpriseDingtalkGroupCustomerMembersRequest, runtime *util.RuntimeOptions) (_result *ListEnterpriseDingtalkGroupCustomerMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenGroupId)) {
		body["OpenGroupId"] = request.OpenGroupId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListEnterpriseDingtalkGroupCustomerMembers"),
		Version:     tea.String("2021-07-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListEnterpriseDingtalkGroupCustomerMembersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListEnterpriseDingtalkGroupCustomerMembers(request *ListEnterpriseDingtalkGroupCustomerMembersRequest) (_result *ListEnterpriseDingtalkGroupCustomerMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListEnterpriseDingtalkGroupCustomerMembersResponse{}
	_body, _err := client.ListEnterpriseDingtalkGroupCustomerMembersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListEnterpriseDingtalkGroupsWithOptions(runtime *util.RuntimeOptions) (_result *ListEnterpriseDingtalkGroupsResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("ListEnterpriseDingtalkGroups"),
		Version:     tea.String("2021-07-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListEnterpriseDingtalkGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListEnterpriseDingtalkGroups() (_result *ListEnterpriseDingtalkGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListEnterpriseDingtalkGroupsResponse{}
	_body, _err := client.ListEnterpriseDingtalkGroupsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListProductByGroupWithOptions(request *ListProductByGroupRequest, runtime *util.RuntimeOptions) (_result *ListProductByGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenGroupId)) {
		query["OpenGroupId"] = request.OpenGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProductByGroup"),
		Version:     tea.String("2021-07-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProductByGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListProductByGroup(request *ListProductByGroupRequest) (_result *ListProductByGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListProductByGroupResponse{}
	_body, _err := client.ListProductByGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryTaskInfoWithOptions(request *QueryTaskInfoRequest, runtime *util.RuntimeOptions) (_result *QueryTaskInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		query["OrderId"] = request.OrderId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryTaskInfo"),
		Version:     tea.String("2021-07-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryTaskInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryTaskInfo(request *QueryTaskInfoRequest) (_result *QueryTaskInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryTaskInfoResponse{}
	_body, _err := client.QueryTaskInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReplyMessageApiWithOptions(request *ReplyMessageApiRequest, runtime *util.RuntimeOptions) (_result *ReplyMessageApiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MsgContent)) {
		query["MsgContent"] = request.MsgContent
	}

	if !tea.BoolValue(util.IsUnset(request.MsgType)) {
		query["MsgType"] = request.MsgType
	}

	if !tea.BoolValue(util.IsUnset(request.OpenGroupId)) {
		query["OpenGroupId"] = request.OpenGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		query["OrderId"] = request.OrderId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReplyMessageApi"),
		Version:     tea.String("2021-07-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReplyMessageApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReplyMessageApi(request *ReplyMessageApiRequest) (_result *ReplyMessageApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReplyMessageApiResponse{}
	_body, _err := client.ReplyMessageApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RestOpenTaskOrderWithOptions(request *RestOpenTaskOrderRequest, runtime *util.RuntimeOptions) (_result *RestOpenTaskOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenGroupId)) {
		query["OpenGroupId"] = request.OpenGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		query["OrderId"] = request.OrderId
	}

	if !tea.BoolValue(util.IsUnset(request.ResetContent)) {
		query["ResetContent"] = request.ResetContent
	}

	if !tea.BoolValue(util.IsUnset(request.ResetType)) {
		query["ResetType"] = request.ResetType
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RestOpenTaskOrder"),
		Version:     tea.String("2021-07-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RestOpenTaskOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RestOpenTaskOrder(request *RestOpenTaskOrderRequest) (_result *RestOpenTaskOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RestOpenTaskOrderResponse{}
	_body, _err := client.RestOpenTaskOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
