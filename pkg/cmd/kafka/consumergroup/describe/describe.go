package describe

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"sort"

	"github.com/redhat-developer/app-services-cli/pkg/cmdutil"
	cgutil "github.com/redhat-developer/app-services-cli/pkg/kafka/consumergroup"
	"github.com/redhat-developer/app-services-cli/pkg/localize"

	"github.com/redhat-developer/app-services-cli/internal/config"
	strimziadminclient "github.com/redhat-developer/app-services-cli/pkg/api/strimzi-admin/client"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/flag"
	"github.com/redhat-developer/app-services-cli/pkg/color"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/dump"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

type Options struct {
	kafkaID      string
	outputFormat string
	id           string

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
	localizer  localize.Localizer
}

type consumerRow struct {
	MemberID      string `json:"memberId,omitempty" header:"Member ID"`
	Partition     int    `json:"partition,omitempty" header:"Partition"`
	Topic         string `json:"topic,omitempty" header:"Topic"`
	LogEndOffset  int    `json:"logEndOffset,omitempty" header:"Log end offset"`
	CurrentOffset int    `json:"offset,omitempty" header:"Current offset"`
	OffsetLag     int    `json:"lag,omitempty" header:"Offset lag"`
}

// NewDescribeConsumerGroupCommand gets a new command for describing a consumer group.
func NewDescribeConsumerGroupCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Connection: f.Connection,
		Config:     f.Config,
		IO:         f.IOStreams,
		localizer:  f.Localizer,
	}
	cmd := &cobra.Command{
		Use:     opts.localizer.MustLocalize("kafka.consumerGroup.describe.cmd.use"),
		Short:   opts.localizer.MustLocalize("kafka.consumerGroup.describe.cmd.shortDescription"),
		Long:    opts.localizer.MustLocalize("kafka.consumerGroup.describe.cmd.longDescription"),
		Example: opts.localizer.MustLocalize("kafka.consumerGroup.describe.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			if opts.outputFormat != "" {
				if err = flag.ValidateOutput(opts.outputFormat); err != nil {
					return err
				}
			}

			if opts.kafkaID != "" {
				return runCmd(opts)
			}

			cfg, err := opts.Config.Load()
			if err != nil {
				return err
			}

			if !cfg.HasKafka() {
				return errors.New(opts.localizer.MustLocalize("kafka.consumerGroup.common.error.noKafkaSelected"))
			}

			opts.kafkaID = cfg.Services.Kafka.ClusterID

			return runCmd(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "", opts.localizer.MustLocalize("kafka.consumerGroup.common.flag.output.description"))
	cmd.Flags().StringVar(&opts.id, "id", "", opts.localizer.MustLocalize("kafka.consumerGroup.common.flag.id.description", localize.NewEntry("Action", "view")))
	_ = cmd.MarkFlagRequired("id")

	// flag based completions for ID
	_ = cmd.RegisterFlagCompletionFunc("id", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return cmdutil.FilterValidConsumerGroupIDs(f, toComplete)
	})

	return cmd
}

func runCmd(opts *Options) error {
	conn, err := opts.Connection(connection.DefaultConfigRequireMasAuth)
	if err != nil {
		return err
	}

	api, kafkaInstance, err := conn.API().TopicAdmin(opts.kafkaID)
	if err != nil {
		return err
	}

	ctx := context.Background()

	consumerGroupData, httpRes, err := api.GetConsumerGroupById(ctx, opts.id).Execute()

	if err != nil {
		if httpRes == nil {
			return err
		}

		cgIDPair := localize.NewEntry("ID", opts.id)
		kafkaNameTmplPair := localize.NewEntry("InstanceName", kafkaInstance.GetName())
		operationTmplPair := localize.NewEntry("Operation", "view")

		switch httpRes.StatusCode {
		case 404:
			return errors.New(opts.localizer.MustLocalize("kafka.consumerGroup.common.error.notFoundError", cgIDPair, kafkaNameTmplPair))
		case 401:
			return errors.New(opts.localizer.MustLocalize("kafka.consumerGroup.common.error.unauthorized", operationTmplPair))
		case 403:
			return errors.New(opts.localizer.MustLocalize("kafka.consumerGroup.common.error.forbidden", operationTmplPair))
		case 500:
			return errors.New(opts.localizer.MustLocalize("kafka.consumerGroup.common.error.internalServerError"))
		case 503:
			return errors.New(opts.localizer.MustLocalize("kafka.consumerGroup.common.error.unableToConnectToKafka", localize.NewEntry("Name", kafkaInstance.GetName())))
		default:
			return err
		}
	}

	stdout := opts.IO.Out
	switch opts.outputFormat {
	case "json":
		data, _ := json.Marshal(consumerGroupData)
		_ = dump.JSON(stdout, data)
	case "yaml", "yml":
		data, _ := yaml.Marshal(consumerGroupData)
		_ = dump.YAML(stdout, data)
	default:
		printConsumerGroupDetails(stdout, consumerGroupData, opts.localizer)
	}

	return nil
}

func mapConsumerGroupDescribeToTableFormat(consumers []strimziadminclient.Consumer) []consumerRow {
	var rows []consumerRow = []consumerRow{}

	for _, consumer := range consumers {

		row := consumerRow{
			Partition:     int(consumer.GetPartition()),
			Topic:         consumer.GetTopic(),
			MemberID:      consumer.GetMemberId(),
			LogEndOffset:  int(consumer.GetLogEndOffset()),
			CurrentOffset: int(consumer.GetOffset()),
			OffsetLag:     int(consumer.GetLag()),
		}
		rows = append(rows, row)
	}

	// sort members by partition number
	sort.Slice(rows, func(i, j int) bool {
		return rows[i].Partition < rows[j].Partition
	})

	return rows
}

// print the consumer grooup details
func printConsumerGroupDetails(w io.Writer, consumerGroupData strimziadminclient.ConsumerGroup, localizer localize.Localizer) {
	fmt.Fprintln(w, "")
	consumers := consumerGroupData.GetConsumers()

	activeMembersCount := cgutil.GetActiveConsumersCount(consumers)
	partitionsWithLagCount := cgutil.GetPartitionsWithLag(consumers)

	fmt.Fprintln(w, color.Bold(localizer.MustLocalize("kafka.consumerGroup.describe.output.activeMembers")), activeMembersCount, "\t", color.Bold(localizer.MustLocalize("kafka.consumerGroup.describe.output.partitionsWithLag")), partitionsWithLagCount)
	fmt.Fprintln(w, "")

	rows := mapConsumerGroupDescribeToTableFormat(consumers)
	dump.Table(w, rows)
}
