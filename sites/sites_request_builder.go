package sites

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i01987e74e1c88ee7412edc4fecc08f77a9bd5b9625489df2dcf277b733bd32a5 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/remove"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i6f81ada91b35c4c8b5de72483e03b83321824cd7301bf2b1faec3e76d10920ab "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/add"
    i78c52f4a1bffb07514485793e74ce9ff957279c79c8593b58f1d24032586a995 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/delta"
    ib81c95029d979466d91938adeca5dddd85b40bf0f696454eb18f46a1732d3a9c "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/count"
)

// SitesRequestBuilder provides operations to manage the collection of site entities.
type SitesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SitesRequestBuilderGetQueryParameters search across a SharePoint tenant for [sites][] that match keywords provided. The only property that works for sorting is **createdDateTime**. The search filter is a free text search that uses multiple properties when retrieving the search results.
type SitesRequestBuilderGetQueryParameters struct {
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
// SitesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SitesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SitesRequestBuilderGetQueryParameters
}
// Add the add property
func (m *SitesRequestBuilder) Add()(*i6f81ada91b35c4c8b5de72483e03b83321824cd7301bf2b1faec3e76d10920ab.AddRequestBuilder) {
    return i6f81ada91b35c4c8b5de72483e03b83321824cd7301bf2b1faec3e76d10920ab.NewAddRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewSitesRequestBuilderInternal instantiates a new SitesRequestBuilder and sets the default values.
func NewSitesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SitesRequestBuilder) {
    m := &SitesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSitesRequestBuilder instantiates a new SitesRequestBuilder and sets the default values.
func NewSitesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SitesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSitesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the count property
func (m *SitesRequestBuilder) Count()(*ib81c95029d979466d91938adeca5dddd85b40bf0f696454eb18f46a1732d3a9c.CountRequestBuilder) {
    return ib81c95029d979466d91938adeca5dddd85b40bf0f696454eb18f46a1732d3a9c.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation search across a SharePoint tenant for [sites][] that match keywords provided. The only property that works for sorting is **createdDateTime**. The search filter is a free text search that uses multiple properties when retrieving the search results.
func (m *SitesRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration search across a SharePoint tenant for [sites][] that match keywords provided. The only property that works for sorting is **createdDateTime**. The search filter is a free text search that uses multiple properties when retrieving the search results.
func (m *SitesRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *SitesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delta provides operations to call the delta method.
func (m *SitesRequestBuilder) Delta()(*i78c52f4a1bffb07514485793e74ce9ff957279c79c8593b58f1d24032586a995.DeltaRequestBuilder) {
    return i78c52f4a1bffb07514485793e74ce9ff957279c79c8593b58f1d24032586a995.NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get search across a SharePoint tenant for [sites][] that match keywords provided. The only property that works for sorting is **createdDateTime**. The search filter is a free text search that uses multiple properties when retrieving the search results.
func (m *SitesRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SiteCollectionResponseable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler search across a SharePoint tenant for [sites][] that match keywords provided. The only property that works for sorting is **createdDateTime**. The search filter is a free text search that uses multiple properties when retrieving the search results.
func (m *SitesRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *SitesRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SiteCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSiteCollectionResponseFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SiteCollectionResponseable), nil
}
// Remove the remove property
func (m *SitesRequestBuilder) Remove()(*i01987e74e1c88ee7412edc4fecc08f77a9bd5b9625489df2dcf277b733bd32a5.RemoveRequestBuilder) {
    return i01987e74e1c88ee7412edc4fecc08f77a9bd5b9625489df2dcf277b733bd32a5.NewRemoveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
