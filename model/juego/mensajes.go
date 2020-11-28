package juego

import (
	"bufio"
	"fmt"
	"os"
)

func pausa() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Presionar [Enter] ")
	reader.ReadString('\n')
}

func mensaje_game_over() {
	fmt.Println(" _____ _____ _____ _____    _____ _____ _____ _____ ")
	fmt.Println("|   __|  _  |     |   __|  |     |  |  |   __| __  |")
	fmt.Println("|  |  |     | | | |   __|  |  |  |  |  |   __|    -|")
	fmt.Println("|_____|__|__|_|_|_|_____|  |_____|\\___/|_____|__|__|")
	fmt.Println("                                                    ")
	pausa()
}

func mensaje_you_win() {
	fmt.Println("____    ____  ______    __    __    ____    __    ____  __  .__   __.        __  ")
	fmt.Println("\\   \\  /   / /  __  \\  |  |  |  |   \\   \\  /  \\  /   / |  | |  \\ |  |       |  | ")
	fmt.Println(" \\   \\/   / |  |  |  | |  |  |  |    \\   \\/    \\/   /  |  | |   \\|  |       |  | ")
	fmt.Println("  \\_    _/  |  |  |  | |  |  |  |     \\            /   |  | |  . `  |       |  | ")
	fmt.Println("    |  |    |  `--'  | |  `--'  |      \\    /\\    /    |  | |  |\\   |       |__| ")
	fmt.Println("    |__|     \\______/   \\______/        \\__/  \\__/     |__| |__| \\__|       (__) ")
	fmt.Println("                                                                                 ")
	pausa()
}

func mensaje_si_pareja() {
	fmt.Println("  _       _____ _                                                   _         _  ")
	fmt.Println(" (_)     / ____(_)                                                 (_)       | | ")
	fmt.Println(" | |    | (___  _       ___  ___  _ __         _ __   __ _ _ __ ___ _  __ _  | | ")
	fmt.Println(" | |     \\___ \\| |     / __|/ _ \\| '_ \\       | '_ \\ / _` | '__/ _ \\ |/ _` | | | ")
	fmt.Println(" | |     ____) | |     \\__ \\ (_) | | | |      | |_) | (_| | | |  __/ | (_| | |_| ")
	fmt.Println(" |_|    |_____/|_|     |___/\\___/|_| |_|      | .__/ \\__,_|_|  \\___| |\\__,_| (_) ")
	fmt.Println("                                              | |                 _/ |           ")
	fmt.Println("                                              |_|                |__/            ")
	pausa()
}

func mensaje_no_son_pareja() {
	fmt.Println("    		                                      _        ")
	fmt.Println("                                                         (_)       ")
	fmt.Println("     _ __   ___    ___  ___  _ __    _ __   __ _ _ __ ___ _  __ _  ")
	fmt.Println("    | '_ \\ / _ \\  / __|/ _ \\| '_ \\  | '_ \\ / _` | '__/ _ \\ |/ _` | ")
	fmt.Println("    | | | | (_) | \\__ \\ (_) | | | | | |_) | (_| | | |  __/ | (_| | ")
	fmt.Println("    |_| |_|\\___/  |___/\\___/|_| |_| | .__/ \\__,_|_|  \\___| |\\__,_| ")
	fmt.Println("                                    | |                 _/ |       ")
	fmt.Println("                                    |_|                |__/        ")
	pausa()
}
