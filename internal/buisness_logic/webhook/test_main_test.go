package webhook

import (
	"assignment-2/internal/constants"
	"assignment-2/internal/database"
	"testing"
)

func TestMain(m *testing.M) {
	database.InitDB("../../../auth.json")
	defer database.CloseFirestore()

	constants.SetTestDB()

	m.Run()
}
