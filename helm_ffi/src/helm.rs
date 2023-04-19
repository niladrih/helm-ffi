use std::ffi::{CStr, CString};
use std::os::raw::c_char;
use std::io::{self, Write};

extern "C" {
    fn CreateHelm(out: &mut dyn Write, args: &[String]) -> *const c_char;
}

pub fn make_helm() {
    let mut out = io::stdout();
    let result = unsafe { CreateHelm(&mut out, &["sample".to_string()]) };
    let c_str = unsafe { CStr::from_ptr(result) };
    println!("{:?}", c_str.to_str())
}