// Code generated by "mapstructure-to-hcl2 -type Config"; DO NOT EDIT.
package googlecomputeimport

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName      *string           `mapstructure:"packer_build_name" cty:"packer_build_name"`
	PackerBuilderType    *string           `mapstructure:"packer_builder_type" cty:"packer_builder_type"`
	PackerDebug          *bool             `mapstructure:"packer_debug" cty:"packer_debug"`
	PackerForce          *bool             `mapstructure:"packer_force" cty:"packer_force"`
	PackerOnError        *string           `mapstructure:"packer_on_error" cty:"packer_on_error"`
	PackerUserVars       map[string]string `mapstructure:"packer_user_variables" cty:"packer_user_variables"`
	PackerSensitiveVars  []string          `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables"`
	AccountFile          *string           `mapstructure:"account_file" cty:"account_file"`
	ProjectId            *string           `mapstructure:"project_id" cty:"project_id"`
	IAP                  *bool             `mapstructure:"iap" cty:"iap"`
	Bucket               *string           `mapstructure:"bucket" cty:"bucket"`
	GCSObjectName        *string           `mapstructure:"gcs_object_name" cty:"gcs_object_name"`
	ImageDescription     *string           `mapstructure:"image_description" cty:"image_description"`
	ImageFamily          *string           `mapstructure:"image_family" cty:"image_family"`
	ImageGuestOsFeatures []string          `mapstructure:"image_guest_os_features" cty:"image_guest_os_features"`
	ImageLabels          map[string]string `mapstructure:"image_labels" cty:"image_labels"`
	ImageName            *string           `mapstructure:"image_name" cty:"image_name"`
	SkipClean            *bool             `mapstructure:"skip_clean" cty:"skip_clean"`
	VaultGCPOauthEngine  *string           `mapstructure:"vault_gcp_oauth_engine" cty:"vault_gcp_oauth_engine"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatConfig)
}

// HCL2Spec returns the hcl spec of a Config.
// This spec is used by HCL to read the fields of Config.
// The decoded values from this spec will then be applied to a FlatConfig.
func (*FlatConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"packer_build_name":          &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":        &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_debug":               &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":               &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":            &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":      &hcldec.AttrSpec{Name: "packer_user_variables", Type: cty.Map(cty.String), Required: false},
		"packer_sensitive_variables": &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"account_file":               &hcldec.AttrSpec{Name: "account_file", Type: cty.String, Required: false},
		"project_id":                 &hcldec.AttrSpec{Name: "project_id", Type: cty.String, Required: false},
		"iap":                        &hcldec.AttrSpec{Name: "iap", Type: cty.Bool, Required: false},
		"bucket":                     &hcldec.AttrSpec{Name: "bucket", Type: cty.String, Required: false},
		"gcs_object_name":            &hcldec.AttrSpec{Name: "gcs_object_name", Type: cty.String, Required: false},
		"image_description":          &hcldec.AttrSpec{Name: "image_description", Type: cty.String, Required: false},
		"image_family":               &hcldec.AttrSpec{Name: "image_family", Type: cty.String, Required: false},
		"image_guest_os_features":    &hcldec.AttrSpec{Name: "image_guest_os_features", Type: cty.List(cty.String), Required: false},
		"image_labels":               &hcldec.AttrSpec{Name: "image_labels", Type: cty.Map(cty.String), Required: false},
		"image_name":                 &hcldec.AttrSpec{Name: "image_name", Type: cty.String, Required: false},
		"skip_clean":                 &hcldec.AttrSpec{Name: "skip_clean", Type: cty.Bool, Required: false},
		"vault_gcp_oauth_engine":     &hcldec.AttrSpec{Name: "vault_gcp_oauth_engine", Type: cty.String, Required: false},
	}
	return s
}
