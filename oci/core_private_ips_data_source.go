// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v31/core"
)

func init() {
	RegisterDatasource("oci_core_private_ips", CorePrivateIpsDataSource())
}

func CorePrivateIpsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCorePrivateIps,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"ip_address": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vlan_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vnic_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"private_ips": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(CorePrivateIpResource()),
			},
		},
	}
}

func readCorePrivateIps(d *schema.ResourceData, m interface{}) error {
	sync := &CorePrivateIpsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient()

	return ReadResource(sync)
}

type CorePrivateIpsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.ListPrivateIpsResponse
}

func (s *CorePrivateIpsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CorePrivateIpsDataSourceCrud) Get() error {
	request := oci_core.ListPrivateIpsRequest{}

	if ipAddress, ok := s.D.GetOkExists("ip_address"); ok {
		tmp := ipAddress.(string)
		request.IpAddress = &tmp
	}

	if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
		tmp := subnetId.(string)
		request.SubnetId = &tmp
	}

	if vlanId, ok := s.D.GetOkExists("vlan_id"); ok {
		tmp := vlanId.(string)
		request.VlanId = &tmp
	}

	if vnicId, ok := s.D.GetOkExists("vnic_id"); ok {
		tmp := vnicId.(string)
		request.VnicId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

	response, err := s.Client.ListPrivateIps(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListPrivateIps(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CorePrivateIpsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("CorePrivateIpsDataSource-", CorePrivateIpsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		privateIp := map[string]interface{}{}

		if r.AvailabilityDomain != nil {
			privateIp["availability_domain"] = *r.AvailabilityDomain
		}

		if r.CompartmentId != nil {
			privateIp["compartment_id"] = *r.CompartmentId
		}

		if r.DefinedTags != nil {
			privateIp["defined_tags"] = definedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			privateIp["display_name"] = *r.DisplayName
		}

		privateIp["freeform_tags"] = r.FreeformTags

		if r.HostnameLabel != nil {
			privateIp["hostname_label"] = *r.HostnameLabel
		}

		if r.Id != nil {
			privateIp["id"] = *r.Id
		}

		if r.IpAddress != nil {
			privateIp["ip_address"] = *r.IpAddress
		}

		if r.IsPrimary != nil {
			privateIp["is_primary"] = *r.IsPrimary
		}

		if r.SubnetId != nil {
			privateIp["subnet_id"] = *r.SubnetId
		}

		if r.TimeCreated != nil {
			privateIp["time_created"] = r.TimeCreated.String()
		}

		if r.VlanId != nil {
			privateIp["vlan_id"] = *r.VlanId
		}

		if r.VnicId != nil {
			privateIp["vnic_id"] = *r.VnicId
		}

		resources = append(resources, privateIp)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, CorePrivateIpsDataSource().Schema["private_ips"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("private_ips", resources); err != nil {
		return err
	}

	return nil
}
