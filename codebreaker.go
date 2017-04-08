package codebreaker


type CodeBreaker struct {
  secret string
}

func evaluar(codeBreaker CodeBreaker, codigo string) string {
  retorno := ""
  size := len(codigo)
  for i := 0; i < size; i++ {
		if codigo[i] == codeBreaker.secret[i]{
      retorno = retorno+"X"
    }
	}
  return retorno


}
