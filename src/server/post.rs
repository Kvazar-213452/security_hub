use serde::{Deserialize, Serialize};
use serde_json::json;
use warp::Filter;
use warp::http::StatusCode;
use std::fs;
use std::collections::HashMap;

mod config;

#[derive(Deserialize)]
struct ConfigChange {
    key: String,
    value: serde_json::Value,
}

pub fn logs_post() -> impl Filter<Extract = impl warp::Reply, Error = warp::Rejection> + Clone {
    warp::path("get_logs")
        .and(warp::post())
        .map(|| {
            let logs = std::fs::read_to_string(config::DATA_LOG).unwrap_or_else(|_| String::from("Не вдалося прочитати файл"));
            let response = json!({ "logs": logs });
            warp::reply::with_status(warp::reply::json(&response), StatusCode::OK)
        })
}

pub fn config_post() -> impl Filter<Extract = impl warp::Reply, Error = warp::Rejection> + Clone {
    warp::path("config")
        .and(warp::post())
        .map(|| {
            let config = std::fs::read_to_string(config::DATA_CONFIG).unwrap_or_else(|_| String::from("Не вдалося прочитати файл"));
            let response = json!({ "config": config });
            warp::reply::with_status(warp::reply::json(&response), StatusCode::OK)
        })
}

pub fn change_config_post() -> impl Filter<Extract = impl warp::Reply, Error = warp::Rejection> + Clone {
    warp::path("change_config")
        .and(warp::post())
        .and(warp::body::json())
        .map(|new_config: ConfigChange| {
            let config_path = config::DATA_CONFIG;
            let mut current_config: HashMap<String, serde_json::Value> = 
                fs::read_to_string(config_path)
                .map(|data| serde_json::from_str(&data).unwrap())
                .unwrap_or_else(|_| {
                    HashMap::new()
                });

            current_config.insert(new_config.key, new_config.value);

            let updated_config = serde_json::to_string(&current_config).unwrap();
            fs::write(config_path, updated_config).expect("Unable to write config file");

            warp::reply::with_status(warp::reply::json(&current_config), StatusCode::OK)
        })
}