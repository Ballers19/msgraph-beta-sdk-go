package beginonboarding

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// BeginOnboardingRequestBuilder provides operations to call the beginOnboarding method.
type BeginOnboardingRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// BeginOnboardingRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BeginOnboardingRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewBeginOnboardingRequestBuilderInternal instantiates a new BeginOnboardingRequestBuilder and sets the default values.
func NewBeginOnboardingRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BeginOnboardingRequestBuilder) {
    m := &BeginOnboardingRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/remoteAssistancePartners/{remoteAssistancePartner%2Did}/microsoft.graph.beginOnboarding";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewBeginOnboardingRequestBuilder instantiates a new BeginOnboardingRequestBuilder and sets the default values.
func NewBeginOnboardingRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BeginOnboardingRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBeginOnboardingRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformationWithRequestConfiguration a request to start onboarding.  Must be coupled with the appropriate TeamViewer account information
func (m *BeginOnboardingRequestBuilder) CreatePostRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(nil);
}
// CreatePostRequestInformationWithRequestConfiguration a request to start onboarding.  Must be coupled with the appropriate TeamViewer account information
func (m *BeginOnboardingRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(requestConfiguration *BeginOnboardingRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// PostWithResponseHandler a request to start onboarding.  Must be coupled with the appropriate TeamViewer account information
func (m *BeginOnboardingRequestBuilder) PostWithResponseHandler(requestConfiguration *BeginOnboardingRequestBuilderPostRequestConfiguration)(error) {
    return m.PostWithResponseHandler(requestConfiguration, nil);
}
// PostWithResponseHandler a request to start onboarding.  Must be coupled with the appropriate TeamViewer account information
func (m *BeginOnboardingRequestBuilder) PostWithResponseHandler(requestConfiguration *BeginOnboardingRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, nil)
    if err != nil {
        return err
    }
    return nil
}
