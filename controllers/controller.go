package controllers

import (
	"chg-gera-script-brad/database"
	"chg-gera-script-brad/models"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var firewalls []models.Firewall
	database.DB.Find(&firewalls)
	c.HTML(http.StatusOK, "index.html", firewalls)
}

func ArquivoGerado(c *gin.Context) {
	chg := models.Chg{}
	chg.Nome = c.PostForm("nome")
	chg.Rdm = c.PostForm("rdm")
	firewall := c.PostForm("firewall")
	var firewallSearch models.Firewall
	database.DB.Where("nome = ?", firewall).First(&firewallSearch)
	chg.Firewalls = append(chg.Firewalls, firewallSearch)
	chg.NumeroDoTicket = c.PostForm("numeroDoTicket")

	chg.DataChg = time.Now()

	content := generateFileContent(chg)
	var nomeDoArquivo string
	nomeDoArquivo = "SCRIPT_IMPLANTAÇÃO_E_ROLLBACK_REGRAS_"
	nomeDoArquivo2 := "SCRIPT_IMPLANTAÇÃO_E_ROLLBACK_REGRAS_"
	for _, fw := range chg.Firewalls {
		nomeDoArquivo += fw.Nome
		nomeDoArquivo2 += fw.Nome
	}
	// Cria o arquivo
	nomeDoArquivo += "_" + chg.Rdm
	nomeDoArquivo += "-" + chg.NumeroDoTicket
	nomeDoArquivo += ".txt"
	filePath := "archives/" + nomeDoArquivo
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Erro ao criar o arquivo:", err)
		return
	}
	defer file.Close()

	// Escreve o conteúdo no arquivo
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo:", err)
		return
	}
	c.HTML(http.StatusOK, "arquivo-gerado.html", gin.H{
		"FileName": filePath,
	})
}

func DownloadArquivo(c *gin.Context) {
	// Obtenha o nome do arquivo a ser baixado do parâmetro da rota
	nomeDoArquivo := c.Query("nomeDoArquivo")

	// Decodificar o nome do arquivo

	filePath := nomeDoArquivo
	fmt.Println("nome do arquivo decodificado: " + filePath)
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename="+nomeDoArquivo)
	c.Header("Content-Type", "application/octet-stream")

	c.File(filePath)

}

func generateFileContent(chg models.Chg) string {
	var nomeFirewall []string
	var ipManager []string
	for _, fw := range chg.Firewalls {
		nomeFirewall = append(nomeFirewall, fw.Nome)
		ipManager = append(ipManager, fw.ManagerIp)
	}
	nome := nomeFirewall[0]
	managerIp := ipManager[0]

	return "! Nome: " + chg.Nome + "\n!RDM: " + chg.Rdm + "\n!Data: " + chg.DataChg.Format("02/01/2006") + "\n\n" +
		"##############################################\n" +
		"#\n" +
		"# CMA PRIMÁRIA - " + managerIp + "\n" +
		"#\n" +
		"##############################################\n\n" +
		"*************************************\n" +
		"****** I M P L A N T A T I O N ******\n" +
		"*************************************\n\n" +
		"01 - Salvar o Database Version.\n" +
		"02 - Aplicar política no firewall da " + nome + "(Manager " + managerIp + ")\n\n" +
		"**************************************************************************************\n" +
		"			>>>>>>>>>> RETORNO <<<<<<<<<<\n" +
		"**************************************************************************************\n\n" +
		"##############################################\n" +
		"#\n" +
		"# CMA PRIMÁRIA - " + managerIp + "\n" +
		"#\n" +
		"##############################################\n\n" +
		"1 - Acessa a Dashboard do Firewall CheckPoint R80.40\n" +
		"02 - clicar em Installation History\n" +
		"03 - Selecionar a Policy Installations (Procurar a última política aplicada)\n" +
		"04 - clicar em Install Specific Version\n"
}
