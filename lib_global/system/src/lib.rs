use std::ffi::CString;
use std::os::raw::c_char;
use sysinfo::{System, SystemExt};

#[no_mangle]
pub extern "C" fn get_os_name() -> *mut c_char {
    let mut system = System::new_all();
    system.refresh_all();
    
    let os_name = system.name().unwrap_or_else(|| "Unknown".to_string());
    
    // Перетворюємо Rust String у C-рядок
    let c_string = CString::new(os_name).unwrap();
    c_string.into_raw()
}

#[no_mangle]
pub extern "C" fn get_cpu_cores() -> usize {
    let mut system = System::new_all();
    system.refresh_all();
    
    // Повертаємо кількість ядер процесора
    system.physical_core_count().unwrap_or(0)
}

#[no_mangle]
pub extern "C" fn get_total_memory() -> u64 {
    let mut system = System::new_all();
    system.refresh_all();
    
    // Повертаємо загальну пам'ять у мегабайтах
    system.total_memory() / 1024
}

#[no_mangle]
pub extern "C" fn get_available_memory() -> u64 {
    let mut system = System::new_all();
    system.refresh_all();
    
    // Повертаємо доступну пам'ять у мегабайтах
    system.available_memory() / 1024
}
