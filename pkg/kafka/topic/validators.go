package topic

import (
	"context"
	"errors"
	"fmt"
	"regexp"

	"strconv"

	strimziadminclient "github.com/redhat-developer/app-services-cli/pkg/api/strimzi-admin/client"
	"github.com/redhat-developer/app-services-cli/pkg/common/commonerr"
)

const (
	legalNameChars       = "^[a-zA-Z0-9\\_\\-]+$"
	maxNameLength        = 249
	minReplicationFactor = 1
	minPartitions        = 1
	maxPartitions        = 100
)

// ValidateName validates the name of the topic
func ValidateName(val interface{}) error {
	name, ok := val.(string)
	if !ok {
		return commonerr.NewCastError(val, "string")
	}

	if len(name) < 1 {
		return errors.New("topic name is required")
	} else if len(name) > maxNameLength {
		return fmt.Errorf("topic name cannot exceed %v characters", maxNameLength)
	}

	matched, _ := regexp.Match(legalNameChars, []byte(name))

	if matched {
		return nil
	}

	return fmt.Errorf(`invalid topic name "%v"; only letters (Aa-Zz), numbers, "_" and "-" are accepted`, name)
}

// ValidatePartitionsN performs validation on the number of partitions v
func ValidatePartitionsN(v interface{}) error {
	partitionsStr := fmt.Sprintf("%v", v)

	partitions, err := strconv.Atoi(partitionsStr)
	if err != nil {
		return commonerr.NewCastError(v, "int32")
	}

	if partitions < minPartitions {
		return fmt.Errorf("invalid partition count %v, minimum value is %v", partitions, minPartitions)
	}

	if partitions > maxPartitions {
		return fmt.Errorf("invalid partition count %v, maximum value is %v", partitions, maxPartitions)
	}

	return nil
}

// ValidationReplicationFactorN performs validation on the number of replicas v
func ValidateReplicationFactorN(v interface{}) error {
	replicas, ok := v.(int32)
	if !ok {
		return commonerr.NewCastError(v, "int32")
	}

	if replicas < minReplicationFactor {
		return fmt.Errorf("invalid replication factor %v, minimum value is %v", replicas, minReplicationFactor)
	}

	return nil
}

// ValidateMessageRetentionPeriod validates the value (ms) of the retention period
// the valid values can range from [-1,...]
func ValidateMessageRetentionPeriod(v interface{}) error {
	retentionPeriodMsStr := fmt.Sprintf("%v", v)

	if retentionPeriodMsStr == "" {
		return nil
	}

	retentionPeriodMs, err := strconv.Atoi(retentionPeriodMsStr)
	if err != nil {
		return commonerr.NewCastError(v, "int")
	}

	if retentionPeriodMs < -1 {
		return fmt.Errorf("invalid retention period %v, minimum value is -1", retentionPeriodMs)
	}

	return nil
}

// ValidateMessageRetentionSize validates the value (bytes) of the retention size
// the valid values can range from [-1,...]
func ValidateMessageRetentionSize(v interface{}) error {
	retentionSizeStr := fmt.Sprintf("%v", v)

	if retentionSizeStr == "" {
		return nil
	}

	retentionPeriodBytes, err := strconv.Atoi(retentionSizeStr)
	if err != nil {
		return commonerr.NewCastError(v, "int")
	}

	if retentionPeriodBytes < -1 {
		return fmt.Errorf("invalid retention size %v, minimum value is -1", retentionPeriodBytes)
	}

	return nil
}

// ValidateNameIsAvailable checks if a topic with the given name already exists
func ValidateNameIsAvailable(api strimziadminclient.DefaultApi, instance string) func(v interface{}) error {
	return func(v interface{}) error {
		name, _ := v.(string)

		_, httpRes, _ := api.GetTopic(context.Background(), name).Execute()

		if httpRes != nil && httpRes.StatusCode == 200 {
			return fmt.Errorf(`topic "%v" already exists in Kafka instance "%v"`, name, instance)
		}

		return nil
	}
}
