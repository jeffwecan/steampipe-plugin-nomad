package nomad

import (
	"context"
	"fmt"

	"github.com/hashicorp/nomad/api"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

//// TABLE DEFINITION

// NodeListStub is a subset of information returned during
// node list operations.
// type NodeListStub struct {
// 	Address               string
// 	ID                    string
// 	Datacenter            string
// 	Name                  string
// 	NodeClass             string
// 	Version               string
// 	Drain                 bool
// 	SchedulingEligibility string
// 	Status                string
// 	StatusDescription     string
// 	Drivers               map[string]*DriverInfo
// 	NodeResources         *NodeResources         `json:",omitempty"`
// 	ReservedResources     *NodeReservedResources `json:",omitempty"`
// 	CreateIndex           uint64
// 	ModifyIndex           uint64
// }

func tableNomadNode(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "nomad_node",
		Description: "Nomad Account",
		List: &plugin.ListConfig{
			Hydrate: listNodes,
		},
		// Get: &plugin.GetConfig{
		// 		KeyColumns: plugin.SingleColumn("id"),
		// 		Hydrate:    getNode,
		// },
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Description: "Node ID",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "name",
				Description: "Node name",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
			{
				Name:        "datacenter",
				Description: "Node datacenter",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Datacenter"),
			},
			{
				Name:        "node_class",
				Description: "Node class",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("NodeClass"),
			},
			{
				Name:        "drain",
				Description: "Is node draining",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Drain"),
			},
			{
				Name:        "scheduling_eligibility",
				Description: "Node scheduling eligibility",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("SchedulingEligibility"),
			},
			{
				Name:        "status",
				Description: "Node status",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Status"),
			},
			{
				Name:        "status_description",
				Description: "Node status description",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("StatusDescription"),
			},
			{
				Name:        "modify_index",
				Description: "Node modify index",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ModifyIndex"),
			},
		},
	}
}

// type nodeData struct {
// 	commonColumnData nomadCommonColumnData
// 	Aliases          []*string
// }

//// LIST FUNCTION

func listNodes(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)
	// Create Session
	client, err := NomadClient(ctx, d, )
	if err != nil {
		return nil, err
	}

	queryOpts := &api.QueryOptions{}
	nodesResp, _, err := client.Nodes().List(queryOpts)
	if err != nil {
		return nil, err
	}

	for _, node := range nodesResp {
		logger.Info(fmt.Sprintf("node: %+v", node))
    d.StreamListItem(ctx, node)
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

// func getOrganizationDetails(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
// 	plugin.Logger(ctx).Trace("getOrganizationDetails")

// 	// Create Session
// 	svc, err := OrganizationService(ctx, d)
// 	if err != nil {
// 		return nil, err
// 	}

// 	op, err := svc.DescribeOrganization(&organizations.DescribeOrganizationInput{})
// 	if err != nil {
// 		if a, ok := err.(nomaderr.Error); ok {
// 			if a.Code() == "NomadOrganizationsNotInUseException" {
// 				return nil, nil
// 			}
// 		}
// 		return nil, err
// 	}

// 	return op, nil
// }

//// Transform Functions

// func nodeDataToTitle(ctx context.Context, d *transform.TransformData) (interface{}, error) {
// 	plugin.Logger(ctx).Trace("getAwsAccountAkas")
// 	nodeInfo := d.HydrateItem.(*nodeData)

// 	if nodeInfo.Aliases != nil && len(nodeInfo.Aliases) > 0 {
// 		return nodeInfo.Aliases[0], nil
// 	}

// 	return nodeInfo.commonColumnData.AccountId, nil
// }

// func nodeDataToAkas(ctx context.Context, d *transform.TransformData) (interface{}, error) {
// 	plugin.Logger(ctx).Trace("nodeDataToTitle")
// 	nodeInfo := d.HydrateItem.(*nodeData)

// 	akas := []string{"arn:" + nodeInfo.commonColumnData.Partition + ":::" + nodeInfo.commonColumnData.AccountId}

// 	return akas, nil
// }
