use warp::Filter;

mod func;
mod config;

#[tokio::main]
async fn main() {
    let port = func::func_shell::find_free_port_dll();
    let _ = func::func_shell::write_article(&port.to_string());

    let mut child = match func::func_shell::start_acwa() {
        Ok(child) => child,
        Err(_e) => {
            return;
        }
    };

    let route = warp::path::end()
        .map(|| warp::reply::html(r#"
            <h1>Веб-сервер працює!</h1>
        "#));

    tokio::spawn(async move {
        warp::serve(route).run(([127, 0, 0, 1], port as u16)).await;
    });

    if let Err(e) = child.wait() {
        eprintln!("Помилка при завершенні ACWA.exe: {}", e);
    }

    println!("ACWA.exe завершено.");
    println!("Отриманий вільний порт: {}", port);
}