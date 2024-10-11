use warp::Filter;
use std::process::Child;

mod func;
mod config;
mod server;

#[tokio::main]
async fn main() {
    let port = func::lib_loader::find_free_port_dll();
    let port_web = func::lib_loader::find_free_port_dll();
    let _ = func::func_shell::write_config_web(&port_web.to_string(), &port.to_string());
    let _ = func::func_shell::write_article(&port_web.to_string());

    let url_port = {
        let path = format!("http://localhost:{}/", port_web);
        path
    };
    
    let mut child: Option<Child> = None;
    const CONFIG_PATH: &str = "data/main_config.json";

    // Виклик з нового модуля
    if let Err(e) = func::func_shell::update_visualization(CONFIG_PATH, &url_port) {
        eprintln!("Помилка при оновленні конфігурації: {}", e);
    }

    // Виклик з нового модуля
    match func::func_shell::load_config(CONFIG_PATH) {
        Ok(config) => {
            if config.visualization == 1 {
                child = match func::func_shell::start_acwa() {
                    Ok(child_process) => Some(child_process),
                    Err(_e) => {
                        return;
                    }
                };
            }
        }
        Err(e) => {
            eprintln!("Помилка при завантаженні конфігурації: {}", e);
        }
    }

    let index_route = server::routes::index_route();
    let logs_post = server::post::logs_post();
    let config_post = server::post::config_post();
    let change_config_post = server::post::change_config_post();
    let logs_post_masege = server::post::log_post_message();

    let routes = index_route.or(logs_post).or(config_post).or(change_config_post).or(logs_post_masege);

    tokio::spawn(async move {
        warp::serve(routes).run(([127, 0, 0, 1], port as u16)).await;
    });

    if let Some(mut child_process) = child {
        if let Err(e) = child_process.wait() {
            eprintln!("Помилка при завершенні ACWA.exe: {}", e);
        }

        println!("ACWA.exe завершено.");
    } else {
        println!("ACWA.exe не було запущено.");
    }

    println!("Отриманий вільний порт: {}", port);
}
