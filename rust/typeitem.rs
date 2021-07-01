fn main() {

    let logical: bool = true;

    let a_float: f64 = 1.0;
    let an_integer   = 3i32;

    println!("{}", logical);
    println!("{}", a_float);
    println!("{}", an_integer);

    let mut inferred_type = 111;
    println!("{}", inferred_type); // 如果111这个值没有使用过，则会发出一个waring

    inferred_type = 222;
    println!("{}", inferred_type);

    // inferred_type = true; // 不能直接覆盖不同类型的值
    
    let inferred_type = true; // 可以通过再次声明的方式覆盖
    println!("{}", inferred_type);


}
