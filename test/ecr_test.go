package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformEcrExample(t *testing.T) {
	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/ecr",
		Vars: map[string]interface{}{
			"image_name":"image",
		},
	})

	terraform.InitAndApply(t, terraformOptions)

	defer terraform.Destroy(t, terraformOptions)

	// ECR Test
	ecr_name := terraform.Output(t, terraformOptions, "ecr_name")
	assert.Equal(t, "image", ecr_name)
}