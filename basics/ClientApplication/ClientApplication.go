import(
	"MiddlewareImplementation/basics/ChatProxy"
	"MiddlewareImplementation/basics/NamingProxy"
	"bufio"
	"fmt"
)
func main() {
	reader := bufio.NewReader(os.Stdin);
	fmt.Println("Bem vindo!");
	fmt.Println("Digite o seu nome");
	user_name, err := reader.ReadString('\n');
	if err!=nil{
		fmt.Println("erro ao tentar ler a entrada");
	}
	fmt.Println("este é um chat com transferência de arquivos");
	fmt.Println("para enviar mensagens basta digitá-las e apertar enter");
	fmt.Println("escreva \":f \" + nome do arquivo para enviar e \":d \" + nome do arquivo para baixar e em seguida aperte enter");
	fmt.Println("escreva \":l \" e aperte enter para ver os arquivos disponiveis");
	var nameServer NamingProxy.NameServiceObject{Aor: "127.0.0.1:1024:0"};
	var chat ChatProxy.Chat{Aor: nameServer.lookup("Chat")};
	for{
		//é preciso definir que tipo é esse chat e colocar como tipo que chama nas funcões Listen e Send
		msgs := chat.Listen();
		for i:=0; i<len(msgs); i++ {
			fmt.Println(msgs[i]);		
		}
		msg, err := reader.ReadString('\n');
		if err!=nil{
			fmt.Println("erro ao tentar ler a entrada");
		}
		resp := chat.Send(msg);
		if(msg[:3]==":d "){
			//tratar download de arquivo
		}else{
			fmt.Println(resp);
		}
}