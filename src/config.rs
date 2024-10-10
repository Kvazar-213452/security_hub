use lazy_static::lazy_static;
use libloading::Library;

// core_web
pub const H_WINDOWS_CORE: &str = "800";
pub const W_WINDOWS_CORE: &str = "1000";
pub const NAME_WINDOWS_CORE: &str = "Security Hub";

pub const SHELL_WEB_CORE: &str = "./shell_web.exe";
pub const START_FILE_CORE: &str = "start_conf.log";

// web_server
pub const CONFIG_WEB: &str = "web/config.toml";

// library
pub const LIBRARY_FOLDER: &str = "library/";
pub const LIBRARY_PORT: &str = "find_free_port.dll";

lazy_static! {
    pub static ref LIBRARY_PORT_LAZY: String = {
        let path = format!("{}{}", LIBRARY_FOLDER, LIBRARY_PORT);
        path
    };

    pub static ref LIBRARY_PORT_PHAT: Library = unsafe {
        Library::new(&*LIBRARY_PORT_LAZY).expect("Не вдалося завантажити бібліотеку")
    };
}

// data
pub const DATA_LOG: &str = "data/main.log";
pub const DATA_CONFIG: &str = "data/main_config.json";