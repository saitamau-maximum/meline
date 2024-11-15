import { useMutation } from "@tanstack/react-query";
import { useRepositories } from "@/hooks/repository";

interface UsePostMessageOptions {
  channelId: number;
}

interface MutationParam {
  content: string;
}

export const usePostMessage = ({ channelId }: UsePostMessageOptions) => {
  const { messageRepository } = useRepositories();

  return useMutation({
    mutationFn: async (param: MutationParam) => {
      return messageRepository.createMessage(channelId, param);
    },
  });
};
