#![windows_subsystem = "windows"]

use warp::Filter;

mod acwa_manager;

#[tokio::main]
async fn main() {
    let mut child = match acwa_manager::start_acwa() {
        Ok(child) => child,
        Err(e) => {
            eprintln!("{}", e);
            return;
        }
    };

    let _ = acwa_manager::write_article();

    let route = warp::path::end()
        .map(|| warp::reply::html(r#"
            <html>
            <head><title>Rust Web Server</title></head>
            <body>
            <h1>Веб-сервер працює!</h1>
            <iframe src="http://127.0.0.1:54685" style="position:fixed;height:100%;width:100%;top:0;left:0;" frameborder="0"></iframe>
            </body>
            </html>
        "#))
        .or(warp::path("about").map(|| {
            warp::reply::html(r#"
                <html>
                <head><title>About</title></head>
                <body>
                <h1>Про проект</h1>
                <p>Це простий веб-сервер на Rust, який запускає ACWA.exe.</p>
                </body>
                </html>
            "#)
        }));

    tokio::spawn(async move {
        warp::serve(route).run(([127, 0, 0, 1], 54685)).await;
    });

    if let Err(e) = child.wait() {
        eprintln!("Помилка при завершенні ACWA.exe: {}", e);
    }

    println!("ACWA.exe завершено.");
}
