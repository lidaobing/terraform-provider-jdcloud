// Copyright 2018 JDCLOUD.COM
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// NOTE: This class is auto generated by the jdcloud code generator program.

package client

import (
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
    rds "github.com/jdcloud-api/jdcloud-sdk-go/services/rds/apis"
    "encoding/json"
    "errors"
)

type RdsClient struct {
    core.JDCloudClient
}

func NewRdsClient(credential *core.Credential) *RdsClient {
    if credential == nil {
        return nil
    }

    config := core.NewConfig()
    config.SetEndpoint("rds.jdcloud-api.com")

    return &RdsClient{
        core.JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "rds",
            Revision:    "0.3.0",
            Logger:      core.NewDefaultLogger(core.LogInfo),
        }}
}

func (c *RdsClient) SetConfig(config *core.Config) {
    c.Config = *config
}

func (c *RdsClient) SetLogger(logger core.Logger) {
    c.Logger = logger
}

/* 获取某个审计文件的下载链接，同时支持内链和外链，链接的有效时间为24小时<br>- 仅支持SQL Server */
func (c *RdsClient) GetAuditDownloadURL(request *rds.GetAuditDownloadURLRequest) (*rds.GetAuditDownloadURLResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &rds.GetAuditDownloadURLResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 对RDS实例进行主备切换。<br>注意：如果实例正在进行备份，那么主备切换将会终止备份操作。可以查看备份策略中的备份开始时间确认是否有备份正在运行。如果确实需要在实例备份时进行主备切换，建议切换完成 后，手工进行一次实例的全备<br>对于SQL Server，主备切换后30分钟内，不支持按时间点恢复/创建，例如在10:05分用户进行了主备切换，那么10:05 ~ 10:35这个时间段不能进行按时间点恢复/创建。<br>- 仅支持SQL Server */
func (c *RdsClient) FailoverInstance(request *rds.FailoverInstanceRequest) (*rds.FailoverInstanceResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &rds.FailoverInstanceResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 删除一个RDS实例或者MySQL的只读实例。删除MySQL主实例时，会同时将对应的MySQL只读实例也删除 [MFA enabled] */
func (c *RdsClient) DeleteInstance(request *rds.DeleteInstanceRequest) (*rds.DeleteInstanceResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &rds.DeleteInstanceResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 获取单库上云工具上传文件的需要的Key。单库上云工具需要正确的key值方能连接到京东云<br>- 仅支持SQL Server */
func (c *RdsClient) GetUploadKey(request *rds.GetUploadKeyRequest) (*rds.GetUploadKeyResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &rds.GetUploadKeyResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 关闭数据库审计。关闭数据库审计后，以前生成的审计结果文件并不会被立即删除。审计结果文件会过期后由系统自动删除，过期时间缺省为6个月<br>- 仅支持SQL Server */
func (c *RdsClient) DeleteAudit(request *rds.DeleteAuditRequest) (*rds.DeleteAuditResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &rds.DeleteAuditResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 根据源实例备份创建一个新实例，并通过追加日志的方式，将新实例的数据恢复到跟源实例指定时间点的数据状态一样。<br>例如根据实例A在“2018-06-18 23:00:00”时间点创建一个实例B，就是新建一个实例B，该实例B的数据跟实例A在“2018-06-18 23:00:00”这个时间点的数据完全一致。<br>对于SQL Server，主备切换后30分钟内，不支持按时间点恢复/创建，例如在10:05分用户进行了主备切换，那么10:05 ~ 10:35这个时间段不能进行按时间点恢复/创建。<br>- 仅支持MySQL */
func (c *RdsClient) CreateInstanceByTime(request *rds.CreateInstanceByTimeRequest) (*rds.CreateInstanceByTimeResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &rds.CreateInstanceByTimeResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 使用实例的全量备份覆盖恢复当前实例<br>- 仅支持MySQL */
func (c *RdsClient) RestoreInstance(request *rds.RestoreInstanceRequest) (*rds.RestoreInstanceResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &rds.RestoreInstanceResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 开启SQL Server的数据库审计功能，目前支持实例级的数据库审计。用户可以根据需要开启、关闭审计、自定义审计策略，并下载审计文件。审计文件为原生的SQL Server审计文件，缺省保存6个月。<br>- 仅支持SQL Server */
func (c *RdsClient) CreateAudit(request *rds.CreateAuditRequest) (*rds.CreateAuditResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &rds.CreateAuditResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 删除RDS实例备份，仅允许删除用户生成的备份，系统自动备份不允许删除。 */
func (c *RdsClient) DeleteBackup(request *rds.DeleteBackupRequest) (*rds.DeleteBackupResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &rds.DeleteBackupResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 修改实例名称，可支持中文，实例名的具体规则可参见帮助中心文档:[名称及密码限制](../../../documentation/Cloud-Database-and-Cache/RDS/Introduction/Restrictions/SQLServer-Restrictions.md)<br>- 仅支持SQL Server */
func (c *RdsClient) SetInstanceName(request *rds.SetInstanceNameRequest) (*rds.SetInstanceNameResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &rds.SetInstanceNameResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 获取整个备份或备份中单个文件的下载链接。<br>- 当输入参数中有文件名时，获取该文件的下载链接。<br>- 输入参数中无文件名时，获取整个备份的下载链接。<br>由于备份机制的差异，使用该接口下载备份时，SQL Server必须输入文件名，每个文件逐一下载，不支持下载整个备份。SQL Server备份中的文件名（不包括后缀）即为备份的数据库名。例如文件名为my_test_db.bak，表示该文件是my_test_db数据库的备份。<br>MySQL可下载整个备份集，但不支持单个文件的下载。<br>- 仅支持SQL Server */
func (c *RdsClient) DescribeBackupDownloadURL(request *rds.DescribeBackupDownloadURLRequest) (*rds.DescribeBackupDownloadURLResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &rds.DescribeBackupDownloadURLResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 开启RDS实例的外网访问功能。开启后，用户可以通过internet访问RDS实例 */
func (c *RdsClient) EnableInternetAccess(request *rds.EnableInternetAccessRequest) (*rds.EnableInternetAccessResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &rds.EnableInternetAccessResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 获取当前账号下所有RDS实例及MySQL只读实例的概要信息，例如实例类型，版本，计费信息等 */
func (c *RdsClient) DescribeInstances(request *rds.DescribeInstancesRequest) (*rds.DescribeInstancesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &rds.DescribeInstancesResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 创建一个数据库。 为了实例的管理和数据恢复，RDS对用户权限进行了限制，用户仅能通过控制台或本接口创建数据库 */
func (c *RdsClient) CreateDatabase(request *rds.CreateDatabaseRequest) (*rds.CreateDatabaseResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &rds.CreateDatabaseResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 根据源实例全量备份创建一个新实例，新实例的数据跟源实例在创建备份时的数据状态一样。<br>例如根据源实例A的一个全量备份“mybak”新建一个实例B，该备份是在“‘2018-8-18 03:23:54”创建的。那么新建实例B的数据状态跟实例A‘2018-8-18 03:23:54’的状态一致<br>- 仅支持MySQL */
func (c *RdsClient) CreateInstanceFromBackup(request *rds.CreateInstanceFromBackupRequest) (*rds.CreateInstanceFromBackupResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &rds.CreateInstanceFromBackupResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 查看RDS实例备份策略。根据数据库类型的不同，支持的备份策略也略有差异，具体请看返回参数中的详细说明 */
func (c *RdsClient) GetBackupPolicy(request *rds.GetBackupPolicyRequest) (*rds.GetBackupPolicyResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &rds.GetBackupPolicyResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 获取用户通过单库上云工具上传到该实例上的文件列表<br>- 仅支持SQL Server */
func (c *RdsClient) DescribeImportFiles(request *rds.DescribeImportFilesRequest) (*rds.DescribeImportFilesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &rds.DescribeImportFilesResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 创建一个RDS实例全量备份，可以对整个实例或者部分数据库（仅SQL Server支持）进行全量备份。同一时间点，只能有一个正在运行的备份任务 */
func (c *RdsClient) CreateBackup(request *rds.CreateBackupRequest) (*rds.CreateBackupResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &rds.CreateBackupResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 根据用户定义的查询条件，获取SQL执行的性能统计信息，例如慢SQL等。用户可以根据这些信息查找与SQL执行相关的性能瓶颈，并进行优化。<br>- 仅支持SQL Server */
func (c *RdsClient) DescribeQueryPerformance(request *rds.DescribeQueryPerformanceRequest) (*rds.DescribeQueryPerformanceResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &rds.DescribeQueryPerformanceResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 获取当前实例下的所有审计结果文件的列表<br>- 仅支持SQL Server */
func (c *RdsClient) GetAuditFiles(request *rds.GetAuditFilesRequest) (*rds.GetAuditFilesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &rds.GetAuditFilesResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 重置数据库账号密码。如果用户忘记账号的密码，可以使用该接口重置指定账号密码。密码重置后，以前的密码将无法使用，必须使用重置后的新密码登录或连接数据库实例。 */
func (c *RdsClient) ResetPassword(request *rds.ResetPasswordRequest) (*rds.ResetPasswordResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &rds.ResetPasswordResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 从RDS实例中删除数据库。为便于管理和数据恢复，RDS对用户权限进行了控制，用户仅能通过控制台或本接口删除数据库 [MFA enabled] */
func (c *RdsClient) DeleteDatabase(request *rds.DeleteDatabaseRequest) (*rds.DeleteDatabaseResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &rds.DeleteDatabaseResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 获取SQL Server 错误日志及下载信息<br>- 仅支持SQL Server */
func (c *RdsClient) DescribeErrorLogs(request *rds.DescribeErrorLogsRequest) (*rds.DescribeErrorLogsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &rds.DescribeErrorLogsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 查看RDS实例当前白名单。白名单是允许访问当前实例的IP/IP段列表，缺省情况下，白名单对本VPC开放。如果用户开启了外网访问的功能，还需要对外网的IP配置白名单。 */
func (c *RdsClient) DescribeWhiteList(request *rds.DescribeWhiteListRequest) (*rds.DescribeWhiteListResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &rds.DescribeWhiteListResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 关闭RDS实例的外网访问功能。关闭后，用户无法通过Internet访问RDS，但可以在京东云内网通过内网域名访问 */
func (c *RdsClient) DisableInternetAccess(request *rds.DisableInternetAccessRequest) (*rds.DisableInternetAccessResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &rds.DisableInternetAccessResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 授予账号的数据库访问权限，即该账号对数据库拥有什么权限。一个账号可以对多个数据库具有访问权限。<br>为便于管理，RDS对权限进行了归类，目前提供以下两种权限<br>- ro：只读权限，用户只能读取数据库中的数据，不能进行创建、插入、删除、更改等操作。<br>- rw：读写权限，用户可以对数据库进行增删改查等操作 */
func (c *RdsClient) GrantPrivilege(request *rds.GrantPrivilegeRequest) (*rds.GrantPrivilegeResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &rds.GrantPrivilegeResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 获取当前实例的所有数据库详细信息的列表 */
func (c *RdsClient) DescribeDatabases(request *rds.DescribeDatabasesRequest) (*rds.DescribeDatabasesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &rds.DescribeDatabasesResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 查看当前实例已开启的审计选项。如当前实例未开启审计，则返回空<br>- 仅支持SQL Server */
func (c *RdsClient) DescribeAudit(request *rds.DescribeAuditRequest) (*rds.DescribeAuditResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &rds.DescribeAuditResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 获取当前系统所支持的各种数据库版本的审计选项及相应的推荐选项<br>- 仅支持SQL Server */
func (c *RdsClient) GetAuditOptions(request *rds.GetAuditOptionsRequest) (*rds.GetAuditOptionsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &rds.GetAuditOptionsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 从用户通过单库上云工具上传到云上的备份文件中恢复单个数据库<br>- 仅支持SQL Server */
func (c *RdsClient) RestoreDatabaseFromFile(request *rds.RestoreDatabaseFromFileRequest) (*rds.RestoreDatabaseFromFileResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &rds.RestoreDatabaseFromFileResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 设置或取消上传文件是否共享给同一账号下的其他实例。缺省情况下，文件仅在上传的实例上可见并可导入，其他实例不可见不可导入。如果需要该文件在其他实例上也可导入，可将此文件设置为共享<br>- 仅支持SQL Server */
func (c *RdsClient) SetImportFileShared(request *rds.SetImportFileSharedRequest) (*rds.SetImportFileSharedResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &rds.SetImportFileSharedResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 根据用户定义的查询条件，获取索引性能的统计信息，并提供缺失索引及索引创建建议。用户可以根据这些信息查找与索引相关的性能瓶颈，并进行优化。<br>- 仅支持SQL Server */
func (c *RdsClient) DescribeIndexPerformance(request *rds.DescribeIndexPerformanceRequest) (*rds.DescribeIndexPerformanceResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &rds.DescribeIndexPerformanceResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 删除数据库账号，账号删除后不可恢复，用户无法再使用该账号登录RDS实例 */
func (c *RdsClient) DeleteAccount(request *rds.DeleteAccountRequest) (*rds.DeleteAccountResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &rds.DeleteAccountResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 从备份中恢复单个数据库，支持从其他实例（但必须是同一个账号下的实例）备份中恢复。例如可以从生产环境的数据库实例的备份恢复到测试环境的数据库中。<br>- 仅支持SQL Server */
func (c *RdsClient) RestoreDatabaseFromBackup(request *rds.RestoreDatabaseFromBackupRequest) (*rds.RestoreDatabaseFromBackupResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &rds.RestoreDatabaseFromBackupResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 查看某个RDS实例下所有账号信息，包括账号名称、对各个数据库的访问权限信息等 */
func (c *RdsClient) DescribeAccounts(request *rds.DescribeAccountsRequest) (*rds.DescribeAccountsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &rds.DescribeAccountsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 创建数据库账号，用户可以使用客户端，应用程序等通过该账号和密码登录RDS数据库实例。<br>为便于管理和恢复，RDS对账号进行了限制，数据库账号只能通过控制台或者OpenAPI进行创建、删除账号以及对账号授权等，用户不能通过SQL语句对账号进行相关操作。 */
func (c *RdsClient) CreateAccount(request *rds.CreateAccountRequest) (*rds.CreateAccountResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &rds.CreateAccountResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 创建一个RDS实例，用户可以使用相应的数据库客户端或者应用程序通过域名和端口链接到该RDS实例上，进行操作。 */
func (c *RdsClient) CreateInstance(request *rds.CreateInstanceRequest) (*rds.CreateInstanceResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &rds.CreateInstanceResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 修改允许访问实例的IP白名单。白名单是允许访问当前实例的IP/IP段列表，缺省情况下，白名单对本VPC开放。如果用户开启了外网访问的功能，还需要对外网的IP配置白名单。 */
func (c *RdsClient) ModifyWhiteList(request *rds.ModifyWhiteListRequest) (*rds.ModifyWhiteListResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &rds.ModifyWhiteListResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 实例扩容，支持升级实例的CPU，内存及磁盘。目前暂不支持实例降配<br>- 仅支持MySQL */
func (c *RdsClient) ModifyInstanceSpec(request *rds.ModifyInstanceSpecRequest) (*rds.ModifyInstanceSpecResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &rds.ModifyInstanceSpecResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 查看该RDS实例下所有备份的详细信息，返回的备份列表按照备份开始时间（backupStartTime）降序排列。 */
func (c *RdsClient) DescribeBackups(request *rds.DescribeBackupsRequest) (*rds.DescribeBackupsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &rds.DescribeBackupsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 修改当前的审计选项。当前已有审计选项可以通过describeAudit获得，支持的全部选项可以通过getAuditOptions获得。<br>- 仅支持SQL Server */
func (c *RdsClient) ModifyAudit(request *rds.ModifyAuditRequest) (*rds.ModifyAuditResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &rds.ModifyAuditResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 重启RDS实例，例如修改了一些配置参数后，需要重启实例才能生效。可以结合主备切换的功能，轮流重启备机，降低对业务的影响<br>**注意：如果实例正在进行备份，那么重启主实例将会终止备份操作。** 可以查看备份策略中的备份开始时间确认是否有备份正在运行。如果确实需要在实例备份时重启主实例，建议重启后，手工进行一次实例的全备。 */
func (c *RdsClient) RebootInstance(request *rds.RebootInstanceRequest) (*rds.RebootInstanceResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &rds.RebootInstanceResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 修改RDS实例备份策略，目前仅支持用户修改“自动备份开始时间窗口”这个参数，其他参数暂不开放修改 */
func (c *RdsClient) SetBackupPolicy(request *rds.SetBackupPolicyRequest) (*rds.SetBackupPolicyResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &rds.SetBackupPolicyResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 查询RDS实例（MySQL、SQL Server等）的详细信息以及MySQL只读实例详细信息 */
func (c *RdsClient) DescribeInstanceAttributes(request *rds.DescribeInstanceAttributesRequest) (*rds.DescribeInstanceAttributesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &rds.DescribeInstanceAttributesResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 取消该账号对某个数据库的所有权限。权限取消后，该账号将不能访问此数据库。取消账号对某个数据库的访问权限，不影响该账号对其他数据库的访问权限<br>- 仅支持MySQL */
func (c *RdsClient) RevokePrivilege(request *rds.RevokePrivilegeRequest) (*rds.RevokePrivilegeResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &rds.RevokePrivilegeResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 从上传到OSS的备份文件中恢复单个数据库<br>- 仅支持SQL Server */
func (c *RdsClient) RestoreDatabaseFromOSS(request *rds.RestoreDatabaseFromOSSRequest) (*rds.RestoreDatabaseFromOSSResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &rds.RestoreDatabaseFromOSSResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

