package tokenlifetimepolicies

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    ie62257f70fec21b9cd339074053fd870d0aec476a2e4106cdb8c727bb20e2bbf "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/tokenlifetimepolicies/ref"
)

// Builds and executes requests for operations under \applications\{application-id}\tokenLifetimePolicies
type TokenLifetimePoliciesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Get
type TokenLifetimePoliciesRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *TokenLifetimePoliciesRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// The tokenLifetimePolicies assigned to this application. Supports $expand.
type TokenLifetimePoliciesRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Include count of items
    Count *bool;
    // Expand related entities
    Expand []string;
    // Filter items by property values
    Filter *string;
    // Order items by property values
    Orderby []string;
    // Search items by search phrases
    Search *string;
    // Select properties to be returned
    Select_escaped []string;
    // Skip the first n items
    Skip *int32;
    // Show only the first n items
    Top *int32;
}
// Instantiates a new TokenLifetimePoliciesRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewTokenLifetimePoliciesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TokenLifetimePoliciesRequestBuilder) {
    m := &TokenLifetimePoliciesRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/applications/{application_id}/tokenLifetimePolicies{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new TokenLifetimePoliciesRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewTokenLifetimePoliciesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TokenLifetimePoliciesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTokenLifetimePoliciesRequestBuilderInternal(urlParams, requestAdapter)
}
// The tokenLifetimePolicies assigned to this application. Supports $expand.
// Parameters:
//  - options : Options for the request
func (m *TokenLifetimePoliciesRequestBuilder) CreateGetRequestInformation(options *TokenLifetimePoliciesRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        err := options.Q.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
    }
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
// The tokenLifetimePolicies assigned to this application. Supports $expand.
// Parameters:
//  - options : Options for the request
func (m *TokenLifetimePoliciesRequestBuilder) Get(options *TokenLifetimePoliciesRequestBuilderGetOptions)(*TokenLifetimePoliciesResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTokenLifetimePoliciesResponse() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*TokenLifetimePoliciesResponse), nil
}
func (m *TokenLifetimePoliciesRequestBuilder) Ref()(*ie62257f70fec21b9cd339074053fd870d0aec476a2e4106cdb8c727bb20e2bbf.RefRequestBuilder) {
    return ie62257f70fec21b9cd339074053fd870d0aec476a2e4106cdb8c727bb20e2bbf.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
