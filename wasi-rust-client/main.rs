fn main()
{
    // println!("Hello, world!");
}

#[no_mangle]
pub extern "C" fn sum(x:i32, y:i32) -> i32
{
    return x + y;
}
