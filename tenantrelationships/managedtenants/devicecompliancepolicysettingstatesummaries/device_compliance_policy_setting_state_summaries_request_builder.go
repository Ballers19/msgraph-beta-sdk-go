package devicecompliancepolicysettingstatesummaries

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
    if1edd47d08b097c27bc827091cc41d48788f587da25946f8efb6fc254e726a6f "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/devicecompliancepolicysettingstatesummaries/count"
)

// DeviceCompliancePolicySettingStateSummariesRequestBuilder provides operations to manage the deviceCompliancePolicySettingStateSummaries property of the microsoft.graph.managedTenants.managedTenant entity.
type DeviceCompliancePolicySettingStateSummariesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceCompliancePolicySettingStateSummariesRequestBuilderGetOptions options for Get
type DeviceCompliancePolicySettingStateSummariesRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceCompliancePolicySettingStateSummariesRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// DeviceCompliancePolicySettingStateSummariesRequestBuilderGetQueryParameters summary information for device compliance policy setting states across managed tenants.
type DeviceCompliancePolicySettingStateSummariesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// DeviceCompliancePolicySettingStateSummariesRequestBuilderPostOptions options for Post
type DeviceCompliancePolicySettingStateSummariesRequestBuilderPostOptions struct {
    // 
    Body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.DeviceCompliancePolicySettingStateSummaryable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// NewDeviceCompliancePolicySettingStateSummariesRequestBuilderInternal instantiates a new DeviceCompliancePolicySettingStateSummariesRequestBuilder and sets the default values.
func NewDeviceCompliancePolicySettingStateSummariesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceCompliancePolicySettingStateSummariesRequestBuilder) {
    m := &DeviceCompliancePolicySettingStateSummariesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/tenantRelationships/managedTenants/deviceCompliancePolicySettingStateSummaries{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceCompliancePolicySettingStateSummariesRequestBuilder instantiates a new DeviceCompliancePolicySettingStateSummariesRequestBuilder and sets the default values.
func NewDeviceCompliancePolicySettingStateSummariesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceCompliancePolicySettingStateSummariesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceCompliancePolicySettingStateSummariesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the count property
func (m *DeviceCompliancePolicySettingStateSummariesRequestBuilder) Count()(*if1edd47d08b097c27bc827091cc41d48788f587da25946f8efb6fc254e726a6f.CountRequestBuilder) {
    return if1edd47d08b097c27bc827091cc41d48788f587da25946f8efb6fc254e726a6f.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation summary information for device compliance policy setting states across managed tenants.
func (m *DeviceCompliancePolicySettingStateSummariesRequestBuilder) CreateGetRequestInformation(options *DeviceCompliancePolicySettingStateSummariesRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePostRequestInformation create new navigation property to deviceCompliancePolicySettingStateSummaries for tenantRelationships
func (m *DeviceCompliancePolicySettingStateSummariesRequestBuilder) CreatePostRequestInformation(options *DeviceCompliancePolicySettingStateSummariesRequestBuilderPostOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Get summary information for device compliance policy setting states across managed tenants.
func (m *DeviceCompliancePolicySettingStateSummariesRequestBuilder) Get(options *DeviceCompliancePolicySettingStateSummariesRequestBuilderGetOptions)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.DeviceCompliancePolicySettingStateSummaryCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateDeviceCompliancePolicySettingStateSummaryCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.DeviceCompliancePolicySettingStateSummaryCollectionResponseable), nil
}
// Post create new navigation property to deviceCompliancePolicySettingStateSummaries for tenantRelationships
func (m *DeviceCompliancePolicySettingStateSummariesRequestBuilder) Post(options *DeviceCompliancePolicySettingStateSummariesRequestBuilderPostOptions)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.DeviceCompliancePolicySettingStateSummaryable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateDeviceCompliancePolicySettingStateSummaryFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.DeviceCompliancePolicySettingStateSummaryable), nil
}
