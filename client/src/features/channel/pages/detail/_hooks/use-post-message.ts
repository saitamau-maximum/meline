import { useMutation } from "@tanstack/react-query";
import { MessageRepositoryImpl } from "@/repositories/message";

interface UsePostMessageOptions {
  channelId: number;
  onCreated?: () => void;
  onFailed?: () => void;
}

interface MutationParam {
  content: string;
}

export const usePostMessage = ({
  channelId,
  onCreated,
  onFailed,
}: UsePostMessageOptions) => {
  return useMutation({
    mutationFn: async (param: MutationParam) => {
      return MessageRepositoryImpl.createMessage(channelId, param);
    },
    onSuccess: () => {
      onCreated?.();
    },
    onError: () => {
      onFailed?.();
    },
  });
};
