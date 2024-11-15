import { useQuery, useQueryClient } from "@tanstack/react-query";
import { MessageRepositoryImpl } from "@/repositories/message";
import { useEffect } from "react";
import { ChatRepositoryImpl } from "@/repositories/chat";
import { logger } from "@/utils/logger";

interface UseMessagesOptions {
  channelId: number;
}

export const useMessages = ({ channelId }: UseMessagesOptions) => {
  const messageQuery = useQuery({
    queryKey: MessageRepositoryImpl.getMessages$$key(channelId),
    queryFn: () => MessageRepositoryImpl.getMessages(channelId),
  });
  const queryClient = useQueryClient();

  useEffect(() => {
    logger.raw(`=== useMessages ===`);
    const chatRepository = new ChatRepositoryImpl(channelId);
    chatRepository.connect();
    logger.info(`Connecting to chat channel ${channelId}`);
    chatRepository.onMessageReceived(() => {
      queryClient.invalidateQueries({
        queryKey: MessageRepositoryImpl.getMessages$$key(channelId),
      });
    });

    return () => {
      chatRepository.disconnect();
      logger.info(`Disconnecting from chat channel ${channelId}`);
    };
  }, [channelId, queryClient]);

  return {
    messages: messageQuery.data?.messages ?? [],
  };
};
