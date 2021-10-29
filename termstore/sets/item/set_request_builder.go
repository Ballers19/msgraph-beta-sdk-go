package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i08ca4c5bb1fe453cf948a9337ebc77517c10861b960c120f6a7bdf151e23091e "github.com/microsoftgraph/msgraph-beta-sdk-go/termstore/sets/item/children"
    ib4bf2188d376329fec86f06954eedc5c031d9700786bc5e49c5d813183216a29 "github.com/microsoftgraph/msgraph-beta-sdk-go/termstore/sets/item/terms"
    ibdb574ef03f0eec0fc47939c5b968e4c54cff1b5a2960cb579f607fb551d04c0 "github.com/microsoftgraph/msgraph-beta-sdk-go/termstore/sets/item/relations"
    icc7c8df3b795dc69afaa4bb6bdd0758e60da8070cbb794629c98ab73eaeb21f6 "github.com/microsoftgraph/msgraph-beta-sdk-go/termstore/sets/item/parentgroup"
    i1fdd91eaeb77b5232aaa6637b51c8335e644143a6bec1447e4639d345a19ce09 "github.com/microsoftgraph/msgraph-beta-sdk-go/termstore/sets/item/terms/item"
    ia9e638bcb5f0cac6ab821e69fff2fbadde36cbb4ae8162280da0e79f201144d2 "github.com/microsoftgraph/msgraph-beta-sdk-go/termstore/sets/item/children/item"
    icad4fc6304f6da07b148cc11a131410e87779371a7fe5d2470d5b412287bc38d "github.com/microsoftgraph/msgraph-beta-sdk-go/termstore/sets/item/relations/item"
)

// Builds and executes requests for operations under \termStore\sets\{set-id}
type SetRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type SetRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type SetRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *SetRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Collection of all sets available in the term store.
type SetRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type SetRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Set;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *SetRequestBuilder) Children()(*i08ca4c5bb1fe453cf948a9337ebc77517c10861b960c120f6a7bdf151e23091e.ChildrenRequestBuilder) {
    return i08ca4c5bb1fe453cf948a9337ebc77517c10861b960c120f6a7bdf151e23091e.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go.termStore.sets.item.children.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *SetRequestBuilder) ChildrenById(id string)(*ia9e638bcb5f0cac6ab821e69fff2fbadde36cbb4ae8162280da0e79f201144d2.TermRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["term_id"] = id
    }
    return ia9e638bcb5f0cac6ab821e69fff2fbadde36cbb4ae8162280da0e79f201144d2.NewTermRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Instantiates a new SetRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewSetRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SetRequestBuilder) {
    m := &SetRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/termStore/sets/{set_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new SetRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewSetRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SetRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSetRequestBuilderInternal(urlParams, requestAdapter)
}
// Collection of all sets available in the term store.
// Parameters:
//  - options : Options for the request
func (m *SetRequestBuilder) CreateDeleteRequestInformation(options *SetRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Collection of all sets available in the term store.
// Parameters:
//  - options : Options for the request
func (m *SetRequestBuilder) CreateGetRequestInformation(options *SetRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Collection of all sets available in the term store.
// Parameters:
//  - options : Options for the request
func (m *SetRequestBuilder) CreatePatchRequestInformation(options *SetRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Collection of all sets available in the term store.
// Parameters:
//  - options : Options for the request
func (m *SetRequestBuilder) Delete(options *SetRequestBuilderDeleteOptions)(error) {
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
// Collection of all sets available in the term store.
// Parameters:
//  - options : Options for the request
func (m *SetRequestBuilder) Get(options *SetRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Set, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewSet() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Set), nil
}
func (m *SetRequestBuilder) ParentGroup()(*icc7c8df3b795dc69afaa4bb6bdd0758e60da8070cbb794629c98ab73eaeb21f6.ParentGroupRequestBuilder) {
    return icc7c8df3b795dc69afaa4bb6bdd0758e60da8070cbb794629c98ab73eaeb21f6.NewParentGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Collection of all sets available in the term store.
// Parameters:
//  - options : Options for the request
func (m *SetRequestBuilder) Patch(options *SetRequestBuilderPatchOptions)(error) {
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
func (m *SetRequestBuilder) Relations()(*ibdb574ef03f0eec0fc47939c5b968e4c54cff1b5a2960cb579f607fb551d04c0.RelationsRequestBuilder) {
    return ibdb574ef03f0eec0fc47939c5b968e4c54cff1b5a2960cb579f607fb551d04c0.NewRelationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go.termStore.sets.item.relations.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *SetRequestBuilder) RelationsById(id string)(*icad4fc6304f6da07b148cc11a131410e87779371a7fe5d2470d5b412287bc38d.RelationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["relation_id"] = id
    }
    return icad4fc6304f6da07b148cc11a131410e87779371a7fe5d2470d5b412287bc38d.NewRelationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *SetRequestBuilder) Terms()(*ib4bf2188d376329fec86f06954eedc5c031d9700786bc5e49c5d813183216a29.TermsRequestBuilder) {
    return ib4bf2188d376329fec86f06954eedc5c031d9700786bc5e49c5d813183216a29.NewTermsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go.termStore.sets.item.terms.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *SetRequestBuilder) TermsById(id string)(*i1fdd91eaeb77b5232aaa6637b51c8335e644143a6bec1447e4639d345a19ce09.TermRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["term_id"] = id
    }
    return i1fdd91eaeb77b5232aaa6637b51c8335e644143a6bec1447e4639d345a19ce09.NewTermRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
