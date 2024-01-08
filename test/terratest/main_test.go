package test

import (
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformInitPlanValidateDestroy(t *testing.T) {
	t.Parallel()

	// Specify the path to the Terraform code that will be tested
	terraformOptions := &terraform.Options{
		// Set the path to the Terraform code that will be tested.
		TerraformDir: "../",
	}

	// Terraform Init and Validate
	terraform.InitAndValidate(t, terraformOptions)

	// Terraform Plan
	terraformPlanOutput := terraform.InitAndPlanAndShow(t, terraformOptions)

	// Validate the Terraform Plan output (e.g., check for no changes, specific resources)
	assert.False(t, strings.Contains(terraformPlanOutput, "to be created"), "Unexpected resources to be created in Terraform plan")

	// No apply step for this example

	// Terraform Destroy
	defer terraform.Destroy(t, terraformOptions)

	// Validate the Terraform Destroy output
	terraform.InitAndPlanAndShow(t, terraformOptions)
}
