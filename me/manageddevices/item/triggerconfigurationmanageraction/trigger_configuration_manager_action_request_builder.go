package triggerconfigurationmanageraction

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// TriggerConfigurationManagerActionRequestBuilder provides operations to call the triggerConfigurationManagerAction method.
type TriggerConfigurationManagerActionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TriggerConfigurationManagerActionRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TriggerConfigurationManagerActionRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTriggerConfigurationManagerActionRequestBuilderInternal instantiates a new TriggerConfigurationManagerActionRequestBuilder and sets the default values.
func NewTriggerConfigurationManagerActionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TriggerConfigurationManagerActionRequestBuilder) {
    m := &TriggerConfigurationManagerActionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/managedDevices/{managedDevice%2Did}/microsoft.graph.triggerConfigurationManagerAction";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTriggerConfigurationManagerActionRequestBuilder instantiates a new TriggerConfigurationManagerActionRequestBuilder and sets the default values.
func NewTriggerConfigurationManagerActionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TriggerConfigurationManagerActionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTriggerConfigurationManagerActionRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformationWithRequestConfiguration trigger action on ConfigurationManager client
func (m *TriggerConfigurationManagerActionRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body TriggerConfigurationManagerActionRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration trigger action on ConfigurationManager client
func (m *TriggerConfigurationManagerActionRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body TriggerConfigurationManagerActionRequestBodyable, requestConfiguration *TriggerConfigurationManagerActionRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// PostWithResponseHandler trigger action on ConfigurationManager client
func (m *TriggerConfigurationManagerActionRequestBuilder) PostWithResponseHandler(body TriggerConfigurationManagerActionRequestBodyable, requestConfiguration *TriggerConfigurationManagerActionRequestBuilderPostRequestConfiguration)(error) {
    return m.PostWithResponseHandler(body, requestConfiguration, nil);
}
// PostWithResponseHandler trigger action on ConfigurationManager client
func (m *TriggerConfigurationManagerActionRequestBuilder) PostWithResponseHandler(body TriggerConfigurationManagerActionRequestBodyable, requestConfiguration *TriggerConfigurationManagerActionRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, nil)
    if err != nil {
        return err
    }
    return nil
}
