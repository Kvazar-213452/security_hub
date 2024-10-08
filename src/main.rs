use warp::Filter;

mod func;
mod config;
mod server;

#[tokio::main]
async fn main() {
    func::func_shell::log_message_add("Ddddw");
    let port = func::func_shell::find_free_port_dll();
    let port_web = func::func_shell::find_free_port_dll();
    let _ = func::func_shell::write_config_web(&port_web.to_string());
    let _ = func::func_shell::write_article(&port.to_string());

    let mut child = match func::func_shell::start_acwa() {
        Ok(child) => child,
        Err(_e) => {
            return;
        }
    };

    let index_route = server::routes::index_route();
    let logs_post = server::post::logs_post();
    let config_post = server::post::config_post();

    let routes = index_route.or(logs_post).or(config_post);

    tokio::spawn(async move {
        warp::serve(routes).run(([127, 0, 0, 1], port as u16)).await;
    });

    if let Err(e) = child.wait() {
        eprintln!("Помилка при завершенні ACWA.exe: {}", e);
    }

    println!("ACWA.exe завершено.");
    println!("Отриманий вільний порт: {}", port);
}