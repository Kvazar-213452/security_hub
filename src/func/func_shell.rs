use std::fs;
use std::process::{Command, Child};
use std::io::Write;
use std::io;
use std::fs::OpenOptions;
use chrono::Local;
use serde::{Deserialize, Serialize};
use serde_json::Error;
use std::fs::File;
use std::io::BufReader;

use crate::config;

pub fn start_acwa() -> std::result::Result<Child, String> {
    let child = Command::new(config::SHELL_WEB_CORE)
        .spawn()
        .map_err(|e| format!("Не вдалося запустити ACWA.exe: {}", e))?;

    println!("ACWA.exe запущено успішно.");
    Ok(child)
}

pub fn write_article(port: &str) -> std::result::Result<(), String> {
    let article_content = format!(
        r#"name = {}
window_h = {}
window_w = {}
html = <style>iframe{{position: fixed;height: 100%;width: 100%;top: 0%;left: 0%;}}</style><iframe src="http://127.0.0.1:{}" frameborder="0"></iframe>"#,
        config::NAME_WINDOWS_CORE,
        config::H_WINDOWS_CORE,
        config::W_WINDOWS_CORE,
        port
    );

    let mut file = fs::File::create(config::START_FILE_CORE)
        .map_err(|e| format!("Не вдалося створити або відкрити файл start_conf.log: {}", e))?;

    file.write_all(article_content.as_bytes())
        .map_err(|e| format!("Не вдалося записати вміст у файл start_conf.log: {}", e))?;

    Ok(())
}

pub fn write_config_web(port: &str, port_main: &str) -> io::Result<()> {
    let config_content = format!(r#"port = {}
port_main = {}"#, port, port_main);

    fs::create_dir_all("web")?;

    let mut file = fs::File::create(config::CONFIG_WEB)?;
    file.write_all(config_content.as_bytes())?;

    Ok(())
}

pub fn log_message_add(message: &str) -> io::Result<()> {
    let current_time = Local::now().format("%Y-%m-%d %H:%M:%S").to_string();
    
    let log_entry = format!("{} || {}\n", message, current_time);
    
    let mut file = OpenOptions::new()
        .write(true)
        .append(true)
        .create(true)
        .open(config::DATA_LOG)?;

    file.write_all(log_entry.as_bytes())?;
    
    Ok(())
}

#[derive(Serialize, Deserialize, Debug)]
pub struct MainConfig {
    pub visualization: u32,
    pub lang: String,
    pub url: String,
}

pub fn load_config(file_path: &str) -> std::result::Result<MainConfig, Error> {
    let file = File::open(file_path).expect("Не вдалося відкрити файл");
    let reader = BufReader::new(file);
    let config: MainConfig = serde_json::from_reader(reader)?;
    Ok(config)
}

pub fn update_visualization(file_path: &str, new_value: &str) -> std::result::Result<(), Error> {
    let mut config = load_config(file_path)?;

    config.url = new_value.to_string();

    let file = File::create(file_path).expect("Не вдалося створити файл");
    serde_json::to_writer(file, &config)?;
    Ok(())
}