use std::io;
 
fn read_int()->Vec<i32> {
    let mut inpt = String::new();
    io::stdin().read_line(&mut inpt).unwrap();
    let mut v : Vec<i32> = inpt.trim().split(" ").map(|x| x.parse().expect("NO")).collect();
    return v;
}
fn read_i64()->Vec<i64> {
    let mut inpt = String::new();
    io::stdin().read_line(&mut inpt).unwrap();
    let mut v : Vec<i64> = inpt.trim().split(" ").map(|x| x.parse().expect("NO")).collect();
    return v;
}

fn main(){
    let n = read_int()[0];
    
}