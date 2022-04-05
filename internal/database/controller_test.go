package database

import (
	"assignment-2/internal/constants"
	"assignment-2/internal/structs"
	"cloud.google.com/go/firestore"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	InitDB("../../auth.json")
	defer CloseFirestore()

	constants.SetTestDB()

	m.Run()
}

func TestInitDB(t *testing.T) {
	assert.NotNil(t, ctx)
	assert.NotNil(t, app)
	assert.NotEqual(t, firestore.Client{}, client)
}

func TestGetDocumentInvalidObjectDoc(t *testing.T) {
	testStruct := structs.PolicyInputFromApi{}

	GetDocument(constants.PolicyDBCollection, "swe2022-10-05", &testStruct)

	assert.Equal(t, structs.PolicyInputFromApi{}, testStruct)
}

func TestGetDocumentNoCollectionOrDoc(t *testing.T) {
	testStruct := structs.CountryStruct{}

	GetDocument("", "", &testStruct)

	assert.Equal(t, structs.CountryStruct{}, testStruct)
}

func TestGetDocumentByHash(t *testing.T) {
	testStruct := structs.WebHookPost{}

	expected := structs.WebHookPost{
		Country: "Sweden",
		Calls:   2,
	}

	GetDocument(constants.WebhookDbCollection,
		"10911bd27492a5be7c1c772c8528f6f207f7da1b35c727669235f74c93e860e2", &testStruct)

	assert.Equal(t, expected, testStruct)
}

func TestGetDocumentCca3(t *testing.T) {
	testStruct := structs.CountryInfo{}

	expected := structs.CountryInfo{
		Common: "Netherlands",
	}

	GetDocument(constants.CountryDbCollection,
		"nld", &testStruct)

	errDel := DeleteDocument(constants.CountryDbCollection, "Netherlands")
	assert.Nil(t, errDel)

	assert.Equal(t, expected.Common, testStruct.Common)
}

func TestGetAllWebhooks(t *testing.T) {
	time.Sleep(time.Second * 3)
	webhooks, err := GetAllWebhooks(constants.WebhookDbCollection, "")

	assert.Nil(t, err)
	assert.Equal(t, 4, len(webhooks))
}

func TestWriteDocument(t *testing.T) {
	toWrite := structs.CountryInfo{
		Common: "Norway",
	}

	err := WriteDocument(constants.CountryDbCollection, "test", toWrite)

	writtenValue := structs.CountryInfo{}
	GetDocument(constants.CountryDbCollection, "test", &writtenValue)

	errDel := DeleteDocument(constants.CountryDbCollection, "test")

	assert.Nil(t, err)
	assert.Nil(t, errDel)
	assert.Equal(t, toWrite.Common, writtenValue.Common)
}

func TestWriteDocumentOverWrite(t *testing.T) {
	toWriteNor := structs.CountryInfo{
		Common: "Norway",
	}

	toWriteSwe := structs.CountryInfo{
		Common: "Sweden",
	}

	errNor := WriteDocument(constants.CountryDbCollection, "test", toWriteNor)

	writtenValueNor := structs.CountryInfo{}
	GetDocument(constants.CountryDbCollection, "test", &writtenValueNor)

	errSwe := WriteDocument(constants.CountryDbCollection, "test", toWriteSwe)

	writtenValueSwe := structs.CountryInfo{}
	GetDocument(constants.CountryDbCollection, "test", &writtenValueSwe)

	errDel := DeleteDocument(constants.CountryDbCollection, "test")

	assert.Nil(t, errNor)
	assert.Nil(t, errSwe)
	assert.Nil(t, errDel)
	assert.Equal(t, toWriteNor.Common, writtenValueNor.Common)
	assert.Equal(t, toWriteSwe.Common, writtenValueSwe.Common)
	assert.NotEqual(t, writtenValueNor.Common, writtenValueSwe.Common)
}

func TestWriteDocumentTimestamp(t *testing.T) {
	toWrite := structs.CountryInfo{
		Common: "Norway",
	}

	err := WriteDocument(constants.CountryDbCollection, "test", toWrite)

	writtenValue := structs.CountryInfo{}
	GetDocument(constants.CountryDbCollection, "test", &writtenValue)

	errDel := DeleteDocument(constants.CountryDbCollection, "test")

	assert.Nil(t, err)
	assert.Nil(t, errDel)
	assert.NotEqual(t, time.Time{}, writtenValue.TimeStamp)
}

func TestDeleteDocumentWithHash(t *testing.T) {
	toWrite := structs.CountryInfo{
		Common: "Norway",
	}

	err := WriteDocument(constants.CountryDbCollection, "test", toWrite)

	errDel := DeleteDocument(constants.CountryDbCollection,
		"3f7933a16a80051e5e219228f7514fc9fa2679c8bad76e1be32f11b488e05dd9")

	writtenValue := structs.CountryInfo{}
	GetDocument(constants.CountryDbCollection, "test", &writtenValue)

	assert.Nil(t, err)
	assert.Nil(t, errDel)
	assert.Equal(t, structs.CountryInfo{}, writtenValue)
}

func TestIncrementCounterNewCounter(t *testing.T) {
	expectedValue := structs.CountryCounter{Count: 1}
	IncrementCounter(constants.CounterDbCollection, "test")

	actualValue := structs.CountryCounter{}
	GetDocument(constants.CounterDbCollection, "test", &actualValue)

	errDel := DeleteDocument(constants.CounterDbCollection, "test")

	assert.Nil(t, errDel)
	assert.Equal(t, expectedValue, actualValue)
}

func TestIncrementCounter(t *testing.T) {
	writtenValue := structs.CountryCounter{Count: 1}
	err := WriteDocument(constants.CounterDbCollection, "test", &writtenValue)

	IncrementCounter(constants.CounterDbCollection, "test")

	actualValue := structs.CountryCounter{}
	GetDocument(constants.CounterDbCollection, "test", &actualValue)

	errDel := DeleteDocument(constants.CounterDbCollection, "test")

	assert.Nil(t, err)
	assert.Nil(t, errDel)
	assert.Equal(t, 2, actualValue.Count)
}
