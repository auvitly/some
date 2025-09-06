gopher template exec [name] --path=internal/config/config.go --object=Config
gopher template exec [name] --path=internal/config/config.go --object=*
gopher template show [name] --go-package --json
gopher template list 

gopher project init [name] --vscode --idea
gopher project show

gopher project layout create [name] --description=... --depend-on=layout
gopher project layout show [name]
gopher project layout list 
gopher project layout delete [name]

gopher project component create [name] --layout=service
gopher project component show [name]
gopher project component list --layout=service
gopher project component delete [name]

gopher template exec [name] 

# темплейты
## функции должены иметь
$args и $returns

## базовые параметры
## базовые импорты
## прочие импорты

string -> float
string -> int
string -> bytes
int hex
int oct
int bin
int dec

$maps.New
$maps.Del
$maps.Set
$maps.Get

$slices.New
$slices.Sort
$slices.Unique