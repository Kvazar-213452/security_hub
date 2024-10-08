use std::fs;
use std::process::{Command, Child};
use std::io::Write;
use std::os::raw::c_int;
use std::io;
use libloading::{Library, Symbol};

use crate::config;

pub fn start_acwa() -> Result<Child, String> {
    let child = Command::new(config::SHELL_WEB_CORE)
        .spawn()
        .map_err(|e| format!("Не вдалося запустити ACWA.exe: {}", e))?;

    println!("ACWA.exe запущено успішно.");
    Ok(child)
}

pub fn write_article(port: &str) -> Result<(), String> {
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

pub fn find_free_port_dll() -> c_int {
    let port = unsafe {
        let lib = Library::new(config::LIBRARY_PORT).expect("Не вдалося завантажити бібліотеку");

        let func: Symbol<unsafe extern fn() -> c_int> = lib.get(b"FindFreePort")
            .expect("Не вдалося знайти функцію FindFreePort");

        func()
    };

    port
}

pub fn write_config(port: &str) -> io::Result<()> {
    let config_content = format!(r#"port = {}"#, port);

    fs::create_dir_all("web")?;

    let mut file = fs::File::create(config::CONFIG_WEB)?;
    file.write_all(config_content.as_bytes())?;

    Ok(())
}