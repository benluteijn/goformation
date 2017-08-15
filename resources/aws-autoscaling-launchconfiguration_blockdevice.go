package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::AutoScaling::LaunchConfiguration.BlockDevice AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig-blockdev-template.html
type AWSAutoScalingLaunchConfiguration_BlockDevice struct {

	// DeleteOnTermination AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig-blockdev-template.html#cfn-as-launchconfig-blockdev-template-deleteonterm

	DeleteOnTermination bool `json:"DeleteOnTermination"`

	// Encrypted AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig-blockdev-template.html#cfn-as-launchconfig-blockdev-template-encrypted

	Encrypted bool `json:"Encrypted"`

	// Iops AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig-blockdev-template.html#cfn-as-launchconfig-blockdev-template-iops

	Iops int64 `json:"Iops"`

	// SnapshotId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig-blockdev-template.html#cfn-as-launchconfig-blockdev-template-snapshotid

	SnapshotId string `json:"SnapshotId"`

	// VolumeSize AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig-blockdev-template.html#cfn-as-launchconfig-blockdev-template-volumesize

	VolumeSize int64 `json:"VolumeSize"`

	// VolumeType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig-blockdev-template.html#cfn-as-launchconfig-blockdev-template-volumetype

	VolumeType string `json:"VolumeType"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSAutoScalingLaunchConfiguration_BlockDevice) AWSCloudFormationType() string {
	return "AWS::AutoScaling::LaunchConfiguration.BlockDevice"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSAutoScalingLaunchConfiguration_BlockDevice) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSAutoScalingLaunchConfiguration_BlockDeviceResources retrieves all AWSAutoScalingLaunchConfiguration_BlockDevice items from a CloudFormation template
func GetAllAWSAutoScalingLaunchConfiguration_BlockDevice(template *Template) map[string]*AWSAutoScalingLaunchConfiguration_BlockDevice {

	results := map[string]*AWSAutoScalingLaunchConfiguration_BlockDevice{}
	for name, resource := range template.Resources {
		result := &AWSAutoScalingLaunchConfiguration_BlockDevice{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSAutoScalingLaunchConfiguration_BlockDeviceWithName retrieves all AWSAutoScalingLaunchConfiguration_BlockDevice items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSAutoScalingLaunchConfiguration_BlockDevice(name string, template *Template) (*AWSAutoScalingLaunchConfiguration_BlockDevice, error) {

	result := &AWSAutoScalingLaunchConfiguration_BlockDevice{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSAutoScalingLaunchConfiguration_BlockDevice{}, errors.New("resource not found")

}