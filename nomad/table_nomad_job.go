package nomad

import (
	"context"

	"github.com/hashicorp/nomad/api"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

//// TABLE DEFINITION

// Region           *string                 `hcl:"region,optional"`
// Namespace        *string                 `hcl:"namespace,optional"`
// ID               *string                 `hcl:"id,optional"`
// Name             *string                 `hcl:"name,optional"`
// Type             *string                 `hcl:"type,optional"`
// Priority         *int                    `hcl:"priority,optional"`
// AllAtOnce        *bool                   `mapstructure:"all_at_once" hcl:"all_at_once,optional"`
// Datacenters      []string                `hcl:"datacenters,optional"`
// Constraints      []*Constraint           `hcl:"constraint,block"`
// Affinities       []*Affinity             `hcl:"affinity,block"`
// TaskGroups       []*TaskGroup            `hcl:"group,block"`
// Update           *UpdateStrategy         `hcl:"update,block"`
// Multiregion      *Multiregion            `hcl:"multiregion,block"`
// Spreads          []*Spread               `hcl:"spread,block"`
// Periodic         *PeriodicConfig         `hcl:"periodic,block"`
// ParameterizedJob *ParameterizedJobConfig `hcl:"parameterized,block"`
// Reschedule       *ReschedulePolicy       `hcl:"reschedule,block"`
// Migrate          *MigrateStrategy        `hcl:"migrate,block"`
// Meta             map[string]string       `hcl:"meta,block"`
// ConsulToken      *string                 `mapstructure:"consul_token" hcl:"consul_token,optional"`
// VaultToken       *string                 `mapstructure:"vault_token" hcl:"vault_token,optional"`

// /* Fields set by server, not sourced from job config file */

// Stop              *bool
// ParentID          *string
// Dispatched        bool
// Payload           []byte
// ConsulNamespace   *string `mapstructure:"consul_namespace"`
// VaultNamespace    *string `mapstructure:"vault_namespace"`
// NomadTokenID      *string `mapstructure:"nomad_token_id"`
// Status            *string
// StatusDescription *string
// Stable            *bool
// Version           *uint64
// SubmitTime        *int64
// CreateIndex       *uint64
// ModifyIndex       *uint64
// JobModifyIndex    *uint64


func tableNomadJob(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "nomad_Job",
		Description: "Nomad jobs",
		List: &plugin.ListConfig{
			Hydrate: listJobs,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getJob,
		},
		Columns: []*plugin.Column{
			{
				Name:        "region",
				Description: "Job region",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Region"),
			},
			{
				Name:        "namespace",
				Description: "Job namespace",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Namespace"),
			},
			{
				Name:        "id",
				Description: "Job ID",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "name",
				Description: "Job name",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
			{
				Name:        "type",
				Description: "Job type",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Type"),
			},
			{
				Name:        "priority",
				Description: "Job priority",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Priority"),
			},
			{
				Name:        "all_at_once",
				Description: "Job allatonce",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("AllAtOnce"),
			},
			{
				Name:        "datacenter",
				Description: "Job datacenters",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Datacenters"),
			},
			{
				Name:        "status",
				Description: "Job status",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Status"),
			},
			{
				Name:        "status_description",
				Description: "Job status description",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("StatusDescription"),
			},
			{
				Name:        "modify_index",
				Description: "Job modify index",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ModifyIndex"),
			},
		},
	}
}

// type JobData struct {
// 	commonColumnData nomadCommonColumnData
// 	Aliases          []*string
// }

//// LIST FUNCTION

func listJobs(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)
	// Create Session
	client, err := NomadClient(ctx, d)
	if err != nil {
		return nil, err
	}

	queryOpts := &api.QueryOptions{}
	jobsResp, _, err := client.Jobs().List(queryOpts)
	if err != nil {
		return nil, err
	}

	for _, job := range jobsResp {
		logger.Trace("listJobs:: Job:", job)
		// d.StreamListItem(ctx, Job)
		jobReadResp, _, err := client.Jobs().Info(job.ID, queryOpts)
		logger.Trace("listJobs:: jobReadResp:", jobReadResp)
		if err != nil {
			return nil, err
		}
		d.StreamListItem(ctx, jobReadResp)
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getJob(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)
	// Create Session
	client, err := NomadClient(ctx, d)
	if err != nil {
		return nil, err
	}

	logger.Trace("getJob:: d.KeyColumnQuals:", d.KeyColumnQuals)
	id := d.KeyColumnQuals["id"].GetStringValue()
	queryOpts := &api.QueryOptions{}
	// JobsResp, _, err := client.GetJobClient(id, queryOpts)
	jobResp, _, err := client.Jobs().Info(id, queryOpts)
	// logger.Trace("getJob:: JobResp:", fmt.Sprintf("%+v", JobResp))
	if err != nil {
		return nil, err
	}
	return jobResp, nil
}

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

// func JobDataToTitle(ctx context.Context, d *transform.TransformData) (interface{}, error) {
// 	plugin.Logger(ctx).Trace("getAwsAccountAkas")
// 	JobInfo := d.HydrateItem.(*JobData)

// 	if JobInfo.Aliases != nil && len(JobInfo.Aliases) > 0 {
// 		return JobInfo.Aliases[0], nil
// 	}

// 	return JobInfo.commonColumnData.AccountId, nil
// }

// func JobDataToAkas(ctx context.Context, d *transform.TransformData) (interface{}, error) {
// 	plugin.Logger(ctx).Trace("JobDataToTitle")
// 	JobInfo := d.HydrateItem.(*JobData)

// 	akas := []string{"arn:" + JobInfo.commonColumnData.Partition + ":::" + JobInfo.commonColumnData.AccountId}

// 	return akas, nil
// }
