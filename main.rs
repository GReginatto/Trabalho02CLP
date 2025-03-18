use std::env;
use std::io::{self, Write};
use std::time::Instant;
use rand::Rng;

const TAMANHO_MAXIMO: usize = 2000;

fn inicializar_matriz_vetor(tamanho: usize) -> (Vec<f32>, Vec<f32>) {
    let mut rng = rand::thread_rng();
    let mut matriz = vec![0.0; tamanho * tamanho];
    let mut vetor_independente = vec![0.0; tamanho];

    for i in 0..tamanho {
        for j in 0..tamanho {
            matriz[i * tamanho + j] = rng.gen::<f32>();
        }
        vetor_independente[i] = rng.gen::<f32>();
    }

    (matriz, vetor_independente)
}

fn eliminacao_gaussiana(tamanho: usize, mut matriz: Vec<f32>, mut vetor_independente: Vec<f32>) -> Vec<f32> {
    let mut vetor_resultado = vec![0.0; tamanho];

    for norma in 0..tamanho - 1 {
        for linha in norma + 1..tamanho {
            let fator = matriz[linha * tamanho + norma] / matriz[norma * tamanho + norma];
            for coluna in norma..tamanho {
                matriz[linha * tamanho + coluna] -= matriz[norma * tamanho + coluna] * fator;
            }
            vetor_independente[linha] -= vetor_independente[norma] * fator;
        }
    }

    for linha in (0..tamanho).rev() {
        vetor_resultado[linha] = vetor_independente[linha];
        for coluna in linha + 1..tamanho {
            vetor_resultado[linha] -= matriz[linha * tamanho + coluna] * vetor_resultado[coluna];
        }
        vetor_resultado[linha] /= matriz[linha * tamanho + linha];
    }

    vetor_resultado
}

fn main() {
    let argumentos: Vec<String> = env::args().collect();
    if argumentos.len() < 2 {
        eprintln!("Uso: {} <tamanho da matriz>", argumentos[0]);
        return;
    }

    let tamanho: usize = argumentos[1].parse().expect("Por favor, insira um número válido");
    if tamanho < 1 || tamanho > TAMANHO_MAXIMO {
        eprintln!("Tamanho inválido. Escolha entre 1 e {}", TAMANHO_MAXIMO);
        return;
    }

    let inicio = Instant::now(); 

    let (matriz, vetor_independente) = inicializar_matriz_vetor(tamanho);
    println!("Matriz e vetor independente inicializados!");
    io::stdout().flush().unwrap();

    let vetor_resultado = eliminacao_gaussiana(tamanho, matriz, vetor_independente);
    println!("Cálculo concluído! Agora imprimindo...");
    io::stdout().flush().unwrap();

    let duracao = inicio.elapsed(); 
    println!("Cálculo concluído em {:.2} segundos", duracao.as_secs_f32());

    
    if tamanho <= 100 {
        println!("Vetor resultado (parcial): {:?}", &vetor_resultado[0..10.min(tamanho)]);
    } else {
        println!("Cálculo finalizado para matriz {}x{}. Primeiros 10 valores do resultado:", tamanho, tamanho);
        println!("{:?}", &vetor_resultado[0..10]);
    }

    io::stdout().flush().unwrap();
}
