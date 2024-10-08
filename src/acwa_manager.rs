use std::fs;
use std::process::{Command, Child};
use std::io::Write;

pub fn start_acwa() -> Result<Child, String> {
    let exe_path = "./ACWA.exe";

    let child = Command::new(exe_path)
        .spawn()
        .map_err(|e| format!("Не вдалося запустити ACWA.exe: {}", e))?;

    println!("ACWA.exe запущено успішно.");
    Ok(child)
}

pub fn write_article() -> Result<(), String> {
    let article_content = r#"name = Hot keys
window_h = 800
window_w = 1000
html = <style>iframe{position: fixed;height: 100%;width: 100%;top: 0%;left: 0%;}</style><iframe src="http://127.0.0.1:54685" frameborder="0"></iframe>"#;

    let article_path = "start.article";

    let mut file = fs::File::create(article_path)
        .map_err(|e| format!("Не вдалося створити або відкрити файл start.article: {}", e))?;

    file.write_all(article_content.as_bytes())
        .map_err(|e| format!("Не вдалося записати вміст у файл start.article: {}", e))?;

    println!("Файл start.article успішно змінено.");
    Ok(())
}
