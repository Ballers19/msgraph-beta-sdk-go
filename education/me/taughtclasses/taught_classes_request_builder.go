package taughtclasses

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i7a6e601f2581d1fd89799b3d3562cb1a60af36439c72822242cf57c2865ffeb5 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/me/taughtclasses/delta"
    ieb6b17e9be2b3e79da372d3ed15040599635eb90f081d0a3d8b20408b3ad3818 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/me/taughtclasses/ref"
)

// Builds and executes requests for operations under \education\me\taughtClasses
type TaughtClassesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Get
type TaughtClassesRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *TaughtClassesRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Classes for which the user is a teacher.
type TaughtClassesRequestBuilderGetQueryParameters struct {
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
// Instantiates a new TaughtClassesRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewTaughtClassesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TaughtClassesRequestBuilder) {
    m := &TaughtClassesRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/education/me/taughtClasses{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new TaughtClassesRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewTaughtClassesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TaughtClassesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTaughtClassesRequestBuilderInternal(urlParams, requestAdapter)
}
// Classes for which the user is a teacher.
// Parameters:
//  - options : Options for the request
func (m *TaughtClassesRequestBuilder) CreateGetRequestInformation(options *TaughtClassesRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Builds and executes requests for operations under \education\me\taughtClasses\microsoft.graph.delta()
func (m *TaughtClassesRequestBuilder) Delta()(*i7a6e601f2581d1fd89799b3d3562cb1a60af36439c72822242cf57c2865ffeb5.DeltaRequestBuilder) {
    return i7a6e601f2581d1fd89799b3d3562cb1a60af36439c72822242cf57c2865ffeb5.NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Classes for which the user is a teacher.
// Parameters:
//  - options : Options for the request
func (m *TaughtClassesRequestBuilder) Get(options *TaughtClassesRequestBuilderGetOptions)(*TaughtClassesResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTaughtClassesResponse() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*TaughtClassesResponse), nil
}
func (m *TaughtClassesRequestBuilder) Ref()(*ieb6b17e9be2b3e79da372d3ed15040599635eb90f081d0a3d8b20408b3ad3818.RefRequestBuilder) {
    return ieb6b17e9be2b3e79da372d3ed15040599635eb90f081d0a3d8b20408b3ad3818.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
