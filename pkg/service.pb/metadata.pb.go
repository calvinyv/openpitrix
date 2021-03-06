// Code generated by protoc-gen-go. DO NOT EDIT.
// source: metadata.proto

package openpitrix

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Const struct {
	ApiPort                *int32  `protobuf:"varint,1,opt,name=api_port,json=apiPort,def=9100" json:"api_port,omitempty"`
	AppPort                *int32  `protobuf:"varint,2,opt,name=app_port,json=appPort,def=9101" json:"app_port,omitempty"`
	RuntimePort            *int32  `protobuf:"varint,3,opt,name=runtime_port,json=runtimePort,def=9102" json:"runtime_port,omitempty"`
	ClusterPort            *int32  `protobuf:"varint,4,opt,name=cluster_port,json=clusterPort,def=9103" json:"cluster_port,omitempty"`
	RepoPort               *int32  `protobuf:"varint,5,opt,name=repo_port,json=repoPort,def=9104" json:"repo_port,omitempty"`
	ApiHost                *string `protobuf:"bytes,11,opt,name=api_host,json=apiHost,def=openpitrix-api" json:"api_host,omitempty"`
	AppHost                *string `protobuf:"bytes,12,opt,name=app_host,json=appHost,def=openpitrix-app" json:"app_host,omitempty"`
	RuntimeHost            *string `protobuf:"bytes,13,opt,name=runtime_host,json=runtimeHost,def=openpitrix-runtime" json:"runtime_host,omitempty"`
	ClusterHost            *string `protobuf:"bytes,14,opt,name=cluster_host,json=clusterHost,def=openpitrix-cluster" json:"cluster_host,omitempty"`
	RepoHost               *string `protobuf:"bytes,15,opt,name=repo_host,json=repoHost,def=openpitrix-repo" json:"repo_host,omitempty"`
	DbType                 *string `protobuf:"bytes,21,opt,name=db_type,json=dbType,def=mysql" json:"db_type,omitempty"`
	DbHost                 *string `protobuf:"bytes,22,opt,name=db_host,json=dbHost,def=openpitrix-db" json:"db_host,omitempty"`
	DbPort                 *int32  `protobuf:"varint,23,opt,name=db_port,json=dbPort,def=3306" json:"db_port,omitempty"`
	DbEncoding             *string `protobuf:"bytes,24,opt,name=db_encoding,json=dbEncoding,def=utf8" json:"db_encoding,omitempty"`
	DbEngine               *string `protobuf:"bytes,25,opt,name=db_engine,json=dbEngine,def=InnoDB" json:"db_engine,omitempty"`
	DbDbname               *string `protobuf:"bytes,26,opt,name=db_dbname,json=dbDbname,def=openpitrix" json:"db_dbname,omitempty"`
	DbAppTableName         *string `protobuf:"bytes,31,opt,name=db_app_table_name,json=dbAppTableName,def=app" json:"db_app_table_name,omitempty"`
	DbRuntimeTableName     *string `protobuf:"bytes,32,opt,name=db_runtime_table_name,json=dbRuntimeTableName,def=runtime" json:"db_runtime_table_name,omitempty"`
	DbClusterTableName     *string `protobuf:"bytes,33,opt,name=db_cluster_table_name,json=dbClusterTableName,def=cluster" json:"db_cluster_table_name,omitempty"`
	DbClusterNodeTableName *string `protobuf:"bytes,34,opt,name=db_cluster_node_table_name,json=dbClusterNodeTableName,def=cluster_node" json:"db_cluster_node_table_name,omitempty"`
	DbRepoTableName        *string `protobuf:"bytes,35,opt,name=db_repo_table_name,json=dbRepoTableName,def=repo" json:"db_repo_table_name,omitempty"`
	XXX_unrecognized       []byte  `json:"-"`
}

func (m *Const) Reset()                    { *m = Const{} }
func (m *Const) String() string            { return proto.CompactTextString(m) }
func (*Const) ProtoMessage()               {}
func (*Const) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

const Default_Const_ApiPort int32 = 9100
const Default_Const_AppPort int32 = 9101
const Default_Const_RuntimePort int32 = 9102
const Default_Const_ClusterPort int32 = 9103
const Default_Const_RepoPort int32 = 9104
const Default_Const_ApiHost string = "openpitrix-api"
const Default_Const_AppHost string = "openpitrix-app"
const Default_Const_RuntimeHost string = "openpitrix-runtime"
const Default_Const_ClusterHost string = "openpitrix-cluster"
const Default_Const_RepoHost string = "openpitrix-repo"
const Default_Const_DbType string = "mysql"
const Default_Const_DbHost string = "openpitrix-db"
const Default_Const_DbPort int32 = 3306
const Default_Const_DbEncoding string = "utf8"
const Default_Const_DbEngine string = "InnoDB"
const Default_Const_DbDbname string = "openpitrix"
const Default_Const_DbAppTableName string = "app"
const Default_Const_DbRuntimeTableName string = "runtime"
const Default_Const_DbClusterTableName string = "cluster"
const Default_Const_DbClusterNodeTableName string = "cluster_node"
const Default_Const_DbRepoTableName string = "repo"

func (m *Const) GetApiPort() int32 {
	if m != nil && m.ApiPort != nil {
		return *m.ApiPort
	}
	return Default_Const_ApiPort
}

func (m *Const) GetAppPort() int32 {
	if m != nil && m.AppPort != nil {
		return *m.AppPort
	}
	return Default_Const_AppPort
}

func (m *Const) GetRuntimePort() int32 {
	if m != nil && m.RuntimePort != nil {
		return *m.RuntimePort
	}
	return Default_Const_RuntimePort
}

func (m *Const) GetClusterPort() int32 {
	if m != nil && m.ClusterPort != nil {
		return *m.ClusterPort
	}
	return Default_Const_ClusterPort
}

func (m *Const) GetRepoPort() int32 {
	if m != nil && m.RepoPort != nil {
		return *m.RepoPort
	}
	return Default_Const_RepoPort
}

func (m *Const) GetApiHost() string {
	if m != nil && m.ApiHost != nil {
		return *m.ApiHost
	}
	return Default_Const_ApiHost
}

func (m *Const) GetAppHost() string {
	if m != nil && m.AppHost != nil {
		return *m.AppHost
	}
	return Default_Const_AppHost
}

func (m *Const) GetRuntimeHost() string {
	if m != nil && m.RuntimeHost != nil {
		return *m.RuntimeHost
	}
	return Default_Const_RuntimeHost
}

func (m *Const) GetClusterHost() string {
	if m != nil && m.ClusterHost != nil {
		return *m.ClusterHost
	}
	return Default_Const_ClusterHost
}

func (m *Const) GetRepoHost() string {
	if m != nil && m.RepoHost != nil {
		return *m.RepoHost
	}
	return Default_Const_RepoHost
}

func (m *Const) GetDbType() string {
	if m != nil && m.DbType != nil {
		return *m.DbType
	}
	return Default_Const_DbType
}

func (m *Const) GetDbHost() string {
	if m != nil && m.DbHost != nil {
		return *m.DbHost
	}
	return Default_Const_DbHost
}

func (m *Const) GetDbPort() int32 {
	if m != nil && m.DbPort != nil {
		return *m.DbPort
	}
	return Default_Const_DbPort
}

func (m *Const) GetDbEncoding() string {
	if m != nil && m.DbEncoding != nil {
		return *m.DbEncoding
	}
	return Default_Const_DbEncoding
}

func (m *Const) GetDbEngine() string {
	if m != nil && m.DbEngine != nil {
		return *m.DbEngine
	}
	return Default_Const_DbEngine
}

func (m *Const) GetDbDbname() string {
	if m != nil && m.DbDbname != nil {
		return *m.DbDbname
	}
	return Default_Const_DbDbname
}

func (m *Const) GetDbAppTableName() string {
	if m != nil && m.DbAppTableName != nil {
		return *m.DbAppTableName
	}
	return Default_Const_DbAppTableName
}

func (m *Const) GetDbRuntimeTableName() string {
	if m != nil && m.DbRuntimeTableName != nil {
		return *m.DbRuntimeTableName
	}
	return Default_Const_DbRuntimeTableName
}

func (m *Const) GetDbClusterTableName() string {
	if m != nil && m.DbClusterTableName != nil {
		return *m.DbClusterTableName
	}
	return Default_Const_DbClusterTableName
}

func (m *Const) GetDbClusterNodeTableName() string {
	if m != nil && m.DbClusterNodeTableName != nil {
		return *m.DbClusterNodeTableName
	}
	return Default_Const_DbClusterNodeTableName
}

func (m *Const) GetDbRepoTableName() string {
	if m != nil && m.DbRepoTableName != nil {
		return *m.DbRepoTableName
	}
	return Default_Const_DbRepoTableName
}

type Default struct {
	DbAdminName           *string `protobuf:"bytes,1,opt,name=db_admin_name,json=dbAdminName,def=root" json:"db_admin_name,omitempty"`
	DbAdminPassword       *string `protobuf:"bytes,2,opt,name=db_admin_password,json=dbAdminPassword,def=password" json:"db_admin_password,omitempty"`
	DbAppUserName         *string `protobuf:"bytes,11,opt,name=db_app_user_name,json=dbAppUserName,def=openpitrix-user-app" json:"db_app_user_name,omitempty"`
	DbAppUserPassword     *string `protobuf:"bytes,12,opt,name=db_app_user_password,json=dbAppUserPassword,def=openpitrix-user-app-password" json:"db_app_user_password,omitempty"`
	DbRuntimeUserName     *string `protobuf:"bytes,13,opt,name=db_runtime_user_name,json=dbRuntimeUserName,def=openpitrix-user-runtime" json:"db_runtime_user_name,omitempty"`
	DbRuntimeUserPassword *string `protobuf:"bytes,14,opt,name=db_runtime_user_password,json=dbRuntimeUserPassword,def=openpitrix-user-runtime-password" json:"db_runtime_user_password,omitempty"`
	DbClusterUserName     *string `protobuf:"bytes,15,opt,name=db_cluster_user_name,json=dbClusterUserName,def=openpitrix-user-cluster" json:"db_cluster_user_name,omitempty"`
	DbClusterUserPassword *string `protobuf:"bytes,16,opt,name=db_cluster_user_password,json=dbClusterUserPassword,def=openpitrix-user-cluster-password" json:"db_cluster_user_password,omitempty"`
	DbRepoUserName        *string `protobuf:"bytes,17,opt,name=db_repo_user_name,json=dbRepoUserName,def=openpitrix-user-repo" json:"db_repo_user_name,omitempty"`
	DbRepoUserPassword    *string `protobuf:"bytes,18,opt,name=db_repo_user_password,json=dbRepoUserPassword,def=openpitrix-user-repo-password" json:"db_repo_user_password,omitempty"`
	GlogLogToStderr       *bool   `protobuf:"varint,21,opt,name=glog_log_to_stderr,json=glogLogToStderr,def=0" json:"glog_log_to_stderr,omitempty"`
	GlogAlsoLogToStderr   *bool   `protobuf:"varint,22,opt,name=glog_also_log_to_stderr,json=glogAlsoLogToStderr,def=0" json:"glog_also_log_to_stderr,omitempty"`
	GlogStderrThreshold   *string `protobuf:"bytes,23,opt,name=glog_stderr_threshold,json=glogStderrThreshold,def=ERROR" json:"glog_stderr_threshold,omitempty"`
	GlogLogDir            *string `protobuf:"bytes,24,opt,name=glog_log_dir,json=glogLogDir,def=" json:"glog_log_dir,omitempty"`
	GlogLogBacktraceAt    *string `protobuf:"bytes,25,opt,name=glog_log_backtrace_at,json=glogLogBacktraceAt,def=" json:"glog_log_backtrace_at,omitempty"`
	GlogV                 *int32  `protobuf:"varint,26,opt,name=glog_v,json=glogV,def=0" json:"glog_v,omitempty"`
	GlogVModule           *string `protobuf:"bytes,27,opt,name=glog_v_module,json=glogVModule,def=" json:"glog_v_module,omitempty"`
	GlogCopyStandardLogTo *string `protobuf:"bytes,28,opt,name=glog_copy_standard_log_to,json=glogCopyStandardLogTo,def=" json:"glog_copy_standard_log_to,omitempty"`
	XXX_unrecognized      []byte  `json:"-"`
}

func (m *Default) Reset()                    { *m = Default{} }
func (m *Default) String() string            { return proto.CompactTextString(m) }
func (*Default) ProtoMessage()               {}
func (*Default) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

const Default_Default_DbAdminName string = "root"
const Default_Default_DbAdminPassword string = "password"
const Default_Default_DbAppUserName string = "openpitrix-user-app"
const Default_Default_DbAppUserPassword string = "openpitrix-user-app-password"
const Default_Default_DbRuntimeUserName string = "openpitrix-user-runtime"
const Default_Default_DbRuntimeUserPassword string = "openpitrix-user-runtime-password"
const Default_Default_DbClusterUserName string = "openpitrix-user-cluster"
const Default_Default_DbClusterUserPassword string = "openpitrix-user-cluster-password"
const Default_Default_DbRepoUserName string = "openpitrix-user-repo"
const Default_Default_DbRepoUserPassword string = "openpitrix-user-repo-password"
const Default_Default_GlogLogToStderr bool = false
const Default_Default_GlogAlsoLogToStderr bool = false
const Default_Default_GlogStderrThreshold string = "ERROR"
const Default_Default_GlogV int32 = 0

func (m *Default) GetDbAdminName() string {
	if m != nil && m.DbAdminName != nil {
		return *m.DbAdminName
	}
	return Default_Default_DbAdminName
}

func (m *Default) GetDbAdminPassword() string {
	if m != nil && m.DbAdminPassword != nil {
		return *m.DbAdminPassword
	}
	return Default_Default_DbAdminPassword
}

func (m *Default) GetDbAppUserName() string {
	if m != nil && m.DbAppUserName != nil {
		return *m.DbAppUserName
	}
	return Default_Default_DbAppUserName
}

func (m *Default) GetDbAppUserPassword() string {
	if m != nil && m.DbAppUserPassword != nil {
		return *m.DbAppUserPassword
	}
	return Default_Default_DbAppUserPassword
}

func (m *Default) GetDbRuntimeUserName() string {
	if m != nil && m.DbRuntimeUserName != nil {
		return *m.DbRuntimeUserName
	}
	return Default_Default_DbRuntimeUserName
}

func (m *Default) GetDbRuntimeUserPassword() string {
	if m != nil && m.DbRuntimeUserPassword != nil {
		return *m.DbRuntimeUserPassword
	}
	return Default_Default_DbRuntimeUserPassword
}

func (m *Default) GetDbClusterUserName() string {
	if m != nil && m.DbClusterUserName != nil {
		return *m.DbClusterUserName
	}
	return Default_Default_DbClusterUserName
}

func (m *Default) GetDbClusterUserPassword() string {
	if m != nil && m.DbClusterUserPassword != nil {
		return *m.DbClusterUserPassword
	}
	return Default_Default_DbClusterUserPassword
}

func (m *Default) GetDbRepoUserName() string {
	if m != nil && m.DbRepoUserName != nil {
		return *m.DbRepoUserName
	}
	return Default_Default_DbRepoUserName
}

func (m *Default) GetDbRepoUserPassword() string {
	if m != nil && m.DbRepoUserPassword != nil {
		return *m.DbRepoUserPassword
	}
	return Default_Default_DbRepoUserPassword
}

func (m *Default) GetGlogLogToStderr() bool {
	if m != nil && m.GlogLogToStderr != nil {
		return *m.GlogLogToStderr
	}
	return Default_Default_GlogLogToStderr
}

func (m *Default) GetGlogAlsoLogToStderr() bool {
	if m != nil && m.GlogAlsoLogToStderr != nil {
		return *m.GlogAlsoLogToStderr
	}
	return Default_Default_GlogAlsoLogToStderr
}

func (m *Default) GetGlogStderrThreshold() string {
	if m != nil && m.GlogStderrThreshold != nil {
		return *m.GlogStderrThreshold
	}
	return Default_Default_GlogStderrThreshold
}

func (m *Default) GetGlogLogDir() string {
	if m != nil && m.GlogLogDir != nil {
		return *m.GlogLogDir
	}
	return ""
}

func (m *Default) GetGlogLogBacktraceAt() string {
	if m != nil && m.GlogLogBacktraceAt != nil {
		return *m.GlogLogBacktraceAt
	}
	return ""
}

func (m *Default) GetGlogV() int32 {
	if m != nil && m.GlogV != nil {
		return *m.GlogV
	}
	return Default_Default_GlogV
}

func (m *Default) GetGlogVModule() string {
	if m != nil && m.GlogVModule != nil {
		return *m.GlogVModule
	}
	return ""
}

func (m *Default) GetGlogCopyStandardLogTo() string {
	if m != nil && m.GlogCopyStandardLogTo != nil {
		return *m.GlogCopyStandardLogTo
	}
	return ""
}

type DBAppTableSchema struct {
	Id               *string `protobuf:"bytes,1,opt,name=id,def=name:id, type:string, size:50, primarykey" json:"id,omitempty"`
	Name             *string `protobuf:"bytes,2,opt,name=name,def=name:name, type:string, size:50" json:"name,omitempty"`
	Description      *string `protobuf:"bytes,3,opt,name=description,def=name:description, type:string, size:1000" json:"description,omitempty"`
	RepoId           *string `protobuf:"bytes,4,opt,name=repo_id,json=repoId,def=name:repo_id, type:string, size:50" json:"repo_id,omitempty"`
	Created          *string `protobuf:"bytes,5,opt,name=created,def=name:created, type:datetime" json:"created,omitempty"`
	LastModified     *string `protobuf:"bytes,6,opt,name=last_modified,json=lastModified,def=name:last_modified, type:datetime" json:"last_modified,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DBAppTableSchema) Reset()                    { *m = DBAppTableSchema{} }
func (m *DBAppTableSchema) String() string            { return proto.CompactTextString(m) }
func (*DBAppTableSchema) ProtoMessage()               {}
func (*DBAppTableSchema) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

const Default_DBAppTableSchema_Id string = "name:id, type:string, size:50, primarykey"
const Default_DBAppTableSchema_Name string = "name:name, type:string, size:50"
const Default_DBAppTableSchema_Description string = "name:description, type:string, size:1000"
const Default_DBAppTableSchema_RepoId string = "name:repo_id, type:string, size:50"
const Default_DBAppTableSchema_Created string = "name:created, type:datetime"
const Default_DBAppTableSchema_LastModified string = "name:last_modified, type:datetime"

func (m *DBAppTableSchema) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return Default_DBAppTableSchema_Id
}

func (m *DBAppTableSchema) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return Default_DBAppTableSchema_Name
}

func (m *DBAppTableSchema) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return Default_DBAppTableSchema_Description
}

func (m *DBAppTableSchema) GetRepoId() string {
	if m != nil && m.RepoId != nil {
		return *m.RepoId
	}
	return Default_DBAppTableSchema_RepoId
}

func (m *DBAppTableSchema) GetCreated() string {
	if m != nil && m.Created != nil {
		return *m.Created
	}
	return Default_DBAppTableSchema_Created
}

func (m *DBAppTableSchema) GetLastModified() string {
	if m != nil && m.LastModified != nil {
		return *m.LastModified
	}
	return Default_DBAppTableSchema_LastModified
}

type DBRuntimeTableSchema struct {
	Id               *string `protobuf:"bytes,1,opt,name=id,def=name:id, type:string, size:50, primarykey" json:"id,omitempty"`
	Name             *string `protobuf:"bytes,2,opt,name=name,def=name:name, type:string, size:50" json:"name,omitempty"`
	Description      *string `protobuf:"bytes,3,opt,name=description,def=name:description, type:string, size:1000" json:"description,omitempty"`
	Url              *string `protobuf:"bytes,4,opt,name=url,def=name:url, type:string, size:255" json:"url,omitempty"`
	Created          *string `protobuf:"bytes,5,opt,name=created,def=name:created, type:datetime" json:"created,omitempty"`
	LastModified     *string `protobuf:"bytes,6,opt,name=last_modified,json=lastModified,def=name:last_modified, type:datetime" json:"last_modified,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DBRuntimeTableSchema) Reset()                    { *m = DBRuntimeTableSchema{} }
func (m *DBRuntimeTableSchema) String() string            { return proto.CompactTextString(m) }
func (*DBRuntimeTableSchema) ProtoMessage()               {}
func (*DBRuntimeTableSchema) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

const Default_DBRuntimeTableSchema_Id string = "name:id, type:string, size:50, primarykey"
const Default_DBRuntimeTableSchema_Name string = "name:name, type:string, size:50"
const Default_DBRuntimeTableSchema_Description string = "name:description, type:string, size:1000"
const Default_DBRuntimeTableSchema_Url string = "name:url, type:string, size:255"
const Default_DBRuntimeTableSchema_Created string = "name:created, type:datetime"
const Default_DBRuntimeTableSchema_LastModified string = "name:last_modified, type:datetime"

func (m *DBRuntimeTableSchema) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return Default_DBRuntimeTableSchema_Id
}

func (m *DBRuntimeTableSchema) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return Default_DBRuntimeTableSchema_Name
}

func (m *DBRuntimeTableSchema) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return Default_DBRuntimeTableSchema_Description
}

func (m *DBRuntimeTableSchema) GetUrl() string {
	if m != nil && m.Url != nil {
		return *m.Url
	}
	return Default_DBRuntimeTableSchema_Url
}

func (m *DBRuntimeTableSchema) GetCreated() string {
	if m != nil && m.Created != nil {
		return *m.Created
	}
	return Default_DBRuntimeTableSchema_Created
}

func (m *DBRuntimeTableSchema) GetLastModified() string {
	if m != nil && m.LastModified != nil {
		return *m.LastModified
	}
	return Default_DBRuntimeTableSchema_LastModified
}

type DBClusterTableSchema struct {
	Id               *string `protobuf:"bytes,1,opt,name=id,def=name:id, type:string, size:50, primarykey" json:"id,omitempty"`
	Name             *string `protobuf:"bytes,2,opt,name=name,def=name:name, type:string, size:50" json:"name,omitempty"`
	Description      *string `protobuf:"bytes,3,opt,name=description,def=name:description, type:string, size:1000" json:"description,omitempty"`
	AppId            *string `protobuf:"bytes,4,opt,name=app_id,json=appId,def=name:app_id, type:string, size:50" json:"app_id,omitempty"`
	AppVersion       *string `protobuf:"bytes,5,opt,name=app_version,json=appVersion,def=name:app_version, type:string, size:50" json:"app_version,omitempty"`
	Status           *string `protobuf:"bytes,6,opt,name=status,def=name:status, type:string, size:50" json:"status,omitempty"`
	Created          *string `protobuf:"bytes,7,opt,name=created,def=name:created, type:datetime" json:"created,omitempty"`
	LastModified     *string `protobuf:"bytes,8,opt,name=last_modified,json=lastModified,def=name:last_modified, type:datetime" json:"last_modified,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DBClusterTableSchema) Reset()                    { *m = DBClusterTableSchema{} }
func (m *DBClusterTableSchema) String() string            { return proto.CompactTextString(m) }
func (*DBClusterTableSchema) ProtoMessage()               {}
func (*DBClusterTableSchema) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{4} }

const Default_DBClusterTableSchema_Id string = "name:id, type:string, size:50, primarykey"
const Default_DBClusterTableSchema_Name string = "name:name, type:string, size:50"
const Default_DBClusterTableSchema_Description string = "name:description, type:string, size:1000"
const Default_DBClusterTableSchema_AppId string = "name:app_id, type:string, size:50"
const Default_DBClusterTableSchema_AppVersion string = "name:app_version, type:string, size:50"
const Default_DBClusterTableSchema_Status string = "name:status, type:string, size:50"
const Default_DBClusterTableSchema_Created string = "name:created, type:datetime"
const Default_DBClusterTableSchema_LastModified string = "name:last_modified, type:datetime"

func (m *DBClusterTableSchema) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return Default_DBClusterTableSchema_Id
}

func (m *DBClusterTableSchema) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return Default_DBClusterTableSchema_Name
}

func (m *DBClusterTableSchema) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return Default_DBClusterTableSchema_Description
}

func (m *DBClusterTableSchema) GetAppId() string {
	if m != nil && m.AppId != nil {
		return *m.AppId
	}
	return Default_DBClusterTableSchema_AppId
}

func (m *DBClusterTableSchema) GetAppVersion() string {
	if m != nil && m.AppVersion != nil {
		return *m.AppVersion
	}
	return Default_DBClusterTableSchema_AppVersion
}

func (m *DBClusterTableSchema) GetStatus() string {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return Default_DBClusterTableSchema_Status
}

func (m *DBClusterTableSchema) GetCreated() string {
	if m != nil && m.Created != nil {
		return *m.Created
	}
	return Default_DBClusterTableSchema_Created
}

func (m *DBClusterTableSchema) GetLastModified() string {
	if m != nil && m.LastModified != nil {
		return *m.LastModified
	}
	return Default_DBClusterTableSchema_LastModified
}

type DBClusterNodeTableSchema struct {
	Id               *string `protobuf:"bytes,1,opt,name=id,def=name:id, type:string, size:50, primarykey" json:"id,omitempty"`
	Name             *string `protobuf:"bytes,2,opt,name=name,def=name:name, type:string, size:50" json:"name,omitempty"`
	Description      *string `protobuf:"bytes,3,opt,name=description,def=name:description, type:string, size:1000" json:"description,omitempty"`
	ClusterId        *string `protobuf:"bytes,4,opt,name=cluster_id,json=clusterId,def=name:cluster_id, type:string, size:50" json:"cluster_id,omitempty"`
	PrivateIp        *string `protobuf:"bytes,5,opt,name=private_ip,json=privateIp,def=name:private_ip, type:string, size:50" json:"private_ip,omitempty"`
	Status           *string `protobuf:"bytes,6,opt,name=status,def=name:status, type:string, size:50" json:"status,omitempty"`
	TransitionStatus *string `protobuf:"bytes,7,opt,name=transition_status,json=transitionStatus,def=name:transition_status, type:string, size:50" json:"transition_status,omitempty"`
	Created          *string `protobuf:"bytes,8,opt,name=created,def=name:created, type:datetime" json:"created,omitempty"`
	LastModified     *string `protobuf:"bytes,9,opt,name=last_modified,json=lastModified,def=name:last_modified, type:datetime" json:"last_modified,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DBClusterNodeTableSchema) Reset()                    { *m = DBClusterNodeTableSchema{} }
func (m *DBClusterNodeTableSchema) String() string            { return proto.CompactTextString(m) }
func (*DBClusterNodeTableSchema) ProtoMessage()               {}
func (*DBClusterNodeTableSchema) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{5} }

const Default_DBClusterNodeTableSchema_Id string = "name:id, type:string, size:50, primarykey"
const Default_DBClusterNodeTableSchema_Name string = "name:name, type:string, size:50"
const Default_DBClusterNodeTableSchema_Description string = "name:description, type:string, size:1000"
const Default_DBClusterNodeTableSchema_ClusterId string = "name:cluster_id, type:string, size:50"
const Default_DBClusterNodeTableSchema_PrivateIp string = "name:private_ip, type:string, size:50"
const Default_DBClusterNodeTableSchema_Status string = "name:status, type:string, size:50"
const Default_DBClusterNodeTableSchema_TransitionStatus string = "name:transition_status, type:string, size:50"
const Default_DBClusterNodeTableSchema_Created string = "name:created, type:datetime"
const Default_DBClusterNodeTableSchema_LastModified string = "name:last_modified, type:datetime"

func (m *DBClusterNodeTableSchema) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return Default_DBClusterNodeTableSchema_Id
}

func (m *DBClusterNodeTableSchema) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return Default_DBClusterNodeTableSchema_Name
}

func (m *DBClusterNodeTableSchema) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return Default_DBClusterNodeTableSchema_Description
}

func (m *DBClusterNodeTableSchema) GetClusterId() string {
	if m != nil && m.ClusterId != nil {
		return *m.ClusterId
	}
	return Default_DBClusterNodeTableSchema_ClusterId
}

func (m *DBClusterNodeTableSchema) GetPrivateIp() string {
	if m != nil && m.PrivateIp != nil {
		return *m.PrivateIp
	}
	return Default_DBClusterNodeTableSchema_PrivateIp
}

func (m *DBClusterNodeTableSchema) GetStatus() string {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return Default_DBClusterNodeTableSchema_Status
}

func (m *DBClusterNodeTableSchema) GetTransitionStatus() string {
	if m != nil && m.TransitionStatus != nil {
		return *m.TransitionStatus
	}
	return Default_DBClusterNodeTableSchema_TransitionStatus
}

func (m *DBClusterNodeTableSchema) GetCreated() string {
	if m != nil && m.Created != nil {
		return *m.Created
	}
	return Default_DBClusterNodeTableSchema_Created
}

func (m *DBClusterNodeTableSchema) GetLastModified() string {
	if m != nil && m.LastModified != nil {
		return *m.LastModified
	}
	return Default_DBClusterNodeTableSchema_LastModified
}

type DBRepoTableSchema struct {
	Id               *string `protobuf:"bytes,1,opt,name=id,def=name:id, type:string, size:50, primarykey" json:"id,omitempty"`
	Name             *string `protobuf:"bytes,2,opt,name=name,def=name:name, type:string, size:50" json:"name,omitempty"`
	Description      *string `protobuf:"bytes,3,opt,name=description,def=name:description, type:string, size:1000" json:"description,omitempty"`
	Url              *string `protobuf:"bytes,4,opt,name=url,def=name:url, type:string, size:255" json:"url,omitempty"`
	Created          *string `protobuf:"bytes,5,opt,name=created,def=name:created, type:datetime" json:"created,omitempty"`
	LastModified     *string `protobuf:"bytes,6,opt,name=last_modified,json=lastModified,def=name:last_modified, type:datetime" json:"last_modified,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DBRepoTableSchema) Reset()                    { *m = DBRepoTableSchema{} }
func (m *DBRepoTableSchema) String() string            { return proto.CompactTextString(m) }
func (*DBRepoTableSchema) ProtoMessage()               {}
func (*DBRepoTableSchema) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{6} }

const Default_DBRepoTableSchema_Id string = "name:id, type:string, size:50, primarykey"
const Default_DBRepoTableSchema_Name string = "name:name, type:string, size:50"
const Default_DBRepoTableSchema_Description string = "name:description, type:string, size:1000"
const Default_DBRepoTableSchema_Url string = "name:url, type:string, size:255"
const Default_DBRepoTableSchema_Created string = "name:created, type:datetime"
const Default_DBRepoTableSchema_LastModified string = "name:last_modified, type:datetime"

func (m *DBRepoTableSchema) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return Default_DBRepoTableSchema_Id
}

func (m *DBRepoTableSchema) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return Default_DBRepoTableSchema_Name
}

func (m *DBRepoTableSchema) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return Default_DBRepoTableSchema_Description
}

func (m *DBRepoTableSchema) GetUrl() string {
	if m != nil && m.Url != nil {
		return *m.Url
	}
	return Default_DBRepoTableSchema_Url
}

func (m *DBRepoTableSchema) GetCreated() string {
	if m != nil && m.Created != nil {
		return *m.Created
	}
	return Default_DBRepoTableSchema_Created
}

func (m *DBRepoTableSchema) GetLastModified() string {
	if m != nil && m.LastModified != nil {
		return *m.LastModified
	}
	return Default_DBRepoTableSchema_LastModified
}

func init() {
	proto.RegisterType((*Const)(nil), "openpitrix.Const")
	proto.RegisterType((*Default)(nil), "openpitrix.Default")
	proto.RegisterType((*DBAppTableSchema)(nil), "openpitrix.DBAppTableSchema")
	proto.RegisterType((*DBRuntimeTableSchema)(nil), "openpitrix.DBRuntimeTableSchema")
	proto.RegisterType((*DBClusterTableSchema)(nil), "openpitrix.DBClusterTableSchema")
	proto.RegisterType((*DBClusterNodeTableSchema)(nil), "openpitrix.DBClusterNodeTableSchema")
	proto.RegisterType((*DBRepoTableSchema)(nil), "openpitrix.DBRepoTableSchema")
}

func init() { proto.RegisterFile("metadata.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 1302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x57, 0x5d, 0x6f, 0x13, 0x47,
	0x14, 0x25, 0x24, 0xfe, 0x1a, 0xe7, 0xcb, 0x03, 0x01, 0xf3, 0xd5, 0x18, 0x53, 0x20, 0x48, 0x01,
	0x1c, 0x42, 0x2a, 0xb2, 0x54, 0xaa, 0x12, 0x4c, 0x4b, 0xaa, 0x42, 0xd1, 0x26, 0x45, 0xe2, 0x69,
	0x35, 0xeb, 0x99, 0x38, 0x23, 0xd6, 0x3b, 0xd3, 0xd9, 0x31, 0xad, 0xfb, 0x56, 0xa9, 0x4f, 0x7d,
	0xe9, 0x6f, 0xe9, 0x7f, 0xe9, 0x7f, 0xe9, 0x6b, 0x35, 0x77, 0x66, 0xc7, 0x1b, 0xc7, 0x7e, 0x01,
	0xa9, 0x95, 0x50, 0x1f, 0x1c, 0x39, 0xf7, 0x9e, 0x73, 0xe6, 0xdc, 0xb9, 0x77, 0x66, 0xd7, 0x68,
	0x79, 0xc0, 0x34, 0xa1, 0x44, 0x93, 0x07, 0x52, 0x09, 0x2d, 0x30, 0x12, 0x92, 0xa5, 0x92, 0x6b,
	0xc5, 0x7f, 0xbe, 0xba, 0xde, 0x17, 0xa2, 0x9f, 0xb0, 0x87, 0x90, 0x89, 0x87, 0xc7, 0x0f, 0x35,
	0x1f, 0xb0, 0x4c, 0x93, 0x81, 0xb4, 0xe0, 0xf6, 0x9f, 0x15, 0x54, 0x7a, 0x26, 0xd2, 0x4c, 0xe3,
	0x75, 0x54, 0x25, 0x92, 0x47, 0x52, 0x28, 0xdd, 0x9c, 0x6b, 0xcd, 0x6d, 0x94, 0x82, 0x85, 0xdd,
	0xad, 0x4e, 0x27, 0xac, 0x10, 0xc9, 0x5f, 0x0b, 0xe5, 0x00, 0xd2, 0x02, 0xce, 0x7b, 0xc0, 0x96,
	0x01, 0x48, 0x00, 0xdc, 0x45, 0x8b, 0x6a, 0x98, 0x9a, 0x15, 0x2c, 0x68, 0xde, 0x83, 0x1e, 0x85,
	0x75, 0x97, 0xc9, 0x81, 0xbd, 0x64, 0x98, 0x69, 0xa6, 0x2c, 0x70, 0xc1, 0x03, 0xb7, 0xc3, 0xba,
	0xcb, 0x00, 0xf0, 0x26, 0xaa, 0x29, 0x26, 0x85, 0x45, 0x95, 0x3c, 0xea, 0x71, 0x58, 0x35, 0x61,
	0x80, 0xdc, 0xb3, 0xb6, 0x4f, 0x44, 0xa6, 0x9b, 0xf5, 0xd6, 0xdc, 0x46, 0x2d, 0x58, 0x1e, 0xd7,
	0x7f, 0x9f, 0x48, 0x0e, 0x05, 0xbc, 0x10, 0x99, 0x83, 0x4a, 0x0b, 0x5d, 0x9c, 0x02, 0x95, 0x50,
	0x0a, 0x40, 0x77, 0xc6, 0xa5, 0x00, 0x7c, 0x09, 0xe0, 0xb8, 0x00, 0x77, 0x69, 0x5f, 0x58, 0x4e,
	0xcb, 0x0b, 0x03, 0xda, 0xf2, 0x19, 0x9a, 0x4b, 0xfb, 0x32, 0x81, 0xb6, 0xe9, 0xca, 0x04, 0xce,
	0x0a, 0x70, 0x56, 0x8a, 0x4b, 0x31, 0x29, 0x6c, 0xc5, 0x80, 0xfe, 0x0c, 0x55, 0x68, 0x1c, 0xe9,
	0x91, 0x64, 0xcd, 0x35, 0xc0, 0x96, 0x06, 0xa3, 0xec, 0xc7, 0x24, 0x2c, 0xd3, 0xf8, 0x68, 0x24,
	0x19, 0xbe, 0x03, 0x79, 0xd0, 0xba, 0x04, 0xf9, 0xa5, 0x82, 0x16, 0x8d, 0x0d, 0x0e, 0x74, 0x6e,
	0x00, 0x0e, 0xb6, 0xf6, 0xb2, 0xdd, 0xda, 0xed, 0xed, 0xce, 0x17, 0x26, 0x0d, 0x1b, 0x7b, 0x1b,
	0xd5, 0x69, 0x1c, 0xb1, 0xb4, 0x27, 0x28, 0x4f, 0xfb, 0xcd, 0x26, 0x48, 0x2d, 0x0c, 0xf5, 0xf1,
	0x93, 0x10, 0xd1, 0xf8, 0xb9, 0x8b, 0xe3, 0x5b, 0xa8, 0x06, 0xb0, 0x3e, 0x4f, 0x59, 0xf3, 0x0a,
	0x80, 0xca, 0x07, 0x69, 0x2a, 0xba, 0xfb, 0x61, 0xd5, 0xc0, 0x4c, 0x1c, 0xdf, 0x05, 0x10, 0x8d,
	0x53, 0x32, 0x60, 0xcd, 0xab, 0x00, 0x2a, 0x4c, 0xa9, 0x01, 0x76, 0x21, 0x87, 0x1f, 0xa0, 0x06,
	0x8d, 0x23, 0xd3, 0x25, 0x4d, 0xe2, 0x84, 0x45, 0x40, 0x58, 0x07, 0xc2, 0xbc, 0x69, 0xd0, 0x32,
	0x8d, 0xf7, 0xa4, 0x3c, 0x32, 0xb9, 0x57, 0x06, 0x1f, 0xa0, 0x35, 0x1a, 0x47, 0x79, 0xab, 0x0a,
	0x9c, 0x16, 0x70, 0x2a, 0x79, 0x97, 0x30, 0x8d, 0x43, 0xfb, 0x75, 0x92, 0x9b, 0xf7, 0xab, 0xc0,
	0xbd, 0x69, 0xb9, 0x79, 0xab, 0x30, 0x8d, 0x9f, 0xd9, 0xaf, 0x63, 0xee, 0x0b, 0x74, 0xb5, 0xc0,
	0x4d, 0x05, 0x3d, 0xb5, 0x78, 0x1b, 0x04, 0x16, 0x8b, 0xe9, 0xf0, 0x92, 0x57, 0x79, 0x25, 0x68,
	0xc1, 0xc5, 0x16, 0xc2, 0xa6, 0x02, 0xd3, 0xfe, 0x82, 0xc2, 0x2d, 0xbb, 0xdb, 0xd0, 0xf9, 0x15,
	0x1a, 0x87, 0x4c, 0x0a, 0x4f, 0x69, 0xff, 0x55, 0x45, 0x95, 0x2e, 0x3b, 0x26, 0xc3, 0x44, 0xe3,
	0x0d, 0xb4, 0x64, 0x36, 0x8c, 0x0e, 0x78, 0x6a, 0x99, 0x73, 0x8e, 0x29, 0x84, 0x0e, 0xeb, 0x34,
	0xde, 0x33, 0x19, 0x58, 0xe8, 0xb1, 0xdd, 0x5a, 0x40, 0x4a, 0x92, 0x65, 0x3f, 0x09, 0x45, 0xe1,
	0x1c, 0xd7, 0x82, 0x6a, 0xfe, 0xbf, 0x59, 0x0b, 0x18, 0xaf, 0x5d, 0x00, 0x7f, 0x89, 0x56, 0x5d,
	0x43, 0x86, 0x99, 0x29, 0xc5, 0x2c, 0x61, 0x8f, 0xd9, 0x85, 0xc2, 0x54, 0x99, 0x1c, 0x1c, 0xa0,
	0x25, 0xe8, 0xcf, 0x0f, 0x19, 0x53, 0xb0, 0xe6, 0x4b, 0x74, 0xb1, 0xc8, 0xf6, 0xcb, 0xda, 0xd3,
	0x77, 0x7d, 0x8a, 0xc2, 0x7d, 0x6f, 0xa5, 0xe1, 0xa5, 0xbc, 0x99, 0x17, 0x20, 0x97, 0x77, 0x7b,
	0x6c, 0xc8, 0x9e, 0xce, 0xcb, 0x93, 0x72, 0x79, 0xf3, 0x1b, 0xbe, 0xf9, 0xde, 0xd8, 0x5b, 0xd4,
	0x9c, 0x54, 0xf2, 0xe6, 0xec, 0xa1, 0x6d, 0xcd, 0x50, 0x1b, 0x1b, 0x5c, 0x3b, 0x25, 0x3b, 0x61,
	0x32, 0xef, 0xfd, 0xd8, 0xe4, 0xca, 0x74, 0x93, 0xf9, 0x94, 0x35, 0xfc, 0x7c, 0x4c, 0x98, 0x3c,
	0xa5, 0xe4, 0x4d, 0xae, 0x4e, 0x37, 0xe9, 0xc0, 0xa7, 0x4c, 0x16, 0x64, 0xbd, 0xc9, 0xaf, 0x60,
	0x18, 0x60, 0xea, 0xc6, 0x0e, 0x1b, 0xa0, 0x79, 0xf1, 0x4c, 0xe1, 0x66, 0x08, 0x97, 0xed, 0x10,
	0x7a, 0x6f, 0xaf, 0xed, 0xc1, 0xf3, 0x02, 0xde, 0x18, 0x06, 0x91, 0x1b, 0xd3, 0x44, 0xc6, 0xae,
	0xf0, 0x58, 0xcd, 0x5b, 0x7a, 0x84, 0x70, 0x3f, 0x11, 0xfd, 0xc8, 0x7c, 0xb4, 0x88, 0x32, 0x4d,
	0x99, 0x52, 0x70, 0xc3, 0x55, 0x83, 0xd2, 0x31, 0x49, 0x32, 0x16, 0xae, 0x18, 0xc0, 0x77, 0xa2,
	0x7f, 0x24, 0x0e, 0x21, 0x8b, 0x9f, 0xa2, 0xcb, 0xc0, 0x21, 0x49, 0x26, 0x26, 0x88, 0x97, 0x8a,
	0xc4, 0x0b, 0x06, 0xb5, 0x97, 0x64, 0xa2, 0x48, 0xde, 0x45, 0x6b, 0x40, 0xb6, 0x84, 0x48, 0x9f,
	0x28, 0x96, 0x9d, 0x88, 0x84, 0xc2, 0x6d, 0x58, 0x0b, 0x4a, 0xcf, 0xc3, 0xf0, 0xfb, 0xd0, 0x52,
	0x2d, 0xe3, 0x28, 0x47, 0xe0, 0x36, 0x5a, 0xf4, 0x5e, 0x29, 0x57, 0xee, 0x72, 0x3c, 0x17, 0x22,
	0x67, 0xb0, 0xcb, 0x15, 0xde, 0x76, 0xf2, 0xe6, 0x13, 0x93, 0xde, 0x3b, 0xad, 0x48, 0x8f, 0x45,
	0x44, 0xbb, 0x4b, 0xf2, 0x5c, 0x88, 0x1d, 0x78, 0x3f, 0x4f, 0xee, 0x69, 0xdc, 0x44, 0x65, 0x20,
	0xbd, 0x87, 0x5b, 0xb2, 0x14, 0xcc, 0x75, 0xc2, 0x92, 0x09, 0xbc, 0xc1, 0x9f, 0xa3, 0x25, 0x9b,
	0x89, 0x06, 0x82, 0x0e, 0x13, 0xd6, 0xbc, 0xe6, 0x64, 0xea, 0x90, 0x7f, 0x09, 0x41, 0x1c, 0xa0,
	0x2b, 0x80, 0xea, 0x09, 0x39, 0x8a, 0x32, 0x4d, 0x52, 0x4a, 0x14, 0x75, 0x3b, 0xd3, 0xbc, 0xee,
	0x18, 0xe0, 0xeb, 0x99, 0x90, 0xa3, 0x43, 0x07, 0x80, 0x5d, 0x69, 0xff, 0x3e, 0x8f, 0x56, 0xbb,
	0xfb, 0xf9, 0xf5, 0x7a, 0xd8, 0x3b, 0x61, 0x03, 0x82, 0x77, 0xd1, 0x79, 0x4e, 0xdd, 0xa5, 0x72,
	0xcf, 0x4c, 0x49, 0xc0, 0xe9, 0x66, 0xcb, 0x3c, 0x7c, 0x82, 0x4c, 0x2b, 0x9e, 0xf6, 0x37, 0x5b,
	0x19, 0xff, 0x85, 0x05, 0x3b, 0x9d, 0xcd, 0x96, 0x54, 0x7c, 0x40, 0xd4, 0xe8, 0x1d, 0x1b, 0x85,
	0xe7, 0x39, 0xc5, 0xdb, 0x68, 0x01, 0xc6, 0xca, 0xde, 0x31, 0xeb, 0x40, 0x36, 0x7f, 0xa6, 0xd3,
	0x43, 0x00, 0xe3, 0x6f, 0x51, 0x9d, 0xb2, 0xac, 0xa7, 0xb8, 0xd4, 0x5c, 0xa4, 0xf0, 0x0a, 0x51,
	0x0b, 0x36, 0x80, 0x5b, 0x88, 0x4f, 0x93, 0xd8, 0xea, 0x74, 0x3a, 0x61, 0x91, 0x8c, 0x9f, 0xa2,
	0x0a, 0x0c, 0x28, 0xa7, 0xf0, 0x86, 0x51, 0x0b, 0xda, 0xa0, 0xe3, 0x62, 0x33, 0x6c, 0x94, 0x4d,
	0xfa, 0x80, 0xe2, 0x1d, 0x54, 0xe9, 0x29, 0x46, 0x34, 0xa3, 0xf0, 0xe2, 0x51, 0x0b, 0xae, 0x01,
	0xd9, 0xc5, 0x1c, 0x99, 0x12, 0xcd, 0xe0, 0x8a, 0xc9, 0xb1, 0xf8, 0x6b, 0xb4, 0x94, 0x90, 0x4c,
	0x9b, 0x26, 0xf1, 0x63, 0xce, 0x68, 0xb3, 0x0c, 0xe4, 0x9b, 0x40, 0x3e, 0x95, 0x99, 0x94, 0x58,
	0x34, 0xd9, 0x97, 0x2e, 0xd9, 0xfe, 0x6d, 0x1e, 0x5d, 0xec, 0xee, 0x17, 0x9f, 0x59, 0x9f, 0x40,
	0x43, 0xb6, 0xd0, 0xfc, 0x50, 0x25, 0xae, 0x19, 0x76, 0xfd, 0xa1, 0x4a, 0xa6, 0x71, 0x1f, 0xed,
	0xec, 0x84, 0x06, 0xfb, 0x5f, 0xb7, 0xe1, 0x8f, 0x05, 0xd3, 0x86, 0xe2, 0xe3, 0xff, 0x13, 0x68,
	0xc3, 0x13, 0x54, 0x36, 0x8f, 0x64, 0x7f, 0x2c, 0xec, 0xae, 0xd8, 0xd0, 0x0c, 0x13, 0x25, 0x22,
	0xe5, 0x01, 0xc5, 0xdf, 0xa0, 0xba, 0x81, 0xbd, 0x67, 0x2a, 0x33, 0x2e, 0x6c, 0x47, 0xee, 0x78,
	0xba, 0x8b, 0xcf, 0xd0, 0x40, 0x44, 0xca, 0x37, 0x16, 0x81, 0x77, 0x51, 0x39, 0xd3, 0x44, 0x0f,
	0xb3, 0x53, 0x8d, 0xb1, 0xa1, 0x59, 0x07, 0xd3, 0x66, 0x8b, 0x13, 0x51, 0xf9, 0x98, 0x89, 0xa8,
	0x7e, 0xd8, 0x44, 0xfc, 0xbd, 0x80, 0x9a, 0x7e, 0x22, 0xfc, 0xab, 0xdc, 0x27, 0x30, 0x15, 0x5d,
	0x84, 0xf2, 0x57, 0x0d, 0x3f, 0x19, 0xb7, 0xed, 0xd6, 0xfa, 0xf0, 0x0c, 0x33, 0x35, 0x87, 0x38,
	0xa0, 0x46, 0x45, 0x2a, 0xfe, 0x9e, 0x68, 0x16, 0x71, 0xe9, 0x06, 0xc4, 0xaa, 0x8c, 0xc3, 0xb3,
	0x54, 0x1c, 0xe2, 0x40, 0x7e, 0xcc, 0x78, 0xbc, 0x45, 0x0d, 0xad, 0x48, 0x9a, 0x71, 0x53, 0x54,
	0xe4, 0x54, 0xec, 0xa0, 0x6c, 0x82, 0xca, 0x99, 0xec, 0x0c, 0xc1, 0xd5, 0x31, 0xf0, 0xf0, 0xcc,
	0xe4, 0x55, 0x3f, 0x66, 0xf2, 0x6a, 0x1f, 0x36, 0x79, 0xbf, 0xce, 0xa3, 0x46, 0x77, 0xdf, 0xff,
	0x14, 0xf8, 0xff, 0x79, 0xf0, 0xef, 0x3f, 0x0f, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0xf6, 0xfe,
	0x10, 0xbd, 0x6c, 0x11, 0x00, 0x00,
}
