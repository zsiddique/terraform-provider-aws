---
subcategory: "ECS"
layout: "aws"
page_title: "AWS: aws_ecs_cluster"
description: |-
  Provides an ECS cluster.
---

# Resource: aws_ecs_cluster

Provides an ECS cluster.

## Example Usage

```hcl
resource "aws_ecs_cluster" "foo" {
  name = "white-hart"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the cluster (up to 255 letters, numbers, hyphens, and underscores)
* `tags` - (Optional) Key-value mapping of resource tags
* `setting` - (Optional) Configuration block(s) with cluster settings. For example, this can be used to enable CloudWatch Container Insights for a cluster. Defined below.
* `capacity_providers` - (Optional) Configuration blocks containing capacity provider configurations. Defined below.
* `default_capacity_provider_strategy` - below
 
## setting

The `setting` configuration block supports the following:

* `name` - (Required) Name of the setting to manage. Valid values: `containerInsights`.
* `value` -  (Required) The value to assign to the setting. Value values are `enabled` and `disabled`.

## capacity_providers

"Only capacity providers using an Auto Scaling group can be created. Amazon ECS tasks on AWS Fargate use the FARGATE and FARGATE_SPOT capacity providers which are already created and available to all accounts in Regions supported by AWS Fargate."

* `autoscaling_group_provider` - below
* `capacity_provider_arn`
* `name`
* `status` - exported?
* `tags`

## autoscaling_group_provider

* `autoscaling_group_arn`
* `managed_scaling` - below
* `managed_termination_protection` - "ENABLED"

## managed_scaling

* maximum_scaling_step_size
* minimum_scaling_step_size
* status - "ENABLED"
* target_capacity

## `default_capacity_provider_strategy`

`default_capacity_provider_strategy` supports the following:

* `base` - (Optional)
* `capacity_provider` - (Required)
* `weight` - (Optional)
 
## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The Amazon Resource Name (ARN) that identifies the cluster
* `arn` - The Amazon Resource Name (ARN) that identifies the cluster

## Import

ECS clusters can be imported using the `name`, e.g.

```
$ terraform import aws_ecs_cluster.stateless stateless-app
```
