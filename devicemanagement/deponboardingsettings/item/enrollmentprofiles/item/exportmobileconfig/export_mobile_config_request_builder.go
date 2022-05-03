package exportmobileconfig

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ExportMobileConfigRequestBuilder provides operations to call the exportMobileConfig method.
type ExportMobileConfigRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ExportMobileConfigRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExportMobileConfigRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewExportMobileConfigRequestBuilderInternal instantiates a new ExportMobileConfigRequestBuilder and sets the default values.
func NewExportMobileConfigRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExportMobileConfigRequestBuilder) {
    m := &ExportMobileConfigRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/depOnboardingSettings/{depOnboardingSetting%2Did}/enrollmentProfiles/{enrollmentProfile%2Did}/microsoft.graph.exportMobileConfig()";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewExportMobileConfigRequestBuilder instantiates a new ExportMobileConfigRequestBuilder and sets the default values.
func NewExportMobileConfigRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExportMobileConfigRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExportMobileConfigRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformationWithRequestConfiguration exports the mobile configuration
func (m *ExportMobileConfigRequestBuilder) CreateGetRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration exports the mobile configuration
func (m *ExportMobileConfigRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *ExportMobileConfigRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// GetWithResponseHandler exports the mobile configuration
func (m *ExportMobileConfigRequestBuilder) GetWithResponseHandler(requestConfiguration *ExportMobileConfigRequestBuilderGetRequestConfiguration)(ExportMobileConfigResponseable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler exports the mobile configuration
func (m *ExportMobileConfigRequestBuilder) GetWithResponseHandler(requestConfiguration *ExportMobileConfigRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ExportMobileConfigResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateExportMobileConfigResponseFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(ExportMobileConfigResponseable), nil
}
