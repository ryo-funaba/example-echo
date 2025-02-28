package mock

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	dynamoDB "github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/ryo-funaba/example_echo/internal/utils/errorutil"
)

type MockDynamodbRepository struct {
	Table string
	PK    string
	SK    string
}

func (m MockDynamodbRepository) BatchExecuteStatement(*dynamoDB.BatchExecuteStatementInput) (*dynamoDB.BatchExecuteStatementOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) BatchExecuteStatementWithContext(aws.Context, *dynamoDB.BatchExecuteStatementInput, ...request.Option) (*dynamoDB.BatchExecuteStatementOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) BatchExecuteStatementRequest(*dynamoDB.BatchExecuteStatementInput) (*request.Request, *dynamoDB.BatchExecuteStatementOutput) {
	return nil, nil
}

func (m MockDynamodbRepository) BatchGetItem(*dynamoDB.BatchGetItemInput) (*dynamoDB.BatchGetItemOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) BatchGetItemWithContext(aws.Context, *dynamoDB.BatchGetItemInput, ...request.Option) (*dynamoDB.BatchGetItemOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) BatchGetItemRequest(*dynamoDB.BatchGetItemInput) (*request.Request, *dynamoDB.BatchGetItemOutput) {
	return nil, nil
}

func (m MockDynamodbRepository) BatchGetItemPages(*dynamoDB.BatchGetItemInput, func(*dynamoDB.BatchGetItemOutput, bool) bool) error {
	return nil
}

func (m MockDynamodbRepository) BatchGetItemPagesWithContext(aws.Context, *dynamoDB.BatchGetItemInput, func(*dynamoDB.BatchGetItemOutput, bool) bool, ...request.Option) error {
	return nil
}

func (m MockDynamodbRepository) BatchWriteItem(*dynamoDB.BatchWriteItemInput) (*dynamoDB.BatchWriteItemOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) BatchWriteItemWithContext(aws.Context, *dynamoDB.BatchWriteItemInput, ...request.Option) (*dynamoDB.BatchWriteItemOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) BatchWriteItemRequest(*dynamoDB.BatchWriteItemInput) (*request.Request, *dynamoDB.BatchWriteItemOutput) {
	return nil, nil
}

func (m MockDynamodbRepository) CreateBackup(*dynamoDB.CreateBackupInput) (*dynamoDB.CreateBackupOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) CreateBackupWithContext(aws.Context, *dynamoDB.CreateBackupInput, ...request.Option) (*dynamoDB.CreateBackupOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) CreateBackupRequest(*dynamoDB.CreateBackupInput) (*request.Request, *dynamoDB.CreateBackupOutput) {
	return nil, nil
}

func (m MockDynamodbRepository) CreateGlobalTable(*dynamoDB.CreateGlobalTableInput) (*dynamoDB.CreateGlobalTableOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) CreateGlobalTableWithContext(aws.Context, *dynamoDB.CreateGlobalTableInput, ...request.Option) (*dynamoDB.CreateGlobalTableOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) CreateGlobalTableRequest(*dynamoDB.CreateGlobalTableInput) (*request.Request, *dynamoDB.CreateGlobalTableOutput) {
	return nil, nil
}

func (m MockDynamodbRepository) CreateTable(*dynamoDB.CreateTableInput) (*dynamoDB.CreateTableOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) CreateTableWithContext(aws.Context, *dynamoDB.CreateTableInput, ...request.Option) (*dynamoDB.CreateTableOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) CreateTableRequest(*dynamoDB.CreateTableInput) (*request.Request, *dynamoDB.CreateTableOutput) {
	return nil, nil
}

func (m MockDynamodbRepository) DeleteBackup(*dynamoDB.DeleteBackupInput) (*dynamoDB.DeleteBackupOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) DeleteBackupWithContext(aws.Context, *dynamoDB.DeleteBackupInput, ...request.Option) (*dynamoDB.DeleteBackupOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) DeleteBackupRequest(*dynamoDB.DeleteBackupInput) (*request.Request, *dynamoDB.DeleteBackupOutput) {
	return nil, nil
}

func (m MockDynamodbRepository) DeleteItem(*dynamoDB.DeleteItemInput) (*dynamoDB.DeleteItemOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) DeleteItemWithContext(aws.Context, *dynamoDB.DeleteItemInput, ...request.Option) (*dynamoDB.DeleteItemOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) DeleteItemRequest(*dynamoDB.DeleteItemInput) (*request.Request, *dynamoDB.DeleteItemOutput) {
	return nil, nil
}

func (m MockDynamodbRepository) DeleteTable(*dynamoDB.DeleteTableInput) (*dynamoDB.DeleteTableOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) DeleteTableWithContext(aws.Context, *dynamoDB.DeleteTableInput, ...request.Option) (*dynamoDB.DeleteTableOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) DeleteTableRequest(*dynamoDB.DeleteTableInput) (*request.Request, *dynamoDB.DeleteTableOutput) {
	return nil, nil
}

func (m MockDynamodbRepository) DescribeBackup(*dynamoDB.DescribeBackupInput) (*dynamoDB.DescribeBackupOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) DescribeBackupWithContext(aws.Context, *dynamoDB.DescribeBackupInput, ...request.Option) (*dynamoDB.DescribeBackupOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) DescribeBackupRequest(*dynamoDB.DescribeBackupInput) (*request.Request, *dynamoDB.DescribeBackupOutput) {
	return nil, nil
}

func (m MockDynamodbRepository) DescribeContinuousBackups(*dynamoDB.DescribeContinuousBackupsInput) (*dynamoDB.DescribeContinuousBackupsOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) DescribeContinuousBackupsWithContext(aws.Context, *dynamoDB.DescribeContinuousBackupsInput, ...request.Option) (*dynamoDB.DescribeContinuousBackupsOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) DescribeContinuousBackupsRequest(*dynamoDB.DescribeContinuousBackupsInput) (*request.Request, *dynamoDB.DescribeContinuousBackupsOutput) {
	return nil, nil
}

func (m MockDynamodbRepository) DescribeContributorInsights(*dynamoDB.DescribeContributorInsightsInput) (*dynamoDB.DescribeContributorInsightsOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) DescribeContributorInsightsWithContext(aws.Context, *dynamoDB.DescribeContributorInsightsInput, ...request.Option) (*dynamoDB.DescribeContributorInsightsOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) DescribeContributorInsightsRequest(*dynamoDB.DescribeContributorInsightsInput) (*request.Request, *dynamoDB.DescribeContributorInsightsOutput) {
	return nil, nil
}

func (m MockDynamodbRepository) DescribeEndpoints(*dynamoDB.DescribeEndpointsInput) (*dynamoDB.DescribeEndpointsOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) DescribeEndpointsWithContext(aws.Context, *dynamoDB.DescribeEndpointsInput, ...request.Option) (*dynamoDB.DescribeEndpointsOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) DescribeEndpointsRequest(*dynamoDB.DescribeEndpointsInput) (*request.Request, *dynamoDB.DescribeEndpointsOutput) {
	return nil, nil
}

func (m MockDynamodbRepository) DescribeExport(*dynamoDB.DescribeExportInput) (*dynamoDB.DescribeExportOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) DescribeExportWithContext(aws.Context, *dynamoDB.DescribeExportInput, ...request.Option) (*dynamoDB.DescribeExportOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) DescribeExportRequest(*dynamoDB.DescribeExportInput) (*request.Request, *dynamoDB.DescribeExportOutput) {
	return nil, nil
}

func (m MockDynamodbRepository) DescribeGlobalTable(*dynamoDB.DescribeGlobalTableInput) (*dynamoDB.DescribeGlobalTableOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) DescribeGlobalTableWithContext(aws.Context, *dynamoDB.DescribeGlobalTableInput, ...request.Option) (*dynamoDB.DescribeGlobalTableOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) DescribeGlobalTableRequest(*dynamoDB.DescribeGlobalTableInput) (*request.Request, *dynamoDB.DescribeGlobalTableOutput) {
	return nil, nil
}

func (m MockDynamodbRepository) DescribeGlobalTableSettings(*dynamoDB.DescribeGlobalTableSettingsInput) (*dynamoDB.DescribeGlobalTableSettingsOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) DescribeGlobalTableSettingsWithContext(aws.Context, *dynamoDB.DescribeGlobalTableSettingsInput, ...request.Option) (*dynamoDB.DescribeGlobalTableSettingsOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) DescribeGlobalTableSettingsRequest(*dynamoDB.DescribeGlobalTableSettingsInput) (*request.Request, *dynamoDB.DescribeGlobalTableSettingsOutput) {
	return nil, nil
}

func (m MockDynamodbRepository) DescribeImport(*dynamoDB.DescribeImportInput) (*dynamoDB.DescribeImportOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) DescribeImportWithContext(aws.Context, *dynamoDB.DescribeImportInput, ...request.Option) (*dynamoDB.DescribeImportOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) DescribeImportRequest(*dynamoDB.DescribeImportInput) (*request.Request, *dynamoDB.DescribeImportOutput) {
	return nil, nil
}

func (m MockDynamodbRepository) DescribeKinesisStreamingDestination(*dynamoDB.DescribeKinesisStreamingDestinationInput) (*dynamoDB.DescribeKinesisStreamingDestinationOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) DescribeKinesisStreamingDestinationWithContext(aws.Context, *dynamoDB.DescribeKinesisStreamingDestinationInput, ...request.Option) (*dynamoDB.DescribeKinesisStreamingDestinationOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) DescribeKinesisStreamingDestinationRequest(*dynamoDB.DescribeKinesisStreamingDestinationInput) (*request.Request, *dynamoDB.DescribeKinesisStreamingDestinationOutput) {
	return nil, nil
}

func (m MockDynamodbRepository) DescribeLimits(*dynamoDB.DescribeLimitsInput) (*dynamoDB.DescribeLimitsOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) DescribeLimitsWithContext(aws.Context, *dynamoDB.DescribeLimitsInput, ...request.Option) (*dynamoDB.DescribeLimitsOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) DescribeLimitsRequest(*dynamoDB.DescribeLimitsInput) (*request.Request, *dynamoDB.DescribeLimitsOutput) {
	return nil, nil
}

func (m MockDynamodbRepository) DescribeTable(*dynamoDB.DescribeTableInput) (*dynamoDB.DescribeTableOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) DescribeTableWithContext(aws.Context, *dynamoDB.DescribeTableInput, ...request.Option) (*dynamoDB.DescribeTableOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) DescribeTableRequest(*dynamoDB.DescribeTableInput) (*request.Request, *dynamoDB.DescribeTableOutput) {
	return nil, nil
}

func (m MockDynamodbRepository) DescribeTableReplicaAutoScaling(*dynamoDB.DescribeTableReplicaAutoScalingInput) (*dynamoDB.DescribeTableReplicaAutoScalingOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) DescribeTableReplicaAutoScalingWithContext(aws.Context, *dynamoDB.DescribeTableReplicaAutoScalingInput, ...request.Option) (*dynamoDB.DescribeTableReplicaAutoScalingOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) DescribeTableReplicaAutoScalingRequest(*dynamoDB.DescribeTableReplicaAutoScalingInput) (*request.Request, *dynamoDB.DescribeTableReplicaAutoScalingOutput) {
	return nil, nil
}

func (m MockDynamodbRepository) DescribeTimeToLive(*dynamoDB.DescribeTimeToLiveInput) (*dynamoDB.DescribeTimeToLiveOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) DescribeTimeToLiveWithContext(aws.Context, *dynamoDB.DescribeTimeToLiveInput, ...request.Option) (*dynamoDB.DescribeTimeToLiveOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) DescribeTimeToLiveRequest(*dynamoDB.DescribeTimeToLiveInput) (*request.Request, *dynamoDB.DescribeTimeToLiveOutput) {
	return nil, nil
}

func (m MockDynamodbRepository) DisableKinesisStreamingDestination(*dynamoDB.DisableKinesisStreamingDestinationInput) (*dynamoDB.DisableKinesisStreamingDestinationOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) DisableKinesisStreamingDestinationWithContext(aws.Context, *dynamoDB.DisableKinesisStreamingDestinationInput, ...request.Option) (*dynamoDB.DisableKinesisStreamingDestinationOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) DisableKinesisStreamingDestinationRequest(*dynamoDB.DisableKinesisStreamingDestinationInput) (*request.Request, *dynamoDB.DisableKinesisStreamingDestinationOutput) {
	return nil, nil
}

func (m MockDynamodbRepository) EnableKinesisStreamingDestination(*dynamoDB.EnableKinesisStreamingDestinationInput) (*dynamoDB.EnableKinesisStreamingDestinationOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) EnableKinesisStreamingDestinationWithContext(aws.Context, *dynamoDB.EnableKinesisStreamingDestinationInput, ...request.Option) (*dynamoDB.EnableKinesisStreamingDestinationOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) EnableKinesisStreamingDestinationRequest(*dynamoDB.EnableKinesisStreamingDestinationInput) (*request.Request, *dynamoDB.EnableKinesisStreamingDestinationOutput) {
	return nil, nil
}

func (m MockDynamodbRepository) ExecuteStatement(*dynamoDB.ExecuteStatementInput) (*dynamoDB.ExecuteStatementOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) ExecuteStatementWithContext(aws.Context, *dynamoDB.ExecuteStatementInput, ...request.Option) (*dynamoDB.ExecuteStatementOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) ExecuteStatementRequest(*dynamoDB.ExecuteStatementInput) (*request.Request, *dynamoDB.ExecuteStatementOutput) {
	return nil, nil
}

func (m MockDynamodbRepository) ExecuteTransaction(*dynamoDB.ExecuteTransactionInput) (*dynamoDB.ExecuteTransactionOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) ExecuteTransactionWithContext(aws.Context, *dynamoDB.ExecuteTransactionInput, ...request.Option) (*dynamoDB.ExecuteTransactionOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) ExecuteTransactionRequest(*dynamoDB.ExecuteTransactionInput) (*request.Request, *dynamoDB.ExecuteTransactionOutput) {
	return nil, nil
}

func (m MockDynamodbRepository) ExportTableToPointInTime(*dynamoDB.ExportTableToPointInTimeInput) (*dynamoDB.ExportTableToPointInTimeOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) ExportTableToPointInTimeWithContext(aws.Context, *dynamoDB.ExportTableToPointInTimeInput, ...request.Option) (*dynamoDB.ExportTableToPointInTimeOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) ExportTableToPointInTimeRequest(*dynamoDB.ExportTableToPointInTimeInput) (*request.Request, *dynamoDB.ExportTableToPointInTimeOutput) {
	return nil, nil
}

func (m MockDynamodbRepository) GetItem(input *dynamoDB.GetItemInput) (*dynamoDB.GetItemOutput, error) {
	if m.Table != *input.TableName {
		return nil, errorutil.Errorf(errorutil.Unknown, "TableName is invalid")
	}

	if m.PK != *input.Key["PK"].S {
		return nil, errorutil.Errorf(errorutil.Unknown, "PK is invalid")
	}

	if m.SK != *input.Key["SK"].S {
		return nil, errorutil.Errorf(errorutil.Unknown, "SK is invalid")
	}

	return &dynamoDB.GetItemOutput{}, nil
}

func (m MockDynamodbRepository) GetItemWithContext(aws.Context, *dynamoDB.GetItemInput, ...request.Option) (*dynamoDB.GetItemOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) GetItemRequest(*dynamoDB.GetItemInput) (*request.Request, *dynamoDB.GetItemOutput) {
	return nil, nil
}

func (m MockDynamodbRepository) ImportTable(*dynamoDB.ImportTableInput) (*dynamoDB.ImportTableOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) ImportTableWithContext(aws.Context, *dynamoDB.ImportTableInput, ...request.Option) (*dynamoDB.ImportTableOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) ImportTableRequest(*dynamoDB.ImportTableInput) (*request.Request, *dynamoDB.ImportTableOutput) {
	return nil, nil
}

func (m MockDynamodbRepository) ListBackups(*dynamoDB.ListBackupsInput) (*dynamoDB.ListBackupsOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) ListBackupsWithContext(aws.Context, *dynamoDB.ListBackupsInput, ...request.Option) (*dynamoDB.ListBackupsOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) ListBackupsRequest(*dynamoDB.ListBackupsInput) (*request.Request, *dynamoDB.ListBackupsOutput) {
	return nil, nil
}

func (m MockDynamodbRepository) ListContributorInsights(*dynamoDB.ListContributorInsightsInput) (*dynamoDB.ListContributorInsightsOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) ListContributorInsightsWithContext(aws.Context, *dynamoDB.ListContributorInsightsInput, ...request.Option) (*dynamoDB.ListContributorInsightsOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) ListContributorInsightsRequest(*dynamoDB.ListContributorInsightsInput) (*request.Request, *dynamoDB.ListContributorInsightsOutput) {
	return nil, nil
}

func (m MockDynamodbRepository) ListContributorInsightsPages(*dynamoDB.ListContributorInsightsInput, func(*dynamoDB.ListContributorInsightsOutput, bool) bool) error {
	return nil
}

func (m MockDynamodbRepository) ListContributorInsightsPagesWithContext(aws.Context, *dynamoDB.ListContributorInsightsInput, func(*dynamoDB.ListContributorInsightsOutput, bool) bool, ...request.Option) error {
	return nil
}

func (m MockDynamodbRepository) ListExports(*dynamoDB.ListExportsInput) (*dynamoDB.ListExportsOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) ListExportsWithContext(aws.Context, *dynamoDB.ListExportsInput, ...request.Option) (*dynamoDB.ListExportsOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) ListExportsRequest(*dynamoDB.ListExportsInput) (*request.Request, *dynamoDB.ListExportsOutput) {
	return nil, nil
}

func (m MockDynamodbRepository) ListExportsPages(*dynamoDB.ListExportsInput, func(*dynamoDB.ListExportsOutput, bool) bool) error {
	return nil
}

func (m MockDynamodbRepository) ListExportsPagesWithContext(aws.Context, *dynamoDB.ListExportsInput, func(*dynamoDB.ListExportsOutput, bool) bool, ...request.Option) error {
	return nil
}

func (m MockDynamodbRepository) ListGlobalTables(*dynamoDB.ListGlobalTablesInput) (*dynamoDB.ListGlobalTablesOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) ListGlobalTablesWithContext(aws.Context, *dynamoDB.ListGlobalTablesInput, ...request.Option) (*dynamoDB.ListGlobalTablesOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) ListGlobalTablesRequest(*dynamoDB.ListGlobalTablesInput) (*request.Request, *dynamoDB.ListGlobalTablesOutput) {
	return nil, nil
}

func (m MockDynamodbRepository) ListImports(*dynamoDB.ListImportsInput) (*dynamoDB.ListImportsOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) ListImportsWithContext(aws.Context, *dynamoDB.ListImportsInput, ...request.Option) (*dynamoDB.ListImportsOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) ListImportsRequest(*dynamoDB.ListImportsInput) (*request.Request, *dynamoDB.ListImportsOutput) {
	return nil, nil
}

func (m MockDynamodbRepository) ListImportsPages(*dynamoDB.ListImportsInput, func(*dynamoDB.ListImportsOutput, bool) bool) error {
	return nil
}

func (m MockDynamodbRepository) ListImportsPagesWithContext(aws.Context, *dynamoDB.ListImportsInput, func(*dynamoDB.ListImportsOutput, bool) bool, ...request.Option) error {
	return nil
}

func (m MockDynamodbRepository) ListTables(*dynamoDB.ListTablesInput) (*dynamoDB.ListTablesOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) ListTablesWithContext(aws.Context, *dynamoDB.ListTablesInput, ...request.Option) (*dynamoDB.ListTablesOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) ListTablesRequest(*dynamoDB.ListTablesInput) (*request.Request, *dynamoDB.ListTablesOutput) {
	return nil, nil
}

func (m MockDynamodbRepository) ListTablesPages(*dynamoDB.ListTablesInput, func(*dynamoDB.ListTablesOutput, bool) bool) error {
	return nil
}

func (m MockDynamodbRepository) ListTablesPagesWithContext(aws.Context, *dynamoDB.ListTablesInput, func(*dynamoDB.ListTablesOutput, bool) bool, ...request.Option) error {
	return nil
}

func (m MockDynamodbRepository) ListTagsOfResource(*dynamoDB.ListTagsOfResourceInput) (*dynamoDB.ListTagsOfResourceOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) ListTagsOfResourceWithContext(aws.Context, *dynamoDB.ListTagsOfResourceInput, ...request.Option) (*dynamoDB.ListTagsOfResourceOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) ListTagsOfResourceRequest(*dynamoDB.ListTagsOfResourceInput) (*request.Request, *dynamoDB.ListTagsOfResourceOutput) {
	return nil, nil
}

func (m MockDynamodbRepository) PutItem(input *dynamoDB.PutItemInput) (*dynamoDB.PutItemOutput, error) {
	if m.Table != *input.TableName {
		return nil, errorutil.Errorf(errorutil.Unknown, "TableName is invalid")
	}

	if m.PK != *input.Item["PK"].S {
		return nil, errorutil.Errorf(errorutil.Unknown, "PK is invalid")
	}

	if m.SK != *input.Item["SK"].S {
		return nil, errorutil.Errorf(errorutil.Unknown, "SK is invalid")
	}

	return &dynamoDB.PutItemOutput{}, nil
}

func (m MockDynamodbRepository) PutItemWithContext(aws.Context, *dynamoDB.PutItemInput, ...request.Option) (*dynamoDB.PutItemOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) PutItemRequest(*dynamoDB.PutItemInput) (*request.Request, *dynamoDB.PutItemOutput) {
	return nil, nil
}

func (m MockDynamodbRepository) Query(*dynamoDB.QueryInput) (*dynamoDB.QueryOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) QueryWithContext(aws.Context, *dynamoDB.QueryInput, ...request.Option) (*dynamoDB.QueryOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) QueryRequest(*dynamoDB.QueryInput) (*request.Request, *dynamoDB.QueryOutput) {
	return nil, nil
}

func (m MockDynamodbRepository) QueryPages(*dynamoDB.QueryInput, func(*dynamoDB.QueryOutput, bool) bool) error {
	return nil
}

func (m MockDynamodbRepository) QueryPagesWithContext(aws.Context, *dynamoDB.QueryInput, func(*dynamoDB.QueryOutput, bool) bool, ...request.Option) error {
	return nil
}

func (m MockDynamodbRepository) RestoreTableFromBackup(*dynamoDB.RestoreTableFromBackupInput) (*dynamoDB.RestoreTableFromBackupOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) RestoreTableFromBackupWithContext(aws.Context, *dynamoDB.RestoreTableFromBackupInput, ...request.Option) (*dynamoDB.RestoreTableFromBackupOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) RestoreTableFromBackupRequest(*dynamoDB.RestoreTableFromBackupInput) (*request.Request, *dynamoDB.RestoreTableFromBackupOutput) {
	return nil, nil
}

func (m MockDynamodbRepository) RestoreTableToPointInTime(*dynamoDB.RestoreTableToPointInTimeInput) (*dynamoDB.RestoreTableToPointInTimeOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) RestoreTableToPointInTimeWithContext(aws.Context, *dynamoDB.RestoreTableToPointInTimeInput, ...request.Option) (*dynamoDB.RestoreTableToPointInTimeOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) RestoreTableToPointInTimeRequest(*dynamoDB.RestoreTableToPointInTimeInput) (*request.Request, *dynamoDB.RestoreTableToPointInTimeOutput) {
	return nil, nil
}

func (m MockDynamodbRepository) Scan(*dynamoDB.ScanInput) (*dynamoDB.ScanOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) ScanWithContext(aws.Context, *dynamoDB.ScanInput, ...request.Option) (*dynamoDB.ScanOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) ScanRequest(*dynamoDB.ScanInput) (*request.Request, *dynamoDB.ScanOutput) {
	return nil, nil
}

func (m MockDynamodbRepository) ScanPages(*dynamoDB.ScanInput, func(*dynamoDB.ScanOutput, bool) bool) error {
	return nil
}

func (m MockDynamodbRepository) ScanPagesWithContext(aws.Context, *dynamoDB.ScanInput, func(*dynamoDB.ScanOutput, bool) bool, ...request.Option) error {
	return nil
}

func (m MockDynamodbRepository) TagResource(*dynamoDB.TagResourceInput) (*dynamoDB.TagResourceOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) TagResourceWithContext(aws.Context, *dynamoDB.TagResourceInput, ...request.Option) (*dynamoDB.TagResourceOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) TagResourceRequest(*dynamoDB.TagResourceInput) (*request.Request, *dynamoDB.TagResourceOutput) {
	return nil, nil
}

func (m MockDynamodbRepository) TransactGetItems(*dynamoDB.TransactGetItemsInput) (*dynamoDB.TransactGetItemsOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) TransactGetItemsWithContext(aws.Context, *dynamoDB.TransactGetItemsInput, ...request.Option) (*dynamoDB.TransactGetItemsOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) TransactGetItemsRequest(*dynamoDB.TransactGetItemsInput) (*request.Request, *dynamoDB.TransactGetItemsOutput) {
	return nil, nil
}

func (m MockDynamodbRepository) TransactWriteItems(*dynamoDB.TransactWriteItemsInput) (*dynamoDB.TransactWriteItemsOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) TransactWriteItemsWithContext(aws.Context, *dynamoDB.TransactWriteItemsInput, ...request.Option) (*dynamoDB.TransactWriteItemsOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) TransactWriteItemsRequest(*dynamoDB.TransactWriteItemsInput) (*request.Request, *dynamoDB.TransactWriteItemsOutput) {
	return nil, nil
}

func (m MockDynamodbRepository) UntagResource(*dynamoDB.UntagResourceInput) (*dynamoDB.UntagResourceOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) UntagResourceWithContext(aws.Context, *dynamoDB.UntagResourceInput, ...request.Option) (*dynamoDB.UntagResourceOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) UntagResourceRequest(*dynamoDB.UntagResourceInput) (*request.Request, *dynamoDB.UntagResourceOutput) {
	return nil, nil
}

func (m MockDynamodbRepository) UpdateContinuousBackups(*dynamoDB.UpdateContinuousBackupsInput) (*dynamoDB.UpdateContinuousBackupsOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) UpdateContinuousBackupsWithContext(aws.Context, *dynamoDB.UpdateContinuousBackupsInput, ...request.Option) (*dynamoDB.UpdateContinuousBackupsOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) UpdateContinuousBackupsRequest(*dynamoDB.UpdateContinuousBackupsInput) (*request.Request, *dynamoDB.UpdateContinuousBackupsOutput) {
	return nil, nil
}

func (m MockDynamodbRepository) UpdateContributorInsights(*dynamoDB.UpdateContributorInsightsInput) (*dynamoDB.UpdateContributorInsightsOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) UpdateContributorInsightsWithContext(aws.Context, *dynamoDB.UpdateContributorInsightsInput, ...request.Option) (*dynamoDB.UpdateContributorInsightsOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) UpdateContributorInsightsRequest(*dynamoDB.UpdateContributorInsightsInput) (*request.Request, *dynamoDB.UpdateContributorInsightsOutput) {
	return nil, nil
}

func (m MockDynamodbRepository) UpdateGlobalTable(*dynamoDB.UpdateGlobalTableInput) (*dynamoDB.UpdateGlobalTableOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) UpdateGlobalTableWithContext(aws.Context, *dynamoDB.UpdateGlobalTableInput, ...request.Option) (*dynamoDB.UpdateGlobalTableOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) UpdateGlobalTableRequest(*dynamoDB.UpdateGlobalTableInput) (*request.Request, *dynamoDB.UpdateGlobalTableOutput) {
	return nil, nil
}

func (m MockDynamodbRepository) UpdateGlobalTableSettings(*dynamoDB.UpdateGlobalTableSettingsInput) (*dynamoDB.UpdateGlobalTableSettingsOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) UpdateGlobalTableSettingsWithContext(aws.Context, *dynamoDB.UpdateGlobalTableSettingsInput, ...request.Option) (*dynamoDB.UpdateGlobalTableSettingsOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) UpdateGlobalTableSettingsRequest(*dynamoDB.UpdateGlobalTableSettingsInput) (*request.Request, *dynamoDB.UpdateGlobalTableSettingsOutput) {
	return nil, nil
}

func (m MockDynamodbRepository) UpdateItem(*dynamoDB.UpdateItemInput) (*dynamoDB.UpdateItemOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) UpdateItemWithContext(aws.Context, *dynamoDB.UpdateItemInput, ...request.Option) (*dynamoDB.UpdateItemOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) UpdateItemRequest(*dynamoDB.UpdateItemInput) (*request.Request, *dynamoDB.UpdateItemOutput) {
	return nil, nil
}

func (m MockDynamodbRepository) UpdateTable(*dynamoDB.UpdateTableInput) (*dynamoDB.UpdateTableOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) UpdateTableWithContext(aws.Context, *dynamoDB.UpdateTableInput, ...request.Option) (*dynamoDB.UpdateTableOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) UpdateTableRequest(*dynamoDB.UpdateTableInput) (*request.Request, *dynamoDB.UpdateTableOutput) {
	return nil, nil
}

func (m MockDynamodbRepository) UpdateTableReplicaAutoScaling(*dynamoDB.UpdateTableReplicaAutoScalingInput) (*dynamoDB.UpdateTableReplicaAutoScalingOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) UpdateTableReplicaAutoScalingWithContext(aws.Context, *dynamoDB.UpdateTableReplicaAutoScalingInput, ...request.Option) (*dynamoDB.UpdateTableReplicaAutoScalingOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) UpdateTableReplicaAutoScalingRequest(*dynamoDB.UpdateTableReplicaAutoScalingInput) (*request.Request, *dynamoDB.UpdateTableReplicaAutoScalingOutput) {
	return nil, nil
}

func (m MockDynamodbRepository) UpdateTimeToLive(*dynamoDB.UpdateTimeToLiveInput) (*dynamoDB.UpdateTimeToLiveOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) UpdateTimeToLiveWithContext(aws.Context, *dynamoDB.UpdateTimeToLiveInput, ...request.Option) (*dynamoDB.UpdateTimeToLiveOutput, error) {
	return nil, nil
}

func (m MockDynamodbRepository) UpdateTimeToLiveRequest(*dynamoDB.UpdateTimeToLiveInput) (*request.Request, *dynamoDB.UpdateTimeToLiveOutput) {
	return nil, nil
}

func (m MockDynamodbRepository) WaitUntilTableExists(*dynamoDB.DescribeTableInput) error {
	return nil
}

func (m MockDynamodbRepository) WaitUntilTableExistsWithContext(aws.Context, *dynamoDB.DescribeTableInput, ...request.WaiterOption) error {
	return nil
}

func (m MockDynamodbRepository) WaitUntilTableNotExists(*dynamoDB.DescribeTableInput) error {
	return nil
}

func (m MockDynamodbRepository) WaitUntilTableNotExistsWithContext(aws.Context, *dynamoDB.DescribeTableInput, ...request.WaiterOption) error {
	return nil
}
