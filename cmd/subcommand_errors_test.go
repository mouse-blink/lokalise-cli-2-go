package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
	"strings"
	"testing"
)

func TestBranchDelete_Error(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(
		fmt.Sprintf("/api2/projects/%s/branches/42", testProjectID),
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			testMethod(t, r, "DELETE")
			testHeader(t, r, "X-Api-Token", testApiToken)

			w.WriteHeader(http.StatusNotFound)
			_, _ = fmt.Fprint(w, `{"error": {"code": 404, "message": "Branch not found"}}`)
		})

	args := []string{"branch", "delete", "--branch-id=42", "--project-id=" + testProjectID}
	rootCmd.SetArgs(args)
	branchDeleteCmd.PreRun = func(cmd *cobra.Command, args []string) {
		Api = client
	}

	cmdErr := rootCmd.Execute()

	if cmdErr == nil || !strings.Contains(cmdErr.Error(), "Branch not found") {
		t.Errorf("Expected the API error 'Branch not found', but got %v", cmdErr)
	}
}
