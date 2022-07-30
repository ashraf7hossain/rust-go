
fn read<T>()->Vec<T> 
where 
    T: std::str::FromStr,
    <T as std::str::FromStr>::Err : std::fmt::Debug,
{
    let mut inpt = String::new();
    std::io::stdin().read_line(&mut inpt).unwrap();
    let v : Vec<T> = inpt.trim().split(" ").map(|x| x.parse().expect("NO")).collect();
    return v;
}


fn main() {
   let z = read::<String>()[0].clone();
   let x: i32 = read()[0];
   let y: i64 = read()[0];
   println!("Sum of x and y = {:?} ",z);
   println!("x = {}",x);
   println!("y = {}",y);
}
