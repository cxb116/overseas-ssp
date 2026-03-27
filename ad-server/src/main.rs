mod errors;
mod config;

use errors::{ApplicationError, ApplicationResult};


/// ## 主要功能
/// - 加载配置文件
/// - 初始化日志系统
/// - 启动后台指标收集器
/// - 启动 HTTP 服务器并监听请求
///
/// ApplicationResult 是应用层的统一返回类型，用于包装可能失败的操作
fn main() -> ApplicationResult<()> {
    let args: Vec<String> = std::env::args().collect();

    let config = parse_config(&args)?;

    println!("Query: {}", config.query);
    println!("File path: {}", config.file_path);

    let contents = std::fs::read_to_string(&config.file_path)?;

    println!("File contents (first 100 chars): {}...",
             contents.chars().take(100).collect::<String>());
    println!("start successfully:");

    Ok(())
}

pub struct Config {
    query: String,
    file_path: String,
}

fn parse_config(args: &[String]) -> ApplicationResult<Config> {
    if args.len() < 3 {
        return Err(ApplicationError::ConfigError(
            format!("需要至少 3 个参数，但得到 {} 个。用法: {} <query> <file_path>",
                    args.len(), args[0])
        ));
    }

    let query = &args[1];
    let file_path = &args[2];

    Ok(Config {
        query: query.clone(),
        file_path: file_path.clone(),
    })
}