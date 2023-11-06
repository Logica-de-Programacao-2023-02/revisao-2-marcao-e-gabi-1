package q4

import "errors"

// Você recebe uma lista de caminhos, onde `caminhos[i] = [cidadeAi, cidadeBi]` significa que existe um caminho direto que
//vai de cidadeAi para cidadeBi. Retorne a cidade de destino, ou seja, a cidade sem nenhum caminho que saia dela.

func Destino(caminhos [][2]string) (string, error) {
	cidadesdeorigem := make(map[string]bool)

	// Marca todas as cidades de partida
	for _, caminho := range caminhos {
		cidadesdeorigem[caminho[0]] = true
	}

	// Encontra a cidade de destino
	for _, caminho := range caminhos {
		cidadededestino := caminho[1]
		if !cidadesdeorigem[cidadededestino] {
			return cidadededestino, nil
		}
	}
	return "", errors.New("cidade de destino não encontrada")
}
