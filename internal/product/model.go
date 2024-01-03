package product

import "time"

const PRODUCT_COLLECTION_NAME = "products"

type Product struct {
	RecordReference         string    `json:"recordReference" bson:"recordReference"`
	DeletionText            string    `json:"deletionText" bson:"deletionText"`
	ProductForm             string    `json:"productForm" bson:"productForm"`
	ProductFormDetail       string    `json:"productFormDetail" bson:"productFormDetail"`
	TitleText               string    `json:"titleText" bson:"titleText"`
	TitleTextKana           string    `json:"titleTextKana" bson:"titleTextKana"`
	Subtitle                string    `json:"subtitle" bson:"subtitle"`
	SubtitleKana            string    `json:"subtitleKana" bson:"subtitleKana"`
	NumberOfPages           int       `json:"numberOfPages" bson:"numberOfPages"`
	ExtentValue             int       `json:"extentValue" bson:"extentValue"`
	SubjectHeadingText      string    `json:"subjectHeadingText" bson:"subjectHeadingText"`
	SubjectCode1            string    `json:"subjectCode1" bson:"subjectCode1"`
	SubjectCodeText1        string    `json:"subjectCodeText1" bson:"subjectCodeText1"`
	SubjectCode2            string    `json:"subjectCode2" bson:"subjectCode2"`
	SubjectCodeText2        string    `json:"subjectCodeText2" bson:"subjectCodeText2"`
	AudienceDescription     string    `json:"audienceDescription" bson:"audienceDescription"`
	ContactCompany          string    `json:"contactCompany" bson:"contactCompany"`
	ContactName             string    `json:"contactName" bson:"contactName"`
	ContactEmail            string    `json:"contactEmail" bson:"contactEmail"`
	AnnouncementDate        time.Time `json:"announcementDate" bson:"announcementDate"`
	PublicationDate         time.Time `json:"publicationDate" bson:"publicationDate"`
	SupplyRestrictionDetail string    `json:"supplyRestrictionDetail" bson:"supplyRestrictionDetail"`
	ExpectedShipDate        time.Time `json:"expectedShipDate" bson:"expectedShipDate"`
	OnSaleDate              time.Time `json:"onSaleDate" bson:"onSaleDate"`
	UnpricedItemType        string    `json:"unpricedItemType" bson:"unpricedItemType"`
	PriceAmount             float64   `json:"priceAmount" bson:"priceAmount"`
	SpecialPriceAmount      float64   `json:"specialPriceAmount" bson:"specialPriceAmount"`
	PriceEffectiveFrom      time.Time `json:"priceEffectiveFrom" bson:"priceEffectiveFrom"`
	PriceEffectiveUntil     time.Time `json:"priceEffectiveUntil" bson:"priceEffectiveUntil"`
	PublicationFlag         bool      `json:"publicationFlag" bson:"publicationFlag"`
	DeleteFlag              bool      `json:"deleteFlag" bson:"deleteFlag"`
	CreateTime              time.Time `json:"createTime" bson:"createTime"`
	UpdateTime              time.Time `json:"updateTime" bson:"updateTime"`
	DeleteTime              time.Time `json:"deleteTime" bson:"deleteTime"`
}
