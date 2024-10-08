use serde_json::json;
use warp::Filter;
use warp::http::StatusCode;

pub fn logs_post() -> impl Filter<Extract = impl warp::Reply, Error = warp::Rejection> + Clone {
    warp::path("get_logs")
        .and(warp::post())
        .map(|| {
            let logs = std::fs::read_to_string("data/main.log").unwrap_or_else(|_| String::from("Не вдалося прочитати файл"));
            let response = json!({ "logs": logs });
            warp::reply::with_status(warp::reply::json(&response), StatusCode::OK)
        })
}
