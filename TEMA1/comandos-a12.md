sudo apt install nasm gdb strace build-essential
mkdir a12
nano hola_enunciado.asm

; x86_64

global _start

section .data
message: db "Hello World!", 10

section .text
_start:
    mov rax, 1
    mov rdi, 1
    mov rsi, message
    mov rdx, 14
    syscall

    mov rax, 60
    xor rdi, rdi
    syscall ;

    nasm -f elf64 hola_enunciado.asm -o hola_enunciado.o

    ld hola_enunciado.o -o hola_enunciado

    ./hola_enunciado

    echo $?
