extern crate rand;

use rand::Rng;
//use std::fmt::LowerHex;

fn main() {
	let mut d = [0u8; 8];

	rand::thread_rng().fill_bytes(&mut d);

    for u in d.iter() {
    	print!("{:x}", u)
    }
    println!("")
}
