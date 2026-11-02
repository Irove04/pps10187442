global _start

section .data
message: db "Hello World!", 10

section .text
_start:
    ; Llamada al sistema sys_write
    mov rax, 1
    mov rdi, 1
    mov rsi, message
    mov rdx, 13
    syscall

    ; Llamada al sistema sys_exit
    mov rax, 60
    xor rdi, rdi
    syscall
;
