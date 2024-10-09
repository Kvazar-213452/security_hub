use warp::Filter;
use sysinfo::{System, SystemExt, ComponentExt};

mod func;
mod config;
mod server;

#[tokio::main]
async fn main() {

    let mut sys = System::new_all();

    // Оновлення інформації про систему
    sys.refresh_all();

    // Отримання назви операційної системи
    if let Some(os_name) = sys.name() {
        println!("Операційна система: {}", os_name);
    }

    // Отримання кількості оперативної пам'яті
    let total_memory = sys.total_memory();
    println!("Загальна оперативна пам'ять: {} KB", total_memory);

    let used_memory = sys.used_memory();
    println!("Використана оперативна пам'ять: {} KB", used_memory);

    // Отримання кількості ядер CPU
    let cpu_cores = sys.cpus().len();
    println!("Кількість ядер процесора: {}", cpu_cores);

    // Інформація про температуру процесора
    for component in sys.components() {
        // Використовуємо temperature() без Option
        let temp = component.temperature();
        println!("Температура {}: {}°C", component.label(), temp);
    }


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
    let change_config_post = server::post::change_config_post();

    let routes = index_route.or(logs_post).or(config_post).or(change_config_post);

    tokio::spawn(async move {
        warp::serve(routes).run(([127, 0, 0, 1], port as u16)).await;
    });

    if let Err(e) = child.wait() {
        eprintln!("Помилка при завершенні ACWA.exe: {}", e);
    }

    println!("ACWA.exe завершено.");
    println!("Отриманий вільний порт: {}", port);
}