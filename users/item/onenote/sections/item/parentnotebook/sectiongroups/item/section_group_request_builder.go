package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i3e248e648310acf5fa9970d0c0b542c1815ff697a20ba4bb30aedfd1d15074e5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onenote/sections/item/parentnotebook/sectiongroups/item/sectiongroups"
    i5c63e8b1040d3f76bb731998b0836d372b398aaba691d6ea1cb4481dab664d7f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onenote/sections/item/parentnotebook/sectiongroups/item/sections"
    ia60151e916ff9a69cfade1045f650439b4bd26b58dd06c8ce799e4a4d38d8c78 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onenote/sections/item/parentnotebook/sectiongroups/item/parentsectiongroup"
    id0b1e287b236a940eaac1a954117bee3464c6d0c09738c80c983d8a2544188c7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onenote/sections/item/parentnotebook/sectiongroups/item/parentnotebook"
    i6398ea1ef1995778bf7711b7807a57b0b5e9ef375adf996b56071267f0381283 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onenote/sections/item/parentnotebook/sectiongroups/item/sectiongroups/item"
    if4f58cf06760f16cdba63ced17982af3fe1e6df86130165765ae9b63bd68cc96 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onenote/sections/item/parentnotebook/sectiongroups/item/sections/item"
)

// Builds and executes requests for operations under \users\{user-id}\onenote\sections\{onenoteSection-id}\parentNotebook\sectionGroups\{sectionGroup-id}
type SectionGroupRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type SectionGroupRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type SectionGroupRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *SectionGroupRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// The section groups in the notebook. Read-only. Nullable.
type SectionGroupRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type SectionGroupRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SectionGroup;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Instantiates a new SectionGroupRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewSectionGroupRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SectionGroupRequestBuilder) {
    m := &SectionGroupRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/onenote/sections/{onenoteSection_id}/parentNotebook/sectionGroups/{sectionGroup_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new SectionGroupRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewSectionGroupRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SectionGroupRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSectionGroupRequestBuilderInternal(urlParams, requestAdapter)
}
// The section groups in the notebook. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *SectionGroupRequestBuilder) CreateDeleteRequestInformation(options *SectionGroupRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The section groups in the notebook. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *SectionGroupRequestBuilder) CreateGetRequestInformation(options *SectionGroupRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The section groups in the notebook. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *SectionGroupRequestBuilder) CreatePatchRequestInformation(options *SectionGroupRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The section groups in the notebook. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *SectionGroupRequestBuilder) Delete(options *SectionGroupRequestBuilderDeleteOptions)(error) {
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
// The section groups in the notebook. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *SectionGroupRequestBuilder) Get(options *SectionGroupRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SectionGroup, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewSectionGroup() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SectionGroup), nil
}
func (m *SectionGroupRequestBuilder) ParentNotebook()(*id0b1e287b236a940eaac1a954117bee3464c6d0c09738c80c983d8a2544188c7.ParentNotebookRequestBuilder) {
    return id0b1e287b236a940eaac1a954117bee3464c6d0c09738c80c983d8a2544188c7.NewParentNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SectionGroupRequestBuilder) ParentSectionGroup()(*ia60151e916ff9a69cfade1045f650439b4bd26b58dd06c8ce799e4a4d38d8c78.ParentSectionGroupRequestBuilder) {
    return ia60151e916ff9a69cfade1045f650439b4bd26b58dd06c8ce799e4a4d38d8c78.NewParentSectionGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The section groups in the notebook. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *SectionGroupRequestBuilder) Patch(options *SectionGroupRequestBuilderPatchOptions)(error) {
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
func (m *SectionGroupRequestBuilder) SectionGroups()(*i3e248e648310acf5fa9970d0c0b542c1815ff697a20ba4bb30aedfd1d15074e5.SectionGroupsRequestBuilder) {
    return i3e248e648310acf5fa9970d0c0b542c1815ff697a20ba4bb30aedfd1d15074e5.NewSectionGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.onenote.sections.item.parentNotebook.sectionGroups.item.sectionGroups.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *SectionGroupRequestBuilder) SectionGroupsById(id string)(*i6398ea1ef1995778bf7711b7807a57b0b5e9ef375adf996b56071267f0381283.SectionGroupRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sectionGroup_id1"] = id
    }
    return i6398ea1ef1995778bf7711b7807a57b0b5e9ef375adf996b56071267f0381283.NewSectionGroupRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *SectionGroupRequestBuilder) Sections()(*i5c63e8b1040d3f76bb731998b0836d372b398aaba691d6ea1cb4481dab664d7f.SectionsRequestBuilder) {
    return i5c63e8b1040d3f76bb731998b0836d372b398aaba691d6ea1cb4481dab664d7f.NewSectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.onenote.sections.item.parentNotebook.sectionGroups.item.sections.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *SectionGroupRequestBuilder) SectionsById(id string)(*if4f58cf06760f16cdba63ced17982af3fe1e6df86130165765ae9b63bd68cc96.OnenoteSectionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteSection_id1"] = id
    }
    return if4f58cf06760f16cdba63ced17982af3fe1e6df86130165765ae9b63bd68cc96.NewOnenoteSectionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
