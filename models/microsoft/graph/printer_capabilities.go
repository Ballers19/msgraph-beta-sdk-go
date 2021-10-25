package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type PrinterCapabilities struct {
    additionalData map[string]interface{};
    bottomMargins []int32;
    collation *bool;
    colorModes []PrintColorMode;
    contentTypes []string;
    copiesPerJob *IntegerRange;
    dpis []int32;
    duplexModes []PrintDuplexMode;
    feedDirections []PrinterFeedDirection;
    feedOrientations []PrinterFeedOrientation;
    finishings []PrintFinishing;
    inputBins []string;
    isColorPrintingSupported *bool;
    isPageRangeSupported *bool;
    leftMargins []int32;
    mediaColors []string;
    mediaSizes []string;
    mediaTypes []string;
    multipageLayouts []PrintMultipageLayout;
    orientations []PrintOrientation;
    outputBins []string;
    pagesPerSheet []int32;
    qualities []PrintQuality;
    rightMargins []int32;
    scalings []PrintScaling;
    supportedColorConfigurations []PrintColorConfiguration;
    supportedCopiesPerJob *IntegerRange;
    supportedDocumentMimeTypes []string;
    supportedDuplexConfigurations []PrintDuplexConfiguration;
    supportedFinishings []PrintFinishing;
    supportedMediaColors []string;
    supportedMediaSizes []string;
    supportedMediaTypes []PrintMediaType;
    supportedOrientations []PrintOrientation;
    supportedOutputBins []string;
    supportedPagesPerSheet *IntegerRange;
    supportedPresentationDirections []PrintPresentationDirection;
    supportedPrintQualities []PrintQuality;
    supportsFitPdfToPage *bool;
    topMargins []int32;
}
func NewPrinterCapabilities()(*PrinterCapabilities) {
    m := &PrinterCapabilities{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *PrinterCapabilities) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *PrinterCapabilities) GetBottomMargins()([]int32) {
    if m == nil {
        return nil
    } else {
        return m.bottomMargins
    }
}
func (m *PrinterCapabilities) GetCollation()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.collation
    }
}
func (m *PrinterCapabilities) GetColorModes()([]PrintColorMode) {
    if m == nil {
        return nil
    } else {
        return m.colorModes
    }
}
func (m *PrinterCapabilities) GetContentTypes()([]string) {
    if m == nil {
        return nil
    } else {
        return m.contentTypes
    }
}
func (m *PrinterCapabilities) GetCopiesPerJob()(*IntegerRange) {
    if m == nil {
        return nil
    } else {
        return m.copiesPerJob
    }
}
func (m *PrinterCapabilities) GetDpis()([]int32) {
    if m == nil {
        return nil
    } else {
        return m.dpis
    }
}
func (m *PrinterCapabilities) GetDuplexModes()([]PrintDuplexMode) {
    if m == nil {
        return nil
    } else {
        return m.duplexModes
    }
}
func (m *PrinterCapabilities) GetFeedDirections()([]PrinterFeedDirection) {
    if m == nil {
        return nil
    } else {
        return m.feedDirections
    }
}
func (m *PrinterCapabilities) GetFeedOrientations()([]PrinterFeedOrientation) {
    if m == nil {
        return nil
    } else {
        return m.feedOrientations
    }
}
func (m *PrinterCapabilities) GetFinishings()([]PrintFinishing) {
    if m == nil {
        return nil
    } else {
        return m.finishings
    }
}
func (m *PrinterCapabilities) GetInputBins()([]string) {
    if m == nil {
        return nil
    } else {
        return m.inputBins
    }
}
func (m *PrinterCapabilities) GetIsColorPrintingSupported()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isColorPrintingSupported
    }
}
func (m *PrinterCapabilities) GetIsPageRangeSupported()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isPageRangeSupported
    }
}
func (m *PrinterCapabilities) GetLeftMargins()([]int32) {
    if m == nil {
        return nil
    } else {
        return m.leftMargins
    }
}
func (m *PrinterCapabilities) GetMediaColors()([]string) {
    if m == nil {
        return nil
    } else {
        return m.mediaColors
    }
}
func (m *PrinterCapabilities) GetMediaSizes()([]string) {
    if m == nil {
        return nil
    } else {
        return m.mediaSizes
    }
}
func (m *PrinterCapabilities) GetMediaTypes()([]string) {
    if m == nil {
        return nil
    } else {
        return m.mediaTypes
    }
}
func (m *PrinterCapabilities) GetMultipageLayouts()([]PrintMultipageLayout) {
    if m == nil {
        return nil
    } else {
        return m.multipageLayouts
    }
}
func (m *PrinterCapabilities) GetOrientations()([]PrintOrientation) {
    if m == nil {
        return nil
    } else {
        return m.orientations
    }
}
func (m *PrinterCapabilities) GetOutputBins()([]string) {
    if m == nil {
        return nil
    } else {
        return m.outputBins
    }
}
func (m *PrinterCapabilities) GetPagesPerSheet()([]int32) {
    if m == nil {
        return nil
    } else {
        return m.pagesPerSheet
    }
}
func (m *PrinterCapabilities) GetQualities()([]PrintQuality) {
    if m == nil {
        return nil
    } else {
        return m.qualities
    }
}
func (m *PrinterCapabilities) GetRightMargins()([]int32) {
    if m == nil {
        return nil
    } else {
        return m.rightMargins
    }
}
func (m *PrinterCapabilities) GetScalings()([]PrintScaling) {
    if m == nil {
        return nil
    } else {
        return m.scalings
    }
}
func (m *PrinterCapabilities) GetSupportedColorConfigurations()([]PrintColorConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.supportedColorConfigurations
    }
}
func (m *PrinterCapabilities) GetSupportedCopiesPerJob()(*IntegerRange) {
    if m == nil {
        return nil
    } else {
        return m.supportedCopiesPerJob
    }
}
func (m *PrinterCapabilities) GetSupportedDocumentMimeTypes()([]string) {
    if m == nil {
        return nil
    } else {
        return m.supportedDocumentMimeTypes
    }
}
func (m *PrinterCapabilities) GetSupportedDuplexConfigurations()([]PrintDuplexConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.supportedDuplexConfigurations
    }
}
func (m *PrinterCapabilities) GetSupportedFinishings()([]PrintFinishing) {
    if m == nil {
        return nil
    } else {
        return m.supportedFinishings
    }
}
func (m *PrinterCapabilities) GetSupportedMediaColors()([]string) {
    if m == nil {
        return nil
    } else {
        return m.supportedMediaColors
    }
}
func (m *PrinterCapabilities) GetSupportedMediaSizes()([]string) {
    if m == nil {
        return nil
    } else {
        return m.supportedMediaSizes
    }
}
func (m *PrinterCapabilities) GetSupportedMediaTypes()([]PrintMediaType) {
    if m == nil {
        return nil
    } else {
        return m.supportedMediaTypes
    }
}
func (m *PrinterCapabilities) GetSupportedOrientations()([]PrintOrientation) {
    if m == nil {
        return nil
    } else {
        return m.supportedOrientations
    }
}
func (m *PrinterCapabilities) GetSupportedOutputBins()([]string) {
    if m == nil {
        return nil
    } else {
        return m.supportedOutputBins
    }
}
func (m *PrinterCapabilities) GetSupportedPagesPerSheet()(*IntegerRange) {
    if m == nil {
        return nil
    } else {
        return m.supportedPagesPerSheet
    }
}
func (m *PrinterCapabilities) GetSupportedPresentationDirections()([]PrintPresentationDirection) {
    if m == nil {
        return nil
    } else {
        return m.supportedPresentationDirections
    }
}
func (m *PrinterCapabilities) GetSupportedPrintQualities()([]PrintQuality) {
    if m == nil {
        return nil
    } else {
        return m.supportedPrintQualities
    }
}
func (m *PrinterCapabilities) GetSupportsFitPdfToPage()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.supportsFitPdfToPage
    }
}
func (m *PrinterCapabilities) GetTopMargins()([]int32) {
    if m == nil {
        return nil
    } else {
        return m.topMargins
    }
}
func (m *PrinterCapabilities) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["bottomMargins"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("int32")
        if err != nil {
            return err
        }
        res := make([]int32, len(val))
        for i, v := range val {
            res[i] = v.(int32)
        }
        m.SetBottomMargins(res)
        return nil
    }
    res["collation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetCollation(val)
        return nil
    }
    res["colorModes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParsePrintColorMode)
        if err != nil {
            return err
        }
        res := make([]PrintColorMode, len(val))
        for i, v := range val {
            res[i] = *(v.(*PrintColorMode))
        }
        m.SetColorModes(res)
        return nil
    }
    res["contentTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetContentTypes(res)
        return nil
    }
    res["copiesPerJob"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIntegerRange() })
        if err != nil {
            return err
        }
        m.SetCopiesPerJob(val.(*IntegerRange))
        return nil
    }
    res["dpis"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("int32")
        if err != nil {
            return err
        }
        res := make([]int32, len(val))
        for i, v := range val {
            res[i] = v.(int32)
        }
        m.SetDpis(res)
        return nil
    }
    res["duplexModes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParsePrintDuplexMode)
        if err != nil {
            return err
        }
        res := make([]PrintDuplexMode, len(val))
        for i, v := range val {
            res[i] = *(v.(*PrintDuplexMode))
        }
        m.SetDuplexModes(res)
        return nil
    }
    res["feedDirections"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParsePrinterFeedDirection)
        if err != nil {
            return err
        }
        res := make([]PrinterFeedDirection, len(val))
        for i, v := range val {
            res[i] = *(v.(*PrinterFeedDirection))
        }
        m.SetFeedDirections(res)
        return nil
    }
    res["feedOrientations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParsePrinterFeedOrientation)
        if err != nil {
            return err
        }
        res := make([]PrinterFeedOrientation, len(val))
        for i, v := range val {
            res[i] = *(v.(*PrinterFeedOrientation))
        }
        m.SetFeedOrientations(res)
        return nil
    }
    res["finishings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParsePrintFinishing)
        if err != nil {
            return err
        }
        res := make([]PrintFinishing, len(val))
        for i, v := range val {
            res[i] = *(v.(*PrintFinishing))
        }
        m.SetFinishings(res)
        return nil
    }
    res["inputBins"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetInputBins(res)
        return nil
    }
    res["isColorPrintingSupported"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsColorPrintingSupported(val)
        return nil
    }
    res["isPageRangeSupported"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsPageRangeSupported(val)
        return nil
    }
    res["leftMargins"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("int32")
        if err != nil {
            return err
        }
        res := make([]int32, len(val))
        for i, v := range val {
            res[i] = v.(int32)
        }
        m.SetLeftMargins(res)
        return nil
    }
    res["mediaColors"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetMediaColors(res)
        return nil
    }
    res["mediaSizes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetMediaSizes(res)
        return nil
    }
    res["mediaTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetMediaTypes(res)
        return nil
    }
    res["multipageLayouts"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParsePrintMultipageLayout)
        if err != nil {
            return err
        }
        res := make([]PrintMultipageLayout, len(val))
        for i, v := range val {
            res[i] = *(v.(*PrintMultipageLayout))
        }
        m.SetMultipageLayouts(res)
        return nil
    }
    res["orientations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParsePrintOrientation)
        if err != nil {
            return err
        }
        res := make([]PrintOrientation, len(val))
        for i, v := range val {
            res[i] = *(v.(*PrintOrientation))
        }
        m.SetOrientations(res)
        return nil
    }
    res["outputBins"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetOutputBins(res)
        return nil
    }
    res["pagesPerSheet"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("int32")
        if err != nil {
            return err
        }
        res := make([]int32, len(val))
        for i, v := range val {
            res[i] = v.(int32)
        }
        m.SetPagesPerSheet(res)
        return nil
    }
    res["qualities"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParsePrintQuality)
        if err != nil {
            return err
        }
        res := make([]PrintQuality, len(val))
        for i, v := range val {
            res[i] = *(v.(*PrintQuality))
        }
        m.SetQualities(res)
        return nil
    }
    res["rightMargins"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("int32")
        if err != nil {
            return err
        }
        res := make([]int32, len(val))
        for i, v := range val {
            res[i] = v.(int32)
        }
        m.SetRightMargins(res)
        return nil
    }
    res["scalings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParsePrintScaling)
        if err != nil {
            return err
        }
        res := make([]PrintScaling, len(val))
        for i, v := range val {
            res[i] = *(v.(*PrintScaling))
        }
        m.SetScalings(res)
        return nil
    }
    res["supportedColorConfigurations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParsePrintColorConfiguration)
        if err != nil {
            return err
        }
        res := make([]PrintColorConfiguration, len(val))
        for i, v := range val {
            res[i] = *(v.(*PrintColorConfiguration))
        }
        m.SetSupportedColorConfigurations(res)
        return nil
    }
    res["supportedCopiesPerJob"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIntegerRange() })
        if err != nil {
            return err
        }
        m.SetSupportedCopiesPerJob(val.(*IntegerRange))
        return nil
    }
    res["supportedDocumentMimeTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetSupportedDocumentMimeTypes(res)
        return nil
    }
    res["supportedDuplexConfigurations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParsePrintDuplexConfiguration)
        if err != nil {
            return err
        }
        res := make([]PrintDuplexConfiguration, len(val))
        for i, v := range val {
            res[i] = *(v.(*PrintDuplexConfiguration))
        }
        m.SetSupportedDuplexConfigurations(res)
        return nil
    }
    res["supportedFinishings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParsePrintFinishing)
        if err != nil {
            return err
        }
        res := make([]PrintFinishing, len(val))
        for i, v := range val {
            res[i] = *(v.(*PrintFinishing))
        }
        m.SetSupportedFinishings(res)
        return nil
    }
    res["supportedMediaColors"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetSupportedMediaColors(res)
        return nil
    }
    res["supportedMediaSizes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetSupportedMediaSizes(res)
        return nil
    }
    res["supportedMediaTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParsePrintMediaType)
        if err != nil {
            return err
        }
        res := make([]PrintMediaType, len(val))
        for i, v := range val {
            res[i] = *(v.(*PrintMediaType))
        }
        m.SetSupportedMediaTypes(res)
        return nil
    }
    res["supportedOrientations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParsePrintOrientation)
        if err != nil {
            return err
        }
        res := make([]PrintOrientation, len(val))
        for i, v := range val {
            res[i] = *(v.(*PrintOrientation))
        }
        m.SetSupportedOrientations(res)
        return nil
    }
    res["supportedOutputBins"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetSupportedOutputBins(res)
        return nil
    }
    res["supportedPagesPerSheet"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIntegerRange() })
        if err != nil {
            return err
        }
        m.SetSupportedPagesPerSheet(val.(*IntegerRange))
        return nil
    }
    res["supportedPresentationDirections"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParsePrintPresentationDirection)
        if err != nil {
            return err
        }
        res := make([]PrintPresentationDirection, len(val))
        for i, v := range val {
            res[i] = *(v.(*PrintPresentationDirection))
        }
        m.SetSupportedPresentationDirections(res)
        return nil
    }
    res["supportedPrintQualities"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParsePrintQuality)
        if err != nil {
            return err
        }
        res := make([]PrintQuality, len(val))
        for i, v := range val {
            res[i] = *(v.(*PrintQuality))
        }
        m.SetSupportedPrintQualities(res)
        return nil
    }
    res["supportsFitPdfToPage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetSupportsFitPdfToPage(val)
        return nil
    }
    res["topMargins"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("int32")
        if err != nil {
            return err
        }
        res := make([]int32, len(val))
        for i, v := range val {
            res[i] = v.(int32)
        }
        m.SetTopMargins(res)
        return nil
    }
    return res
}
func (m *PrinterCapabilities) IsNil()(bool) {
    return m == nil
}
func (m *PrinterCapabilities) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteCollectionOfInt32Values("bottomMargins", m.GetBottomMargins())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("collation", m.GetCollation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("colorModes", SerializePrintColorMode(m.GetColorModes()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("contentTypes", m.GetContentTypes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("copiesPerJob", m.GetCopiesPerJob())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfInt32Values("dpis", m.GetDpis())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("duplexModes", SerializePrintDuplexMode(m.GetDuplexModes()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("feedDirections", SerializePrinterFeedDirection(m.GetFeedDirections()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("feedOrientations", SerializePrinterFeedOrientation(m.GetFeedOrientations()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("finishings", SerializePrintFinishing(m.GetFinishings()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("inputBins", m.GetInputBins())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isColorPrintingSupported", m.GetIsColorPrintingSupported())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isPageRangeSupported", m.GetIsPageRangeSupported())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfInt32Values("leftMargins", m.GetLeftMargins())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("mediaColors", m.GetMediaColors())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("mediaSizes", m.GetMediaSizes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("mediaTypes", m.GetMediaTypes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("multipageLayouts", SerializePrintMultipageLayout(m.GetMultipageLayouts()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("orientations", SerializePrintOrientation(m.GetOrientations()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("outputBins", m.GetOutputBins())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfInt32Values("pagesPerSheet", m.GetPagesPerSheet())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("qualities", SerializePrintQuality(m.GetQualities()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfInt32Values("rightMargins", m.GetRightMargins())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("scalings", SerializePrintScaling(m.GetScalings()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("supportedColorConfigurations", SerializePrintColorConfiguration(m.GetSupportedColorConfigurations()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("supportedCopiesPerJob", m.GetSupportedCopiesPerJob())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("supportedDocumentMimeTypes", m.GetSupportedDocumentMimeTypes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("supportedDuplexConfigurations", SerializePrintDuplexConfiguration(m.GetSupportedDuplexConfigurations()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("supportedFinishings", SerializePrintFinishing(m.GetSupportedFinishings()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("supportedMediaColors", m.GetSupportedMediaColors())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("supportedMediaSizes", m.GetSupportedMediaSizes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("supportedMediaTypes", SerializePrintMediaType(m.GetSupportedMediaTypes()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("supportedOrientations", SerializePrintOrientation(m.GetSupportedOrientations()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("supportedOutputBins", m.GetSupportedOutputBins())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("supportedPagesPerSheet", m.GetSupportedPagesPerSheet())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("supportedPresentationDirections", SerializePrintPresentationDirection(m.GetSupportedPresentationDirections()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("supportedPrintQualities", SerializePrintQuality(m.GetSupportedPrintQualities()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("supportsFitPdfToPage", m.GetSupportsFitPdfToPage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfInt32Values("topMargins", m.GetTopMargins())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *PrinterCapabilities) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *PrinterCapabilities) SetBottomMargins(value []int32)() {
    m.bottomMargins = value
}
func (m *PrinterCapabilities) SetCollation(value *bool)() {
    m.collation = value
}
func (m *PrinterCapabilities) SetColorModes(value []PrintColorMode)() {
    m.colorModes = value
}
func (m *PrinterCapabilities) SetContentTypes(value []string)() {
    m.contentTypes = value
}
func (m *PrinterCapabilities) SetCopiesPerJob(value *IntegerRange)() {
    m.copiesPerJob = value
}
func (m *PrinterCapabilities) SetDpis(value []int32)() {
    m.dpis = value
}
func (m *PrinterCapabilities) SetDuplexModes(value []PrintDuplexMode)() {
    m.duplexModes = value
}
func (m *PrinterCapabilities) SetFeedDirections(value []PrinterFeedDirection)() {
    m.feedDirections = value
}
func (m *PrinterCapabilities) SetFeedOrientations(value []PrinterFeedOrientation)() {
    m.feedOrientations = value
}
func (m *PrinterCapabilities) SetFinishings(value []PrintFinishing)() {
    m.finishings = value
}
func (m *PrinterCapabilities) SetInputBins(value []string)() {
    m.inputBins = value
}
func (m *PrinterCapabilities) SetIsColorPrintingSupported(value *bool)() {
    m.isColorPrintingSupported = value
}
func (m *PrinterCapabilities) SetIsPageRangeSupported(value *bool)() {
    m.isPageRangeSupported = value
}
func (m *PrinterCapabilities) SetLeftMargins(value []int32)() {
    m.leftMargins = value
}
func (m *PrinterCapabilities) SetMediaColors(value []string)() {
    m.mediaColors = value
}
func (m *PrinterCapabilities) SetMediaSizes(value []string)() {
    m.mediaSizes = value
}
func (m *PrinterCapabilities) SetMediaTypes(value []string)() {
    m.mediaTypes = value
}
func (m *PrinterCapabilities) SetMultipageLayouts(value []PrintMultipageLayout)() {
    m.multipageLayouts = value
}
func (m *PrinterCapabilities) SetOrientations(value []PrintOrientation)() {
    m.orientations = value
}
func (m *PrinterCapabilities) SetOutputBins(value []string)() {
    m.outputBins = value
}
func (m *PrinterCapabilities) SetPagesPerSheet(value []int32)() {
    m.pagesPerSheet = value
}
func (m *PrinterCapabilities) SetQualities(value []PrintQuality)() {
    m.qualities = value
}
func (m *PrinterCapabilities) SetRightMargins(value []int32)() {
    m.rightMargins = value
}
func (m *PrinterCapabilities) SetScalings(value []PrintScaling)() {
    m.scalings = value
}
func (m *PrinterCapabilities) SetSupportedColorConfigurations(value []PrintColorConfiguration)() {
    m.supportedColorConfigurations = value
}
func (m *PrinterCapabilities) SetSupportedCopiesPerJob(value *IntegerRange)() {
    m.supportedCopiesPerJob = value
}
func (m *PrinterCapabilities) SetSupportedDocumentMimeTypes(value []string)() {
    m.supportedDocumentMimeTypes = value
}
func (m *PrinterCapabilities) SetSupportedDuplexConfigurations(value []PrintDuplexConfiguration)() {
    m.supportedDuplexConfigurations = value
}
func (m *PrinterCapabilities) SetSupportedFinishings(value []PrintFinishing)() {
    m.supportedFinishings = value
}
func (m *PrinterCapabilities) SetSupportedMediaColors(value []string)() {
    m.supportedMediaColors = value
}
func (m *PrinterCapabilities) SetSupportedMediaSizes(value []string)() {
    m.supportedMediaSizes = value
}
func (m *PrinterCapabilities) SetSupportedMediaTypes(value []PrintMediaType)() {
    m.supportedMediaTypes = value
}
func (m *PrinterCapabilities) SetSupportedOrientations(value []PrintOrientation)() {
    m.supportedOrientations = value
}
func (m *PrinterCapabilities) SetSupportedOutputBins(value []string)() {
    m.supportedOutputBins = value
}
func (m *PrinterCapabilities) SetSupportedPagesPerSheet(value *IntegerRange)() {
    m.supportedPagesPerSheet = value
}
func (m *PrinterCapabilities) SetSupportedPresentationDirections(value []PrintPresentationDirection)() {
    m.supportedPresentationDirections = value
}
func (m *PrinterCapabilities) SetSupportedPrintQualities(value []PrintQuality)() {
    m.supportedPrintQualities = value
}
func (m *PrinterCapabilities) SetSupportsFitPdfToPage(value *bool)() {
    m.supportsFitPdfToPage = value
}
func (m *PrinterCapabilities) SetTopMargins(value []int32)() {
    m.topMargins = value
}
