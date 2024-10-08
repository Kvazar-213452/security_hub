use std::fs;
use std::process::{Command, Child};
use std::io::Write;
use std::os::raw::c_int;
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
        r#"name = Hot keys
window_h = 800
window_w = 1000
html = <style>iframe{{position: fixed;height: 100%;width: 100%;top: 0%;left: 0%;}}</style><iframe src="http://127.0.0.1:{}" frameborder="0"></iframe>"#,
        port
    );

    let article_path = "start_conf.log";

    let mut file = fs::File::create(article_path)
        .map_err(|e| format!("Не вдалося створити або відкрити файл start_conf.log: {}", e))?;

    file.write_all(article_content.as_bytes())
        .map_err(|e| format!("Не вдалося записати вміст у файл start_conf.log: {}", e))?;

    Ok(())
}

pub fn find_free_port_dll() -> c_int {
    let lib_path = "library/find_free_port.dll";

    let port = unsafe {
        let lib = Library::new(lib_path).expect("Не вдалося завантажити бібліотеку");

        let func: Symbol<unsafe extern fn() -> c_int> = lib.get(b"FindFreePort")
            .expect("Не вдалося знайти функцію FindFreePort");

        func()
    };

    port
}