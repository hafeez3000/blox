// Copyright 2016 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the License). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the license file accompanying this file. This file is distributed
// on an AS IS BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package e2etasksteps

import (
	"github.com/aws/amazon-ecs-event-stream-handler/internal/models"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/pkg/errors"
)

func ValidateTasksMatch(ecsTask ecs.Task, eshTask models.TaskModel) error {
	if *ecsTask.TaskArn != *eshTask.TaskARN ||
		*ecsTask.ClusterArn != *eshTask.ClusterARN ||
		*ecsTask.ContainerInstanceArn != *eshTask.ContainerInstanceARN {
		return errors.New("Tasks don't match")
	}
	return nil
}

func ValidateListContainsTask(ecsTask ecs.Task, eshTaskList []models.TaskModel) error {
	taskARN := *ecsTask.TaskArn
	var eshTask models.TaskModel
	for _, t := range eshTaskList {
		if *t.TaskARN == taskARN {
			eshTask = t
			break
		}
	}
	if eshTask.TaskARN == nil {
		return errors.Errorf("Task with ARN '%s' not found in response", taskARN)
	}
	return ValidateTasksMatch(ecsTask, eshTask)
}