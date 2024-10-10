use libloading::{Symbol};
use std::os::raw::c_int;

use crate::config;

pub fn find_free_port_dll() -> c_int {
    unsafe {
        let func: Symbol<unsafe extern fn() -> c_int> = config::LIBRARY_PORT_PHAT.get(b"FindFreePort")
            .expect("Не вдалося знайти функцію FindFreePort");

        func()
    }
}