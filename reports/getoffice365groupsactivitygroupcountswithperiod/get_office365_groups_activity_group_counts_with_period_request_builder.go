package getoffice365groupsactivitygroupcountswithperiod

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilder provides operations to call the getOffice365GroupsActivityGroupCounts method.
type GetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilderInternal instantiates a new GetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilder and sets the default values.
func NewGetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, period *string)(*GetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilder) {
    m := &GetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/reports/microsoft.graph.getOffice365GroupsActivityGroupCounts(period='{period}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if period != nil {
        urlTplParams[""] = *period
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilder instantiates a new GetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilder and sets the default values.
func NewGetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformationWithRequestConfiguration invoke function getOffice365GroupsActivityGroupCounts
func (m *GetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilder) CreateGetRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration invoke function getOffice365GroupsActivityGroupCounts
func (m *GetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *GetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// GetWithResponseHandler invoke function getOffice365GroupsActivityGroupCounts
func (m *GetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilder) GetWithResponseHandler(requestConfiguration *GetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilderGetRequestConfiguration)(GetOffice365GroupsActivityGroupCountsWithPeriodResponseable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler invoke function getOffice365GroupsActivityGroupCounts
func (m *GetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilder) GetWithResponseHandler(requestConfiguration *GetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(GetOffice365GroupsActivityGroupCountsWithPeriodResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetOffice365GroupsActivityGroupCountsWithPeriodResponseFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetOffice365GroupsActivityGroupCountsWithPeriodResponseable), nil
}
