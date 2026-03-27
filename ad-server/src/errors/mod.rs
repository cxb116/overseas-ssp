use std::io;

/// 应用级错误类型
#[derive(Debug, thiserror::Error)]
pub enum ApplicationError {
    /// 配置文件加载错误
    #[error("配置加载失败: {0}")]
    ConfigError(String),

    /// 数据库连接错误
    #[error("数据库连接失败: {0}")]
    DatabaseError(String),

    /// IO 错误（文件读写等）
    #[error("IO 错误: {0}")]
    IoError(#[from] io::Error),

    /// HTTP 客户端错误
    #[error("HTTP 客户端错误: {0}")]
    HttpClientError(String),

    /// 初始化失败
    #[error("应用初始化失败: {0}")]
    InitializationError(String),
}

/// 应用级结果类型
pub type ApplicationResult<T> = Result<T, ApplicationError>;
