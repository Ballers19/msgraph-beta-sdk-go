package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0084c64e69f2de1db99e3db4eebd9d8a41444eb2f0f9d19cc03286f0b896875a "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicydefinitions/item/nextversiondefinition"
    i91fccd52e0d78587541f985ebb0b59be750da1e83e56b1da10b05e4eb4dc73ac "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicydefinitions/item/definitionfile"
    ia5547c0682358c94012164a326f029053ec5f1ddfbb91573107909f094b955bc "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicydefinitions/item/presentations"
    ib5e5e1e82d57aaa3cb2752c79932c323ee118f915bbc0dc5e5ab7600a66a7f27 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicydefinitions/item/category"
    ide1c03ea82181c831c7c8a172c8a94f8d4edd80cc28a64d3a11e7c2c58c4ab3e "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicydefinitions/item/previousversiondefinition"
    i755dd99b2492e3a413ac1208e4de2ede29dfb284de02d498fb55c3247fe139f0 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicydefinitions/item/presentations/item"
)

// Builds and executes requests for operations under \deviceManagement\groupPolicyDefinitions\{groupPolicyDefinition-id}
type GroupPolicyDefinitionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type GroupPolicyDefinitionRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type GroupPolicyDefinitionRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *GroupPolicyDefinitionRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// The available group policy definitions for this account.
type GroupPolicyDefinitionRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type GroupPolicyDefinitionRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GroupPolicyDefinition;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *GroupPolicyDefinitionRequestBuilder) Category()(*ib5e5e1e82d57aaa3cb2752c79932c323ee118f915bbc0dc5e5ab7600a66a7f27.CategoryRequestBuilder) {
    return ib5e5e1e82d57aaa3cb2752c79932c323ee118f915bbc0dc5e5ab7600a66a7f27.NewCategoryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new GroupPolicyDefinitionRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewGroupPolicyDefinitionRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GroupPolicyDefinitionRequestBuilder) {
    m := &GroupPolicyDefinitionRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/deviceManagement/groupPolicyDefinitions/{groupPolicyDefinition_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new GroupPolicyDefinitionRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewGroupPolicyDefinitionRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GroupPolicyDefinitionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupPolicyDefinitionRequestBuilderInternal(urlParams, requestAdapter)
}
// The available group policy definitions for this account.
// Parameters:
//  - options : Options for the request
func (m *GroupPolicyDefinitionRequestBuilder) CreateDeleteRequestInformation(options *GroupPolicyDefinitionRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
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
// The available group policy definitions for this account.
// Parameters:
//  - options : Options for the request
func (m *GroupPolicyDefinitionRequestBuilder) CreateGetRequestInformation(options *GroupPolicyDefinitionRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The available group policy definitions for this account.
// Parameters:
//  - options : Options for the request
func (m *GroupPolicyDefinitionRequestBuilder) CreatePatchRequestInformation(options *GroupPolicyDefinitionRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
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
func (m *GroupPolicyDefinitionRequestBuilder) DefinitionFile()(*i91fccd52e0d78587541f985ebb0b59be750da1e83e56b1da10b05e4eb4dc73ac.DefinitionFileRequestBuilder) {
    return i91fccd52e0d78587541f985ebb0b59be750da1e83e56b1da10b05e4eb4dc73ac.NewDefinitionFileRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The available group policy definitions for this account.
// Parameters:
//  - options : Options for the request
func (m *GroupPolicyDefinitionRequestBuilder) Delete(options *GroupPolicyDefinitionRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// The available group policy definitions for this account.
// Parameters:
//  - options : Options for the request
func (m *GroupPolicyDefinitionRequestBuilder) Get(options *GroupPolicyDefinitionRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GroupPolicyDefinition, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewGroupPolicyDefinition() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GroupPolicyDefinition), nil
}
func (m *GroupPolicyDefinitionRequestBuilder) NextVersionDefinition()(*i0084c64e69f2de1db99e3db4eebd9d8a41444eb2f0f9d19cc03286f0b896875a.NextVersionDefinitionRequestBuilder) {
    return i0084c64e69f2de1db99e3db4eebd9d8a41444eb2f0f9d19cc03286f0b896875a.NewNextVersionDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The available group policy definitions for this account.
// Parameters:
//  - options : Options for the request
func (m *GroupPolicyDefinitionRequestBuilder) Patch(options *GroupPolicyDefinitionRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *GroupPolicyDefinitionRequestBuilder) Presentations()(*ia5547c0682358c94012164a326f029053ec5f1ddfbb91573107909f094b955bc.PresentationsRequestBuilder) {
    return ia5547c0682358c94012164a326f029053ec5f1ddfbb91573107909f094b955bc.NewPresentationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go.deviceManagement.groupPolicyDefinitions.item.presentations.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *GroupPolicyDefinitionRequestBuilder) PresentationsById(id string)(*i755dd99b2492e3a413ac1208e4de2ede29dfb284de02d498fb55c3247fe139f0.GroupPolicyPresentationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["groupPolicyPresentation_id"] = id
    }
    return i755dd99b2492e3a413ac1208e4de2ede29dfb284de02d498fb55c3247fe139f0.NewGroupPolicyPresentationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *GroupPolicyDefinitionRequestBuilder) PreviousVersionDefinition()(*ide1c03ea82181c831c7c8a172c8a94f8d4edd80cc28a64d3a11e7c2c58c4ab3e.PreviousVersionDefinitionRequestBuilder) {
    return ide1c03ea82181c831c7c8a172c8a94f8d4edd80cc28a64d3a11e7c2c58c4ab3e.NewPreviousVersionDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
