# checkCpfGo

Command line tool to verify physical person registers status on Brazilian entities.<br>
A pure GO (Golang) solution that manages HTTP requests.

The binary receives a CPF (Brazilian Physical Person Register) to be checked.<br>
Then the CPF is validated on SERPRO (Brazilian Federal Data Processing Service) that uses a RESTful API.<br>

The HTTP request implementation uses GO native libraries.

## Quick Start (on Linux)

### Install GO compiler
```bash
sudo apt install gccgo-go
```
This application was tested with GO release `1.12.2`. To check the installed version of GO compiler, run:
```bash
go version
```

### Build and run checkCpfGo (from the project root folder)
```bash
go build
./checkCpfGo 123456789
```
Assuming that *123456789* is a CPF to be checked.

## Usage

Simply call the binary passing a CPF as parameter, the verdict will be printed on the command line.

### Success example

```
CPF:      40442820135
SUCCESS:  True
MESSAGE:  Regular CPF (OK)

```

### Error example

```
CPF:      11334739706
SUCCESS:  False
MESSAGE:  CPF not regular or not existant

```

## Testing the Binary

### RUN Unit tests (from the project root folder)

The file `app_test.go` contains a collection of unit tests. Run the tests by calling the command below.

``` bash
go test -v
```

### Testing manually

#### Call the binary passing a test CPF as parameter:

```bash
./checkCpfGo 40442820135
```

#### List of testing CPFs provided by SERPRO:

The application uses a service provided by SERPRO, so any test depends on this service being online.<br>
The CPFs provided by SERPRO for testing are:

```
40442820135: Regular (Register OK)
63017285995: Regular (Register OK)
91708635203: Regular (Register OK)
58136053391: Regular (Register OK)
40532176871: Suspended (Problem with the registry)
47123586964: Suspended (Problem with the registry)
07691852312: Regularization Pending (Problem with the registry)
10975384600: Regularization Pending (Problem with the registry)
01648527949: Canceled by Multiplicity (Problem with the registry)
47893062592: Canceled by Multiplicity (Problem with the registry)
98302514705: Null (Problem with the registry)
18025346790: Null (Problem with the registry)
64913872591: Registration Canceled (Problem with the registry)
52389071686: Registration Canceled (Problem with the registry)
05137518743: Deceased Holder (Problem with the registry)
08849979878: Deceased Holder (Problem with the registry)
```

## Implementation

### File structure

The application is composed by a set of source code files:

**app.go** : Main file, with `main()` function.<br>
**app_test.go** : File containing all application tests.

**middleware_model.go** : SERPRO service structures.<br>
**middleware.go** : Core implementation of the REQUEST to SERPRO service.

**README.md** : This file, explaining how the application works.<br>
**.gitignore** : List of non-versioned files (such as the compiled binary).

After the compilation process, the non-versioned binary file `checkCpfGo` will be generated.

### API Credentials

The application is configured with a personal TOKEN, please change it.<br>
Its possible to generate a new TOKEN on SERPRO service page: [servicos.serpro.gov.br](https://servicos.serpro.gov.br/inteligencia-de-negocios-serpro/biblioteca/consulta-cpf/teste.html).

The in use TOKEN is specified on the top of `middleware.go` file.<br>
If going to production the SERPRO service URL must also be changed (at the same file).

## Application Info

### Author

Heitor Peralles<br>
[heitorgp@gmail.com](mailto:heitorgp@gmail.com)<br>
[linkedin.com/in/heitorperalles](https://www.linkedin.com/in/heitorperalles)

### Source

[github.com/heitorperalles/checkCpfGo](https://www.github.com/heitorperalles/checkCpfGo)

### License

MIT licensed. See the **LICENSE** file for details.
