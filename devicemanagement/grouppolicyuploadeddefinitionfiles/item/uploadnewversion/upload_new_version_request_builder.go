package uploadnewversion

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// Builds and executes requests for operations under \deviceManagement\groupPolicyUploadedDefinitionFiles\{groupPolicyUploadedDefinitionFile-id}\microsoft.graph.uploadNewVersion
type UploadNewVersionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Post
type UploadNewVersionRequestBuilderPostOptions struct {
    // 
    Body *UploadNewVersionRequestBody;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Instantiates a new UploadNewVersionRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewUploadNewVersionRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UploadNewVersionRequestBuilder) {
    m := &UploadNewVersionRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/deviceManagement/groupPolicyUploadedDefinitionFiles/{groupPolicyUploadedDefinitionFile_id}/microsoft.graph.uploadNewVersion";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new UploadNewVersionRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewUploadNewVersionRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UploadNewVersionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUploadNewVersionRequestBuilderInternal(urlParams, requestAdapter)
}
// Invoke action uploadNewVersion
// Parameters:
//  - options : Options for the request
func (m *UploadNewVersionRequestBuilder) CreatePostRequestInformation(options *UploadNewVersionRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.POST
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Invoke action uploadNewVersion
// Parameters:
//  - options : Options for the request
func (m *UploadNewVersionRequestBuilder) Post(options *UploadNewVersionRequestBuilderPostOptions)(error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
