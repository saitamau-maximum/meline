import { useMutation } from "@tanstack/react-query";
import { MessageRepositoryImpl } from "@/repositories/message";

interface UsePostMessageOptions {
  channelId: number;
}

interface MutationParam {
  content: string;
}

export const usePostMessage = ({ channelId }: UsePostMessageOptions) => {
  return useMutation({
    mutationFn: async (param: MutationParam) => {
      return MessageRepositoryImpl.createMessage(channelId, param);
    },
  });
};
