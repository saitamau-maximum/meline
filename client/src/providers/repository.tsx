import { DefaultRepositories } from "@/repositories";
import { IAuthUserRepository } from "@/repositories/auth-user";
import { IChannelRepository } from "@/repositories/channel";
import { IChatRepository } from "@/repositories/chat";
import { IMessageRepository } from "@/repositories/message";
import { INotificationRepository } from "@/repositories/notification";
import React, { createContext } from "react";

export interface RepositoryContextProps {
  channelRepository: IChannelRepository;
  chatRepository: IChatRepository;
  authUserRepository: IAuthUserRepository;
  messageRepository: IMessageRepository;
  notificationRepository: INotificationRepository;
}

export const RepositoryContext =
  createContext<RepositoryContextProps>(DefaultRepositories);

interface LoadingProviderProps {
  children: React.ReactNode;
}

export const RepositoryProvider = ({ children }: LoadingProviderProps) => {
  return (
    <RepositoryContext.Provider value={DefaultRepositories}>
      {children}
    </RepositoryContext.Provider>
  );
};
