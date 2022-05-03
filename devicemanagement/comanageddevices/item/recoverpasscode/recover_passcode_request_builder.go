package recoverpasscode

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// RecoverPasscodeRequestBuilder provides operations to call the recoverPasscode method.
type RecoverPasscodeRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// RecoverPasscodeRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RecoverPasscodeRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewRecoverPasscodeRequestBuilderInternal instantiates a new RecoverPasscodeRequestBuilder and sets the default values.
func NewRecoverPasscodeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RecoverPasscodeRequestBuilder) {
    m := &RecoverPasscodeRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/comanagedDevices/{managedDevice%2Did}/microsoft.graph.recoverPasscode";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRecoverPasscodeRequestBuilder instantiates a new RecoverPasscodeRequestBuilder and sets the default values.
func NewRecoverPasscodeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RecoverPasscodeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRecoverPasscodeRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformationWithRequestConfiguration recover passcode
func (m *RecoverPasscodeRequestBuilder) CreatePostRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(nil);
}
// CreatePostRequestInformationWithRequestConfiguration recover passcode
func (m *RecoverPasscodeRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(requestConfiguration *RecoverPasscodeRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// PostWithResponseHandler recover passcode
func (m *RecoverPasscodeRequestBuilder) PostWithResponseHandler(requestConfiguration *RecoverPasscodeRequestBuilderPostRequestConfiguration)(error) {
    return m.PostWithResponseHandler(requestConfiguration, nil);
}
// PostWithResponseHandler recover passcode
func (m *RecoverPasscodeRequestBuilder) PostWithResponseHandler(requestConfiguration *RecoverPasscodeRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
