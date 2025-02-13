package riskyserviceprincipals

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i418a1a5558948b4662e488386dc90986e3fc661f4e22322afbb353d02d78d298 "github.com/microsoftgraph/msgraph-beta-sdk-go/identityprotection/riskyserviceprincipals/dismiss"
    i85dee32ed1d0f963b95f0e6653515c81e7dc0dc00db2e6cc11adfd6b4476e28f "github.com/microsoftgraph/msgraph-beta-sdk-go/identityprotection/riskyserviceprincipals/confirmcompromised"
    ib48ad338b29836e614d9379952d94ca511eae67832181629704a82189ddf853c "github.com/microsoftgraph/msgraph-beta-sdk-go/identityprotection/riskyserviceprincipals/count"
)

// RiskyServicePrincipalsRequestBuilder provides operations to manage the riskyServicePrincipals property of the microsoft.graph.identityProtectionRoot entity.
type RiskyServicePrincipalsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// RiskyServicePrincipalsRequestBuilderGetQueryParameters azure AD service principals that are at risk.
type RiskyServicePrincipalsRequestBuilderGetQueryParameters struct {
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
// RiskyServicePrincipalsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RiskyServicePrincipalsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RiskyServicePrincipalsRequestBuilderGetQueryParameters
}
// RiskyServicePrincipalsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RiskyServicePrincipalsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ConfirmCompromised the confirmCompromised property
func (m *RiskyServicePrincipalsRequestBuilder) ConfirmCompromised()(*i85dee32ed1d0f963b95f0e6653515c81e7dc0dc00db2e6cc11adfd6b4476e28f.ConfirmCompromisedRequestBuilder) {
    return i85dee32ed1d0f963b95f0e6653515c81e7dc0dc00db2e6cc11adfd6b4476e28f.NewConfirmCompromisedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewRiskyServicePrincipalsRequestBuilderInternal instantiates a new RiskyServicePrincipalsRequestBuilder and sets the default values.
func NewRiskyServicePrincipalsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RiskyServicePrincipalsRequestBuilder) {
    m := &RiskyServicePrincipalsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityProtection/riskyServicePrincipals{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRiskyServicePrincipalsRequestBuilder instantiates a new RiskyServicePrincipalsRequestBuilder and sets the default values.
func NewRiskyServicePrincipalsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RiskyServicePrincipalsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRiskyServicePrincipalsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the count property
func (m *RiskyServicePrincipalsRequestBuilder) Count()(*ib48ad338b29836e614d9379952d94ca511eae67832181629704a82189ddf853c.CountRequestBuilder) {
    return ib48ad338b29836e614d9379952d94ca511eae67832181629704a82189ddf853c.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation azure AD service principals that are at risk.
func (m *RiskyServicePrincipalsRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration azure AD service principals that are at risk.
func (m *RiskyServicePrincipalsRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *RiskyServicePrincipalsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to riskyServicePrincipals for identityProtection
func (m *RiskyServicePrincipalsRequestBuilder) CreatePostRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RiskyServicePrincipalable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration create new navigation property to riskyServicePrincipals for identityProtection
func (m *RiskyServicePrincipalsRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RiskyServicePrincipalable, requestConfiguration *RiskyServicePrincipalsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Dismiss the dismiss property
func (m *RiskyServicePrincipalsRequestBuilder) Dismiss()(*i418a1a5558948b4662e488386dc90986e3fc661f4e22322afbb353d02d78d298.DismissRequestBuilder) {
    return i418a1a5558948b4662e488386dc90986e3fc661f4e22322afbb353d02d78d298.NewDismissRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get azure AD service principals that are at risk.
func (m *RiskyServicePrincipalsRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RiskyServicePrincipalCollectionResponseable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler azure AD service principals that are at risk.
func (m *RiskyServicePrincipalsRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *RiskyServicePrincipalsRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RiskyServicePrincipalCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateRiskyServicePrincipalCollectionResponseFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RiskyServicePrincipalCollectionResponseable), nil
}
// Post create new navigation property to riskyServicePrincipals for identityProtection
func (m *RiskyServicePrincipalsRequestBuilder) Post(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RiskyServicePrincipalable)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RiskyServicePrincipalable, error) {
    return m.PostWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler create new navigation property to riskyServicePrincipals for identityProtection
func (m *RiskyServicePrincipalsRequestBuilder) PostWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RiskyServicePrincipalable, requestConfiguration *RiskyServicePrincipalsRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RiskyServicePrincipalable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateRiskyServicePrincipalFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RiskyServicePrincipalable), nil
}
