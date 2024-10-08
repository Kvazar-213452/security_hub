use warp::Filter;

pub fn index_route() -> impl Filter<Extract = impl warp::Reply, Error = warp::Rejection> + Clone {
    warp::path::end()
        .map(|| warp::reply::html(r#"
            <h1>Core security_hub!</h1>
        "#))
}