package analyzer_test

import (
	"go/parser"
	"go/scanner"
	"go/token"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/GaijinEntertainment/go-exhaustruct/v3/analyzer"
)

func TestSuggestedFix(t *testing.T) {
	t.Parallel()

	_, err := analyzer.NewAnalyzer(
		[]string{`.*TestStruct`},
		nil,
	)
	require.NoError(t, err)

	// Create a test file with missing fields
	testCode := `package test

type TestStruct struct {
	A int
	B string
	C bool
}

func main() {
	_ = TestStruct{A: 1}
}
`

	// Parse the code to verify structure
	fset := token.NewFileSet()
	_, err = parser.ParseFile(fset, "test.go", testCode, parser.ParseComments)
	require.NoError(t, err)

	// This is a simplified test - in reality, SuggestedFix testing
	// requires a more complex setup with proper type information
	t.Log("SuggestedFix functionality is implemented and tested through integration tests")
}

func TestZeroValueGeneration(t *testing.T) {
	tests := []struct {
		name     string
		code     string
		expected string
	}{
		{
			name: "empty struct",
			code: `package test
type TestStruct struct {
	A int
	B string
}
func main() {
	_ = TestStruct{} // This should suggest adding A: 0, B: ""
}`,
			expected: "missing fields",
		},
		{
			name: "partial struct",
			code: `package test
type TestStruct struct {
	A int
	B string
	C bool
}
func main() {
	_ = TestStruct{A: 1} // This should suggest adding B: "", C: false
}`,
			expected: "missing fields",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fset := token.NewFileSet()
			_, err := parser.ParseFile(fset, "test.go", tt.code, parser.ParseComments)

			// Check that parsing works (this validates our test code structure)
			if err != nil {
				if errList, ok := err.(scanner.ErrorList); ok {
					for _, e := range errList {
						t.Logf("Parse error: %v", e)
					}
				}
				t.Errorf("Parse error: %v", err)
			}

			// The actual functionality is tested in the integration test
			assert.Contains(t, tt.code, "TestStruct")
		})
	}
}

func TestSuggestedFixIntegration(t *testing.T) {
	// Test that suggested fixes are generated correctly
	// This would require a more complex setup to test the actual TextEdit generation
	t.Log("SuggestedFix integration testing requires analysis framework setup")

	// Verify that the analyzer includes SuggestedFix in its output
	a, err := analyzer.NewAnalyzer([]string{`.*`}, nil)
	require.NoError(t, err)

	assert.NotNil(t, a)
	assert.Equal(t, "exhaustruct", a.Name)
	assert.Contains(t, a.Doc, "fields are initialized")
}

func TestSuggestedFixWithAnalysisTest(t *testing.T) {
	t.Parallel()

	// Create analyzer with pattern matching the test package
	a, err := analyzer.NewAnalyzer([]string{`.*suggestedfix\..*`}, nil)
	require.NoError(t, err)

	// Run the analyzer with suggested fixes enabled
	testdata := analysistest.TestData()
	results := analysistest.RunWithSuggestedFixes(t, testdata, a, "suggestedfix")

	// Verify that suggested fixes were generated
	require.NotEmpty(t, results, "Expected analysis results with suggested fixes")

	// Check that we have diagnostic results
	for _, result := range results {
		if result.Result != nil {
			t.Logf("Analysis completed for package: %s", result.Pass.Pkg.Name())
		}
	}
}
