fn reverse(pair: (i32, bool)) -> (bool, i32) {
    
    let (boolean, integer) = pair;

    return (integer, boolean)
}

#[fmt::Display]
#[derive(Debug)]
struct Matrix(f32, f32, f32, f32);


fn main() {


    let long_tuple = (1u8, 2u16, 3u32,
                      3i32, -6i8,
                      0.1f32, 0.2f64,
                      'a', "hello", true);

    println!("long_tuple first: {}", long_tuple.1);
    println!("long_tuple second: {}", long_tuple.2);
    
    // println!("{:?}", long_tuple);

    // Tuples can be tuple members
    let tuple_of_tuples = ((1u8, 2u16, 3u32), (4u64, -1i32), (true, false));
    println!("tuple of tuples: {:?}", tuple_of_tuples);

    let pair = (1, true);
    println!("pair is {:?}", pair);

    // reverse tuple
    println!("reverse tuple: {:?}", reverse(pair));

    // too long tuple can't be print
    // let too_long_tuple = (1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13);
    // println!("{}", too_long_tuple);

    let tuple = (1, "hello", 4.5, true);
    
    let (a, b, c, d) = tuple;
    println!("{:?}, {:?}, {:?}, {:?}", a, b, c, d);

    let matrix = Matrix(1.1, 1.2, 1.3, 1.4);
    println!("{:?}", matrix);
    println!("{}", matrix);

}
