package appmanagementpolicies

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i494befe9e10024b14b7a34c156ca01699d6c87fdb1afdafa57b3483cbc4b0025 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/appmanagementpolicies/count"
    if8d921a2197e1c2db547d5610111532d1766c5db9692b825698130064c138c12 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/appmanagementpolicies/ref"
)

// AppManagementPoliciesRequestBuilder provides operations to manage the appManagementPolicies property of the microsoft.graph.application entity.
type AppManagementPoliciesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AppManagementPoliciesRequestBuilderGetQueryParameters the appManagementPolicy applied to this application.
type AppManagementPoliciesRequestBuilderGetQueryParameters struct {
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
// AppManagementPoliciesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppManagementPoliciesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AppManagementPoliciesRequestBuilderGetQueryParameters
}
// NewAppManagementPoliciesRequestBuilderInternal instantiates a new AppManagementPoliciesRequestBuilder and sets the default values.
func NewAppManagementPoliciesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AppManagementPoliciesRequestBuilder) {
    m := &AppManagementPoliciesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/applications/{application%2Did}/appManagementPolicies{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAppManagementPoliciesRequestBuilder instantiates a new AppManagementPoliciesRequestBuilder and sets the default values.
func NewAppManagementPoliciesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AppManagementPoliciesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAppManagementPoliciesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the count property
func (m *AppManagementPoliciesRequestBuilder) Count()(*i494befe9e10024b14b7a34c156ca01699d6c87fdb1afdafa57b3483cbc4b0025.CountRequestBuilder) {
    return i494befe9e10024b14b7a34c156ca01699d6c87fdb1afdafa57b3483cbc4b0025.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the appManagementPolicy applied to this application.
func (m *AppManagementPoliciesRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the appManagementPolicy applied to this application.
func (m *AppManagementPoliciesRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *AppManagementPoliciesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get the appManagementPolicy applied to this application.
func (m *AppManagementPoliciesRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppManagementPolicyCollectionResponseable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler the appManagementPolicy applied to this application.
func (m *AppManagementPoliciesRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *AppManagementPoliciesRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppManagementPolicyCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAppManagementPolicyCollectionResponseFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppManagementPolicyCollectionResponseable), nil
}
// Ref the ref property
func (m *AppManagementPoliciesRequestBuilder) Ref()(*if8d921a2197e1c2db547d5610111532d1766c5db9692b825698130064c138c12.RefRequestBuilder) {
    return if8d921a2197e1c2db547d5610111532d1766c5db9692b825698130064c138c12.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
