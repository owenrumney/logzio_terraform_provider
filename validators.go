package main

import (
	"github.com/jonboydell/logzio_client"
	"fmt"
)

func contains(slice []string, s string) bool {

	for _, value := range slice {
		if value == s {
			return true
		}
	}
	return false
}

func validateOperation(v interface{}, k string) (ws []string, errors []error) {

	value := v.(string)

	validOperations := []string{
		logzio_client.OperatorGreaterThanOrEquals,
		logzio_client.OperatorGreaterThan,
		logzio_client.OperatorEquals,
		logzio_client.OperatorLessThan,
		logzio_client.OperatorLessThanOrEquals,
		logzio_client.OperatorNotEquals,
	}

	if len(value) == 0 {
		errors = append(errors, fmt.Errorf("Operation %q must not be blank and be one of %s", k, validOperations))
	}

	valid := false
	for _, op := range validOperations {
		if op == value {
			valid = true
		}
	}

	if !valid {
		errors = append(errors, fmt.Errorf("Operation %q must be one of %s", k, validOperations))
	}
	return
}

func validAggregationTypes(v interface{}, k string) (ws []string, errors []error) {

	value := v.(string)

	validAggregationTypes := []string{logzio_client.AggregationTypeUniqueCount,
	logzio_client.Avg, logzio_client.AggregationTypeMaz, logzio_client.AggregationTypeNone,
	logzio_client.AggregationTypeSum, logzio_client.AggregationTypeCount, logzio_client.AggregationTypeMin}

	if !contains(validAggregationTypes, value) {
		errors = append(errors, fmt.Errorf("valueAggregationType %q must be one of %s", k, validAggregationTypes))
	}
	return
}