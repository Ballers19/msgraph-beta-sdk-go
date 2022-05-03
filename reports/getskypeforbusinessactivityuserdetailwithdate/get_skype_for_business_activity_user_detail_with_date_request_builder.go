package getskypeforbusinessactivityuserdetailwithdate

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GetSkypeForBusinessActivityUserDetailWithDateRequestBuilder provides operations to call the getSkypeForBusinessActivityUserDetail method.
type GetSkypeForBusinessActivityUserDetailWithDateRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetSkypeForBusinessActivityUserDetailWithDateRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetSkypeForBusinessActivityUserDetailWithDateRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetSkypeForBusinessActivityUserDetailWithDateRequestBuilderInternal instantiates a new GetSkypeForBusinessActivityUserDetailWithDateRequestBuilder and sets the default values.
func NewGetSkypeForBusinessActivityUserDetailWithDateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*GetSkypeForBusinessActivityUserDetailWithDateRequestBuilder) {
    m := &GetSkypeForBusinessActivityUserDetailWithDateRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/reports/microsoft.graph.getSkypeForBusinessActivityUserDetail(date={date})";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if date != nil {
        urlTplParams[""] = (*date).String()
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetSkypeForBusinessActivityUserDetailWithDateRequestBuilder instantiates a new GetSkypeForBusinessActivityUserDetailWithDateRequestBuilder and sets the default values.
func NewGetSkypeForBusinessActivityUserDetailWithDateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetSkypeForBusinessActivityUserDetailWithDateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetSkypeForBusinessActivityUserDetailWithDateRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformationWithRequestConfiguration invoke function getSkypeForBusinessActivityUserDetail
func (m *GetSkypeForBusinessActivityUserDetailWithDateRequestBuilder) CreateGetRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration invoke function getSkypeForBusinessActivityUserDetail
func (m *GetSkypeForBusinessActivityUserDetailWithDateRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *GetSkypeForBusinessActivityUserDetailWithDateRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// GetWithResponseHandler invoke function getSkypeForBusinessActivityUserDetail
func (m *GetSkypeForBusinessActivityUserDetailWithDateRequestBuilder) GetWithResponseHandler(requestConfiguration *GetSkypeForBusinessActivityUserDetailWithDateRequestBuilderGetRequestConfiguration)(GetSkypeForBusinessActivityUserDetailWithDateResponseable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler invoke function getSkypeForBusinessActivityUserDetail
func (m *GetSkypeForBusinessActivityUserDetailWithDateRequestBuilder) GetWithResponseHandler(requestConfiguration *GetSkypeForBusinessActivityUserDetailWithDateRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(GetSkypeForBusinessActivityUserDetailWithDateResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetSkypeForBusinessActivityUserDetailWithDateResponseFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetSkypeForBusinessActivityUserDetailWithDateResponseable), nil
}
