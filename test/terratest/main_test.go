package test

import (
	"os"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformInitAndValidate(t *testing.T) {
	t.Parallel()

	// Path to the Terraform code that will be tested
	terraformOptions := &terraform.Options{
		TerraformDir: "../",
	}

	// Run `terraform init` and `terraform validate` to initialize the working directory and validate the Terraform configuration
	terraform.InitAndValidate(t, terraformOptions)

	// Clean up by running `terraform destroy`
	defer terraform.Destroy(t, terraformOptions)

	// Validate the Terraform configuration using `terraform validate`
	terraform.Validate(t, terraformOptions)
}

func TestTerraformPlan(t *testing.T) {
	t.Parallel()

	// Path to the Terraform code that will be tested
	terraformOptions := &terraform.Options{
		// Set the path to the Terraform code that will be tested.
		TerraformDir: "../",
	}

	// Run `terraform init` and `terraform plan`. Fail the test if there are any errors.
	terraform.InitAndPlan(t, terraformOptions)

	// Validate that the plan does not contain any changes
	changes := terraform.Plan(t, terraformOptions)
	assert.Empty(t, changes)

}

// This code will run when the test suite is being cleaned up
func TestTerraformPlanCleanup(t *testing.T) {
	t.Parallel()

	// Path to the Terraform code that will be tested
	terraformOptions := &terraform.Options{
		// Set the path to the Terraform code that will be tested.
		TerraformDir: "../",
	}

	// Run `terraform init` and `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// Validate that there are no changes to apply.
	terraform.InitAndPlan(t, terraformOptions)
	changes := terraform.Plan(t, terraformOptions)
	assert.Empty(t, changes)
}

// This function will be called to run the test
func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	os.Exit(m.Run())
}
