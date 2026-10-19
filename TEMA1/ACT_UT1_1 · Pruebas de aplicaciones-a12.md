### Ejercicio 3.3.1

**Comandos utilizados:**
sudo apt install nasm gdb strace build-essential
mkdir a12
nano hola_enunciado.asm
nasm -f elf64 hola_enunciado.asm -o hola_enunciado.o
ld hola_enunciado.o -o hola_enunciado
./hola_enunciado
echo $?

**Salida obtenida:**
Hello World!
Violación de segmento (`core' generado)

**Código de salida:**
139


### Ejercicio 3.3.2

**Comandos utilizados:**
cp hola_enunciado.asm hola.asm
nano hola.asm
nasm -f elf64 hola.asm -o hola.o
ld hola.o -o hola
./hola
echo $?

**Salida obtenida:**
Hello World!

**Código de salida:**
0


### Ejercicio 3.3.3


**Comandos utilizados:**


**Análisis de la ejecución:**
