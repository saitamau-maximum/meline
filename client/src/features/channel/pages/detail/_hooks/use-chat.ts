import { useEffect } from "react";
import { ChatRepositoryImpl, MessageResponse } from "@/repositories/chat";

interface UseChatOptions {
  channelId: number;
  onMessageReceived?: (message: MessageResponse) => void;
}

export const useChat = ({ channelId, onMessageReceived }: UseChatOptions) => {
  const chatRepository = new ChatRepositoryImpl(channelId);

  useEffect(() => {
    chatRepository.connect();
    chatRepository.onMessageReceived((message) => {
      console.log("message", message);
      onMessageReceived?.(message);
    });

    return () => {
      chatRepository.disconnect();
    };
  }, []);
};
