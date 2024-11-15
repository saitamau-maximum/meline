import { useQuery, useQueryClient } from "@tanstack/react-query";
import { MessageRepositoryImpl } from "@/repositories/message";
import { useEffect, useMemo } from "react";
import { ChatRepositoryImpl } from "@/repositories/chat";
import { logger } from "@/utils/logger";

interface UseMessagesOptions {
  channelId: number;
}

export const useMessages = ({ channelId }: UseMessagesOptions) => {
  const messageRepository = useMemo(() => new MessageRepositoryImpl(), []);

  const messageQuery = useQuery({
    queryKey: messageRepository.getMessages$$key(channelId),
    queryFn: () => messageRepository.getMessages(channelId),
  });
  const queryClient = useQueryClient();

  useEffect(() => {
    logger.raw(`=== useMessages ===`);
    const chatRepository = new ChatRepositoryImpl();
    chatRepository.connect(channelId);
    logger.info(`Connecting to chat channel ${channelId}`);
    chatRepository.onMessageReceived(() => {
      queryClient.invalidateQueries({
        queryKey: messageRepository.getMessages$$key(channelId),
      });
    });

    return () => {
      chatRepository.disconnect();
      logger.info(`Disconnecting from chat channel ${channelId}`);
    };
  }, [channelId, queryClient, messageRepository]);

  return {
    messages: messageQuery.data?.messages ?? [],
  };
};
