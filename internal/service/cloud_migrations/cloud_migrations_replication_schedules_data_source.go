// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package cloud_migrations

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_cloud_migrations "github.com/oracle/oci-go-sdk/v65/cloudmigrations"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CloudMigrationsReplicationSchedulesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCloudMigrationsReplicationSchedules,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"replication_schedule_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"replication_schedule_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(CloudMigrationsReplicationScheduleResource()),
						},
					},
				},
			},
		},
	}
}

func readCloudMigrationsReplicationSchedules(d *schema.ResourceData, m interface{}) error {
	sync := &CloudMigrationsReplicationSchedulesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MigrationClient()

	return tfresource.ReadResource(sync)
}

type CloudMigrationsReplicationSchedulesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_cloud_migrations.MigrationClient
	Res    *oci_cloud_migrations.ListReplicationSchedulesResponse
}

func (s *CloudMigrationsReplicationSchedulesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CloudMigrationsReplicationSchedulesDataSourceCrud) Get() error {
	request := oci_cloud_migrations.ListReplicationSchedulesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if replicationScheduleId, ok := s.D.GetOkExists("id"); ok {
		tmp := replicationScheduleId.(string)
		request.ReplicationScheduleId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_cloud_migrations.ReplicationScheduleLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "cloud_migrations")

	response, err := s.Client.ListReplicationSchedules(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListReplicationSchedules(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CloudMigrationsReplicationSchedulesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CloudMigrationsReplicationSchedulesDataSource-", CloudMigrationsReplicationSchedulesDataSource(), s.D))
	resources := []map[string]interface{}{}
	replicationSchedule := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ReplicationScheduleSummaryToMap(item))
	}
	replicationSchedule["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, CloudMigrationsReplicationSchedulesDataSource().Schema["replication_schedule_collection"].Elem.(*schema.Resource).Schema)
		replicationSchedule["items"] = items
	}

	resources = append(resources, replicationSchedule)
	if err := s.D.Set("replication_schedule_collection", resources); err != nil {
		return err
	}

	return nil
}
