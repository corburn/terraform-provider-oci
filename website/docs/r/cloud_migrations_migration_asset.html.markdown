---
subcategory: "Cloud Migrations"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_migrations_migration_asset"
sidebar_current: "docs-oci-resource-cloud_migrations-migration_asset"
description: |-
  Provides the Migration Asset resource in Oracle Cloud Infrastructure Cloud Migrations service
---

# oci_cloud_migrations_migration_asset
This resource provides the Migration Asset resource in Oracle Cloud Infrastructure Cloud Migrations service.

Creates a migration asset.


## Example Usage

```hcl
resource "oci_cloud_migrations_migration_asset" "test_migration_asset" {
	#Required
	availability_domain = var.migration_asset_availability_domain
	inventory_asset_id = oci_cloud_migrations_inventory_asset.test_inventory_asset.id
	migration_id = oci_cloud_migrations_migration.test_migration.id
	replication_compartment_id = oci_identity_compartment.test_compartment.id
	snap_shot_bucket_name = oci_objectstorage_bucket.test_bucket.name

	#Optional
	depends_on = var.migration_asset_depends_on
	display_name = var.migration_asset_display_name
	replication_schedule_id = oci_cloud_migrations_replication_schedule.test_replication_schedule.id
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Required) Availability domain
* `depends_on` - (Optional) (Updatable) List of migration assets that depends on this asset.
* `display_name` - (Optional) (Updatable) A user-friendly name. If empty, then source asset name will be used. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `inventory_asset_id` - (Required) OCID of an asset for an inventory.
* `migration_id` - (Required) OCID of the associated migration.
* `replication_compartment_id` - (Required) Replication compartment identifier
* `replication_schedule_id` - (Optional) (Updatable) Replication schedule identifier
* `snap_shot_bucket_name` - (Required) Name of snapshot bucket


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `availability_domain` - Availability domain
* `compartment_id` - Compartment Identifier
* `depended_on_by` - List of migration assets that depend on the asset.
* `depends_on` - List of migration assets that depends on the asset.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `id` - Asset ID generated by mirgration service. It is used in the mirgration service pipeline.
* `lifecycle_details` - A message describing the current state in more detail. For example, it can be used to provide actionable information for a resource in Failed state.
* `migration_id` - OCID of the associated migration.
* `notifications` - List of notifications
* `parent_snapshot` - The parent snapshot of the migration asset to be used by the replication task.
* `replication_compartment_id` - Replication compartment identifier
* `replication_schedule_id` - Replication schedule identifier
* `snap_shot_bucket_name` - Name of snapshot bucket
* `snapshots` - Key-value pair representing disks ID mapped to the OCIDs of replicated or hydration server volume snapshots. Example: `{"bar-key": "value"}` 
	* `unmodified_volume_id` - ID of the unmodified volume
	* `uuid` - ID of the vCenter disk obtained from Inventory.
	* `volume_id` - ID of the hydration server volume
	* `volume_type` - The hydration server volume type
* `source_asset_data` - Key-value pair representing asset metadata keys and values scoped to a namespace. Example: `{"bar-key": "value"}` 
* `source_asset_id` - OCID that is referenced to an asset for an inventory.
* `state` - The current state of the migration asset.
* `tenancy_id` - Tenancy identifier
* `time_created` - The time when the migration asset was created. An RFC3339 formatted datetime string.
* `time_updated` - The time when the migration asset was updated. An RFC3339 formatted datetime string.
* `type` - The type of asset referenced for inventory.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Migration Asset
	* `update` - (Defaults to 20 minutes), when updating the Migration Asset
	* `delete` - (Defaults to 20 minutes), when destroying the Migration Asset


## Import

MigrationAssets can be imported using the `id`, e.g.

```
$ terraform import oci_cloud_migrations_migration_asset.test_migration_asset "id"
```

