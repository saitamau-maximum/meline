import { AuthUserRepositoryImpl } from "./auth-user";
import { ChannelRepositoryImpl } from "./channel";
import { ChatRepositoryImpl } from "./chat";
import { MessageRepositoryImpl } from "./message";

export const DefaultRepositories = {
  channelRepository: new ChannelRepositoryImpl(),
  chatRepository: new ChatRepositoryImpl(),
  authUserRepository: new AuthUserRepositoryImpl(),
  messageRepository: new MessageRepositoryImpl(),
};
