package shareforschooldatasyncservice

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ShareForSchoolDataSyncServiceRequestBuilder provides operations to call the shareForSchoolDataSyncService method.
type ShareForSchoolDataSyncServiceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ShareForSchoolDataSyncServiceRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ShareForSchoolDataSyncServiceRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewShareForSchoolDataSyncServiceRequestBuilderInternal instantiates a new ShareForSchoolDataSyncServiceRequestBuilder and sets the default values.
func NewShareForSchoolDataSyncServiceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ShareForSchoolDataSyncServiceRequestBuilder) {
    m := &ShareForSchoolDataSyncServiceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/depOnboardingSettings/{depOnboardingSetting%2Did}/microsoft.graph.shareForSchoolDataSyncService";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewShareForSchoolDataSyncServiceRequestBuilder instantiates a new ShareForSchoolDataSyncServiceRequestBuilder and sets the default values.
func NewShareForSchoolDataSyncServiceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ShareForSchoolDataSyncServiceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewShareForSchoolDataSyncServiceRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformationWithRequestConfiguration invoke action shareForSchoolDataSyncService
func (m *ShareForSchoolDataSyncServiceRequestBuilder) CreatePostRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action shareForSchoolDataSyncService
func (m *ShareForSchoolDataSyncServiceRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(requestConfiguration *ShareForSchoolDataSyncServiceRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// PostWithResponseHandler invoke action shareForSchoolDataSyncService
func (m *ShareForSchoolDataSyncServiceRequestBuilder) PostWithResponseHandler(requestConfiguration *ShareForSchoolDataSyncServiceRequestBuilderPostRequestConfiguration)(error) {
    return m.PostWithResponseHandler(requestConfiguration, nil);
}
// PostWithResponseHandler invoke action shareForSchoolDataSyncService
func (m *ShareForSchoolDataSyncServiceRequestBuilder) PostWithResponseHandler(requestConfiguration *ShareForSchoolDataSyncServiceRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
