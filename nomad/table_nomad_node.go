package nomad

import (
	"context"

	"github.com/nomad/nomad-sdk-go/nomad/nomaderr"
	"github.com/nomad/nomad-sdk-go/service/iam"
	"github.com/nomad/nomad-sdk-go/service/organizations"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

//// TABLE DEFINITION

func tableNomadNode(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "nomad_node",
		Description: "Nomad Account",
		List: &plugin.ListConfig{
			Hydrate: listNodeAlias,
		},
		Columns: []*plugin.Column{
			{
				Name:        "node_aliases",
				Description: "A list of aliases associated with the node, if applicable.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Aliases"),
			},
			{
				Name:        "organization_id",
				Description: "The unique identifier (ID) of an organization, if applicable.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getOrganizationDetails,
				Transform:   transform.FromField("Organization.Id"),
			},
			{
				Name:        "organization_arn",
				Description: "The Amazon Resource Name (ARN) of an organization.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getOrganizationDetails,
				Transform:   transform.FromField("Organization.Arn"),
			},
			{
				Name:        "organization_feature_set",
				Description: "Specifies the functionality that currently is available to the organization. If set to \"ALL\", then all features are enabled and policies can be applied to nodes in the organization. If set to \"CONSOLIDATED_BILLING\", then only consolidated billing functionality is available.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getOrganizationDetails,
				Transform:   transform.FromField("Organization.FeatureSet"),
			},
			{
				Name:        "organization_master_node_arn",
				Description: "The Amazon Resource Name (ARN) of the node that is designated as the management node for the organization",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getOrganizationDetails,
				Transform:   transform.FromField("Organization.MasterNodeArn"),
			},
			{
				Name:        "organization_master_node_email",
				Description: "The email address that is associated with the Nomad node that is designated as the management node for the organization",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getOrganizationDetails,
				Transform:   transform.FromField("Organization.MasterNodeEmail"),
			},
			{
				Name:        "organization_master_node_id",
				Description: "The unique identifier (ID) of the management node of an organization",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getOrganizationDetails,
				Transform:   transform.FromField("Organization.MasterNodeId"),
			},
			{
				Name:        "organization_available_policy_types",
				Description: "The Region opt-in status. The possible values are opt-in-not-required, opted-in, and not-opted-in",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getOrganizationDetails,
				Transform:   transform.FromField("Organization.AvailablePolicyTypes"),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(nodeDataToTitle),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(nodeDataToAkas),
			},
		},
	}
}

type nodeData struct {
	commonColumnData nomadCommonColumnData
	Aliases          []*string
}

//// LIST FUNCTION

func listAccountAlias(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	// Create Session
	svc, err := IAMService(ctx, d)
	if err != nil {
		return nil, err
	}

	commonData, err := getCommonColumns(ctx, d, h)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*nomadCommonColumnData)

	// execute list call
	op, err := svc.ListAccountAliases(&iam.ListAccountAliasesInput{})
	if err != nil {
		return nil, err
	}

	if op.AccountAliases != nil {
		d.StreamListItem(ctx, &nodeData{
			commonColumnData: *commonColumnData,
			Aliases:          op.AccountAliases,
		})
		return nil, nil
	}

	d.StreamListItem(ctx, &nodeData{
		commonColumnData: *commonColumnData,
	})

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getOrganizationDetails(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getOrganizationDetails")

	// Create Session
	svc, err := OrganizationService(ctx, d)
	if err != nil {
		return nil, err
	}

	op, err := svc.DescribeOrganization(&organizations.DescribeOrganizationInput{})
	if err != nil {
		if a, ok := err.(nomaderr.Error); ok {
			if a.Code() == "NomadOrganizationsNotInUseException" {
				return nil, nil
			}
		}
		return nil, err
	}

	return op, nil
}

//// Transform Functions

func nodeDataToTitle(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getAwsAccountAkas")
	nodeInfo := d.HydrateItem.(*nodeData)

	if nodeInfo.Aliases != nil && len(nodeInfo.Aliases) > 0 {
		return nodeInfo.Aliases[0], nil
	}

	return nodeInfo.commonColumnData.AccountId, nil
}

func nodeDataToAkas(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	plugin.Logger(ctx).Trace("nodeDataToTitle")
	nodeInfo := d.HydrateItem.(*nodeData)

	akas := []string{"arn:" + nodeInfo.commonColumnData.Partition + ":::" + nodeInfo.commonColumnData.AccountId}

	return akas, nil
}
