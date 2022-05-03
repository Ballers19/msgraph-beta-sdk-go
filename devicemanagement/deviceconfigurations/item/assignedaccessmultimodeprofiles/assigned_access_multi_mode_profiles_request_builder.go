package assignedaccessmultimodeprofiles

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// AssignedAccessMultiModeProfilesRequestBuilder provides operations to call the assignedAccessMultiModeProfiles method.
type AssignedAccessMultiModeProfilesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AssignedAccessMultiModeProfilesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AssignedAccessMultiModeProfilesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAssignedAccessMultiModeProfilesRequestBuilderInternal instantiates a new AssignedAccessMultiModeProfilesRequestBuilder and sets the default values.
func NewAssignedAccessMultiModeProfilesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AssignedAccessMultiModeProfilesRequestBuilder) {
    m := &AssignedAccessMultiModeProfilesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceConfigurations/{deviceConfiguration%2Did}/microsoft.graph.assignedAccessMultiModeProfiles";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAssignedAccessMultiModeProfilesRequestBuilder instantiates a new AssignedAccessMultiModeProfilesRequestBuilder and sets the default values.
func NewAssignedAccessMultiModeProfilesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AssignedAccessMultiModeProfilesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAssignedAccessMultiModeProfilesRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformationWithRequestConfiguration invoke action assignedAccessMultiModeProfiles
func (m *AssignedAccessMultiModeProfilesRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body AssignedAccessMultiModeProfilesRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action assignedAccessMultiModeProfiles
func (m *AssignedAccessMultiModeProfilesRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body AssignedAccessMultiModeProfilesRequestBodyable, requestConfiguration *AssignedAccessMultiModeProfilesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// PostWithResponseHandler invoke action assignedAccessMultiModeProfiles
func (m *AssignedAccessMultiModeProfilesRequestBuilder) PostWithResponseHandler(body AssignedAccessMultiModeProfilesRequestBodyable, requestConfiguration *AssignedAccessMultiModeProfilesRequestBuilderPostRequestConfiguration)(error) {
    return m.PostWithResponseHandler(body, requestConfiguration, nil);
}
// PostWithResponseHandler invoke action assignedAccessMultiModeProfiles
func (m *AssignedAccessMultiModeProfilesRequestBuilder) PostWithResponseHandler(body AssignedAccessMultiModeProfilesRequestBodyable, requestConfiguration *AssignedAccessMultiModeProfilesRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
