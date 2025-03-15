use std::env;
use rand::Rng;

const MAXN: usize = 2000;

fn initialize_inputs(n: usize) -> (Vec<Vec<f32>>, Vec<f32>) {
    let mut rng = rand::thread_rng();
    let mut a = vec![vec![0.0; n]; n];
    let mut b = vec![0.0; n];

    for i in 0..n {
        for j in 0..n {
            a[i][j] = rng.gen::<f32>();
        }
        b[i] = rng.gen::<f32>();
    }

    (a, b)
}

fn gauss(n: usize, mut a: Vec<Vec<f32>>, mut b: Vec<f32>) -> Vec<f32> {
    let mut x = vec![0.0; n];

    for norm in 0..n-1 {
        for row in norm+1..n {
            let multiplier = a[row][norm] / a[norm][norm];
            for col in norm..n {
                a[row][col] -= a[norm][col] * multiplier;
            }
            b[row] -= b[norm] * multiplier;
        }
    }

    for row in (0..n).rev() {
        x[row] = b[row];
        for col in row+1..n {
            x[row] -= a[row][col] * x[col];
        }
        x[row] /= a[row][row];
    }

    x
}

fn main() {
    let args: Vec<String> = env::args().collect();
    if args.len() < 2 {
        eprintln!("Uso: {} <tamanho da matriz>", args[0]);
        return;
    }

    let n: usize = args[1].parse().expect("Por favor, insira um número válido");
    if n < 1 || n > MAXN {
        eprintln!("Tamanho inválido. Escolha entre 1 e {}", MAXN);
        return;
    }

    let (a, b) = initialize_inputs(n);
    let x = gauss(n, a, b);

    if n < 100 {
        println!("X = {:?}", x);
    }
}