package mapping

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestExtractTags tests the ExtractTags function.
func TestExtractTags(t *testing.T) {
	m := testingModelMap(t)

	err := m.RegisterModels(&Model1WithMany2Many{}, &Model2WithMany2Many{}, &JoinModel{})
	require.NoError(t, err)

	first, ok := m.GetModelStruct(&Model1WithMany2Many{})
	require.True(t, ok)

	synced, ok := first.RelationByName("Synced")
	require.True(t, ok)

	extracted := synced.ExtractFieldTags("neuron")
	assert.Len(t, extracted, 3)
}
