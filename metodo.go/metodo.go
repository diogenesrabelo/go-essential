package main

type user struct {
	name  string
	idade uint8
}

func (u user) salvar() {

}

func (u *user) fezAniversario() {
	u.idade++
}

func main() {

}
