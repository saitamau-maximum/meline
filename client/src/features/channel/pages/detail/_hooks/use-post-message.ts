import { useMutation } from "@tanstack/react-query";
import { MessageRepositoryImpl } from "@/repositories/message";
import { useMemo } from "react";

interface UsePostMessageOptions {
  channelId: number;
}

interface MutationParam {
  content: string;
}

export const usePostMessage = ({ channelId }: UsePostMessageOptions) => {
  const messageRepository = useMemo(() => new MessageRepositoryImpl(), []);

  return useMutation({
    mutationFn: async (param: MutationParam) => {
      return messageRepository.createMessage(channelId, param);
    },
  });
};
