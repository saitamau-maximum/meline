import { useQuery, useQueryClient } from "@tanstack/react-query";
import { useEffect } from "react";
import { logger } from "@/utils/logger";
import { useRepositories } from "@/hooks/repository";

interface UseMessagesOptions {
  channelId: number;
}

export const useMessages = ({ channelId }: UseMessagesOptions) => {
  const { messageRepository, chatRepository } = useRepositories();

  const messageQuery = useQuery({
    queryKey: messageRepository.getMessages$$key(channelId),
    queryFn: () => messageRepository.getMessages(channelId),
  });
  const queryClient = useQueryClient();

  useEffect(() => {
    logger.raw(`=== useMessages ===`);
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
  }, [channelId, queryClient, chatRepository, messageRepository]);

  return {
    messages: messageQuery.data?.messages ?? [],
  };
};
