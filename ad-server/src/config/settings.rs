use std::path::PathBuf;
use crate::errors::{ApplicationError, ApplicationResult};

/// 命令行参数配置
///
/// 通过 `clap` 库解析命令行参数，支持以下选项：
///
/// ```bash
/// router -f config/custom.toml
/// router --config-path config/custom.toml
/// ```
#[derive(clap::Parser, Default)]
pub struct CmdLineConf {
    /// 配置文件路径
    ///
    /// 如果未指定此选项，应用程序将查找 `config/config.toml`
    /// 短选项：`-f`
    /// 长选项：`--config-path`
    #[arg(short = 'f', long, value_name = "FILE")]
    pub config_path: Option<PathBuf>,
}



// pub fn with_config_path(config_path: Option<PathBuf>)->ApplicationResult<Self> {
//     
// }