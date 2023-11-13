package handlers

func (*BotHandlers) help() string {
	return "/help - list commands\n" +
		"/list <?page> <?limit> <?order> <?direction> - list data\n" +
		"/get <id> - get player\n" +
		"/create <name> <club> <?games> <?goals> <?assists> - add player\n" +
		"/update <id> <name> <?club> <?games> <?goals> <?assists> - update player\n" +
		"/delete <id> - delete player"
}
