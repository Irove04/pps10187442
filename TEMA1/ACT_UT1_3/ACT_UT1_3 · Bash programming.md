# ACT_UT1_3 Bash programming

## 3.3 Script en Bash simple y Debug

Para esta actividad, he creado un script básico, le he dado permisos de ejecución y he utilizado las herramientas de debug de Bash.

**1. Creación del script:**
```bash
touch script1.sh
nano script1.sh
```

## 3.4: Tutorial de Bash

 Basado en el repositorio de [krother/bash_tutorial](https://github.com/krother/bash_tutorial).

### Preparación

Antes de comenzar:

```bash
# Clonar el repositorio de GitHub
git clone [https://github.com/krother/bash_tutorial.git](https://github.com/krother/bash_tutorial.git)

# Entrar a la carpeta de los ejercicios
cd bash_tutorial/exercises/
```

# EJERCICIO 1 BASH

## La primera y segunda letra

```bash
torial/exercises/exercise_1$ ls -a directoryB/solution_1.1/
.  ..  the_first_character_is_A  .the_second_character_is_P
```
## Tercera letra

```bash
alumno@pps10187442:~/Escritorio/pps-entregas/pps10187442/TEMA1/ACT_UT1_3/bash_tutorial/exercises/exercise_1$ cd directoryB
alumno@pps10187442:~/Escritorio/pps-entregas/pps10187442/TEMA1/ACT_UT1_3/bash_tutorial/exercises/exercise_1/directoryB$ source program.sh 
The third character is 'T'.
```
## Cuarta letra E por que es 69 en la tabla ASCII

```bash
alumno@pps10187442:~/Escritorio/pps-entregas/pps10187442/TEMA1/ACT_UT1_3/bash_tutorial/exercises/exercise_1/directoryC$ ls -l
total 4
-rw-rw-r-- 1 alumno alumno 69 nov  8 23:47 text_file.txt

```

# EJERCICIO 2 BASH

## Quinta letra  

```bash
alumno@pps10187442:~/Escritorio/pps-entregas/pps10187442/TEMA1/ACT_UT1_3/bash_tutorial/exercises/exercise_2$ less solution_2.1.txt ^C 

The fifth character of the solution is 'N'.


Did you know:
    The solution is usually well-dressed.
```

## Sexta letra es O por que para guardar un documento es control+O

```bash
nano caracteres.txt
```

# EJERCICIO 3 BASH

## Septima y octava letra es la D y la Y por que es donde  esta el simbolo +

```bash
alumno@pps10187442:~/Escritorio/pps-entregas/pps10187442/TEMA1/ACT_UT1_3/bash_tutorial/exercises/exercise_3$ ls
data  part1  part2
alumno@pps10187442:~/Escritorio/pps-entregas/pps10187442/TEMA1/ACT_UT1_3/bash_tutorial/exercises/exercise_3$ mkdir solution
alumno@pps10187442:~/Escritorio/pps-entregas/pps10187442/TEMA1/ACT_UT1_3/bash_tutorial/exercises/exercise_3$ cp part1/* part2/* solution/
alumno@pps10187442:~/Escritorio/pps-entregas/pps10187442/TEMA1/ACT_UT1_3/bash_tutorial/exercises/exercise_3$ ls -l solution/*
-rw-rw-r-- 1 alumno alumno 0 nov  9 00:22 solution/___+______+_.txt
-rw-rw-r-- 1 alumno alumno 0 nov  9 00:22 solution/UIMDELGKBCYF.txt
alumno@pps10187442:~/Escritorio/pps-entregas/pps10187442/TEMA1/ACT_UT1_3/bash_tutorial/exercises/exercise_3$ 

```
## Novena y décima letra es la T y la E por que son las letras diferentes

```bash
alumno@pps10187442:~/Escritorio/pps-entregas/pps10187442/TEMA1/ACT_UT1_3/bash_tutorial/exercises/exercise_3$ cd data/
alumno@pps10187442:~/Escritorio/pps-entregas/pps10187442/TEMA1/ACT_UT1_3/bash_tutorial/exercises/exercise_3/data$ rm *Y*
alumno@pps10187442:~/Escritorio/pps-entregas/pps10187442/TEMA1/ACT_UT1_3/bash_tutorial/exercises/exercise_3/data$ ls
X_X_X_T_X_X_X  X_X_X_X_X_E_X
alumno@pps10187442:~/Escritorio/pps-entregas/pps10187442/TEMA1/ACT_UT1_3/bash_tutorial/exercises/exercise_3/data$ 

```
# EJERCICIO 4 BASH

## La decimoprimera letra es la S por que falta en "stupid"

```bash
alumno@pps10187442:~/Escritorio/pps-entregas/pps10187442/TEMA1/ACT_UT1_3/bash_tutorial/exercises$ cd exercise_4
alumno@pps10187442:~/Escritorio/pps-entregas/pps10187442/TEMA1/ACT_UT1_3/bash_tutorial/exercises/exercise_4$ diff ai.txt artificial_intelligence.txt 
6c6
< that they're too stupid
---
> that they're too tupid
alumno@pps10187442:~/Escritorio/pps-entregas/pps10187442/TEMA1/ACT_UT1_3/bash_tutorial/exercises/exercise_4$ 
```
## Decimosegunda letra es la P por que la ultima palabra es "parameters"

```bash
alumno@pps10187442:~/Escritorio/pps-entregas/pps10187442/TEMA1/ACT_UT1_3/bash_tutorial/exercises/exercise_4$ sort elephant.txt 

and with five
I can fit an elephant
it will wiggle its trunk.
(John von Neumann)
With four parameters
alumno@pps10187442:~/Escritorio/pps-entregas/pps10187442/TEMA1/ACT_UT1_3/bash_tutorial/exercises/exercise_4$ 
```

## Decimotercera letra es la A por que la frase empieza por "a"

```bash
alumno@pps10187442:~/Escritorio/pps-entregas/pps10187442/TEMA1/ACT_UT1_3/bash_tutorial/exercises/exercise_4$ grep fire datascience.txt 
a fire station
alumno@pps10187442:~/Escritorio/pps-entregas/pps10187442/TEMA1/ACT_UT1_3/bash_tutorial/exercises/exercise_4$ 
```

# EJERCICIO 5 BASH

## Decimocuarta letra es T y la decimoquinta es A

```bash
alumno@pps10187442:~/Escritorio/pps-entregas/pps10187442/TEMA1/ACT_UT1_3/bash_tutorial/exercises$ cd exercise_5
alumno@pps10187442:~/Escritorio/pps-entregas/pps10187442/TEMA1/ACT_UT1_3/bash_tutorial/exercises/exercise_5$ unzip wrapped.zip 
Archive:  wrapped.zip
  inflating: wrapped.tar             
alumno@pps10187442:~/Escritorio/pps-entregas/pps10187442/TEMA1/ACT_UT1_3/bash_tutorial/exercises/exercise_5$ tar -xf wrapped.tar 
alumno@pps10187442:~/Escritorio/pps-entregas/pps10187442/TEMA1/ACT_UT1_3/bash_tutorial/exercises/exercise_5$ gunzip solution/solution.txt.gz 
alumno@pps10187442:~/Escritorio/pps-entregas/pps10187442/TEMA1/ACT_UT1_3/bash_tutorial/exercises/exercise_5$ cat solution/solution.txt 

Congratulations. You unwrapped everything.

The 14th character is: 'T'

The 15th character is: 'A'

alumno@pps10187442:~/Escritorio/pps-entregas/pps10187442/TEMA1/ACT_UT1_3/bash_tutorial/exercises/exercise_5$ 
```

# EJERCICIO 6 BASH

## Decimosexta letra es G y la decimoseptima es O

```bash
alumno@pps10187442:~/Escritorio/pps-entregas/pps10187442/TEMA1/ACT_UT1_3/bash_tutorial/exercises$ cd exercise_6/
alumno@pps10187442:~/Escritorio/pps-entregas/pps10187442/TEMA1/ACT_UT1_3/bash_tutorial/exercises/exercise_6$ chmod a+x permissions.sh
alumno@pps10187442:~/Escritorio/pps-entregas/pps10187442/TEMA1/ACT_UT1_3/bash_tutorial/exercises/exercise_6$ ./permissions.sh 
character 16 is:
G
character 17 is:
O
alumno@pps10187442:~/Escritorio/pps-entregas/pps10187442/TEMA1/ACT_UT1_3/bash_tutorial/exercises/exercise_6$ 
```

## Decimooctava letra es la N por que es la ultima letra del primer autor

```bash
alumno@pps10187442:~/Escritorio/pps-entregas/pps10187442/TEMA1/ACT_UT1_3/bash_tutorial/exercises/exercise_6$ df --version 
df (GNU coreutils) 9.4
Copyright © 2023 Free Software Foundation, Inc.
License GPLv3+: GNU GPL version 3 or later <https://gnu.org/licenses/gpl.html>.
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

Escrito por Torbjörn Granlund, David MacKenzie, y Paul Eggert.
alumno@pps10187442:~/Escritorio/pps-entregas/pps10187442/TEMA1/ACT_UT1_3/bash_tutorial/exercises/exercise_6$ 
```
## Decimonovena letra es I por que es la novena letra del alfabeto

```bash
alumno@pps10187442:~/Escritorio/pps-entregas/pps10187442/TEMA1/ACT_UT1_3/bash_tutorial/exercises/exercise_6$ export GIVEME=SOLUTION
alumno@pps10187442:~/Escritorio/pps-entregas/pps10187442/TEMA1/ACT_UT1_3/bash_tutorial/exercises/exercise_6$ echo $GIVEME | wc -c
9
```

## La vigesima letra es la C por que es la opcion de numero maximo de paquetes

```bash
alumno@pps10187442:~/Escritorio/pps-entregas/pps10187442/TEMA1/ACT_UT1_3/bash_tutorial/exercises/exercise_6$ man ping
```
