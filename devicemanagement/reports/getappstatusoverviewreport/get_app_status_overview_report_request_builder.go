package getappstatusoverviewreport

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetAppStatusOverviewReportRequestBuilder provides operations to call the getAppStatusOverviewReport method.
type GetAppStatusOverviewReportRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetAppStatusOverviewReportRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetAppStatusOverviewReportRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetAppStatusOverviewReportRequestBuilderInternal instantiates a new GetAppStatusOverviewReportRequestBuilder and sets the default values.
func NewGetAppStatusOverviewReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetAppStatusOverviewReportRequestBuilder) {
    m := &GetAppStatusOverviewReportRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/reports/microsoft.graph.getAppStatusOverviewReport";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetAppStatusOverviewReportRequestBuilder instantiates a new GetAppStatusOverviewReportRequestBuilder and sets the default values.
func NewGetAppStatusOverviewReportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetAppStatusOverviewReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetAppStatusOverviewReportRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformationWithRequestConfiguration invoke action getAppStatusOverviewReport
func (m *GetAppStatusOverviewReportRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body GetAppStatusOverviewReportRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action getAppStatusOverviewReport
func (m *GetAppStatusOverviewReportRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body GetAppStatusOverviewReportRequestBodyable, requestConfiguration *GetAppStatusOverviewReportRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// PostWithResponseHandler invoke action getAppStatusOverviewReport
func (m *GetAppStatusOverviewReportRequestBuilder) PostWithResponseHandler(body GetAppStatusOverviewReportRequestBodyable, requestConfiguration *GetAppStatusOverviewReportRequestBuilderPostRequestConfiguration)(GetAppStatusOverviewReportResponseable, error) {
    return m.PostWithResponseHandler(body, requestConfiguration, nil);
}
// PostWithResponseHandler invoke action getAppStatusOverviewReport
func (m *GetAppStatusOverviewReportRequestBuilder) PostWithResponseHandler(body GetAppStatusOverviewReportRequestBodyable, requestConfiguration *GetAppStatusOverviewReportRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(GetAppStatusOverviewReportResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetAppStatusOverviewReportResponseFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetAppStatusOverviewReportResponseable), nil
}
