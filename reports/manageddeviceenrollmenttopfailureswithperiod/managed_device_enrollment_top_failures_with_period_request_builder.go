package manageddeviceenrollmenttopfailureswithperiod

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilder provides operations to call the managedDeviceEnrollmentTopFailures method.
type ManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilderInternal instantiates a new ManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilder and sets the default values.
func NewManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, period *string)(*ManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilder) {
    m := &ManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/reports/microsoft.graph.managedDeviceEnrollmentTopFailures(period='{period}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if period != nil {
        urlTplParams["period"] = *period
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilder instantiates a new ManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilder and sets the default values.
func NewManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformation invoke function managedDeviceEnrollmentTopFailures
func (m *ManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration invoke function managedDeviceEnrollmentTopFailures
func (m *ManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *ManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get invoke function managedDeviceEnrollmentTopFailures
func (m *ManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Reportable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler invoke function managedDeviceEnrollmentTopFailures
func (m *ManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *ManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Reportable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateReportFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Reportable), nil
}
