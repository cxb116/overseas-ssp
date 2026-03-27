pub fn add_one(x: i16) -> i16 {
    x + 1
}

#[cfg(test)]
mod tests {
    use super::*;
    
    #[test]
    fn it_works() {
        assert_eq!(3, add_one(2));
    }
}