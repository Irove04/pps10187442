### Ejercicio 3.3.1

**Comandos utilizados:**
1. sudo apt install nasm gdb strace build-essential
2. mkdir a12
3. nano hola_enunciado.asm
4. nasm -f elf64 hola_enunciado.asm -o hola_enunciado.o
5. ld hola_enunciado.o -o hola_enunciado
6. ./hola_enunciado
7. echo $?

**Salida obtenida:**
Hello World!
Violación de segmento (`core' generado)

**Código de salida:**
139


### Ejercicio 3.3.2

**Comandos utilizados:**
1. cp hola_enunciado.asm hola.asm
2. nano hola.asm
3. nasm -f elf64 hola.asm -o hola.o
4. ld hola.o -o hola
5. ./hola
6. echo $?

**Salida obtenida:**
Hello World!

**Código de salida:**
0


### Ejercicio 3.3.3


**Comandos utilizados:**


**Análisis de la ejecución:**
