---
subcategory: "ECS"
layout: "aws"
page_title: "AWS: aws_ecs_capacity_provider"
description: |-
  Provides an ECS Capacity Provider.
---

# Resource: aws_ecs_capacity_provider

Provides an ECS Capacity Provider.

## Example Usage

```hcl
resource "aws_ecs_capacity_provider" "example" {
  name = "example"

  auto_scaling_group_provider {
    auto_scaling_group_arn = aws_autoscaling_group.example.arn
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the capacity provider (up to 255 letters, numbers, hyphens, and underscores).
* `tags` - (Optional) Key-value mapping of resource tags.
* `auto_scaling_group_provider` - (Required) Configuration block describing the auto scaling group.

## auto_scaling_group_provider

The `auto_scaling_group_provider` configuration block supports the following:

* `auto_scaling_group_arn` - (Required) The Amazon Resource Name (ARN) of the auto scaling group.
* `managed_scaling` - (Optional) Configuration block describing the auto scaling group scaling policies.
* `managed_termination_protection` -  (Optional) Whether to enable scale-in protection. Valid values are `ENABLED` and `DISABLED`.

## managed_scaling

The `managed_scaling` configuration block supports the following:

* `maximum_scaling_step_size` - (Optional) The maximum number of container instances that ECS will scale in or out at a time. A number between 1 and 10,000.
* `minimum_scaling_step_size` - (Optional) The minimum number of container instances that ECS will scale in or out at a time. A number between 1 and 10,000.
* `status` - (Optional) Whether to enable managed scaling. Valid values are `ENABLED` and `DISABLED`.
* `target_capacity` - (Optional) The target capacity for the capacity provider. A number between 1 and 100.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The Amazon Resource Name (ARN) that identifies the capacity provider.

## Import

ECS clusters can be imported using the `name`, e.g.

```hcl
$ terraform import aws_ecs_capacity_provider.example example
```
